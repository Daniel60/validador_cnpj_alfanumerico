package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	TAMANHO_CNPJ_SEM_DV = 12
	VALOR_BASE          = '0'
)

var (
	PESOS_DV                    = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	REGEX_CARACTERES_FORMATACAO = regexp.MustCompile(`[./-]`)
	REGEX_FORMACAO_BASE_CNPJ    = regexp.MustCompile(`[A-Z\d]{12}`)
	REGEX_FORMACAO_DV           = regexp.MustCompile(`[\d]{2}`)
	REGEX_VALOR_ZERADO          = regexp.MustCompile(`[0]+$`)
)

func ValidadorCNPJ(fl validator.FieldLevel) bool {
	cnpj := fl.Field().String()
	return IsValidCNPJ(cnpj)
}

func IsValidCNPJ(cnpj string) bool {
	cnpj = RemoveCaracteresFormatacao(cnpj)

	if len(cnpj) != TAMANHO_CNPJ_SEM_DV+2 {
		return false
	}

	if isCNPJFormacaoValidaComDV(cnpj) {
		dvInformado := cnpj[TAMANHO_CNPJ_SEM_DV:]
		dvCalculado := calculaDV(cnpj[:TAMANHO_CNPJ_SEM_DV])
		return dvCalculado == dvInformado
	}

	return false
}

func calculaDV(baseCNPJ string) string {
	if isCNPJFormacaoValidaSemDV(baseCNPJ) {
		dv1 := strconv.Itoa(calculaDigito(baseCNPJ))
		dv2 := strconv.Itoa(calculaDigito(baseCNPJ + dv1))
		return dv1 + dv2
	}
	fmt.Printf("Error no CalculoDV, CNPJ %s não é válido para o cálculo do DV", baseCNPJ)
	return ""
}
func calculaDigito(cnpj string) int {
	soma := 0
	for i := len(cnpj) - 1; i >= 0; i-- {
		valorCaracter := int(cnpj[i]) - VALOR_BASE
		soma += valorCaracter * PESOS_DV[len(PESOS_DV)-len(cnpj)+i]
	}
	if soma%11 < 2 {
		return 0
	}
	return 11 - (soma % 11)
}

func RemoveCaracteresFormatacao(cnpj string) string {
	return REGEX_CARACTERES_FORMATACAO.ReplaceAllString(strings.TrimSpace(cnpj), "")
}

func isCNPJFormacaoValidaSemDV(cnpj string) bool {
	return REGEX_FORMACAO_BASE_CNPJ.MatchString(cnpj) && !REGEX_VALOR_ZERADO.MatchString(cnpj)
}

func isCNPJFormacaoValidaComDV(cnpj string) bool {
	return REGEX_FORMACAO_BASE_CNPJ.MatchString(cnpj[:TAMANHO_CNPJ_SEM_DV]) &&
		REGEX_FORMACAO_DV.MatchString(cnpj[TAMANHO_CNPJ_SEM_DV:]) &&
		!REGEX_VALOR_ZERADO.MatchString(cnpj)
}

func main() {
	cnpj := "28.UV8.YY9/0001-74"
	fmt.Println("CNPJ limpo: ", RemoveCaracteresFormatacao(cnpj))
	fmt.Println(IsValidCNPJ(cnpj))
}
