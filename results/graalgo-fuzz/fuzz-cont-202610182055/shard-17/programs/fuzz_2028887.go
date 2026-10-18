package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []int{42, 1, -1, -63}
    var v2 int16 = -1
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 127}
    v2 -= -1
    if (-2147483648 <= -2147483648) {
        v2 = (v2 - (^(v2)))
        v5 := GenS1{F0: 2}
        _ = v5
    }
    fmt.Println(v2)
    var v6 float32 = float32(((float64(1.5) / 1.0) / (1.5 * 0.1)))
    for fi7 := 0; fi7 < 12; fi7++ {
        v6 = float32(1.5)
        var v8 int32 = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: v2}
    var v10 int16 = v2
    if ((0 <= 948) && (false && true)) {
        v11 := v3[65535]
        delete(v3, 866)
        if ((2 <= 127) || (774 != 65535)) {
            v10 -= v2
            v10 = (v10 * (v10 + v10))
            var v12 uint8 = 128
            _ = v12
        } else {
            v11 = "foo bar"
            var v13 uint8 = 128
            _ = v13
        }
        _ = v11
    }
    v2 = -31
    switch -1 {
    case 9223372036854775807:
        v10 = (v10 ^ (-(v10)))
    default:
        v2 = (v2 & v10)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
