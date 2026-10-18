package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint32, p1 string, p2 uint8) string {
    p2 = 128
    if false {
        if (69 >= 9223372036854775807) {
            v6 := GenS1{F0: 689, F1: 42, F2: (p2 & 255)}
            v7 := []int{}
            _ = v6
            _ = v7
        } else {
            if (!false) {
                for fi8 := 0; fi8 < 6; fi8++ {
                    v9 := GenS1{F0: -1, F1: int8(p2), F2: (p2 & p2)}
                    _ = fi8
                    _ = v9
                }
            }
            v10 := []int{127, 9223372036854775807, 42, 62}
            var v11 uint8 = (^(p2))
            _ = v10
            _ = v11
        }
        p0 -= p0
    }
    switch p1 {
    case "":
        if ((0.1 <= ((-2.25 + -2.25) - 1e10)) || true) {
            var v12 string = "\n\t"
            p0 -= 2147483648
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 float64 = 0.1
    v4 := make(map[uint32]int)
    var v5 bool = (2147483647 != 32768)
    if v5 {
        var v13 uint32 = 719
        _ = v13
    }
    v14 := GenS1{F0: -32768, F1: 127, F2: 1}
    v15 := GenS2{F0: 2, F1: -1, F2: 2}
    var v16 int8 = 127
    v17 := GenS1{F0: 42, F1: v16, F2: 128}
    var v18 uint32 = 2147483647
    for fi19 := 0; fi19 < 12; fi19++ {
        v20 := GenS1{F0: -1, F1: v16, F2: 50}
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: v16, F1: 42, F2: 25}
    if (!true) {
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := []int{(fi22 | (fi22 | 42)), (fi22 ^ fi22), (fi22 - (fi22 / (((fi22 * (fi22 << ((737) & 63)))) | 1))), (fi22 & (-(fi22)))}
            var v24 int16 = -171
            _ = fi22
            _ = v23
            _ = v24
        }
        delete(v1, ("a" + "a"))
    }
    v25 := GenS4{F0: v16}
    v26 := []int{128, -2147483648, 1, 9223372036854775807}
    v3 = v3
    switch "a" {
    case "hello":
        if false {
            if v5 {
                if false {
                    var v27 uint32 = 396
                    _ = v27
                } else {
                    switch -396 {
                    case -1:
                        v3 -= -2.25
                    }
                }
                switch -104 {
                case 610:
                    v5 = false
                default:
                    if v5 {
                        fmt.Println(v16, v3, v16)
                        var v28 int64 = -9223372036854775808
                        fmt.Println(v3)
                        _ = v28
                    }
                }
            } else {
                switch "foo bar" {
                case "\n\t":
                    var v29 uint32 = (v18 * 2)
                    _ = v29
                case "hello":
                    for fi30 := 0; fi30 < 7; fi30++ {
                        for fi31 := 0; fi31 < 12; fi31++ {
                            v18 = 318
                            _ = fi31
                        }
                        var v32 uint8 = 127
                        _ = fi30
                        _ = v32
                    }
                case "":
                    v5 = (-2.25 <= 0.0)
                default:
                    switch "" {
                    case "foo bar":
                        if v5 {
                            v5 = (255 <= 42)
                            var v33 uint32 = 42
                            _ = v33
                        } else {
                            v34 := GenS5{F0: 324, F1: 0}
                            _ = v34
                        }
                    case "\n\t":
                        var v35 uint32 = (v18 * (v18 / ((v18) | 1)))
                        _ = v35
                    case "hello":
                        v16 += (v16 | 1)
                    default:
                        var v36 int16 = -958
                        _ = v36
                    }
                }
                v5 = v5
            }
            v3 += float64(1e10)
        }
    case "":
        var v37 uint8 = 2
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        _ = len(v26)
        for fi39 := 0; fi39 < 14; fi39++ {
            v26 = append(v26, (fi38 ^ (fi38 - fi38)))
            _ = fi39
        }
        _ = fi38
    }
    fmt.Println(v18, v3, v5)
    for fi40 := 0; fi40 < 2; fi40++ {
        var v41 uint32 = v18
        _ = fi40
        _ = v41
    }
    if (42 == (v16 << ((946) & 63))) {
        var v42 int32 = -1
        fmt.Println(v18, v42)
        _ = v42
    } else {
        var v43 int64 = -1
        var v44 float64 = v3
        _ = v43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
