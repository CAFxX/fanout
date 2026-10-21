package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 2
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 bool = false
    var v5 uint16 = 42
    var v6 int16 = -113
    var v7 float64 = float64(1e10)
    if (float32(((-2.25 * (v7 + float64((float32(v7) - 1e10)))) / -2.25)) > float32(3.14159)) {
        var v8 string = "hello"
        if ((2147483647 >= -775) || true) {
            v4 = false
        } else {
            v2[1] = 42
        }
        _ = v8
    } else {
        v1 = 359
        var v9 float64 = v7
        _ = v9
    }
    v1 += (v1 % (((v1 | v1)) | 1))
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int32 = 432
        var v12 uint16 = v5
        v13 := v3[(fi10 >> (((^(v1))) & 63))]
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    if (!v4) {
        switch "a" {
        case "\n\t":
            v14 := GenS1{F0: -893}
            _ = v14
        default:
            var v15 float64 = ((3.14159 * 0.0) * (v7 * 0.1))
            _ = v15
        }
        fmt.Println(v6, v5)
        v6 = v6
    } else {
        v16 := v3[2147483647]
        for fi17 := 0; fi17 < 9; fi17++ {
            switch "a" {
            case "\n\t":
                v1 = (v1 & 1)
            }
            _ = fi17
        }
        _ = v16
    }
    v1 -= (v1 % ((v1) | 1))
    v18 := GenS2{F0: (^(v1))}
    v19 := GenS1{F0: 42}
    v4 = (-1 < 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
