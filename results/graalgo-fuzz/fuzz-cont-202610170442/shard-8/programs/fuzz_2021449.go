package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float32 = ((1e10 - -2.25) + 3.14159)
    var v2 float64 = 0.1
    var v3 string = "\n\t"
    v4 := []uint32{361, 42, 2147483647}
    v3 = v3
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 float32 = ((1e10 - (1e10 - v1)) - 1.5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 += (v2 / ((1.5 * float64(float32(v2))) / v2))
        _ = fi7
    }
    v8 := GenS1{F0: 42, F1: 65, F2: 1000}
    fmt.Println(v1, v3, v1)
    fmt.Println(v3)
    _ = len(v4)
    if (!(749 <= 627)) {
        var v9 uint32 = 65535
        var v10 float32 = (float32(float64(v1)) / (v1 + float32(1.5)))
        _ = v10
        _ = v9
    } else {
        v11 := GenS2{F0: 297}
        _ = v11
    }
    if (!(42 != -2147483648)) {
        switch 32767 {
        case 32768:
            var v12 float64 = (float64(float32(v2)) / -2.25)
            _ = v12
        case 42:
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 string = ("\n\t" + v3)
                var v15 bool = (true || true)
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            for fi16 := 0; fi16 < 8; fi16++ {
                v2 = v2
                v4 = append(v4, 961)
                v2 += 3.14159
                _ = fi16
            }
        }
        v1 = -2.25
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 float32 = 1.5
            _ = fi17
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
