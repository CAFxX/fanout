package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int16; F2 uint32 }

func main() {
    v1 := []int{9223372036854775807, 128, 128, -2147483648}
    var v2 int8 = 1
    var v3 uint8 = 255
    var v4 uint16 = uint16(v3)
    v5 := make(map[uint32]bool)
    var v6 int64 = 1
    var v7 uint32 = 4294967295
    v2 = int8(v3)
    _ = len(v1)
    if true {
        var v8 bool = (true || (false || (float64(1.5) >= 0.1)))
        var v9 float32 = 0.0
        var v10 uint = 1
        _ = v10
        _ = v8
        _ = v9
    } else {
        if false {
            switch 0 {
            case 128:
                v11 := GenS1{F0: 32767}
                _ = v11
            case 442:
                fmt.Println(v6, v7)
            case 32767:
                v1 = append(v1, int(v2))
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                _ = len(v1)
                _ = fi12
            }
        }
    }
    v2 += (v2 ^ (v2 ^ (v2 % (((v2 | (v2 << ((18446744073709551615) & 63)))) | 1))))
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS2{F0: (-(v6))}
        _ = fi13
        _ = v14
    }
    v3 = (v3 & v3)
    v4 = 65535
    if false {
        v15 := v5[v7]
        _ = v15
    }
    v7 -= v7
    v16 := GenS2{F0: int64(v4)}
    v4 -= v4
    for fi17 := 0; fi17 < 14; fi17++ {
        switch fi17 {
        case -167:
            v18 := GenS3{F0: 676, F1: -287, F2: (v7 + v7)}
            _ = v18
        case -384:
            var v19 int64 = -144
            _ = v19
        case -2147483648:
            v1 = append(v1, (fi17 | fi17))
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 float64 = ((0.0 + (3.14159 - (1.5 + (0.0 * float64(1e10))))) * (0.0 + ((-2.25 + 0.1) + (((0.1 - float64(0.1)) - 0.0) - 1.5))))
            _ = fi20
            _ = v21
        }
        v2 = (v2 * -128)
        _ = fi17
    }
    v22 := GenS2{F0: v6}
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 string = "\n\t"
        _ = fi23
        _ = v24
    }
    var v25 int = 307
    var v26 uint32 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
