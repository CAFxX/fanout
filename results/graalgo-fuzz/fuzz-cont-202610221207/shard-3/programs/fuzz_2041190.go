package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]bool)
    var v3 int8 = -1
    v4 := []int{9223372036854775807, 32768}
    var v5 int8 = -1
    v6 := GenS1{F0: 4294967295, F1: -32768}
    var v7 bool = false
    delete(v2, (^(v1)))
    var v8 uint32 = v1
    _ = len(v4)
    var v9 uint = 0
    switch "a" {
    case "hello":
        v7 = v7
    case "a":
        var v10 int64 = int64(v3)
        _ = v10
    }
    fmt.Println(v7)
    v7 = v7
    var v11 float32 = ((float32((-2.25 - 1.5)) - (1e10 * -2.25)) - (3.14159 + ((0.0 + 1.5) + (-2.25 / (float32(-2.25) + 1e10)))))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: v9, F1: 1000}
        var v14 int32 = 0
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
