package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 uint64, p1 string) uint32 {
    v5 := GenS1{F0: 790, F1: -829}
    var v6 float32 = 1e10
    v7 := GenS1{F0: 0, F1: 1}
    switch p1 {
    case "\n\t":
        v8 := GenS2{F0: 253}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int32) int32 {
    fmt.Println(p0)
    var v9 int16 = 42
    _ = p0
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2147483648, 906, 179, 254}
    var v4 float64 = (-2.25 - (0.0 - -2.25))
    var v10 int64 = 42
    v10 = 0
    v4 -= ((0.1 - (v4 - v4)) - (1.5 + (((((float64(0.1) + v4) - -2.25) * float64((0.1 + 0.1))) + 0.1) / 1.0)))
    v4 = v4
    v10 = v10
    delete(v1, 395)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v11 int64 = (v10 % (((^(v10))) | 1))
    var v12 float64 = 0.0
    var v13 int16 = 42
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int32 = -87
        v4 += 0.1
        var v16 float64 = (v12 / (1.5 + v12))
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v13, v10, v4)
    if ((float32(v4) * (1.5 + (1e10 - float32(v4)))) > 0.0) {
        v10 = v10
        delete(v1, 4294967295)
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            if (160 <= 128) {
                var v18 uint = 647
                var v19 uint16 = 65535
                _ = v18
                _ = v19
            }
            if (-9223372036854775808 != (fi17 >> ((4294967295) & 63))) {
                v20 := GenS3{F0: 0, F1: v11}
                v21 := GenS3{F0: 2, F1: v10}
                if (743 == 1) {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        var v23 float32 = 3.14159
                        if ((v13 >> ((458) & 63)) <= 379) {
                            var v24 uint = 4294967295
                            _ = v24
                        } else {
                            v25 := GenS3{F0: 66, F1: (v11 % ((-1) | 1))}
                            var v26 uint = 4294967295
                            _ = v25
                            _ = v26
                        }
                        v11 = (v11 / (((v11 | v10)) | 1))
                        _ = fi22
                        _ = v23
                    }
                    v13 += v13
                    v13 = 32767
                }
                _ = v20
                _ = v21
            }
            _ = fi17
        }
    }
    v11 = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
