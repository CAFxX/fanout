package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 bool, p1 uint64, p2 string) uint8 {
    if (-112 > 30) {
        var v7 uint8 = 2
        if p0 {
            fmt.Println(p2, p0, p1)
            var v8 float64 = ((((float64(1.5) - float64(float32(0.1))) / (float64((-2.25 - 0.1)) + (0.0 - -2.25))) + 1e10) - 1e10)
            v7 = v7
            _ = v8
        } else {
            switch 1 {
            case 42:
                v9 := []int{32768, -9223372036854775808}
                _ = v9
            default:
                for fi10 := 0; fi10 < 6; fi10++ {
                    v11 := GenS1{F0: (v7 ^ (v7 ^ v7)), F1: -460, F2: (v7 << ((747) & 63))}
                    v12 := []int{fi10}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        }
        p2 = p2
        _ = v7
    } else {
        v13 := []int{}
        _ = v13
    }
    var v14 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 255
}

func main() {
    var v1 int = -2147483648
    var v2 uint64 = 735
    var v3 uint16 = 2
    v4 := []string{"", "hello", "\n\t", "hello"}
    var v5 int8 = -1
    var v6 bool = ((v5 >= v5) && ((v5 & (v5 * (-(v5)))) > 42))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v15 := GenS1{F0: 42, F1: -692, F2: 73}
    var v16 uint8 = 142
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint64 = (^(v2))
    var v18 uint32 = 2147483647
    var v19 uint32 = v18
    _ = len(v4)
    if v6 {
        switch v1 {
        case 522:
            var v20 uint64 = v17
            _ = v20
        case 9223372036854775807:
            var v21 uint16 = (v3 + uint16(v17))
            _ = v21
        default:
            var v22 uint32 = uint32(v3)
            _ = v22
        }
    }
    var v23 int32 = 315
    if (v1 == v1) {
        var v24 int8 = (^(v5))
        for fi25 := 0; fi25 < 11; fi25++ {
            v26 := GenS2{F0: 1}
            _ = fi25
            _ = v26
        }
        v5 -= v24
        _ = v24
    } else {
        v27 := GenS1{F0: (^(v16)), F1: -933, F2: (^(v16))}
        _ = v27
    }
    v28 := GenS1{F0: v16, F1: v23, F2: v16}
    var v29 int16 = -816
    if true {
        var v30 uint = 255
        _ = v30
    }
    var v31 float64 = (3.14159 - float64(((((((1.5 / 0.1) * float32((float64(1e10) + 0.1))) + (-2.25 + 3.14159)) * (((3.14159 * ((0.0 + ((1e10 + (1e10 * ((0.0 - (3.14159 / 0.1)) / 0.1))) + ((1.5 / float32(0.1)) * (float32(1.5) / (1.5 + -2.25))))) / (3.14159 * 0.1))) * -2.25) / 1.5)) - (-2.25 * ((-2.25 - 1.5) - (0.0 + 3.14159)))) * (((-2.25 * 0.0) - 0.0) * (1.5 * 0.0)))))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
