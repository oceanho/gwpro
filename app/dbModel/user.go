package dbModel

import (
	"github.com/oceanho/gw/backend/gwdb"
)

type UserCategory uint8

const (
	GenericUser   UserCategory = 1
	TenancyAdmin               = 2
	Administrator              = 3
)

//func (category UserCategory) Scan(value interface{}) error {
//	val, ok := value.(int64)
//	if !ok {
//		return fmt.Errorf("value is not uint8, value: %v", value)
//	}
//	if val < 1 || val > 3 {
//		return fmt.Errorf("value are invalid range(1-3), value: %v", value)
//	}
//	return nil
//}
//
//func (category UserCategory) Value() (driver.Value, error) {
//	return category, nil
//}

type User struct {
	gwdb.Model
	gwdb.HasTenantState
	Category UserCategory `gorm:"not null;default:1"`
	Passport string       `gorm:"type:varchar(32);not null"`
	Secret   string       `gorm:"type:varchar(64);not null"`
	gwdb.HasCreationState
	gwdb.HasModificationState
}

func (*User) TableName() string {
	return "gw_pro_user"
}

type UserProfile struct {
	gwdb.Model
	gwdb.HasTenantState
	UserId uint64
	Email  string `gorm:"type:varchar(32);not null"`
	Phone  string `gorm:"type:varchar(12);not null"`
	gwdb.HasCreationState
	gwdb.HasModificationState
}

func (UserProfile) TableName() string {
	return "gw_pro_user_profile"
}
