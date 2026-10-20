package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 string; F1 uint; F2 int64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint16; F1 int }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 bool; F1 int32 }
type GenS15 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS16 struct { F0 int8; F1 uint64 }
type GenS17 struct { F0 uint32; F1 int16; F2 string }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS20 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS21 struct { F0 int32; F1 string }
type GenS22 struct { F0 uint64; F1 string }
type GenS23 struct { F0 int8; F1 uint8 }
type GenS24 struct { F0 int32; F1 int8 }
type GenS25 struct { F0 int32; F1 int; F2 int }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 bool; F1 string; F2 uint32 }
type GenS28 struct { F0 uint; F1 uint8; F2 uint }
type GenS29 struct { F0 uint32 }
type GenS30 struct { F0 string }

func main() {
    var v1 uint16 = 65535
    var v2 uint16 = v1
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 uint64 = 18446744073709551615
            _ = fi5
            _ = v6
        }
        v2 = v2
        _ = fi4
    }
    fmt.Println(v2, v1)
    if (0 >= -1) {
        switch 722 {
        case 4294967295:
            if ((v2 >> ((2) & 63)) == 656) {
                for fi7 := 0; fi7 < 1; fi7++ {
                    if ("" != (("\n\t" + (("foo bar" + "hello") + ("a" + ("hello" + "a")))) + "\n\t")) {
                        v8 := []int{}
                        var v9 bool = true
                        var v10 int64 = 1
                        _ = v10
                        _ = v8
                        _ = v9
                    } else {
                        var v11 uint32 = 0
                        fmt.Println(v11)
                        _ = len(v3)
                        _ = v11
                    }
                    fmt.Println(fi7)
                    v12 := GenS1{F0: 42, F1: (("\n\t" + "\n\t") + "")}
                    _ = fi7
                    _ = v12
                }
                switch 9223372036854775807 {
                case -2147483648:
                    v2 = (v1 * (v2 - v1))
                default:
                    if (((((0.0 + 3.14159) * -2.25) * 0.1) >= 3.14159) && (!true)) {
                        v13 := []int{32768, -889}
                        v1 += 762
                        _ = v13
                    } else {
                        v14 := GenS2{F0: 127}
                        _ = v14
                    }
                }
                switch ("foo bar" + "a") {
                case "":
                    fmt.Println(v1, v2)
                case "hello":
                    fmt.Println(v1, v1)
                }
            }
        case -748:
            v15 := []int{32768, 127}
            _ = v15
        case 0:
            fmt.Println(v2, v1)
        default:
            v16 := []int{127}
            _ = v16
        }
        v17 := GenS1{F0: 4294967295, F1: (("hello" + ("" + "hello")) + ("\n\t" + "a"))}
        switch -2147483648 {
        case 752:
            for fi18 := 0; fi18 < 18; fi18++ {
                v1 += 255
                if false {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v20 := GenS1{F0: 896, F1: (("" + ("foo bar" + "hello")) + "\n\t")}
                        v1 = v2
                        switch ("hello" + (("\n\t" + "hello") + "a")) {
                        case "foo bar":
                            v21 := []int{fi19}
                            _ = v21
                        case "\n\t":
                            if ((1.5 * -2.25) > float32(3.14159)) {
                                var v22 uint64 = 42
                                v23 := []int{(fi19 ^ (fi19 ^ 680)), -1, 169}
                                if true {
                                    if true {
                                        fmt.Println(fi18, fi18, fi19)
                                    }
                                    if false {
                                        var v24 int16 = 32767
                                        switch 32768 {
                                        case -1:
                                            if (((v24 | v24) < v24) || true) {
                                                if (!true) {
                                                    var v25 uint = 1
                                                    _ = v25
                                                }
                                                var v26 int = fi18
                                                _ = v26
                                            }
                                        case 9223372036854775807:
                                            var v27 float32 = (1e10 / 1e10)
                                            var v28 int64 = -518
                                            _ = v27
                                            _ = v28
                                        case 128:
                                            fmt.Println(v22, v1, fi18)
                                        }
                                        fmt.Println(fi18)
                                        _ = v24
                                    } else {
                                        v1 = 255
                                    }
                                    if false {
                                        if false {
                                            v29 := GenS3{F0: 65535, F1: fi18, F2: -439}
                                            v30 := GenS2{F0: 42}
                                            _ = v29
                                            _ = v30
                                        }
                                        var v31 int16 = -1
                                        _ = v31
                                    } else {
                                        v1 = v1
                                    }
                                    v2 = 391
                                } else {
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    if (true || false) {
                                        fmt.Println(v1, v22, v1)
                                        var v32 uint = 42
                                        _ = v32
                                    }
                                }
                                var v33 uint = 255
                                v3["\n\t"] = "\n\t"
                                _ = v22
                                _ = v23
                                _ = v33
                            }
                        default:
                            v1 -= 32767
                        }
                        var v34 uint8 = 128
                        var v35 string = "foo bar"
                        _ = fi19
                        _ = v20
                        _ = v34
                        _ = v35
                    }
                } else {
                    if ((fi18 & fi18) < fi18) {
                        v36 := []int{}
                        switch 4294967295 {
                        case 128:
                            for fi37 := 0; fi37 < 2; fi37++ {
                                var v38 uint64 = 923
                                var v39 float32 = float32(float64((0.1 + 0.1)))
                                _ = len(v36)
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        case 1:
                            fmt.Println(v1)
                        }
                        v40 := GenS2{F0: 122}
                        _ = v36
                        _ = v40
                    } else {
                        v41 := GenS2{F0: 42}
                        var v42 float32 = 0.0
                        v43 := GenS2{F0: 127}
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    if (!(false || true)) {
                        for fi44 := 0; fi44 < 18; fi44++ {
                            if true {
                                switch "a" {
                                case "\n\t":
                                    v45 := []int{(fi18 * (fi18 % ((fi18) | 1))), (fi18 << ((2) & 63)), -133, (fi18 - (fi18 + fi18))}
                                    _ = v45
                                case "a":
                                    v46 := GenS4{F0: 2}
                                    _ = v46
                                }
                            }
                            v47 := GenS4{F0: 824}
                            v3["hello"] = "a"
                            _ = fi44
                            _ = v47
                        }
                    } else {
                        var v48 float64 = 0.0
                        switch fi18 {
                        case 127:
                            for fi49 := 0; fi49 < 10; fi49++ {
                                fmt.Println(v1, fi49)
                                _ = fi49
                            }
                        default:
                            v50 := GenS4{F0: 2147483648}
                            _ = v50
                        }
                        var v51 int32 = 1
                        _ = v48
                        _ = v51
                    }
                    switch ("hello" + "") {
                    case "\n\t":
                        for fi52 := 0; fi52 < 5; fi52++ {
                            if (0.0 == 1.5) {
                                var v53 int = fi52
                                var v54 int64 = -9223372036854775808
                                _ = v53
                                _ = v54
                            }
                            v2 -= (v1 & 408)
                            var v55 uint32 = 2147483648
                            _ = fi52
                            _ = v55
                        }
                    case "":
                        v56 := []int{fi18, (fi18 ^ 42)}
                        _ = v56
                    default:
                        for fi57 := 0; fi57 < 1; fi57++ {
                            v1 = v2
                            v2 = (^(v1))
                            delete(v3, "\n\t")
                            var v58 int64 = -150
                            v59 := []int{(fi57 - fi18)}
                            var v60 int32 = 2147483647
                            v61 := GenS1{F0: 413, F1: ("a" + "\n\t")}
                            v62 := GenS5{F0: (v58 / (((v58 ^ v58)) | 1)), F1: false}
                            if false {
                                if true {
                                    var v63 uint32 = 2147483647
                                    _ = v63
                                } else {
                                    v64 := GenS5{F0: (-(v58)), F1: (0 < 0)}
                                    v1 += v2
                                    v65 := GenS3{F0: 2, F1: -9223372036854775808, F2: (v58 - 1)}
                                    _ = v64
                                    _ = v65
                                }
                            }
                            if (!((!false) || (!(false && false)))) {
                                v58 += 0
                                if len(v59) > 0 {
                                    v59[0] = fi57
                                }
                            }
                            v59 = append(v59, fi18)
                            if (!(42 == 0)) {
                                if (v60 != (v60 % ((v60) | 1))) {
                                    if false {
                                        v2 += v2
                                        if false {
                                            v3["foo bar"] = "foo bar"
                                            _ = len(v59)
                                            var v66 int = fi18
                                            v60 += v60
                                            if true {
                                                if false {
                                                    v67 := GenS3{F0: 42, F1: fi18, F2: v58}
                                                    _ = v67
                                                } else {
                                                    v1 = (v2 * (v1 | v1))
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v68 float64 = (0.1 / (1.5 * (-2.25 / 0.1)))
                                                        _ = v68
                                                    default:
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v69 uint64 = 19
                                                            _ = v69
                                                        case "a":
                                                            var v70 uint16 = v1
                                                            _ = v70
                                                        }
                                                    }
                                                    v71 := GenS3{F0: 339, F1: fi57, F2: v58}
                                                    _ = v71
                                                }
                                                if false {
                                                    var v72 uint = 1
                                                    _ = v72
                                                } else {
                                                    switch "" {
                                                    case "hello":
                                                        if true {
                                                            if true {
                                                                if (false && true) {
                                                                    if false {
                                                                        switch fi18 {
                                                                        case -9223372036854775808:
                                                                            var v73 float32 = float32((-2.25 + ((0.1 + 0.1) + 1e10)))
                                                                            _ = v73
                                                                        case 128:
                                                                            switch (("foo bar" + "\n\t") + ("\n\t" + "foo bar")) {
                                                                            case "":
                                                                                switch 32767 {
                                                                                case 128:
                                                                                    var v74 int = v66
                                                                                    _ = v74
                                                                                case 127:
                                                                                    v60 += 2147483647
                                                                                }
                                                                            case "a":
                                                                                switch fi57 {
                                                                                case -392:
                                                                                    var v75 uint16 = v2
                                                                                    _ = v75
                                                                                case -904:
                                                                                    if true {
                                                                                        fmt.Println(fi18, v58, v2)
                                                                                    } else {
                                                                                        v60 -= (-(v60))
                                                                                        v2 = v1
                                                                                    }
                                                                                default:
                                                                                    var v76 uint64 = 4294967295
                                                                                    _ = v76
                                                                                }
                                                                            default:
                                                                                if (!(!(!(false || ((fi18 << ((18446744073709551615) & 63)) <= (v66 - fi57)))))) {
                                                                                    v2 += v1
                                                                                    _ = len(v59)
                                                                                    var v77 uint = 255
                                                                                    _ = v77
                                                                                }
                                                                            }
                                                                        default:
                                                                            v60 += (v60 * -1)
                                                                        }
                                                                        var v78 int = fi18
                                                                        _ = v78
                                                                    }
                                                                    if len(v59) > 0 {
                                                                        v59[0] = fi18
                                                                    }
                                                                    fmt.Println(v58, v58, v2)
                                                                }
                                                                if (true || true) {
                                                                    var v79 uint16 = (v2 | v1)
                                                                    if false {
                                                                        if (3.14159 >= 3.14159) {
                                                                            if false {
                                                                                v80 := GenS5{F0: (v58 >> ((42) & 63)), F1: false}
                                                                                switch (fi57 * (fi18 - -1)) {
                                                                                case 128:
                                                                                    var v81 uint = 1
                                                                                    _ = v81
                                                                                case 4294967295:
                                                                                    v82 := GenS6{F0: 1, F1: 42}
                                                                                    _ = v82
                                                                                case 32768:
                                                                                    v79 += (v79 % (((v2 + v79)) | 1))
                                                                                }
                                                                                var v83 uint16 = (v2 - v79)
                                                                                v59 = append(v59, (fi57 >> ((636) & 63)))
                                                                                if false {
                                                                                    if len(v59) > 0 {
                                                                                        v59[0] = (fi18 - 1)
                                                                                    }
                                                                                    if (fi18 >= -238) {
                                                                                        if false {
                                                                                            if true {
                                                                                                v58 = (v58 & -320)
                                                                                                switch ("foo bar" + "") {
                                                                                                case "\n\t":
                                                                                                    var v84 float64 = 0.0
                                                                                                    _ = v84
                                                                                                default:
                                                                                                    var v85 uint32 = 2147483647
                                                                                                    _ = v85
                                                                                                }
                                                                                            } else {
                                                                                                var v86 int64 = (v58 - (v58 / ((-9223372036854775808) | 1)))
                                                                                                _ = v86
                                                                                            }
                                                                                        } else {
                                                                                            switch "foo bar" {
                                                                                            case "hello":
                                                                                                v2 -= 32767
                                                                                                v87 := GenS7{F0: (("a" + "\n\t") + "foo bar"), F1: uint(v79), F2: 779}
                                                                                                _ = v87
                                                                                            case "":
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    v88 := GenS2{F0: -128}
                                                                                                    _ = v88
                                                                                                case "":
                                                                                                    fmt.Println(v58, v58, v66)
                                                                                                default:
                                                                                                    v89 := v3["a"]
                                                                                                    _ = v89
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v90 float64 = -2.25
                                                                                                _ = v90
                                                                                            }
                                                                                            var v91 uint64 = 4294967295
                                                                                            _ = v91
                                                                                        }
                                                                                    }
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        if (65535 >= 65535) {
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                if len(v59) > 0 {
                                                                                                    v59[0] = fi57
                                                                                                }
                                                                                            case "hello":
                                                                                                v3["foo bar"] = "a"
                                                                                            default:
                                                                                                fmt.Println(v60, v58)
                                                                                            }
                                                                                            var v92 float64 = float64(1e10)
                                                                                            var v93 uint32 = 2
                                                                                            _ = v92
                                                                                            _ = v93
                                                                                        } else {
                                                                                            fmt.Println(v60, v58)
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                fmt.Println(v79, v66, v58)
                                                                                            case "a":
                                                                                                v94 := GenS4{F0: 156}
                                                                                                _ = v94
                                                                                            case "foo bar":
                                                                                                var v95 float64 = -2.25
                                                                                                _ = v95
                                                                                            }
                                                                                            var v96 float32 = float32(-2.25)
                                                                                            _ = v96
                                                                                        }
                                                                                    case "foo bar":
                                                                                        fmt.Println(v79)
                                                                                    case "\n\t":
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            v97 := GenS3{F0: 352, F1: fi18, F2: v58}
                                                                                            _ = v97
                                                                                        case "hello":
                                                                                            var v98 int16 = 0
                                                                                            _ = v98
                                                                                        case "a":
                                                                                            v60 = v60
                                                                                        default:
                                                                                            v2 = 0
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v80
                                                                                _ = v83
                                                                            }
                                                                            v66 -= int(v2)
                                                                            var v99 int64 = (^(v58))
                                                                            _ = v99
                                                                        } else {
                                                                            v2 = v2
                                                                            fmt.Println(v1, v79)
                                                                            var v100 int64 = v58
                                                                            _ = v100
                                                                        }
                                                                    }
                                                                    v79 -= 42
                                                                    if true {
                                                                        var v101 float64 = -2.25
                                                                        _ = v101
                                                                    }
                                                                    _ = v79
                                                                }
                                                                v66 += fi57
                                                            } else {
                                                                switch v66 {
                                                                case 32767:
                                                                    if false {
                                                                        v1 += (^(v1))
                                                                        var v102 int = fi18
                                                                        var v103 uint32 = 668
                                                                        _ = v102
                                                                        _ = v103
                                                                    } else {
                                                                        switch v66 {
                                                                        case 9223372036854775807:
                                                                            if false {
                                                                                var v104 float32 = 0.1
                                                                                var v105 uint16 = (v2 / ((1) | 1))
                                                                                var v106 float32 = float32((0.0 - 1.5))
                                                                                _ = v104
                                                                                _ = v105
                                                                                _ = v106
                                                                            } else {
                                                                                if false {
                                                                                    var v107 float64 = 0.0
                                                                                    _ = v107
                                                                                } else {
                                                                                    var v108 uint32 = 719
                                                                                    if len(v59) > 0 {
                                                                                        v59[0] = (v66 ^ -81)
                                                                                    }
                                                                                    _ = v108
                                                                                }
                                                                            }
                                                                        case 619:
                                                                            v109 := GenS2{F0: 100}
                                                                            _ = v109
                                                                        case -9223372036854775808:
                                                                            v66 -= (-(fi18))
                                                                        }
                                                                    }
                                                                case 0:
                                                                    if true {
                                                                        var v110 int8 = -128
                                                                        v58 = (v58 - (v58 ^ (v58 << ((999) & 63))))
                                                                        _ = v110
                                                                    } else {
                                                                        if (1.5 < 1e10) {
                                                                            switch 9223372036854775807 {
                                                                            case 128:
                                                                                delete(v3, "\n\t")
                                                                            case 1:
                                                                                if len(v59) > 0 {
                                                                                    v59[0] = (fi57 + fi57)
                                                                                }
                                                                            case 42:
                                                                                var v111 float32 = float32(((1.5 / (0.1 / 3.14159)) * -2.25))
                                                                                _ = v111
                                                                            }
                                                                            var v112 uint32 = 2147483648
                                                                            var v113 uint = 65535
                                                                            _ = v112
                                                                            _ = v113
                                                                        } else {
                                                                            if (!false) {
                                                                                if false {
                                                                                    v58 = (^(v58))
                                                                                } else {
                                                                                    if true {
                                                                                        if len(v59) >= 2 {
                                                                                            v59 = v59[1:len(v59)]
                                                                                        }
                                                                                        v114 := GenS8{F0: 125}
                                                                                        _ = v114
                                                                                    } else {
                                                                                        var v115 int32 = (v60 + 1)
                                                                                        _ = v115
                                                                                    }
                                                                                }
                                                                                _ = len(v3)
                                                                                switch 128 {
                                                                                case 32767:
                                                                                    var v116 uint8 = 128
                                                                                    _ = v116
                                                                                case 435:
                                                                                    _ = len(v3)
                                                                                case -2147483648:
                                                                                    v2 = (v2 | v1)
                                                                                default:
                                                                                    if false {
                                                                                        delete(v3, ("" + "hello"))
                                                                                        var v117 uint16 = 893
                                                                                        _ = v117
                                                                                    } else {
                                                                                        _ = len(v59)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if false {
                                                                                    if true {
                                                                                        fmt.Println(v60, v1, v60)
                                                                                        v66 -= (v66 << ((65535) & 63))
                                                                                    }
                                                                                    var v118 float64 = float64(3.14159)
                                                                                    _ = v118
                                                                                } else {
                                                                                    var v119 uint = 42
                                                                                    v1 = 65535
                                                                                    _ = v119
                                                                                }
                                                                            }
                                                                            if ((v2 << ((255) & 63)) != (v2 << ((42) & 63))) {
                                                                                v120 := GenS7{F0: "foo bar", F1: 118, F2: v58}
                                                                                v121 := GenS6{F0: 71, F1: -522}
                                                                                _ = v120
                                                                                _ = v121
                                                                            }
                                                                            if true {
                                                                                v1 = 32767
                                                                                v3["a"] = "a"
                                                                                _ = len(v59)
                                                                                var v122 uint8 = 127
                                                                                _ = v122
                                                                            }
                                                                            v123 := GenS8{F0: uint(v2)}
                                                                            _ = v123
                                                                        }
                                                                        v124 := GenS9{F0: 42, F1: 42}
                                                                        _ = v124
                                                                    }
                                                                case -166:
                                                                    var v125 uint64 = 137
                                                                    _ = v125
                                                                default:
                                                                    var v126 uint32 = 2147483648
                                                                    _ = v126
                                                                }
                                                            }
                                                        }
                                                    case "a":
                                                        var v127 int16 = -722
                                                        _ = v127
                                                    case "foo bar":
                                                        v3["\n\t"] = "hello"
                                                    default:
                                                        v59 = append(v59, -2147483648)
                                                    }
                                                    fmt.Println(v66)
                                                }
                                            } else {
                                                if false {
                                                    var v128 int16 = 0
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v58 -= int64(v1)
                                                    case "\n\t":
                                                        v66 = fi18
                                                    case "hello":
                                                        v60 = v60
                                                    }
                                                    _ = v128
                                                }
                                                v66 = (-(fi18))
                                            }
                                            _ = v66
                                        }
                                        fmt.Println(v60, v2)
                                    }
                                    v58 = (v58 / ((v58) | 1))
                                } else {
                                    fmt.Println(fi18, fi18, fi18)
                                    v58 = (v58 - (^(v58)))
                                }
                                fmt.Println(v2, fi18)
                            }
                            _ = fi57
                            _ = v58
                            _ = v59
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                    }
                }
                switch fi18 {
                case 4294967295:
                    v129 := GenS5{F0: 403, F1: (1 >= 0)}
                    _ = v129
                }
                _ = fi18
            }
        case 639:
            fmt.Println(v1, v1)
        }
        _ = v17
    }
    switch "foo bar" {
    case "hello":
        if (true || (true || false)) {
            if true {
                fmt.Println(v1, v1)
                switch ("hello" + ("hello" + (("foo bar" + "foo bar") + ("hello" + "foo bar")))) {
                case "a":
                    switch (("" + "hello") + ("" + "foo bar")) {
                    case "\n\t":
                        v130 := []int{42, -2147483648}
                        _ = v130
                    case "foo bar":
                        for fi131 := 0; fi131 < 12; fi131++ {
                            v132 := GenS8{F0: 4294967295}
                            v2 = (v2 >> ((255) & 63))
                            fmt.Println(v2, v2)
                            _ = fi131
                            _ = v132
                        }
                    case "":
                        for fi133 := 0; fi133 < 15; fi133++ {
                            v134 := GenS10{F0: 2, F1: 1, F2: (v2 + (v2 ^ (^(v1))))}
                            _ = fi133
                            _ = v134
                        }
                    }
                case "\n\t":
                    switch "\n\t" {
                    case "a":
                        var v135 int32 = 2147483647
                        _ = v135
                    case "foo bar":
                        var v136 int8 = -105
                        _ = v136
                    case "hello":
                        var v137 uint = 18446744073709551615
                        _ = v137
                    }
                case "":
                    if (true && false) {
                        v138 := GenS11{F0: 100}
                        if (4294967295 > 772) {
                            if false {
                                switch 42 {
                                case 1:
                                    var v139 int16 = 42
                                    _ = v139
                                case 0:
                                    v1 -= (v1 << ((4294967295) & 63))
                                case 625:
                                    if ((818 < 4294967295) || (-53 > 2147483647)) {
                                        var v140 uint = 0
                                        v141 := GenS7{F0: "", F1: v140, F2: -292}
                                        _ = v140
                                        _ = v141
                                    }
                                default:
                                    if false {
                                        if true {
                                            v142 := GenS11{F0: 0}
                                            if false {
                                                var v143 int8 = 127
                                                if true {
                                                    fmt.Println(v2, v1, v1)
                                                } else {
                                                    var v144 float32 = -2.25
                                                    _ = v144
                                                }
                                                v145 := GenS8{F0: 464}
                                                _ = v143
                                                _ = v145
                                            } else {
                                                if false {
                                                    v1 = (v1 % ((v2) | 1))
                                                    var v146 uint64 = 2
                                                    _ = v146
                                                } else {
                                                    for fi147 := 0; fi147 < 6; fi147++ {
                                                        v1 += (v1 >> ((0) & 63))
                                                        for fi148 := 0; fi148 < 1; fi148++ {
                                                            if false {
                                                                if false {
                                                                    _ = len(v3)
                                                                    switch ("a" + "foo bar") {
                                                                    case "\n\t":
                                                                        v149 := GenS3{F0: uint(v2), F1: fi147, F2: 42}
                                                                        switch fi148 {
                                                                        case -1:
                                                                            if (true || false) {
                                                                                var v150 float32 = 0.0
                                                                                _ = v150
                                                                            }
                                                                        default:
                                                                            if (!(!false)) {
                                                                                _ = len(v3)
                                                                                fmt.Println(fi147, v1, v2)
                                                                                if (1000 <= 42) {
                                                                                    if false {
                                                                                        switch (fi148 >> ((2) & 63)) {
                                                                                        case -640:
                                                                                            var v151 int16 = 156
                                                                                            _ = v151
                                                                                        case 127:
                                                                                            v2 = v1
                                                                                        default:
                                                                                            v152 := GenS8{F0: 42}
                                                                                            _ = v152
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch -100 {
                                                                                    case -204:
                                                                                        v153 := GenS12{F0: v1, F1: (fi148 / ((fi148) | 1))}
                                                                                        _ = v153
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v154 := []int{}
                                                                        _ = v149
                                                                        _ = v154
                                                                    }
                                                                    var v155 int8 = 42
                                                                    switch fi148 {
                                                                    case 128:
                                                                        v155 += v155
                                                                    }
                                                                    _ = v155
                                                                } else {
                                                                    var v156 int32 = 42
                                                                    fmt.Println(v1, fi148, v2)
                                                                    v3["a"] = "hello"
                                                                    _ = v156
                                                                }
                                                                v157 := GenS12{F0: 1, F1: (fi148 / ((0) | 1))}
                                                                var v158 uint32 = 2147483648
                                                                _ = v157
                                                                _ = v158
                                                            } else {
                                                                v159 := GenS3{F0: 4294967295, F1: -745, F2: 257}
                                                                switch 32768 {
                                                                case 4294967295:
                                                                    if (((740 > -365) || true) || (!(((!false) || (255 >= 2)) && (!true)))) {
                                                                        v3[(("\n\t" + "\n\t") + ("a" + ""))] = (("hello" + "foo bar") + (("hello" + "foo bar") + "hello"))
                                                                        if false {
                                                                            var v160 uint = 42
                                                                            v160 = (v160 % ((672) | 1))
                                                                            v2 = v2
                                                                            _ = v160
                                                                        }
                                                                    }
                                                                case 997:
                                                                    fmt.Println(fi147)
                                                                }
                                                                _ = v159
                                                            }
                                                            if (fi148 != (fi147 ^ 32767)) {
                                                                v161 := GenS2{F0: 1}
                                                                var v162 int = fi147
                                                                if false {
                                                                    if (!(uint8(v2) >= 144)) {
                                                                        var v163 uint64 = 42
                                                                        _ = v163
                                                                    } else {
                                                                        var v164 float64 = 1.5
                                                                        _ = v164
                                                                    }
                                                                } else {
                                                                    v162 = (fi147 << ((255) & 63))
                                                                }
                                                                _ = v161
                                                                _ = v162
                                                            }
                                                            var v165 uint16 = v1
                                                            v2 += (^(v2))
                                                            if true {
                                                                v1 -= v165
                                                                v166 := GenS6{F0: 100, F1: -216}
                                                                _ = v166
                                                            } else {
                                                                var v167 string = "hello"
                                                                _ = v167
                                                            }
                                                            v3[""] = (("foo bar" + "") + ("a" + ("hello" + "hello")))
                                                            if (false && (0.0 != -2.25)) {
                                                                fmt.Println(v1)
                                                            } else {
                                                                var v168 float32 = 1e10
                                                                fmt.Println(fi148, fi147)
                                                                _ = v168
                                                            }
                                                            _ = fi148
                                                            _ = v165
                                                        }
                                                        if ((42 <= 46) || (9223372036854775807 >= (^(fi147)))) {
                                                            v1 = 628
                                                        }
                                                        _ = fi147
                                                    }
                                                    var v169 uint8 = 128
                                                    _ = v169
                                                }
                                                v170 := GenS10{F0: 4294967295, F1: 0, F2: v1}
                                                var v171 uint16 = 1
                                                _ = v170
                                                _ = v171
                                            }
                                            var v172 float32 = 1.5
                                            _ = v142
                                            _ = v172
                                        } else {
                                            for fi173 := 0; fi173 < 9; fi173++ {
                                                var v174 int32 = -1
                                                delete(v3, "hello")
                                                _ = fi173
                                                _ = v174
                                            }
                                        }
                                    } else {
                                        v2 -= v1
                                        if (1 != 18446744073709551615) {
                                            var v175 int = 256
                                            v176 := GenS11{F0: 100}
                                            _ = v175
                                            _ = v176
                                        } else {
                                            if (!true) {
                                                var v177 int = -1
                                                switch v177 {
                                                case -1:
                                                    switch v177 {
                                                    case 32767:
                                                        v178 := GenS8{F0: 42}
                                                        _ = v178
                                                    case 4294967295:
                                                        v179 := GenS5{F0: 0, F1: (!false)}
                                                        _ = v179
                                                    default:
                                                        if (4294967295 > 224) {
                                                            if false {
                                                                for fi180 := 0; fi180 < 9; fi180++ {
                                                                    v1 = v2
                                                                    for fi181 := 0; fi181 < 19; fi181++ {
                                                                        if true {
                                                                            v177 = fi181
                                                                        } else {
                                                                            v1 = 985
                                                                            if false {
                                                                                _ = len(v3)
                                                                            } else {
                                                                                if false {
                                                                                    fmt.Println(v177, v1)
                                                                                    if (18446744073709551615 >= 0) {
                                                                                        v2 = v2
                                                                                        v182 := GenS13{F0: 9223372036854775807}
                                                                                        _ = v182
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            if (2147483648 != 2147483648) {
                                                                                                v177 = (^(fi180))
                                                                                            } else {
                                                                                                fmt.Println(fi181)
                                                                                                v183 := []int{}
                                                                                                fmt.Println(fi181, v2, v2)
                                                                                                v177 += fi181
                                                                                                _ = v183
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v184 int64 = 509
                                                                                    _ = v184
                                                                                } else {
                                                                                    v1 -= 264
                                                                                    var v185 string = "foo bar"
                                                                                    fmt.Println(fi180, v185)
                                                                                    _ = v185
                                                                                }
                                                                                fmt.Println(fi180)
                                                                                var v186 uint64 = 18446744073709551615
                                                                                _ = v186
                                                                            }
                                                                            var v187 int = fi180
                                                                            _ = v187
                                                                        }
                                                                        _ = fi181
                                                                    }
                                                                    for fi188 := 0; fi188 < 5; fi188++ {
                                                                        v189 := []int{}
                                                                        if (!true) {
                                                                            if (-128 <= 100) {
                                                                                v2 = (^(v2))
                                                                                var v190 int16 = -630
                                                                                var v191 uint32 = 0
                                                                                _ = v190
                                                                                _ = v191
                                                                            }
                                                                            var v192 int16 = -1
                                                                            _ = v192
                                                                        }
                                                                        _ = fi188
                                                                        _ = v189
                                                                    }
                                                                    _ = fi180
                                                                }
                                                            } else {
                                                                var v193 uint64 = 42
                                                                _ = v193
                                                            }
                                                            v194 := GenS14{F0: true, F1: 2147483647}
                                                            _ = v194
                                                        }
                                                    }
                                                case 9223372036854775807:
                                                    var v195 uint8 = 2
                                                    _ = v195
                                                case 907:
                                                    v196 := []int{(v177 % (((v177 + (^(v177)))) | 1)), 1, v177, v177}
                                                    _ = v196
                                                }
                                                _ = v177
                                            } else {
                                                v197 := GenS9{F0: -128, F1: 1}
                                                if (869 < 4294967295) {
                                                    if (3.14159 > 0.1) {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v198 string = ""
                                                            _ = v198
                                                        case "":
                                                            for fi199 := 0; fi199 < 13; fi199++ {
                                                                switch "a" {
                                                                case "":
                                                                    for fi200 := 0; fi200 < 12; fi200++ {
                                                                        v201 := v3[("hello" + "\n\t")]
                                                                        var v202 string = ("" + v201)
                                                                        _ = fi200
                                                                        _ = v201
                                                                        _ = v202
                                                                    }
                                                                case "hello":
                                                                    var v203 int = (fi199 & 9223372036854775807)
                                                                    _ = v203
                                                                case "a":
                                                                    for fi204 := 0; fi204 < 6; fi204++ {
                                                                        v3["a"] = ("foo bar" + "")
                                                                        _ = fi204
                                                                    }
                                                                }
                                                                v205 := GenS11{F0: 42}
                                                                _ = fi199
                                                                _ = v205
                                                            }
                                                        default:
                                                            v1 += v1
                                                        }
                                                        fmt.Println(v2, v1)
                                                    } else {
                                                        if true {
                                                            switch 0 {
                                                            case 580:
                                                                v2 += v2
                                                            case 4294967295:
                                                                var v206 int8 = 42
                                                                _ = v206
                                                            }
                                                            var v207 uint8 = 135
                                                            _ = v207
                                                        } else {
                                                            var v208 int16 = 0
                                                            _ = v208
                                                        }
                                                    }
                                                    v1 = (^(v1))
                                                    v209 := GenS15{F0: 42, F1: 0, F2: 0}
                                                    _ = v209
                                                } else {
                                                    switch -548 {
                                                    case 32768:
                                                        var v210 string = "\n\t"
                                                        _ = v210
                                                    case 0:
                                                        if (532 > 468) {
                                                            v1 += v2
                                                        } else {
                                                            if ("" == "foo bar") {
                                                                for fi211 := 0; fi211 < 6; fi211++ {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        v212 := GenS10{F0: 1, F1: 2147483647, F2: (v2 ^ 255)}
                                                                        _ = v212
                                                                    case "hello":
                                                                        for fi213 := 0; fi213 < 17; fi213++ {
                                                                            if (-100 == 1) {
                                                                                if (2147483648 >= 41) {
                                                                                    v214 := []int{32767}
                                                                                    _ = v214
                                                                                } else {
                                                                                    var v215 bool = (100 != 59)
                                                                                    var v216 uint32 = 659
                                                                                    var v217 string = "hello"
                                                                                    _ = v215
                                                                                    _ = v216
                                                                                    _ = v217
                                                                                }
                                                                            } else {
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    switch (fi213 / (((fi211 / (((fi211 / (((fi213 | (fi211 >> ((1) & 63)))) | 1))) | 1))) | 1)) {
                                                                                    case 32768:
                                                                                        var v218 uint64 = 883
                                                                                        _ = v218
                                                                                    case 2147483647:
                                                                                        switch (fi213 << ((2) & 63)) {
                                                                                        case 127:
                                                                                            v1 -= 255
                                                                                        case 128:
                                                                                            fmt.Println(v1, v2)
                                                                                        default:
                                                                                            if ("a" != "a") {
                                                                                                fmt.Println(fi211)
                                                                                            } else {
                                                                                                fmt.Println(fi211)
                                                                                                var v219 uint8 = 127
                                                                                                delete(v3, ("a" + (("\n\t" + ("hello" + "")) + "a")))
                                                                                                _ = v219
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v2 = (v1 ^ v2)
                                                                                    }
                                                                                default:
                                                                                    v2 -= 42
                                                                                    switch 32768 {
                                                                                    case 279:
                                                                                        v1 += v1
                                                                                    }
                                                                                    v220 := GenS3{F0: 0, F1: (fi213 / (((fi213 + -1)) | 1)), F2: 9223372036854775807}
                                                                                    _ = v220
                                                                                }
                                                                                var v221 int8 = -128
                                                                                _ = v221
                                                                            }
                                                                            switch ("hello" + "\n\t") {
                                                                            case "hello":
                                                                                if true {
                                                                                    var v222 uint8 = 127
                                                                                    _ = v222
                                                                                } else {
                                                                                    if (42 >= 42) {
                                                                                        v223 := []int{}
                                                                                        _ = v223
                                                                                    } else {
                                                                                        var v224 int16 = 42
                                                                                        _ = v224
                                                                                    }
                                                                                }
                                                                            case "":
                                                                                if (!true) {
                                                                                    switch 2147483647 {
                                                                                    case 32767:
                                                                                        var v225 int64 = 662
                                                                                        _ = v225
                                                                                    }
                                                                                } else {
                                                                                    v226 := GenS16{F0: 127, F1: 1}
                                                                                    _ = v226
                                                                                }
                                                                            default:
                                                                                v2 = v1
                                                                            }
                                                                            _ = fi213
                                                                        }
                                                                    case "":
                                                                        v1 += v2
                                                                    }
                                                                    v1 = v1
                                                                    v2 += 711
                                                                    _ = fi211
                                                                }
                                                                v227 := GenS17{F0: 4294967295, F1: 1000, F2: "\n\t"}
                                                                var v228 int32 = 2147483647
                                                                _ = v227
                                                                _ = v228
                                                            } else {
                                                                v2 += (^(v1))
                                                                var v229 string = "foo bar"
                                                                v1 -= 520
                                                                _ = v229
                                                            }
                                                            var v230 uint8 = 255
                                                            _ = v230
                                                        }
                                                    default:
                                                        v1 = (v1 & 2)
                                                    }
                                                    for fi231 := 0; fi231 < 12; fi231++ {
                                                        var v232 float64 = 1.5
                                                        for fi233 := 0; fi233 < 1; fi233++ {
                                                            if true {
                                                                switch -188 {
                                                                case 32767:
                                                                    var v234 uint16 = v2
                                                                    v232 += v232
                                                                    if (408 >= 11) {
                                                                        if (18446744073709551615 >= 0) {
                                                                            if (!true) {
                                                                                v234 += 32767
                                                                                v235 := GenS18{F0: 255}
                                                                                if false {
                                                                                    var v236 int = fi233
                                                                                    v237 := []int{fi231, 9223372036854775807, (v236 + (fi231 % (((fi233 >> ((18446744073709551615) & 63))) | 1)))}
                                                                                    _ = v236
                                                                                    _ = v237
                                                                                } else {
                                                                                    var v238 string = ""
                                                                                    _ = v238
                                                                                }
                                                                                _ = v235
                                                                            }
                                                                            v1 = 255
                                                                            v239 := GenS19{F0: 531, F1: 18, F2: -689}
                                                                            _ = v239
                                                                        }
                                                                        v240 := GenS7{F0: "hello", F1: 710, F2: 42}
                                                                        v241 := []int{(fi233 * fi231), 32768, fi231, 42}
                                                                        _ = v240
                                                                        _ = v241
                                                                    }
                                                                    _ = v234
                                                                }
                                                            }
                                                            if true {
                                                                fmt.Println(v232, fi233)
                                                                v1 = v2
                                                                fmt.Println(v2, v1)
                                                                switch "a" {
                                                                case "":
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        switch "" {
                                                                        case "":
                                                                            v242 := GenS20{F0: 2, F1: -66, F2: v2}
                                                                            fmt.Println(v2)
                                                                            _ = v242
                                                                        }
                                                                    case "a":
                                                                        var v243 int32 = 2147483647
                                                                        _ = v243
                                                                    }
                                                                case "foo bar":
                                                                    var v244 float32 = (0.0 + -2.25)
                                                                    _ = v244
                                                                case "hello":
                                                                    v1 = v2
                                                                    var v245 int16 = 42
                                                                    var v246 int8 = 42
                                                                    _ = v245
                                                                    _ = v246
                                                                }
                                                            } else {
                                                                switch fi233 {
                                                                case -1:
                                                                    v247 := GenS19{F0: 42, F1: 564, F2: 42}
                                                                    _ = v247
                                                                case 2147483647:
                                                                    switch ((("foo bar" + "") + "") + ("a" + "foo bar")) {
                                                                    case "":
                                                                        switch (fi233 << ((255) & 63)) {
                                                                        case -198:
                                                                            switch "" {
                                                                            case "hello":
                                                                                var v248 bool = (!true)
                                                                                _ = v248
                                                                            case "a":
                                                                                delete(v3, "\n\t")
                                                                            case "foo bar":
                                                                                if true {
                                                                                    v232 -= (v232 / (float64(3.14159) / v232))
                                                                                    var v249 int8 = 100
                                                                                    var v250 int64 = 1
                                                                                    _ = v249
                                                                                    _ = v250
                                                                                } else {
                                                                                    v232 = 1.5
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        var v251 float32 = 3.14159
                                                                                        _ = v251
                                                                                    default:
                                                                                        v252 := []int{(fi231 * fi233)}
                                                                                        _ = v252
                                                                                    }
                                                                                    v1 = (v2 - v2)
                                                                                }
                                                                            default:
                                                                                v3[(("" + "a") + ("a" + ("hello" + ("" + ("\n\t" + "foo bar")))))] = "hello"
                                                                            }
                                                                        case 65:
                                                                            var v253 int = fi231
                                                                            var v254 uint8 = 255
                                                                            switch fi231 {
                                                                            case 20:
                                                                                v255 := GenS5{F0: -9223372036854775808, F1: ((960 == (fi233 | 670)) || (!false))}
                                                                                _ = v255
                                                                            case 0:
                                                                                var v256 uint64 = 0
                                                                                _ = v256
                                                                            }
                                                                            v1 = v1
                                                                            _ = v253
                                                                            _ = v254
                                                                        case 0:
                                                                            var v257 bool = (2 != 4294967295)
                                                                            _ = v257
                                                                        }
                                                                        v258 := GenS6{F0: -1, F1: -32768}
                                                                        fmt.Println(v1, v2)
                                                                        _ = v258
                                                                    case "a":
                                                                        var v259 uint = 223
                                                                        _ = v259
                                                                    }
                                                                    var v260 int = fi231
                                                                    _ = v260
                                                                default:
                                                                    var v261 uint64 = 1
                                                                    var v262 int16 = -545
                                                                    _ = v261
                                                                    _ = v262
                                                                }
                                                            }
                                                            _ = fi233
                                                        }
                                                        _ = fi231
                                                        _ = v232
                                                    }
                                                }
                                                v263 := []int{-313, 995, -654, -2147483648}
                                                _ = v197
                                                _ = v263
                                            }
                                            v264 := []int{}
                                            _ = v264
                                        }
                                        v1 += 2
                                    }
                                }
                                var v265 uint = 2
                                _ = v265
                            } else {
                                if false {
                                    v3["foo bar"] = "foo bar"
                                    v1 = v2
                                    if (255 > 127) {
                                        v3["\n\t"] = "foo bar"
                                        var v266 int64 = 669
                                        if false {
                                            v267 := GenS17{F0: 4294967295, F1: 0, F2: "hello"}
                                            switch 128 {
                                            case 42:
                                                var v268 float64 = float64(float32(0.1))
                                                _ = v268
                                            default:
                                                v269 := GenS18{F0: 18446744073709551615}
                                                _ = v269
                                            }
                                            v270 := []int{}
                                            _ = v267
                                            _ = v270
                                        }
                                        _ = v266
                                    }
                                } else {
                                    if true {
                                        for fi271 := 0; fi271 < 8; fi271++ {
                                            if false {
                                                v1 += (v1 % ((v2) | 1))
                                                v272 := GenS12{F0: v1, F1: -280}
                                                if false {
                                                    delete(v3, ("\n\t" + (("" + "\n\t") + ("a" + (("a" + "a") + ("a" + (((("hello" + "foo bar") + ("hello" + "")) + "hello") + "a")))))))
                                                    v273 := []int{(fi271 ^ 582), fi271}
                                                    var v274 bool = true
                                                    _ = v273
                                                    _ = v274
                                                } else {
                                                    v1 = 122
                                                    v2 = 1
                                                }
                                                _ = v272
                                            } else {
                                                v1 += v2
                                                var v275 int32 = 2147483647
                                                var v276 uint16 = v2
                                                _ = v275
                                                _ = v276
                                            }
                                            _ = fi271
                                        }
                                        v277 := GenS21{F0: 947, F1: (("foo bar" + "hello") + "")}
                                        _ = v277
                                    }
                                    var v278 uint16 = 1
                                    switch ("hello" + "hello") {
                                    case "\n\t":
                                        switch ("" + "hello") {
                                        case "foo bar":
                                            v278 = (^(v278))
                                        case "":
                                            var v279 int = 4294967295
                                            _ = v279
                                        case "a":
                                            if false {
                                                for fi280 := 0; fi280 < 13; fi280++ {
                                                    var v281 uint16 = v2
                                                    _ = fi280
                                                    _ = v281
                                                }
                                                fmt.Println(v2, v278, v278)
                                            }
                                        }
                                    case "a":
                                        v2 = (v2 / ((0) | 1))
                                    case "":
                                        v282 := []int{-1, -893, 594}
                                        _ = v282
                                    }
                                    _ = v278
                                }
                                v283 := GenS17{F0: 65535, F1: -32768, F2: ("\n\t" + "\n\t")}
                                v2 = (v1 | v2)
                                _ = v283
                            }
                            for fi284 := 0; fi284 < 4; fi284++ {
                                switch ("hello" + "a") {
                                case "foo bar":
                                    for fi285 := 0; fi285 < 2; fi285++ {
                                        v2 -= 0
                                        var v286 uint64 = 42
                                        _ = fi285
                                        _ = v286
                                    }
                                case "a":
                                    for fi287 := 0; fi287 < 18; fi287++ {
                                        v1 -= (v2 / ((v2) | 1))
                                        switch (^(fi287)) {
                                        case 746:
                                            switch "" {
                                            case "":
                                                switch fi287 {
                                                case -1:
                                                    switch "hello" {
                                                    case "a":
                                                        var v288 uint8 = 0
                                                        _ = v288
                                                    case "\n\t":
                                                        switch (fi284 / ((fi287) | 1)) {
                                                        case 908:
                                                            fmt.Println(v1, fi284, v1)
                                                        case -2147483648:
                                                            v289 := []int{}
                                                            _ = v289
                                                        case 127:
                                                            v290 := GenS17{F0: 897, F1: 32767, F2: "\n\t"}
                                                            var v291 int8 = 42
                                                            fmt.Println(v291, v2)
                                                            _ = v290
                                                            _ = v291
                                                        }
                                                    }
                                                case -2147483648:
                                                    var v292 int32 = 2147483647
                                                    _ = v292
                                                default:
                                                    fmt.Println(fi287, v1)
                                                }
                                            case "foo bar":
                                                v293 := []int{(fi284 ^ (fi287 - (fi287 >> ((687) & 63)))), fi284, (fi284 % ((-1) | 1))}
                                                _ = v293
                                            default:
                                                if (v1 <= 65535) {
                                                    switch ("hello" + "") {
                                                    case "a":
                                                        fmt.Println(v2, fi287, v1)
                                                    case "foo bar":
                                                        switch (fi284 >> ((4294967295) & 63)) {
                                                        case -621:
                                                            v294 := GenS22{F0: 221, F1: ""}
                                                            _ = v294
                                                        }
                                                    default:
                                                        var v295 uint32 = 65535
                                                        _ = v295
                                                    }
                                                    if false {
                                                        if false {
                                                            var v296 int64 = 1
                                                            v2 = v2
                                                            var v297 string = (("foo bar" + "\n\t") + "")
                                                            _ = v296
                                                            _ = v297
                                                        } else {
                                                            v298 := []int{fi287}
                                                            if len(v298) > 0 {
                                                                v298[0] = fi287
                                                            }
                                                            v299 := GenS23{F0: 42, F1: 174}
                                                            _ = v298
                                                            _ = v299
                                                        }
                                                        v300 := GenS24{F0: 1, F1: -110}
                                                        v2 = v2
                                                        _ = v300
                                                    } else {
                                                        var v301 string = "\n\t"
                                                        v1 = (v2 ^ (v1 + 154))
                                                        var v302 uint = 4294967295
                                                        v303 := GenS25{F0: 1, F1: (fi284 + (-(fi287))), F2: fi284}
                                                        v301 = v301
                                                        _ = v301
                                                        _ = v302
                                                        _ = v303
                                                    }
                                                } else {
                                                    v2 += 42
                                                    v304 := []int{(fi284 - (fi287 | 2147483647)), fi287, (^(fi287)), (fi284 + (fi284 | -700))}
                                                    if len(v304) >= 2 {
                                                        v304 = v304[1:len(v304)]
                                                    }
                                                    _ = v304
                                                }
                                            }
                                        case 127:
                                            v305 := GenS25{F0: 1, F1: fi287, F2: (fi287 / ((fi287) | 1))}
                                            v3[((("hello" + "hello") + "hello") + "foo bar")] = ""
                                            _ = v305
                                        default:
                                            if false {
                                                fmt.Println(v1, fi284, fi284)
                                                var v306 bool = false
                                                fmt.Println(v2, fi284, v1)
                                                _ = v306
                                            }
                                        }
                                        _ = fi287
                                    }
                                case "\n\t":
                                    if (18446744073709551615 >= 1) {
                                        v307 := v3[(("a" + ("" + "hello")) + ("foo bar" + ""))]
                                        v2 += v2
                                        _ = v307
                                    }
                                }
                                v308 := []int{(fi284 >> ((205) & 63))}
                                if ("" < "hello") {
                                    for fi309 := 0; fi309 < 13; fi309++ {
                                        var v310 uint32 = 737
                                        if false {
                                            fmt.Println(fi309)
                                            if (!false) {
                                                v311 := GenS26{F0: (!(408 == 18446744073709551615))}
                                                fmt.Println(v310, fi284, v1)
                                                fmt.Println(v310, v310, v1)
                                                _ = v311
                                            }
                                            var v312 int = fi309
                                            var v313 int = int(v310)
                                            _ = v312
                                            _ = v313
                                        } else {
                                            var v314 int16 = 42
                                            var v315 uint8 = 0
                                            v314 = v314
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v316 string = "\n\t"
                                                _ = v316
                                            case "foo bar":
                                                switch "a" {
                                                case "":
                                                    v1 = (v1 * uint16(v310))
                                                }
                                            case "":
                                                var v317 uint16 = (v1 * 903)
                                                _ = v317
                                            }
                                            v3[""] = ""
                                            var v318 float64 = 1.5
                                            _ = v314
                                            _ = v315
                                            _ = v318
                                        }
                                        v1 += 2
                                        v319 := v3["\n\t"]
                                        v308 = append(v308, 888)
                                        _ = fi309
                                        _ = v310
                                        _ = v319
                                    }
                                    for fi320 := 0; fi320 < 4; fi320++ {
                                        v2 -= v1
                                        if (true || ((v1 < (v2 - 0)) && false)) {
                                            var v321 uint16 = 0
                                            _ = v321
                                        }
                                        var v322 string = ("" + "a")
                                        _ = fi320
                                        _ = v322
                                    }
                                    v323 := v3["a"]
                                    _ = v323
                                }
                                _ = fi284
                                _ = v308
                            }
                            for fi324 := 0; fi324 < 13; fi324++ {
                                if true {
                                    for fi325 := 0; fi325 < 2; fi325++ {
                                        if true {
                                            v326 := []int{}
                                            _ = v326
                                        }
                                        switch (fi324 >> ((18446744073709551615) & 63)) {
                                        case 128:
                                            v327 := GenS27{F0: false, F1: "foo bar", F2: 1}
                                            _ = v327
                                        case 459:
                                            fmt.Println(fi324, fi324, fi325)
                                            var v328 uint = 65535
                                            _ = v328
                                        }
                                        v329 := GenS8{F0: 65535}
                                        _ = fi325
                                        _ = v329
                                    }
                                } else {
                                    if true {
                                        if false {
                                            if (4294967295 >= 1) {
                                                var v330 int64 = 1
                                                _ = v330
                                            } else {
                                                for fi331 := 0; fi331 < 16; fi331++ {
                                                    var v332 uint8 = 192
                                                    _ = fi331
                                                    _ = v332
                                                }
                                            }
                                            v333 := []int{fi324, (fi324 ^ (fi324 ^ (^(fi324)))), 32768}
                                            _ = v333
                                        } else {
                                            v334 := []int{fi324, fi324}
                                            for fi335 := 0; fi335 < 6; fi335++ {
                                                var v336 int32 = 996
                                                v1 = v2
                                                _ = fi335
                                                _ = v336
                                            }
                                            _ = len(v334)
                                            _ = v334
                                        }
                                    }
                                    v337 := GenS6{F0: 100, F1: 159}
                                    if (true || (true && (true && false))) {
                                        fmt.Println(v1, v2, v2)
                                        var v338 uint8 = 127
                                        var v339 int16 = -32768
                                        _ = v338
                                        _ = v339
                                    } else {
                                        v340 := GenS18{F0: 1}
                                        _ = len(v3)
                                        _ = v340
                                    }
                                    _ = v337
                                }
                                var v341 float64 = 0.0
                                if (false || false) {
                                    v342 := []int{}
                                    _ = v342
                                }
                                _ = fi324
                                _ = v341
                            }
                        } else {
                            switch "" {
                            case "a":
                                switch "a" {
                                case "":
                                    var v343 int = 4294967295
                                    _ = v343
                                case "a":
                                    v344 := GenS28{F0: 784, F1: 2, F2: 418}
                                    _ = v344
                                default:
                                    v345 := GenS20{F0: 106, F1: -20, F2: 892}
                                    _ = v345
                                }
                            }
                            switch 127 {
                            case 127:
                                var v346 int = 0
                                _ = v346
                            case 128:
                                if (!false) {
                                    v1 = 2
                                }
                            }
                        }
                        for fi347 := 0; fi347 < 11; fi347++ {
                            var v348 int8 = 1
                            fmt.Println(v1)
                            _ = fi347
                            _ = v348
                        }
                        _ = v138
                    }
                }
                fmt.Println(v1, v2)
            }
            _ = len(v3)
            v2 += 897
        }
    case "a":
        switch 853 {
        case 1:
            v2 += v2
        }
    case "":
        v1 = 42
    }
    v2 += 32767
    v2 = v2
    var v349 string = ("\n\t" + "foo bar")
    v349 = v349
    if (false || false) {
        switch (v349 + v349) {
        case "foo bar":
            for fi350 := 0; fi350 < 10; fi350++ {
                var v351 float32 = float32(3.14159)
                fmt.Println(v2)
                var v352 int64 = 1
                _ = fi350
                _ = v351
                _ = v352
            }
        }
    } else {
        _ = len(v3)
        fmt.Println(v349)
        if false {
            v1 = v2
            for fi353 := 0; fi353 < 20; fi353++ {
                var v354 int16 = 1000
                for fi355 := 0; fi355 < 19; fi355++ {
                    var v356 int32 = 42
                    delete(v3, "\n\t")
                    _ = fi355
                    _ = v356
                }
                _ = fi353
                _ = v354
            }
        }
    }
    var v357 float32 = 1e10
    v358 := GenS29{F0: 65535}
    v357 = (v357 / v357)
    v359 := GenS19{F0: 0, F1: 4294967295, F2: -1}
    v360 := GenS30{F0: (v349 + "a")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v349:", v349)
    fmt.Println("v357:", v357)
    _ = v1
    _ = v2
    _ = v3
    _ = v349
    _ = v357
    _ = v358
    _ = v359
    _ = v360
}
