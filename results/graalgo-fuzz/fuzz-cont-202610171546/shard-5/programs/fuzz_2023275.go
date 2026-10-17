package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 int16; F2 int64 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string; F1 int64 }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint32; F1 int }
type GenS15 struct { F0 int16 }
func genFn0(p0 int8, p1 uint) bool {
    v7 := GenS1{F0: false}
    var v8 float64 = ((3.14159 / 1.0) - float64(3.14159))
    var v9 uint16 = 324
    switch -2147483648 {
    case 42:
        v10 := []int{4294967295}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1() uint64 {
    switch 2147483647 {
    case 32768:
        for fi11 := 0; fi11 < 19; fi11++ {
            switch fi11 {
            case -32:
                var v12 uint8 = 0
                _ = v12
            case 42:
                var v13 string = "\n\t"
                _ = v13
            }
            var v14 int32 = 42
            _ = fi11
            _ = v14
        }
    default:
        v15 := GenS2{F0: true, F1: 97, F2: 2147483647}
        _ = v15
    }
    v16 := []int{-539, -923, -2147483648}
    if ((((3.14159 / (0.0 - 0.1)) - 1.5) - 0.1) >= 1e10) {
        if false {
            switch 9223372036854775807 {
            case 9223372036854775807:
                if ((4294967295 == 0) && false) {
                    var v17 string = ""
                    v17 = "hello"
                    _ = v17
                } else {
                    if true {
                        v18 := GenS3{F0: 42, F1: -32768, F2: 239}
                        var v19 int8 = -128
                        v19 += 127
                        _ = v18
                        _ = v19
                    } else {
                        switch 127 {
                        case 32767:
                            for fi20 := 0; fi20 < 20; fi20++ {
                                var v21 uint = 0
                                v21 -= v21
                                v22 := GenS3{F0: 2, F1: 32767, F2: 9223372036854775807}
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case 1:
                            var v23 uint8 = 128
                            _ = v23
                        case 32768:
                            v24 := GenS3{F0: 0, F1: 1, F2: 0}
                            _ = v24
                        default:
                            var v25 int16 = 42
                            _ = v25
                        }
                        if false {
                            var v26 uint = 0
                            for fi27 := 0; fi27 < 14; fi27++ {
                                switch ("" + (("foo bar" + ("foo bar" + "\n\t")) + (("foo bar" + "") + "foo bar"))) {
                                case "\n\t":
                                    v26 += (v26 % ((v26) | 1))
                                case "foo bar":
                                    v26 -= (v26 >> ((v26) & 63))
                                default:
                                    switch int(v26) {
                                    case 9223372036854775807:
                                        var v28 float32 = 0.0
                                        _ = v28
                                    default:
                                        for fi29 := 0; fi29 < 18; fi29++ {
                                            if (42 >= 384) {
                                                fmt.Println(fi27)
                                            }
                                            _ = fi29
                                        }
                                    }
                                }
                                switch "hello" {
                                case "foo bar":
                                    switch "hello" {
                                    case "a":
                                        var v30 uint8 = 42
                                        _ = v30
                                    case "":
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                    case "hello":
                                        v26 += v26
                                    default:
                                        v31 := GenS3{F0: 2, F1: 1000, F2: 9223372036854775807}
                                        _ = v31
                                    }
                                case "hello":
                                    var v32 uint8 = 141
                                    _ = v32
                                }
                                v26 -= (v26 % ((v26) | 1))
                                _ = fi27
                            }
                            _ = v26
                        } else {
                            var v33 string = "\n\t"
                            v34 := GenS1{F0: true}
                            _ = v33
                            _ = v34
                        }
                        var v35 string = "a"
                        _ = v35
                    }
                    var v36 int16 = 0
                    v16 = append(v16, 42)
                    _ = v36
                }
            }
        }
        v37 := GenS4{F0: "hello", F1: 403}
        _ = v37
    } else {
        for fi38 := 0; fi38 < 13; fi38++ {
            if true {
                v39 := GenS5{F0: 42, F1: 69}
                switch 9223372036854775807 {
                case 1:
                    var v40 bool = (((-32768 < 32767) || false) && (true || (false && (42 == 2))))
                    _ = v40
                case 9223372036854775807:
                    var v41 uint64 = 2
                    _ = v41
                case 2147483647:
                    var v42 int64 = 519
                    _ = v42
                }
                var v43 uint64 = 47
                _ = v39
                _ = v43
            } else {
                v16 = append(v16, -539)
                var v44 int64 = 42
                fmt.Println(fi38)
                _ = v44
            }
            if len(v16) > 0 {
                v16[0] = fi38
            }
            _ = fi38
        }
        var v45 uint64 = 18446744073709551615
        switch -138 {
        case 42:
            var v46 uint64 = (^(v45))
            _ = v46
        default:
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        _ = v45
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        var v48 int64 = -9223372036854775808
        _ = fi47
        _ = v48
    }
    var v49 uint64 = 189
    _ = v16
    _ = v49
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2147483648, 0}
    var v3 int8 = 109
    v4 := []int{}
    v5 := make(map[string]bool)
    v6 := []string{"foo bar", "hello", ("a" + "hello"), ("" + "hello")}
    switch -2147483648 {
    case 42:
        _ = len(v2)
    }
    switch "hello" {
    case "a":
        for fi50 := 0; fi50 < 11; fi50++ {
            var v51 int32 = -2147483648
            _ = fi50
            _ = v51
        }
    case "\n\t":
        if (0.1 < 1e10) {
            v52 := GenS6{F0: 1, F1: false, F2: 0}
            fmt.Println(v3)
            v53 := GenS2{F0: (4294967295 >= 9223372036854775807), F1: -819, F2: 42}
            _ = v52
            _ = v53
        } else {
            var v54 uint16 = 32767
            fmt.Println(v3)
            switch "\n\t" {
            case "foo bar":
                if ((("\n\t" + "foo bar") == "\n\t") || (100 == (v3 * v3))) {
                    v3 = 42
                    if (-2147483648 <= 2147483647) {
                        v55 := GenS7{F0: genFn1(), F1: false, F2: v54}
                        switch "\n\t" {
                        case "a":
                            v56 := v1["\n\t"]
                            _ = v56
                        }
                        var v57 float64 = (0.1 / 0.1)
                        _ = v55
                        _ = v57
                    }
                    var v58 int64 = 0
                    _ = v58
                } else {
                    for fi59 := 0; fi59 < 12; fi59++ {
                        for fi60 := 0; fi60 < 17; fi60++ {
                            var v61 uint32 = 4294967295
                            var v62 bool = false
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                        v3 = 0
                        _ = fi59
                    }
                    fmt.Println(v3, v3)
                }
            case "":
                v63 := GenS6{F0: 128, F1: (118 < 392), F2: 42}
                _ = v63
            case "a":
                var v64 int32 = -1
                _ = v64
            }
            _ = v54
        }
    case "":
        for fi65 := 0; fi65 < 9; fi65++ {
            var v66 int64 = 1
            if true {
                for fi67 := 0; fi67 < 16; fi67++ {
                    fmt.Println(v66)
                    _ = fi67
                }
            } else {
                var v68 int = fi65
                var v69 uint64 = 203
                v5["a"] = (false || (0.1 > -2.25))
                _ = v68
                _ = v69
            }
            _ = fi65
            _ = v66
        }
    default:
        if false {
            var v70 uint8 = 1
            for fi71 := 0; fi71 < 9; fi71++ {
                switch "a" {
                case "\n\t":
                    v70 = 127
                case "hello":
                    var v72 uint = 964
                    _ = v72
                case "foo bar":
                    v70 += v70
                }
                _ = fi71
            }
            _ = v70
        }
    }
    for fi73 := 0; fi73 < 4; fi73++ {
        if false {
            fmt.Println(fi73)
        } else {
            v3 += (v3 << ((0) & 63))
        }
        _ = len(v1)
        if false {
            var v74 int = fi73
            _ = v74
        }
        _ = fi73
    }
    v3 = (-(v3))
    var v75 uint = 65535
    var v76 int64 = -1
    var v77 bool = (!(true || true))
    var v78 uint16 = 485
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if v77 {
        switch ("" + ("" + "foo bar")) {
        case "hello":
            switch -2147483648 {
            case 42:
                if false {
                    v6 = append(v6, "")
                }
            case 128:
                if v77 {
                    var v79 int32 = -1
                    v80 := GenS1{F0: ((v3 + -1) != (v3 - (v3 & int8(v78))))}
                    v79 = v79
                    _ = v79
                    _ = v80
                }
            case 9223372036854775807:
                for fi81 := 0; fi81 < 4; fi81++ {
                    fmt.Println(v75, fi81, v77)
                    for fi82 := 0; fi82 < 1; fi82++ {
                        if (true && (v75 < (v75 % ((v75) | 1)))) {
                            fmt.Println(v77, v76, v3)
                            _ = len(v1)
                        }
                        _ = fi82
                    }
                    _ = fi81
                }
            }
        case "":
            v3 += v3
        default:
            for fi83 := 0; fi83 < 3; fi83++ {
                var v84 int = fi83
                var v85 uint = v75
                switch (((((("hello" + "foo bar") + ((("hello" + "\n\t") + "hello") + ("\n\t" + "hello"))) + "\n\t") + "foo bar") + ("" + "")) + ((("foo bar" + "a") + "hello") + "a")) {
                case "foo bar":
                    if false {
                        for fi86 := 0; fi86 < 14; fi86++ {
                            if v77 {
                                v2 = append(v2, 994)
                                v75 = 237
                                switch "\n\t" {
                                case "a":
                                    var v87 int = (^(fi83))
                                    _ = v87
                                case "":
                                    v6 = append(v6, "\n\t")
                                case "foo bar":
                                    if (!(v75 == 42)) {
                                        v88 := GenS7{F0: 42, F1: (2 <= 307), F2: v78}
                                        v89 := GenS8{F0: v76}
                                        var v90 float64 = (-2.25 * 0.0)
                                        v91 := GenS9{F0: v76, F1: (false || v77)}
                                        var v92 uint32 = 2147483648
                                        if false {
                                            var v93 uint32 = (v92 << (((v85 / ((65535) | 1))) & 63))
                                            if true {
                                                fmt.Println(fi86, v76)
                                                if false {
                                                    v77 = (42 > 704)
                                                    var v94 int8 = 64
                                                    delete(v1, "")
                                                    _ = v94
                                                }
                                                v3 = 1
                                                fmt.Println(v90, fi86, v76)
                                                if true {
                                                    v95 := v1["hello"]
                                                    var v96 float64 = float64((-2.25 * float32(1.5)))
                                                    _ = v95
                                                    _ = v96
                                                }
                                            }
                                            var v97 int64 = -226
                                            _ = v93
                                            _ = v97
                                        }
                                        _ = v88
                                        _ = v89
                                        _ = v90
                                        _ = v91
                                        _ = v92
                                    }
                                }
                            }
                            switch (("" + ("foo bar" + "a")) + ("" + "foo bar")) {
                            case "":
                                var v98 float32 = 3.14159
                                _ = v98
                            default:
                                v99 := GenS8{F0: int64(v78)}
                                _ = v99
                            }
                            _ = fi86
                        }
                    }
                case "\n\t":
                    for fi100 := 0; fi100 < 11; fi100++ {
                        if v77 {
                            v101 := GenS10{F0: v77}
                            if ((float64((float32(1.5) * float32((1e10 + 3.14159)))) - (((3.14159 + -2.25) / float64(float32((1.5 * 3.14159)))) + float64(1e10))) >= -2.25) {
                                v84 = (fi83 % ((v84) | 1))
                                v77 = (733 == uint32(v85))
                                v77 = (v77 || v77)
                                v2 = append(v2, 826)
                                var v102 int64 = 0
                                _ = v102
                            }
                            var v103 uint8 = 0
                            if v77 {
                                switch ("a" + "\n\t") {
                                case "a":
                                    fmt.Println(v85, v85)
                                case "hello":
                                    _ = len(v1)
                                }
                                fmt.Println(v75)
                            }
                            switch ("a" + "a") {
                            case "hello":
                                v75 = (v75 + v75)
                            case "":
                                switch (fi100 >> (((v85 / ((v85) | 1))) & 63)) {
                                case 2147483647:
                                    var v104 uint32 = 2147483647
                                    _ = v104
                                case -706:
                                    switch (^(v84)) {
                                    case -914:
                                        _ = len(v1)
                                    default:
                                        if v77 {
                                            var v105 uint = 255
                                            var v106 uint16 = (^(v78))
                                            if ((((1.5 / 1.0) * (float64((float32(3.14159) + (0.0 - 1.5))) + 1e10)) - 3.14159) > 3.14159) {
                                                var v107 int64 = 17
                                                if true {
                                                    fmt.Println(v84, fi83)
                                                    var v108 float32 = float32(float64(3.14159))
                                                    if (true && (v105 < 1)) {
                                                        if v77 {
                                                            switch ("\n\t" + "hello") {
                                                            case "":
                                                                var v109 uint = v105
                                                                _ = v109
                                                            case "hello":
                                                                var v110 float32 = ((float32(0.1) * (3.14159 / (0.1 - (v108 * float32(0.1))))) - v108)
                                                                _ = v110
                                                            default:
                                                                fmt.Println(v77)
                                                            }
                                                            v76 += (v76 * -9223372036854775808)
                                                        } else {
                                                            switch (^(fi100)) {
                                                            case 32767:
                                                                if (!((fi83 / ((int(v85)) | 1)) != 42)) {
                                                                    _ = len(v2)
                                                                }
                                                            case 32768:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                            v103 -= 128
                                                            if (619 < -9223372036854775808) {
                                                                var v111 uint16 = 1
                                                                delete(v1, ("hello" + ("hello" + "hello")))
                                                                v103 -= v103
                                                                _ = v111
                                                            } else {
                                                                v84 = (v84 % ((fi83) | 1))
                                                                var v112 uint64 = 1
                                                                v113 := GenS11{F0: "foo bar", F1: v76}
                                                                _ = v112
                                                                _ = v113
                                                            }
                                                        }
                                                        v84 -= fi100
                                                        v77 = ("a" >= "hello")
                                                    }
                                                    _ = v108
                                                } else {
                                                    switch fi83 {
                                                    case -793:
                                                        if (v77 || true) {
                                                            var v114 uint64 = 4294967295
                                                            _ = v114
                                                        } else {
                                                            switch v84 {
                                                            case -365:
                                                                switch (("hello" + "") + ("\n\t" + "hello")) {
                                                                case "\n\t":
                                                                    switch (^(fi100)) {
                                                                    case 128:
                                                                        v75 -= v75
                                                                    }
                                                                case "a":
                                                                    v77 = v77
                                                                default:
                                                                    if v77 {
                                                                        v75 = (v105 ^ (^(v85)))
                                                                        var v115 int64 = v107
                                                                        _ = v115
                                                                    }
                                                                }
                                                            case 42:
                                                                delete(v5, "\n\t")
                                                            case 32767:
                                                                v4 = append(v4, int(v106))
                                                            default:
                                                                if false {
                                                                    v77 = (0 >= 2)
                                                                    v116 := GenS11{F0: "\n\t", F1: v76}
                                                                    _ = v116
                                                                } else {
                                                                    v107 = v76
                                                                }
                                                            }
                                                            v117 := GenS1{F0: v77}
                                                            _ = v117
                                                        }
                                                        v118 := GenS9{F0: v107, F1: v77}
                                                        var v119 string = "hello"
                                                        _ = v118
                                                        _ = v119
                                                    case 42:
                                                        var v120 int16 = 42
                                                        _ = v120
                                                    case 2147483647:
                                                        v105 = 42
                                                    default:
                                                        v121 := GenS3{F0: 47, F1: -1, F2: v76}
                                                        _ = v121
                                                    }
                                                }
                                                var v122 uint16 = 2
                                                _ = v107
                                                _ = v122
                                            } else {
                                                switch (fi100 + fi83) {
                                                case -515:
                                                    switch fi100 {
                                                    case 4294967295:
                                                        v84 = (v84 | v84)
                                                        if len(v6) > 0 {
                                                            v6[0] = ("\n\t" + "a")
                                                        }
                                                        v1[""] = (!("foo bar" >= "\n\t"))
                                                        switch (fi83 << (((v75 / ((18446744073709551615) | 1))) & 63)) {
                                                        case 0:
                                                            fmt.Println(v3)
                                                        case 9223372036854775807:
                                                            var v123 uint32 = 1
                                                            _ = v123
                                                        case 32768:
                                                            var v124 float32 = float32(1.5)
                                                            _ = v124
                                                        }
                                                        v125 := GenS12{F0: v3, F1: 0}
                                                        _ = v125
                                                    case -2147483648:
                                                        var v126 int64 = (v76 * (v76 / ((v76) | 1)))
                                                        _ = v126
                                                    case 1:
                                                        if ((^(v106)) > 42) {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    default:
                                                        v127 := GenS2{F0: ((v84 <= 32767) || (v85 != 1)), F1: (v76 + (^(v76))), F2: 42}
                                                        _ = v127
                                                    }
                                                case 127:
                                                    fmt.Println(v78)
                                                case 1:
                                                    v85 -= (v75 ^ v85)
                                                default:
                                                    switch "a" {
                                                    case "foo bar":
                                                        v128 := GenS9{F0: v76, F1: ((v106 ^ 65535) != 624)}
                                                        _ = v128
                                                    case "hello":
                                                        var v129 float64 = 1e10
                                                        _ = v129
                                                    case "a":
                                                        v77 = (true && ("hello" >= "\n\t"))
                                                    }
                                                }
                                                var v130 int16 = 76
                                                fmt.Println(v85)
                                                _ = v130
                                            }
                                            v3 = 42
                                            v131 := GenS2{F0: (!(117 > 4294967295)), F1: 1, F2: -1}
                                            _ = v105
                                            _ = v106
                                            _ = v131
                                        } else {
                                            switch ("\n\t" + ("hello" + "\n\t")) {
                                            case "\n\t":
                                                switch fi83 {
                                                case -968:
                                                    v6 = append(v6, "\n\t")
                                                case -923:
                                                    switch (v84 * fi83) {
                                                    case 32768:
                                                        if v77 {
                                                            var v132 bool = ((((1e10 + (3.14159 / 0.1)) >= 3.14159) || v77) && (v77 || (42 >= -702)))
                                                            _ = v132
                                                        }
                                                    default:
                                                        v75 += (^(v75))
                                                    }
                                                }
                                            }
                                            if ((v76 << ((v75) & 63)) == v76) {
                                                if (!v77) {
                                                    delete(v5, ("" + "foo bar"))
                                                    v133 := GenS13{F0: (v78 * v78)}
                                                    _ = v133
                                                } else {
                                                    v3 = 0
                                                    v85 -= (v85 << (((v75 * 1)) & 63))
                                                    fmt.Println(fi83, v85, v85)
                                                }
                                                v76 += v76
                                                v134 := GenS14{F0: 593, F1: fi100}
                                                if (!(v3 >= 100)) {
                                                    var v135 int8 = (-(v3))
                                                    var v136 uint16 = v78
                                                    fmt.Println(v103)
                                                    v75 += (v85 % (((^(v85))) | 1))
                                                    _ = v135
                                                    _ = v136
                                                } else {
                                                    v137 := v5[""]
                                                    if (v137 && true) {
                                                        if v77 {
                                                            fmt.Println(v75)
                                                        } else {
                                                            if v137 {
                                                                var v138 bool = (!v137)
                                                                v139 := v5["\n\t"]
                                                                v84 = (fi83 / (((fi83 % ((fi100) | 1))) | 1))
                                                                v85 = v85
                                                                _ = v138
                                                                _ = v139
                                                            }
                                                            v140 := GenS15{F0: 42}
                                                            _ = v140
                                                        }
                                                        v75 = 255
                                                        v141 := GenS12{F0: 100, F1: 0}
                                                        fmt.Println(v76)
                                                        _ = v141
                                                    } else {
                                                        var v142 int8 = (v3 / (((v3 * 100)) | 1))
                                                        _ = v142
                                                    }
                                                    _ = v137
                                                }
                                                fmt.Println(v85)
                                                _ = len(v5)
                                                var v143 uint8 = 29
                                                switch v84 {
                                                case -167:
                                                    var v144 int64 = -626
                                                    _ = v144
                                                case 339:
                                                    _ = len(v1)
                                                }
                                                _ = v134
                                                _ = v143
                                            }
                                            v77 = false
                                        }
                                    }
                                }
                            case "a":
                                v85 -= (v75 ^ (v75 & v85))
                                fmt.Println(v76, fi100)
                                var v145 float32 = (0.1 - 3.14159)
                                v103 = 0
                                _ = v145
                            }
                            _ = v101
                            _ = v103
                        }
                        var v146 int32 = 1
                        _ = fi100
                        _ = v146
                    }
                default:
                    v4 = append(v4, -9223372036854775808)
                }
                _ = fi83
                _ = v84
                _ = v85
            }
        }
        if true {
            _ = len(v6)
        } else {
            v147 := GenS1{F0: (v76 != (v76 * v76))}
            if v77 {
                v148 := GenS11{F0: "foo bar", F1: -9223372036854775808}
                v75 = v75
                _ = v148
            } else {
                v149 := v1[("\n\t" + "hello")]
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v149
            }
            _ = v147
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
