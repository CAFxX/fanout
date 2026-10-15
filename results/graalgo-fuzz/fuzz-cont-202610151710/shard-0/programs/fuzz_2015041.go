package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int64) int64 {
    v5 := []int{}
    switch 9223372036854775807 {
    case 552:
        v6 := GenS1{F0: 980, F1: 1}
        _ = v6
    }
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    if true {
        _ = len(v5)
        v7 := GenS1{F0: 270, F1: 42}
        if len(v5) > 0 {
            v5[0] = 32768
        }
        _ = v7
    }
    _ = p0
    _ = v5
    return (p0 << ((255) & 63))
}
func genFn1(p0 uint, p1 int16) int8 {
    var v8 int16 = p1
    fmt.Println(v8, p1, p0)
    _ = p0
    _ = p1
    _ = v8
    return -103
}

func main() {
    var v1 int8 = -123
    v2 := make(map[int]bool)
    var v3 int = 309
    var v4 uint64 = 42
    v1 = v1
    var v9 int8 = (v1 | (v1 << ((255) & 63)))
    fmt.Println(v1, v3, v9)
    var v10 uint = 210
    v11 := []int{}
    fmt.Println(v4, v10)
    var v12 int8 = (v9 / (((v1 >> (((v10 % ((v10) | 1))) & 63))) | 1))
    v4 = v4
    fmt.Println(v4)
    _ = len(v11)
    if (-58 <= 0) {
        if true {
            v4 += (v4 % ((v4) | 1))
            var v13 uint8 = 241
            _ = v13
        }
        v12 += genFn1((v10 % ((v10) | 1)), -32768)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
