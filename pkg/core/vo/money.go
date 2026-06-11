package vo


import(
	"github.com/shopspring/decimal"
	
	"github.com/mustang/pkg/core/types"
)

type Money struct {
	Amount   decimal.Decimal
	Currency types.CurrencyType
}