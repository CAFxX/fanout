package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 uint16; F1 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int; F1 int8; F2 int16 }
type GenS11 struct { F0 uint16; F1 bool }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32; F1 uint16 }
type GenS15 struct { F0 int16; F1 int32; F2 uint32 }
type GenS16 struct { F0 int8; F1 int16; F2 uint32 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int32; F1 int64 }
type GenS19 struct { F0 uint64; F1 int64 }
type GenS20 struct { F0 bool; F1 bool }
type GenS21 struct { F0 uint32; F1 uint16; F2 int }
type GenS22 struct { F0 uint16; F1 bool }
type GenS23 struct { F0 uint8; F1 uint16 }
type GenS24 struct { F0 uint }
type GenS25 struct { F0 int32; F1 uint }
type GenS26 struct { F0 int64; F1 bool }
type GenS27 struct { F0 uint16; F1 int16 }
type GenS28 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS29 struct { F0 uint }
type GenS30 struct { F0 int16; F1 int8; F2 bool }
type GenS31 struct { F0 string }
type GenS32 struct { F0 string }
type GenS33 struct { F0 int16 }
type GenS34 struct { F0 int64 }
type GenS35 struct { F0 uint8; F1 uint; F2 int32 }

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{}
    v3 := []uint32{2147483648, 66, 2147483648}
    var v4 bool = false
    var v5 int8 = 117
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 14; fi7++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        var v8 uint64 = 1
        v9 := GenS1{F0: 631, F1: 65535}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v5, v1)
    for fi10 := 0; fi10 < 13; fi10++ {
        if v4 {
            v11 := GenS1{F0: 433, F1: 578}
            v1 -= v1
            _ = v11
        } else {
            v1 -= (v1 >> ((255) & 63))
            for fi12 := 0; fi12 < 3; fi12++ {
                v5 = 54
                if (v4 || ((v1 == 373) || ((fi12 <= (-(fi10))) && (32 == 65535)))) {
                    switch fi12 {
                    case 32768:
                        v4 = false
                    case 127:
                        if v4 {
                            v13 := GenS2{F0: (fi10 - -494)}
                            fmt.Println(fi12, fi12)
                            _ = v13
                        } else {
                            switch "" {
                            case "":
                                v14 := GenS1{F0: 65535, F1: 2147483647}
                                _ = v14
                            case "foo bar":
                                v15 := GenS2{F0: (fi10 << ((0) & 63))}
                                _ = v15
                            case "a":
                                switch "\n\t" {
                                case "\n\t":
                                    if v4 {
                                        fmt.Println(fi12)
                                        v16 := GenS1{F0: 18, F1: 651}
                                        _ = v16
                                    }
                                }
                            default:
                                fmt.Println(fi12, fi10, v4)
                            }
                            switch (fi12 % ((fi10) | 1)) {
                            case 32767:
                                if true {
                                    var v17 uint64 = 175
                                    fmt.Println(v5, v17, v5)
                                    if ((!true) || ((v4 && true) || (!true))) {
                                        switch -362 {
                                        case -1:
                                            fmt.Println(fi12, v17, fi12)
                                        }
                                        if v4 {
                                            switch -9223372036854775808 {
                                            case -427:
                                                var v18 int = fi10
                                                _ = v18
                                            case 1:
                                                var v19 int64 = 42
                                                _ = v19
                                            default:
                                                fmt.Println(v4, v4)
                                            }
                                            var v20 int8 = v5
                                            var v21 uint16 = 32767
                                            _ = v20
                                            _ = v21
                                        } else {
                                            v17 += (v17 << ((0) & 63))
                                        }
                                        var v22 uint8 = uint8(v17)
                                        _ = v22
                                    }
                                    _ = v17
                                }
                            case 1:
                                fmt.Println(fi12, v1, fi12)
                            case -1:
                                v1 += -2147483648
                            }
                            if false {
                                var v23 float32 = 3.14159
                                v1 += 1
                                v4 = v4
                                switch "hello" {
                                case "\n\t":
                                    var v24 uint64 = 764
                                    _ = v24
                                }
                                v5 += (v5 * v5)
                                _ = v23
                            }
                            var v25 int32 = v1
                            _ = v25
                        }
                    case -9223372036854775808:
                        v26 := GenS3{F0: 1}
                        _ = v26
                    default:
                        fmt.Println(v5, v4)
                    }
                    switch "foo bar" {
                    case "":
                        switch 960 {
                        case 0:
                            v2 = append(v2, 42)
                        case 4294967295:
                            v1 += v1
                        default:
                            switch "foo bar" {
                            case "foo bar":
                                switch 233 {
                                case 1:
                                    v27 := GenS3{F0: (v5 & (v5 + 127))}
                                    _ = v27
                                case 9223372036854775807:
                                    if v4 {
                                        if true {
                                            fmt.Println(v4, v5)
                                            if (((65535 <= 278) && false) || v4) {
                                                delete(v6, ("\n\t" + "hello"))
                                            }
                                            fmt.Println(v5, v5)
                                        } else {
                                            switch fi12 {
                                            case 0:
                                                var v28 string = ("hello" + "foo bar")
                                                _ = v28
                                            case 924:
                                                v29 := GenS4{F0: false}
                                                _ = v29
                                            }
                                            var v30 int = fi12
                                            fmt.Println(v30, v4, v4)
                                            _ = v30
                                        }
                                    } else {
                                        v4 = (-9223372036854775808 < -9223372036854775808)
                                        fmt.Println(v5, v4, v4)
                                        v4 = (127 < 1)
                                        if true {
                                            _ = len(v3)
                                            var v31 int64 = 0
                                            _ = v31
                                        }
                                    }
                                case 863:
                                    v1 -= v1
                                default:
                                    delete(v6, ("" + "a"))
                                }
                            case "\n\t":
                                var v32 int64 = -1
                                _ = v32
                            default:
                                v5 = (v5 - v5)
                            }
                        }
                    case "foo bar":
                        var v33 uint32 = 65535
                        fmt.Println(v5, v4)
                        _ = v33
                    default:
                        delete(v6, "hello")
                    }
                    if (((("foo bar" + "hello") + "hello") + "hello") != "") {
                        var v34 float64 = (0.0 / (-2.25 + 0.0))
                        _ = v34
                    } else {
                        v5 = (^(v5))
                    }
                }
                switch "foo bar" {
                case "":
                    if v4 {
                        switch "a" {
                        case "\n\t":
                            v5 = (v5 - (v5 & (^(v5))))
                        case "":
                            if v4 {
                                v35 := GenS5{F0: 65535, F1: 42}
                                v36 := GenS4{F0: (3.14159 >= -2.25)}
                                if (true && v4) {
                                    v1 += v1
                                    fmt.Println(v1)
                                    v1 += v1
                                } else {
                                    v37 := GenS1{F0: 0, F1: 2147483647}
                                    _ = v37
                                }
                                switch fi12 {
                                case 32767:
                                    if true {
                                        v1 -= v1
                                        var v38 uint32 = 1
                                        _ = v38
                                    } else {
                                        if v4 {
                                            v6["foo bar"] = 42
                                        }
                                    }
                                case 4294967295:
                                    _ = len(v6)
                                }
                                _ = v35
                                _ = v36
                            }
                        default:
                            v5 -= v5
                        }
                        var v39 uint32 = 42
                        if true {
                            if (42 <= 994) {
                                delete(v6, "a")
                                _ = len(v6)
                            }
                            v5 = v5
                            v5 = 102
                        }
                        _ = v39
                    }
                    v4 = ("a" <= "")
                    if (-891 != 0) {
                        switch "a" {
                        case "foo bar":
                            var v40 int = (^(fi10))
                            _ = v40
                        case "\n\t":
                            v1 = -2147483648
                        case "a":
                            v2 = append(v2, 636)
                        }
                    }
                case "hello":
                    v5 = v5
                case "foo bar":
                    if (1 <= 787) {
                        var v41 int64 = 1
                        _ = len(v6)
                        var v42 float32 = (-2.25 + 1e10)
                        _ = v41
                        _ = v42
                    } else {
                        var v43 int = (fi10 * fi10)
                        switch ("" + "a") {
                        case "hello":
                            v1 = (v1 & (^(v1)))
                        }
                        _ = v43
                    }
                default:
                    switch "foo bar" {
                    case "hello":
                        if v4 {
                            if false {
                                v4 = (!(v4 && false))
                                var v44 uint8 = 1
                                delete(v6, ("a" + "foo bar"))
                                var v45 string = (("hello" + ("a" + ("\n\t" + (("hello" + "\n\t") + "hello")))) + "a")
                                switch "a" {
                                case "foo bar":
                                    switch "hello" {
                                    case "hello":
                                        v45 = v45
                                    case "foo bar":
                                        v44 -= v44
                                    default:
                                        _ = len(v6)
                                        var v46 string = ""
                                        if len(v2) > 0 {
                                            v2[0] = 1
                                        }
                                        switch (-(fi12)) {
                                        case -90:
                                            v45 = v45
                                            v46 = v46
                                            switch "" {
                                            case "\n\t":
                                                switch fi12 {
                                                case 0:
                                                    v47 := GenS5{F0: 0, F1: 0}
                                                    _ = v47
                                                case 2147483647:
                                                    switch int(v44) {
                                                    case 9223372036854775807:
                                                        v48 := GenS6{F0: 65535, F1: 1}
                                                        _ = v48
                                                    default:
                                                        var v49 float32 = float32(0.1)
                                                        _ = v49
                                                    }
                                                case 127:
                                                    if true {
                                                        var v50 float32 = -2.25
                                                        if v4 {
                                                            fmt.Println(v45, v4, v1)
                                                            _ = len(v6)
                                                            _ = len(v3)
                                                            v45 = v46
                                                        } else {
                                                            if true {
                                                                fmt.Println(fi12)
                                                                var v51 int = (-(fi10))
                                                                _ = v51
                                                            } else {
                                                                var v52 int32 = v1
                                                                switch (fi10 >> ((556) & 63)) {
                                                                case 0:
                                                                    if v4 {
                                                                        v5 = v5
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v53 := v6[((v45 + v45) + (v46 + (((("\n\t" + "a") + "") + v45) + ("hello" + v45))))]
                                                                            _ = v53
                                                                        }
                                                                        fmt.Println(v45, v46, v52)
                                                                        switch v46 {
                                                                        case "":
                                                                            fmt.Println(v52, v52, v5)
                                                                            _ = len(v2)
                                                                        }
                                                                    } else {
                                                                        v54 := GenS3{F0: v5}
                                                                        var v55 float64 = 1.5
                                                                        v5 = v5
                                                                        delete(v6, (("\n\t" + ((v45 + v46) + "foo bar")) + ((v45 + "hello") + v46)))
                                                                        _ = v54
                                                                        _ = v55
                                                                    }
                                                                    v6["hello"] = 0
                                                                    switch v46 {
                                                                    case "a":
                                                                        fmt.Println(v52, fi12, v52)
                                                                    case "hello":
                                                                        fmt.Println(v1)
                                                                    case "foo bar":
                                                                        if (651 == 0) {
                                                                            fmt.Println(v52)
                                                                            if ((9223372036854775807 == -9223372036854775808) && ((fi10 >= fi10) && true)) {
                                                                                v56 := GenS1{F0: 576, F1: 65535}
                                                                                var v57 string = v45
                                                                                if (v45 > (v45 + "a")) {
                                                                                    fmt.Println(v5, v4, v44)
                                                                                } else {
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = 42
                                                                                    }
                                                                                }
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                v5 = -128
                                                                            }
                                                                            switch (-(fi12)) {
                                                                            case 4294967295:
                                                                                v58 := GenS4{F0: ((fi12 <= (fi12 ^ (fi10 * 4294967295))) && (-596 <= 948))}
                                                                                _ = v58
                                                                            }
                                                                            v52 = (^(v52))
                                                                            var v59 uint32 = 146
                                                                            _ = v59
                                                                        } else {
                                                                            if v4 {
                                                                                v6["hello"] = 65535
                                                                            } else {
                                                                                v4 = v4
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v4)
                                                                    if (!(-1 <= (v1 % ((2147483647) | 1)))) {
                                                                        var v60 uint16 = 2
                                                                        var v61 int32 = (v1 >> ((65535) & 63))
                                                                        fmt.Println(fi12, v4, v60)
                                                                        v62 := GenS7{F0: v60, F1: v44}
                                                                        switch v46 {
                                                                        case "foo bar":
                                                                            if v4 {
                                                                                v4 = v4
                                                                            }
                                                                        default:
                                                                            switch (fi12 & fi10) {
                                                                            case 9223372036854775807:
                                                                                var v63 uint16 = (v60 + v60)
                                                                                _ = v63
                                                                            case -2147483648:
                                                                                if (v4 || (0 >= v60)) {
                                                                                    var v64 uint64 = 1
                                                                                    var v65 int = (fi10 >> ((0) & 63))
                                                                                    if (float64(v50) < (float64(v50) * 3.14159)) {
                                                                                        v65 = fi10
                                                                                        v66 := GenS8{F0: v1}
                                                                                        _ = v66
                                                                                    }
                                                                                    fmt.Println(v46, v46, v61)
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                }
                                                                            default:
                                                                                var v67 uint32 = 2
                                                                                _ = v67
                                                                            }
                                                                        }
                                                                        var v68 int8 = (v5 & (v5 << ((0) & 63)))
                                                                        v69 := GenS9{F0: 2}
                                                                        _ = v60
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v68
                                                                        _ = v69
                                                                    } else {
                                                                        var v70 float64 = -2.25
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        if v4 {
                                                                            v45 = v45
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                if v4 {
                                                                                    if (2 <= 152) {
                                                                                        if (((v50 - v50) - (v50 - -2.25)) < (v50 - (v50 - v50))) {
                                                                                            v3 = append(v3, 1)
                                                                                        } else {
                                                                                            if ((1 < -9223372036854775808) || (fi12 >= fi12)) {
                                                                                                v50 = (v50 * (-2.25 / 1.5))
                                                                                                v71 := GenS10{F0: 32767, F1: (v5 + (v5 ^ -1)), F2: 42}
                                                                                                var v72 uint8 = v44
                                                                                                _ = v71
                                                                                                _ = v72
                                                                                            }
                                                                                            fmt.Println(fi12, fi12)
                                                                                        }
                                                                                        var v73 int8 = (v5 - (v5 << ((513) & 63)))
                                                                                        _ = v73
                                                                                    } else {
                                                                                        switch fi12 {
                                                                                        case -854:
                                                                                            if (2 < 1) {
                                                                                                v52 = v52
                                                                                                v1 = 42
                                                                                            }
                                                                                        case 127:
                                                                                            v74 := v6[v46]
                                                                                            _ = v74
                                                                                        case -1:
                                                                                            switch fi10 {
                                                                                            case 127:
                                                                                                if (fi12 > int(v44)) {
                                                                                                    switch "foo bar" {
                                                                                                    case "a":
                                                                                                        v1 += (v52 & 42)
                                                                                                    }
                                                                                                    v75 := GenS3{F0: 1}
                                                                                                    if true {
                                                                                                        var v76 float64 = -2.25
                                                                                                        v76 = 1e10
                                                                                                        fmt.Println(v5, v1)
                                                                                                        _ = v76
                                                                                                    }
                                                                                                    _ = v75
                                                                                                }
                                                                                                v50 += v50
                                                                                                var v77 int8 = v5
                                                                                                _ = v77
                                                                                            }
                                                                                        default:
                                                                                            var v78 uint64 = 1
                                                                                            _ = v78
                                                                                        }
                                                                                    }
                                                                                    _ = len(v3)
                                                                                    v79 := GenS11{F0: 65535, F1: v4}
                                                                                    var v80 uint16 = 255
                                                                                    if (127 < (fi12 * fi12)) {
                                                                                        v6[v46] = 657
                                                                                        delete(v6, (("a" + (v46 + v45)) + v46))
                                                                                        v45 = v46
                                                                                    }
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                }
                                                                            case "foo bar":
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            case "":
                                                                                if false {
                                                                                    v2 = append(v2, 123)
                                                                                    delete(v6, v45)
                                                                                    v45 = v45
                                                                                    fmt.Println(v5, v44)
                                                                                } else {
                                                                                    switch (("foo bar" + (((v46 + v46) + "foo bar") + (v45 + v46))) + v45) {
                                                                                    case "foo bar":
                                                                                        v81 := v6[(v45 + v45)]
                                                                                        _ = v81
                                                                                    case "":
                                                                                        var v82 bool = (2 == 18446744073709551615)
                                                                                        _ = v82
                                                                                    default:
                                                                                        var v83 int32 = (v52 * (v52 % (((v52 + (^(v52)))) | 1)))
                                                                                        _ = v83
                                                                                    }
                                                                                    v84 := GenS7{F0: 0, F1: 0}
                                                                                    v70 = (1.5 * (0.1 / v70))
                                                                                    _ = v84
                                                                                }
                                                                            default:
                                                                                v45 = ("a" + v46)
                                                                            }
                                                                        }
                                                                        _ = v70
                                                                    }
                                                                case -774:
                                                                    switch "" {
                                                                    case "a":
                                                                        v85 := GenS4{F0: (!v4)}
                                                                        _ = v85
                                                                    default:
                                                                        if true {
                                                                            var v86 float32 = v50
                                                                            var v87 uint16 = 38
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                    }
                                                                case 947:
                                                                    v88 := v6[""]
                                                                    _ = v88
                                                                default:
                                                                    if false {
                                                                        var v89 uint = 202
                                                                        fmt.Println(v45, v45)
                                                                        _ = v89
                                                                    }
                                                                }
                                                                switch fi10 {
                                                                case -564:
                                                                    var v90 int = fi10
                                                                    _ = v90
                                                                case 9223372036854775807:
                                                                    if v4 {
                                                                        if v4 {
                                                                            v91 := GenS12{F0: 349, F1: (!(!(!false)))}
                                                                            if (2 == 65535) {
                                                                                switch fi12 {
                                                                                case 4294967295:
                                                                                    switch ("\n\t" + v45) {
                                                                                    case "hello":
                                                                                        v4 = v4
                                                                                    case "foo bar":
                                                                                        if ((!false) && (42 != 1)) {
                                                                                            if (true || ((775 >= 1) || (2 > 4294967295))) {
                                                                                                v46 = ((v46 + v46) + "foo bar")
                                                                                                v52 -= (v52 - (^(v1)))
                                                                                            } else {
                                                                                                fmt.Println(v1, v1)
                                                                                                v92 := GenS13{F0: 66}
                                                                                                if v4 {
                                                                                                    if v4 {
                                                                                                        fmt.Println(v45, fi10)
                                                                                                    } else {
                                                                                                        var v93 int16 = 0
                                                                                                        var v94 uint64 = 1
                                                                                                        _ = v93
                                                                                                        _ = v94
                                                                                                    }
                                                                                                    v1 = v1
                                                                                                } else {
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    var v95 int32 = 1
                                                                                                    _ = v95
                                                                                                }
                                                                                                _ = v92
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        switch ((((("" + (v45 + (v46 + ""))) + "") + ("\n\t" + "a")) + "\n\t") + (v45 + "hello")) {
                                                                                        case "foo bar":
                                                                                            v52 += (v1 >> ((255) & 63))
                                                                                        case "hello":
                                                                                            v5 -= (-(v5))
                                                                                        case "":
                                                                                            fmt.Println(v50, v1)
                                                                                            if v4 {
                                                                                                if (false && true) {
                                                                                                    v4 = (44 <= 65535)
                                                                                                    var v96 uint = 0
                                                                                                    _ = len(v6)
                                                                                                    _ = v96
                                                                                                } else {
                                                                                                    v97 := GenS13{F0: 290}
                                                                                                    fmt.Println(v1, v1, v45)
                                                                                                    v5 += (v5 | (v5 / (((v5 + 0)) | 1)))
                                                                                                    v5 = v5
                                                                                                    var v98 string = v45
                                                                                                    _ = v97
                                                                                                    _ = v98
                                                                                                }
                                                                                                delete(v6, v46)
                                                                                                v45 = "a"
                                                                                            }
                                                                                        default:
                                                                                            switch fi10 {
                                                                                            case 128:
                                                                                                var v99 uint = 4294967295
                                                                                                _ = v99
                                                                                            default:
                                                                                                v100 := GenS11{F0: 2, F1: v4}
                                                                                                var v101 bool = false
                                                                                                var v102 int64 = 9223372036854775807
                                                                                                v6["a"] = 65535
                                                                                                var v103 int32 = 42
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                                _ = v102
                                                                                                _ = v103
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v1 = (v52 & (v52 * 261))
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v104 int64 = -9223372036854775808
                                                                                var v105 int32 = (-(v1))
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                            v1 = (^(v52))
                                                                            _ = v91
                                                                        }
                                                                        var v106 bool = (v4 || v4)
                                                                        var v107 bool = v4
                                                                        v5 -= 1
                                                                        fmt.Println(v106, v1)
                                                                        _ = v106
                                                                        _ = v107
                                                                    } else {
                                                                        switch v45 {
                                                                        case "\n\t":
                                                                            var v108 int64 = -9223372036854775808
                                                                            _ = v108
                                                                        case "foo bar":
                                                                            var v109 bool = (!(4294967295 <= 42))
                                                                            if ((fi12 / ((0) | 1)) < -710) {
                                                                                var v110 int64 = -186
                                                                                fmt.Println(v44, v109)
                                                                                _ = v110
                                                                            } else {
                                                                                v44 = (v44 & (v44 << ((18446744073709551615) & 63)))
                                                                                var v111 float32 = (0.1 + ((-2.25 * -2.25) / v50))
                                                                                _ = v111
                                                                            }
                                                                            var v112 uint = 42
                                                                            _ = v109
                                                                            _ = v112
                                                                        default:
                                                                            v113 := GenS14{F0: -891, F1: 297}
                                                                            _ = v113
                                                                        }
                                                                        if (v5 == v5) {
                                                                            v114 := GenS15{F0: -32768, F1: v1, F2: 42}
                                                                            _ = v114
                                                                        } else {
                                                                            var v115 int16 = 285
                                                                            switch (fi12 / ((fi10) | 1)) {
                                                                            case 9223372036854775807:
                                                                                var v116 uint8 = (v44 % ((v44) | 1))
                                                                                _ = v116
                                                                            case 0:
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    var v117 uint64 = 42
                                                                                    _ = v117
                                                                                }
                                                                            }
                                                                            v115 += 32767
                                                                            delete(v6, ((("hello" + ((v46 + (((v45 + "hello") + "foo bar") + (v45 + v45))) + v45)) + (v46 + v45)) + "\n\t"))
                                                                            var v118 int16 = (v115 | v115)
                                                                            _ = v115
                                                                            _ = v118
                                                                        }
                                                                    }
                                                                }
                                                                _ = v52
                                                            }
                                                            v2 = append(v2, 366)
                                                            v50 = v50
                                                        }
                                                        v119 := GenS10{F0: fi12, F1: 1, F2: -32768}
                                                        switch "a" {
                                                        case "\n\t":
                                                            var v120 uint32 = 974
                                                            _ = v120
                                                        case "a":
                                                            if (v4 || (!v4)) {
                                                                v45 = v46
                                                            }
                                                        case "foo bar":
                                                            var v121 uint = 18446744073709551615
                                                            _ = v121
                                                        }
                                                        _ = v119
                                                        _ = v50
                                                    }
                                                default:
                                                    v5 = (v5 / (((v5 - int8(v44))) | 1))
                                                }
                                            }
                                        case 127:
                                            v122 := GenS9{F0: 0}
                                            _ = v122
                                        }
                                        var v123 int8 = (v5 << ((65535) & 63))
                                        if (v4 || false) {
                                            var v124 uint32 = 2147483647
                                            var v125 float64 = (3.14159 * 1e10)
                                            _ = v124
                                            _ = v125
                                        } else {
                                            _ = len(v6)
                                        }
                                        _ = v123
                                        _ = v46
                                    }
                                case "hello":
                                    if v4 {
                                        var v126 string = (v45 + "")
                                        v127 := GenS14{F0: -1, F1: 42}
                                        _ = v126
                                        _ = v127
                                    } else {
                                        var v128 float64 = (((((((3.14159 / (0.1 * (((1e10 * 3.14159) / 0.1) * ((0.0 + (0.1 - (((3.14159 - 1.5) - ((0.0 - 1e10) * ((3.14159 * -2.25) * 1e10))) + -2.25))) / 0.1)))) + 0.0) + (-2.25 / 3.14159)) / (1.5 / 0.1)) + 0.0) * (3.14159 * 3.14159)) - ((1e10 - ((float64(((((0.1 + float32(3.14159)) * (1e10 / 1.0)) / 3.14159) / 1.5)) / -2.25) * float64(float32(1e10)))) - 3.14159))
                                        var v129 uint32 = 317
                                        delete(v6, "\n\t")
                                        if (v4 || (1 > 641)) {
                                            var v130 float32 = 0.0
                                            var v131 int32 = v1
                                            v44 += 255
                                            if ((true || ((v130 / (1e10 * float32((((((v128 - 3.14159) - ((1.5 + (0.1 - -2.25)) - 1e10)) / v128) - v128) - (((3.14159 * float64(v130)) + v128) * (float64(1e10) / 0.1)))))) > (float32(v128) - float32(((((1e10 / (v128 - (0.1 + v128))) * 1e10) / float64(v130)) / 1.0))))) || ((!(784 >= 2)) || v4)) {
                                                fmt.Println(v129, v45)
                                            }
                                            v130 = ((0.0 * v130) * v130)
                                            v1 = 1
                                            if (!(1 == 4294967295)) {
                                                var v132 int = (fi12 % ((fi10) | 1))
                                                _ = v132
                                            }
                                            fmt.Println(fi12)
                                            _ = v130
                                            _ = v131
                                        }
                                        _ = v128
                                        _ = v129
                                    }
                                case "":
                                    switch (fi10 % ((fi10) | 1)) {
                                    case 0:
                                        switch (fi10 << ((2) & 63)) {
                                        case 42:
                                            v4 = true
                                        case 127:
                                            if (v5 >= v5) {
                                                v3 = append(v3, 331)
                                                v2 = append(v2, 4294967295)
                                                switch fi10 {
                                                case -909:
                                                    fmt.Println(v4, fi10)
                                                case -2147483648:
                                                    var v133 int64 = 9223372036854775807
                                                    _ = v133
                                                case 2147483647:
                                                    var v134 uint32 = 65535
                                                    if ((-(v5)) <= 42) {
                                                        switch fi10 {
                                                        case -559:
                                                            fmt.Println(fi12, v1)
                                                        }
                                                    }
                                                    _ = v134
                                                }
                                                v44 = 0
                                                switch 1 {
                                                case -57:
                                                    var v135 float32 = ((3.14159 - 1.5) / ((0.1 - (3.14159 + 1e10)) * 1e10))
                                                    _ = v135
                                                case -501:
                                                    v136 := GenS3{F0: (v5 + (v5 / (((v5 & (^(v5)))) | 1)))}
                                                    _ = v136
                                                case 942:
                                                    delete(v6, (v45 + v45))
                                                default:
                                                    _ = len(v3)
                                                }
                                            }
                                        case 0:
                                            if ((176 > 255) && false) {
                                                switch (fi10 * fi10) {
                                                case 32767:
                                                    var v137 uint = 65535
                                                    _ = v137
                                                }
                                            } else {
                                                var v138 float32 = 0.0
                                                var v139 uint32 = 65535
                                                if v4 {
                                                    switch ("hello" + ("foo bar" + v45)) {
                                                    case "":
                                                        v44 = (v44 + v44)
                                                    case "hello":
                                                        fmt.Println(v138, fi10, v44)
                                                    case "foo bar":
                                                        v140 := GenS10{F0: 32767, F1: 42, F2: 42}
                                                        _ = v140
                                                    default:
                                                        v141 := GenS14{F0: -2147483648, F1: 2}
                                                        var v142 string = "\n\t"
                                                        var v143 uint64 = 1
                                                        _ = v141
                                                        _ = v142
                                                        _ = v143
                                                    }
                                                    v4 = (!v4)
                                                } else {
                                                    var v144 float32 = float32(-2.25)
                                                    var v145 uint8 = (v44 / ((0) | 1))
                                                    _ = v144
                                                    _ = v145
                                                }
                                                _ = v138
                                                _ = v139
                                            }
                                        default:
                                            v1 -= 166
                                            switch v45 {
                                            case "\n\t":
                                                fmt.Println(v45)
                                            case "hello":
                                                v44 += (v44 << ((1) & 63))
                                                v146 := GenS11{F0: 255, F1: false}
                                                v147 := GenS4{F0: true}
                                                v148 := GenS15{F0: 833, F1: v1, F2: 636}
                                                _ = v146
                                                _ = v147
                                                _ = v148
                                            default:
                                                if v4 {
                                                    var v149 float32 = 1.5
                                                    _ = v149
                                                }
                                            }
                                            if ((241 != v44) || (v45 != "foo bar")) {
                                                var v150 uint = uint(v44)
                                                if true {
                                                    var v151 string = "hello"
                                                    v1 += -71
                                                    _ = v151
                                                } else {
                                                    var v152 float32 = 1e10
                                                    v1 = v1
                                                    _ = v152
                                                }
                                                v4 = (!v4)
                                                _ = v150
                                            } else {
                                                var v153 uint16 = 0
                                                _ = v153
                                            }
                                        }
                                    case 2147483647:
                                        v154 := GenS2{F0: (fi12 + fi12)}
                                        _ = v154
                                    }
                                default:
                                    switch fi10 {
                                    case 125:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 629:
                                        v1 -= -738
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v155 := GenS16{F0: (v5 << ((65535) & 63)), F1: 679, F2: 0}
                                        _ = v155
                                    case -9223372036854775808:
                                        var v156 uint32 = 86
                                        _ = v156
                                    }
                                }
                                _ = v44
                                _ = v45
                            } else {
                                v4 = v4
                            }
                            fmt.Println(fi12)
                            delete(v6, ("\n\t" + "hello"))
                        } else {
                            if v4 {
                                var v157 float64 = float64(0.1)
                                switch int(v5) {
                                case 128:
                                    var v158 uint = 699
                                    _ = v158
                                case -370:
                                    var v159 uint = 18446744073709551615
                                    _ = v159
                                default:
                                    v4 = v4
                                }
                                _ = v157
                            } else {
                                var v160 int16 = 0
                                _ = v160
                            }
                        }
                    default:
                        if (false || (-32768 <= 32767)) {
                            if v4 {
                                v2 = append(v2, 65535)
                            }
                            switch "a" {
                            case "foo bar":
                                if ((v5 * -72) == int8(fi10)) {
                                    if ((("foo bar" + "") + ("a" + ("hello" + ("\n\t" + ("" + ""))))) > (("" + "a") + "foo bar")) {
                                        var v161 uint32 = 112
                                        var v162 int16 = 32767
                                        _ = v161
                                        _ = v162
                                    }
                                    var v163 int16 = -1
                                    _ = v163
                                } else {
                                    v4 = (!(1 < 1))
                                    if (9223372036854775807 >= -9223372036854775808) {
                                        switch fi10 {
                                        case -521:
                                            fmt.Println(v1)
                                        case -901:
                                            switch ("\n\t" + "foo bar") {
                                            case "a":
                                                if (222 < 2147483648) {
                                                    v1 = 1
                                                    if (("hello" + "\n\t") <= "foo bar") {
                                                        var v164 uint64 = 4294967295
                                                        v165 := GenS3{F0: -1}
                                                        var v166 int8 = int8(fi12)
                                                        _ = v164
                                                        _ = v165
                                                        _ = v166
                                                    }
                                                    switch (fi12 % (((fi12 + 9223372036854775807)) | 1)) {
                                                    case 913:
                                                        var v167 float32 = 0.0
                                                        _ = v167
                                                    case 272:
                                                        switch "a" {
                                                        case "":
                                                            switch fi10 {
                                                            case 32768:
                                                                v1 = v1
                                                            case 42:
                                                                switch "\n\t" {
                                                                case "":
                                                                    v1 = -2147483648
                                                                case "foo bar":
                                                                    v168 := GenS17{F0: (!(v4 || v4))}
                                                                    _ = v168
                                                                default:
                                                                    var v169 uint16 = 42
                                                                    _ = v169
                                                                }
                                                            default:
                                                                if (0.0 <= float32(1.5)) {
                                                                    var v170 uint8 = 1
                                                                    _ = v170
                                                                }
                                                            }
                                                        case "\n\t":
                                                            switch (^(fi10)) {
                                                            case 42:
                                                                if false {
                                                                    if (!v4) {
                                                                        if true {
                                                                            v171 := v6["hello"]
                                                                            v5 -= v5
                                                                            _ = v171
                                                                        } else {
                                                                            if ((2 == 1) && (true && v4)) {
                                                                                switch (fi12 >> ((255) & 63)) {
                                                                                case 127:
                                                                                    v1 += v1
                                                                                default:
                                                                                    var v172 uint32 = 42
                                                                                    _ = v172
                                                                                }
                                                                            } else {
                                                                                var v173 uint = 1
                                                                                var v174 uint64 = 4294967295
                                                                                _ = v173
                                                                                _ = v174
                                                                            }
                                                                            var v175 int64 = -9223372036854775808
                                                                            _ = v175
                                                                        }
                                                                        fmt.Println(fi12)
                                                                        var v176 int32 = -2147483648
                                                                        _ = v176
                                                                    }
                                                                } else {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v177 int16 = 32767
                                                                    _ = v177
                                                                }
                                                            default:
                                                                v2 = append(v2, 4294967295)
                                                            }
                                                        case "a":
                                                            v1 -= (v1 / ((42) | 1))
                                                        }
                                                    }
                                                } else {
                                                    var v178 uint64 = 1
                                                    _ = v178
                                                }
                                            }
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = 564
                                            }
                                        }
                                        v1 += (v1 ^ v1)
                                        v1 = (v1 * (v1 & (v1 | (v1 + (-(v1))))))
                                    } else {
                                        var v179 int32 = 769
                                        var v180 uint8 = 127
                                        v181 := GenS7{F0: 42, F1: 176}
                                        _ = v179
                                        _ = v180
                                        _ = v181
                                    }
                                    v5 = 30
                                }
                            case "hello":
                                var v182 int8 = (v5 - (v5 & (v5 + v5)))
                                _ = v182
                            case "a":
                                v1 -= v1
                            }
                        } else {
                            switch fi12 {
                            case -9223372036854775808:
                                switch ("a" + "\n\t") {
                                case "a":
                                    switch ("\n\t" + "\n\t") {
                                    case "foo bar":
                                        var v183 string = ("a" + "")
                                        _ = v183
                                    case "":
                                        fmt.Println(fi12)
                                    }
                                default:
                                    var v184 uint32 = 4294967295
                                    _ = v184
                                }
                                v5 = (-(v5))
                            case 2147483647:
                                v185 := GenS18{F0: -2147483648, F1: 1}
                                _ = v185
                            case 128:
                                var v186 float64 = 1e10
                                _ = v186
                            }
                            v187 := GenS15{F0: 1, F1: v1, F2: 0}
                            v6[""] = 42
                            _ = v187
                        }
                    }
                }
                _ = fi12
            }
            var v188 bool = false
            _ = v188
        }
        fmt.Println(v1, v5, v1)
        _ = fi10
    }
    v4 = true
    if v4 {
        v189 := GenS19{F0: 1, F1: 903}
        var v190 uint32 = 4294967295
        if (618 == 732) {
            fmt.Println(v190, v4)
        }
        _ = v189
        _ = v190
    } else {
        if len(v2) > 0 {
            v2[0] = 0
        }
    }
    for fi191 := 0; fi191 < 19; fi191++ {
        var v192 uint32 = 259
        switch "" {
        case "":
            if (true && v4) {
                v5 = v5
            } else {
                switch "" {
                case "\n\t":
                    switch ("a" + "foo bar") {
                    case "":
                        for fi193 := 0; fi193 < 16; fi193++ {
                            if true {
                                v192 = 2147483647
                            } else {
                                _ = len(v2)
                                switch "hello" {
                                case "hello":
                                    v194 := GenS20{F0: (-9223372036854775808 < -9223372036854775808), F1: (100 < 57)}
                                    _ = v194
                                case "a":
                                    var v195 int64 = -777
                                    _ = v195
                                }
                            }
                            if true {
                                v196 := GenS18{F0: (v1 | (^(v1))), F1: -1}
                                if (2 >= 18446744073709551615) {
                                    var v197 int32 = (^(v1))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v197
                                }
                                var v198 uint32 = v192
                                var v199 bool = v4
                                v2 = append(v2, (v198 + (v198 & (v192 | v198))))
                                _ = v196
                                _ = v198
                                _ = v199
                            } else {
                                v4 = v4
                                var v200 uint16 = 937
                                v200 -= (v200 + 0)
                                v4 = (false && v4)
                                switch ("foo bar" + (("hello" + "") + "")) {
                                case "hello":
                                    if (42 > -32768) {
                                        switch (fi193 % ((fi191) | 1)) {
                                        case 4294967295:
                                            if (2 > 42) {
                                                v5 += v5
                                            } else {
                                                var v201 uint = 0
                                                switch (-(fi193)) {
                                                case 42:
                                                    v202 := v6["hello"]
                                                    _ = v202
                                                default:
                                                    v200 = v200
                                                }
                                                _ = v201
                                            }
                                        case 721:
                                            switch "a" {
                                            case "foo bar":
                                                delete(v6, "")
                                            case "\n\t":
                                                var v203 uint = 2
                                                _ = v203
                                            case "a":
                                                v204 := GenS21{F0: 65535, F1: (^(v200)), F2: (fi191 - -812)}
                                                _ = v204
                                            default:
                                                fmt.Println(fi191, v200)
                                                if true {
                                                    _ = len(v2)
                                                    v205 := GenS14{F0: (v1 << ((0) & 63)), F1: v200}
                                                    if false {
                                                        v2 = append(v2, (v192 << ((1) & 63)))
                                                    }
                                                    switch 9223372036854775807 {
                                                    case 42:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if ("hello" < "hello") {
                                                                var v206 bool = v4
                                                                v207 := GenS22{F0: 2, F1: true}
                                                                _ = v206
                                                                _ = v207
                                                            } else {
                                                                if v4 {
                                                                    fmt.Println(v1, v200)
                                                                } else {
                                                                    v1 = (v1 / ((v1) | 1))
                                                                    _ = len(v6)
                                                                    fmt.Println(v5, fi191)
                                                                }
                                                                if (!v4) {
                                                                    v208 := GenS8{F0: (-(v1))}
                                                                    if (-96 >= (v5 ^ (v5 | v5))) {
                                                                        fmt.Println(fi193, v5)
                                                                    }
                                                                    _ = v208
                                                                }
                                                                var v209 float64 = 1e10
                                                                _ = v209
                                                            }
                                                        case "":
                                                            v210 := GenS4{F0: (5 != 4294967295)}
                                                            _ = v210
                                                        case "\n\t":
                                                            var v211 int = (fi193 ^ 32768)
                                                            _ = v211
                                                        default:
                                                            delete(v6, ("hello" + ((("" + ("a" + "")) + (("foo bar" + ("\n\t" + "a")) + (("a" + "a") + "a"))) + "foo bar")))
                                                        }
                                                    case -2147483648:
                                                        var v212 int64 = -561
                                                        _ = v212
                                                    case 362:
                                                        v200 = v200
                                                    }
                                                    v1 += v1
                                                    _ = v205
                                                } else {
                                                    var v213 uint8 = 2
                                                    _ = v213
                                                }
                                                v214 := GenS21{F0: (v192 << ((4294967295) & 63)), F1: 986, F2: 434}
                                                _ = v214
                                            }
                                        case -124:
                                            v215 := GenS18{F0: 42, F1: 1}
                                            _ = v215
                                        }
                                        var v216 uint64 = 559
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v216
                                    } else {
                                        _ = len(v3)
                                        v4 = (v200 < 65535)
                                        var v217 float64 = 0.0
                                        _ = v217
                                    }
                                case "":
                                    var v218 uint = 141
                                    _ = v218
                                case "foo bar":
                                    if (!(!true)) {
                                        fmt.Println(fi193, v5, fi191)
                                        fmt.Println(fi191, v4)
                                        v2 = append(v2, v192)
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            if v4 {
                                                var v219 int8 = (v5 | (v5 - (^(v5))))
                                                _ = v219
                                            }
                                        case "\n\t":
                                            if v4 {
                                                var v220 float32 = (-2.25 - float32(0.0))
                                                if ((v5 & v5) != (^(v5))) {
                                                    var v221 int8 = 100
                                                    switch ((("hello" + ("foo bar" + ("" + (("" + "\n\t") + "")))) + "") + "\n\t") {
                                                    case "\n\t":
                                                        if true {
                                                            v192 = (v192 & (^(v192)))
                                                            if (4294967295 != 18446744073709551615) {
                                                                v5 -= v221
                                                            }
                                                        }
                                                    default:
                                                        v1 += 42
                                                    }
                                                    var v222 uint = 695
                                                    _ = v221
                                                    _ = v222
                                                } else {
                                                    v223 := GenS23{F0: 1, F1: 42}
                                                    _ = len(v6)
                                                    _ = v223
                                                }
                                                v5 = 42
                                                if ((0 == 228) || (4294967295 == 255)) {
                                                    fmt.Println(v1, fi191, v1)
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = v192
                                                    }
                                                    var v224 string = ""
                                                    _ = v224
                                                }
                                                v4 = false
                                                _ = v220
                                            }
                                        case "foo bar":
                                            v225 := GenS24{F0: 240}
                                            _ = v225
                                        default:
                                            var v226 float32 = float32(float64((1e10 + 1e10)))
                                            _ = v226
                                        }
                                    }
                                default:
                                    var v227 bool = v4
                                    _ = v227
                                }
                                _ = v200
                            }
                            var v228 uint32 = 0
                            if v4 {
                                v228 += (v228 | (v228 >> ((18446744073709551615) & 63)))
                            } else {
                                if ((42 != 42) && ((3.14159 - 1e10) < 1.5)) {
                                    fmt.Println(fi191, v228, fi191)
                                }
                            }
                            v228 += v192
                            _ = fi193
                            _ = v228
                        }
                    case "foo bar":
                        if false {
                            fmt.Println(v1, v192)
                            var v229 int = (fi191 / (((fi191 * fi191)) | 1))
                            _ = v229
                        }
                    case "a":
                        switch (fi191 * 42) {
                        case 32768:
                            fmt.Println(v192)
                        case 127:
                            v230 := GenS10{F0: (fi191 / ((32767) | 1)), F1: (v5 * (v5 * v5)), F2: -348}
                            _ = v230
                        }
                    }
                case "foo bar":
                    for fi231 := 0; fi231 < 16; fi231++ {
                        v2 = append(v2, 1)
                        var v232 int64 = 9223372036854775807
                        if (18446744073709551615 >= 711) {
                            v4 = true
                        } else {
                            fmt.Println(v192, v4)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        if ("\n\t" == "hello") {
                            if false {
                                fmt.Println(fi231, v1)
                            } else {
                                var v233 float32 = 1.5
                                var v234 uint16 = 1
                                v5 = 1
                                _ = v233
                                _ = v234
                            }
                        } else {
                            v235 := GenS20{F0: v4, F1: false}
                            var v236 uint64 = 4294967295
                            var v237 uint64 = (v236 & v236)
                            _ = v235
                            _ = v236
                            _ = v237
                        }
                        _ = len(v6)
                        _ = fi231
                        _ = v232
                    }
                default:
                    if ((!true) || false) {
                        var v238 int16 = 42
                        v238 += v238
                        v239 := v6["a"]
                        _ = v238
                        _ = v239
                    } else {
                        if (!(-294 <= -430)) {
                            for fi240 := 0; fi240 < 6; fi240++ {
                                var v241 int16 = -293
                                var v242 int64 = -9223372036854775808
                                var v243 uint64 = 640
                                v244 := GenS17{F0: (0.0 <= 1.5)}
                                _ = fi240
                                _ = v241
                                _ = v242
                                _ = v243
                                _ = v244
                            }
                            var v245 float32 = 0.0
                            switch fi191 {
                            case 4294967295:
                                fmt.Println(v5, v5)
                            case 2147483647:
                                if false {
                                    v246 := GenS24{F0: 772}
                                    _ = v246
                                }
                            case 42:
                                v1 = (v1 ^ (v1 >> ((2) & 63)))
                            default:
                                if (0 < 255) {
                                    switch 292 {
                                    case 849:
                                        var v247 int8 = (v5 & v5)
                                        _ = v247
                                    default:
                                        fmt.Println(fi191, v192, v5)
                                    }
                                    for fi248 := 0; fi248 < 12; fi248++ {
                                        switch 1 {
                                        case -1:
                                            fmt.Println(v1, v4, v245)
                                        case 32768:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case 141:
                                            var v249 string = ("a" + ("foo bar" + "\n\t"))
                                            _ = v249
                                        }
                                        _ = fi248
                                    }
                                } else {
                                    v250 := GenS25{F0: (v1 * -418), F1: 702}
                                    v4 = v4
                                    _ = v250
                                }
                            }
                            _ = v245
                        } else {
                            var v251 uint32 = v192
                            v252 := GenS6{F0: 432, F1: v251}
                            _ = v251
                            _ = v252
                        }
                        v5 += v5
                    }
                }
                for fi253 := 0; fi253 < 19; fi253++ {
                    var v254 int32 = v1
                    v192 -= v192
                    _ = fi253
                    _ = v254
                }
                for fi255 := 0; fi255 < 15; fi255++ {
                    v2 = append(v2, (v192 / (((v192 ^ 2)) | 1)))
                    _ = fi255
                }
            }
        case "a":
            for fi256 := 0; fi256 < 16; fi256++ {
                switch "hello" {
                case "hello":
                    var v257 uint16 = 42
                    _ = v257
                case "":
                    if false {
                        fmt.Println(fi256, v192)
                        var v258 int16 = 15
                        var v259 int64 = -9223372036854775808
                        _ = v258
                        _ = v259
                    }
                case "a":
                    v260 := GenS2{F0: fi256}
                    _ = v260
                }
                if (42 < 4294967295) {
                    fmt.Println(fi191, fi191)
                    _ = len(v2)
                    v192 = (v192 ^ v192)
                }
                _ = fi256
            }
        case "hello":
            for fi261 := 0; fi261 < 19; fi261++ {
                var v262 int32 = (-(v1))
                _ = fi261
                _ = v262
            }
        default:
            if ((-1 != 0) || v4) {
                if (!((false && false) && false)) {
                    v263 := GenS19{F0: 56, F1: -990}
                    _ = len(v6)
                    _ = v263
                }
                if v4 {
                    var v264 uint32 = (v192 - v192)
                    _ = v264
                } else {
                    switch ("foo bar" + (("" + "hello") + ("" + ""))) {
                    case "\n\t":
                        if ((v192 & v192) != 624) {
                            var v265 uint16 = 957
                            v266 := GenS14{F0: v1, F1: (v265 - v265)}
                            _ = v265
                            _ = v266
                        } else {
                            fmt.Println(v5)
                            for fi267 := 0; fi267 < 18; fi267++ {
                                switch (fi191 >> ((42) & 63)) {
                                case -9223372036854775808:
                                    v192 += (v192 * (v192 - (v192 ^ v192)))
                                case 1:
                                    v1 -= (v1 % (((-(v1))) | 1))
                                case 128:
                                    var v268 uint8 = 253
                                    _ = v268
                                default:
                                    switch fi267 {
                                    case -2147483648:
                                        v269 := GenS1{F0: 974, F1: v192}
                                        _ = v269
                                    case 558:
                                        if (!(true && ((float32(0.1) == ((0.0 / ((0.1 - 0.1) + -2.25)) / 1.5)) && ((!false) || v4)))) {
                                            var v270 int = 128
                                            _ = v270
                                        }
                                    case 4294967295:
                                        delete(v6, "hello")
                                    }
                                }
                                _ = fi267
                            }
                            v192 = (v192 << ((4294967295) & 63))
                        }
                    case "hello":
                        if v4 {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v271 string = "\n\t"
                            fmt.Println(fi191, v271, v4)
                            _ = v271
                        } else {
                            switch (fi191 / ((fi191) | 1)) {
                            case -316:
                                v272 := v6["\n\t"]
                                _ = v272
                            case -1:
                                v5 = v5
                            }
                        }
                    case "foo bar":
                        var v273 bool = false
                        _ = v273
                    default:
                        for fi274 := 0; fi274 < 20; fi274++ {
                            v4 = v4
                            if (v4 && false) {
                                v5 = 127
                            } else {
                                if v4 {
                                    if v4 {
                                        v275 := GenS13{F0: 776}
                                        var v276 uint8 = 128
                                        fmt.Println(v192)
                                        _ = v275
                                        _ = v276
                                    }
                                    fmt.Println(v1, v192, fi191)
                                }
                                v277 := GenS15{F0: -744, F1: v1, F2: (v192 >> ((4294967295) & 63))}
                                v192 = 65535
                                _ = v277
                            }
                            _ = fi274
                        }
                    }
                }
            } else {
                switch fi191 {
                case -56:
                    for fi278 := 0; fi278 < 20; fi278++ {
                        switch "" {
                        case "\n\t":
                            var v279 int32 = v1
                            _ = v279
                        default:
                            switch (fi278 ^ (fi278 * fi191)) {
                            case -385:
                                fmt.Println(fi191)
                            }
                        }
                        if true {
                            var v280 string = "a"
                            _ = len(v6)
                            if (!v4) {
                                v5 = -54
                            } else {
                                fmt.Println(fi278)
                                if len(v2) > 0 {
                                    v2[0] = (^(v192))
                                }
                                var v281 int16 = 604
                                v282 := GenS3{F0: (^(v5))}
                                v2 = append(v2, 2)
                                _ = v281
                                _ = v282
                            }
                            v4 = (42 <= (^(v1)))
                            v5 -= 100
                            _ = v280
                        } else {
                            v283 := GenS17{F0: v4}
                            fmt.Println(fi278, v4)
                            v284 := GenS14{F0: v1, F1: 42}
                            _ = len(v6)
                            if false {
                                v285 := GenS22{F0: 255, F1: v4}
                                _ = v285
                            } else {
                                switch fi278 {
                                case -138:
                                    v6[(("" + "a") + ("a" + ("\n\t" + "hello")))] = v192
                                case 32768:
                                    switch fi278 {
                                    case 32768:
                                        v286 := GenS26{F0: 482, F1: (!true)}
                                        _ = v286
                                    case 9223372036854775807:
                                        v3 = append(v3, v192)
                                    case 1:
                                        v287 := GenS1{F0: 1, F1: v192}
                                        _ = v287
                                    }
                                default:
                                    var v288 uint64 = 2
                                    _ = v288
                                }
                                _ = len(v2)
                                var v289 int = 1
                                if (623 >= 0) {
                                    v4 = (!false)
                                    v1 = -2147483648
                                } else {
                                    if (!v4) {
                                        if false {
                                            v5 = (v5 % ((v5) | 1))
                                            v290 := GenS16{F0: (v5 / (((v5 - 1)) | 1)), F1: 545, F2: (v192 << ((584) & 63))}
                                            fmt.Println(fi278, v1, v289)
                                            _ = v290
                                        } else {
                                            if v4 {
                                                v6["\n\t"] = 65535
                                                fmt.Println(v4)
                                            } else {
                                                v289 = fi278
                                                v5 += (v5 + 9)
                                                if (255 >= 500) {
                                                    v4 = v4
                                                    v192 = (v192 ^ (v192 & v192))
                                                    v192 -= v192
                                                } else {
                                                    if v4 {
                                                        delete(v6, "a")
                                                        var v291 int8 = (v5 << ((750) & 63))
                                                        switch 1 {
                                                        case 409:
                                                            fmt.Println(v291)
                                                        case 302:
                                                            v292 := GenS23{F0: 0, F1: 80}
                                                            _ = v292
                                                        case -578:
                                                            if len(v2) > 0 {
                                                                v2[0] = (^(v192))
                                                            }
                                                        default:
                                                            _ = len(v6)
                                                        }
                                                        _ = v291
                                                    }
                                                    v293 := GenS18{F0: v1, F1: -963}
                                                    v4 = (1 != 905)
                                                    v294 := GenS27{F0: 92, F1: 0}
                                                    _ = v293
                                                    _ = v294
                                                }
                                            }
                                        }
                                        var v295 string = "hello"
                                        v5 = v5
                                        _ = v295
                                    } else {
                                        switch 2147483647 {
                                        case 9223372036854775807:
                                            v296 := GenS28{F0: 0, F1: 2, F2: 1}
                                            _ = v296
                                        case -815:
                                            var v297 string = ""
                                            _ = v297
                                        default:
                                            v192 = v192
                                        }
                                        v5 += -1
                                        _ = len(v3)
                                        var v298 uint64 = 4294967295
                                        v289 = v289
                                        _ = v298
                                    }
                                    v299 := GenS29{F0: 18446744073709551615}
                                    fmt.Println(fi191, v192)
                                    var v300 uint32 = v192
                                    if v4 {
                                        v3 = append(v3, (v192 + v300))
                                        if len(v2) > 0 {
                                            v2[0] = (v192 / (((v300 * 2147483648)) | 1))
                                        }
                                    } else {
                                        if (266 <= 2) {
                                            if false {
                                                switch (fi191 ^ (fi278 - (^(fi191)))) {
                                                case 32767:
                                                    var v301 uint = 42
                                                    var v302 int8 = v5
                                                    fmt.Println(v1, v302, v302)
                                                    v300 = v192
                                                    _ = v301
                                                    _ = v302
                                                case -1:
                                                    var v303 int32 = -575
                                                    v1 -= v303
                                                    _ = v303
                                                default:
                                                    fmt.Println(v192, fi278)
                                                }
                                            } else {
                                                if v4 {
                                                    fmt.Println(fi191, v192, fi191)
                                                    var v304 bool = (false && (138 < 42))
                                                    _ = v304
                                                }
                                                if (847 <= 42) {
                                                    var v305 int64 = 42
                                                    if false {
                                                        v306 := GenS4{F0: (!((v192 * 0) > (v192 + v192)))}
                                                        _ = v306
                                                    }
                                                    _ = v305
                                                }
                                                v307 := GenS30{F0: 0, F1: 100, F2: v4}
                                                _ = v307
                                            }
                                            var v308 uint = 230
                                            _ = v308
                                        }
                                    }
                                    _ = v299
                                    _ = v300
                                }
                                var v309 uint64 = 42
                                _ = v289
                                _ = v309
                            }
                            _ = v283
                            _ = v284
                        }
                        switch -2147483648 {
                        case 32767:
                            switch "" {
                            case "hello":
                                if v4 {
                                    v5 = v5
                                    v4 = ((^(v5)) < 93)
                                    var v310 uint = 124
                                    _ = v310
                                } else {
                                    var v311 uint8 = 2
                                    v5 += 42
                                    delete(v6, ("foo bar" + "a"))
                                    v1 += v1
                                    _ = v311
                                }
                            case "a":
                                switch (("a" + "foo bar") + (("\n\t" + ((("foo bar" + ((((("hello" + (((("hello" + "foo bar") + "foo bar") + "") + "foo bar")) + "") + ("hello" + "hello")) + "") + ("foo bar" + ("hello" + ("" + "a"))))) + (("\n\t" + ("hello" + "")) + "foo bar")) + (("hello" + "foo bar") + "hello"))) + ("\n\t" + "\n\t"))) {
                                case "":
                                    if ((v192 * v192) < (v192 * (v192 << ((685) & 63)))) {
                                        var v312 int16 = 32767
                                        var v313 uint32 = (^(v192))
                                        var v314 int8 = (v5 / (((-(v5))) | 1))
                                        _ = v312
                                        _ = v313
                                        _ = v314
                                    } else {
                                        v4 = (v192 <= (v192 / (((v192 ^ v192)) | 1)))
                                        v315 := GenS31{F0: ""}
                                        _ = v315
                                    }
                                case "\n\t":
                                    v316 := GenS13{F0: 2}
                                    _ = v316
                                case "hello":
                                    v5 -= v5
                                }
                            default:
                                if v4 {
                                    v317 := v6["\n\t"]
                                    _ = v317
                                } else {
                                    v318 := GenS12{F0: 128, F1: v4}
                                    v1 += (v1 ^ (v1 << ((255) & 63)))
                                    v1 = int32(v5)
                                    _ = v318
                                }
                            }
                        case 0:
                            var v319 int32 = (^(v1))
                            _ = v319
                        case 9223372036854775807:
                            v192 = (v192 % (((^(v192))) | 1))
                        }
                        _ = fi278
                    }
                case -2147483648:
                    if v4 {
                        var v320 int64 = -460
                        switch "hello" {
                        case "a":
                            fmt.Println(v192, v5)
                        case "\n\t":
                            var v321 int64 = -48
                            _ = v321
                        case "foo bar":
                            if (0.1 >= (0.0 - (float64((-2.25 * -2.25)) + 0.0))) {
                                for fi322 := 0; fi322 < 15; fi322++ {
                                    if (255 > 1) {
                                        v3 = append(v3, 594)
                                    } else {
                                        if ((!((fi191 - 32768) > (^(fi191)))) || ((v320 << ((18446744073709551615) & 63)) >= v320)) {
                                            fmt.Println(v5, v320)
                                        } else {
                                            v5 = v5
                                            fmt.Println(fi322)
                                        }
                                    }
                                    var v323 int16 = 42
                                    v5 = 103
                                    _ = fi322
                                    _ = v323
                                }
                                var v324 float64 = 0.0
                                _ = v324
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v320
                    }
                case -669:
                    if v4 {
                        var v325 string = "hello"
                        for fi326 := 0; fi326 < 3; fi326++ {
                            fmt.Println(fi191, v1, fi191)
                            if true {
                                if ((fi191 << ((525) & 63)) <= (fi326 - 32768)) {
                                    fmt.Println(fi191, fi326)
                                    switch v325 {
                                    case "foo bar":
                                        var v327 int = (fi191 % ((-1) | 1))
                                        _ = v327
                                    case "a":
                                        if false {
                                            var v328 uint16 = 1
                                            var v329 float32 = 1.5
                                            v192 += uint32(v328)
                                            _ = v328
                                            _ = v329
                                        }
                                    case "\n\t":
                                        var v330 string = ""
                                        _ = v330
                                    }
                                    fmt.Println(fi326, v192, v4)
                                    var v331 uint16 = 2
                                    _ = v331
                                }
                                v3 = append(v3, (v192 / (((^(v192))) | 1)))
                                fmt.Println(fi191, v1)
                            }
                            _ = fi326
                        }
                        switch (fi191 - fi191) {
                        case -2147483648:
                            var v332 int16 = 852
                            _ = v332
                        case 9223372036854775807:
                            v6[v325] = 2147483647
                        default:
                            for fi333 := 0; fi333 < 6; fi333++ {
                                var v334 float64 = 0.1
                                fmt.Println(v192, v334)
                                _ = fi333
                                _ = v334
                            }
                        }
                        _ = v325
                    }
                default:
                    fmt.Println(v5, fi191)
                }
                if (127 >= 128) {
                    var v335 int16 = 1
                    _ = v335
                } else {
                    var v336 uint8 = 0
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v337 int16 = 1000
                    _ = v336
                    _ = v337
                }
                v6[((("\n\t" + "") + "foo bar") + ("foo bar" + ("hello" + "foo bar")))] = 2
            }
        }
        _ = fi191
        _ = v192
    }
    switch "hello" {
    case "":
        var v338 bool = v4
        _ = v338
    case "\n\t":
        switch 127 {
        case 0:
            fmt.Println(v5)
        case -102:
            v1 += v1
        }
    }
    if v4 {
        if (("foo bar" < "a") || (42 != 65535)) {
            var v339 float32 = 1e10
            _ = v339
        } else {
            v340 := GenS32{F0: (("hello" + "\n\t") + ("hello" + "\n\t"))}
            _ = v340
        }
        var v341 uint64 = 1
        _ = v341
    } else {
        for fi342 := 0; fi342 < 16; fi342++ {
            v4 = (fi342 != (fi342 & fi342))
            v343 := GenS21{F0: 833, F1: 42, F2: fi342}
            v344 := GenS33{F0: -1}
            _ = fi342
            _ = v343
            _ = v344
        }
        v4 = v4
    }
    v4 = true
    if v4 {
        if v4 {
            var v345 int32 = -2147483648
            for fi346 := 0; fi346 < 16; fi346++ {
                _ = len(v6)
                _ = fi346
            }
            _ = len(v6)
            _ = v345
        } else {
            for fi347 := 0; fi347 < 12; fi347++ {
                for fi348 := 0; fi348 < 4; fi348++ {
                    v1 -= v1
                    if (255 == 128) {
                        v1 -= v1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v349 := GenS32{F0: ((("\n\t" + ("a" + "")) + "a") + ("a" + ("foo bar" + "\n\t")))}
                        var v350 uint16 = 1
                        _ = v349
                        _ = v350
                    }
                    _ = fi348
                }
                v351 := GenS34{F0: -9223372036854775808}
                v5 -= (v5 % ((v5) | 1))
                _ = fi347
                _ = v351
            }
        }
    } else {
        v4 = v4
        for fi352 := 0; fi352 < 1; fi352++ {
            for fi353 := 0; fi353 < 20; fi353++ {
                var v354 int8 = (v5 + -50)
                v5 -= (v5 ^ v5)
                _ = fi353
                _ = v354
            }
            _ = fi352
        }
        v4 = (2 != 145)
    }
    for fi355 := 0; fi355 < 11; fi355++ {
        if ((42 <= 0) && v4) {
            fmt.Println(v1)
            v5 = 113
            if ((v5 >> ((2) & 63)) > 42) {
                for fi356 := 0; fi356 < 4; fi356++ {
                    delete(v6, "a")
                    var v357 uint8 = 255
                    v1 += (v1 - v1)
                    _ = fi356
                    _ = v357
                }
                for fi358 := 0; fi358 < 1; fi358++ {
                    v6[("a" + "a")] = 200
                    if ("foo bar" >= "") {
                        fmt.Println(v4, fi355)
                    }
                    v5 -= -128
                    _ = fi358
                }
            }
        }
        fmt.Println(fi355, fi355)
        switch (fi355 << ((228) & 63)) {
        case 32767:
            var v359 uint = 4294967295
            _ = v359
        case -658:
            v360 := GenS1{F0: 0, F1: 1}
            _ = v360
        case -399:
            if v4 {
                v361 := GenS25{F0: (v1 - (v1 - v1)), F1: 18446744073709551615}
                for fi362 := 0; fi362 < 18; fi362++ {
                    var v363 int64 = 184
                    v1 = v1
                    if (v363 < int64(fi355)) {
                        if (!(v363 != (v363 ^ v363))) {
                            var v364 int64 = int64(v1)
                            _ = v364
                        }
                    } else {
                        v365 := GenS21{F0: 2, F1: 728, F2: 4294967295}
                        if true {
                            if (134 <= 127) {
                                var v366 int16 = 0
                                v366 = (v366 ^ 758)
                                v366 = 0
                                _ = v366
                            }
                        } else {
                            v367 := GenS19{F0: 2, F1: v363}
                            if true {
                                v368 := v6["\n\t"]
                                if true {
                                    v368 += 333
                                    v369 := v6[""]
                                    v5 -= 42
                                    var v370 int32 = v1
                                    v369 = (v368 & v368)
                                    switch fi355 {
                                    case 32767:
                                        switch (^(fi362)) {
                                        case 128:
                                            fmt.Println(v1)
                                        }
                                    default:
                                        v5 = 1
                                    }
                                    _ = v369
                                    _ = v370
                                } else {
                                    v5 += v5
                                }
                                v363 = (v363 & v363)
                                var v371 uint32 = v368
                                _ = v368
                                _ = v371
                            } else {
                                var v372 bool = (v4 && false)
                                var v373 string = ""
                                _ = v372
                                _ = v373
                            }
                            fmt.Println(fi355, v363, v5)
                            _ = v367
                        }
                        _ = v365
                    }
                    _ = fi362
                    _ = v363
                }
                v374 := GenS35{F0: 128, F1: 375, F2: (v1 / (((^(v1))) | 1))}
                _ = v361
                _ = v374
            }
        }
        _ = fi355
    }
    v1 += v1
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
