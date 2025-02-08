package transaction_test

import (
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/transactions"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/portfolio/document"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/portfolio/instrument"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/portfolio/transaction"
	"github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes"
	details_test "github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes/details"
)

var BenefitsSavebackExecution01 = fakes.TransactionTestCase{
	Enabled: true,
	Source:  details_test.BenefitsSavebackExecution01,
	EventType: transactions.EventTypeBenefitsSavebackExecution,
	Result: transaction.Model{
		UUID: "73fc417a-62ef-4179-a85e-9f3b29224567",
		Instrument: instrument.Model{
			ISIN: "XF000DOT0011",
			Name: "Polkadot",
			Icon: "logos/XF000DOT0011/v2",
			Type: instrument.TypeCryptocurrency,
		},
		Type:   transaction.TypeSaveback,
		Status: "executed",
		Shares: 2.270212,
		Rate:   6.61,
		Total:  15,
		Documents: []document.Model{
			{
				TransactionUUID: "73fc417a-62ef-4179-a85e-9f3b29224567",
				ID:              "3a54ce6c-7bf7-4db5-a79e-5c24dbc71594",
				URL:             "https://traderepublic-data-production.s3.eu-central-1.amazonaws.com/timeline/postbox/",
				Detail:          "22.03.2024",
				Title:           "Abrechnung Ausführung",
				Filepath:        "2024-03/73fc417a-62ef-4179-a85e-9f3b29224567/Abrechnung Ausführung.pdf",
			},
			{
				TransactionUUID: "73fc417a-62ef-4179-a85e-9f3b29224567",
				ID:              "70776ac6-b87e-4c73-a8cb-558466234f0d",
				URL:             "https://traderepublic-data-production.s3.eu-central-1.amazonaws.com/timeline/postbox/",
				Detail:          "21.03.2024",
				Title:           "Kosteninformation",
				Filepath:        "2024-03/73fc417a-62ef-4179-a85e-9f3b29224567/Kosteninformation.pdf",
			},
		},
	},
}

func init() {
	fakes.TransactionTestCasesSupported["BenefitsSavebackExecution01"] = BenefitsSavebackExecution01
}
