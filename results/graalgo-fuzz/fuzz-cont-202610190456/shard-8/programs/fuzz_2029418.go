package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 uint64; F1 int; F2 int }
type GenS9 struct { F0 uint64; F1 uint8; F2 bool }
type GenS10 struct { F0 int16; F1 int64 }
type GenS11 struct { F0 int; F1 int16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint; F1 int16 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint64 = 42
        fmt.Println(fi6, fi6)
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 128
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(fi10, fi9, v8)
            fmt.Println(v8, v8, fi10)
            _ = fi10
        }
        var v11 bool = false
        for fi12 := 0; fi12 < 10; fi12++ {
            switch "foo bar" {
            case "hello":
                v11 = v11
            case "a":
                var v13 bool = v11
                _ = v13
            default:
                var v14 float64 = 1e10
                v15 := []int{fi12, 837}
                _ = v14
                _ = v15
            }
            fmt.Println(v11)
            fmt.Println(fi9, fi12)
            var v16 float64 = 0.1
            _ = fi12
            _ = v16
        }
        _ = fi9
        _ = v11
    }
    _ = v8
    return 42
}
func genFn1(p0 int, p1 string) uint64 {
    p1 = ("hello" + "")
    switch (-(p0)) {
    case -9223372036854775808:
        var v17 int8 = -97
        _ = v17
    case 870:
        if false {
            var v18 uint64 = 0
            p0 = (^(p0))
            _ = v18
        } else {
            if (42 <= 805) {
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 uint32 = 0
                    p1 = "foo bar"
                    for fi21 := 0; fi21 < 12; fi21++ {
                        fmt.Println(fi21, v20)
                        if false {
                            if false {
                                v22 := GenS1{F0: 0, F1: true, F2: -1}
                                switch (p1 + "a") {
                                case "":
                                    v23 := []int{fi21, 42, (p0 + (p0 << ((1) & 63)))}
                                    _ = v23
                                }
                                _ = v22
                            }
                        }
                        _ = fi21
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                fmt.Println(p0, p0)
                for fi24 := 0; fi24 < 11; fi24++ {
                    v25 := GenS1{F0: -1, F1: (-275 >= 86), F2: 9223372036854775807}
                    _ = fi24
                    _ = v25
                }
                switch p1 {
                case "\n\t":
                    v26 := GenS2{F0: (!false)}
                    _ = v26
                case "hello":
                    v27 := GenS3{F0: 1, F1: 9223372036854775807, F2: true}
                    _ = v27
                }
            }
            switch (p0 - (p0 ^ p0)) {
            case -2147483648:
                v28 := GenS3{F0: 141, F1: 767, F2: (1.5 < (1.5 + ((float32((1e10 / -2.25)) + (float32(-2.25) + (((1.5 * float32(3.14159)) * ((-2.25 + float32(0.1)) * (0.1 - ((0.0 / 0.1) * float32((((-2.25 + 1e10) - 0.0) * -2.25)))))) * 3.14159))) * 1e10)))}
                _ = v28
            case 127:
                p0 = -474
            }
        }
    case 128:
        v29 := GenS3{F0: genFn0(), F1: 0, F2: (1 >= 1)}
        _ = v29
    default:
        var v30 int8 = 100
        _ = v30
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        p1 = p1
        _ = fi31
    }
    var v32 int64 = 1
    _ = p0
    _ = p1
    _ = v32
    return 4294967295
}

func main() {
    v1 := []string{"foo bar", "a", "foo bar"}
    var v2 uint64 = 42
    var v3 uint64 = (v2 % ((v2) | 1))
    var v4 uint32 = 2147483647
    var v5 int = -685
    if true {
        for fi33 := 0; fi33 < 3; fi33++ {
            v3 -= (^(v2))
            _ = fi33
        }
        var v34 int32 = -847
        _ = v34
    } else {
        var v35 string = ""
        _ = v35
    }
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "a")
    }
    var v36 int8 = -122
    fmt.Println(v3, v36)
    for fi37 := 0; fi37 < 5; fi37++ {
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 int8 = v36
            switch 42 {
            case -9223372036854775808:
                var v40 int32 = 42
                _ = v40
            case -2147483648:
                if (true || false) {
                    var v41 int32 = 0
                    switch 2147483647 {
                    case 32768:
                        var v42 bool = (127 > 46)
                        _ = v42
                    default:
                        v3 += v2
                    }
                    v2 += 897
                    _ = v41
                } else {
                    fmt.Println(v39, v2, v3)
                    switch "a" {
                    case "\n\t":
                        v43 := GenS2{F0: (1 <= -1)}
                        _ = v43
                    case "hello":
                        v44 := GenS2{F0: (-9223372036854775808 <= 1)}
                        _ = v44
                    case "":
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    }
                    fmt.Println(v3, v4)
                }
            case 171:
                v1 = append(v1, "a")
                switch (fi38 + 99) {
                case 448:
                    fmt.Println(fi38, v39, v2)
                case 42:
                    v45 := GenS4{F0: 196, F1: "foo bar"}
                    if (!true) {
                        if (0 > 255) {
                            var v46 int8 = (v36 | v36)
                            _ = v46
                        }
                        if (-1 == 748) {
                            if (-1 > -2147483648) {
                                if (((fi38 << ((0) & 63)) < v5) && (-2.25 != -2.25)) {
                                    fmt.Println(v2, fi37)
                                } else {
                                    v36 -= v39
                                }
                                var v47 int32 = 0
                                fmt.Println(v3, v39, fi38)
                                v48 := GenS5{F0: 255}
                                switch v5 {
                                case 0:
                                    if (1.5 >= 1e10) {
                                        fmt.Println(v39, v2)
                                        switch "a" {
                                        case "a":
                                            if (0 == v47) {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v49 uint16 = 32767
                                                    _ = v49
                                                default:
                                                    _ = len(v1)
                                                }
                                                switch "hello" {
                                                case "hello":
                                                    var v50 string = "foo bar"
                                                    _ = v50
                                                case "\n\t":
                                                    v51 := GenS3{F0: 527, F1: -342, F2: true}
                                                    _ = v51
                                                case "foo bar":
                                                    if true {
                                                        if true {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v52 int8 = (v39 >> ((727) & 63))
                                                            v47 = (v47 >> ((1) & 63))
                                                            _ = v52
                                                        }
                                                    }
                                                }
                                                v53 := GenS2{F0: false}
                                                _ = v53
                                            } else {
                                                if (0.0 < 1.5) {
                                                    v54 := GenS5{F0: 1}
                                                    if (false || false) {
                                                        v47 = v47
                                                        fmt.Println(fi37, v5, v5)
                                                    }
                                                    _ = v54
                                                }
                                                fmt.Println(v3, v36, v36)
                                            }
                                        case "\n\t":
                                            fmt.Println(v36, v2, v5)
                                        case "hello":
                                            if (v47 == (v47 ^ (v47 << ((65535) & 63)))) {
                                                if false {
                                                    v55 := GenS4{F0: 128, F1: ""}
                                                    var v56 int64 = 0
                                                    v5 += (fi37 * -9223372036854775808)
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    if (1 > 999) {
                                                        v57 := GenS6{F0: -1}
                                                        _ = v57
                                                    }
                                                    if false {
                                                        switch -266 {
                                                        case -2147483648:
                                                            var v58 bool = (!(1.5 < (3.14159 / (1e10 * (float64(-2.25) / 3.14159)))))
                                                            _ = v58
                                                        case 4294967295:
                                                            if true {
                                                                var v59 uint16 = 889
                                                                _ = v59
                                                            } else {
                                                                v60 := GenS7{F0: (v3 << ((358) & 63)), F1: (v4 - (v4 % ((v4) | 1)))}
                                                                _ = v60
                                                            }
                                                        }
                                                        fmt.Println(fi37, v4)
                                                    } else {
                                                        v61 := GenS2{F0: false}
                                                        if (!(32767 > -32768)) {
                                                            v62 := GenS5{F0: 18446744073709551615}
                                                            fmt.Println(v3, fi37, v4)
                                                            v3 = (v2 - (v3 * 42))
                                                            v47 += v47
                                                            switch 128 {
                                                            case -1:
                                                                switch (fi37 * (v5 >> ((1) & 63))) {
                                                                case -577:
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    if (((266 != v4) && (true || false)) && ((1.5 / (1.5 - (((((float64((3.14159 / 0.1)) * (0.0 - -2.25)) + 1e10) - 0.1) + 1e10) / float64((-2.25 / (0.1 * (0.1 / 1e10))))))) <= 0.1)) {
                                                                        switch ("\n\t" + ("hello" + ((((("" + (("\n\t" + (((("a" + "\n\t") + "\n\t") + "foo bar") + "hello")) + "")) + ("foo bar" + "foo bar")) + ("a" + (("\n\t" + (("a" + (("hello" + "") + "a")) + "\n\t")) + (("\n\t" + "hello") + ("\n\t" + "hello"))))) + ("hello" + ("\n\t" + ""))) + ("foo bar" + "\n\t")))) {
                                                                        case "":
                                                                            v63 := GenS8{F0: 42, F1: (v5 ^ fi38), F2: fi37}
                                                                            _ = v63
                                                                        }
                                                                    }
                                                                case 857:
                                                                    v36 -= 0
                                                                }
                                                            case 1:
                                                                v3 = (^(v3))
                                                            }
                                                            _ = v62
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = ("a" + "")
                                                        }
                                                        var v64 int32 = -736
                                                        if (!(32767 <= -32768)) {
                                                            var v65 float64 = 3.14159
                                                            _ = v65
                                                        }
                                                        _ = v61
                                                        _ = v64
                                                    }
                                                    switch ("hello" + ("a" + "a")) {
                                                    case "\n\t":
                                                        v5 += fi37
                                                    case "a":
                                                        if true {
                                                            if (!(true || true)) {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    var v66 float32 = ((3.14159 / 1.5) - 1.5)
                                                                    _ = v66
                                                                case "":
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                case "a":
                                                                    var v67 uint32 = 2
                                                                    _ = v67
                                                                }
                                                                v4 += 2147483647
                                                                v47 = v47
                                                                v68 := GenS9{F0: 0, F1: 158, F2: false}
                                                                _ = v68
                                                            }
                                                            if true {
                                                                v69 := GenS9{F0: (^(v3)), F1: 42, F2: true}
                                                                v39 += (v36 >> ((255) & 63))
                                                                _ = v69
                                                            }
                                                            v36 = v39
                                                            var v70 float64 = 1.5
                                                            _ = v70
                                                        }
                                                    case "hello":
                                                        v1 = append(v1, "hello")
                                                    default:
                                                        switch v5 {
                                                        case -852:
                                                            switch "" {
                                                            case "foo bar":
                                                                switch v5 {
                                                                case -9223372036854775808:
                                                                    _ = len(v1)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "hello"
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                case 9223372036854775807:
                                                                    if (-1 < -9223372036854775808) {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v2 += uint64(v4)
                                                                    } else {
                                                                        var v71 uint64 = v3
                                                                        _ = v71
                                                                    }
                                                                default:
                                                                    var v72 int64 = 583
                                                                    _ = v72
                                                                }
                                                            case "hello":
                                                                if false {
                                                                    var v73 int32 = (v47 ^ (v47 * v47))
                                                                    v74 := GenS5{F0: 42}
                                                                    fmt.Println(fi37, v3)
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                            case "a":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            default:
                                                                v2 -= (v2 | 18446744073709551615)
                                                            }
                                                        case 0:
                                                            if true {
                                                                fmt.Println(v47, v47, v39)
                                                                v5 -= 127
                                                                var v75 int16 = -891
                                                                var v76 uint16 = 390
                                                                _ = v75
                                                                _ = v76
                                                            } else {
                                                                if (2 < 1) {
                                                                    var v77 uint32 = v4
                                                                    _ = v77
                                                                }
                                                                var v78 uint = uint(v4)
                                                                _ = v78
                                                            }
                                                        }
                                                        v1 = append(v1, (("foo bar" + ("" + "a")) + "\n\t"))
                                                    }
                                                }
                                                v5 = (fi38 - fi38)
                                            } else {
                                                switch "a" {
                                                case "\n\t":
                                                    v79 := GenS5{F0: 4294967295}
                                                    _ = v79
                                                case "hello":
                                                    fmt.Println(v47, v4, fi38)
                                                }
                                            }
                                        default:
                                            if (false && true) {
                                                var v80 uint64 = v3
                                                if (0.1 > (-2.25 - ((0.1 * -2.25) - 1.5))) {
                                                    var v81 bool = (!false)
                                                    _ = v81
                                                } else {
                                                    var v82 float32 = (1.5 - 3.14159)
                                                    var v83 bool = (!true)
                                                    var v84 int16 = int16(v39)
                                                    _ = v82
                                                    _ = v83
                                                    _ = v84
                                                }
                                                _ = v80
                                            }
                                        }
                                    } else {
                                        v85 := GenS8{F0: (v2 - (v2 - v3)), F1: 509, F2: (v5 << ((994) & 63))}
                                        var v86 uint = 0
                                        var v87 int64 = -9223372036854775808
                                        var v88 uint16 = 0
                                        v36 = -128
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    }
                                case 649:
                                    v89 := GenS10{F0: 32767, F1: 0}
                                    _ = v89
                                }
                                v3 -= v3
                                _ = v47
                                _ = v48
                            }
                            var v90 float64 = (1e10 * (1.5 + ((1.5 - 0.1) / 0.1)))
                            _ = v90
                        } else {
                            var v91 uint8 = 93
                            v5 = v5
                            v1 = append(v1, "foo bar")
                            _ = v91
                        }
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    }
                    if (((!((true && true) || false)) || (65535 == 946)) && true) {
                        v36 += -32
                    }
                    fmt.Println(fi38)
                    v92 := GenS11{F0: 42, F1: -622}
                    var v93 int16 = 498
                    _ = v45
                    _ = v92
                    _ = v93
                case 0:
                    if (fi38 < fi37) {
                        var v94 uint32 = (v4 >> ((4294967295) & 63))
                        if len(v1) > 0 {
                            v1[0] = ("hello" + "foo bar")
                        }
                        fmt.Println(fi37)
                        _ = v94
                    }
                }
            }
            var v95 float32 = float32((0.0 / 3.14159))
            _ = fi38
            _ = v39
            _ = v95
        }
        _ = fi37
    }
    if len(v1) > 0 {
        v1[0] = ((("\n\t" + "\n\t") + ("a" + ("foo bar" + "hello"))) + "a")
    }
    if false {
        switch 127 {
        case 9223372036854775807:
            var v96 string = "\n\t"
            _ = v96
        case -9223372036854775808:
            var v97 int64 = 1
            _ = v97
        case 0:
            var v98 int32 = -2147483648
            _ = v98
        default:
            var v99 uint = 4294967295
            _ = v99
        }
        v100 := GenS12{F0: true}
        _ = v100
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi101 := 0; fi101 < 16; fi101++ {
        v1 = append(v1, "foo bar")
        for fi102 := 0; fi102 < 16; fi102++ {
            var v103 bool = (2147483647 > -1)
            if true {
                var v104 int = 0
                var v105 int16 = 1000
                var v106 int32 = 0
                _ = v104
                _ = v105
                _ = v106
            } else {
                var v107 uint32 = 1
                v3 += (v2 | v2)
                fmt.Println(fi101)
                v107 = genFn0()
                if false {
                    v3 = v2
                    switch 42 {
                    case 42:
                        var v108 int = fi102
                        _ = v108
                    default:
                        var v109 int32 = 2147483647
                        _ = v109
                    }
                    v110 := GenS6{F0: 0}
                    _ = v110
                } else {
                    if len(v1) > 0 {
                        v1[0] = ("hello" + ("a" + "foo bar"))
                    }
                    var v111 float32 = 0.1
                    switch ("foo bar" + ("\n\t" + ("" + "foo bar"))) {
                    case "foo bar":
                        v112 := GenS11{F0: (^(fi102)), F1: -32768}
                        _ = v112
                    case "\n\t":
                        switch (v5 >> ((4294967295) & 63)) {
                        case 127:
                            var v113 float64 = 1e10
                            _ = v113
                        case -596:
                            if (!(!(32767 <= -1))) {
                                if ((v103 || v103) && true) {
                                    if (32767 > -1) {
                                        v107 -= genFn0()
                                        if len(v1) > 0 {
                                            v1[0] = ("\n\t" + "\n\t")
                                        }
                                    } else {
                                        v103 = v103
                                        var v114 int = -902
                                        _ = v114
                                    }
                                } else {
                                    v3 += 42
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        v115 := GenS12{F0: (-1 < -510)}
                                        _ = v115
                                    case "foo bar":
                                        var v116 int16 = 0
                                        _ = v116
                                    case "\n\t":
                                        fmt.Println(v107)
                                    }
                                }
                                v111 += 3.14159
                            } else {
                                v117 := GenS13{F0: 65535}
                                _ = v117
                            }
                        case 2147483647:
                            v118 := GenS14{F0: 255, F1: -1}
                            _ = v118
                        default:
                            var v119 int64 = 9223372036854775807
                            _ = v119
                        }
                    }
                    v111 += 0.0
                    if ("" != (("foo bar" + "foo bar") + "")) {
                        if false {
                            v120 := GenS4{F0: 9, F1: "a"}
                            var v121 int8 = (v36 << ((936) & 63))
                            _ = v120
                            _ = v121
                        }
                        v5 -= (v5 & fi101)
                        var v122 float32 = v111
                        _ = v122
                    }
                    _ = v111
                }
                _ = v107
            }
            var v123 uint = 4294967295
            _ = fi102
            _ = v103
            _ = v123
        }
        _ = fi101
    }
    for fi124 := 0; fi124 < 2; fi124++ {
        fmt.Println(v2, fi124)
        _ = fi124
    }
    switch ("hello" + "a") {
    case "":
        fmt.Println(v5, v5, v2)
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v125 int32 = 2147483647
        _ = v125
    case "foo bar":
        v5 -= 320
    case "a":
        for fi126 := 0; fi126 < 6; fi126++ {
            if false {
                v1 = append(v1, ("a" + "foo bar"))
                fmt.Println(v36, v5, v4)
                v3 = (v3 / ((v2) | 1))
            } else {
                v127 := GenS3{F0: v4, F1: -9223372036854775808, F2: (true && true)}
                _ = v127
            }
            _ = fi126
        }
    default:
        var v128 uint = 545
        _ = v128
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
