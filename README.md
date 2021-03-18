![Build Status](https://github.com/ppreeper/pad/actions/workflows/go.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/ppreeper/pad/badge.svg?branch=master)](https://coveralls.io/github/ppreeper/pad?branch=master)

# pad
Pad Library for padding string with various patterns. This has been tuned for high speed operation.

## Quick Ref

This library provides several common functions:

Common Padding

* Left
* LeftLen
* Right
* RightLen

Zero Fill Pattern

* ZFill
* ZFillLen

Justified text

* LJust
* LJustLen
* RJust
* RJustLen.

The functions Left, Right, LJust, RJust, ZFill add a pattern to the string and return the result, which is of size input + pattern * count.

The functions LeftLen, RightLen, LJustLen, RJustLen, ZFillLen add a pattern to the string and return a string of the specified size.

```go
Left(s string, p string, c int)
Right(s string, p string, c int)
ZFill(s string, c int)
LJust(s string, c int)
RJust(s string, c int)
```

```go
LeftLen(s string, p string, l int)
RightLen(s string, p string, l int)
ZFillLen(s string, l int)
LJustLen(s string, l int)
RJustLen(s string, l int)
```

## Example Usage

Please look at the example code.

### main.go
```go
package main

import (
	"fmt"

	"github.com/ppreeper/pad"
)

func main() {
	fmt.Println(pad.Right("Hello", "!", 20))
	fmt.Println(pad.Left("Exit now", "→", 20))
	fmt.Println("-----")
	fmt.Println(pad.Left("D", "a", 4))
	fmt.Println(pad.Right("D", "a", 4))
	fmt.Println(pad.ZFill("D", 4))
	fmt.Println(pad.LJust("D", 4))
	fmt.Println(pad.RJust("D", 4))
	fmt.Println("-----")
	fmt.Println(pad.Left("ABCD", "a", 5))
	fmt.Println(pad.Right("ABCD", "a", 5))
	fmt.Println(pad.ZFill("ABCD", 5))
	fmt.Println(pad.LJust("ABCD", 5))
	fmt.Println(pad.RJust("ABCD", 5))
	fmt.Println("-----")
	fmt.Println(pad.LeftLen("D", "a", 4))
	fmt.Println(pad.RightLen("D", "a", 4))
	fmt.Println(pad.ZFillLen("D", 4))
	fmt.Println(pad.LJustLen("D", 4))
	fmt.Println(pad.RJustLen("D", 4))
	fmt.Println("-----")
	fmt.Println(pad.LeftLen("ABCD", "a", 5))
	fmt.Println(pad.RightLen("ABCD", "a", 5))
	fmt.Println(pad.ZFillLen("ABCD", 5))
	fmt.Println(pad.LJustLen("ABCD", 5))
	fmt.Println(pad.RJustLen("ABCD", 5))
}
```

### output
```
Hello!!!!!!!!!!!!!!!!!!!!
→→→→→→→→→→→→→→→→→→→→Exit now
-----
aaaaD
Daaaa
0000D
D    
    D
-----
aaaaaABCD
ABCDaaaaa
00000ABCD
ABCD     
     ABCD
-----
aaaD
Daaa
000D
D   
   D
-----
aABCD
ABCDa
0ABCD
ABCD 
 ABCD
```


## Benchmarks

Current benchmarks 

```
goos: linux
goarch: amd64
pkg: github.com/ppreeper/pad
BenchmarkLeft_D_utf8_4-12               10987552               109 ns/op
BenchmarkLeft_D_a_4-12                  18355047                71.9 ns/op
BenchmarkLeft_D_a_200-12                 3388698               329 ns/op
BenchmarkLeftLen_D_utf8_4-12            12383883                92.7 ns/op
BenchmarkLeftLen_D_a_4-12               14235428                80.8 ns/op
BenchmarkLeftLen_D_a_200-12              3804021               308 ns/op
BenchmarkLJust_utf8_4-12                14946373                77.5 ns/op
BenchmarkLJust_D_4-12                   15604454                69.7 ns/op
BenchmarkLJust_D_a_200-12                4065147               319 ns/op
BenchmarkLJustLen_utf8_4-12             14027293                80.5 ns/op
BenchmarkLJustLen_D_4-12                14452222                79.4 ns/op
BenchmarkLJustLen_D_a_200-12             3666148               318 ns/op
BenchmarkRight_D_utf8_4-12              11142861               107 ns/op
BenchmarkRight_D_a_4-12                 17843385                71.3 ns/op
BenchmarkRight_D_a_200-12                3416127               341 ns/op
BenchmarkRightLen_D_utf8_4-12           12343040                89.9 ns/op
BenchmarkRightLen_D_a_4-12              14721405                80.6 ns/op
BenchmarkRightLen_D_a_200-12             3458692               354 ns/op
BenchmarkRJust_utf8_4-12                13767662                82.1 ns/op
BenchmarkRJust_D_4-12                   17698326                70.0 ns/op
BenchmarkRJust_D_a_200-12                3509358               342 ns/op
BenchmarkRJustLen_utf8_4-12             15495074                81.3 ns/op
BenchmarkRJustLen_D_4-12                13987210                83.6 ns/op
BenchmarkRJustLen_D_a_200-12             3430438               350 ns/op
BenchmarkZFill_D_4-12                   15297421                69.6 ns/op
BenchmarkZFill_DDDD_4-12                13453654                80.0 ns/op
BenchmarkZFill_D_a_200-12                3537516               352 ns/op
BenchmarkZFillLen_D_4-12                14793592                79.0 ns/op
BenchmarkZFillLen_DDDD_4-12             17648059                81.6 ns/op
BenchmarkZFillLen_D_a_200-12             3490700               349 ns/op
PASS
ok      github.com/ppreeper/pad 40.849s
```

