package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int32; F2 string }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 int16; F2 uint32 }
type GenS8 struct { F0 uint; F1 int32; F2 uint }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int8; F1 int64; F2 int32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int; F1 int8 }
type GenS16 struct { F0 int16; F1 int16; F2 uint }
type GenS17 struct { F0 string; F1 int64 }
type GenS18 struct { F0 int32; F1 int8 }
type GenS19 struct { F0 uint16; F1 uint }
type GenS20 struct { F0 uint16; F1 int64; F2 uint64 }
func genFn0(p0 int, p1 uint32) string {
    for fi5 := 0; fi5 < 14; fi5++ {
        p1 = (^(p1))
        if true {
            var v6 uint32 = (p1 ^ 42)
            _ = v6
        }
        var v7 uint = 238
        _ = fi5
        _ = v7
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        v8 := []int{p0}
        _ = v8
    case "\n\t":
        switch "a" {
        case "a":
            v9 := []int{-746, (p0 - -456)}
            _ = v9
        case "":
            v10 := GenS1{F0: 42, F1: 2, F2: ("\n\t" + "hello")}
            _ = v10
        case "\n\t":
            fmt.Println(p0, p1)
        default:
            v11 := []int{(p0 / (((p0 >> ((18446744073709551615) & 63))) | 1))}
            _ = v11
        }
    case "hello":
        var v12 int8 = -100
        _ = v12
    default:
        if false {
            var v13 int16 = 32767
            var v14 uint16 = 1
            _ = v13
            _ = v14
        }
    }
    v15 := GenS1{F0: 1, F1: 1, F2: ("a" + "\n\t")}
    _ = p0
    _ = p1
    _ = v15
    return "hello"
}

