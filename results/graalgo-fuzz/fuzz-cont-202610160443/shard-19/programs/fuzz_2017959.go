package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int64; F1 int; F2 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 int8; F2 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 int8; F1 int; F2 uint64 }
type GenS11 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS12 struct { F0 string; F1 uint16; F2 int }
type GenS13 struct { F0 uint64; F1 uint64; F2 int }
type GenS14 struct { F0 uint8; F1 uint16; F2 string }
func genFn0(p0 uint16, p1 uint16, p2 int) uint {
    p2 = 128
    v4 := []int{-1, p2}
    switch (-(p2)) {
    case 42:
        switch (p2 % ((p2) | 1)) {
        case 127:
            _ = len(v4)
        default:
            v5 := GenS1{F0: ("foo bar" + ("foo bar" + (("" + "\n\t") + ("\n\t" + ("\n\t" + "\n\t"))))), F1: -608}
            _ = v5
        }
    case 9223372036854775807:
        switch p2 {
        case 32767:
            switch (p2 % ((p2) | 1)) {
            case 1:
                switch 9223372036854775807 {
                case 32768:
                    v6 := GenS2{F0: 0, F1: 1}
                    _ = v6
                case -876:
                    var v7 uint8 = 0
                    _ = v7
                case 998:
                    if true {
                        var v8 float64 = (3.14159 + -2.25)
                        for fi9 := 0; fi9 < 17; fi9++ {
                            v10 := GenS1{F0: ("foo bar" + ""), F1: -1}
                            _ = fi9
                            _ = v10
                        }
                        var v11 string = ("\n\t" + "foo bar")
                        _ = v11
                        _ = v8
                    } else {
                        p2 = (-(p2))
                    }
                default:
                    v4 = append(v4, p2)
                }
            case 2147483647:
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 float32 = 1.5
                    _ = fi12
                    _ = v13
                }
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}
func genFn1(p0 string) uint64 {
    v14 := []int{32767, 0, 359}
    fmt.Println(p0)
    _ = len(v14)
    p0 = ("" + "")
    p0 = (p0 + "a")
    _ = p0
    _ = v14
    return 264
}

func main() {
    v1 := []int{0}
    v2 := []uint32{42, 2147483648}
    var v3 int16 = -1
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 bool = false
        for fi17 := 0; fi17 < 20; fi17++ {
            if v16 {
                var v18 string = ("hello" + "")
                _ = v18
            }
            _ = len(v1)
            _ = fi17
        }
        fmt.Println(v16, v16, v16)
        _ = fi15
        _ = v16
    }
    v19 := GenS3{F0: -9223372036854775808, F1: 32768, F2: 0}
    switch "a" {
    case "foo bar":
        if (0.0 != ((1e10 * ((0.0 / 3.14159) - ((-2.25 * (1e10 - 0.1)) / 1.0))) / 1.0)) {
            var v20 float32 = (3.14159 / 3.14159)
            var v21 float32 = (0.1 * ((1e10 + v20) - v20))
            v22 := GenS3{F0: -9223372036854775808, F1: 4294967295, F2: 0}
            _ = v20
            _ = v21
            _ = v22
        } else {
            switch "\n\t" {
            case "a":
                for fi23 := 0; fi23 < 16; fi23++ {
                    v1 = append(v1, 127)
                    switch fi23 {
                    case 100:
                        fmt.Println(fi23, v3)
                    }
                    if (1 < 0) {
                        if true {
                            var v24 uint = genFn0(255, 65535, 32767)
                            var v25 uint8 = 127
                            _ = v24
                            _ = v25
                        } else {
                            fmt.Println(fi23, fi23)
                            v3 = v3
                        }
                        v3 += (v3 % ((v3) | 1))
                    }
                    _ = fi23
                }
            case "foo bar":
                v26 := GenS3{F0: -9223372036854775808, F1: 9223372036854775807, F2: 421}
                _ = v26
            default:
                if false {
                    var v27 uint64 = 580
                    switch "hello" {
                    case "a":
                        v27 = (^(v27))
                    case "foo bar":
                        if (-2.25 > (1e10 + 0.0)) {
                            if (((float64(0.1) <= (0.0 * 0.0)) && true) || true) {
                                v28 := GenS4{F0: 2147483647}
                                v29 := GenS5{F0: 32, F1: 127, F2: 665}
                                var v30 uint64 = (v27 + v27)
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                if (true || (42 <= 2147483647)) {
                                    var v31 int16 = 32767
                                    var v32 int64 = 0
                                    _ = v31
                                    _ = v32
                                }
                                v27 = 0
                                if true {
                                    var v33 int64 = 0
                                    switch (("\n\t" + "hello") + "hello") {
                                    case "":
                                        var v34 uint = 4294967295
                                        _ = v34
                                    }
                                    _ = v33
                                }
                            }
                        }
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        var v35 float32 = 0.0
                        _ = v35
                    }
                    if (v27 != v27) {
                        v27 += v27
                        for fi36 := 0; fi36 < 16; fi36++ {
                            for fi37 := 0; fi37 < 13; fi37++ {
                                fmt.Println(fi36, fi37, v3)
                                if (-9223372036854775808 > -108) {
                                    v27 = 301
                                    if (255 <= 255) {
                                        v38 := GenS5{F0: 2147483647, F1: -63, F2: 294}
                                        v3 += (v3 + int16(fi36))
                                        _ = v38
                                    }
                                    _ = len(v1)
                                } else {
                                    var v39 int = fi37
                                    var v40 uint = 42
                                    _ = v39
                                    _ = v40
                                }
                                _ = fi37
                            }
                            _ = fi36
                        }
                    } else {
                        fmt.Println(v3, v27)
                        if (304 < 791) {
                            var v41 uint8 = 0
                            _ = v41
                        } else {
                            switch "hello" {
                            case "":
                                v27 = (v27 % ((v27) | 1))
                            case "foo bar":
                                if ((v27 | (v27 & (v27 | genFn1("hello")))) < (^(v27))) {
                                    if ((v3 - v3) != -635) {
                                        var v42 int8 = 127
                                        var v43 int16 = v3
                                        var v44 int = 128
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    for fi45 := 0; fi45 < 7; fi45++ {
                                        v27 = (v27 << ((18446744073709551615) & 63))
                                        if (!true) {
                                            v46 := GenS6{F0: -128}
                                            var v47 int8 = 1
                                            var v48 string = "foo bar"
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                        fmt.Println(v27, fi45, v3)
                                        _ = fi45
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            default:
                                for fi49 := 0; fi49 < 10; fi49++ {
                                    _ = len(v2)
                                    switch 4294967295 {
                                    case 32767:
                                        fmt.Println(v3, fi49, fi49)
                                    case -1:
                                        var v50 uint = 255
                                        _ = v50
                                    case 32768:
                                        var v51 int64 = 42
                                        _ = v51
                                    default:
                                        v3 -= (-(v3))
                                    }
                                    _ = fi49
                                }
                            }
                            if (!false) {
                                switch "foo bar" {
                                case "a":
                                    var v52 uint16 = 42
                                    _ = v52
                                case "foo bar":
                                    switch "hello" {
                                    case "":
                                        switch ("foo bar" + "\n\t") {
                                        case "hello":
                                            v53 := GenS7{F0: 2147483647, F1: ((2147483647 > 2147483647) && (v27 > v27))}
                                            _ = v53
                                        case "foo bar":
                                            if len(v2) > 0 {
                                                v2[0] = 0
                                            }
                                        case "a":
                                            if true {
                                                v54 := GenS7{F0: -9223372036854775808, F1: true}
                                                switch 616 {
                                                case -649:
                                                    var v55 int = 32767
                                                    _ = v55
                                                case 1:
                                                    var v56 bool = (("" + (("foo bar" + "\n\t") + (("a" + "hello") + "foo bar"))) != "")
                                                    _ = v56
                                                }
                                                if (true || ((((0.1 * (-2.25 * 3.14159)) + 0.0) < 1e10) || (-9223372036854775808 >= 4294967295))) {
                                                    fmt.Println(v3, v3)
                                                } else {
                                                    switch 2147483647 {
                                                    case -1:
                                                        var v57 int32 = -234
                                                        _ = v57
                                                    case -2147483648:
                                                        v27 -= (v27 >> ((0) & 63))
                                                    case 912:
                                                        for fi58 := 0; fi58 < 5; fi58++ {
                                                            var v59 float32 = (float32(float64(0.1)) + 3.14159)
                                                            _ = fi58
                                                            _ = v59
                                                        }
                                                    default:
                                                        for fi60 := 0; fi60 < 18; fi60++ {
                                                            v61 := GenS8{F0: (42 >= -677)}
                                                            for fi62 := 0; fi62 < 1; fi62++ {
                                                                var v63 uint = genFn0(2, 42, (fi60 | fi60))
                                                                _ = fi62
                                                                _ = v63
                                                            }
                                                            v64 := GenS7{F0: (^(fi60)), F1: true}
                                                            _ = fi60
                                                            _ = v61
                                                            _ = v64
                                                        }
                                                    }
                                                    v65 := GenS9{F0: ((true && (2 == 11)) && true), F1: -9223372036854775808}
                                                    v27 = (v27 / (((v27 - 335)) | 1))
                                                    _ = v65
                                                }
                                                _ = v54
                                            }
                                        default:
                                            v27 = (v27 * 4294967295)
                                        }
                                    case "foo bar":
                                        var v66 uint32 = 42
                                        _ = v66
                                    default:
                                        if true {
                                            var v67 int = 1
                                            if (2 <= 2) {
                                                var v68 uint16 = 42
                                                _ = v68
                                            }
                                            _ = v67
                                        } else {
                                            v27 += 1
                                            _ = len(v2)
                                        }
                                    }
                                case "":
                                    var v69 uint8 = 42
                                    _ = v69
                                default:
                                    if ((-222 < 483) && (3.14159 > 0.1)) {
                                        for fi70 := 0; fi70 < 7; fi70++ {
                                            v71 := GenS10{F0: 71, F1: -9223372036854775808, F2: genFn1("foo bar")}
                                            _ = fi70
                                            _ = v71
                                        }
                                        fmt.Println(v3, v27)
                                        var v72 int16 = -871
                                        _ = v72
                                    }
                                }
                            }
                        }
                    }
                    _ = v27
                }
            }
            if len(v2) > 0 {
                v2[0] = 42
            }
        }
    case "\n\t":
        fmt.Println(v3)
    default:
        var v73 int = 42
        _ = v73
    }
    v74 := GenS6{F0: 5}
    if ((-973 == -703) || true) {
        fmt.Println(v3)
        if (!true) {
            v75 := GenS11{F0: 52, F1: 157, F2: 42}
            if (!(9223372036854775807 <= -812)) {
                switch -2147483648 {
                case 2147483647:
                    var v76 float32 = 0.0
                    _ = v76
                case 4294967295:
                    v3 = (v3 * -1)
                case 167:
                    v77 := GenS11{F0: 186, F1: 120, F2: 205}
                    _ = v77
                }
                v1 = append(v1, 0)
                _ = len(v1)
            }
            for fi78 := 0; fi78 < 8; fi78++ {
                var v79 float64 = 0.1
                v80 := GenS10{F0: 42, F1: (^(fi78)), F2: 4294967295}
                _ = fi78
                _ = v79
                _ = v80
            }
            _ = v75
        } else {
            switch -9223372036854775808 {
            case -1:
                if false {
                    v81 := GenS12{F0: ("hello" + "hello"), F1: 953, F2: -689}
                    v3 -= -249
                    v82 := GenS13{F0: 286, F1: 4294967295, F2: 42}
                    _ = v81
                    _ = v82
                } else {
                    if true {
                        var v83 string = "hello"
                        _ = v83
                    }
                }
            case 32768:
                var v84 int8 = 0
                _ = v84
            case 2147483647:
                var v85 float64 = 1.5
                _ = v85
            default:
                fmt.Println(v3)
            }
            for fi86 := 0; fi86 < 8; fi86++ {
                fmt.Println(v3)
                if (!(!true)) {
                    if (976 >= 2) {
                        v3 = 32767
                    }
                    var v87 bool = (-905 <= 581)
                    var v88 float32 = (float32(float64(((0.1 + 1.5) * 0.1))) + (-2.25 + ((float32((0.0 / 3.14159)) / (0.1 / 1e10)) * (0.0 + -2.25))))
                    _ = v87
                    _ = v88
                } else {
                    if true {
                        v3 = 1000
                    } else {
                        v89 := GenS11{F0: 2, F1: 0, F2: 1}
                        _ = v89
                    }
                    if true {
                        switch (fi86 & 42) {
                        case -1:
                            for fi90 := 0; fi90 < 6; fi90++ {
                                var v91 int64 = 100
                                v92 := GenS8{F0: true}
                                _ = fi90
                                _ = v91
                                _ = v92
                            }
                        default:
                            v93 := GenS14{F0: 128, F1: 32767, F2: ""}
                            _ = v93
                        }
                    } else {
                        v94 := GenS6{F0: 79}
                        v3 -= v3
                        _ = v94
                    }
                    v95 := GenS2{F0: 4294967295, F1: -1}
                    _ = v95
                }
                var v96 uint16 = 308
                _ = fi86
                _ = v96
            }
        }
    } else {
        var v97 int16 = (v3 + v3)
        switch ("foo bar" + "\n\t") {
        case "a":
            var v98 uint = 0
            _ = v98
        }
        _ = v97
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        v3 = 32767
    default:
        v99 := GenS7{F0: -246, F1: false}
        _ = v99
    }
    _ = len(v2)
    if (((!(!false)) || (-2147483648 != 794)) || true) {
        v100 := GenS14{F0: 178, F1: 0, F2: (("\n\t" + ((("" + "a") + "a") + "hello")) + "foo bar")}
        v3 = v3
        var v101 uint64 = 83
        _ = v100
        _ = v101
    } else {
        if (!true) {
            var v102 int32 = -2147483648
            for fi103 := 0; fi103 < 2; fi103++ {
                v3 = v3
                fmt.Println(v102)
                v3 = v3
                _ = fi103
            }
            _ = v102
        }
        for fi104 := 0; fi104 < 3; fi104++ {
            var v105 uint16 = 255
            var v106 int8 = -20
            var v107 uint8 = 128
            _ = fi104
            _ = v105
            _ = v106
            _ = v107
        }
    }
    var v108 int8 = 100
    for fi109 := 0; fi109 < 2; fi109++ {
        if false {
            var v110 bool = false
            var v111 int = (fi109 + -231)
            _ = v110
            _ = v111
        }
        _ = fi109
    }
    v3 = 1
    fmt.Println(v108)
    var v112 int32 = 42
    if false {
        var v113 bool = true
        _ = v113
    }
    var v114 int64 = 42
    _ = len(v1)
    v115 := GenS12{F0: "", F1: 32767, F2: 392}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v108:", v108)
    fmt.Println("v112:", v112)
    fmt.Println("v114:", v114)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v108
    _ = v112
    _ = v114
    _ = v115
    _ = v19
    _ = v2
    _ = v3
    _ = v74
}
