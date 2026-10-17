package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint = 689
    var v2 float32 = 0.1
    var v3 int8 = 76
    var v4 int64 = 9223372036854775807
    var v5 int64 = v4
    v3 += (v3 ^ v3)
    var v6 float64 = (0.1 / (((-2.25 - 1.5) * (((1e10 + 1.5) * (1.5 * 0.1)) / float64(((v2 * 0.1) - (v2 + -2.25))))) + 1.5))
    var v7 int8 = v3
    var v8 int32 = 865
    fmt.Println(v2)
    v9 := GenS1{F0: 0}
    v10 := []int{-439, 2147483647}
    v2 += 1.5
    for fi11 := 0; fi11 < 3; fi11++ {
        if len(v10) > 0 {
            v10[0] = (^(fi11))
        }
        _ = fi11
    }
    v8 = (v8 + v8)
    for fi12 := 0; fi12 < 6; fi12++ {
        if true {
            var v13 uint = 1
            switch fi12 {
            case -333:
                fmt.Println(v4, v8)
            case 128:
                v2 = float32(v6)
            }
            _ = v13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        _ = len(v10)
        v1 = v1
        var v15 int8 = (^(v7))
        _ = fi14
        _ = v15
    }
    var v16 int32 = v8
    var v17 int16 = -781
    fmt.Println(v3)
    v17 -= (v17 % ((v17) | 1))
    if (127 != 255) {
        var v18 int8 = (v7 & int8(v1))
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := GenS1{F0: 140}
            var v21 uint32 = 4294967295
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v18
    }
    var v22 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
