package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint16 }
func genFn0(p0 int32) int {
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            p0 -= p0
            if (!(2 > 2147483647)) {
                if ((3.14159 > -2.25) || true) {
                    if true {
                        v7 := GenS1{F0: 2, F1: false, F2: 0}
                        _ = v7
                    }
                    fmt.Println(fi5, fi5, p0)
                } else {
                    v8 := GenS2{F0: 1, F1: 0}
                    _ = v8
                }
                switch fi6 {
                case 42:
                    var v9 int8 = 100
                    _ = v9
                }
                var v10 int64 = -9223372036854775808
                _ = v10
            } else {
                var v11 int8 = 1
                p0 = p0
                _ = v11
            }
            v12 := GenS3{F0: 601, F1: 255}
            _ = fi6
            _ = v12
        }
        _ = fi5
    }
    var v13 float64 = 1.5
    var v14 int16 = 42
    switch 831 {
    case 1:
        v14 = v14
    }
    if (-332 <= 2147483647) {
        v13 = (1.5 * v13)
        var v15 int = -2147483648
        var v16 uint8 = 127
        _ = v15
        _ = v16
    }
    _ = p0
    _ = v13
    _ = v14
    return -1
}

func main() {
    var v1 int32 = 286
    v2 := []string{"foo bar", "a", "a", "a"}
    var v3 uint8 = 126
    var v4 uint16 = 399
    v17 := GenS4{F0: 42, F1: 91}
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 float64 = 1.5
        if (139 > v4) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v3 -= (v3 / ((0) | 1))
        }
        _ = fi18
        _ = v19
    }
    v20 := GenS3{F0: 18446744073709551615, F1: (v4 - (v4 | 32767))}
    v3 = v3
    var v21 uint = 367
    var v22 int = -1
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 int8 = 127
        v25 := GenS1{F0: uint64(v21), F1: false, F2: v21}
        v26 := GenS2{F0: 42, F1: 2147483647}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    switch ((("" + "\n\t") + "\n\t") + "") {
    case "hello":
        v27 := GenS1{F0: 835, F1: false, F2: (v21 / (((v21 << (((v21 / (((^(v21))) | 1))) & 63))) | 1))}
        _ = v27
    case "":
        v4 = 837
    case "\n\t":
        var v28 int64 = -573
        _ = v28
    default:
        v4 += 715
    }
    v4 = uint16(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
