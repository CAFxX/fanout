package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 uint32, p1 uint, p2 int64) int64 {
    p0 += (^(p0))
    v4 := GenS1{F0: int16(p1)}
    var v5 uint = (p1 % (((p1 << ((uint(p0)) & 63))) | 1))
    var v6 uint32 = p0
    v7 := GenS1{F0: -887}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]int)
    var v3 uint32 = 2147483647
    var v8 uint = 322
    v3 -= (v3 - 986)
    var v9 bool = ((^(v1)) == v1)
    v1 = (-(v1))
    var v10 int8 = 1
    if v9 {
        var v11 int = -2147483648
        delete(v2, (v3 % (((v3 << ((v8) & 63))) | 1)))
        v9 = (1 <= uint8(v8))
        _ = v11
    } else {
        v12 := GenS1{F0: 0}
        switch "a" {
        case "a":
            if true {
                if true {
                    if false {
                        fmt.Println(v1, v8, v3)
                        fmt.Println(v3)
                        var v13 float64 = (-2.25 / (0.1 - 1e10))
                        _ = v13
                    }
                } else {
                    v9 = ((!v9) || (((-2.25 - ((3.14159 - (0.1 + 0.0)) + -2.25)) + 1.5) == 0.0))
                }
            }
        case "hello":
            if v9 {
                v14 := GenS2{F0: v1, F1: v3, F2: (v9 && false)}
                fmt.Println(v8)
                _ = v14
            }
        default:
            v3 -= v3
        }
        delete(v2, v3)
        _ = v12
    }
    switch -823 {
    case 4294967295:
        v15 := []int{127}
        _ = v15
    case 2147483647:
        for fi16 := 0; fi16 < 18; fi16++ {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                v17 := GenS2{F0: (v1 ^ (v1 ^ v1)), F1: 42, F2: (!false)}
                _ = v17
            }
            v18 := GenS2{F0: v1, F1: 396, F2: ((v3 * v3) < (v3 * 0))}
            _ = fi16
            _ = v18
        }
    case -2147483648:
        v8 += 2
    }
    v19 := GenS3{F0: 42}
    v20 := GenS1{F0: -32768}
    v9 = v9
    var v21 int64 = -1
    var v22 int16 = -1
    var v23 float64 = float64(3.14159)
    v10 = 12
    var v24 int = -2147483648
    var v25 int32 = (v1 * v1)
    if (2 <= 526) {
        if ((v10 == v10) || v9) {
            for fi26 := 0; fi26 < 7; fi26++ {
                v8 += (v8 ^ v8)
                var v27 int8 = v10
                _ = fi26
                _ = v27
            }
            var v28 int32 = (v1 & (-(v1)))
            _ = v28
        }
    }
    v29 := GenS4{F0: 1, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v8
    _ = v9
}
