package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 3.14159
    v3 := make(map[uint32]uint32)
    var v4 float64 = -2.25
    v5 := make(map[string]int)
    v6 := make(map[int]string)
    v4 = ((1e10 / 1.0) / float64(float32(v4)))
    v7 := GenS1{F0: 2147483647, F1: -9223372036854775808, F2: (458 == 42)}
    var v8 int = -9223372036854775808
    var v9 uint16 = 255
    switch "" {
    case "\n\t":
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int16 = 32767
            _ = fi10
            _ = v11
        }
    case "foo bar":
        var v12 uint64 = 1
        _ = v12
    }
    v9 = (v9 ^ (v9 % ((v9) | 1)))
    v9 -= 34
    fmt.Println(v9, v4, v4)
    for fi13 := 0; fi13 < 5; fi13++ {
        v8 += -2147483648
        var v14 int8 = -99
        var v15 uint = 4294967295
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
