package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/henryeffiong/gobank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if currency is supported
		return util.IsSupportedCurrency(currency)
	}

	return false
}
