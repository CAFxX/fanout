package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 int64 }
type GenS3 struct { F0 uint32; F1 int; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int64; F1 int64; F2 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 int; F1 int32; F2 int32 }
type GenS12 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = -2.25
    var v3 uint64 = 961
    switch 413 {
    case 1:
        var v4 uint16 = 65535
        _ = v4
    case 32767:
        v5 := []int{}
        _ = v5
    case 785:
        fmt.Println(v2, v2)
    }
    switch 313 {
    case 4294967295:
        var v6 int32 = -2147483648
        _ = v6
    case 0:
        var v7 bool = false
        _ = v7
    default:
        v2 += v2
    }
    fmt.Println(v2, v3)
    v3 += v3
    switch -464 {
    case -920:
        v8 := []int{-9223372036854775808}
        _ = v8
    case 1:
        v2 += v2
    default:
        if (!true) {
            v2 = float64(0.0)
            v9 := GenS1{F0: true, F1: 2}
            fmt.Println(v2, v3)
            _ = v9
        } else {
            var v10 int16 = 966
            for fi11 := 0; fi11 < 17; fi11++ {
                switch 32767 {
                case -395:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        delete(v1, "foo bar")
                        v3 += v3
                        v13 := GenS1{F0: ((v2 / 1.0) < (v2 - v2)), F1: 807}
                        var v14 int64 = 0
                        var v15 int32 = 1
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                }
                var v16 float32 = 0.0
                _ = fi11
                _ = v16
            }
            var v17 bool = true
            _ = v10
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        switch fi18 {
        case 1:
            fmt.Println(fi18, v2, v2)
        case 9223372036854775807:
            v2 = float64(float32(v2))
        case 155:
            var v19 int = fi18
            _ = v19
        }
        fmt.Println(v3, v2)
        _ = fi18
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int = fi20
        _ = fi20
        _ = v21
    }
    v22 := []int{0, -9223372036854775808, 0, 42}
    if len(v22) > 0 {
        v22[0] = -9223372036854775808
    }
    switch -9223372036854775808 {
    case 127:
        v3 += v3
    case 42:
        fmt.Println(v2, v3)
    case -564:
        for fi23 := 0; fi23 < 11; fi23++ {
            switch "foo bar" {
            case "a":
                switch 32767 {
                case 42:
                    if false {
                        var v24 string = "a"
                        var v25 uint16 = 255
                        _ = v24
                        _ = v25
                    }
                case 4294967295:
                    v3 = (^(v3))
                case -2147483648:
                    if true {
                        for fi26 := 0; fi26 < 14; fi26++ {
                            var v27 int8 = 0
                            fmt.Println(fi26, v27, v3)
                            if true {
                                v27 -= 127
                                fmt.Println(v2)
                                if true {
                                    var v28 uint8 = 24
                                    var v29 int32 = 42
                                    v22 = append(v22, (fi26 - (fi23 ^ fi26)))
                                    _ = v28
                                    _ = v29
                                }
                                v3 = v3
                            } else {
                                var v30 int16 = 32767
                                var v31 uint = 33
                                switch "" {
                                case "foo bar":
                                    var v32 int16 = 568
                                    _ = v32
                                case "\n\t":
                                    v33 := GenS2{F0: (fi23 % (((fi26 + fi26)) | 1)), F1: -9223372036854775808, F2: 42}
                                    _ = v33
                                case "hello":
                                    switch ("a" + "foo bar") {
                                    case "foo bar":
                                        var v34 uint8 = 128
                                        _ = v34
                                    default:
                                        v30 += (-(v30))
                                    }
                                }
                                _ = v30
                                _ = v31
                            }
                            _ = fi26
                            _ = v27
                        }
                    }
                }
            default:
                v35 := GenS1{F0: (!true), F1: 1}
                _ = v35
            }
            var v36 int = (^(fi23))
            _ = fi23
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        _ = len(v22)
        var v38 float32 = -2.25
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        v3 = 2
        var v40 uint16 = 774
        var v41 int = (fi39 % (((fi39 % (((-(fi39))) | 1))) | 1))
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := GenS3{F0: 0, F1: 32768, F2: (42 == 860)}
    fmt.Println(v3, v2)
    if (-32768 != 32767) {
        switch "\n\t" {
        case "":
            for fi43 := 0; fi43 < 12; fi43++ {
                if true {
                    v2 += v2
                }
                for fi44 := 0; fi44 < 6; fi44++ {
                    fmt.Println(fi44)
                    var v45 float64 = v2
                    v22 = append(v22, fi44)
                    v1["foo bar"] = (^(fi43))
                    if true {
                        var v46 int8 = 42
                        var v47 string = ((((("a" + ("foo bar" + "a")) + "") + "") + "hello") + "hello")
                        v48 := v1[(("hello" + (v47 + ((("a" + v47) + v47) + (("a" + (v47 + v47)) + v47)))) + (v47 + (v47 + ((((("foo bar" + v47) + (v47 + v47)) + v47) + (v47 + "")) + "hello"))))]
                        _ = v46
                        _ = v47
                        _ = v48
                    } else {
                        fmt.Println(v2, v45, v2)
                        delete(v1, "hello")
                        fmt.Println(v2, v3)
                        v2 += float64(0.0)
                        if true {
                            if (!(!((0.0 + -2.25) <= 0.0))) {
                                var v49 int32 = -1
                                v50 := GenS4{F0: -151}
                                _ = v49
                                _ = v50
                            } else {
                                v51 := v1["hello"]
                                switch (("" + (("hello" + "") + "a")) + "\n\t") {
                                case "hello":
                                    switch 32767 {
                                    case 128:
                                        _ = len(v22)
                                    }
                                default:
                                    var v52 int8 = 0
                                    _ = v52
                                }
                                _ = v51
                            }
                            delete(v1, "hello")
                            switch -9223372036854775808 {
                            case 243:
                                var v53 int8 = 42
                                _ = v53
                            case -1:
                                switch "\n\t" {
                                case "foo bar":
                                    v54 := GenS5{F0: -44, F1: -899, F2: 0}
                                    _ = v54
                                case "a":
                                    var v55 uint8 = 1
                                    _ = v55
                                default:
                                    v56 := GenS6{F0: v3, F1: 42}
                                    _ = v56
                                }
                            case 127:
                                v3 += (v3 & 42)
                            default:
                                _ = len(v22)
                            }
                            v57 := GenS5{F0: 1, F1: (fi44 % ((127) | 1)), F2: 17}
                            if ((true || (-2147483648 == 0)) && (78 <= 0)) {
                                var v58 uint8 = 255
                                _ = v58
                            } else {
                                fmt.Println(fi44, v2, v3)
                            }
                            _ = v57
                        } else {
                            v45 = ((v45 + 0.1) + (1.5 * (float64(1e10) - (v2 + -2.25))))
                            var v59 int8 = -36
                            _ = v59
                        }
                        if true {
                            if (0 != 255) {
                                v45 = (float64(0.0) * (((v2 / (v45 + (v2 - float64(-2.25)))) / v45) * v2))
                                switch "hello" {
                                case "hello":
                                    var v60 int16 = 42
                                    _ = v60
                                case "foo bar":
                                    if true {
                                        if true {
                                            v61 := v1[((("foo bar" + "") + ("\n\t" + "hello")) + ("hello" + "hello"))]
                                            v62 := GenS4{F0: -211}
                                            if (v61 <= fi44) {
                                                var v63 string = "a"
                                                _ = len(v1)
                                                fmt.Println(v63)
                                                _ = v63
                                            }
                                            _ = v61
                                            _ = v62
                                        } else {
                                            var v64 uint32 = 511
                                            switch 9223372036854775807 {
                                            case 418:
                                                v65 := GenS7{F0: 149, F1: 42}
                                                _ = v65
                                            }
                                            fmt.Println(v64)
                                            v66 := GenS5{F0: -1, F1: (fi43 & fi43), F2: 73}
                                            if (0 >= -1) {
                                                if false {
                                                    var v67 float32 = (3.14159 + (1.5 * (-2.25 * -2.25)))
                                                    v68 := GenS8{F0: int64(v3), F1: 0, F2: 1}
                                                    _ = v67
                                                    _ = v68
                                                }
                                                var v69 float32 = 1.5
                                                _ = v69
                                            } else {
                                                v3 += v3
                                                switch (fi44 + fi44) {
                                                case -385:
                                                    v22 = append(v22, (fi43 / ((fi43) | 1)))
                                                }
                                            }
                                            _ = v64
                                            _ = v66
                                        }
                                    }
                                case "a":
                                    var v70 uint32 = 655
                                    _ = v70
                                default:
                                    if true {
                                        if true {
                                            switch "a" {
                                            case "a":
                                                if (!false) {
                                                    v45 = (v2 * v45)
                                                }
                                            case "hello":
                                                switch (^(fi43)) {
                                                case 32768:
                                                    if (v45 >= ((-2.25 * v45) + v2)) {
                                                        if true {
                                                            if false {
                                                                if false {
                                                                    if len(v22) >= 2 {
                                                                        v22 = v22[1:len(v22)]
                                                                    }
                                                                    var v71 uint16 = uint16(v3)
                                                                    switch fi44 {
                                                                    case -187:
                                                                        var v72 uint16 = v71
                                                                        _ = v72
                                                                    }
                                                                    if (-834 > 1) {
                                                                        var v73 int32 = 15
                                                                        v71 -= (^(v71))
                                                                        var v74 float64 = (v2 - -2.25)
                                                                        _ = v73
                                                                        _ = v74
                                                                    }
                                                                    if true {
                                                                        v45 = v45
                                                                        if (-1 <= -559) {
                                                                            var v75 float32 = (1e10 + 1e10)
                                                                            v1["a"] = -2147483648
                                                                            _ = v75
                                                                        } else {
                                                                            fmt.Println(v2, v71)
                                                                            v22 = append(v22, fi43)
                                                                            fmt.Println(v71, fi43, fi43)
                                                                        }
                                                                    }
                                                                    _ = v71
                                                                }
                                                            } else {
                                                                v45 += 0.0
                                                                var v76 string = "\n\t"
                                                                if false {
                                                                    var v77 int64 = -1
                                                                    _ = v77
                                                                } else {
                                                                    if ((fi43 << ((518) & 63)) < (fi44 << ((4294967295) & 63))) {
                                                                        v76 = v76
                                                                        var v78 int32 = 2147483647
                                                                        _ = v78
                                                                    }
                                                                }
                                                                _ = len(v22)
                                                                _ = v76
                                                            }
                                                            var v79 int16 = -667
                                                            fmt.Println(v2, v45, v2)
                                                            _ = v79
                                                        }
                                                        if false {
                                                            var v80 int64 = -288
                                                            v22 = append(v22, fi44)
                                                            var v81 uint32 = 2147483648
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                    } else {
                                                        v82 := GenS8{F0: 46, F1: -9223372036854775808, F2: 89}
                                                        if (true && true) {
                                                            v83 := GenS9{F0: fi43}
                                                            fmt.Println(fi44, v3, fi43)
                                                            _ = v83
                                                        }
                                                        if false {
                                                            v45 += v45
                                                            switch fi44 {
                                                            case -1:
                                                                v3 = (v3 << ((42) & 63))
                                                            }
                                                        } else {
                                                            v2 -= v2
                                                            var v84 uint64 = (v3 / ((v3) | 1))
                                                            v85 := GenS10{F0: int64(fi43), F1: 9223372036854775807}
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                        _ = v82
                                                    }
                                                case -175:
                                                    v45 += (v2 + 1e10)
                                                }
                                            case "":
                                                if true {
                                                    var v86 float32 = (0.1 / 1.0)
                                                    var v87 uint8 = 1
                                                    v87 -= 64
                                                    _ = v86
                                                    _ = v87
                                                }
                                            default:
                                                v45 -= v45
                                            }
                                            v88 := GenS11{F0: -1, F1: -822, F2: -1}
                                            v2 -= (v45 + v45)
                                            _ = v88
                                        }
                                        if false {
                                            v3 -= (v3 ^ v3)
                                            fmt.Println(v3, fi43, fi43)
                                        }
                                        var v89 float64 = v2
                                        _ = v89
                                    } else {
                                        v45 -= (v45 * ((v45 + v2) / 1.0))
                                        v90 := v1[(("a" + ("foo bar" + ("a" + "\n\t"))) + "")]
                                        v3 = 18446744073709551615
                                        if false {
                                            v3 -= v3
                                        } else {
                                            if (((0.0 + (0.0 * ((1e10 * (1e10 - 1e10)) - ((float32((v45 + v45)) * (1.5 / 3.14159)) / -2.25)))) / -2.25) >= (-2.25 * (float32(((((v2 * (3.14159 - v45)) + (v2 / v2)) * 3.14159) * v2)) / 3.14159))) {
                                                v91 := GenS11{F0: (fi44 % ((fi43) | 1)), F1: 204, F2: -883}
                                                _ = v91
                                            } else {
                                                if len(v22) > 0 {
                                                    v22[0] = v90
                                                }
                                                fmt.Println(v45, v2, fi43)
                                            }
                                            if false {
                                                if len(v22) > 0 {
                                                    v22[0] = 366
                                                }
                                                v2 = 1.5
                                                switch "" {
                                                case "foo bar":
                                                    v90 += -779
                                                case "a":
                                                    var v92 int = fi43
                                                    _ = v92
                                                }
                                                v3 -= v3
                                                v93 := GenS7{F0: 255, F1: 484}
                                                var v94 uint64 = (v3 * 27)
                                                if (65535 <= 260) {
                                                    var v95 float32 = float32(0.0)
                                                    if false {
                                                        if (((false && false) || false) && true) {
                                                            fmt.Println(v3)
                                                            var v96 string = "a"
                                                            _ = v96
                                                        } else {
                                                            switch (v90 * fi44) {
                                                            case 127:
                                                                if (!false) {
                                                                    delete(v1, "")
                                                                    v1[("a" + "\n\t")] = v90
                                                                }
                                                                v2 += v2
                                                                var v97 uint64 = 42
                                                                _ = v97
                                                            }
                                                            v45 = (v45 * v2)
                                                        }
                                                    } else {
                                                        fmt.Println(v3, v45, v3)
                                                    }
                                                    _ = v95
                                                } else {
                                                    v90 = -2147483648
                                                }
                                                var v98 uint16 = 42
                                                _ = v93
                                                _ = v94
                                                _ = v98
                                            }
                                            var v99 bool = false
                                            var v100 uint = 255
                                            v101 := GenS9{F0: (fi44 + fi43)}
                                            _ = v100
                                            _ = v101
                                            _ = v99
                                        }
                                        _ = v90
                                    }
                                }
                                var v102 int32 = 42
                                var v103 int32 = v102
                                switch (fi44 + fi43) {
                                case 0:
                                    switch (fi44 * 32768) {
                                    case 2147483647:
                                        var v104 float32 = 0.1
                                        _ = v104
                                    default:
                                        v105 := GenS7{F0: 65535, F1: 1}
                                        _ = v105
                                    }
                                case 9223372036854775807:
                                    v45 -= v45
                                }
                                if false {
                                    var v106 uint = 4294967295
                                    var v107 int8 = -39
                                    if true {
                                        switch "" {
                                        case "a":
                                            v3 = 1
                                        case "\n\t":
                                            if (((true || false) && false) || true) {
                                                v45 = v45
                                            }
                                        case "":
                                            _ = len(v1)
                                        default:
                                            fmt.Println(v102, fi43)
                                            var v108 int8 = (v107 * 1)
                                            if false {
                                                _ = len(v1)
                                                fmt.Println(v107)
                                            }
                                            _ = v108
                                        }
                                    }
                                    switch 1 {
                                    case 32767:
                                        v1["\n\t"] = fi43
                                    case -31:
                                        fmt.Println(v107)
                                    case -320:
                                        v109 := GenS12{F0: 9223372036854775807}
                                        _ = v109
                                    }
                                    _ = v106
                                    _ = v107
                                }
                                _ = v102
                                _ = v103
                            } else {
                                v2 = v2
                                v110 := GenS1{F0: true, F1: 1}
                                v111 := GenS10{F0: 9223372036854775807, F1: 1}
                                fmt.Println(v3, fi44)
                                _ = v110
                                _ = v111
                            }
                            var v112 uint64 = (v3 - 58)
                            _ = v112
                        } else {
                            var v113 uint32 = 2
                            _ = v113
                        }
                        var v114 uint64 = v3
                        _ = v114
                    }
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
        default:
            switch "foo bar" {
            case "":
                v115 := GenS12{F0: 848}
                _ = v115
            default:
                var v116 string = "foo bar"
                _ = v116
            }
        }
        var v117 string = ""
        _ = v117
    } else {
        if (-236 > 9223372036854775807) {
            for fi118 := 0; fi118 < 20; fi118++ {
                for fi119 := 0; fi119 < 9; fi119++ {
                    var v120 int64 = 973
                    v3 -= (v3 % ((42) | 1))
                    _ = fi119
                    _ = v120
                }
                _ = fi118
            }
            for fi121 := 0; fi121 < 14; fi121++ {
                v22 = append(v22, (fi121 & fi121))
                var v122 float32 = 3.14159
                if false {
                    var v123 uint64 = (^(v3))
                    _ = v123
                } else {
                    var v124 int32 = 572
                    v122 = 1e10
                    v2 = 3.14159
                    _ = v124
                }
                _ = fi121
                _ = v122
            }
        }
        v22 = append(v22, 690)
    }
    var v125 int = 32768
    fmt.Println(v2)
    v126 := GenS5{F0: -90, F1: 32768, F2: 27}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v125:", v125)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v125
    _ = v126
    _ = v2
    _ = v22
    _ = v3
    _ = v42
}
