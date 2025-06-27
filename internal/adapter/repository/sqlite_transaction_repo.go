package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/PersonalFinanceManagement/cli/adapter/repository/models"
	"github.com/PersonalFinanceManagement/domain/entity"
	"github.com/PersonalFinanceManagement/domain/repository"
	"github.com/aarondl/sqlboiler/v4/boil"
)

// SQLiteTransactionRepository is the concrete implementation for storing transactions in SQLite.
// It acts as a "secondary adapter", plugging into the port defined by the domain.
type SQLiteTransactionRepository struct {
	db *sql.DB
}

// NewSQLiteTransactionRepository creates a new instance of the repository.
func NewSQLiteTransactionRepository(db *sql.DB) *SQLiteTransactionRepository {
	return &SQLiteTransactionRepository{
		db: db,
	}
}

// Ensure our implementation satisfies the interface at compile time.
var _ repository.TransactionRepository = &SQLiteTransactionRepository{}

// Create inserts a new transaction into the database.
func (r *SQLiteTransactionRepository) Create(tx *entity.Transaction) error {
	// Create a SQLBoiler model from your domain entity. This is the translation step.
	model := models.Transaction{
		ID:                   tx.ID,
		Amount:               tx.Amount,
		Pending:              tx.Pending,
		Type:                 string(tx.Type),
		SourceAccountID:      tx.SourceAccountID,
		DestinationAccountID: tx.DestinationAccountID,
		Payee:                tx.Payee,
		CategoryID:           tx.CategoryID,
		Description:          tx.Description,
		MethodOfPayment:      string(tx.MethodOfPayment),
		Created:              tx.Created,
		Updated:              tx.Updated,
	}

	// Use SQLBoiler's generated Insert function for a type-safe insertion.
	// boil.Infer() tells SQLBoiler to use the database's default values for any zero-value fields.
	err := model.Insert(context.Background(), r.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("could not insert transaction: %w", err)
	}
	return nil
}

// GetByID retrieves a transaction by its ID.
func (r *SQLiteTransactionRepository) GetByID(id string) (*entity.Transaction, error) {
	// Use SQLBoiler's generated FindTransaction function.
	model, err := models.FindTransaction(context.Background(), r.db, id)
	if err != nil {
		// Handle the case where the record is not found gracefully.
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction with ID %s not found", id)
		}
		return nil, fmt.Errorf("could not find transaction: %w", err)
	}

	// Translate the SQLBoiler model back to your pure domain entity.
	return &entity.Transaction{
		ID:                   model.ID,
		Amount:               model.Amount,
		Pending:              model.Pending,
		Type:                 entity.TransactionType(model.Type),
		SourceAccountID:      model.SourceAccountID,
		DestinationAccountID: model.DestinationAccountID,
		Payee:                model.Payee,
		CategoryID:           model.CategoryID,
		Description:          model.Description,
		MethodOfPayment:      entity.MethodOfPayment(model.MethodOfPayment),
		Created:              model.Created,
		Updated:              model.Updated,
	}, nil
}

// Update is a placeholder for now.
// A real implementation would convert the entity to a model and call model.Update().
func (r *SQLiteTransactionRepository) Update(transaction *entity.Transaction) error {
	// TODO: Implement the update logic.
	return fmt.Errorf("update not implemented")
}

// Delete is a placeholder for now.
// A real implementation would find the model and call model.Delete().
func (r *SQLiteTransactionRepository) Delete(id string) error {
	// TODO: Implement the delete logic.
	return fmt.Errorf("delete not implemented")
}

// Clone is a placeholder for now.
func (r *SQLiteTransactionRepository) Clone(newTransactionName string, transaction *entity.Transaction) (*entity.Transaction, error) {
	// TODO: Implement the clone logic.
	return nil, fmt.Errorf("clone not implemented")
}
