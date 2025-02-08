package timeline_test

import (
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/transactions"
	"github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes"
)

var InterestPayout02 = fakes.TimelineTestCase{
	Enabled: true,
	JSON:    `{
			"action": {
			"payload": "05d83d0f-0810-4f1e-a2c5-5833644904d5",
			"type": "timelineDetail"
			},
			"amount": {
			"currency": "EUR",
			"fractionDigits": 2,
			"value": 1
			},
			"badge": null,
			"cashAccountNumber": null,
			"deleted": false,
			"eventType": "INTEREST_PAYOUT",
			"hidden": false,
			"icon": "logos/timeline_interest_new/v2",
			"id": "05d83d0f-0810-4f1e-a2c5-5833644904d5",
			"status": "EXECUTED",
			"subAmount": null,
			"subtitle": "3.5 % p.a.",
			"timestamp": "2024-10-01T03:23:47.140+0000",
			"title": "Zinsen"
		}`,
	Result:  transactions.ResponseItem{
		Action: transactions.ResponseItemAction{
			Payload: "05d83d0f-0810-4f1e-a2c5-5833644904d5",
			Type:    "timelineDetail",
		},
		Amount: transactions.ResponseItemAmount{
			Currency:       "EUR",
			FractionDigits: 2,
			Value:          1,
		},
		EventType: "INTEREST_PAYOUT",
		Icon:      "logos/timeline_interest_new/v2",
		ID:        "05d83d0f-0810-4f1e-a2c5-5833644904d5",
		Status:    "EXECUTED",
		Subtitle:  "3.5 % p.a.",
		Timestamp: "2024-10-01T03:23:47.140+0000",
		Title:     "Zinsen",
	},
}

func init() {
	fakes.TimelineTestCasesSupported["InterestPayout02"] = InterestPayout02
}