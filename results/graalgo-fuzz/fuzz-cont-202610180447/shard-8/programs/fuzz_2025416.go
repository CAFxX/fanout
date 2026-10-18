package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint; F1 int64; F2 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int16; F1 uint }
type GenS13 struct { F0 bool; F1 int32; F2 uint16 }
type GenS14 struct { F0 int8; F1 uint16; F2 int64 }
type GenS15 struct { F0 int; F1 bool; F2 string }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint64; F1 bool }
type GenS19 struct { F0 uint; F1 uint8 }
type GenS20 struct { F0 uint16; F1 string; F2 int }
type GenS21 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS22 struct { F0 uint; F1 uint8 }
type GenS23 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint16, p1 uint) int64 {
    p0 += (p0 ^ (p0 | 256))
    p1 += p1
    fmt.Println(p0, p0)
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 933
}
func genFn1(p0 uint, p1 string) bool {
    v6 := []int{32768, 0}
    p0 = p0
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            if true {
                var v9 int32 = -602
                v10 := GenS1{F0: (p0 << ((11) & 63)), F1: (^(p0)), F2: fi7}
                v9 -= (v9 % (((v9 - (^(v9)))) | 1))
                p0 -= p0
                if (247 == 852) {
                    if true {
                        fmt.Println(fi8, fi8)
                    } else {
                        switch fi7 {
                        case 9223372036854775807:
                            fmt.Println(p1)
                        case 32767:
                            v6 = append(v6, 732)
                        }
                        p0 -= 4294967295
                        p0 = (^(p0))
                    }
                    fmt.Println(v9, p0, v9)
                } else {
                    if (!(0.0 <= (1.5 / 1.5))) {
                        if false {
                            if (!(fi8 >= (fi8 | (fi7 & (fi8 % ((fi8) | 1)))))) {
                                v11 := GenS1{F0: p0, F1: p0, F2: (fi7 >> (((^(p0))) & 63))}
                                v12 := GenS2{F0: 32767}
                                _ = v11
                                _ = v12
                            } else {
                                v13 := GenS3{F0: (^(p0))}
                                if ((false || false) || true) {
                                    v14 := GenS4{F0: 255, F1: p1}
                                    var v15 int64 = -447
                                    _ = v14
                                    _ = v15
                                } else {
                                    v16 := GenS4{F0: 1, F1: p1}
                                    if true {
                                        if (898 != 985) {
                                            fmt.Println(p0, fi7)
                                        } else {
                                            if (v9 < v9) {
                                                if false {
                                                    var v17 uint16 = uint16(p0)
                                                    v9 = (v9 * v9)
                                                    switch p1 {
                                                    case "\n\t":
                                                        v9 = v9
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v18 int16 = 1000
                                                            _ = v18
                                                        case "":
                                                            switch fi8 {
                                                            case -2147483648:
                                                                v19 := GenS1{F0: p0, F1: p0, F2: (fi7 + int(v17))}
                                                                _ = v19
                                                            case 128:
                                                                var v20 int64 = 1
                                                                _ = v20
                                                            case -612:
                                                                var v21 bool = true
                                                                _ = v21
                                                            }
                                                        case "\n\t":
                                                            if ((615 != -1) || true) {
                                                                v17 = (v17 << (((p0 % ((p0) | 1))) & 63))
                                                                switch 1 {
                                                                case -687:
                                                                    v22 := GenS5{F0: 1}
                                                                    _ = v22
                                                                case 1:
                                                                    var v23 string = (p1 + "\n\t")
                                                                    _ = v23
                                                                case 42:
                                                                    v24 := GenS2{F0: 42}
                                                                    _ = v24
                                                                default:
                                                                    v6 = append(v6, fi8)
                                                                }
                                                                v6 = append(v6, int(v9))
                                                                fmt.Println(p1, v9, v17)
                                                            } else {
                                                                if false {
                                                                    v25 := GenS6{F0: -186, F1: (v17 << (((^(p0))) & 63)), F2: (v9 + 0)}
                                                                    switch ((((("" + (p1 + (("" + ((p1 + "foo bar") + (p1 + (p1 + "a")))) + "foo bar"))) + ("\n\t" + "hello")) + (p1 + (("\n\t" + ("hello" + p1)) + ((("hello" + ("" + "\n\t")) + "\n\t") + p1)))) + p1) + ((p1 + p1) + p1)) {
                                                                    case "hello":
                                                                        v6 = append(v6, fi8)
                                                                    case "\n\t":
                                                                        p0 += p0
                                                                    case "foo bar":
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            var v26 int16 = 1000
                                                                            _ = v26
                                                                        case "a":
                                                                            var v27 int64 = 0
                                                                            _ = v27
                                                                        case "\n\t":
                                                                            var v28 int8 = 0
                                                                            if (!(!false)) {
                                                                                fmt.Println(v9)
                                                                                if (!false) {
                                                                                    var v29 uint8 = 2
                                                                                    p0 = (^(p0))
                                                                                    _ = v29
                                                                                }
                                                                            }
                                                                            _ = v28
                                                                        }
                                                                    default:
                                                                        var v30 uint32 = 777
                                                                        _ = v30
                                                                    }
                                                                    if len(v6) > 0 {
                                                                        v6[0] = (fi7 ^ fi7)
                                                                    }
                                                                    if false {
                                                                        switch (p1 + p1) {
                                                                        case "hello":
                                                                            v31 := GenS7{F0: (p0 + (p0 & p0)), F1: -842, F2: 127}
                                                                            _ = v31
                                                                        default:
                                                                            p0 = p0
                                                                        }
                                                                        switch (p1 + ("\n\t" + "foo bar")) {
                                                                        case "":
                                                                            if true {
                                                                                var v32 uint32 = 42
                                                                                switch -2147483648 {
                                                                                case 1:
                                                                                    if (true || true) {
                                                                                        if ((!true) || (2 == 4294967295)) {
                                                                                            v33 := GenS3{F0: 20}
                                                                                            v34 := GenS8{F0: v9}
                                                                                            p1 = "hello"
                                                                                            _ = v33
                                                                                            _ = v34
                                                                                        } else {
                                                                                            v35 := GenS5{F0: 15}
                                                                                            p1 = "foo bar"
                                                                                            if ((fi7 >= fi8) && (true || (true || ((p1 + "hello") == ((p1 + ("a" + ((p1 + ("\n\t" + "")) + "\n\t"))) + (p1 + ((p1 + ((p1 + ("\n\t" + p1)) + (p1 + (p1 + "foo bar")))) + ((p1 + ((p1 + "foo bar") + p1)) + (p1 + p1))))))))) {
                                                                                                if (0 == 161) {
                                                                                                    v9 = (v9 & 2147483647)
                                                                                                    p0 = 18446744073709551615
                                                                                                    var v36 bool = (uint8(p0) != 114)
                                                                                                    _ = v36
                                                                                                } else {
                                                                                                    p1 = (("foo bar" + "a") + ((((p1 + p1) + p1) + p1) + ""))
                                                                                                }
                                                                                            } else {
                                                                                                v32 = 4294967295
                                                                                                v6 = append(v6, int(p0))
                                                                                            }
                                                                                            _ = v35
                                                                                        }
                                                                                        var v37 int16 = 32767
                                                                                        var v38 string = "hello"
                                                                                        _ = v37
                                                                                        _ = v38
                                                                                    } else {
                                                                                        if false {
                                                                                            var v39 uint8 = 236
                                                                                            fmt.Println(v17)
                                                                                            _ = v39
                                                                                        }
                                                                                        v40 := GenS4{F0: v17, F1: (p1 + ("foo bar" + ("" + p1)))}
                                                                                        _ = v40
                                                                                    }
                                                                                case 128:
                                                                                    p0 -= uint(v17)
                                                                                case 9223372036854775807:
                                                                                    var v41 float32 = 1e10
                                                                                    _ = v41
                                                                                default:
                                                                                    p1 = ("" + p1)
                                                                                }
                                                                                _ = v32
                                                                            } else {
                                                                                v42 := GenS5{F0: -128}
                                                                                v9 = v9
                                                                                v43 := GenS4{F0: (v17 << (((p0 | p0)) & 63)), F1: (p1 + p1)}
                                                                                v17 += uint16(p0)
                                                                                fmt.Println(fi7, p0, v17)
                                                                                if (false || true) {
                                                                                    v44 := GenS6{F0: 4294967295, F1: 757, F2: 42}
                                                                                    v17 = v17
                                                                                    v45 := GenS9{F0: (false && (!(0 > 207)))}
                                                                                    _ = v44
                                                                                    _ = v45
                                                                                }
                                                                                p1 = p1
                                                                                _ = v42
                                                                                _ = v43
                                                                            }
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            var v46 uint32 = 42
                                                                            _ = v46
                                                                        }
                                                                        p0 -= (p0 / (((p0 / (((p0 << (((p0 / ((1) | 1))) & 63))) | 1))) | 1))
                                                                        _ = len(v6)
                                                                    }
                                                                    var v47 float64 = 1e10
                                                                    fmt.Println(v47, v47)
                                                                    _ = v25
                                                                    _ = v47
                                                                }
                                                                var v48 int64 = -724
                                                                _ = v48
                                                            }
                                                        }
                                                    default:
                                                        p0 -= 947
                                                    }
                                                    v9 += v9
                                                    _ = v17
                                                }
                                            } else {
                                                _ = len(v6)
                                                if true {
                                                    v9 += v9
                                                }
                                                var v49 int16 = -380
                                                _ = v49
                                            }
                                            v9 = (-(v9))
                                        }
                                        v50 := GenS10{F0: -1}
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        p1 = p1
                                        _ = v50
                                    } else {
                                        v51 := GenS2{F0: -99}
                                        switch "a" {
                                        case "":
                                            v52 := GenS11{F0: (v9 * (-(v9)))}
                                            var v53 int16 = -32768
                                            v54 := GenS9{F0: false}
                                            var v55 int = (^(fi8))
                                            if false {
                                                if true {
                                                    v53 += (v53 % ((v53) | 1))
                                                    if (uint64(p0) == 1) {
                                                        v9 = (-(v9))
                                                    }
                                                } else {
                                                    p1 = p1
                                                    if (-1 <= genFn0(1, p0)) {
                                                        p1 = (p1 + (p1 + "a"))
                                                        switch (^(v55)) {
                                                        case 9223372036854775807:
                                                            v56 := GenS9{F0: true}
                                                            _ = v56
                                                        default:
                                                            if true {
                                                                var v57 uint64 = 4294967295
                                                                var v58 uint32 = 2147483647
                                                                if false {
                                                                    v59 := GenS12{F0: 995, F1: (^(p0))}
                                                                    switch p1 {
                                                                    case "\n\t":
                                                                        v53 = int16(v9)
                                                                    case "hello":
                                                                        v60 := GenS9{F0: (!("hello" < ((p1 + ((p1 + p1) + p1)) + p1)))}
                                                                        _ = v60
                                                                    default:
                                                                        var v61 int = v55
                                                                        _ = v61
                                                                    }
                                                                    _ = v59
                                                                } else {
                                                                    v62 := GenS4{F0: 814, F1: "foo bar"}
                                                                    _ = v62
                                                                }
                                                                _ = v57
                                                                _ = v58
                                                            }
                                                            v63 := GenS11{F0: v9}
                                                            _ = v63
                                                        }
                                                    } else {
                                                        v64 := GenS13{F0: false, F1: (v9 ^ 42), F2: 322}
                                                        _ = v64
                                                    }
                                                }
                                            }
                                            fmt.Println(v53)
                                            v9 = 42
                                            if (1.5 >= ((0.1 - -2.25) - 3.14159)) {
                                                var v65 uint = p0
                                                _ = v65
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        case "foo bar":
                                            switch (fi8 >> ((2) & 63)) {
                                            case 1:
                                                p0 = (p0 + p0)
                                            }
                                            v9 -= (-(v9))
                                        case "\n\t":
                                            p0 -= (p0 << ((940) & 63))
                                        }
                                        _ = v51
                                    }
                                    var v66 uint8 = 0
                                    switch (fi7 >> ((18446744073709551615) & 63)) {
                                    case 778:
                                        var v67 uint8 = v66
                                        var v68 uint16 = 42
                                        _ = v67
                                        _ = v68
                                    case 32767:
                                        fmt.Println(p1, p1, fi8)
                                    case -1:
                                        fmt.Println(v66, v9)
                                    }
                                    p0 = (p0 ^ p0)
                                    switch fi8 {
                                    case 0:
                                        v69 := GenS2{F0: -806}
                                        _ = v69
                                    }
                                    _ = v16
                                    _ = v66
                                }
                                _ = v13
                            }
                        }
                        _ = len(v6)
                        if false {
                            var v70 uint64 = 805
                            var v71 uint16 = 972
                            p1 = p1
                            _ = v70
                            _ = v71
                        }
                    } else {
                        p0 = (p0 + (p0 >> ((65535) & 63)))
                        if true {
                            switch 242 {
                            case 42:
                                p0 -= p0
                            case 127:
                                v6 = append(v6, fi8)
                            case 128:
                                var v72 int64 = -9223372036854775808
                                _ = v72
                            }
                            switch p1 {
                            case "":
                                var v73 uint = 18446744073709551615
                                _ = v73
                            case "hello":
                                switch (p1 + p1) {
                                case "foo bar":
                                    if (127 == 15) {
                                        _ = len(v6)
                                        switch fi7 {
                                        case 9223372036854775807:
                                            switch (fi8 | (fi7 >> ((p0) & 63))) {
                                            case 1:
                                                if true {
                                                    if true {
                                                        if (1.5 != 0.0) {
                                                            p1 = ("hello" + p1)
                                                        } else {
                                                            p0 = 636
                                                        }
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        v74 := GenS14{F0: -116, F1: 32767, F2: -1}
                                                        if len(v6) > 0 {
                                                            v6[0] = fi7
                                                        }
                                                        fmt.Println(p0)
                                                        var v75 int8 = int8(fi8)
                                                        v76 := GenS15{F0: fi7, F1: false, F2: p1}
                                                        _ = v74
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                    var v77 uint16 = 0
                                                    _ = v77
                                                }
                                            case 470:
                                                p0 = p0
                                            case 9223372036854775807:
                                                var v78 int32 = (-(v9))
                                                _ = v78
                                            }
                                        }
                                    } else {
                                        if (0 >= 1) {
                                            var v79 int64 = 9223372036854775807
                                            if ((1e10 / ((3.14159 / float32(float64(3.14159))) + -2.25)) >= ((1e10 * 3.14159) * -2.25)) {
                                                v80 := GenS1{F0: 65535, F1: 911, F2: fi7}
                                                fmt.Println(fi7)
                                                switch p1 {
                                                case "hello":
                                                    var v81 int = (fi7 >> ((862) & 63))
                                                    _ = v81
                                                default:
                                                    var v82 uint = p0
                                                    _ = v82
                                                }
                                                _ = v80
                                            } else {
                                                if len(v6) > 0 {
                                                    v6[0] = 128
                                                }
                                            }
                                            v9 += (v9 & (v9 | (v9 ^ v9)))
                                            _ = v79
                                        }
                                        var v83 int = 9223372036854775807
                                        _ = v83
                                    }
                                case "a":
                                    p1 = (((((p1 + p1) + p1) + "") + (p1 + (p1 + "a"))) + p1)
                                case "hello":
                                    fmt.Println(fi7, fi8)
                                }
                                var v84 int = (-(fi8))
                                p1 = (("hello" + (p1 + "")) + p1)
                                _ = v84
                            case "foo bar":
                                switch p1 {
                                case "":
                                    var v85 int16 = 0
                                    _ = v85
                                case "\n\t":
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                fmt.Println(p1)
                                var v86 bool = false
                                _ = v86
                            }
                        }
                        switch (p1 + "") {
                        case "":
                            var v87 int32 = 204
                            _ = v87
                        }
                    }
                }
                _ = len(v6)
                _ = v10
                _ = v9
            } else {
                v88 := GenS7{F0: p0, F1: int64(p0), F2: 42}
                _ = v88
            }
            _ = fi8
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if false {
            switch fi7 {
            case 42:
                for fi89 := 0; fi89 < 8; fi89++ {
                    switch (p1 + p1) {
                    case "a":
                        var v90 float32 = float32((-2.25 - 3.14159))
                        _ = v90
                    case "\n\t":
                        var v91 int8 = 42
                        _ = v91
                    case "":
                        p0 = (p0 << (((p0 ^ (p0 ^ (p0 ^ (^(p0)))))) & 63))
                    }
                    _ = fi89
                }
            case -930:
                for fi92 := 0; fi92 < 17; fi92++ {
                    switch p1 {
                    case "hello":
                        switch fi7 {
                        case 9223372036854775807:
                            p1 = p1
                        case 0:
                            _ = len(v6)
                        default:
                            v93 := GenS16{F0: 128}
                            var v94 bool = true
                            if v94 {
                                p0 += p0
                            } else {
                                var v95 int16 = 42
                                v95 += 1000
                                _ = v95
                            }
                            if v94 {
                                v6 = append(v6, (fi7 << (((p0 >> ((p0) & 63))) & 63)))
                            } else {
                                var v96 uint64 = 1
                                if len(v6) > 0 {
                                    v6[0] = fi92
                                }
                                if v94 {
                                    var v97 int32 = -2147483648
                                    switch (fi92 << ((581) & 63)) {
                                    case 4294967295:
                                        switch p1 {
                                        case "hello":
                                            if ((float32(float64((0.1 + float32(1.5)))) - -2.25) < ((((3.14159 - (1e10 + 0.1)) + (((3.14159 + 1.5) / (0.0 - (0.1 / (1e10 / 1.5)))) / -2.25)) * 1e10) + (-2.25 / 3.14159))) {
                                                if (1 > -32768) {
                                                    var v98 string = "foo bar"
                                                    if false {
                                                        if false {
                                                            v99 := GenS17{F0: v97}
                                                            v97 = v97
                                                            _ = len(v6)
                                                            _ = v99
                                                        } else {
                                                            v96 += (v96 % ((18446744073709551615) | 1))
                                                            v96 = (v96 * v96)
                                                            var v100 int16 = 1000
                                                            switch 128 {
                                                            case 1:
                                                                v97 += (v97 * v97)
                                                            case 4294967295:
                                                                var v101 string = "hello"
                                                                _ = v101
                                                            default:
                                                                fmt.Println(v98, fi7)
                                                            }
                                                            _ = v100
                                                        }
                                                        var v102 int32 = (^(v97))
                                                        var v103 int16 = -18
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    fmt.Println(p0, fi92, v96)
                                                    _ = v98
                                                }
                                                switch -230 {
                                                case -710:
                                                    var v104 int64 = 1
                                                    _ = v104
                                                case -181:
                                                    v97 = (v97 & v97)
                                                case 1:
                                                    var v105 uint64 = (v96 - uint64(p0))
                                                    _ = v105
                                                }
                                            }
                                        case "":
                                            switch p1 {
                                            case "a":
                                                switch fi92 {
                                                case 4294967295:
                                                    var v106 uint = 65535
                                                    _ = v106
                                                case 381:
                                                    var v107 int32 = (v97 * -231)
                                                    _ = v107
                                                default:
                                                    if (42 <= 218) {
                                                        if v94 {
                                                            var v108 int8 = -1
                                                            _ = v108
                                                        }
                                                        var v109 bool = (1000 == -383)
                                                        _ = v109
                                                    } else {
                                                        var v110 int = fi92
                                                        var v111 uint16 = 32767
                                                        v94 = v94
                                                        if v94 {
                                                            v112 := GenS18{F0: (v96 + 1), F1: (true && (102 != 0))}
                                                            _ = v112
                                                        }
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                    var v113 int8 = -3
                                                    _ = v113
                                                }
                                            case "hello":
                                                if v94 {
                                                    v114 := GenS1{F0: p0, F1: (p0 ^ p0), F2: (fi7 ^ fi92)}
                                                    _ = v114
                                                }
                                            case "":
                                                var v115 float32 = (0.0 + ((0.1 - 0.0) / (-2.25 / ((((((float32(1e10) * 1.5) - (((1e10 * ((((1e10 * 1.5) * (0.0 - float32((float64((0.1 - -2.25)) / 1.0)))) / 0.1) + ((float32((((-2.25 - 0.0) + 1e10) - (1.5 / 3.14159))) + 1.5) + -2.25))) - 1e10) * 1e10)) / ((float32((1e10 + 0.0)) + (0.0 / 3.14159)) - -2.25)) - -2.25) + (-2.25 * 0.1)) + 3.14159))))
                                                _ = v115
                                            }
                                        case "\n\t":
                                            if v94 {
                                                fmt.Println(fi7)
                                            } else {
                                                if (255 == 2) {
                                                    if (65535 == 88) {
                                                        if false {
                                                            switch 9223372036854775807 {
                                                            case 2147483647:
                                                                var v116 uint = p0
                                                                _ = v116
                                                            }
                                                            v97 += v97
                                                            v94 = v94
                                                            switch ("hello" + (p1 + "foo bar")) {
                                                            case "\n\t":
                                                                v117 := GenS19{F0: 18446744073709551615, F1: 127}
                                                                fmt.Println(fi92, p0, fi7)
                                                                _ = v117
                                                            }
                                                            if len(v6) > 0 {
                                                                v6[0] = 2147483647
                                                            }
                                                            if (!(!false)) {
                                                                var v118 float32 = -2.25
                                                                _ = v118
                                                            }
                                                        }
                                                    }
                                                    v97 = (-(v97))
                                                }
                                                if (!v94) {
                                                    if (1 > genFn0(42, 370)) {
                                                        v119 := GenS14{F0: 115, F1: 1, F2: genFn0(15, p0)}
                                                        _ = v119
                                                    }
                                                } else {
                                                    p1 = "hello"
                                                }
                                                var v120 int = int(p0)
                                                _ = v120
                                            }
                                        default:
                                            v97 -= (v97 + v97)
                                        }
                                    case -629:
                                        if (127 < -9223372036854775808) {
                                            if v94 {
                                                var v121 int = (fi7 & fi7)
                                                var v122 uint32 = 65535
                                                _ = v121
                                                _ = v122
                                            } else {
                                                v97 += (v97 >> (((p0 + 0)) & 63))
                                            }
                                            _ = len(v6)
                                            fmt.Println(fi92)
                                        }
                                    case 913:
                                        v6 = append(v6, -587)
                                    default:
                                        if v94 {
                                            if v94 {
                                                v96 += 18446744073709551615
                                                v123 := GenS14{F0: -128, F1: 255, F2: -337}
                                                v96 = v96
                                                fmt.Println(fi7, p1, v96)
                                                _ = v123
                                            } else {
                                                switch -69 {
                                                case 9223372036854775807:
                                                    switch p1 {
                                                    case "a":
                                                        v124 := GenS12{F0: -32768, F1: 2}
                                                        _ = v124
                                                    case "":
                                                        _ = len(v6)
                                                    }
                                                    if ((!((3.14159 * (1e10 / -2.25)) < -2.25)) || (!v94)) {
                                                        if len(v6) > 0 {
                                                            v6[0] = fi92
                                                        }
                                                    } else {
                                                        var v125 uint8 = 42
                                                        switch 127 {
                                                        case 251:
                                                            v97 = (v97 ^ (v97 << ((p0) & 63)))
                                                            var v126 int = (-(fi7))
                                                            _ = v126
                                                        }
                                                        if false {
                                                            if (v125 < (v125 + uint8(v96))) {
                                                                v6 = append(v6, -2147483648)
                                                            }
                                                            v125 = (^(v125))
                                                        } else {
                                                            v96 += 52
                                                        }
                                                        _ = v125
                                                    }
                                                case 190:
                                                    v127 := GenS13{F0: (true && (v94 && v94)), F1: 480, F2: 779}
                                                    _ = v127
                                                case 0:
                                                    var v128 float32 = float32((0.1 / -2.25))
                                                    _ = v128
                                                default:
                                                    if ((p0 >> ((p0) & 63)) <= p0) {
                                                        v97 = (v97 - 1)
                                                    } else {
                                                        fmt.Println(fi92, p1)
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                }
                                                if (!v94) {
                                                    var v129 uint8 = 2
                                                    _ = v129
                                                } else {
                                                    var v130 int8 = -71
                                                    var v131 uint8 = 128
                                                    _ = v130
                                                    _ = v131
                                                }
                                                if v94 {
                                                    var v132 int16 = -1
                                                    _ = v132
                                                }
                                                if (0 != -1) {
                                                    if (4294967295 < 65535) {
                                                        p1 = "a"
                                                        v133 := GenS20{F0: 0, F1: p1, F2: -2147483648}
                                                        _ = v133
                                                    } else {
                                                        if true {
                                                            var v134 float32 = float32(1e10)
                                                            _ = v134
                                                        }
                                                        v135 := GenS21{F0: 4294967295, F1: 1, F2: 51}
                                                        var v136 int16 = 0
                                                        v97 = v97
                                                        if (v94 && v94) {
                                                            v137 := GenS22{F0: (p0 * (p0 + (p0 % ((65535) | 1)))), F1: 135}
                                                            var v138 uint = 255
                                                            if (v96 == 1) {
                                                                var v139 int32 = (v97 * (v97 % (((v97 ^ v97)) | 1)))
                                                                _ = v139
                                                            } else {
                                                                v140 := GenS19{F0: uint(v96), F1: 2}
                                                                var v141 uint = 883
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                            _ = v137
                                                            _ = v138
                                                        }
                                                        _ = v135
                                                        _ = v136
                                                    }
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = fi92
                                                    }
                                                } else {
                                                    p1 = p1
                                                    v94 = ((fi7 | fi92) == (fi7 ^ (fi7 - (fi7 * fi92))))
                                                    p1 = p1
                                                    var v142 float32 = 0.1
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    var v143 uint8 = 1
                                                    fmt.Println(fi92)
                                                    switch (fi92 - -2147483648) {
                                                    case 942:
                                                        p1 = ((("a" + (p1 + p1)) + "") + "hello")
                                                    }
                                                    var v144 int = fi92
                                                    v145 := GenS12{F0: -32768, F1: p0}
                                                    _ = len(v6)
                                                    _ = v142
                                                    _ = v143
                                                    _ = v144
                                                    _ = v145
                                                }
                                            }
                                            var v146 int32 = (v97 ^ v97)
                                            _ = len(v6)
                                            _ = v146
                                        }
                                    }
                                    _ = v97
                                } else {
                                    if true {
                                        v96 -= v96
                                        var v147 uint = 2
                                        _ = v147
                                    }
                                }
                                _ = v96
                            }
                            var v148 uint32 = 0
                            v149 := GenS2{F0: -145}
                            _ = v148
                            _ = v149
                            _ = v93
                            _ = v94
                        }
                    case "\n\t":
                        p0 -= (p0 & 398)
                    }
                    p0 = p0
                    _ = len(v6)
                    _ = fi92
                }
            case 2147483647:
                v150 := GenS6{F0: fi7, F1: 255, F2: 299}
                _ = v150
            default:
                if false {
                    if false {
                        if true {
                            var v151 uint8 = 1
                            fmt.Println(p0)
                            _ = v151
                        } else {
                            var v152 string = ("a" + "hello")
                            for fi153 := 0; fi153 < 17; fi153++ {
                                p0 = p0
                                var v154 int16 = 825
                                _ = fi153
                                _ = v154
                            }
                            var v155 float32 = 0.1
                            _ = v152
                            _ = v155
                        }
                        v156 := GenS7{F0: 65535, F1: genFn0(uint16(p0), 1), F2: 134}
                        for fi157 := 0; fi157 < 4; fi157++ {
                            var v158 int32 = 42
                            _ = fi157
                            _ = v158
                        }
                        _ = v156
                    } else {
                        if (1.5 == float64(0.0)) {
                            p1 = "foo bar"
                            v159 := GenS20{F0: 255, F1: p1, F2: (fi7 * fi7)}
                            p1 = (p1 + "hello")
                            _ = v159
                        } else {
                            var v160 uint64 = 2
                            p1 = "\n\t"
                            switch (fi7 - (fi7 - fi7)) {
                            case 32768:
                                var v161 uint16 = 623
                                _ = v161
                            }
                            _ = v160
                        }
                        p1 = "\n\t"
                    }
                    p1 = ""
                    p1 = ("a" + p1)
                }
            }
        }
        _ = fi7
    }
    var v162 float32 = 0.1
    _ = p0
    _ = p1
    _ = v162
    _ = v6
    return (100 == -34)
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := []string{"", "\n\t", ("a" + (("\n\t" + "foo bar") + "\n\t"))}
    var v4 int64 = 834
    var v5 uint32 = 643
    if true {
        fmt.Println(v5)
        var v163 int16 = -32768
        _ = v163
    }
    delete(v1, (v5 | (v5 >> ((775) & 63))))
    if true {
        v5 += v5
    } else {
        v4 += 1
        for fi164 := 0; fi164 < 1; fi164++ {
            fmt.Println(fi164)
            _ = fi164
        }
        v165 := v1[4294967295]
        _ = v165
    }
    var v166 float64 = (1e10 / -2.25)
    fmt.Println(v166)
    if (0 > 127) {
        v167 := GenS6{F0: 32767, F1: 255, F2: -968}
        v5 = (v5 | v5)
        _ = v167
    } else {
        delete(v1, (v5 % ((2147483648) | 1)))
        var v168 uint64 = 4294967295
        _ = v168
    }
    var v169 int = 9223372036854775807
    v5 += 2
    for fi170 := 0; fi170 < 16; fi170++ {
        switch v169 {
        case 127:
            for fi171 := 0; fi171 < 14; fi171++ {
                var v172 uint64 = 18446744073709551615
                fmt.Println(v5, fi171, fi170)
                _ = fi171
                _ = v172
            }
        case 2147483647:
            var v173 float64 = ((v166 / v166) + v166)
            _ = v173
        default:
            if ((v4 <= v4) || (0 > 2)) {
                var v174 uint8 = 128
                if true {
                    for fi175 := 0; fi175 < 8; fi175++ {
                        var v176 string = "foo bar"
                        var v177 int32 = -1
                        v178 := GenS22{F0: 343, F1: v174}
                        _ = fi175
                        _ = v176
                        _ = v177
                        _ = v178
                    }
                    var v179 uint = 899
                    _ = v179
                } else {
                    if false {
                        switch "a" {
                        case "foo bar":
                            var v180 bool = false
                            _ = v180
                        case "a":
                            var v181 int64 = v4
                            _ = v181
                        }
                        switch v169 {
                        case -820:
                            v182 := GenS17{F0: 42}
                            _ = v182
                        default:
                            fmt.Println(v166)
                        }
                        var v183 int = fi170
                        _ = v183
                    } else {
                        v174 += v174
                        switch 363 {
                        case -9223372036854775808:
                            var v184 uint8 = v174
                            _ = v184
                        case 127:
                            v4 -= (v4 | (-(v4)))
                        }
                    }
                    v185 := GenS10{F0: 42}
                    switch -528 {
                    case 9223372036854775807:
                        if false {
                            for fi186 := 0; fi186 < 19; fi186++ {
                                if (127 < v174) {
                                    var v187 int64 = (v4 / ((v4) | 1))
                                    if true {
                                        fmt.Println(v169)
                                        if false {
                                            var v188 int32 = 0
                                            v187 += (v187 << ((1) & 63))
                                            v189 := GenS23{F0: 68, F1: 42}
                                            _ = v188
                                            _ = v189
                                        } else {
                                            _ = len(v3)
                                            v187 = v187
                                        }
                                    }
                                    var v190 int64 = (v187 + -1)
                                    v191 := GenS6{F0: v169, F1: 0, F2: 913}
                                    var v192 string = "\n\t"
                                    v193 := GenS11{F0: 537}
                                    delete(v2, fi186)
                                    var v194 bool = (false || true)
                                    _ = v187
                                    _ = v190
                                    _ = v191
                                    _ = v192
                                    _ = v193
                                    _ = v194
                                }
                                switch ("\n\t" + "foo bar") {
                                case "\n\t":
                                    fmt.Println(fi186, v4)
                                case "a":
                                    var v195 int16 = 0
                                    _ = v195
                                }
                                v196 := GenS4{F0: 647, F1: "\n\t"}
                                _ = fi186
                                _ = v196
                            }
                        }
                    case 42:
                        if true {
                            fmt.Println(v174)
                            var v197 int16 = -1
                            for fi198 := 0; fi198 < 17; fi198++ {
                                switch ("\n\t" + (("foo bar" + ("foo bar" + "\n\t")) + "a")) {
                                case "\n\t":
                                    var v199 string = ("\n\t" + (("" + "\n\t") + "hello"))
                                    _ = v199
                                case "hello":
                                    if ((277 == 900) && (!true)) {
                                        v3 = append(v3, ((("" + "a") + "") + ((("\n\t" + ("" + "foo bar")) + "") + "")))
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                        if (fi198 > (fi198 * (fi170 + 0))) {
                                            v166 += (float64((-2.25 - 1.5)) + (-2.25 * 3.14159))
                                            if (int32(v174) <= 42) {
                                                var v200 uint8 = 60
                                                if true {
                                                    switch (fi170 | (v169 / ((9223372036854775807) | 1))) {
                                                    case 32767:
                                                        v201 := GenS8{F0: -2147483648}
                                                        _ = v201
                                                    case 127:
                                                        v4 = -1
                                                    }
                                                }
                                                v200 += v174
                                                _ = v200
                                            }
                                        } else {
                                            if (v197 != v197) {
                                                var v202 string = ""
                                                var v203 uint32 = v5
                                                v203 = (^(v5))
                                                _ = v202
                                                _ = v203
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = ""
                                            }
                                            v197 += v197
                                            var v204 int64 = v4
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v205 := GenS20{F0: 563, F1: "hello", F2: 127}
                                            _ = v204
                                            _ = v205
                                        }
                                        switch ((("\n\t" + "\n\t") + "a") + "") {
                                        case "a":
                                            if (!false) {
                                                v166 = 0.1
                                                v3 = append(v3, "")
                                                var v206 int8 = -43
                                                _ = v206
                                            } else {
                                                v207 := GenS20{F0: 65535, F1: "foo bar", F2: (fi170 ^ -890)}
                                                var v208 int16 = (v197 ^ (v197 | (v197 ^ v197)))
                                                _ = v207
                                                _ = v208
                                            }
                                        case "hello":
                                            var v209 string = "hello"
                                            _ = v209
                                        case "":
                                            v210 := GenS4{F0: 0, F1: "a"}
                                            _ = v210
                                        }
                                    }
                                default:
                                    var v211 uint16 = uint16(v174)
                                    _ = v211
                                }
                                _ = fi198
                            }
                            _ = v197
                        } else {
                            v166 = ((0.1 - float64(0.0)) * v166)
                            var v212 float64 = v166
                            _ = v212
                        }
                    default:
                        v5 += v5
                    }
                    _ = v185
                }
                _ = v174
            }
        }
        _ = fi170
    }
    delete(v2, v169)
    if (v4 == genFn0(0, 65535)) {
        delete(v1, v5)
    }
    var v213 bool = (v166 == 0.1)
    v213 = true
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v214 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v166:", v166)
    fmt.Println("v169:", v169)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v213:", v213)
    fmt.Println("v214:", v214)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v166
    _ = v169
    _ = v2
    _ = v213
    _ = v214
    _ = v3
    _ = v4
    _ = v5
}
