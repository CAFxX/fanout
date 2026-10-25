package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 string }
type GenS2 struct { F0 uint; F1 uint16; F2 uint }
type GenS3 struct { F0 int8; F1 uint64; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint32 = 177
    var v2 uint32 = (v1 ^ v1)
    var v3 int64 = 0
    v1 = v2
    var v4 int8 = 108
    var v5 float32 = 1.5
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            var v7 uint = 484
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v3, v3, v3)
    v2 -= (v1 * v1)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint64 = 18446744073709551615
        var v10 int16 = -32768
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(v3, v2)
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(v2, fi11, fi12)
            v13 := []int{fi12, 127, 32768}
            v14 := GenS1{F0: fi12, F1: 1, F2: (("hello" + "\n\t") + "")}
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(fi11)
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS1{F0: fi11, F1: 217, F2: "foo bar"}
            var v17 uint16 = 1
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        if true {
            if false {
                v19 := GenS2{F0: 4294967295, F1: 2, F2: 255}
                var v20 uint32 = (v1 - (v2 % ((v1) | 1)))
                var v21 string = "foo bar"
                _ = v19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                if true {
                    if (!false) {
                        if (true || true) {
                            switch "\n\t" {
                            case "\n\t":
                                v5 = v5
                                v1 = (v2 % ((65535) | 1))
                                if ((false && false) || true) {
                                    fmt.Println(fi22, v2)
                                    v23 := GenS1{F0: (-(fi18)), F1: 28, F2: "foo bar"}
                                    if false {
                                        if true {
                                            fmt.Println(v3, v4)
                                        } else {
                                            var v24 uint32 = v1
                                            if (1e10 == 1.5) {
                                                v25 := []int{(fi22 + (fi18 >> ((67) & 63))), int(v4)}
                                                var v26 int = 827
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v27 := []int{(^(fi22))}
                                                _ = v27
                                            }
                                            if false {
                                                v28 := []int{32768}
                                                var v29 uint16 = 65535
                                                _ = v28
                                                _ = v29
                                            }
                                            _ = v24
                                        }
                                    } else {
                                        if (!true) {
                                            if true {
                                                var v30 float64 = (3.14159 - 0.1)
                                                v31 := []int{-9223372036854775808, fi22}
                                                _ = v30
                                                _ = v31
                                            }
                                        } else {
                                            var v32 uint = 4294967295
                                            v33 := []int{}
                                            switch fi18 {
                                            case 4294967295:
                                                fmt.Println(v3)
                                            default:
                                                if len(v33) > 0 {
                                                    v33[0] = 32767
                                                }
                                            }
                                            var v34 uint64 = 42
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        }
                                        if (!((false || (v5 > v5)) && (723 < 65535))) {
                                            fmt.Println(fi22, v3, v3)
                                            var v35 uint64 = 1
                                            v36 := GenS2{F0: 2, F1: 255, F2: 65535}
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    _ = v23
                                }
                                var v37 float32 = ((v5 / v5) + v5)
                                _ = v37
                            case "foo bar":
                                var v38 uint32 = 1
                                _ = v38
                            case "":
                                var v39 bool = (fi18 > fi18)
                                _ = v39
                            default:
                                v40 := GenS1{F0: fi22, F1: 127, F2: "foo bar"}
                                _ = v40
                            }
                            fmt.Println(v1)
                        } else {
                            switch "" {
                            case "a":
                                v41 := GenS2{F0: 18446744073709551615, F1: 1, F2: 42}
                                _ = v41
                            }
                            v42 := []int{fi22, 518, fi22}
                            if true {
                                var v43 uint16 = 255
                                var v44 uint16 = 639
                                _ = v43
                                _ = v44
                            } else {
                                switch (fi22 * (fi22 >> ((18446744073709551615) & 63))) {
                                case 32767:
                                    if true {
                                        var v45 int8 = (v4 >> ((951) & 63))
                                        _ = v45
                                    }
                                case -192:
                                    var v46 int64 = -1
                                    _ = v46
                                case 42:
                                    v47 := GenS1{F0: -1, F1: 128, F2: "hello"}
                                    _ = v47
                                }
                            }
                            _ = v42
                        }
                    } else {
                        v48 := []int{fi18, fi22, 268, fi22}
                        v5 = (v5 / 1e10)
                        var v49 float32 = v5
                        _ = v48
                        _ = v49
                    }
                    v50 := GenS2{F0: 18446744073709551615, F1: 42, F2: 0}
                    _ = v50
                }
                var v51 int16 = 1000
                _ = fi22
                _ = v51
            }
            for fi52 := 0; fi52 < 9; fi52++ {
                switch ("hello" + "") {
                case "\n\t":
                    var v53 float32 = 1.5
                    _ = v53
                default:
                    switch (fi18 / (((-(fi52))) | 1)) {
                    case 127:
                        if (0 == 255) {
                            v5 = float32(3.14159)
                            v3 = v3
                            var v54 int16 = -32768
                            _ = v54
                        } else {
                            var v55 uint64 = 42
                            fmt.Println(v3, v2, v1)
                            _ = v55
                        }
                    default:
                        v1 = v1
                    }
                }
                if (-1 != 42) {
                    fmt.Println(fi52, fi52)
                    v56 := GenS3{F0: v4, F1: 1, F2: fi18}
                    v3 = -9223372036854775808
                    _ = v56
                }
                fmt.Println(fi18, fi52, v2)
                _ = fi52
            }
        }
        v57 := []int{-1}
        fmt.Println(v3)
        _ = fi18
        _ = v57
    }
    for fi58 := 0; fi58 < 7; fi58++ {
        fmt.Println(v4, v5, v5)
        fmt.Println(v1)
        if true {
            if true {
                for fi59 := 0; fi59 < 5; fi59++ {
                    v60 := GenS1{F0: fi58, F1: 128, F2: (("\n\t" + "\n\t") + ("" + "hello"))}
                    _ = fi59
                    _ = v60
                }
                switch int(v3) {
                case 127:
                    switch fi58 {
                    case 1:
                        var v61 int64 = -148
                        _ = v61
                    case -111:
                        switch "\n\t" {
                        case "hello":
                            v1 -= v2
                        default:
                            for fi62 := 0; fi62 < 11; fi62++ {
                                v63 := []int{1, 128, 32768}
                                v64 := GenS2{F0: 255, F1: 65535, F2: uint(v1)}
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                        }
                    case 2147483647:
                        v65 := GenS4{F0: (v2 >> ((18446744073709551615) & 63))}
                        _ = v65
                    }
                case 4294967295:
                    if (!(("foo bar" < "hello") || false)) {
                        for fi66 := 0; fi66 < 18; fi66++ {
                            v67 := []int{fi66, 128}
                            _ = fi66
                            _ = v67
                        }
                        v3 -= (v3 ^ (v3 / ((0) | 1)))
                    }
                case 2147483647:
                    var v68 uint = 18446744073709551615
                    _ = v68
                default:
                    switch "" {
                    case "\n\t":
                        if true {
                            if false {
                                v3 = (v3 % ((v3) | 1))
                                if (true && ((255 <= 0) || (v3 > (v3 + v3)))) {
                                    v69 := GenS2{F0: 42, F1: 65535, F2: 18446744073709551615}
                                    _ = v69
                                } else {
                                    if (4294967295 >= v2) {
                                        v70 := []int{fi58, (fi58 & 455)}
                                        _ = v70
                                    }
                                    if true {
                                        switch (fi58 / (((fi58 ^ -9223372036854775808)) | 1)) {
                                        case -1:
                                            for fi71 := 0; fi71 < 6; fi71++ {
                                                v72 := []int{}
                                                _ = fi71
                                                _ = v72
                                            }
                                        case -2147483648:
                                            var v73 float32 = (v5 + (v5 * float32(-2.25)))
                                            _ = v73
                                        case 416:
                                            var v74 float64 = -2.25
                                            _ = v74
                                        }
                                        var v75 int64 = (v3 * -706)
                                        for fi76 := 0; fi76 < 3; fi76++ {
                                            v77 := GenS1{F0: fi76, F1: 127, F2: ("\n\t" + (("foo bar" + "foo bar") + ""))}
                                            _ = fi76
                                            _ = v77
                                        }
                                        _ = v75
                                    } else {
                                        switch (fi58 & -9223372036854775808) {
                                        case 4294967295:
                                            for fi78 := 0; fi78 < 9; fi78++ {
                                                if true {
                                                    v79 := []int{32767, (fi58 + int(v2)), fi58}
                                                    var v80 int32 = 0
                                                    v5 -= (v5 / 0.1)
                                                    _ = v79
                                                    _ = v80
                                                }
                                                v81 := GenS5{F0: 42, F1: -572}
                                                var v82 string = "hello"
                                                switch 0 {
                                                case 127:
                                                    v1 = v2
                                                case 0:
                                                    var v83 uint = 911
                                                    _ = v83
                                                case 42:
                                                    v84 := GenS2{F0: 280, F1: 499, F2: 1}
                                                    _ = v84
                                                default:
                                                    v85 := []int{(fi58 >> ((255) & 63))}
                                                    _ = v85
                                                }
                                                v86 := []int{int(v4)}
                                                _ = fi78
                                                _ = v81
                                                _ = v82
                                                _ = v86
                                            }
                                        }
                                    }
                                    for fi87 := 0; fi87 < 17; fi87++ {
                                        var v88 uint16 = 2
                                        v4 = v4
                                        _ = fi87
                                        _ = v88
                                    }
                                }
                                for fi89 := 0; fi89 < 7; fi89++ {
                                    v90 := []int{(fi58 | fi89), (-(fi58)), -9223372036854775808, 976}
                                    v5 += (v5 + v5)
                                    v90 = append(v90, (fi89 & (fi89 >> ((2) & 63))))
                                    _ = fi89
                                    _ = v90
                                }
                            }
                            var v91 int16 = 1
                            for fi92 := 0; fi92 < 9; fi92++ {
                                v93 := GenS6{F0: 956}
                                v5 += 3.14159
                                _ = fi92
                                _ = v93
                            }
                            _ = v91
                        }
                    case "a":
                        switch ("" + "foo bar") {
                        case "":
                            v3 = -701
                        case "foo bar":
                            if false {
                                var v94 float32 = v5
                                if false {
                                    var v95 uint8 = 255
                                    var v96 uint = 1
                                    v97 := GenS5{F0: 1000, F1: -1}
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                }
                                v3 = (v3 + (v3 % ((v3) | 1)))
                                _ = v94
                            } else {
                                v4 -= -54
                                for fi98 := 0; fi98 < 6; fi98++ {
                                    fmt.Println(v1, v5)
                                    switch (("" + "\n\t") + "foo bar") {
                                    case "\n\t":
                                        switch fi98 {
                                        case 244:
                                            v99 := []int{fi98}
                                            _ = v99
                                        }
                                    case "":
                                        v1 += 0
                                    }
                                    if true {
                                        if true {
                                            if false {
                                                if (!false) {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if ((1 >= 2) || (v3 >= (v3 * -9223372036854775808))) {
                                                            v100 := []int{0, fi98, (fi98 / (((-(fi98))) | 1)), (fi58 - fi58)}
                                                            v1 = v2
                                                            if true {
                                                                var v101 float32 = v5
                                                                var v102 int = (fi98 % (((fi98 + (fi98 - 1))) | 1))
                                                                v103 := GenS6{F0: 35}
                                                                switch int(v1) {
                                                                case 127:
                                                                    var v104 float64 = (((0.0 * ((0.0 / (1e10 / 3.14159)) + (-2.25 * 0.1))) / 1e10) / float64(((v5 / (v101 / (float32((((-2.25 + 1.5) - -2.25) - 1e10)) / 1e10))) * v5)))
                                                                    _ = v104
                                                                }
                                                                v4 = 95
                                                                var v105 uint16 = 800
                                                                fmt.Println(v5)
                                                                v105 = (v105 >> ((51) & 63))
                                                                _ = v101
                                                                _ = v102
                                                                _ = v103
                                                                _ = v105
                                                            }
                                                            _ = v100
                                                        }
                                                    default:
                                                        var v106 uint32 = v2
                                                        _ = v106
                                                    }
                                                    v107 := []int{fi98, (-(fi58))}
                                                    if true {
                                                        v2 += v1
                                                    }
                                                    switch (fi58 + (fi58 / (((fi58 | fi98)) | 1))) {
                                                    case 127:
                                                        if (1e10 == (float64(3.14159) * 1.5)) {
                                                            v5 -= v5
                                                            v107 = append(v107, fi58)
                                                            switch ("\n\t" + "\n\t") {
                                                            case "":
                                                                var v108 int = fi98
                                                                _ = v108
                                                            case "hello":
                                                                if true {
                                                                    if (!false) {
                                                                        v4 += 66
                                                                        if true {
                                                                            v3 = (v3 << ((656) & 63))
                                                                        }
                                                                    }
                                                                } else {
                                                                    v3 = v3
                                                                    fmt.Println(fi98, v4)
                                                                    _ = len(v107)
                                                                }
                                                            case "foo bar":
                                                                if (1 == 1) {
                                                                    fmt.Println(v3, v1, fi58)
                                                                    if false {
                                                                        v3 -= (v3 ^ v3)
                                                                    } else {
                                                                        switch (fi58 + fi98) {
                                                                        case -2147483648:
                                                                            var v109 uint64 = 0
                                                                            _ = v109
                                                                        case 127:
                                                                            v107 = append(v107, 42)
                                                                        }
                                                                    }
                                                                    v1 -= v2
                                                                } else {
                                                                    fmt.Println(v3, v3, v2)
                                                                }
                                                            }
                                                        } else {
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                _ = len(v107)
                                                            default:
                                                                v3 -= v3
                                                            }
                                                            var v110 string = ("foo bar" + "hello")
                                                            _ = v110
                                                        }
                                                    case 1:
                                                        var v111 uint16 = 255
                                                        _ = v111
                                                    case 4294967295:
                                                        switch (fi98 / ((-478) | 1)) {
                                                        case -1:
                                                            v2 = (v1 >> ((2) & 63))
                                                        }
                                                    }
                                                    _ = v107
                                                }
                                                v3 += (v3 % ((v3) | 1))
                                                v112 := GenS5{F0: -1, F1: -2147483648}
                                                _ = v112
                                            }
                                        } else {
                                            v3 -= v3
                                            v113 := []int{(-(fi98)), fi58, fi58}
                                            if true {
                                                v114 := GenS4{F0: v2}
                                                var v115 int8 = -128
                                                _ = v114
                                                _ = v115
                                            }
                                            _ = v113
                                        }
                                        var v116 int32 = 1
                                        _ = v116
                                    } else {
                                        v4 = (-(v4))
                                        fmt.Println(v1, v2)
                                    }
                                    _ = fi98
                                }
                            }
                        default:
                            v117 := GenS5{F0: -32768, F1: 1}
                            _ = v117
                        }
                    case "foo bar":
                        var v118 uint32 = 2
                        _ = v118
                    default:
                        var v119 uint16 = 869
                        _ = v119
                    }
                }
            }
        } else {
            if true {
                var v120 uint64 = 42
                if ("\n\t" == ("a" + "foo bar")) {
                    v5 = float32((((0.1 / 0.1) * (1.5 / (0.0 - 1e10))) + 1e10))
                }
                if true {
                    switch "" {
                    case "foo bar":
                        v121 := []int{32767, (-(fi58)), fi58, fi58}
                        _ = v121
                    }
                    v4 = (v4 ^ (v4 / ((53) | 1)))
                } else {
                    var v122 uint16 = 867
                    if true {
                        var v123 float32 = -2.25
                        v124 := GenS6{F0: 0}
                        _ = v123
                        _ = v124
                    } else {
                        v122 = (v122 & 255)
                    }
                    _ = v122
                }
                _ = v120
            }
        }
        _ = fi58
    }
    v125 := GenS3{F0: v4, F1: 971, F2: 0}
    var v126 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v125
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
