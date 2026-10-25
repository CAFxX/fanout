package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint8 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }

func main() {
    var v1 int8 = -99
    v2 := make(map[uint32]int)
    var v3 int8 = 119
    v4 := make(map[string]int)
    var v5 string = ""
    var v6 float64 = 3.14159
    switch 0 {
    case -9223372036854775808:
        v7 := []int{-9223372036854775808, -9223372036854775808, -2147483648, 127}
        _ = v7
    case 4294967295:
        v5 = v5
    default:
        v1 += v1
    }
    _ = len(v2)
    if (((v6 != (1.5 + 1.5)) || true) && (!(-9223372036854775808 <= 9223372036854775807))) {
        fmt.Println(v3)
        if false {
            fmt.Println(v6)
        } else {
            v8 := GenS1{F0: 122, F1: true, F2: 2}
            var v9 bool = false
            v3 -= -1
            _ = v8
            _ = v9
        }
        var v10 int16 = 32767
        _ = v10
    }
    v11 := []int{42, 32768, -2147483648, 4294967295}
    fmt.Println(v3, v3, v3)
    v3 += v1
    v12 := GenS2{F0: 1, F1: -1, F2: 255}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    switch ("\n\t" + v5) {
    case "a":
        var v13 uint64 = 42
        _ = v13
    }
    switch 32768 {
    case -618:
        if false {
            v14 := GenS1{F0: (^(v3)), F1: false, F2: 42}
            _ = v14
        }
    case 128:
        if len(v11) > 0 {
            v11[0] = -381
        }
    case 32767:
        _ = len(v2)
    }
    var v15 uint8 = 0
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
