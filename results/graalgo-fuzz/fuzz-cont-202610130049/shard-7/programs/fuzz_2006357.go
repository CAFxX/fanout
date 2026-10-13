package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8; F1 int; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint32; F1 uint64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int16; F1 bool; F2 int32 }
type GenS16 struct { F0 uint; F1 int8 }
type GenS17 struct { F0 int16; F1 string }
type GenS18 struct { F0 int16; F1 int16; F2 int64 }
type GenS19 struct { F0 int; F1 int32; F2 int64 }
type GenS20 struct { F0 int8; F1 bool; F2 int32 }
type GenS21 struct { F0 int8; F1 uint32; F2 uint }
type GenS22 struct { F0 uint64; F1 int8 }
type GenS23 struct { F0 string; F1 uint64; F2 int }
type GenS24 struct { F0 bool; F1 int; F2 int8 }
type GenS25 struct { F0 int8; F1 int8 }
type GenS26 struct { F0 int8; F1 uint64; F2 uint64 }
func genFn0(p0 int32, p1 bool) int16 {
    if (42 >= 1) {
        v4 := []int{-517, 1, 128}
        v5 := GenS1{F0: (-2.25 > (-2.25 / -2.25))}
        switch (((("\n\t" + "") + ("" + ("foo bar" + ("foo bar" + "\n\t")))) + ("hello" + ("hello" + ""))) + "") {
        case "hello":
            var v6 uint32 = 42
            _ = v6
        case "foo bar":
            var v7 string = ("foo bar" + (("a" + ("\n\t" + ("hello" + "\n\t"))) + ""))
            _ = v7
        case "\n\t":
            fmt.Println(p1)
        default:
            _ = len(v4)
        }
        _ = v4
        _ = v5
    } else {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint32 = 2
            p0 = -615
            v10 := GenS2{F0: -306, F1: "hello"}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS2{F0: -9223372036854775808, F1: (("" + "a") + "a")}
        fmt.Println(fi11, fi11, p1)
        _ = fi11
        _ = v12
    }
    p1 = (p1 && (false && (2 < 787)))
    fmt.Println(p1, p1)
    switch 42 {
    case 9223372036854775807:
        if p1 {
            var v13 uint8 = 127
            _ = v13
        }
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint64) int8 {
    var v14 int32 = 1
    var v15 uint16 = 255
    v14 -= v14
    _ = p0
    _ = v14
    _ = v15
    return -61
}

func main() {
    var v1 bool = true
    v2 := make(map[int]bool)
    var v3 float32 = 0.1
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int = (fi16 - 2147483647)
        for fi18 := 0; fi18 < 5; fi18++ {
            v2[314] = (1 < 65535)
            var v19 string = "hello"
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    v20 := []int{4294967295, 32767, 9223372036854775807, -2147483648}
    fmt.Println(v3)
    v20 = append(v20, -9223372036854775808)
    if (v3 != v3) {
        for fi21 := 0; fi21 < 3; fi21++ {
            switch ("\n\t" + (("a" + "\n\t") + "a")) {
            case "hello":
                fmt.Println(v3, v3)
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                v1 = true
                fmt.Println(v3, fi21, v3)
                _ = fi22
            }
            _ = fi21
        }
        if (!true) {
            if (((0 >= 65535) || v1) && ((!(v1 || true)) && ((v1 || true) && (4294967295 != 2147483647)))) {
                v23 := GenS3{F0: 255, F1: -2147483648, F2: 127}
                var v24 int32 = -1
                _ = v23
                _ = v24
            } else {
                var v25 int8 = -1
                for fi26 := 0; fi26 < 6; fi26++ {
                    v2[-744] = ((0.0 - (0.0 + 0.1)) <= -2.25)
                    var v27 int64 = 42
                    fmt.Println(fi26, v1, v3)
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
            switch -2147483648 {
            case 4294967295:
                v28 := GenS4{F0: 1, F1: 42}
                _ = v28
            case 128:
                fmt.Println(v3, v1)
            default:
                if len(v20) > 0 {
                    v20[0] = 42
                }
            }
            v3 = 1e10
        } else {
            var v29 uint = 65535
            _ = v29
        }
    } else {
        switch "a" {
        case "a":
            var v30 int32 = 0
            _ = v30
        case "hello":
            fmt.Println(v1)
        case "foo bar":
            v31 := GenS5{F0: 745}
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            v1 = v1
            if true {
                _ = len(v2)
            }
            switch fi33 {
            case 770:
                v34 := GenS6{F0: (fi33 * 910)}
                _ = v34
            case -2147483648:
                if true {
                    v3 = v3
                    v1 = v1
                    v1 = (0 > 515)
                }
            }
            _ = fi33
        }
        var v35 int16 = 350
        switch "" {
        case "":
            var v36 uint64 = 963
            _ = v36
        }
        _ = fi32
        _ = v35
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        switch ("hello" + "foo bar") {
        case "hello":
            switch fi37 {
            case 127:
                v3 = (1.5 / 0.1)
            case 756:
                v1 = ("" >= "a")
            }
        }
        if (false && false) {
            var v38 string = "hello"
            _ = v38
        } else {
            v1 = true
            var v39 float32 = v3
            _ = v39
        }
        _ = fi37
    }
    v40 := GenS1{F0: (float64((float32(1.5) + (v3 + 0.0))) != -2.25)}
    if (("\n\t" + (("a" + "hello") + "a")) < "") {
        var v41 int32 = -2147483648
        _ = len(v20)
        switch -973 {
        case 4294967295:
            v42 := GenS7{F0: 492}
            _ = v42
        case 32768:
            var v43 uint = 4294967295
            _ = v43
        }
        _ = v41
    } else {
        v1 = false
        if ("\n\t" >= "\n\t") {
            v44 := GenS4{F0: 1, F1: 1}
            _ = v44
        }
        for fi45 := 0; fi45 < 8; fi45++ {
            switch "foo bar" {
            case "\n\t":
                var v46 int = -809
                _ = v46
            case "":
                var v47 uint8 = 127
                _ = v47
            case "hello":
                if (!((fi45 <= -9223372036854775808) || (!false))) {
                    if true {
                        v3 += (v3 / 3.14159)
                        v1 = (v1 || ((-1 != 0) && (v1 && false)))
                        v48 := GenS5{F0: 859}
                        _ = v48
                    } else {
                        v3 = (((v3 - v3) * v3) + v3)
                        v3 -= ((1e10 / v3) + -2.25)
                    }
                    for fi49 := 0; fi49 < 18; fi49++ {
                        if (!(0.0 < (1.5 + v3))) {
                            var v50 string = ""
                            _ = v50
                        } else {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            switch (fi45 * fi45) {
                            case 725:
                                if len(v20) > 0 {
                                    v20[0] = 128
                                }
                            case -9223372036854775808:
                                var v51 float64 = 0.1
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                var v52 string = ""
                                v3 += float32(3.14159)
                                _ = v51
                                _ = v52
                            case 128:
                                var v53 int8 = 42
                                if v1 {
                                    if (1 <= 128) {
                                        switch "\n\t" {
                                        case "a":
                                            v1 = ((!(2147483647 >= -1)) && v1)
                                        case "hello":
                                            var v54 bool = true
                                            fmt.Println(v53, fi45, v54)
                                            _ = v54
                                        default:
                                            var v55 uint64 = 42
                                            _ = v55
                                        }
                                    }
                                    v53 -= v53
                                }
                                _ = v53
                            }
                            if (-238 > 42) {
                                var v56 string = ("" + ("" + ("hello" + "")))
                                _ = v56
                            }
                            switch "" {
                            case "hello":
                                v57 := GenS1{F0: v1}
                                _ = v57
                            default:
                                v20 = append(v20, -860)
                            }
                        }
                        _ = fi49
                    }
                    for fi58 := 0; fi58 < 20; fi58++ {
                        var v59 int32 = 1
                        _ = len(v2)
                        switch fi58 {
                        case -304:
                            v60 := GenS1{F0: ((v3 + v3) != v3)}
                            _ = v60
                        default:
                            fmt.Println(fi45)
                        }
                        _ = fi58
                        _ = v59
                    }
                } else {
                    v61 := GenS5{F0: 0}
                    var v62 int16 = 1000
                    v63 := GenS8{F0: ("" + "hello")}
                    _ = v61
                    _ = v62
                    _ = v63
                }
            }
            if v1 {
                if (0 <= 423) {
                    v3 = (0.0 * (v3 / v3))
                    if v1 {
                        for fi64 := 0; fi64 < 5; fi64++ {
                            if len(v20) > 0 {
                                v20[0] = fi64
                            }
                            v65 := GenS6{F0: 127}
                            v2[fi64] = (!true)
                            if ((((-2.25 + 0.1) - (1e10 / (0.0 + 0.1))) == (0.0 / float64(((v3 - v3) - (v3 - 1.5))))) || (!((!v1) || true))) {
                                var v66 uint = 565
                                var v67 uint = (^(v66))
                                _ = v66
                                _ = v67
                            }
                            _ = fi64
                            _ = v65
                        }
                        var v68 int16 = -1
                        var v69 float32 = float32(-2.25)
                        _ = v68
                        _ = v69
                    } else {
                        if (-552 < 1) {
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v3, v3)
                            case "foo bar":
                                for fi70 := 0; fi70 < 4; fi70++ {
                                    fmt.Println(fi45)
                                    var v71 uint32 = 42
                                    v71 = v71
                                    _ = fi70
                                    _ = v71
                                }
                            case "\n\t":
                                if (673 >= 0) {
                                    var v72 int64 = 0
                                    fmt.Println(v1, v3, fi45)
                                    for fi73 := 0; fi73 < 15; fi73++ {
                                        var v74 int64 = (v72 % (((v72 ^ (v72 << ((254) & 63)))) | 1))
                                        if false {
                                            v75 := GenS4{F0: 1, F1: 42}
                                            var v76 uint = 4294967295
                                            _ = len(v20)
                                            if len(v20) > 0 {
                                                v20[0] = fi73
                                            }
                                            if v1 {
                                                var v77 uint64 = 42
                                                v72 = v74
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(v72)
                                                default:
                                                    switch ("\n\t" + ("\n\t" + ("a" + "\n\t"))) {
                                                    case "a":
                                                        v77 += (v77 >> ((v76) & 63))
                                                    case "":
                                                        if false {
                                                            var v78 uint8 = 127
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                            _ = v78
                                                        }
                                                    default:
                                                        var v79 uint64 = 42
                                                        fmt.Println(v79, v1, v79)
                                                        if len(v20) >= 2 {
                                                            v20 = v20[1:len(v20)]
                                                        }
                                                        var v80 int64 = (v74 ^ -9223372036854775808)
                                                        if len(v20) > 0 {
                                                            v20[0] = -283
                                                        }
                                                        var v81 float64 = 0.1
                                                        if v1 {
                                                            v3 -= v3
                                                            v82 := GenS8{F0: "foo bar"}
                                                            _ = v82
                                                        } else {
                                                            v76 = v76
                                                        }
                                                        switch ("\n\t" + (("a" + ("" + "")) + "hello")) {
                                                        case "hello":
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v77 -= v77
                                                            case "":
                                                                v83 := GenS9{F0: 2}
                                                                _ = v83
                                                            case "hello":
                                                                var v84 float32 = 0.1
                                                                _ = v84
                                                            }
                                                        case "\n\t":
                                                            v1 = (!v1)
                                                        case "foo bar":
                                                            fmt.Println(v79, v81)
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                        var v85 int64 = v74
                                                        switch (fi73 % (((fi45 % ((fi45) | 1))) | 1)) {
                                                        case -609:
                                                            v80 += (v72 & 545)
                                                        default:
                                                            switch "\n\t" {
                                                            case "":
                                                                switch fi45 {
                                                                case -647:
                                                                    var v86 int = (fi73 << (((v76 - 93)) & 63))
                                                                    _ = v86
                                                                case -1:
                                                                    if v1 {
                                                                        var v87 uint64 = 0
                                                                        v88 := GenS10{F0: 0}
                                                                        var v89 uint32 = 4294967295
                                                                        v1 = ((!(((fi73 - fi45) != fi45) && (!(-2147483648 != -2147483648)))) && true)
                                                                        _ = v87
                                                                        _ = v88
                                                                        _ = v89
                                                                    } else {
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                        var v90 bool = v1
                                                                        _ = v90
                                                                    }
                                                                    v91 := GenS10{F0: 53}
                                                                    _ = v91
                                                                }
                                                            case "hello":
                                                                fmt.Println(v3, v77)
                                                            case "\n\t":
                                                                fmt.Println(v76, fi45)
                                                            }
                                                        }
                                                        _ = v79
                                                        _ = v80
                                                        _ = v81
                                                        _ = v85
                                                    }
                                                }
                                                delete(v2, 1)
                                                var v92 int16 = 1
                                                _ = v77
                                                _ = v92
                                            } else {
                                                v93 := GenS8{F0: "hello"}
                                                fmt.Println(v3, fi73)
                                                _ = len(v20)
                                                var v94 int8 = 0
                                                _ = v93
                                                _ = v94
                                            }
                                            fmt.Println(fi45, v1, v1)
                                            _ = v75
                                            _ = v76
                                        } else {
                                            v3 -= v3
                                            if false {
                                                v1 = (65535 > 2147483648)
                                                delete(v2, (fi45 >> ((1) & 63)))
                                                if false {
                                                    v20 = append(v20, fi73)
                                                    v95 := GenS5{F0: 2}
                                                    _ = v95
                                                }
                                            }
                                            switch "" {
                                            case "\n\t":
                                                if false {
                                                    _ = len(v20)
                                                    if (-2.25 != 1.5) {
                                                        if false {
                                                            var v96 int16 = genFn0(2147483647, false)
                                                            if v1 {
                                                                var v97 float64 = 1.5
                                                                var v98 uint32 = 0
                                                                _ = v97
                                                                _ = v98
                                                            } else {
                                                                v99 := GenS1{F0: false}
                                                                if (float32(float64((((((float32(0.1) - float32((float64((v3 + ((((((1.5 * (0.0 - v3)) / ((v3 + -2.25) / -2.25)) + float32((3.14159 + (3.14159 + (-2.25 - 1.5))))) + float32(-2.25)) - 3.14159) / (v3 - 0.0)))) * float64((v3 - (-2.25 + v3)))))) + v3) + v3) / float32((1e10 / (((1e10 + float64(v3)) - (1.5 + (0.0 + 1.5))) + (0.0 + 0.1))))) + (0.0 / v3)))) == 3.14159) {
                                                                    var v100 int16 = 32767
                                                                    delete(v2, fi45)
                                                                    _ = v100
                                                                } else {
                                                                    var v101 string = ("foo bar" + "foo bar")
                                                                    var v102 uint = 65535
                                                                    _ = v101
                                                                    _ = v102
                                                                }
                                                                if (v1 && (38 > 2)) {
                                                                    if ((1 != -35) && true) {
                                                                        v103 := GenS11{F0: 0, F1: 1}
                                                                        _ = v103
                                                                    }
                                                                    v104 := GenS7{F0: 1}
                                                                    fmt.Println(fi73)
                                                                    _ = v104
                                                                } else {
                                                                    v105 := GenS2{F0: v74, F1: ""}
                                                                    v106 := GenS3{F0: 2, F1: 2147483647, F2: 255}
                                                                    if (true && v1) {
                                                                        v96 += genFn0(-1, (!(1 != v96)))
                                                                        delete(v2, (fi45 ^ 32767))
                                                                    }
                                                                    v3 += v3
                                                                    v107 := GenS12{F0: 366}
                                                                    _ = v105
                                                                    _ = v106
                                                                    _ = v107
                                                                }
                                                                if v1 {
                                                                    fmt.Println(v1, v74)
                                                                    switch (^(fi45)) {
                                                                    case -9223372036854775808:
                                                                        v20 = append(v20, (-(fi73)))
                                                                    }
                                                                    if false {
                                                                        if false {
                                                                            if len(v20) > 0 {
                                                                                v20[0] = (fi73 >> ((963) & 63))
                                                                            }
                                                                        } else {
                                                                            switch fi73 {
                                                                            case -708:
                                                                                v74 = 9223372036854775807
                                                                                v3 = v3
                                                                                fmt.Println(fi73, v96, fi73)
                                                                                if (255 >= 1) {
                                                                                    v108 := GenS13{F0: 128}
                                                                                    if true {
                                                                                        if true {
                                                                                            if true {
                                                                                                if (-128 > 0) {
                                                                                                    v109 := GenS14{F0: (v72 % ((v74) | 1))}
                                                                                                    v110 := GenS13{F0: fi45}
                                                                                                    fmt.Println(v96)
                                                                                                    _ = v109
                                                                                                    _ = v110
                                                                                                }
                                                                                            }
                                                                                            var v111 int16 = v96
                                                                                            if true {
                                                                                                var v112 bool = (("\n\t" + "hello") >= ("" + ""))
                                                                                                v96 = v96
                                                                                                var v113 float64 = 1e10
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            } else {
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    v114 := GenS3{F0: 128, F1: (fi73 / ((fi73) | 1)), F2: 205}
                                                                                                    _ = v114
                                                                                                default:
                                                                                                    v115 := GenS15{F0: v96, F1: true, F2: -2147483648}
                                                                                                    _ = v115
                                                                                                }
                                                                                                if len(v20) > 0 {
                                                                                                    v20[0] = (fi73 % ((461) | 1))
                                                                                                }
                                                                                            }
                                                                                            _ = v111
                                                                                        }
                                                                                    }
                                                                                    v116 := GenS16{F0: 18446744073709551615, F1: 1}
                                                                                    _ = v108
                                                                                    _ = v116
                                                                                } else {
                                                                                    if false {
                                                                                        var v117 int8 = -4
                                                                                        var v118 int16 = (v96 ^ v96)
                                                                                        if (9223372036854775807 > (fi45 & 128)) {
                                                                                            v119 := GenS6{F0: (fi73 & fi45)}
                                                                                            if v1 {
                                                                                                _ = len(v2)
                                                                                            }
                                                                                            if (65535 < 819) {
                                                                                                var v120 int64 = -1
                                                                                                _ = v120
                                                                                            } else {
                                                                                                if false {
                                                                                                    v1 = (2147483647 > 1)
                                                                                                    if v1 {
                                                                                                        switch fi73 {
                                                                                                        case -9223372036854775808:
                                                                                                            v121 := GenS17{F0: 32767, F1: ""}
                                                                                                            _ = v121
                                                                                                        case -176:
                                                                                                            v122 := GenS18{F0: 42, F1: 32767, F2: (v74 >> ((4294967295) & 63))}
                                                                                                            if len(v20) > 0 {
                                                                                                                v20[0] = fi45
                                                                                                            }
                                                                                                            var v123 int = -9223372036854775808
                                                                                                            _ = v122
                                                                                                            _ = v123
                                                                                                        default:
                                                                                                            var v124 float32 = v3
                                                                                                            _ = v124
                                                                                                        }
                                                                                                    }
                                                                                                    var v125 string = "\n\t"
                                                                                                    var v126 bool = (v1 || v1)
                                                                                                    var v127 int16 = v118
                                                                                                    _ = v125
                                                                                                    _ = v126
                                                                                                    _ = v127
                                                                                                }
                                                                                                var v128 uint = 255
                                                                                                v129 := GenS15{F0: genFn0(560, true), F1: (65535 == 42), F2: -1}
                                                                                                _ = v128
                                                                                                _ = v129
                                                                                            }
                                                                                            v117 = v117
                                                                                            v74 -= v72
                                                                                            var v130 string = "hello"
                                                                                            var v131 uint8 = 127
                                                                                            switch v130 {
                                                                                            case "foo bar":
                                                                                                v132 := GenS6{F0: 128}
                                                                                                _ = v132
                                                                                            case "a":
                                                                                                if (!true) {
                                                                                                    delete(v2, fi73)
                                                                                                }
                                                                                            default:
                                                                                                var v133 uint8 = (v131 + v131)
                                                                                                _ = v133
                                                                                            }
                                                                                            _ = v119
                                                                                            _ = v130
                                                                                            _ = v131
                                                                                        }
                                                                                        _ = v117
                                                                                        _ = v118
                                                                                    }
                                                                                }
                                                                                var v134 bool = (18446744073709551615 != 143)
                                                                                _ = v134
                                                                            default:
                                                                                fmt.Println(v74, fi45, v1)
                                                                                delete(v2, 128)
                                                                            }
                                                                            switch -589 {
                                                                            case -2147483648:
                                                                                v74 += 437
                                                                                switch "" {
                                                                                case "a":
                                                                                    switch -626 {
                                                                                    case -9223372036854775808:
                                                                                        var v135 int64 = v74
                                                                                        _ = v135
                                                                                    case 42:
                                                                                        if false {
                                                                                            v136 := GenS9{F0: 0}
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                switch 1 {
                                                                                                case -922:
                                                                                                    v137 := GenS7{F0: 262}
                                                                                                    _ = v137
                                                                                                case 304:
                                                                                                    v96 -= v96
                                                                                                    v138 := GenS14{F0: v72}
                                                                                                    var v139 int64 = (-(v74))
                                                                                                    var v140 int32 = 42
                                                                                                    fmt.Println(v140, fi45, fi73)
                                                                                                    _ = v138
                                                                                                    _ = v139
                                                                                                    _ = v140
                                                                                                default:
                                                                                                    v1 = (!true)
                                                                                                }
                                                                                            case "hello":
                                                                                                v141 := GenS1{F0: v1}
                                                                                                _ = v141
                                                                                            default:
                                                                                                _ = len(v2)
                                                                                            }
                                                                                            _ = v136
                                                                                        } else {
                                                                                            if true {
                                                                                                var v142 int64 = int64(v96)
                                                                                                _ = v142
                                                                                            } else {
                                                                                                if v1 {
                                                                                                    var v143 string = ((("" + "hello") + (("foo bar" + "\n\t") + "a")) + ("a" + "foo bar"))
                                                                                                    _ = v143
                                                                                                } else {
                                                                                                    v72 = v72
                                                                                                    v3 += (v3 * v3)
                                                                                                    fmt.Println(v74, v72)
                                                                                                    if (!v1) {
                                                                                                        v3 += (((((v3 * v3) * (-2.25 - v3)) - 0.1) + 0.1) / ((v3 / (0.0 * v3)) / (v3 * v3)))
                                                                                                        switch (-(fi45)) {
                                                                                                        case 0:
                                                                                                            if true {
                                                                                                                v3 -= float32(0.0)
                                                                                                                v3 += v3
                                                                                                                if v1 {
                                                                                                                    v96 = (v96 + v96)
                                                                                                                    if len(v20) > 0 {
                                                                                                                        v20[0] = 4294967295
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v144 string = "hello"
                                                                                                                    v3 -= 0.0
                                                                                                                    _ = v144
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v74 = 391
                                                                                                    }
                                                                                                    if v1 {
                                                                                                        if (-1 > -1) {
                                                                                                            v145 := GenS10{F0: 24}
                                                                                                            var v146 uint = 373
                                                                                                            if v1 {
                                                                                                                fmt.Println(v96)
                                                                                                                v1 = v1
                                                                                                                v74 -= v72
                                                                                                            } else {
                                                                                                                var v147 int64 = (v72 * v74)
                                                                                                                fmt.Println(v147, v3)
                                                                                                                var v148 uint64 = 2
                                                                                                                _ = v147
                                                                                                                _ = v148
                                                                                                            }
                                                                                                            _ = v145
                                                                                                            _ = v146
                                                                                                        }
                                                                                                        v149 := GenS2{F0: (v74 + 42), F1: "a"}
                                                                                                        var v150 uint64 = 18446744073709551615
                                                                                                        _ = v149
                                                                                                        _ = v150
                                                                                                    }
                                                                                                    if true {
                                                                                                        var v151 uint16 = 32767
                                                                                                        fmt.Println(v151, v72)
                                                                                                        switch 605 {
                                                                                                        case -694:
                                                                                                            v20 = append(v20, (fi73 >> ((95) & 63)))
                                                                                                        }
                                                                                                        _ = v151
                                                                                                    } else {
                                                                                                        var v152 int16 = v96
                                                                                                        var v153 float32 = ((((v3 + v3) / 1.0) * (v3 * v3)) - v3)
                                                                                                        var v154 uint8 = 255
                                                                                                        v153 = 0.1
                                                                                                        _ = v152
                                                                                                        _ = v153
                                                                                                        _ = v154
                                                                                                    }
                                                                                                    switch (fi45 / (((fi45 * 1)) | 1)) {
                                                                                                    case 1:
                                                                                                        if v1 {
                                                                                                            var v155 bool = false
                                                                                                            _ = v155
                                                                                                        }
                                                                                                        v156 := GenS19{F0: fi45, F1: 0, F2: 0}
                                                                                                        _ = v156
                                                                                                    case 85:
                                                                                                        var v157 int64 = (v72 | (v74 << ((4294967295) & 63)))
                                                                                                        _ = v157
                                                                                                    case -9223372036854775808:
                                                                                                        if (2147483648 != 0) {
                                                                                                            var v158 int8 = -105
                                                                                                            v1 = true
                                                                                                            _ = v158
                                                                                                        } else {
                                                                                                            var v159 uint32 = 125
                                                                                                            _ = v159
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    switch (-(fi45)) {
                                                                                                    case 127:
                                                                                                        v74 = (v74 * v74)
                                                                                                    default:
                                                                                                        v160 := v2[(fi73 / (((fi73 << ((1) & 63))) | 1))]
                                                                                                        _ = v160
                                                                                                    }
                                                                                                }
                                                                                                delete(v2, (fi45 >> ((65535) & 63)))
                                                                                            }
                                                                                            var v161 string = ""
                                                                                            v72 -= (v72 << ((2) & 63))
                                                                                            v96 = (v96 | v96)
                                                                                            v162 := GenS20{F0: 1, F1: v1, F2: 693}
                                                                                            var v163 bool = v1
                                                                                            _ = v161
                                                                                            _ = v162
                                                                                            _ = v163
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        if v1 {
                                                                                            if ((0.0 / ((((1e10 - 1.5) + (((((3.14159 - 1e10) / 1.5) / 1.0) * ((((float64(v3) / 0.1) * 1.5) + (0.1 / 1.5)) / float64(v3))) - (1.5 * 3.14159))) - float64(v3)) + 0.0)) == float64(3.14159)) {
                                                                                                if len(v20) >= 2 {
                                                                                                    v20 = v20[1:len(v20)]
                                                                                                }
                                                                                            }
                                                                                            var v164 float32 = v3
                                                                                            if true {
                                                                                                v72 += (v74 ^ (v74 / ((42) | 1)))
                                                                                                if v1 {
                                                                                                    v74 = (-(v74))
                                                                                                    switch "" {
                                                                                                    case "hello":
                                                                                                        fmt.Println(fi73, v74)
                                                                                                    case "a":
                                                                                                        fmt.Println(v164, v164)
                                                                                                    case "foo bar":
                                                                                                        switch (("foo bar" + "") + "a") {
                                                                                                        case "hello":
                                                                                                            if v1 {
                                                                                                                switch fi45 {
                                                                                                                case 9223372036854775807:
                                                                                                                    v165 := v2[fi45]
                                                                                                                    _ = v165
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v74, v74, v72)
                                                                                                                var v166 uint = 479
                                                                                                                _ = v166
                                                                                                            }
                                                                                                        case "a":
                                                                                                            delete(v2, -9223372036854775808)
                                                                                                        case "":
                                                                                                            var v167 int32 = -867
                                                                                                            _ = v167
                                                                                                        }
                                                                                                        fmt.Println(v72, v74)
                                                                                                    }
                                                                                                    v72 = (-(v72))
                                                                                                } else {
                                                                                                    var v168 int32 = -2147483648
                                                                                                    switch fi45 {
                                                                                                    case -9223372036854775808:
                                                                                                        delete(v2, (fi73 * (fi45 * int(v168))))
                                                                                                    case -1:
                                                                                                        v3 = v3
                                                                                                    default:
                                                                                                        if v1 {
                                                                                                            var v169 uint8 = 255
                                                                                                            _ = v169
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v168
                                                                                                }
                                                                                                v170 := GenS10{F0: 0}
                                                                                                switch "\n\t" {
                                                                                                case "hello":
                                                                                                    var v171 int64 = v72
                                                                                                    _ = v171
                                                                                                case "a":
                                                                                                    fmt.Println(v74, fi45, v3)
                                                                                                }
                                                                                                _ = v170
                                                                                            }
                                                                                            _ = v164
                                                                                        }
                                                                                    default:
                                                                                        v96 = v96
                                                                                    }
                                                                                }
                                                                            case -871:
                                                                                v172 := GenS12{F0: 2}
                                                                                _ = v172
                                                                            }
                                                                        }
                                                                        v173 := v2[int(v72)]
                                                                        _ = v173
                                                                    } else {
                                                                        if (255 != 2) {
                                                                            _ = len(v2)
                                                                            fmt.Println(v1)
                                                                            var v174 uint8 = 0
                                                                            var v175 uint16 = 0
                                                                            _ = v174
                                                                            _ = v175
                                                                        }
                                                                        fmt.Println(fi73, v96, v72)
                                                                        v176 := v2[9223372036854775807]
                                                                        _ = v176
                                                                    }
                                                                } else {
                                                                    _ = len(v2)
                                                                }
                                                                if (false && ((!v1) || v1)) {
                                                                    var v177 string = ""
                                                                    if false {
                                                                        var v178 float32 = v3
                                                                        var v179 int32 = 1
                                                                        _ = v178
                                                                        _ = v179
                                                                    } else {
                                                                        var v180 float64 = (3.14159 / float64(v3))
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            if (42 >= 0) {
                                                                                switch fi73 {
                                                                                case 1:
                                                                                    var v181 int32 = -41
                                                                                    _ = v181
                                                                                case 0:
                                                                                    v182 := GenS21{F0: -128, F1: 2, F2: 255}
                                                                                    _ = v182
                                                                                case 9223372036854775807:
                                                                                    _ = len(v2)
                                                                                }
                                                                            }
                                                                        case "hello":
                                                                            switch (fi73 & fi73) {
                                                                            case -1:
                                                                                if v1 {
                                                                                    v20 = append(v20, 451)
                                                                                    var v183 uint16 = 32767
                                                                                    v2[int(v74)] = (!v1)
                                                                                    fmt.Println(v183, fi45)
                                                                                    var v184 int64 = -1
                                                                                    _ = v183
                                                                                    _ = v184
                                                                                }
                                                                                v1 = ((float64(1e10) + v180) >= v180)
                                                                            case -2147483648:
                                                                                var v185 uint64 = 964
                                                                                _ = v185
                                                                            case -9223372036854775808:
                                                                                v1 = v1
                                                                                if v1 {
                                                                                    v177 = v177
                                                                                } else {
                                                                                    if v1 {
                                                                                        var v186 int64 = (-(v74))
                                                                                        if (v3 == 1e10) {
                                                                                            var v187 int16 = (v96 << ((2) & 63))
                                                                                            _ = v187
                                                                                        }
                                                                                        _ = v186
                                                                                    }
                                                                                    v3 = float32(((v180 * (v180 / v180)) * ((v180 * (v180 * v180)) / 1.5)))
                                                                                    var v188 int64 = v72
                                                                                    _ = v188
                                                                                }
                                                                                delete(v2, fi45)
                                                                            default:
                                                                                delete(v2, (-(fi73)))
                                                                            }
                                                                        case "\n\t":
                                                                            v189 := GenS15{F0: (v96 & (-(v96))), F1: v1, F2: int32(v96)}
                                                                            _ = v189
                                                                        default:
                                                                            if (fi45 < (fi73 / ((-627) | 1))) {
                                                                                v190 := v2[fi73]
                                                                                v1 = (!v190)
                                                                                var v191 bool = (v3 != v3)
                                                                                _ = v190
                                                                                _ = v191
                                                                            } else {
                                                                                if v1 {
                                                                                    var v192 float32 = v3
                                                                                    if true {
                                                                                        fmt.Println(fi45, v192)
                                                                                        var v193 float32 = v3
                                                                                        _ = v193
                                                                                    } else {
                                                                                        if v1 {
                                                                                            v192 = v192
                                                                                        }
                                                                                        var v194 uint32 = 1
                                                                                        var v195 float32 = v3
                                                                                        _ = v194
                                                                                        _ = v195
                                                                                    }
                                                                                    _ = v192
                                                                                }
                                                                                var v196 float32 = v3
                                                                                _ = v196
                                                                            }
                                                                        }
                                                                        _ = v180
                                                                    }
                                                                    v1 = (250 > 463)
                                                                    _ = v177
                                                                } else {
                                                                    var v197 string = "a"
                                                                    var v198 uint16 = 807
                                                                    switch fi45 {
                                                                    case 428:
                                                                        var v199 float32 = v3
                                                                        _ = v199
                                                                    case 128:
                                                                        v3 = v3
                                                                    case -396:
                                                                        v200 := GenS16{F0: 255, F1: 100}
                                                                        v197 = v197
                                                                        var v201 float64 = 1e10
                                                                        _ = v200
                                                                        _ = v201
                                                                    default:
                                                                        v202 := GenS17{F0: v96, F1: v197}
                                                                        _ = v202
                                                                    }
                                                                    _ = v197
                                                                    _ = v198
                                                                }
                                                                _ = v99
                                                            }
                                                            var v203 int32 = -2147483648
                                                            _ = v203
                                                            _ = v96
                                                        }
                                                        if (481 < 134) {
                                                            v72 -= v72
                                                        } else {
                                                            var v204 int8 = 42
                                                            v205 := v2[fi73]
                                                            _ = v204
                                                            _ = v205
                                                        }
                                                        if false {
                                                            if len(v20) > 0 {
                                                                v20[0] = (-(fi73))
                                                            }
                                                        }
                                                    }
                                                    if ("hello" >= "hello") {
                                                        if false {
                                                            var v206 float32 = v3
                                                            fmt.Println(fi45, fi73)
                                                            var v207 int16 = 1000
                                                            _ = v206
                                                            _ = v207
                                                        } else {
                                                            v74 -= (v74 + 481)
                                                        }
                                                    } else {
                                                        if (187 != 781) {
                                                            v208 := GenS5{F0: 0}
                                                            var v209 int8 = 42
                                                            if v1 {
                                                                var v210 uint = 65535
                                                                _ = v210
                                                            } else {
                                                                v211 := GenS22{F0: 2, F1: genFn1(1)}
                                                                _ = v211
                                                            }
                                                            _ = v208
                                                            _ = v209
                                                        } else {
                                                            v212 := GenS19{F0: (fi45 * (fi45 / (((fi73 + (fi73 - (fi73 - (fi73 << ((411) & 63)))))) | 1))), F1: 2147483647, F2: 522}
                                                            _ = v212
                                                        }
                                                        v1 = (!(true && (42 < 338)))
                                                        v72 = (v72 - (v74 - v74))
                                                        v213 := GenS23{F0: "hello", F1: 42, F2: 9223372036854775807}
                                                        var v214 uint64 = 452
                                                        _ = v213
                                                        _ = v214
                                                    }
                                                    if v1 {
                                                        fmt.Println(v72, v72, fi73)
                                                        if (255 != 42) {
                                                            var v215 uint64 = 2
                                                            _ = v215
                                                        } else {
                                                            v74 -= v72
                                                            var v216 uint32 = 2
                                                            _ = v216
                                                        }
                                                    }
                                                }
                                            case "a":
                                                v3 = 1e10
                                            case "":
                                                v217 := v2[(fi45 - (fi45 + fi45))]
                                                _ = v217
                                            default:
                                                fmt.Println(v74, v72, v74)
                                            }
                                        }
                                        if v1 {
                                            v3 = (v3 + 0.0)
                                            var v218 uint = 18446744073709551615
                                            _ = v218
                                        }
                                        _ = fi73
                                        _ = v74
                                    }
                                    _ = v72
                                } else {
                                    var v219 float64 = -2.25
                                    var v220 uint16 = 943
                                    fmt.Println(v219, v3)
                                    _ = v219
                                    _ = v220
                                }
                            default:
                                v1 = (float32(0.0) <= 1e10)
                            }
                            v1 = ((!v1) || v1)
                        }
                    }
                    if false {
                        switch fi45 {
                        case -2147483648:
                            var v221 uint16 = 411
                            _ = v221
                        case 126:
                            v1 = v1
                        case 256:
                            if (32767 != genFn0(695, false)) {
                                v20 = append(v20, fi45)
                            }
                        default:
                            v1 = v1
                        }
                        v222 := GenS24{F0: ((float32((float64((float32(1e10) - v3)) - 1e10)) - (0.1 - v3)) < (3.14159 - 0.1)), F1: fi45, F2: 0}
                        switch (^(fi45)) {
                        case 816:
                            v3 = (v3 + 0.0)
                        case 127:
                            var v223 float64 = float64(v3)
                            _ = v223
                        case 32767:
                            v3 -= float32(1e10)
                        default:
                            v3 -= v3
                        }
                        _ = v222
                    } else {
                        var v224 int8 = 42
                        v3 -= (v3 * v3)
                        _ = v224
                    }
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v225 string = "foo bar"
                        _ = v225
                    case "":
                        fmt.Println(fi45, v1, v1)
                    case "foo bar":
                        var v226 float32 = (float32(0.1) / v3)
                        _ = v226
                    default:
                        if v1 {
                            delete(v2, (fi45 / (((fi45 / ((fi45) | 1))) | 1)))
                            for fi227 := 0; fi227 < 6; fi227++ {
                                v3 -= (v3 - ((v3 + 0.1) / v3))
                                var v228 float64 = (-2.25 - 1.5)
                                _ = fi227
                                _ = v228
                            }
                        }
                    }
                    if v1 {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v3, v3)
                        case "\n\t":
                            v229 := GenS18{F0: -1, F1: 0, F2: 572}
                            _ = v229
                        }
                        var v230 bool = (-1 >= 42)
                        _ = v230
                    }
                }
                if v1 {
                    if v1 {
                        v231 := GenS20{F0: genFn1(42), F1: (0 <= 42), F2: -2147483648}
                        v232 := GenS25{F0: genFn1(502), F1: 42}
                        _ = v231
                        _ = v232
                    } else {
                        fmt.Println(v3)
                        var v233 float64 = 1.5
                        _ = v233
                    }
                } else {
                    var v234 int16 = 475
                    _ = v234
                }
                v3 += 0.1
            } else {
                fmt.Println(fi45, fi45)
                var v235 float64 = ((float64((0.1 * v3)) * -2.25) * 1e10)
                _ = v235
            }
            _ = len(v2)
            _ = fi45
        }
    }
    var v236 uint8 = 236
    switch "a" {
    case "a":
        v1 = true
    case "foo bar":
        _ = len(v2)
    case "":
        v237 := GenS17{F0: 0, F1: ""}
        _ = v237
    }
    v3 = v3
    for fi238 := 0; fi238 < 5; fi238++ {
        fmt.Println(v3, v1, v3)
        if v1 {
            var v239 int64 = 0
            _ = v239
        }
        _ = fi238
    }
    switch "a" {
    case "\n\t":
        v240 := GenS23{F0: "\n\t", F1: 42, F2: 127}
        _ = v240
    case "":
        if ((128 > (v236 | v236)) && (65535 > 0)) {
            v236 = (v236 + v236)
        } else {
            v1 = v1
            for fi241 := 0; fi241 < 18; fi241++ {
                if v1 {
                    v1 = (2147483647 != -452)
                }
                _ = fi241
            }
            v242 := GenS26{F0: 0, F1: 248, F2: 2}
            _ = v242
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v236:", v236)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v236
    _ = v3
    _ = v40
}
