package util

import (
	"strconv"
)

const (
	EN = iota
	CN = iota
)

var defaultErrorId = -1000

var getErrorMessageId = LookupFuncWithDefaultKey(M_ERROR_CODES_CN, strconv.Itoa(defaultErrorId))

var getErrorMessageEn = LookupFuncWithDefaultKey(M_ERROR_CODES_EN, strconv.Itoa(defaultErrorId))

func GetErrorString(language int, id int) string {

	if language == CN {
		return getErrorMessageId(strconv.Itoa(id))
	}

	if language == EN {
		return getErrorMessageEn(strconv.Itoa(id))
	}

	return getErrorMessageEn(strconv.Itoa(id))
}
