package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
type GenS2 struct { F0 int8; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int16; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint16, p1 int8, p2 uint16) uint8 {
    switch ("hello" + "a") {
    case "\n\t":
        fmt.Println(p0)
    case "foo bar":
        var v6 float64 = 1e10
        _ = v6
    }
    v7 := []int{127, -279, 290}
    v8 := GenS1{F0: (p1 & p1), F1: 50, F2: 42}
    if ((^(p2)) <= (^(p0))) {
        v9 := GenS1{F0: (p1 ^ p1), F1: 379, F2: 0}
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                switch "hello" {
                case "\n\t":
                    var v12 float32 = float32((1.5 * 0.0))
                    var v13 int8 = (p1 | 0)
                    fmt.Println(fi11, p1)
                    _ = v12
                    _ = v13
                }
                _ = fi11
            }
            _ = fi10
        }
        p1 -= p1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint64, p1 bool) int {
    p0 -= p0
    p1 = ((!(p1 || (!(p1 && p1)))) || p1)
    var v14 int16 = -541
    _ = p0
    _ = p1
    _ = v14
    return 127
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]uint32)
    var v3 int = -2147483648
    v4 := make(map[string]string)
    var v5 int64 = 867
    v15 := []int{v3, v3, v3, -2147483648}
    v3 += -2147483648
    switch "\n\t" {
    case "a":
        if true {
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 uint8 = 128
                if (2 > (v17 + v17)) {
                    if true {
                        delete(v4, "\n\t")
                    } else {
                        v15 = append(v15, v3)
                        if (!(int8(v5) < 127)) {
                            if false {
                                v18 := GenS2{F0: 1, F1: 1, F2: v5}
                                _ = v18
                            }
                        } else {
                            if true {
                                v3 -= (v3 & 32767)
                                fmt.Println(fi16, v3, v1)
                                v5 -= -9223372036854775808
                            }
                            if (32767 < 255) {
                                var v19 uint32 = 2147483647
                                _ = v19
                            }
                        }
                        v20 := v2[v3]
                        _ = v20
                    }
                }
                _ = fi16
                _ = v17
            }
        }
    case "\n\t":
        v5 = (^(v5))
    case "foo bar":
        v5 += (^(v5))
    default:
        v21 := GenS1{F0: -128, F1: (v1 / ((v1) | 1)), F2: -1}
        _ = v21
    }
    switch "foo bar" {
    case "a":
        for fi22 := 0; fi22 < 12; fi22++ {
            _ = len(v4)
            var v23 int8 = -95
            var v24 uint16 = 820
            _ = fi22
            _ = v23
            _ = v24
        }
    case "":
        var v25 uint8 = 1
        _ = v25
    case "\n\t":
        switch v3 {
        case 612:
            for fi26 := 0; fi26 < 3; fi26++ {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                fmt.Println(v3, v5, v1)
                _ = len(v15)
                _ = fi26
            }
        default:
            fmt.Println(v1, v5)
        }
    }
    v27 := GenS3{F0: 255, F1: 1000}
    var v28 float32 = 0.1
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    if true {
        fmt.Println(v1)
        fmt.Println(v5, v1, v28)
    } else {
        if false {
            var v29 int32 = (v1 & (v1 ^ 1))
            _ = v29
        }
        v30 := GenS1{F0: -1, F1: v1, F2: 1000}
        v1 = v1
        _ = v30
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := GenS1{F0: 42, F1: v1, F2: 1}
        switch "hello" {
        case "":
            v33 := GenS4{F0: 32767, F1: 65535, F2: 42}
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    if len(v15) > 0 {
        v15[0] = genFn1(629, (437 <= 1))
    }
    if (1e10 == v28) {
        var v34 uint32 = 1
        _ = v34
    } else {
        v35 := GenS1{F0: 127, F1: 2147483647, F2: 446}
        _ = v35
    }
    v36 := GenS5{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
