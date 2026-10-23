package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    var v3 uint = 65535
    v4 := []int{128, 9223372036854775807, 1, -2147483648}
    var v5 int = 536
    v6 := GenS1{F0: 42}
    v2[("hello" + "hello")] = 2
    v3 += 255
    v7 := GenS2{F0: 4294967295}
    fmt.Println(v3, v3)
    v8 := GenS3{F0: "\n\t", F1: -701}
    v5 = 9223372036854775807
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
