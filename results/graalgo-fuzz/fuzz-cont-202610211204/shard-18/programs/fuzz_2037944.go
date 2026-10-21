package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint }
func genFn0(p0 uint8, p1 int32) int32 {
    if (1.5 == float32(float64(1.5))) {
        var v6 float32 = 0.0
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 1, F1: 65535}
        var v9 uint32 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = (p0 >> ((550) & 63))
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 int64 = -789
    var v4 float64 = (-2.25 / float64(((float32(((0.1 * float64((float32(1e10) + ((1.5 - -2.25) / 3.14159)))) + (3.14159 + ((1e10 * 3.14159) * 0.0)))) * 0.1) + (0.0 * 0.1))))
    v5 := make(map[int]bool)
    v10 := GenS2{F0: true, F1: 614}
    var v11 uint16 = 436
    fmt.Println(v3, v3, v3)
    for fi12 := 0; fi12 < 9; fi12++ {
        v4 = 0.1
        _ = fi12
    }
    v3 = (v3 + v3)
    v13 := GenS3{F0: 722}
    v14 := []int{-9223372036854775808, 563, -609, -185}
    var v15 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
