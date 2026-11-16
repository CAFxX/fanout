package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int8 }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 1e10
    var v3 uint16 = 1
    v4 := []int{-1, -484, -915}
    v5 := []uint32{2147483647, 0}
    var v6 float64 = 1.5
    v2 = v2
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 65535, F1: 704, F2: 42}
        v3 -= 255
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := v1["\n\t"]
            var v11 int8 = 42
            v5 = append(v5, 907)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    var v12 uint64 = 18446744073709551615
    fmt.Println(v6, v6, v2)
    _ = len(v1)
    v13 := GenS1{F0: 65535, F1: 855, F2: -18}
    var v14 int32 = -1
    if (((v3 & (v3 & (v3 * v3))) >= v3) && (!false)) {
        v3 -= 1
        v6 = float64(1e10)
        var v15 int64 = 0
        _ = v15
    }
    var v16 int64 = -936
    v17 := GenS2{F0: -109, F1: 2}
    var v18 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
