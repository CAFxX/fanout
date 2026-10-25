package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 int8 }

func main() {
    var v1 bool = true
    var v2 uint16 = 2
    var v3 uint8 = 128
    v4 := []int{0, 882, 32768, 9223372036854775807}
    v1 = false
    var v5 float64 = 1.5
    _ = len(v4)
    v2 += (v2 / (((v2 % ((928) | 1))) | 1))
    fmt.Println(v1, v1)
    switch "foo bar" {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = 814
        }
    case "\n\t":
        v5 = v5
    }
    v6 := GenS1{F0: 1, F1: 802, F2: -522}
    fmt.Println(v5)
    var v7 float32 = 0.1
    var v8 int32 = 789
    var v9 int8 = 42
    _ = len(v4)
    var v10 int64 = -1
    if v1 {
        if (1 > 42) {
            fmt.Println(v9, v5)
            v11 := GenS1{F0: (v10 / (((v10 - 831)) | 1)), F1: v8, F2: v8}
            _ = v11
        }
        v12 := GenS2{F0: uint16(v3), F1: 1, F2: (^(v9))}
        _ = v12
    }
    var v13 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
