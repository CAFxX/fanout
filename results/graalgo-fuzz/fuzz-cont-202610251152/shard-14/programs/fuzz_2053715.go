package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
func genFn0(p0 uint32, p1 uint32, p2 int64) int {
    if false {
        var v5 int64 = (p2 / ((-1) | 1))
        _ = v5
    }
    v6 := []int{127, -2147483648}
    var v7 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 605
}
func genFn1(p0 uint16) int64 {
    p0 -= p0
    v8 := []int{}
    _ = p0
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 float32 = (0.1 - (-2.25 - (0.0 + 0.1)))
    v4 := make(map[uint32]string)
    v9 := GenS1{F0: 65535}
    v10 := []int{0, -2147483648}
    v11 := GenS2{F0: 42}
    if len(v10) > 0 {
        v10[0] = 42
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int64 = 42
        var v14 int64 = (v13 * (^(v13)))
        fmt.Println(v13, fi12, v2)
        _ = fi12
        _ = v13
        _ = v14
    }
    switch -2147483648 {
    case -346:
        v15 := GenS2{F0: 127}
        _ = v15
    }
    var v16 int32 = 1
    v17 := GenS3{F0: 1, F1: v2, F2: 71}
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(v3, v3, v16)
        fmt.Println(fi18, v2, fi18)
        var v19 uint16 = 32767
        _ = fi18
        _ = v19
    }
    v16 -= (v16 % ((-1) | 1))
    fmt.Println(v2, v16, v2)
    var v20 uint32 = 360
    v21 := GenS1{F0: v20}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
