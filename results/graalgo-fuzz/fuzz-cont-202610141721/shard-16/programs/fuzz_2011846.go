package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
func genFn0(p0 int32) uint32 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 9; fi4++ {
        p0 = (^(p0))
        _ = fi4
    }
    if ((!false) || (-666 <= 9223372036854775807)) {
        p0 = (p0 << ((62) & 63))
        if false {
            var v5 int8 = 1
            v6 := GenS1{F0: 1, F1: (v5 >> ((467) & 63))}
            _ = v5
            _ = v6
        } else {
            p0 = p0
        }
        v7 := []int{-701, 32767}
        _ = v7
    } else {
        p0 += p0
        var v8 int8 = 1
        p0 = p0
        _ = v8
    }
    _ = p0
    return 65535
}

func main() {
    var v1 int16 = 42
    var v2 int = 683
    v3 := make(map[string]bool)
    fmt.Println(v1)
    var v9 float64 = float64((0.1 * 1.5))
    var v10 float64 = v9
    var v11 int16 = (v1 * (v1 - 1))
    v12 := []int{-9223372036854775808, v2}
    var v13 uint32 = 445
    v14 := GenS2{F0: (^(v13)), F1: (^(v1)), F2: 1}
    switch "foo bar" {
    case "foo bar":
        var v15 int32 = -2147483648
        _ = v15
    case "\n\t":
        switch (v2 / (((-(v2))) | 1)) {
        case 0:
            var v16 uint16 = 32767
            _ = v16
        default:
            v17 := v3[("" + "a")]
            _ = v17
        }
    }
    v13 -= v13
    var v18 float32 = 1.5
    var v19 uint32 = v13
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 uint8 = 255
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 bool = true
    var v24 bool = (-523 > int32(v2))
    v25 := GenS2{F0: (^(v13)), F1: (^(v11)), F2: v13}
    var v26 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v9
}
