package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := []string{"foo bar", "", "foo bar"}
    v2 := make(map[string]string)
    v3 := []int{}
    var v4 int64 = -9223372036854775808
    var v5 uint32 = 4294967295
    v6 := []uint32{(v5 << ((2) & 63)), v5, v5}
    v5 = (v5 * (v5 - (v5 & (v5 % (((v5 << ((255) & 63))) | 1)))))
    v7 := GenS1{F0: -32768}
    v8 := GenS2{F0: 254, F1: 191}
    delete(v2, "")
    if len(v6) > 0 {
        v6[0] = (v5 & v5)
    }
    if (2 == 255) {
        var v9 int8 = 127
        var v10 bool = true
        _ = v10
        _ = v9
    }
    _ = len(v1)
    v11 := v2[(("\n\t" + "a") + (("" + "hello") + "a"))]
    for fi12 := 0; fi12 < 9; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(fi12)
            v4 += (v4 - v4)
            _ = fi13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint8 = uint8(v5)
            fmt.Println(fi12, v11)
            switch (^(fi12)) {
            case 4294967295:
                v5 = (^(v5))
            case -422:
                v16 := GenS2{F0: v15, F1: 230}
                _ = v16
            default:
                var v17 float32 = (0.0 / (3.14159 / float32(float64(float32(1e10)))))
                _ = v17
            }
            _ = fi14
            _ = v15
        }
        var v18 uint8 = 1
        _ = fi12
        _ = v18
    }
    v5 = (v5 | v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
