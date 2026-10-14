package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 float32 = (((1.5 / 1e10) * ((0.0 / -2.25) * float32(3.14159))) + (float32(((3.14159 - 0.0) + (1.5 + (-2.25 - (-2.25 / -2.25))))) * (((((-2.25 + (0.1 * (float32((((float64((3.14159 - 1.5)) - float64(1e10)) - (float64((-2.25 + 1e10)) * 1.5)) * ((0.1 - float64(-2.25)) + float64(0.1)))) + 3.14159))) + (((-2.25 * ((0.0 - 3.14159) + (float32((float64(0.0) / float64(-2.25))) * 1e10))) * float32(1.5)) + (((float32(-2.25) + (0.1 / 1e10)) / 1.0) + (((1.5 / 1.0) + 0.1) / -2.25)))) / -2.25) - 1.5) * 1.5)))
    v2 := []uint32{}
    var v3 uint8 = 2
    v1 = v1
    v3 -= v3
    for fi4 := 0; fi4 < 7; fi4++ {
        fmt.Println(fi4)
        var v5 bool = (!true)
        _ = fi4
        _ = v5
    }
    if false {
        for fi6 := 0; fi6 < 8; fi6++ {
            switch fi6 {
            case 42:
                fmt.Println(fi6, v1, v3)
            case -9223372036854775808:
                v1 += 3.14159
            }
            fmt.Println(v3, v1, v3)
            var v7 int32 = 2147483647
            _ = fi6
            _ = v7
        }
        var v8 uint64 = 0
        _ = v8
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        if len(v2) > 0 {
            v2[0] = 65535
        }
        v3 = v3
        _ = fi9
    }
    v10 := GenS1{F0: 2, F1: 0}
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint8 = 127
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
