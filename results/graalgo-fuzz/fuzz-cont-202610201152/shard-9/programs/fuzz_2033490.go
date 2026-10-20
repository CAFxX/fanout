package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int64; F1 uint; F2 uint8 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int16; F1 int64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint32; F1 int8; F2 int }
type GenS15 struct { F0 int; F1 int32; F2 int8 }
type GenS16 struct { F0 int16; F1 int; F2 int16 }
type GenS17 struct { F0 int32; F1 int32; F2 int64 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 int }
type GenS20 struct { F0 bool; F1 uint; F2 uint64 }
type GenS21 struct { F0 int16; F1 uint16; F2 int16 }
type GenS22 struct { F0 int16 }
func genFn0() uint64 {
    v4 := []int{9223372036854775807, 42, -1, 32768}
    if ("a" != "a") {
        var v5 uint64 = 838
        if false {
            var v6 float32 = 3.14159
            v6 = 3.14159
            v6 = (v6 + v6)
            _ = v6
        } else {
            v7 := GenS1{F0: false, F1: true}
            if (0 < 0) {
                v8 := GenS2{F0: 32767}
                fmt.Println(v5)
                var v9 uint8 = 127
                _ = v8
                _ = v9
            } else {
                v5 = v5
            }
            v10 := GenS3{F0: 255, F1: 85}
            _ = v10
            _ = v7
        }
        v5 -= v5
        _ = v5
    } else {
        var v11 uint16 = 860
        switch (("hello" + "foo bar") + ("" + "")) {
        case "a":
            if (2 < 2) {
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 uint8 = 2
                    var v14 uint8 = v13
                    for fi15 := 0; fi15 < 14; fi15++ {
                        fmt.Println(v11, fi15, v14)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi15
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                if false {
                    if true {
                        fmt.Println(v11)
                        fmt.Println(v11)
                    } else {
                        v4 = append(v4, -410)
                        if false {
                            fmt.Println(v11)
                            var v16 int = 4294967295
                            switch "\n\t" {
                            case "foo bar":
                                v17 := GenS4{F0: -767, F1: 2, F2: 167}
                                _ = v17
                            case "\n\t":
                                v11 -= 42
                            case "":
                                if (!false) {
                                    if len(v4) > 0 {
                                        v4[0] = 447
                                    }
                                    for fi18 := 0; fi18 < 5; fi18++ {
                                        var v19 uint32 = 2147483648
                                        _ = fi18
                                        _ = v19
                                    }
                                    v16 = (v16 / ((v16) | 1))
                                } else {
                                    var v20 float64 = 3.14159
                                    _ = v20
                                }
                            default:
                                var v21 bool = ((1e10 - 1e10) > (1.5 / -2.25))
                                _ = v21
                            }
                            _ = v16
                        } else {
                            if (false || false) {
                                switch "hello" {
                                case "a":
                                    fmt.Println(v11)
                                }
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    var v23 int32 = -1
                                    _ = fi22
                                    _ = v23
                                }
                            } else {
                                v11 = (^(v11))
                            }
                            if false {
                                var v24 int = 128
                                switch v24 {
                                case 895:
                                    for fi25 := 0; fi25 < 17; fi25++ {
                                        for fi26 := 0; fi26 < 7; fi26++ {
                                            v27 := GenS2{F0: -2147483648}
                                            fmt.Println(v11, v11, v11)
                                            switch (fi26 + v24) {
                                            case 9223372036854775807:
                                                v28 := GenS5{F0: 1, F1: 42}
                                                _ = v28
                                            case 32767:
                                                var v29 uint = 42
                                                _ = v29
                                            case 0:
                                                v30 := GenS1{F0: (2147483647 >= 2147483648), F1: true}
                                                _ = v30
                                            default:
                                                fmt.Println(fi26, v11)
                                            }
                                            if true {
                                                var v31 uint16 = (v11 * (v11 * v11))
                                                switch (^(fi25)) {
                                                case -347:
                                                    switch ("hello" + "a") {
                                                    case "a":
                                                        var v32 int8 = 0
                                                        _ = v32
                                                    case "":
                                                        fmt.Println(v11, fi26)
                                                    }
                                                default:
                                                    var v33 string = "a"
                                                    _ = v33
                                                }
                                                if true {
                                                    var v34 int16 = 1000
                                                    if (int16(fi26) <= (-(v34))) {
                                                        if (((false && true) || (!true)) && (!true)) {
                                                            v4 = append(v4, (fi25 << ((1) & 63)))
                                                            v35 := GenS6{F0: 421, F1: 1}
                                                            switch ("\n\t" + "hello") {
                                                            case "":
                                                                if (!(("" >= ("foo bar" + "")) && false)) {
                                                                    v31 += (v31 << ((0) & 63))
                                                                }
                                                            default:
                                                                var v36 float64 = -2.25
                                                                _ = v36
                                                            }
                                                            _ = v35
                                                        } else {
                                                            var v37 int16 = v34
                                                            if false {
                                                                var v38 int32 = -1
                                                                var v39 int16 = 1000
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = fi26
                                                            }
                                                            _ = v37
                                                        }
                                                    }
                                                    _ = v34
                                                }
                                                _ = v31
                                            } else {
                                                if false {
                                                    var v40 int = fi25
                                                    _ = v40
                                                } else {
                                                    var v41 uint32 = 65535
                                                    if false {
                                                        _ = len(v4)
                                                        var v42 int32 = -2147483648
                                                        _ = v42
                                                    }
                                                    _ = v41
                                                }
                                                fmt.Println(fi25)
                                                if (!false) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v43 int64 = -9223372036854775808
                                                    _ = v43
                                                } else {
                                                    var v44 int64 = 1
                                                    v45 := GenS4{F0: (v44 + (v44 | v44)), F1: 753, F2: 255}
                                                    if false {
                                                        if (109 == 127) {
                                                            v46 := GenS7{F0: 42, F1: 2}
                                                            var v47 uint = 476
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            fmt.Println(fi26, fi25, v24)
                                                            if true {
                                                                v11 += 392
                                                                var v48 uint8 = 255
                                                                var v49 int8 = 113
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                var v50 float64 = 1e10
                                                                switch fi25 {
                                                                case 2147483647:
                                                                    if true {
                                                                        var v51 float64 = ((0.0 - (float64(3.14159) * (v50 - v50))) * ((0.0 * v50) / (-2.25 + 3.14159)))
                                                                        v52 := GenS8{F0: (v44 << ((942) & 63))}
                                                                        v44 -= (v44 & (v44 & v44))
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                case 127:
                                                                    switch 32767 {
                                                                    case 4294967295:
                                                                        var v53 float64 = (1.5 - 3.14159)
                                                                        _ = v53
                                                                    default:
                                                                        switch "a" {
                                                                        case "a":
                                                                            switch ((("" + "foo bar") + ("\n\t" + "")) + "a") {
                                                                            case "\n\t":
                                                                                v44 -= (-(v44))
                                                                                fmt.Println(v50, fi26)
                                                                            case "hello":
                                                                                var v54 int = (-(fi26))
                                                                                _ = v54
                                                                            case "a":
                                                                                var v55 int8 = -48
                                                                                _ = v55
                                                                            default:
                                                                                var v56 float32 = (0.1 * 1.5)
                                                                                _ = v56
                                                                            }
                                                                        default:
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (fi26 * (fi25 | 42))
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                v57 := GenS5{F0: (v44 / (((v44 * v44)) | 1)), F1: -9223372036854775808}
                                                                                _ = v57
                                                                            case "a":
                                                                                fmt.Println(v50, v50)
                                                                            case "hello":
                                                                                _ = len(v4)
                                                                            default:
                                                                                v4 = append(v4, fi25)
                                                                            }
                                                                            if (!true) {
                                                                                switch fi25 {
                                                                                case 2147483647:
                                                                                    if false {
                                                                                        if (1 <= -2147483648) {
                                                                                            var v58 uint64 = 394
                                                                                            v50 -= (v50 + 0.0)
                                                                                            if false {
                                                                                                fmt.Println(v11, v44)
                                                                                                var v59 uint16 = 281
                                                                                                v4 = append(v4, 128)
                                                                                                _ = v59
                                                                                            }
                                                                                            _ = v58
                                                                                        }
                                                                                        fmt.Println(v24, v44, fi26)
                                                                                    } else {
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = (-(v24))
                                                                                            }
                                                                                            switch ("a" + "") {
                                                                                            case "":
                                                                                                if (("\n\t" >= ("" + "")) || false) {
                                                                                                    v44 -= 317
                                                                                                    switch (fi25 << ((4294967295) & 63)) {
                                                                                                    case 0:
                                                                                                        var v60 string = "a"
                                                                                                        _ = v60
                                                                                                    }
                                                                                                } else {
                                                                                                    switch "hello" {
                                                                                                    case "":
                                                                                                        v24 -= fi25
                                                                                                    case "foo bar":
                                                                                                        var v61 float64 = (1e10 - (v50 - 1.5))
                                                                                                        _ = v61
                                                                                                    default:
                                                                                                        var v62 uint = uint(v11)
                                                                                                        _ = v62
                                                                                                    }
                                                                                                    v63 := GenS5{F0: v44, F1: 0}
                                                                                                    _ = v63
                                                                                                }
                                                                                            case "a":
                                                                                                if (false || (1 == 1)) {
                                                                                                    v11 += v11
                                                                                                    v11 -= v11
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        var v64 uint16 = 42
                                                                                                        _ = v64
                                                                                                    case "foo bar":
                                                                                                        var v65 string = ((("hello" + ("a" + "\n\t")) + "a") + "\n\t")
                                                                                                        _ = v65
                                                                                                    }
                                                                                                }
                                                                                            case "\n\t":
                                                                                                fmt.Println(v50, fi26, fi26)
                                                                                            default:
                                                                                                v66 := GenS9{F0: 100}
                                                                                                _ = v66
                                                                                            }
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                        case "\n\t":
                                                                                            fmt.Println(fi26, v50)
                                                                                        default:
                                                                                            v4 = append(v4, 4294967295)
                                                                                        }
                                                                                        v67 := GenS1{F0: (!false), F1: true}
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        _ = v67
                                                                                    }
                                                                                case -803:
                                                                                    switch "" {
                                                                                    case "":
                                                                                        v50 = 0.1
                                                                                    case "\n\t":
                                                                                        v68 := GenS10{F0: 187, F1: 699}
                                                                                        _ = v68
                                                                                    }
                                                                                }
                                                                                var v69 uint = 895
                                                                                v70 := GenS11{F0: (v11 + (^(v11)))}
                                                                                _ = v69
                                                                                _ = v70
                                                                            } else {
                                                                                fmt.Println(v11)
                                                                                var v71 bool = false
                                                                                _ = v71
                                                                            }
                                                                        }
                                                                    }
                                                                case -1:
                                                                    var v72 bool = (!false)
                                                                    _ = v72
                                                                }
                                                                v73 := GenS12{F0: 1, F1: v44}
                                                                _ = v50
                                                                _ = v73
                                                            }
                                                            var v74 uint8 = 161
                                                            if (true && (((v74 % (((v74 >> ((382) & 63))) | 1)) != 128) || (9223372036854775807 <= (v44 % ((v44) | 1))))) {
                                                                if false {
                                                                    _ = len(v4)
                                                                    if true {
                                                                        v74 -= v74
                                                                        var v75 bool = ((v44 & v44) >= (v44 + (v44 & (v44 >> ((0) & 63)))))
                                                                        var v76 uint32 = 438
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                } else {
                                                                    fmt.Println(fi26, fi26, v74)
                                                                    fmt.Println(v11)
                                                                }
                                                            } else {
                                                                _ = len(v4)
                                                                if true {
                                                                    v77 := GenS13{F0: 42}
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    switch v24 {
                                                                    case 0:
                                                                        v24 += 662
                                                                    case 740:
                                                                        v44 += v44
                                                                    default:
                                                                        if false {
                                                                            var v78 int = 42
                                                                            fmt.Println(v11)
                                                                            _ = v78
                                                                        }
                                                                    }
                                                                    _ = v77
                                                                } else {
                                                                    fmt.Println(v74)
                                                                    if false {
                                                                        fmt.Println(fi25, fi25)
                                                                        v79 := GenS9{F0: 43}
                                                                        var v80 float64 = -2.25
                                                                        var v81 string = "foo bar"
                                                                        if false {
                                                                            switch fi26 {
                                                                            case -627:
                                                                                switch v24 {
                                                                                case 32767:
                                                                                    if (((v81 + v81) + "") == v81) {
                                                                                        var v82 int16 = 32767
                                                                                        _ = v82
                                                                                    } else {
                                                                                        fmt.Println(v81, v24, v11)
                                                                                        var v83 int16 = 42
                                                                                        _ = v83
                                                                                    }
                                                                                case -973:
                                                                                    var v84 string = v81
                                                                                    _ = v84
                                                                                default:
                                                                                    v85 := GenS14{F0: 2, F1: 0, F2: (fi26 + (fi25 | fi25))}
                                                                                    _ = v85
                                                                                }
                                                                                fmt.Println(fi25, v74)
                                                                                v86 := GenS15{F0: v24, F1: 0, F2: -1}
                                                                                _ = v86
                                                                            case 9223372036854775807:
                                                                                if true {
                                                                                    if (0 == (v74 % (((v74 * 255)) | 1))) {
                                                                                        var v87 int = (^(fi26))
                                                                                        var v88 int = fi26
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                    } else {
                                                                                        v74 = v74
                                                                                    }
                                                                                    v81 = ("foo bar" + "\n\t")
                                                                                    v80 += ((0.1 * (1.5 + ((float64(0.0) + float64(1.5)) * float64(0.1)))) - (v80 + v80))
                                                                                } else {
                                                                                    v89 := GenS16{F0: 0, F1: (fi25 << ((1) & 63)), F2: 0}
                                                                                    v80 += v80
                                                                                    fmt.Println(v81)
                                                                                    _ = v89
                                                                                }
                                                                            }
                                                                            fmt.Println(v44, v74, v44)
                                                                            _ = len(v4)
                                                                        } else {
                                                                            fmt.Println(v74, fi26, v24)
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        }
                                                                        v44 = 0
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                    }
                                                                    fmt.Println(v24, v44)
                                                                }
                                                                v44 += 0
                                                            }
                                                            v44 = (v44 % (((^(v44))) | 1))
                                                            _ = v74
                                                        }
                                                        if (("foo bar" + "foo bar") < "hello") {
                                                            if true {
                                                                var v90 uint8 = 197
                                                                _ = v90
                                                            } else {
                                                                var v91 int8 = 127
                                                                v11 += v11
                                                                var v92 int64 = 0
                                                                var v93 int64 = v44
                                                                _ = v91
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                            switch "a" {
                                                            case "":
                                                                var v94 uint = 1
                                                                _ = v94
                                                            }
                                                            var v95 uint8 = 42
                                                            _ = v95
                                                        }
                                                        var v96 int64 = v44
                                                        _ = v96
                                                    } else {
                                                        v44 = v44
                                                    }
                                                    _ = v44
                                                    _ = v45
                                                }
                                            }
                                            _ = fi26
                                            _ = v27
                                        }
                                        for fi97 := 0; fi97 < 3; fi97++ {
                                            v24 += (fi97 & fi25)
                                            if true {
                                                fmt.Println(fi25)
                                            } else {
                                                if true {
                                                    if ((v24 << ((0) & 63)) != fi25) {
                                                        v24 -= (fi97 >> ((65535) & 63))
                                                    }
                                                } else {
                                                    v11 += 32767
                                                }
                                            }
                                            _ = fi97
                                        }
                                        v98 := GenS17{F0: 707, F1: -2147483648, F2: 0}
                                        _ = fi25
                                        _ = v98
                                    }
                                default:
                                    switch v24 {
                                    case 128:
                                        fmt.Println(v11, v24)
                                    default:
                                        var v99 bool = false
                                        _ = v99
                                    }
                                }
                                _ = v24
                            } else {
                                if (true && true) {
                                    for fi100 := 0; fi100 < 13; fi100++ {
                                        fmt.Println(v11, fi100)
                                        _ = fi100
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (-2.25 >= 3.14159) {
                                        var v101 uint = 809
                                        _ = v101
                                    } else {
                                        switch 0 {
                                        case 1:
                                            v4 = append(v4, -1)
                                        case -693:
                                            switch 0 {
                                            case 4294967295:
                                                var v102 uint64 = 2
                                                _ = v102
                                            case 32767:
                                                if false {
                                                    if (783 > v11) {
                                                        fmt.Println(v11)
                                                        if len(v4) > 0 {
                                                            v4[0] = 32767
                                                        }
                                                        for fi103 := 0; fi103 < 10; fi103++ {
                                                            if (!(!false)) {
                                                                fmt.Println(v11, fi103)
                                                            } else {
                                                                var v104 float64 = (3.14159 - 0.1)
                                                                v105 := GenS9{F0: -119}
                                                                _ = v104
                                                                _ = v105
                                                            }
                                                            _ = fi103
                                                        }
                                                    } else {
                                                        for fi106 := 0; fi106 < 20; fi106++ {
                                                            var v107 float32 = 0.1
                                                            _ = fi106
                                                            _ = v107
                                                        }
                                                        for fi108 := 0; fi108 < 6; fi108++ {
                                                            v109 := GenS10{F0: 2147483647, F1: 4294967295}
                                                            var v110 uint32 = 2147483647
                                                            _ = fi108
                                                            _ = v109
                                                            _ = v110
                                                        }
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = -116
                                                    }
                                                } else {
                                                    v111 := GenS2{F0: -1}
                                                    _ = v111
                                                }
                                            default:
                                                var v112 string = ("\n\t" + "a")
                                                _ = v112
                                            }
                                        }
                                        v11 -= v11
                                    }
                                } else {
                                    v11 = (v11 >> ((1) & 63))
                                    var v113 float32 = 1e10
                                    v113 = float32(-2.25)
                                    _ = v113
                                }
                            }
                        }
                    }
                }
            }
        case "hello":
            fmt.Println(v11)
        }
        v4 = append(v4, 1)
        _ = v11
    }
    v114 := GenS18{F0: 128}
    _ = len(v4)
    _ = v114
    _ = v4
    return 1
}

