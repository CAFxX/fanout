package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int16; F2 uint16 }
type GenS6 struct { F0 int32; F1 int; F2 uint32 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint8; F1 int }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int32; F1 uint64 }
type GenS16 struct { F0 uint64; F1 bool; F2 int32 }
type GenS17 struct { F0 string; F1 bool; F2 uint }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint64; F1 int8; F2 int64 }
type GenS20 struct { F0 uint32; F1 int8 }
type GenS21 struct { F0 uint; F1 uint16 }
func genFn0(p0 int64) int {
    switch "hello" {
    case "a":
        if true {
            switch "hello" {
            case "\n\t":
                var v7 uint32 = 65535
                _ = v7
            case "":
                p0 = 9223372036854775807
            }
            var v8 float64 = 3.14159
            fmt.Println(p0)
            _ = v8
        } else {
            var v9 uint = 0
            p0 = 9223372036854775807
            v10 := []int{128, 32768, -60, -446}
            _ = v10
            _ = v9
        }
    default:
        p0 = (p0 ^ (p0 << ((18446744073709551615) & 63)))
    }
    v11 := []int{4294967295, -2147483648, 2147483647}
    if true {
        switch 127 {
        case 128:
            if len(v11) > 0 {
                v11[0] = 32768
            }
        case -1:
            var v12 float32 = (0.1 - (((-2.25 / 1.5) - (float32(0.0) * 0.0)) - 0.0))
            _ = v12
        case 2147483647:
            var v13 string = "hello"
            _ = v13
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            switch "\n\t" {
            case "foo bar":
                v11 = append(v11, (fi14 >> ((1) & 63)))
            case "\n\t":
                var v15 int64 = p0
                _ = v15
            case "":
                p0 -= p0
            }
            if (0.0 == 0.0) {
                var v16 int16 = 1000
                var v17 bool = (0 < 65535)
                _ = v16
                _ = v17
            } else {
                p0 = p0
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 uint = 540
                _ = fi18
                _ = v19
            }
            _ = fi14
        }
    }
    if (1.5 != (0.1 + 1.5)) {
        var v20 float32 = 0.0
        switch ("foo bar" + "a") {
        case "\n\t":
            var v21 uint = 4294967295
            _ = v21
        case "foo bar":
            var v22 uint64 = 0
            _ = v22
        }
        _ = v20
    }
    _ = p0
    _ = v11
    return 32767
}
func genFn1(p0 int64, p1 bool) uint8 {
    var v23 int64 = (p0 ^ -211)
    p1 = (false || p1)
    if p1 {
        switch ("hello" + "") {
        case "hello":
            v23 -= -1
        }
        if p1 {
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 int32 = -2147483648
                v26 := []int{fi24, fi24, 42, 42}
                v23 += v23
                _ = fi24
                _ = v25
                _ = v26
            }
            switch ("" + (("foo bar" + ((("foo bar" + ("foo bar" + "")) + ("\n\t" + "hello")) + "foo bar")) + "")) {
            case "hello":
                var v27 int = 625
                _ = v27
            case "":
                if false {
                    var v28 string = "a"
                    _ = v28
                } else {
                    var v29 uint = 1
                    if ((!p1) || false) {
                        p1 = p1
                        if ((0.0 / 1.5) != float32((3.14159 / (1e10 / float64(1e10))))) {
                            for fi30 := 0; fi30 < 20; fi30++ {
                                for fi31 := 0; fi31 < 10; fi31++ {
                                    v32 := GenS1{F0: v29, F1: -382, F2: v29}
                                    fmt.Println(fi31, p0)
                                    _ = fi31
                                    _ = v32
                                }
                                var v33 float64 = -2.25
                                var v34 bool = ((v29 << ((v29) & 63)) != 2)
                                _ = fi30
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v23, p1)
                            v35 := GenS1{F0: (^(v29)), F1: -32768, F2: (^(v29))}
                            _ = v35
                        } else {
                            v29 -= 18446744073709551615
                            var v36 int8 = 70
                            _ = v36
                        }
                    } else {
                        fmt.Println(v23, p1)
                        v37 := GenS1{F0: (^(v29)), F1: 42, F2: (v29 + (v29 ^ 255))}
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 3; fi38++ {
                        var v39 int = -2147483648
                        v40 := GenS2{F0: 0, F1: 1}
                        for fi41 := 0; fi41 < 2; fi41++ {
                            if true {
                                v42 := GenS3{F0: 4294967295}
                                var v43 bool = (p1 && p1)
                                _ = v42
                                _ = v43
                            }
                            if true {
                                v44 := []int{255, genFn0(-661), fi41, genFn0(p0)}
                                v45 := GenS3{F0: 4294967295}
                                var v46 string = "a"
                                var v47 uint = 42
                                _ = len(v44)
                                v23 -= (p0 % ((702) | 1))
                                _ = v44
                                _ = v45
                                _ = v46
                                _ = v47
                            } else {
                                v48 := GenS2{F0: 2147483647, F1: 32767}
                                p0 -= p0
                                _ = v48
                            }
                            v39 -= genFn0(1)
                            _ = fi41
                        }
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = v29
                }
            case "a":
                v49 := []int{32767, 128, 127, 127}
                _ = v49
            default:
                v50 := []int{9223372036854775807, 2147483647, 575, 324}
                _ = v50
            }
        }
    }
    v23 = -691
    var v51 int8 = -58
    _ = p0
    _ = p1
    _ = v23
    _ = v51
    return 2
}

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]int)
    var v3 uint32 = 888
    var v4 float64 = (3.14159 / 3.14159)
    var v5 uint32 = v3
    var v6 float64 = v4
    var v52 uint16 = 65535
    switch "\n\t" {
    case "foo bar":
        for fi53 := 0; fi53 < 14; fi53++ {
            v54 := GenS4{F0: -507}
            v55 := []int{genFn0(42), (fi53 << ((513) & 63))}
            _ = fi53
            _ = v54
            _ = v55
        }
    case "\n\t":
        _ = len(v2)
    case "":
        var v56 uint8 = 127
        _ = v56
    default:
        v57 := []int{128}
        _ = v57
    }
    switch 127 {
    case 4294967295:
        for fi58 := 0; fi58 < 18; fi58++ {
            v59 := GenS1{F0: 980, F1: 42, F2: 2}
            switch "foo bar" {
            case "\n\t":
                if (v6 <= 1.5) {
                    v60 := []int{fi58}
                    _ = v60
                }
            case "":
                v2[(v5 >> ((478) & 63))] = (^(fi58))
            default:
                v6 = v6
            }
            _ = fi58
            _ = v59
        }
    case -2147483648:
        if true {
            var v61 uint16 = (v52 * (v52 - v52))
            _ = v61
        }
    default:
        if (1 >= 770) {
            v62 := GenS4{F0: 0}
            for fi63 := 0; fi63 < 3; fi63++ {
                v4 += v4
                for fi64 := 0; fi64 < 2; fi64++ {
                    if true {
                        fmt.Println(v5, v6)
                        v6 = v4
                        v65 := []int{(-(fi64)), genFn0(205)}
                        _ = v65
                    }
                    if (false || (!((-1 > -846) || (9223372036854775807 <= 9223372036854775807)))) {
                        var v66 int32 = int32(fi64)
                        switch "hello" {
                        case "foo bar":
                            var v67 bool = true
                            _ = v67
                        }
                        switch 4294967295 {
                        case -1:
                            v68 := []int{-2147483648, fi63, (^(fi63))}
                            _ = v68
                        default:
                            switch (("foo bar" + (("foo bar" + "hello") + (("foo bar" + "hello") + ""))) + "hello") {
                            case "":
                                if (12 != 2) {
                                    fmt.Println(v1)
                                    var v69 int8 = v1
                                    v52 = (^(v52))
                                    _ = v69
                                }
                            case "foo bar":
                                if false {
                                    switch "foo bar" {
                                    case "a":
                                        _ = len(v2)
                                    case "foo bar":
                                        v66 -= (v66 % ((v66) | 1))
                                    case "hello":
                                        v66 = v66
                                    }
                                    if (2147483648 <= 838) {
                                        var v70 uint = 1
                                        _ = v70
                                    } else {
                                        var v71 int = (fi64 << ((893) & 63))
                                        _ = v71
                                    }
                                    if true {
                                        v3 += v3
                                        v72 := GenS3{F0: (v3 / (((^(v5))) | 1))}
                                        var v73 int64 = 42
                                        v52 = v52
                                        _ = v72
                                        _ = v73
                                    }
                                }
                            case "a":
                                if (!true) {
                                    v74 := []int{genFn0(int64(fi63))}
                                    fmt.Println(v52, v66, v4)
                                    _ = v74
                                } else {
                                    v4 = v4
                                    fmt.Println(v5, v66, v52)
                                    var v75 int32 = (v66 * (v66 / ((int32(fi64)) | 1)))
                                    fmt.Println(v6)
                                    v76 := GenS4{F0: 864}
                                    _ = v75
                                    _ = v76
                                }
                            default:
                                if (!true) {
                                    v77 := []int{fi63, fi63, (fi63 ^ fi63), fi64}
                                    if (!(((float64(0.1) / v4) > (-2.25 / v6)) || (((((float64((1.5 / 0.1)) - v6) - (v6 - v6)) - v6) - 0.1) == v6))) {
                                        v78 := GenS5{F0: (v66 * (-(v66))), F1: 32767, F2: v52}
                                        if false {
                                            var v79 uint8 = 255
                                            if (!(!true)) {
                                                if false {
                                                    v5 += uint32(v79)
                                                    fmt.Println(v6, v6)
                                                    v80 := v2[308]
                                                    _ = v80
                                                }
                                                v66 = -746
                                                if (-668 >= 0) {
                                                    v77 = append(v77, genFn0(-9223372036854775808))
                                                } else {
                                                    v6 = -2.25
                                                    var v81 int8 = v1
                                                    switch ("\n\t" + "a") {
                                                    case "":
                                                        var v82 uint16 = (^(v52))
                                                        _ = v82
                                                    default:
                                                        v4 = -2.25
                                                    }
                                                    _ = v81
                                                }
                                                _ = len(v77)
                                                if true {
                                                    switch "hello" {
                                                    case "hello":
                                                        if (!(!true)) {
                                                            fmt.Println(v6, v3)
                                                            v83 := GenS6{F0: v66, F1: (fi64 >> ((270) & 63)), F2: (v3 << ((129) & 63))}
                                                            _ = v83
                                                        } else {
                                                            v5 -= 1
                                                            if true {
                                                                var v84 uint64 = 854
                                                                v85 := v2[846]
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                v52 = v52
                                                                fmt.Println(v52, v6, v66)
                                                                if false {
                                                                    var v86 uint32 = v3
                                                                    _ = v86
                                                                } else {
                                                                    if (("foo bar" <= "") && (v52 > (v52 % (((^(v52))) | 1)))) {
                                                                        switch fi63 {
                                                                        case -2147483648:
                                                                            var v87 int16 = -32768
                                                                            v87 = v87
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                v79 -= (v79 + v79)
                                                                                _ = len(v2)
                                                                                v3 -= v5
                                                                            case "foo bar":
                                                                                fmt.Println(v3, v4)
                                                                            default:
                                                                                v88 := GenS6{F0: (v66 ^ int32(v79)), F1: 1, F2: v3}
                                                                                _ = v88
                                                                            }
                                                                            _ = v87
                                                                        case 32768:
                                                                            v89 := GenS7{F0: (v79 >> ((4294967295) & 63)), F1: 430}
                                                                            _ = v89
                                                                        case -856:
                                                                            var v90 uint64 = 168
                                                                            _ = v90
                                                                        default:
                                                                            if len(v77) > 0 {
                                                                                v77[0] = (fi64 / (((fi63 & genFn0(0))) | 1))
                                                                            }
                                                                        }
                                                                        if ("\n\t" >= ("" + (("\n\t" + "\n\t") + "hello"))) {
                                                                            switch (fi63 / ((fi64) | 1)) {
                                                                            case -348:
                                                                                switch (((("" + "a") + "") + "hello") + "a") {
                                                                                case "\n\t":
                                                                                    delete(v2, (v3 | (v3 >> ((544) & 63))))
                                                                                }
                                                                            default:
                                                                                switch (fi64 & genFn0(1)) {
                                                                                case -2147483648:
                                                                                    switch (fi64 * (fi64 & (fi64 * 4294967295))) {
                                                                                    case 13:
                                                                                        fmt.Println(v1, v52)
                                                                                    case 9223372036854775807:
                                                                                        v3 += v5
                                                                                    case -9223372036854775808:
                                                                                        v1 = v1
                                                                                    default:
                                                                                        fmt.Println(fi63)
                                                                                    }
                                                                                case 32767:
                                                                                    var v91 string = ""
                                                                                    _ = v91
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v92 := GenS5{F0: 0, F1: 1000, F2: (v52 & 270)}
                                                                    fmt.Println(v1, v52, v6)
                                                                    _ = v92
                                                                }
                                                            }
                                                        }
                                                        delete(v2, uint32(v79))
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            switch "" {
                                                            case "a":
                                                                switch fi63 {
                                                                case 42:
                                                                    v66 = v66
                                                                default:
                                                                    delete(v2, v5)
                                                                }
                                                            case "\n\t":
                                                                var v93 int16 = 32767
                                                                _ = v93
                                                            default:
                                                                v79 -= (v79 / ((127) | 1))
                                                            }
                                                        case "foo bar":
                                                            _ = len(v77)
                                                        case "hello":
                                                            v79 = (v79 ^ v79)
                                                        }
                                                        fmt.Println(v3)
                                                        var v94 float32 = 0.1
                                                        if (v1 != v1) {
                                                            var v95 uint = 65535
                                                            _ = v95
                                                        }
                                                        _ = v94
                                                    case "a":
                                                        var v96 uint16 = (v52 >> ((uint(v5)) & 63))
                                                        _ = v96
                                                    case "\n\t":
                                                        switch "a" {
                                                        case "":
                                                            fmt.Println(v4, v6)
                                                        case "\n\t":
                                                            if false {
                                                                fmt.Println(v5, v3, v79)
                                                                if len(v77) > 0 {
                                                                    v77[0] = -76
                                                                }
                                                            }
                                                        }
                                                    default:
                                                        if false {
                                                            var v97 int32 = (v66 % (((v66 / ((v66) | 1))) | 1))
                                                            var v98 int8 = 23
                                                            _ = v97
                                                            _ = v98
                                                        } else {
                                                            var v99 uint8 = v79
                                                            v79 -= 185
                                                            _ = v99
                                                        }
                                                    }
                                                    if (0 < 1) {
                                                        v1 -= int8(v79)
                                                        fmt.Println(fi63, v79)
                                                        if false {
                                                            var v100 uint = 834
                                                            fmt.Println(v6)
                                                            if (221 != 167) {
                                                                fmt.Println(v100, v4, v52)
                                                                var v101 float64 = 0.1
                                                                _ = v101
                                                            }
                                                            _ = v100
                                                        }
                                                    }
                                                    if (!true) {
                                                        switch (fi64 >> ((589) & 63)) {
                                                        case 223:
                                                            switch "\n\t" {
                                                            case "a":
                                                                fmt.Println(v5, v6)
                                                            case "foo bar":
                                                                v5 = 1
                                                                switch (^(fi64)) {
                                                                case 190:
                                                                    switch (^(fi63)) {
                                                                    case 988:
                                                                        if ("a" <= "foo bar") {
                                                                            var v102 int16 = 42
                                                                            if true {
                                                                                var v103 int64 = 0
                                                                                v104 := v2[v3]
                                                                                var v105 int16 = (v102 - v102)
                                                                                _ = v103
                                                                                _ = v104
                                                                                _ = v105
                                                                            } else {
                                                                                if false {
                                                                                    switch 42 {
                                                                                    case 0:
                                                                                        if (127 < fi64) {
                                                                                            fmt.Println(v102)
                                                                                            _ = len(v77)
                                                                                            if (v102 <= (v102 << ((1) & 63))) {
                                                                                                var v106 bool = (true && ((^(v79)) > v79))
                                                                                                fmt.Println(v1, v4)
                                                                                                v107 := GenS2{F0: (v3 - (v3 & v3)), F1: (v52 | v52)}
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            } else {
                                                                                                if (int16(v52) == v102) {
                                                                                                    var v108 bool = false
                                                                                                    v1 = (v1 * 100)
                                                                                                    _ = v108
                                                                                                }
                                                                                                var v109 int = (fi63 | int(v3))
                                                                                                v110 := GenS8{F0: (float64(0.0) >= float64(0.1)), F1: (v3 / ((v5) | 1))}
                                                                                                _ = v109
                                                                                                _ = v110
                                                                                            }
                                                                                            v3 -= (v5 & 0)
                                                                                            v111 := GenS9{F0: 774}
                                                                                            _ = v111
                                                                                        } else {
                                                                                            var v112 uint16 = 483
                                                                                            var v113 int8 = v1
                                                                                            switch fi64 {
                                                                                            case 42:
                                                                                                v1 += v113
                                                                                            case 32768:
                                                                                                delete(v2, 725)
                                                                                            case -9223372036854775808:
                                                                                                if true {
                                                                                                    var v114 uint8 = v79
                                                                                                    _ = v114
                                                                                                } else {
                                                                                                    v112 = v112
                                                                                                    v115 := GenS8{F0: ((v79 << ((465) & 63)) <= v79), F1: 987}
                                                                                                    _ = v115
                                                                                                }
                                                                                            }
                                                                                            switch ("hello" + (("hello" + "\n\t") + "a")) {
                                                                                            case "hello":
                                                                                                if true {
                                                                                                    v116 := GenS3{F0: (v5 << ((891) & 63))}
                                                                                                    var v117 string = ("\n\t" + ("\n\t" + ""))
                                                                                                    _ = v116
                                                                                                    _ = v117
                                                                                                } else {
                                                                                                    if (!false) {
                                                                                                        v118 := GenS8{F0: ((v66 << ((18446744073709551615) & 63)) != v66), F1: v3}
                                                                                                        _ = v118
                                                                                                    } else {
                                                                                                        v77 = append(v77, -42)
                                                                                                    }
                                                                                                    var v119 uint = 384
                                                                                                    var v120 int8 = -1
                                                                                                    _ = v119
                                                                                                    _ = v120
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v121 uint32 = (v3 + (v3 | 2))
                                                                                                _ = v121
                                                                                            }
                                                                                            _ = v112
                                                                                            _ = v113
                                                                                        }
                                                                                    default:
                                                                                        v6 = v6
                                                                                        if true {
                                                                                            switch "a" {
                                                                                            case "foo bar":
                                                                                                switch "foo bar" {
                                                                                                case "hello":
                                                                                                    switch genFn0(-1) {
                                                                                                    case 127:
                                                                                                        var v122 uint32 = 47
                                                                                                        _ = v122
                                                                                                    case 670:
                                                                                                        v123 := GenS1{F0: 2, F1: 1000, F2: 759}
                                                                                                        var v124 int = 2147483647
                                                                                                        v125 := GenS6{F0: (v66 + (^(v66))), F1: genFn0(0), F2: 2}
                                                                                                        if (!true) {
                                                                                                            v5 -= (v3 << ((4294967295) & 63))
                                                                                                            switch ("\n\t" + "") {
                                                                                                            case "\n\t":
                                                                                                                _ = len(v2)
                                                                                                            case "a":
                                                                                                                var v126 uint64 = 100
                                                                                                                _ = v126
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    if len(v77) > 0 {
                                                                                                                        v77[0] = int(v1)
                                                                                                                    }
                                                                                                                    var v127 int8 = (v1 & 42)
                                                                                                                    _ = v127
                                                                                                                } else {
                                                                                                                    if false {
                                                                                                                        var v128 string = "hello"
                                                                                                                        v129 := GenS10{F0: (v102 / (((-(v102))) | 1))}
                                                                                                                        _ = v128
                                                                                                                        _ = v129
                                                                                                                    } else {
                                                                                                                        v6 += (v4 / 0.1)
                                                                                                                        v130 := GenS7{F0: (v79 >> ((65535) & 63)), F1: v66}
                                                                                                                        v102 += -32768
                                                                                                                        _ = v130
                                                                                                                    }
                                                                                                                    v131 := GenS11{F0: v102}
                                                                                                                    fmt.Println(v52)
                                                                                                                    _ = v131
                                                                                                                }
                                                                                                            }
                                                                                                            v124 = 0
                                                                                                        }
                                                                                                        if (v4 <= 1e10) {
                                                                                                            v132 := GenS12{F0: (2 == 1)}
                                                                                                            fmt.Println(v102, v1)
                                                                                                            _ = v132
                                                                                                        }
                                                                                                        _ = v123
                                                                                                        _ = v124
                                                                                                        _ = v125
                                                                                                    case 33:
                                                                                                        if len(v77) > 0 {
                                                                                                            v77[0] = fi64
                                                                                                        }
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v5 += 0
                                                                                                default:
                                                                                                    switch (fi63 >> ((0) & 63)) {
                                                                                                    case 9223372036854775807:
                                                                                                        v1 = (v1 | (v1 * v1))
                                                                                                    case 0:
                                                                                                        fmt.Println(v52, v3)
                                                                                                    }
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v3 += 208
                                                                                            case "a":
                                                                                                var v133 uint32 = v3
                                                                                                _ = v133
                                                                                            }
                                                                                            if false {
                                                                                                var v134 float64 = v4
                                                                                                if true {
                                                                                                    fmt.Println(fi63, v4)
                                                                                                    fmt.Println(v1)
                                                                                                } else {
                                                                                                    var v135 uint16 = (v52 >> ((1) & 63))
                                                                                                    _ = v135
                                                                                                }
                                                                                                switch genFn0(0) {
                                                                                                case 4294967295:
                                                                                                    v136 := GenS4{F0: 1}
                                                                                                    v137 := GenS11{F0: v102}
                                                                                                    _ = v136
                                                                                                    _ = v137
                                                                                                case -916:
                                                                                                    if len(v77) > 0 {
                                                                                                        v77[0] = (fi64 & fi64)
                                                                                                    }
                                                                                                case 128:
                                                                                                    if (650 != 255) {
                                                                                                        v4 = 0.0
                                                                                                        v138 := GenS9{F0: 42}
                                                                                                        v79 += genFn1(-692, false)
                                                                                                        _ = v138
                                                                                                    }
                                                                                                }
                                                                                                v139 := GenS13{F0: genFn1(1, true), F1: (fi63 >> ((628) & 63))}
                                                                                                _ = v134
                                                                                                _ = v139
                                                                                            } else {
                                                                                                var v140 float32 = -2.25
                                                                                                if true {
                                                                                                    if (((v66 | (-(v66))) != (v66 << ((1) & 63))) || true) {
                                                                                                        v79 = v79
                                                                                                        var v141 int64 = 820
                                                                                                        var v142 int64 = (v141 % ((v141) | 1))
                                                                                                        _ = v141
                                                                                                        _ = v142
                                                                                                    }
                                                                                                    v79 -= genFn1(9223372036854775807, ((-(fi63)) != (-(fi64))))
                                                                                                    var v143 int = genFn0(-9223372036854775808)
                                                                                                    _ = v143
                                                                                                } else {
                                                                                                    var v144 float64 = (v6 / -2.25)
                                                                                                    if len(v77) >= 2 {
                                                                                                        v77 = v77[1:len(v77)]
                                                                                                    }
                                                                                                    _ = v144
                                                                                                }
                                                                                                _ = v140
                                                                                            }
                                                                                        } else {
                                                                                            v145 := GenS1{F0: 805, F1: 0, F2: 4294967295}
                                                                                            _ = v145
                                                                                        }
                                                                                        v6 += (float64(0.0) - (1.5 * v4))
                                                                                        var v146 string = ("\n\t" + "hello")
                                                                                        _ = v146
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    fmt.Println(v4, v1, v79)
                                                                                    v77 = append(v77, 0)
                                                                                } else {
                                                                                    if false {
                                                                                        var v147 bool = ((v5 << ((18446744073709551615) & 63)) != v3)
                                                                                        if (!false) {
                                                                                            v6 -= v4
                                                                                            var v148 float64 = v6
                                                                                            _ = v148
                                                                                        } else {
                                                                                            v4 += v6
                                                                                            var v149 uint = 2
                                                                                            switch (("hello" + "") + "foo bar") {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v1, v1)
                                                                                            case "foo bar":
                                                                                                v5 += v3
                                                                                            case "":
                                                                                                var v150 uint8 = uint8(v52)
                                                                                                _ = v150
                                                                                            default:
                                                                                                v151 := v2[596]
                                                                                                _ = v151
                                                                                            }
                                                                                            _ = v149
                                                                                        }
                                                                                        var v152 float32 = 1.5
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            var v153 int = -2147483648
                                                                                            _ = v153
                                                                                        case "\n\t":
                                                                                            v154 := GenS1{F0: 42, F1: v102, F2: 18446744073709551615}
                                                                                            _ = v154
                                                                                        case "hello":
                                                                                            var v155 uint32 = (v3 * 4294967295)
                                                                                            _ = v155
                                                                                        default:
                                                                                            var v156 int16 = -278
                                                                                            _ = v156
                                                                                        }
                                                                                        v4 = -2.25
                                                                                        fmt.Println(v152)
                                                                                        _ = v147
                                                                                        _ = v152
                                                                                    }
                                                                                }
                                                                                v4 += float64(0.0)
                                                                            }
                                                                            if ((genFn0(80) != fi63) || ((!true) || (v52 > 906))) {
                                                                                switch fi63 {
                                                                                case 234:
                                                                                    if (genFn0(int64(v102)) > 1) {
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            v157 := GenS14{F0: 1}
                                                                                            _ = v157
                                                                                        }
                                                                                        v5 = (v3 + v5)
                                                                                    } else {
                                                                                        var v158 bool = true
                                                                                        _ = v158
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        if true {
                                                                                            v52 += (v52 * v52)
                                                                                            var v159 int32 = (-(v66))
                                                                                            if (v66 >= (v159 & (v66 << ((20) & 63)))) {
                                                                                                v3 = v3
                                                                                                var v160 int64 = -408
                                                                                                _ = v160
                                                                                            } else {
                                                                                                if ((v102 != -32768) && false) {
                                                                                                    if ((v102 - 0) == v102) {
                                                                                                        v161 := GenS15{F0: (v159 % ((int32(v3)) | 1)), F1: 42}
                                                                                                        switch "hello" {
                                                                                                        case "a":
                                                                                                            v52 += (v52 | 32767)
                                                                                                        case "":
                                                                                                            fmt.Println(v4, v3)
                                                                                                        default:
                                                                                                            v162 := GenS12{F0: (uint8(v52) >= (v79 * (^(v79))))}
                                                                                                            _ = v162
                                                                                                        }
                                                                                                        _ = v161
                                                                                                    } else {
                                                                                                        v163 := v2[(^(v3))]
                                                                                                        v159 = v159
                                                                                                        v163 = genFn0(1)
                                                                                                        _ = v163
                                                                                                    }
                                                                                                    fmt.Println(fi64, v3, fi64)
                                                                                                    var v164 int8 = (-(v1))
                                                                                                    _ = v164
                                                                                                }
                                                                                                v6 = v4
                                                                                                v79 = 255
                                                                                                var v165 int = (fi64 - fi64)
                                                                                                _ = v165
                                                                                            }
                                                                                            v166 := GenS16{F0: 42, F1: (255 > 0), F2: (-(v159))}
                                                                                            _ = v159
                                                                                            _ = v166
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v6)
                                                                                var v167 bool = (true && false)
                                                                                _ = v167
                                                                            }
                                                                            _ = v102
                                                                        }
                                                                    case -2147483648:
                                                                        var v168 float32 = -2.25
                                                                        _ = v168
                                                                    case 42:
                                                                        fmt.Println(v66, v6, v3)
                                                                    default:
                                                                        var v169 uint8 = v79
                                                                        _ = v169
                                                                    }
                                                                case 948:
                                                                    if (18446744073709551615 < 107) {
                                                                        v1 -= (v1 ^ 19)
                                                                    }
                                                                case 1:
                                                                    v2[(^(v3))] = (fi64 * fi64)
                                                                default:
                                                                    if true {
                                                                        var v170 float64 = ((0.1 + 0.0) + ((0.1 / v4) * 0.0))
                                                                        v171 := GenS15{F0: v66, F1: 2}
                                                                        if (!false) {
                                                                            var v172 int8 = v1
                                                                            _ = v172
                                                                        }
                                                                        _ = v170
                                                                        _ = v171
                                                                    }
                                                                    if len(v77) >= 2 {
                                                                        v77 = v77[1:len(v77)]
                                                                    }
                                                                }
                                                                fmt.Println(v79)
                                                            case "":
                                                                v6 += 3.14159
                                                            }
                                                        case 2147483647:
                                                            v173 := GenS17{F0: (("hello" + "foo bar") + "\n\t"), F1: false, F2: 255}
                                                            _ = v173
                                                        case -743:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v174 := v2[uint32(v79)]
                                                                _ = v174
                                                            default:
                                                                if false {
                                                                    v175 := GenS18{F0: ((("foo bar" + "a") + ("" + ("a" + ("\n\t" + ("\n\t" + ("foo bar" + "hello")))))) + ("" + "\n\t"))}
                                                                    _ = v175
                                                                } else {
                                                                    if false {
                                                                        v79 = 128
                                                                        v6 -= (v6 * 0.1)
                                                                        fmt.Println(v52)
                                                                    } else {
                                                                        if false {
                                                                            if len(v77) > 0 {
                                                                                v77[0] = (^(fi64))
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v1, v4, v79)
                                                                            fmt.Println(v6, fi63, v52)
                                                                        }
                                                                        switch ("\n\t" + ("foo bar" + "a")) {
                                                                        case "a":
                                                                            if true {
                                                                                v52 -= (^(v52))
                                                                            }
                                                                        case "\n\t":
                                                                            v2[v5] = genFn0(1)
                                                                        default:
                                                                            var v176 uint64 = 42
                                                                            _ = v176
                                                                        }
                                                                        switch (fi63 + (fi63 >> ((2) & 63))) {
                                                                        case 0:
                                                                            v4 -= (v4 + v4)
                                                                        }
                                                                        _ = len(v77)
                                                                    }
                                                                    v77 = append(v77, (fi63 ^ genFn0(0)))
                                                                }
                                                            }
                                                        default:
                                                            v177 := GenS7{F0: v79, F1: 303}
                                                            _ = v177
                                                        }
                                                    } else {
                                                        if len(v77) > 0 {
                                                            v77[0] = fi64
                                                        }
                                                        v5 += 1
                                                        fmt.Println(v79)
                                                        var v178 uint = 42
                                                        _ = v178
                                                    }
                                                }
                                            } else {
                                                var v179 uint = 4294967295
                                                v3 += v3
                                                _ = v179
                                            }
                                            _ = v79
                                        } else {
                                            delete(v2, v5)
                                        }
                                        _ = v78
                                    } else {
                                        v180 := GenS19{F0: 170, F1: (v1 / ((42) | 1)), F2: 9223372036854775807}
                                        var v181 int64 = -549
                                        var v182 uint = 18446744073709551615
                                        _ = v180
                                        _ = v181
                                        _ = v182
                                    }
                                    if ((true || false) && (((^(v52)) > (v52 * 1)) || true)) {
                                        v6 = (v6 / v4)
                                        v2[v5] = 752
                                    }
                                    _ = v77
                                }
                            }
                        }
                        _ = v66
                    }
                    _ = fi64
                }
                _ = fi63
            }
            _ = v62
        }
    }
    switch 32767 {
    case 9223372036854775807:
        v52 = v52
    case 4294967295:
        switch 128 {
        case -553:
            if true {
                v1 = v1
            }
        default:
            _ = len(v2)
        }
    case -1:
        for fi183 := 0; fi183 < 12; fi183++ {
            var v184 bool = true
            for fi185 := 0; fi185 < 10; fi185++ {
                fmt.Println(v3, v4)
                _ = fi185
            }
            v3 -= 212
            _ = fi183
            _ = v184
        }
    default:
        var v186 uint16 = (v52 % (((v52 - v52)) | 1))
        _ = v186
    }
    for fi187 := 0; fi187 < 9; fi187++ {
        v188 := GenS20{F0: 70, F1: (^(v1))}
        var v189 int16 = -544
        var v190 bool = ((-(v189)) == v189)
        _ = fi187
        _ = v188
        _ = v189
        _ = v190
    }
    for fi191 := 0; fi191 < 9; fi191++ {
        if true {
            v6 = float64(((1.5 / (3.14159 + 3.14159)) / 3.14159))
            v192 := []int{}
            v5 += 42
            _ = v192
        }
        _ = fi191
    }
    for fi193 := 0; fi193 < 2; fi193++ {
        for fi194 := 0; fi194 < 13; fi194++ {
            fmt.Println(v1, v1)
            delete(v2, (v3 * 569))
            var v195 float32 = (1.5 - ((3.14159 / (-2.25 + (1.5 - ((0.0 + 0.1) * 0.0)))) * 3.14159))
            _ = fi194
            _ = v195
        }
        v4 += float64((3.14159 * 1.5))
        _ = fi193
    }
    var v196 float64 = (float64(0.0) - v4)
    v2[(v3 << ((573) & 63))] = 4294967295
    v1 += (v1 ^ v1)
    v197 := GenS21{F0: 255, F1: v52}
    fmt.Println("v1:", v1)
    fmt.Println("v196:", v196)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v196
    _ = v197
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
