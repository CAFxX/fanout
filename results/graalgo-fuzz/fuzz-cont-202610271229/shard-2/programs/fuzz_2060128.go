package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int8; F2 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 bool; F2 int }
type GenS7 struct { F0 uint; F1 int32; F2 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 int64; F2 int16 }
type GenS10 struct { F0 int8; F1 int; F2 uint32 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 bool; F1 uint64 }
type GenS13 struct { F0 uint; F1 uint64; F2 int8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint8; F1 uint64 }
type GenS16 struct { F0 string; F1 uint16 }
type GenS17 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    var v3 int = 32767
    var v4 float32 = ((-2.25 / 3.14159) + 0.0)
    var v5 int16 = 1
    switch "" {
    case "a":
        var v6 int8 = 100
        _ = v6
    case "hello":
        switch "\n\t" {
        case "a":
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 uint8 = 228
                _ = fi7
                _ = v8
            }
        case "hello":
            switch "\n\t" {
            case "":
                var v9 uint8 = 1
                _ = v9
            }
        case "\n\t":
            v10 := v1[0]
            _ = v10
        }
    default:
        var v11 int16 = (^(v5))
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 float32 = v4
        switch "foo bar" {
        case "":
            if false {
                if false {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v5 = (v5 ^ -32768)
                        fmt.Println(fi14, v13, v5)
                        if true {
                            v3 -= fi14
                            v15 := []int{(fi14 << ((65535) & 63)), (-(fi14)), -41}
                            _ = v15
                        }
                        if (0 <= 42) {
                            var v16 float32 = 1e10
                            switch "hello" {
                            case "a":
                                if false {
                                    if (1 > 18446744073709551615) {
                                        if true {
                                            v3 = v3
                                        }
                                        var v17 int32 = -1
                                        v4 -= v13
                                        _ = v17
                                    }
                                    if (true && (!(!(true && (false || (true && ((float64(((0.0 / 1.5) - 1e10)) / 0.1) > 0.0))))))) {
                                        v18 := GenS1{F0: "foo bar"}
                                        fmt.Println(fi14)
                                        _ = v18
                                    }
                                }
                            default:
                                var v19 float32 = 3.14159
                                _ = v19
                            }
                            _ = v16
                        } else {
                            if (545 > 441) {
                                var v20 int8 = 100
                                v2[42] = false
                                var v21 int8 = v20
                                _ = v20
                                _ = v21
                            } else {
                                if (("" != "\n\t") || (-32768 >= (v5 + v5))) {
                                    v5 = v5
                                    var v22 int = (v3 / (((v3 >> ((21) & 63))) | 1))
                                    fmt.Println(v4, v5)
                                    _ = v22
                                } else {
                                    if (!(false && true)) {
                                        if (100 == 100) {
                                            v13 += v4
                                            fmt.Println(fi12)
                                            fmt.Println(v13, v3)
                                            v13 -= 1e10
                                        }
                                    }
                                    switch "\n\t" {
                                    case "":
                                        v1[0] = "\n\t"
                                        var v23 uint16 = 255
                                        _ = v23
                                    }
                                    var v24 uint8 = 42
                                    v5 = 32767
                                    v5 -= 0
                                    _ = v24
                                }
                                var v25 int64 = 970
                                v26 := GenS1{F0: ("" + "foo bar")}
                                _ = v25
                                _ = v26
                            }
                            switch "" {
                            case "hello":
                                var v27 float64 = (0.1 / 1.5)
                                _ = v27
                            case "":
                                v28 := []int{v3, fi12, (-(fi14))}
                                v5 += 32767
                                _ = v28
                            default:
                                v1[2147483648] = "foo bar"
                            }
                            var v29 int = -9223372036854775808
                            _ = v29
                        }
                        var v30 uint16 = 2
                        _ = fi14
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 14; fi31++ {
                        v13 = v4
                        _ = fi31
                    }
                } else {
                    v32 := GenS2{F0: (957 >= 0)}
                    for fi33 := 0; fi33 < 4; fi33++ {
                        v3 = -29
                        if (2147483647 < int32(v5)) {
                            v4 = ((0.0 + 3.14159) * 1e10)
                            fmt.Println(v4, v4, v5)
                            var v34 int8 = -1
                            _ = v34
                        } else {
                            fmt.Println(fi33, fi12, v3)
                            if (1000 > (v5 & (v5 - v5))) {
                                v5 -= v5
                                var v35 int = (fi12 | 32768)
                                delete(v2, 589)
                                _ = v35
                            }
                            switch (((("" + "\n\t") + ("foo bar" + "a")) + "foo bar") + ("" + "")) {
                            case "foo bar":
                                fmt.Println(v13)
                            case "":
                                var v36 int64 = -888
                                _ = v36
                            }
                        }
                        _ = fi33
                    }
                    if ((298 == 4294967295) && false) {
                        switch int(v5) {
                        case 32767:
                            if true {
                                for fi37 := 0; fi37 < 10; fi37++ {
                                    fmt.Println(v13, v3, v3)
                                    var v38 uint64 = 1
                                    fmt.Println(v4, v5)
                                    _ = fi37
                                    _ = v38
                                }
                            }
                        }
                        switch (fi12 + fi12) {
                        case -2147483648:
                            v39 := []int{(fi12 - (v3 & v3))}
                            _ = v39
                        default:
                            if false {
                                v40 := GenS1{F0: ("a" + "\n\t")}
                                var v41 int16 = (^(v5))
                                if (1 < 256) {
                                    v41 -= v41
                                    if (0 >= 1) {
                                        var v42 int16 = v5
                                        var v43 float64 = -2.25
                                        _ = v42
                                        _ = v43
                                    }
                                } else {
                                    for fi44 := 0; fi44 < 10; fi44++ {
                                        v45 := GenS2{F0: (1 <= -1)}
                                        _ = fi44
                                        _ = v45
                                    }
                                    var v46 float32 = (v4 - v4)
                                    _ = v46
                                }
                                _ = v40
                                _ = v41
                            }
                        }
                        if (42 < -817) {
                            var v47 int = v3
                            _ = v47
                        } else {
                            switch 4294967295 {
                            case 9223372036854775807:
                                v48 := GenS1{F0: "a"}
                                _ = v48
                            }
                            var v49 uint8 = 42
                            _ = v49
                        }
                    } else {
                        v50 := GenS2{F0: (-2147483648 > int32(v5))}
                        if ((v5 << ((1) & 63)) >= -1) {
                            if false {
                                fmt.Println(v5)
                                v3 = fi12
                                var v51 int16 = (v5 * -1)
                                _ = v51
                            } else {
                                v52 := GenS1{F0: (("hello" + "") + (("\n\t" + (("a" + "\n\t") + "\n\t")) + "foo bar"))}
                                _ = v52
                            }
                            var v53 uint64 = 4294967295
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 13; fi54++ {
                            if false {
                                if true {
                                    if false {
                                        v3 -= v3
                                        var v55 uint64 = 18446744073709551615
                                        var v56 int = (fi12 + (fi54 | (fi54 + -9223372036854775808)))
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v13 += 0.0
                                        var v57 int32 = 42
                                        fmt.Println(v4, v13)
                                        switch ("foo bar" + "\n\t") {
                                        case "\n\t":
                                            delete(v2, 42)
                                        case "hello":
                                            var v58 uint = 18446744073709551615
                                            _ = v58
                                        case "foo bar":
                                            v5 -= -982
                                        }
                                        _ = v57
                                    }
                                    fmt.Println(fi54)
                                    switch -9223372036854775808 {
                                    case 273:
                                        v59 := GenS3{F0: "foo bar"}
                                        v60 := GenS3{F0: ((("hello" + "\n\t") + ("\n\t" + "foo bar")) + "hello")}
                                        fmt.Println(fi54, v13)
                                        _ = v59
                                        _ = v60
                                    default:
                                        switch ("hello" + ("foo bar" + ("\n\t" + "\n\t"))) {
                                        case "hello":
                                            switch "hello" {
                                            case "\n\t":
                                                v3 = (fi12 | 0)
                                            case "hello":
                                                v61 := GenS4{F0: 706, F1: 100, F2: ((0 > 14) && (!((0.0 / ((1e10 / (float64((v4 / v13)) - 3.14159)) / -2.25)) != 3.14159)))}
                                                _ = v61
                                            case "a":
                                                switch (fi12 + 42) {
                                                case 482:
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v62 float64 = 1e10
                                                        _ = v62
                                                    case "hello":
                                                        v13 = 3.14159
                                                    case "foo bar":
                                                        if true {
                                                            v63 := GenS2{F0: false}
                                                            if false {
                                                                v64 := GenS4{F0: 202, F1: -9, F2: (-9223372036854775808 < 42)}
                                                                if ((!false) && true) {
                                                                    switch fi54 {
                                                                    case 32768:
                                                                        var v65 string = (("a" + "hello") + "foo bar")
                                                                        _ = v65
                                                                    case -537:
                                                                        v2[2] = (1 < 382)
                                                                    case 9223372036854775807:
                                                                        v66 := []int{fi12, 127, (fi12 + (-(v3)))}
                                                                        _ = v66
                                                                    }
                                                                    v13 -= 0.1
                                                                    v67 := GenS5{F0: true}
                                                                    _ = v67
                                                                }
                                                                var v68 uint = 42
                                                                var v69 float64 = 0.1
                                                                _ = v64
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            _ = v63
                                                        } else {
                                                            if (!(true || (true || (true && ("\n\t" != ("a" + "\n\t")))))) {
                                                                var v70 float32 = v13
                                                                _ = v70
                                                            }
                                                            if (943 == 1) {
                                                                if false {
                                                                    v5 = (v5 & (v5 + v5))
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            v2[730] = false
                                                                        case "a":
                                                                            delete(v1, 791)
                                                                        }
                                                                    case "a":
                                                                        v5 = v5
                                                                    case "foo bar":
                                                                        v2[65535] = false
                                                                    }
                                                                    switch ("" + "") {
                                                                    case "foo bar":
                                                                        var v71 float32 = 0.1
                                                                        _ = v71
                                                                    }
                                                                } else {
                                                                    v4 += (float32(0.0) + (((v13 - v13) / v4) * 0.0))
                                                                    fmt.Println(v3, v5, v5)
                                                                    v4 += v4
                                                                }
                                                                v72 := v1[128]
                                                                _ = v72
                                                            }
                                                            v73 := []int{(fi54 / ((fi54) | 1))}
                                                            _ = v73
                                                        }
                                                    default:
                                                        v3 = int(v5)
                                                        v4 = v4
                                                        switch (("foo bar" + ((("foo bar" + (("" + "a") + "\n\t")) + "") + "\n\t")) + "hello") {
                                                        case "hello":
                                                            fmt.Println(v3)
                                                            v5 = (v5 | (^(v5)))
                                                            switch (v3 << ((4294967295) & 63)) {
                                                            case 42:
                                                                var v74 uint16 = 2
                                                                _ = v74
                                                            case 2147483647:
                                                                var v75 uint8 = 128
                                                                _ = v75
                                                            case 50:
                                                                v13 = (3.14159 / (float32(0.0) + 3.14159))
                                                            default:
                                                                v76 := []int{}
                                                                _ = v76
                                                            }
                                                        case "\n\t":
                                                            switch ("hello" + "foo bar") {
                                                            case "foo bar":
                                                                var v77 uint32 = 966
                                                                _ = v77
                                                            case "hello":
                                                                v78 := GenS6{F0: 255, F1: false, F2: (fi54 / (((v3 / (((v3 ^ fi12)) | 1))) | 1))}
                                                                _ = v78
                                                            case "\n\t":
                                                                v3 = fi12
                                                                switch 760 {
                                                                case -63:
                                                                    fmt.Println(v4, v3, v3)
                                                                case 128:
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v79 := GenS2{F0: false}
                                                                        _ = v79
                                                                    case "\n\t":
                                                                        switch ("foo bar" + ("hello" + ("a" + "hello"))) {
                                                                        case "a":
                                                                            if (-9223372036854775808 >= 904) {
                                                                                fmt.Println(fi54)
                                                                                var v80 float64 = (((1.5 + 1e10) * 0.1) * -2.25)
                                                                                _ = v80
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(v4)
                                                                        case "":
                                                                            switch ("\n\t" + "") {
                                                                            case "\n\t":
                                                                                if (v13 != float32(0.0)) {
                                                                                    v3 += -1
                                                                                } else {
                                                                                    if false {
                                                                                        fmt.Println(fi54)
                                                                                        v3 += v3
                                                                                        v13 += v13
                                                                                    }
                                                                                    var v81 bool = false
                                                                                    v81 = false
                                                                                    _ = v81
                                                                                }
                                                                            default:
                                                                                v4 = (((float32(3.14159) + float32(3.14159)) - 0.1) - (0.0 / v13))
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                if (42 < 162) {
                                                                                    switch (("a" + "") + "a") {
                                                                                    case "a":
                                                                                        switch fi54 {
                                                                                        case -1:
                                                                                            var v82 uint64 = 66
                                                                                            _ = v82
                                                                                        case 628:
                                                                                            v83 := GenS2{F0: (false || (2 < 1))}
                                                                                            _ = v83
                                                                                        default:
                                                                                            v84 := GenS7{F0: 4294967295, F1: -828, F2: "foo bar"}
                                                                                            _ = v84
                                                                                        }
                                                                                    }
                                                                                    var v85 uint32 = 32
                                                                                    var v86 uint16 = 0
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                                v5 += 0
                                                                                v4 = (v4 - (-2.25 / ((float32(0.0) * (v4 * v13)) + v4)))
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    var v87 float32 = ((v13 - 3.14159) + ((v13 + v4) / -2.25))
                                                                    _ = v87
                                                                }
                                                            default:
                                                                if (18446744073709551615 <= 42) {
                                                                    fmt.Println(v5)
                                                                }
                                                            }
                                                        }
                                                    }
                                                default:
                                                    if true {
                                                        if false {
                                                            if (false || true) {
                                                                delete(v1, 408)
                                                                v88 := GenS6{F0: 127, F1: true, F2: v3}
                                                                fmt.Println(v5, v4)
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v89 := GenS7{F0: 1, F1: 947, F2: "a"}
                                                                    _ = v89
                                                                }
                                                                v90 := GenS7{F0: 18446744073709551615, F1: -1, F2: "a"}
                                                                _ = v88
                                                                _ = v90
                                                            }
                                                            var v91 float64 = (1.5 - 3.14159)
                                                            v13 += ((0.1 - v4) - v13)
                                                            _ = v91
                                                        }
                                                        fmt.Println(fi12, v13)
                                                        v92 := GenS4{F0: 1, F1: -1, F2: false}
                                                        var v93 float64 = 0.0
                                                        _ = v92
                                                        _ = v93
                                                    } else {
                                                        var v94 uint32 = 559
                                                        switch (fi12 + (fi12 % ((fi12) | 1))) {
                                                        case -492:
                                                            v95 := GenS8{F0: 127}
                                                            _ = v95
                                                        }
                                                        _ = v94
                                                    }
                                                }
                                            default:
                                                var v96 float32 = v4
                                                _ = v96
                                            }
                                        default:
                                            v5 -= v5
                                        }
                                    }
                                }
                                v5 = v5
                            } else {
                                var v97 uint16 = 32767
                                var v98 int8 = 100
                                _ = v97
                                _ = v98
                            }
                            v99 := v1[34]
                            v5 = v5
                            _ = fi54
                            _ = v99
                        }
                        _ = v50
                    }
                    _ = v32
                }
                switch ("hello" + "\n\t") {
                case "a":
                    if true {
                        switch "\n\t" {
                        case "\n\t":
                            var v100 uint = 4294967295
                            _ = v100
                        }
                        if true {
                            switch v3 {
                            case 9223372036854775807:
                                for fi101 := 0; fi101 < 10; fi101++ {
                                    v5 -= v5
                                    delete(v2, 72)
                                    var v102 uint = 65535
                                    _ = fi101
                                    _ = v102
                                }
                            }
                            v103 := GenS4{F0: 535, F1: -1, F2: true}
                            v3 = -358
                            _ = v103
                        } else {
                            v13 += 1e10
                            v3 += (fi12 - 9223372036854775807)
                            v3 += (^(v3))
                        }
                        v1[2147483648] = "hello"
                    } else {
                        fmt.Println(fi12, v3, v4)
                    }
                case "":
                    for fi104 := 0; fi104 < 6; fi104++ {
                        fmt.Println(fi104)
                        _ = fi104
                    }
                case "foo bar":
                    for fi105 := 0; fi105 < 12; fi105++ {
                        var v106 int32 = 2147483647
                        var v107 string = ""
                        var v108 int16 = v5
                        _ = fi105
                        _ = v106
                        _ = v107
                        _ = v108
                    }
                default:
                    for fi109 := 0; fi109 < 20; fi109++ {
                        fmt.Println(fi12)
                        _ = fi109
                    }
                }
                fmt.Println(v13, v13)
            } else {
                v3 = (v3 + fi12)
                _ = len(v2)
                for fi110 := 0; fi110 < 20; fi110++ {
                    switch 897 {
                    case -949:
                        if (!(!true)) {
                            delete(v2, 42)
                            if false {
                                v5 = -212
                                var v111 int64 = 42
                                var v112 uint16 = 862
                                _ = v111
                                _ = v112
                            }
                        } else {
                            v4 = 0.1
                            v4 = (0.0 * ((((v4 * (v13 / v4)) * 0.1) - (v4 - (v13 - 1.5))) + v4))
                        }
                    default:
                        var v113 uint = 4294967295
                        _ = v113
                    }
                    if (true || ((736 == 65535) || false)) {
                        if (!(-1 > 9223372036854775807)) {
                            v114 := GenS4{F0: 2147483647, F1: 127, F2: false}
                            fmt.Println(v3)
                            v3 += fi110
                            _ = v114
                        } else {
                            if (0 >= 255) {
                                v115 := GenS9{F0: 903, F1: 1, F2: (v5 + -351)}
                                switch "a" {
                                case "":
                                    if (-1 > 0) {
                                        if (!false) {
                                            var v116 uint8 = 255
                                            fmt.Println(v4)
                                            v117 := []int{}
                                            _ = v116
                                            _ = v117
                                        }
                                        switch "\n\t" {
                                        case "":
                                            v13 -= v4
                                        case "\n\t":
                                            v118 := []int{(fi110 & fi12)}
                                            _ = v118
                                        default:
                                            v119 := GenS8{F0: 100}
                                            _ = v119
                                        }
                                        switch "a" {
                                        case "hello":
                                            var v120 uint64 = 106
                                            _ = v120
                                        case "foo bar":
                                            if (42 > 100) {
                                                v121 := GenS10{F0: 42, F1: fi110, F2: 634}
                                                if true {
                                                    v122 := []int{fi12}
                                                    var v123 uint64 = 636
                                                    _ = v122
                                                    _ = v123
                                                }
                                                _ = v121
                                            }
                                        }
                                    }
                                case "a":
                                    v124 := []int{(-(fi12)), 127, v3, (fi12 / (((fi12 / ((v3) | 1))) | 1))}
                                    _ = v124
                                default:
                                    fmt.Println(v3, fi12)
                                }
                                _ = v115
                            } else {
                                var v125 int16 = (v5 - v5)
                                if true {
                                    v4 -= (v13 * 1e10)
                                    var v126 float64 = 0.0
                                    v127 := GenS11{F0: true, F1: 328}
                                    _ = v126
                                    _ = v127
                                }
                                _ = v125
                            }
                        }
                        var v128 float32 = 0.0
                        if true {
                            v3 = (fi110 - (v3 % ((fi12) | 1)))
                            var v129 int8 = 1
                            v130 := GenS6{F0: 128, F1: (fi12 >= 1), F2: (v3 + 745)}
                            _ = v129
                            _ = v130
                        } else {
                            fmt.Println(fi12, v128, fi12)
                            v131 := []int{128}
                            _ = v131
                        }
                        _ = v128
                    }
                    if true {
                        switch "a" {
                        case "":
                            if true {
                                v132 := GenS3{F0: "a"}
                                var v133 int32 = 42
                                _ = v132
                                _ = v133
                            }
                        case "\n\t":
                            delete(v2, 2)
                        default:
                            fmt.Println(v13, v5)
                            v13 -= 1e10
                            v134 := []int{fi12, v3}
                            _ = v134
                        }
                        if false {
                            if true {
                                switch v3 {
                                case -664:
                                    var v135 uint8 = 42
                                    _ = v135
                                case 2147483647:
                                    v136 := []int{(fi12 & fi12), (-(fi110)), fi110, fi12}
                                    _ = v136
                                default:
                                    v137 := GenS12{F0: true, F1: 0}
                                    _ = v137
                                }
                                v1[42] = "foo bar"
                            }
                            v13 = v13
                        } else {
                            var v138 uint8 = 42
                            switch fi12 {
                            case 9223372036854775807:
                                switch 459 {
                                case 1:
                                    v139 := GenS10{F0: 89, F1: fi110, F2: 666}
                                    _ = v139
                                case -324:
                                    var v140 uint32 = 4294967295
                                    _ = v140
                                case 0:
                                    v141 := []int{(fi110 | -607)}
                                    _ = v141
                                }
                            default:
                                v142 := GenS4{F0: 948, F1: 42, F2: (100 != -128)}
                                _ = v142
                            }
                            _ = v138
                        }
                        var v143 uint8 = 127
                        v4 = float32(0.0)
                        v144 := GenS13{F0: 4294967295, F1: 39, F2: 1}
                        _ = v143
                        _ = v144
                    }
                    _ = fi110
                }
            }
        case "hello":
            fmt.Println(v13)
        case "a":
            v145 := GenS14{F0: "\n\t"}
            _ = v145
        default:
            if false {
                v13 = v13
            } else {
                if true {
                    v146 := []int{(fi12 ^ (v3 + v3))}
                    if (!true) {
                        if (2 < 187) {
                            if false {
                                if (-32768 <= v5) {
                                    var v147 uint16 = 32767
                                    _ = v147
                                } else {
                                    v3 = 9223372036854775807
                                }
                                switch "foo bar" {
                                case "hello":
                                    for fi148 := 0; fi148 < 4; fi148++ {
                                        if true {
                                            fmt.Println(v13, v4)
                                            if false {
                                                var v149 uint8 = 0
                                                var v150 float64 = 1e10
                                                if len(v146) >= 2 {
                                                    v146 = v146[1:len(v146)]
                                                }
                                                var v151 int64 = 1
                                                _ = v149
                                                _ = v150
                                                _ = v151
                                            }
                                            var v152 string = "a"
                                            fmt.Println(v13, v5, v5)
                                            _ = v152
                                        }
                                        v146 = append(v146, (fi148 * (fi12 - -2147483648)))
                                        var v153 bool = false
                                        _ = fi148
                                        _ = v153
                                    }
                                default:
                                    _ = len(v2)
                                }
                                v154 := GenS7{F0: 18446744073709551615, F1: 276, F2: "\n\t"}
                                _ = v154
                            } else {
                                if true {
                                    for fi155 := 0; fi155 < 15; fi155++ {
                                        v5 += v5
                                        var v156 bool = false
                                        var v157 float32 = v13
                                        _ = fi155
                                        _ = v156
                                        _ = v157
                                    }
                                    for fi158 := 0; fi158 < 17; fi158++ {
                                        if (true && true) {
                                            var v159 int16 = v5
                                            _ = v159
                                        } else {
                                            switch 1 {
                                            case -2147483648:
                                                if true {
                                                    switch "a" {
                                                    case "":
                                                        var v160 bool = (!true)
                                                        _ = v160
                                                    case "a":
                                                        switch (fi12 & (^(v3))) {
                                                        case 128:
                                                            switch (-(v3)) {
                                                            case 42:
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v146 = append(v146, (v3 & (v3 & 42)))
                                                                case "\n\t":
                                                                    var v161 string = ("\n\t" + "foo bar")
                                                                    _ = v161
                                                                default:
                                                                    var v162 float32 = v4
                                                                    var v163 float32 = v13
                                                                    v164 := GenS14{F0: ("\n\t" + "a")}
                                                                    switch (^(fi158)) {
                                                                    case 0:
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v162 -= -2.25
                                                                        default:
                                                                            fmt.Println(fi12, v3)
                                                                        }
                                                                    case 2147483647:
                                                                        var v165 int16 = v5
                                                                        _ = v165
                                                                    case 4294967295:
                                                                        v4 = float32(-2.25)
                                                                    }
                                                                    _ = v162
                                                                    _ = v163
                                                                    _ = v164
                                                                }
                                                            }
                                                        default:
                                                            v166 := GenS11{F0: (0 > 353), F1: -106}
                                                            _ = v166
                                                        }
                                                    default:
                                                        v167 := GenS15{F0: 82, F1: 42}
                                                        _ = v167
                                                    }
                                                    v4 = v13
                                                    fmt.Println(v13, fi158, v5)
                                                    switch fi158 {
                                                    case -837:
                                                        v168 := GenS16{F0: ("foo bar" + ""), F1: 2}
                                                        _ = v168
                                                    default:
                                                        var v169 float64 = 1.5
                                                        _ = v169
                                                    }
                                                }
                                            case 782:
                                                v170 := GenS17{F0: -1}
                                                _ = v170
                                            }
                                            if len(v146) >= 2 {
                                                v146 = v146[1:len(v146)]
                                            }
                                        }
                                        _ = fi158
                                    }
                                } else {
                                    v171 := v2[2147483648]
                                    if v171 {
                                        var v172 int32 = -1
                                        if v171 {
                                            switch fi12 {
                                            case -1:
                                                v171 = v171
                                            case 0:
                                                var v173 int32 = v172
                                                _ = v173
                                            case 1:
                                                var v174 bool = ((true && true) || ((9223372036854775807 == 42) || (!((-674 < 1) && v171))))
                                                _ = v174
                                            }
                                        }
                                        _ = v172
                                    } else {
                                        var v175 uint8 = 255
                                        v4 = v4
                                        if v171 {
                                            v176 := GenS7{F0: 4294967295, F1: 42, F2: "hello"}
                                            v146 = append(v146, (fi12 / ((fi12) | 1)))
                                            _ = v176
                                        }
                                        _ = v175
                                    }
                                    _ = v171
                                }
                                v3 = 32767
                            }
                            for fi177 := 0; fi177 < 4; fi177++ {
                                fmt.Println(fi177)
                                _ = fi177
                            }
                        } else {
                            var v178 int16 = 32767
                            _ = v178
                        }
                    }
                    v179 := GenS6{F0: 86, F1: true, F2: v3}
                    _ = v146
                    _ = v179
                } else {
                    var v180 uint = 65535
                    var v181 bool = ((!(4294967295 > 0)) || false)
                    _ = v180
                    _ = v181
                }
                var v182 uint16 = 255
                _ = v182
            }
        }
        _ = fi12
        _ = v13
    }
    if false {
        v183 := GenS4{F0: 4294967295, F1: 100, F2: false}
        var v184 int8 = 100
        _ = v183
        _ = v184
    }
    if (-2147483648 >= -1) {
        if true {
            if (!(!(true && true))) {
                var v185 float64 = float64((1e10 + 0.0))
                _ = v185
            } else {
                var v186 int8 = 1
                _ = v186
            }
            v187 := GenS5{F0: false}
            _ = v187
        } else {
            for fi188 := 0; fi188 < 4; fi188++ {
                v3 = (v3 ^ (fi188 % ((9223372036854775807) | 1)))
                var v189 bool = ((-(v5)) >= (-(v5)))
                _ = fi188
                _ = v189
            }
        }
        if (1 < 1) {
            if false {
                var v190 uint8 = 14
                _ = v190
            }
        } else {
            v3 += v3
        }
    }
    if false {
        v5 -= 934
        v191 := v1[2147483647]
        var v192 int = v3
        _ = v191
        _ = v192
    } else {
        v193 := GenS4{F0: 2, F1: 0, F2: ((!true) || true)}
        for fi194 := 0; fi194 < 18; fi194++ {
            var v195 int = fi194
            for fi196 := 0; fi196 < 7; fi196++ {
                var v197 bool = (0 > 42)
                _ = fi196
                _ = v197
            }
            _ = fi194
            _ = v195
        }
        v3 = 264
        _ = v193
    }
    v5 += (^(v5))
    var v198 int = v3
    fmt.Println(v198, v3, v4)
    var v199 uint8 = 0
    v1[42] = ("\n\t" + ("\n\t" + "\n\t"))
    switch "a" {
    case "hello":
        var v200 int32 = 1
        _ = v200
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v198:", v198)
    fmt.Println("v199:", v199)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v198
    _ = v199
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
