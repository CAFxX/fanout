package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS8 struct { F0 bool; F1 uint8; F2 int64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32; F1 int; F2 bool }
type GenS11 struct { F0 string; F1 uint32; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 61
    var v3 uint8 = 42
    var v4 uint16 = 0
    var v5 int16 = 477
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: -1, F1: 2147483647, F2: 1}
    switch 127 {
    case 9223372036854775807:
        v4 -= v4
    default:
        switch "\n\t" {
        case "a":
            switch 829 {
            case 4294967295:
                if (1 == 910) {
                    for fi8 := 0; fi8 < 10; fi8++ {
                        var v9 uint16 = 2
                        _ = fi8
                        _ = v9
                    }
                    v10 := []int{-49}
                    v11 := GenS2{F0: -1, F1: "hello", F2: 2}
                    _ = v10
                    _ = v11
                }
            }
        }
    }
    if (!(!(false || true))) {
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := []int{fi12, fi13}
                if (0.1 < (((1.5 - 0.0) * ((float64(1.5) + 1e10) + 1e10)) - 0.0)) {
                    if true {
                        _ = len(v14)
                        v15 := v1[-2147483648]
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v15
                    }
                    if false {
                        switch 32768 {
                        case -867:
                            var v16 uint16 = v4
                            if true {
                                if false {
                                    var v17 int32 = 1
                                    var v18 float32 = ((float32((3.14159 - 1e10)) / (1.5 / (0.1 * 0.1))) + float32(-2.25))
                                    if true {
                                        var v19 uint32 = 2147483647
                                        _ = v19
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                fmt.Println(fi13)
                                if ((-1 > 0) || false) {
                                    v20 := GenS2{F0: 520, F1: "hello", F2: 255}
                                    v5 += v5
                                    var v21 uint = 65535
                                    switch fi13 {
                                    case -1:
                                        if (false || (false || true)) {
                                            v16 -= v16
                                        } else {
                                            fmt.Println(v21)
                                        }
                                    }
                                    if (48 == 58) {
                                        _ = len(v14)
                                        v2 = (v2 >> ((v21) & 63))
                                        switch fi13 {
                                        case 795:
                                            fmt.Println(fi13, v16, fi13)
                                        case 2147483647:
                                            v2 += 0
                                        case 540:
                                            if (0 > -1) {
                                                switch ("a" + "hello") {
                                                case "hello":
                                                    var v22 string = "a"
                                                    fmt.Println(v16, v4, v21)
                                                    var v23 uint64 = 0
                                                    v24 := GenS3{F0: 4294967295, F1: ("" + (v22 + (("foo bar" + v22) + (("a" + (v22 + (("hello" + v22) + v22))) + v22))))}
                                                    _ = v22
                                                    _ = v23
                                                    _ = v24
                                                case "a":
                                                    fmt.Println(v16, v4)
                                                case "":
                                                    v25 := GenS4{F0: 1, F1: 373}
                                                    _ = v25
                                                }
                                            } else {
                                                if (57 <= v2) {
                                                    v16 -= 2
                                                    var v26 int64 = -9223372036854775808
                                                    var v27 int8 = int8(v26)
                                                    fmt.Println(v16, v5)
                                                    _ = v26
                                                    _ = v27
                                                }
                                            }
                                        default:
                                            if len(v14) >= 2 {
                                                v14 = v14[1:len(v14)]
                                            }
                                            if true {
                                                fmt.Println(v21, v16, fi12)
                                                v1[fi13] = ((v5 | v5) <= int16(v2))
                                            } else {
                                                if (fi13 != 128) {
                                                    var v28 int64 = -932
                                                    v29 := GenS2{F0: v5, F1: ("hello" + ("" + "a")), F2: 18446744073709551615}
                                                    v30 := GenS5{F0: 127}
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                }
                                                var v31 int16 = (v5 / ((-225) | 1))
                                                var v32 int = (fi13 ^ fi13)
                                                fmt.Println(v31, fi12, v32)
                                                switch ("hello" + ("a" + (("hello" + "\n\t") + "foo bar"))) {
                                                case "\n\t":
                                                    fmt.Println(v2, v3, v31)
                                                    v3 = v3
                                                case "a":
                                                    fmt.Println(fi13, v31)
                                                case "foo bar":
                                                    v21 += v21
                                                }
                                                _ = v31
                                                _ = v32
                                            }
                                            v6[("\n\t" + "foo bar")] = 65535
                                            v33 := GenS5{F0: 42}
                                            var v34 int = int(v3)
                                            if true {
                                                fmt.Println(v21)
                                                fmt.Println(fi12, fi12)
                                            } else {
                                                v35 := GenS6{F0: (^(v34))}
                                                if true {
                                                    fmt.Println(v21, v2)
                                                    fmt.Println(v16)
                                                    var v36 int32 = 2147483647
                                                    v3 = (^(v2))
                                                    v14 = append(v14, fi13)
                                                    switch (("\n\t" + (("a" + "foo bar") + ("foo bar" + ""))) + ("hello" + "foo bar")) {
                                                    case "":
                                                        v36 = (v36 & (v36 >> (((^(v21))) & 63)))
                                                    case "hello":
                                                        if false {
                                                            if true {
                                                                var v37 float32 = 3.14159
                                                                v38 := GenS6{F0: (fi13 | (fi13 ^ fi13))}
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                            delete(v1, v34)
                                                        } else {
                                                            if (!(true || true)) {
                                                                v16 = v16
                                                            } else {
                                                                if false {
                                                                    v21 -= v21
                                                                    var v39 uint8 = v2
                                                                    var v40 bool = false
                                                                    var v41 int64 = 9223372036854775807
                                                                    v16 -= 2
                                                                    _ = v39
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                                switch "" {
                                                                case "":
                                                                    if ((true && false) && true) {
                                                                        var v42 string = ("hello" + "\n\t")
                                                                        _ = v42
                                                                    } else {
                                                                        var v43 uint64 = 0
                                                                        if len(v14) > 0 {
                                                                            v14[0] = (fi12 & v34)
                                                                        }
                                                                        var v44 uint32 = uint32(v43)
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                    v45 := GenS4{F0: 9223372036854775807, F1: uint64(v21)}
                                                                    _ = v45
                                                                case "\n\t":
                                                                    v3 = (v2 - (^(v2)))
                                                                case "foo bar":
                                                                    if true {
                                                                        var v46 int = v34
                                                                        v47 := GenS3{F0: 65535, F1: "\n\t"}
                                                                        v48 := GenS6{F0: -952}
                                                                        v49 := GenS4{F0: 9223372036854775807, F1: 2}
                                                                        _ = v46
                                                                        _ = v47
                                                                        _ = v48
                                                                        _ = v49
                                                                    } else {
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            var v50 uint8 = v2
                                                                            _ = v50
                                                                        default:
                                                                            if (1e10 == 3.14159) {
                                                                                var v51 uint = 1
                                                                                delete(v6, "")
                                                                                _ = v51
                                                                            } else {
                                                                                v21 = v21
                                                                                var v52 string = "a"
                                                                                _ = v52
                                                                            }
                                                                            v53 := GenS4{F0: -2, F1: 612}
                                                                            fmt.Println(fi13, v3, fi12)
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            if ((207 <= -1) && false) {
                                                                fmt.Println(v4, fi13, v36)
                                                                v34 = (v34 & v34)
                                                            }
                                                            if len(v14) > 0 {
                                                                v14[0] = fi13
                                                            }
                                                        }
                                                    default:
                                                        if len(v14) >= 2 {
                                                            v14 = v14[1:len(v14)]
                                                        }
                                                    }
                                                    v4 -= (v4 - 251)
                                                    _ = v36
                                                } else {
                                                    switch ("a" + "hello") {
                                                    case "hello":
                                                        var v54 int16 = (v5 / ((32767) | 1))
                                                        var v55 int16 = (v54 / ((-32768) | 1))
                                                        var v56 int8 = -128
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    case "a":
                                                        if false {
                                                            if false {
                                                                v34 = (fi13 | (-(fi13)))
                                                                var v57 uint = (v21 ^ (v21 << ((v21) & 63)))
                                                                v58 := GenS7{F0: (v2 + v3), F1: 2147483647, F2: (^(v4))}
                                                                fmt.Println(fi13, v21)
                                                                if false {
                                                                    v3 = (v3 - (^(v3)))
                                                                    var v59 uint = v57
                                                                    _ = len(v14)
                                                                    _ = v59
                                                                }
                                                                _ = v57
                                                                _ = v58
                                                            }
                                                            _ = len(v14)
                                                            fmt.Println(v4, v16, fi13)
                                                        }
                                                    }
                                                }
                                                if (42 == (v5 ^ v5)) {
                                                    if (v16 > (^(v4))) {
                                                        var v60 float64 = 3.14159
                                                        _ = len(v1)
                                                        _ = v60
                                                    } else {
                                                        var v61 uint16 = v4
                                                        v62 := GenS6{F0: fi13}
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    var v63 int8 = 42
                                                    _ = v63
                                                } else {
                                                    var v64 uint = 18446744073709551615
                                                    _ = len(v14)
                                                    v21 -= 650
                                                    _ = v64
                                                }
                                                v4 += (v16 % (((v4 % (((^(v16))) | 1))) | 1))
                                                var v65 float64 = 3.14159
                                                _ = v35
                                                _ = v65
                                            }
                                            var v66 float32 = 3.14159
                                            var v67 uint32 = 422
                                            if len(v14) >= 2 {
                                                v14 = v14[1:len(v14)]
                                            }
                                            _ = v33
                                            _ = v34
                                            _ = v66
                                            _ = v67
                                        }
                                        if true {
                                            v5 = (v5 * v5)
                                            fmt.Println(v4, fi12)
                                        }
                                    } else {
                                        var v68 int32 = 2147483647
                                        _ = v68
                                    }
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    _ = v20
                                    _ = v21
                                } else {
                                    delete(v1, fi12)
                                    var v69 int8 = 21
                                    _ = v69
                                }
                            } else {
                                var v70 float32 = 1.5
                                if true {
                                    var v71 int32 = -990
                                    switch ("a" + (("foo bar" + "a") + "hello")) {
                                    case "\n\t":
                                        v71 -= -939
                                    case "foo bar":
                                        v72 := GenS8{F0: (!false), F1: 42, F2: 0}
                                        _ = v72
                                    case "hello":
                                        var v73 int8 = 100
                                        _ = v73
                                    default:
                                        switch ("\n\t" + "a") {
                                        case "hello":
                                            v71 = v71
                                        case "a":
                                            var v74 uint16 = uint16(v3)
                                            _ = v74
                                        default:
                                            var v75 int64 = -337
                                            _ = v75
                                        }
                                    }
                                    _ = len(v1)
                                    _ = v71
                                } else {
                                    v76 := GenS5{F0: 72}
                                    _ = v76
                                }
                                v77 := GenS4{F0: 9223372036854775807, F1: 951}
                                v2 = v2
                                _ = v70
                                _ = v77
                            }
                            switch fi12 {
                            case 42:
                                switch (((((("a" + "hello") + "hello") + "foo bar") + "hello") + ("a" + ("" + ((("foo bar" + "hello") + ("hello" + "a")) + "")))) + (("a" + ("hello" + "")) + "a")) {
                                case "\n\t":
                                    var v78 int16 = v5
                                    _ = v78
                                case "a":
                                    v79 := GenS9{F0: -30}
                                    _ = v79
                                case "":
                                    if (0 == 42) {
                                        v80 := GenS10{F0: 2, F1: (fi12 | fi12), F2: false}
                                        if true {
                                            fmt.Println(v4, v4, v4)
                                        }
                                        v14 = append(v14, fi12)
                                        _ = v80
                                    } else {
                                        var v81 int8 = -1
                                        _ = v81
                                    }
                                    fmt.Println(v3, v5, v5)
                                default:
                                    v16 = v4
                                }
                            default:
                                if false {
                                    v82 := GenS7{F0: 255, F1: 204, F2: (^(v4))}
                                    if true {
                                        v5 = v5
                                        v83 := GenS7{F0: (v2 * 255), F1: 2147483648, F2: (v4 << ((255) & 63))}
                                        v16 += (v16 % (((v4 ^ 169)) | 1))
                                        v4 += (v4 ^ (v16 % ((v16) | 1)))
                                        _ = v83
                                    } else {
                                        if (true || false) {
                                            v4 += 224
                                            switch ("foo bar" + ((("a" + ("" + "\n\t")) + "hello") + ("foo bar" + "foo bar"))) {
                                            case "foo bar":
                                                v16 = v16
                                            default:
                                                var v84 float64 = 1.5
                                                switch "hello" {
                                                case "foo bar":
                                                    var v85 float32 = -2.25
                                                    _ = v85
                                                }
                                                v84 = (v84 - v84)
                                                v4 = v16
                                                _ = v84
                                            }
                                            v2 -= (v2 >> ((902) & 63))
                                            if (!true) {
                                                fmt.Println(fi12, v5)
                                            } else {
                                                v4 = 910
                                            }
                                        }
                                    }
                                    _ = v82
                                } else {
                                    var v86 uint32 = 42
                                    _ = len(v1)
                                    if true {
                                        var v87 int64 = 591
                                        v88 := v6["hello"]
                                        _ = v87
                                        _ = v88
                                    }
                                    _ = v86
                                }
                            }
                            _ = v16
                        case 128:
                            v89 := GenS11{F0: "hello", F1: 65535, F2: 19}
                            _ = v89
                        default:
                            if false {
                                if ((1 != 42) && true) {
                                    v90 := GenS10{F0: uint32(v3), F1: fi12, F2: (-15 <= 127)}
                                    v91 := v1[fi12]
                                    _ = v90
                                    _ = v91
                                }
                                var v92 float64 = float64(1e10)
                                _ = v92
                            } else {
                                v93 := GenS7{F0: v2, F1: 42, F2: 255}
                                v94 := v6[""]
                                var v95 int64 = 42
                                _ = v93
                                _ = v94
                                _ = v95
                            }
                        }
                        delete(v1, fi12)
                        var v96 uint8 = v2
                        _ = v96
                    } else {
                        var v97 int16 = v5
                        var v98 uint8 = (v2 * 50)
                        v99 := GenS7{F0: (v98 + (v3 - v3)), F1: 751, F2: v4}
                        _ = v97
                        _ = v98
                        _ = v99
                    }
                    var v100 int = 32768
                    var v101 uint = 4294967295
                    _ = v100
                    _ = v101
                }
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        fmt.Println(v2)
        switch (("foo bar" + "\n\t") + ("foo bar" + "")) {
        case "a":
            var v102 uint32 = 601
            _ = v102
        }
    }
    if false {
        v3 = v3
        fmt.Println(v4, v5, v3)
    }
    if ((true && ((933 != -1) && false)) || (-104 > 127)) {
        v4 = v4
        for fi103 := 0; fi103 < 3; fi103++ {
            for fi104 := 0; fi104 < 3; fi104++ {
                if (!(-2.25 < (-2.25 + ((0.1 / (((3.14159 / (-2.25 - 3.14159)) * (-2.25 - 0.0)) * 0.1)) / (0.1 / 0.1))))) {
                    fmt.Println(v4, v4, v5)
                } else {
                    var v105 int16 = -1
                    _ = v105
                }
                switch "foo bar" {
                case "foo bar":
                    var v106 int8 = 42
                    _ = v106
                default:
                    switch "hello" {
                    case "foo bar":
                        var v107 int = fi103
                        _ = v107
                    default:
                        if true {
                            fmt.Println(v4)
                            fmt.Println(fi104, v3)
                            if (-686 == 0) {
                                v4 = (v4 ^ 32767)
                                v108 := v1[fi104]
                                v109 := GenS8{F0: (550 <= 0), F1: uint8(v4), F2: -27}
                                _ = v108
                                _ = v109
                            }
                        }
                    }
                }
                _ = fi104
            }
            var v110 uint16 = (v4 << ((4294967295) & 63))
            _ = fi103
            _ = v110
        }
    }
    v111 := GenS5{F0: 127}
    v3 = 128
    v2 -= (v3 + (^(v3)))
    switch "a" {
    case "\n\t":
        fmt.Println(v4, v5)
    }
    for fi112 := 0; fi112 < 1; fi112++ {
        switch fi112 {
        case 1:
            v4 += (^(v4))
        case 156:
            v113 := GenS11{F0: "\n\t", F1: 451, F2: int8(fi112)}
            _ = v113
        default:
            for fi114 := 0; fi114 < 9; fi114++ {
                var v115 bool = true
                switch (^(fi114)) {
                case -865:
                    var v116 float32 = -2.25
                    _ = v116
                }
                if (v115 || false) {
                    var v117 uint = 18446744073709551615
                    v2 = (v3 / (((v3 * v3)) | 1))
                    fmt.Println(fi114, v115)
                    _ = v117
                }
                fmt.Println(v2, v115)
                var v118 float32 = 1e10
                _ = fi114
                _ = v115
                _ = v118
            }
        }
        _ = fi112
    }
    switch (("a" + ("\n\t" + "\n\t")) + "\n\t") {
    case "hello":
        for fi119 := 0; fi119 < 11; fi119++ {
            delete(v6, ("a" + "foo bar"))
            switch "\n\t" {
            case "\n\t":
                _ = len(v1)
            case "a":
                v5 = v5
            }
            _ = fi119
        }
    case "\n\t":
        for fi120 := 0; fi120 < 5; fi120++ {
            v121 := []int{-278, (fi120 ^ 42)}
            var v122 bool = (false && ((true && false) && true))
            for fi123 := 0; fi123 < 15; fi123++ {
                switch (fi123 % ((fi120) | 1)) {
                case 301:
                    var v124 uint16 = v4
                    _ = v124
                case 1:
                    _ = len(v121)
                case 9223372036854775807:
                    v125 := GenS1{F0: (fi120 ^ fi120), F1: 4294967295, F2: 963}
                    _ = v125
                default:
                    v5 = v5
                }
                v126 := v1[-9223372036854775808]
                _ = fi123
                _ = v126
            }
            _ = fi120
            _ = v121
            _ = v122
        }
    case "foo bar":
        var v127 uint64 = 236
        _ = v127
    }
    for fi128 := 0; fi128 < 12; fi128++ {
        switch "" {
        case "":
            v2 -= v3
        case "a":
            var v129 uint8 = 255
            _ = v129
        case "\n\t":
            v4 = 65535
        default:
            if (true || ((v4 >= 1) && (!(!false)))) {
                var v130 int16 = (v5 / ((1) | 1))
                v131 := []int{fi128}
                _ = v130
                _ = v131
            } else {
                v132 := []int{fi128, fi128}
                var v133 uint = 844
                _ = v132
                _ = v133
            }
        }
        _ = fi128
    }
    var v134 int16 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v134:", v134)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v111
    _ = v134
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
