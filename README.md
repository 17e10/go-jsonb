# go-jsonb

[![GoDev][godev-image]][godev-url]

go-jsonb は JSON ファイルの読み書きを提供します.

## Usage

```go
import "github.com/17e10/go-jsonb"

type person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
john := person{"john", 20}
jsonb.Save(filename, &john)

var readed person
jsonb.Load(filename, &readed)
readed == john
```

## License

This software is released under the MIT License, see LICENSE.

## Author

17e10

[godev-image]: https://pkg.go.dev/badge/github.com/17e10/go-jsonb
[godev-url]: https://pkg.go.dev/github.com/17e10/go-jsonb
