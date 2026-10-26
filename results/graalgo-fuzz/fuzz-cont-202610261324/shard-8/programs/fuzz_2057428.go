package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 int32 = 42
    var v2 uint32 = 0
    v3 := make(map[uint32]uint32)
    var v4 uint32 = (v2 / (((v2 << ((887) & 63))) | 1))
    var v5 uint = 0
    v6 := make(map[int]int)
    if ((62 >= -1) || (v4 < v2)) {
        v2 = 1
        v5 = 18446744073709551615
    } else {
        v1 -= (v1 / (((v1 << (((v5 << ((65535) & 63))) & 63))) | 1))
    }
    v7 := GenS1{F0: v1, F1: 160, F2: 42}
    var v8 int32 = v1
    var v9 uint8 = 1
    v10 := []int{0}
    fmt.Println(v5, v2)
    v4 -= v2
    for fi11 := 0; fi11 < 13; fi11++ {
        v10 = append(v10, (fi11 >> (((v5 ^ v5)) & 63)))
        var v12 int = -457
        switch 128 {
        case 1:
            if ((127 >= -1) || (!false)) {
                if (false || false) {
                    v5 += v5
                } else {
                    var v13 uint32 = (v4 / ((0) | 1))
                    var v14 float32 = (((-2.25 - float32(0.0)) - (float32(1e10) + (0.1 / 0.1))) / -2.25)
                    _ = v13
                    _ = v14
                }
                var v15 int64 = 1
                _ = v15
            }
        default:
            v6[-9223372036854775808] = (fi11 / ((1) | 1))
        }
        _ = fi11
        _ = v12
    }
    if (2147483647 <= v1) {
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                switch fi17 {
                case -2147483648:
                    if ((true || (747 != 255)) || false) {
                        var v18 uint16 = 1
                        switch 127 {
                        case 128:
                            switch ("" + ("hello" + "a")) {
                            case "hello":
                                var v19 float32 = ((((((1e10 / 1.0) * 1.5) * ((1.5 * float32(((0.0 / 3.14159) - 1e10))) * (0.1 - 0.0))) * ((0.1 * ((float32(float64(3.14159)) - 1.5) + 1.5)) / (0.0 - 0.1))) / 0.1) * ((1.5 * 0.0) * 1e10))
                                _ = v19
                            }
                        case 80:
                            v10 = append(v10, fi17)
                        default:
                            var v20 bool = false
                            _ = v20
                        }
                        v10 = append(v10, fi17)
                        _ = v18
                    } else {
                        var v21 int = fi16
                        _ = v21
                    }
                case 32768:
                    v22 := GenS1{F0: (v8 % ((-1) | 1)), F1: 144, F2: (v9 ^ v9)}
                    _ = v22
                case 1:
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi16, v8, v1)
                    default:
                        v2 = 49
                    }
                default:
                    v23 := GenS2{F0: 18446744073709551615, F1: (^(v9))}
                    _ = v23
                }
                _ = fi17
            }
            _ = fi16
        }
    } else {
        _ = len(v3)
        if (!(!("foo bar" >= "hello"))) {
            v5 = v5
            v24 := v6[42]
            var v25 float32 = -2.25
            _ = v24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
