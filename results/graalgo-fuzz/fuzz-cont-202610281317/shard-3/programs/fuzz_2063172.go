package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int32 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 string; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint32; F1 int; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int8; F1 string; F2 uint32 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int64; F1 uint; F2 uint32 }
type GenS11 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int }
type GenS14 struct { F0 bool; F1 uint64; F2 int16 }
type GenS15 struct { F0 bool; F1 bool }
func genFn0() uint8 {
    v5 := []int{1, 128, 1}
    var v6 bool = (-1 == -128)
    var v7 string = ""
    for fi8 := 0; fi8 < 17; fi8++ {
        _ = len(v5)
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(fi9, fi8)
            switch v7 {
            case "":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "a":
                v6 = v6
            case "hello":
                var v10 uint32 = 179
                _ = v10
            }
            if false {
                var v11 int = fi8
                fmt.Println(fi8, v11)
                _ = v11
            } else {
                var v12 uint = 0
                v5 = append(v5, fi9)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v12
            }
            v7 = v7
            _ = fi9
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            v6 = v6
            if (true && (1 >= -9223372036854775808)) {
                v7 = (v7 + v7)
                var v14 uint16 = 255
                var v15 int32 = 2147483647
                var v16 uint8 = 1
                if (false && v6) {
                    v17 := GenS1{F0: (v14 + 255), F1: -1, F2: (v15 << ((42) & 63))}
                    v6 = ((v7 + ((v7 + v7) + v7)) != v7)
                    switch (fi13 & fi13) {
                    case 32768:
                        v16 = v16
                        v18 := GenS1{F0: v14, F1: 592, F2: int32(v16)}
                        _ = v18
                    case -909:
                        switch v7 {
                        case "hello":
                            var v19 int64 = -399
                            _ = v19
                        case "":
                            v16 -= v16
                        case "\n\t":
                            if v6 {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v20 int = (^(fi8))
                                _ = v20
                            } else {
                                v7 = ("a" + "\n\t")
                            }
                        }
                    case -9223372036854775808:
                        if v6 {
                            var v21 bool = v6
                            v14 += v14
                            switch ("hello" + (v7 + v7)) {
                            case "foo bar":
                                v5 = append(v5, 4294967295)
                            case "hello":
                                if v21 {
                                    if v6 {
                                        switch "a" {
                                        case "hello":
                                            if len(v5) > 0 {
                                                v5[0] = -2147483648
                                            }
                                        case "\n\t":
                                            if len(v5) > 0 {
                                                v5[0] = (fi8 - (-(fi8)))
                                            }
                                        default:
                                            var v22 string = ((v7 + (v7 + v7)) + v7)
                                            _ = v22
                                        }
                                        var v23 uint = 1
                                        fmt.Println(v21, v7, v23)
                                        _ = v23
                                    }
                                    var v24 int8 = 100
                                    _ = v24
                                }
                            default:
                                v25 := GenS2{F0: 522, F1: 33}
                                _ = v25
                            }
                            if v6 {
                                if (439 > 42) {
                                    v26 := GenS2{F0: 18446744073709551615, F1: fi8}
                                    var v27 int = 412
                                    _ = v26
                                    _ = v27
                                } else {
                                    switch (v7 + v7) {
                                    case "\n\t":
                                        v28 := GenS3{F0: v7, F1: 576, F2: 42}
                                        _ = v28
                                    case "":
                                        var v29 int8 = -82
                                        var v30 int8 = v29
                                        _ = v29
                                        _ = v30
                                    case "foo bar":
                                        fmt.Println(fi8, v6)
                                    default:
                                        fmt.Println(v16, fi13)
                                    }
                                    v31 := GenS3{F0: v7, F1: 4294967295, F2: -904}
                                    _ = v31
                                }
                            } else {
                                v32 := GenS4{F0: 689, F1: fi13, F2: -536}
                                var v33 uint8 = v16
                                if v21 {
                                    switch v7 {
                                    case "foo bar":
                                        v6 = (1 <= 1)
                                        if (v7 == ("\n\t" + "a")) {
                                            fmt.Println(v14, v16)
                                        } else {
                                            if v6 {
                                                v34 := GenS4{F0: uint32(v33), F1: 128, F2: 424}
                                                if false {
                                                    var v35 uint32 = 2147483648
                                                    if (!(v14 < v14)) {
                                                        fmt.Println(v14, v33, v14)
                                                        v36 := GenS1{F0: (v14 ^ 718), F1: int64(v16), F2: v15}
                                                        if (105 != int8(fi13)) {
                                                            switch ("a" + v7) {
                                                            case "foo bar":
                                                                if (!(!true)) {
                                                                    _ = len(v5)
                                                                    var v37 int8 = 66
                                                                    var v38 uint = 1
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                            case "":
                                                                fmt.Println(v16, v35, v14)
                                                            }
                                                            var v39 float64 = 3.14159
                                                            var v40 int16 = 42
                                                            _ = v39
                                                            _ = v40
                                                        } else {
                                                            v7 = "foo bar"
                                                            switch fi8 {
                                                            case 2147483647:
                                                                if len(v5) > 0 {
                                                                    v5[0] = fi8
                                                                }
                                                            case -1:
                                                                v41 := GenS1{F0: 255, F1: -1, F2: (v15 << ((42) & 63))}
                                                                _ = v41
                                                            case 32767:
                                                                if v6 {
                                                                    v42 := GenS3{F0: (((v7 + v7) + v7) + "hello"), F1: 4294967295, F2: 42}
                                                                    v43 := GenS4{F0: v35, F1: (fi13 - (-(fi8))), F2: 42}
                                                                    _ = v42
                                                                    _ = v43
                                                                } else {
                                                                    v44 := GenS4{F0: (v35 << ((2) & 63)), F1: 598, F2: 0}
                                                                    switch fi13 {
                                                                    case -2147483648:
                                                                        v16 -= 193
                                                                    }
                                                                    _ = len(v5)
                                                                    fmt.Println(v15, fi8)
                                                                    _ = v44
                                                                }
                                                            default:
                                                                switch ("\n\t" + (("a" + "foo bar") + (v7 + v7))) {
                                                                case "\n\t":
                                                                    switch int(v14) {
                                                                    case 4294967295:
                                                                        v45 := GenS5{F0: (v33 ^ 127)}
                                                                        _ = v45
                                                                    default:
                                                                        v14 = v14
                                                                    }
                                                                case "":
                                                                    if ((v15 | (v15 + (v15 * v15))) != 2147483647) {
                                                                        var v46 float64 = 3.14159
                                                                        if (v21 && ((^(v33)) != v16)) {
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        }
                                                                        var v47 uint16 = (v14 >> ((910) & 63))
                                                                        fmt.Println(v21)
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                case "hello":
                                                                    switch fi13 {
                                                                    case 1:
                                                                        var v48 uint = 18446744073709551615
                                                                        _ = v48
                                                                    case -9223372036854775808:
                                                                        v35 = (v35 * 2147483647)
                                                                    case 32767:
                                                                        var v49 float32 = 1.5
                                                                        _ = v49
                                                                    default:
                                                                        var v50 int64 = -1
                                                                        _ = v50
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if true {
                                                            v35 -= (^(v35))
                                                            v51 := GenS6{F0: -1}
                                                            var v52 string = v7
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        var v53 uint8 = (^(v16))
                                                        switch v7 {
                                                        case "foo bar":
                                                            var v54 uint8 = v33
                                                            _ = v54
                                                        case "hello":
                                                            var v55 float32 = 1e10
                                                            _ = v55
                                                        case "a":
                                                            v56 := GenS1{F0: v14, F1: -1, F2: (v15 | v15)}
                                                            _ = v56
                                                        default:
                                                            var v57 string = (("" + (((((((v7 + v7) + v7) + "a") + v7) + "foo bar") + (((("foo bar" + v7) + "a") + "") + "a")) + v7)) + v7)
                                                            _ = v57
                                                        }
                                                        var v58 float64 = -2.25
                                                        _ = v36
                                                        _ = v53
                                                        _ = v58
                                                    }
                                                    var v59 string = v7
                                                    _ = v35
                                                    _ = v59
                                                } else {
                                                    var v60 int8 = 57
                                                    _ = v60
                                                }
                                                var v61 int16 = 0
                                                _ = v34
                                                _ = v61
                                            }
                                        }
                                    case "a":
                                        if v21 {
                                            v5 = append(v5, 44)
                                            var v62 uint16 = (v14 ^ (v14 & 32767))
                                            _ = v62
                                        }
                                    case "hello":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    default:
                                        var v63 string = ("foo bar" + "hello")
                                        _ = v63
                                    }
                                    v64 := GenS7{F0: (v14 >> ((18446744073709551615) & 63)), F1: (!false)}
                                    v33 = v16
                                    _ = v64
                                }
                                var v65 uint32 = 1
                                _ = v32
                                _ = v33
                                _ = v65
                            }
                            _ = v21
                        }
                    }
                    _ = v17
                }
                _ = v14
                _ = v15
                _ = v16
            } else {
                v6 = v6
                if len(v5) > 0 {
                    v5[0] = (-(fi8))
                }
            }
            if len(v5) > 0 {
                v5[0] = (fi13 | (fi13 - fi13))
            }
            _ = fi13
        }
        _ = fi8
    }
    var v66 int = 701
    _ = v5
    _ = v6
    _ = v66
    _ = v7
    return 96
}
func genFn1(p0 string, p1 bool) string {
    for fi67 := 0; fi67 < 19; fi67++ {
        switch ((p0 + p0) + p0) {
        case "a":
            switch p0 {
            case "foo bar":
                fmt.Println(fi67, p0, fi67)
            default:
                v68 := GenS8{F0: -22, F1: p0, F2: 1}
                _ = v68
            }
        case "":
            var v69 int32 = 0
            _ = v69
        }
        _ = fi67
    }
    var v70 int8 = 53
    v70 += (v70 << ((695) & 63))
    var v71 uint64 = 42
    _ = p0
    _ = p1
    _ = v70
    _ = v71
    return ""
}

func main() {
    v1 := []int{4294967295, 32767, 9223372036854775807}
    v2 := make(map[string]string)
    var v3 float64 = ((-2.25 / (1.5 + (1.5 * -2.25))) - -2.25)
    v4 := make(map[int]string)
    v72 := GenS9{F0: 595, F1: 1}
    switch 2147483647 {
    case 172:
        v3 = -2.25
    case -1:
        fmt.Println(v3)
    case 9223372036854775807:
        v73 := GenS2{F0: 18446744073709551615, F1: -224}
        _ = v73
    }
    for fi74 := 0; fi74 < 6; fi74++ {
        if ((float32((0.0 * v3)) / (1.5 + (1e10 / 1e10))) != -2.25) {
            for fi75 := 0; fi75 < 4; fi75++ {
                if true {
                    if (0 <= 1000) {
                        if ("\n\t" != "\n\t") {
                            if true {
                                v76 := GenS3{F0: ("foo bar" + ("hello" + ("foo bar" + ""))), F1: 543, F2: 9223372036854775807}
                                _ = v76
                            } else {
                                v77 := GenS5{F0: 52}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v4[fi75] = ""
                                _ = v77
                            }
                        }
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v3)
                            fmt.Println(v3, v3)
                            var v78 int = (fi74 ^ (-(fi74)))
                            _ = v78
                        }
                    }
                } else {
                    if (-9223372036854775808 > 42) {
                        var v79 int64 = -547
                        v4[1] = ""
                        _ = v79
                    } else {
                        var v80 int64 = -913
                        fmt.Println(v3)
                        var v81 uint16 = 1
                        if false {
                            if len(v1) > 0 {
                                v1[0] = (fi75 - 32768)
                            }
                            if (2 > 137) {
                                if true {
                                    if true {
                                        v82 := v2[(("a" + "\n\t") + "\n\t")]
                                        v82 = v82
                                        var v83 string = v82
                                        _ = v82
                                        _ = v83
                                    } else {
                                        v84 := v2["hello"]
                                        var v85 bool = true
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v84
                                        _ = v85
                                    }
                                    if ((2147483647 != -239) && (true || true)) {
                                        v81 = (v81 - 874)
                                        fmt.Println(v81)
                                        v80 -= (v80 / ((v80) | 1))
                                    } else {
                                        _ = len(v2)
                                        if true {
                                            if false {
                                                v3 -= (3.14159 * (v3 - (v3 - v3)))
                                                v1 = append(v1, fi74)
                                            }
                                            v81 += (v81 * (v81 % (((v81 % (((v81 >> ((2) & 63))) | 1))) | 1)))
                                        } else {
                                            var v86 uint16 = (v81 ^ (v81 & (v81 ^ v81)))
                                            v81 = v86
                                            fmt.Println(v86, v81)
                                            _ = v86
                                        }
                                    }
                                }
                                v87 := v2["foo bar"]
                                if true {
                                    v87 = v87
                                } else {
                                    var v88 int8 = 100
                                    if (99 != v88) {
                                        v3 = ((v3 * 1e10) * (1e10 - 0.0))
                                        v89 := GenS10{F0: 0, F1: 255, F2: 2}
                                        fmt.Println(v87)
                                        _ = v89
                                    }
                                    switch "" {
                                    case "":
                                        v90 := GenS11{F0: v80, F1: 2, F2: 128}
                                        _ = v90
                                    default:
                                        v91 := GenS12{F0: 705}
                                        _ = v91
                                    }
                                    _ = v88
                                }
                                _ = v87
                            }
                            _ = len(v1)
                        } else {
                            if true {
                                if true {
                                    v80 = v80
                                } else {
                                    v80 = (v80 % ((0) | 1))
                                }
                                if (!true) {
                                    fmt.Println(fi74)
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v92 uint32 = 2147483647
                                        _ = v92
                                    case "hello":
                                        v80 -= (v80 & (v80 << ((4294967295) & 63)))
                                    case "a":
                                        v2[("\n\t" + ("a" + "foo bar"))] = "hello"
                                    }
                                    if true {
                                        var v93 int32 = -174
                                        v80 += -9223372036854775808
                                        var v94 float64 = 0.0
                                        _ = v93
                                        _ = v94
                                    }
                                    var v95 int = (fi75 & -2147483648)
                                    _ = v95
                                }
                            }
                            if true {
                                v96 := GenS4{F0: 2, F1: fi75, F2: v80}
                                _ = v96
                            }
                            if (-128 != 24) {
                                v97 := GenS13{F0: (-(fi74))}
                                _ = len(v4)
                                switch "a" {
                                case "foo bar":
                                    v98 := GenS11{F0: v80, F1: 255, F2: 81}
                                    _ = v98
                                }
                                _ = v97
                            } else {
                                fmt.Println(v81, v80)
                                v99 := GenS10{F0: 42, F1: 0, F2: 42}
                                _ = v99
                            }
                        }
                        v80 -= (v80 % (((v80 % ((v80) | 1))) | 1))
                        _ = v80
                        _ = v81
                    }
                    v3 = ((v3 / (float64(float32(1e10)) + v3)) * v3)
                }
                switch ("\n\t" + (("foo bar" + ("hello" + "\n\t")) + "")) {
                case "a":
                    if true {
                        v3 = v3
                        var v100 int = (fi74 << ((1) & 63))
                        v100 = (v100 + v100)
                        _ = v100
                    }
                case "":
                    fmt.Println(fi74)
                case "foo bar":
                    var v101 string = ""
                    _ = v101
                }
                v102 := GenS8{F0: 0, F1: "foo bar", F2: 245}
                v3 = v3
                _ = fi75
                _ = v102
            }
            v3 = v3
        }
        _ = fi74
    }
    v3 += (float64(0.0) + ((0.0 + 0.1) * float64(((1.5 + (1.5 / float32(-2.25))) / 0.1))))
    switch 128 {
    case 2147483647:
        if (63 >= 204) {
            v103 := GenS14{F0: true, F1: 0, F2: 0}
            v104 := GenS7{F0: 566, F1: (171 != 813)}
            _ = v103
            _ = v104
        }
    case -2147483648:
        switch -1 {
        case -741:
            if (!true) {
                var v105 int32 = 1
                var v106 bool = (23 >= 255)
                _ = v105
                _ = v106
            }
        case -2147483648:
            v3 += (v3 - v3)
        case 42:
            var v107 uint64 = 2
            _ = v107
        }
    default:
        v1 = append(v1, -1)
    }
    var v108 uint64 = 2
    v3 += v3
    switch 32768 {
    case 128:
        v3 -= 1e10
    case -2147483648:
        var v109 int32 = -2147483648
        _ = v109
    default:
        var v110 int64 = 1
        _ = v110
    }
    var v111 int8 = -1
    for fi112 := 0; fi112 < 1; fi112++ {
        fmt.Println(fi112, v108, v3)
        switch -2147483648 {
        case 32768:
            v3 = float64(-2.25)
        case 42:
            var v113 uint16 = 1
            _ = v113
        }
        for fi114 := 0; fi114 < 2; fi114++ {
            if ((float32(v3) - (float32(v3) / 3.14159)) > -2.25) {
                var v115 int = fi114
                v115 = fi112
                _ = v115
            } else {
                v111 = v111
            }
            _ = fi114
        }
        _ = fi112
    }
    if (128 < -2147483648) {
        switch "hello" {
        case "a":
            v116 := GenS3{F0: "\n\t", F1: v108, F2: 1}
            _ = v116
        case "foo bar":
            v117 := GenS15{F0: (!((v108 % (((v108 & 1)) | 1)) < (^(v108)))), F1: false}
            _ = v117
        case "hello":
            var v118 int = 4294967295
            _ = v118
        }
        var v119 int8 = v111
        v3 += v3
        _ = v119
    }
    var v120 uint64 = (v108 / (((v108 | v108)) | 1))
    v111 += v111
    for fi121 := 0; fi121 < 10; fi121++ {
        v3 += v3
        _ = fi121
    }
    v122 := v4[0]
    var v123 uint = 4294967295
    if false {
        var v124 uint64 = v120
        var v125 uint32 = 0
        v111 = 65
        _ = v124
        _ = v125
    } else {
        delete(v4, 42)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v108:", v108)
    fmt.Println("v111:", v111)
    fmt.Println("v120:", v120)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v108
    _ = v111
    _ = v120
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v72
}
