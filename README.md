# Clean interface to calculate percentages in Go

[![Go version](https://img.shields.io/github/go-mod/go-version/rorecek/percentage?style=flat-square)](https://github.com/rorecek/percentage/blob/master/go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/rorecek/percentage?style=flat-square)](https://goreportcard.com/report/github.com/rorecek/percentage)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/rorecek/percentage)
[![Releases](https://img.shields.io/github/v/release/rorecek/percentage?sort=semver&style=flat-square)](https://github.com/rorecek/percentage/releases)
[![LICENSE](https://img.shields.io/github/license/rorecek/percentage?style=flat-square)](https://github.com/rorecek/percentage/blob/master/LICENSE)

A clean interface to calculate percentages in Go: changes between values, percentage increases and partial values.
Think of it as a very lightweight math library for the basic stuff. Aimed to make your code more readable and easier to understand.
Inspired by [mattiasgeniar/php-percentages](https://github.com/mattiasgeniar/php-percentages)

## Installation

You can install this module with ```go get``` command:

``` bash
go get -u github.com/rorecek/percentage
```

## Usage

``` go
import "github.com/rorecek/percentage"

// What's the percentage increase from 100 to 120?
percentage.DifferenceBetween(100, 120);         // 20

// What's the absolute percentage change from 100 to 80?
percentage.AbsoluteDifferenceBetween(100, 80);  // 20, not -20

// How much is 120 compared to 100?
percentage.Calculate(120, 100);                 // 120%

// How much is 50 compared to 100?
percentage.Calculate(50, 100);                  // 50%

// What is 20% of 200?
percentage.Of(20, 200);                         // 40
```

## Testing

``` bash
go test
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](CONTRIBUTING.md) for details.

## Security

If you discover any security related issues, please email pavel@rorecek.com instead of using the issue tracker.

## Credits

- [Pavel Rorecek](https://github.com/rorecek)
- [Mattias Geniar](https://github.com/mattiasgeniar/php-percentages)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
