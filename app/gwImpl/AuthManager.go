package gwImpl

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oceanho/gw"
	"github.com/oceanho/gwpro/app/dbModel"
	"time"
)

type AuthManager struct {
	store             gw.Store
	permissionManager gw.IPermissionManager
	hash              gw.ICryptoHash
	passwordSigner    gw.IPasswordSigner
	userCachePrefix   string
	expiration        time.Duration
	permPagerExpr     gw.PagerExpr
}

func (a AuthManager) getUserCacheKey(passport string) string {
	return fmt.Sprintf("%s.%s", a.userCachePrefix, passport)
}

func (a AuthManager) Login(passport, secret, credType, verifyCode string) (gw.User, error) {
	var gwUser gw.User
	password := a.passwordSigner.Sign(secret)
	cache := a.store.GetCacheStoreByName("gwPro")
	userCacheKey := a.getUserCacheKey(passport)
	bytes, err := cache.Get(context.Background(), userCacheKey).Bytes()
	if err == nil && len(bytes) > 0 {
		err := json.Unmarshal(bytes, &gwUser)
		return gwUser, err
	}
	var user dbModel.User
	db := a.store.GetDbStoreByName("gwPro")
	err = db.Where("passport = ? and secret = ?", passport, password).First(&user).Error
	if err != nil {
		return gw.EmptyUser, err
	}
	// secret/password checker
	_, perms, err := a.permissionManager.QueryByUser(user.TenantId, user.ID, a.permPagerExpr)
	if err != nil {
		return gw.EmptyUser, err
	}
	gwUser.Id = user.ID
	gwUser.TenantId = user.TenantId
	gwUser.Passport = user.Passport
	gwUser.Permissions = perms
	_ = cache.Set(context.Background(), userCacheKey, gwUser, a.expiration).Err()
	return gwUser, nil
}

func (a AuthManager) Logout(user gw.User) bool {
	return true
}

func DefaultAuthManager(initCtx gw.ServerInitializationContext) AuthManager {
	return AuthManager{
		store:             initCtx.Store,
		hash:              initCtx.Hash,
		userCachePrefix:   "gw-pro-user",
		expiration:        time.Hour * 168,
		permissionManager: initCtx.PermissionManager,
		passwordSigner:    initCtx.PasswordSigner,
		permPagerExpr:     gw.DefaultPagerExpr(2048, 1),
	}
}
