package main

import (
    "fmt"
)


func main() {
    var v1 uint64 = 888
    v2 := []uint32{}
    v3 := make(map[string]string)
    var v4 int32 = 0
    var v5 int32 = 462
    switch "foo bar" {
    case "":
        v5 = -2147483648
    case "hello":
        fmt.Println(v4, v5, v5)
    case "a":
        v5 += 42
    }
    var v6 int = -564
    var v7 uint64 = (v1 | 42)
    var v8 uint16 = 2
    var v9 int16 = 1000
    v7 += (v1 << ((4294967295) & 63))
    var v10 float64 = (float64((-2.25 * 0.1)) / ((1.5 / ((0.1 * (3.14159 * (-2.25 + 0.0))) * 3.14159)) / 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
