package fakes

import (
	"time"

	"github.com/dhojayev/traderepublic-portfolio-downloader/internal"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/filesystem"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/details"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/transactions"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/portfolio/document"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/portfolio/instrument"
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/portfolio/transaction"
)

var InterestPayout02 = LegacyTransactionTestCase{
	TimelineDetailsData: TimelineDetailsTestData{Raw: []byte(`{
    "id": "05d83d0f-0810-4f1e-a2c5-5833644904d5",
    "sections": [
        {
            "data": {
                "icon": "logos/timeline_interest_new/v2",
                "status": "executed",
                "timestamp": "2024-10-01T05:30:16.685746+02:00"
            },
            "title": "Du hast €1 erhalten",
            "type": "header"
        },
        {
            "data": [
                {
                    "detail": {
                        "functionalStyle": "EXECUTED",
                        "text": "Ausgeführt",
                        "type": "status"
                    },
                    "style": "plain",
                    "title": "Status"
                },
                {
                    "detail": {
                        "text": "€1.00",
                        "type": "text"
                    },
                    "style": "plain",
                    "title": "Durchschnittssaldo"
                },
                {
                    "detail": {
                        "text": "3.50 %",
                        "type": "text"
                    },
                    "style": "plain",
                    "title": "Jährliche Rate"
                },
                {
                    "detail": {
                        "text": "Cash",
                        "type": "text"
                    },
                    "style": "plain",
                    "title": "Asset"
                }
            ],
            "title": "Übersicht",
            "type": "table"
        },
        {
            "steps": [
                {
                    "content": {
                        "cta": null,
                        "subtitle": null,
                        "timestamp": "2024-10-01T05:30:16.685746+02:00",
                        "title": "Zinsberechnung"
                    },
                    "leading": {
                        "avatar": {
                            "status": "completed",
                            "type": "bullet"
                        },
                        "connection": {
                            "order": "first"
                        }
                    }
                },
                {
                    "content": {
                        "cta": null,
                        "subtitle": null,
                        "timestamp": "2024-10-01T05:30:16.685746+02:00",
                        "title": "Zinszahlung"
                    },
                    "leading": {
                        "avatar": {
                            "status": "completed",
                            "type": "bullet"
                        },
                        "connection": {
                            "order": "last"
                        }
                    }
                }
            ],
            "title": "Status",
            "type": "steps"
        },
        {
            "data": [
                {
                    "detail": {
                        "text": "€1.00",
                        "type": "text"
                    },
                    "style": "plain",
                    "title": "Angesammelt"
                },
                {
                    "detail": {
                        "text": "€0.00",
                        "type": "text"
                    },
                    "style": "plain",
                    "title": "Steuern"
                },
                {
                    "detail": {
                        "text": "€1.00",
                        "type": "text"
                    },
                    "style": "plain",
                    "title": "Gesamt"
                }
            ],
            "title": "Transaktion",
            "type": "table"
        },
        {
            "data": [
                {
                    "action": {
                        "payload": "https://traderepublic-postbox-platform-production.s3.eu-central-1.amazonaws.com/timeline/postbox/2024/10/1/dbçlablçablabsddoaju2057820937502375024",
                        "type": "browserModal"
                    },
                    "id": "ecc7bf31-e271-4e92-9e14-3c4cee9443db",
                    "postboxType": "INTEREST_PAYOUT_INVOICE",
                    "title": "Abrechnung"
                }
            ],
            "title": "Dokument",
            "type": "documents"
        },
        {
            "data": [
                {
                    "detail": {
                        "action": {
                            "payload": {
                                "contextCategory": "NHC",
                                "contextParams": {
                                    "chat_flow_key": "NHC_0020_interest_past_interest_payout",
                                    "timelineEventId": "ecc7bf31-e271-4e92-9e14-3c4cee9443db"
                                }
                            },
                            "type": "customerSupportChat"
                        },
                        "icon": "",
                        "style": "highlighted",
                        "type": "listItemAvatarDefault"
                    },
                    "style": "plain",
                    "title": ""
                }
            ],
            "title": "",
            "type": "table"
        }
    ]
}`),
		Normalized: details.NormalizedResponse{
			ID:          "",
			Header:      details.NormalizedResponseHeaderSection{},
			Overview:    details.NormalizedResponseOverviewSection{},
			Performance: details.NormalizedResponsePerformanceSection{},
			Transaction: details.NormalizedResponseTransactionSection{},
			Documents:   details.NormalizedResponseDocumentsSection{},
		},
	},
	EventType: transactions.EventTypeSavingsPlanInvoiceCreated,
	Transaction: transaction.Model{
		UUID: "05d83d0f-0810-4f1e-a2c5-5833644904d5",
		Instrument: instrument.Model{
			Name: "Cash",
			Icon: "logos/timeline_interest_new/v2",
			Type: instrument.TypeCash,
		},
		Documents: []document.Model{
			{
				TransactionUUID: "05d83d0f-0810-4f1e-a2c5-5833644904d5",
				ID:              "ecc7bf31-e271-4e92-9e14-3c4cee9443db",
				URL:             "https://traderepublic-postbox-platform-production.s3.eu-central-1.amazonaws.com/timeline/postbox/2024/10/1/dbçlablçablabsddoaju2057820937502375024",
				Title:           "Abrechnung",
				Filepath:        "/05d83d0f-0810-4f1e-a2c5-5833644904d5/Abrechnung.pdf",
			},
		},
		Type:   transaction.TypeInterestPayout,
		Status: "executed",
		Total:  1,
	},
	CSVEntry: filesystem.CSVEntry{
		ID:        "05d83d0f-0810-4f1e-a2c5-5833644904d5",
		Status:    "executed",
		Type:      transaction.TypeInterestPayout,
		AssetType: string(instrument.TypeCash),
		Credit:    1,
		Documents: []string{
			"/05d83d0f-0810-4f1e-a2c5-5833644904d5/Abrechnung.pdf",
		},
	},
}

func init() {
	InterestPayout02.Transaction.Timestamp, _ = time.Parse(details.ResponseTimeFormat, "2024-10-01T05:30:16.685746+02:00")
	InterestPayout02.CSVEntry.Timestamp = internal.DateTime{Time: InterestPayout02.Transaction.Timestamp}

	RegisterSupported("InterestPayout02", InterestPayout02)
}
