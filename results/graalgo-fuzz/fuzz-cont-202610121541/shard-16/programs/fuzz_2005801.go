package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint32; F1 int64; F2 int }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 int8; F1 int64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint32; F1 uint64 }
type GenS15 struct { F0 uint64; F1 uint64 }
type GenS16 struct { F0 int }
func genFn0() uint64 {
    v4 := GenS1{F0: 255}
    var v5 uint64 = 4294967295
    _ = v4
    _ = v5
    return 875
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    v3 := []string{"foo bar", (("\n\t" + "") + (("" + "\n\t") + "")), (("hello" + ("a" + "")) + ("\n\t" + ""))}
    for fi6 := 0; fi6 < 15; fi6++ {
        v2["foo bar"] = 264
        switch fi6 {
        case 4294967295:
            v7 := GenS1{F0: 55}
            _ = v7
        case 780:
            if (87 != 4294967295) {
                var v8 float32 = 1.5
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 uint = 639
                    _ = len(v3)
                    _ = fi9
                    _ = v10
                }
                _ = v8
            }
        }
        var v11 int32 = -2147483648
        _ = fi6
        _ = v11
    }
    if (!((1.5 * 1e10) > (0.0 - 0.0))) {
        for fi12 := 0; fi12 < 1; fi12++ {
            if true {
                var v13 string = ("foo bar" + "")
                for fi14 := 0; fi14 < 3; fi14++ {
                    v13 = v13
                    v13 = ""
                    var v15 string = v13
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            fmt.Println(fi12)
            var v16 bool = false
            _ = fi12
            _ = v16
        }
        _ = len(v2)
    }
    if true {
        var v17 int64 = 9223372036854775807
        _ = v17
    } else {
        switch "a" {
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            var v18 uint32 = 337
            _ = v18
        }
        if (!(true || (true && ((2147483647 != -2147483648) && false)))) {
            if true {
                var v19 float32 = (((0.1 * 3.14159) + (3.14159 - ((-2.25 - (((0.0 * 1.5) / 1.0) + 1.5)) - (1e10 * 1.5)))) + 3.14159)
                _ = v19
            } else {
                var v20 int32 = 42
                fmt.Println(v20)
                _ = v20
            }
            if true {
                var v21 uint32 = 2147483648
                for fi22 := 0; fi22 < 4; fi22++ {
                    fmt.Println(v21)
                    _ = fi22
                }
                _ = v21
            }
        }
        v3 = append(v3, "\n\t")
    }
    v23 := GenS2{F0: 10, F1: -974}
    var v24 uint8 = 0
    switch "hello" {
    case "hello":
        switch (("a" + "\n\t") + "") {
        case "foo bar":
            v24 = v24
        case "a":
            for fi25 := 0; fi25 < 6; fi25++ {
                v3 = append(v3, "a")
                var v26 uint64 = 953
                if (!((false || true) && (9223372036854775807 < 1))) {
                    for fi27 := 0; fi27 < 3; fi27++ {
                        if (-2147483648 >= 0) {
                            v26 -= (v26 * 184)
                        } else {
                            v26 += (v26 | v26)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if ((v26 ^ v26) > v26) {
                                var v28 int32 = -323
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                                v26 = (v26 | (^(v26)))
                                _ = v28
                            } else {
                                v29 := v2[("hello" + ("\n\t" + "foo bar"))]
                                if (42 < 0) {
                                    v3 = append(v3, ((((("foo bar" + ("foo bar" + "")) + "\n\t") + "foo bar") + "foo bar") + "a"))
                                    if false {
                                        var v30 uint8 = (v24 << ((255) & 63))
                                        _ = v30
                                    }
                                    v24 -= 1
                                } else {
                                    if true {
                                        v31 := GenS3{F0: 49}
                                        switch "a" {
                                        case "\n\t":
                                            v32 := GenS2{F0: 0, F1: 42}
                                            _ = v32
                                        case "":
                                            v26 += v26
                                        case "a":
                                            v3 = append(v3, "\n\t")
                                        default:
                                            if (1 <= 42) {
                                                fmt.Println(fi27, v26)
                                            }
                                        }
                                        v24 += 17
                                        _ = v31
                                    } else {
                                        if true {
                                            if ("\n\t" >= (("" + "hello") + "a")) {
                                                v26 += (v26 - 42)
                                            }
                                        } else {
                                            var v33 int64 = 42
                                            var v34 string = "a"
                                            _ = v33
                                            _ = v34
                                        }
                                        v26 += (v26 % ((genFn0()) | 1))
                                        if len(v3) > 0 {
                                            v3[0] = ("hello" + "a")
                                        }
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = len(v3)
                                _ = v29
                            }
                            if (true || true) {
                                _ = len(v1)
                            } else {
                                v35 := v1[("" + "hello")]
                                fmt.Println(v35, v26)
                                _ = v35
                            }
                            if (!(-559 >= int64(v26))) {
                                var v36 uint = 564
                                if false {
                                    if ("a" >= ("\n\t" + ("a" + "\n\t"))) {
                                        switch -537 {
                                        case -1:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case -2147483648:
                                            v36 += 65535
                                            v36 += (v36 >> ((v36) & 63))
                                            v37 := GenS4{F0: -650}
                                            _ = v37
                                        default:
                                            if (!false) {
                                                var v38 uint8 = 55
                                                if ((!true) || true) {
                                                    if true {
                                                        if true {
                                                            v36 -= (^(v36))
                                                            var v39 uint64 = uint64(v36)
                                                            switch 127 {
                                                            case -2147483648:
                                                                var v40 uint = v36
                                                                _ = v40
                                                            case 32768:
                                                                v26 -= (v39 / (((v26 + (v26 & genFn0()))) | 1))
                                                            case -9223372036854775808:
                                                                switch (fi27 >> (((v36 * (v36 % (((v36 >> ((4294967295) & 63))) | 1)))) & 63)) {
                                                                case 32767:
                                                                    v26 += 953
                                                                }
                                                            default:
                                                                v39 = (v26 ^ v26)
                                                            }
                                                            v41 := GenS2{F0: 100, F1: -9223372036854775808}
                                                            fmt.Println(v38)
                                                            var v42 bool = (486 <= -1)
                                                            var v43 int32 = -1
                                                            _ = v39
                                                            _ = v41
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            var v44 float32 = 0.1
                                                            _ = v44
                                                        }
                                                        fmt.Println(v26)
                                                        var v45 uint8 = (v24 ^ (v24 & (v24 & v24)))
                                                        _ = v45
                                                    } else {
                                                        v46 := GenS4{F0: -32768}
                                                        if ((-1 > 254) || true) {
                                                            v26 = (v26 % (((v26 * (v26 + v26))) | 1))
                                                            switch 2147483647 {
                                                            case 4294967295:
                                                                var v47 float32 = ((float32(1.5) + float32(0.0)) + (float32(1.5) / ((1.5 / 1.0) * 0.1)))
                                                                _ = v47
                                                            case 42:
                                                                if true {
                                                                    fmt.Println(fi25, v26)
                                                                    switch 2147483647 {
                                                                    case -2147483648:
                                                                        v26 += v26
                                                                    case -993:
                                                                        v48 := GenS2{F0: 22, F1: -1}
                                                                        _ = v48
                                                                    default:
                                                                        var v49 uint = uint(v38)
                                                                        _ = v49
                                                                    }
                                                                    switch "a" {
                                                                    case "":
                                                                        if (-516 < 0) {
                                                                            var v50 uint64 = 42
                                                                            v38 = v38
                                                                            var v51 float32 = -2.25
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                    case "foo bar":
                                                                        if (v24 < (^(v38))) {
                                                                            var v52 uint64 = genFn0()
                                                                            if false {
                                                                                fmt.Println(v24)
                                                                            }
                                                                            var v53 string = "a"
                                                                            _ = v52
                                                                            _ = v53
                                                                        } else {
                                                                            switch ("foo bar" + (("\n\t" + "\n\t") + "hello")) {
                                                                            case "hello":
                                                                                if ((false && (187 >= -1)) || false) {
                                                                                    v26 = 148
                                                                                }
                                                                            default:
                                                                                switch 32767 {
                                                                                case -523:
                                                                                    if false {
                                                                                        v54 := GenS3{F0: -1}
                                                                                        var v55 string = ("foo bar" + ("\n\t" + ((("a" + "hello") + (("" + "a") + "a")) + "\n\t")))
                                                                                        switch fi25 {
                                                                                        case 0:
                                                                                            fmt.Println(fi25, fi25, fi27)
                                                                                        case -1:
                                                                                            var v56 uint = (v36 * v36)
                                                                                            _ = v56
                                                                                        case 127:
                                                                                            v36 -= v36
                                                                                        default:
                                                                                            if false {
                                                                                                _ = len(v1)
                                                                                                var v57 float64 = ((-2.25 * 1e10) - (1.5 - (1e10 / (3.14159 - 1e10))))
                                                                                                if false {
                                                                                                    var v58 uint64 = genFn0()
                                                                                                    _ = v58
                                                                                                }
                                                                                                if (!false) {
                                                                                                    var v59 int8 = 42
                                                                                                    var v60 int16 = -1
                                                                                                    _ = v59
                                                                                                    _ = v60
                                                                                                }
                                                                                                v61 := GenS3{F0: 127}
                                                                                                var v62 float64 = float64(1.5)
                                                                                                _ = v57
                                                                                                _ = v61
                                                                                                _ = v62
                                                                                            } else {
                                                                                                v63 := GenS5{F0: ((v36 % (((v36 / ((v36) | 1))) | 1)) <= v36)}
                                                                                                var v64 string = "\n\t"
                                                                                                fmt.Println(v26)
                                                                                                _ = v63
                                                                                                _ = v64
                                                                                            }
                                                                                        }
                                                                                        switch v55 {
                                                                                        case "foo bar":
                                                                                            var v65 int32 = 42
                                                                                            _ = v65
                                                                                        default:
                                                                                            v66 := GenS4{F0: -998}
                                                                                            _ = v66
                                                                                        }
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                    }
                                                                                case 128:
                                                                                    var v67 uint16 = 9
                                                                                    _ = v67
                                                                                case -2147483648:
                                                                                    if false {
                                                                                        v36 += v36
                                                                                        fmt.Println(v26)
                                                                                        v68 := GenS3{F0: int8(v24)}
                                                                                        _ = v68
                                                                                    }
                                                                                default:
                                                                                    v24 -= 163
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        if (-966 >= 881) {
                                                                            v26 = (v26 >> (((v36 % ((v36) | 1))) & 63))
                                                                            v69 := GenS2{F0: -1, F1: 448}
                                                                            fmt.Println(fi27, v36)
                                                                            _ = v69
                                                                        } else {
                                                                            switch "a" {
                                                                            case "":
                                                                                var v70 uint64 = v26
                                                                                _ = v70
                                                                            case "foo bar":
                                                                                v26 -= (^(v26))
                                                                            case "hello":
                                                                                v24 -= 99
                                                                            }
                                                                            v26 = genFn0()
                                                                            v38 = 116
                                                                        }
                                                                    }
                                                                    v36 = 255
                                                                }
                                                            default:
                                                                if false {
                                                                    if (-1 != 1000) {
                                                                        var v71 bool = false
                                                                        var v72 int64 = 9223372036854775807
                                                                        switch fi27 {
                                                                        case 433:
                                                                            fmt.Println(fi25, v71)
                                                                        default:
                                                                            var v73 uint16 = 32767
                                                                            _ = v73
                                                                        }
                                                                        _ = v71
                                                                        _ = v72
                                                                    } else {
                                                                        v74 := GenS5{F0: ("foo bar" > "")}
                                                                        _ = v74
                                                                    }
                                                                }
                                                            }
                                                            v75 := GenS3{F0: -128}
                                                            v76 := v1[""]
                                                            v24 -= (v38 * v38)
                                                            if true {
                                                                var v77 int16 = 1000
                                                                if v76 {
                                                                    switch (fi27 & 9223372036854775807) {
                                                                    case 32768:
                                                                        if false {
                                                                            var v78 string = ("a" + "a")
                                                                            v24 += v24
                                                                            v26 += genFn0()
                                                                            if (600 < 2147483647) {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = "\n\t"
                                                                                }
                                                                                v38 = (^(v24))
                                                                            }
                                                                            var v79 string = "foo bar"
                                                                            fmt.Println(fi25, v78)
                                                                            var v80 int32 = -606
                                                                            _ = v78
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                    if (("foo bar" <= "\n\t") && (!v76)) {
                                                                        v77 = 975
                                                                        v36 = (v36 - (^(v36)))
                                                                        v81 := GenS2{F0: 1, F1: 42}
                                                                        _ = v81
                                                                    } else {
                                                                        v82 := GenS6{F0: v77, F1: 1}
                                                                        v38 = (^(v38))
                                                                        v36 = (v36 & 0)
                                                                        var v83 int8 = int8(fi25)
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                    v84 := GenS7{F0: -1}
                                                                    var v85 string = "foo bar"
                                                                    v77 -= v77
                                                                    switch fi25 {
                                                                    case 1:
                                                                        v24 = (^(v38))
                                                                    default:
                                                                        _ = len(v2)
                                                                    }
                                                                    var v86 uint16 = 637
                                                                    v24 -= 254
                                                                    _ = v84
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                if v76 {
                                                                    var v87 string = ""
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v77 = v77
                                                                    }
                                                                    _ = v87
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                if ("" == "") {
                                                                    if (!(v76 || v76)) {
                                                                        v88 := GenS6{F0: v77, F1: 1}
                                                                        v1["a"] = v76
                                                                        var v89 uint8 = (^(v24))
                                                                        v26 = (v26 >> ((2) & 63))
                                                                        _ = v88
                                                                        _ = v89
                                                                    } else {
                                                                        v26 = genFn0()
                                                                    }
                                                                    var v90 float64 = 3.14159
                                                                    switch fi27 {
                                                                    case 42:
                                                                        switch (fi25 / ((127) | 1)) {
                                                                        case -258:
                                                                            switch "hello" {
                                                                            case "":
                                                                                v90 = 3.14159
                                                                            case "a":
                                                                                delete(v1, "hello")
                                                                            case "foo bar":
                                                                                _ = len(v1)
                                                                            }
                                                                        default:
                                                                            v91 := GenS4{F0: v77}
                                                                            _ = v91
                                                                        }
                                                                        v92 := GenS8{F0: 858}
                                                                        _ = v92
                                                                    case -419:
                                                                        fmt.Println(v77, v76, v24)
                                                                    case -9223372036854775808:
                                                                        var v93 uint32 = 0
                                                                        _ = v93
                                                                    }
                                                                    _ = v90
                                                                } else {
                                                                    v26 = (v26 / ((genFn0()) | 1))
                                                                    if true {
                                                                        var v94 int8 = 0
                                                                        switch int(v77) {
                                                                        case 329:
                                                                            if v76 {
                                                                                var v95 uint64 = v26
                                                                                _ = v95
                                                                            }
                                                                        case 1:
                                                                            v96 := GenS9{F0: 974, F1: -1, F2: (fi25 + fi27)}
                                                                            _ = v96
                                                                        }
                                                                        delete(v1, "foo bar")
                                                                        v26 = v26
                                                                        var v97 int64 = -639
                                                                        _ = v94
                                                                        _ = v97
                                                                    } else {
                                                                        v76 = (v76 && (!true))
                                                                        v24 += v38
                                                                    }
                                                                }
                                                                _ = v77
                                                            } else {
                                                                fmt.Println(v26)
                                                            }
                                                            var v98 float64 = (0.0 / 1.5)
                                                            v99 := GenS10{F0: v38, F1: v24}
                                                            var v100 float64 = (float64((((0.0 / 1e10) + float32(1e10)) / 3.14159)) * float64(-2.25))
                                                            _ = v100
                                                            _ = v75
                                                            _ = v76
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                        _ = v46
                                                    }
                                                } else {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        switch "a" {
                                                        case "foo bar":
                                                            if (1 >= (v38 >> (((v36 & 844)) & 63))) {
                                                                if true {
                                                                    v101 := GenS9{F0: 534, F1: 0, F2: fi25}
                                                                    var v102 float32 = 1.5
                                                                    _ = v101
                                                                    _ = v102
                                                                } else {
                                                                    var v103 int64 = 672
                                                                    if (genFn0() >= (v26 | 1)) {
                                                                        v103 += (v103 - 9223372036854775807)
                                                                        switch fi25 {
                                                                        case 0:
                                                                            v104 := GenS7{F0: 590}
                                                                            _ = v104
                                                                        case 1:
                                                                            if true {
                                                                                var v105 string = ("\n\t" + "foo bar")
                                                                                v103 = (v103 + (v103 ^ (v103 % ((42) | 1))))
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = ""
                                                                                }
                                                                                v106 := GenS2{F0: -128, F1: -9223372036854775808}
                                                                                var v107 int8 = 93
                                                                                _ = v105
                                                                                _ = v106
                                                                                _ = v107
                                                                            } else {
                                                                                var v108 string = "a"
                                                                                var v109 int8 = 24
                                                                                fmt.Println(v26, fi27)
                                                                                _ = v108
                                                                                _ = v109
                                                                            }
                                                                        case 2147483647:
                                                                            v24 = v24
                                                                        default:
                                                                            if ((v103 >> ((v36) & 63)) >= 0) {
                                                                                v110 := GenS11{F0: 56, F1: 9223372036854775807}
                                                                                _ = len(v1)
                                                                                _ = v110
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v26, fi25, v36)
                                                                    _ = v103
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v3 = append(v3, "a")
                                                        default:
                                                            var v111 int = (fi25 - (^(fi25)))
                                                            _ = v111
                                                        }
                                                    case "hello":
                                                        v112 := GenS2{F0: -1, F1: 0}
                                                        _ = v112
                                                    }
                                                    var v113 string = "hello"
                                                    v38 -= v24
                                                    _ = v113
                                                }
                                                v26 = (v26 + 4294967295)
                                                var v114 float64 = 0.1
                                                _ = v114
                                                _ = v38
                                            }
                                        }
                                        var v115 float32 = 3.14159
                                        _ = v115
                                    } else {
                                        v26 = v26
                                        var v116 uint = v36
                                        v117 := GenS7{F0: -9223372036854775808}
                                        _ = v116
                                        _ = v117
                                    }
                                }
                                _ = v36
                            }
                        }
                        _ = fi27
                    }
                    switch fi25 {
                    case 1:
                        if ((v24 + 2) >= (^(v24))) {
                            _ = len(v2)
                            v118 := GenS6{F0: int16(fi25), F1: 255}
                            switch "a" {
                            case "hello":
                                switch (("" + "\n\t") + "a") {
                                case "a":
                                    var v119 uint = 0
                                    _ = v119
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch "hello" {
                                        case "hello":
                                            var v120 int64 = 42
                                            _ = v120
                                        case "":
                                            var v121 float64 = 1.5
                                            _ = v121
                                        case "a":
                                            switch 547 {
                                            case 9223372036854775807:
                                                for fi122 := 0; fi122 < 14; fi122++ {
                                                    v24 = v24
                                                    if (-9223372036854775808 <= 9223372036854775807) {
                                                        var v123 int = -9223372036854775808
                                                        _ = v123
                                                    } else {
                                                        var v124 bool = (!(float64(0.1) != 1.5))
                                                        _ = v124
                                                    }
                                                    _ = fi122
                                                }
                                            default:
                                                var v125 uint64 = (^(v26))
                                                _ = v125
                                            }
                                        default:
                                            for fi126 := 0; fi126 < 7; fi126++ {
                                                if (!(0.0 < 3.14159)) {
                                                    v127 := GenS4{F0: 1}
                                                    if false {
                                                        v128 := GenS12{F0: (-(fi25))}
                                                        var v129 float32 = -2.25
                                                        var v130 uint64 = 42
                                                        _ = v128
                                                        _ = v129
                                                        _ = v130
                                                    }
                                                    var v131 int64 = 776
                                                    if (false && true) {
                                                        v24 = v24
                                                        v131 = 42
                                                        var v132 uint64 = uint64(v24)
                                                        _ = v132
                                                    } else {
                                                        fmt.Println(fi25, fi25, fi25)
                                                        if (((^(v24)) > 84) && (4294967295 >= 18446744073709551615)) {
                                                            v24 -= 121
                                                        } else {
                                                            switch ("\n\t" + ("foo bar" + "hello")) {
                                                            case "hello":
                                                                v131 -= (v131 / (((v131 ^ -510)) | 1))
                                                            case "":
                                                                switch "\n\t" {
                                                                case "":
                                                                    if true {
                                                                        v131 = (v131 % ((v131) | 1))
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                if false {
                                                                    if true {
                                                                        v133 := GenS6{F0: -32768, F1: 167}
                                                                        _ = v133
                                                                    } else {
                                                                        fmt.Println(fi25)
                                                                        v26 -= v26
                                                                        if true {
                                                                            fmt.Println(v26, v24)
                                                                            var v134 uint = 2
                                                                            _ = v134
                                                                        }
                                                                    }
                                                                    v24 -= uint8(v26)
                                                                }
                                                            default:
                                                                v26 += v26
                                                            }
                                                            if false {
                                                                if ((1 == 1) && false) {
                                                                    switch (fi25 | (fi126 - (fi25 << ((uint(v26)) & 63)))) {
                                                                    case 127:
                                                                        var v135 uint32 = 505
                                                                        _ = v135
                                                                    case -474:
                                                                        v1[(("" + "hello") + "hello")] = true
                                                                    }
                                                                } else {
                                                                    if (0.1 == (-2.25 / (-2.25 + -2.25))) {
                                                                        var v136 int16 = -32768
                                                                        v137 := GenS4{F0: (v136 - -357)}
                                                                        var v138 int = (fi25 + fi126)
                                                                        _ = v136
                                                                        _ = v137
                                                                        _ = v138
                                                                    } else {
                                                                        _ = len(v3)
                                                                        var v139 float64 = 0.0
                                                                        var v140 uint = 13
                                                                        _ = v139
                                                                        _ = v140
                                                                    }
                                                                }
                                                                delete(v1, "")
                                                                v24 = v24
                                                            } else {
                                                                var v141 int16 = 32767
                                                                _ = v141
                                                            }
                                                            _ = len(v3)
                                                        }
                                                    }
                                                    v26 += 2
                                                    fmt.Println(fi126)
                                                    var v142 int = fi126
                                                    _ = v127
                                                    _ = v131
                                                    _ = v142
                                                }
                                                _ = fi126
                                            }
                                        }
                                    case "":
                                        v24 += (v24 | v24)
                                    default:
                                        var v143 int16 = -1
                                        _ = v143
                                    }
                                }
                            case "":
                                v24 += (v24 | v24)
                            case "foo bar":
                                for fi144 := 0; fi144 < 2; fi144++ {
                                    v26 -= (v26 << ((42) & 63))
                                    if (42 == 42) {
                                        v145 := GenS8{F0: 1}
                                        fmt.Println(fi144, v24, v24)
                                        _ = v145
                                    }
                                    _ = fi144
                                }
                            default:
                                if true {
                                    v24 = v24
                                } else {
                                    for fi146 := 0; fi146 < 17; fi146++ {
                                        fmt.Println(fi146, fi25)
                                        var v147 int = 1
                                        fmt.Println(v147, fi146, v24)
                                        v26 = 42
                                        _ = fi146
                                        _ = v147
                                    }
                                }
                            }
                            _ = v118
                        } else {
                            switch 32768 {
                            case 128:
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                            default:
                                v26 -= 0
                            }
                            if (100 <= 3) {
                                v148 := GenS13{F0: (("" + ("\n\t" + "")) + "\n\t")}
                                fmt.Println(v26, v26, v26)
                                _ = v148
                            } else {
                                fmt.Println(v26)
                                var v149 int64 = -9223372036854775808
                                _ = v149
                            }
                        }
                    case 778:
                        fmt.Println(v26, v26, v24)
                    default:
                        var v150 float32 = float32(1e10)
                        _ = v150
                    }
                } else {
                    if false {
                        if true {
                            fmt.Println(v24, fi25, fi25)
                            v24 += (^(v24))
                        }
                        _ = len(v1)
                        if ("" > "foo bar") {
                            var v151 int16 = 42
                            _ = len(v1)
                            _ = v151
                        } else {
                            for fi152 := 0; fi152 < 19; fi152++ {
                                var v153 int = fi25
                                v26 += v26
                                _ = fi152
                                _ = v153
                            }
                            var v154 uint32 = 2
                            _ = v154
                        }
                    } else {
                        fmt.Println(v26, v26, v24)
                        _ = len(v3)
                        v26 += v26
                    }
                    switch "a" {
                    case "foo bar":
                        if (!(true || false)) {
                            switch (("\n\t" + "\n\t") + "") {
                            case "a":
                                v155 := GenS8{F0: 1}
                                _ = v155
                            }
                        }
                    }
                    for fi156 := 0; fi156 < 16; fi156++ {
                        v157 := GenS12{F0: (fi25 * 0)}
                        v3 = append(v3, "")
                        v24 += 0
                        _ = fi156
                        _ = v157
                    }
                }
                _ = fi25
                _ = v26
            }
        case "":
            for fi158 := 0; fi158 < 13; fi158++ {
                for fi159 := 0; fi159 < 18; fi159++ {
                    if false {
                        switch (fi158 >> ((51) & 63)) {
                        case 9223372036854775807:
                            switch "" {
                            case "":
                                v160 := v2["hello"]
                                _ = v160
                            }
                        case 42:
                            v24 += 128
                        }
                        var v161 uint8 = v24
                        v162 := GenS14{F0: 4294967295, F1: 473}
                        _ = v161
                        _ = v162
                    }
                    _ = fi159
                }
                _ = fi158
            }
        default:
            var v163 uint = 939
            _ = v163
        }
    default:
        v24 = (v24 ^ (^(v24)))
    }
    if (true && ("\n\t" <= (((("foo bar" + (("hello" + "\n\t") + "")) + "a") + "") + "\n\t"))) {
        if (true || ((!true) || (true && true))) {
            v164 := GenS4{F0: 0}
            switch 426 {
            case -1:
                var v165 uint32 = 380
                _ = v165
            }
            var v166 uint = 28
            _ = v164
            _ = v166
        } else {
            v24 = 42
            v24 = 255
            var v167 uint8 = v24
            _ = v167
        }
    }
    v168 := GenS7{F0: 382}
    v169 := GenS15{F0: 0, F1: 1}
    for fi170 := 0; fi170 < 2; fi170++ {
        var v171 int8 = -1
        switch "hello" {
        case "":
            if false {
                var v172 uint = 0
                _ = v172
            } else {
                v173 := GenS16{F0: fi170}
                _ = v173
            }
        default:
            _ = len(v3)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi170
        _ = v171
    }
    switch -9223372036854775808 {
    case 1:
        for fi174 := 0; fi174 < 18; fi174++ {
            v24 = v24
            _ = fi174
        }
    case 42:
        v175 := GenS1{F0: (^(v24))}
        _ = v175
    case -966:
        for fi176 := 0; fi176 < 2; fi176++ {
            fmt.Println(fi176, v24)
            _ = fi176
        }
    }
    fmt.Println(v24)
    var v177 uint64 = 4294967295
    v177 = (v177 - v177)
    if true {
        v178 := GenS10{F0: (v24 ^ v24), F1: (v24 + (^(v24)))}
        _ = v178
    }
    var v179 uint32 = uint32(v177)
    v24 -= (^(v24))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v177:", v177)
    fmt.Println("v179:", v179)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v168
    _ = v169
    _ = v177
    _ = v179
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
