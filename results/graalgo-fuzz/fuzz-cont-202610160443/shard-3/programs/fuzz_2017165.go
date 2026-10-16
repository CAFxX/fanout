package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0() bool {
    var v5 uint16 = 181
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(v5)
        var v7 int64 = 27
        _ = fi6
        _ = v7
    }
    v5 = v5
    v8 := []int{128}
    _ = v5
    _ = v8
    return true
}
func genFn1() uint64 {
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint16 = 32767
        var v11 int8 = -1
        var v12 bool = (false || ((false || (4294967295 > 18446744073709551615)) || ((-1 >= 0) && false)))
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 string = "hello"
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 bool = false
        var v16 uint64 = 42
        if true {
            var v17 uint8 = 42
            _ = v17
        } else {
            var v18 uint16 = 364
            for fi19 := 0; fi19 < 20; fi19++ {
                v15 = false
                fmt.Println(v13, v15, v15)
                var v20 uint8 = 42
                _ = fi19
                _ = v20
            }
            v18 += (v18 + v18)
            _ = v18
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = v13
    return 0
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{0, 2147483648, 4294967295, 966}
    v3 := make(map[int]string)
    var v4 float32 = (1e10 - (-2.25 - 1e10))
    v21 := GenS1{F0: 227, F1: 1000}
    v1 = 407
    if true {
        if true {
            v22 := GenS2{F0: (v1 ^ 90), F1: 32767}
            delete(v3, 551)
            _ = v22
        }
    } else {
        v23 := GenS3{F0: 0, F1: 2147483647, F2: (v1 & (v1 << ((854) & 63)))}
        var v24 uint64 = 4294967295
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 float32 = -2.25
        var v27 uint32 = 42
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v1, v1)
    v28 := GenS2{F0: (v1 - (v1 ^ 65535)), F1: v1}
    if len(v2) > 0 {
        v2[0] = 800
    }
    var v29 int32 = -2147483648
    var v30 float64 = 0.0
    v4 -= (((v4 / (-2.25 / v4)) * float32((v30 / (v30 + float64(float32(-2.25)))))) + (float32(0.0) - v4))
    v4 += v4
    v1 -= (v1 << ((0) & 63))
    v31 := v3[1]
    var v32 uint = 1
    v31 = (v31 + (v31 + (v31 + "hello")))
    if false {
        _ = len(v3)
    }
    v32 -= (^(v32))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
