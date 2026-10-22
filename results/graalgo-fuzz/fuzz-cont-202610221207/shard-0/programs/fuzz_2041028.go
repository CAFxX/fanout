package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 int64; F2 int }
func genFn0(p0 int8) uint16 {
    v7 := []int{617, -1, 58, 4294967295}
    if (0.1 >= 3.14159) {
        p0 += p0
        var v8 float64 = ((1.5 + 1e10) + (0.1 * (3.14159 * 1.5)))
        var v9 float32 = (0.0 + 0.0)
        _ = v8
        _ = v9
    }
    p0 -= (p0 % (((^(p0))) | 1))
    _ = p0
    _ = v7
    return 1000
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    v3 := []string{"", ""}
    v4 := make(map[string]string)
    var v5 float32 = ((-2.25 / 1.5) * -2.25)
    var v6 uint64 = 1
    _ = len(v4)
    var v10 bool = (1 <= 65535)
    switch int(v6) {
    case 0:
        if (v6 == (v6 >> ((707) & 63))) {
            for fi11 := 0; fi11 < 10; fi11++ {
                if (1 <= 42) {
                    v6 += (v6 >> ((0) & 63))
                } else {
                    fmt.Println(v5, v5, fi11)
                    delete(v4, (("" + "\n\t") + ("a" + "")))
                    var v12 uint8 = 128
                    _ = v12
                }
                v10 = v10
                v13 := GenS1{F0: 32767}
                _ = fi11
                _ = v13
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 int8 = -128
                v15 += (v15 - v15)
                _ = fi14
                _ = v15
            }
        } else {
            v5 -= (3.14159 + (v5 * ((0.1 + (float32(float64(1e10)) - float32(float64(0.1)))) + 1.5)))
        }
    case 32768:
        v16 := GenS2{F0: (v6 % ((18446744073709551615) | 1)), F1: 0}
        _ = v16
    case -856:
        var v17 int32 = 2147483647
        _ = v17
    default:
        v18 := GenS3{F0: 2147483648}
        _ = v18
    }
    var v19 int64 = -9223372036854775808
    v5 = v5
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 float32 = v5
            v5 = v5
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    var v23 uint8 = 0
    v24 := GenS4{F0: (v6 - v6), F1: int64(v6), F2: -2147483648}
    if (255 >= v23) {
        switch -1 {
        case 127:
            if v10 {
                var v25 int64 = (v19 / (((v19 * 454)) | 1))
                v23 -= v23
                _ = v25
            } else {
                var v26 float32 = 1.5
                _ = v26
            }
        case -9223372036854775808:
            var v27 float64 = 1e10
            _ = v27
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
