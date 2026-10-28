package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 100
    var v3 bool = false
    v4 := make(map[string]uint32)
    var v5 uint = 984
    var v6 int32 = 0
    var v7 int64 = 946
    v7 = v7
    v8 := []int{}
    var v9 string = "hello"
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v5, v3, v9)
        _ = fi10
    }
    v9 = "foo bar"
    var v11 int64 = v7
    v12 := GenS1{F0: 623, F1: 464}
    v5 = 148
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint64 = 18446744073709551615
        fmt.Println(v14)
        _ = fi13
        _ = v14
    }
    v3 = (255 > 838)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
