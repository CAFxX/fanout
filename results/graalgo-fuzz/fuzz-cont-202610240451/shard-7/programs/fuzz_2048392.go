package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8; F1 uint64; F2 int16 }
type GenS13 struct { F0 string; F1 int8; F2 uint }
type GenS14 struct { F0 int; F1 int8 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int; F1 int8 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0(p0 string, p1 int, p2 int32) uint64 {
    v6 := GenS1{F0: p2}
    fmt.Println(p0, p0, p1)
    var v7 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1() uint {
    var v8 string = "a"
    if false {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 bool = (!true)
            _ = fi9
            _ = v10
        }
    } else {
        switch 42 {
        case 9223372036854775807:
            v8 = ((("a" + v8) + v8) + v8)
        case 32768:
            switch (("" + v8) + (v8 + v8)) {
            case "hello":
                switch 184 {
                case 2147483647:
                    for fi11 := 0; fi11 < 17; fi11++ {
                        if false {
                            fmt.Println(v8, fi11)
                        }
                        _ = fi11
                    }
                case 86:
                    v8 = ("foo bar" + (("\n\t" + v8) + v8))
                case 1:
                    v8 = v8
                default:
                    v8 = v8
                }
            case "":
                v12 := []int{-206}
                _ = v12
            case "\n\t":
                fmt.Println(v8)
            }
        case 1:
            var v13 uint8 = 128
            _ = v13
        default:
            v14 := []int{974, 373, -349}
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint8 = 128
        v17 := GenS2{F0: 127, F1: -9223372036854775808}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 bool = (((true && false) && (675 == 62)) && true)
    if v18 {
        v19 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807}
        if v18 {
            v18 = v18
        } else {
            v20 := []int{-333}
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 uint16 = 138
                _ = fi21
                _ = v22
            }
            _ = v20
        }
        v8 = "foo bar"
        _ = v19
    } else {
        v8 = v8
        switch -2147483648 {
        case 42:
            v23 := GenS3{F0: 972, F1: ""}
            _ = v23
        default:
            var v24 uint64 = 2
            _ = v24
        }
        v25 := []int{-2147483648, -266}
        _ = v25
    }
    _ = v18
    _ = v8
    return 308
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 255
    var v3 bool = (2 < 2)
    var v4 int64 = -1
    var v5 float64 = ((0.1 / 1e10) / -2.25)
    v2 = genFn1()
    v26 := GenS2{F0: 32768, F1: v4}
    v2 -= genFn1()
    for fi27 := 0; fi27 < 7; fi27++ {
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS4{F0: 2147483648}
            v2 += (v2 % ((genFn1()) | 1))
            var v30 int = 521
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 int32 = -1
        _ = fi27
        _ = v31
    }
    delete(v1, "\n\t")
    v32 := GenS5{F0: 1000}
    v33 := []int{718, 25, -1, 128}
    if true {
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 uint = (v2 / ((v2) | 1))
            _ = fi34
            _ = v35
        }
        var v36 int8 = 0
        _ = v36
    }
    if true {
        v5 += (v5 / 0.1)
        if (true && true) {
            fmt.Println(v4, v4, v3)
        }
        if v3 {
            if (0 >= -32768) {
                if true {
                    var v37 uint64 = 2
                    var v38 float64 = v5
                    _ = v37
                    _ = v38
                } else {
                    delete(v1, "foo bar")
                    for fi39 := 0; fi39 < 19; fi39++ {
                        if v3 {
                            for fi40 := 0; fi40 < 13; fi40++ {
                                v3 = true
                                switch "\n\t" {
                                case "hello":
                                    fmt.Println(v5, fi40, v2)
                                }
                                switch -9223372036854775808 {
                                case 1:
                                    delete(v1, (("foo bar" + "hello") + "hello"))
                                default:
                                    if true {
                                        var v41 uint64 = 115
                                        _ = v41
                                    } else {
                                        v5 += v5
                                        if (492 < -117) {
                                            if len(v33) > 0 {
                                                v33[0] = (fi39 & -9223372036854775808)
                                            }
                                            var v42 float32 = (1e10 - -2.25)
                                            if v3 {
                                                var v43 bool = v3
                                                v43 = v3
                                                _ = v43
                                            }
                                            v5 -= (v5 / v5)
                                            var v44 bool = (((!(("hello" + "") == "a")) || false) || v3)
                                            _ = v42
                                            _ = v44
                                        } else {
                                            _ = len(v33)
                                        }
                                    }
                                    v3 = v3
                                }
                                if true {
                                    v5 = v5
                                    v5 = 1.5
                                }
                                _ = fi40
                            }
                        } else {
                            if v3 {
                                var v45 uint = (v2 - v2)
                                v46 := GenS1{F0: 815}
                                _ = v45
                                _ = v46
                            } else {
                                v5 += (v5 + (-2.25 - v5))
                            }
                            v47 := GenS5{F0: 905}
                            fmt.Println(v4, v4, v3)
                            _ = v47
                        }
                        _ = fi39
                    }
                }
                var v48 int64 = v4
                v5 = v5
                _ = v48
            }
            v3 = (-2.25 <= (1.5 - (3.14159 - 0.0)))
        } else {
            if (!(2 < 162)) {
                v5 -= (3.14159 + v5)
                v49 := GenS4{F0: 549}
                v2 += 1
                _ = v49
            } else {
                var v50 string = "hello"
                var v51 uint16 = 861
                _ = v50
                _ = v51
            }
            fmt.Println(v4, v4)
            if (false && false) {
                var v52 float64 = v5
                v2 -= 18446744073709551615
                _ = v52
            }
        }
    }
    if v3 {
        v53 := v1["hello"]
        fmt.Println(v53)
        switch -2147483648 {
        case 42:
            v5 = v5
        case -1:
            if (!v3) {
                var v54 uint16 = 2
                if ((float32(-2.25) < 3.14159) && (!v3)) {
                    if v3 {
                        v5 = v5
                        if ((2 == genFn0("foo bar", -299, 2147483647)) || false) {
                            for fi55 := 0; fi55 < 12; fi55++ {
                                if (v5 <= v5) {
                                    if ((^(v2)) < v2) {
                                        if v3 {
                                            if false {
                                                if (-32768 > 1) {
                                                    if true {
                                                        v56 := v1[""]
                                                        v54 -= 32767
                                                        v57 := GenS6{F0: "\n\t", F1: 42}
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    for fi58 := 0; fi58 < 15; fi58++ {
                                                        v1["hello"] = v53
                                                        v5 = v5
                                                        var v59 int32 = -320
                                                        _ = len(v33)
                                                        v3 = (32767 <= 1000)
                                                        fmt.Println(v53)
                                                        _ = fi58
                                                        _ = v59
                                                    }
                                                    _ = len(v1)
                                                }
                                                v4 = (v4 * v4)
                                            } else {
                                                fmt.Println(v2)
                                            }
                                        } else {
                                            switch ("hello" + ("" + ("a" + ("foo bar" + "hello")))) {
                                            case "":
                                                fmt.Println(v2)
                                            case "foo bar":
                                                for fi60 := 0; fi60 < 13; fi60++ {
                                                    var v61 uint8 = 127
                                                    var v62 int16 = -1
                                                    var v63 uint = 52
                                                    _ = fi60
                                                    _ = v61
                                                    _ = v62
                                                    _ = v63
                                                }
                                            default:
                                                fmt.Println(v3, fi55, fi55)
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                var v64 int64 = (v4 | v4)
                                                _ = v64
                                            case "":
                                                if v3 {
                                                    switch fi55 {
                                                    case 9223372036854775807:
                                                        if (genFn0("hello", (fi55 * 825), 0) == genFn0("\n\t", int(v4), 2147483647)) {
                                                            v65 := GenS3{F0: fi55, F1: (("foo bar" + ("a" + "hello")) + ("" + ("a" + "foo bar")))}
                                                            v4 = (v4 & (v4 % (((v4 << (((^(v2))) & 63))) | 1)))
                                                            var v66 int8 = -103
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    case 128:
                                                        fmt.Println(fi55, v3)
                                                    case -338:
                                                        for fi67 := 0; fi67 < 9; fi67++ {
                                                            fmt.Println(fi55, fi55)
                                                            v68 := GenS7{F0: -9223372036854775808, F1: fi67}
                                                            _ = fi67
                                                            _ = v68
                                                        }
                                                    default:
                                                        if (v3 || true) {
                                                            v69 := GenS7{F0: -861, F1: 4294967295}
                                                            v1["hello"] = 1
                                                            _ = v69
                                                        }
                                                    }
                                                    fmt.Println(v54, v3)
                                                }
                                            default:
                                                v70 := GenS6{F0: "", F1: 813}
                                                _ = v70
                                            }
                                        }
                                        v4 += (v4 ^ (-(v4)))
                                        switch (("hello" + ("\n\t" + "\n\t")) + (("a" + "foo bar") + (((("hello" + "") + ("a" + "\n\t")) + "a") + ""))) {
                                        case "":
                                            v71 := GenS7{F0: -1, F1: fi55}
                                            _ = v71
                                        default:
                                            v4 -= 9223372036854775807
                                        }
                                    }
                                    if true {
                                        if v3 {
                                            if (("\n\t" + "") > "hello") {
                                                switch 128 {
                                                case 9223372036854775807:
                                                    v53 = (v53 % ((v53) | 1))
                                                case -370:
                                                    var v72 int32 = -643
                                                    _ = v72
                                                default:
                                                    if (!true) {
                                                        var v73 uint = genFn1()
                                                        _ = v73
                                                    }
                                                }
                                                _ = len(v1)
                                                v54 -= v54
                                            }
                                        } else {
                                            v2 += (v2 + v2)
                                            var v74 bool = true
                                            var v75 uint64 = 42
                                            _ = v74
                                            _ = v75
                                        }
                                    } else {
                                        fmt.Println(v3)
                                        switch fi55 {
                                        case 32767:
                                            for fi76 := 0; fi76 < 2; fi76++ {
                                                switch (fi76 - -2147483648) {
                                                case 4294967295:
                                                    v77 := GenS7{F0: (^(v4)), F1: 127}
                                                    _ = v77
                                                case -2147483648:
                                                    var v78 float32 = 0.1
                                                    _ = v78
                                                }
                                                if v3 {
                                                    var v79 uint8 = 150
                                                    _ = v79
                                                }
                                                fmt.Println(v2)
                                                _ = fi76
                                            }
                                        }
                                        var v80 uint8 = 135
                                        _ = v80
                                    }
                                } else {
                                    var v81 uint32 = (^(v53))
                                    _ = v81
                                }
                                var v82 int16 = 0
                                _ = fi55
                                _ = v82
                            }
                        } else {
                            v1["hello"] = v53
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            var v83 uint = v2
                            _ = v83
                        }
                        switch "a" {
                        case "":
                            switch 128 {
                            case 2147483647:
                                v84 := v1[("hello" + ("foo bar" + ("" + "\n\t")))]
                                _ = v84
                            case 649:
                                if (0 >= -2147483648) {
                                    if (v3 || v3) {
                                        if (42 >= -32768) {
                                            for fi85 := 0; fi85 < 18; fi85++ {
                                                v86 := GenS8{F0: 699, F1: int8(v54)}
                                                v87 := GenS5{F0: -1}
                                                var v88 bool = ((v3 || false) && (((3.14159 - 0.0) != float32(v5)) || v3))
                                                _ = fi85
                                                _ = v86
                                                _ = v87
                                                _ = v88
                                            }
                                            v2 += v2
                                        } else {
                                            for fi89 := 0; fi89 < 13; fi89++ {
                                                v90 := v1[("foo bar" + (("foo bar" + ("\n\t" + "a")) + ""))]
                                                var v91 uint = (^(v2))
                                                var v92 bool = v3
                                                _ = fi89
                                                _ = v90
                                                _ = v91
                                                _ = v92
                                            }
                                            var v93 float32 = 3.14159
                                            for fi94 := 0; fi94 < 2; fi94++ {
                                                var v95 string = "\n\t"
                                                _ = fi94
                                                _ = v95
                                            }
                                            _ = v93
                                        }
                                        v33 = append(v33, 0)
                                        v96 := GenS9{F0: 564}
                                        _ = v96
                                    } else {
                                        v53 = (v53 << ((v2) & 63))
                                        var v97 int = 2147483647
                                        _ = len(v33)
                                        _ = v97
                                    }
                                    if (667 == 32767) {
                                        if len(v33) > 0 {
                                            v33[0] = 9223372036854775807
                                        }
                                        var v98 uint32 = (v53 * (v53 - 4294967295))
                                        v98 -= 0
                                        _ = v98
                                    } else {
                                        v53 = v53
                                        var v99 int32 = -1
                                        for fi100 := 0; fi100 < 10; fi100++ {
                                            var v101 int64 = v4
                                            if (42 != uint8(v54)) {
                                                var v102 string = "foo bar"
                                                v101 += v101
                                                if (false && v3) {
                                                    switch 32768 {
                                                    case 1:
                                                        var v103 int8 = 0
                                                        _ = v103
                                                    default:
                                                        switch fi100 {
                                                        case 127:
                                                            var v104 uint32 = v53
                                                            _ = v104
                                                        case 42:
                                                            v33 = append(v33, fi100)
                                                        case -854:
                                                            var v105 int16 = 1000
                                                            _ = v105
                                                        default:
                                                            if v3 {
                                                                for fi106 := 0; fi106 < 7; fi106++ {
                                                                    v107 := v1[v102]
                                                                    if len(v33) >= 2 {
                                                                        v33 = v33[1:len(v33)]
                                                                    }
                                                                    _ = fi106
                                                                    _ = v107
                                                                }
                                                            } else {
                                                                var v108 int = 128
                                                                _ = v108
                                                            }
                                                        }
                                                    }
                                                    var v109 float32 = ((1.5 + 0.1) * 3.14159)
                                                    var v110 int64 = (-(v101))
                                                    _ = v109
                                                    _ = v110
                                                }
                                                _ = v102
                                            } else {
                                                var v111 uint16 = (v54 | v54)
                                                var v112 string = "hello"
                                                fmt.Println(v53, v54)
                                                _ = v111
                                                _ = v112
                                            }
                                            if (972 != 42) {
                                                for fi113 := 0; fi113 < 11; fi113++ {
                                                    switch "hello" {
                                                    case "hello":
                                                        if (v101 >= (v101 >> ((v2) & 63))) {
                                                            var v114 int8 = -1
                                                            if false {
                                                                _ = len(v33)
                                                            }
                                                            _ = v114
                                                        } else {
                                                            v54 = 605
                                                        }
                                                    case "foo bar":
                                                        switch (^(fi113)) {
                                                        case -9223372036854775808:
                                                            var v115 int16 = 441
                                                            _ = v115
                                                        }
                                                    case "\n\t":
                                                        var v116 float64 = (v5 * float64(0.0))
                                                        _ = v116
                                                    }
                                                    if (true || v3) {
                                                        v101 = v101
                                                        var v117 float64 = 3.14159
                                                        _ = v117
                                                    } else {
                                                        switch fi113 {
                                                        case -878:
                                                            v118 := GenS6{F0: "foo bar", F1: (v99 | -2147483648)}
                                                            _ = v118
                                                        case 182:
                                                            if ((v101 != int64(v54)) || (true || false)) {
                                                                delete(v1, "hello")
                                                                var v119 int8 = 127
                                                                _ = v119
                                                            } else {
                                                                var v120 int = (fi113 | fi113)
                                                                var v121 string = ""
                                                                v122 := GenS1{F0: (v99 << (((v2 >> ((255) & 63))) & 63))}
                                                                _ = v120
                                                                _ = v121
                                                                _ = v122
                                                            }
                                                            v101 = (v4 >> ((genFn1()) & 63))
                                                        case 0:
                                                            fmt.Println(fi100, v53)
                                                        default:
                                                            var v123 int64 = -978
                                                            _ = v123
                                                        }
                                                    }
                                                    _ = fi113
                                                }
                                            } else {
                                                v4 = 42
                                                v4 += (v101 - (v4 / ((v4) | 1)))
                                                fmt.Println(v5)
                                            }
                                            _ = fi100
                                            _ = v101
                                        }
                                        _ = v99
                                    }
                                }
                            case -422:
                                var v124 int8 = -1
                                _ = v124
                            }
                        case "a":
                            switch "foo bar" {
                            case "hello":
                                var v125 float64 = v5
                                _ = v125
                            case "foo bar":
                                var v126 uint32 = v53
                                _ = v126
                            default:
                                var v127 float32 = -2.25
                                _ = v127
                            }
                        case "hello":
                            v4 += 0
                        }
                    } else {
                        v128 := GenS10{F0: 127, F1: 177}
                        v2 -= 1
                        _ = v128
                    }
                    v2 -= (v2 / ((v2) | 1))
                    v33 = append(v33, -1)
                }
                if v3 {
                    fmt.Println(v54, v4, v3)
                    for fi129 := 0; fi129 < 5; fi129++ {
                        if false {
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            switch "a" {
                            case "foo bar":
                                v130 := GenS2{F0: (fi129 - fi129), F1: (v4 / ((704) | 1))}
                                _ = v130
                            default:
                                switch (^(fi129)) {
                                case 956:
                                    if (v3 && v3) {
                                        v5 = (0.1 + v5)
                                    }
                                case 128:
                                    fmt.Println(v5, v53)
                                case 9223372036854775807:
                                    for fi131 := 0; fi131 < 7; fi131++ {
                                        if len(v33) > 0 {
                                            v33[0] = (fi129 - fi129)
                                        }
                                        v132 := GenS11{F0: (v2 << ((genFn1()) & 63))}
                                        v33 = append(v33, fi129)
                                        _ = fi131
                                        _ = v132
                                    }
                                default:
                                    for fi133 := 0; fi133 < 11; fi133++ {
                                        if v3 {
                                            var v134 uint32 = v53
                                            v135 := GenS12{F0: 92, F1: 42, F2: 1}
                                            var v136 float32 = 1e10
                                            _ = v134
                                            _ = v135
                                            _ = v136
                                        }
                                        if len(v33) > 0 {
                                            v33[0] = int(v4)
                                        }
                                        _ = fi133
                                    }
                                }
                            }
                        }
                        v137 := GenS2{F0: 42, F1: -1}
                        v138 := GenS9{F0: 1}
                        _ = fi129
                        _ = v137
                        _ = v138
                    }
                    v33 = append(v33, -514)
                }
                _ = v54
            }
        default:
            if true {
                for fi139 := 0; fi139 < 15; fi139++ {
                    fmt.Println(v53, v4, v3)
                    fmt.Println(v2)
                    _ = fi139
                }
                var v140 int16 = 32767
                v141 := GenS13{F0: (((("" + "\n\t") + "hello") + ("\n\t" + "foo bar")) + "\n\t"), F1: 1, F2: uint(v53)}
                _ = v140
                _ = v141
            } else {
                var v142 int16 = 42
                switch 32768 {
                case 938:
                    for fi143 := 0; fi143 < 2; fi143++ {
                        v144 := GenS14{F0: 127, F1: -128}
                        fmt.Println(v4, v3, v3)
                        var v145 int = (fi143 % (((fi143 * (fi143 ^ 32768))) | 1))
                        _ = fi143
                        _ = v144
                        _ = v145
                    }
                }
                v53 += v53
                _ = v142
            }
        }
        _ = v53
    } else {
        v146 := GenS15{F0: 0}
        for fi147 := 0; fi147 < 9; fi147++ {
            if (2147483647 <= 0) {
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
            } else {
                var v148 int8 = -128
                _ = v148
            }
            _ = fi147
        }
        for fi149 := 0; fi149 < 15; fi149++ {
            if true {
                for fi150 := 0; fi150 < 3; fi150++ {
                    if ((((3.14159 + (0.0 * float64(1.5))) * v5) > ((1e10 * v5) / 1.5)) || false) {
                        if v3 {
                            var v151 float32 = -2.25
                            _ = v151
                        } else {
                            if v3 {
                                var v152 uint16 = 255
                                if v3 {
                                    v4 = -987
                                    v153 := GenS1{F0: 0}
                                    v154 := GenS16{F0: "\n\t"}
                                    _ = v153
                                    _ = v154
                                } else {
                                    if ("\n\t" < "\n\t") {
                                        v2 -= 747
                                    }
                                }
                                var v155 uint = (v2 >> ((uint(v152)) & 63))
                                _ = v152
                                _ = v155
                            } else {
                                delete(v1, (((("" + "foo bar") + "foo bar") + "foo bar") + ""))
                                fmt.Println(fi149)
                                v4 += v4
                            }
                        }
                        var v156 uint16 = 32767
                        v2 = (v2 >> ((v2) & 63))
                        v156 = (^(v156))
                        _ = v156
                    } else {
                        if (!(fi150 != fi150)) {
                            if (509 >= 4294967295) {
                                v157 := GenS17{F0: -2}
                                v5 += v5
                                _ = v157
                            }
                            if v3 {
                                v3 = (!v3)
                            }
                            v33 = append(v33, fi150)
                        }
                    }
                    var v158 int16 = -32768
                    switch 1 {
                    case 0:
                        v3 = v3
                    }
                    v33 = append(v33, (fi149 >> (((v2 * v2)) & 63)))
                    _ = fi150
                    _ = v158
                }
            } else {
                v5 = float64((0.1 * (0.1 - (1e10 + float32(v5)))))
                var v159 int64 = 1
                _ = v159
            }
            var v160 float32 = 0.0
            v160 = (v160 + v160)
            _ = fi149
            _ = v160
        }
        _ = v146
    }
    for fi161 := 0; fi161 < 5; fi161++ {
        v162 := GenS18{F0: -2147483648, F1: 100}
        var v163 float32 = 1.5
        _ = fi161
        _ = v162
        _ = v163
    }
    for fi164 := 0; fi164 < 5; fi164++ {
        for fi165 := 0; fi165 < 9; fi165++ {
            var v166 uint = (v2 << (((v2 >> (((v2 & (^(v2)))) & 63))) & 63))
            _ = fi165
            _ = v166
        }
        v167 := GenS19{F0: 2}
        v33 = append(v33, 128)
        _ = fi164
        _ = v167
    }
    v3 = (!(65535 <= 65535))
    v5 += ((v5 * (v5 * v5)) * float64((float32(0.1) - 1.5)))
    switch "" {
    case "a":
        var v168 string = ("hello" + ("a" + ""))
        _ = v168
    case "":
        if v3 {
            v169 := GenS20{F0: genFn1()}
            _ = v169
        }
    case "\n\t":
        var v170 uint32 = 97
        _ = v170
    }
    if true {
        fmt.Println(v3, v2)
    }
    switch 667 {
    case 0:
        switch 32768 {
        case 2147483647:
            if (42 == 394) {
                var v171 uint = (v2 << ((v2) & 63))
                for fi172 := 0; fi172 < 8; fi172++ {
                    v33 = append(v33, (^(fi172)))
                    _ = fi172
                }
                if v3 {
                    var v173 uint32 = 0
                    _ = v173
                } else {
                    fmt.Println(v171, v3, v3)
                }
                _ = v171
            }
        case -9223372036854775808:
            delete(v1, "hello")
        default:
            if v3 {
                var v174 uint8 = 128
                if (!true) {
                    var v175 float64 = (v5 * (v5 * ((float64(1.5) / (v5 * v5)) / float64(float32(v5)))))
                    v176 := v1["foo bar"]
                    _ = v175
                    _ = v176
                } else {
                    v3 = true
                }
                switch "" {
                case "hello":
                    var v177 uint16 = 32767
                    _ = v177
                case "":
                    v3 = true
                case "a":
                    v2 -= uint(v174)
                }
                _ = v174
            }
        }
    }
    v178 := GenS21{F0: 0, F1: 228, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v178
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
