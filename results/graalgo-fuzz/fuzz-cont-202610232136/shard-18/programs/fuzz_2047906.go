package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 bool, p1 uint64) bool {
    p1 = 429
    if p0 {
        p1 = (p1 >> ((4294967295) & 63))
    }
    _ = p0
    _ = p1
    return false
}
func genFn1(p0 uint32) uint8 {
    v5 := []int{881, 429, -2147483648, 127}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 128
}

func main() {
    var v1 int16 = 1
    var v2 uint64 = 0
    var v3 int16 = v1
    var v4 bool = false
    v6 := GenS1{F0: 568, F1: 706, F2: 255}
    fmt.Println(v2, v3)
    if v4 {
        var v7 bool = (1 != 85)
        v2 = (v2 - (v2 & v2))
        for fi8 := 0; fi8 < 7; fi8++ {
            if (1 != 0) {
                var v9 int = fi8
                _ = v9
            }
            _ = fi8
        }
        _ = v7
    }
    v2 = v2
    var v10 float64 = 0.0
    v11 := GenS2{F0: (v1 >> ((65535) & 63))}
    if (!v4) {
        v12 := GenS2{F0: 1000}
        _ = v12
    }
    if (v2 < v2) {
        var v13 int64 = 0
        v10 = v10
        _ = v13
    } else {
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS1{F0: 66, F1: 9223372036854775807, F2: 1}
                v10 = (0.1 * 1.5)
                v17 := []int{317, fi15, fi14}
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 uint64 = v2
            _ = fi14
            _ = v18
        }
    }
    fmt.Println(v1, v4)
    v3 += (v3 % (((v3 + v1)) | 1))
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
