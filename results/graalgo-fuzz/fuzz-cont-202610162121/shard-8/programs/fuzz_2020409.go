package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 uint16 = 0
    var v4 float64 = 1.5
    v5 := make(map[string]string)
    var v6 uint16 = (v3 << ((42) & 63))
    v7 := []int{}
    var v8 string = ""
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 uint = 0
        v11 := GenS1{F0: 32767}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = len(v1)
    if (1 > 5) {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint64 = 297
            _ = fi12
            _ = v13
        }
    }
    v6 += v6
    var v14 int16 = 102
    if (!false) {
        v14 -= (v14 / ((1) | 1))
        v15 := GenS1{F0: v14}
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 uint64 = 42
            v18 := GenS2{F0: 42}
            fmt.Println(v17, v17)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    } else {
        v6 += v3
        var v19 float64 = (v4 * v4)
        _ = v19
    }
    v8 = v8
    var v20 uint32 = 1
    v21 := GenS3{F0: -1, F1: 670, F2: 100}
    v3 = (v6 % (((v6 & v6)) | 1))
    v4 = (0.0 + v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