func main() {
    v1 := []uint32{1}
    var v2 bool = (false || false)
    var v3 uint = 664
    var v4 int64 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v16 := GenS2{F0: (1 == v4)}
    v3 -= 4294967295
    for fi17 := 0; fi17 < 3; fi17++ {
        v1 = append(v1, 2)
        var v18 int = fi17
        for fi19 := 0; fi19 < 10; fi19++ {
            if v2 {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                var v20 float32 = 0.0
                if false {
                    v4 = v4
                    v18 = 0
                } else {
                    fmt.Println(v20, fi17, v2)
                    v3 = v3
                    var v21 uint8 = 2
                    _ = v21
                }
                v2 = v2
                _ = v20
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    if (!true) {
        _ = len(v1)
        var v22 int64 = (v4 ^ v4)
        v2 = ((-2.25 * (float64((0.1 + (1.5 + 0.0))) * (1.5 - ((1e10 + 0.1) / 1.5)))) > float64(float32(0.1)))
        _ = v22
    }
    v23 := GenS3{F0: 255}
    var v24 float32 = 0.0
    v3 -= v3
    if v2 {
        if v2 {
            var v25 float32 = (float32(0.1) * v24)
            if true {
                for fi26 := 0; fi26 < 8; fi26++ {
                    fmt.Println(v3, v4)
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v1 = append(v1, 0)
                        var v28 uint32 = 1
                        _ = fi27
                        _ = v28
                    }
                    var v29 string = (("\n\t" + "hello") + "a")
                    _ = fi26
                    _ = v29
                }
            } else {
                switch ("a" + ("foo bar" + "foo bar")) {
                case "hello":
                    if false {
                        _ = len(v1)
                    }
                case "foo bar":
                    var v30 int64 = v4
                    _ = v30
                case "\n\t":
                    if (317 < 42) {
                        for fi31 := 0; fi31 < 5; fi31++ {
                            if (1 != 0) {
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    fmt.Println(fi32)
                                    var v33 uint16 = 777
                                    v34 := GenS4{F0: 4294967295, F1: -2147483648, F2: ("foo bar" + ("\n\t" + "hello"))}
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch (-(fi31)) {
                                case 26:
                                    for fi35 := 0; fi35 < 7; fi35++ {
                                        if v2 {
                                            v36 := GenS5{F0: 219, F1: 230}
                                            v2 = (1 >= 2147483647)
                                            _ = v36
                                        }
                                        v4 -= -747
                                        v24 = v25
                                        if len(v1) > 0 {
                                            v1[0] = 0
                                        }
                                        v37 := GenS6{F0: 2}
                                        _ = fi35
                                        _ = v37
                                    }
                                }
                            }
                            if ((^(fi31)) >= (fi31 % ((32768) | 1))) {
                                switch ("a" + "a") {
                                case "hello":
                                    var v38 float64 = float64(v25)
                                    _ = v38
                                case "foo bar":
                                    for fi39 := 0; fi39 < 2; fi39++ {
                                        v40 := GenS7{F0: v3, F1: 630, F2: 0}
                                        v24 += ((-2.25 * 1.5) / (v24 + (v25 * v25)))
                                        if true {
                                            fmt.Println(v3, fi39)
                                            v4 -= 1
                                            v2 = true
                                        }
                                        v41 := GenS8{F0: v3, F1: 2147483647, F2: (v3 >> ((v3) & 63))}
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                case "\n\t":
                                    for fi42 := 0; fi42 < 1; fi42++ {
                                        var v43 int8 = 127
                                        _ = fi42
                                        _ = v43
                                    }
                                default:
                                    v44 := GenS3{F0: (v3 >> (((v3 & v3)) & 63))}
                                    _ = v44
                                }
                                if ((v24 + 3.14159) >= (3.14159 + v24)) {
                                    var v45 string = ("a" + ("foo bar" + "\n\t"))
                                    for fi46 := 0; fi46 < 4; fi46++ {
                                        var v47 float64 = 1.5
                                        _ = fi46
                                        _ = v47
                                    }
                                    _ = v45
                                }
                                fmt.Println(fi31, v3)
                            } else {
                                v4 = v4
                            }
                            for fi48 := 0; fi48 < 15; fi48++ {
                                var v49 float64 = ((((float64((0.1 / (v25 * 1e10))) - (((0.0 / (((3.14159 - -2.25) + 1.5) * 1e10)) - 1e10) - 3.14159)) - 3.14159) + 1e10) - (1.5 * 1.5))
                                v50 := GenS3{F0: (v3 >> (((v3 & (v3 * (v3 | v3)))) & 63))}
                                if (!v2) {
                                    var v51 int64 = (v4 & v4)
                                    _ = v51
                                }
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            _ = fi31
                        }
                        switch 2147483647 {
                        case -466:
                            var v52 int8 = -1
                            _ = v52
                        case 485:
                            var v53 uint64 = 4294967295
                            _ = v53
                        }
                    } else {
                        switch ("\n\t" + "hello") {
                        case "hello":
                            if v2 {
                                var v54 int64 = v4
                                v55 := GenS3{F0: (v3 + v3)}
                                _ = v54
                                _ = v55
                            } else {
                                var v56 int16 = -1
                                for fi57 := 0; fi57 < 18; fi57++ {
                                    var v58 bool = (2147483647 != 42)
                                    v4 -= 0
                                    _ = len(v1)
                                    _ = fi57
                                    _ = v58
                                }
                                if v2 {
                                    _ = len(v1)
                                } else {
                                    var v59 int32 = -2147483648
                                    for fi60 := 0; fi60 < 4; fi60++ {
                                        var v61 int16 = (v56 >> ((v3) & 63))
                                        if false {
                                            for fi62 := 0; fi62 < 14; fi62++ {
                                                var v63 int8 = 100
                                                v64 := GenS1{F0: 897, F1: 0, F2: "hello"}
                                                if (v24 >= 1e10) {
                                                    switch "hello" {
                                                    case "":
                                                        v65 := GenS9{F0: 0, F1: 0}
                                                        _ = v65
                                                    case "hello":
                                                        switch 32768 {
                                                        case 9223372036854775807:
                                                            if (v63 > v63) {
                                                                if (fi60 == 4294967295) {
                                                                    var v66 uint = (^(v3))
                                                                    var v67 int16 = v61
                                                                    var v68 string = "a"
                                                                    _ = v66
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = len(v1)
                                                            } else {
                                                                _ = len(v1)
                                                                if false {
                                                                    v4 -= v4
                                                                } else {
                                                                    v59 -= -2147483648
                                                                }
                                                            }
                                                        case 102:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    default:
                                                        v61 = (v61 % (((v56 + v61)) | 1))
                                                    }
                                                    v69 := GenS10{F0: v63, F1: 0, F2: (-(v59))}
                                                    _ = v69
                                                } else {
                                                    switch (fi62 << (((v3 - v3)) & 63)) {
                                                    case 0:
                                                        var v70 int32 = v59
                                                        _ = v70
                                                    case 2147483647:
                                                        var v71 uint = v3
                                                        _ = v71
                                                    }
                                                }
                                                switch "a" {
                                                case "\n\t":
                                                    var v72 int64 = -9223372036854775808
                                                    _ = v72
                                                }
                                                _ = fi62
                                                _ = v63
                                                _ = v64
                                            }
                                            if v2 {
                                                if len(v1) > 0 {
                                                    v1[0] = 65535
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                for fi73 := 0; fi73 < 7; fi73++ {
                                                    if len(v1) > 0 {
                                                        v1[0] = 24
                                                    }
                                                    var v74 bool = v2
                                                    v1 = append(v1, 2147483648)
                                                    if len(v1) > 0 {
                                                        v1[0] = 4294967295
                                                    }
                                                    switch (fi60 / ((32767) | 1)) {
                                                    case 32767:
                                                        var v75 float32 = v24
                                                        _ = v75
                                                    case 4294967295:
                                                        var v76 string = "foo bar"
                                                        _ = v76
                                                    case 0:
                                                        var v77 uint32 = 1
                                                        _ = v77
                                                    }
                                                    switch fi60 {
                                                    case 1:
                                                        if false {
                                                            switch fi73 {
                                                            case -1:
                                                                var v78 string = ("" + "a")
                                                                if (int16(v4) == v61) {
                                                                    var v79 int32 = v59
                                                                    v2 = v2
                                                                    _ = v79
                                                                }
                                                                _ = v78
                                                            }
                                                            if false {
                                                                v56 = v61
                                                                if (1e10 != float64(v24)) {
                                                                    switch (fi60 * fi60) {
                                                                    case 0:
                                                                        if (!v2) {
                                                                            var v80 float64 = 1.5
                                                                            _ = v80
                                                                        }
                                                                    case -1:
                                                                        var v81 int32 = v59
                                                                        _ = v81
                                                                    default:
                                                                        fmt.Println(v24, v24, v56)
                                                                        var v82 int8 = -53
                                                                        var v83 float64 = ((float64(1e10) - -2.25) - 1e10)
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                } else {
                                                                    var v84 int16 = (v56 * v61)
                                                                    v85 := GenS3{F0: 1}
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                            } else {
                                                                v2 = v2
                                                                switch (fi60 | fi73) {
                                                                case -871:
                                                                    var v86 uint16 = 255
                                                                    _ = v86
                                                                case 498:
                                                                    v2 = v74
                                                                case -1:
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        fmt.Println(v25)
                                                                    case "\n\t":
                                                                        var v87 int16 = int16(v59)
                                                                        _ = v87
                                                                    case "foo bar":
                                                                        v59 -= -621
                                                                    default:
                                                                        v88 := GenS7{F0: 241, F1: v56, F2: 65535}
                                                                        _ = v88
                                                                    }
                                                                }
                                                                var v89 string = (("" + "hello") + "\n\t")
                                                                switch (fi73 >> (((v3 / (((v3 / (((v3 + 927)) | 1))) | 1))) & 63)) {
                                                                case 1:
                                                                    if (v89 != ("foo bar" + v89)) {
                                                                        var v90 uint16 = 1
                                                                        _ = v90
                                                                    } else {
                                                                        fmt.Println(v56, fi60)
                                                                        v24 -= float32((-2.25 + 1.5))
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 2
                                                                        }
                                                                    }
                                                                }
                                                                _ = v89
                                                            }
                                                            v91 := GenS11{F0: (v3 >> ((641) & 63))}
                                                            _ = v91
                                                        }
                                                    case -9223372036854775808:
                                                        v4 += v4
                                                    case 32767:
                                                        v59 += (v59 & v59)
                                                    }
                                                    if (-128 != -1) {
                                                        switch (fi73 & fi60) {
                                                        case 2147483647:
                                                            var v92 bool = (fi60 == (fi73 << ((v3) & 63)))
                                                            _ = v92
                                                        default:
                                                            v93 := GenS12{F0: -9223372036854775808}
                                                            _ = v93
                                                        }
                                                    } else {
                                                        var v94 string = ""
                                                        _ = v94
                                                    }
                                                    _ = fi73
                                                    _ = v74
                                                }
                                            }
                                            switch (fi60 ^ int(v3)) {
                                            case 9223372036854775807:
                                                for fi95 := 0; fi95 < 17; fi95++ {
                                                    switch "hello" {
                                                    case "a":
                                                        var v96 uint64 = 4294967295
                                                        _ = v96
                                                    case "foo bar":
                                                        v24 = v24
                                                    case "":
                                                        v97 := GenS13{F0: 128}
                                                        _ = v97
                                                    }
                                                    var v98 int16 = 0
                                                    _ = len(v1)
                                                    _ = fi95
                                                    _ = v98
                                                }
                                            case -45:
                                                var v99 bool = v2
                                                _ = v99
                                            default:
                                                fmt.Println(v25, v24)
                                            }
                                        }
                                        if v2 {
                                            v100 := GenS14{F0: 100}
                                            fmt.Println(fi60, fi60)
                                            _ = v100
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                for fi101 := 0; fi101 < 9; fi101++ {
                                                    v1 = append(v1, 1)
                                                    if (!false) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        switch "hello" {
                                                        case "foo bar":
                                                            fmt.Println(v56, v24)
                                                        case "\n\t":
                                                            fmt.Println(v2, v3)
                                                        case "":
                                                            var v102 uint32 = 2147483648
                                                            _ = v102
                                                        }
                                                    } else {
                                                        v56 += v61
                                                        var v103 int32 = (^(v59))
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v104 uint = (v3 & v3)
                                                            _ = v104
                                                        case "\n\t":
                                                            if (2147483648 != 2147483648) {
                                                                if (v2 || (v3 > (^(v3)))) {
                                                                    fmt.Println(v61)
                                                                    v61 = -324
                                                                } else {
                                                                    if (2147483648 > 1) {
                                                                        var v105 uint16 = 2
                                                                        fmt.Println(v61, v3)
                                                                        fmt.Println(v105, v24)
                                                                        v24 = v25
                                                                        if (v105 >= v105) {
                                                                            switch 127 {
                                                                            case 333:
                                                                                switch fi101 {
                                                                                case 32767:
                                                                                    switch (fi101 & fi60) {
                                                                                    case 1:
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 917
                                                                                        }
                                                                                    }
                                                                                case 534:
                                                                                    if v2 {
                                                                                        if v2 {
                                                                                            if v2 {
                                                                                                v106 := GenS8{F0: (v3 * 18446744073709551615), F1: v59, F2: 18446744073709551615}
                                                                                                _ = len(v1)
                                                                                                _ = v106
                                                                                            } else {
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = 4294967295
                                                                                                }
                                                                                                var v107 uint8 = 127
                                                                                                _ = v107
                                                                                            }
                                                                                        }
                                                                                        if (!(v2 && v2)) {
                                                                                            v108 := GenS15{F0: (fi60 ^ fi101), F1: 127}
                                                                                            v4 -= (v4 * -294)
                                                                                            _ = v108
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v3 = (v3 - (v3 | v3))
                                                                                }
                                                                            case -2147483648:
                                                                                v109 := GenS5{F0: 503, F1: 172}
                                                                                _ = len(v1)
                                                                                var v110 int = (fi60 << (((v3 * 355)) & 63))
                                                                                _ = v109
                                                                                _ = v110
                                                                            case 127:
                                                                                var v111 uint16 = 731
                                                                                _ = v111
                                                                            }
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v59 -= int32(v4)
                                                                            case "":
                                                                                var v112 int32 = v103
                                                                                _ = v112
                                                                            case "foo bar":
                                                                                v105 = (v105 & v105)
                                                                            }
                                                                            v113 := GenS2{F0: (((("hello" + "a") + "\n\t") + "a") != (("a" + "a") + "foo bar"))}
                                                                            _ = v113
                                                                        }
                                                                        if false {
                                                                            switch fi101 {
                                                                            case 1:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 2147483647
                                                                                }
                                                                            case 128:
                                                                                fmt.Println(v103)
                                                                            default:
                                                                                v24 -= (v25 / v24)
                                                                            }
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                switch (^(fi60)) {
                                                                                case 0:
                                                                                    if true {
                                                                                        v2 = v2
                                                                                    }
                                                                                case 353:
                                                                                    if (float64((v25 * float32(1e10))) >= 1.5) {
                                                                                        fmt.Println(v3, v4)
                                                                                        var v114 int8 = 42
                                                                                        switch fi60 {
                                                                                        case 0:
                                                                                            fmt.Println(fi101, v3, v59)
                                                                                        case -759:
                                                                                            var v115 bool = v2
                                                                                            _ = v115
                                                                                        default:
                                                                                            var v116 int8 = v114
                                                                                            _ = v116
                                                                                        }
                                                                                        _ = v114
                                                                                    }
                                                                                case 32767:
                                                                                    if (!false) {
                                                                                        if (v2 || ((v3 ^ v3) == 255)) {
                                                                                            v103 += int32(v56)
                                                                                            if (true || true) {
                                                                                                fmt.Println(v56)
                                                                                                fmt.Println(fi60, v105, v56)
                                                                                                v117 := GenS16{F0: v56, F1: v56, F2: v3}
                                                                                                _ = v117
                                                                                            }
                                                                                            v118 := GenS9{F0: 127, F1: (^(v103))}
                                                                                            _ = v118
                                                                                        } else {
                                                                                            var v119 int64 = v4
                                                                                            switch -1 {
                                                                                            case 32768:
                                                                                                var v120 int8 = -128
                                                                                                _ = v120
                                                                                            }
                                                                                            _ = v119
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 65535
                                                                                    }
                                                                                }
                                                                            case "foo bar":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 2147483648
                                                                                }
                                                                            }
                                                                        }
                                                                        if (v2 || (127 > 127)) {
                                                                            v3 += v3
                                                                        } else {
                                                                            var v121 float64 = ((0.0 / 1.0) * 0.1)
                                                                            _ = v121
                                                                        }
                                                                        _ = v105
                                                                    }
                                                                }
                                                                if (!(((uint8(v3) != 77) && v2) && (32767 < 753))) {
                                                                    if v2 {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        fmt.Println(fi101, v103)
                                                                        v61 = 0
                                                                    } else {
                                                                        v2 = false
                                                                    }
                                                                    if v2 {
                                                                        v122 := GenS17{F0: "\n\t", F1: (v4 & int64(v103))}
                                                                        var v123 int = (fi101 << ((v3) & 63))
                                                                        _ = v122
                                                                        _ = v123
                                                                    } else {
                                                                        v124 := GenS15{F0: -718, F1: 27}
                                                                        _ = v124
                                                                    }
                                                                    var v125 bool = false
                                                                    _ = v125
                                                                } else {
                                                                    var v126 int16 = (v61 >> ((v3) & 63))
                                                                    var v127 int = -999
                                                                    _ = v126
                                                                    _ = v127
                                                                }
                                                            } else {
                                                                var v128 float64 = 1.5
                                                                if v2 {
                                                                    var v129 int = fi60
                                                                    var v130 int32 = -2147483648
                                                                    fmt.Println(v2)
                                                                    _ = v129
                                                                    _ = v130
                                                                }
                                                                _ = v128
                                                            }
                                                        }
                                                        if (!(0 != 2)) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            v131 := GenS8{F0: (v3 << ((v3) & 63)), F1: 42, F2: v3}
                                                            _ = v131
                                                        } else {
                                                            v132 := GenS14{F0: -128}
                                                            if (-128 == 69) {
                                                                v1 = append(v1, 65535)
                                                                var v133 float32 = (v25 + (v24 / -2.25))
                                                                v25 = 0.0
                                                                _ = v133
                                                            }
                                                            var v134 float64 = 3.14159
                                                            var v135 int64 = v4
                                                            _ = v132
                                                            _ = v134
                                                            _ = v135
                                                        }
                                                        _ = len(v1)
                                                        _ = v103
                                                    }
                                                    v59 -= int32(v3)
                                                    if ((false || (51 == uint8(v3))) || (!((255 == 0) || false))) {
                                                        v2 = ((v59 >> (((v3 | (v3 + v3))) & 63)) <= 1)
                                                    } else {
                                                        var v136 int = (fi60 + (fi101 & int(v56)))
                                                        _ = v136
                                                    }
                                                    v25 -= (v24 / (3.14159 - 0.1))
                                                    _ = fi101
                                                }
                                            case "hello":
                                                var v137 int16 = -32768
                                                _ = v137
                                            default:
                                                if ((428 < 32767) && (!(uint8(v3) < 2))) {
                                                    var v138 string = ""
                                                    if (true && true) {
                                                        var v139 uint32 = 42
                                                        _ = v139
                                                    } else {
                                                        v140 := GenS18{F0: 42, F1: 100}
                                                        _ = v140
                                                    }
                                                    v141 := GenS19{F0: 42, F1: (v3 % (((^(v3))) | 1))}
                                                    _ = v138
                                                    _ = v141
                                                }
                                            }
                                        }
                                        _ = fi60
                                        _ = v61
                                    }
                                    for fi142 := 0; fi142 < 12; fi142++ {
                                        if len(v1) > 0 {
                                            v1[0] = 4294967295
                                        }
                                        _ = fi142
                                    }
                                    _ = v59
                                }
                                _ = v56
                            }
                        default:
                            if (((((((float32(1.5) / v24) - (v25 + v25)) + v25) / v25) != 3.14159) || (!v2)) || true) {
                                for fi143 := 0; fi143 < 8; fi143++ {
                                    var v144 float32 = (-2.25 * float32((-2.25 / 3.14159)))
                                    v145 := GenS20{F0: 353, F1: v4, F2: 42}
                                    v4 += (v4 % ((-281) | 1))
                                    _ = fi143
                                    _ = v144
                                    _ = v145
                                }
                            } else {
                                fmt.Println(v2, v25, v3)
                                switch "\n\t" {
                                case "hello":
                                    var v146 int = 846
                                    _ = v146
                                case "":
                                    var v147 uint32 = 65535
                                    _ = v147
                                case "foo bar":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                        }
                    }
                }
                var v148 int32 = 690
                var v149 bool = v2
                _ = v148
                _ = v149
            }
            _ = v25
        } else {
            switch -324 {
            case -857:
                v150 := GenS19{F0: 1, F1: 18446744073709551615}
                _ = v150
            case 128:
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
            case -177:
                fmt.Println(v4, v24)
            default:
                switch -1 {
                case 9223372036854775807:
                    var v151 uint64 = 207
                    _ = v151
                case -9223372036854775808:
                    var v152 int = 396
                    _ = v152
                case 0:
                    if (("a" + "foo bar") < "") {
                        v4 = 60
                        switch "\n\t" {
                        case "hello":
                            if v2 {
                                _ = len(v1)
                            }
                        default:
                            v2 = false
                        }
                        switch "hello" {
                        case "hello":
                            if ((1e10 < 1.5) || (!(!(!((!v2) || false))))) {
                                var v153 int16 = -819
                                var v154 float64 = 0.0
                                _ = v153
                                _ = v154
                            }
                        case "a":
                            v3 = (v3 - v3)
                        default:
                            v155 := GenS7{F0: 506, F1: 0, F2: 282}
                            _ = v155
                        }
                    }
                default:
                    var v156 uint8 = 230
                    _ = v156
                }
            }
        }
        fmt.Println(v2)
    }
    switch "a" {
    case "":
        var v157 int64 = v4
        _ = v157
    case "a":
        if (v3 <= (v3 % ((v3) | 1))) {
            switch "foo bar" {
            case "hello":
                v158 := GenS10{F0: 0, F1: v4, F2: 673}
                _ = v158
            }
            for fi159 := 0; fi159 < 14; fi159++ {
                var v160 uint = 2
                fmt.Println(v2)
                _ = fi159
                _ = v160
            }
        }
    }
    if v2 {
        var v161 uint16 = 633
        _ = v161
    } else {
        v1 = append(v1, 584)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
