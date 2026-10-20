package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8, p1 uint8) int {
    v5 := GenS1{F0: 42, F1: -473}
    var v6 int16 = 167
    v7 := []int{}
    v8 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 uint64) string {
    v9 := []int{-2147483648, genFn0(42, 128)}
    v9 = append(v9, 547)
    _ = p0
    _ = v9
    return "hello"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 127
    var v3 uint16 = 32767
    var v4 uint16 = (v3 & (v3 + 65535))
    v4 = v4
    v3 = (v4 % (((^(v3))) | 1))
    var v10 int = v2
    var v11 uint = uint(v3)
    var v12 uint64 = 1
    var v13 int16 = 369
    v2 -= v10
    v14 := GenS3{F0: (v13 / ((int16(v2)) | 1))}
    fmt.Println(v11, v10, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
