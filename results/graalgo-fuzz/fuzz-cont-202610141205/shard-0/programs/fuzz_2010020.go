package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0() uint64 {
    v7 := []int{}
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    _ = v7
    return 809
}
func genFn1(p0 int32, p1 int64, p2 int16) int {
    v8 := GenS1{F0: 0, F1: (p0 >> ((113) & 63))}
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint = 42
        var v11 uint = v10
        if true {
            var v12 bool = true
            _ = v12
        } else {
            v13 := []int{(fi9 / ((-9223372036854775808) | 1)), (^(fi9)), (fi9 / ((fi9) | 1))}
            var v14 uint64 = 4294967295
            var v15 int = fi9
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 127
}

func main() {
    var v1 int16 = -1
    v2 := []string{}
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    var v6 string = "a"
    delete(v3, 4294967295)
    if (100 > 127) {
        for fi16 := 0; fi16 < 18; fi16++ {
            v1 += 871
            switch fi16 {
            case -2147483648:
                v2 = append(v2, ("" + (v6 + v6)))
            case 721:
                if true {
                    var v17 int16 = (-(v1))
                    _ = v17
                }
            default:
                for fi18 := 0; fi18 < 8; fi18++ {
                    v5[fi16] = 65535
                    var v19 bool = false
                    var v20 int64 = 0
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            _ = fi16
        }
    }
    if false {
        v21 := GenS1{F0: 993, F1: 2147483647}
        v6 = (((v6 + v6) + v6) + "")
        v1 += (v1 + (v1 + (v1 - 1)))
        _ = v21
    } else {
        switch v6 {
        case "\n\t":
            for fi22 := 0; fi22 < 12; fi22++ {
                fmt.Println(v6, v1, fi22)
                v1 = (-(v1))
                _ = fi22
            }
        case "":
            var v23 int32 = 2147483647
            _ = v23
        }
        var v24 uint8 = 124
        _ = v24
    }
    var v25 int8 = -128
    fmt.Println(v6, v6)
    v25 += (v25 / (((v25 % ((v25) | 1))) | 1))
    if false {
        fmt.Println(v6)
    } else {
        var v26 uint32 = 2147483647
        _ = v26
    }
    if (869 > 1) {
        delete(v5, 1)
    } else {
        if (((-2147483648 == -2147483648) && true) && false) {
            var v27 uint64 = genFn0()
            _ = v27
        } else {
            v4[410] = (0.0 <= (3.14159 / float64(-2.25)))
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 uint64 = 647
                _ = fi28
                _ = v29
            }
            v30 := GenS2{F0: 562, F1: 255}
            _ = v30
        }
    }
    v31 := GenS2{F0: genFn1(-2147483648, 1, v1), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
