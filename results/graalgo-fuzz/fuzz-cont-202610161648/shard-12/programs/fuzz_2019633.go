package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int64; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8; F1 int; F2 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint64; F1 int32; F2 bool }
type GenS13 struct { F0 int16; F1 uint; F2 int16 }
type GenS14 struct { F0 uint64; F1 int16; F2 int64 }
type GenS15 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS16 struct { F0 bool; F1 uint32; F2 int }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int32; F1 bool }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 bool }
func genFn0(p0 string, p1 uint, p2 bool) int64 {
    if (!(0 <= -1)) {
        var v7 float64 = 1e10
        v8 := GenS1{F0: 593, F1: 580}
        var v9 bool = false
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        p2 = true
        var v11 int32 = 1
        _ = fi10
        _ = v11
    }
    fmt.Println(p0, p1, p0)
    for fi12 := 0; fi12 < 3; fi12++ {
        p2 = (true && true)
        switch p0 {
        case "a":
            v13 := []int{}
            _ = v13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        switch (fi14 * fi14) {
        case 822:
            if false {
                fmt.Println(p0, p2)
            }
        case 128:
            var v15 int64 = 0
            _ = v15
        }
        v16 := []int{-9223372036854775808, fi14, 4294967295, 127}
        _ = fi14
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 bool = ("a" == "foo bar")
    v2 := make(map[string]string)
    var v3 uint32 = 4294967295
    var v4 int8 = 42
    var v5 int = 32767
    v6 := make(map[int]bool)
    var v17 int64 = -9223372036854775808
    switch -1 {
    case 128:
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS1{F0: 0, F1: (^(v3))}
            _ = fi18
            _ = v19
        }
    case -1:
        if (!v1) {
            for fi20 := 0; fi20 < 8; fi20++ {
                switch (-(v5)) {
                case 9223372036854775807:
                    if ((("a" + "") + (("a" + "foo bar") + "")) != "\n\t") {
                        v2["\n\t"] = "\n\t"
                        fmt.Println(v5)
                    } else {
                        if false {
                            v1 = (v1 && (0 < 1))
                            if (v1 || true) {
                                v1 = (fi20 == 32768)
                                var v21 string = "a"
                                switch v21 {
                                case "\n\t":
                                    v22 := GenS1{F0: 916, F1: 65535}
                                    _ = v22
                                case "hello":
                                    var v23 bool = (1e10 < ((1e10 / 0.1) * 1.5))
                                    _ = v23
                                case "foo bar":
                                    _ = len(v6)
                                default:
                                    if true {
                                        for fi24 := 0; fi24 < 6; fi24++ {
                                            var v25 int32 = 2147483647
                                            if (!(true || v1)) {
                                                v3 += v3
                                                switch v21 {
                                                case "hello":
                                                    v25 = -2147483648
                                                case "foo bar":
                                                    v26 := GenS1{F0: 255, F1: 0}
                                                    _ = v26
                                                }
                                                var v27 bool = false
                                                _ = v27
                                            }
                                            _ = fi24
                                            _ = v25
                                        }
                                        v28 := GenS2{F0: 761}
                                        var v29 bool = (v17 != (v17 / ((v17) | 1)))
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                _ = v21
                            } else {
                                if (!(((!v1) && v1) && (-32768 >= 1000))) {
                                    v30 := GenS3{F0: (false || false), F1: 123}
                                    _ = len(v6)
                                    _ = v30
                                } else {
                                    if (((1e10 + float32(3.14159)) / ((((0.1 + ((1.5 - (-2.25 + -2.25)) + 0.0)) + 0.0) / ((0.1 + (1.5 - 1e10)) / -2.25)) - 0.0)) < 3.14159) {
                                        _ = len(v2)
                                        v6[(v5 - fi20)] = (!v1)
                                    } else {
                                        for fi31 := 0; fi31 < 6; fi31++ {
                                            if v1 {
                                                v32 := GenS1{F0: 32767, F1: v3}
                                                switch "hello" {
                                                case "foo bar":
                                                    v4 -= (-(v4))
                                                case "":
                                                    v33 := GenS4{F0: (("hello" + "foo bar") + "\n\t"), F1: 593}
                                                    _ = v33
                                                case "hello":
                                                    if (v1 || (v1 || (!v1))) {
                                                        fmt.Println(v3)
                                                        var v34 string = "foo bar"
                                                        var v35 uint = 1
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        fmt.Println(v4)
                                                        if ((v17 + genFn0("a", 42, v1)) == genFn0(("\n\t" + "foo bar"), 18446744073709551615, (!(1 != -128)))) {
                                                            v36 := []int{(^(v5)), (fi31 ^ (v5 * (^(fi20))))}
                                                            v37 := v6[fi31]
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        if (!((v1 && (((-2.25 + 0.0) - 3.14159) == -2.25)) || false)) {
                                                            fmt.Println(v17)
                                                            var v38 int16 = 1000
                                                            _ = v38
                                                        } else {
                                                            v39 := GenS5{F0: v17, F1: -2147483648, F2: v3}
                                                            v40 := GenS5{F0: v17, F1: 28, F2: v3}
                                                            var v41 uint16 = 42
                                                            var v42 float32 = (3.14159 / 1.0)
                                                            switch fi20 {
                                                            case 2147483647:
                                                                v43 := []int{}
                                                                _ = v43
                                                            case -9223372036854775808:
                                                                v44 := []int{v5, (v5 + 127), fi31}
                                                                _ = v44
                                                            case 32768:
                                                                var v45 float32 = 3.14159
                                                                _ = v45
                                                            }
                                                            fmt.Println(v17, v5, v5)
                                                            fmt.Println(v3, fi31, v1)
                                                            fmt.Println(v17)
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    }
                                                }
                                                v17 = (^(v17))
                                                _ = v32
                                            }
                                            v5 += fi20
                                            fmt.Println(fi20, v1, v3)
                                            _ = fi31
                                        }
                                        if (536 != 18446744073709551615) {
                                            v1 = ((true || (32767 >= 42)) && false)
                                            v3 = (v3 << ((897) & 63))
                                        }
                                        if true {
                                            var v46 string = ""
                                            for fi47 := 0; fi47 < 7; fi47++ {
                                                fmt.Println(v46, v5)
                                                _ = fi47
                                            }
                                            var v48 float32 = 0.1
                                            _ = v46
                                            _ = v48
                                        } else {
                                            var v49 int64 = int64(v3)
                                            var v50 int8 = 114
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    if ((!v1) || (2 <= 311)) {
                                        fmt.Println(v17, v1)
                                        fmt.Println(v5, fi20, v1)
                                        if ((!false) || (39 != 604)) {
                                            v6[(^(fi20))] = v1
                                            var v51 uint64 = 87
                                            _ = v51
                                        }
                                    } else {
                                        var v52 bool = ("" != ("hello" + "foo bar"))
                                        for fi53 := 0; fi53 < 13; fi53++ {
                                            v52 = v52
                                            v3 -= v3
                                            v3 += 322
                                            _ = fi53
                                        }
                                        v54 := GenS6{F0: 131}
                                        _ = v52
                                        _ = v54
                                    }
                                }
                                for fi55 := 0; fi55 < 13; fi55++ {
                                    if false {
                                        v56 := v6[-1]
                                        _ = v56
                                    } else {
                                        var v57 uint64 = 2
                                        delete(v2, "foo bar")
                                        _ = v57
                                    }
                                    var v58 string = ("\n\t" + ("" + (("" + "") + "a")))
                                    v5 = (v5 & (v5 * v5))
                                    _ = fi55
                                    _ = v58
                                }
                                delete(v6, (v5 + v5))
                            }
                            fmt.Println(fi20, fi20, v5)
                        }
                    }
                case 127:
                    switch v5 {
                    case 482:
                        var v59 int8 = (v4 - (v4 >> ((1) & 63)))
                        _ = v59
                    case 127:
                        for fi60 := 0; fi60 < 10; fi60++ {
                            var v61 bool = (1 <= 667)
                            _ = len(v2)
                            fmt.Println(v4)
                            _ = fi60
                            _ = v61
                        }
                    }
                case 2147483647:
                    v17 += v17
                }
                fmt.Println(fi20, v4)
                _ = fi20
            }
            switch -9223372036854775808 {
            case -30:
                v4 = v4
            case 847:
                v17 = genFn0("a", 18446744073709551615, true)
            }
        }
    case 508:
        for fi62 := 0; fi62 < 14; fi62++ {
            v63 := GenS5{F0: v17, F1: 0, F2: (v3 % (((v3 << ((0) & 63))) | 1))}
            _ = fi62
            _ = v63
        }
    }
    if ((v17 / (((-(v17))) | 1)) == -9223372036854775808) {
        fmt.Println(v17)
        for fi64 := 0; fi64 < 12; fi64++ {
            if true {
                if (!(65535 < 255)) {
                    v4 = 2
                    v4 += v4
                }
                if (v17 == 9223372036854775807) {
                    for fi65 := 0; fi65 < 16; fi65++ {
                        var v66 float64 = 0.0
                        v3 -= 0
                        var v67 uint32 = (v3 & (v3 ^ (v3 / (((v3 / (((v3 % ((64) | 1))) | 1))) | 1))))
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                }
                fmt.Println(v5, v1)
            }
            _ = fi64
        }
        switch v5 {
        case 128:
            v3 = v3
        default:
            if true {
                fmt.Println(v1, v4)
            } else {
                for fi68 := 0; fi68 < 6; fi68++ {
                    var v69 int16 = -237
                    _ = fi68
                    _ = v69
                }
                var v70 uint64 = 1
                var v71 uint16 = 32767
                _ = v70
                _ = v71
            }
        }
    }
    v72 := []int{}
    switch "foo bar" {
    case "a":
        fmt.Println(v5, v3)
    default:
        fmt.Println(v4, v17)
    }
    v6[(v5 ^ v5)] = v1
    for fi73 := 0; fi73 < 5; fi73++ {
        if true {
            var v74 uint16 = 2
            v75 := GenS7{F0: 18446744073709551615}
            v76 := GenS2{F0: 1000}
            _ = v74
            _ = v75
            _ = v76
        }
        for fi77 := 0; fi77 < 10; fi77++ {
            var v78 float64 = float64(float32((1e10 / 1.0)))
            v78 = (v78 - (3.14159 * ((((v78 / v78) + 1e10) / v78) + v78)))
            v17 += v17
            if len(v72) >= 2 {
                v72 = v72[1:len(v72)]
            }
            v2["hello"] = "\n\t"
            _ = fi77
            _ = v78
        }
        _ = fi73
    }
    v79 := GenS2{F0: 32767}
    v6[128] = v1
    var v80 uint = 65535
    var v81 uint8 = 2
    switch -1 {
    case 42:
        var v82 uint8 = (v81 << ((v80) & 63))
        _ = v82
    case 0:
        if (9223372036854775807 >= (v17 * (v17 % ((v17) | 1)))) {
            var v83 bool = true
            _ = v83
        } else {
            v80 += v80
        }
    }
    if true {
        for fi84 := 0; fi84 < 8; fi84++ {
            switch 2147483647 {
            case -369:
                switch v5 {
                case -1:
                    for fi85 := 0; fi85 < 8; fi85++ {
                        if (v1 || v1) {
                            v86 := GenS2{F0: 29}
                            v1 = (32767 != -120)
                            switch fi85 {
                            case 42:
                                delete(v2, ("foo bar" + ("\n\t" + "\n\t")))
                            default:
                                v87 := GenS8{F0: v5}
                                _ = v87
                            }
                            _ = v86
                        }
                        _ = fi85
                    }
                case 42:
                    var v88 float64 = 3.14159
                    _ = v88
                case 127:
                    if false {
                        switch fi84 {
                        case -9223372036854775808:
                            v72 = append(v72, v5)
                        }
                        if (!(!(v1 || (1 < (v4 >> (((v80 ^ 65535)) & 63)))))) {
                            _ = len(v72)
                            var v89 bool = false
                            for fi90 := 0; fi90 < 20; fi90++ {
                                if (-114 <= v4) {
                                    var v91 bool = true
                                    switch (("" + "a") + (("" + "a") + "")) {
                                    case "a":
                                        switch (fi84 % ((fi90) | 1)) {
                                        case 1:
                                            var v92 int8 = v4
                                            _ = v92
                                        case -605:
                                            var v93 uint8 = v81
                                            _ = v93
                                        case -9223372036854775808:
                                            if ((!v89) && ("foo bar" == "\n\t")) {
                                                if (!false) {
                                                    v94 := GenS9{F0: 62, F1: v5, F2: 2}
                                                    v17 += -1
                                                    _ = v94
                                                } else {
                                                    if (758 <= 0) {
                                                        v81 += v81
                                                        if ((v4 <= v4) && (fi90 < v5)) {
                                                            v80 = v80
                                                            _ = len(v6)
                                                        } else {
                                                            switch int(v80) {
                                                            case 4294967295:
                                                                fmt.Println(v17, v4, fi84)
                                                            case 32767:
                                                                if v89 {
                                                                    switch 0 {
                                                                    case 2147483647:
                                                                        v72 = append(v72, (fi90 + fi84))
                                                                    case 171:
                                                                        if (1 < 2) {
                                                                            fmt.Println(v3, v1)
                                                                            var v95 int = (fi90 + fi84)
                                                                            if (!false) {
                                                                                v96 := GenS3{F0: v91, F1: 42}
                                                                                _ = v96
                                                                            } else {
                                                                                if v89 {
                                                                                    if len(v72) > 0 {
                                                                                        v72[0] = (v95 >> ((372) & 63))
                                                                                    }
                                                                                    v97 := v6[-525]
                                                                                    v3 -= v3
                                                                                    _ = v97
                                                                                } else {
                                                                                    var v98 float32 = (((float32(((0.0 + 0.0) + ((1e10 * 1e10) * -2.25))) * -2.25) / 1.5) - 1.5)
                                                                                    _ = v98
                                                                                }
                                                                                if false {
                                                                                    v99 := GenS10{F0: ((127 <= 2) || (v89 || false))}
                                                                                    fmt.Println(v4, v89, v5)
                                                                                    fmt.Println(v91, v91, v80)
                                                                                    _ = v99
                                                                                }
                                                                            }
                                                                            v81 -= v81
                                                                            v6[(fi84 + v95)] = v91
                                                                            v100 := GenS4{F0: ("foo bar" + ("\n\t" + ("" + "a"))), F1: 454}
                                                                            _ = v100
                                                                            _ = v95
                                                                        } else {
                                                                            var v101 uint8 = 127
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                if (42 > 0) {
                                                                                    v102 := GenS11{F0: -2147483648}
                                                                                    var v103 bool = (!v89)
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                }
                                                                            default:
                                                                                v3 = 384
                                                                            }
                                                                            if v1 {
                                                                                var v104 uint64 = 880
                                                                                v105 := v2["\n\t"]
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                            _ = v101
                                                                        }
                                                                    case 738:
                                                                        v6[127] = (v89 || v91)
                                                                    default:
                                                                        switch (("a" + ("\n\t" + "foo bar")) + "\n\t") {
                                                                        case "\n\t":
                                                                            var v106 int32 = 42
                                                                            _ = v106
                                                                        case "foo bar":
                                                                            var v107 uint64 = uint64(v80)
                                                                            _ = v107
                                                                        case "hello":
                                                                            switch v5 {
                                                                            case 127:
                                                                                var v108 uint8 = (v81 % (((v81 - 151)) | 1))
                                                                                _ = v108
                                                                            case -9223372036854775808:
                                                                                var v109 uint = 4294967295
                                                                                _ = v109
                                                                            }
                                                                        default:
                                                                            _ = len(v2)
                                                                        }
                                                                    }
                                                                    v17 -= (v17 >> ((18446744073709551615) & 63))
                                                                } else {
                                                                    fmt.Println(v5, fi84, v4)
                                                                    v110 := GenS2{F0: 1000}
                                                                    v4 = v4
                                                                    _ = v110
                                                                }
                                                            case 926:
                                                                if len(v72) >= 2 {
                                                                    v72 = v72[1:len(v72)]
                                                                }
                                                            default:
                                                                switch "hello" {
                                                                case "":
                                                                    if len(v72) > 0 {
                                                                        v72[0] = v5
                                                                    }
                                                                case "foo bar":
                                                                    var v111 int64 = genFn0(("a" + ("hello" + ("" + "\n\t"))), v80, v89)
                                                                    _ = v111
                                                                default:
                                                                    v81 += 173
                                                                }
                                                            }
                                                            switch ("hello" + ("foo bar" + "a")) {
                                                            case "":
                                                                switch ("a" + ("" + "foo bar")) {
                                                                case "foo bar":
                                                                    v112 := GenS6{F0: (v81 & uint8(v3))}
                                                                    var v113 uint8 = v81
                                                                    var v114 int32 = 2147483647
                                                                    _ = v112
                                                                    _ = v113
                                                                    _ = v114
                                                                default:
                                                                    var v115 string = "\n\t"
                                                                    _ = v115
                                                                }
                                                            case "a":
                                                                v81 -= (v81 & v81)
                                                            }
                                                            _ = len(v6)
                                                            var v116 int16 = 235
                                                            if true {
                                                                if (3.14159 < 1e10) {
                                                                    fmt.Println(v89, v3, v5)
                                                                    switch -520 {
                                                                    case -1:
                                                                        var v117 int64 = v17
                                                                        _ = v117
                                                                    default:
                                                                        v81 += 1
                                                                    }
                                                                    v89 = true
                                                                } else {
                                                                    var v118 string = ""
                                                                    var v119 uint8 = (v81 - v81)
                                                                    _ = len(v2)
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                            } else {
                                                                var v120 int = (-(v5))
                                                                _ = v120
                                                            }
                                                            var v121 int = fi84
                                                            switch (v5 ^ v5) {
                                                            case 4294967295:
                                                                v122 := GenS4{F0: "foo bar", F1: 32767}
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if v89 {
                                                                        if (!("\n\t" != "foo bar")) {
                                                                            if (!(v89 && (-128 == -128))) {
                                                                                var v123 uint32 = (^(v3))
                                                                                _ = v123
                                                                            }
                                                                            v124 := GenS9{F0: (v4 - -103), F1: fi84, F2: 878}
                                                                            fmt.Println(v91, v4, v17)
                                                                            switch (v121 << ((v80) & 63)) {
                                                                            case 127:
                                                                                v125 := GenS12{F0: 588, F1: 2147483647, F2: (!true)}
                                                                                _ = v125
                                                                            }
                                                                            v3 = (v3 - v3)
                                                                            if v89 {
                                                                                var v126 int64 = int64(v3)
                                                                                v3 = uint32(v80)
                                                                                v1 = v1
                                                                                _ = v126
                                                                            } else {
                                                                                v3 = (v3 / ((v3) | 1))
                                                                                if (!(!v91)) {
                                                                                    v80 = 18446744073709551615
                                                                                } else {
                                                                                    var v127 int = fi84
                                                                                    v128 := GenS13{F0: (v116 & v116), F1: 65535, F2: v116}
                                                                                    if (0 != 799) {
                                                                                        v129 := GenS14{F0: 614, F1: (^(v116)), F2: v17}
                                                                                        _ = v129
                                                                                    } else {
                                                                                        v130 := GenS2{F0: (v116 + (v116 - 893))}
                                                                                        _ = v130
                                                                                    }
                                                                                    _ = v127
                                                                                    _ = v128
                                                                                }
                                                                                fmt.Println(v89)
                                                                                switch 614 {
                                                                                case 32768:
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        v1 = v91
                                                                                    }
                                                                                case 210:
                                                                                    fmt.Println(v89, v81, v5)
                                                                                case 32767:
                                                                                    v72 = append(v72, v5)
                                                                                default:
                                                                                    var v131 uint = 65535
                                                                                    _ = v131
                                                                                }
                                                                            }
                                                                            _ = v124
                                                                        } else {
                                                                            v3 = 4294967295
                                                                            _ = len(v6)
                                                                            var v132 int64 = 42
                                                                            switch -9223372036854775808 {
                                                                            case -166:
                                                                                v133 := GenS15{F0: (v3 | v3), F1: 0, F2: 1}
                                                                                _ = v133
                                                                            }
                                                                            var v134 int8 = (v4 - (v4 | v4))
                                                                            v135 := GenS7{F0: uint64(v81)}
                                                                            _ = v132
                                                                            _ = v134
                                                                            _ = v135
                                                                        }
                                                                        var v136 uint = (v80 - (v80 / ((v80) | 1)))
                                                                        fmt.Println(v5)
                                                                        v137 := GenS13{F0: (-(v116)), F1: (v136 | (^(v80))), F2: (v116 * v116)}
                                                                        _ = v136
                                                                        _ = v137
                                                                    } else {
                                                                        if v89 {
                                                                            var v138 uint = (v80 & (v80 % (((v80 << (((v80 | (v80 & v80))) & 63))) | 1)))
                                                                            fmt.Println(v81)
                                                                            _ = v138
                                                                        }
                                                                        v91 = true
                                                                        v139 := GenS16{F0: v89, F1: (v3 | v3), F2: v121}
                                                                        _ = v139
                                                                    }
                                                                }
                                                                var v140 uint16 = 42
                                                                _ = v122
                                                                _ = v140
                                                            case -2147483648:
                                                                v141 := GenS17{F0: (v81 % ((v81) | 1))}
                                                                _ = v141
                                                            default:
                                                                fmt.Println(v4, v4, v5)
                                                            }
                                                            _ = v116
                                                            _ = v121
                                                        }
                                                    }
                                                }
                                                v4 = v4
                                            } else {
                                                fmt.Println(v4, v1, v3)
                                            }
                                        default:
                                            v5 += 32768
                                        }
                                    default:
                                        switch ("a" + "a") {
                                        case "a":
                                            fmt.Println(v5)
                                        default:
                                            v6[fi90] = (!v89)
                                        }
                                    }
                                    _ = v91
                                }
                                v5 = (-(fi84))
                                v142 := GenS9{F0: (v4 << (((v80 + (v80 * (v80 & v80)))) & 63)), F1: (v5 / ((int(v81)) | 1)), F2: 139}
                                _ = fi90
                                _ = v142
                            }
                            _ = v89
                        }
                        v143 := GenS18{F0: -606, F1: false}
                        _ = v143
                    }
                default:
                    switch "a" {
                    case "foo bar":
                        for fi144 := 0; fi144 < 12; fi144++ {
                            v145 := GenS17{F0: 255}
                            _ = fi144
                            _ = v145
                        }
                    }
                }
            case -623:
                v5 = (fi84 >> ((uint(v3)) & 63))
            case -1:
                var v146 int64 = (v17 | (-(v17)))
                _ = v146
            }
            for fi147 := 0; fi147 < 2; fi147++ {
                switch v5 {
                case 32767:
                    if v1 {
                        var v148 float32 = 1.5
                        _ = v148
                    } else {
                        v149 := GenS12{F0: 975, F1: -2147483648, F2: true}
                        v4 += (v4 << ((v80) & 63))
                        if v1 {
                            if len(v72) > 0 {
                                v72[0] = (-(fi84))
                            }
                            var v150 bool = (1 > -1)
                            switch "" {
                            case "a":
                                _ = len(v72)
                            case "foo bar":
                                v151 := GenS13{F0: 1, F1: uint(v81), F2: 1}
                                _ = v151
                            case "":
                                switch "" {
                                case "hello":
                                    v150 = true
                                case "\n\t":
                                    v152 := GenS19{F0: (-763 > 0)}
                                    _ = v152
                                case "a":
                                    v81 += (v81 ^ 2)
                                }
                            }
                            v81 = 135
                            v3 += v3
                            _ = v150
                        } else {
                            var v153 int32 = -2147483648
                            switch (fi84 / ((v5) | 1)) {
                            case 1:
                                fmt.Println(v17)
                            case 9223372036854775807:
                                var v154 int8 = (^(v4))
                                _ = v154
                            case 4294967295:
                                fmt.Println(v3, fi84, v81)
                            }
                            v155 := v6[(v5 + -957)]
                            _ = v153
                            _ = v155
                        }
                        _ = v149
                    }
                    v156 := GenS4{F0: "a", F1: 65535}
                    fmt.Println(v81, v1, fi84)
                    v81 -= (^(v81))
                    delete(v6, (^(fi84)))
                    _ = v156
                default:
                    if false {
                        var v157 float32 = (0.1 - float32(-2.25))
                        _ = v157
                    }
                }
                var v158 int32 = -1
                var v159 int32 = (v158 >> (((v80 + 4294967295)) & 63))
                v160 := GenS20{F0: v1}
                var v161 float64 = (((3.14159 / (3.14159 / 1.5)) / -2.25) * 3.14159)
                _ = fi147
                _ = v158
                _ = v159
                _ = v160
                _ = v161
            }
            v17 = 9223372036854775807
            _ = fi84
        }
    }
    v6[v5] = (v3 != v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v79
    _ = v80
    _ = v81
}
