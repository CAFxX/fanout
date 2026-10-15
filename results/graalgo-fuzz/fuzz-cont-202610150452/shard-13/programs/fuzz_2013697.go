package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int {
    var v5 uint16 = 0
    v6 := []int{-1}
    v5 = (v5 % (((v5 * 469)) | 1))
    _ = v5
    _ = v6
    return 14
}

func main() {
    var v1 uint16 = 574
    v2 := make(map[string]uint32)
    var v3 uint8 = 42
    var v4 int64 = -1
    v7 := []int{9223372036854775807, 0, 127, 127}
    v8 := GenS1{F0: int8(v1)}
    v9 := GenS1{F0: 127}
    v4 -= (v4 + 0)
    fmt.Println(v4)
    var v10 int64 = v4
    if (65535 < 769) {
        var v11 uint = 65535
        fmt.Println(v10, v10, v4)
        _ = v11
    }
    v10 -= 424
    v10 = 42
    v4 = (v10 >> ((255) & 63))
    switch "foo bar" {
    case "foo bar":
        v12 := GenS1{F0: 100}
        _ = v12
    }
    fmt.Println(v4, v1, v10)
    if true {
        var v13 float32 = 1e10
        v14 := GenS1{F0: 1}
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
