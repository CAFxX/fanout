package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 string }

func main() {
    var v1 int8 = -128
    var v2 int16 = 42
    var v3 bool = false
    v4 := make(map[int]uint32)
    var v5 uint = 559
    var v6 int = 1
    switch "" {
    case "foo bar":
        var v7 float32 = (-2.25 - (1.5 + 3.14159))
        _ = v7
    case "hello":
        fmt.Println(v5)
    }
    v6 = (^(v6))
    if v3 {
        for fi8 := 0; fi8 < 5; fi8++ {
            switch ("hello" + "hello") {
            case "":
                var v9 uint8 = uint8(v5)
                _ = v9
            case "foo bar":
                var v10 uint64 = 42
                _ = v10
            case "hello":
                v3 = v3
            }
            v11 := []int{(fi8 & 9223372036854775807)}
            _ = fi8
            _ = v11
        }
    }
    if ((v2 & 1) <= v2) {
        var v12 int16 = 0
        v13 := v4[v6]
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 211, F1: 128}
    v6 = (-(v6))
    var v15 uint64 = 709
    switch v6 {
    case -2147483648:
        fmt.Println(v3)
    case 150:
        v3 = (65535 <= v5)
    default:
        var v16 uint16 = 255
        _ = v16
    }
    fmt.Println(v15, v3)
    _ = len(v4)
    v17 := GenS1{F0: 0, F1: 255}
    v18 := GenS2{F0: 70, F1: -2147483648, F2: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
