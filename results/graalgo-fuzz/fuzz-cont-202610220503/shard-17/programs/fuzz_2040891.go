package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int64; F1 int8; F2 uint }
type GenS7 struct { F0 uint8; F1 bool; F2 int }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint16; F1 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32; F1 int32; F2 int8 }
type GenS12 struct { F0 int64; F1 int; F2 string }
type GenS13 struct { F0 int16; F1 int64; F2 uint64 }
type GenS14 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS15 struct { F0 int; F1 int; F2 bool }
func genFn0(p0 uint16, p1 uint8) int8 {
    var v4 string = ("a" + ("foo bar" + "\n\t"))
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []uint32{1, 65535, 42}
    v2 := make(map[int]bool)
    var v3 bool = false
    if (float64((0.1 * 3.14159)) <= ((-2.25 / -2.25) + 0.1)) {
        for fi6 := 0; fi6 < 4; fi6++ {
            switch ("a" + ((((((("foo bar" + "foo bar") + "\n\t") + "hello") + ("a" + "a")) + "") + "a") + ("" + "a"))) {
            case "a":
                v7 := GenS1{F0: "", F1: 42, F2: 65535}
                _ = v7
            default:
                v3 = ((234 != 462) || ((((genFn0(1, 108) == -128) && (!(!v3))) && ((!v3) && true)) && (2147483647 <= 2147483647)))
            }
            _ = fi6
        }
        var v8 uint = 65535
        var v9 uint32 = 4294967295
        _ = v8
        _ = v9
    } else {
        v3 = (127 <= genFn0(1, 128))
        var v10 float32 = 1e10
        _ = v10
    }
    if v3 {
        if v3 {
            var v11 int = -2147483648
            _ = v11
        } else {
            var v12 float32 = (-2.25 * 1.5)
            var v13 int = 42
            _ = v12
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int64 = 1
            var v16 bool = v3
            _ = fi14
            _ = v15
            _ = v16
        }
        if v3 {
            v2[127] = v3
            if v3 {
                var v17 uint32 = 42
                var v18 uint16 = 42
                if v3 {
                    switch ("foo bar" + "a") {
                    case "\n\t":
                        var v19 uint64 = 0
                        _ = v19
                    default:
                        if v3 {
                            var v20 uint16 = (v18 & v18)
                            _ = v20
                        }
                    }
                } else {
                    fmt.Println(v3, v18, v3)
                    var v21 uint16 = 65535
                    _ = v21
                }
                _ = v17
                _ = v18
            } else {
                if false {
                    if v3 {
                        var v22 int8 = -1
                        _ = v22
                    }
                } else {
                    v23 := GenS2{F0: 32767}
                    v3 = ("hello" <= "\n\t")
                    _ = v23
                }
            }
            switch 2147483647 {
            case 32768:
                var v24 int8 = -128
                _ = v24
            case 191:
                if (true || (1000 < 104)) {
                    for fi25 := 0; fi25 < 13; fi25++ {
                        if (v3 || (81 > 42)) {
                            var v26 uint16 = 694
                            for fi27 := 0; fi27 < 19; fi27++ {
                                v28 := GenS3{F0: v3, F1: 807, F2: ("a" + (("\n\t" + "hello") + "hello"))}
                                v3 = ("" == "foo bar")
                                switch "hello" {
                                case "":
                                    var v29 uint8 = 1
                                    _ = v29
                                default:
                                    v30 := GenS1{F0: "", F1: 496, F2: 915}
                                    _ = v30
                                }
                                v3 = (42 >= 169)
                                v31 := GenS3{F0: (((float64(1.5) != 0.0) && true) && (!v3)), F1: 255, F2: ""}
                                _ = fi27
                                _ = v28
                                _ = v31
                            }
                            _ = v26
                        }
                        var v32 uint = 135
                        v33 := GenS4{F0: 10}
                        _ = fi25
                        _ = v32
                        _ = v33
                    }
                    _ = len(v2)
                } else {
                    v34 := GenS5{F0: -1, F1: 2, F2: 536}
                    v35 := GenS6{F0: -636, F1: 100, F2: 549}
                    if v3 {
                        fmt.Println(v3)
                    }
                    _ = v34
                    _ = v35
                }
            }
        } else {
            v3 = (4294967295 < 65535)
            v3 = v3
        }
        v3 = v3
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        if false {
            if (!(!(4294967295 == 4294967295))) {
                fmt.Println(fi36, v3)
                switch ("hello" + "a") {
                case "\n\t":
                    v3 = v3
                case "":
                    var v37 int64 = -9223372036854775808
                    _ = v37
                case "hello":
                    v3 = (false && (42 > 127))
                }
            } else {
                for fi38 := 0; fi38 < 10; fi38++ {
                    v3 = true
                    var v39 int64 = -117
                    _ = fi38
                    _ = v39
                }
                fmt.Println(fi36, v3)
            }
        } else {
            v40 := GenS6{F0: -245, F1: 127, F2: 274}
            for fi41 := 0; fi41 < 15; fi41++ {
                if v3 {
                    switch (fi41 + fi36) {
                    case 2147483647:
                        if v3 {
                            var v42 string = ("a" + ("a" + ("a" + "hello")))
                            fmt.Println(v42)
                            _ = v42
                        } else {
                            v3 = v3
                            _ = len(v1)
                        }
                    case 406:
                        v43 := GenS6{F0: 9223372036854775807, F1: genFn0(2, 155), F2: 4294967295}
                        _ = v43
                    case -623:
                        fmt.Println(fi41)
                    default:
                        v44 := GenS6{F0: 1, F1: genFn0(32767, 42), F2: 2}
                        _ = v44
                    }
                    v45 := GenS1{F0: ("a" + "\n\t"), F1: 0, F2: 853}
                    _ = v45
                }
                _ = fi41
            }
            var v46 uint16 = 2
            _ = v40
            _ = v46
        }
        _ = fi36
    }
    if (369 < 1) {
        var v47 uint16 = 126
        _ = v47
    } else {
        fmt.Println(v3)
        if true {
            var v48 uint64 = 308
            if v3 {
                v49 := GenS6{F0: 89, F1: 1, F2: 2}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch 32768 {
                case 127:
                    switch 1 {
                    case 2147483647:
                        if (("\n\t" + "") > "hello") {
                            var v50 uint16 = 357
                            v51 := GenS4{F0: (v50 % ((223) | 1))}
                            _ = v50
                            _ = v51
                        } else {
                            for fi52 := 0; fi52 < 16; fi52++ {
                                var v53 uint = 1
                                if v3 {
                                    var v54 int32 = 1
                                    var v55 int = 127
                                    _ = v54
                                    _ = v55
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        for fi56 := 0; fi56 < 11; fi56++ {
                                            var v57 int8 = 62
                                            _ = fi56
                                            _ = v57
                                        }
                                    case "":
                                        v58 := GenS4{F0: 401}
                                        _ = v58
                                    case "foo bar":
                                        for fi59 := 0; fi59 < 3; fi59++ {
                                            v60 := GenS7{F0: 255, F1: v3, F2: (fi59 & (fi52 - fi59))}
                                            var v61 string = "\n\t"
                                            if true {
                                                switch v61 {
                                                case "a":
                                                    var v62 string = (v61 + "hello")
                                                    var v63 uint64 = 42
                                                    var v64 float64 = (3.14159 + -2.25)
                                                    fmt.Println(v63, fi59)
                                                    fmt.Println(fi59, v64, v61)
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                case "\n\t":
                                                    if true {
                                                        v65 := GenS5{F0: -1, F1: v48, F2: 2147483648}
                                                        _ = v65
                                                    }
                                                case "":
                                                    v61 = (v61 + v61)
                                                default:
                                                    if v3 {
                                                        fmt.Println(fi59, fi59, v48)
                                                        var v66 int16 = -1
                                                        _ = v66
                                                    }
                                                }
                                                if v3 {
                                                    if false {
                                                        v2[fi59] = v3
                                                        var v67 uint8 = 127
                                                        if v3 {
                                                            switch "hello" {
                                                            case "hello":
                                                                v68 := GenS2{F0: 881}
                                                                _ = v68
                                                            case "\n\t":
                                                                if false {
                                                                    if (0.0 >= 0.1) {
                                                                        _ = len(v2)
                                                                        v61 = v61
                                                                        var v69 int8 = 104
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v69
                                                                    }
                                                                    if (!v3) {
                                                                        v70 := GenS7{F0: (v67 % (((v67 & (v67 >> ((v53) & 63)))) | 1)), F1: v3, F2: (-(fi59))}
                                                                        v71 := GenS8{F0: 290, F1: 255}
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                } else {
                                                                    v72 := GenS9{F0: 65535, F1: 603}
                                                                    if (9223372036854775807 >= -426) {
                                                                        switch 2147483647 {
                                                                        case 0:
                                                                            _ = len(v1)
                                                                        case 513:
                                                                            v48 += 4294967295
                                                                        case 9223372036854775807:
                                                                            if ((-2147483648 != 2147483647) && false) {
                                                                                v73 := GenS10{F0: (v48 < 42)}
                                                                                fmt.Println(v53)
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 42
                                                                                }
                                                                                _ = v73
                                                                            } else {
                                                                                var v74 bool = v3
                                                                                v53 += uint(v67)
                                                                                v75 := GenS6{F0: 715, F1: 126, F2: (v53 % ((v53) | 1))}
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                        v3 = (false || v3)
                                                                        _ = len(v1)
                                                                    }
                                                                    v3 = v3
                                                                    _ = v72
                                                                }
                                                            }
                                                        } else {
                                                            v76 := GenS5{F0: fi59, F1: v48, F2: 584}
                                                            v3 = (uint8(v53) >= (^(v67)))
                                                            var v77 uint16 = 0
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                        _ = v67
                                                    } else {
                                                        v78 := GenS10{F0: (42 <= -109)}
                                                        switch "" {
                                                        case "":
                                                            v61 = ("\n\t" + v61)
                                                        case "\n\t":
                                                            if v3 {
                                                                fmt.Println(fi52)
                                                                if (42 < 1) {
                                                                    v79 := GenS11{F0: 623, F1: -128, F2: genFn0(57, 0)}
                                                                    var v80 uint64 = (v48 | 42)
                                                                    var v81 int16 = 42
                                                                    _ = v79
                                                                    _ = v80
                                                                    _ = v81
                                                                }
                                                                fmt.Println(v61, fi52)
                                                            } else {
                                                                var v82 string = v61
                                                                _ = v82
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = 1
                                                            }
                                                        default:
                                                            switch 392 {
                                                            case -1:
                                                                if false {
                                                                    var v83 float32 = float32((0.0 - float64(3.14159)))
                                                                    _ = v83
                                                                }
                                                            case 42:
                                                                v84 := GenS12{F0: 0, F1: fi52, F2: v61}
                                                                var v85 float32 = float32(1e10)
                                                                v86 := GenS4{F0: 1}
                                                                v87 := GenS7{F0: 127, F1: ((270 < 0) || false), F2: 9223372036854775807}
                                                                var v88 float64 = 0.1
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                                _ = v87
                                                                _ = v88
                                                            case 127:
                                                                var v89 int16 = 32767
                                                                _ = v89
                                                            default:
                                                                var v90 int64 = 1
                                                                _ = v90
                                                            }
                                                        }
                                                        _ = v78
                                                    }
                                                    if v3 {
                                                        var v91 int = fi52
                                                        _ = v91
                                                    }
                                                    switch ("" + v61) {
                                                    case "":
                                                        switch v61 {
                                                        case "\n\t":
                                                            switch 32768 {
                                                            case -559:
                                                                v61 = ""
                                                            }
                                                            fmt.Println(fi59)
                                                        case "hello":
                                                            var v92 uint64 = 0
                                                            _ = v92
                                                        case "a":
                                                            delete(v2, (-(fi52)))
                                                        default:
                                                            var v93 float32 = 1e10
                                                            _ = v93
                                                        }
                                                    default:
                                                        switch 9223372036854775807 {
                                                        case 4294967295:
                                                            var v94 uint16 = 666
                                                            _ = v94
                                                        case -1:
                                                            if v3 {
                                                                v3 = v3
                                                                if (255 != 1) {
                                                                    var v95 uint8 = 127
                                                                    var v96 uint16 = 42
                                                                    _ = v95
                                                                    _ = v96
                                                                } else {
                                                                    delete(v2, -2147483648)
                                                                    if (0.1 == (1.5 + (0.1 - (-2.25 - -2.25)))) {
                                                                        var v97 uint8 = 116
                                                                        v98 := GenS8{F0: 999, F1: v97}
                                                                        var v99 float64 = (1.5 / 1e10)
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 2
                                                                        }
                                                                        fmt.Println(v48, v99, fi52)
                                                                        var v100 int64 = -635
                                                                        if (4294967295 <= 1) {
                                                                            var v101 int32 = 1
                                                                            fmt.Println(v3, v53, fi59)
                                                                            _ = v101
                                                                        }
                                                                        switch fi52 {
                                                                        case 725:
                                                                            fmt.Println(fi59, v99)
                                                                        default:
                                                                            switch (fi52 | -9223372036854775808) {
                                                                            case 128:
                                                                                _ = len(v1)
                                                                            }
                                                                        }
                                                                        if (42 != 127) {
                                                                            fmt.Println(v100, fi52, v100)
                                                                            v61 = "foo bar"
                                                                            switch fi59 {
                                                                            case -1:
                                                                                v100 = v100
                                                                            case 4294967295:
                                                                                var v102 uint = (^(v53))
                                                                                _ = v102
                                                                            default:
                                                                                if (42 <= v100) {
                                                                                    v61 = (v61 + v61)
                                                                                    v103 := v2[int(v100)]
                                                                                    var v104 string = "\n\t"
                                                                                    if (v53 == (v53 & v53)) {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 825
                                                                                        }
                                                                                        var v105 uint16 = 255
                                                                                        _ = v105
                                                                                    } else {
                                                                                        var v106 float32 = ((0.1 / 0.1) - (1e10 - 0.1))
                                                                                        v107 := GenS13{F0: 1000, F1: (v100 + (v100 ^ -852)), F2: (v48 ^ 505)}
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                    }
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                }
                                                                            }
                                                                        }
                                                                        v97 += 255
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        fmt.Println(v97)
                                                                        if (9223372036854775807 == 42) {
                                                                            if true {
                                                                                v1 = append(v1, 114)
                                                                                v108 := GenS14{F0: (42 != 410), F1: 92, F2: 42}
                                                                                _ = len(v1)
                                                                                _ = v108
                                                                            }
                                                                            var v109 int32 = -1
                                                                            _ = v109
                                                                        }
                                                                        fmt.Println(fi59, v61, v61)
                                                                        _ = v100
                                                                        _ = v97
                                                                        _ = v98
                                                                        _ = v99
                                                                    } else {
                                                                        var v110 bool = (!v3)
                                                                        _ = v110
                                                                    }
                                                                    v2[(fi59 ^ fi52)] = (2 > 42)
                                                                }
                                                            } else {
                                                                v53 -= (v53 >> ((v53) & 63))
                                                            }
                                                        default:
                                                            var v111 int32 = 0
                                                            _ = v111
                                                        }
                                                    }
                                                    var v112 float32 = 0.1
                                                    _ = v112
                                                }
                                                v113 := GenS15{F0: 0, F1: 128, F2: (v3 || true)}
                                                _ = v113
                                            }
                                            _ = fi59
                                            _ = v60
                                            _ = v61
                                        }
                                    default:
                                        v114 := GenS5{F0: (fi52 & fi52), F1: (v48 << ((v53) & 63)), F2: 2147483647}
                                        _ = v114
                                    }
                                    _ = len(v1)
                                    v48 = v48
                                }
                                _ = fi52
                                _ = v53
                            }
                            v3 = true
                            v48 = (v48 * (^(v48)))
                        }
                    default:
                        switch ("\n\t" + "a") {
                        case "":
                            v115 := v2[2147483647]
                            _ = v115
                        case "foo bar":
                            v48 = v48
                        case "\n\t":
                            if (35 > -32768) {
                                var v116 float32 = 1.5
                                _ = v116
                            }
                        }
                    }
                case 9223372036854775807:
                    v3 = true
                default:
                    if false {
                        var v117 int32 = 485
                        _ = v117
                    } else {
                        var v118 uint32 = 403
                        fmt.Println(v48, v48)
                        if (127 < 12) {
                            for fi119 := 0; fi119 < 9; fi119++ {
                                if v3 {
                                    v120 := GenS15{F0: 1, F1: -921, F2: (255 < 396)}
                                    _ = v120
                                }
                                v121 := GenS15{F0: 32768, F1: (-(fi119)), F2: (!v3)}
                                fmt.Println(v48)
                                _ = fi119
                                _ = v121
                            }
                            var v122 uint = 63
                            var v123 int64 = -1
                            _ = v122
                            _ = v123
                        } else {
                            var v124 int64 = -833
                            _ = v124
                        }
                        _ = v118
                    }
                }
                _ = v49
            }
            for fi125 := 0; fi125 < 15; fi125++ {
                fmt.Println(fi125, v3)
                _ = fi125
            }
            _ = v48
        }
    }
    v3 = (!v3)
    fmt.Println(v3)
    if true {
        for fi126 := 0; fi126 < 16; fi126++ {
            v3 = v3
            var v127 bool = true
            v1 = append(v1, 988)
            _ = fi126
            _ = v127
        }
    }
    for fi128 := 0; fi128 < 3; fi128++ {
        v3 = v3
        _ = fi128
    }
    fmt.Println(v3)
    v3 = (false || (2 == 1))
    var v129 uint = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v129:", v129)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v129
    _ = v2
    _ = v3
}
