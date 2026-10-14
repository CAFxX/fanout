package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 int8; F1 int }
func genFn0() int {
    var v6 bool = true
    var v7 uint16 = 1
    v6 = (42 > 2)
    v7 -= (^(v7))
    v7 += (v7 | v7)
    _ = v6
    _ = v7
    return -658
}

func main() {
    var v1 float32 = float32(3.14159)
    var v2 uint8 = 121
    var v3 string = "a"
    v4 := make(map[int]uint32)
    v5 := []string{""}
    v1 -= ((v1 / -2.25) + v1)
    var v8 uint16 = 0
    var v9 int = 128
    switch (v9 | 389) {
    case 2147483647:
        var v10 int32 = 42
        _ = v10
    case -2147483648:
        fmt.Println(v2, v2, v8)
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 float32 = float32(1.5)
        switch 9223372036854775807 {
        case -2147483648:
            v13 := GenS1{F0: 1, F1: false, F2: uint(v2)}
            _ = v13
        default:
            v14 := GenS1{F0: v8, F1: false, F2: 590}
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    if false {
        if (-32768 < 42) {
            switch "\n\t" {
            case "":
                var v15 uint = 1
                _ = v15
            default:
                for fi16 := 0; fi16 < 9; fi16++ {
                    _ = len(v4)
                    switch (fi16 >> ((2) & 63)) {
                    case 312:
                        v3 = v3
                    case 127:
                        v17 := GenS1{F0: v8, F1: true, F2: 0}
                        _ = v17
                    case -842:
                        var v18 int = 9223372036854775807
                        _ = v18
                    }
                    v1 = v1
                    _ = fi16
                }
            }
            fmt.Println(v9)
            var v19 float64 = -2.25
            _ = v19
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                if false {
                    var v21 int = (fi20 * (v9 - genFn0()))
                    switch v3 {
                    case "a":
                        var v22 uint32 = 1
                        _ = v22
                    }
                    v21 = (fi20 >> ((255) & 63))
                    _ = v21
                }
                if true {
                    v23 := v4[v9]
                    v24 := GenS2{F0: 100, F1: (v9 - (v9 - v9))}
                    _ = v23
                    _ = v24
                }
                _ = fi20
            }
            v25 := GenS2{F0: -128, F1: (v9 - v9)}
            _ = v25
        }
    } else {
        var v26 uint16 = v8
        for fi27 := 0; fi27 < 6; fi27++ {
            v4[127] = 2147483647
            _ = fi27
        }
        fmt.Println(v9, v1, v8)
        _ = v26
    }
    var v28 int = (v9 & -2147483648)
    fmt.Println(v3, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
