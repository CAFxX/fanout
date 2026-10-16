package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int32 = 944
    var v2 int8 = -12
    var v3 int = 1
    v1 = -701
    v1 = 0
    v4 := []int{}
    var v5 uint8 = 76
    v6 := GenS1{F0: 975, F1: (v1 % ((v1) | 1))}
    v2 -= 0
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 float32 = 0.0
        v5 = v5
        fmt.Println(v1, fi7, v3)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            if ((182 == 2) || (0.1 > (3.14159 + 1.5))) {
                v1 += -1
                var v11 int16 = 32767
                v12 := GenS1{F0: 2, F1: v1}
                _ = v11
                _ = v12
            } else {
                v1 = 903
            }
            _ = fi10
        }
        _ = fi9
    }
    fmt.Println(v2, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
