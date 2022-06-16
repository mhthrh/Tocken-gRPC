package User

import (
	"GitHub.com/mhthrh/JWT/Server/CryptoUtil"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id          uuid.UUID `json:"id"gorm:"type:uuid;primaryKey;column:Id"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null;column:Name"`
	LastName    string    `json:"lastName" gorm:"type:varchar(255);not null;column:LastName"`
	UserName    string    `json:"userName" gorm:"type:varchar(255);unique;uniqueIndex;not null;column:UserName"`
	PassWord    string    `json:"passWord" gorm:"type:varchar(255);not null;column:Password"`
	Email       string    `json:"email" gorm:"type:varchar(50);not null;column:Email"`
	PhoneNumber string    `json:"phoneNumber" gorm:"type:varchar(50);not null;column:PhoneNumber"`
	IsActive    bool      `json:"isActive" gorm:"type:bool;default:false;not null;column:IsActive"`
	CreatDate   int64     `json:"creatDate"gorm:"autoCreateTime:nano;not null;column:CreateDate"`
}
type Request struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type Response struct {
	ValidTill string `json:"validTill"`
	SignedKey string `json:"signedKey"`
}

type tool struct {
	db  *gorm.DB
	fnc func(string) string
}

func New(db *gorm.DB, f func(string) string) *tool {
	return &tool{db: db, fnc: f}
}

func (t *tool) SignIn(l *Request) (Response, error) {
	var user User
	tx := t.db.Where("\"UserName\" = ? and \"Password\"=? and \"IsActive\"=true", l.Username, t.fnc(l.Password)).First(&user)
	if tx.Error != nil {
		return Response{}, tx.Error
	}
	if tx.RowsAffected == 1 {
		till := time.Now().Add(180 * time.Second).Format(time.UnixDate)
		k := CryptoUtil.NewKey()
		k.Text = fmt.Sprintf("%s#%s", user.UserName, till)
		k.Encrypt()
		return Response{
			ValidTill: till,
			SignedKey: k.Result,
		}, nil
	}
	return Response{}, fmt.Errorf("general error")
}
