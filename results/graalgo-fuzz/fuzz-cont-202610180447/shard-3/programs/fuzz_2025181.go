package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 bool; F2 int64 }
type GenS3 struct { F0 int; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 uint32; F2 string }
type GenS5 struct { F0 uint8; F1 uint16; F2 bool }
type GenS6 struct { F0 int8; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int8; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS9 struct { F0 int64; F1 int16; F2 int64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32; F1 uint }
type GenS12 struct { F0 uint }
func genFn0(p0 uint, p1 string) int8 {
    if true {
        v5 := GenS1{F0: "a", F1: 4294967295, F2: 42}
        if (1 == 722) {
            v6 := GenS1{F0: ("\n\t" + p1), F1: uint32(p0), F2: 1}
            if (42 == 18446744073709551615) {
                for fi7 := 0; fi7 < 12; fi7++ {
                    v8 := []int{}
                    for fi9 := 0; fi9 < 16; fi9++ {
                        switch (p1 + p1) {
                        case "a":
                            var v10 float32 = 1e10
                            _ = v10
                        default:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                        _ = fi9
                    }
                    p1 = "a"
                    _ = fi7
                    _ = v8
                }
            }
            _ = v6
        } else {
            if true {
                v11 := GenS1{F0: p1, F1: 672, F2: 255}
                _ = v11
            } else {
                v12 := GenS2{F0: 127, F1: (-32768 <= 1), F2: 265}
                _ = v12
            }
            for fi13 := 0; fi13 < 2; fi13++ {
                for fi14 := 0; fi14 < 1; fi14++ {
                    p0 = p0
                    v15 := GenS2{F0: 0, F1: true, F2: 116}
                    var v16 bool = true
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                fmt.Println(fi13)
                fmt.Println(fi13, p0)
                _ = fi13
            }
        }
        _ = v5
    } else {
        p0 = p0
        fmt.Println(p1, p0)
        for fi17 := 0; fi17 < 11; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                fmt.Println(fi18, fi17, fi18)
                var v19 int16 = -1
                p0 -= p0
                _ = fi18
                _ = v19
            }
            if (!(!(!true))) {
                var v20 float64 = ((1e10 + ((1.5 - 1e10) + 1.5)) - 3.14159)
                for fi21 := 0; fi21 < 10; fi21++ {
                    if false {
                        var v22 int8 = -38
                        v22 += v22
                        v23 := GenS3{F0: fi21, F1: 367, F2: fi21}
                        _ = v22
                        _ = v23
                    }
                    var v24 int64 = 0
                    if true {
                        v25 := []int{}
                        v26 := GenS3{F0: fi17, F1: 0, F2: fi17}
                        _ = v25
                        _ = v26
                    }
                    _ = fi21
                    _ = v24
                }
                _ = v20
            } else {
                p1 = p1
                for fi27 := 0; fi27 < 7; fi27++ {
                    v28 := GenS2{F0: 0, F1: (fi17 > (fi17 << ((65535) & 63))), F2: 42}
                    if true {
                        var v29 float64 = (3.14159 + (0.0 * (3.14159 - (1e10 + (0.0 - -2.25)))))
                        _ = v29
                    } else {
                        p1 = ((p1 + ("foo bar" + p1)) + p1)
                    }
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi17
        }
    }
    switch (p1 + p1) {
    case "a":
        var v30 int = 741
        _ = v30
    default:
        v31 := []int{9223372036854775807, 1, 32768, 692}
        _ = v31
    }
    var v32 bool = false
    if false {
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 float32 = 0.1
            v35 := []int{-9223372036854775808}
            v34 = v34
            _ = fi33
            _ = v34
            _ = v35
        }
        fmt.Println(v32, p1)
    }
    v36 := []int{127}
    _ = p0
    _ = p1
    _ = v32
    _ = v36
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 int8 = -128
    var v4 uint8 = 128
    var v37 uint16 = 255
    for fi38 := 0; fi38 < 2; fi38++ {
        if true {
            fmt.Println(v37, v4, v37)
        }
        v3 += (v3 - (-(v3)))
        _ = fi38
    }
    v37 = 363
    v3 += (v3 >> ((2) & 63))
    for fi39 := 0; fi39 < 10; fi39++ {
        v40 := []int{2147483647, -244, 32767, (fi39 + (fi39 % (((fi39 + fi39)) | 1)))}
        fmt.Println(v37, fi39)
        v4 = v4
        _ = fi39
        _ = v40
    }
    if (0.0 != -2.25) {
        v41 := GenS3{F0: 81, F1: 911, F2: 9223372036854775807}
        _ = v41
    } else {
        if (0 < -1) {
            if true {
                if true {
                    if false {
                        var v42 bool = (42 == 32767)
                        if false {
                            v43 := GenS1{F0: "a", F1: 4294967295, F2: (v37 * (^(v37)))}
                            if v42 {
                                if (v37 <= v37) {
                                    var v44 uint32 = 4294967295
                                    _ = v44
                                } else {
                                    v3 += v3
                                    var v45 bool = false
                                    if ((uint32(v37) <= 2147483648) && v42) {
                                        if false {
                                            if v45 {
                                                switch ("\n\t" + ("a" + ((("\n\t" + ("hello" + "a")) + "a") + ("" + "foo bar")))) {
                                                case "foo bar":
                                                    fmt.Println(v45, v42, v37)
                                                case "":
                                                    for fi46 := 0; fi46 < 20; fi46++ {
                                                        fmt.Println(v37, v4, v45)
                                                        _ = fi46
                                                    }
                                                case "\n\t":
                                                    delete(v2, -2147483648)
                                                }
                                                switch "a" {
                                                case "hello":
                                                    v42 = (0 == -1)
                                                case "\n\t":
                                                    delete(v1, "")
                                                case "a":
                                                    v45 = v42
                                                default:
                                                    for fi47 := 0; fi47 < 20; fi47++ {
                                                        switch fi47 {
                                                        case 42:
                                                            var v48 uint = 18446744073709551615
                                                            _ = v48
                                                        default:
                                                            if v42 {
                                                                v49 := GenS2{F0: v4, F1: (-2.25 == 3.14159), F2: -818}
                                                                for fi50 := 0; fi50 < 16; fi50++ {
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        switch "hello" {
                                                                        case "a":
                                                                            v51 := GenS3{F0: fi50, F1: 18446744073709551615, F2: fi47}
                                                                            _ = v51
                                                                        case "hello":
                                                                            var v52 bool = v45
                                                                            _ = v52
                                                                        }
                                                                    default:
                                                                        v4 = 2
                                                                    }
                                                                    v53 := []int{}
                                                                    _ = fi50
                                                                    _ = v53
                                                                }
                                                                _ = v49
                                                            } else {
                                                                fmt.Println(v3, v37)
                                                                v54 := []int{(fi47 / (((^(fi47))) | 1)), 127}
                                                                _ = v54
                                                            }
                                                        }
                                                        _ = fi47
                                                    }
                                                }
                                                v37 = 42
                                            } else {
                                                var v55 int = int(v4)
                                                _ = v55
                                            }
                                            for fi56 := 0; fi56 < 8; fi56++ {
                                                v3 += (v3 + (v3 - (v3 + (v3 << ((42) & 63)))))
                                                switch "" {
                                                case "\n\t":
                                                    v37 += v37
                                                default:
                                                    for fi57 := 0; fi57 < 11; fi57++ {
                                                        var v58 uint32 = 556
                                                        v37 = 2
                                                        if v45 {
                                                            _ = len(v2)
                                                        }
                                                        _ = fi57
                                                        _ = v58
                                                    }
                                                }
                                                _ = fi56
                                            }
                                        }
                                    } else {
                                        if v45 {
                                            switch -1 {
                                            case 32768:
                                                switch 32768 {
                                                case -747:
                                                    for fi59 := 0; fi59 < 5; fi59++ {
                                                        var v60 int = -1
                                                        v61 := GenS4{F0: (v4 - v4), F1: 2147483648, F2: ("\n\t" + ("\n\t" + (("foo bar" + ("hello" + (("a" + "a") + "\n\t"))) + "\n\t")))}
                                                        v4 = (^(v4))
                                                        _ = fi59
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                }
                                            case 260:
                                                v62 := []int{}
                                                _ = v62
                                            }
                                            for fi63 := 0; fi63 < 16; fi63++ {
                                                v64 := []int{}
                                                for fi65 := 0; fi65 < 19; fi65++ {
                                                    var v66 int16 = 277
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        if v45 {
                                                            fmt.Println(v4, fi65, v4)
                                                        } else {
                                                            if true {
                                                                var v67 float32 = 3.14159
                                                                switch "hello" {
                                                                case "a":
                                                                    v2[128] = ("hello" + "\n\t")
                                                                case "hello":
                                                                    v68 := GenS5{F0: (v4 - (v4 >> ((65535) & 63))), F1: (v37 % (((^(v37))) | 1)), F2: (v3 >= v3)}
                                                                    v4 -= (v4 / (((v4 % ((228) | 1))) | 1))
                                                                    v69 := GenS1{F0: ("hello" + "foo bar"), F1: 631, F2: v37}
                                                                    _ = v68
                                                                    _ = v69
                                                                case "foo bar":
                                                                    if len(v64) > 0 {
                                                                        v64[0] = (fi65 << ((2) & 63))
                                                                    }
                                                                }
                                                                _ = v67
                                                            }
                                                            if (!(("a" + "a") == "")) {
                                                                _ = len(v1)
                                                                var v70 uint16 = (v37 & 255)
                                                                v37 -= 42
                                                                _ = v70
                                                            }
                                                            v45 = (!true)
                                                        }
                                                    case "":
                                                        v71 := GenS4{F0: v4, F1: 2, F2: ((("hello" + "") + "hello") + ("foo bar" + "a"))}
                                                        _ = v71
                                                    default:
                                                        var v72 bool = false
                                                        _ = v72
                                                    }
                                                    if (!false) {
                                                        if true {
                                                            v37 -= v37
                                                            switch 0 {
                                                            case -2147483648:
                                                                switch (fi65 & 9223372036854775807) {
                                                                case 350:
                                                                    delete(v1, (("a" + ("hello" + "\n\t")) + ""))
                                                                case 32768:
                                                                    var v73 float32 = (-2.25 - 3.14159)
                                                                    _ = v73
                                                                default:
                                                                    v74 := GenS6{F0: -128, F1: 495, F2: 504}
                                                                    _ = v74
                                                                }
                                                            }
                                                        } else {
                                                            v45 = ((v37 / ((v37) | 1)) > 65535)
                                                        }
                                                        if (!v45) {
                                                            v75 := GenS7{F0: v4, F1: 42, F2: v66}
                                                            _ = v75
                                                        } else {
                                                            switch ("" + "foo bar") {
                                                            case "hello":
                                                                fmt.Println(v4, v66, v45)
                                                            case "foo bar":
                                                                if len(v64) >= 2 {
                                                                    v64 = v64[1:len(v64)]
                                                                }
                                                            case "\n\t":
                                                                v76 := GenS6{F0: genFn0(255, (((("\n\t" + "") + "") + (("a" + "a") + ("a" + "\n\t"))) + "foo bar")), F1: 0, F2: 42}
                                                                _ = v76
                                                            default:
                                                                v45 = (v42 || (("a" <= ((("hello" + "a") + "foo bar") + ("" + ("" + ("" + ""))))) && (false || ((1e10 * 3.14159) <= float64((0.0 * ((1.5 + 3.14159) + 0.1)))))))
                                                            }
                                                        }
                                                    }
                                                    _ = fi65
                                                    _ = v66
                                                }
                                                _ = fi63
                                                _ = v64
                                            }
                                            v42 = true
                                        } else {
                                            var v77 int32 = int32(v3)
                                            v37 = 2
                                            _ = v77
                                        }
                                        for fi78 := 0; fi78 < 18; fi78++ {
                                            for fi79 := 0; fi79 < 4; fi79++ {
                                                fmt.Println(fi78)
                                                if true {
                                                    v80 := GenS8{F0: 158, F1: 0, F2: (v37 >> ((774) & 63))}
                                                    v81 := []int{fi78, (fi78 + (fi78 & (fi79 | (fi79 << ((1) & 63)))))}
                                                    _ = v80
                                                    _ = v81
                                                } else {
                                                    var v82 uint16 = v37
                                                    switch (^(fi78)) {
                                                    case -23:
                                                        v4 -= v4
                                                    case -269:
                                                        if v42 {
                                                            v37 += v82
                                                        } else {
                                                            v3 += (v3 + v3)
                                                            v42 = false
                                                        }
                                                    default:
                                                        v3 = -120
                                                    }
                                                    var v83 int16 = 1
                                                    _ = v82
                                                    _ = v83
                                                }
                                                _ = fi79
                                            }
                                            if false {
                                                if (v45 && (819 < 0)) {
                                                    v45 = v42
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        switch 4294967295 {
                                                        case -9223372036854775808:
                                                            fmt.Println(fi78, v42)
                                                        }
                                                    }
                                                }
                                                var v84 uint8 = v4
                                                _ = v84
                                            }
                                            _ = fi78
                                        }
                                    }
                                    _ = v45
                                }
                            } else {
                                for fi85 := 0; fi85 < 1; fi85++ {
                                    switch fi85 {
                                    case 1:
                                        for fi86 := 0; fi86 < 3; fi86++ {
                                            v87 := []int{(fi86 << ((255) & 63))}
                                            var v88 int64 = 1
                                            _ = fi86
                                            _ = v87
                                            _ = v88
                                        }
                                    case 42:
                                        if false {
                                            v89 := GenS7{F0: v4, F1: genFn0(65535, "a"), F2: -1}
                                            _ = v89
                                        }
                                    case -206:
                                        v90 := GenS1{F0: "", F1: 42, F2: (v37 & v37)}
                                        _ = v90
                                    }
                                    _ = fi85
                                }
                                var v91 float64 = (float64(float32(float64(1e10))) * 1.5)
                                switch "foo bar" {
                                case "a":
                                    if true {
                                        v92 := GenS9{F0: 42, F1: 32767, F2: -223}
                                        switch (("" + "foo bar") + ("a" + "\n\t")) {
                                        case "\n\t":
                                            fmt.Println(v4, v4)
                                        case "hello":
                                            v93 := GenS4{F0: v4, F1: 1, F2: "foo bar"}
                                            _ = v93
                                        default:
                                            for fi94 := 0; fi94 < 15; fi94++ {
                                                if (v42 || v42) {
                                                    v95 := []int{fi94}
                                                    if (v37 >= v37) {
                                                        var v96 int = 32768
                                                        if len(v95) >= 2 {
                                                            v95 = v95[1:len(v95)]
                                                        }
                                                        _ = v96
                                                    } else {
                                                        if v42 {
                                                            v42 = (0 <= (v4 | (v4 >> ((912) & 63))))
                                                        } else {
                                                            var v97 float32 = (float32(v91) * 1.5)
                                                            v98 := GenS10{F0: (true || (!(2147483647 != 2)))}
                                                            var v99 uint = 18446744073709551615
                                                            _ = v97
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                    }
                                                    _ = v95
                                                } else {
                                                    if v42 {
                                                        v42 = v42
                                                    }
                                                    var v100 int64 = 42
                                                    _ = v100
                                                }
                                                var v101 int64 = int64(v37)
                                                fmt.Println(v4)
                                                _ = fi94
                                                _ = v101
                                            }
                                        }
                                        _ = v92
                                    } else {
                                        if (false || v42) {
                                            v91 -= 1e10
                                            var v102 float64 = v91
                                            _ = v102
                                        }
                                        switch -171 {
                                        case -369:
                                            for fi103 := 0; fi103 < 14; fi103++ {
                                                fmt.Println(v91, v91)
                                                if v42 {
                                                    v42 = v42
                                                    delete(v1, (("hello" + (((("" + (("hello" + ((("foo bar" + "a") + ("\n\t" + "")) + "")) + ("" + "a"))) + ("a" + "")) + "a") + "foo bar")) + "foo bar"))
                                                    v104 := []int{(fi103 - fi103)}
                                                    _ = v104
                                                } else {
                                                    switch (fi103 | (fi103 | fi103)) {
                                                    case 32767:
                                                        var v105 uint32 = 131
                                                        _ = v105
                                                    }
                                                    fmt.Println(v42, v4)
                                                }
                                                delete(v1, "a")
                                                _ = fi103
                                            }
                                        case 180:
                                            if v42 {
                                                v106 := GenS6{F0: genFn0(1, "\n\t"), F1: 42, F2: 18446744073709551615}
                                                v37 -= v37
                                                v37 = (^(v37))
                                                _ = v106
                                            } else {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v4 -= (v4 & (v4 * v4))
                                                case "hello":
                                                    for fi107 := 0; fi107 < 9; fi107++ {
                                                        if v42 {
                                                            if (891 > -1) {
                                                                if (18446744073709551615 == 34) {
                                                                    fmt.Println(v4, v37)
                                                                    v2[(fi107 % (((fi107 ^ 32767)) | 1))] = "a"
                                                                    var v108 float32 = 1e10
                                                                    _ = v108
                                                                }
                                                                if (0 != -688) {
                                                                    v109 := GenS1{F0: "a", F1: 65535, F2: v37}
                                                                    v91 += -2.25
                                                                    _ = v109
                                                                }
                                                                if (v42 || (!((!(!v42)) && ((v3 >> ((221) & 63)) < v3)))) {
                                                                    v110 := GenS1{F0: (("hello" + ((("a" + "a") + ("a" + "")) + ("" + ""))) + "\n\t"), F1: 225, F2: (v37 & 32767)}
                                                                    if v42 {
                                                                        if v42 {
                                                                            v111 := []int{(^(fi107)), 163}
                                                                            fmt.Println(v3)
                                                                            v91 = (float64((float32(3.14159) + 3.14159)) + v91)
                                                                            _ = v111
                                                                        }
                                                                    }
                                                                    _ = v110
                                                                }
                                                            }
                                                        } else {
                                                            switch "" {
                                                            case "\n\t":
                                                                var v112 int16 = 1000
                                                                _ = v112
                                                            case "foo bar":
                                                                v113 := GenS7{F0: 44, F1: genFn0(483, ("hello" + "hello")), F2: 1000}
                                                                _ = v113
                                                            default:
                                                                var v114 float64 = v91
                                                                _ = v114
                                                            }
                                                            fmt.Println(fi107)
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v115 int64 = -274
                                                                _ = v115
                                                            default:
                                                                v4 = v4
                                                            }
                                                        }
                                                        v2[(fi107 * fi107)] = "\n\t"
                                                        _ = fi107
                                                    }
                                                default:
                                                    v116 := GenS11{F0: 1, F1: 18446744073709551615}
                                                    _ = v116
                                                }
                                                var v117 int32 = -2147483648
                                                var v118 int16 = 42
                                                _ = v117
                                                _ = v118
                                            }
                                        case 42:
                                            delete(v2, 42)
                                        }
                                    }
                                case "":
                                    switch -89 {
                                    case 1:
                                        var v119 int8 = (^(v3))
                                        _ = v119
                                    }
                                }
                                _ = v91
                            }
                            fmt.Println(v4, v3)
                            _ = v43
                        }
                        for fi120 := 0; fi120 < 19; fi120++ {
                            v37 += v37
                            if ("a" != "a") {
                                switch 32767 {
                                case 4294967295:
                                    if (false && false) {
                                        var v121 int32 = 2147483647
                                        if (306 > int64(v37)) {
                                            var v122 int64 = 811
                                            _ = v122
                                        } else {
                                            v42 = (0 != v37)
                                            var v123 uint32 = 65535
                                            v2[(fi120 / ((1) | 1))] = ("a" + (("" + "a") + "hello"))
                                            _ = v123
                                        }
                                        _ = v121
                                    } else {
                                        fmt.Println(v37, fi120, v37)
                                    }
                                }
                                if (255 == (v4 % (((v4 | v4)) | 1))) {
                                    switch fi120 {
                                    case 32767:
                                        for fi124 := 0; fi124 < 2; fi124++ {
                                            v125 := []int{}
                                            var v126 uint32 = 65535
                                            v127 := GenS11{F0: int32(v4), F1: 0}
                                            _ = fi124
                                            _ = v125
                                            _ = v126
                                            _ = v127
                                        }
                                    case -1:
                                        v37 -= (v37 * (v37 + v37))
                                    case 1:
                                        v128 := GenS12{F0: 4294967295}
                                        _ = v128
                                    default:
                                        for fi129 := 0; fi129 < 3; fi129++ {
                                            if v42 {
                                                var v130 int64 = 0
                                                v131 := []int{(fi120 / ((fi129) | 1)), fi129}
                                                _ = v130
                                                _ = v131
                                            }
                                            if v42 {
                                                var v132 string = "\n\t"
                                                v37 = 1
                                                v3 += (^(v3))
                                                _ = v132
                                            } else {
                                                if v42 {
                                                    var v133 float32 = (3.14159 * 1e10)
                                                    if v42 {
                                                        var v134 string = "hello"
                                                        v135 := []int{(fi129 % ((-2147483648) | 1))}
                                                        v37 = v37
                                                        v134 = v134
                                                        _ = v134
                                                        _ = v135
                                                    } else {
                                                        v136 := []int{}
                                                        _ = v136
                                                    }
                                                    v137 := GenS3{F0: -283, F1: 0, F2: fi120}
                                                    _ = v133
                                                    _ = v137
                                                } else {
                                                    v138 := GenS4{F0: (v4 << ((836) & 63)), F1: 1, F2: "foo bar"}
                                                    _ = v138
                                                }
                                            }
                                            var v139 uint = 0
                                            _ = fi129
                                            _ = v139
                                        }
                                    }
                                } else {
                                    v140 := []int{(fi120 >> ((42) & 63)), fi120, (fi120 * 32768), fi120}
                                    var v141 uint = 1
                                    _ = v140
                                    _ = v141
                                }
                            }
                            _ = fi120
                        }
                        _ = v42
                    } else {
                        var v142 uint8 = (v4 % (((v4 ^ (^(v4)))) | 1))
                        v143 := []int{int(v37), 32767, -50}
                        var v144 uint16 = (v37 + (v37 | v37))
                        _ = v142
                        _ = v143
                        _ = v144
                    }
                    if (42 == -9223372036854775808) {
                        fmt.Println(v37, v3)
                    }
                }
                var v145 int16 = 1000
                fmt.Println(v4, v4, v4)
                _ = v145
            } else {
                v3 += 82
                v146 := GenS6{F0: 127, F1: 1, F2: 965}
                _ = v146
            }
            v37 = 903
            v147 := []int{}
            _ = v147
        } else {
            v4 = v4
        }
        for fi148 := 0; fi148 < 11; fi148++ {
            for fi149 := 0; fi149 < 4; fi149++ {
                switch ("" + "foo bar") {
                case "foo bar":
                    if true {
                        v150 := []int{fi149, 4294967295}
                        switch fi149 {
                        case 42:
                            v3 = (v3 / ((v3) | 1))
                        }
                        if false {
                            v4 -= 11
                            var v151 int32 = 2147483647
                            _ = v151
                        }
                        _ = v150
                    }
                }
                v3 -= genFn0(65535, "\n\t")
                _ = fi149
            }
            v152 := []int{fi148, (fi148 % (((^(fi148))) | 1)), fi148}
            _ = fi148
            _ = v152
        }
    }
    v2[42] = ("hello" + "hello")
    for fi153 := 0; fi153 < 2; fi153++ {
        v4 = 0
        fmt.Println(v37, v4, fi153)
        if true {
            var v154 int16 = -69
            var v155 bool = true
            fmt.Println(v37, v155, v155)
            _ = v154
            _ = v155
        } else {
            v3 = (v3 ^ (v3 + (v3 * v3)))
        }
        _ = fi153
    }
    for fi156 := 0; fi156 < 7; fi156++ {
        var v157 uint16 = (v37 >> ((4294967295) & 63))
        _ = fi156
        _ = v157
    }
    for fi158 := 0; fi158 < 3; fi158++ {
        var v159 int32 = -535
        var v160 uint = 4294967295
        var v161 int32 = -683
        _ = fi158
        _ = v159
        _ = v160
        _ = v161
    }
    v162 := []int{}
    switch -2147483648 {
    case -1:
        v37 += 32767
    case -723:
        v2[-435] = (((("a" + "\n\t") + "\n\t") + "hello") + "foo bar")
    case -738:
        var v163 float32 = -2.25
        _ = v163
    default:
        if (true || false) {
            v3 = 1
        } else {
            if (true || true) {
                if len(v162) >= 2 {
                    v162 = v162[1:len(v162)]
                }
            } else {
                v4 = v4
                for fi164 := 0; fi164 < 17; fi164++ {
                    v162 = append(v162, fi164)
                    _ = fi164
                }
            }
            if (-554 > 536) {
                fmt.Println(v4)
            }
        }
    }
    if true {
        for fi165 := 0; fi165 < 18; fi165++ {
            if false {
                var v166 uint64 = 42
                var v167 bool = true
                _ = v166
                _ = v167
            }
            for fi168 := 0; fi168 < 10; fi168++ {
                _ = len(v1)
                if (("\n\t" + ("" + "\n\t")) != "\n\t") {
                    fmt.Println(v4, v37)
                }
                v4 += v4
                v37 = v37
                _ = len(v162)
                v169 := GenS8{F0: v4, F1: 2, F2: v37}
                v170 := GenS6{F0: v3, F1: 42, F2: 26}
                v4 = 66
                v162 = append(v162, 128)
                _ = fi168
                _ = v169
                _ = v170
            }
            var v171 int64 = 9223372036854775807
            _ = fi165
            _ = v171
        }
        fmt.Println(v3, v4, v4)
        for fi172 := 0; fi172 < 16; fi172++ {
            v173 := v1[("" + "\n\t")]
            _ = fi172
            _ = v173
        }
    }
    v37 = (v37 % (((^(v37))) | 1))
    v174 := GenS10{F0: false}
    v3 = (v3 & genFn0(117, ("\n\t" + (("foo bar" + (("\n\t" + "hello") + ("" + ""))) + "hello"))))
    var v175 uint64 = 595
    fmt.Println(v4, v37)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v162:", len(v162), v162)
    fmt.Println("v175:", v175)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v162
    _ = v174
    _ = v175
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
