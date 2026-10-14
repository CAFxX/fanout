package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    v1 := []int{32768, -1}
    var v2 int = 4294967295
    var v3 int16 = 260
    v4 := []int{v2}
    v3 = -1
    var v5 int8 = 100
    v6 := GenS1{F0: 2, F1: 2}
    v5 += -128
    var v7 int8 = (^(v5))
    v8 := GenS2{F0: 2, F1: 1, F2: 65535}
    v7 = (^(v5))
    if len(v4) > 0 {
        v4[0] = v2
    }
    fmt.Println(v5, v5)
    var v9 string = "foo bar"
    v5 += 100
    var v10 uint16 = 32767
    v5 = 1
    v10 = v10
    switch "hello" {
    case "hello":
        v10 -= (v10 & (v10 | 552))
    case "\n\t":
        switch (v2 << ((454) & 63)) {
        case 32767:
            fmt.Println(v2)
        case -9223372036854775808:
            var v11 int8 = v7
            _ = v11
        case 9223372036854775807:
            for fi12 := 0; fi12 < 5; fi12++ {
                if (true && (1 > v5)) {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 bool = (v7 >= v7)
                        _ = len(v4)
                        v2 += fi13
                        var v15 uint64 = 42
                        var v16 float32 = float32(0.0)
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                } else {
                    v1 = append(v1, 9223372036854775807)
                    var v17 uint32 = 659
                    var v18 float32 = float32(((((1e10 + 0.0) * 0.0) - (float64(((0.0 - (float32((((3.14159 - ((-2.25 + 0.1) / 1e10)) - (((float64(0.0) - (((3.14159 - 0.1) + 1e10) / 3.14159)) - (1.5 - (1.5 / 1.5))) + (1.5 - 1e10))) * (float64((1.5 + -2.25)) * -2.25))) * 0.1)) * -2.25)) - 0.1)) - ((1e10 * 3.14159) / (-2.25 + (0.1 - 3.14159)))))
                    _ = v17
                    _ = v18
                }
                _ = fi12
            }
        }
    default:
        if (!false) {
            var v19 uint32 = uint32(v10)
            if len(v4) > 0 {
                v4[0] = (v2 * (v2 + -53))
            }
            _ = v19
        }
    }
    var v20 uint32 = 42
    v21 := GenS1{F0: 0, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
