package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint64, p2 uint64) bool {
    if (p2 >= p1) {
        v5 := GenS1{F0: 0}
        v6 := []int{32768, -2147483648}
        var v7 bool = (379 > -1)
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 uint16 = 255
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint8 = 127
        v11 := GenS1{F0: v10}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return false
}

func main() {
    v1 := []uint32{1, 2147483648, 2147483648, 2}
    var v2 int = 9223372036854775807
    var v3 int = 2147483647
    var v4 uint8 = 111
    v3 += (v3 ^ (v2 ^ v3))
    v3 += v2
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint16 = 1
        var v14 float32 = ((-2.25 / float32(0.1)) / ((0.0 - 3.14159) / 3.14159))
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int16 = 0
    v15 -= (v15 % ((-1) | 1))
    var v16 uint32 = 715
    v17 := GenS1{F0: v4}
    var v18 float64 = ((float64(float32(1.5)) - 0.0) / 0.1)
    for fi19 := 0; fi19 < 1; fi19++ {
        fmt.Println(v4, v16)
        _ = fi19
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
