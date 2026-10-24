package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint16; F2 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 int64; F1 uint64 }
type GenS12 struct { F0 int32; F1 int64 }
type GenS13 struct { F0 string; F1 uint32; F2 int16 }
type GenS14 struct { F0 int32; F1 uint64 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint32; F1 int8 }
type GenS17 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int8 = -18
    v2 := make(map[uint32]int)
    var v3 uint8 = 42
    var v4 uint8 = v3
    if false {
        fmt.Println(v1, v3, v4)
    } else {
        fmt.Println(v4)
        var v5 float32 = 0.0
        _ = v5
    }
    var v6 int16 = 32767
    if false {
        v1 += v1
        _ = len(v2)
        switch "a" {
        case "foo bar":
            v1 += v1
        default:
            var v7 int16 = v6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            switch (fi8 - 128) {
            case 1:
                v3 += v4
            default:
                v9 := GenS1{F0: 227}
                _ = v9
            }
            v10 := []int{(fi8 * -305), fi8, (fi8 << ((65535) & 63))}
            _ = fi8
            _ = v10
        }
    }
    var v11 int16 = v6
    v4 -= 5
    var v12 uint16 = 65535
    var v13 string = "a"
    var v14 uint8 = 108
    var v15 bool = (v13 < (("foo bar" + v13) + v13))
    switch 42 {
    case 4294967295:
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := GenS1{F0: 0}
            v14 -= (v14 << ((2) & 63))
            var v18 uint16 = (v12 >> ((18446744073709551615) & 63))
            _ = fi16
            _ = v17
            _ = v18
        }
    case 127:
        fmt.Println(v15, v11)
    case 32768:
        for fi19 := 0; fi19 < 13; fi19++ {
            if (!((v14 * (v3 + (^(v14)))) < v3)) {
                switch (fi19 << ((1) & 63)) {
                case -805:
                    v11 = 42
                }
                v6 = 32767
                if v15 {
                    _ = len(v2)
                } else {
                    switch (fi19 / (((fi19 | fi19)) | 1)) {
                    case 32767:
                        for fi20 := 0; fi20 < 14; fi20++ {
                            v6 = -384
                            if (1 == 0) {
                                switch v13 {
                                case "foo bar":
                                    switch v13 {
                                    case "hello":
                                        if (!(0.0 < (float64((((1e10 * 1.5) + (0.1 - 3.14159)) / (1.5 - (3.14159 / 1e10)))) + -2.25))) {
                                            var v21 int32 = -1
                                            _ = v21
                                        } else {
                                            if (!(((v13 + (v13 + v13)) > v13) && v15)) {
                                                if (false && true) {
                                                    if v15 {
                                                        var v22 uint8 = v4
                                                        v23 := []int{(^(fi20)), (fi20 + fi19), -9223372036854775808, fi19}
                                                        _ = v22
                                                        _ = v23
                                                    } else {
                                                        var v24 bool = false
                                                        v25 := GenS2{F0: v24}
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                }
                                            } else {
                                                if (0 >= 1) {
                                                    v11 = 42
                                                    var v26 string = v13
                                                    if v15 {
                                                        fmt.Println(v4, fi20)
                                                    }
                                                    _ = v26
                                                }
                                                v27 := GenS1{F0: 661}
                                                _ = v27
                                            }
                                            v28 := GenS2{F0: ((3.14159 / -2.25) <= (float32((3.14159 * 3.14159)) / 1.5))}
                                            _ = v28
                                        }
                                        v14 = 128
                                        v12 = v12
                                    default:
                                        if ((v11 ^ (v11 | v11)) < v11) {
                                            if true {
                                                if (128 >= fi20) {
                                                    v29 := GenS2{F0: true}
                                                    v30 := GenS1{F0: 18446744073709551615}
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    var v31 int16 = (v6 | -195)
                                                    v32 := GenS2{F0: (!(18446744073709551615 >= 4294967295))}
                                                    _ = v31
                                                    _ = v32
                                                }
                                                v6 = (v11 % ((32767) | 1))
                                            } else {
                                                var v33 int64 = 42
                                                if v15 {
                                                    fmt.Println(fi19)
                                                    fmt.Println(v6, v4, fi20)
                                                    var v34 int16 = (v11 & -75)
                                                    _ = v34
                                                }
                                                var v35 uint8 = 220
                                                if v15 {
                                                    var v36 bool = false
                                                    _ = len(v2)
                                                    v37 := GenS2{F0: v15}
                                                    var v38 uint8 = v14
                                                    var v39 int = fi19
                                                    _ = v36
                                                    _ = v37
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    v40 := v2[878]
                                                    var v41 int32 = 42
                                                    if true {
                                                        v11 += v6
                                                    } else {
                                                        v42 := GenS1{F0: 18446744073709551615}
                                                        switch 32767 {
                                                        case 2147483647:
                                                            v15 = (!(0.1 != ((((1e10 + 3.14159) + 0.0) - 0.0) - -2.25)))
                                                        case 32768:
                                                            var v43 int32 = -1
                                                            _ = v43
                                                        case 32767:
                                                            var v44 float64 = 0.0
                                                            _ = v44
                                                        }
                                                        switch (^(fi20)) {
                                                        case -780:
                                                            v14 = (v4 << ((0) & 63))
                                                        case 2147483647:
                                                            v45 := []int{2147483647, v40, 42}
                                                            _ = v45
                                                        case -2147483648:
                                                            v46 := GenS3{F0: v1, F1: (!v15), F2: v4}
                                                            _ = v46
                                                        }
                                                        fmt.Println(v12)
                                                        _ = v42
                                                    }
                                                    if v15 {
                                                        v41 = 583
                                                        if v15 {
                                                            v47 := v2[4294967295]
                                                            var v48 float64 = (3.14159 + (float64((((3.14159 + -2.25) - float32(0.0)) / 0.1)) - 0.0))
                                                            v3 = 255
                                                            v49 := GenS3{F0: (v1 + (v1 >> ((255) & 63))), F1: false, F2: (v4 << ((4294967295) & 63))}
                                                            _ = v47
                                                            _ = v48
                                                            _ = v49
                                                        } else {
                                                            fmt.Println(fi19)
                                                            var v50 int32 = v41
                                                            if v15 {
                                                                switch fi20 {
                                                                case 128:
                                                                    v51 := v2[512]
                                                                    _ = v51
                                                                case 32767:
                                                                    if true {
                                                                        if ((!v15) || v15) {
                                                                            v2[0] = 1
                                                                        }
                                                                        fmt.Println(fi19, v1, fi19)
                                                                    } else {
                                                                        v2[1] = (fi20 % ((fi19) | 1))
                                                                        v14 += 42
                                                                    }
                                                                default:
                                                                    if (((fi19 & (^(fi20))) == fi20) || (4294967295 < 42)) {
                                                                        v52 := GenS4{F0: v33, F1: 42}
                                                                        var v53 int64 = -1
                                                                        _ = v52
                                                                        _ = v53
                                                                    } else {
                                                                        if false {
                                                                            v54 := v2[uint32(v4)]
                                                                            delete(v2, 739)
                                                                            _ = v54
                                                                        } else {
                                                                            v12 = (v12 * 744)
                                                                            var v55 bool = false
                                                                            _ = v55
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            if v15 {
                                                                v2[517] = fi19
                                                                var v56 int8 = (v1 << ((4294967295) & 63))
                                                                if true {
                                                                    delete(v2, 1)
                                                                    switch (-(v40)) {
                                                                    case 33:
                                                                        var v57 int64 = (v33 ^ (-(v33)))
                                                                        _ = v57
                                                                    case -9223372036854775808:
                                                                        v50 -= (v50 / ((-2147483648) | 1))
                                                                    case -880:
                                                                        v58 := GenS5{F0: 783, F1: (v12 - v12), F2: v1}
                                                                        _ = v58
                                                                    }
                                                                    var v59 uint64 = 42
                                                                    delete(v2, 0)
                                                                    if false {
                                                                        if true {
                                                                            var v60 uint = 732
                                                                            var v61 float32 = float32(0.0)
                                                                            delete(v2, 2)
                                                                            v62 := GenS1{F0: 18446744073709551615}
                                                                            var v63 uint8 = v35
                                                                            _ = v60
                                                                            _ = v61
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            fmt.Println(v4, v56, v11)
                                                                            var v64 bool = (v15 || (v3 <= (v35 & (^(v14)))))
                                                                            v65 := []int{9223372036854775807, fi20, (fi19 % (((fi19 % ((v40) | 1))) | 1)), fi20}
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    } else {
                                                                        v2[42] = int(v11)
                                                                        switch 32767 {
                                                                        case 32768:
                                                                            var v66 string = "a"
                                                                            _ = v66
                                                                        case 369:
                                                                            var v67 int32 = int32(v59)
                                                                            _ = v67
                                                                        case 957:
                                                                            delete(v2, 536)
                                                                            v68 := GenS6{F0: int64(fi20)}
                                                                            _ = v68
                                                                        default:
                                                                            v59 = 42
                                                                            if (((0.0 / 1e10) > float64(0.0)) && v15) {
                                                                                v4 = v3
                                                                                var v69 float64 = ((-2.25 / 1e10) / 1.5)
                                                                                if (0 < 255) {
                                                                                    v13 = v13
                                                                                    if false {
                                                                                        v70 := []int{-9223372036854775808}
                                                                                        switch int(v6) {
                                                                                        case 32767:
                                                                                            v3 = 255
                                                                                        }
                                                                                        var v71 string = ""
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                    } else {
                                                                                        v6 += (-(v11))
                                                                                    }
                                                                                } else {
                                                                                    var v72 int16 = (v6 & v6)
                                                                                    switch v13 {
                                                                                    case "foo bar":
                                                                                        v73 := GenS1{F0: 2}
                                                                                        var v74 bool = (!false)
                                                                                        _ = v73
                                                                                        _ = v74
                                                                                    }
                                                                                    v12 = 65535
                                                                                    _ = v72
                                                                                }
                                                                                v15 = v15
                                                                                _ = v69
                                                                            }
                                                                            if false {
                                                                                fmt.Println(v40, v12, v11)
                                                                            } else {
                                                                                v15 = v15
                                                                                v75 := v2[2]
                                                                                if v15 {
                                                                                    switch ("" + v13) {
                                                                                    case "foo bar":
                                                                                        v40 = 1
                                                                                    case "":
                                                                                        switch "" {
                                                                                        case "a":
                                                                                            v35 = v35
                                                                                            switch fi20 {
                                                                                            case -1:
                                                                                                fmt.Println(v35)
                                                                                            default:
                                                                                                var v76 int = -1
                                                                                                fmt.Println(v41, v40)
                                                                                                switch 9223372036854775807 {
                                                                                                case 32767:
                                                                                                    v76 = v40
                                                                                                }
                                                                                                _ = v76
                                                                                            }
                                                                                        default:
                                                                                            if (("\n\t" >= "a") && true) {
                                                                                                v77 := []int{1, fi20, (fi20 * 651), 128}
                                                                                                v12 = uint16(v14)
                                                                                                fmt.Println(v33, v15, v14)
                                                                                                v78 := GenS2{F0: false}
                                                                                                v79 := GenS7{F0: (v12 << ((42) & 63)), F1: v15}
                                                                                                _ = v77
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            }
                                                                                        }
                                                                                        v80 := []int{128}
                                                                                        _ = v80
                                                                                    default:
                                                                                        v81 := GenS3{F0: (v56 & 81), F1: false, F2: 73}
                                                                                        _ = v81
                                                                                    }
                                                                                }
                                                                                if (true && false) {
                                                                                    var v82 bool = (-233 != v41)
                                                                                    v83 := GenS7{F0: uint16(v3), F1: (v50 <= 42)}
                                                                                    _ = v82
                                                                                    _ = v83
                                                                                }
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                    }
                                                                    var v84 uint64 = (v59 % (((v59 + 42)) | 1))
                                                                    fmt.Println(v4, v6)
                                                                    _ = v59
                                                                    _ = v84
                                                                }
                                                                if (0 == v6) {
                                                                    if (((!v15) && (v33 > v33)) && false) {
                                                                        v33 += v33
                                                                    }
                                                                    v85 := []int{v40, fi20}
                                                                    var v86 string = ""
                                                                    fmt.Println(fi20, v15)
                                                                    _ = v85
                                                                    _ = v86
                                                                } else {
                                                                    v87 := []int{v40, (v40 / ((fi19) | 1))}
                                                                    if (65535 < 42) {
                                                                        _ = len(v87)
                                                                        v88 := GenS8{F0: (^(v14)), F1: v1}
                                                                        _ = v88
                                                                    } else {
                                                                        var v89 int8 = v1
                                                                        v90 := GenS9{F0: ((v12 & 0) >= (v12 * 2))}
                                                                        var v91 bool = (((786 < 882) && v15) && v15)
                                                                        _ = v89
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                    _ = v87
                                                                }
                                                                var v92 bool = (1e10 != -2.25)
                                                                fmt.Println(v3, v33)
                                                                _ = v56
                                                                _ = v92
                                                            } else {
                                                                v93 := GenS10{F0: v11, F1: 4294967295}
                                                                v94 := GenS11{F0: 42, F1: 4294967295}
                                                                var v95 string = v13
                                                                _ = v93
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                            fmt.Println(v11)
                                                            _ = v50
                                                        }
                                                        var v96 string = ("hello" + v13)
                                                        _ = v96
                                                    }
                                                    _ = v40
                                                    _ = v41
                                                }
                                                if (v33 < v33) {
                                                    var v97 uint8 = v3
                                                    v12 = v12
                                                    if v15 {
                                                        switch v13 {
                                                        case "a":
                                                            var v98 int16 = (v6 + v11)
                                                            _ = v98
                                                        default:
                                                            var v99 int8 = v1
                                                            _ = v99
                                                        }
                                                        var v100 float32 = ((0.1 * 1.5) - (((0.0 / ((1e10 + float32(((-2.25 * 3.14159) - 0.1))) / 1.0)) / ((-2.25 / 1.0) / 3.14159)) / -2.25))
                                                        _ = v100
                                                    }
                                                    _ = v97
                                                }
                                                v12 = (^(v12))
                                                _ = v33
                                                _ = v35
                                            }
                                        }
                                    }
                                case "\n\t":
                                    if false {
                                        v2[1] = (fi20 - fi20)
                                        v101 := []int{}
                                        switch "hello" {
                                        case "hello":
                                            switch ((((v13 + (((("\n\t" + ("\n\t" + v13)) + "foo bar") + (v13 + v13)) + "\n\t")) + v13) + "") + "a") {
                                            case "a":
                                                if false {
                                                    switch fi19 {
                                                    case 32767:
                                                        var v102 uint8 = (^(v3))
                                                        _ = v102
                                                    case 128:
                                                        var v103 uint8 = 127
                                                        _ = v103
                                                    case 0:
                                                        if v15 {
                                                            delete(v2, 4294967295)
                                                            var v104 float64 = -2.25
                                                            var v105 int8 = 42
                                                            _ = len(v2)
                                                            v104 = v104
                                                            _ = v104
                                                            _ = v105
                                                        }
                                                    default:
                                                        var v106 uint64 = 259
                                                        _ = v106
                                                    }
                                                    v14 = (^(v14))
                                                }
                                            }
                                        case "":
                                            var v107 uint16 = 32767
                                            _ = v107
                                        default:
                                            var v108 int32 = 0
                                            _ = v108
                                        }
                                        _ = v101
                                    } else {
                                        v13 = ((("\n\t" + v13) + v13) + v13)
                                        v13 = ((("" + v13) + v13) + "\n\t")
                                        switch 921 {
                                        case 4294967295:
                                            fmt.Println(v4)
                                        case 0:
                                            v3 += (^(v14))
                                        default:
                                            v109 := GenS12{F0: -2147483648, F1: -978}
                                            _ = v109
                                        }
                                        v6 -= (v6 & (v11 - -1))
                                        if v15 {
                                            switch fi20 {
                                            case 9223372036854775807:
                                                switch int(v1) {
                                                case 1:
                                                    var v110 float64 = ((0.0 / (float64(3.14159) / 1.0)) * ((1.5 - 0.0) + 1.5))
                                                    _ = v110
                                                case 645:
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v3, v3)
                                                    case "\n\t":
                                                        v111 := GenS11{F0: 42, F1: 18446744073709551615}
                                                        if v15 {
                                                            v112 := GenS13{F0: (((v13 + "") + "a") + ((v13 + v13) + v13)), F1: 42, F2: v11}
                                                            _ = v112
                                                        }
                                                        if v15 {
                                                            v113 := GenS4{F0: 9223372036854775807, F1: 2}
                                                            v114 := GenS14{F0: 1, F1: 684}
                                                            var v115 int32 = 42
                                                            var v116 float64 = (1.5 * -2.25)
                                                            var v117 uint32 = 2147483648
                                                            _ = v113
                                                            _ = v114
                                                            _ = v115
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                        var v118 int64 = 11
                                                        _ = v111
                                                        _ = v118
                                                    default:
                                                        if (v1 > v1) {
                                                            var v119 float64 = 3.14159
                                                            var v120 uint = uint(v12)
                                                            _ = v119
                                                            _ = v120
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    var v121 float32 = float32(-2.25)
                                                    _ = v121
                                                }
                                            case -246:
                                                var v122 uint8 = (v3 + 127)
                                                _ = v122
                                            }
                                            if ((v1 - -1) < (v1 - v1)) {
                                                v1 -= (v1 & v1)
                                                v123 := GenS4{F0: -231, F1: 2}
                                                _ = v123
                                            } else {
                                                switch 4294967295 {
                                                case 178:
                                                    v13 = ("" + v13)
                                                default:
                                                    v11 = v11
                                                }
                                                v11 += v11
                                                switch fi20 {
                                                case -2147483648:
                                                    var v124 uint8 = v4
                                                    _ = v124
                                                default:
                                                    v125 := GenS6{F0: 1}
                                                    _ = v125
                                                }
                                            }
                                            v1 -= v1
                                            if false {
                                                var v126 int64 = -9223372036854775808
                                                v14 -= v3
                                                var v127 uint16 = 408
                                                _ = v126
                                                _ = v127
                                            } else {
                                                v128 := GenS15{F0: 2147483647}
                                                switch v13 {
                                                case "":
                                                    var v129 string = v13
                                                    _ = v129
                                                case "hello":
                                                    _ = len(v2)
                                                }
                                                v130 := GenS16{F0: 2147483647, F1: (v1 + (^(v1)))}
                                                _ = v128
                                                _ = v130
                                            }
                                            var v131 uint = 42
                                            _ = v131
                                        }
                                    }
                                }
                                if ((v12 & v12) == 81) {
                                    v13 = "a"
                                    var v132 uint = 255
                                    var v133 int8 = (-(v1))
                                    _ = v132
                                    _ = v133
                                }
                            } else {
                                switch fi20 {
                                case 518:
                                    switch (fi20 >> ((0) & 63)) {
                                    case -893:
                                        fmt.Println(v1)
                                    }
                                case 128:
                                    if v15 {
                                        var v134 int = (-(fi20))
                                        var v135 int32 = 878
                                        v1 += v1
                                        _ = v134
                                        _ = v135
                                    }
                                case 9223372036854775807:
                                    if v15 {
                                        var v136 float32 = float32(1.5)
                                        var v137 float32 = v136
                                        var v138 bool = (v12 <= v12)
                                        _ = v136
                                        _ = v137
                                        _ = v138
                                    }
                                }
                                v139 := GenS10{F0: v6, F1: 139}
                                _ = v139
                            }
                            _ = fi20
                        }
                    case -1:
                        v1 -= 127
                    default:
                        var v140 uint8 = 255
                        _ = v140
                    }
                    v141 := []int{fi19, 127, fi19, (fi19 + 128)}
                    _ = v141
                }
            }
            _ = fi19
        }
    default:
        if (42 < 403) {
            switch "a" {
            case "a":
                if true {
                    switch "\n\t" {
                    case "\n\t":
                        v142 := GenS8{F0: v3, F1: (-(v1))}
                        _ = v142
                    default:
                        v143 := GenS17{F0: v14, F1: (v12 ^ (^(v12)))}
                        _ = v143
                    }
                }
            case "foo bar":
                v144 := []int{814, -2147483648}
                _ = v144
            case "":
                switch (v13 + ("a" + v13)) {
                case "\n\t":
                    v15 = ("foo bar" != (("foo bar" + ((v13 + v13) + v13)) + v13))
                case "":
                    for fi145 := 0; fi145 < 17; fi145++ {
                        for fi146 := 0; fi146 < 12; fi146++ {
                            if (float64(1.5) == (3.14159 / 1.0)) {
                                v1 += -54
                                v147 := GenS16{F0: 4294967295, F1: 0}
                                v12 -= 65535
                                _ = v147
                            }
                            _ = fi146
                        }
                        _ = fi145
                    }
                }
            }
        } else {
            for fi148 := 0; fi148 < 20; fi148++ {
                v149 := []int{fi148, (fi148 >> ((uint(v14)) & 63)), (fi148 * (fi148 ^ 128))}
                _ = fi148
                _ = v149
            }
            switch 32768 {
            case 4294967295:
                for fi150 := 0; fi150 < 12; fi150++ {
                    for fi151 := 0; fi151 < 12; fi151++ {
                        var v152 uint32 = uint32(v12)
                        _ = fi151
                        _ = v152
                    }
                    var v153 uint64 = uint64(v3)
                    _ = fi150
                    _ = v153
                }
            default:
                for fi154 := 0; fi154 < 8; fi154++ {
                    switch ("hello" + (v13 + v13)) {
                    case "hello":
                        v6 = v6
                    case "\n\t":
                        v2[2] = 0
                    case "":
                        v3 = (v3 % ((uint8(v12)) | 1))
                    default:
                        v155 := v2[578]
                        _ = v155
                    }
                    var v156 int32 = int32(v14)
                    v157 := GenS5{F0: 4294967295, F1: (v12 >> ((2) & 63)), F2: v1}
                    _ = fi154
                    _ = v156
                    _ = v157
                }
            }
            for fi158 := 0; fi158 < 14; fi158++ {
                for fi159 := 0; fi159 < 20; fi159++ {
                    v13 = "\n\t"
                    v14 += (v3 / (((v4 ^ v14)) | 1))
                    v6 = 42
                    v15 = (((0.0 + (1.5 + (0.0 + 3.14159))) + float32(1e10)) != ((float32((0.1 * float64((3.14159 + 1.5)))) - (0.0 - 0.1)) - 3.14159))
                    _ = fi159
                }
                for fi160 := 0; fi160 < 6; fi160++ {
                    var v161 uint16 = v12
                    var v162 int64 = 1
                    v162 -= (v162 ^ (v162 * v162))
                    _ = fi160
                    _ = v161
                    _ = v162
                }
                _ = fi158
            }
        }
    }
    fmt.Println(v14, v15, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
