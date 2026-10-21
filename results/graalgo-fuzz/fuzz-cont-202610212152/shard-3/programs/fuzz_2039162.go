package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint; F1 int }
func genFn0(p0 uint8) uint16 {
    v6 := GenS1{F0: -285, F1: 943, F2: 1}
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: -1, F1: 562, F2: 4294967295}
        _ = fi7
        _ = v8
    }
    p0 -= (p0 << ((2) & 63))
    var v9 int8 = 42
    _ = p0
    _ = v6
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    var v3 int = 232
    var v4 float32 = -2.25
    v5 := make(map[uint32]string)
    fmt.Println(v4, v3, v2)
    v4 = 3.14159
    v4 = v4
    switch (v3 >> ((1) & 63)) {
    case 9223372036854775807:
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                if true {
                    if false {
                        v12 := GenS2{F0: 415, F1: (fi11 - fi10)}
                        _ = v12
                    } else {
                        var v13 uint64 = 845
                        var v14 float64 = (float64(1e10) + 0.0)
                        _ = v13
                        _ = v14
                    }
                    if (false && (false && (1 < 42))) {
                        fmt.Println(v4)
                        switch fi11 {
                        case 32767:
                            v15 := GenS2{F0: 2, F1: fi11}
                            _ = v15
                        case 737:
                            v3 -= 32767
                        case 9223372036854775807:
                            var v16 int32 = 34
                            v2 -= (v2 / (((^(v2))) | 1))
                            _ = len(v5)
                            _ = v16
                        }
                    } else {
                        var v17 int16 = 0
                        delete(v5, 42)
                        _ = v17
                    }
                }
                _ = fi11
            }
            var v18 uint8 = 1
            _ = fi10
            _ = v18
        }
    default:
        switch (v3 * (v3 + (v3 & 169))) {
        case 127:
            v2 = (v2 / ((v2) | 1))
        case 1:
            v3 += (v3 * 809)
        default:
            v2 -= (v2 << ((4294967295) & 63))
        }
    }
    fmt.Println(v3, v2, v4)
    var v19 int32 = -2147483648
    if (true && true) {
        for fi20 := 0; fi20 < 13; fi20++ {
            switch "\n\t" {
            case "":
                var v21 uint64 = 42
                _ = v21
            case "foo bar":
                delete(v1, 4294967295)
            case "hello":
                v3 = (-(fi20))
            default:
                v22 := v5[2147483648]
                _ = v22
            }
            _ = fi20
        }
        v23 := GenS2{F0: 349, F1: (v3 + (v3 | (-(v3))))}
        _ = v23
    }
    var v24 uint16 = genFn0(1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
