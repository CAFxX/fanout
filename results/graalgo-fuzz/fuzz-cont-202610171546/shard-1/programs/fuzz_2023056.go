package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint {
    var v4 int16 = 42
    v5 := []int{-2147483648}
    _ = v4
    _ = v5
    return 905
}

func main() {
    var v1 float64 = 0.1
    v2 := []int{-2147483648, 9223372036854775807, 32767}
    var v3 int16 = -32768
    var v6 uint32 = 967
    var v7 uint16 = 1
    fmt.Println(v1)
    if false {
        for fi8 := 0; fi8 < 10; fi8++ {
            v1 = v1
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint32 = (v6 << ((4294967295) & 63))
                _ = fi9
                _ = v10
            }
            fmt.Println(v3)
            _ = fi8
        }
    } else {
        v11 := GenS1{F0: 2147483647}
        _ = v11
    }
    v1 += (v1 - (float64(0.0) / v1))
    fmt.Println(v1)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint16 = v7
        _ = fi12
        _ = v13
    }
    var v14 int32 = 7
    if false {
        var v15 int = -2147483648
        var v16 float32 = 1e10
        if len(v2) > 0 {
            v2[0] = v15
        }
        _ = v15
        _ = v16
    } else {
        v1 += float64(((0.0 / float32(v1)) + ((1.5 / ((0.1 / (((0.0 / -2.25) + (0.0 * (float32(0.0) - 0.1))) * float32(((v1 - (-2.25 - (v1 / v1))) * 0.0)))) * (0.1 - 3.14159))) - 3.14159)))
        if len(v2) > 0 {
            v2[0] = -313
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
