package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
func genFn0(p0 uint, p1 string) int8 {
    v7 := []int{4294967295, 0, int(p0)}
    p0 -= p0
    v7 = append(v7, -1)
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 float32 = (1.5 / ((((1e10 / 1.0) * 1.5) / (1e10 / (1.5 + 3.14159))) / 1.0))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return -128
}
func genFn1(p0 uint32, p1 bool, p2 string) uint8 {
    var v10 string = "hello"
    p1 = p1
    var v11 float64 = (3.14159 - (1.5 + (float64((3.14159 - 3.14159)) / float64(float32((1e10 * -2.25))))))
    v12 := GenS1{F0: 0, F1: v10}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 227
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    var v4 int16 = -828
    v5 := make(map[int]int)
    var v6 int16 = (v4 >> ((280) & 63))
    var v13 uint64 = 18446744073709551615
    v13 = v13
    v14 := GenS1{F0: 1, F1: "hello"}
    v15 := v2[123]
    switch "foo bar" {
    case "\n\t":
        delete(v5, 941)
    default:
        v4 = -1
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v5[fi16] = fi16
        v17 := v2[v15]
        _ = fi16
        _ = v17
    }
    var v18 int64 = -9223372036854775808
    var v19 uint64 = (^(v13))
    var v20 int8 = 87
    v13 += 1
    switch 32767 {
    case 2147483647:
        switch -524 {
        case 473:
            if len(v3) > 0 {
                v3[0] = 0
            }
        default:
            var v21 int8 = (v20 | (v20 << ((2) & 63)))
            _ = v21
        }
    case -662:
        fmt.Println(v4)
    default:
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    switch -963 {
    case -351:
        var v22 float64 = float64(-2.25)
        _ = v22
    case -1:
        if true {
            var v23 float64 = 0.1
            _ = v23
        }
    case 0:
        if true {
            if true {
                if (!false) {
                    if true {
                        fmt.Println(v18)
                        v1 += v18
                    }
                    v20 -= (v20 << ((18446744073709551615) & 63))
                }
                if (2 > 255) {
                    v24 := GenS1{F0: 834, F1: (("a" + "") + "")}
                    _ = v24
                }
                for fi25 := 0; fi25 < 10; fi25++ {
                    if ((v15 / ((v15) | 1)) >= v15) {
                        for fi26 := 0; fi26 < 18; fi26++ {
                            var v27 uint16 = 625
                            _ = fi26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 4; fi28++ {
                            v18 = -374
                            switch fi28 {
                            case 4294967295:
                                v4 -= -505
                            case -47:
                                v13 -= (v19 << ((2) & 63))
                            default:
                                fmt.Println(v4)
                            }
                            _ = fi28
                        }
                    }
                    _ = fi25
                }
            }
            fmt.Println(v19, v20, v15)
        }
    }
    var v29 string = "hello"
    if false {
        var v30 int16 = v6
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
