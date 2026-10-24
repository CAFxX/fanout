package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint64; F2 bool }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32; F1 bool }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 uint32; F1 int; F2 string }
type GenS14 struct { F0 uint; F1 int; F2 int64 }
type GenS15 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS16 struct { F0 uint64; F1 int32 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int64; F1 int8; F2 int16 }
type GenS19 struct { F0 int; F1 string; F2 uint16 }
type GenS20 struct { F0 int16; F1 int16; F2 int }
type GenS21 struct { F0 int32; F1 bool; F2 uint16 }
type GenS22 struct { F0 uint16; F1 uint8 }
type GenS23 struct { F0 int }
type GenS24 struct { F0 int32; F1 int; F2 bool }
type GenS25 struct { F0 uint64 }
type GenS26 struct { F0 uint64 }
type GenS27 struct { F0 uint8; F1 int8 }
type GenS28 struct { F0 int32; F1 uint32; F2 int16 }
type GenS29 struct { F0 int16; F1 uint8 }
type GenS30 struct { F0 int64; F1 int16; F2 bool }
type GenS31 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 int16) int64 {
    p0 = p0
    var v7 int64 = 42
    if (4294967295 <= 42) {
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: -61, F1: 0, F2: 1}
            var v10 int16 = p0
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint16 = 326
        _ = v11
    } else {
        v12 := []int{42, 32768}
        for fi13 := 0; fi13 < 15; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(fi14, fi14, p0)
                if (!true) {
                    if false {
                        _ = len(v12)
                    }
                    v15 := GenS1{F0: v7, F1: 231, F2: 255}
                    v16 := GenS1{F0: int64(p0), F1: 769, F2: 4294967295}
                    _ = v15
                    _ = v16
                }
                _ = fi14
            }
            v7 -= (v7 ^ int64(fi13))
            if true {
                if (737 != 65535) {
                    v12 = append(v12, fi13)
                } else {
                    fmt.Println(fi13, v7)
                }
                v12 = append(v12, fi13)
                for fi17 := 0; fi17 < 20; fi17++ {
                    fmt.Println(p0, fi13)
                    p0 = (p0 - p0)
                    if true {
                        v7 -= v7
                        fmt.Println(p0, v7, p0)
                    }
                    _ = fi17
                }
            }
            _ = fi13
        }
        var v18 string = "a"
        _ = v12
        _ = v18
    }
    switch "a" {
    case "\n\t":
        var v19 uint = 0
        _ = v19
    case "":
        if (42 >= 551) {
            switch "\n\t" {
            case "":
                if false {
                    var v20 int64 = (v7 & (v7 - 42))
                    for fi21 := 0; fi21 < 8; fi21++ {
                        switch "a" {
                        case "a":
                            var v22 uint32 = 0
                            _ = v22
                        case "foo bar":
                            v20 -= v20
                        case "\n\t":
                            if false {
                                v23 := []int{}
                                p0 = 32767
                                var v24 uint8 = 2
                                _ = v23
                                _ = v24
                            } else {
                                v25 := GenS2{F0: 42, F1: 42}
                                _ = v25
                            }
                        default:
                            if ((v7 << ((0) & 63)) > (v7 - (v20 << ((255) & 63)))) {
                                v26 := []int{}
                                v27 := GenS1{F0: 9223372036854775807, F1: 2, F2: 1}
                                _ = v26
                                _ = v27
                            } else {
                                if (("hello" + "foo bar") == "foo bar") {
                                    fmt.Println(p0, p0)
                                }
                                v28 := GenS3{F0: (^(p0))}
                                switch "a" {
                                case "a":
                                    switch fi21 {
                                    case -1:
                                        v29 := GenS4{F0: (v20 << ((811) & 63))}
                                        _ = v29
                                    case 2147483647:
                                        fmt.Println(fi21)
                                    default:
                                        fmt.Println(v7)
                                    }
                                case "hello":
                                    fmt.Println(p0, fi21)
                                }
                                _ = v28
                            }
                        }
                        var v30 uint16 = 949
                        _ = fi21
                        _ = v30
                    }
                    _ = v20
                } else {
                    if ((3.14159 - 3.14159) < (-2.25 * 3.14159)) {
                        v7 -= v7
                        v7 += -9223372036854775808
                    } else {
                        for fi31 := 0; fi31 < 13; fi31++ {
                            if true {
                                p0 += -1
                            } else {
                                v32 := GenS2{F0: 42, F1: 255}
                                _ = v32
                            }
                            var v33 uint16 = 73
                            for fi34 := 0; fi34 < 15; fi34++ {
                                v7 -= (v7 >> ((65535) & 63))
                                var v35 int16 = p0
                                _ = fi34
                                _ = v35
                            }
                            _ = fi31
                            _ = v33
                        }
                        var v36 float32 = 0.0
                        var v37 int8 = -128
                        _ = v36
                        _ = v37
                    }
                    v38 := []int{-58, 32767}
                    _ = v38
                }
            default:
                if false {
                    switch 127 {
                    case 0:
                        var v39 string = ""
                        _ = v39
                    }
                    if true {
                        for fi40 := 0; fi40 < 14; fi40++ {
                            switch ("a" + "") {
                            case "foo bar":
                                for fi41 := 0; fi41 < 12; fi41++ {
                                    switch "a" {
                                    case "hello":
                                        v42 := []int{1, (fi41 | fi41), fi41, fi41}
                                        _ = v42
                                    case "\n\t":
                                        var v43 uint16 = 0
                                        _ = v43
                                    }
                                    _ = fi41
                                }
                            case "\n\t":
                                var v44 string = "hello"
                                _ = v44
                            case "a":
                                var v45 int = (fi40 >> ((2) & 63))
                                _ = v45
                            }
                            _ = fi40
                        }
                        for fi46 := 0; fi46 < 12; fi46++ {
                            for fi47 := 0; fi47 < 10; fi47++ {
                                v48 := GenS5{F0: (172 >= 42), F1: 882, F2: (0 != 127)}
                                if (1 > 18446744073709551615) {
                                    v49 := GenS6{F0: 233, F1: true}
                                    if true {
                                        var v50 int16 = p0
                                        _ = v50
                                    } else {
                                        var v51 int32 = 42
                                        switch 1 {
                                        case 32767:
                                            if false {
                                                switch fi46 {
                                                case -1:
                                                    p0 = (p0 + -132)
                                                case 4294967295:
                                                    if true {
                                                        var v52 float64 = 0.0
                                                        var v53 bool = ((p0 & p0) <= p0)
                                                        fmt.Println(fi46, v51, p0)
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                default:
                                                    v54 := GenS2{F0: 255, F1: 42}
                                                    _ = v54
                                                }
                                                var v55 int8 = -1
                                                _ = v55
                                            }
                                        case -786:
                                            v56 := []int{fi46, (-(fi46)), fi46}
                                            _ = v56
                                        }
                                        v57 := []int{fi47, -329}
                                        _ = v51
                                        _ = v57
                                    }
                                    _ = v49
                                } else {
                                    v7 += v7
                                    switch fi47 {
                                    case -9223372036854775808:
                                        if ((-2147483648 < 1) && (true && (p0 != (p0 - -1)))) {
                                            v58 := GenS7{F0: ((fi47 * fi46) != (^(fi47)))}
                                            v59 := []int{32767}
                                            _ = v58
                                            _ = v59
                                        } else {
                                            if false {
                                                var v60 int16 = 1000
                                                if (32767 > 1) {
                                                    v7 = 0
                                                    v61 := []int{(fi47 / (((fi46 - fi47)) | 1)), -2147483648, -9223372036854775808}
                                                    _ = v61
                                                } else {
                                                    v60 = v60
                                                    var v62 int16 = (v60 / ((991) | 1))
                                                    _ = v62
                                                }
                                                _ = v60
                                            } else {
                                                var v63 int64 = (-(v7))
                                                fmt.Println(p0, v7)
                                                if (!(13 <= 1)) {
                                                    var v64 uint16 = 11
                                                    v65 := []int{fi46, (fi47 * (fi46 - fi46)), fi47, (fi47 * fi47)}
                                                    _ = v64
                                                    _ = v65
                                                }
                                                v7 = v63
                                                if true {
                                                    if (!(!false)) {
                                                        var v66 float32 = 0.0
                                                        _ = v66
                                                    } else {
                                                        v7 -= v63
                                                        switch (fi47 + 32767) {
                                                        case 4294967295:
                                                            p0 += (p0 - 1000)
                                                        case -2147483648:
                                                            if (true || false) {
                                                                fmt.Println(fi46, p0)
                                                                if false {
                                                                    var v67 string = (("hello" + "\n\t") + ("\n\t" + "a"))
                                                                    var v68 uint = 18446744073709551615
                                                                    if false {
                                                                        v69 := GenS6{F0: 255, F1: true}
                                                                        if false {
                                                                            v70 := GenS8{F0: 246, F1: p0}
                                                                            var v71 int32 = -927
                                                                            var v72 int = (fi47 ^ fi46)
                                                                            v67 = (v67 + "foo bar")
                                                                            switch (v67 + ((v67 + "\n\t") + "foo bar")) {
                                                                            case "":
                                                                                v73 := []int{}
                                                                                _ = v73
                                                                            case "a":
                                                                                var v74 uint = (v68 ^ v68)
                                                                                _ = v74
                                                                            }
                                                                            _ = v70
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                        var v75 float64 = float64(((float32(-2.25) + ((((0.0 - 1e10) + ((3.14159 + 1.5) * (3.14159 - 1e10))) * (3.14159 + (1e10 * (1.5 / (((-2.25 + (-2.25 + (-2.25 / ((0.0 - float32(0.1)) / -2.25)))) + ((1e10 / 1e10) * (3.14159 - 0.0))) + (3.14159 * 3.14159)))))) * 0.1)) + float32(0.0)))
                                                                        _ = v69
                                                                        _ = v75
                                                                    } else {
                                                                        var v76 float64 = ((0.1 * -2.25) - 0.1)
                                                                        v77 := []int{fi47}
                                                                        var v78 int16 = 0
                                                                        v7 = 539
                                                                        if len(v77) >= 2 {
                                                                            v77 = v77[1:len(v77)]
                                                                        }
                                                                        _ = v76
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    fmt.Println(v63, fi46)
                                                                    if (42 > 1) {
                                                                        fmt.Println(v63, fi46)
                                                                    } else {
                                                                        var v79 int16 = 225
                                                                        _ = v79
                                                                    }
                                                                    if ((fi47 - (fi46 ^ fi46)) != (fi47 ^ fi46)) {
                                                                        fmt.Println(fi46, v7, fi46)
                                                                        p0 = (-(p0))
                                                                    } else {
                                                                        var v80 int16 = (-(p0))
                                                                        var v81 bool = false
                                                                        _ = v80
                                                                        _ = v81
                                                                    }
                                                                    p0 -= 42
                                                                    p0 += (p0 * p0)
                                                                    v82 := []int{fi46}
                                                                    v83 := GenS4{F0: (v7 >> ((4294967295) & 63))}
                                                                    if (!true) {
                                                                        p0 = (p0 & 32767)
                                                                        if (!((p0 - p0) >= 1)) {
                                                                            v7 -= 804
                                                                            v82 = append(v82, fi46)
                                                                            switch (^(fi47)) {
                                                                            case -813:
                                                                                v7 = (^(v63))
                                                                                fmt.Println(p0, fi47, v7)
                                                                                if true {
                                                                                    v84 := GenS6{F0: 239, F1: false}
                                                                                    if true {
                                                                                        var v85 int8 = -128
                                                                                        _ = v85
                                                                                    }
                                                                                    if false {
                                                                                        if (2 != 569) {
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                var v86 float32 = float32(0.1)
                                                                                                _ = v86
                                                                                            default:
                                                                                                var v87 bool = true
                                                                                                _ = v87
                                                                                            }
                                                                                            var v88 uint = 18446744073709551615
                                                                                            v7 = (v7 & (v63 & v63))
                                                                                            var v89 int8 = 38
                                                                                            _ = v88
                                                                                            _ = v89
                                                                                        }
                                                                                        fmt.Println(v7, fi47, v63)
                                                                                        v90 := GenS8{F0: 1, F1: (p0 >> ((1) & 63))}
                                                                                        fmt.Println(fi46, v7, fi47)
                                                                                        _ = v90
                                                                                    } else {
                                                                                        v7 += v63
                                                                                    }
                                                                                    _ = v84
                                                                                } else {
                                                                                    fmt.Println(fi46)
                                                                                    v63 = (v7 / ((v7) | 1))
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    switch (fi46 / ((128) | 1)) {
                                                                                    case 206:
                                                                                        switch (-(fi47)) {
                                                                                        case 1:
                                                                                            p0 = (p0 >> ((18446744073709551615) & 63))
                                                                                        case 42:
                                                                                            fmt.Println(v7, p0)
                                                                                            if (-2.25 > 1.5) {
                                                                                                fmt.Println(p0, v7, p0)
                                                                                                fmt.Println(v63, v63)
                                                                                                v63 -= v63
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            if true {
                                                                                                _ = len(v82)
                                                                                                v91 := GenS9{F0: 122, F1: 99}
                                                                                                p0 += (p0 % (((p0 / ((407) | 1))) | 1))
                                                                                                switch 42 {
                                                                                                case -2147483648:
                                                                                                    v92 := GenS10{F0: 2}
                                                                                                    _ = v92
                                                                                                case 836:
                                                                                                    v93 := GenS11{F0: -1, F1: false}
                                                                                                    _ = v93
                                                                                                case 32768:
                                                                                                    var v94 float32 = 3.14159
                                                                                                    _ = v94
                                                                                                default:
                                                                                                    switch fi47 {
                                                                                                    case 296:
                                                                                                        fmt.Println(p0, fi47)
                                                                                                        fmt.Println(v63)
                                                                                                        switch -979 {
                                                                                                        case 418:
                                                                                                            switch (("foo bar" + "") + "hello") {
                                                                                                            case "":
                                                                                                                if true {
                                                                                                                    if len(v82) > 0 {
                                                                                                                        v82[0] = (^(fi47))
                                                                                                                    }
                                                                                                                    var v95 int = (fi47 / (((^(fi47))) | 1))
                                                                                                                    if len(v82) >= 2 {
                                                                                                                        v82 = v82[1:len(v82)]
                                                                                                                    }
                                                                                                                    _ = v95
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                if len(v82) > 0 {
                                                                                                                    v82[0] = fi46
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                var v96 uint8 = 127
                                                                                                                _ = v96
                                                                                                            default:
                                                                                                                if true {
                                                                                                                    v82 = append(v82, -2147483648)
                                                                                                                    p0 = (p0 * (-(p0)))
                                                                                                                    var v97 int16 = p0
                                                                                                                    _ = v97
                                                                                                                }
                                                                                                            }
                                                                                                        case 127:
                                                                                                            fmt.Println(fi46)
                                                                                                        }
                                                                                                    case 127:
                                                                                                        p0 -= p0
                                                                                                    }
                                                                                                }
                                                                                                _ = v91
                                                                                            } else {
                                                                                                fmt.Println(fi47)
                                                                                            }
                                                                                        }
                                                                                    case 361:
                                                                                        switch -389 {
                                                                                        case 32768:
                                                                                            switch (("hello" + ("a" + "\n\t")) + ("" + "foo bar")) {
                                                                                            case "a":
                                                                                                var v98 int16 = 0
                                                                                                _ = v98
                                                                                            }
                                                                                        case 32767:
                                                                                            v7 = (v7 * (v7 + (v7 % (((v7 & (v7 + v7))) | 1))))
                                                                                            fmt.Println(p0, fi46, fi46)
                                                                                            if (910 < 2) {
                                                                                                v7 = 1
                                                                                            } else {
                                                                                                if true {
                                                                                                    if false {
                                                                                                        if (584 <= (p0 - p0)) {
                                                                                                            p0 = 1
                                                                                                            if ((-(p0)) != p0) {
                                                                                                                v82 = append(v82, (fi47 * fi47))
                                                                                                            } else {
                                                                                                                _ = len(v82)
                                                                                                                if (!(true && true)) {
                                                                                                                    fmt.Println(p0, p0)
                                                                                                                    switch ("a" + ("\n\t" + ((("hello" + "hello") + "\n\t") + ""))) {
                                                                                                                    case "":
                                                                                                                        if true {
                                                                                                                            var v99 bool = true
                                                                                                                            var v100 int8 = -1
                                                                                                                            _ = v100
                                                                                                                            _ = v99
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        var v101 uint = 65535
                                                                                                                        _ = v101
                                                                                                                    case "a":
                                                                                                                        var v102 uint64 = 0
                                                                                                                        _ = v102
                                                                                                                    }
                                                                                                                }
                                                                                                                v82 = append(v82, (fi47 & (fi46 + -9223372036854775808)))
                                                                                                            }
                                                                                                            fmt.Println(fi47)
                                                                                                            fmt.Println(fi46)
                                                                                                            var v103 bool = (!(!false))
                                                                                                            v104 := GenS4{F0: 9223372036854775807}
                                                                                                            _ = v103
                                                                                                            _ = v104
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch fi47 {
                                                                                                        case -2147483648:
                                                                                                            if len(v82) > 0 {
                                                                                                                v82[0] = fi47
                                                                                                            }
                                                                                                        case 0:
                                                                                                            v105 := GenS12{F0: int8(fi47), F1: fi46}
                                                                                                            _ = v105
                                                                                                        default:
                                                                                                            fmt.Println(fi47)
                                                                                                        }
                                                                                                        v7 = (v7 + (-(v7)))
                                                                                                        var v106 uint = 4294967295
                                                                                                        _ = v106
                                                                                                    }
                                                                                                }
                                                                                                if (!(42 != 174)) {
                                                                                                    var v107 bool = false
                                                                                                    v63 = (v7 - 0)
                                                                                                    _ = v107
                                                                                                } else {
                                                                                                    switch fi46 {
                                                                                                    case 9223372036854775807:
                                                                                                        var v108 int64 = -839
                                                                                                        _ = v108
                                                                                                    case -2147483648:
                                                                                                        v82 = append(v82, fi46)
                                                                                                    case 128:
                                                                                                        if (220 < 255) {
                                                                                                            p0 = p0
                                                                                                        } else {
                                                                                                            if len(v82) > 0 {
                                                                                                                v82[0] = fi47
                                                                                                            }
                                                                                                            var v109 int8 = 100
                                                                                                            v63 = (^(v63))
                                                                                                            p0 += (p0 << ((600) & 63))
                                                                                                            v109 += (v109 & v109)
                                                                                                            v63 -= (v63 + -866)
                                                                                                            _ = v109
                                                                                                        }
                                                                                                    default:
                                                                                                        var v110 uint = 2
                                                                                                        _ = v110
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if true {
                                                                                                var v111 uint16 = 65535
                                                                                                _ = v111
                                                                                            } else {
                                                                                                v82 = append(v82, 4294967295)
                                                                                            }
                                                                                        case -908:
                                                                                            fmt.Println(v7, fi46, v63)
                                                                                        }
                                                                                    case -2147483648:
                                                                                        switch ("\n\t" + "\n\t") {
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                v112 := GenS13{F0: 869, F1: (fi47 / ((fi46) | 1)), F2: ((("hello" + ("\n\t" + ("\n\t" + ("\n\t" + (("a" + "\n\t") + "a"))))) + "\n\t") + "")}
                                                                                                v113 := GenS14{F0: 430, F1: fi46, F2: (v63 + 9223372036854775807)}
                                                                                                if ((true && false) || (!((("foo bar" + "a") < "\n\t") || false))) {
                                                                                                    p0 = (p0 >> ((42) & 63))
                                                                                                } else {
                                                                                                    var v114 int8 = -1
                                                                                                    _ = v114
                                                                                                }
                                                                                                if len(v82) >= 2 {
                                                                                                    v82 = v82[1:len(v82)]
                                                                                                }
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            } else {
                                                                                                var v115 int32 = 1
                                                                                                fmt.Println(v63, v63)
                                                                                                v82 = append(v82, fi47)
                                                                                                _ = v115
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v116 int = (fi46 * (fi47 * fi47))
                                                                                        _ = v116
                                                                                    }
                                                                                    if len(v82) >= 2 {
                                                                                        v82 = v82[1:len(v82)]
                                                                                    }
                                                                                    var v117 int = fi47
                                                                                    _ = v117
                                                                                } else {
                                                                                    var v118 uint64 = 2
                                                                                    _ = v118
                                                                                }
                                                                            }
                                                                            switch (fi46 ^ (fi47 | 106)) {
                                                                            case 42:
                                                                                fmt.Println(p0)
                                                                            case 1:
                                                                                var v119 uint = 837
                                                                                _ = v119
                                                                            case 128:
                                                                                v7 += v63
                                                                            }
                                                                        }
                                                                        p0 = (p0 / ((81) | 1))
                                                                    } else {
                                                                        v7 = (v63 << ((42) & 63))
                                                                        if len(v82) > 0 {
                                                                            v82[0] = (fi47 - (fi47 >> ((18446744073709551615) & 63)))
                                                                        }
                                                                        fmt.Println(v7)
                                                                    }
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                                v63 = (^(v7))
                                                            }
                                                        default:
                                                            v63 = v63
                                                        }
                                                    }
                                                    p0 -= 0
                                                    if false {
                                                        v120 := GenS15{F0: -2147483648, F1: 111, F2: 2147483647}
                                                        var v121 string = "hello"
                                                        var v122 int = fi46
                                                        switch "a" {
                                                        case "":
                                                            var v123 int16 = (p0 - (^(p0)))
                                                            _ = v123
                                                        case "\n\t":
                                                            v124 := []int{(fi47 ^ (fi46 - -931))}
                                                            _ = v124
                                                        default:
                                                            var v125 int32 = -2147483648
                                                            _ = v125
                                                        }
                                                        var v126 string = v121
                                                        _ = v120
                                                        _ = v121
                                                        _ = v122
                                                        _ = v126
                                                    } else {
                                                        v63 = (v63 * 42)
                                                        switch fi46 {
                                                        case 2147483647:
                                                            if false {
                                                                if (1 > 678) {
                                                                    v127 := GenS11{F0: 315, F1: (307 > 639)}
                                                                    fmt.Println(fi46)
                                                                    var v128 int = fi47
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                                switch 45 {
                                                                case 128:
                                                                    v7 += 9223372036854775807
                                                                case -124:
                                                                    if (1 != 65535) {
                                                                        v129 := []int{fi47, fi47}
                                                                        _ = v129
                                                                    } else {
                                                                        var v130 int = fi47
                                                                        _ = v130
                                                                    }
                                                                case 127:
                                                                    v7 = 9223372036854775807
                                                                }
                                                                v63 = 42
                                                            } else {
                                                                var v131 uint8 = 1
                                                                if (3.14159 != ((float32(-2.25) + (0.0 * float32(0.0))) - ((0.1 * 0.1) / 3.14159))) {
                                                                    var v132 uint32 = 4294967295
                                                                    _ = v132
                                                                } else {
                                                                    var v133 uint64 = 999
                                                                    var v134 float32 = (1e10 + float32((3.14159 - (1.5 - ((1.5 - (1e10 + 3.14159)) / 0.1)))))
                                                                    _ = v133
                                                                    _ = v134
                                                                }
                                                                if (false || false) {
                                                                    v135 := GenS8{F0: v131, F1: (p0 * (-(p0)))}
                                                                    v136 := GenS16{F0: 2, F1: -1}
                                                                    _ = v135
                                                                    _ = v136
                                                                } else {
                                                                    v137 := GenS17{F0: 116}
                                                                    _ = v137
                                                                }
                                                                _ = v131
                                                            }
                                                        case 674:
                                                            v138 := GenS18{F0: v7, F1: -128, F2: 1000}
                                                            _ = v138
                                                        case -9223372036854775808:
                                                            var v139 int16 = 117
                                                            _ = v139
                                                        }
                                                        fmt.Println(v63)
                                                        switch "foo bar" {
                                                        case "":
                                                            fmt.Println(v7, p0)
                                                        case "\n\t":
                                                            v140 := GenS14{F0: 18446744073709551615, F1: fi46, F2: (v7 % ((-9223372036854775808) | 1))}
                                                            _ = v140
                                                        default:
                                                            v141 := []int{fi47, fi47, 2147483647}
                                                            v7 = (v7 / (((v63 & (v63 % ((v63) | 1)))) | 1))
                                                            _ = v141
                                                        }
                                                        var v142 float64 = float64(3.14159)
                                                        _ = v142
                                                    }
                                                }
                                                _ = v63
                                            }
                                        }
                                    case -2147483648:
                                        v143 := []int{(fi47 * fi46), (fi47 << ((172) & 63))}
                                        _ = v143
                                    }
                                }
                                var v144 int8 = 100
                                _ = fi47
                                _ = v144
                                _ = v48
                            }
                            v145 := []int{fi46, 0, fi46, 846}
                            _ = fi46
                            _ = v145
                        }
                        for fi146 := 0; fi146 < 11; fi146++ {
                            if true {
                                switch "\n\t" {
                                case "":
                                    if false {
                                        p0 = (^(p0))
                                        for fi147 := 0; fi147 < 14; fi147++ {
                                            v7 = -191
                                            var v148 int = (fi146 % ((128) | 1))
                                            if ((v7 ^ (v7 + v7)) != v7) {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    fmt.Println(v148, fi147, v7)
                                                    switch fi146 {
                                                    case 32767:
                                                        v148 = fi146
                                                    case -1:
                                                        if (42 == 4294967295) {
                                                            var v149 int16 = 32767
                                                            if (!true) {
                                                                if (!(-128 != -1)) {
                                                                    var v150 uint64 = 42
                                                                    if true {
                                                                        v151 := []int{v148, fi147, (fi147 << ((349) & 63)), v148}
                                                                        v149 = v149
                                                                        var v152 int32 = 1
                                                                        _ = v151
                                                                        _ = v152
                                                                    } else {
                                                                        v7 -= 0
                                                                        v153 := GenS19{F0: (fi147 & (fi147 % (((-(v148))) | 1))), F1: "\n\t", F2: 2}
                                                                        _ = v153
                                                                    }
                                                                    _ = v150
                                                                } else {
                                                                    v7 += v7
                                                                    v154 := []int{fi147, 9223372036854775807, (fi146 - (v148 - fi146)), 128}
                                                                    fmt.Println(v149, v149)
                                                                    var v155 uint8 = 1
                                                                    p0 = (v149 / ((v149) | 1))
                                                                    _ = v154
                                                                    _ = v155
                                                                }
                                                                fmt.Println(v7, v148)
                                                            } else {
                                                                v156 := GenS4{F0: 657}
                                                                _ = v156
                                                            }
                                                            var v157 uint64 = 1
                                                            _ = v149
                                                            _ = v157
                                                        }
                                                    }
                                                case "\n\t":
                                                    if true {
                                                        p0 = (p0 >> ((937) & 63))
                                                        var v158 int8 = 127
                                                        _ = v158
                                                    } else {
                                                        var v159 string = ("\n\t" + "foo bar")
                                                        _ = v159
                                                    }
                                                case "":
                                                    v160 := GenS20{F0: (p0 & 179), F1: (-(p0)), F2: (v148 << ((340) & 63))}
                                                    _ = v160
                                                default:
                                                    v161 := GenS18{F0: v7, F1: -9, F2: p0}
                                                    _ = v161
                                                }
                                            } else {
                                                v162 := GenS20{F0: 42, F1: p0, F2: v148}
                                                var v163 int8 = -1
                                                var v164 uint8 = 255
                                                _ = v162
                                                _ = v163
                                                _ = v164
                                            }
                                            _ = fi147
                                            _ = v148
                                        }
                                    } else {
                                        var v165 int16 = p0
                                        if (!false) {
                                            if false {
                                                v166 := []int{(fi146 + (fi146 + (fi146 >> ((65535) & 63)))), -755}
                                                fmt.Println(v165)
                                                fmt.Println(p0, fi146, v165)
                                                _ = v166
                                            } else {
                                                v167 := GenS21{F0: 0, F1: true, F2: 255}
                                                p0 += 1000
                                                for fi168 := 0; fi168 < 10; fi168++ {
                                                    var v169 bool = false
                                                    v170 := GenS22{F0: 65535, F1: 128}
                                                    _ = fi168
                                                    _ = v169
                                                    _ = v170
                                                }
                                                _ = v167
                                            }
                                            if false {
                                                v7 = v7
                                            } else {
                                                v171 := GenS23{F0: fi146}
                                                fmt.Println(v165, p0)
                                                _ = v171
                                            }
                                            var v172 uint16 = 65535
                                            _ = v172
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                v165 += (-(p0))
                                            case "foo bar":
                                                var v173 int16 = 32767
                                                _ = v173
                                            default:
                                                v174 := GenS11{F0: -2147483648, F1: ((4294967295 > 1) && false)}
                                                _ = v174
                                            }
                                        }
                                        var v175 float32 = 0.0
                                        _ = v165
                                        _ = v175
                                    }
                                case "foo bar":
                                    var v176 uint16 = 0
                                    _ = v176
                                }
                                fmt.Println(fi146, v7)
                                switch fi146 {
                                case 958:
                                    var v177 float64 = (3.14159 * ((1e10 + 0.1) + 1e10))
                                    _ = v177
                                case -872:
                                    v178 := GenS11{F0: 976, F1: false}
                                    _ = v178
                                case -1:
                                    var v179 uint16 = 255
                                    _ = v179
                                }
                            } else {
                                var v180 string = ""
                                v181 := []int{(fi146 % ((32767) | 1))}
                                _ = v180
                                _ = v181
                            }
                            var v182 int32 = -2147483648
                            _ = fi146
                            _ = v182
                        }
                    }
                } else {
                    var v183 float32 = 3.14159
                    v184 := []int{128, -9223372036854775808, 128, -625}
                    v7 -= (v7 - (v7 % ((v7) | 1)))
                    _ = v183
                    _ = v184
                }
            }
            if (true && true) {
                switch -2147483648 {
                case 9223372036854775807:
                    if (((1 >= -565) && false) || (127 > 1)) {
                        v185 := []int{-447, -362, 343}
                        var v186 string = "a"
                        _ = v185
                        _ = v186
                    }
                }
                v187 := []int{}
                v188 := GenS24{F0: -2147483648, F1: 0, F2: false}
                _ = v187
                _ = v188
            }
            if (!(2147483647 <= 1)) {
                switch 4294967295 {
                case 32768:
                    v189 := []int{}
                    _ = v189
                default:
                    v7 = v7
                }
            } else {
                v190 := []int{-465, 42}
                v191 := GenS25{F0: 2}
                var v192 uint16 = 255
                _ = v190
                _ = v191
                _ = v192
            }
        }
    default:
        p0 -= (p0 / ((-1) | 1))
    }
    _ = p0
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int32 = -1
    v4 := []uint32{}
    v5 := []int{-9223372036854775808, -725, 1}
    var v6 uint16 = 330
    for fi193 := 0; fi193 < 18; fi193++ {
        delete(v1, int(v3))
        v3 = (v3 & (v3 | v3))
        switch "a" {
        case "hello":
            var v194 float64 = -2.25
            _ = v194
        case "foo bar":
            v195 := GenS24{F0: v3, F1: 32767, F2: true}
            _ = v195
        case "\n\t":
            v3 += (v3 + (v3 + v3))
        }
        _ = fi193
    }
    v3 = (v3 >> ((546) & 63))
    var v196 bool = false
    var v197 string = "a"
    if v196 {
        var v198 uint64 = 0
        for fi199 := 0; fi199 < 17; fi199++ {
            var v200 int64 = -422
            v201 := GenS3{F0: 1000}
            v6 = v6
            _ = fi199
            _ = v200
            _ = v201
        }
        _ = v198
    } else {
        var v202 bool = v196
        var v203 float64 = 1.5
        switch -2147483648 {
        case 32767:
            v204 := GenS26{F0: 766}
            _ = v204
        case 127:
            v202 = false
        case -712:
            switch "a" {
            case "foo bar":
                v205 := GenS27{F0: 127, F1: 100}
                _ = v205
            case "\n\t":
                if false {
                    delete(v2, 1)
                    v203 -= (v203 * (v203 - v203))
                } else {
                    switch "hello" {
                    case "foo bar":
                        switch ("\n\t" + v197) {
                        case "a":
                            for fi206 := 0; fi206 < 10; fi206++ {
                                if true {
                                    fmt.Println(v197, v202, v3)
                                    fmt.Println(fi206, v3, v203)
                                    if v196 {
                                        var v207 float64 = ((1e10 + 0.0) / 1e10)
                                        fmt.Println(v196)
                                        switch ("hello" + v197) {
                                        case "foo bar":
                                            _ = len(v2)
                                        case "":
                                            v196 = v202
                                        default:
                                            var v208 int = (fi206 % (((fi206 & 0)) | 1))
                                            _ = v208
                                        }
                                        _ = v207
                                    } else {
                                        fmt.Println(v197, v197, v3)
                                        if v196 {
                                            v209 := v2[1]
                                            v197 = v197
                                            _ = v209
                                        }
                                    }
                                }
                                _ = fi206
                            }
                        case "foo bar":
                            var v210 uint32 = 42
                            _ = v210
                        case "":
                            var v211 float64 = v203
                            _ = v211
                        }
                    }
                    for fi212 := 0; fi212 < 14; fi212++ {
                        var v213 float32 = 3.14159
                        switch -1 {
                        case -2147483648:
                            var v214 int16 = -570
                            _ = v214
                        }
                        var v215 uint8 = 127
                        _ = fi212
                        _ = v213
                        _ = v215
                    }
                }
            }
        default:
            v216 := GenS28{F0: v3, F1: 616, F2: 151}
            _ = v216
        }
        _ = v202
        _ = v203
    }
    if (-9223372036854775808 == -1) {
        _ = len(v4)
        var v217 uint64 = 482
        _ = v217
    }
    v218 := GenS7{F0: v196}
    switch "\n\t" {
    case "hello":
        v196 = v196
    case "":
        var v219 int = 4294967295
        _ = v219
    }
    v6 += 32767
    if true {
        var v220 int = -2147483648
        fmt.Println(v197, v3)
        _ = v220
    }
    v6 += v6
    v197 = v197
    v221 := GenS9{F0: 245, F1: 933}
    if v196 {
        if v196 {
            v222 := GenS16{F0: 2, F1: v3}
            _ = v222
        } else {
            _ = len(v4)
        }
    } else {
        fmt.Println(v196)
        if v196 {
            v3 = (^(v3))
            for fi223 := 0; fi223 < 3; fi223++ {
                switch (v197 + "\n\t") {
                case "hello":
                    var v224 int = fi223
                    _ = v224
                }
                _ = fi223
            }
        }
        if (!(2147483647 != -1)) {
            for fi225 := 0; fi225 < 16; fi225++ {
                var v226 int8 = 100
                v227 := GenS29{F0: 32767, F1: 168}
                v228 := GenS23{F0: fi225}
                _ = fi225
                _ = v226
                _ = v227
                _ = v228
            }
        } else {
            var v229 uint16 = v6
            _ = v229
        }
    }
    v197 = "a"
    if false {
        v230 := GenS11{F0: v3, F1: ((1.5 + 1.5) > 0.0)}
        v4 = append(v4, uint32(v6))
        switch 1 {
        case -2147483648:
            switch "" {
            case "foo bar":
                var v231 int64 = -727
                _ = v231
            case "\n\t":
                v3 -= (v3 + -2147483648)
            default:
                if (-381 <= v3) {
                    for fi232 := 0; fi232 < 16; fi232++ {
                        if v196 {
                            v233 := GenS4{F0: -1}
                            _ = v233
                        } else {
                            for fi234 := 0; fi234 < 6; fi234++ {
                                v235 := GenS30{F0: 672, F1: 1, F2: ((fi232 / ((-1) | 1)) != 32767)}
                                v3 = 2147483647
                                var v236 uint16 = (v6 - v6)
                                _ = fi234
                                _ = v235
                                _ = v236
                            }
                            delete(v2, 700)
                        }
                        if (int16(v3) == -1) {
                            v6 += 42
                            fmt.Println(v6)
                        }
                        _ = fi232
                    }
                    for fi237 := 0; fi237 < 13; fi237++ {
                        v3 = 1
                        fmt.Println(v196, v196, fi237)
                        var v238 uint32 = 42
                        _ = fi237
                        _ = v238
                    }
                } else {
                    if v196 {
                        switch v197 {
                        case "hello":
                            v239 := GenS20{F0: 170, F1: 1000, F2: -1}
                            _ = v239
                        }
                        v3 = (v3 & (v3 >> ((255) & 63)))
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch v197 {
                    case "hello":
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = 931
                            }
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    case "a":
                        v196 = v196
                    }
                }
            }
        case 646:
            v240 := GenS31{F0: v197, F1: 2147483647, F2: -78}
            _ = v240
        }
        _ = v230
    } else {
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        var v241 uint = 206
        _ = v241
    }
    if (2147483647 != 1) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v242 int = 42
        var v243 uint32 = 2147483647
        _ = v242
        _ = v243
    } else {
        v3 += int32(v6)
        switch -9223372036854775808 {
        case 127:
            _ = len(v5)
        default:
            v6 += v6
        }
    }
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v196:", v196)
    fmt.Println("v197:", v197)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v196
    _ = v197
    _ = v2
    _ = v218
    _ = v221
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
