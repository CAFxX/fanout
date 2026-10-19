package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }
type GenS2 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int = 211
    var v2 uint = 4294967295
    v3 := []uint32{757, 730, 2, 236}
    if (!true) {
        v3 = append(v3, 2147483648)
        if ("a" < "hello") {
            v4 := GenS1{F0: 42, F1: 1000, F2: (1000 >= -32768)}
            fmt.Println(v1)
            _ = v4
        }
        v2 -= (^(v2))
    } else {
        if ((v2 << (((v2 + (v2 >> ((v2) & 63)))) & 63)) <= (^(v2))) {
            fmt.Println(v1, v2)
        }
        v1 += (v1 | v1)
    }
    var v5 uint8 = 75
    v3 = append(v3, 2147483648)
    for fi6 := 0; fi6 < 2; fi6++ {
        switch (-(v1)) {
        case 32767:
            if true {
                for fi7 := 0; fi7 < 3; fi7++ {
                    if (2 < 42) {
                        v8 := GenS2{F0: 556, F1: 2, F2: 621}
                        v9 := GenS1{F0: 149, F1: -598, F2: (1e10 <= float32(float64(3.14159)))}
                        _ = v8
                        _ = v9
                    }
                    var v10 uint8 = v5
                    _ = fi7
                    _ = v10
                }
                v1 += fi6
                v11 := GenS1{F0: -661, F1: -1, F2: (!(-1 < -1))}
                _ = v11
            }
        default:
            var v12 uint32 = 2
            _ = v12
        }
        _ = fi6
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS1{F0: -86, F1: -794, F2: true}
        var v15 int16 = 1
        v16 := GenS3{F0: 0}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 uint32 = 0
        var v19 int16 = 1
        _ = len(v3)
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 uint32 = 2147483648
    v5 = (v5 + 1)
    fmt.Println(v5)
    if len(v3) > 0 {
        v3[0] = (^(v20))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v5
}
