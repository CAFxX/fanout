package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint64 = 160
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 uint = 2
    v5 := make(map[uint32]bool)
    fmt.Println(v1, v4)
    v6 := GenS1{F0: 246, F1: v4}
    switch 32768 {
    case 0:
        v4 += (^(v4))
    case 4294967295:
        v1 -= v1
    }
    v7 := GenS1{F0: 2, F1: 911}
    fmt.Println(v1, v1)
    switch 0 {
    case 1:
        v8 := GenS1{F0: 234, F1: 0}
        _ = v8
    case -217:
        var v9 bool = false
        _ = v9
    case -538:
        fmt.Println(v1, v4)
    }
    switch 2147483647 {
    case 297:
        switch ("a" + "\n\t") {
        case "\n\t":
            v10 := []int{127, 555, 2147483647}
            _ = v10
        case "foo bar":
            v11 := []int{-1, 32767, -1}
            _ = v11
        }
    case 128:
        fmt.Println(v1, v4)
    case -9223372036854775808:
        var v12 int = 4294967295
        _ = v12
    default:
        v4 -= 18446744073709551615
    }
    v1 += (v1 - (v1 & (v1 ^ v1)))
    v4 -= 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
