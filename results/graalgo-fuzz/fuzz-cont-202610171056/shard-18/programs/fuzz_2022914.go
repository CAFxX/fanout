package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint8; F2 bool }
type GenS5 struct { F0 int64; F1 int16; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int8; F1 int64; F2 int32 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 uint64; F1 int; F2 int64 }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS14 struct { F0 uint64; F1 int64; F2 int64 }
type GenS15 struct { F0 int; F1 uint64 }
type GenS16 struct { F0 uint64; F1 int }
type GenS17 struct { F0 int; F1 uint }
type GenS18 struct { F0 int; F1 uint16 }
type GenS19 struct { F0 uint32; F1 uint8 }
type GenS20 struct { F0 int16; F1 uint16 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 int; F1 int64 }
func genFn0() int32 {
    var v6 int = 127
    v6 = -940
    v7 := GenS1{F0: ((("hello" + ("hello" + "\n\t")) + "foo bar") + "\n\t"), F1: 0}
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: "hello", F1: 66}
        v6 = (fi8 << ((42) & 63))
        _ = fi8
        _ = v9
    }
    switch v6 {
    case 4294967295:
        var v10 int32 = 2147483647
        _ = v10
    case 127:
        var v11 int8 = 109
        _ = v11
    }
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 uint, p1 bool, p2 uint8) uint16 {
    switch 1 {
    case 634:
        p0 = p0
    case 32767:
        fmt.Println(p2)
    }
    if false {
        v12 := []int{32767}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        v13 := GenS2{F0: -1}
        _ = v12
        _ = v13
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return uint16(p0)
}

func main() {
    var v1 float64 = 0.1
    var v2 int16 = 1
    v3 := []uint32{2}
    var v4 int8 = 100
    v5 := []int{4294967295, 0}
    v14 := GenS2{F0: (-(v2))}
    for fi15 := 0; fi15 < 15; fi15++ {
        switch "hello" {
        case "":
            var v16 int8 = (^(v4))
            _ = v16
        case "\n\t":
            v17 := GenS3{F0: (v4 / (((-(v4))) | 1))}
            _ = v17
        case "hello":
            if (!(((fi15 == (fi15 % ((fi15) | 1))) || true) && true)) {
                if len(v5) > 0 {
                    v5[0] = (fi15 + (fi15 ^ fi15))
                }
                var v18 bool = true
                _ = v18
            } else {
                if (42 != -9223372036854775808) {
                    switch fi15 {
                    case 32768:
                        switch "\n\t" {
                        case "":
                            v1 = v1
                        case "hello":
                            var v19 int64 = -1
                            _ = v19
                        case "\n\t":
                            var v20 float32 = 1e10
                            _ = v20
                        default:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case 9223372036854775807:
                        var v21 uint32 = 4294967295
                        _ = v21
                    case -959:
                        v22 := GenS4{F0: (true || ((1.5 / v1) == (1e10 + float64(((0.1 * (1.5 + (-2.25 * float32(0.0)))) / (1.5 - 1e10)))))), F1: 123, F2: (!(!(("" + "\n\t") > "foo bar")))}
                        _ = v22
                    }
                    if (!((!(false && (false && false))) || false)) {
                        fmt.Println(v4, v1)
                        if true {
                            v23 := GenS1{F0: ("a" + "\n\t"), F1: 2}
                            switch (fi15 / (((fi15 & fi15)) | 1)) {
                            case 366:
                                for fi24 := 0; fi24 < 15; fi24++ {
                                    v25 := GenS5{F0: 0, F1: (v2 & (v2 * v2)), F2: 42}
                                    v2 = int16(v4)
                                    v26 := GenS1{F0: "foo bar", F1: 2}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                            case -2147483648:
                                v4 += (v4 * (v4 - int8(v2)))
                            case 2147483647:
                                if true {
                                    v4 -= int8(fi15)
                                    var v27 bool = (v1 < ((v1 - 0.0) * 0.1))
                                    v28 := GenS6{F0: 185, F1: 4294967295}
                                    _ = v27
                                    _ = v28
                                } else {
                                    v1 = (v1 * float64((3.14159 * ((0.0 + 0.0) / float32((0.0 + (((-2.25 * v1) / v1) + 0.1)))))))
                                    v2 = (v2 * (-(v2)))
                                }
                            default:
                                for fi29 := 0; fi29 < 18; fi29++ {
                                    switch (fi15 << ((911) & 63)) {
                                    case 0:
                                        _ = len(v5)
                                    case 1:
                                        switch (("hello" + "hello") + ("hello" + "a")) {
                                        case "\n\t":
                                            fmt.Println(v1, v2, fi15)
                                        default:
                                            v2 = -634
                                        }
                                    default:
                                        v2 = (v2 & (v2 | v2))
                                    }
                                    v2 -= (v2 / ((v2) | 1))
                                    v30 := GenS6{F0: 0, F1: 4294967295}
                                    _ = fi29
                                    _ = v30
                                }
                            }
                            v2 = v2
                            _ = v23
                        }
                    }
                }
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            }
        default:
            var v31 bool = (true && false)
            _ = v31
        }
        _ = fi15
    }
    var v32 int8 = v4
    for fi33 := 0; fi33 < 19; fi33++ {
        var v34 float32 = 1e10
        _ = fi33
        _ = v34
    }
    var v35 bool = false
    switch 108 {
    case -911:
        if v35 {
            if v35 {
                v35 = v35
            }
        } else {
            var v36 int8 = -1
            if true {
                v1 = ((v1 * 3.14159) * -2.25)
                var v37 uint16 = 42
                _ = v37
            }
            for fi38 := 0; fi38 < 7; fi38++ {
                for fi39 := 0; fi39 < 2; fi39++ {
                    var v40 uint32 = 2147483648
                    _ = fi39
                    _ = v40
                }
                _ = fi38
            }
            _ = v36
        }
    }
    switch -44 {
    case -428:
        var v41 bool = false
        _ = v41
    case 673:
        _ = len(v5)
    case 0:
        var v42 uint64 = 42
        _ = v42
    default:
        for fi43 := 0; fi43 < 16; fi43++ {
            switch 127 {
            case -692:
                v1 = (v1 - v1)
            case 1:
                var v44 uint8 = 2
                _ = v44
            }
            if true {
                v2 += (-(v2))
                var v45 bool = true
                v46 := GenS7{F0: -362}
                _ = v45
                _ = v46
            }
            _ = fi43
        }
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        for fi48 := 0; fi48 < 11; fi48++ {
            var v49 int16 = (v2 + 32767)
            var v50 uint32 = 65535
            if (255 == 65535) {
                switch 0 {
                case 127:
                    v2 += 32767
                case 0:
                    var v51 float32 = 1.5
                    _ = v51
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            } else {
                fmt.Println(v32, v32, fi47)
                v32 += (-(v32))
            }
            var v52 string = "\n\t"
            _ = fi48
            _ = v49
            _ = v50
            _ = v52
        }
        fmt.Println(v32, v2, v32)
        if (!v35) {
            for fi53 := 0; fi53 < 16; fi53++ {
                fmt.Println(v35, v32, v2)
                _ = fi53
            }
            fmt.Println(v4, v32)
            for fi54 := 0; fi54 < 20; fi54++ {
                switch fi47 {
                case 581:
                    fmt.Println(v2)
                case 127:
                    var v55 int32 = 0
                    _ = v55
                }
                var v56 bool = (!true)
                _ = fi54
                _ = v56
            }
        }
        _ = fi47
    }
    switch "a" {
    case "foo bar":
        if (float64((-2.25 - float32(v1))) <= 3.14159) {
            v57 := GenS8{F0: 0}
            switch 246 {
            case 1:
                for fi58 := 0; fi58 < 5; fi58++ {
                    v59 := GenS4{F0: (128 == 102), F1: 1, F2: (!((-(v32)) == 127))}
                    v60 := GenS1{F0: "\n\t", F1: 0}
                    var v61 int16 = v2
                    _ = fi58
                    _ = v59
                    _ = v60
                    _ = v61
                }
            case 127:
                var v62 float64 = v1
                _ = v62
            case 32767:
                var v63 int16 = -573
                _ = v63
            }
            if true {
                if (false && v35) {
                    switch ("a" + "\n\t") {
                    case "hello":
                        v1 += (v1 - 1.5)
                    case "foo bar":
                        for fi64 := 0; fi64 < 2; fi64++ {
                            _ = len(v5)
                            if false {
                                if v35 {
                                    v2 += (-(v2))
                                    if v35 {
                                        switch fi64 {
                                        case -618:
                                            for fi65 := 0; fi65 < 19; fi65++ {
                                                v1 += ((v1 * v1) - ((v1 - 0.1) - (3.14159 * (((float64((((3.14159 / 1.5) - (0.0 / -2.25)) * (3.14159 - 1.5))) * v1) + 3.14159) - (((((v1 + 1e10) * (v1 - ((0.1 / -2.25) - (v1 - v1)))) + v1) * float64(0.0)) - v1)))))
                                                var v66 uint8 = 255
                                                switch (("foo bar" + "foo bar") + "") {
                                                case "hello":
                                                    v35 = v35
                                                case "a":
                                                    var v67 int32 = -2147483648
                                                    _ = v67
                                                }
                                                _ = fi65
                                                _ = v66
                                            }
                                        default:
                                            for fi68 := 0; fi68 < 13; fi68++ {
                                                v4 = v4
                                                _ = len(v5)
                                                if (!true) {
                                                    v32 = v32
                                                    fmt.Println(v32, v1)
                                                } else {
                                                    fmt.Println(v4, fi64)
                                                    v5 = append(v5, -458)
                                                    if (0.1 >= 0.1) {
                                                        fmt.Println(fi64, v4)
                                                    } else {
                                                        if v35 {
                                                            v35 = ((1.5 - v1) < (v1 * 1e10))
                                                        }
                                                    }
                                                    v4 = 16
                                                }
                                                var v69 uint16 = 256
                                                var v70 uint32 = 2
                                                _ = fi68
                                                _ = v69
                                                _ = v70
                                            }
                                        }
                                        v71 := GenS3{F0: 1}
                                        _ = v71
                                    } else {
                                        if v35 {
                                            fmt.Println(v2)
                                            if ((1 > 0) && (-1 > int64(fi64))) {
                                                v2 = v2
                                                v72 := GenS6{F0: 1, F1: 247}
                                                fmt.Println(v35, fi64)
                                                _ = v72
                                            } else {
                                                v4 = (-(v4))
                                                if true {
                                                    var v73 float32 = 0.0
                                                    for fi74 := 0; fi74 < 9; fi74++ {
                                                        v2 -= 275
                                                        if true {
                                                            v1 = v1
                                                        }
                                                        _ = fi74
                                                    }
                                                    v75 := GenS9{F0: v32, F1: -1, F2: -2147483648}
                                                    _ = v73
                                                    _ = v75
                                                } else {
                                                    var v76 int16 = v2
                                                    fmt.Println(v1, v32)
                                                    fmt.Println(v4, v2, fi64)
                                                    _ = v76
                                                }
                                                var v77 int = fi64
                                                _ = v77
                                            }
                                            var v78 uint32 = 2147483648
                                            _ = v78
                                        } else {
                                            if v35 {
                                                if (!(0 != 671)) {
                                                    for fi79 := 0; fi79 < 14; fi79++ {
                                                        _ = len(v3)
                                                        fmt.Println(fi64)
                                                        var v80 float32 = 0.1
                                                        _ = fi79
                                                        _ = v80
                                                    }
                                                    _ = len(v3)
                                                    switch "" {
                                                    case "a":
                                                        v2 -= 42
                                                    case "\n\t":
                                                        if v35 {
                                                            v32 += -128
                                                            var v81 int8 = v4
                                                            switch (("foo bar" + ("hello" + ("foo bar" + ("a" + "")))) + "") {
                                                            case "":
                                                                if v35 {
                                                                    var v82 bool = ((v2 | (v2 + -578)) > 0)
                                                                    _ = v82
                                                                }
                                                            case "hello":
                                                                for fi83 := 0; fi83 < 9; fi83++ {
                                                                    if (!("foo bar" < ((((("" + ("foo bar" + "hello")) + "\n\t") + "foo bar") + "hello") + "hello"))) {
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        var v84 uint8 = 81
                                                                        var v85 bool = (!(255 <= 65535))
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                    _ = fi83
                                                                }
                                                            case "\n\t":
                                                                if (2 > 2) {
                                                                    var v86 int32 = 2147483647
                                                                    switch fi64 {
                                                                    case 338:
                                                                        if v35 {
                                                                            if (-2.25 > 3.14159) {
                                                                                for fi87 := 0; fi87 < 2; fi87++ {
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = 216
                                                                                    }
                                                                                    fmt.Println(v4)
                                                                                    _ = fi87
                                                                                }
                                                                                fmt.Println(v2)
                                                                                v4 += (v81 + (v32 * (v4 >> ((42) & 63))))
                                                                            }
                                                                        } else {
                                                                            var v88 uint32 = 4294967295
                                                                            var v89 uint64 = 42
                                                                            for fi90 := 0; fi90 < 7; fi90++ {
                                                                                v86 = 42
                                                                                v86 += v86
                                                                                v35 = v35
                                                                                _ = fi90
                                                                            }
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                    _ = v86
                                                                } else {
                                                                    for fi91 := 0; fi91 < 12; fi91++ {
                                                                        if true {
                                                                            v1 = ((1.5 / 1.0) * v1)
                                                                            if v35 {
                                                                                switch "a" {
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        var v92 int8 = v81
                                                                                        switch 42 {
                                                                                        case 0:
                                                                                            if v35 {
                                                                                                var v93 uint16 = 65535
                                                                                                v94 := GenS10{F0: 860, F1: 248}
                                                                                                v93 -= genFn1(65535, (1 == -326), 1)
                                                                                                _ = v93
                                                                                                _ = v94
                                                                                            } else {
                                                                                                var v95 int64 = -809
                                                                                                if v35 {
                                                                                                    var v96 int32 = 277
                                                                                                    if ("\n\t" <= "foo bar") {
                                                                                                        _ = len(v5)
                                                                                                        var v97 int16 = (v2 | v2)
                                                                                                        switch int(v97) {
                                                                                                        case 650:
                                                                                                            switch "" {
                                                                                                            case "foo bar":
                                                                                                                _ = len(v3)
                                                                                                            case "a":
                                                                                                                v81 += (v4 / (((v81 / ((1) | 1))) | 1))
                                                                                                            default:
                                                                                                                fmt.Println(v35, fi91)
                                                                                                            }
                                                                                                        case 127:
                                                                                                            if v35 {
                                                                                                                switch "a" {
                                                                                                                case "a":
                                                                                                                    var v98 float64 = v1
                                                                                                                    _ = v98
                                                                                                                case "hello":
                                                                                                                    v81 = v92
                                                                                                                case "\n\t":
                                                                                                                    var v99 int = fi64
                                                                                                                    _ = v99
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v97
                                                                                                    } else {
                                                                                                        var v100 float32 = float32(v1)
                                                                                                        var v101 uint16 = 255
                                                                                                        _ = v100
                                                                                                        _ = v101
                                                                                                    }
                                                                                                    switch fi64 {
                                                                                                    case 9223372036854775807:
                                                                                                        if (!v35) {
                                                                                                            v95 += v95
                                                                                                            var v102 bool = false
                                                                                                            _ = v102
                                                                                                        }
                                                                                                        v103 := GenS11{F0: 876, F1: fi91, F2: (-(v95))}
                                                                                                        var v104 uint32 = 998
                                                                                                        var v105 int64 = int64(v96)
                                                                                                        _ = v103
                                                                                                        _ = v104
                                                                                                        _ = v105
                                                                                                    case -2147483648:
                                                                                                        v35 = (!true)
                                                                                                    default:
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = 2147483648
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v96
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        var v106 uint = 42
                                                                                                        _ = v106
                                                                                                    }
                                                                                                    if v35 {
                                                                                                        var v107 string = ("foo bar" + "\n\t")
                                                                                                        v1 -= 0.1
                                                                                                        v108 := GenS12{F0: 4294967295, F1: 0}
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        var v109 uint16 = 0
                                                                                                        var v110 uint64 = 1
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                        _ = v109
                                                                                                        _ = v110
                                                                                                    } else {
                                                                                                        var v111 int32 = 42
                                                                                                        _ = v111
                                                                                                    }
                                                                                                    if v35 {
                                                                                                        var v112 int = fi91
                                                                                                        var v113 uint8 = 0
                                                                                                        switch "hello" {
                                                                                                        case "":
                                                                                                            switch fi91 {
                                                                                                            case -113:
                                                                                                                fmt.Println(fi91)
                                                                                                            default:
                                                                                                                var v114 uint8 = v113
                                                                                                                _ = v114
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v112
                                                                                                        _ = v113
                                                                                                    } else {
                                                                                                        fmt.Println(fi91, v81, v32)
                                                                                                        var v115 int64 = (v95 | v95)
                                                                                                        var v116 int = -1
                                                                                                        _ = v115
                                                                                                        _ = v116
                                                                                                    }
                                                                                                    if (4294967295 != 18446744073709551615) {
                                                                                                        switch fi64 {
                                                                                                        case 128:
                                                                                                            var v117 string = ""
                                                                                                            _ = v117
                                                                                                        case -1:
                                                                                                            var v118 int64 = v95
                                                                                                            _ = v118
                                                                                                        case 127:
                                                                                                            var v119 int32 = 42
                                                                                                            _ = v119
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v92, v81, v2)
                                                                                                    switch "foo bar" {
                                                                                                    case "foo bar":
                                                                                                        _ = len(v3)
                                                                                                    case "\n\t":
                                                                                                        var v120 uint8 = 76
                                                                                                        _ = v120
                                                                                                    case "":
                                                                                                        var v121 uint32 = 49
                                                                                                        _ = v121
                                                                                                    default:
                                                                                                        switch (^(fi64)) {
                                                                                                        case 9223372036854775807:
                                                                                                            v122 := GenS13{F0: v2, F1: 0, F2: 1}
                                                                                                            _ = v122
                                                                                                        case -61:
                                                                                                            switch fi64 {
                                                                                                            case 2147483647:
                                                                                                                _ = len(v3)
                                                                                                            default:
                                                                                                                v81 += v92
                                                                                                            }
                                                                                                        case 128:
                                                                                                            switch (^(fi91)) {
                                                                                                            case 1:
                                                                                                                fmt.Println(v92)
                                                                                                            case 128:
                                                                                                                var v123 uint64 = 0
                                                                                                                _ = v123
                                                                                                            case -1:
                                                                                                                switch "a" {
                                                                                                                case "a":
                                                                                                                    if (!v35) {
                                                                                                                        fmt.Println(v35, v1, v2)
                                                                                                                        v4 = -1
                                                                                                                        if v35 {
                                                                                                                            v95 += -1
                                                                                                                            if (true || false) {
                                                                                                                                v4 += (v4 >> ((2) & 63))
                                                                                                                            } else {
                                                                                                                                if v35 {
                                                                                                                                    var v124 int16 = (v2 - v2)
                                                                                                                                    var v125 uint64 = 2
                                                                                                                                    var v126 uint16 = 758
                                                                                                                                    _ = v124
                                                                                                                                    _ = v125
                                                                                                                                    _ = v126
                                                                                                                                }
                                                                                                                                var v127 float64 = (((v1 - ((v1 - (float64(float32(0.0)) - -2.25)) + v1)) * (((v1 / (v1 + v1)) * float64(1e10)) + 1e10)) - 0.1)
                                                                                                                                _ = v127
                                                                                                                            }
                                                                                                                            var v128 float64 = (v1 * v1)
                                                                                                                            _ = v128
                                                                                                                        }
                                                                                                                        v95 += -9223372036854775808
                                                                                                                        v129 := GenS1{F0: ("foo bar" + ""), F1: 2}
                                                                                                                        _ = v129
                                                                                                                    } else {
                                                                                                                        if (fi91 >= (fi91 / ((fi91) | 1))) {
                                                                                                                            v3 = append(v3, 134)
                                                                                                                            v32 = v32
                                                                                                                            v130 := GenS14{F0: 18446744073709551615, F1: (v95 + (v95 | (v95 >> ((411) & 63)))), F2: (v95 * (v95 >> ((18446744073709551615) & 63)))}
                                                                                                                            _ = v130
                                                                                                                        } else {
                                                                                                                            var v131 float64 = v1
                                                                                                                            var v132 float32 = ((1.5 - ((3.14159 * (1e10 + 3.14159)) / (((1e10 - 1e10) / 3.14159) + -2.25))) - 0.1)
                                                                                                                            if true {
                                                                                                                                var v133 uint16 = 328
                                                                                                                                if (2 >= (v133 - (^(v133)))) {
                                                                                                                                    v5 = append(v5, (fi91 >> ((161) & 63)))
                                                                                                                                }
                                                                                                                                _ = v133
                                                                                                                            } else {
                                                                                                                                switch (fi64 - (fi64 % (((fi64 & 32767)) | 1))) {
                                                                                                                                case -1:
                                                                                                                                    _ = len(v3)
                                                                                                                                }
                                                                                                                                switch "hello" {
                                                                                                                                case "\n\t":
                                                                                                                                    var v134 float64 = -2.25
                                                                                                                                    _ = v134
                                                                                                                                case "a":
                                                                                                                                    fmt.Println(fi91, v4, v131)
                                                                                                                                }
                                                                                                                                v135 := GenS2{F0: 0}
                                                                                                                                _ = v135
                                                                                                                            }
                                                                                                                            _ = v131
                                                                                                                            _ = v132
                                                                                                                        }
                                                                                                                        v136 := GenS15{F0: fi64, F1: 18446744073709551615}
                                                                                                                        _ = v136
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            v35 = v35
                                                                                                            v3 = append(v3, 23)
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v95
                                                                                            }
                                                                                        case 4294967295:
                                                                                            fmt.Println(v1)
                                                                                        case 32767:
                                                                                            if false {
                                                                                                v4 -= 97
                                                                                                v137 := GenS6{F0: 255, F1: 42}
                                                                                                _ = v137
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(fi64, v92, v2)
                                                                                        _ = v92
                                                                                    }
                                                                                case "foo bar":
                                                                                    var v138 int8 = (-(v32))
                                                                                    _ = v138
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v4, v4, v81)
                                                                                v139 := GenS16{F0: 1, F1: (fi91 + fi64)}
                                                                                _ = v139
                                                                            }
                                                                            v4 = v81
                                                                        }
                                                                        if len(v5) > 0 {
                                                                            v5[0] = (fi64 >> ((0) & 63))
                                                                        }
                                                                        _ = len(v5)
                                                                        _ = fi91
                                                                    }
                                                                }
                                                            }
                                                            _ = v81
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = 403
                                                            }
                                                            v35 = true
                                                            v140 := GenS17{F0: 4294967295, F1: 42}
                                                            _ = v140
                                                        }
                                                    case "":
                                                        fmt.Println(v2, v2)
                                                    default:
                                                        v3 = append(v3, 1)
                                                    }
                                                }
                                                var v141 uint64 = 18446744073709551615
                                                for fi142 := 0; fi142 < 14; fi142++ {
                                                    switch fi64 {
                                                    case -898:
                                                        var v143 float64 = (1e10 / ((((v1 - (float64(1.5) / (v1 + v1))) - (1.5 / v1)) * (-2.25 * v1)) * float64(0.1)))
                                                        _ = v143
                                                    default:
                                                        switch "hello" {
                                                        case "a":
                                                            if (!(255 != 127)) {
                                                                fmt.Println(fi142)
                                                                v144 := GenS5{F0: 694, F1: 554, F2: 9223372036854775807}
                                                                var v145 int32 = 0
                                                                _ = v144
                                                                _ = v145
                                                            } else {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v146 bool = (v35 && (false && (1000 <= v2)))
                                                                    _ = v146
                                                                case "hello":
                                                                    var v147 uint64 = v141
                                                                    _ = v147
                                                                }
                                                                fmt.Println(v32)
                                                                fmt.Println(v4, v141, v32)
                                                            }
                                                        case "":
                                                            v141 -= 18446744073709551615
                                                        default:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                    }
                                                    var v148 int64 = -9223372036854775808
                                                    v149 := GenS18{F0: (fi64 << ((42) & 63)), F1: 795}
                                                    _ = fi142
                                                    _ = v148
                                                    _ = v149
                                                }
                                                _ = v141
                                            } else {
                                                var v150 uint = 740
                                                _ = v150
                                            }
                                            for fi151 := 0; fi151 < 19; fi151++ {
                                                var v152 int8 = 100
                                                switch (fi64 >> ((18446744073709551615) & 63)) {
                                                case 4294967295:
                                                    if (4294967295 > 0) {
                                                        v153 := GenS19{F0: 543, F1: 251}
                                                        _ = v153
                                                    } else {
                                                        v154 := GenS20{F0: (v2 * -32768), F1: 65535}
                                                        if (989 >= -1) {
                                                            _ = len(v5)
                                                            var v155 uint16 = 2
                                                            var v156 int16 = (^(v2))
                                                            _ = v155
                                                            _ = v156
                                                        }
                                                        _ = v154
                                                    }
                                                case 1:
                                                    if ((0.1 / 1.0) < 3.14159) {
                                                        var v157 float64 = 1e10
                                                        fmt.Println(v32, v4, fi151)
                                                        _ = v157
                                                    } else {
                                                        v3 = append(v3, 711)
                                                        var v158 int64 = 1
                                                        _ = v158
                                                    }
                                                default:
                                                    fmt.Println(v2, v2, v152)
                                                }
                                                var v159 float64 = v1
                                                _ = fi151
                                                _ = v152
                                                _ = v159
                                            }
                                            v160 := GenS7{F0: genFn0()}
                                            _ = v160
                                        }
                                        if v35 {
                                            _ = len(v3)
                                            if len(v5) > 0 {
                                                v5[0] = 933
                                            }
                                        }
                                        _ = len(v3)
                                    }
                                } else {
                                    v1 += float64(0.0)
                                    switch fi64 {
                                    case -2147483648:
                                        switch int(v32) {
                                        case 32768:
                                            for fi161 := 0; fi161 < 17; fi161++ {
                                                if true {
                                                    var v162 int16 = v2
                                                    _ = v162
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = 959
                                                    }
                                                    var v163 uint16 = 65535
                                                    var v164 float32 = 0.0
                                                    _ = v163
                                                    _ = v164
                                                }
                                                _ = fi161
                                            }
                                        case -499:
                                            v165 := GenS5{F0: -104, F1: v2, F2: 654}
                                            _ = v165
                                        case 9223372036854775807:
                                            var v166 int32 = genFn0()
                                            _ = v166
                                        }
                                    case 128:
                                        if len(v3) > 0 {
                                            v3[0] = 273
                                        }
                                    case 266:
                                        var v167 int8 = v4
                                        _ = v167
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v168 uint32 = 876
                                        _ = v168
                                    case "":
                                        switch (fi64 * fi64) {
                                        case 127:
                                            var v169 float32 = 0.0
                                            _ = v169
                                        case 9223372036854775807:
                                            v4 = -17
                                        default:
                                            for fi170 := 0; fi170 < 18; fi170++ {
                                                v5 = append(v5, (^(fi170)))
                                                v4 = (-(v32))
                                                fmt.Println(v4)
                                                if false {
                                                    if (155 < 1) {
                                                        if ((v35 || (!((true && (true || (32767 > 100))) || v35))) || (0 > v2)) {
                                                            v4 = int8(v2)
                                                            v3 = append(v3, 2147483648)
                                                        }
                                                    } else {
                                                        if v35 {
                                                            if (-2.25 < (0.1 + float32((v1 + (((v1 - v1) / v1) - v1))))) {
                                                                v171 := GenS4{F0: (!true), F1: 128, F2: ((539 < 2147483647) && (true || ((2 != 40) || v35)))}
                                                                switch "a" {
                                                                case "\n\t":
                                                                    v1 = 0.0
                                                                case "a":
                                                                    var v172 string = "a"
                                                                    _ = v172
                                                                case "foo bar":
                                                                    _ = len(v5)
                                                                }
                                                                _ = v171
                                                            } else {
                                                                var v173 int16 = (v2 ^ v2)
                                                                _ = v173
                                                            }
                                                        }
                                                        v174 := GenS21{F0: 1}
                                                        _ = v174
                                                    }
                                                }
                                                _ = fi170
                                            }
                                        }
                                    case "a":
                                        var v175 uint64 = 0
                                        _ = v175
                                    default:
                                        if (false || (815 != 1)) {
                                            if ((-2147483648 != 2147483647) || (!v35)) {
                                                var v176 int32 = -650
                                                _ = v176
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            for fi177 := 0; fi177 < 6; fi177++ {
                                                if len(v3) > 0 {
                                                    v3[0] = 2
                                                }
                                                _ = fi177
                                            }
                                            v178 := GenS10{F0: 820, F1: 127}
                                            _ = v178
                                        }
                                    }
                                }
                                var v179 uint32 = 1
                                _ = v179
                            } else {
                                v5 = append(v5, -972)
                            }
                            _ = fi64
                        }
                    }
                }
                fmt.Println(v32, v32)
            }
            _ = v57
        }
    }
    switch 32768 {
    case 42:
        var v180 uint64 = 1
        _ = v180
    case 127:
        for fi181 := 0; fi181 < 4; fi181++ {
            if (v35 && true) {
                v1 = v1
                v4 += 127
            } else {
                for fi182 := 0; fi182 < 11; fi182++ {
                    var v183 uint16 = 555
                    fmt.Println(v2, v35, v2)
                    v1 = v1
                    _ = fi182
                    _ = v183
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi181
        }
    case 0:
        var v184 uint8 = 128
        _ = v184
    }
    var v185 int64 = 285
    v186 := GenS11{F0: 1, F1: -38, F2: v185}
    for fi187 := 0; fi187 < 8; fi187++ {
        v188 := GenS22{F0: fi187, F1: 672}
        v1 -= float64(float32((((1.5 / v1) - (v1 + (0.1 / v1))) * v1)))
        v3 = append(v3, 2)
        _ = fi187
        _ = v188
    }
    fmt.Println("v1:", v1)
    fmt.Println("v185:", v185)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v185
    _ = v186
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}
