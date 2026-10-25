package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 uint32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 string; F1 bool; F2 int64 }
type GenS14 struct { F0 uint; F1 int32 }
type GenS15 struct { F0 int16; F1 uint; F2 uint }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 int8; F1 int64 }

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]bool)
    var v3 string = ""
    v4 := make(map[string]int)
    var v5 float64 = 1.5
    v6 := v2[1]
    v7 := GenS1{F0: 4294967295, F1: 42}
    if (!(false && false)) {
        v8 := GenS1{F0: 65535, F1: 9223372036854775807}
        if v6 {
            if (128 < 0) {
                if v6 {
                    v9 := []int{}
                    if ((false || (!(4 > 2))) || (!(((float32(v5) - 1e10) < 0.0) || false))) {
                        v10 := GenS2{F0: -934, F1: false}
                        _ = v10
                    } else {
                        var v11 uint8 = 127
                        fmt.Println(v5, v11)
                        var v12 int64 = -41
                        _ = v11
                        _ = v12
                    }
                    _ = v9
                }
                v1 = (v1 ^ v1)
                v13 := GenS1{F0: 255, F1: -9223372036854775808}
                _ = v13
            }
        }
        switch -1 {
        case 805:
            var v14 float64 = v5
            _ = v14
        case -997:
            for fi15 := 0; fi15 < 16; fi15++ {
                if (778 < 0) {
                    v1 = 452
                    v16 := GenS3{F0: 857, F1: 42}
                    _ = v16
                } else {
                    var v17 int16 = -1
                    v18 := GenS4{F0: (-32768 != (v17 + (v17 ^ (^(v17))))), F1: 0, F2: 2}
                    var v19 uint64 = (^(v1))
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v1 -= v1
                for fi20 := 0; fi20 < 14; fi20++ {
                    if ((true || v6) && v6) {
                        var v21 int = 1
                        v22 := GenS3{F0: 0, F1: 9223372036854775807}
                        _ = v21
                        _ = v22
                    }
                    _ = fi20
                }
                _ = fi15
            }
        case 128:
            delete(v2, 2)
        }
        _ = v8
    } else {
        if (34 < 1) {
            v23 := GenS5{F0: 4, F1: 0}
            v1 = 0
            var v24 int = 0
            _ = v23
            _ = v24
        } else {
            v6 = (((65535 != 255) && v6) && ((32767 < 1) || true))
            var v25 string = "a"
            _ = v25
        }
        v26 := GenS3{F0: -9223372036854775808, F1: 0}
        if true {
            v5 -= v5
            for fi27 := 0; fi27 < 2; fi27++ {
                var v28 float32 = -2.25
                _ = fi27
                _ = v28
            }
        }
        _ = v26
    }
    v29 := GenS6{F0: 1, F1: 1}
    for fi30 := 0; fi30 < 19; fi30++ {
        v5 = v5
        _ = fi30
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        v32 := []int{fi31}
        _ = fi31
        _ = v32
    }
    var v33 float64 = float64(0.0)
    if false {
        if (-2.25 <= float32(float64(1e10))) {
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int = 1
                _ = fi34
                _ = v35
            }
            var v36 uint8 = 128
            var v37 uint64 = 1
            _ = v36
            _ = v37
        }
        v33 = (v5 + v33)
        var v38 uint64 = 2
        _ = v38
    } else {
        v39 := GenS4{F0: (v5 > v5), F1: 2147483647, F2: 639}
        var v40 string = ""
        _ = v39
        _ = v40
    }
    if (1 <= 42) {
        var v41 uint = 1
        _ = v41
    } else {
        fmt.Println(v33, v6, v6)
    }
    v6 = v6
    fmt.Println(v3, v3, v3)
    v42 := []int{441, 0}
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    v43 := GenS1{F0: 42, F1: -1}
    if (!((v6 || v6) || ((1 < -32768) || v6))) {
        v5 -= (((1e10 - ((v33 - v33) / (v33 - float64(float32(((1.5 + 0.0) * v5)))))) * ((-2.25 / 1e10) * ((v5 / v5) * (v33 * v33)))) / float64(1e10))
        v33 = 1.5
        var v44 int32 = -2147483648
        _ = v44
    }
    switch ((v3 + v3) + ("" + v3)) {
    case "foo bar":
        for fi45 := 0; fi45 < 6; fi45++ {
            if (1 <= 17) {
                fmt.Println(v1, v6, v3)
            } else {
                for fi46 := 0; fi46 < 20; fi46++ {
                    if ((-612 > 42) && (2147483647 == -652)) {
                        switch "\n\t" {
                        case "hello":
                            if (344 > 2) {
                                v6 = (!true)
                                if v6 {
                                    _ = len(v42)
                                    var v47 string = ((((v3 + v3) + (v3 + "")) + (v3 + (v3 + v3))) + "hello")
                                    var v48 uint = 255
                                    _ = v47
                                    _ = v48
                                } else {
                                    v5 += (v5 - (1.5 / float64(1e10)))
                                    switch (fi46 * 239) {
                                    case 2147483647:
                                        v5 -= (v33 * v33)
                                    case -1:
                                        var v49 int64 = 104
                                        _ = v49
                                    case 431:
                                        if (0 < 550) {
                                            v50 := GenS3{F0: 9223372036854775807, F1: 672}
                                            _ = v50
                                        } else {
                                            fmt.Println(v6)
                                            v3 = v3
                                        }
                                    default:
                                        v3 = "a"
                                    }
                                    var v51 float32 = 1.5
                                    _ = v51
                                }
                                v52 := GenS7{F0: 2}
                                _ = v52
                            } else {
                                var v53 int16 = -896
                                v54 := GenS5{F0: -1, F1: 0}
                                v1 = v1
                                _ = v53
                                _ = v54
                            }
                            switch (("hello" + v3) + v3) {
                            case "foo bar":
                                v42 = append(v42, fi45)
                            case "hello":
                                if v6 {
                                    v55 := GenS6{F0: int64(fi46), F1: 2147483647}
                                    _ = len(v2)
                                    _ = v55
                                }
                            default:
                                if true {
                                    var v56 bool = ((fi45 ^ -513) > (fi46 ^ (fi46 & fi45)))
                                    v5 -= (v33 / 1.5)
                                    v57 := v4[v3]
                                    if true {
                                        var v58 uint32 = 65535
                                        if v56 {
                                            v5 += v5
                                            v1 -= 998
                                            v59 := GenS8{F0: 127, F1: v6, F2: v58}
                                            _ = v59
                                        } else {
                                            switch v3 {
                                            case "":
                                                var v60 float64 = 3.14159
                                                _ = v60
                                            }
                                            switch v57 {
                                            case 249:
                                                var v61 uint32 = v58
                                                _ = v61
                                            case 127:
                                                var v62 int64 = -1
                                                v58 = 457
                                                _ = v62
                                            case 9223372036854775807:
                                                switch (^(fi46)) {
                                                case 127:
                                                    var v63 string = ((v3 + ((((v3 + (("\n\t" + "foo bar") + ((((v3 + "") + v3) + (("foo bar" + "") + "foo bar")) + v3))) + (("hello" + "") + v3)) + (v3 + v3)) + v3)) + v3)
                                                    _ = v63
                                                }
                                            }
                                        }
                                        _ = v58
                                    } else {
                                        var v64 int32 = 997
                                        _ = v64
                                    }
                                    _ = v56
                                    _ = v57
                                }
                            }
                        case "\n\t":
                            v65 := GenS1{F0: 170, F1: 212}
                            _ = v65
                        }
                    }
                    v6 = false
                    switch (^(fi45)) {
                    case 32768:
                        fmt.Println(v6, v1, v3)
                    case -511:
                        if (((!v6) && (243 <= 255)) || true) {
                            var v66 float32 = (0.0 - (((0.1 / (3.14159 - ((1e10 * float32(((v5 * ((0.0 + ((3.14159 / (1e10 + (3.14159 + 0.1))) - ((v33 / (v33 + v33)) * v33))) + v33)) + (v33 / v33)))) - (float32((3.14159 * (((v33 - v33) * (v5 + v5)) / v33))) / (-2.25 / 1.0))))) - 0.0) / 1.0))
                            switch fi45 {
                            case -1:
                                v5 += 3.14159
                            case -675:
                                v67 := GenS9{F0: -9223372036854775808, F1: 137}
                                _ = v67
                            case 32768:
                                v68 := GenS1{F0: 172, F1: 9223372036854775807}
                                _ = v68
                            }
                            _ = len(v2)
                            v69 := v4[((v3 + v3) + v3)]
                            var v70 int16 = 1
                            _ = v66
                            _ = v69
                            _ = v70
                        } else {
                            if (65535 <= 4294967295) {
                                switch v3 {
                                case "":
                                    if (((216 <= 128) && v6) && (4294967295 == 65535)) {
                                        var v71 int16 = -32768
                                        delete(v4, "foo bar")
                                        _ = v71
                                    } else {
                                        switch -148 {
                                        case -9223372036854775808:
                                            fmt.Println(v5, v1, v3)
                                        case -978:
                                            var v72 float64 = (1e10 + v33)
                                            _ = v72
                                        }
                                        var v73 uint64 = v1
                                        _ = v73
                                    }
                                default:
                                    fmt.Println(v3, fi46)
                                }
                                v33 -= (v33 + v33)
                                switch ("\n\t" + "") {
                                case "a":
                                    v74 := GenS7{F0: 1}
                                    _ = v74
                                case "foo bar":
                                    if true {
                                        v1 -= (^(v1))
                                        var v75 int = (^(fi46))
                                        _ = v75
                                    } else {
                                        delete(v2, 2147483648)
                                    }
                                case "hello":
                                    var v76 int32 = 42
                                    _ = v76
                                }
                            } else {
                                v1 = (v1 | (^(v1)))
                            }
                            v77 := GenS9{F0: -1, F1: 255}
                            v3 = v3
                            _ = v77
                        }
                    }
                    _ = fi46
                }
                var v78 int64 = -1
                var v79 int8 = 42
                _ = v78
                _ = v79
            }
            v42 = append(v42, fi45)
            var v80 uint64 = (v1 ^ (^(v1)))
            _ = fi45
            _ = v80
        }
    case "hello":
        var v81 int8 = -96
        _ = v81
    default:
        for fi82 := 0; fi82 < 18; fi82++ {
            if (1000 <= 1000) {
                if (-1 > 0) {
                    v83 := GenS4{F0: v6, F1: 2147483648, F2: 2}
                    if v6 {
                        var v84 int64 = int64(fi82)
                        var v85 float64 = 1.5
                        if len(v42) > 0 {
                            v42[0] = -654
                        }
                        _ = v84
                        _ = v85
                    } else {
                        var v86 int16 = -498
                        v87 := GenS7{F0: 128}
                        if (true && (v33 <= (0.1 * v33))) {
                            v88 := GenS3{F0: -9223372036854775808, F1: 275}
                            _ = len(v42)
                            v6 = true
                            _ = v88
                        } else {
                            if true {
                                fmt.Println(v33, v5, v86)
                            } else {
                                fmt.Println(v3, fi82)
                                switch "\n\t" {
                                case "":
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v5, v86)
                                    case "a":
                                        var v89 float64 = 0.1
                                        _ = v89
                                    case "":
                                        v33 = (v33 + v5)
                                    default:
                                        var v90 string = "hello"
                                        _ = v90
                                    }
                                case "a":
                                    for fi91 := 0; fi91 < 16; fi91++ {
                                        fmt.Println(fi91)
                                        var v92 int32 = 1
                                        var v93 uint16 = 2
                                        if ((!(-112 >= 0)) && (!v6)) {
                                            if (false || (1e10 >= -2.25)) {
                                                v94 := GenS10{F0: "foo bar"}
                                                v95 := GenS11{F0: 2147483647, F1: 2147483648}
                                                _ = v94
                                                _ = v95
                                            } else {
                                                if (v6 && ((v1 != v1) && ((v33 + -2.25) <= float64(1e10)))) {
                                                    if len(v42) >= 2 {
                                                        v42 = v42[1:len(v42)]
                                                    }
                                                    if v6 {
                                                        fmt.Println(v3, fi82)
                                                        switch fi91 {
                                                        case 9223372036854775807:
                                                            switch (fi82 | (fi91 & (fi91 ^ (^(fi91))))) {
                                                            case 2147483647:
                                                                v96 := GenS12{F0: (v93 / ((42) | 1))}
                                                                _ = v96
                                                            case 9223372036854775807:
                                                                fmt.Println(v3, v92, v93)
                                                            default:
                                                                var v97 uint16 = v93
                                                                _ = v97
                                                            }
                                                        default:
                                                            var v98 uint32 = 2147483647
                                                            _ = v98
                                                        }
                                                        if v6 {
                                                            var v99 int64 = -113
                                                            _ = v99
                                                        }
                                                    } else {
                                                        if v6 {
                                                            if (0.1 != v5) {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v100 := GenS13{F0: (((v3 + (v3 + (v3 + ""))) + v3) + ((v3 + "foo bar") + v3)), F1: ((!((fi82 < (fi82 & (fi91 & (fi91 - 128)))) || (255 < 65535))) || (!v6)), F2: -9223372036854775808}
                                                                    _ = v100
                                                                default:
                                                                    v42 = append(v42, fi91)
                                                                }
                                                                var v101 int32 = (v92 & 42)
                                                                v5 = v5
                                                                delete(v4, v3)
                                                                if ((7 < 1) && ((^(v1)) < v1)) {
                                                                    switch ("hello" + v3) {
                                                                    case "\n\t":
                                                                        v102 := GenS8{F0: 67, F1: v6, F2: 356}
                                                                        _ = v102
                                                                    case "a":
                                                                        v6 = true
                                                                    case "hello":
                                                                        if true {
                                                                            if (fi91 <= fi82) {
                                                                                v1 = (^(v1))
                                                                                var v103 int64 = -428
                                                                                switch (-(fi82)) {
                                                                                case -960:
                                                                                    v42 = append(v42, fi91)
                                                                                case -2147483648:
                                                                                    if v6 {
                                                                                        v104 := GenS14{F0: uint(v93), F1: (v101 / ((0) | 1))}
                                                                                        _ = v104
                                                                                    }
                                                                                }
                                                                                fmt.Println(v86)
                                                                                v103 = -9223372036854775808
                                                                                _ = v103
                                                                            }
                                                                            delete(v4, v3)
                                                                        } else {
                                                                            v3 = v3
                                                                            fmt.Println(v3, fi82, v6)
                                                                        }
                                                                    default:
                                                                        switch -548 {
                                                                        case -1:
                                                                            _ = len(v42)
                                                                        case 32767:
                                                                            v92 += int32(fi91)
                                                                        case 4294967295:
                                                                            var v105 int = fi82
                                                                            _ = v105
                                                                        }
                                                                    }
                                                                    delete(v2, 596)
                                                                } else {
                                                                    if v6 {
                                                                        if (-76 <= -1) {
                                                                            if v6 {
                                                                                switch v3 {
                                                                                case "\n\t":
                                                                                    v101 = v101
                                                                                }
                                                                                fmt.Println(fi82, v92)
                                                                                v106 := GenS15{F0: v86, F1: 255, F2: 2}
                                                                                _ = v106
                                                                            } else {
                                                                                var v107 bool = v6
                                                                                _ = v107
                                                                            }
                                                                        }
                                                                        v93 += v93
                                                                        var v108 uint32 = 659
                                                                        _ = v108
                                                                    }
                                                                }
                                                                _ = v101
                                                            } else {
                                                                var v109 float64 = v5
                                                                _ = len(v42)
                                                                _ = v109
                                                            }
                                                            if (false || (v3 != v3)) {
                                                                var v110 string = (v3 + v3)
                                                                _ = v110
                                                            }
                                                            var v111 int = 21
                                                            _ = v111
                                                        }
                                                        if (v6 && false) {
                                                            v1 = v1
                                                            fmt.Println(v1)
                                                        }
                                                        var v112 uint = 255
                                                        v113 := GenS16{F0: v93}
                                                        _ = len(v42)
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                }
                                                v92 -= int32(fi82)
                                                var v114 uint8 = 127
                                                _ = v114
                                            }
                                            var v115 uint = 65535
                                            if (-1 <= -1) {
                                                var v116 int64 = -9223372036854775808
                                                var v117 string = ((v3 + v3) + v3)
                                                _ = v116
                                                _ = v117
                                            }
                                            _ = v115
                                        } else {
                                            var v118 int16 = (v86 << ((18446744073709551615) & 63))
                                            _ = len(v2)
                                            var v119 float32 = 1.5
                                            fmt.Println(v86, v5, fi91)
                                            var v120 bool = v6
                                            if false {
                                                if v120 {
                                                    var v121 int16 = -686
                                                    v93 -= 21
                                                    v122 := GenS3{F0: -9223372036854775808, F1: -1}
                                                    var v123 uint8 = 127
                                                    _ = v121
                                                    _ = v122
                                                    _ = v123
                                                }
                                            }
                                            _ = v118
                                            _ = v119
                                            _ = v120
                                        }
                                        if v6 {
                                            fmt.Println(v1, v6, v86)
                                        }
                                        _ = fi91
                                        _ = v92
                                        _ = v93
                                    }
                                default:
                                    for fi124 := 0; fi124 < 18; fi124++ {
                                        fmt.Println(v86)
                                        v1 -= 116
                                        var v125 bool = (42 > 192)
                                        _ = fi124
                                        _ = v125
                                    }
                                }
                            }
                        }
                        _ = v86
                        _ = v87
                    }
                    if v6 {
                        v5 += ((1e10 + float64(-2.25)) - 1.5)
                        _ = len(v42)
                    }
                    _ = v83
                } else {
                    v33 += (v5 * ((v5 - v33) / ((0.1 / v5) - v5)))
                    switch fi82 {
                    case 42:
                        var v126 float32 = 0.1
                        _ = v126
                    case -2147483648:
                        if v6 {
                            for fi127 := 0; fi127 < 4; fi127++ {
                                _ = len(v42)
                                v128 := v4[v3]
                                _ = fi127
                                _ = v128
                            }
                        }
                    default:
                        switch v3 {
                        case "foo bar":
                            v3 = v3
                        case "\n\t":
                            switch fi82 {
                            case -866:
                                if ((!v6) || true) {
                                    for fi129 := 0; fi129 < 8; fi129++ {
                                        if (!(!(!((false || (32767 > 1)) || (!v6))))) {
                                            v130 := GenS15{F0: 295, F1: 1, F2: 222}
                                            fmt.Println(v3, fi129, fi129)
                                            v131 := GenS6{F0: 1, F1: 42}
                                            _ = v130
                                            _ = v131
                                        } else {
                                            v132 := GenS14{F0: 795, F1: 2147483647}
                                            var v133 uint = 663
                                            _ = v132
                                            _ = v133
                                        }
                                        v1 = 4294967295
                                        _ = fi129
                                    }
                                    if true {
                                        v3 = "hello"
                                        var v134 float32 = 3.14159
                                        v134 -= -2.25
                                        _ = v134
                                    } else {
                                        var v135 int16 = 32767
                                        if len(v42) > 0 {
                                            v42[0] = (fi82 | fi82)
                                        }
                                        _ = v135
                                    }
                                    var v136 uint64 = v1
                                    _ = v136
                                }
                            case 313:
                                v33 -= (-2.25 + v33)
                            case 9223372036854775807:
                                var v137 int = fi82
                                _ = v137
                            }
                        case "a":
                            var v138 uint8 = 181
                            _ = v138
                        }
                    }
                    v5 = -2.25
                }
                _ = len(v42)
                var v139 uint32 = 42
                _ = v139
            } else {
                for fi140 := 0; fi140 < 18; fi140++ {
                    v141 := GenS1{F0: 4294967295, F1: -1}
                    if v6 {
                        var v142 uint64 = 567
                        if true {
                            fmt.Println(v33, v33, v33)
                        }
                        _ = v142
                    } else {
                        v33 = v33
                        if (v3 < v3) {
                            fmt.Println(fi140, v3, fi140)
                            if (100 > 42) {
                                fmt.Println(v33)
                                v42 = append(v42, fi140)
                            } else {
                                if v6 {
                                    if false {
                                        var v143 float64 = v33
                                        if (fi82 == (fi82 - (fi82 * (fi82 >> ((1) & 63))))) {
                                            if len(v42) > 0 {
                                                v42[0] = (-(fi82))
                                            }
                                            fmt.Println(v3, v1)
                                        } else {
                                            v144 := GenS8{F0: 184, F1: false, F2: 0}
                                            _ = v144
                                        }
                                        switch v3 {
                                        case "a":
                                            v145 := GenS11{F0: 840, F1: 65535}
                                            _ = v145
                                        case "hello":
                                            var v146 int32 = -2147483648
                                            _ = v146
                                        }
                                        _ = v143
                                    }
                                } else {
                                    switch 201 {
                                    case -377:
                                        v2[65535] = (-124 >= 1000)
                                    }
                                    v5 -= (0.1 + float64((0.1 + 0.0)))
                                    var v147 float32 = ((1e10 / 1e10) * float32(1.5))
                                    _ = v147
                                }
                            }
                            var v148 int64 = -1
                            _ = v148
                        } else {
                            v3 = ""
                        }
                    }
                    v3 = ("hello" + v3)
                    v149 := GenS7{F0: 109}
                    switch fi140 {
                    case 300:
                        if (9223372036854775807 >= 1) {
                            var v150 int8 = 42
                            v151 := GenS17{F0: (v150 >> ((132) & 63)), F1: 1}
                            if len(v42) >= 2 {
                                v42 = v42[1:len(v42)]
                            }
                            _ = v150
                            _ = v151
                        }
                    case 9223372036854775807:
                        var v152 int64 = 9223372036854775807
                        _ = v152
                    }
                    _ = fi140
                    _ = v141
                    _ = v149
                }
                var v153 int64 = -216
                switch v3 {
                case "hello":
                    var v154 int8 = 45
                    _ = v154
                case "\n\t":
                    v155 := GenS17{F0: 42, F1: -871}
                    _ = v155
                case "":
                    var v156 float32 = -2.25
                    _ = v156
                }
                _ = v153
            }
            fmt.Println(fi82, fi82, v33)
            switch ((v3 + v3) + ("foo bar" + ((("" + v3) + v3) + ((((v3 + ("a" + v3)) + (v3 + (v3 + ("" + (((((v3 + (("a" + (v3 + "a")) + ("\n\t" + ""))) + "") + v3) + v3) + v3))))) + v3) + (("foo bar" + v3) + (v3 + v3)))))) {
            case "foo bar":
                _ = len(v4)
            case "\n\t":
                switch (^(fi82)) {
                case 32768:
                    var v157 int32 = 0
                    _ = v157
                default:
                    for fi158 := 0; fi158 < 18; fi158++ {
                        fmt.Println(v6)
                        _ = fi158
                    }
                }
            default:
                fmt.Println(v3, v33)
            }
            _ = fi82
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
