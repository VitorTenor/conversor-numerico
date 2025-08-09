# Conversor Numérico

Este projeto é uma biblioteca Go para converter números em texto por extenso (em português) e vice-versa.

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
    num, err := numbers.TextToNumber("quinhentos e vinte e três mil e onze")
    if err != nil {
        panic(err)
    }
    fmt.Println(num) // Output: 523011

    text, err := numbers.NumberToText(523011)
    if err != nil {
        panic(err)
    }
    fmt.Println(text) // Output: quinhentos e vinte e três mil e onze
}
```

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
