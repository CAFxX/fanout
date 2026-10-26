package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint }
func genFn0(p0 int8, p1 string, p2 string) int64 {
    v7 := GenS1{F0: 255, F1: 1}
    switch -916 {
    case -1:
        p1 = "a"
    }
    v8 := GenS2{F0: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 4294967295
    var v3 int8 = 0
    var v4 uint8 = 42
    var v5 int32 = 42
    v6 := make(map[int]bool)
    v9 := []int{-462}
    v2 += 109
    v1 = v1
    v10 := GenS2{F0: v1}
    for fi11 := 0; fi11 < 14; fi11++ {
        fmt.Println(v2, fi11, v4)
        for fi12 := 0; fi12 < 12; fi12++ {
            v3 = 0
            _ = fi12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v2 -= 2
        _ = fi13
    }
    v14 := GenS2{F0: v1}
    var v15 uint32 = 4294967295
    v16 := GenS3{F0: 255, F1: 213, F2: 2}
    v5 = v5
    v3 = 100
    var v17 uint64 = 42
    var v18 uint16 = 0
    switch 9223372036854775807 {
    case 437:
        v19 := GenS2{F0: v1}
        _ = v19
    case 670:
        v17 = 637
    default:
        v5 += v5
    }
    var v20 int8 = 11
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
