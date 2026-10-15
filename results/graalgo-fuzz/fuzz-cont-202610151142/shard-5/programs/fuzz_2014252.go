package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 976
    v3 := []string{}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v4 int64 = -9223372036854775808
    fmt.Println(v4, v2)
    var v5 uint16 = 255
    fmt.Println(v5, v5, v4)
    var v6 uint8 = 63
    var v7 float64 = (((float64(0.0) * float64(1.5)) + float64(float32((1e10 - 1.5)))) + 1.5)
    for fi8 := 0; fi8 < 6; fi8++ {
        if (-724 <= 42) {
            var v9 bool = false
            switch (fi8 << ((172) & 63)) {
            case -2147483648:
                var v10 uint = 4294967295
                _ = v10
            case -1:
                var v11 uint64 = (v2 + (^(v2)))
                _ = v11
            default:
                var v12 uint64 = (v2 / ((354) | 1))
                _ = v12
            }
            _ = v9
        }
        v13 := GenS1{F0: int32(v2)}
        _ = fi8
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v4 += (v4 & (v4 - (v4 | -9223372036854775808)))
    fmt.Println(v7, v5, v2)
    v6 -= v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
