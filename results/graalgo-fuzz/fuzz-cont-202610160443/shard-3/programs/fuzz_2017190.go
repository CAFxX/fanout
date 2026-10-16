package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint; F2 string }
func genFn0(p0 uint8, p1 uint64, p2 bool) uint32 {
    v6 := []int{1, 2147483647, -981}
    for fi7 := 0; fi7 < 20; fi7++ {
        p1 = (p1 - p1)
        fmt.Println(p2, p1)
        _ = fi7
    }
    var v8 int64 = 496
    p2 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return 1
}

func main() {
    var v1 int8 = -128
    var v2 int64 = -517
    var v3 uint8 = 128
    v4 := []uint32{797}
    var v5 float32 = 0.1
    if (0 <= 32767) {
        var v9 int8 = 0
        v10 := GenS1{F0: 51, F1: (v9 % (((v9 * v1)) | 1))}
        _ = v10
        _ = v9
    } else {
        if false {
            v5 -= v5
            v4 = append(v4, genFn0(v3, 168, false))
            var v11 uint16 = 2
            _ = v11
        }
        v12 := GenS2{F0: -1}
        _ = v12
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v2 = (v2 * v2)
    var v13 int16 = -32768
    v14 := GenS3{F0: (^(v3))}
    var v15 int8 = v1
    var v16 string = ""
    v17 := GenS4{F0: -9223372036854775808, F1: 4294967295, F2: ""}
    var v18 string = "hello"
    if false {
        v19 := GenS3{F0: 63}
        _ = v19
    }
    v13 += -32768
    v20 := GenS4{F0: 0, F1: 255, F2: v16}
    fmt.Println(v15, v5, v5)
    v1 -= int8(v2)
    if false {
        v13 += v13
        var v21 uint = 4294967295
        if len(v4) > 0 {
            v4[0] = 1
        }
        _ = v21
    } else {
        var v22 int32 = 42
        v23 := GenS1{F0: -752, F1: v15}
        _ = len(v4)
        _ = v22
        _ = v23
    }
    var v24 int = 32768
    var v25 uint16 = 263
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
