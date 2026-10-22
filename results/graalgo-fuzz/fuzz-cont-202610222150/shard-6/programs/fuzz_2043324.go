package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 string; F1 uint32; F2 string }
type GenS7 struct { F0 uint; F1 int8; F2 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16; F1 int; F2 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 int16; F1 uint; F2 int }
type GenS13 struct { F0 bool; F1 uint8; F2 int8 }

func main() {
    var v1 uint8 = 0
    var v2 int32 = 507
    var v3 uint32 = 579
    var v4 int64 = 746
    var v5 int = -18
    v4 -= v4
    if false {
        v6 := []int{-462, v5, v5}
        if len(v6) > 0 {
            v6[0] = v5
        }
        _ = v6
    }
    switch "\n\t" {
    case "\n\t":
        var v7 uint16 = 1
        _ = v7
    case "hello":
        if (1.5 > (-2.25 - float64((1e10 - 1e10)))) {
            switch "" {
            case "foo bar":
                if ((!((((("hello" + "foo bar") + ("foo bar" + "foo bar")) + ("a" + "hello")) != ((("a" + ("hello" + ("foo bar" + ""))) + "") + ("\n\t" + ""))) || false)) && (42 < 32767)) {
                    if false {
                        var v8 bool = true
                        _ = v8
                    } else {
                        if false {
                            v9 := []int{int(v4), 127, 9223372036854775807, v5}
                            for fi10 := 0; fi10 < 14; fi10++ {
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                switch v5 {
                                case -1:
                                    var v11 uint32 = (v3 + uint32(v1))
                                    _ = v11
                                default:
                                    v3 += (^(v3))
                                }
                                var v12 int = v5
                                _ = fi10
                                _ = v12
                            }
                            _ = v9
                        }
                    }
                    for fi13 := 0; fi13 < 2; fi13++ {
                        v4 += v4
                        _ = fi13
                    }
                } else {
                    var v14 float64 = 1e10
                    _ = v14
                }
            case "a":
                if false {
                    v2 -= v2
                } else {
                    if false {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            v3 = v3
                            for fi16 := 0; fi16 < 8; fi16++ {
                                switch "a" {
                                case "":
                                    var v17 uint32 = v3
                                    _ = v17
                                case "a":
                                    v1 = v1
                                case "\n\t":
                                    v18 := GenS1{F0: 0, F1: (-(v4)), F2: 4294967295}
                                    _ = v18
                                }
                                v1 = (^(v1))
                                _ = fi16
                            }
                            v19 := GenS2{F0: (v1 / (((v1 + (v1 % ((v1) | 1)))) | 1)), F1: (v1 << ((255) & 63)), F2: -800}
                            _ = fi15
                            _ = v19
                        }
                        switch "foo bar" {
                        case "":
                            v3 -= uint32(v1)
                        case "a":
                            var v20 float32 = 0.0
                            _ = v20
                        default:
                            for fi21 := 0; fi21 < 4; fi21++ {
                                if (-2.25 <= ((0.0 / (1e10 / (1.5 * (0.1 * float64((0.0 - ((((((0.0 / 1.0) * 1e10) + 1.5) + (1e10 / (1e10 / -2.25))) * ((3.14159 - -2.25) + 1e10)) + 3.14159))))))) / (0.1 - (0.1 + -2.25)))) {
                                    if (!false) {
                                        fmt.Println(v2)
                                        v22 := GenS1{F0: 224, F1: v4, F2: 0}
                                        _ = v22
                                    } else {
                                        v23 := []int{(v5 + (v5 % (((fi21 << ((2) & 63))) | 1))), int(v3), (^(v5))}
                                        var v24 bool = (0.1 <= (1e10 - float64(1.5)))
                                        _ = v23
                                        _ = v24
                                    }
                                    if (true || (((-183 <= -1) || false) || false)) {
                                        for fi25 := 0; fi25 < 10; fi25++ {
                                            v4 += (v4 / ((v4) | 1))
                                            if ((v2 ^ v2) < (-(v2))) {
                                                v26 := []int{(-(v5))}
                                                _ = v26
                                            }
                                            v27 := GenS2{F0: 2, F1: v1, F2: 1}
                                            _ = fi25
                                            _ = v27
                                        }
                                    } else {
                                        if (v2 == (^(v2))) {
                                            v28 := []int{fi21}
                                            fmt.Println(v5)
                                            _ = v28
                                        } else {
                                            switch fi21 {
                                            case 32767:
                                                if (-32768 > 1) {
                                                    v29 := GenS2{F0: v1, F1: v1, F2: 32767}
                                                    fmt.Println(v1, fi21, v4)
                                                    v5 = (v5 >> ((18446744073709551615) & 63))
                                                    _ = v29
                                                }
                                            }
                                            v1 = (v1 | v1)
                                        }
                                    }
                                } else {
                                    if true {
                                        var v30 int = 100
                                        for fi31 := 0; fi31 < 11; fi31++ {
                                            v32 := GenS1{F0: -924, F1: (-(v4)), F2: 42}
                                            _ = fi31
                                            _ = v32
                                        }
                                        var v33 uint8 = (v1 * v1)
                                        _ = v30
                                        _ = v33
                                    } else {
                                        v2 = (v2 - -2147483648)
                                    }
                                    var v34 float32 = (1e10 + -2.25)
                                    _ = v34
                                }
                                _ = fi21
                            }
                        }
                    }
                    v5 += int(v2)
                    if true {
                        fmt.Println(v5, v2, v2)
                        var v35 int8 = 42
                        _ = v35
                    }
                }
            case "\n\t":
                v5 = -887
            default:
                for fi36 := 0; fi36 < 8; fi36++ {
                    fmt.Println(v4, v4)
                    var v37 int16 = int16(v4)
                    _ = fi36
                    _ = v37
                }
            }
            if ((960 == 0) || false) {
                v1 += 2
                if (!true) {
                    for fi38 := 0; fi38 < 3; fi38++ {
                        var v39 uint32 = (v3 * 2147483648)
                        v40 := GenS3{F0: 137, F1: -529}
                        v41 := GenS2{F0: (v1 >> ((0) & 63)), F1: v1, F2: -218}
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                }
            } else {
                var v42 int64 = -9223372036854775808
                fmt.Println(v4)
                _ = v42
            }
        }
    case "a":
        if false {
            v2 = v2
        } else {
            v43 := GenS2{F0: (v1 << ((255) & 63)), F1: v1, F2: 305}
            switch ("\n\t" + "\n\t") {
            case "a":
                v3 = (v3 & uint32(v1))
            case "":
                for fi44 := 0; fi44 < 14; fi44++ {
                    v5 += fi44
                    _ = fi44
                }
            default:
                fmt.Println(v4, v2)
            }
            _ = v43
        }
    default:
        fmt.Println(v2, v3, v2)
    }
    var v45 int = 32767
    var v46 uint64 = 42
    for fi47 := 0; fi47 < 16; fi47++ {
        v4 = v4
        v4 += (v4 + (v4 * (v4 + v4)))
        fmt.Println(v5, v1, v4)
        _ = fi47
    }
    v48 := GenS1{F0: -32768, F1: 496, F2: v46}
    if (uint64(v1) >= 0) {
        fmt.Println(v1, v46)
        for fi49 := 0; fi49 < 12; fi49++ {
            if false {
                var v50 uint16 = 0
                v50 -= v50
                v51 := GenS3{F0: (v50 + 32767), F1: -345}
                _ = v50
                _ = v51
            }
            _ = fi49
        }
        fmt.Println(v3, v4)
    } else {
        for fi52 := 0; fi52 < 2; fi52++ {
            v1 = (v1 & v1)
            switch 32767 {
            case 32767:
                if (false && true) {
                    for fi53 := 0; fi53 < 4; fi53++ {
                        v54 := []int{(v5 * (^(fi53))), (-(v45)), v5, fi52}
                        _ = fi53
                        _ = v54
                    }
                    v2 += v2
                    switch "" {
                    case "":
                        if (true && false) {
                            v2 += (v2 / ((1) | 1))
                            var v55 int = (v5 * (fi52 - (fi52 << ((229) & 63))))
                            v56 := []int{}
                            _ = v55
                            _ = v56
                        }
                    case "foo bar":
                        var v57 int64 = 0
                        _ = v57
                    }
                } else {
                    for fi58 := 0; fi58 < 11; fi58++ {
                        v59 := []int{v45, 1}
                        _ = fi58
                        _ = v59
                    }
                    var v60 int64 = (v4 / ((v4) | 1))
                    fmt.Println(v1, v45)
                    _ = v60
                }
            }
            _ = fi52
        }
        var v61 uint8 = v1
        var v62 uint16 = uint16(v46)
        _ = v61
        _ = v62
    }
    if false {
        if true {
            if ("hello" >= "a") {
                for fi63 := 0; fi63 < 8; fi63++ {
                    for fi64 := 0; fi64 < 2; fi64++ {
                        fmt.Println(v2, v1)
                        _ = fi64
                    }
                    _ = fi63
                }
                var v65 int16 = 1000
                _ = v65
            } else {
                if ((0.1 >= (((3.14159 + 1.5) * 1.5) * 0.0)) && true) {
                    if false {
                        if (!false) {
                            var v66 float32 = 1.5
                            if false {
                                switch ("" + "foo bar") {
                                case "hello":
                                    v67 := []int{(v5 / ((v45) | 1)), (-(v45)), v5, v5}
                                    _ = v67
                                default:
                                    v68 := GenS2{F0: 42, F1: uint8(v3), F2: -1}
                                    _ = v68
                                }
                            }
                            _ = v66
                        } else {
                            v5 -= -1
                        }
                    } else {
                        v69 := GenS4{F0: v5, F1: v46}
                        switch v5 {
                        case 0:
                            v4 += v4
                        }
                        var v70 uint16 = 42
                        _ = v69
                        _ = v70
                    }
                    for fi71 := 0; fi71 < 8; fi71++ {
                        for fi72 := 0; fi72 < 14; fi72++ {
                            fmt.Println(v3)
                            if false {
                                if (!true) {
                                    var v73 uint32 = 63
                                    switch "foo bar" {
                                    case "\n\t":
                                        v3 = v73
                                    case "":
                                        fmt.Println(fi71)
                                    }
                                    var v74 uint32 = v73
                                    v75 := GenS2{F0: v1, F1: (v1 / (((^(v1))) | 1)), F2: 0}
                                    if false {
                                        var v76 int32 = (-(v2))
                                        v76 -= (v76 << ((18446744073709551615) & 63))
                                        _ = v76
                                    }
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                }
                            } else {
                                var v77 float64 = 1.5
                                v78 := GenS5{F0: 144, F1: -45}
                                _ = v77
                                _ = v78
                            }
                            _ = fi72
                        }
                        var v79 int64 = 9223372036854775807
                        if true {
                            v80 := GenS6{F0: "foo bar", F1: (v3 % ((uint32(v46)) | 1)), F2: ""}
                            _ = v80
                        }
                        _ = fi71
                        _ = v79
                    }
                    fmt.Println(v45)
                }
            }
            v45 -= (v45 - 32767)
            var v81 int64 = (v4 | 9223372036854775807)
            _ = v81
        }
        for fi82 := 0; fi82 < 6; fi82++ {
            if true {
                switch ("a" + (("foo bar" + "a") + "")) {
                case "a":
                    switch "\n\t" {
                    case "a":
                        var v83 uint8 = v1
                        _ = v83
                    case "foo bar":
                        v84 := GenS7{F0: 312, F1: -1, F2: v46}
                        _ = v84
                    case "\n\t":
                        if true {
                            v45 -= fi82
                        }
                    default:
                        v85 := GenS7{F0: 761, F1: 100, F2: (v46 + 874)}
                        _ = v85
                    }
                case "":
                    v1 = v1
                case "foo bar":
                    if ((v4 + (v4 - int64(v3))) >= v4) {
                        v86 := []int{(^(v45)), fi82, 127, (v45 ^ (v5 % ((v5) | 1)))}
                        if len(v86) > 0 {
                            v86[0] = v5
                        }
                        _ = v86
                    }
                }
                v87 := GenS1{F0: -32768, F1: (-(v4)), F2: v46}
                for fi88 := 0; fi88 < 14; fi88++ {
                    switch "hello" {
                    case "hello":
                        if true {
                            fmt.Println(v45, v4)
                        }
                    default:
                        if (false || true) {
                            v89 := GenS5{F0: 0, F1: 0}
                            _ = v89
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            fmt.Println(v5, fi82)
                            v2 = (v2 | (v2 - v2))
                            fmt.Println(v4, v5, v4)
                        } else {
                            v2 = (-(v2))
                            if true {
                                var v90 uint = 255
                                fmt.Println(v5, v45, v90)
                                v3 += (v3 / (((v3 | (v3 - (v3 >> ((v90) & 63))))) | 1))
                                _ = v90
                            }
                        }
                    case "hello":
                        if (0 <= 570) {
                            v91 := GenS2{F0: v1, F1: (v1 % (((^(v1))) | 1)), F2: 1}
                            _ = v91
                        } else {
                            if false {
                                var v92 float64 = 1e10
                                _ = v92
                            }
                            if ((1.5 + -2.25) <= (3.14159 * 1.5)) {
                                v46 = v46
                                if (v4 < v4) {
                                    v93 := GenS8{F0: (v46 << ((159) & 63))}
                                    _ = v93
                                } else {
                                    v2 = v2
                                    if true {
                                        v94 := GenS9{F0: 32767, F1: int(v2), F2: (fi82 << ((2) & 63))}
                                        if ((528 != 1) || (v1 <= (^(v1)))) {
                                            fmt.Println(v4)
                                        } else {
                                            if false {
                                                switch -9223372036854775808 {
                                                case 4294967295:
                                                    switch -2147483648 {
                                                    case 4294967295:
                                                        var v95 uint = 255
                                                        _ = v95
                                                    case 78:
                                                        v96 := []int{0, (fi82 / ((int(v1)) | 1)), v45, v5}
                                                        _ = v96
                                                    case -801:
                                                        if false {
                                                            v97 := []int{}
                                                            var v98 float64 = (float64(0.1) * (0.0 / 1.5))
                                                            if len(v97) > 0 {
                                                                v97[0] = fi88
                                                            }
                                                            _ = v97
                                                            _ = v98
                                                        } else {
                                                            v4 = v4
                                                            if (18446744073709551615 < (v46 + v46)) {
                                                                v99 := GenS10{F0: 32767}
                                                                var v100 int8 = 8
                                                                v101 := []int{}
                                                                v45 += 953
                                                                v102 := GenS9{F0: 1000, F1: (v45 ^ v45), F2: v5}
                                                                _ = v100
                                                                _ = v101
                                                                _ = v102
                                                                _ = v99
                                                            } else {
                                                                var v103 string = ((("hello" + ("\n\t" + "a")) + ("" + "a")) + (("hello" + ("hello" + "\n\t")) + "foo bar"))
                                                                var v104 int64 = v4
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                            v2 = 42
                                                            fmt.Println(v1, v1, v3)
                                                        }
                                                    }
                                                case 344:
                                                    if true {
                                                        v45 = fi88
                                                        switch v5 {
                                                        case 32768:
                                                            switch (((("a" + (("\n\t" + "a") + "foo bar")) + "foo bar") + "") + ("foo bar" + "")) {
                                                            case "foo bar":
                                                                var v105 uint8 = 1
                                                                _ = v105
                                                            default:
                                                                if true {
                                                                    fmt.Println(v45, v46)
                                                                    switch ("a" + ((("a" + ("hello" + "")) + "foo bar") + "hello")) {
                                                                    case "foo bar":
                                                                        v106 := []int{fi82, (fi88 + fi82), (fi88 & (v45 ^ v45))}
                                                                        _ = v106
                                                                    case "hello":
                                                                        if (((v3 + 42) != (v3 % ((v3) | 1))) && (v46 > v46)) {
                                                                            v46 = (v46 + v46)
                                                                            v107 := GenS10{F0: 0}
                                                                            _ = v107
                                                                        } else {
                                                                            v108 := []int{fi82, (fi88 | (^(fi88))), v45}
                                                                            v3 = (^(v3))
                                                                            _ = v108
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case 0:
                                                            var v109 string = ("\n\t" + "foo bar")
                                                            _ = v109
                                                        case -252:
                                                            v110 := GenS8{F0: (v46 % ((4294967295) | 1))}
                                                            _ = v110
                                                        default:
                                                            if (("hello" + "") >= "\n\t") {
                                                                var v111 int = fi82
                                                                _ = v111
                                                            } else {
                                                                switch "hello" {
                                                                case "":
                                                                    if (false && (((("" + "hello") + "") + "\n\t") == "hello")) {
                                                                        var v112 uint64 = (v46 ^ (v46 & (v46 % (((v46 & (v46 + 1))) | 1))))
                                                                        _ = v112
                                                                    }
                                                                case "a":
                                                                    if true {
                                                                        v46 += v46
                                                                        var v113 bool = (936 <= 1000)
                                                                        _ = v113
                                                                    }
                                                                default:
                                                                    v5 = (v45 >> ((0) & 63))
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        if true {
                                                            v114 := []int{v5, (^(fi88)), v5, (v5 * v45)}
                                                            v115 := GenS3{F0: 1, F1: -32768}
                                                            _ = v114
                                                            _ = v115
                                                        } else {
                                                            v116 := GenS11{F0: ("\n\t" + ("\n\t" + "foo bar")), F1: 758}
                                                            var v117 string = ("\n\t" + ("foo bar" + "\n\t"))
                                                            var v118 uint16 = 255
                                                            _ = v116
                                                            _ = v117
                                                            _ = v118
                                                        }
                                                        switch ("" + "\n\t") {
                                                        case "hello":
                                                            if (!false) {
                                                                if (!((2147483648 == v3) || false)) {
                                                                    fmt.Println(v3)
                                                                    v119 := GenS3{F0: 255, F1: -32768}
                                                                    _ = v119
                                                                }
                                                                fmt.Println(v1)
                                                            } else {
                                                                v120 := []int{(^(fi82))}
                                                                _ = v120
                                                            }
                                                        case "foo bar":
                                                            v121 := GenS10{F0: 720}
                                                            _ = v121
                                                        default:
                                                            if true {
                                                                v122 := GenS4{F0: (-(fi88)), F1: 1}
                                                                var v123 string = "foo bar"
                                                                switch (^(fi88)) {
                                                                case 295:
                                                                    if false {
                                                                        if (0 >= 65535) {
                                                                            var v124 int8 = 0
                                                                            var v125 int16 = 1
                                                                            var v126 uint8 = v1
                                                                            var v127 float64 = 0.0
                                                                            if false {
                                                                                fmt.Println(v124, v5, v4)
                                                                                var v128 int64 = int64(v124)
                                                                                fmt.Println(v3)
                                                                                _ = v128
                                                                            }
                                                                            _ = v124
                                                                            _ = v125
                                                                            _ = v126
                                                                            _ = v127
                                                                        }
                                                                        if true {
                                                                            var v129 float32 = 0.0
                                                                            v130 := GenS12{F0: 32767, F1: 73, F2: fi88}
                                                                            v5 -= fi82
                                                                            _ = v129
                                                                            _ = v130
                                                                        }
                                                                        fmt.Println(v46)
                                                                    } else {
                                                                        v123 = "a"
                                                                    }
                                                                case 4294967295:
                                                                    fmt.Println(fi82)
                                                                case 127:
                                                                    var v131 uint32 = 441
                                                                    _ = v131
                                                                }
                                                                _ = v122
                                                                _ = v123
                                                            }
                                                        }
                                                    }
                                                }
                                                v3 += (v3 | 351)
                                                if true {
                                                    var v132 float32 = (-2.25 + float32(1.5))
                                                    _ = v132
                                                } else {
                                                    v1 -= v1
                                                    fmt.Println(v45, v4)
                                                    v1 = (v1 | 128)
                                                }
                                            }
                                            var v133 int32 = 0
                                            _ = v133
                                        }
                                        _ = v94
                                    } else {
                                        v45 = v5
                                    }
                                }
                                v3 = v3
                            }
                            fmt.Println(v4, v4)
                        }
                    }
                    v134 := []int{}
                    _ = fi88
                    _ = v134
                }
                _ = v87
            } else {
                switch 4294967295 {
                case 211:
                    if true {
                        v135 := GenS6{F0: ("foo bar" + "a"), F1: (v3 ^ (v3 | (v3 + v3))), F2: (("\n\t" + (("foo bar" + (((("hello" + "a") + ("" + "a")) + "foo bar") + "hello")) + "hello")) + "hello")}
                        v2 += v2
                        _ = v135
                    }
                case -213:
                    v136 := GenS13{F0: true, F1: (v1 % ((v1) | 1)), F2: 42}
                    _ = v136
                case -1:
                    var v137 int8 = -128
                    _ = v137
                default:
                    fmt.Println(v1)
                }
            }
            _ = fi82
        }
    } else {
        var v138 bool = (v4 > (v4 % (((v4 | v4)) | 1)))
        _ = v138
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v5
}
