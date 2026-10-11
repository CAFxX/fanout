package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int8; F1 int; F2 int16 }
type GenS7 struct { F0 int64; F1 uint; F2 int64 }
type GenS8 struct { F0 int8; F1 int8; F2 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 int32 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 bool; F1 int; F2 uint8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint16 }

func main() {
    var v1 float64 = 0.1
    var v2 string = "\n\t"
    v3 := make(map[uint32]int)
    var v4 bool = (511 == 2)
    switch 192 {
    case 127:
        v5 := []int{}
        _ = v5
    case 4294967295:
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(v4, fi6)
            _ = fi6
        }
    default:
        v2 = "a"
    }
    fmt.Println(v2, v2, v4)
    if v4 {
        for fi7 := 0; fi7 < 8; fi7++ {
            switch "a" {
            case "a":
                if true {
                    v8 := []int{797, (fi7 & 9223372036854775807), (fi7 % ((2147483647) | 1)), fi7}
                    var v9 int64 = 9223372036854775807
                    for fi10 := 0; fi10 < 12; fi10++ {
                        fmt.Println(v4)
                        _ = fi10
                    }
                    _ = v8
                    _ = v9
                }
            case "\n\t":
                var v11 int16 = 0
                _ = v11
            case "foo bar":
                v2 = "hello"
            default:
                if (v1 >= 1.5) {
                    if (767 == -665) {
                        for fi12 := 0; fi12 < 8; fi12++ {
                            if (v4 || (((v2 == v2) && ((v4 && (!((float32(v1) >= -2.25) || v4))) || (624 >= 42))) && v4)) {
                                var v13 int16 = -32768
                                v13 -= v13
                                var v14 float32 = (float32(float64((1.5 - 1e10))) + float32(v1))
                                _ = v13
                                _ = v14
                            } else {
                                var v15 float32 = 1e10
                                if false {
                                    var v16 int32 = 2147483647
                                    fmt.Println(v2, v15, v16)
                                    var v17 int16 = 42
                                    if v4 {
                                        switch 454 {
                                        case -9223372036854775808:
                                            switch ("a" + ("" + v2)) {
                                            case "":
                                                fmt.Println(v2)
                                            case "foo bar":
                                                v15 = v15
                                            case "\n\t":
                                                fmt.Println(fi7, v1, v1)
                                                var v18 int32 = 0
                                                v1 = ((v1 * v1) / 1.0)
                                                _ = v18
                                            }
                                            v15 = v15
                                            if (v4 || (!v4)) {
                                                v1 += 1.5
                                                fmt.Println(v17)
                                                if (42 > (v17 & (v17 - v17))) {
                                                    v17 = (v17 | (v17 - (v17 & v17)))
                                                }
                                            }
                                        case 9223372036854775807:
                                            var v19 int = (^(fi7))
                                            _ = v19
                                        case -2147483648:
                                            var v20 int16 = 0
                                            _ = v20
                                        default:
                                            fmt.Println(v2, fi7)
                                        }
                                    }
                                    if (!v4) {
                                        if v4 {
                                            fmt.Println(v17)
                                            fmt.Println(v4)
                                            v21 := []int{(fi7 * (fi12 - fi7)), (fi12 >> ((484) & 63))}
                                            v22 := GenS1{F0: 65535, F1: 731}
                                            _ = v21
                                            _ = v22
                                        }
                                        var v23 int8 = 1
                                        v16 = (v16 % (((v16 + v16)) | 1))
                                        var v24 uint8 = 2
                                        v24 -= 1
                                        _ = v23
                                        _ = v24
                                    } else {
                                        if (127 > 64) {
                                            if false {
                                                fmt.Println(v16)
                                                fmt.Println(fi12, v17, v4)
                                                fmt.Println(fi12)
                                                v2 = v2
                                                var v25 string = v2
                                                _ = v25
                                            }
                                            var v26 int32 = v16
                                            _ = v26
                                        } else {
                                            var v27 uint8 = 42
                                            v2 = "\n\t"
                                            v4 = true
                                            _ = v27
                                        }
                                        if v4 {
                                            switch (fi12 << ((363) & 63)) {
                                            case 175:
                                                v28 := GenS2{F0: 1}
                                                _ = v28
                                            case -783:
                                                var v29 int32 = (v16 + 27)
                                                v4 = v4
                                                _ = v29
                                            }
                                            fmt.Println(v2, fi12, v4)
                                        } else {
                                            var v30 uint32 = 659
                                            var v31 uint8 = 127
                                            if v4 {
                                                if v4 {
                                                    v32 := GenS2{F0: fi7}
                                                    _ = v32
                                                }
                                                var v33 float32 = float32(3.14159)
                                                if (!(-607 < (v16 / (((v16 | (v16 % (((-(v16))) | 1)))) | 1)))) {
                                                    fmt.Println(fi12, v16)
                                                    v17 = -499
                                                } else {
                                                    v17 = (v17 * (v17 ^ v17))
                                                    v15 = v33
                                                    if v4 {
                                                        var v34 int16 = (v17 % ((v17) | 1))
                                                        switch fi7 {
                                                        case -9223372036854775808:
                                                            switch fi7 {
                                                            case 125:
                                                                var v35 int8 = 100
                                                                _ = v35
                                                            default:
                                                                v36 := []int{(fi7 << ((18446744073709551615) & 63)), fi7, (fi7 - (fi7 % (((-(fi7))) | 1))), (-(fi12))}
                                                                _ = v36
                                                            }
                                                            v16 -= (v16 - 42)
                                                        case -295:
                                                            var v37 uint32 = v30
                                                            _ = v37
                                                        }
                                                        _ = v34
                                                    }
                                                    v38 := []int{-9223372036854775808, fi12}
                                                    _ = v38
                                                }
                                                _ = v33
                                            } else {
                                                var v39 int32 = -33
                                                _ = v39
                                            }
                                            _ = v30
                                            _ = v31
                                        }
                                        var v40 bool = (v15 > 3.14159)
                                        _ = v40
                                    }
                                    _ = v16
                                    _ = v17
                                }
                                v4 = v4
                                _ = v15
                            }
                            _ = fi12
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if (1e10 != 1e10) {
                                v41 := GenS2{F0: fi7}
                                var v42 string = (v2 + ("\n\t" + ""))
                                if false {
                                    fmt.Println(v2, v4, v4)
                                    var v43 int = fi7
                                    v44 := GenS1{F0: 65535, F1: 4294967295}
                                    _ = v43
                                    _ = v44
                                } else {
                                    _ = len(v3)
                                    for fi45 := 0; fi45 < 16; fi45++ {
                                        v1 = (v1 - v1)
                                        delete(v3, 2)
                                        fmt.Println(v42, v2)
                                        _ = fi45
                                    }
                                }
                                _ = v41
                                _ = v42
                            }
                        case "a":
                            v46 := GenS1{F0: 718, F1: 1}
                            _ = v46
                        default:
                            var v47 bool = ((v4 && false) && (!v4))
                            _ = v47
                        }
                    } else {
                        v2 = (v2 + (v2 + v2))
                    }
                    v48 := GenS1{F0: 300, F1: 18446744073709551615}
                    _ = v48
                }
            }
            var v49 uint8 = 1
            v50 := []int{0}
            _ = fi7
            _ = v49
            _ = v50
        }
    } else {
        var v51 string = v2
        fmt.Println(v1, v2, v2)
        _ = v51
    }
    v4 = ((1 != 0) || v4)
    var v52 uint32 = 2
    v53 := GenS3{F0: 65535}
    for fi54 := 0; fi54 < 7; fi54++ {
        v55 := GenS3{F0: 304}
        _ = fi54
        _ = v55
    }
    v56 := []int{1, 2147483647, 1, -501}
    delete(v3, v52)
    v4 = (57 < 18446744073709551615)
    if ((!v4) || v4) {
        switch 1 {
        case 32767:
            v57 := GenS2{F0: 247}
            _ = v57
        case 9223372036854775807:
            delete(v3, v52)
        case 1:
            v2 = ""
        default:
            v56 = append(v56, 128)
        }
        for fi58 := 0; fi58 < 19; fi58++ {
            for fi59 := 0; fi59 < 6; fi59++ {
                var v60 int64 = 42
                if (1 >= 210) {
                    switch (fi59 << ((207) & 63)) {
                    case 127:
                        if (!v4) {
                            var v61 bool = (-1 >= 1)
                            var v62 uint64 = 2
                            v63 := v3[v52]
                            _ = v61
                            _ = v62
                            _ = v63
                        } else {
                            if (-76 >= 6) {
                                _ = len(v3)
                            } else {
                                if len(v56) > 0 {
                                    v56[0] = (fi59 | fi59)
                                }
                            }
                            v1 = (((v1 - v1) + v1) * (-2.25 / v1))
                        }
                    case -76:
                        if (!false) {
                            var v64 uint = 65535
                            _ = v64
                        } else {
                            var v65 uint64 = 85
                            v56 = append(v56, (^(fi58)))
                            _ = v65
                        }
                    case -329:
                        var v66 string = v2
                        _ = v66
                    }
                    v67 := GenS1{F0: 1, F1: 4294967295}
                    v1 = 0.0
                    _ = v67
                }
                var v68 int = (fi58 << ((1) & 63))
                fmt.Println(v60, v2, v68)
                _ = fi59
                _ = v60
                _ = v68
            }
            var v69 float64 = float64(0.0)
            if v4 {
                switch "\n\t" {
                case "\n\t":
                    v70 := GenS4{F0: (v52 <= (^(v52))), F1: 1}
                    _ = v70
                }
                switch v2 {
                case "foo bar":
                    v71 := v3[(v52 / ((352) | 1))]
                    _ = v71
                default:
                    v4 = ((1.5 < 3.14159) || v4)
                }
                var v72 uint = 2
                _ = v72
            } else {
                v73 := GenS2{F0: (fi58 | (^(fi58)))}
                switch fi58 {
                case -600:
                    var v74 int32 = -1
                    _ = v74
                default:
                    for fi75 := 0; fi75 < 18; fi75++ {
                        if (-9223372036854775808 > -1) {
                            var v76 float32 = -2.25
                            _ = v76
                        } else {
                            v2 = (v2 + v2)
                            switch v2 {
                            case "\n\t":
                                fmt.Println(v69, fi75)
                            case "":
                                switch (fi58 | fi58) {
                                case 9223372036854775807:
                                    fmt.Println(v69, v2)
                                case 128:
                                    var v77 bool = (!(!(2 >= 128)))
                                    _ = v77
                                }
                            case "hello":
                                switch v2 {
                                case "\n\t":
                                    if v4 {
                                        var v78 int8 = 0
                                        if v4 {
                                            delete(v3, (v52 & v52))
                                        } else {
                                            if false {
                                                v79 := GenS3{F0: 255}
                                                if v4 {
                                                    switch -2147483648 {
                                                    case 378:
                                                        v56 = append(v56, fi75)
                                                    default:
                                                        var v80 float32 = float32(v1)
                                                        _ = v80
                                                    }
                                                    v56 = append(v56, 175)
                                                }
                                                _ = v79
                                            }
                                            fmt.Println(fi58)
                                            v2 = "foo bar"
                                            switch v2 {
                                            case "foo bar":
                                                if len(v56) > 0 {
                                                    v56[0] = 32768
                                                }
                                            case "hello":
                                                var v81 uint64 = 734
                                                _ = v81
                                            case "a":
                                                var v82 uint = 42
                                                _ = v82
                                            }
                                            fmt.Println(v2, v2)
                                        }
                                        _ = v78
                                    }
                                case "":
                                    var v83 int = (fi75 | fi75)
                                    _ = v83
                                case "a":
                                    switch (v2 + v2) {
                                    case "\n\t":
                                        v1 = v69
                                    case "hello":
                                        fmt.Println(v1)
                                    case "foo bar":
                                        var v84 int16 = -32768
                                        _ = v84
                                    default:
                                        if false {
                                            if v4 {
                                                if v4 {
                                                    v69 += v1
                                                } else {
                                                    var v85 int = (fi75 + fi75)
                                                    var v86 float32 = float32((0.1 + (float64(0.1) + float64(((0.0 - 0.1) * 0.1)))))
                                                    _ = v85
                                                    _ = v86
                                                }
                                                v4 = (0.0 > float64((((((-2.25 / 1.0) + (0.1 + 0.0)) - 1e10) / -2.25) + 1.5)))
                                                if v4 {
                                                    var v87 int32 = -2147483648
                                                    _ = v87
                                                }
                                            }
                                            switch v2 {
                                            case "foo bar":
                                                if (1.5 >= v69) {
                                                    switch (fi75 + (fi58 | fi75)) {
                                                    case -779:
                                                        if (!v4) {
                                                            var v88 uint8 = 255
                                                            var v89 uint16 = 32767
                                                            v69 -= 0.1
                                                            _ = v88
                                                            _ = v89
                                                        } else {
                                                            if v4 {
                                                                if (0 != -690) {
                                                                    if v4 {
                                                                        v2 = (v2 + "hello")
                                                                        v1 = float64(0.0)
                                                                    } else {
                                                                        v1 -= ((v1 / (v69 + v69)) + 3.14159)
                                                                        var v90 float32 = ((float32(v69) * -2.25) + 1e10)
                                                                        fmt.Println(fi58)
                                                                        _ = v90
                                                                    }
                                                                    var v91 uint = 1
                                                                    var v92 int8 = 1
                                                                    _ = v91
                                                                    _ = v92
                                                                } else {
                                                                    var v93 float32 = 3.14159
                                                                    var v94 int = (fi75 / (((-(fi75))) | 1))
                                                                    _ = v93
                                                                    _ = v94
                                                                }
                                                                fmt.Println(v2, v2, fi58)
                                                            } else {
                                                                v2 = (((v2 + v2) + v2) + (((v2 + "") + "foo bar") + v2))
                                                                v52 = v52
                                                                v4 = false
                                                            }
                                                        }
                                                    case 43:
                                                        var v95 int64 = 209
                                                        _ = v95
                                                    }
                                                    if len(v56) >= 2 {
                                                        v56 = v56[1:len(v56)]
                                                    }
                                                    var v96 bool = v4
                                                    _ = v96
                                                } else {
                                                    var v97 bool = v4
                                                    if true {
                                                        var v98 uint16 = 0
                                                        v69 = 1.5
                                                        _ = v98
                                                    } else {
                                                        v4 = (127 < 3)
                                                        v1 = v1
                                                    }
                                                    v3[(v52 & (^(v52)))] = 127
                                                    var v99 int32 = 61
                                                    v1 = v1
                                                    switch ("a" + v2) {
                                                    case "foo bar":
                                                        if (2147483648 <= v52) {
                                                            v4 = (0 == 4)
                                                            if true {
                                                                v52 -= 0
                                                                if (!(!v97)) {
                                                                    switch fi75 {
                                                                    case 1:
                                                                        var v100 uint32 = v52
                                                                        _ = v100
                                                                    case -9223372036854775808:
                                                                        var v101 uint64 = 4294967295
                                                                        _ = v101
                                                                    case 32767:
                                                                        if (18446744073709551615 < 0) {
                                                                            v2 = "foo bar"
                                                                            if v97 {
                                                                                fmt.Println(v69)
                                                                                fmt.Println(v99, fi58)
                                                                                switch 128 {
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(fi58, v1)
                                                                                case 128:
                                                                                    var v102 uint64 = 42
                                                                                    _ = v102
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v103 int16 = -308
                                                                            v103 -= v103
                                                                            _ = v103
                                                                        }
                                                                        v104 := v3[v52]
                                                                        _ = v104
                                                                    default:
                                                                        if true {
                                                                            var v105 string = ""
                                                                            _ = v105
                                                                        }
                                                                    }
                                                                    fmt.Println(v4, v1, v97)
                                                                } else {
                                                                    v106 := GenS3{F0: 18446744073709551615}
                                                                    v56 = append(v56, fi75)
                                                                    v2 = "hello"
                                                                    _ = v106
                                                                }
                                                            } else {
                                                                v107 := GenS1{F0: 138, F1: 18446744073709551615}
                                                                _ = v107
                                                            }
                                                            v108 := GenS3{F0: 65535}
                                                            _ = v108
                                                        } else {
                                                            v56 = append(v56, (fi58 / (((fi75 >> ((249) & 63))) | 1)))
                                                            if (0.0 != 3.14159) {
                                                                var v109 bool = v4
                                                                v110 := GenS4{F0: v109, F1: v99}
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                        }
                                                    }
                                                    v52 += (v52 >> ((42) & 63))
                                                    if (42 < 42) {
                                                        v69 -= v69
                                                    } else {
                                                        v111 := GenS5{F0: 617, F1: v2}
                                                        _ = v111
                                                    }
                                                    _ = v97
                                                    _ = v99
                                                }
                                            }
                                        } else {
                                            v112 := GenS5{F0: 167, F1: "foo bar"}
                                            _ = v112
                                        }
                                    }
                                default:
                                    if (v4 || v4) {
                                        v113 := GenS6{F0: 42, F1: fi58, F2: 1}
                                        var v114 float64 = float64(1.5)
                                        if false {
                                            fmt.Println(v1, v52, v2)
                                        }
                                        var v115 uint8 = 250
                                        _ = v113
                                        _ = v114
                                        _ = v115
                                    } else {
                                        if v4 {
                                            v4 = ((95 != 1) || v4)
                                            var v116 int32 = 76
                                            if len(v56) >= 2 {
                                                v56 = v56[1:len(v56)]
                                            }
                                            _ = v116
                                        } else {
                                            fmt.Println(fi58, v69)
                                            v4 = (0 < 1)
                                            v117 := GenS2{F0: -993}
                                            _ = v117
                                        }
                                    }
                                    _ = len(v3)
                                    if v4 {
                                        v118 := GenS7{F0: 0, F1: 4294967295, F2: 318}
                                        fmt.Println(v52, fi75, fi58)
                                        if v4 {
                                            v4 = true
                                            fmt.Println(v52)
                                            var v119 bool = ((!true) && v4)
                                            _ = v119
                                        } else {
                                            if true {
                                                if (-9223372036854775808 < fi58) {
                                                    v4 = v4
                                                    v69 -= v1
                                                    if len(v56) > 0 {
                                                        v56[0] = 1
                                                    }
                                                    fmt.Println(v52, v4)
                                                    v1 = 0.0
                                                } else {
                                                    switch (fi58 / (((fi58 * fi58)) | 1)) {
                                                    case 32767:
                                                        v52 = (^(v52))
                                                    case 880:
                                                        v120 := GenS8{F0: 86, F1: 100, F2: -128}
                                                        _ = v120
                                                    case 4294967295:
                                                        v4 = (!(!false))
                                                    default:
                                                        switch -586 {
                                                        case -9223372036854775808:
                                                            v4 = v4
                                                        case 127:
                                                            v121 := GenS3{F0: 325}
                                                            _ = v121
                                                        }
                                                    }
                                                }
                                                var v122 float32 = -2.25
                                                _ = v122
                                            } else {
                                                var v123 uint = 255
                                                fmt.Println(v1, v2, v52)
                                                v124 := GenS2{F0: fi75}
                                                _ = v123
                                                _ = v124
                                            }
                                            var v125 int8 = 42
                                            _ = v125
                                        }
                                        _ = v118
                                    } else {
                                        if len(v56) >= 2 {
                                            v56 = v56[1:len(v56)]
                                        }
                                    }
                                    var v126 string = "foo bar"
                                    fmt.Println(v126, v69)
                                    var v127 int64 = 88
                                    if ((-128 == 0) || v4) {
                                        fmt.Println(v2, fi75, v4)
                                    }
                                    var v128 uint8 = 0
                                    _ = v126
                                    _ = v127
                                    _ = v128
                                }
                            }
                            if true {
                                v129 := GenS1{F0: 32767, F1: 18446744073709551615}
                                v69 = v1
                                v130 := GenS9{F0: -128}
                                _ = v129
                                _ = v130
                            }
                        }
                        fmt.Println(v69)
                        _ = fi75
                    }
                }
                _ = v73
            }
            _ = fi58
            _ = v69
        }
    } else {
        switch v2 {
        case "hello":
            v131 := GenS5{F0: 65535, F1: (v2 + (v2 + "foo bar"))}
            _ = v131
        case "":
            var v132 int32 = 2147483647
            _ = v132
        case "foo bar":
            if len(v56) >= 2 {
                v56 = v56[1:len(v56)]
            }
        }
    }
    if (2 > 2147483647) {
        if v4 {
            switch 2147483647 {
            case 32767:
                var v133 float32 = 3.14159
                _ = v133
            case 128:
                delete(v3, 2147483648)
            case 429:
                for fi134 := 0; fi134 < 8; fi134++ {
                    if ("foo bar" != ("hello" + v2)) {
                        if (42 == -60) {
                            v135 := GenS3{F0: 0}
                            _ = v135
                        }
                        v136 := GenS10{F0: 895}
                        _ = v136
                    }
                    if v4 {
                        _ = len(v56)
                        fmt.Println(v1, v52)
                    } else {
                        v137 := GenS7{F0: 976, F1: 749, F2: 507}
                        _ = v137
                    }
                    _ = fi134
                }
            default:
                if true {
                    if len(v56) > 0 {
                        v56[0] = -1
                    }
                    fmt.Println(v1)
                    v1 = v1
                } else {
                    v138 := GenS11{F0: v2, F1: -1}
                    var v139 int16 = 601
                    if v4 {
                        if v4 {
                            v140 := GenS1{F0: 2, F1: 633}
                            _ = v140
                        }
                        for fi141 := 0; fi141 < 17; fi141++ {
                            var v142 int64 = 9223372036854775807
                            v143 := GenS9{F0: -128}
                            fmt.Println(v4, v2)
                            _ = fi141
                            _ = v142
                            _ = v143
                        }
                        var v144 uint = 65535
                        _ = v144
                    } else {
                        v145 := GenS10{F0: 0}
                        v146 := GenS12{F0: 42}
                        _ = v145
                        _ = v146
                    }
                    _ = v138
                    _ = v139
                }
            }
        } else {
            for fi147 := 0; fi147 < 18; fi147++ {
                v52 = v52
                _ = fi147
            }
        }
        if ((v52 & (v52 << ((2) & 63))) >= (v52 + (v52 << ((319) & 63)))) {
            fmt.Println(v2)
        }
        fmt.Println(v52, v52)
    } else {
        switch -726 {
        case 127:
            if v4 {
                var v148 bool = true
                _ = len(v3)
                v149 := GenS12{F0: 42}
                _ = v148
                _ = v149
            } else {
                var v150 uint8 = 102
                v151 := GenS13{F0: v4, F1: -9223372036854775808, F2: (v150 - v150)}
                v152 := GenS14{F0: -1}
                _ = v150
                _ = v151
                _ = v152
            }
        case 32767:
            v52 = (^(v52))
        case 42:
            var v153 float64 = v1
            _ = v153
        }
    }
    v56 = append(v56, 4294967295)
    if (v2 >= v2) {
        fmt.Println(v2)
        switch v2 {
        case "":
            v1 = (1.5 / float64(1.5))
        default:
            var v154 int32 = -1
            _ = v154
        }
        v155 := GenS15{F0: 885}
        _ = v155
    }
    for fi156 := 0; fi156 < 10; fi156++ {
        v2 = (v2 + (((("foo bar" + v2) + v2) + v2) + (v2 + v2)))
        _ = fi156
    }
    fmt.Println(v52, v52, v52)
    fmt.Println(v52, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v56:", len(v56), v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v56
}
