package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]bool)
    var v3 uint32 = 977
    v4 := []string{"\n\t"}
    for fi5 := 0; fi5 < 8; fi5++ {
        if (128 < 58) {
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 float32 = (3.14159 * (1.5 - 1e10))
                _ = len(v2)
                v8 := GenS1{F0: -2147483648, F1: -1, F2: 2}
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 int8 = 42
            _ = v9
        }
        v1 = v1
        _ = fi5
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v3 -= 4294967295
        _ = len(v2)
        _ = fi10
    }
    var v11 int32 = 42
    var v12 int64 = -911
    _ = len(v2)
    var v13 uint64 = 1
    var v14 float64 = v1
    v4 = append(v4, "hello")
    for fi15 := 0; fi15 < 17; fi15++ {
        _ = len(v2)
        if false {
            var v16 int8 = 3
            _ = v16
        } else {
            v3 = v3
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            v14 += ((float64(float32(-2.25)) / v14) - v14)
            v3 -= 0
            _ = len(v4)
            _ = fi17
        }
        _ = fi15
    }
    if (!false) {
        var v18 uint = 42
        _ = v18
    } else {
        v19 := GenS1{F0: v11, F1: 42, F2: v3}
        var v20 int64 = v12
        _ = v19
        _ = v20
    }
    fmt.Println(v3, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
