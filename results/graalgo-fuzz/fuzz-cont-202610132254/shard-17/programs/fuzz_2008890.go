package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int64; F2 int64 }
type GenS3 struct { F0 bool; F1 int; F2 uint16 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int8; F1 uint; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string; F1 int8; F2 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 uint8 }
func genFn0(p0 int, p1 int64) uint32 {
    fmt.Println(p1, p0)
    var v4 uint8 = 255
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 12; fi6++ {
        v4 = v4
        var v7 uint8 = v4
        var v8 int64 = p1
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 224 {
    case 776:
        var v9 int64 = p1
        _ = v9
    default:
        if (0.0 == 1e10) {
            v10 := GenS2{F0: -9223372036854775808, F1: (p1 | -901), F2: (p1 >> ((833) & 63))}
            var v11 int16 = 504
            _ = v10
            _ = v11
        } else {
            var v12 int32 = 2147483647
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1() int32 {
    v13 := GenS1{F0: 918}
    if true {
        if true {
            if (true && false) {
                if (-32768 >= 32767) {
                    v14 := GenS2{F0: 410, F1: 1, F2: -9223372036854775808}
                    var v15 uint = 42
                    v16 := []int{1, int(v15), 2147483647}
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    switch ("" + "foo bar") {
                    case "foo bar":
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if (false || (0 >= 0)) {
                                v18 := []int{fi17, 9223372036854775807, fi17}
                                _ = v18
                            }
                            _ = fi17
                        }
                    default:
                        v19 := GenS3{F0: (!true), F1: 32768, F2: 32767}
                        _ = v19
                    }
                }
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(fi20)
                    _ = fi20
                }
                var v21 int32 = -1
                _ = v21
            } else {
                var v22 float32 = (0.1 / 0.1)
                v22 -= (float32(float64(v22)) * float32(0.0))
                fmt.Println(v22)
                _ = v22
            }
            switch -1 {
            case 1:
                var v23 uint = 255
                _ = v23
            default:
                var v24 float32 = (1.5 * 1.5)
                _ = v24
            }
            v25 := []int{127, 1}
            _ = v25
        } else {
            if false {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := GenS4{F0: true, F1: (^(fi26))}
                    var v28 float64 = (1.5 / 1.5)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                var v29 int = 32768
                _ = v29
            }
        }
        var v30 bool = true
        for fi31 := 0; fi31 < 16; fi31++ {
            v30 = v30
            if (2 >= 18446744073709551615) {
                for fi32 := 0; fi32 < 11; fi32++ {
                    v30 = true
                    _ = fi32
                }
            } else {
                v30 = true
                v33 := []int{}
                _ = v33
            }
            _ = fi31
        }
        _ = v30
    }
    var v34 float32 = (1e10 / -2.25)
    if (1000 != 42) {
        v35 := GenS5{F0: 100, F1: 255, F2: -32768}
        var v36 uint = 827
        _ = v35
        _ = v36
    } else {
        if true {
            v37 := []int{1, 883, 490}
            if len(v37) >= 2 {
                v37 = v37[1:len(v37)]
            }
            _ = v37
        } else {
            if false {
                var v38 int = 128
                v39 := []int{}
                v38 = 127
                _ = v38
                _ = v39
            }
        }
        v40 := []int{}
        _ = v40
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        v34 += v34
        var v42 uint = 18446744073709551615
        switch 128 {
        case 32767:
            v43 := []int{fi41, fi41, fi41, 38}
            _ = v43
        case -690:
            v44 := []int{}
            _ = v44
        default:
            if false {
                if true {
                    for fi45 := 0; fi45 < 9; fi45++ {
                        if false {
                            var v46 uint16 = 316
                            var v47 int32 = -989
                            _ = v46
                            _ = v47
                        }
                        if (true && (genFn0((fi45 + fi45), int64(v42)) < 4294967295)) {
                            var v48 uint = v42
                            _ = v48
                        } else {
                            v34 += v34
                            var v49 uint = v42
                            var v50 uint64 = 0
                            _ = v49
                            _ = v50
                        }
                        var v51 string = ""
                        _ = fi45
                        _ = v51
                    }
                    var v52 int8 = -128
                    v34 = (v34 * ((v34 + (((v34 * float32(1e10)) * 1.5) - (v34 - -2.25))) / ((((v34 - float32(float64(float32((0.1 * 1.5))))) / v34) / float32((0.1 * float64(1e10)))) + 0.1)))
                    _ = v52
                }
                v53 := GenS6{F0: (fi41 >> ((65535) & 63))}
                _ = v53
            }
        }
        _ = fi41
        _ = v42
    }
    _ = v13
    _ = v34
    return -920
}

func main() {
    var v1 int8 = -1
    var v2 string = ""
    var v3 float32 = ((1e10 / (((1e10 - -2.25) / (0.1 + 0.0)) / 0.1)) + 3.14159)
    if true {
        v3 += v3
        var v54 uint32 = 329
        _ = v54
    }
    v55 := GenS4{F0: (v1 == 1), F1: 4294967295}
    for fi56 := 0; fi56 < 8; fi56++ {
        if true {
            switch (v2 + v2) {
            case "hello":
                v57 := []int{(fi56 ^ fi56), fi56}
                _ = v57
            }
            var v58 string = v2
            var v59 int64 = -9223372036854775808
            _ = v58
            _ = v59
        }
        _ = fi56
    }
    v2 = ("\n\t" + (v2 + v2))
    if (42 != 85) {
        for fi60 := 0; fi60 < 16; fi60++ {
            var v61 uint = 4294967295
            var v62 int32 = 42
            switch "a" {
            case "\n\t":
                if false {
                    for fi63 := 0; fi63 < 20; fi63++ {
                        if ((((v62 & 42) < v62) && true) || (v3 > 0.1)) {
                            var v64 uint16 = 255
                            if true {
                                switch 32768 {
                                case 4294967295:
                                    v61 -= (v61 | v61)
                                case 943:
                                    switch (v2 + "hello") {
                                    case "hello":
                                        if false {
                                            v65 := GenS2{F0: fi63, F1: 42, F2: -1}
                                            v66 := GenS4{F0: (42 == 2147483647), F1: 2147483647}
                                            v67 := GenS7{F0: 1, F1: (v64 + v64), F2: 255}
                                            _ = v65
                                            _ = v66
                                            _ = v67
                                        }
                                    default:
                                        v68 := GenS8{F0: 0}
                                        _ = v68
                                    }
                                case 128:
                                    v2 = "\n\t"
                                }
                                fmt.Println(v2)
                                fmt.Println(fi63, v3)
                            }
                            var v69 uint8 = uint8(v64)
                            switch (("foo bar" + v2) + ((v2 + v2) + ((v2 + v2) + v2))) {
                            case "a":
                                if false {
                                    var v70 bool = ((v1 <= (v1 * (v1 / ((100) | 1)))) && (1000 <= int16(v64)))
                                    switch "hello" {
                                    case "":
                                        var v71 int = (fi63 & fi63)
                                        _ = v71
                                    case "foo bar":
                                        v61 = 65535
                                    case "hello":
                                        v72 := GenS7{F0: 1, F1: v64, F2: (v64 >> (((v61 & (v61 % (((v61 ^ (v61 ^ (^(v61))))) | 1)))) & 63))}
                                        _ = v72
                                    default:
                                        if (!true) {
                                            if v70 {
                                                if true {
                                                    var v73 uint = (v61 & v61)
                                                    _ = v73
                                                }
                                                switch fi63 {
                                                case -144:
                                                    if v70 {
                                                        if false {
                                                            if (v1 == 1) {
                                                                fmt.Println(fi63, v1)
                                                                v62 += (v62 + (v62 + v62))
                                                            }
                                                            if (3.14159 >= (3.14159 - 0.0)) {
                                                                var v74 int8 = (v1 - (v1 % ((1) | 1)))
                                                                v70 = v70
                                                                v1 -= v74
                                                                _ = v74
                                                            } else {
                                                                var v75 int64 = 1
                                                                switch -816 {
                                                                case 4294967295:
                                                                    v61 += (v61 % ((1) | 1))
                                                                case 0:
                                                                    var v76 uint16 = (v64 | 253)
                                                                    _ = v76
                                                                default:
                                                                    var v77 int = fi63
                                                                    _ = v77
                                                                }
                                                                _ = v75
                                                            }
                                                        }
                                                    }
                                                case -32:
                                                    var v78 float64 = 1.5
                                                    _ = v78
                                                case 4294967295:
                                                    if (-9223372036854775808 != -383) {
                                                        v79 := []int{42, fi63, fi60, fi60}
                                                        _ = v79
                                                    }
                                                default:
                                                    var v80 uint = (v61 ^ v61)
                                                    _ = v80
                                                }
                                                if v70 {
                                                    if v70 {
                                                        v1 = v1
                                                        var v81 float64 = (float64(float32(((1e10 * (0.1 - float64((v3 - (float32(float64(((float32(-2.25) + float32(0.1)) * 0.0))) / 0.1))))) * (3.14159 - float64((v3 / -2.25)))))) * 1.5)
                                                        _ = v81
                                                    }
                                                    v2 = "foo bar"
                                                } else {
                                                    switch fi63 {
                                                    case 370:
                                                        v82 := []int{fi60, (fi63 / (((fi63 * (fi63 * (^(fi63))))) | 1)), fi60}
                                                        _ = v82
                                                    case 0:
                                                        v1 += (v1 / (((-(v1))) | 1))
                                                    }
                                                    v83 := []int{32768, 512, (fi60 >> (((v61 >> ((v61) & 63))) & 63)), 4294967295}
                                                    _ = v83
                                                }
                                                if ((v69 >> ((934) & 63)) >= v69) {
                                                    fmt.Println(v1, v70, fi63)
                                                }
                                                switch fi63 {
                                                case 786:
                                                    v2 = "a"
                                                default:
                                                    switch ((v2 + "\n\t") + v2) {
                                                    case "\n\t":
                                                        v1 = v1
                                                    case "hello":
                                                        fmt.Println(v64)
                                                        fmt.Println(fi60, v64)
                                                        v3 += v3
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v70)
                                                var v84 uint = v61
                                                if v70 {
                                                    switch "\n\t" {
                                                    case "":
                                                        v85 := GenS2{F0: fi60, F1: 9223372036854775807, F2: 0}
                                                        _ = v85
                                                    case "hello":
                                                        fmt.Println(v2, v64, fi63)
                                                    }
                                                    var v86 bool = v70
                                                    var v87 uint = (v61 % (((v84 - 2)) | 1))
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    if v70 {
                                                        v88 := GenS8{F0: 42}
                                                        _ = v88
                                                    }
                                                    v84 += v61
                                                    switch 32767 {
                                                    case 1:
                                                        v69 -= v69
                                                    case 2147483647:
                                                        v89 := GenS3{F0: v70, F1: fi63, F2: (v64 >> ((198) & 63))}
                                                        _ = v89
                                                    case -1:
                                                        var v90 uint8 = 153
                                                        _ = v90
                                                    default:
                                                        var v91 uint = uint(v69)
                                                        v92 := []int{}
                                                        var v93 uint16 = 1
                                                        v94 := GenS9{F0: v2, F1: v1, F2: 2147483648}
                                                        switch (fi63 | int(v62)) {
                                                        case -1:
                                                            fmt.Println(v84, v61)
                                                        case 4294967295:
                                                            if (v70 && v70) {
                                                                v95 := GenS8{F0: 42}
                                                                var v96 uint8 = (^(v69))
                                                                if len(v92) > 0 {
                                                                    v92[0] = (fi63 / ((fi60) | 1))
                                                                }
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                            fmt.Println(v93, v93, fi60)
                                                            fmt.Println(v91, fi63)
                                                            var v97 bool = (v3 != v3)
                                                            _ = v97
                                                        case 127:
                                                            if true {
                                                                if v70 {
                                                                    switch (fi63 >> (((^(v91))) & 63)) {
                                                                    case -1:
                                                                        if v70 {
                                                                            fmt.Println(v84)
                                                                            v84 += (v61 >> ((994) & 63))
                                                                        }
                                                                        switch ("foo bar" + ((((v2 + (v2 + ((v2 + "\n\t") + "\n\t"))) + (v2 + "a")) + "foo bar") + v2)) {
                                                                        case "a":
                                                                            v98 := GenS10{F0: 1}
                                                                            _ = v98
                                                                        default:
                                                                            if v70 {
                                                                                v70 = ((((v70 || (!v70)) || v70) && false) && v70)
                                                                            }
                                                                        }
                                                                        var v99 int16 = 1
                                                                        _ = v99
                                                                    default:
                                                                        v93 += (v64 % ((1) | 1))
                                                                        v92 = append(v92, fi63)
                                                                    }
                                                                } else {
                                                                    var v100 uint32 = 532
                                                                    var v101 uint64 = 42
                                                                    _ = v100
                                                                    _ = v101
                                                                }
                                                                v91 = v91
                                                            } else {
                                                                switch ("foo bar" + "a") {
                                                                case "hello":
                                                                    var v102 int64 = 42
                                                                    _ = v102
                                                                default:
                                                                    if (v70 && true) {
                                                                        if false {
                                                                            v69 = v69
                                                                            v92 = append(v92, (fi63 * (fi60 >> (((v84 * v84)) & 63))))
                                                                            v3 -= v3
                                                                            var v103 int = (fi60 << ((159) & 63))
                                                                            _ = v103
                                                                        } else {
                                                                            v2 = v2
                                                                        }
                                                                        v70 = v70
                                                                        var v104 int = 32767
                                                                        _ = v104
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v91
                                                        _ = v92
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                }
                                                _ = v84
                                            }
                                        }
                                    }
                                    var v105 int8 = v1
                                    _ = v105
                                    _ = v70
                                }
                            case "\n\t":
                                switch fi60 {
                                case 127:
                                    v106 := GenS7{F0: -327, F1: v64, F2: (v64 ^ (v64 % ((v64) | 1)))}
                                    fmt.Println(v2, fi60, v61)
                                    switch fi60 {
                                    case 946:
                                        v107 := []int{(fi60 << ((v61) & 63))}
                                        _ = v107
                                    case 128:
                                        var v108 float64 = (0.1 + 1.5)
                                        _ = v108
                                    default:
                                        v109 := GenS11{F0: 2147483647, F1: 119}
                                        _ = v109
                                    }
                                    var v110 uint64 = 42
                                    _ = v106
                                    _ = v110
                                }
                            case "hello":
                                v111 := GenS6{F0: -2147483648}
                                v112 := GenS4{F0: ((v61 - (^(v61))) > (v61 ^ (v61 >> ((v61) & 63)))), F1: -2147483648}
                                _ = v111
                                _ = v112
                            }
                            fmt.Println(v61)
                            _ = v64
                            _ = v69
                        } else {
                            v113 := GenS12{F0: 0}
                            fmt.Println(fi60, v61, v3)
                            _ = v113
                        }
                        _ = fi63
                    }
                    var v114 int32 = v62
                    v115 := GenS1{F0: 2}
                    _ = v114
                    _ = v115
                }
            case "hello":
                switch v2 {
                case "hello":
                    v116 := []int{4294967295, fi60, fi60, 731}
                    _ = v116
                case "foo bar":
                    for fi117 := 0; fi117 < 4; fi117++ {
                        fmt.Println(v1, v61, v62)
                        _ = fi117
                    }
                case "a":
                    v118 := GenS1{F0: 105}
                    _ = v118
                }
            case "a":
                fmt.Println(fi60, v2)
            default:
                v61 -= v61
            }
            _ = fi60
            _ = v61
            _ = v62
        }
        switch "hello" {
        case "\n\t":
            fmt.Println(v2)
        case "a":
            if false {
                var v119 bool = ((false || (v1 == -1)) || (1 <= 0))
                var v120 bool = true
                v119 = ((false || true) || ((float32(((1.5 * (((((1e10 - (3.14159 / 3.14159)) + 0.0) / float64(((0.0 * v3) * 3.14159))) + -2.25) / -2.25)) + 0.0)) * 1.5) <= 0.0))
                _ = v119
                _ = v120
            } else {
                if ((^(v1)) <= 0) {
                    if false {
                        v121 := []int{0, -1, -9223372036854775808}
                        v3 -= 1e10
                        _ = len(v121)
                        _ = v121
                    }
                }
                switch 42 {
                case -713:
                    v122 := []int{0, 42}
                    _ = v122
                default:
                    v123 := []int{127, -2147483648}
                    _ = v123
                }
            }
        case "hello":
            if (2 < 1) {
                var v124 bool = false
                v125 := []int{}
                _ = v124
                _ = v125
            }
        }
    } else {
        fmt.Println(v3, v2)
        v3 = (v3 / (0.0 + 3.14159))
        v126 := GenS4{F0: false, F1: 127}
        _ = v126
    }
    v127 := []int{}
    v127 = append(v127, 0)
    switch 9223372036854775807 {
    case -2147483648:
        v128 := GenS12{F0: 128}
        _ = v128
    case 9223372036854775807:
        if true {
            v3 -= 3.14159
            v2 = ("hello" + ((v2 + "") + ((v2 + "\n\t") + v2)))
        }
    default:
        v3 += float32(float64(1.5))
    }
    _ = len(v127)
    var v129 uint16 = 975
    v3 += v3
    v130 := GenS1{F0: 193}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v127:", len(v127), v127)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v127
    _ = v129
    _ = v130
    _ = v2
    _ = v3
    _ = v55
}
