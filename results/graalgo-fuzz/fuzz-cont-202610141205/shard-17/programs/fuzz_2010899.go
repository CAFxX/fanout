package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int; F1 int8; F2 int16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64; F1 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 int8 = -1
    var v2 int = -145
    v3 := make(map[string]int)
    var v4 int64 = -9223372036854775808
    var v5 float32 = 3.14159
    var v6 uint8 = 255
    if false {
        v4 += (-(v4))
        var v7 int32 = 42
        _ = v7
    } else {
        _ = len(v3)
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if (!(false || false)) {
            var v9 int8 = (v1 + v1)
            v2 -= 4294967295
            _ = v9
        }
        if false {
            if false {
                for fi10 := 0; fi10 < 2; fi10++ {
                    if false {
                        v11 := []int{}
                        if (-1 < (-(v4))) {
                            if (((false || (32767 > -695)) || (false || ((466 >= 42) && (v6 >= (^(v6)))))) || ("" >= "")) {
                                _ = len(v11)
                                var v12 uint64 = 412
                                _ = v12
                            }
                        } else {
                            if ((v1 - 0) != v1) {
                                if (0.0 >= 1e10) {
                                    v3["a"] = v2
                                    var v13 int = (fi10 << ((255) & 63))
                                    _ = v13
                                } else {
                                    var v14 uint64 = 18446744073709551615
                                    _ = len(v3)
                                    _ = v14
                                }
                                if false {
                                    var v15 uint64 = 522
                                    v16 := GenS1{F0: 679, F1: v1}
                                    var v17 string = "a"
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                } else {
                                    v4 += (v4 * (v4 & v4))
                                    fmt.Println(v5, fi10)
                                    if false {
                                        var v18 int64 = 42
                                        switch v2 {
                                        case 1:
                                            var v19 uint = 18446744073709551615
                                            _ = v19
                                        case 32768:
                                            var v20 uint16 = 65535
                                            _ = v20
                                        case -9223372036854775808:
                                            _ = len(v3)
                                        default:
                                            var v21 uint = 222
                                            v5 -= (float32(float64(v5)) + float32(3.14159))
                                            _ = v21
                                        }
                                        v22 := GenS2{F0: 0, F1: v1}
                                        _ = v18
                                        _ = v22
                                    }
                                    var v23 float32 = v5
                                    _ = v23
                                }
                            }
                        }
                        _ = v11
                    }
                    v1 = (v1 | v1)
                    v24 := GenS2{F0: uint64(v6), F1: (^(v1))}
                    var v25 uint = 42
                    _ = fi10
                    _ = v24
                    _ = v25
                }
                var v26 int8 = -128
                _ = v26
            }
            v27 := GenS2{F0: 1, F1: v1}
            if false {
                v28 := []int{fi8, -229, fi8}
                fmt.Println(fi8, v6, v5)
                _ = v28
            }
            _ = v27
        }
        v4 -= -1
        _ = fi8
    }
    var v29 float64 = 0.1
    switch (^(v2)) {
    case -205:
        v30 := v3[""]
        _ = v30
    case -9223372036854775808:
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 float32 = -2.25
            _ = fi31
            _ = v32
        }
    case 4294967295:
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 int8 = v1
            var v35 int8 = (v34 | v34)
            _ = fi33
            _ = v34
            _ = v35
        }
    default:
        var v36 float64 = (((1.5 * (v29 * ((0.0 + (v29 / 1.0)) - float64(0.1)))) * v29) * ((((v29 * (((((((((v29 + v29) * ((v29 * (3.14159 - 1e10)) * ((0.1 / 1.0) - (((1e10 + (float64(v5) - -2.25)) * (v29 + v29)) / (1e10 / v29))))) + v29) + v29) + 3.14159) * 0.1) * 1e10) * v29) - v29)) / 1e10) + v29) + (((((v29 / 1.5) - v29) - v29) + ((((v29 - v29) * float64((v5 + 3.14159))) / 1.5) + (((((float64(1e10) + 0.1) - v29) - 0.1) / 1.5) - (1e10 - v29)))) * 0.1)))
        _ = v36
    }
    var v37 int16 = 32767
    v38 := []int{(v2 / (((v2 / ((-900) | 1))) | 1)), 4294967295, v2, int(v1)}
    var v39 float64 = 1.5
    var v40 float32 = (((v5 / (v5 - (-2.25 / 3.14159))) * (v5 - 3.14159)) / (1e10 - v5))
    switch ("hello" + "foo bar") {
    case "\n\t":
        v41 := GenS1{F0: 65535, F1: (v1 ^ (v1 ^ v1))}
        _ = v41
    }
    switch "\n\t" {
    case "foo bar":
        v42 := GenS1{F0: 2, F1: v1}
        _ = v42
    case "hello":
        fmt.Println(v5)
    case "a":
        switch (("hello" + "") + ("\n\t" + ("a" + (("\n\t" + "\n\t") + "hello")))) {
        case "":
            v6 = v6
        default:
            var v43 float64 = (v39 / (v39 / (v39 / (v39 * 3.14159))))
            _ = v43
        }
    }
    switch ("\n\t" + "\n\t") {
    case "a":
        var v44 bool = true
        _ = v44
    case "":
        for fi45 := 0; fi45 < 13; fi45++ {
            if false {
                switch "hello" {
                case "":
                    for fi46 := 0; fi46 < 9; fi46++ {
                        v39 = v39
                        _ = fi46
                    }
                case "\n\t":
                    for fi47 := 0; fi47 < 16; fi47++ {
                        var v48 int16 = 0
                        v49 := GenS3{F0: 0, F1: v1}
                        switch ("a" + "") {
                        case "":
                            var v50 uint32 = 239
                            _ = v50
                        default:
                            v40 -= (v40 / v40)
                        }
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                default:
                    var v51 uint = 18446744073709551615
                    _ = v51
                }
                if false {
                    v52 := GenS1{F0: 555, F1: 127}
                    _ = v52
                }
            } else {
                v38 = append(v38, (fi45 | 0))
                switch "a" {
                case "foo bar":
                    fmt.Println(v4, v1, fi45)
                default:
                    if false {
                        v5 = v40
                        if true {
                            for fi53 := 0; fi53 < 7; fi53++ {
                                v5 = v5
                                v54 := GenS1{F0: 32767, F1: 100}
                                _ = fi53
                                _ = v54
                            }
                            switch ("foo bar" + "\n\t") {
                            case "":
                                v55 := v3["hello"]
                                _ = v55
                            case "hello":
                                v37 += v37
                            case "foo bar":
                                v40 = 1.5
                            default:
                                var v56 float32 = ((v40 + ((v5 / v40) - ((-2.25 / (v5 * v40)) - (v5 + 1e10)))) / 1e10)
                                _ = v56
                            }
                        } else {
                            for fi57 := 0; fi57 < 20; fi57++ {
                                var v58 int16 = v37
                                var v59 float32 = (((v5 + v40) + 1.5) * v40)
                                fmt.Println(v40, v6)
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            var v60 uint16 = 16
                            _ = v60
                        }
                        var v61 int8 = v1
                        _ = v61
                    } else {
                        if false {
                            if len(v38) >= 2 {
                                v38 = v38[1:len(v38)]
                            }
                            v4 = -9223372036854775808
                        } else {
                            fmt.Println(v37)
                            v62 := GenS4{F0: 0, F1: 4294967295}
                            _ = v62
                        }
                        if (!(false || true)) {
                            v29 = ((((0.1 * 1.5) + v39) * float64(v5)) - v39)
                            if (true || false) {
                                switch fi45 {
                                case 128:
                                    for fi63 := 0; fi63 < 3; fi63++ {
                                        if true {
                                            switch "" {
                                            case "\n\t":
                                                switch (("hello" + ("hello" + "hello")) + "foo bar") {
                                                case "hello":
                                                    switch fi45 {
                                                    case -858:
                                                        switch fi45 {
                                                        case 2147483647:
                                                            var v64 float64 = v39
                                                            _ = v64
                                                        case 0:
                                                            fmt.Println(v37, v39, v5)
                                                            v29 = (((-2.25 - v29) * (3.14159 + v39)) * v29)
                                                            if (1 <= 718) {
                                                                if (fi63 == (^(fi63))) {
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v1 = (v1 >> ((42) & 63))
                                                                    }
                                                                    v65 := GenS3{F0: 42, F1: (v1 ^ v1)}
                                                                    _ = v65
                                                                }
                                                            } else {
                                                                switch (((("\n\t" + "\n\t") + (("" + "") + "foo bar")) + (("a" + ("hello" + "foo bar")) + "foo bar")) + "hello") {
                                                                case "":
                                                                    var v66 uint32 = 1
                                                                    _ = v66
                                                                case "\n\t":
                                                                    if true {
                                                                        switch "" {
                                                                        case "":
                                                                            _ = len(v38)
                                                                        default:
                                                                            if true {
                                                                                fmt.Println(v5, v1, v5)
                                                                                v67 := v3["foo bar"]
                                                                                switch ("\n\t" + ("a" + ("a" + "\n\t"))) {
                                                                                case "a":
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        v68 := GenS4{F0: (v37 ^ (v37 >> ((255) & 63))), F1: 4294967295}
                                                                                        _ = v68
                                                                                    default:
                                                                                        v69 := GenS2{F0: 18446744073709551615, F1: (^(v1))}
                                                                                        _ = v69
                                                                                    }
                                                                                    if true {
                                                                                        var v70 int8 = 42
                                                                                        var v71 uint = 635
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                    } else {
                                                                                        _ = len(v38)
                                                                                        if len(v38) > 0 {
                                                                                            v38[0] = 1
                                                                                        }
                                                                                        v6 += 0
                                                                                    }
                                                                                default:
                                                                                    v72 := GenS5{F0: (v4 & -546), F1: (v4 >> ((uint(v6)) & 63))}
                                                                                    _ = v72
                                                                                }
                                                                                _ = v67
                                                                            } else {
                                                                                v37 = 1
                                                                                var v73 uint8 = (v6 / ((v6) | 1))
                                                                                _ = v73
                                                                            }
                                                                        }
                                                                        var v74 float64 = v29
                                                                        v75 := GenS5{F0: v4, F1: (v4 ^ v4)}
                                                                        _ = v74
                                                                        _ = v75
                                                                    } else {
                                                                        if ("hello" < "\n\t") {
                                                                            fmt.Println(v6, v40, fi45)
                                                                        } else {
                                                                            if (!false) {
                                                                                if false {
                                                                                    v76 := GenS6{F0: v37, F1: ("hello" + ("" + "hello"))}
                                                                                    v77 := GenS1{F0: 65535, F1: 36}
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        switch (fi63 % ((294) | 1)) {
                                                                                        case 32767:
                                                                                            switch 694 {
                                                                                            case -63:
                                                                                                v3[""] = int(v1)
                                                                                                fmt.Println(v1, v4)
                                                                                            default:
                                                                                                _ = len(v38)
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v78 uint8 = 253
                                                                                        _ = v78
                                                                                    default:
                                                                                        var v79 bool = (false && false)
                                                                                        _ = v79
                                                                                    }
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                                v80 := v3[""]
                                                                                var v81 uint = 65535
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    fmt.Println(v29)
                                                                                default:
                                                                                    var v82 uint16 = 0
                                                                                    _ = v82
                                                                                }
                                                                                v2 = 32768
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                if (!true) {
                                                                                    v4 = 0
                                                                                    var v83 uint16 = 0
                                                                                    if len(v38) > 0 {
                                                                                        v38[0] = 127
                                                                                    }
                                                                                    fmt.Println(v4, v1)
                                                                                    v84 := GenS3{F0: 890, F1: (^(v1))}
                                                                                    fmt.Println(v1)
                                                                                    _ = v83
                                                                                    _ = v84
                                                                                }
                                                                                if false {
                                                                                    _ = len(v38)
                                                                                    v4 += (-(v4))
                                                                                } else {
                                                                                    v2 -= (v2 & 128)
                                                                                }
                                                                                var v85 bool = true
                                                                                v38 = append(v38, fi45)
                                                                                switch (v2 * fi63) {
                                                                                case 361:
                                                                                    v86 := GenS5{F0: v4, F1: (v4 << ((18446744073709551615) & 63))}
                                                                                    _ = v86
                                                                                case -9223372036854775808:
                                                                                    var v87 string = "a"
                                                                                    _ = v87
                                                                                }
                                                                                var v88 int32 = -2147483648
                                                                                _ = v85
                                                                                _ = v88
                                                                            }
                                                                            var v89 uint8 = (v6 >> ((1) & 63))
                                                                            var v90 string = "foo bar"
                                                                            var v91 uint32 = 65535
                                                                            if (false || false) {
                                                                                if (-1 == (v37 - 609)) {
                                                                                    if false {
                                                                                        v6 = (^(v89))
                                                                                        var v92 float32 = 1e10
                                                                                        _ = v92
                                                                                    } else {
                                                                                        v39 = 0.0
                                                                                    }
                                                                                    var v93 uint = 42
                                                                                    _ = v93
                                                                                }
                                                                                var v94 bool = (433 != 874)
                                                                                var v95 int32 = 1
                                                                                if v94 {
                                                                                    var v96 float32 = v5
                                                                                    if ((!false) || false) {
                                                                                        var v97 uint = 65535
                                                                                        var v98 uint8 = v89
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    }
                                                                                    _ = v96
                                                                                }
                                                                                _ = v94
                                                                                _ = v95
                                                                            }
                                                                            var v99 int = fi63
                                                                            fmt.Println(v2, v37)
                                                                            switch v90 {
                                                                            case "hello":
                                                                                if false {
                                                                                    v39 = float64(((0.0 - v40) + ((v40 / (3.14159 + float32(float64(v5)))) * v40)))
                                                                                }
                                                                            case "":
                                                                                var v100 uint64 = 432
                                                                                _ = v100
                                                                            }
                                                                            _ = v89
                                                                            _ = v90
                                                                            _ = v91
                                                                            _ = v99
                                                                        }
                                                                        var v101 uint = 65535
                                                                        _ = v101
                                                                    }
                                                                }
                                                                if (1 <= 2) {
                                                                    if len(v38) >= 2 {
                                                                        v38 = v38[1:len(v38)]
                                                                    }
                                                                    var v102 float64 = (0.0 - 0.1)
                                                                    _ = v102
                                                                }
                                                                v4 = (v4 ^ v4)
                                                            }
                                                        }
                                                    case 0:
                                                        if true {
                                                            v5 = (v40 + v5)
                                                        }
                                                    case 127:
                                                        if (100 != (^(v6))) {
                                                            fmt.Println(v29)
                                                            switch fi63 {
                                                            case 0:
                                                                var v103 int16 = v37
                                                                _ = v103
                                                            case 42:
                                                                var v104 string = "a"
                                                                _ = v104
                                                            case -890:
                                                                fmt.Println(v2)
                                                            default:
                                                                if false {
                                                                    if (uint(v6) == 2) {
                                                                        v105 := GenS3{F0: -257, F1: (v1 | v1)}
                                                                        _ = v105
                                                                    } else {
                                                                        v106 := GenS7{F0: 127, F1: 1, F2: v37}
                                                                        v107 := GenS8{F0: 47}
                                                                        var v108 float32 = (3.14159 + -2.25)
                                                                        _ = v106
                                                                        _ = v107
                                                                        _ = v108
                                                                    }
                                                                } else {
                                                                    if (false && (0 == 42)) {
                                                                        v109 := GenS5{F0: (v4 ^ (v4 & v4)), F1: v4}
                                                                        _ = v109
                                                                    } else {
                                                                        v2 = v2
                                                                        if true {
                                                                            fmt.Println(v5, v40, v39)
                                                                            v29 -= (v29 - v29)
                                                                        } else {
                                                                            if false {
                                                                                if (((53 < 689) && (!false)) || (!true)) {
                                                                                    var v110 uint8 = v6
                                                                                    v111 := GenS2{F0: 551, F1: (^(v1))}
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                                if (v4 != v4) {
                                                                                    delete(v3, "a")
                                                                                }
                                                                                if true {
                                                                                    v5 = v5
                                                                                }
                                                                            } else {
                                                                                if false {
                                                                                    v38 = append(v38, fi45)
                                                                                }
                                                                            }
                                                                        }
                                                                        if ((18446744073709551615 >= 18446744073709551615) && false) {
                                                                            if false {
                                                                                v112 := GenS9{F0: v4, F1: 4294967295}
                                                                                _ = v112
                                                                            }
                                                                            if true {
                                                                                fmt.Println(v4, v1)
                                                                            } else {
                                                                                if true {
                                                                                    v37 -= -32768
                                                                                    v113 := GenS8{F0: int8(v6)}
                                                                                    _ = v113
                                                                                }
                                                                                delete(v3, "foo bar")
                                                                                v114 := GenS6{F0: (v37 >> ((319) & 63)), F1: "a"}
                                                                                _ = v114
                                                                            }
                                                                            var v115 float32 = v40
                                                                            v37 = v37
                                                                            _ = v115
                                                                        } else {
                                                                            v3[(("hello" + "\n\t") + "")] = v2
                                                                            fmt.Println(v5, fi63, v40)
                                                                            fmt.Println(v40, v37, v39)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            switch "hello" {
                                                            case "a":
                                                                fmt.Println(v5)
                                                            case "":
                                                                var v116 string = "foo bar"
                                                                _ = v116
                                                            case "foo bar":
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if true {
                                                                        v117 := GenS4{F0: 32767, F1: 2147483648}
                                                                        _ = v117
                                                                    }
                                                                    switch (fi45 % ((32768) | 1)) {
                                                                    case 1:
                                                                        if false {
                                                                            switch (fi45 | v2) {
                                                                            case 127:
                                                                                if (!true) {
                                                                                    v118 := GenS3{F0: 0, F1: v1}
                                                                                    _ = v118
                                                                                }
                                                                            case 4294967295:
                                                                                switch fi63 {
                                                                                case -2147483648:
                                                                                    var v119 string = ((("foo bar" + "hello") + "\n\t") + "a")
                                                                                    _ = v119
                                                                                default:
                                                                                    var v120 uint16 = uint16(v6)
                                                                                    _ = v120
                                                                                }
                                                                            case -9223372036854775808:
                                                                                if len(v38) > 0 {
                                                                                    v38[0] = v2
                                                                                }
                                                                            default:
                                                                                _ = len(v38)
                                                                            }
                                                                            v4 += (v4 / (((v4 & (v4 % ((v4) | 1)))) | 1))
                                                                        } else {
                                                                            v121 := v3["foo bar"]
                                                                            v2 = int(v37)
                                                                            fmt.Println(fi45, v37, v121)
                                                                            if false {
                                                                                var v122 uint8 = (v6 % (((v6 - (v6 >> ((0) & 63)))) | 1))
                                                                                _ = v122
                                                                            }
                                                                            _ = v121
                                                                        }
                                                                    case 2147483647:
                                                                        var v123 int = (fi45 * -258)
                                                                        _ = v123
                                                                    case 9223372036854775807:
                                                                        fmt.Println(fi45, v37, v29)
                                                                    }
                                                                    var v124 uint16 = 1
                                                                    v125 := GenS9{F0: v4, F1: 4294967295}
                                                                    switch -995 {
                                                                    case -583:
                                                                        if false {
                                                                            v5 = v5
                                                                            v40 += v40
                                                                            var v126 int16 = 42
                                                                            _ = v126
                                                                        } else {
                                                                            fmt.Println(v29, v4)
                                                                            v29 = v29
                                                                        }
                                                                    }
                                                                    v39 = v39
                                                                    _ = v124
                                                                    _ = v125
                                                                }
                                                            default:
                                                                v37 += v37
                                                            }
                                                        }
                                                        _ = len(v38)
                                                    }
                                                }
                                            case "foo bar":
                                                fmt.Println(v40)
                                            }
                                            var v127 uint64 = 2
                                            v128 := GenS10{F0: -1}
                                            _ = v127
                                            _ = v128
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                if true {
                                                    var v129 uint64 = 42
                                                    _ = v129
                                                } else {
                                                    v3[("\n\t" + ("" + ("a" + "hello")))] = fi45
                                                    fmt.Println(v2, v39, v4)
                                                    var v130 float32 = (v5 - v40)
                                                    _ = v130
                                                }
                                            case "foo bar":
                                                v5 = (3.14159 - 0.0)
                                            case "hello":
                                                var v131 int16 = (v37 / (((v37 >> ((42) & 63))) | 1))
                                                _ = v131
                                            default:
                                                v39 = float64(v40)
                                            }
                                            v39 = (float64(float32(v39)) + v39)
                                            v132 := v3[(((("\n\t" + "\n\t") + (("" + "a") + ("\n\t" + ""))) + "\n\t") + "\n\t")]
                                            _ = v132
                                        }
                                        _ = fi63
                                    }
                                case 32768:
                                    var v133 bool = true
                                    _ = v133
                                case -9223372036854775808:
                                    for fi134 := 0; fi134 < 13; fi134++ {
                                        if false {
                                            var v135 uint32 = 65535
                                            _ = v135
                                        } else {
                                            if true {
                                                var v136 int32 = -1
                                                var v137 int64 = v4
                                                v1 -= 127
                                                _ = v136
                                                _ = v137
                                            }
                                            if ((v6 ^ 127) >= 2) {
                                                var v138 uint = 2
                                                v139 := GenS11{F0: -2147483648, F1: 42}
                                                _ = v138
                                                _ = v139
                                            }
                                        }
                                        _ = fi134
                                    }
                                }
                                v140 := GenS3{F0: 2147483647, F1: (v1 << ((65535) & 63))}
                                _ = v140
                            } else {
                                if len(v38) >= 2 {
                                    v38 = v38[1:len(v38)]
                                }
                            }
                            var v141 float64 = v39
                            _ = v141
                        }
                        var v142 int32 = -881
                        _ = v142
                    }
                }
                var v143 int16 = (v37 & v37)
                _ = v143
            }
            switch fi45 {
            case -1:
                if len(v38) > 0 {
                    v38[0] = (v2 ^ fi45)
                }
            default:
                fmt.Println(v29, v37)
            }
            _ = fi45
        }
    default:
        for fi144 := 0; fi144 < 7; fi144++ {
            v2 -= 1
            v4 = 0
            for fi145 := 0; fi145 < 19; fi145++ {
                v1 = (v1 * (v1 % ((v1) | 1)))
                var v146 uint8 = 76
                var v147 uint8 = 0
                switch "\n\t" {
                case "a":
                    delete(v3, "")
                case "foo bar":
                    v146 = (^(v146))
                    fmt.Println(v1, fi144)
                    v3[""] = v2
                }
                fmt.Println(v2, fi145, v40)
                switch "\n\t" {
                case "a":
                    v3["hello"] = fi144
                default:
                    _ = len(v38)
                }
                switch (fi145 | (^(fi144))) {
                case -2147483648:
                    if (!true) {
                        delete(v3, (("foo bar" + "") + "foo bar"))
                    } else {
                        switch (v2 - (fi144 * (fi145 << ((0) & 63)))) {
                        case -544:
                            var v148 int = -510
                            _ = v148
                        case 32767:
                            v5 += v5
                        case -1:
                            v149 := GenS12{F0: 421}
                            _ = v149
                        }
                        switch ("hello" + ("a" + "hello")) {
                        case "a":
                            var v150 string = "hello"
                            _ = v150
                        case "foo bar":
                            v2 = 128
                        }
                        fmt.Println(v6)
                    }
                case -1:
                    v3[("" + "hello")] = (v2 | (-(fi144)))
                case 32767:
                    fmt.Println(fi145, v147)
                default:
                    if len(v38) > 0 {
                        v38[0] = (^(fi145))
                    }
                }
                _ = fi145
                _ = v146
                _ = v147
            }
            _ = fi144
        }
    }
    fmt.Println(v29, v40, v29)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
