package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0() int64 {
    var v6 uint32 = 4294967295
    v7 := []int{-1, 1, 433}
    for fi8 := 0; fi8 < 5; fi8++ {
        v6 = v6
        _ = fi8
    }
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 int32, p2 uint8) uint16 {
    var v9 uint8 = (p2 >> ((383) & 63))
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(p1)
        var v11 uint = 1
        p0 = (!((!p0) || p0))
        _ = fi10
        _ = v11
    }
    switch ("hello" + "a") {
    case "\n\t":
        fmt.Println(p1, p0, p1)
    default:
        v12 := []int{-1, 488}
        _ = v12
    }
    var v13 uint8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v9
    return 444
}

func main() {
    v1 := []int{-1, -845, 42, 194}
    var v2 uint8 = 127
    var v3 float32 = (0.0 * 1e10)
    var v4 bool = false
    var v5 uint32 = 982
    v2 = 253
    for fi14 := 0; fi14 < 4; fi14++ {
        switch (fi14 << ((991) & 63)) {
        case -678:
            v4 = (0.0 != -2.25)
        default:
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := GenS1{F0: -1}
                if false {
                    var v17 int16 = 42
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        }
        _ = fi14
    }
    var v18 int64 = -840
    if false {
        v19 := GenS2{F0: 53, F1: (931 <= 1)}
        if true {
            for fi20 := 0; fi20 < 9; fi20++ {
                v3 = ((((v3 - ((float32((0.0 + (((1.5 - 3.14159) - 0.0) / -2.25))) - v3) / (v3 - 3.14159))) - float32(-2.25)) + float32((1e10 + (0.1 * float64((v3 * v3)))))) + v3)
                _ = fi20
            }
            var v21 int64 = v18
            _ = v21
        }
        var v22 int32 = 438
        _ = v19
        _ = v22
    }
    var v23 int8 = 60
    for fi24 := 0; fi24 < 2; fi24++ {
        v5 -= (v5 * 985)
        _ = fi24
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    v23 -= 127
    fmt.Println(v5, v3)
    switch (((("a" + "hello") + ("" + "a")) + "\n\t") + "foo bar") {
    case "\n\t":
        var v25 int = 2147483647
        _ = v25
    case "foo bar":
        var v26 uint64 = 387
        _ = v26
    case "hello":
        var v27 float32 = (0.0 + v3)
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 int = int(v2)
        v3 -= (-2.25 * v3)
        _ = fi28
        _ = v29
    }
    var v30 float64 = float64(((float32(0.1) - v3) - 3.14159))
    for fi31 := 0; fi31 < 3; fi31++ {
        v1 = append(v1, (fi31 | fi31))
        v32 := GenS2{F0: (v2 % (((v2 & (v2 & (v2 - v2)))) | 1)), F1: false}
        _ = fi31
        _ = v32
    }
    v2 = (v2 & 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
