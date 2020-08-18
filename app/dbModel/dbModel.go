package dbModel

import "github.com/oceanho/gw"

func AutoMigrate(store gw.Store) {
	db := store.GetDbStoreByName("gwPro")
	_ = db.AutoMigrate(&User{}, &UserProfile{}, &Role{}, &UserRole{})
}
