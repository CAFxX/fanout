package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"foo bar", "\n\t", ""}
    var v3 float64 = ((1e10 + 0.1) / 0.1)
    v4 := make(map[uint32]uint32)
    v5 := []string{(("foo bar" + "foo bar") + (("foo bar" + "a") + "")), "hello", "hello", ("\n\t" + "")}
    v3 = v3
    v6 := GenS1{F0: 197, F1: 1}
    var v7 float64 = (v3 - 1.5)
    switch -9223372036854775808 {
    case 1:
        var v8 uint8 = 255
        _ = v8
    default:
        var v9 int32 = -2147483648
        _ = v9
    }
    var v10 int8 = 1
    v10 += v10
    v7 = v3
    v7 = v7
    var v11 uint32 = 2147483647
    fmt.Println(v10, v7, v3)
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 float32 = 3.14159
        _ = fi12
        _ = v13
    }
    var v14 int8 = v10
    for fi15 := 0; fi15 < 9; fi15++ {
        v1[v11] = (-(fi15))
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
