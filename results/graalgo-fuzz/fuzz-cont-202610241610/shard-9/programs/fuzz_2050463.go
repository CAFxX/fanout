package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint = 659
    v2 := []uint32{418, 4294967295, 42}
    v3 := make(map[int]bool)
    var v4 int8 = 123
    var v5 float32 = 1e10
    v5 = 1e10
    v6 := GenS1{F0: 2}
    switch "\n\t" {
    case "a":
        switch "" {
        case "\n\t":
            var v7 uint8 = 2
            _ = v7
        }
    case "\n\t":
        v4 = v4
    }
    var v8 int32 = 1
    fmt.Println(v5)
    v1 -= 248
    v3[128] = (!true)
    v8 += (-(v8))
    if (232 < 255) {
        delete(v3, -390)
    }
    if false {
        v8 += v8
        fmt.Println(v1, v4)
    }
    v5 -= v5
    var v9 int8 = 42
    v10 := GenS2{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
