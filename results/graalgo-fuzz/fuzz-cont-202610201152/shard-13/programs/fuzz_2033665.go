package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint32 }
func genFn0(p0 uint8) uint16 {
    var v6 uint8 = (^(p0))
    if (false && (4294967295 > 42)) {
        v6 = 157
        v7 := []int{9223372036854775807, 48, 32767, 2147483647}
        _ = v7
    }
    _ = p0
    _ = v6
    return 32767
}

func main() {
    var v1 uint32 = 65535
    var v2 uint16 = 474
    var v3 int8 = 1
    var v4 float32 = (1e10 * float32(((3.14159 + (0.0 * 0.1)) * float64(0.0))))
    var v5 uint32 = 10
    for fi8 := 0; fi8 < 19; fi8++ {
        v5 -= (v5 << ((2) & 63))
        v9 := []int{fi8, (fi8 / ((145) | 1)), fi8, fi8}
        _ = fi8
        _ = v9
    }
    fmt.Println(v1, v3)
    var v10 uint64 = 875
    v11 := GenS1{F0: 127, F1: 9223372036854775807, F2: (v1 | (v5 - v5))}
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(v2, v5)
        switch ("" + "foo bar") {
        case "\n\t":
            v13 := []int{42, (fi12 - fi12), (^(fi12)), 9223372036854775807}
            _ = v13
        }
        fmt.Println(v5, fi12)
        _ = fi12
    }
    v3 = 0
    v5 += v1
    var v14 uint = 2
    v15 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
