package fakes

import (
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/filesystem"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/activitylog"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/details"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/transactions"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/portfolio/transaction"
)

var (
	ActivityLogTestCasesSupported = make(map[string]ActivityLogTestCase)

	TimelineTestCasesSupported      = make(map[string]TimelineTestCase)
	TimelineTestCasesUnsupported    = make(map[string]TimelineTestCase)
	DetailsTestCasesSupported       = make(map[string]DetailsTestCase)
	DetailsTestCasesUnsupported     = make(map[string]DetailsTestCase)
	TransactionTestCasesSupported   = make(map[string]TransactionTestCase)
	TransactionTestCasesUnsupported = make(map[string]TransactionTestCase)
	TransactionTestCasesUnknown     = make(map[string]TransactionTestCase)
	CSVTestCasesSupported           = make(map[string]CSVTestCase)
)

type TimelineTestCase struct {
	Enabled bool
	JSON    string
	Result  transactions.ResponseItem
}

type DetailsTestCase struct {
	Enabed bool
	Parent TimelineTestCase
	JSON   string
	Result details.NormalizedResponse
}

type TransactionTestCase struct {
	Enabled   bool
	Source    DetailsTestCase
	EventType transactions.EventType
	Result    transaction.Model
}

type CSVTestCase struct {
	Enabled     bool
	Transaction transaction.Model
	Result      filesystem.CSVEntry
}

type LegacyTransactionTestCase struct {
	TimelineTransactionsData TimelineTransactionsTestData
	TimelineDetailsData      TimelineDetailsTestData
	EventType                transactions.EventType
	Transaction              transaction.Model
	CSVEntry                 filesystem.CSVEntry
}

type ActivityLogTestCase struct {
	ActivityLogData     ActivityLogTestData
	TimelineDetailsData TimelineDetailsTestData
}

type TimelineTransactionsTestData struct {
	Raw          []byte
	Unmarshalled transactions.ResponseItem
}

type ActivityLogTestData struct {
	Raw          []byte
	Unmarshalled activitylog.ResponseItem
}

type TimelineDetailsTestData struct {
	Raw        []byte
	Normalized details.NormalizedResponse
}

func RegisterSupported(name string, testCase LegacyTransactionTestCase) {
	// TransactionTestCasesSupported[name] = testCase
}

func RegisterUnsupported(name string, testCase LegacyTransactionTestCase) {
	// TransactionTestCasesUnsupported[name] = testCase
}

func RegisterUnknown(name string, testCase LegacyTransactionTestCase) {
	// TransactionTestCasesUnknown[name] = testCase
}

func RegisterActivityLogSupported(name string, testCase ActivityLogTestCase) {
	ActivityLogTestCasesSupported[name] = testCase
}
