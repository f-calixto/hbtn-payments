package repository

import (
	"database/sql"
	"fmt"
	"os"

	"control-pago-backend/internal/entity"
	"control-pago-backend/internal/errors"
	"control-pago-backend/log"
)

type Repository interface {
	RegisterPayment(pmt *entity.RegisterPaymentRequest) error
	StoreReceipt(receipt entity.Receipt) error
	// GetPayments()
}

var (
	registerPaymentQueryWithCompany = `INSERT INTO payments(month, amount, receipt_url, company)
										VALUES($1, $2, $3, $4)`
	registerPaymentQueryWithoutCompany = `INSERT INTO payments(month, amount, receipt_url)
										VALUES($1, $2, $3)`
)

type repository struct {
	logger         log.Logger
	db             *sql.DB
	receiptsFolder string
}

func (r *repository) RegisterPayment(pmt *entity.RegisterPaymentRequest) error {
	if pmt.Company == nil {
		_, err := r.db.Exec(registerPaymentQueryWithoutCompany, pmt.Month, pmt.Amount, pmt.Receipt)
		if err != nil {
			return err
		}
	} else {
		_, err := r.db.Exec(registerPaymentQueryWithCompany, pmt.Month, pmt.Amount, pmt.Receipt, pmt.Company)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) StoreReceipt(receipt entity.Receipt) error {
	f, err := os.Create(fmt.Sprintf("%s/%s", r.receiptsFolder, receipt.Name))
	if err != nil {
		return errors.NewFileError("Error creating file")
	}
	defer f.Close()

	_, err = f.Write(receipt.Data)
	if err != nil {
		return errors.NewFileError("Error writing content to file")
	}

	return nil
}

func NewRepository(lgr log.Logger, db *sql.DB, receiptsFolder string) Repository {
	return &repository{
		logger:         lgr,
		db:             db,
		receiptsFolder: receiptsFolder,
	}
}