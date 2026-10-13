package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint8 }
func genFn0(p0 bool, p1 uint8) bool {
    fmt.Println(p0)
    v5 := []int{-335, 4294967295, -2147483648}
    if (true && false) {
        switch "foo bar" {
        case "\n\t":
            p0 = (true || false)
        case "a":
            v6 := GenS1{F0: 238, F1: 154}
            _ = v6
        }
    }
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = v5
    return (594 != 366)
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int32 = 729
    var v3 bool = false
    v4 := make(map[uint32]int)
    var v7 uint16 = 42
    if (0.1 < (0.1 - (1.5 + 3.14159))) {
        v8 := []int{}
        v9 := GenS1{F0: 829, F1: 1}
        v10 := GenS2{F0: "\n\t"}
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int64 = 9223372036854775807
    v12 := v4[v1]
    v13 := GenS3{F0: 2147483647, F1: 2}
    v1 += 65535
    v14 := []int{-1}
    v1 = (v1 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
