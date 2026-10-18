package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32) uint64 {
    fmt.Println(p0)
    fmt.Println(p0)
    v7 := GenS1{F0: (p0 + p0), F1: 2147483647}
    var v8 int32 = -12
    var v9 float32 = (((((1e10 * (0.0 + 0.1)) - (0.1 - 0.0)) / (0.1 * ((3.14159 * float32(((3.14159 * 0.1) / 1.5))) * 1.5))) / (3.14159 + (((float32(3.14159) - -2.25) * (0.1 * (1.5 / -2.25))) + (((float32(((float64(-2.25) + 1.5) + 0.1)) * (0.0 + 0.0)) + 0.0) - 1e10)))) - (1e10 * 1.5))
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint64, p1 uint) uint {
    if true {
        v10 := GenS1{F0: 108, F1: -627}
        v11 := GenS1{F0: 604, F1: 663}
        p0 = (p0 - (p0 - p0))
        _ = v10
        _ = v11
    }
    p1 += 0
    fmt.Println(p1, p1)
    v12 := []int{42, -130}
    _ = p0
    _ = p1
    _ = v12
    return (^(p1))
}

func main() {
    v1 := []int{0, 128, 9223372036854775807, -1}
    var v2 uint16 = 32767
    var v3 uint8 = 42
    var v4 int8 = -52
    v5 := []uint32{}
    v6 := []int{59, 0, -2147483648, -9223372036854775808}
    v13 := GenS1{F0: 65535, F1: 1}
    fmt.Println(v4, v3)
    v5 = append(v5, 2147483647)
    for fi14 := 0; fi14 < 2; fi14++ {
        switch fi14 {
        case -948:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case -2147483648:
            v3 -= (v3 / (((v3 - v3)) | 1))
        }
        var v15 int32 = 929
        _ = fi14
        _ = v15
    }
    if false {
        if (0 <= -2147483648) {
            fmt.Println(v3, v3, v2)
            if (uint16(v3) != 934) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                v5 = append(v5, 180)
                var v16 bool = false
                for fi17 := 0; fi17 < 17; fi17++ {
                    if v16 {
                        fmt.Println(v2, v2, v2)
                        for fi18 := 0; fi18 < 11; fi18++ {
                            var v19 uint64 = uint64(v3)
                            _ = fi18
                            _ = v19
                        }
                        v20 := GenS1{F0: uint32(v2), F1: fi17}
                        _ = v20
                    }
                    v21 := GenS2{F0: -1}
                    var v22 int8 = (v4 - v4)
                    _ = fi17
                    _ = v21
                    _ = v22
                }
                _ = v16
            }
        }
    } else {
        var v23 uint8 = v3
        var v24 float64 = float64(float32((1.5 - (3.14159 * 0.0))))
        v23 = 128
        _ = v23
        _ = v24
    }
    v1 = append(v1, 743)
    var v25 int = -1
    v25 = 294
    v26 := GenS3{F0: 683, F1: int16(v4)}
    v6 = append(v6, -2147483648)
    if false {
        var v27 int32 = 42
        _ = v27
    } else {
        fmt.Println(v3, v4, v3)
    }
    var v28 uint16 = (v2 * (v2 ^ (v2 << ((18446744073709551615) & 63))))
    v3 = 0
    v2 -= (v28 << ((185) & 63))
    var v29 uint = 255
    var v30 uint8 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
