package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint32; F1 uint64 }
type GenS10 struct { F0 int16; F1 int64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string; F1 uint64; F2 int64 }
type GenS14 struct { F0 uint8; F1 uint }
type GenS15 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS16 struct { F0 uint; F1 int }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 string; F1 int32; F2 uint32 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint64; F1 int16 }
type GenS21 struct { F0 string; F1 int32 }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 int; F1 int8 }
type GenS24 struct { F0 int8; F1 int32; F2 uint8 }
type GenS25 struct { F0 string; F1 int64; F2 uint }
type GenS26 struct { F0 int64; F1 uint16 }
type GenS27 struct { F0 int64; F1 uint8 }
type GenS28 struct { F0 uint; F1 int16; F2 int16 }
type GenS29 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS30 struct { F0 uint8 }
type GenS31 struct { F0 uint8 }
type GenS32 struct { F0 string; F1 int16; F2 int16 }
type GenS33 struct { F0 uint64; F1 bool }
type GenS34 struct { F0 uint32 }
type GenS35 struct { F0 uint64; F1 int16 }
type GenS36 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]string)
    var v3 int16 = -298
    v4 := make(map[string]string)
    if true {
        v5 := GenS1{F0: 65535, F1: 42}
        _ = v5
    } else {
        v6 := GenS2{F0: 255, F1: v1, F2: -1}
        var v7 float64 = ((-2.25 / float64((-2.25 - 1.5))) + float64((float32((float64(1e10) * 1.5)) - (((1e10 * (float32(-2.25) - (1.5 - -2.25))) / 0.1) + 1.5))))
        switch -785 {
        case 9223372036854775807:
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 int = (^(fi8))
                fmt.Println(fi8)
                _ = fi8
                _ = v9
            }
        case -496:
            var v10 int16 = (v3 >> ((42) & 63))
            _ = v10
        case 127:
            v11 := GenS3{F0: 0}
            _ = v11
        }
        _ = v6
        _ = v7
    }
    switch 127 {
    case 32767:
        for fi12 := 0; fi12 < 14; fi12++ {
            switch "hello" {
            case "hello":
                var v13 uint32 = (v1 >> ((65535) & 63))
                _ = v13
            }
            var v14 string = ("\n\t" + "\n\t")
            _ = fi12
            _ = v14
        }
    case 883:
        v4[(("hello" + "foo bar") + "a")] = ""
    default:
        if false {
            v3 = v3
            v3 -= v3
            var v15 uint = 18446744073709551615
            _ = v15
        }
    }
    if (false || (!false)) {
        var v16 int32 = 42
        fmt.Println(v16, v16)
        if true {
            v16 = (v16 % (((v16 - v16)) | 1))
            if ((42 != 0) && false) {
                if true {
                    var v17 uint8 = 1
                    _ = v17
                } else {
                    v18 := GenS4{F0: uint16(v1), F1: uint(v1), F2: (v1 % ((v1) | 1))}
                    _ = v18
                }
                fmt.Println(v16)
                if true {
                    v19 := GenS5{F0: "", F1: 200}
                    var v20 int32 = 1
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v3, v16)
                }
            } else {
                v1 = 1
                v3 -= (^(v3))
                for fi21 := 0; fi21 < 19; fi21++ {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        v23 := []int{fi21, fi22, (fi22 & fi22)}
                        var v24 int16 = (v3 * v3)
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v3 -= v3
                    _ = fi21
                }
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 uint32 = v1
                _ = fi25
                _ = v26
            }
        } else {
            var v27 uint = 1
            _ = v27
        }
        _ = v16
    }
    v3 += 0
    var v28 float64 = (1e10 + -2.25)
    v28 -= (0.0 / v28)
    if true {
        for fi29 := 0; fi29 < 14; fi29++ {
            v30 := v2[v1]
            v31 := GenS1{F0: 2, F1: 0}
            _ = fi29
            _ = v30
            _ = v31
        }
        if (372 <= 926) {
            if false {
                v28 = v28
                var v32 uint = 18446744073709551615
                for fi33 := 0; fi33 < 12; fi33++ {
                    v34 := []int{(fi33 * (fi33 * (fi33 | (-(fi33)))))}
                    for fi35 := 0; fi35 < 8; fi35++ {
                        if false {
                            v36 := GenS3{F0: 1}
                            _ = len(v4)
                            _ = v36
                        } else {
                            var v37 int32 = -1
                            v28 = v28
                            _ = v37
                        }
                        _ = fi35
                    }
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
        } else {
            var v38 float32 = -2.25
            var v39 int32 = -792
            v3 += v3
            _ = v38
            _ = v39
        }
        v1 = v1
    } else {
        var v40 float64 = ((v28 / v28) * 0.1)
        delete(v2, v1)
        delete(v2, (v1 - 0))
        _ = v40
    }
    if (!(!(!(v28 <= (v28 * -2.25))))) {
        var v41 uint32 = (v1 - (v1 % (((^(v1))) | 1)))
        _ = v41
    } else {
        v3 = 0
    }
    switch -2147483648 {
    case -413:
        if ((-2.25 + 1e10) == 1.5) {
            v28 = (v28 + (v28 * (v28 + 1.5)))
            for fi42 := 0; fi42 < 12; fi42++ {
                for fi43 := 0; fi43 < 4; fi43++ {
                    var v44 uint64 = 2
                    v45 := GenS6{F0: 42}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                _ = fi42
            }
            var v46 float32 = ((float32((0.1 + (v28 - v28))) - (0.1 - 3.14159)) - 1.5)
            _ = v46
        } else {
            v1 += (v1 | (v1 | (v1 - v1)))
        }
    case 0:
        for fi47 := 0; fi47 < 11; fi47++ {
            for fi48 := 0; fi48 < 2; fi48++ {
                var v49 float32 = (0.1 + 1e10)
                v28 += (v28 / ((3.14159 / v28) * v28))
                _ = fi48
                _ = v49
            }
            v50 := GenS7{F0: 900}
            v3 -= v3
            _ = fi47
            _ = v50
        }
    }
    if false {
        delete(v2, v1)
    }
    var v51 uint16 = 65535
    v1 = (v1 ^ v1)
    v52 := []int{32767, -1, 32767, 42}
    if false {
        for fi53 := 0; fi53 < 4; fi53++ {
            var v54 float64 = (1e10 / ((0.0 * v28) - (3.14159 / ((1.5 / (0.1 - v28)) * ((v28 * -2.25) * 0.1)))))
            v55 := GenS5{F0: "a", F1: 42}
            _ = fi53
            _ = v54
            _ = v55
        }
    }
    var v56 int32 = 1
    for fi57 := 0; fi57 < 11; fi57++ {
        if true {
            for fi58 := 0; fi58 < 19; fi58++ {
                v59 := GenS2{F0: 0, F1: v1, F2: (-(v56))}
                _ = fi58
                _ = v59
            }
            v60 := v2[v1]
            _ = v60
        }
        switch ("hello" + "\n\t") {
        case "foo bar":
            if len(v52) > 0 {
                v52[0] = (fi57 << ((456) & 63))
            }
        case "":
            _ = len(v2)
        }
        if false {
            switch fi57 {
            case 4294967295:
                v61 := GenS8{F0: -32768}
                _ = v61
            case 1:
                for fi62 := 0; fi62 < 2; fi62++ {
                    if false {
                        var v63 uint8 = 42
                        var v64 string = "a"
                        v63 = v63
                        _ = v63
                        _ = v64
                    }
                    v65 := GenS9{F0: (v1 + (v1 - (v1 + v1))), F1: 1}
                    _ = fi62
                    _ = v65
                }
            case 0:
                if (!(v1 >= v1)) {
                    if (-1 <= -81) {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi57, v28, v3)
                        case "foo bar":
                            if (42 <= 18446744073709551615) {
                                fmt.Println(v56, v56)
                                for fi66 := 0; fi66 < 18; fi66++ {
                                    var v67 uint16 = v51
                                    if true {
                                        _ = len(v52)
                                        v68 := GenS8{F0: (v3 * v3)}
                                        var v69 int8 = -46
                                        _ = v68
                                        _ = v69
                                    }
                                    _ = fi66
                                    _ = v67
                                }
                                var v70 string = (("" + "foo bar") + "hello")
                                _ = v70
                            } else {
                                for fi71 := 0; fi71 < 9; fi71++ {
                                    if (237 > (v1 * v1)) {
                                        delete(v2, v1)
                                        var v72 uint8 = 0
                                        var v73 uint = 894
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = fi71
                                }
                                var v74 int16 = -105
                                switch "\n\t" {
                                case "hello":
                                    var v75 uint8 = 127
                                    _ = v75
                                }
                                _ = v74
                            }
                        case "hello":
                            fmt.Println(fi57)
                        default:
                            fmt.Println(v56, v56, v28)
                        }
                        if true {
                            var v76 uint16 = (v51 + (^(v51)))
                            _ = v76
                        } else {
                            var v77 uint64 = 664
                            _ = v77
                        }
                        switch "\n\t" {
                        case "\n\t":
                            if ((71 <= 127) && ((fi57 << ((65535) & 63)) != fi57)) {
                                var v78 int32 = (^(v56))
                                _ = v78
                            }
                        default:
                            var v79 uint8 = 54
                            _ = v79
                        }
                    }
                    var v80 int = int(v3)
                    v81 := GenS8{F0: 1000}
                    _ = v80
                    _ = v81
                } else {
                    v82 := GenS10{F0: 1, F1: 839}
                    _ = v82
                }
            default:
                if true {
                    v1 -= 42
                }
            }
        } else {
            fmt.Println(v1, v51, v56)
            for fi83 := 0; fi83 < 18; fi83++ {
                var v84 int = int(v3)
                v56 = 42
                _ = fi83
                _ = v84
            }
            if (100 < -128) {
                if false {
                    v52 = append(v52, (fi57 | (fi57 % ((9223372036854775807) | 1))))
                } else {
                    for fi85 := 0; fi85 < 17; fi85++ {
                        v56 -= v56
                        _ = fi85
                    }
                    switch fi57 {
                    case 2147483647:
                        if ((false && true) && false) {
                            if (-486 >= (fi57 | (fi57 * -1))) {
                                v28 += v28
                                var v86 int64 = 42
                                if ((v3 > 0) && false) {
                                    switch ("hello" + "foo bar") {
                                    case "\n\t":
                                        v28 = 0.0
                                    default:
                                        for fi87 := 0; fi87 < 13; fi87++ {
                                            if ((!(("a" + "") > ("foo bar" + "a"))) || (v86 <= (v86 ^ v86))) {
                                                switch (("hello" + ("a" + "")) + "") {
                                                case "\n\t":
                                                    if len(v52) >= 2 {
                                                        v52 = v52[1:len(v52)]
                                                    }
                                                case "a":
                                                    if len(v52) > 0 {
                                                        v52[0] = 361
                                                    }
                                                }
                                                if ((957 <= 313) && true) {
                                                    var v88 uint16 = 0
                                                    v28 -= (((v28 * (((float64(float32(0.1)) + float64(-2.25)) * float64(1e10)) - float64(-2.25))) + (float64(-2.25) - 1.5)) - 0.1)
                                                    _ = v88
                                                } else {
                                                    switch ("\n\t" + "foo bar") {
                                                    case "hello":
                                                        v1 -= v1
                                                    case "":
                                                        v89 := v2[(v1 & v1)]
                                                        v90 := GenS11{F0: 100}
                                                        _ = v89
                                                        _ = v90
                                                    case "foo bar":
                                                        switch ("" + ("a" + "\n\t")) {
                                                        case "":
                                                            v56 += (v56 >> ((4294967295) & 63))
                                                        case "hello":
                                                            fmt.Println(v86)
                                                        default:
                                                            if (!(-2.25 >= float64(((((float32(0.1) + 0.0) - 1e10) - -2.25) * 0.1)))) {
                                                                v86 += (v86 ^ 1)
                                                                var v91 uint = 1
                                                                v28 += (((3.14159 - v28) + float64((-2.25 * (0.1 * 3.14159)))) + 0.0)
                                                                _ = v91
                                                            } else {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    fmt.Println(v86)
                                                                }
                                                                v3 = v3
                                                            }
                                                        }
                                                        var v92 int32 = v56
                                                        _ = v92
                                                    }
                                                }
                                            }
                                            if len(v52) >= 2 {
                                                v52 = v52[1:len(v52)]
                                            }
                                            delete(v2, v1)
                                            if (188 >= 0) {
                                                if (false || true) {
                                                    switch "foo bar" {
                                                    case "a":
                                                        if (119 <= -1) {
                                                            if (1e10 >= ((-2.25 - 0.1) / 1.0)) {
                                                                if (0 <= v51) {
                                                                    if true {
                                                                        v93 := GenS6{F0: 127}
                                                                        v94 := GenS7{F0: (v86 / ((v86) | 1))}
                                                                        _ = v93
                                                                        _ = v94
                                                                    }
                                                                    v3 -= v3
                                                                    var v95 float64 = v28
                                                                    _ = v95
                                                                } else {
                                                                    if true {
                                                                        v28 += 1.5
                                                                        if (v51 <= (v51 & (v51 / ((v51) | 1)))) {
                                                                            v96 := GenS12{F0: 127}
                                                                            _ = v96
                                                                        } else {
                                                                            var v97 float32 = 1e10
                                                                            fmt.Println(v28)
                                                                            _ = v97
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(v3)
                                                                        } else {
                                                                            if false {
                                                                                if false {
                                                                                    v51 += v51
                                                                                    v98 := GenS13{F0: "hello", F1: 1, F2: (v86 + v86)}
                                                                                    _ = v98
                                                                                } else {
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        v56 -= 2147483647
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            v99 := GenS10{F0: -266, F1: v86}
                                                                                            if len(v52) >= 2 {
                                                                                                v52 = v52[1:len(v52)]
                                                                                            }
                                                                                            var v100 float64 = ((float64(1.5) / float64(1.5)) / v28)
                                                                                            _ = v100
                                                                                            _ = v99
                                                                                        }
                                                                                    case "":
                                                                                        var v101 int16 = v3
                                                                                        _ = v101
                                                                                    default:
                                                                                        switch ("" + "") {
                                                                                        case "\n\t":
                                                                                            v56 -= 889
                                                                                        default:
                                                                                            var v102 uint16 = 32767
                                                                                            _ = v102
                                                                                        }
                                                                                    }
                                                                                    v4["hello"] = "\n\t"
                                                                                    _ = len(v52)
                                                                                }
                                                                                if (v1 <= v1) {
                                                                                    if ((!true) && ((v3 & v3) > v3)) {
                                                                                        v103 := GenS14{F0: 42, F1: 915}
                                                                                        _ = v103
                                                                                    }
                                                                                    v104 := GenS15{F0: (v1 ^ v1), F1: v1, F2: v86}
                                                                                    _ = v104
                                                                                }
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    switch ("" + "hello") {
                                                                                    case "hello":
                                                                                        v105 := GenS3{F0: v51}
                                                                                        _ = v105
                                                                                    }
                                                                                case "hello":
                                                                                    var v106 uint64 = 705
                                                                                    _ = v106
                                                                                default:
                                                                                    if (!((2 > 4294967295) || (true || ((!true) && false)))) {
                                                                                        var v107 string = "\n\t"
                                                                                        _ = v107
                                                                                    } else {
                                                                                        var v108 uint8 = 145
                                                                                        v109 := GenS5{F0: "foo bar", F1: (v108 >> ((0) & 63))}
                                                                                        switch 2147483647 {
                                                                                        case 32768:
                                                                                            var v110 float32 = float32(v28)
                                                                                            _ = v110
                                                                                        case 582:
                                                                                            var v111 int64 = (v86 * -1)
                                                                                            _ = v111
                                                                                        }
                                                                                        _ = v108
                                                                                        _ = v109
                                                                                    }
                                                                                }
                                                                                switch int(v3) {
                                                                                case 1:
                                                                                    v112 := v4[""]
                                                                                    _ = v112
                                                                                case 42:
                                                                                    var v113 int = (fi87 & (fi87 | -1))
                                                                                    _ = v113
                                                                                default:
                                                                                    v1 = 2147483648
                                                                                }
                                                                                fmt.Println(v51)
                                                                            } else {
                                                                                var v114 uint16 = (v51 / ((42) | 1))
                                                                                _ = v114
                                                                            }
                                                                            if false {
                                                                                v1 += v1
                                                                                if len(v52) >= 2 {
                                                                                    v52 = v52[1:len(v52)]
                                                                                }
                                                                                switch (fi57 * (fi57 * (fi87 << ((0) & 63)))) {
                                                                                case 0:
                                                                                    var v115 int32 = v56
                                                                                    _ = v115
                                                                                case 1:
                                                                                    v3 -= -304
                                                                                case -1:
                                                                                    v28 = v28
                                                                                default:
                                                                                    var v116 uint16 = (v51 - v51)
                                                                                    _ = v116
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    fmt.Println(v3, v3, v51)
                                                                                case "\n\t":
                                                                                    v117 := GenS16{F0: 65535, F1: (^(fi87))}
                                                                                    _ = v117
                                                                                case "a":
                                                                                    var v118 uint8 = 1
                                                                                    _ = v118
                                                                                }
                                                                            }
                                                                        }
                                                                        if true {
                                                                            fmt.Println(fi57, v1, v86)
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v86 -= (v86 & 0)
                                                                        v119 := GenS17{F0: v51}
                                                                        var v120 int = 1
                                                                        _ = v119
                                                                        _ = v120
                                                                    }
                                                                    fmt.Println(fi87)
                                                                    switch (fi57 / ((fi57) | 1)) {
                                                                    case 843:
                                                                        v51 -= v51
                                                                    case 42:
                                                                        var v121 uint = 4294967295
                                                                        _ = v121
                                                                    default:
                                                                        switch fi57 {
                                                                        case 128:
                                                                            v122 := GenS10{F0: (v3 >> ((1) & 63)), F1: v86}
                                                                            _ = v122
                                                                        }
                                                                        if false {
                                                                            if (1 == int8(v51)) {
                                                                                v56 = (v56 ^ (v56 << ((0) & 63)))
                                                                                var v123 int16 = v3
                                                                                if (!(v51 > (v51 - 748))) {
                                                                                    v124 := v2[2147483648]
                                                                                    v86 += 1
                                                                                    fmt.Println(v124, v56, v28)
                                                                                    _ = v124
                                                                                }
                                                                                _ = v123
                                                                            } else {
                                                                                v1 -= v1
                                                                                v125 := v2[v1]
                                                                                switch v125 {
                                                                                case "":
                                                                                    var v126 string = ("foo bar" + v125)
                                                                                    _ = v126
                                                                                case "foo bar":
                                                                                    var v127 bool = false
                                                                                    _ = v127
                                                                                default:
                                                                                    v86 = v86
                                                                                }
                                                                                var v128 uint64 = 42
                                                                                _ = v125
                                                                                _ = v128
                                                                            }
                                                                        }
                                                                        fmt.Println(v51)
                                                                    }
                                                                    v129 := GenS15{F0: 2, F1: v1, F2: v86}
                                                                    _ = v129
                                                                }
                                                                v1 -= 924
                                                                fmt.Println(v28, v1, v1)
                                                                if true {
                                                                    if false {
                                                                        v86 = int64(v1)
                                                                    }
                                                                } else {
                                                                    switch fi87 {
                                                                    case 32768:
                                                                        switch (("a" + ("hello" + "a")) + ((("foo bar" + ("hello" + "a")) + "hello") + ("foo bar" + ""))) {
                                                                        case "foo bar":
                                                                            v130 := GenS18{F0: "hello", F1: (v56 ^ (v56 << ((508) & 63))), F2: v1}
                                                                            _ = v130
                                                                        case "\n\t":
                                                                            if true {
                                                                                var v131 uint16 = (v51 / ((v51) | 1))
                                                                                if len(v52) >= 2 {
                                                                                    v52 = v52[1:len(v52)]
                                                                                }
                                                                                fmt.Println(v131, v1, fi57)
                                                                                _ = v131
                                                                            } else {
                                                                                v132 := GenS17{F0: (v51 | (v51 % ((2) | 1)))}
                                                                                v133 := GenS12{F0: 49}
                                                                                _ = v132
                                                                                _ = v133
                                                                            }
                                                                        default:
                                                                            v3 -= (^(v3))
                                                                        }
                                                                    case 658:
                                                                        var v134 int = fi57
                                                                        _ = v134
                                                                    default:
                                                                        v4["\n\t"] = ("\n\t" + "\n\t")
                                                                    }
                                                                    v135 := GenS19{F0: v51}
                                                                    v52 = append(v52, (-(fi57)))
                                                                    _ = v135
                                                                }
                                                                v56 = v56
                                                            } else {
                                                                var v136 uint8 = 42
                                                                _ = v136
                                                            }
                                                        } else {
                                                            var v137 uint64 = 433
                                                            fmt.Println(v3, fi57, v28)
                                                            _ = v137
                                                        }
                                                    case "hello":
                                                        var v138 bool = (int(v56) < (fi57 * fi57))
                                                        _ = v138
                                                    }
                                                    v86 -= (v86 * (v86 - v86))
                                                } else {
                                                    v139 := GenS4{F0: (^(v51)), F1: 18446744073709551615, F2: (v1 >> ((236) & 63))}
                                                    var v140 string = "a"
                                                    if len(v52) >= 2 {
                                                        v52 = v52[1:len(v52)]
                                                    }
                                                    var v141 uint32 = v1
                                                    if len(v52) > 0 {
                                                        v52[0] = (fi87 % ((fi87) | 1))
                                                    }
                                                    v142 := GenS20{F0: 1, F1: v3}
                                                    _ = v139
                                                    _ = v140
                                                    _ = v141
                                                    _ = v142
                                                }
                                            }
                                            _ = fi87
                                        }
                                    }
                                    for fi143 := 0; fi143 < 18; fi143++ {
                                        v144 := GenS1{F0: 255, F1: 255}
                                        var v145 int64 = (v86 >> ((42) & 63))
                                        _ = fi143
                                        _ = v144
                                        _ = v145
                                    }
                                }
                                _ = v86
                            }
                        }
                    }
                    if true {
                        v146 := GenS21{F0: ("hello" + "hello"), F1: v56}
                        v28 += (((float64((3.14159 / (3.14159 + (0.0 + 1e10)))) / (0.1 / (float64((1.5 + ((((0.1 * 1.5) - -2.25) * 0.0) - 1e10))) - float64((3.14159 - 1.5))))) / (((((1e10 + -2.25) * v28) - (((v28 - v28) - v28) - 3.14159)) / v28) - float64(0.1))) / float64(1e10))
                        for fi147 := 0; fi147 < 10; fi147++ {
                            v2[v1] = "a"
                            switch (fi147 + fi147) {
                            case 42:
                                var v148 int64 = -1
                                _ = v148
                            }
                            if true {
                                v56 = v56
                                var v149 uint8 = 0
                                if true {
                                    fmt.Println(v56)
                                    if true {
                                        delete(v2, 410)
                                        var v150 int8 = 127
                                        if len(v52) > 0 {
                                            v52[0] = 368
                                        }
                                        _ = v150
                                    }
                                } else {
                                    v151 := GenS22{F0: 772}
                                    var v152 uint = 565
                                    v56 = v56
                                    _ = v151
                                    _ = v152
                                }
                                switch fi57 {
                                case 128:
                                    switch "" {
                                    case "":
                                        if (v56 == -1) {
                                            if true {
                                                var v153 string = ("foo bar" + (("hello" + "hello") + ("a" + ("hello" + "hello"))))
                                                switch fi147 {
                                                case 2147483647:
                                                    if true {
                                                        v51 -= v51
                                                        var v154 uint16 = v51
                                                        _ = v154
                                                    }
                                                case 128:
                                                    fmt.Println(v153, fi147, v1)
                                                case 32767:
                                                    fmt.Println(v51)
                                                }
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v3, v28)
                                                }
                                                _ = v153
                                            }
                                            v28 -= v28
                                        } else {
                                            var v155 uint = uint(v1)
                                            v156 := GenS23{F0: (fi147 & 353), F1: 42}
                                            fmt.Println(v149, v1, v56)
                                            _ = v155
                                            _ = v156
                                        }
                                        var v157 int = fi57
                                        if true {
                                            if false {
                                                v4[("\n\t" + ("a" + ("foo bar" + ("" + (("foo bar" + ("hello" + (("foo bar" + "hello") + ("a" + "\n\t")))) + (("a" + (("" + "\n\t") + "a")) + "hello"))))))] = "hello"
                                                var v158 int = (fi147 & int(v56))
                                                _ = v158
                                            } else {
                                                if (false && true) {
                                                    v159 := GenS24{F0: 100, F1: 329, F2: (v149 / ((v149) | 1))}
                                                    var v160 bool = true
                                                    _ = v159
                                                    _ = v160
                                                }
                                                v157 += (v157 * fi147)
                                                v161 := GenS25{F0: "a", F1: -1, F2: 4294967295}
                                                fmt.Println(v56, v3)
                                                _ = v161
                                            }
                                            fmt.Println(v3, v1)
                                        }
                                        _ = v157
                                    default:
                                        if len(v52) >= 2 {
                                            v52 = v52[1:len(v52)]
                                        }
                                    }
                                case -823:
                                    fmt.Println(fi57, v56)
                                default:
                                    v162 := GenS26{F0: 0, F1: v51}
                                    _ = v162
                                }
                                _ = v149
                            } else {
                                var v163 string = "a"
                                _ = v163
                            }
                            fmt.Println(v3, v3, v3)
                            if true {
                                v164 := GenS27{F0: -1, F1: 255}
                                v51 -= v51
                                _ = v164
                            } else {
                                if (-1 != 42) {
                                    switch 756 {
                                    case 42:
                                        v56 = v56
                                    }
                                    var v165 uint64 = 18446744073709551615
                                    var v166 int16 = 849
                                    _ = v165
                                    _ = v166
                                }
                                var v167 bool = (false && true)
                                var v168 uint16 = 255
                                _ = v167
                                _ = v168
                            }
                            _ = fi147
                        }
                        _ = v146
                    } else {
                        switch fi57 {
                        case -605:
                            if false {
                                if (3.14159 > 1.5) {
                                    v56 = (v56 + v56)
                                    v1 -= (v1 / (((^(v1))) | 1))
                                    v1 += v1
                                } else {
                                    if ((!true) && false) {
                                        switch (fi57 - (fi57 % (((fi57 << ((164) & 63))) | 1))) {
                                        case 32768:
                                            if true {
                                                _ = len(v2)
                                                v169 := GenS28{F0: 13, F1: v3, F2: int16(v51)}
                                                if true {
                                                    if true {
                                                        v1 += (v1 << ((65535) & 63))
                                                    } else {
                                                        var v170 int16 = (-(v3))
                                                        for fi171 := 0; fi171 < 18; fi171++ {
                                                            fmt.Println(fi57)
                                                            var v172 float32 = 0.0
                                                            switch 4294967295 {
                                                            case -2147483648:
                                                                v173 := GenS7{F0: 873}
                                                                _ = v173
                                                            case 128:
                                                                v56 = v56
                                                            case 108:
                                                                var v174 int32 = (v56 * (v56 % (((v56 | 2147483647)) | 1)))
                                                                _ = v174
                                                            }
                                                            _ = fi171
                                                            _ = v172
                                                        }
                                                        var v175 int16 = (v3 + (^(v170)))
                                                        _ = v170
                                                        _ = v175
                                                    }
                                                    fmt.Println(v56)
                                                } else {
                                                    v176 := GenS29{F0: v1, F1: 128, F2: (v51 / ((v51) | 1))}
                                                    _ = v176
                                                }
                                                _ = v169
                                            }
                                        case -9223372036854775808:
                                            var v177 float64 = 1.5
                                            _ = v177
                                        case 128:
                                            for fi178 := 0; fi178 < 1; fi178++ {
                                                switch ("" + "foo bar") {
                                                case "foo bar":
                                                    var v179 uint = 783
                                                    _ = v179
                                                default:
                                                    var v180 int8 = -81
                                                    _ = v180
                                                }
                                                var v181 int64 = -191
                                                var v182 string = "foo bar"
                                                _ = fi178
                                                _ = v181
                                                _ = v182
                                            }
                                        }
                                    } else {
                                        switch 32768 {
                                        case 42:
                                            if (((fi57 ^ fi57) == (fi57 - 801)) && (100 < 100)) {
                                                fmt.Println(v1, v28)
                                                for fi183 := 0; fi183 < 10; fi183++ {
                                                    v1 = 334
                                                    if (((^(v51)) == 42) || false) {
                                                        var v184 int8 = 42
                                                        switch "hello" {
                                                        case "a":
                                                            if true {
                                                                var v185 int16 = (v3 % ((v3) | 1))
                                                                var v186 float64 = 3.14159
                                                                _ = v185
                                                                _ = v186
                                                            } else {
                                                                if len(v52) >= 2 {
                                                                    v52 = v52[1:len(v52)]
                                                                }
                                                                var v187 uint16 = v51
                                                                _ = v187
                                                            }
                                                        case "hello":
                                                            var v188 uint32 = v1
                                                            _ = v188
                                                        case "\n\t":
                                                            delete(v2, v1)
                                                        default:
                                                            v28 = v28
                                                        }
                                                        _ = v184
                                                    } else {
                                                        v51 += (v51 >> ((160) & 63))
                                                        var v189 uint = 979
                                                        _ = v189
                                                    }
                                                    _ = len(v4)
                                                    var v190 float32 = 0.0
                                                    if (v190 >= (v190 / v190)) {
                                                        if true {
                                                            if true {
                                                                _ = len(v4)
                                                                v191 := v4[("foo bar" + "")]
                                                                _ = v191
                                                            }
                                                            switch "\n\t" {
                                                            case "":
                                                                v190 += float32(((((v28 - ((v28 - 1.5) - v28)) - -2.25) * v28) + v28))
                                                            }
                                                            v28 += 1.5
                                                            var v192 int16 = v3
                                                            switch "a" {
                                                            case "\n\t":
                                                                if true {
                                                                    v51 -= 648
                                                                    switch ("\n\t" + "a") {
                                                                    case "foo bar":
                                                                        if true {
                                                                            var v193 int64 = 9223372036854775807
                                                                            _ = v193
                                                                        }
                                                                    case "\n\t":
                                                                        var v194 uint64 = 4294967295
                                                                        _ = v194
                                                                    case "hello":
                                                                        var v195 string = "foo bar"
                                                                        _ = v195
                                                                    default:
                                                                        if true {
                                                                            var v196 int = -1
                                                                            _ = v196
                                                                        }
                                                                    }
                                                                } else {
                                                                    v51 = (^(v51))
                                                                    v192 = v3
                                                                    v197 := v4["a"]
                                                                    v198 := GenS17{F0: (^(v51))}
                                                                    _ = v197
                                                                    _ = v198
                                                                }
                                                            case "a":
                                                                var v199 int = 1
                                                                _ = v199
                                                            case "":
                                                                var v200 float64 = v28
                                                                _ = v200
                                                            default:
                                                                var v201 int32 = v56
                                                                _ = v201
                                                            }
                                                            if false {
                                                                v202 := GenS30{F0: 2}
                                                                _ = v202
                                                            } else {
                                                                var v203 uint16 = v51
                                                                v2[v1] = "a"
                                                                fmt.Println(v28, fi57)
                                                                _ = v203
                                                            }
                                                            _ = v192
                                                        } else {
                                                            v2[v1] = (("" + "") + "\n\t")
                                                            v204 := GenS31{F0: 124}
                                                            if true {
                                                                if (4294967295 != 1) {
                                                                    v205 := GenS12{F0: 161}
                                                                    var v206 int64 = 200
                                                                    _ = v205
                                                                    _ = v206
                                                                }
                                                                var v207 uint32 = v1
                                                                _ = v207
                                                            } else {
                                                                v208 := GenS20{F0: 582, F1: (v3 >> ((621) & 63))}
                                                                switch 32768 {
                                                                case 127:
                                                                    fmt.Println(fi183, v56, fi183)
                                                                default:
                                                                    _ = len(v2)
                                                                }
                                                                switch (fi57 % ((fi183) | 1)) {
                                                                case -9223372036854775808:
                                                                    var v209 int8 = -128
                                                                    _ = v209
                                                                case 42:
                                                                    delete(v4, (("" + (("a" + "foo bar") + "hello")) + ("foo bar" + "hello")))
                                                                case 32768:
                                                                    if len(v52) > 0 {
                                                                        v52[0] = 486
                                                                    }
                                                                }
                                                                _ = v208
                                                            }
                                                            switch int(v51) {
                                                            case 0:
                                                                v28 += (((((0.1 + 1e10) - 1.5) - 3.14159) * float64(((v190 / 1.0) / ((3.14159 + v190) * ((v190 * v190) / ((v190 + (v190 + float32(1e10))) * (float32(v28) + v190))))))) - v28)
                                                            case 32767:
                                                                fmt.Println(v51)
                                                            case 9223372036854775807:
                                                                if true {
                                                                    var v210 float32 = 0.0
                                                                    _ = v210
                                                                }
                                                            default:
                                                                fmt.Println(fi183, v190)
                                                            }
                                                            _ = v204
                                                        }
                                                        var v211 int = fi57
                                                        if (!((v1 << ((18446744073709551615) & 63)) > v1)) {
                                                            var v212 bool = (true && (!(255 >= 181)))
                                                            v190 += v190
                                                            _ = v212
                                                        }
                                                        v190 = 1e10
                                                        if (true && false) {
                                                            _ = len(v4)
                                                            var v213 uint64 = 18446744073709551615
                                                            fmt.Println(v51, v56, v1)
                                                            _ = v213
                                                        }
                                                        _ = v211
                                                    }
                                                    _ = fi183
                                                    _ = v190
                                                }
                                                for fi214 := 0; fi214 < 14; fi214++ {
                                                    if len(v52) > 0 {
                                                        v52[0] = fi57
                                                    }
                                                    _ = fi214
                                                }
                                            }
                                        default:
                                            if true {
                                                v51 += v51
                                                if true {
                                                    var v215 int = fi57
                                                    _ = v215
                                                } else {
                                                    var v216 int = (^(fi57))
                                                    switch "" {
                                                    case "":
                                                        var v217 int32 = 434
                                                        _ = v217
                                                    case "hello":
                                                        if (!(0 >= v216)) {
                                                            for fi218 := 0; fi218 < 18; fi218++ {
                                                                v219 := GenS32{F0: "", F1: 0, F2: v3}
                                                                v220 := GenS33{F0: 272, F1: true}
                                                                _ = fi218
                                                                _ = v219
                                                                _ = v220
                                                            }
                                                            v221 := GenS34{F0: (v1 >> ((4294967295) & 63))}
                                                            _ = v221
                                                        } else {
                                                            fmt.Println(v1, v216, fi57)
                                                            var v222 uint = 0
                                                            v51 += 42
                                                            _ = v222
                                                        }
                                                    }
                                                    var v223 uint32 = 126
                                                    _ = v216
                                                    _ = v223
                                                }
                                            }
                                        }
                                        _ = len(v2)
                                        v224 := GenS20{F0: 607, F1: v3}
                                        _ = v224
                                    }
                                    v3 -= (v3 + v3)
                                    for fi225 := 0; fi225 < 10; fi225++ {
                                        v226 := GenS26{F0: -9223372036854775808, F1: v51}
                                        var v227 bool = true
                                        v28 = 1e10
                                        _ = fi225
                                        _ = v226
                                        _ = v227
                                    }
                                }
                                v51 = (v51 ^ (v51 >> ((1) & 63)))
                                fmt.Println(v56, v28)
                            }
                        case -9223372036854775808:
                            for fi228 := 0; fi228 < 16; fi228++ {
                                fmt.Println(v1, v1)
                                if true {
                                    fmt.Println(v28)
                                    v229 := GenS28{F0: 534, F1: (v3 - (^(v3))), F2: (v3 / ((v3) | 1))}
                                    _ = v229
                                } else {
                                    var v230 uint16 = v51
                                    if false {
                                        fmt.Println(v1, fi228)
                                        v52 = append(v52, -354)
                                        v3 = v3
                                    } else {
                                        if ((^(v56)) == (v56 / ((-1) | 1))) {
                                            v231 := GenS34{F0: 2147483647}
                                            _ = v231
                                        } else {
                                            fmt.Println(v230, fi57, fi57)
                                            if true {
                                                var v232 uint16 = (v51 / ((v230) | 1))
                                                delete(v2, 2147483647)
                                                v51 = (v230 + v51)
                                                var v233 uint = 4294967295
                                                if (true || true) {
                                                    fmt.Println(v230)
                                                    var v234 float64 = v28
                                                    _ = v234
                                                } else {
                                                    v235 := GenS35{F0: 225, F1: 42}
                                                    if true {
                                                        v236 := GenS30{F0: 78}
                                                        v237 := GenS36{F0: (fi228 << (((v233 ^ v233)) & 63)), F1: v1}
                                                        _ = v236
                                                        _ = v237
                                                    }
                                                    _ = v235
                                                }
                                                _ = v232
                                                _ = v233
                                            } else {
                                                fmt.Println(fi57, v1, fi228)
                                                v238 := v2[v1]
                                                var v239 int8 = 100
                                                _ = v238
                                                _ = v239
                                            }
                                            v28 += v28
                                        }
                                        fmt.Println(v51)
                                        if len(v52) > 0 {
                                            v52[0] = fi57
                                        }
                                    }
                                    _ = v230
                                }
                                _ = fi228
                            }
                        }
                        fmt.Println(fi57, fi57, fi57)
                    }
                }
                if (!false) {
                    fmt.Println(v28, v1)
                } else {
                    fmt.Println(v51, fi57)
                    fmt.Println(v3)
                    var v240 uint16 = v51
                    _ = v240
                }
            } else {
                if (18446744073709551615 != 4294967295) {
                    delete(v2, (v1 & (v1 % (((v1 << ((18446744073709551615) & 63))) | 1))))
                    v241 := GenS22{F0: 18446744073709551615}
                    _ = v241
                }
                v3 += 512
            }
        }
        _ = fi57
    }
    v242 := GenS25{F0: "\n\t", F1: 9223372036854775807, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v242
    _ = v28
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v56
}
