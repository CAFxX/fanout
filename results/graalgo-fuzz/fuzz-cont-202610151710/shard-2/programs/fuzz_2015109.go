package main

import (
    "fmt"
)


func main() {
    var v1 float32 = (float32(-2.25) + 1e10)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    v1 -= (-2.25 - v1)
    fmt.Println(v1)
    v1 = v1
    v5 := []int{239, 2147483647}
    v1 = 1e10
    switch "hello" {
    case "hello":
        switch 162 {
        case 287:
            var v6 uint32 = 972
            _ = v6
        case 32767:
            v4["a"] = 874
        case -9223372036854775808:
            var v7 int8 = 42
            _ = v7
        default:
            var v8 uint = 0
            _ = v8
        }
    case "a":
        v1 += 3.14159
    default:
        v1 = v1
    }
    var v9 float32 = (v1 / -2.25)
    v1 = (0.1 + float32(((1.5 / 1e10) * 0.1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
