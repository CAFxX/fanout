package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int32; F2 uint64 }
type GenS4 struct { F0 string; F1 uint; F2 string }
type GenS5 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint32; F2 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 uint32; F1 uint }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 uint; F1 int8 }
func genFn0() string {
    if (((true || (-128 == 0)) && (!true)) || true) {
        v6 := GenS1{F0: 521, F1: 42}
        var v7 int = 1
        _ = v6
        _ = v7
    } else {
        var v8 int32 = -1
        _ = v8
    }
    var v9 int32 = 1
    for fi10 := 0; fi10 < 3; fi10++ {
        v9 -= v9
        _ = fi10
    }
    switch 32768 {
    case 2147483647:
        var v11 int32 = v9
        _ = v11
    case 337:
        fmt.Println(v9)
    default:
        if true {
            switch 32768 {
            case 128:
                v9 -= -1
            case 1:
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 int16 = -116
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 uint8 = 128
                        var v16 uint32 = 0
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    _ = fi12
                    _ = v13
                }
            }
        } else {
            if true {
                var v17 uint = 975
                _ = v17
            }
            switch 32767 {
            case -302:
                if false {
                    var v18 int64 = 1
                    switch 172 {
                    case 2147483647:
                        v19 := GenS2{F0: (128 <= 127)}
                        _ = v19
                    case 32768:
                        var v20 string = "a"
                        _ = v20
                    case -105:
                        if true {
                            var v21 uint64 = 1
                            for fi22 := 0; fi22 < 2; fi22++ {
                                for fi23 := 0; fi23 < 19; fi23++ {
                                    v24 := GenS1{F0: (v18 >> ((42) & 63)), F1: int16(v9)}
                                    v21 -= 751
                                    var v25 uint8 = uint8(v21)
                                    var v26 int = fi23
                                    v27 := []int{v26}
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                v28 := GenS3{F0: (fi22 ^ fi22), F1: v9, F2: v21}
                                _ = fi22
                                _ = v28
                            }
                            _ = v21
                        }
                    default:
                        v29 := GenS3{F0: 127, F1: -9, F2: 2}
                        _ = v29
                    }
                    _ = v18
                } else {
                    v30 := []int{-196, 341, -157}
                    _ = v30
                }
            case -2147483648:
                fmt.Println(v9)
            }
            v31 := GenS4{F0: "\n\t", F1: 0, F2: ("foo bar" + "hello")}
            _ = v31
        }
    }
    var v32 uint64 = 42
    _ = v32
    _ = v9
    return ("hello" + "hello")
}
func genFn1(p0 int) string {
    v33 := GenS4{F0: "hello", F1: 906, F2: ("foo bar" + ("" + "\n\t"))}
    for fi34 := 0; fi34 < 20; fi34++ {
        switch ("" + "hello") {
        case "":
            p0 += (p0 & fi34)
        default:
            v35 := GenS5{F0: 42, F1: 42, F2: 1}
            _ = v35
        }
        _ = fi34
    }
    fmt.Println(p0)
    for fi36 := 0; fi36 < 6; fi36++ {
        p0 = (fi36 % ((p0) | 1))
        v37 := []int{}
        var v38 int16 = -324
        _ = fi36
        _ = v37
        _ = v38
    }
    _ = p0
    _ = v33
    return "a"
}

