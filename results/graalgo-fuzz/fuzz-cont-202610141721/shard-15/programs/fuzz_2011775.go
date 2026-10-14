package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 int64, p2 string) uint {
    var v6 int16 = 32767
    var v7 uint8 = 255
    for fi8 := 0; fi8 < 2; fi8++ {
        switch fi8 {
        case 218:
            var v9 uint16 = 65535
            _ = v9
        default:
            p1 += p1
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 835
}

func main() {
    var v1 int8 = -53
    var v2 int = 557
    v3 := []uint32{0, 42}
    var v4 float64 = ((0.0 + float64(float32(float64(((0.1 * 0.0) / -2.25))))) * (0.0 + (0.1 / ((((1.5 + -2.25) + (1e10 * 0.1)) - (3.14159 + (-2.25 / 0.1))) / (0.0 - -2.25)))))
    var v5 bool = false
    v5 = v5
    switch "" {
    case "a":
        var v10 int32 = -438
        _ = v10
    default:
        var v11 uint8 = 128
        _ = v11
    }
    switch v2 {
    case 9223372036854775807:
        if (v1 == 70) {
            fmt.Println(v4)
            var v12 uint = 4294967295
            _ = v12
        }
    case -1:
        var v13 float32 = 0.0
        _ = v13
    default:
        v4 = ((float64(3.14159) + ((1e10 * 1e10) - v4)) * (v4 + (1e10 + 0.0)))
    }
    v4 -= v4
    v14 := GenS1{F0: (v5 && (2 > 42)), F1: 577}
    var v15 bool = true
    v16 := GenS2{F0: 42}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, 704)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
