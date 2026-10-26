package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int32 {
    var v5 int32 = -217
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 float64 = (1.5 / -2.25)
        var v8 uint = 2
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v5
    return -1
}
func genFn1(p0 int32, p1 int32, p2 uint16) int8 {
    fmt.Println(p2, p1, p0)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return -126
}

func main() {
    var v1 uint = 203
    var v2 float32 = 1.5
    var v3 uint32 = 2147483648
    var v4 int16 = -32768
    var v9 float64 = (((float64(v2) / ((1e10 - float64((3.14159 + float32(float64(1e10))))) - (0.0 - (0.1 * float64((v2 + (float32(1.5) + 1e10))))))) - 0.0) * float64(1e10))
    v10 := []int{127, -9223372036854775808, -302, 42}
    if (false || (!true)) {
        var v11 uint64 = 18446744073709551615
        v11 += 499
        v4 += (v4 / ((v4) | 1))
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint8 = 42
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 float64 = 0.0
        _ = fi14
        _ = v15
    }
    fmt.Println(v9, v2)
    v9 += v9
    for fi16 := 0; fi16 < 8; fi16++ {
        if len(v10) > 0 {
            v10[0] = (fi16 | (fi16 % ((fi16) | 1)))
        }
        _ = fi16
    }
    v17 := GenS1{F0: 395}
    var v18 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
