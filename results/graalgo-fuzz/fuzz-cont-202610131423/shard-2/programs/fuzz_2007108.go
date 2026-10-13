package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int, p1 int16, p2 int16) bool {
    var v4 int = (-(p0))
    var v5 uint16 = 42
    p0 -= (p0 / ((995) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (2 == 255)
}

func main() {
    var v1 float32 = 1.5
    var v2 int32 = -507
    var v3 float64 = (1e10 - 0.0)
    var v6 float64 = (v3 / (v3 * v3))
    v2 = v2
    v7 := GenS1{F0: -32768, F1: 100, F2: v2}
    v8 := []int{-521, 9223372036854775807, -1, 128}
    v8 = append(v8, -966)
    v2 = v2
    var v9 string = ""
    var v10 uint32 = 65535
    switch v9 {
    case "\n\t":
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 12; fi12++ {
                fmt.Println(fi12, fi12)
                v2 += -268
                _ = fi12
            }
            if (!(int8(fi11) < -82)) {
                if ((v10 > v10) || false) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v13 string = ((v9 + v9) + v9)
                    _ = v13
                }
            }
            fmt.Println(v10, v2)
            _ = fi11
        }
    case "hello":
        var v14 int = 32768
        _ = v14
    default:
        v8 = append(v8, 574)
    }
    v15 := GenS1{F0: -32768, F1: 58, F2: (v2 >> ((763) & 63))}
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 uint16 = 42
        if false {
            v18 := GenS2{F0: -683}
            var v19 int = fi16
            if len(v8) > 0 {
                v8[0] = 345
            }
            _ = v18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    v2 = v2
    v20 := GenS2{F0: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
