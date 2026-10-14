package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 bool; F2 int8 }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 uint16; F1 int32; F2 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64; F1 uint64; F2 int16 }
type GenS12 struct { F0 uint32; F1 int; F2 int16 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint64; F1 int8 }
type GenS17 struct { F0 uint64; F1 uint }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int16; F1 uint64 }
type GenS20 struct { F0 uint; F1 uint16 }
type GenS21 struct { F0 int32; F1 int16 }

func main() {
    var v1 float64 = 1e10
    v2 := []int{1, 481, 0, 2147483647}
    v3 := make(map[uint32]string)
    v4 := []uint32{1}
    var v5 string = "a"
    var v6 bool = (9223372036854775807 <= 830)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint8 = 232
        var v9 int = (fi7 << ((18) & 63))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = (v5 + ("foo bar" + ("a" + v5)))
    for fi11 := 0; fi11 < 3; fi11++ {
        if ((2147483647 != 374) && v6) {
            if v6 {
                if v6 {
                    fmt.Println(v1)
                } else {
                    fmt.Println(v1, v5, v6)
                    var v12 uint16 = 32767
                    v4 = append(v4, 2147483648)
                    _ = v12
                }
            } else {
                var v13 uint8 = 0
                _ = v13
            }
        } else {
            var v14 uint8 = 128
            for fi15 := 0; fi15 < 17; fi15++ {
                v6 = true
                _ = fi15
            }
            _ = v14
        }
        _ = fi11
    }
    if ((4294967295 != 128) || ((65535 == 4294967295) && v6)) {
        var v16 int = 463
        if false {
            var v17 int32 = 670
            v1 -= v1
            _ = v17
        } else {
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := GenS1{F0: -150, F1: 100}
                delete(v3, 0)
                fmt.Println(v5, v10, v1)
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                for fi21 := 0; fi21 < 2; fi21++ {
                    if v6 {
                        var v22 uint64 = 292
                        switch v16 {
                        case 4294967295:
                            fmt.Println(v22, fi20)
                        case -113:
                            fmt.Println(v22, fi20, v10)
                        case 9223372036854775807:
                            switch (fi21 % ((fi20) | 1)) {
                            case 32768:
                                switch (((v10 + v5) + v10) + (v5 + (("hello" + (v10 + v5)) + v10))) {
                                case "hello":
                                    if (v22 == 2) {
                                        v16 -= 866
                                        var v23 uint32 = 1
                                        fmt.Println(fi20, v10)
                                        _ = v23
                                    }
                                case "":
                                    var v24 uint8 = 15
                                    _ = v24
                                case "foo bar":
                                    switch -1 {
                                    case -196:
                                        v6 = v6
                                    case 128:
                                        fmt.Println(v1, fi21)
                                    }
                                default:
                                    fmt.Println(v5, v5)
                                    if (1.5 == float64(0.0)) {
                                        var v25 int64 = 42
                                        v6 = v6
                                        _ = v25
                                    }
                                    v26 := GenS1{F0: (v16 | int(v22)), F1: 48}
                                    _ = v26
                                }
                            case 9223372036854775807:
                                var v27 uint16 = 243
                                _ = v27
                            case -9223372036854775808:
                                var v28 int8 = 0
                                if false {
                                    fmt.Println(v10, fi21, v6)
                                } else {
                                    switch (fi21 % (((-(v16))) | 1)) {
                                    case -2147483648:
                                        switch "foo bar" {
                                        case "a":
                                            v28 -= (^(v28))
                                        case "foo bar":
                                            v22 += (^(v22))
                                        }
                                    case 32768:
                                        v29 := GenS1{F0: (fi20 / ((-2147483648) | 1)), F1: (v28 * (v28 << ((42) & 63)))}
                                        _ = v29
                                    case 0:
                                        fmt.Println(v28, fi21, fi20)
                                    }
                                    switch (fi21 * (v16 ^ 127)) {
                                    case 0:
                                        if false {
                                            if v6 {
                                                _ = len(v3)
                                                var v30 uint = 188
                                                var v31 int64 = 42
                                                _ = v30
                                                _ = v31
                                            } else {
                                                v32 := GenS1{F0: fi20, F1: (v28 % (((v28 | (^(v28)))) | 1))}
                                                _ = v32
                                            }
                                        }
                                    case -9223372036854775808:
                                        var v33 int64 = -992
                                        _ = v33
                                    case 1:
                                        v34 := GenS2{F0: 0, F1: 513}
                                        _ = v34
                                    default:
                                        fmt.Println(v6, fi21, v6)
                                    }
                                }
                                var v35 float64 = (v1 - 1e10)
                                _ = v28
                                _ = v35
                            default:
                                var v36 string = v10
                                _ = v36
                            }
                            v22 += (v22 ^ 2)
                            switch v5 {
                            case "":
                                v16 += v16
                            }
                        }
                        var v37 float32 = 1.5
                        _ = v22
                        _ = v37
                    } else {
                        var v38 int = (-(fi21))
                        if (305 == 826) {
                            v10 = (v5 + (v10 + "hello"))
                            v1 -= float64(3.14159)
                            var v39 uint = 4294967295
                            v40 := GenS2{F0: 0, F1: 278}
                            _ = v39
                            _ = v40
                        }
                        switch "a" {
                        case "":
                            delete(v3, 374)
                        case "foo bar":
                            fmt.Println(v6)
                        case "\n\t":
                            if v6 {
                                v1 -= v1
                                v41 := v3[998]
                                if (v6 && (635 > 255)) {
                                    _ = len(v2)
                                    v41 = v41
                                    v1 = ((v1 / v1) / v1)
                                    if ((-(fi20)) >= (v16 / ((9223372036854775807) | 1))) {
                                        switch (fi20 + (fi20 << ((255) & 63))) {
                                        case 32767:
                                            if len(v2) > 0 {
                                                v2[0] = (fi20 / ((v16) | 1))
                                            }
                                        case 9223372036854775807:
                                            if len(v2) > 0 {
                                                v2[0] = v38
                                            }
                                        default:
                                            v1 = float64((1.5 / 3.14159))
                                        }
                                    } else {
                                        if v6 {
                                            if v6 {
                                                fmt.Println(v6, v16, fi21)
                                                if v6 {
                                                    var v42 bool = true
                                                    var v43 uint32 = 0
                                                    v3[0] = (v41 + v10)
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    var v44 uint32 = 2147483648
                                                    var v45 bool = false
                                                    _ = v44
                                                    _ = v45
                                                }
                                                switch fi20 {
                                                case 9223372036854775807:
                                                    fmt.Println(fi21, v5)
                                                case 503:
                                                    switch v10 {
                                                    case "":
                                                        var v46 float32 = float32(0.0)
                                                        _ = v46
                                                    case "a":
                                                        var v47 uint8 = 128
                                                        _ = v47
                                                    default:
                                                        v16 = (v16 | (v16 * (fi21 / ((v16) | 1))))
                                                    }
                                                case -9223372036854775808:
                                                    v41 = v5
                                                default:
                                                    if (!((v1 - -2.25) == 3.14159)) {
                                                        v6 = (!(1000 >= 42))
                                                    } else {
                                                        if (!(255 >= 2)) {
                                                            var v48 uint32 = 109
                                                            _ = v48
                                                        } else {
                                                            v16 = fi20
                                                            fmt.Println(fi21, v10, v38)
                                                            if len(v4) > 0 {
                                                                v4[0] = 42
                                                            }
                                                            if false {
                                                                v2 = append(v2, -2147483648)
                                                            }
                                                        }
                                                        if v6 {
                                                            if len(v2) > 0 {
                                                                v2[0] = fi20
                                                            }
                                                            fmt.Println(v5, v6, v41)
                                                            switch "foo bar" {
                                                            case "":
                                                                var v49 int = 42
                                                                _ = v49
                                                            case "hello":
                                                                v50 := GenS2{F0: 42, F1: 1}
                                                                _ = v50
                                                            default:
                                                                switch v16 {
                                                                case 234:
                                                                    var v51 int8 = 0
                                                                    _ = v51
                                                                case 32767:
                                                                    var v52 uint16 = 2
                                                                    _ = v52
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v38, fi20, fi20)
                                                            switch v10 {
                                                            case "\n\t":
                                                                v6 = (3.14159 < v1)
                                                            default:
                                                                v10 = v10
                                                            }
                                                            v53 := GenS1{F0: (^(fi21)), F1: 97}
                                                            _ = v53
                                                        }
                                                        v38 += (fi21 % ((fi21) | 1))
                                                        switch (v16 << ((42) & 63)) {
                                                        case -1:
                                                            var v54 int = fi21
                                                            _ = v54
                                                        case 2147483647:
                                                            var v55 string = "\n\t"
                                                            var v56 int8 = -128
                                                            v16 = int(v56)
                                                            switch (v55 + (v5 + (v10 + ("hello" + (v5 + (v5 + v41)))))) {
                                                            case "hello":
                                                                if v6 {
                                                                    var v57 string = "a"
                                                                    v58 := GenS2{F0: -32768, F1: 255}
                                                                    var v59 int64 = -1
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                            }
                                                            v5 = v41
                                                            if (622 != 18446744073709551615) {
                                                                v3[2147483648] = "\n\t"
                                                                if true {
                                                                    v56 = (v56 + (v56 * 40))
                                                                    fmt.Println(v16, v55)
                                                                }
                                                                v60 := GenS1{F0: fi20, F1: (v56 & (v56 << ((255) & 63)))}
                                                                _ = v60
                                                            } else {
                                                                v61 := GenS2{F0: -32768, F1: 2}
                                                                _ = v61
                                                            }
                                                            switch (((v5 + v41) + v5) + v41) {
                                                            case "foo bar":
                                                                var v62 float32 = 1.5
                                                                _ = v62
                                                            }
                                                            _ = v55
                                                            _ = v56
                                                        default:
                                                            if (65535 != 1) {
                                                                fmt.Println(v41, v16, v6)
                                                                v41 = "hello"
                                                            } else {
                                                                switch (v16 + fi20) {
                                                                case 128:
                                                                    v10 = ((v10 + "foo bar") + v5)
                                                                case -310:
                                                                    if (false || v6) {
                                                                        _ = len(v2)
                                                                        _ = len(v3)
                                                                        var v63 int = (v16 | fi20)
                                                                        _ = v63
                                                                    } else {
                                                                        v5 = v5
                                                                    }
                                                                case 32768:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        var v64 int32 = 0
                                                                        _ = v64
                                                                    case "":
                                                                        var v65 bool = v6
                                                                        _ = v65
                                                                    case "hello":
                                                                        v38 = (v16 ^ 42)
                                                                    }
                                                                }
                                                                if len(v2) > 0 {
                                                                    v2[0] = (fi20 + v16)
                                                                }
                                                                if (22 < 255) {
                                                                    var v66 int8 = -128
                                                                    switch (^(v16)) {
                                                                    case 4294967295:
                                                                        var v67 float32 = -2.25
                                                                        _ = v67
                                                                    case -2147483648:
                                                                        var v68 int = (v38 << ((18446744073709551615) & 63))
                                                                        if v6 {
                                                                            var v69 string = v41
                                                                            var v70 string = v69
                                                                            v71 := GenS3{F0: -63, F1: v6}
                                                                            _ = v69
                                                                            _ = v70
                                                                            _ = v71
                                                                        } else {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        }
                                                                        var v72 uint32 = 2147483647
                                                                        _ = v68
                                                                        _ = v72
                                                                    }
                                                                    var v73 bool = (fi20 < -9223372036854775808)
                                                                    _ = v66
                                                                    _ = v73
                                                                }
                                                                v10 = "\n\t"
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v5)
                                                        var v74 uint8 = 139
                                                        _ = v74
                                                    }
                                                }
                                            } else {
                                                if v6 {
                                                    v16 += v16
                                                    v16 = (-(fi21))
                                                    v75 := GenS2{F0: -1, F1: 2}
                                                    v3[4294967295] = (v5 + v5)
                                                    v76 := GenS4{F0: 2147483648}
                                                    if (!(true && false)) {
                                                        v77 := GenS5{F0: 0}
                                                        _ = v77
                                                    } else {
                                                        v38 += 988
                                                        v6 = (fi20 < (fi20 ^ 2147483647))
                                                        v78 := GenS4{F0: 0}
                                                        _ = v78
                                                    }
                                                    if v6 {
                                                        v41 = (v41 + v41)
                                                        if (2 == 742) {
                                                            v41 = ((v10 + ((v10 + "") + "a")) + "a")
                                                        } else {
                                                            var v79 uint16 = 42
                                                            var v80 uint = 4294967295
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        switch v16 {
                                                        case -625:
                                                            switch "a" {
                                                            case "a":
                                                                v16 = v38
                                                            }
                                                        default:
                                                            var v81 int32 = 1
                                                            _ = v81
                                                        }
                                                    }
                                                    _ = v75
                                                    _ = v76
                                                }
                                                switch (((("\n\t" + "\n\t") + v5) + "") + v41) {
                                                case "foo bar":
                                                    v16 = 2147483647
                                                default:
                                                    if true {
                                                        switch v5 {
                                                        case "\n\t":
                                                            if v6 {
                                                                var v82 int64 = -9223372036854775808
                                                                v83 := GenS4{F0: 946}
                                                                v84 := v3[4294967295]
                                                                _ = v82
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                        case "hello":
                                                            if true {
                                                                if true {
                                                                    if ((-2.25 + float32(float64(3.14159))) == (0.0 - (((0.1 / (-2.25 + 3.14159)) / (-2.25 / (((0.1 / 0.1) + float32(((v1 - 1e10) / 1.0))) - -2.25))) - 0.1))) {
                                                                        delete(v3, 2)
                                                                    } else {
                                                                        var v85 uint64 = 4294967295
                                                                        v1 -= v1
                                                                        _ = v85
                                                                    }
                                                                    fmt.Println(v10, fi21, v16)
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v86 := GenS4{F0: 641}
                                                                        _ = v86
                                                                    case "\n\t":
                                                                        switch v10 {
                                                                        case "foo bar":
                                                                            if v6 {
                                                                                if v6 {
                                                                                    v38 = (^(v16))
                                                                                } else {
                                                                                    switch v41 {
                                                                                    case "a":
                                                                                        switch 0 {
                                                                                        case -309:
                                                                                            if false {
                                                                                                switch fi20 {
                                                                                                case 4294967295:
                                                                                                    var v87 uint16 = 32767
                                                                                                    _ = v87
                                                                                                default:
                                                                                                    v88 := GenS5{F0: 65535}
                                                                                                    _ = v88
                                                                                                }
                                                                                                v4 = append(v4, 2147483647)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v1 -= v1
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            v89 := GenS6{F0: 943, F1: ("hello" < v5), F2: -128}
                                                                            _ = v89
                                                                        }
                                                                        switch -670 {
                                                                        case 610:
                                                                            var v90 uint64 = 18446744073709551615
                                                                            switch fi21 {
                                                                            case 32768:
                                                                                if v6 {
                                                                                    v1 = (v1 * v1)
                                                                                }
                                                                            }
                                                                            _ = v90
                                                                        case 9223372036854775807:
                                                                            v10 = v41
                                                                            v5 = "\n\t"
                                                                            v91 := v3[350]
                                                                            _ = v91
                                                                        case 267:
                                                                            switch ("\n\t" + "foo bar") {
                                                                            case "":
                                                                                v3[141] = v5
                                                                            case "\n\t":
                                                                                fmt.Println(v10, v16, v10)
                                                                            case "a":
                                                                                fmt.Println(v1)
                                                                            default:
                                                                                var v92 string = v10
                                                                                var v93 int64 = -1
                                                                                _ = v92
                                                                                _ = v93
                                                                            }
                                                                        }
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    default:
                                                                        var v94 uint8 = 139
                                                                        _ = v94
                                                                    }
                                                                } else {
                                                                    v95 := GenS7{F0: 42, F1: -2147483648}
                                                                    if (885 < 781) {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    } else {
                                                                        v10 = "hello"
                                                                    }
                                                                    v41 = "\n\t"
                                                                    _ = v95
                                                                }
                                                                fmt.Println(v16, v38)
                                                            } else {
                                                                fmt.Println(v6)
                                                            }
                                                            var v96 bool = (1.5 > 0.1)
                                                            _ = v96
                                                        }
                                                        v97 := GenS8{F0: 42, F1: 511}
                                                        v98 := GenS7{F0: 238, F1: -2147483648}
                                                        _ = v97
                                                        _ = v98
                                                    } else {
                                                        var v99 uint16 = 32767
                                                        _ = v99
                                                    }
                                                    v100 := GenS2{F0: -32768, F1: 65535}
                                                    _ = len(v3)
                                                    if v6 {
                                                        if false {
                                                            fmt.Println(v6)
                                                            var v101 int32 = 1
                                                            _ = v101
                                                        } else {
                                                            if ((855 == 713) && v6) {
                                                                v10 = v5
                                                                v5 = v5
                                                                v1 -= (v1 * v1)
                                                            } else {
                                                                if (v6 || false) {
                                                                    v3[2] = v5
                                                                }
                                                                v38 -= 2147483647
                                                                var v102 int16 = -32768
                                                                _ = v102
                                                            }
                                                            switch 1 {
                                                            case 32768:
                                                                v103 := GenS9{F0: 32767, F1: 2147483647, F2: 0}
                                                                _ = v103
                                                            case 4294967295:
                                                                if (!(!v6)) {
                                                                    v41 = v10
                                                                    if (v6 || v6) {
                                                                        switch v41 {
                                                                        case "":
                                                                            var v104 int16 = -1
                                                                            _ = v104
                                                                        default:
                                                                            v105 := GenS8{F0: -9223372036854775808, F1: 738}
                                                                            _ = v105
                                                                        }
                                                                    } else {
                                                                        var v106 int16 = 32767
                                                                        switch -892 {
                                                                        case -425:
                                                                            if v6 {
                                                                                var v107 uint64 = 4294967295
                                                                                if true {
                                                                                    v10 = v5
                                                                                }
                                                                                v2 = append(v2, 127)
                                                                                _ = v107
                                                                            } else {
                                                                                if (!((((0.1 + -2.25) * ((3.14159 + 1.5) / (1e10 - 3.14159))) <= -2.25) || (v6 && v6))) {
                                                                                    var v108 uint8 = 122
                                                                                    _ = v108
                                                                                }
                                                                            }
                                                                        case 127:
                                                                            var v109 int16 = v106
                                                                            _ = v109
                                                                        case 1:
                                                                            _ = len(v3)
                                                                        default:
                                                                            v6 = v6
                                                                        }
                                                                        var v110 uint8 = 244
                                                                        _ = v106
                                                                        _ = v110
                                                                    }
                                                                    switch fi20 {
                                                                    case -9223372036854775808:
                                                                        var v111 int16 = -840
                                                                        _ = v111
                                                                    case -514:
                                                                        var v112 uint64 = 2
                                                                        _ = v112
                                                                    default:
                                                                        v1 = float64((0.1 + (0.1 * 0.1)))
                                                                        switch fi20 {
                                                                        case 0:
                                                                            var v113 int8 = 127
                                                                            _ = v113
                                                                        default:
                                                                            fmt.Println(v38, v6)
                                                                            v6 = (0 < 42)
                                                                            v114 := v3[65535]
                                                                            _ = v114
                                                                        }
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v4 = append(v4, 42)
                                                                        }
                                                                        switch -118 {
                                                                        case -433:
                                                                            _ = len(v2)
                                                                        case 709:
                                                                            var v115 uint64 = 2
                                                                            _ = v115
                                                                        }
                                                                    }
                                                                }
                                                                var v116 int64 = 546
                                                                _ = v116
                                                            case 658:
                                                                v4 = append(v4, 2147483648)
                                                            }
                                                        }
                                                    } else {
                                                        if v6 {
                                                            var v117 float64 = v1
                                                            v118 := GenS10{F0: 2147483647}
                                                            _ = v117
                                                            _ = v118
                                                        } else {
                                                            fmt.Println(v16)
                                                            var v119 int64 = -9223372036854775808
                                                            _ = v119
                                                        }
                                                    }
                                                    switch ((v10 + (v10 + v10)) + v10) {
                                                    case "a":
                                                        if true {
                                                            v6 = (!v6)
                                                            var v120 int = v38
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v38 += (v38 % (((v38 / (((^(fi21))) | 1))) | 1))
                                                            default:
                                                                if true {
                                                                    v1 -= 3.14159
                                                                }
                                                            }
                                                            v5 = v10
                                                            if (((v5 <= (v10 + (((v10 + (v10 + v10)) + (("\n\t" + (("" + ((v10 + v41) + v10)) + v10)) + "\n\t")) + v41))) && ("\n\t" > v41)) && (713 < 42)) {
                                                                v121 := GenS7{F0: 65535, F1: 1}
                                                                v122 := v3[2]
                                                                _ = v121
                                                                _ = v122
                                                            } else {
                                                                v123 := GenS11{F0: 0, F1: 2, F2: 717}
                                                                v1 += v1
                                                                v124 := GenS4{F0: 4294967295}
                                                                _ = v123
                                                                _ = v124
                                                            }
                                                            _ = v120
                                                        }
                                                    default:
                                                        v16 += v38
                                                        if false {
                                                            var v125 uint = 189
                                                            _ = v125
                                                        }
                                                        switch v5 {
                                                        case "a":
                                                            if false {
                                                                var v126 uint32 = 1
                                                                var v127 int32 = -297
                                                                _ = v126
                                                                _ = v127
                                                            } else {
                                                                v128 := GenS2{F0: -1, F1: 301}
                                                                if true {
                                                                    var v129 float32 = float32(v1)
                                                                    if v6 {
                                                                        switch v41 {
                                                                        case "":
                                                                            var v130 uint32 = 876
                                                                            _ = v130
                                                                        case "a":
                                                                            v131 := GenS12{F0: 495, F1: -2147483648, F2: 21}
                                                                            _ = v131
                                                                        case "foo bar":
                                                                            var v132 float32 = (v129 * ((((v129 - v129) - 1e10) + 0.0) + ((3.14159 - (((float32(v1) / 0.1) / v129) * (v129 / 1.0))) - v129)))
                                                                            _ = v132
                                                                        }
                                                                    } else {
                                                                        v5 = (v10 + "")
                                                                    }
                                                                    var v133 int8 = 1
                                                                    _ = v129
                                                                    _ = v133
                                                                }
                                                                v134 := GenS2{F0: -32768, F1: 570}
                                                                v135 := GenS10{F0: 1}
                                                                _ = v128
                                                                _ = v134
                                                                _ = v135
                                                            }
                                                        case "foo bar":
                                                            fmt.Println(v38, fi21)
                                                        default:
                                                            var v136 string = v41
                                                            _ = v136
                                                        }
                                                    }
                                                    _ = v100
                                                }
                                                v5 = "a"
                                            }
                                            v10 = v5
                                        }
                                    }
                                } else {
                                    if (!v6) {
                                        var v137 uint16 = 65535
                                        v3[42] = "\n\t"
                                        _ = v137
                                    } else {
                                        if (1 != 2) {
                                            switch v5 {
                                            case "hello":
                                                var v138 uint32 = 42
                                                _ = v138
                                            case "foo bar":
                                                v1 += v1
                                            default:
                                                _ = len(v4)
                                            }
                                            v139 := GenS13{F0: -5}
                                            _ = v139
                                        }
                                        fmt.Println(v41, v16, fi21)
                                    }
                                }
                                _ = v41
                            } else {
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = -2147483648
                                    }
                                    v6 = v6
                                    v140 := GenS3{F0: 1, F1: false}
                                    _ = v140
                                }
                                if v6 {
                                    var v141 uint32 = 599
                                    _ = v141
                                }
                                switch (fi20 ^ v38) {
                                case 32767:
                                    var v142 int16 = -165
                                    _ = v142
                                case 677:
                                    v6 = v6
                                }
                            }
                        }
                        _ = v38
                    }
                    fmt.Println(v16, v10, v6)
                    switch "a" {
                    case "":
                        var v143 string = v5
                        _ = v143
                    case "\n\t":
                        v144 := GenS14{F0: 4294967295, F1: 1, F2: 4294967295}
                        _ = v144
                    case "hello":
                        var v145 uint16 = 300
                        _ = v145
                    default:
                        v10 = "foo bar"
                        v3[97] = (v10 + v10)
                        v10 = v10
                    }
                    _ = len(v2)
                    v6 = v6
                    _ = fi21
                }
                switch fi20 {
                case 0:
                    if true {
                        var v146 int16 = 1000
                        if (-9223372036854775808 != 9223372036854775807) {
                            var v147 string = (v5 + (v5 + ("" + v5)))
                            _ = v147
                        } else {
                            fmt.Println(v1, fi20)
                            for fi148 := 0; fi148 < 9; fi148++ {
                                if v6 {
                                    if (65535 != 823) {
                                        v5 = v5
                                        v146 = v146
                                        var v149 uint8 = 255
                                        v150 := GenS14{F0: 42, F1: -32, F2: 34}
                                        var v151 uint8 = v149
                                        _ = v149
                                        _ = v150
                                        _ = v151
                                    }
                                    if (((false || false) || false) || (true && (!v6))) {
                                        var v152 uint = 42
                                        v5 = v5
                                        _ = v152
                                    }
                                    if v6 {
                                        v1 = v1
                                        if (v6 && v6) {
                                            switch ("foo bar" + "\n\t") {
                                            case "foo bar":
                                                var v153 float64 = (float64((1e10 - (-2.25 / (1e10 + (-2.25 + 1.5))))) / float64(1e10))
                                                _ = v153
                                            case "a":
                                                if (-128 < -36) {
                                                    v1 = 0.0
                                                    fmt.Println(v6, v5, v5)
                                                    fmt.Println(v6, v146)
                                                } else {
                                                    v154 := GenS15{F0: 421}
                                                    if false {
                                                        if v6 {
                                                            v1 = ((1.5 * v1) - (v1 + v1))
                                                            if false {
                                                                if ((true && (!v6)) && ((true || (41 <= 42)) || ((!(((0.0 * (1e10 - 3.14159)) / (3.14159 - ((1.5 / (0.0 * ((((((1e10 / float32(((((3.14159 - (float64(1.5) / float64(float32(0.1)))) - v1) + (v1 - float64(1e10))) - (v1 + float64((-2.25 / (1e10 + float32((v1 * v1))))))))) / 1.0) * 3.14159) + 3.14159) + 3.14159) - -2.25))) + (3.14159 / 1.5)))) > 3.14159)) || true))) {
                                                                    var v155 bool = (!((379 < 829) && false))
                                                                    v146 = (v146 >> ((0) & 63))
                                                                    var v156 bool = ((v146 | v146) != v146)
                                                                    v157 := GenS16{F0: 18446744073709551615, F1: -1}
                                                                    if v156 {
                                                                        var v158 float64 = v1
                                                                        var v159 int8 = 1
                                                                        var v160 float32 = 1e10
                                                                        _ = v158
                                                                        _ = v159
                                                                        _ = v160
                                                                    }
                                                                    _ = v155
                                                                    _ = v156
                                                                    _ = v157
                                                                }
                                                                var v161 float32 = 0.0
                                                                _ = v161
                                                            } else {
                                                                v1 = (v1 - -2.25)
                                                                var v162 int8 = 100
                                                                _ = v162
                                                            }
                                                            v163 := GenS6{F0: 18446744073709551615, F1: v6, F2: -1}
                                                            v6 = (!(!v6))
                                                            _ = v163
                                                        }
                                                    } else {
                                                        fmt.Println(v6)
                                                    }
                                                    _ = v154
                                                }
                                            case "hello":
                                                fmt.Println(v5, fi20, v146)
                                            default:
                                                v3[2] = v5
                                                switch v5 {
                                                case "\n\t":
                                                    if v6 {
                                                        v146 = (v146 % ((-32768) | 1))
                                                    }
                                                }
                                            }
                                        } else {
                                            v164 := GenS11{F0: 1, F1: 42, F2: v146}
                                            fmt.Println(v1)
                                            _ = v164
                                        }
                                        var v165 int16 = (v146 * (v146 >> ((187) & 63)))
                                        _ = v165
                                    }
                                } else {
                                    v3[2] = (((("" + v10) + "foo bar") + v5) + (v5 + v5))
                                    v16 -= (fi20 * fi20)
                                    var v166 string = ""
                                    switch "hello" {
                                    case "a":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "hello":
                                        v167 := GenS17{F0: 18446744073709551615, F1: 42}
                                        _ = v167
                                    case "":
                                        v168 := v3[1]
                                        _ = v168
                                    default:
                                        delete(v3, 65535)
                                    }
                                    _ = v166
                                }
                                var v169 int8 = -1
                                _ = fi148
                                _ = v169
                            }
                            var v170 float32 = 1e10
                            _ = v170
                        }
                        _ = v146
                    } else {
                        var v171 uint32 = 29
                        _ = v171
                    }
                default:
                    var v172 uint8 = 0
                    _ = v172
                }
                if (42 > 0) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi173 := 0; fi173 < 17; fi173++ {
                        switch ("hello" + (v10 + v10)) {
                        case "hello":
                            v1 += float64(float32(((((v1 / 1e10) * (v1 / float64((3.14159 * ((((3.14159 + 1.5) - 0.0) / (float32((1e10 / float64((3.14159 / (-2.25 + 1.5))))) - 1.5)) * 1.5))))) / 3.14159) / v1)))
                        case "":
                            switch 1 {
                            case 2147483647:
                                fmt.Println(v1, v10)
                            case 127:
                                switch (v16 | 727) {
                                case 128:
                                    fmt.Println(fi20, v1)
                                }
                            case 1:
                                if false {
                                    switch (fi173 << ((1) & 63)) {
                                    case 32767:
                                        v174 := GenS18{F0: -1}
                                        _ = v174
                                    default:
                                        if (v6 && v6) {
                                            var v175 uint16 = 0
                                            v1 = (float64((0.0 / 1e10)) - v1)
                                            _ = v175
                                        } else {
                                            if (!(!(2 <= 266))) {
                                                v1 -= 0.1
                                                var v176 uint32 = 1
                                                var v177 uint64 = 18446744073709551615
                                                v178 := GenS17{F0: v177, F1: 307}
                                                _ = v176
                                                _ = v177
                                                _ = v178
                                            } else {
                                                v10 = "a"
                                            }
                                            fmt.Println(v1, v1, v1)
                                        }
                                    }
                                } else {
                                    var v179 uint8 = 128
                                    fmt.Println(v1)
                                    if (!(!v6)) {
                                        switch "hello" {
                                        case "a":
                                            var v180 string = v5
                                            _ = v180
                                        }
                                        var v181 uint64 = 0
                                        v1 += v1
                                        _ = v181
                                    } else {
                                        delete(v3, 2)
                                        v182 := GenS17{F0: 18446744073709551615, F1: 42}
                                        _ = v182
                                    }
                                    _ = v179
                                }
                            }
                        case "a":
                            var v183 string = ("foo bar" + ("foo bar" + v5))
                            _ = v183
                        }
                        var v184 int = fi20
                        _ = fi173
                        _ = v184
                    }
                }
                _ = fi20
            }
        }
        _ = v16
    } else {
        var v185 uint8 = 127
        switch (v10 + v5) {
        case "\n\t":
            fmt.Println(v10)
        case "hello":
            v185 -= v185
        }
        _ = v185
    }
    v186 := GenS19{F0: 1, F1: 2}
    fmt.Println(v6)
    switch 2147483647 {
    case 32768:
        for fi187 := 0; fi187 < 17; fi187++ {
            for fi188 := 0; fi188 < 19; fi188++ {
                switch v10 {
                case "hello":
                    v189 := GenS20{F0: 18446744073709551615, F1: 65535}
                    _ = v189
                case "foo bar":
                    var v190 int8 = 0
                    _ = v190
                case "":
                    if ((18446744073709551615 != 42) || (((3.14159 * 0.1) - 0.1) != 0.1)) {
                        v191 := v3[2147483647]
                        var v192 float32 = -2.25
                        v10 = (v191 + v5)
                        _ = v191
                        _ = v192
                    }
                    v5 = v10
                    var v193 int32 = -366
                    v194 := GenS11{F0: -1, F1: 4294967295, F2: 42}
                    if (0.1 > (((float64(3.14159) + -2.25) * v1) * float64((((0.0 + 0.1) * 1.5) / (-2.25 + 0.1))))) {
                        var v195 int16 = -84
                        v195 += (v195 + 1000)
                        _ = v195
                    }
                    _ = v193
                    _ = v194
                }
                v196 := GenS2{F0: -1, F1: 896}
                _ = fi188
                _ = v196
            }
            _ = fi187
        }
    case -193:
        var v197 int8 = 73
        _ = v197
    }
    if v6 {
        if v6 {
            v198 := GenS1{F0: -2147483648, F1: 11}
            fmt.Println(v1, v6)
            v6 = false
            _ = v198
        } else {
            v10 = ((v5 + v10) + (v10 + v10))
            v1 += float64(((1e10 / (-2.25 + 1e10)) + ((1e10 * (-2.25 / 1e10)) - -2.25)))
        }
    } else {
        v5 = v10
        switch 4294967295 {
        case 2147483647:
            var v199 uint = 625
            _ = v199
        }
        v200 := GenS9{F0: 2, F1: -108, F2: 1}
        _ = v200
    }
    v5 = "\n\t"
    for fi201 := 0; fi201 < 19; fi201++ {
        if (128 != 255) {
            v1 = v1
        } else {
            v202 := GenS18{F0: -2147483648}
            if v6 {
                v203 := GenS21{F0: 0, F1: 0}
                fmt.Println(v10, v10)
                _ = v203
            }
            _ = v202
        }
        _ = fi201
    }
    var v204 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v204:", v204)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v186
    _ = v2
    _ = v204
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
