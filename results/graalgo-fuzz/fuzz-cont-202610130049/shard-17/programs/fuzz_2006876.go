package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }

func main() {
    var v1 float32 = 0.1
    var v2 bool = (!(true || true))
    v3 := make(map[uint32]uint32)
    var v4 int32 = -2147483648
    var v5 uint16 = 606
    fmt.Println(v2)
    fmt.Println(v5)
    fmt.Println(v5, v4, v1)
    switch "hello" {
    case "":
        if ((!((0.0 * ((3.14159 - (float32(1e10) / (-2.25 - ((v1 / ((float32(0.1) * (0.0 / v1)) * 1e10)) * (v1 - float32((float64(0.1) + 3.14159))))))) - (v1 - float32((((1.5 / 1e10) * (1e10 * 1.5)) + float64(((v1 + (1.5 - (v1 * float32(0.0)))) + v1))))))) != v1)) || true) {
            var v6 bool = false
            _ = v6
        } else {
            fmt.Println(v4, v5)
            fmt.Println(v1, v2)
        }
    default:
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint8 = 255
            var v9 float64 = 3.14159
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 float32 = 1.5
    fmt.Println(v2)
    _ = len(v3)
    var v11 uint = 0
    v5 = (v5 & (v5 << ((uint(v5)) & 63)))
    v11 = (v11 * v11)
    fmt.Println(v5, v4, v11)
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            fmt.Println(v2, v2, v11)
            v5 = v5
            v10 = v10
            var v14 bool = ((!v2) && true)
            v4 = (v4 % (((v4 + (v4 - v4))) | 1))
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    switch int(v4) {
    case 163:
        var v15 float32 = ((float32(float64(float32(3.14159))) + ((1.5 + v10) - 1.5)) * float32(1.5))
        _ = v15
    }
    v16 := GenS1{F0: 128, F1: 9223372036854775807, F2: v5}
    switch 32768 {
    case -9223372036854775808:
        var v17 uint = (v11 - v11)
        _ = v17
    case -490:
        var v18 string = "foo bar"
        _ = v18
    case -2147483648:
        v11 += (^(v11))
    }
    if (!v2) {
        v4 += (v4 ^ v4)
        v3[215] = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
