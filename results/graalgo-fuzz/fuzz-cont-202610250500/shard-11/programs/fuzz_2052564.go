package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64, p1 int8) uint {
    v6 := GenS1{F0: p1, F1: 0, F2: (2 != 147)}
    var v7 uint16 = 32767
    var v8 int16 = 713
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int16 = -809
    var v2 bool = (32767 == 9223372036854775807)
    var v3 int8 = 100
    var v4 float64 = 1e10
    var v5 int8 = v3
    v9 := GenS2{F0: -32768}
    if ((!(v2 && ((!(v2 && true)) || ((v3 != (-(v3))) && v2)))) && false) {
        var v10 int64 = 9223372036854775807
        _ = v10
    } else {
        var v11 float32 = (3.14159 + float32(0.1))
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := []int{(fi12 & -981), (^(fi13)), 128}
                _ = fi13
                _ = v14
            }
            var v15 uint8 = 42
            var v16 int16 = v1
            _ = fi12
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 float32 = v11
            for fi19 := 0; fi19 < 15; fi19++ {
                v1 -= (^(v1))
                fmt.Println(v2, v5)
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
        _ = v11
    }
    v20 := []int{4294967295}
    v21 := GenS3{F0: (v1 << ((1) & 63))}
    fmt.Println(v4, v5)
    v5 += -1
    if v2 {
        v5 -= (v3 % ((v5) | 1))
        var v22 uint = 757
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 int64 = 655
        fmt.Println(fi23, v4)
        _ = fi23
        _ = v24
    }
    var v25 int64 = 9223372036854775807
    switch -812 {
    case -659:
        v4 += (((v4 / (((0.0 * ((float64(1e10) + ((0.1 / v4) - float64((float32((v4 / v4)) + 1.5)))) * v4)) + ((v4 * (1e10 - (((float64(-2.25) / 1e10) + 1.5) / (float64((((((3.14159 + 1.5) + (0.0 + ((1.5 * float32((1.5 + v4))) * float32((((v4 + v4) * (v4 * 1e10)) - -2.25))))) + (1e10 / float32(v4))) * ((-2.25 - 1.5) + 0.0)) + 0.0)) / (v4 + 3.14159))))) / (v4 / 3.14159))) * v4)) - ((v4 * v4) / 3.14159)) / 3.14159)
    case -140:
        var v26 string = "\n\t"
        _ = v26
    case -489:
        var v27 int8 = -128
        _ = v27
    }
    v3 += -1
    _ = len(v20)
    v3 += (v3 | 127)
    fmt.Println(v5)
    v3 += (v5 >> ((1) & 63))
    for fi28 := 0; fi28 < 9; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            v25 = (v25 | v25)
            var v30 uint = 65535
            fmt.Println(v2)
            v20 = append(v20, (^(fi28)))
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    var v31 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
