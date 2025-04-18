## Validador de CNPJ Alfanumérico

Este projeto é uma implementação em Go para validação de CNPJs Alfanuméricos (Cadastro Nacional da Pessoa Jurídica) no Brasil. Ele utiliza a biblioteca go-playground/validator para validações customizadas e inclui funções para limpar, validar e calcular os dígitos verificadores de um CNPJ.

### Funcionalidades

- **Remoção de formatação**: remove caracteres como pontos, barras e traços de um CNPJ.
- **Validação de CNPJ**: verifica se um CNPJ é válido, considerando sua formação e dígitos verificadores.
- **Cálculo de dígitos verificadores (DV)**: calcula os dois dígitos verificadores de um CNPJ com base nos pesos definidos.

### Estrutura do Projeto

- **validador.go:** Contém a lógica principal para validação de CNPJs.
- **validador_test.go**: Inclui testes unitários para validar o comportamento das funções.

```bash
validador_cnpj_alfanumerico/
├── go.mod
├── go.sum
├── .gitignore
├── README.md
├── LICENSE
└── validador/
    ├── validador.go
    └── validador_test.go
```

### Pré-requisitos

- Go 1.20 ou superior.
- Biblioteca go-playground/validator para validações customizadas.
- Biblioteca testify para testes unitários.

### Instalação

1. Instale o package:

```shell
go get github.com/Daniel60/validador_cnpj_alfanumerico
```

2. Usar as dependências:

```go
import (
	"fmt"

	val "github.com/Daniel60/validador_cnpj_alfanumerico/validadorCNPJ"
	"github.com/go-playground/validator/v10"
)

type Empresa struct {
	CNPJ string `validate:"cnpj"`
}

func main() {
	// como usar o validador com tags nos structs
	v10 := validator.New(validator.WithRequiredStructEnabled())
	v10.RegisterValidation("cnpj", val.ValidadorCNPJField)

	// struct
	empresa := Empresa{CNPJ: "12.345.678/0001-95"}
	err := v10.Struct(empresa)

	// como usar as funções do package
	a := val.IsValidCNPJ("122334")
	b := val.RemoveCaracteresFormatacao("asas_22")

	fmt.Println(a, b, empresa, err == nil)
}
```

### Uso

Importe o package e utilize:

1. A função  `ValidadorCNPJ(fl validator.FieldLevel) bool` para validar com retorno boleano e é usado com a biblioteca de validaçao v10 ou o binding do gin gonic

2. A função `RemoveCaracteresFormatacao(cnpj string) string` para tirar a mascara do cnpj e devolve string

3. A função `IsValidCNPJ(cnpj string) bool` valida a string e devolver o boleano

### Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

### Licença

Este projeto está licenciado sob a [MIT License](https://opensource.org/license/mit).
