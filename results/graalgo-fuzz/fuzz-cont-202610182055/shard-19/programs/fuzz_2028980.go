package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0(p0 int, p1 uint32) int8 {
    v5 := GenS1{F0: (0 <= 1), F1: 65535}
    v6 := GenS2{F0: 672, F1: 2, F2: -2147483648}
    v7 := []int{p0, (p0 - p0)}
    var v8 int8 = 1
    p0 = 127
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 bool = true
    v1 = (32767 > -364)
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS3{F0: -2147483648, F1: -2147483648}
        v4 = (v4 && true)
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 1, F1: 42, F2: 42}
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := []int{-433, -1, fi12}
        fmt.Println(v1, fi12, v1)
        _ = fi12
        _ = v13
    }
    v1 = false
    fmt.Println(v4, v1)
    fmt.Println(v4, v4)
    for fi14 := 0; fi14 < 16; fi14++ {
        v1 = (-32768 >= 42)
        _ = fi14
    }
    v4 = (9223372036854775807 <= -1)
    var v15 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
