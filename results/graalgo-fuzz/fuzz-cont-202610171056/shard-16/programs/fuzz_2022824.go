package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int32 }

func main() {
    v1 := []int{1, -609, 127}
    v2 := []uint32{2147483647, 50}
    v3 := []uint32{65535, 742}
    v4 := make(map[int]uint32)
    var v5 int8 = 97
    v6 := make(map[int]bool)
    v7 := GenS1{F0: 42, F1: 2}
    var v8 int16 = 1
    v9 := GenS2{F0: true}
    v5 += (v5 | 42)
    v10 := GenS1{F0: 0, F1: 65535}
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v8 -= (^(v8))
            var v13 int8 = v5
            _ = fi12
            _ = v13
        }
        v5 = -119
        if true {
            fmt.Println(fi11, v8)
        }
        _ = fi11
    }
    v14 := GenS3{F0: 77, F1: 0}
    var v15 int16 = -1
    v16 := GenS3{F0: 4294967295, F1: 0}
    var v17 float32 = 1.5
    v17 = v17
    v2 = append(v2, 2147483647)
    v17 -= (v17 + (((0.1 - ((((0.1 - v17) - v17) + v17) * v17)) + ((v17 - (3.14159 + 0.1)) + v17)) + ((v17 + (v17 * v17)) + v17)))
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := v6[fi18]
        var v20 uint16 = 816
        _ = fi18
        _ = v19
        _ = v20
    }
    v8 = v15
    v5 = 0
    _ = len(v4)
    for fi21 := 0; fi21 < 3; fi21++ {
        fmt.Println(v8, v8)
        if (false && ("foo bar" < "hello")) {
            var v22 uint64 = 42
            _ = v22
        } else {
            v23 := v4[fi21]
            _ = v23
        }
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
