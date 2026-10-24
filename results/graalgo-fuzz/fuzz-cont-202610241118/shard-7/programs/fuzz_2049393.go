package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 bool; F2 int8 }
type GenS7 struct { F0 uint; F1 int; F2 uint16 }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint; F2 uint8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool; F1 int64; F2 uint8 }
type GenS12 struct { F0 bool; F1 int8; F2 uint32 }
type GenS13 struct { F0 uint32; F1 int64; F2 string }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int; F1 uint32 }
type GenS16 struct { F0 uint16; F1 int8 }
type GenS17 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS18 struct { F0 uint32; F1 string }
type GenS19 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0(p0 int, p1 uint8) uint16 {
    v5 := GenS1{F0: 2147483647}
    v6 := []int{}
    v7 := GenS1{F0: -911}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int16) bool {
    fmt.Println(p0)
    p0 += (p0 % (((p0 << ((326) & 63))) | 1))
    _ = p0
    return (false || true)
}

func main() {
    v1 := []string{"a", (("foo bar" + "\n\t") + "")}
    v2 := []int{603, 2147483647}
    var v3 uint = 4294967295
    var v4 int = 127
    if false {
        for fi8 := 0; fi8 < 6; fi8++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (((genFn0(-1, 127) != 338) && false) || ((^(v4)) > v4)) {
                v9 := GenS1{F0: 651}
                fmt.Println(v3, v3, fi8)
                if false {
                    v10 := GenS2{F0: -2147483648}
                    _ = v10
                } else {
                    switch "\n\t" {
                    case "":
                        switch fi8 {
                        case 300:
                            var v11 int16 = 1
                            _ = v11
                        case 150:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    default:
                        if false {
                            for fi12 := 0; fi12 < 1; fi12++ {
                                v3 += (v3 << ((v3) & 63))
                                if true {
                                    switch fi8 {
                                    case -1:
                                        var v13 uint32 = 326
                                        _ = v13
                                    case 42:
                                        if false {
                                            if ((true && (!(127 > 222))) && (878 < 1)) {
                                                var v14 float64 = 1e10
                                                var v15 int32 = 0
                                                if false {
                                                    fmt.Println(v3)
                                                } else {
                                                    switch ("\n\t" + "\n\t") {
                                                    case "foo bar":
                                                        v3 -= (^(v3))
                                                    case "a":
                                                        var v16 string = (((("a" + "a") + ("a" + "a")) + "") + "a")
                                                        _ = v16
                                                    default:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v3 = 255
                                                        v15 -= v15
                                                    }
                                                    v2 = append(v2, 799)
                                                    v17 := GenS1{F0: (^(v15))}
                                                    _ = v17
                                                }
                                                _ = v14
                                                _ = v15
                                            }
                                            v3 = v3
                                        } else {
                                            if (100 == 110) {
                                                fmt.Println(v3, v3)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v18 := GenS1{F0: 0}
                                                v3 = (v3 << (((v3 / ((71) | 1))) & 63))
                                                v3 = (v3 >> ((v3) & 63))
                                                _ = v18
                                            } else {
                                                v3 += 42
                                                var v19 uint = v3
                                                _ = v19
                                            }
                                            v4 = v4
                                        }
                                    default:
                                        if (false && (!(1 > 42))) {
                                            var v20 float32 = 3.14159
                                            if (1.5 == 1e10) {
                                                if (!false) {
                                                    v2 = append(v2, fi8)
                                                    var v21 int32 = 73
                                                    _ = v21
                                                }
                                                fmt.Println(v4, fi12)
                                                v4 = (-(fi8))
                                                var v22 int64 = 42
                                                v22 = v22
                                                fmt.Println(v20)
                                                _ = v22
                                            }
                                            var v23 float64 = -2.25
                                            _ = v20
                                            _ = v23
                                        }
                                    }
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v24 float64 = (3.14159 * ((float64((1.5 / 3.14159)) * float64(float32(0.0))) * (3.14159 - 0.0)))
                                _ = fi12
                                _ = v24
                            }
                            if (false && ((418 >= v3) || (0 > 52))) {
                                v4 = fi8
                                var v25 uint64 = 18446744073709551615
                                _ = v25
                            }
                            if ((!("hello" < "a")) && false) {
                                v3 += 1
                                if (804 <= 0) {
                                    v4 += (fi8 / ((fi8) | 1))
                                }
                                switch (("hello" + ("hello" + (((("\n\t" + ("a" + ("\n\t" + "hello"))) + "\n\t") + "\n\t") + "hello"))) + "hello") {
                                case "foo bar":
                                    for fi26 := 0; fi26 < 9; fi26++ {
                                        var v27 uint64 = 0
                                        v4 = (fi26 + (fi8 << ((42) & 63)))
                                        v3 = (v3 | (v3 << ((v3) & 63)))
                                        if false {
                                            var v28 int32 = -473
                                            v29 := GenS1{F0: 1}
                                            v30 := GenS1{F0: -719}
                                            switch -1 {
                                            case -320:
                                                var v31 string = "foo bar"
                                                var v32 string = v31
                                                _ = v31
                                                _ = v32
                                            case 181:
                                                _ = len(v1)
                                            }
                                            v3 += v3
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                        _ = fi26
                                        _ = v27
                                    }
                                case "a":
                                    var v33 int = (-(fi8))
                                    _ = v33
                                default:
                                    var v34 int16 = -32768
                                    _ = v34
                                }
                            } else {
                                var v35 int = v4
                                _ = v35
                            }
                        } else {
                            v2 = append(v2, fi8)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v36 string = ("foo bar" + "hello")
                            _ = v36
                        }
                    }
                    v37 := GenS3{F0: ("a" + "hello"), F1: 0}
                    if len(v2) > 0 {
                        v2[0] = (v4 | (v4 % (((^(fi8))) | 1)))
                    }
                    _ = v37
                }
                _ = v9
            } else {
                v38 := GenS2{F0: -568}
                _ = v38
            }
            _ = fi8
        }
    } else {
        v1 = append(v1, "a")
        if ((false || false) || true) {
            switch "foo bar" {
            case "":
                v39 := GenS1{F0: -2147483648}
                _ = v39
            case "hello":
                switch v4 {
                case 128:
                    var v40 uint64 = 212
                    _ = v40
                }
            case "a":
                var v41 uint = (v3 - v3)
                _ = v41
            default:
                for fi42 := 0; fi42 < 5; fi42++ {
                    v4 = -1
                    v43 := GenS1{F0: -131}
                    v3 -= 4294967295
                    _ = fi42
                    _ = v43
                }
            }
            fmt.Println(v3, v4)
        }
        if (!(true || true)) {
            var v44 float64 = 0.1
            _ = v44
        }
    }
    v2 = append(v2, v4)
    fmt.Println(v3, v4)
    for fi45 := 0; fi45 < 15; fi45++ {
        var v46 uint64 = 2
        switch "a" {
        case "\n\t":
            _ = len(v1)
        }
        _ = fi45
        _ = v46
    }
    var v47 uint = (v3 / (((v3 | (v3 + (v3 & 2)))) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v48 bool = (!true)
    for fi49 := 0; fi49 < 10; fi49++ {
        var v50 uint8 = 42
        v50 = (v50 & v50)
        _ = fi49
        _ = v50
    }
    v47 = (v47 * (v47 << (((v47 - (v47 % ((v47) | 1)))) & 63)))
    switch "a" {
    case "a":
        switch v4 {
        case 1:
            if v48 {
                if true {
                    var v51 uint8 = 127
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v51
                } else {
                    v2 = append(v2, v4)
                    var v52 uint = v3
                    _ = v52
                }
            } else {
                v3 += (^(v47))
                _ = len(v1)
                var v53 int16 = 706
                _ = v53
            }
        case 0:
            if v48 {
                switch "" {
                case "\n\t":
                    v3 += v3
                }
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v54 := GenS4{F0: 42, F1: 778, F2: 65535}
                _ = v54
            }
        default:
            for fi55 := 0; fi55 < 5; fi55++ {
                var v56 uint64 = 0
                _ = fi55
                _ = v56
            }
        }
    case "":
        switch v4 {
        case -426:
            var v57 uint16 = 496
            _ = v57
        case -1:
            v58 := GenS5{F0: 1}
            _ = v58
        case 127:
            for fi59 := 0; fi59 < 9; fi59++ {
                if (!v48) {
                    var v60 uint = 912
                    switch v4 {
                    case 9223372036854775807:
                        for fi61 := 0; fi61 < 5; fi61++ {
                            var v62 int32 = 236
                            if v48 {
                                var v63 int32 = int32(v3)
                                _ = v63
                            }
                            v64 := GenS6{F0: 374, F1: (true || true), F2: 100}
                            _ = fi61
                            _ = v62
                            _ = v64
                        }
                    case -932:
                        v65 := GenS6{F0: genFn0(4294967295, 255), F1: true, F2: -128}
                        _ = v65
                    case -9223372036854775808:
                        if (127 > -84) {
                            v66 := GenS1{F0: -166}
                            v60 = 2
                            _ = v66
                        }
                    }
                    var v67 uint8 = 0
                    _ = v60
                    _ = v67
                } else {
                    if len(v2) > 0 {
                        v2[0] = v4
                    }
                    for fi68 := 0; fi68 < 7; fi68++ {
                        v4 += fi59
                        var v69 int8 = 42
                        _ = fi68
                        _ = v69
                    }
                    v3 = (v47 ^ (v3 + 405))
                }
                if (-32768 > 42) {
                    v70 := GenS1{F0: 2147483647}
                    v71 := GenS7{F0: v3, F1: v4, F2: 42}
                    _ = v70
                    _ = v71
                }
                _ = fi59
            }
        }
    }
    v72 := GenS8{F0: ("foo bar" + ("" + "foo bar")), F1: 350}
    switch 42 {
    case 4294967295:
        var v73 int8 = 89
        _ = v73
    }
    for fi74 := 0; fi74 < 12; fi74++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v75 uint = 241
        if v48 {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v76 uint16 = 1
            switch "foo bar" {
            case "hello":
                fmt.Println(fi74, v76, fi74)
            default:
                switch "" {
                case "\n\t":
                    v77 := GenS1{F0: -206}
                    _ = v77
                case "a":
                    for fi78 := 0; fi78 < 13; fi78++ {
                        if ((((("" + "a") + "hello") == ("" + "")) && (!("" != ""))) && v48) {
                            var v79 uint64 = 30
                            v75 += 4294967295
                            v3 = (v47 % ((v47) | 1))
                            var v80 int16 = 0
                            switch fi74 {
                            case 0:
                                v81 := GenS8{F0: ("a" + ("a" + "foo bar")), F1: 42}
                                _ = v81
                            }
                            _ = v79
                            _ = v80
                        }
                        _ = fi78
                    }
                case "":
                    switch (v4 | -9223372036854775808) {
                    case -965:
                        for fi82 := 0; fi82 < 17; fi82++ {
                            if ((0.0 - float64((float32(((0.0 / 0.1) * float64(((float32(((-2.25 / (((1e10 / (3.14159 + 1.5)) / 1e10) - ((0.0 - (0.1 / (float64(1.5) + -2.25))) - (float64(-2.25) + 0.0)))) * 0.1)) - float32(((0.0 + (0.0 * float64((3.14159 * (((0.1 / (float32(0.0) - 3.14159)) + (0.1 - -2.25)) * 0.0))))) - 3.14159))) + (float32(((3.14159 - float64(0.1)) + (-2.25 * (float64(-2.25) * 0.1)))) - 0.1))))) + float32((1.5 - 0.0))))) == 1e10) {
                                var v83 uint8 = 2
                                if v48 {
                                    switch "" {
                                    case "foo bar":
                                        if (v48 || (38 <= 36)) {
                                            if (v48 && (!(2147483647 == 42))) {
                                                var v84 int32 = 42
                                                if ((1e10 < (0.1 + 1e10)) && (uint64(v3) < uint64(v83))) {
                                                    fmt.Println(v84, v75, v84)
                                                    v47 += (v75 ^ v47)
                                                    fmt.Println(fi74)
                                                }
                                                if (1 <= 1) {
                                                    v85 := GenS3{F0: ("a" + "\n\t"), F1: 1000}
                                                    var v86 uint = (v3 + v47)
                                                    v87 := GenS3{F0: "foo bar", F1: 655}
                                                    fmt.Println(v76, v75)
                                                    v83 = v83
                                                    v75 -= v47
                                                    _ = v85
                                                    _ = v86
                                                    _ = v87
                                                }
                                                v88 := GenS5{F0: -1}
                                                v84 += v84
                                                _ = v84
                                                _ = v88
                                            }
                                            var v89 float64 = -2.25
                                            var v90 int16 = 32767
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v3 = v75
                                            v90 = v90
                                            _ = v89
                                            _ = v90
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "\n\t"
                                        }
                                    case "hello":
                                        v91 := GenS9{F0: 1, F1: v75, F2: v83}
                                        _ = v91
                                    case "":
                                        if (!v48) {
                                            var v92 uint16 = 841
                                            v93 := GenS2{F0: 42}
                                            fmt.Println(v92, v47, v47)
                                            _ = v92
                                            _ = v93
                                        } else {
                                            switch 128 {
                                            case 32767:
                                                if v48 {
                                                    v4 -= 127
                                                    v4 = v4
                                                    v83 = v83
                                                    fmt.Println(v76, v75)
                                                } else {
                                                    switch "" {
                                                    case "foo bar":
                                                        var v94 bool = v48
                                                        _ = v94
                                                    }
                                                }
                                                v95 := GenS7{F0: 932, F1: 227, F2: (^(v76))}
                                                v96 := GenS10{F0: v83}
                                                _ = v95
                                                _ = v96
                                            case -1:
                                                _ = len(v2)
                                            }
                                            v97 := GenS1{F0: int32(v47)}
                                            var v98 uint16 = (v76 & v76)
                                            _ = v97
                                            _ = v98
                                        }
                                    default:
                                        var v99 float64 = 1e10
                                        _ = v99
                                    }
                                } else {
                                    if v48 {
                                        var v100 int32 = 1
                                        var v101 int8 = -1
                                        _ = v100
                                        _ = v101
                                    } else {
                                        var v102 int32 = -662
                                        _ = v102
                                    }
                                    if (42 == 42) {
                                        switch (((("" + "a") + (("\n\t" + "\n\t") + "a")) + "a") + "") {
                                        case "foo bar":
                                            v4 = (^(fi74))
                                        case "a":
                                            var v103 int64 = 267
                                            _ = v103
                                        default:
                                            fmt.Println(v83, v75)
                                        }
                                        if (371 <= 4294967295) {
                                            v83 = (v83 >> ((42) & 63))
                                            switch -1 {
                                            case 32768:
                                                if (v4 == fi74) {
                                                    switch ("hello" + "") {
                                                    case "foo bar":
                                                        if (!true) {
                                                            fmt.Println(v76)
                                                            switch (fi82 - -1) {
                                                            case 839:
                                                                v2 = append(v2, fi74)
                                                            }
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = ("hello" + "a")
                                                            }
                                                            var v104 float64 = 3.14159
                                                            _ = len(v1)
                                                            _ = v104
                                                        }
                                                    default:
                                                        v48 = false
                                                    }
                                                    fmt.Println(v83, v4)
                                                } else {
                                                    var v105 string = "hello"
                                                    v3 -= v75
                                                    _ = v105
                                                }
                                            default:
                                                if (!(-1 == -128)) {
                                                    if (!false) {
                                                        var v106 int = 32768
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v3 -= v47
                                                            var v107 uint16 = 457
                                                            _ = v107
                                                        }
                                                        var v108 float32 = float32(1.5)
                                                        _ = v106
                                                        _ = v108
                                                    }
                                                    if (genFn0((^(fi82)), (^(v83))) <= v76) {
                                                        v47 = 598
                                                        if (-1 != 42) {
                                                            v109 := GenS11{F0: v48, F1: -80, F2: uint8(v3)}
                                                            _ = v109
                                                        }
                                                        if (int16(v76) != 0) {
                                                            var v110 uint64 = 0
                                                            switch "a" {
                                                            case "\n\t":
                                                                fmt.Println(fi74, v110)
                                                            case "hello":
                                                                fmt.Println(fi74)
                                                            default:
                                                                v2 = append(v2, fi74)
                                                            }
                                                            _ = v110
                                                        }
                                                    }
                                                    if ((!(v48 && (!true))) && ("hello" == "hello")) {
                                                        fmt.Println(fi82, fi82, v83)
                                                        v111 := GenS9{F0: 34, F1: (v75 | (^(v3))), F2: (v83 * 127)}
                                                        _ = v111
                                                    }
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v112 int = 187
                                                        _ = v112
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = "hello"
                                                    }
                                                }
                                            }
                                        } else {
                                            switch "hello" {
                                            case "":
                                                if true {
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v113 string = "hello"
                                                        _ = v113
                                                    case "":
                                                        fmt.Println(v47, fi74, v76)
                                                        if v48 {
                                                            if (806 > 4294967295) {
                                                                var v114 int = fi82
                                                                v115 := GenS9{F0: 1, F1: (^(v75)), F2: (v83 % ((v83) | 1))}
                                                                _ = v114
                                                                _ = v115
                                                            } else {
                                                                if ((v48 || (!true)) && (-1 < -27)) {
                                                                    v116 := GenS12{F0: v48, F1: -84, F2: 653}
                                                                    _ = v116
                                                                } else {
                                                                    if ((^(v75)) != 947) {
                                                                        v47 += (v47 * v3)
                                                                    }
                                                                }
                                                                var v117 float32 = (1.5 / -2.25)
                                                                switch "foo bar" {
                                                                case "":
                                                                    var v118 uint = (v75 * 65535)
                                                                    _ = v118
                                                                case "foo bar":
                                                                    switch ("hello" + "") {
                                                                    case "a":
                                                                        _ = len(v1)
                                                                    }
                                                                default:
                                                                    if (-809 >= 42) {
                                                                        var v119 float32 = (((-2.25 * v117) + v117) + 3.14159)
                                                                        v75 -= 18446744073709551615
                                                                        v120 := GenS13{F0: 2147483648, F1: -1, F2: "a"}
                                                                        _ = v119
                                                                        _ = v120
                                                                    }
                                                                }
                                                                _ = v117
                                                            }
                                                            if v48 {
                                                                var v121 float32 = 1e10
                                                                if (0 < v75) {
                                                                    fmt.Println(v47, v83, v76)
                                                                } else {
                                                                    v122 := GenS14{F0: 1}
                                                                    v75 -= (v47 & v75)
                                                                    _ = v122
                                                                }
                                                                if false {
                                                                    if (1 < 62) {
                                                                        fmt.Println(fi82, fi82, v4)
                                                                        v47 = (v75 % ((uint(v83)) | 1))
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (fi74 + (^(v4)))
                                                                        }
                                                                        v47 += 18446744073709551615
                                                                    }
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        if v48 {
                                                                            switch (("hello" + "hello") + "a") {
                                                                            case "a":
                                                                                _ = len(v1)
                                                                            }
                                                                            v83 += (^(v83))
                                                                            v47 -= (v3 << ((0) & 63))
                                                                            fmt.Println(v76, v3)
                                                                            v123 := GenS15{F0: (fi74 | v4), F1: 0}
                                                                            _ = v123
                                                                        }
                                                                    case "hello":
                                                                        v47 = 1
                                                                        switch "a" {
                                                                        case "a":
                                                                            switch fi74 {
                                                                            case 2147483647:
                                                                                v121 += ((v121 / v121) / -2.25)
                                                                            case -1:
                                                                                var v124 uint = (v47 << ((18446744073709551615) & 63))
                                                                                _ = v124
                                                                            }
                                                                        case "\n\t":
                                                                            var v125 int = (v4 ^ (fi82 << ((4294967295) & 63)))
                                                                            _ = v125
                                                                        }
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (-(fi74))
                                                                        }
                                                                        var v126 uint = v47
                                                                        var v127 uint8 = (v83 << ((v47) & 63))
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ("a" + (("" + (("foo bar" + "\n\t") + "foo bar")) + "foo bar"))
                                                                        }
                                                                        _ = v126
                                                                        _ = v127
                                                                    case "foo bar":
                                                                        var v128 string = ("a" + (("\n\t" + ("a" + ("\n\t" + ("hello" + ("a" + "a"))))) + ""))
                                                                        _ = v128
                                                                    default:
                                                                        v129 := GenS11{F0: (2 <= 2147483647), F1: int64(fi74), F2: (v83 & (v83 ^ (v83 / ((127) | 1))))}
                                                                        _ = v129
                                                                    }
                                                                }
                                                                _ = v121
                                                            }
                                                            switch 32767 {
                                                            case 0:
                                                                var v130 uint = (v47 * v75)
                                                                _ = v130
                                                            }
                                                        } else {
                                                            v131 := GenS15{F0: (fi82 & (fi82 - 0)), F1: uint32(v76)}
                                                            _ = v131
                                                        }
                                                        var v132 int64 = -1
                                                        var v133 int = 20
                                                        var v134 float32 = (((1.5 / -2.25) * 1.5) * (1.5 * 1e10))
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    case "a":
                                                        v47 = v75
                                                    }
                                                    v48 = v48
                                                    var v135 bool = v48
                                                    _ = v135
                                                } else {
                                                    v1 = append(v1, ("foo bar" + "foo bar"))
                                                    var v136 int8 = 100
                                                    v3 += v47
                                                    _ = v136
                                                }
                                            default:
                                                var v137 uint16 = (v76 | (v76 | genFn0(v4, (v83 - 165))))
                                                _ = v137
                                            }
                                        }
                                        v4 -= (fi74 * v4)
                                    } else {
                                        var v138 uint8 = 0
                                        if v48 {
                                            switch fi74 {
                                            case 325:
                                                var v139 bool = false
                                                _ = v139
                                            default:
                                                var v140 uint32 = 688
                                                _ = v140
                                            }
                                            var v141 float32 = 1.5
                                            _ = v141
                                        } else {
                                            fmt.Println(fi74, fi82, v47)
                                        }
                                        if false {
                                            var v142 int = (v4 + v4)
                                            _ = v142
                                        }
                                        _ = v138
                                    }
                                    v83 = 127
                                }
                                _ = v83
                            } else {
                                _ = len(v2)
                                v143 := GenS7{F0: 255, F1: v4, F2: v76}
                                var v144 string = "hello"
                                if ((!v48) || v48) {
                                    fmt.Println(fi82, fi74, fi74)
                                    if (-1 > -2147483648) {
                                        switch (fi74 ^ (v4 | fi82)) {
                                        case 0:
                                            var v145 uint8 = 42
                                            _ = v145
                                        case 1:
                                            v146 := GenS5{F0: 193}
                                            _ = v146
                                        default:
                                            if (((v3 * v3) == (v75 / ((0) | 1))) && false) {
                                                v144 = (v144 + v144)
                                            }
                                        }
                                        v47 = v75
                                        var v147 int64 = 42
                                        var v148 uint = (v47 % (((v47 ^ v47)) | 1))
                                        if len(v2) > 0 {
                                            v2[0] = fi74
                                        }
                                        _ = v147
                                        _ = v148
                                    } else {
                                        v48 = (((!true) && (v48 || v48)) && (v4 < 32768))
                                        fmt.Println(v47)
                                    }
                                    var v149 float32 = ((0.0 * -2.25) - 0.0)
                                    _ = v149
                                } else {
                                    if (2 >= 128) {
                                        switch (fi74 ^ (v4 | fi82)) {
                                        case 2147483647:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        case -828:
                                            switch ("\n\t" + "foo bar") {
                                            case "hello":
                                                fmt.Println(fi82, v3, v4)
                                            case "foo bar":
                                                fmt.Println(v144, fi74)
                                            case "a":
                                                var v150 float64 = 1e10
                                                _ = v150
                                            }
                                        default:
                                            fmt.Println(v4)
                                        }
                                    } else {
                                        fmt.Println(fi82)
                                        if (!true) {
                                            if (32767 < -595) {
                                                fmt.Println(v144, v4, v47)
                                                var v151 int = fi82
                                                v47 = v3
                                                v152 := GenS16{F0: (^(v76)), F1: -1}
                                                _ = v151
                                                _ = v152
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = v144
                                            }
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v76, v3, fi74)
                                            case "\n\t":
                                                if ((((v144 + (v144 + v144)) + "") + ((("hello" + v144) + ((v144 + "foo bar") + "\n\t")) + ((v144 + ("" + v144)) + ("" + ((v144 + "\n\t") + "a"))))) == ("" + v144)) {
                                                    var v153 float64 = 3.14159
                                                    v48 = (v3 > (^(v47)))
                                                    _ = v153
                                                } else {
                                                    v76 -= v76
                                                    _ = len(v2)
                                                    switch ("hello" + "\n\t") {
                                                    case "":
                                                        v154 := GenS6{F0: v76, F1: (587 == 1), F2: -128}
                                                        _ = v154
                                                    default:
                                                        if v48 {
                                                            fmt.Println(v4)
                                                            if v48 {
                                                                fmt.Println(v75)
                                                                fmt.Println(v47, v76, v4)
                                                            }
                                                        }
                                                    }
                                                    if v48 {
                                                        v47 = (v47 << (((v47 + 327)) & 63))
                                                        v3 += v47
                                                    } else {
                                                        fmt.Println(v48)
                                                    }
                                                    v47 = v75
                                                }
                                            case "hello":
                                                switch ("a" + "a") {
                                                case "":
                                                    v155 := GenS15{F0: (fi82 / ((-1000) | 1)), F1: 603}
                                                    _ = v155
                                                case "\n\t":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "a":
                                                    if (65535 == 0) {
                                                        fmt.Println(v4, v76)
                                                    }
                                                default:
                                                    var v156 int16 = 42
                                                    var v157 uint8 = 65
                                                    _ = v156
                                                    _ = v157
                                                }
                                                fmt.Println(v47, v3)
                                                if len(v1) > 0 {
                                                    v1[0] = (v144 + ((v144 + ("\n\t" + "\n\t")) + v144))
                                                }
                                            default:
                                                var v158 float64 = 1e10
                                                _ = v158
                                            }
                                            if (1.5 >= 1e10) {
                                                var v159 string = v144
                                                if ((!(((1.5 + (float32(float64(float32(1e10))) + 0.1)) - (0.0 - (0.1 - -2.25))) < 1.5)) && v48) {
                                                    v160 := GenS17{F0: 1, F1: 32767, F2: 128}
                                                    fmt.Println(v144, v3)
                                                    var v161 uint8 = 255
                                                    _ = v160
                                                    _ = v161
                                                } else {
                                                    fmt.Println(v4, v76, fi82)
                                                    v144 = v144
                                                }
                                                _ = v159
                                            } else {
                                                v1 = append(v1, "a")
                                            }
                                        }
                                        var v162 int16 = 1
                                        _ = v162
                                    }
                                    switch "a" {
                                    case "hello":
                                        var v163 bool = false
                                        _ = v163
                                    default:
                                        _ = len(v2)
                                    }
                                    v3 = v3
                                    v144 = (((("hello" + v144) + v144) + (v144 + v144)) + v144)
                                }
                                _ = v143
                                _ = v144
                            }
                            switch "\n\t" {
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v164 := GenS18{F0: 65535, F1: ((("a" + ("foo bar" + "foo bar")) + "hello") + "\n\t")}
                            if true {
                                fmt.Println(fi82, fi74, v3)
                                v47 = 1
                            }
                            _ = fi82
                            _ = v164
                        }
                    case 32768:
                        if v48 {
                            var v165 float64 = float64((0.0 - 1e10))
                            _ = v165
                        } else {
                            for fi166 := 0; fi166 < 16; fi166++ {
                                switch "a" {
                                case "foo bar":
                                    v167 := GenS19{F0: uint64(v76), F1: (v76 / (((v76 | (^(v76)))) | 1)), F2: 1}
                                    _ = v167
                                case "a":
                                    v75 = 42
                                case "\n\t":
                                    _ = len(v2)
                                default:
                                    v168 := GenS13{F0: 2147483647, F1: -736, F2: (("hello" + "foo bar") + "foo bar")}
                                    _ = v168
                                }
                                switch 271 {
                                case 32768:
                                    if false {
                                        switch "\n\t" {
                                        case "":
                                            var v169 int = (-(fi166))
                                            _ = v169
                                        case "\n\t":
                                            v170 := GenS2{F0: 0}
                                            _ = v170
                                        case "a":
                                            v1 = append(v1, ("hello" + "foo bar"))
                                        default:
                                            v47 = 64
                                        }
                                        v75 -= (v75 >> ((v47) & 63))
                                    } else {
                                        if ((1 != uint32(v75)) || (!v48)) {
                                            v3 = v75
                                            var v171 uint = (v75 << ((v47) & 63))
                                            fmt.Println(v3)
                                            _ = v171
                                        }
                                        v47 -= (^(v47))
                                    }
                                    var v172 int = fi74
                                    switch "\n\t" {
                                    case "hello":
                                        v75 = (v75 - v75)
                                    case "a":
                                        fmt.Println(v4, v76, v4)
                                    case "":
                                        switch (v172 + fi74) {
                                        case 0:
                                            var v173 int = -139
                                            _ = v173
                                        default:
                                            var v174 string = ("\n\t" + "")
                                            _ = v174
                                        }
                                    }
                                    _ = v172
                                case 42:
                                    fmt.Println(v76)
                                case 2147483647:
                                    var v175 float32 = 0.1
                                    _ = v175
                                }
                                v4 -= -373
                                _ = fi166
                            }
                            switch "" {
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                            case "foo bar":
                                v176 := GenS7{F0: 42, F1: fi74, F2: 497}
                                _ = v176
                            case "\n\t":
                                if v48 {
                                    var v177 uint = v75
                                    fmt.Println(v48, v47, v3)
                                    for fi178 := 0; fi178 < 8; fi178++ {
                                        fmt.Println(v177, v3, v47)
                                        switch (fi74 ^ (fi178 + fi74)) {
                                        case 128:
                                            var v179 uint = 0
                                            var v180 int64 = 0
                                            _ = v179
                                            _ = v180
                                        }
                                        _ = fi178
                                    }
                                    _ = v177
                                }
                            default:
                                var v181 int = (fi74 << (((v47 * v47)) & 63))
                                _ = v181
                            }
                        }
                    case 127:
                        v76 -= (v76 % (((v76 / ((32767) | 1))) | 1))
                    default:
                        v2 = append(v2, (fi74 ^ (v4 / ((1) | 1))))
                    }
                default:
                    switch -9223372036854775808 {
                    case 917:
                        v182 := GenS4{F0: 127, F1: 926, F2: 616}
                        _ = v182
                    default:
                        fmt.Println(v76, v3)
                    }
                }
            }
            _ = v76
        }
        _ = fi74
        _ = v75
    }
    var v183 uint16 = 32767
    var v184 int = (^(v4))
    var v185 uint = v3
    var v186 float64 = (0.1 + 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v183:", v183)
    fmt.Println("v184:", v184)
    fmt.Println("v185:", v185)
    fmt.Println("v186:", v186)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v183
    _ = v184
    _ = v185
    _ = v186
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v72
}
