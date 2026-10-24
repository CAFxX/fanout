package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 string; F2 int }
type GenS5 struct { F0 string; F1 uint64; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 int64 }
type GenS11 struct { F0 uint64; F1 bool }
type GenS12 struct { F0 uint32; F1 uint64 }
type GenS13 struct { F0 int8; F1 uint32 }
type GenS14 struct { F0 uint16; F1 int8 }
type GenS15 struct { F0 int64; F1 int8 }
type GenS16 struct { F0 string; F1 uint64; F2 int8 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 uint32; F1 string; F2 uint16 }
func genFn0(p0 int8) int32 {
    v6 := []int{}
    v7 := GenS1{F0: 1, F1: false}
    var v8 int8 = (p0 & 127)
    var v9 uint8 = 0
    if ((v8 << ((4294967295) & 63)) < (v8 << ((255) & 63))) {
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 float64 = ((3.14159 / 3.14159) - -2.25)
            var v12 float32 = 1.5
            if (false && (0 != uint32(v9))) {
                fmt.Println(p0, p0, v8)
                var v13 float32 = (v12 - (v12 * ((v12 - -2.25) - v12)))
                v6 = append(v6, fi10)
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1() uint32 {
    v14 := []int{-209, -871}
    v14 = append(v14, 4294967295)
    if false {
        switch 127 {
        case -9223372036854775808:
            if len(v14) > 0 {
                v14[0] = 0
            }
        case 400:
            var v15 int64 = -9223372036854775808
            _ = v15
        default:
            var v16 int = -2147483648
            _ = v16
        }
        var v17 int64 = -463
        switch 0 {
        case -9223372036854775808:
            var v18 int32 = genFn0(127)
            _ = v18
        case 0:
            v17 -= -410
        }
        _ = v17
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        switch 0 {
        case 32767:
            var v20 int = fi19
            _ = v20
        case 0:
            var v21 bool = (921 >= -2147483648)
            _ = v21
        }
        v22 := GenS1{F0: 65535, F1: ((-(fi19)) <= (fi19 * fi19))}
        var v23 uint = 52
        _ = fi19
        _ = v22
        _ = v23
    }
    v14 = append(v14, 128)
    _ = v14
    return 2147483648
}

func main() {
    var v1 int8 = -128
    var v2 uint = 384
    v3 := make(map[int]int)
    var v4 int = 42
    v5 := []string{"a", ("" + ("a" + "\n\t")), ("foo bar" + "a")}
    v3[v4] = (v4 % ((v4) | 1))
    v24 := GenS1{F0: (v2 << ((65535) & 63)), F1: (!true)}
    fmt.Println(v4)
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v4, v2)
        var v26 int = 127
        _ = fi25
        _ = v26
    }
    v1 -= -67
    if (2147483647 <= 2147483647) {
        fmt.Println(v4, v2)
        var v27 string = "hello"
        _ = v27
    } else {
        if true {
            var v28 int8 = 42
            fmt.Println(v28, v2)
            switch (^(v4)) {
            case 128:
                var v29 float64 = 3.14159
                _ = v29
            case -9223372036854775808:
                switch "a" {
                case "hello":
                    var v30 uint64 = 1
                    _ = v30
                default:
                    if (true || ((!((1 <= -32768) || true)) || true)) {
                        v1 += v1
                        for fi31 := 0; fi31 < 10; fi31++ {
                            v4 -= -1
                            v5 = append(v5, "hello")
                            _ = fi31
                        }
                        if (true || false) {
                            if false {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if (2 == 704) {
                                    if ((0 != 527) && ("foo bar" >= "\n\t")) {
                                        if (1000 == int16(v2)) {
                                            for fi32 := 0; fi32 < 8; fi32++ {
                                                if false {
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(v4, v4, v2)
                                                    case "":
                                                        var v33 uint8 = 255
                                                        _ = v33
                                                    }
                                                    if (318 >= 0) {
                                                        var v34 int32 = 0
                                                        var v35 bool = false
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        v1 = v28
                                                        if len(v5) > 0 {
                                                            v5[0] = "foo bar"
                                                        }
                                                        if false {
                                                            v36 := v3[fi32]
                                                            for fi37 := 0; fi37 < 11; fi37++ {
                                                                _ = len(v5)
                                                                switch ("hello" + "hello") {
                                                                case "":
                                                                    v36 = (^(v4))
                                                                case "\n\t":
                                                                    v38 := GenS1{F0: (v2 % ((v2) | 1)), F1: (-1 < 9223372036854775807)}
                                                                    _ = v38
                                                                default:
                                                                    var v39 string = "hello"
                                                                    _ = v39
                                                                }
                                                                fmt.Println(v36, v28)
                                                                v1 += 42
                                                                _ = fi37
                                                            }
                                                            switch fi32 {
                                                            case 9223372036854775807:
                                                                v4 = 2147483647
                                                            case 107:
                                                                for fi40 := 0; fi40 < 3; fi40++ {
                                                                    var v41 uint = (v2 << ((v2) & 63))
                                                                    var v42 int8 = (v28 ^ v1)
                                                                    fmt.Println(v42)
                                                                    _ = fi40
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                            case 4294967295:
                                                                v28 += 61
                                                            }
                                                            _ = v36
                                                        }
                                                    }
                                                    v1 += 0
                                                } else {
                                                    switch "foo bar" {
                                                    case "":
                                                        v1 += -1
                                                    }
                                                }
                                                _ = fi32
                                            }
                                            var v43 uint64 = 4294967295
                                            for fi44 := 0; fi44 < 19; fi44++ {
                                                for fi45 := 0; fi45 < 18; fi45++ {
                                                    var v46 float64 = (1e10 - (0.1 - (3.14159 - 0.1)))
                                                    _ = fi45
                                                    _ = v46
                                                }
                                                if false {
                                                    v47 := GenS2{F0: 255, F1: -1}
                                                    _ = v47
                                                }
                                                _ = fi44
                                            }
                                            _ = v43
                                        }
                                        v4 = v4
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = ""
                                    }
                                } else {
                                    fmt.Println(v1, v1)
                                    if (32767 >= -1) {
                                        for fi48 := 0; fi48 < 17; fi48++ {
                                            var v49 int = fi48
                                            _ = fi48
                                            _ = v49
                                        }
                                    }
                                }
                            }
                            if (!true) {
                                if false {
                                    for fi50 := 0; fi50 < 4; fi50++ {
                                        if false {
                                            if (!((v4 ^ (-(v4))) == (fi50 & fi50))) {
                                                var v51 int64 = 343
                                                if false {
                                                    delete(v3, v4)
                                                    var v52 uint = 255
                                                    _ = v52
                                                } else {
                                                    v5 = append(v5, ("hello" + ""))
                                                    switch v4 {
                                                    case -9223372036854775808:
                                                        if true {
                                                            v53 := GenS3{F0: 18446744073709551615}
                                                            v51 = -1
                                                            _ = v53
                                                        } else {
                                                            delete(v3, (^(fi50)))
                                                        }
                                                    case 956:
                                                        var v54 int32 = -1
                                                        _ = v54
                                                    default:
                                                        v51 = (v51 / (((v51 % (((v51 / ((0) | 1))) | 1))) | 1))
                                                    }
                                                    v4 = 127
                                                }
                                                _ = v51
                                            }
                                            if (!((743 == 42) || (-2147483648 < genFn0(1)))) {
                                                if (v2 >= (v2 % (((^(v2))) | 1))) {
                                                    if (v4 < (v4 / ((fi50) | 1))) {
                                                        var v55 uint64 = 2
                                                        fmt.Println(v55)
                                                        fmt.Println(fi50, v1, v1)
                                                        _ = v55
                                                    }
                                                    var v56 string = "a"
                                                    v57 := GenS2{F0: 1, F1: 0}
                                                    _ = v56
                                                    _ = v57
                                                }
                                                var v58 uint16 = 78
                                                fmt.Println(v1, v2)
                                                _ = v58
                                            }
                                        } else {
                                            for fi59 := 0; fi59 < 3; fi59++ {
                                                var v60 uint64 = 689
                                                switch (fi50 % ((v4) | 1)) {
                                                case -311:
                                                    var v61 bool = ("" <= ("foo bar" + ""))
                                                    _ = v61
                                                case 128:
                                                    switch 825 {
                                                    case -576:
                                                        v62 := GenS4{F0: v2, F1: "foo bar", F2: (fi50 * (fi59 * (v4 % ((264) | 1))))}
                                                        _ = v62
                                                    case 1:
                                                        switch ("a" + ("hello" + "")) {
                                                        case "hello":
                                                            switch (fi50 * (v4 * v4)) {
                                                            case 128:
                                                                if (!false) {
                                                                    v63 := GenS5{F0: "hello", F1: v60, F2: 42}
                                                                    _ = len(v3)
                                                                    _ = v63
                                                                } else {
                                                                    if false {
                                                                        var v64 uint32 = 65535
                                                                        if len(v5) > 0 {
                                                                            v5[0] = "hello"
                                                                        }
                                                                        if true {
                                                                            var v65 uint64 = 1
                                                                            if len(v5) > 0 {
                                                                                v5[0] = ""
                                                                            }
                                                                            switch ((("hello" + "a") + ("hello" + "\n\t")) + "hello") {
                                                                            case "foo bar":
                                                                                v66 := GenS4{F0: (v2 >> (((v2 >> ((v2) & 63))) & 63)), F1: ("hello" + ""), F2: (v4 & fi50)}
                                                                                _ = v66
                                                                            case "\n\t":
                                                                                v5 = append(v5, "")
                                                                            case "":
                                                                                if (true && ((-(fi50)) >= (fi59 ^ -376))) {
                                                                                    _ = len(v3)
                                                                                    fmt.Println(v4)
                                                                                }
                                                                            }
                                                                            _ = v65
                                                                        } else {
                                                                            var v67 uint64 = v60
                                                                            var v68 uint64 = (v60 - 2)
                                                                            if true {
                                                                                if (!false) {
                                                                                    var v69 uint32 = v64
                                                                                    if true {
                                                                                        v70 := GenS6{F0: 32767}
                                                                                        var v71 int64 = -9223372036854775808
                                                                                        v72 := GenS7{F0: 0, F1: v71}
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                        _ = v72
                                                                                    } else {
                                                                                        switch fi50 {
                                                                                        case 1:
                                                                                            v73 := GenS4{F0: 255, F1: ("\n\t" + "hello"), F2: (fi59 / ((int(v2)) | 1))}
                                                                                            _ = v73
                                                                                        case 32767:
                                                                                            var v74 int8 = -1
                                                                                            _ = v74
                                                                                        case -488:
                                                                                            var v75 uint = (^(v2))
                                                                                            _ = v75
                                                                                        }
                                                                                        fmt.Println(v4, v28, v4)
                                                                                        fmt.Println(v68, v67)
                                                                                    }
                                                                                    _ = v69
                                                                                }
                                                                            } else {
                                                                                var v76 uint16 = 42
                                                                                _ = v76
                                                                            }
                                                                            var v77 bool = true
                                                                            var v78 string = (("hello" + "a") + "")
                                                                            fmt.Println(v78, v68, fi50)
                                                                            _ = v67
                                                                            _ = v68
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                        _ = v64
                                                                    } else {
                                                                        var v79 uint16 = 347
                                                                        if (false || true) {
                                                                            switch fi59 {
                                                                            case 1:
                                                                                v4 += -198
                                                                            }
                                                                            v79 += (v79 - v79)
                                                                            fmt.Println(v4, v60)
                                                                        } else {
                                                                            v4 -= v4
                                                                        }
                                                                        var v80 uint32 = 65535
                                                                        _ = v79
                                                                        _ = v80
                                                                    }
                                                                    if false {
                                                                        v5 = append(v5, "hello")
                                                                    }
                                                                }
                                                            case -976:
                                                                v60 = (v60 + v60)
                                                            }
                                                        case "":
                                                            v2 = 18446744073709551615
                                                        default:
                                                            if (4294967295 > v60) {
                                                                v81 := GenS2{F0: uint8(v60), F1: -32768}
                                                                _ = v81
                                                            }
                                                        }
                                                    case -604:
                                                        if (!true) {
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v82 uint = v2
                                                                _ = v82
                                                            }
                                                        } else {
                                                            if ((v60 + (v60 & 1)) >= (v60 - (v60 | 306))) {
                                                                v60 -= (v60 << ((v2) & 63))
                                                                if false {
                                                                    v60 = v60
                                                                    var v83 uint8 = 2
                                                                    fmt.Println(v83, v28, v4)
                                                                    _ = v83
                                                                } else {
                                                                    v3[(fi59 + fi50)] = v4
                                                                    var v84 uint32 = 2
                                                                    _ = v84
                                                                }
                                                                switch 32767 {
                                                                case 1:
                                                                    v3[(fi50 >> ((v2) & 63))] = (v4 + (fi59 - v4))
                                                                case -2147483648:
                                                                    var v85 float32 = float32(1.5)
                                                                    _ = v85
                                                                default:
                                                                    var v86 float32 = 3.14159
                                                                    _ = v86
                                                                }
                                                                switch "hello" {
                                                                case "":
                                                                    if (!true) {
                                                                        var v87 bool = false
                                                                        _ = v87
                                                                    }
                                                                case "\n\t":
                                                                    fmt.Println(v1)
                                                                case "a":
                                                                    if true {
                                                                        v2 = (v2 << ((1) & 63))
                                                                    } else {
                                                                        v3[(v4 >> ((4294967295) & 63))] = v4
                                                                        v4 = 127
                                                                        var v88 uint16 = 848
                                                                        v5 = append(v5, ("foo bar" + ""))
                                                                        _ = v88
                                                                    }
                                                                }
                                                                v28 -= (-(v28))
                                                            }
                                                            v28 = (v1 << (((v2 ^ v2)) & 63))
                                                            v89 := GenS3{F0: (v60 + (v60 & (v60 * v60)))}
                                                            _ = v89
                                                        }
                                                    default:
                                                        fmt.Println(v60)
                                                    }
                                                }
                                                _ = len(v5)
                                                _ = fi59
                                                _ = v60
                                            }
                                            if (float32(((((3.14159 * -2.25) * (((float64(-2.25) / ((3.14159 + float64((1e10 / 1.5))) - 1.5)) + 1e10) + (float64(1.5) + (1e10 - 0.1)))) / 1e10) * 3.14159)) < (0.1 * 1.5)) {
                                                if (18446744073709551615 < 2) {
                                                    fmt.Println(v28)
                                                    v2 = (v2 / ((655) | 1))
                                                    var v90 float32 = float32(-2.25)
                                                    _ = v90
                                                } else {
                                                    v2 = (v2 & v2)
                                                }
                                                var v91 int = fi50
                                                v91 -= 4294967295
                                                _ = v91
                                            } else {
                                                v1 = v1
                                                v28 -= (v1 % (((^(v28))) | 1))
                                            }
                                        }
                                        var v92 int8 = v1
                                        _ = fi50
                                        _ = v92
                                    }
                                    var v93 int32 = genFn0(v28)
                                    _ = v93
                                }
                                for fi94 := 0; fi94 < 13; fi94++ {
                                    fmt.Println(v2)
                                    _ = fi94
                                }
                                fmt.Println(v1)
                            } else {
                                var v95 float64 = -2.25
                                _ = v95
                            }
                            if len(v5) > 0 {
                                v5[0] = ("hello" + ("foo bar" + (("foo bar" + "") + "a")))
                            }
                        } else {
                            for fi96 := 0; fi96 < 6; fi96++ {
                                v28 = v28
                                if true {
                                    switch 9223372036854775807 {
                                    case -1:
                                        for fi97 := 0; fi97 < 13; fi97++ {
                                            switch "foo bar" {
                                            case "\n\t":
                                                fmt.Println(v2, v1)
                                            case "foo bar":
                                                var v98 uint8 = 255
                                                _ = v98
                                            case "a":
                                                var v99 int32 = 1
                                                _ = v99
                                            }
                                            var v100 string = "a"
                                            v1 -= (v28 ^ 1)
                                            _ = fi97
                                            _ = v100
                                        }
                                    case -308:
                                        var v101 string = "hello"
                                        _ = v101
                                    case 40:
                                        var v102 int8 = (v1 ^ -24)
                                        _ = v102
                                    default:
                                        v4 = 4294967295
                                    }
                                    fmt.Println(fi96, v28, fi96)
                                    v4 = (-(v4))
                                } else {
                                    for fi103 := 0; fi103 < 9; fi103++ {
                                        if false {
                                            if (!(!(!true))) {
                                                v2 = (^(v2))
                                            } else {
                                                fmt.Println(fi103, v28)
                                                v104 := GenS3{F0: 1}
                                                _ = v104
                                            }
                                            var v105 float32 = -2.25
                                            _ = v105
                                        } else {
                                            var v106 uint32 = 2147483647
                                            v106 -= (^(v106))
                                            if true {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v107 uint32 = genFn1()
                                                _ = v107
                                            } else {
                                                var v108 uint8 = 42
                                                switch fi103 {
                                                case 2147483647:
                                                    var v109 string = "foo bar"
                                                    _ = v109
                                                case 776:
                                                    var v110 int = int(v106)
                                                    _ = v110
                                                default:
                                                    v4 = (v4 * fi96)
                                                }
                                                var v111 int = 1
                                                _ = v108
                                                _ = v111
                                            }
                                            _ = v106
                                        }
                                        var v112 uint16 = 1
                                        v113 := GenS8{F0: -972}
                                        _ = fi103
                                        _ = v112
                                        _ = v113
                                    }
                                    var v114 uint32 = 513
                                    fmt.Println(v1, v28, v4)
                                    _ = v114
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi96
                            }
                        }
                    } else {
                        var v115 int8 = 1
                        for fi116 := 0; fi116 < 8; fi116++ {
                            var v117 bool = false
                            _ = fi116
                            _ = v117
                        }
                        switch "" {
                        case "a":
                            switch "hello" {
                            case "hello":
                                if false {
                                    if ("a" < ("hello" + ("foo bar" + "a"))) {
                                        v28 -= (v28 & v1)
                                        if true {
                                            v118 := v3[v4]
                                            _ = v118
                                        }
                                    } else {
                                        v5 = append(v5, "")
                                    }
                                    if true {
                                        for fi119 := 0; fi119 < 10; fi119++ {
                                            v2 -= (v2 ^ (^(v2)))
                                            var v120 string = "a"
                                            _ = fi119
                                            _ = v120
                                        }
                                        switch v4 {
                                        case 805:
                                            var v121 string = "\n\t"
                                            _ = v121
                                        }
                                        if (192 < 1000) {
                                            v122 := GenS7{F0: genFn0(v1), F1: 0}
                                            if len(v5) > 0 {
                                                v5[0] = "hello"
                                            }
                                            _ = v122
                                        }
                                    } else {
                                        if (false || false) {
                                            if len(v5) > 0 {
                                                v5[0] = "a"
                                            }
                                        } else {
                                            if (!(42 >= 0)) {
                                                v1 -= (v1 / (((^(v115))) | 1))
                                                v1 += (v28 >> ((v2) & 63))
                                            }
                                            var v123 uint = (^(v2))
                                            _ = v123
                                        }
                                        v2 += (^(v2))
                                        for fi124 := 0; fi124 < 1; fi124++ {
                                            v115 = 127
                                            var v125 int = int(v115)
                                            _ = fi124
                                            _ = v125
                                        }
                                    }
                                    v126 := GenS2{F0: 128, F1: 0}
                                    _ = v126
                                }
                            case "":
                                v127 := GenS7{F0: 42, F1: 42}
                                _ = v127
                            }
                        default:
                            v4 = (-(v4))
                        }
                        _ = v115
                    }
                }
            case -939:
                if true {
                    if false {
                        fmt.Println(v1, v28, v28)
                        v4 = (v4 & v4)
                        v2 = v2
                    } else {
                        var v128 uint = (v2 - (v2 * v2))
                        v129 := GenS9{F0: 2}
                        _ = v128
                        _ = v129
                    }
                    var v130 string = ""
                    fmt.Println(v28, v28)
                    _ = v130
                }
            default:
                var v131 uint16 = 2
                _ = v131
            }
            _ = v28
        }
        for fi132 := 0; fi132 < 2; fi132++ {
            fmt.Println(v1, v4)
            var v133 uint32 = 2147483648
            var v134 float32 = 1e10
            _ = fi132
            _ = v133
            _ = v134
        }
    }
    if ((42 == 42) && (false && false)) {
        for fi135 := 0; fi135 < 18; fi135++ {
            v136 := GenS8{F0: 811}
            var v137 uint = v2
            var v138 int64 = -9223372036854775808
            _ = fi135
            _ = v136
            _ = v137
            _ = v138
        }
    }
    if true {
        v1 -= (v1 & v1)
        v1 = v1
    }
    v139 := GenS5{F0: "\n\t", F1: 756, F2: 126}
    if false {
        v140 := GenS10{F0: v1, F1: 9223372036854775807}
        if (!false) {
            fmt.Println(v1, v2, v4)
            v141 := GenS10{F0: v1, F1: -9223372036854775808}
            _ = v141
        } else {
            var v142 uint = 1
            var v143 int = 62
            _ = v142
            _ = v143
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v140
    } else {
        switch v4 {
        case 186:
            _ = len(v5)
        case 32767:
            var v144 uint64 = 18446744073709551615
            _ = v144
        default:
            var v145 uint32 = 2147483648
            _ = v145
        }
        v146 := GenS11{F0: 227, F1: true}
        _ = v146
    }
    var v147 int8 = (v1 - 127)
    var v148 uint32 = 4294967295
    v2 = v2
    for fi149 := 0; fi149 < 10; fi149++ {
        if ((v148 / (((v148 % (((v148 | v148)) | 1))) | 1)) == (v148 - 838)) {
            v1 += v147
        }
        var v150 float32 = (0.1 / 1.0)
        var v151 int8 = 1
        _ = fi149
        _ = v150
        _ = v151
    }
    switch v4 {
    case 32767:
        _ = len(v5)
    case -9223372036854775808:
        if (((v147 * int8(v4)) >= -54) && ((1e10 - float64((0.1 - (1e10 * 3.14159)))) != float64((-2.25 / -2.25)))) {
            switch v4 {
            case -660:
                for fi152 := 0; fi152 < 6; fi152++ {
                    var v153 uint16 = 32767
                    v2 -= (v2 & v2)
                    _ = fi152
                    _ = v153
                }
            case 128:
                switch ("hello" + "foo bar") {
                case "hello":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "\n\t":
                    v1 = (v147 | v1)
                }
            default:
                if true {
                    for fi154 := 0; fi154 < 9; fi154++ {
                        if true {
                            var v155 float32 = ((-2.25 / 1e10) * 0.0)
                            _ = v155
                        } else {
                            if false {
                                for fi156 := 0; fi156 < 10; fi156++ {
                                    v2 = (v2 << ((v2) & 63))
                                    switch 0 {
                                    case 2147483647:
                                        var v157 int8 = int8(v4)
                                        _ = v157
                                    case 519:
                                        switch "" {
                                        case "hello":
                                            v158 := GenS12{F0: genFn1(), F1: 4294967295}
                                            _ = v158
                                        }
                                    case 9223372036854775807:
                                        v159 := v3[fi154]
                                        _ = v159
                                    }
                                    var v160 int = (v4 / ((-2147483648) | 1))
                                    _ = fi156
                                    _ = v160
                                }
                                switch fi154 {
                                case 4294967295:
                                    v4 = v4
                                case 1:
                                    v2 += (v2 - 240)
                                case 393:
                                    var v161 uint64 = 2
                                    _ = v161
                                }
                            }
                        }
                        delete(v3, v4)
                        _ = fi154
                    }
                }
            }
            v162 := GenS6{F0: 2}
            _ = v162
        } else {
            if true {
                v148 = (v148 ^ 64)
            } else {
                var v163 float32 = 0.0
                switch v4 {
                case 1:
                    if false {
                        switch (-(v4)) {
                        case -9223372036854775808:
                            if true {
                                if ((v2 - 91) < (v2 & (v2 | 2))) {
                                    fmt.Println(v1, v163, v148)
                                    var v164 int16 = 278
                                    _ = v164
                                } else {
                                    fmt.Println(v148, v148)
                                }
                                v165 := GenS1{F0: 42, F1: (!true)}
                                switch v4 {
                                case -232:
                                    if len(v5) > 0 {
                                        v5[0] = (("a" + (("" + "\n\t") + ("\n\t" + ("" + "hello")))) + ("foo bar" + "hello"))
                                    }
                                case -566:
                                    fmt.Println(v2)
                                case -113:
                                    for fi166 := 0; fi166 < 18; fi166++ {
                                        fmt.Println(v2, fi166, v1)
                                        v163 -= ((v163 * 1.5) + ((v163 + ((0.1 / 1.0) / (-2.25 - (v163 + (float32(1.5) * v163))))) + ((v163 + v163) * (0.0 / v163))))
                                        _ = fi166
                                    }
                                }
                                _ = v165
                            }
                        default:
                            v167 := GenS13{F0: (v147 | v1), F1: 4294967295}
                            _ = v167
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v163 -= v163
                    }
                case 0:
                    switch "\n\t" {
                    case "hello":
                        var v168 string = "foo bar"
                        _ = v168
                    case "":
                        if true {
                            switch (v4 | (v4 * 661)) {
                            case -31:
                                var v169 int32 = 1
                                _ = v169
                            case 128:
                                v148 = genFn1()
                            case 245:
                                for fi170 := 0; fi170 < 2; fi170++ {
                                    var v171 bool = true
                                    v148 = v148
                                    _ = fi170
                                    _ = v171
                                }
                            default:
                                if ((!(4294967295 > 18446744073709551615)) || false) {
                                    v172 := GenS13{F0: v1, F1: 1}
                                    _ = v172
                                }
                            }
                            for fi173 := 0; fi173 < 14; fi173++ {
                                v5 = append(v5, "hello")
                                v163 = v163
                                _ = fi173
                            }
                            for fi174 := 0; fi174 < 9; fi174++ {
                                v1 = (v147 << (((v2 & (^(v2)))) & 63))
                                switch 0 {
                                case 32768:
                                    switch ("foo bar" + "hello") {
                                    case "a":
                                        for fi175 := 0; fi175 < 12; fi175++ {
                                            if (-614 == 266) {
                                                var v176 int8 = (v147 >> ((v2) & 63))
                                                v4 = fi175
                                                v177 := GenS1{F0: v2, F1: true}
                                                _ = v176
                                                _ = v177
                                            }
                                            var v178 float64 = -2.25
                                            v147 = (v1 / ((63) | 1))
                                            v163 = 0.1
                                            v147 += (v1 * (^(v1)))
                                            _ = fi175
                                            _ = v178
                                        }
                                    default:
                                        if true {
                                            v148 = (v148 >> ((4294967295) & 63))
                                            var v179 float32 = v163
                                            _ = v179
                                        } else {
                                            if (9223372036854775807 < 1) {
                                                v147 = -116
                                                v4 += -906
                                            } else {
                                                if false {
                                                    switch "\n\t" {
                                                    case "a":
                                                        switch v4 {
                                                        case -712:
                                                            if true {
                                                                switch "hello" {
                                                                case "":
                                                                    fmt.Println(v2, v163, v1)
                                                                case "foo bar":
                                                                    for fi180 := 0; fi180 < 13; fi180++ {
                                                                        if true {
                                                                            fmt.Println(v1)
                                                                            var v181 int32 = -2147483648
                                                                            var v182 uint16 = 317
                                                                            _ = v181
                                                                            _ = v182
                                                                        }
                                                                        v147 = (v147 >> (((^(v2))) & 63))
                                                                        _ = fi180
                                                                    }
                                                                }
                                                                if (0 == 4294967295) {
                                                                    fmt.Println(v147, v2, v163)
                                                                }
                                                                v183 := GenS14{F0: 614, F1: 7}
                                                                _ = v183
                                                            }
                                                        case -2147483648:
                                                            v184 := GenS6{F0: 1}
                                                            _ = v184
                                                        case -553:
                                                            v185 := v3[128]
                                                            _ = v185
                                                        default:
                                                            switch int(v148) {
                                                            case 1:
                                                                for fi186 := 0; fi186 < 19; fi186++ {
                                                                    v2 = (v2 << (((v2 + v2)) & 63))
                                                                    _ = fi186
                                                                }
                                                            default:
                                                                v148 -= v148
                                                            }
                                                        }
                                                    case "hello":
                                                        v147 -= (v147 & v1)
                                                    }
                                                    for fi187 := 0; fi187 < 16; fi187++ {
                                                        fmt.Println(fi174, v148, v1)
                                                        var v188 uint = 42
                                                        switch fi174 {
                                                        case 127:
                                                            var v189 uint32 = genFn1()
                                                            _ = v189
                                                        case 0:
                                                            var v190 uint = v188
                                                            _ = v190
                                                        default:
                                                            v163 -= (1e10 + float32(3.14159))
                                                            v163 = v163
                                                            switch 127 {
                                                            case 32768:
                                                                var v191 uint32 = (v148 * (v148 - (v148 * genFn1())))
                                                                _ = v191
                                                            default:
                                                                v192 := GenS13{F0: (v1 | (v1 ^ 127)), F1: v148}
                                                                _ = v192
                                                            }
                                                        }
                                                        _ = fi187
                                                        _ = v188
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v193 := GenS14{F0: 238, F1: 42}
                                                    _ = v193
                                                case "\n\t":
                                                    v194 := GenS2{F0: 2, F1: 198}
                                                    _ = v194
                                                default:
                                                    v195 := GenS15{F0: -945, F1: -114}
                                                    _ = v195
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = fi174
                            }
                        } else {
                            if (!true) {
                                fmt.Println(v1)
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case "a":
                        if ("" == "hello") {
                            if true {
                                v196 := GenS16{F0: (("hello" + ("foo bar" + (("\n\t" + "hello") + ("\n\t" + "hello")))) + "\n\t"), F1: uint64(v2), F2: v147}
                                _ = v196
                            } else {
                                _ = len(v3)
                            }
                            var v197 uint32 = (v148 >> ((1) & 63))
                            switch v4 {
                            case 4294967295:
                                v148 = 2
                            case 32768:
                                v1 = int8(v197)
                            }
                            _ = v197
                        } else {
                            var v198 uint = 2
                            if (1.5 != (((0.0 + 3.14159) - 1.5) + 1e10)) {
                                v199 := GenS17{F0: "a"}
                                _ = len(v5)
                                v200 := GenS8{F0: 2147483647}
                                _ = v199
                                _ = v200
                            }
                            switch (v4 / ((v4) | 1)) {
                            case -608:
                                v201 := GenS16{F0: "a", F1: 42, F2: v147}
                                _ = v201
                            default:
                                var v202 uint64 = 985
                                _ = v202
                            }
                            _ = v198
                        }
                    default:
                        var v203 bool = (v163 >= 1e10)
                        _ = v203
                    }
                }
                _ = v163
            }
        }
    case 9223372036854775807:
        v204 := GenS18{F0: 42, F1: "a", F2: 0}
        _ = v204
    default:
        v205 := GenS17{F0: "\n\t"}
        _ = v205
    }
    switch "" {
    case "foo bar":
        switch v4 {
        case 27:
            v206 := GenS6{F0: 688}
            _ = v206
        case 0:
            v148 -= (^(v148))
        case 1:
            fmt.Println(v4, v2)
        }
    case "hello":
        var v207 uint8 = 117
        _ = v207
    }
    v1 += v147
    fmt.Println("v1:", v1)
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v139
    _ = v147
    _ = v148
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
