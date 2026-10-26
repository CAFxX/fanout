package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool; F1 string; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 uint; F2 int32 }
type GenS6 struct { F0 uint16; F1 string; F2 int64 }
type GenS7 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0(p0 int64) int {
    var v4 uint64 = 1
    for fi5 := 0; fi5 < 19; fi5++ {
        v4 = 2
        for fi6 := 0; fi6 < 4; fi6++ {
            if false {
                p0 -= (p0 + -9223372036854775808)
                var v7 int8 = -128
                _ = v7
            }
            _ = fi6
        }
        switch "foo bar" {
        case "hello":
            var v8 int32 = -1
            _ = v8
        case "foo bar":
            if (1000 <= -1) {
                if false {
                    switch fi5 {
                    case -520:
                        switch (fi5 - (fi5 / ((127) | 1))) {
                        case 974:
                            if true {
                                fmt.Println(p0)
                                v4 = v4
                                var v9 float32 = (-2.25 * 1.5)
                                _ = v9
                            } else {
                                var v10 string = "\n\t"
                                var v11 float32 = 1e10
                                v10 = "hello"
                                _ = v10
                                _ = v11
                            }
                        case 42:
                            v12 := GenS1{F0: 0, F1: int64(v4)}
                            _ = v12
                        }
                    case -849:
                        fmt.Println(v4, fi5, v4)
                    }
                    var v13 float32 = float32(3.14159)
                    _ = v13
                }
                v14 := GenS1{F0: 592, F1: (-(p0))}
                var v15 uint16 = 985
                _ = v14
                _ = v15
            }
        }
        _ = fi5
    }
    fmt.Println(p0, v4)
    for fi16 := 0; fi16 < 2; fi16++ {
        p0 -= (p0 >> ((229) & 63))
        fmt.Println(fi16, fi16)
        _ = fi16
    }
    p0 = (p0 / ((-9223372036854775808) | 1))
    _ = p0
    _ = v4
    return -1
}
func genFn1(p0 uint8) int8 {
    p0 += (p0 >> ((42) & 63))
    if (4294967295 > 2) {
        v17 := GenS1{F0: 42, F1: 1}
        if true {
            var v18 uint32 = 2147483648
            _ = v18
        } else {
            var v19 int = -758
            _ = v19
        }
        _ = v17
    } else {
        var v20 int8 = 127
        fmt.Println(v20, v20)
        switch "hello" {
        case "":
            var v21 int64 = 0
            _ = v21
        case "a":
            v22 := GenS1{F0: 42, F1: -9223372036854775808}
            _ = v22
        }
        _ = v20
    }
    p0 += 128
    _ = p0
    return -128
}

