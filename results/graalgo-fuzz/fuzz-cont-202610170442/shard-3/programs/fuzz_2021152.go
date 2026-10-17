package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
func genFn0() uint {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int = -1
        v7 = v7
        if false {
            v8 := GenS1{F0: 2147483647, F1: "\n\t"}
            _ = v8
        } else {
            var v9 float32 = (0.1 + float32((float64(-2.25) * (float64(float32((((((1.5 - 0.0) + 0.0) * ((3.14159 + (1e10 / 1.5)) - 0.0)) / ((((((0.0 - ((0.0 - 3.14159) + 1e10)) - 0.1) / (1.5 / (1e10 * 1.5))) - 1.5) / ((0.1 + 0.1) - 1.5)) - (0.0 - (3.14159 - 3.14159)))) - -2.25))) / (1.5 - 0.0)))))
            v10 := []int{v7, fi6, (v7 ^ fi6)}
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v11 int8 = -128
    v11 = -49
    _ = v11
    return 999
}

func main() {
    v1 := []uint32{65535, 42, 65535}
    var v2 uint32 = 2
    var v3 int32 = 0
    v4 := []uint32{}
    v5 := []uint32{v2, v2, (^(v2))}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v3, v2)
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: (v3 ^ v3), F1: "\n\t"}
        for fi14 := 0; fi14 < 19; fi14++ {
            v2 -= v2
            var v15 int64 = 42
            _ = fi14
            _ = v15
        }
        switch (fi12 - -9223372036854775808) {
        case -2147483648:
            v3 = v3
        default:
            var v16 int64 = int64(fi12)
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    var v17 uint64 = 817
    fmt.Println(v2, v3, v17)
    if len(v1) > 0 {
        v1[0] = v2
    }
    if (!(false && false)) {
        var v18 uint64 = 1
        fmt.Println(v18, v3, v17)
        _ = v18
    }
    v1 = append(v1, 551)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
