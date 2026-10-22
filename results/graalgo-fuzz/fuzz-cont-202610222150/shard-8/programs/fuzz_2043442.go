package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 int; F1 int; F2 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64; F1 uint8 }
type GenS13 struct { F0 int; F1 bool }
func genFn0(p0 int16, p1 uint) string {
    for fi5 := 0; fi5 < 2; fi5++ {
        p0 += (p0 - p0)
        v6 := GenS1{F0: p1, F1: ""}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: (p1 * p1), F1: ("foo bar" + ("\n\t" + "a"))}
    fmt.Println(p0, p1)
    if true {
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                if (((!(fi8 != (fi9 / ((fi9) | 1)))) || (0.1 == 3.14159)) || (572 > 135)) {
                    v10 := GenS1{F0: 497, F1: ("" + ("" + "foo bar"))}
                    _ = v10
                } else {
                    if ((-1 == 827) && true) {
                        fmt.Println(fi8)
                        var v11 bool = (true && false)
                        _ = v11
                    }
                }
                _ = fi9
            }
            v12 := []int{(fi8 | fi8), 144}
            _ = fi8
            _ = v12
        }
        if (false || false) {
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 int64 = -9223372036854775808
                _ = fi13
                _ = v14
            }
            v15 := []int{1, 902, 4294967295, 210}
            v16 := GenS1{F0: p1, F1: "a"}
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: p1, F1: "hello"}
            p0 -= 0
            var v18 uint32 = 18
            _ = v17
            _ = v18
        }
    } else {
        if ((p1 % (((p1 & p1)) | 1)) != (p1 ^ p1)) {
            var v19 int = 398
            fmt.Println(p1)
            _ = v19
        }
        if false {
            var v20 bool = true
            var v21 uint64 = 1
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 int64 = -9223372036854775808
                _ = fi22
                _ = v23
            }
        }
        if false {
            switch 4294967295 {
            case 485:
                var v24 uint32 = 535
                _ = v24
            default:
                v25 := []int{0}
                _ = v25
            }
            fmt.Println(p1, p0)
        } else {
            var v26 float64 = 0.0
            _ = v26
        }
    }
    p0 = (p0 % (((p0 + (^(p0)))) | 1))
    _ = p0
    _ = p1
    _ = v7
    return ("foo bar" + ("\n\t" + "foo bar"))
}
func genFn1() int8 {
    if (true && (3.14159 < 1e10)) {
        var v27 uint64 = 42
        _ = v27
    } else {
        for fi28 := 0; fi28 < 4; fi28++ {
            v29 := []int{(fi28 - fi28), fi28, fi28, fi28}
            _ = fi28
            _ = v29
        }
    }
    var v30 int32 = -1
    for fi31 := 0; fi31 < 6; fi31++ {
        switch ("a" + (((("a" + "") + ("\n\t" + "")) + ("" + "a")) + "\n\t")) {
        case "hello":
            v30 = -1
        }
        fmt.Println(v30, fi31)
        if true {
            fmt.Println(fi31)
            var v32 uint64 = 233
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 uint64 = v32
                v35 := []int{fi31, (fi31 / ((0) | 1))}
                var v36 bool = (!true)
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            _ = v32
        }
        _ = fi31
    }
    var v37 int16 = -870
    _ = v30
    _ = v37
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 int = 1
    switch ("a" + (((("a" + "hello") + ("\n\t" + "\n\t")) + ("\n\t" + (("a" + ("a" + "foo bar")) + "hello"))) + "\n\t")) {
    case "":
        v4 = (v4 >> ((738) & 63))
    }
    switch "hello" {
    case "a":
        for fi38 := 0; fi38 < 7; fi38++ {
            var v39 uint64 = 453
            var v40 uint16 = 65535
            _ = fi38
            _ = v39
            _ = v40
        }
    case "foo bar":
        for fi41 := 0; fi41 < 18; fi41++ {
            if (true && (65535 >= 1)) {
                for fi42 := 0; fi42 < 14; fi42++ {
                    var v43 bool = true
                    _ = fi42
                    _ = v43
                }
                for fi44 := 0; fi44 < 12; fi44++ {
                    v45 := GenS2{F0: (("a" + ("" + (("\n\t" + "\n\t") + ("a" + ("foo bar" + (("foo bar" + "a") + "a")))))) + "a"), F1: (fi44 | v4), F2: 1}
                    v46 := GenS3{F0: 4294967295, F1: 1, F2: 127}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                _ = len(v3)
            }
            var v47 uint16 = 773
            v4 = -728
            _ = fi41
            _ = v47
        }
    case "":
        switch "a" {
        case "hello":
            switch "a" {
            case "":
                v4 -= v4
            }
        case "\n\t":
            v4 = v4
        default:
            fmt.Println(v4)
        }
    default:
        switch (v4 + v4) {
        case 128:
            v48 := GenS4{F0: 1}
            _ = v48
        case 127:
            fmt.Println(v4)
        default:
            switch ("hello" + "a") {
            case "foo bar":
                fmt.Println(v4)
            }
        }
    }
    fmt.Println(v4)
    v4 = 739
    v4 -= (v4 | (v4 % ((v4) | 1)))
    v4 += v4
    v49 := []int{(v4 * (^(v4))), (v4 + v4), 1, (v4 ^ v4)}
    var v50 bool = false
    switch (v4 << ((18446744073709551615) & 63)) {
    case 52:
        fmt.Println(v50, v50)
    case -827:
        for fi51 := 0; fi51 < 12; fi51++ {
            var v52 uint16 = 32767
            v3["foo bar"] = (v4 * v4)
            var v53 int8 = 1
            _ = fi51
            _ = v52
            _ = v53
        }
    case -243:
        var v54 uint8 = 2
        _ = v54
    default:
        var v55 float64 = (float64((-2.25 / 0.1)) * -2.25)
        _ = v55
    }
    for fi56 := 0; fi56 < 10; fi56++ {
        if (!((false && false) || v50)) {
            if (((("hello" + "") + "hello") + "a") >= "a") {
                fmt.Println(fi56, v4, v4)
                _ = len(v49)
                v50 = ((32767 != 42) && (!v50))
            }
            v4 += (fi56 ^ (fi56 % ((-352) | 1)))
            for fi57 := 0; fi57 < 7; fi57++ {
                var v58 float64 = (0.0 - -2.25)
                switch v4 {
                case 221:
                    var v59 int32 = -2147483648
                    _ = v59
                case 778:
                    switch "\n\t" {
                    case "":
                        if false {
                            v50 = v50
                            if false {
                                if (fi56 <= 2147483647) {
                                    if v50 {
                                        fmt.Println(v4, fi57)
                                        var v60 uint = 255
                                        v4 += (^(v4))
                                        v49 = append(v49, fi57)
                                        if len(v49) > 0 {
                                            v49[0] = (^(fi57))
                                        }
                                        var v61 int8 = -77
                                        if (!true) {
                                            var v62 uint = v60
                                            _ = v62
                                        } else {
                                            switch ("foo bar" + ("foo bar" + "foo bar")) {
                                            case "\n\t":
                                                var v63 int8 = (-(v61))
                                                _ = v63
                                            case "":
                                                fmt.Println(v50)
                                            case "foo bar":
                                                if ((v4 ^ (fi56 >> ((v60) & 63))) < 42) {
                                                    v61 += (v61 % ((v61) | 1))
                                                }
                                            default:
                                                var v64 uint64 = 2
                                                _ = v64
                                            }
                                            fmt.Println(fi56, fi57, v61)
                                        }
                                        if v50 {
                                            switch -115 {
                                            case -531:
                                                var v65 uint = (v60 & (v60 / ((v60) | 1)))
                                                _ = v65
                                            case 32767:
                                                switch ("a" + "\n\t") {
                                                case "hello":
                                                    if false {
                                                        fmt.Println(v4)
                                                        var v66 int16 = 42
                                                        if false {
                                                            v58 = (v58 - 1e10)
                                                        } else {
                                                            var v67 int32 = 0
                                                            if (float32(v58) == float32(v58)) {
                                                                v68 := GenS4{F0: 0}
                                                                if (231 <= 195) {
                                                                    if true {
                                                                        switch v4 {
                                                                        case -355:
                                                                            var v69 uint64 = uint64(v60)
                                                                            _ = v69
                                                                        case -748:
                                                                            var v70 uint64 = 4294967295
                                                                            _ = v70
                                                                        default:
                                                                            v66 = v66
                                                                        }
                                                                        fmt.Println(v67)
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            var v71 string = ""
                                                                            _ = v71
                                                                        case "a":
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "foo bar":
                                                                                fmt.Println(fi57)
                                                                            case "hello":
                                                                                var v72 float64 = v58
                                                                                _ = v72
                                                                            }
                                                                        default:
                                                                            v73 := GenS1{F0: (v60 & (v60 >> (((v60 * v60)) & 63))), F1: ""}
                                                                            _ = v73
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v58, v67)
                                                                        fmt.Println(fi56, v58)
                                                                    }
                                                                }
                                                                _ = v68
                                                            } else {
                                                                v67 += v67
                                                                v66 += v66
                                                            }
                                                            if v50 {
                                                                v67 = (v67 | (v67 + 803))
                                                                v50 = (v50 || v50)
                                                                var v74 int = fi56
                                                                _ = v74
                                                            } else {
                                                                var v75 uint32 = 1
                                                                v76 := GenS5{F0: (v67 >> (((^(v60))) & 63))}
                                                                delete(v3, "foo bar")
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            _ = v67
                                                        }
                                                        _ = v66
                                                    } else {
                                                        delete(v1, "")
                                                    }
                                                default:
                                                    fmt.Println(fi57)
                                                }
                                            case 128:
                                                v77 := GenS4{F0: 903}
                                                _ = v77
                                            }
                                            switch fi56 {
                                            case 2147483647:
                                                v78 := v1[""]
                                                _ = v78
                                            }
                                        } else {
                                            var v79 int16 = 1
                                            if (!false) {
                                                v80 := v3[("hello" + "\n\t")]
                                                v81 := GenS6{F0: 1}
                                                _ = v80
                                                _ = v81
                                            } else {
                                                var v82 uint32 = 4294967295
                                                _ = v82
                                            }
                                            _ = v79
                                        }
                                        _ = v60
                                        _ = v61
                                    } else {
                                        var v83 uint8 = 42
                                        var v84 uint16 = 521
                                        var v85 string = ""
                                        v84 -= v84
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    switch (fi57 % ((fi57) | 1)) {
                                    case 128:
                                        if len(v49) > 0 {
                                            v49[0] = fi57
                                        }
                                    }
                                }
                                if v50 {
                                    v58 -= 1.5
                                    var v86 bool = (1e10 == 0.1)
                                    v86 = ((float64((0.1 + -2.25)) * -2.25) != -2.25)
                                    _ = v86
                                } else {
                                    if v50 {
                                        if v50 {
                                            switch 33 {
                                            case 127:
                                                var v87 float64 = (3.14159 - v58)
                                                _ = v87
                                            case 9223372036854775807:
                                                if len(v49) > 0 {
                                                    v49[0] = v4
                                                }
                                            case 32767:
                                                switch 36 {
                                                case 2147483647:
                                                    v88 := GenS6{F0: -633}
                                                    _ = v88
                                                case 4294967295:
                                                    if v50 {
                                                        var v89 uint16 = 65535
                                                        if len(v49) >= 2 {
                                                            v49 = v49[1:len(v49)]
                                                        }
                                                        fmt.Println(fi57, v50)
                                                        _ = v89
                                                    } else {
                                                        var v90 uint32 = 2
                                                        switch fi57 {
                                                        case -683:
                                                            v58 = v58
                                                        case -658:
                                                            var v91 uint16 = 1
                                                            _ = v91
                                                        default:
                                                            switch (fi56 | fi56) {
                                                            case 127:
                                                                v92 := GenS1{F0: 65535, F1: "foo bar"}
                                                                _ = v92
                                                            case 32768:
                                                                if v50 {
                                                                    v93 := GenS2{F0: "hello", F1: (fi57 % (((fi56 ^ (v4 << ((161) & 63)))) | 1)), F2: 42}
                                                                    _ = v93
                                                                }
                                                            default:
                                                                var v94 uint16 = 943
                                                                _ = v94
                                                            }
                                                        }
                                                        _ = v90
                                                    }
                                                }
                                            default:
                                                v4 = -846
                                            }
                                        } else {
                                            _ = len(v49)
                                            v95 := v2["\n\t"]
                                            switch fi56 {
                                            case 349:
                                                fmt.Println(v58, fi56, fi56)
                                            case 942:
                                                v95 = v95
                                            case 32768:
                                                var v96 int64 = 207
                                                _ = v96
                                            }
                                            _ = v95
                                        }
                                    }
                                    v4 += (v4 & 0)
                                    if false {
                                        switch ("foo bar" + ("foo bar" + ("hello" + "hello"))) {
                                        case "hello":
                                            fmt.Println(v4)
                                            var v97 int = v4
                                            if false {
                                                var v98 int16 = 0
                                                var v99 string = "\n\t"
                                                var v100 int8 = genFn1()
                                                _ = v100
                                                _ = v98
                                                _ = v99
                                            }
                                            _ = v97
                                        case "foo bar":
                                            v101 := v1[("a" + "a")]
                                            _ = v101
                                        }
                                    } else {
                                        if (128 == 127) {
                                            if true {
                                                v58 += (v58 - (1e10 * (3.14159 + (v58 * v58))))
                                                fmt.Println(v50, fi57, fi57)
                                                v4 -= (^(fi56))
                                            }
                                            if v50 {
                                                fmt.Println(v50, fi56, v58)
                                                switch (v4 + (fi56 - fi56)) {
                                                case -732:
                                                    v102 := v3[(("a" + ("\n\t" + ("\n\t" + "hello"))) + ("foo bar" + "hello"))]
                                                    _ = v102
                                                case -711:
                                                    fmt.Println(v58, v4, fi56)
                                                default:
                                                    v49 = append(v49, fi57)
                                                }
                                            } else {
                                                if v50 {
                                                    var v103 int64 = 0
                                                    v58 = (((-2.25 + (v58 / (((v58 * (v58 + v58)) + (-2.25 - 1.5)) - ((1.5 + ((float64(0.0) + v58) - (-2.25 / v58))) * (0.1 - float64((3.14159 / 1e10))))))) - v58) - (v58 + 0.0))
                                                    switch ("a" + ("" + "\n\t")) {
                                                    case "":
                                                        _ = len(v49)
                                                    case "a":
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v104 string = ("foo bar" + "foo bar")
                                                            _ = v104
                                                        default:
                                                            switch ("" + "") {
                                                            case "foo bar":
                                                                v4 = (fi56 << ((4294967295) & 63))
                                                            case "\n\t":
                                                                v105 := GenS1{F0: 18446744073709551615, F1: (("hello" + "foo bar") + "foo bar")}
                                                                _ = v105
                                                            default:
                                                                fmt.Println(v103, v58)
                                                            }
                                                        }
                                                    case "hello":
                                                        v58 += ((v58 / ((v58 + (v58 * v58)) + (3.14159 * v58))) * float64(3.14159))
                                                    }
                                                    _ = v103
                                                } else {
                                                    var v106 bool = v50
                                                    var v107 float64 = v58
                                                    _ = v106
                                                    _ = v107
                                                }
                                                if v50 {
                                                    if len(v49) >= 2 {
                                                        v49 = v49[1:len(v49)]
                                                    }
                                                    if true {
                                                        var v108 int16 = 708
                                                        _ = v108
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            switch "" {
                                                            case "":
                                                                if v50 {
                                                                    var v109 float32 = 1.5
                                                                    _ = v109
                                                                }
                                                            case "\n\t":
                                                                if (0 < 127) {
                                                                    v110 := GenS7{F0: v4, F1: 42}
                                                                    v50 = v50
                                                                    if len(v49) >= 2 {
                                                                        v49 = v49[1:len(v49)]
                                                                    }
                                                                    _ = v110
                                                                }
                                                            case "foo bar":
                                                                v111 := GenS4{F0: 42}
                                                                _ = v111
                                                            default:
                                                                fmt.Println(fi56)
                                                            }
                                                        default:
                                                            v50 = (((255 > 0) && ((false && (4294967295 >= 0)) || false)) || (-128 != 106))
                                                        }
                                                        v112 := GenS8{F0: 2, F1: 965, F2: 65535}
                                                        _ = v112
                                                    }
                                                    if (!(("\n\t" + "") < "")) {
                                                        var v113 uint = 4294967295
                                                        _ = v113
                                                    } else {
                                                        var v114 uint64 = 617
                                                        v4 = -9223372036854775808
                                                        v58 = v58
                                                        _ = len(v49)
                                                        v58 += v58
                                                        _ = v114
                                                    }
                                                } else {
                                                    fmt.Println(fi57)
                                                    v58 = (1e10 * v58)
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        if (!(260 <= 1)) {
                                                            fmt.Println(fi56)
                                                            var v115 uint32 = 454
                                                            _ = v115
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            v4 = (fi56 >> ((65535) & 63))
                                                            v50 = (!true)
                                                        }
                                                    case "hello":
                                                        v58 = 1.5
                                                    }
                                                }
                                                v58 -= (((v58 * (v58 / float64((float32(v58) * ((((3.14159 / 1e10) * float32(-2.25)) - float32((0.0 / 3.14159))) * 0.1))))) - v58) - float64(0.0))
                                            }
                                            if true {
                                                if len(v49) > 0 {
                                                    v49[0] = -1
                                                }
                                                if (!(0.1 < v58)) {
                                                    v116 := GenS7{F0: (fi56 | fi56), F1: 42}
                                                    var v117 float64 = -2.25
                                                    _ = v116
                                                    _ = v117
                                                }
                                                var v118 string = "foo bar"
                                                _ = v118
                                            } else {
                                                v58 = v58
                                                if len(v49) > 0 {
                                                    v49[0] = (fi57 - v4)
                                                }
                                                v58 = v58
                                                v50 = v50
                                            }
                                            if len(v49) > 0 {
                                                v49[0] = fi57
                                            }
                                            fmt.Println(v50)
                                        } else {
                                            fmt.Println(fi56, fi56)
                                            v1["\n\t"] = false
                                            var v119 uint = 876
                                            _ = v119
                                        }
                                        v58 = v58
                                        fmt.Println(fi57, v4)
                                    }
                                }
                                v120 := GenS9{F0: v4, F1: v50}
                                v4 = (v4 % ((fi57) | 1))
                                switch "\n\t" {
                                case "a":
                                    v121 := GenS7{F0: fi57, F1: 42}
                                    _ = v121
                                case "hello":
                                    v50 = true
                                }
                                _ = v120
                            }
                        } else {
                            v4 += 9223372036854775807
                        }
                        switch v4 {
                        case -919:
                            if false {
                                fmt.Println(v50, fi57, v4)
                                if (-1 > 375) {
                                    v122 := GenS7{F0: 2147483647, F1: 511}
                                    v50 = v50
                                    _ = v122
                                } else {
                                    v123 := GenS7{F0: fi57, F1: 18446744073709551615}
                                    var v124 int = 32767
                                    _ = v123
                                    _ = v124
                                }
                            }
                        case 170:
                            var v125 uint = 4294967295
                            _ = v125
                        }
                        v126 := GenS10{F0: (fi56 % ((427) | 1)), F1: 32768, F2: 42}
                        _ = v126
                    case "foo bar":
                        delete(v3, ("\n\t" + ("foo bar" + ("foo bar" + "\n\t"))))
                    default:
                        v127 := GenS3{F0: 42, F1: 42, F2: 1}
                        _ = v127
                    }
                case 42:
                    var v128 int64 = 9223372036854775807
                    _ = v128
                }
                switch "a" {
                case "foo bar":
                    v4 = (fi56 ^ (^(fi56)))
                    var v129 int16 = 436
                    var v130 uint32 = 2147483647
                    _ = v129
                    _ = v130
                case "hello":
                    var v131 uint16 = 255
                    _ = v131
                case "\n\t":
                    v50 = false
                default:
                    fmt.Println(fi56, fi56)
                    var v132 bool = true
                    if false {
                        switch "" {
                        case "hello":
                            if v132 {
                                if (-1 > -1) {
                                    if len(v49) > 0 {
                                        v49[0] = (fi56 * 379)
                                    }
                                    var v133 float32 = (float32((v58 / v58)) / 1.5)
                                    _ = v133
                                }
                                fmt.Println(v50, fi57)
                                fmt.Println(fi56)
                            }
                            var v134 string = (((("a" + ("hello" + "hello")) + "a") + "a") + "hello")
                            _ = v134
                        case "\n\t":
                            if v132 {
                                fmt.Println(v132, v50)
                                v132 = false
                                delete(v2, "a")
                                var v135 int32 = -1
                                _ = v135
                            } else {
                                v3["\n\t"] = (v4 ^ (fi56 | (^(fi57))))
                            }
                            v136 := GenS11{F0: -630}
                            _ = v136
                        default:
                            v137 := GenS12{F0: 940, F1: 42}
                            _ = v137
                        }
                        fmt.Println(fi56, v50, v58)
                        var v138 uint8 = 1
                        fmt.Println(v58, v132, fi56)
                        v139 := GenS13{F0: (fi57 * 812), F1: v132}
                        _ = v138
                        _ = v139
                    }
                    _ = v132
                }
                _ = fi57
                _ = v58
            }
        }
        delete(v2, ("foo bar" + "hello"))
        var v140 int32 = 900
        _ = fi56
        _ = v140
    }
    var v141 float64 = 3.14159
    var v142 string = ("foo bar" + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
}
