package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 string }
type GenS5 struct { F0 int16; F1 int32; F2 int8 }
type GenS6 struct { F0 string; F1 uint32; F2 uint }
type GenS7 struct { F0 int; F1 string; F2 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 bool; F1 string; F2 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 3.14159
    v3 := make(map[string]bool)
    v4 := []int{4294967295, 4294967295, 0, 2147483647}
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v5 uint8 = 0
    var v6 float32 = (((v2 / 3.14159) - -2.25) + v2)
    var v7 uint32 = uint32(v5)
    fmt.Println(v2, v5, v7)
    fmt.Println(v2)
    fmt.Println(v2, v5)
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(v5)
        _ = fi8
    }
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    v9 := GenS1{F0: (994 == -948), F1: v7, F2: false}
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint64 = 0
            switch "foo bar" {
            case "\n\t":
                switch "\n\t" {
                case "":
                    v7 = 0
                case "\n\t":
                    v6 -= (3.14159 / v6)
                default:
                    if len(v4) > 0 {
                        v4[0] = (fi11 + fi11)
                    }
                }
                fmt.Println(v6, fi10)
            case "hello":
                v7 -= v7
            default:
                if true {
                    v4 = append(v4, (^(fi11)))
                    var v13 int = (-(fi10))
                    var v14 int8 = -128
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(fi11, v7, fi10)
                    if (247 <= (v7 ^ (v7 - (v7 / (((^(v7))) | 1))))) {
                        var v15 uint8 = 89
                        if false {
                            if false {
                                var v16 uint32 = (^(v7))
                                _ = v16
                            } else {
                                _ = len(v1)
                                v2 += v6
                                if (!true) {
                                    var v17 float64 = 0.0
                                    var v18 bool = (59 <= -1)
                                    _ = v17
                                    _ = v18
                                } else {
                                    if true {
                                        v1[(v7 * v7)] = ("a" + ("foo bar" + "foo bar"))
                                        v19 := GenS1{F0: false, F1: v7, F2: true}
                                        _ = v19
                                    } else {
                                        v5 += 128
                                        v20 := GenS1{F0: false, F1: v7, F2: false}
                                        v4 = append(v4, fi11)
                                        switch (fi10 - 9223372036854775807) {
                                        case -421:
                                            var v21 int = fi10
                                            _ = v21
                                        case -583:
                                            fmt.Println(fi10, fi11)
                                        case 95:
                                            if true {
                                                v5 += v5
                                                var v22 uint16 = 1
                                                _ = v22
                                            } else {
                                                fmt.Println(v7, v2, v2)
                                                v5 = 179
                                            }
                                        }
                                        _ = v20
                                    }
                                }
                                v23 := GenS1{F0: ((2 >= 987) || true), F1: (v7 % ((uint32(v5)) | 1)), F2: false}
                                _ = v23
                            }
                            var v24 string = ("foo bar" + ("\n\t" + "a"))
                            _ = v24
                        }
                        _ = v15
                    } else {
                        var v25 string = "hello"
                        v26 := GenS2{F0: v7, F1: 42}
                        var v27 uint16 = 65535
                        if (!true) {
                            v6 += (v2 + 0.0)
                        } else {
                            if (false || (false || (!(0.1 != 1e10)))) {
                                var v28 uint16 = (v27 | v27)
                                _ = v28
                            } else {
                                v25 = v25
                                switch (fi10 >> ((42) & 63)) {
                                case 9223372036854775807:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case -1:
                                    v29 := GenS2{F0: 4294967295, F1: 18446744073709551615}
                                    _ = v29
                                case 4294967295:
                                    v12 -= (^(v12))
                                }
                                var v30 float64 = 3.14159
                                _ = v30
                            }
                        }
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                }
            }
            _ = fi11
            _ = v12
        }
        v2 = (0.0 * v2)
        switch ("a" + ("hello" + "")) {
        case "":
            for fi31 := 0; fi31 < 13; fi31++ {
                v32 := GenS3{F0: true}
                v7 -= 65535
                v4 = append(v4, 549)
                _ = fi31
                _ = v32
            }
        case "hello":
            if true {
                v33 := v1[0]
                _ = v33
            }
        default:
            v7 = 865
        }
        _ = fi10
    }
    if true {
        for fi34 := 0; fi34 < 16; fi34++ {
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := v1[v7]
                v37 := GenS2{F0: v7, F1: uint64(v7)}
                if (!false) {
                    var v38 int32 = 42
                    var v39 string = ((((v36 + v36) + (v36 + v36)) + (("\n\t" + ("hello" + v36)) + "a")) + "foo bar")
                    if false {
                        fmt.Println(v7)
                        fmt.Println(v7, v36)
                    } else {
                        switch int(v5) {
                        case 32768:
                            var v40 uint16 = 1
                            _ = v40
                        case 42:
                            if false {
                                switch v39 {
                                case "\n\t":
                                    if false {
                                        if (0 <= fi35) {
                                            switch "a" {
                                            case "a":
                                                v41 := GenS1{F0: (2 < 2), F1: (^(v7)), F2: false}
                                                _ = v41
                                            default:
                                                v42 := GenS4{F0: 65535, F1: 42, F2: ((((v39 + "foo bar") + (v39 + "foo bar")) + "\n\t") + (v36 + ""))}
                                                _ = v42
                                            }
                                        } else {
                                            v7 += 415
                                            if (((v5 % ((128) | 1)) > uint8(v7)) && true) {
                                                v2 += (v6 + v6)
                                                v39 = "hello"
                                                fmt.Println(v7, v5, v2)
                                                if (v38 > (v38 % ((v38) | 1))) {
                                                    switch (fi35 ^ 1) {
                                                    case 4294967295:
                                                        fmt.Println(v2, fi34)
                                                    case 32768:
                                                        if (!true) {
                                                            fmt.Println(v38)
                                                            v7 = 1
                                                            fmt.Println(v2, v38, v2)
                                                        }
                                                    case 128:
                                                        v43 := GenS2{F0: (v7 ^ v7), F1: 593}
                                                        _ = v43
                                                    default:
                                                        if true {
                                                            v7 += (^(v7))
                                                            var v44 int32 = v38
                                                            v38 = (v38 << ((4294967295) & 63))
                                                            _ = v44
                                                        }
                                                    }
                                                    if ((!true) || false) {
                                                        fmt.Println(v38)
                                                        fmt.Println(v2, fi35)
                                                        delete(v3, v39)
                                                        v45 := GenS1{F0: true, F1: v7, F2: false}
                                                        var v46 int = (fi35 << ((0) & 63))
                                                        v38 = (v38 / (((v38 / (((-(v38))) | 1))) | 1))
                                                        v47 := GenS2{F0: 387, F1: 683}
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    v38 += v38
                                                    fmt.Println(v6)
                                                    _ = len(v4)
                                                    var v48 bool = true
                                                    switch (fi35 * fi34) {
                                                    case 32767:
                                                        var v49 string = "\n\t"
                                                        _ = v49
                                                    }
                                                    if v48 {
                                                        v50 := GenS1{F0: ((v38 * 229) <= (v38 ^ v38)), F1: (v7 | 151), F2: v48}
                                                        v6 = (v2 / (-2.25 * float32(-2.25)))
                                                        if v48 {
                                                            v36 = "foo bar"
                                                        } else {
                                                            v36 = ""
                                                            var v51 bool = false
                                                            _ = v51
                                                        }
                                                        var v52 uint64 = 60
                                                        _ = v50
                                                        _ = v52
                                                    }
                                                    _ = v48
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = (fi35 ^ 42)
                                                }
                                                var v53 int8 = 127
                                                delete(v1, v7)
                                                v54 := GenS5{F0: 1, F1: v38, F2: -116}
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        v2 -= (v6 / (v2 * ((v6 + float32((((0.0 - (3.14159 - ((1.5 * 0.0) - (3.14159 / ((1.5 * 3.14159) * float64(v2)))))) + (1.5 + float64(v6))) - 0.1))) - float32(float64((0.1 * 3.14159))))))
                                    }
                                case "foo bar":
                                    switch fi35 {
                                    case -833:
                                        var v55 int16 = -1
                                        _ = v55
                                    case -9223372036854775808:
                                        if (v36 < (v36 + "")) {
                                            switch ("\n\t" + (v39 + (v39 + "a"))) {
                                            case "a":
                                                switch fi34 {
                                                case 9223372036854775807:
                                                    if false {
                                                        fmt.Println(v36, v39)
                                                        v7 = 1
                                                    }
                                                }
                                            case "":
                                                fmt.Println(v7, fi35)
                                            }
                                            v56 := GenS2{F0: v7, F1: 305}
                                            _ = v56
                                        }
                                    case 0:
                                        v5 += v5
                                    }
                                }
                            } else {
                                v36 = v39
                                fmt.Println(v2, v7, v7)
                                if true {
                                    var v57 int64 = 149
                                    _ = v57
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = fi34
                                    }
                                    v4 = append(v4, fi34)
                                }
                                switch v36 {
                                case "hello":
                                    fmt.Println(v2, v39)
                                    v2 += v6
                                default:
                                    switch 32767 {
                                    case 9223372036854775807:
                                        var v58 uint16 = 42
                                        _ = v58
                                    case -9223372036854775808:
                                        var v59 uint64 = 923
                                        _ = v59
                                    }
                                }
                            }
                        case 127:
                            v38 = (v38 / ((42) | 1))
                        default:
                            v38 = (v38 + (v38 & -2147483648))
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v38 += (v38 + 0)
                    }
                    _ = v38
                    _ = v39
                }
                var v60 uint = 2
                _ = fi35
                _ = v36
                _ = v37
                _ = v60
            }
            var v61 int = 799
            for fi62 := 0; fi62 < 8; fi62++ {
                var v63 int8 = 36
                var v64 int8 = v63
                _ = fi62
                _ = v63
                _ = v64
            }
            _ = fi34
            _ = v61
        }
        switch 880 {
        case 795:
            var v65 uint64 = 4294967295
            _ = v65
        case 1:
            v66 := GenS4{F0: 42, F1: v5, F2: "\n\t"}
            _ = v66
        }
        v67 := GenS1{F0: false, F1: 1, F2: (!((3.14159 - 3.14159) < (((3.14159 + 1e10) * 1.5) + (3.14159 * 1.5))))}
        _ = v67
    } else {
        v5 -= (^(v5))
    }
    var v68 bool = (false || false)
    if v68 {
        var v69 float32 = ((float32(1e10) / v6) + float32(-2.25))
        _ = v69
    } else {
        if ((v6 - float32((0.1 + 0.1))) == (v2 * 0.0)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            if (18446744073709551615 > 4) {
                switch "\n\t" {
                case "":
                    fmt.Println(v7)
                case "foo bar":
                    for fi70 := 0; fi70 < 17; fi70++ {
                        v71 := GenS5{F0: 1, F1: 2147483647, F2: -3}
                        _ = fi70
                        _ = v71
                    }
                default:
                    var v72 float32 = v2
                    _ = v72
                }
                var v73 float64 = 1.5
                _ = v73
            } else {
                var v74 float64 = 1e10
                if false {
                    fmt.Println(v7, v5, v5)
                    var v75 string = "hello"
                    v7 -= v7
                    _ = v75
                } else {
                    for fi76 := 0; fi76 < 12; fi76++ {
                        if v68 {
                            var v77 float32 = float32(v74)
                            for fi78 := 0; fi78 < 5; fi78++ {
                                if (v68 && v68) {
                                    v1[v7] = "hello"
                                } else {
                                    fmt.Println(fi76, v6)
                                    switch fi78 {
                                    case 0:
                                        v2 = 1e10
                                    case -9223372036854775808:
                                        v68 = ((42 == 746) && (42 < 127))
                                    }
                                    v79 := GenS2{F0: (^(v7)), F1: 2}
                                    var v80 bool = (!(!false))
                                    var v81 int32 = 179
                                    var v82 int8 = 1
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                                if false {
                                    v2 -= 3.14159
                                    if (true || true) {
                                        var v83 uint64 = 477
                                        if (!true) {
                                            var v84 string = "\n\t"
                                            if v68 {
                                                v83 -= v83
                                            } else {
                                                v68 = v68
                                                v83 = 2
                                                v7 = (v7 >> ((uint(v7)) & 63))
                                            }
                                            _ = v84
                                        } else {
                                            v77 -= v6
                                        }
                                        _ = v83
                                    }
                                    if true {
                                        switch "foo bar" {
                                        case "\n\t":
                                            if false {
                                                if true {
                                                    v4 = append(v4, 2147483647)
                                                    var v85 uint32 = v7
                                                    var v86 string = ""
                                                    _ = v85
                                                    _ = v86
                                                }
                                                var v87 int8 = 106
                                                if (int64(v87) >= -9223372036854775808) {
                                                    switch 1 {
                                                    case 9223372036854775807:
                                                        if (((v6 / -2.25) * ((0.1 - ((float32((v74 * v74)) * 0.0) - v6)) + 1.5)) != (0.1 * ((0.1 - 1.5) * -2.25))) {
                                                            var v88 int = (fi78 & fi76)
                                                            _ = v88
                                                        }
                                                    case 608:
                                                        if (!(false || false)) {
                                                            v5 -= (v5 & (v5 & v5))
                                                            var v89 float64 = v74
                                                            switch (-(fi78)) {
                                                            case 926:
                                                                v68 = (!((v87 >> ((2) & 63)) < (v87 | v87)))
                                                            case -873:
                                                                if (1 >= 1) {
                                                                    var v90 bool = (v5 < 75)
                                                                    v74 = 0.0
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v91 int = fi76
                                                                        _ = v91
                                                                    case "hello":
                                                                        v92 := GenS4{F0: 32767, F1: v5, F2: "hello"}
                                                                        _ = v92
                                                                    }
                                                                    _ = v90
                                                                } else {
                                                                    v93 := GenS1{F0: true, F1: v7, F2: v68}
                                                                    if v68 {
                                                                        switch ("\n\t" + "hello") {
                                                                        case "\n\t":
                                                                            var v94 int = 127
                                                                            _ = v94
                                                                        }
                                                                        fmt.Println(v5, fi76, v68)
                                                                        switch "hello" {
                                                                        case "a":
                                                                            fmt.Println(v87)
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v77, v68, v6)
                                                                        if (v77 != 1.5) {
                                                                            if v68 {
                                                                                v87 -= -1
                                                                                v5 = 1
                                                                            }
                                                                        } else {
                                                                            var v95 int32 = -482
                                                                            if true {
                                                                                if false {
                                                                                    var v96 float64 = v89
                                                                                    switch ((("foo bar" + "foo bar") + "") + "") {
                                                                                    case "foo bar":
                                                                                        var v97 uint16 = 2
                                                                                        _ = v97
                                                                                    case "\n\t":
                                                                                        v98 := GenS3{F0: true}
                                                                                        _ = v98
                                                                                    }
                                                                                    _ = v96
                                                                                }
                                                                                switch "a" {
                                                                                case "a":
                                                                                    v99 := GenS1{F0: ((uint(v7) >= 4294967295) && v68), F1: uint32(v5), F2: (!(-1 <= v95))}
                                                                                    _ = v99
                                                                                case "hello":
                                                                                    v100 := GenS6{F0: "a", F1: v7, F2: 255}
                                                                                    _ = v100
                                                                                case "foo bar":
                                                                                    v87 += v87
                                                                                default:
                                                                                    v3[""] = v68
                                                                                }
                                                                                v68 = (!(false || true))
                                                                            } else {
                                                                                var v101 int64 = 42
                                                                                _ = v101
                                                                            }
                                                                            _ = v95
                                                                        }
                                                                    }
                                                                    if true {
                                                                        var v102 int64 = 42
                                                                        _ = v102
                                                                    } else {
                                                                        var v103 uint8 = (v5 ^ 2)
                                                                        v103 = v103
                                                                        var v104 string = "hello"
                                                                        _ = v103
                                                                        _ = v104
                                                                    }
                                                                    v105 := GenS7{F0: 127, F1: "\n\t", F2: 18446744073709551615}
                                                                    v7 -= (v7 * v7)
                                                                    _ = v105
                                                                    _ = v93
                                                                }
                                                            }
                                                            _ = v89
                                                        }
                                                    case 127:
                                                        _ = len(v4)
                                                        var v106 uint64 = 42
                                                        _ = v106
                                                    }
                                                    v1[4294967295] = ""
                                                    switch (fi76 + (fi78 << ((255) & 63))) {
                                                    case 32768:
                                                        v77 = v6
                                                    default:
                                                        if v68 {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v5 -= v5
                                                            case "foo bar":
                                                                fmt.Println(v68)
                                                            case "hello":
                                                                switch ("a" + "hello") {
                                                                case "hello":
                                                                    if false {
                                                                        if (!(2147483647 >= v7)) {
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                var v107 uint64 = 472
                                                                                _ = v107
                                                                            }
                                                                        } else {
                                                                            v108 := GenS6{F0: "\n\t", F1: (v7 ^ (v7 | v7)), F2: 0}
                                                                            _ = v108
                                                                        }
                                                                    } else {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    v68 = (v87 >= v87)
                                                                case "":
                                                                    if true {
                                                                        if (!(!(!v68))) {
                                                                            if (("hello" + (("hello" + "") + "a")) != "hello") {
                                                                                v87 += v87
                                                                                v68 = (v74 == v74)
                                                                                switch (fi76 + fi78) {
                                                                                case -2147483648:
                                                                                    v109 := GenS8{F0: 1000}
                                                                                    v4 = append(v4, fi76)
                                                                                    _ = v109
                                                                                default:
                                                                                    fmt.Println(v6, v77, v7)
                                                                                }
                                                                            } else {
                                                                                var v110 int64 = 9223372036854775807
                                                                                _ = v110
                                                                            }
                                                                            v2 = 3.14159
                                                                        }
                                                                    }
                                                                    fmt.Println(v6, v5)
                                                                    v1[65535] = "\n\t"
                                                                }
                                                                if (!(!true)) {
                                                                    var v111 uint = 4294967295
                                                                    fmt.Println(v5, v111, v87)
                                                                    _ = v111
                                                                } else {
                                                                    var v112 uint8 = 42
                                                                    delete(v3, "")
                                                                    _ = v112
                                                                }
                                                                fmt.Println(v5)
                                                            }
                                                            var v113 uint16 = 2
                                                            var v114 uint8 = v5
                                                            _ = v113
                                                            _ = v114
                                                        } else {
                                                            if (v74 <= float64(-2.25)) {
                                                                v115 := GenS8{F0: -1}
                                                                _ = v115
                                                            }
                                                            v116 := GenS1{F0: ("a" == "a"), F1: v7, F2: v68}
                                                            var v117 int64 = 212
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                    }
                                                } else {
                                                    if v68 {
                                                        var v118 int32 = 0
                                                        switch fi78 {
                                                        case 32768:
                                                            fmt.Println(fi76, v77)
                                                        case 354:
                                                            v6 -= float32(1.5)
                                                        case 0:
                                                            v119 := GenS7{F0: (-(fi78)), F1: ("" + "foo bar"), F2: 485}
                                                            _ = v119
                                                        default:
                                                            var v120 int32 = (v118 % ((1) | 1))
                                                            _ = v120
                                                        }
                                                        var v121 uint64 = 2
                                                        switch ("" + "foo bar") {
                                                        case "hello":
                                                            fmt.Println(fi76)
                                                        case "foo bar":
                                                            v2 -= (v2 - (v77 * v2))
                                                        }
                                                        var v122 uint64 = uint64(v5)
                                                        v68 = v68
                                                        var v123 uint8 = v5
                                                        _ = v118
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                    var v124 int16 = 42
                                                    _ = len(v4)
                                                    v125 := GenS1{F0: true, F1: (v7 << ((uint(v5)) & 63)), F2: (0 <= fi76)}
                                                    _ = v124
                                                    _ = v125
                                                }
                                                _ = len(v3)
                                                fmt.Println(v68)
                                                var v126 uint32 = (v7 % ((2) | 1))
                                                _ = v126
                                                _ = v87
                                            }
                                        case "":
                                            v5 += v5
                                        default:
                                            var v127 uint32 = v7
                                            _ = v127
                                        }
                                    } else {
                                        var v128 float64 = v74
                                        v5 += 42
                                        _ = v128
                                    }
                                    v74 = v74
                                } else {
                                    v2 += (v77 * 0.1)
                                }
                                var v129 uint16 = 65535
                                v6 = v2
                                _ = fi78
                                _ = v129
                            }
                            _ = v77
                        } else {
                            var v130 int64 = -9223372036854775808
                            _ = v130
                        }
                        _ = fi76
                    }
                }
                var v131 string = ("" + (("foo bar" + "") + ("a" + "\n\t")))
                _ = v131
                _ = v74
            }
            delete(v1, (v7 & (v7 >> ((255) & 63))))
            if v68 {
                if (!(9223372036854775807 > -1)) {
                    if v68 {
                        var v132 int64 = -295
                        v68 = v68
                        if (65535 == 921) {
                            v68 = v68
                        } else {
                            for fi133 := 0; fi133 < 20; fi133++ {
                                var v134 bool = (true && (!v68))
                                var v135 bool = v134
                                if ((1e10 >= (0.0 * v2)) && (!false)) {
                                    if v135 {
                                        if (int8(fi133) >= 100) {
                                            for fi136 := 0; fi136 < 3; fi136++ {
                                                var v137 int16 = 0
                                                fmt.Println(v2, v132, v134)
                                                fmt.Println(fi133, v2)
                                                if (2147483647 <= -1) {
                                                    var v138 string = ("" + "foo bar")
                                                    var v139 float64 = 3.14159
                                                    _ = v138
                                                    _ = v139
                                                }
                                                _ = fi136
                                                _ = v137
                                            }
                                            if (0.1 <= v6) {
                                                for fi140 := 0; fi140 < 12; fi140++ {
                                                    delete(v1, (v7 % ((v7) | 1)))
                                                    v5 -= v5
                                                    _ = fi140
                                                }
                                                if v134 {
                                                    if true {
                                                        var v141 uint32 = (v7 * v7)
                                                        var v142 int = 133
                                                        _ = v141
                                                        _ = v142
                                                    } else {
                                                        var v143 bool = (!true)
                                                        var v144 uint32 = (v7 % ((1) | 1))
                                                        _ = v143
                                                        _ = v144
                                                    }
                                                }
                                                v68 = (!(v134 || ("" > (("hello" + "\n\t") + (("" + ("\n\t" + "")) + "foo bar")))))
                                            }
                                            switch -2147483648 {
                                            case -2147483648:
                                                v5 -= (v5 * v5)
                                            case 32767:
                                                if v68 {
                                                    var v145 int = fi133
                                                    var v146 uint32 = (^(v7))
                                                    switch "a" {
                                                    case "a":
                                                        var v147 uint32 = 1
                                                        _ = v147
                                                    case "hello":
                                                        if ((v68 && v68) && (0 > 1)) {
                                                            var v148 int = (v145 ^ -16)
                                                            _ = v148
                                                        }
                                                    }
                                                    _ = v145
                                                    _ = v146
                                                }
                                            default:
                                                v3[("\n\t" + "foo bar")] = (42 <= -52)
                                            }
                                        }
                                    }
                                    v149 := GenS9{F0: v5, F1: "hello"}
                                    fmt.Println(v68, fi133, fi133)
                                    _ = v149
                                } else {
                                    _ = len(v1)
                                    v150 := GenS9{F0: (^(v5)), F1: "\n\t"}
                                    _ = v150
                                }
                                _ = fi133
                                _ = v134
                                _ = v135
                            }
                            var v151 bool = ((!(42 != 2147483647)) && false)
                            _ = v151
                        }
                        _ = v132
                    } else {
                        var v152 int = -9223372036854775808
                        v7 += (^(v7))
                        _ = v152
                    }
                    if true {
                        v5 = (^(v5))
                        var v153 int64 = -1
                        v5 -= (v5 ^ (v5 * 255))
                        _ = v153
                    }
                }
            } else {
                v1[(v7 / (((^(v7))) | 1))] = (("foo bar" + ("" + "")) + ((("\n\t" + "\n\t") + "foo bar") + "\n\t"))
                var v154 int = 32767
                _ = v154
            }
        }
        for fi155 := 0; fi155 < 16; fi155++ {
            v156 := GenS8{F0: -362}
            v68 = true
            v4 = append(v4, (^(fi155)))
            _ = fi155
            _ = v156
        }
        v157 := GenS10{F0: (!(v6 <= 0.0)), F1: (("a" + ("a" + "hello")) + (("a" + (("a" + "") + "hello")) + "\n\t")), F2: (v7 * (^(v7)))}
        _ = v157
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v7
    _ = v9
}
