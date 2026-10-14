package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 string }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int64 = 42
    var v2 int8 = 80
    v3 := make(map[string]uint32)
    var v4 int = 429
    v4 += (v4 / ((42) | 1))
    var v5 float32 = -2.25
    if true {
        fmt.Println(v5, v1, v5)
    } else {
        switch v4 {
        case 42:
            if ((-1 <= -749) && (1000 <= 0)) {
                v2 -= 1
                if (false && false) {
                    var v6 int8 = v2
                    _ = v6
                }
            } else {
                v5 = float32(0.1)
            }
        case 9223372036854775807:
            v7 := GenS1{F0: (v2 >> ((255) & 63)), F1: 255}
            _ = v7
        case 127:
            switch "a" {
            case "":
                v8 := GenS1{F0: (-(v2)), F1: 622}
                _ = v8
            default:
                for fi9 := 0; fi9 < 8; fi9++ {
                    v4 += v4
                    var v10 float64 = (3.14159 + (3.14159 - 1e10))
                    _ = fi9
                    _ = v10
                }
            }
        }
    }
    v4 -= 909
    switch v4 {
    case -1:
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := []int{v4}
                v2 = (v2 * (v2 / (((v2 ^ (v2 << ((42) & 63)))) | 1)))
                v13 = append(v13, fi11)
                _ = fi12
                _ = v13
            }
            var v14 int64 = v1
            _ = fi11
            _ = v14
        }
    case -584:
        v1 += v1
    case 4294967295:
        v5 += (v5 * (0.0 * v5))
    default:
        switch (v4 + v4) {
        case -773:
            var v15 uint16 = 350
            _ = v15
        case 128:
            var v16 int8 = (v2 / (((v2 + v2)) | 1))
            _ = v16
        }
    }
    v1 = (v1 & (-(v1)))
    v17 := GenS2{F0: 328, F1: (-(v1)), F2: "\n\t"}
    for fi18 := 0; fi18 < 8; fi18++ {
        v5 += (v5 - (((v5 / 1.0) * v5) + v5))
        if (18446744073709551615 != 0) {
            if (149 == 42) {
                var v19 uint = 0
                _ = v19
            } else {
                var v20 float32 = float32(1e10)
                if (!false) {
                    _ = len(v3)
                    fmt.Println(v5, v4, v1)
                }
                for fi21 := 0; fi21 < 8; fi21++ {
                    if ((-(v1)) > (v1 + v1)) {
                        var v22 float32 = v5
                        _ = v22
                    }
                    var v23 int32 = -1
                    _ = fi21
                    _ = v23
                }
                _ = v20
            }
        } else {
            v1 -= v1
            v2 += (v2 & -128)
            var v24 int64 = v1
            _ = v24
        }
        _ = fi18
    }
    var v25 uint32 = 496
    v26 := GenS3{F0: (v1 << ((255) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
