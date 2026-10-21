package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 bool; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 int8; F2 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 int8; F1 uint16 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS17 struct { F0 int8; F1 int64; F2 int64 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 string; F1 uint64 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 int8; F1 bool; F2 bool }

func main() {
    var v1 bool = (42 <= -128)
    v2 := make(map[string]uint32)
    var v3 int = -9223372036854775808
    v4 := []string{"\n\t"}
    var v5 float64 = float64(0.1)
    var v6 int64 = 234
    var v7 float32 = 3.14159
    fmt.Println(v3)
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(v7, v1)
            v10 := GenS1{F0: v3}
            _ = len(v4)
            _ = fi9
            _ = v10
        }
        if (!v1) {
            if len(v4) > 0 {
                v4[0] = ("a" + "a")
            }
            for fi11 := 0; fi11 < 18; fi11++ {
                v5 += (1.5 + float64(((v7 / 1.5) + v7)))
                if v1 {
                    if (1 < 42) {
                        if (1 > -1) {
                            switch fi11 {
                            case 128:
                                v5 -= 3.14159
                            case 2147483647:
                                var v12 int8 = 42
                                _ = v12
                            case -1:
                                _ = len(v2)
                            }
                            switch (("foo bar" + "") + "") {
                            case "foo bar":
                                if (!v1) {
                                    v13 := GenS1{F0: 509}
                                    _ = v13
                                }
                            case "a":
                                v3 -= fi8
                            case "hello":
                                v3 -= (^(v3))
                            }
                            switch "hello" {
                            case "\n\t":
                                switch "" {
                                case "\n\t":
                                    if (0 > 653) {
                                        if (!v1) {
                                            var v14 uint64 = 352
                                            _ = v14
                                        }
                                        switch ("\n\t" + "foo bar") {
                                        case "":
                                            if (v6 < -1) {
                                                _ = len(v4)
                                                if false {
                                                    fmt.Println(fi8)
                                                    v7 -= ((1e10 - 0.1) - (float32((float64(((float32(3.14159) - ((float32(float64(v7)) / (v7 * float32(1.5))) + -2.25)) + (v7 + (v7 / ((v7 * v7) + (v7 * v7)))))) * v5)) + 0.0))
                                                } else {
                                                    v4 = append(v4, "hello")
                                                }
                                                v2[("\n\t" + ("\n\t" + ("foo bar" + "")))] = 971
                                            } else {
                                                v15 := GenS1{F0: 0}
                                                v3 = fi11
                                                var v16 int16 = -1
                                                _ = v15
                                                _ = v16
                                            }
                                        }
                                        delete(v2, ("hello" + "\n\t"))
                                        switch "foo bar" {
                                        case "foo bar":
                                            v17 := v2[""]
                                            _ = v17
                                        case "hello":
                                            v1 = v1
                                        case "":
                                            switch "hello" {
                                            case "foo bar":
                                                var v18 float32 = 3.14159
                                                _ = v18
                                            }
                                        }
                                    }
                                case "":
                                    if (42 <= 221) {
                                        fmt.Println(v5, v3)
                                    } else {
                                        v7 = 3.14159
                                    }
                                default:
                                    v3 = -2147483648
                                }
                            default:
                                fmt.Println(v1, v3, v6)
                            }
                        } else {
                            v6 = 1
                            if v1 {
                                var v19 int16 = -1
                                fmt.Println(v3, v1)
                                _ = v19
                            }
                        }
                        var v20 string = "hello"
                        _ = v20
                    }
                    v6 = -1
                    var v21 int32 = -312
                    _ = v21
                }
                v22 := GenS1{F0: 127}
                _ = fi11
                _ = v22
            }
        }
        switch "hello" {
        case "a":
            v7 += ((((float32((v5 / v5)) - v7) * 1.5) - v7) - float32(0.1))
        case "hello":
            if v1 {
                switch fi8 {
                case 128:
                    var v23 uint64 = 869
                    _ = v23
                default:
                    if (0 >= 70) {
                        v7 = (v7 - ((float32(v5) - v7) * float32((0.0 - float64(float32(v5))))))
                    }
                }
                if v1 {
                    if v1 {
                        var v24 string = (("a" + (("hello" + "\n\t") + "hello")) + "hello")
                        v25 := GenS2{F0: (fi8 % ((v3) | 1)), F1: v1, F2: -128}
                        _ = v24
                        _ = v25
                    }
                }
            }
        default:
            v26 := GenS3{F0: 499}
            _ = v26
        }
        _ = fi8
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := GenS2{F0: -234, F1: (v3 == fi27), F2: 1}
        for fi29 := 0; fi29 < 4; fi29++ {
            v5 += float64(0.1)
            var v30 int = (v3 * fi29)
            fmt.Println(v1, v5)
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        v32 := v2[("hello" + ("hello" + ("a" + ("foo bar" + ("a" + ("\n\t" + "\n\t"))))))]
        _ = fi31
        _ = v32
    }
    v7 += v7
    switch "\n\t" {
    case "foo bar":
        var v33 uint16 = 255
        _ = v33
    case "a":
        v34 := GenS4{F0: 2147483648, F1: 0, F2: -434}
        _ = v34
    }
    _ = len(v4)
    fmt.Println(v5, v7)
    fmt.Println(v3, v7)
    if (32767 > 1) {
        var v35 int16 = 535
        _ = v35
    } else {
        if (!(true && false)) {
            v4 = append(v4, ("foo bar" + "foo bar"))
            v3 = v3
            fmt.Println(v5, v7)
        }
    }
    var v36 int = v3
    var v37 int64 = v6
    v38 := GenS4{F0: 2147483647, F1: 4294967295, F2: int32(v3)}
    switch 4294967295 {
    case -256:
        v39 := GenS5{F0: 0}
        _ = v39
    default:
        v4 = append(v4, "a")
    }
    if ((v37 + (v37 | (^(v37)))) != -585) {
        v6 = -35
        for fi40 := 0; fi40 < 16; fi40++ {
            _ = len(v4)
            v41 := GenS6{F0: 72, F1: 1, F2: 4294967295}
            v42 := GenS7{F0: 0}
            _ = fi40
            _ = v41
            _ = v42
        }
        for fi43 := 0; fi43 < 4; fi43++ {
            var v44 string = ""
            _ = fi43
            _ = v44
        }
    } else {
        for fi45 := 0; fi45 < 11; fi45++ {
            for fi46 := 0; fi46 < 13; fi46++ {
                var v47 string = ""
                v5 = (0.1 * 0.1)
                _ = fi46
                _ = v47
            }
            for fi48 := 0; fi48 < 18; fi48++ {
                switch ((("a" + ("\n\t" + "a")) + ("hello" + "a")) + ("\n\t" + "a")) {
                case "a":
                    if (!v1) {
                        v49 := v2["hello"]
                        if (!true) {
                            v36 += (^(fi45))
                            delete(v2, ("" + "hello"))
                            v3 += (fi48 ^ fi45)
                        } else {
                            fmt.Println(v5)
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                            fmt.Println(v49, v7, v36)
                        }
                        var v50 uint16 = 2
                        switch 4294967295 {
                        case -952:
                            switch ("" + "\n\t") {
                            case "hello":
                                delete(v2, "")
                            default:
                                switch ("a" + "\n\t") {
                                case "":
                                    v37 = -1
                                case "\n\t":
                                    switch "\n\t" {
                                    case "a":
                                        v4 = append(v4, "")
                                    default:
                                        var v51 float64 = (float64(float32(-2.25)) + ((float64(3.14159) * v5) + -2.25))
                                        _ = v51
                                    }
                                case "a":
                                    fmt.Println(v6, v49, v36)
                                default:
                                    v36 = v36
                                }
                            }
                        case -1:
                            var v52 int64 = (v37 * v6)
                            _ = v52
                        }
                        if v1 {
                            var v53 uint16 = 0
                            if (-32768 != -1) {
                                v37 = int64(fi45)
                                var v54 uint32 = 42
                                _ = v54
                            } else {
                                v55 := GenS1{F0: v3}
                                _ = v55
                            }
                            switch 42 {
                            case -560:
                                if v1 {
                                    v56 := v2["\n\t"]
                                    _ = v56
                                } else {
                                    v4 = append(v4, "")
                                }
                            }
                            _ = v53
                        }
                        if ((v1 && (255 < 2)) && false) {
                            fmt.Println(v6)
                            var v57 int8 = 25
                            var v58 uint32 = 4294967295
                            _ = v57
                            _ = v58
                        }
                        switch "foo bar" {
                        case "":
                            v59 := GenS1{F0: fi45}
                            _ = v59
                        default:
                            if (((v49 != 306) && (!(true && true))) || (-125 != 42)) {
                                v60 := GenS8{F0: 6}
                                delete(v2, ((("a" + "hello") + ("a" + "\n\t")) + ("\n\t" + "")))
                                _ = v60
                            }
                        }
                        var v61 uint32 = v49
                        _ = v49
                        _ = v50
                        _ = v61
                    }
                }
                switch ("hello" + "\n\t") {
                case "hello":
                    var v62 bool = (0 >= 100)
                    fmt.Println(v36, v1)
                    if (v62 && false) {
                        v4 = append(v4, "foo bar")
                        v7 = v7
                    }
                    _ = v62
                case "foo bar":
                    v2["a"] = 12
                default:
                    if (-1 >= -93) {
                        var v63 int16 = 32767
                        _ = v63
                    }
                }
                _ = fi48
            }
            _ = fi45
        }
        if (v1 && false) {
            if (1 < 65535) {
                for fi64 := 0; fi64 < 17; fi64++ {
                    v37 = v6
                    v65 := GenS3{F0: 2}
                    if ("a" < "a") {
                        v36 = (fi64 / ((796) | 1))
                        v5 += 1e10
                        fmt.Println(fi64)
                    } else {
                        var v66 float32 = v7
                        v67 := GenS4{F0: 98, F1: 0, F2: -2147483648}
                        _ = v66
                        _ = v67
                    }
                    _ = fi64
                    _ = v65
                }
                switch 984 {
                case 0:
                    if (!true) {
                        v68 := GenS9{F0: 935, F1: "\n\t"}
                        var v69 bool = (!(!false))
                        _ = v68
                        _ = v69
                    } else {
                        var v70 int8 = 1
                        v71 := GenS10{F0: 42, F1: 2}
                        _ = v70
                        _ = v71
                    }
                case -9223372036854775808:
                    var v72 int16 = -387
                    _ = v72
                default:
                    if (v1 || false) {
                        switch -9223372036854775808 {
                        case -958:
                            for fi73 := 0; fi73 < 8; fi73++ {
                                fmt.Println(fi73, v37, v7)
                                v37 += (v6 & v6)
                                v5 = v5
                                _ = fi73
                            }
                        default:
                            var v74 string = "\n\t"
                            _ = v74
                        }
                        v75 := GenS11{F0: 1, F1: 436}
                        v76 := GenS12{F0: 1}
                        _ = v75
                        _ = v76
                    } else {
                        for fi77 := 0; fi77 < 12; fi77++ {
                            switch v36 {
                            case -846:
                                v4 = append(v4, ("\n\t" + ""))
                            case 9223372036854775807:
                                var v78 float64 = v5
                                _ = v78
                            case 2147483647:
                                v1 = false
                            default:
                                v7 = 0.0
                            }
                            _ = fi77
                        }
                        if (v5 <= v5) {
                            for fi79 := 0; fi79 < 13; fi79++ {
                                var v80 uint32 = 2147483648
                                _ = fi79
                                _ = v80
                            }
                            _ = len(v4)
                            var v81 uint32 = 4294967295
                            _ = v81
                        }
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                    }
                }
            } else {
                var v82 int32 = 2147483647
                var v83 string = ((("foo bar" + "a") + "\n\t") + ("\n\t" + (("a" + "") + "\n\t")))
                for fi84 := 0; fi84 < 8; fi84++ {
                    for fi85 := 0; fi85 < 14; fi85++ {
                        if len(v4) > 0 {
                            v4[0] = v83
                        }
                        var v86 float32 = 1e10
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi85
                        _ = v86
                    }
                    _ = fi84
                }
                _ = v82
                _ = v83
            }
            v37 = (v6 & v6)
        }
        if true {
            delete(v2, "hello")
            if (((v3 % (((v3 & v3)) | 1)) != v3) || false) {
                for fi87 := 0; fi87 < 20; fi87++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi87
                }
                if ((!(0 > 0)) && (v5 < ((v5 - (v5 * (v5 / (0.1 / 0.1)))) / (0.1 / (0.1 + v5))))) {
                    if false {
                        for fi88 := 0; fi88 < 10; fi88++ {
                            v5 = float64(((((0.0 * (v7 * (1e10 - -2.25))) - v7) - v7) * ((-2.25 + (v7 + ((-2.25 / v7) - 1e10))) * v7)))
                            for fi89 := 0; fi89 < 19; fi89++ {
                                switch int(v6) {
                                case 671:
                                    switch "\n\t" {
                                    case "a":
                                        _ = len(v4)
                                    case "foo bar":
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                    case "":
                                        var v90 int16 = 1
                                        _ = v90
                                    default:
                                        if v1 {
                                            switch (v3 * fi88) {
                                            case -122:
                                                v91 := GenS10{F0: 32767, F1: 4294967295}
                                                _ = v91
                                            case 1:
                                                delete(v2, "foo bar")
                                                switch fi89 {
                                                case -2147483648:
                                                    fmt.Println(v7, v5)
                                                case -1:
                                                    v92 := GenS12{F0: 42}
                                                    _ = v92
                                                case 4294967295:
                                                    var v93 uint32 = 573
                                                    _ = v93
                                                }
                                                v94 := GenS2{F0: (fi89 * (v3 >> ((1) & 63))), F1: (!true), F2: 42}
                                                var v95 int64 = v37
                                                _ = len(v4)
                                                _ = v94
                                                _ = v95
                                            default:
                                                var v96 float64 = v5
                                                _ = v96
                                            }
                                        } else {
                                            if false {
                                                switch ("\n\t" + ("foo bar" + "\n\t")) {
                                                case "foo bar":
                                                    v5 += v5
                                                }
                                            } else {
                                                v5 = (((v5 - (-2.25 + 0.1)) + 3.14159) / v5)
                                            }
                                            if v1 {
                                                fmt.Println(v37)
                                                var v97 float64 = (1.5 / (v5 * ((v5 * v5) + (3.14159 + (1e10 + 0.0)))))
                                                _ = len(v4)
                                                if v1 {
                                                    if (2147483647 > 369) {
                                                        switch (("" + "a") + ("foo bar" + "\n\t")) {
                                                        case "":
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        case "\n\t":
                                                            var v98 uint8 = 128
                                                            _ = v98
                                                        case "foo bar":
                                                            v99 := GenS10{F0: 2, F1: 18446744073709551615}
                                                            _ = v99
                                                        default:
                                                            fmt.Println(v36)
                                                        }
                                                        if v1 {
                                                            var v100 uint8 = 1
                                                            _ = v100
                                                        }
                                                    } else {
                                                        switch "" {
                                                        case "foo bar":
                                                            if len(v4) > 0 {
                                                                v4[0] = "\n\t"
                                                            }
                                                        case "a":
                                                            fmt.Println(v5)
                                                        case "hello":
                                                            var v101 string = "a"
                                                            _ = v101
                                                        default:
                                                            v102 := v2[("a" + ("" + "a"))]
                                                            _ = v102
                                                        }
                                                    }
                                                    v7 = 1.5
                                                    switch v36 {
                                                    case 4294967295:
                                                        v103 := GenS4{F0: 2147483647, F1: 776, F2: 621}
                                                        _ = v103
                                                    case 42:
                                                        _ = len(v2)
                                                    case 128:
                                                        if v1 {
                                                            v6 = (v37 | -1)
                                                        } else {
                                                            switch ("\n\t" + "\n\t") {
                                                            case "\n\t":
                                                                if (0.1 > float32((v5 * float64(0.1)))) {
                                                                    if (42 < 2) {
                                                                        v6 += (-(v37))
                                                                        var v104 string = "foo bar"
                                                                        v3 -= fi89
                                                                        _ = v104
                                                                    }
                                                                } else {
                                                                    _ = len(v2)
                                                                    v105 := GenS13{F0: 127}
                                                                    switch "" {
                                                                    case "hello":
                                                                        if v1 {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        } else {
                                                                            var v106 bool = (!true)
                                                                            v107 := v2["foo bar"]
                                                                            var v108 string = ""
                                                                            _ = v106
                                                                            _ = v107
                                                                            _ = v108
                                                                        }
                                                                    case "a":
                                                                        switch 983 {
                                                                        case 9223372036854775807:
                                                                            if (0 > 2) {
                                                                                var v109 uint64 = 18446744073709551615
                                                                                if v1 {
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = ("a" + (("" + (("" + "") + "")) + "\n\t"))
                                                                                    }
                                                                                    _ = len(v4)
                                                                                    v97 = v97
                                                                                }
                                                                                _ = v109
                                                                            } else {
                                                                                fmt.Println(v7, fi88, v3)
                                                                            }
                                                                        case 32767:
                                                                            var v110 int = 9223372036854775807
                                                                            _ = v110
                                                                        case -1:
                                                                            var v111 int16 = -1
                                                                            _ = v111
                                                                        default:
                                                                            switch "hello" {
                                                                            case "a":
                                                                                _ = len(v2)
                                                                            default:
                                                                                v112 := GenS14{F0: 32767}
                                                                                _ = v112
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        if v1 {
                                                                            var v113 uint32 = 2147483647
                                                                            switch v3 {
                                                                            case 128:
                                                                                v114 := GenS4{F0: v113, F1: 0, F2: 137}
                                                                                if v1 {
                                                                                    v36 = (-(v36))
                                                                                    if ((913 == 0) && v1) {
                                                                                        v115 := GenS4{F0: v113, F1: 0, F2: -2147483648}
                                                                                        var v116 int8 = 42
                                                                                        var v117 int8 = -31
                                                                                        var v118 uint8 = 1
                                                                                        if (float64((v7 * (v7 * ((v7 / v7) + 1.5)))) >= v97) {
                                                                                            if v1 {
                                                                                                fmt.Println(fi88)
                                                                                                var v119 uint32 = (v113 ^ 775)
                                                                                                v1 = v1
                                                                                                var v120 float64 = float64(v7)
                                                                                                _ = v119
                                                                                                _ = v120
                                                                                            } else {
                                                                                                if (!v1) {
                                                                                                    _ = len(v4)
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v4 = append(v4, "\n\t")
                                                                                            var v121 float32 = 3.14159
                                                                                            v117 -= (v116 + (-(v116)))
                                                                                            _ = v121
                                                                                        }
                                                                                        v3 += (v36 ^ -2147483648)
                                                                                        var v122 uint = 721
                                                                                        _ = v115
                                                                                        _ = v116
                                                                                        _ = v117
                                                                                        _ = v118
                                                                                        _ = v122
                                                                                    }
                                                                                    switch ("\n\t" + "a") {
                                                                                    case "\n\t":
                                                                                        v5 = v97
                                                                                    default:
                                                                                        var v123 string = "hello"
                                                                                        _ = v123
                                                                                    }
                                                                                } else {
                                                                                    v3 += fi88
                                                                                    v124 := GenS9{F0: 1, F1: (("\n\t" + "foo bar") + "foo bar")}
                                                                                    _ = v124
                                                                                }
                                                                                if (v36 == fi89) {
                                                                                    v125 := GenS15{F0: int16(v6)}
                                                                                    v97 += 3.14159
                                                                                    v6 += 42
                                                                                    _ = v125
                                                                                }
                                                                                _ = v114
                                                                            case -333:
                                                                                _ = len(v2)
                                                                            case -298:
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = "a"
                                                                                }
                                                                            default:
                                                                                var v126 uint32 = 42
                                                                                v4 = append(v4, "")
                                                                                _ = v126
                                                                            }
                                                                            v5 = v97
                                                                            _ = v113
                                                                        } else {
                                                                            var v127 float32 = 1e10
                                                                            fmt.Println(v7, v5)
                                                                            _ = v127
                                                                        }
                                                                    default:
                                                                        v128 := GenS6{F0: 127, F1: 127, F2: 62}
                                                                        _ = v128
                                                                    }
                                                                    if v1 {
                                                                        v129 := GenS12{F0: 42}
                                                                        var v130 uint = 2
                                                                        _ = v129
                                                                        _ = v130
                                                                    }
                                                                    var v131 float32 = (v7 * v7)
                                                                    _ = v105
                                                                    _ = v131
                                                                }
                                                            case "hello":
                                                                v36 += (v3 >> ((18446744073709551615) & 63))
                                                            case "foo bar":
                                                                v37 -= v37
                                                            default:
                                                                v132 := GenS10{F0: 32767, F1: 4294967295}
                                                                _ = v132
                                                            }
                                                            v1 = (-39 == v6)
                                                        }
                                                    default:
                                                        v3 -= (v3 & (v36 ^ (-(v36))))
                                                    }
                                                } else {
                                                    var v133 int64 = v6
                                                    if ((181 <= 42) || (943 <= 0)) {
                                                        v37 -= (v133 - v133)
                                                    }
                                                    v7 -= (v7 + (3.14159 * v7))
                                                    _ = v133
                                                }
                                                _ = v97
                                            } else {
                                                var v134 float64 = v5
                                                v36 = v36
                                                _ = v134
                                            }
                                            var v135 float32 = (v7 * v7)
                                            _ = v135
                                        }
                                    }
                                case -9223372036854775808:
                                    v136 := GenS16{F0: 2, F1: 2147483647, F2: 4294967295}
                                    _ = v136
                                case -2147483648:
                                    switch (("hello" + ("" + ("a" + "foo bar"))) + "") {
                                    case "hello":
                                        v36 -= (^(fi88))
                                    case "foo bar":
                                        if v1 {
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                if (-1 > -530) {
                                                    var v137 int32 = 42
                                                    _ = len(v2)
                                                    _ = v137
                                                }
                                            }
                                        } else {
                                            if true {
                                                v138 := GenS17{F0: 0, F1: v37, F2: 1}
                                                v139 := GenS7{F0: 127}
                                                v37 += v37
                                                _ = v138
                                                _ = v139
                                            } else {
                                                var v140 bool = true
                                                v141 := GenS15{F0: 42}
                                                v142 := GenS18{F0: 127}
                                                _ = v140
                                                _ = v141
                                                _ = v142
                                            }
                                            var v143 uint = 0
                                            _ = v143
                                        }
                                    case "":
                                        if v1 {
                                            v4 = append(v4, "foo bar")
                                            fmt.Println(v3, v3, v7)
                                        }
                                    }
                                default:
                                    switch fi88 {
                                    case -9223372036854775808:
                                        v144 := GenS12{F0: -1}
                                        _ = v144
                                    }
                                }
                                v1 = (18446744073709551615 < 706)
                                switch "\n\t" {
                                case "foo bar":
                                    var v145 uint64 = 1
                                    _ = v145
                                case "":
                                    var v146 string = "hello"
                                    _ = v146
                                case "\n\t":
                                    v5 -= 3.14159
                                default:
                                    v6 += v37
                                }
                                fmt.Println(v3, fi88, fi89)
                                v7 -= v7
                                _ = len(v4)
                                _ = fi89
                            }
                            _ = fi88
                        }
                        var v147 int64 = (v6 % ((v37) | 1))
                        _ = v147
                    } else {
                        for fi148 := 0; fi148 < 10; fi148++ {
                            v149 := GenS19{F0: -1}
                            fmt.Println(v1, fi148)
                            _ = fi148
                            _ = v149
                        }
                        switch v36 {
                        case -1:
                            v1 = true
                        case 764:
                            v6 = -1
                        default:
                            if (!false) {
                                v150 := GenS17{F0: -44, F1: v37, F2: (v37 * v37)}
                                _ = v150
                            } else {
                                var v151 float32 = float32(1e10)
                                v152 := GenS14{F0: 255}
                                for fi153 := 0; fi153 < 12; fi153++ {
                                    if v1 {
                                        var v154 int32 = 42
                                        for fi155 := 0; fi155 < 3; fi155++ {
                                            _ = len(v2)
                                            switch ("hello" + "foo bar") {
                                            case "hello":
                                                v156 := GenS20{F0: (("" + "hello") + "a"), F1: 514}
                                                _ = v156
                                            case "foo bar":
                                                var v157 float32 = (v151 + v7)
                                                _ = v157
                                            default:
                                                if len(v4) > 0 {
                                                    v4[0] = ("foo bar" + (("" + "hello") + ""))
                                                }
                                            }
                                            v158 := GenS13{F0: 0}
                                            var v159 int32 = (v154 - 2147483647)
                                            _ = fi155
                                            _ = v158
                                            _ = v159
                                        }
                                        var v160 uint32 = 2147483647
                                        _ = v154
                                        _ = v160
                                    }
                                    _ = fi153
                                }
                                _ = v151
                                _ = v152
                            }
                        }
                        _ = len(v4)
                    }
                    if (!v1) {
                        for fi161 := 0; fi161 < 18; fi161++ {
                            for fi162 := 0; fi162 < 7; fi162++ {
                                if (3.14159 != v5) {
                                    switch "" {
                                    case "":
                                        if ((!(!v1)) || (2 > 976)) {
                                            v4 = append(v4, "\n\t")
                                            _ = len(v4)
                                            _ = len(v2)
                                        } else {
                                            var v163 uint8 = 70
                                            v164 := GenS10{F0: 255, F1: 2}
                                            if (v7 >= float32(v5)) {
                                                v4 = append(v4, "")
                                                v3 -= 4294967295
                                            } else {
                                                var v165 uint32 = 2147483648
                                                var v166 float64 = ((v5 - v5) + 0.1)
                                                v167 := GenS17{F0: -79, F1: -566, F2: (v6 / (((v37 % ((42) | 1))) | 1))}
                                                _ = v165
                                                _ = v166
                                                _ = v167
                                            }
                                            _ = v163
                                            _ = v164
                                        }
                                    case "foo bar":
                                        _ = len(v4)
                                    default:
                                        v37 = -9223372036854775808
                                    }
                                    switch (-(fi161)) {
                                    case -9223372036854775808:
                                        v2["foo bar"] = 67
                                    case 1:
                                        var v168 int = v3
                                        _ = v168
                                    default:
                                        v36 -= (v3 >> ((0) & 63))
                                    }
                                    switch "" {
                                    case "hello":
                                        v2[("" + ("\n\t" + ("\n\t" + "\n\t")))] = 42
                                    case "\n\t":
                                        v169 := GenS16{F0: 42, F1: 2147483648, F2: 1}
                                        _ = v169
                                    case "":
                                        var v170 int16 = -32768
                                        _ = v170
                                    }
                                }
                                v171 := GenS15{F0: 0}
                                fmt.Println(v7, v7, v3)
                                _ = fi162
                                _ = v171
                            }
                            for fi172 := 0; fi172 < 1; fi172++ {
                                fmt.Println(v5, fi172, v7)
                                _ = fi172
                            }
                            for fi173 := 0; fi173 < 3; fi173++ {
                                switch "foo bar" {
                                case "\n\t":
                                    var v174 uint8 = 87
                                    _ = v174
                                case "a":
                                    v37 -= (^(v6))
                                default:
                                    fmt.Println(v37)
                                }
                                _ = fi173
                            }
                            _ = fi161
                        }
                        for fi175 := 0; fi175 < 20; fi175++ {
                            if (0 != 0) {
                                v7 = 3.14159
                                v3 += fi175
                            } else {
                                for fi176 := 0; fi176 < 11; fi176++ {
                                    v3 = v3
                                    v6 = (v37 % ((v37) | 1))
                                    fmt.Println(v1)
                                    _ = fi176
                                }
                            }
                            v5 = float64((0.1 + v7))
                            _ = fi175
                        }
                    }
                    for fi177 := 0; fi177 < 16; fi177++ {
                        fmt.Println(v36, v5, fi177)
                        switch (v3 / ((v3) | 1)) {
                        case -158:
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v36, v36, v6)
                            default:
                                v178 := v2["hello"]
                                _ = v178
                            }
                        case 32768:
                            v36 = v3
                        }
                        _ = fi177
                    }
                } else {
                    for fi179 := 0; fi179 < 1; fi179++ {
                        fmt.Println(v1)
                        _ = fi179
                    }
                    v180 := GenS14{F0: 561}
                    var v181 uint8 = 128
                    _ = v180
                    _ = v181
                }
            } else {
                v182 := GenS21{F0: 2}
                var v183 bool = false
                var v184 uint64 = 18446744073709551615
                _ = v182
                _ = v183
                _ = v184
            }
        } else {
            for fi185 := 0; fi185 < 13; fi185++ {
                var v186 int = v36
                _ = fi185
                _ = v186
            }
            v187 := GenS22{F0: int8(v36), F1: (true && true), F2: (!((((v5 * v5) + v5) + ((v5 + 0.0) / v5)) < float64(float32((v5 / v5)))))}
            v188 := GenS6{F0: 42, F1: -56, F2: 468}
            _ = v187
            _ = v188
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
