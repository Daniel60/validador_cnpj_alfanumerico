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

Importe o package e utilize:
1. A função  `ValidadorCNPJ(cnpj string)` para validar com retorno boleano

2. A função `RemoveCaracteresFormatacao(cnpj string)` para tirar a mascara do cnpj

### Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

### Licença

Este projeto está licenciado sob a [MIT License](https://opensource.org/license/mit).
