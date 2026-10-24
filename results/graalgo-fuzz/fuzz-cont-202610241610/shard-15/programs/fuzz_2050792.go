package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint8; F1 int; F2 int }
type GenS5 struct { F0 uint64; F1 bool; F2 int8 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 string }
type GenS9 struct { F0 uint64; F1 bool; F2 bool }
type GenS10 struct { F0 int64; F1 int16 }
type GenS11 struct { F0 bool; F1 string }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS14 struct { F0 int; F1 uint32; F2 int64 }
func genFn0() int8 {
    switch -5 {
    case 0:
        var v5 uint8 = 128
        _ = v5
    case 2147483647:
        var v6 int = 32768
        _ = v6
    case 864:
        var v7 uint64 = 2
        _ = v7
    }
    var v8 uint16 = 208
    var v9 uint16 = v8
    v10 := GenS1{F0: 0, F1: 42}
    _ = v10
    _ = v8
    _ = v9
    return 126
}
func genFn1(p0 uint16, p1 int, p2 int32) uint16 {
    var v11 int32 = (p2 | p2)
    if ("foo bar" <= "") {
        if false {
            p0 = 581
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: 2, F1: (p1 << ((95) & 63))}
                _ = fi12
                _ = v13
            }
            var v14 uint32 = 860
            _ = v14
        }
        p2 += v11
        fmt.Println(v11, p0, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 32767
}

