package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 int8 = 37
    v2 := make(map[uint32]bool)
    var v3 bool = (925 != 1)
    var v4 uint = 206
    var v5 int16 = -508
    if false {
        v6 := []int{48}
        v3 = v3
        _ = v6
    } else {
        v7 := []int{520, 127, 813, -2147483648}
        _ = v7
    }
    v3 = v3
    switch "a" {
    case "a":
        if ((v4 % ((v4) | 1)) > v4) {
            fmt.Println(v4, v3)
            var v8 uint8 = 134
            switch 32768 {
            case -845:
                var v9 int32 = 2147483647
                _ = v9
            }
            _ = v8
        }
    }
    v10 := GenS1{F0: 2, F1: 2}
    var v11 string = "a"
    var v12 int32 = 42
    var v13 string = (v11 + ("" + ("foo bar" + v11)))
    if false {
        var v14 uint64 = 18446744073709551615
        var v15 float32 = float32((((0.1 * 0.1) - float64(((3.14159 - 0.0) * (float32(0.1) - 0.1)))) + 0.1))
        v13 = v11
        _ = v14
        _ = v15
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
