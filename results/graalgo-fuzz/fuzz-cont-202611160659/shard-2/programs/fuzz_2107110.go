package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint32) uint64 {
    p0 = 1
    var v5 int8 = 42
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]string)
    v3 := []int{2147483647, 0}
    v4 := make(map[int]string)
    var v6 string = "hello"
    v7 := GenS1{F0: 530, F1: 1, F2: 18446744073709551615}
    v6 = (v6 + v6)
    v8 := GenS1{F0: 255, F1: -1, F2: 2}
    v9 := GenS1{F0: 524, F1: -2147483648, F2: 255}
    v4[11] = v6
    var v10 float64 = -2.25
    v11 := GenS2{F0: v6}
    var v12 int = 79
    var v13 uint32 = 83
    v14 := GenS3{F0: v1, F1: 42}
    v15 := GenS1{F0: 65535, F1: 2147483647, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
