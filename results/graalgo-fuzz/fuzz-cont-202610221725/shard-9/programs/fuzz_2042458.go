package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int8; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint; F1 int16; F2 bool }
type GenS8 struct { F0 int; F1 int16; F2 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool; F1 int; F2 int64 }
type GenS11 struct { F0 uint; F1 uint; F2 int }
type GenS12 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS13 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"a", ("\n\t" + "foo bar"), "foo bar"}
    var v3 int = 128
    v4 := []int{v3}
    var v5 string = ("foo bar" + "\n\t")
    if ((!true) && (2 == 42)) {
        switch v3 {
        case -1:
            v3 = v3
        case 503:
            v3 = 0
        default:
            v6 := GenS1{F0: 2, F1: (v3 + -9223372036854775808), F2: 42}
            _ = v6
        }
    }
    v3 += (^(v3))
    for fi7 := 0; fi7 < 4; fi7++ {
        if (0 < 18446744073709551615) {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 uint32 = 389
                fmt.Println(v3)
                _ = fi8
                _ = v9
            }
            if (v5 >= "\n\t") {
                switch (v3 - 0) {
                case 32768:
                    for fi10 := 0; fi10 < 5; fi10++ {
                        var v11 int8 = 1
                        if (v11 > 127) {
                            var v12 uint8 = 42
                            _ = v12
                        }
                        delete(v1, 42)
                        v5 = (v5 + "a")
                        var v13 float32 = -2.25
                        _ = fi10
                        _ = v11
                        _ = v13
                    }
                default:
                    v14 := GenS1{F0: 2147483647, F1: (^(v3)), F2: -1}
                    _ = v14
                }
                v4 = append(v4, (fi7 - 1))
                fmt.Println(v3, fi7)
            } else {
                v3 = (v3 & -2147483648)
            }
        } else {
            var v15 uint8 = 33
            switch ("" + "a") {
            case "\n\t":
                v16 := GenS1{F0: 156, F1: fi7, F2: -1}
                _ = v16
            case "a":
                var v17 int32 = 1
                _ = v17
            }
            switch (v3 | (v3 * fi7)) {
            case -245:
                v18 := GenS1{F0: uint32(v15), F1: 42, F2: -374}
                _ = v18
            case 417:
                var v19 uint = 65535
                _ = v19
            default:
                var v20 int16 = 781
                _ = v20
            }
            _ = v15
        }
        switch 4294967295 {
        case 32768:
            v3 = (v3 - v3)
        }
        _ = fi7
    }
    switch (v3 % (((v3 | v3)) | 1)) {
    case 9223372036854775807:
        for fi21 := 0; fi21 < 9; fi21++ {
            switch (v3 % (((fi21 ^ fi21)) | 1)) {
            case 2147483647:
                v22 := GenS2{F0: (v5 + "hello"), F1: 4294967295}
                _ = v22
            case 0:
                for fi23 := 0; fi23 < 8; fi23++ {
                    _ = len(v1)
                    fmt.Println(fi21, fi21)
                    fmt.Println(v5, fi23)
                    v3 = (v3 | (v3 & fi23))
                    if false {
                        switch (fi23 & 9223372036854775807) {
                        case -9223372036854775808:
                            var v24 float32 = 3.14159
                            _ = v24
                        case -188:
                            switch fi23 {
                            case 32767:
                                switch v5 {
                                case "":
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                    if true {
                                        v25 := GenS1{F0: 42, F1: fi23, F2: -1}
                                        _ = v25
                                    }
                                    v26 := GenS2{F0: "a", F1: 509}
                                    var v27 uint32 = 547
                                    v27 -= v27
                                    fmt.Println(fi21)
                                    v28 := v1[259]
                                    v1[v27] = v28
                                    fmt.Println(fi21)
                                    var v29 uint16 = 65535
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                case "\n\t":
                                    var v30 uint = 761
                                    _ = v30
                                case "foo bar":
                                    v3 += 1
                                default:
                                    if false {
                                        v31 := GenS3{F0: 4294967295, F1: 1}
                                        switch -9223372036854775808 {
                                        case -9223372036854775808:
                                            v32 := GenS1{F0: 65535, F1: 0, F2: 42}
                                            _ = v32
                                        case 32767:
                                            v33 := GenS4{F0: 1, F1: 132, F2: 56}
                                            _ = v33
                                        default:
                                            v5 = (v5 + v5)
                                        }
                                        _ = v31
                                    }
                                }
                            }
                        case 32768:
                            var v34 float32 = 1e10
                            _ = v34
                        }
                    } else {
                        v3 -= 0
                        if false {
                            fmt.Println(v5, v5)
                            switch "hello" {
                            case "a":
                                var v35 int8 = -71
                                _ = v35
                            case "\n\t":
                                if false {
                                    v36 := GenS5{F0: 0}
                                    v3 -= (fi21 * v3)
                                    _ = v36
                                } else {
                                    if (fi21 < fi23) {
                                        var v37 int16 = -1
                                        _ = v37
                                    }
                                    v38 := v1[2]
                                    _ = v38
                                }
                            default:
                                v39 := v1[2147483647]
                                _ = v39
                            }
                        } else {
                            switch ((v5 + v5) + v5) {
                            case "hello":
                                fmt.Println(v5)
                                var v40 float64 = -2.25
                                switch 40 {
                                case 32767:
                                    v41 := GenS5{F0: 255}
                                    _ = v41
                                case 0:
                                    _ = len(v4)
                                }
                                _ = v40
                            case "foo bar":
                                v3 += (fi23 % (((-(fi21))) | 1))
                            default:
                                _ = len(v4)
                            }
                            _ = len(v4)
                        }
                        v4 = append(v4, 793)
                        var v42 uint = 42
                        _ = v42
                    }
                    _ = fi23
                }
            case 42:
                v43 := GenS2{F0: ("foo bar" + ("a" + "")), F1: 0}
                _ = v43
            default:
                v3 -= 9223372036854775807
            }
            v44 := GenS5{F0: 65535}
            _ = fi21
            _ = v44
        }
    case -144:
        var v45 string = v5
        _ = v45
    case 4294967295:
        var v46 uint64 = 276
        _ = v46
    default:
        for fi47 := 0; fi47 < 2; fi47++ {
            fmt.Println(v5, v5)
            var v48 int64 = 669
            _ = fi47
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        if false {
            v50 := GenS1{F0: 42, F1: (-(fi49)), F2: -1}
            switch (v3 << ((0) & 63)) {
            case 4294967295:
                v51 := GenS6{F0: -188, F1: -128}
                _ = v51
            case 32768:
                var v52 uint8 = 42
                _ = v52
            }
            var v53 int = fi49
            _ = v50
            _ = v53
        } else {
            var v54 uint16 = 255
            if false {
                var v55 int8 = -128
                var v56 int = 927
                for fi57 := 0; fi57 < 6; fi57++ {
                    var v58 uint8 = 1
                    switch 32767 {
                    case 42:
                        var v59 int8 = v55
                        v2 = append(v2, v5)
                        _ = v59
                    default:
                        v60 := GenS7{F0: 0, F1: -32768, F2: (255 != 42)}
                        _ = v60
                    }
                    if (uint64(v54) <= 716) {
                        v3 += fi49
                    } else {
                        if (true && true) {
                            v61 := GenS5{F0: (v54 + 32767)}
                            if (!(831 < 42)) {
                                v62 := v1[4294967295]
                                if (v58 >= (v58 >> ((597) & 63))) {
                                    v58 += (v58 ^ (v58 - 128))
                                    if true {
                                        var v63 float64 = (-2.25 - 1e10)
                                        _ = v63
                                    } else {
                                        v1[4294967295] = "foo bar"
                                        v55 += (v55 / ((v55) | 1))
                                        var v64 uint32 = 42
                                        _ = v64
                                    }
                                }
                                switch v5 {
                                case "hello":
                                    v65 := GenS7{F0: 0, F1: 32767, F2: false}
                                    _ = v65
                                case "\n\t":
                                    if len(v2) > 0 {
                                        v2[0] = v5
                                    }
                                case "a":
                                    v5 = ("hello" + v5)
                                default:
                                    if false {
                                        delete(v1, 2147483647)
                                    } else {
                                        v5 = v62
                                        fmt.Println(v3, fi49, v62)
                                    }
                                }
                                _ = v62
                            }
                            var v66 int64 = 936
                            v5 = v5
                            _ = v61
                            _ = v66
                        } else {
                            switch 42 {
                            case 0:
                                var v67 uint64 = 774
                                _ = v67
                            case 127:
                                if false {
                                    var v68 uint16 = (v54 + v54)
                                    v3 = fi57
                                    v56 -= (^(v56))
                                    _ = v68
                                } else {
                                    var v69 int64 = -9223372036854775808
                                    _ = len(v2)
                                    v70 := GenS8{F0: v3, F1: 42, F2: (^(v55))}
                                    _ = v69
                                    _ = v70
                                }
                            case 940:
                                fmt.Println(v58)
                            }
                            if true {
                                var v71 int32 = 42
                                v54 = (v54 & 1)
                                v3 += v56
                                _ = v71
                            } else {
                                if (177 >= 0) {
                                    switch v5 {
                                    case "":
                                        var v72 int32 = 325
                                        fmt.Println(v3, fi49, v72)
                                        v58 = v58
                                        if len(v4) > 0 {
                                            v4[0] = fi49
                                        }
                                        _ = v72
                                    case "foo bar":
                                        switch int(v58) {
                                        case 474:
                                            fmt.Println(fi57, v58, v58)
                                        case 169:
                                            v73 := GenS2{F0: ((v5 + (v5 + v5)) + "hello"), F1: 11}
                                            _ = v73
                                        default:
                                            var v74 int8 = v55
                                            if len(v4) > 0 {
                                                v4[0] = v56
                                            }
                                            _ = v74
                                        }
                                    case "\n\t":
                                        switch v5 {
                                        case "hello":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "foo bar":
                                            switch v3 {
                                            case 700:
                                                if false {
                                                    fmt.Println(v5, v55, v5)
                                                    if false {
                                                        if ((!true) || ((0.1 < 1.5) || true)) {
                                                            var v75 uint32 = 0
                                                            v58 += (v58 * (v58 >> ((4294967295) & 63)))
                                                            _ = v75
                                                        }
                                                    }
                                                } else {
                                                    var v76 int8 = v55
                                                    _ = v76
                                                }
                                            case 4294967295:
                                                var v77 int64 = 0
                                                v77 -= (^(v77))
                                                if (0 == 529) {
                                                    fmt.Println(v56)
                                                }
                                                _ = v77
                                            case -610:
                                                v78 := GenS9{F0: false}
                                                _ = v78
                                            }
                                        default:
                                            v55 = v55
                                        }
                                        v79 := GenS1{F0: 42, F1: (-(v56)), F2: -754}
                                        var v80 int16 = 538
                                        _ = v79
                                        _ = v80
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        v58 += v58
                                    case "":
                                        if true {
                                            if ((v5 + (v5 + v5)) != v5) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v81 int = fi49
                                                _ = v81
                                            } else {
                                                v55 += (v55 / (((v55 / ((v55) | 1))) | 1))
                                                v82 := GenS10{F0: false, F1: fi49, F2: -1}
                                                _ = v82
                                            }
                                            var v83 string = ""
                                            switch int(v55) {
                                            case -325:
                                                delete(v1, 42)
                                                var v84 uint64 = 0
                                                if (4294967295 == 4294967295) {
                                                    switch ((v5 + "") + v5) {
                                                    case "hello":
                                                        fmt.Println(v5, fi57, v84)
                                                    case "a":
                                                        var v85 uint = 416
                                                        _ = v85
                                                    }
                                                    v83 = v83
                                                    v84 = 18446744073709551615
                                                    v86 := GenS11{F0: uint(v58), F1: 2, F2: (-(v3))}
                                                    var v87 uint16 = v54
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    fmt.Println(v55, v83)
                                                }
                                                _ = v84
                                            }
                                            var v88 float64 = ((1e10 + ((0.0 / (3.14159 / 0.1)) * (0.0 - float64((0.0 * 1.5))))) - 1.5)
                                            _ = v83
                                            _ = v88
                                        } else {
                                            switch (("a" + "hello") + v5) {
                                            case "foo bar":
                                                var v89 bool = true
                                                _ = v89
                                            default:
                                                var v90 uint64 = 2
                                                _ = v90
                                            }
                                        }
                                    }
                                }
                            }
                            var v91 uint64 = 42
                            v92 := GenS3{F0: 65535, F1: -373}
                            _ = v91
                            _ = v92
                        }
                        if ((v5 + v5) == v5) {
                            if ((fi49 % (((v3 % ((127) | 1))) | 1)) >= 0) {
                                delete(v1, 65535)
                            }
                        } else {
                            var v93 uint8 = v58
                            v4 = append(v4, (v3 - (-(fi57))))
                            switch "" {
                            case "":
                                fmt.Println(fi57)
                            case "a":
                                if (0 > 0) {
                                    v3 = v3
                                    v94 := v1[831]
                                    fmt.Println(fi49, v93)
                                    _ = v94
                                } else {
                                    var v95 int32 = -1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v95
                                }
                            }
                            _ = v93
                        }
                        v3 += (v56 << ((4) & 63))
                    }
                    _ = fi57
                    _ = v58
                }
                _ = v55
                _ = v56
            }
            _ = v54
        }
        _ = fi49
    }
    switch v3 {
    case -759:
        v5 = (((v5 + ((v5 + "") + v5)) + v5) + v5)
    }
    fmt.Println(v3, v5)
    var v96 float64 = float64((((0.1 / 1.0) - (3.14159 + float32(1.5))) / -2.25))
    delete(v1, 62)
    fmt.Println(v5, v96, v5)
    for fi97 := 0; fi97 < 20; fi97++ {
        fmt.Println(v96, v96, v3)
        switch v5 {
        case "\n\t":
            for fi98 := 0; fi98 < 4; fi98++ {
                switch "\n\t" {
                case "a":
                    var v99 uint16 = 1
                    _ = v99
                default:
                    v96 += float64((1e10 / 1e10))
                }
                v5 = "a"
                _ = fi98
            }
        case "foo bar":
            switch v5 {
            case "":
                v5 = ("\n\t" + (v5 + v5))
            case "hello":
                if false {
                    v96 -= (v96 / 1e10)
                } else {
                    var v100 bool = (v5 >= (("hello" + v5) + "a"))
                    fmt.Println(v3, v3)
                    fmt.Println(v96, v96, fi97)
                    _ = v100
                }
            }
        default:
            if false {
                if (255 >= 18446744073709551615) {
                    for fi101 := 0; fi101 < 13; fi101++ {
                        var v102 uint32 = 4294967295
                        _ = fi101
                        _ = v102
                    }
                }
                v3 = 139
            } else {
                var v103 float64 = 1.5
                var v104 int32 = 0
                _ = v103
                _ = v104
            }
        }
        _ = fi97
    }
    v105 := GenS12{F0: 0, F1: 2, F2: 0}
    var v106 int16 = 1
    switch v5 {
    case "foo bar":
        v96 = ((v96 - (1.5 - (float64(-2.25) - v96))) - v96)
    case "hello":
        var v107 int = v3
        _ = v107
    default:
        fmt.Println(v5)
    }
    v4 = append(v4, (v3 ^ -1))
    if false {
        fmt.Println(v106, v96, v106)
        switch (v3 % ((-1) | 1)) {
        case 158:
            switch (v3 % (((v3 << ((0) & 63))) | 1)) {
            case 128:
                var v108 int32 = 2147483647
                _ = v108
            case 9223372036854775807:
                switch v5 {
                case "":
                    for fi109 := 0; fi109 < 3; fi109++ {
                        v96 -= (v96 * 3.14159)
                        for fi110 := 0; fi110 < 16; fi110++ {
                            v4 = append(v4, 800)
                            v2 = append(v2, "hello")
                            fmt.Println(v106, v96)
                            _ = fi110
                        }
                        v111 := GenS13{F0: -1}
                        _ = fi109
                        _ = v111
                    }
                }
            default:
                var v112 int = v3
                _ = v112
            }
        case 374:
            if true {
                switch "a" {
                case "\n\t":
                    v113 := GenS8{F0: (v3 - -990), F1: (v106 >> ((0) & 63)), F2: -1}
                    _ = v113
                case "a":
                    v1[361] = v5
                case "hello":
                    v114 := GenS13{F0: 42}
                    _ = v114
                }
                var v115 uint64 = 345
                _ = v115
            } else {
                var v116 uint = 1
                _ = v116
            }
        case 1:
            if len(v4) > 0 {
                v4[0] = v3
            }
        }
        var v117 uint64 = 0
        _ = v117
    }
    if ((0 >= 767) && (1 > 4294967295)) {
        if true {
            var v118 int16 = v106
            _ = v118
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v96
}
