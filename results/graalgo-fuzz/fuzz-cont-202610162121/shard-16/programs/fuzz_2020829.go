package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int64; F2 int16 }
type GenS5 struct { F0 uint16; F1 bool; F2 int16 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint8; F1 int8 }
type GenS12 struct { F0 int32; F1 int16; F2 uint8 }
type GenS13 struct { F0 int32; F1 uint; F2 int16 }

func main() {
    var v1 uint8 = 207
    var v2 bool = (564 == 9223372036854775807)
    var v3 int32 = 735
    v4 := []int{}
    for fi5 := 0; fi5 < 10; fi5++ {
        switch (fi5 >> ((65535) & 63)) {
        case 42:
            switch (("a" + "a") + (("hello" + (("a" + ("\n\t" + "a")) + ("foo bar" + ""))) + ((("hello" + "a") + "foo bar") + ("" + (("hello" + "") + ("hello" + "a")))))) {
            case "a":
                fmt.Println(v2, v3, v3)
            case "foo bar":
                for fi6 := 0; fi6 < 20; fi6++ {
                    fmt.Println(v3, fi6)
                    v7 := GenS1{F0: (v3 / ((v3) | 1)), F1: 2}
                    var v8 uint32 = uint32(v1)
                    if ((!false) && (369 <= 0)) {
                        var v9 int = (fi5 + fi5)
                        _ = v9
                    } else {
                        switch (("" + ((("" + "hello") + "a") + "hello")) + "foo bar") {
                        case "foo bar":
                            var v10 int16 = 42
                            var v11 bool = true
                            _ = v10
                            _ = v11
                        case "hello":
                            v12 := GenS1{F0: v3, F1: 4294967295}
                            _ = v12
                        case "\n\t":
                            v2 = ((0.1 / (1e10 + 3.14159)) == (float32(0.1) / (0.0 - 1e10)))
                        default:
                            v13 := GenS1{F0: (v3 % ((v3) | 1)), F1: 0}
                            _ = v13
                        }
                        var v14 string = ("" + "foo bar")
                        switch "hello" {
                        case "":
                            var v15 float64 = ((1e10 - 3.14159) - 0.1)
                            _ = v15
                        case "a":
                            v4 = append(v4, 1)
                            var v16 int8 = -127
                            v4 = append(v4, (fi5 / (((fi5 * 4294967295)) | 1)))
                            var v17 float64 = (1.5 + (0.1 - (0.0 / 1.0)))
                            _ = v16
                            _ = v17
                        default:
                            var v18 string = (v14 + (v14 + v14))
                            _ = v18
                        }
                        if (-1 < 68) {
                            if (v3 < (v3 | (-(v3)))) {
                                switch v14 {
                                case "a":
                                    var v19 uint = 42
                                    _ = v19
                                case "foo bar":
                                    if ((v2 || v2) || v2) {
                                        if false {
                                            if v2 {
                                                switch 0 {
                                                case 4294967295:
                                                    switch v14 {
                                                    case "hello":
                                                        v20 := GenS1{F0: v3, F1: 181}
                                                        _ = v20
                                                    case "":
                                                        if v2 {
                                                            var v21 uint16 = 65535
                                                            var v22 int8 = 42
                                                            var v23 float64 = 0.0
                                                            switch v14 {
                                                            case "foo bar":
                                                                v24 := GenS1{F0: (v3 * (v3 >> ((349) & 63))), F1: uint64(v21)}
                                                                _ = v24
                                                            case "hello":
                                                                var v25 string = "a"
                                                                _ = v25
                                                            case "a":
                                                                v26 := GenS2{F0: (-(v3))}
                                                                _ = v26
                                                            }
                                                            _ = v21
                                                            _ = v22
                                                            _ = v23
                                                        }
                                                    }
                                                case -209:
                                                    if (2 > 18446744073709551615) {
                                                        v2 = (v2 || v2)
                                                        if (0 < 1) {
                                                            var v27 uint32 = (v8 | v8)
                                                            v28 := GenS2{F0: 1}
                                                            v27 -= (v27 << ((255) & 63))
                                                            if (4294967295 < 2) {
                                                                if (-1 > 1) {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = fi6
                                                                    }
                                                                    v29 := GenS1{F0: v3, F1: 371}
                                                                    _ = v29
                                                                }
                                                            } else {
                                                                _ = len(v4)
                                                                switch (fi5 ^ 985) {
                                                                case 1:
                                                                    if (((v14 + ("\n\t" + "foo bar")) + ("hello" + (("hello" + "\n\t") + (v14 + ("hello" + v14))))) >= v14) {
                                                                        var v30 string = "hello"
                                                                        v31 := GenS1{F0: -1, F1: 1}
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                case 831:
                                                                    if false {
                                                                        v32 := GenS2{F0: (-(v3))}
                                                                        _ = v32
                                                                    } else {
                                                                        fmt.Println(v27, v27, v8)
                                                                    }
                                                                case 2147483647:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                default:
                                                                    v33 := GenS3{F0: 0}
                                                                    _ = v33
                                                                }
                                                            }
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                        var v34 int8 = 74
                                                        if (42 < 1) {
                                                            v8 = uint32(v1)
                                                            if len(v4) > 0 {
                                                                v4[0] = (fi6 % (((-(fi6))) | 1))
                                                            }
                                                        }
                                                        _ = v34
                                                    } else {
                                                        var v35 uint = 2
                                                        if v2 {
                                                            fmt.Println(fi5, v14)
                                                            switch v14 {
                                                            case "a":
                                                                _ = len(v4)
                                                            case "hello":
                                                                if false {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v3, v3, v14)
                                                            }
                                                            v36 := GenS2{F0: -13}
                                                            _ = v36
                                                        }
                                                        if (v2 && ((v14 != v14) && v2)) {
                                                            switch v14 {
                                                            case "foo bar":
                                                                var v37 int32 = v3
                                                                _ = v37
                                                            default:
                                                                v14 = (v14 + v14)
                                                                var v38 int32 = v3
                                                                _ = len(v4)
                                                                _ = v38
                                                            }
                                                            v39 := GenS3{F0: -1}
                                                            _ = v39
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = (^(fi6))
                                                            }
                                                        }
                                                        _ = v35
                                                    }
                                                default:
                                                    if false {
                                                        switch v14 {
                                                        case "a":
                                                            fmt.Println(fi5, v14, v3)
                                                        case "\n\t":
                                                            v14 = v14
                                                        case "":
                                                            v40 := GenS1{F0: (v3 ^ v3), F1: 398}
                                                            _ = v40
                                                        default:
                                                            var v41 uint32 = (v8 << ((42) & 63))
                                                            _ = v41
                                                        }
                                                    }
                                                }
                                                var v42 uint8 = 42
                                                _ = v42
                                            } else {
                                                var v43 string = (v14 + v14)
                                                _ = len(v4)
                                                if (!(!false)) {
                                                    var v44 string = ("foo bar" + (v43 + "a"))
                                                    v14 = ("hello" + v14)
                                                    _ = v44
                                                }
                                                _ = v43
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = 993
                                            }
                                        }
                                    }
                                    fmt.Println(v2)
                                    v45 := GenS4{F0: (fi5 / (((fi5 | (fi6 * (-(fi6))))) | 1)), F1: -9223372036854775808, F2: -985}
                                    _ = v45
                                default:
                                    var v46 uint16 = 56
                                    _ = v46
                                }
                                fmt.Println(v8, v3)
                            } else {
                                fmt.Println(v3, fi5, v3)
                                switch v14 {
                                case "a":
                                    v1 += v1
                                default:
                                    var v47 int32 = v3
                                    _ = v47
                                }
                            }
                            switch ("" + "\n\t") {
                            case "hello":
                                if (((float64((0.0 * (((0.1 - 0.1) + 1.5) / float32(3.14159)))) * 1e10) * (0.1 - 0.0)) != (0.0 / ((-2.25 + 1.5) / float64(1e10)))) {
                                    var v48 string = "a"
                                    _ = len(v4)
                                    var v49 uint8 = 42
                                    _ = v48
                                    _ = v49
                                } else {
                                    v50 := GenS4{F0: fi6, F1: 508, F2: -809}
                                    _ = v50
                                }
                            case "foo bar":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                if (!(0 >= -700)) {
                                    var v51 int64 = 42
                                    v52 := GenS5{F0: 2, F1: false, F2: 0}
                                    switch "foo bar" {
                                    case "":
                                        v1 = 131
                                    case "foo bar":
                                        fmt.Println(v3, v1, fi5)
                                    case "a":
                                        var v53 int32 = -2147483648
                                        _ = v53
                                    default:
                                        if v2 {
                                            v51 += 1
                                        } else {
                                            switch -276 {
                                            case -555:
                                                fmt.Println(v2)
                                            case -825:
                                                v54 := GenS3{F0: -1}
                                                _ = v54
                                            case 4294967295:
                                                fmt.Println(v8, fi6, v14)
                                                v8 -= v8
                                                var v55 bool = (-62 == 0)
                                                _ = v55
                                            default:
                                                if false {
                                                    v14 = ((v14 + "a") + v14)
                                                }
                                            }
                                            v3 += v3
                                            var v56 int32 = int32(v8)
                                            _ = len(v4)
                                            _ = v56
                                        }
                                    }
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v57 int = 32767
                                    fmt.Println(fi6)
                                    var v58 uint8 = 179
                                    _ = v57
                                    _ = v58
                                }
                            }
                        } else {
                            var v59 uint = 4294967295
                            _ = v59
                        }
                        _ = v14
                    }
                    switch "a" {
                    case "foo bar":
                        if v2 {
                            if v2 {
                                v3 -= (v3 << ((659) & 63))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            switch -2147483648 {
                            case 32767:
                                if false {
                                    var v60 int16 = -1
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v60
                                }
                            case -262:
                                switch ((("a" + "foo bar") + "foo bar") + "\n\t") {
                                case "hello":
                                    v61 := GenS4{F0: 197, F1: 183, F2: 32767}
                                    _ = v61
                                }
                            case 2147483647:
                                v8 = 2147483647
                            default:
                                v2 = (-1 == -38)
                            }
                        }
                    }
                    if (-619 > 42) {
                        v3 = 145
                        fmt.Println(fi5)
                    } else {
                        var v62 string = "\n\t"
                        if v2 {
                            if (v62 <= ("\n\t" + v62)) {
                                v4 = append(v4, (^(fi6)))
                            } else {
                                fmt.Println(v8, v2, v62)
                            }
                            if true {
                                fmt.Println(v1, fi6, fi5)
                                fmt.Println(fi6, fi5, fi5)
                            } else {
                                var v63 bool = false
                                v2 = v63
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v63
                            }
                            switch fi5 {
                            case 1:
                                v1 += (v1 ^ (v1 & v1))
                                v8 -= (v8 - (v8 * (v8 + v8)))
                                v4 = append(v4, 1)
                            }
                        } else {
                            switch "hello" {
                            case "a":
                                if (3.14159 == 0.0) {
                                    var v64 int32 = v3
                                    _ = v64
                                } else {
                                    switch (^(fi5)) {
                                    case 32767:
                                        fmt.Println(fi5)
                                        var v65 int64 = 1
                                        _ = v65
                                    default:
                                        var v66 int8 = 1
                                        _ = v66
                                    }
                                }
                            case "":
                                v67 := GenS6{F0: "a", F1: 100}
                                _ = v67
                            }
                        }
                        var v68 int16 = 1000
                        _ = v62
                        _ = v68
                    }
                    v2 = true
                    if false {
                        fmt.Println(v1, v8)
                        fmt.Println(v1, v3)
                        var v69 float64 = ((1.5 / -2.25) - -2.25)
                        _ = v69
                    } else {
                        if (!v2) {
                            v4 = append(v4, (fi6 % (((fi5 >> ((255) & 63))) | 1)))
                            var v70 int64 = 233
                            v71 := GenS7{F0: (v1 < v1), F1: (v1 / (((v1 ^ 255)) | 1)), F2: 1}
                            _ = v70
                            _ = v71
                        } else {
                            if false {
                                _ = len(v4)
                            }
                            if (v3 <= (^(v3))) {
                                var v72 bool = (false || v2)
                                _ = v72
                            }
                            fmt.Println(v8, v8, v1)
                        }
                        v4 = append(v4, (fi5 + fi6))
                        v2 = v2
                    }
                    fmt.Println(v2, fi5, v3)
                    v2 = ("hello" <= "a")
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            case "hello":
                for fi73 := 0; fi73 < 1; fi73++ {
                    var v74 int = (fi5 * fi5)
                    v4 = append(v4, (fi5 ^ (fi5 + (fi5 * fi5))))
                    _ = fi73
                    _ = v74
                }
            }
        case 127:
            if (65535 > 42) {
                var v75 int = fi5
                for fi76 := 0; fi76 < 15; fi76++ {
                    v3 += (v3 % ((v3) | 1))
                    var v77 int32 = v3
                    _ = fi76
                    _ = v77
                }
                if (false || false) {
                    var v78 float32 = 3.14159
                    _ = v78
                } else {
                    var v79 int = fi5
                    _ = v79
                }
                _ = v75
            } else {
                if (v2 && (1 != 14)) {
                    v3 += -2147483648
                } else {
                    switch fi5 {
                    case 0:
                        v80 := GenS6{F0: "a", F1: 100}
                        _ = v80
                    }
                    switch "" {
                    case "\n\t":
                        if true {
                            v1 = v1
                        } else {
                            var v81 int32 = v3
                            v1 = 156
                            _ = v81
                        }
                    case "foo bar":
                        v82 := GenS7{F0: v2, F1: v1, F2: v1}
                        _ = v82
                    case "a":
                        v83 := GenS5{F0: 1, F1: (65535 >= 193), F2: 461}
                        _ = v83
                    }
                }
                for fi84 := 0; fi84 < 6; fi84++ {
                    var v85 int64 = -1
                    v86 := GenS2{F0: v3}
                    fmt.Println(v2)
                    var v87 int = 32768
                    _ = fi84
                    _ = v85
                    _ = v86
                    _ = v87
                }
                if (0 <= 938) {
                    var v88 int16 = 1
                    _ = v88
                } else {
                    v2 = v2
                    for fi89 := 0; fi89 < 16; fi89++ {
                        var v90 int64 = 1
                        _ = fi89
                        _ = v90
                    }
                    if v2 {
                        var v91 float32 = float32(((3.14159 / 3.14159) / 1.0))
                        var v92 uint16 = 32767
                        _ = v91
                        _ = v92
                    } else {
                        fmt.Println(v3, fi5)
                        var v93 int16 = -629
                        fmt.Println(v3, fi5)
                        _ = v93
                    }
                }
            }
        default:
            var v94 float64 = 0.1
            _ = v94
        }
        var v95 bool = (v3 >= v3)
        _ = fi5
        _ = v95
    }
    v96 := GenS4{F0: 128, F1: 183, F2: 61}
    v1 -= v1
    v97 := GenS8{F0: 1000, F1: 815}
    v98 := GenS9{F0: 1}
    var v99 uint32 = 2147483647
    var v100 bool = (!(!v2))
    for fi101 := 0; fi101 < 3; fi101++ {
        for fi102 := 0; fi102 < 2; fi102++ {
            v103 := GenS8{F0: 23, F1: 18446744073709551615}
            if v2 {
                var v104 uint8 = v1
                var v105 uint8 = v1
                if (!v2) {
                    var v106 string = (("a" + "foo bar") + "")
                    var v107 int8 = 40
                    if ((v100 || v2) || false) {
                        v108 := GenS10{F0: ""}
                        v104 = v1
                        _ = v108
                    } else {
                        var v109 int32 = v3
                        var v110 bool = v2
                        fmt.Println(v1, v106)
                        switch fi102 {
                        case 127:
                            v111 := GenS11{F0: 42, F1: v107}
                            _ = v111
                        case 42:
                            switch (v106 + (v106 + v106)) {
                            case "hello":
                                v104 -= v104
                            default:
                                var v112 int32 = (v3 >> ((700) & 63))
                                _ = v112
                            }
                        case -639:
                            if v2 {
                                v3 += (v3 % ((42) | 1))
                                if len(v4) > 0 {
                                    v4[0] = 127
                                }
                            } else {
                                var v113 uint16 = 374
                                if len(v4) > 0 {
                                    v4[0] = -567
                                }
                                _ = v113
                            }
                        }
                        _ = v109
                        _ = v110
                    }
                    _ = v106
                    _ = v107
                } else {
                    switch "foo bar" {
                    case "":
                        fmt.Println(fi101, v2)
                    case "a":
                        var v114 int64 = 53
                        _ = v114
                    case "\n\t":
                        if (!(v2 || v100)) {
                            fmt.Println(v1, v2, v99)
                        } else {
                            fmt.Println(fi102, v1, v1)
                            var v115 bool = ((fi102 & (-(fi101))) <= fi101)
                            v3 = 441
                            _ = v115
                        }
                    }
                }
                _ = v104
                _ = v105
            } else {
                if true {
                    var v116 uint = 1
                    v117 := GenS10{F0: "a"}
                    switch (^(fi101)) {
                    case 32768:
                        v118 := GenS12{F0: (v3 / ((v3) | 1)), F1: -474, F2: 42}
                        _ = v118
                    }
                    v119 := GenS10{F0: "hello"}
                    _ = v116
                    _ = v117
                    _ = v119
                } else {
                    fmt.Println(v2)
                }
                var v120 float64 = 1e10
                _ = v120
            }
            var v121 uint8 = (^(v1))
            _ = fi102
            _ = v103
            _ = v121
        }
        switch fi101 {
        case 0:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case -1:
            fmt.Println(v3, v2, v100)
        case 351:
            var v122 float64 = (((-2.25 * 3.14159) - 3.14159) / -2.25)
            _ = v122
        default:
            for fi123 := 0; fi123 < 13; fi123++ {
                if (false && false) {
                    if (!(!v100)) {
                        _ = len(v4)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        v4 = append(v4, fi101)
                        if true {
                            var v124 uint32 = uint32(v1)
                            _ = v124
                        }
                    }
                    v2 = v2
                } else {
                    var v125 bool = (true || (0 <= 42))
                    v99 = v99
                    _ = v125
                }
                switch 42 {
                case 128:
                    v1 = (v1 + 42)
                case -376:
                    v126 := GenS12{F0: (v3 / ((v3) | 1)), F1: -1, F2: (v1 ^ (v1 << ((684) & 63)))}
                    v127 := GenS4{F0: (fi101 / (((fi101 / ((fi101) | 1))) | 1)), F1: 1, F2: -32768}
                    _ = v126
                    _ = v127
                }
                v128 := GenS13{F0: (v3 * (v3 * (-(v3)))), F1: 255, F2: 42}
                _ = fi123
                _ = v128
            }
        }
        _ = fi101
    }
    v99 = v99
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
