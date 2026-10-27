package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 string; F1 int16; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 uint32; F2 int32 }
type GenS8 struct { F0 string; F1 bool; F2 uint16 }
type GenS9 struct { F0 int8; F1 string }
type GenS10 struct { F0 uint16; F1 int32 }
type GenS11 struct { F0 int64; F1 int }
type GenS12 struct { F0 uint32; F1 int; F2 uint8 }
type GenS13 struct { F0 uint64; F1 int64; F2 string }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int32; F1 int8; F2 uint8 }
func genFn0(p0 int16) int16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 53, F1: -58}
        p0 = -921
        switch (fi7 >> ((42) & 63)) {
        case -961:
            p0 = (p0 * 0)
        case -147:
            v9 := GenS2{F0: "foo bar", F1: (^(p0)), F2: fi7}
            _ = v9
        case 1:
            p0 = (^(p0))
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return (p0 | -106)
}
func genFn1(p0 string, p1 uint8) uint {
    v10 := []int{}
    if false {
        var v11 float64 = 0.0
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint64 = 0
            for fi14 := 0; fi14 < 10; fi14++ {
                fmt.Println(fi14, fi14)
                var v15 uint32 = 42
                switch fi12 {
                case 327:
                    fmt.Println(p0, p0)
                default:
                    var v16 int8 = 42
                    _ = v16
                }
                switch "\n\t" {
                case "foo bar":
                    var v17 int16 = 1000
                    _ = v17
                case "hello":
                    var v18 uint = 92
                    _ = v18
                default:
                    if ((fi12 ^ fi12) <= fi14) {
                        if true {
                            if (0 >= 1) {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                v19 := GenS1{F0: 127, F1: -32768}
                                _ = v19
                            } else {
                                switch fi12 {
                                case 42:
                                    fmt.Println(p0, fi12, v11)
                                default:
                                    if len(v10) > 0 {
                                        v10[0] = 625
                                    }
                                }
                            }
                            v20 := GenS2{F0: p0, F1: 32767, F2: -2147483648}
                            if ((v15 | (v15 | (v15 % ((4294967295) | 1)))) <= 42) {
                                switch (^(fi14)) {
                                case 32768:
                                    switch fi14 {
                                    case 0:
                                        switch fi12 {
                                        case 4294967295:
                                            v15 -= v15
                                        case -562:
                                            v21 := GenS1{F0: 127, F1: genFn0(1)}
                                            _ = v21
                                        default:
                                            p1 += (^(p1))
                                        }
                                    case 259:
                                        v22 := GenS1{F0: 42, F1: int16(v13)}
                                        fmt.Println(v11, fi14)
                                        _ = v22
                                    case -2147483648:
                                        var v23 uint32 = v15
                                        _ = v23
                                    }
                                case 128:
                                    var v24 uint = 42
                                    _ = len(v10)
                                    v24 = v24
                                    if false {
                                        if (false && false) {
                                            var v25 string = p0
                                            var v26 uint = 1
                                            v27 := GenS1{F0: -111, F1: 32767}
                                            _ = v25
                                            _ = v26
                                            _ = v27
                                        } else {
                                            var v28 int = 128
                                            _ = v28
                                        }
                                        v29 := GenS2{F0: p0, F1: 302, F2: fi14}
                                        _ = v29
                                    } else {
                                        v13 = v13
                                    }
                                    if false {
                                        switch int(v24) {
                                        case 127:
                                            if true {
                                                fmt.Println(v15)
                                                fmt.Println(fi14, v24)
                                                if ((1.5 / -2.25) < (v11 / (v11 / v11))) {
                                                    v30 := GenS2{F0: p0, F1: 1000, F2: fi14}
                                                    _ = v30
                                                }
                                            }
                                        case 0:
                                            v10 = append(v10, 9223372036854775807)
                                        case 32767:
                                            fmt.Println(v13, v15, v11)
                                        }
                                        p0 = p0
                                        v31 := GenS3{F0: v13}
                                        _ = v31
                                    } else {
                                        switch "a" {
                                        case "":
                                            v11 = 0.1
                                        default:
                                            switch (("hello" + "foo bar") + "\n\t") {
                                            case "\n\t":
                                                if (0.0 > (0.0 + 3.14159)) {
                                                    var v32 float64 = ((((v11 * (v11 - v11)) + (v11 + 3.14159)) * float64(1.5)) * (v11 / (1e10 + 0.0)))
                                                    switch p0 {
                                                    case "hello":
                                                        if false {
                                                            v13 += (v13 >> ((775) & 63))
                                                            var v33 uint32 = (v15 % ((v15) | 1))
                                                            v34 := GenS2{F0: p0, F1: 1000, F2: fi14}
                                                            _ = v33
                                                            _ = v34
                                                        } else {
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            fmt.Println(v24, fi12)
                                                        }
                                                    }
                                                    _ = v32
                                                }
                                            case "foo bar":
                                                if len(v10) > 0 {
                                                    v10[0] = (fi14 % ((604) | 1))
                                                }
                                            case "":
                                                var v35 uint = v24
                                                v36 := GenS4{F0: (fi14 + (fi14 | fi12))}
                                                _ = v35
                                                _ = v36
                                            }
                                        }
                                    }
                                    _ = v24
                                default:
                                    var v37 uint32 = 0
                                    _ = v37
                                }
                                v15 = (^(v15))
                            } else {
                                if false {
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                } else {
                                    var v38 int64 = 0
                                    fmt.Println(v38, fi12, p1)
                                    var v39 uint16 = 636
                                    _ = v38
                                    _ = v39
                                }
                                fmt.Println(v11, fi12, v13)
                                var v40 uint64 = v13
                                _ = v40
                            }
                            switch p0 {
                            case "hello":
                                var v41 uint64 = (v13 << ((957) & 63))
                                _ = v41
                            case "a":
                                var v42 uint32 = uint32(v13)
                                _ = v42
                            }
                            switch fi12 {
                            case 127:
                                var v43 string = p0
                                _ = v43
                            case 42:
                                switch p0 {
                                case "":
                                    var v44 int64 = 1
                                    _ = v44
                                case "hello":
                                    if true {
                                        if false {
                                            v13 -= 18446744073709551615
                                        }
                                        var v45 int8 = -62
                                        _ = v45
                                    } else {
                                        if false {
                                            p1 += 42
                                            fmt.Println(p1, v15, fi12)
                                            var v46 uint64 = (v13 << ((65535) & 63))
                                            if (0 <= 859) {
                                                var v47 int32 = -528
                                                switch "foo bar" {
                                                case "":
                                                    if len(v10) > 0 {
                                                        v10[0] = fi14
                                                    }
                                                case "hello":
                                                    v15 = 23
                                                }
                                                _ = v47
                                            } else {
                                                v48 := GenS5{F0: (v15 >> ((255) & 63))}
                                                _ = v48
                                            }
                                            v13 = v13
                                            var v49 int64 = -818
                                            var v50 int16 = 32767
                                            fmt.Println(v13)
                                            v51 := GenS3{F0: 42}
                                            _ = v46
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        } else {
                                            if (-9223372036854775808 > -9223372036854775808) {
                                                if ((!((v11 * (v11 - ((((v11 + v11) * (v11 + v11)) + v11) + (0.1 - 3.14159)))) <= (1.5 + ((-2.25 - 1e10) + v11)))) && true) {
                                                    v52 := GenS4{F0: (fi12 ^ (fi12 | int(v13)))}
                                                    var v53 uint16 = 2
                                                    _ = v52
                                                    _ = v53
                                                }
                                                v11 -= v11
                                                if false {
                                                    fmt.Println(fi14, v13)
                                                }
                                            }
                                            var v54 float32 = (1e10 / 1.5)
                                            _ = v54
                                        }
                                    }
                                }
                            case -346:
                                fmt.Println(p0)
                            default:
                                p0 = p0
                            }
                            _ = v20
                        } else {
                            var v55 uint32 = v15
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v55
                        }
                    }
                }
                switch p0 {
                case "foo bar":
                    var v56 uint16 = 32767
                    _ = v56
                case "":
                    if len(v10) > 0 {
                        v10[0] = int(v13)
                    }
                default:
                    fmt.Println(fi12, v15, v11)
                }
                if false {
                    fmt.Println(p0, p0)
                    if (42 <= 631) {
                        switch p0 {
                        case "foo bar":
                            fmt.Println(p0, p1, v13)
                        case "hello":
                            var v57 int16 = -593
                            _ = v57
                        case "\n\t":
                            var v58 uint16 = 388
                            _ = v58
                        default:
                            if (-1 != 42) {
                                switch fi14 {
                                case 2147483647:
                                    var v59 int8 = -128
                                    _ = v59
                                }
                            }
                        }
                        v15 = (^(v15))
                        v13 += 18446744073709551615
                    }
                }
                _ = fi14
                _ = v15
            }
            var v60 int64 = 0
            _ = fi12
            _ = v13
            _ = v60
        }
        if false {
            v11 -= (-2.25 + 0.1)
            for fi61 := 0; fi61 < 12; fi61++ {
                p0 = "\n\t"
                _ = fi61
            }
            var v62 float32 = float32(3.14159)
            _ = v62
        }
        _ = v11
    } else {
        if len(v10) > 0 {
            v10[0] = 32767
        }
        if true {
            if false {
                var v63 string = "foo bar"
                var v64 uint = 267
                _ = v63
                _ = v64
            } else {
                var v65 uint32 = 907
                if len(v10) > 0 {
                    v10[0] = int(p1)
                }
                _ = v65
            }
            fmt.Println(p0, p0)
            v66 := GenS1{F0: 1, F1: 1000}
            _ = v66
        }
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = v10
    return 4294967295
}

