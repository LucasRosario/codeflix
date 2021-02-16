package repository

import (
	"fmt"

	"github.com/LucasRosario/codiflix/domain/model"
	"github.com/jinzhu/gorm"
)

//TransactionRepositoryDb Repositorio de trançaões no banco de dados
type TransactionRepositoryDb struct {
	Db *gorm.DB
}

//Register Registros no BD
func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

//Save Editando transações no BD
func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

//Find Procurando transações no banco de dados
func (t *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}
	return &transaction, nil
}
