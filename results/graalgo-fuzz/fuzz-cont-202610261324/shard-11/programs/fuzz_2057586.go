package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 int8) int8 {
    p0 -= (p0 * p0)
    for fi6 := 0; fi6 < 17; fi6++ {
        p0 = p0
        _ = fi6
    }
    _ = p0
    return (p0 / ((p0) | 1))
}

func main() {
    var v1 bool = (2 <= 255)
    var v2 uint32 = 2
    var v3 int = -591
    var v4 int8 = 0
    v5 := []uint32{2, (^(v2)), (v2 + 4294967295), v2}
    v7 := GenS1{F0: 65535, F1: (42 < 1), F2: v2}
    var v8 uint = 2
    var v9 int16 = 1000
    v10 := GenS1{F0: 2147483648, F1: v1, F2: v2}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v11 int = v3
    for fi12 := 0; fi12 < 9; fi12++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v11 += v3
        _ = fi12
    }
    var v13 float32 = ((1e10 - -2.25) * (float32(0.0) + 0.1))
    var v14 int32 = 42
    v13 -= float32(0.1)
    fmt.Println(v9, v3, v8)
    var v15 uint16 = uint16(v2)
    v16 := GenS2{F0: "foo bar", F1: 0, F2: v1}
    v17 := GenS3{F0: uint8(v15), F1: v14}
    v15 = (v15 % ((32767) | 1))
    v15 -= (^(v15))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
