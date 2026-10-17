package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }

func main() {
    var v1 uint8 = 0
    var v2 float64 = (3.14159 - 0.0)
    var v3 float32 = 1.5
    var v4 int64 = -1
    var v5 int = -1
    v6 := make(map[int]string)
    fmt.Println(v1, v4)
    v7 := GenS1{F0: 255, F1: "foo bar"}
    _ = len(v6)
    var v8 uint8 = v1
    fmt.Println(v2)
    fmt.Println(v8, v3)
    v4 = v4
    v8 += (v8 >> ((175) & 63))
    switch (v5 + (v5 << ((42) & 63))) {
    case 617:
        var v9 uint64 = 18446744073709551615
        _ = v9
    case -356:
        v1 = v1
    case 947:
        v10 := v6[(v5 % ((v5) | 1))]
        _ = v10
    default:
        v6[(v5 * (v5 >> ((138) & 63)))] = ("hello" + "")
    }
    var v11 uint8 = (v8 * v1)
    v11 = (v11 >> ((304) & 63))
    if ((!true) || false) {
        v12 := GenS1{F0: 448, F1: "hello"}
        _ = v12
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