func main() {
    var v1 uint32 = 1
    var v2 uint16 = 32767
    var v3 int16 = 1
    v4 := make(map[string]int)
    var v5 uint32 = 4294967295
    var v6 uint = 255
    v67 := v4["\n\t"]
    v68 := GenS6{F0: uint(v2)}
    if (!false) {
        var v69 string = "a"
        fmt.Println(v1)
        var v70 int8 = -105
        _ = v69
        _ = v70
    } else {
        for fi71 := 0; fi71 < 20; fi71++ {
            v2 += v2
            v6 -= (^(v6))
            var v72 bool = false
            _ = fi71
            _ = v72
        }
        delete(v4, "hello")
    }
    if (false || false) {
        for fi73 := 0; fi73 < 12; fi73++ {
            for fi74 := 0; fi74 < 12; fi74++ {
                v75 := []int{}
                var v76 uint = v6
                _ = fi74
                _ = v75
                _ = v76
            }
            _ = fi73
        }
        for fi77 := 0; fi77 < 14; fi77++ {
            v2 += v2
            for fi78 := 0; fi78 < 5; fi78++ {
                if false {
                    v3 = (v3 + v3)
                    if false {
                        if (!true) {
                            switch "\n\t" {
                            case "\n\t":
                                if false {
                                    if false {
                                        v2 -= v2
                                    } else {
                                        var v79 uint16 = v2
                                        fmt.Println(v67)
                                        _ = v79
                                    }
                                    v1 += (v5 - v1)
                                } else {
                                    _ = len(v4)
                                    if false {
                                        v80 := GenS7{F0: ("hello" != "foo bar"), F1: (v5 ^ v1), F2: 975}
                                        if true {
                                            v81 := GenS4{F0: fi77}
                                            switch (v67 - 179) {
                                            case 4294967295:
                                                if true {
                                                    v82 := GenS8{F0: (("\n\t" + "\n\t") + "hello"), F1: false, F2: v2}
                                                    _ = v82
                                                }
                                            case -1:
                                                var v83 float64 = ((-2.25 - -2.25) / 1.5)
                                                _ = v83
                                            }
                                            v84 := GenS4{F0: fi77}
                                            _ = v81
                                            _ = v84
                                        }
                                        if false {
                                            v2 -= (^(v2))
                                        }
                                        switch "hello" {
                                        case "a":
                                            v85 := []int{(v67 + 2147483647), -702}
                                            _ = v85
                                        default:
                                            var v86 float64 = 0.1
                                            _ = v86
                                        }
                                        _ = len(v4)
                                        var v87 bool = ((42 >= 482) && (true || false))
                                        _ = v80
                                        _ = v87
                                    } else {
                                        var v88 uint8 = 48
                                        switch "a" {
                                        case "hello":
                                            v89 := GenS6{F0: (v6 ^ v6)}
                                            _ = v89
                                        case "a":
                                            var v90 float32 = (-2.25 * 0.0)
                                            _ = v90
                                        default:
                                            if true {
                                                v6 = 42
                                                v2 = v2
                                            } else {
                                                fmt.Println(v3, fi77)
                                                v5 -= (^(v5))
                                                switch (v67 | fi77) {
                                                case 4294967295:
                                                    v91 := GenS5{F0: 1}
                                                    _ = v91
                                                case -521:
                                                    if (v5 < v1) {
                                                        var v92 uint8 = uint8(v6)
                                                        _ = v92
                                                    }
                                                case 0:
                                                    var v93 uint32 = (v5 ^ v5)
                                                    _ = v93
                                                default:
                                                    fmt.Println(v3, fi78)
                                                }
                                                v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
                                                v94 := GenS9{F0: 1, F1: ("" + "a")}
                                                if false {
                                                    if (false && (v5 > v1)) {
                                                        fmt.Println(v88, v2, v6)
                                                        v95 := GenS10{F0: 32767, F1: -1}
                                                        _ = v95
                                                    }
                                                    v96 := []int{55, v67, (v67 + fi78), v67}
                                                    v96 = append(v96, 32767)
                                                    _ = v96
                                                } else {
                                                    fmt.Println(fi78)
                                                    fmt.Println(v6)
                                                    fmt.Println(v5)
                                                }
                                                v3 -= v3
                                                _ = v94
                                            }
                                        }
                                        _ = v88
                                    }
                                }
                            }
                            v97 := v4["\n\t"]
                            var v98 bool = false
                            fmt.Println(v97)
                            if (!true) {
                                v99 := GenS11{F0: 42, F1: int(v2)}
                                if false {
                                    var v100 uint16 = (v2 & uint16(v6))
                                    v101 := GenS3{F0: 2}
                                    if true {
                                        if (v3 == 42) {
                                            if v98 {
                                                v5 = 4294967295
                                            }
                                            fmt.Println(v100)
                                            v6 -= uint(v2)
                                            if true {
                                                var v102 uint32 = (v1 & (^(v1)))
                                                _ = v102
                                            } else {
                                                if ((v6 <= v6) || v98) {
                                                    v6 = v6
                                                    if v98 {
                                                        if (v6 >= (v6 << ((v6) & 63))) {
                                                            v103 := []int{(fi78 << ((v6) & 63)), (v67 + (fi77 % ((fi78) | 1)))}
                                                            _ = v103
                                                        } else {
                                                            v4[("\n\t" + "foo bar")] = (fi77 + (fi77 % ((-1) | 1)))
                                                        }
                                                        v3 = (v3 / (((v3 << ((uint(v100)) & 63))) | 1))
                                                        var v104 uint8 = 42
                                                        _ = v104
                                                    }
                                                    var v105 int64 = -799
                                                    _ = v105
                                                } else {
                                                    v106 := []int{fi78, v97, v97}
                                                    fmt.Println(v100, fi78, v6)
                                                    v107 := GenS3{F0: 69}
                                                    _ = v106
                                                    _ = v107
                                                }
                                                var v108 int16 = (-(v3))
                                                _ = v108
                                            }
                                        }
                                        v109 := []int{fi77, (v97 | (^(v67))), (v97 & (v67 ^ -330)), 128}
                                        fmt.Println(v3)
                                        _ = v109
                                    }
                                    _ = v100
                                    _ = v101
                                } else {
                                    if v98 {
                                        v5 -= (^(v1))
                                        fmt.Println(v1, v6)
                                    } else {
                                        var v110 uint = 1
                                        if false {
                                            fmt.Println(v110)
                                            var v111 int8 = int8(v5)
                                            if v98 {
                                                v112 := GenS12{F0: v5, F1: v67, F2: 255}
                                                _ = v112
                                            } else {
                                                v4["hello"] = -1
                                                v67 -= 9223372036854775807
                                                v1 = 65535
                                                var v113 uint16 = (v2 >> ((0) & 63))
                                                _ = v113
                                            }
                                            _ = v111
                                        } else {
                                            var v114 int = (fi77 - v67)
                                            _ = v114
                                        }
                                        v115 := GenS13{F0: 42, F1: int64(v67), F2: ("" + "")}
                                        _ = v110
                                        _ = v115
                                    }
                                    fmt.Println(v67)
                                    fmt.Println(v67, v3)
                                    _ = len(v4)
                                    if (!((!(!v98)) || ((v1 >> (((v6 / ((v6) | 1))) & 63)) >= (^(v5))))) {
                                        if v98 {
                                            var v116 int = 127
                                            _ = v116
                                        }
                                        v3 -= (v3 & (v3 * (v3 / (((-(v3))) | 1))))
                                    } else {
                                        v117 := []int{32768, (fi78 >> (((v6 + genFn1("a", 2))) & 63)), 42}
                                        _ = v117
                                    }
                                }
                                v118 := GenS11{F0: int64(fi78), F1: 32768}
                                v119 := []int{-9223372036854775808, -2147483648, (fi77 % ((fi77) | 1))}
                                var v120 uint32 = (v1 / ((v5) | 1))
                                _ = v118
                                _ = v119
                                _ = v120
                                _ = v99
                            } else {
                                v121 := GenS9{F0: -1, F1: "a"}
                                var v122 string = (("a" + "\n\t") + ("hello" + ""))
                                var v123 bool = (2147483647 > 2147483647)
                                _ = v121
                                _ = v122
                                _ = v123
                            }
                            delete(v4, ("a" + ""))
                            _ = v97
                            _ = v98
                        } else {
                            v2 = (v2 + (v2 * (v2 - v2)))
                        }
                    } else {
                        v5 -= 0
                        var v124 uint8 = 183
                        _ = v124
                    }
                    var v125 float32 = float32(((3.14159 / 1.5) * 1.5))
                    _ = v125
                } else {
                    if (v67 <= v67) {
                        v6 -= v6
                    }
                }
                v126 := GenS4{F0: (fi77 >> ((v6) & 63))}
                _ = fi78
                _ = v126
            }
            if (-1 == -892) {
                v2 = v2
                var v127 uint8 = 43
                v5 -= uint32(v2)
                _ = v127
            } else {
                var v128 uint32 = (v5 + 2147483647)
                if (true || ("foo bar" >= ("a" + (("foo bar" + "foo bar") + ("a" + ""))))) {
                    var v129 uint8 = 128
                    _ = v129
                } else {
                    fmt.Println(v128, v2, v67)
                    fmt.Println(fi77, v1, v3)
                    v130 := []int{(v67 % (((fi77 * (fi77 % ((127) | 1)))) | 1)), (fi77 ^ (fi77 ^ 127)), fi77, 654}
                    _ = v130
                }
                _ = v128
            }
            _ = fi77
        }
    } else {
        v131 := []int{v67, v67, (v67 << ((genFn1("hello", 18)) & 63)), v67}
        _ = v131
    }
    if ((v3 / (((v3 & 1)) | 1)) != (v3 / (((v3 + -32768)) | 1))) {
        delete(v4, "")
    }
    v67 -= (v67 << ((uint(v2)) & 63))
    for fi132 := 0; fi132 < 2; fi132++ {
        switch (-(fi132)) {
        case 42:
            v133 := []int{(v67 / (((fi132 & (fi132 & v67))) | 1)), v67, -2147483648, (v67 & (-(fi132)))}
            _ = v133
        case 4294967295:
            var v134 float64 = 1e10
            _ = v134
        default:
            for fi135 := 0; fi135 < 3; fi135++ {
                switch fi132 {
                case -2147483648:
                    v3 = v3
                case 2147483647:
                    var v136 int = -711
                    _ = v136
                }
                v137 := GenS14{F0: (^(v1))}
                var v138 uint16 = (v2 * (^(v2)))
                fmt.Println(v67, v67, v6)
                var v139 int8 = 1
                delete(v4, (("\n\t" + "a") + (("hello" + "a") + ("hello" + "\n\t"))))
                _ = fi135
                _ = v137
                _ = v138
                _ = v139
            }
        }
        v140 := GenS15{F0: 2147483647, F1: -128, F2: 247}
        v2 = (v2 | v2)
        _ = fi132
        _ = v140
    }
    var v141 bool = (636 < 1)
    fmt.Println(v5, v1)
    var v142 bool = v141
    fmt.Println(v6, v3)
    var v143 float64 = (1.5 + 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v141
    _ = v142
    _ = v143
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
}
