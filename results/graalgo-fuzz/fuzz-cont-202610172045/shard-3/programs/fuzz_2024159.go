package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 int; F1 uint; F2 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int64 }

func main() {
    var v1 int16 = 1000
    var v2 float32 = 1e10
    v3 := make(map[uint32]string)
    var v4 uint16 = 0
    var v5 float64 = -2.25
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    v3[2] = "hello"
    if true {
        if false {
            fmt.Println(v5)
            for fi8 := 0; fi8 < 4; fi8++ {
                if (float32((1e10 * -2.25)) != float32(v5)) {
                    if true {
                        fmt.Println(v5, fi8)
                        v5 -= 1e10
                        for fi9 := 0; fi9 < 10; fi9++ {
                            delete(v3, 0)
                            if (false && (!false)) {
                                if true {
                                    v10 := []int{}
                                    _ = v10
                                } else {
                                    var v11 int64 = 1
                                    _ = v11
                                }
                            } else {
                                v12 := GenS1{F0: -1}
                                if (65 == 1) {
                                    var v13 uint64 = 0
                                    fmt.Println(v13)
                                    var v14 bool = (!false)
                                    _ = v13
                                    _ = v14
                                }
                                v15 := GenS2{F0: 42, F1: false}
                                _ = v12
                                _ = v15
                            }
                            v4 -= (^(v4))
                            delete(v3, 2147483648)
                            fmt.Println(v4, v1)
                            _ = fi9
                        }
                    }
                    fmt.Println(v1)
                    var v16 float32 = (v2 - (float32(v5) + v2))
                    _ = v16
                } else {
                    switch 32767 {
                    case 127:
                        var v17 int = 42
                        _ = v17
                    case -153:
                        if true {
                            for fi18 := 0; fi18 < 7; fi18++ {
                                v1 = v1
                                fmt.Println(v2, v1, fi8)
                                v19 := GenS2{F0: 18446744073709551615, F1: true}
                                _ = fi18
                                _ = v19
                            }
                        }
                    }
                    var v20 uint64 = 1
                    switch (fi8 * -1) {
                    case 42:
                        var v21 bool = true
                        _ = v21
                    case -541:
                        fmt.Println(v1, v2)
                    case 1:
                        var v22 int32 = -896
                        _ = v22
                    default:
                        switch fi8 {
                        case -398:
                            v23 := []int{1, fi8}
                            _ = v23
                        case 42:
                            if true {
                                v24 := []int{2147483647}
                                v25 := v3[1]
                                _ = v24
                                _ = v25
                            } else {
                                v26 := GenS1{F0: 100}
                                _ = v26
                            }
                        }
                    }
                    _ = v20
                }
                for fi27 := 0; fi27 < 6; fi27++ {
                    v28 := GenS3{F0: 42, F1: (v1 & (v1 * v1)), F2: 1}
                    _ = fi27
                    _ = v28
                }
                v29 := []int{684, 2147483647}
                _ = fi8
                _ = v29
            }
        }
    } else {
        switch "hello" {
        case "\n\t":
            v30 := GenS4{F0: 284, F1: 1, F2: (v1 << ((42) & 63))}
            _ = v30
        default:
            v31 := GenS5{F0: 42}
            _ = v31
        }
        delete(v3, 4294967295)
        if (!true) {
            for fi32 := 0; fi32 < 17; fi32++ {
                for fi33 := 0; fi33 < 3; fi33++ {
                    v34 := v3[758]
                    var v35 uint64 = 4294967295
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                var v36 int64 = 9223372036854775807
                var v37 int64 = 0
                _ = fi32
                _ = v36
                _ = v37
            }
        } else {
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 uint64 = 2
                v39 = (v39 * v39)
                for fi40 := 0; fi40 < 10; fi40++ {
                    v39 += v39
                    _ = fi40
                }
                _ = fi38
                _ = v39
            }
            var v41 uint = 1
            _ = v41
        }
    }
    var v42 uint = 2
    if false {
        var v43 uint64 = 1
        if true {
            v44 := GenS1{F0: -1}
            var v45 uint8 = uint8(v4)
            v46 := []int{381, 42, 32768}
            _ = v44
            _ = v45
            _ = v46
        } else {
            switch 32768 {
            case -963:
                switch ("hello" + "a") {
                case "\n\t":
                    for fi47 := 0; fi47 < 19; fi47++ {
                        if false {
                            v48 := GenS2{F0: (v43 >> (((v42 ^ v42)) & 63)), F1: false}
                            _ = v48
                        }
                        v49 := GenS2{F0: (v43 << ((v42) & 63)), F1: (1 >= 283)}
                        switch (("foo bar" + "hello") + "hello") {
                        case "foo bar":
                            switch "foo bar" {
                            case "hello":
                                v50 := GenS4{F0: (fi47 ^ (fi47 & (fi47 * fi47))), F1: (v42 | 65535), F2: (^(v1))}
                                _ = v50
                            case "\n\t":
                                v51 := GenS6{F0: v43}
                                _ = v51
                            case "a":
                                v43 = (v43 - (^(v43)))
                            default:
                                v43 += 4294967295
                            }
                        case "":
                            v3[2147483648] = "\n\t"
                        default:
                            var v52 int8 = -1
                            _ = v52
                        }
                        _ = fi47
                        _ = v49
                    }
                case "a":
                    for fi53 := 0; fi53 < 5; fi53++ {
                        if false {
                            v2 = (1.5 * ((((((((0.0 / -2.25) - (float32((1e10 + v5)) / (0.0 - v2))) + 0.0) - (1.5 - v2)) * v2) * -2.25) + 0.0) - 3.14159))
                        }
                        var v54 int32 = -668
                        v42 += (v42 * 65535)
                        _ = fi53
                        _ = v54
                    }
                case "hello":
                    v2 = (((v2 - 1e10) - v2) / float32((float64((v2 + (v2 - ((-2.25 * ((v2 / v2) + float32(float64(((v2 * v2) + 0.1))))) - v2)))) / v5)))
                default:
                    v4 = (v4 / ((v4) | 1))
                }
            case 2147483647:
                for fi55 := 0; fi55 < 9; fi55++ {
                    v56 := GenS3{F0: uint64(v4), F1: (v1 - (v1 * (^(v1)))), F2: v43}
                    var v57 uint32 = 878
                    switch "a" {
                    case "\n\t":
                        var v58 uint16 = 2
                        _ = v58
                    default:
                        for fi59 := 0; fi59 < 5; fi59++ {
                            var v60 uint32 = (v57 / (((^(v57))) | 1))
                            _ = fi59
                            _ = v60
                        }
                    }
                    _ = fi55
                    _ = v56
                    _ = v57
                }
            default:
                v61 := GenS2{F0: 1, F1: false}
                _ = v61
            }
            fmt.Println(v5, v1)
            var v62 uint8 = 128
            _ = v62
        }
        _ = v43
    }
    for fi63 := 0; fi63 < 5; fi63++ {
        if true {
            for fi64 := 0; fi64 < 16; fi64++ {
                var v65 uint8 = 220
                _ = fi64
                _ = v65
            }
        } else {
            v66 := []int{}
            _ = v66
        }
        _ = fi63
    }
    v67 := GenS5{F0: 12}
    switch "hello" {
    case "foo bar":
        for fi68 := 0; fi68 < 6; fi68++ {
            var v69 float64 = ((v5 - v5) * v5)
            if false {
                if (100 <= -1) {
                    switch ("" + ("hello" + "")) {
                    case "a":
                        switch int(v1) {
                        case -629:
                            var v70 bool = false
                            _ = v70
                        default:
                            delete(v3, 42)
                        }
                    case "foo bar":
                        switch (-(fi68)) {
                        case 32767:
                            fmt.Println(v2, v4)
                        case -802:
                            if false {
                                for fi71 := 0; fi71 < 6; fi71++ {
                                    delete(v3, 2)
                                    switch fi68 {
                                    case 42:
                                        if true {
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v1, v42)
                                            case "foo bar":
                                                if true {
                                                    v72 := v3[65535]
                                                    v73 := GenS7{F0: 890, F1: (fi71 ^ 9223372036854775807)}
                                                    _ = v72
                                                    _ = v73
                                                } else {
                                                    switch (fi71 + (fi68 * fi68)) {
                                                    case 1:
                                                        v74 := GenS6{F0: 115}
                                                        _ = v74
                                                    case -1:
                                                        fmt.Println(fi71, v69, v69)
                                                    case -9223372036854775808:
                                                        switch (fi68 & fi68) {
                                                        case 32768:
                                                            var v75 float32 = (v2 + ((v2 / v2) * 3.14159))
                                                            _ = v75
                                                        case -165:
                                                            v69 = (v5 + v5)
                                                        case -9223372036854775808:
                                                            if (false || (false && false)) {
                                                                var v76 int8 = 127
                                                                switch (fi68 % ((-2147483648) | 1)) {
                                                                case 0:
                                                                    if false {
                                                                        var v77 uint32 = 246
                                                                        if (2147483647 != -2147483648) {
                                                                            switch fi71 {
                                                                            case 128:
                                                                                if (834 != -1) {
                                                                                    var v78 int32 = 0
                                                                                    _ = v78
                                                                                }
                                                                            case 9223372036854775807:
                                                                                if (-2147483648 >= -910) {
                                                                                    v79 := []int{(^(fi68))}
                                                                                    v80 := GenS2{F0: 4294967295, F1: false}
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                } else {
                                                                                    var v81 float64 = v69
                                                                                    fmt.Println(v1)
                                                                                    v76 = (v76 & (v76 & v76))
                                                                                    switch (fi71 - 0) {
                                                                                    case -979:
                                                                                        v82 := GenS3{F0: 0, F1: -589, F2: 665}
                                                                                        _ = v82
                                                                                    case -2147483648:
                                                                                        if true {
                                                                                            fmt.Println(v5, v81)
                                                                                            var v83 bool = false
                                                                                            v42 = (v42 | 151)
                                                                                            _ = v83
                                                                                        }
                                                                                    case -1:
                                                                                        var v84 uint16 = 1
                                                                                        _ = v84
                                                                                    }
                                                                                    _ = v81
                                                                                }
                                                                            }
                                                                            delete(v3, v77)
                                                                            fmt.Println(v5, v69)
                                                                        } else {
                                                                            v85 := GenS8{F0: -875}
                                                                            v2 += -2.25
                                                                            _ = v85
                                                                        }
                                                                        _ = v77
                                                                    } else {
                                                                        switch fi68 {
                                                                        case 2147483647:
                                                                            v86 := v3[2147483647]
                                                                            _ = v86
                                                                        case 32767:
                                                                            v2 = 0.1
                                                                        }
                                                                    }
                                                                }
                                                                _ = v76
                                                            } else {
                                                                var v87 int = fi68
                                                                var v88 int32 = 143
                                                                v89 := GenS5{F0: 42}
                                                                _ = v87
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                        }
                                                    default:
                                                        v90 := GenS6{F0: 541}
                                                        _ = v90
                                                    }
                                                    if false {
                                                        var v91 uint32 = 42
                                                        _ = v91
                                                    } else {
                                                        switch ("foo bar" + ("hello" + "")) {
                                                        case "":
                                                            v4 = v4
                                                        default:
                                                            v92 := v3[uint32(v4)]
                                                            _ = v92
                                                        }
                                                    }
                                                    if (42 == 18446744073709551615) {
                                                        v5 -= -2.25
                                                    }
                                                    var v93 int = fi71
                                                    _ = v93
                                                }
                                            }
                                            if (-9223372036854775808 > 1) {
                                                v94 := []int{fi68, (fi68 / (((-(fi68))) | 1)), -710, 127}
                                                _ = v94
                                            } else {
                                                var v95 float32 = (v2 - v2)
                                                v1 = -143
                                                v95 += 1.5
                                                fmt.Println(fi71, v95)
                                                _ = v95
                                            }
                                        } else {
                                            v42 -= (v42 % ((542) | 1))
                                            v4 -= v4
                                        }
                                    default:
                                        var v96 bool = (-1 > -482)
                                        _ = v96
                                    }
                                    var v97 uint32 = 0
                                    v42 += 2
                                    _ = fi71
                                    _ = v97
                                }
                            } else {
                                v98 := []int{fi68, fi68, fi68, (fi68 * (fi68 + 9223372036854775807))}
                                var v99 bool = (65535 == 0)
                                _ = v98
                                _ = v99
                            }
                        case 459:
                            v3[4294967295] = ("a" + ("hello" + ("hello" + "")))
                        }
                    case "hello":
                        v100 := []int{0, (^(fi68)), (-(fi68))}
                        _ = v100
                    }
                    var v101 int64 = 478
                    switch fi68 {
                    case -259:
                        fmt.Println(v1, v42)
                    }
                    _ = v101
                }
                if false {
                    v102 := []int{fi68, (fi68 | (fi68 << ((65535) & 63))), fi68, fi68}
                    v69 += float64(v2)
                    _ = v102
                }
            } else {
                for fi103 := 0; fi103 < 12; fi103++ {
                    var v104 int64 = 42
                    var v105 uint = (v42 * (v42 & (v42 * (v42 & 255))))
                    v4 = 1
                    var v106 int16 = -32768
                    if false {
                        var v107 uint32 = 2147483648
                        switch "a" {
                        case "":
                            switch "" {
                            case "hello":
                                v108 := []int{}
                                _ = v108
                            case "\n\t":
                                v42 += 18446744073709551615
                                _ = len(v3)
                            case "a":
                                var v109 int16 = v106
                                _ = v109
                            }
                        case "hello":
                            v110 := []int{128, (^(fi68)), -659}
                            _ = v110
                        case "foo bar":
                            v1 = (v106 | -32768)
                        }
                        v4 = (v4 + 111)
                        _ = v107
                    } else {
                        v2 = 1.5
                        if true {
                            fmt.Println(v105, v69)
                        } else {
                            v111 := GenS7{F0: 4294967295, F1: fi103}
                            fmt.Println(v2, fi68, v105)
                            fmt.Println(v106)
                            fmt.Println(fi103, v2, fi103)
                            _ = v111
                        }
                    }
                    _ = fi103
                    _ = v104
                    _ = v105
                    _ = v106
                }
            }
            _ = fi68
            _ = v69
        }
    case "a":
        v3[42] = "hello"
    case "\n\t":
        v2 = (v2 * float32((v5 - (v5 - (1.5 + float64((v2 - float32(v5))))))))
    }
    if false {
        v112 := []int{}
        v5 += 0.0
        v1 = 1000
        _ = v112
    }
    v113 := GenS7{F0: 18446744073709551615, F1: 42}
    for fi114 := 0; fi114 < 2; fi114++ {
        v115 := GenS5{F0: 741}
        _ = fi114
        _ = v115
    }
    fmt.Println(v42, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v67
}
