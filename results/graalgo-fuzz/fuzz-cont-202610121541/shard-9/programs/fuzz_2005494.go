package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 1
    var v3 int16 = 1
    v4 := []int{547, 9223372036854775807, 127, 42}
    var v5 int = 2147483647
    var v6 float64 = (0.0 - 1e10)
    fmt.Println(v3, v3, v3)
    v6 = v6
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int16 = (v3 | int16(v2))
        var v9 int32 = 222
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 1}
    switch "foo bar" {
    case "hello":
        var v11 int16 = -96
        _ = v11
    }
    var v12 int8 = 42
    fmt.Println(v6)
    fmt.Println(v3, v12)
    fmt.Println(v6, v2)
    var v13 uint32 = 1
    fmt.Println(v2, v3, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
