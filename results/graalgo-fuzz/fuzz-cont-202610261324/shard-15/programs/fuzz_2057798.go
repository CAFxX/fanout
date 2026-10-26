package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint; F1 int64; F2 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint8 }
type GenS11 struct { F0 int16; F1 int8; F2 uint16 }
type GenS12 struct { F0 uint; F1 int64 }
type GenS13 struct { F0 int32; F1 int8; F2 int8 }
type GenS14 struct { F0 int; F1 uint64; F2 uint16 }
type GenS15 struct { F0 uint32; F1 uint8 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint; F1 string }
func genFn0() int {
    var v7 int16 = 42
    v7 = 352
    if true {
        var v8 uint16 = 1
        _ = v8
    } else {
        v9 := []int{32768}
        if true {
            fmt.Println(v7)
            v7 = (v7 & (v7 + 0))
            if false {
                for fi10 := 0; fi10 < 9; fi10++ {
                    for fi11 := 0; fi11 < 19; fi11++ {
                        var v12 uint8 = 127
                        switch -9 {
                        case 691:
                            v9 = append(v9, (fi11 % ((4294967295) | 1)))
                        default:
                            v12 = (^(v12))
                        }
                        _ = fi11
                        _ = v12
                    }
                    var v13 int8 = 100
                    _ = fi10
                    _ = v13
                }
                var v14 uint64 = 18446744073709551615
                _ = v14
            }
        }
        _ = v9
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v7 += v7
        fmt.Println(fi15, v7)
        _ = fi15
    }
    _ = v7
    return 32767
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    v3 := []int{-2147483648}
    v4 := []string{("" + ("hello" + "")), "hello"}
    v5 := make(map[uint32]string)
    var v6 uint64 = 770
    if true {
        for fi16 := 0; fi16 < 3; fi16++ {
            v6 = (v6 * (v6 << ((0) & 63)))
            _ = fi16
        }
    } else {
        if (!true) {
            var v17 string = "hello"
            _ = v17
        }
        if (("hello" + "\n\t") <= "foo bar") {
            v6 = (v6 ^ (v6 / ((1) | 1)))
            delete(v2, 782)
            for fi18 := 0; fi18 < 19; fi18++ {
                for fi19 := 0; fi19 < 17; fi19++ {
                    v20 := GenS1{F0: 0, F1: 4294967295}
                    v6 = v6
                    _ = fi19
                    _ = v20
                }
                var v21 uint = 946
                _ = fi18
                _ = v21
            }
        }
    }
    switch -2147483648 {
    case 0:
        if true {
            v6 -= 4294967295
            v22 := GenS2{F0: "a", F1: -9223372036854775808}
            if (0 <= -9223372036854775808) {
                fmt.Println(v6)
                for fi23 := 0; fi23 < 12; fi23++ {
                    v24 := GenS3{F0: false, F1: (("a" + "") + ("hello" + ""))}
                    v6 -= v6
                    if false {
                        fmt.Println(fi23)
                        for fi25 := 0; fi25 < 7; fi25++ {
                            v26 := GenS4{F0: 61, F1: 255}
                            var v27 string = "a"
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        v28 := v2[922]
                        _ = v28
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            var v29 float32 = ((1e10 * 3.14159) - 0.0)
                            _ = v29
                        case "hello":
                            delete(v2, 65535)
                        }
                        v4 = append(v4, "")
                        v6 -= v6
                    }
                    _ = fi23
                    _ = v24
                }
                v30 := GenS5{F0: 18446744073709551615}
                _ = v30
            } else {
                if (true || false) {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        switch genFn0() {
                        case 565:
                            switch 127 {
                            case -2147483648:
                                v6 -= (v6 / ((19) | 1))
                            }
                        default:
                            switch int(v6) {
                            case -9223372036854775808:
                                var v31 uint8 = 1
                                _ = v31
                            case -2147483648:
                                v32 := GenS4{F0: 65535, F1: 417}
                                _ = v32
                            case 342:
                                for fi33 := 0; fi33 < 12; fi33++ {
                                    var v34 string = (("foo bar" + (("\n\t" + "") + "a")) + (("" + "hello") + "hello"))
                                    v3 = append(v3, genFn0())
                                    v34 = ("\n\t" + v34)
                                    _ = fi33
                                    _ = v34
                                }
                            default:
                                var v35 int32 = -765
                                _ = v35
                            }
                        }
                    case "foo bar":
                        var v36 float32 = 0.1
                        _ = v36
                    case "a":
                        if false {
                            var v37 uint8 = 140
                            _ = v37
                        } else {
                            var v38 bool = (1 <= 4294967295)
                            _ = v38
                        }
                    default:
                        var v39 uint8 = 0
                        _ = v39
                    }
                    var v40 int8 = -128
                    _ = v40
                } else {
                    var v41 uint32 = 282
                    _ = v41
                }
                switch "\n\t" {
                case "a":
                    for fi42 := 0; fi42 < 13; fi42++ {
                        for fi43 := 0; fi43 < 20; fi43++ {
                            switch "\n\t" {
                            case "":
                                v3 = append(v3, fi43)
                            default:
                                var v44 int16 = -97
                                var v45 int = (fi42 - (fi43 << ((42) & 63)))
                                switch fi43 {
                                case 127:
                                    var v46 uint = 217
                                    fmt.Println(fi42, v46)
                                    var v47 float64 = 1e10
                                    _ = v46
                                    _ = v47
                                case -1:
                                    v6 = v6
                                case 680:
                                    v6 -= v6
                                }
                                _ = v44
                                _ = v45
                            }
                            delete(v2, 4294967295)
                            var v48 int8 = 0
                            v48 = -100
                            _ = fi43
                            _ = v48
                        }
                        _ = fi42
                    }
                case "\n\t":
                    _ = len(v2)
                case "hello":
                    for fi49 := 0; fi49 < 15; fi49++ {
                        v6 -= (^(v6))
                        _ = fi49
                    }
                }
                for fi50 := 0; fi50 < 14; fi50++ {
                    for fi51 := 0; fi51 < 15; fi51++ {
                        fmt.Println(fi51, fi51)
                        _ = fi51
                    }
                    _ = fi50
                }
            }
            _ = v22
        }
    case -586:
        if true {
            fmt.Println(v6)
            for fi52 := 0; fi52 < 19; fi52++ {
                switch ("" + "foo bar") {
                case "a":
                    var v53 uint64 = (v6 / (((v6 / ((v6) | 1))) | 1))
                    _ = v53
                case "foo bar":
                    v6 -= (^(v6))
                }
                var v54 int = fi52
                var v55 uint64 = v6
                _ = fi52
                _ = v54
                _ = v55
            }
            if true {
                var v56 uint32 = 2147483647
                _ = v56
            }
        }
    }
    fmt.Println(v6)
    v6 += (v6 | v6)
    v57 := GenS6{F0: -1, F1: 100}
    var v58 float64 = 0.1
    fmt.Println(v58, v58)
    var v59 int64 = -965
    if (!true) {
        switch 42 {
        case 128:
            var v60 int16 = 177
            _ = v60
        case 936:
            v58 -= 3.14159
        case 2147483647:
            var v61 bool = false
            _ = v61
        }
    } else {
        if true {
            var v62 uint16 = 2
            for fi63 := 0; fi63 < 7; fi63++ {
                switch fi63 {
                case -9223372036854775808:
                    var v64 int64 = (v59 | (v59 - (^(v59))))
                    _ = v64
                default:
                    var v65 int = 32768
                    _ = v65
                }
                for fi66 := 0; fi66 < 16; fi66++ {
                    delete(v5, 0)
                    var v67 int16 = 42
                    if (!((361 < 2147483647) || true)) {
                        switch ("hello" + ("hello" + "a")) {
                        case "":
                            switch "" {
                            case "foo bar":
                                var v68 uint8 = 42
                                if false {
                                    v68 = 25
                                    var v69 uint64 = (v6 >> ((255) & 63))
                                    if ((0 > v59) && true) {
                                        if (0 >= v67) {
                                            var v70 float32 = float32(((float64(1e10) / ((v58 + v58) / ((v58 + (v58 + 1.5)) * (-2.25 - (float64((((-2.25 / 1e10) + 3.14159) + -2.25)) / (float64((0.0 - 0.0)) + v58)))))) + float64(((0.0 - float32(v58)) * 0.0))))
                                            v68 = (v68 << ((840) & 63))
                                            var v71 int = fi63
                                            _ = v70
                                            _ = v71
                                        } else {
                                            delete(v5, 2)
                                            v72 := GenS7{F0: 42, F1: 42}
                                            if true {
                                                v58 = 3.14159
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    v6 = 0
                                                case "a":
                                                    if false {
                                                        if ((v67 * 1000) != int16(v62)) {
                                                            var v73 uint8 = (v68 * 10)
                                                            _ = v73
                                                        }
                                                        if (true || true) {
                                                            var v74 int8 = int8(v6)
                                                            _ = v74
                                                        }
                                                        v59 = (-(v59))
                                                    }
                                                case "\n\t":
                                                    var v75 uint16 = v62
                                                    _ = v75
                                                default:
                                                    if false {
                                                        v69 = v6
                                                    }
                                                }
                                                v68 += (v68 & (v68 / ((v68) | 1)))
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            _ = v72
                                        }
                                        if true {
                                            var v76 int = (^(fi63))
                                            fmt.Println(v62, fi66)
                                            _ = len(v3)
                                            _ = v76
                                        } else {
                                            v2[42] = (fi63 % (((fi63 - int(v62))) | 1))
                                            v58 = -2.25
                                            var v77 bool = (!((v68 | v68) > (v68 | 128)))
                                            _ = v77
                                        }
                                    }
                                    _ = v69
                                } else {
                                    v3 = append(v3, (fi63 & (fi66 ^ 127)))
                                    fmt.Println(v67, v59, v58)
                                    var v78 string = "foo bar"
                                    fmt.Println(v6, v62, v62)
                                    _ = v78
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v68
                            case "hello":
                                fmt.Println(v59, v62, v58)
                            }
                        case "\n\t":
                            switch fi63 {
                            case 750:
                                v58 = float64(0.1)
                            case 9223372036854775807:
                                switch fi66 {
                                case 32767:
                                    switch fi63 {
                                    case -2147483648:
                                        v79 := GenS8{F0: 4294967295, F1: v59, F2: 1}
                                        _ = v79
                                    case 42:
                                        v80 := GenS3{F0: false, F1: ""}
                                        _ = v80
                                    }
                                    v81 := GenS9{F0: genFn0()}
                                    if (!((v62 > v62) || true)) {
                                        v59 = (v59 - (-(v59)))
                                        _ = len(v1)
                                        v62 += 255
                                        switch "foo bar" {
                                        case "a":
                                            v82 := GenS2{F0: "foo bar", F1: -39}
                                            _ = v82
                                        case "\n\t":
                                            var v83 float64 = (float64(1e10) * 1.5)
                                            _ = v83
                                        case "foo bar":
                                            var v84 uint32 = 2147483648
                                            _ = v84
                                        }
                                    } else {
                                        v59 = -9223372036854775808
                                    }
                                    _ = v81
                                case 4294967295:
                                    v58 = v58
                                case 127:
                                    if (!((true || false) && (((!(!true)) || (false && false)) && true))) {
                                        var v85 int16 = v67
                                        v6 = (v6 >> ((65535) & 63))
                                        switch "hello" {
                                        case "":
                                            fmt.Println(v6, v6)
                                            fmt.Println(v85, v85, v6)
                                        case "\n\t":
                                            var v86 string = "a"
                                            var v87 uint8 = uint8(v6)
                                            _ = v86
                                            _ = v87
                                        case "a":
                                            v88 := GenS10{F0: (-(fi63)), F1: 42}
                                            _ = v88
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        _ = v85
                                    } else {
                                        switch (fi66 % ((genFn0()) | 1)) {
                                        case -720:
                                            v67 = (v67 << ((18446744073709551615) & 63))
                                        }
                                        _ = len(v4)
                                        var v89 int = genFn0()
                                        _ = v89
                                    }
                                default:
                                    v62 += v62
                                }
                            case 670:
                                v67 += 32767
                                fmt.Println(v58)
                                var v90 string = (("" + ("\n\t" + "a")) + ("a" + "\n\t"))
                                _ = v90
                            default:
                                v91 := GenS11{F0: int16(v62), F1: -128, F2: v62}
                                _ = v91
                            }
                        case "foo bar":
                            v6 = (v6 & v6)
                        }
                        v59 -= v59
                    } else {
                        var v92 float32 = ((3.14159 * 1.5) / 1.5)
                        _ = v92
                    }
                    if false {
                        if false {
                            if ((v67 ^ (-(v67))) != v67) {
                                var v93 int16 = 42
                                _ = v93
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(fi63, v6, v58)
                                case "\n\t":
                                    v94 := GenS12{F0: 0, F1: (v59 % (((v59 << ((807) & 63))) | 1))}
                                    _ = v94
                                case "a":
                                    var v95 bool = (fi63 == 2147483647)
                                    _ = v95
                                default:
                                    v96 := GenS3{F0: false, F1: ("" + "hello")}
                                    _ = v96
                                }
                                v58 -= v58
                                _ = len(v2)
                                if true {
                                    v62 -= 2
                                    v6 = (v6 | v6)
                                    if false {
                                        var v97 int16 = 32767
                                        if true {
                                            if (!false) {
                                                if true {
                                                    v58 = v58
                                                    switch (("foo bar" + "foo bar") + "\n\t") {
                                                    case "hello":
                                                        var v98 int64 = 42
                                                        _ = v98
                                                    case "":
                                                        v58 = 3.14159
                                                    case "foo bar":
                                                        if (!((-2.25 / 0.1) >= 3.14159)) {
                                                            fmt.Println(v67, v59, v6)
                                                        }
                                                    }
                                                    v99 := GenS13{F0: 948, F1: -84, F2: 127}
                                                    fmt.Println(v97)
                                                    _ = v99
                                                } else {
                                                    v100 := GenS14{F0: fi66, F1: 0, F2: (v62 % (((v62 | v62)) | 1))}
                                                    _ = v100
                                                }
                                                var v101 int = fi66
                                                v102 := GenS1{F0: 399, F1: 0}
                                                switch fi63 {
                                                case 0:
                                                    var v103 int64 = 1
                                                    _ = v103
                                                default:
                                                    var v104 int = fi63
                                                    _ = v104
                                                }
                                                var v105 uint16 = (v62 & v62)
                                                _ = v101
                                                _ = v102
                                                _ = v105
                                            } else {
                                                fmt.Println(fi63, v6, v59)
                                            }
                                            v106 := GenS5{F0: 0}
                                            v97 -= v97
                                            _ = v106
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            switch fi66 {
                                            case -2147483648:
                                                var v107 uint = 65535
                                                _ = v107
                                            case -321:
                                                v97 = 42
                                            case 0:
                                                v58 = 1.5
                                            default:
                                                if (!("a" > "a")) {
                                                    var v108 int = fi63
                                                    var v109 int16 = v97
                                                    _ = v108
                                                    _ = v109
                                                }
                                            }
                                        case "foo bar":
                                            v59 = 1
                                        }
                                        delete(v2, 603)
                                        _ = v97
                                    }
                                    var v110 int64 = (v59 >> ((1) & 63))
                                    if false {
                                        v67 += (v67 | (v67 + (v67 * 737)))
                                        v62 = (^(v62))
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(fi66)
                                        }
                                    } else {
                                        fmt.Println(v6, v62)
                                    }
                                    _ = v110
                                } else {
                                    if false {
                                        v67 = -1
                                    }
                                    var v111 int64 = (v59 & v59)
                                    v111 += (v111 - 42)
                                    _ = v111
                                }
                                delete(v2, 1)
                            }
                            v112 := GenS10{F0: (fi63 ^ 4294967295), F1: 2}
                            v6 += v6
                            _ = v112
                        }
                        v113 := v2[2]
                        switch v113 {
                        case 32767:
                            if (42 == 255) {
                                v114 := GenS15{F0: 522, F1: 42}
                                fmt.Println(fi66, fi66, v6)
                                _ = v114
                            }
                        case -1:
                            var v115 uint8 = uint8(v62)
                            _ = v115
                        case 0:
                            v116 := GenS16{F0: true}
                            _ = v116
                        }
                        v58 -= (0.1 / (((float64(-2.25) / 1.5) + v58) * v58))
                        fmt.Println(v59, v113)
                        _ = v113
                    }
                    var v117 int32 = -726
                    v118 := GenS3{F0: false, F1: ("\n\t" + ("a" + "\n\t"))}
                    fmt.Println(v62, v59, fi66)
                    var v119 int8 = -87
                    _ = fi66
                    _ = v117
                    _ = v118
                    _ = v119
                    _ = v67
                }
                _ = fi63
            }
            _ = v62
        }
    }
    if false {
        var v120 uint16 = 42
        v121 := GenS14{F0: 32768, F1: 931, F2: v120}
        var v122 int8 = 0
        _ = v120
        _ = v121
        _ = v122
    } else {
        for fi123 := 0; fi123 < 16; fi123++ {
            var v124 uint = 121
            for fi125 := 0; fi125 < 10; fi125++ {
                v126 := GenS5{F0: (v6 + (v6 | (v6 + v6)))}
                fmt.Println(fi123)
                _ = fi125
                _ = v126
            }
            _ = fi123
            _ = v124
        }
        var v127 bool = ((v6 > 42) && true)
        _ = v127
    }
    v6 = (v6 | (v6 | 0))
    var v128 int32 = -1
    if true {
        fmt.Println(v6)
        v3 = append(v3, 4294967295)
        var v129 float32 = 1e10
        _ = v129
    } else {
        v130 := GenS17{F0: 654, F1: "foo bar"}
        _ = v130
    }
    for fi131 := 0; fi131 < 4; fi131++ {
        for fi132 := 0; fi132 < 14; fi132++ {
            var v133 int64 = 767
            var v134 int = fi131
            _ = fi132
            _ = v133
            _ = v134
        }
        switch ("a" + "foo bar") {
        case "a":
            if true {
                for fi135 := 0; fi135 < 1; fi135++ {
                    v59 += -1
                    _ = fi135
                }
            }
        case "":
            var v136 int64 = v59
            _ = v136
        default:
            if ((45 >= -128) && true) {
                fmt.Println(v6)
                if len(v1) > 0 {
                    v1[0] = 1
                }
            }
        }
        switch fi131 {
        case 32767:
            var v137 int32 = 2147483647
            _ = v137
        case -737:
            for fi138 := 0; fi138 < 19; fi138++ {
                if ((v59 + v59) != (v59 - v59)) {
                    var v139 int8 = 127
                    _ = v139
                }
                var v140 bool = false
                v128 += (^(v128))
                _ = fi138
                _ = v140
            }
        }
        _ = fi131
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v128:", v128)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
}
