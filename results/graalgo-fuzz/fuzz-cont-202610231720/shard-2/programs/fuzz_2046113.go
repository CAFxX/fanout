package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint16; F2 string }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int; F1 int64 }
func genFn0(p0 int, p1 int, p2 uint16) int64 {
    switch "\n\t" {
    case "":
        v5 := []int{429}
        _ = v5
    case "foo bar":
        var v6 int32 = 467
        _ = v6
    default:
        p0 = int(p2)
    }
    var v7 int32 = -1
    v8 := []int{709, p1, p1, p1}
    p1 = p0
    if (443 < 138) {
        fmt.Println(p2)
    } else {
        v7 += (v7 % ((42) | 1))
        if ((!false) || false) {
            if (true && true) {
                v7 += (^(v7))
                for fi9 := 0; fi9 < 19; fi9++ {
                    v10 := GenS1{F0: 2, F1: 213, F2: 62}
                    _ = fi9
                    _ = v10
                }
                var v11 uint8 = 2
                _ = v11
            } else {
                p0 = p1
            }
            v8 = append(v8, 32768)
            p1 = p0
        }
        switch (p1 ^ p1) {
        case 9223372036854775807:
            switch 9223372036854775807 {
            case 353:
                p1 = p1
            default:
                var v12 int64 = 1
                _ = v12
            }
        default:
            var v13 uint8 = 127
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = (^(v1))
    var v3 int16 = -32768
    v4 := make(map[int]int)
    switch ("hello" + "a") {
    case "a":
        if false {
            if true {
                if false {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v15 := []int{(fi14 / (((fi14 / ((fi14) | 1))) | 1))}
                        _ = fi14
                        _ = v15
                    }
                    var v16 int8 = 1
                    _ = v16
                }
            }
            switch 128 {
            case 127:
                v2 = v1
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                v18 := GenS2{F0: 75, F1: 409, F2: "\n\t"}
                v3 = v3
                v19 := GenS2{F0: 519, F1: 0, F2: "foo bar"}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    case "":
        if false {
            for fi20 := 0; fi20 < 6; fi20++ {
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 uint = 407
                    if true {
                        switch fi21 {
                        case 127:
                            v23 := GenS1{F0: 152, F1: 128, F2: 2}
                            _ = v23
                        case 9223372036854775807:
                            var v24 int32 = -2147483648
                            _ = v24
                        default:
                            switch 114 {
                            case -731:
                                switch fi21 {
                                case 441:
                                    if (-994 <= 1) {
                                        fmt.Println(fi21, v3)
                                        if true {
                                            var v25 int = (^(fi20))
                                            _ = v25
                                        }
                                    }
                                case 32768:
                                    v26 := GenS2{F0: 812, F1: 0, F2: ("" + (("hello" + "\n\t") + "foo bar"))}
                                    _ = v26
                                }
                            case -179:
                                var v27 int = 264
                                var v28 uint = v22
                                var v29 int8 = int8(v2)
                                _ = v27
                                _ = v28
                                _ = v29
                            default:
                                v22 = (^(v22))
                            }
                        }
                    }
                    fmt.Println(fi21, v3, v22)
                    if (false && true) {
                        if false {
                            v30 := GenS3{F0: v1, F1: "foo bar"}
                            var v31 float64 = 1.5
                            fmt.Println(fi20, v1, v2)
                            var v32 float32 = ((0.0 + 1.5) + 1e10)
                            if (true || false) {
                                if ((1.5 * (v32 * 0.0)) >= v32) {
                                    var v33 int16 = -885
                                    var v34 int64 = v2
                                    fmt.Println(v33, v31)
                                    _ = v33
                                    _ = v34
                                } else {
                                    v35 := GenS3{F0: (v1 << (((v22 / (((^(v22))) | 1))) & 63)), F1: "foo bar"}
                                    _ = v35
                                }
                            }
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        delete(v4, fi20)
                        v36 := []int{-137, (fi21 * fi20), fi21, fi21}
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v1)
                        case "foo bar":
                            var v37 uint8 = 34
                            _ = v37
                        default:
                            v1 -= genFn0(fi20, -189, 0)
                        }
                        _ = v36
                    }
                    _ = fi21
                    _ = v22
                }
                if ((87 > 279) || ((0.0 + 1.5) > 1e10)) {
                    v38 := v4[(fi20 >> ((423) & 63))]
                    _ = v38
                } else {
                    v39 := GenS4{F0: 1}
                    fmt.Println(fi20, v3)
                    _ = v39
                }
                v3 = (-(v3))
                _ = fi20
            }
            fmt.Println(v3, v2)
        } else {
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 uint16 = 2
                v2 += (^(v1))
                _ = fi40
                _ = v41
            }
        }
    }
    fmt.Println(v1, v3)
    v42 := GenS5{F0: ((("hello" + "\n\t") + "\n\t") + "\n\t"), F1: 42}
    var v43 int32 = 1
    for fi44 := 0; fi44 < 7; fi44++ {
        switch "\n\t" {
        case "a":
            fmt.Println(v2)
        case "\n\t":
            v43 = (^(v43))
        case "foo bar":
            v45 := GenS2{F0: 660, F1: 65535, F2: "a"}
            _ = v45
        default:
            v2 = (-(v2))
        }
        _ = fi44
    }
    v46 := GenS5{F0: "foo bar", F1: 184}
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 uint64 = 858
        v49 := []int{-9223372036854775808, fi47, 42}
        _ = fi47
        _ = v48
        _ = v49
    }
    v4[-256] = 1
    v1 -= (v2 << ((2) & 63))
    if true {
        var v50 int8 = -1
        _ = v50
    } else {
        v51 := GenS6{F0: -1, F1: (v1 * genFn0(0, 127, 838))}
        v52 := []int{-680, 2147483647, 42}
        if true {
            v3 += (v3 - v3)
            _ = len(v52)
            v3 += -683
        } else {
            var v53 int32 = (v43 >> ((1) & 63))
            _ = v53
        }
        _ = v51
        _ = v52
    }
    v2 += -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v46
}
