package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }

func main() {
    var v1 float32 = (1e10 * 3.14159)
    var v2 uint8 = 42
    var v3 float32 = v1
    var v4 int32 = -644
    var v5 int32 = v4
    fmt.Println(v4, v3, v1)
    v6 := GenS1{F0: ((127 >= 127) || (1 <= (v4 | 2147483647)))}
    v7 := []int{}
    if false {
        v8 := GenS2{F0: 65535, F1: 765, F2: 0}
        v9 := GenS2{F0: 2147483647, F1: 32768, F2: 0}
        fmt.Println(v4, v4, v2)
        _ = v8
        _ = v9
    } else {
        fmt.Println(v3)
    }
    v7 = append(v7, -85)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v5, v2)
    if (!false) {
        v5 = 293
        var v10 float32 = v1
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v1 = (float32((((-2.25 + 1e10) * 0.0) / float64(1.5))) + ((v1 + 1.5) - float32(-2.25)))
        for fi12 := 0; fi12 < 3; fi12++ {
            v2 += (v2 | (v2 % (((^(v2))) | 1)))
            var v13 uint16 = 42
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    fmt.Println(v4)
    if false {
        v5 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
