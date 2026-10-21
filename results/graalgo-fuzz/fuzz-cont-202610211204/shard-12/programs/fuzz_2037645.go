package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }
func genFn0(p0 int8, p1 int32, p2 int8) bool {
    var v4 int8 = 0
    var v5 float32 = float32(0.1)
    p0 += (p2 - 112)
    v5 = v5
    fmt.Println(v4, p0, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (!false)
}
func genFn1(p0 uint8, p1 bool, p2 uint64) int {
    if ((!p1) && (-634 <= 1)) {
        var v6 int16 = 32767
        v7 := GenS1{F0: 520}
        v6 = 310
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        p1 = false
        for fi9 := 0; fi9 < 7; fi9++ {
            p0 = (p0 | (p0 ^ (^(p0))))
            v10 := GenS2{F0: 127, F1: p1, F2: 100}
            v11 := []int{-9223372036854775808, fi9, fi8}
            var v12 int16 = -295
            v12 -= (v12 >> ((18446744073709551615) & 63))
            v13 := GenS2{F0: 1, F1: p1, F2: 1}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    p1 = true
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := []uint32{109, 107}
    var v2 float64 = 3.14159
    v3 := make(map[int]uint32)
    var v14 uint = 225
    if (genFn1(240, true, 1) > genFn1(42, false, 129)) {
        switch 9223372036854775807 {
        case -2147483648:
            switch 2147483647 {
            case -1:
                switch 32767 {
                case 128:
                    for fi15 := 0; fi15 < 16; fi15++ {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        var v16 uint = v14
                        if true {
                            v16 = (v16 + (v14 % (((v16 >> ((2) & 63))) | 1)))
                        }
                        _ = fi15
                        _ = v16
                    }
                case 42:
                    if true {
                        var v17 int64 = -677
                        _ = v17
                    }
                }
            case 32768:
                v18 := GenS1{F0: 0}
                _ = v18
            case 2147483647:
                v3[722] = 975
            }
        }
    } else {
        var v19 int64 = 42
        _ = v19
    }
    v20 := GenS3{F0: 0, F1: 9223372036854775807, F2: 42}
    v2 += 1e10
    if false {
        v14 += v14
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            if (0.1 != 0.1) {
                var v22 uint64 = 18446744073709551615
                _ = v22
            }
            _ = fi21
        }
        v14 = 943
        fmt.Println(v14)
    }
    var v23 float32 = (1e10 + ((float32(float64(float32(float64(0.0)))) * 3.14159) - (0.1 - 0.1)))
    var v24 int32 = -99
    v23 = ((v23 - v23) * 0.1)
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := v3[(fi25 << (((^(v14))) & 63))]
        fmt.Println(v23)
        v27 := GenS1{F0: uint32(v14)}
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
}
