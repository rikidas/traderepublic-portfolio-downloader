// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/dhojayev/traderepublic-portfolio-downloader/cmd/portfoliodownloader"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/api"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/api/auth"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/api/timeline/details"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/api/timeline/transactions"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/api/websocket"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/filesystem"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/portfolio"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/portfolio/transaction"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/writer"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

func CreateLocalApp(baseDir string, logger *logrus.Logger) (portfoliodownloader.App, error) {
	jsonReader := filesystem.NewJSONReader(baseDir, logger)
	client := transactions.NewClient(jsonReader)
	detailsClient := details.NewClient(jsonReader)
	typeResolver := transaction.NewTypeResolver(logger)
	detailsDeserializer := transaction.NewDetailsDeserializer(typeResolver, logger)
	csvEntryFactory := transaction.NewCSVEntryFactory(logger)
	csvReader := filesystem.NewCSVReader(logger)
	csvWriter := filesystem.NewCSVWriter(logger)
	processor := transaction.NewProcessor(detailsDeserializer, csvEntryFactory, csvReader, csvWriter, logger)
	app := portfoliodownloader.NewApp(client, detailsClient, processor, logger)
	return app, nil
}

func CreateRemoteApp(phoneNumber auth.PhoneNumber, pin auth.Pin, logger *logrus.Logger) (portfoliodownloader.App, error) {
	client := api.NewClient(logger)
	authClient, err := auth.NewClient(phoneNumber, pin, client, logger)
	if err != nil {
		return portfoliodownloader.App{}, err
	}
	jsonWriter := filesystem.NewJSONWriter(logger)
	reader, err := websocket.NewReader(authClient, jsonWriter, logger)
	if err != nil {
		return portfoliodownloader.App{}, err
	}
	transactionsClient := transactions.NewClient(reader)
	detailsClient := details.NewClient(reader)
	typeResolver := transaction.NewTypeResolver(logger)
	detailsDeserializer := transaction.NewDetailsDeserializer(typeResolver, logger)
	csvEntryFactory := transaction.NewCSVEntryFactory(logger)
	csvReader := filesystem.NewCSVReader(logger)
	csvWriter := filesystem.NewCSVWriter(logger)
	processor := transaction.NewProcessor(detailsDeserializer, csvEntryFactory, csvReader, csvWriter, logger)
	app := portfoliodownloader.NewApp(transactionsClient, detailsClient, processor, logger)
	return app, nil
}

// wire.go:

var (
	DefaultSet = wire.NewSet(portfoliodownloader.NewApp, transactions.NewClient, details.NewClient, transaction.NewTypeResolver, transaction.NewDetailsDeserializer, transaction.NewCSVEntryFactory, filesystem.NewCSVReader, filesystem.NewCSVWriter, transaction.NewProcessor, wire.Bind(new(transaction.DetailsDeserializerInterface), new(transaction.DetailsDeserializer)), wire.Bind(new(filesystem.FactoryInterface), new(transaction.CSVEntryFactory)))

	RemoteSet = wire.NewSet(
		DefaultSet, api.NewClient, auth.NewClient, filesystem.NewJSONWriter, websocket.NewReader, wire.Bind(new(auth.ClientInterface), new(*auth.Client)), wire.Bind(new(writer.Interface), new(filesystem.JSONWriter)), wire.Bind(new(portfolio.ReaderInterface), new(*websocket.Reader)),
	)

	LocalSet = wire.NewSet(
		DefaultSet, writer.NewNilWriter, filesystem.NewJSONReader, wire.Bind(new(writer.Interface), new(writer.NilWriter)), wire.Bind(new(portfolio.ReaderInterface), new(filesystem.JSONReader)),
	)
)
