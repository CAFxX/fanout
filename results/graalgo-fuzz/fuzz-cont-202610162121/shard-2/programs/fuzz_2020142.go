package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int32; F2 int32 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int16; F1 uint32 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 int8 = -128
    if (v1 > (v1 + v1)) {
        if false {
            fmt.Println(v1, v1)
        } else {
            for fi5 := 0; fi5 < 16; fi5++ {
                v6 := GenS1{F0: "foo bar"}
                _ = fi5
                _ = v6
            }
        }
    }
    switch v1 {
    case "hello":
        if false {
            if (2 < 255) {
                if (!false) {
                    if ("hello" == "hello") {
                        switch "\n\t" {
                        case "hello":
                            v7 := GenS1{F0: v1}
                            _ = v7
                        case "":
                            v8 := GenS2{F0: 64}
                            _ = v8
                        case "\n\t":
                            if false {
                                v9 := []int{32768}
                                var v10 int16 = 1000
                                if (1 != 128) {
                                    var v11 uint16 = 0
                                    _ = v11
                                }
                                _ = v10
                                _ = v9
                            }
                        }
                    } else {
                        var v12 uint32 = 175
                        v2[-575] = "hello"
                        var v13 float32 = 0.0
                        _ = v12
                        _ = v13
                    }
                } else {
                    if (false && true) {
                        v1 = ((v1 + "\n\t") + (v1 + ""))
                        var v14 uint8 = 128
                        _ = v14
                    }
                }
            } else {
                for fi15 := 0; fi15 < 15; fi15++ {
                    if true {
                        v16 := v2[86]
                        _ = v16
                    }
                    _ = fi15
                }
                for fi17 := 0; fi17 < 7; fi17++ {
                    switch ("a" + ((v1 + v1) + "\n\t")) {
                    case "":
                        for fi18 := 0; fi18 < 1; fi18++ {
                            var v19 uint32 = 4294967295
                            var v20 int32 = 608
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    case "hello":
                        var v21 string = (v1 + ("" + v1))
                        _ = v21
                    }
                    fmt.Println(v1, fi17, v4)
                    _ = fi17
                }
            }
            if false {
                v22 := []int{}
                switch 32767 {
                case 127:
                    v1 = (("a" + "\n\t") + "\n\t")
                }
                switch 32767 {
                case 32768:
                    v4 += -76
                case 1:
                    v4 = (v4 + (v4 % (((v4 << ((4294967295) & 63))) | 1)))
                case 9223372036854775807:
                    if (42 > 326) {
                        if (float64(0.0) >= (0.0 - 3.14159)) {
                            switch v1 {
                            case "":
                                var v23 uint16 = 2
                                _ = v23
                            case "hello":
                                var v24 uint64 = 0
                                _ = v24
                            default:
                                v1 = "\n\t"
                            }
                        } else {
                            fmt.Println(v4, v1)
                            fmt.Println(v1)
                        }
                        v4 = (v4 & (-(v4)))
                        for fi25 := 0; fi25 < 18; fi25++ {
                            v4 -= v4
                            _ = fi25
                        }
                    }
                default:
                    if (-32768 == 0) {
                        var v26 uint64 = 42
                        _ = v26
                    } else {
                        _ = len(v22)
                        for fi27 := 0; fi27 < 17; fi27++ {
                            var v28 uint64 = 4294967295
                            _ = fi27
                            _ = v28
                        }
                        fmt.Println(v4, v4)
                    }
                }
                _ = v22
            } else {
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 uint = 4294967295
                    if (false && false) {
                        switch fi29 {
                        case 2147483647:
                            var v31 float64 = (1e10 - 3.14159)
                            _ = v31
                        case -1:
                            delete(v3, v1)
                        }
                        v32 := GenS1{F0: ""}
                        v33 := GenS1{F0: v1}
                        _ = v32
                        _ = v33
                    } else {
                        v34 := GenS3{F0: 0, F1: 1}
                        _ = v34
                    }
                    v35 := GenS2{F0: 2}
                    _ = fi29
                    _ = v30
                    _ = v35
                }
                if true {
                    delete(v3, "\n\t")
                } else {
                    var v36 uint64 = 601
                    delete(v2, 313)
                    v36 = 42
                    _ = v36
                }
                if true {
                    var v37 float64 = (((0.1 * (3.14159 + float64(0.0))) - 0.1) - 3.14159)
                    v1 = ("a" + (v1 + v1))
                    switch 4294967295 {
                    case 32768:
                        if (!false) {
                            if (!(!(2147483647 <= 18))) {
                                var v38 uint8 = 88
                                _ = len(v2)
                                _ = v38
                            }
                        }
                    case 200:
                        delete(v3, "")
                    }
                    _ = v37
                }
            }
            v4 += (^(v4))
        }
    default:
        for fi39 := 0; fi39 < 12; fi39++ {
            v40 := GenS1{F0: v1}
            var v41 uint = 42
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    switch (v1 + "a") {
    case "a":
        if false {
            delete(v2, 532)
            v4 = v4
        } else {
            switch -339 {
            case 2147483647:
                if true {
                    switch (v1 + v1) {
                    case "":
                        fmt.Println(v4)
                    case "foo bar":
                        var v42 int64 = -1
                        _ = v42
                    case "hello":
                        v43 := []int{112, 4294967295}
                        _ = v43
                    default:
                        switch 0 {
                        case -574:
                            if (true || true) {
                                switch v1 {
                                case "":
                                    var v44 int16 = 32767
                                    _ = v44
                                default:
                                    fmt.Println(v4)
                                }
                                _ = len(v2)
                            }
                        }
                    }
                    switch 585 {
                    case -237:
                        v45 := GenS3{F0: 32767, F1: 1}
                        _ = v45
                    default:
                        v46 := []int{458, int(v4)}
                        _ = v46
                    }
                    v1 = (("foo bar" + "hello") + v1)
                }
            }
            switch 4294967295 {
            case 1:
                v47 := GenS4{F0: 0}
                _ = v47
            case -62:
                switch -318 {
                case 4294967295:
                    v4 += (v4 | (v4 << ((2) & 63)))
                case 314:
                    v48 := GenS2{F0: 42}
                    _ = v48
                default:
                    if true {
                        var v49 uint32 = 42
                        for fi50 := 0; fi50 < 6; fi50++ {
                            if false {
                                var v51 bool = false
                                if false {
                                    if true {
                                        var v52 uint8 = 127
                                        _ = v52
                                    }
                                    if true {
                                        v53 := GenS1{F0: ""}
                                        var v54 bool = v51
                                        _ = v53
                                        _ = v54
                                    } else {
                                        var v55 uint64 = 892
                                        v56 := GenS1{F0: (v1 + ((v1 + v1) + ((v1 + ("\n\t" + (v1 + ((v1 + v1) + v1)))) + v1)))}
                                        _ = v55
                                        _ = v56
                                    }
                                } else {
                                    var v57 int = (fi50 / ((fi50) | 1))
                                    var v58 uint8 = 127
                                    _ = v57
                                    _ = v58
                                }
                                switch (fi50 ^ fi50) {
                                case -825:
                                    for fi59 := 0; fi59 < 4; fi59++ {
                                        fmt.Println(v1, v51, v1)
                                        v1 = "foo bar"
                                        fmt.Println(fi50, v51, v4)
                                        _ = fi59
                                    }
                                case 9223372036854775807:
                                    if v51 {
                                        v60 := GenS5{F0: 1000, F1: -1, F2: 1}
                                        var v61 string = "a"
                                        for fi62 := 0; fi62 < 1; fi62++ {
                                            v63 := GenS6{F0: 42, F1: 0}
                                            v64 := GenS2{F0: 18446744073709551615}
                                            _ = fi62
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = v60
                                        _ = v61
                                    } else {
                                        fmt.Println(v4, v1, fi50)
                                        var v65 uint64 = 18446744073709551615
                                        _ = v65
                                    }
                                default:
                                    if (125 > 255) {
                                        v66 := v3[v1]
                                        v66 += 2147483647
                                        var v67 uint8 = 2
                                        _ = v66
                                        _ = v67
                                    } else {
                                        for fi68 := 0; fi68 < 20; fi68++ {
                                            var v69 uint64 = 4294967295
                                            if (1 > 518) {
                                                delete(v3, v1)
                                                var v70 uint64 = 0
                                                _ = v70
                                            } else {
                                                v71 := GenS6{F0: v49, F1: 200}
                                                var v72 bool = ((2 < 65535) && v51)
                                                var v73 float64 = 0.0
                                                v51 = ((fi50 >> ((18446744073709551615) & 63)) < (fi50 ^ fi50))
                                                var v74 bool = (4294967295 > v49)
                                                v75 := GenS7{F0: 2}
                                                switch v1 {
                                                case "":
                                                    v76 := []int{4294967295, -716, 0}
                                                    if (v72 && (true && true)) {
                                                        fmt.Println(v51, v4)
                                                    } else {
                                                        if v72 {
                                                            v77 := GenS7{F0: 4294967295}
                                                            v78 := GenS3{F0: 1000, F1: -1}
                                                            fmt.Println(v49)
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                    }
                                                    if (!(((v49 % ((v49) | 1)) <= (v49 << ((943) & 63))) || (!false))) {
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v79 int32 = -195
                                                            _ = v79
                                                        case "foo bar":
                                                            v49 = 701
                                                        case "\n\t":
                                                            var v80 uint32 = 2147483647
                                                            var v81 uint = 413
                                                            _ = v80
                                                            _ = v81
                                                        default:
                                                            var v82 int16 = 32767
                                                            _ = v82
                                                        }
                                                    } else {
                                                        _ = len(v76)
                                                        v72 = (v69 == v69)
                                                    }
                                                    var v83 int32 = -1
                                                    _ = v76
                                                    _ = v83
                                                case "hello":
                                                    fmt.Println(v51, v73, v74)
                                                }
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                                _ = v74
                                                _ = v75
                                            }
                                            _ = fi68
                                            _ = v69
                                        }
                                    }
                                }
                                _ = v51
                            } else {
                                switch v1 {
                                case "":
                                    switch v1 {
                                    case "foo bar":
                                        v2[(fi50 & 128)] = (v1 + "hello")
                                    }
                                case "foo bar":
                                    for fi84 := 0; fi84 < 4; fi84++ {
                                        fmt.Println(v49, v4)
                                        fmt.Println(v4, v49, fi50)
                                        v85 := []int{-714, fi84, (fi50 * int(v49))}
                                        _ = fi84
                                        _ = v85
                                    }
                                }
                                v49 = v49
                                fmt.Println(v1)
                            }
                            v1 = "a"
                            _ = fi50
                        }
                        for fi86 := 0; fi86 < 10; fi86++ {
                            v87 := GenS7{F0: 42}
                            if true {
                                v4 = (v4 % ((v4) | 1))
                                switch (v1 + v1) {
                                case "a":
                                    if true {
                                        fmt.Println(v1, v4, v49)
                                    } else {
                                        _ = len(v3)
                                        v88 := GenS5{F0: -32768, F1: -38, F2: 920}
                                        if (1.5 > 0.0) {
                                            var v89 int64 = 1
                                            v90 := v2[(fi86 << ((438) & 63))]
                                            _ = v89
                                            _ = v90
                                        }
                                        _ = v88
                                    }
                                default:
                                    v49 = v49
                                }
                            }
                            _ = fi86
                            _ = v87
                        }
                        _ = v49
                    }
                }
            case 0:
                if true {
                    for fi91 := 0; fi91 < 14; fi91++ {
                        v92 := []int{}
                        v1 = ("foo bar" + v1)
                        v1 = "\n\t"
                        _ = fi91
                        _ = v92
                    }
                } else {
                    fmt.Println(v4, v4)
                    var v93 int16 = 826
                    v1 = (v1 + ("" + "hello"))
                    _ = v93
                }
            }
        }
    case "hello":
        v4 += (v4 << ((2) & 63))
    default:
        switch "\n\t" {
        case "hello":
            var v94 uint8 = 166
            _ = v94
        case "foo bar":
            var v95 uint = 662
            _ = v95
        case "\n\t":
            if true {
                var v96 uint32 = 332
                var v97 int32 = 1
                _ = v96
                _ = v97
            }
        }
    }
    v98 := []int{}
    var v99 bool = (float32((((-2.25 * 1.5) - 1.5) * float64((1.5 * 3.14159)))) <= -2.25)
    fmt.Println(v99, v1, v99)
    v3[v1] = 9223372036854775807
    var v100 int8 = v4
    switch int(v4) {
    case -2147483648:
        var v101 int8 = v100
        _ = v101
    default:
        v102 := v3[v1]
        _ = v102
    }
    _ = len(v98)
    v1 = v1
    if (-1 < 0) {
        v3["hello"] = -375
    } else {
        v99 = v99
        v99 = (true || v99)
        var v103 uint64 = 599
        _ = v103
    }
    v104 := GenS5{F0: 32767, F1: 42, F2: 1}
    v99 = false
    v4 += (v100 % ((127) | 1))
    if v99 {
        if (0 == -9223372036854775808) {
            if (!(2147483647 > -2147483648)) {
                var v105 uint64 = 18446744073709551615
                v106 := GenS6{F0: 538, F1: 158}
                _ = v105
                _ = v106
            } else {
                if true {
                    if false {
                        var v107 int32 = 1
                        switch 128 {
                        case 0:
                            v108 := GenS8{F0: 0}
                            _ = v108
                        case -1:
                            for fi109 := 0; fi109 < 2; fi109++ {
                                if v99 {
                                    v110 := GenS9{F0: 0}
                                    if (v1 > v1) {
                                        var v111 bool = (!(1.5 != -2.25))
                                        _ = v111
                                    }
                                    _ = v110
                                }
                                _ = fi109
                            }
                        default:
                            v98 = append(v98, -2147483648)
                        }
                        v107 = -2147483648
                        _ = v107
                    }
                } else {
                    fmt.Println(v100, v99, v1)
                }
                v112 := GenS10{F0: 1}
                delete(v2, 32767)
                _ = v112
            }
            for fi113 := 0; fi113 < 2; fi113++ {
                v99 = v99
                switch ("" + v1) {
                case "a":
                    if len(v98) >= 2 {
                        v98 = v98[1:len(v98)]
                    }
                case "foo bar":
                    var v114 uint32 = 42
                    _ = v114
                case "":
                    var v115 int16 = -1
                    _ = v115
                }
                var v116 bool = v99
                _ = fi113
                _ = v116
            }
            v4 = (v100 & v100)
        }
        v98 = append(v98, -610)
    }
    if true {
        v99 = v99
        if (9223372036854775807 == -9223372036854775808) {
            v117 := GenS11{F0: -32768, F1: 2}
            _ = v117
        }
        switch 32768 {
        case 42:
            if false {
                var v118 int32 = -2147483648
                _ = len(v3)
                _ = v118
            } else {
                if v99 {
                    v1 = ((v1 + (v1 + "a")) + "")
                } else {
                    if (!true) {
                        switch 32768 {
                        case 0:
                            v100 += v100
                        case -588:
                            v1 = v1
                        }
                    }
                    v100 += (v4 | v100)
                    for fi119 := 0; fi119 < 12; fi119++ {
                        v3["\n\t"] = (fi119 - fi119)
                        v120 := v2[fi119]
                        var v121 uint = 2
                        _ = fi119
                        _ = v120
                        _ = v121
                    }
                }
                var v122 uint = 45
                fmt.Println(v4, v1)
                _ = v122
            }
        }
    } else {
        v123 := GenS10{F0: 1}
        _ = v123
    }
    var v124 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v124:", v124)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v98:", len(v98), v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v104
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v98
    _ = v99
}
