package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 537
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v5 string = "\n\t"
    v2 = 2147483647
    if true {
        switch 32767 {
        case 1:
            fmt.Println(v1, v5, v1)
        case 42:
            v6 := GenS1{F0: v2}
            _ = v6
        default:
            var v7 uint64 = v1
            _ = v7
        }
        var v8 uint = 943
        _ = v8
    } else {
        var v9 uint32 = (v2 >> ((2) & 63))
        _ = v9
    }
    fmt.Println(v1, v5)
    fmt.Println(v2, v1, v2)
    fmt.Println(v2, v2)
    v10 := []int{-9223372036854775808, 9223372036854775807}
    _ = len(v4)
    fmt.Println(v1)
    v3[v2] = v2
    var v11 int = 584
    v1 = (^(v1))
    if len(v10) > 0 {
        v10[0] = -129
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
