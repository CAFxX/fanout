package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 int16 = 340
    var v2 int = 899
    var v3 int8 = 0
    v4 := []uint32{}
    v5 := make(map[string]string)
    v1 = 1
    v3 += (v3 % ((v3) | 1))
    var v6 uint = 1
    fmt.Println(v6, v2)
    if ((3.14159 + (1.5 / ((1e10 * 1.5) + -2.25))) == (-2.25 - 1.5)) {
        fmt.Println(v1)
        switch "foo bar" {
        case "a":
            switch ("foo bar" + "") {
            case "a":
                v6 = 0
            case "foo bar":
                fmt.Println(v3, v3, v2)
            }
        case "":
            v7 := v5[""]
            _ = v7
        }
    } else {
        fmt.Println(v3, v3)
        fmt.Println(v1, v2, v3)
        var v8 int8 = (^(v3))
        _ = v8
    }
    var v9 float64 = 0.1
    v10 := GenS1{F0: 136, F1: 0}
    var v11 int64 = -9223372036854775808
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
