package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int8; F1 uint16; F2 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32; F1 uint; F2 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16; F1 int16 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int8; F1 uint16; F2 int }
type GenS15 struct { F0 uint; F1 uint64 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 bool; F1 uint16 }
type GenS18 struct { F0 string; F1 uint32 }
type GenS19 struct { F0 int8; F1 uint8 }
func genFn0(p0 int8, p1 int, p2 uint32) string {
    fmt.Println(p2)
    if (255 < 2) {
        var v7 int = p1
        v8 := GenS1{F0: 255}
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := []int{fi9, fi9}
            if (42 <= 779) {
                p0 = 100
                v11 := GenS2{F0: 585}
                switch "hello" {
                case "a":
                    if true {
                        var v12 bool = (1 <= -1)
                        _ = v12
                    } else {
                        switch 42 {
                        case 1:
                            v13 := GenS1{F0: 65535}
                            _ = v13
                        case -1:
                            v14 := GenS3{F0: (^(p2))}
                            _ = v14
                        default:
                            var v15 uint16 = 65535
                            _ = v15
                        }
                        p1 += (fi9 & -676)
                        if true {
                            p2 += p2
                            v7 = 14
                        }
                    }
                case "foo bar":
                    fmt.Println(v7, p1)
                default:
                    var v16 bool = false
                    _ = v16
                }
                _ = v11
            } else {
                for fi17 := 0; fi17 < 11; fi17++ {
                    fmt.Println(fi9, p1)
                    if ((p2 << ((686) & 63)) >= 2147483647) {
                        if false {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        } else {
                            v18 := GenS2{F0: 1000}
                            p1 = fi9
                            if (((42 < 1) && true) && false) {
                                if false {
                                    p2 = p2
                                    if (2 == 116) {
                                        switch 878 {
                                        case 1:
                                            if true {
                                                var v19 uint64 = 18446744073709551615
                                                var v20 uint = 881
                                                _ = v19
                                                _ = v20
                                            }
                                        case 0:
                                            v21 := GenS4{F0: -894, F1: 9223372036854775807}
                                            _ = v21
                                        case 4294967295:
                                            if true {
                                                var v22 int16 = 1
                                                fmt.Println(fi9, p2)
                                                if false {
                                                    v22 += -626
                                                }
                                                _ = v22
                                            } else {
                                                p0 += -124
                                            }
                                        }
                                        fmt.Println(fi17, p0, v7)
                                        var v23 string = "foo bar"
                                        if (fi17 == (^(v7))) {
                                            switch fi17 {
                                            case 9223372036854775807:
                                                fmt.Println(p1)
                                            case 32768:
                                                v7 = (-(fi9))
                                            case 4294967295:
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                            default:
                                                v24 := GenS5{F0: p0, F1: 65535, F2: 283}
                                                _ = v24
                                            }
                                            if (1 >= 2) {
                                                if false {
                                                    var v25 uint8 = 115
                                                    v7 -= 0
                                                    v25 += v25
                                                    fmt.Println(p1, v7, v23)
                                                    if true {
                                                        fmt.Println(fi9, p0)
                                                        var v26 float32 = (0.1 * ((1.5 + float32((0.1 / 1.0))) - -2.25))
                                                        var v27 uint16 = 1
                                                        _ = v26
                                                        _ = v27
                                                    } else {
                                                        var v28 uint32 = 0
                                                        _ = v28
                                                    }
                                                    _ = v25
                                                } else {
                                                    if true {
                                                        var v29 uint8 = 120
                                                        v30 := GenS6{F0: v29}
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        if (false && true) {
                                                            v7 = (fi9 + p1)
                                                            var v31 uint64 = 2
                                                            _ = v31
                                                        } else {
                                                            if true {
                                                                v23 = (v23 + "foo bar")
                                                                var v32 int8 = (p0 + p0)
                                                                var v33 int16 = -1
                                                                p2 = 4294967295
                                                                var v34 uint64 = 0
                                                                _ = v32
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                            fmt.Println(p1, p2, p2)
                                                        }
                                                    }
                                                }
                                                var v35 int64 = 0
                                                v7 = 110
                                                v23 = "hello"
                                                _ = v35
                                            }
                                        } else {
                                            fmt.Println(p1, v23, p1)
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                            if ((p0 | 100) != 0) {
                                                var v36 int16 = 1
                                                switch "hello" {
                                                case "":
                                                    v36 += 1
                                                case "\n\t":
                                                    if true {
                                                        if (true || (v36 <= (v36 ^ (^(v36))))) {
                                                            var v37 uint = 18446744073709551615
                                                            _ = v37
                                                        }
                                                    } else {
                                                        var v38 uint16 = 255
                                                        switch (-(p1)) {
                                                        case -2147483648:
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                        default:
                                                            var v39 int8 = p0
                                                            _ = v39
                                                        }
                                                        _ = v38
                                                    }
                                                case "a":
                                                    if true {
                                                        if false {
                                                            _ = len(v10)
                                                        }
                                                        v23 = ""
                                                        _ = len(v10)
                                                    } else {
                                                        var v40 uint16 = 860
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                        if true {
                                                            if ((!(!false)) || (45 == 1)) {
                                                                if (42 < 195) {
                                                                    var v41 uint8 = 2
                                                                    _ = v41
                                                                } else {
                                                                    v10 = append(v10, -9223372036854775808)
                                                                    var v42 uint = 1
                                                                    if false {
                                                                        var v43 uint32 = p2
                                                                        p0 -= p0
                                                                        _ = v43
                                                                    } else {
                                                                        var v44 uint16 = (v40 >> ((391) & 63))
                                                                        _ = v44
                                                                    }
                                                                    _ = v42
                                                                }
                                                                var v45 int8 = int8(fi9)
                                                                _ = v45
                                                            }
                                                            var v46 bool = (0 <= 4294967295)
                                                            if len(v10) > 0 {
                                                                v10[0] = 42
                                                            }
                                                            var v47 string = v23
                                                            var v48 int8 = p0
                                                            if v46 {
                                                                if len(v10) >= 2 {
                                                                    v10 = v10[1:len(v10)]
                                                                }
                                                                v49 := GenS1{F0: 255}
                                                                var v50 uint32 = (^(p2))
                                                                var v51 uint64 = 984
                                                                if false {
                                                                    v23 = v23
                                                                    v52 := GenS7{F0: (v36 | -32768)}
                                                                    if (!(128 == 42)) {
                                                                        if v46 {
                                                                            fmt.Println(p2, v47, p0)
                                                                        } else {
                                                                            v53 := GenS4{F0: 1000, F1: p1}
                                                                            p1 = fi17
                                                                            _ = v53
                                                                        }
                                                                    } else {
                                                                        var v54 int64 = -866
                                                                        switch v7 {
                                                                        case 127:
                                                                            var v55 int32 = 1
                                                                            _ = v55
                                                                        case 0:
                                                                            fmt.Println(fi9)
                                                                        }
                                                                        var v56 string = (("hello" + ((v47 + v47) + (v23 + v23))) + v23)
                                                                        _ = v54
                                                                        _ = v56
                                                                    }
                                                                    _ = v52
                                                                } else {
                                                                    var v57 int16 = (v36 + int16(p0))
                                                                    fmt.Println(v23, fi9)
                                                                    _ = v57
                                                                }
                                                                _ = v49
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                        } else {
                                                            v10 = append(v10, (fi17 % ((v7) | 1)))
                                                        }
                                                        _ = v40
                                                    }
                                                default:
                                                    if true {
                                                        var v58 int64 = -1
                                                        _ = v58
                                                    }
                                                    if len(v10) > 0 {
                                                        v10[0] = fi17
                                                    }
                                                    if (true && true) {
                                                        if (4294967295 < 2) {
                                                            var v59 int32 = 1
                                                            switch v23 {
                                                            case "\n\t":
                                                                if (true || (((1e10 + 0.1) != (((float32(1.5) + 0.1) + 3.14159) - 3.14159)) || (true || (143 == 0)))) {
                                                                    v36 -= v36
                                                                    if false {
                                                                        if len(v10) > 0 {
                                                                            v10[0] = p1
                                                                        }
                                                                        if (false && true) {
                                                                            v36 = (^(v36))
                                                                            v60 := GenS8{F0: 2147483647, F1: 711, F2: "a"}
                                                                            v23 = v23
                                                                            fmt.Println(v59, p2, p0)
                                                                            var v61 int16 = 445
                                                                            var v62 int64 = 1
                                                                            var v63 uint64 = 0
                                                                            var v64 string = "\n\t"
                                                                            v65 := GenS1{F0: 116}
                                                                            v63 = (v63 >> ((2) & 63))
                                                                            var v66 string = "foo bar"
                                                                            _ = v60
                                                                            _ = v61
                                                                            _ = v62
                                                                            _ = v63
                                                                            _ = v64
                                                                            _ = v65
                                                                            _ = v66
                                                                        } else {
                                                                            if true {
                                                                                var v67 int16 = (v36 << ((4294967295) & 63))
                                                                                var v68 bool = (false && false)
                                                                                _ = v67
                                                                                _ = v68
                                                                            } else {
                                                                                var v69 uint16 = 255
                                                                                v69 = (^(v69))
                                                                                fmt.Println(v59, p1, fi9)
                                                                                _ = v69
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch (v23 + (v23 + (((v23 + (v23 + ("hello" + v23))) + "a") + ""))) {
                                                                        case "foo bar":
                                                                            fmt.Println(v59, v23, v59)
                                                                        case "a":
                                                                            v70 := GenS5{F0: p0, F1: 1, F2: v36}
                                                                            switch v23 {
                                                                            case "hello":
                                                                                switch fi17 {
                                                                                case -90:
                                                                                    v23 = v23
                                                                                case 32767:
                                                                                    if len(v10) > 0 {
                                                                                        v10[0] = 0
                                                                                    }
                                                                                case -2147483648:
                                                                                    v10 = append(v10, fi9)
                                                                                }
                                                                                if (fi17 > (fi17 << ((916) & 63))) {
                                                                                    switch (v7 & (fi17 | (v7 ^ int(p2)))) {
                                                                                    case -622:
                                                                                        if false {
                                                                                            var v71 bool = true
                                                                                            var v72 bool = false
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                        }
                                                                                    default:
                                                                                        var v73 uint64 = 241
                                                                                        _ = v73
                                                                                    }
                                                                                }
                                                                            case "":
                                                                                v74 := GenS9{F0: 0}
                                                                                _ = v74
                                                                            default:
                                                                                v10 = append(v10, (v7 ^ -2147483648))
                                                                                if len(v10) >= 2 {
                                                                                    v10 = v10[1:len(v10)]
                                                                                }
                                                                                _ = len(v10)
                                                                            }
                                                                            v75 := GenS8{F0: (p2 << ((666) & 63)), F1: 18446744073709551615, F2: v23}
                                                                            _ = v70
                                                                            _ = v75
                                                                        case "\n\t":
                                                                            if ((v59 << ((689) & 63)) == (v59 << ((872) & 63))) {
                                                                                var v76 int16 = -32768
                                                                                v36 = (-(v76))
                                                                                _ = v76
                                                                            } else {
                                                                                switch (fi17 ^ (p1 << ((4294967295) & 63))) {
                                                                                case 127:
                                                                                    if len(v10) >= 2 {
                                                                                        v10 = v10[1:len(v10)]
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    if ("a" == "a") {
                                                                                        if (170 >= 1) {
                                                                                            _ = len(v10)
                                                                                            _ = len(v10)
                                                                                            var v77 uint16 = 0
                                                                                            _ = v77
                                                                                        } else {
                                                                                            v78 := GenS1{F0: 18446744073709551615}
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(fi9, v36, v23)
                                                                                            default:
                                                                                                fmt.Println(p1, p2)
                                                                                            }
                                                                                            v79 := GenS6{F0: 128}
                                                                                            _ = v78
                                                                                            _ = v79
                                                                                        }
                                                                                    } else {
                                                                                        var v80 uint8 = 2
                                                                                        _ = v80
                                                                                    }
                                                                                    fmt.Println(v59)
                                                                                    v81 := GenS10{F0: p2}
                                                                                    _ = v81
                                                                                }
                                                                                if false {
                                                                                    var v82 float64 = 0.1
                                                                                    _ = v82
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi17, v7)
                                                                        }
                                                                        var v83 uint8 = 42
                                                                        v84 := GenS8{F0: 447, F1: 42, F2: (("foo bar" + v23) + (("a" + (v23 + "")) + ("a" + v23)))}
                                                                        switch 32768 {
                                                                        case 0:
                                                                            var v85 float32 = float32((0.1 - -2.25))
                                                                            _ = v85
                                                                        default:
                                                                            if true {
                                                                                fmt.Println(v59, v83)
                                                                                if len(v10) > 0 {
                                                                                    v10[0] = (fi9 / ((42) | 1))
                                                                                }
                                                                            } else {
                                                                                if (false && (v83 >= 128)) {
                                                                                    if false {
                                                                                        switch "a" {
                                                                                        case "\n\t":
                                                                                            var v86 int = -525
                                                                                            _ = v86
                                                                                        case "":
                                                                                            if true {
                                                                                                var v87 uint = 999
                                                                                                _ = v87
                                                                                            }
                                                                                        case "foo bar":
                                                                                            if (true || false) {
                                                                                                var v88 uint32 = p2
                                                                                                p2 = p2
                                                                                                v83 -= v83
                                                                                                _ = v88
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if false {
                                                                                            fmt.Println(p1)
                                                                                            if ((^(p2)) >= p2) {
                                                                                                v7 -= 659
                                                                                            } else {
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    var v89 int64 = 0
                                                                                                    _ = v89
                                                                                                default:
                                                                                                    fmt.Println(v7, p2)
                                                                                                }
                                                                                                v90 := GenS11{F0: p2}
                                                                                                _ = v90
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(fi9)
                                                                                        }
                                                                                        var v91 int = v7
                                                                                        _ = len(v10)
                                                                                        if false {
                                                                                            if len(v10) >= 2 {
                                                                                                v10 = v10[1:len(v10)]
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v36)
                                                                                            switch (p1 % (((p1 / (((v7 % ((fi9) | 1))) | 1))) | 1)) {
                                                                                            case 128:
                                                                                                switch fi17 {
                                                                                                case -810:
                                                                                                    var v92 bool = (836 < (p2 % ((p2) | 1)))
                                                                                                    _ = v92
                                                                                                case -2147483648:
                                                                                                    fmt.Println(v36, v36)
                                                                                                default:
                                                                                                    v7 = (v91 + (fi17 & (^(v7))))
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                v59 = v59
                                                                                            default:
                                                                                                if (126 == 9223372036854775807) {
                                                                                                    v91 += p1
                                                                                                    _ = len(v10)
                                                                                                    fmt.Println(fi9, p2)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v93 := GenS12{F0: 65535, F1: v36}
                                                                                        if len(v10) > 0 {
                                                                                            v10[0] = fi9
                                                                                        }
                                                                                        fmt.Println(v59, v83)
                                                                                        var v94 uint32 = p2
                                                                                        _ = v91
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    }
                                                                                } else {
                                                                                    v95 := GenS13{F0: v7}
                                                                                    _ = v95
                                                                                }
                                                                                v96 := GenS9{F0: 4294967295}
                                                                                _ = v96
                                                                            }
                                                                        }
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                    if (2 == 104) {
                                                                        var v97 int = 1
                                                                        p2 += p2
                                                                        v7 -= p1
                                                                        var v98 uint = 4294967295
                                                                        var v99 int = fi9
                                                                        var v100 uint64 = 0
                                                                        _ = v100
                                                                        _ = v97
                                                                        _ = v98
                                                                        _ = v99
                                                                    } else {
                                                                        v101 := GenS14{F0: 0, F1: 138, F2: (fi9 * v7)}
                                                                        var v102 uint8 = 14
                                                                        _ = len(v10)
                                                                        _ = v101
                                                                        _ = v102
                                                                    }
                                                                }
                                                            default:
                                                                var v103 int8 = 1
                                                                var v104 uint8 = 221
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                            var v105 string = (("\n\t" + "") + v23)
                                                            _ = v105
                                                            _ = v59
                                                        } else {
                                                            var v106 float64 = 3.14159
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            v107 := GenS15{F0: 65535, F1: 42}
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                        v36 = (-(v36))
                                                        _ = len(v10)
                                                    } else {
                                                        var v108 int32 = 1
                                                        fmt.Println(fi17, fi9)
                                                        fmt.Println(p0)
                                                        _ = v108
                                                    }
                                                    var v109 int64 = 1
                                                    _ = v109
                                                }
                                                _ = v36
                                            }
                                            switch -2147483648 {
                                            case 42:
                                                var v110 uint16 = 2
                                                _ = v110
                                            case -830:
                                                _ = len(v10)
                                            }
                                        }
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        if true {
                                            if true {
                                                switch fi9 {
                                                case 0:
                                                    if (2147483647 >= 534) {
                                                        var v111 float64 = 1e10
                                                        var v112 uint = 963
                                                        _ = v111
                                                        _ = v112
                                                    }
                                                case -2147483648:
                                                    switch "" {
                                                    case "\n\t":
                                                        v113 := GenS12{F0: 393, F1: 935}
                                                        _ = v113
                                                    }
                                                case -9223372036854775808:
                                                    var v114 uint = 1
                                                    _ = v114
                                                default:
                                                    v115 := GenS5{F0: p0, F1: 65535, F2: -1}
                                                    _ = v115
                                                }
                                                var v116 float32 = (1e10 + 1e10)
                                                fmt.Println(p2, p2, v7)
                                                _ = v116
                                            } else {
                                                v117 := GenS2{F0: 786}
                                                p0 += 42
                                                switch (v7 * -167) {
                                                case 0:
                                                    switch "hello" {
                                                    case "hello":
                                                        p1 = v7
                                                    case "a":
                                                        v7 = (^(fi17))
                                                    case "":
                                                        fmt.Println(v7, p0, p0)
                                                        v118 := GenS3{F0: p2}
                                                        var v119 uint = 214
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                case 375:
                                                    if false {
                                                        if (false && (true || true)) {
                                                            fmt.Println(v23, p2, p2)
                                                        }
                                                        if true {
                                                            var v120 int8 = p0
                                                            var v121 uint8 = 1
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "":
                                                    v122 := GenS16{F0: 9223372036854775807}
                                                    _ = v122
                                                case "hello":
                                                    var v123 int64 = 42
                                                    _ = v123
                                                case "a":
                                                    switch "hello" {
                                                    case "":
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                    case "foo bar":
                                                        v7 += (fi9 << ((0) & 63))
                                                    default:
                                                        var v124 uint64 = 399
                                                        _ = v124
                                                    }
                                                default:
                                                    p1 = (-(fi9))
                                                }
                                                switch fi9 {
                                                case -1:
                                                    fmt.Println(p0, fi9)
                                                }
                                                _ = v117
                                            }
                                        } else {
                                            v125 := GenS17{F0: false, F1: 65535}
                                            p0 += p0
                                            _ = v125
                                        }
                                        _ = len(v10)
                                        _ = v23
                                    }
                                    p1 = fi17
                                }
                                var v126 uint8 = 0
                                _ = v126
                            }
                            _ = v18
                        }
                    } else {
                        var v127 float64 = (1e10 / 0.1)
                        var v128 int = fi9
                        if (1000 >= -556) {
                            if true {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            } else {
                                var v129 int8 = 11
                                _ = v129
                            }
                            v7 = (v7 << ((747) & 63))
                        }
                        _ = v127
                        _ = v128
                    }
                    var v130 uint = 992
                    switch (fi9 * fi9) {
                    case 128:
                        if true {
                            p2 = (p2 * 259)
                        }
                    default:
                        fmt.Println(p0, fi17, fi17)
                    }
                    _ = fi17
                    _ = v130
                }
                if true {
                    var v131 uint = 2
                    for fi132 := 0; fi132 < 16; fi132++ {
                        v133 := GenS2{F0: 799}
                        _ = fi132
                        _ = v133
                    }
                    _ = v131
                }
                if len(v10) > 0 {
                    v10[0] = fi9
                }
            }
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        fmt.Println(p1, p1, p2)
        if ((!true) || (1000 <= 32767)) {
            v134 := GenS11{F0: (^(p2))}
            v135 := []int{}
            if len(v135) >= 2 {
                v135 = v135[1:len(v135)]
            }
            _ = v134
            _ = v135
        } else {
            var v136 int8 = (-(p0))
            v137 := []int{}
            _ = v136
            _ = v137
        }
    }
    switch 127 {
    case 128:
        v138 := GenS18{F0: "", F1: (p2 >> ((18446744073709551615) & 63))}
        _ = v138
    case 127:
        var v139 uint8 = 255
        _ = v139
    default:
        for fi140 := 0; fi140 < 12; fi140++ {
            if (!true) {
                var v141 uint = 642
                if (1 <= 42) {
                    p2 = p2
                    for fi142 := 0; fi142 < 7; fi142++ {
                        if false {
                            var v143 int = p1
                            v144 := GenS17{F0: false, F1: 65535}
                            _ = v143
                            _ = v144
                        }
                        v145 := []int{}
                        var v146 uint = 18446744073709551615
                        _ = fi142
                        _ = v145
                        _ = v146
                    }
                }
                _ = v141
            }
            _ = fi140
        }
    }
    p2 += 299
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 0
    var v3 int16 = -1
    var v4 int64 = 9223372036854775807
    var v5 uint = 255
    v6 := make(map[uint32]bool)
    v5 -= v5
    v2 = (v2 * (v2 + (v2 % ((100) | 1))))
    for fi147 := 0; fi147 < 19; fi147++ {
        var v148 uint8 = 255
        _ = fi147
        _ = v148
    }
    v149 := GenS19{F0: (v2 % ((v2) | 1)), F1: 1}
    v150 := GenS12{F0: 2, F1: -32768}
    fmt.Println(v4, v5, v3)
    var v151 int16 = (v3 * -32768)
    v152 := []int{4294967295, -2147483648, -793, 32768}
    for fi153 := 0; fi153 < 10; fi153++ {
        switch 42 {
        case 32767:
            delete(v6, 330)
        case 9223372036854775807:
            v151 = (v3 + int16(v4))
        }
        fmt.Println(v2, v4, v3)
        _ = fi153
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v151:", v151)
    fmt.Println("slicelen v152:", len(v152), v152)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v149
    _ = v150
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
