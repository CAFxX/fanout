package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = ((0.0 * (1.5 * ((1.5 + 3.14159) * 0.1))) * ((3.14159 - -2.25) - (((1.5 * (1e10 + 1e10)) * -2.25) + 1.5)))
    var v3 uint32 = 0
    var v4 int = 0
    var v5 float64 = ((1.5 / 1e10) - 0.0)
    var v6 float32 = float32(v5)
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 = (float32(0.0) / -2.25)
        _ = fi7
    }
    var v8 int16 = 42
    v3 -= (^(v3))
    v4 = v4
    v9 := []int{(-(v4)), v4, (v4 % (((v4 << ((4294967295) & 63))) | 1)), (v4 >> ((1) & 63))}
    v10 := GenS1{F0: -640, F1: v4, F2: 255}
    var v11 float64 = 3.14159
    v6 = (v6 - (float32((-2.25 - v5)) - ((v6 / v6) / v6)))
    v12 := GenS2{F0: "foo bar", F1: -2147483648}
    switch ("" + "") {
    case "foo bar":
        if (v2 > 0.0) {
            v2 = 3.14159
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS2{F0: "", F1: -1}
                var v15 uint64 = 18446744073709551615
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            if true {
                v11 += 3.14159
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