func main() {
    v1 := []int{127, -97, -244}
    var v2 uint8 = 161
    var v3 uint32 = 0
    var v115 int8 = 121
    for fi116 := 0; fi116 < 5; fi116++ {
        v117 := GenS19{F0: (fi116 >> ((2) & 63))}
        var v118 float32 = 0.1
        _ = fi116
        _ = v117
        _ = v118
    }
    switch "\n\t" {
    case "":
        var v119 int32 = -2147483648
        _ = v119
    case "a":
        v3 = v3
    }
    var v120 float32 = 0.0
    switch -1 {
    case 0:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if ((((-311 <= -976) && true) || (-1 >= 581)) && true) {
            for fi121 := 0; fi121 < 3; fi121++ {
                var v122 uint8 = 16
                _ = len(v1)
                var v123 float64 = (1e10 * (((-2.25 * ((0.1 + (3.14159 + 1e10)) + 3.14159)) - (-2.25 + ((float64(v120) - -2.25) - (float64(3.14159) * (0.0 - -2.25))))) - ((((float64(v120) / float64(v120)) - 1e10) / -2.25) + 0.0)))
                _ = fi121
                _ = v122
                _ = v123
            }
        }
    }
    var v124 bool = (false && (((float32(1e10) + (v120 - ((0.0 + v120) * 3.14159))) != v120) && (true && true)))
    var v125 uint = 4294967295
    for fi126 := 0; fi126 < 11; fi126++ {
        switch fi126 {
        case 9223372036854775807:
            switch "foo bar" {
            case "foo bar":
                if false {
                    switch ((("\n\t" + ((("\n\t" + ("foo bar" + "a")) + ("\n\t" + "foo bar")) + (("hello" + "hello") + (("hello" + "hello") + "\n\t")))) + (("\n\t" + "hello") + "foo bar")) + "\n\t") {
                    case "":
                        var v127 float64 = (((0.1 * 1.5) - ((0.1 * 3.14159) * 3.14159)) - -2.25)
                        _ = v127
                    case "\n\t":
                        switch ("foo bar" + "") {
                        case "":
                            v3 = 0
                        case "a":
                            v120 -= v120
                        case "foo bar":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            var v128 int8 = (-(v115))
                            _ = v128
                        }
                    case "a":
                        for fi129 := 0; fi129 < 14; fi129++ {
                            fmt.Println(fi126)
                            _ = fi129
                        }
                    default:
                        v2 = v2
                    }
                    v130 := GenS18{F0: (v2 % (((v2 << ((634) & 63))) | 1))}
                    _ = v130
                } else {
                    var v131 bool = (-896 <= 42)
                    for fi132 := 0; fi132 < 11; fi132++ {
                        if (!(42 != (^(v2)))) {
                            var v133 int16 = 1
                            _ = v133
                        }
                        switch (fi132 << ((0) & 63)) {
                        case 32767:
                            if false {
                                var v134 float64 = 0.0
                                _ = v134
                            } else {
                                v131 = v131
                            }
                        case 1:
                            if (!((fi132 * fi132) <= (fi126 >> (((^(v125))) & 63)))) {
                                v3 = 509
                                v135 := GenS19{F0: fi132}
                                var v136 float64 = (-2.25 - (1.5 / 1.5))
                                _ = v135
                                _ = v136
                            }
                        default:
                            if ((v124 && v131) || v131) {
                                v137 := GenS12{F0: -1, F1: 1}
                                var v138 float64 = float64(0.0)
                                _ = v137
                                _ = v138
                            } else {
                                v3 += v3
                                if v124 {
                                    if ((v2 % ((42) | 1)) != v2) {
                                        var v139 int = fi126
                                        _ = v139
                                    }
                                }
                                v140 := GenS20{F0: (313 < fi132), F1: (v125 >> ((2) & 63)), F2: 4294967295}
                                _ = v140
                            }
                        }
                        v131 = false
                        var v141 uint = 18446744073709551615
                        _ = fi132
                        _ = v141
                    }
                    _ = v131
                }
            case "a":
                var v142 float64 = (float64(1e10) + (((1.5 + float64(0.0)) - 1e10) - 1.5))
                _ = v142
            case "":
                if false {
                    var v143 bool = false
                    switch "\n\t" {
                    case "hello":
                        v144 := GenS8{F0: 771}
                        _ = v144
                    case "":
                        var v145 float32 = (v120 * float32(((float64((-2.25 - (0.1 * (v120 / 0.1)))) * 0.0) / -2.25)))
                        _ = v145
                    default:
                        fmt.Println(v2, v2, v143)
                    }
                    _ = v143
                }
            }
        case 32767:
            var v146 uint8 = v2
            _ = v146
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v124 = (!(false && v124))
        }
        if ((32767 <= 42) && false) {
            v115 = v115
        }
        _ = fi126
    }
    v2 = 0
    var v147 uint32 = 1
    switch 32767 {
    case 25:
        switch 4294967295 {
        case 127:
            if ((!(0.1 > (-2.25 / 1e10))) || (v115 < v115)) {
                var v148 uint = 0
                switch "a" {
                case "a":
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                case "\n\t":
                    v2 += v2
                }
                _ = v148
            }
        default:
            for fi149 := 0; fi149 < 20; fi149++ {
                for fi150 := 0; fi150 < 9; fi150++ {
                    switch "hello" {
                    case "a":
                        var v151 uint8 = v2
                        _ = v151
                    case "hello":
                        if (false && (!(-897 == 32768))) {
                            var v152 float32 = (3.14159 / (float32((0.1 * (1e10 - (((0.1 - ((1.5 * 0.1) / 1e10)) - -2.25) + 1.5)))) * 1.5))
                            _ = v152
                        }
                        _ = len(v1)
                    }
                    v120 -= v120
                    var v153 uint16 = 441
                    _ = fi150
                    _ = v153
                }
                _ = fi149
            }
        }
    case -528:
        v154 := GenS3{F0: v2, F1: 2147483647}
        _ = v154
    default:
        var v155 uint64 = uint64(v2)
        _ = v155
    }
    var v156 float32 = v120
    if ((v115 | (-(v115))) <= (v115 + (v115 ^ 0))) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((32767 <= -498) && false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v1 = append(v1, 636)
            var v157 uint8 = v2
            _ = v157
        } else {
            if false {
                var v158 uint64 = 4294967295
                switch 42 {
                case -9223372036854775808:
                    v159 := GenS8{F0: 1}
                    _ = v159
                case 537:
                    var v160 bool = (!v124)
                    _ = v160
                default:
                    if (255 < v125) {
                        v161 := GenS9{F0: 127}
                        if (1 == -1) {
                            v125 = uint(v2)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        if v124 {
                            if (float64(v156) > float64(0.0)) {
                                var v162 uint64 = v158
                                _ = v162
                            }
                            fmt.Println(v115)
                        }
                        _ = v161
                    }
                }
                _ = v158
            } else {
                var v163 uint64 = 931
                switch "\n\t" {
                case "hello":
                    if true {
                        fmt.Println(v125, v120, v163)
                        var v164 int16 = 1000
                        fmt.Println(v164, v125)
                        _ = v164
                    } else {
                        v124 = v124
                        v165 := GenS11{F0: 0}
                        _ = v165
                    }
                case "":
                    fmt.Println(v125, v2, v115)
                default:
                    var v166 float64 = 0.0
                    _ = v166
                }
                for fi167 := 0; fi167 < 10; fi167++ {
                    if (v3 == (v3 << ((v125) & 63))) {
                        v2 -= uint8(v125)
                        var v168 int8 = (v115 + v115)
                        _ = v168
                    } else {
                        for fi169 := 0; fi169 < 11; fi169++ {
                            if false {
                                _ = len(v1)
                            }
                            var v170 uint64 = v163
                            _ = fi169
                            _ = v170
                        }
                        fmt.Println(v120, fi167, v3)
                        v147 = (v3 ^ (v3 & uint32(v163)))
                    }
                    _ = fi167
                }
                _ = v163
            }
            var v171 bool = (v2 != (v2 % ((42) | 1)))
            for fi172 := 0; fi172 < 9; fi172++ {
                v2 += 0
                if v124 {
                    v3 -= (v147 * (v3 << (((v125 / ((0) | 1))) & 63)))
                } else {
                    v173 := GenS10{F0: (fi172 ^ fi172), F1: 42}
                    v125 = v125
                    _ = v173
                }
                _ = fi172
            }
            _ = v171
        }
    } else {
        fmt.Println(v124, v147)
        if false {
            var v174 string = "hello"
            _ = v174
        } else {
            if ("\n\t" <= ("" + "hello")) {
                v2 += v2
                v175 := GenS8{F0: -9223372036854775808}
                _ = v175
            } else {
                v124 = v124
                v120 = float32((0.1 - 1.5))
                for fi176 := 0; fi176 < 3; fi176++ {
                    v115 += 42
                    for fi177 := 0; fi177 < 13; fi177++ {
                        v125 -= 333
                        if v124 {
                            switch ("hello" + ("" + "hello")) {
                            case "a":
                                if true {
                                    var v178 uint32 = (v3 / ((42) | 1))
                                    var v179 uint8 = v2
                                    _ = v178
                                    _ = v179
                                }
                            case "":
                                if (42 < 1) {
                                    fmt.Println(fi177)
                                    v147 = v147
                                    var v180 uint32 = (^(v3))
                                    _ = v180
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch -55 {
                                    case 1:
                                        v1 = append(v1, fi177)
                                    case 0:
                                        switch "foo bar" {
                                        case "a":
                                            v1 = append(v1, 32768)
                                        case "foo bar":
                                            if len(v1) > 0 {
                                                v1[0] = -9223372036854775808
                                            }
                                        case "hello":
                                            v181 := GenS21{F0: 1000, F1: 1, F2: 1}
                                            _ = v181
                                        default:
                                            var v182 int = (fi177 ^ fi177)
                                            _ = v182
                                        }
                                    }
                                    var v183 uint16 = 0
                                    fmt.Println(v147)
                                    var v184 float64 = 1e10
                                    _ = v183
                                    _ = v184
                                }
                            case "hello":
                                fmt.Println(v115, fi176, v115)
                            }
                        } else {
                            v185 := GenS22{F0: 599}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v186 uint = v125
                            var v187 float32 = v156
                            var v188 uint = uint(v2)
                            fmt.Println(v187, v3)
                            _ = v185
                            _ = v186
                            _ = v187
                            _ = v188
                        }
                        var v189 float64 = ((3.14159 * (1.5 / 1.0)) * 0.0)
                        var v190 uint64 = 4294967295
                        v191 := GenS12{F0: 42, F1: -584}
                        if len(v1) > 0 {
                            v1[0] = fi176
                        }
                        _ = fi177
                        _ = v189
                        _ = v190
                        _ = v191
                    }
                    _ = fi176
                }
            }
        }
    }
    v124 = false
    v192 := GenS2{F0: -715}
    var v193 int32 = -278
    for fi194 := 0; fi194 < 14; fi194++ {
        v115 = (v115 & (v115 / ((v115) | 1)))
        v193 = v193
        var v195 string = "a"
        _ = fi194
        _ = v195
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v115:", v115)
    fmt.Println("v120:", v120)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v147:", v147)
    fmt.Println("v156:", v156)
    fmt.Println("v193:", v193)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v115
    _ = v120
    _ = v124
    _ = v125
    _ = v147
    _ = v156
    _ = v192
    _ = v193
    _ = v2
    _ = v3
}
