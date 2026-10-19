package main

import (
    "fmt"
)


func main() {
    var v1 int = 128
    v2 := make(map[string]uint32)
    var v3 float32 = 3.14159
    v4 := make(map[string]uint32)
    v5 := make(map[int]int)
    var v6 uint64 = 42
    fmt.Println(v6, v1, v1)
    var v7 uint32 = 2
    switch "foo bar" {
    case "hello":
        fmt.Println(v6, v7)
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 float32 = v3
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 127
    var v12 float64 = 1e10
    v13 := []int{-1, 9223372036854775807, (v1 ^ v1), (v1 >> ((18446744073709551615) & 63))}
    var v14 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
