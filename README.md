# go-stream

Jednostavni go package za map, reduce, filter i pipe operacije.

## Requirements

- Go (1.24.0)

## Upute za korištenje

Package se može koristiti/isprobati na dva načina.

### Prvi način - `go get`

U vlastitom Golang projektu/modulu pokrenuti sljedeću komandu:

```zsh
go get -u github.com/dsabljic/go-stream
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

Svi ostali primjeri su dostupni u [`main.go`](https://gist.github.com/dsabljic/388dfcb92b7dae0e66f1a9c0f21bbdc0) gist-u.

### Drugi način

Klonirajte repozitorij:

```bash
git clone https://github.com/dsabljic/go-stream.git
cd go-stream
# go mod tidy # nepotrebno jer se koristi samo std lib
```

Nakon toga možete pokrenuti:

```bash
make cover
```

te otvoriti `coverage.html` u browseru kako bi vidjeli pokrivenost testovima, te možete i pokrenuti unit testove sa:

```bash
go test ./... # ili go test -v ./... za detaljniji pregled
```

Ukoliko iz nekog razloga želite izbaciti report za određene datoteke, trebate dodati zapis u `exclude.txt` datoteku u sljedećem obliku:

```
github.com/dsabljic/go-stream/[filename].[extension]
```
