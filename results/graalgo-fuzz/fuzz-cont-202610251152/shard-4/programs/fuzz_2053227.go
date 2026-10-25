package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 int16; F2 int8 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 int16; F1 uint32; F2 int32 }
type GenS9 struct { F0 uint; F1 int64; F2 int32 }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 int8; F1 uint }

func main() {
    var v1 uint64 = 243
    v2 := []int{270, 1, -774, 653}
    v3 := make(map[int]uint32)
    v4 := []string{("hello" + "")}
    var v5 float64 = 0.0
    if true {
        v2 = append(v2, -2147483648)
        if (2147483647 <= 0) {
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 string = "\n\t"
                v7 = v7
                _ = fi6
                _ = v7
            }
        } else {
            var v8 uint16 = 0
            for fi9 := 0; fi9 < 16; fi9++ {
                switch "" {
                case "hello":
                    v10 := GenS1{F0: ("foo bar" + "a")}
                    _ = v10
                case "\n\t":
                    if true {
                        switch fi9 {
                        case 32768:
                            for fi11 := 0; fi11 < 11; fi11++ {
                                fmt.Println(fi11, v8)
                                var v12 string = ""
                                if true {
                                    v12 = ""
                                    if false {
                                        var v13 int = 32767
                                        var v14 int64 = -9223372036854775808
                                        v15 := GenS2{F0: -1, F1: 0, F2: ""}
                                        _ = v13
                                        _ = v14
                                        _ = v15
                                    } else {
                                        var v16 uint64 = 42
                                        v16 += v1
                                        var v17 float64 = ((1.5 + v5) / float64((1.5 + 0.1)))
                                        _ = v16
                                        _ = v17
                                    }
                                    fmt.Println(v1, v8)
                                }
                                _ = fi11
                                _ = v12
                            }
                        default:
                            fmt.Println(v1, v8, fi9)
                        }
                        var v18 int32 = 228
                        _ = v18
                    }
                }
                _ = fi9
            }
            _ = v8
        }
        v3[-63] = 0
    }
    switch 32768 {
    case 32768:
        fmt.Println(v1, v1)
    case 32767:
        fmt.Println(v1, v5)
    case 1:
        v19 := GenS1{F0: "\n\t"}
        _ = v19
    default:
        var v20 uint16 = 65535
        _ = v20
    }
    v21 := GenS3{F0: 631}
    for fi22 := 0; fi22 < 10; fi22++ {
        _ = len(v3)
        v23 := GenS3{F0: 63}
        switch "" {
        case "\n\t":
            fmt.Println(v1, v1, v5)
        default:
            var v24 uint16 = 0
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    switch ("hello" + ("hello" + ("\n\t" + "foo bar"))) {
    case "hello":
        v25 := GenS4{F0: 415, F1: -32768}
        _ = v25
    case "a":
        var v26 uint8 = 125
        _ = v26
    case "\n\t":
        v1 += (v1 << ((65535) & 63))
    default:
        _ = len(v4)
    }
    v5 -= v5
    v27 := GenS2{F0: 127, F1: 501, F2: ""}
    for fi28 := 0; fi28 < 20; fi28++ {
        v1 -= 18446744073709551615
        for fi29 := 0; fi29 < 10; fi29++ {
            if true {
                var v30 uint8 = 255
                _ = v30
            } else {
                var v31 string = "a"
                _ = v31
            }
            v1 = v1
            fmt.Println(v1)
            _ = fi29
        }
        _ = fi28
    }
    v1 -= v1
    for fi32 := 0; fi32 < 18; fi32++ {
        v33 := GenS3{F0: 23}
        v34 := GenS4{F0: 32767, F1: -32768}
        if true {
            for fi35 := 0; fi35 < 4; fi35++ {
                switch fi32 {
                case -65:
                    v36 := v3[32768]
                    _ = v36
                }
                delete(v3, -1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi35
            }
        } else {
            var v37 int = -985
            if ((!false) || (-54 != -39)) {
                v1 = (^(v1))
            } else {
                v37 = fi32
                if (255 != 1) {
                    v38 := v3[(v37 / (((v37 & fi32)) | 1))]
                    v39 := v3[-2147483648]
                    switch (fi32 << ((4294967295) & 63)) {
                    case -697:
                        switch fi32 {
                        case 9223372036854775807:
                            var v40 uint32 = (^(v39))
                            _ = v40
                        case 0:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 32768:
                            v37 = (fi32 + (fi32 & fi32))
                        default:
                            v39 = v39
                        }
                    }
                    _ = v38
                    _ = v39
                } else {
                    for fi41 := 0; fi41 < 13; fi41++ {
                        v1 = (v1 - v1)
                        _ = fi41
                    }
                    v1 -= v1
                    v5 += (v5 * v5)
                }
                v1 += (v1 >> ((700) & 63))
            }
            v37 -= v37
            _ = v37
        }
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi42 := 0; fi42 < 5; fi42++ {
        v43 := GenS5{F0: -5}
        switch (fi42 | (fi42 / ((fi42) | 1))) {
        case 1:
            v44 := GenS6{F0: 2, F1: 42, F2: -128}
            _ = v44
        case -936:
            var v45 uint8 = 2
            _ = v45
        case 794:
            v5 = v5
        }
        _ = fi42
        _ = v43
    }
    if true {
        if true {
            var v46 uint = 4294967295
            var v47 uint16 = 2
            _ = v46
            _ = v47
        } else {
            delete(v3, 0)
            var v48 int8 = 100
            if (128 <= 127) {
                v5 -= 0.1
            } else {
                var v49 int32 = 42
                v5 = v5
                switch "hello" {
                case "\n\t":
                    for fi50 := 0; fi50 < 18; fi50++ {
                        if (!true) {
                            for fi51 := 0; fi51 < 19; fi51++ {
                                switch -261 {
                                case 0:
                                    v52 := v3[fi51]
                                    _ = v52
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    v53 := GenS6{F0: 133, F1: 0, F2: (v48 ^ 0)}
                                    _ = v53
                                case "":
                                    var v54 uint16 = 255
                                    _ = v54
                                case "hello":
                                    fmt.Println(v5, v48, fi51)
                                    var v55 int64 = -978
                                    _ = v55
                                default:
                                    v48 = v48
                                }
                                v49 = (-(v49))
                                _ = fi51
                            }
                            for fi56 := 0; fi56 < 5; fi56++ {
                                if (false || (878 != 2)) {
                                    var v57 int32 = 542
                                    switch 125 {
                                    case -9:
                                        var v58 uint32 = 2
                                        v1 = v1
                                        switch fi50 {
                                        case 6:
                                            switch ("foo bar" + "hello") {
                                            case "\n\t":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v59 bool = (false || ((false && true) || (true || false)))
                                                _ = v59
                                            }
                                        default:
                                            if (true || true) {
                                                v60 := GenS3{F0: 32767}
                                                _ = v60
                                            }
                                        }
                                        _ = v58
                                    case 32767:
                                        switch 128 {
                                        case 42:
                                            v2 = append(v2, 32768)
                                        }
                                    default:
                                        v61 := GenS5{F0: v48}
                                        var v62 uint64 = v1
                                        var v63 uint16 = 902
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = v57
                                }
                                if true {
                                    if (-589 >= 0) {
                                        v5 = 1.5
                                        switch (fi56 * fi50) {
                                        case 399:
                                            fmt.Println(v5)
                                        default:
                                            fmt.Println(v5, v49, fi56)
                                        }
                                        v5 = ((float64(0.0) - v5) * v5)
                                    } else {
                                        var v64 uint32 = 2
                                        var v65 int32 = (v49 | (v49 / ((v49) | 1)))
                                        v3[128] = 0
                                        _ = v64
                                        _ = v65
                                    }
                                    _ = len(v3)
                                    _ = len(v2)
                                } else {
                                    fmt.Println(v49, v48, v5)
                                }
                                var v66 bool = true
                                if (803 >= 902) {
                                    if (("foo bar" + ("a" + "")) <= "foo bar") {
                                        var v67 bool = true
                                        fmt.Println(v5, v5, v5)
                                        _ = v67
                                    }
                                } else {
                                    v68 := v3[(^(fi56))]
                                    _ = v68
                                }
                                _ = fi56
                                _ = v66
                            }
                        } else {
                            v5 -= (1.5 * float64(0.0))
                            switch "foo bar" {
                            case "\n\t":
                                var v69 string = ("\n\t" + "a")
                                _ = v69
                            case "a":
                                v3[fi50] = 2
                            default:
                                v70 := v3[fi50]
                                _ = v70
                            }
                        }
                        fmt.Println(v48, v1, v48)
                        for fi71 := 0; fi71 < 15; fi71++ {
                            switch (fi71 / ((128) | 1)) {
                            case -447:
                                var v72 uint = 0
                                _ = v72
                            default:
                                var v73 uint64 = (^(v1))
                                _ = v73
                            }
                            v3[(fi71 / ((fi71) | 1))] = 0
                            v74 := GenS7{F0: "\n\t", F1: 0}
                            _ = fi71
                            _ = v74
                        }
                        _ = fi50
                    }
                case "a":
                    var v75 uint = 2
                    _ = v75
                }
                _ = v49
            }
            _ = v48
        }
        if true {
            if len(v2) > 0 {
                v2[0] = -553
            }
            switch ("" + "hello") {
            case "hello":
                var v76 uint64 = 18446744073709551615
                _ = v76
            default:
                v5 = (v5 * v5)
            }
            v1 += 2
        }
        var v77 int = 9223372036854775807
        _ = v77
    } else {
        v5 -= (v5 + 1.5)
    }
    fmt.Println(v1, v5)
    if (true && (2 == 203)) {
        var v78 int8 = 127
        if false {
            v5 = v5
            var v79 uint8 = 0
            _ = v79
        }
        _ = v78
    }
    if (128 != 42) {
        if false {
            v80 := GenS5{F0: 1}
            var v81 string = "foo bar"
            _ = v80
            _ = v81
        }
        for fi82 := 0; fi82 < 1; fi82++ {
            v5 += 0.0
            if false {
                var v83 float32 = 0.0
                _ = v83
            } else {
                if (1 != -1) {
                    if false {
                        for fi84 := 0; fi84 < 14; fi84++ {
                            var v85 uint64 = (v1 + v1)
                            v1 += v1
                            var v86 float32 = 0.0
                            _ = fi84
                            _ = v85
                            _ = v86
                        }
                        if false {
                            switch 42 {
                            case -810:
                                delete(v3, (-(fi82)))
                            default:
                                v87 := GenS8{F0: 1, F1: 370, F2: 0}
                                _ = v87
                            }
                            v5 = (0.0 + v5)
                        }
                    } else {
                        for fi88 := 0; fi88 < 20; fi88++ {
                            v1 -= v1
                            var v89 uint = 0
                            v5 = ((v5 + (v5 / 1e10)) + 0.0)
                            _ = fi88
                            _ = v89
                        }
                        _ = len(v3)
                        v1 = (^(v1))
                    }
                    v1 += (v1 / (((v1 / ((v1) | 1))) | 1))
                } else {
                    switch 4294967295 {
                    case -1000:
                        v1 = (v1 / ((v1) | 1))
                    }
                    if (!false) {
                        v1 = v1
                        delete(v3, fi82)
                        v1 = v1
                    }
                }
                if (!true) {
                    v90 := GenS4{F0: 0, F1: 42}
                    for fi91 := 0; fi91 < 4; fi91++ {
                        if false {
                            if true {
                                v5 = v5
                                if len(v4) > 0 {
                                    v4[0] = ("\n\t" + "hello")
                                }
                                fmt.Println(fi82, v1, v1)
                            } else {
                                v92 := GenS9{F0: 4294967295, F1: -690, F2: 42}
                                v5 = v5
                                _ = v92
                            }
                            v93 := GenS1{F0: ((("foo bar" + "a") + "a") + "foo bar")}
                            v5 -= 0.1
                            _ = v93
                        } else {
                            if ("hello" < "hello") {
                                v4 = append(v4, "\n\t")
                                v5 = (float64((0.0 + (((3.14159 / (1.5 + 3.14159)) - ((0.0 - 0.1) + 1.5)) * 0.0))) + (-2.25 / v5))
                                v94 := GenS10{F0: 1, F1: 1}
                                _ = v94
                            }
                        }
                        if ((fi82 + -9223372036854775808) <= fi82) {
                            v1 += v1
                            v95 := GenS11{F0: "foo bar", F1: 127}
                            switch "hello" {
                            case "a":
                                if false {
                                    switch "foo bar" {
                                    case "\n\t":
                                        if true {
                                            var v96 uint64 = (v1 >> ((753) & 63))
                                            _ = v96
                                        }
                                    case "hello":
                                        var v97 int64 = 1
                                        var v98 uint8 = 0
                                        _ = v97
                                        _ = v98
                                    default:
                                        if true {
                                            var v99 uint8 = 1
                                            _ = v99
                                        } else {
                                            v1 = v1
                                            fmt.Println(v5)
                                            switch "foo bar" {
                                            case "hello":
                                                v100 := GenS2{F0: 0, F1: 438, F2: "hello"}
                                                _ = v100
                                            case "":
                                                v1 = (v1 / ((v1) | 1))
                                            case "\n\t":
                                                fmt.Println(fi91, fi82)
                                            }
                                        }
                                    }
                                    v1 = (v1 & (v1 + v1))
                                }
                            case "":
                                v4 = append(v4, (("" + "hello") + "a"))
                            }
                            switch (fi82 % (((fi91 >> ((18446744073709551615) & 63))) | 1)) {
                            case 42:
                                var v101 int32 = 42
                                _ = v101
                            case 473:
                                v3[0] = 2
                            case 797:
                                var v102 uint32 = 65535
                                _ = v102
                            default:
                                v103 := GenS7{F0: (("a" + "") + "\n\t"), F1: 1}
                                _ = v103
                            }
                            _ = v95
                        } else {
                            v2 = append(v2, (fi91 - fi82))
                            v104 := GenS12{F0: 127, F1: 540}
                            _ = v104
                        }
                        v5 += (-2.25 - 3.14159)
                        _ = fi91
                    }
                    _ = v90
                }
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
            }
            _ = len(v3)
            _ = fi82
        }
        fmt.Println(v1)
    }
    _ = len(v3)
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 = (3.14159 - v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
