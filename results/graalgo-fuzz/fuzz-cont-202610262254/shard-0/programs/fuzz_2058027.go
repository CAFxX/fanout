package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int64 = 588
    var v2 uint = 255
    v3 := make(map[uint32]uint32)
    v2 -= v2
    fmt.Println(v1, v1)
    var v4 float64 = 0.0
    var v5 float64 = 1e10
    var v6 int8 = 38
    if true {
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 128, F1: (v6 | (v6 / ((-97) | 1))), F2: 253}
            if (42 >= (fi7 >> (((v2 * v2)) & 63))) {
                v9 := GenS1{F0: 179, F1: v6, F2: 197}
                v5 -= v5
                _ = v9
            } else {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := v3[269]
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi7
            _ = v8
        }
    } else {
        v12 := GenS1{F0: 2, F1: v6, F2: 140}
        var v13 int64 = (v1 * (v1 + (v1 - (v1 + 0))))
        switch "\n\t" {
        case "":
            fmt.Println(v1)
        case "foo bar":
            switch "" {
            case "a":
                v2 += 606
            case "\n\t":
                v14 := GenS1{F0: 0, F1: (v6 / (((v6 | 127)) | 1)), F2: 190}
                _ = v14
            case "hello":
                v4 -= (3.14159 - float64(((float32(0.1) + ((-2.25 * (((-2.25 - 1.5) - (float32(v5) - (3.14159 * 1e10))) + (1.5 + (3.14159 + (1e10 / 1.5))))) - 3.14159)) + 1e10)))
            default:
                v15 := GenS2{F0: -1}
                _ = v15
            }
        case "hello":
            var v16 uint = (v2 ^ 65535)
            _ = v16
        }
        _ = v12
        _ = v13
    }
    v1 = (^(v1))
    v6 += (v6 << ((v2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
