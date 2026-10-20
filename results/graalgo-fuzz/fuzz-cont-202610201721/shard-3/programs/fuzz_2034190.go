package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 uint8; F2 bool }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 uint; F1 uint64; F2 uint }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint8; F1 int; F2 uint64 }
type GenS14 struct { F0 int64; F1 uint16; F2 int8 }
type GenS15 struct { F0 uint32; F1 uint16 }
type GenS16 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int32; F1 int; F2 int64 }

func main() {
    var v1 string = ("\n\t" + ("hello" + ""))
    var v2 uint32 = 1
    var v3 float32 = ((((1.5 + 0.0) - 1e10) + 1.5) / 3.14159)
    v4 := []int{-181, 1, 4294967295}
    switch 127 {
    case 9223372036854775807:
        for fi5 := 0; fi5 < 10; fi5++ {
            switch fi5 {
            case -417:
                for fi6 := 0; fi6 < 10; fi6++ {
                    if true {
                        fmt.Println(fi6, v3, v3)
                    } else {
                        switch v1 {
                        case "foo bar":
                            v7 := GenS1{F0: "foo bar", F1: -78}
                            _ = v7
                        default:
                            var v8 uint = 255
                            _ = v8
                        }
                        if true {
                            switch "\n\t" {
                            case "a":
                                var v9 uint32 = 2
                                _ = v9
                            case "\n\t":
                                v3 += 1e10
                                if (fi6 < (fi6 ^ 20)) {
                                    if true {
                                        v3 += v3
                                        switch (fi6 & fi6) {
                                        case 9223372036854775807:
                                            var v10 uint16 = 2
                                            _ = v10
                                        case 4294967295:
                                            var v11 uint = 42
                                            _ = v11
                                        case -1:
                                            var v12 uint16 = 65535
                                            _ = v12
                                        }
                                        var v13 uint16 = 65535
                                        _ = v13
                                    }
                                    v14 := GenS2{F0: 65535, F1: 75}
                                    var v15 string = v1
                                    _ = v14
                                    _ = v15
                                } else {
                                    fmt.Println(v2, v2, fi5)
                                    if len(v4) > 0 {
                                        v4[0] = (fi6 | fi5)
                                    }
                                    v2 = v2
                                    v4 = append(v4, 32768)
                                    var v16 int64 = -1
                                    if false {
                                        v17 := GenS1{F0: (("hello" + (v1 + v1)) + v1), F1: 127}
                                        switch v1 {
                                        case "foo bar":
                                            var v18 uint16 = 65535
                                            _ = v18
                                        }
                                        _ = v17
                                    }
                                    _ = v16
                                }
                                var v19 int8 = 1
                                _ = v19
                            }
                            var v20 int64 = 9223372036854775807
                            _ = v20
                        } else {
                            v3 = 1.5
                            fmt.Println(v3, fi6)
                            fmt.Println(fi6, v1, v3)
                        }
                        fmt.Println(fi6, fi6)
                    }
                    v4 = append(v4, (fi6 | (^(fi6))))
                    fmt.Println(v2, v2, v1)
                    _ = fi6
                }
            case 32768:
                for fi21 := 0; fi21 < 10; fi21++ {
                    if len(v4) > 0 {
                        v4[0] = (fi21 & (fi5 % (((fi21 - (^(fi5)))) | 1)))
                    }
                    switch "" {
                    case "foo bar":
                        v3 += -2.25
                    }
                    if (!true) {
                        if (v3 < float32(1e10)) {
                            _ = len(v4)
                            fmt.Println(v2, v2, v2)
                            var v22 int64 = -9223372036854775808
                            switch v1 {
                            case "a":
                                v22 = v22
                            case "\n\t":
                                v2 += 2
                            default:
                                if false {
                                    var v23 int32 = -620
                                    _ = v23
                                } else {
                                    if true {
                                        v3 += (float32((0.1 / float64(v3))) * v3)
                                    }
                                }
                            }
                            v1 = (v1 + ("foo bar" + v1))
                            v24 := GenS3{F0: (v2 | (v2 % ((v2) | 1))), F1: 0, F2: 4294967295}
                            var v25 int32 = -2147483648
                            var v26 uint16 = uint16(v2)
                            _ = v22
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v2 += v2
                    } else {
                        v2 += (v2 >> ((441) & 63))
                    }
                    _ = fi21
                }
            case -2147483648:
                v3 -= (1e10 / v3)
            default:
                if (821 <= 1000) {
                    if (true && (100 <= 42)) {
                        var v27 uint64 = 4294967295
                        var v28 uint16 = 32767
                        v27 += v27
                        _ = v27
                        _ = v28
                    }
                    switch "\n\t" {
                    case "a":
                        v1 = "a"
                    }
                } else {
                    if false {
                        fmt.Println(v2, v1)
                    } else {
                        v1 = ("a" + "\n\t")
                        v3 -= v3
                        v2 -= v2
                    }
                }
            }
            switch ((("" + "foo bar") + (((v1 + v1) + "") + v1)) + ((((v1 + "\n\t") + v1) + "a") + v1)) {
            case "hello":
                fmt.Println(fi5)
            default:
                var v29 float32 = 1e10
                _ = v29
            }
            _ = fi5
        }
    case -168:
        for fi30 := 0; fi30 < 14; fi30++ {
            _ = len(v4)
            switch v1 {
            case "foo bar":
                var v31 int = (fi30 * fi30)
                _ = v31
            case "":
                if ((fi30 << ((255) & 63)) != fi30) {
                    v2 -= (v2 & v2)
                    for fi32 := 0; fi32 < 6; fi32++ {
                        switch "hello" {
                        case "":
                            if (true && (1e10 != (((0.0 * float64(-2.25)) + 0.1) + float64(0.0)))) {
                                fmt.Println(fi30, v3, v3)
                            }
                        case "\n\t":
                            switch "a" {
                            case "foo bar":
                                switch v1 {
                                case "hello":
                                    v4 = append(v4, (fi32 << ((1) & 63)))
                                default:
                                    if (false && (1.5 == v3)) {
                                        v1 = ""
                                        var v33 int64 = 9223372036854775807
                                        v4 = append(v4, fi30)
                                        _ = v33
                                    } else {
                                        fmt.Println(v1, fi30, fi30)
                                        v2 -= (v2 << ((42) & 63))
                                        fmt.Println(v3, fi32, v2)
                                    }
                                }
                            case "a":
                                v34 := GenS4{F0: false}
                                _ = v34
                            default:
                                var v35 bool = true
                                _ = v35
                            }
                        case "hello":
                            if false {
                                v4 = append(v4, fi30)
                            }
                        default:
                            if ((0.1 * 3.14159) != -2.25) {
                                v2 -= (v2 & 2147483647)
                                if false {
                                    var v36 string = "hello"
                                    _ = v36
                                } else {
                                    var v37 int16 = -17
                                    var v38 int64 = -1
                                    _ = v37
                                    _ = v38
                                }
                                var v39 float32 = 0.0
                                _ = v39
                            } else {
                                fmt.Println(v2, v3)
                                var v40 int64 = -822
                                v4 = append(v4, fi30)
                                v1 = v1
                                _ = v40
                            }
                        }
                        fmt.Println(fi30, fi32, fi32)
                        _ = fi32
                    }
                } else {
                    if false {
                        for fi41 := 0; fi41 < 19; fi41++ {
                            var v42 int64 = 0
                            switch 9223372036854775807 {
                            case 127:
                                v43 := GenS1{F0: "\n\t", F1: 42}
                                _ = v43
                            default:
                                v1 = v1
                            }
                            _ = fi41
                            _ = v42
                        }
                    }
                    fmt.Println(v1, v3)
                }
            case "hello":
                v4 = append(v4, fi30)
            }
            _ = fi30
        }
    }
    var v44 string = "a"
    fmt.Println(v3, v44, v2)
    if false {
        switch v1 {
        case "foo bar":
            if false {
                var v45 uint64 = 397
                v46 := GenS5{F0: 197}
                switch v1 {
                case "a":
                    v2 += 273
                case "":
                    switch 9223372036854775807 {
                    case 1:
                        v4 = append(v4, 127)
                    }
                case "hello":
                    fmt.Println(v2, v45)
                default:
                    fmt.Println(v44, v44)
                }
                _ = v45
                _ = v46
            }
        default:
            for fi47 := 0; fi47 < 17; fi47++ {
                if len(v4) > 0 {
                    v4[0] = (fi47 | 2147483647)
                }
                fmt.Println(v2, v44)
                for fi48 := 0; fi48 < 19; fi48++ {
                    var v49 uint16 = 1
                    v50 := GenS1{F0: "", F1: 100}
                    switch (v1 + v44) {
                    case "a":
                        if ((fi47 & (fi47 / (((fi48 ^ fi47)) | 1))) != 0) {
                            var v51 int8 = -113
                            if false {
                                var v52 float32 = 1e10
                                _ = v52
                            } else {
                                switch 1 {
                                case 32767:
                                    var v53 uint32 = (v2 << ((18446744073709551615) & 63))
                                    _ = v53
                                case 4294967295:
                                    v3 += v3
                                default:
                                    var v54 int32 = 42
                                    _ = v54
                                }
                                if (3.14159 < v3) {
                                    if false {
                                        v3 = (v3 / v3)
                                        if len(v4) > 0 {
                                            v4[0] = (fi48 * fi48)
                                        }
                                        v55 := GenS2{F0: 18446744073709551615, F1: (v51 % (((v51 * v51)) | 1))}
                                        _ = v55
                                    }
                                }
                            }
                            _ = v51
                        } else {
                            v4 = append(v4, fi48)
                            if (v49 > 1) {
                                if len(v4) > 0 {
                                    v4[0] = fi47
                                }
                                v56 := GenS4{F0: false}
                                v57 := GenS3{F0: v2, F1: -32768, F2: 42}
                                _ = v56
                                _ = v57
                            } else {
                                if false {
                                    v58 := GenS1{F0: ("hello" + v44), F1: 59}
                                    _ = v58
                                } else {
                                    v44 = "hello"
                                    _ = len(v4)
                                    switch (fi47 - (fi47 >> ((417) & 63))) {
                                    case -576:
                                        if true {
                                            if false {
                                                v1 = v1
                                            } else {
                                                if (!((v2 << ((2) & 63)) >= uint32(v49))) {
                                                    v4 = append(v4, fi47)
                                                } else {
                                                    var v59 int16 = -32768
                                                    var v60 uint16 = v49
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v59 = v59
                                                    v61 := GenS6{F0: 0, F1: 117, F2: true}
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                        } else {
                                            if true {
                                                var v62 int64 = -9223372036854775808
                                                v63 := GenS1{F0: ((v1 + "\n\t") + (v1 + "hello")), F1: 100}
                                                switch (fi47 % ((42) | 1)) {
                                                case 725:
                                                    fmt.Println(v62)
                                                default:
                                                    v64 := GenS2{F0: 255, F1: 57}
                                                    if (v44 > v44) {
                                                        var v65 float32 = (v3 - v3)
                                                        if ((!(false && false)) || true) {
                                                            v49 = (v49 >> ((722) & 63))
                                                        }
                                                        _ = v65
                                                    }
                                                    var v66 uint64 = 4294967295
                                                    _ = v64
                                                    _ = v66
                                                }
                                                _ = v62
                                                _ = v63
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = (fi48 >> ((uint(v49)) & 63))
                                            }
                                            v49 += (v49 ^ v49)
                                            var v67 int32 = 979
                                            if true {
                                                switch fi48 {
                                                case 32767:
                                                    v1 = (v44 + v44)
                                                case 128:
                                                    v68 := GenS7{F0: 0, F1: 134}
                                                    _ = v68
                                                }
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = len(v4)
                                            if true {
                                                if ((true && (-1 >= 1)) && true) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v69 int32 = (v67 + (-(v67)))
                                                    if true {
                                                        if len(v4) > 0 {
                                                            v4[0] = (-(fi47))
                                                        }
                                                        if ((((!false) || (0.0 > 3.14159)) || (42 != 0)) || (!true)) {
                                                            v1 = v1
                                                            switch (fi48 ^ 128) {
                                                            case 9223372036854775807:
                                                                if (float32(1.5) < v3) {
                                                                    var v70 uint = 1
                                                                    var v71 int16 = 1000
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                            default:
                                                                var v72 int = (fi47 * fi47)
                                                                _ = v72
                                                            }
                                                            v73 := GenS8{F0: -610}
                                                            switch v44 {
                                                            case "":
                                                                var v74 int32 = 0
                                                                _ = v74
                                                            case "a":
                                                                v49 = 42
                                                                v75 := GenS4{F0: (((1.5 * -2.25) >= 1.5) && (0 >= (v67 / (((-(v67))) | 1))))}
                                                                _ = v75
                                                            default:
                                                                v1 = (v1 + v1)
                                                            }
                                                            _ = v73
                                                        } else {
                                                            if ((32767 != fi48) || (!(127 == 84))) {
                                                                var v76 uint8 = 125
                                                                _ = v76
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            if ((float64((float32(((float64(((v3 * (-2.25 + (0.0 / 1e10))) - v3)) * 0.0) + 1.5)) - (v3 * float32(0.1)))) / 0.1) <= (1.5 * (1.5 + (1e10 + (1.5 * 0.1))))) {
                                                                if ((false || true) && ((v49 & v49) != v49)) {
                                                                    fmt.Println(v44)
                                                                } else {
                                                                    if (v2 < (v2 ^ (v2 & 124))) {
                                                                        if true {
                                                                            var v77 int = fi47
                                                                            v1 = "foo bar"
                                                                            _ = v77
                                                                        }
                                                                        v78 := GenS5{F0: 255}
                                                                        _ = v78
                                                                    } else {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = fi48
                                                                        }
                                                                        v79 := GenS9{F0: -831}
                                                                        if true {
                                                                            _ = len(v4)
                                                                            if true {
                                                                                var v80 int = (fi48 << ((4294967295) & 63))
                                                                                if true {
                                                                                    var v81 int8 = 10
                                                                                    _ = v81
                                                                                }
                                                                                _ = v80
                                                                            }
                                                                        }
                                                                        _ = v79
                                                                    }
                                                                }
                                                                v3 -= v3
                                                            }
                                                            fmt.Println(v3)
                                                        }
                                                    } else {
                                                        var v82 uint16 = (v49 % ((0) | 1))
                                                        if ((v44 + (("\n\t" + v1) + (v1 + v1))) > "hello") {
                                                            v44 = (v44 + v44)
                                                            var v83 int16 = 183
                                                            _ = v83
                                                        }
                                                        _ = v82
                                                    }
                                                    _ = len(v4)
                                                    _ = v69
                                                } else {
                                                    switch -9223372036854775808 {
                                                    case 610:
                                                        v2 -= (v2 / ((v2) | 1))
                                                    }
                                                    v3 += v3
                                                }
                                                v2 = 2147483647
                                            } else {
                                                if (true || (v49 != 255)) {
                                                    v1 = v44
                                                    switch (fi48 | fi48) {
                                                    case 2147483647:
                                                        v84 := GenS10{F0: false, F1: "hello"}
                                                        _ = v84
                                                    case -1:
                                                        v3 = (0.1 * 1e10)
                                                    case 32767:
                                                        var v85 int32 = -384
                                                        _ = v85
                                                    }
                                                }
                                                if (!((false || (true && (true || false))) && true)) {
                                                    v86 := GenS7{F0: 100, F1: 42}
                                                    _ = v86
                                                } else {
                                                    v49 -= v49
                                                    switch v1 {
                                                    case "\n\t":
                                                        v49 -= 1
                                                    case "a":
                                                        switch "a" {
                                                        case "a":
                                                            var v87 uint = 113
                                                            _ = v87
                                                        case "hello":
                                                            var v88 uint32 = (^(v2))
                                                            _ = v88
                                                        case "":
                                                            var v89 int8 = -128
                                                            _ = v89
                                                        default:
                                                            if true {
                                                                v1 = (v44 + (v44 + ""))
                                                                var v90 string = ("\n\t" + v1)
                                                                if len(v4) > 0 {
                                                                    v4[0] = fi47
                                                                }
                                                                _ = v90
                                                            } else {
                                                                var v91 int16 = 1000
                                                                v91 = v91
                                                                v92 := GenS9{F0: v91}
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v93 int = 379
                                                        _ = v93
                                                    default:
                                                        v3 -= v3
                                                    }
                                                }
                                                fmt.Println(v67)
                                                if (!(4294967295 == 724)) {
                                                    v94 := GenS11{F0: 4294967295, F1: 0, F2: 556}
                                                    _ = v94
                                                } else {
                                                    v95 := GenS6{F0: 127, F1: 0, F2: true}
                                                    switch (fi47 << ((42) & 63)) {
                                                    case 0:
                                                        fmt.Println(v1)
                                                    }
                                                    _ = v95
                                                }
                                            }
                                            _ = v67
                                        }
                                    case 32767:
                                        fmt.Println(v49, v2)
                                    case 42:
                                        v96 := GenS6{F0: -31, F1: 42, F2: (935 >= 18446744073709551615)}
                                        _ = v96
                                    }
                                    if ((100 <= 0) && false) {
                                        v4 = append(v4, -717)
                                        if (true || true) {
                                            v97 := GenS9{F0: 110}
                                            if (0 <= -1) {
                                                fmt.Println(v49)
                                            } else {
                                                var v98 int8 = 0
                                                _ = len(v4)
                                                _ = v98
                                            }
                                            _ = v97
                                        }
                                        v99 := GenS12{F0: true}
                                        _ = v99
                                    } else {
                                        _ = len(v4)
                                        v49 -= (^(v49))
                                        switch (-(fi47)) {
                                        case 2147483647:
                                            v4 = append(v4, fi47)
                                        case 128:
                                            fmt.Println(v49)
                                        case 42:
                                            var v100 int = (fi48 / ((fi47) | 1))
                                            _ = v100
                                        default:
                                            if true {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch 9223372036854775807 {
                                                case 32767:
                                                    if true {
                                                        switch v44 {
                                                        case "":
                                                            if false {
                                                                v1 = "foo bar"
                                                                v101 := GenS3{F0: 2, F1: 108, F2: 18446744073709551615}
                                                                if false {
                                                                    _ = len(v4)
                                                                }
                                                                _ = v101
                                                            }
                                                        case "\n\t":
                                                            if len(v4) > 0 {
                                                                v4[0] = 42
                                                            }
                                                        }
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                case 4294967295:
                                                    v3 = float32((float64(((v3 * 0.0) * v3)) + 3.14159))
                                                case 0:
                                                    fmt.Println(fi48)
                                                }
                                                v2 = v2
                                                switch 4294967295 {
                                                case 261:
                                                    v102 := GenS13{F0: 37, F1: 191, F2: 666}
                                                    _ = v102
                                                case 2147483647:
                                                    if len(v4) > 0 {
                                                        v4[0] = fi48
                                                    }
                                                case -1:
                                                    var v103 int64 = -1
                                                    _ = v103
                                                default:
                                                    switch int(v49) {
                                                    case -815:
                                                        var v104 uint32 = (v2 - 754)
                                                        _ = v104
                                                    case 2147483647:
                                                        var v105 uint64 = 1
                                                        _ = v105
                                                    case 0:
                                                        v106 := GenS14{F0: 213, F1: (v49 << ((uint(v2)) & 63)), F2: -55}
                                                        _ = v106
                                                    }
                                                }
                                            } else {
                                                if true {
                                                    var v107 int16 = 1000
                                                    _ = v107
                                                }
                                                v1 = "a"
                                                var v108 uint = 18446744073709551615
                                                if false {
                                                    v109 := GenS15{F0: 65535, F1: v49}
                                                    if (!(42 >= -9223372036854775808)) {
                                                        v44 = ("\n\t" + (v44 + ""))
                                                        _ = len(v4)
                                                    } else {
                                                        v110 := GenS12{F0: (0.0 < v3)}
                                                        if (!((v49 >> (((v108 | (^(v108)))) & 63)) == (v49 % (((v49 + v49)) | 1)))) {
                                                            var v111 bool = true
                                                            if v111 {
                                                                var v112 float64 = (-2.25 + 3.14159)
                                                                _ = v112
                                                            } else {
                                                                if v111 {
                                                                    if (127 == 100) {
                                                                        fmt.Println(v44)
                                                                        v113 := GenS12{F0: (true && true)}
                                                                        _ = v113
                                                                    }
                                                                    fmt.Println(v49, v1, fi48)
                                                                }
                                                                v108 = v108
                                                            }
                                                            v2 = (v2 ^ v2)
                                                            _ = v111
                                                        }
                                                        _ = v110
                                                    }
                                                    _ = v109
                                                }
                                                v108 += (^(v108))
                                                if (434 <= v108) {
                                                    var v114 uint32 = v2
                                                    var v115 int64 = 9223372036854775807
                                                    var v116 int8 = -122
                                                    _ = v114
                                                    _ = v115
                                                    _ = v116
                                                }
                                                _ = v108
                                            }
                                        }
                                        v117 := GenS4{F0: (((v3 != 1.5) && false) && true)}
                                        _ = v117
                                    }
                                    var v118 int64 = 1
                                    var v119 uint32 = (v2 % (((v2 ^ (v2 << ((1) & 63)))) | 1))
                                    v44 = "hello"
                                    _ = v118
                                    _ = v119
                                }
                                v49 += (v49 >> ((2) & 63))
                            }
                            v120 := GenS11{F0: 94, F1: 1, F2: 887}
                            _ = v120
                        }
                    case "":
                        v4 = append(v4, (fi48 << ((255) & 63)))
                    default:
                        var v121 int = (fi47 - fi47)
                        _ = v121
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                _ = fi47
            }
        }
    } else {
        var v122 int64 = 0
        v1 = ("\n\t" + v1)
        _ = v122
    }
    v3 = v3
    if ((v3 > 1e10) && ((1 >= 2) && true)) {
        var v123 uint8 = 2
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi124 := 0; fi124 < 9; fi124++ {
            v123 = (v123 & v123)
            _ = fi124
        }
        _ = v123
    } else {
        var v125 int8 = 1
        v44 = v44
        _ = v125
    }
    if (2 <= 0) {
        if (v1 != "hello") {
            v1 = ("hello" + ((("hello" + (v1 + (v44 + v1))) + v1) + v1))
            var v126 int8 = 45
            if false {
                if true {
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if false {
                            for fi127 := 0; fi127 < 5; fi127++ {
                                switch "a" {
                                case "\n\t":
                                    var v128 float64 = 1.5
                                    _ = v128
                                }
                                if (!(0 == 255)) {
                                    switch (v44 + "") {
                                    case "\n\t":
                                        v126 -= (v126 * (v126 * v126))
                                    case "a":
                                        var v129 float64 = (-2.25 * -2.25)
                                        _ = v129
                                    }
                                    if (v44 >= ("foo bar" + ("hello" + v1))) {
                                        v2 = 65535
                                        if true {
                                            _ = len(v4)
                                            var v130 int = (fi127 << ((65535) & 63))
                                            _ = v130
                                        } else {
                                            if (1 <= 1000) {
                                                v131 := GenS16{F0: 42, F1: ((26 <= v126) || ((((1.5 - ((1.5 * (1.5 * 0.1)) / 1.0)) * 1.5) < 0.0) || (784 > -112))), F2: 239}
                                                for fi132 := 0; fi132 < 1; fi132++ {
                                                    var v133 float64 = 0.1
                                                    v1 = (v44 + v44)
                                                    _ = fi132
                                                    _ = v133
                                                }
                                                var v134 int16 = -32768
                                                _ = v131
                                                _ = v134
                                            }
                                            for fi135 := 0; fi135 < 4; fi135++ {
                                                if true {
                                                    if false {
                                                        v44 = "a"
                                                        switch "a" {
                                                        case "":
                                                            v136 := GenS15{F0: 0, F1: 1}
                                                            _ = v136
                                                        default:
                                                            var v137 int8 = 0
                                                            _ = v137
                                                        }
                                                        if (0 == 47) {
                                                            if false {
                                                                var v138 uint16 = 2
                                                                var v139 int8 = v126
                                                                switch fi127 {
                                                                case 128:
                                                                    fmt.Println(fi135, v126, v1)
                                                                case 127:
                                                                    v140 := GenS17{F0: 1000}
                                                                    _ = v140
                                                                case 567:
                                                                    if false {
                                                                        var v141 uint64 = 968
                                                                        _ = v141
                                                                    }
                                                                }
                                                                var v142 float32 = v3
                                                                v143 := GenS9{F0: 32767}
                                                                _ = v138
                                                                _ = v139
                                                                _ = v142
                                                                _ = v143
                                                            }
                                                            v4 = append(v4, fi135)
                                                            if true {
                                                                if ((v126 * (v126 - -23)) <= 14) {
                                                                    var v144 int8 = v126
                                                                    var v145 uint16 = 1
                                                                    fmt.Println(v1)
                                                                    _ = v144
                                                                    _ = v145
                                                                }
                                                            }
                                                        } else {
                                                            if ((float32(((0.1 - 0.1) * ((float64(0.0) - 0.1) - 3.14159))) * 1.5) >= v3) {
                                                                v146 := GenS10{F0: ((false || true) && false), F1: ("\n\t" + v1)}
                                                                var v147 float32 = v3
                                                                fmt.Println(v44)
                                                                fmt.Println(v126, v1)
                                                                _ = v146
                                                                _ = v147
                                                            }
                                                            var v148 int = (fi127 * 42)
                                                            _ = v148
                                                        }
                                                    } else {
                                                        fmt.Println(fi127, fi127)
                                                    }
                                                    if (-811 != fi127) {
                                                        var v149 int8 = 127
                                                        fmt.Println(v149)
                                                        v150 := GenS14{F0: 9223372036854775807, F1: 32767, F2: (v126 & 1)}
                                                        _ = v149
                                                        _ = v150
                                                    } else {
                                                        v151 := GenS18{F0: int32(fi127), F1: fi127, F2: 376}
                                                        _ = v151
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v152 float32 = 0.0
                                                    _ = v152
                                                }
                                                if (float64(v3) >= ((0.0 - float64((float32(0.1) / ((float32(0.1) / v3) + float32(1.5))))) + 1e10)) {
                                                    if false {
                                                        var v153 int8 = v126
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        switch (^(fi127)) {
                                                        case 537:
                                                            var v154 uint8 = 89
                                                            _ = v154
                                                        case 4294967295:
                                                            if false {
                                                                v155 := GenS11{F0: 1, F1: 27, F2: 4294967295}
                                                                if (!true) {
                                                                    v153 += 42
                                                                } else {
                                                                    fmt.Println(v2, v126)
                                                                    var v156 int = (fi135 + 128)
                                                                    _ = len(v4)
                                                                    _ = v156
                                                                }
                                                                v126 = (v126 / ((1) | 1))
                                                                _ = v155
                                                            }
                                                        case 2147483647:
                                                            v153 = (v153 & v153)
                                                        }
                                                        _ = v153
                                                    }
                                                    v157 := GenS2{F0: 42, F1: (v126 - (-(v126)))}
                                                    if false {
                                                        fmt.Println(fi127, v44, v126)
                                                        v2 -= 529
                                                        fmt.Println(v2, v126)
                                                        if (true || false) {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                    } else {
                                                        var v158 uint8 = 80
                                                        var v159 uint16 = 818
                                                        var v160 uint = uint(v158)
                                                        _ = v158
                                                        _ = v159
                                                        _ = v160
                                                    }
                                                    _ = v157
                                                }
                                                _ = fi135
                                            }
                                        }
                                    } else {
                                        if ((true || (127 >= 248)) || ((((-2147483648 > 0) && false) && true) && false)) {
                                            fmt.Println(v126)
                                            _ = len(v4)
                                            fmt.Println(v126, v126, v126)
                                        }
                                    }
                                    if true {
                                        switch fi127 {
                                        case -2147483648:
                                            if true {
                                                v2 = 2147483647
                                            } else {
                                                v161 := GenS12{F0: (false || (-2147483648 == 1))}
                                                if false {
                                                    var v162 int = fi127
                                                    _ = v162
                                                }
                                                _ = v161
                                            }
                                        case 0:
                                            for fi163 := 0; fi163 < 9; fi163++ {
                                                v164 := GenS7{F0: (v126 % ((127) | 1)), F1: 18446744073709551615}
                                                _ = fi163
                                                _ = v164
                                            }
                                        }
                                    } else {
                                        if false {
                                            var v165 uint8 = 64
                                            _ = v165
                                        }
                                        var v166 uint64 = 18446744073709551615
                                        fmt.Println(v2, v44, v2)
                                        _ = v166
                                    }
                                }
                                v44 = v44
                                _ = fi127
                            }
                            v2 += (v2 + v2)
                        }
                        var v167 int = -2147483648
                        _ = v167
                    } else {
                        if (false || (((v2 >> ((18446744073709551615) & 63)) >= 2) && (((((v3 + v3) / 0.1) - v3) - 1e10) <= 0.1))) {
                            v1 = v1
                            var v168 float32 = v3
                            v169 := GenS13{F0: 1, F1: 9223372036854775807, F2: 42}
                            _ = v168
                            _ = v169
                        } else {
                            switch v1 {
                            case "":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        var v170 int = 1
                        _ = v170
                    }
                } else {
                    var v171 int = 127
                    switch (v171 + (^(v171))) {
                    case 0:
                        if false {
                            v3 = ((0.0 + v3) - v3)
                            v1 = (v44 + "a")
                        } else {
                            if (115 != -894) {
                                v171 = -2147483648
                                var v172 uint8 = 0
                                for fi173 := 0; fi173 < 12; fi173++ {
                                    var v174 int = v171
                                    _ = fi173
                                    _ = v174
                                }
                                _ = v172
                            }
                            fmt.Println(v171)
                        }
                    }
                    if (v44 <= ("" + (v44 + v1))) {
                        var v175 int64 = 0
                        for fi176 := 0; fi176 < 5; fi176++ {
                            for fi177 := 0; fi177 < 14; fi177++ {
                                var v178 string = ((v44 + (v44 + v1)) + (v44 + ""))
                                _ = fi177
                                _ = v178
                            }
                            _ = fi176
                        }
                        _ = v175
                    }
                    _ = v171
                }
                var v179 float32 = v3
                _ = v179
            } else {
                fmt.Println(v126)
            }
            _ = v126
        }
        var v180 int64 = 124
        _ = v180
    } else {
        v181 := GenS6{F0: 0, F1: 0, F2: false}
        var v182 int32 = 2147483647
        _ = v181
        _ = v182
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
}
