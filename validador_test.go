package validador

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type TestStructCNPJ struct {
	Document string `validate:"cnpj"`
}

func TestIsCNPJ(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("cnpj", ValidadorCNPJ)

	tests := []struct {
		name     string
		document string
		expected bool
	}{
		{"Valido CNPJ", "12345678000195", true},
		{"Valido CNPJ com mascara", "12.345.678/0001-95", true},
		{"Valido CNPJ alfanumerico", "28UV8YY9000174", true},
		{"Valido CNPJ alfanumerico com mascara", "28.UV8.YY9/0001-74", true},
		{"Invalido CNPJ", "12345678900111", false},
		{"Invalido CNPJ alfanumerico", "28.UV8.YY9/0001-70", false},
		{"Invalido CNPJ alfanumerico com mascara", "28.UV8.YY9/0001-70", false},
		{"Invalido CNPJ alfanumerico com mascara", "28.UV8.YY9/0001-70", false},
		{"Invalido todos os digitos iguais", "1111111111111", false},
		{"Invalido vazio", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testStruct := TestStructCNPJ{
				Document: tt.document,
			}
			err := validate.Struct(testStruct)
			if tt.expected {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
