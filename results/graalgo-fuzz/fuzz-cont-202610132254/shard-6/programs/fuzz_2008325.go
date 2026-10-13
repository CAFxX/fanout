package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint32 }

func main() {
    var v1 float32 = float32(float64((-2.25 + ((1e10 * float32((float64(0.1) / -2.25))) * ((((3.14159 * (0.0 + (3.14159 * 3.14159))) * 1.5) + 1.5) + 1e10)))))
    v2 := []uint32{2147483648, 664}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    var v5 int8 = 42
    var v6 uint8 = 128
    v7 := GenS1{F0: -9223372036854775808, F1: (v5 % (((v5 << ((516) & 63))) | 1))}
    switch -2147483648 {
    case 42:
        var v8 int32 = 42
        _ = v8
    case 149:
        switch -443 {
        case 32767:
            var v9 int32 = 42
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = fi10
    }
    var v11 int64 = -9223372036854775808
    v12 := GenS1{F0: 1, F1: (v5 / ((v5) | 1))}
    var v13 int32 = -2147483648
    if (128 < v6) {
        fmt.Println(v5, v1, v6)
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 string = "foo bar"
            _ = fi14
            _ = v15
        }
        v6 += v6
    }
    v16 := GenS2{F0: 42, F1: 1, F2: 669}
    var v17 int16 = 0
    v11 -= (v11 / ((v11) | 1))
    if false {
        for fi18 := 0; fi18 < 4; fi18++ {
            v1 -= v1
            if (!(2 == 918)) {
                v5 -= (v5 | (v5 % (((-(v5))) | 1)))
            } else {
                fmt.Println(v5)
            }
            if false {
                var v19 float32 = (v1 / (v1 - (v1 / 0.1)))
                v19 -= (((0.0 * float32(float64(((0.0 * ((1.5 / v1) / 1.5)) / 1.0)))) - v19) * v1)
                _ = v19
            }
            _ = fi18
        }
        var v20 uint16 = 1
        var v21 int8 = (-(v5))
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
