package validador

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	tamanhoCNPJSemDV = 12
	valorBase        = '0'
)

var (
	pesosDV                   = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	regexCaracteresFormatacao = regexp.MustCompile(`[./-]`)
	regexFormacaoBaseCNPJ     = regexp.MustCompile(`[A-Z\d]{12}`)
	regexFormacaoDV           = regexp.MustCompile(`[\d]{2}`)
	regexValorZerado          = regexp.MustCompile(`[0]+$`)
)

// ValidadorCNPJ is a custom validator function for validating CNPJ fields using the validator package.
func ValidadorCNPJ(fl validator.FieldLevel) bool {
	cnpj := fl.Field().String()
	return IsValidCNPJ(cnpj)
}

// IsValidCNPJ validates a given CNPJ string by checking its length, format, and verification digits.
func IsValidCNPJ(cnpj string) bool {
	cnpj = RemoveCaracteresFormatacao(cnpj)

	if len(cnpj) != tamanhoCNPJSemDV+2 {
		return false
	}

	if isCNPJFormacaoValidaComDV(cnpj) {
		dvInformado := cnpj[tamanhoCNPJSemDV:]
		dvCalculado := calculaDV(cnpj[:tamanhoCNPJSemDV])
		return dvCalculado == dvInformado
	}

	return false
}

// RemoveCaracteresFormatacao removes formatting characters (e.g., ".", "/", "-") from a CNPJ string.
func RemoveCaracteresFormatacao(cnpj string) string {
	return regexCaracteresFormatacao.ReplaceAllString(strings.TrimSpace(cnpj), "")
}

// calculaDV calculates the verification digits (DV) for a given CNPJ base.
func calculaDV(baseCNPJ string) string {
	if isCNPJFormacaoValidaSemDV(baseCNPJ) {
		dv1 := strconv.Itoa(calculaDigito(baseCNPJ))
		dv2 := strconv.Itoa(calculaDigito(baseCNPJ + dv1))
		return dv1 + dv2
	}
	fmt.Printf("Error no CalculoDV, CNPJ %s não é válido para o cálculo do DV", baseCNPJ)
	return ""
}

// calculaDigito calculates a single verification digit for a given CNPJ string.
func calculaDigito(cnpj string) int {
	soma := 0
	for i := len(cnpj) - 1; i >= 0; i-- {
		valorCaracter := int(cnpj[i]) - valorBase
		soma += valorCaracter * pesosDV[len(pesosDV)-len(cnpj)+i]
	}
	if soma%11 < 2 {
		return 0
	}
	return 11 - (soma % 11)
}

// isCNPJFormacaoValidaSemDV checks if the base CNPJ (without DV) has a valid format.
func isCNPJFormacaoValidaSemDV(cnpj string) bool {
	return regexFormacaoBaseCNPJ.MatchString(cnpj) && !regexValorZerado.MatchString(cnpj)
}

// isCNPJFormacaoValidaComDV checks if the full CNPJ (base + DV) has a valid format.
func isCNPJFormacaoValidaComDV(cnpj string) bool {
	return regexFormacaoBaseCNPJ.MatchString(cnpj[:tamanhoCNPJSemDV]) &&
		regexFormacaoDV.MatchString(cnpj[tamanhoCNPJSemDV:]) &&
		!regexValorZerado.MatchString(cnpj)
}
