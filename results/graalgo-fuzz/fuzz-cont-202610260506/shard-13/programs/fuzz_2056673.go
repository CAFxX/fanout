package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint }
func genFn0() int16 {
    if false {
        var v6 int = 2147483647
        switch -1 {
        case 32767:
            v7 := GenS1{F0: (^(v6)), F1: -883}
            _ = v7
        default:
            v8 := GenS1{F0: v6, F1: -999}
            _ = v8
        }
        v9 := []int{}
        _ = v6
        _ = v9
    }
    v10 := GenS1{F0: -961, F1: 1}
    var v11 uint16 = 255
    v11 -= (v11 >> ((4294967295) & 63))
    fmt.Println(v11)
    _ = v10
    _ = v11
    return -933
}

func main() {
    var v1 int8 = 127
    var v2 uint32 = 4294967295
    var v3 float32 = (1.5 * ((1.5 * 3.14159) + 3.14159))
    var v4 uint32 = v2
    var v5 int = 1
    for fi12 := 0; fi12 < 9; fi12++ {
        v1 = (v1 << ((958) & 63))
        _ = fi12
    }
    v13 := []int{(v5 - (^(v5))), -706, (v5 / ((v5) | 1)), 1}
    _ = len(v13)
    fmt.Println(v3)
    v14 := GenS2{F0: v1, F1: true, F2: 4294967295}
    var v15 uint16 = 822
    fmt.Println(v5, v3)
    var v16 uint = 42
    switch "" {
    case "\n\t":
        fmt.Println(v4)
    case "a":
        v4 = (v2 / (((v2 % ((2147483648) | 1))) | 1))
    }
    v15 += v15
    if len(v13) > 0 {
        v13[0] = int(v16)
    }
    v17 := GenS1{F0: -2147483648, F1: 129}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
