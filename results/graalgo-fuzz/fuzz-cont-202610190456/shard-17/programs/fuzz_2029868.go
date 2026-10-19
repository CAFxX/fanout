package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
func genFn0(p0 string) uint {
    var v4 uint16 = 691
    v5 := GenS1{F0: -17, F1: 1, F2: ("" <= p0)}
    _ = p0
    _ = v4
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint64, p1 int) uint64 {
    v6 := []int{p1, p1, (p1 >> ((0) & 63))}
    v7 := GenS1{F0: 12, F1: 255, F2: (3.14159 >= 0.1)}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 uint = 852
    var v2 int32 = 792
    v3 := []int{-96}
    _ = len(v3)
    v1 = (v1 % ((v1) | 1))
    for fi8 := 0; fi8 < 12; fi8++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v1, v1)
        case "foo bar":
            switch ("foo bar" + "a") {
            case "foo bar":
                var v9 float32 = -2.25
                _ = v9
            case "a":
                for fi10 := 0; fi10 < 6; fi10++ {
                    v2 = (v2 | -2147483648)
                    v1 -= v1
                    _ = fi10
                }
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v11 int8 = 100
        _ = fi8
        _ = v11
    }
    var v12 uint64 = 97
    v1 = (v1 | 18446744073709551615)
    var v13 float64 = 1.5
    v3 = append(v3, 32767)
    if (1.5 > v13) {
        var v14 float32 = (0.0 - 3.14159)
        v15 := GenS1{F0: 127, F1: uint(v12), F2: (!false)}
        v1 = 0
        _ = v14
        _ = v15
    } else {
        var v16 float64 = (v13 / float64((float32((v13 - float64(((1.5 - -2.25) + (0.1 + -2.25))))) * (1.5 + ((3.14159 / (((1e10 + ((-2.25 - 3.14159) / 1e10)) / 1.5) + (-2.25 / 3.14159))) - (1.5 + ((1e10 * 0.1) - (float32(0.1) * 3.14159))))))))
        var v17 float32 = 3.14159
        _ = v16
        _ = v17
    }
    if (1 >= uint16(v12)) {
        if true {
            for fi18 := 0; fi18 < 6; fi18++ {
                if (!true) {
                    var v19 int8 = -37
                    fmt.Println(fi18, v1, v2)
                    _ = v19
                }
                _ = fi18
            }
        }
    } else {
        fmt.Println(v2)
    }
    v20 := GenS2{F0: 42, F1: 183, F2: ((-128 != 42) || true)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
}
