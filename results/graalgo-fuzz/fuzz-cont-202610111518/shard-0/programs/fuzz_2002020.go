package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }

func main() {
    var v1 int16 = -32768
    var v2 int16 = 0
    v3 := make(map[uint32]uint32)
    var v4 uint = 26
    var v5 int32 = 2147483647
    v4 += (^(v4))
    var v6 float64 = 1.5
    v7 := GenS1{F0: 2147483647, F1: v2, F2: true}
    var v8 string = ""
    for fi9 := 0; fi9 < 11; fi9++ {
        _ = len(v3)
        _ = fi9
    }
    v5 += (v5 / ((-816) | 1))
    v5 += (v5 | v5)
    if (127 != 255) {
        if true {
            var v10 int = 843
            _ = v10
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint32 = 4294967295
            v13 := []int{}
            fmt.Println(v2, v12, fi11)
            _ = fi11
            _ = v12
            _ = v13
        }
        v4 -= v4
    }
    if true {
        v3[577] = 0
        v14 := GenS1{F0: 42, F1: v2, F2: (1 > 0)}
        for fi15 := 0; fi15 < 5; fi15++ {
            if ((v1 + -48) <= v1) {
                v1 = v2
            }
            var v16 float32 = 0.1
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
