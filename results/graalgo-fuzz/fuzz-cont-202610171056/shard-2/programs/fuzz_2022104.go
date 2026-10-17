package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int8; F1 uint32; F2 bool }
type GenS8 struct { F0 string; F1 uint; F2 int32 }
type GenS9 struct { F0 int; F1 uint8; F2 uint8 }
type GenS10 struct { F0 int8; F1 int32; F2 int64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8; F1 uint }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint32; F1 uint32 }
func genFn0() string {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int8 = 32
        fmt.Println(v8)
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 613
    _ = v9
    return "foo bar"
}
func genFn1(p0 string) int8 {
    var v10 int8 = 114
    p0 = ((p0 + p0) + p0)
    v11 := []int{1}
    fmt.Println(v10)
    _ = p0
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint32 = 0
    var v2 int16 = 1
    v3 := []string{("\n\t" + ("\n\t" + "\n\t")), ""}
    var v4 uint16 = 32767
    v5 := make(map[uint32]string)
    var v6 int64 = 1
    v12 := GenS1{F0: 1, F1: v2, F2: v4}
    v13 := GenS2{F0: v1, F1: 0, F2: 803}
    delete(v5, 65535)
    switch "foo bar" {
    case "\n\t":
        switch "" {
        case "a":
            if true {
                var v14 uint32 = (v1 / (((v1 / (((v1 % (((v1 >> ((2) & 63))) | 1))) | 1))) | 1))
                var v15 float64 = ((0.0 * 0.1) * (-2.25 - 1.5))
                _ = v14
                _ = v15
            } else {
                switch (("" + "a") + "a") {
                case "foo bar":
                    v4 = (v4 ^ uint16(v1))
                case "hello":
                    var v16 uint32 = (v1 - 4294967295)
                    _ = v16
                case "\n\t":
                    if true {
                        var v17 string = ("a" + "foo bar")
                        if false {
                            v18 := GenS3{F0: -938}
                            _ = v18
                        }
                        _ = v17
                    }
                }
            }
        default:
            var v19 uint8 = 210
            _ = v19
        }
    case "a":
        var v20 uint8 = uint8(v1)
        _ = v20
    case "":
        v4 = v4
    default:
        switch -48 {
        case -946:
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS4{F0: 2, F1: 2, F2: 1}
                for fi23 := 0; fi23 < 20; fi23++ {
                    v24 := GenS5{F0: 0}
                    v1 += (v1 << ((255) & 63))
                    var v25 float32 = ((((1.5 - -2.25) - 0.0) - (float32((((-2.25 + 3.14159) + float64((((1.5 + 1.5) * 1e10) - 0.0))) + 1.5)) + -2.25)) * 0.1)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        }
    }
    switch ("\n\t" + "") {
    case "hello":
        var v26 bool = false
        _ = v26
    case "foo bar":
        v27 := v5[(v1 / ((v1) | 1))]
        _ = v27
    case "\n\t":
        v28 := GenS6{F0: 9223372036854775807, F1: 9223372036854775807}
        _ = v28
    default:
        v2 = v2
    }
    if true {
        v4 -= 255
        v4 = v4
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 float32 = 0.0
            var v31 int64 = v6
            fmt.Println(v30)
            _ = fi29
            _ = v30
            _ = v31
        }
    } else {
        v32 := v5[v1]
        var v33 int32 = 2147483647
        fmt.Println(v4, v6, v1)
        _ = v32
        _ = v33
    }
    var v34 uint8 = 0
    var v35 bool = (true || true)
    var v36 float64 = (0.1 / 1.5)
    v5[v1] = "hello"
    for fi37 := 0; fi37 < 5; fi37++ {
        v36 = (v36 * v36)
        var v38 uint16 = (^(v4))
        _ = fi37
        _ = v38
    }
    if v35 {
        for fi39 := 0; fi39 < 8; fi39++ {
            switch (fi39 << ((205) & 63)) {
            case 0:
                for fi40 := 0; fi40 < 20; fi40++ {
                    var v41 int = fi39
                    if false {
                        v41 += (fi39 + (fi39 * (^(fi40))))
                        if (!v35) {
                            var v42 uint8 = uint8(v4)
                            var v43 float64 = (0.0 * 1e10)
                            _ = v42
                            _ = v43
                        } else {
                            v1 += 261
                            if false {
                                v35 = (!v35)
                            }
                            var v44 float32 = 1e10
                            _ = v44
                        }
                        if v35 {
                            if false {
                                v41 += (v41 * (-(fi40)))
                            } else {
                                switch v41 {
                                case 451:
                                    var v45 int = (v41 << ((42) & 63))
                                    _ = v45
                                case 870:
                                    if (int8(v6) == 42) {
                                        switch "\n\t" {
                                        case "":
                                            if false {
                                                var v46 float64 = -2.25
                                                _ = v46
                                            } else {
                                                if (!false) {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v47 string = ""
                                                        _ = v47
                                                    }
                                                } else {
                                                    var v48 uint8 = v34
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = len(v3)
                                                    _ = v48
                                                }
                                                var v49 bool = true
                                                _ = v49
                                            }
                                        case "a":
                                            switch (v41 >> ((0) & 63)) {
                                            case 128:
                                                switch ("" + "hello") {
                                                case "\n\t":
                                                    switch (fi39 | 2147483647) {
                                                    case -690:
                                                        switch v41 {
                                                        case -906:
                                                            fmt.Println(v2, v4)
                                                            v35 = v35
                                                            if true {
                                                                var v50 string = ("hello" + "")
                                                                _ = len(v3)
                                                                if v35 {
                                                                    v50 = "\n\t"
                                                                } else {
                                                                    if (false && v35) {
                                                                        delete(v5, v1)
                                                                        _ = len(v5)
                                                                        var v51 int16 = (^(v2))
                                                                        _ = v51
                                                                    } else {
                                                                        var v52 uint64 = 42
                                                                        fmt.Println(v1)
                                                                        if ((v36 / 1.0) <= v36) {
                                                                            if (("foo bar" + (v50 + (((v50 + ((v50 + ((v50 + v50) + "\n\t")) + (("foo bar" + v50) + v50))) + ("foo bar" + (v50 + v50))) + ((v50 + ((v50 + v50) + v50)) + "\n\t")))) == v50) {
                                                                                var v53 uint16 = uint16(v1)
                                                                                switch 9223372036854775807 {
                                                                                case 42:
                                                                                    v36 -= v36
                                                                                case 4294967295:
                                                                                    var v54 string = v50
                                                                                    _ = v54
                                                                                default:
                                                                                    delete(v5, (v1 % (((v1 * v1)) | 1)))
                                                                                }
                                                                                _ = v53
                                                                            } else {
                                                                                v55 := GenS6{F0: (v6 << ((255) & 63)), F1: (v6 & -643)}
                                                                                var v56 int8 = -7
                                                                                v57 := GenS5{F0: v2}
                                                                                _ = v55
                                                                                _ = v56
                                                                                _ = v57
                                                                            }
                                                                            if v35 {
                                                                                var v58 float64 = float64(float32(v36))
                                                                                _ = v58
                                                                            }
                                                                            v59 := GenS5{F0: v2}
                                                                            _ = v59
                                                                        } else {
                                                                            v1 -= (v1 + v1)
                                                                            var v60 int8 = 1
                                                                            if v35 {
                                                                                var v61 int16 = (v2 | (v2 * (v2 * v2)))
                                                                                _ = v61
                                                                            } else {
                                                                                switch (fi40 % ((fi40) | 1)) {
                                                                                case 9223372036854775807:
                                                                                    var v62 int16 = int16(fi39)
                                                                                    _ = v62
                                                                                case 591:
                                                                                    delete(v5, (v1 ^ 4294967295))
                                                                                default:
                                                                                    v63 := GenS5{F0: v2}
                                                                                    _ = v63
                                                                                }
                                                                                v2 = int16(v1)
                                                                                if v35 {
                                                                                    var v64 int8 = (-(v60))
                                                                                    if (true || false) {
                                                                                        fmt.Println(v34)
                                                                                    } else {
                                                                                        var v65 uint64 = uint64(v4)
                                                                                        fmt.Println(v35)
                                                                                        v3 = append(v3, (v50 + "a"))
                                                                                        _ = v65
                                                                                    }
                                                                                    _ = v64
                                                                                }
                                                                                _ = len(v3)
                                                                            }
                                                                            if v35 {
                                                                                fmt.Println(v50)
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    fmt.Println(v2, fi40)
                                                                                default:
                                                                                    _ = len(v3)
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = ""
                                                                                }
                                                                            } else {
                                                                                switch "a" {
                                                                                case "":
                                                                                    v60 = v60
                                                                                case "a":
                                                                                    v60 = v60
                                                                                }
                                                                                v50 = v50
                                                                                switch 4294967295 {
                                                                                case 42:
                                                                                    v2 = v2
                                                                                case -1:
                                                                                    v4 = (v4 ^ v4)
                                                                                case -983:
                                                                                    var v66 string = (v50 + "foo bar")
                                                                                    _ = v66
                                                                                }
                                                                                var v67 int64 = (v6 - (^(v6)))
                                                                                v68 := GenS2{F0: (v1 / ((v1) | 1)), F1: 159, F2: 4294967295}
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                            _ = v60
                                                                        }
                                                                        _ = v52
                                                                    }
                                                                    var v69 int16 = (v2 << ((2) & 63))
                                                                    _ = v69
                                                                }
                                                                _ = v50
                                                            }
                                                        }
                                                    case -602:
                                                        if true {
                                                            var v70 uint8 = 42
                                                            switch "hello" {
                                                            case "":
                                                                v41 = fi39
                                                            case "\n\t":
                                                                var v71 int8 = 110
                                                                _ = v71
                                                            default:
                                                                v5[(v1 * 0)] = "foo bar"
                                                            }
                                                            _ = v70
                                                        }
                                                    case -2147483648:
                                                        var v72 uint64 = 2
                                                        _ = v72
                                                    }
                                                case "":
                                                    fmt.Println(fi40, v6)
                                                case "foo bar":
                                                    var v73 int32 = 1
                                                    _ = v73
                                                default:
                                                    var v74 uint = 0
                                                    _ = v74
                                                }
                                            default:
                                                var v75 string = "hello"
                                                _ = v75
                                            }
                                        }
                                        var v76 uint = 0
                                        _ = len(v3)
                                        var v77 int = (fi39 % ((int(v76)) | 1))
                                        _ = v76
                                        _ = v77
                                    } else {
                                        v35 = (42 <= 42)
                                    }
                                }
                                if true {
                                    v78 := GenS3{F0: -446}
                                    v34 = v34
                                    v79 := GenS7{F0: 92, F1: v1, F2: (!(genFn1(("a" + ((("foo bar" + ("" + "foo bar")) + (("\n\t" + "a") + "")) + "\n\t"))) == 100))}
                                    _ = v78
                                    _ = v79
                                }
                                if len(v3) > 0 {
                                    v3[0] = (("" + "a") + "\n\t")
                                }
                            }
                        }
                    } else {
                        if v35 {
                            v4 = (v4 << ((117) & 63))
                            switch (("foo bar" + ("\n\t" + "\n\t")) + (("foo bar" + "hello") + ("hello" + ""))) {
                            case "\n\t":
                                fmt.Println(v2, v2, v36)
                                var v80 uint32 = v1
                                var v81 int16 = (v2 % ((v2) | 1))
                                _ = v80
                                _ = v81
                            case "hello":
                                switch (fi40 * fi40) {
                                case 977:
                                    if false {
                                        v82 := GenS5{F0: (-(v2))}
                                        v2 = v2
                                        var v83 uint64 = 4294967295
                                        _ = v82
                                        _ = v83
                                    } else {
                                        var v84 uint8 = v34
                                        if (-9223372036854775808 == (v6 | (v6 | (v6 / (((^(v6))) | 1))))) {
                                            if false {
                                                v6 = v6
                                                v3 = append(v3, ("hello" + ("a" + ("a" + "foo bar"))))
                                                switch int(v1) {
                                                case 1:
                                                    v3 = append(v3, "\n\t")
                                                case 194:
                                                    var v85 bool = v35
                                                    _ = v85
                                                case 760:
                                                    v86 := GenS8{F0: "\n\t", F1: 2, F2: 42}
                                                    _ = v86
                                                default:
                                                    v87 := GenS9{F0: 42, F1: (v34 ^ uint8(v4)), F2: v34}
                                                    _ = v87
                                                }
                                            }
                                        } else {
                                            switch fi39 {
                                            case 0:
                                                fmt.Println(fi39, v84, fi39)
                                            case -114:
                                                if false {
                                                    fmt.Println(v36)
                                                    switch (^(v41)) {
                                                    case 42:
                                                        _ = len(v3)
                                                    case 4294967295:
                                                        v88 := GenS1{F0: 553, F1: 32767, F2: 42}
                                                        v35 = v35
                                                        if v35 {
                                                            fmt.Println(v4, v41, fi39)
                                                            var v89 float64 = 0.1
                                                            _ = v89
                                                        } else {
                                                            if (42 <= 596) {
                                                                var v90 uint32 = 42
                                                                var v91 int32 = -2147483648
                                                                v91 += v91
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                            if (v4 > 711) {
                                                                _ = len(v5)
                                                                if len(v3) > 0 {
                                                                    v3[0] = "a"
                                                                }
                                                            } else {
                                                                v6 = -686
                                                                fmt.Println(v34, fi39, fi39)
                                                            }
                                                        }
                                                        _ = v88
                                                    case 128:
                                                        fmt.Println(v4)
                                                    default:
                                                        v34 += v34
                                                    }
                                                    var v92 uint = 4294967295
                                                    v93 := GenS2{F0: 397, F1: 0, F2: 2}
                                                    if (v35 || true) {
                                                        var v94 string = "foo bar"
                                                        switch v41 {
                                                        case -9223372036854775808:
                                                            var v95 int32 = 2147483647
                                                            _ = v95
                                                        default:
                                                            v92 -= v92
                                                        }
                                                        var v96 uint16 = 632
                                                        _ = v94
                                                        _ = v96
                                                    } else {
                                                        v97 := GenS10{F0: genFn1(("a" + "foo bar")), F1: 0, F2: (v6 << ((826) & 63))}
                                                        _ = v97
                                                    }
                                                    _ = v92
                                                    _ = v93
                                                } else {
                                                    v98 := GenS4{F0: (v84 * 188), F1: v34, F2: 78}
                                                    v35 = (2147483647 > (fi39 / ((127) | 1)))
                                                    _ = v98
                                                }
                                            }
                                        }
                                        var v99 uint16 = (v4 | 2)
                                        switch ("hello" + ("a" + "a")) {
                                        case "a":
                                            var v100 bool = (!(!v35))
                                            var v101 float32 = 0.0
                                            _ = v100
                                            _ = v101
                                        case "\n\t":
                                            v1 = (v1 + 2147483647)
                                        default:
                                            if false {
                                                v102 := v5[1]
                                                _ = v102
                                            } else {
                                                fmt.Println(v36)
                                                if ((3.14159 - (-2.25 - (0.1 + (1.5 - -2.25)))) < 0.0) {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v103 float32 = -2.25
                                                        _ = v103
                                                    case "hello":
                                                        var v104 uint = 501
                                                        var v105 uint32 = v1
                                                        _ = v104
                                                        _ = v105
                                                    default:
                                                        _ = len(v5)
                                                    }
                                                    fmt.Println(v1, v41)
                                                    v35 = ((false || false) || (!(v34 != v34)))
                                                    if (!((!(-2.25 <= float32(float64(0.1)))) && (v36 >= v36))) {
                                                        if (v35 && false) {
                                                            v4 += v4
                                                        } else {
                                                            v106 := GenS1{F0: 443, F1: (v2 + v2), F2: (^(v4))}
                                                            fmt.Println(v1, v6)
                                                            _ = v106
                                                        }
                                                        var v107 string = ("a" + ("hello" + "a"))
                                                        _ = v107
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        fmt.Println(fi39, v34, v4)
                                                    }
                                                }
                                                if false {
                                                    v35 = ((v99 / ((v99) | 1)) == 282)
                                                    var v108 uint32 = (v1 + 292)
                                                    fmt.Println(v84, fi39, v34)
                                                    _ = v108
                                                }
                                            }
                                        }
                                        _ = v84
                                        _ = v99
                                    }
                                case -9223372036854775808:
                                    v6 = (^(v6))
                                default:
                                    v2 += (v2 * v2)
                                }
                            default:
                                if (v35 && (true || false)) {
                                    if (v41 < (v41 & (v41 << ((18446744073709551615) & 63)))) {
                                        v2 -= 1
                                        v36 += ((v36 + 0.1) - 3.14159)
                                        v2 += (v2 / ((v2) | 1))
                                    } else {
                                        var v109 bool = false
                                        switch v41 {
                                        case -2147483648:
                                            var v110 int16 = -1
                                            _ = v110
                                        case 503:
                                            var v111 uint8 = (v34 / (((^(v34))) | 1))
                                            _ = v111
                                        }
                                        var v112 int32 = 2147483647
                                        _ = v109
                                        _ = v112
                                    }
                                    fmt.Println(v4, v36, v36)
                                    switch "\n\t" {
                                    case "hello":
                                        var v113 uint64 = 18446744073709551615
                                        _ = v113
                                    default:
                                        switch "a" {
                                        case "hello":
                                            if v35 {
                                                if (!(v35 || (!(((!v35) || (!(false && ((v4 ^ v4) < v4)))) && false)))) {
                                                    v36 = (v36 - 3.14159)
                                                    switch 9223372036854775807 {
                                                    case -9223372036854775808:
                                                        if (!false) {
                                                            _ = len(v5)
                                                        } else {
                                                            var v114 uint16 = (v4 - (^(v4)))
                                                            v3 = append(v3, "\n\t")
                                                            if (v6 >= (v6 / ((int64(v114)) | 1))) {
                                                                var v115 int32 = 2147483647
                                                                _ = v115
                                                            } else {
                                                                v116 := GenS11{F0: 0}
                                                                _ = v116
                                                            }
                                                            _ = v114
                                                        }
                                                    case 32768:
                                                        v117 := GenS12{F0: -128, F1: 1}
                                                        _ = v117
                                                    }
                                                    v41 -= v41
                                                }
                                                v4 -= (v4 / ((32767) | 1))
                                                var v118 bool = v35
                                                _ = v118
                                            } else {
                                                if false {
                                                    var v119 int32 = -2147483648
                                                    v35 = v35
                                                    fmt.Println(v34)
                                                    _ = v119
                                                }
                                                var v120 uint8 = (v34 % (((^(v34))) | 1))
                                                _ = v120
                                            }
                                        case "":
                                            var v121 uint32 = 806
                                            v122 := GenS11{F0: 42}
                                            _ = v121
                                            _ = v122
                                        case "\n\t":
                                            if (!((v2 + (v2 % ((v2) | 1))) >= (v2 << ((0) & 63)))) {
                                                if ((("a" + "a") + ("a" + "hello")) != "a") {
                                                    fmt.Println(v41)
                                                    fmt.Println(v1)
                                                }
                                                switch (-(fi39)) {
                                                case 4294967295:
                                                    v123 := GenS13{F0: 661}
                                                    _ = v123
                                                }
                                                v6 -= (^(v6))
                                            }
                                        }
                                    }
                                    if (true || ((float64(0.0) * 1e10) >= (3.14159 / v36))) {
                                        v3 = append(v3, "a")
                                        _ = len(v5)
                                        v36 = (-2.25 * v36)
                                        var v124 bool = (v41 < fi40)
                                        fmt.Println(v1, v36)
                                        _ = v124
                                    }
                                    delete(v5, v1)
                                } else {
                                    if (v35 || ((v2 & v2) > (v2 + 1))) {
                                        var v125 int64 = v6
                                        v34 -= (v34 * 255)
                                        _ = v125
                                    }
                                }
                            }
                            if v35 {
                                var v126 uint8 = v34
                                _ = v126
                            } else {
                                var v127 int = (fi39 % (((fi40 >> ((255) & 63))) | 1))
                                fmt.Println(v6, v35, v2)
                                v34 += v34
                                if v35 {
                                    var v128 uint = uint(v34)
                                    _ = v128
                                }
                                _ = v127
                            }
                            var v129 uint = 4294967295
                            _ = v129
                        }
                    }
                    var v130 uint64 = 469
                    _ = fi40
                    _ = v130
                    _ = v41
                }
            case 32768:
                if (257 > (v6 * (v6 + 813))) {
                    if (2 != 65535) {
                        for fi131 := 0; fi131 < 16; fi131++ {
                            _ = len(v5)
                            switch fi131 {
                            case -227:
                                var v132 uint = 350
                                _ = v132
                            }
                            var v133 uint32 = 65535
                            _ = fi131
                            _ = v133
                        }
                        v5[v1] = "hello"
                    }
                    v134 := GenS14{F0: 1000}
                    _ = v134
                }
            case 822:
                if v35 {
                    v36 += float64(((((-2.25 * 0.0) + 1.5) / ((1.5 - float32(1e10)) * -2.25)) - -2.25))
                    if ((!(!v35)) || (((0.0 / (0.0 - (0.0 - float32(float64(float32(((v36 / v36) - v36))))))) - -2.25) == (0.0 + 0.1))) {
                        if v35 {
                            switch 0 {
                            case 127:
                                v34 -= (v34 >> ((0) & 63))
                            }
                        }
                        v135 := v5[uint32(v4)]
                        _ = v135
                    } else {
                        for fi136 := 0; fi136 < 17; fi136++ {
                            switch (fi136 / (((fi39 & 597)) | 1)) {
                            case -930:
                                if (1.5 <= float32(0.1)) {
                                    v5[(v1 ^ 1)] = "\n\t"
                                    _ = len(v3)
                                }
                            case -1:
                                v137 := GenS5{F0: v2}
                                _ = v137
                            }
                            fmt.Println(v6)
                            _ = fi136
                        }
                        v2 += v2
                        fmt.Println(v36, fi39, v1)
                    }
                    v138 := GenS2{F0: v1, F1: -1, F2: 2}
                    _ = v138
                }
            default:
                if v35 {
                    v3 = append(v3, "a")
                } else {
                    v34 -= (v34 - 41)
                    for fi139 := 0; fi139 < 17; fi139++ {
                        v36 += (v36 - v36)
                        switch "" {
                        case "\n\t":
                            var v140 float32 = (0.0 * 3.14159)
                            _ = v140
                        }
                        _ = fi139
                    }
                }
            }
            var v141 float32 = 1e10
            var v142 uint16 = 2
            _ = fi39
            _ = v141
            _ = v142
        }
        fmt.Println(v35, v4)
        for fi143 := 0; fi143 < 19; fi143++ {
            v144 := GenS15{F0: v1, F1: (v1 % (((v1 / ((v1) | 1))) | 1))}
            _ = fi143
            _ = v144
        }
    }
    var v145 int = -1
    for fi146 := 0; fi146 < 6; fi146++ {
        var v147 uint32 = (^(v1))
        if v35 {
            v148 := GenS6{F0: v6, F1: v6}
            var v149 uint32 = (^(v147))
            _ = v148
            _ = v149
        } else {
            fmt.Println(v147, v2, v147)
            var v150 uint16 = (v4 << ((18446744073709551615) & 63))
            _ = v150
        }
        v3 = append(v3, "a")
        _ = fi146
        _ = v147
    }
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v145
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