func main() {
    v1 := []int{-9223372036854775808}
    var v2 uint32 = 2147483647
    var v3 uint16 = 255
    var v4 int16 = 42
    var v5 uint16 = 65535
    for fi39 := 0; fi39 < 20; fi39++ {
        for fi40 := 0; fi40 < 9; fi40++ {
            var v41 float32 = -2.25
            v4 = v4
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    v42 := GenS2{F0: false}
    if ((true || true) || false) {
        for fi43 := 0; fi43 < 5; fi43++ {
            var v44 uint32 = (v2 << ((255) & 63))
            _ = fi43
            _ = v44
        }
    } else {
        var v45 string = ("foo bar" + ("\n\t" + "hello"))
        var v46 uint64 = 2
        _ = v45
        _ = v46
    }
    fmt.Println(v3, v4)
    var v47 int8 = -128
    if false {
        v3 = v5
    } else {
        var v48 int64 = 846
        _ = v48
    }
    if false {
        if (128 < int(v5)) {
            if (-2.25 != 1e10) {
                v2 -= (v2 + 4294967295)
                for fi49 := 0; fi49 < 2; fi49++ {
                    switch fi49 {
                    case 32768:
                        switch "hello" {
                        case "":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "foo bar":
                            v50 := GenS3{F0: fi49, F1: 0, F2: 737}
                            _ = v50
                        case "a":
                            fmt.Println(fi49, v3)
                        default:
                            if false {
                                v2 += (v2 & (^(v2)))
                                fmt.Println(v3)
                            } else {
                                var v51 int64 = 9223372036854775807
                                switch ("hello" + "hello") {
                                case "hello":
                                    if (("\n\t" == (("a" + "") + "hello")) && true) {
                                        if true {
                                            fmt.Println(v47, fi49)
                                        } else {
                                            if (!((v4 << ((2) & 63)) < int16(v5))) {
                                                v52 := GenS6{F0: v3, F1: 2147483647, F2: ""}
                                                _ = v52
                                            } else {
                                                var v53 uint32 = (^(v2))
                                                for fi54 := 0; fi54 < 15; fi54++ {
                                                    var v55 string = ""
                                                    _ = fi54
                                                    _ = v55
                                                }
                                                if ((fi49 - (-(fi49))) != -681) {
                                                    for fi56 := 0; fi56 < 2; fi56++ {
                                                        if (2 <= 962) {
                                                            v1 = append(v1, 42)
                                                            if (!true) {
                                                                v57 := GenS7{F0: 18446744073709551615}
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v58 := GenS2{F0: false}
                                                                var v59 int = fi49
                                                                v60 := GenS1{F0: 42, F1: (-(v4))}
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                        v51 -= v51
                                                        _ = fi56
                                                    }
                                                } else {
                                                    v3 = v5
                                                    v1 = append(v1, fi49)
                                                    v51 = v51
                                                }
                                                _ = v53
                                            }
                                            v2 = (v2 >> ((4294967295) & 63))
                                            if ((42 < 1) && (fi49 > -2147483648)) {
                                                var v61 uint32 = v2
                                                _ = v61
                                            } else {
                                                if false {
                                                    var v62 uint16 = (^(v5))
                                                    var v63 int64 = v51
                                                    _ = v62
                                                    _ = v63
                                                }
                                            }
                                        }
                                    }
                                case "foo bar":
                                    switch "" {
                                    case "hello":
                                        _ = len(v1)
                                    case "":
                                        for fi64 := 0; fi64 < 16; fi64++ {
                                            if len(v1) > 0 {
                                                v1[0] = fi64
                                            }
                                            v51 += v51
                                            v3 -= (^(v5))
                                            _ = fi64
                                        }
                                    default:
                                        for fi65 := 0; fi65 < 17; fi65++ {
                                            switch "" {
                                            case "hello":
                                                v2 -= (^(v2))
                                            default:
                                                v4 += 84
                                            }
                                            var v66 int8 = v47
                                            v67 := GenS1{F0: -9223372036854775808, F1: -1}
                                            var v68 uint64 = 1
                                            _ = fi65
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        }
                                    }
                                case "\n\t":
                                    if len(v1) > 0 {
                                        v1[0] = (fi49 - 2147483647)
                                    }
                                default:
                                    fmt.Println(v4, fi49)
                                }
                                _ = v51
                            }
                        }
                    case 1:
                        switch "\n\t" {
                        case "":
                            v47 -= (v47 - -85)
                        case "foo bar":
                            v3 = (v3 + uint16(v2))
                        }
                    }
                    switch "a" {
                    case "foo bar":
                        if false {
                            var v69 int = fi49
                            switch (fi49 * fi49) {
                            case 0:
                                v70 := GenS3{F0: 32767, F1: int32(v2), F2: 520}
                                _ = v70
                            case -943:
                                for fi71 := 0; fi71 < 10; fi71++ {
                                    v3 = v5
                                    _ = fi71
                                }
                            case 9223372036854775807:
                                var v72 float32 = ((0.0 * 1.5) * 0.0)
                                _ = v72
                            default:
                                v73 := GenS8{F0: 45, F1: "foo bar"}
                                _ = v73
                            }
                            var v74 uint = 65535
                            _ = v69
                            _ = v74
                        } else {
                            fmt.Println(v5)
                            v75 := GenS6{F0: v5, F1: 65535, F2: ("\n\t" + ("a" + "foo bar"))}
                            v3 = (v5 % ((273) | 1))
                            _ = v75
                        }
                    case "":
                        for fi76 := 0; fi76 < 18; fi76++ {
                            v77 := GenS9{F0: v2, F1: 255}
                            v1 = append(v1, (fi76 / ((503) | 1)))
                            _ = fi76
                            _ = v77
                        }
                    }
                    if false {
                        if len(v1) > 0 {
                            v1[0] = fi49
                        }
                        for fi78 := 0; fi78 < 18; fi78++ {
                            if false {
                                fmt.Println(fi49, v47)
                                _ = len(v1)
                                if len(v1) > 0 {
                                    v1[0] = int(v2)
                                }
                                if (2 <= 18446744073709551615) {
                                    fmt.Println(v2, fi78, fi78)
                                    var v79 float64 = (1e10 * (1.5 + (0.1 * 0.1)))
                                    if len(v1) > 0 {
                                        v1[0] = -2147483648
                                    }
                                    _ = v79
                                } else {
                                    switch (fi49 >> ((255) & 63)) {
                                    case 4294967295:
                                        v80 := GenS3{F0: fi78, F1: 0, F2: 918}
                                        _ = v80
                                    case 746:
                                        var v81 uint16 = 0
                                        _ = v81
                                    case -686:
                                        var v82 int16 = (v4 ^ v4)
                                        var v83 string = ("foo bar" + ("a" + ""))
                                        _ = v82
                                        _ = v83
                                    default:
                                        v1 = append(v1, fi78)
                                        v2 -= (v2 % ((1) | 1))
                                        v47 = 127
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            v2 = 2147483647
                            _ = fi78
                        }
                        if false {
                            if ((3.14159 > (1e10 * ((((0.0 / 3.14159) / 1.0) * -2.25) / 0.1))) && false) {
                                _ = len(v1)
                                v84 := GenS1{F0: 42, F1: v4}
                                _ = v84
                            }
                            switch "foo bar" {
                            case "a":
                                switch (fi49 ^ fi49) {
                                case 584:
                                    var v85 int32 = 2147483647
                                    _ = v85
                                }
                            case "foo bar":
                                _ = len(v1)
                            }
                        }
                    } else {
                        for fi86 := 0; fi86 < 13; fi86++ {
                            v87 := GenS5{F0: 1, F1: (v2 << ((0) & 63)), F2: 0}
                            var v88 float64 = 3.14159
                            if true {
                                v2 = (v2 | (v2 / ((332) | 1)))
                            }
                            switch "" {
                            case "hello":
                                var v89 uint16 = (^(v3))
                                _ = v89
                            case "":
                                fmt.Println(v88)
                            default:
                                switch "a" {
                                case "foo bar":
                                    var v90 int32 = 42
                                    _ = v90
                                }
                            }
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                        for fi91 := 0; fi91 < 4; fi91++ {
                            v5 += v5
                            switch 127 {
                            case -1:
                                v92 := GenS6{F0: v5, F1: 899, F2: "foo bar"}
                                _ = v92
                            case 915:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case -9223372036854775808:
                                if false {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483647
                                    }
                                    var v93 uint = 116
                                    _ = v93
                                }
                            }
                            _ = fi91
                        }
                    }
                    _ = fi49
                }
            }
            v94 := GenS8{F0: -1, F1: "hello"}
            if len(v1) > 0 {
                v1[0] = 691
            }
            _ = v94
        } else {
            v95 := GenS1{F0: -1, F1: (v4 >> ((18446744073709551615) & 63))}
            _ = v95
        }
        v96 := GenS4{F0: "", F1: 0, F2: "a"}
        v97 := GenS10{F0: true, F1: 42}
        _ = v96
        _ = v97
    } else {
        v98 := GenS11{F0: 4294967295, F1: (v47 - (v47 << ((4294967295) & 63)))}
        _ = v98
    }
    v99 := GenS5{F0: -2147483648, F1: v2, F2: 4294967295}
    fmt.Println(v4, v4)
    v5 += v5
    for fi100 := 0; fi100 < 2; fi100++ {
        for fi101 := 0; fi101 < 18; fi101++ {
            if len(v1) > 0 {
                v1[0] = fi100
            }
            var v102 uint32 = (^(v2))
            _ = fi101
            _ = v102
        }
        var v103 uint32 = (v2 * v2)
        _ = fi100
        _ = v103
    }
    switch -995 {
    case 26:
        v2 = v2
    case -2147483648:
        if (v4 >= 86) {
            switch ("\n\t" + ("\n\t" + "hello")) {
            case "hello":
                v2 = 2147483647
            case "":
                var v104 uint64 = 891
                _ = v104
            case "foo bar":
                v105 := GenS5{F0: 42, F1: v2, F2: 18446744073709551615}
                _ = v105
            }
            v106 := GenS3{F0: 32768, F1: 919, F2: 1}
            _ = v106
        }
    case 127:
        if (v47 == 29) {
            if (603 >= 65535) {
                if ((1.5 / 3.14159) < float64(1e10)) {
                    var v107 uint = 42
                    switch 32768 {
                    case -9223372036854775808:
                        for fi108 := 0; fi108 < 18; fi108++ {
                            var v109 float32 = float32(0.1)
                            v110 := GenS5{F0: 2147483647, F1: (^(v2)), F2: 42}
                            _ = fi108
                            _ = v109
                            _ = v110
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = v107
                } else {
                    var v111 bool = ((v2 | (v2 % ((2147483648) | 1))) == 1)
                    var v112 int8 = v47
                    _ = v111
                    _ = v112
                }
            }
        }
    }
    if ((!true) && (!false)) {
        for fi113 := 0; fi113 < 1; fi113++ {
            if false {
                var v114 uint = 260
                _ = v114
            } else {
                for fi115 := 0; fi115 < 3; fi115++ {
                    if false {
                        if false {
                            var v116 int = fi113
                            _ = v116
                        }
                        v4 = (v4 + -751)
                        if true {
                            v1 = append(v1, (fi115 >> ((255) & 63)))
                            v117 := GenS11{F0: 4294967295, F1: v47}
                            v118 := GenS8{F0: v47, F1: ""}
                            _ = v117
                            _ = v118
                        }
                    }
                    v119 := GenS5{F0: 0, F1: (v2 - 436), F2: 1}
                    if false {
                        v4 -= v4
                    } else {
                        var v120 bool = true
                        _ = v120
                    }
                    _ = fi115
                    _ = v119
                }
            }
            v121 := GenS4{F0: ("hello" + "\n\t"), F1: 760, F2: (("a" + ("foo bar" + "hello")) + "")}
            for fi122 := 0; fi122 < 16; fi122++ {
                v3 -= v5
                var v123 int64 = -1
                var v124 string = "foo bar"
                _ = fi122
                _ = v123
                _ = v124
            }
            _ = fi113
            _ = v121
        }
    } else {
        var v125 uint16 = 0
        for fi126 := 0; fi126 < 8; fi126++ {
            var v127 uint = 42
            v2 = (v2 - (v2 << ((v127) & 63)))
            _ = fi126
            _ = v127
        }
        var v128 float64 = 3.14159
        _ = v125
        _ = v128
    }
    if false {
        if true {
            var v129 float64 = (1e10 * (3.14159 - 0.0))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v129
        } else {
            _ = len(v1)
            fmt.Println(v4, v4, v4)
        }
    } else {
        if true {
            fmt.Println(v5, v5, v5)
            var v130 int16 = v4
            _ = v130
        }
        v4 = 32767
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v99
}
