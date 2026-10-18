package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []int{63, -486, -9223372036854775808, 512}
    var v2 uint16 = 518
    var v3 int32 = 42
    v4 := make(map[int]int)
    var v5 uint = 255
    var v6 bool = false
    var v7 uint16 = (^(v2))
    v8 := GenS1{F0: int16(v3), F1: 854}
    var v9 int32 = v3
    if (1e10 >= (3.14159 * 1e10)) {
        v2 -= v2
        v7 += v7
    }
    var v10 uint = 18446744073709551615
    fmt.Println(v9, v2)
    fmt.Println(v10, v3, v10)
    switch ("hello" + ("a" + "a")) {
    case "a":
        fmt.Println(v2, v5)
    case "foo bar":
        v11 := GenS2{F0: 0}
        _ = v11
    case "\n\t":
        v7 += v2
    }
    var v12 uint8 = 102
    var v13 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
