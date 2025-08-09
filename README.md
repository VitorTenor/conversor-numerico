# Conversor Numérico

Este projeto é uma biblioteca Go para conversão bidirecional entre números e texto em português, com suporte a decimais e valores monetários!

Transforme palavras em números e números em palavras.

## Instalação

Para instalar a versão estável mais recente, utilize o comando:

```bash
go get github.com/vitortenor/conversor-numerico@v1.0.0
```

Ou para a versão mais recente:

```bash
go get github.com/vitortenor/conversor-numerico@latest
```

## Uso como biblioteca

Importe o pacote no seu projeto Go:

```go
import (
    "fmt"
    numbers "github.com/vitortenor/conversor-numerico"
)

func main() {
    // Conversão de inteiros
    num, err := numbers.TextToNumber("quinhentos e vinte e três mil e onze")
    if err != nil {
        panic(err)
    }
    fmt.Println(num) // Saída: 523011

    text, err := numbers.NumberToText(523011)
    if err != nil {
        panic(err)
    }
    fmt.Println(text) // Saída: quinhentos e vinte e três mil e onze

    // Conversão decimal
    decimal, err := numbers.TextToDecimal("cinquenta e dois vírgula vinte e cinco")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%.2f\n", decimal.ToFloat64()) // Saída: 52.25

    // Conversão monetária
    money, err := numbers.TextToDecimal("cinquenta reais e vinte e cinco centavos")
    if err != nil {
        panic(err)
    }
    fmt.Printf("R$ %.2f\n", money.ToFloat64()) // Saída: R$ 50.25

    // Converter decimal para texto
    decimalText, err := numbers.DecimalToText(52.25, false)
    if err != nil {
        panic(err)
    }
    fmt.Println(decimalText) // Saída: cinquenta e dois vírgula vinte e cinco

    // Converter valor monetário para texto
    moneyText, err := numbers.DecimalToText(50.25, true)
    if err != nil {
        panic(err)
    }
    fmt.Println(moneyText) // Saída: cinquenta reais e vinte e cinco centavos
}
```

## Funcionalidades

### Números Inteiros
- Converter texto para números: `"vinte e cinco"` → `25`
- Converter números para texto: `25` → `"vinte e cinco"`

### Números Decimais
- Converter texto decimal para números: `"três vírgula quatorze"` → `3.14`
- Converter números para texto decimal: `3.14` → `"três vírgula quatorze"`

### Valores Monetários
- Converter texto monetário para valores: `"cinquenta reais e vinte e cinco centavos"` → `50.25`
- Converter valores para texto monetário: `50.25` → `"cinquenta reais e vinte e cinco centavos"`
- Trata formas singulares/plurais: `"um real"`, `"dois reais"`, `"um centavo"`, `"dois centavos"`

### Formatos Suportados
- **Inteiros**: `"quarenta e dois"` → `42`
- **Decimais com vírgula**: `"cinquenta e dois vírgula vinte e cinco"` → `52.25`
- **Decimais com ponto**: `"três ponto quatorze"` → `3.14`
- **Monetário**: `"cinquenta reais e vinte e cinco centavos"` → `50.25`
- **Apenas reais**: `"cem reais"` → `100.00`
- **Apenas centavos**: `"cinquenta centavos"` → `0.50`

## Estrutura do projeto

```
conversor-numerico/
├── number_to_text.go
├── text_to_number.go
├── types.go
├── utils.go
└── go.mod
```

## Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests. Este projeto é feito para ajudar!

## Licença

MIT