func main() {
    var v1 uint32 = 0
    var v2 float64 = 1.5
    var v3 string = "foo bar"
    fmt.Println(v3, v1)
    var v23 int = -9223372036854775808
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int32 = 42
        _ = fi24
        _ = v25
    }
    switch (v23 >> ((4294967295) & 63)) {
    case -787:
        v2 = 1.5
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        if false {
            v27 := GenS2{F0: ((false || (false || (((!(v23 < 871)) || (v3 < v3)) && (1 < fi26)))) || true), F1: "", F2: 127}
            var v28 uint32 = (v1 + 2)
            v29 := GenS1{F0: 1, F1: 712}
            _ = v27
            _ = v28
            _ = v29
        } else {
            v30 := []int{(v23 | genFn0(9223372036854775807)), fi26, 4294967295}
            var v31 string = v3
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            _ = v30
            _ = v31
        }
        if false {
            fmt.Println(fi26, v23, v23)
        } else {
            for fi32 := 0; fi32 < 3; fi32++ {
                var v33 uint64 = 18446744073709551615
                _ = fi32
                _ = v33
            }
            for fi34 := 0; fi34 < 9; fi34++ {
                if false {
                    var v35 int = (v23 | fi26)
                    var v36 string = v3
                    switch genFn0(0) {
                    case 32767:
                        var v37 float32 = float32(v2)
                        _ = v37
                    case 9223372036854775807:
                        v38 := []int{-730, (fi26 + (v35 % ((v23) | 1))), (fi26 ^ genFn0(9223372036854775807)), 158}
                        _ = v38
                    case 0:
                        v2 = (((((v2 / v2) - float64((((float32(3.14159) * (0.0 + ((((0.0 / float32(v2)) - float32(float64((float32((v2 + v2)) - (((1.5 + -2.25) * 1.5) / ((float32(v2) * 0.0) - (3.14159 - -2.25))))))) * 1.5) - float32((v2 - v2))))) * 1e10) + (((((float32(v2) * 1.5) * (float32((((((3.14159 - v2) - v2) / 0.1) + (v2 + v2)) * v2)) * (float32((v2 * 0.1)) * ((0.0 + (0.1 / 1.0)) * ((0.1 / 1e10) / 1.0))))) + 1.5) / (1.5 * 3.14159)) + (1e10 - ((-2.25 * float32(v2)) - (((-2.25 / 1e10) - 1.5) + 1.5))))))) + v2) + (v2 + (v2 - (((v2 + 3.14159) - float64(((1e10 + (0.0 * 3.14159)) / 1.5))) - float64((float32(float64(float32(v2))) / -2.25)))))) / (0.1 / ((0.0 / float64(((1.5 + 0.0) - -2.25))) + v2)))
                    default:
                        var v39 bool = (false || (int(v1) > 656))
                        _ = v39
                    }
                    _ = v35
                    _ = v36
                } else {
                    if true {
                        if (255 != 238) {
                            v23 -= v23
                        }
                        v40 := GenS3{F0: 729, F1: 2}
                        v2 -= (v2 - v2)
                        _ = v40
                    }
                    if true {
                        switch (v3 + v3) {
                        case "\n\t":
                            var v41 bool = false
                            _ = v41
                        }
                    }
                }
                _ = fi34
            }
            for fi42 := 0; fi42 < 10; fi42++ {
                var v43 string = (v3 + ((v3 + v3) + v3))
                _ = fi42
                _ = v43
            }
        }
        v44 := GenS4{F0: v23}
        _ = fi26
        _ = v44
    }
    for fi45 := 0; fi45 < 3; fi45++ {
        var v46 string = v3
        fmt.Println(v3)
        v47 := GenS5{F0: -498, F1: 42, F2: 1}
        _ = fi45
        _ = v46
        _ = v47
    }
    if true {
        if (241 < 2) {
            fmt.Println(v3, v23)
            v23 += v23
        }
        for fi48 := 0; fi48 < 15; fi48++ {
            if (-16 > -2147483648) {
                if false {
                    v23 += fi48
                    switch v3 {
                    case "\n\t":
                        v3 = v3
                    case "a":
                        if (false && true) {
                            for fi49 := 0; fi49 < 6; fi49++ {
                                v50 := []int{127, (fi48 >> ((23) & 63))}
                                _ = fi49
                                _ = v50
                            }
                        }
                    case "foo bar":
                        v51 := GenS6{F0: 2, F1: v3, F2: 9223372036854775807}
                        _ = v51
                    }
                }
                v1 += v1
                var v52 float32 = ((0.1 / ((0.1 / (float32((0.1 + (v2 + v2))) + -2.25)) * (0.1 - (0.1 + -2.25)))) + 3.14159)
                _ = v52
            } else {
                v53 := GenS7{F0: 203, F1: 127, F2: 32767}
                _ = v53
            }
            _ = fi48
        }
        v3 = v3
    }
    v54 := []int{(^(v23))}
    fmt.Println(v3)
    v55 := GenS8{F0: 4294967295, F1: 65535}
    for fi56 := 0; fi56 < 1; fi56++ {
        for fi57 := 0; fi57 < 16; fi57++ {
            var v58 uint32 = 683
            switch (-(v23)) {
            case -9223372036854775808:
                switch -1 {
                case -14:
                    if true {
                        if true {
                            v23 = fi56
                        }
                        var v59 int32 = int32(fi57)
                        switch (v3 + v3) {
                        case "":
                            var v60 uint64 = 515
                            _ = v60
                        case "\n\t":
                            fmt.Println(fi57)
                        case "hello":
                            if true {
                                fmt.Println(fi57, v1)
                            }
                        }
                        _ = v59
                    }
                case 524:
                    if true {
                        switch 4294967295 {
                        case 9223372036854775807:
                            v58 = 42
                        case -2147483648:
                            if false {
                                var v61 uint8 = 2
                                _ = v61
                            } else {
                                if true {
                                    if (float32(0.0) >= -2.25) {
                                        v62 := GenS1{F0: 1000, F1: 0}
                                        switch v3 {
                                        case "hello":
                                            var v63 uint16 = 730
                                            _ = v63
                                        case "":
                                            v58 = 410
                                        default:
                                            v2 = (0.1 / v2)
                                        }
                                        _ = v62
                                    }
                                    if len(v54) > 0 {
                                        v54[0] = (fi57 << ((2) & 63))
                                    }
                                    v23 = 1
                                    fmt.Println(v58, fi57, v58)
                                    var v64 int = fi56
                                    _ = v64
                                }
                                v65 := GenS3{F0: v58, F1: 680}
                                _ = v65
                            }
                        }
                        v3 = ""
                        var v66 int = v23
                        _ = v66
                    } else {
                        var v67 uint8 = 1
                        v2 = v2
                        _ = v67
                    }
                }
            case -2147483648:
                v68 := GenS9{F0: (!(42 >= 18446744073709551615))}
                _ = v68
            case 128:
                var v69 uint8 = 97
                _ = v69
            }
            fmt.Println(fi57, v23, v2)
            _ = fi57
            _ = v58
        }
        _ = fi56
    }
    var v70 int8 = 73
    switch v23 {
    case -674:
        if ((!false) && false) {
            var v71 float64 = 1e10
            for fi72 := 0; fi72 < 9; fi72++ {
                switch ((((v3 + v3) + "a") + "") + "hello") {
                case "a":
                    fmt.Println(v23, v3)
                default:
                    if false {
                        for fi73 := 0; fi73 < 6; fi73++ {
                            switch v3 {
                            case "foo bar":
                                var v74 string = v3
                                _ = v74
                            case "":
                                if ((v1 << ((18446744073709551615) & 63)) != 2147483648) {
                                    var v75 int8 = 42
                                    _ = v75
                                } else {
                                    v71 = ((v2 * v2) - 0.0)
                                    v76 := GenS10{F0: v1, F1: 1}
                                    if len(v54) > 0 {
                                        v54[0] = genFn0(-439)
                                    }
                                    _ = v76
                                }
                            }
                            var v77 int32 = 2147483647
                            _ = fi73
                            _ = v77
                        }
                        if (true && (!false)) {
                            fmt.Println(v70)
                            v23 = (fi72 % ((genFn0(1)) | 1))
                            var v78 string = v3
                            _ = v78
                        } else {
                            var v79 uint = 334
                            for fi80 := 0; fi80 < 13; fi80++ {
                                v81 := GenS11{F0: 42, F1: 42, F2: (v70 & (^(v70)))}
                                _ = fi80
                                _ = v81
                            }
                            _ = v79
                        }
                        var v82 uint64 = 1
                        _ = v82
                    } else {
                        switch "" {
                        case "\n\t":
                            v83 := GenS9{F0: false}
                            _ = v83
                        }
                    }
                }
                v71 += float64(0.0)
                _ = fi72
            }
            var v84 uint = 4294967295
            _ = v71
            _ = v84
        }
    }
    var v85 string = (v3 + "")
    var v86 float32 = 0.1
    fmt.Println(v70, v85)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v70:", v70)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v54
    _ = v55
    _ = v70
    _ = v85
    _ = v86
}
