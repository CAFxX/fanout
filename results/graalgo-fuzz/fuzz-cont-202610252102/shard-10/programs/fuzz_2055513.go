package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 uint32; F2 int }
type GenS7 struct { F0 int16; F1 bool; F2 string }
type GenS8 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint; F1 uint }
type GenS14 struct { F0 uint32 }
func genFn0() int8 {
    var v7 uint16 = 65535
    v8 := GenS1{F0: 0, F1: 32767, F2: 42}
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int64) bool {
    for fi9 := 0; fi9 < 15; fi9++ {
        switch fi9 {
        case 4294967295:
            v10 := GenS2{F0: (-(fi9)), F1: -1, F2: (p0 + 9223372036854775807)}
            _ = v10
        case 210:
            var v11 string = "hello"
            _ = v11
        default:
            v12 := GenS2{F0: fi9, F1: -566, F2: p0}
            _ = v12
        }
        v13 := GenS3{F0: 2}
        var v14 uint8 = 0
        _ = fi9
        _ = v13
        _ = v14
    }
    var v15 float64 = 1e10
    fmt.Println(p0)
    _ = p0
    _ = v15
    return (0 == 2)
}

func main() {
    v1 := []string{"", ("\n\t" + ((("\n\t" + "foo bar") + ("" + "foo bar")) + (("foo bar" + ("foo bar" + "")) + "\n\t"))), ("hello" + "hello"), ""}
    v2 := make(map[int]uint32)
    v3 := []int{-877, 1, 32767}
    var v4 uint8 = 255
    v5 := []int{2147483647, -631, 1}
    v6 := make(map[string]int)
    v3 = append(v3, 32768)
    v3 = append(v3, 4294967295)
    if len(v3) > 0 {
        v3[0] = 623
    }
    if true {
        switch 1 {
        case -1:
            v1 = append(v1, "")
        case 127:
            if (387 == 255) {
                fmt.Println(v4)
                var v16 int = 0
                switch -335 {
                case 921:
                    if (!false) {
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v4 += (v4 - v4)
                            _ = fi17
                        }
                        if true {
                            switch "hello" {
                            case "":
                                fmt.Println(v4)
                            case "foo bar":
                                v16 = (v16 - v16)
                            case "\n\t":
                                v18 := GenS3{F0: (v4 - v4)}
                                _ = v18
                            default:
                                for fi19 := 0; fi19 < 16; fi19++ {
                                    if true {
                                        v4 -= v4
                                        v1 = append(v1, "foo bar")
                                    }
                                    if (true || (!false)) {
                                        fmt.Println(v16)
                                        for fi20 := 0; fi20 < 15; fi20++ {
                                            var v21 int = (^(fi20))
                                            if true {
                                                v21 = 127
                                                if (0 <= -77) {
                                                    var v22 int16 = 1
                                                    v23 := GenS1{F0: 9223372036854775807, F1: 2, F2: v4}
                                                    if (!(!(true || (1.5 > 1.5)))) {
                                                        var v24 float64 = (float64(-2.25) - 3.14159)
                                                        var v25 uint32 = 757
                                                        v3 = append(v3, fi20)
                                                        _ = v24
                                                        _ = v25
                                                    } else {
                                                        var v26 uint32 = 526
                                                        _ = v26
                                                    }
                                                    v27 := GenS1{F0: 9223372036854775807, F1: 65535, F2: (^(v4))}
                                                    _ = v22
                                                    _ = v23
                                                    _ = v27
                                                } else {
                                                    switch (fi20 / (((-(fi20))) | 1)) {
                                                    case -354:
                                                        v28 := GenS2{F0: -768, F1: 74, F2: 42}
                                                        _ = v28
                                                    case -734:
                                                        var v29 string = "a"
                                                        _ = v29
                                                    default:
                                                        var v30 int = (v16 ^ 29)
                                                        _ = v30
                                                    }
                                                }
                                            } else {
                                                switch ("\n\t" + "") {
                                                case "a":
                                                    v21 = fi20
                                                case "hello":
                                                    v21 += (v21 << ((65535) & 63))
                                                case "foo bar":
                                                    fmt.Println(v16, fi19, fi19)
                                                default:
                                                    if true {
                                                        if ("foo bar" < "\n\t") {
                                                            var v31 int8 = 1
                                                            v31 = 14
                                                            var v32 float64 = (0.1 - 1.5)
                                                            v4 += (v4 | v4)
                                                            _ = v31
                                                            _ = v32
                                                        } else {
                                                            _ = len(v2)
                                                        }
                                                        if (!(!false)) {
                                                            fmt.Println(v4, v21)
                                                            switch (v21 - v21) {
                                                            case 32768:
                                                                var v33 int16 = 42
                                                                _ = v33
                                                            default:
                                                                v34 := GenS1{F0: 1, F1: 1, F2: 216}
                                                                _ = v34
                                                            }
                                                        } else {
                                                            v21 -= v16
                                                            if true {
                                                                var v35 uint32 = 0
                                                                _ = v35
                                                            } else {
                                                                if true {
                                                                    var v36 uint32 = 2147483647
                                                                    if ((fi19 >> ((uint(v36)) & 63)) >= v16) {
                                                                        v21 = (fi19 * (v21 % (((v21 << ((1) & 63))) | 1)))
                                                                        v37 := GenS1{F0: -1, F1: 2, F2: (v4 >> ((362) & 63))}
                                                                        _ = v37
                                                                    }
                                                                    _ = v36
                                                                } else {
                                                                    v2[127] = 4294967295
                                                                }
                                                                var v38 int16 = -1
                                                                var v39 uint16 = 541
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            fmt.Println(v4, v4, v21)
                                                            v3 = append(v3, fi20)
                                                            v1 = append(v1, "foo bar")
                                                        }
                                                        v4 -= (v4 - v4)
                                                    } else {
                                                        fmt.Println(v4)
                                                        v4 += v4
                                                        var v40 int8 = 16
                                                        switch (("\n\t" + "foo bar") + ("foo bar" + ("foo bar" + "\n\t"))) {
                                                        case "\n\t":
                                                            switch "a" {
                                                            case "\n\t":
                                                                if ((fi20 + fi20) != (v16 >> ((188) & 63))) {
                                                                    var v41 uint8 = (v4 % ((128) | 1))
                                                                    v21 = fi20
                                                                    var v42 uint16 = 534
                                                                    _ = v41
                                                                    _ = v42
                                                                }
                                                            }
                                                        case "foo bar":
                                                            if len(v5) > 0 {
                                                                v5[0] = 42
                                                            }
                                                        case "a":
                                                            var v43 uint = 4294967295
                                                            _ = v43
                                                        }
                                                        v21 = (v21 - 127)
                                                        _ = v40
                                                    }
                                                }
                                            }
                                            v21 = v21
                                            var v44 bool = (2147483647 == 2)
                                            _ = fi20
                                            _ = v21
                                            _ = v44
                                        }
                                    } else {
                                        if (32767 > 42) {
                                            if false {
                                                v45 := GenS1{F0: 9223372036854775807, F1: 969, F2: 1}
                                                _ = v45
                                            }
                                            v46 := GenS4{F0: true, F1: (fi19 * v16)}
                                            for fi47 := 0; fi47 < 17; fi47++ {
                                                v48 := GenS5{F0: 470}
                                                var v49 int16 = 1
                                                v50 := GenS6{F0: v4, F1: 4294967295, F2: v16}
                                                var v51 int32 = 0
                                                v16 += fi47
                                                _ = fi47
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            }
                                            _ = v46
                                        } else {
                                            v52 := GenS5{F0: 392}
                                            v53 := GenS7{F0: -173, F1: true, F2: ("" + "")}
                                            for fi54 := 0; fi54 < 2; fi54++ {
                                                v55 := GenS6{F0: 153, F1: 603, F2: fi54}
                                                _ = len(v2)
                                                _ = fi54
                                                _ = v55
                                            }
                                            _ = v52
                                            _ = v53
                                        }
                                    }
                                    _ = fi19
                                }
                            }
                            v16 = v16
                            fmt.Println(v16, v16)
                        }
                        if false {
                            if false {
                                v16 -= 32767
                            } else {
                                var v56 int = (v16 * int(v4))
                                _ = v56
                            }
                            if true {
                                v16 = v16
                            } else {
                                for fi57 := 0; fi57 < 12; fi57++ {
                                    v58 := GenS6{F0: (v4 ^ (v4 + 42)), F1: 2147483647, F2: v16}
                                    _ = len(v6)
                                    _ = fi57
                                    _ = v58
                                }
                                if ((v16 * 699) != (v16 / (((v16 ^ v16)) | 1))) {
                                    v16 -= 32768
                                }
                            }
                        }
                    } else {
                        v16 -= (v16 | v16)
                    }
                case -236:
                    _ = len(v2)
                default:
                    v59 := GenS1{F0: int64(v16), F1: 814, F2: (v4 | (v4 | 166))}
                    _ = v59
                }
                _ = v16
            } else {
                var v60 uint = 4294967295
                _ = v60
            }
        }
    }
    switch -673 {
    case 32768:
        v61 := GenS5{F0: 232}
        _ = v61
    case 1:
        var v62 uint64 = 0
        _ = v62
    }
    v4 -= v4
    for fi63 := 0; fi63 < 4; fi63++ {
        if (138 >= 0) {
            fmt.Println(v4, fi63)
            v64 := v2[fi63]
            v2[fi63] = v64
            _ = v64
        } else {
            for fi65 := 0; fi65 < 1; fi65++ {
                v4 += v4
                v66 := GenS7{F0: -32768, F1: (42 < 632), F2: ((("" + ("" + ("a" + ("foo bar" + "hello")))) + "foo bar") + ("a" + "foo bar"))}
                v4 += v4
                _ = fi65
                _ = v66
            }
            delete(v6, "")
        }
        if len(v1) > 0 {
            v1[0] = "a"
        }
        _ = fi63
    }
    fmt.Println(v4)
    for fi67 := 0; fi67 < 14; fi67++ {
        var v68 int32 = 0
        _ = fi67
        _ = v68
    }
    v4 -= v4
    var v69 int = -376
    var v70 int32 = 0
    if true {
        if (!true) {
            var v71 int16 = -1
            _ = v71
        } else {
            if ("foo bar" != "\n\t") {
                var v72 int32 = v70
                if (-47 < 1) {
                    v73 := v2[32767]
                    fmt.Println(v72)
                    _ = v73
                } else {
                    switch "\n\t" {
                    case "a":
                        for fi74 := 0; fi74 < 13; fi74++ {
                            v70 = (v72 ^ (v72 + (-(v70))))
                            _ = fi74
                        }
                    case "foo bar":
                        delete(v6, ("" + (("foo bar" + "") + "\n\t")))
                    case "hello":
                        var v75 uint = 65535
                        _ = v75
                    }
                }
                _ = v72
            } else {
                switch ("" + (("hello" + "a") + "hello")) {
                case "hello":
                    var v76 uint16 = 0
                    _ = v76
                case "\n\t":
                    var v77 uint32 = 2
                    _ = v77
                case "":
                    var v78 int64 = 1
                    _ = v78
                }
                v69 = (v69 ^ 42)
                var v79 float64 = 1e10
                _ = v79
            }
        }
        if (!false) {
            for fi80 := 0; fi80 < 5; fi80++ {
                switch (fi80 / ((v69) | 1)) {
                case 926:
                    var v81 int = (v69 | fi80)
                    _ = v81
                case 633:
                    switch 128 {
                    case 507:
                        switch ("\n\t" + "\n\t") {
                        case "":
                            switch "\n\t" {
                            case "a":
                                var v82 float32 = float32(-2.25)
                                _ = v82
                            case "":
                                _ = len(v6)
                            }
                        case "\n\t":
                            for fi83 := 0; fi83 < 15; fi83++ {
                                _ = len(v1)
                                if ((true || (float32(0.1) == 1.5)) || true) {
                                    if (false || (false || (1e10 >= -2.25))) {
                                        v4 = 42
                                        v70 = v70
                                    } else {
                                        switch (fi83 - -874) {
                                        case -1:
                                            var v84 uint32 = 7
                                            _ = v84
                                        }
                                    }
                                } else {
                                    switch 0 {
                                    case 128:
                                        v70 -= -1
                                    case 32768:
                                        v4 = v4
                                    }
                                    _ = len(v6)
                                    switch v69 {
                                    case -1:
                                        if (true || ((true || true) && (true || (0.0 < (float64(3.14159) * ((3.14159 * (0.1 / 1e10)) / 1.0)))))) {
                                            var v85 int16 = 0
                                            var v86 float64 = 1.5
                                            switch (fi80 >> ((18446744073709551615) & 63)) {
                                            case 634:
                                                var v87 float32 = ((1e10 / 1e10) + -2.25)
                                                _ = v87
                                            case 733:
                                                v88 := GenS1{F0: 1, F1: 32767, F2: v4}
                                                _ = v88
                                            }
                                            _ = v85
                                            _ = v86
                                        } else {
                                            v4 = 14
                                            v89 := v2[fi80]
                                            v90 := GenS6{F0: 78, F1: 2147483647, F2: -1}
                                            _ = v89
                                            _ = v90
                                        }
                                    }
                                }
                                switch "\n\t" {
                                case "hello":
                                    if (2147483648 == 2147483647) {
                                        var v91 int8 = 73
                                        _ = v91
                                    }
                                default:
                                    v4 = 65
                                }
                                if false {
                                    switch "a" {
                                    case "":
                                        switch "" {
                                        case "foo bar":
                                            if true {
                                                v92 := GenS7{F0: -1, F1: false, F2: ("" + "a")}
                                                _ = v92
                                            } else {
                                                v69 = (fi83 * (v69 + fi83))
                                                v4 = (^(v4))
                                                if (42 <= 987) {
                                                    var v93 uint32 = 65535
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v94 uint32 = v93
                                                        _ = v94
                                                    case "":
                                                        if false {
                                                            if len(v5) > 0 {
                                                                v5[0] = (fi80 & (-(fi80)))
                                                            }
                                                            v93 = (^(v93))
                                                        }
                                                    default:
                                                        switch (-(fi83)) {
                                                        case 128:
                                                            if false {
                                                                if (false && (!(4294967295 >= 42))) {
                                                                    var v95 int32 = (v70 % ((0) | 1))
                                                                    var v96 int32 = v95
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                                v93 = (v93 | v93)
                                                            }
                                                        case 2147483647:
                                                            v4 -= (v4 & v4)
                                                        }
                                                    }
                                                    var v97 int64 = -9223372036854775808
                                                    _ = v93
                                                    _ = v97
                                                } else {
                                                    v98 := GenS6{F0: (v4 | v4), F1: 4294967295, F2: (fi80 >> ((326) & 63))}
                                                    var v99 float64 = 1.5
                                                    if false {
                                                        if (1e10 < 0.0) {
                                                            if true {
                                                                var v100 bool = (65535 <= 1)
                                                                _ = v100
                                                            }
                                                            if false {
                                                                var v101 float32 = -2.25
                                                                var v102 float32 = float32((v99 * v99))
                                                                var v103 uint8 = 128
                                                                _ = v101
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                        }
                                                        var v104 int = -811
                                                        var v105 int32 = -927
                                                        _ = v104
                                                        _ = v105
                                                    } else {
                                                        v69 = v69
                                                        v70 -= v70
                                                        switch "a" {
                                                        case "hello":
                                                            var v106 float32 = (3.14159 + -2.25)
                                                            _ = v106
                                                        case "":
                                                            switch (v69 | 673) {
                                                            case 0:
                                                                if true {
                                                                    if (!true) {
                                                                        v69 -= 32767
                                                                        var v107 uint8 = v4
                                                                        var v108 float32 = 1.5
                                                                        v69 = fi83
                                                                        switch (v69 & (v69 - fi83)) {
                                                                        case 1:
                                                                            if ((!false) || false) {
                                                                                switch (v69 >> ((18446744073709551615) & 63)) {
                                                                                case 32767:
                                                                                    if (!true) {
                                                                                        var v109 float32 = 1.5
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = fi83
                                                                                        }
                                                                                        fmt.Println(v107, v69, fi80)
                                                                                        _ = v109
                                                                                    }
                                                                                case 2147483647:
                                                                                    var v110 uint = 4294967295
                                                                                    _ = v110
                                                                                case 1:
                                                                                    var v111 uint16 = 0
                                                                                    _ = v111
                                                                                default:
                                                                                    var v112 uint64 = 42
                                                                                    _ = v112
                                                                                }
                                                                            } else {
                                                                                var v113 int32 = -744
                                                                                fmt.Println(v70, v113)
                                                                                var v114 int64 = -1
                                                                                _ = v113
                                                                                _ = v114
                                                                            }
                                                                        }
                                                                        _ = v107
                                                                        _ = v108
                                                                    }
                                                                    var v115 uint16 = 32767
                                                                    _ = v115
                                                                }
                                                            case -2147483648:
                                                                if (773 > 0) {
                                                                    if false {
                                                                        _ = len(v1)
                                                                        var v116 int32 = (v70 - (v70 % ((-1) | 1)))
                                                                        switch (^(fi83)) {
                                                                        case -817:
                                                                            var v117 uint8 = (v4 ^ (v4 % ((1) | 1)))
                                                                            _ = v117
                                                                        case 282:
                                                                            if ("\n\t" <= ("hello" + "foo bar")) {
                                                                                switch ("foo bar" + ("\n\t" + "")) {
                                                                                case "foo bar":
                                                                                    var v118 int = v69
                                                                                    _ = v118
                                                                                case "":
                                                                                    v69 = v69
                                                                                case "\n\t":
                                                                                    fmt.Println(v69)
                                                                                default:
                                                                                    v119 := GenS8{F0: v4, F1: (v4 / (((v4 - (v4 ^ v4))) | 1)), F2: 0}
                                                                                    _ = v119
                                                                                }
                                                                                v120 := GenS9{F0: "a", F1: 42}
                                                                                fmt.Println(v99)
                                                                                _ = v120
                                                                            } else {
                                                                                switch ("hello" + "") {
                                                                                case "\n\t":
                                                                                    v69 -= 0
                                                                                case "":
                                                                                    if false {
                                                                                        var v121 int8 = 30
                                                                                        v70 = (v70 + v116)
                                                                                        _ = v121
                                                                                    }
                                                                                }
                                                                                fmt.Println(v69, fi80, v99)
                                                                                var v122 int32 = (v70 << ((255) & 63))
                                                                                _ = v122
                                                                            }
                                                                        case -948:
                                                                            v69 = fi83
                                                                        default:
                                                                            v123 := v2[(^(fi83))]
                                                                            _ = v123
                                                                        }
                                                                        _ = v116
                                                                    } else {
                                                                        if (((1 >= v4) || (-128 == 100)) || false) {
                                                                            v69 = fi80
                                                                            if (!false) {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            } else {
                                                                                switch (fi83 % (((fi80 & v69)) | 1)) {
                                                                                case -806:
                                                                                    var v124 int = fi83
                                                                                    _ = v124
                                                                                default:
                                                                                    v99 = v99
                                                                                }
                                                                            }
                                                                            var v125 int32 = (v70 << ((693) & 63))
                                                                            _ = v125
                                                                        }
                                                                        var v126 uint = 255
                                                                        if true {
                                                                            var v127 int8 = genFn0()
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                switch (v69 + fi83) {
                                                                                case 32767:
                                                                                    v128 := GenS10{F0: (v127 * v127), F1: (^(v126))}
                                                                                    _ = v128
                                                                                }
                                                                            case "a":
                                                                                v99 = (((0.0 * v99) * v99) - v99)
                                                                            default:
                                                                                v129 := GenS7{F0: -683, F1: false, F2: "hello"}
                                                                                _ = v129
                                                                            }
                                                                            var v130 string = ""
                                                                            _ = v127
                                                                            _ = v130
                                                                        }
                                                                        _ = v126
                                                                    }
                                                                    v131 := GenS8{F0: (v4 ^ v4), F1: v4, F2: 4294967295}
                                                                    switch fi83 {
                                                                    case -2147483648:
                                                                        v70 = (v70 * v70)
                                                                    case -208:
                                                                        v132 := GenS3{F0: 1}
                                                                        _ = v132
                                                                    case 42:
                                                                        var v133 uint = 65535
                                                                        _ = v133
                                                                    }
                                                                    _ = v131
                                                                }
                                                            case -9223372036854775808:
                                                                var v134 int8 = 127
                                                                _ = v134
                                                            }
                                                        default:
                                                            if true {
                                                                if false {
                                                                    fmt.Println(fi83, v99)
                                                                    if true {
                                                                        v1 = append(v1, ("hello" + "foo bar"))
                                                                        var v135 float32 = ((float32(0.0) + 0.1) + (1e10 - float32(float64(((((0.0 + (float32(v99) + -2.25)) / 0.1) + -2.25) + 1.5)))))
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        v136 := GenS1{F0: -1, F1: 811, F2: 122}
                                                                        var v137 uint64 = 469
                                                                        _ = v135
                                                                        _ = v136
                                                                        _ = v137
                                                                    }
                                                                    switch fi83 {
                                                                    case 2147483647:
                                                                        v69 -= fi83
                                                                    case -997:
                                                                        switch 9223372036854775807 {
                                                                        case 285:
                                                                            _ = len(v5)
                                                                            v138 := v2[(fi83 - (-(v69)))]
                                                                            _ = v138
                                                                        case -776:
                                                                            var v139 string = "foo bar"
                                                                            _ = v139
                                                                        case 4294967295:
                                                                            var v140 uint32 = 2
                                                                            _ = v140
                                                                        }
                                                                    case 319:
                                                                        var v141 string = "\n\t"
                                                                        _ = v141
                                                                    }
                                                                }
                                                                v142 := GenS6{F0: (v4 >> ((255) & 63)), F1: 396, F2: (fi83 ^ 4294967295)}
                                                                v70 += -1
                                                                _ = v142
                                                            } else {
                                                                var v143 bool = (!true)
                                                                v144 := GenS3{F0: 127}
                                                                var v145 int = (fi83 | (fi80 | -968))
                                                                _ = v143
                                                                _ = v144
                                                                _ = v145
                                                            }
                                                        }
                                                        var v146 int = (fi80 % ((0) | 1))
                                                        v147 := GenS11{F0: true, F1: 725}
                                                        _ = v146
                                                        _ = v147
                                                    }
                                                    fmt.Println(fi80, v99, fi80)
                                                    v69 += -523
                                                    _ = v98
                                                    _ = v99
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = "hello"
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = -1
                                                }
                                            }
                                        case "":
                                            v148 := GenS12{F0: 2}
                                            _ = v148
                                        }
                                    default:
                                        var v149 uint16 = 484
                                        _ = v149
                                    }
                                    var v150 uint = 472
                                    _ = v150
                                }
                                v70 = 2147483647
                                _ = fi83
                            }
                        case "hello":
                            _ = len(v6)
                        default:
                            switch "" {
                            case "hello":
                                v4 = 1
                            }
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case -955:
                    var v151 int = (v69 / (((fi80 - fi80)) | 1))
                    _ = v151
                default:
                    fmt.Println(fi80)
                }
                if (65535 > 2147483648) {
                    v69 -= fi80
                    for fi152 := 0; fi152 < 1; fi152++ {
                        v2[(-(fi80))] = 2
                        _ = fi152
                    }
                } else {
                    var v153 float64 = (((3.14159 / (3.14159 * 1.5)) + -2.25) + 0.1)
                    var v154 bool = (false || (true || (fi80 > (^(fi80)))))
                    v154 = (!("\n\t" != ("foo bar" + "hello")))
                    _ = v153
                    _ = v154
                }
                fmt.Println(v4, v70, v4)
                _ = fi80
            }
            var v155 bool = true
            _ = v155
        } else {
            v156 := GenS13{F0: 0, F1: 1}
            fmt.Println(v69, v70, v4)
            _ = v156
        }
    }
    v6["a"] = -2147483648
    v4 = v4
    if true {
        v4 = v4
    }
    var v157 int32 = 926
    if false {
        var v158 float32 = (1.5 - 1.5)
        switch "foo bar" {
        case "":
            v4 -= (v4 % ((v4) | 1))
        case "hello":
            var v159 uint16 = 255
            _ = v159
        case "foo bar":
            fmt.Println(v157, v70)
        }
        for fi160 := 0; fi160 < 17; fi160++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi161 := 0; fi161 < 9; fi161++ {
                var v162 uint32 = 2
                v4 = v4
                if false {
                    if false {
                        v163 := GenS14{F0: 2}
                        switch 32768 {
                        case -507:
                            if (float64(((v158 + v158) * -2.25)) < ((3.14159 / 1.0) / -2.25)) {
                                var v164 uint = 18446744073709551615
                                if (true || (!true)) {
                                    v2[v69] = v162
                                    var v165 uint64 = 42
                                    _ = v165
                                }
                                _ = v164
                            } else {
                                v166 := v2[32768]
                                _ = v166
                            }
                        case 42:
                            v157 += (v157 % ((-569) | 1))
                        default:
                            var v167 uint64 = 0
                            _ = v167
                        }
                        _ = v163
                    } else {
                        v70 += v157
                    }
                    var v168 float32 = (((v158 + 3.14159) / 3.14159) + (((v158 + 3.14159) - float32(0.0)) - 0.0))
                    _ = v168
                }
                _ = fi161
                _ = v162
            }
            for fi169 := 0; fi169 < 17; fi169++ {
                if len(v3) > 0 {
                    v3[0] = v69
                }
                v170 := GenS2{F0: fi160, F1: -845, F2: -9223372036854775808}
                _ = fi169
                _ = v170
            }
            _ = fi160
        }
        _ = v158
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v157:", v157)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v157
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
}
