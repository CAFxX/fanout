package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint8 }
func genFn0(p0 string) int64 {
    v5 := []int{-503, 128}
    var v6 float64 = 3.14159
    if (804 >= 46) {
        v6 = v6
        v6 += (0.1 / 1e10)
        if false {
            p0 = ""
            v7 := GenS1{F0: 2, F1: 739, F2: 0}
            v6 -= -2.25
            _ = v7
        } else {
            v6 = v6
            p0 = "a"
        }
    }
    var v8 uint32 = 595
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return 1
}
func genFn1(p0 bool, p1 int32) uint {
    fmt.Println(p1)
    p1 = 1
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 int32 = 0
    var v2 uint = 0
    var v3 int64 = -9223372036854775808
    v4 := []uint32{}
    v3 += (v3 + v3)
    v9 := GenS1{F0: 65535, F1: 52, F2: (v2 ^ v2)}
    v10 := GenS1{F0: 170, F1: 42, F2: (v2 << (((v2 << (((^(v2))) & 63))) & 63))}
    v3 += (v3 ^ v3)
    var v11 uint16 = 42
    v2 = (v2 - v2)
    switch 32767 {
    case 32767:
        var v12 uint16 = v11
        _ = v12
    case -971:
        var v13 string = ""
        _ = v13
    default:
        v14 := GenS1{F0: v2, F1: 27, F2: v2}
        _ = v14
    }
    if ((!(127 != 0)) || (false && (("\n\t" + "hello") >= "a"))) {
        fmt.Println(v3)
        if true {
            v3 = (^(v3))
            var v15 uint16 = (^(v11))
            if false {
                for fi16 := 0; fi16 < 10; fi16++ {
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                    _ = fi16
                }
            } else {
                v17 := GenS1{F0: genFn1(true, (v1 * (v1 / ((-1) | 1)))), F1: uint32(v15), F2: (v2 & (^(v2)))}
                _ = v17
            }
            _ = v15
        }
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := GenS1{F0: genFn1(((true && (!true)) || ((v2 + 119) <= 795)), v1), F1: 2, F2: v2}
            v21 := GenS2{F0: 4294967295, F1: v11, F2: 127}
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
