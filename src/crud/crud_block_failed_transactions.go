package crud

import (
	"reflect"
	"sync"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/geometry-labs/icon-blocks/models"
)

// BlockFailedTransactionModel - type for block table model
type BlockFailedTransactionModel struct {
	db            *gorm.DB
	model         *models.BlockFailedTransaction
	modelORM      *models.BlockFailedTransactionORM
	LoaderChannel chan *models.BlockFailedTransaction
}

var blockFailedTransactionModel *BlockFailedTransactionModel
var blockFailedTransactionModelOnce sync.Once

// GetBlockModel - create and/or return the blocks table model
func GetBlockFailedTransactionModel() *BlockFailedTransactionModel {
	blockFailedTransactionModelOnce.Do(func() {
		dbConn := getPostgresConn()
		if dbConn == nil {
			zap.S().Fatal("Cannot connect to postgres database")
		}

		blockFailedTransactionModel = &BlockFailedTransactionModel{
			db:            dbConn,
			model:         &models.BlockFailedTransaction{},
			LoaderChannel: make(chan *models.BlockFailedTransaction, 1),
		}

		err := blockFailedTransactionModel.Migrate()
		if err != nil {
			zap.S().Fatal("BlockFailedTransactionModel: Unable migrate postgres table: ", err.Error())
		}

		StartBlockFailedTransactionLoader()
	})

	return blockFailedTransactionModel
}

// Migrate - migrate blockFailedTransactions table
func (m *BlockFailedTransactionModel) Migrate() error {
	// Only using BlockFailedTransactionRawORM (ORM version of the proto generated struct) to create the TABLE
	err := m.db.AutoMigrate(m.modelORM) // Migration and Index creation
	return err
}

// Insert - Insert blockFailedTransaction into table
func (m *BlockFailedTransactionModel) Insert(blockFailedTransaction *models.BlockFailedTransaction) error {
	db := m.db

	// Set table
	db = db.Model(&models.BlockFailedTransaction{})

	db = db.Create(blockFailedTransaction)

	return db.Error
}

// SelectOne - select from blockFailedTransactions table
func (m *BlockFailedTransactionModel) SelectOne(transactionHash string) (*models.BlockFailedTransaction, error) {
	db := m.db

	// Set table
	db = db.Model(&models.BlockFailedTransaction{})

	// Transaction hash
	db = db.Where("transaction_hash = ?", transactionHash)

	blockFailedTransaction := &models.BlockFailedTransaction{}
	db = db.First(blockFailedTransaction)

	return blockFailedTransaction, db.Error
}

// SelectMany - select many from blockFailedTransactions table by block number
func (m *BlockFailedTransactionModel) SelectMany(number uint32) (*[]models.BlockFailedTransaction, error) {
	db := m.db

	// Set table
	db = db.Model(&models.BlockFailedTransaction{})

	// Number
	db = db.Where("number = ?", number)

	blockFailedTransactions := &[]models.BlockFailedTransaction{}
	db = db.Find(blockFailedTransactions)

	return blockFailedTransactions, db.Error
}

// UpdateOne - update in blockFailedTransactions table
func (m *BlockFailedTransactionModel) UpdateOne(blockFailedTransaction *models.BlockFailedTransaction) error {
	db := m.db

	// Set table
	db = db.Model(&models.BlockFailedTransaction{})

	// Transaction hash
	db = db.Where("transaction_hash = ?", blockFailedTransaction.TransactionHash)

	db = db.First(blockFailedTransaction)

	return db.Error
}

func (m *BlockFailedTransactionModel) UpsertOne(
	blockFailedTransaction *models.BlockFailedTransaction,
) error {
	db := m.db

	// map[string]interface{}
	updateOnConflictValues := extractFilledFieldsFromModel(
		reflect.ValueOf(*blockFailedTransaction),
		reflect.TypeOf(*blockFailedTransaction),
	)

	// Upsert
	db = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "transaction_hash"}}, // NOTE set to primary keys for table
		DoUpdates: clause.Assignments(updateOnConflictValues),
	}).Create(blockFailedTransaction)

	return db.Error
}

// StartBlockFailedTransactionLoader starts loader
func StartBlockFailedTransactionLoader() {
	go func() {

		for {
			// Read newBlockFailedTransaction
			newBlockFailedTransaction := <-GetBlockFailedTransactionModel().LoaderChannel

			//////////////////////
			// Load to postgres //
			//////////////////////
			err := GetBlockFailedTransactionModel().UpsertOne(newBlockFailedTransaction)
			zap.S().Debug("Loader=BlockFailedTransaction, TransactionHash=", newBlockFailedTransaction.TransactionHash, " - Upserted")
			if err != nil {
				// Error
				zap.S().Fatal("Loader=BlockFailedTransaction, TransactionHash=", newBlockFailedTransaction.TransactionHash, " - Error: ", err.Error())
			}

			///////////////////////
			// Force enrichments //
			///////////////////////
			err = reloadBlock(newBlockFailedTransaction.Number)
			if err != nil {
				zap.S().Fatal(err.Error())
			}
		}
	}()
}
