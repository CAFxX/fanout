package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int32 = -486
    var v2 uint32 = 65535
    v3 := make(map[uint32]int)
    var v4 int = -508
    var v5 int16 = 32767
    var v6 bool = false
    v1 = -2147483648
    if (!true) {
        v3[648] = v4
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        if (false && true) {
            switch fi7 {
            case 0:
                var v8 uint8 = 216
                _ = v8
            case 159:
                fmt.Println(v4)
            }
        }
        v4 += (fi7 % ((32767) | 1))
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint32 = v2
            _ = fi9
            _ = v10
        }
        _ = fi7
    }
    fmt.Println(v2, v5)
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: v2, F1: -1}
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint = 18446744073709551615
            var v15 int8 = 127
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            switch (fi11 & (fi11 << ((390) & 63))) {
            case 32768:
                v17 := v3[(v2 % ((v2) | 1))]
                _ = v17
            case -1:
                var v18 int = 99
                _ = v18
            case 4294967295:
                v1 -= 0
            }
            var v19 float64 = ((((1.5 + ((3.14159 - (0.1 * (0.0 * 3.14159))) * -2.25)) + (0.0 * 0.0)) - 3.14159) / 3.14159)
            _ = fi16
            _ = v19
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v5)
    var v20 int8 = 1
    v21 := GenS2{F0: (^(v5))}
    var v22 int32 = (v1 / (((v1 >> ((255) & 63))) | 1))
    for fi23 := 0; fi23 < 5; fi23++ {
        for fi24 := 0; fi24 < 15; fi24++ {
            v25 := v3[(^(v2))]
            v6 = (37 > 0)
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v26 := GenS2{F0: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
