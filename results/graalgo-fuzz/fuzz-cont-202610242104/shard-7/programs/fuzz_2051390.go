package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int16; F1 bool }
func genFn0(p0 int64, p1 bool) bool {
    var v5 int8 = 42
    v6 := []int{32768, 4294967295, 1}
    var v7 float64 = 3.14159
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (v7 <= v7)
}
func genFn1() uint8 {
    if false {
        switch "hello" {
        case "foo bar":
            var v8 uint16 = 1
            _ = v8
        }
        if (!(false || false)) {
            v9 := GenS1{F0: 332, F1: 0, F2: 491}
            _ = v9
        } else {
            var v10 uint16 = 857
            if (-617 >= -1) {
                v10 = 624
            }
            _ = v10
        }
        var v11 uint = 0
        _ = v11
    } else {
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := []int{fi12}
            for fi14 := 0; fi14 < 14; fi14++ {
                if ((false && (!true)) && false) {
                    _ = len(v13)
                    fmt.Println(fi14, fi14)
                }
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        var v15 float64 = 0.0
        _ = v15
    }
    v16 := GenS2{F0: -1, F1: 560}
    var v17 int64 = 9223372036854775807
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 uint = 0
    v2 := []uint32{550, 4294967295, 1, 904}
    var v3 int32 = -2147483648
    var v4 float64 = 3.14159
    v18 := GenS3{F0: -1, F1: true}
    fmt.Println(v3, v1)
    v19 := GenS2{F0: -1, F1: 9223372036854775807}
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v3 += v3
    v20 := GenS3{F0: 1000, F1: true}
    for fi21 := 0; fi21 < 4; fi21++ {
        v1 = 1
        _ = fi21
    }
    v4 = ((v4 / ((v4 / v4) / v4)) - float64(((0.1 * ((1.5 / ((0.0 * ((3.14159 * 1e10) + (((((1.5 / -2.25) * (1.5 - (3.14159 + 3.14159))) / 1.0) / 1.0) + (1e10 - 1.5)))) / ((-2.25 - ((((float32((v4 / (0.0 + v4))) / (3.14159 * 0.1)) + (float32(v4) + 0.1)) / 1e10) / 1.5)) - 0.0))) * 1.5)) * (3.14159 / ((((1.5 - (1e10 - 1e10)) / 1.0) / 1.0) + 1.5)))))
    if ((672 == 1) && (1 >= 0)) {
        v22 := GenS1{F0: 458, F1: (v3 / ((v3) | 1)), F2: v3}
        _ = v22
    }
    v23 := GenS3{F0: 1, F1: false}
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 float32 = (-2.25 - float32(1e10))
        _ = fi24
        _ = v25
    }
    v3 = v3
    for fi26 := 0; fi26 < 9; fi26++ {
        v3 -= v3
        v1 = (v1 << (((v1 % ((193) | 1))) & 63))
        var v27 bool = ((!false) || false)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v3 = v3
        var v29 int64 = -1
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: 42, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v30
    _ = v4
}
