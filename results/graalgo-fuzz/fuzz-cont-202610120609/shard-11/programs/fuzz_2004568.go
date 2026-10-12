package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]int)
    var v3 uint8 = 42
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{uint32(v4), 0, 2}
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: v3, F1: 300}
        fmt.Println(v4, fi6, v4)
        _ = fi6
        _ = v7
    }
    v1 = v1
    var v8 int32 = -249
    v3 += v3
    switch 9223372036854775807 {
    case 127:
        if true {
            v1 -= v1
        }
    }
    v3 -= (v3 ^ v3)
    v9 := GenS2{F0: -128}
    v3 += 0
    v10 := v2[((("" + "\n\t") + "") + ((("hello" + "") + (("\n\t" + "\n\t") + "a")) + ("a" + ((("a" + (("\n\t" + ("" + ("a" + ""))) + ("a" + ""))) + "") + "\n\t"))))]
    var v11 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
