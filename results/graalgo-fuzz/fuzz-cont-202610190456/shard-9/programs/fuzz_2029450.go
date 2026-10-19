package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint8; F2 string }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 string }

func main() {
    var v1 uint64 = 365
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    var v5 uint32 = 0
    var v6 uint = 0
    v7 := []int{128, -744, 1, 0}
    v8 := GenS1{F0: -1, F1: int32(v6)}
    var v9 int32 = 1
    fmt.Println(v5, v1, v5)
    var v10 string = (((("a" + "hello") + ("foo bar" + "a")) + "") + (("foo bar" + ("\n\t" + "foo bar")) + "\n\t"))
    if (true && false) {
        var v11 int8 = 42
        var v12 string = "\n\t"
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        switch (fi13 | fi13) {
        case -450:
            var v14 uint32 = v5
            _ = v14
        case -2147483648:
            var v15 uint32 = (v5 & v5)
            _ = v15
        case -9223372036854775808:
            var v16 float64 = (1e10 - float64(0.0))
            _ = v16
        }
        v1 = uint64(v5)
        _ = fi13
    }
    if len(v7) > 0 {
        v7[0] = 127
    }
    _ = len(v7)
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int8 = 111
        _ = fi17
        _ = v18
    }
    switch -2147483648 {
    case -202:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case 42:
        if false {
            if len(v7) > 0 {
                v7[0] = 2147483647
            }
            fmt.Println(v6, v9)
            v19 := GenS2{F0: v1, F1: 128, F2: v10}
            _ = v19
        } else {
            switch 32767 {
            case -9223372036854775808:
                if (float32(0.1) < float32((((3.14159 + float64(-2.25)) / 1.5) / (0.1 * (1.5 - ((0.1 * 0.0) / (1e10 / ((float64(1.5) / (float64((((1e10 + float32(1e10)) / 1e10) / 1.0)) - 1.5)) - (-2.25 - 3.14159))))))))) {
                    switch (v10 + (v10 + (v10 + (v10 + "\n\t")))) {
                    case "\n\t":
                        var v20 uint = 578
                        _ = v20
                    case "":
                        v1 = (^(v1))
                    default:
                        v21 := GenS1{F0: v9, F1: v9}
                        _ = v21
                    }
                    v6 -= (v6 | v6)
                    v5 += 781
                } else {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        v1 += 4294967295
                        _ = fi22
                    }
                    v5 += v5
                }
            case 42:
                v1 += v1
            }
            fmt.Println(v5, v6, v1)
        }
    case 4294967295:
        if false {
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS3{F0: uint(v5)}
                var v25 int8 = 0
                var v26 int32 = v9
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            v27 := GenS3{F0: v6}
            _ = v27
        } else {
            v28 := GenS2{F0: v1, F1: 42, F2: (v10 + "a")}
            v7 = append(v7, -1)
            _ = v28
        }
    default:
        switch -9223372036854775808 {
        case 1:
            for fi29 := 0; fi29 < 11; fi29++ {
                switch 1 {
                case 488:
                    if (4294967295 > (^(v1))) {
                        var v30 uint64 = (^(v1))
                        _ = v30
                    } else {
                        switch ("" + (v10 + v10)) {
                        case "":
                            var v31 float64 = ((0.1 - (-2.25 - (1.5 * float64(1.5)))) * 1.5)
                            _ = v31
                        case "foo bar":
                            v10 = "hello"
                        default:
                            if false {
                                fmt.Println(fi29)
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                v1 = 0
                            } else {
                                if (1e10 >= (((-2.25 / 1.0) / float64((3.14159 / float32(((-2.25 - (float64(float32(((1.5 - 1e10) / 1e10))) * -2.25)) / 1e10))))) - (0.1 * 0.1))) {
                                    v6 = 65535
                                }
                            }
                        }
                        fmt.Println(fi29, fi29, v10)
                        var v32 int8 = -6
                        _ = v32
                    }
                default:
                    v33 := GenS2{F0: (v1 / (((v1 << (((v6 * v6)) & 63))) | 1)), F1: (fi29 >> (((v6 >> ((2) & 63))) & 63)), F2: v10}
                    _ = v33
                }
                _ = fi29
            }
        case -266:
            fmt.Println(v5, v5, v5)
        default:
            v10 = v10
        }
    }
    if (127 < 9223372036854775807) {
        for fi34 := 0; fi34 < 19; fi34++ {
            if (-1 == 100) {
                var v35 uint32 = 1
                var v36 uint32 = v35
                fmt.Println(v36)
                _ = v35
                _ = v36
            } else {
                v37 := v2[9223372036854775807]
                fmt.Println(fi34)
                if (true && (!(false || (false || false)))) {
                    var v38 uint32 = v5
                    _ = v38
                }
                _ = v37
            }
            delete(v3, v10)
            fmt.Println(fi34)
            _ = fi34
        }
        v39 := GenS3{F0: v6}
        _ = v39
    } else {
        if ((false && (!((v10 > v10) || (!(0.0 == 0.0))))) || true) {
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                var v40 int8 = 42
                _ = v40
            case "":
                var v41 uint16 = 872
                _ = v41
            case "\n\t":
                for fi42 := 0; fi42 < 10; fi42++ {
                    v43 := GenS3{F0: v6}
                    fmt.Println(v6, v6)
                    fmt.Println(v5, v10)
                    _ = fi42
                    _ = v43
                }
            default:
                var v44 uint = 4294967295
                _ = v44
            }
            var v45 string = v10
            _ = v45
        }
        if (42 > -128) {
            var v46 int8 = -128
            switch 256 {
            case 931:
                for fi47 := 0; fi47 < 14; fi47++ {
                    switch int(v1) {
                    case -2147483648:
                        if (0 > 660) {
                            v48 := GenS4{F0: (v10 + "foo bar"), F1: (-(fi47))}
                            v46 += (v46 << (((v6 ^ (v6 + (v6 ^ uint(v5))))) & 63))
                            _ = v48
                        }
                    case 26:
                        if true {
                            fmt.Println(v10, v46, v46)
                            switch fi47 {
                            case -1:
                                switch v10 {
                                case "hello":
                                    v46 += v46
                                case "":
                                    if ((true || (float32(-2.25) >= (1e10 * 0.0))) && (!(true && false))) {
                                        fmt.Println(v6)
                                    } else {
                                        fmt.Println(v46)
                                    }
                                case "foo bar":
                                    for fi49 := 0; fi49 < 18; fi49++ {
                                        v50 := GenS2{F0: (v1 - (^(v1))), F1: (fi47 % (((fi49 * (fi47 ^ (fi47 & int(v46))))) | 1)), F2: (v10 + ((v10 + ((v10 + (v10 + v10)) + v10)) + v10))}
                                        if (false && true) {
                                            if (1 > int(v5)) {
                                                var v51 int32 = -111
                                                if (3.14159 > 1.5) {
                                                    v52 := v2[128]
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    if true {
                                                        var v53 uint16 = 65535
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                        _ = v53
                                                    }
                                                    _ = v52
                                                } else {
                                                    switch v10 {
                                                    case "":
                                                        if len(v7) > 0 {
                                                            v7[0] = fi47
                                                        }
                                                    default:
                                                        v3[(v10 + (v10 + (v10 + "a")))] = (v46 < 1)
                                                    }
                                                    if true {
                                                        if (v10 < "hello") {
                                                            var v54 int64 = -9223372036854775808
                                                            v46 -= v46
                                                            v5 -= uint32(v1)
                                                            _ = v54
                                                        } else {
                                                            switch ((v10 + "foo bar") + "") {
                                                            case "foo bar":
                                                                v51 = (^(v51))
                                                            default:
                                                                var v55 int = (fi49 * (fi49 | 128))
                                                                _ = v55
                                                            }
                                                            v46 += (v46 / (((v46 ^ (v46 + -44))) | 1))
                                                            if len(v7) >= 2 {
                                                                v7 = v7[1:len(v7)]
                                                            }
                                                        }
                                                        var v56 bool = true
                                                        _ = v56
                                                    }
                                                }
                                                _ = v51
                                            }
                                        }
                                        if true {
                                            var v57 uint64 = (v1 / ((2) | 1))
                                            if (-908 < -1) {
                                                v58 := GenS4{F0: ("foo bar" + ("a" + v10)), F1: 4294967295}
                                                v57 += 18446744073709551615
                                                switch "hello" {
                                                case "a":
                                                    _ = len(v7)
                                                case "\n\t":
                                                    var v59 uint = 1
                                                    _ = v59
                                                case "":
                                                    if true {
                                                        if false {
                                                            v60 := v2[838]
                                                            v61 := GenS2{F0: (v1 * v1), F1: (fi49 * (^(v60))), F2: (v10 + v10)}
                                                            var v62 uint = (v6 - (^(v6)))
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            if (fi47 == fi49) {
                                                                if false {
                                                                    fmt.Println(v9)
                                                                    if (0.1 >= float32(1.5)) {
                                                                        switch (fi47 >> ((65535) & 63)) {
                                                                        case -694:
                                                                            var v63 float64 = -2.25
                                                                            _ = v63
                                                                        case 128:
                                                                            v57 = v1
                                                                        }
                                                                        if false {
                                                                            v64 := GenS1{F0: v9, F1: (v9 + -1)}
                                                                            fmt.Println(v10)
                                                                            _ = v64
                                                                        }
                                                                        v65 := GenS5{F0: 912}
                                                                        _ = v65
                                                                    } else {
                                                                        if ((int64(v57) != 1) && (!true)) {
                                                                            var v66 int16 = int16(v46)
                                                                            fmt.Println(v6)
                                                                            if false {
                                                                                fmt.Println(fi49, v6)
                                                                            }
                                                                            _ = v66
                                                                        } else {
                                                                            v67 := GenS2{F0: v57, F1: fi49, F2: ""}
                                                                            v3[((v10 + v10) + v10)] = false
                                                                            _ = v67
                                                                        }
                                                                        fmt.Println(v6, v6, v5)
                                                                    }
                                                                    fmt.Println(v6, v9, v9)
                                                                }
                                                                var v68 uint8 = 128
                                                                v69 := GenS6{F0: 1, F1: uint8(v6), F2: (v10 + ("foo bar" + ("hello" + (v10 + "\n\t"))))}
                                                                _ = v68
                                                                _ = v69
                                                            } else {
                                                                fmt.Println(v5)
                                                                var v70 uint = v6
                                                                _ = v70
                                                            }
                                                            if (1 < 0) {
                                                                v71 := GenS5{F0: 32767}
                                                                var v72 uint = 865
                                                                _ = v71
                                                                _ = v72
                                                            } else {
                                                                var v73 float64 = 1.5
                                                                v74 := GenS3{F0: (^(v6))}
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                            v75 := GenS1{F0: v9, F1: v9}
                                                            _ = v75
                                                        }
                                                        _ = len(v3)
                                                        var v76 int16 = -576
                                                        fmt.Println(fi47, v1, v10)
                                                        if false {
                                                            if (-2147483648 > (fi47 ^ 2147483647)) {
                                                                var v77 int64 = 9223372036854775807
                                                                var v78 uint = (v6 % (((v6 * v6)) | 1))
                                                                var v79 float64 = 0.0
                                                                _ = v77
                                                                _ = v78
                                                                _ = v79
                                                            } else {
                                                                var v80 uint64 = (v1 >> ((v6) & 63))
                                                                if false {
                                                                    v6 = 2
                                                                    v81 := GenS6{F0: 1000, F1: 128, F2: "foo bar"}
                                                                    v2[(fi49 >> ((v6) & 63))] = (fi49 % ((fi49) | 1))
                                                                    var v82 float64 = -2.25
                                                                    v5 += v5
                                                                    _ = v81
                                                                    _ = v82
                                                                } else {
                                                                    v6 += uint(v57)
                                                                    var v83 uint32 = 0
                                                                    _ = v83
                                                                }
                                                                switch fi49 {
                                                                case 2147483647:
                                                                    v84 := v2[fi47]
                                                                    _ = v84
                                                                case -9223372036854775808:
                                                                    var v85 int64 = -9223372036854775808
                                                                    _ = v85
                                                                }
                                                                _ = v80
                                                            }
                                                        } else {
                                                            v3[""] = (true || (70 >= 1))
                                                            switch v10 {
                                                            case "a":
                                                                var v86 float32 = 0.0
                                                                _ = v86
                                                            case "\n\t":
                                                                var v87 uint32 = (^(v5))
                                                                _ = v87
                                                            case "hello":
                                                                if (uint64(v5) >= 18446744073709551615) {
                                                                    if (!(!false)) {
                                                                        var v88 int32 = (v9 | (v9 - v9))
                                                                        _ = v88
                                                                    }
                                                                } else {
                                                                    var v89 int8 = (-(v46))
                                                                    v90 := GenS4{F0: ((v10 + v10) + (v10 + "\n\t")), F1: (^(fi47))}
                                                                    var v91 int8 = v89
                                                                    if ((v6 & (v6 % ((v6) | 1))) < (v6 ^ 2)) {
                                                                        var v92 bool = false
                                                                        var v93 int = -585
                                                                        _ = v92
                                                                        _ = v93
                                                                    } else {
                                                                        v94 := GenS7{F0: (fi47 * (fi49 & (^(fi47)))), F1: 42}
                                                                        if true {
                                                                            fmt.Println(fi47)
                                                                            v95 := GenS1{F0: (v9 >> ((42) & 63)), F1: -2147483648}
                                                                            _ = v95
                                                                        } else {
                                                                            if true {
                                                                                v89 = 86
                                                                                v89 = v91
                                                                                v57 = uint64(v5)
                                                                                fmt.Println(fi49, v91)
                                                                            } else {
                                                                                fmt.Println(v91, v10)
                                                                                if (0.0 <= 1e10) {
                                                                                    delete(v2, fi49)
                                                                                    _ = len(v7)
                                                                                } else {
                                                                                    var v96 uint16 = uint16(v57)
                                                                                    _ = v96
                                                                                }
                                                                                if (false && true) {
                                                                                    v7 = append(v7, fi47)
                                                                                    v91 -= v91
                                                                                }
                                                                            }
                                                                        }
                                                                        var v97 uint64 = v57
                                                                        v5 -= (^(v5))
                                                                        v98 := GenS2{F0: 18446744073709551615, F1: (fi49 & (fi47 + (fi49 - -1))), F2: "a"}
                                                                        _ = v94
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                    _ = v89
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                            default:
                                                                if false {
                                                                    fmt.Println(v1)
                                                                    if (((!false) || (((true || (255 != uint8(v6))) || false) || ((0.0 / 3.14159) == 3.14159))) && false) {
                                                                        switch (v10 + (v10 + v10)) {
                                                                        case "\n\t":
                                                                            switch (fi47 | fi47) {
                                                                            case 42:
                                                                                var v99 string = v10
                                                                                _ = v99
                                                                            case -427:
                                                                                v100 := GenS8{F0: "a"}
                                                                                _ = v100
                                                                            case 127:
                                                                                fmt.Println(v46, fi47)
                                                                            default:
                                                                                var v101 int16 = v76
                                                                                _ = v101
                                                                            }
                                                                        default:
                                                                            var v102 uint32 = v5
                                                                            _ = v102
                                                                        }
                                                                    } else {
                                                                        var v103 int = fi49
                                                                        var v104 uint32 = (v5 * v5)
                                                                        _ = v103
                                                                        _ = v104
                                                                    }
                                                                    var v105 float32 = (3.14159 - -2.25)
                                                                    _ = v105
                                                                }
                                                            }
                                                        }
                                                        _ = v76
                                                    } else {
                                                        v3[(v10 + v10)] = (int8(v5) == (v46 << ((762) & 63)))
                                                        v9 = (v9 * v9)
                                                        var v106 bool = false
                                                        _ = v106
                                                    }
                                                    _ = len(v7)
                                                default:
                                                    switch ((v10 + "a") + ((("a" + ("a" + v10)) + ((v10 + v10) + v10)) + v10)) {
                                                    case "hello":
                                                        fmt.Println(v57)
                                                    }
                                                }
                                                _ = v58
                                            }
                                            v107 := GenS1{F0: -744, F1: -652}
                                            _ = v107
                                            _ = v57
                                        }
                                        if (!false) {
                                            switch (((v10 + v10) + (v10 + v10)) + "\n\t") {
                                            case "\n\t":
                                                v108 := GenS4{F0: "foo bar", F1: (fi47 & (fi49 * (fi49 + -135)))}
                                                _ = v108
                                            case "a":
                                                switch v10 {
                                                case "foo bar":
                                                    var v109 int8 = v46
                                                    _ = v109
                                                case "":
                                                    if (!(!(255 >= 255))) {
                                                        var v110 int16 = int16(v6)
                                                        var v111 string = (v10 + ((v10 + "a") + (v10 + ("foo bar" + v10))))
                                                        _ = v110
                                                        _ = v111
                                                    } else {
                                                        fmt.Println(v10, v9, v6)
                                                        v1 = (^(v1))
                                                        v46 -= v46
                                                    }
                                                    v112 := GenS7{F0: fi49, F1: 255}
                                                    _ = v112
                                                default:
                                                    var v113 uint8 = uint8(v1)
                                                    _ = v113
                                                }
                                            default:
                                                var v114 bool = (-791 < 0)
                                                _ = v114
                                            }
                                            v115 := GenS8{F0: "a"}
                                            _ = v115
                                        } else {
                                            v10 = v10
                                        }
                                        _ = fi49
                                        _ = v50
                                    }
                                }
                            case 4294967295:
                                var v116 int16 = 32767
                                _ = v116
                            }
                        }
                    }
                    fmt.Println(v46)
                    _ = fi47
                }
            case 24:
                v1 = v1
            case 127:
                var v117 uint = v6
                _ = v117
            }
            v118 := v3[(v10 + (v10 + "hello"))]
            _ = v118
            _ = v46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
