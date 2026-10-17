package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 string }
type GenS2 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 int64 }
func genFn0(p0 int16, p1 int64, p2 bool) int8 {
    p1 += 85
    if (164 > 42) {
        if p2 {
            v4 := GenS1{F0: 4294967295, F1: 0, F2: "hello"}
            var v5 int32 = 2147483647
            switch -1 {
            case 1:
                v6 := []int{int(v5)}
                _ = v6
            case 680:
                if (838 <= 42) {
                    for fi7 := 0; fi7 < 6; fi7++ {
                        v5 = 42
                        _ = fi7
                    }
                }
            case 128:
                v8 := GenS2{F0: 255, F1: (v5 * v5), F2: 128}
                _ = v8
            default:
                if (!p2) {
                    for fi9 := 0; fi9 < 19; fi9++ {
                        p0 = p0
                        _ = fi9
                    }
                    var v10 uint64 = 0
                    _ = v10
                } else {
                    var v11 uint8 = 127
                    switch 638 {
                    case 898:
                        v12 := GenS2{F0: (v11 ^ (v11 / ((140) | 1))), F1: v5, F2: v11}
                        _ = v12
                    }
                    _ = v11
                }
            }
            _ = v4
            _ = v5
        } else {
            if (657 != 2) {
                v13 := GenS3{F0: -128, F1: 1, F2: "a"}
                var v14 float32 = 3.14159
                v14 += v14
                _ = v13
                _ = v14
            }
            if (false || p2) {
                for fi15 := 0; fi15 < 16; fi15++ {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        switch (fi16 * fi15) {
                        case 674:
                            var v17 bool = (1 <= 0)
                            _ = v17
                        default:
                            p1 += 9223372036854775807
                        }
                        _ = fi16
                    }
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v19 := []int{(fi18 / ((9223372036854775807) | 1))}
                        if true {
                            _ = len(v19)
                            v20 := GenS2{F0: 0, F1: -1, F2: 164}
                            var v21 int8 = 0
                            _ = v20
                            _ = v21
                        }
                        v22 := GenS1{F0: 885, F1: 108, F2: "a"}
                        var v23 bool = p2
                        p0 -= -735
                        _ = fi18
                        _ = v19
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(p0, p1)
                    _ = fi15
                }
                p0 -= p0
            }
            var v24 float64 = (0.1 + -2.25)
            _ = v24
        }
        p2 = (p2 && ((0.0 - 0.0) <= 0.0))
    } else {
        if ((!(false && true)) && (p0 <= p0)) {
            var v25 uint8 = 255
            switch 0 {
            case 128:
                v26 := GenS4{F0: p1}
                _ = v26
            default:
                if ((2 == 861) || p2) {
                    p1 = 0
                } else {
                    for fi27 := 0; fi27 < 9; fi27++ {
                        p1 += -40
                        _ = fi27
                    }
                    p0 -= (p0 & p0)
                }
            }
            fmt.Println(v25, v25, p1)
            _ = v25
        } else {
            v28 := GenS5{F0: 241, F1: 127}
            for fi29 := 0; fi29 < 6; fi29++ {
                v30 := GenS6{F0: (p1 % (((-(p1))) | 1))}
                _ = fi29
                _ = v30
            }
            _ = v28
        }
    }
    var v31 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    return 100
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v32 uint64 = 951
    var v33 float64 = -2.25
    switch "hello" {
    case "":
        fmt.Println(v32, v32)
    case "hello":
        var v34 float64 = (v33 * 1e10)
        _ = v34
    default:
        var v35 uint8 = 144
        _ = v35
    }
    var v36 string = "a"
    v37 := []int{42, 4294967295, 4294967295, 369}
    fmt.Println(v33, v32, v32)
    var v38 uint16 = 42
    v33 -= v33
    fmt.Println(v38, v38, v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
}
