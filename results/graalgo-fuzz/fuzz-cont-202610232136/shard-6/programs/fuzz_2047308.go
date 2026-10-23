package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int32; F1 int16; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int; F2 uint }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 string }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int64; F1 uint64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 string; F1 uint16; F2 int64 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint64; F1 uint32 }
type GenS16 struct { F0 uint64; F1 uint16 }
type GenS17 struct { F0 uint64; F1 int; F2 bool }
type GenS18 struct { F0 int8; F1 int64 }
type GenS19 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int64, p1 int64) int64 {
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            if false {
                switch ((("\n\t" + "") + ("" + "hello")) + ("\n\t" + "a")) {
                case "hello":
                    p0 -= 42
                default:
                    p0 += 0
                }
            } else {
                var v10 string = ""
                _ = v10
            }
            v11 := GenS1{F0: (-(p0)), F1: 42, F2: 3}
            fmt.Println(fi9, p0, p1)
            if false {
                var v12 float64 = (-2.25 * 0.0)
                v13 := []int{127, (fi9 & fi8), fi9}
                var v14 int16 = 42
                v14 += (v14 << ((18446744073709551615) & 63))
                var v15 bool = true
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
            if true {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(p1)
                case "hello":
                    v16 := GenS1{F0: int64(v7), F1: 1, F2: 262}
                    _ = v16
                case "a":
                    p1 = p0
                }
            } else {
                fmt.Println(p0)
                if true {
                    var v17 uint8 = 1
                    v18 := []int{(fi8 / ((int(v17)) | 1)), (fi8 << ((65535) & 63))}
                    if len(v18) > 0 {
                        v18[0] = (-(fi9))
                    }
                    _ = v17
                    _ = v18
                } else {
                    v19 := []int{(fi8 | 0), 0}
                    v7 += v7
                    _ = v19
                }
            }
            _ = fi9
            _ = v11
        }
        _ = fi8
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v7
    return (p1 * p0)
}
func genFn1() uint8 {
    v20 := GenS1{F0: 42, F1: 1, F2: 660}
    if true {
        var v21 int8 = 100
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 uint16 = 255
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 16; fi25++ {
                v26 := GenS1{F0: 423, F1: 65535, F2: 0}
                v27 := []int{(-(fi22))}
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi22
        }
        if (true && false) {
            v28 := []int{9223372036854775807, -291}
            var v29 uint64 = 0
            _ = v28
            _ = v29
        } else {
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := GenS2{F0: fi30, F1: false}
                for fi32 := 0; fi32 < 2; fi32++ {
                    switch "foo bar" {
                    case "a":
                        fmt.Println(fi30, fi30, fi30)
                    default:
                        v33 := GenS2{F0: -2147483648, F1: (false && true)}
                        _ = v33
                    }
                    fmt.Println(fi32)
                    v34 := GenS1{F0: 0, F1: 2, F2: 2}
                    _ = fi32
                    _ = v34
                }
                if true {
                    if true {
                        v35 := GenS2{F0: fi30, F1: true}
                        _ = v35
                    }
                } else {
                    switch (-(fi30)) {
                    case 127:
                        var v36 uint8 = 2
                        _ = v36
                    }
                    v37 := GenS1{F0: 9223372036854775807, F1: 65535, F2: 4294967295}
                    _ = v37
                }
                _ = fi30
                _ = v31
            }
        }
        _ = v21
    } else {
        var v38 uint8 = 255
        if (((-2.25 / -2.25) - 0.1) >= float32((1.5 * 0.0))) {
            if true {
                v38 -= v38
                v39 := GenS2{F0: 9223372036854775807, F1: true}
                for fi40 := 0; fi40 < 15; fi40++ {
                    var v41 int = (fi40 - 128)
                    if true {
                        switch fi40 {
                        case 1:
                            v42 := []int{(v41 - (fi40 << ((42) & 63)))}
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 4; fi43++ {
                            if false {
                                var v44 uint = 42
                                _ = v44
                            }
                            var v45 uint8 = (v38 | (v38 - v38))
                            fmt.Println(v45, fi40)
                            _ = fi43
                            _ = v45
                        }
                    }
                    v46 := GenS1{F0: 1, F1: 42, F2: 2}
                    _ = fi40
                    _ = v41
                    _ = v46
                }
                _ = v39
            } else {
                v38 += 42
                var v47 uint8 = v38
                for fi48 := 0; fi48 < 5; fi48++ {
                    v49 := []int{fi48, fi48, fi48, fi48}
                    var v50 uint64 = 9
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                _ = v47
            }
            fmt.Println(v38)
            v38 = 215
        } else {
            v38 += 0
            var v51 uint16 = 884
            fmt.Println(v38)
            _ = v51
        }
        var v52 int8 = 0
        _ = v38
        _ = v52
    }
    if ((127 < -25) && false) {
        if false {
            v53 := []int{779, 177, 42, -425}
            var v54 bool = (-9223372036854775808 < 374)
            switch "\n\t" {
            case "a":
                v55 := GenS1{F0: -1, F1: 2147483648, F2: 1}
                _ = v55
            case "\n\t":
                if (!((42 != 549) && true)) {
                    for fi56 := 0; fi56 < 16; fi56++ {
                        v53 = append(v53, (fi56 / (((fi56 << ((42) & 63))) | 1)))
                        _ = fi56
                    }
                    v57 := GenS3{F0: 42, F1: -872, F2: v54}
                    _ = v57
                } else {
                    for fi58 := 0; fi58 < 9; fi58++ {
                        v59 := GenS1{F0: 1, F1: 2, F2: 42}
                        v60 := GenS1{F0: -965, F1: 4294967295, F2: 4294967295}
                        for fi61 := 0; fi61 < 2; fi61++ {
                            var v62 uint32 = 2
                            switch ("" + ("a" + "hello")) {
                            case "\n\t":
                                var v63 uint64 = 18446744073709551615
                                _ = v63
                            case "":
                                v62 -= (^(v62))
                            default:
                                v54 = (!(!v54))
                                if (((329 <= 1) && ("a" <= "foo bar")) || true) {
                                    if (v54 || (true && (599 == 32767))) {
                                        var v64 uint8 = 2
                                        var v65 uint8 = v64
                                        var v66 uint64 = 592
                                        var v67 float32 = (3.14159 / 1.0)
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                    v62 = v62
                                    if (825 >= 257) {
                                        v54 = v54
                                        if v54 {
                                            v62 -= 776
                                            var v68 bool = (42 == -1)
                                            switch ("" + "a") {
                                            case "foo bar":
                                                switch "" {
                                                case "\n\t":
                                                    if (802 >= 18446744073709551615) {
                                                        if true {
                                                            v68 = false
                                                        } else {
                                                            var v69 float64 = (1.5 * 3.14159)
                                                            _ = v69
                                                        }
                                                        fmt.Println(v54)
                                                        var v70 uint16 = 1
                                                        v70 = 255
                                                        if false {
                                                            var v71 uint32 = 4294967295
                                                            var v72 uint = 18446744073709551615
                                                            _ = len(v53)
                                                            _ = v71
                                                            _ = v72
                                                        } else {
                                                            fmt.Println(fi61)
                                                            v73 := GenS4{F0: fi61}
                                                            if (127 <= 0) {
                                                                if v54 {
                                                                    switch 128 {
                                                                    case -1:
                                                                        _ = len(v53)
                                                                    case 756:
                                                                        var v74 int = -314
                                                                        _ = v74
                                                                    default:
                                                                        v53 = append(v53, (-(fi61)))
                                                                    }
                                                                    fmt.Println(v70, v54)
                                                                } else {
                                                                    if ((v62 + v62) <= (v62 ^ (v62 - v62))) {
                                                                        v70 = (v70 - 411)
                                                                        switch fi61 {
                                                                        case 128:
                                                                            v75 := GenS5{F0: v62, F1: 4294967295, F2: 42}
                                                                            _ = v75
                                                                        case 127:
                                                                            fmt.Println(v70)
                                                                        default:
                                                                            v76 := GenS6{F0: -1, F1: ("\n\t" + "foo bar")}
                                                                            _ = v76
                                                                        }
                                                                        if v54 {
                                                                            v77 := GenS7{F0: 0, F1: ""}
                                                                            _ = v77
                                                                        } else {
                                                                            fmt.Println(v68)
                                                                        }
                                                                    } else {
                                                                        var v78 int16 = 960
                                                                        v62 = 712
                                                                        _ = v78
                                                                    }
                                                                    if len(v53) > 0 {
                                                                        v53[0] = fi58
                                                                    }
                                                                }
                                                                v53 = append(v53, fi61)
                                                            } else {
                                                                v54 = v68
                                                                if len(v53) > 0 {
                                                                    v53[0] = fi61
                                                                }
                                                                var v79 int64 = genFn0(0, 9223372036854775807)
                                                                _ = v79
                                                            }
                                                            _ = v73
                                                        }
                                                        if (0.0 != 3.14159) {
                                                            v68 = (-126 >= -1)
                                                        } else {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v80 int32 = 1
                                                                _ = v80
                                                            case "":
                                                                var v81 int64 = 1
                                                                _ = v81
                                                            }
                                                            switch (fi58 + (fi58 << ((369) & 63))) {
                                                            case 2147483647:
                                                                if len(v53) > 0 {
                                                                    v53[0] = (fi61 | 0)
                                                                }
                                                            case 0:
                                                                if false {
                                                                    var v82 float32 = 3.14159
                                                                    _ = v82
                                                                } else {
                                                                    var v83 uint8 = 128
                                                                    _ = v83
                                                                }
                                                            case -9223372036854775808:
                                                                v84 := GenS6{F0: -1, F1: "hello"}
                                                                _ = v84
                                                            default:
                                                                if len(v53) >= 2 {
                                                                    v53 = v53[1:len(v53)]
                                                                }
                                                            }
                                                        }
                                                        _ = v70
                                                    } else {
                                                        var v85 bool = v54
                                                        v86 := GenS8{F0: -9223372036854775808}
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                case "a":
                                                    var v87 uint32 = v62
                                                    _ = v87
                                                }
                                            default:
                                                if len(v53) >= 2 {
                                                    v53 = v53[1:len(v53)]
                                                }
                                            }
                                            _ = v68
                                        } else {
                                            var v88 string = "a"
                                            v89 := GenS7{F0: 4294967295, F1: "a"}
                                            switch "hello" {
                                            case "\n\t":
                                                if (175 != -2147483648) {
                                                    if v54 {
                                                        if v54 {
                                                            switch fi61 {
                                                            case 4294967295:
                                                                v90 := GenS7{F0: 42, F1: ("hello" + v88)}
                                                                _ = v90
                                                            case 1:
                                                                var v91 uint = 135
                                                                _ = v91
                                                            case -9223372036854775808:
                                                                v88 = v88
                                                                if ((0.1 * (((-2.25 + 0.0) - (1.5 / -2.25)) * 1.5)) < float32(3.14159)) {
                                                                    var v92 uint64 = 4294967295
                                                                    var v93 int64 = 9223372036854775807
                                                                    _ = v92
                                                                    _ = v93
                                                                }
                                                                fmt.Println(v62)
                                                            default:
                                                                switch fi58 {
                                                                case 9223372036854775807:
                                                                    _ = len(v53)
                                                                case 116:
                                                                    if true {
                                                                        fmt.Println(v54, v88)
                                                                        fmt.Println(v54, v88)
                                                                        if ("\n\t" > v88) {
                                                                            if true {
                                                                                fmt.Println(v88, v62, v54)
                                                                                if (v62 != (v62 - 2147483648)) {
                                                                                    switch (v88 + "hello") {
                                                                                    case "hello":
                                                                                        var v94 uint64 = 2
                                                                                        _ = v94
                                                                                    case "\n\t":
                                                                                        v54 = (!false)
                                                                                    }
                                                                                    v62 += 2
                                                                                    v95 := GenS9{F0: 32767, F1: v88}
                                                                                    _ = v95
                                                                                }
                                                                            }
                                                                            switch fi58 {
                                                                            case 32767:
                                                                                switch "hello" {
                                                                                case "":
                                                                                    v96 := GenS6{F0: 0, F1: v88}
                                                                                    _ = v96
                                                                                case "a":
                                                                                    v62 = 0
                                                                                }
                                                                            case 549:
                                                                                if v54 {
                                                                                    switch 4294967295 {
                                                                                    case 844:
                                                                                        var v97 string = ((v88 + v88) + v88)
                                                                                        _ = v97
                                                                                    case 127:
                                                                                        if len(v53) >= 2 {
                                                                                            v53 = v53[1:len(v53)]
                                                                                        }
                                                                                    default:
                                                                                        var v98 uint = 0
                                                                                        _ = v98
                                                                                    }
                                                                                    if v54 {
                                                                                        if v54 {
                                                                                            fmt.Println(v88)
                                                                                        }
                                                                                    } else {
                                                                                        var v99 uint16 = 320
                                                                                        v100 := GenS9{F0: (v99 / ((v99) | 1)), F1: v88}
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    }
                                                                                    if true {
                                                                                        switch (fi61 * fi58) {
                                                                                        case 128:
                                                                                            var v101 float32 = float32(1.5)
                                                                                            _ = v101
                                                                                        }
                                                                                        if ((-128 >= 0) && false) {
                                                                                            if len(v53) >= 2 {
                                                                                                v53 = v53[1:len(v53)]
                                                                                            }
                                                                                            if true {
                                                                                                if (2 != 88) {
                                                                                                    var v102 uint16 = 2
                                                                                                    switch v88 {
                                                                                                    case "hello":
                                                                                                        var v103 uint = 18446744073709551615
                                                                                                        _ = v103
                                                                                                    case "a":
                                                                                                        v104 := GenS10{F0: (fi61 ^ 9223372036854775807)}
                                                                                                        _ = v104
                                                                                                    default:
                                                                                                        switch fi61 {
                                                                                                        case 32767:
                                                                                                            var v105 int16 = 42
                                                                                                            _ = v105
                                                                                                        case -1:
                                                                                                            v62 = uint32(v102)
                                                                                                        case 909:
                                                                                                            switch fi58 {
                                                                                                            case -21:
                                                                                                                fmt.Println(v62, fi58)
                                                                                                            case 637:
                                                                                                                if len(v53) >= 2 {
                                                                                                                    v53 = v53[1:len(v53)]
                                                                                                                }
                                                                                                            default:
                                                                                                                switch fi61 {
                                                                                                                case -1:
                                                                                                                    var v106 uint64 = 42
                                                                                                                    _ = v106
                                                                                                                default:
                                                                                                                    switch "a" {
                                                                                                                    case "\n\t":
                                                                                                                        if (1 >= -965) {
                                                                                                                            switch v88 {
                                                                                                                            case "hello":
                                                                                                                                v107 := GenS11{F0: 436, F1: uint64(v102)}
                                                                                                                                _ = v107
                                                                                                                            case "":
                                                                                                                                v108 := GenS6{F0: 42, F1: v88}
                                                                                                                                _ = v108
                                                                                                                            case "a":
                                                                                                                                switch fi58 {
                                                                                                                                case -727:
                                                                                                                                    v88 = v88
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                v53 = append(v53, 1)
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v88 = (("\n\t" + "hello") + v88)
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        if len(v53) >= 2 {
                                                                                                                            v53 = v53[1:len(v53)]
                                                                                                                        }
                                                                                                                    case "a":
                                                                                                                        if len(v53) > 0 {
                                                                                                                            v53[0] = fi58
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v54 = v54
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(v54, v62)
                                                                                                            }
                                                                                                            var v109 float32 = (-2.25 + float32((1e10 * (((0.0 / 0.1) * 3.14159) - 0.1))))
                                                                                                            if len(v53) > 0 {
                                                                                                                v53[0] = fi61
                                                                                                            }
                                                                                                            _ = v109
                                                                                                        default:
                                                                                                            v53 = append(v53, 127)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v102
                                                                                                } else {
                                                                                                    v54 = v54
                                                                                                    var v110 int = (fi58 & 32768)
                                                                                                    _ = v110
                                                                                                }
                                                                                                if len(v53) > 0 {
                                                                                                    v53[0] = fi61
                                                                                                }
                                                                                            }
                                                                                            var v111 bool = (!true)
                                                                                            _ = v111
                                                                                        }
                                                                                    } else {
                                                                                        switch v88 {
                                                                                        case "\n\t":
                                                                                            var v112 uint32 = (v62 | 848)
                                                                                            _ = v112
                                                                                        case "hello":
                                                                                            switch fi61 {
                                                                                            case 42:
                                                                                                fmt.Println(fi58, v88, fi61)
                                                                                            case -1:
                                                                                                fmt.Println(v54, v88, v88)
                                                                                                v113 := GenS12{F0: 2147483647}
                                                                                                v62 = (v62 & (v62 / ((2) | 1)))
                                                                                                _ = v113
                                                                                            case -77:
                                                                                                var v114 uint = 2
                                                                                                _ = v114
                                                                                            default:
                                                                                                if v54 {
                                                                                                    v54 = (-41 >= 67)
                                                                                                    fmt.Println(v54, fi58, v54)
                                                                                                    switch -232 {
                                                                                                    case 1:
                                                                                                        if (!false) {
                                                                                                            v115 := GenS1{F0: -303, F1: (v62 << ((255) & 63)), F2: 208}
                                                                                                            var v116 uint64 = 18446744073709551615
                                                                                                            switch fi58 {
                                                                                                            case 32768:
                                                                                                                v117 := GenS7{F0: 65535, F1: ("hello" + v88)}
                                                                                                                _ = v117
                                                                                                            default:
                                                                                                                v88 = v88
                                                                                                            }
                                                                                                            _ = v115
                                                                                                            _ = v116
                                                                                                        } else {
                                                                                                            v88 = ("hello" + ("foo bar" + v88))
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    var v118 uint64 = 92
                                                                                                    _ = v118
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            if false {
                                                                                                v88 = "hello"
                                                                                            }
                                                                                            var v119 uint16 = 66
                                                                                            _ = v119
                                                                                        } else {
                                                                                            v120 := GenS5{F0: v62, F1: (fi58 << ((787) & 63)), F2: 255}
                                                                                            _ = v120
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch 9223372036854775807 {
                                                                                    case -129:
                                                                                        v121 := GenS10{F0: (fi58 ^ 32767)}
                                                                                        _ = v121
                                                                                    default:
                                                                                        v88 = "a"
                                                                                        v122 := GenS11{F0: -1, F1: 1}
                                                                                        _ = v122
                                                                                    }
                                                                                }
                                                                            case 346:
                                                                                v62 += 4294967295
                                                                            default:
                                                                                var v123 int32 = 42
                                                                                _ = v123
                                                                            }
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                var v124 float32 = -2.25
                                                                                _ = v124
                                                                            case "":
                                                                                if v54 {
                                                                                    switch fi61 {
                                                                                    case -2147483648:
                                                                                        v54 = v54
                                                                                    case 32768:
                                                                                        v125 := GenS7{F0: 2, F1: v88}
                                                                                        v126 := GenS12{F0: 172}
                                                                                        if false {
                                                                                            v127 := GenS6{F0: genFn0(42, -422), F1: "\n\t"}
                                                                                            v62 -= v62
                                                                                            switch "hello" {
                                                                                            case "":
                                                                                                if len(v53) >= 2 {
                                                                                                    v53 = v53[1:len(v53)]
                                                                                                }
                                                                                            case "a":
                                                                                                _ = len(v53)
                                                                                            }
                                                                                            v53 = append(v53, fi58)
                                                                                            _ = v127
                                                                                        }
                                                                                        _ = v125
                                                                                        _ = v126
                                                                                    default:
                                                                                        if len(v53) >= 2 {
                                                                                            v53 = v53[1:len(v53)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if len(v53) > 0 {
                                                                                    v53[0] = fi61
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    switch -65 {
                                                                    case 128:
                                                                        switch (fi58 << ((200) & 63)) {
                                                                        case 1:
                                                                            switch 32768 {
                                                                            case -1:
                                                                                v62 -= v62
                                                                            case 42:
                                                                                if v54 {
                                                                                    var v128 int8 = 109
                                                                                    fmt.Println(v62, v62)
                                                                                    _ = v128
                                                                                } else {
                                                                                    var v129 int = (fi61 * fi58)
                                                                                    _ = v129
                                                                                }
                                                                            case 384:
                                                                                if false {
                                                                                    v54 = v54
                                                                                    var v130 int16 = 32767
                                                                                    var v131 bool = false
                                                                                    _ = v130
                                                                                    _ = v131
                                                                                }
                                                                            }
                                                                        }
                                                                    case -11:
                                                                        var v132 uint = 0
                                                                        _ = v132
                                                                    case 4294967295:
                                                                        if (v62 > (v62 + 609)) {
                                                                            v133 := GenS6{F0: 9223372036854775807, F1: (v88 + "\n\t")}
                                                                            _ = v133
                                                                        } else {
                                                                            if (-2.25 <= ((1.5 * 1e10) - 3.14159)) {
                                                                                switch "hello" {
                                                                                case "":
                                                                                    if (1e10 == 3.14159) {
                                                                                        v54 = ((v54 && v54) || false)
                                                                                        _ = len(v53)
                                                                                        v134 := GenS6{F0: -974, F1: (v88 + ((v88 + v88) + v88))}
                                                                                        _ = v134
                                                                                    } else {
                                                                                        var v135 uint = 65535
                                                                                        _ = v135
                                                                                    }
                                                                                }
                                                                                v62 += v62
                                                                            } else {
                                                                                fmt.Println(v88, fi58)
                                                                            }
                                                                            v54 = true
                                                                            var v136 int16 = -524
                                                                            _ = v136
                                                                        }
                                                                    default:
                                                                        v88 = v88
                                                                        if v54 {
                                                                            v53 = append(v53, (-(fi58)))
                                                                        }
                                                                        fmt.Println(v88, fi58)
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v88 = v88
                                                                        case "foo bar":
                                                                            var v137 string = (v88 + v88)
                                                                            _ = v137
                                                                        default:
                                                                            v88 = v88
                                                                        }
                                                                        if (fi58 >= 1) {
                                                                            var v138 uint16 = 255
                                                                            v139 := GenS13{F0: v88, F1: v138, F2: -1}
                                                                            var v140 uint32 = v62
                                                                            _ = v138
                                                                            _ = v139
                                                                            _ = v140
                                                                        } else {
                                                                            v141 := GenS14{F0: 102}
                                                                            var v142 int32 = 42
                                                                            _ = v141
                                                                            _ = v142
                                                                        }
                                                                        var v143 float64 = 0.1
                                                                        _ = v143
                                                                    }
                                                                }
                                                            }
                                                            _ = len(v53)
                                                            v144 := GenS15{F0: 18446744073709551615, F1: (v62 & v62)}
                                                            _ = v144
                                                        } else {
                                                            switch v88 {
                                                            case "hello":
                                                                switch v88 {
                                                                case "\n\t":
                                                                    v53 = append(v53, (fi61 - fi58))
                                                                }
                                                            case "a":
                                                                var v145 string = "hello"
                                                                _ = v145
                                                            case "foo bar":
                                                                fmt.Println(fi61)
                                                            }
                                                            if ((fi58 % (((fi61 << ((18446744073709551615) & 63))) | 1)) < (fi58 + (^(fi61)))) {
                                                                _ = len(v53)
                                                                if true {
                                                                    v146 := GenS16{F0: 698, F1: 65535}
                                                                    v147 := GenS17{F0: 4294967295, F1: (fi58 & 32767), F2: (!v54)}
                                                                    var v148 int64 = -170
                                                                    _ = v146
                                                                    _ = v147
                                                                    _ = v148
                                                                }
                                                                fmt.Println(fi61)
                                                            } else {
                                                                var v149 float64 = (1.5 + (-2.25 / (1e10 - -2.25)))
                                                                _ = v149
                                                            }
                                                            v150 := GenS1{F0: 1, F1: v62, F2: 203}
                                                            _ = v150
                                                        }
                                                        fmt.Println(v62, fi58)
                                                    }
                                                    if (v54 || ((fi61 + fi61) < -1)) {
                                                        fmt.Println(v62, v54, v88)
                                                        if true {
                                                            if v54 {
                                                                var v151 string = (v88 + "foo bar")
                                                                var v152 int = (fi61 * (fi58 ^ fi58))
                                                                _ = v151
                                                                _ = v152
                                                            }
                                                        }
                                                        fmt.Println(v54)
                                                    }
                                                }
                                            case "foo bar":
                                                v153 := GenS15{F0: 252, F1: 945}
                                                _ = v153
                                            default:
                                                v154 := GenS18{F0: -128, F1: genFn0(-9223372036854775808, 42)}
                                                _ = v154
                                            }
                                            _ = v88
                                            _ = v89
                                        }
                                        v62 = v62
                                        fmt.Println(fi58, v62)
                                    }
                                }
                            }
                            _ = fi61
                            _ = v62
                        }
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                }
            case "":
                if ((-1 == genFn0(-654, 0)) && v54) {
                    var v155 bool = v54
                    for fi156 := 0; fi156 < 3; fi156++ {
                        for fi157 := 0; fi157 < 12; fi157++ {
                            var v158 uint8 = 0
                            _ = fi157
                            _ = v158
                        }
                        _ = fi156
                    }
                    _ = v155
                }
            }
            _ = v53
            _ = v54
        } else {
            for fi159 := 0; fi159 < 16; fi159++ {
                var v160 int16 = 1
                switch (fi159 << ((42) & 63)) {
                case 1:
                    for fi161 := 0; fi161 < 16; fi161++ {
                        fmt.Println(fi161, v160)
                        _ = fi161
                    }
                case -1:
                    v162 := GenS19{F0: 1, F1: 255}
                    _ = v162
                case 334:
                    v163 := []int{32767}
                    _ = v163
                }
                _ = fi159
                _ = v160
            }
        }
        v164 := []int{32768, -795}
        _ = v164
    }
    if true {
        var v165 int32 = -1
        v165 += -78
        _ = v165
    } else {
        var v166 float32 = 3.14159
        v166 += float32((((0.0 * 0.1) / -2.25) + 1e10))
        v167 := []int{127, 128, -2147483648}
        _ = v166
        _ = v167
    }
    _ = v20
    return 255
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.0
    v3 := make(map[int]bool)
    var v4 uint32 = 2147483647
    var v5 float64 = ((((3.14159 - 0.0) - 0.1) / float64((0.0 - 0.1))) - v2)
    v6 := make(map[int]int)
    v5 += v5
    for fi168 := 0; fi168 < 6; fi168++ {
        v4 = v4
        var v169 int16 = -1
        _ = fi168
        _ = v169
    }
    v4 += (^(v4))
    v6[936] = -1
    if false {
        v6[32768] = 32767
        var v170 uint32 = (^(v4))
        _ = v170
    }
    var v171 bool = true
    v171 = true
    v5 = v5
    var v172 int8 = 1
    for fi173 := 0; fi173 < 3; fi173++ {
        fmt.Println(v171)
        _ = fi173
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v171:", v171)
    fmt.Println("v172:", v172)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v171
    _ = v172
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
