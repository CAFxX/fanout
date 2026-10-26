package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 uint16; F1 int; F2 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint32 }

func main() {
    var v1 float32 = 1.5
    var v2 uint64 = 1
    v3 := []int{32768, -576, -511, 1}
    var v4 int = 32767
    v5 := []int{}
    v6 := make(map[int]int)
    v4 = (v4 + v4)
    v1 -= (float32(((0.0 - float64(v1)) * 0.0)) - 3.14159)
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 11; fi8++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if false {
            var v9 uint8 = 0
            v2 += 2
            v10 := v6[fi8]
            _ = v10
            _ = v9
        }
        v2 -= 1
        _ = fi8
    }
    v6[v4] = 42
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 string = "\n\t"
        _ = fi11
        _ = v12
    }
    switch "\n\t" {
    case "hello":
        for fi13 := 0; fi13 < 20; fi13++ {
            v1 = 1e10
            switch "hello" {
            case "hello":
                v1 += 1.5
            case "a":
                switch v4 {
                case -9223372036854775808:
                    v1 -= v1
                case -1:
                    fmt.Println(v4, v2)
                default:
                    if (!(true && false)) {
                        var v14 uint16 = 446
                        var v15 int32 = 297
                        _ = v14
                        _ = v15
                    }
                }
            case "foo bar":
                var v16 int16 = 1
                _ = v16
            default:
                _ = len(v5)
            }
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 uint32 = 1
                if (27 < 95) {
                    v19 := GenS1{F0: -489}
                    v18 -= (v18 >> ((4294967295) & 63))
                    v20 := GenS1{F0: -1}
                    delete(v6, fi17)
                    fmt.Println(fi17, v18)
                    _ = v19
                    _ = v20
                }
                v21 := GenS1{F0: -1}
                _ = fi17
                _ = v18
                _ = v21
            }
            _ = fi13
        }
    default:
        fmt.Println(v2, v1, v4)
    }
    if (!(true && (127 != 255))) {
        v4 = (-(v4))
        _ = len(v5)
    } else {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 int32 = -2147483648
            v23 = (^(v23))
            _ = len(v5)
            _ = fi22
            _ = v23
        }
        var v24 int = (v4 % (((v4 & -9223372036854775808)) | 1))
        v25 := v6[v4]
        _ = v24
        _ = v25
    }
    switch v4 {
    case -2147483648:
        var v26 float32 = v1
        _ = v26
    default:
        _ = len(v5)
    }
    var v27 string = "\n\t"
    switch v27 {
    case "":
        v28 := GenS2{F0: 648}
        _ = v28
    default:
        switch 1 {
        case 2147483647:
            var v29 string = ("a" + "\n\t")
            _ = v29
        case 32767:
            v30 := GenS2{F0: -2147483648}
            _ = v30
        case 42:
            if false {
                switch v4 {
                case 189:
                    var v31 int16 = -493
                    _ = v31
                case 4294967295:
                    v32 := GenS2{F0: 490}
                    _ = v32
                case -800:
                    switch v27 {
                    case "\n\t":
                        v33 := GenS3{F0: (v2 & v2)}
                        _ = v33
                    case "hello":
                        for fi34 := 0; fi34 < 17; fi34++ {
                            if true {
                                if true {
                                    v27 = (("" + v27) + v27)
                                    switch v27 {
                                    case "a":
                                        v35 := GenS4{F0: v4, F1: 1}
                                        _ = v35
                                    case "\n\t":
                                        if (950 > -1) {
                                            v36 := GenS5{F0: 76, F1: 2147483647}
                                            var v37 int64 = -9223372036854775808
                                            _ = v36
                                            _ = v37
                                        }
                                    default:
                                        v27 = "a"
                                    }
                                    fmt.Println(fi34, v2, v2)
                                } else {
                                    v38 := GenS2{F0: 1}
                                    _ = v38
                                }
                                v4 -= v4
                                if false {
                                    v1 += (1e10 - v1)
                                    if true {
                                        if true {
                                            v1 = v1
                                            v4 += fi34
                                            v6[42] = v4
                                        }
                                    } else {
                                        var v39 uint = 4294967295
                                        for fi40 := 0; fi40 < 10; fi40++ {
                                            if false {
                                                if (!false) {
                                                    var v41 bool = (!((((v27 + "a") + "foo bar") + v27) < "\n\t"))
                                                    var v42 uint32 = 919
                                                    var v43 int32 = 2147483647
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                                v1 -= v1
                                            }
                                            _ = fi40
                                        }
                                        _ = v39
                                    }
                                }
                            }
                            v4 = 628
                            if ((361 == 1) && (32767 >= (v4 & (fi34 >> ((42) & 63))))) {
                                v4 = fi34
                            } else {
                                fmt.Println(v27)
                                v44 := GenS2{F0: 42}
                                fmt.Println(fi34)
                                _ = v44
                            }
                            _ = fi34
                        }
                    default:
                        if true {
                            _ = len(v6)
                            _ = len(v6)
                        }
                    }
                }
                var v45 float64 = 0.0
                var v46 int64 = -1
                _ = v45
                _ = v46
            } else {
                var v47 uint = uint(v2)
                for fi48 := 0; fi48 < 4; fi48++ {
                    switch 32767 {
                    case -9223372036854775808:
                        v47 += (v47 | 252)
                    case 1:
                        v27 = ""
                    case 32767:
                        var v49 int64 = -1
                        _ = v49
                    }
                    v2 = v2
                    _ = fi48
                }
                var v50 int32 = 1
                _ = v47
                _ = v50
            }
        default:
            var v51 float32 = float32(float64(((1e10 * (v1 + ((0.0 - v1) * 1.5))) - v1)))
            _ = v51
        }
    }
    if false {
        if (v27 < ("" + v27)) {
            fmt.Println(v4, v4, v1)
        }
        fmt.Println(v2)
        if true {
            var v52 float64 = ((float64(v1) / float64(((1.5 + float32(-2.25)) * v1))) + 3.14159)
            var v53 float64 = float64((float32(1e10) + float32((float64(((v1 * v1) + v1)) * 1.5))))
            _ = v52
            _ = v53
        }
    } else {
        switch ((v27 + "a") + (v27 + v27)) {
        case "hello":
            for fi54 := 0; fi54 < 12; fi54++ {
                v1 = (0.0 * float32(float64(0.0)))
                v55 := v6[(v4 >> ((18446744073709551615) & 63))]
                if true {
                    switch "foo bar" {
                    case "hello":
                        if true {
                            for fi56 := 0; fi56 < 18; fi56++ {
                                v57 := v6[(v4 | int(v2))]
                                _ = fi56
                                _ = v57
                            }
                            var v58 float64 = ((0.1 / 0.1) + 3.14159)
                            _ = v58
                        }
                    }
                    v59 := GenS6{F0: 2, F1: (v55 / ((v4) | 1)), F2: 127}
                    var v60 int = (v55 * fi54)
                    _ = v59
                    _ = v60
                } else {
                    v55 -= 4294967295
                    var v61 float64 = -2.25
                    _ = v61
                }
                _ = fi54
                _ = v55
            }
        case "\n\t":
            for fi62 := 0; fi62 < 8; fi62++ {
                var v63 string = "hello"
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        v1 = v1
                    }
                    var v64 int32 = 0
                    _ = v64
                } else {
                    v27 = "\n\t"
                }
                _ = fi62
                _ = v63
            }
        case "foo bar":
            if (2147483647 > v4) {
                var v65 uint = uint(v2)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v65
            } else {
                v4 += (^(v4))
                var v66 bool = false
                var v67 float64 = 1.5
                _ = v66
                _ = v67
            }
        default:
            v1 += -2.25
        }
    }
    v1 += ((0.1 * (v1 - float32(0.1))) + v1)
    switch (v4 << ((1) & 63)) {
    case 9223372036854775807:
        fmt.Println(v27, v1, v4)
    case -2147483648:
        v68 := GenS5{F0: 2, F1: 4294967295}
        _ = v68
    case 286:
        for fi69 := 0; fi69 < 17; fi69++ {
            fmt.Println(v27, fi69, fi69)
            v2 += v2
            if true {
                var v70 uint16 = 0
                _ = v70
            } else {
                for fi71 := 0; fi71 < 11; fi71++ {
                    switch 1 {
                    case -804:
                        var v72 int16 = 1000
                        _ = v72
                    case 0:
                        _ = len(v3)
                        var v73 int16 = 1000
                        if (-1 != 127) {
                            var v74 int64 = -837
                            if ((42 >= 42) || (2147483648 > 2147483647)) {
                                if ((v27 < v27) || true) {
                                    _ = len(v6)
                                    delete(v6, (fi71 * fi69))
                                    v2 = 642
                                }
                            }
                            if ((-2.25 / 3.14159) != 3.14159) {
                                v73 = (v73 | v73)
                                var v75 float32 = v1
                                _ = v75
                            } else {
                                if (false && true) {
                                    var v76 uint32 = 658
                                    v77 := GenS6{F0: 7, F1: (fi69 % (((fi71 & (v4 << ((252) & 63)))) | 1)), F2: (-(v4))}
                                    _ = v76
                                    _ = v77
                                } else {
                                    var v78 uint8 = 93
                                    _ = v78
                                }
                                var v79 bool = (!(0 <= -128))
                                if true {
                                    v27 = ("" + v27)
                                }
                                _ = v79
                            }
                            var v80 int8 = -99
                            if (2 != 4294967295) {
                                switch (v27 + v27) {
                                case "":
                                    v74 = v74
                                case "hello":
                                    fmt.Println(v1, v80, v1)
                                case "foo bar":
                                    if true {
                                        fmt.Println(v4)
                                        fmt.Println(v4, v73, v74)
                                        v74 -= int64(fi69)
                                        v81 := GenS4{F0: v4, F1: fi69}
                                        v1 = (1.5 - 0.1)
                                        v2 = v2
                                        v1 -= v1
                                        _ = v81
                                    } else {
                                        fmt.Println(v4, v1, v4)
                                        var v82 int16 = (v73 % ((-32768) | 1))
                                        _ = v82
                                    }
                                }
                                if (255 != 47) {
                                    var v83 int8 = 100
                                    switch (v27 + "a") {
                                    case "":
                                        if (1 != 42) {
                                            var v84 uint32 = 2
                                            v85 := GenS4{F0: (v4 | v4), F1: 1}
                                            v74 -= v74
                                            _ = v84
                                            _ = v85
                                        }
                                    }
                                    v86 := GenS3{F0: v2}
                                    _ = v83
                                    _ = v86
                                } else {
                                    if ((!true) || true) {
                                        _ = len(v6)
                                    } else {
                                        var v87 bool = (false || false)
                                        _ = v87
                                    }
                                    v88 := GenS3{F0: (v2 + (v2 | v2))}
                                    v80 += (v80 ^ v80)
                                    _ = v88
                                }
                            } else {
                                v89 := GenS7{F0: (v80 % ((-1) | 1))}
                                _ = v89
                            }
                            v90 := GenS1{F0: 252}
                            _ = v74
                            _ = v80
                            _ = v90
                        } else {
                            var v91 bool = (1 < 42)
                            fmt.Println(v27)
                            _ = v91
                        }
                        _ = v73
                    default:
                        switch "hello" {
                        case "\n\t":
                            _ = len(v5)
                            var v92 bool = (false && (-223 == 1))
                            var v93 int64 = 987
                            _ = v92
                            _ = v93
                        case "hello":
                            _ = len(v6)
                        }
                    }
                    if (v27 >= (((((v27 + "\n\t") + ("hello" + "a")) + "") + (v27 + (v27 + "a"))) + v27)) {
                        var v94 uint16 = 42
                        switch -871 {
                        case 774:
                            if true {
                                v5 = append(v5, v4)
                                v95 := GenS8{F0: -1, F1: (v1 < 1e10)}
                                _ = v95
                            } else {
                                fmt.Println(v1)
                                v94 -= v94
                                v96 := GenS1{F0: 808}
                                _ = v96
                            }
                        case -483:
                            v1 = (-2.25 * v1)
                        case 32768:
                            switch (((v27 + v27) + ("\n\t" + "a")) + v27) {
                            case "hello":
                                var v97 int = fi71
                                _ = v97
                            }
                        default:
                            switch "a" {
                            case "foo bar":
                                if (fi71 <= fi71) {
                                    v98 := GenS6{F0: (v94 ^ (v94 - v94)), F1: fi71, F2: fi71}
                                    _ = v98
                                }
                            }
                        }
                        _ = v94
                    }
                    _ = fi71
                }
                fmt.Println(fi69)
                switch (((v27 + "") + v27) + v27) {
                case "hello":
                    fmt.Println(fi69)
                case "a":
                    v99 := GenS9{F0: 2}
                    _ = v99
                case "\n\t":
                    if true {
                        var v100 bool = false
                        _ = v100
                    } else {
                        delete(v6, v4)
                    }
                }
            }
            _ = fi69
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
