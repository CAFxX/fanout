package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 uint32; F1 int32; F2 string }
type GenS13 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS14 struct { F0 uint32; F1 int16 }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 uint16; F1 int16; F2 int64 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int64; F1 uint8; F2 string }
type GenS19 struct { F0 uint; F1 uint64 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 uint16; F1 uint32 }
type GenS22 struct { F0 int32; F1 int16; F2 uint8 }
func genFn0(p0 bool) uint32 {
    var v7 uint16 = 599
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint = 1
        v10 := []int{fi8, (-(fi8))}
        _ = fi8
        _ = v10
        _ = v9
    }
    p0 = p0
    v7 = v7
    v7 += v7
    _ = p0
    _ = v7
    return 2147483648
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    v4 := []string{"\n\t", "\n\t", "\n\t", (("hello" + "a") + ("foo bar" + "foo bar"))}
    var v5 uint64 = 0
    var v6 int64 = 974
    v11 := GenS1{F0: v5, F1: 872}
    v12 := GenS2{F0: (v6 << ((2) & 63))}
    var v13 uint64 = 146
    if false {
        for fi14 := 0; fi14 < 20; fi14++ {
            if (!false) {
                if (false || ((true && false) && (!(v5 <= 1)))) {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        if true {
                            v5 += 0
                            if ((-1 < -1) && true) {
                                v2 += v2
                                v16 := GenS1{F0: 1, F1: 4294967295}
                                _ = v16
                            } else {
                                v2 += v2
                                var v17 int16 = 1000
                                v5 = (v5 / ((v13) | 1))
                                _ = v17
                            }
                            var v18 float64 = v2
                            _ = v18
                        }
                        _ = fi15
                    }
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        v19 := GenS3{F0: 119}
                        _ = v19
                    }
                } else {
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v21 := GenS2{F0: -480}
                        v4 = append(v4, ((("hello" + "a") + "") + ""))
                        v13 = (v13 % ((v5) | 1))
                        _ = fi20
                        _ = v21
                    }
                }
                fmt.Println(fi14)
                v6 -= (v6 % ((v6) | 1))
            } else {
                var v22 int8 = -1
                _ = v22
            }
            var v23 int16 = 42
            for fi24 := 0; fi24 < 14; fi24++ {
                if false {
                    fmt.Println(v5, v2)
                    v25 := GenS1{F0: (v13 << ((890) & 63)), F1: 778}
                    var v26 int = (fi14 * (fi24 / ((9223372036854775807) | 1)))
                    _ = v25
                    _ = v26
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v27 := GenS4{F0: 1, F1: 42, F2: 2147483647}
                    _ = v27
                }
                if (v5 >= v5) {
                    _ = len(v4)
                    v28 := GenS5{F0: 0}
                    var v29 uint16 = 32767
                    _ = v28
                    _ = v29
                }
                _ = fi24
            }
            _ = fi14
            _ = v23
        }
        v2 += (v2 / 1e10)
    } else {
        var v30 uint16 = 65535
        var v31 float64 = float64(0.0)
        v30 = (v30 * 65535)
        _ = v30
        _ = v31
    }
    var v32 float32 = (0.0 * 0.0)
    var v33 float32 = 1.5
    for fi34 := 0; fi34 < 9; fi34++ {
        for fi35 := 0; fi35 < 16; fi35++ {
            var v36 uint16 = 65535
            if (!(!(v6 > 0))) {
                if (0.1 <= ((v32 - v33) - v32)) {
                    var v37 float32 = 3.14159
                    fmt.Println(v13)
                    var v38 int16 = -118
                    switch 128 {
                    case 128:
                        switch "a" {
                        case "\n\t":
                            if (42 != 42) {
                                fmt.Println(v37)
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v13 += (v13 % ((v5) | 1))
                                    var v39 uint32 = 955
                                    if ((255 > uint8(v5)) && false) {
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                    } else {
                                        fmt.Println(fi35, v13, fi34)
                                        if ("a" != "\n\t") {
                                            v6 -= 796
                                            if (v32 == v37) {
                                                v39 = v39
                                                fmt.Println(v37, fi35)
                                                v32 -= v33
                                            } else {
                                                var v40 uint32 = (^(v39))
                                                if ((v36 - 2) != (v36 - 925)) {
                                                    v41 := GenS2{F0: v6}
                                                    v42 := GenS2{F0: v6}
                                                    v2 += v2
                                                    _ = v41
                                                    _ = v42
                                                }
                                                _ = v40
                                            }
                                        } else {
                                            if (!(!((v36 * (v36 * v36)) > (v36 - v36)))) {
                                                delete(v1, (fi34 % ((fi34) | 1)))
                                                switch ("hello" + ("\n\t" + "a")) {
                                                case "hello":
                                                    var v43 uint16 = v36
                                                    _ = v43
                                                case "\n\t":
                                                    v36 -= 255
                                                    v44 := GenS6{F0: 29}
                                                    fmt.Println(v6, v2)
                                                    _ = v44
                                                }
                                                if (uint8(v36) < 2) {
                                                    var v45 uint8 = 127
                                                    var v46 uint = 0
                                                    if true {
                                                        if (1 <= 127) {
                                                            v13 = v5
                                                            v47 := GenS3{F0: v38}
                                                            _ = v47
                                                        }
                                                        if (true && false) {
                                                            var v48 uint16 = (v36 % (((^(v36))) | 1))
                                                            _ = v48
                                                        } else {
                                                            v4 = append(v4, "")
                                                            fmt.Println(v38)
                                                        }
                                                        if true {
                                                            var v49 int8 = 100
                                                            v50 := GenS1{F0: (v5 % ((uint64(v39)) | 1)), F1: 2147483648}
                                                            if false {
                                                                v37 = v33
                                                                if false {
                                                                    fmt.Println(v37)
                                                                    switch (fi35 % (((fi35 - 42)) | 1)) {
                                                                    case -523:
                                                                        v46 -= (v46 * (^(v46)))
                                                                    case 9223372036854775807:
                                                                        v32 = (float32(((v2 - (v2 * float64(1e10))) - ((0.1 / -2.25) / v2))) + 0.1)
                                                                    default:
                                                                        var v51 int64 = (^(v6))
                                                                        _ = v51
                                                                    }
                                                                    v49 -= 100
                                                                    if ("hello" == "") {
                                                                        v45 = v45
                                                                        fmt.Println(fi34, v38)
                                                                        v13 = (^(v5))
                                                                    }
                                                                    var v52 float32 = (v32 / (v32 + 0.0))
                                                                    if (!(!false)) {
                                                                        v39 = 65535
                                                                    } else {
                                                                        var v53 bool = false
                                                                        v39 -= v39
                                                                        switch 32768 {
                                                                        case 24:
                                                                            switch "a" {
                                                                            case "":
                                                                                var v54 float32 = (float32(v2) + float32(v2))
                                                                                _ = v54
                                                                            case "foo bar":
                                                                                var v55 int8 = v49
                                                                                _ = v55
                                                                            case "\n\t":
                                                                                v52 = (float32(3.14159) + (3.14159 + v32))
                                                                            }
                                                                        }
                                                                        _ = v53
                                                                    }
                                                                    var v56 int32 = -1
                                                                    switch (fi34 >> ((1) & 63)) {
                                                                    case -93:
                                                                        v57 := GenS7{F0: v36, F1: 2147483647}
                                                                        _ = v57
                                                                    case 739:
                                                                        switch ("hello" + "a") {
                                                                        case "hello":
                                                                            v32 += v37
                                                                        case "foo bar":
                                                                            var v58 float32 = ((0.1 / (v32 - v37)) / v33)
                                                                            _ = v58
                                                                        default:
                                                                            var v59 float32 = (float32(v2) * (1.5 * v52))
                                                                            _ = v59
                                                                        }
                                                                    case -873:
                                                                        v46 = 4294967295
                                                                    }
                                                                    _ = v52
                                                                    _ = v56
                                                                }
                                                                if false {
                                                                    if true {
                                                                        var v60 uint8 = 0
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 float64 = ((float64(v37) / 3.14159) / (float64((v33 * float32(v2))) + float64(1e10)))
                                                                        v62 := GenS1{F0: v5, F1: 749}
                                                                        v63 := GenS8{F0: (v39 - 4294967295), F1: -93, F2: v45}
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                } else {
                                                                    var v64 int32 = 561
                                                                    fmt.Println(v49, v38)
                                                                    v65 := v1[fi34]
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                                switch (fi35 & -444) {
                                                                case -9223372036854775808:
                                                                    v66 := GenS4{F0: 874, F1: 765, F2: (v39 + genFn0(((v46 + v46) < 65535)))}
                                                                    _ = v66
                                                                case 4294967295:
                                                                    var v67 string = "hello"
                                                                    _ = v67
                                                                case 32768:
                                                                    v45 = v45
                                                                default:
                                                                    fmt.Println(v46, v2)
                                                                }
                                                            } else {
                                                                var v68 uint16 = uint16(v45)
                                                                _ = v68
                                                            }
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            switch 9223372036854775807 {
                                                            case 9223372036854775807:
                                                                switch (fi35 / ((fi35) | 1)) {
                                                                case 4294967295:
                                                                    fmt.Println(v37, v13, v46)
                                                                case 9223372036854775807:
                                                                    switch (("hello" + ("foo bar" + "a")) + (("foo bar" + ("hello" + "a")) + "")) {
                                                                    case "foo bar":
                                                                        var v69 bool = (v39 != 73)
                                                                        _ = v69
                                                                    }
                                                                case -798:
                                                                    v70 := GenS9{F0: v46}
                                                                    fmt.Println(v5)
                                                                    var v71 uint64 = uint64(v46)
                                                                    _ = v70
                                                                    _ = v71
                                                                default:
                                                                    v5 = (^(v13))
                                                                }
                                                            case 2147483647:
                                                                v72 := GenS1{F0: v5, F1: v39}
                                                                _ = v72
                                                            case 128:
                                                                switch ("foo bar" + "hello") {
                                                                case "foo bar":
                                                                    var v73 float32 = v37
                                                                    _ = v73
                                                                case "hello":
                                                                    var v74 int32 = int32(v36)
                                                                    _ = v74
                                                                default:
                                                                    delete(v3, (v39 + (^(v39))))
                                                                }
                                                            default:
                                                                v33 = -2.25
                                                            }
                                                            var v75 int16 = v38
                                                            _ = v75
                                                        }
                                                    } else {
                                                        var v76 string = "foo bar"
                                                        if len(v4) > 0 {
                                                            v4[0] = (v76 + v76)
                                                        }
                                                        _ = v76
                                                    }
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    var v77 uint8 = 252
                                                    v78 := GenS5{F0: 65535}
                                                    switch "" {
                                                    case "a":
                                                        if false {
                                                            v79 := GenS5{F0: v36}
                                                            _ = v79
                                                        } else {
                                                            var v80 uint32 = (v39 >> ((915) & 63))
                                                            _ = v80
                                                        }
                                                    case "foo bar":
                                                        if true {
                                                            v5 -= v13
                                                            fmt.Println(v38)
                                                        } else {
                                                            var v81 float32 = (1.5 + v33)
                                                            _ = v81
                                                        }
                                                    }
                                                    _ = v77
                                                    _ = v78
                                                }
                                                switch "" {
                                                case "foo bar":
                                                    if ((true || true) && (2 >= 42)) {
                                                        switch fi35 {
                                                        case 42:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        case -2147483648:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        default:
                                                            _ = len(v4)
                                                        }
                                                        v32 -= 0.0
                                                        var v82 float64 = (v2 * (((((1e10 + 1e10) - v2) - v2) * (v2 / float64(v37))) * 0.1))
                                                        _ = v82
                                                    }
                                                case "hello":
                                                    _ = len(v4)
                                                case "a":
                                                    v13 -= v5
                                                }
                                                v83 := GenS4{F0: 1, F1: -2147483648, F2: (^(v39))}
                                                _ = v83
                                            }
                                            var v84 int32 = -1
                                            fmt.Println(v37, fi34)
                                            switch (("a" + "hello") + "hello") {
                                            case "a":
                                                v33 -= v37
                                            case "":
                                                if false {
                                                    if (4294967295 > (v13 | v13)) {
                                                        fmt.Println(v2, v37)
                                                    }
                                                    if true {
                                                        var v85 int8 = -128
                                                        fmt.Println(v13, v13)
                                                        _ = v85
                                                    } else {
                                                        v37 += v32
                                                        v2 = v2
                                                    }
                                                }
                                            case "\n\t":
                                                switch fi35 {
                                                case -780:
                                                    fmt.Println(v37, fi35, v13)
                                                case 0:
                                                    fmt.Println(v6)
                                                case -1:
                                                    fmt.Println(v36)
                                                }
                                            default:
                                                if (v37 < 0.0) {
                                                    v86 := GenS3{F0: v38}
                                                    if (false && true) {
                                                        switch "a" {
                                                        case "\n\t":
                                                            if true {
                                                                v5 = (^(v13))
                                                            } else {
                                                                _ = len(v3)
                                                                v36 = 255
                                                                switch "foo bar" {
                                                                case "":
                                                                    v84 += v84
                                                                case "a":
                                                                    if true {
                                                                        var v87 uint8 = 2
                                                                        _ = v87
                                                                    }
                                                                default:
                                                                    fmt.Println(v38, v37)
                                                                    _ = len(v3)
                                                                }
                                                            }
                                                        default:
                                                            if false {
                                                                switch 573 {
                                                                case -1:
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        fmt.Println(v36, v32)
                                                                    case "a":
                                                                        var v88 int64 = 278
                                                                        _ = v88
                                                                    case "hello":
                                                                        v89 := GenS4{F0: (v39 & uint32(v36)), F1: (v84 | -2147483648), F2: 0}
                                                                        _ = v89
                                                                    default:
                                                                        fmt.Println(v38, v32, v36)
                                                                    }
                                                                case -260:
                                                                    fmt.Println(v2, v39)
                                                                default:
                                                                    v90 := GenS6{F0: 222}
                                                                    fmt.Println(fi35)
                                                                    _ = v90
                                                                }
                                                                var v91 uint32 = (v39 << ((782) & 63))
                                                                v36 -= 32767
                                                                _ = v91
                                                            }
                                                        }
                                                        var v92 float64 = -2.25
                                                        switch "hello" {
                                                        case "":
                                                            switch (fi35 + 4294967295) {
                                                            case 570:
                                                                if false {
                                                                    if false {
                                                                        var v93 int64 = 42
                                                                        switch int(v39) {
                                                                        case 32767:
                                                                            v36 += 0
                                                                        }
                                                                        _ = v93
                                                                    } else {
                                                                        v39 = v39
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        switch (fi34 / ((fi35) | 1)) {
                                                                        case -667:
                                                                            _ = len(v3)
                                                                        case 0:
                                                                            var v94 uint32 = v39
                                                                            _ = v94
                                                                        default:
                                                                            if false {
                                                                                if (576 == 255) {
                                                                                    var v95 int = fi35
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        v92 = v92
                                                                                    case "":
                                                                                        v96 := GenS10{F0: 735, F1: v84}
                                                                                        _ = v96
                                                                                    default:
                                                                                        v13 = 4294967295
                                                                                    }
                                                                                    if true {
                                                                                        fmt.Println(v6, v6, v13)
                                                                                    } else {
                                                                                        v97 := GenS6{F0: 0}
                                                                                        v13 = 2
                                                                                        _ = v97
                                                                                    }
                                                                                    if false {
                                                                                        if false {
                                                                                            v6 -= (-(v6))
                                                                                        }
                                                                                    }
                                                                                    _ = v95
                                                                                }
                                                                            } else {
                                                                                var v98 int32 = (-(v84))
                                                                                _ = v98
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v4 = append(v4, "\n\t")
                                                                        var v99 int8 = 100
                                                                        _ = len(v4)
                                                                        _ = v99
                                                                    }
                                                                    fmt.Println(v38)
                                                                }
                                                            case -2147483648:
                                                                fmt.Println(v5, v84)
                                                            }
                                                        }
                                                        _ = v92
                                                    }
                                                    _ = v86
                                                } else {
                                                    v100 := GenS8{F0: uint32(v36), F1: -10, F2: 128}
                                                    _ = v100
                                                }
                                                var v101 int16 = 0
                                                v102 := GenS4{F0: uint32(v36), F1: v84, F2: genFn0((18446744073709551615 <= 1))}
                                                v2 -= v2
                                                _ = v101
                                                _ = v102
                                            }
                                            _ = v84
                                        }
                                        var v103 float32 = (0.0 * float32((v2 / v2)))
                                        _ = v103
                                    }
                                    var v104 int16 = int16(v36)
                                    _ = v104
                                    _ = v39
                                }
                                v36 += 2
                            }
                        default:
                            var v105 uint8 = 128
                            _ = v105
                        }
                        v13 = v5
                    case -416:
                        switch fi35 {
                        case 42:
                            if true {
                                if (4294967295 <= 0) {
                                    switch (-(fi35)) {
                                    case 32767:
                                        fmt.Println(v37, v33, v13)
                                    case 0:
                                        if false {
                                            v106 := v1[fi34]
                                            _ = v106
                                        } else {
                                            var v107 int32 = 1
                                            fmt.Println(v6, v36, v5)
                                            _ = v107
                                        }
                                    case 42:
                                        switch (fi35 & (^(fi35))) {
                                        case 9223372036854775807:
                                            fmt.Println(v5, v38)
                                        }
                                    default:
                                        v13 = (v13 | v5)
                                    }
                                    if false {
                                        fmt.Println(fi35)
                                        if (v5 < uint64(v36)) {
                                            v13 -= 18446744073709551615
                                            var v108 bool = false
                                            _ = v108
                                        } else {
                                            if false {
                                                v33 += v32
                                                v109 := GenS9{F0: 276}
                                                fmt.Println(v2)
                                                _ = v109
                                            }
                                        }
                                        if true {
                                            if (v6 <= 1) {
                                                v32 += v33
                                                if len(v4) > 0 {
                                                    v4[0] = ("\n\t" + "a")
                                                }
                                            } else {
                                                v110 := GenS11{F0: ((v6 & (v6 & v6)) <= 42), F1: (v6 % ((v6) | 1))}
                                                _ = v110
                                            }
                                            var v111 int64 = (v6 * v6)
                                            _ = v111
                                        }
                                        if true {
                                            v2 = v2
                                            var v112 string = (("" + ("foo bar" + "foo bar")) + "\n\t")
                                            _ = v112
                                        }
                                    }
                                } else {
                                    _ = len(v3)
                                    v5 = 42
                                    v38 -= v38
                                }
                                v33 = 0.0
                                if (((genFn0((!true)) == 65535) || ((((v2 - (v2 / 3.14159)) / v2) + (((v2 + 0.0) + 1.5) + (v2 - ((((v2 * (v2 * v2)) / v2) * 3.14159) - float64((v37 + v32)))))) <= 1.5)) && false) {
                                    if ((v6 > (v6 % (((v6 / ((v6) | 1))) | 1))) && true) {
                                        var v113 uint16 = (v36 % ((v36) | 1))
                                        fmt.Println(v32)
                                        switch 127 {
                                        case -9223372036854775808:
                                            v114 := GenS12{F0: 0, F1: -1, F2: "a"}
                                            _ = v114
                                        default:
                                            if true {
                                                if (1 <= 1) {
                                                    if false {
                                                        fmt.Println(v113, v13, v32)
                                                        v5 = (^(v5))
                                                    } else {
                                                        fmt.Println(fi34)
                                                        var v115 float32 = 0.1
                                                        v33 -= v33
                                                        _ = v115
                                                    }
                                                    if false {
                                                        var v116 int = (fi35 & fi35)
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v4 = append(v4, "hello")
                                                        default:
                                                            var v117 uint = 18446744073709551615
                                                            _ = v117
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v116
                                                    } else {
                                                        if ((!false) || true) {
                                                            v118 := GenS11{F0: (!true), F1: (v6 << ((2) & 63))}
                                                            _ = v118
                                                        } else {
                                                            switch 42 {
                                                            case 32767:
                                                                v32 = ((1e10 * v37) * v33)
                                                            case 162:
                                                                v113 = (^(v36))
                                                            }
                                                        }
                                                        fmt.Println(v38)
                                                        if false {
                                                            if len(v4) > 0 {
                                                                v4[0] = "foo bar"
                                                            }
                                                            switch ("" + ("foo bar" + "")) {
                                                            case "\n\t":
                                                                delete(v1, 128)
                                                            default:
                                                                v36 -= v36
                                                            }
                                                            fmt.Println(v113, v2, v37)
                                                            _ = len(v3)
                                                        }
                                                        var v119 int32 = int32(v5)
                                                        fmt.Println(fi34, fi34)
                                                        _ = v119
                                                    }
                                                } else {
                                                    if true {
                                                        fmt.Println(v33, v5, v37)
                                                    }
                                                    switch (fi35 + fi35) {
                                                    case -1:
                                                        var v120 uint16 = v113
                                                        _ = v120
                                                    case 32767:
                                                        if false {
                                                            if false {
                                                                v2 = (0.0 - float64(v37))
                                                            } else {
                                                                v121 := GenS10{F0: 601, F1: -904}
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                _ = v121
                                                            }
                                                            if true {
                                                                switch ("foo bar" + "a") {
                                                                case "foo bar":
                                                                    var v122 uint = 65535
                                                                    _ = v122
                                                                case "a":
                                                                    var v123 bool = (true && false)
                                                                    _ = v123
                                                                }
                                                                switch -374 {
                                                                case 145:
                                                                    fmt.Println(fi34, v33, v32)
                                                                case 127:
                                                                    v124 := v1[(fi34 & fi35)]
                                                                    _ = v124
                                                                case 0:
                                                                    if true {
                                                                        if true {
                                                                            var v125 uint8 = 182
                                                                            fmt.Println(v33, v5, fi34)
                                                                            switch 1 {
                                                                            case -1:
                                                                                if (42 == (v38 | (^(v38)))) {
                                                                                    var v126 uint16 = (v113 ^ v36)
                                                                                    _ = v126
                                                                                } else {
                                                                                    v37 = v32
                                                                                }
                                                                                v127 := GenS6{F0: v125}
                                                                                var v128 uint32 = 2147483648
                                                                                _ = v127
                                                                                _ = v128
                                                                            case 1:
                                                                                v129 := GenS13{F0: 65535, F1: 2147483647, F2: v125}
                                                                                _ = v129
                                                                            default:
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                var v130 uint64 = (^(v13))
                                                                                v131 := v1[(^(fi35))]
                                                                                if (!(v113 > v36)) {
                                                                                    if true {
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                    } else {
                                                                                        var v132 float64 = 1e10
                                                                                        _ = v132
                                                                                    }
                                                                                    if true {
                                                                                        if v131 {
                                                                                            fmt.Println(v6)
                                                                                            if v131 {
                                                                                                v113 = (v113 ^ v113)
                                                                                                switch -201 {
                                                                                                case 128:
                                                                                                    v113 = v36
                                                                                                default:
                                                                                                    var v133 uint16 = v36
                                                                                                    _ = v133
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v134 := GenS12{F0: 432, F1: 2147483647, F2: ""}
                                                                                            if ((v113 * (v36 % ((v36) | 1))) == v36) {
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = ("foo bar" + "foo bar")
                                                                                                }
                                                                                            } else {
                                                                                                var v135 int32 = 300
                                                                                                switch (fi35 - 32767) {
                                                                                                case 0:
                                                                                                    if v131 {
                                                                                                        var v136 int32 = v135
                                                                                                        if false {
                                                                                                            v36 += v113
                                                                                                            var v137 uint64 = (v5 | 421)
                                                                                                            _ = v137
                                                                                                        } else {
                                                                                                            v37 += v33
                                                                                                            v38 = v38
                                                                                                            var v138 int64 = v6
                                                                                                            v130 = v13
                                                                                                            _ = v138
                                                                                                        }
                                                                                                        _ = v136
                                                                                                    }
                                                                                                case 822:
                                                                                                    var v139 uint32 = 42
                                                                                                    _ = v139
                                                                                                default:
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = ""
                                                                                                    }
                                                                                                }
                                                                                                v131 = false
                                                                                                v2 -= v2
                                                                                                v140 := GenS10{F0: 65535, F1: 42}
                                                                                                _ = v135
                                                                                                _ = v140
                                                                                            }
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            _ = v134
                                                                                        }
                                                                                        var v141 float32 = v37
                                                                                        var v142 int64 = (v6 - (-(v6)))
                                                                                        _ = v141
                                                                                        _ = v142
                                                                                    } else {
                                                                                        delete(v3, 42)
                                                                                        if false {
                                                                                            v131 = ((v38 % ((int16(v36)) | 1)) != v38)
                                                                                        } else {
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = "\n\t"
                                                                                            }
                                                                                            var v143 float64 = (v2 - 3.14159)
                                                                                            _ = v143
                                                                                        }
                                                                                    }
                                                                                    var v144 int8 = int8(v6)
                                                                                    if (!(int64(v38) == -1)) {
                                                                                        _ = len(v4)
                                                                                    } else {
                                                                                        v32 -= (v32 / 1e10)
                                                                                        var v145 int16 = (v38 << ((311) & 63))
                                                                                        _ = v145
                                                                                    }
                                                                                    _ = v144
                                                                                }
                                                                                _ = v130
                                                                                _ = v131
                                                                            }
                                                                            var v146 int8 = -128
                                                                            _ = v125
                                                                            _ = v146
                                                                        }
                                                                    }
                                                                }
                                                                v147 := GenS6{F0: 0}
                                                                v38 = v38
                                                                _ = v147
                                                            } else {
                                                                var v148 int8 = 1
                                                                _ = v148
                                                            }
                                                            v149 := GenS14{F0: genFn0((!(fi35 <= (fi34 >> ((65535) & 63))))), F1: v38}
                                                            _ = v149
                                                        }
                                                        if true {
                                                            fmt.Println(fi34)
                                                        } else {
                                                            v150 := GenS1{F0: 4294967295, F1: 2147483648}
                                                            _ = v150
                                                        }
                                                        var v151 float32 = v32
                                                        v151 += v32
                                                        switch "a" {
                                                        case "\n\t":
                                                            if (false || (((!(!(false || (!true)))) && ((!(v5 <= v5)) || ((^(v38)) > -32768))) || true)) {
                                                                v3[4294967295] = 2
                                                                fmt.Println(v2)
                                                                switch fi34 {
                                                                case 9223372036854775807:
                                                                    var v152 int64 = (-(v6))
                                                                    _ = v152
                                                                case 32767:
                                                                    v37 -= ((float32(v2) - v151) + float32(float64(v32)))
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v153 := GenS15{F0: 2, F1: int32(v38)}
                                                                        _ = v153
                                                                    case "foo bar":
                                                                        var v154 uint32 = 1
                                                                        _ = v154
                                                                    default:
                                                                        switch (fi34 % (((fi34 & (^(fi35)))) | 1)) {
                                                                        case -9223372036854775808:
                                                                            var v155 float32 = (3.14159 / float32(1.5))
                                                                            var v156 uint = 2
                                                                            var v157 uint32 = 2
                                                                            _ = v155
                                                                            _ = v156
                                                                            _ = v157
                                                                        default:
                                                                            fmt.Println(fi35, v2)
                                                                        }
                                                                    }
                                                                    delete(v1, 85)
                                                                default:
                                                                    _ = len(v4)
                                                                }
                                                                var v158 int64 = v6
                                                                _ = v158
                                                            }
                                                        case "foo bar":
                                                            switch "hello" {
                                                            case "":
                                                                v13 -= v13
                                                            }
                                                        default:
                                                            var v159 int32 = 2147483647
                                                            _ = v159
                                                        }
                                                        _ = v151
                                                    default:
                                                        delete(v3, 2147483648)
                                                        v160 := GenS16{F0: 255, F1: (v38 / ((v38) | 1)), F2: v6}
                                                        var v161 string = "\n\t"
                                                        _ = v160
                                                        _ = v161
                                                    }
                                                }
                                                v38 += (^(v38))
                                            }
                                            if ((v113 & v113) != (^(v113))) {
                                                v2 -= v2
                                            } else {
                                                v162 := GenS17{F0: (v113 >> ((0) & 63))}
                                                switch "" {
                                                case "foo bar":
                                                    fmt.Println(fi34)
                                                case "a":
                                                    v13 += 822
                                                case "hello":
                                                    v5 = (v13 / ((18446744073709551615) | 1))
                                                }
                                                v5 += 2
                                                if true {
                                                    _ = len(v1)
                                                    v163 := v1[int(v113)]
                                                    _ = v163
                                                } else {
                                                    v164 := GenS18{F0: (v6 / ((v6) | 1)), F1: 2, F2: "a"}
                                                    if false {
                                                        var v165 uint = 117
                                                        v36 += (v36 << (((v165 + v165)) & 63))
                                                        if true {
                                                            var v166 uint64 = v5
                                                            v167 := GenS9{F0: v165}
                                                            fmt.Println(v113, fi35)
                                                            var v168 uint = 65535
                                                            _ = v166
                                                            _ = v167
                                                            _ = v168
                                                        } else {
                                                            v165 = (v165 << ((v165) & 63))
                                                        }
                                                        if true {
                                                            v169 := GenS19{F0: (v165 + v165), F1: v5}
                                                            if true {
                                                                v5 = (v5 * 2)
                                                                switch (-(fi35)) {
                                                                case 1:
                                                                    fmt.Println(fi35, v6, v37)
                                                                case 32767:
                                                                    var v170 uint64 = (v5 % ((v5) | 1))
                                                                    var v171 int32 = 1
                                                                    v37 = v37
                                                                    _ = v170
                                                                    _ = v171
                                                                case -9223372036854775808:
                                                                    v38 += v38
                                                                }
                                                                var v172 int64 = v6
                                                                _ = v172
                                                            } else {
                                                                var v173 int = 42
                                                                _ = v173
                                                            }
                                                            _ = v169
                                                        }
                                                        if (((!true) || ((v36 & v36) < v113)) || false) {
                                                            v32 -= (3.14159 + (v32 * (1.5 * float32((0.1 + v2)))))
                                                        }
                                                        if (uint8(v165) < 208) {
                                                            v4 = append(v4, "")
                                                            fmt.Println(v37, v33, v6)
                                                        }
                                                        var v174 float32 = (-2.25 * v33)
                                                        v5 = uint64(v113)
                                                        _ = v165
                                                        _ = v174
                                                    }
                                                    _ = v164
                                                }
                                                _ = v162
                                            }
                                            var v175 int = (fi34 + (fi35 % ((4294967295) | 1)))
                                            switch (fi34 - fi34) {
                                            case 406:
                                                v2 -= v2
                                            case -9223372036854775808:
                                                v32 = 0.1
                                            default:
                                                fmt.Println(v37, v33, v5)
                                            }
                                            if (v2 == ((v2 + float64(((v37 * float32(-2.25)) + ((v33 + ((((float32(v2) - (1e10 + 1.5)) / (v32 - (v32 * 1.5))) * 1e10) + v37)) - (0.0 + v32))))) / 1e10)) {
                                                switch (("foo bar" + ("\n\t" + "")) + "\n\t") {
                                                case "\n\t":
                                                    v5 += v5
                                                case "foo bar":
                                                    fmt.Println(v6, v113)
                                                default:
                                                    switch fi35 {
                                                    case -9223372036854775808:
                                                        var v176 int64 = (v6 >> ((575) & 63))
                                                        _ = v176
                                                    default:
                                                        if false {
                                                            v2 += (v2 - (((v2 + float64(3.14159)) / v2) + ((3.14159 - 0.0) + v2)))
                                                            switch 32768 {
                                                            case 32767:
                                                                v3[2147483648] = 42
                                                                v32 += (v37 + float32((v2 + v2)))
                                                                switch 127 {
                                                                case 4294967295:
                                                                    fmt.Println(v13)
                                                                case 9223372036854775807:
                                                                    v177 := v3[42]
                                                                    _ = v177
                                                                case 127:
                                                                    v178 := GenS20{F0: (27 != v38)}
                                                                    _ = v178
                                                                default:
                                                                    v36 = v113
                                                                }
                                                            }
                                                            fmt.Println(v6, v32)
                                                        } else {
                                                            v38 = v38
                                                        }
                                                    }
                                                }
                                                var v179 int32 = 42
                                                var v180 int8 = 36
                                                _ = v179
                                                _ = v180
                                            }
                                            _ = v175
                                        }
                                        _ = v113
                                    } else {
                                        v37 += (-2.25 + (3.14159 * v33))
                                    }
                                    if false {
                                        v33 = v32
                                        v181 := GenS11{F0: true, F1: -1}
                                        var v182 int8 = -128
                                        _ = v181
                                        _ = v182
                                    } else {
                                        switch (^(fi34)) {
                                        case 534:
                                            if true {
                                                switch "a" {
                                                case "a":
                                                    var v183 float64 = 3.14159
                                                    _ = v183
                                                case "":
                                                    v38 = (-(v38))
                                                case "\n\t":
                                                    if len(v4) > 0 {
                                                        v4[0] = ("\n\t" + ("a" + ""))
                                                    }
                                                }
                                            } else {
                                                var v184 float64 = (1e10 + v2)
                                                _ = v184
                                            }
                                        default:
                                            var v185 uint64 = 4294967295
                                            _ = v185
                                        }
                                        var v186 uint = 2
                                        var v187 int16 = (v38 * v38)
                                        v3[925] = 2147483648
                                        _ = v186
                                        _ = v187
                                    }
                                    var v188 int8 = -128
                                    _ = v188
                                }
                                if false {
                                    fmt.Println(fi34)
                                } else {
                                    v37 += v37
                                    v38 -= v38
                                    var v189 uint16 = 523
                                    _ = v189
                                }
                            }
                        case 0:
                            v38 = (v38 * v38)
                        default:
                            v33 = v32
                        }
                    }
                    v2 += v2
                    var v190 uint = 928
                    _ = v190
                    _ = v37
                    _ = v38
                }
                if true {
                    switch (fi34 << ((4294967295) & 63)) {
                    case 678:
                        var v191 uint8 = 42
                        _ = v191
                    default:
                        v192 := GenS16{F0: 311, F1: 0, F2: v6}
                        _ = v192
                    }
                    var v193 uint32 = 65535
                    var v194 string = ("foo bar" + "foo bar")
                    _ = v193
                    _ = v194
                }
                var v195 uint16 = (v36 + v36)
                _ = v195
            }
            if false {
                _ = len(v4)
                v196 := GenS8{F0: genFn0(false), F1: -89, F2: 42}
                if (42 >= 22) {
                    switch int(v6) {
                    case 42:
                        switch fi34 {
                        case -2147483648:
                            fmt.Println(v6, v13, fi35)
                        case 128:
                            if (true || true) {
                                v197 := v3[2147483648]
                                v33 -= v32
                                if (42 == 100) {
                                    fmt.Println(v13, v33, fi35)
                                    v198 := GenS21{F0: 42, F1: genFn0(true)}
                                    switch 4294967295 {
                                    case 2147483647:
                                        v199 := GenS22{F0: 0, F1: -348, F2: 255}
                                        _ = v199
                                    case 128:
                                        var v200 float64 = (v2 - float64((1.5 * v33)))
                                        _ = v200
                                    case 42:
                                        _ = len(v1)
                                        _ = len(v3)
                                    default:
                                        var v201 uint8 = 217
                                        _ = v201
                                    }
                                    _ = v198
                                }
                                _ = v197
                            }
                        case 434:
                            var v202 string = ("\n\t" + "hello")
                            _ = v202
                        }
                    }
                } else {
                    v203 := GenS14{F0: 233, F1: 382}
                    v6 -= -9223372036854775808
                    _ = v203
                }
                v2 = ((v2 - v2) + float64(float32(float64(1.5))))
                v6 = v6
                _ = v196
            }
            _ = fi35
            _ = v36
        }
        for fi204 := 0; fi204 < 8; fi204++ {
            if true {
                var v205 int8 = -45
                _ = v205
            } else {
                v4 = append(v4, "hello")
                v1[fi34] = false
            }
            _ = fi204
        }
        for fi206 := 0; fi206 < 11; fi206++ {
            switch (fi34 << ((81) & 63)) {
            case -1:
                fmt.Println(v13, v5)
            case -450:
                v32 += v33
            }
            var v207 uint16 = 65535
            _ = fi206
            _ = v207
        }
        _ = fi34
    }
    if true {
        fmt.Println(v33, v13, v13)
    }
    if true {
        v5 -= 42
        if true {
            v4 = append(v4, "a")
            if (127 >= 122) {
                fmt.Println(v32, v32)
                v6 += 1
            } else {
                fmt.Println(v32, v32)
                if (!(v6 != (v6 & (v6 ^ v6)))) {
                    v33 = v32
                } else {
                    switch 42 {
                    case -1000:
                        switch "\n\t" {
                        case "foo bar":
                            if (255 >= 2) {
                                v6 = v6
                                if false {
                                    v208 := GenS2{F0: (v6 << ((0) & 63))}
                                    v209 := GenS9{F0: 524}
                                    for fi210 := 0; fi210 < 6; fi210++ {
                                        for fi211 := 0; fi211 < 17; fi211++ {
                                            var v212 int32 = 1
                                            _ = len(v4)
                                            v5 = v13
                                            _ = fi211
                                            _ = v212
                                        }
                                        _ = fi210
                                    }
                                    _ = v208
                                    _ = v209
                                }
                            }
                        case "a":
                            fmt.Println(v2)
                        default:
                            v13 = (v13 - (v13 | 42))
                        }
                    default:
                        var v213 int8 = -128
                        _ = v213
                    }
                }
            }
        }
    }
    fmt.Println(v6, v13)
    var v214 int = 0
    v6 += v6
    var v215 int = v214
    v2 += (v2 / (float64(((v32 * float32((v2 + float64(v32)))) + ((v33 / -2.25) - v32))) * 3.14159))
    v4 = append(v4, "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v214:", v214)
    fmt.Println("v215:", v215)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v214
    _ = v215
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
