package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{978}
    var v3 int = 32768
    v3 -= -2147483648
    v3 += (-(v3))
    var v4 uint8 = 255
    var v5 int32 = -333
    switch "a" {
    case "foo bar":
        fmt.Println(v4, v3)
    case "hello":
        v6 := GenS1{F0: 42}
        _ = v6
    case "a":
        v5 -= (v5 << ((334) & 63))
    }
    delete(v1, (v3 << ((80) & 63)))
    var v7 uint64 = 505
    var v8 uint16 = 1
    v2 = append(v2, 42)
    v9 := GenS2{F0: 617}
    var v10 int64 = -944
    delete(v1, 128)
    var v11 bool = (!false)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
