package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint32 }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    var v1 uint32 = 0
    var v2 bool = ((((0.1 * ((-2.25 * (float32(((0.0 * 1e10) / 1e10)) + ((((float32(0.1) / ((1e10 / 1.0) + 1e10)) + (float32(-2.25) - 0.1)) + float32(-2.25)) + 0.1))) + 0.0)) / 1.5) - 0.0) <= (3.14159 / 1.5))
    v3 := make(map[uint32]uint32)
    var v4 bool = (1 <= -128)
    var v5 string = "\n\t"
    var v6 int64 = 1
    var v7 string = ((((("\n\t" + (v5 + v5)) + v5) + v5) + (v5 + v5)) + ("" + ((v5 + v5) + (v5 + ((("a" + ((("hello" + v5) + v5) + v5)) + "a") + "\n\t")))))
    var v8 uint32 = (v1 >> ((2) & 63))
    v5 = ("a" + "hello")
    fmt.Println(v2, v8, v1)
    v7 = (("foo bar" + (v5 + (v5 + v7))) + "hello")
    v9 := []int{-286, 725}
    switch 9223372036854775807 {
    case 598:
        var v10 float64 = 1.5
        _ = v10
    }
    var v11 uint16 = 0
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: 42, F1: v6, F2: (^(v8))}
        var v14 int8 = -1
        v15 := GenS2{F0: fi12, F1: (v14 % (((v14 ^ int8(v11))) | 1))}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int = 817
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
