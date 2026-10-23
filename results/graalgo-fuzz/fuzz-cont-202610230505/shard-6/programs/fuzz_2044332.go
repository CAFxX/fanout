package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 string }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint; F2 bool }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 uint16; F2 uint8 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 int8; F1 string }
func genFn0(p0 bool) int {
    for fi6 := 0; fi6 < 3; fi6++ {
        p0 = true
        v7 := GenS1{F0: 127, F1: 127}
        _ = fi6
        _ = v7
    }
    p0 = (0 == -1)
    switch -64 {
    case 4294967295:
        fmt.Println(p0)
    case 2147483647:
        var v8 int32 = -118
        _ = v8
    default:
        fmt.Println(p0)
    }
    v9 := GenS2{F0: 2, F1: 703, F2: "foo bar"}
    _ = p0
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 uint, p1 string) uint32 {
    v10 := []int{}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v10 = append(v10, 128)
    _ = p0
    _ = p1
    _ = v10
    return uint32(p0)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 0
    v3 := make(map[string]string)
    var v4 uint = 1
    var v5 uint64 = 42
    v11 := []int{}
    if true {
        var v12 int16 = 1
        v13 := GenS1{F0: v2, F1: -1}
        v12 = v12
        _ = v12
        _ = v13
    }
    var v14 uint8 = 102
    switch -2147483648 {
    case 580:
        switch -2147483648 {
        case 32768:
            _ = len(v11)
        case 42:
            var v15 uint8 = (v14 % (((v14 | v14)) | 1))
            _ = v15
        case -1:
            switch 128 {
            case 127:
                if ((v5 == (v5 << ((255) & 63))) || false) {
                    v2 -= v2
                } else {
                    if ((false && true) && (float32(1.5) == (float32(float64(1.5)) * (3.14159 / 1.5)))) {
                        var v16 uint16 = 255
                        if (!true) {
                            switch 2147483647 {
                            case 32767:
                                for fi17 := 0; fi17 < 16; fi17++ {
                                    var v18 string = "\n\t"
                                    _ = fi17
                                    _ = v18
                                }
                            }
                        } else {
                            v16 -= (^(v16))
                        }
                        if (false || true) {
                            if (false && (v4 > v4)) {
                                switch "a" {
                                case "hello":
                                    if (42 >= 112) {
                                        var v19 uint16 = (v16 / (((v16 & (^(v16)))) | 1))
                                        for fi20 := 0; fi20 < 9; fi20++ {
                                            v19 += 440
                                            var v21 int = (-(fi20))
                                            _ = fi20
                                            _ = v21
                                        }
                                        for fi22 := 0; fi22 < 7; fi22++ {
                                            var v23 uint16 = 32767
                                            v24 := GenS2{F0: (v14 & (^(v14))), F1: (v19 << ((v4) & 63)), F2: "\n\t"}
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = v19
                                    } else {
                                        v5 = v5
                                        v25 := GenS1{F0: (v2 & (v2 * 1)), F1: 42}
                                        _ = v25
                                    }
                                case "a":
                                    delete(v1, -157)
                                case "foo bar":
                                    if (v5 <= (v5 >> ((v4) & 63))) {
                                        for fi26 := 0; fi26 < 19; fi26++ {
                                            switch ("a" + ("a" + "foo bar")) {
                                            case "hello":
                                                for fi27 := 0; fi27 < 8; fi27++ {
                                                    v28 := GenS3{F0: uint64(v14), F1: 65535}
                                                    _ = fi27
                                                    _ = v28
                                                }
                                            case "":
                                                switch ("a" + "hello") {
                                                case "a":
                                                    v14 -= v14
                                                case "hello":
                                                    fmt.Println(fi26)
                                                }
                                            default:
                                                fmt.Println(v2)
                                            }
                                            switch fi26 {
                                            case -2147483648:
                                                switch (fi26 ^ 127) {
                                                case 2147483647:
                                                    switch ("foo bar" + ("\n\t" + "hello")) {
                                                    case "foo bar":
                                                        var v29 int16 = -1
                                                        _ = v29
                                                    case "\n\t":
                                                        if true {
                                                            if ((v14 != v14) && (1 >= 597)) {
                                                                if true {
                                                                    if true {
                                                                        if true {
                                                                            v14 += v14
                                                                        } else {
                                                                            var v30 uint8 = (v14 * v14)
                                                                            var v31 uint = (v4 + 749)
                                                                            v2 += v2
                                                                            _ = v30
                                                                            _ = v31
                                                                        }
                                                                        v32 := GenS2{F0: (v14 | (v14 << ((536) & 63))), F1: (v16 + (^(v16))), F2: "foo bar"}
                                                                        v5 += v5
                                                                        _ = v32
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v33 int8 = (v2 - -1)
                                                            _ = v33
                                                        }
                                                    }
                                                case 624:
                                                    var v34 uint16 = (v16 % ((uint16(v4)) | 1))
                                                    _ = v34
                                                case 42:
                                                    if len(v11) > 0 {
                                                        v11[0] = fi26
                                                    }
                                                default:
                                                    if ((v14 & v14) == v14) {
                                                        if true {
                                                            if false {
                                                                var v35 int64 = -488
                                                                _ = v35
                                                            } else {
                                                                v14 = v14
                                                            }
                                                            for fi36 := 0; fi36 < 20; fi36++ {
                                                                v2 = (v2 % ((-1) | 1))
                                                                var v37 string = "foo bar"
                                                                switch genFn0(((42 <= 527) || false)) {
                                                                case 32767:
                                                                    v38 := GenS4{F0: uint8(v5), F1: 757, F2: (false || (true || true))}
                                                                    _ = v38
                                                                case -417:
                                                                    v39 := GenS5{F0: "", F1: -1}
                                                                    _ = v39
                                                                }
                                                                switch (fi36 / ((fi26) | 1)) {
                                                                case 898:
                                                                    v37 = v37
                                                                case 493:
                                                                    var v40 float64 = (0.0 + ((0.1 + 3.14159) - -2.25))
                                                                    _ = v40
                                                                default:
                                                                    var v41 uint64 = 361
                                                                    _ = v41
                                                                }
                                                                var v42 uint16 = uint16(v14)
                                                                _ = fi36
                                                                _ = v37
                                                                _ = v42
                                                            }
                                                            v14 -= (v14 + v14)
                                                        }
                                                        for fi43 := 0; fi43 < 16; fi43++ {
                                                            v5 -= 0
                                                            v44 := v1[(^(fi26))]
                                                            if true {
                                                                if (!(v5 >= v5)) {
                                                                    var v45 bool = false
                                                                    v4 = (v4 % ((v4) | 1))
                                                                    v1[127] = false
                                                                    switch (^(fi43)) {
                                                                    case 42:
                                                                        if v45 {
                                                                            fmt.Println(v2, v14, v45)
                                                                            switch fi26 {
                                                                            case 550:
                                                                                v16 += v16
                                                                            case 348:
                                                                                v46 := GenS6{F0: 0, F1: 127}
                                                                                v4 -= v4
                                                                                _ = v46
                                                                            case 943:
                                                                                v4 += v4
                                                                            }
                                                                            v4 = v4
                                                                            v14 = (v14 | uint8(v16))
                                                                        } else {
                                                                            var v47 float64 = float64(((float32(float64((-2.25 - ((1e10 - 0.0) * 0.1)))) / 0.1) + 0.0))
                                                                            fmt.Println(fi43)
                                                                            _ = v47
                                                                        }
                                                                        var v48 int = fi43
                                                                        var v49 int = v48
                                                                        var v50 int8 = int8(v4)
                                                                        v16 = v16
                                                                        v16 = 901
                                                                        _ = v48
                                                                        _ = v49
                                                                        _ = v50
                                                                    case 428:
                                                                        v1[int(v5)] = v44
                                                                        var v51 int8 = v2
                                                                        fmt.Println(v45, v45)
                                                                        _ = v51
                                                                    case -1:
                                                                        var v52 int32 = -1
                                                                        v16 = v16
                                                                        v53 := GenS7{F0: v2}
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                    var v54 int8 = (v2 | v2)
                                                                    var v55 string = ("\n\t" + "\n\t")
                                                                    switch ((v55 + v55) + "a") {
                                                                    case "":
                                                                        if ((v16 + (v16 % ((v16) | 1))) < v16) {
                                                                            fmt.Println(fi43)
                                                                        }
                                                                    case "foo bar":
                                                                        switch 664 {
                                                                        case 127:
                                                                            if false {
                                                                                var v56 float32 = 3.14159
                                                                                var v57 uint8 = (^(v14))
                                                                                if len(v11) > 0 {
                                                                                    v11[0] = -9223372036854775808
                                                                                }
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                if ((v16 * 255) <= v16) {
                                                                                    v58 := GenS4{F0: (v14 * (^(v14))), F1: (v4 >> ((407) & 63)), F2: true}
                                                                                    if v44 {
                                                                                        var v59 uint = (v4 + (v4 | (v4 - (v4 ^ 4294967295))))
                                                                                        switch ((v55 + v55) + ((v55 + "\n\t") + v55)) {
                                                                                        case "foo bar":
                                                                                            if (-444 != 609) {
                                                                                                var v60 uint16 = (^(v16))
                                                                                                if false {
                                                                                                    v44 = (!true)
                                                                                                    if v45 {
                                                                                                        if false {
                                                                                                            var v61 bool = v44
                                                                                                            var v62 int = fi26
                                                                                                            if (v54 <= v54) {
                                                                                                                switch v55 {
                                                                                                                case "":
                                                                                                                    var v63 uint16 = v16
                                                                                                                    _ = v63
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v61
                                                                                                            _ = v62
                                                                                                        } else {
                                                                                                            _ = len(v11)
                                                                                                            if false {
                                                                                                                var v64 int = (^(fi26))
                                                                                                                var v65 float64 = 3.14159
                                                                                                                _ = v64
                                                                                                                _ = v65
                                                                                                            } else {
                                                                                                                v66 := GenS6{F0: 42, F1: v54}
                                                                                                                var v67 int64 = -352
                                                                                                                var v68 int = fi43
                                                                                                                _ = v66
                                                                                                                _ = v67
                                                                                                                _ = v68
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v11) >= 2 {
                                                                                                            v11 = v11[1:len(v11)]
                                                                                                        }
                                                                                                        var v69 uint16 = v16
                                                                                                        _ = v69
                                                                                                    }
                                                                                                } else {
                                                                                                    v16 -= (v60 << (((v4 / ((4294967295) | 1))) & 63))
                                                                                                }
                                                                                                v3[v55] = v55
                                                                                                _ = v60
                                                                                            }
                                                                                        case "a":
                                                                                            v4 = (v4 * (v59 | v59))
                                                                                        case "hello":
                                                                                            var v70 int64 = 42
                                                                                            _ = v70
                                                                                        }
                                                                                        v14 += 128
                                                                                        var v71 bool = v45
                                                                                        _ = v59
                                                                                        _ = v71
                                                                                    } else {
                                                                                        v45 = v45
                                                                                        if ((fi26 ^ (fi43 | (fi43 | -9223372036854775808))) != int(v4)) {
                                                                                            fmt.Println(v2, v16)
                                                                                        } else {
                                                                                            switch v55 {
                                                                                            case "a":
                                                                                                v16 -= v16
                                                                                            }
                                                                                            v72 := GenS4{F0: v14, F1: 87, F2: v45}
                                                                                            fmt.Println(v55, v45, v44)
                                                                                            _ = v72
                                                                                        }
                                                                                        var v73 int16 = 508
                                                                                        _ = v73
                                                                                    }
                                                                                    _ = v58
                                                                                }
                                                                            }
                                                                        case 128:
                                                                            if len(v11) >= 2 {
                                                                                v11 = v11[1:len(v11)]
                                                                            }
                                                                        default:
                                                                            if len(v11) > 0 {
                                                                                v11[0] = fi26
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        var v74 bool = v44
                                                                        _ = v74
                                                                    default:
                                                                        var v75 string = v55
                                                                        _ = v75
                                                                    }
                                                                    _ = v45
                                                                    _ = v54
                                                                    _ = v55
                                                                } else {
                                                                    v76 := v1[genFn0((0 >= -630))]
                                                                    _ = v76
                                                                }
                                                                v16 += v16
                                                            }
                                                            if len(v11) > 0 {
                                                                v11[0] = 0
                                                            }
                                                            var v77 int32 = 2147483647
                                                            _ = fi43
                                                            _ = v44
                                                            _ = v77
                                                        }
                                                        v2 = v2
                                                    }
                                                }
                                            default:
                                                fmt.Println(v5, v16)
                                            }
                                            _ = fi26
                                        }
                                        for fi78 := 0; fi78 < 3; fi78++ {
                                            v14 += 97
                                            _ = fi78
                                        }
                                        var v79 int = 762
                                        _ = v79
                                    }
                                default:
                                    var v80 float64 = 1e10
                                    _ = v80
                                }
                                v81 := GenS8{F0: "a", F1: 0, F2: 1}
                                if true {
                                    for fi82 := 0; fi82 < 16; fi82++ {
                                        var v83 float32 = 1.5
                                        fmt.Println(fi82, v2)
                                        _ = fi82
                                        _ = v83
                                    }
                                    var v84 string = ("" + "a")
                                    _ = v84
                                } else {
                                    var v85 int64 = 42
                                    _ = v85
                                }
                                _ = v81
                            }
                            if false {
                                switch "foo bar" {
                                case "hello":
                                    if len(v11) > 0 {
                                        v11[0] = 39
                                    }
                                case "\n\t":
                                    if true {
                                        v16 = v16
                                        if true {
                                            switch 22 {
                                            case 0:
                                                if len(v11) > 0 {
                                                    v11[0] = 10
                                                }
                                            case -1:
                                                if true {
                                                    v14 += 1
                                                } else {
                                                    fmt.Println(v16, v5)
                                                    var v86 float64 = float64(0.1)
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(v14)
                                                    default:
                                                        for fi87 := 0; fi87 < 1; fi87++ {
                                                            for fi88 := 0; fi88 < 18; fi88++ {
                                                                switch (^(fi87)) {
                                                                case 32768:
                                                                    v89 := GenS8{F0: "\n\t", F1: v16, F2: v14}
                                                                    _ = v89
                                                                case -766:
                                                                    if len(v11) >= 2 {
                                                                        v11 = v11[1:len(v11)]
                                                                    }
                                                                case 1:
                                                                    var v90 uint16 = 65535
                                                                    _ = v90
                                                                }
                                                                _ = fi88
                                                            }
                                                            _ = fi87
                                                        }
                                                    }
                                                    _ = v86
                                                }
                                            }
                                        } else {
                                            v5 = v5
                                            switch 128 {
                                            case 4294967295:
                                                switch -2147483648 {
                                                case 4294967295:
                                                    for fi91 := 0; fi91 < 8; fi91++ {
                                                        v4 = 65535
                                                        v92 := v1[(fi91 & fi91)]
                                                        _ = fi91
                                                        _ = v92
                                                    }
                                                case -180:
                                                    if true {
                                                        v4 -= v4
                                                    }
                                                }
                                            case 226:
                                                var v93 uint16 = (v16 ^ v16)
                                                _ = v93
                                            case 42:
                                                var v94 bool = (("" < "") && true)
                                                _ = v94
                                            default:
                                                var v95 float64 = 3.14159
                                                _ = v95
                                            }
                                            var v96 int8 = -1
                                            _ = v96
                                        }
                                    }
                                default:
                                    v3["\n\t"] = "foo bar"
                                }
                            }
                            fmt.Println(v5, v16)
                        }
                        _ = v16
                    }
                    v97 := GenS5{F0: "hello", F1: 224}
                    if false {
                        var v98 float32 = (-2.25 - (1.5 + -2.25))
                        _ = v98
                    }
                    _ = v97
                }
            case -828:
                v4 = v4
            default:
                for fi99 := 0; fi99 < 13; fi99++ {
                    for fi100 := 0; fi100 < 18; fi100++ {
                        switch "foo bar" {
                        case "\n\t":
                            var v101 int = fi99
                            _ = v101
                        }
                        v5 = v5
                        _ = fi100
                    }
                    v102 := GenS5{F0: "", F1: 32767}
                    v5 += (v5 % ((v5) | 1))
                    _ = fi99
                    _ = v102
                }
            }
        default:
            fmt.Println(v4, v2)
        }
    case -1:
        v103 := GenS6{F0: -645, F1: 100}
        _ = v103
    case 950:
        var v104 int16 = -1
        _ = v104
    default:
        var v105 uint = (v4 % (((v4 << ((v4) & 63))) | 1))
        _ = v105
    }
    fmt.Println(v4, v4)
    v4 += v4
    fmt.Println(v4, v5, v2)
    fmt.Println(v4)
    switch -425 {
    case -2147483648:
        v14 = 255
    case 4294967295:
        if true {
            if true {
                if (true || (9223372036854775807 > 0)) {
                    var v106 uint64 = (v5 - (^(v5)))
                    _ = v106
                } else {
                    v107 := GenS9{F0: 250, F1: 56}
                    _ = v107
                }
            } else {
                if false {
                    if (1 != -266) {
                        if true {
                            var v108 float32 = (-2.25 * (3.14159 / 1.5))
                            var v109 uint16 = 1
                            _ = v108
                            _ = v109
                        }
                        fmt.Println(v4, v14, v2)
                    } else {
                        for fi110 := 0; fi110 < 14; fi110++ {
                            switch "hello" {
                            case "\n\t":
                                var v111 bool = (v2 > (v2 % (((v2 << ((v4) & 63))) | 1)))
                                _ = v111
                            default:
                                v1[(^(fi110))] = (!(float64(1.5) != 0.1))
                            }
                            if (!false) {
                                v112 := v1[(fi110 >> ((v4) & 63))]
                                _ = v112
                            } else {
                                fmt.Println(v5, fi110, v14)
                                fmt.Println(v2)
                                var v113 int64 = 0
                                _ = v113
                            }
                            v114 := GenS10{F0: 42, F1: (("foo bar" + ("a" + "a")) + ("foo bar" + ""))}
                            _ = fi110
                            _ = v114
                        }
                        v115 := GenS6{F0: 42, F1: v2}
                        _ = v115
                    }
                }
            }
        } else {
            var v116 float64 = (1.5 + 3.14159)
            _ = v116
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
