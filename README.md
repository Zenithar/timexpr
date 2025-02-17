# TimExpr 

This simple library uses a custom [PEG grammar](internal/parser/grammar.peg) to parse and evaluate time expressions.
It allows you to write time expressions in a natural language and evaluate them 
to a `time.Time` object.

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/Zenithar/timexpr"
)

func main() {
    // Parse the expression
    ts, err := timexpr.Parse("tomorrow")
    if err != nil {
        panic(err)
    }

    fmt.Println(ts.Format(time.RFC3339))
}
```

## Supported expressions

The following expressions are supported:

- `now`: returns the current time
- `today`: returns the current date at midnight
- `tomorrow`: returns the current time + 1 day
- `yesterday`: returns the current time - 1 day
- `next (NUMBER)? UNIT`: returns the current time + NUMBER UNIT 
- `last (NUMBER)? UNIT`: returns the crrent time - NUMBER UNIT
- `NUMBER UNIT (ago|later|sooner)` : returns the reference time - NUMBER UNIT

Where `UNIT` can be one of the following:

- `s` / `sec` / `second` / `seconds` for seconds
- `m` / `min`/ `minute`/ `minutes` for minutes
- `h` / `hour` / `hours` for hours
- `d` / `day` / `days` for days
- `w` / `week`/ `weeks` for weeks
- `M` / `month` / `months` for months
- `y` / `year` / `years` for years

## Development

To generate the parser, you need to install the `pigeon` tool:

```bash
go install github.com/mna/pigeon@latest
```

Then you can generate the parser:

```bash
cd internal/parser
pigeon -optimize-parser -optimize-basic-latin -o grammar.go grammar.peg
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
