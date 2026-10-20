package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0() uint32 {
    if true {
        var v6 int16 = 1000
        if (!(!false)) {
            v6 -= (v6 - v6)
            v7 := GenS1{F0: 0, F1: 360}
            switch 9223372036854775807 {
            case -2147483648:
                var v8 int = -2
                _ = v8
            case 2147483647:
                if false {
                    fmt.Println(v6)
                    var v9 int = 440
                    _ = v9
                }
            }
            _ = v7
        } else {
            v6 = v6
        }
        v10 := []int{}
        _ = v10
        _ = v6
    }
    if true {
        var v11 int16 = 42
        _ = v11
    } else {
        var v12 uint32 = 0
        _ = v12
    }
    v13 := []int{}
    v14 := GenS1{F0: 42, F1: 1}
    _ = v13
    _ = v14
    return 2
}

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = (v1 >> ((286) & 63))
    v3 := make(map[uint32]bool)
    var v4 float32 = (0.1 + (0.0 + 1e10))
    var v5 int32 = 592
    v15 := []int{1}
    var v16 int32 = v5
    v17 := GenS2{F0: v2}
    fmt.Println(v16, v2)
    v4 -= (((-2.25 - 0.1) / v4) - (((((1.5 / (v4 + (v4 / v4))) - v4) / 3.14159) + (float32(float64(1.5)) * 1.5)) + 0.0))
    v18 := GenS1{F0: -1, F1: uint64(v2)}
    for fi19 := 0; fi19 < 1; fi19++ {
        v5 = (^(v16))
        switch (^(fi19)) {
        case 4294967295:
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 float32 = ((1.5 * v4) - 3.14159)
                if (-479 == 1) {
                    v16 = (v16 >> ((601) & 63))
                }
                _ = fi20
                _ = v21
            }
        default:
            if false {
                v1 = v2
                var v22 uint16 = (v1 | 0)
                switch fi19 {
                case -489:
                    for fi23 := 0; fi23 < 19; fi23++ {
                        v22 = (v2 | (v22 % ((v1) | 1)))
                        _ = fi23
                    }
                }
                _ = v22
            }
        }
        fmt.Println(v5, v1)
        _ = fi19
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        v1 -= (v1 / ((v1) | 1))
        _ = fi24
    }
    var v25 int16 = -32768
    for fi26 := 0; fi26 < 8; fi26++ {
        fmt.Println(v5)
        _ = fi26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 uint16 = 346
        _ = fi27
        _ = v28
    }
    var v29 uint16 = (v2 ^ v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
