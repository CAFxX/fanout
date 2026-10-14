package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint; F1 int16; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 uint; F1 uint8; F2 int64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8; F1 int64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint; F1 uint8; F2 string }
type GenS18 struct { F0 int8; F1 bool; F2 uint16 }
type GenS19 struct { F0 int32; F1 int8; F2 int16 }
type GenS20 struct { F0 uint32; F1 uint16 }
type GenS21 struct { F0 string; F1 int64; F2 int32 }
func genFn0(p0 int8) string {
    p0 = p0
    if true {
        v4 := GenS1{F0: 1000}
        _ = v4
    } else {
        v5 := []int{0, -669, -636}
        _ = v5
    }
    if (false && (((996 > 18446744073709551615) || true) && ("" <= "a"))) {
        var v6 int64 = -1
        _ = v6
    } else {
        p0 += -1
        p0 += (p0 % ((p0) | 1))
    }
    _ = p0
    return ("a" + "\n\t")
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := []uint32{465, 1, 2147483647}
    var v7 uint = 90
    v8 := GenS1{F0: 1000}
    v7 += 555
    v7 -= (v7 - (v7 >> ((v7) & 63)))
    fmt.Println(v7)
    var v9 string = "foo bar"
    var v10 int = 2147483647
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch ((v9 + v9) + "\n\t") {
    case "":
        fmt.Println(v9, v7)
    case "a":
        if ((v10 % ((v10) | 1)) <= (v10 * 482)) {
            if false {
                switch 127 {
                case 42:
                    v11 := GenS1{F0: 32767}
                    _ = v11
                case -1:
                    v10 += (^(v10))
                case 128:
                    if (65535 <= 2) {
                        if len(v3) > 0 {
                            v3[0] = 925
                        }
                        if true {
                            v7 = v7
                            if false {
                                fmt.Println(v9, v9, v7)
                                var v12 float64 = 0.1
                                _ = v12
                            }
                            v9 = (v9 + v9)
                        }
                        var v13 uint64 = 1
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v9 = v9
                            var v15 int8 = 1
                            _ = fi14
                            _ = v15
                        }
                        if true {
                            var v16 uint16 = 1
                            for fi17 := 0; fi17 < 15; fi17++ {
                                v2[392] = 1
                                _ = fi17
                            }
                            _ = v16
                        } else {
                            if (1 != 1) {
                                for fi18 := 0; fi18 < 18; fi18++ {
                                    for fi19 := 0; fi19 < 16; fi19++ {
                                        switch v9 {
                                        case "":
                                            v20 := GenS1{F0: 42}
                                            _ = v20
                                        case "foo bar":
                                            var v21 uint = v7
                                            _ = v21
                                        case "hello":
                                            if true {
                                                fmt.Println(v7, fi18, v9)
                                                var v22 int32 = 58
                                                if (4294967295 < 18446744073709551615) {
                                                    fmt.Println(v9, v22, v7)
                                                    var v23 float32 = -2.25
                                                    _ = v23
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = 0
                                                    }
                                                }
                                                _ = v22
                                            } else {
                                                var v24 float64 = 1.5
                                                if (-9223372036854775808 == -9223372036854775808) {
                                                    v10 -= (v10 & fi18)
                                                    if true {
                                                        if true {
                                                            v3 = append(v3, 1)
                                                            if false {
                                                                v10 -= v10
                                                                v25 := GenS2{F0: fi18, F1: -765}
                                                                _ = v25
                                                            } else {
                                                                var v26 int16 = -1
                                                                v10 = v10
                                                                var v27 int = fi19
                                                                _ = v26
                                                                _ = v27
                                                            }
                                                            _ = len(v2)
                                                            v7 += 4294967295
                                                            var v28 int16 = 372
                                                            _ = v28
                                                        } else {
                                                            var v29 bool = false
                                                            v30 := GenS3{F0: 2147483648, F1: 0, F2: 65535}
                                                            fmt.Println(v29, v24)
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                    } else {
                                                        var v31 int64 = 557
                                                        switch ("hello" + (v9 + (("foo bar" + ("\n\t" + v9)) + "\n\t"))) {
                                                        case "\n\t":
                                                            var v32 uint8 = 18
                                                            _ = v32
                                                        }
                                                        _ = v31
                                                    }
                                                } else {
                                                    v10 += fi19
                                                    var v33 uint64 = 4294967295
                                                    _ = v33
                                                }
                                                fmt.Println(v24, v10)
                                                switch "foo bar" {
                                                case "hello":
                                                    var v34 float32 = (1e10 + 0.1)
                                                    _ = v34
                                                }
                                                _ = v24
                                            }
                                        default:
                                            v10 = 356
                                        }
                                        _ = fi19
                                    }
                                    if (float64(((0.0 - ((float32((0.0 * 3.14159)) * 1e10) + ((1e10 + (1e10 - (3.14159 / 1.0))) + ((1e10 / 3.14159) + float32(float64(1e10)))))) + (3.14159 / 0.1))) < 1.5) {
                                        var v35 uint16 = 152
                                        v36 := GenS4{F0: (v7 % (((v7 * (v7 * 18446744073709551615))) | 1)), F1: 42, F2: -2147483648}
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v37 := GenS1{F0: -929}
                                        switch (v9 + (v9 + (v9 + (((("a" + ((v9 + ((v9 + ("a" + v9)) + ((v9 + ((("a" + v9) + (v9 + ("foo bar" + ""))) + v9)) + v9))) + v9)) + v9) + ("a" + v9)) + v9)))) {
                                        case "a":
                                            if true {
                                                v10 -= (fi18 / (((fi18 * -862)) | 1))
                                            } else {
                                                for fi38 := 0; fi38 < 13; fi38++ {
                                                    if (v10 != fi38) {
                                                        v2[-1] = 877
                                                    }
                                                    v39 := GenS5{F0: 2}
                                                    if (1e10 == 1e10) {
                                                        v40 := GenS2{F0: fi38, F1: 32767}
                                                        _ = v40
                                                    } else {
                                                        v7 = 1
                                                        var v41 int = (-(fi18))
                                                        _ = v41
                                                    }
                                                    v42 := GenS1{F0: 1}
                                                    delete(v1, 2147483648)
                                                    _ = fi38
                                                    _ = v39
                                                    _ = v42
                                                }
                                                v43 := v1[2]
                                                _ = v43
                                            }
                                        default:
                                            if (((!(false || false)) || (-235 == 0)) || false) {
                                                if (false || false) {
                                                    switch 887 {
                                                    case 32767:
                                                        if (fi18 != (fi18 + fi18)) {
                                                            switch fi18 {
                                                            case -429:
                                                                var v44 float32 = 1.5
                                                                _ = v44
                                                            default:
                                                                switch "a" {
                                                                case "\n\t":
                                                                    v3 = append(v3, 508)
                                                                case "foo bar":
                                                                    var v45 uint8 = 127
                                                                    _ = v45
                                                                case "hello":
                                                                    switch (v10 % ((2147483647) | 1)) {
                                                                    case 0:
                                                                        fmt.Println(v10, v10)
                                                                    case 4294967295:
                                                                        v46 := GenS3{F0: 889, F1: 1, F2: 42}
                                                                        _ = v46
                                                                    case 1:
                                                                        fmt.Println(v7, v9, v9)
                                                                    }
                                                                default:
                                                                    v10 += v10
                                                                }
                                                            }
                                                            v47 := GenS1{F0: 1000}
                                                            _ = v47
                                                        } else {
                                                            if (-128 == -1) {
                                                                var v48 uint16 = 65535
                                                                fmt.Println(fi18)
                                                                if false {
                                                                    var v49 uint64 = 42
                                                                    fmt.Println(v7, v49)
                                                                    if (0 <= (v7 % (((v7 << ((65535) & 63))) | 1))) {
                                                                        _ = len(v1)
                                                                        v9 = "\n\t"
                                                                    } else {
                                                                        v50 := GenS6{F0: 2147483647}
                                                                        fmt.Println(v10)
                                                                        v7 -= v7
                                                                        _ = v50
                                                                    }
                                                                    _ = v49
                                                                }
                                                                _ = v48
                                                            }
                                                            v9 = v9
                                                        }
                                                    case 9223372036854775807:
                                                        var v51 bool = true
                                                        _ = v51
                                                    case -427:
                                                        var v52 uint = 583
                                                        _ = v52
                                                    default:
                                                        for fi53 := 0; fi53 < 3; fi53++ {
                                                            if ((v10 << ((928) & 63)) != -9223372036854775808) {
                                                                var v54 int8 = 1
                                                                if false {
                                                                    v55 := GenS4{F0: v7, F1: 0, F2: 0}
                                                                    if (!false) {
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            delete(v1, 2)
                                                                        case "":
                                                                            var v56 int16 = -1
                                                                            _ = v56
                                                                        case "a":
                                                                            _ = len(v3)
                                                                            v7 = 880
                                                                            v10 += (fi53 | v10)
                                                                        }
                                                                        v10 = fi53
                                                                        v54 = -1
                                                                    }
                                                                    var v57 uint = v7
                                                                    _ = v55
                                                                    _ = v57
                                                                }
                                                                if (42 <= 0) {
                                                                    switch (^(fi53)) {
                                                                    case 4294967295:
                                                                        var v58 bool = false
                                                                        _ = v58
                                                                    case 128:
                                                                        var v59 int16 = 32767
                                                                        _ = v59
                                                                    case 0:
                                                                        v60 := GenS5{F0: 128}
                                                                        _ = v60
                                                                    default:
                                                                        v61 := GenS7{F0: 171, F1: 2, F2: 255}
                                                                        _ = v61
                                                                    }
                                                                    v10 += fi18
                                                                } else {
                                                                    fmt.Println(v10)
                                                                    fmt.Println(v10, v10, v54)
                                                                    v62 := GenS8{F0: "foo bar"}
                                                                    _ = v62
                                                                }
                                                                if (!true) {
                                                                    if false {
                                                                        if true {
                                                                            fmt.Println(v54, v7)
                                                                            v63 := GenS7{F0: 233, F1: 42, F2: 1}
                                                                            _ = v63
                                                                        }
                                                                        v10 += v10
                                                                    } else {
                                                                        v64 := GenS8{F0: ("\n\t" + v9)}
                                                                        var v65 bool = true
                                                                        if (292 >= 1) {
                                                                            var v66 float32 = 3.14159
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 1
                                                                            }
                                                                            if v65 {
                                                                                v7 = v7
                                                                                v54 = v54
                                                                            } else {
                                                                                v67 := GenS5{F0: 2}
                                                                                fmt.Println(v65, v10)
                                                                                var v68 int8 = (v54 - v54)
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                            _ = v66
                                                                        }
                                                                        _ = v64
                                                                        _ = v65
                                                                    }
                                                                } else {
                                                                    var v69 int32 = 8
                                                                    var v70 float64 = 0.1
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                _ = v54
                                                            } else {
                                                                var v71 uint64 = 598
                                                                var v72 int32 = 1
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                            _ = len(v2)
                                                            if (!false) {
                                                                v73 := v2[-2147483648]
                                                                var v74 int16 = 837
                                                                fmt.Println(fi18, v10)
                                                                v3 = append(v3, (v73 % (((v73 / ((v73) | 1))) | 1)))
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                            _ = fi53
                                                        }
                                                    }
                                                    for fi75 := 0; fi75 < 15; fi75++ {
                                                        if false {
                                                            var v76 int8 = 116
                                                            _ = v76
                                                        } else {
                                                            v9 = (v9 + v9)
                                                            switch (("hello" + "hello") + "a") {
                                                            case "a":
                                                                v7 = (v7 >> (((v7 - 1)) & 63))
                                                            case "":
                                                                if (((((fi75 | (v10 / ((68) | 1))) != fi18) || true) || false) || (1 <= 2147483647)) {
                                                                    fmt.Println(v10)
                                                                    v7 = v7
                                                                    delete(v1, 2147483647)
                                                                }
                                                            default:
                                                                v3 = append(v3, 2147483648)
                                                            }
                                                            v3 = append(v3, 0)
                                                        }
                                                        v77 := GenS9{F0: 562, F1: -2147483648}
                                                        _ = len(v3)
                                                        _ = fi75
                                                        _ = v77
                                                    }
                                                }
                                                var v78 int16 = 833
                                                if true {
                                                    if true {
                                                        v3 = append(v3, 182)
                                                        for fi79 := 0; fi79 < 12; fi79++ {
                                                            delete(v2, -2147483648)
                                                            v78 = 42
                                                            _ = fi79
                                                        }
                                                        var v80 uint32 = 1
                                                        _ = v80
                                                    } else {
                                                        var v81 uint8 = 91
                                                        for fi82 := 0; fi82 < 12; fi82++ {
                                                            if ((float64(-2.25) + 3.14159) >= 0.1) {
                                                                if (false || (!(0 < 710))) {
                                                                    var v83 float32 = (1.5 / -2.25)
                                                                    var v84 bool = true
                                                                    if v84 {
                                                                        fmt.Println(v7, v83)
                                                                        v85 := GenS6{F0: -2147483648}
                                                                        if (false || v84) {
                                                                            v81 = 131
                                                                            var v86 int16 = -581
                                                                            _ = v86
                                                                        } else {
                                                                            fmt.Println(v83, v9, v9)
                                                                        }
                                                                        _ = v85
                                                                    } else {
                                                                        v83 = v83
                                                                        var v87 int32 = -42
                                                                        _ = v87
                                                                    }
                                                                    switch v9 {
                                                                    case "\n\t":
                                                                        v7 = v7
                                                                    case "":
                                                                        var v88 int16 = v78
                                                                        _ = v88
                                                                    case "foo bar":
                                                                        var v89 int8 = 0
                                                                        _ = v89
                                                                    default:
                                                                        fmt.Println(v84, fi18, v83)
                                                                    }
                                                                    v90 := v1[2]
                                                                    _ = v83
                                                                    _ = v84
                                                                    _ = v90
                                                                }
                                                                switch (v9 + "foo bar") {
                                                                case "\n\t":
                                                                    if (int8(fi18) <= 1) {
                                                                        fmt.Println(fi82)
                                                                        if (2147483647 >= 1) {
                                                                            var v91 int64 = int64(v78)
                                                                            v2[(fi18 + int(v78))] = 924
                                                                            _ = v91
                                                                        }
                                                                    } else {
                                                                        switch (v9 + ("hello" + ("\n\t" + v9))) {
                                                                        case "hello":
                                                                            var v92 uint8 = (v81 >> ((v7) & 63))
                                                                            _ = v92
                                                                        case "a":
                                                                            var v93 uint64 = 1
                                                                            _ = v93
                                                                        case "":
                                                                            v94 := v1[985]
                                                                            _ = v94
                                                                        }
                                                                        v95 := GenS10{F0: 960, F1: v81, F2: 1}
                                                                        fmt.Println(v81, fi82)
                                                                        _ = v95
                                                                    }
                                                                case "":
                                                                    switch v9 {
                                                                    case "\n\t":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    case "":
                                                                        if (!true) {
                                                                            var v96 uint = v7
                                                                            switch fi18 {
                                                                            case -989:
                                                                                switch (^(fi82)) {
                                                                                case 127:
                                                                                    if true {
                                                                                        v97 := GenS11{F0: fi18}
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        switch (^(fi82)) {
                                                                                        case 127:
                                                                                            delete(v2, (v10 - v10))
                                                                                        }
                                                                                        _ = v97
                                                                                    } else {
                                                                                        v81 = v81
                                                                                        _ = len(v3)
                                                                                        v98 := GenS12{F0: v81, F1: -1}
                                                                                        _ = v98
                                                                                    }
                                                                                default:
                                                                                    v99 := v2[v10]
                                                                                    _ = v99
                                                                                }
                                                                            case 127:
                                                                                v81 = v81
                                                                            }
                                                                            v100 := v2[fi82]
                                                                            _ = v100
                                                                            _ = v96
                                                                        } else {
                                                                            if false {
                                                                                if (((v9 + (v9 + "\n\t")) >= "") && (!(false || (0 >= -2147483648)))) {
                                                                                    var v101 int = v10
                                                                                    fmt.Println(v7, v81, v7)
                                                                                    v101 -= v101
                                                                                    v102 := GenS3{F0: 32, F1: v78, F2: 255}
                                                                                    var v103 int64 = -292
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                } else {
                                                                                    var v104 uint8 = v81
                                                                                    _ = v104
                                                                                }
                                                                                switch "a" {
                                                                                case "a":
                                                                                    v81 = (v81 % ((v81) | 1))
                                                                                    v105 := GenS12{F0: 125, F1: -70}
                                                                                    v78 += 42
                                                                                    _ = v105
                                                                                case "\n\t":
                                                                                    v106 := GenS9{F0: (^(v78)), F1: -787}
                                                                                    _ = v106
                                                                                case "hello":
                                                                                    var v107 int16 = 324
                                                                                    _ = v107
                                                                                }
                                                                                fmt.Println(v81)
                                                                            } else {
                                                                                var v108 uint = (v7 % (((v7 / ((v7) | 1))) | 1))
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    switch (^(fi18)) {
                                                                                    case 128:
                                                                                        v7 += (^(v7))
                                                                                    case 2147483647:
                                                                                        if true {
                                                                                            var v109 uint64 = 4294967295
                                                                                            if (-1 < 572) {
                                                                                                var v110 uint32 = 424
                                                                                                _ = v110
                                                                                            }
                                                                                            var v111 int32 = 2147483647
                                                                                            _ = v109
                                                                                            _ = v111
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    v112 := GenS13{F0: v81}
                                                                                    _ = v112
                                                                                }
                                                                                _ = v108
                                                                            }
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    v9 = "hello"
                                                                default:
                                                                    if true {
                                                                        if (!((1.5 + 0.0) != -2.25)) {
                                                                            var v113 float32 = 0.1
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            _ = v113
                                                                        } else {
                                                                            fmt.Println(v81, v78, v10)
                                                                            v7 = (^(v7))
                                                                        }
                                                                        var v114 uint8 = v81
                                                                        if (((^(v81)) < (v114 * (v81 << ((v7) & 63)))) && (182 != 65535)) {
                                                                            switch v9 {
                                                                            case "\n\t":
                                                                                v115 := v2[(v10 | 32768)]
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                        _ = v114
                                                                    } else {
                                                                        v7 += (v7 * (v7 + 0))
                                                                    }
                                                                }
                                                            } else {
                                                                v78 += -1
                                                                v116 := GenS6{F0: 2147483647}
                                                                _ = v116
                                                            }
                                                            _ = fi82
                                                        }
                                                        _ = v81
                                                    }
                                                    v7 += v7
                                                    var v117 float32 = 1.5
                                                    _ = v117
                                                } else {
                                                    v7 += (^(v7))
                                                    for fi118 := 0; fi118 < 4; fi118++ {
                                                        if (32767 == 755) {
                                                            v10 = v10
                                                            _ = len(v3)
                                                        }
                                                        _ = fi118
                                                    }
                                                    var v119 uint16 = 42
                                                    _ = v119
                                                }
                                                _ = v78
                                            } else {
                                                if ((0.0 + 0.1) <= 0.1) {
                                                    var v120 uint = v7
                                                    v120 -= (v7 * v7)
                                                    if ((0.1 / (((float32(1e10) + 0.0) * 0.0) + float32((0.1 - float64(-2.25))))) < 1e10) {
                                                        v121 := GenS4{F0: 316, F1: 32767, F2: -487}
                                                        _ = v121
                                                    } else {
                                                        v122 := GenS14{F0: -1}
                                                        v9 = v9
                                                        _ = v122
                                                    }
                                                    _ = v120
                                                }
                                                switch (fi18 >> ((v7) & 63)) {
                                                case 42:
                                                    switch (v9 + v9) {
                                                    case "a":
                                                        var v123 int = v10
                                                        _ = v123
                                                    case "\n\t":
                                                        if (!false) {
                                                            fmt.Println(v10, v10, v10)
                                                            fmt.Println(v7, v9)
                                                            var v124 string = v9
                                                            _ = v124
                                                        } else {
                                                            for fi125 := 0; fi125 < 20; fi125++ {
                                                                v9 = ""
                                                                if true {
                                                                    var v126 uint32 = 236
                                                                    _ = v126
                                                                }
                                                                if true {
                                                                    v10 = fi125
                                                                } else {
                                                                    if true {
                                                                        var v127 uint16 = 42
                                                                        var v128 float32 = 1e10
                                                                        fmt.Println(fi125)
                                                                        _ = v127
                                                                        _ = v128
                                                                    } else {
                                                                        v1[65535] = ((true && true) || (880 > 42))
                                                                        if ((false || true) && false) {
                                                                            fmt.Println(fi125, fi125, fi18)
                                                                            var v129 string = ""
                                                                            v129 = ""
                                                                            _ = v129
                                                                        } else {
                                                                            if (((1e10 > ((1e10 - 3.14159) * 1.5)) || false) && (0 > 95)) {
                                                                                switch (v9 + (v9 + "\n\t")) {
                                                                                case "a":
                                                                                    v9 = v9
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        v1[uint32(v7)] = ((0.1 + (float64(float32((0.0 * (0.1 / (-2.25 + 0.0))))) - (3.14159 / 1.0))) > (-2.25 - 1e10))
                                                                                        var v130 string = v9
                                                                                        _ = v130
                                                                                    } else {
                                                                                        if (!true) {
                                                                                            var v131 uint32 = 2147483648
                                                                                            _ = v131
                                                                                        }
                                                                                        v132 := GenS15{F0: 2}
                                                                                        if (-2147483648 == v10) {
                                                                                            v7 -= v7
                                                                                            v133 := GenS4{F0: (v7 - v7), F1: 32767, F2: 1}
                                                                                            _ = v133
                                                                                        }
                                                                                        _ = v132
                                                                                    }
                                                                                default:
                                                                                    v10 = int(v7)
                                                                                    v134 := GenS16{F0: 0}
                                                                                    _ = v134
                                                                                }
                                                                                _ = len(v2)
                                                                            }
                                                                            v135 := GenS17{F0: 255, F1: 180, F2: (v9 + ("foo bar" + "a"))}
                                                                            _ = v135
                                                                        }
                                                                        v136 := GenS12{F0: 0, F1: 1}
                                                                        _ = v136
                                                                    }
                                                                    var v137 float64 = (0.1 - 1.5)
                                                                    var v138 uint = v7
                                                                    var v139 float64 = v137
                                                                    var v140 uint32 = 1
                                                                    v137 += v137
                                                                    var v141 int = fi18
                                                                    v142 := GenS18{F0: 127, F1: (42 > -552), F2: 32767}
                                                                    _ = v137
                                                                    _ = v138
                                                                    _ = v139
                                                                    _ = v140
                                                                    _ = v141
                                                                    _ = v142
                                                                }
                                                                if false {
                                                                    if (!true) {
                                                                        fmt.Println(fi125, v10, v9)
                                                                    }
                                                                    switch (fi125 & fi18) {
                                                                    case -1:
                                                                        v9 = "foo bar"
                                                                    }
                                                                    v10 -= -9223372036854775808
                                                                    switch ((v9 + ("foo bar" + v9)) + (((((v9 + "foo bar") + "") + (v9 + v9)) + ((v9 + v9) + ("\n\t" + "a"))) + v9)) {
                                                                    case "":
                                                                        switch fi18 {
                                                                        case 32767:
                                                                            v10 += (fi18 + fi18)
                                                                        case -1:
                                                                            v3 = append(v3, 870)
                                                                        default:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 467
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        switch ("hello" + ((v9 + v9) + "\n\t")) {
                                                                        case "foo bar":
                                                                            fmt.Println(fi18, v7)
                                                                        case "hello":
                                                                            var v143 uint32 = 2147483648
                                                                            _ = v143
                                                                        }
                                                                    case "foo bar":
                                                                        if (42 > 42) {
                                                                            var v144 float64 = 1.5
                                                                            v7 += v7
                                                                            _ = v144
                                                                        } else {
                                                                            v2[fi18] = 42
                                                                        }
                                                                    default:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 327
                                                                        }
                                                                    }
                                                                }
                                                                var v145 uint8 = 127
                                                                switch (fi125 * -46) {
                                                                case 128:
                                                                    v146 := GenS19{F0: -2147483648, F1: 5, F2: -32768}
                                                                    _ = v146
                                                                }
                                                                _ = fi125
                                                                _ = v145
                                                            }
                                                            var v147 float64 = 3.14159
                                                            var v148 int32 = 42
                                                            _ = v147
                                                            _ = v148
                                                        }
                                                    case "hello":
                                                        if (v7 > (v7 / (((v7 + v7)) | 1))) {
                                                            v10 = fi18
                                                            if true {
                                                                switch v10 {
                                                                case 629:
                                                                    v7 = v7
                                                                case -226:
                                                                    for fi149 := 0; fi149 < 10; fi149++ {
                                                                        var v150 uint64 = 18446744073709551615
                                                                        var v151 int16 = 0
                                                                        _ = fi149
                                                                        _ = v150
                                                                        _ = v151
                                                                    }
                                                                case 481:
                                                                    fmt.Println(fi18)
                                                                }
                                                                fmt.Println(v9)
                                                            }
                                                        } else {
                                                            switch v9 {
                                                            case "hello":
                                                                v10 += (v10 | (v10 % ((v10) | 1)))
                                                            default:
                                                                var v152 float64 = (float64(((0.1 + -2.25) - float32(3.14159))) * 3.14159)
                                                                _ = v152
                                                            }
                                                            var v153 uint8 = 13
                                                            _ = v153
                                                        }
                                                    }
                                                }
                                                var v154 int64 = -495
                                                _ = v154
                                            }
                                        }
                                        fmt.Println(v7)
                                        _ = v37
                                    }
                                    _ = fi18
                                }
                                if false {
                                    if (true || (false && (0 != 18446744073709551615))) {
                                        fmt.Println(v10, v10, v7)
                                        v155 := GenS18{F0: 1, F1: (false && false), F2: 2}
                                        _ = v155
                                    } else {
                                        if (!true) {
                                            fmt.Println(v9, v7)
                                            fmt.Println(v7, v10)
                                            var v156 int = v10
                                            _ = v156
                                        } else {
                                            v157 := GenS20{F0: 0, F1: 42}
                                            if ((false || true) || (487 == -1)) {
                                                var v158 string = v9
                                                var v159 float32 = (((1e10 / (float32(0.1) - 0.0)) - float32((1e10 - (3.14159 - float64((1.5 * -2.25)))))) + 1.5)
                                                _ = v158
                                                _ = v159
                                            }
                                            _ = v157
                                        }
                                    }
                                    for fi160 := 0; fi160 < 15; fi160++ {
                                        v10 = fi160
                                        switch v10 {
                                        case 1:
                                            v10 = 2147483647
                                        case -2147483648:
                                            v9 = v9
                                        case 215:
                                            v7 = (^(v7))
                                        default:
                                            var v161 uint64 = 18446744073709551615
                                            _ = v161
                                        }
                                        _ = fi160
                                    }
                                    if true {
                                        v10 = v10
                                        if (true && false) {
                                            fmt.Println(v10, v10)
                                        }
                                        fmt.Println(v9, v10)
                                    }
                                }
                                fmt.Println(v7, v9, v9)
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            for fi162 := 0; fi162 < 11; fi162++ {
                                for fi163 := 0; fi163 < 16; fi163++ {
                                    var v164 int8 = 121
                                    v9 = "a"
                                    _ = fi163
                                    _ = v164
                                }
                                _ = fi162
                            }
                        }
                        v165 := GenS6{F0: -1}
                        _ = v165
                    }
                }
            }
            if false {
                var v166 uint32 = 1
                var v167 uint8 = uint8(v7)
                for fi168 := 0; fi168 < 14; fi168++ {
                    v169 := GenS21{F0: "hello", F1: 9223372036854775807, F2: -1}
                    _ = fi168
                    _ = v169
                }
                _ = v166
                _ = v167
            }
            var v170 uint8 = 42
            _ = v170
        } else {
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
        }
    }
    if true {
        for fi171 := 0; fi171 < 5; fi171++ {
            for fi172 := 0; fi172 < 19; fi172++ {
                fmt.Println(v7)
                fmt.Println(fi172)
                _ = fi172
            }
            _ = fi171
        }
        v7 = v7
    }
    v9 = ((v9 + "\n\t") + (((v9 + v9) + ("" + "foo bar")) + "\n\t"))
    switch ("hello" + (v9 + "")) {
    case "foo bar":
        var v173 uint32 = 4294967295
        _ = v173
    case "\n\t":
        v9 = v9
    case "a":
        fmt.Println(v9)
    }
    fmt.Println(v10, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
