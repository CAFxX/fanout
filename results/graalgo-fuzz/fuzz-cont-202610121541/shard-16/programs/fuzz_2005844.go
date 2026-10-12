package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 uint; F2 uint }
func genFn0(p0 int8, p1 uint8, p2 bool) int32 {
    v6 := []int{-2147483648}
    var v7 string = "hello"
    var v8 int16 = -282
    for fi9 := 0; fi9 < 20; fi9++ {
        if p2 {
            _ = len(v6)
        } else {
            fmt.Println(v8, fi9, v8)
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int16 = -32768
        _ = len(v6)
        var v12 int16 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint, p1 int64) uint16 {
    fmt.Println(p1, p1)
    v13 := GenS1{F0: 0}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v13
    return 503
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 42
    var v5 uint = 0
    fmt.Println(v1)
    v4 = v4
    v2[4294967295] = 886
    v14 := []int{-767, -1, 32767, int(v4)}
    fmt.Println(v5, v1)
    var v15 uint = v5
    var v16 uint = v5
    var v17 float64 = (((0.0 + (0.0 + -2.25)) / float64(float32(1.5))) - 3.14159)
    switch 4294967295 {
    case -2147483648:
        v18 := GenS2{F0: 128, F1: 339}
        _ = v18
    case 799:
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    case 4294967295:
        var v19 int8 = 42
        _ = v19
    default:
        _ = len(v14)
    }
    v5 = v15
    var v20 int64 = 0
    v21 := GenS3{F0: 1, F1: v15, F2: v5}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
