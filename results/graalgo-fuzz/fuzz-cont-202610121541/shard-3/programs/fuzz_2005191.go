package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int8; F1 uint8; F2 int }
type GenS6 struct { F0 string; F1 int; F2 uint32 }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 uint; F1 string; F2 uint8 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 int32; F1 int64 }
type GenS12 struct { F0 uint32; F1 uint32; F2 int }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool; F1 int16 }
type GenS15 struct { F0 int32; F1 string }
type GenS16 struct { F0 bool; F1 int }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint; F1 int8 }
type GenS19 struct { F0 uint32 }
func genFn0(p0 uint8) uint {
    fmt.Println(p0)
    p0 = (p0 ^ (p0 << ((1) & 63)))
    if (0 > 1) {
        v4 := []int{673, 0, 32768}
        _ = v4
    }
    p0 = (p0 ^ p0)
    _ = p0
    return 251
}

func main() {
    v1 := []uint32{1, 660, 767, 2}
    var v2 uint32 = 921
    var v3 uint16 = 65535
    fmt.Println(v3, v2)
    fmt.Println(v2, v3)
    v2 = 2
    if (-1 <= -1) {
        switch "\n\t" {
        case "foo bar":
            for fi5 := 0; fi5 < 12; fi5++ {
                var v6 string = ""
                _ = len(v1)
                _ = fi5
                _ = v6
            }
        case "a":
            for fi7 := 0; fi7 < 7; fi7++ {
                v3 = 0
                v1 = append(v1, v2)
                _ = fi7
            }
        }
        var v8 int16 = 1000
        _ = v8
    } else {
        if true {
            if true {
                var v9 int = -2147483648
                _ = v9
            }
            var v10 float32 = 0.0
            v10 = ((v10 / (v10 - (v10 - v10))) / ((1.5 * (v10 * (0.0 * v10))) / v10))
            _ = v10
        } else {
            v3 = 65535
            switch 4294967295 {
            case 42:
                v11 := GenS1{F0: genFn0(25)}
                _ = v11
            default:
                v12 := GenS2{F0: 42}
                _ = v12
            }
            var v13 uint16 = (v3 * (^(v3)))
            _ = v13
        }
        fmt.Println(v2)
        if (!true) {
            for fi14 := 0; fi14 < 11; fi14++ {
                fmt.Println(v2, v3, fi14)
                fmt.Println(v2)
                v2 = v2
                _ = fi14
            }
        } else {
            switch 898 {
            case -202:
                var v15 int16 = 32767
                _ = v15
            case 9223372036854775807:
                var v16 string = "a"
                _ = v16
            case 109:
                if true {
                    v2 = 406
                    v17 := GenS3{F0: 100, F1: ("\n\t" + "\n\t")}
                    v3 = v3
                    _ = v17
                } else {
                    v3 = v3
                    if true {
                        switch 0 {
                        case 2147483647:
                            switch "\n\t" {
                            case "":
                                var v18 int16 = -32768
                                _ = v18
                            case "a":
                                _ = len(v1)
                            case "foo bar":
                                if ((((0.1 * 1.5) / 1.5) + 3.14159) == (1e10 + (0.0 * 1e10))) {
                                    if (42 < (^(v3))) {
                                        var v19 uint32 = 932
                                        v3 = (v3 * 567)
                                        _ = v19
                                    } else {
                                        var v20 uint8 = 145
                                        if false {
                                            v2 += (v2 * (^(v2)))
                                        }
                                        if true {
                                            var v21 int = 0
                                            for fi22 := 0; fi22 < 4; fi22++ {
                                                var v23 uint = 935
                                                switch 4294967295 {
                                                case 750:
                                                    if ("\n\t" > "foo bar") {
                                                        switch "foo bar" {
                                                        case "a":
                                                            if (v2 >= (v2 | 836)) {
                                                                v24 := GenS4{F0: v20, F1: 0}
                                                                for fi25 := 0; fi25 < 9; fi25++ {
                                                                    var v26 uint = v23
                                                                    _ = fi25
                                                                    _ = v26
                                                                }
                                                                switch fi22 {
                                                                case 32767:
                                                                    v27 := GenS5{F0: 1, F1: (v20 | 1), F2: (fi22 + (fi22 & 0))}
                                                                    _ = v27
                                                                default:
                                                                    for fi28 := 0; fi28 < 7; fi28++ {
                                                                        var v29 int16 = 731
                                                                        v30 := GenS6{F0: ("\n\t" + "\n\t"), F1: (fi22 & v21), F2: uint32(v23)}
                                                                        _ = fi28
                                                                        _ = v29
                                                                        _ = v30
                                                                    }
                                                                }
                                                                _ = v24
                                                            }
                                                        case "foo bar":
                                                            for fi31 := 0; fi31 < 16; fi31++ {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                var v32 uint16 = 0
                                                                v33 := GenS7{F0: -1, F1: v23}
                                                                _ = fi31
                                                                _ = v32
                                                                _ = v33
                                                            }
                                                        case "":
                                                            for fi34 := 0; fi34 < 10; fi34++ {
                                                                v2 = v2
                                                                if len(v1) > 0 {
                                                                    v1[0] = (v2 % ((v2) | 1))
                                                                }
                                                                if true {
                                                                    switch (-(fi22)) {
                                                                    case -1:
                                                                        v35 := GenS8{F0: genFn0(v20), F1: "hello", F2: v20}
                                                                        _ = v35
                                                                    default:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (v2 - (v2 / ((uint32(v23)) | 1)))
                                                                        }
                                                                    }
                                                                } else {
                                                                    switch (fi22 & int(v3)) {
                                                                    case 42:
                                                                        if (1 >= 138) {
                                                                            v3 -= 32767
                                                                            v36 := GenS2{F0: 127}
                                                                            _ = v36
                                                                        }
                                                                    case 9223372036854775807:
                                                                        if false {
                                                                            if (v20 < v20) {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (v2 << ((v23) & 63))
                                                                                }
                                                                                var v37 uint8 = v20
                                                                                v2 = v2
                                                                                fmt.Println(v2, v2, v20)
                                                                                var v38 int64 = 9223372036854775807
                                                                                _ = v37
                                                                                _ = v38
                                                                            }
                                                                            switch ("a" + ("foo bar" + "hello")) {
                                                                            case "":
                                                                                v21 -= (^(fi22))
                                                                            default:
                                                                                var v39 uint = (v23 - (v23 >> (((v23 >> ((v23) & 63))) & 63)))
                                                                                _ = v39
                                                                            }
                                                                        } else {
                                                                            var v40 bool = ((32767 <= 42) || ((-434 == -1) || (!false)))
                                                                            v1 = append(v1, (v2 << (((v23 | v23)) & 63)))
                                                                            _ = v40
                                                                        }
                                                                    case 1:
                                                                        v21 = 153
                                                                    }
                                                                    v41 := GenS9{F0: (v21 / ((fi22) | 1)), F1: 1}
                                                                    _ = v41
                                                                }
                                                                _ = fi34
                                                            }
                                                        }
                                                    } else {
                                                        switch (fi22 & (v21 << (((v23 % ((v23) | 1))) & 63))) {
                                                        case -2147483648:
                                                            for fi42 := 0; fi42 < 7; fi42++ {
                                                                if false {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = v2
                                                                    }
                                                                    v43 := GenS6{F0: "\n\t", F1: 127, F2: 1}
                                                                    v20 = (v20 & (v20 - v20))
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v20 = (v20 & 127)
                                                                    default:
                                                                        v44 := GenS1{F0: v23}
                                                                        _ = v44
                                                                    }
                                                                    v2 = 1
                                                                    v2 = (v2 - (^(v2)))
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v1 = append(v1, v2)
                                                                    _ = v43
                                                                }
                                                                if true {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v45 uint64 = 2
                                                                        _ = v45
                                                                    }
                                                                    fmt.Println(fi42, v20)
                                                                    v21 = 42
                                                                } else {
                                                                    var v46 uint32 = 2147483647
                                                                    var v47 int8 = 1
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                                if (!((v3 <= (v3 / (((v3 << (((v23 - 2)) & 63))) | 1))) && true)) {
                                                                    v20 = v20
                                                                    if false {
                                                                        var v48 string = "\n\t"
                                                                        var v49 string = v48
                                                                        v20 -= (v20 - (^(v20)))
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                }
                                                                _ = fi42
                                                            }
                                                        case 127:
                                                            var v50 int8 = 0
                                                            _ = v50
                                                        case 543:
                                                            for fi51 := 0; fi51 < 15; fi51++ {
                                                                _ = len(v1)
                                                                if (!false) {
                                                                    switch -827 {
                                                                    case 430:
                                                                        v52 := GenS10{F0: v2, F1: 42}
                                                                        _ = v52
                                                                    case 9223372036854775807:
                                                                        v3 = v3
                                                                    case 434:
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            var v53 float32 = 1.5
                                                                            _ = v53
                                                                        case "":
                                                                            fmt.Println(v23, fi22, v2)
                                                                            var v54 uint = genFn0(v20)
                                                                            var v55 uint32 = (^(v2))
                                                                            _ = v54
                                                                            _ = v55
                                                                        default:
                                                                            _ = len(v1)
                                                                        }
                                                                    default:
                                                                        _ = len(v1)
                                                                    }
                                                                    var v56 int = fi51
                                                                    if (true || (9223372036854775807 == 9223372036854775807)) {
                                                                        fmt.Println(v23, v56, fi22)
                                                                        if false {
                                                                            var v57 uint64 = uint64(v23)
                                                                            var v58 int32 = 42
                                                                            _ = v57
                                                                            _ = v58
                                                                        }
                                                                        switch ("a" + ("" + (("" + "") + ("a" + ("\n\t" + ("foo bar" + ("hello" + ("foo bar" + "")))))))) {
                                                                        case "a":
                                                                            v59 := GenS11{F0: 0, F1: 0}
                                                                            _ = v59
                                                                        case "foo bar":
                                                                            if true {
                                                                                v2 = v2
                                                                                v60 := GenS12{F0: (^(v2)), F1: (v2 >> (((^(v23))) & 63)), F2: (fi22 & 4294967295)}
                                                                                _ = v60
                                                                            } else {
                                                                                var v61 uint64 = 42
                                                                                if true {
                                                                                    var v62 string = "a"
                                                                                    v62 = (("\n\t" + (v62 + ("a" + "a"))) + v62)
                                                                                    _ = v62
                                                                                }
                                                                                if (1.5 == float64((3.14159 * (1.5 - 1e10)))) {
                                                                                    if (!(v20 > v20)) {
                                                                                        switch v21 {
                                                                                        case -170:
                                                                                            switch "\n\t" {
                                                                                            case "foo bar":
                                                                                                v63 := GenS13{F0: v2}
                                                                                                _ = v63
                                                                                            case "a":
                                                                                                fmt.Println(v3)
                                                                                            case "":
                                                                                                if (uint64(v23) >= v61) {
                                                                                                    switch ("foo bar" + "hello") {
                                                                                                    case "\n\t":
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                    }
                                                                                                    v64 := GenS11{F0: 2147483647, F1: -196}
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        if (!((v3 << (((v23 & (^(v23)))) & 63)) > (v3 % (((^(v3))) | 1)))) {
                                                                                                            var v65 float64 = 3.14159
                                                                                                            v20 = (v20 / ((v20) | 1))
                                                                                                            fmt.Println(v23, v56)
                                                                                                            _ = v65
                                                                                                        } else {
                                                                                                            var v66 float32 = 1.5
                                                                                                            _ = v66
                                                                                                        }
                                                                                                    case "a":
                                                                                                        var v67 int64 = -24
                                                                                                        _ = v67
                                                                                                    case "hello":
                                                                                                        var v68 float32 = -2.25
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    var v69 uint8 = v20
                                                                                                    _ = v64
                                                                                                    _ = v69
                                                                                                } else {
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                if (32768 != fi22) {
                                                                                                    switch ("hello" + ("foo bar" + "a")) {
                                                                                                    case "foo bar":
                                                                                                        switch fi22 {
                                                                                                        case 9223372036854775807:
                                                                                                            var v70 uint8 = v20
                                                                                                            _ = v70
                                                                                                        default:
                                                                                                            var v71 string = "foo bar"
                                                                                                            _ = v71
                                                                                                        }
                                                                                                    default:
                                                                                                        if (v23 > v23) {
                                                                                                            v23 += (v23 % (((v23 << ((v23) & 63))) | 1))
                                                                                                            var v72 int16 = 0
                                                                                                            var v73 int32 = 1
                                                                                                            _ = v72
                                                                                                            _ = v73
                                                                                                        }
                                                                                                    }
                                                                                                    var v74 uint32 = uint32(v23)
                                                                                                    _ = v74
                                                                                                }
                                                                                            }
                                                                                        case 304:
                                                                                            fmt.Println(fi22, v21, v3)
                                                                                        case 1:
                                                                                            v75 := GenS14{F0: true, F1: 160}
                                                                                            _ = v75
                                                                                        }
                                                                                        v1 = append(v1, (v2 ^ (v2 >> ((v23) & 63))))
                                                                                        if true {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = (v2 | v2)
                                                                                            }
                                                                                            if (((-2.25 - 3.14159) - 1e10) < ((-2.25 / 1.0) * -2.25)) {
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                v56 += -9223372036854775808
                                                                                                switch 42 {
                                                                                                case -1:
                                                                                                    v23 = v23
                                                                                                case 42:
                                                                                                    var v76 int16 = 1
                                                                                                    _ = v76
                                                                                                case 128:
                                                                                                    fmt.Println(v23, v23, fi51)
                                                                                                default:
                                                                                                    if ("hello" <= ("\n\t" + "\n\t")) {
                                                                                                        v77 := GenS9{F0: 971, F1: -2147483648}
                                                                                                        var v78 uint16 = 32767
                                                                                                        switch (v21 & 191) {
                                                                                                        case -744:
                                                                                                            if (1 < 0) {
                                                                                                                if true {
                                                                                                                    if (!(65535 > v3)) {
                                                                                                                        switch ("hello" + "") {
                                                                                                                        case "a":
                                                                                                                            v79 := GenS14{F0: (!false), F1: -32768}
                                                                                                                            _ = v79
                                                                                                                        case "":
                                                                                                                            v23 = (v23 | genFn0((v20 + 1)))
                                                                                                                        case "hello":
                                                                                                                            if ("\n\t" == "hello") {
                                                                                                                                fmt.Println(v78, v78, v3)
                                                                                                                                var v80 uint32 = v2
                                                                                                                                var v81 uint = 255
                                                                                                                                _ = v80
                                                                                                                                _ = v81
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if len(v1) >= 2 {
                                                                                                                            v1 = v1[1:len(v1)]
                                                                                                                        }
                                                                                                                        v82 := GenS5{F0: 19, F1: (v20 / ((v20) | 1)), F2: v21}
                                                                                                                        _ = v82
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            v2 = uint32(v23)
                                                                                                                            var v83 float64 = float64(0.1)
                                                                                                                            _ = v83
                                                                                                                        } else {
                                                                                                                            if (!false) {
                                                                                                                                v84 := GenS15{F0: 42, F1: ""}
                                                                                                                                switch "\n\t" {
                                                                                                                                case "":
                                                                                                                                    if len(v1) > 0 {
                                                                                                                                        v1[0] = (v2 << ((v23) & 63))
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    var v85 int8 = 74
                                                                                                                                    _ = v85
                                                                                                                                }
                                                                                                                                var v86 int32 = 42
                                                                                                                                _ = v84
                                                                                                                                _ = v86
                                                                                                                            }
                                                                                                                            v87 := GenS2{F0: 1}
                                                                                                                            _ = v87
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v88 string = ((((("\n\t" + ("a" + "")) + "foo bar") + "hello") + ("hello" + "\n\t")) + "")
                                                                                                                    v78 = (^(v78))
                                                                                                                    _ = v88
                                                                                                                } else {
                                                                                                                    var v89 uint32 = v2
                                                                                                                    _ = v89
                                                                                                                }
                                                                                                            }
                                                                                                        case 128:
                                                                                                            _ = len(v1)
                                                                                                        case 32767:
                                                                                                            v2 = 4294967295
                                                                                                        default:
                                                                                                            switch ("a" + "\n\t") {
                                                                                                            case "a":
                                                                                                                v90 := GenS15{F0: 42, F1: (("a" + "") + "")}
                                                                                                                _ = v90
                                                                                                            case "":
                                                                                                                v61 = 855
                                                                                                            default:
                                                                                                                var v91 uint64 = v61
                                                                                                                _ = v91
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v77
                                                                                                        _ = v78
                                                                                                    } else {
                                                                                                        if ((((v61 + (v61 << ((961) & 63))) != v61) || true) || false) {
                                                                                                            if true {
                                                                                                                if len(v1) > 0 {
                                                                                                                    v1[0] = (v2 & v2)
                                                                                                                }
                                                                                                                if (100 != 119) {
                                                                                                                    fmt.Println(v23)
                                                                                                                    v92 := GenS16{F0: (-516 >= 9223372036854775807), F1: v56}
                                                                                                                    v23 -= v23
                                                                                                                    _ = v92
                                                                                                                } else {
                                                                                                                    var v93 int = v56
                                                                                                                    v61 = (v61 >> ((v23) & 63))
                                                                                                                    _ = v93
                                                                                                                }
                                                                                                                v23 = uint(v3)
                                                                                                            }
                                                                                                            switch ("a" + "hello") {
                                                                                                            case "hello":
                                                                                                                var v94 int16 = 1
                                                                                                                _ = v94
                                                                                                            }
                                                                                                        }
                                                                                                        var v95 bool = (v20 != (^(v20)))
                                                                                                        _ = v95
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v2, v20)
                                                                                                _ = len(v1)
                                                                                                switch fi51 {
                                                                                                case 42:
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                default:
                                                                                                    v2 += v2
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v96 int8 = -78
                                                                                    v21 += (fi22 << ((v23) & 63))
                                                                                    _ = v96
                                                                                }
                                                                                var v97 uint64 = (v61 | uint64(v23))
                                                                                var v98 int8 = 1
                                                                                _ = v61
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                        case "\n\t":
                                                                            var v99 uint32 = uint32(v3)
                                                                            _ = v99
                                                                        default:
                                                                            if false {
                                                                                if false {
                                                                                    if ((!false) && (!(false && false))) {
                                                                                        if (!false) {
                                                                                            v20 = (v20 | (v20 << ((v23) & 63)))
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            var v100 int64 = 518
                                                                                            var v101 float32 = (0.1 + 0.1)
                                                                                            _ = v100
                                                                                            _ = v101
                                                                                        } else {
                                                                                            v2 = (v2 >> (((v23 - (v23 & (v23 & genFn0(v20))))) & 63))
                                                                                            fmt.Println(v20, v3)
                                                                                        }
                                                                                        v2 -= v2
                                                                                    } else {
                                                                                        v2 = v2
                                                                                    }
                                                                                    var v102 uint8 = 127
                                                                                    _ = len(v1)
                                                                                    v23 -= (v23 * (v23 / ((65535) | 1)))
                                                                                    _ = v102
                                                                                }
                                                                                var v103 int16 = 1000
                                                                                _ = v103
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 4294967295
                                                                        }
                                                                        if (1e10 > 0.0) {
                                                                            var v104 uint32 = uint32(v20)
                                                                            v1 = append(v1, (v2 + v2))
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v23 = uint(v3)
                                                                                var v105 string = "foo bar"
                                                                                _ = v105
                                                                            case "a":
                                                                                var v106 bool = false
                                                                                _ = v106
                                                                            case "":
                                                                                switch (("a" + "a") + "hello") {
                                                                                case "\n\t":
                                                                                    var v107 int8 = 42
                                                                                    _ = v107
                                                                                default:
                                                                                    fmt.Println(v21, v3)
                                                                                }
                                                                            }
                                                                            _ = v104
                                                                        }
                                                                        fmt.Println(v21, v21)
                                                                    }
                                                                    _ = v56
                                                                } else {
                                                                    switch ((("hello" + "") + ("hello" + ("" + ""))) + "hello") {
                                                                    case "":
                                                                        v108 := GenS9{F0: fi22, F1: -988}
                                                                        _ = v108
                                                                    case "foo bar":
                                                                        var v109 int32 = 2147483647
                                                                        _ = v109
                                                                    default:
                                                                        _ = len(v1)
                                                                    }
                                                                    var v110 string = "foo bar"
                                                                    _ = v110
                                                                }
                                                                _ = fi51
                                                            }
                                                        default:
                                                            var v111 bool = true
                                                            _ = v111
                                                        }
                                                    }
                                                case 2147483647:
                                                    var v112 int32 = -1
                                                    _ = v112
                                                default:
                                                    v113 := GenS17{F0: 65535}
                                                    _ = v113
                                                }
                                                if (((1.5 / (-2.25 + float32(1.5))) * 3.14159) != ((1.5 - -2.25) + 1e10)) {
                                                    if len(v1) > 0 {
                                                        v1[0] = (v2 & v2)
                                                    }
                                                } else {
                                                    var v114 int16 = -524
                                                    switch v21 {
                                                    case 9223372036854775807:
                                                        var v115 bool = (246 != (v20 - v20))
                                                        _ = v115
                                                    case -2147483648:
                                                        _ = len(v1)
                                                    case 574:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    _ = v114
                                                }
                                                _ = fi22
                                                _ = v23
                                            }
                                            if true {
                                                var v116 int64 = 102
                                                fmt.Println(v3)
                                                var v117 uint64 = 42
                                                _ = v116
                                                _ = v117
                                            } else {
                                                var v118 int64 = 42
                                                _ = v118
                                            }
                                            _ = v21
                                        } else {
                                            fmt.Println(v2)
                                            if (!false) {
                                                v2 += uint32(v3)
                                                switch ((("foo bar" + ("\n\t" + (("hello" + ("foo bar" + (("\n\t" + ("\n\t" + "")) + ("hello" + "hello")))) + "foo bar"))) + "foo bar") + "\n\t") {
                                                case "hello":
                                                    var v119 uint64 = 0
                                                    _ = v119
                                                }
                                            } else {
                                                v120 := GenS8{F0: 448, F1: "a", F2: 3}
                                                _ = v120
                                            }
                                        }
                                        _ = v20
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                    switch 1 {
                                    case 218:
                                        v121 := GenS14{F0: false, F1: 24}
                                        _ = v121
                                    case 0:
                                        v3 += v3
                                    default:
                                        fmt.Println(v3, v3)
                                    }
                                } else {
                                    v122 := GenS18{F0: 726, F1: 100}
                                    fmt.Println(v3, v2)
                                    _ = v122
                                }
                            }
                        case 4294967295:
                            v123 := GenS5{F0: 100, F1: 0, F2: 42}
                            _ = v123
                        case -2147483648:
                            if (255 <= v3) {
                                fmt.Println(v2, v3)
                                v3 -= v3
                            } else {
                                if ((2 < 2) || true) {
                                    if (false || false) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v124 bool = (0.0 <= float32(1.5))
                                        _ = v124
                                    } else {
                                        for fi125 := 0; fi125 < 3; fi125++ {
                                            v126 := GenS5{F0: -128, F1: 128, F2: fi125}
                                            var v127 int = fi125
                                            _ = fi125
                                            _ = v126
                                            _ = v127
                                        }
                                        var v128 bool = true
                                        _ = v128
                                    }
                                    v2 = (v2 / ((42) | 1))
                                }
                                var v129 uint8 = 42
                                var v130 uint64 = 4294967295
                                _ = v129
                                _ = v130
                            }
                        default:
                            for fi131 := 0; fi131 < 5; fi131++ {
                                v3 += (v3 - (v3 % ((v3) | 1)))
                                fmt.Println(v2)
                                _ = fi131
                            }
                        }
                        fmt.Println(v2, v2)
                    } else {
                        var v132 int = 127
                        var v133 bool = false
                        _ = v132
                        _ = v133
                    }
                }
            }
        }
    }
    for fi134 := 0; fi134 < 20; fi134++ {
        var v135 bool = false
        _ = fi134
        _ = v135
    }
    for fi136 := 0; fi136 < 13; fi136++ {
        v2 += v2
        switch "hello" {
        case "a":
            if (true && (!(fi136 <= int(v2)))) {
                fmt.Println(fi136, v3)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (!false) {
                    if false {
                        if (!false) {
                            var v137 string = ("hello" + "a")
                            v2 += (^(v2))
                            v137 = v137
                            _ = v137
                        }
                        v2 = v2
                    }
                    var v138 int32 = 1
                    switch ("" + "hello") {
                    case "a":
                        var v139 uint16 = 336
                        _ = v139
                    case "\n\t":
                        var v140 string = "\n\t"
                        _ = v140
                    case "foo bar":
                        var v141 int = (fi136 ^ 305)
                        _ = v141
                    }
                    _ = v138
                } else {
                    var v142 uint32 = (^(v2))
                    v3 = (v3 << ((1) & 63))
                    if len(v1) > 0 {
                        v1[0] = (v2 >> ((713) & 63))
                    }
                    _ = v142
                }
            }
        case "\n\t":
            v2 = v2
        case "foo bar":
            for fi143 := 0; fi143 < 1; fi143++ {
                v2 -= 2
                var v144 uint32 = v2
                _ = fi143
                _ = v144
            }
        default:
            v3 -= (v3 - 255)
        }
        _ = fi136
    }
    v145 := GenS17{F0: 32767}
    fmt.Println(v2, v2)
    v2 -= v2
    v146 := GenS19{F0: 546}
    var v147 bool = (0.1 != 1.5)
    var v148 uint32 = (^(v2))
    var v149 float32 = 0.1
    v149 -= v149
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v145
    _ = v146
    _ = v147
    _ = v148
    _ = v149
    _ = v2
    _ = v3
}
