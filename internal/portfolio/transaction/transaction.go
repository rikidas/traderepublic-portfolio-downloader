package transaction

import (
	"strings"
	"time"
)

const (
	TypePurchase       = "Purchase"
	TypeSale           = "Sale"
	TypeDividendPayout = "Dividends"
	TypeRoundUp        = "Round up"
	TypeSaveback       = "Saveback"

	AssetTypeStocks         = "Stocks"
	AssetTypeETF            = "ETF"
	AssetTypeCryptocurrency = "Cryptocurrency"
	AssetTypeLending        = "Lending"
	AssetTypeOther          = "Other"

	instrumentPrefixLending = "XS"
	instrumentPrefixCrypto  = "XF000"
	instrumentSuffixDist    = "(Dist)"
	instrumentSuffixAcc     = "(Acc)"
)

type Purchase struct {
	BaseTransaction BaseTransaction
	Asset           Asset
	MonetaryValues  MonetaryValues
	Documents       []Document
}

func NewPurchase(
	transaction BaseTransaction,
	asset Asset,
	monetaryValues MonetaryValues,
	documents []Document,
) Purchase {
	transaction.Type = TypePurchase

	return Purchase{
		BaseTransaction: transaction,
		Asset:           asset,
		MonetaryValues:  monetaryValues,
		Documents:       documents,
	}
}

type Sale struct {
	Purchase
	Yield  float64
	Profit float64
}

func NewSale(
	yield, profit float64,
	purchase Purchase,
) Sale {
	purchase.BaseTransaction.Type = TypeSale

	return Sale{
		Purchase: purchase,
		Yield:    yield,
		Profit:   profit,
	}
}

type Benefit struct {
	Purchase
}

func (b Benefit) IsTypeRoundUp() bool {
	return b.BaseTransaction.Type == TypeRoundUp
}

func NewBenefit(benefitType string, purchase Purchase) Benefit {
	purchase.BaseTransaction.Type = benefitType

	return Benefit{
		Purchase: purchase,
	}
}

type DividendPayout struct {
	Sale
}

func NewDividendPayout(sale Sale) DividendPayout {
	sale.BaseTransaction.Type = TypeDividendPayout

	return DividendPayout{
		sale,
	}
}

type BaseTransaction struct {
	ID        string
	Type      string
	Timestamp time.Time
	Status    string
}

func NewBaseTransaction(id, status string, timestamp time.Time) BaseTransaction {
	return BaseTransaction{
		ID:        id,
		Timestamp: timestamp,
		Status:    status,
	}
}

type Asset struct {
	Type       string
	Instrument string
	Name       string
	Shares     float64
}

func NewAsset(instrument, name string, shares float64) Asset {
	assetType := AssetTypeOther

	switch {
	case strings.HasSuffix(name, instrumentSuffixDist), strings.HasSuffix(name, instrumentSuffixAcc):
		assetType = AssetTypeETF
	case strings.HasPrefix(instrument, instrumentPrefixCrypto):
		assetType = AssetTypeCryptocurrency
	case strings.HasPrefix(instrument, instrumentPrefixLending):
		assetType = AssetTypeLending
	}

	return Asset{
		Type:       assetType,
		Instrument: instrument,
		Name:       name,
		Shares:     shares,
	}
}

type MonetaryValues struct {
	Rate       float64
	Commission float64
	Total      float64
}

func NewMonetaryValues(rate, commission, total float64) MonetaryValues {
	return MonetaryValues{
		Rate:       rate,
		Commission: commission,
		Total:      total,
	}
}
