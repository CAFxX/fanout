package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 42
    var v3 int16 = 301
    v4 := make(map[string]string)
    if true {
        v2 = v2
    }
    switch "foo bar" {
    case "\n\t":
        v5 := GenS1{F0: 115}
        _ = v5
    case "hello":
        v6 := []int{}
        _ = v6
    default:
        fmt.Println(v3, v3)
    }
    _ = len(v1)
    if (1 <= 134) {
        v2 -= (v2 | (v2 >> (((v2 ^ v2)) & 63)))
        fmt.Println(v3, v3)
        var v7 int8 = 100
        _ = v7
    }
    fmt.Println(v2, v2)
    var v8 uint = (v2 - (v2 & (v2 >> ((v2) & 63))))
    var v9 uint64 = 42
    fmt.Println(v9, v8)
    var v10 int = -2147483648
    var v11 int8 = 127
    if false {
        var v12 uint8 = 1
        v8 = (v8 >> ((v2) & 63))
        var v13 int = (^(v10))
        _ = v12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
