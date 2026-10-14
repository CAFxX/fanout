package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int32 }
func genFn0() bool {
    if (-407 >= -9223372036854775808) {
        switch 1 {
        case -1:
            v7 := GenS1{F0: -1, F1: 563, F2: -32768}
            _ = v7
        }
        var v8 uint16 = 2
        _ = v8
    }
    var v9 uint32 = 65535
    v10 := []int{32768, 4294967295}
    var v11 uint8 = uint8(v9)
    _ = v10
    _ = v11
    _ = v9
    return true
}
func genFn1(p0 int8, p1 int8, p2 int32) bool {
    v12 := []int{}
    p2 += (p2 + p2)
    v13 := GenS2{F0: 255, F1: 42}
    p1 = (p0 + (p0 + p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return false
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 416
    var v3 float64 = ((((float64(float32(1.5)) - 0.0) * 0.1) + 1.5) * (0.0 * (0.0 / 3.14159)))
    var v4 string = "a"
    v5 := make(map[uint32]string)
    v6 := make(map[int]string)
    v4 = (v4 + v4)
    v4 = ((v4 + "hello") + v4)
    var v14 int64 = (v2 * v2)
    v15 := GenS1{F0: v2, F1: -504, F2: 314}
    v16 := GenS2{F0: 318, F1: 42}
    var v17 int16 = -784
    v18 := GenS3{F0: 16, F1: 1, F2: 2147483647}
    v4 = ((v4 + "\n\t") + ((v4 + (((v4 + ("hello" + v4)) + "foo bar") + (("hello" + v4) + ("a" + "hello")))) + v4))
    v19 := v1[v4]
    fmt.Println(v4)
    var v20 uint64 = 42
    v21 := GenS2{F0: uint16(v19), F1: (^(v17))}
    v14 += (v14 - v2)
    if false {
        _ = len(v5)
        v22 := GenS2{F0: 65535, F1: (v17 << ((444) & 63))}
        _ = v22
    } else {
        v2 = (v2 % (((v14 - v2)) | 1))
    }
    v4 = v4
    var v23 int = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
