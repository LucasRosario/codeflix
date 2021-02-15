package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

//Bank TIPO DE CONTA
type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" gorm:"type:varchar(20)" valdid:"notnull"`
	Name     string     `json:"name" gorm:"type:varchar(255)" valdid:"notnull"`
	Accounts []*Account `gorm:"ForeingKey:BankID" valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return err
	}
	return nil
}

// NewBank NOVA CONTA
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	return &bank, nil
}
