package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 int32 }
type GenS3 struct { F0 int16 }
func genFn0() string {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{}
        var v8 uint32 = 4294967295
        if false {
            var v9 float64 = 1e10
            _ = v9
        } else {
            if false {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v7 = append(v7, fi6)
                v10 := GenS1{F0: 255, F1: 1}
                _ = v10
            }
            var v11 int64 = 76
            var v12 int32 = 2147483647
            _ = v11
            _ = v12
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int = 42
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 0, F1: 1, F2: 313}
    _ = v15
    return "\n\t"
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 uint8 = 165
    v4 := []int{}
    var v5 int32 = 0
    var v16 int16 = -62
    v3 = (v3 / (((v3 * (v3 ^ v3))) | 1))
    v5 = 2147483647
    if (!((v16 << ((235) & 63)) >= 1000)) {
        var v17 string = ("hello" + "foo bar")
        var v18 float32 = float32((1e10 / 1.0))
        var v19 uint16 = 576
        _ = v17
        _ = v18
        _ = v19
    } else {
        v16 = (v16 + v16)
        var v20 string = "a"
        _ = v20
    }
    v21 := v1[-2147483648]
    var v22 int16 = v16
    v23 := GenS3{F0: -728}
    var v24 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
