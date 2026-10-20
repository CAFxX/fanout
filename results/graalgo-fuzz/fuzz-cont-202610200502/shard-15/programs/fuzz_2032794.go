package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    var v3 float64 = (((3.14159 * -2.25) + (0.0 - 0.1)) / 1.0)
    delete(v1, 2147483647)
    fmt.Println(v3)
    var v4 float64 = (-2.25 * v3)
    fmt.Println(v4, v3)
    var v5 uint = 1
    var v6 bool = (42 == 0)
    if v6 {
        v6 = false
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi7
        }
        _ = len(v2)
        v4 += float64(((3.14159 * -2.25) / (((-2.25 * 0.0) + 0.0) - (-2.25 - 3.14159))))
    }
    v8 := GenS1{F0: 876, F1: v5}
    v9 := GenS1{F0: 1, F1: (v5 + 255)}
    v10 := GenS2{F0: 0, F1: 1}
    _ = len(v1)
    v11 := GenS3{F0: -2147483648}
    if v6 {
        var v12 uint64 = 359
        var v13 float32 = 1e10
        _ = v12
        _ = v13
    }
    v4 += v3
    var v14 uint16 = 0
    v3 = v4
    v15 := GenS2{F0: uint32(v5), F1: 1000}
    var v16 float64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
