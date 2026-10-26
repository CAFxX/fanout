package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    v3 := []string{"a", "foo bar", ("" + "\n\t")}
    var v4 uint64 = 653
    v5 := make(map[uint32]bool)
    v4 = 1
    var v6 float32 = (float32(-2.25) - float32(1e10))
    if false {
        v7 := GenS1{F0: 42, F1: (v4 ^ v4)}
        v6 += v6
        var v8 uint8 = 197
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1000, F1: (v4 & v4)}
    _ = len(v3)
    fmt.Println(v4, v2, v6)
    var v10 int16 = -32768
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v4 -= 42
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 uint64 = 4294967295
            v10 = (v10 % ((v10) | 1))
            var v14 uint16 = 102
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
