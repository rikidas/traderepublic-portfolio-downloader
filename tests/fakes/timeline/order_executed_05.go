package timeline_test

import (
	"github.com/dhojayev/traderepublic-portfolio-downloader/internal/traderepublic/api/timeline/transactions"
	"github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes"
)

var OrderExecuted05 = fakes.TimelineTestCase{
	Enabled: true,
	JSON: `{
	"action": {
		"payload": "eb6ee8c7-2cff-4dcc-ab70-3ca7f31f0371",
		"type": "timelineDetail"
	},
	"amount": {
		"currency": "EUR",
		"fractionDigits": 2,
		"value": -5001.01
	},
	"badge": null,
	"eventType": "ORDER_EXECUTED",
	"icon": "logos/DE0007500001/v2",
	"id": "eb6ee8c7-2cff-4dcc-ab70-3ca7f31f0371",
	"status": "EXECUTED",
	"subAmount": null,
	"subtitle": "Kauforder",
	"timestamp": "2023-09-12T06:35:52.879+0000",
	"title": "Anleihe Feb. 2024"
	}`,
	Result: transactions.ResponseItem{
		Action: transactions.ResponseItemAction{
			Payload: "eb6ee8c7-2cff-4dcc-ab70-3ca7f31f0371",
			Type:    "timelineDetail",
		},
		Amount: transactions.ResponseItemAmount{
			Currency:       "EUR",
			FractionDigits: 2,
			Value:          -5001.01,
		},
		EventType: "ORDER_EXECUTED",
		Icon:      "logos/DE0007500001/v2",
		ID:        "eb6ee8c7-2cff-4dcc-ab70-3ca7f31f0371",
		Status:    "EXECUTED",
		Subtitle:  "Kauforder",
		Timestamp: "2023-09-12T06:35:52.879+0000",
		Title:     "Anleihe Feb. 2024",
	},
}

func init() {
	fakes.TimelineTestCasesSupported["OrderExecuted05"] = OrderExecuted05
}
