package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int16, p1 string, p2 int16) uint32 {
    fmt.Println(p2, p1, p0)
    v5 := []int{399, -234, -309}
    p2 -= (p2 % ((42) | 1))
    p2 = -340
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}

func main() {
    v1 := []string{}
    var v2 float64 = float64(-2.25)
    v3 := make(map[string]bool)
    v4 := []uint32{4294967295, 13}
    v2 = v2
    v6 := GenS1{F0: -201, F1: 242, F2: 89}
    if (!((genFn0(560, ("a" + "a"), 32767) < 433) || (-2147483648 != 128))) {
        if false {
            v2 = float64(((0.0 - ((((1.5 - (0.1 - -2.25)) / 1.0) + (float32(1.5) * 0.0)) * 1e10)) - (((((0.1 + (((float32(float64(-2.25)) * -2.25) * 3.14159) + (1e10 / (float32(-2.25) / 1.5)))) + (0.0 + -2.25)) + -2.25) - (1.5 - 3.14159)) * (1.5 + 1.5))))
        }
    } else {
        var v7 float32 = -2.25
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v2 = v2
        if (1 > 30) {
            v2 += v2
            switch -866 {
            case 32768:
                for fi9 := 0; fi9 < 2; fi9++ {
                    v4 = append(v4, 0)
                    _ = fi9
                }
            }
        }
        _ = fi8
    }
    _ = len(v4)
    fmt.Println(v2)
    fmt.Println(v2)
    if (0 >= 1) {
        for fi10 := 0; fi10 < 1; fi10++ {
            if (!true) {
                v2 = v2
            }
            _ = fi10
        }
        v2 -= (v2 * (v2 / v2))
        delete(v3, "foo bar")
    } else {
        if (!false) {
            var v11 int64 = -511
            var v12 uint32 = 42
            _ = v11
            _ = v12
        } else {
            var v13 uint16 = 794
            var v14 int16 = 950
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(v2, fi15, v14)
                _ = fi15
            }
            _ = v13
            _ = v14
        }
    }
    v16 := GenS1{F0: 32767, F1: 2, F2: 0}
    v2 -= ((1e10 - (v2 * v2)) + (v2 + v2))
    v17 := GenS1{F0: -424, F1: 127, F2: 2}
    v18 := GenS2{F0: 42, F1: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
