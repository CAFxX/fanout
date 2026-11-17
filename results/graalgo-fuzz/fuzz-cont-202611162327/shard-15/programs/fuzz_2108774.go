package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int8; F1 int64; F2 int16 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string; F1 int8; F2 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS12 struct { F0 int32; F1 int64; F2 int32 }
type GenS13 struct { F0 uint16; F1 int }

func main() {
    var v1 uint8 = 1
    v2 := []string{("hello" + ("foo bar" + "hello"))}
    var v3 int32 = -173
    fmt.Println(v1, v3)
    var v4 int32 = (v3 + v3)
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(v1, v1, fi5)
        _ = fi5
    }
    var v6 uint64 = 0
    if false {
        if true {
            var v7 uint = 2
            for fi8 := 0; fi8 < 13; fi8++ {
                v6 += 0
                var v9 float64 = ((-2.25 / 1e10) - 1.5)
                for fi10 := 0; fi10 < 18; fi10++ {
                    v4 -= -574
                    v11 := GenS1{F0: 768, F1: false, F2: 42}
                    v3 -= int32(v7)
                    _ = fi10
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
            fmt.Println(v4)
            _ = v7
        } else {
            var v12 uint64 = v6
            _ = v12
        }
    } else {
        if false {
            v1 = 128
            if false {
                v13 := GenS1{F0: 255, F1: (!(false || true)), F2: 255}
                if (65535 == 42) {
                    switch -2147483648 {
                    case 498:
                        for fi14 := 0; fi14 < 5; fi14++ {
                            for fi15 := 0; fi15 < 12; fi15++ {
                                fmt.Println(fi14, v4, v1)
                                v1 = (v1 + v1)
                                if len(v2) > 0 {
                                    v2[0] = "foo bar"
                                }
                                _ = fi15
                            }
                            _ = fi14
                        }
                    }
                }
                switch -3 {
                case -2147483648:
                    var v16 int32 = (v4 * v3)
                    _ = v16
                case 4294967295:
                    if ("a" >= "hello") {
                        if (false && ((-145 <= -575) && (!false))) {
                            v17 := GenS2{F0: "hello", F1: true}
                            switch -713 {
                            case 127:
                                if (!true) {
                                    var v18 int64 = -155
                                    _ = v18
                                }
                            case -2147483648:
                                if (false && false) {
                                    if false {
                                        var v19 uint8 = (^(v1))
                                        _ = v19
                                    } else {
                                        for fi20 := 0; fi20 < 16; fi20++ {
                                            var v21 float32 = 0.0
                                            var v22 uint32 = 440
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                    v2 = append(v2, "hello")
                                    if (true || true) {
                                        fmt.Println(v3, v3)
                                    }
                                }
                            case -901:
                                v6 = 0
                            default:
                                v4 += (v4 & v4)
                            }
                            _ = v17
                        } else {
                            v23 := GenS3{F0: 805, F1: -307, F2: (v6 - (v6 >> ((87) & 63)))}
                            var v24 uint8 = v1
                            _ = v23
                            _ = v24
                        }
                    } else {
                        if false {
                            for fi25 := 0; fi25 < 14; fi25++ {
                                for fi26 := 0; fi26 < 14; fi26++ {
                                    if true {
                                        var v27 int32 = (v3 % (((v4 / ((v3) | 1))) | 1))
                                        _ = v27
                                    } else {
                                        v1 += 42
                                        switch "a" {
                                        case "hello":
                                            switch "hello" {
                                            case "foo bar":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    fmt.Println(fi26)
                                                case "":
                                                    fmt.Println(fi26, v6)
                                                case "hello":
                                                    var v28 int8 = -1
                                                    _ = v28
                                                default:
                                                    v4 = v4
                                                }
                                            case "\n\t":
                                                var v29 uint = 42
                                                _ = v29
                                            case "a":
                                                var v30 int64 = -1
                                                _ = v30
                                            default:
                                                var v31 int64 = 42
                                                _ = v31
                                            }
                                        case "\n\t":
                                            var v32 int = (fi25 * (fi25 << ((2) & 63)))
                                            _ = v32
                                        default:
                                            if true {
                                                if false {
                                                    var v33 uint16 = 924
                                                    var v34 float32 = float32((1e10 + 3.14159))
                                                    _ = len(v2)
                                                    var v35 uint32 = 4294967295
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                }
                                                v4 -= v3
                                                v1 += v1
                                                if (!((float32(3.14159) + (3.14159 / 1e10)) == (float32(0.1) / float32(3.14159)))) {
                                                    v36 := GenS3{F0: 2, F1: 0, F2: (v6 + v6)}
                                                    _ = v36
                                                }
                                                if true {
                                                    fmt.Println(fi26, v4)
                                                }
                                            }
                                        }
                                    }
                                    v37 := GenS1{F0: 2, F1: true, F2: 65535}
                                    v2 = append(v2, "\n\t")
                                    _ = fi26
                                    _ = v37
                                }
                                v38 := GenS4{F0: 2, F1: "hello"}
                                switch (^(fi25)) {
                                case 2147483647:
                                    for fi39 := 0; fi39 < 8; fi39++ {
                                        v1 += 40
                                        if false {
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                            if false {
                                                var v40 uint8 = 42
                                                v1 = (v1 / (((v1 | v40)) | 1))
                                                if (815 == -689) {
                                                    v41 := GenS5{F0: -110, F1: 42, F2: 1000}
                                                    if true {
                                                        fmt.Println(v40, v1)
                                                        if (-717 != -1) {
                                                            var v42 int64 = 0
                                                            if (2 != 2) {
                                                                switch (fi39 / ((fi25) | 1)) {
                                                                case -2147483648:
                                                                    v43 := GenS4{F0: 1, F1: ""}
                                                                    fmt.Println(v3, v42)
                                                                    v3 = (v3 * (v4 * v3))
                                                                    _ = v43
                                                                case 4294967295:
                                                                    if (!(531 == 177)) {
                                                                        v2 = append(v2, "")
                                                                        var v44 uint64 = v6
                                                                        _ = v44
                                                                    } else {
                                                                        v1 = (^(v40))
                                                                    }
                                                                case -500:
                                                                    if true {
                                                                        v45 := GenS6{F0: (v6 - 40), F1: -1}
                                                                        switch 42 {
                                                                        case 42:
                                                                            fmt.Println(v3)
                                                                        case 9223372036854775807:
                                                                            var v46 uint8 = (^(v1))
                                                                            _ = v46
                                                                        case 32768:
                                                                            v47 := GenS7{F0: 694}
                                                                            _ = v47
                                                                        default:
                                                                            v48 := GenS4{F0: 65535, F1: "\n\t"}
                                                                            _ = v48
                                                                        }
                                                                        if len(v2) > 0 {
                                                                            v2[0] = ("a" + "\n\t")
                                                                        }
                                                                        _ = v45
                                                                    } else {
                                                                        if false {
                                                                            v4 -= (v4 << ((uint(v1)) & 63))
                                                                            var v49 string = "hello"
                                                                            _ = v49
                                                                        } else {
                                                                            fmt.Println(v4)
                                                                        }
                                                                        if (32767 >= 487) {
                                                                            var v50 string = ("" + "")
                                                                            v51 := GenS3{F0: v6, F1: -660, F2: (v6 * 2)}
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                default:
                                                                    v52 := GenS8{F0: (1 <= 42)}
                                                                    _ = v52
                                                                }
                                                            } else {
                                                                switch 270 {
                                                                case 0:
                                                                    var v53 float64 = -2.25
                                                                    _ = v53
                                                                default:
                                                                    var v54 string = "hello"
                                                                    _ = v54
                                                                }
                                                                if true {
                                                                    var v55 int8 = 0
                                                                    v56 := GenS1{F0: 915, F1: false, F2: 902}
                                                                    _ = v55
                                                                    _ = v56
                                                                } else {
                                                                    var v57 int = fi39
                                                                    v1 = (v1 & 1)
                                                                    _ = v57
                                                                }
                                                            }
                                                            v2 = append(v2, ((("hello" + "") + ("a" + ("a" + "a"))) + "a"))
                                                            if (false || (42 >= 32767)) {
                                                                if false {
                                                                    fmt.Println(v3, v3)
                                                                    var v58 uint8 = (v1 * 255)
                                                                    _ = v58
                                                                }
                                                                var v59 int64 = int64(v6)
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v59
                                                            }
                                                            _ = v42
                                                        } else {
                                                            _ = len(v2)
                                                            v60 := GenS2{F0: "", F1: (255 <= 1)}
                                                            v4 += -1
                                                            _ = v60
                                                        }
                                                        if (false || (9223372036854775807 == -9223372036854775808)) {
                                                            if false {
                                                                v6 = v6
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v61 uint64 = v6
                                                                    _ = v61
                                                                case "\n\t":
                                                                    v1 -= 127
                                                                case "a":
                                                                    v62 := GenS7{F0: 1}
                                                                    _ = v62
                                                                default:
                                                                    if ((v6 == 4294967295) || (-2.25 < (1.5 * 0.1))) {
                                                                        v2 = append(v2, "a")
                                                                    } else {
                                                                        v6 += v6
                                                                        if (false || (v3 >= (^(v3)))) {
                                                                            var v63 uint8 = v1
                                                                            v64 := GenS9{F0: ("a" + "a"), F1: 121, F2: (359 > -1)}
                                                                            _ = v63
                                                                            _ = v64
                                                                        } else {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            if false {
                                                                                v65 := GenS3{F0: 990, F1: -32768, F2: v6}
                                                                                var v66 int8 = int8(v4)
                                                                                _ = v65
                                                                                _ = v66
                                                                            }
                                                                            var v67 int32 = (v3 << ((385) & 63))
                                                                            _ = v67
                                                                        }
                                                                    }
                                                                }
                                                                var v68 uint64 = v6
                                                                _ = v68
                                                            } else {
                                                                var v69 int8 = -128
                                                                v70 := GenS4{F0: 973, F1: ("hello" + ("\n\t" + "\n\t"))}
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            if ((v6 | 935) >= (^(v6))) {
                                                                var v71 float32 = float32(-2.25)
                                                                v72 := GenS8{F0: (1 != -128)}
                                                                var v73 int = fi25
                                                                var v74 int8 = 100
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v71
                                                                _ = v72
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                            if false {
                                                                if (!(2 <= 487)) {
                                                                    var v75 string = ("foo bar" + "a")
                                                                    var v76 string = v75
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    _ = v75
                                                                    _ = v76
                                                                } else {
                                                                    _ = len(v2)
                                                                    fmt.Println(v1)
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    fmt.Println(v1)
                                                                }
                                                            }
                                                            v77 := GenS9{F0: "\n\t", F1: 42, F2: false}
                                                            _ = v77
                                                        } else {
                                                            if false {
                                                                var v78 uint64 = v6
                                                                var v79 uint32 = 287
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                            if (true && false) {
                                                                v6 = v6
                                                            } else {
                                                                _ = len(v2)
                                                                if len(v2) > 0 {
                                                                    v2[0] = "a"
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = len(v2)
                                                    _ = v41
                                                } else {
                                                    if true {
                                                        var v80 uint32 = 951
                                                        _ = v80
                                                    } else {
                                                        v3 = (v4 & -449)
                                                        v40 = v40
                                                    }
                                                }
                                                v81 := GenS10{F0: 18446744073709551615}
                                                _ = v40
                                                _ = v81
                                            }
                                        } else {
                                            v1 = v1
                                            var v82 int8 = 124
                                            _ = v82
                                        }
                                        _ = fi39
                                    }
                                case 62:
                                    for fi83 := 0; fi83 < 4; fi83++ {
                                        v6 += v6
                                        v2 = append(v2, (("hello" + "hello") + "hello"))
                                        _ = fi83
                                    }
                                default:
                                    v84 := GenS9{F0: "a", F1: 100, F2: (9223372036854775807 >= 9223372036854775807)}
                                    _ = v84
                                }
                                _ = fi25
                                _ = v38
                            }
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            v3 = int32(v1)
                        } else {
                            var v85 int64 = -1
                            _ = v85
                        }
                        v86 := GenS10{F0: 521}
                        for fi87 := 0; fi87 < 8; fi87++ {
                            v88 := GenS7{F0: 1}
                            _ = fi87
                            _ = v88
                        }
                        _ = v86
                    }
                case 128:
                    var v89 int64 = -1
                    _ = v89
                }
                _ = v13
            } else {
                var v90 int32 = 1
                v2 = append(v2, "a")
                if (707 >= 2147483647) {
                    var v91 int64 = -952
                    var v92 uint8 = v1
                    for fi93 := 0; fi93 < 11; fi93++ {
                        v94 := GenS11{F0: uint8(v6), F1: false, F2: 0}
                        if true {
                            v91 = (v91 % ((v91) | 1))
                            v90 = v3
                        }
                        fmt.Println(v3)
                        _ = fi93
                        _ = v94
                    }
                    _ = v91
                    _ = v92
                }
                _ = v90
            }
        } else {
            var v95 int16 = 42
            _ = v95
        }
        var v96 uint16 = 952
        fmt.Println(v3, v1, v96)
        _ = v96
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v97 bool = ((4294967295 <= uint(v6)) || ((!true) || true))
    if false {
        var v98 float64 = ((((((1e10 / 1.0) * 3.14159) + 0.0) * (3.14159 / (float64(float32(0.0)) - -2.25))) + (1.5 - 3.14159)) - (1.5 - (1e10 + 1.5)))
        _ = v98
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi99 := 0; fi99 < 20; fi99++ {
        v3 += (v3 | v4)
        var v100 int64 = -9223372036854775808
        var v101 float64 = float64(0.0)
        _ = fi99
        _ = v100
        _ = v101
    }
    if true {
        v97 = ((!(42 <= v1)) || v97)
    } else {
        var v102 string = ("\n\t" + "")
        _ = v102
    }
    fmt.Println(v97)
    if (!false) {
        v97 = (v97 || v97)
    } else {
        fmt.Println(v97)
    }
    if (!false) {
        if true {
            if false {
                switch 42 {
                case -2147483648:
                    for fi103 := 0; fi103 < 9; fi103++ {
                        for fi104 := 0; fi104 < 18; fi104++ {
                            var v105 int = 127
                            var v106 string = (((("foo bar" + "foo bar") + "hello") + "foo bar") + ("\n\t" + "a"))
                            _ = fi104
                            _ = v105
                            _ = v106
                        }
                        for fi107 := 0; fi107 < 15; fi107++ {
                            var v108 uint16 = 65535
                            _ = fi107
                            _ = v108
                        }
                        _ = fi103
                    }
                }
            } else {
                for fi109 := 0; fi109 < 2; fi109++ {
                    v2 = append(v2, ("hello" + "a"))
                    _ = fi109
                }
                var v110 int8 = 100
                _ = v110
            }
            for fi111 := 0; fi111 < 5; fi111++ {
                for fi112 := 0; fi112 < 8; fi112++ {
                    v4 = 36
                    v113 := GenS5{F0: 102, F1: -423, F2: 199}
                    _ = fi112
                    _ = v113
                }
                if v97 {
                    if ((v1 + v1) == (^(v1))) {
                        switch int(v6) {
                        case 32767:
                            if (!true) {
                                var v114 uint32 = 0
                                _ = v114
                            }
                        }
                    }
                } else {
                    if false {
                        var v115 bool = (42 <= -9223372036854775808)
                        _ = v115
                    } else {
                        var v116 int8 = 63
                        switch (fi111 / ((fi111) | 1)) {
                        case 293:
                            if v97 {
                                var v117 uint16 = 1
                                for fi118 := 0; fi118 < 1; fi118++ {
                                    fmt.Println(v6, v117)
                                    v119 := GenS12{F0: 2147483647, F1: -9223372036854775808, F2: v3}
                                    _ = fi118
                                    _ = v119
                                }
                                _ = v117
                            } else {
                                if v97 {
                                    var v120 string = ""
                                    _ = v120
                                } else {
                                    v121 := GenS13{F0: 255, F1: -1}
                                    _ = len(v2)
                                    _ = v121
                                }
                                v122 := GenS5{F0: v116, F1: 42, F2: 57}
                                _ = v122
                            }
                        }
                        _ = v116
                    }
                    fmt.Println(v4)
                    fmt.Println(v3)
                }
                v123 := GenS11{F0: (v1 - (v1 << ((255) & 63))), F1: v97, F2: 323}
                _ = fi111
                _ = v123
            }
        }
        var v124 int64 = -9223372036854775808
        _ = v124
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (!(v97 || (813 < 4294967295))) {
            _ = len(v2)
            if v97 {
                var v125 int8 = -27
                _ = v125
            } else {
                fmt.Println(v6, v4, v3)
                switch 32768 {
                case -9223372036854775808:
                    var v126 uint8 = 1
                    _ = v126
                case 879:
                    var v127 int32 = v3
                    _ = v127
                default:
                    var v128 float64 = 0.0
                    _ = v128
                }
                v1 = v1
            }
            var v129 uint8 = v1
            _ = v129
        } else {
            _ = len(v2)
            v130 := GenS6{F0: (v6 & 4294967295), F1: 1000}
            v97 = v97
            _ = v130
        }
    }
    if len(v2) > 0 {
        v2[0] = ("hello" + "foo bar")
    }
    v97 = (255 <= 467)
    v4 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v97
}
