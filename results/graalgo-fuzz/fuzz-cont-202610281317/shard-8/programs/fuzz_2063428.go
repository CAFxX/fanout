package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int64; F2 int16 }
type GenS6 struct { F0 int16; F1 int32; F2 uint16 }
type GenS7 struct { F0 int32; F1 int8; F2 uint16 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint16; F1 uint; F2 int32 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 int64; F1 int16; F2 uint8 }
type GenS12 struct { F0 uint16 }

func main() {
    var v1 float64 = float64(0.1)
    var v2 string = "\n\t"
    var v3 int8 = -93
    var v4 int32 = -406
    v5 := []int{712, 42}
    for fi6 := 0; fi6 < 15; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(fi6)
        if true {
            v7 := GenS1{F0: -128, F1: 4294967295, F2: -1}
            _ = v7
        } else {
            var v8 int16 = -348
            v4 -= v4
            fmt.Println(v8)
            _ = v8
        }
        _ = fi6
    }
    switch -5 {
    case 0:
        for fi9 := 0; fi9 < 3; fi9++ {
            _ = len(v5)
            for fi10 := 0; fi10 < 13; fi10++ {
                if false {
                    v4 = -1
                    v1 = (-2.25 + 1e10)
                    fmt.Println(v1, v3)
                }
                _ = fi10
            }
            _ = fi9
        }
    case 32767:
        v5 = append(v5, 943)
    }
    switch -849 {
    case 4294967295:
        v11 := GenS1{F0: (v3 & v3), F1: 2, F2: 4294967295}
        _ = v11
    case 9223372036854775807:
        var v12 float32 = 1.5
        _ = v12
    default:
        v13 := GenS2{F0: 32767}
        _ = v13
    }
    v14 := GenS1{F0: v3, F1: 42, F2: 660}
    switch v2 {
    case "hello":
        for fi15 := 0; fi15 < 13; fi15++ {
            v4 = (-(v4))
            v2 = v2
            for fi16 := 0; fi16 < 6; fi16++ {
                if false {
                    v4 = v4
                }
                _ = fi16
            }
            _ = fi15
        }
    case "foo bar":
        for fi17 := 0; fi17 < 1; fi17++ {
            if (false || true) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (false && true) {
                    if false {
                        v3 -= (-(v3))
                        for fi18 := 0; fi18 < 8; fi18++ {
                            v19 := GenS2{F0: 42}
                            switch fi18 {
                            case -161:
                                v20 := GenS2{F0: 32767}
                                _ = v20
                            case 127:
                                if (false && ((!false) && ((-(fi18)) > (fi17 / (((fi18 & fi17)) | 1))))) {
                                    if (true || false) {
                                        var v21 float64 = v1
                                        var v22 uint16 = 42
                                        fmt.Println(v3, v2, v2)
                                        _ = v21
                                        _ = v22
                                    } else {
                                        switch 42 {
                                        case 9223372036854775807:
                                            var v23 int64 = -641
                                            _ = v23
                                        case 42:
                                            var v24 string = "\n\t"
                                            _ = v24
                                        }
                                        var v25 uint8 = 127
                                        if true {
                                            if false {
                                                var v26 float32 = (0.0 + (-2.25 * 3.14159))
                                                v27 := GenS3{F0: 128, F1: 1000}
                                                v1 = (v1 / (((v1 * v1) - ((v1 / v1) + v1)) - v1))
                                                _ = v26
                                                _ = v27
                                            } else {
                                                var v28 uint64 = 0
                                                if (true && true) {
                                                    switch v2 {
                                                    case "":
                                                        v2 = (v2 + (v2 + ("a" + "foo bar")))
                                                    case "hello":
                                                        if len(v5) > 0 {
                                                            v5[0] = -2147483648
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            v29 := GenS4{F0: 42, F1: (v25 << ((18446744073709551615) & 63))}
                                                            v25 -= v25
                                                            v28 -= v28
                                                            v30 := GenS4{F0: 255, F1: 255}
                                                            var v31 uint64 = (v28 ^ (v28 ^ 221))
                                                            v32 := GenS3{F0: 128, F1: 0}
                                                            _ = v29
                                                            _ = v30
                                                            _ = v31
                                                            _ = v32
                                                        } else {
                                                            v33 := GenS3{F0: (v25 * v25), F1: 0}
                                                            fmt.Println(v25)
                                                            _ = v33
                                                        }
                                                    default:
                                                        v4 -= v4
                                                    }
                                                }
                                                var v34 int16 = 1
                                                _ = v28
                                                _ = v34
                                            }
                                            v25 = (v25 % ((1) | 1))
                                            v5 = append(v5, (fi17 / (((fi18 | fi18)) | 1)))
                                            if len(v5) > 0 {
                                                v5[0] = int(v25)
                                            }
                                        } else {
                                            var v35 bool = (!(!((v4 + int32(fi18)) != v4)))
                                            _ = v35
                                        }
                                        _ = v25
                                    }
                                    var v36 uint = 0
                                    if len(v5) > 0 {
                                        v5[0] = 4294967295
                                    }
                                    v37 := GenS4{F0: 18446744073709551615, F1: 197}
                                    v36 = (v36 - v36)
                                    switch "hello" {
                                    case "hello":
                                        v4 -= 1
                                    }
                                    _ = v36
                                    _ = v37
                                } else {
                                    v5 = append(v5, (fi18 - (fi17 & fi17)))
                                    fmt.Println(v1, v3)
                                }
                            }
                            v38 := GenS5{F0: 263, F1: 1, F2: -278}
                            _ = fi18
                            _ = v19
                            _ = v38
                        }
                    }
                    v39 := GenS3{F0: 166, F1: 533}
                    v40 := GenS6{F0: 42, F1: v4, F2: 2}
                    _ = v39
                    _ = v40
                }
                if true {
                    v41 := GenS7{F0: (-(v4)), F1: v3, F2: 0}
                    _ = v41
                } else {
                    if false {
                        var v42 int = (fi17 | 0)
                        var v43 int = -1
                        _ = v42
                        _ = v43
                    } else {
                        var v44 int64 = 9223372036854775807
                        v45 := GenS7{F0: v4, F1: v3, F2: 1}
                        switch "\n\t" {
                        case "\n\t":
                            for fi46 := 0; fi46 < 5; fi46++ {
                                fmt.Println(v1)
                                var v47 int32 = 42
                                var v48 int64 = (v44 | 702)
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                        }
                        _ = v44
                        _ = v45
                    }
                }
            }
            for fi49 := 0; fi49 < 7; fi49++ {
                v2 = v2
                _ = fi49
            }
            v4 = int32(v3)
            _ = fi17
        }
    case "a":
        v3 -= -128
    default:
        var v50 int8 = v3
        _ = v50
    }
    fmt.Println(v3)
    if (9223372036854775807 < -9223372036854775808) {
        var v51 float32 = ((float32(((v1 / ((0.1 - v1) * v1)) - 0.1)) * (float32(v1) / 1.0)) / 0.1)
        var v52 string = v2
        v52 = v2
        _ = v51
        _ = v52
    } else {
        if (2147483648 < 4294967295) {
            var v53 int8 = v3
            if (!(!true)) {
                fmt.Println(v3)
                v54 := GenS1{F0: (v3 & v3), F1: 1, F2: 9223372036854775807}
                for fi55 := 0; fi55 < 19; fi55++ {
                    var v56 uint64 = 2
                    fmt.Println(v4, v3, fi55)
                    if false {
                        var v57 bool = (!false)
                        var v58 int8 = 100
                        _ = v57
                        _ = v58
                    } else {
                        switch (fi55 - -503) {
                        case 176:
                            var v59 int64 = -1
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 20; fi60++ {
                            var v61 int8 = 100
                            fmt.Println(fi55, fi60, v61)
                            var v62 uint8 = 2
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                    }
                    _ = fi55
                    _ = v56
                }
                _ = v54
            }
            _ = v53
        } else {
            if (!(386 != 18446744073709551615)) {
                for fi63 := 0; fi63 < 7; fi63++ {
                    v1 = v1
                    var v64 string = ("\n\t" + "hello")
                    _ = fi63
                    _ = v64
                }
            } else {
                var v65 uint16 = 57
                var v66 float32 = ((float32(v1) - (0.0 - ((1.5 / ((((0.0 - float32(((v1 / 3.14159) / -2.25))) * float32(0.0)) + 0.1) * 3.14159)) * (((0.0 * 1e10) + 0.1) / 1.0)))) - (-2.25 + 0.0))
                fmt.Println(v1, v65)
                _ = v65
                _ = v66
            }
        }
        if true {
            for fi67 := 0; fi67 < 14; fi67++ {
                for fi68 := 0; fi68 < 18; fi68++ {
                    var v69 bool = false
                    var v70 uint = 557
                    var v71 float32 = 1e10
                    _ = fi68
                    _ = v69
                    _ = v70
                    _ = v71
                }
                _ = len(v5)
                _ = fi67
            }
            if len(v5) > 0 {
                v5[0] = 0
            }
        }
    }
    if (!(1 == 1)) {
        v72 := GenS1{F0: v3, F1: 1, F2: 0}
        var v73 int8 = 42
        if true {
            v3 = 0
            switch v2 {
            case "\n\t":
                _ = len(v5)
            default:
                v2 = "a"
            }
        }
        _ = v72
        _ = v73
    } else {
        v74 := GenS8{F0: 219, F1: (v2 + "")}
        switch 2147483647 {
        case 2147483647:
            if false {
                v3 += v3
                if ((v2 >= v2) || false) {
                    if false {
                        var v75 float64 = v1
                        var v76 bool = (64 >= 130)
                        _ = v75
                        _ = v76
                    }
                    v2 = (v2 + v2)
                } else {
                    v4 -= (v4 % (((^(v4))) | 1))
                    v4 = 42
                }
            } else {
                if (false || (!(42 >= 655))) {
                    for fi77 := 0; fi77 < 17; fi77++ {
                        fmt.Println(v4, v3)
                        v4 = v4
                        _ = fi77
                    }
                } else {
                    var v78 uint64 = 0
                    v3 += (-(v3))
                    _ = v78
                }
                if (!true) {
                    switch "foo bar" {
                    case "foo bar":
                        v79 := GenS5{F0: 2147483647, F1: 9223372036854775807, F2: -1}
                        _ = v79
                    case "hello":
                        fmt.Println(v3)
                    }
                    if true {
                        v5 = append(v5, -271)
                        switch 16 {
                        case 9223372036854775807:
                            if (2147483647 == 1) {
                                switch "" {
                                case "a":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                if false {
                                    var v80 float32 = (3.14159 - 3.14159)
                                    _ = v80
                                } else {
                                    if (!(!false)) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    } else {
                                        fmt.Println(v2, v2)
                                    }
                                    v81 := GenS5{F0: 2, F1: 0, F2: 1}
                                    v4 = v4
                                    _ = v81
                                }
                            }
                        case -9223372036854775808:
                            v5 = append(v5, 1)
                        default:
                            if false {
                                switch "a" {
                                case "\n\t":
                                    v3 += 42
                                }
                                var v82 string = "foo bar"
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v82
                            } else {
                                switch -339 {
                                case 128:
                                    var v83 int8 = v3
                                    _ = v83
                                case -2147483648:
                                    var v84 int64 = 9223372036854775807
                                    _ = v84
                                }
                                var v85 uint8 = 128
                                _ = v85
                            }
                        }
                    } else {
                        fmt.Println(v1, v4, v1)
                        switch 127 {
                        case -528:
                            var v86 int8 = (v3 ^ (-(v3)))
                            _ = v86
                        }
                    }
                    switch 2147483647 {
                    case -78:
                        switch -9223372036854775808 {
                        case 32768:
                            if ((0 > v4) && false) {
                                v1 += ((v1 - 0.0) * (v1 * 1e10))
                                if (!((^(v4)) != 2147483647)) {
                                    v1 += 1.5
                                    v4 = (v4 + (v4 % ((v4) | 1)))
                                    if false {
                                        v5 = append(v5, 721)
                                        if true {
                                            var v87 uint8 = 128
                                            if true {
                                                v3 -= v3
                                                v5 = append(v5, 127)
                                            }
                                            fmt.Println(v1)
                                            _ = v87
                                        } else {
                                            switch -2147483648 {
                                            case 32767:
                                                switch v2 {
                                                case "a":
                                                    fmt.Println(v2, v1)
                                                default:
                                                    v1 = (v1 + (v1 * (v1 - float64(((0.0 - (1e10 + 0.0)) / (0.1 + ((float32(v1) - (3.14159 + (1.5 + (-2.25 - (0.0 * ((float32(((v1 / float64(1.5)) / 1e10)) / 1.0) / 0.1)))))) * (-2.25 - 1e10))))))))
                                                }
                                            case 127:
                                                if false {
                                                    v88 := GenS3{F0: 2, F1: 902}
                                                    var v89 uint32 = 978
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    v1 = v1
                                                    if len(v5) > 0 {
                                                        v5[0] = 9223372036854775807
                                                    }
                                                }
                                            }
                                            if (v4 == (v4 | (v4 * v4))) {
                                                var v90 bool = true
                                                _ = v90
                                            } else {
                                                if false {
                                                    v2 = (v2 + v2)
                                                } else {
                                                    for fi91 := 0; fi91 < 5; fi91++ {
                                                        var v92 float32 = 3.14159
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        switch ("\n\t" + "foo bar") {
                                                        case "":
                                                            switch (v2 + v2) {
                                                            case "":
                                                                var v93 float32 = (1.5 - float32((v1 - float64(((float32(0.0) + (float32(v1) + 0.1)) + v92)))))
                                                                _ = v93
                                                            case "a":
                                                                if (990 > 421) {
                                                                    v4 -= v4
                                                                    v94 := GenS1{F0: v3, F1: 42, F2: fi91}
                                                                    for fi95 := 0; fi95 < 7; fi95++ {
                                                                        if (42 > (fi91 * fi91)) {
                                                                            v92 += 1.5
                                                                            if false {
                                                                                v96 := GenS9{F0: 143, F1: 58, F2: (v4 & (v4 * 2147483647))}
                                                                                _ = v96
                                                                            }
                                                                            v1 = v1
                                                                            v2 = ("foo bar" + (("a" + v2) + v2))
                                                                            if true {
                                                                                v97 := GenS9{F0: 783, F1: 42, F2: v4}
                                                                                v98 := GenS10{F0: 463, F1: 195}
                                                                                _ = v97
                                                                                _ = v98
                                                                            } else {
                                                                                if true {
                                                                                    var v99 int = 9223372036854775807
                                                                                    if (!true) {
                                                                                        v4 -= (-(v4))
                                                                                        var v100 uint = 255
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = v99
                                                                                        }
                                                                                        _ = v100
                                                                                    }
                                                                                    v99 -= fi95
                                                                                    _ = v99
                                                                                }
                                                                                v101 := GenS11{F0: 42, F1: 42, F2: 1}
                                                                                _ = len(v5)
                                                                                _ = v101
                                                                            }
                                                                            var v102 int32 = v4
                                                                            var v103 uint16 = 1
                                                                            var v104 int16 = 730
                                                                            _ = v102
                                                                            _ = v103
                                                                            _ = v104
                                                                        }
                                                                        if ((^(fi95)) < 4294967295) {
                                                                            var v105 string = ("a" + "hello")
                                                                            _ = v105
                                                                        }
                                                                        var v106 uint64 = 1
                                                                        _ = fi95
                                                                        _ = v106
                                                                    }
                                                                    _ = v94
                                                                } else {
                                                                    _ = len(v5)
                                                                }
                                                            case "hello":
                                                                var v107 int8 = v3
                                                                _ = v107
                                                            }
                                                        case "foo bar":
                                                            v92 = (((float32(v1) / 0.1) * float32((3.14159 * 0.1))) - v92)
                                                        }
                                                        _ = fi91
                                                        _ = v92
                                                    }
                                                    for fi108 := 0; fi108 < 20; fi108++ {
                                                        v5 = append(v5, fi108)
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = fi108
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        if true {
                                            v3 = v3
                                            for fi109 := 0; fi109 < 15; fi109++ {
                                                if len(v5) > 0 {
                                                    v5[0] = fi109
                                                }
                                                _ = fi109
                                            }
                                        }
                                    }
                                }
                                var v110 int = 42
                                _ = v110
                            }
                        case -1:
                            for fi111 := 0; fi111 < 1; fi111++ {
                                fmt.Println(fi111, v4)
                                _ = fi111
                            }
                        }
                    case -447:
                        var v112 float64 = (0.1 - v1)
                        _ = v112
                    }
                } else {
                    v4 -= 42
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
        case 128:
            var v113 float32 = -2.25
            _ = v113
        case -1:
            fmt.Println(v1, v3, v2)
        default:
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
        }
        _ = v74
    }
    v2 = "\n\t"
    var v114 float32 = 1e10
    fmt.Println(v2, v3)
    v2 = (v2 + "foo bar")
    v115 := GenS12{F0: 65535}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v114
    _ = v115
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
