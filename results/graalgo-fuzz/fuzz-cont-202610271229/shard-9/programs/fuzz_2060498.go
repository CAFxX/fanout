package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int16; F2 uint }
type GenS5 struct { F0 int8; F1 string; F2 uint }
type GenS6 struct { F0 int; F1 int16; F2 uint }
type GenS7 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int32; F1 int; F2 uint64 }
type GenS12 struct { F0 string; F1 int }
type GenS13 struct { F0 string; F1 string }
type GenS14 struct { F0 string; F1 string }
type GenS15 struct { F0 int64; F1 int16 }
type GenS16 struct { F0 int8; F1 bool; F2 bool }
type GenS17 struct { F0 bool; F1 uint8; F2 int }
type GenS18 struct { F0 uint16; F1 uint32; F2 uint64 }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint16 = 1
    v3 := make(map[uint32]int)
    var v4 int16 = -1
    var v5 uint8 = 40
    var v6 bool = false
    if (0 >= 42) {
        if v6 {
            if v6 {
                var v7 int8 = 100
                if (!(int32(v4) >= -500)) {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        var v9 uint64 = 0
                        v1 = 459
                        _ = fi8
                        _ = v9
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        delete(v3, 1)
                    case "\n\t":
                        if v6 {
                            v2 = v2
                            fmt.Println(v1, v6, v7)
                            var v10 float64 = 0.1
                            _ = v10
                        } else {
                            fmt.Println(v4, v2)
                            var v11 string = "\n\t"
                            _ = v11
                        }
                    case "foo bar":
                        if (32768 < 512) {
                            var v12 string = "a"
                            v7 -= (v7 ^ (^(v7)))
                            _ = v12
                        } else {
                            v2 += 0
                        }
                    default:
                        var v13 int32 = 0
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v6 = v6
                        v7 -= -128
                        var v15 int = int(v1)
                        _ = fi14
                        _ = v15
                    }
                    if (true || true) {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v17 := GenS1{F0: 936}
                            _ = fi16
                            _ = v17
                        }
                        if true {
                            var v18 int = 32768
                            v4 -= (v4 ^ v4)
                            _ = v18
                        } else {
                            fmt.Println(v6, v2, v5)
                        }
                        if (2 >= 469) {
                            v19 := GenS2{F0: ("\n\t" + "a"), F1: 42, F2: (((("hello" + "\n\t") + "\n\t") + "\n\t") + "foo bar")}
                            _ = v19
                        } else {
                            v5 = 42
                        }
                    }
                }
                fmt.Println(v6)
                _ = v7
            }
        }
    }
    if v6 {
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 4; fi21++ {
                switch ("a" + ("" + ("hello" + "a"))) {
                case "\n\t":
                    v22 := GenS3{F0: (v1 >> ((65535) & 63))}
                    _ = v22
                case "":
                    var v23 int8 = -1
                    _ = v23
                }
                _ = fi21
            }
            var v24 bool = (v4 == (v4 << ((844) & 63)))
            v1 = 2147483648
            _ = fi20
            _ = v24
        }
        v1 = (^(v1))
    }
    v25 := GenS4{F0: 42, F1: (v4 % (((v4 - v4)) | 1)), F2: 2}
    var v26 int16 = 912
    v4 = v26
    v2 = v2
    v27 := []int{}
    var v28 string = "foo bar"
    v1 -= 42
    for fi29 := 0; fi29 < 5; fi29++ {
        if v6 {
            if false {
                var v30 int32 = -1
                fmt.Println(v28, v6)
                v31 := GenS2{F0: v28, F1: 18446744073709551615, F2: "hello"}
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 uint32 = 411
                if (v5 != 42) {
                    _ = len(v27)
                } else {
                    switch "foo bar" {
                    case "":
                        if ((1.5 * 0.1) < 0.1) {
                            switch v28 {
                            case "foo bar":
                                if (float32(0.1) != ((3.14159 / float32(1.5)) + (1e10 + 1e10))) {
                                    var v34 int8 = -1
                                    var v35 int16 = v4
                                    _ = v34
                                    _ = v35
                                } else {
                                    v36 := v3[(^(v33))]
                                    v37 := GenS5{F0: int8(v2), F1: ("hello" + ((v28 + v28) + (v28 + ("" + "a")))), F2: 551}
                                    if len(v27) > 0 {
                                        v27[0] = (fi29 >> ((0) & 63))
                                    }
                                    _ = v36
                                    _ = v37
                                }
                            }
                            v3[v1] = (^(fi32))
                            v26 += v4
                            v38 := GenS6{F0: fi32, F1: -32768, F2: 4294967295}
                            if (v6 || (42 >= int8(v5))) {
                                if (((1e10 - 3.14159) / 1e10) <= -2.25) {
                                    v1 = v1
                                    var v39 int8 = int8(v2)
                                    _ = v39
                                }
                                v6 = (float32((1e10 + 3.14159)) > 1e10)
                                if true {
                                    if true {
                                        _ = len(v3)
                                    }
                                }
                                v40 := GenS7{F0: (^(v33)), F1: 100, F2: 161}
                                v41 := GenS4{F0: 836, F1: (v4 | (^(v4))), F2: 0}
                                var v42 uint16 = v2
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            _ = v38
                        } else {
                            var v43 uint = 18446744073709551615
                            switch v28 {
                            case "foo bar":
                                _ = len(v3)
                            default:
                                if (v5 != 127) {
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    v28 = ""
                                    var v44 float64 = (3.14159 * 1e10)
                                    v5 += (v5 >> ((2) & 63))
                                    _ = v44
                                } else {
                                    if false {
                                        var v45 uint64 = 42
                                        fmt.Println(v4)
                                        if (42 == 42) {
                                            v26 -= (v26 / ((v26) | 1))
                                            switch (v28 + (v28 + "")) {
                                            case "a":
                                                switch 128 {
                                                case 128:
                                                    fmt.Println(v33)
                                                    var v46 float32 = 0.0
                                                    _ = v46
                                                case 529:
                                                    delete(v3, (v1 % ((4294967295) | 1)))
                                                case 0:
                                                    var v47 uint32 = v1
                                                    _ = v47
                                                }
                                            default:
                                                if (-709 != 42) {
                                                    switch v28 {
                                                    case "foo bar":
                                                        switch ("hello" + "foo bar") {
                                                        case "foo bar":
                                                            switch ("foo bar" + v28) {
                                                            case "\n\t":
                                                                var v48 float32 = 1.5
                                                                _ = v48
                                                            case "":
                                                                v33 = 2
                                                            default:
                                                                var v49 float64 = 1.5
                                                                _ = v49
                                                            }
                                                        default:
                                                            var v50 int16 = v4
                                                            _ = v50
                                                        }
                                                    case "a":
                                                        v43 += 255
                                                    case "hello":
                                                        fmt.Println(v1, v33)
                                                    default:
                                                        if (554 >= (v2 | (v2 + v2))) {
                                                            _ = len(v3)
                                                        } else {
                                                            var v51 uint8 = (v5 | v5)
                                                            v28 = v28
                                                            switch fi32 {
                                                            case 169:
                                                                switch (fi32 + fi32) {
                                                                case 32767:
                                                                    v52 := v3[v1]
                                                                    _ = v52
                                                                case 1:
                                                                    v6 = v6
                                                                case -451:
                                                                    fmt.Println(v6)
                                                                }
                                                                if v6 {
                                                                    v53 := GenS8{F0: -46, F1: (v43 ^ (v43 / (((v43 << ((698) & 63))) | 1)))}
                                                                    v54 := GenS3{F0: v1}
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                var v55 int16 = v4
                                                                var v56 float32 = (((float32(1.5) * (3.14159 * 1.5)) - 3.14159) * 1e10)
                                                                _ = v55
                                                                _ = v56
                                                            case 1:
                                                                var v57 uint = (^(v43))
                                                                _ = v57
                                                            }
                                                            _ = v51
                                                        }
                                                    }
                                                }
                                                if len(v27) >= 2 {
                                                    v27 = v27[1:len(v27)]
                                                }
                                            }
                                            var v58 uint16 = v2
                                            _ = v58
                                        } else {
                                            v1 = (^(v33))
                                        }
                                        _ = v45
                                    } else {
                                        fmt.Println(v5)
                                        fmt.Println(v28, v2)
                                        fmt.Println(v6)
                                    }
                                    v1 = v33
                                    var v59 bool = v6
                                    v60 := GenS2{F0: v28, F1: v43, F2: "hello"}
                                    _ = v59
                                    _ = v60
                                }
                            }
                            if true {
                                v26 = v4
                                v2 += 1
                            }
                            _ = v43
                        }
                    default:
                        if v6 {
                            v4 = v26
                            if (50 == v5) {
                                fmt.Println(v28, v33, v33)
                                v1 += 54
                                var v61 bool = v6
                                v62 := GenS9{F0: 195}
                                v63 := GenS8{F0: 1, F1: 65535}
                                _ = v61
                                _ = v62
                                _ = v63
                            } else {
                                var v64 uint32 = 0
                                _ = v64
                            }
                            v6 = v6
                        } else {
                            fmt.Println(v5, fi29)
                            v2 = (^(v2))
                            var v65 int32 = 1
                            _ = v65
                        }
                    }
                }
                var v66 int8 = 127
                switch v28 {
                case "\n\t":
                    switch (fi29 - (fi32 >> ((1) & 63))) {
                    case -1:
                        v4 -= 0
                    case 127:
                        var v67 uint = 2
                        _ = v67
                    }
                    var v68 int32 = -1
                    v28 = (v28 + (v28 + v28))
                    _ = v68
                case "hello":
                    v2 = v2
                default:
                    v1 += v1
                }
                var v69 string = (v28 + (((v28 + (v28 + "")) + "foo bar") + (((v28 + v28) + v28) + ("\n\t" + ("\n\t" + v28)))))
                if (!(!true)) {
                    _ = len(v27)
                    switch 4294967295 {
                    case 2147483647:
                        var v70 float64 = (1.5 / 1.0)
                        _ = v70
                    case 127:
                        var v71 uint = 705
                        _ = v71
                    case 4294967295:
                        v72 := GenS1{F0: uint(v2)}
                        _ = v72
                    default:
                        if v6 {
                            var v73 string = v69
                            v27 = append(v27, int(v1))
                            if (-2.25 <= 0.0) {
                                var v74 float32 = (float32(((float64(3.14159) + (1.5 + 1.5)) - 0.1)) / ((0.0 + 3.14159) / (1e10 * float32((float64(-2.25) + float64(0.0))))))
                                if ((^(v1)) <= (v1 / ((4294967295) | 1))) {
                                    var v75 float64 = (float64(0.0) / 1e10)
                                    _ = v75
                                }
                                _ = v74
                            } else {
                                if v6 {
                                    v1 = v1
                                    delete(v3, uint32(v2))
                                    v76 := GenS8{F0: -2147483648, F1: 2}
                                    _ = v76
                                } else {
                                    v27 = append(v27, (fi32 | (fi29 / ((fi29) | 1))))
                                    if (true || v6) {
                                        var v77 string = ((((v28 + "") + (v28 + (((v69 + v69) + ("hello" + "")) + v28))) + v28) + "hello")
                                        v78 := GenS1{F0: uint(v2)}
                                        if v6 {
                                            if false {
                                                v77 = "\n\t"
                                                fmt.Println(v77)
                                            } else {
                                                var v79 string = ("foo bar" + v73)
                                                var v80 uint8 = (v5 & v5)
                                                v81 := GenS6{F0: -1, F1: v26, F2: uint(v5)}
                                                _ = v79
                                                _ = v80
                                                _ = v81
                                            }
                                            if v6 {
                                                v82 := GenS10{F0: 127}
                                                var v83 uint = 0
                                                v33 += 2
                                                _ = v82
                                                _ = v83
                                            }
                                        }
                                        _ = v77
                                        _ = v78
                                    }
                                    var v84 uint64 = 242
                                    _ = v84
                                }
                            }
                            _ = v73
                        } else {
                            var v85 int32 = 2147483647
                            _ = v85
                        }
                    }
                } else {
                    v26 += v4
                    if v6 {
                        v1 -= 1
                        v86 := GenS6{F0: 98, F1: (v4 >> ((uint(v5)) & 63)), F2: 613}
                        var v87 int64 = 0
                        _ = v86
                        _ = v87
                    }
                }
                var v88 uint32 = (v1 * v1)
                _ = fi32
                _ = v33
                _ = v66
                _ = v69
                _ = v88
            }
        }
        v89 := v3[(v1 | 996)]
        v5 += (v5 >> ((255) & 63))
        _ = fi29
        _ = v89
    }
    if (!v6) {
        _ = len(v3)
        var v90 string = (v28 + "\n\t")
        v5 = (v5 / (((^(v5))) | 1))
        _ = v90
    }
    if true {
        v91 := GenS4{F0: -127, F1: 1000, F2: 2}
        for fi92 := 0; fi92 < 8; fi92++ {
            var v93 int16 = (-(v4))
            _ = fi92
            _ = v93
        }
        v26 += (-(v4))
        _ = v91
    } else {
        switch -9223372036854775808 {
        case 32768:
            for fi94 := 0; fi94 < 16; fi94++ {
                if v6 {
                    fmt.Println(v2, v28, v4)
                    switch 32768 {
                    case 940:
                        v26 = v4
                    case -493:
                        _ = len(v27)
                    case -984:
                        _ = len(v27)
                    }
                    var v95 int8 = 44
                    _ = v95
                } else {
                    fmt.Println(v28, v4)
                    if (-1 >= 814) {
                        switch "foo bar" {
                        case "a":
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                        default:
                            for fi96 := 0; fi96 < 4; fi96++ {
                                var v97 bool = (9223372036854775807 > -1)
                                v3[v1] = 128
                                switch (fi94 % ((fi96) | 1)) {
                                case -800:
                                    switch "" {
                                    case "a":
                                        if (true && true) {
                                            fmt.Println(v28, v1, fi94)
                                            v97 = v6
                                            var v98 int16 = 0
                                            _ = v98
                                        } else {
                                            v99 := GenS5{F0: -1, F1: v28, F2: 255}
                                            _ = v99
                                        }
                                    case "foo bar":
                                        if v6 {
                                            var v100 uint16 = (v2 << ((0) & 63))
                                            v101 := GenS6{F0: 128, F1: (v26 << ((775) & 63)), F2: 0}
                                            _ = v100
                                            _ = v101
                                        }
                                    case "":
                                        switch "foo bar" {
                                        case "hello":
                                            if v97 {
                                                var v102 string = "hello"
                                                var v103 float32 = 0.0
                                                var v104 float32 = v103
                                                _ = v102
                                                _ = v103
                                                _ = v104
                                            }
                                        case "":
                                            var v105 int = fi94
                                            _ = v105
                                        default:
                                            v1 += 2147483648
                                        }
                                    }
                                case 847:
                                    var v106 int64 = 42
                                    _ = v106
                                default:
                                    v6 = (false || (v2 == v2))
                                }
                                _ = fi96
                                _ = v97
                            }
                        }
                        if (-457 <= -1) {
                            for fi107 := 0; fi107 < 15; fi107++ {
                                switch (fi94 - 1) {
                                case 0:
                                    v1 += v1
                                default:
                                    if true {
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(v1, fi94)
                                        case "\n\t":
                                            if (v6 || (v2 > 1)) {
                                                fmt.Println(v28, v5)
                                                switch (fi107 ^ fi94) {
                                                case 42:
                                                    v28 = v28
                                                }
                                                v28 = (v28 + (("hello" + (v28 + "")) + ((v28 + "hello") + "a")))
                                            }
                                        default:
                                            var v108 int = 1
                                            var v109 float32 = 3.14159
                                            _ = v108
                                            _ = v109
                                        }
                                        fmt.Println(fi94, fi107, v4)
                                        if v6 {
                                            if v6 {
                                                switch (fi107 - (fi107 ^ fi94)) {
                                                case 1:
                                                    if v6 {
                                                        switch (v28 + ("" + v28)) {
                                                        case "":
                                                            v110 := GenS8{F0: 0, F1: 2}
                                                            _ = v110
                                                        case "foo bar":
                                                            switch v28 {
                                                            case "foo bar":
                                                                fmt.Println(v2)
                                                            case "":
                                                                fmt.Println(v6, v28, v26)
                                                            }
                                                        case "a":
                                                            v27 = append(v27, (fi94 >> ((65535) & 63)))
                                                        default:
                                                            v28 = v28
                                                        }
                                                        var v111 int32 = 0
                                                        v1 -= v1
                                                        _ = v111
                                                    } else {
                                                        v26 -= 1000
                                                        if false {
                                                            v3[65535] = 0
                                                            if v6 {
                                                                v112 := GenS9{F0: 790}
                                                                _ = v112
                                                            } else {
                                                                var v113 int32 = 42
                                                                _ = v113
                                                            }
                                                        }
                                                        var v114 uint32 = v1
                                                        var v115 int64 = 666
                                                        var v116 int16 = v26
                                                        var v117 int = (fi94 | (fi107 & (fi94 + (fi94 * 2147483647))))
                                                        v114 += (v1 + (v1 & 2147483648))
                                                        v118 := GenS2{F0: "hello", F1: 0, F2: ""}
                                                        var v119 string = "\n\t"
                                                        var v120 int32 = 145
                                                        if (-68 < int8(v2)) {
                                                            v121 := GenS2{F0: (((v119 + v28) + (v28 + ("a" + ((v28 + ((v119 + (v119 + v119)) + ((v119 + ("" + v119)) + v28))) + "\n\t")))) + v119), F1: 487, F2: v119}
                                                            v122 := GenS10{F0: 42}
                                                            v115 = v115
                                                            _ = v121
                                                            _ = v122
                                                        }
                                                        v123 := GenS11{F0: v120, F1: (fi107 - 4294967295), F2: 0}
                                                        _ = v114
                                                        _ = v115
                                                        _ = v116
                                                        _ = v117
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                        _ = v123
                                                    }
                                                case 32768:
                                                    v124 := GenS12{F0: (v28 + v28), F1: (fi107 & (fi94 & (^(fi107))))}
                                                    _ = v124
                                                case 42:
                                                    switch "a" {
                                                    case "foo bar":
                                                        v2 -= (v2 + v2)
                                                    case "hello":
                                                        v125 := GenS13{F0: ("hello" + "a"), F1: v28}
                                                        _ = v125
                                                    case "a":
                                                        if ((0 >= (v1 + v1)) && v6) {
                                                            switch v28 {
                                                            case "\n\t":
                                                                v26 = v4
                                                            case "":
                                                                if len(v27) >= 2 {
                                                                    v27 = v27[1:len(v27)]
                                                                }
                                                            }
                                                            v126 := GenS4{F0: 1, F1: v26, F2: 1}
                                                            fmt.Println(fi107)
                                                            _ = v126
                                                        } else {
                                                            if true {
                                                                fmt.Println(v26)
                                                                _ = len(v3)
                                                                v127 := GenS14{F0: v28, F1: ((("\n\t" + ("" + "\n\t")) + v28) + "hello")}
                                                                _ = v127
                                                            }
                                                            if (-871 == 42) {
                                                                v6 = v6
                                                                v6 = (6 < 2)
                                                            } else {
                                                                if len(v27) > 0 {
                                                                    v27[0] = fi107
                                                                }
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v128 := GenS8{F0: int32(v1), F1: 615}
                                                    _ = v128
                                                }
                                                var v129 int8 = int8(fi107)
                                                _ = v129
                                            } else {
                                                v27 = append(v27, fi94)
                                                fmt.Println(v2)
                                                v130 := GenS6{F0: (fi94 % ((fi107) | 1)), F1: (v26 / ((32767) | 1)), F2: 146}
                                                _ = v130
                                            }
                                            if v6 {
                                                switch v28 {
                                                case "foo bar":
                                                    v131 := GenS15{F0: 742, F1: -798}
                                                    _ = v131
                                                case "hello":
                                                    if v6 {
                                                        fmt.Println(v26, v26, v26)
                                                        v5 = (v5 >> ((18446744073709551615) & 63))
                                                        var v132 string = v28
                                                        _ = v132
                                                    }
                                                case "\n\t":
                                                    switch (fi107 % ((1) | 1)) {
                                                    case 520:
                                                        v28 = v28
                                                    case 854:
                                                        v5 = (^(v5))
                                                    default:
                                                        v133 := GenS6{F0: fi94, F1: -32768, F2: 42}
                                                        _ = v133
                                                    }
                                                default:
                                                    v26 = (v4 * (v26 ^ (-(v26))))
                                                }
                                                var v134 int32 = 594
                                                var v135 uint16 = (^(v2))
                                                _ = v134
                                                _ = v135
                                            } else {
                                                v3[v1] = (fi107 & 772)
                                                var v136 bool = ((2147483647 != 42) || v6)
                                                _ = v136
                                            }
                                        }
                                    }
                                }
                                v26 -= v4
                                _ = fi107
                            }
                            var v137 float32 = 0.0
                            _ = v137
                        } else {
                            switch "a" {
                            case "foo bar":
                                for fi138 := 0; fi138 < 19; fi138++ {
                                    v4 = v4
                                    _ = fi138
                                }
                            case "a":
                                if (v4 != -791) {
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                } else {
                                    if (!v6) {
                                        for fi139 := 0; fi139 < 4; fi139++ {
                                            var v140 string = "a"
                                            var v141 uint8 = 255
                                            var v142 int64 = 1
                                            _ = fi139
                                            _ = v140
                                            _ = v141
                                            _ = v142
                                        }
                                    }
                                    if (9223372036854775807 >= -118) {
                                        var v143 uint16 = (v2 + 255)
                                        _ = v143
                                    } else {
                                        v4 += (v4 * 42)
                                    }
                                }
                            case "":
                                var v144 uint8 = 42
                                _ = v144
                            }
                            v6 = (1 == 1)
                        }
                    } else {
                        var v145 int = fi94
                        v146 := GenS16{F0: 0, F1: v6, F2: (!(v1 < (v1 >> ((646) & 63))))}
                        _ = v145
                        _ = v146
                    }
                    if (!false) {
                        v28 = ""
                        v1 -= v1
                    }
                }
                switch fi94 {
                case 42:
                    for fi147 := 0; fi147 < 11; fi147++ {
                        fmt.Println(v4, v26, v1)
                        if (v5 <= v5) {
                            v148 := GenS17{F0: true, F1: (v5 >> ((255) & 63)), F2: fi94}
                            _ = v148
                        }
                        _ = fi147
                    }
                }
                v4 = (-(v4))
                _ = fi94
            }
        default:
            switch 127 {
            case -36:
                for fi149 := 0; fi149 < 15; fi149++ {
                    _ = len(v27)
                    _ = fi149
                }
            case 876:
                var v150 int8 = 127
                _ = v150
            }
        }
        if (!v6) {
            var v151 string = (((v28 + (v28 + v28)) + (v28 + (v28 + "a"))) + (("foo bar" + "a") + v28))
            _ = v151
        } else {
            switch "\n\t" {
            case "hello":
                if v6 {
                    for fi152 := 0; fi152 < 2; fi152++ {
                        v4 = v4
                        _ = fi152
                    }
                    fmt.Println(v2)
                    var v153 bool = v6
                    _ = v153
                }
            case "":
                v154 := GenS17{F0: v6, F1: 1, F2: -9223372036854775808}
                _ = v154
            case "\n\t":
                v28 = ("hello" + v28)
            }
            fmt.Println(v28)
            switch 127 {
            case 4294967295:
                v6 = false
            }
        }
        switch 0 {
        case 127:
            v6 = true
        default:
            v155 := GenS18{F0: uint16(v1), F1: v1, F2: 698}
            _ = v155
        }
    }
    var v156 int32 = -1
    v157 := GenS16{F0: 127, F1: (v156 >= (v156 >> ((1) & 63))), F2: (!(42 == 65535))}
    fmt.Println("v1:", v1)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v156
    _ = v157
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
