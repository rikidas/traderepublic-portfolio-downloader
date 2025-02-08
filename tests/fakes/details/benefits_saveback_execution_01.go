package details_test

import "github.com/dhojayev/traderepublic-portfolio-downloader/tests/fakes"

var BenefitsSavebackExecution01 = fakes.DetailsTestCase{
	Enabed: false,
	JSON: `{
		"id": "73fc417a-62ef-4179-a85e-9f3b29224567",
		"sections": [
		  {
			"action": null,
			"data": {
			  "icon": "logos/XF000DOT0011/v2",
			  "status": "executed",
			  "subtitleText": null,
			  "timestamp": "2024-03-22T18:15:06.448+0000"
			},
			"title": "Dein Bonus von 15,00 € wurde investiert",
			"type": "header"
		  },
		  {
			"action": null,
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
				  "action": null,
				  "text": "Saveback",
				  "trend": null,
				  "type": "text"
				},
				"style": "plain",
				"title": "Auftragsart"
			  },
			  {
				"detail": {
				  "action": null,
				  "text": "Polkadot",
				  "trend": null,
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
			"action": null,
			"data": [
			  {
				"detail": {
				  "action": null,
				  "text": "2.270212",
				  "trend": null,
				  "type": "text"
				},
				"style": "plain",
				"title": "Aktien"
			  },
			  {
				"detail": {
				  "action": null,
				  "text": "6,61 €",
				  "trend": null,
				  "type": "text"
				},
				"style": "plain",
				"title": "Aktienkurs"
			  },
			  {
				"detail": {
				  "action": null,
				  "text": "Kostenlos",
				  "trend": null,
				  "type": "text"
				},
				"style": "plain",
				"title": "Gebühr"
			  },
			  {
				"detail": {
				  "action": null,
				  "text": "15,00 €",
				  "trend": null,
				  "type": "text"
				},
				"style": "highlighted",
				"title": "Gesamt"
			  }
			],
			"title": "Transaktion",
			"type": "table"
		  },
		  {
			"action": null,
			"data": [
			  {
				"action": {
				  "payload": "https://traderepublic-data-production.s3.eu-central-1.amazonaws.com/timeline/postbox/",
				  "type": "browserModal"
				},
				"detail": "22.03.2024",
				"id": "3a54ce6c-7bf7-4db5-a79e-5c24dbc71594",
				"postboxType": "SAVINGS_PLAN_EXECUTED_V2",
				"title": "Abrechnung Ausführung"
			  },
			  {
				"action": {
				  "payload": "https://traderepublic-data-production.s3.eu-central-1.amazonaws.com/timeline/postbox/",
				  "type": "browserModal"
				},
				"detail": "21.03.2024",
				"id": "70776ac6-b87e-4c73-a8cb-558466234f0d",
				"postboxType": "COSTS_INFO_SAVINGS_PLAN_V2",
				"title": "Kosteninformation"
			  }
			],
			"title": "Dokumente",
			"type": "documents"
		  }
		]
	  }`,
}

func init() {
	fakes.DetailsTestCasesSupported["BenefitsSavebackExecution01"] = BenefitsSavebackExecution01
}
