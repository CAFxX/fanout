package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS3 struct { F0 int16; F1 uint }

func main() {
    var v1 uint16 = 65535
    var v2 int32 = 409
    v3 := []int{127, 356, -2147483648}
    v4 := make(map[uint32]int)
    var v5 int8 = -1
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := v4[2147483647]
        _ = fi6
        _ = v7
    }
    switch 4294967295 {
    case -9223372036854775808:
        v2 = v2
    default:
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 int64 = -1
                var v11 float64 = 0.1
                var v12 int8 = (^(v5))
                v13 := GenS1{F0: fi8, F1: 419, F2: 32767}
                var v14 uint64 = 660
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            v15 := GenS2{F0: 77, F1: 2, F2: 0}
            v4[0] = (fi8 / ((fi8) | 1))
            _ = fi8
            _ = v15
        }
    }
    var v16 float64 = 1.5
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    if true {
        if (-1 >= -32768) {
            if (4294967295 <= 2) {
                switch -43 {
                case -319:
                    var v17 uint16 = 255
                    _ = v17
                case 0:
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 bool = (true || false)
                        if len(v3) > 0 {
                            v3[0] = (fi18 % ((9223372036854775807) | 1))
                        }
                        _ = fi18
                        _ = v19
                    }
                }
                v2 += (v2 * 42)
            } else {
                v1 = v1
                delete(v4, 0)
            }
            var v20 uint = 255
            _ = v20
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 int64 = -1
            var v23 uint16 = 303
            v24 := GenS3{F0: -32768, F1: 4294967295}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        switch 32768 {
        case 32767:
            switch 32768 {
            case -675:
                var v25 int = 42
                _ = v25
            case -2147483648:
                if ((90 >= 188) || true) {
                    v2 = v2
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v5 -= (v5 >> ((172) & 63))
                        var v27 uint64 = 42
                        _ = fi26
                        _ = v27
                    }
                }
            }
        default:
            v28 := v4[4294967295]
            _ = v28
        }
    }
    fmt.Println(v5, v16)
    v2 += 42
    fmt.Println(v16)
    v5 -= v5
    v1 = (v1 + (^(v1)))
    var v29 int8 = (v5 >> ((1) & 63))
    v16 += ((3.14159 * (v16 - (v16 - 0.0))) + float64(float32(v16)))
    v5 -= -128
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
