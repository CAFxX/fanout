package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int32; F1 bool; F2 uint8 }
type GenS3 struct { F0 int16; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint; F1 string; F2 uint }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
type GenS6 struct { F0 int8; F1 int64; F2 uint8 }
type GenS7 struct { F0 int; F1 bool; F2 bool }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 int8; F1 int8 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 bool; F1 uint8; F2 int8 }
type GenS12 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS13 struct { F0 int8; F1 int64 }
type GenS14 struct { F0 int; F1 int64; F2 uint }
type GenS15 struct { F0 bool; F1 uint64 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint64; F1 string }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 uint; F1 uint64; F2 string }
func genFn0() uint {
    var v6 bool = false
    if true {
        v7 := []int{2147483647, -9223372036854775808, -783, 127}
        if v6 {
            var v8 bool = (9223372036854775807 >= 9223372036854775807)
            var v9 int8 = 13
            _ = v8
            _ = v9
        }
        _ = v7
    } else {
        v10 := GenS1{F0: 819, F1: "a"}
        for fi11 := 0; fi11 < 16; fi11++ {
            if (2 == 42) {
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 uint64 = 2
                    v14 := []int{-1}
                    switch "foo bar" {
                    case "a":
                        if v6 {
                            fmt.Println(v13, v6, v13)
                        }
                        if false {
                            switch (fi12 & int(v13)) {
                            case 32768:
                                v14 = append(v14, (^(fi12)))
                            }
                            v13 += (^(v13))
                            v15 := GenS2{F0: 1, F1: v6, F2: 128}
                            v6 = v6
                            if (!v6) {
                                fmt.Println(v13, v13)
                                var v16 uint16 = 255
                                _ = v16
                            } else {
                                v13 = 42
                                fmt.Println(fi12, fi11)
                                var v17 uint8 = 128
                                _ = v17
                            }
                            _ = v15
                        } else {
                            if v6 {
                                var v18 uint8 = 2
                                _ = v18
                            }
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            v14 = append(v14, -842)
                        }
                        v13 = (v13 % ((18446744073709551615) | 1))
                    case "hello":
                        var v19 float64 = (1e10 - 3.14159)
                        _ = v19
                    case "":
                        fmt.Println(v13)
                    }
                    v14 = append(v14, fi11)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            v20 := GenS3{F0: -1, F1: -2147483648, F2: 42}
            v6 = v6
            _ = fi11
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            v6 = false
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 int32 = 938
                if v6 {
                    v24 := []int{}
                    if v6 {
                        var v25 uint32 = 542
                        _ = v25
                    }
                    v6 = (((((0.0 - 1e10) * (0.0 / 1.5)) + 3.14159) != 1.5) || (fi21 == fi21))
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            if v6 {
                for fi26 := 0; fi26 < 2; fi26++ {
                    v27 := []int{}
                    fmt.Println(v6)
                    switch "a" {
                    case "a":
                        v6 = ((226 != 597) && v6)
                    case "hello":
                        switch "a" {
                        case "":
                            v28 := GenS4{F0: 4294967295, F1: "\n\t", F2: 255}
                            _ = v28
                        }
                    case "":
                        if len(v27) > 0 {
                            v27[0] = -1
                        }
                    }
                    _ = fi26
                    _ = v27
                }
                for fi29 := 0; fi29 < 8; fi29++ {
                    var v30 uint = 65535
                    switch 128 {
                    case -1:
                        fmt.Println(fi29, fi21, v30)
                    default:
                        v31 := []int{42, 23, fi21, 0}
                        _ = v31
                    }
                    _ = fi29
                    _ = v30
                }
                v32 := []int{fi21, 64, (fi21 + (fi21 << ((42) & 63)))}
                _ = v32
            } else {
                var v33 int = 4294967295
                var v34 float32 = 1e10
                _ = v33
                _ = v34
            }
            _ = fi21
        }
        _ = v10
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := []int{-838, fi35, 42}
        fmt.Println(fi35, v6)
        _ = fi35
        _ = v36
    }
    if false {
        v6 = v6
        var v37 int8 = 0
        v38 := GenS4{F0: 281, F1: (("hello" + "\n\t") + ("foo bar" + "hello")), F2: 77}
        _ = v37
        _ = v38
    }
    if v6 {
        v6 = v6
        var v39 uint8 = 128
        _ = v39
    } else {
        v40 := GenS3{F0: 32767, F1: 42, F2: 0}
        v41 := []int{-494, -178}
        _ = v40
        _ = v41
    }
    _ = v6
    return 65535
}
func genFn1(p0 int16, p1 uint64, p2 uint8) uint8 {
    p2 = 0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int8 = -85
    var v2 uint32 = 914
    var v3 int64 = 957
    var v4 uint16 = 0
    var v5 bool = false
    v42 := GenS3{F0: 32767, F1: -1, F2: v2}
    var v43 string = "a"
    v44 := []int{32768, 2147483647, 32767}
    for fi45 := 0; fi45 < 19; fi45++ {
        v44 = append(v44, 141)
        v5 = v5
        _ = fi45
    }
    v46 := GenS5{F0: 42, F1: -9223372036854775808, F2: v4}
    if false {
        var v47 float64 = ((((3.14159 * 3.14159) - ((-2.25 - 1.5) - 1.5)) + 1e10) - ((float64(float32(3.14159)) + (1.5 - 1.5)) * 1.5))
        if v5 {
            v4 -= (v4 | 32767)
            var v48 string = v43
            _ = v48
        } else {
            v44 = append(v44, -9223372036854775808)
            var v49 int = 0
            _ = v49
        }
        v1 -= v1
        _ = v47
    }
    switch 4294967295 {
    case -9223372036854775808:
        var v50 float64 = 0.0
        _ = v50
    case 798:
        if true {
            if (false || v5) {
                var v51 uint = 477
                var v52 int32 = 2147483647
                v2 = (v2 | v2)
                _ = v51
                _ = v52
            } else {
                if v5 {
                    v2 = v2
                    v44 = append(v44, -138)
                }
                for fi53 := 0; fi53 < 14; fi53++ {
                    v3 -= -1
                    _ = fi53
                }
                var v54 float32 = (0.0 - ((-2.25 / 1.0) / (0.0 + 3.14159)))
                _ = v54
            }
            for fi55 := 0; fi55 < 14; fi55++ {
                var v56 int8 = 43
                for fi57 := 0; fi57 < 8; fi57++ {
                    v1 -= (v1 - v56)
                    if v5 {
                        v43 = v43
                        if v5 {
                            v58 := GenS6{F0: (v56 / ((-128) | 1)), F1: v3, F2: 17}
                            var v59 uint32 = v2
                            var v60 int16 = 125
                            _ = v58
                            _ = v59
                            _ = v60
                        } else {
                            v4 = v4
                        }
                        var v61 float32 = -2.25
                        _ = v61
                    } else {
                        v2 -= v2
                        if v5 {
                            v4 += (v4 >> ((0) & 63))
                            if len(v44) >= 2 {
                                v44 = v44[1:len(v44)]
                            }
                            _ = len(v44)
                            if (-1 != v3) {
                                if ((0.1 * 0.1) > 3.14159) {
                                    v62 := GenS6{F0: 0, F1: v3, F2: 1}
                                    _ = v62
                                } else {
                                    v63 := GenS1{F0: 255, F1: "hello"}
                                    var v64 int32 = 1
                                    var v65 bool = (!false)
                                    if (127 == fi55) {
                                        switch fi55 {
                                        case 757:
                                            var v66 uint32 = 4294967295
                                            _ = v66
                                        case -1:
                                            v44 = append(v44, fi55)
                                            var v67 uint8 = 42
                                            _ = v67
                                        default:
                                            if ((v64 ^ v64) > int32(fi55)) {
                                                if v65 {
                                                    fmt.Println(v43, v4, v5)
                                                    v4 = 65535
                                                    var v68 int8 = -1
                                                    _ = v68
                                                }
                                                var v69 uint16 = 65535
                                                if ((v5 || false) && true) {
                                                    v64 = (v64 ^ (v64 + 774))
                                                    v70 := GenS5{F0: 726, F1: (v3 << ((11) & 63)), F2: v69}
                                                    _ = v70
                                                }
                                                _ = v69
                                            } else {
                                                v1 -= (v1 >> ((1) & 63))
                                                if (((-(v64)) >= 42) || ((v43 == (v43 + v43)) && false)) {
                                                    v44 = append(v44, (^(fi57)))
                                                    v71 := GenS7{F0: fi57, F1: (v5 || (442 < (^(v2)))), F2: v5}
                                                    if (1 > 0) {
                                                        fmt.Println(fi55, fi55)
                                                        fmt.Println(v64, v64)
                                                        v3 = (v3 << ((65535) & 63))
                                                    } else {
                                                        if (v64 >= 42) {
                                                            if len(v44) >= 2 {
                                                                v44 = v44[1:len(v44)]
                                                            }
                                                            switch fi57 {
                                                            case 42:
                                                                var v72 uint32 = v2
                                                                _ = v72
                                                            case 987:
                                                                if true {
                                                                    if len(v44) >= 2 {
                                                                        v44 = v44[1:len(v44)]
                                                                    }
                                                                    _ = len(v44)
                                                                    if true {
                                                                        v73 := GenS5{F0: v64, F1: v3, F2: (v4 % (((v4 * 15)) | 1))}
                                                                        var v74 int64 = (v3 * (-(v3)))
                                                                        _ = v73
                                                                        _ = v74
                                                                    }
                                                                }
                                                            default:
                                                                var v75 float32 = 0.0
                                                                _ = v75
                                                            }
                                                        } else {
                                                            v76 := GenS2{F0: (-(v64)), F1: (32767 >= -32768), F2: 255}
                                                            var v77 uint16 = 442
                                                            var v78 int32 = v64
                                                            v77 = 65535
                                                            v64 = v64
                                                            var v79 uint8 = 87
                                                            v80 := GenS8{F0: (v43 + "hello"), F1: ((v56 / ((-53) | 1)) == (v1 | 0))}
                                                            var v81 uint8 = 128
                                                            switch v43 {
                                                            case "hello":
                                                                if false {
                                                                    var v82 int32 = v78
                                                                    v83 := GenS1{F0: (v77 + (v77 << ((genFn0()) & 63))), F1: v43}
                                                                    var v84 uint8 = 42
                                                                    _ = v82
                                                                    _ = v83
                                                                    _ = v84
                                                                } else {
                                                                    switch fi55 {
                                                                    case 195:
                                                                        v85 := GenS3{F0: -1, F1: -31, F2: 4294967295}
                                                                        _ = v85
                                                                    default:
                                                                        var v86 int64 = -9223372036854775808
                                                                        _ = v86
                                                                    }
                                                                    _ = len(v44)
                                                                }
                                                            case "a":
                                                                v64 = (v78 / ((int32(v1)) | 1))
                                                            default:
                                                                v44 = append(v44, fi57)
                                                            }
                                                            v3 = (v3 % ((258) | 1))
                                                            _ = v76
                                                            _ = v77
                                                            _ = v78
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        fmt.Println(v56)
                                                        v87 := GenS5{F0: -887, F1: v3, F2: (^(v4))}
                                                        _ = v87
                                                    }
                                                    _ = v71
                                                }
                                            }
                                            v44 = append(v44, 4294967295)
                                            switch v43 {
                                            case "foo bar":
                                                v3 += (v3 >> ((0) & 63))
                                            default:
                                                var v88 uint32 = (v2 + 812)
                                                _ = v88
                                            }
                                        }
                                        var v89 uint32 = (^(v2))
                                        if v5 {
                                            _ = len(v44)
                                            if v5 {
                                                v90 := GenS8{F0: "a", F1: (0.0 < 1e10)}
                                                _ = v90
                                            } else {
                                                fmt.Println(fi55, v65, v2)
                                                v65 = v65
                                                var v91 uint8 = 42
                                                _ = v91
                                            }
                                            v1 -= v56
                                            var v92 int16 = 98
                                            if ((!((v43 + "a") <= v43)) || (!(619 != v64))) {
                                                v3 += int64(v4)
                                                if (v5 && true) {
                                                    v5 = false
                                                } else {
                                                    var v93 uint8 = 42
                                                    v56 += v56
                                                    switch v43 {
                                                    case "\n\t":
                                                        if v5 {
                                                            v92 += 32767
                                                            fmt.Println(v89, v92, v89)
                                                            fmt.Println(v2, fi57)
                                                            fmt.Println(v43, v64, v4)
                                                        } else {
                                                            v4 -= v4
                                                        }
                                                    case "foo bar":
                                                        var v94 uint16 = v4
                                                        _ = v94
                                                    default:
                                                        v1 = (^(v1))
                                                    }
                                                    _ = v93
                                                }
                                                v3 = (v3 >> ((716) & 63))
                                            }
                                            v64 -= (v64 - (v64 ^ (v64 & v64)))
                                            _ = v92
                                        } else {
                                            v43 = v43
                                            fmt.Println(v4)
                                        }
                                        _ = v89
                                    } else {
                                        switch ((v43 + v43) + "a") {
                                        case "\n\t":
                                            v2 -= 2
                                        case "a":
                                            v64 = v64
                                        case "foo bar":
                                            if v65 {
                                                if v5 {
                                                    if false {
                                                        if v5 {
                                                            fmt.Println(v5, v56)
                                                        } else {
                                                            v95 := GenS5{F0: 1, F1: v3, F2: v4}
                                                            if (-2.25 > (-2.25 - 0.1)) {
                                                                var v96 float32 = 0.1
                                                                var v97 uint = 1
                                                                if (v1 > (-(v1))) {
                                                                    if len(v44) >= 2 {
                                                                        v44 = v44[1:len(v44)]
                                                                    }
                                                                    v98 := GenS3{F0: 32767, F1: v64, F2: v2}
                                                                    fmt.Println(v64, v1, v5)
                                                                    if (!true) {
                                                                        var v99 uint16 = 560
                                                                        v97 += (v97 >> ((v97) & 63))
                                                                        fmt.Println(v96)
                                                                        _ = v99
                                                                    } else {
                                                                        if (v65 && (v97 <= v97)) {
                                                                            var v100 int = 244
                                                                            _ = v100
                                                                        } else {
                                                                            fmt.Println(v1, v56)
                                                                            _ = len(v44)
                                                                        }
                                                                        v101 := GenS3{F0: 32767, F1: int32(v2), F2: 2}
                                                                        _ = v101
                                                                    }
                                                                    v56 = (v1 + (v1 % ((v1) | 1)))
                                                                    _ = v98
                                                                }
                                                                _ = v96
                                                                _ = v97
                                                            }
                                                            var v102 float32 = (-2.25 + float32(0.0))
                                                            _ = v102
                                                            _ = v95
                                                        }
                                                        v2 -= v2
                                                        if v65 {
                                                            var v103 uint16 = (v4 | v4)
                                                            _ = v103
                                                        }
                                                        v4 = 790
                                                        if (2 == 55) {
                                                            if ((fi55 - fi57) <= fi55) {
                                                                fmt.Println(v4, v3)
                                                                v3 += v3
                                                                var v104 string = ""
                                                                _ = v104
                                                            }
                                                            _ = len(v44)
                                                        } else {
                                                            var v105 float32 = 0.1
                                                            _ = v105
                                                        }
                                                    } else {
                                                        if (v56 != 5) {
                                                            if (32767 == -1) {
                                                                v106 := GenS4{F0: 240, F1: v43, F2: 4294967295}
                                                                _ = v106
                                                            } else {
                                                                var v107 bool = (428 == 18446744073709551615)
                                                                switch 117 {
                                                                case 127:
                                                                    var v108 int16 = 423
                                                                    _ = v108
                                                                case -577:
                                                                    v109 := GenS1{F0: v4, F1: (((((v43 + v43) + "foo bar") + ("a" + v43)) + v43) + "hello")}
                                                                    _ = v109
                                                                default:
                                                                    v1 -= v1
                                                                }
                                                                _ = v107
                                                            }
                                                            if true {
                                                                v4 += v4
                                                                if len(v44) >= 2 {
                                                                    v44 = v44[1:len(v44)]
                                                                }
                                                            }
                                                        } else {
                                                            var v110 uint = 255
                                                            fmt.Println(v56, fi57, fi55)
                                                            _ = v110
                                                        }
                                                        if len(v44) > 0 {
                                                            v44[0] = fi55
                                                        }
                                                    }
                                                }
                                                v56 += (v56 | (v1 - v56))
                                                fmt.Println(fi57, v56)
                                            } else {
                                                v3 = v3
                                            }
                                        }
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                    }
                                    v56 = 42
                                    if false {
                                        if len(v44) >= 2 {
                                            v44 = v44[1:len(v44)]
                                        }
                                        v56 += -119
                                    }
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                                v5 = v5
                            }
                            v1 -= v56
                        } else {
                            var v111 float64 = 0.1
                            v44 = append(v44, (^(fi57)))
                            _ = v111
                        }
                        if (true && v5) {
                            if false {
                                v112 := GenS7{F0: -9223372036854775808, F1: false, F2: (v5 || ((0 > 1000) || (2 == 0)))}
                                _ = v112
                            }
                        }
                    }
                    _ = fi57
                }
                _ = fi55
                _ = v56
            }
            if true {
                _ = len(v44)
            }
        }
    case 500:
        v113 := GenS9{F0: (v1 / (((v1 * v1)) | 1)), F1: int8(v2)}
        _ = v113
    default:
        var v114 uint = 18446744073709551615
        _ = v114
    }
    _ = len(v44)
    switch 4294967295 {
    case 42:
        switch "\n\t" {
        case "foo bar":
            v43 = ((v43 + v43) + v43)
        }
    case 302:
        v44 = append(v44, -447)
    }
    for fi115 := 0; fi115 < 18; fi115++ {
        v3 += (v3 >> ((255) & 63))
        var v116 int64 = v3
        _ = fi115
        _ = v116
    }
    v2 = 65535
    var v117 string = "foo bar"
    var v118 float64 = ((0.0 / float64(float32(3.14159))) / 1.0)
    v1 = -37
    for fi119 := 0; fi119 < 12; fi119++ {
        var v120 float32 = 1.5
        _ = fi119
        _ = v120
    }
    var v121 float32 = 0.1
    for fi122 := 0; fi122 < 17; fi122++ {
        v123 := GenS3{F0: 32767, F1: -2147483648, F2: (^(v2))}
        switch (v43 + (v117 + v117)) {
        case "a":
            v43 = "foo bar"
        case "foo bar":
            switch fi122 {
            case 0:
                switch v43 {
                case "foo bar":
                    fmt.Println(v5, v5, v2)
                case "a":
                    switch v117 {
                    case "":
                        var v124 int32 = 42
                        _ = v124
                    case "hello":
                        if (false || v5) {
                            v125 := GenS1{F0: 0, F1: ((v43 + "\n\t") + "\n\t")}
                            fmt.Println(v121, v121, v5)
                            _ = v125
                        } else {
                            var v126 uint = 65535
                            v127 := GenS10{F0: 0}
                            if (v118 <= ((v118 / (v118 + 1e10)) + 1e10)) {
                                for fi128 := 0; fi128 < 8; fi128++ {
                                    v4 = v4
                                    switch v43 {
                                    case "a":
                                        v3 = v3
                                    case "foo bar":
                                        v129 := GenS2{F0: 1, F1: v5, F2: 44}
                                        _ = v129
                                    case "hello":
                                        if false {
                                            _ = len(v44)
                                        } else {
                                            var v130 int8 = (-(v1))
                                            var v131 int16 = 42
                                            switch (fi122 - fi128) {
                                            case 128:
                                                fmt.Println(v43, v118, v117)
                                            case -588:
                                                v132 := GenS5{F0: -2147483648, F1: (v3 % (((v3 >> ((v126) & 63))) | 1)), F2: (v4 / ((v4) | 1))}
                                                _ = v132
                                            }
                                            _ = v130
                                            _ = v131
                                        }
                                    default:
                                        if v5 {
                                            var v133 uint8 = uint8(v4)
                                            if (-482 > -2147483648) {
                                                switch 9223372036854775807 {
                                                case 9223372036854775807:
                                                    v121 = -2.25
                                                case -958:
                                                    var v134 int32 = 1
                                                    fmt.Println(fi128, v121)
                                                    v135 := GenS11{F0: false, F1: (^(v133)), F2: v1}
                                                    v44 = append(v44, 128)
                                                    _ = v134
                                                    _ = v135
                                                default:
                                                    v136 := GenS4{F0: v126, F1: v43, F2: v126}
                                                    _ = v136
                                                }
                                                if len(v44) > 0 {
                                                    v44[0] = (fi122 / ((fi128) | 1))
                                                }
                                                var v137 int64 = int64(v4)
                                                _ = v137
                                            }
                                            _ = v133
                                        }
                                    }
                                    var v138 int32 = 42
                                    _ = fi128
                                    _ = v138
                                }
                            }
                            _ = v126
                            _ = v127
                        }
                    }
                }
            case 346:
                v1 += -128
            default:
                v4 = (v4 | v4)
            }
        default:
            if (-2147483648 <= 2147483647) {
                if (420 == 1) {
                    var v139 uint = uint(v2)
                    for fi140 := 0; fi140 < 2; fi140++ {
                        _ = len(v44)
                        v44 = append(v44, 42)
                        var v141 float64 = 3.14159
                        _ = fi140
                        _ = v141
                    }
                    _ = v139
                }
                v44 = append(v44, (fi122 - (fi122 / ((int(v4)) | 1))))
                var v142 int32 = 2147483647
                _ = v142
            } else {
                for fi143 := 0; fi143 < 11; fi143++ {
                    fmt.Println(v118, v121)
                    switch fi143 {
                    case 0:
                        if false {
                            if (0 > (v1 >> ((345) & 63))) {
                                v144 := GenS4{F0: 198, F1: "\n\t", F2: 4294967295}
                                var v145 uint16 = (v4 - v4)
                                _ = v144
                                _ = v145
                            } else {
                                switch -191 {
                                case 128:
                                    _ = len(v44)
                                }
                                v146 := GenS6{F0: 42, F1: 819, F2: uint8(v2)}
                                _ = v146
                            }
                            var v147 uint8 = 128
                            if true {
                                var v148 int16 = 1000
                                var v149 uint32 = v2
                                if v5 {
                                    v150 := GenS12{F0: v5, F1: 1, F2: v147}
                                    var v151 float32 = v121
                                    v148 -= v148
                                    _ = v150
                                    _ = v151
                                } else {
                                    v152 := GenS13{F0: v1, F1: (v3 / (((v3 % (((v3 / ((v3) | 1))) | 1))) | 1))}
                                    v148 = v148
                                    v153 := GenS14{F0: -518, F1: int64(v147), F2: 661}
                                    _ = v152
                                    _ = v153
                                }
                                _ = len(v44)
                                _ = v148
                                _ = v149
                            }
                            var v154 uint32 = v2
                            fmt.Println(v5, v147, v154)
                            _ = v147
                            _ = v154
                        } else {
                            if len(v44) > 0 {
                                v44[0] = int(v1)
                            }
                            switch (^(fi122)) {
                            case -1:
                                if len(v44) >= 2 {
                                    v44 = v44[1:len(v44)]
                                }
                            case -244:
                                v44 = append(v44, fi122)
                            case 0:
                                if true {
                                    v2 = v2
                                }
                            default:
                                var v155 uint16 = 65535
                                _ = v155
                            }
                        }
                    case -1:
                        var v156 int64 = (-(v3))
                        _ = v156
                    case 32768:
                        v43 = v43
                        v157 := GenS15{F0: false, F1: 0}
                        var v158 uint32 = 42
                        v44 = append(v44, 35)
                        _ = v157
                        _ = v158
                    default:
                        switch (fi143 / ((199) | 1)) {
                        case 32767:
                            if ((v117 + v117) > "hello") {
                                if v5 {
                                    if (-9223372036854775808 < 9223372036854775807) {
                                        if (781 == 656) {
                                            v159 := GenS1{F0: (v4 >> ((312) & 63)), F1: "a"}
                                            v1 += v1
                                            var v160 uint8 = genFn1(-644, 531, 127)
                                            v161 := GenS16{F0: 2}
                                            var v162 uint16 = v4
                                            var v163 int32 = -2147483648
                                            var v164 uint8 = 0
                                            _ = v159
                                            _ = v160
                                            _ = v161
                                            _ = v162
                                            _ = v163
                                            _ = v164
                                        } else {
                                            var v165 int16 = -783
                                            _ = v165
                                        }
                                        var v166 uint32 = (^(v2))
                                        var v167 uint32 = v166
                                        if len(v44) > 0 {
                                            v44[0] = fi122
                                        }
                                        v167 = v166
                                        v168 := GenS12{F0: (v5 || (1 != int16(v1))), F1: 0, F2: 1}
                                        v169 := GenS6{F0: (v1 & (-(v1))), F1: -1, F2: genFn1(42, 348, 2)}
                                        _ = v166
                                        _ = v167
                                        _ = v168
                                        _ = v169
                                    }
                                    fmt.Println(v121, fi143)
                                    switch (fi143 / (((fi143 * (fi122 & (^(fi122))))) | 1)) {
                                    case 4294967295:
                                        if len(v44) > 0 {
                                            v44[0] = (^(fi143))
                                        }
                                    }
                                } else {
                                    switch v43 {
                                    case "hello":
                                        fmt.Println(v4)
                                    default:
                                        if v5 {
                                            fmt.Println(v3, v5, v3)
                                            _ = len(v44)
                                            v1 = (-(v1))
                                            switch "foo bar" {
                                            case "hello":
                                                v170 := GenS3{F0: -859, F1: -2147483648, F2: (v2 << ((65535) & 63))}
                                                _ = v170
                                            case "\n\t":
                                                v5 = (1 >= -2147483648)
                                            case "":
                                                if len(v44) > 0 {
                                                    v44[0] = int(v3)
                                                }
                                            }
                                        } else {
                                            v2 = 393
                                            if true {
                                                var v171 uint32 = v2
                                                _ = v171
                                            } else {
                                                if v5 {
                                                    if (-285 <= 1) {
                                                        switch (fi143 / (((^(fi143))) | 1)) {
                                                        case 9223372036854775807:
                                                            var v172 int8 = v1
                                                            _ = v172
                                                        case 42:
                                                            fmt.Println(v121)
                                                        case 32767:
                                                            if len(v44) >= 2 {
                                                                v44 = v44[1:len(v44)]
                                                            }
                                                        }
                                                    }
                                                    _ = len(v44)
                                                    v4 += 0
                                                    v43 = ("" + "foo bar")
                                                }
                                            }
                                        }
                                    }
                                    var v173 float32 = (v121 * (((1e10 / v121) / (1e10 - float32(float64(v121)))) / (v121 - v121)))
                                    if (v5 || (v118 >= 0.0)) {
                                        fmt.Println(v2, fi122, v118)
                                        v173 = float32((v118 / v118))
                                    } else {
                                        v43 = ((v43 + (v117 + (("a" + "\n\t") + "\n\t"))) + v43)
                                        if (v118 == v118) {
                                            if false {
                                                if v5 {
                                                    if false {
                                                        if v5 {
                                                            if (v5 || true) {
                                                                v4 -= 652
                                                                var v174 float64 = (v118 - float64((float32(1.5) + (v121 * v121))))
                                                                var v175 uint64 = 42
                                                                v4 += (v4 | (^(v4)))
                                                                _ = v174
                                                                _ = v175
                                                            }
                                                        }
                                                        v173 -= 0.1
                                                    } else {
                                                        v117 = v43
                                                        _ = len(v44)
                                                        var v176 uint = 0
                                                        _ = v176
                                                    }
                                                }
                                            }
                                            var v177 uint8 = 90
                                            if (v177 < (v177 ^ v177)) {
                                                v44 = append(v44, (fi122 * (fi122 % ((924) | 1))))
                                                v177 = 1
                                            } else {
                                                var v178 int = -2147483648
                                                v3 = v3
                                                var v179 string = "foo bar"
                                                switch 32767 {
                                                case 128:
                                                    v180 := GenS5{F0: 425, F1: v3, F2: (v4 & v4)}
                                                    _ = v180
                                                case 32767:
                                                    v181 := GenS17{F0: 0, F1: v117}
                                                    var v182 int = fi143
                                                    var v183 bool = v5
                                                    var v184 uint16 = v4
                                                    fmt.Println(v2, v3, v184)
                                                    _ = v181
                                                    _ = v182
                                                    _ = v183
                                                    _ = v184
                                                case 802:
                                                    var v185 int32 = 1
                                                    _ = v185
                                                }
                                                _ = v178
                                                _ = v179
                                            }
                                            _ = v177
                                        }
                                    }
                                    _ = v173
                                }
                            } else {
                                if false {
                                    fmt.Println(v1)
                                    v186 := GenS1{F0: v4, F1: v43}
                                    switch v117 {
                                    case "":
                                        switch fi122 {
                                        case 4294967295:
                                            fmt.Println(v4, v118, v4)
                                        case 414:
                                            if v5 {
                                                v43 = (v43 + (v117 + ((("a" + (v43 + v117)) + (("\n\t" + ("a" + ("" + "a"))) + ("a" + v43))) + "\n\t")))
                                            }
                                        }
                                    case "foo bar":
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(v2)
                                        case "\n\t":
                                            _ = len(v44)
                                        }
                                    case "\n\t":
                                        v117 = "hello"
                                    default:
                                        var v187 bool = false
                                        _ = v187
                                    }
                                    _ = v186
                                } else {
                                    if ((v5 && v5) && v5) {
                                        fmt.Println(v1, v2, v4)
                                        v188 := GenS1{F0: 0, F1: (v117 + ((v117 + "a") + "foo bar"))}
                                        _ = v188
                                    }
                                    if ((fi122 << ((364) & 63)) == -1) {
                                        switch (v43 + ("a" + "\n\t")) {
                                        case "hello":
                                            v44 = append(v44, fi122)
                                        case "foo bar":
                                            if len(v44) > 0 {
                                                v44[0] = -814
                                            }
                                        case "a":
                                            var v189 uint64 = 4294967295
                                            _ = v189
                                        default:
                                            var v190 int32 = 506
                                            _ = v190
                                        }
                                    } else {
                                        v44 = append(v44, (fi122 - (fi122 + (fi122 * (fi122 + fi143)))))
                                    }
                                    var v191 uint32 = 65535
                                    v44 = append(v44, (fi122 ^ fi143))
                                    _ = v191
                                }
                                v3 = int64(v4)
                            }
                        case -1:
                            v192 := GenS18{F0: 1000}
                            _ = v192
                        case -725:
                            v193 := GenS19{F0: 18446744073709551615, F1: 785, F2: v117}
                            _ = v193
                        default:
                            v5 = v5
                        }
                    }
                    _ = fi143
                }
                var v194 int64 = 0
                _ = v194
            }
        }
        v195 := GenS10{F0: 42}
        _ = fi122
        _ = v123
        _ = v195
    }
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v117
    _ = v118
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v46
    _ = v5
}
