// Package math is a package with the validador
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	// TAMANHO_CNPJ_SEM_DV defines the length of the CNPJ without the verification digits (DV).
	TAMANHO_CNPJ_SEM_DV = 12
	// VALOR_BASE is the base value used for character-to-integer conversion.
	VALOR_BASE = '0'
)

var (
	// PESOS_DV defines the weights used for calculating the verification digits (DV) of the CNPJ.
	PESOS_DV = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	// REGEX_CARACTERES_FORMATACAO is a regular expression to match formatting characters in a CNPJ (e.g., ".", "/", "-").
	REGEX_CARACTERES_FORMATACAO = regexp.MustCompile(`[./-]`)
	// REGEX_FORMACAO_BASE_CNPJ is a regular expression to validate the base formation of a CNPJ (12 alphanumeric characters).
	REGEX_FORMACAO_BASE_CNPJ = regexp.MustCompile(`[A-Z\d]{12}`)
	// REGEX_FORMACAO_DV is a regular expression to validate the formation of the verification digits (2 numeric characters).
	REGEX_FORMACAO_DV = regexp.MustCompile(`[\d]{2}`)
	// REGEX_VALOR_ZERADO is a regular expression to check if a CNPJ consists only of zeros.
	REGEX_VALOR_ZERADO = regexp.MustCompile(`[0]+$`)
)

// ValidadorCNPJ is a custom validator function for validating CNPJ fields using the validator package.
func ValidadorCNPJ(fl validator.FieldLevel) bool {
	cnpj := fl.Field().String()
	return IsValidCNPJ(cnpj)
}

// IsValidCNPJ validates a given CNPJ string by checking its length, format, and verification digits.
func IsValidCNPJ(cnpj string) bool {
	// Remove formatting characters from the CNPJ.
	cnpj = RemoveCaracteresFormatacao(cnpj)

	// Check if the CNPJ has the correct length (base + DV).
	if len(cnpj) != TAMANHO_CNPJ_SEM_DV+2 {
		return false
	}

	// Validate the CNPJ format and calculate the verification digits.
	if isCNPJFormacaoValidaComDV(cnpj) {
		dvInformado := cnpj[TAMANHO_CNPJ_SEM_DV:]
		dvCalculado := calculaDV(cnpj[:TAMANHO_CNPJ_SEM_DV])
		return dvCalculado == dvInformado
	}

	return false
}

// calculaDV calculates the verification digits (DV) for a given CNPJ base.
func calculaDV(baseCNPJ string) string {
	// Check if the base CNPJ has a valid format.
	if isCNPJFormacaoValidaSemDV(baseCNPJ) {
		// Calculate the first and second verification digits.
		dv1 := strconv.Itoa(calculaDigito(baseCNPJ))
		dv2 := strconv.Itoa(calculaDigito(baseCNPJ + dv1))
		return dv1 + dv2
	}
	// Log an error if the base CNPJ is invalid for DV calculation.
	fmt.Printf("Error no CalculoDV, CNPJ %s não é válido para o cálculo do DV", baseCNPJ)
	return ""
}

// calculaDigito calculates a single verification digit for a given CNPJ string.
func calculaDigito(cnpj string) int {
	soma := 0
	// Iterate over the CNPJ characters and calculate the weighted sum.
	for i := len(cnpj) - 1; i >= 0; i-- {
		valorCaracter := int(cnpj[i]) - VALOR_BASE
		soma += valorCaracter * PESOS_DV[len(PESOS_DV)-len(cnpj)+i]
	}
	// Calculate the verification digit based on the weighted sum.
	if soma%11 < 2 {
		return 0
	}
	return 11 - (soma % 11)
}

// RemoveCaracteresFormatacao removes formatting characters (e.g., ".", "/", "-") from a CNPJ string.
func RemoveCaracteresFormatacao(cnpj string) string {
	return REGEX_CARACTERES_FORMATACAO.ReplaceAllString(strings.TrimSpace(cnpj), "")
}

// isCNPJFormacaoValidaSemDV checks if the base CNPJ (without DV) has a valid format.
func isCNPJFormacaoValidaSemDV(cnpj string) bool {
	return REGEX_FORMACAO_BASE_CNPJ.MatchString(cnpj) && !REGEX_VALOR_ZERADO.MatchString(cnpj)
}

// isCNPJFormacaoValidaComDV checks if the full CNPJ (base + DV) has a valid format.
func isCNPJFormacaoValidaComDV(cnpj string) bool {
	return REGEX_FORMACAO_BASE_CNPJ.MatchString(cnpj[:TAMANHO_CNPJ_SEM_DV]) &&
		REGEX_FORMACAO_DV.MatchString(cnpj[TAMANHO_CNPJ_SEM_DV:]) &&
		!REGEX_VALOR_ZERADO.MatchString(cnpj)
}
