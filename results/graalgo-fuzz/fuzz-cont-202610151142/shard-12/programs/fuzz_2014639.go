package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64; F1 bool; F2 string }
func genFn0(p0 int8, p1 string, p2 uint) uint32 {
    switch 142 {
    case 194:
        p1 = "\n\t"
    case 230:
        p1 = p1
    case 32767:
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: 1000, F1: -2147483648}
            _ = fi6
            _ = v7
        }
    default:
        fmt.Println(p2, p2)
    }
    if (1 <= -32768) {
        v8 := GenS1{F0: 32767, F1: 0}
        _ = v8
    }
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    return 357
}

func main() {
    var v1 uint64 = 0
    v2 := []int{}
    v3 := []uint32{1, 1}
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            _ = len(v5)
            _ = fi10
        }
        v1 -= v1
        var v11 uint8 = 0
        _ = fi9
        _ = v11
    }
    var v12 int8 = 127
    v5[2] = 2147483647
    for fi13 := 0; fi13 < 7; fi13++ {
        if true {
            var v14 float64 = 1.5
            if false {
                if true {
                    var v15 int = fi13
                    fmt.Println(v12)
                    _ = v15
                } else {
                    if (!true) {
                        var v16 int16 = 0
                        var v17 uint = 2
                        v16 = int16(v1)
                        _ = v16
                        _ = v17
                    }
                }
                fmt.Println(v14, v14)
            }
            _ = v14
        } else {
            v5[1] = fi13
        }
        if false {
            v2 = append(v2, 32767)
            var v18 uint8 = 42
            _ = v18
        } else {
            v1 -= v1
        }
        _ = fi13
    }
    switch 42 {
    case 127:
        _ = len(v2)
    default:
        var v19 uint32 = genFn0((-(v12)), "", 0)
        _ = v19
    }
    var v20 int64 = -9223372036854775808
    _ = len(v3)
    v21 := GenS1{F0: -595, F1: -959}
    v20 += (v20 | v20)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v22 bool = (4294967295 != 2147483647)
    var v23 uint8 = 1
    var v24 float64 = (-2.25 + (0.0 + 0.0))
    switch -933 {
    case 0:
        fmt.Println(v23)
    case -799:
        for fi25 := 0; fi25 < 4; fi25++ {
            v2 = append(v2, fi25)
            v26 := GenS1{F0: 1, F1: 2147483647}
            _ = fi25
            _ = v26
        }
    case 2147483647:
        var v27 uint64 = 0
        _ = v27
    }
    v28 := GenS1{F0: -32768, F1: 1}
    fmt.Println(v1, v1)
    if len(v2) > 0 {
        v2[0] = -841
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        fmt.Println(v20)
        v30 := GenS2{F0: v20, F1: v22, F2: ""}
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
