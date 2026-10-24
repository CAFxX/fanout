package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 uint8 = 128
    v2 := []int{}
    v3 := make(map[string]string)
    var v4 uint = 65535
    for fi5 := 0; fi5 < 15; fi5++ {
        if true {
            v6 := GenS1{F0: v4}
            _ = v6
        }
        switch fi5 {
        case -1:
            var v7 uint8 = (v1 + (v1 * 128))
            _ = v7
        case 9223372036854775807:
            var v8 float32 = float32((((-2.25 + -2.25) + (0.0 + (0.1 * (0.0 + 0.1)))) + 0.1))
            _ = v8
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi5
    }
    fmt.Println(v1)
    v9 := GenS2{F0: 856, F1: (v1 / ((42) | 1))}
    for fi10 := 0; fi10 < 7; fi10++ {
        if true {
            if (!(373 == 42)) {
                var v11 uint8 = (v1 >> ((4294967295) & 63))
                v4 -= v4
                v4 -= (v4 ^ v4)
                _ = v11
            }
        } else {
            v1 -= v1
            v12 := GenS1{F0: 0}
            if (4294967295 > 18446744073709551615) {
                var v13 int8 = 42
                switch fi10 {
                case -715:
                    v4 += v4
                case 1:
                    var v14 int32 = -1
                    _ = v14
                }
                v15 := GenS2{F0: 0, F1: v1}
                _ = v13
                _ = v15
            } else {
                fmt.Println(v4)
                v4 = 4294967295
                fmt.Println(fi10, fi10, fi10)
            }
            _ = v12
        }
        _ = fi10
    }
    v4 += 4294967295
    v4 = (v4 | v4)
    v16 := v3["hello"]
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
