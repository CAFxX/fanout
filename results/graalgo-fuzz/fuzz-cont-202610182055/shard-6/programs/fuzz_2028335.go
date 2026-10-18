package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 int16 }
type GenS4 struct { F0 int32; F1 bool; F2 int8 }
func genFn0() int {
    var v5 uint16 = 32767
    var v6 float32 = (1.5 + float32(0.1))
    v6 = (3.14159 + ((0.1 + (v6 + ((v6 - v6) * -2.25))) - float32(1e10)))
    for fi7 := 0; fi7 < 2; fi7++ {
        v5 -= (^(v5))
        switch fi7 {
        case 669:
            var v8 uint = 179
            _ = v8
        case 479:
            v6 += (v6 - 0.0)
        case -9223372036854775808:
            v9 := GenS1{F0: -1, F1: (v5 ^ v5)}
            _ = v9
        }
        var v10 int16 = -781
        _ = fi7
        _ = v10
    }
    _ = v5
    _ = v6
    return 68
}
func genFn1() uint8 {
    v11 := GenS2{F0: false, F1: 987, F2: 42}
    var v12 uint16 = 375
    var v13 int = 2147483647
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int8 = 127
    var v2 float32 = 1e10
    v3 := []uint32{1}
    var v4 int64 = 42
    if (2147483647 != 0) {
        _ = len(v3)
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(v4, fi14, v1)
            v2 = (v2 * (((((v2 * v2) / -2.25) / 1e10) - 3.14159) - ((v2 - v2) * v2)))
            v15 := GenS3{F0: (v1 > -1), F1: 127, F2: 865}
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS3{F0: false, F1: v1, F2: 0}
    var v17 float32 = ((3.14159 * (v2 + v2)) - -2.25)
    if (255 != 65535) {
        v1 = -128
    } else {
        if len(v3) > 0 {
            v3[0] = 629
        }
        fmt.Println(v2, v17)
    }
    var v18 uint16 = 1
    var v19 int64 = v4
    v20 := GenS3{F0: false, F1: (^(v1)), F2: 638}
    var v21 uint32 = 53
    fmt.Println(v17, v2, v4)
    switch 1 {
    case -299:
        fmt.Println(v2, v4)
    }
    v22 := GenS4{F0: 42, F1: (-9223372036854775808 < -258), F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
