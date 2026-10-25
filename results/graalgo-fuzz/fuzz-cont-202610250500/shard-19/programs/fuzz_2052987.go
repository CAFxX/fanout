package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 65535
    var v2 uint8 = 0
    var v3 float64 = (-2.25 - 0.1)
    v4 := []int{}
    v3 -= 1e10
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: -927, F1: -1, F2: 9223372036854775807}
        var v7 uint8 = v2
        if (1.5 == 0.1) {
            var v8 uint32 = 2147483648
            var v9 int16 = 0
            _ = v8
            _ = v9
        } else {
            var v10 float32 = ((float32(0.1) - 0.1) - 1e10)
            _ = v10
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v11 int8 = -5
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(fi12, fi12)
        fmt.Println(v2)
        v1 += (^(v1))
        _ = fi12
    }
    v2 = (^(v2))
    if false {
        fmt.Println(v3, v11)
        var v13 string = "\n\t"
        _ = v13
    } else {
        v14 := GenS1{F0: 2147483647, F1: -974, F2: 1}
        var v15 uint16 = 291
        v11 = (v11 + (-(v11)))
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v2 -= v2
        _ = fi16
    }
    v17 := GenS2{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
