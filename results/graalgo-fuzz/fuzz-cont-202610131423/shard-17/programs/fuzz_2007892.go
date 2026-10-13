package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint64 }
func genFn0(p0 uint32, p1 uint16, p2 bool) int32 {
    var v6 uint32 = p0
    p2 = (2 != 127)
    v7 := []int{-9223372036854775808, 32768, 4294967295}
    p1 = (^(p1))
    var v8 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 uint32 = 163
    var v5 uint16 = 1
    var v9 uint64 = 1
    fmt.Println(v5)
    v10 := GenS1{F0: 2147483647, F1: -1}
    v4 += v4
    v11 := GenS2{F0: 1, F1: v4}
    var v12 uint = 8
    switch ("\n\t" + (("a" + "a") + "a")) {
    case "hello":
        fmt.Println(v5)
    }
    if (0.0 < float32(0.1)) {
        v13 := GenS3{F0: -15, F1: v9, F2: v9}
        _ = v13
    } else {
        v14 := []int{208, 9223372036854775807, 755}
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
