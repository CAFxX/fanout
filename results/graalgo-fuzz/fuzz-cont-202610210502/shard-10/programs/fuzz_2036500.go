package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
func genFn0(p0 bool, p1 uint8, p2 string) uint16 {
    v6 := []int{42, 0}
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 string = "hello"
    var v2 bool = true
    v3 := make(map[int]int)
    var v4 int16 = -1
    v5 := make(map[int]uint32)
    v7 := []int{-2147483648}
    v4 += (v4 * v4)
    v8 := GenS1{F0: 658, F1: 128}
    var v9 int32 = 42
    var v10 uint16 = 832
    var v11 bool = v2
    var v12 int32 = (v9 / (((-(v9))) | 1))
    v7 = append(v7, -23)
    delete(v5, -731)
    var v13 uint32 = 65535
    v9 = v9
    v2 = (!v2)
    var v14 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
