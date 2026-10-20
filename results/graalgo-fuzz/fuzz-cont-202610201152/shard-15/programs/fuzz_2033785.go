package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 uint8 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint8; F1 int8; F2 int64 }
type GenS9 struct { F0 int32; F1 string; F2 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint32; F1 int32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 string; F1 int64; F2 uint16 }
type GenS15 struct { F0 int32; F1 uint64; F2 int16 }
type GenS16 struct { F0 uint8; F1 uint32; F2 int64 }

func main() {
    var v1 int = 9223372036854775807
    var v2 float64 = (((3.14159 - 0.1) / 1.0) * (((3.14159 * 1e10) / 1e10) + 0.1))
    v3 := make(map[string]int)
    var v4 uint64 = 1
    v5 := make(map[uint32]int)
    var v6 float32 = ((3.14159 * ((1.5 * ((1.5 + 0.0) / 1.0)) + 1.5)) / 3.14159)
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 42, F1: false}
        if (false || (397 >= 42)) {
            if (false || true) {
                v6 = 1.5
                v9 := []int{(fi7 + v1), 128, (fi7 + (v1 & 42))}
                var v10 int64 = -1
                _ = v10
                _ = v9
            }
            switch v1 {
            case 1:
                v6 = 1e10
            case 4294967295:
                fmt.Println(fi7, v1, v6)
            }
        }
        fmt.Println(v4, v1)
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: 32767, F1: true}
        switch "foo bar" {
        case "foo bar":
            v13 := GenS1{F0: 485, F1: true}
            _ = v13
        }
        for fi14 := 0; fi14 < 5; fi14++ {
            if (1e10 != ((v2 / 1.5) / (float64(v6) / v2))) {
                v2 -= float64(v6)
                v1 = (-(fi11))
                fmt.Println(v2, v6, v2)
            }
            v2 = (0.0 * 1e10)
            if (2147483647 > 42) {
                v15 := GenS1{F0: 42, F1: false}
                var v16 int64 = 9223372036854775807
                v2 += v2
                _ = v15
                _ = v16
            } else {
                if ((true || false) || (32767 <= 42)) {
                    if (2 <= 42) {
                        if true {
                            if true {
                                v17 := []int{}
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                switch "hello" {
                                case "\n\t":
                                    if false {
                                        v18 := GenS1{F0: 1000, F1: false}
                                        if len(v17) > 0 {
                                            v17[0] = (fi14 + 9223372036854775807)
                                        }
                                        v17 = append(v17, (v1 / ((fi11) | 1)))
                                        _ = v18
                                    }
                                }
                                _ = len(v17)
                                _ = v17
                            }
                        }
                        v19 := GenS2{F0: 127, F1: "a"}
                        var v20 int16 = -1
                        _ = v19
                        _ = v20
                    } else {
                        v21 := []int{fi11}
                        var v22 int8 = 1
                        _ = v21
                        _ = v22
                    }
                    if true {
                        var v23 int16 = -1
                        if (true || (1 <= 278)) {
                            v2 += v2
                            if false {
                                switch -606 {
                                case 42:
                                    v1 = fi14
                                }
                            } else {
                                if ((v1 + v1) < (v1 + v1)) {
                                    var v24 uint = 0
                                    fmt.Println(v24, v6, v24)
                                    _ = v24
                                } else {
                                    v25 := GenS2{F0: 0, F1: "hello"}
                                    var v26 uint32 = 2
                                    fmt.Println(v26)
                                    fmt.Println(v6, v26, v4)
                                    v6 -= v6
                                    switch "hello" {
                                    case "":
                                        var v27 int8 = -1
                                        _ = v27
                                    case "\n\t":
                                        v1 -= (^(v1))
                                        fmt.Println(v2, fi11, fi14)
                                    default:
                                        if (!((!(9223372036854775807 == -212)) || false)) {
                                            fmt.Println(v6)
                                            v23 = (v23 % ((v23) | 1))
                                            var v28 uint8 = 247
                                            v26 = 65535
                                            _ = v28
                                        }
                                    }
                                    var v29 int16 = (v23 >> ((0) & 63))
                                    if (2147483647 < fi11) {
                                        v30 := GenS2{F0: 127, F1: "hello"}
                                        _ = v30
                                    } else {
                                        var v31 string = "hello"
                                        var v32 uint16 = 32767
                                        v33 := GenS3{F0: 57, F1: (-(v29)), F2: v26}
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v25
                                    _ = v26
                                    _ = v29
                                }
                            }
                            if true {
                                if true {
                                    v5[489] = (fi11 / ((4294967295) | 1))
                                    v34 := GenS2{F0: 1, F1: ("" + "foo bar")}
                                    var v35 int8 = -128
                                    _ = v34
                                    _ = v35
                                }
                                v23 -= 42
                                var v36 uint8 = 42
                                _ = v36
                            }
                        }
                        switch "" {
                        case "foo bar":
                            if false {
                                if true {
                                    if false {
                                        var v37 int16 = v23
                                        if (!false) {
                                            v4 -= 1
                                            var v38 int16 = 547
                                            _ = v38
                                        }
                                        v6 += (((float32(v2) / (1.5 * float32(1e10))) * v6) * v6)
                                        if true {
                                            v37 -= v23
                                            if (true && (("hello" != ("foo bar" + "a")) || false)) {
                                                v2 = (v2 / ((v2 - (v2 / -2.25)) - v2))
                                            }
                                            v4 = v4
                                        } else {
                                            if false {
                                                if false {
                                                    v2 -= ((v2 + 3.14159) - (v2 / 0.1))
                                                } else {
                                                    v37 = (v23 % (((-(v37))) | 1))
                                                    var v39 uint64 = (v4 % ((v4) | 1))
                                                    v4 = 4294967295
                                                    _ = v39
                                                }
                                            }
                                            v23 -= -1
                                        }
                                        v4 -= v4
                                        _ = v37
                                    } else {
                                        v23 = v23
                                        v2 += (-2.25 + 0.0)
                                        if (0 > 1) {
                                            v40 := GenS4{F0: 968, F1: 1}
                                            if false {
                                                var v41 uint16 = 0
                                                v1 -= (fi14 + fi11)
                                                _ = v41
                                            }
                                            v23 = v23
                                            _ = v40
                                        } else {
                                            v42 := v5[2147483647]
                                            _ = v42
                                        }
                                        v1 += (v1 * (fi14 * 32767))
                                        v23 = (v23 >> ((547) & 63))
                                        if (226 <= 2) {
                                            fmt.Println(v1, fi11)
                                            fmt.Println(v1, fi11, v6)
                                            switch -2147483648 {
                                            case 674:
                                                var v43 uint32 = 246
                                                _ = v43
                                            case 4294967295:
                                                v23 = v23
                                            case 32768:
                                                v44 := GenS3{F0: 250, F1: v23, F2: 2147483647}
                                                _ = v44
                                            }
                                        } else {
                                            var v45 uint = 906
                                            _ = v45
                                        }
                                    }
                                } else {
                                    if (187 == 42) {
                                        v46 := GenS1{F0: (v23 - v23), F1: false}
                                        if (!false) {
                                            if ((!true) || (false && (false && true))) {
                                                v5[4294967295] = (fi11 + (fi11 ^ v1))
                                                if (65535 == 216) {
                                                    v2 += v2
                                                }
                                            } else {
                                                if true {
                                                    switch -2147483648 {
                                                    case 2147483647:
                                                        var v47 float32 = float32(float64((0.0 * v6)))
                                                        _ = v47
                                                    case -1:
                                                        v48 := GenS1{F0: 0, F1: true}
                                                        _ = v48
                                                    }
                                                    delete(v5, 675)
                                                    v1 = -9223372036854775808
                                                    switch "foo bar" {
                                                    case "":
                                                        v1 += (fi14 % ((966) | 1))
                                                    default:
                                                        switch (v1 ^ (v1 * 42)) {
                                                        case 42:
                                                            if ((743 <= 255) && true) {
                                                                var v49 uint = 651
                                                                var v50 uint32 = 65535
                                                                v4 = v4
                                                                fmt.Println(v6, fi14)
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            switch v1 {
                                                            case 353:
                                                                var v51 uint32 = 2147483648
                                                                _ = v51
                                                            }
                                                            v6 = v6
                                                        case 4294967295:
                                                            var v52 int16 = v23
                                                            _ = v52
                                                        }
                                                    }
                                                    v53 := []int{-9223372036854775808}
                                                    _ = v53
                                                } else {
                                                    if ((1e10 > ((v6 + (-2.25 * ((v6 + v6) + v6))) - 0.0)) && false) {
                                                        if true {
                                                            delete(v3, "a")
                                                            v54 := GenS4{F0: 788, F1: -1}
                                                            _ = v54
                                                        }
                                                        fmt.Println(v6)
                                                    }
                                                }
                                                switch (v1 >> ((42) & 63)) {
                                                case -9223372036854775808:
                                                    var v55 float64 = (0.1 - (-2.25 / v2))
                                                    _ = v55
                                                case 647:
                                                    v56 := GenS5{F0: 2, F1: 1}
                                                    _ = v56
                                                default:
                                                    if true {
                                                        var v57 uint32 = 0
                                                        switch (v1 + 128) {
                                                        case 4294967295:
                                                            v1 += fi11
                                                        case 127:
                                                            var v58 uint32 = (v57 & 57)
                                                            _ = v58
                                                        }
                                                        _ = v57
                                                    } else {
                                                        if true {
                                                            if false {
                                                                switch (fi14 & (fi14 / (((fi11 | 38)) | 1))) {
                                                                case 2147483647:
                                                                    if ("foo bar" == "hello") {
                                                                        switch fi14 {
                                                                        case 4294967295:
                                                                            v59 := GenS2{F0: 0, F1: "hello"}
                                                                            _ = v59
                                                                        case -921:
                                                                            v1 += 0
                                                                        case 0:
                                                                            if (4294967295 == 0) {
                                                                                var v60 uint8 = 42
                                                                                _ = v60
                                                                            }
                                                                        default:
                                                                            v61 := GenS3{F0: 127, F1: (v23 >> ((2) & 63)), F2: 4294967295}
                                                                            _ = v61
                                                                        }
                                                                    } else {
                                                                        var v62 uint8 = 1
                                                                        _ = v62
                                                                    }
                                                                    v6 = (((v6 / float32(v2)) + v6) / v6)
                                                                    v63 := []int{(fi11 ^ (fi14 >> ((4294967295) & 63)))}
                                                                    _ = v63
                                                                case -599:
                                                                    v64 := GenS6{F0: (v23 - v23), F1: 2147483647, F2: 0}
                                                                    _ = v64
                                                                }
                                                                var v65 uint64 = (v4 - 1)
                                                                v6 = (v6 - 3.14159)
                                                                _ = v65
                                                            } else {
                                                                v6 -= (((((1.5 * v6) + float32((v2 - v2))) / v6) / (0.1 / float32((v2 - v2)))) + v6)
                                                                fmt.Println(v23)
                                                            }
                                                            if true {
                                                                fmt.Println(fi14)
                                                            }
                                                            v23 = v23
                                                        } else {
                                                            var v66 string = ("hello" + ("hello" + "hello"))
                                                            v67 := GenS3{F0: 20, F1: v23, F2: 42}
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        if false {
                                                            var v68 uint8 = 242
                                                            v1 += fi11
                                                            switch (("" + ("hello" + "")) + "hello") {
                                                            case "hello":
                                                                switch -444 {
                                                                case -9223372036854775808:
                                                                    if (1e10 > 1.5) {
                                                                        var v69 uint8 = v68
                                                                        v4 = v4
                                                                        _ = v69
                                                                    } else {
                                                                        v4 = v4
                                                                    }
                                                                case 50:
                                                                    if ((!("" != "")) || true) {
                                                                        fmt.Println(v4, fi14, v2)
                                                                        var v70 float64 = (v2 / v2)
                                                                        fmt.Println(v68, fi11)
                                                                        _ = v70
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v23, v6, v2)
                                                            default:
                                                                var v71 string = "\n\t"
                                                                _ = v71
                                                            }
                                                            _ = v68
                                                        } else {
                                                            v6 = float32(v2)
                                                            var v72 uint64 = 0
                                                            var v73 int32 = 1
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        v4 += v4
                                                    }
                                                }
                                                fmt.Println(v4)
                                            }
                                            switch ("foo bar" + ("a" + ("a" + "hello"))) {
                                            case "hello":
                                                if (128 != 127) {
                                                    v2 -= float64(v6)
                                                    if true {
                                                        if (406 == 2147483647) {
                                                            if (335 == -533) {
                                                                v23 -= 42
                                                                var v74 float64 = v2
                                                                _ = v74
                                                            } else {
                                                                switch v1 {
                                                                case -32:
                                                                    v4 -= v4
                                                                case 530:
                                                                    v75 := GenS1{F0: (v23 * (v23 + 1000)), F1: (1 > 42)}
                                                                    _ = v75
                                                                default:
                                                                    if false {
                                                                        if false {
                                                                            if true {
                                                                                if ((3.14159 != v6) && (!true)) {
                                                                                    var v76 uint8 = 255
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        var v77 float64 = 1e10
                                                                                        _ = v77
                                                                                    default:
                                                                                        v78 := []int{-699, -1, v1}
                                                                                        _ = v78
                                                                                    }
                                                                                    _ = v76
                                                                                }
                                                                            } else {
                                                                                var v79 int8 = 100
                                                                                v2 -= v2
                                                                                var v80 int8 = v79
                                                                                _ = v79
                                                                                _ = v80
                                                                            }
                                                                        } else {
                                                                            var v81 uint64 = v4
                                                                            _ = v81
                                                                        }
                                                                        if (244 < 255) {
                                                                            switch fi14 {
                                                                            case 2147483647:
                                                                                v82 := []int{}
                                                                                v1 = (fi11 % ((v1) | 1))
                                                                                v82 = append(v82, v1)
                                                                                _ = v82
                                                                            }
                                                                            if (((float32(-2.25) != ((v6 - v6) + (v6 / v6))) && ("\n\t" != "a")) && false) {
                                                                                var v83 int16 = (v23 - v23)
                                                                                _ = v83
                                                                            } else {
                                                                                if (!(false || false)) {
                                                                                    v84 := []int{fi11, int(v23), fi14, (fi11 / ((fi11) | 1))}
                                                                                    var v85 string = "\n\t"
                                                                                    v84 = append(v84, v1)
                                                                                    _ = v84
                                                                                    _ = v85
                                                                                } else {
                                                                                    delete(v3, "")
                                                                                    var v86 uint = 4294967295
                                                                                    var v87 uint32 = 4294967295
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                }
                                                                                if false {
                                                                                    v88 := GenS2{F0: 1, F1: "hello"}
                                                                                    _ = v88
                                                                                } else {
                                                                                    _ = len(v5)
                                                                                }
                                                                                fmt.Println(fi14)
                                                                            }
                                                                            var v89 uint16 = 443
                                                                            _ = v89
                                                                        }
                                                                        v6 += v6
                                                                    }
                                                                }
                                                                if (255 <= 42) {
                                                                    fmt.Println(v6, v1)
                                                                    v1 = v1
                                                                    if true {
                                                                        switch (fi14 >> ((0) & 63)) {
                                                                        case -15:
                                                                            v90 := GenS1{F0: (v23 + v23), F1: false}
                                                                            _ = v90
                                                                        default:
                                                                            v91 := GenS7{F0: 42, F1: 127}
                                                                            _ = v91
                                                                        }
                                                                    } else {
                                                                        var v92 string = ("" + "hello")
                                                                        v93 := []int{(fi14 - (fi11 | v1)), fi11}
                                                                        _ = v92
                                                                        _ = v93
                                                                    }
                                                                    v6 += -2.25
                                                                    var v94 uint = 65535
                                                                    fmt.Println(v23, v2)
                                                                    if (!((fi14 <= -1) || (true || false))) {
                                                                        v95 := GenS6{F0: -1, F1: 1, F2: 255}
                                                                        v23 += (v23 >> ((v94) & 63))
                                                                        var v96 int64 = 0
                                                                        _ = v95
                                                                        _ = v96
                                                                    }
                                                                    _ = v94
                                                                } else {
                                                                    if (v6 != v6) {
                                                                        fmt.Println(v23, v1)
                                                                        fmt.Println(v4, v6)
                                                                    } else {
                                                                        var v97 int64 = -1
                                                                        var v98 int32 = -1
                                                                        delete(v5, 4294967295)
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                }
                                                                var v99 string = "foo bar"
                                                                _ = v99
                                                            }
                                                            fmt.Println(v6, v6, fi11)
                                                        } else {
                                                            var v100 int8 = -1
                                                            var v101 uint8 = 255
                                                            v100 += (-(v100))
                                                            _ = v100
                                                            _ = v101
                                                        }
                                                    } else {
                                                        var v102 uint64 = v4
                                                        v103 := GenS8{F0: 128, F1: 2, F2: 42}
                                                        v1 += 669
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    v104 := []int{(fi14 + 4294967295)}
                                                    v105 := GenS7{F0: 42, F1: 127}
                                                    if false {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if true {
                                                                v4 = v4
                                                                delete(v5, 65535)
                                                                v106 := GenS9{F0: -861, F1: "", F2: (fi14 * v1)}
                                                                v6 = 1.5
                                                                if false {
                                                                    v1 = v1
                                                                } else {
                                                                    v4 = v4
                                                                    _ = len(v104)
                                                                    var v107 string = "foo bar"
                                                                    var v108 int8 = 1
                                                                    _ = v107
                                                                    _ = v108
                                                                }
                                                                fmt.Println(v4, v6)
                                                                if true {
                                                                    var v109 int64 = 9223372036854775807
                                                                    fmt.Println(v6, v23)
                                                                    _ = v109
                                                                } else {
                                                                    var v110 int = 127
                                                                    fmt.Println(v1, v1)
                                                                    v2 = v2
                                                                    _ = v110
                                                                }
                                                                var v111 int64 = -178
                                                                _ = v106
                                                                _ = v111
                                                            }
                                                        default:
                                                            v4 += 1
                                                        }
                                                        var v112 float64 = (((v2 - v2) / ((v2 * float64((v6 - v6))) / (float64(v6) - (v2 + (v2 * 3.14159))))) / float64(0.1))
                                                        _ = v112
                                                    }
                                                    _ = v104
                                                    _ = v105
                                                }
                                            case "foo bar":
                                                if ((!(107 <= -1)) && true) {
                                                    v113 := []int{(v1 & (fi14 << ((65535) & 63))), fi14, fi14, (fi14 * -728)}
                                                    var v114 uint8 = 242
                                                    _ = v113
                                                    _ = v114
                                                } else {
                                                    v115 := GenS3{F0: 127, F1: -32768, F2: 2147483648}
                                                    var v116 int64 = 1
                                                    v5[1] = (fi14 ^ v1)
                                                    var v117 int16 = v23
                                                    v118 := GenS5{F0: 2, F1: 1}
                                                    v119 := GenS10{F0: 42}
                                                    _ = v115
                                                    _ = v116
                                                    _ = v117
                                                    _ = v118
                                                    _ = v119
                                                }
                                            }
                                            v5[1] = v1
                                        } else {
                                            var v120 uint16 = 2
                                            v121 := []int{0, -469, (v1 | fi14)}
                                            _ = v120
                                            _ = v121
                                        }
                                        _ = v46
                                    } else {
                                        var v122 int8 = -77
                                        _ = v122
                                    }
                                    v123 := GenS11{F0: 1}
                                    var v124 int = fi14
                                    _ = v123
                                    _ = v124
                                }
                            }
                        }
                        _ = v23
                    } else {
                        fmt.Println(v6)
                    }
                    var v125 uint64 = v4
                    var v126 int8 = 42
                    v3["\n\t"] = fi14
                    _ = v125
                    _ = v126
                }
            }
            _ = fi14
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v6, v6)
    fmt.Println(v6, v2, v6)
    for fi127 := 0; fi127 < 15; fi127++ {
        for fi128 := 0; fi128 < 19; fi128++ {
            v4 += 0
            if (1 < 907) {
                var v129 uint16 = 2
                fmt.Println(v1, fi127, v4)
                var v130 float64 = 0.1
                _ = v129
                _ = v130
            }
            fmt.Println(fi127, fi128)
            _ = fi128
        }
        switch fi127 {
        case 42:
            for fi131 := 0; fi131 < 13; fi131++ {
                v2 += float64(3.14159)
                v3[("a" + ("foo bar" + (((("hello" + ("" + "\n\t")) + ("" + "a")) + "\n\t") + ("\n\t" + "\n\t"))))] = (fi131 % ((-188) | 1))
                switch (("a" + "\n\t") + "hello") {
                case "\n\t":
                    if false {
                        if true {
                            v2 = v2
                            v132 := GenS9{F0: 563, F1: "\n\t", F2: fi127}
                            if (!(-9223372036854775808 > -519)) {
                                var v133 float64 = v2
                                _ = v133
                            }
                            var v134 uint32 = 42
                            switch ("\n\t" + "hello") {
                            case "":
                                var v135 int8 = -128
                                _ = v135
                            case "a":
                                v136 := GenS9{F0: -2147483648, F1: (("hello" + ("foo bar" + "hello")) + "\n\t"), F2: v1}
                                _ = v136
                            case "\n\t":
                                var v137 int8 = -7
                                _ = v137
                            }
                            v134 += 817
                            if true {
                                v138 := []int{32768, 32767, fi131, 4294967295}
                                switch "" {
                                case "hello":
                                    var v139 int64 = 0
                                    _ = v139
                                case "":
                                    if true {
                                        fmt.Println(fi127)
                                    }
                                default:
                                    v138 = append(v138, 1)
                                }
                                _ = v138
                            } else {
                                v4 = 42
                                if false {
                                    v1 += fi127
                                }
                            }
                            if false {
                                fmt.Println(v134, v2, fi131)
                                v140 := []int{906}
                                _ = v140
                            } else {
                                switch fi131 {
                                case 535:
                                    var v141 uint8 = 128
                                    _ = v141
                                case -579:
                                    var v142 float32 = v6
                                    _ = v142
                                case 1:
                                    v4 -= v4
                                default:
                                    v143 := GenS8{F0: 164, F1: -122, F2: -9223372036854775808}
                                    if false {
                                        fmt.Println(v2, v6)
                                        v134 = v134
                                    } else {
                                        v2 += 3.14159
                                        var v144 uint16 = 65535
                                        v2 = float64(float32(v2))
                                        _ = v144
                                    }
                                    v145 := GenS11{F0: 1}
                                    _ = v143
                                    _ = v145
                                }
                            }
                            v134 = v134
                            var v146 int = fi131
                            if false {
                                var v147 float64 = 3.14159
                                _ = v147
                            }
                            v148 := GenS12{F0: (v134 ^ 2), F1: 2147483647}
                            if true {
                                v146 = (fi131 ^ (fi131 - v1))
                                var v149 uint32 = (v134 - 2147483648)
                                _ = v149
                            } else {
                                var v150 int32 = -579
                                _ = v150
                            }
                            v146 = -9223372036854775808
                            if true {
                                v146 = v1
                                if false {
                                    if true {
                                        v151 := GenS13{F0: 2147483647}
                                        if true {
                                            var v152 int8 = 0
                                            v134 += (^(v134))
                                            v153 := v3["a"]
                                            _ = v152
                                            _ = v153
                                        } else {
                                            var v154 int = v146
                                            var v155 int32 = 2147483647
                                            fmt.Println(fi127, v154)
                                            _ = v154
                                            _ = v155
                                        }
                                        _ = v151
                                    } else {
                                        var v156 uint8 = 0
                                        _ = v156
                                    }
                                    var v157 int = (^(fi127))
                                    v158 := GenS3{F0: 2, F1: 359, F2: uint32(v4)}
                                    var v159 int16 = 1
                                    var v160 float32 = (0.0 + v6)
                                    _ = v157
                                    _ = v158
                                    _ = v159
                                    _ = v160
                                }
                                v146 = int(v4)
                            } else {
                                if true {
                                    var v161 int16 = 42
                                    var v162 int8 = 66
                                    _ = v161
                                    _ = v162
                                }
                                var v163 int32 = 1
                                v164 := GenS14{F0: "hello", F1: 9223372036854775807, F2: 255}
                                _ = v163
                                _ = v164
                            }
                            _ = v132
                            _ = v134
                            _ = v146
                            _ = v148
                        }
                        if true {
                            v165 := []int{(v1 >> ((2) & 63))}
                            if true {
                                var v166 int32 = -2147483648
                                if true {
                                    v6 = v6
                                }
                                var v167 uint32 = 2147483648
                                _ = v166
                                _ = v167
                            }
                            v2 += 3.14159
                            var v168 int8 = -128
                            _ = v165
                            _ = v168
                        } else {
                            fmt.Println(v1)
                            switch "\n\t" {
                            case "":
                                v169 := GenS9{F0: -2147483648, F1: ("\n\t" + "foo bar"), F2: 2147483647}
                                _ = v169
                            case "a":
                                var v170 int16 = 32767
                                _ = v170
                            }
                            v4 += v4
                            v171 := []int{(-(v1)), fi127, (^(fi131))}
                            _ = v171
                        }
                        fmt.Println(v2, v1)
                    }
                case "hello":
                    v172 := []int{(v1 * fi127), (v1 / (((fi131 % ((v1) | 1))) | 1))}
                    _ = v172
                case "a":
                    var v173 uint64 = (^(v4))
                    _ = v173
                }
                _ = fi131
            }
        case 4294967295:
            var v174 int = fi127
            _ = v174
        default:
            v2 += v2
        }
        if true {
            if false {
                v175 := GenS15{F0: -839, F1: 18446744073709551615, F2: 32767}
                v6 = (-2.25 + v6)
                _ = v175
            }
            v176 := GenS16{F0: 255, F1: 1, F2: 9223372036854775807}
            _ = v176
        } else {
            var v177 float32 = (((((((v6 - v6) / (0.0 / v6)) - 0.1) / 1.0) / 1.0) + v6) - ((float32(0.1) - v6) / 1.0))
            _ = v177
        }
        _ = fi127
    }
    switch ("" + "") {
    case "foo bar":
        fmt.Println(v1, v2, v1)
    default:
        var v178 uint8 = 100
        _ = v178
    }
    v1 = v1
    var v179 float32 = 1.5
    if (true || (259 == 741)) {
        v6 -= 1.5
        v1 += v1
        if (42 > int32(v1)) {
            v5[0] = (v1 - v1)
        }
    }
    var v180 int16 = 1000
    v5[2147483648] = (-(v1))
    v181 := GenS16{F0: 255, F1: 4294967295, F2: -1}
    v179 -= (1e10 / 1.5)
    v182 := []int{(v1 >> ((547) & 63)), (v1 >> ((18446744073709551615) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v179:", v179)
    fmt.Println("v180:", v180)
    fmt.Println("slicelen v182:", len(v182), v182)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v179
    _ = v180
    _ = v181
    _ = v182
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
