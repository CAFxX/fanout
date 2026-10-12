package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 1
    var v3 int32 = -1
    var v4 int16 = -656
    var v5 int8 = 42
    v6 := GenS1{F0: 0}
    var v7 uint8 = 37
    switch 9223372036854775807 {
    case -1:
        v8 := v1[int(v2)]
        _ = v8
    }
    v9 := GenS1{F0: v7}
    fmt.Println(v3)
    if true {
        fmt.Println(v3, v5)
        var v10 uint8 = (v7 & (v7 % ((v7) | 1)))
        _ = v10
    }
    var v11 uint8 = v7
    var v12 int8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
