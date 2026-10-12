package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 uint) uint64 {
    var v4 int64 = 504
    v4 = v4
    var v5 string = "\n\t"
    v6 := GenS1{F0: -1, F1: -907}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 bool) uint64 {
    var v7 uint8 = 128
    v8 := []int{-1, 32768}
    for fi9 := 0; fi9 < 17; fi9++ {
        v7 += v7
        if p0 {
            fmt.Println(v7, v7)
            v8 = append(v8, 169)
        } else {
            var v10 uint16 = 32767
            _ = v10
        }
        p0 = (42 >= 2147483647)
        _ = fi9
    }
    if p0 {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: -9223372036854775808, F1: 1}
            _ = fi11
            _ = v12
        }
        fmt.Println(v7, p0)
        for fi13 := 0; fi13 < 8; fi13++ {
            switch 127 {
            case -478:
                switch fi13 {
                case 9223372036854775807:
                    if len(v8) > 0 {
                        v8[0] = -408
                    }
                default:
                    var v14 uint = 18446744073709551615
                    _ = v14
                }
            case 4294967295:
                v7 -= v7
            }
            _ = fi13
        }
    } else {
        v7 = v7
        v7 = v7
    }
    var v15 uint32 = uint32(v7)
    _ = p0
    _ = v15
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 string = ("\n\t" + ("\n\t" + "\n\t"))
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    fmt.Println(v1)
    var v16 int32 = 2147483647
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint64 = genFn0(2)
        _ = fi17
        _ = v18
    }
    v1 = (v1 + ("" + v1))
    v16 = (v16 / ((2147483647) | 1))
    if false {
        v1 = v1
    } else {
        v19 := GenS1{F0: int64(v16), F1: 1}
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 float64 = ((-2.25 - (-2.25 - 0.0)) - (3.14159 * 1e10))
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS1{F0: 772, F1: 4}
            var v24 float64 = 1.5
            _ = fi22
            _ = v23
            _ = v24
        }
        fmt.Println(v21, v16)
        _ = fi20
        _ = v21
    }
    fmt.Println(v16)
    var v25 float64 = 1.5
    v25 += 1e10
    v16 -= (^(v16))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
