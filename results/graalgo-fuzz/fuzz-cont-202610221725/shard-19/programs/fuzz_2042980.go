package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8; F1 uint32; F2 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int8; F2 int64 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 string; F1 string; F2 int32 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 int32; F1 int64 }
type GenS13 struct { F0 uint8; F1 string }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int64; F1 int32; F2 int8 }
type GenS20 struct { F0 uint; F1 int; F2 int32 }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            switch fi5 {
            case 190:
                v7 := []int{(fi6 & fi6), (fi6 ^ fi6), (fi6 * 42)}
                _ = v7
            case 1:
                var v8 uint32 = 65535
                _ = v8
            case -92:
                v9 := GenS1{F0: 1000, F1: true}
                _ = v9
            default:
                var v10 int64 = -1
                _ = v10
            }
            v11 := []int{fi5, fi5}
            _ = fi6
            _ = v11
        }
        var v12 int8 = 42
        _ = fi5
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 float32 = 1e10
        v14 = v14
        fmt.Println(v14)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        switch (^(fi15)) {
        case 2147483647:
            v16 := GenS1{F0: -32768, F1: false}
            _ = v16
        case -9223372036854775808:
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 float32 = -2.25
                _ = fi17
                _ = v18
            }
        case 128:
            var v19 uint64 = 886
            _ = v19
        default:
            v20 := []int{fi15}
            _ = v20
        }
        _ = fi15
    }
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 2
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    for fi21 := 0; fi21 < 15; fi21++ {
        switch (fi21 | fi21) {
        case 1:
            var v22 int16 = -159
            _ = v22
        case 9223372036854775807:
            switch "hello" {
            case "foo bar":
                v23 := GenS2{F0: 42}
                _ = v23
            default:
                if true {
                    v1[2] = 2147483648
                    switch fi21 {
                    case 127:
                        var v24 int32 = 1
                        _ = v24
                    case 4294967295:
                        var v25 uint8 = v2
                        _ = v25
                    case -9223372036854775808:
                        if true {
                            switch ("foo bar" + "foo bar") {
                            case "foo bar":
                                var v26 uint = 70
                                _ = v26
                            case "a":
                                var v27 int16 = 42
                                _ = v27
                            default:
                                var v28 uint64 = 42
                                _ = v28
                            }
                        }
                    }
                } else {
                    if (4294967295 != 175) {
                        if true {
                            v2 = v2
                            if true {
                                var v29 float64 = ((1.5 / 3.14159) + -2.25)
                                v30 := GenS3{F0: 2, F1: uint(v2)}
                                if (false && false) {
                                    v29 = 0.0
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        var v32 int = 1
                                        delete(v1, 65535)
                                        _ = fi31
                                        _ = v32
                                    }
                                } else {
                                    for fi33 := 0; fi33 < 6; fi33++ {
                                        if (true || ("foo bar" > "foo bar")) {
                                            if (true && false) {
                                                var v34 uint64 = 768
                                                var v35 int = fi21
                                                v36 := GenS1{F0: 32767, F1: (!(65535 >= 161))}
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                            }
                                            v29 = 1e10
                                            v29 += (0.1 * float64(((1e10 * (float32(v29) + 1.5)) / (-2.25 / 1.0))))
                                        } else {
                                            v37 := []int{(^(fi21)), (fi33 / ((fi21) | 1))}
                                            v2 += 0
                                            fmt.Println(fi33)
                                            _ = v37
                                        }
                                        if (42 < 9223372036854775807) {
                                            v2 -= (^(v2))
                                            fmt.Println(v29)
                                            fmt.Println(v2)
                                        }
                                        if true {
                                            v2 -= v2
                                            v38 := []int{-1, 0}
                                            if true {
                                                var v39 uint = 4294967295
                                                _ = v39
                                            }
                                            _ = v38
                                        } else {
                                            var v40 int8 = 0
                                            v41 := GenS4{F0: 100, F1: 836, F2: 255}
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = fi33
                                    }
                                    fmt.Println(fi21, fi21)
                                }
                                _ = v29
                                _ = v30
                            } else {
                                for fi42 := 0; fi42 < 18; fi42++ {
                                    if true {
                                        v43 := []int{fi42, (fi42 % ((fi42) | 1)), fi21, 42}
                                        var v44 uint = 255
                                        _ = v43
                                        _ = v44
                                    } else {
                                        var v45 int64 = -579
                                        var v46 uint8 = 114
                                        var v47 uint = 255
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                    fmt.Println(fi42, fi42)
                                    v2 -= (^(v2))
                                    switch fi42 {
                                    case 254:
                                        if ((0 != 18446744073709551615) && false) {
                                            var v48 int32 = int32(v2)
                                            v2 += 2
                                            fmt.Println(fi42)
                                            _ = v48
                                        }
                                    case 42:
                                        v49 := []int{(fi42 | (fi21 & (fi42 - (fi21 - (fi42 + fi42))))), (fi42 >> ((65535) & 63))}
                                        _ = v49
                                    case 32768:
                                        var v50 string = "hello"
                                        v51 := GenS5{F0: -524}
                                        _ = v50
                                        _ = v51
                                    }
                                    if (-758 <= -2147483648) {
                                        v2 = 0
                                        switch ("hello" + "foo bar") {
                                        case "a":
                                            var v52 string = ("\n\t" + "hello")
                                            _ = v52
                                        case "hello":
                                            var v53 float32 = 3.14159
                                            _ = v53
                                        }
                                        switch (-(fi42)) {
                                        case 1:
                                            fmt.Println(v2, v2, v2)
                                        case 9223372036854775807:
                                            v54 := []int{9223372036854775807, (fi42 | 128)}
                                            fmt.Println(v2, fi42, fi42)
                                            _ = v54
                                        }
                                        v2 = (^(v2))
                                        var v55 string = "hello"
                                        _ = v55
                                    }
                                    _ = fi42
                                }
                                if false {
                                    if (285 > 1) {
                                        v2 -= (v2 / ((128) | 1))
                                        var v56 uint8 = (^(v2))
                                        _ = v56
                                    } else {
                                        v57 := []int{(fi21 | fi21), fi21, 314, fi21}
                                        _ = v57
                                    }
                                    v2 -= 128
                                }
                                for fi58 := 0; fi58 < 19; fi58++ {
                                    var v59 uint8 = 255
                                    var v60 uint32 = 184
                                    v60 += (v60 % ((1) | 1))
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                            }
                            v2 = v2
                        } else {
                            v61 := GenS1{F0: -1, F1: false}
                            _ = v61
                        }
                    }
                    _ = len(v4)
                    var v62 uint16 = 564
                    _ = v62
                }
            }
        case -9223372036854775808:
            if (false || (-97 >= -89)) {
                if false {
                    if true {
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            var v63 bool = true
                            _ = v63
                        }
                    }
                    v64 := []int{fi21}
                    _ = v64
                } else {
                    v2 += (v2 >> ((255) & 63))
                    fmt.Println(fi21)
                    if ((244 != 128) || true) {
                        v65 := []int{(fi21 / ((-9223372036854775808) | 1))}
                        for fi66 := 0; fi66 < 5; fi66++ {
                            fmt.Println(fi66, fi66, fi66)
                            _ = fi66
                        }
                        _ = v65
                    } else {
                        var v67 uint8 = v2
                        v68 := GenS4{F0: -128, F1: 1, F2: 0}
                        if (true || (!false)) {
                            var v69 int32 = 762
                            _ = v69
                        } else {
                            var v70 bool = ((!true) && (true && (!true)))
                            _ = v70
                        }
                        _ = v67
                        _ = v68
                    }
                }
                v2 = v2
            } else {
                for fi71 := 0; fi71 < 8; fi71++ {
                    if false {
                        v72 := v3[(^(fi21))]
                        v72 -= 65535
                        var v73 int32 = 1
                        v74 := GenS4{F0: -1, F1: 749, F2: 4294967295}
                        _ = v72
                        _ = v73
                        _ = v74
                    } else {
                        if ((!(!(false || ("hello" > "hello")))) || true) {
                            v2 = v2
                        } else {
                            var v75 uint16 = 0
                            _ = v75
                        }
                    }
                    _ = fi71
                }
            }
        }
        switch ("hello" + "\n\t") {
        case "\n\t":
            v1[2147483647] = 188
        default:
            for fi76 := 0; fi76 < 7; fi76++ {
                var v77 int8 = -99
                v77 += v77
                switch fi21 {
                case 9223372036854775807:
                    var v78 int32 = 1
                    _ = v78
                case 0:
                    var v79 uint32 = 42
                    _ = v79
                }
                _ = fi76
                _ = v77
            }
        }
        _ = fi21
    }
    for fi80 := 0; fi80 < 17; fi80++ {
        fmt.Println(fi80)
        _ = fi80
    }
    v2 -= v2
    var v81 int8 = 12
    if false {
        v1[1] = 658
        fmt.Println(v2, v2)
        switch 2147483647 {
        case -54:
            for fi82 := 0; fi82 < 17; fi82++ {
                for fi83 := 0; fi83 < 12; fi83++ {
                    v81 = int8(fi82)
                    var v84 int64 = 1
                    v81 += v81
                    _ = fi83
                    _ = v84
                }
                _ = fi82
            }
        default:
            var v85 uint8 = (v2 + 128)
            _ = v85
        }
    } else {
        if true {
            v86 := GenS6{F0: true, F1: (v81 ^ (v81 & (v81 ^ 100))), F2: 1}
            v2 -= v2
            _ = v86
        } else {
            if (!(!false)) {
                v87 := []int{42, 2147483647, 32768, 42}
                fmt.Println(v2, v2)
                v81 -= -1
                _ = v87
            } else {
                var v88 int8 = (v81 % ((v81) | 1))
                fmt.Println(v2)
                v88 -= v81
                _ = v88
            }
        }
        v81 += (-(v81))
        if ((float64(3.14159) == 0.1) || (845 < 4294967295)) {
            fmt.Println(v2, v81)
            v89 := GenS7{F0: 128, F1: (42 != 42)}
            _ = v89
        }
    }
    v90 := []int{2147483647, 1, 9223372036854775807, 682}
    v91 := GenS7{F0: (^(v2)), F1: false}
    v92 := GenS6{F0: false, F1: (^(v81)), F2: -9223372036854775808}
    for fi93 := 0; fi93 < 19; fi93++ {
        for fi94 := 0; fi94 < 13; fi94++ {
            v95 := GenS8{F0: 42}
            _ = fi94
            _ = v95
        }
        _ = fi93
    }
    switch "hello" {
    case "hello":
        v3[-881] = 1
    case "foo bar":
        for fi96 := 0; fi96 < 2; fi96++ {
            switch 1 {
            case 0:
                if false {
                    for fi97 := 0; fi97 < 2; fi97++ {
                        v2 += (v2 / (((v2 | 0)) | 1))
                        if false {
                            switch (-(fi97)) {
                            case 2147483647:
                                fmt.Println(fi97, fi96, fi97)
                            default:
                                _ = len(v1)
                            }
                            if (!(2 <= 4294967295)) {
                                _ = len(v90)
                                if true {
                                    v98 := GenS9{F0: 0}
                                    var v99 int32 = 0
                                    if false {
                                        if (379 == 115) {
                                            var v100 float32 = (3.14159 * 0.1)
                                            if false {
                                                v100 += ((float32(-2.25) - 1.5) / (v100 - (float32((1e10 / (((float64(-2.25) * 0.1) - ((1.5 / 0.1) / 1.0)) / float64((((1.5 * 0.0) + v100) / v100))))) / v100)))
                                                switch "a" {
                                                case "foo bar":
                                                    var v101 int = fi96
                                                    _ = v101
                                                case "\n\t":
                                                    switch (("" + ((("\n\t" + "foo bar") + (("hello" + "hello") + "a")) + ("" + "\n\t"))) + ("\n\t" + ("hello" + "a"))) {
                                                    case "a":
                                                        v100 = v100
                                                    case "\n\t":
                                                        var v102 float32 = v100
                                                        _ = v102
                                                    case "hello":
                                                        var v103 int64 = 421
                                                        var v104 uint8 = (v2 >> ((255) & 63))
                                                        fmt.Println(fi96, v81, fi97)
                                                        if (((v103 % ((v103) | 1)) > 9223372036854775807) || (v103 == (v103 + (^(v103))))) {
                                                            switch (fi96 << ((18446744073709551615) & 63)) {
                                                            case 1:
                                                                v81 += v81
                                                                if false {
                                                                    if len(v90) >= 2 {
                                                                        v90 = v90[1:len(v90)]
                                                                    }
                                                                    v105 := GenS7{F0: v104, F1: false}
                                                                    switch ("" + ("" + "foo bar")) {
                                                                    case "hello":
                                                                        if (false && false) {
                                                                            var v106 float64 = 1e10
                                                                            _ = v106
                                                                        }
                                                                    case "":
                                                                        switch "" {
                                                                        case "a":
                                                                            var v107 int16 = 729
                                                                            _ = v107
                                                                        case "\n\t":
                                                                            v81 -= v81
                                                                        case "hello":
                                                                            v108 := GenS10{F0: "a", F1: "", F2: v99}
                                                                            _ = v108
                                                                        default:
                                                                            v81 -= (v81 / ((100) | 1))
                                                                        }
                                                                    case "a":
                                                                        var v109 uint = 4294967295
                                                                        if false {
                                                                            v90 = append(v90, fi97)
                                                                        } else {
                                                                            v103 -= -9223372036854775808
                                                                        }
                                                                        v81 -= (v81 / (((v81 + (v81 ^ int8(v109)))) | 1))
                                                                        _ = v109
                                                                    }
                                                                    fmt.Println(v100, v100)
                                                                    var v110 int16 = 1000
                                                                    _ = v105
                                                                    _ = v110
                                                                }
                                                            case 32768:
                                                                v111 := GenS8{F0: 18446744073709551615}
                                                                _ = v111
                                                            }
                                                            delete(v4, fi97)
                                                        } else {
                                                            v104 = 255
                                                            if ((((v100 / 1.0) < 0.0) || true) && (1e10 >= 1.5)) {
                                                                if len(v90) > 0 {
                                                                    v90[0] = 127
                                                                }
                                                            } else {
                                                                if (!true) {
                                                                    switch fi96 {
                                                                    case 586:
                                                                        delete(v4, (fi96 >> ((103) & 63)))
                                                                    case 4294967295:
                                                                        v112 := GenS7{F0: (v2 ^ (^(v104))), F1: false}
                                                                        _ = v112
                                                                    case -755:
                                                                        delete(v1, 2)
                                                                    }
                                                                }
                                                                if true {
                                                                    if ((float64(0.0) - -2.25) >= -2.25) {
                                                                        if false {
                                                                            switch (fi97 | (fi96 + (fi96 << ((65535) & 63)))) {
                                                                            case -66:
                                                                                switch ("" + "\n\t") {
                                                                                case "a":
                                                                                    v81 = (v81 + (v81 ^ 48))
                                                                                case "\n\t":
                                                                                    var v113 string = (("hello" + "a") + "foo bar")
                                                                                    _ = v113
                                                                                case "foo bar":
                                                                                    var v114 int8 = (v81 & (v81 * 127))
                                                                                    _ = v114
                                                                                default:
                                                                                    var v115 uint32 = 0
                                                                                    _ = v115
                                                                                }
                                                                                var v116 int64 = (v103 << ((65535) & 63))
                                                                                _ = v116
                                                                            default:
                                                                                var v117 int = fi96
                                                                                _ = v117
                                                                            }
                                                                        } else {
                                                                            v99 += v99
                                                                            switch (fi96 >> ((739) & 63)) {
                                                                            case 128:
                                                                                if true {
                                                                                    if len(v90) > 0 {
                                                                                        v90[0] = fi96
                                                                                    }
                                                                                    fmt.Println(v104, v2, v81)
                                                                                } else {
                                                                                    fmt.Println(v104)
                                                                                    v99 = (^(v99))
                                                                                    v81 += (v81 ^ v81)
                                                                                    v118 := GenS11{F0: v103, F1: v81}
                                                                                    _ = v118
                                                                                }
                                                                            default:
                                                                                var v119 uint = 0
                                                                                _ = v119
                                                                            }
                                                                        }
                                                                        v120 := GenS12{F0: (v99 % ((int32(v81)) | 1)), F1: v103}
                                                                        var v121 float64 = ((3.14159 - 0.1) / (-2.25 - float64((v100 + v100))))
                                                                        _ = v120
                                                                        _ = v121
                                                                    } else {
                                                                        if ((-2.25 / v100) != ((v100 - v100) * v100)) {
                                                                            if (925 < 315) {
                                                                                if (!((true && (((0 >= 10) || true) && false)) || true)) {
                                                                                    _ = len(v3)
                                                                                } else {
                                                                                    v103 = v103
                                                                                }
                                                                            } else {
                                                                                switch (^(fi97)) {
                                                                                case -738:
                                                                                    switch fi97 {
                                                                                    case 2147483647:
                                                                                        fmt.Println(v104)
                                                                                    case 38:
                                                                                        fmt.Println(v103)
                                                                                    }
                                                                                case 899:
                                                                                    var v122 uint16 = 797
                                                                                    _ = v122
                                                                                }
                                                                            }
                                                                            v90 = append(v90, (fi97 | 1))
                                                                            if (float64(v100) <= -2.25) {
                                                                                var v123 bool = false
                                                                                switch (("\n\t" + "\n\t") + "") {
                                                                                case "foo bar":
                                                                                    v99 += v99
                                                                                case "hello":
                                                                                    v3[(fi96 & (fi96 >> ((762) & 63)))] = 42
                                                                                default:
                                                                                    fmt.Println(v2, v104)
                                                                                }
                                                                                _ = v123
                                                                            } else {
                                                                                var v124 float32 = (float32(float64(v100)) + ((((((v100 / (float32(3.14159) * v100)) - 1.5) / v100) / float32(float64(1.5))) - v100) - ((float32(1e10) - (1.5 + 0.0)) / v100)))
                                                                                if false {
                                                                                    var v125 int16 = 1
                                                                                    if false {
                                                                                        if true {
                                                                                            var v126 float64 = 1e10
                                                                                            var v127 int32 = (v99 >> ((131) & 63))
                                                                                            var v128 uint = 4294967295
                                                                                            _ = v126
                                                                                            _ = v127
                                                                                            _ = v128
                                                                                        } else {
                                                                                            switch (("a" + "foo bar") + "") {
                                                                                            case "":
                                                                                                var v129 int16 = int16(v2)
                                                                                                _ = v129
                                                                                            default:
                                                                                                v130 := GenS2{F0: int8(v99)}
                                                                                                _ = v130
                                                                                            }
                                                                                            v103 = (v103 & v103)
                                                                                            v131 := GenS10{F0: ("\n\t" + "\n\t"), F1: ("\n\t" + (("\n\t" + ("\n\t" + ("a" + "\n\t"))) + "")), F2: v99}
                                                                                            _ = v131
                                                                                        }
                                                                                    } else {
                                                                                        v99 += v99
                                                                                    }
                                                                                    _ = v125
                                                                                }
                                                                                v132 := GenS11{F0: -1, F1: -34}
                                                                                _ = v124
                                                                                _ = v132
                                                                            }
                                                                        }
                                                                        if (!true) {
                                                                            var v133 uint64 = 1
                                                                            if (((v103 != (v103 & (v103 ^ (v103 / (((v103 >> ((2) & 63))) | 1))))) || (1 < v133)) && false) {
                                                                                v134 := GenS13{F0: v2, F1: "a"}
                                                                                v135 := GenS13{F0: v2, F1: ((((("foo bar" + "a") + "") + "a") + "\n\t") + "foo bar")}
                                                                                var v136 uint64 = (v133 | v133)
                                                                                var v137 string = ("" + "")
                                                                                if len(v90) >= 2 {
                                                                                    v90 = v90[1:len(v90)]
                                                                                }
                                                                                if ((1 != v99) || (fi97 == (fi97 / ((fi96) | 1)))) {
                                                                                    var v138 float64 = (0.1 * 1.5)
                                                                                    _ = len(v90)
                                                                                    v2 += (v2 + 128)
                                                                                    _ = v138
                                                                                } else {
                                                                                    v133 -= v133
                                                                                    if (true && true) {
                                                                                        switch (fi96 - fi97) {
                                                                                        case -991:
                                                                                            v90 = append(v90, (fi97 & (fi97 | (fi97 - (fi97 | fi96)))))
                                                                                        case -9223372036854775808:
                                                                                            v139 := GenS14{F0: 43}
                                                                                            _ = v139
                                                                                        case -2147483648:
                                                                                            v2 -= v2
                                                                                        default:
                                                                                            if (1.5 < (1.5 - 0.0)) {
                                                                                                v140 := GenS15{F0: ((v137 + v137) + ("foo bar" + v137))}
                                                                                                v141 := GenS16{F0: (v81 & v81)}
                                                                                                _ = v140
                                                                                                _ = v141
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v142 := GenS1{F0: -1, F1: true}
                                                                                    _ = v142
                                                                                }
                                                                                v143 := GenS3{F0: 739, F1: 42}
                                                                                if ((4294967295 != 361) || ((^(v133)) > v133)) {
                                                                                    v144 := GenS13{F0: (v104 | (v2 / (((^(v104))) | 1))), F1: (v137 + "hello")}
                                                                                    if false {
                                                                                        v100 = 3.14159
                                                                                        if (false && ("\n\t" < "a")) {
                                                                                            var v145 int8 = int8(fi96)
                                                                                            v145 -= -1
                                                                                            v137 = (v137 + v137)
                                                                                            _ = v145
                                                                                        }
                                                                                        v104 = uint8(v136)
                                                                                        var v146 uint64 = v133
                                                                                        var v147 uint64 = v146
                                                                                        fmt.Println(v146, v2, v104)
                                                                                        _ = v146
                                                                                        _ = v147
                                                                                    } else {
                                                                                        var v148 string = v137
                                                                                        v133 = v136
                                                                                        _ = v148
                                                                                    }
                                                                                    _ = v144
                                                                                } else {
                                                                                    fmt.Println(v81, fi96)
                                                                                    fmt.Println(v133, v103, v100)
                                                                                    switch (v137 + v137) {
                                                                                    case "":
                                                                                        if (v100 <= ((0.0 / (v100 / 3.14159)) * -2.25)) {
                                                                                            switch v137 {
                                                                                            case "a":
                                                                                                if ((42 <= 0) && (!true)) {
                                                                                                    if (v81 != (v81 + (v81 * (v81 << ((743) & 63))))) {
                                                                                                        fmt.Println(fi96, fi97)
                                                                                                    }
                                                                                                    var v149 int32 = 216
                                                                                                    var v150 float32 = (v100 * ((1.5 * v100) - v100))
                                                                                                    _ = v149
                                                                                                    _ = v150
                                                                                                }
                                                                                            default:
                                                                                                if false {
                                                                                                    if true {
                                                                                                        fmt.Println(v104)
                                                                                                        var v151 int8 = (v81 & v81)
                                                                                                        _ = v151
                                                                                                    }
                                                                                                } else {
                                                                                                    if len(v90) > 0 {
                                                                                                        v90[0] = fi97
                                                                                                    }
                                                                                                    v152 := GenS3{F0: 0, F1: 463}
                                                                                                    delete(v1, 2147483647)
                                                                                                    _ = v152
                                                                                                }
                                                                                                v153 := v3[fi96]
                                                                                                v133 = 691
                                                                                                _ = v153
                                                                                            }
                                                                                            var v154 int16 = -551
                                                                                            _ = v154
                                                                                        }
                                                                                    default:
                                                                                        var v155 uint32 = 0
                                                                                        _ = v155
                                                                                    }
                                                                                }
                                                                                var v156 uint64 = 18446744073709551615
                                                                                fmt.Println(v99, v133)
                                                                                _ = v134
                                                                                _ = v135
                                                                                _ = v136
                                                                                _ = v137
                                                                                _ = v143
                                                                                _ = v156
                                                                            } else {
                                                                                if false {
                                                                                    if len(v90) > 0 {
                                                                                        v90[0] = fi96
                                                                                    }
                                                                                    var v157 uint8 = (v2 | uint8(v133))
                                                                                    var v158 int = fi97
                                                                                    v99 = v99
                                                                                    _ = v157
                                                                                    _ = v158
                                                                                } else {
                                                                                    v90 = append(v90, (fi97 + 127))
                                                                                    fmt.Println(v100, v100, fi96)
                                                                                    fmt.Println(v2)
                                                                                    if (v81 < (v81 << ((1) & 63))) {
                                                                                        fmt.Println(v103, v2, v2)
                                                                                        if false {
                                                                                            switch (fi96 << ((796) & 63)) {
                                                                                            case 991:
                                                                                                fmt.Println(v104, v99)
                                                                                            case -1:
                                                                                                v159 := GenS17{F0: 32767}
                                                                                                _ = v159
                                                                                            case 0:
                                                                                                fmt.Println(v103, fi97)
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v133, v100)
                                                                                            if ((v133 * v133) >= (^(v133))) {
                                                                                                v99 = 42
                                                                                                if (("\n\t" + "") <= "a") {
                                                                                                    var v160 int64 = 32
                                                                                                    var v161 int32 = v99
                                                                                                    v133 = 18446744073709551615
                                                                                                    _ = v160
                                                                                                    _ = v161
                                                                                                } else {
                                                                                                    _ = len(v1)
                                                                                                    v100 = v100
                                                                                                }
                                                                                            }
                                                                                            var v162 float64 = 0.0
                                                                                            if ((2 < 967) && ((994 >= 302) || false)) {
                                                                                                if len(v90) > 0 {
                                                                                                    v90[0] = (fi96 & (fi96 * (fi97 / ((-1) | 1))))
                                                                                                }
                                                                                            } else {
                                                                                                v90 = append(v90, int(v81))
                                                                                            }
                                                                                            _ = v162
                                                                                        }
                                                                                    } else {
                                                                                        _ = len(v90)
                                                                                    }
                                                                                    var v163 uint64 = uint64(v2)
                                                                                    switch int(v81) {
                                                                                    case 42:
                                                                                        if true {
                                                                                            v164 := GenS3{F0: 255, F1: 65535}
                                                                                            fmt.Println(v103)
                                                                                            _ = len(v4)
                                                                                            _ = v164
                                                                                        }
                                                                                    default:
                                                                                        var v165 uint32 = uint32(v163)
                                                                                        _ = v165
                                                                                    }
                                                                                    _ = v163
                                                                                }
                                                                                if (!false) {
                                                                                    delete(v4, int(v133))
                                                                                    fmt.Println(v104, fi96, fi96)
                                                                                } else {
                                                                                    var v166 int8 = (v81 & v81)
                                                                                    var v167 uint16 = 65535
                                                                                    v168 := GenS8{F0: (v133 % ((v133) | 1))}
                                                                                    _ = v166
                                                                                    _ = v167
                                                                                    _ = v168
                                                                                }
                                                                                if false {
                                                                                    if true {
                                                                                        fmt.Println(v99, v81, fi96)
                                                                                        if (false && false) {
                                                                                            fmt.Println(v2)
                                                                                            var v169 int16 = 42
                                                                                            _ = v169
                                                                                        } else {
                                                                                            var v170 string = ("" + "hello")
                                                                                            var v171 uint32 = 42
                                                                                            v172 := v1[v171]
                                                                                            _ = v170
                                                                                            _ = v171
                                                                                            _ = v172
                                                                                        }
                                                                                    }
                                                                                    var v173 uint8 = v2
                                                                                    _ = v173
                                                                                } else {
                                                                                    fmt.Println(v103, v133, v99)
                                                                                    var v174 int = fi96
                                                                                    var v175 float64 = ((((0.0 / (-2.25 * ((0.1 - (-2.25 + (3.14159 - float64(v100)))) * ((1e10 - 1.5) / 1e10)))) + 0.0) / 1.5) + ((1.5 - 0.1) / 3.14159))
                                                                                    v176 := GenS18{F0: 42}
                                                                                    _ = v174
                                                                                    _ = v175
                                                                                    _ = v176
                                                                                }
                                                                            }
                                                                            var v177 int8 = (v81 + (v81 * (v81 & -4)))
                                                                            _ = v133
                                                                            _ = v177
                                                                        } else {
                                                                            v178 := GenS19{F0: v103, F1: (v99 * (^(v99))), F2: v81}
                                                                            if false {
                                                                                if (0.1 == v100) {
                                                                                    if len(v90) >= 2 {
                                                                                        v90 = v90[1:len(v90)]
                                                                                    }
                                                                                    fmt.Println(fi97, v100)
                                                                                }
                                                                            }
                                                                            var v179 uint32 = 42
                                                                            _ = v178
                                                                            _ = v179
                                                                        }
                                                                    }
                                                                    var v180 bool = (!false)
                                                                    _ = v180
                                                                } else {
                                                                    v90 = append(v90, fi97)
                                                                    var v181 string = (("foo bar" + "\n\t") + ("foo bar" + "\n\t"))
                                                                    _ = v181
                                                                }
                                                            }
                                                        }
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                case "hello":
                                                    if false {
                                                        v81 += 100
                                                        if ((true && (false && true)) && true) {
                                                            var v182 uint16 = 2
                                                            _ = v182
                                                        }
                                                    }
                                                }
                                            } else {
                                                v100 = v100
                                                v183 := GenS6{F0: (-1 > 42), F1: 0, F2: 9223372036854775807}
                                                _ = v183
                                            }
                                            if false {
                                                v184 := GenS20{F0: 132, F1: fi97, F2: 1}
                                                _ = v184
                                            }
                                            _ = v100
                                        } else {
                                            v4[fi96] = ("a" + "a")
                                            _ = len(v90)
                                            v185 := GenS12{F0: (v99 * v99), F1: -1}
                                            _ = v185
                                        }
                                        v4[(fi97 - -939)] = ""
                                    }
                                    _ = v98
                                    _ = v99
                                }
                            } else {
                                var v186 string = "hello"
                                _ = v186
                            }
                        }
                        _ = fi97
                    }
                }
            case 128:
                fmt.Println(v81, v2)
            case 95:
                if false {
                    var v187 int = (fi96 << ((581) & 63))
                    _ = v187
                }
            default:
                switch ("a" + (("\n\t" + "\n\t") + ("a" + "a"))) {
                case "":
                    v2 = v2
                case "hello":
                    v81 -= v81
                }
            }
            _ = fi96
        }
    case "":
        v2 -= (^(v2))
    }
    if false {
        var v188 uint8 = v2
        _ = v188
    } else {
        if false {
            for fi189 := 0; fi189 < 5; fi189++ {
                var v190 string = "a"
                if true {
                    v191 := GenS3{F0: 42, F1: 65535}
                    var v192 int64 = 1
                    _ = v191
                    _ = v192
                } else {
                    v190 = v190
                    if true {
                        for fi193 := 0; fi193 < 4; fi193++ {
                            var v194 int64 = -9223372036854775808
                            switch (^(fi193)) {
                            case -1:
                                fmt.Println(v194, fi193)
                            case 48:
                                var v195 float32 = (0.0 + 3.14159)
                                _ = v195
                            case -9223372036854775808:
                                var v196 float64 = 0.0
                                _ = v196
                            }
                            var v197 uint16 = 65535
                            v81 = int8(v2)
                            v197 = (^(v197))
                            _ = fi193
                            _ = v194
                            _ = v197
                        }
                        if (!false) {
                            v198 := GenS17{F0: -557}
                            var v199 uint = 662
                            for fi200 := 0; fi200 < 2; fi200++ {
                                v201 := GenS16{F0: 127}
                                v190 = v190
                                _ = len(v90)
                                _ = fi200
                                _ = v201
                            }
                            _ = v198
                            _ = v199
                        }
                        v81 = (v81 >> ((364) & 63))
                    }
                }
                _ = len(v1)
                _ = fi189
                _ = v190
            }
            var v202 uint8 = v2
            _ = v202
        }
    }
    var v203 float32 = ((((((float32(1e10) - (0.0 / ((1e10 - 3.14159) + 0.0))) / ((0.1 + -2.25) - (((((0.1 - 0.0) + -2.25) * 0.0) + 1.5) + 0.0))) * ((1.5 / 1e10) / (3.14159 * -2.25))) / (0.0 - (0.1 + (0.0 - (0.1 - 1e10))))) - 0.1) * (1.5 + 1.5))
    v81 -= v81
    var v204 uint = 65535
    fmt.Println(v204, v204)
    v204 = v204
    switch ("foo bar" + (("" + "hello") + ("a" + ((("\n\t" + ("foo bar" + ((("foo bar" + ("a" + "")) + "a") + "\n\t"))) + (("a" + "foo bar") + ("hello" + "hello"))) + ((("\n\t" + (("\n\t" + ("foo bar" + "\n\t")) + ("foo bar" + "hello"))) + "hello") + "\n\t"))))) {
    case "foo bar":
        for fi205 := 0; fi205 < 20; fi205++ {
            for fi206 := 0; fi206 < 20; fi206++ {
                fmt.Println(v203)
                _ = fi206
            }
            if ("foo bar" != "hello") {
                v2 = (v2 / ((v2) | 1))
                var v207 string = "foo bar"
                _ = v207
            }
            var v208 uint = 65535
            _ = fi205
            _ = v208
        }
    case "a":
        var v209 uint64 = 0
        _ = v209
    default:
        for fi210 := 0; fi210 < 3; fi210++ {
            _ = len(v1)
            _ = fi210
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("v204:", v204)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v81:", v81)
    fmt.Println("slicelen v90:", len(v90), v90)
    _ = v1
    _ = v2
    _ = v203
    _ = v204
    _ = v3
    _ = v4
    _ = v81
    _ = v90
    _ = v91
    _ = v92
}
