package main

import (
    "fmt"
)


func main() {
    v1 := make(map[string]bool)
    var v2 int64 = -404
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 0
    var v5 uint = 255
    var v6 uint = (v5 ^ v5)
    v2 = (v2 >> ((v6) & 63))
    v4 -= 303
    v7 := []int{128, 2147483647, 1, 127}
    var v8 int32 = 172
    var v9 uint16 = 0
    var v10 uint64 = 220
    fmt.Println(v4)
    var v11 uint8 = 141
    switch "a" {
    case "hello":
        v2 += v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
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
