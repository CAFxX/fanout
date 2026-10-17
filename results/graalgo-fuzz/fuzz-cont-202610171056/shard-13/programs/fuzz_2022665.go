package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int16; F1 uint16; F2 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 uint64; F2 uint64 }
type GenS8 struct { F0 bool; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint; F1 int8; F2 int32 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint; F1 int32 }
type GenS13 struct { F0 int32; F1 int }
type GenS14 struct { F0 bool; F1 bool }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint8; F1 uint8; F2 uint32 }

func main() {
    v1 := []uint32{603, 154}
    var v2 int = -104
    var v3 int = (v2 ^ (v2 >> ((106) & 63)))
    var v4 string = ((("hello" + "\n\t") + ("" + ((("a" + "foo bar") + "foo bar") + "\n\t"))) + "\n\t")
    var v5 string = (v4 + "")
    switch 968 {
    case 0:
        var v6 uint64 = 0
        _ = v6
    case 4294967295:
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int32 = 1
                _ = fi8
                _ = v9
            }
            fmt.Println(v4)
            switch v4 {
            case "hello":
                v4 = ""
            }
            _ = fi7
        }
    }
    v5 = v4
    switch (v4 + ((v5 + v4) + ("a" + v5))) {
    case "a":
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 uint8 = 127
                v13 := GenS1{F0: (v12 * (v12 >> ((42) & 63)))}
                var v14 uint64 = 365
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            fmt.Println(v4, v5)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
        }
    case "foo bar":
        var v15 int8 = 127
        _ = v15
    case "\n\t":
        v3 -= (^(v2))
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(v5, v4)
        for fi17 := 0; fi17 < 9; fi17++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v18 int = 189
            v4 = ("foo bar" + "")
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if len(v1) > 0 {
        v1[0] = 386
    }
    v19 := GenS1{F0: 255}
    if ((true && (921 > -9223372036854775808)) || false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v20 int8 = -128
        for fi21 := 0; fi21 < 11; fi21++ {
            if (true && false) {
                v22 := GenS2{F0: false, F1: v20}
                if (((v3 * -9223372036854775808) > -643) || true) {
                    v3 -= v2
                    v3 += fi21
                } else {
                    var v23 bool = (42 != -218)
                    if (42 < -230) {
                        var v24 uint8 = 52
                        switch v5 {
                        case "\n\t":
                            var v25 float32 = (0.1 - -2.25)
                            _ = v25
                        case "hello":
                            if v23 {
                                switch v5 {
                                case "hello":
                                    if true {
                                        for fi26 := 0; fi26 < 11; fi26++ {
                                            v20 = 42
                                            if v23 {
                                                v27 := GenS2{F0: (1 < 32767), F1: v20}
                                                v3 = 32767
                                                var v28 uint = 51
                                                _ = v27
                                                _ = v28
                                            }
                                            var v29 int8 = (v20 % (((v20 << ((1) & 63))) | 1))
                                            _ = fi26
                                            _ = v29
                                        }
                                        for fi30 := 0; fi30 < 13; fi30++ {
                                            if (991 != 42) {
                                                var v31 string = v5
                                                if len(v1) > 0 {
                                                    v1[0] = 119
                                                }
                                                var v32 uint16 = 32767
                                                _ = v31
                                                _ = v32
                                            }
                                            fmt.Println(v2, v5, fi21)
                                            if (!v23) {
                                                var v33 float32 = (0.0 - float32(0.0))
                                                if v23 {
                                                    fmt.Println(v4, fi21, v20)
                                                    v33 = float32((-2.25 - 0.0))
                                                } else {
                                                    v34 := GenS2{F0: false, F1: (^(v20))}
                                                    if v23 {
                                                        var v35 int = 42
                                                        if len(v1) > 0 {
                                                            v1[0] = 413
                                                        }
                                                        _ = v35
                                                    }
                                                    v36 := GenS3{F0: -583, F1: 65535}
                                                    _ = v34
                                                    _ = v36
                                                }
                                                _ = v33
                                            }
                                            _ = fi30
                                        }
                                    }
                                default:
                                    fmt.Println(fi21, v20)
                                }
                            } else {
                                v5 = "\n\t"
                            }
                        default:
                            switch ((v5 + ((v5 + v4) + v5)) + "foo bar") {
                            case "hello":
                                for fi37 := 0; fi37 < 19; fi37++ {
                                    var v38 string = "hello"
                                    _ = len(v1)
                                    var v39 uint = 512
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                            default:
                                var v40 uint16 = 42
                                _ = v40
                            }
                        }
                        _ = v24
                    }
                    if v23 {
                        for fi41 := 0; fi41 < 14; fi41++ {
                            var v42 uint64 = 18446744073709551615
                            v43 := GenS3{F0: -113, F1: 2}
                            v4 = (v5 + v4)
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        var v44 int8 = v20
                        _ = v44
                    } else {
                        for fi45 := 0; fi45 < 15; fi45++ {
                            switch (-(fi21)) {
                            case 127:
                                v46 := GenS2{F0: v23, F1: v20}
                                _ = v46
                            }
                            _ = fi45
                        }
                        fmt.Println(v23, v2, fi21)
                    }
                    _ = v23
                }
                var v47 int = (fi21 & v2)
                _ = v22
                _ = v47
            } else {
                var v48 uint8 = 60
                _ = v48
            }
            if false {
                if true {
                    var v49 float32 = 0.0
                    _ = v49
                }
                if false {
                    switch -1 {
                    case 32767:
                        fmt.Println(v2)
                    }
                    fmt.Println(v3, v5)
                    v2 -= -1
                } else {
                    v3 = (^(fi21))
                }
            } else {
                for fi50 := 0; fi50 < 5; fi50++ {
                    var v51 string = v5
                    v5 = v51
                    var v52 uint8 = 2
                    _ = len(v1)
                    var v53 uint16 = 770
                    if true {
                        v54 := GenS4{F0: 266, F1: v2}
                        v55 := GenS2{F0: true, F1: v20}
                        v20 -= 42
                        v20 = 0
                        switch 128 {
                        case -671:
                            var v56 uint16 = uint16(v52)
                            _ = v56
                        case 128:
                            if (42 != 183) {
                                switch v2 {
                                case 949:
                                    if false {
                                        switch fi50 {
                                        case 128:
                                            if len(v1) > 0 {
                                                v1[0] = 2147483647
                                            }
                                        case -129:
                                            v1 = append(v1, 2)
                                        }
                                        v20 = v20
                                        if (false || (true || (65535 == 2147483647))) {
                                            var v57 int = fi21
                                            v51 = "\n\t"
                                            _ = v57
                                        } else {
                                            v1 = append(v1, 0)
                                            v58 := GenS1{F0: v52}
                                            switch (v51 + ((("a" + v51) + (("" + ("\n\t" + ("hello" + v51))) + ("hello" + v4))) + "a")) {
                                            case "":
                                                _ = len(v1)
                                            case "\n\t":
                                                fmt.Println(fi21, v2, v4)
                                            }
                                            _ = v58
                                        }
                                    } else {
                                        switch fi21 {
                                        case 32768:
                                            v51 = "hello"
                                        case 1:
                                            v2 -= 2147483647
                                        case 9223372036854775807:
                                            if true {
                                                _ = len(v1)
                                                if false {
                                                    if len(v1) > 0 {
                                                        v1[0] = 430
                                                    }
                                                    fmt.Println(v52, fi21, v20)
                                                } else {
                                                    v51 = v5
                                                    var v59 string = "a"
                                                    _ = v59
                                                }
                                                v60 := GenS4{F0: 1000, F1: v2}
                                                _ = v60
                                            }
                                        }
                                        v61 := GenS5{F0: 401, F1: 42, F2: (-(v3))}
                                        switch 0 {
                                        case 517:
                                            var v62 int8 = (v20 / ((-105) | 1))
                                            _ = v62
                                        case 32767:
                                            if true {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                fmt.Println(fi50)
                                            } else {
                                                var v63 float32 = -2.25
                                                v3 += (-(v3))
                                                var v64 float64 = 3.14159
                                                if (false && ((v53 % (((v53 << ((757) & 63))) | 1)) >= 1)) {
                                                    switch 1 {
                                                    case 127:
                                                        fmt.Println(fi50, v2)
                                                    case 677:
                                                        switch (v2 << ((uint(v53)) & 63)) {
                                                        case 548:
                                                            v20 = (v20 >> ((0) & 63))
                                                        default:
                                                            v53 -= 32767
                                                        }
                                                    }
                                                }
                                                _ = v63
                                                _ = v64
                                            }
                                        case 9223372036854775807:
                                            switch "a" {
                                            case "":
                                                fmt.Println(v51, v52)
                                            case "foo bar":
                                                if (false || (2147483647 <= 217)) {
                                                    v1 = append(v1, 250)
                                                    v65 := GenS1{F0: v52}
                                                    _ = v65
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = 1
                                                    }
                                                    if (float64(1.5) >= 0.0) {
                                                        v66 := GenS1{F0: v52}
                                                        _ = v66
                                                    } else {
                                                        v53 = v53
                                                    }
                                                    if (-1 > -9223372036854775808) {
                                                        _ = len(v1)
                                                    } else {
                                                        v4 = (v4 + v5)
                                                        v53 = 130
                                                        if (1 == 2) {
                                                            var v67 uint64 = 4294967295
                                                            v20 -= 0
                                                            switch v5 {
                                                            case "a":
                                                                if ((v53 << ((uint(v53)) & 63)) == v53) {
                                                                    v68 := GenS1{F0: (v52 ^ v52)}
                                                                    switch (v3 << ((255) & 63)) {
                                                                    case 128:
                                                                        switch v3 {
                                                                        case 231:
                                                                            _ = len(v1)
                                                                        }
                                                                    }
                                                                    _ = v68
                                                                } else {
                                                                    v3 = (v2 ^ fi21)
                                                                    if false {
                                                                        switch v4 {
                                                                        case "":
                                                                            _ = len(v1)
                                                                        case "\n\t":
                                                                            if (1 == int32(v20)) {
                                                                                v69 := GenS3{F0: 9223372036854775807, F1: 32767}
                                                                                v4 = ("\n\t" + v4)
                                                                                _ = v69
                                                                            }
                                                                        default:
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                    }
                                                                    v51 = (v51 + "hello")
                                                                }
                                                            case "foo bar":
                                                                var v70 uint8 = v52
                                                                _ = v70
                                                            case "\n\t":
                                                                var v71 uint8 = (v52 * v52)
                                                                _ = v71
                                                            default:
                                                                switch ((v5 + ((v5 + (((v5 + "") + v5) + "hello")) + "hello")) + "") {
                                                                case "":
                                                                    v2 = (^(fi21))
                                                                case "hello":
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 2
                                                                    }
                                                                case "\n\t":
                                                                    var v72 float32 = (((3.14159 + 3.14159) * float32(3.14159)) - 1.5)
                                                                    fmt.Println(v4, v53, fi50)
                                                                    _ = v72
                                                                default:
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 65535
                                                                        }
                                                                    case "hello":
                                                                        fmt.Println(v20)
                                                                    default:
                                                                        if (1 != uint32(v52)) {
                                                                            if (int16(v67) != 1) {
                                                                                v52 -= (v52 ^ 1)
                                                                                if (-771 != 42) {
                                                                                    var v73 string = (("a" + (v4 + "hello")) + v5)
                                                                                    switch (v51 + (v4 + v51)) {
                                                                                    case "hello":
                                                                                        v74 := GenS1{F0: (v52 >> ((uint(v67)) & 63))}
                                                                                        _ = v74
                                                                                    }
                                                                                    var v75 uint32 = 2147483648
                                                                                    _ = v73
                                                                                    _ = v75
                                                                                } else {
                                                                                    if false {
                                                                                        if true {
                                                                                            v76 := GenS6{F0: 42}
                                                                                            v5 = v51
                                                                                            if true {
                                                                                                var v77 int32 = 0
                                                                                                _ = v77
                                                                                            } else {
                                                                                                v2 -= 0
                                                                                                fmt.Println(v67)
                                                                                            }
                                                                                            _ = v76
                                                                                        } else {
                                                                                            if false {
                                                                                                v52 = v52
                                                                                                var v78 uint64 = (v67 + (v67 >> ((18446744073709551615) & 63)))
                                                                                                switch v51 {
                                                                                                case "a":
                                                                                                    v53 = 2
                                                                                                case "foo bar":
                                                                                                    switch 337 {
                                                                                                    case -1:
                                                                                                        if false {
                                                                                                            v2 += (fi21 | v2)
                                                                                                            v53 -= 351
                                                                                                            var v79 uint = 18446744073709551615
                                                                                                            _ = v79
                                                                                                        }
                                                                                                    case 0:
                                                                                                        fmt.Println(fi21)
                                                                                                    default:
                                                                                                        fmt.Println(v20, v53)
                                                                                                        _ = len(v1)
                                                                                                        switch "" {
                                                                                                        case "hello":
                                                                                                            fmt.Println(v52)
                                                                                                        case "":
                                                                                                            v4 = "hello"
                                                                                                        case "\n\t":
                                                                                                            var v80 int = (fi21 ^ 202)
                                                                                                            _ = v80
                                                                                                        }
                                                                                                        fmt.Println(v51)
                                                                                                        v81 := GenS2{F0: (2147483647 < -600), F1: (v20 << ((1) & 63))}
                                                                                                        _ = v81
                                                                                                    }
                                                                                                }
                                                                                                var v82 bool = (925 <= 0)
                                                                                                if (v82 || (2147483647 != 663)) {
                                                                                                    var v83 uint8 = (v52 + v52)
                                                                                                    _ = v83
                                                                                                } else {
                                                                                                    var v84 int64 = 0
                                                                                                    fmt.Println(v4, v53, v82)
                                                                                                    _ = v84
                                                                                                }
                                                                                                var v85 int32 = -298
                                                                                                _ = v78
                                                                                                _ = v82
                                                                                                _ = v85
                                                                                            } else {
                                                                                                v51 = v4
                                                                                                fmt.Println(v53)
                                                                                                if (750 != 32767) {
                                                                                                    v86 := GenS7{F0: "", F1: v67, F2: v67}
                                                                                                    var v87 int = -7
                                                                                                    switch 127 {
                                                                                                    case 356:
                                                                                                        v88 := GenS6{F0: 373}
                                                                                                        _ = v88
                                                                                                    case 32767:
                                                                                                        v4 = ("\n\t" + "a")
                                                                                                    }
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = 65535
                                                                                                    }
                                                                                                    fmt.Println(v67, v67, v5)
                                                                                                    var v89 bool = true
                                                                                                    switch v5 {
                                                                                                    case "\n\t":
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = 184
                                                                                                        }
                                                                                                    default:
                                                                                                        var v90 int64 = -346
                                                                                                        _ = v90
                                                                                                    }
                                                                                                    var v91 int32 = -1
                                                                                                    _ = v86
                                                                                                    _ = v87
                                                                                                    _ = v89
                                                                                                    _ = v91
                                                                                                }
                                                                                            }
                                                                                            v20 += v20
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = 2
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if (false || (-942 != -1)) {
                                                                                            v3 += -1
                                                                                            if (42 <= 1) {
                                                                                                v92 := GenS1{F0: v52}
                                                                                                v93 := GenS4{F0: -32768, F1: v2}
                                                                                                var v94 string = (v51 + v4)
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                                _ = v94
                                                                                            } else {
                                                                                                fmt.Println(v2, v52)
                                                                                            }
                                                                                            if (65535 == 42) {
                                                                                                v51 = "foo bar"
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v95 float64 = (((0.1 / -2.25) - ((-2.25 / 1.5) / 3.14159)) / (1.5 - ((-2.25 / 1.0) / 1.0)))
                                                                                _ = v95
                                                                            }
                                                                            var v96 int64 = 1
                                                                            v3 += (v3 | 32767)
                                                                            _ = v96
                                                                        } else {
                                                                            v4 = "\n\t"
                                                                            if false {
                                                                                v20 += (v20 >> ((4294967295) & 63))
                                                                                v97 := GenS1{F0: (v52 >> ((65535) & 63))}
                                                                                _ = v97
                                                                            }
                                                                            v53 -= (v53 % ((v53) | 1))
                                                                        }
                                                                    }
                                                                    v1 = append(v1, 42)
                                                                }
                                                            }
                                                            _ = v67
                                                        }
                                                    }
                                                }
                                            default:
                                                if ((v52 << ((255) & 63)) < 2) {
                                                    if true {
                                                        v2 -= (fi21 ^ -1)
                                                    }
                                                    var v98 uint64 = 276
                                                    if false {
                                                        v52 += (v52 >> ((2) & 63))
                                                        if (v53 <= (^(v53))) {
                                                            v5 = ((v51 + ((v4 + v5) + ("foo bar" + v5))) + v51)
                                                        } else {
                                                            switch "" {
                                                            case "a":
                                                                v5 = ("\n\t" + "hello")
                                                            case "foo bar":
                                                                fmt.Println(v52, v4, fi50)
                                                            }
                                                            var v99 int32 = 103
                                                            _ = v99
                                                        }
                                                        var v100 int16 = 1
                                                        _ = v100
                                                    }
                                                    v101 := GenS8{F0: ((!(!(42 >= (v53 % (((v53 << ((361) & 63))) | 1))))) && (false || true)), F1: 2, F2: (v53 * v53)}
                                                    v2 = (-(v2))
                                                    _ = v101
                                                    _ = v98
                                                } else {
                                                    var v102 float64 = 3.14159
                                                    _ = v102
                                                }
                                            }
                                        default:
                                            var v103 float32 = (float32(0.0) - 1e10)
                                            _ = v103
                                        }
                                        if false {
                                            var v104 uint16 = 457
                                            fmt.Println(v2, fi21)
                                            v105 := GenS9{F0: (^(v52))}
                                            _ = v104
                                            _ = v105
                                        }
                                        var v106 uint8 = (v52 + v52)
                                        if (!(!false)) {
                                            if (true || (v20 > 127)) {
                                                if false {
                                                    var v107 bool = (2 >= 18446744073709551615)
                                                    var v108 float64 = 0.0
                                                    var v109 int64 = -9223372036854775808
                                                    _ = v107
                                                    _ = v108
                                                    _ = v109
                                                } else {
                                                    v110 := GenS4{F0: -32768, F1: (-(fi50))}
                                                    switch fi21 {
                                                    case 9223372036854775807:
                                                        v1 = append(v1, 2147483648)
                                                        var v111 uint = 18446744073709551615
                                                        v2 = (^(fi21))
                                                        _ = v111
                                                    case 4294967295:
                                                        fmt.Println(v3)
                                                    case 498:
                                                        var v112 int32 = 2147483647
                                                        _ = v112
                                                    }
                                                    _ = v110
                                                }
                                                fmt.Println(v52, v20, fi50)
                                                if (!true) {
                                                    v113 := GenS7{F0: (v51 + v51), F1: 4294967295, F2: 250}
                                                    _ = v113
                                                }
                                            }
                                            v1 = append(v1, 42)
                                            switch ((v4 + "foo bar") + v4) {
                                            case "hello":
                                                var v114 uint8 = (v52 >> ((252) & 63))
                                                _ = v114
                                            case "":
                                                if false {
                                                    v52 -= v52
                                                    if true {
                                                        fmt.Println(v52, v51)
                                                    }
                                                }
                                            }
                                            v3 -= fi50
                                            if len(v1) > 0 {
                                                v1[0] = 2
                                            }
                                        }
                                        switch (fi50 % ((32767) | 1)) {
                                        case 42:
                                            var v115 uint64 = 42
                                            _ = v115
                                        }
                                        _ = v106
                                        _ = v61
                                    }
                                case 42:
                                    var v116 int64 = 260
                                    _ = v116
                                case 27:
                                    var v117 int8 = -70
                                    _ = v117
                                }
                            } else {
                                var v118 bool = true
                                var v119 float32 = 0.1
                                var v120 uint = 156
                                if (v53 > v53) {
                                    if (2 < 1) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if false {
                                            if (false || (563 != uint32(v52))) {
                                                v20 = v20
                                            }
                                            fmt.Println(v119, v52)
                                        } else {
                                            v121 := GenS10{F0: uint(v52), F1: -128, F2: 1}
                                            _ = v121
                                        }
                                        var v122 uint64 = 0
                                        _ = v122
                                    }
                                }
                                _ = v118
                                _ = v119
                                _ = v120
                            }
                        case 0:
                            v5 = v5
                        }
                        _ = v54
                        _ = v55
                    } else {
                        v123 := GenS7{F0: v4, F1: 4294967295, F2: 1}
                        var v124 int64 = -9223372036854775808
                        _ = v123
                        _ = v124
                    }
                    var v125 int = (fi21 >> ((18446744073709551615) & 63))
                    _ = fi50
                    _ = v125
                    _ = v51
                    _ = v52
                    _ = v53
                }
                v126 := GenS11{F0: v3}
                v20 -= v20
                _ = v126
            }
            _ = fi21
        }
        _ = v20
    } else {
        if false {
            if (false && true) {
                var v127 int64 = -1
                _ = v127
            } else {
                var v128 uint64 = 927
                v129 := GenS4{F0: 0, F1: (v2 / (((v2 / ((v3) | 1))) | 1))}
                _ = v128
                _ = v129
            }
            var v130 int64 = -1
            v4 = v5
            _ = v130
        }
        var v131 string = v5
        _ = v131
    }
    v4 = v4
    switch (v2 - v3) {
    case 9223372036854775807:
        switch v5 {
        case "foo bar":
            v5 = "foo bar"
        case "a":
            v5 = v5
        case "hello":
            var v132 uint64 = 2
            _ = v132
        default:
            var v133 int64 = 1
            _ = v133
        }
    case 398:
        fmt.Println(v2, v3, v5)
    case -9223372036854775808:
        for fi134 := 0; fi134 < 4; fi134++ {
            v5 = ((v5 + "\n\t") + ("\n\t" + (v4 + ("foo bar" + ("hello" + v5)))))
            var v135 string = ((v5 + "a") + "hello")
            _ = fi134
            _ = v135
        }
    }
    v5 = ("" + v5)
    fmt.Println(v5, v5, v2)
    if (true && (128 == 160)) {
        switch v4 {
        case "":
            fmt.Println(v3, v5, v2)
        case "hello":
            switch v2 {
            case 9223372036854775807:
                for fi136 := 0; fi136 < 12; fi136++ {
                    if (255 <= 128) {
                        v1 = append(v1, 0)
                    } else {
                        switch v4 {
                        case "foo bar":
                            _ = len(v1)
                        case "a":
                            fmt.Println(v5)
                        case "\n\t":
                            v137 := GenS5{F0: 32767, F1: 0, F2: -2147483648}
                            _ = v137
                        default:
                            for fi138 := 0; fi138 < 5; fi138++ {
                                fmt.Println(v2, v5)
                                _ = fi138
                            }
                        }
                        if true {
                            v2 += fi136
                            v5 = v5
                            switch v4 {
                            case "a":
                                for fi139 := 0; fi139 < 14; fi139++ {
                                    if false {
                                        var v140 int64 = 1
                                        v3 += int(v140)
                                        v141 := GenS2{F0: ((-531 != (-(v140))) && (32767 <= 42)), F1: 94}
                                        _ = v140
                                        _ = v141
                                    } else {
                                        v142 := GenS7{F0: ("\n\t" + ""), F1: 1, F2: 0}
                                        v143 := GenS1{F0: 1}
                                        _ = v142
                                        _ = v143
                                    }
                                    var v144 uint32 = 1
                                    v4 = ("hello" + "a")
                                    _ = fi139
                                    _ = v144
                                }
                            case "hello":
                                if (42 <= -24) {
                                    for fi145 := 0; fi145 < 11; fi145++ {
                                        var v146 int64 = 368
                                        var v147 int = fi136
                                        _ = fi145
                                        _ = v146
                                        _ = v147
                                    }
                                } else {
                                    var v148 bool = true
                                    v4 = ("foo bar" + ("" + ""))
                                    if v148 {
                                        for fi149 := 0; fi149 < 1; fi149++ {
                                            switch v4 {
                                            case "\n\t":
                                                var v150 int16 = 32767
                                                _ = v150
                                            }
                                            _ = len(v1)
                                            var v151 string = ("hello" + "")
                                            if (-32768 == -17) {
                                                v2 = 32767
                                            }
                                            var v152 uint = 597
                                            if true {
                                                var v153 uint64 = 1
                                                v151 = v4
                                                if (42 == 2) {
                                                    var v154 float64 = -2.25
                                                    _ = v154
                                                } else {
                                                    v155 := GenS10{F0: v152, F1: -1, F2: -453}
                                                    switch 42 {
                                                    case 42:
                                                        if false {
                                                            if v148 {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 968
                                                                }
                                                                var v156 int16 = 281
                                                                _ = v156
                                                            } else {
                                                                v157 := GenS8{F0: (!v148), F1: v152, F2: 226}
                                                                if v148 {
                                                                    v5 = ((v5 + v5) + v5)
                                                                } else {
                                                                    v158 := GenS5{F0: 1, F1: 431, F2: fi136}
                                                                    _ = v158
                                                                }
                                                                var v159 uint16 = 42
                                                                _ = v157
                                                                _ = v159
                                                            }
                                                            switch v151 {
                                                            case "":
                                                                v153 -= 18446744073709551615
                                                            default:
                                                                v148 = false
                                                            }
                                                        } else {
                                                            v2 = 1
                                                            if (9223372036854775807 == 9223372036854775807) {
                                                                v3 -= v3
                                                                v160 := GenS11{F0: 2147483647}
                                                                var v161 uint = 2
                                                                _ = v160
                                                                _ = v161
                                                            }
                                                            v152 = v152
                                                        }
                                                    case 2147483647:
                                                        switch (fi136 ^ 127) {
                                                        case 676:
                                                            if len(v1) > 0 {
                                                                v1[0] = 0
                                                            }
                                                        case -2147483648:
                                                            var v162 float32 = 0.1
                                                            _ = v162
                                                        }
                                                    case 287:
                                                        fmt.Println(fi136, v148, v3)
                                                    default:
                                                        var v163 int8 = -128
                                                        _ = v163
                                                    }
                                                    _ = v155
                                                }
                                                _ = v153
                                            } else {
                                                if false {
                                                    v3 = v2
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        fmt.Println(v5, v2, v151)
                                                    case "":
                                                        var v164 string = ""
                                                        _ = v164
                                                    default:
                                                        fmt.Println(v151)
                                                    }
                                                }
                                                var v165 uint64 = 18446744073709551615
                                                _ = v165
                                            }
                                            v148 = true
                                            _ = fi149
                                            _ = v151
                                            _ = v152
                                        }
                                        v5 = v5
                                    }
                                    _ = v148
                                }
                            case "\n\t":
                                switch v2 {
                                case 42:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 128:
                                    for fi166 := 0; fi166 < 11; fi166++ {
                                        if false {
                                            var v167 bool = ((-866 <= 1) || false)
                                            v168 := GenS11{F0: (v3 - -9223372036854775808)}
                                            var v169 string = v4
                                            _ = v167
                                            _ = v168
                                            _ = v169
                                        } else {
                                            v2 += 9223372036854775807
                                            if false {
                                                _ = len(v1)
                                                var v170 int32 = 385
                                                _ = v170
                                            }
                                            v4 = v5
                                        }
                                        v2 -= v3
                                        _ = fi166
                                    }
                                case 127:
                                    v5 = v4
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = 65535
                                    }
                                }
                            }
                        }
                        for fi171 := 0; fi171 < 8; fi171++ {
                            var v172 int = (^(fi136))
                            _ = len(v1)
                            var v173 bool = (true && false)
                            _ = fi171
                            _ = v172
                            _ = v173
                        }
                    }
                    for fi174 := 0; fi174 < 10; fi174++ {
                        var v175 uint64 = 1
                        if (true || true) {
                            fmt.Println(fi136, v3)
                            v176 := GenS12{F0: 42, F1: 42}
                            _ = v176
                        }
                        if (959 == v2) {
                            if false {
                                var v177 float32 = 1e10
                                _ = v177
                            } else {
                                fmt.Println(v4, v175, v5)
                            }
                            var v178 uint = 4294967295
                            v5 = "foo bar"
                            _ = v178
                        } else {
                            var v179 float32 = -2.25
                            var v180 int = (-(fi174))
                            _ = v179
                            _ = v180
                        }
                        _ = fi174
                        _ = v175
                    }
                    var v181 int16 = -32768
                    _ = fi136
                    _ = v181
                }
            case 32767:
                v182 := GenS13{F0: -1, F1: v2}
                _ = v182
            case -2147483648:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                var v183 uint16 = 42
                _ = v183
            }
        default:
            if true {
                for fi184 := 0; fi184 < 12; fi184++ {
                    switch v3 {
                    case 638:
                        var v185 uint16 = 1
                        _ = v185
                    case -601:
                        if (4294967295 >= 450) {
                            v5 = ("a" + (v4 + "hello"))
                            fmt.Println(v2, fi184, v3)
                            v3 = v3
                        }
                    }
                    if false {
                        v2 += (fi184 * v2)
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v3, v5)
                    }
                    _ = fi184
                }
            } else {
                fmt.Println(v2, v5, v2)
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi186 := 0; fi186 < 11; fi186++ {
        v2 = (v2 ^ fi186)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi186
    }
    fmt.Println(v3, v2, v3)
    if true {
        switch -221 {
        case 42:
            if false {
                v187 := GenS10{F0: 42, F1: 0, F2: 42}
                var v188 int64 = -1
                _ = v187
                _ = v188
            } else {
                v2 -= (-(v2))
            }
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -9223372036854775808:
            v189 := GenS9{F0: 2}
            _ = v189
        default:
            switch (v3 + (v2 ^ -1)) {
            case -781:
                if ((!true) || false) {
                    var v190 uint = 723
                    v191 := GenS4{F0: 1, F1: v2}
                    for fi192 := 0; fi192 < 18; fi192++ {
                        v3 = 392
                        _ = fi192
                    }
                    _ = v190
                    _ = v191
                } else {
                    for fi193 := 0; fi193 < 7; fi193++ {
                        for fi194 := 0; fi194 < 12; fi194++ {
                            switch (v4 + v4) {
                            case "\n\t":
                                v1 = append(v1, 1)
                                if (!(319 == 0)) {
                                    if true {
                                        var v195 uint = 845
                                        var v196 float64 = 1.5
                                        _ = v195
                                        _ = v196
                                    } else {
                                        v2 = (v3 << ((18446744073709551615) & 63))
                                        v1 = append(v1, 535)
                                    }
                                    switch fi193 {
                                    case -982:
                                        v197 := GenS14{F0: (true && (!true)), F1: true}
                                        _ = v197
                                    case 32768:
                                        var v198 uint = 2
                                        _ = v198
                                    case -1:
                                        v199 := GenS15{F0: 1}
                                        _ = v199
                                    default:
                                        if len(v1) > 0 {
                                            v1[0] = 2147483647
                                        }
                                    }
                                }
                                if (563 < 2) {
                                    v4 = ("" + (("" + ((("hello" + (v4 + (("foo bar" + ("a" + "hello")) + ((v5 + v5) + v5)))) + "a") + "")) + v4))
                                } else {
                                    var v200 int8 = 1
                                    var v201 int32 = -1
                                    var v202 int16 = 32767
                                    v203 := GenS16{F0: 0}
                                    _ = v200
                                    _ = v201
                                    _ = v202
                                    _ = v203
                                }
                            case "hello":
                                v204 := GenS17{F0: 98, F1: 255, F2: 42}
                                _ = v204
                            default:
                                var v205 int64 = 42
                                _ = v205
                            }
                            if true {
                                var v206 uint8 = 243
                                switch v4 {
                                case "":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    v4 = "a"
                                }
                                v2 = (v2 >> ((698) & 63))
                                _ = v206
                            }
                            var v207 float32 = ((0.1 + 3.14159) / 1.0)
                            fmt.Println(v3)
                            _ = fi194
                            _ = v207
                        }
                        v208 := GenS4{F0: 1000, F1: v3}
                        var v209 float64 = 0.1
                        _ = fi193
                        _ = v208
                        _ = v209
                    }
                    _ = len(v1)
                    if (!(!false)) {
                        v2 = (v3 * (-(v3)))
                    }
                }
            }
        }
        v3 = (^(v3))
        for fi210 := 0; fi210 < 11; fi210++ {
            fmt.Println(v4, v4)
            fmt.Println(v5, v3, v5)
            _ = fi210
        }
    }
    for fi211 := 0; fi211 < 10; fi211++ {
        switch ("foo bar" + "a") {
        case "hello":
            if len(v1) > 0 {
                v1[0] = 407
            }
        case "a":
            switch (fi211 >> ((18446744073709551615) & 63)) {
            case -523:
                if (!false) {
                    fmt.Println(fi211, v2, v2)
                }
            case -125:
                var v212 float32 = -2.25
                _ = v212
            }
        default:
            v3 = fi211
        }
        _ = fi211
    }
    v1 = append(v1, 65535)
    if true {
        switch 127 {
        case -324:
            var v213 uint8 = 160
            _ = v213
        case -864:
            _ = len(v1)
        case -386:
            fmt.Println(v3, v5)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
