package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0(p0 int32, p1 uint16) int64 {
    v6 := []int{9223372036854775807, 9223372036854775807, 0}
    var v7 int = 2147483647
    var v8 uint32 = 540
    p0 = (p0 / (((p0 / ((-313) | 1))) | 1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 9223372036854775807
}
func genFn1(p0 uint32) uint32 {
    p0 -= (p0 ^ 874)
    var v9 uint = 255
    _ = p0
    _ = v9
    return p0
}

func main() {
    var v1 int16 = 1000
    var v2 uint32 = 402
    var v3 string = ("\n\t" + "a")
    var v4 int8 = -1
    v5 := []uint32{v2}
    if ((0 >= -9223372036854775808) && (2147483647 < -2147483648)) {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int = (fi10 >> ((0) & 63))
            _ = fi10
            _ = v11
        }
    }
    if true {
        var v12 uint16 = 255
        v13 := GenS1{F0: 4294967295}
        _ = v12
        _ = v13
    }
    var v14 uint16 = 456
    v4 = v4
    v1 += (v1 | v1)
    var v15 bool = true
    v16 := GenS2{F0: v14, F1: 853}
    if (!true) {
        for fi17 := 0; fi17 < 17; fi17++ {
            if v15 {
                switch 9223372036854775807 {
                case 42:
                    fmt.Println(v1, v3, v4)
                case 127:
                    var v18 float32 = 3.14159
                    _ = v18
                default:
                    _ = len(v5)
                }
                v1 -= v1
            } else {
                fmt.Println(v2, v3, v2)
                var v19 uint = 42
                switch 6 {
                case 128:
                    for fi20 := 0; fi20 < 6; fi20++ {
                        switch fi20 {
                        case 32767:
                            var v21 uint64 = 42
                            _ = v21
                        default:
                            v15 = false
                        }
                        _ = len(v5)
                        v5 = append(v5, v2)
                        _ = fi20
                    }
                default:
                    v4 += (v4 << ((527) & 63))
                }
                _ = v19
            }
            _ = fi17
        }
        v5 = append(v5, v2)
    } else {
        v22 := GenS2{F0: v14, F1: uint(v14)}
        _ = v22
    }
    var v23 uint64 = 1
    if (255 == uint8(v14)) {
        v4 = v4
        for fi24 := 0; fi24 < 1; fi24++ {
            v23 += v23
            _ = fi24
        }
        v25 := GenS2{F0: (v14 | v14), F1: 65535}
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        switch "hello" {
        case "hello":
            v1 = 994
        default:
            for fi27 := 0; fi27 < 20; fi27++ {
                switch v3 {
                case "foo bar":
                    var v28 uint64 = v23
                    _ = v28
                case "\n\t":
                    var v29 float32 = 1e10
                    _ = v29
                }
                v1 += v1
                _ = fi27
            }
        }
        fmt.Println(v2, v15)
        v30 := GenS2{F0: 1, F1: 916}
        _ = fi26
        _ = v30
    }
    if (1 == -2147483648) {
        v3 = (v3 + v3)
    }
    var v31 float64 = 1.5
    v1 += (v1 ^ (v1 % ((v1) | 1)))
    v4 = (v4 / ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
