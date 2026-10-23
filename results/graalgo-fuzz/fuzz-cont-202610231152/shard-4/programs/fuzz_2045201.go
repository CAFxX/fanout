package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint32; F1 int32; F2 int16 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint16; F1 int }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 int64; F1 int64 }
type GenS14 struct { F0 uint8; F1 int }
func genFn0(p0 string, p1 uint8) uint32 {
    v6 := []int{128, 2147483647, 195, -1}
    var v7 uint16 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 161
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    var v3 int64 = 666
    var v4 uint64 = 949
    var v5 int32 = 798
    if (255 != 32767) {
        v8 := GenS1{F0: 32767, F1: 536, F2: ("\n\t" + "foo bar")}
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(v5)
            var v10 uint32 = 2147483648
            _ = fi9
            _ = v10
        }
        switch -2147483648 {
        case 9223372036854775807:
            var v11 int = -2147483648
            _ = v11
        default:
            var v12 int32 = 2147483647
            _ = v12
        }
        _ = v8
    } else {
        v13 := GenS2{F0: (v4 | v4), F1: (v3 ^ (v3 | v3))}
        v5 -= 1
        _ = v13
    }
    var v14 uint = 0
    var v15 int16 = 42
    var v16 float32 = 3.14159
    var v17 int16 = v15
    if (uint16(v4) > 32767) {
        fmt.Println(v17, v16, v3)
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int = (fi18 | fi18)
            _ = fi18
            _ = v19
        }
        v20 := GenS2{F0: v4, F1: v3}
        _ = v20
    } else {
        var v21 int = int(v4)
        _ = v21
    }
    if ((!(true || (0.0 <= 0.1))) && true) {
        for fi22 := 0; fi22 < 5; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                fmt.Println(v4)
                v5 = (v5 % ((0) | 1))
                _ = fi23
            }
            _ = fi22
        }
    } else {
        v3 = v3
        switch 191 {
        case 2147483647:
            var v24 uint8 = 42
            _ = v24
        case 0:
            var v25 int32 = v5
            _ = v25
        case 128:
            fmt.Println(v14)
        }
    }
    if (6 <= 127) {
        v17 -= (^(v17))
    }
    v17 = (v15 & v17)
    var v26 uint32 = 2
    for fi27 := 0; fi27 < 19; fi27++ {
        v1[-702] = ""
        v28 := []int{(fi27 & -91), 9223372036854775807, 0, 9223372036854775807}
        _ = fi27
        _ = v28
    }
    if (true && false) {
        switch 4294967295 {
        case 32767:
            var v29 int64 = -9223372036854775808
            _ = v29
        case 0:
            switch 32768 {
            case 9223372036854775807:
                v15 = -1
            default:
                if true {
                    v30 := []int{4294967295, 32767, -9223372036854775808}
                    _ = v30
                } else {
                    v14 -= (v14 * (^(v14)))
                }
            }
        }
    } else {
        if ((!true) && true) {
            var v31 uint64 = v4
            _ = v31
        } else {
            switch 880 {
            case 2147483647:
                switch ("" + "a") {
                case "foo bar":
                    fmt.Println(v14)
                }
            case 32768:
                v32 := GenS3{F0: 612, F1: 1}
                _ = v32
            case 32767:
                v17 = v17
            }
        }
        v33 := GenS3{F0: 413, F1: (v3 << ((4294967295) & 63))}
        _ = v33
    }
    for fi34 := 0; fi34 < 9; fi34++ {
        var v35 int8 = 0
        var v36 int8 = (v35 | -1)
        _ = fi34
        _ = v35
        _ = v36
    }
    if false {
        if ((!true) || false) {
            if (false && true) {
                var v37 float32 = v16
                v2[(v26 - (v26 ^ v26))] = "a"
                v15 = v17
                _ = v37
            } else {
                var v38 int = -861
                delete(v2, (v26 >> (((^(v14))) & 63)))
                v16 += v16
                _ = v38
            }
        }
        fmt.Println(v14, v26)
        if (2147483647 >= 127) {
            var v39 int64 = (^(v3))
            _ = v39
        } else {
            if (-2.25 < 0.0) {
                v14 = 255
                for fi40 := 0; fi40 < 10; fi40++ {
                    switch (fi40 >> ((v14) & 63)) {
                    case 897:
                        var v41 int32 = v5
                        _ = v41
                    case 32767:
                        v14 = (v14 >> (((v14 - v14)) & 63))
                    case 125:
                        for fi42 := 0; fi42 < 3; fi42++ {
                            v4 = v4
                            _ = fi42
                        }
                    }
                    for fi43 := 0; fi43 < 10; fi43++ {
                        v44 := GenS2{F0: (v4 ^ 42), F1: 42}
                        _ = fi43
                        _ = v44
                    }
                    _ = fi40
                }
                for fi45 := 0; fi45 < 15; fi45++ {
                    v17 = (-(v15))
                    if (true || (42 != v5)) {
                        v16 += ((0.0 * -2.25) * (1.5 / v16))
                        fmt.Println(v15)
                        var v46 int8 = 42
                        _ = v46
                    }
                    _ = fi45
                }
            } else {
                switch "\n\t" {
                case "a":
                    var v47 uint32 = (^(v26))
                    _ = v47
                case "\n\t":
                    if true {
                        _ = len(v1)
                        if true {
                            v48 := []int{127, int(v14)}
                            _ = v48
                        } else {
                            var v49 int = 0
                            fmt.Println(v3, v14, v49)
                            _ = v49
                        }
                        v1[-1] = "hello"
                    } else {
                        var v50 uint8 = uint8(v26)
                        _ = v50
                    }
                default:
                    v14 -= 42
                }
                if false {
                    if (((!true) && (v4 > 4294967295)) && false) {
                        var v51 int16 = (v15 | 638)
                        var v52 int16 = v51
                        for fi53 := 0; fi53 < 16; fi53++ {
                            if true {
                                if false {
                                    v14 = v14
                                    if (false && true) {
                                        var v54 bool = true
                                        for fi55 := 0; fi55 < 12; fi55++ {
                                            if (127 == fi53) {
                                                v56 := v1[(-(fi55))]
                                                if (0.1 > (v16 / (v16 - 3.14159))) {
                                                    if v54 {
                                                        v57 := GenS3{F0: 1, F1: (-(v3))}
                                                        switch ("\n\t" + ((v56 + ((v56 + (v56 + "\n\t")) + ((v56 + v56) + v56))) + "foo bar")) {
                                                        case "foo bar":
                                                            var v58 uint16 = 32767
                                                            _ = v58
                                                        }
                                                        _ = v57
                                                    } else {
                                                        fmt.Println(v51)
                                                        if v54 {
                                                            v59 := v2[505]
                                                            _ = v59
                                                        }
                                                        var v60 uint64 = (v4 ^ v4)
                                                        _ = v60
                                                    }
                                                    _ = len(v2)
                                                    v16 += (v16 / ((v16 - v16) * 0.0))
                                                    v61 := GenS4{F0: (v26 * (v26 ^ uint32(v14))), F1: -1, F2: -1}
                                                    v2[(v26 & (v26 * v26))] = "hello"
                                                    var v62 bool = (9223372036854775807 != (v3 / (((v3 / (((^(v3))) | 1))) | 1)))
                                                    v26 = (v26 - (v26 & (v26 ^ genFn0("", 2))))
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    if (v54 || true) {
                                                        v52 -= v51
                                                    }
                                                }
                                                _ = v56
                                            }
                                            fmt.Println(v51)
                                            _ = fi55
                                        }
                                        v15 = (-(v51))
                                        _ = v54
                                    } else {
                                        for fi63 := 0; fi63 < 20; fi63++ {
                                            var v64 int32 = v5
                                            delete(v1, (fi53 << ((4294967295) & 63)))
                                            v3 = v3
                                            _ = fi63
                                            _ = v64
                                        }
                                        v65 := GenS5{F0: false, F1: (v3 << ((v14) & 63))}
                                        _ = v65
                                    }
                                }
                            } else {
                                v66 := GenS5{F0: true, F1: v3}
                                var v67 bool = true
                                for fi68 := 0; fi68 < 14; fi68++ {
                                    v52 = v17
                                    _ = fi68
                                }
                                _ = v66
                                _ = v67
                            }
                            _ = fi53
                        }
                        _ = v51
                        _ = v52
                    }
                    if ("a" >= "\n\t") {
                        if false {
                            var v69 bool = (((240 != v5) || ("" >= "")) && (true && false))
                            for fi70 := 0; fi70 < 15; fi70++ {
                                switch (fi70 * (fi70 ^ (fi70 >> ((uint(v4)) & 63)))) {
                                case 134:
                                    v4 -= 4294967295
                                }
                                if true {
                                    if (v14 > v14) {
                                        v15 = 0
                                        var v71 int = 158
                                        _ = v71
                                    }
                                    if v69 {
                                        v3 = (v3 + v3)
                                        fmt.Println(v5, v4)
                                    } else {
                                        if (fi70 < fi70) {
                                            v5 = (v5 ^ v5)
                                            v72 := GenS6{F0: 42}
                                            for fi73 := 0; fi73 < 3; fi73++ {
                                                v5 += v5
                                                _ = fi73
                                            }
                                            _ = v72
                                        } else {
                                            fmt.Println(v5)
                                            v26 += genFn0("", 105)
                                        }
                                        fmt.Println(fi70)
                                        if v69 {
                                            fmt.Println(v5)
                                            var v74 int64 = -213
                                            v75 := GenS7{F0: 518, F1: 127}
                                            _ = v74
                                            _ = v75
                                        } else {
                                            v76 := GenS5{F0: (v69 && v69), F1: (v3 << ((v14) & 63))}
                                            var v77 uint64 = v4
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                    var v78 int16 = int16(v14)
                                    _ = v78
                                } else {
                                    var v79 uint16 = 0
                                    v3 = v3
                                    _ = v79
                                }
                                _ = fi70
                            }
                            switch -564 {
                            case -2147483648:
                                var v80 int32 = v5
                                _ = v80
                            case 32767:
                                v81 := []int{453, 32768}
                                _ = v81
                            case 476:
                                v3 += v3
                            default:
                                v14 += (v14 % (((v14 >> (((v14 & 255)) & 63))) | 1))
                            }
                            _ = v69
                        }
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v17)
                        case "foo bar":
                            if true {
                                if false {
                                    var v82 uint64 = (v4 % ((v4) | 1))
                                    if (v17 <= -32768) {
                                        for fi83 := 0; fi83 < 12; fi83++ {
                                            var v84 int32 = (^(v5))
                                            var v85 float32 = 1e10
                                            _ = fi83
                                            _ = v84
                                            _ = v85
                                        }
                                    } else {
                                        for fi86 := 0; fi86 < 1; fi86++ {
                                            v4 += (v82 % ((5) | 1))
                                            var v87 bool = true
                                            switch "" {
                                            case "a":
                                                if v87 {
                                                    if v87 {
                                                        v88 := GenS8{F0: 448, F1: 1}
                                                        _ = v88
                                                    } else {
                                                        v89 := GenS9{F0: 376, F1: 255}
                                                        _ = v89
                                                    }
                                                    switch "" {
                                                    case "":
                                                        switch fi86 {
                                                        case 4294967295:
                                                            fmt.Println(v15, v16)
                                                        case 2147483647:
                                                            var v90 uint32 = v26
                                                            _ = v90
                                                        case 128:
                                                            if false {
                                                                for fi91 := 0; fi91 < 13; fi91++ {
                                                                    v16 -= 3.14159
                                                                    v92 := GenS10{F0: v15}
                                                                    var v93 int64 = v3
                                                                    var v94 bool = false
                                                                    _ = fi91
                                                                    _ = v92
                                                                    _ = v93
                                                                    _ = v94
                                                                }
                                                                var v95 float32 = v16
                                                                _ = v95
                                                            } else {
                                                                for fi96 := 0; fi96 < 6; fi96++ {
                                                                    var v97 uint64 = v4
                                                                    v98 := v2[65535]
                                                                    fmt.Println(fi86, v87, v87)
                                                                    v14 += (v14 << (((v14 - v14)) & 63))
                                                                    _ = fi96
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                                var v99 uint32 = 65535
                                                                _ = v99
                                                            }
                                                        }
                                                    case "\n\t":
                                                        if (v3 < v3) {
                                                            _ = len(v2)
                                                            for fi100 := 0; fi100 < 20; fi100++ {
                                                                v101 := []int{(fi100 & fi100), (fi86 >> ((2) & 63)), (-(fi100)), fi86}
                                                                fmt.Println(v15)
                                                                var v102 int8 = -128
                                                                _ = fi100
                                                                _ = v101
                                                                _ = v102
                                                            }
                                                            var v103 string = ""
                                                            _ = v103
                                                        } else {
                                                            for fi104 := 0; fi104 < 6; fi104++ {
                                                                v82 = v82
                                                                var v105 uint32 = v26
                                                                _ = fi104
                                                                _ = v105
                                                            }
                                                            switch ("a" + ("foo bar" + "")) {
                                                            case "a":
                                                                v106 := GenS10{F0: v17}
                                                                _ = v106
                                                            }
                                                            switch (((("hello" + "foo bar") + ("\n\t" + "\n\t")) + "") + ("a" + "hello")) {
                                                            case "foo bar":
                                                                var v107 int64 = -1
                                                                _ = v107
                                                            case "hello":
                                                                v108 := []int{}
                                                                _ = v108
                                                            case "\n\t":
                                                                v109 := []int{}
                                                                _ = v109
                                                            }
                                                        }
                                                    case "hello":
                                                        v17 += v15
                                                    }
                                                    switch "hello" {
                                                    case "a":
                                                        var v110 uint8 = 128
                                                        _ = v110
                                                    case "\n\t":
                                                        v17 = 32767
                                                    case "foo bar":
                                                        for fi111 := 0; fi111 < 7; fi111++ {
                                                            fmt.Println(v26)
                                                            _ = fi111
                                                        }
                                                    }
                                                }
                                            default:
                                                v87 = v87
                                            }
                                            _ = fi86
                                            _ = v87
                                        }
                                        var v112 uint8 = uint8(v4)
                                        v113 := GenS1{F0: 314, F1: 140, F2: "\n\t"}
                                        _ = v112
                                        _ = v113
                                    }
                                    v114 := GenS6{F0: 0}
                                    _ = v114
                                    _ = v82
                                }
                            } else {
                                v115 := GenS11{F0: 42, F1: 587}
                                var v116 uint16 = 1
                                var v117 string = ("" + "hello")
                                _ = v115
                                _ = v116
                                _ = v117
                            }
                        case "hello":
                            v118 := GenS12{F0: -800, F1: 42}
                            _ = v118
                        default:
                            v14 = v14
                        }
                        v26 = (v26 << ((v14) & 63))
                    }
                    if true {
                        v119 := GenS11{F0: 2, F1: 127}
                        if false {
                            v120 := GenS13{F0: (-(v3)), F1: v3}
                            v16 += 1e10
                            v26 -= (v26 * (v26 & 253))
                            _ = v120
                        }
                        if (647 == (v26 + genFn0((("\n\t" + ("\n\t" + "\n\t")) + "\n\t"), 0))) {
                            var v121 int64 = 198
                            var v122 bool = true
                            _ = v121
                            _ = v122
                        } else {
                            fmt.Println(v15)
                            fmt.Println(v14, v26)
                            v15 -= 42
                        }
                        _ = v119
                    } else {
                        v14 -= 304
                        if ((v5 <= -1) && (true || true)) {
                            fmt.Println(v5, v3)
                            v5 = v5
                            for fi123 := 0; fi123 < 16; fi123++ {
                                v124 := []int{}
                                _ = fi123
                                _ = v124
                            }
                        } else {
                            var v125 int64 = v3
                            v14 -= (v14 & v14)
                            v126 := GenS14{F0: 128, F1: -398}
                            _ = v125
                            _ = v126
                        }
                    }
                }
                switch "foo bar" {
                case "hello":
                    v127 := v2[genFn0(("hello" + ("" + "")), 1)]
                    _ = v127
                case "\n\t":
                    if true {
                        v128 := GenS6{F0: -1}
                        switch "foo bar" {
                        case "hello":
                            v26 -= v26
                        }
                        v129 := GenS2{F0: v4, F1: v3}
                        _ = v128
                        _ = v129
                    } else {
                        var v130 int = 128
                        var v131 int32 = (v5 + (v5 % ((int32(v26)) | 1)))
                        _ = v130
                        _ = v131
                    }
                case "":
                    v4 = v4
                }
            }
        }
    } else {
        v2[65535] = "hello"
        fmt.Println(v3, v16, v4)
    }
    var v132 float32 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v132:", v132)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v132
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
