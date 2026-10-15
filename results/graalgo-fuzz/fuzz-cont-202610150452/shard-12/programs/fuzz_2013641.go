package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint32; F2 bool }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint16) int64 {
    p0 = p0
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(p0, fi7)
        _ = fi7
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint16 = 737
    var v2 string = ("" + "")
    var v3 int64 = 0
    var v4 int8 = 27
    var v5 float32 = -2.25
    var v6 float64 = 1e10
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int64 = v3
        var v10 uint = 4294967295
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        if false {
            var v12 uint16 = (v1 + 32767)
            v13 := []int{fi11, fi11, fi11, fi11}
            v2 = "foo bar"
            _ = v12
            _ = v13
        }
        if (125 >= 127) {
            switch v2 {
            case "":
                for fi14 := 0; fi14 < 14; fi14++ {
                    v3 -= (-(v3))
                    v15 := GenS1{F0: v2, F1: 65535, F2: (v4 * v4)}
                    var v16 string = v2
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "a":
                if false {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        var v18 int64 = -1
                        v19 := []int{fi11, fi17, 51}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    var v20 float64 = (((v6 * 1.5) / 3.14159) / (v6 - v6))
                    v6 = v20
                    _ = v20
                } else {
                    if ((v2 + v2) == ("foo bar" + "hello")) {
                        for fi21 := 0; fi21 < 9; fi21++ {
                            fmt.Println(v1, v1, fi11)
                            v6 = (v6 + (v6 + ((1e10 + 0.1) + (v6 + float64(v5)))))
                            switch (v2 + ((((v2 + v2) + (("" + ((v2 + v2) + ((v2 + v2) + v2))) + v2)) + "") + v2)) {
                            case "\n\t":
                                if (!(false || false)) {
                                    switch fi11 {
                                    case -1:
                                        v1 = (v1 * (^(v1)))
                                    default:
                                        var v22 int8 = (v4 & (v4 % (((v4 << ((2) & 63))) | 1)))
                                        _ = v22
                                    }
                                }
                            default:
                                switch (fi11 >> ((943) & 63)) {
                                case 127:
                                    v2 = ""
                                default:
                                    v4 -= v4
                                }
                            }
                            _ = fi21
                        }
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v24 := []int{}
                            var v25 int8 = (^(v4))
                            v26 := GenS1{F0: ("" + v2), F1: 2147483648, F2: -1}
                            v27 := GenS1{F0: v2, F1: 2147483647, F2: (v25 | int8(fi23))}
                            v28 := GenS1{F0: ((v2 + "a") + ((v2 + "") + v2)), F1: 42, F2: (-(v4))}
                            switch v2 {
                            case "a":
                                fmt.Println(v2, v2)
                            }
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    } else {
                        var v29 int = fi11
                        v30 := []int{v29, fi11}
                        for fi31 := 0; fi31 < 14; fi31++ {
                            if (true && (!false)) {
                                switch (v2 + v2) {
                                case "a":
                                    switch fi31 {
                                    case 392:
                                        var v32 float32 = (v5 - (1e10 / 1.5))
                                        _ = v32
                                    case 9223372036854775807:
                                        var v33 string = v2
                                        _ = v33
                                    default:
                                        var v34 uint64 = 704
                                        _ = v34
                                    }
                                }
                                fmt.Println(v29, v1, v6)
                            }
                            _ = fi31
                        }
                        _ = v29
                        _ = v30
                    }
                    fmt.Println(v5, fi11)
                    var v35 bool = ((v4 & v4) > (v4 & v4))
                    _ = v35
                }
            case "foo bar":
                if (180 == 0) {
                    v36 := GenS1{F0: v2, F1: 238, F2: v4}
                    v1 = (v1 >> ((665) & 63))
                    _ = v36
                } else {
                    if true {
                        var v37 float64 = v6
                        var v38 uint16 = (v1 + 42)
                        fmt.Println(v4)
                        _ = v37
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 14; fi39++ {
                        v40 := []int{fi39, 419, int(v1), fi11}
                        _ = fi39
                        _ = v40
                    }
                }
            default:
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v5 += v5
                default:
                    v41 := GenS2{F0: 0, F1: 378}
                    _ = v41
                }
            }
            v42 := GenS1{F0: "hello", F1: 831, F2: v4}
            v43 := GenS3{F0: false, F1: 1000}
            _ = v42
            _ = v43
        } else {
            switch ((v2 + "\n\t") + "") {
            case "":
                fmt.Println(v2)
            default:
                v44 := []int{42}
                _ = v44
            }
            v6 -= -2.25
            for fi45 := 0; fi45 < 4; fi45++ {
                var v46 int64 = int64(v1)
                _ = fi45
                _ = v46
            }
        }
        fmt.Println(fi11, v5)
        _ = fi11
    }
    if true {
        fmt.Println(v3, v4)
        v47 := []int{32768, 42, 32768, 42}
        v1 += (^(v1))
        _ = v47
    }
    v48 := GenS1{F0: v2, F1: 2147483647, F2: (v4 / ((1) | 1))}
    var v49 uint8 = 42
    if true {
        for fi50 := 0; fi50 < 3; fi50++ {
            v51 := GenS1{F0: (v2 + ("a" + ((("a" + "\n\t") + ((v2 + (v2 + (((v2 + (v2 + ("foo bar" + "foo bar"))) + v2) + v2))) + "hello")) + ((v2 + "foo bar") + v2)))), F1: 1, F2: 1}
            switch "a" {
            case "a":
                for fi52 := 0; fi52 < 6; fi52++ {
                    if (!false) {
                        v49 = v49
                        if false {
                            var v53 int64 = v3
                            if ((int8(fi50) >= 100) && (true || (((0.0 / (v5 / float32(float64(float32((v6 - 3.14159)))))) >= v5) && true))) {
                                v49 += uint8(v1)
                            }
                            _ = v53
                        } else {
                            v54 := []int{(fi50 << ((219) & 63))}
                            var v55 int64 = 1
                            if true {
                                if len(v54) > 0 {
                                    v54[0] = 853
                                }
                                switch v2 {
                                case "a":
                                    v49 = v49
                                case "foo bar":
                                    switch ("foo bar" + ((((v2 + v2) + v2) + (v2 + v2)) + ("" + v2))) {
                                    case "a":
                                        v55 = -508
                                    default:
                                        var v56 uint16 = 806
                                        _ = v56
                                    }
                                }
                                if false {
                                    if false {
                                        var v57 string = v2
                                        _ = v57
                                    } else {
                                        if false {
                                            v58 := GenS3{F0: ((v6 - 1e10) <= 1.5), F1: 0}
                                            _ = v58
                                        } else {
                                            _ = len(v54)
                                            var v59 string = (("\n\t" + "hello") + v2)
                                            if (42 >= -1) {
                                                if true {
                                                    var v60 uint32 = 651
                                                    fmt.Println(v6)
                                                    if (v5 > v5) {
                                                        v61 := GenS3{F0: false, F1: 1}
                                                        if false {
                                                            var v62 int32 = -263
                                                            _ = v62
                                                        } else {
                                                            if (!(true && (fi50 > (fi50 >> ((65535) & 63))))) {
                                                                if (v3 != (v55 >> ((0) & 63))) {
                                                                    v4 = -1
                                                                    switch fi52 {
                                                                    case 9223372036854775807:
                                                                        var v63 bool = false
                                                                        _ = v63
                                                                    case 128:
                                                                        v64 := GenS4{F0: false, F1: -406}
                                                                        _ = v64
                                                                    case 180:
                                                                        var v65 uint8 = (v49 - (v49 | 0))
                                                                        _ = v65
                                                                    default:
                                                                        v3 = 811
                                                                    }
                                                                }
                                                                var v66 uint = 4294967295
                                                                fmt.Println(fi50, v2)
                                                                if ((v4 - v4) != (v4 << (((v66 ^ (v66 & (v66 % ((v66) | 1))))) & 63))) {
                                                                    var v67 uint8 = 0
                                                                    v4 = int8(v55)
                                                                    if ((!(uint16(v67) < (v1 ^ 76))) || false) {
                                                                        var v68 uint = (v66 % ((v66) | 1))
                                                                        _ = v68
                                                                    } else {
                                                                        var v69 float64 = 0.1
                                                                        v70 := GenS3{F0: false, F1: -1}
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                    var v71 uint64 = uint64(v66)
                                                                    _ = v67
                                                                    _ = v71
                                                                } else {
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        var v72 float32 = v5
                                                                        _ = v72
                                                                    }
                                                                    var v73 string = ("" + "\n\t")
                                                                    _ = v73
                                                                }
                                                                _ = v66
                                                            } else {
                                                                v2 = ""
                                                                fmt.Println(fi50, v4, v3)
                                                                fmt.Println(v60)
                                                            }
                                                            switch (fi50 / (((fi50 ^ fi50)) | 1)) {
                                                            case 0:
                                                                var v74 uint = 2
                                                                _ = v74
                                                            case 42:
                                                                if len(v54) >= 2 {
                                                                    v54 = v54[1:len(v54)]
                                                                }
                                                            default:
                                                                _ = len(v54)
                                                            }
                                                            v54 = append(v54, 17)
                                                        }
                                                        _ = v61
                                                    }
                                                    _ = v60
                                                } else {
                                                    switch v59 {
                                                    case "a":
                                                        fmt.Println(fi50, v49)
                                                    case "foo bar":
                                                        var v75 uint8 = (v49 >> ((662) & 63))
                                                        var v76 int32 = -361
                                                        var v77 float64 = (v6 * 0.0)
                                                        _ = v75
                                                        _ = v76
                                                        _ = v77
                                                    default:
                                                        var v78 int = (fi52 % (((fi52 + fi50)) | 1))
                                                        _ = v78
                                                    }
                                                    var v79 uint16 = v1
                                                    _ = v79
                                                }
                                                fmt.Println(v2, v5, v1)
                                                v6 = v6
                                            }
                                            _ = v59
                                        }
                                        v80 := GenS3{F0: (float32(v6) == v5), F1: 1000}
                                        _ = v80
                                    }
                                    var v81 uint16 = (v1 % ((v1) | 1))
                                    var v82 int64 = (v55 - (v55 / ((v55) | 1)))
                                    _ = v81
                                    _ = v82
                                }
                                _ = len(v54)
                            } else {
                                var v83 int = (fi50 ^ (-(fi52)))
                                _ = len(v54)
                                v84 := GenS3{F0: (128 >= v49), F1: -440}
                                _ = v83
                                _ = v84
                            }
                            switch v2 {
                            case "":
                                v4 -= 127
                            case "\n\t":
                                _ = len(v54)
                            default:
                                v54 = append(v54, (fi50 << ((419) & 63)))
                            }
                            fmt.Println(fi50, v6)
                            _ = v54
                            _ = v55
                        }
                    }
                    if ((v4 / ((100) | 1)) == v4) {
                        v85 := GenS5{F0: 42, F1: 1}
                        if (!true) {
                            v86 := []int{}
                            _ = v86
                        }
                        _ = v85
                    }
                    v4 = v4
                    _ = fi52
                }
            case "foo bar":
                switch 320 {
                case -52:
                    v4 = 42
                case 755:
                    switch v2 {
                    case "hello":
                        if true {
                            var v87 int64 = (^(v3))
                            if (-2147483648 < 0) {
                                var v88 float32 = v5
                                _ = v88
                            } else {
                                v89 := GenS4{F0: ((0 < 2) || false), F1: -2147483648}
                                fmt.Println(v1, v5)
                                if (!false) {
                                    for fi90 := 0; fi90 < 17; fi90++ {
                                        switch -20 {
                                        case 589:
                                            if (!false) {
                                                var v91 float64 = v6
                                                v92 := []int{}
                                                if false {
                                                    var v93 float64 = float64(3.14159)
                                                    var v94 uint32 = 1
                                                    v95 := GenS5{F0: v94, F1: 0}
                                                    _ = v93
                                                    _ = v94
                                                    _ = v95
                                                }
                                                _ = v91
                                                _ = v92
                                            }
                                        case -2147483648:
                                            var v96 int64 = genFn0((^(v1)))
                                            _ = v96
                                        default:
                                            v97 := []int{int(v49), fi50, fi50}
                                            v98 := GenS3{F0: (((v4 <= -128) || (fi90 == (fi50 >> ((1) & 63)))) || false), F1: 1000}
                                            fmt.Println(v49, v6, v1)
                                            _ = v97
                                            _ = v98
                                        }
                                        fmt.Println(v6, v4)
                                        fmt.Println(v87, v2)
                                        _ = fi90
                                    }
                                    v99 := []int{(fi50 / ((int(v4)) | 1)), (fi50 % ((fi50) | 1))}
                                    _ = v99
                                } else {
                                    var v100 uint64 = 4294967295
                                    v101 := []int{}
                                    _ = v100
                                    _ = v101
                                }
                                _ = v89
                            }
                            _ = v87
                        }
                    case "a":
                        v102 := GenS6{F0: 874, F1: v3}
                        _ = v102
                    default:
                        var v103 int16 = 1
                        _ = v103
                    }
                default:
                    for fi104 := 0; fi104 < 4; fi104++ {
                        v2 = v2
                        _ = fi104
                    }
                }
            case "hello":
                var v105 bool = (-678 == v3)
                _ = v105
            }
            switch (v2 + v2) {
            case "\n\t":
                for fi106 := 0; fi106 < 13; fi106++ {
                    var v107 bool = ((1 > 42) || (false && true))
                    _ = fi106
                    _ = v107
                }
            default:
                fmt.Println(v2, v2)
            }
            _ = fi50
            _ = v51
        }
        fmt.Println(v6)
        v6 = 1e10
    } else {
        fmt.Println(v1, v2, v2)
        switch v2 {
        case "hello":
            v108 := GenS5{F0: 518, F1: int16(v49)}
            _ = v108
        case "\n\t":
            var v109 uint64 = 4294967295
            _ = v109
        case "":
            var v110 uint = 943
            _ = v110
        default:
            v111 := GenS2{F0: v3, F1: 1}
            _ = v111
        }
        v112 := GenS7{F0: v1, F1: uint32(v49), F2: true}
        _ = v112
    }
    if (!(-2147483648 == 42)) {
        v113 := GenS4{F0: false, F1: 0}
        var v114 int32 = 136
        for fi115 := 0; fi115 < 20; fi115++ {
            fmt.Println(v1, fi115, v49)
            _ = fi115
        }
        _ = v113
        _ = v114
    }
    v116 := GenS8{F0: 127, F1: -191}
    var v117 bool = true
    var v118 int32 = -518
    var v119 float32 = v5
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v116
    _ = v117
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
