package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 string }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint16 = 255
    var v2 bool = false
    var v3 float32 = (1.5 / -2.25)
    var v4 int64 = 0
    var v5 int8 = 0
    var v6 uint16 = 42
    v5 -= -128
    switch "foo bar" {
    case "":
        v6 -= (v6 - 856)
    case "hello":
        v7 := GenS1{F0: v5, F1: 4294967295, F2: "a"}
        _ = v7
    case "\n\t":
        if false {
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := []int{}
                v1 = 2
                var v10 uint32 = 0
                _ = fi8
                _ = v10
                _ = v9
            }
            fmt.Println(v3, v4, v5)
        }
    }
    switch "foo bar" {
    case "foo bar":
        v6 += v6
    default:
        fmt.Println(v1)
    }
    v11 := GenS1{F0: v5, F1: 18446744073709551615, F2: (("" + ("hello" + "")) + ("foo bar" + (("\n\t" + "foo bar") + ("foo bar" + "foo bar"))))}
    fmt.Println(v3)
    v4 -= (v4 / ((v4) | 1))
    v12 := GenS2{F0: 0, F1: -1}
    var v13 int = 176
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
