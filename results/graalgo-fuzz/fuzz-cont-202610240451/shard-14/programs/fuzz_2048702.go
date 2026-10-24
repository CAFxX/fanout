package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 uint8; F2 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int; F1 int; F2 int16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint; F1 uint32; F2 int64 }
type GenS12 struct { F0 int64; F1 uint8; F2 string }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]bool)
    var v3 uint = 364
    v4 := make(map[int]uint32)
    v5 := []uint32{4294967295, 0}
    var v6 int = 128
    switch (v6 - (v6 + (v6 ^ v6))) {
    case -303:
        if (!false) {
            v6 -= v6
            v3 = v3
        } else {
            switch (v6 & (v6 + v6)) {
            case -378:
                if ((false || false) || (-57 == 42)) {
                    v3 = v3
                    v7 := GenS1{F0: 42, F1: 42, F2: -960}
                    _ = v7
                } else {
                    var v8 int64 = -1
                    if (v1 == (1e10 * v1)) {
                        var v9 string = "a"
                        _ = v9
                    } else {
                        var v10 float32 = v1
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v10
                    }
                    if (false && ((v8 >= 0) && true)) {
                        v11 := GenS1{F0: -1, F1: 127, F2: 991}
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 int8 = 127
                            v13 -= (^(v13))
                            _ = fi12
                            _ = v13
                        }
                        v14 := GenS2{F0: "foo bar"}
                        _ = v11
                        _ = v14
                    } else {
                        var v15 bool = false
                        v16 := GenS1{F0: -1, F1: 236, F2: -1}
                        for fi17 := 0; fi17 < 1; fi17++ {
                            for fi18 := 0; fi18 < 12; fi18++ {
                                var v19 int32 = 1
                                _ = fi18
                                _ = v19
                            }
                            v1 -= -2.25
                            _ = fi17
                        }
                        _ = v15
                        _ = v16
                    }
                    _ = v8
                }
            case -1:
                v1 -= (v1 / 1e10)
            case 340:
                if (42 >= 0) {
                    if (!false) {
                        for fi20 := 0; fi20 < 14; fi20++ {
                            for fi21 := 0; fi21 < 15; fi21++ {
                                var v22 string = ("hello" + "a")
                                if true {
                                    var v23 float32 = 1.5
                                    v24 := GenS2{F0: (v22 + v22)}
                                    var v25 uint8 = 76
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    if true {
                                        v22 = "\n\t"
                                        v1 = (v1 * (0.0 * 1e10))
                                    } else {
                                        if false {
                                            var v26 uint16 = 1
                                            v5 = append(v5, 42)
                                            if (!false) {
                                                var v27 int8 = -37
                                                _ = v27
                                            } else {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v28 uint32 = 536
                                                    _ = v28
                                                case "a":
                                                    v6 = fi21
                                                case "hello":
                                                    v29 := GenS1{F0: -99, F1: 2, F2: 730}
                                                    _ = v29
                                                }
                                            }
                                            _ = v26
                                        }
                                        v22 = (v22 + (v22 + "foo bar"))
                                        v30 := GenS3{F0: (((v22 + "hello") + (v22 + "\n\t")) + (v22 + "\n\t")), F1: 843}
                                        _ = v30
                                    }
                                    v5 = append(v5, 948)
                                }
                                _ = len(v5)
                                _ = fi21
                                _ = v22
                            }
                            v31 := GenS2{F0: ("\n\t" + "\n\t")}
                            var v32 bool = (873 == 1)
                            _ = fi20
                            _ = v31
                            _ = v32
                        }
                    } else {
                        v6 = (v6 | (v6 + (v6 - v6)))
                        if false {
                            if ((float32((1.5 * (float64(v1) + -2.25))) - v1) == float32(1e10)) {
                                var v33 uint64 = 0
                                switch "foo bar" {
                                case "":
                                    var v34 uint64 = v33
                                    _ = v34
                                case "\n\t":
                                    for fi35 := 0; fi35 < 4; fi35++ {
                                        v36 := GenS4{F0: 53, F1: -886}
                                        v37 := v4[(fi35 | 42)]
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                    }
                                case "hello":
                                    var v38 uint = v3
                                    _ = v38
                                }
                                _ = v33
                            } else {
                                switch v6 {
                                case 9223372036854775807:
                                    if false {
                                        v3 += v3
                                    }
                                case 4294967295:
                                    if (v3 < v3) {
                                        switch v6 {
                                        case 738:
                                            v2[128] = (!(v6 == v6))
                                        case 32767:
                                            if false {
                                                switch "foo bar" {
                                                case "hello":
                                                    var v39 int64 = 9223372036854775807
                                                    _ = v39
                                                }
                                            } else {
                                                fmt.Println(v3)
                                                v6 += v6
                                                var v40 bool = false
                                                _ = v40
                                            }
                                        case -9223372036854775808:
                                            fmt.Println(v3, v3, v3)
                                        default:
                                            if ("" != "a") {
                                                fmt.Println(v1)
                                                var v41 int32 = -393
                                                _ = v41
                                            } else {
                                                v42 := GenS2{F0: "foo bar"}
                                                _ = v42
                                            }
                                        }
                                    }
                                }
                            }
                            v1 -= 1e10
                        } else {
                            v3 = (v3 | v3)
                            v1 -= v1
                            for fi43 := 0; fi43 < 7; fi43++ {
                                _ = len(v4)
                                var v44 float32 = v1
                                _ = fi43
                                _ = v44
                            }
                        }
                    }
                }
            default:
                v45 := GenS5{F0: "a"}
                _ = v45
            }
            var v46 int16 = 513
            var v47 uint8 = 0
            _ = v46
            _ = v47
        }
    }
    if true {
        if true {
            switch "\n\t" {
            case "hello":
                delete(v4, v6)
            case "\n\t":
                for fi48 := 0; fi48 < 10; fi48++ {
                    var v49 int64 = -1
                    _ = fi48
                    _ = v49
                }
            }
        }
        var v50 int8 = -2
        v5 = append(v5, 321)
        _ = v50
    } else {
        v51 := GenS4{F0: 127, F1: 575}
        if (!true) {
            delete(v4, (v6 - v6))
        }
        var v52 uint64 = 0
        _ = v51
        _ = v52
    }
    fmt.Println(v6)
    var v53 int = v6
    fmt.Println(v3)
    fmt.Println(v3, v3, v1)
    for fi54 := 0; fi54 < 5; fi54++ {
        switch v6 {
        case 9223372036854775807:
            var v55 bool = (127 <= 0)
            _ = v55
        case 2147483647:
            if (false && (v6 < v6)) {
                v56 := GenS6{F0: 2, F1: 1, F2: 127}
                switch ("a" + "") {
                case "a":
                    v1 = (v1 - (float32(-2.25) * v1))
                case "":
                    v1 = ((1e10 + 1e10) / (v1 + (3.14159 - (1e10 - (v1 / (v1 + 1e10))))))
                default:
                    if len(v5) > 0 {
                        v5[0] = 935
                    }
                }
                _ = v56
            }
        }
        switch "" {
        case "a":
            var v57 int8 = int8(v6)
            _ = v57
        case "":
            var v58 int64 = 1
            _ = v58
        }
        v59 := GenS7{F0: 42}
        _ = fi54
        _ = v59
    }
    v3 = 18446744073709551615
    var v60 int32 = 130
    for fi61 := 0; fi61 < 6; fi61++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi62 := 0; fi62 < 7; fi62++ {
            v60 -= v60
            fmt.Println(v3)
            if (406 != 94) {
                if ((!(true || ((true || false) || true))) && (1 > 1)) {
                    var v63 float64 = ((-2.25 + -2.25) - 1e10)
                    if (65535 >= (v3 * v3)) {
                        _ = len(v2)
                        var v64 int8 = -128
                        _ = v64
                    }
                    v1 = 1.5
                    _ = v63
                }
                if (!true) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                var v65 float32 = ((v1 * v1) - (1e10 + float32(-2.25)))
                _ = v65
            }
            _ = fi62
        }
        var v66 int16 = -1
        _ = fi61
        _ = v66
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    switch ("a" + "") {
    case "":
        if true {
            var v67 uint32 = 325
            for fi68 := 0; fi68 < 6; fi68++ {
                var v69 uint16 = 722
                if false {
                    v3 = (v3 * (v3 << ((v3) & 63)))
                    switch "hello" {
                    case "hello":
                        for fi70 := 0; fi70 < 13; fi70++ {
                            var v71 int8 = -9
                            v72 := GenS5{F0: "a"}
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                    }
                } else {
                    var v73 uint32 = (v67 + (v67 * 4294967295))
                    _ = len(v2)
                    for fi74 := 0; fi74 < 17; fi74++ {
                        v73 -= v67
                        _ = fi74
                    }
                    _ = v73
                }
                _ = len(v2)
                _ = fi68
                _ = v69
            }
            if false {
                v60 = v60
                switch "\n\t" {
                case "":
                    if false {
                        v75 := GenS8{F0: 42, F1: v6}
                        switch ("\n\t" + "\n\t") {
                        case "":
                            v3 += v3
                        case "a":
                            var v76 uint32 = v67
                            _ = v76
                        default:
                            var v77 int16 = -405
                            _ = v77
                        }
                        v3 = v3
                        _ = v75
                    } else {
                        delete(v2, (v6 / (((^(v53))) | 1)))
                        v78 := GenS1{F0: -1, F1: 182, F2: 42}
                        for fi79 := 0; fi79 < 11; fi79++ {
                            switch 540 {
                            case -991:
                                if len(v5) > 0 {
                                    v5[0] = v67
                                }
                            }
                            _ = fi79
                        }
                        _ = v78
                    }
                case "a":
                    v67 = 2147483648
                default:
                    if len(v5) > 0 {
                        v5[0] = (v67 / ((uint32(v3)) | 1))
                    }
                }
            } else {
                fmt.Println(v1)
                if (((v3 % ((42) | 1)) >= 18446744073709551615) && true) {
                    v60 += int32(v67)
                    fmt.Println(v1, v67)
                    for fi80 := 0; fi80 < 5; fi80++ {
                        var v81 int16 = 0
                        var v82 int32 = (v60 & 42)
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                } else {
                    v83 := GenS9{F0: v6, F1: v53, F2: 0}
                    if (127 == 127) {
                        fmt.Println(v1, v3)
                        if len(v5) > 0 {
                            v5[0] = v67
                        }
                    } else {
                        delete(v4, -791)
                        var v84 uint16 = 32767
                        _ = v84
                    }
                    _ = v83
                }
                var v85 float32 = 0.1
                _ = v85
            }
            _ = v67
        } else {
            var v86 uint32 = 2147483647
            var v87 float32 = v1
            _ = v86
            _ = v87
        }
    default:
        var v88 int16 = -783
        _ = v88
    }
    switch v6 {
    case 701:
        var v89 uint32 = 2147483648
        _ = v89
    case 32768:
        if (1 >= 1) {
            if false {
                for fi90 := 0; fi90 < 15; fi90++ {
                    v3 = 106
                    if (10 < 40) {
                        switch (("a" + ("" + "a")) + ("\n\t" + "\n\t")) {
                        case "foo bar":
                            var v91 bool = (false || true)
                            _ = v91
                        case "\n\t":
                            if (("a" + "") != "foo bar") {
                                if true {
                                    v3 -= (v3 & (v3 | (v3 << ((180) & 63))))
                                    for fi92 := 0; fi92 < 7; fi92++ {
                                        v93 := GenS10{F0: (v3 * v3)}
                                        _ = fi92
                                        _ = v93
                                    }
                                    fmt.Println(v53, fi90)
                                }
                            }
                        case "a":
                            for fi94 := 0; fi94 < 17; fi94++ {
                                if false {
                                    var v95 bool = true
                                    _ = v95
                                } else {
                                    var v96 uint16 = 693
                                    var v97 int16 = -32768
                                    var v98 uint16 = (v96 & v96)
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                }
                                fmt.Println(fi90)
                                _ = fi94
                            }
                        }
                        if false {
                            var v99 bool = false
                            var v100 uint32 = 410
                            fmt.Println(v53)
                            _ = v100
                            _ = v99
                        }
                        fmt.Println(v6)
                    }
                    _ = fi90
                }
                v60 = (v60 & 42)
                for fi101 := 0; fi101 < 13; fi101++ {
                    var v102 int8 = 42
                    fmt.Println(v60, fi101, fi101)
                    _ = fi101
                    _ = v102
                }
            } else {
                var v103 uint64 = 116
                v5 = append(v5, 187)
                for fi104 := 0; fi104 < 20; fi104++ {
                    switch v53 {
                    case 128:
                        if false {
                            v60 = (v60 * v60)
                            v1 = (v1 - (v1 + (v1 / v1)))
                            if (true && false) {
                                fmt.Println(v1)
                                if (255 <= 178) {
                                    var v105 int32 = v60
                                    _ = v105
                                }
                                if ((((v1 * 1e10) - (((v1 / 0.1) * v1) - v1)) - (-2.25 - float32((float64(v1) * 0.0)))) == ((v1 - 1e10) * 1e10)) {
                                    v106 := GenS11{F0: (v3 & v3), F1: 4294967295, F2: -814}
                                    fmt.Println(v6, v3)
                                    _ = v106
                                }
                            }
                        }
                    case -9223372036854775808:
                        for fi107 := 0; fi107 < 19; fi107++ {
                            switch "a" {
                            case "":
                                fmt.Println(v6)
                            case "hello":
                                v108 := v2[fi104]
                                _ = v108
                            case "foo bar":
                                v53 = fi104
                            default:
                                fmt.Println(fi107)
                            }
                            v109 := GenS12{F0: 1, F1: 255, F2: ("hello" + ("hello" + "foo bar"))}
                            var v110 int32 = 2147483647
                            _ = fi107
                            _ = v109
                            _ = v110
                        }
                    case 32767:
                        var v111 bool = true
                        _ = v111
                    }
                    if (0 == 588) {
                        fmt.Println(v60, v1)
                        v5 = append(v5, 2147483648)
                        v1 -= (v1 - v1)
                    }
                    var v112 float32 = 1.5
                    _ = fi104
                    _ = v112
                }
                _ = v103
            }
            v53 += v53
        }
    case 32767:
        for fi113 := 0; fi113 < 6; fi113++ {
            v6 -= v6
            v1 -= 0.1
            _ = fi113
        }
    }
    v3 += (v3 * v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v60
}
