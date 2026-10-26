package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 int16; F1 bool; F2 uint64 }
type GenS7 struct { F0 int32; F1 string; F2 int }
type GenS8 struct { F0 string; F1 int; F2 string }
type GenS9 struct { F0 string; F1 int; F2 uint16 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int; F1 string; F2 uint64 }
type GenS12 struct { F0 int32; F1 uint64; F2 int8 }
func genFn0() uint32 {
    var v5 int64 = 0
    fmt.Println(v5)
    v6 := GenS1{F0: -53, F1: 1, F2: 128}
    v7 := []int{42}
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 int16, p1 bool, p2 int16) int32 {
    p0 = (^(p2))
    p2 += (p2 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint8 = 255
    var v2 int = -854
    var v3 int8 = 0
    v4 := make(map[string]bool)
    if ((v2 / (((v2 * (v2 * (v2 << ((1) & 63))))) | 1)) > (v2 / ((v2) | 1))) {
        _ = len(v4)
    }
    fmt.Println(v1)
    v8 := GenS1{F0: (v3 % (((-(v3))) | 1)), F1: 1, F2: v2}
    if false {
        var v9 string = ("foo bar" + ("a" + ""))
        var v10 string = v9
        if ((true || true) && (1 != -114)) {
            var v11 float64 = 1.5
            _ = v11
        }
        _ = v10
        _ = v9
    } else {
        switch 0 {
        case 42:
            v12 := GenS1{F0: -118, F1: 696, F2: -198}
            _ = v12
        case -2147483648:
            v1 -= v1
        default:
            v2 -= v2
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := GenS1{F0: -7, F1: 0, F2: v2}
            if false {
                if false {
                    v15 := GenS2{F0: "hello", F1: -1}
                    _ = v15
                }
                v16 := GenS1{F0: (v3 % (((v3 ^ (v3 >> ((18446744073709551615) & 63)))) | 1)), F1: 2147483647, F2: 32767}
                for fi17 := 0; fi17 < 1; fi17++ {
                    v1 += v1
                    var v18 uint16 = 65535
                    if (597 < uint32(v1)) {
                        var v19 uint16 = v18
                        _ = v19
                    }
                    v20 := GenS2{F0: "a", F1: -1}
                    _ = fi17
                    _ = v18
                    _ = v20
                }
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        if ((42 < 171) || (false && (18446744073709551615 >= 0))) {
            for fi21 := 0; fi21 < 7; fi21++ {
                switch "" {
                case "a":
                    switch (^(fi21)) {
                    case -724:
                        _ = len(v4)
                    default:
                        v3 = (v3 - (^(v3)))
                    }
                case "\n\t":
                    switch fi21 {
                    case 516:
                        for fi22 := 0; fi22 < 16; fi22++ {
                            v23 := GenS2{F0: "a", F1: 60}
                            v24 := GenS1{F0: (v3 ^ v3), F1: -2147483648, F2: fi22}
                            if false {
                                var v25 int16 = 442
                                _ = v25
                            } else {
                                v2 += (-(v2))
                            }
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    case 32768:
                        v3 -= v3
                    }
                case "foo bar":
                    v26 := v4[(("" + (("foo bar" + "") + "hello")) + "\n\t")]
                    _ = v26
                }
                _ = fi21
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                fmt.Println(v2)
                switch (v2 | fi27) {
                case 2147483647:
                    switch "hello" {
                    case "a":
                        for fi28 := 0; fi28 < 5; fi28++ {
                            v2 = fi28
                            var v29 uint32 = 65535
                            _ = fi28
                            _ = v29
                        }
                    }
                case 1:
                    var v30 int = v2
                    _ = v30
                case 0:
                    for fi31 := 0; fi31 < 10; fi31++ {
                        v2 -= v2
                        var v32 float32 = 1e10
                        _ = fi31
                        _ = v32
                    }
                }
                _ = fi27
            }
            var v33 int32 = 0
            _ = v33
        } else {
            var v34 bool = ("foo bar" > "\n\t")
            var v35 float32 = ((0.1 + 3.14159) * 0.1)
            _ = v34
            _ = v35
        }
    }
    v36 := GenS1{F0: (v3 & v3), F1: -1, F2: 128}
    if false {
        for fi37 := 0; fi37 < 7; fi37++ {
            switch 642 {
            case -9223372036854775808:
                fmt.Println(v2, v2)
            case 324:
                if true {
                    var v38 int16 = 42
                    if ("a" <= "") {
                        var v39 int64 = 789
                        for fi40 := 0; fi40 < 20; fi40++ {
                            v39 = v39
                            v39 += -817
                            _ = fi40
                        }
                        _ = v39
                    } else {
                        switch 32768 {
                        case 128:
                            var v41 int64 = 42
                            _ = v41
                        case 127:
                            v42 := v4[((("" + "\n\t") + (("" + "a") + ("hello" + ""))) + "\n\t")]
                            _ = v42
                        case 2147483647:
                            if false {
                                v3 = (^(v3))
                                v38 = (-(v38))
                                v43 := []int{v2, (fi37 >> ((220) & 63)), 0, v2}
                                _ = v43
                            } else {
                                var v44 int64 = 0
                                _ = v44
                            }
                        }
                    }
                    _ = v38
                } else {
                    delete(v4, "")
                    switch "foo bar" {
                    case "a":
                        if (!true) {
                            var v45 uint = 42
                            v46 := GenS3{F0: true, F1: ("hello" + "a")}
                            var v47 int = fi37
                            _ = v45
                            _ = v46
                            _ = v47
                        } else {
                            if (false && true) {
                                var v48 uint16 = 65535
                                _ = v48
                            } else {
                                var v49 int = (fi37 % ((v2) | 1))
                                var v50 int8 = (v3 >> ((657) & 63))
                                switch v2 {
                                case 32768:
                                    for fi51 := 0; fi51 < 5; fi51++ {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v52 int16 = -1
                                            _ = v52
                                        }
                                        _ = fi51
                                    }
                                case 9223372036854775807:
                                    var v53 int64 = 130
                                    _ = v53
                                case 716:
                                    v54 := GenS1{F0: v3, F1: 0, F2: (v49 * (v49 ^ (fi37 + (v2 & fi37))))}
                                    _ = v54
                                default:
                                    v55 := []int{4294967295}
                                    _ = v55
                                }
                                _ = v49
                                _ = v50
                            }
                            if true {
                                v56 := GenS2{F0: "", F1: 0}
                                v3 = -62
                                _ = v56
                            } else {
                                v4["hello"] = false
                                v3 -= -31
                            }
                            switch (fi37 ^ -488) {
                            case -1:
                                if true {
                                    v57 := GenS4{F0: -9223372036854775808, F1: (!false)}
                                    fmt.Println(v3)
                                    _ = v57
                                }
                            case 32767:
                                if ((v2 * v2) > -1) {
                                    v58 := GenS5{F0: true, F1: 9223372036854775807}
                                    var v59 uint32 = 1
                                    switch (fi37 % ((v2) | 1)) {
                                    case 32767:
                                        v60 := GenS6{F0: -577, F1: true, F2: 4294967295}
                                        _ = v60
                                    case 128:
                                        switch (v2 - 32767) {
                                        case 32768:
                                            v61 := []int{(fi37 ^ (fi37 % (((fi37 / (((fi37 * fi37)) | 1))) | 1))), 656, -2147483648}
                                            _ = v61
                                        default:
                                            v62 := []int{v2, 4294967295}
                                            _ = v62
                                        }
                                    default:
                                        var v63 bool = (!true)
                                        _ = v63
                                    }
                                    _ = v58
                                    _ = v59
                                }
                            case 128:
                                if (-985 <= -48) {
                                    for fi64 := 0; fi64 < 17; fi64++ {
                                        v3 = v3
                                        v65 := GenS5{F0: true, F1: 814}
                                        v4["a"] = true
                                        v1 = (v1 & v1)
                                        _ = fi64
                                        _ = v65
                                    }
                                }
                            }
                        }
                    case "hello":
                        switch "a" {
                        case "":
                            for fi66 := 0; fi66 < 8; fi66++ {
                                v4["\n\t"] = (false || (v3 == v3))
                                fmt.Println(v3)
                                _ = fi66
                            }
                        case "foo bar":
                            fmt.Println(v1, v2, fi37)
                        case "a":
                            switch (fi37 ^ (^(v2))) {
                            case 42:
                                for fi67 := 0; fi67 < 2; fi67++ {
                                    v68 := []int{v2, -713, (^(v2))}
                                    _ = fi67
                                    _ = v68
                                }
                            case -9223372036854775808:
                                v2 -= fi37
                            default:
                                v69 := GenS7{F0: genFn1(1, (v1 <= (v1 ^ (v1 * 0))), 1), F1: ((("foo bar" + ("hello" + ("a" + "hello"))) + "") + "hello"), F2: fi37}
                                _ = v69
                            }
                        }
                    case "foo bar":
                        var v70 uint32 = 2
                        _ = v70
                    }
                }
            }
            var v71 int32 = 2147483647
            _ = fi37
            _ = v71
        }
        v3 -= v3
        for fi72 := 0; fi72 < 2; fi72++ {
            v2 -= v2
            for fi73 := 0; fi73 < 3; fi73++ {
                var v74 int64 = -63
                _ = fi73
                _ = v74
            }
            v75 := GenS6{F0: 32767, F1: true, F2: 2}
            _ = fi72
            _ = v75
        }
    }
    var v76 float32 = (0.0 + 0.1)
    v77 := []int{v2}
    if false {
        switch "foo bar" {
        case "\n\t":
            if len(v77) >= 2 {
                v77 = v77[1:len(v77)]
            }
        case "foo bar":
            if true {
                if (32767 <= -367) {
                    v3 += v3
                    v76 = ((0.0 + v76) * float32(0.0))
                }
                v3 = v3
                var v78 float32 = v76
                _ = v78
            }
        case "a":
            var v79 int32 = 1
            _ = v79
        default:
            for fi80 := 0; fi80 < 12; fi80++ {
                var v81 uint64 = 0
                for fi82 := 0; fi82 < 17; fi82++ {
                    v83 := GenS1{F0: 53, F1: -2147483648, F2: (fi80 % ((-1) | 1))}
                    v3 -= (v3 * v3)
                    var v84 uint32 = 1
                    _ = fi82
                    _ = v83
                    _ = v84
                }
                _ = fi80
                _ = v81
            }
        }
    }
    var v85 bool = (true && true)
    v77 = append(v77, 128)
    if true {
        v86 := GenS8{F0: (("\n\t" + ("hello" + "hello")) + "foo bar"), F1: int(v3), F2: "\n\t"}
        for fi87 := 0; fi87 < 11; fi87++ {
            v85 = ((v1 << ((51) & 63)) > (v1 * v1))
            var v88 float64 = 0.1
            _ = fi87
            _ = v88
        }
        var v89 float64 = float64(1e10)
        _ = v86
        _ = v89
    } else {
        switch ("foo bar" + "") {
        case "\n\t":
            for fi90 := 0; fi90 < 3; fi90++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(v2)
                default:
                    v91 := GenS9{F0: "\n\t", F1: -2147483648, F2: 0}
                    _ = v91
                }
                v3 = v3
                _ = fi90
            }
        case "foo bar":
            var v92 float32 = v76
            _ = v92
        default:
            switch "a" {
            case "foo bar":
                switch v2 {
                case 32768:
                    for fi93 := 0; fi93 < 18; fi93++ {
                        v76 = 3.14159
                        if (true && true) {
                            var v94 uint64 = 225
                            for fi95 := 0; fi95 < 13; fi95++ {
                                if len(v77) > 0 {
                                    v77[0] = v2
                                }
                                var v96 float32 = v76
                                v96 = float32((-2.25 * float64(v76)))
                                _ = fi95
                                _ = v96
                            }
                            _ = v94
                        }
                        if v85 {
                            if v85 {
                                var v97 int = (fi93 * v2)
                                v98 := GenS4{F0: (v2 & v2), F1: v85}
                                _ = v97
                                _ = v98
                            } else {
                                v1 = 42
                                if true {
                                    switch "a" {
                                    case "":
                                        v1 += v1
                                    case "foo bar":
                                        switch (("foo bar" + (("a" + (("foo bar" + "hello") + ("foo bar" + "\n\t"))) + (("" + "\n\t") + "foo bar"))) + (("foo bar" + ("foo bar" + ("hello" + "a"))) + "\n\t")) {
                                        case "a":
                                            v85 = false
                                        default:
                                            switch 127 {
                                            case 830:
                                                v99 := v4[""]
                                                _ = v99
                                            case 114:
                                                var v100 bool = true
                                                _ = v100
                                            default:
                                                delete(v4, "foo bar")
                                            }
                                        }
                                    case "hello":
                                        fmt.Println(v3, v3)
                                    default:
                                        var v101 int32 = 2147483647
                                        _ = v101
                                    }
                                    var v102 float32 = (float32(-2.25) * float32(1.5))
                                    var v103 uint16 = 65535
                                    _ = v102
                                    _ = v103
                                } else {
                                    var v104 uint64 = 449
                                    if v85 {
                                        fmt.Println(v3, v76)
                                    } else {
                                        for fi105 := 0; fi105 < 6; fi105++ {
                                            v1 += v1
                                            fmt.Println(v1)
                                            if v85 {
                                                _ = len(v4)
                                                v2 = v2
                                            } else {
                                                if ((v85 && v85) || false) {
                                                    var v106 float32 = (3.14159 / (v76 / 1.0))
                                                    var v107 float32 = 3.14159
                                                    _ = v106
                                                    _ = v107
                                                }
                                            }
                                            _ = fi105
                                        }
                                        if (((-(v2)) >= v2) || ((fi93 | v2) >= 4294967295)) {
                                            switch ((("\n\t" + "\n\t") + "\n\t") + "foo bar") {
                                            case "":
                                                v108 := GenS10{F0: "hello"}
                                                _ = v108
                                            case "\n\t":
                                                v109 := GenS7{F0: int32(v2), F1: "a", F2: fi93}
                                                _ = v109
                                            case "a":
                                                var v110 string = "\n\t"
                                                _ = v110
                                            default:
                                                v111 := GenS9{F0: ((((("foo bar" + ("" + "a")) + ("\n\t" + "a")) + "a") + "\n\t") + ("foo bar" + "a")), F1: int(v104), F2: 65535}
                                                _ = v111
                                            }
                                        }
                                        fmt.Println(v76, v76)
                                    }
                                    _ = v104
                                }
                            }
                            if v85 {
                                v112 := GenS9{F0: ("\n\t" + ""), F1: 4294967295, F2: 419}
                                for fi113 := 0; fi113 < 14; fi113++ {
                                    v3 = (v3 >> ((933) & 63))
                                    var v114 uint8 = 127
                                    if (((1.5 * -2.25) * float64((v76 + v76))) > float64(v76)) {
                                        if v85 {
                                            if len(v77) > 0 {
                                                v77[0] = (fi93 << ((18446744073709551615) & 63))
                                            }
                                            var v115 uint16 = 66
                                            _ = v115
                                        } else {
                                            delete(v4, "")
                                        }
                                    } else {
                                        delete(v4, "a")
                                        fmt.Println(v2, v85)
                                        fmt.Println(v114, v3, v114)
                                    }
                                    _ = fi113
                                    _ = v114
                                }
                                for fi116 := 0; fi116 < 16; fi116++ {
                                    switch "hello" {
                                    case "foo bar":
                                        var v117 int = v2
                                        _ = v117
                                    default:
                                        var v118 int = fi116
                                        _ = v118
                                    }
                                    _ = fi116
                                }
                                _ = v112
                            } else {
                                for fi119 := 0; fi119 < 4; fi119++ {
                                    var v120 int = (^(v2))
                                    _ = fi119
                                    _ = v120
                                }
                                v77 = append(v77, (fi93 * (fi93 ^ (v2 << ((2) & 63)))))
                            }
                        }
                        _ = fi93
                    }
                case 0:
                    v4[(("" + "a") + ("\n\t" + "foo bar"))] = v85
                case -402:
                    var v121 uint16 = 65535
                    _ = v121
                default:
                    var v122 uint16 = 0
                    _ = v122
                }
            case "":
                v123 := v4["foo bar"]
                _ = v123
            }
        }
        var v124 int = 954
        _ = v124
    }
    for fi125 := 0; fi125 < 9; fi125++ {
        var v126 uint8 = (v1 - (v1 * v1))
        if v85 {
            v76 -= v76
        } else {
            for fi127 := 0; fi127 < 12; fi127++ {
                v77 = append(v77, (-(fi125)))
                var v128 uint64 = 0
                _ = fi127
                _ = v128
            }
        }
        _ = fi125
        _ = v126
    }
    v3 -= v3
    var v129 int64 = 1
    switch (v2 - v2) {
    case 32768:
        switch v2 {
        case 127:
            if false {
                v1 += v1
                if v85 {
                    delete(v4, "foo bar")
                    var v130 float32 = (((((0.1 * (v76 * ((1.5 - 0.1) + v76))) - 3.14159) / 1e10) + -2.25) * ((((0.0 / 3.14159) + v76) / float32(float64(1e10))) * v76))
                    switch v2 {
                    case 9223372036854775807:
                        fmt.Println(v1)
                    default:
                        v131 := GenS11{F0: (v2 + v2), F1: "hello", F2: 4294967295}
                        _ = v131
                    }
                    _ = v130
                } else {
                    v85 = (v85 || (2 != 627))
                    v2 = -1
                }
                if v85 {
                    var v132 int16 = -112
                    v133 := GenS2{F0: "hello", F1: -1}
                    for fi134 := 0; fi134 < 2; fi134++ {
                        v135 := GenS4{F0: fi134, F1: v85}
                        _ = fi134
                        _ = v135
                    }
                    _ = v132
                    _ = v133
                } else {
                    for fi136 := 0; fi136 < 5; fi136++ {
                        v137 := GenS3{F0: v85, F1: ""}
                        v129 -= v129
                        _ = fi136
                        _ = v137
                    }
                    if v85 {
                        v138 := GenS11{F0: 4294967295, F1: ("\n\t" + (((("\n\t" + "") + "a") + (("\n\t" + "") + "a")) + "")), F2: 18446744073709551615}
                        _ = v138
                    }
                }
            } else {
                for fi139 := 0; fi139 < 1; fi139++ {
                    if (v85 && ((-2.25 != ((-2.25 + 0.0) / (v76 - (1e10 / v76)))) && v85)) {
                        v140 := GenS4{F0: (fi139 | (fi139 ^ fi139)), F1: (((("\n\t" + "a") + "\n\t") == "") && (!(!v85)))}
                        _ = v140
                    }
                    _ = fi139
                }
                v129 += (v129 >> ((643) & 63))
            }
        case 32768:
            switch v2 {
            case 128:
                v77 = append(v77, (v2 % (((v2 % ((v2) | 1))) | 1)))
            case 4294967295:
                switch (v2 | (v2 % ((-1) | 1))) {
                case 146:
                    if (0 >= 1000) {
                        var v141 int8 = 120
                        v142 := GenS12{F0: 2147483647, F1: 2, F2: (v141 % ((v3) | 1))}
                        _ = v141
                        _ = v142
                    }
                case 473:
                    var v143 int32 = 531
                    _ = v143
                default:
                    v4["a"] = (v76 <= (v76 / 0.1))
                }
            default:
                v3 = 1
            }
        case 2147483647:
            switch ("a" + "a") {
            case "\n\t":
                for fi144 := 0; fi144 < 6; fi144++ {
                    var v145 int32 = 71
                    if (v85 || ((v85 && v85) || (-677 >= 9223372036854775807))) {
                        fmt.Println(v129, v2, v1)
                        switch "foo bar" {
                        case "":
                            for fi146 := 0; fi146 < 10; fi146++ {
                                if len(v77) > 0 {
                                    v77[0] = (fi144 >> ((255) & 63))
                                }
                                var v147 float64 = -2.25
                                _ = fi146
                                _ = v147
                            }
                        }
                        if v85 {
                            var v148 uint64 = 42
                            var v149 int16 = 32767
                            _ = v148
                            _ = v149
                        }
                    }
                    v1 = 2
                    _ = fi144
                    _ = v145
                }
            case "hello":
                switch int(v129) {
                case 601:
                    fmt.Println(v129, v76)
                case 32767:
                    var v150 int32 = genFn1(1000, v85, -32768)
                    _ = v150
                case 9223372036854775807:
                    var v151 uint = 65535
                    _ = v151
                default:
                    v77 = append(v77, 127)
                }
            default:
                var v152 float32 = v76
                _ = v152
            }
        }
    default:
        v85 = (42 >= -890)
    }
    switch (-(v2)) {
    case 32767:
        for fi153 := 0; fi153 < 11; fi153++ {
            if v85 {
                fmt.Println(v76)
                var v154 string = "\n\t"
                for fi155 := 0; fi155 < 14; fi155++ {
                    v2 -= int(v1)
                    if v85 {
                        switch (v2 & 128) {
                        case 4294967295:
                            if len(v77) >= 2 {
                                v77 = v77[1:len(v77)]
                            }
                        default:
                            var v156 uint16 = 356
                            _ = v156
                        }
                        v1 = v1
                    }
                    fmt.Println(fi153, fi153)
                    _ = fi155
                }
                _ = v154
            } else {
                if len(v77) >= 2 {
                    v77 = v77[1:len(v77)]
                }
                fmt.Println(v2, v129, v129)
            }
            if (v85 && v85) {
                v77 = append(v77, v2)
                _ = len(v4)
            }
            _ = fi153
        }
    case 468:
        fmt.Println(v1, v129, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v76:", v76)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v129
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v76
    _ = v77
    _ = v8
    _ = v85
}
