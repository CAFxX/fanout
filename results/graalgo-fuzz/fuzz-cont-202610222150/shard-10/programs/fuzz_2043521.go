package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }

func main() {
    var v1 uint16 = 0
    var v2 int8 = 0
    var v3 uint = 18446744073709551615
    v4 := []int{1, -9223372036854775808}
    var v5 int8 = -128
    if len(v4) > 0 {
        v4[0] = -552
    }
    v6 := GenS1{F0: "\n\t", F1: -2147483648}
    v2 -= v5
    v7 := GenS1{F0: "\n\t", F1: int(v1)}
    var v8 string = "\n\t"
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v2 = (-(v5))
            v3 = (^(v3))
            v11 := GenS1{F0: "foo bar", F1: -43}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 string = v8
    v5 = v5
    var v13 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
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
