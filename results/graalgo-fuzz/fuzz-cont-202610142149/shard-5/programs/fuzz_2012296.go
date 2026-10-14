package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0(p0 int16) uint8 {
    if true {
        v6 := []int{32767}
        p0 -= (p0 / (((-(p0))) | 1))
        _ = v6
    } else {
        v7 := []int{2147483647, 263}
        var v8 uint = 255
        var v9 float32 = float32(1.5)
        _ = v7
        _ = v8
        _ = v9
    }
    if (false && false) {
        var v10 uint64 = 18446744073709551615
        v11 := GenS1{F0: 127, F1: p0}
        _ = v10
        _ = v11
    } else {
        var v12 int = 62
        _ = v12
    }
    _ = p0
    return 20
}
func genFn1(p0 uint8, p1 int, p2 uint32) string {
    p2 = (p2 % (((p2 + p2)) | 1))
    v13 := []int{p1, 32767, p1, (p1 >> ((42) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return ""
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 float32 = ((0.0 * (float32((float64(((0.0 - 0.1) * 1e10)) * (1e10 + ((1.5 * float64(-2.25)) - -2.25)))) * 0.0)) - (((3.14159 * 1.5) / 0.1) / 1.5))
    v5 := make(map[uint32]uint32)
    fmt.Println(v4, v1)
    v4 = v4
    v4 -= (1e10 / v4)
    fmt.Println(v1, v4)
    v4 = 0.0
    var v14 int64 = 9223372036854775807
    for fi15 := 0; fi15 < 9; fi15++ {
        if true {
            var v16 bool = (-2147483648 > 1)
            v14 -= -9223372036854775808
            for fi17 := 0; fi17 < 13; fi17++ {
                fmt.Println(v16)
                _ = fi17
            }
            _ = v16
        }
        fmt.Println(fi15, v14)
        v4 = (v4 - ((v4 * ((v4 / v4) * 1e10)) / v4))
        _ = fi15
    }
    var v18 float32 = v4
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v4)
        switch (^(fi19)) {
        case 32768:
            _ = len(v3)
        case 641:
            var v20 uint32 = 289
            _ = v20
        }
        var v21 float32 = float32(float64(v4))
        _ = fi19
        _ = v21
    }
    v4 = (v18 / (1.5 * 3.14159))
    var v22 uint32 = 42
    if false {
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 uint = 65535
            var v25 uint32 = (v22 + (v22 + (v22 >> (((v24 + v24)) & 63))))
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
