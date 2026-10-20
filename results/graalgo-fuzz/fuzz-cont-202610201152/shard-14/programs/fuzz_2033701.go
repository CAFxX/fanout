package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint16; F1 int64; F2 int8 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 bool; F1 uint; F2 uint }
type GenS15 struct { F0 uint8; F1 int8; F2 bool }
type GenS16 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS17 struct { F0 int8; F1 int16; F2 int }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 int8; F1 int }
type GenS21 struct { F0 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = false
    var v3 int32 = -2147483648
    v4 := GenS1{F0: ((1 <= 52) && v2), F1: 4294967295}
    if ((42 <= 557) && false) {
        if true {
            v5 := GenS2{F0: 1, F1: 0, F2: -128}
            if true {
                v6 := []int{1, -532, -1, -2147483648}
                v3 = (v3 * 697)
                _ = v6
            }
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := GenS1{F0: (32768 > -2147483648), F1: 0}
                _ = fi7
                _ = v8
            }
            _ = v5
        } else {
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := v1[("a" + "foo bar")]
                for fi11 := 0; fi11 < 1; fi11++ {
                    fmt.Println(v10)
                    switch fi11 {
                    case -2147483648:
                        v10 -= (v10 | (v10 * v10))
                    case -629:
                        fmt.Println(v2, v3, v2)
                    case 1:
                        if false {
                            if (0.1 >= 3.14159) {
                                if v2 {
                                    var v12 string = "a"
                                    var v13 uint = 65535
                                    _ = v12
                                    _ = v13
                                } else {
                                    v14 := GenS2{F0: 955, F1: 4294967295, F2: -128}
                                    v15 := GenS1{F0: false, F1: v10}
                                    _ = v14
                                    _ = v15
                                }
                                fmt.Println(fi11)
                            }
                            var v16 float64 = float64(float32(3.14159))
                            if (912 < 65535) {
                                var v17 uint64 = 1
                                v16 += v16
                                v16 -= v16
                                _ = v17
                            }
                            _ = v16
                        }
                    default:
                        if ((1 == 42) || (!v2)) {
                            if ((!(false || (((v2 && ((float64(1.5) / 1.0) == 0.0)) || v2) || (-128 > 0)))) && ("hello" >= ("\n\t" + "foo bar"))) {
                                if (false || ((v2 && v2) && true)) {
                                    switch (^(fi11)) {
                                    case 42:
                                        var v18 int8 = 0
                                        if v2 {
                                            v18 -= (v18 ^ (-(v18)))
                                        }
                                        _ = v18
                                    case 0:
                                        delete(v1, "hello")
                                    case -881:
                                        var v19 int64 = 560
                                        _ = v19
                                    }
                                    if (!false) {
                                        v3 -= v3
                                    }
                                    var v20 float64 = 1.5
                                    _ = len(v1)
                                    _ = v20
                                }
                                switch (fi11 & (fi11 / ((1) | 1))) {
                                case 9223372036854775807:
                                    v3 -= 2147483647
                                case 4294967295:
                                    v3 += 78
                                default:
                                    var v21 bool = (1e10 != (3.14159 / 3.14159))
                                    _ = v21
                                }
                            } else {
                                if false {
                                    v22 := []int{fi9, (fi9 * (fi9 + (fi9 >> ((351) & 63)))), fi9}
                                    var v23 string = "foo bar"
                                    v24 := GenS3{F0: 18446744073709551615, F1: 0}
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                if ((0.1 + 1e10) < 0.1) {
                                    v3 -= v3
                                }
                                if (false || (708 >= 4294967295)) {
                                    var v25 int64 = -1
                                    if (!v2) {
                                        var v26 int = -2147483648
                                        _ = v26
                                    }
                                    switch -48 {
                                    case 0:
                                        delete(v1, "hello")
                                    case 128:
                                        var v27 bool = (true || v2)
                                        _ = v27
                                    case 1:
                                        fmt.Println(fi9)
                                        var v28 string = ("" + "hello")
                                        if v2 {
                                            if true {
                                                var v29 string = ("" + (v28 + v28))
                                                _ = v29
                                            }
                                        }
                                        _ = v28
                                    }
                                    _ = v25
                                }
                                if v2 {
                                    v1[""] = 269
                                    v3 -= 42
                                } else {
                                    v10 = (^(v10))
                                    if false {
                                        var v30 float64 = (0.1 - (1e10 * (1.5 * float64(float32(1.5)))))
                                        var v31 int32 = (v3 - -141)
                                        var v32 uint16 = 2
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    } else {
                                        delete(v1, ("" + "hello"))
                                    }
                                }
                            }
                            if false {
                                var v33 float32 = (-2.25 - (0.0 / float32(-2.25)))
                                v3 -= v3
                                _ = v33
                            }
                            _ = len(v1)
                        } else {
                            if v2 {
                                var v34 bool = v2
                                v35 := GenS2{F0: 1, F1: (^(v10)), F2: 0}
                                _ = v34
                                _ = v35
                            }
                            var v36 int8 = -1
                            _ = v36
                        }
                    }
                    _ = fi11
                }
                for fi37 := 0; fi37 < 12; fi37++ {
                    if v2 {
                        if (fi37 >= fi37) {
                            v38 := GenS4{F0: 2, F1: v2}
                            v3 -= v3
                            v2 = false
                            _ = v38
                        } else {
                            switch (-(fi9)) {
                            case 214:
                                v3 -= (v3 / (((v3 << ((1) & 63))) | 1))
                                var v39 float64 = -2.25
                                _ = v39
                            case 0:
                                fmt.Println(v3, fi37)
                            case 42:
                                v2 = (1 > -97)
                                if v2 {
                                    v40 := []int{(fi37 ^ fi37), int(v3)}
                                    if v2 {
                                        v3 -= v3
                                        var v41 uint64 = 42
                                        if v2 {
                                            var v42 uint = 255
                                            fmt.Println(v10)
                                            _ = v42
                                        }
                                        if true {
                                            if (int64(v41) == 737) {
                                                var v43 int16 = 32767
                                                if (-1 != -1) {
                                                    var v44 int8 = -57
                                                    v2 = v2
                                                    _ = v44
                                                } else {
                                                    v45 := GenS5{F0: 32767, F1: 1, F2: 100}
                                                    v3 -= (v3 ^ (v3 * -1))
                                                    if (1.5 != 3.14159) {
                                                        v2 = v2
                                                        if v2 {
                                                            if len(v40) >= 2 {
                                                                v40 = v40[1:len(v40)]
                                                            }
                                                        }
                                                        v3 += 2147483647
                                                    }
                                                    _ = v45
                                                }
                                                _ = v43
                                            } else {
                                                fmt.Println(v2, v41, v10)
                                                v46 := GenS1{F0: (v3 == (v3 | 2147483647)), F1: (v10 | 4294967295)}
                                                _ = v46
                                            }
                                            var v47 float32 = 3.14159
                                            _ = v47
                                        }
                                        v40 = append(v40, fi37)
                                        _ = v41
                                    } else {
                                        v48 := GenS6{F0: "a", F1: 1}
                                        _ = v48
                                    }
                                    _ = v40
                                }
                                fmt.Println(v2, v10)
                            }
                            v3 = v3
                            v49 := []int{int(v10), 4294967295, 161}
                            switch fi9 {
                            case 42:
                                fmt.Println(v2)
                            }
                            if (v3 < -1) {
                                var v50 string = (("a" + "\n\t") + "")
                                if ((985 > 2) || v2) {
                                    v51 := GenS7{F0: 42, F1: ((-2.25 * -2.25) > float64((((1e10 / 1e10) * (1e10 * 1.5)) - -2.25))), F2: (^(v10))}
                                    v52 := v1[("a" + "hello")]
                                    if true {
                                        v53 := GenS8{F0: -1, F1: -32768}
                                        v3 += (v3 + int32(v10))
                                        if (v2 && (!(9223372036854775807 > -1))) {
                                            if false {
                                                v54 := GenS9{F0: v3}
                                                _ = v54
                                            }
                                        } else {
                                            var v55 uint64 = uint64(v10)
                                            if (-2.25 != 3.14159) {
                                                v56 := GenS10{F0: 42}
                                                fmt.Println(v2, v3)
                                                v1["\n\t"] = (v10 | uint32(v55))
                                                _ = v56
                                            }
                                            _ = v55
                                        }
                                        _ = v53
                                    } else {
                                        delete(v1, "a")
                                    }
                                    _ = v51
                                    _ = v52
                                }
                                var v57 uint16 = 65535
                                _ = v50
                                _ = v57
                            } else {
                                v58 := GenS6{F0: ("\n\t" + "hello"), F1: 180}
                                v59 := GenS11{F0: 554, F1: 316}
                                if (v2 && false) {
                                    var v60 bool = (140 < 255)
                                    var v61 uint16 = 299
                                    if v2 {
                                        v60 = true
                                        v49 = append(v49, fi9)
                                    } else {
                                        v49 = append(v49, 0)
                                    }
                                    v2 = (!(!((649 < 275) && v60)))
                                    var v62 int64 = 275
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                                _ = v58
                                _ = v59
                            }
                            fmt.Println(v2)
                            _ = v49
                        }
                    } else {
                        var v63 int = fi9
                        v64 := v1["a"]
                        _ = v63
                        _ = v64
                    }
                    var v65 uint8 = 42
                    if v2 {
                        v1["foo bar"] = uint32(v65)
                        if v2 {
                            v65 = (v65 * v65)
                            v3 = v3
                        } else {
                            var v66 float32 = -2.25
                            fmt.Println(v3, fi37)
                            fmt.Println(fi37, v2, v3)
                            _ = v66
                        }
                        if v2 {
                            var v67 bool = v2
                            if v67 {
                                v67 = (!(v67 || v67))
                                v68 := GenS12{F0: 42}
                                v3 = int32(fi37)
                                v69 := GenS5{F0: 42, F1: 502, F2: 100}
                                _ = v68
                                _ = v69
                            } else {
                                v70 := GenS13{F0: -465}
                                fmt.Println(v3, v2)
                                v67 = v2
                                _ = v70
                            }
                            _ = v67
                        }
                        switch "\n\t" {
                        case "":
                            if v2 {
                                v71 := GenS5{F0: 255, F1: -1, F2: 0}
                                _ = v71
                            } else {
                                v72 := GenS14{F0: v2, F1: 327, F2: 458}
                                v1["a"] = v10
                                _ = v72
                            }
                        case "hello":
                            if (-2.25 <= 0.1) {
                                v1[("" + "\n\t")] = v10
                                if true {
                                    if true {
                                        if v2 {
                                            if (!true) {
                                                switch "a" {
                                                case "hello":
                                                    v3 = v3
                                                }
                                            }
                                        } else {
                                            v73 := []int{2147483647}
                                            v74 := GenS3{F0: 42, F1: 700}
                                            fmt.Println(v3, v10, fi9)
                                            _ = v73
                                            _ = v74
                                        }
                                        var v75 uint32 = (v10 ^ v10)
                                        _ = v75
                                    } else {
                                        v10 += v10
                                        var v76 int8 = int8(v10)
                                        v3 -= (v3 * v3)
                                        _ = v76
                                    }
                                } else {
                                    fmt.Println(fi9)
                                    if v2 {
                                        v77 := GenS15{F0: 2, F1: 1, F2: ("\n\t" >= "a")}
                                        _ = v77
                                    }
                                }
                                fmt.Println(v10, fi9, v2)
                            } else {
                                v78 := GenS5{F0: 255, F1: int64(v3), F2: 38}
                                _ = v78
                            }
                        case "a":
                            v79 := GenS13{F0: 32767}
                            _ = v79
                        default:
                            v1[("foo bar" + ("\n\t" + "\n\t"))] = (v10 | v10)
                        }
                        v2 = (v3 >= v3)
                    }
                    v80 := GenS16{F0: v10, F1: (v10 + (^(v10))), F2: v10}
                    v2 = (int8(fi9) == int8(v10))
                    _ = fi37
                    _ = v65
                    _ = v80
                }
                _ = fi9
                _ = v10
            }
        }
        if v2 {
            v1["hello"] = 2147483648
            v81 := []int{-2147483648, 1, -9223372036854775808, 127}
            _ = v81
        }
    } else {
        fmt.Println(v3)
    }
    v3 = -1
    v3 -= (^(v3))
    var v82 int16 = 32767
    for fi83 := 0; fi83 < 4; fi83++ {
        switch "\n\t" {
        case "\n\t":
            for fi84 := 0; fi84 < 19; fi84++ {
                v82 = v82
                var v85 int8 = -17
                _ = fi84
                _ = v85
            }
        case "":
            delete(v1, (("hello" + ("a" + "")) + "a"))
        }
        v3 = (v3 / ((-13) | 1))
        v82 -= (v82 - (v82 - (-(v82))))
        _ = fi83
    }
    v86 := []int{42, 0}
    for fi87 := 0; fi87 < 7; fi87++ {
        for fi88 := 0; fi88 < 10; fi88++ {
            var v89 float64 = 0.1
            var v90 uint64 = 4294967295
            _ = fi88
            _ = v89
            _ = v90
        }
        _ = fi87
    }
    var v91 float32 = (0.0 * 1.5)
    var v92 int = 2147483647
    for fi93 := 0; fi93 < 3; fi93++ {
        v1[("\n\t" + (("hello" + "hello") + ""))] = 267
        var v94 int8 = -128
        _ = fi93
        _ = v94
    }
    switch ("a" + "hello") {
    case "\n\t":
        fmt.Println(v2, v82, v82)
    case "foo bar":
        var v95 int32 = v3
        _ = v95
    case "hello":
        if v2 {
            for fi96 := 0; fi96 < 6; fi96++ {
                v2 = (!((1 < 0) || v2))
                _ = fi96
            }
            if (127 < 42) {
                switch "hello" {
                case "a":
                    var v97 int8 = 0
                    _ = v97
                case "hello":
                    v86 = append(v86, v92)
                case "":
                    var v98 uint64 = 857
                    _ = v98
                default:
                    var v99 float64 = 1.5
                    _ = v99
                }
                if v2 {
                    v100 := GenS17{F0: 100, F1: (v82 * (v82 * (v82 | v82))), F2: v92}
                    _ = v100
                }
                if false {
                    v91 = ((0.0 + v91) / v91)
                    if true {
                        if (316 <= 4294967295) {
                            v101 := GenS12{F0: (v82 * v82)}
                            _ = v101
                        } else {
                            var v102 uint64 = 4294967295
                            v103 := GenS18{F0: 0}
                            _ = v102
                            _ = v103
                        }
                    } else {
                        if true {
                            var v104 uint64 = 4294967295
                            for fi105 := 0; fi105 < 10; fi105++ {
                                for fi106 := 0; fi106 < 14; fi106++ {
                                    fmt.Println(v92, v2)
                                    var v107 float32 = 1.5
                                    v2 = (!(!v2))
                                    _ = fi106
                                    _ = v107
                                }
                                if v2 {
                                    v108 := GenS11{F0: 32767, F1: (v104 | (v104 & 4294967295))}
                                    _ = v108
                                } else {
                                    v109 := GenS19{F0: 100}
                                    var v110 uint64 = v104
                                    for fi111 := 0; fi111 < 17; fi111++ {
                                        if (3.14159 >= (v91 / -2.25)) {
                                            if v2 {
                                                v112 := GenS20{F0: 0, F1: fi105}
                                                var v113 float64 = float64(v91)
                                                var v114 float32 = v91
                                                _ = v112
                                                _ = v113
                                                _ = v114
                                            } else {
                                                v82 -= v82
                                                var v115 float64 = 1e10
                                                v82 += (v82 + 862)
                                                v110 = v104
                                                _ = v115
                                            }
                                            var v116 uint32 = 1
                                            _ = v116
                                        }
                                        v104 = 18446744073709551615
                                        if true {
                                            if true {
                                                if v2 {
                                                    var v117 float64 = 0.1
                                                    _ = v117
                                                }
                                                var v118 int8 = -1
                                                _ = v118
                                            }
                                            fmt.Println(v2)
                                            if v2 {
                                                _ = len(v1)
                                                v2 = v2
                                                var v119 int32 = v3
                                                if true {
                                                    var v120 bool = v2
                                                    _ = v120
                                                }
                                                var v121 int16 = (^(v82))
                                                _ = v119
                                                _ = v121
                                            } else {
                                                delete(v1, "hello")
                                                v110 = v104
                                                switch int(v110) {
                                                case -2147483648:
                                                    if v2 {
                                                        v122 := GenS7{F0: 1, F1: (!((0.1 * 0.1) < 3.14159)), F2: 2}
                                                        _ = len(v1)
                                                        v123 := GenS21{F0: 352}
                                                        _ = v122
                                                        _ = v123
                                                    } else {
                                                        fmt.Println(v110, fi105, fi111)
                                                        var v124 int64 = int64(v3)
                                                        v125 := v1[("a" + "a")]
                                                        var v126 int32 = (v3 | (v3 << ((18446744073709551615) & 63)))
                                                        _ = v124
                                                        _ = v125
                                                        _ = v126
                                                    }
                                                case 2147483647:
                                                    var v127 int16 = v82
                                                    _ = v127
                                                }
                                                fmt.Println(v91, v2, v104)
                                                var v128 uint64 = 1
                                                if (42 < 255) {
                                                    v128 = 18446744073709551615
                                                    if (42 != 1) {
                                                        if (uint(v110) < 40) {
                                                            if v2 {
                                                                var v129 int32 = v3
                                                                switch (fi111 | 0) {
                                                                case 0:
                                                                    switch (fi105 - v92) {
                                                                    case 0:
                                                                        switch "" {
                                                                        case "a":
                                                                            var v130 int8 = -1
                                                                            _ = v130
                                                                        case "hello":
                                                                            v104 = v110
                                                                        case "":
                                                                            v129 = -526
                                                                        default:
                                                                            v92 = fi105
                                                                        }
                                                                    case 459:
                                                                        var v131 float32 = ((v91 / (-2.25 + v91)) * (1.5 / v91))
                                                                        _ = v131
                                                                    case 42:
                                                                        var v132 uint32 = 143
                                                                        _ = v132
                                                                    default:
                                                                        v129 = (v3 * 2147483647)
                                                                    }
                                                                case -9223372036854775808:
                                                                    v92 -= 1
                                                                case 1:
                                                                    v133 := GenS17{F0: int8(v104), F1: 32767, F2: (-(v92))}
                                                                    _ = v133
                                                                default:
                                                                    var v134 uint32 = 319
                                                                    _ = v134
                                                                }
                                                                v110 = (^(v128))
                                                                _ = v129
                                                            }
                                                            fmt.Println(fi111, v104, v92)
                                                        }
                                                        var v135 uint16 = 255
                                                        _ = v135
                                                    }
                                                    fmt.Println(v3, fi111, v3)
                                                } else {
                                                    v136 := GenS6{F0: "\n\t", F1: 1}
                                                    _ = v136
                                                }
                                                _ = v128
                                            }
                                        } else {
                                            var v137 float64 = 0.1
                                            var v138 int32 = (v3 | (-(v3)))
                                            switch "hello" {
                                            case "":
                                                var v139 float32 = 1.5
                                                _ = v139
                                            default:
                                                v92 += fi111
                                            }
                                            v140 := GenS6{F0: ("foo bar" + (("hello" + ((("foo bar" + (("foo bar" + "\n\t") + "foo bar")) + ("a" + ("" + ("" + "")))) + (("hello" + "\n\t") + "\n\t"))) + ("\n\t" + "foo bar"))), F1: 0}
                                            _ = v137
                                            _ = v138
                                            _ = v140
                                        }
                                        _ = fi111
                                    }
                                    _ = v109
                                    _ = v110
                                }
                                _ = fi105
                            }
                            _ = v104
                        } else {
                            switch (v92 | 4294967295) {
                            case 32767:
                                v2 = (v2 || (853 == 0))
                            case 42:
                                v92 = v92
                            case 9223372036854775807:
                                v82 = (v82 - v82)
                            default:
                                if v2 {
                                    fmt.Println(v91)
                                    var v141 float64 = ((float64(((v91 - v91) - 3.14159)) + -2.25) + 3.14159)
                                    if (!v2) {
                                        v142 := GenS3{F0: 529, F1: 65535}
                                        v143 := GenS13{F0: 1000}
                                        _ = v142
                                        _ = v143
                                    }
                                    _ = v141
                                }
                            }
                        }
                    }
                    switch ("a" + "foo bar") {
                    case "hello":
                        v144 := GenS12{F0: (v82 << ((36) & 63))}
                        _ = v144
                    case "\n\t":
                        if (!(v2 || ((int8(v92) < -1) && v2))) {
                            var v145 uint8 = 35
                            _ = v145
                        } else {
                            v82 = -970
                        }
                    case "":
                        delete(v1, ("" + "a"))
                    }
                }
            }
            fmt.Println(v3)
        } else {
            if false {
                v82 = v82
            }
        }
    }
    v2 = (255 == 65535)
    var v146 int16 = v82
    fmt.Println(v146)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v146:", v146)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v82:", v82)
    fmt.Println("slicelen v86:", len(v86), v86)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v146
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v86
    _ = v91
    _ = v92
}
