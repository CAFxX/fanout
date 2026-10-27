package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int32; F2 int8 }

func main() {
    v1 := []string{}
    var v2 float32 = ((float32(-2.25) - 0.0) * (-2.25 * float32(((1e10 * 1e10) / (float64(-2.25) - ((float64(-2.25) + -2.25) - 0.0))))))
    var v3 uint64 = 634
    var v4 int8 = 100
    var v5 int16 = 241
    v4 -= v4
    v3 = v3
    var v6 int16 = -680
    switch 1 {
    case 4294967295:
        v7 := GenS1{F0: v4}
        _ = v7
    case 32768:
        var v8 int16 = (v5 / (((v5 | 400)) | 1))
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            v1 = append(v1, ("hello" + "a"))
            _ = fi9
        }
        var v10 bool = (-271 < 328)
        if (v6 <= v5) {
            if (!(-2147483648 != 0)) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 int8 = -1
                    v13 := GenS2{F0: (v4 >> ((2) & 63)), F1: 42, F2: 17}
                    v1 = append(v1, "\n\t")
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                fmt.Println(v4, v2, v5)
            }
            var v14 int8 = 100
            _ = v14
        }
        _ = v10
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS2{F0: v4, F1: 318, F2: (-(v4))}
        _ = fi15
        _ = v16
    }
    var v17 bool = (65535 == 0)
    v4 -= -1
    _ = len(v1)
    var v18 uint32 = 29
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
