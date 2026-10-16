package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint8; F2 string }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 bool; F1 string }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int32; F1 string; F2 uint }
type GenS14 struct { F0 uint; F1 bool; F2 string }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int; F1 uint }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 uint16; F1 uint8; F2 uint }
func genFn0() uint64 {
    switch "a" {
    case "\n\t":
        var v7 bool = (4294967295 == 815)
        _ = v7
    }
    var v8 uint = 65535
    switch -2147483648 {
    case 32767:
        fmt.Println(v8)
    case 128:
        v9 := []int{0, 0, 964, 0}
        _ = v9
    case 1:
        for fi10 := 0; fi10 < 8; fi10++ {
            switch fi10 {
            case 1:
                var v11 int64 = 0
                _ = v11
            case -9223372036854775808:
                for fi12 := 0; fi12 < 17; fi12++ {
                    v13 := []int{-9223372036854775808, (-(fi10)), (fi12 * (fi10 - (fi10 - (fi12 / ((fi10) | 1)))))}
                    _ = len(v13)
                    switch -762 {
                    case 4294967295:
                        v8 = (v8 | (v8 % ((v8) | 1)))
                    default:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    }
                    _ = fi12
                    _ = v13
                }
            }
            _ = fi10
        }
    }
    switch ((("a" + (("" + "") + "")) + "hello") + "a") {
    case "hello":
        v8 -= (v8 ^ 65535)
    }
    _ = v8
    return 4294967295
}
func genFn1(p0 uint8) uint8 {
    switch "hello" {
    case "":
        var v14 int8 = -1
        _ = v14
    default:
        var v15 float32 = (0.0 + ((3.14159 - 0.1) + 1e10))
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        p0 = p0
        fmt.Println(p0)
        _ = fi16
    }
    p0 += p0
    p0 = p0
    _ = p0
    return (p0 >> ((18446744073709551615) & 63))
}

