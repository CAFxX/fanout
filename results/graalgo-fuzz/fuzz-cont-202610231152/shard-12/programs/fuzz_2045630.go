package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 int; F2 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool; F1 int32; F2 string }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 int64; F1 string }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int64; F1 uint8; F2 int64 }
func genFn0() int16 {
    var v5 int64 = -9223372036854775808
    v5 += v5
    v6 := GenS1{F0: 0}
    fmt.Println(v5)
    if (!false) {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 uint32 = 44
            switch (("\n\t" + "") + ("\n\t" + "")) {
            case "foo bar":
                var v9 float64 = (0.1 - 1e10)
                _ = v9
            default:
                v10 := GenS1{F0: 925}
                _ = v10
            }
            v11 := []int{(fi7 / (((-(fi7))) | 1))}
            _ = fi7
            _ = v11
            _ = v8
        }
    } else {
        if true {
            v5 = v5
        } else {
            switch 1 {
            case 0:
                var v12 uint64 = 1
                _ = v12
            default:
                var v13 string = "hello"
                _ = v13
            }
            var v14 bool = (false && (794 >= 32768))
            _ = v14
        }
        if true {
            var v15 uint8 = 173
            switch 9223372036854775807 {
            case -9223372036854775808:
                v16 := []int{}
                _ = v16
            case 32768:
                var v17 int = 203
                _ = v17
            default:
                var v18 string = "hello"
                _ = v18
            }
            _ = v15
        }
        v5 = 0
    }
    _ = v5
    _ = v6
    return -32768
}
func genFn1() int8 {
    if (false && (1 <= 2)) {
        var v19 int64 = -1
        var v20 int = 134
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 float64 = 3.14159
            if ((false && (0 < -9223372036854775808)) && (1 > -116)) {
                v20 = (v20 * fi21)
            } else {
                v23 := GenS2{F0: 2147483647, F1: -826, F2: 65}
                if true {
                    switch "hello" {
                    case "a":
                        var v24 int32 = 1
                        _ = v24
                    case "hello":
                        fmt.Println(fi21)
                    case "foo bar":
                        v22 = (v22 + v22)
                    }
                    v25 := []int{-182}
                    _ = v25
                }
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 string = ("hello" + "a")
                    _ = fi26
                    _ = v27
                }
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    }
    if (true || true) {
        var v28 int8 = -57
        _ = v28
    } else {
        v29 := []int{}
        switch -9223372036854775808 {
        case 2147483647:
            v29 = append(v29, 2147483647)
        case -2147483648:
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
        case 0:
            var v30 uint64 = 684
            _ = v30
        }
        if true {
            var v31 uint32 = 2147483647
            switch 466 {
            case 128:
                v31 = v31
            }
            for fi32 := 0; fi32 < 2; fi32++ {
                for fi33 := 0; fi33 < 16; fi33++ {
                    var v34 uint32 = (v31 - 158)
                    v35 := GenS3{F0: 9223372036854775807, F1: 784}
                    var v36 uint16 = 712
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
                _ = fi32
            }
            _ = v31
        }
        _ = v29
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v38 := GenS4{F0: ("" + ((("a" + (("a" + (("foo bar" + "hello") + ("\n\t" + "\n\t"))) + "")) + "a") + ("" + ("a" + ("a" + "foo bar")))))}
        var v39 uint32 = 0
        if (false && ((^(v39)) <= (^(v39)))) {
            var v40 float32 = 1.5
            _ = v40
        } else {
            v41 := GenS2{F0: 719, F1: 1000, F2: 128}
            _ = v41
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    var v42 int32 = 2147483647
    if true {
        switch "" {
        case "foo bar":
            v43 := GenS4{F0: ("foo bar" + "\n\t")}
            _ = v43
        }
        var v44 int = -2147483648
        _ = v44
    } else {
        var v45 uint32 = 736
        var v46 bool = true
        _ = v45
        _ = v46
    }
    _ = v42
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = ("hello" + ("foo bar" + "a"))
    v3 := []string{((v2 + (v2 + v2)) + v2), v2}
    var v4 uint8 = 127
    var v47 int8 = genFn1()
    switch "hello" {
    case "":
        for fi48 := 0; fi48 < 4; fi48++ {
            v1 = 895
            var v49 uint16 = 42
            _ = fi48
            _ = v49
        }
    case "foo bar":
        if (false && (1 < 42)) {
            v4 -= (^(v4))
            var v50 uint = 18446744073709551615
            _ = v50
        }
    case "hello":
        v4 = 42
    }
    for fi51 := 0; fi51 < 4; fi51++ {
        fmt.Println(v47, v1)
        fmt.Println(v1, v4, v4)
        for fi52 := 0; fi52 < 7; fi52++ {
            if len(v3) > 0 {
                v3[0] = v2
            }
            var v53 string = ""
            _ = fi52
            _ = v53
        }
        _ = fi51
    }
    if true {
        if true {
            for fi54 := 0; fi54 < 19; fi54++ {
                switch (fi54 * (fi54 >> ((42) & 63))) {
                case -2147483648:
                    var v55 int32 = -1
                    _ = v55
                case 4294967295:
                    var v56 float32 = (3.14159 - 3.14159)
                    _ = v56
                case -451:
                    v4 += (v4 << ((4294967295) & 63))
                }
                switch (fi54 % (((fi54 - fi54)) | 1)) {
                case 165:
                    v57 := GenS1{F0: -1}
                    _ = v57
                case 42:
                    v2 = (v2 + "\n\t")
                case 128:
                    var v58 string = v2
                    _ = v58
                }
                _ = fi54
            }
        } else {
            var v59 bool = true
            if true {
                v60 := GenS5{F0: 65535}
                for fi61 := 0; fi61 < 20; fi61++ {
                    var v62 uint64 = (v1 - (v1 | v1))
                    for fi63 := 0; fi63 < 12; fi63++ {
                        v64 := GenS5{F0: 889}
                        var v65 int = (^(fi63))
                        switch v2 {
                        case "foo bar":
                            if v59 {
                                v66 := GenS6{F0: 42}
                                _ = v66
                            }
                        case "\n\t":
                            if false {
                                var v67 uint16 = 1
                                var v68 string = v2
                                var v69 int = -619
                                _ = v67
                                _ = v68
                                _ = v69
                            } else {
                                var v70 int = v65
                                v71 := GenS2{F0: 2147483647, F1: genFn0(), F2: (^(v4))}
                                _ = v70
                                _ = v71
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                    _ = fi61
                    _ = v62
                }
                fmt.Println(v2, v2)
                _ = v60
            } else {
                var v72 int8 = (v47 & v47)
                if (false || (false || (708 <= 937))) {
                    switch 128 {
                    case 877:
                        switch v2 {
                        case "foo bar":
                            for fi73 := 0; fi73 < 19; fi73++ {
                                fmt.Println(v1, v1, v4)
                                if (0 > 0) {
                                    for fi74 := 0; fi74 < 4; fi74++ {
                                        fmt.Println(v1, fi73)
                                        switch (fi74 - (fi74 + fi74)) {
                                        case 994:
                                            if false {
                                                var v75 int8 = v47
                                                if len(v3) > 0 {
                                                    v3[0] = (v2 + "a")
                                                }
                                                if true {
                                                    v59 = (float64((1e10 + float32(1.5))) > (1e10 / -2.25))
                                                    _ = len(v3)
                                                    v76 := GenS6{F0: -9223372036854775808}
                                                    fmt.Println(v59, v1, v72)
                                                    _ = v76
                                                }
                                                var v77 uint = 42
                                                _ = v75
                                                _ = v77
                                            }
                                        case -384:
                                            if (((2 > 1) || false) && true) {
                                                switch v2 {
                                                case "a":
                                                    if v59 {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v3 = append(v3, ("foo bar" + (v2 + v2)))
                                                        }
                                                        v78 := GenS7{F0: 4294967295, F1: (fi73 + 32768), F2: v1}
                                                        var v79 float32 = (0.1 - (float32(-2.25) / 0.1))
                                                        v72 += v47
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                case "hello":
                                                    fmt.Println(v4, fi73)
                                                }
                                                fmt.Println(v72, v72, fi73)
                                            }
                                        default:
                                            var v80 int = (fi74 << ((42) & 63))
                                            _ = v80
                                        }
                                        _ = fi74
                                    }
                                    var v81 bool = false
                                    for fi82 := 0; fi82 < 5; fi82++ {
                                        v59 = (-2.25 <= (1.5 - -2.25))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch -901 {
                                        case -1:
                                            if len(v3) > 0 {
                                                v3[0] = v2
                                            }
                                            switch (v2 + ((v2 + v2) + (("foo bar" + v2) + v2))) {
                                            case "":
                                                switch v2 {
                                                case "a":
                                                    v1 -= v1
                                                    fmt.Println(v4, v47, fi73)
                                                    var v83 uint32 = 2147483648
                                                    _ = v83
                                                case "foo bar":
                                                    v59 = false
                                                default:
                                                    switch (v2 + v2) {
                                                    case "\n\t":
                                                        if len(v3) > 0 {
                                                            v3[0] = (v2 + v2)
                                                        }
                                                    case "":
                                                        var v84 int16 = -804
                                                        _ = v84
                                                    case "hello":
                                                        var v85 bool = ((-2.25 * 3.14159) != -2.25)
                                                        _ = v85
                                                    default:
                                                        v86 := GenS8{F0: 42}
                                                        _ = v86
                                                    }
                                                }
                                            case "hello":
                                                if len(v3) > 0 {
                                                    v3[0] = v2
                                                }
                                            default:
                                                var v87 string = v2
                                                _ = v87
                                            }
                                        case 128:
                                            if true {
                                                var v88 string = ("" + v2)
                                                if v81 {
                                                    if (!true) {
                                                        switch fi73 {
                                                        case 32767:
                                                            if (!(1 > 255)) {
                                                                var v89 float64 = 3.14159
                                                                v90 := GenS2{F0: 42, F1: -32768, F2: v4}
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                v91 := GenS9{F0: v81, F1: 685, F2: v2}
                                                                var v92 float32 = 0.0
                                                                if ((^(v47)) != (-(v72))) {
                                                                    v59 = v81
                                                                    if ((^(v72)) != v47) {
                                                                        switch "" {
                                                                        case "":
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                if (v81 && (65535 < 0)) {
                                                                                    v4 = (v4 | (v4 | (v4 >> ((255) & 63))))
                                                                                    v47 -= v72
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            if true {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                v93 := GenS10{F0: (^(v47)), F1: 4294967295}
                                                                                _ = v93
                                                                            } else {
                                                                                if (v59 || (((!v81) && (false || v81)) || v59)) {
                                                                                    var v94 uint16 = 2
                                                                                    fmt.Println(v47)
                                                                                    _ = v94
                                                                                } else {
                                                                                    if true {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "foo bar"
                                                                                        }
                                                                                        var v95 uint16 = 65535
                                                                                        _ = v95
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            v96 := GenS10{F0: 1, F1: 1}
                                                                            _ = v96
                                                                        default:
                                                                            v97 := GenS5{F0: 0}
                                                                            _ = v97
                                                                        }
                                                                    } else {
                                                                        if v59 {
                                                                            _ = len(v3)
                                                                            var v98 uint64 = uint64(v4)
                                                                            _ = v98
                                                                        } else {
                                                                            v92 = 0.0
                                                                            fmt.Println(v88, v72)
                                                                        }
                                                                        v92 += 0.0
                                                                    }
                                                                } else {
                                                                    if (fi73 < (fi73 << ((1) & 63))) {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v4 = v4
                                                                        default:
                                                                            var v99 uint8 = 2
                                                                            _ = v99
                                                                        }
                                                                    }
                                                                }
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                        case -997:
                                                            v72 = 1
                                                        case 9223372036854775807:
                                                            v100 := GenS11{F0: 0, F1: v88}
                                                            _ = v100
                                                        default:
                                                            v59 = ("\n\t" < (v88 + v88))
                                                        }
                                                    }
                                                } else {
                                                    v2 = ((("foo bar" + v2) + v88) + v2)
                                                    var v101 int32 = 2147483647
                                                    if true {
                                                        if len(v3) > 0 {
                                                            v3[0] = v88
                                                        }
                                                        var v102 int8 = (v72 & 101)
                                                        _ = v102
                                                    }
                                                    _ = v101
                                                }
                                                v103 := GenS11{F0: 9223372036854775807, F1: "hello"}
                                                _ = v103
                                                _ = v88
                                            } else {
                                                switch "hello" {
                                                case "hello":
                                                    v81 = v81
                                                case "a":
                                                    if (-2.25 < 1e10) {
                                                        var v104 int = (fi73 / ((-648) | 1))
                                                        v72 += v47
                                                        if len(v3) > 0 {
                                                            v3[0] = "hello"
                                                        }
                                                        _ = v104
                                                    }
                                                case "\n\t":
                                                    switch v2 {
                                                    case "a":
                                                        if (v47 <= v47) {
                                                            switch fi82 {
                                                            case -398:
                                                                switch (fi82 - fi82) {
                                                                case 0:
                                                                    v72 = v47
                                                                default:
                                                                    v105 := GenS5{F0: 0}
                                                                    _ = v105
                                                                }
                                                            case -2147483648:
                                                                v106 := GenS12{F0: v81}
                                                                _ = v106
                                                            case 1:
                                                                var v107 uint64 = v1
                                                                _ = v107
                                                            }
                                                        }
                                                    case "hello":
                                                        v3 = append(v3, ("a" + ((v2 + v2) + ("hello" + (("a" + (((v2 + v2) + v2) + ("a" + v2))) + "\n\t")))))
                                                    case "foo bar":
                                                        v72 = genFn1()
                                                    default:
                                                        v108 := GenS11{F0: -1, F1: ("\n\t" + "\n\t")}
                                                        _ = v108
                                                    }
                                                default:
                                                    fmt.Println(fi82, v1, v81)
                                                }
                                                var v109 uint16 = 255
                                                _ = v109
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = v2
                                            }
                                        }
                                        _ = fi82
                                    }
                                    _ = v81
                                } else {
                                    var v110 int8 = 100
                                    var v111 float64 = -2.25
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    _ = v110
                                    _ = v111
                                }
                                var v112 int32 = 42
                                _ = fi73
                                _ = v112
                            }
                        case "a":
                            var v113 float64 = 1.5
                            _ = v113
                        default:
                            v1 -= (v1 ^ (v1 - 879))
                        }
                    case 128:
                        v47 += v47
                    }
                    for fi114 := 0; fi114 < 10; fi114++ {
                        v72 += v72
                        if v59 {
                            v59 = (v59 || ((true && (false && ((v59 && v59) || false))) || (255 > 4294967295)))
                        } else {
                            var v115 uint16 = 194
                            _ = v115
                        }
                        if v59 {
                            if v59 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v116 := GenS5{F0: 2}
                                for fi117 := 0; fi117 < 8; fi117++ {
                                    var v118 string = v2
                                    _ = fi117
                                    _ = v118
                                }
                                _ = v116
                            } else {
                                if (true || true) {
                                    if true {
                                        if len(v3) > 0 {
                                            v3[0] = v2
                                        }
                                        v1 -= v1
                                    } else {
                                        v2 = v2
                                        for fi119 := 0; fi119 < 10; fi119++ {
                                            _ = len(v3)
                                            _ = fi119
                                        }
                                        var v120 float64 = (3.14159 / 3.14159)
                                        _ = v120
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = v2
                                }
                                v2 = v2
                            }
                            switch (fi114 >> ((431) & 63)) {
                            case -871:
                                v1 -= v1
                            case 127:
                                var v121 uint = 2
                                _ = v121
                            case 256:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        _ = fi114
                    }
                } else {
                    v59 = (591 != 42)
                    for fi122 := 0; fi122 < 7; fi122++ {
                        switch 9223372036854775807 {
                        case 42:
                            v59 = v59
                        case -9223372036854775808:
                            var v123 uint = 2
                            _ = v123
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi122
                    }
                    var v124 uint16 = 878
                    _ = v124
                }
                _ = v72
            }
            _ = v59
        }
        v125 := GenS11{F0: 0, F1: "hello"}
        _ = v125
    } else {
        fmt.Println(v47, v2)
    }
    switch 1 {
    case 1:
        v3 = append(v3, ("\n\t" + v2))
    }
    v126 := GenS12{F0: (!true)}
    for fi127 := 0; fi127 < 13; fi127++ {
        var v128 int = fi127
        switch 32767 {
        case -837:
            for fi129 := 0; fi129 < 3; fi129++ {
                v47 = (v47 / (((v47 >> ((0) & 63))) | 1))
                _ = fi129
            }
        case 1:
            _ = len(v3)
        case 4294967295:
            v1 += 2
        }
        v4 += v4
        _ = fi127
        _ = v128
    }
    var v130 int32 = 801
    v131 := GenS13{F0: 125, F1: (v4 % (((v4 - v4)) | 1)), F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v126
    _ = v130
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v47
}
