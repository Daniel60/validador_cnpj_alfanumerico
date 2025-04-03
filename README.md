## Validador de CNPJ Alfanumérico

Este projeto é uma implementação em Go para validação de CNPJs Alfanuméricos (Cadastro Nacional da Pessoa Jurídica) no Brasil. Ele utiliza a biblioteca go-playground/validator para validações customizadas e inclui funções para limpar, validar e calcular os dígitos verificadores de um CNPJ.

### Funcionalidades

- Remoção de formatação: Remove caracteres como pontos, barras e traços de um CNPJ.
- Validação de CNPJ: Verifica se um CNPJ é válido, considerando sua formação e dígitos verificadores.
- Cálculo de dígitos verificadores (DV): Calcula os dois dígitos verificadores de um CNPJ com base nos pesos definidos.

### Estrutura do Projeto

- **main.go:** Contém a lógica principal para validação de CNPJs.
- **main_test.go**: Inclui testes unitários para validar o comportamento das funções.

### Pré-requisitos

- Go 1.18 ou superior.
- Biblioteca go-playground/validator para validações customizadas.
- Biblioteca testify para testes unitários.

### Instalação

1. Clone o repositório:

```shell
git clone https://github.com/seu-usuario/validador-cnpj.git
cd validador-cnpj
```

2. Instale as dependências:

```shell
go mod tidy
```

### Uso

#### Executar o programa principal

1. Edite o arquivo main.go para incluir o CNPJ que deseja validar.
2. Execute o programa:

```shell
go run main.go
```

Exemplo de saída:

```shell
CNPJ limpo:  28UV8YY9000174
true
```

### Testes

Para rodar os testes unitários, execute:

```shell
go test ./...
```

Exemplo de saída:

```shell
ok  	_/path/to/project	0.123s
```

### Exemplo de Código

#### Validação de CNPJ

```go
cnpj := "28.UV8.YY9/0001-74"
fmt.Println("CNPJ limpo: ", RemoveCaracteresFormatacao(cnpj))
fmt.Println(IsValidCNPJ(cnpj))
```

#### Teste Unitário

```go
func TestIsCNPJ(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("cnpj", ValidadorCNPJ)

	tests := []struct {
		name     string
		document string
		expected bool
	}{
		{"Valido CNPJ", "12345678000195", true},
		{"Invalido CNPJ", "12345678900111", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testStruct := TestStructCNPJ{Document: tt.document}
			err := validate.Struct(testStruct)
			if tt.expected {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
```

### Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

### Licença

Este projeto está licenciado sob a [MIT License](https://opensource.org/license/mit).
