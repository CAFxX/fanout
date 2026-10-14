package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int; F2 bool }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int8; F1 uint8; F2 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 int16 }
type GenS10 struct { F0 int16; F1 int32; F2 int }
type GenS11 struct { F0 uint32; F1 uint32 }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 int32; F1 uint16; F2 int }
type GenS14 struct { F0 uint16; F1 string; F2 int16 }
type GenS15 struct { F0 uint; F1 uint32 }
type GenS16 struct { F0 uint8; F1 uint8 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := []int{-580, 32768, 2147483647, -1}
    var v2 uint = 18446744073709551615
    v3 := []uint32{}
    var v4 uint8 = 1
    for fi5 := 0; fi5 < 13; fi5++ {
        if false {
            v6 := GenS1{F0: (fi5 % (((fi5 ^ fi5)) | 1)), F1: 527}
            var v7 int64 = 9223372036854775807
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 15; fi8++ {
                if true {
                    v4 = 53
                }
                switch fi8 {
                case -942:
                    var v9 uint = (v2 >> ((0) & 63))
                    _ = v9
                case -985:
                    if ((v2 ^ (v2 % (((v2 / (((v2 & (v2 + 1))) | 1))) | 1))) != (^(v2))) {
                        v10 := GenS1{F0: (fi5 - fi8), F1: 1}
                        v4 = 1
                        fmt.Println(v2, fi8, fi8)
                        _ = v10
                    }
                case 127:
                    switch "\n\t" {
                    case "\n\t":
                        v4 += v4
                    case "foo bar":
                        if (1000 == 1) {
                            switch "\n\t" {
                            case "foo bar":
                                switch ("\n\t" + "a") {
                                case "hello":
                                    _ = len(v1)
                                case "foo bar":
                                    var v11 uint = v2
                                    _ = v11
                                case "":
                                    v12 := GenS2{F0: (-32768 > -782), F1: 0}
                                    _ = v12
                                default:
                                    if ((42 < 513) || false) {
                                        if true {
                                            v4 += (v4 ^ v4)
                                            var v13 int32 = -2147483648
                                            _ = v13
                                        } else {
                                            switch 485 {
                                            case 77:
                                                fmt.Println(v4)
                                            case 1:
                                                v14 := GenS2{F0: true, F1: 1}
                                                _ = v14
                                            case -716:
                                                fmt.Println(v4)
                                            default:
                                                if (!(v2 >= v2)) {
                                                    v4 -= 227
                                                    if false {
                                                        v15 := GenS1{F0: fi8, F1: 87}
                                                        if len(v1) > 0 {
                                                            v1[0] = (fi5 % ((fi5) | 1))
                                                        }
                                                        _ = v15
                                                    } else {
                                                        if true {
                                                            var v16 int64 = 1
                                                            if len(v3) > 0 {
                                                                v3[0] = 2
                                                            }
                                                            var v17 uint = v2
                                                            _ = v16
                                                            _ = v17
                                                        }
                                                        var v18 uint32 = 2147483647
                                                        if len(v3) > 0 {
                                                            v3[0] = 615
                                                        }
                                                        _ = v18
                                                    }
                                                    var v19 bool = true
                                                    _ = v19
                                                }
                                            }
                                            if true {
                                                if false {
                                                    var v20 uint64 = 18446744073709551615
                                                    _ = v20
                                                }
                                            } else {
                                                var v21 int = 32768
                                                _ = v21
                                            }
                                            if false {
                                                if false {
                                                    v2 = (v2 - 2)
                                                }
                                                if (!false) {
                                                    if true {
                                                        switch "" {
                                                        case "a":
                                                            fmt.Println(v2, v2)
                                                        case "\n\t":
                                                            var v22 string = ("foo bar" + "\n\t")
                                                            _ = v22
                                                        case "foo bar":
                                                            v4 = v4
                                                            v1 = append(v1, (^(fi8)))
                                                        }
                                                        _ = len(v1)
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = uint32(v2)
                                                        }
                                                    }
                                                    var v23 int32 = 986
                                                    _ = v23
                                                } else {
                                                    if false {
                                                        var v24 int8 = -128
                                                        switch (fi8 | -1) {
                                                        case 9:
                                                            v4 = (v4 ^ 187)
                                                        default:
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v2 = 122
                                                            case "\n\t":
                                                                if true {
                                                                    var v25 uint = 612
                                                                    v3 = append(v3, 2147483647)
                                                                    _ = v25
                                                                }
                                                            case "a":
                                                                fmt.Println(fi8, v24, fi8)
                                                            }
                                                        }
                                                        _ = v24
                                                    } else {
                                                        v4 -= v4
                                                        if ((false && false) && false) {
                                                            switch (fi8 % (((fi5 >> ((1) & 63))) | 1)) {
                                                            case 832:
                                                                var v26 uint = 881
                                                                _ = v26
                                                            case 1:
                                                                v27 := GenS3{F0: 0}
                                                                _ = v27
                                                            case -85:
                                                                var v28 uint = 255
                                                                _ = v28
                                                            default:
                                                                var v29 uint32 = 4294967295
                                                                _ = v29
                                                            }
                                                            if true {
                                                                switch (-(fi5)) {
                                                                case -1:
                                                                    fmt.Println(fi8, fi8, v4)
                                                                case 128:
                                                                    var v30 int = (-(fi5))
                                                                    _ = v30
                                                                case 1:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                default:
                                                                    switch (fi8 >> ((v2) & 63)) {
                                                                    case 4294967295:
                                                                        v31 := GenS4{F0: 42, F1: (fi8 - fi5), F2: (!false)}
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        switch (fi5 << ((656) & 63)) {
                                                                        case 127:
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                        _ = v31
                                                                    default:
                                                                        var v32 int = (fi5 % ((fi5) | 1))
                                                                        _ = v32
                                                                    }
                                                                    var v33 int8 = 1
                                                                    v34 := GenS2{F0: false, F1: 32767}
                                                                    _ = v33
                                                                    _ = v34
                                                                }
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v2 = (v2 + (v2 >> ((369) & 63)))
                                                                    v2 += (v2 - 2)
                                                                    var v35 float64 = 1.5
                                                                    _ = v35
                                                                case "":
                                                                    var v36 int16 = -2
                                                                    _ = v36
                                                                default:
                                                                    if true {
                                                                        if true {
                                                                            v4 += v4
                                                                            v37 := GenS3{F0: -2147483648}
                                                                            v2 = (v2 - 435)
                                                                            _ = v37
                                                                        }
                                                                        var v38 int16 = 123
                                                                        v2 = v2
                                                                        v39 := GenS1{F0: -1, F1: 4294967295}
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 779
                                                                        }
                                                                        var v40 int = (-(fi5))
                                                                        _ = v38
                                                                        _ = v39
                                                                        _ = v40
                                                                    }
                                                                }
                                                            } else {
                                                                switch (^(fi8)) {
                                                                case 32768:
                                                                    if false {
                                                                        v2 = v2
                                                                    }
                                                                default:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                                var v41 uint16 = 0
                                                                if (((2147483648 != uint32(v41)) && (!(!true))) && true) {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v42 uint64 = 1
                                                                    switch (fi5 / ((fi5) | 1)) {
                                                                    case 128:
                                                                        v43 := GenS2{F0: true, F1: 255}
                                                                        if false {
                                                                            v42 += (v42 - 2)
                                                                            v44 := GenS5{F0: v4, F1: fi5}
                                                                            _ = v44
                                                                        } else {
                                                                            if false {
                                                                                var v45 float64 = (1.5 - 0.0)
                                                                                var v46 float64 = 1e10
                                                                                fmt.Println(v41, v45)
                                                                                var v47 float64 = (v45 + -2.25)
                                                                                v46 = 0.1
                                                                                _ = v45
                                                                                _ = v46
                                                                                _ = v47
                                                                            } else {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v48 := GenS5{F0: v4, F1: (fi5 - fi8)}
                                                                                switch fi5 {
                                                                                case 0:
                                                                                    v1 = append(v1, 2147483647)
                                                                                case -764:
                                                                                    if false {
                                                                                        if false {
                                                                                            if false {
                                                                                                v49 := GenS1{F0: (fi5 - fi8), F1: v42}
                                                                                                v41 += (v41 % ((v41) | 1))
                                                                                                if (688 == (v41 << ((928) & 63))) {
                                                                                                    if ((("hello" + "\n\t") > ("hello" + "a")) && true) {
                                                                                                        fmt.Println(v42, v41, v41)
                                                                                                        v50 := GenS6{F0: 127, F1: 255, F2: true}
                                                                                                        if true {
                                                                                                            var v51 uint8 = v4
                                                                                                            var v52 uint32 = 847
                                                                                                            if true {
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = 42
                                                                                                                }
                                                                                                                if false {
                                                                                                                    var v53 int = fi5
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = v52
                                                                                                                    }
                                                                                                                    var v54 uint = (v2 * 497)
                                                                                                                    v3 = append(v3, 42)
                                                                                                                    v41 = (v41 % (((v41 + (v41 & v41))) | 1))
                                                                                                                    _ = v53
                                                                                                                    _ = v54
                                                                                                                }
                                                                                                                var v55 int = (fi5 + (^(fi8)))
                                                                                                                v55 += (v55 | -775)
                                                                                                                _ = v55
                                                                                                            } else {
                                                                                                                v56 := GenS7{F0: 32767}
                                                                                                                var v57 float32 = (0.1 / 1e10)
                                                                                                                v58 := GenS8{F0: (42 == v42)}
                                                                                                                if false {
                                                                                                                    switch -224 {
                                                                                                                    case 1:
                                                                                                                        v59 := GenS9{F0: fi5, F1: 0}
                                                                                                                        _ = v59
                                                                                                                    case 42:
                                                                                                                        if (((v2 << ((4294967295) & 63)) < (v2 - uint(v52))) && true) {
                                                                                                                            v41 += 0
                                                                                                                            switch (fi8 ^ (-(fi8))) {
                                                                                                                            case 128:
                                                                                                                                v60 := GenS10{F0: 674, F1: -789, F2: 128}
                                                                                                                                var v61 float64 = 3.14159
                                                                                                                                _ = v60
                                                                                                                                _ = v61
                                                                                                                            case 1:
                                                                                                                                if false {
                                                                                                                                    if true {
                                                                                                                                        fmt.Println(v4, fi5)
                                                                                                                                    }
                                                                                                                                    fmt.Println(fi5, v41)
                                                                                                                                    v2 = uint(v52)
                                                                                                                                }
                                                                                                                            case 292:
                                                                                                                                v41 -= 104
                                                                                                                            default:
                                                                                                                                v51 += (v51 % (((v4 ^ v51)) | 1))
                                                                                                                            }
                                                                                                                            if len(v1) >= 2 {
                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                v42 = v42
                                                                                                                _ = v56
                                                                                                                _ = v57
                                                                                                                _ = v58
                                                                                                            }
                                                                                                            _ = v51
                                                                                                            _ = v52
                                                                                                        }
                                                                                                        _ = v50
                                                                                                    }
                                                                                                }
                                                                                                _ = v49
                                                                                            } else {
                                                                                                v1 = append(v1, (fi5 >> (((v2 - v2)) & 63)))
                                                                                            }
                                                                                        }
                                                                                        var v62 uint32 = 848
                                                                                        _ = v62
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v4, v42, fi5)
                                                                                }
                                                                                v41 = v41
                                                                                _ = v48
                                                                            }
                                                                            if false {
                                                                                if true {
                                                                                    fmt.Println(fi8)
                                                                                    var v63 uint64 = v42
                                                                                    v64 := GenS11{F0: 0, F1: 2}
                                                                                    _ = v63
                                                                                    _ = v64
                                                                                } else {
                                                                                    if true {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 4294967295
                                                                                        }
                                                                                        v4 = (^(v4))
                                                                                    }
                                                                                    v4 -= 255
                                                                                    v65 := GenS12{F0: int(v4), F1: 127}
                                                                                    _ = v65
                                                                                }
                                                                            }
                                                                        }
                                                                        var v66 int = fi8
                                                                        _ = v43
                                                                        _ = v66
                                                                    default:
                                                                        switch "a" {
                                                                        case "":
                                                                            var v67 string = "foo bar"
                                                                            _ = v67
                                                                        default:
                                                                            v68 := GenS5{F0: (v4 % ((128) | 1)), F1: (-(fi8))}
                                                                            _ = v68
                                                                        }
                                                                    }
                                                                    _ = v42
                                                                }
                                                                _ = v41
                                                            }
                                                            switch (("" + "foo bar") + "\n\t") {
                                                            case "foo bar":
                                                                if len(v3) > 0 {
                                                                    v3[0] = 30
                                                                }
                                                            }
                                                        }
                                                        if false {
                                                            var v69 uint64 = 42
                                                            v69 = 0
                                                            switch fi8 {
                                                            case 128:
                                                                v70 := GenS13{F0: 2147483647, F1: 2, F2: fi5}
                                                                _ = v70
                                                            case -372:
                                                                if false {
                                                                    var v71 uint8 = (v4 | v4)
                                                                    var v72 uint64 = (v69 % (((v69 ^ (v69 % ((1) | 1)))) | 1))
                                                                    v73 := GenS5{F0: (v71 - (v71 >> ((v2) & 63))), F1: (fi8 & 32767)}
                                                                    _ = v71
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                            case 42:
                                                                if true {
                                                                    if (9223372036854775807 == 9223372036854775807) {
                                                                        if false {
                                                                            fmt.Println(fi8)
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                v4 = v4
                                                                            case "hello":
                                                                                v69 = (v69 & 1)
                                                                            case "a":
                                                                                v2 = (v2 >> (((v2 % (((v2 & (v2 % ((485) | 1)))) | 1))) & 63))
                                                                            default:
                                                                                _ = len(v1)
                                                                            }
                                                                            switch ("hello" + "a") {
                                                                            case "a":
                                                                                var v74 int16 = 1
                                                                                _ = v74
                                                                            case "foo bar":
                                                                                v75 := GenS14{F0: 1, F1: "\n\t", F2: -32768}
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                        var v76 int32 = -825
                                                                        _ = v76
                                                                    }
                                                                    switch ("hello" + "a") {
                                                                    case "":
                                                                        fmt.Println(v2, v4)
                                                                    case "foo bar":
                                                                        v77 := GenS6{F0: -90, F1: (v4 / (((^(v4))) | 1)), F2: false}
                                                                        var v78 int32 = 42
                                                                        _ = len(v1)
                                                                        var v79 bool = (0 < v78)
                                                                        _ = v77
                                                                        _ = v78
                                                                        _ = v79
                                                                    case "a":
                                                                        v80 := GenS3{F0: 0}
                                                                        v81 := GenS12{F0: fi8, F1: 127}
                                                                        if true {
                                                                            v2 = v2
                                                                            v2 = 1
                                                                        } else {
                                                                            _ = len(v3)
                                                                            fmt.Println(v4)
                                                                            switch (fi5 >> ((0) & 63)) {
                                                                            case -935:
                                                                                var v82 uint16 = 30
                                                                                _ = v82
                                                                            case -1:
                                                                                var v83 uint = (v2 | uint(v69))
                                                                                _ = v83
                                                                            case 128:
                                                                                v4 += v4
                                                                            default:
                                                                                v4 = 2
                                                                            }
                                                                        }
                                                                        _ = v80
                                                                        _ = v81
                                                                    default:
                                                                        fmt.Println(v2, v4)
                                                                    }
                                                                    var v84 int64 = -9223372036854775808
                                                                    var v85 uint8 = (v4 / ((v4) | 1))
                                                                    _ = v84
                                                                    _ = v85
                                                                } else {
                                                                    _ = len(v1)
                                                                }
                                                                v69 -= (^(v69))
                                                            default:
                                                                var v86 uint = (v2 | (v2 & v2))
                                                                _ = v86
                                                            }
                                                            _ = v69
                                                        }
                                                    }
                                                }
                                            } else {
                                                v87 := GenS15{F0: (v2 % (((^(v2))) | 1)), F1: 631}
                                                var v88 float64 = 0.0
                                                _ = v87
                                                _ = v88
                                            }
                                            if (!false) {
                                                v89 := GenS4{F0: 0, F1: -377, F2: ((int32(v4) > 0) || false)}
                                                switch -9223372036854775808 {
                                                case 42:
                                                    v2 += 0
                                                case 128:
                                                    fmt.Println(fi5, fi5)
                                                    v90 := GenS12{F0: (^(fi8)), F1: -128}
                                                    _ = v90
                                                default:
                                                    v2 += v2
                                                }
                                                v4 += 127
                                                _ = v89
                                            } else {
                                                fmt.Println(fi5)
                                            }
                                        }
                                    }
                                    v1 = append(v1, -9223372036854775808)
                                }
                            case "":
                                v2 = (v2 ^ (v2 % (((v2 | (v2 - v2))) | 1)))
                            }
                            v91 := GenS13{F0: 0, F1: 747, F2: (fi5 + (fi8 ^ -2147483648))}
                            var v92 bool = (!true)
                            _ = v91
                            _ = v92
                        }
                    case "hello":
                        v2 -= (v2 << ((v2) & 63))
                    default:
                        if true {
                            var v93 float64 = float64(1.5)
                            var v94 int16 = 1
                            if ((v93 / v93) == 0.0) {
                                if (-1 > 9223372036854775807) {
                                    if false {
                                        v95 := GenS16{F0: v4, F1: v4}
                                        _ = v95
                                    } else {
                                        fmt.Println(fi8, fi8, v93)
                                        if (128 >= -1) {
                                            var v96 bool = false
                                            _ = v96
                                        } else {
                                            v97 := GenS10{F0: v94, F1: int32(fi8), F2: (fi5 << ((v2) & 63))}
                                            _ = v97
                                        }
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        v93 -= v93
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = len(v1)
                                }
                                v4 = v4
                                _ = len(v3)
                                if false {
                                    _ = len(v3)
                                } else {
                                    switch fi8 {
                                    case -609:
                                        v3 = append(v3, 0)
                                    case -2147483648:
                                        var v98 int16 = v94
                                        _ = v98
                                    case 363:
                                        v2 += (v2 + v2)
                                        var v99 int64 = 0
                                        v100 := GenS3{F0: 832}
                                        v4 = (^(v4))
                                        if (!true) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            switch (fi5 | 1) {
                                            case -9223372036854775808:
                                                if ((v2 << ((128) & 63)) < uint(v4)) {
                                                    fmt.Println(v99)
                                                }
                                            default:
                                                v101 := GenS17{F0: v99}
                                                _ = v101
                                            }
                                            fmt.Println(fi8)
                                            var v102 string = "foo bar"
                                            v103 := GenS18{F0: 2}
                                            var v104 int64 = v99
                                            switch (fi5 + (fi8 + 32767)) {
                                            case 72:
                                                var v105 int64 = v99
                                                _ = v105
                                            case 4294967295:
                                                switch 128 {
                                                case 2147483647:
                                                    var v106 uint16 = 42
                                                    _ = v106
                                                }
                                            }
                                            _ = v102
                                            _ = v103
                                            _ = v104
                                        } else {
                                            v4 += 127
                                            v4 -= v4
                                        }
                                        _ = v100
                                        _ = v99
                                    default:
                                        if true {
                                            v107 := GenS16{F0: 2, F1: (v4 * 128)}
                                            v108 := GenS14{F0: 255, F1: ("foo bar" + ""), F2: 1}
                                            v109 := GenS19{F0: v4, F1: 139}
                                            _ = v107
                                            _ = v108
                                            _ = v109
                                        }
                                    }
                                    var v110 int8 = 0
                                    var v111 uint8 = (v4 ^ (v4 ^ (v4 * (v4 + v4))))
                                    _ = v110
                                    _ = v111
                                }
                            } else {
                                if (0 < 728) {
                                    var v112 bool = (0.0 != 1e10)
                                    _ = v112
                                } else {
                                    fmt.Println(v94)
                                    v113 := GenS11{F0: uint32(v4), F1: 0}
                                    v4 -= v4
                                    _ = v113
                                }
                                v114 := GenS18{F0: 475}
                                var v115 float64 = (3.14159 / 1.5)
                                _ = v114
                                _ = v115
                            }
                            fmt.Println(v4, v2)
                            _ = v93
                            _ = v94
                        }
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                }
                _ = len(v1)
                _ = fi8
            }
            var v116 int = 32767
            _ = v116
        }
        v2 += 625
        _ = fi5
    }
    switch 32768 {
    case -578:
        v2 = v2
    case -141:
        v2 -= (v2 >> ((v2) & 63))
    case 128:
        v2 = (v2 / ((18446744073709551615) | 1))
    }
    var v117 bool = false
    var v118 int8 = -1
    for fi119 := 0; fi119 < 1; fi119++ {
        fmt.Println(fi119, v117, v4)
        _ = len(v3)
        var v120 int64 = -140
        _ = fi119
        _ = v120
    }
    if true {
        for fi121 := 0; fi121 < 3; fi121++ {
            v2 += v2
            _ = fi121
        }
    } else {
        var v122 bool = (3.14159 != (0.1 - 0.1))
        var v123 int8 = v118
        var v124 float64 = 1.5
        _ = v122
        _ = v123
        _ = v124
    }
    v117 = (1000 >= -867)
    v118 = 100
    var v125 int32 = 42
    v126 := GenS11{F0: 2147483648, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v125:", v125)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v117
    _ = v118
    _ = v125
    _ = v126
    _ = v2
    _ = v3
    _ = v4
}
