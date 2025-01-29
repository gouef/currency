<img align=right width="168" src="docs/gouef_logo.png">

# Currency
The `currency` package provides functionality to manage and validate different currency codes and their details. It contains a list of currency codes with information such as the name, symbol, and minor units (decimal places). You can use this package to retrieve currency information and validate whether a currency code exists. ([ISO 4217](https://en.wikipedia.org/wiki/ISO_4217)) 

[![GoDoc](https://pkg.go.dev/badge/github.com/gouef/currency.svg)](https://pkg.go.dev/github.com/gouef/currency)
[![GitHub stars](https://img.shields.io/github/stars/gouef/currency?style=social)](https://github.com/gouef/currency/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouef/currency)](https://goreportcard.com/report/github.com/gouef/currency)
[![codecov](https://codecov.io/github/gouef/currency/branch/main/graph/badge.svg?token=YUG8EMH6Q8)](https://codecov.io/github/gouef/currency)

## Versions
![Stable Version](https://img.shields.io/github/v/release/gouef/currency?label=Stable&labelColor=green)
![GitHub Release](https://img.shields.io/github/v/release/gouef/currency?label=RC&include_prereleases&filter=*rc*&logoSize=diago)
![GitHub Release](https://img.shields.io/github/v/release/gouef/currency?label=Beta&include_prereleases&filter=*beta*&logoSize=diago)

## Features

- Search for a currency by its code (e.g., USD, EUR).
- Validate if a currency code exists.
- Return detailed information such as currency name, symbol, and minor units.

## Installation

```shell
go get -u github.com/gouef/currency
```

## Usages

### Finding a currency
You can use the `FindCurrency` function to search for a currency by its code. It returns a pointer to the `Currency` struct, which contains details like the name, symbol, and minor units.

```go
package main

import (
	"fmt"
	"log"
	"github.com/gouef/currency"
)

func main() {
	code := "USD"
	curr, err := currency.FindCurrency(code)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Currency Code: %s\n", curr.Code)
	fmt.Printf("Currency Name: %s\n", curr.Name)
	fmt.Printf("Currency Symbol: %s\n", curr.Symbol)
	fmt.Printf("Minor Units: %d\n", curr.MinorUnits)
}

```

### Validating a Currency Code
To check whether a currency code is valid, use the `ValidateCurrency` function. It returns `true` if the code exists and `false` otherwise.

```go
package main

import (
	"fmt"
	"log"
	"github.com/gouef/currency"
)

func main() {
	code := "EUR"
	if currency.ValidateCurrency(code) {
		fmt.Printf("The currency code %s is valid!\n", code)
	} else {
		log.Printf("The currency code %s is not valid.\n", code)
	}
}

```

## Contributing

Read [Contributing](CONTRIBUTING.md)

## Contributors

<div>
<span>
  <a href="https://github.com/JanGalek"><img src="https://raw.githubusercontent.com/gouef/currency/refs/heads/contributors-svg/.github/contributors/JanGalek.svg" alt="JanGalek" /></a>
</span>
<span>
  <a href="https://github.com/actions-user"><img src="https://raw.githubusercontent.com/gouef/currency/refs/heads/contributors-svg/.github/contributors/actions-user.svg" alt="actions-user" /></a>
</span>
</div>

