package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 127
    v3 := make(map[uint32]bool)
    v2 -= (v2 + (v2 * (v2 ^ (v2 << ((0) & 63)))))
    var v4 int8 = (v2 / ((v2) | 1))
    if (569 != 745) {
        v5 := v1[1]
        var v6 int32 = 1
        var v7 int64 = -1
        _ = v5
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: 2, F1: 255}
            _ = fi8
            _ = v9
        }
        if true {
            var v10 float64 = 0.1
            _ = v10
        } else {
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint = 523
                var v13 uint = (v12 ^ v12)
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                if false {
                    v15 := GenS1{F0: 127, F1: 8}
                    var v16 uint16 = 32767
                    var v17 string = "\n\t"
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = fi14
            }
        }
    }
    var v18 uint32 = 2147483647
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            fmt.Println(fi19)
            v2 = int8(v18)
            _ = fi20
        }
        fmt.Println(fi19, v18)
        var v21 uint16 = 1
        _ = fi19
        _ = v21
    }
    delete(v1, (v18 << ((18446744073709551615) & 63)))
    switch "" {
    case "hello":
        var v22 float64 = ((1e10 + ((0.0 * (3.14159 / -2.25)) / (1e10 + float64(1e10)))) / float64((float32(1.5) / (float32((0.0 + ((3.14159 * (1.5 / ((0.1 + 3.14159) * 3.14159))) * 1e10))) * ((1.5 * ((float32(float64(1.5)) / 1e10) - 1.5)) * -2.25)))))
        _ = v22
    case "\n\t":
        for fi23 := 0; fi23 < 20; fi23++ {
            v2 = -1
            if (-1 != (v4 % ((v4) | 1))) {
                var v24 int32 = 2147483647
                fmt.Println(v18, v24)
                _ = v24
            }
            var v25 int64 = -516
            _ = fi23
            _ = v25
        }
    }
    v2 = -1
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS1{F0: 255, F1: 127}
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := GenS2{F0: (v18 | v18), F1: (^(v18)), F2: 42}
            _ = fi28
            _ = v29
        }
        switch fi26 {
        case 586:
            var v30 int32 = -2147483648
            _ = v30
        case 282:
            var v31 uint32 = (v18 & v18)
            _ = v31
        case -879:
            delete(v1, (v18 ^ (v18 * v18)))
        default:
            v2 += v2
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
