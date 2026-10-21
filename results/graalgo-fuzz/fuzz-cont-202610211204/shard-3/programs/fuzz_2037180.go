package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 string }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 0
    var v3 bool = (1.5 <= 0.1)
    fmt.Println(v2, v3)
    if (2 != 65535) {
        switch 128 {
        case 1:
            var v4 uint32 = 891
            _ = v4
        case -9223372036854775808:
            delete(v1, 4294967295)
        case 0:
            var v5 int64 = 0
            _ = v5
        default:
            var v6 uint64 = 802
            _ = v6
        }
        var v7 int64 = -1
        _ = v7
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 42, F1: 154}
            v2 = 127
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int8 = (v2 - v2)
        _ = fi10
        _ = v11
    }
    _ = len(v1)
    var v12 bool = (915 >= 0)
    if ((v2 * v2) != (v2 >> ((240) & 63))) {
        v2 += 0
        if (0 == 42) {
            for fi13 := 0; fi13 < 1; fi13++ {
                if ("foo bar" != (("\n\t" + "a") + "foo bar")) {
                    switch fi13 {
                    case 0:
                        v14 := v1[(-(fi13))]
                        _ = v14
                    case 4294967295:
                        v15 := GenS1{F0: 42, F1: 1}
                        _ = v15
                    case 2147483647:
                        for fi16 := 0; fi16 < 1; fi16++ {
                            if (-651 >= 42) {
                                var v17 float32 = 0.1
                                if (42 != 65535) {
                                    v18 := GenS1{F0: 264, F1: 2}
                                    var v19 uint16 = 0
                                    var v20 uint32 = 750
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                _ = v17
                            }
                            if v12 {
                                if true {
                                    if v3 {
                                        var v21 int64 = 9223372036854775807
                                        if (2 < 692) {
                                            v12 = false
                                            v2 = -1
                                        }
                                        v22 := []int{}
                                        _ = v21
                                        _ = v22
                                    }
                                    v23 := GenS1{F0: 2, F1: 1}
                                    v24 := []int{fi13, (fi13 * fi16)}
                                    v3 = v12
                                    if len(v24) > 0 {
                                        v24[0] = (fi13 / ((fi16) | 1))
                                    }
                                    if v12 {
                                        if (((36 > -9223372036854775808) || (v3 || true)) && v12) {
                                            if (v3 && v12) {
                                                v25 := GenS1{F0: 724, F1: 42}
                                                if len(v24) > 0 {
                                                    v24[0] = 32768
                                                }
                                                _ = v25
                                            } else {
                                                var v26 bool = v12
                                                v3 = (v26 || (482 >= 122))
                                                v27 := GenS2{F0: 0, F1: -1, F2: ("hello" + ("foo bar" + "a"))}
                                                _ = v26
                                                _ = v27
                                            }
                                            _ = len(v1)
                                            if false {
                                                if v12 {
                                                    v28 := GenS3{F0: (0.1 != (3.14159 * 1e10)), F1: v3, F2: 1}
                                                    v12 = (!v12)
                                                    _ = v28
                                                } else {
                                                    if true {
                                                        v29 := v1[60]
                                                        var v30 int64 = -9223372036854775808
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        var v31 uint64 = 952
                                                        v1[fi16] = 0
                                                        _ = v31
                                                    }
                                                    switch ("foo bar" + ("foo bar" + "a")) {
                                                    case "":
                                                        v3 = (!(1 == 4294967295))
                                                    case "a":
                                                        v12 = (-744 != int64(v2))
                                                    }
                                                    fmt.Println(v12)
                                                }
                                                v3 = (!false)
                                            } else {
                                                v12 = (32767 <= 65535)
                                                var v32 string = ""
                                                var v33 float64 = (0.0 * 3.14159)
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                        var v34 uint = 0
                                        v2 = (v2 * v2)
                                        v35 := GenS2{F0: 2147483648, F1: 754, F2: ""}
                                        v2 += v2
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v36 := GenS4{F0: 18446744073709551615, F1: 32767, F2: 0}
                                        _ = v36
                                    }
                                    _ = v23
                                    _ = v24
                                }
                                v37 := v1[32767]
                                _ = v37
                            }
                            _ = fi16
                        }
                    default:
                        v38 := GenS4{F0: 42, F1: -32768, F2: 13}
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 16; fi39++ {
                        if true {
                            v12 = (v3 || v3)
                            var v40 float32 = 1.5
                            _ = v40
                        }
                        var v41 int32 = 42
                        var v42 int8 = 127
                        fmt.Println(fi13)
                        if (163 <= 255) {
                            v41 = 288
                            v2 += v42
                            v41 -= 0
                        } else {
                            var v43 float64 = 1e10
                            _ = v43
                        }
                        _ = fi39
                        _ = v41
                        _ = v42
                    }
                    var v44 int = fi13
                    _ = v44
                } else {
                    switch "foo bar" {
                    case "hello":
                        switch ("" + "") {
                        case "":
                            v2 = 127
                        case "a":
                            fmt.Println(v2, v3, fi13)
                        }
                    case "foo bar":
                        for fi45 := 0; fi45 < 4; fi45++ {
                            v46 := []int{(-(fi13)), fi13}
                            delete(v1, fi13)
                            _ = fi45
                            _ = v46
                        }
                    }
                    if true {
                        fmt.Println(v3, v3)
                    } else {
                        if ("" != "hello") {
                            fmt.Println(v12, v3)
                            if v12 {
                                v3 = (42 != -2147483648)
                            }
                        } else {
                            switch 42 {
                            case 127:
                                for fi47 := 0; fi47 < 20; fi47++ {
                                    var v48 float32 = (float32(1.5) / (float32(1.5) * 0.1))
                                    _ = fi47
                                    _ = v48
                                }
                            case 11:
                                var v49 float32 = -2.25
                                _ = v49
                            default:
                                for fi50 := 0; fi50 < 7; fi50++ {
                                    fmt.Println(v3, v3, fi50)
                                    fmt.Println(v3, v12, v12)
                                    _ = fi50
                                }
                            }
                            var v51 int16 = 1000
                            _ = v51
                        }
                    }
                }
                _ = fi13
            }
            if false {
                var v52 int32 = 0
                switch 42 {
                case 0:
                    for fi53 := 0; fi53 < 2; fi53++ {
                        v54 := v1[fi53]
                        switch "a" {
                        case "":
                            v55 := []int{fi53}
                            _ = v55
                        case "hello":
                            if (true || v3) {
                                var v56 int32 = (v52 * -2147483648)
                                _ = v56
                            }
                        case "a":
                            v1[fi53] = (v54 % ((4294967295) | 1))
                        default:
                            v54 -= (v54 * v54)
                        }
                        if v3 {
                            v12 = v12
                            for fi57 := 0; fi57 < 9; fi57++ {
                                if true {
                                    v2 -= v2
                                    v12 = v3
                                    v58 := GenS5{F0: 52}
                                    _ = v58
                                } else {
                                    v52 = (v52 << ((65535) & 63))
                                }
                                var v59 uint32 = v54
                                _ = fi57
                                _ = v59
                            }
                        }
                        _ = fi53
                        _ = v54
                    }
                case 877:
                    _ = len(v1)
                case 727:
                    var v60 bool = ((^(v52)) == (v52 / (((^(v52))) | 1)))
                    _ = v60
                }
                _ = v52
            } else {
                if false {
                    if (!(999 > 42)) {
                        fmt.Println(v3)
                        var v61 uint32 = 0
                        if v12 {
                            v62 := GenS2{F0: (v61 - (v61 + (v61 % ((v61) | 1)))), F1: -2147483648, F2: ("a" + "foo bar")}
                            if (!true) {
                                var v63 float32 = 1.5
                                v61 -= (v61 >> ((779) & 63))
                                v12 = true
                                _ = v63
                            } else {
                                v64 := GenS3{F0: (v12 && (v12 && (-128 == (v2 << ((0) & 63))))), F1: true, F2: 2}
                                _ = v64
                            }
                            v65 := []int{0}
                            _ = v62
                            _ = v65
                        } else {
                            switch "hello" {
                            case "a":
                                for fi66 := 0; fi66 < 9; fi66++ {
                                    for fi67 := 0; fi67 < 9; fi67++ {
                                        v68 := []int{4294967295, (fi67 ^ (fi66 - fi67)), fi67, 0}
                                        var v69 uint64 = 42
                                        v12 = v12
                                        _ = fi67
                                        _ = v68
                                        _ = v69
                                    }
                                    for fi70 := 0; fi70 < 16; fi70++ {
                                        v12 = v12
                                        var v71 uint64 = 4294967295
                                        var v72 int = fi70
                                        _ = fi70
                                        _ = v71
                                        _ = v72
                                    }
                                    if (fi66 >= fi66) {
                                        fmt.Println(fi66, fi66, v61)
                                    }
                                    _ = fi66
                                }
                            case "\n\t":
                                v1[9223372036854775807] = (v61 + v61)
                            }
                            for fi73 := 0; fi73 < 15; fi73++ {
                                if true {
                                    v61 = 2147483647
                                    var v74 uint8 = 2
                                    _ = v74
                                }
                                _ = fi73
                            }
                            if (!true) {
                                if v12 {
                                    fmt.Println(v2, v12, v61)
                                    v12 = ((270 < v61) || (862 <= 1))
                                }
                                for fi75 := 0; fi75 < 15; fi75++ {
                                    var v76 float64 = float64(float32(1e10))
                                    var v77 uint64 = 0
                                    v78 := GenS1{F0: 18446744073709551615, F1: 242}
                                    _ = fi75
                                    _ = v76
                                    _ = v77
                                    _ = v78
                                }
                                switch "hello" {
                                case "\n\t":
                                    for fi79 := 0; fi79 < 6; fi79++ {
                                        v61 -= v61
                                        _ = fi79
                                    }
                                case "foo bar":
                                    var v80 float64 = ((0.0 * float64(float32(float64(0.1)))) - 1.5)
                                    _ = v80
                                case "hello":
                                    for fi81 := 0; fi81 < 9; fi81++ {
                                        v3 = v3
                                        _ = fi81
                                    }
                                }
                            } else {
                                v82 := v1[-1]
                                _ = v82
                            }
                        }
                        _ = v61
                    }
                    var v83 uint8 = 40
                    _ = v83
                } else {
                    var v84 float32 = 1e10
                    v85 := GenS4{F0: 126, F1: -523, F2: 0}
                    var v86 bool = (-9223372036854775808 == 42)
                    _ = v84
                    _ = v85
                    _ = v86
                }
                switch 32768 {
                case 127:
                    v87 := GenS3{F0: (true || (100 > (-(v2)))), F1: (!(!v12)), F2: 62}
                    _ = v87
                case 32768:
                    if false {
                        var v88 int = 4294967295
                        v89 := GenS5{F0: 0}
                        var v90 int16 = 270
                        _ = v88
                        _ = v89
                        _ = v90
                    } else {
                        switch 42 {
                        case 128:
                            v12 = true
                        }
                    }
                case 4294967295:
                    _ = len(v1)
                default:
                    for fi91 := 0; fi91 < 11; fi91++ {
                        v2 = v2
                        if v12 {
                            if (!(42 > 42)) {
                                v92 := GenS4{F0: 4294967295, F1: 288, F2: 3}
                                if (2 <= 42) {
                                    v3 = (!(-880 >= 567))
                                    if (true && v12) {
                                        v93 := []int{fi91, 169, fi91}
                                        v94 := GenS5{F0: (-(v2))}
                                        _ = v93
                                        _ = v94
                                    } else {
                                        delete(v1, (fi91 >> ((42) & 63)))
                                    }
                                }
                                v2 = 0
                                _ = v92
                            }
                            for fi95 := 0; fi95 < 9; fi95++ {
                                fmt.Println(fi95, fi95)
                                if (v2 != (v2 / ((v2) | 1))) {
                                    v96 := []int{fi95, (fi95 % ((127) | 1)), fi91, (fi91 + fi91)}
                                    _ = v96
                                } else {
                                    if (v3 || false) {
                                        v2 = (v2 & (v2 + v2))
                                        var v97 int = (fi95 / (((fi91 ^ 384)) | 1))
                                        _ = v97
                                    }
                                }
                                v98 := []int{fi91, fi95}
                                _ = fi95
                                _ = v98
                            }
                            v2 = v2
                        }
                        v99 := GenS3{F0: (4294967295 > 0), F1: (v3 && v12), F2: 127}
                        _ = fi91
                        _ = v99
                    }
                }
                switch -306 {
                case -2147483648:
                    fmt.Println(v3)
                case 4294967295:
                    v100 := []int{-730, 127, 32768, 1}
                    _ = v100
                case 127:
                    switch "foo bar" {
                    case "foo bar":
                        v101 := []int{425, -301, 690, -2147483648}
                        _ = v101
                    case "a":
                        var v102 float32 = 1e10
                        _ = v102
                    default:
                        var v103 uint64 = 0
                        _ = v103
                    }
                }
            }
        }
        switch 32768 {
        case 359:
            v104 := GenS4{F0: 91, F1: 1000, F2: 128}
            _ = v104
        case 4294967295:
            var v105 float64 = (1.5 - float64(0.1))
            _ = v105
        }
    }
    v2 += 127
    fmt.Println(v12)
    v3 = v3
    var v106 uint16 = 1
    var v107 float64 = 0.1
    v107 -= v107
    v12 = (255 > 32767)
    v108 := GenS6{F0: 0}
    if (-1 > 42) {
        v109 := []int{1, 1, -1, -146}
        fmt.Println(v12, v12)
        _ = v109
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v106
    _ = v107
    _ = v108
    _ = v12
    _ = v2
    _ = v3
}
