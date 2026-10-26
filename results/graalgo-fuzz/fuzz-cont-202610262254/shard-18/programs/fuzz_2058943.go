package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint8 = 42
    v2 := []string{"", ("" + "\n\t")}
    var v3 uint16 = 377
    var v4 bool = (false && (((((1e10 - ((1.5 / ((-2.25 + 1.5) / 1e10)) - 1e10)) - 1.5) + -2.25) - 3.14159) < 1e10))
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 int8 = 100
        _ = fi5
        _ = v6
    }
    v3 = 2
    switch -1 {
    case 0:
        if v4 {
            v7 := GenS1{F0: 2}
            _ = v7
        } else {
            var v8 string = "\n\t"
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        _ = len(v2)
        v1 = 42
        v2 = append(v2, ("hello" + ""))
        _ = fi9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 float64 = -2.25
        v11 -= v11
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v4, v1, v1)
        switch "foo bar" {
        case "":
            fmt.Println(v4)
        }
        switch (("hello" + "a") + "\n\t") {
        case "\n\t":
            switch (fi12 ^ (fi12 >> ((65535) & 63))) {
            case 4294967295:
                v13 := GenS2{F0: true, F1: (^(v3))}
                _ = v13
            }
        case "foo bar":
            v1 = 1
        }
        _ = fi12
    }
    fmt.Println(v4, v1, v4)
    v14 := GenS2{F0: (v1 <= (v1 + 36)), F1: 588}
    if (!v4) {
        var v15 uint16 = v3
        v16 := GenS2{F0: (!(!v4)), F1: (v15 - v3)}
        if (!(v4 && v4)) {
            fmt.Println(v3, v4)
            v15 -= 602
            fmt.Println(v1)
        }
        _ = v15
        _ = v16
    } else {
        var v17 bool = (("\n\t" == "\n\t") || false)
        v17 = ("\n\t" == "hello")
        v4 = false
        _ = v17
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
