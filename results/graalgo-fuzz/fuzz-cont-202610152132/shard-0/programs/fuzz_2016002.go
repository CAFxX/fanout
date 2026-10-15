package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int8 }
func genFn0(p0 uint64) uint64 {
    var v4 float64 = 1.5
    var v5 uint = 2
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint8) int8 {
    switch 0 {
    case 355:
        v6 := GenS1{F0: 42, F1: (p0 ^ (p0 ^ (p0 * p0))), F2: 127}
        _ = v6
    }
    if ((("\n\t" + "a") < "a") && false) {
        p0 = p0
        v7 := GenS1{F0: genFn0(42), F1: 128, F2: int8(p0)}
        _ = v7
    } else {
        p0 -= p0
        p0 -= 95
        var v8 int16 = -32768
        _ = v8
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint32 = 716
    var v2 uint64 = 4294967295
    v3 := []string{}
    if false {
        v2 = genFn0((v2 >> ((255) & 63)))
        var v9 int = 42
        v9 = v9
        _ = v9
    }
    v2 = v2
    var v10 uint16 = 469
    v3 = append(v3, "foo bar")
    _ = len(v3)
    if (((!false) || false) || (-9223372036854775808 != -9223372036854775808)) {
        var v11 uint = 255
        _ = len(v3)
        _ = v11
    } else {
        v10 -= 136
        if false {
            v2 = (v2 - (^(v2)))
            for fi12 := 0; fi12 < 17; fi12++ {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                var v13 float64 = (float64(3.14159) / 1.0)
                _ = fi12
                _ = v13
            }
        }
        var v14 uint8 = 42
        _ = v14
    }
    if ((1 == 0) || true) {
        v2 = genFn0((v2 & (^(v2))))
        var v15 int8 = 42
        _ = v15
    } else {
        fmt.Println(v1, v10)
    }
    var v16 bool = true
    var v17 float32 = float32(-2.25)
    fmt.Println(v2)
    var v18 float64 = (3.14159 * (0.0 - (0.1 - (3.14159 - float64(float32(((3.14159 - 0.0) + (3.14159 + (((0.1 * -2.25) / 1.0) + -2.25)))))))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
