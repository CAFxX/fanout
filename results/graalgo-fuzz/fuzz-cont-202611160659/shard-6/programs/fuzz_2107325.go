package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 bool }
func genFn0(p0 uint16, p1 uint64) bool {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := []int{-2147483648, (^(fi4)), -500, fi4}
        var v6 uint16 = (p0 * p0)
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint32 = 2147483648
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 uint = 42
    var v2 float32 = 1.5
    var v3 string = "foo bar"
    if true {
        var v9 uint16 = 0
        v10 := GenS1{F0: v3, F1: 75, F2: (false || ((0.1 > -2.25) && (false && (344 == 2147483648))))}
        _ = v10
        _ = v9
    }
    var v11 uint8 = 95
    v3 = (("" + v3) + v3)
    switch -9223372036854775808 {
    case 32768:
        v12 := []int{2147483647}
        _ = v12
    case -681:
        v2 = (1.5 - v2)
    case 1:
        v13 := []int{601, -2147483648, -2147483648}
        _ = v13
    default:
        var v14 int32 = -2147483648
        _ = v14
    }
    v1 = v1
    v3 = ((("a" + v3) + "\n\t") + v3)
    var v15 uint16 = 672
    var v16 int = 2147483647
    v2 += -2.25
    v17 := []int{(v16 + v16), v16}
    if (-1 != 9223372036854775807) {
        _ = len(v17)
    } else {
        fmt.Println(v1, v3)
        v18 := GenS1{F0: ((v3 + "foo bar") + "\n\t"), F1: 0, F2: true}
        v19 := GenS1{F0: v3, F1: -128, F2: (0 < 4294967295)}
        _ = v18
        _ = v19
    }
    v3 = "foo bar"
    fmt.Println(v1, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
