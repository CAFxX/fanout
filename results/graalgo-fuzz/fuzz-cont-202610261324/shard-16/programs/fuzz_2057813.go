package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 int8 = -24
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]string)
    v4 := []int{-9223372036854775808, -608, 9223372036854775807, -98}
    v5 := GenS1{F0: 18446744073709551615, F1: -577}
    var v6 int16 = 1
    var v7 uint16 = 0
    var v8 string = "foo bar"
    switch "foo bar" {
    case "foo bar":
        for fi9 := 0; fi9 < 14; fi9++ {
            v6 = v6
            v6 += v6
            _ = fi9
        }
    case "":
        v1 += (v1 + (-(v1)))
    case "\n\t":
        var v10 uint64 = 18446744073709551615
        _ = v10
    }
    delete(v3, "hello")
    v11 := GenS2{F0: (v7 | (^(v7))), F1: 8}
    var v12 int64 = v2
    v13 := GenS2{F0: (^(v7)), F1: v7}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
