package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int16; F1 int64; F2 int32 }

func main() {
    var v1 int8 = -128
    var v2 uint32 = 8
    v3 := []int{127}
    var v4 float32 = (0.0 - 1.5)
    v4 -= (v4 - v4)
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int32 = -841
        v1 += (v1 ^ (v1 + (v1 << ((4294967295) & 63))))
        _ = fi5
        _ = v6
    }
    fmt.Println(v1, v2, v2)
    for fi7 := 0; fi7 < 19; fi7++ {
        if (!true) {
            fmt.Println(fi7, v1, v1)
            fmt.Println(v1, v4, v2)
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: (true || (!false)), F1: (^(v1))}
                _ = fi8
                _ = v9
            }
        } else {
            _ = len(v3)
            var v10 uint16 = 42
            _ = v10
        }
        var v11 uint16 = 2
        _ = fi7
        _ = v11
    }
    var v12 int32 = 0
    var v13 uint16 = 0
    v4 = float32(0.0)
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            v1 += (^(v1))
            fmt.Println(v12, v2)
            v13 += v13
            _ = fi15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            if false {
                var v17 int16 = -1
                _ = v17
            } else {
                if (1 >= 0) {
                    v18 := GenS2{F0: 1, F1: 42, F2: (v12 << ((65535) & 63))}
                    v12 -= (-(v12))
                    _ = len(v3)
                    _ = v18
                }
                var v19 int16 = 611
                _ = len(v3)
                _ = v19
            }
            var v20 uint8 = 42
            fmt.Println(v13)
            v4 -= v4
            _ = fi16
            _ = v20
        }
        _ = fi14
    }
    switch "" {
    case "a":
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
