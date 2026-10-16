package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0(p0 int64, p1 bool, p2 int) uint32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        switch (fi4 >> ((0) & 63)) {
        case 4294967295:
            p1 = p1
        case -9223372036854775808:
            var v5 int64 = (p0 % (((p0 << ((4294967295) & 63))) | 1))
            _ = v5
        }
        _ = fi4
    }
    var v6 int16 = 1000
    var v7 uint16 = 2
    var v8 int = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int8 = -66
    v2 := make(map[uint32]uint32)
    var v3 float32 = ((3.14159 + (((0.0 / 1.0) - -2.25) + -2.25)) / ((float32((0.1 + 0.0)) / -2.25) / ((float32(float64((1e10 / -2.25))) - 0.1) + ((-2.25 * (3.14159 - 1e10)) * (0.0 - -2.25)))))
    v1 = -117
    if false {
        var v9 bool = false
        delete(v2, 1)
        _ = v9
    }
    v10 := GenS1{F0: 9223372036854775807, F1: 42}
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(v3)
            v3 = 1.5
            fmt.Println(fi11, v3, v1)
            _ = fi12
        }
        var v13 uint64 = 725
        v1 -= v1
        _ = fi11
        _ = v13
    }
    _ = len(v2)
    v3 = v3
    v3 += (v3 - v3)
    v3 += (v3 - float32((-2.25 * 0.0)))
    v3 -= (float32(((((3.14159 - ((3.14159 * 3.14159) * 3.14159)) * 0.0) * float64((1.5 + (float32(float64(0.1)) - v3)))) - float64((v3 - ((v3 + (v3 * v3)) - v3))))) + (((float32(float64(1e10)) * v3) - v3) * (v3 + v3)))
    var v14 bool = (v1 <= v1)
    v1 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
}
