package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint8 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []string{"\n\t"}
    v2 := []uint32{65535, 1, 427}
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    var v5 int64 = -203
    var v6 int8 = 0
    v4[128] = (false && false)
    v5 -= 1
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int8 = 22
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int8 = v6
            var v11 uint = 65535
            _ = fi9
            _ = v10
            _ = v11
        }
        switch (fi7 / (((fi7 + int(v6))) | 1)) {
        case 127:
            for fi12 := 0; fi12 < 16; fi12++ {
                if ((42 > 1) || (v5 > (v5 << ((65535) & 63)))) {
                    _ = len(v2)
                    v13 := GenS1{F0: 18446744073709551615, F1: fi7, F2: 193}
                    _ = v13
                } else {
                    switch -851 {
                    case 4294967295:
                        v5 += -811
                    case 128:
                        switch "" {
                        case "\n\t":
                            v2 = append(v2, 595)
                        case "a":
                            var v14 float64 = (1e10 - (0.1 * float64((0.1 + v3))))
                            _ = v14
                        case "foo bar":
                            v4[fi12] = (32767 != 833)
                        default:
                            var v15 int = fi12
                            _ = v15
                        }
                    default:
                        v6 = (v8 & (v8 - -128))
                    }
                    var v16 int64 = 1
                    fmt.Println(v16, v8)
                    switch -976 {
                    case 42:
                        fmt.Println(fi7, v8)
                    case 519:
                        fmt.Println(fi7, v8)
                    }
                    var v17 bool = true
                    v16 -= v5
                    var v18 int64 = (v16 | v5)
                    _ = v16
                    _ = v17
                    _ = v18
                }
                var v19 int64 = v5
                v5 = v5
                _ = fi12
                _ = v19
            }
        }
        _ = fi7
        _ = v8
    }
    if (-1 < (v6 ^ v6)) {
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 float32 = (v3 + v3)
            _ = fi20
            _ = v21
        }
        _ = len(v4)
        delete(v4, 973)
    } else {
        fmt.Println(v3, v3, v3)
    }
    v3 = 0.1
    v2 = append(v2, 0)
    var v22 int32 = 42
    v22 -= (v22 + v22)
    v3 = float32(0.1)
    fmt.Println(v3, v22, v5)
    if false {
        v4[32767] = (!false)
        var v23 float32 = ((1.5 * (v3 / 3.14159)) / v3)
        _ = v23
    }
    var v24 int32 = 894
    v3 = ((v3 + -2.25) * v3)
    v25 := GenS2{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
