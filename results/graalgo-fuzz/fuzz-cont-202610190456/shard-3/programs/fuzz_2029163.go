package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int64; F2 int32 }

func main() {
    v1 := []int{187, 2147483647, -2147483648}
    var v2 uint = 18446744073709551615
    var v3 float64 = 1.5
    v4 := make(map[string]int)
    v5 := []int{1, 42, 32768, -861}
    v6 := []uint32{65535, 238}
    fmt.Println(v3, v2)
    for fi7 := 0; fi7 < 15; fi7++ {
        v2 = (v2 % (((v2 >> (((v2 & (^(v2)))) & 63))) | 1))
        for fi8 := 0; fi8 < 1; fi8++ {
            if ((!false) || (-1 < 455)) {
                switch (fi7 << (((v2 << (((v2 % ((v2) | 1))) & 63))) & 63)) {
                case -107:
                    v3 += -2.25
                }
                fmt.Println(fi7, v3)
                v9 := GenS1{F0: -613}
                _ = v9
            }
            if len(v6) > 0 {
                v6[0] = 0
            }
            var v10 uint = v2
            if (-706 < -1) {
                fmt.Println(fi8, v2)
                var v11 float64 = v3
                v11 = (1e10 + 0.1)
                _ = v11
            }
            var v12 bool = false
            _ = fi8
            _ = v10
            _ = v12
        }
        _ = fi7
    }
    v13 := GenS1{F0: 1}
    for fi14 := 0; fi14 < 2; fi14++ {
        switch ("\n\t" + "hello") {
        case "\n\t":
            v2 -= (v2 & v2)
        case "hello":
            switch (^(fi14)) {
            case 32768:
                if true {
                    var v15 float32 = ((0.1 - 1e10) - 0.1)
                    var v16 uint = (v2 - (v2 % ((2) | 1)))
                    _ = v15
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        if (-241 > 0) {
                            var v18 uint = 18446744073709551615
                            v19 := v4[((("\n\t" + ("" + "hello")) + "foo bar") + "")]
                            _ = v18
                            _ = v19
                        } else {
                            var v20 int64 = -520
                            _ = len(v4)
                            _ = v20
                        }
                        _ = fi17
                    }
                }
            case -2147483648:
                var v21 int64 = 321
                _ = v21
            default:
                var v22 int16 = 552
                _ = v22
            }
        default:
            var v23 int32 = 1
            _ = v23
        }
        _ = fi14
    }
    if false {
        for fi24 := 0; fi24 < 17; fi24++ {
            v2 += 1
            _ = fi24
        }
        _ = len(v4)
        v3 -= ((1e10 * (3.14159 * float64(1e10))) - (-2.25 / (float64(3.14159) / v3)))
    } else {
        v25 := GenS1{F0: 1}
        _ = v25
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v26 := GenS2{F0: 2147483647, F1: 151, F2: -976}
    var v27 uint8 = 112
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
