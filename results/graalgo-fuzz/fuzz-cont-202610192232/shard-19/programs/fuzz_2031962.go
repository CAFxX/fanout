package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 bool; F1 int32 }

func main() {
    var v1 int16 = 0
    var v2 uint8 = 250
    v3 := []string{"", "a", ("foo bar" + "\n\t")}
    v1 += 1
    v4 := GenS1{F0: "\n\t", F1: 2, F2: 9223372036854775807}
    var v5 uint = 42
    fmt.Println(v2)
    var v6 uint = (v5 % (((v5 >> ((18446744073709551615) & 63))) | 1))
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v6)
    case "hello":
        v7 := GenS2{F0: (18446744073709551615 != 1), F1: -523}
        _ = v7
    default:
        fmt.Println(v1, v1)
    }
    var v8 uint64 = 0
    v8 = v8
    v2 = (v2 & 255)
    v8 -= 430
    if true {
        v6 += 1
        switch 4294967295 {
        case 4294967295:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 32767:
            v5 -= (v6 / ((v5) | 1))
        case 1:
            var v9 int16 = (v1 - (v1 % ((1000) | 1)))
            _ = v9
        default:
            fmt.Println(v5, v6, v6)
        }
        var v10 bool = (!false)
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
