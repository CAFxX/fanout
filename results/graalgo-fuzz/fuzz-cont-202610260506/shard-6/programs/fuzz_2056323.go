package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 bool }

func main() {
    var v1 int = 128
    var v2 uint64 = 1
    var v3 float32 = 0.0
    if (0.1 == ((1e10 + (0.0 + (1e10 + -2.25))) - 0.0)) {
        fmt.Println(v2)
        var v4 uint8 = 0
        var v5 int32 = -166
        _ = v4
        _ = v5
    }
    v3 = v3
    switch "foo bar" {
    case "\n\t":
        var v6 int8 = 0
        _ = v6
    case "foo bar":
        var v7 uint8 = 0
        _ = v7
    case "a":
        v8 := []int{(^(v1)), v1, v1, v1}
        _ = v8
    default:
        var v9 int16 = -659
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(fi10, v2)
        _ = fi10
    }
    fmt.Println(v3)
    var v11 int32 = -247
    v11 = (v11 * 440)
    v12 := GenS1{F0: "\n\t"}
    if (-1 == (^(v11))) {
        v3 += v3
    }
    v13 := GenS2{F0: 65535, F1: (true || false)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
