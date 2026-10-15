package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 string; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int; F1 uint16; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 uint32; F1 int16; F2 int64 }
type GenS11 struct { F0 uint64; F1 string; F2 int16 }
type GenS12 struct { F0 uint; F1 bool }
func genFn0() int8 {
    var v7 float32 = 0.1
    var v8 uint8 = 255
    v9 := GenS1{F0: 42}
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int, p1 int64, p2 int) string {
    if ((0 > 2) || (true && ((!(1000 < 4294967295)) && false))) {
        v10 := []int{(p0 ^ p0), (^(p0)), 1, p0}
        var v11 int = (p0 - p2)
        switch p2 {
        case -142:
            var v12 float32 = 1.5
            _ = v12
        case 42:
            if (!false) {
                var v13 bool = false
                if len(v10) > 0 {
                    v10[0] = -1
                }
                v14 := GenS1{F0: p1}
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 18; fi15++ {
                    if len(v10) > 0 {
                        v10[0] = p0
                    }
                    if len(v10) > 0 {
                        v10[0] = (p2 << ((737) & 63))
                    }
                    var v16 int16 = -650
                    _ = fi15
                    _ = v16
                }
            }
        default:
            var v17 uint16 = 65535
            _ = v17
        }
        _ = v10
        _ = v11
    } else {
        v18 := []int{p0, p0, 4294967295}
        var v19 int16 = 32767
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        if ((42 != -32768) || false) {
            v21 := GenS2{F0: false, F1: "foo bar", F2: 740}
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 int32 = -2147483648
                p1 -= (p1 * p1)
                _ = fi22
                _ = v23
            }
            v24 := GenS1{F0: -608}
            _ = v21
            _ = v24
        }
        p2 = (p2 % (((p0 + fi20)) | 1))
        _ = fi20
    }
    var v25 int16 = -77
    var v26 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return "\n\t"
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 813
    var v3 string = "a"
    var v4 int32 = -2147483648
    var v5 bool = true
    var v6 uint64 = 0
    if v5 {
        switch 4294967295 {
        case 32768:
            v6 = v6
        default:
            var v27 uint8 = 42
            _ = v27
        }
        for fi28 := 0; fi28 < 10; fi28++ {
            v1[v3] = (fi28 << (((v2 << (((v2 + v2)) & 63))) & 63))
            switch (fi28 % ((32768) | 1)) {
            case 128:
                if (!((v6 & 694) > (v6 / (((v6 >> (((v2 | v2)) & 63))) | 1)))) {
                    v1[v3] = (fi28 & (^(fi28)))
                    if true {
                        v29 := GenS1{F0: -1}
                        var v30 int64 = 697
                        _ = v29
                        _ = v30
                    }
                    var v31 uint32 = 42
                    _ = v31
                }
            }
            for fi32 := 0; fi32 < 7; fi32++ {
                v4 += (v4 >> (((v2 * (^(v2)))) & 63))
                v33 := GenS3{F0: 32767, F1: 128}
                var v34 uint8 = 2
                var v35 float32 = 1.5
                v2 = 255
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            _ = fi28
        }
    }
    v4 -= -372
    if (-9223372036854775808 >= 9223372036854775807) {
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 uint = v2
            if (32767 > 1000) {
                if (true && ((v6 >= v6) && v5)) {
                    v4 = v4
                    v38 := []int{fi36, 9223372036854775807}
                    _ = v38
                } else {
                    var v39 uint8 = 157
                    _ = v39
                }
                var v40 int32 = 859
                for fi41 := 0; fi41 < 17; fi41++ {
                    var v42 int32 = v40
                    if (!v5) {
                        v4 = 1
                        v4 += (-(v4))
                    } else {
                        switch (fi36 - -9223372036854775808) {
                        case -9223372036854775808:
                            v6 -= v6
                        case 892:
                            var v43 string = v3
                            _ = v43
                        }
                    }
                    _ = fi41
                    _ = v42
                }
                _ = v40
            }
            _ = fi36
            _ = v37
        }
        var v44 int8 = 1
        _ = v44
    } else {
        var v45 float64 = (-2.25 * -2.25)
        _ = v45
    }
    fmt.Println(v6, v5, v3)
    if v5 {
        if v5 {
            if (!v5) {
                for fi46 := 0; fi46 < 18; fi46++ {
                    var v47 float64 = 3.14159
                    _ = fi46
                    _ = v47
                }
                if (128 > 939) {
                    switch "" {
                    case "a":
                        var v48 float32 = (1e10 * 0.0)
                        _ = v48
                    case "hello":
                        v49 := []int{521, 4294967295, -533}
                        _ = v49
                    default:
                        if true {
                            v5 = (65535 == 42)
                            v50 := []int{851, -2147483648, 32768, 127}
                            _ = v50
                        } else {
                            v51 := GenS4{F0: v6}
                            _ = v51
                        }
                    }
                    var v52 uint32 = 2147483647
                    var v53 int16 = -32768
                    _ = v52
                    _ = v53
                }
                v54 := GenS1{F0: -182}
                _ = v54
            } else {
                if v5 {
                    for fi55 := 0; fi55 < 16; fi55++ {
                        v2 = (v2 >> ((v2) & 63))
                        v56 := GenS5{F0: 169, F1: (((^(v6)) > v6) && false)}
                        var v57 float64 = 1e10
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                }
            }
            switch "\n\t" {
            case "hello":
                if (569 < 2) {
                    if ((v2 > v2) && (!true)) {
                        switch -2147483648 {
                        case -9223372036854775808:
                            v6 = 2
                        default:
                            v2 += v2
                        }
                        for fi58 := 0; fi58 < 15; fi58++ {
                            var v59 uint = 483
                            v4 += 182
                            _ = fi58
                            _ = v59
                        }
                    } else {
                        for fi60 := 0; fi60 < 7; fi60++ {
                            if (6 > 204) {
                                switch ((v3 + v3) + v3) {
                                case "hello":
                                    _ = len(v1)
                                case "\n\t":
                                    var v61 int8 = -1
                                    _ = v61
                                case "":
                                    var v62 uint = (v2 % ((v2) | 1))
                                    _ = v62
                                default:
                                    v63 := GenS1{F0: -9223372036854775808}
                                    _ = v63
                                }
                                v5 = true
                            }
                            var v64 int8 = genFn0()
                            v64 = (v64 ^ 127)
                            _ = fi60
                            _ = v64
                        }
                        v65 := GenS1{F0: 9223372036854775807}
                        v2 -= v2
                        _ = v65
                    }
                    fmt.Println(v6)
                    v66 := v1["foo bar"]
                    _ = v66
                } else {
                    if false {
                        var v67 uint64 = 2
                        _ = v67
                    } else {
                        if (v5 && v5) {
                            switch v3 {
                            case "":
                                if true {
                                    v68 := []int{1, -987}
                                    _ = v68
                                }
                            default:
                                v69 := GenS6{F0: -9223372036854775808, F1: 2, F2: genFn0()}
                                _ = v69
                            }
                            v70 := GenS6{F0: -2147483648, F1: 14, F2: 42}
                            _ = v70
                        } else {
                            fmt.Println(v3, v4)
                        }
                    }
                    v6 = (v6 / ((v6) | 1))
                    for fi71 := 0; fi71 < 5; fi71++ {
                        switch (fi71 / (((fi71 | fi71)) | 1)) {
                        case 664:
                            v2 += (v2 | (v2 + (v2 % (((^(v2))) | 1))))
                        case -945:
                            _ = len(v1)
                        }
                        v3 = ""
                        _ = fi71
                    }
                }
            case "\n\t":
                for fi72 := 0; fi72 < 4; fi72++ {
                    switch v3 {
                    case "hello":
                        fmt.Println(v5)
                    case "foo bar":
                        var v73 int32 = -2147483648
                        _ = v73
                    case "":
                        for fi74 := 0; fi74 < 13; fi74++ {
                            if (!true) {
                                if (42 > -14) {
                                    if false {
                                        v6 -= (v6 / (((^(v6))) | 1))
                                        if v5 {
                                            if (v5 || v5) {
                                                if (1 >= 0) {
                                                    var v75 string = "\n\t"
                                                    fmt.Println(fi72)
                                                    switch (("" + ((("foo bar" + "") + "hello") + ((v75 + v3) + (v3 + v3)))) + "\n\t") {
                                                    case "\n\t":
                                                        v76 := []int{-2147483648}
                                                        _ = v76
                                                    case "a":
                                                        if ((9223372036854775807 == 9223372036854775807) && ((v75 > ("a" + v3)) && false)) {
                                                            v77 := []int{}
                                                            v75 = v3
                                                            v6 = 2
                                                            _ = v77
                                                        }
                                                    case "hello":
                                                        v78 := GenS1{F0: 1}
                                                        _ = v78
                                                    default:
                                                        var v79 float32 = ((1.5 * 3.14159) + float32(3.14159))
                                                        _ = v79
                                                    }
                                                    _ = v75
                                                }
                                                v80 := GenS7{F0: 1000}
                                                _ = v80
                                            } else {
                                                v81 := GenS5{F0: 0, F1: ((v3 + ("a" + (v3 + v3))) <= (v3 + "\n\t"))}
                                                v6 += (v6 + v6)
                                                v4 += (v4 + (v4 | (v4 - v4)))
                                                _ = v81
                                            }
                                        }
                                    }
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        switch v3 {
                                        case "hello":
                                            switch v3 {
                                            case "":
                                                v82 := GenS8{F0: (v4 ^ 42)}
                                                _ = v82
                                            }
                                        case "a":
                                            var v83 uint8 = uint8(v2)
                                            _ = v83
                                        case "\n\t":
                                            switch ((v3 + ((v3 + (v3 + ("foo bar" + v3))) + "")) + v3) {
                                            case "":
                                                v84 := GenS9{F0: (v6 & 224), F1: (-(fi72))}
                                                _ = v84
                                            case "foo bar":
                                                v6 -= v6
                                            default:
                                                var v85 uint16 = 32767
                                                _ = v85
                                            }
                                        }
                                    case "hello":
                                        if true {
                                            switch (v3 + "foo bar") {
                                            case "":
                                                v86 := GenS9{F0: v6, F1: fi72}
                                                _ = v86
                                            case "a":
                                                v87 := v1[v3]
                                                _ = v87
                                            }
                                            v5 = (true || ((uint16(v6) >= 0) || ((v2 - 1) >= v2)))
                                            if v5 {
                                                switch 61 {
                                                case -1:
                                                    v88 := GenS10{F0: 2147483647, F1: -424, F2: 0}
                                                    _ = v88
                                                default:
                                                    v89 := []int{fi72, fi72}
                                                    _ = v89
                                                }
                                                v4 -= (v4 + v4)
                                            }
                                            fmt.Println(v6, v5, v2)
                                            if (1e10 < float32((3.14159 * (float64((3.14159 - 1e10)) * 1.5)))) {
                                                v90 := GenS11{F0: (^(v6)), F1: "\n\t", F2: 32767}
                                                _ = v90
                                            } else {
                                                v4 -= 940
                                            }
                                        } else {
                                            v1["foo bar"] = (fi74 & fi74)
                                            if v5 {
                                                v1[((v3 + v3) + ("" + v3))] = fi74
                                                var v91 int32 = v4
                                                _ = v91
                                            }
                                        }
                                    case "":
                                        _ = len(v1)
                                    default:
                                        v92 := []int{fi72, (fi74 * fi74), fi74, fi74}
                                        _ = v92
                                    }
                                    v4 = v4
                                    if true {
                                        v93 := GenS2{F0: (2 >= 255), F1: v3, F2: 42}
                                        v94 := v1[((v3 + v3) + v3)]
                                        var v95 string = (v3 + ("foo bar" + v3))
                                        _ = len(v1)
                                        fmt.Println(fi74, v3, v5)
                                        _ = v93
                                        _ = v94
                                        _ = v95
                                    } else {
                                        v4 += (v4 % (((v4 | (v4 << ((v2) & 63)))) | 1))
                                        v5 = v5
                                    }
                                }
                            }
                            _ = fi74
                        }
                    default:
                        v96 := GenS9{F0: v6, F1: 526}
                        _ = v96
                    }
                    var v97 uint64 = (v6 >> ((1) & 63))
                    var v98 int = 128
                    _ = fi72
                    _ = v97
                    _ = v98
                }
            default:
                var v99 float64 = -2.25
                _ = v99
            }
            fmt.Println(v5)
        } else {
            for fi100 := 0; fi100 < 11; fi100++ {
                if ((v5 && v5) && (0 < 255)) {
                    var v101 int8 = 1
                    for fi102 := 0; fi102 < 16; fi102++ {
                        v103 := GenS4{F0: v6}
                        fmt.Println(v5, v2)
                        v104 := GenS3{F0: 65535, F1: 255}
                        if true {
                            var v105 float64 = 1.5
                            switch (fi100 | (fi100 / ((fi102) | 1))) {
                            case 1:
                                var v106 uint16 = uint16(v6)
                                _ = v106
                            default:
                                v107 := []int{fi102, (fi102 << (((v2 % (((v2 >> (((v2 / (((v2 << ((619) & 63))) | 1))) & 63))) | 1))) & 63)), (fi102 - (fi102 ^ (fi102 / (((fi102 / ((fi100) | 1))) | 1)))), (fi100 << ((v2) & 63))}
                                _ = v107
                            }
                            var v108 int = fi102
                            _ = v105
                            _ = v108
                        }
                        _ = fi102
                        _ = v103
                        _ = v104
                    }
                    v2 = (v2 << (((^(v2))) & 63))
                    _ = v101
                }
                v109 := []int{32767, 42, fi100, (fi100 / (((fi100 * fi100)) | 1))}
                _ = fi100
                _ = v109
            }
            if (v4 <= -22) {
                switch ("hello" + "\n\t") {
                case "":
                    var v110 uint64 = v6
                    _ = v110
                default:
                    for fi111 := 0; fi111 < 16; fi111++ {
                        v3 = ("hello" + v3)
                        v112 := GenS7{F0: 971}
                        _ = fi111
                        _ = v112
                    }
                }
            }
            v2 = (^(v2))
        }
        if v5 {
            if ("foo bar" >= v3) {
                v113 := GenS12{F0: (v2 | 18446744073709551615), F1: v5}
                _ = v113
            } else {
                v3 = v3
                _ = len(v1)
            }
        }
    }
    v114 := v1[v3]
    v5 = ((-1 < 1) || (!(v2 <= v2)))
    var v115 int8 = -43
    v6 = (v6 | (v6 + (v6 + v6)))
    var v116 uint8 = 255
    var v117 int64 = -106
    if v5 {
        v118 := GenS6{F0: v114, F1: 1, F2: genFn0()}
        v6 = (v6 << ((v2) & 63))
        var v119 float64 = 0.0
        _ = v118
        _ = v119
    }
    var v120 uint16 = 1
    var v121 uint32 = 1
    v122 := v1[v3]
    v123 := []int{(v114 - v122), -427}
    for fi124 := 0; fi124 < 16; fi124++ {
        var v125 float64 = 1.5
        _ = fi124
        _ = v125
    }
    v120 = (v120 ^ (^(v120)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("slicelen v123:", len(v123), v123)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v115
    _ = v116
    _ = v117
    _ = v120
    _ = v121
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
