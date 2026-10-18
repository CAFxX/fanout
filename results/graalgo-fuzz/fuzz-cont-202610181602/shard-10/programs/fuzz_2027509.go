package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int32; F1 int8 }

func main() {
    v1 := []int{-9223372036854775808, 42}
    var v2 int8 = 0
    v3 := make(map[string]int)
    var v4 int16 = -1
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: (v2 | (v2 | -1)), F1: int64(v4), F2: -128}
    var v7 uint = 695
    v7 = (^(v7))
    _ = len(v3)
    v8 := GenS1{F0: v2, F1: -226, F2: (v2 >> ((v7) & 63))}
    if (int8(v4) == v2) {
        v7 += (^(v7))
    }
    var v9 int64 = 9223372036854775807
    for fi10 := 0; fi10 < 17; fi10++ {
        if false {
            fmt.Println(v4, v4, v9)
            v11 := GenS2{F0: 131, F1: 9}
            _ = v11
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int16 = v4
            v1 = append(v1, fi12)
            _ = fi13
            _ = v14
        }
        v15 := GenS3{F0: int32(fi12), F1: v2}
        _ = fi12
        _ = v15
    }
    fmt.Println(v9, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
