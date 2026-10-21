package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint8; F1 bool; F2 int32 }
func genFn0(p0 int64, p1 uint16) int16 {
    v7 := GenS1{F0: 1, F1: 60, F2: 44}
    if true {
        var v8 uint = 859
        var v9 uint8 = 108
        if false {
            v8 = (v8 % ((586) | 1))
        }
        _ = v8
        _ = v9
    }
    fmt.Println(p0, p0)
    v10 := GenS2{F0: p0, F1: 100}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 32767
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    v3 := []uint32{2147483648, 42}
    v4 := make(map[uint32]int)
    var v5 float64 = 0.0
    var v6 float64 = v5
    v11 := GenS1{F0: 120, F1: -128, F2: 128}
    var v12 uint16 = 503
    var v13 float32 = ((-2.25 * 1.5) - 1.5)
    _ = len(v3)
    _ = len(v2)
    fmt.Println(v1)
    var v14 string = "foo bar"
    fmt.Println(v1)
    v15 := GenS2{F0: 42, F1: -1}
    v16 := GenS3{F0: 214, F1: v1, F2: -2147483648}
    v13 -= (3.14159 - float32(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
