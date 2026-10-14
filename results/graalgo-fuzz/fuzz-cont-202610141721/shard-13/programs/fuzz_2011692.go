package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0() uint32 {
    var v5 int64 = 120
    v6 := GenS1{F0: (v5 / ((1) | 1)), F1: 65535}
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -123
    v3 := []int{4294967295, 1, -973, 416}
    var v4 float32 = float32(0.0)
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v4 -= -2.25
            var v9 float32 = (v4 - (v4 + (1.5 * 3.14159)))
            var v10 uint32 = genFn0()
            _ = fi8
            _ = v10
            _ = v9
        }
        if (true && (fi7 != fi7)) {
            v11 := GenS1{F0: (v2 << ((255) & 63)), F1: 32767}
            switch fi7 {
            case 1:
                delete(v1, ("a" + "foo bar"))
            case 42:
                for fi12 := 0; fi12 < 13; fi12++ {
                    v1["\n\t"] = (("" + "\n\t") + "foo bar")
                    v4 = ((float32(((((1e10 + (0.1 * (0.0 / 1.5))) * (float64(v4) - float64(v4))) * 0.1) * float64(v4))) / v4) + 1e10)
                    _ = fi12
                }
            case -2147483648:
                fmt.Println(fi7)
            default:
                v13 := GenS1{F0: (v2 * v2), F1: 255}
                _ = v13
            }
            v14 := GenS1{F0: v2, F1: 65535}
            _ = v11
            _ = v14
        } else {
            var v15 int16 = 0
            _ = v15
        }
        _ = fi7
    }
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    v2 += 0
    fmt.Println(v2, v2)
    if len(v3) > 0 {
        v3[0] = 125
    }
    if len(v3) > 0 {
        v3[0] = 145
    }
    v2 -= (v2 & v2)
    fmt.Println(v4, v2)
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
