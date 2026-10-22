package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 string; F2 uint8 }
type GenS4 struct { F0 uint8; F1 string; F2 uint64 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint32; F1 int8; F2 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint8; F1 string }
type GenS14 struct { F0 string; F1 int; F2 int }
type GenS15 struct { F0 uint8; F1 int32 }
type GenS16 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS17 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int16; F1 int32 }
type GenS21 struct { F0 bool; F1 int; F2 string }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS24 struct { F0 bool; F1 string }
type GenS25 struct { F0 uint }
type GenS26 struct { F0 uint8 }
type GenS27 struct { F0 int8 }

func main() {
    var v1 uint = 654
    v2 := make(map[string]bool)
    var v3 uint64 = 1
    var v4 uint32 = 164
    var v5 bool = (!true)
    v6 := make(map[int]string)
    if v5 {
        fmt.Println(v4, v5)
        for fi7 := 0; fi7 < 12; fi7++ {
            v3 = 4294967295
            _ = fi7
        }
    }
    switch -9223372036854775808 {
    case 2147483647:
        if v5 {
            var v8 float64 = (0.0 * 3.14159)
            v9 := GenS1{F0: uint64(v4), F1: -9223372036854775808, F2: v4}
            for fi10 := 0; fi10 < 1; fi10++ {
                if v5 {
                    v5 = v5
                    delete(v2, ("hello" + ""))
                    v11 := GenS1{F0: v3, F1: 397, F2: 760}
                    _ = v11
                } else {
                    if ((!(!v5)) || v5) {
                        v12 := []int{-638}
                        _ = v12
                    }
                }
                switch "foo bar" {
                case "":
                    var v13 uint16 = 431
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v4 += v4
                        var v15 uint8 = 127
                        _ = fi14
                        _ = v15
                    }
                }
                _ = fi10
            }
            _ = v8
            _ = v9
        } else {
            v4 -= (v4 - 1)
            v16 := GenS2{F0: -75}
            _ = v16
        }
    case 128:
        if v5 {
            v17 := GenS1{F0: v3, F1: 42, F2: v4}
            _ = v17
        }
    case 4294967295:
        for fi18 := 0; fi18 < 5; fi18++ {
            delete(v6, (fi18 ^ (-(fi18))))
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int8 = -87
                _ = fi19
                _ = v20
            }
            v21 := GenS2{F0: -128}
            _ = fi18
            _ = v21
        }
    default:
        v22 := GenS1{F0: v3, F1: -1, F2: 56}
        _ = v22
    }
    switch -9223372036854775808 {
    case 1:
        if (((0.1 + (0.1 / ((0.1 * (1e10 / 0.1)) / (1.5 / ((3.14159 / float64(3.14159)) - 0.0))))) + 0.1) != (1e10 - 0.0)) {
            for fi23 := 0; fi23 < 10; fi23++ {
                for fi24 := 0; fi24 < 2; fi24++ {
                    switch "" {
                    case "a":
                        if (((v5 || v5) && true) || v5) {
                            fmt.Println(v1, v4)
                            var v25 string = "\n\t"
                            var v26 int32 = 42
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = fi24
                }
                switch 127 {
                case 0:
                    if false {
                        fmt.Println(fi23)
                        v27 := []int{}
                        v28 := GenS1{F0: (v3 & v3), F1: 42, F2: 438}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 uint16 = 2
                        _ = v29
                    }
                default:
                    v30 := GenS3{F0: -1, F1: "foo bar", F2: 128}
                    _ = v30
                }
                v3 = 42
                _ = fi23
            }
            v1 = v1
        }
    case 4294967295:
        _ = len(v2)
    case 305:
        v31 := []int{}
        _ = v31
    }
    switch "hello" {
    case "a":
        v32 := GenS1{F0: (v3 >> ((v1) & 63)), F1: -65, F2: v4}
        _ = v32
    }
    var v33 int16 = -32768
    v1 = v1
    switch 128 {
    case 4294967295:
        v34 := GenS4{F0: 128, F1: "a", F2: (^(v3))}
        _ = v34
    }
    fmt.Println(v5, v3)
    for fi35 := 0; fi35 < 7; fi35++ {
        if false {
            if (!v5) {
                var v36 int = fi35
                _ = v36
            }
        } else {
            v6[(fi35 % ((fi35) | 1))] = "foo bar"
            var v37 uint8 = 42
            _ = v37
        }
        _ = fi35
    }
    v38 := []int{42, int(v1), -463, 0}
    v39 := GenS3{F0: -1, F1: "foo bar", F2: 2}
    var v40 uint = (v1 * 18446744073709551615)
    if v5 {
        if true {
            switch "a" {
            case "\n\t":
                fmt.Println(v40, v40, v5)
            case "":
                for fi41 := 0; fi41 < 19; fi41++ {
                    var v42 int8 = 100
                    v38 = append(v38, 0)
                    _ = fi41
                    _ = v42
                }
            default:
                v38 = append(v38, -1)
            }
            v43 := GenS2{F0: 127}
            var v44 int8 = 127
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 2; fi45++ {
            var v46 int = 127
            for fi47 := 0; fi47 < 12; fi47++ {
                if (-2147483648 != int32(v40)) {
                    switch "\n\t" {
                    case "\n\t":
                        if (!(0.0 != -2.25)) {
                            v48 := GenS5{F0: v1, F1: 255}
                            var v49 int32 = 1
                            _ = v48
                            _ = v49
                        }
                    case "":
                        v50 := GenS6{F0: -1}
                        _ = v50
                    default:
                        v51 := GenS3{F0: -2147483648, F1: "\n\t", F2: 127}
                        _ = v51
                    }
                } else {
                    var v52 int8 = 127
                    if v5 {
                        v5 = v5
                        var v53 int8 = int8(v40)
                        if (!(v3 < 2)) {
                            v54 := GenS4{F0: 142, F1: (("hello" + ("" + "foo bar")) + "a"), F2: (^(v3))}
                            if (false || v5) {
                                v46 = (^(v46))
                                fmt.Println(fi45, v5)
                                fmt.Println(v40, v33, v40)
                            } else {
                                if (65535 == 1) {
                                    v38 = append(v38, (fi47 - fi45))
                                    v55 := v6[585]
                                    v56 := GenS7{F0: v4, F1: (v46 % ((32768) | 1))}
                                    if (float64(1.5) >= 1e10) {
                                        if v5 {
                                            if ((1.5 + (((0.0 + -2.25) / -2.25) / 1e10)) == ((1e10 * (3.14159 + -2.25)) * (0.0 / ((3.14159 + (float32(1.5) - float32(-2.25))) + (-2.25 - -2.25))))) {
                                                var v57 float32 = (1e10 - 0.0)
                                                if v5 {
                                                    if false {
                                                        switch v55 {
                                                        case "a":
                                                            var v58 string = v55
                                                            _ = v58
                                                        }
                                                    }
                                                    var v59 int64 = -1
                                                    v60 := GenS8{F0: v4, F1: (v52 % (((v53 * v52)) | 1)), F2: v5}
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    if (!((1e10 / 1e10) != v57)) {
                                                        switch (fi47 >> ((65535) & 63)) {
                                                        case 1:
                                                            var v61 int64 = -9223372036854775808
                                                            _ = v61
                                                        default:
                                                            switch 32767 {
                                                            case 32768:
                                                                v62 := GenS9{F0: uint(v4)}
                                                                _ = v62
                                                            case 127:
                                                                fmt.Println(v55, v3)
                                                            case -958:
                                                                switch ("a" + "") {
                                                                case "\n\t":
                                                                    var v63 bool = (!v5)
                                                                    _ = v63
                                                                case "foo bar":
                                                                    switch (-(v46)) {
                                                                    case 127:
                                                                        var v64 uint8 = 2
                                                                        _ = v64
                                                                    case 446:
                                                                        var v65 int64 = -83
                                                                        _ = v65
                                                                    default:
                                                                        if true {
                                                                            if (985 <= v46) {
                                                                                v66 := GenS10{F0: (v4 | v4)}
                                                                                if (177 >= 65535) {
                                                                                    var v67 uint32 = (^(v4))
                                                                                    fmt.Println(v52, v55, v55)
                                                                                    v52 = v53
                                                                                    _ = v67
                                                                                } else {
                                                                                    switch 32767 {
                                                                                    case -706:
                                                                                        switch (fi45 & fi45) {
                                                                                        case 0:
                                                                                            v68 := GenS4{F0: 0, F1: "a", F2: v3}
                                                                                            _ = v68
                                                                                        case 795:
                                                                                            fmt.Println(v46, v46, fi45)
                                                                                        default:
                                                                                            var v69 uint = 0
                                                                                            _ = v69
                                                                                        }
                                                                                    default:
                                                                                        v46 = v46
                                                                                    }
                                                                                    fmt.Println(v1)
                                                                                    v70 := GenS5{F0: (v1 / (((v40 % ((456) | 1))) | 1)), F1: 1}
                                                                                    _ = v70
                                                                                }
                                                                                _ = v66
                                                                            }
                                                                            v40 = (v1 << ((v1) & 63))
                                                                        }
                                                                    }
                                                                case "":
                                                                    var v71 int16 = v33
                                                                    _ = v71
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v55, v53, v3)
                                                        fmt.Println(v5)
                                                    }
                                                    if v5 {
                                                        if (-2147483648 > 42) {
                                                            v72 := GenS7{F0: v4, F1: (-(fi47))}
                                                            v73 := GenS8{F0: (v4 + v4), F1: v52, F2: (v40 >= (v40 ^ v40))}
                                                            _ = v72
                                                            _ = v73
                                                        } else {
                                                            switch ((v55 + "\n\t") + v55) {
                                                            case "a":
                                                                if (true || true) {
                                                                    v52 += 117
                                                                    if false {
                                                                        var v74 uint32 = (^(v4))
                                                                        var v75 int64 = -9223372036854775808
                                                                        v76 := GenS7{F0: v4, F1: int(v40)}
                                                                        _ = v74
                                                                        _ = v75
                                                                        _ = v76
                                                                    } else {
                                                                        v33 = (^(v33))
                                                                        switch v55 {
                                                                        case "":
                                                                            v1 = (v1 + (v40 * v1))
                                                                        }
                                                                        v77 := GenS1{F0: 710, F1: 9223372036854775807, F2: (v4 << ((39) & 63))}
                                                                        _ = v77
                                                                    }
                                                                    var v78 uint8 = 64
                                                                    _ = v78
                                                                }
                                                            case "\n\t":
                                                                var v79 int16 = (v33 / ((-1) | 1))
                                                                _ = v79
                                                            }
                                                        }
                                                        if (v5 && (v5 || (3.14159 >= 1.5))) {
                                                            v80 := GenS7{F0: v4, F1: -426}
                                                            if (!v5) {
                                                                switch fi47 {
                                                                case 42:
                                                                    v46 += (fi45 * (-(fi47)))
                                                                }
                                                            } else {
                                                                v81 := GenS9{F0: v40}
                                                                v46 = -629
                                                                _ = v81
                                                            }
                                                            _ = v80
                                                        } else {
                                                            switch (((v55 + "hello") + ("a" + v55)) + "a") {
                                                            case "\n\t":
                                                                fmt.Println(fi47, v40, fi47)
                                                            case "hello":
                                                                if len(v38) > 0 {
                                                                    v38[0] = -779
                                                                }
                                                            case "":
                                                                fmt.Println(v57)
                                                            }
                                                            v52 -= (-(v53))
                                                            switch fi45 {
                                                            case 9223372036854775807:
                                                                v6[363] = v55
                                                            case -720:
                                                                fmt.Println(v46, v55, v40)
                                                            case 454:
                                                                var v82 float64 = 0.0
                                                                _ = v82
                                                            default:
                                                                var v83 uint16 = 42
                                                                _ = v83
                                                            }
                                                            var v84 string = "hello"
                                                            _ = v84
                                                        }
                                                    }
                                                }
                                                _ = v57
                                            }
                                            switch v46 {
                                            case 9223372036854775807:
                                                switch fi45 {
                                                case 2147483647:
                                                    var v85 uint32 = 65535
                                                    _ = v85
                                                case 1:
                                                    v3 -= 0
                                                }
                                            case 42:
                                                v46 = (v46 % ((fi45) | 1))
                                            }
                                        }
                                    }
                                    if (v55 > ("" + v55)) {
                                        v86 := GenS11{F0: (v40 * v1), F1: fi47}
                                        var v87 uint = (v1 / (((v1 & v1)) | 1))
                                        if (2147483647 <= 2147483647) {
                                            v88 := GenS12{F0: 42}
                                            fmt.Println(v33)
                                            _ = len(v38)
                                            _ = v88
                                        } else {
                                            var v89 int64 = 9223372036854775807
                                            fmt.Println(v55)
                                            _ = v89
                                        }
                                        var v90 int32 = 651
                                        _ = v86
                                        _ = v87
                                        _ = v90
                                    } else {
                                        if false {
                                            v91 := GenS13{F0: 40, F1: ""}
                                            v46 -= int(v33)
                                            fmt.Println(v4, v46)
                                            _ = v91
                                        } else {
                                            if (255 == 128) {
                                                v92 := GenS14{F0: v55, F1: (-(fi47)), F2: (fi45 * (fi47 >> ((v1) & 63)))}
                                                var v93 float32 = float32(1e10)
                                                var v94 uint64 = 454
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            } else {
                                                if v5 {
                                                    if v5 {
                                                        switch v55 {
                                                        case "hello":
                                                            var v95 int64 = 983
                                                            switch fi47 {
                                                            case 4294967295:
                                                                v53 = v53
                                                                v5 = (!(-1 != -426))
                                                                v40 -= v40
                                                            case 42:
                                                                var v96 uint8 = 1
                                                                _ = v96
                                                            default:
                                                                if true {
                                                                    v38 = append(v38, (^(v46)))
                                                                    if false {
                                                                        v97 := GenS10{F0: v4}
                                                                        _ = v97
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v95, v52)
                                                            var v98 float64 = float64(-2.25)
                                                            if true {
                                                                var v99 float32 = 3.14159
                                                                _ = v99
                                                            } else {
                                                                if v5 {
                                                                    v38 = append(v38, 128)
                                                                    var v100 uint = (v1 + 981)
                                                                    v38 = append(v38, fi47)
                                                                    fmt.Println(fi45)
                                                                    _ = v100
                                                                } else {
                                                                    fmt.Println(v4)
                                                                }
                                                                if (!(2 <= 42)) {
                                                                    switch v46 {
                                                                    case -198:
                                                                        var v101 string = (("hello" + (("" + ("a" + "\n\t")) + v55)) + "foo bar")
                                                                        v98 -= 3.14159
                                                                        _ = v101
                                                                    case -9223372036854775808:
                                                                        v102 := GenS1{F0: (^(v3)), F1: (v95 / (((v95 % ((451) | 1))) | 1)), F2: uint32(v40)}
                                                                        _ = v102
                                                                    default:
                                                                        delete(v2, v55)
                                                                    }
                                                                    _ = len(v38)
                                                                }
                                                            }
                                                            v103 := GenS7{F0: (v4 * v4), F1: (v46 / ((fi47) | 1))}
                                                            v33 = v33
                                                            _ = v103
                                                            _ = v95
                                                            _ = v98
                                                        case "foo bar":
                                                            if (v5 && (42 >= uint8(v4))) {
                                                                v3 -= 924
                                                                v40 += (v1 | (v40 ^ v1))
                                                            }
                                                            v4 -= uint32(v40)
                                                            fmt.Println(v40)
                                                        case "\n\t":
                                                            v2["\n\t"] = (-2.25 >= float32(float64(0.0)))
                                                        default:
                                                            var v104 uint32 = (v4 - v4)
                                                            _ = v104
                                                        }
                                                    }
                                                }
                                                switch v55 {
                                                case "":
                                                    v105 := GenS12{F0: v53}
                                                    _ = v105
                                                case "foo bar":
                                                    fmt.Println(v3, v52)
                                                }
                                                switch (fi45 / ((v46) | 1)) {
                                                case -2147483648:
                                                    v106 := GenS1{F0: 42, F1: 1, F2: 835}
                                                    _ = v106
                                                case 1:
                                                    v107 := GenS11{F0: 1, F1: 9223372036854775807}
                                                    _ = v107
                                                }
                                            }
                                            switch v55 {
                                            case "a":
                                                fmt.Println(v33)
                                            case "hello":
                                                fmt.Println(v53)
                                            }
                                            v108 := GenS15{F0: 42, F1: -1}
                                            _ = v108
                                        }
                                        fmt.Println(v3, v52)
                                        if v5 {
                                            v33 += v33
                                            if ((!false) && false) {
                                                v109 := GenS16{F0: 2, F1: v5, F2: 1}
                                                _ = v109
                                            } else {
                                                v4 = (v4 >> ((v1) & 63))
                                            }
                                            v52 = int8(v1)
                                        } else {
                                            if (!((v33 - int16(v1)) == (v33 & -32768))) {
                                                var v110 uint16 = 32767
                                                _ = len(v38)
                                                fmt.Println(v3)
                                                _ = v110
                                            } else {
                                                var v111 uint64 = v3
                                                v112 := GenS17{F0: (v33 / ((1000) | 1)), F1: 209, F2: v3}
                                                if v5 {
                                                    v33 -= v33
                                                    if v5 {
                                                        if v5 {
                                                            if (1.5 >= ((-2.25 + 0.1) / 1.0)) {
                                                                if (!(false && (false && v5))) {
                                                                    v33 -= 42
                                                                    if (int16(v1) != int16(v4)) {
                                                                        var v113 uint16 = 0
                                                                        v114 := GenS8{F0: v4, F1: (v53 << ((v1) & 63)), F2: ((v53 * v53) <= (v53 | (-(v52))))}
                                                                        var v115 int64 = -633
                                                                        var v116 int32 = -436
                                                                        if v5 {
                                                                            var v117 int8 = -128
                                                                            _ = v117
                                                                        }
                                                                        _ = v113
                                                                        _ = v114
                                                                        _ = v115
                                                                        _ = v116
                                                                    } else {
                                                                        var v118 int = fi47
                                                                        v55 = (((((v55 + v55) + v55) + "hello") + "") + "\n\t")
                                                                        v119 := GenS7{F0: 2, F1: 128}
                                                                        var v120 int = 42
                                                                        fmt.Println(v55, v3)
                                                                        fmt.Println(v46, fi45, v5)
                                                                        _ = v118
                                                                        _ = v119
                                                                        _ = v120
                                                                    }
                                                                    v121 := GenS15{F0: 128, F1: -2147483648}
                                                                    if true {
                                                                        v4 = v4
                                                                        v2[(v55 + "hello")] = v5
                                                                        var v122 int8 = -1
                                                                        _ = v122
                                                                    }
                                                                    _ = v121
                                                                }
                                                                v123 := GenS17{F0: -972, F1: 2, F2: v3}
                                                                _ = v123
                                                            }
                                                            switch v46 {
                                                            case 0:
                                                                v1 = v1
                                                            case -9223372036854775808:
                                                                switch (-(v46)) {
                                                                case -192:
                                                                    fmt.Println(v111, fi45, fi45)
                                                                default:
                                                                    v53 = (v52 - (v53 | int8(v4)))
                                                                    if ((v4 & (v4 | (v4 << ((v40) & 63)))) > 1) {
                                                                        var v124 int = 32767
                                                                        var v125 int32 = -2147483648
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                }
                                                            case 128:
                                                                if (1 < -2147483648) {
                                                                    v46 += 2147483647
                                                                    fmt.Println(v33)
                                                                    var v126 float64 = 1e10
                                                                    _ = v126
                                                                }
                                                            default:
                                                                v127 := GenS15{F0: 128, F1: 535}
                                                                _ = v127
                                                            }
                                                            v40 = (v40 << (((^(v1))) & 63))
                                                        }
                                                        fmt.Println(v55, v40)
                                                        var v128 uint = v40
                                                        var v129 int8 = v53
                                                        if v5 {
                                                            switch "hello" {
                                                            case "hello":
                                                                if v5 {
                                                                    var v130 uint = 1
                                                                    _ = v130
                                                                } else {
                                                                    switch (fi47 | fi45) {
                                                                    case 128:
                                                                        var v131 string = v55
                                                                        _ = v131
                                                                    case -2147483648:
                                                                        v132 := GenS18{F0: -1}
                                                                        _ = v132
                                                                    case 4294967295:
                                                                        if ("" < (v55 + ("\n\t" + (v55 + v55)))) {
                                                                            delete(v2, "hello")
                                                                        } else {
                                                                            if (uint8(v1) >= 96) {
                                                                                v53 -= (v52 >> (((v40 >> ((721) & 63))) & 63))
                                                                            }
                                                                        }
                                                                    default:
                                                                        v38 = append(v38, fi45)
                                                                        fmt.Println(v1, v4, v33)
                                                                        var v133 int8 = v129
                                                                        _ = v133
                                                                    }
                                                                    var v134 float32 = (1e10 + (0.1 / 1e10))
                                                                    v135 := GenS19{F0: v5}
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                            }
                                                            v1 += (v1 - v128)
                                                            if v5 {
                                                                if true {
                                                                    var v136 uint16 = 1
                                                                    _ = v136
                                                                }
                                                            }
                                                            if len(v38) >= 2 {
                                                                v38 = v38[1:len(v38)]
                                                            }
                                                            v1 = v40
                                                            v33 = v33
                                                        }
                                                        _ = v128
                                                        _ = v129
                                                    }
                                                    var v137 string = v55
                                                    v138 := GenS20{F0: (v33 + 42), F1: 527}
                                                    _ = v137
                                                    _ = v138
                                                } else {
                                                    if true {
                                                        fmt.Println(v5)
                                                        var v139 int64 = 42
                                                        _ = v139
                                                    }
                                                }
                                                _ = v111
                                                _ = v112
                                            }
                                            v140 := GenS1{F0: 42, F1: 1, F2: (v4 ^ v4)}
                                            v141 := GenS10{F0: 394}
                                            _ = v140
                                            _ = v141
                                        }
                                        if (!(v55 <= ((v55 + "hello") + "a"))) {
                                            var v142 int = (v46 & fi45)
                                            v142 -= v142
                                            v143 := GenS9{F0: (v40 >> ((uint(v3)) & 63))}
                                            _ = v142
                                            _ = v143
                                        }
                                        var v144 int16 = -771
                                        _ = v144
                                    }
                                    var v145 uint32 = 387
                                    _ = v145
                                    _ = v55
                                    _ = v56
                                } else {
                                    v6[(^(fi45))] = ("hello" + ("" + "a"))
                                    var v146 uint16 = 1
                                    var v147 float32 = 0.1
                                    _ = len(v2)
                                    _ = v146
                                    _ = v147
                                }
                                var v148 uint = (^(v40))
                                var v149 int = 32768
                                _ = v148
                                _ = v149
                            }
                            _ = v54
                        } else {
                            v150 := GenS19{F0: v5}
                            var v151 int64 = 42
                            _ = v150
                            _ = v151
                        }
                        _ = v53
                    }
                    _ = v52
                }
                _ = len(v38)
                _ = len(v2)
                _ = fi47
            }
            _ = fi45
            _ = v46
        }
        v4 += v4
    } else {
        var v152 bool = v5
        if v152 {
            for fi153 := 0; fi153 < 11; fi153++ {
                for fi154 := 0; fi154 < 5; fi154++ {
                    v155 := GenS21{F0: ((v1 << ((686) & 63)) < (v1 / ((957) | 1))), F1: (fi154 << ((v1) & 63)), F2: ""}
                    v156 := GenS21{F0: ((!(0 == 314)) && (true && v152)), F1: fi153, F2: "foo bar"}
                    v1 = (^(v1))
                    _ = fi154
                    _ = v155
                    _ = v156
                }
                v4 = v4
                _ = fi153
            }
            var v157 int = 43
            if (v1 <= v40) {
                v158 := GenS22{F0: v5}
                v152 = (!true)
                switch int(v40) {
                case -1:
                    var v159 string = "\n\t"
                    _ = v159
                case 0:
                    var v160 uint64 = 2
                    _ = v160
                case -9223372036854775808:
                    if v152 {
                        v157 = 0
                        v40 = v1
                    }
                default:
                    v161 := v6[(v157 * (v157 - (^(v157))))]
                    _ = v161
                }
                _ = v158
            } else {
                for fi162 := 0; fi162 < 5; fi162++ {
                    for fi163 := 0; fi163 < 17; fi163++ {
                        fmt.Println(v4, v33, v40)
                        if true {
                            v164 := v6[(fi163 | v157)]
                            _ = v164
                        }
                        if ((fi162 ^ (fi162 >> (((v1 & 18446744073709551615)) & 63))) <= 1) {
                            _ = len(v38)
                            if true {
                                v5 = (v5 || v5)
                                v165 := GenS4{F0: 43, F1: "hello", F2: (v3 << ((v1) & 63))}
                                var v166 string = "foo bar"
                                _ = v165
                                _ = v166
                            } else {
                                v40 = (^(v40))
                                v4 += (^(v4))
                                var v167 int16 = (v33 * v33)
                                _ = v167
                            }
                        } else {
                            if true {
                                v5 = false
                            }
                        }
                        _ = fi163
                    }
                    switch (fi162 * fi162) {
                    case -2147483648:
                        if false {
                            switch "\n\t" {
                            case "":
                                for fi168 := 0; fi168 < 9; fi168++ {
                                    v152 = ((((((("" + "") + ("a" + (("a" + ("\n\t" + "foo bar")) + ("foo bar" + "\n\t")))) + "") + "") + "") + "a") != "foo bar")
                                    v1 -= v40
                                    _ = fi168
                                }
                            case "foo bar":
                                for fi169 := 0; fi169 < 13; fi169++ {
                                    v170 := GenS23{F0: 0, F1: 2147483648, F2: -128}
                                    if v5 {
                                        switch ("hello" + "") {
                                        case "\n\t":
                                            if v152 {
                                                v171 := GenS24{F0: v5, F1: ""}
                                                if (((v3 - 4294967295) <= (v3 & v3)) || v152) {
                                                    if len(v38) > 0 {
                                                        v38[0] = fi169
                                                    }
                                                    if len(v38) >= 2 {
                                                        v38 = v38[1:len(v38)]
                                                    }
                                                }
                                                _ = v171
                                            }
                                        default:
                                            v172 := GenS11{F0: 139, F1: fi162}
                                            _ = v172
                                        }
                                    } else {
                                        v173 := GenS19{F0: v152}
                                        var v174 uint64 = uint64(v4)
                                        _ = v173
                                        _ = v174
                                    }
                                    var v175 uint64 = (v3 - v3)
                                    v176 := GenS5{F0: 640, F1: 210}
                                    switch (^(fi162)) {
                                    case 9223372036854775807:
                                        var v177 uint32 = 0
                                        _ = v177
                                    case 32768:
                                        var v178 float64 = ((0.1 / float64((1.5 * 0.1))) * float64(3.14159))
                                        _ = v178
                                    }
                                    _ = fi169
                                    _ = v170
                                    _ = v175
                                    _ = v176
                                }
                            }
                            if v5 {
                                var v179 int32 = 2147483647
                                v4 = (v4 + (v4 * v4))
                                _ = v179
                            } else {
                                v180 := GenS25{F0: 2}
                                _ = v180
                            }
                            for fi181 := 0; fi181 < 6; fi181++ {
                                switch ("a" + "") {
                                case "hello":
                                    var v182 uint = 4294967295
                                    _ = v182
                                case "\n\t":
                                    switch (((("a" + "") + "") + "") + "foo bar") {
                                    case "":
                                        v38 = append(v38, fi181)
                                    case "\n\t":
                                        v152 = ((fi162 < v157) && (int8(v33) <= 0))
                                    default:
                                        switch "" {
                                        case "hello":
                                            var v183 int16 = 1
                                            _ = v183
                                        }
                                    }
                                    if v5 {
                                        fmt.Println(fi181)
                                        v33 -= v33
                                        var v184 uint16 = 0
                                        _ = v184
                                    } else {
                                        if v152 {
                                            v185 := GenS2{F0: 1}
                                            _ = v185
                                        } else {
                                            fmt.Println(fi162, v4)
                                        }
                                    }
                                    v186 := GenS26{F0: 1}
                                    _ = v186
                                case "":
                                    var v187 uint = v1
                                    _ = v187
                                }
                                v33 = 0
                                _ = fi181
                            }
                        } else {
                            var v188 int16 = 1000
                            for fi189 := 0; fi189 < 16; fi189++ {
                                v4 -= 993
                                if (v33 < 464) {
                                    var v190 uint32 = (^(v4))
                                    var v191 float32 = 3.14159
                                    _ = v190
                                    _ = v191
                                } else {
                                    var v192 int32 = int32(v4)
                                    _ = v192
                                }
                                _ = fi189
                            }
                            if true {
                                var v193 float32 = ((0.0 / (-2.25 / ((((float32(1e10) - (0.0 + 1.5)) / 3.14159) + 1.5) / 1.0))) * 1.5)
                                _ = len(v38)
                                var v194 string = (("a" + "") + ("hello" + "hello"))
                                _ = v193
                                _ = v194
                            }
                            _ = v188
                        }
                    case 107:
                        v5 = false
                    default:
                        v195 := v2[("" + "")]
                        _ = v195
                    }
                    v38 = append(v38, (v157 << ((1) & 63)))
                    _ = fi162
                }
                v1 += (v1 & (v40 ^ v1))
            }
            _ = v157
        } else {
            _ = len(v38)
        }
        _ = v152
    }
    if (127 < -128) {
        for fi196 := 0; fi196 < 15; fi196++ {
            switch fi196 {
            case 127:
                var v197 uint32 = uint32(v3)
                _ = v197
            case -2147483648:
                v198 := GenS17{F0: v33, F1: 255, F2: (v3 | v3)}
                _ = v198
            case -421:
                var v199 uint32 = v4
                _ = v199
            }
            if v5 {
                if (679 <= 9223372036854775807) {
                    v4 += v4
                    if ((!(-9223372036854775808 >= 1)) || (!v5)) {
                        v200 := GenS27{F0: 1}
                        fmt.Println(v40, fi196, v33)
                        var v201 int8 = 1
                        _ = v200
                        _ = v201
                    }
                    fmt.Println(v4)
                }
                var v202 int16 = 407
                for fi203 := 0; fi203 < 2; fi203++ {
                    var v204 int32 = 1
                    if true {
                        v4 += (v4 >> (((v1 % ((1) | 1))) & 63))
                        _ = len(v38)
                    } else {
                        fmt.Println(v4, v202, v5)
                        _ = len(v38)
                        v204 = (v204 >> (((^(v40))) & 63))
                    }
                    fmt.Println(v33, v1, v202)
                    _ = fi203
                    _ = v204
                }
                _ = v202
            }
            var v205 uint = v40
            _ = fi196
            _ = v205
        }
        var v206 int8 = -128
        _ = v206
    } else {
        fmt.Println(v3, v1)
        for fi207 := 0; fi207 < 6; fi207++ {
            if len(v38) >= 2 {
                v38 = v38[1:len(v38)]
            }
            _ = fi207
        }
    }
    fmt.Println(v4, v4, v40)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
