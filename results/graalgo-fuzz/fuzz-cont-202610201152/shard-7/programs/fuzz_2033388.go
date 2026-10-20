package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 uint32; F1 bool; F2 int8 }
type GenS11 struct { F0 int16; F1 uint64 }
type GenS12 struct { F0 int8; F1 uint; F2 bool }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 string; F1 uint8; F2 uint8 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int16; F1 int32 }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int32; F1 uint; F2 int16 }
type GenS20 struct { F0 string; F1 int16; F2 string }
type GenS21 struct { F0 int32; F1 uint32 }
type GenS22 struct { F0 string; F1 uint }

func main() {
    v1 := []int{}
    var v2 int32 = 42
    var v3 string = "hello"
    var v4 float32 = 0.0
    var v5 int16 = 42
    var v6 int16 = (v5 + (^(v5)))
    v7 := GenS1{F0: 8}
    v8 := GenS2{F0: 30}
    var v9 int32 = (v2 / (((v2 % (((v2 / ((484) | 1))) | 1))) | 1))
    v10 := GenS3{F0: 1, F1: 1}
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: 0}
        v13 := GenS4{F0: (true || (-63 > 120)), F1: fi11}
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(v9, v6, v3)
        fmt.Println(v3)
        _ = fi14
    }
    v15 := GenS5{F0: -1, F1: v3}
    if ((2147483647 != v9) || true) {
        if ("foo bar" <= (v3 + "foo bar")) {
            for fi16 := 0; fi16 < 3; fi16++ {
                if (v3 != (("" + ("\n\t" + "foo bar")) + "hello")) {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        _ = len(v1)
                        fmt.Println(v5, v3, fi16)
                        if false {
                            var v18 int8 = 1
                            v1 = append(v1, (fi17 + (fi17 | 2147483647)))
                            _ = v18
                        } else {
                            fmt.Println(v5)
                            if (!false) {
                                switch v3 {
                                case "foo bar":
                                    if true {
                                        v19 := GenS1{F0: -26}
                                        if (954 != 2147483647) {
                                            var v20 int32 = v9
                                            _ = v20
                                        } else {
                                            var v21 int64 = -9223372036854775808
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v22 bool = false
                                            _ = v21
                                            _ = v22
                                        }
                                        _ = v19
                                    } else {
                                        v2 += (v9 >> ((0) & 63))
                                        switch fi16 {
                                        case 357:
                                            if true {
                                                var v23 int32 = v9
                                                _ = v23
                                            }
                                        case 0:
                                            if (!true) {
                                                v24 := GenS6{F0: 216, F1: (-(v2))}
                                                v6 = v5
                                                var v25 int8 = -19
                                                v26 := GenS2{F0: (v25 - 42)}
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v1 = append(v1, (-(fi17)))
                                                if false {
                                                    if false {
                                                        if (1 > 127) {
                                                            if len(v1) > 0 {
                                                                v1[0] = fi17
                                                            }
                                                            var v27 uint8 = 2
                                                            v3 = ((v3 + v3) + "a")
                                                            v2 = v2
                                                            _ = v27
                                                        } else {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                        v3 = v3
                                                        if (false || ((((("foo bar" + (v3 + (((v3 + "") + v3) + (v3 + ((("\n\t" + (((v3 + (v3 + "foo bar")) + v3) + (v3 + "foo bar"))) + v3) + v3))))) + ("" + ("\n\t" + "foo bar"))) + "a") + (v3 + (v3 + ("" + "a")))) > ("" + v3))) {
                                                            v28 := GenS7{F0: 255, F1: (-2.25 >= -2.25)}
                                                            var v29 int8 = 19
                                                            _ = v28
                                                            _ = v29
                                                        } else {
                                                            if true {
                                                                if (4294967295 == 2147483648) {
                                                                    v3 = (v3 + (v3 + (("\n\t" + (v3 + "a")) + v3)))
                                                                    var v30 bool = (35 == 331)
                                                                    if ((^(v9)) != 42) {
                                                                        v6 += (v6 / (((v6 % ((-426) | 1))) | 1))
                                                                    } else {
                                                                        var v31 int = -219
                                                                        var v32 uint8 = 128
                                                                        _ = v31
                                                                        _ = v32
                                                                    }
                                                                    _ = v30
                                                                }
                                                                if false {
                                                                    switch 156 {
                                                                    case 128:
                                                                        v9 -= v2
                                                                    default:
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                    var v33 float32 = float32(0.1)
                                                                    _ = v33
                                                                }
                                                            } else {
                                                                switch v3 {
                                                                case "foo bar":
                                                                    v9 += v2
                                                                default:
                                                                    _ = len(v1)
                                                                }
                                                                var v34 uint32 = 65535
                                                                _ = v34
                                                            }
                                                            var v35 float64 = (0.1 * 3.14159)
                                                            v5 -= 923
                                                            _ = v35
                                                        }
                                                    } else {
                                                        switch fi17 {
                                                        case -9223372036854775808:
                                                            v5 = (v5 % (((v5 << ((0) & 63))) | 1))
                                                        case -242:
                                                            v4 = 3.14159
                                                            v9 = -280
                                                        }
                                                        v1 = append(v1, fi17)
                                                    }
                                                    var v36 uint8 = 105
                                                    _ = v36
                                                }
                                                if true {
                                                    var v37 uint64 = 4294967295
                                                    _ = v37
                                                }
                                            }
                                        case 4294967295:
                                            var v38 float64 = 1e10
                                            _ = v38
                                        }
                                        var v39 int64 = 9223372036854775807
                                        _ = v39
                                    }
                                case "a":
                                    var v40 int16 = -744
                                    _ = v40
                                case "":
                                    if (((((("\n\t" + v3) + (("" + v3) + v3)) + "\n\t") + (v3 + (v3 + (v3 + "foo bar")))) + "hello") != v3) {
                                        v3 = "hello"
                                        var v41 uint16 = 65535
                                        _ = v41
                                    } else {
                                        if false {
                                            v4 -= ((1e10 / (3.14159 * float32((1e10 - 1.5)))) + 0.0)
                                            var v42 string = ("hello" + v3)
                                            switch v42 {
                                            case "hello":
                                                v42 = v3
                                            case "\n\t":
                                                var v43 int16 = (^(v5))
                                                _ = v43
                                            default:
                                                fmt.Println(fi16, v9)
                                            }
                                            _ = v42
                                        }
                                    }
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v6, fi17, fi16)
                            } else {
                                v44 := GenS6{F0: 128, F1: v2}
                                _ = v44
                            }
                        }
                        _ = fi17
                    }
                    v45 := GenS8{F0: ((true || (757 > 1)) || (9223372036854775807 == 1)), F1: 279, F2: 2}
                    _ = v45
                } else {
                    fmt.Println(v4, v3, v3)
                    var v46 float32 = (v4 - v4)
                    _ = v46
                }
                v5 += v5
                if ((v4 / ((v4 - v4) * 3.14159)) > (-2.25 + v4)) {
                    if (!(0.0 != (0.0 + float64(v4)))) {
                        for fi47 := 0; fi47 < 7; fi47++ {
                            var v48 bool = true
                            _ = fi47
                            _ = v48
                        }
                        v6 = v6
                        v3 = v3
                    }
                } else {
                    for fi49 := 0; fi49 < 2; fi49++ {
                        v50 := GenS9{F0: 42, F1: -23}
                        if true {
                            var v51 uint64 = 1
                            switch v3 {
                            case "foo bar":
                                switch fi49 {
                                case -513:
                                    fmt.Println(v9, fi49)
                                default:
                                    if true {
                                        fmt.Println(v6)
                                        switch (v3 + v3) {
                                        case "\n\t":
                                            if (!(!false)) {
                                                if false {
                                                    if (!(910 > 42)) {
                                                        var v52 float64 = 1e10
                                                        _ = v52
                                                    }
                                                } else {
                                                    v4 = float32(1.5)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                        case "a":
                                            switch (fi16 + fi49) {
                                            case 9223372036854775807:
                                                _ = len(v1)
                                                var v53 float64 = -2.25
                                                switch "a" {
                                                case "\n\t":
                                                    if ((!(-566 >= fi49)) && false) {
                                                        v2 += (v9 % ((v9) | 1))
                                                    } else {
                                                        if false {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v54 := GenS3{F0: v2, F1: (-(v2))}
                                                                _ = v54
                                                            }
                                                        } else {
                                                            v55 := GenS1{F0: 1}
                                                            switch ((v3 + v3) + "foo bar") {
                                                            case "":
                                                                if (!true) {
                                                                    v4 -= (v4 * ((float32(v53) - -2.25) + v4))
                                                                } else {
                                                                    var v56 uint32 = 592
                                                                    var v57 int8 = -9
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = 32767
                                                            }
                                                            _ = v55
                                                        }
                                                        switch ((("a" + (((v3 + v3) + ((v3 + (v3 + "\n\t")) + "hello")) + (v3 + v3))) + ((v3 + v3) + (v3 + ("\n\t" + "foo bar")))) + ((((v3 + (v3 + "a")) + (v3 + v3)) + v3) + v3)) {
                                                        case "":
                                                            v58 := GenS6{F0: 2, F1: 0}
                                                            _ = v58
                                                        case "hello":
                                                            fmt.Println(v2, v2)
                                                        case "foo bar":
                                                            _ = len(v1)
                                                        }
                                                        fmt.Println(v51)
                                                        var v59 int = fi16
                                                        _ = v59
                                                    }
                                                default:
                                                    v60 := GenS8{F0: (!(42 < uint16(v51))), F1: 902, F2: v51}
                                                    _ = v60
                                                }
                                                var v61 uint32 = 2147483647
                                                _ = len(v1)
                                                _ = v53
                                                _ = v61
                                            default:
                                                v9 = v2
                                                fmt.Println(fi16, v5)
                                                var v62 int8 = -128
                                                v9 = -693
                                                _ = v62
                                            }
                                            v2 += -2147483648
                                        }
                                        if true {
                                            switch 1 {
                                            case 32767:
                                                var v63 int16 = v5
                                                _ = v63
                                            case 9223372036854775807:
                                                var v64 int = fi16
                                                _ = v64
                                            case 128:
                                                switch (v3 + ("a" + v3)) {
                                                case "\n\t":
                                                    switch (fi16 + -2147483648) {
                                                    case -2147483648:
                                                        if (true && false) {
                                                            switch (^(fi49)) {
                                                            case 9223372036854775807:
                                                                var v65 int16 = (^(v5))
                                                                _ = v65
                                                            case 2147483647:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            default:
                                                                if true {
                                                                    v6 += -627
                                                                    var v66 uint64 = (^(v51))
                                                                    var v67 uint8 = 128
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                            }
                                                        } else {
                                                            v1 = append(v1, fi16)
                                                        }
                                                    default:
                                                        var v68 bool = true
                                                        _ = v68
                                                    }
                                                }
                                            default:
                                                fmt.Println(v6)
                                            }
                                            if false {
                                                if true {
                                                    v5 = (v6 - v5)
                                                    fmt.Println(v51, fi16, v9)
                                                    v6 = v6
                                                }
                                                v5 = (^(v6))
                                            } else {
                                                fmt.Println(v2)
                                                var v69 uint8 = 128
                                                v9 = (v2 | -1)
                                                _ = v69
                                            }
                                        } else {
                                            v51 = (^(v51))
                                        }
                                    } else {
                                        var v70 uint64 = (^(v51))
                                        v71 := GenS1{F0: 127}
                                        v51 -= 315
                                        if ((638 < 2) && (0 > 285)) {
                                            if ((42 <= 2) && false) {
                                                v4 += -2.25
                                            } else {
                                                var v72 float32 = 0.0
                                                var v73 int32 = v9
                                                _ = v72
                                                _ = v73
                                            }
                                            v74 := GenS2{F0: -128}
                                            _ = v74
                                        } else {
                                            var v75 string = v3
                                            _ = v75
                                        }
                                        _ = v70
                                        _ = v71
                                    }
                                }
                            case "":
                                v76 := GenS3{F0: -881, F1: (v9 | v9)}
                                _ = v76
                            }
                            _ = v51
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                var v77 uint64 = 682
                                _ = v77
                            case "hello":
                                if ((true && false) && ((1.5 / 0.1) > 1.5)) {
                                    fmt.Println(fi16, v3)
                                    v1 = append(v1, (fi49 % ((fi49) | 1)))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    var v78 float32 = v4
                                    if true {
                                        var v79 uint8 = 121
                                        _ = v79
                                    }
                                    v80 := GenS9{F0: 0, F1: 127}
                                    _ = v78
                                    _ = v80
                                }
                            case "":
                                var v81 uint64 = 18446744073709551615
                                _ = v81
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            fmt.Println(v3)
                            _ = len(v1)
                        }
                        var v82 uint = 4294967295
                        var v83 float32 = (v4 * v4)
                        var v84 uint16 = 255
                        _ = fi49
                        _ = v50
                        _ = v82
                        _ = v83
                        _ = v84
                    }
                    switch fi16 {
                    case 2147483647:
                        v9 += 2147483647
                    default:
                        v9 += v2
                    }
                    v6 = (v6 >> ((4294967295) & 63))
                }
                _ = fi16
            }
            var v85 uint16 = 0
            _ = v85
        }
        var v86 uint64 = 2
        fmt.Println(v5, v4, v3)
        _ = v86
    } else {
        switch 1 {
        case 128:
            switch ("\n\t" + v3) {
            case "a":
                if (255 >= 42) {
                    switch -9223372036854775808 {
                    case -2147483648:
                        v3 = v3
                    default:
                        for fi87 := 0; fi87 < 9; fi87++ {
                            if true {
                                fmt.Println(v2, v4, v9)
                                v88 := GenS10{F0: 2147483647, F1: true, F2: 1}
                                for fi89 := 0; fi89 < 17; fi89++ {
                                    var v90 float64 = 0.1
                                    var v91 int8 = 1
                                    _ = fi89
                                    _ = v90
                                    _ = v91
                                }
                                _ = v88
                            }
                            v92 := GenS2{F0: -63}
                            _ = fi87
                            _ = v92
                        }
                    }
                    v93 := GenS2{F0: 127}
                    for fi94 := 0; fi94 < 3; fi94++ {
                        if true {
                            v95 := GenS11{F0: v5, F1: 2}
                            v2 += v9
                            fmt.Println(v4)
                            _ = v95
                        } else {
                            fmt.Println(v3, v9, v3)
                            var v96 uint8 = 2
                            _ = v96
                        }
                        if false {
                            switch (v3 + v3) {
                            case "a":
                                var v97 uint64 = 731
                                _ = v97
                            case "hello":
                                for fi98 := 0; fi98 < 17; fi98++ {
                                    switch ("a" + v3) {
                                    case "foo bar":
                                        v99 := GenS6{F0: 1, F1: v9}
                                        _ = v99
                                    default:
                                        if (32767 < 214) {
                                            if ((0 == 2147483648) && false) {
                                                v100 := GenS3{F0: (v9 * 1), F1: 2147483647}
                                                if (false && true) {
                                                    v4 = v4
                                                } else {
                                                    fmt.Println(v4, fi94, v4)
                                                    var v101 float32 = v4
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v101
                                                }
                                                v102 := GenS5{F0: 0, F1: (v3 + v3)}
                                                _ = v100
                                                _ = v102
                                            }
                                            v4 += v4
                                            var v103 string = v3
                                            v4 -= v4
                                            if true {
                                                var v104 uint64 = 468
                                                v105 := GenS12{F0: 42, F1: 42, F2: false}
                                                _ = v104
                                                _ = v105
                                            } else {
                                                if false {
                                                    v6 -= (-(v5))
                                                    var v106 int16 = (v6 << ((42) & 63))
                                                    if false {
                                                        v107 := GenS13{F0: 1}
                                                        v108 := GenS10{F0: 2147483648, F1: ("" <= (v3 + (v103 + "foo bar"))), F2: -128}
                                                        _ = len(v1)
                                                        var v109 string = v3
                                                        _ = v107
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v110 int64 = -687
                                                        _ = v110
                                                    default:
                                                        fmt.Println(fi98, v103)
                                                    }
                                                    _ = v106
                                                }
                                            }
                                            _ = len(v1)
                                            _ = v103
                                        }
                                    }
                                    _ = fi98
                                }
                            }
                            v111 := GenS14{F0: 155}
                            _ = v111
                        }
                        _ = fi94
                    }
                    _ = v93
                } else {
                    fmt.Println(v4, v5, v3)
                    v1 = append(v1, 1)
                }
            case "hello":
                v4 = 0.0
            default:
                for fi112 := 0; fi112 < 1; fi112++ {
                    switch v3 {
                    case "":
                        for fi113 := 0; fi113 < 16; fi113++ {
                            switch v3 {
                            case "\n\t":
                                v4 = ((float32(3.14159) / 1.0) + (((v4 + 1e10) / 1e10) / v4))
                            default:
                                v114 := GenS9{F0: 128, F1: 127}
                                _ = v114
                            }
                            if true {
                                var v115 bool = (!true)
                                _ = v115
                            } else {
                                v116 := GenS15{F0: ((v3 + "") + v3), F1: 2, F2: 1}
                                v5 -= v5
                                if (317 <= 42) {
                                    fmt.Println(v6, v2)
                                    fmt.Println(fi113, v4)
                                    v4 += (0.0 - (1.5 + v4))
                                } else {
                                    var v117 float64 = 1e10
                                    _ = len(v1)
                                    _ = v117
                                }
                                _ = v116
                            }
                            _ = fi113
                        }
                    case "a":
                        var v118 int8 = 70
                        _ = v118
                    default:
                        v9 -= int32(v5)
                    }
                    if false {
                        if false {
                            v3 = ("hello" + (v3 + (v3 + v3)))
                            v3 = (v3 + v3)
                            switch (fi112 & fi112) {
                            case 32768:
                                switch v3 {
                                case "":
                                    v2 = v9
                                }
                            }
                        }
                        if ((^(fi112)) != fi112) {
                            v6 = -378
                        }
                    } else {
                        var v119 int8 = -128
                        v3 = (v3 + v3)
                        _ = v119
                    }
                    var v120 uint = 9
                    _ = fi112
                    _ = v120
                }
            }
        default:
            v121 := GenS16{F0: 0}
            _ = v121
        }
        fmt.Println(v9, v9)
    }
    v5 += v6
    for fi122 := 0; fi122 < 9; fi122++ {
        var v123 int16 = 13
        v124 := GenS17{F0: v123, F1: (v9 - (-(v9)))}
        switch fi122 {
        case -282:
            v5 -= (v123 | v123)
        case 32768:
            for fi125 := 0; fi125 < 2; fi125++ {
                v5 -= v123
                _ = fi125
            }
        case 344:
            v2 = v9
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi122
        _ = v123
        _ = v124
    }
    var v126 string = (v3 + "")
    if ((1.5 * v4) == 0.1) {
        fmt.Println(v6)
    } else {
        if (148 != 255) {
            for fi127 := 0; fi127 < 4; fi127++ {
                switch v126 {
                case "hello":
                    if (false || (0.0 != 0.0)) {
                        for fi128 := 0; fi128 < 19; fi128++ {
                            switch fi127 {
                            case 32767:
                                var v129 uint64 = 42
                                _ = v129
                            }
                            _ = fi128
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 127
                        }
                        if false {
                            if true {
                                var v130 string = ("\n\t" + v3)
                                v131 := GenS18{F0: (v2 & 1)}
                                if len(v1) > 0 {
                                    v1[0] = -325
                                }
                                _ = v130
                                _ = v131
                            }
                        } else {
                            if ((((v126 + v3) + v126) + v3) != v3) {
                                var v132 int8 = -39
                                _ = v132
                            } else {
                                if false {
                                    for fi133 := 0; fi133 < 8; fi133++ {
                                        _ = len(v1)
                                        v126 = v3
                                        if (4294967295 != 18446744073709551615) {
                                            v5 -= v5
                                            v134 := GenS4{F0: false, F1: 9223372036854775807}
                                            if true {
                                                var v135 uint16 = 42
                                                v136 := GenS1{F0: 100}
                                                _ = v135
                                                _ = v136
                                            } else {
                                                fmt.Println(v5)
                                                v126 = ("\n\t" + (v126 + "hello"))
                                            }
                                            v126 = ("foo bar" + ("hello" + ("foo bar" + (v3 + (v126 + (v3 + "a"))))))
                                            _ = v134
                                        } else {
                                            _ = len(v1)
                                        }
                                        _ = fi133
                                    }
                                }
                            }
                            for fi137 := 0; fi137 < 8; fi137++ {
                                v1 = append(v1, 9223372036854775807)
                                var v138 uint32 = 42
                                _ = fi137
                                _ = v138
                            }
                            switch (v3 + (((("\n\t" + v126) + ((v3 + v3) + v126)) + "a") + "")) {
                            case "a":
                                var v139 int8 = 100
                                _ = v139
                            case "foo bar":
                                v140 := GenS19{F0: -1, F1: 668, F2: (v5 + v6)}
                                _ = v140
                            default:
                                v6 = v5
                            }
                        }
                        v141 := GenS6{F0: 128, F1: -811}
                        _ = v141
                    }
                case "\n\t":
                    v142 := GenS20{F0: v3, F1: v6, F2: "foo bar"}
                    _ = v142
                case "":
                    for fi143 := 0; fi143 < 10; fi143++ {
                        var v144 uint16 = 255
                        if (v144 == v144) {
                            if (545 != 65535) {
                                if false {
                                    v9 = v2
                                } else {
                                    if (v5 != (^(v5))) {
                                        var v145 uint8 = 239
                                        var v146 uint32 = 1
                                        var v147 uint8 = 127
                                        _ = v145
                                        _ = v146
                                        _ = v147
                                    }
                                }
                            }
                            v4 += v4
                        }
                        _ = fi143
                        _ = v144
                    }
                }
                var v148 int16 = v5
                v3 = ""
                _ = fi127
                _ = v148
            }
            v5 += 163
        } else {
            var v149 bool = false
            for fi150 := 0; fi150 < 10; fi150++ {
                for fi151 := 0; fi151 < 20; fi151++ {
                    var v152 int = 32767
                    _ = fi151
                    _ = v152
                }
                var v153 bool = ("foo bar" >= v3)
                v154 := GenS19{F0: v2, F1: 0, F2: 32767}
                _ = fi150
                _ = v153
                _ = v154
            }
            _ = v149
        }
    }
    for fi155 := 0; fi155 < 1; fi155++ {
        v4 = float32(float64(((1.5 + v4) + (0.1 + (0.0 * 0.1)))))
        var v156 uint16 = 0
        for fi157 := 0; fi157 < 17; fi157++ {
            if (fi155 <= (fi155 >> ((921) & 63))) {
                if len(v1) > 0 {
                    v1[0] = (fi157 & 0)
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v4 -= ((v4 * ((-2.25 / float32(-2.25)) / 0.1)) / (-2.25 / 1.5))
                _ = len(v1)
            }
            switch -563 {
            case -454:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 32768:
                var v158 uint16 = v156
                _ = v158
            case 125:
                v9 += (-(v9))
            }
            _ = fi157
        }
        _ = fi155
        _ = v156
    }
    var v159 bool = (!(42 != 100))
    if (109 == 197) {
        v1 = append(v1, int(v9))
    }
    switch (v126 + (("foo bar" + v3) + "\n\t")) {
    case "foo bar":
        if ((float64((3.14159 - (v4 / v4))) <= 0.0) && (57 <= 46)) {
            var v160 uint = 220
            _ = v160
        }
    case "":
        _ = len(v1)
    case "hello":
        if false {
            switch ("\n\t" + v3) {
            case "":
                v161 := GenS8{F0: true, F1: 42, F2: 18446744073709551615}
                _ = v161
            case "a":
                v3 = v3
            case "hello":
                if len(v1) > 0 {
                    v1[0] = 138
                }
            default:
                v1 = append(v1, 127)
            }
            switch -1 {
            case 126:
                v162 := GenS4{F0: (v159 && true), F1: 32768}
                _ = v162
            case 42:
                if v159 {
                    _ = len(v1)
                    switch 128 {
                    case 0:
                        v163 := GenS21{F0: (v9 - (v2 | (-(v9)))), F1: 209}
                        _ = v163
                    case 2147483647:
                        for fi164 := 0; fi164 < 12; fi164++ {
                            v2 = (v2 + (v9 | v9))
                            v4 += (v4 / (((v4 / v4) + v4) + v4))
                            _ = fi164
                        }
                    default:
                        v165 := GenS22{F0: v126, F1: 4294967295}
                        _ = v165
                    }
                }
            }
        }
    default:
        v2 = (^(v9))
    }
    var v166 int = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v126:", v126)
    fmt.Println("v159:", v159)
    fmt.Println("v166:", v166)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v126
    _ = v15
    _ = v159
    _ = v166
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
