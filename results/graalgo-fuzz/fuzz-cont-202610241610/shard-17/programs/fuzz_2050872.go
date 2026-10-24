package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint16 }
func genFn0(p0 uint) string {
    fmt.Println(p0)
    fmt.Println(p0)
    var v7 int32 = 1
    switch "foo bar" {
    case "foo bar":
        if (false || false) {
            v8 := []int{}
            _ = v8
        }
    }
    _ = p0
    _ = v7
    return "\n\t"
}

func main() {
    var v1 bool = (4294967295 <= 185)
    var v2 uint16 = 154
    v3 := make(map[uint32]int)
    var v4 int32 = 0
    v5 := []uint32{0, 42, 1, 917}
    var v6 bool = v1
    v9 := GenS1{F0: 42, F1: 916, F2: 42}
    v6 = (!(!v6))
    var v10 uint8 = 1
    var v11 float32 = -2.25
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = 2
    }
    var v12 int16 = 0
    delete(v3, 0)
    var v13 float32 = (float32((float64(v11) / (((3.14159 * (-2.25 * -2.25)) + (((1.5 - ((1.5 + 0.1) - 1e10)) - 0.0) - -2.25)) + ((float64(0.0) / float64(((1e10 - (v11 - v11)) + (((v11 + (v11 / 3.14159)) - v11) / v11)))) - (0.1 * (-2.25 * (((1e10 + 1.5) * float64(float32((-2.25 / (1.5 * (((0.0 + 0.1) + (1.5 * 3.14159)) + float64((0.1 * (3.14159 - (((1e10 - v11) - v11) - -2.25)))))))))) - ((3.14159 * float64(-2.25)) - 1e10)))))))) - (0.0 - 0.1))
    if (1 <= -62) {
        for fi14 := 0; fi14 < 12; fi14++ {
            v12 -= v12
            v15 := v3[42]
            _ = fi14
            _ = v15
        }
        fmt.Println(v11, v2)
        v6 = v6
    } else {
        v3[237] = 336
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 float32 = 1.5
            v18 := GenS2{F0: (v10 & (v10 ^ (v10 >> ((65535) & 63)))), F1: (v2 == (v2 + v2)), F2: (v2 + 2)}
            var v19 bool = (v2 < v2)
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        var v20 bool = (float64(float32(3.14159)) >= (((1e10 - ((float64(v11) + 3.14159) / float64(((v13 - (((((v11 * v11) * v13) - 1.5) + (1e10 / (1e10 * (v13 - v13)))) + float32(0.0))) * v13)))) - 1.5) / (0.1 / 0.1)))
        _ = v20
    }
    var v21 int16 = (v12 - v12)
    fmt.Println(v4, v13, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
