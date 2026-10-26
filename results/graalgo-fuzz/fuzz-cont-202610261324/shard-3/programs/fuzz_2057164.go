package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 bool; F2 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 string }

func main() {
    v1 := []int{2147483647, 0}
    var v2 bool = true
    var v3 uint8 = 1
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    switch 2147483647 {
    case 127:
        v2 = v2
    case 1:
        v6 := GenS1{F0: true, F1: 0}
        _ = v6
    case -391:
        v2 = (!v2)
    default:
        switch ("a" + "a") {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            if (((!(-1 != 42)) && true) || (!false)) {
                fmt.Println(v3, v3)
                v2 = v2
                switch "" {
                case "\n\t":
                    if true {
                        if (float32(1e10) > -2.25) {
                            if false {
                                fmt.Println(v2)
                            } else {
                                var v7 int32 = -614
                                var v8 int8 = 90
                                var v9 uint = 206
                                _ = v7
                                _ = v8
                                _ = v9
                            }
                            for fi10 := 0; fi10 < 16; fi10++ {
                                var v11 int64 = 0
                                var v12 uint16 = 2
                                v13 := GenS2{F0: 74, F1: (v12 - (v12 & 32767))}
                                _ = fi10
                                _ = v11
                                _ = v12
                                _ = v13
                            }
                        }
                    } else {
                        fmt.Println(v3, v3)
                        fmt.Println(v3)
                    }
                case "a":
                    var v14 int16 = 1
                    _ = v14
                case "foo bar":
                    v1 = append(v1, 553)
                }
            } else {
                switch 32767 {
                case -484:
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v16 := GenS2{F0: -1, F1: 42}
                        v17 := GenS3{F0: 32767}
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case 2147483647:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 32768:
                    if len(v1) > 0 {
                        v1[0] = 641
                    }
                default:
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 int64 = 867
                        _ = fi18
                        _ = v19
                    }
                }
                v20 := GenS2{F0: 0, F1: uint16(v3)}
                v2 = v2
                _ = v20
            }
        default:
            if (true && false) {
                var v21 uint8 = v3
                v22 := GenS4{F0: 127, F1: v2, F2: 813}
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 float32 = 1e10
                    switch 66 {
                    case 42:
                        if false {
                            for fi25 := 0; fi25 < 4; fi25++ {
                                v3 = v3
                                v26 := v5[((("" + "foo bar") + "foo bar") + "")]
                                _ = fi25
                                _ = v26
                            }
                            var v27 bool = false
                            _ = v27
                        } else {
                            delete(v4, "hello")
                        }
                    case -9223372036854775808:
                        v3 = 128
                    case 2147483647:
                        v28 := GenS4{F0: -15, F1: (((904 == fi23) || v2) || (((v3 * (v21 / ((0) | 1))) == 255) && false)), F2: -1}
                        _ = v28
                    default:
                        var v29 uint = 431
                        _ = v29
                    }
                    _ = fi23
                    _ = v24
                }
                _ = v21
                _ = v22
            }
        }
    }
    switch "\n\t" {
    case "foo bar":
        v30 := GenS5{F0: 42}
        _ = v30
    case "a":
        switch ("foo bar" + ("" + "foo bar")) {
        case "\n\t":
            fmt.Println(v2, v3)
        case "hello":
            v3 = (v3 / ((v3) | 1))
        default:
            if v2 {
                v4["a"] = 42
                for fi31 := 0; fi31 < 8; fi31++ {
                    v3 += (v3 + v3)
                    v32 := GenS5{F0: 32767}
                    _ = fi31
                    _ = v32
                }
                v2 = false
            }
        }
    case "":
        var v33 float32 = 1.5
        _ = v33
    default:
        if v2 {
            delete(v5, "\n\t")
            var v34 uint32 = 0
            _ = v34
        } else {
            v2 = v2
            for fi35 := 0; fi35 < 8; fi35++ {
                v2 = (9223372036854775807 >= 1)
                if (!(fi35 >= fi35)) {
                    var v36 int32 = -2147483648
                    if v2 {
                        for fi37 := 0; fi37 < 20; fi37++ {
                            var v38 int = fi35
                            v39 := GenS3{F0: 255}
                            if len(v1) > 0 {
                                v1[0] = (fi35 << ((0) & 63))
                            }
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 6; fi40++ {
                            var v41 float32 = (((1e10 + (-2.25 / (1e10 * float32(((0.0 - ((1e10 + 1e10) * 1e10)) / float64(1e10)))))) * 0.0) - ((1.5 * 3.14159) - 1e10))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v42 uint64 = 421
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        var v43 float32 = 0.1
                        _ = v43
                    } else {
                        switch (("" + "hello") + (((("hello" + "hello") + "\n\t") + "hello") + "\n\t")) {
                        case "hello":
                            for fi44 := 0; fi44 < 17; fi44++ {
                                v2 = false
                                if v2 {
                                    if (!(!(v2 || false))) {
                                        v45 := GenS2{F0: (v36 << ((2) & 63)), F1: 65535}
                                        switch (-(fi44)) {
                                        case -207:
                                            if v2 {
                                                v36 = v36
                                            } else {
                                                var v46 uint = 1
                                                var v47 float32 = (0.1 - 1.5)
                                                var v48 int16 = 32767
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                            }
                                        case 42:
                                            v36 = v36
                                        case -796:
                                            v2 = (42 != 250)
                                        default:
                                            var v49 uint8 = v3
                                            _ = v49
                                        }
                                        v3 -= v3
                                        _ = v45
                                    }
                                    switch 0 {
                                    case -982:
                                        v2 = (127 != 100)
                                    case 820:
                                        v3 = v3
                                    case 32768:
                                        var v50 int64 = 42
                                        _ = v50
                                    default:
                                        var v51 int32 = -1
                                        _ = v51
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (fi44 | fi44)
                                    }
                                    var v52 int32 = (v36 + (-(v36)))
                                    _ = v52
                                } else {
                                    switch ((("" + "") + ((("foo bar" + "\n\t") + "") + ("\n\t" + ("" + "\n\t")))) + "a") {
                                    case "a":
                                        v36 += v36
                                    default:
                                        v3 = v3
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        v2 = ((v36 > v36) || (!(!true)))
                                    case "":
                                        v5["hello"] = ("a" + "")
                                    default:
                                        var v53 int32 = v36
                                        _ = v53
                                    }
                                }
                                var v54 uint32 = 0
                                _ = fi44
                                _ = v54
                            }
                        default:
                            if false {
                                var v55 int64 = 9223372036854775807
                                fmt.Println(v36, v3)
                                _ = v55
                            }
                        }
                    }
                    _ = v36
                } else {
                    if (271 < 65535) {
                        fmt.Println(v2, v2)
                    } else {
                        if ((fi35 + fi35) != 1) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    v2 = v2
                    var v56 bool = v2
                    _ = v56
                }
                switch (^(fi35)) {
                case 32767:
                    fmt.Println(fi35, v2, fi35)
                default:
                    var v57 float32 = (1e10 / 3.14159)
                    _ = v57
                }
                _ = fi35
            }
            v58 := GenS4{F0: 0, F1: false, F2: -1}
            _ = v58
        }
    }
    v2 = v2
    if (!true) {
        if (!(!(v2 || (255 < 255)))) {
            v2 = (42 < 1)
            v3 += (v3 ^ v3)
            v2 = (((-1 > 42) && false) && ((-2147483648 == 1) && false))
        }
        if true {
            if ((0.0 - (-2.25 * 3.14159)) < 1.5) {
                v1 = append(v1, 127)
                var v59 float32 = -2.25
                for fi60 := 0; fi60 < 1; fi60++ {
                    fmt.Println(v59, v59, v59)
                    for fi61 := 0; fi61 < 15; fi61++ {
                        var v62 bool = v2
                        if (v62 && false) {
                            if (false || (757 == 255)) {
                                var v63 uint = 2
                                v63 += v63
                                v64 := GenS1{F0: true, F1: 314}
                                _ = v63
                                _ = v64
                            } else {
                                var v65 int16 = -1
                                _ = v65
                            }
                        } else {
                            switch fi60 {
                            case 32768:
                                _ = len(v4)
                            case 1:
                                v59 += v59
                            }
                            switch (fi60 % ((4294967295) | 1)) {
                            case 32767:
                                v3 = v3
                            case 42:
                                if v62 {
                                    var v66 float64 = (1e10 / 3.14159)
                                    v62 = v62
                                    fmt.Println(v66)
                                    _ = v66
                                } else {
                                    if (2147483647 >= 1) {
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v67 int16 = 1000
                                            _ = v67
                                        case "foo bar":
                                            if (true || v2) {
                                                var v68 int16 = 31
                                                _ = v68
                                            }
                                        default:
                                            fmt.Println(fi61, v2)
                                        }
                                        if v62 {
                                            v4[("" + "hello")] = 655
                                            var v69 bool = v2
                                            _ = v69
                                        } else {
                                            switch (fi61 + 127) {
                                            case 128:
                                                v3 = 128
                                            default:
                                                switch (fi60 >> ((765) & 63)) {
                                                case -1:
                                                    v70 := GenS6{F0: (fi61 >> ((0) & 63))}
                                                    _ = v70
                                                default:
                                                    var v71 uint16 = 255
                                                    _ = v71
                                                }
                                            }
                                        }
                                        v72 := GenS1{F0: (!(!true)), F1: 42}
                                        fmt.Println(v62, fi60)
                                        var v73 bool = (!true)
                                        var v74 int64 = -395
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                }
                            default:
                                switch fi60 {
                                case -48:
                                    fmt.Println(v2)
                                case -988:
                                    v75 := v5[("\n\t" + (("\n\t" + "\n\t") + "foo bar"))]
                                    _ = v75
                                case -2147483648:
                                    var v76 float64 = 0.0
                                    _ = v76
                                }
                            }
                            v59 -= ((v59 + v59) * 3.14159)
                            fmt.Println(v59, v2, v62)
                            var v77 bool = true
                            var v78 uint16 = 155
                            var v79 uint16 = (^(v78))
                            switch fi60 {
                            case 531:
                                switch fi60 {
                                case 1000:
                                    var v80 int16 = 1
                                    _ = v80
                                case -136:
                                    var v81 int8 = -128
                                    _ = v81
                                case 756:
                                    _ = len(v1)
                                }
                            default:
                                var v82 int8 = 115
                                _ = v82
                            }
                            v77 = v62
                            _ = v77
                            _ = v78
                            _ = v79
                        }
                        _ = fi61
                        _ = v62
                    }
                    v83 := GenS5{F0: 1}
                    _ = fi60
                    _ = v83
                }
                _ = v59
            }
        } else {
            var v84 uint16 = 904
            _ = v84
        }
        switch 42 {
        case 42:
            _ = len(v5)
        }
    }
    v3 += 128
    v85 := GenS4{F0: 7, F1: (v2 && v2), F2: 1000}
    switch 127 {
    case 128:
        for fi86 := 0; fi86 < 2; fi86++ {
            fmt.Println(v3, v3, v2)
            var v87 int16 = -32768
            _ = fi86
            _ = v87
        }
    default:
        var v88 uint16 = 32767
        _ = v88
    }
    switch "hello" {
    case "":
        for fi89 := 0; fi89 < 6; fi89++ {
            switch fi89 {
            case 32767:
                v90 := GenS7{F0: 272, F1: 2147483648}
                _ = v90
            case -666:
                if v2 {
                    for fi91 := 0; fi91 < 12; fi91++ {
                        v3 += 232
                        v2 = v2
                        if (164 != 9223372036854775807) {
                            var v92 uint32 = 0
                            v93 := GenS8{F0: -9223372036854775808}
                            var v94 int16 = 1
                            _ = v92
                            _ = v93
                            _ = v94
                        } else {
                            v95 := GenS6{F0: fi89}
                            if v2 {
                                var v96 int = (-(fi89))
                                _ = v96
                            } else {
                                delete(v5, "")
                            }
                            _ = v95
                        }
                        _ = fi91
                    }
                    v97 := GenS9{F0: 218}
                    for fi98 := 0; fi98 < 20; fi98++ {
                        if false {
                            if (42 > -1) {
                                var v99 float64 = (3.14159 + 0.1)
                                var v100 uint16 = 265
                                _ = v100
                                _ = v99
                            }
                            if (((v3 * v3) != 100) || (!(315 != 42))) {
                                if v2 {
                                    switch (-(fi98)) {
                                    case 2147483647:
                                        var v101 int16 = 1
                                        _ = v101
                                    default:
                                        v4["\n\t"] = 1
                                    }
                                    _ = len(v4)
                                    v3 += (v3 % ((v3) | 1))
                                } else {
                                    v102 := GenS9{F0: -1}
                                    _ = v102
                                }
                                var v103 float64 = 0.0
                                _ = v103
                            }
                        } else {
                            if v2 {
                                v2 = (65535 >= 0)
                            }
                            v3 -= v3
                        }
                        switch ("a" + "\n\t") {
                        case "\n\t":
                            v2 = true
                        case "foo bar":
                            fmt.Println(v3, fi89, fi98)
                        }
                        _ = fi98
                    }
                    _ = v97
                } else {
                    v3 += (^(v3))
                    var v104 uint8 = (v3 * v3)
                    _ = v104
                }
            default:
                v2 = v2
            }
            _ = fi89
        }
    case "hello":
        v3 = (v3 & (v3 + v3))
    case "\n\t":
        if v2 {
            fmt.Println(v2, v3)
        }
    }
    var v105 uint32 = 2
    v106 := GenS1{F0: (!v2), F1: -171}
    v2 = v2
    switch "a" {
    case "a":
        switch -2147483648 {
        case 32767:
            for fi107 := 0; fi107 < 18; fi107++ {
                v108 := v5["hello"]
                v109 := GenS10{F0: 461, F1: (v108 + v108)}
                var v110 float64 = (float64(3.14159) * 0.0)
                _ = fi107
                _ = v108
                _ = v109
                _ = v110
            }
        }
    default:
        switch "hello" {
        case "a":
            v105 -= (v105 % ((v105) | 1))
        case "\n\t":
            if true {
                v111 := GenS6{F0: 2147483647}
                switch 128 {
                case 9223372036854775807:
                    v2 = (0 >= 65535)
                default:
                    v112 := GenS1{F0: v2, F1: 551}
                    _ = v112
                }
                _ = v111
            }
        case "":
            v1 = append(v1, 42)
        }
    }
    v113 := v5["hello"]
    v113 = v113
    fmt.Println(v105, v105, v2)
    v114 := v4[v113]
    for fi115 := 0; fi115 < 8; fi115++ {
        var v116 string = "\n\t"
        v2 = false
        v113 = ((v116 + (v113 + ((v113 + (v113 + "hello")) + (v116 + (v113 + ((v116 + "") + "a")))))) + (v113 + (v116 + "foo bar")))
        _ = fi115
        _ = v116
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v105:", v105)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v105
    _ = v106
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
}