func main() {
    var v1 int8 = -123
    var v2 uint64 = 718
    var v3 int64 = 42
    var v4 uint = 0
    var v5 int8 = 100
    var v6 string = "hello"
    var v17 int8 = v5
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS1{F0: 255, F1: (^(v17))}
        _ = fi18
        _ = v19
    }
    v20 := []int{4294967295, 0, -9223372036854775808, 4294967295}
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 20; fi22++ {
            fmt.Println(fi21)
            if len(v20) > 0 {
                v20[0] = (fi22 / (((fi21 * 9223372036854775807)) | 1))
            }
            v1 = (v17 ^ (v5 >> (((v4 ^ 0)) & 63)))
            fmt.Println(v2, fi21)
            v23 := GenS1{F0: 179, F1: (v5 + v17)}
            _ = fi22
            _ = v23
        }
        var v24 uint16 = 42
        for fi25 := 0; fi25 < 8; fi25++ {
            if false {
                switch v6 {
                case "foo bar":
                    if true {
                        var v26 uint = (v4 >> (((v4 ^ (v4 * 18446744073709551615))) & 63))
                        var v27 uint = 18446744073709551615
                        _ = v26
                        _ = v27
                    } else {
                        _ = len(v20)
                        var v28 uint16 = 756
                        _ = v28
                    }
                default:
                    var v29 int32 = -944
                    _ = v29
                }
                var v30 int64 = (^(v3))
                _ = v30
            }
            _ = fi25
        }
        _ = fi21
        _ = v24
    }
    switch (v6 + ((v6 + "\n\t") + ("hello" + v6))) {
    case "hello":
        _ = len(v20)
    case "foo bar":
        switch "hello" {
        case "a":
            var v31 uint64 = v2
            _ = v31
        }
    default:
        if false {
            _ = len(v20)
        } else {
            var v32 uint16 = 0
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 int8 = (-(v17))
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                var v35 int64 = v3
                _ = fi33
                _ = v34
                _ = v35
            }
            var v36 int8 = (v1 >> ((v4) & 63))
            _ = v32
            _ = v36
        }
    }
    if (uint32(v4) != 2) {
        v1 = v17
        for fi37 := 0; fi37 < 17; fi37++ {
            fmt.Println(fi37)
            v5 += 105
            var v38 bool = (-953 < -1)
            _ = fi37
            _ = v38
        }
    }
    v2 -= (v2 - (v2 | uint64(v4)))
    switch 739 {
    case 128:
        fmt.Println(v2, v3)
    default:
        if false {
            v39 := GenS2{F0: (v3 | -1)}
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        v41 := GenS3{F0: (^(v4))}
        if (!true) {
            v3 -= (v3 >> (((v4 >> (((v4 | 693)) & 63))) & 63))
            switch v6 {
            case "a":
                if true {
                    if (-1 <= 1) {
                        v5 -= 100
                        var v42 int32 = 539
                        if true {
                            if (!true) {
                                v6 = v6
                            } else {
                                if true {
                                    for fi43 := 0; fi43 < 2; fi43++ {
                                        var v44 int16 = 32767
                                        v6 = "hello"
                                        v45 := GenS1{F0: 180, F1: (v5 >> (((v4 - 620)) & 63))}
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                    }
                                } else {
                                    _ = len(v20)
                                }
                                for fi46 := 0; fi46 < 7; fi46++ {
                                    v2 += v2
                                    if (317 != v2) {
                                        fmt.Println(v6, v3, v1)
                                        if true {
                                            if false {
                                                switch (fi40 >> ((2) & 63)) {
                                                case 0:
                                                    v47 := GenS2{F0: (^(v3))}
                                                    _ = v47
                                                }
                                                v4 = (v4 | 255)
                                                var v48 float64 = (1.5 * 1e10)
                                                _ = v48
                                            } else {
                                                switch v6 {
                                                case "foo bar":
                                                    v49 := GenS2{F0: v3}
                                                    _ = v49
                                                case "a":
                                                    if (!(true && false)) {
                                                        if (fi46 < fi40) {
                                                            if (1e10 > (1.5 - ((0.1 + 1e10) / 3.14159))) {
                                                                switch (fi40 % (((fi46 & -2147483648)) | 1)) {
                                                                case 42:
                                                                    v5 += (v17 | v1)
                                                                case -665:
                                                                    v3 = -9223372036854775808
                                                                default:
                                                                    if true {
                                                                        fmt.Println(v1)
                                                                        if len(v20) >= 2 {
                                                                            v20 = v20[1:len(v20)]
                                                                        }
                                                                        var v50 float32 = 3.14159
                                                                        _ = v50
                                                                    }
                                                                    if ((v6 == ((v6 + v6) + "foo bar")) && false) {
                                                                        v17 -= (v1 & v1)
                                                                        if true {
                                                                            if true {
                                                                                fmt.Println(v42, v1)
                                                                                if len(v20) > 0 {
                                                                                    v20[0] = 32768
                                                                                }
                                                                            }
                                                                            if (false && false) {
                                                                                if (((133 >= 255) || (2147483648 <= 42)) && true) {
                                                                                    var v51 string = ((v6 + ((((v6 + v6) + (("\n\t" + ((((((v6 + v6) + "a") + v6) + ((v6 + v6) + "")) + v6) + (v6 + v6))) + (v6 + v6))) + (((v6 + v6) + "") + ((("a" + ((v6 + v6) + "foo bar")) + v6) + "a"))) + v6)) + v6)
                                                                                    if true {
                                                                                        fmt.Println(fi46)
                                                                                    } else {
                                                                                        v52 := GenS4{F0: v4, F1: 446}
                                                                                        if false {
                                                                                            var v53 int64 = v3
                                                                                            _ = v53
                                                                                        }
                                                                                        _ = v52
                                                                                    }
                                                                                    if true {
                                                                                        if true {
                                                                                            var v54 bool = true
                                                                                            v20 = append(v20, (-(fi46)))
                                                                                            _ = v54
                                                                                        }
                                                                                        if true {
                                                                                            v2 -= (v2 / (((v2 & genFn0())) | 1))
                                                                                        }
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            v20 = append(v20, (fi46 % ((fi40) | 1)))
                                                                                        case "hello":
                                                                                            v5 += v5
                                                                                        case "foo bar":
                                                                                            var v55 uint32 = uint32(v4)
                                                                                            _ = v55
                                                                                        default:
                                                                                            var v56 int8 = v1
                                                                                            _ = v56
                                                                                        }
                                                                                    } else {
                                                                                        var v57 int8 = (^(v5))
                                                                                        if len(v20) >= 2 {
                                                                                            v20 = v20[1:len(v20)]
                                                                                        }
                                                                                        _ = v57
                                                                                    }
                                                                                    _ = v51
                                                                                } else {
                                                                                    if true {
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            var v58 int16 = -1
                                                                                            _ = v58
                                                                                        default:
                                                                                            if true {
                                                                                                v59 := GenS1{F0: 20, F1: v1}
                                                                                                v1 = (v5 * (^(v17)))
                                                                                                v1 = v5
                                                                                                _ = v59
                                                                                            }
                                                                                            fmt.Println(v6, v4, v4)
                                                                                        }
                                                                                        v20 = append(v20, (fi46 << (((v4 ^ v4)) & 63)))
                                                                                        v60 := GenS4{F0: (^(v4)), F1: 340}
                                                                                        _ = v60
                                                                                    }
                                                                                    _ = len(v20)
                                                                                }
                                                                                if (!false) {
                                                                                    if true {
                                                                                        var v61 int64 = (-(v3))
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            v17 -= (v1 & v1)
                                                                                        default:
                                                                                            if (!true) {
                                                                                                var v62 uint = 720
                                                                                                switch "" {
                                                                                                case "":
                                                                                                    if (true || (false || (true && true))) {
                                                                                                        var v63 uint32 = 2147483648
                                                                                                        v6 = v6
                                                                                                        _ = v63
                                                                                                    }
                                                                                                default:
                                                                                                    v64 := GenS2{F0: 177}
                                                                                                    _ = v64
                                                                                                }
                                                                                                _ = v62
                                                                                            } else {
                                                                                                if (42 < 823) {
                                                                                                    if true {
                                                                                                        fmt.Println(v2)
                                                                                                    } else {
                                                                                                        v65 := GenS5{F0: v3, F1: 182, F2: v6}
                                                                                                        _ = v65
                                                                                                    }
                                                                                                    if len(v20) >= 2 {
                                                                                                        v20 = v20[1:len(v20)]
                                                                                                    }
                                                                                                }
                                                                                                var v66 int64 = 9223372036854775807
                                                                                                v67 := GenS6{F0: v1, F1: fi46}
                                                                                                _ = v66
                                                                                                _ = v67
                                                                                            }
                                                                                        }
                                                                                        v6 = v6
                                                                                        _ = v61
                                                                                    }
                                                                                    if len(v20) > 0 {
                                                                                        v20[0] = 0
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi46, v5, v42)
                                                                                var v68 int64 = v3
                                                                                _ = v68
                                                                            } else {
                                                                                v69 := GenS7{F0: 628}
                                                                                fmt.Println(v6, v1, v5)
                                                                                _ = v69
                                                                            }
                                                                            if (((true || true) && true) || true) {
                                                                                if ((0.1 < 0.1) || (32767 <= 65535)) {
                                                                                    var v70 int = (-(fi46))
                                                                                    var v71 int64 = v3
                                                                                    _ = v70
                                                                                    _ = v71
                                                                                }
                                                                            } else {
                                                                                switch fi46 {
                                                                                case -9223372036854775808:
                                                                                    v72 := GenS8{F0: (v17 + v17)}
                                                                                    _ = v72
                                                                                case 42:
                                                                                    v2 -= v2
                                                                                default:
                                                                                    v17 += -1
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if (42 >= -11) {
                                                                                if (v4 <= v4) {
                                                                                    fmt.Println(v4, v3, fi40)
                                                                                    if true {
                                                                                        fmt.Println(v1)
                                                                                        fmt.Println(v5)
                                                                                    }
                                                                                    _ = len(v20)
                                                                                }
                                                                                if (211 < 106) {
                                                                                    v5 -= (v17 | 127)
                                                                                    if len(v20) > 0 {
                                                                                        v20[0] = int(v42)
                                                                                    }
                                                                                    _ = len(v20)
                                                                                    switch fi40 {
                                                                                    case -2147483648:
                                                                                        v73 := GenS6{F0: v5, F1: (-(fi46))}
                                                                                        _ = v73
                                                                                    case 531:
                                                                                        fmt.Println(v3, v4, v17)
                                                                                        v74 := GenS9{F0: (v6 != ("foo bar" + ((v6 + "a") + v6)))}
                                                                                        _ = v74
                                                                                    default:
                                                                                        var v75 uint8 = 127
                                                                                        _ = v75
                                                                                    }
                                                                                    var v76 uint32 = 733
                                                                                    _ = v76
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v6, v2)
                                                                            }
                                                                            v77 := GenS10{F0: genFn0(), F1: -159}
                                                                            v6 = "a"
                                                                            switch fi40 {
                                                                            case 785:
                                                                                v78 := GenS11{F0: false, F1: v6}
                                                                                _ = v78
                                                                            case 1:
                                                                                if (2147483648 <= 982) {
                                                                                    v5 = (v17 / ((-1) | 1))
                                                                                }
                                                                            case 235:
                                                                                var v79 uint8 = 127
                                                                                _ = v79
                                                                            default:
                                                                                if len(v20) >= 2 {
                                                                                    v20 = v20[1:len(v20)]
                                                                                }
                                                                            }
                                                                            switch v6 {
                                                                            case "foo bar":
                                                                                if len(v20) > 0 {
                                                                                    v20[0] = fi46
                                                                                }
                                                                            case "\n\t":
                                                                                fmt.Println(v42, v5, v17)
                                                                            }
                                                                            _ = v77
                                                                        }
                                                                        v80 := GenS5{F0: v3, F1: 127, F2: (v6 + ("\n\t" + (v6 + v6)))}
                                                                        _ = v80
                                                                    }
                                                                }
                                                                v4 = (v4 << (((v4 / (((v4 << ((v4) & 63))) | 1))) & 63))
                                                            }
                                                        } else {
                                                            var v81 int = fi46
                                                            if (v3 == (^(v3))) {
                                                                v82 := GenS12{F0: (v3 + v3)}
                                                                _ = v82
                                                            } else {
                                                                var v83 int16 = 32767
                                                                _ = v83
                                                            }
                                                            v42 += (v42 - -2147483648)
                                                            _ = v81
                                                        }
                                                        var v84 float32 = 1e10
                                                        _ = v84
                                                    } else {
                                                        _ = len(v20)
                                                    }
                                                case "\n\t":
                                                    v42 += (v42 / ((v42) | 1))
                                                default:
                                                    var v85 uint8 = 255
                                                    _ = v85
                                                }
                                            }
                                            v6 = "hello"
                                        }
                                        var v86 int8 = v1
                                        if false {
                                            var v87 uint32 = 2
                                            if len(v20) > 0 {
                                                v20[0] = fi40
                                            }
                                            switch (fi40 | (fi40 & (fi46 ^ fi40))) {
                                            case 127:
                                                var v88 float32 = 0.0
                                                _ = v88
                                            case -275:
                                                if len(v20) > 0 {
                                                    v20[0] = (fi40 * (fi46 + int(v42)))
                                                }
                                                if true {
                                                    fmt.Println(v1, v1, v4)
                                                    if (!false) {
                                                        switch (fi40 | fi46) {
                                                        case 1:
                                                            var v89 float64 = (0.0 + -2.25)
                                                            _ = v89
                                                        }
                                                    } else {
                                                        fmt.Println(v6, v2, v4)
                                                        switch fi46 {
                                                        case 9223372036854775807:
                                                            if false {
                                                                v90 := GenS8{F0: v17}
                                                                if len(v20) > 0 {
                                                                    v20[0] = (-(fi40))
                                                                }
                                                                if ((v4 * v4) < 65535) {
                                                                    if len(v20) >= 2 {
                                                                        v20 = v20[1:len(v20)]
                                                                    }
                                                                    v91 := GenS1{F0: 10, F1: v17}
                                                                    _ = v91
                                                                }
                                                                _ = v90
                                                            }
                                                        case -1:
                                                            v92 := GenS5{F0: v3, F1: uint8(v4), F2: v6}
                                                            _ = v92
                                                        default:
                                                            v42 += int32(fi46)
                                                        }
                                                        v93 := GenS3{F0: (v4 / ((v4) | 1))}
                                                        _ = v93
                                                    }
                                                    var v94 int8 = (v5 >> ((255) & 63))
                                                    v3 = 0
                                                    fmt.Println(v3, v86)
                                                    v95 := GenS13{F0: int32(v3), F1: v6, F2: v4}
                                                    var v96 bool = true
                                                    _ = v94
                                                    _ = v95
                                                    _ = v96
                                                } else {
                                                    var v97 int16 = -1
                                                    v1 -= (v1 - v5)
                                                    v6 = "foo bar"
                                                    _ = v97
                                                }
                                                v17 = -128
                                            case -494:
                                                if len(v20) > 0 {
                                                    v20[0] = fi46
                                                }
                                            default:
                                                var v98 float64 = 0.1
                                                var v99 float64 = (1.5 - 0.0)
                                                _ = v98
                                                _ = v99
                                            }
                                            _ = v87
                                        } else {
                                            v100 := GenS14{F0: v4, F1: false, F2: "foo bar"}
                                            _ = v100
                                        }
                                        v101 := GenS8{F0: -4}
                                        var v102 uint32 = 2
                                        var v103 uint32 = v102
                                        if false {
                                            var v104 bool = (true && (!true))
                                            v105 := GenS9{F0: v104}
                                            _ = v104
                                            _ = v105
                                        } else {
                                            fmt.Println(fi46)
                                        }
                                        _ = v101
                                        _ = v102
                                        _ = v103
                                        _ = v86
                                    } else {
                                        if (false || false) {
                                            v106 := GenS15{F0: 255}
                                            _ = v106
                                        }
                                        var v107 uint64 = v2
                                        _ = v107
                                    }
                                    var v108 bool = ((!(true && (0 > 0))) && true)
                                    _ = fi46
                                    _ = v108
                                }
                            }
                            var v109 uint16 = 2
                            _ = v109
                        } else {
                            if (3.14159 == 0.0) {
                                var v110 uint16 = 0
                                v1 -= v17
                                v2 = 0
                                _ = v110
                            } else {
                                v111 := GenS2{F0: v3}
                                var v112 int64 = v3
                                v5 += (v1 | v1)
                                _ = v111
                                _ = v112
                            }
                            var v113 bool = true
                            _ = v113
                        }
                        _ = v42
                    }
                    var v114 uint = (v4 % ((42) | 1))
                    v4 = (v114 | 65535)
                    _ = v114
                } else {
                    v115 := GenS11{F0: (739 >= 1), F1: v6}
                    var v116 uint8 = genFn1(42)
                    v116 = (v116 << (((v4 | uint(v116))) & 63))
                    _ = v115
                    _ = v116
                }
            case "":
                for fi117 := 0; fi117 < 19; fi117++ {
                    v4 = v4
                    _ = fi117
                }
            }
        } else {
            v4 += uint(v2)
            v118 := GenS1{F0: 1, F1: (^(v17))}
            if ((-(v3)) >= -308) {
                var v119 int16 = 523
                _ = v119
            }
            _ = v118
        }
        _ = fi40
        _ = v41
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v120 := GenS15{F0: 1}
    for fi121 := 0; fi121 < 6; fi121++ {
        if false {
            v122 := GenS15{F0: 42}
            _ = v122
        }
        v123 := GenS16{F0: fi121, F1: (v4 + v4)}
        if true {
            v4 += (v4 - v4)
        } else {
            switch fi121 {
            case -975:
                var v124 uint8 = 141
                _ = v124
            default:
                v2 = 1
            }
            v3 = (^(v3))
        }
        _ = fi121
        _ = v123
    }
    v125 := GenS17{F0: (3 <= genFn1(128))}
    var v126 uint = v4
    switch 0 {
    case 1:
        var v127 uint = v4
        _ = v127
    default:
        v128 := GenS16{F0: -9223372036854775808, F1: (^(v126))}
        _ = v128
    }
    v1 = v5
    for fi129 := 0; fi129 < 12; fi129++ {
        v130 := GenS15{F0: 255}
        v131 := GenS18{F0: 255, F1: 141, F2: 255}
        _ = fi129
        _ = v130
        _ = v131
    }
    v4 += (^(v126))
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v120
    _ = v125
    _ = v126
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
