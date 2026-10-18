package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 int8; F1 uint }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 bool; F1 uint32 }
type GenS14 struct { F0 uint32; F1 int64 }
type GenS15 struct { F0 int8; F1 int8; F2 uint16 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 bool; F1 bool }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 string; F1 string; F2 uint }
type GenS20 struct { F0 uint16; F1 uint }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 int; F1 uint8 }

func main() {
    v1 := make(map[string]int)
    v2 := []int{4294967295, -2147483648, -2147483648, 32767}
    var v3 int16 = -1
    var v4 int64 = 42
    var v5 string = (((("a" + "a") + ("hello" + "a")) + "foo bar") + "foo bar")
    fmt.Println(v5, v5, v4)
    for fi6 := 0; fi6 < 18; fi6++ {
        v4 -= v4
        if false {
            switch (fi6 | (fi6 >> ((1) & 63))) {
            case 2147483647:
                v2 = append(v2, (fi6 % ((fi6) | 1)))
            case 128:
                var v7 float64 = (0.0 + 1e10)
                _ = v7
            }
            if false {
                if (42 == 934) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            switch fi6 {
            case -121:
                v4 += 891
            case 4294967295:
                v8 := GenS1{F0: 4294967295, F1: 455, F2: 128}
                _ = v8
            default:
                switch 9223372036854775807 {
                case 1:
                    if (1 == 2147483648) {
                        if true {
                            for fi9 := 0; fi9 < 3; fi9++ {
                                if (!false) {
                                    v10 := GenS2{F0: ("a" + v5), F1: -128}
                                    var v11 int32 = -908
                                    if false {
                                        switch v5 {
                                        case "\n\t":
                                            v4 -= (v4 % ((42) | 1))
                                        case "":
                                            v12 := GenS3{F0: 2}
                                            _ = v12
                                        default:
                                            var v13 uint64 = 18446744073709551615
                                            _ = v13
                                        }
                                        if (true || true) {
                                            fmt.Println(v11)
                                        }
                                        v14 := GenS4{F0: v5}
                                        _ = v14
                                    }
                                    _ = v10
                                    _ = v11
                                } else {
                                    var v15 int8 = 1
                                    v15 = -107
                                    _ = v15
                                }
                                if true {
                                    switch -52 {
                                    case -2147483648:
                                        var v16 string = v5
                                        _ = v16
                                    }
                                    v4 = 933
                                    var v17 uint8 = 2
                                    _ = v17
                                } else {
                                    var v18 string = (v5 + v5)
                                    v19 := GenS5{F0: (v4 << ((276) & 63)), F1: (v3 * v3)}
                                    _ = v18
                                    _ = v19
                                }
                                _ = fi9
                            }
                            var v20 int = fi6
                            if false {
                                if (float64(float32(0.0)) < (((3.14159 + ((3.14159 * 0.1) + 0.1)) / -2.25) - 1e10)) {
                                    var v21 int16 = 42
                                    _ = v21
                                } else {
                                    fmt.Println(v20, v3)
                                }
                                for fi22 := 0; fi22 < 2; fi22++ {
                                    var v23 uint64 = 4294967295
                                    var v24 float32 = ((0.0 + 3.14159) / ((-2.25 + (float32(1.5) + (3.14159 / -2.25))) / (0.1 * -2.25)))
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                            }
                            _ = v20
                        } else {
                            if false {
                                var v25 int16 = v3
                                _ = v25
                            } else {
                                v2 = append(v2, (fi6 ^ fi6))
                                v26 := GenS4{F0: v5}
                                var v27 int = fi6
                                _ = v26
                                _ = v27
                            }
                            if (1.5 == 1.5) {
                                switch (v5 + ((((((("a" + "\n\t") + (v5 + ("foo bar" + v5))) + (("\n\t" + v5) + (("a" + ((v5 + "foo bar") + v5)) + v5))) + "") + v5) + "a") + v5)) {
                                case "":
                                    v28 := GenS6{F0: fi6, F1: (v5 + ((v5 + "foo bar") + "hello"))}
                                    _ = v28
                                case "foo bar":
                                    fmt.Println(fi6, fi6)
                                }
                                if false {
                                    v29 := GenS7{F0: (v4 * v4), F1: v5}
                                    if len(v2) > 0 {
                                        v2[0] = -336
                                    }
                                    switch "a" {
                                    case "hello":
                                        v4 -= v4
                                    case "\n\t":
                                        var v30 uint16 = 124
                                        _ = v30
                                    }
                                    _ = v29
                                } else {
                                    v31 := GenS4{F0: v5}
                                    var v32 int = fi6
                                    _ = v31
                                    _ = v32
                                }
                            }
                        }
                    }
                case 4294967295:
                    for fi33 := 0; fi33 < 15; fi33++ {
                        if true {
                            fmt.Println(v5, v5, fi33)
                            var v34 float64 = float64(((0.1 * 1e10) + -2.25))
                            switch (fi33 >> ((18446744073709551615) & 63)) {
                            case 32767:
                                v35 := GenS8{F0: 2, F1: -21}
                                _ = v35
                            default:
                                if len(v2) > 0 {
                                    v2[0] = fi6
                                }
                            }
                            _ = v34
                        } else {
                            v1[v5] = (-(fi33))
                            v36 := v1[v5]
                            _ = v36
                        }
                        if true {
                            var v37 uint16 = 899
                            v38 := GenS9{F0: 1, F1: 740}
                            _ = v37
                            _ = v38
                        } else {
                            var v39 int64 = v4
                            _ = v39
                        }
                        switch (fi33 + (fi6 ^ (fi33 / ((fi33) | 1)))) {
                        case -115:
                            var v40 string = (v5 + (v5 + "foo bar"))
                            _ = v40
                        case 4294967295:
                            var v41 float64 = 0.1
                            _ = v41
                        default:
                            if ((0.1 * (-2.25 - 1e10)) >= 0.1) {
                                fmt.Println(v4, v5)
                                switch "hello" {
                                case "\n\t":
                                    switch "\n\t" {
                                    case "":
                                        v5 = ""
                                    case "a":
                                        v42 := GenS5{F0: 42, F1: (v3 - 0)}
                                        v4 = (^(v4))
                                        var v43 int16 = 810
                                        _ = v42
                                        _ = v43
                                    default:
                                        switch (fi6 - -9223372036854775808) {
                                        case 32767:
                                            if ((3.14159 > (float32(-2.25) / -2.25)) || true) {
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch fi33 {
                                                    case 32768:
                                                        v3 = v3
                                                    case -2147483648:
                                                        if (v5 == v5) {
                                                            switch (fi6 / ((fi6) | 1)) {
                                                            case 32768:
                                                                if (true || (!(!true))) {
                                                                    v4 += v4
                                                                    fmt.Println(v4)
                                                                    var v44 int16 = 538
                                                                    _ = v44
                                                                } else {
                                                                    var v45 string = v5
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    switch fi33 {
                                                                    case 1:
                                                                        v4 = (v4 | v4)
                                                                    case -425:
                                                                        v4 += 62
                                                                    default:
                                                                        var v46 uint8 = 1
                                                                        _ = v46
                                                                    }
                                                                    _ = v45
                                                                }
                                                                if (0 != int32(fi33)) {
                                                                    if (!false) {
                                                                        v47 := GenS6{F0: fi33, F1: (v5 + "a")}
                                                                        var v48 int32 = 930
                                                                        if (129 == 2) {
                                                                            v49 := GenS7{F0: (v4 / ((v4) | 1)), F1: ((("" + v5) + v5) + v5)}
                                                                            if true {
                                                                                v50 := GenS10{F0: 128}
                                                                                if ((v48 / (((v48 | -111)) | 1)) == v48) {
                                                                                    if (true && ((42 != 214) || (false && (!false)))) {
                                                                                        var v51 float64 = 0.1
                                                                                        _ = v51
                                                                                    } else {
                                                                                        var v52 bool = false
                                                                                        fmt.Println(v4)
                                                                                        _ = v52
                                                                                    }
                                                                                    v48 -= (v48 / (((v48 % (((v48 % ((v48) | 1))) | 1))) | 1))
                                                                                    if false {
                                                                                        switch (-(fi6)) {
                                                                                        case 1:
                                                                                            var v53 int32 = v48
                                                                                            if (v3 == v3) {
                                                                                                switch "\n\t" {
                                                                                                case "hello":
                                                                                                    v53 = (v53 - 408)
                                                                                                default:
                                                                                                    var v54 int8 = -1
                                                                                                    _ = v54
                                                                                                }
                                                                                            }
                                                                                            v55 := GenS11{F0: 42}
                                                                                            v53 -= v53
                                                                                            v3 = 42
                                                                                            v4 += v4
                                                                                            var v56 uint8 = 2
                                                                                            if false {
                                                                                                fmt.Println(fi33, fi33)
                                                                                                if true {
                                                                                                    var v57 int16 = (v3 << ((0) & 63))
                                                                                                    v48 = v48
                                                                                                    _ = len(v1)
                                                                                                    _ = v57
                                                                                                } else {
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                }
                                                                                                if (998 >= v4) {
                                                                                                    v53 = (v53 / ((246) | 1))
                                                                                                } else {
                                                                                                    v48 = (v53 ^ -1)
                                                                                                    if true {
                                                                                                        var v58 int64 = v4
                                                                                                        var v59 int8 = 1
                                                                                                        v60 := GenS12{F0: -13}
                                                                                                        _ = v58
                                                                                                        _ = v59
                                                                                                        _ = v60
                                                                                                    }
                                                                                                    v61 := GenS8{F0: v56, F1: (v48 | 42)}
                                                                                                    _ = v61
                                                                                                }
                                                                                            } else {
                                                                                                v3 = -726
                                                                                                v4 += v4
                                                                                            }
                                                                                            v56 = v56
                                                                                            if (false && true) {
                                                                                                v3 = (v3 + v3)
                                                                                                v53 -= (v53 * 2147483647)
                                                                                            }
                                                                                            var v62 uint16 = 32767
                                                                                            _ = v53
                                                                                            _ = v55
                                                                                            _ = v56
                                                                                            _ = v62
                                                                                        default:
                                                                                            delete(v1, v5)
                                                                                        }
                                                                                        var v63 uint16 = 133
                                                                                        fmt.Println(v48)
                                                                                        _ = v63
                                                                                    } else {
                                                                                        var v64 int = 0
                                                                                        if true {
                                                                                            fmt.Println(v48, v5, v3)
                                                                                            var v65 int = 4294967295
                                                                                            v2 = append(v2, v65)
                                                                                            _ = v65
                                                                                        }
                                                                                        var v66 float32 = ((1e10 + (float32((0.1 * float64(float32(0.1)))) * 3.14159)) * 3.14159)
                                                                                        _ = v64
                                                                                        _ = v66
                                                                                    }
                                                                                } else {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = fi33
                                                                                    }
                                                                                    v67 := GenS13{F0: (0.1 == 0.0), F1: 2}
                                                                                    _ = v67
                                                                                }
                                                                                switch (("\n\t" + (v5 + ((v5 + "hello") + ""))) + (v5 + "hello")) {
                                                                                case "foo bar":
                                                                                    var v68 float32 = -2.25
                                                                                    if false {
                                                                                        if (!(!true)) {
                                                                                            v5 = v5
                                                                                            v69 := v1[(("foo bar" + v5) + "")]
                                                                                            delete(v1, v5)
                                                                                            _ = v69
                                                                                        } else {
                                                                                            fmt.Println(v48)
                                                                                            v70 := GenS9{F0: -128, F1: 114}
                                                                                            v5 = ""
                                                                                            _ = v70
                                                                                        }
                                                                                        var v71 bool = (527 < 2)
                                                                                        var v72 int = 42
                                                                                        _ = v71
                                                                                        _ = v72
                                                                                    }
                                                                                    _ = v68
                                                                                case "a":
                                                                                    var v73 uint32 = 42
                                                                                    _ = v73
                                                                                case "":
                                                                                    var v74 float32 = 1e10
                                                                                    _ = v74
                                                                                }
                                                                                _ = v50
                                                                            }
                                                                            var v75 uint64 = 4294967295
                                                                            var v76 uint32 = 201
                                                                            _ = v49
                                                                            _ = v75
                                                                            _ = v76
                                                                        } else {
                                                                            var v77 float32 = 1.5
                                                                            _ = v77
                                                                        }
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    var v78 string = (v5 + v5)
                                                                    _ = v78
                                                                } else {
                                                                    var v79 float32 = -2.25
                                                                    v4 += v4
                                                                    if (!(!(!true))) {
                                                                        v4 = v4
                                                                        v80 := GenS9{F0: 88, F1: 4294967295}
                                                                        _ = v80
                                                                    } else {
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            fmt.Println(fi6, v5, v3)
                                                                        case "":
                                                                            v81 := GenS2{F0: ("a" + ("foo bar" + v5)), F1: 127}
                                                                            _ = v81
                                                                        case "hello":
                                                                            if (false && true) {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 32767
                                                                                }
                                                                                v82 := v1[v5]
                                                                                var v83 uint64 = 42
                                                                                _ = v82
                                                                                _ = v83
                                                                            } else {
                                                                                if true {
                                                                                    if true {
                                                                                        var v84 uint16 = 255
                                                                                        var v85 int = 0
                                                                                        v79 = (1e10 * v79)
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    }
                                                                                    var v86 uint32 = 42
                                                                                    v87 := GenS1{F0: 4294967295, F1: 31, F2: 2}
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                }
                                                                                v88 := GenS14{F0: 4294967295, F1: (v4 / ((1) | 1))}
                                                                                v4 = 581
                                                                                _ = v88
                                                                            }
                                                                        default:
                                                                            fmt.Println(v4, v3)
                                                                        }
                                                                        var v89 string = ((v5 + v5) + "")
                                                                        _ = v89
                                                                    }
                                                                    switch (fi6 % ((fi33) | 1)) {
                                                                    case 32767:
                                                                        var v90 int = fi33
                                                                        _ = v90
                                                                    case 442:
                                                                        var v91 int8 = 92
                                                                        _ = v91
                                                                    case 594:
                                                                        if ((true || (!false)) && ((-(fi33)) != 504)) {
                                                                            if true {
                                                                                var v92 uint = 231
                                                                                if true {
                                                                                    fmt.Println(fi33, fi6, fi33)
                                                                                    fmt.Println(v79)
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v3)
                                                                                    _ = len(v2)
                                                                                    v79 = -2.25
                                                                                } else {
                                                                                    if ((-373 < v4) || (false || false)) {
                                                                                        var v93 int = (fi33 - fi6)
                                                                                        fmt.Println(fi33)
                                                                                        switch (("hello" + "a") + v5) {
                                                                                        case "hello":
                                                                                            v94 := GenS8{F0: 1, F1: 1}
                                                                                            _ = v94
                                                                                        case "a":
                                                                                            switch v5 {
                                                                                            case "\n\t":
                                                                                                var v95 int8 = 42
                                                                                                _ = v95
                                                                                            default:
                                                                                                var v96 float64 = float64(v79)
                                                                                                _ = v96
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v92, v3)
                                                                                        }
                                                                                        v1[(v5 + v5)] = 127
                                                                                        _ = v93
                                                                                    } else {
                                                                                        if true {
                                                                                            fmt.Println(v4, v5, fi6)
                                                                                            fmt.Println(v79)
                                                                                            v5 = v5
                                                                                        } else {
                                                                                            var v97 uint64 = 1
                                                                                            v5 = "hello"
                                                                                            var v98 uint16 = 2
                                                                                            fmt.Println(v79, v5, v98)
                                                                                            if true {
                                                                                                v5 = (v5 + v5)
                                                                                                v99 := v1["a"]
                                                                                                fmt.Println(fi33, v5)
                                                                                                _ = v99
                                                                                            } else {
                                                                                                var v100 float32 = 3.14159
                                                                                                var v101 uint8 = 66
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                            }
                                                                                            _ = v97
                                                                                            _ = v98
                                                                                        }
                                                                                        v92 = 2
                                                                                        v3 = v3
                                                                                    }
                                                                                    switch fi33 {
                                                                                    case 42:
                                                                                        switch fi6 {
                                                                                        case -666:
                                                                                            switch (fi6 | 1) {
                                                                                            case 2147483647:
                                                                                                v92 += (v92 * (v92 >> (((v92 >> ((v92) & 63))) & 63)))
                                                                                            case 128:
                                                                                                fmt.Println(fi6)
                                                                                            case 9223372036854775807:
                                                                                                v102 := GenS15{F0: -1, F1: 49, F2: 842}
                                                                                                _ = v102
                                                                                            default:
                                                                                                var v103 bool = (v3 != 1)
                                                                                                _ = v103
                                                                                            }
                                                                                        case 971:
                                                                                            v104 := v1[v5]
                                                                                            _ = v104
                                                                                        case 263:
                                                                                            var v105 float64 = 0.0
                                                                                            _ = v105
                                                                                        default:
                                                                                            if true {
                                                                                                var v106 string = v5
                                                                                                fmt.Println(v106, fi33)
                                                                                                v107 := GenS15{F0: 42, F1: 89, F2: 2}
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            }
                                                                                        }
                                                                                    case 127:
                                                                                        v5 = "foo bar"
                                                                                        v1[v5] = 102
                                                                                        switch v5 {
                                                                                        case "hello":
                                                                                            var v108 float32 = 1.5
                                                                                            _ = v108
                                                                                        default:
                                                                                            switch "a" {
                                                                                            case "\n\t":
                                                                                                var v109 uint = (v92 / ((v92) | 1))
                                                                                                _ = v109
                                                                                            case "hello":
                                                                                                fmt.Println(v92)
                                                                                            case "":
                                                                                                fmt.Println(v92)
                                                                                            default:
                                                                                                fmt.Println(v92, fi6, fi33)
                                                                                            }
                                                                                            v92 -= (v92 / ((2) | 1))
                                                                                            var v110 int = -2147483648
                                                                                            _ = v110
                                                                                        }
                                                                                    }
                                                                                    _ = len(v2)
                                                                                }
                                                                                _ = v92
                                                                            }
                                                                        } else {
                                                                            if false {
                                                                                v111 := GenS13{F0: false, F1: 0}
                                                                                _ = v111
                                                                            }
                                                                        }
                                                                    }
                                                                    if (true && (42 != 398)) {
                                                                        var v112 string = ((v5 + ("hello" + "a")) + v5)
                                                                        _ = v112
                                                                    }
                                                                    _ = v79
                                                                }
                                                            }
                                                            v4 += -1
                                                        } else {
                                                            var v113 uint64 = 1
                                                            var v114 uint32 = 2147483648
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                        var v115 string = (v5 + v5)
                                                        var v116 int32 = 0
                                                        _ = v115
                                                        _ = v116
                                                    }
                                                    var v117 float32 = ((((0.0 + (0.0 * -2.25)) * 1.5) - ((1e10 + ((-2.25 + (3.14159 - ((0.1 - 1e10) / 1e10))) + (1e10 - 1e10))) / (1.5 / -2.25))) + (float32(1.5) + 1.5))
                                                    fmt.Println(fi6, fi6, fi33)
                                                    var v118 uint16 = 695
                                                    switch 128 {
                                                    case 424:
                                                        switch fi33 {
                                                        case -9223372036854775808:
                                                            fmt.Println(v3, v118)
                                                            v117 -= float32(1.5)
                                                            fmt.Println(v5, v5, v118)
                                                        case 127:
                                                            if (false && (true && false)) {
                                                                var v119 float32 = 0.0
                                                                var v120 int32 = -1
                                                                v121 := GenS16{F0: 0}
                                                                _ = v119
                                                                _ = v120
                                                                _ = v121
                                                            } else {
                                                                v5 = ("\n\t" + ("" + (((v5 + v5) + v5) + "hello")))
                                                            }
                                                        default:
                                                            v5 = v5
                                                        }
                                                    case -927:
                                                        switch (v5 + v5) {
                                                        case "foo bar":
                                                            v118 += (v118 & (v118 & (v118 ^ 42)))
                                                        }
                                                    case 4294967295:
                                                        if false {
                                                            v1[("\n\t" + ("hello" + v5))] = (fi33 % ((fi6) | 1))
                                                            fmt.Println(v118, v117, v117)
                                                            var v122 uint32 = uint32(v118)
                                                            var v123 uint = 42
                                                            v124 := GenS2{F0: ("a" + ""), F1: -1}
                                                            var v125 uint64 = 18446744073709551615
                                                            var v126 uint = (^(v123))
                                                            _ = v122
                                                            _ = v123
                                                            _ = v124
                                                            _ = v125
                                                            _ = v126
                                                        } else {
                                                            v117 += 0.1
                                                        }
                                                    default:
                                                        fmt.Println(fi6, fi33)
                                                    }
                                                    _ = len(v2)
                                                    v118 -= 42
                                                    _ = v117
                                                    _ = v118
                                                } else {
                                                    var v127 float32 = float32(1e10)
                                                    _ = v127
                                                }
                                                v128 := GenS17{F0: (!((817 != 1) && false)), F1: false}
                                                v2 = append(v2, fi6)
                                                _ = v128
                                            } else {
                                                var v129 uint64 = 2
                                                _ = v129
                                            }
                                        }
                                    }
                                case "":
                                    v3 = v3
                                case "hello":
                                    v130 := GenS18{F0: false}
                                    _ = v130
                                default:
                                    if (419 == -116) {
                                        v3 -= 590
                                        v5 = (v5 + v5)
                                        var v131 float64 = 1.5
                                        if (255 >= 255) {
                                            var v132 uint8 = 0
                                            var v133 uint64 = 18446744073709551615
                                            _ = v132
                                            _ = v133
                                        }
                                        _ = v131
                                    } else {
                                        if true {
                                            v4 = (^(v4))
                                        } else {
                                            delete(v1, v5)
                                        }
                                        v134 := GenS19{F0: ((((v5 + (((v5 + v5) + "") + v5)) + v5) + ((("a" + v5) + "hello") + "")) + v5), F1: v5, F2: 590}
                                        _ = v134
                                    }
                                }
                                if false {
                                    v135 := GenS20{F0: 255, F1: 65535}
                                    v5 = "foo bar"
                                    _ = v135
                                } else {
                                    var v136 uint = 18446744073709551615
                                    v137 := GenS21{F0: (3.14159 >= (0.1 + (3.14159 + -2.25)))}
                                    _ = v136
                                    _ = v137
                                }
                            }
                        }
                        _ = len(v2)
                        _ = fi33
                    }
                }
            }
        }
        _ = fi6
    }
    v138 := GenS7{F0: v4, F1: v5}
    if false {
        v3 -= 0
        var v139 uint32 = 9
        _ = v139
    }
    var v140 uint16 = 255
    v2 = append(v2, 32768)
    v141 := GenS22{F0: -72, F1: 97}
    v3 = 32767
    v140 -= v140
    var v142 string = v5
    fmt.Println(v4)
    v140 = (v140 / ((32767) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v140:", v140)
    fmt.Println("v142:", v142)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v138
    _ = v140
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
