package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint; F1 int64; F2 int }
func genFn0(p0 int8) int8 {
    for fi7 := 0; fi7 < 16; fi7++ {
        switch fi7 {
        case 9223372036854775807:
            v8 := GenS1{F0: ((396 <= 2147483647) && false), F1: 367}
            _ = v8
        case -1:
            p0 += (p0 * p0)
        case -2147483648:
            if (0 != 20) {
                var v9 uint = 716
                var v10 uint = v9
                v9 += (^(v9))
                _ = v10
                _ = v9
            }
        default:
            switch 453 {
            case 9223372036854775807:
                v11 := []int{(fi7 - fi7), fi7, fi7, (fi7 ^ fi7)}
                _ = v11
            }
        }
        _ = fi7
    }
    fmt.Println(p0)
    if ((1.5 - (0.0 * -2.25)) >= (0.1 + -2.25)) {
        if ((!(2 == 204)) && (!(42 >= 115))) {
            switch "a" {
            case "hello":
                if true {
                    if true {
                        for fi12 := 0; fi12 < 7; fi12++ {
                            p0 = (p0 - p0)
                            v13 := GenS2{F0: 18446744073709551615, F1: 0, F2: ""}
                            p0 = (p0 + p0)
                            _ = fi12
                            _ = v13
                        }
                        p0 = p0
                    }
                } else {
                    switch "a" {
                    case "a":
                        if (127 > 255) {
                            var v14 bool = (110 <= 164)
                            var v15 int32 = 531
                            _ = v14
                            _ = v15
                        }
                    case "hello":
                        for fi16 := 0; fi16 < 13; fi16++ {
                            if false {
                                v17 := []int{}
                                var v18 uint64 = 4294967295
                                switch fi16 {
                                case -9223372036854775808:
                                    v19 := GenS1{F0: true, F1: 9223372036854775807}
                                    _ = v19
                                case 32768:
                                    v20 := GenS1{F0: (855 < -1), F1: -48}
                                    _ = v20
                                default:
                                    var v21 int = fi16
                                    _ = v21
                                }
                                _ = v17
                                _ = v18
                            } else {
                                fmt.Println(fi16)
                            }
                            v22 := GenS2{F0: 813, F1: 2, F2: "\n\t"}
                            p0 = (p0 + (-(p0)))
                            _ = fi16
                            _ = v22
                        }
                    }
                    fmt.Println(p0)
                }
            case "":
                for fi23 := 0; fi23 < 2; fi23++ {
                    fmt.Println(fi23, p0)
                    var v24 int32 = 1
                    _ = fi23
                    _ = v24
                }
            case "foo bar":
                fmt.Println(p0)
            default:
                fmt.Println(p0)
            }
            switch 127 {
            case -254:
                v25 := []int{-393}
                _ = v25
            case 42:
                v26 := GenS1{F0: false, F1: -874}
                _ = v26
            default:
                v27 := []int{}
                _ = v27
            }
            v28 := GenS3{F0: "a", F1: 65535}
            _ = v28
        }
        var v29 int16 = -343
        _ = v29
    }
    _ = p0
    return (-(p0))
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -1
    var v3 int = 253
    v4 := make(map[string]bool)
    v5 := make(map[uint32]uint32)
    var v6 float64 = (-2.25 / (1.5 + 3.14159))
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 uint64 = 4294967295
        var v32 int = 1
        var v33 float32 = (1e10 * 0.0)
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        switch 32768 {
        case 127:
            var v35 uint8 = 2
            _ = v35
        }
        v36 := GenS4{F0: 65535, F1: 42, F2: (fi34 / (((v3 & (^(fi34)))) | 1))}
        _ = fi34
        _ = v36
    }
    var v37 float64 = (0.0 * v6)
    v3 = v3
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 int = -9223372036854775808
        fmt.Println(v2, v3, v3)
        fmt.Println(fi38)
        _ = fi38
        _ = v39
    }
    v40 := []int{9223372036854775807, v3, (v3 & (v3 % ((v3) | 1)))}
    fmt.Println(v2)
    _ = len(v40)
    var v41 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
