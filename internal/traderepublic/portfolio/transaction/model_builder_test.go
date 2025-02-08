package transaction_test

import (
	"fmt"
	"io"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/reader"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/details"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/transactions"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/portfolio/transaction"
	"github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes"

	_ "github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes/transaction"
)

func TestModelBuilderBuildSupported(t *testing.T) {
	t.Parallel()

	logger := log.New()
	logger.Out = io.Discard
	controller := gomock.NewController(t)
	readerMock := reader.NewMockInterface(controller)
	detailsClient := details.NewClient(readerMock, logger)
	normalizer := details.NewTransactionResponseNormalizer(logger)
	builderFactory := transaction.ProvideModelBuilderFactory(logger)

	for testCaseName, testCase := range fakes.TransactionTestCasesSupported {
		readerMock.
			EXPECT().
			Read("timelineDetailV2", gomock.Any()).
			DoAndReturn(func(_ string, _ map[string]any) (reader.JSONResponse, error) {
				return reader.NewJSONResponse([]byte(testCase.Source.JSON)), nil
			})

		var response details.Response

		err := detailsClient.Details(testCase.Result.UUID, &response)
		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))

		normalizedResponse, err := normalizer.Normalize(response)
		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))

		builder, err := builderFactory.Create(transactions.EventType(testCase.Source.Parent.Result.EventType), normalizedResponse)
		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))

		if err != nil {
			return
		}

		actual, err := builder.Build()
		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))
		assert.Equal(t, testCase.Result, actual, fmt.Sprintf("case '%s'", testCaseName))
	}
}

func TestModelBuilderBuildUnsupported(t *testing.T) {
	t.Parallel()

	testCases := fakes.TransactionTestCasesUnsupported
	logger := log.New()
	controller := gomock.NewController(t)
	readerMock := reader.NewMockInterface(controller)
	detailsClient := details.NewClient(readerMock, logger)
	normalizer := details.NewTransactionResponseNormalizer(logger)
	builderFactory := transaction.ProvideModelBuilderFactory(logger)

	for testCaseName, testCase := range testCases {
		readerMock.
			EXPECT().
			Read("timelineDetailV2", gomock.Any()).
			DoAndReturn(func(_ string, _ map[string]any) (reader.JSONResponse, error) {
				return reader.NewJSONResponse([]byte(testCase.Source.JSON)), nil
			})

		var response details.Response

		err := detailsClient.Details(testCase.Result.UUID, &response)
		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))

		normalizedResponse, _ := normalizer.Normalize(response)

		_, err = builderFactory.Create(testCase.EventType, normalizedResponse)
		assert.Error(t, err, fmt.Sprintf("case '%s'", testCaseName))
	}
}

func TestModelBuilderBuildUnknown(t *testing.T) {
	t.Parallel()

	logger := log.New()
	controller := gomock.NewController(t)
	readerMock := reader.NewMockInterface(controller)
	detailsClient := details.NewClient(readerMock, logger)
	normalizer := details.NewTransactionResponseNormalizer(logger)
	builderFactory := transaction.ProvideModelBuilderFactory(logger)

	for testCaseName, testCase := range fakes.TransactionTestCasesUnknown {
		readerMock.
			EXPECT().
			Read("timelineDetailV2", gomock.Any()).
			DoAndReturn(func(_ string, _ map[string]any) (reader.JSONResponse, error) {
				return reader.NewJSONResponse([]byte(testCase.Source.JSON)), nil
			})

		var response details.Response

		err := detailsClient.Details(testCase.Result.UUID, &response)
		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))

		normalizedResponse, err := normalizer.Normalize(response)
		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))

		builder, err := builderFactory.Create(testCase.EventType, normalizedResponse)
		assert.NoError(t, err, fmt.Sprintf("case '%s'", testCaseName))

		if err != nil {
			return
		}

		_, err = builder.Build()
		assert.ErrorIs(t, err, transaction.ErrModelBuilderInsufficientDataResolved, fmt.Sprintf("case '%s'", testCaseName))
	}
}
