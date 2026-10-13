package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS3 struct { F0 string; F1 uint }

func main() {
    var v1 int8 = -128
    v2 := make(map[string]bool)
    var v3 float64 = (-2.25 + -2.25)
    v3 = ((((v3 - v3) / -2.25) - (0.1 / v3)) - v3)
    v3 -= -2.25
    var v4 int32 = 1
    for fi5 := 0; fi5 < 18; fi5++ {
        v1 += (v1 + (v1 << ((4294967295) & 63)))
        v3 = (v3 / 0.1)
        v6 := GenS1{F0: (18446744073709551615 < 5), F1: 1}
        _ = fi5
        _ = v6
    }
    var v7 int8 = -45
    v1 -= (v1 | -1)
    var v8 bool = false
    v9 := []int{int(v7), 128, -538}
    var v10 uint = 0
    var v11 int = 42
    var v12 float64 = v3
    var v13 uint8 = 251
    v3 -= float64(((1e10 + 1e10) / -2.25))
    v14 := GenS2{F0: v13, F1: v10, F2: 1}
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint8 = (v13 | (v13 & v13))
        _ = fi15
        _ = v16
    }
    if ((v13 * v13) < (v13 & (v13 / ((v13) | 1)))) {
        fmt.Println(v10)
        v12 -= v3
    } else {
        if ((((1.5 - v12) + 1.5) / v3) <= (1e10 - v3)) {
            switch v11 {
            case -395:
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 int16 = -571
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 float64 = float64(1e10)
                _ = v19
            }
            var v20 int64 = -1
            v21 := GenS3{F0: "a", F1: v10}
            _ = v20
            _ = v21
        } else {
            fmt.Println(v3)
            v13 = (v13 * (^(v13)))
        }
        v22 := v2[""]
        _ = v22
    }
    var v23 uint32 = 581
    fmt.Println(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
