package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int32; F2 int }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]bool)
    var v3 uint8 = 128
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint8 = 127
        _ = fi5
        _ = v6
    }
    if false {
        v1 = -762
        v7 := v2[2147483647]
        _ = v7
    }
    if true {
        var v8 uint8 = (v3 / ((v3) | 1))
        var v9 bool = false
        v10 := v2[2147483647]
        _ = v10
        _ = v8
        _ = v9
    } else {
        v3 = (v3 * 46)
        var v11 float32 = 3.14159
        if false {
            v12 := []int{110, 812}
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                var v13 uint = 0
                _ = v13
            }
            v11 -= (v11 * v11)
            _ = v12
        } else {
            v14 := GenS1{F0: 1, F1: -1}
            var v15 uint16 = 65535
            v16 := GenS2{F0: (float64(1e10) <= ((3.14159 - (3.14159 * 1e10)) / 0.1))}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v11
    }
    v17 := GenS1{F0: 128, F1: 360}
    if false {
        v1 = (v1 * v1)
        v3 = 128
    } else {
        var v18 uint16 = 0
        v18 -= v18
        v19 := GenS3{F0: 239, F1: -1, F2: 42}
        _ = v18
        _ = v19
    }
    v1 -= (-(v1))
    var v20 int16 = 1000
    v1 -= -206
    v21 := GenS3{F0: 2, F1: -709, F2: -2147483648}
    v22 := []int{42}
    if len(v22) > 0 {
        v22[0] = -752
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint64 = 4294967295
        _ = fi23
        _ = v24
    }
    v3 += (v3 - (v3 + v3))
    v3 = (v3 | v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
