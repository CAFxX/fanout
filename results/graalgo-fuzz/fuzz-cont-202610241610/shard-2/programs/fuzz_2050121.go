package main

import (
    "fmt"
)


func main() {
    var v1 int8 = 97
    var v2 float64 = (0.0 + (0.0 - 1e10))
    v3 := make(map[int]uint32)
    var v4 uint16 = 42
    var v5 int = 1
    var v6 uint64 = 18446744073709551615
    fmt.Println(v2, v5)
    var v7 uint8 = 244
    var v8 bool = (false && true)
    v1 = (v1 >> ((239) & 63))
    var v9 int64 = int64(v4)
    delete(v3, (^(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
