# Conversor Numérico

Este projeto é uma biblioteca Go para converter números em texto por extenso (em português) e vice-versa.

## Instalação

Para instalar, utilize o comando:

```bash
go get github.com/vitortenor/conversor-numerico/numbers
```

## Uso como biblioteca

Importe o pacote `numbers` no seu projeto Go:

```go
import "github.com/vitortenor/conversor-numerico/numbers"

func main() {
    num, err := numbers.TextToNumber("quinhentos e vinte e três mil e onze")
    if err != nil {
        panic(err)
    }
    fmt.Println(num)

    text, err := numbers.NumberToText(523011)
    if err != nil {
        panic(err)
    }
    fmt.Println(text)
}
```

## Estrutura do projeto

```
conversor-numerico/
├── number_to_text.go
├── text_to_number.go
├── types.go
├── utils.go
├── go.mod
└── .gitignore
```

## Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests. Este projeto é feito para ajudar!

## Licença

MIT
