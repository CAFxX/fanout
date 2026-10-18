package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 string; F1 string; F2 uint64 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int8 }
func genFn0() uint {
    v7 := GenS1{F0: 18446744073709551615, F1: ((true && true) && true), F2: 2}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: 337, F1: true, F2: 2}
        _ = fi8
        _ = v9
    }
    _ = v7
    return 4294967295
}
func genFn1(p0 uint8, p1 int64, p2 string) int16 {
    var v10 float64 = float64((0.0 - (-2.25 * (float32((float64(((3.14159 / (((float32(1e10) * 0.0) - -2.25) * 0.1)) + 0.0)) * -2.25)) * 0.1))))
    for fi11 := 0; fi11 < 18; fi11++ {
        v10 += ((v10 + 0.0) * float64(float32(v10)))
        var v12 int8 = -26
        _ = fi11
        _ = v12
    }
    v13 := []int{-1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    return 529
}

func main() {
    v1 := []int{-779}
    v2 := make(map[int]string)
    var v3 float32 = 3.14159
    var v4 int32 = 9
    var v5 float64 = 1e10
    v6 := make(map[int]bool)
    v4 -= (v4 ^ (v4 ^ 1))
    v14 := GenS1{F0: 0, F1: (42 <= 2), F2: 127}
    v5 += (float64(v3) - v5)
    var v15 uint64 = 74
    switch 32767 {
    case 127:
        if false {
            v4 = v4
            if (0 < -19) {
                v15 = (^(v15))
                if true {
                    var v16 float64 = float64(-2.25)
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 uint64 = (v15 << ((269) & 63))
                        _ = fi17
                        _ = v18
                    }
                    var v19 float64 = 0.0
                    _ = v16
                    _ = v19
                } else {
                    v3 = float32(3.14159)
                    v3 += v3
                    for fi20 := 0; fi20 < 14; fi20++ {
                        if (false && false) {
                            switch "foo bar" {
                            case "a":
                                var v21 uint8 = 217
                                _ = v21
                            case "\n\t":
                                if ((true || false) && (!true)) {
                                    v3 = v3
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case "foo bar":
                                v4 -= v4
                            default:
                                v3 += (1.5 - v3)
                            }
                            if (!false) {
                                var v22 uint64 = (v15 ^ v15)
                                _ = v22
                            }
                            v23 := GenS1{F0: 150, F1: false, F2: 127}
                            _ = v23
                        } else {
                            var v24 float32 = (0.1 * 0.1)
                            v25 := GenS1{F0: 112, F1: true, F2: 1}
                            v26 := GenS2{F0: 127}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v27 := GenS2{F0: 128}
                        fmt.Println(v15)
                        _ = fi20
                        _ = v27
                    }
                }
            }
            v15 = v15
        }
    }
    v4 += v4
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int8 = 100
        _ = fi28
        _ = v29
    }
    if false {
        for fi30 := 0; fi30 < 20; fi30++ {
            if false {
                var v31 uint64 = 18446744073709551615
                _ = v31
            }
            v15 += v15
            _ = fi30
        }
        if len(v1) > 0 {
            v1[0] = 42
        }
        if false {
            fmt.Println(v3, v4)
            v32 := GenS3{F0: v4, F1: 385}
            switch "a" {
            case "hello":
                switch 127 {
                case 127:
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v34 := GenS4{F0: (("\n\t" + (("a" + ("\n\t" + "a")) + "")) + ("" + "")), F1: ("\n\t" + "hello"), F2: v15}
                        _ = fi33
                        _ = v34
                    }
                case -2147483648:
                    var v35 uint8 = 0
                    _ = v35
                }
            case "\n\t":
                for fi36 := 0; fi36 < 13; fi36++ {
                    v4 += (v4 * (v4 & 2147483647))
                    v37 := GenS1{F0: 1, F1: false, F2: 0}
                    for fi38 := 0; fi38 < 8; fi38++ {
                        v4 -= (^(v4))
                        _ = fi38
                    }
                    _ = fi36
                    _ = v37
                }
            }
            _ = v32
        }
    } else {
        var v39 uint16 = 415
        _ = v39
    }
    fmt.Println(v5)
    switch "foo bar" {
    case "\n\t":
        v5 = -2.25
    }
    v5 -= v5
    for fi40 := 0; fi40 < 16; fi40++ {
        v41 := GenS4{F0: (("a" + "foo bar") + "hello"), F1: "a", F2: 0}
        v42 := GenS5{F0: v4, F1: (fi40 >= fi40)}
        var v43 uint = 255
        _ = fi40
        _ = v41
        _ = v42
        _ = v43
    }
    if false {
        switch "hello" {
        case "foo bar":
            for fi44 := 0; fi44 < 6; fi44++ {
                var v45 int64 = 325
                switch 2147483647 {
                case 9223372036854775807:
                    if (v5 == v5) {
                        var v46 float32 = v3
                        _ = v46
                    } else {
                        if false {
                            v5 = ((3.14159 - v5) * 0.0)
                        } else {
                            v47 := GenS6{F0: -84}
                            for fi48 := 0; fi48 < 3; fi48++ {
                                var v49 int = fi48
                                _ = fi48
                                _ = v49
                            }
                            _ = v47
                        }
                        v45 = (v45 ^ 380)
                        var v50 int32 = v4
                        _ = v50
                    }
                case 128:
                    v51 := GenS1{F0: 965, F1: (0 <= 840), F2: 222}
                    _ = v51
                default:
                    v45 = (v45 / ((v45) | 1))
                }
                var v52 int64 = v45
                _ = fi44
                _ = v45
                _ = v52
            }
        default:
            v5 -= v5
        }
        _ = len(v1)
        if (!((940 < 65535) && true)) {
            fmt.Println(v5, v5, v3)
            var v53 uint = 1
            var v54 int8 = 1
            _ = v53
            _ = v54
        }
    }
    v1 = append(v1, 32767)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
