package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int64; F1 uint32; F2 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 uint64; F1 uint32 }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 int8; F1 string; F2 uint64 }
type GenS12 struct { F0 int32; F1 int32 }
type GenS13 struct { F0 uint64; F1 int16 }
type GenS14 struct { F0 uint16; F1 uint8 }
type GenS15 struct { F0 bool; F1 int16 }
type GenS16 struct { F0 int64; F1 uint8; F2 uint }
type GenS17 struct { F0 int; F1 uint8 }

func main() {
    var v1 bool = false
    var v2 float64 = 1e10
    var v3 int = 2147483647
    var v4 uint16 = 65535
    var v5 uint16 = v4
    var v6 string = "a"
    for fi7 := 0; fi7 < 5; fi7++ {
        switch (v3 >> ((255) & 63)) {
        case 32767:
            switch (v3 | (fi7 * fi7)) {
            case 127:
                if v1 {
                    v8 := []int{(-(fi7))}
                    switch (v6 + ((v6 + (("" + v6) + v6)) + "")) {
                    case "a":
                        v6 = "hello"
                    case "hello":
                        for fi9 := 0; fi9 < 3; fi9++ {
                            if len(v8) > 0 {
                                v8[0] = fi7
                            }
                            var v10 int = v3
                            var v11 int64 = -289
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    }
                    _ = v8
                }
            case -563:
                v12 := GenS1{F0: 351}
                _ = v12
            case 32768:
                fmt.Println(v6, v1, v5)
            default:
                switch -836 {
                case 4294967295:
                    for fi13 := 0; fi13 < 4; fi13++ {
                        fmt.Println(fi13, v3)
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                if (!((v2 * v2) != 1e10)) {
                                    if (false && (0 <= 830)) {
                                        v14 := []int{v3}
                                        v15 := GenS1{F0: 2}
                                        _ = v14
                                        _ = v15
                                    } else {
                                        v16 := GenS1{F0: 322}
                                        _ = v16
                                    }
                                } else {
                                    fmt.Println(fi7)
                                    v2 = v2
                                    fmt.Println(v4)
                                    var v17 int32 = -1
                                    v18 := []int{1, (^(fi13)), -228}
                                    if v1 {
                                        var v19 uint64 = 42
                                        _ = v19
                                    }
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                if (!false) {
                                    var v20 bool = false
                                    fmt.Println(v4, v20, v5)
                                    _ = v20
                                }
                            }
                        case "hello":
                            if ((42 != 127) && v1) {
                                switch fi13 {
                                case -1:
                                    if false {
                                        switch ("foo bar" + v6) {
                                        case "":
                                            switch 2147483647 {
                                            case 42:
                                                v21 := GenS2{F0: -51, F1: 1, F2: 737}
                                                _ = v21
                                            default:
                                                var v22 int16 = 1
                                                _ = v22
                                            }
                                        case "hello":
                                            v23 := []int{(fi7 << ((701) & 63)), (fi13 | 42), int(v5), fi13}
                                            _ = v23
                                        case "\n\t":
                                            v24 := []int{(fi7 ^ 127), v3}
                                            _ = v24
                                        default:
                                            v5 -= (^(v5))
                                        }
                                    } else {
                                        if false {
                                            var v25 uint = 1
                                            _ = v25
                                        } else {
                                            fmt.Println(v2)
                                            if true {
                                                if ((42 != 2) && true) {
                                                    v26 := []int{fi7, (v3 - -9223372036854775808)}
                                                    _ = v26
                                                }
                                                var v27 float32 = 3.14159
                                                fmt.Println(v2, v4)
                                                _ = v27
                                            }
                                            if (-76 <= 41) {
                                                if v1 {
                                                    if v1 {
                                                        v5 = v4
                                                    } else {
                                                        fmt.Println(v3, v1, v3)
                                                        v28 := []int{4294967295}
                                                        if v1 {
                                                            v5 = v4
                                                            var v29 int16 = -896
                                                            fmt.Println(v2, v29)
                                                            var v30 float32 = (((1e10 + (((((0.0 * (1.5 * 3.14159)) + float32(1e10)) * ((0.1 * -2.25) - 1.5)) / float32(v2)) + (0.1 - 0.0))) + 0.1) / 0.1)
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                        _ = v28
                                                    }
                                                    fmt.Println(v1, fi7, v2)
                                                } else {
                                                    if false {
                                                        var v31 int32 = 1
                                                        v32 := GenS2{F0: 22, F1: 65535, F2: (v31 * 0)}
                                                        v2 += (((v2 / (1e10 + (v2 - 1.5))) / 1.0) / float64(1.5))
                                                        v5 += 32767
                                                        var v33 float64 = v2
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        v1 = (!(!v1))
                                                    }
                                                    switch -1 {
                                                    case 128:
                                                        if true {
                                                            var v34 uint = 65535
                                                            if (42 > 24) {
                                                                v35 := []int{fi13}
                                                                v4 -= v5
                                                                switch (v6 + v6) {
                                                                case "hello":
                                                                    v36 := GenS1{F0: (v34 - (v34 - uint(v4)))}
                                                                    _ = v36
                                                                case "":
                                                                    if true {
                                                                        var v37 uint8 = 0
                                                                        var v38 string = "hello"
                                                                        var v39 int32 = 481
                                                                        _ = v37
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                case "a":
                                                                    _ = len(v35)
                                                                    switch v6 {
                                                                    case "\n\t":
                                                                        switch fi13 {
                                                                        case 32767:
                                                                            v6 = ("hello" + v6)
                                                                        case -240:
                                                                            v3 -= fi13
                                                                        default:
                                                                            v4 += 276
                                                                        }
                                                                    case "foo bar":
                                                                        var v40 bool = v1
                                                                        _ = v40
                                                                    }
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        switch fi7 {
                                                                        case 128:
                                                                            v34 = v34
                                                                        }
                                                                    case "foo bar":
                                                                        switch 2147483647 {
                                                                        case 1:
                                                                            v2 += v2
                                                                        case 128:
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                switch fi7 {
                                                                                case 2147483647:
                                                                                    if len(v35) > 0 {
                                                                                        v35[0] = (v3 % ((fi13) | 1))
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(v4, fi13, v3)
                                                                                case 127:
                                                                                    v1 = v1
                                                                                    v41 := GenS3{F0: v6, F1: -128, F2: v6}
                                                                                    var v42 uint64 = 4294967295
                                                                                    var v43 int8 = 127
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                    _ = v43
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    switch 32767 {
                                                                                    case -56:
                                                                                        if v1 {
                                                                                            var v44 int16 = -1
                                                                                            var v45 uint8 = 127
                                                                                            if (!false) {
                                                                                                var v46 uint32 = 311
                                                                                                var v47 uint32 = (v46 << ((255) & 63))
                                                                                                _ = v46
                                                                                                _ = v47
                                                                                            }
                                                                                            _ = v44
                                                                                            _ = v45
                                                                                        }
                                                                                    case 857:
                                                                                        fmt.Println(v4)
                                                                                    default:
                                                                                        v34 = uint(v4)
                                                                                        v48 := GenS4{F0: -1}
                                                                                        _ = len(v35)
                                                                                        var v49 int8 = -125
                                                                                        if ((-(v49)) <= v49) {
                                                                                            switch v3 {
                                                                                            case 9223372036854775807:
                                                                                                v50 := GenS5{F0: 1, F1: (fi7 % (((-(fi7))) | 1))}
                                                                                                _ = v50
                                                                                            case 2147483647:
                                                                                                var v51 int = v3
                                                                                                _ = v51
                                                                                            case 42:
                                                                                                switch (v3 / (((^(v3))) | 1)) {
                                                                                                case 4294967295:
                                                                                                    if (1 >= 127) {
                                                                                                        v52 := GenS2{F0: (v49 * 100), F1: 2147483647, F2: -1}
                                                                                                        _ = v52
                                                                                                    }
                                                                                                case -160:
                                                                                                    fmt.Println(v49)
                                                                                                default:
                                                                                                    v53 := GenS4{F0: -32768}
                                                                                                    _ = v53
                                                                                                }
                                                                                            }
                                                                                            var v54 string = ((v6 + ((("a" + (v6 + v6)) + "a") + v6)) + "a")
                                                                                            var v55 int = (fi13 % (((fi13 / (((fi7 - v3)) | 1))) | 1))
                                                                                            _ = v54
                                                                                            _ = v55
                                                                                        } else {
                                                                                            v49 = 1
                                                                                            v56 := GenS6{F0: -9223372036854775808, F1: 1, F2: 0}
                                                                                            if (("a" + "a") >= v6) {
                                                                                                if ((v34 < v34) && true) {
                                                                                                    var v57 uint8 = 0
                                                                                                    _ = v57
                                                                                                } else {
                                                                                                    var v58 uint8 = 128
                                                                                                    _ = v58
                                                                                                }
                                                                                                fmt.Println(v6, v4, v2)
                                                                                                fmt.Println(v1)
                                                                                            }
                                                                                            _ = v56
                                                                                        }
                                                                                        var v59 float32 = 3.14159
                                                                                        var v60 uint32 = 1
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    }
                                                                                }
                                                                                v5 = v5
                                                                                var v61 int = fi13
                                                                                if len(v35) > 0 {
                                                                                    v35[0] = (-(v61))
                                                                                }
                                                                                _ = v61
                                                                            }
                                                                            if v1 {
                                                                                v2 = v2
                                                                                if len(v35) >= 2 {
                                                                                    v35 = v35[1:len(v35)]
                                                                                }
                                                                                var v62 uint8 = 1
                                                                                _ = v62
                                                                            } else {
                                                                                switch (v3 % (((^(v3))) | 1)) {
                                                                                case 2147483647:
                                                                                    v2 += 0.1
                                                                                case -741:
                                                                                    var v63 int64 = 1
                                                                                    _ = v63
                                                                                case -9223372036854775808:
                                                                                    var v64 int64 = -64
                                                                                    if v1 {
                                                                                        v65 := GenS7{F0: fi7}
                                                                                        var v66 uint16 = (^(v4))
                                                                                        if (v34 > (^(v34))) {
                                                                                            var v67 uint32 = 2147483647
                                                                                            var v68 uint64 = 42
                                                                                            v64 -= v64
                                                                                            _ = v67
                                                                                            _ = v68
                                                                                        }
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    } else {
                                                                                        fmt.Println(v34)
                                                                                        v69 := GenS4{F0: -1}
                                                                                        _ = v69
                                                                                    }
                                                                                    _ = v64
                                                                                }
                                                                            }
                                                                            v35 = append(v35, (v3 * (fi13 | fi13)))
                                                                        case 0:
                                                                            var v70 float32 = 0.0
                                                                            _ = v70
                                                                        default:
                                                                            switch fi13 {
                                                                            case 4294967295:
                                                                                v71 := GenS8{F0: -1, F1: (fi7 * int(v34))}
                                                                                _ = v71
                                                                            default:
                                                                                v72 := GenS7{F0: fi7}
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                    default:
                                                                        v73 := GenS9{F0: 1, F1: 2}
                                                                        _ = v73
                                                                    }
                                                                default:
                                                                    var v74 int64 = 147
                                                                    _ = v74
                                                                }
                                                                fmt.Println(v2, v4, v2)
                                                                fmt.Println(v4)
                                                                var v75 int32 = 1
                                                                v76 := GenS10{F0: (v6 + v6), F1: v75}
                                                                if v1 {
                                                                    var v77 int8 = 42
                                                                    switch (((v6 + ("\n\t" + ("\n\t" + v6))) + "foo bar") + ("a" + v6)) {
                                                                    case "":
                                                                        fmt.Println(fi7, v3)
                                                                    case "a":
                                                                        v77 -= int8(fi13)
                                                                    case "hello":
                                                                        switch ("hello" + (("foo bar" + v6) + v6)) {
                                                                        case "foo bar":
                                                                            if false {
                                                                                var v78 bool = false
                                                                                _ = v78
                                                                            } else {
                                                                                switch -242 {
                                                                                case -516:
                                                                                    v3 -= (v3 * (v3 & (-(fi7))))
                                                                                case -9223372036854775808:
                                                                                    switch v3 {
                                                                                    case 1:
                                                                                        v79 := GenS11{F0: v77, F1: v6, F2: 265}
                                                                                        _ = v79
                                                                                    case -591:
                                                                                        if len(v35) >= 2 {
                                                                                            v35 = v35[1:len(v35)]
                                                                                        }
                                                                                    case 733:
                                                                                        v80 := GenS7{F0: (fi7 | (^(fi7)))}
                                                                                        _ = v80
                                                                                    }
                                                                                case -735:
                                                                                    if len(v35) > 0 {
                                                                                        v35[0] = (fi7 | 627)
                                                                                    }
                                                                                    v81 := GenS12{F0: 42, F1: (v75 << (((v34 << ((v34) & 63))) & 63))}
                                                                                    _ = v81
                                                                                }
                                                                                v35 = append(v35, fi7)
                                                                                v82 := GenS6{F0: 0, F1: 0, F2: -1}
                                                                                _ = v82
                                                                            }
                                                                            fmt.Println(v1, fi13)
                                                                        case "":
                                                                            if true {
                                                                                v83 := GenS13{F0: uint64(v34), F1: 32767}
                                                                                var v84 int64 = int64(v75)
                                                                                switch v6 {
                                                                                case "":
                                                                                    var v85 int64 = int64(v75)
                                                                                    _ = v85
                                                                                }
                                                                                var v86 bool = (((0.1 * v2) + (v2 - 1e10)) == 1e10)
                                                                                v2 -= v2
                                                                                _ = v83
                                                                                _ = v84
                                                                                _ = v86
                                                                            }
                                                                        default:
                                                                            v34 -= v34
                                                                        }
                                                                    }
                                                                    _ = v77
                                                                }
                                                                _ = v35
                                                                _ = v75
                                                                _ = v76
                                                            } else {
                                                                v3 = (v3 >> ((4294967295) & 63))
                                                                if (0.1 < ((v2 * v2) * v2)) {
                                                                    if false {
                                                                        switch fi13 {
                                                                        case 128:
                                                                            v87 := []int{fi7, fi13}
                                                                            _ = v87
                                                                        case -640:
                                                                            if (127 != 23) {
                                                                                switch 42 {
                                                                                case 932:
                                                                                    v2 += float64(0.1)
                                                                                }
                                                                                v2 -= v2
                                                                            }
                                                                        default:
                                                                            v2 += v2
                                                                        }
                                                                        switch 32768 {
                                                                        case 913:
                                                                            fmt.Println(v6, v1)
                                                                        case -71:
                                                                            var v88 bool = (-608 >= v3)
                                                                            _ = v88
                                                                        case 697:
                                                                            fmt.Println(v4)
                                                                        }
                                                                        v89 := GenS7{F0: (fi13 - (fi13 ^ (v3 >> (((v34 ^ (^(v34)))) & 63))))}
                                                                        _ = v89
                                                                    } else {
                                                                        v2 = ((float64(1.5) / float64((1e10 / 1.0))) * (((v2 * v2) * 1e10) * 3.14159))
                                                                        switch (v6 + "foo bar") {
                                                                        case "hello":
                                                                            fmt.Println(fi7, fi7)
                                                                        default:
                                                                            switch 0 {
                                                                            case 702:
                                                                                var v90 uint8 = 2
                                                                                _ = v90
                                                                            case 9223372036854775807:
                                                                                var v91 int = (fi7 & fi7)
                                                                                _ = v91
                                                                            case 127:
                                                                                var v92 uint8 = 2
                                                                                _ = v92
                                                                            }
                                                                        }
                                                                    }
                                                                    v5 += (v5 ^ 32767)
                                                                    var v93 int8 = -128
                                                                    _ = v93
                                                                }
                                                                v3 = 1
                                                            }
                                                            v94 := GenS3{F0: v6, F1: 82, F2: v6}
                                                            _ = v34
                                                            _ = v94
                                                        } else {
                                                            v2 += v2
                                                            switch (v6 + v6) {
                                                            case "":
                                                                var v95 int8 = 53
                                                                _ = v95
                                                            case "a":
                                                                v96 := []int{(fi7 / ((fi13) | 1)), (v3 + fi13), (fi7 + (fi7 % (((-(fi7))) | 1))), fi7}
                                                                _ = v96
                                                            case "hello":
                                                                v97 := GenS14{F0: (v5 << ((626) & 63)), F1: 0}
                                                                _ = v97
                                                            default:
                                                                var v98 int = v3
                                                                _ = v98
                                                            }
                                                            v99 := []int{128, (fi7 & v3), (v3 | (fi7 & (v3 & (fi13 ^ (^(fi7)))))), -9223372036854775808}
                                                            if len(v99) >= 2 {
                                                                v99 = v99[1:len(v99)]
                                                            }
                                                            if v1 {
                                                                var v100 bool = (-622 != -799)
                                                                _ = v100
                                                            }
                                                            _ = v99
                                                        }
                                                    }
                                                }
                                                v101 := GenS14{F0: 65535, F1: 127}
                                                _ = v101
                                            } else {
                                                var v102 uint8 = 2
                                                _ = v102
                                            }
                                        }
                                        v5 = (^(v4))
                                    }
                                case 4294967295:
                                    if v1 {
                                        v103 := []int{}
                                        if len(v103) >= 2 {
                                            v103 = v103[1:len(v103)]
                                        }
                                        _ = len(v103)
                                        _ = v103
                                    }
                                case -2147483648:
                                    var v104 uint16 = v5
                                    _ = v104
                                default:
                                    if (!(-952 >= -89)) {
                                        v105 := GenS13{F0: 42, F1: -650}
                                        v106 := GenS15{F0: ((v6 > "\n\t") || ((v4 / ((0) | 1)) == 1)), F1: 0}
                                        if ((v4 - 31) != (^(v4))) {
                                            switch fi7 {
                                            case 32767:
                                                v107 := []int{(v3 ^ (v3 - fi7)), (v3 >> ((42) & 63)), v3}
                                                var v108 string = v6
                                                var v109 uint64 = 18446744073709551615
                                                _ = v107
                                                _ = v108
                                                _ = v109
                                            case 4294967295:
                                                if (0.0 < (((0.0 / v2) + 3.14159) * (v2 - 0.1))) {
                                                    var v110 int64 = 1
                                                    v1 = (219 == 1)
                                                    _ = v110
                                                } else {
                                                    fmt.Println(v1, v4)
                                                }
                                            case 32768:
                                                v111 := []int{fi7}
                                                _ = v111
                                            default:
                                                v6 = v6
                                            }
                                        } else {
                                            v6 = "a"
                                            fmt.Println(v4, fi13)
                                            var v112 int64 = 206
                                            _ = v112
                                        }
                                        _ = v105
                                        _ = v106
                                    }
                                }
                            }
                            v3 -= (fi13 - fi13)
                            var v113 float32 = (-2.25 + 3.14159)
                            _ = v113
                        default:
                            fmt.Println(fi13, v1)
                        }
                        v1 = (!(v5 == 42))
                        _ = fi13
                    }
                }
            }
        default:
            v114 := GenS11{F0: 0, F1: (((("" + (((v6 + v6) + (v6 + "\n\t")) + v6)) + v6) + "a") + "foo bar"), F2: 762}
            _ = v114
        }
        _ = fi7
    }
    if false {
        var v115 bool = true
        v4 = (v5 << ((1) & 63))
        var v116 int8 = 36
        _ = v115
        _ = v116
    }
    fmt.Println(v5, v3, v6)
    var v117 int = (-(v3))
    v6 = ""
    for fi118 := 0; fi118 < 8; fi118++ {
        switch v6 {
        case "":
            var v119 uint16 = (v5 | v4)
            _ = v119
        case "a":
            var v120 int16 = -1
            _ = v120
        case "foo bar":
            v121 := []int{4294967295, (fi118 | (fi118 / ((v3) | 1))), v3, (^(v3))}
            _ = v121
        default:
            var v122 int8 = 100
            _ = v122
        }
        var v123 float32 = 3.14159
        v3 = (fi118 << ((74) & 63))
        _ = fi118
        _ = v123
    }
    fmt.Println(v6, v6, v117)
    v2 = v2
    v124 := []int{v117, v117, (v3 >> ((849) & 63))}
    for fi125 := 0; fi125 < 15; fi125++ {
        v126 := GenS16{F0: 73, F1: 0, F2: 671}
        fmt.Println(v5)
        _ = fi125
        _ = v126
    }
    v117 -= (v117 % (((-(v3))) | 1))
    v124 = append(v124, (v117 % ((v3) | 1)))
    v124 = append(v124, -1)
    if false {
        fmt.Println(v117, v117, v6)
        var v127 uint16 = 65535
        v128 := GenS17{F0: (v117 / ((0) | 1)), F1: 0}
        _ = v127
        _ = v128
    } else {
        if len(v124) > 0 {
            v124[0] = 9223372036854775807
        }
    }
    var v129 string = ((((((((v6 + ((v6 + (v6 + "")) + "")) + v6) + (v6 + v6)) + "hello") + ((v6 + v6) + v6)) + "a") + ("a" + v6)) + "\n\t")
    v124 = append(v124, v117)
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("slicelen v124:", len(v124), v124)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v117
    _ = v124
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
