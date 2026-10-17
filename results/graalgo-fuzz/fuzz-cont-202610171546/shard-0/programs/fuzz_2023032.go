package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint }
func genFn0() int32 {
    if (0 >= 100) {
        var v7 int32 = -1
        fmt.Println(v7)
        v8 := GenS1{F0: 106, F1: -128}
        _ = v7
        _ = v8
    }
    var v9 uint16 = 2
    fmt.Println(v9)
    _ = v9
    return -609
}
func genFn1() int8 {
    var v10 float64 = 1.5
    v10 = (v10 - (((1e10 + v10) / v10) / v10))
    v10 = float64(1e10)
    _ = v10
    return 127
}

func main() {
    var v1 int = -1
    var v2 uint = 643
    var v3 uint16 = 199
    var v4 int8 = -128
    var v5 string = "foo bar"
    var v6 uint16 = (v3 & v3)
    v3 = 172
    if true {
        var v11 uint = (v2 >> ((58) & 63))
        v12 := []int{v1}
        _ = v11
        _ = v12
    }
    v4 -= (v4 | v4)
    v13 := GenS1{F0: (v1 | (v1 - -105)), F1: v4}
    fmt.Println(v5, v1)
    if true {
        v14 := GenS2{F0: 2, F1: 32767, F2: 4294967295}
        fmt.Println(v2)
        fmt.Println(v4, v4)
        _ = v14
    }
    switch v5 {
    case "\n\t":
        switch ("foo bar" + v5) {
        case "a":
            fmt.Println(v4, v1, v6)
        default:
            var v15 int32 = 1
            _ = v15
        }
    default:
        v16 := GenS2{F0: 42, F1: int16(v3), F2: v2}
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int16 = 919
        switch (v1 * fi17) {
        case -986:
            v19 := []int{}
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
