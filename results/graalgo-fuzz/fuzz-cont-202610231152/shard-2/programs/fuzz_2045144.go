package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint16 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    var v4 int8 = -98
    v5 := make(map[string]string)
    v1 -= v1
    v6 := GenS1{F0: ("foo bar" + "hello"), F1: 2, F2: 2}
    var v7 int = 2147483647
    v4 += v4
    var v8 string = "hello"
    v4 += (v4 ^ v4)
    v9 := GenS2{F0: 42}
    var v10 float64 = 0.1
    switch ("hello" + (v8 + v8)) {
    case "foo bar":
        var v11 uint8 = 79
        _ = v11
    }
    switch (v7 - v7) {
    case 520:
        v12 := []int{32767, v7, v7}
        _ = v12
    case 4294967295:
        fmt.Println(v1, v10)
    case 1:
        v13 := GenS3{F0: 1, F1: 0, F2: 149}
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
