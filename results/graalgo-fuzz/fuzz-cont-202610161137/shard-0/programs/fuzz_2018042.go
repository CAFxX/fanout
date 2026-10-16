package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 string; F1 uint8; F2 int }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint8; F1 int8; F2 int64 }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 int; F1 uint8; F2 uint64 }
type GenS12 struct { F0 int32; F1 uint32 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int32; F1 uint }
type GenS15 struct { F0 int64; F1 int8; F2 int8 }
func genFn0(p0 int) int8 {
    var v4 int16 = 673
    if (347 < 0) {
        v5 := GenS1{F0: 25, F1: -456, F2: 0}
        var v6 uint64 = 536
        _ = v5
        _ = v6
    }
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 int32, p1 uint, p2 uint16) uint {
    if true {
        p2 = 1
    }
    p0 = (^(p0))
    var v7 uint = uint(p2)
    fmt.Println(p2, p2, p2)
    var v8 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(v7))
}

func main() {
    var v1 uint8 = 127
    v2 := []int{9223372036854775807, 128, 0, 465}
    v3 := make(map[string]int)
    if ((1 > 2147483647) || (1 >= 286)) {
        v1 += (v1 / ((v1) | 1))
        fmt.Println(v1)
    } else {
        var v9 float64 = (((-2.25 * ((1.5 + (0.0 * (0.0 - 3.14159))) - float64((((1.5 + float32(float64(float32(0.0)))) + 0.1) / 1.5)))) / 1e10) / 1e10)
        var v10 int8 = -1
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: "foo bar"}
    if len(v2) > 0 {
        v2[0] = 42
    }
    if true {
        for fi12 := 0; fi12 < 15; fi12++ {
            v1 -= 0
            v13 := GenS3{F0: ""}
            for fi14 := 0; fi14 < 19; fi14++ {
                if ((42 < 9223372036854775807) && (937 >= 42)) {
                    v15 := GenS4{F0: (v1 >> ((genFn1(2147483647, 65535, 97)) & 63)), F1: -1}
                    _ = v15
                } else {
                    v1 = v1
                    fmt.Println(fi12, fi14)
                }
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        if false {
            var v16 int16 = -1
            v16 += (v16 >> ((65535) & 63))
            fmt.Println(v16)
            _ = v16
        }
    } else {
        var v17 string = ("a" + (("\n\t" + "a") + "a"))
        if false {
            delete(v3, ((((((v17 + v17) + v17) + ((v17 + v17) + "foo bar")) + v17) + v17) + v17))
            if (42 <= 9223372036854775807) {
                fmt.Println(v1)
                if false {
                    v17 = "hello"
                    switch ("\n\t" + v17) {
                    case "a":
                        for fi18 := 0; fi18 < 2; fi18++ {
                            var v19 uint = 4294967295
                            v1 += v1
                            _ = fi18
                            _ = v19
                        }
                    case "\n\t":
                        v20 := v3["foo bar"]
                        _ = v20
                    }
                }
            } else {
                for fi21 := 0; fi21 < 13; fi21++ {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        if ((!((-9223372036854775808 > fi21) || ((float32(0.1) - 1e10) == 0.1))) && true) {
                            fmt.Println(fi22, fi21)
                        } else {
                            fmt.Println(fi21, v17)
                            v1 -= v1
                            var v23 string = v17
                            _ = v23
                        }
                        switch "a" {
                        case "\n\t":
                            v1 -= (v1 ^ v1)
                        case "foo bar":
                            v1 -= v1
                        case "a":
                            v1 = (v1 % ((v1) | 1))
                        default:
                            v24 := GenS4{F0: (v1 / (((v1 ^ 2)) | 1)), F1: 127}
                            _ = v24
                        }
                        var v25 uint64 = 222
                        if (v17 < (((v17 + ("hello" + ((v17 + ((v17 + v17) + v17)) + ""))) + v17) + v17)) {
                            v25 = v25
                            v26 := GenS3{F0: v17}
                            if false {
                                var v27 uint16 = 353
                                _ = v27
                            } else {
                                v1 = (v1 & (v1 % ((79) | 1)))
                                var v28 int8 = 42
                                _ = v28
                            }
                            _ = v26
                        }
                        _ = fi22
                        _ = v25
                    }
                    _ = len(v2)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi21
                }
            }
        } else {
            var v29 float64 = 3.14159
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if ((255 == 32767) && true) {
                if (true || (2147483647 == 32768)) {
                    var v30 uint16 = 255
                    _ = v30
                }
                for fi31 := 0; fi31 < 14; fi31++ {
                    v2 = append(v2, (^(fi31)))
                    v29 = 0.1
                    if false {
                        switch fi31 {
                        case 4294967295:
                            var v32 uint = 42
                            _ = v32
                        case -148:
                            switch "a" {
                            case "":
                                for fi33 := 0; fi33 < 13; fi33++ {
                                    switch ("a" + ((v17 + v17) + "\n\t")) {
                                    case "a":
                                        switch v17 {
                                        case "hello":
                                            fmt.Println(v1, v17, v17)
                                            var v34 uint16 = 533
                                            _ = v34
                                        }
                                    case "":
                                        if false {
                                            v17 = v17
                                            if len(v2) > 0 {
                                                v2[0] = fi33
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = (fi33 - fi33)
                                            }
                                        } else {
                                            var v35 int = (fi31 >> ((42) & 63))
                                            _ = v35
                                        }
                                    default:
                                        var v36 uint64 = 0
                                        _ = v36
                                    }
                                    _ = fi33
                                }
                            case "a":
                                if false {
                                    var v37 uint = 4294967295
                                    v37 = v37
                                    _ = v37
                                } else {
                                    for fi38 := 0; fi38 < 4; fi38++ {
                                        fmt.Println(v17)
                                        v39 := GenS4{F0: 47, F1: 127}
                                        _ = fi38
                                        _ = v39
                                    }
                                }
                            default:
                                for fi40 := 0; fi40 < 11; fi40++ {
                                    delete(v3, v17)
                                    if (!true) {
                                        _ = len(v3)
                                        delete(v3, v17)
                                    } else {
                                        v29 -= v29
                                    }
                                    _ = fi40
                                }
                            }
                        case 42:
                            if (0 <= -463) {
                                if (true && (true && (!false))) {
                                    _ = len(v2)
                                    v29 = v29
                                }
                            }
                        }
                    } else {
                        v1 = (v1 * (v1 | v1))
                    }
                    _ = fi31
                }
                v41 := GenS5{F0: "a", F1: 912}
                _ = v41
            }
            _ = v29
        }
        _ = v17
    }
    if true {
        if ("a" == "\n\t") {
            for fi42 := 0; fi42 < 3; fi42++ {
                for fi43 := 0; fi43 < 2; fi43++ {
                    if true {
                        _ = len(v3)
                    }
                    v1 -= (^(v1))
                    _ = fi43
                }
                switch ("a" + "") {
                case "foo bar":
                    switch "a" {
                    case "":
                        var v44 bool = (2147483648 >= 625)
                        _ = v44
                    case "a":
                        var v45 int16 = -1
                        _ = v45
                    case "hello":
                        var v46 int = (fi42 & (fi42 + (fi42 % ((fi42) | 1))))
                        _ = v46
                    }
                case "a":
                    v47 := GenS5{F0: "\n\t", F1: 32767}
                    _ = v47
                case "":
                    delete(v3, "foo bar")
                default:
                    var v48 int64 = 9223372036854775807
                    _ = v48
                }
                _ = fi42
            }
            var v49 int8 = 107
            if (!false) {
                delete(v3, "")
                switch 989 {
                case 9223372036854775807:
                    for fi50 := 0; fi50 < 6; fi50++ {
                        fmt.Println(v1, v1, fi50)
                        _ = fi50
                    }
                case 32768:
                    for fi51 := 0; fi51 < 19; fi51++ {
                        if (true && ((42 >= 715) || false)) {
                            var v52 uint32 = 2
                            switch "foo bar" {
                            case "\n\t":
                                v1 = v1
                            case "hello":
                                v53 := GenS3{F0: ""}
                                _ = v53
                            default:
                                if true {
                                    var v54 int64 = 645
                                    _ = v54
                                } else {
                                    if (255 > 276) {
                                        v55 := GenS5{F0: "a", F1: -32768}
                                        v49 = (v49 / ((v49) | 1))
                                        fmt.Println(v52, v49, v49)
                                        _ = v55
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = (fi51 | (fi51 << ((2) & 63)))
                                        }
                                        v56 := GenS6{F0: (("\n\t" + (("\n\t" + ("" + ("hello" + (((("\n\t" + (("\n\t" + ("" + "\n\t")) + "\n\t")) + "hello") + ("hello" + "a")) + "")))) + "")) + "foo bar"), F1: 0, F2: fi51}
                                        _ = v56
                                    }
                                    for fi57 := 0; fi57 < 17; fi57++ {
                                        v1 = (v1 - (^(v1)))
                                        if (3.14159 > 1.5) {
                                            var v58 int32 = 2147483647
                                            var v59 uint = 1
                                            var v60 uint8 = (v1 >> (((^(v59))) & 63))
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        } else {
                                            v49 += -7
                                        }
                                        _ = fi57
                                    }
                                    fmt.Println(v1, fi51)
                                }
                            }
                            _ = v52
                        } else {
                            v1 = (^(v1))
                        }
                        for fi61 := 0; fi61 < 9; fi61++ {
                            v49 -= 68
                            v49 = 42
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                            _ = fi61
                        }
                        _ = fi51
                    }
                case 2147483647:
                    if (true || false) {
                        if (false || true) {
                            switch -257 {
                            case -216:
                                switch 484 {
                                case 1:
                                    fmt.Println(v49)
                                case 42:
                                    if (true || true) {
                                        for fi62 := 0; fi62 < 18; fi62++ {
                                            v1 += 0
                                            v3["hello"] = fi62
                                            _ = fi62
                                        }
                                        if ((^(v1)) <= (^(v1))) {
                                            var v63 float32 = (1.5 + -2.25)
                                            if (42 <= 42) {
                                                v63 += float32(1.5)
                                                var v64 bool = (false && (genFn0(127) == (v49 + 42)))
                                                _ = v64
                                            }
                                            _ = v63
                                        } else {
                                            var v65 bool = ((false && (!false)) || (!false))
                                            _ = v65
                                        }
                                    }
                                case -852:
                                    v66 := GenS7{F0: 65535, F1: 2}
                                    _ = v66
                                }
                            case -9223372036854775808:
                                for fi67 := 0; fi67 < 14; fi67++ {
                                    v49 += -30
                                    _ = fi67
                                }
                            case 9223372036854775807:
                                for fi68 := 0; fi68 < 16; fi68++ {
                                    v1 -= v1
                                    _ = fi68
                                }
                            default:
                                var v69 float64 = 0.1
                                _ = v69
                            }
                        } else {
                            if (-748 < 32767) {
                                var v70 int8 = 3
                                fmt.Println(v1, v1, v1)
                                _ = v70
                            } else {
                                switch "\n\t" {
                                case "a":
                                    var v71 int = 32767
                                    _ = v71
                                case "hello":
                                    v2 = append(v2, -138)
                                default:
                                    var v72 float64 = 3.14159
                                    _ = v72
                                }
                            }
                        }
                        switch 32767 {
                        case -2147483648:
                            if true {
                                _ = len(v3)
                                fmt.Println(v49, v1)
                            } else {
                                v73 := GenS7{F0: 42, F1: 1}
                                _ = v73
                            }
                        case -54:
                            v74 := GenS1{F0: 0, F1: 42, F2: (v49 * -1)}
                            _ = v74
                        case 42:
                            switch "foo bar" {
                            case "":
                                for fi75 := 0; fi75 < 15; fi75++ {
                                    fmt.Println(v1, v49)
                                    var v76 uint = 2
                                    for fi77 := 0; fi77 < 4; fi77++ {
                                        var v78 uint = v76
                                        fmt.Println(v1, fi77, fi75)
                                        _ = fi77
                                        _ = v78
                                    }
                                    _ = fi75
                                    _ = v76
                                }
                            case "hello":
                                if false {
                                    v79 := v3["\n\t"]
                                    switch ("hello" + "hello") {
                                    case "":
                                        v80 := GenS8{F0: 18446744073709551615, F1: 0}
                                        _ = v80
                                    case "foo bar":
                                        v1 += 82
                                    }
                                    _ = v79
                                }
                            default:
                                v1 += v1
                            }
                        default:
                            if (float32(float64(0.1)) <= 1.5) {
                                switch "hello" {
                                case "a":
                                    var v81 int8 = genFn0(128)
                                    _ = v81
                                case "hello":
                                    var v82 uint16 = 526
                                    _ = v82
                                case "foo bar":
                                    v1 += v1
                                }
                            }
                        }
                    } else {
                        if true {
                            v83 := v3["hello"]
                            v1 = (^(v1))
                            _ = v83
                        }
                        for fi84 := 0; fi84 < 9; fi84++ {
                            for fi85 := 0; fi85 < 10; fi85++ {
                                v1 += (v1 & v1)
                                switch 1 {
                                case 32768:
                                    switch (^(fi84)) {
                                    case 0:
                                        if false {
                                            var v86 int32 = 2147483647
                                            var v87 uint64 = 269
                                            if len(v2) > 0 {
                                                v2[0] = fi84
                                            }
                                            _ = v86
                                            _ = v87
                                        } else {
                                            v88 := GenS9{F0: 1, F1: v49, F2: -1}
                                            if false {
                                                if true {
                                                    var v89 int8 = v49
                                                    v90 := GenS10{F0: 9223372036854775807, F1: v89}
                                                    _ = v89
                                                    _ = v90
                                                }
                                                v1 = (v1 % ((42) | 1))
                                            } else {
                                                v91 := GenS11{F0: fi85, F1: (v1 * v1), F2: 740}
                                                switch -848 {
                                                case 143:
                                                    fmt.Println(fi85)
                                                default:
                                                    if len(v2) > 0 {
                                                        v2[0] = fi85
                                                    }
                                                }
                                                _ = v91
                                            }
                                            switch ("hello" + "hello") {
                                            case "foo bar":
                                                v1 = v1
                                            }
                                            v1 += 2
                                            v92 := GenS3{F0: "foo bar"}
                                            if (!true) {
                                                var v93 int64 = -947
                                                _ = v93
                                            } else {
                                                v1 = (v1 % ((v1) | 1))
                                                var v94 int64 = -1
                                                _ = v94
                                            }
                                            _ = v88
                                            _ = v92
                                        }
                                    case -1:
                                        v1 = v1
                                    }
                                case -312:
                                    fmt.Println(fi85)
                                default:
                                    if (!true) {
                                        var v95 float64 = float64(1.5)
                                        v95 = v95
                                        switch (^(fi85)) {
                                        case -2147483648:
                                            v2 = append(v2, (fi85 << ((18446744073709551615) & 63)))
                                        default:
                                            var v96 int64 = 384
                                            _ = v96
                                        }
                                        v49 += genFn0(2147483647)
                                        v2 = append(v2, 127)
                                        if true {
                                            if (!(float32(float64(((float32((((v95 + ((v95 * v95) * float64(0.1))) * (v95 / (((v95 - float64(((float32(0.1) - -2.25) / -2.25))) + (0.1 * float64((-2.25 + 3.14159)))) - v95))) + 1.5)) / 0.1) * -2.25))) > 1e10)) {
                                                var v97 float32 = 1.5
                                                var v98 uint64 = 4294967295
                                                _ = v97
                                                _ = v98
                                            } else {
                                                v99 := GenS7{F0: 65535, F1: 4294967295}
                                                var v100 float32 = 1.5
                                                v49 = (v49 << ((1) & 63))
                                                _ = v100
                                                _ = v99
                                            }
                                        } else {
                                            v49 = genFn0(fi84)
                                            if ((v49 ^ int8(fi85)) > v49) {
                                                v101 := v3[""]
                                                _ = v101
                                            } else {
                                                switch (fi84 | fi84) {
                                                case 602:
                                                    v102 := GenS12{F0: -527, F1: 2}
                                                    _ = v102
                                                case -180:
                                                    fmt.Println(fi85, fi85, fi84)
                                                case 337:
                                                    if (false && true) {
                                                        v3[(("hello" + (("foo bar" + "a") + (("a" + "a") + "a"))) + "foo bar")] = (^(fi85))
                                                        var v103 float32 = (3.14159 / 0.1)
                                                        v104 := GenS13{F0: -2147483648}
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                }
                                                v1 -= v1
                                            }
                                            v95 += v95
                                        }
                                        var v105 uint64 = 167
                                        _ = v105
                                        _ = v95
                                    }
                                }
                                v106 := GenS14{F0: 879, F1: 18446744073709551615}
                                _ = fi85
                                _ = v106
                            }
                            _ = fi84
                        }
                        fmt.Println(v1)
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = -657
                    }
                }
                switch ("\n\t" + "a") {
                case "foo bar":
                    v49 += v49
                }
            } else {
                v49 += 1
                v49 -= 0
            }
            _ = v49
        } else {
            var v107 uint32 = 629
            _ = v107
        }
    } else {
        if true {
            _ = len(v2)
            _ = len(v3)
        } else {
            if len(v2) > 0 {
                v2[0] = 1
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v1)
            case "a":
                if true {
                    for fi108 := 0; fi108 < 3; fi108++ {
                        var v109 string = ("a" + "foo bar")
                        if false {
                            v110 := GenS3{F0: v109}
                            _ = v110
                        } else {
                            var v111 bool = (((v1 % ((v1) | 1)) == (v1 * v1)) && (636 < 86))
                            if v111 {
                                v111 = ((fi108 + fi108) < (^(fi108)))
                                for fi112 := 0; fi112 < 7; fi112++ {
                                    v109 = v109
                                    _ = fi112
                                }
                                v109 = ("hello" + ("a" + v109))
                            }
                            _ = v111
                        }
                        _ = fi108
                        _ = v109
                    }
                    v1 = v1
                    var v113 float32 = float32(((0.0 * -2.25) + 1e10))
                    _ = v113
                } else {
                    var v114 int64 = 0
                    v114 -= (v114 << ((2) & 63))
                    _ = v114
                }
            case "\n\t":
                for fi115 := 0; fi115 < 5; fi115++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v116 float64 = 1.5
                        _ = v116
                    case "":
                        if (65535 != 0) {
                            for fi117 := 0; fi117 < 20; fi117++ {
                                v1 -= v1
                                var v118 int32 = 688
                                _ = fi117
                                _ = v118
                            }
                            fmt.Println(fi115)
                        }
                    }
                    if false {
                        if (v1 == v1) {
                            switch (-(fi115)) {
                            case 32767:
                                var v119 float64 = 1e10
                                _ = v119
                            }
                            if (18446744073709551615 < 481) {
                                fmt.Println(fi115)
                            }
                            if false {
                                var v120 int = fi115
                                var v121 bool = (!(-522 <= -716))
                                for fi122 := 0; fi122 < 4; fi122++ {
                                    v123 := v3[("\n\t" + "\n\t")]
                                    _ = fi122
                                    _ = v123
                                }
                                _ = v120
                                _ = v121
                            }
                        }
                    }
                    v1 -= 164
                    _ = fi115
                }
            }
        }
    }
    var v124 uint16 = 182
    v124 = 32767
    var v125 uint64 = 18446744073709551615
    v1 -= 26
    if false {
        v126 := GenS7{F0: 2147483647, F1: 893}
        v127 := GenS15{F0: 42, F1: -8, F2: 42}
        if false {
            v124 += 65535
        }
        _ = v126
        _ = v127
    }
    if false {
        var v128 uint = 389
        v125 = (v125 * v125)
        v125 -= (v125 + (v125 << (((v128 >> (((v128 | (^(v128)))) & 63))) & 63)))
        _ = v128
    }
    for fi129 := 0; fi129 < 20; fi129++ {
        var v130 float32 = 1.5
        v131 := v3[""]
        _ = fi129
        _ = v130
        _ = v131
    }
    var v132 uint64 = (^(v125))
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v132:", v132)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v124
    _ = v125
    _ = v132
    _ = v2
    _ = v3
}
