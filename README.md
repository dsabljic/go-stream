# go-stream

Jednostavni go package za map, reduce, filter i pipe operacije.

## Requirements

- Go (1.24.0)

## Upute za korištenje

Package se može koristiti i isprobati na dva načina.

### Prvi način - `go get`

U vlastitom Golang projektu/modulu pokrenuti sljedeću komandu:

```zsh
go get github.com/dsabljic/go-stream
```

Zatim uvezite package sa sljedećim `import` statementom.

```go
import stream "github.com/dsabljic/go-stream"
```

Nakon toga, sve je spremno za upotrebu.

Primjer upotrebe

```go
a := []int{1, 2, 3, 4, 5}

doubledA := Map(a, func(el, index int, array []int) int {
    x := array[index] * 2
    return x * el
})
```

Svi ostali primjeri su dostupni u `main.go` datoteci.

### Drugi način

Klonirajte repozitorij:

```bash
git clone https://github.com/dsabljic/go-stream.git
cd go-stream
# go mod tidy # nepotrebno jer se koristi samo std lib
```

Zatim pokrenite primjere sa:

```bash
go run .
```

Nakon toga možete pokrenuti:

```bash
make prep
```

te otvoriti `coverage.html` u browseru kako bi vidjeli pokrivenost testovima, te možete i pokrenuti unit testove sa:

```bash
go test ./... # ili go test -v ./... za detaljniji pregled
```