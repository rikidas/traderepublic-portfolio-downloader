// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package transaction

import (
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/api/timeline/details"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/api/timeline/transactions"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/database"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/filesystem"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/portfolio/document"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/reader"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/writer"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func ProvideHandler(responseReader reader.Interface, responseWriter writer.Interface, dbConnection *gorm.DB, logger *logrus.Logger) (Handler, error) {
	client := transactions.NewClient(responseReader, logger)
	detailsClient := details.NewClient(responseReader, logger)
	transactionResponseNormalizer := details.NewTransactionResponseNormalizer(logger)
	eventTypeResolver := transactions.NewEventTypeResolver(logger)
	typeResolver := details.NewTypeResolver(logger)
	dateResolver := document.NewDateResolver(logger)
	modelBuilder := document.NewModelBuilder(dateResolver, logger)
	modelBuilderFactory := NewModelBuilderFactory(typeResolver, modelBuilder, logger)
	repository, err := ProvideTransactionRepository(dbConnection, logger)
	if err != nil {
		return Handler{}, err
	}
	csvEntryFactory := NewCSVEntryFactory(logger)
	csvReader := filesystem.NewCSVReader(logger)
	csvWriter := filesystem.NewCSVWriter(logger)
	downloader := document.NewDownloader(logger)
	processor := NewProcessor(modelBuilderFactory, repository, csvEntryFactory, csvReader, csvWriter, downloader, logger)
	handler := NewHandler(client, detailsClient, transactionResponseNormalizer, eventTypeResolver, processor, logger)
	return handler, nil
}

// wire.go:

func ProvideTransactionRepository(db *gorm.DB, logger *logrus.Logger) (*database.Repository[*Model], error) {
	return database.NewRepository[*Model](db, logger)
}

func ProvideInstrumentRepository(db *gorm.DB, logger *logrus.Logger) (*database.Repository[*Instrument], error) {
	return database.NewRepository[*Instrument](db, logger)
}
