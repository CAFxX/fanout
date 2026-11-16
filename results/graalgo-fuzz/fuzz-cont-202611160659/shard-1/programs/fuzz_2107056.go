package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint = 0
    var v2 uint8 = 42
    var v3 float32 = float32(-2.25)
    var v4 float64 = ((-2.25 / 3.14159) + float64(0.1))
    var v5 uint = 4294967295
    v6 := GenS1{F0: -1, F1: 9223372036854775807}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int16 = -923
        v4 = float64((((3.14159 / v3) * v3) / 3.14159))
        _ = fi7
        _ = v8
    }
    if ((!(((((0.0 * float64(0.0)) * v4) / 0.1) - v4) == v4)) || (127 <= 33)) {
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(fi9, v2, v2)
            v10 := GenS1{F0: -1, F1: 42}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: -1, F1: 0}
        _ = v11
    }
    v12 := GenS1{F0: -367, F1: -1}
    var v13 uint8 = (v2 + v2)
    v3 = (v3 + ((v3 + v3) * v3))
    v3 -= 0.1
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int = fi14
        if (1 <= 18446744073709551615) {
            v2 += 2
            v16 := []int{(v15 / ((v15) | 1)), (fi14 >> ((v5) & 63))}
            for fi17 := 0; fi17 < 17; fi17++ {
                v4 += v4
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi17
            }
            _ = v16
        }
        v18 := GenS2{F0: -545, F1: 2147483648}
        _ = fi14
        _ = v15
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
