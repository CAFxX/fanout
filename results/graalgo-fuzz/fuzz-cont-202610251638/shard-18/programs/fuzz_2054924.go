package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint64; F1 int; F2 bool }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 uint; F1 int32; F2 int64 }
type GenS10 struct { F0 bool; F1 uint }
type GenS11 struct { F0 int; F1 uint16 }
type GenS12 struct { F0 int32; F1 int8; F2 int }

func main() {
    v1 := []string{((("\n\t" + ("a" + (("\n\t" + "a") + "\n\t"))) + "a") + "hello"), "a"}
    v2 := []string{"a", "", "\n\t", ("\n\t" + (("foo bar" + "\n\t") + ("a" + ("foo bar" + ("a" + ("foo bar" + ("\n\t" + "a")))))))}
    var v3 int64 = -9223372036854775808
    v4 := []string{("hello" + "a"), "hello", "hello", (("a" + "a") + (("" + (("foo bar" + "foo bar") + "foo bar")) + "foo bar"))}
    v5 := make(map[uint32]int)
    fmt.Println(v3)
    for fi6 := 0; fi6 < 4; fi6++ {
        if (true && false) {
            var v7 uint32 = 676
            _ = v7
        } else {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: 2147483647, F1: 4294967295}
                _ = fi8
                _ = v9
            }
        }
        var v10 bool = false
        var v11 uint8 = 42
        _ = fi6
        _ = v10
        _ = v11
    }
    if true {
        var v12 bool = true
        for fi13 := 0; fi13 < 11; fi13++ {
            for fi14 := 0; fi14 < 3; fi14++ {
                v12 = v12
                switch 32768 {
                case 0:
                    fmt.Println(fi13, v3)
                case 32768:
                    if v12 {
                        _ = len(v4)
                    } else {
                        switch "\n\t" {
                        case "a":
                            switch fi14 {
                            case -756:
                                switch (fi13 & 1) {
                                case 9223372036854775807:
                                    if (2 != 2) {
                                        v15 := GenS2{F0: 29}
                                        v12 = v12
                                        _ = v15
                                    } else {
                                        v16 := GenS2{F0: 853}
                                        if ((fi13 << ((2) & 63)) != (-(fi14))) {
                                            v5[65535] = (fi14 / ((-263) | 1))
                                            var v17 uint32 = 164
                                            v18 := GenS3{F0: -1, F1: "hello"}
                                            if v12 {
                                                var v19 float32 = float32((((-2.25 * 0.1) + 0.0) / (1e10 * 3.14159)))
                                                v3 += (v3 ^ -57)
                                                v20 := GenS4{F0: v17, F1: 255, F2: 127}
                                                _ = v19
                                                _ = v20
                                            } else {
                                                var v21 float64 = ((0.0 + -2.25) + 3.14159)
                                                _ = v21
                                            }
                                            _ = v17
                                            _ = v18
                                        }
                                        _ = v16
                                    }
                                case 603:
                                    var v22 uint8 = 127
                                    _ = v22
                                case 128:
                                    v12 = (59 <= 42)
                                default:
                                    v3 = (v3 ^ v3)
                                }
                            case -188:
                                _ = len(v1)
                            case -457:
                                v23 := GenS5{F0: 4294967295, F1: 2147483647, F2: false}
                                _ = v23
                            }
                        case "":
                            fmt.Println(fi14)
                        default:
                            var v24 string = "hello"
                            _ = v24
                        }
                        v3 = (v3 / ((-9223372036854775808) | 1))
                        var v25 float64 = 3.14159
                        _ = v25
                    }
                default:
                    v3 += (^(v3))
                }
                var v26 uint64 = 2
                switch "hello" {
                case "":
                    var v27 uint = 596
                    _ = v27
                case "a":
                    v3 = v3
                case "\n\t":
                    switch fi13 {
                    case -760:
                        var v28 float64 = 3.14159
                        _ = v28
                    case 128:
                        if true {
                            var v29 uint16 = 65535
                            v30 := GenS1{F0: 4294967295, F1: v26}
                            fmt.Println(v12, v29)
                            _ = v29
                            _ = v30
                        } else {
                            v26 = 0
                        }
                    case -561:
                        v12 = (v12 && false)
                    }
                default:
                    switch fi13 {
                    case 9223372036854775807:
                        if ("\n\t" == "foo bar") {
                            if false {
                                v26 -= 4294967295
                                var v31 string = "foo bar"
                                _ = v31
                            } else {
                                if v12 {
                                    v32 := GenS3{F0: int64(v26), F1: "\n\t"}
                                    var v33 uint = 521
                                    v33 = (v33 / ((v33) | 1))
                                    _ = v32
                                    _ = v33
                                }
                                _ = len(v4)
                            }
                            v26 = (v26 | v26)
                            if (true && true) {
                                switch "" {
                                case "hello":
                                    v2 = append(v2, (("a" + "") + (("\n\t" + "\n\t") + "hello")))
                                }
                                fmt.Println(v26, fi13)
                                v12 = ((!(42 >= 255)) && false)
                            } else {
                                v34 := GenS3{F0: -1, F1: "foo bar"}
                                v26 += 666
                                _ = v34
                            }
                        }
                    case 127:
                        var v35 uint32 = 2147483647
                        _ = v35
                    case 2147483647:
                        var v36 int = fi13
                        _ = v36
                    }
                }
                v3 += v3
                var v37 int16 = 1
                v38 := GenS5{F0: v26, F1: int(v37), F2: v12}
                _ = fi14
                _ = v26
                _ = v37
                _ = v38
            }
            v12 = (109 < 79)
            switch "hello" {
            case "foo bar":
                if v12 {
                    var v39 uint64 = 997
                    switch 0 {
                    case 1:
                        v12 = (false && v12)
                    default:
                        var v40 int16 = -32768
                        _ = v40
                    }
                    switch (fi13 % (((-(fi13))) | 1)) {
                    case -89:
                        v41 := GenS2{F0: v39}
                        _ = v41
                    case 0:
                        v42 := GenS1{F0: 933, F1: v39}
                        _ = v42
                    case -782:
                        switch "foo bar" {
                        case "":
                            fmt.Println(v12, v39, fi13)
                        case "a":
                            if (60 < 255) {
                                _ = len(v5)
                            } else {
                                v3 -= 42
                                var v43 float64 = 3.14159
                                _ = v43
                            }
                        case "hello":
                            v39 = 377
                        }
                    }
                    _ = v39
                }
            case "\n\t":
                switch fi13 {
                case 128:
                    v12 = ("hello" > "foo bar")
                }
            }
            _ = fi13
        }
        _ = v12
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        v3 = (-(v3))
        fmt.Println(fi44, fi44)
        _ = fi44
    }
    fmt.Println(v3)
    var v45 uint32 = 4294967295
    fmt.Println(v45)
    switch "" {
    case "\n\t":
        for fi46 := 0; fi46 < 19; fi46++ {
            v45 += 2
            _ = fi46
        }
    case "":
        v3 = -9223372036854775808
    default:
        var v47 int16 = 1000
        _ = v47
    }
    var v48 int32 = 1
    for fi49 := 0; fi49 < 18; fi49++ {
        v3 += v3
        if true {
            if true {
                if true {
                    if (false && true) {
                        v3 = (-(v3))
                        for fi50 := 0; fi50 < 15; fi50++ {
                            _ = len(v4)
                            if true {
                                var v51 int32 = v48
                                _ = v51
                            }
                            var v52 float64 = ((((0.1 - 0.1) * 1.5) + -2.25) / (1.5 + 0.0))
                            _ = fi50
                            _ = v52
                        }
                    }
                    switch fi49 {
                    case 784:
                        v3 -= 779
                    default:
                        switch 32 {
                        case -927:
                            for fi53 := 0; fi53 < 19; fi53++ {
                                _ = len(v5)
                                _ = fi53
                            }
                        case 42:
                            var v54 bool = false
                            _ = v54
                        case -813:
                            switch -236 {
                            case -1:
                                v45 = (v45 + (v45 | (^(v45))))
                            case -885:
                                fmt.Println(v3)
                            case -925:
                                v55 := GenS6{F0: true, F1: 255}
                                _ = v55
                            default:
                                switch 4294967295 {
                                case 2147483647:
                                    for fi56 := 0; fi56 < 7; fi56++ {
                                        fmt.Println(v48)
                                        v48 -= v48
                                        switch fi56 {
                                        case 818:
                                            v45 = (^(v45))
                                        default:
                                            v57 := GenS4{F0: (v45 & v45), F1: uint8(v45), F2: 127}
                                            _ = v57
                                        }
                                        _ = fi56
                                    }
                                case -334:
                                    delete(v5, v45)
                                }
                            }
                        default:
                            for fi58 := 0; fi58 < 18; fi58++ {
                                var v59 int64 = 1
                                fmt.Println(v59, fi58, fi49)
                                var v60 uint = 809
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                        }
                    }
                } else {
                    for fi61 := 0; fi61 < 4; fi61++ {
                        var v62 uint8 = 144
                        fmt.Println(v3)
                        if (false && true) {
                            var v63 int8 = 77
                            v48 = 0
                            var v64 uint = 0
                            v64 = 694
                            _ = v63
                            _ = v64
                        }
                        v65 := GenS7{F0: 0}
                        v45 -= v45
                        _ = fi61
                        _ = v62
                        _ = v65
                    }
                    switch (fi49 | (^(fi49))) {
                    case -426:
                        var v66 string = "hello"
                        _ = v66
                    }
                }
                switch ("a" + "foo bar") {
                case "":
                    var v67 uint64 = 18446744073709551615
                    _ = v67
                }
                delete(v5, (^(v45)))
            } else {
                if true {
                    v68 := GenS5{F0: 1, F1: (fi49 | 32768), F2: (!(!false))}
                    _ = v68
                }
                for fi69 := 0; fi69 < 16; fi69++ {
                    v3 = -9223372036854775808
                    fmt.Println(v48, v48)
                    v70 := GenS2{F0: 18446744073709551615}
                    _ = fi69
                    _ = v70
                }
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi71 := 0; fi71 < 7; fi71++ {
                v3 += v3
                delete(v5, 481)
                var v72 uint = 0
                if ((float64((1e10 - (-2.25 * 1e10))) - ((-2.25 * (1.5 * (1e10 - (0.1 / -2.25)))) - (0.0 * (1e10 - -2.25)))) != (0.0 - 0.1)) {
                    if (32767 < 777) {
                        switch (fi71 / ((fi71) | 1)) {
                        case 1:
                            if (3.14159 == 0.0) {
                                if true {
                                    if false {
                                        v72 = (v72 / ((v72) | 1))
                                    } else {
                                        v73 := v5[v45]
                                        var v74 uint = v72
                                        _ = v73
                                        _ = v74
                                    }
                                    fmt.Println(fi71, v48, fi49)
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                }
                                if false {
                                    var v75 bool = false
                                    if false {
                                        v76 := GenS7{F0: int8(v48)}
                                        v75 = v75
                                        _ = v76
                                    }
                                    var v77 int16 = -662
                                    _ = v75
                                    _ = v77
                                }
                                v78 := GenS3{F0: -1, F1: "hello"}
                                _ = v78
                            }
                        case 443:
                            if true {
                                v72 = 983
                                if false {
                                    fmt.Println(fi49, v72)
                                    v5[(v45 & (v45 >> ((v72) & 63)))] = (fi49 - fi49)
                                } else {
                                    v45 = (^(v45))
                                }
                            } else {
                                if true {
                                    if (v48 < (v48 + v48)) {
                                        if false {
                                            var v79 int32 = v48
                                            _ = v79
                                        }
                                    }
                                    var v80 int = (-(fi71))
                                    v81 := GenS8{F0: 1, F1: (fi71 >> (((v72 | v72)) & 63))}
                                    _ = v80
                                    _ = v81
                                }
                                v82 := GenS9{F0: (v72 << ((v72) & 63)), F1: v48, F2: (-(v3))}
                                _ = v82
                            }
                        default:
                            fmt.Println(v45)
                        }
                    }
                    switch "" {
                    case "":
                        if ((0 < 18446744073709551615) || false) {
                            if (1e10 < float64(0.0)) {
                                v83 := v5[v45]
                                v45 = v45
                                v84 := GenS5{F0: 0, F1: (v83 ^ fi71), F2: false}
                                _ = v83
                                _ = v84
                            }
                        }
                    case "hello":
                        var v85 float32 = 0.0
                        _ = v85
                    case "foo bar":
                        switch (fi49 / ((fi71) | 1)) {
                        case 1:
                            switch -803 {
                            case 0:
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    v48 += -257
                                    fmt.Println(fi71, v48)
                                } else {
                                    v72 -= (^(v72))
                                    fmt.Println(v72, fi49, fi49)
                                }
                            case -1:
                                var v86 float32 = (1e10 + 0.1)
                                _ = v86
                            case 9223372036854775807:
                                var v87 uint = v72
                                if (4294967295 < 981) {
                                    fmt.Println(fi49, v45, v72)
                                    v45 -= 1
                                    var v88 uint16 = 0
                                    _ = v88
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        _ = len(v5)
                                    case "hello":
                                        v3 += v3
                                    default:
                                        var v89 int16 = -32768
                                        _ = v89
                                    }
                                    v48 = 377
                                }
                                v48 += 2147483647
                                _ = v87
                            }
                        case 32768:
                            v72 -= (v72 << ((419) & 63))
                        case 203:
                            if true {
                                if true {
                                    v3 = v3
                                }
                                var v90 int64 = -1
                                switch "\n\t" {
                                case "":
                                    v3 += (v3 << (((^(v72))) & 63))
                                }
                                switch (fi49 >> (((^(v72))) & 63)) {
                                case -954:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 128:
                                    if (!(float64((1e10 * 0.0)) >= 0.0)) {
                                        v90 -= int64(v45)
                                        switch (fi71 & (fi71 << ((v72) & 63))) {
                                        case 1:
                                            v91 := GenS7{F0: -1}
                                            _ = v91
                                        case -275:
                                            var v92 uint64 = 42
                                            _ = v92
                                        case -9223372036854775808:
                                            v90 -= v3
                                        }
                                        if false {
                                            if true {
                                                v93 := GenS8{F0: v48, F1: fi49}
                                                _ = v93
                                            }
                                            v48 = (v48 | -2147483648)
                                            fmt.Println(v90, fi49)
                                        } else {
                                            v3 += 329
                                            v94 := GenS8{F0: (^(v48)), F1: -1}
                                            _ = v94
                                        }
                                    }
                                case 0:
                                    v72 = 4
                                }
                                switch (^(fi71)) {
                                case 128:
                                    switch 9223372036854775807 {
                                    case -2147483648:
                                        v95 := GenS10{F0: (true && ((971 < 1) || (817 < 65535))), F1: (v72 % ((v72) | 1))}
                                        _ = v95
                                    case -466:
                                        var v96 bool = true
                                        _ = v96
                                    case 127:
                                        v97 := GenS7{F0: 13}
                                        _ = v97
                                    }
                                case -1:
                                    if ((!true) || (true || true)) {
                                        switch 128 {
                                        case 32767:
                                            v98 := GenS7{F0: int8(v45)}
                                            _ = v98
                                        }
                                        switch (("foo bar" + "foo bar") + "\n\t") {
                                        case "hello":
                                            var v99 int8 = 0
                                            _ = v99
                                        case "a":
                                            v48 -= (v48 % ((2147483647) | 1))
                                        }
                                        var v100 int32 = v48
                                        _ = v100
                                    } else {
                                        var v101 float64 = 3.14159
                                        _ = v101
                                    }
                                case 32767:
                                    if true {
                                        if (v72 <= (v72 | (v72 | v72))) {
                                            if (false && false) {
                                                v48 = -1
                                                if true {
                                                    var v102 int = -22
                                                    var v103 uint16 = 0
                                                    v104 := v5[v45]
                                                    switch ("hello" + "foo bar") {
                                                    case "hello":
                                                        v48 += (v48 / ((1) | 1))
                                                    default:
                                                        v102 += fi71
                                                    }
                                                    v105 := GenS6{F0: (v72 != 65535), F1: (v103 | (v103 >> ((v72) & 63)))}
                                                    if false {
                                                        fmt.Println(fi49, fi49, v48)
                                                        var v106 bool = true
                                                        _ = v106
                                                    }
                                                    _ = v102
                                                    _ = v103
                                                    _ = v104
                                                    _ = v105
                                                } else {
                                                    v107 := GenS3{F0: (v90 * (v3 >> ((v72) & 63))), F1: ("foo bar" + "foo bar")}
                                                    v90 += 1
                                                    _ = v107
                                                }
                                            }
                                        }
                                        var v108 uint = (v72 << (((^(v72))) & 63))
                                        var v109 int8 = 42
                                        var v110 uint64 = 4294967295
                                        v3 = v90
                                        _ = v108
                                        _ = v109
                                        _ = v110
                                    }
                                }
                                fmt.Println(fi49, v45, fi71)
                                if true {
                                    v45 -= v45
                                    v1 = append(v1, "a")
                                    var v111 uint = v72
                                    v112 := GenS7{F0: 42}
                                    var v113 float32 = (0.0 - 0.1)
                                    _ = v111
                                    _ = v112
                                    _ = v113
                                }
                                var v114 string = "a"
                                _ = v114
                                _ = v90
                            }
                            switch "a" {
                            case "":
                                v5[v45] = -2147483648
                            case "\n\t":
                                v115 := GenS11{F0: -1, F1: 952}
                                fmt.Println(v3, v45, v3)
                                _ = v115
                            case "a":
                                var v116 bool = (1 < 742)
                                _ = v116
                            }
                            if len(v2) > 0 {
                                v2[0] = ("a" + ("a" + "foo bar"))
                            }
                        default:
                            v72 -= 65535
                        }
                    default:
                        switch fi49 {
                        case 756:
                            var v117 bool = (-2.25 < ((1.5 * (((float64(float32(1.5)) * ((3.14159 + float64(1e10)) / (1.5 * -2.25))) + 1.5) / (0.1 / -2.25))) - (3.14159 + (1.5 / 1.5))))
                            _ = v117
                        }
                    }
                    v118 := GenS12{F0: (v48 / (((^(v48))) | 1)), F1: 127, F2: fi49}
                    var v119 bool = (0.1 >= (3.14159 * 1.5))
                    v120 := GenS9{F0: v72, F1: (v48 + -2147483648), F2: v3}
                    _ = v118
                    _ = v119
                    _ = v120
                }
                v121 := GenS9{F0: 42, F1: (v48 - v48), F2: (v3 | (v3 + v3))}
                _ = fi71
                _ = v121
                _ = v72
            }
            v3 = v3
        }
        for fi122 := 0; fi122 < 19; fi122++ {
            v123 := GenS8{F0: -1, F1: (^(fi122))}
            var v124 uint16 = 65535
            v125 := GenS3{F0: (v3 % (((v3 / ((int64(v48)) | 1))) | 1)), F1: "foo bar"}
            _ = fi122
            _ = v123
            _ = v124
            _ = v125
        }
        _ = fi49
    }
    fmt.Println(v45, v3)
    var v126 float64 = ((((1e10 / 1.0) / 1.5) / 1.5) + (0.0 * -2.25))
    var v127 uint64 = 608
    var v128 int64 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v126
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
}
