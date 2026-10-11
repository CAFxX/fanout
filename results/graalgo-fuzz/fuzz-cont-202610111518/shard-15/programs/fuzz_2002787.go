package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int8; F1 uint64; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint32; F1 string; F2 int }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint; F1 uint16 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int8; F1 int64 }
type GenS17 struct { F0 string; F1 bool; F2 uint64 }
type GenS18 struct { F0 int64; F1 int8 }
type GenS19 struct { F0 string; F1 uint32 }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 int8; F1 bool }
type GenS22 struct { F0 int32; F1 uint16; F2 uint }
type GenS23 struct { F0 int8; F1 bool; F2 uint64 }
func genFn0(p0 int64) uint32 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 uint8 = 165
            var v7 uint64 = 266
            v8 := GenS1{F0: true, F1: true}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    var v9 int = 9223372036854775807
    if false {
        var v10 bool = true
        var v11 uint = 65535
        v12 := []int{(-(v9)), (v9 / ((-9223372036854775808) | 1))}
        _ = v10
        _ = v11
        _ = v12
    } else {
        if false {
            fmt.Println(v9, v9)
            var v13 float32 = 1.5
            var v14 int32 = 2147483647
            _ = v13
            _ = v14
        } else {
            v15 := GenS2{F0: 42, F1: -1, F2: 42}
            _ = v15
        }
        p0 = 1
    }
    _ = p0
    _ = v9
    return 0
}
func genFn1() uint64 {
    for fi16 := 0; fi16 < 15; fi16++ {
        switch fi16 {
        case 0:
            var v17 uint16 = 1
            _ = v17
        case 422:
            fmt.Println(fi16)
        default:
            if (38 < 127) {
                for fi18 := 0; fi18 < 10; fi18++ {
                    v19 := GenS1{F0: false, F1: true}
                    var v20 int8 = 0
                    if true {
                        if (!false) {
                            switch (fi18 - fi16) {
                            case -704:
                                v21 := GenS2{F0: 2, F1: -423, F2: 1}
                                _ = v21
                            case 1:
                                switch 32768 {
                                case 0:
                                    switch 128 {
                                    case -2147483648:
                                        if (249 <= 42) {
                                            v22 := []int{fi18}
                                            if false {
                                                v20 = (v20 & (v20 + v20))
                                                var v23 int64 = -1
                                                var v24 uint8 = 55
                                                if (0 < 692) {
                                                    fmt.Println(v20, v23, v24)
                                                } else {
                                                    if len(v22) > 0 {
                                                        v22[0] = 128
                                                    }
                                                    var v25 uint16 = 985
                                                    _ = v25
                                                }
                                                _ = v23
                                                _ = v24
                                            } else {
                                                fmt.Println(fi18, fi18, fi16)
                                                switch int(v20) {
                                                case 42:
                                                    var v26 int = (fi18 / (((-(fi16))) | 1))
                                                    _ = v26
                                                default:
                                                    fmt.Println(v20)
                                                }
                                            }
                                            v20 -= 1
                                            _ = v22
                                        } else {
                                            if true {
                                                if false {
                                                    v27 := []int{1, fi16, fi16}
                                                    fmt.Println(fi18)
                                                    if true {
                                                        v27 = append(v27, (fi18 * -1))
                                                        var v28 int64 = 9223372036854775807
                                                        if len(v27) > 0 {
                                                            v27[0] = fi16
                                                        }
                                                        _ = v28
                                                    } else {
                                                        if true {
                                                            if false {
                                                                switch (^(fi18)) {
                                                                case 1:
                                                                    v29 := GenS3{F0: 42}
                                                                    v20 = v20
                                                                    if len(v27) >= 2 {
                                                                        v27 = v27[1:len(v27)]
                                                                    }
                                                                    _ = v29
                                                                case 805:
                                                                    v30 := GenS3{F0: 0}
                                                                    _ = v30
                                                                case 127:
                                                                    var v31 int8 = v20
                                                                    _ = v31
                                                                default:
                                                                    if len(v27) >= 2 {
                                                                        v27 = v27[1:len(v27)]
                                                                    }
                                                                }
                                                                v20 += (v20 >> ((2) & 63))
                                                                if false {
                                                                    fmt.Println(fi18)
                                                                    v20 = -55
                                                                    var v32 string = ("foo bar" + (("foo bar" + "foo bar") + ""))
                                                                    _ = v32
                                                                } else {
                                                                    var v33 uint16 = 65535
                                                                    _ = len(v27)
                                                                    _ = v33
                                                                }
                                                            } else {
                                                                v20 = v20
                                                                if (((fi18 & fi16) == (fi16 >> ((81) & 63))) && false) {
                                                                    var v34 uint16 = 1
                                                                    _ = v34
                                                                }
                                                                var v35 uint8 = 127
                                                                _ = v35
                                                            }
                                                            switch fi18 {
                                                            case 4294967295:
                                                                if len(v27) > 0 {
                                                                    v27[0] = -131
                                                                }
                                                            default:
                                                                v36 := GenS4{F0: 0, F1: true}
                                                                _ = v36
                                                            }
                                                        } else {
                                                            var v37 int16 = -266
                                                            var v38 uint16 = 32767
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                    }
                                                    _ = v27
                                                }
                                            } else {
                                                v39 := []int{fi18}
                                                if true {
                                                    switch "" {
                                                    case "hello":
                                                        v20 -= (^(v20))
                                                    case "a":
                                                        v20 -= v20
                                                    case "":
                                                        if len(v39) > 0 {
                                                            v39[0] = 368
                                                        }
                                                    }
                                                    var v40 float64 = 0.0
                                                    _ = v40
                                                }
                                                _ = v39
                                            }
                                            var v41 float64 = (-2.25 + 1e10)
                                            v42 := []int{fi18, (fi18 % ((-657) | 1)), fi18, fi18}
                                            _ = v41
                                            _ = v42
                                        }
                                    default:
                                        if true {
                                            switch "hello" {
                                            case "hello":
                                                v43 := GenS5{F0: 100, F1: 4294967295, F2: -9223372036854775808}
                                                _ = v43
                                            case "\n\t":
                                                if (395 <= 0) {
                                                    switch (fi16 % ((32768) | 1)) {
                                                    case -731:
                                                        fmt.Println(fi16, v20, fi18)
                                                    case -772:
                                                        if true {
                                                            var v44 bool = true
                                                            v44 = (v20 < v20)
                                                            switch (fi16 ^ (fi18 >> ((0) & 63))) {
                                                            case 9223372036854775807:
                                                                var v45 float32 = -2.25
                                                                _ = v45
                                                            default:
                                                                v44 = v44
                                                            }
                                                            _ = v44
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        if (false && (127 < fi16)) {
                                                            var v46 bool = (!false)
                                                            _ = v46
                                                        }
                                                    case "\n\t":
                                                        v47 := []int{fi16, fi18, fi18, fi16}
                                                        _ = v47
                                                    case "foo bar":
                                                        var v48 int64 = -1
                                                        _ = v48
                                                    default:
                                                        var v49 uint = 881
                                                        _ = v49
                                                    }
                                                    v20 -= 1
                                                } else {
                                                    if false {
                                                        v20 = v20
                                                        v50 := GenS5{F0: v20, F1: 620, F2: -9223372036854775808}
                                                        _ = v50
                                                    }
                                                    if (!false) {
                                                        fmt.Println(fi16)
                                                        v20 = 100
                                                        v51 := GenS6{F0: 128, F1: 42, F2: (fi16 / ((42) | 1))}
                                                        _ = v51
                                                    } else {
                                                        if (4294967295 < 18446744073709551615) {
                                                            v20 -= (v20 << ((4294967295) & 63))
                                                            switch fi18 {
                                                            case -808:
                                                                if false {
                                                                    v20 -= v20
                                                                } else {
                                                                    if false {
                                                                        var v52 int64 = 42
                                                                        var v53 uint64 = 204
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            var v54 bool = true
                                                                            _ = v54
                                                                        default:
                                                                            v20 = (v20 << ((0) & 63))
                                                                        }
                                                                        _ = v52
                                                                        _ = v53
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(fi18, fi16)
                                                                            if (760 < 9223372036854775807) {
                                                                                var v55 int = (fi16 ^ (fi16 * (fi16 << ((178) & 63))))
                                                                                var v56 uint8 = 14
                                                                                fmt.Println(fi16, fi18, v55)
                                                                                _ = v55
                                                                                _ = v56
                                                                            } else {
                                                                                v57 := []int{fi18, (^(fi18)), -640}
                                                                                if len(v57) >= 2 {
                                                                                    v57 = v57[1:len(v57)]
                                                                                }
                                                                                if true {
                                                                                    var v58 uint16 = 255
                                                                                    _ = v58
                                                                                } else {
                                                                                    v20 = (^(v20))
                                                                                    var v59 uint8 = 104
                                                                                    _ = v59
                                                                                }
                                                                                _ = v57
                                                                            }
                                                                        } else {
                                                                            v20 += v20
                                                                            v60 := []int{-9223372036854775808}
                                                                            fmt.Println(fi16, v20, fi18)
                                                                            _ = v60
                                                                        }
                                                                        v61 := []int{fi16, -575}
                                                                        if len(v61) > 0 {
                                                                            v61[0] = fi16
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                    v62 := []int{(fi18 + int(v20))}
                                                                    _ = len(v62)
                                                                    v20 -= (v20 >> ((18446744073709551615) & 63))
                                                                    v20 = 42
                                                                    if len(v62) >= 2 {
                                                                        v62 = v62[1:len(v62)]
                                                                    }
                                                                    _ = v62
                                                                }
                                                                var v63 int8 = (v20 | v20)
                                                                v64 := GenS7{F0: (("a" + ("hello" + "")) + "foo bar"), F1: 42}
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                        } else {
                                                            v65 := GenS8{F0: ""}
                                                            _ = v65
                                                        }
                                                        v66 := GenS7{F0: "", F1: 1}
                                                        fmt.Println(v20, fi18)
                                                        if (true || false) {
                                                            v67 := []int{fi16, 664}
                                                            var v68 int16 = 225
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            v69 := GenS9{F0: (-(v20)), F1: 42}
                                                            v70 := []int{fi16}
                                                            var v71 uint16 = 931
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        switch ("a" + "") {
                                                        case "foo bar":
                                                            v72 := []int{}
                                                            _ = v72
                                                        case "hello":
                                                            if true {
                                                                v73 := GenS10{F0: 18446744073709551615}
                                                                v20 -= v20
                                                                _ = v73
                                                            }
                                                        case "a":
                                                            fmt.Println(fi16, v20, fi16)
                                                        }
                                                        _ = v66
                                                    }
                                                }
                                            }
                                        } else {
                                            if true {
                                                var v74 uint16 = 42
                                                var v75 int32 = -1
                                                switch fi18 {
                                                case 9223372036854775807:
                                                    v76 := GenS11{F0: 2147483648, F1: ("" + "\n\t"), F2: (fi18 * 518)}
                                                    _ = v76
                                                case 32768:
                                                    var v77 int32 = v75
                                                    _ = v77
                                                case 18:
                                                    if false {
                                                        fmt.Println(fi18, v74, v74)
                                                    } else {
                                                        v78 := GenS5{F0: (v20 / (((v20 >> ((0) & 63))) | 1)), F1: 564, F2: 9223372036854775807}
                                                        _ = v78
                                                    }
                                                }
                                                _ = v74
                                                _ = v75
                                            }
                                            if false {
                                                if ((float64((0.1 + 3.14159)) / -2.25) != (0.0 + ((1.5 + 0.0) - ((0.0 + 1e10) * 1e10)))) {
                                                    switch (fi18 / ((fi18) | 1)) {
                                                    case -201:
                                                        v20 = 0
                                                    default:
                                                        var v79 string = "\n\t"
                                                        _ = v79
                                                    }
                                                    fmt.Println(fi18)
                                                    if ((!((!(((1.5 + 0.1) < (float64(0.0) * (0.1 - 1e10))) && (42 <= 255))) && (2147483647 == 856))) || false) {
                                                        fmt.Println(fi18)
                                                        v20 = (v20 % ((-90) | 1))
                                                    }
                                                    var v80 int64 = -906
                                                    var v81 int64 = v80
                                                    v80 -= -642
                                                    var v82 uint = 4294967295
                                                    _ = v80
                                                    _ = v81
                                                    _ = v82
                                                } else {
                                                    fmt.Println(v20)
                                                    var v83 int8 = (v20 / ((v20) | 1))
                                                    var v84 float64 = 3.14159
                                                    _ = v83
                                                    _ = v84
                                                }
                                                v85 := GenS4{F0: 9223372036854775807, F1: (!((v20 ^ v20) <= v20))}
                                                var v86 uint16 = 0
                                                _ = v85
                                                _ = v86
                                            } else {
                                                v20 = -31
                                            }
                                        }
                                    }
                                case 2147483647:
                                    v87 := GenS3{F0: 255}
                                    _ = v87
                                case 985:
                                    var v88 uint32 = 1
                                    _ = v88
                                }
                            case -2147483648:
                                if ("hello" == "hello") {
                                    var v89 int8 = v20
                                    _ = v89
                                } else {
                                    v90 := []int{(fi16 & (fi18 % ((0) | 1)))}
                                    _ = v90
                                }
                            }
                        }
                        v91 := []int{-276, 127}
                        var v92 int64 = 1
                        _ = v91
                        _ = v92
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                switch fi16 {
                case 0:
                    var v93 bool = true
                    _ = v93
                case 32768:
                    v94 := []int{}
                    _ = v94
                case -9223372036854775808:
                    v95 := []int{fi16, fi16}
                    _ = v95
                default:
                    for fi96 := 0; fi96 < 7; fi96++ {
                        var v97 string = ("a" + ("\n\t" + "hello"))
                        _ = fi96
                        _ = v97
                    }
                }
                var v98 int32 = 0
                _ = v98
            }
        }
        var v99 float64 = 0.1
        _ = fi16
        _ = v99
    }
    var v100 int = 9223372036854775807
    v100 += 0
    var v101 bool = false
    _ = v100
    _ = v101
    return 4294967295
}

func main() {
    var v1 int = 2147483647
    var v2 uint16 = 203
    v3 := make(map[uint32]string)
    v1 = 9223372036854775807
    if (-863 <= -643) {
        for fi102 := 0; fi102 < 9; fi102++ {
            v103 := GenS10{F0: 890}
            if (!(!(3.14159 != 1e10))) {
                v2 -= (v2 & v2)
            }
            v104 := []int{}
            _ = fi102
            _ = v103
            _ = v104
        }
        if false {
            if false {
                v2 = (v2 << ((255) & 63))
                fmt.Println(v2)
            } else {
                if ((!((!true) || true)) && false) {
                    v105 := GenS12{F0: 32767, F1: false}
                    var v106 string = ""
                    _ = v105
                    _ = v106
                }
            }
            var v107 uint64 = 42
            _ = v107
        }
        if true {
            v2 = (v2 % ((v2) | 1))
            switch -2147483648 {
            case 32767:
                switch (v1 & (-(v1))) {
                case -158:
                    switch v1 {
                    case -1:
                        if false {
                            v2 = (v2 | (v2 << ((901) & 63)))
                            var v108 uint8 = 0
                            var v109 int = v1
                            _ = v108
                            _ = v109
                        }
                    default:
                        v1 = (-(v1))
                    }
                case 32768:
                    v3[genFn0(-463)] = "\n\t"
                case 394:
                    switch (v1 - (^(v1))) {
                    case 0:
                        switch (v1 & v1) {
                        case -9223372036854775808:
                            if true {
                                v2 = v2
                            }
                        case 0:
                            if true {
                                v2 = (v2 & 2)
                                var v110 uint64 = genFn1()
                                _ = v110
                            } else {
                                for fi111 := 0; fi111 < 10; fi111++ {
                                    var v112 float32 = (-2.25 * (((((1e10 + (1e10 * -2.25)) / 1.0) - 3.14159) + -2.25) - -2.25))
                                    _ = fi111
                                    _ = v112
                                }
                                var v113 int8 = 100
                                _ = v113
                            }
                        default:
                            v114 := GenS13{F0: 718}
                            _ = v114
                        }
                    case 127:
                        if false {
                            var v115 uint16 = v2
                            _ = v115
                        }
                    }
                default:
                    switch v1 {
                    case -431:
                        v116 := GenS14{F0: 65535, F1: v2}
                        _ = v116
                    case -9223372036854775808:
                        v117 := GenS9{F0: 1, F1: 2}
                        _ = v117
                    case 381:
                        if true {
                            if (-827 <= 1) {
                                v2 -= (v2 & v2)
                                v2 = (v2 & v2)
                                var v118 float32 = 0.0
                                _ = v118
                            } else {
                                v1 = (v1 ^ 128)
                                fmt.Println(v2)
                                var v119 float32 = -2.25
                                _ = v119
                            }
                            v120 := v3[2147483648]
                            _ = v120
                        } else {
                            if (1 == 70) {
                                var v121 int64 = 0
                                var v122 int64 = (v121 % (((v121 % ((v121) | 1))) | 1))
                                _ = v121
                                _ = v122
                            } else {
                                v2 += 111
                            }
                            var v123 uint32 = 1
                            _ = v123
                        }
                    }
                }
            case -2147483648:
                v124 := GenS15{F0: -1}
                _ = v124
            case 3:
                if (42 <= 203) {
                    if true {
                        for fi125 := 0; fi125 < 13; fi125++ {
                            if false {
                                v126 := []int{(v1 + (fi125 + fi125))}
                                _ = len(v3)
                                _ = v126
                            } else {
                                v2 += (v2 << ((4294967295) & 63))
                                fmt.Println(v1, v2, fi125)
                            }
                            v1 += v1
                            for fi127 := 0; fi127 < 12; fi127++ {
                                delete(v3, 65535)
                                fmt.Println(fi127)
                                var v128 int64 = -97
                                v129 := []int{(v1 >> ((0) & 63))}
                                v1 = v1
                                _ = fi127
                                _ = v128
                                _ = v129
                            }
                            _ = fi125
                        }
                    }
                    if (0 <= genFn1()) {
                        v130 := GenS5{F0: 20, F1: 42, F2: 42}
                        if true {
                            var v131 uint32 = 2147483647
                            _ = v131
                        } else {
                            var v132 bool = ((78 == 196) || true)
                            _ = v132
                        }
                        for fi133 := 0; fi133 < 2; fi133++ {
                            switch "hello" {
                            case "\n\t":
                                v2 = v2
                            case "":
                                if false {
                                    for fi134 := 0; fi134 < 11; fi134++ {
                                        v2 -= (v2 | (v2 * (v2 + (v2 + v2))))
                                        _ = fi134
                                    }
                                    var v135 bool = false
                                    fmt.Println(fi133)
                                    _ = v135
                                } else {
                                    v136 := GenS10{F0: 0}
                                    var v137 bool = false
                                    v1 = -2147483648
                                    _ = v136
                                    _ = v137
                                }
                            case "foo bar":
                                var v138 uint16 = v2
                                _ = v138
                            }
                            var v139 float32 = 0.1
                            v2 += (^(v2))
                            _ = fi133
                            _ = v139
                        }
                        _ = v130
                    } else {
                        if true {
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(v2)
                            case "a":
                                var v140 int = v1
                                _ = v140
                            case "":
                                var v141 uint32 = 42
                                _ = v141
                            default:
                                switch v1 {
                                case 9223372036854775807:
                                    v2 = 425
                                case 0:
                                    var v142 float32 = 1.5
                                    _ = v142
                                case 2147483647:
                                    if false {
                                        switch (v1 >> ((628) & 63)) {
                                        case -935:
                                            fmt.Println(v1)
                                        case 2147483647:
                                            var v143 uint64 = 565
                                            _ = v143
                                        default:
                                            var v144 int64 = 9223372036854775807
                                            _ = v144
                                        }
                                        v145 := GenS13{F0: 587}
                                        _ = v145
                                    }
                                }
                            }
                            if ("foo bar" == (("\n\t" + ("foo bar" + "foo bar")) + "")) {
                                if true {
                                    if (!(false && (42 != 4294967295))) {
                                        var v146 float64 = 0.0
                                        var v147 uint64 = 4294967295
                                        _ = v146
                                        _ = v147
                                    } else {
                                        v1 += (-(v1))
                                        for fi148 := 0; fi148 < 12; fi148++ {
                                            var v149 string = (("\n\t" + "") + ("foo bar" + "\n\t"))
                                            var v150 float64 = -2.25
                                            _ = fi148
                                            _ = v149
                                            _ = v150
                                        }
                                        v151 := []int{}
                                        _ = v151
                                    }
                                }
                            } else {
                                for fi152 := 0; fi152 < 17; fi152++ {
                                    v153 := GenS13{F0: 2147483647}
                                    for fi154 := 0; fi154 < 17; fi154++ {
                                        if (!(false || false)) {
                                            if false {
                                                if (312 != 65535) {
                                                    if (true || (42 <= (v2 * (v2 ^ 65535)))) {
                                                        var v155 float64 = 0.0
                                                        _ = v155
                                                    } else {
                                                        switch ("a" + "foo bar") {
                                                        case "foo bar":
                                                            var v156 uint32 = 0
                                                            _ = v156
                                                        case "a":
                                                            v157 := GenS16{F0: -1, F1: 9223372036854775807}
                                                            _ = v157
                                                        }
                                                        v3[genFn0(673)] = ""
                                                    }
                                                } else {
                                                    v1 += (fi154 >> ((4294967295) & 63))
                                                    v2 = v2
                                                }
                                                if true {
                                                    var v158 uint16 = (v2 ^ v2)
                                                    _ = v158
                                                } else {
                                                    v2 = (v2 >> ((4294967295) & 63))
                                                }
                                                if (4294967295 > 658) {
                                                    var v159 int32 = -749
                                                    _ = v159
                                                } else {
                                                    delete(v3, genFn0(1))
                                                    var v160 uint64 = 165
                                                    v160 += v160
                                                    _ = v160
                                                }
                                            } else {
                                                fmt.Println(v2, fi152, fi154)
                                                fmt.Println(v2)
                                            }
                                        }
                                        delete(v3, 0)
                                        if ((false || false) && (65535 != 1)) {
                                            fmt.Println(fi152, v2)
                                            var v161 int32 = -1
                                            v162 := []int{128, int(v161), (^(v1)), v1}
                                            _ = v161
                                            _ = v162
                                        } else {
                                            v163 := []int{}
                                            if false {
                                                if len(v163) > 0 {
                                                    v163[0] = 490
                                                }
                                                var v164 uint16 = (v2 ^ 42)
                                                _ = v164
                                            }
                                            if len(v163) > 0 {
                                                v163[0] = -9223372036854775808
                                            }
                                            var v165 int32 = -323
                                            _ = v163
                                            _ = v165
                                        }
                                        _ = fi154
                                    }
                                    fmt.Println(v2, v1)
                                    _ = fi152
                                    _ = v153
                                }
                                v166 := v3[378]
                                v167 := []int{v1}
                                _ = v166
                                _ = v167
                            }
                            switch 0 {
                            case 220:
                                v168 := GenS11{F0: 42, F1: (("hello" + "") + ("" + "hello")), F2: int(v2)}
                                _ = v168
                            default:
                                fmt.Println(v2)
                            }
                        }
                        var v169 int16 = -1
                        if (false && false) {
                            for fi170 := 0; fi170 < 4; fi170++ {
                                v171 := GenS17{F0: "", F1: (3.14159 > float32((1.5 - 1e10))), F2: 442}
                                if true {
                                    for fi172 := 0; fi172 < 1; fi172++ {
                                        var v173 bool = true
                                        var v174 float64 = 1e10
                                        v175 := GenS1{F0: v173, F1: (204 == 1)}
                                        _ = fi172
                                        _ = v173
                                        _ = v174
                                        _ = v175
                                    }
                                    var v176 int32 = 42
                                    _ = v176
                                } else {
                                    var v177 uint = 863
                                    v177 = 178
                                    _ = v177
                                }
                                _ = fi170
                                _ = v171
                            }
                            switch v1 {
                            case -529:
                                for fi178 := 0; fi178 < 5; fi178++ {
                                    v169 = 1000
                                    _ = fi178
                                }
                            case 0:
                                if (82 >= 0) {
                                    for fi179 := 0; fi179 < 19; fi179++ {
                                        var v180 uint16 = v2
                                        var v181 int = 32768
                                        _ = fi179
                                        _ = v180
                                        _ = v181
                                    }
                                    fmt.Println(v1, v2, v1)
                                } else {
                                    var v182 float64 = float64(3.14159)
                                    v2 = (v2 >> ((1) & 63))
                                    _ = len(v3)
                                    _ = v182
                                }
                            case 100:
                                v169 = (v169 + v169)
                            }
                            v2 -= (v2 >> ((2) & 63))
                        } else {
                            v2 += (^(v2))
                        }
                        _ = v169
                    }
                    v183 := GenS18{F0: 891, F1: -38}
                    _ = v183
                }
            default:
                v2 += v2
            }
        }
    }
    var v184 uint8 = 171
    if (!true) {
        v185 := []int{(v1 / ((1) | 1)), v1, 4294967295}
        _ = v185
    }
    fmt.Println(v2, v2)
    v186 := GenS17{F0: ((("" + "") + (("" + "foo bar") + ("a" + "\n\t"))) + ""), F1: (2147483647 > 2147483647), F2: 116}
    for fi187 := 0; fi187 < 8; fi187++ {
        var v188 float32 = 1.5
        _ = fi187
        _ = v188
    }
    var v189 int8 = 106
    if true {
        var v190 uint8 = (v184 + (v184 - v184))
        _ = v190
    }
    switch ((("" + "") + "a") + "\n\t") {
    case "hello":
        v1 += -2147483648
    case "":
        v191 := GenS15{F0: 42}
        _ = v191
    case "a":
        v184 = (v184 | 128)
    default:
        if (-1 > -9223372036854775808) {
            fmt.Println(v184, v1)
            switch "foo bar" {
            case "a":
                var v192 uint8 = v184
                _ = v192
            case "hello":
                switch v1 {
                case 42:
                    v193 := v3[42]
                    _ = v193
                }
            }
        }
    }
    switch "hello" {
    case "":
        var v194 int16 = -1
        _ = v194
    case "\n\t":
        if (true && (!(18446744073709551615 != 4294967295))) {
            switch ("foo bar" + "") {
            case "foo bar":
                var v195 int16 = 523
                _ = v195
            }
            v1 = (v1 << ((2) & 63))
        } else {
            if (true && false) {
                var v196 int32 = 118
                if (-1 >= 0) {
                    v197 := GenS6{F0: 204, F1: v2, F2: v1}
                    for fi198 := 0; fi198 < 6; fi198++ {
                        v199 := GenS19{F0: "foo bar", F1: 1}
                        _ = fi198
                        _ = v199
                    }
                    _ = v197
                } else {
                    switch (-(v1)) {
                    case 32768:
                        v2 = (v2 + 816)
                    default:
                        _ = len(v3)
                    }
                }
                _ = v196
            }
        }
    default:
        var v200 uint32 = 1
        _ = v200
    }
    v201 := v3[1]
    v189 = (v189 >> ((4294967295) & 63))
    v201 = v201
    delete(v3, 344)
    switch int(v189) {
    case 0:
        var v202 int8 = v189
        _ = v202
    case -2147483648:
        v201 = v201
    case -1:
        v2 += v2
    default:
        v203 := []int{(v1 | v1), (v1 | (-(v1))), v1, 32768}
        _ = v203
    }
    v201 = (((v201 + ((v201 + v201) + v201)) + v201) + "")
    switch ("hello" + v201) {
    case "foo bar":
        for fi204 := 0; fi204 < 6; fi204++ {
            if true {
                var v205 int = int(v184)
                v189 += v189
                _ = v205
            } else {
                var v206 uint64 = 4294967295
                v207 := GenS6{F0: 65, F1: 714, F2: 0}
                _ = v206
                _ = v207
            }
            _ = fi204
        }
    case "\n\t":
        if true {
            v189 = (v189 << ((939) & 63))
            switch ((v201 + ("\n\t" + "a")) + v201) {
            case "foo bar":
                if false {
                    switch v1 {
                    case 42:
                        if ((1e10 >= -2.25) && (!(true && false))) {
                            if ((^(v184)) >= (v184 - (^(v184)))) {
                                fmt.Println(v201)
                            }
                            v208 := GenS20{F0: -128}
                            if (false && ((3.14159 * float64((((0.1 / 1e10) + 3.14159) - float32(1.5)))) >= 3.14159)) {
                                for fi209 := 0; fi209 < 8; fi209++ {
                                    if (422 > 18446744073709551615) {
                                        v1 = (v1 | 4294967295)
                                        fmt.Println(v2, v189)
                                        v189 = v189
                                    } else {
                                        for fi210 := 0; fi210 < 12; fi210++ {
                                            switch 861 {
                                            case 749:
                                                var v211 uint8 = v184
                                                _ = v211
                                            default:
                                                switch 32767 {
                                                case 32767:
                                                    switch (("hello" + v201) + (("a" + v201) + "\n\t")) {
                                                    case "\n\t":
                                                        v212 := []int{(fi209 & fi210)}
                                                        _ = v212
                                                    case "foo bar":
                                                        var v213 int32 = -1
                                                        if (834 != 0) {
                                                            v189 += v189
                                                        }
                                                        _ = v213
                                                    }
                                                case 128:
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v214 := []int{fi209}
                                                        _ = v214
                                                    case "":
                                                        v1 = 32767
                                                    case "a":
                                                        var v215 int16 = -1
                                                        _ = v215
                                                    default:
                                                        v216 := GenS20{F0: v189}
                                                        _ = v216
                                                    }
                                                default:
                                                    v189 += 0
                                                }
                                            }
                                            v217 := GenS21{F0: 1, F1: (-1 > -9223372036854775808)}
                                            v2 += (v2 << ((255) & 63))
                                            _ = fi210
                                            _ = v217
                                        }
                                        v2 = (v2 * v2)
                                        v218 := GenS22{F0: -2147483648, F1: v2, F2: 500}
                                        _ = v218
                                    }
                                    fmt.Println(v201, v201)
                                    _ = fi209
                                }
                                fmt.Println(v1, v189)
                                v219 := []int{v1, (v1 + -9223372036854775808), (v1 * (v1 + (-(v1)))), (v1 % (((^(v1))) | 1))}
                                _ = v219
                            }
                            _ = v208
                        }
                    default:
                        _ = len(v3)
                    }
                    var v220 uint32 = 65535
                    v221 := []int{(v1 | v1), v1, (v1 + (^(v1)))}
                    _ = v220
                    _ = v221
                } else {
                    for fi222 := 0; fi222 < 15; fi222++ {
                        if (!false) {
                            for fi223 := 0; fi223 < 13; fi223++ {
                                var v224 int8 = v189
                                _ = fi223
                                _ = v224
                            }
                            var v225 bool = ((true || ((1 > 42) && (genFn0(0) != 1))) || false)
                            if (65535 >= (v2 + 65535)) {
                                v226 := []int{fi222, -28}
                                if len(v226) > 0 {
                                    v226[0] = (-(v1))
                                }
                                v227 := GenS15{F0: -1}
                                _ = v226
                                _ = v227
                            }
                            _ = v225
                        }
                        _ = fi222
                    }
                }
            case "\n\t":
                v228 := GenS23{F0: 0, F1: false, F2: 18446744073709551615}
                _ = v228
            case "hello":
                v3[genFn0(1)] = v201
            }
            if (4294967295 != 42) {
                v189 += (-(v189))
                v229 := GenS6{F0: v184, F1: (^(v2)), F2: (v1 | (v1 * (v1 - -413)))}
                switch (v201 + (((v201 + v201) + v201) + v201)) {
                case "a":
                    var v230 string = "\n\t"
                    _ = v230
                case "foo bar":
                    v189 += v189
                }
                _ = v229
            } else {
                for fi231 := 0; fi231 < 9; fi231++ {
                    fmt.Println(fi231)
                    _ = fi231
                }
                if true {
                    v1 = v1
                }
                var v232 uint64 = 362
                _ = v232
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v184:", v184)
    fmt.Println("v189:", v189)
    fmt.Println("v2:", v2)
    fmt.Println("v201:", v201)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v184
    _ = v186
    _ = v189
    _ = v2
    _ = v201
    _ = v3
}
