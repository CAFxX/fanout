package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int, p1 uint8, p2 int16) int64 {
    if true {
        if false {
            switch (-(p0)) {
            case 42:
                v7 := []int{p0, 840, 0}
                _ = v7
            case 32767:
                p2 -= p2
            }
            switch "hello" {
            case "foo bar":
                var v8 int8 = 42
                _ = v8
            }
            if (false && (((float32(float64((((((float32(((1.5 * (float64(1.5) + 1.5)) + 0.1)) / 3.14159) * ((3.14159 + ((0.0 * (1e10 / 1e10)) / 3.14159)) + 1e10)) - 0.0) / 1e10) + float32((1.5 / 3.14159))))) + 0.1) + 0.1) < 1.5)) {
                for fi9 := 0; fi9 < 14; fi9++ {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        v11 := GenS1{F0: 0, F1: 364}
                        switch fi9 {
                        case 2147483647:
                            if (!false) {
                                v12 := GenS1{F0: -9223372036854775808, F1: 1}
                                var v13 uint64 = 1
                                var v14 bool = false
                                var v15 int = (-(p0))
                                _ = v12
                                _ = v13
                                _ = v14
                                _ = v15
                            } else {
                                var v16 int32 = 2147483647
                                v17 := []int{}
                                v16 += (v16 >> ((807) & 63))
                                _ = v16
                                _ = v17
                            }
                        case 0:
                            p1 -= p1
                        case 32768:
                            v18 := GenS1{F0: 609, F1: 65535}
                            _ = v18
                        default:
                            var v19 uint8 = (p1 + p1)
                            _ = v19
                        }
                        _ = fi10
                        _ = v11
                    }
                    for fi20 := 0; fi20 < 3; fi20++ {
                        if ((0.1 / 3.14159) > 1e10) {
                            p1 += 12
                            v21 := []int{}
                            p2 -= (p2 >> ((18446744073709551615) & 63))
                            _ = v21
                        } else {
                            var v22 uint8 = p1
                            _ = v22
                        }
                        var v23 uint8 = p1
                        p1 += v23
                        _ = fi20
                        _ = v23
                    }
                    v24 := []int{}
                    _ = fi9
                    _ = v24
                }
                fmt.Println(p1, p2, p1)
                var v25 bool = (p1 == 2)
                _ = v25
            }
        }
    } else {
        p0 = p0
        p0 = p0
        p0 -= p0
    }
    p2 -= 603
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 int8 = 127
    var v5 uint16 = 32767
    v6 := []uint32{2, 622, 458}
    v1 = false
    v26 := v3[-9223372036854775808]
    switch 742 {
    case 991:
        var v27 bool = (1.5 == 1.5)
        _ = v27
    case 0:
        v28 := GenS1{F0: -1, F1: 673}
        _ = v28
    default:
        for fi29 := 0; fi29 < 18; fi29++ {
            delete(v2, 2)
            if true {
                v30 := GenS2{F0: (^(v26))}
                _ = len(v6)
                _ = v30
            } else {
                fmt.Println(v4, v26, v5)
                if (false || v1) {
                    if len(v6) > 0 {
                        v6[0] = v26
                    }
                }
            }
            if v1 {
                fmt.Println(v26, fi29, v1)
                for fi31 := 0; fi31 < 15; fi31++ {
                    v6 = append(v6, v26)
                    _ = fi31
                }
                fmt.Println(v26)
            }
            _ = fi29
        }
    }
    for fi32 := 0; fi32 < 9; fi32++ {
        v1 = (int8(v26) == v4)
        v5 += (v5 * v5)
        _ = fi32
    }
    delete(v2, 216)
    fmt.Println(v4)
    var v33 int64 = 0
    v34 := GenS2{F0: v26}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
