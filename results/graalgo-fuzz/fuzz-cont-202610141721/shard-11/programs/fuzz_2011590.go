package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 int32; F2 uint }
type GenS3 struct { F0 uint16; F1 int64; F2 bool }
type GenS4 struct { F0 uint8; F1 int; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int16; F1 uint16; F2 int8 }
type GenS14 struct { F0 int16; F1 int8; F2 int8 }
type GenS15 struct { F0 int16; F1 uint8; F2 int }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint64, p1 uint16) uint8 {
    p0 -= p0
    v7 := GenS1{F0: "foo bar", F1: 2}
    v8 := GenS1{F0: "\n\t", F1: 42}
    var v9 int8 = 1
    v10 := GenS2{F0: 77, F1: 2147483647, F2: 65535}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 132
}

func main() {
    var v1 bool = (1 >= 42)
    v2 := []uint32{65535, 65535, 4294967295}
    var v3 int64 = -423
    var v4 int = -2147483648
    var v5 int64 = (v3 % (((v3 % (((v3 / (((v3 & 9223372036854775807)) | 1))) | 1))) | 1))
    var v6 bool = (134 >= 0)
    v11 := GenS1{F0: "foo bar", F1: 128}
    if (((v1 && v1) || (v1 && (!v1))) || v6) {
        v2 = append(v2, 866)
        var v12 uint16 = 42
        if v1 {
            var v13 int16 = 1000
            _ = v13
        } else {
            var v14 int = v4
            if v6 {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v15 uint16 = v12
                _ = v15
            } else {
                fmt.Println(v5, v1)
                if (!(779 != 18446744073709551615)) {
                    fmt.Println(v3)
                } else {
                    switch v4 {
                    case 127:
                        switch "" {
                        case "foo bar":
                            v16 := GenS3{F0: 421, F1: v3, F2: ((-2.25 * float64(1e10)) > 1e10)}
                            _ = v16
                        case "a":
                            for fi17 := 0; fi17 < 8; fi17++ {
                                v12 += v12
                                _ = fi17
                            }
                        }
                    case 9223372036854775807:
                        switch 32768 {
                        case -1:
                            if (165 == 56) {
                                v6 = (!((v12 >> ((1) & 63)) != 32767))
                                v14 -= (^(v14))
                                var v18 int16 = 1000
                                _ = v18
                            } else {
                                if (float32(-2.25) >= float32(3.14159)) {
                                    for fi19 := 0; fi19 < 7; fi19++ {
                                        var v20 int64 = v3
                                        _ = fi19
                                        _ = v20
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = 4294967295
                                    }
                                } else {
                                    for fi21 := 0; fi21 < 12; fi21++ {
                                        for fi22 := 0; fi22 < 13; fi22++ {
                                            v23 := GenS4{F0: 1, F1: (v4 & (v14 / ((fi22) | 1))), F2: v12}
                                            v5 += (v5 >> ((65535) & 63))
                                            v4 = 9223372036854775807
                                            if v1 {
                                                fmt.Println(v4, fi22, v1)
                                                _ = len(v2)
                                                if v6 {
                                                    _ = len(v2)
                                                    var v24 uint16 = (v12 + v12)
                                                    _ = v24
                                                } else {
                                                    var v25 float64 = ((0.1 + 3.14159) / 3.14159)
                                                    if ((false || ((42 != -2147483648) && (!true))) || v6) {
                                                        v1 = (v1 || v1)
                                                        if v6 {
                                                            if v1 {
                                                                if ((v14 * (fi21 / ((fi21) | 1))) < (fi22 * 128)) {
                                                                    var v26 int = fi21
                                                                    _ = v26
                                                                }
                                                            }
                                                            var v27 int32 = 913
                                                            if true {
                                                                v5 -= -9223372036854775808
                                                            }
                                                            switch "foo bar" {
                                                            case "":
                                                                fmt.Println(fi21, v12, v6)
                                                            case "\n\t":
                                                                if (127 != 100) {
                                                                    if true {
                                                                        if (false || (!(42 < 4294967295))) {
                                                                            var v28 uint64 = 718
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 597
                                                                            }
                                                                            var v29 int8 = 100
                                                                            if (true || (2 < 18446744073709551615)) {
                                                                                v30 := GenS5{F0: 8}
                                                                                var v31 bool = v6
                                                                                var v32 uint16 = 1
                                                                                _ = v30
                                                                                _ = v31
                                                                                _ = v32
                                                                            }
                                                                            switch (fi21 >> ((4294967295) & 63)) {
                                                                            case 127:
                                                                                v33 := GenS1{F0: "a", F1: genFn0(v28, v12)}
                                                                                _ = v33
                                                                            case 2147483647:
                                                                                v34 := GenS4{F0: 55, F1: (fi21 / ((fi21) | 1)), F2: (^(v12))}
                                                                                _ = v34
                                                                            default:
                                                                                if false {
                                                                                    v28 = 18446744073709551615
                                                                                } else {
                                                                                    v35 := GenS2{F0: -498, F1: v27, F2: 871}
                                                                                    fmt.Println(v27)
                                                                                    _ = v35
                                                                                }
                                                                            }
                                                                            v36 := GenS6{F0: (v5 + 1), F1: 0}
                                                                            v37 := GenS6{F0: v5, F1: (^(v5))}
                                                                            var v38 uint64 = v28
                                                                            _ = v28
                                                                            _ = v29
                                                                            _ = v36
                                                                            _ = v37
                                                                            _ = v38
                                                                        } else {
                                                                            var v39 int64 = v5
                                                                            _ = v39
                                                                        }
                                                                        v12 = (v12 ^ (v12 | (v12 - 2)))
                                                                        switch ("" + "foo bar") {
                                                                        case "hello":
                                                                            v40 := GenS4{F0: 2, F1: v4, F2: 825}
                                                                            v5 = (-(v3))
                                                                            _ = v40
                                                                        case "\n\t":
                                                                            v14 = (v14 >> ((255) & 63))
                                                                        case "foo bar":
                                                                            v12 = (v12 | (^(v12)))
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v1)
                                                                        switch fi22 {
                                                                        case -684:
                                                                            switch (v14 - (v14 * v4)) {
                                                                            case -1:
                                                                                if v6 {
                                                                                    var v41 uint32 = 65535
                                                                                    var v42 int64 = (v3 ^ 9223372036854775807)
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (v41 | v41)
                                                                                    }
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                } else {
                                                                                    if ((42 <= fi22) || (false && false)) {
                                                                                        if ((v5 >= -9223372036854775808) || v1) {
                                                                                            if true {
                                                                                                var v43 uint64 = 2
                                                                                                var v44 int32 = v27
                                                                                                _ = v43
                                                                                                _ = v44
                                                                                            }
                                                                                            var v45 int64 = v3
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            _ = v45
                                                                                        } else {
                                                                                            if v1 {
                                                                                                var v46 uint16 = (v12 << ((0) & 63))
                                                                                                v47 := GenS2{F0: 1000, F1: v27, F2: 65535}
                                                                                                _ = v46
                                                                                                _ = v47
                                                                                            }
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "foo bar":
                                                                                            v48 := GenS7{F0: (("hello" + "\n\t") + "")}
                                                                                            _ = v48
                                                                                        case "a":
                                                                                            var v49 uint16 = 68
                                                                                            _ = v49
                                                                                        case "":
                                                                                            if ((1.5 / (0.0 + (0.0 + 1e10))) < 1e10) {
                                                                                                var v50 uint = 42
                                                                                                var v51 uint8 = 2
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    v4 -= fi22
                                                                                                case "":
                                                                                                    if v1 {
                                                                                                        var v52 int16 = -32768
                                                                                                        v50 -= (v50 - 4294967295)
                                                                                                        var v53 string = ""
                                                                                                        _ = v52
                                                                                                        _ = v53
                                                                                                    }
                                                                                                }
                                                                                                _ = v50
                                                                                                _ = v51
                                                                                            } else {
                                                                                                if (1.5 <= float64(float32(0.1))) {
                                                                                                    v54 := GenS8{F0: "a", F1: 42}
                                                                                                    _ = v54
                                                                                                } else {
                                                                                                    if v6 {
                                                                                                        var v55 uint = 42
                                                                                                        _ = v55
                                                                                                    } else {
                                                                                                        v56 := GenS7{F0: "foo bar"}
                                                                                                        v57 := GenS6{F0: v3, F1: (v3 + 618)}
                                                                                                        _ = v56
                                                                                                        _ = v57
                                                                                                    }
                                                                                                    var v58 int16 = -31
                                                                                                    _ = v58
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v3 -= v5
                                                                                            v14 = v14
                                                                                            v59 := GenS9{F0: false}
                                                                                            var v60 string = "\n\t"
                                                                                            fmt.Println(v3, v4)
                                                                                            _ = v59
                                                                                            _ = v60
                                                                                        }
                                                                                    } else {
                                                                                        var v61 uint = 0
                                                                                        _ = v61
                                                                                    }
                                                                                    v14 = -2147483648
                                                                                    if (true && false) {
                                                                                        v62 := GenS10{F0: 42, F1: 4294967295, F2: 100}
                                                                                        if (1.5 > v25) {
                                                                                            if (!(v4 != -2147483648)) {
                                                                                                v63 := GenS11{F0: 128}
                                                                                                fmt.Println(v6)
                                                                                                fmt.Println(fi21, v12, v4)
                                                                                                v64 := GenS10{F0: 4294967295, F1: 230, F2: -1}
                                                                                                var v65 int = int(v12)
                                                                                                if (v1 && true) {
                                                                                                    var v66 uint8 = 209
                                                                                                    _ = v66
                                                                                                } else {
                                                                                                    var v67 int16 = -32768
                                                                                                    v12 += (v12 >> ((255) & 63))
                                                                                                    var v68 bool = (543 != 0)
                                                                                                    _ = v67
                                                                                                    _ = v68
                                                                                                }
                                                                                                _ = v63
                                                                                                _ = v64
                                                                                                _ = v65
                                                                                            } else {
                                                                                                var v69 uint16 = v12
                                                                                                var v70 string = ((("\n\t" + ("" + ("" + ("hello" + "foo bar")))) + ("\n\t" + ("a" + "foo bar"))) + "a")
                                                                                                fmt.Println(v4)
                                                                                                v71 := GenS9{F0: false}
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                            }
                                                                                            switch 32767 {
                                                                                            case 0:
                                                                                                var v72 uint8 = 2
                                                                                                _ = v72
                                                                                            }
                                                                                        }
                                                                                        _ = v62
                                                                                    }
                                                                                    v5 = v5
                                                                                    v1 = (true && (2 > 2))
                                                                                }
                                                                            case 128:
                                                                                v73 := GenS10{F0: 42, F1: 2, F2: 33}
                                                                                _ = v73
                                                                            case 9223372036854775807:
                                                                                v1 = (-128 != 42)
                                                                            default:
                                                                                var v74 int = int(v27)
                                                                                _ = v74
                                                                            }
                                                                        case 127:
                                                                            v75 := GenS4{F0: 42, F1: -2147483648, F2: (v12 | v12)}
                                                                            _ = v75
                                                                        case 32768:
                                                                            v76 := GenS12{F0: -128}
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                    v14 = (fi21 << ((uint(v12)) & 63))
                                                                } else {
                                                                    var v77 int16 = 42
                                                                    switch v14 {
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v27, v25, v5)
                                                                    }
                                                                    _ = v77
                                                                }
                                                            }
                                                            _ = v27
                                                        } else {
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if v6 {
                                                                    fmt.Println(v25)
                                                                    v14 = -9223372036854775808
                                                                    v12 = 1
                                                                    fmt.Println(v12, v14, fi21)
                                                                    fmt.Println(fi21, v6)
                                                                } else {
                                                                    var v78 float32 = -2.25
                                                                    if v1 {
                                                                        v79 := GenS13{F0: -324, F1: (v12 % ((v12) | 1)), F2: -1}
                                                                        var v80 int8 = -1
                                                                        _ = len(v2)
                                                                        var v81 uint16 = 871
                                                                        var v82 string = ("a" + "hello")
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                        _ = v82
                                                                    } else {
                                                                        v14 += (fi22 % ((2147483647) | 1))
                                                                        switch (fi21 & fi21) {
                                                                        case 128:
                                                                            v83 := GenS10{F0: 18446744073709551615, F1: 18446744073709551615, F2: 78}
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                    v84 := GenS14{F0: -32768, F1: 42, F2: -38}
                                                                    _ = v78
                                                                    _ = v84
                                                                }
                                                            case "a":
                                                                v25 += v25
                                                            case "":
                                                                fmt.Println(v5)
                                                            }
                                                            if (42 == int8(v5)) {
                                                                v5 = (v5 * (v3 & 521))
                                                            } else {
                                                                switch "" {
                                                                case "a":
                                                                    switch ("\n\t" + "") {
                                                                    case "foo bar":
                                                                        if v1 {
                                                                            var v85 int32 = -32
                                                                            if ((fi22 + fi21) > fi22) {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            } else {
                                                                                var v86 uint16 = 215
                                                                                var v87 uint = 0
                                                                                _ = v86
                                                                                _ = v87
                                                                            }
                                                                            v12 -= (v12 % (((v12 + (v12 * (^(v12))))) | 1))
                                                                            _ = len(v2)
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 138
                                                                            }
                                                                            _ = v85
                                                                        } else {
                                                                            fmt.Println(fi22, v6)
                                                                            v88 := GenS15{F0: 42, F1: uint8(v12), F2: (v4 << ((4294967295) & 63))}
                                                                            v89 := GenS15{F0: 442, F1: 42, F2: fi21}
                                                                            v14 -= (v14 + (v4 / ((v14) | 1)))
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    var v90 uint32 = 2
                                                                    _ = v90
                                                                }
                                                                fmt.Println(v3)
                                                            }
                                                            var v91 int8 = 127
                                                            _ = v91
                                                        }
                                                        v6 = v6
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    v6 = ((true && v6) || (!(false || ((float64((3.14159 + 0.1)) + v25) == v25))))
                                                    _ = v25
                                                }
                                            }
                                            _ = fi22
                                            _ = v23
                                        }
                                        v6 = v6
                                        _ = fi21
                                    }
                                    var v92 float32 = ((0.1 - (((0.0 / -2.25) + 0.0) - -2.25)) - 0.1)
                                    _ = v92
                                }
                                v6 = v1
                                fmt.Println(v1, v14, v1)
                            }
                        default:
                            v6 = false
                        }
                    case 1:
                        v3 -= (v3 >> ((0) & 63))
                    }
                    for fi93 := 0; fi93 < 20; fi93++ {
                        if false {
                            v94 := GenS4{F0: 42, F1: v4, F2: 42}
                            for fi95 := 0; fi95 < 4; fi95++ {
                                if false {
                                    var v96 uint64 = 18446744073709551615
                                    var v97 int64 = (^(v5))
                                    _ = v96
                                    _ = v97
                                } else {
                                    var v98 uint = 263
                                    _ = v98
                                }
                                if ((2 != v12) && v1) {
                                    var v99 bool = ((176 != 172) || v6)
                                    _ = v99
                                }
                                _ = fi95
                            }
                            for fi100 := 0; fi100 < 4; fi100++ {
                                v101 := GenS16{F0: -968}
                                _ = fi100
                                _ = v101
                            }
                            _ = v94
                        } else {
                            for fi102 := 0; fi102 < 7; fi102++ {
                                fmt.Println(v4, fi102)
                                _ = len(v2)
                                if (v1 && false) {
                                    fmt.Println(v6, v14, fi102)
                                } else {
                                    var v103 uint16 = 65535
                                    v104 := GenS2{F0: 42, F1: 369, F2: 1}
                                    _ = v103
                                    _ = v104
                                }
                                _ = fi102
                            }
                            if (v1 || v6) {
                                var v105 bool = false
                                _ = v105
                            } else {
                                var v106 bool = (!((546 == 4294967295) && v6))
                                _ = v106
                            }
                        }
                        v5 = (^(v5))
                        _ = fi93
                    }
                    v14 = 1
                }
            }
            _ = v14
        }
        _ = v12
    } else {
        switch v4 {
        case 1:
            var v107 uint8 = 127
            _ = v107
        default:
            var v108 int = 32767
            _ = v108
        }
        if (!((v1 && (-1 >= 2147483647)) && (4294967295 != 0))) {
            switch int(v5) {
            case -506:
                v109 := GenS1{F0: "\n\t", F1: 127}
                _ = v109
            }
            switch "hello" {
            case "hello":
                if v6 {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v110 int = v4
                    v1 = false
                    _ = v110
                } else {
                    for fi111 := 0; fi111 < 12; fi111++ {
                        for fi112 := 0; fi112 < 8; fi112++ {
                            if true {
                                if true {
                                    v6 = true
                                    var v113 uint32 = 4294967295
                                    v4 = (fi112 / (((fi112 >> ((281) & 63))) | 1))
                                    var v114 string = ""
                                    var v115 float32 = float32((float64(0.1) * 0.1))
                                    _ = v113
                                    _ = v114
                                    _ = v115
                                } else {
                                    v116 := GenS7{F0: "foo bar"}
                                    _ = v116
                                }
                                fmt.Println(v4)
                            } else {
                                v3 += v3
                                v4 = v4
                            }
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            var v117 int64 = v5
                            var v118 string = (("\n\t" + "hello") + "")
                            _ = fi112
                            _ = v117
                            _ = v118
                        }
                        switch 1 {
                        case 32767:
                            var v119 float32 = 0.1
                            _ = v119
                        case 128:
                            for fi120 := 0; fi120 < 17; fi120++ {
                                var v121 int16 = 861
                                v5 += v3
                                _ = len(v2)
                                v121 = (v121 | (v121 << ((65535) & 63)))
                                fmt.Println(v5, v1, v121)
                                _ = fi120
                                _ = v121
                            }
                        case 2147483647:
                            v4 = (-(fi111))
                        default:
                            var v122 uint8 = 42
                            _ = v122
                        }
                        v2 = append(v2, 683)
                        _ = fi111
                    }
                    var v123 int16 = 1
                    _ = v123
                }
            case "\n\t":
                if v6 {
                    if false {
                        if v1 {
                            v6 = v1
                            v5 = v5
                            switch ("hello" + "hello") {
                            case "\n\t":
                                v1 = (true || (!(0 >= -791)))
                            case "foo bar":
                                v6 = (1.5 > -2.25)
                            default:
                                v2 = append(v2, 0)
                            }
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v124 := GenS11{F0: (v4 + (^(v4)))}
                        _ = v124
                    }
                    if v1 {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        v125 := GenS6{F0: (v5 + (^(v5))), F1: v5}
                        var v126 float32 = 3.14159
                        var v127 string = "hello"
                        _ = v125
                        _ = v126
                        _ = v127
                    }
                    v128 := GenS17{F0: 2147483647}
                    _ = v128
                }
            default:
                var v129 uint16 = 0
                _ = v129
            }
            v130 := GenS18{F0: -64}
            _ = v130
        }
        if false {
            v131 := GenS19{F0: 100, F1: 71}
            for fi132 := 0; fi132 < 10; fi132++ {
                switch (((("" + "foo bar") + "\n\t") + ("a" + "a")) + "a") {
                case "\n\t":
                    var v133 int8 = 100
                    _ = v133
                default:
                    v1 = false
                }
                for fi134 := 0; fi134 < 11; fi134++ {
                    var v135 int16 = 1000
                    _ = fi134
                    _ = v135
                }
                fmt.Println(v6)
                _ = fi132
            }
            if false {
                if (298 < 413) {
                    v6 = false
                    v2 = append(v2, 42)
                    if false {
                        switch (v4 & -391) {
                        case -1:
                            v1 = v6
                        case 127:
                            v136 := GenS18{F0: -25}
                            _ = v136
                        case 407:
                            v2 = append(v2, 4294967295)
                        }
                    }
                }
            } else {
                v4 += 783
                if (-128 < 0) {
                    v3 = (v3 >> ((255) & 63))
                }
            }
            _ = v131
        }
    }
    var v137 bool = (!(!(1 >= -72)))
    switch "\n\t" {
    case "foo bar":
        for fi138 := 0; fi138 < 9; fi138++ {
            var v139 uint32 = 301
            for fi140 := 0; fi140 < 2; fi140++ {
                v6 = (0.0 >= 1e10)
                var v141 int8 = -84
                _ = fi140
                _ = v141
            }
            v5 = (^(v5))
            _ = fi138
            _ = v139
        }
    }
    var v142 string = ""
    var v143 string = v142
    var v144 uint8 = 127
    var v145 float32 = -2.25
    for fi146 := 0; fi146 < 19; fi146++ {
        var v147 string = ("a" + "\n\t")
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = fi146
        _ = v147
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v148 string = (v142 + ("hello" + v142))
    var v149 uint8 = (v144 & (^(v144)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v142 = v142
    v143 = ("a" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v137
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v148
    _ = v149
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
