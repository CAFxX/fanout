package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 int16; F1 bool; F2 string }
type GenS13 struct { F0 uint; F1 string }
type GenS14 struct { F0 int64; F1 int16; F2 uint }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int16; F1 uint8 }
type GenS17 struct { F0 uint; F1 string; F2 uint8 }
type GenS18 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS19 struct { F0 int32; F1 string; F2 string }
type GenS20 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS21 struct { F0 int8 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 14; fi5++ {
        switch ("a" + (("" + ("" + "a")) + "foo bar")) {
        case "foo bar":
            v6 := GenS1{F0: 0}
            _ = v6
        }
        fmt.Println(fi5)
        var v7 int16 = int16(fi5)
        _ = fi5
        _ = v7
    }
    var v8 bool = false
    v9 := GenS2{F0: -268}
    v8 = (v8 && (-752 == -2147483648))
    for fi10 := 0; fi10 < 13; fi10++ {
        if (v8 || (127 < 1)) {
            var v11 int16 = -784
            v12 := GenS3{F0: 0}
            if v8 {
                switch fi10 {
                case 187:
                    if v8 {
                        switch "foo bar" {
                        case "hello":
                            if (-128 >= 0) {
                                switch (-(fi10)) {
                                case 619:
                                    for fi13 := 0; fi13 < 20; fi13++ {
                                        v8 = (!(1 <= 27))
                                        var v14 uint64 = 18446744073709551615
                                        _ = fi13
                                        _ = v14
                                    }
                                case 1000:
                                    var v15 string = "foo bar"
                                    _ = v15
                                case -2147483648:
                                    if v8 {
                                        var v16 uint32 = 759
                                        v11 += v11
                                        _ = v16
                                    } else {
                                        for fi17 := 0; fi17 < 4; fi17++ {
                                            v11 = (v11 + v11)
                                            var v18 int32 = 406
                                            fmt.Println(v11, v8)
                                            _ = fi17
                                            _ = v18
                                        }
                                    }
                                }
                                v19 := GenS4{F0: 255, F1: 2, F2: 1}
                                _ = v19
                            } else {
                                var v20 uint = 2
                                if (65535 > (v20 * v20)) {
                                    for fi21 := 0; fi21 < 9; fi21++ {
                                        var v22 int64 = 1
                                        _ = fi21
                                        _ = v22
                                    }
                                    v23 := GenS1{F0: (v11 * 42)}
                                    _ = v23
                                } else {
                                    v24 := GenS5{F0: (!((!true) || v8))}
                                    v20 -= v20
                                    switch 32768 {
                                    case -2147483648:
                                        if (-800 > 0) {
                                            for fi25 := 0; fi25 < 6; fi25++ {
                                                v26 := []int{32768, (fi25 << ((18446744073709551615) & 63)), (fi10 % ((fi25) | 1)), 2147483647}
                                                fmt.Println(v20)
                                                if (1 != 42) {
                                                    if true {
                                                        v11 += v11
                                                        if len(v26) > 0 {
                                                            v26[0] = (fi10 & fi10)
                                                        }
                                                        var v27 uint16 = 163
                                                        _ = v27
                                                    }
                                                    var v28 uint8 = 185
                                                    var v29 string = "\n\t"
                                                    var v30 int64 = -885
                                                    if false {
                                                        switch v29 {
                                                        case "a":
                                                            if v8 {
                                                                var v31 int32 = 0
                                                                v32 := GenS6{F0: 0}
                                                                if (((v30 - 285) < -712) || (v28 <= v28)) {
                                                                    v33 := GenS6{F0: 42}
                                                                    v26 = append(v26, (^(fi25)))
                                                                    _ = v33
                                                                }
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                        default:
                                                            fmt.Println(fi25, v8, v20)
                                                        }
                                                        if len(v26) >= 2 {
                                                            v26 = v26[1:len(v26)]
                                                        }
                                                        var v34 int = 32767
                                                        var v35 int64 = (v30 * (v30 & (v30 / (((v30 >> (((v20 | (v20 / (((v20 >> ((v20) & 63))) | 1)))) & 63))) | 1))))
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    var v36 uint = v20
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                    _ = v36
                                                }
                                                _ = fi25
                                                _ = v26
                                            }
                                        } else {
                                            var v37 int16 = v11
                                            _ = v37
                                        }
                                    case 1:
                                        if v8 {
                                            for fi38 := 0; fi38 < 15; fi38++ {
                                                if (-2.25 != 0.1) {
                                                    switch (fi10 + (fi38 & (fi10 & (fi10 >> (((v20 * (^(v20)))) & 63))))) {
                                                    case -9223372036854775808:
                                                        switch "" {
                                                        case "hello":
                                                            fmt.Println(fi10)
                                                        case "foo bar":
                                                            if (float64((0.0 / 1.0)) > 3.14159) {
                                                                var v39 uint16 = 678
                                                                fmt.Println(v20, v39)
                                                                _ = v39
                                                            } else {
                                                                v40 := GenS2{F0: -702}
                                                                var v41 int8 = 41
                                                                if v8 {
                                                                    fmt.Println(v8)
                                                                }
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                        case "a":
                                                            var v42 int = (fi38 ^ fi38)
                                                            _ = v42
                                                        default:
                                                            fmt.Println(fi38, fi10)
                                                        }
                                                    case 42:
                                                        v20 = (v20 + (^(v20)))
                                                        var v43 uint8 = 233
                                                        _ = v43
                                                    case -554:
                                                        switch 945 {
                                                        case 9223372036854775807:
                                                            var v44 uint16 = 255
                                                            _ = v44
                                                        case -2147483648:
                                                            switch (fi38 ^ (fi38 - fi38)) {
                                                            case 2147483647:
                                                                v45 := []int{}
                                                                _ = v45
                                                            case 9223372036854775807:
                                                                var v46 int = fi38
                                                                _ = v46
                                                            }
                                                        }
                                                    }
                                                    if (true && ("" < "\n\t")) {
                                                        v47 := GenS1{F0: 315}
                                                        v48 := []int{fi38, fi38}
                                                        v49 := GenS7{F0: 29}
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        v20 -= (v20 * (^(v20)))
                                                        fmt.Println(v11, v11)
                                                        switch "hello" {
                                                        case "a":
                                                            switch ("foo bar" + "hello") {
                                                            case "a":
                                                                v50 := []int{4294967295, 0}
                                                                _ = v50
                                                            case "foo bar":
                                                                switch (fi38 * fi38) {
                                                                case 32768:
                                                                    fmt.Println(fi10, v8, v8)
                                                                case -909:
                                                                    var v51 uint16 = 0
                                                                    _ = v51
                                                                case 546:
                                                                    switch ("foo bar" + "foo bar") {
                                                                    case "\n\t":
                                                                        v52 := []int{4294967295, fi10}
                                                                        _ = v52
                                                                    case "a":
                                                                        if (18446744073709551615 != 498) {
                                                                            v20 = 255
                                                                            v11 = v11
                                                                            var v53 float32 = (float32(0.1) - 1.5)
                                                                            _ = v53
                                                                        } else {
                                                                            var v54 string = "\n\t"
                                                                            if v8 {
                                                                                v55 := []int{fi38, -1, (fi10 / ((fi38) | 1)), (fi10 ^ fi38)}
                                                                                var v56 uint8 = 255
                                                                                var v57 uint64 = 10
                                                                                _ = v55
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                if false {
                                                                                    if true {
                                                                                        var v58 string = "foo bar"
                                                                                        v59 := GenS7{F0: 255}
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    } else {
                                                                                        v60 := []int{(^(fi38))}
                                                                                        var v61 uint32 = 2147483647
                                                                                        if ((-2.25 * float64(1.5)) < float64(0.1)) {
                                                                                            v8 = false
                                                                                            if v8 {
                                                                                                if (v8 || false) {
                                                                                                    v11 = v11
                                                                                                } else {
                                                                                                    switch (fi38 + (fi38 & (fi10 & fi10))) {
                                                                                                    case -783:
                                                                                                        fmt.Println(v8, v54)
                                                                                                    }
                                                                                                    v11 = -822
                                                                                                    switch ((("foo bar" + v54) + v54) + ("foo bar" + v54)) {
                                                                                                    case "\n\t":
                                                                                                        v62 := GenS5{F0: (127 != -1)}
                                                                                                        _ = v62
                                                                                                    case "hello":
                                                                                                        if false {
                                                                                                            if (!v8) {
                                                                                                                if len(v60) >= 2 {
                                                                                                                    v60 = v60[1:len(v60)]
                                                                                                                }
                                                                                                            }
                                                                                                            v20 = (v20 % ((65535) | 1))
                                                                                                            var v63 uint32 = (v61 - v61)
                                                                                                            _ = v63
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                if (2147483647 <= (fi10 * fi38)) {
                                                                                                    if len(v60) > 0 {
                                                                                                        v60[0] = fi38
                                                                                                    }
                                                                                                }
                                                                                                var v64 uint32 = (^(v61))
                                                                                                _ = v64
                                                                                            } else {
                                                                                                switch (fi38 % (((fi38 + (fi10 / ((-375) | 1)))) | 1)) {
                                                                                                case -1:
                                                                                                    v65 := GenS8{F0: 71}
                                                                                                    _ = v65
                                                                                                case 1:
                                                                                                    if false {
                                                                                                        if (false || (0 != fi38)) {
                                                                                                            if false {
                                                                                                                if len(v60) >= 2 {
                                                                                                                    v60 = v60[1:len(v60)]
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v66 bool = (0.0 <= (((0.1 - 0.0) + 1e10) - 1e10))
                                                                                                                _ = v66
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v67 float32 = 3.14159
                                                                                                            _ = v67
                                                                                                        }
                                                                                                        v8 = v8
                                                                                                    }
                                                                                                default:
                                                                                                    if len(v60) >= 2 {
                                                                                                        v60 = v60[1:len(v60)]
                                                                                                    }
                                                                                                    if (1 < 138) {
                                                                                                        v61 -= v61
                                                                                                    } else {
                                                                                                        var v68 bool = ((float32(((0.1 * ((1.5 * 0.0) / float64(1.5))) - float64(0.1))) + float32(0.1)) == 1e10)
                                                                                                        if ((42 < 542) && false) {
                                                                                                            v8 = (true && v68)
                                                                                                            var v69 int64 = -910
                                                                                                            _ = len(v60)
                                                                                                            _ = v69
                                                                                                        } else {
                                                                                                            v68 = false
                                                                                                        }
                                                                                                        _ = v68
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v60
                                                                                        _ = v61
                                                                                    }
                                                                                    v70 := []int{fi38, fi10, 42, fi10}
                                                                                    var v71 uint = v20
                                                                                    _ = v70
                                                                                    _ = v71
                                                                                }
                                                                                var v72 uint64 = 42
                                                                                _ = v72
                                                                            }
                                                                            switch (v54 + "\n\t") {
                                                                            case "foo bar":
                                                                                var v73 int64 = 327
                                                                                _ = v73
                                                                            case "":
                                                                                v11 -= (v11 * v11)
                                                                            }
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                    if (0 == 255) {
                                                                        v11 = (^(v11))
                                                                    } else {
                                                                        var v74 int = 935
                                                                        _ = v74
                                                                    }
                                                                    switch (fi38 % ((42) | 1)) {
                                                                    case 0:
                                                                        var v75 int8 = 0
                                                                        _ = v75
                                                                    case -2147483648:
                                                                        v76 := []int{}
                                                                        _ = v76
                                                                    default:
                                                                        fmt.Println(v8, v20)
                                                                    }
                                                                default:
                                                                    v20 = v20
                                                                }
                                                            }
                                                        }
                                                        v8 = v8
                                                        var v77 uint16 = 72
                                                        if true {
                                                            var v78 int = (fi10 << ((42) & 63))
                                                            _ = v78
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            var v79 int64 = -1
                                                            _ = v79
                                                        case "\n\t":
                                                            v80 := []int{}
                                                            _ = v80
                                                        default:
                                                            v8 = (30 != v11)
                                                        }
                                                        v8 = v8
                                                        var v81 int8 = 1
                                                        _ = v77
                                                        _ = v81
                                                    }
                                                    var v82 float64 = float64(0.0)
                                                    _ = v82
                                                }
                                                if true {
                                                    var v83 uint = 255
                                                    v84 := GenS7{F0: 0}
                                                    fmt.Println(v11)
                                                    var v85 uint8 = 171
                                                    fmt.Println(v11, v20)
                                                    _ = v83
                                                    _ = v84
                                                    _ = v85
                                                }
                                                _ = fi38
                                            }
                                            var v86 uint8 = 128
                                            _ = v86
                                        } else {
                                            v87 := []int{(fi10 + fi10), (fi10 * fi10), (fi10 ^ (fi10 % ((fi10) | 1))), (fi10 % (((fi10 * (fi10 >> ((v20) & 63)))) | 1))}
                                            v88 := GenS9{F0: (("a" + "foo bar") + ("\n\t" + "a")), F1: 84}
                                            _ = v87
                                            _ = v88
                                        }
                                    }
                                    _ = v24
                                }
                                _ = v20
                            }
                        default:
                            v89 := GenS4{F0: 18, F1: 990, F2: 42}
                            _ = v89
                        }
                        var v90 int16 = 0
                        _ = v90
                    } else {
                        for fi91 := 0; fi91 < 1; fi91++ {
                            if v8 {
                                var v92 bool = true
                                _ = v92
                            }
                            switch (fi91 - fi10) {
                            case 0:
                                fmt.Println(fi91, fi10)
                            default:
                                switch fi10 {
                                case 4294967295:
                                    var v93 int32 = 42
                                    _ = v93
                                }
                            }
                            _ = fi91
                        }
                        switch "" {
                        case "hello":
                            v94 := []int{-231, (fi10 % (((^(fi10))) | 1)), fi10, 42}
                            _ = v94
                        default:
                            switch -9223372036854775808 {
                            case 1:
                                v95 := []int{fi10, fi10, fi10}
                                _ = v95
                            case 393:
                                switch (("\n\t" + ("a" + "a")) + "a") {
                                case "foo bar":
                                    for fi96 := 0; fi96 < 10; fi96++ {
                                        v97 := GenS10{F0: v11}
                                        var v98 uint64 = 229
                                        switch (fi96 + fi10) {
                                        case -949:
                                            switch fi10 {
                                            case 127:
                                                v99 := []int{(fi10 | int(v11)), (fi96 ^ (fi96 - -2147483648)), fi10, (fi10 << ((295) & 63))}
                                                _ = v99
                                            default:
                                                v100 := GenS11{F0: 22, F1: (400 > 904)}
                                                _ = v100
                                            }
                                        default:
                                            if false {
                                                switch "a" {
                                                case "hello":
                                                    if false {
                                                        v101 := GenS11{F0: -42, F1: (1 >= 42)}
                                                        if (!(!true)) {
                                                            v102 := []int{fi10, 2147483647}
                                                            _ = v102
                                                        } else {
                                                            v103 := GenS11{F0: 100, F1: (255 > 2)}
                                                            if ((0 >= -819) || (!v8)) {
                                                                var v104 float64 = 0.1
                                                                if v8 {
                                                                    v8 = true
                                                                    v8 = ((!false) && (!v8))
                                                                }
                                                                _ = v104
                                                            }
                                                            var v105 int8 = 0
                                                            _ = v103
                                                            _ = v105
                                                        }
                                                        _ = v101
                                                    } else {
                                                        var v106 int32 = 42
                                                        _ = v106
                                                    }
                                                case "":
                                                    v8 = v8
                                                }
                                                v107 := []int{-790}
                                                _ = v107
                                            } else {
                                                if true {
                                                    v108 := []int{(fi10 * -668), fi96, (fi10 / ((fi96) | 1)), -2147483648}
                                                    if (((-2.25 / -2.25) * 0.0) < ((0.1 * 0.1) / ((((1e10 / 1e10) * ((0.0 / 0.1) + -2.25)) * 3.14159) + 0.1))) {
                                                        v109 := GenS12{F0: v11, F1: (43 > -121), F2: ""}
                                                        fmt.Println(fi10)
                                                        _ = v109
                                                    }
                                                    fmt.Println(fi96, fi96, v11)
                                                    _ = v108
                                                } else {
                                                    var v110 float32 = 0.1
                                                    var v111 int16 = (v11 % ((v11) | 1))
                                                    _ = v110
                                                    _ = v111
                                                }
                                                var v112 uint64 = (v98 % (((^(v98))) | 1))
                                                v113 := []int{fi10, (fi96 * (fi96 % ((fi10) | 1))), (fi10 + (fi10 / (((fi96 / (((^(fi10))) | 1))) | 1))), (fi96 % ((-900) | 1))}
                                                _ = v112
                                                _ = v113
                                            }
                                        }
                                        v8 = false
                                        v114 := []int{fi10}
                                        v11 = (v11 / ((v11) | 1))
                                        if (v8 && v8) {
                                            v8 = false
                                            var v115 string = "a"
                                            var v116 int32 = -541
                                            _ = v115
                                            _ = v116
                                        }
                                        _ = fi96
                                        _ = v114
                                        _ = v97
                                        _ = v98
                                    }
                                case "a":
                                    if (1 <= 0) {
                                        v117 := GenS13{F0: 134, F1: ("hello" + "a")}
                                        _ = v117
                                    } else {
                                        var v118 uint16 = 400
                                        _ = v118
                                    }
                                }
                            }
                        }
                    }
                case 2147483647:
                    switch ("hello" + ("a" + "a")) {
                    case "a":
                        var v119 float64 = float64(0.0)
                        _ = v119
                    case "foo bar":
                        var v120 int = fi10
                        _ = v120
                    case "hello":
                        for fi121 := 0; fi121 < 10; fi121++ {
                            if v8 {
                                var v122 uint = 42
                                switch "\n\t" {
                                case "foo bar":
                                    var v123 uint8 = 13
                                    _ = v123
                                case "a":
                                    if (4294967295 >= 0) {
                                        v8 = v8
                                        v124 := GenS2{F0: -1}
                                        fmt.Println(fi10, fi121)
                                        _ = v124
                                    }
                                case "\n\t":
                                    var v125 uint32 = 480
                                    _ = v125
                                default:
                                    if (3.14159 < 0.1) {
                                        fmt.Println(v122, v8)
                                        switch (^(fi121)) {
                                        case -584:
                                            switch ("hello" + "hello") {
                                            case "foo bar":
                                                fmt.Println(v11, v122, fi10)
                                            case "":
                                                fmt.Println(fi10, fi121, v8)
                                                v122 = (v122 ^ (v122 | (v122 / ((v122) | 1))))
                                            }
                                        default:
                                            var v126 float32 = 0.0
                                            _ = v126
                                        }
                                    } else {
                                        var v127 int16 = -32768
                                        _ = v127
                                    }
                                }
                                _ = v122
                            }
                            var v128 int16 = 320
                            _ = fi121
                            _ = v128
                        }
                    default:
                        v11 -= v11
                    }
                case -221:
                    if ((1.5 - ((1e10 + (0.0 + 3.14159)) - ((-2.25 - -2.25) / 3.14159))) <= 0.1) {
                        var v129 uint16 = 42
                        if false {
                            v130 := GenS12{F0: (^(v11)), F1: v8, F2: "\n\t"}
                            var v131 string = ("" + "hello")
                            _ = v130
                            _ = v131
                        } else {
                            var v132 uint16 = 42
                            v8 = false
                            _ = v132
                        }
                        _ = v129
                    } else {
                        switch ("\n\t" + "hello") {
                        case "hello":
                            var v133 float64 = ((-2.25 * float64(((1.5 / 3.14159) + 3.14159))) - (float64(3.14159) - (1e10 - 3.14159)))
                            _ = v133
                        default:
                            if (!v8) {
                                if (!(42 < 1)) {
                                    var v134 string = ""
                                    if v8 {
                                        v11 -= (v11 ^ 0)
                                    }
                                    _ = v134
                                }
                                v135 := GenS14{F0: -275, F1: 1, F2: 747}
                                _ = v135
                            } else {
                                for fi136 := 0; fi136 < 14; fi136++ {
                                    v137 := GenS15{F0: 255}
                                    if v8 {
                                        var v138 uint32 = 489
                                        _ = v138
                                    } else {
                                        var v139 float64 = -2.25
                                        fmt.Println(fi136, fi10)
                                        var v140 uint = 1
                                        _ = v139
                                        _ = v140
                                    }
                                    v141 := []int{(fi10 | fi136), 4294967295}
                                    _ = fi136
                                    _ = v137
                                    _ = v141
                                }
                                var v142 float32 = 3.14159
                                var v143 float32 = ((1.5 - v142) - v142)
                                _ = v142
                                _ = v143
                            }
                        }
                        if (!((fi10 ^ 95) < (fi10 & fi10))) {
                            switch "foo bar" {
                            case "hello":
                                v8 = false
                            case "a":
                                fmt.Println(v8, v11)
                            default:
                                for fi144 := 0; fi144 < 7; fi144++ {
                                    v8 = (9223372036854775807 <= -1)
                                    _ = fi144
                                }
                            }
                            if v8 {
                                v145 := []int{fi10, (fi10 % ((-244) | 1))}
                                if len(v145) >= 2 {
                                    v145 = v145[1:len(v145)]
                                }
                                _ = v145
                            } else {
                                var v146 uint = 812
                                if (42 < 0) {
                                    var v147 uint64 = 879
                                    if ((!(1000 >= (-(v11)))) || ((!(("hello" + "") >= "foo bar")) || (v11 != v11))) {
                                        fmt.Println(v11)
                                        v148 := []int{fi10, fi10}
                                        v149 := GenS16{F0: (v11 - (v11 % (((v11 * v11)) | 1))), F1: 2}
                                        _ = v148
                                        _ = v149
                                    }
                                    v147 -= v147
                                    _ = v147
                                } else {
                                    v8 = false
                                    for fi150 := 0; fi150 < 16; fi150++ {
                                        v151 := []int{4294967295, (fi10 | fi10)}
                                        var v152 uint32 = 2
                                        v151 = append(v151, fi10)
                                        v146 = (^(v146))
                                        fmt.Println(fi10)
                                        if len(v151) > 0 {
                                            v151[0] = -218
                                        }
                                        var v153 uint64 = 1
                                        _ = fi150
                                        _ = v151
                                        _ = v152
                                        _ = v153
                                    }
                                    v8 = ((false && (3.14159 > ((1.5 - -2.25) - 1.5))) || false)
                                }
                                switch fi10 {
                                case -2147483648:
                                    v8 = v8
                                }
                                _ = v146
                            }
                        }
                    }
                default:
                    for fi154 := 0; fi154 < 15; fi154++ {
                        v155 := []int{}
                        v156 := GenS17{F0: 4294967295, F1: ("hello" + "a"), F2: 0}
                        var v157 bool = v8
                        if v8 {
                            v155 = append(v155, -941)
                            v8 = ((0.1 / 1.0) <= (1.5 * 0.0))
                            var v158 bool = v157
                            _ = v158
                        } else {
                            var v159 int = 89
                            _ = v159
                        }
                        var v160 float32 = (0.0 + float32(3.14159))
                        _ = fi154
                        _ = v155
                        _ = v156
                        _ = v157
                        _ = v160
                    }
                }
                if (!v8) {
                    v11 = (v11 >> ((764) & 63))
                    switch fi10 {
                    case -229:
                        v8 = true
                    }
                } else {
                    v161 := GenS18{F0: 42, F1: 2147483647, F2: 32767}
                    fmt.Println(fi10, v11, v8)
                    _ = v161
                }
                switch "\n\t" {
                case "":
                    v162 := []int{941, (fi10 >> ((0) & 63)), fi10, fi10}
                    _ = v162
                case "\n\t":
                    v163 := []int{}
                    _ = v163
                default:
                    for fi164 := 0; fi164 < 18; fi164++ {
                        var v165 bool = (0 <= 128)
                        _ = fi164
                        _ = v165
                    }
                }
            } else {
                if (2147483647 == -901) {
                    v166 := []int{}
                    fmt.Println(v11, v8, v8)
                    _ = v166
                }
                if (fi10 < fi10) {
                    v8 = v8
                    if true {
                        var v167 int8 = -128
                        for fi168 := 0; fi168 < 1; fi168++ {
                            v167 += v167
                            v169 := []int{int(v167), (fi168 & (-(fi168)))}
                            v167 += (v167 + (v167 + v167))
                            _ = fi168
                            _ = v169
                        }
                        v11 = (v11 % ((v11) | 1))
                        _ = v167
                    }
                } else {
                    var v170 int16 = (v11 ^ -32768)
                    fmt.Println(v170, v8)
                    v11 += v11
                    _ = v170
                }
                var v171 uint = 2
                _ = v171
            }
            _ = v11
            _ = v12
        } else {
            var v172 float64 = -2.25
            _ = v172
        }
        for fi173 := 0; fi173 < 9; fi173++ {
            fmt.Println(fi10, fi10)
            fmt.Println(v8, fi173, fi173)
            _ = fi173
        }
        var v174 uint8 = 13
        _ = fi10
        _ = v174
    }
    _ = v8
    _ = v9
    return 34
}
func genFn1(p0 int16, p1 bool, p2 int8) int64 {
    v175 := []int{805, 0, 32767}
    _ = len(v175)
    _ = p0
    _ = p1
    _ = p2
    _ = v175
    return genFn0()
}

func main() {
    v1 := []int{-1, 9223372036854775807, 4294967295, 2147483647}
    var v2 bool = true
    v3 := make(map[string]bool)
    var v4 int = 267
    v4 = v4
    v4 = 997
    if true {
        var v176 bool = (1 <= -9223372036854775808)
        if (!v2) {
            v177 := GenS19{F0: 2147483647, F1: "\n\t", F2: "a"}
            if v176 {
                if (0 != 1) {
                    v178 := GenS20{F0: (v176 && false), F1: 1, F2: 42}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v178
                } else {
                    for fi179 := 0; fi179 < 8; fi179++ {
                        var v180 int16 = 162
                        v181 := GenS21{F0: -100}
                        _ = fi179
                        _ = v180
                        _ = v181
                    }
                    v2 = (true || (65535 <= 42))
                }
                v176 = (("a" == "a") || (1e10 <= -2.25))
                v4 += -2147483648
            } else {
                v182 := v3["hello"]
                v2 = true
                v4 = 128
                _ = v182
            }
            v176 = v2
            _ = v177
        }
        for fi183 := 0; fi183 < 3; fi183++ {
            var v184 float64 = 3.14159
            _ = fi183
            _ = v184
        }
        _ = v176
    }
    var v185 int16 = 1000
    v2 = (4294967295 < 61)
    switch "a" {
    case "foo bar":
        var v186 bool = (245 >= 132)
        _ = v186
    case "":
        _ = len(v1)
    case "\n\t":
        v185 -= (v185 / (((^(v185))) | 1))
    }
    v187 := GenS18{F0: 127, F1: 657, F2: v185}
    delete(v3, (("foo bar" + ("" + "")) + "a"))
    v4 -= v4
    if (!(v2 && ((v185 & v185) <= v185))) {
        var v188 int8 = 80
        _ = len(v3)
        for fi189 := 0; fi189 < 2; fi189++ {
            if (v2 && (!((-(v4)) < (v4 * (fi189 + 32768))))) {
                if ((v188 >> ((2) & 63)) > int8(v185)) {
                    v190 := GenS14{F0: 42, F1: v185, F2: 0}
                    var v191 int64 = 9223372036854775807
                    _ = v190
                    _ = v191
                }
                for fi192 := 0; fi192 < 10; fi192++ {
                    var v193 uint16 = 0
                    var v194 string = "foo bar"
                    v185 -= v185
                    _ = fi192
                    _ = v193
                    _ = v194
                }
            }
            _ = fi189
        }
        _ = v188
    }
    var v195 uint16 = 255
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v185:", v185)
    fmt.Println("v195:", v195)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v185
    _ = v187
    _ = v195
    _ = v2
    _ = v3
    _ = v4
}