func main() {
    var v1 bool = (false || (468 <= 820))
    v2 := make(map[int]int)
    v3 := []uint32{2, 2, 2147483648}
    v4 := []int{42}
    if v1 {
        v1 = true
    } else {
        var v15 uint64 = 275
        if v1 {
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: 4294967295, F1: int(v15)}
                _ = fi16
                _ = v17
            }
            v18 := GenS1{F0: 2, F1: 32767}
            _ = v18
        } else {
            var v19 int32 = 0
            var v20 int32 = (v19 ^ -469)
            v21 := GenS1{F0: 435, F1: -1}
            _ = v19
            _ = v20
            _ = v21
        }
        _ = v15
    }
    fmt.Println(v1)
    fmt.Println(v1)
    v3 = append(v3, 1)
    for fi22 := 0; fi22 < 8; fi22++ {
        if false {
            if v1 {
                v23 := GenS1{F0: 255, F1: 41}
                var v24 uint32 = 42
                _ = v23
                _ = v24
            }
        }
        _ = fi22
    }
    if (177 == 2) {
        for fi25 := 0; fi25 < 3; fi25++ {
            for fi26 := 0; fi26 < 1; fi26++ {
                v1 = v1
                if (18446744073709551615 <= 0) {
                    if v1 {
                        v1 = v1
                        var v27 uint64 = 2
                        _ = v27
                    } else {
                        fmt.Println(fi25)
                    }
                    switch "hello" {
                    case "foo bar":
                        var v28 uint8 = 0
                        _ = v28
                    case "\n\t":
                        delete(v2, 1)
                    case "hello":
                        v29 := GenS2{F0: 65535}
                        _ = v29
                    default:
                        if v1 {
                            v1 = true
                        } else {
                            if (!(-457 != 42)) {
                                v1 = (!v1)
                                if v1 {
                                    fmt.Println(v1)
                                    v1 = (717 == genFn1(genFn1(2, 32768, -2147483648), fi26, -251))
                                } else {
                                    v30 := GenS1{F0: 4294967295, F1: fi25}
                                    _ = v30
                                }
                            } else {
                                if (2 <= 1) {
                                    fmt.Println(v1, v1, v1)
                                    v1 = v1
                                }
                                var v31 uint16 = 557
                                _ = v31
                            }
                        }
                    }
                } else {
                    v1 = v1
                    v32 := v2[42]
                    _ = v32
                }
                _ = fi26
            }
            _ = fi25
        }
        v33 := GenS2{F0: 42}
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 int16 = 32767
            if v1 {
                var v36 int64 = 1
                v2[(fi34 + fi34)] = 127
                _ = v36
            } else {
                if (-9223372036854775808 != -1) {
                    if (!false) {
                        if false {
                            var v37 uint8 = 100
                            _ = v37
                        } else {
                            if (!false) {
                                if v1 {
                                    if len(v3) > 0 {
                                        v3[0] = 80
                                    }
                                }
                                v1 = false
                                v38 := v2[fi34]
                                _ = v38
                            } else {
                                switch "" {
                                case "foo bar":
                                    if ((^(fi34)) == fi34) {
                                        delete(v2, fi34)
                                        v39 := GenS2{F0: 255}
                                        _ = v39
                                    } else {
                                        var v40 uint8 = 42
                                        v4 = append(v4, fi34)
                                        if v1 {
                                            v41 := GenS2{F0: 84}
                                            var v42 uint16 = 65535
                                            _ = v41
                                            _ = v42
                                        }
                                        _ = v40
                                    }
                                case "hello":
                                    v35 = v35
                                case "\n\t":
                                    v1 = v1
                                }
                                var v43 uint16 = 952
                                fmt.Println(v35, v43, fi34)
                                _ = v43
                            }
                        }
                        if v1 {
                            v35 -= v35
                            if true {
                                _ = len(v2)
                                for fi44 := 0; fi44 < 19; fi44++ {
                                    var v45 int64 = 1
                                    _ = fi44
                                    _ = v45
                                }
                            }
                        }
                    }
                    if false {
                        if (127 > -1) {
                            v35 += 42
                            if v1 {
                                v46 := GenS2{F0: 44}
                                var v47 int32 = int32(fi34)
                                _ = v46
                                _ = v47
                            }
                        } else {
                            v3 = append(v3, 480)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v1, fi34)
                        }
                        switch "hello" {
                        case "foo bar":
                            v48 := GenS3{F0: 0, F1: v1}
                            _ = v48
                        case "hello":
                            v49 := GenS4{F0: 42, F1: 42, F2: (fi34 >> ((257) & 63))}
                            _ = v49
                        case "a":
                            for fi50 := 0; fi50 < 9; fi50++ {
                                v3 = append(v3, 65535)
                                fmt.Println(fi50, fi50)
                                _ = fi50
                            }
                        default:
                            fmt.Println(v35, v1, fi34)
                        }
                    } else {
                        var v51 int64 = -9223372036854775808
                        var v52 int = (fi34 << ((18446744073709551615) & 63))
                        v53 := GenS5{F0: 1, F1: (!v1), F2: -1}
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    var v54 int = fi34
                    _ = v54
                }
                switch "" {
                case "foo bar":
                    if v1 {
                        for fi55 := 0; fi55 < 14; fi55++ {
                            v1 = (-1 > 2147483647)
                            _ = fi55
                        }
                    }
                default:
                    v1 = true
                }
                v56 := GenS3{F0: fi34, F1: v1}
                _ = v56
            }
            v1 = ((!(0 <= 42)) && (-2.25 == 1.5))
            _ = fi34
            _ = v35
        }
        _ = v33
    } else {
        var v57 int16 = 32767
        _ = v57
    }
    if (!(("" <= ("foo bar" + ("foo bar" + ("a" + "a")))) || ("hello" != "foo bar"))) {
        var v58 string = ((("" + "foo bar") + "foo bar") + "\n\t")
        for fi59 := 0; fi59 < 6; fi59++ {
            fmt.Println(fi59, v58)
            v60 := GenS1{F0: 337, F1: fi59}
            if ((v1 && v1) && (42 > 895)) {
                if len(v4) > 0 {
                    v4[0] = (fi59 + (fi59 % ((fi59) | 1)))
                }
                v58 = "a"
                fmt.Println(fi59, fi59, fi59)
            } else {
                v1 = (!v1)
                v61 := GenS6{F0: 520, F1: 259}
                _ = v61
            }
            _ = fi59
            _ = v60
        }
        _ = v58
    } else {
        v2[1] = -357
        if (((-117 > 100) || (852 == 0)) || v1) {
            if v1 {
                v1 = (true || v1)
                if (-1 <= -97) {
                    delete(v2, 42)
                    var v62 string = "foo bar"
                    _ = v62
                } else {
                    if (-1 == -572) {
                        v63 := GenS7{F0: 719}
                        fmt.Println(v1)
                        v2[2147483647] = 2147483647
                        _ = v63
                    }
                    if (!false) {
                        switch "a" {
                        case "\n\t":
                            for fi64 := 0; fi64 < 6; fi64++ {
                                v65 := GenS6{F0: 136, F1: 415}
                                _ = fi64
                                _ = v65
                            }
                        case "foo bar":
                            var v66 int32 = -2147483648
                            _ = v66
                        case "a":
                            switch 32768 {
                            case 2147483647:
                                v1 = (!(v1 || (!(282 < 1))))
                            case 4294967295:
                                switch "" {
                                case "\n\t":
                                    var v67 uint8 = 7
                                    _ = v67
                                case "foo bar":
                                    for fi68 := 0; fi68 < 6; fi68++ {
                                        var v69 int32 = 312
                                        v70 := GenS2{F0: 65535}
                                        _ = fi68
                                        _ = v69
                                        _ = v70
                                    }
                                case "":
                                    switch "\n\t" {
                                    case "":
                                        if v1 {
                                            if (!v1) {
                                                for fi71 := 0; fi71 < 3; fi71++ {
                                                    switch 1 {
                                                    case -165:
                                                        var v72 float64 = (-2.25 / 1.5)
                                                        _ = v72
                                                    default:
                                                        v1 = (1000 >= -20)
                                                    }
                                                    var v73 int = fi71
                                                    _ = fi71
                                                    _ = v73
                                                }
                                                v1 = (false && v1)
                                                v1 = (v1 || (824 == 255))
                                            } else {
                                                var v74 uint32 = 42
                                                for fi75 := 0; fi75 < 10; fi75++ {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v76 uint64 = 4294967295
                                                        _ = v76
                                                    case "a":
                                                        if (("\n\t" <= ("hello" + (("" + "\n\t") + "foo bar"))) || (!v1)) {
                                                            switch (^(fi75)) {
                                                            case -246:
                                                                var v77 int = fi75
                                                                _ = v77
                                                            case -311:
                                                                fmt.Println(v74)
                                                            case 128:
                                                                for fi78 := 0; fi78 < 2; fi78++ {
                                                                    delete(v2, 42)
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 32767
                                                                    }
                                                                    if (1.5 <= 0.1) {
                                                                        switch (fi78 >> ((65535) & 63)) {
                                                                        case -2147483648:
                                                                            _ = len(v3)
                                                                        case 788:
                                                                            v79 := GenS8{F0: (^(v74)), F1: 0, F2: ""}
                                                                            _ = v79
                                                                        case 726:
                                                                            var v80 uint64 = 1
                                                                            fmt.Println(v74, fi78)
                                                                            var v81 int = (fi75 & fi75)
                                                                            _ = v80
                                                                            _ = v81
                                                                        default:
                                                                            var v82 uint32 = v74
                                                                            _ = v82
                                                                        }
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (fi75 * (fi78 >> ((2) & 63)))
                                                                        }
                                                                    }
                                                                    _ = fi78
                                                                }
                                                            }
                                                            _ = len(v2)
                                                        } else {
                                                            v83 := GenS9{F0: 1, F1: (42 == 42), F2: (140 < 127)}
                                                            v74 -= v74
                                                            _ = v83
                                                        }
                                                    default:
                                                        for fi84 := 0; fi84 < 2; fi84++ {
                                                            if v1 {
                                                                if true {
                                                                    if v1 {
                                                                        var v85 int = 32768
                                                                        v86 := v2[(-(fi84))]
                                                                        var v87 uint8 = 13
                                                                        _ = v85
                                                                        _ = v86
                                                                        _ = v87
                                                                    }
                                                                    switch fi75 {
                                                                    case 128:
                                                                        v88 := GenS9{F0: 0, F1: ("hello" != "hello"), F2: v1}
                                                                        _ = v88
                                                                    case -1:
                                                                        v74 -= v74
                                                                    }
                                                                    var v89 int32 = 0
                                                                    v1 = ((v74 / ((880) | 1)) > (^(v74)))
                                                                    v90 := GenS7{F0: 255}
                                                                    _ = v89
                                                                    _ = v90
                                                                } else {
                                                                    var v91 uint64 = 2
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    fmt.Println(v91, fi84, v74)
                                                                    var v92 uint64 = (v91 + (^(v91)))
                                                                    var v93 int8 = genFn0()
                                                                    switch (fi75 | 650) {
                                                                    case -106:
                                                                        var v94 int16 = 42
                                                                        _ = v94
                                                                    case 4294967295:
                                                                        v95 := GenS5{F0: (v91 ^ 1), F1: (!false), F2: (v93 << ((109) & 63))}
                                                                        _ = v95
                                                                    case 2147483647:
                                                                        switch ((("" + ("" + "\n\t")) + "hello") + "a") {
                                                                        case "a":
                                                                            var v96 uint64 = (^(v92))
                                                                            _ = v96
                                                                        case "":
                                                                            var v97 int64 = -407
                                                                            _ = v97
                                                                        }
                                                                    default:
                                                                        v98 := GenS9{F0: (^(v91)), F1: v1, F2: v1}
                                                                        _ = v98
                                                                    }
                                                                    var v99 uint = 18446744073709551615
                                                                    fmt.Println(v1, v99)
                                                                    _ = v91
                                                                    _ = v92
                                                                    _ = v93
                                                                    _ = v99
                                                                }
                                                                var v100 uint64 = uint64(v74)
                                                                _ = v100
                                                            } else {
                                                                v1 = v1
                                                                v4 = append(v4, 127)
                                                            }
                                                            var v101 int16 = -437
                                                            _ = fi84
                                                            _ = v101
                                                        }
                                                    }
                                                    _ = fi75
                                                }
                                                _ = v74
                                            }
                                            v1 = false
                                            switch "foo bar" {
                                            case "foo bar":
                                                _ = len(v2)
                                            case "hello":
                                                switch 2147483647 {
                                                case -271:
                                                    _ = len(v2)
                                                case 162:
                                                    var v102 int64 = 9223372036854775807
                                                    _ = v102
                                                case -2147483648:
                                                    fmt.Println(v1)
                                                default:
                                                    v1 = false
                                                }
                                            }
                                        } else {
                                            v1 = ("foo bar" == "hello")
                                            if false {
                                                v103 := GenS7{F0: 558}
                                                _ = v103
                                            } else {
                                                if (!true) {
                                                    _ = len(v4)
                                                    v104 := GenS5{F0: 4294967295, F1: false, F2: 100}
                                                    if true {
                                                        v1 = ((float64(3.14159) / (((3.14159 * -2.25) - 3.14159) + (1e10 / 1.0))) == (1.5 / 1.5))
                                                    } else {
                                                        for fi105 := 0; fi105 < 13; fi105++ {
                                                            fmt.Println(v1, v1)
                                                            _ = fi105
                                                        }
                                                        v1 = v1
                                                        v1 = ((1.5 - (((((0.1 * (1e10 * -2.25)) / -2.25) / 1.0) / 1.5) / -2.25)) < -2.25)
                                                    }
                                                    _ = v104
                                                } else {
                                                    v1 = v1
                                                    v1 = (3.14159 < (float64(float32(((1.5 * (((-2.25 - 1.5) + float64(0.0)) + 0.1)) / 3.14159))) + 1e10))
                                                }
                                                v1 = (-588 == 42)
                                                for fi106 := 0; fi106 < 12; fi106++ {
                                                    for fi107 := 0; fi107 < 8; fi107++ {
                                                        var v108 int64 = 42
                                                        var v109 string = ("a" + "foo bar")
                                                        _ = fi107
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                    for fi110 := 0; fi110 < 4; fi110++ {
                                                        if v1 {
                                                            _ = len(v2)
                                                        }
                                                        v1 = ((true || v1) || (!true))
                                                        _ = fi110
                                                    }
                                                    _ = fi106
                                                }
                                            }
                                            var v111 uint16 = 2
                                            _ = v111
                                        }
                                    case "foo bar":
                                        for fi112 := 0; fi112 < 2; fi112++ {
                                            switch "" {
                                            case "":
                                                switch "a" {
                                                case "":
                                                    v1 = v1
                                                case "\n\t":
                                                    v1 = v1
                                                case "hello":
                                                    for fi113 := 0; fi113 < 14; fi113++ {
                                                        var v114 int64 = -708
                                                        fmt.Println(fi112, fi113, v1)
                                                        _ = fi113
                                                        _ = v114
                                                    }
                                                default:
                                                    for fi115 := 0; fi115 < 20; fi115++ {
                                                        var v116 bool = (2 >= 42)
                                                        var v117 uint = 278
                                                        v116 = (2 != 869)
                                                        _ = fi115
                                                        _ = v116
                                                        _ = v117
                                                    }
                                                }
                                            case "foo bar":
                                                if (!(false || v1)) {
                                                    fmt.Println(v1, fi112)
                                                }
                                            case "hello":
                                                if v1 {
                                                    v1 = false
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            default:
                                                var v118 int32 = -2147483648
                                                _ = v118
                                            }
                                            _ = fi112
                                        }
                                    }
                                default:
                                    v1 = v1
                                }
                            case -2147483648:
                                v1 = v1
                            default:
                                for fi119 := 0; fi119 < 2; fi119++ {
                                    v1 = (0 < 1)
                                    for fi120 := 0; fi120 < 16; fi120++ {
                                        v121 := GenS9{F0: 0, F1: ((236 >= 144) || v1), F2: ("a" == "foo bar")}
                                        _ = fi120
                                        _ = v121
                                    }
                                    fmt.Println(v1, fi119)
                                    _ = fi119
                                }
                            }
                        default:
                            switch "a" {
                            case "\n\t":
                                for fi122 := 0; fi122 < 8; fi122++ {
                                    v1 = (!(v1 || (!v1)))
                                    _ = fi122
                                }
                            default:
                                v1 = v1
                            }
                        }
                    } else {
                        v123 := GenS6{F0: 2147483648, F1: 255}
                        var v124 int16 = 207
                        fmt.Println(v1, v1)
                        _ = v123
                        _ = v124
                    }
                }
            } else {
                for fi125 := 0; fi125 < 4; fi125++ {
                    var v126 uint64 = 2
                    for fi127 := 0; fi127 < 11; fi127++ {
                        var v128 int32 = -2147483648
                        fmt.Println(fi127, fi127, v126)
                        _ = fi127
                        _ = v128
                    }
                    for fi129 := 0; fi129 < 1; fi129++ {
                        var v130 uint32 = 2147483647
                        var v131 int = (fi129 << ((0) & 63))
                        v132 := GenS1{F0: 495, F1: fi125}
                        v133 := GenS10{F0: -9223372036854775808, F1: 42}
                        _ = fi129
                        _ = v130
                        _ = v131
                        _ = v132
                        _ = v133
                    }
                    _ = fi125
                    _ = v126
                }
            }
        }
        delete(v2, 9223372036854775807)
    }
    switch "a" {
    case "hello":
        if ((1.5 < 1.5) || (((!v1) && v1) && (-1 > -1))) {
            v1 = false
            var v134 bool = false
            _ = v134
        } else {
            v135 := GenS11{F0: v1, F1: "\n\t"}
            var v136 uint8 = 2
            _ = v135
            _ = v136
        }
    case "foo bar":
        v137 := GenS12{F0: 42}
        _ = v137
    case "\n\t":
        if (255 > 221) {
            v138 := GenS4{F0: 128, F1: 42, F2: 32768}
            var v139 uint64 = 0
            if v1 {
                for fi140 := 0; fi140 < 18; fi140++ {
                    for fi141 := 0; fi141 < 15; fi141++ {
                        v139 -= (^(v139))
                        _ = fi141
                    }
                    _ = fi140
                }
                for fi142 := 0; fi142 < 7; fi142++ {
                    switch fi142 {
                    case -9223372036854775808:
                        for fi143 := 0; fi143 < 20; fi143++ {
                            var v144 uint16 = 42
                            v139 += (v139 / ((v139) | 1))
                            _ = fi143
                            _ = v144
                        }
                    }
                    var v145 uint32 = 1
                    v1 = (9223372036854775807 < (fi142 + fi142))
                    _ = fi142
                    _ = v145
                }
                if v1 {
                    v139 += v139
                }
            } else {
                switch 1 {
                case 0:
                    switch -350 {
                    case -666:
                        var v146 uint64 = (^(v139))
                        _ = v146
                    }
                }
                var v147 uint8 = 42
                _ = v147
            }
            _ = v138
            _ = v139
        }
    }
    switch 128 {
    case -9223372036854775808:
        v1 = v1
    case 9223372036854775807:
        v1 = v1
    default:
        v148 := GenS13{F0: 2, F1: -1, F2: 318}
        _ = v148
    }
    var v149 int = 128
    v1 = (v149 <= (^(v149)))
    v149 = v149
    if false {
        _ = len(v4)
        switch "a" {
        case "hello":
            fmt.Println(v149, v1)
        case "":
            var v150 int64 = 0
            _ = v150
        default:
            for fi151 := 0; fi151 < 19; fi151++ {
                v1 = (1.5 >= (float64(0.1) + 0.1))
                fmt.Println(v1)
                _ = fi151
            }
        }
        var v152 bool = v1
        _ = v152
    }
    for fi153 := 0; fi153 < 4; fi153++ {
        v154 := GenS14{F0: 0, F1: 2147483648, F2: -643}
        fmt.Println(v1, v149, fi153)
        v1 = (128 != 4294967295)
        _ = fi153
        _ = v154
    }
    for fi155 := 0; fi155 < 10; fi155++ {
        v149 = fi155
        var v156 uint64 = 18446744073709551615
        v156 = (v156 % ((v156) | 1))
        _ = fi155
        _ = v156
    }
    for fi157 := 0; fi157 < 5; fi157++ {
        for fi158 := 0; fi158 < 7; fi158++ {
            var v159 float32 = float32((1.5 / 0.1))
            if true {
                v159 -= 0.0
                if v1 {
                    _ = len(v4)
                }
            }
            _ = fi158
            _ = v159
        }
        fmt.Println(fi157)
        _ = fi157
    }
    delete(v2, v149)
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v149
    _ = v2
    _ = v3
    _ = v4
}
