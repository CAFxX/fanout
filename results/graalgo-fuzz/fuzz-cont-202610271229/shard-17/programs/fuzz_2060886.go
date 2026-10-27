package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int8 = 42
    v2 := make(map[int]bool)
    v3 := []uint32{282, 1, 1}
    v4 := []uint32{2147483648, 244, 2147483648, 1}
    var v5 int32 = -888
    var v6 int16 = 837
    if true {
        if true {
            var v7 uint = 1
            _ = v7
        } else {
            if (42 > -1) {
                v1 -= v1
                var v8 uint64 = 2
                _ = v8
            }
        }
    }
    v6 = (-(v6))
    if ((-2.25 - (1.5 + 0.0)) > ((0.0 - float32((1e10 / ((-2.25 + 0.0) / 3.14159)))) * 0.1)) {
        _ = len(v2)
        switch 2147483647 {
        case 4294967295:
            if ((((((0.0 / 1e10) + -2.25) + float32(float64(0.0))) / 1.0) == (float32(-2.25) * float32(0.0))) && false) {
                if len(v3) > 0 {
                    v3[0] = 2
                }
            }
        case -9223372036854775808:
            if (65535 <= 2147483647) {
                var v9 bool = true
                fmt.Println(v5, v9)
                v2[888] = (0.1 > 1e10)
                _ = v9
            } else {
                v10 := GenS1{F0: 65535}
                _ = v10
            }
        }
    } else {
        v11 := GenS1{F0: 2147483648}
        _ = v11
    }
    var v12 uint8 = 103
    var v13 bool = (4294967295 != 155)
    if false {
        v14 := GenS1{F0: 825}
        v15 := GenS2{F0: (-(v6)), F1: 360}
        v16 := GenS3{F0: (-(v5))}
        _ = v14
        _ = v15
        _ = v16
    } else {
        if v13 {
            fmt.Println(v13, v6)
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            fmt.Println(v5)
            fmt.Println(fi17, fi17, v12)
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: 4294967295}
                v6 += (v6 << ((4294967295) & 63))
                v5 = (v5 % ((v5) | 1))
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        var v20 uint = 595
        _ = v20
    }
    _ = len(v2)
    v12 = v12
    v21 := GenS3{F0: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
