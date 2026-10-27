package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint32 = 0
    var v2 uint = 478
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1.5
    var v5 string = "\n\t"
    var v6 float64 = v4
    v7 := GenS1{F0: 4294967295, F1: (v1 - (v1 ^ 4294967295))}
    var v8 uint = (^(v2))
    v8 += 18446744073709551615
    v5 = (((v5 + v5) + v5) + "foo bar")
    fmt.Println(v4, v5, v6)
    var v9 uint = (v8 | (v2 % ((v2) | 1)))
    var v10 int8 = 127
    if ((0.0 / 1e10) >= 3.14159) {
        v2 += uint(v1)
        switch (v5 + v5) {
        case "\n\t":
            var v11 float32 = (0.0 * -2.25)
            _ = v11
        default:
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := []int{fi12, (-(fi12))}
                _ = fi12
                _ = v13
            }
        }
        var v14 string = ("foo bar" + v5)
        _ = v14
    }
    v15 := []int{-637, -1}
    if (-128 >= 81) {
        var v16 float32 = 3.14159
        var v17 uint8 = 130
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
