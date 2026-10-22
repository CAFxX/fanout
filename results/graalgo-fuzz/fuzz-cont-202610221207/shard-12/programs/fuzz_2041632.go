package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 int8 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 uint64; F1 uint }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint; F1 string; F2 bool }
type GenS15 struct { F0 uint8; F1 string }
type GenS16 struct { F0 string; F1 uint32; F2 int }
type GenS17 struct { F0 uint16; F1 int }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 int64; F1 int32; F2 uint8 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 int16; F1 int64; F2 uint64 }
type GenS22 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS23 struct { F0 int }
type GenS24 struct { F0 uint8; F1 bool; F2 int8 }
type GenS25 struct { F0 int64; F1 uint8 }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS28 struct { F0 uint64 }
func genFn0() int16 {
    var v7 bool = (false || false)
    if v7 {
        v8 := []int{127, 127, 42}
        _ = v8
    } else {
        var v9 uint16 = 76
        _ = v9
    }
    _ = v7
    return -833
}
func genFn1(p0 uint) string {
    v10 := []int{479, 9223372036854775807}
    p0 = 255
    fmt.Println(p0)
    v11 := GenS1{F0: 0, F1: 39}
    if false {
        switch "hello" {
        case "a":
            if len(v10) > 0 {
                v10[0] = -2147483648
            }
        case "foo bar":
            for fi12 := 0; fi12 < 12; fi12++ {
                if (2 >= 632) {
                    p0 = (p0 | (p0 + (^(p0))))
                    if len(v10) > 0 {
                        v10[0] = -9223372036854775808
                    }
                    if false {
                        p0 -= (^(p0))
                        var v13 uint64 = 1
                        switch "foo bar" {
                        case "a":
                            var v14 uint8 = 255
                            _ = v14
                        }
                        _ = v13
                    }
                } else {
                    var v15 float32 = -2.25
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = v15
                }
                switch fi12 {
                case 875:
                    if ((9223372036854775807 >= -967) && true) {
                        fmt.Println(p0)
                    } else {
                        var v16 int16 = -1
                        var v17 int64 = -9223372036854775808
                        for fi18 := 0; fi18 < 11; fi18++ {
                            if false {
                                v17 = (v17 % ((v17) | 1))
                                if len(v10) > 0 {
                                    v10[0] = fi12
                                }
                                if false {
                                    if len(v10) > 0 {
                                        v10[0] = 127
                                    }
                                    switch (fi12 & (fi18 ^ (fi18 / ((fi12) | 1)))) {
                                    case -2147483648:
                                        p0 += p0
                                    case 4294967295:
                                        p0 = p0
                                    case 1:
                                        v19 := GenS2{F0: ""}
                                        _ = v19
                                    }
                                } else {
                                    var v20 string = ("" + "a")
                                    _ = v20
                                }
                            } else {
                                fmt.Println(v17, v16, v17)
                                v21 := GenS1{F0: 0, F1: 42}
                                _ = v21
                            }
                            if false {
                                if (163 == 168) {
                                    var v22 float64 = 3.14159
                                    v16 -= v16
                                    _ = v22
                                }
                                var v23 float64 = (1e10 + float64((0.0 - 0.0)))
                                fmt.Println(v17)
                                v24 := GenS1{F0: 2, F1: 42}
                                var v25 int64 = (v17 * v17)
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            _ = fi18
                        }
                        _ = v16
                        _ = v17
                    }
                case -695:
                    var v26 bool = false
                    _ = v26
                case -1:
                    switch "hello" {
                    case "\n\t":
                        _ = len(v10)
                    }
                default:
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 float32 = 1e10
                        var v29 int64 = -9223372036854775808
                        var v30 int8 = -3
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi12
            }
        default:
            v31 := GenS3{F0: 32767, F1: 2}
            _ = v31
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = v10
    _ = v11
    return ""
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 918
    var v3 bool = (-1 != -9223372036854775808)
    var v4 int16 = 1
    var v5 uint = 0
    var v6 uint32 = 42
    fmt.Println(v2)
    v4 = v4
    if (float32((1e10 + float64(0.0))) <= ((float32(1.5) * (1.5 + 0.0)) - 0.0)) {
        switch -383 {
        case -9223372036854775808:
            v32 := GenS4{F0: genFn0(), F1: v2, F2: 83}
            _ = v32
        case 140:
            switch 127 {
            case 127:
                v33 := GenS1{F0: 0, F1: 1}
                _ = v33
            case -2147483648:
                v2 -= (v2 >> (((v5 - 1)) & 63))
            case 9223372036854775807:
                v4 = (^(v4))
            default:
                switch "" {
                case "":
                    v34 := GenS3{F0: 43, F1: (v6 + v6)}
                    _ = v34
                case "hello":
                    var v35 uint64 = 4294967295
                    _ = v35
                }
            }
        case -386:
            var v36 uint64 = 0
            _ = v36
        default:
            if v3 {
                v5 = v5
            }
        }
        v37 := GenS3{F0: 32767, F1: 42}
        _ = v37
    } else {
        switch 32767 {
        case 969:
            for fi38 := 0; fi38 < 2; fi38++ {
                if v3 {
                    switch ("" + ("\n\t" + "a")) {
                    case "foo bar":
                        for fi39 := 0; fi39 < 8; fi39++ {
                            fmt.Println(v4)
                            var v40 uint32 = v6
                            v6 = v6
                            _ = fi39
                            _ = v40
                        }
                    case "a":
                        v41 := []int{}
                        _ = v41
                    default:
                        v42 := v1["\n\t"]
                        _ = v42
                    }
                    var v43 uint16 = 32767
                    v5 = 524
                    _ = v43
                } else {
                    v44 := GenS1{F0: 255, F1: -110}
                    if (0.1 <= (-2.25 / 1.5)) {
                        v45 := GenS5{F0: 127, F1: v4}
                        _ = v45
                    }
                    _ = v44
                }
                if true {
                    var v46 int16 = -353
                    _ = v46
                }
                var v47 uint32 = v6
                _ = fi38
                _ = v47
            }
        }
        v48 := []int{}
        _ = v48
    }
    v49 := GenS6{F0: 4294967295}
    for fi50 := 0; fi50 < 18; fi50++ {
        for fi51 := 0; fi51 < 7; fi51++ {
            var v52 int32 = (v2 ^ (v2 % (((v2 * (^(v2)))) | 1)))
            _ = fi51
            _ = v52
        }
        var v53 bool = (!true)
        for fi54 := 0; fi54 < 20; fi54++ {
            v55 := GenS6{F0: 1}
            _ = fi54
            _ = v55
        }
        _ = fi50
        _ = v53
    }
    v6 = 1
    fmt.Println(v6, v2)
    var v56 uint32 = 0
    var v57 float32 = 0.0
    if (float64((v57 * 1.5)) == -2.25) {
        if v3 {
            v4 = (v4 << (((v5 * v5)) & 63))
            var v58 uint16 = 255
            _ = len(v1)
            _ = v58
        }
        v57 = v57
        for fi59 := 0; fi59 < 14; fi59++ {
            for fi60 := 0; fi60 < 2; fi60++ {
                v61 := GenS2{F0: (("hello" + "") + "foo bar")}
                if v3 {
                    switch "" {
                    case "":
                        fmt.Println(v4)
                    case "a":
                        if (float32((-2.25 / (-2.25 - float64(v57)))) < v57) {
                            var v62 int8 = -1
                            _ = v62
                        }
                    case "\n\t":
                        if false {
                            v6 += v56
                            fmt.Println(v5, v57, v4)
                            var v63 uint16 = 0
                            _ = v63
                        } else {
                            v5 = (v5 << ((18446744073709551615) & 63))
                            if false {
                                switch fi59 {
                                case 4294967295:
                                    v64 := GenS1{F0: 43, F1: 42}
                                    _ = v64
                                case 0:
                                    if ((("foo bar" + "\n\t") != "") && v3) {
                                        var v65 uint8 = 143
                                        _ = v65
                                    } else {
                                        switch -407 {
                                        case 4294967295:
                                            v66 := GenS7{F0: v2, F1: 18}
                                            _ = v66
                                        case 584:
                                            fmt.Println(fi60, v6)
                                        case -1:
                                            if (0 >= 100) {
                                                var v67 int32 = -863
                                                if v3 {
                                                    var v68 float64 = (float64(v57) + ((1e10 - (float64(v57) + 1.5)) * float64((1e10 + v57))))
                                                    if (!v3) {
                                                        var v69 int32 = (v67 % ((-1) | 1))
                                                        _ = v69
                                                    } else {
                                                        var v70 int64 = 1
                                                        if (((v68 + 1.5) - (-2.25 + (-2.25 + 0.0))) != float64(0.0)) {
                                                            switch (fi59 >> (((v5 % ((v5) | 1))) & 63)) {
                                                            case 263:
                                                                v4 -= genFn0()
                                                            default:
                                                                v71 := []int{(-(fi59)), (-(fi60)), -9223372036854775808}
                                                                _ = v71
                                                            }
                                                            v72 := GenS8{F0: (int32(v70) > (v67 * -2147483648))}
                                                            v3 = (v68 == (v68 - v68))
                                                            _ = v72
                                                        }
                                                        _ = v70
                                                    }
                                                    if ((!true) && true) {
                                                        fmt.Println(v5, v57, v5)
                                                    } else {
                                                        v73 := v1["foo bar"]
                                                        v74 := GenS2{F0: v73}
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                    v2 -= v67
                                                    if (1 <= uint64(v56)) {
                                                        v67 -= v2
                                                        v1["\n\t"] = (("hello" + "hello") + "foo bar")
                                                    } else {
                                                        switch 1 {
                                                        case 2147483647:
                                                            delete(v1, "a")
                                                            v75 := GenS4{F0: v4, F1: (v2 << ((94) & 63)), F2: 127}
                                                            if (v3 && v3) {
                                                                var v76 int8 = int8(v5)
                                                                var v77 int16 = (v4 ^ genFn0())
                                                                v78 := GenS4{F0: v4, F1: int32(fi60), F2: (v76 - v76)}
                                                                _ = v76
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            _ = v75
                                                        case 1:
                                                            v1[("hello" + ("foo bar" + ("" + "hello")))] = "\n\t"
                                                        case 0:
                                                            var v79 uint = 255
                                                            _ = v79
                                                        }
                                                        switch "" {
                                                        case "\n\t":
                                                            switch "" {
                                                            case "\n\t":
                                                                if true {
                                                                    v80 := GenS9{F0: (^(fi60)), F1: 42, F2: 282}
                                                                    var v81 int32 = v2
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    v3 = false
                                                                    if (77 > -128) {
                                                                        v5 = (v5 << ((0) & 63))
                                                                        var v82 int8 = 127
                                                                        v1["\n\t"] = (("a" + "foo bar") + ("\n\t" + ("a" + ("a" + ""))))
                                                                        v83 := GenS3{F0: 42, F1: v56}
                                                                        _ = v82
                                                                        _ = v83
                                                                    } else {
                                                                        v84 := GenS2{F0: "hello"}
                                                                        var v85 uint8 = 2
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                    var v86 float64 = ((v68 + (0.0 * 3.14159)) + float64(3.14159))
                                                                    _ = v86
                                                                }
                                                            case "hello":
                                                                var v87 uint = v5
                                                                _ = v87
                                                            }
                                                        case "foo bar":
                                                            if (v3 && (v3 || v3)) {
                                                                if (!(("foo bar" == "a") || v3)) {
                                                                    v88 := GenS10{F0: 255, F1: 212}
                                                                    v5 = (v5 >> (((v5 - v5)) & 63))
                                                                    v4 = genFn0()
                                                                    _ = v88
                                                                } else {
                                                                    v89 := []int{fi60, (fi60 | 458)}
                                                                    _ = v89
                                                                }
                                                            } else {
                                                                if ((^(v5)) > 42) {
                                                                    switch fi60 {
                                                                    case 9223372036854775807:
                                                                        if (v67 != v67) {
                                                                            fmt.Println(v3, v6)
                                                                            v1[""] = "a"
                                                                            v90 := GenS8{F0: v3}
                                                                            _ = v90
                                                                        }
                                                                    case 4294967295:
                                                                        fmt.Println(v3, v56)
                                                                    }
                                                                    v2 = (v67 & (v2 * -2147483648))
                                                                    var v91 bool = ((("" + "\n\t") + "") <= "hello")
                                                                    v3 = ((-(v4)) <= 0)
                                                                    fmt.Println(fi59, v91, v91)
                                                                    v6 -= v56
                                                                    var v92 int64 = 9223372036854775807
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    switch fi60 {
                                                                    case 42:
                                                                        v5 += v5
                                                                    case 0:
                                                                        v1["\n\t"] = (("" + "a") + "a")
                                                                    }
                                                                    v4 += (v4 | -596)
                                                                }
                                                                v93 := GenS10{F0: 1, F1: 204}
                                                                _ = v93
                                                            }
                                                        case "":
                                                            switch "hello" {
                                                            case "hello":
                                                                v3 = ((v4 << ((v5) & 63)) <= (v4 >> ((v5) & 63)))
                                                                v3 = (36 > 42)
                                                            case "\n\t":
                                                                v94 := GenS1{F0: 127, F1: -1}
                                                                _ = v94
                                                            }
                                                        }
                                                        if v3 {
                                                            if (!((v57 / (1e10 / (v57 + v57))) > (1e10 * float32(v68)))) {
                                                                if false {
                                                                    v95 := []int{(-(fi60)), fi60, fi60, fi59}
                                                                    _ = v95
                                                                } else {
                                                                    v96 := GenS11{F0: 18446744073709551615, F1: (v5 * (v5 >> ((871) & 63)))}
                                                                    v4 -= 459
                                                                    _ = v96
                                                                }
                                                            } else {
                                                                v6 = 928
                                                            }
                                                            v97 := []int{}
                                                            _ = v97
                                                        }
                                                    }
                                                    _ = v68
                                                } else {
                                                    var v98 int16 = (v4 & v4)
                                                    if ((("" + "hello") > "\n\t") || false) {
                                                        v99 := GenS2{F0: "\n\t"}
                                                        v100 := GenS12{F0: int8(v5)}
                                                        fmt.Println(v57)
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                    v4 += (v98 / (((v98 | v98)) | 1))
                                                    _ = v98
                                                }
                                                var v101 int64 = 191
                                                _ = v101
                                                _ = v67
                                            }
                                        default:
                                            var v102 float32 = v57
                                            _ = v102
                                        }
                                        v103 := GenS13{F0: (v4 << ((v5) & 63))}
                                        var v104 uint16 = 42
                                        _ = v103
                                        _ = v104
                                    }
                                default:
                                    switch "a" {
                                    case "hello":
                                        var v105 uint8 = 33
                                        _ = v105
                                    case "foo bar":
                                        v106 := GenS14{F0: (v5 ^ (^(v5))), F1: "\n\t", F2: false}
                                        _ = v106
                                    default:
                                        if (-37 == (v4 | (v4 - 196))) {
                                            if v3 {
                                                v57 = (-2.25 - 0.0)
                                                var v107 float64 = (-2.25 * float64(v57))
                                                _ = v107
                                            } else {
                                                if v3 {
                                                    if (2 > 255) {
                                                        fmt.Println(v5)
                                                        delete(v1, "")
                                                        switch ("\n\t" + "a") {
                                                        case "foo bar":
                                                            v2 = v2
                                                        case "":
                                                            if ((v56 | (^(v56))) < v56) {
                                                                v56 = (v6 * (v6 * (^(v6))))
                                                            }
                                                        case "hello":
                                                            v4 += v4
                                                        default:
                                                            v108 := []int{fi60}
                                                            _ = v108
                                                        }
                                                        v109 := GenS3{F0: uint16(v56), F1: (v6 | (v6 / (((v56 - (v6 / (((^(v56))) | 1)))) | 1)))}
                                                        _ = v109
                                                    }
                                                    v110 := GenS15{F0: 0, F1: "\n\t"}
                                                    var v111 uint16 = 0
                                                    var v112 uint32 = (v6 * (v6 >> ((uint(v56)) & 63)))
                                                    _ = v110
                                                    _ = v111
                                                    _ = v112
                                                }
                                                v113 := v1[("\n\t" + "foo bar")]
                                                switch 2147483647 {
                                                case 933:
                                                    switch (v113 + v113) {
                                                    case "":
                                                        v3 = v3
                                                    case "hello":
                                                        var v114 uint8 = 128
                                                        _ = v114
                                                    case "\n\t":
                                                        if ((3.14159 - ((((0.0 + (3.14159 - (1e10 + 1.5))) * 3.14159) / float64(v57)) * float64((v57 - v57)))) > 3.14159) {
                                                            v115 := GenS7{F0: v2, F1: 1}
                                                            var v116 uint64 = 607
                                                            var v117 uint32 = (v56 | (v56 ^ 65535))
                                                            v118 := []int{fi59, int(v4), (fi60 & 4294967295), (fi60 << ((611) & 63))}
                                                            v2 = v2
                                                            _ = v115
                                                            _ = v116
                                                            _ = v117
                                                            _ = v118
                                                        } else {
                                                            if v3 {
                                                                var v119 uint64 = 433
                                                                var v120 float64 = (((3.14159 + (3.14159 + float64((v57 * v57)))) * float64(float32(3.14159))) / ((1e10 * 1.5) - 0.1))
                                                                var v121 int64 = 0
                                                                _ = v119
                                                                _ = v120
                                                                _ = v121
                                                            } else {
                                                                v56 = v56
                                                                switch (fi59 & 32767) {
                                                                case 128:
                                                                    v122 := []int{fi59, (fi59 / (((^(fi60))) | 1))}
                                                                    _ = v122
                                                                case 32768:
                                                                    if (fi60 == (-(fi59))) {
                                                                        var v123 bool = v3
                                                                        _ = v123
                                                                    }
                                                                case -9223372036854775808:
                                                                    v3 = (!v3)
                                                                default:
                                                                    var v124 float64 = 0.0
                                                                    _ = v124
                                                                }
                                                            }
                                                            fmt.Println(v4, v57, fi59)
                                                        }
                                                        switch (v113 + "hello") {
                                                        case "":
                                                            if v3 {
                                                                v57 = v57
                                                            } else {
                                                                var v125 int8 = 120
                                                                if (9223372036854775807 < -9223372036854775808) {
                                                                    v2 -= -641
                                                                    var v126 float64 = (-2.25 - 1e10)
                                                                    delete(v1, (v113 + v113))
                                                                    _ = v126
                                                                }
                                                                fmt.Println(v113, v3, v6)
                                                                _ = v125
                                                            }
                                                        case "a":
                                                            var v127 int32 = v2
                                                            _ = v127
                                                        }
                                                        fmt.Println(v3, v3)
                                                    default:
                                                        fmt.Println(v113, v2)
                                                    }
                                                default:
                                                    var v128 int = -315
                                                    _ = v128
                                                }
                                                v129 := []int{}
                                                _ = v113
                                                _ = v129
                                            }
                                        }
                                    }
                                }
                                v6 = uint32(v5)
                                v3 = v3
                            }
                        }
                    default:
                        v130 := []int{(fi60 & (fi59 / ((12) | 1))), (-(fi60)), fi60}
                        _ = v130
                    }
                    var v131 bool = true
                    _ = v131
                }
                _ = fi60
                _ = v61
            }
            v132 := GenS16{F0: ("foo bar" + "a"), F1: (^(v6)), F2: 0}
            _ = fi59
            _ = v132
        }
    }
    switch ("hello" + "hello") {
    case "foo bar":
        for fi133 := 0; fi133 < 4; fi133++ {
            switch "foo bar" {
            case "a":
                if v3 {
                    for fi134 := 0; fi134 < 13; fi134++ {
                        switch (fi134 - (-(fi134))) {
                        case 32768:
                            if v3 {
                                if true {
                                    fmt.Println(v57)
                                    var v135 uint16 = 42
                                    v1["foo bar"] = ("" + ((("" + "foo bar") + "foo bar") + (("hello" + ("foo bar" + "")) + "\n\t")))
                                    var v136 int64 = 1
                                    delete(v1, (("hello" + "a") + "hello"))
                                    fmt.Println(v56)
                                    _ = v135
                                    _ = v136
                                }
                                v137 := GenS17{F0: 251, F1: (fi133 / ((fi134) | 1))}
                                switch "" {
                                case "hello":
                                    fmt.Println(v56, v4)
                                default:
                                    v138 := []int{int(v4), fi134}
                                    _ = v138
                                }
                                _ = v137
                            } else {
                                v139 := GenS10{F0: 128, F1: 5}
                                var v140 int64 = 9223372036854775807
                                _ = v139
                                _ = v140
                            }
                        case 9223372036854775807:
                            v4 += v4
                        default:
                            v141 := GenS18{F0: 128}
                            _ = v141
                        }
                        if true {
                            if false {
                                fmt.Println(v5, v57, v5)
                                var v142 int = (fi134 * (fi134 - fi134))
                                fmt.Println(v142, fi134, v6)
                                _ = v142
                            }
                            if (-1 != -1) {
                                v5 += (v5 & 2)
                                if (v3 && (!((v4 ^ 0) > (-(v4))))) {
                                    if ((v5 - (v5 << ((v5) & 63))) != v5) {
                                        v143 := GenS14{F0: (v5 / ((v5) | 1)), F1: "\n\t", F2: (3.14159 >= ((((0.0 + 1.5) + 1.5) / (0.0 - float64(v57))) / -2.25))}
                                        if (-39 <= 1) {
                                            fmt.Println(v4, v5, v4)
                                        }
                                        fmt.Println(v3, v57, v4)
                                        _ = v143
                                    }
                                } else {
                                    switch -9223372036854775808 {
                                    case 2147483647:
                                        v56 -= 65535
                                    case 0:
                                        if ((v2 | -2147483648) <= v2) {
                                            var v144 uint8 = 200
                                            _ = v144
                                        }
                                    }
                                    var v145 uint8 = 128
                                    fmt.Println(v57, v145, fi133)
                                    _ = v145
                                }
                                fmt.Println(fi134)
                            }
                            if (!(v2 != v2)) {
                                v57 -= -2.25
                                fmt.Println(fi134, v5, v4)
                                var v146 string = ("a" + "foo bar")
                                if (!((v57 / (v57 * (1.5 + -2.25))) > float32(3.14159))) {
                                    v6 = (^(v56))
                                    if ("a" != v146) {
                                        var v147 uint = (v5 * (v5 & v5))
                                        if (!(!v3)) {
                                            if ((v5 * (v147 * 1)) == v147) {
                                                v148 := GenS19{F0: -693, F1: (v2 & v2), F2: 103}
                                                _ = v148
                                            } else {
                                                var v149 int32 = v2
                                                switch (^(fi134)) {
                                                case 42:
                                                    var v150 string = v146
                                                    switch v150 {
                                                    case "":
                                                        var v151 int8 = 42
                                                        _ = v151
                                                    }
                                                    _ = v150
                                                case -9223372036854775808:
                                                    if (fi134 <= -28) {
                                                        v5 += 18446744073709551615
                                                        if v3 {
                                                            v152 := GenS18{F0: 205}
                                                            if false {
                                                                v153 := GenS20{F0: fi134}
                                                                _ = v153
                                                            }
                                                            _ = v152
                                                        } else {
                                                            var v154 float32 = 1e10
                                                            if true {
                                                                if v3 {
                                                                    v155 := []int{(fi134 + (fi134 % (((fi134 & (-(fi133)))) | 1))), (-(fi134)), 538}
                                                                    var v156 float64 = 0.0
                                                                    _ = v155
                                                                    _ = v156
                                                                }
                                                                var v157 uint = v147
                                                                if v3 {
                                                                    var v158 uint32 = 648
                                                                    _ = v158
                                                                } else {
                                                                    v149 += v149
                                                                    if (117 == 1) {
                                                                        v159 := GenS20{F0: int(v56)}
                                                                        _ = v159
                                                                    } else {
                                                                        v160 := GenS3{F0: 23, F1: 42}
                                                                        switch (v146 + v146) {
                                                                        case "":
                                                                            var v161 uint32 = 65535
                                                                            _ = v161
                                                                        case "\n\t":
                                                                            v162 := []int{}
                                                                            _ = v162
                                                                        default:
                                                                            if (v146 == v146) {
                                                                                var v163 float64 = -2.25
                                                                                v164 := GenS7{F0: (v2 / ((v2) | 1)), F1: 1}
                                                                                _ = v163
                                                                                _ = v164
                                                                            } else {
                                                                                if (int8(v149) >= -1) {
                                                                                    var v165 float32 = 3.14159
                                                                                    v166 := []int{(fi133 * (fi134 << ((v5) & 63))), fi133, 32767}
                                                                                    _ = v165
                                                                                    _ = v166
                                                                                }
                                                                                v167 := GenS21{F0: v4, F1: 9223372036854775807, F2: uint64(v5)}
                                                                                _ = v167
                                                                            }
                                                                        }
                                                                        _ = v160
                                                                    }
                                                                    if (32767 >= 590) {
                                                                        v57 += 0.1
                                                                    }
                                                                }
                                                                v57 = v154
                                                                v4 += int16(v6)
                                                                _ = v157
                                                            } else {
                                                                fmt.Println(v5, v3, fi134)
                                                                if (!(((v154 / ((v57 / ((-2.25 / 1e10) / v57)) - ((v57 + v154) / ((v57 - float32((3.14159 + 0.1))) - ((v57 * (1.5 + v154)) - ((((((v57 - v154) * (v154 / float32(1e10))) + v154) * 0.0) / ((v57 + (3.14159 * (float32(float64(((v154 * 1e10) * -2.25))) - v57))) / ((0.1 + v154) / v57))) + -2.25)))))) * v57) >= (v57 + (v154 / v154)))) {
                                                                    v168 := []int{412, (fi134 / ((fi134) | 1)), (fi134 - 973), (fi133 ^ fi133)}
                                                                    _ = v168
                                                                } else {
                                                                    var v169 int64 = 1
                                                                    delete(v1, "")
                                                                    var v170 uint16 = 32767
                                                                    _ = v169
                                                                    _ = v170
                                                                }
                                                                v171 := GenS15{F0: 147, F1: ("\n\t" + (v146 + "hello"))}
                                                                _ = v171
                                                            }
                                                            fmt.Println(v146, v149)
                                                            _ = v154
                                                        }
                                                    }
                                                case -188:
                                                    v149 = (^(v2))
                                                default:
                                                    v5 = 261
                                                }
                                                delete(v1, v146)
                                                _ = v149
                                            }
                                            var v172 float32 = float32(0.0)
                                            if v3 {
                                                var v173 uint8 = 2
                                                v56 += (^(v6))
                                                _ = v173
                                            }
                                            _ = v172
                                        }
                                        v174 := GenS22{F0: 255, F1: 809, F2: 18446744073709551615}
                                        if ((v56 >= (^(v56))) && v3) {
                                            v146 = v146
                                        }
                                        fmt.Println(v146)
                                        _ = v147
                                        _ = v174
                                    }
                                } else {
                                    v5 += v5
                                    delete(v1, v146)
                                }
                                v175 := GenS23{F0: fi133}
                                _ = v146
                                _ = v175
                            } else {
                                switch ("\n\t" + ("\n\t" + "a")) {
                                case "":
                                    switch 106 {
                                    case 63:
                                        v5 += v5
                                    default:
                                        v56 = (^(v6))
                                    }
                                case "foo bar":
                                    if (v57 != v57) {
                                        _ = len(v1)
                                        v4 = v4
                                        if ((v2 ^ -2147483648) < (^(v2))) {
                                            v56 = (v56 + v56)
                                            switch fi134 {
                                            case 32767:
                                                var v176 uint16 = 329
                                                _ = v176
                                            }
                                        } else {
                                            if (1 >= 255) {
                                                v2 -= (v2 + (v2 >> ((4294967295) & 63)))
                                                v177 := []int{}
                                                _ = v177
                                            } else {
                                                fmt.Println(v5, v6, v6)
                                                if false {
                                                    delete(v1, "")
                                                    var v178 bool = v3
                                                    _ = v178
                                                } else {
                                                    if false {
                                                        v57 -= 0.0
                                                        v5 += (^(v5))
                                                    }
                                                    if (0.1 >= 1e10) {
                                                        fmt.Println(v56)
                                                    }
                                                    v179 := []int{fi133, (fi134 & -1)}
                                                    _ = v179
                                                }
                                                var v180 int64 = 578
                                                if ((^(v6)) == v6) {
                                                    var v181 float64 = ((float64(3.14159) / 1.5) * -2.25)
                                                    v182 := GenS18{F0: 42}
                                                    fmt.Println(v5)
                                                    v3 = (255 >= 0)
                                                    if false {
                                                        delete(v1, "a")
                                                        var v183 int = fi134
                                                        v180 += 42
                                                        _ = v183
                                                    }
                                                    fmt.Println(v4, v3)
                                                    _ = v181
                                                    _ = v182
                                                }
                                                _ = v180
                                            }
                                            v184 := GenS8{F0: ((v3 && ((true && false) || v3)) || v3)}
                                            _ = v184
                                        }
                                        v2 += (v2 % ((v2) | 1))
                                        switch fi133 {
                                        case -97:
                                            if v3 {
                                                var v185 uint64 = 4294967295
                                                v186 := GenS24{F0: 1, F1: v3, F2: 1}
                                                v187 := GenS1{F0: 128, F1: 127}
                                                _ = v185
                                                _ = v186
                                                _ = v187
                                            } else {
                                                v188 := GenS9{F0: 32768, F1: 100, F2: 1}
                                                if v3 {
                                                    if (float32(-2.25) <= 0.0) {
                                                        v2 = int32(v56)
                                                        v3 = ((v5 >> ((uint(v56)) & 63)) > v5)
                                                        fmt.Println(v56, v57)
                                                        var v189 uint8 = 128
                                                        _ = v189
                                                    }
                                                    var v190 bool = true
                                                    fmt.Println(v57, v4)
                                                    _ = v190
                                                } else {
                                                    if true {
                                                        if (!(((1e10 + v57) != v57) || (v3 || v3))) {
                                                            v57 -= -2.25
                                                            if true {
                                                                v2 += v2
                                                                v191 := GenS25{F0: 9223372036854775807, F1: uint8(v5)}
                                                                switch ("foo bar" + "hello") {
                                                                case "":
                                                                    v2 -= int32(v5)
                                                                case "a":
                                                                    v5 += v5
                                                                case "\n\t":
                                                                    var v192 uint32 = (v56 - v56)
                                                                    _ = v192
                                                                default:
                                                                    fmt.Println(v3, v2)
                                                                }
                                                                _ = v191
                                                            }
                                                        }
                                                        _ = len(v1)
                                                        var v193 uint64 = 2
                                                        _ = v193
                                                    } else {
                                                        switch 9223372036854775807 {
                                                        case -2147483648:
                                                            v194 := GenS15{F0: 128, F1: "\n\t"}
                                                            _ = v194
                                                        case 42:
                                                            var v195 string = ""
                                                            _ = v195
                                                        case -977:
                                                            var v196 float32 = ((v57 / v57) * v57)
                                                            _ = v196
                                                        }
                                                        fmt.Println(v2, v3, fi134)
                                                    }
                                                }
                                                _ = v188
                                            }
                                        case -1:
                                            v5 += (v5 | (v5 << ((v5) & 63)))
                                        default:
                                            _ = len(v1)
                                        }
                                        if ("foo bar" == "a") {
                                            var v197 int16 = v4
                                            v198 := v1["\n\t"]
                                            v199 := GenS25{F0: 1, F1: 225}
                                            v56 = v56
                                            var v200 bool = v3
                                            _ = v197
                                            _ = v198
                                            _ = v199
                                            _ = v200
                                        }
                                    }
                                    var v201 int64 = 674
                                    if false {
                                        var v202 string = ((("foo bar" + "foo bar") + "\n\t") + "a")
                                        _ = v202
                                    }
                                    _ = v201
                                default:
                                    v203 := GenS10{F0: 1, F1: 128}
                                    _ = v203
                                }
                            }
                        }
                        v204 := GenS15{F0: 181, F1: "\n\t"}
                        var v205 int = (fi133 & fi134)
                        _ = fi134
                        _ = v204
                        _ = v205
                    }
                    v1[""] = "foo bar"
                    switch 128 {
                    case -2147483648:
                        var v206 float64 = 0.0
                        _ = v206
                    case 2147483647:
                        v3 = ((v57 == v57) || v3)
                    }
                } else {
                    var v207 string = ((((((((((("hello" + "foo bar") + "\n\t") + "foo bar") + "\n\t") + "a") + ("a" + ("hello" + ""))) + ("" + ("hello" + "hello"))) + "hello") + "a") + "hello") + "a")
                    if true {
                        fmt.Println(fi133, fi133)
                    } else {
                        v208 := GenS26{F0: 100}
                        switch v207 {
                        case "":
                            var v209 int64 = -1
                            _ = v209
                        default:
                            if (255 != 128) {
                                fmt.Println(v5, v3)
                                var v210 int = fi133
                                if ((v5 >> ((v5) & 63)) == (v5 * (^(v5)))) {
                                    v211 := GenS27{F0: 1, F1: v56, F2: (v56 & (v6 / ((683) | 1)))}
                                    delete(v1, "a")
                                    _ = v211
                                } else {
                                    v212 := GenS2{F0: v207}
                                    if v3 {
                                        v1[v207] = (v207 + "")
                                        for fi213 := 0; fi213 < 16; fi213++ {
                                            var v214 bool = (1 != 65535)
                                            v1[""] = v207
                                            _ = fi213
                                            _ = v214
                                        }
                                    } else {
                                        v2 += 53
                                        if ((v5 & v5) == 255) {
                                            if v3 {
                                                v215 := GenS8{F0: ((fi133 >= (fi133 << ((65535) & 63))) || (!(!(float32(3.14159) >= -2.25))))}
                                                _ = v215
                                            } else {
                                                if v3 {
                                                    v216 := GenS2{F0: ""}
                                                    fmt.Println(v210)
                                                    v217 := GenS21{F0: (v4 * v4), F1: int64(v210), F2: 18446744073709551615}
                                                    _ = v216
                                                    _ = v217
                                                } else {
                                                    var v218 float64 = 1.5
                                                    _ = v218
                                                }
                                            }
                                            v2 = 2147483647
                                            var v219 int = fi133
                                            _ = v219
                                        }
                                        var v220 float64 = float64(((v57 + v57) + (v57 - -2.25)))
                                        _ = v220
                                    }
                                    v56 -= 4294967295
                                    _ = v212
                                }
                                _ = v210
                            } else {
                                v221 := []int{}
                                fmt.Println(v56)
                                _ = v221
                            }
                        }
                        _ = v208
                    }
                    _ = v207
                }
            default:
                var v222 uint32 = (v56 >> (((^(v5))) & 63))
                _ = v222
            }
            _ = fi133
        }
    case "":
        var v223 string = "hello"
        _ = v223
    case "\n\t":
        if false {
            var v224 float64 = float64(v57)
            for fi225 := 0; fi225 < 14; fi225++ {
                if (!(float64(0.0) <= v224)) {
                    if v3 {
                        if v3 {
                            if v3 {
                                var v226 uint32 = (v56 | v56)
                                if v3 {
                                    v227 := []int{(-(fi225)), (fi225 * fi225)}
                                    switch (fi225 << (((v5 / ((v5) | 1))) & 63)) {
                                    case 42:
                                        switch fi225 {
                                        case 9223372036854775807:
                                            v3 = (65535 != 4294967295)
                                        case -9223372036854775808:
                                            fmt.Println(v56, v5)
                                        case -222:
                                            v6 = (^(v6))
                                        default:
                                            var v228 float32 = 1e10
                                            _ = v228
                                        }
                                    }
                                    _ = v227
                                }
                                var v229 int16 = genFn0()
                                _ = v226
                                _ = v229
                            } else {
                                switch 127 {
                                case 32767:
                                    v230 := GenS19{F0: 9223372036854775807, F1: -571, F2: 128}
                                    _ = v230
                                default:
                                    v3 = v3
                                }
                                switch fi225 {
                                case 372:
                                    v231 := []int{(fi225 + -9223372036854775808), fi225, 4294967295}
                                    _ = v231
                                case 32768:
                                    v232 := GenS24{F0: 1, F1: v3, F2: -128}
                                    _ = v232
                                case 4294967295:
                                    var v233 uint64 = 4294967295
                                    _ = v233
                                }
                                if (true && (!v3)) {
                                    fmt.Println(v2, v3)
                                    v56 = (v6 >> (((^(v5))) & 63))
                                    v4 = int16(v5)
                                } else {
                                    v234 := []int{(fi225 & fi225), 128, 128, (fi225 | (-(fi225)))}
                                    v3 = false
                                    _ = v234
                                }
                            }
                            for fi235 := 0; fi235 < 13; fi235++ {
                                if (0 < 127) {
                                    _ = len(v1)
                                    v236 := v1[""]
                                    v2 = (v2 + 0)
                                    v224 -= 0.0
                                    if ((v3 && ((!(v6 <= v6)) && (1 < 18446744073709551615))) || v3) {
                                        if (36 > 192) {
                                            v56 += (^(v56))
                                            fmt.Println(v3, v2)
                                            var v237 uint16 = 257
                                            _ = v237
                                        }
                                        v238 := v1[v236]
                                        _ = v238
                                    }
                                    fmt.Println(v2, fi235, v4)
                                    _ = v236
                                }
                                var v239 float64 = 0.0
                                _ = fi235
                                _ = v239
                            }
                            var v240 int64 = 9223372036854775807
                            _ = v240
                        } else {
                            if v3 {
                                fmt.Println(v224, fi225, v5)
                            } else {
                                var v241 int = fi225
                                _ = v241
                            }
                            v242 := GenS28{F0: 42}
                            fmt.Println(v224, v2, v4)
                            _ = v242
                        }
                    } else {
                        for fi243 := 0; fi243 < 4; fi243++ {
                            switch 32767 {
                            case 1:
                                fmt.Println(v3)
                            default:
                                var v244 int16 = v4
                                _ = v244
                            }
                            var v245 uint16 = 836
                            _ = fi243
                            _ = v245
                        }
                    }
                    if false {
                        var v246 int32 = (v2 / ((v2) | 1))
                        var v247 int16 = v4
                        v5 -= v5
                        _ = v246
                        _ = v247
                    }
                    var v248 uint16 = 32767
                    _ = v248
                }
                _ = fi225
            }
            _ = v224
        } else {
            var v249 int32 = 42
            _ = v249
        }
    default:
        var v250 float64 = -2.25
        _ = v250
    }
    v56 += (v56 & (^(v6)))
    delete(v1, ("hello" + ""))
    v251 := GenS5{F0: -1, F1: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v251
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
