package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 bool }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int16; F1 int16; F2 int32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 int16 }
type GenS10 struct { F0 uint32; F1 string; F2 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16; F1 uint8; F2 int8 }
type GenS14 struct { F0 uint; F1 string }
type GenS15 struct { F0 uint16; F1 uint; F2 int }
func genFn0(p0 uint8) int {
    fmt.Println(p0)
    if true {
        v7 := GenS1{F0: -9223372036854775808, F1: -32768, F2: false}
        v8 := []int{}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            if true {
                fmt.Println(p0)
                v10 := GenS2{F0: (((212 >= (^(p0))) || (!true)) && true), F1: 42}
                _ = v10
            } else {
                v11 := []int{}
                if false {
                    var v12 uint32 = 507
                    _ = v12
                } else {
                    p0 -= (p0 * p0)
                    v13 := GenS1{F0: -9223372036854775808, F1: 700, F2: ((p0 / ((p0) | 1)) <= p0)}
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v11 = append(v11, -2147483648)
                        _ = fi14
                    }
                    _ = v13
                }
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = v11
            }
            _ = fi9
        }
        var v15 float64 = 1e10
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        if ((42 != 4) || false) {
            if true {
                p0 = 86
            } else {
                var v17 uint64 = 4294967295
                v18 := []int{fi16, fi16}
                _ = v17
                _ = v18
            }
            fmt.Println(fi16)
            v19 := []int{fi16, fi16}
            _ = v19
        }
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS3{F0: 171, F1: 776}
            v22 := GenS3{F0: 389, F1: 4294967295}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi16
    }
    _ = p0
    return 32768
}
func genFn1() int16 {
    var v23 float64 = (((0.0 - 1.5) / -2.25) - 1e10)
    var v24 float64 = ((1.5 / v23) * (v23 + v23))
    v25 := GenS3{F0: 590, F1: 0}
    v26 := []int{-1, 127, -1, -9223372036854775808}
    v24 += (((1.5 + (3.14159 * 1e10)) / v24) - v23)
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{9223372036854775807, 1, 1}
    var v3 int64 = -1
    var v4 float32 = (1.5 / 0.1)
    var v5 string = ("hello" + "")
    var v6 uint16 = 583
    if false {
        fmt.Println(v4)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 = ((v5 + ("" + ("hello" + v5))) + ("foo bar" + ""))
    } else {
        v27 := GenS2{F0: (1 != 42), F1: 307}
        _ = v27
    }
    var v28 uint64 = 453
    switch 0 {
    case -9223372036854775808:
        var v29 int8 = 0
        _ = v29
    case -146:
        if true {
            if (2147483647 == -1) {
                v28 += v28
            } else {
                v3 = -980
                if ((genFn0(27) != 245) || false) {
                    v30 := GenS4{F0: 1000, F1: -609, F2: -2147483648}
                    v31 := GenS2{F0: (-128 > -128), F1: v28}
                    _ = v30
                    _ = v31
                }
            }
            fmt.Println(v6, v3)
        }
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS5{F0: 911, F1: uint32(v6)}
            _ = fi32
            _ = v33
        }
    case "":
        switch ((v5 + v5) + "") {
        case "\n\t":
            if (569 <= 128) {
                var v34 int64 = v3
                if false {
                    if (v5 == ((v5 + "foo bar") + v5)) {
                        var v35 uint8 = 128
                        if false {
                            var v36 int16 = 654
                            var v37 float32 = v4
                            v1[v5] = 680
                            _ = v36
                            _ = v37
                        }
                        _ = v35
                    }
                    v2 = append(v2, 32767)
                }
                _ = v34
            }
        }
    }
    if (!(42 != 42)) {
        var v38 float64 = -2.25
        _ = len(v1)
        _ = v38
    } else {
        v5 = "foo bar"
        var v39 int64 = v3
        _ = v39
    }
    switch ("" + "foo bar") {
    case "foo bar":
        switch v5 {
        case "":
            if false {
                var v40 int8 = 127
                if true {
                    var v41 uint32 = 2147483647
                    _ = v41
                } else {
                    fmt.Println(v3, v40)
                    var v42 int64 = (v3 / (((v3 + v3)) | 1))
                    var v43 uint = 0
                    _ = v42
                    _ = v43
                }
                _ = len(v2)
                _ = v40
            }
        default:
            var v44 int32 = 1
            _ = v44
        }
    default:
        switch "foo bar" {
        case "a":
            switch 32767 {
            case -9223372036854775808:
                for fi45 := 0; fi45 < 18; fi45++ {
                    var v46 int = (fi45 % ((810) | 1))
                    _ = fi45
                    _ = v46
                }
            default:
                v3 = v3
            }
        case "":
            if false {
                v4 -= (v4 / ((v4 / v4) / v4))
            } else {
                v4 = v4
                var v47 uint16 = v6
                fmt.Println(v3)
                _ = v47
            }
        default:
            v5 = "\n\t"
        }
    }
    if ((42 != 255) && (0 == 410)) {
        v6 = 1
        v4 -= 1.5
    } else {
        if true {
            if (v3 <= v3) {
                var v48 uint64 = 0
                delete(v1, v5)
                if false {
                    if true {
                        v5 = ("" + "foo bar")
                        for fi49 := 0; fi49 < 3; fi49++ {
                            if false {
                                switch v5 {
                                case "foo bar":
                                    var v50 uint = 18446744073709551615
                                    _ = v50
                                }
                                switch fi49 {
                                case 32768:
                                    v4 += v4
                                case -2147483648:
                                    if (42 > 858) {
                                        if (false && false) {
                                            fmt.Println(v4)
                                        }
                                        var v51 uint16 = (^(v6))
                                        v51 = v6
                                        _ = v51
                                    }
                                }
                                switch (v5 + ("\n\t" + ((("foo bar" + v5) + ("foo bar" + v5)) + (("" + (((v5 + v5) + "") + v5)) + ((v5 + v5) + (v5 + ((((("" + (v5 + v5)) + ((v5 + ("hello" + (v5 + "foo bar"))) + (v5 + (("" + (v5 + ("foo bar" + "\n\t"))) + v5)))) + v5) + (v5 + v5)) + v5))))))) {
                                case "a":
                                    fmt.Println(v6)
                                case "":
                                    var v52 string = ((v5 + v5) + "foo bar")
                                    _ = v52
                                case "hello":
                                    for fi53 := 0; fi53 < 2; fi53++ {
                                        var v54 bool = (1 <= -1)
                                        var v55 uint32 = 65535
                                        v56 := GenS5{F0: -2147483648, F1: (v55 / (((v55 << ((65535) & 63))) | 1))}
                                        fmt.Println(v3, v3, v48)
                                        _ = fi53
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                }
                            } else {
                                if true {
                                    v1[("\n\t" + "hello")] = 2147483648
                                    var v57 uint32 = 623
                                    v58 := v1["\n\t"]
                                    _ = v57
                                    _ = v58
                                } else {
                                    v6 -= 32767
                                    var v59 uint16 = (v6 ^ (v6 >> ((819) & 63)))
                                    if false {
                                        for fi60 := 0; fi60 < 8; fi60++ {
                                            switch fi60 {
                                            case 32767:
                                                if false {
                                                    switch (fi49 & fi60) {
                                                    case -9223372036854775808:
                                                        var v61 string = v5
                                                        _ = v61
                                                    default:
                                                        switch fi60 {
                                                        case 32767:
                                                            var v62 float32 = v4
                                                            _ = v62
                                                        case -315:
                                                            v5 = v5
                                                        case 117:
                                                            var v63 uint32 = 2147483648
                                                            _ = v63
                                                        default:
                                                            v64 := GenS1{F0: 9223372036854775807, F1: 42, F2: (false || ((0 > 970) && false))}
                                                            _ = v64
                                                        }
                                                        if (false || (42 > 1)) {
                                                            if false {
                                                                switch ((((v5 + (v5 + "a")) + "a") + "a") + v5) {
                                                                case "hello":
                                                                    var v65 string = "\n\t"
                                                                    _ = v65
                                                                }
                                                                v6 = (v6 >> ((255) & 63))
                                                                if (!true) {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 42
                                                                    }
                                                                    if true {
                                                                        var v66 float64 = (0.0 + 3.14159)
                                                                        v5 = "hello"
                                                                        if false {
                                                                            fmt.Println(v66, v28)
                                                                            fmt.Println(fi60)
                                                                        }
                                                                        _ = v66
                                                                    } else {
                                                                        if false {
                                                                            var v67 uint32 = 0
                                                                            var v68 int16 = -606
                                                                            _ = v67
                                                                            _ = v68
                                                                        } else {
                                                                            v69 := GenS6{F0: (!true)}
                                                                            switch (fi60 % ((fi60) | 1)) {
                                                                            case -281:
                                                                                var v70 bool = false
                                                                                _ = v70
                                                                            case 4294967295:
                                                                                delete(v1, v5)
                                                                            case 32767:
                                                                                v3 = (v3 - 769)
                                                                            }
                                                                            _ = v69
                                                                        }
                                                                        v71 := GenS7{F0: 127, F1: -596}
                                                                        _ = v71
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = fi60
                                                                    }
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v3 -= (v3 % ((0) | 1))
                                                            }
                                                        }
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v72 := GenS1{F0: 42, F1: 32767, F2: (!true)}
                                                    _ = v72
                                                } else {
                                                    v73 := GenS1{F0: (^(v3)), F1: 42, F2: true}
                                                    _ = v73
                                                }
                                            }
                                            fmt.Println(v48, v3, v5)
                                            _ = fi60
                                        }
                                        _ = len(v2)
                                    }
                                    _ = v59
                                }
                                if (-128 == 1) {
                                    v28 -= 42
                                    for fi74 := 0; fi74 < 12; fi74++ {
                                        var v75 string = (v5 + "\n\t")
                                        _ = fi74
                                        _ = v75
                                    }
                                    v28 += 734
                                } else {
                                    var v76 uint32 = 4294967295
                                    _ = v76
                                }
                                if (127 >= 99) {
                                    for fi77 := 0; fi77 < 12; fi77++ {
                                        var v78 int32 = 1
                                        fmt.Println(v5, v28)
                                        if (fi49 >= fi49) {
                                            v4 += float32((1e10 - 1e10))
                                            switch v5 {
                                            case "":
                                                var v79 uint64 = v48
                                                _ = v79
                                            case "foo bar":
                                                v48 = (v48 + v28)
                                            case "hello":
                                                v6 += 1
                                            default:
                                                var v80 float64 = float64(v4)
                                                _ = v80
                                            }
                                            v4 = float32(3.14159)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v81 int16 = -224
                                            _ = v81
                                        }
                                        _ = fi77
                                        _ = v78
                                    }
                                    var v82 uint = 255
                                    _ = v82
                                } else {
                                    if (589 <= 42) {
                                        for fi83 := 0; fi83 < 20; fi83++ {
                                            v4 -= 3.14159
                                            if (-1 <= 352) {
                                                if (true || false) {
                                                    fmt.Println(v48)
                                                    var v84 int64 = v3
                                                    _ = v84
                                                } else {
                                                    v85 := GenS8{F0: int16(fi83)}
                                                    fmt.Println(fi49, v3)
                                                    _ = v85
                                                }
                                                fmt.Println(v4, fi83, v5)
                                            }
                                            if (!(3.14159 >= (float32(float64(float32((0.1 / 0.1)))) - v4))) {
                                                if ((v28 | 0) >= (v28 & 656)) {
                                                    var v86 bool = false
                                                    delete(v1, "\n\t")
                                                    v6 = (v6 >> ((10) & 63))
                                                    if (-9223372036854775808 == fi49) {
                                                        var v87 int32 = 1
                                                        var v88 uint = uint(v48)
                                                        if len(v2) > 0 {
                                                            v2[0] = 4294967295
                                                        }
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        switch v5 {
                                                        case "hello":
                                                            v4 = 3.14159
                                                        case "\n\t":
                                                            var v89 string = (v5 + "a")
                                                            _ = v89
                                                        default:
                                                            switch fi49 {
                                                            case 4294967295:
                                                                var v90 int = 2147483647
                                                                _ = v90
                                                            default:
                                                                var v91 uint16 = (^(v6))
                                                                _ = v91
                                                            }
                                                        }
                                                        if (!true) {
                                                            var v92 string = v5
                                                            v93 := GenS9{F0: v86, F1: 0}
                                                            _ = len(v2)
                                                            if v86 {
                                                                if v86 {
                                                                    var v94 int64 = int64(v6)
                                                                    if (104 >= -128) {
                                                                        v48 = 4294967295
                                                                        var v95 int32 = -2147483648
                                                                        switch v92 {
                                                                        case "\n\t":
                                                                            v96 := GenS10{F0: 922, F1: v92, F2: 72}
                                                                            _ = v96
                                                                        case "foo bar":
                                                                            if (v6 > (v6 + (v6 % ((v6) | 1)))) {
                                                                                fmt.Println(v86, v94)
                                                                                if (true || (0 != (v94 & v3))) {
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        v2 = append(v2, (fi49 << ((4294967295) & 63)))
                                                                                    case "a":
                                                                                        if (!(true && v86)) {
                                                                                            fmt.Println(v95, v92, v92)
                                                                                            _ = len(v1)
                                                                                            var v97 int = (fi49 & genFn0(27))
                                                                                            _ = v97
                                                                                        }
                                                                                    }
                                                                                    var v98 uint64 = 2
                                                                                    v6 += (v6 + 42)
                                                                                    _ = v98
                                                                                } else {
                                                                                    if ((v6 & (v6 - (^(v6)))) != (^(v6))) {
                                                                                        var v99 uint16 = v6
                                                                                        var v100 int8 = -121
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    } else {
                                                                                        var v101 string = v92
                                                                                        switch (fi83 - (fi49 + (fi83 | (fi83 * -2147483648)))) {
                                                                                        case -644:
                                                                                            v48 = 137
                                                                                        case 128:
                                                                                            delete(v1, (v92 + "a"))
                                                                                        case -224:
                                                                                            v3 = 0
                                                                                        default:
                                                                                            v102 := GenS1{F0: (v94 % ((v94) | 1)), F1: genFn1(), F2: false}
                                                                                            v103 := GenS11{F0: 1}
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                        }
                                                                                        _ = v101
                                                                                    }
                                                                                    var v104 int = (fi83 ^ genFn0(127))
                                                                                    if (0 >= genFn1()) {
                                                                                        switch v104 {
                                                                                        case -2147483648:
                                                                                            if v86 {
                                                                                                v104 = (fi83 << ((803) & 63))
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                var v105 int = (v104 << ((711) & 63))
                                                                                                _ = v105
                                                                                            }
                                                                                        case 348:
                                                                                            switch (v5 + "") {
                                                                                            case "\n\t":
                                                                                                var v106 int64 = -255
                                                                                                _ = v106
                                                                                            case "":
                                                                                                var v107 uint32 = 65535
                                                                                                _ = v107
                                                                                            case "foo bar":
                                                                                                var v108 uint32 = 2
                                                                                                _ = v108
                                                                                            default:
                                                                                                v109 := GenS10{F0: 65535, F1: (v5 + ((v92 + "\n\t") + "\n\t")), F2: 1}
                                                                                                _ = v109
                                                                                            }
                                                                                        case 2147483647:
                                                                                            if (float64(float32(3.14159)) < 0.0) {
                                                                                                v6 = v6
                                                                                                switch (v104 | (fi83 >> ((18446744073709551615) & 63))) {
                                                                                                case -2147483648:
                                                                                                    var v110 bool = v86
                                                                                                    _ = v110
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v6 -= v6
                                                                                    _ = v104
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v92, v28)
                                                                                if (v86 || (v48 == v28)) {
                                                                                    v6 += 65535
                                                                                    var v111 int8 = 100
                                                                                    if (3.14159 == -2.25) {
                                                                                        v4 = 0.1
                                                                                    } else {
                                                                                        if v86 {
                                                                                            v3 += v94
                                                                                        } else {
                                                                                            _ = len(v1)
                                                                                            if (v4 == v4) {
                                                                                                v112 := v1[(v92 + (("" + (("hello" + (("foo bar" + (v5 + "")) + (((v92 + v5) + v92) + v5))) + v92)) + ((v5 + (("" + v92) + v92)) + ("hello" + ("hello" + (v5 + v5))))))]
                                                                                                if (!((!(true || false)) && v86)) {
                                                                                                    v113 := GenS7{F0: 1, F1: 2147483647}
                                                                                                    var v114 uint32 = 42
                                                                                                    _ = v113
                                                                                                    _ = v114
                                                                                                }
                                                                                                fmt.Println(v4, v94)
                                                                                                var v115 uint32 = 2147483648
                                                                                                _ = v112
                                                                                                _ = v115
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v111
                                                                                } else {
                                                                                    switch -806 {
                                                                                    case 965:
                                                                                        if true {
                                                                                            if (-2147483648 == -1) {
                                                                                                switch 2147483647 {
                                                                                                case -918:
                                                                                                    v48 += 4294967295
                                                                                                    v95 += v95
                                                                                                }
                                                                                                var v116 bool = true
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    v92 = v92
                                                                                                case "\n\t":
                                                                                                    switch v92 {
                                                                                                    case "a":
                                                                                                        v4 = v4
                                                                                                    }
                                                                                                default:
                                                                                                    switch (fi49 ^ -1) {
                                                                                                    case 127:
                                                                                                        var v117 float32 = 1.5
                                                                                                        _ = v117
                                                                                                    case 128:
                                                                                                        v2 = append(v2, 2147483647)
                                                                                                    default:
                                                                                                        var v118 int64 = v94
                                                                                                        _ = v118
                                                                                                    }
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    v94 = (-(v94))
                                                                                                case "hello":
                                                                                                    if v116 {
                                                                                                        if v86 {
                                                                                                            switch "" {
                                                                                                            case "foo bar":
                                                                                                                v94 += (v94 | 42)
                                                                                                            }
                                                                                                            v119 := GenS12{F0: 430}
                                                                                                            _ = v119
                                                                                                        } else {
                                                                                                            v120 := GenS13{F0: 1, F1: 1, F2: 127}
                                                                                                            v121 := GenS14{F0: uint(v28), F1: ""}
                                                                                                            v6 += v6
                                                                                                            var v122 int32 = 1
                                                                                                            if v86 {
                                                                                                                switch (^(fi49)) {
                                                                                                                case -582:
                                                                                                                    fmt.Println(fi83)
                                                                                                                }
                                                                                                            }
                                                                                                            v48 = v48
                                                                                                            switch (v5 + v5) {
                                                                                                            case "":
                                                                                                                fmt.Println(v122, v48, v5)
                                                                                                            case "a":
                                                                                                                if (v116 && (1 <= 1)) {
                                                                                                                    v1[""] = 917
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v122 -= v122
                                                                                                            default:
                                                                                                                var v123 int = (fi49 - fi83)
                                                                                                                _ = v123
                                                                                                            }
                                                                                                            _ = len(v2)
                                                                                                            _ = v120
                                                                                                            _ = v121
                                                                                                            _ = v122
                                                                                                        }
                                                                                                    } else {
                                                                                                        _ = len(v2)
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = fi49
                                                                                                            }
                                                                                                        }
                                                                                                        switch v92 {
                                                                                                        case "a":
                                                                                                            v116 = (100 > -119)
                                                                                                        case "\n\t":
                                                                                                            v2 = append(v2, genFn0(127))
                                                                                                        case "foo bar":
                                                                                                            if ((v48 >> ((1) & 63)) >= (v28 * v48)) {
                                                                                                                var v124 int8 = 127
                                                                                                                _ = v124
                                                                                                            }
                                                                                                        default:
                                                                                                            var v125 uint = uint(v28)
                                                                                                            _ = v125
                                                                                                        }
                                                                                                    }
                                                                                                case "":
                                                                                                    v5 = ""
                                                                                                default:
                                                                                                    var v126 string = v5
                                                                                                    _ = v126
                                                                                                }
                                                                                                switch genFn0(198) {
                                                                                                case 127:
                                                                                                    fmt.Println(v94, fi49)
                                                                                                case 458:
                                                                                                    var v127 int64 = (v3 - 42)
                                                                                                    _ = v127
                                                                                                case 4294967295:
                                                                                                    var v128 uint32 = 2
                                                                                                    _ = v128
                                                                                                }
                                                                                                _ = v116
                                                                                            }
                                                                                            fmt.Println(v92)
                                                                                        } else {
                                                                                            v1["a"] = 2147483648
                                                                                        }
                                                                                    case -800:
                                                                                        v48 = (v48 << ((524) & 63))
                                                                                    case 4294967295:
                                                                                        v129 := GenS8{F0: 0}
                                                                                        _ = v129
                                                                                    }
                                                                                    var v130 string = ""
                                                                                    var v131 uint64 = v28
                                                                                    _ = v130
                                                                                    _ = v131
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            v132 := GenS11{F0: 1}
                                                                            _ = v132
                                                                        default:
                                                                            v133 := GenS15{F0: (v6 << ((2) & 63)), F1: 419, F2: (fi49 + (^(fi83)))}
                                                                            _ = v133
                                                                        }
                                                                        _ = v95
                                                                    }
                                                                    _ = v94
                                                                } else {
                                                                    v6 -= (v6 / (((v6 + v6)) | 1))
                                                                }
                                                                v2 = append(v2, (fi83 ^ 9223372036854775807))
                                                            } else {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                if v86 {
                                                                    switch fi49 {
                                                                    case 4294967295:
                                                                        fmt.Println(v4)
                                                                    case 127:
                                                                        fmt.Println(fi49, fi49)
                                                                    case 0:
                                                                        v4 = 0.1
                                                                    default:
                                                                        fmt.Println(v28, v3, v28)
                                                                    }
                                                                    if false {
                                                                        var v134 uint64 = (v48 >> ((571) & 63))
                                                                        _ = v134
                                                                    }
                                                                    v92 = v5
                                                                } else {
                                                                    v6 += v6
                                                                }
                                                                fmt.Println(v4)
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v135 uint16 = (v6 << ((1) & 63))
                                                                    _ = v135
                                                                }
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = v92
                                                            _ = v93
                                                        }
                                                    }
                                                    _ = v86
                                                } else {
                                                    var v136 int32 = -1
                                                    if true {
                                                        v28 += (v48 + (v48 + 42))
                                                        fmt.Println(fi83)
                                                        v1[((("a" + ((((v5 + "") + v5) + v5) + "")) + (v5 + v5)) + (v5 + "a"))] = 4294967295
                                                    } else {
                                                        fmt.Println(fi49)
                                                    }
                                                    if false {
                                                        v5 = v5
                                                        var v137 uint = 0
                                                        fmt.Println(v28, v4)
                                                        _ = v137
                                                    } else {
                                                        v138 := GenS7{F0: 127, F1: (-(v136))}
                                                        _ = v138
                                                    }
                                                    _ = v136
                                                }
                                            }
                                            _ = fi83
                                        }
                                        var v139 bool = false
                                        switch "foo bar" {
                                        case "hello":
                                            v4 = v4
                                        case "foo bar":
                                            var v140 uint64 = (v28 / ((v28) | 1))
                                            _ = v140
                                        case "a":
                                            var v141 uint = 0
                                            _ = v141
                                        }
                                        _ = v139
                                    } else {
                                        switch 0 {
                                        case 908:
                                            var v142 float32 = v4
                                            _ = v142
                                        case 4294967295:
                                            var v143 int32 = -21
                                            _ = v143
                                        case 204:
                                            var v144 uint32 = 591
                                            _ = v144
                                        default:
                                            v145 := GenS5{F0: 991, F1: 2}
                                            _ = v145
                                        }
                                    }
                                    v146 := GenS11{F0: 0}
                                    _ = v146
                                }
                            }
                            if false {
                                switch 42 {
                                case 32768:
                                    fmt.Println(v5, fi49, v5)
                                }
                                fmt.Println(v5, v5)
                            }
                            _ = fi49
                        }
                    }
                    switch 2147483647 {
                    case 32768:
                        if len(v2) > 0 {
                            v2[0] = 818
                        }
                    }
                    v5 = v5
                } else {
                    if true {
                        var v147 bool = false
                        if (v3 < v3) {
                            var v148 bool = v147
                            _ = v148
                        }
                        v149 := GenS5{F0: 881, F1: 4294967295}
                        _ = v147
                        _ = v149
                    }
                }
                _ = v48
            }
            var v150 float64 = 3.14159
            _ = v150
        } else {
            for fi151 := 0; fi151 < 5; fi151++ {
                v5 = v5
                var v152 uint32 = 65535
                _ = fi151
                _ = v152
            }
        }
    }
    v6 = 1
    var v153 string = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v153:", v153)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v153
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
