package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint16; F1 int8; F2 int16 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 int8; F2 int8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint16; F1 int32; F2 int16 }
type GenS14 struct { F0 uint32; F1 int; F2 uint8 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int64; F1 int64 }
type GenS19 struct { F0 uint32; F1 int; F2 bool }
type GenS20 struct { F0 uint32; F1 bool }
type GenS21 struct { F0 uint; F1 int8; F2 uint64 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint8 }
type GenS24 struct { F0 uint8; F1 uint8; F2 string }
type GenS25 struct { F0 uint8 }

func main() {
    var v1 int = 128
    var v2 string = "a"
    var v3 uint8 = 112
    v4 := []int{-9223372036854775808}
    var v5 uint = 18446744073709551615
    fmt.Println(v2, v1, v1)
    v5 += v5
    if false {
        var v6 uint = 667
        _ = v6
    }
    if (-2147483648 > 424) {
        var v7 int16 = 32767
        _ = v7
    } else {
        var v8 bool = (v1 != (-(v1)))
        _ = v8
    }
    var v9 int64 = -495
    if ((-721 <= 2147483647) && true) {
        var v10 bool = false
        switch v2 {
        case "":
            for fi11 := 0; fi11 < 2; fi11++ {
                for fi12 := 0; fi12 < 14; fi12++ {
                    if (!false) {
                        v1 = -317
                        v1 = (-(v1))
                        if v10 {
                            v13 := GenS1{F0: 1}
                            fmt.Println(fi11, fi11, v2)
                            _ = v13
                        }
                    } else {
                        var v14 uint32 = 1
                        switch 42 {
                        case -1:
                            if len(v4) > 0 {
                                v4[0] = (v1 << (((v5 - v5)) & 63))
                            }
                        }
                        _ = v14
                    }
                    v2 = v2
                    _ = fi12
                }
                v1 -= -1
                if (v10 && (v9 <= v9)) {
                    if v10 {
                        v15 := GenS1{F0: -1}
                        var v16 int16 = 0
                        _ = v15
                        _ = v16
                    }
                } else {
                    if false {
                        v3 = 128
                        v17 := GenS1{F0: 2147483647}
                        _ = v17
                    }
                    v1 = int(v5)
                }
                _ = fi11
            }
        }
        v1 += (-(v1))
        _ = v10
    }
    v18 := GenS1{F0: -1}
    var v19 float32 = -2.25
    if false {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint32 = uint32(v3)
            _ = fi20
            _ = v21
        }
        if (!(v5 >= 907)) {
            var v22 float64 = (3.14159 / -2.25)
            _ = v22
        }
    }
    if false {
        v9 = (v9 ^ v9)
        v19 += (v19 / v19)
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            switch (("hello" + v2) + v2) {
            case "hello":
                if (true || (((v3 * 0) <= v3) && false)) {
                    v5 = (v5 + v5)
                    if false {
                        var v25 float64 = 1.5
                        if (2 <= 512) {
                            v26 := GenS1{F0: -2147483648}
                            fmt.Println(v25, v1, fi23)
                            _ = v26
                        }
                        v19 = 0.1
                        fmt.Println(fi24, fi24)
                        _ = v25
                    } else {
                        if true {
                            switch fi24 {
                            case -564:
                                if len(v4) > 0 {
                                    v4[0] = (fi24 - fi24)
                                }
                            case 413:
                                switch (v2 + ("" + "\n\t")) {
                                case "\n\t":
                                    v4 = append(v4, (v1 >> (((^(v5))) & 63)))
                                case "foo bar":
                                    if true {
                                        fmt.Println(fi24)
                                        var v27 uint8 = 26
                                        switch (((v2 + (v2 + "foo bar")) + v2) + v2) {
                                        case "hello":
                                            v28 := GenS1{F0: -1}
                                            _ = v28
                                        case "foo bar":
                                            if false {
                                                if len(v4) > 0 {
                                                    v4[0] = (-(v1))
                                                }
                                                if (-616 <= -1) {
                                                    v29 := GenS2{F0: 42, F1: 4294967295, F2: 794}
                                                    v5 -= v5
                                                    _ = v29
                                                } else {
                                                    var v30 uint64 = 4294967295
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v30
                                                }
                                                var v31 uint32 = 42
                                                var v32 string = (v2 + ("" + (("" + "a") + "a")))
                                                v31 -= (^(v31))
                                                _ = v31
                                                _ = v32
                                            } else {
                                                if true {
                                                    if true {
                                                        var v33 float32 = v19
                                                        if true {
                                                            v1 += v1
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            v1 = 1
                                                        } else {
                                                            var v34 uint = (v5 % ((65535) | 1))
                                                            _ = v34
                                                        }
                                                        var v35 string = v2
                                                        _ = v33
                                                        _ = v35
                                                    } else {
                                                        v19 += 0.1
                                                        var v36 uint16 = 0
                                                        if len(v4) > 0 {
                                                            v4[0] = int(v27)
                                                        }
                                                        var v37 float64 = 1.5
                                                        if false {
                                                            v38 := GenS2{F0: 42, F1: 587, F2: (v36 % (((v36 + (^(v36)))) | 1))}
                                                            _ = v38
                                                        } else {
                                                            if true {
                                                                if (1 >= v9) {
                                                                    var v39 int32 = 42
                                                                    _ = v39
                                                                } else {
                                                                    if (!true) {
                                                                        var v40 int64 = (-(v9))
                                                                        v19 = (v19 + float32(((v37 * v37) / 1.5)))
                                                                        _ = v40
                                                                    } else {
                                                                        var v41 bool = (false && false)
                                                                        _ = v41
                                                                    }
                                                                    if ((v9 * v9) <= (v9 ^ (v9 << (((v5 - uint(v36))) & 63)))) {
                                                                        switch (fi23 * (v1 ^ (v1 - (-(fi23))))) {
                                                                        case -314:
                                                                            _ = len(v4)
                                                                            v1 -= (-(fi23))
                                                                            var v42 uint8 = v27
                                                                            _ = v42
                                                                        case 2147483647:
                                                                            v19 += v19
                                                                        case 128:
                                                                            fmt.Println(fi24)
                                                                        default:
                                                                            if true {
                                                                                v4 = append(v4, 173)
                                                                                _ = len(v4)
                                                                                var v43 int = 779
                                                                                _ = v43
                                                                            }
                                                                        }
                                                                        if (((-1 != -1) || (true || false)) && true) {
                                                                            if false {
                                                                                if (42 != 0) {
                                                                                    if (!(((v19 - 0.0) < float32(v37)) || false)) {
                                                                                        var v44 string = "\n\t"
                                                                                        _ = v44
                                                                                    }
                                                                                    if false {
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            v45 := GenS3{F0: v27, F1: -1}
                                                                                            _ = v45
                                                                                        }
                                                                                        var v46 uint = 1
                                                                                        if false {
                                                                                            var v47 uint = 65535
                                                                                            switch fi24 {
                                                                                            case 128:
                                                                                                v2 = "\n\t"
                                                                                            case -2147483648:
                                                                                                if false {
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                    var v48 uint8 = 2
                                                                                                    _ = v48
                                                                                                } else {
                                                                                                    v49 := GenS1{F0: -381}
                                                                                                    var v50 bool = false
                                                                                                    _ = v49
                                                                                                    _ = v50
                                                                                                }
                                                                                            default:
                                                                                                v3 = v3
                                                                                            }
                                                                                            _ = v47
                                                                                        }
                                                                                        _ = v46
                                                                                    }
                                                                                    v51 := GenS4{F0: v36}
                                                                                    var v52 int = 582
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                } else {
                                                                                    v9 -= (v9 + int64(v5))
                                                                                    if ((v5 ^ 255) >= (v5 - (^(v5)))) {
                                                                                        var v53 uint32 = 219
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            var v54 int32 = 1
                                                                                            _ = v54
                                                                                        case "\n\t":
                                                                                            fmt.Println(v19, v2)
                                                                                        case "":
                                                                                            v55 := GenS4{F0: v36}
                                                                                            _ = v55
                                                                                        default:
                                                                                            v19 += float32(v37)
                                                                                        }
                                                                                        v3 = (v3 % (((v27 - v27)) | 1))
                                                                                        _ = v53
                                                                                    } else {
                                                                                        var v56 uint32 = 297
                                                                                        _ = len(v4)
                                                                                        _ = v56
                                                                                    }
                                                                                    v2 = "foo bar"
                                                                                    switch (-(v1)) {
                                                                                    case 1:
                                                                                        _ = len(v4)
                                                                                    default:
                                                                                        var v57 int64 = 268
                                                                                        _ = v57
                                                                                    }
                                                                                    v1 = fi23
                                                                                    v5 -= (v5 + v5)
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if (v2 >= (v2 + (((v2 + "") + ("hello" + (v2 + v2))) + v2))) {
                                                                            var v58 bool = true
                                                                            v4 = append(v4, -226)
                                                                            v5 -= v5
                                                                            _ = v58
                                                                        } else {
                                                                            switch ((("a" + (((("a" + "foo bar") + "\n\t") + v2) + v2)) + v2) + v2) {
                                                                            case "hello":
                                                                                v59 := GenS5{F0: 953, F1: (v27 + v3), F2: 485}
                                                                                _ = v59
                                                                            case "foo bar":
                                                                                var v60 float64 = float64(v19)
                                                                                _ = v60
                                                                            case "a":
                                                                                switch ("\n\t" + ("\n\t" + v2)) {
                                                                                case "hello":
                                                                                    if true {
                                                                                        var v61 bool = (v19 > ((v19 - v19) - (-2.25 + v19)))
                                                                                        if (1e10 != 0.0) {
                                                                                            var v62 uint32 = 1
                                                                                            v3 = v27
                                                                                            if v61 {
                                                                                                var v63 string = (v2 + (v2 + v2))
                                                                                                var v64 uint32 = v62
                                                                                                switch ((v63 + (v2 + v63)) + ((v2 + v63) + v63)) {
                                                                                                case "":
                                                                                                    v4 = append(v4, v1)
                                                                                                }
                                                                                                _ = v63
                                                                                                _ = v64
                                                                                            } else {
                                                                                                if v61 {
                                                                                                    if ((1000 >= -73) && v61) {
                                                                                                        v65 := GenS6{F0: (^(v36)), F1: 1, F2: 497}
                                                                                                        v2 = ("hello" + v2)
                                                                                                        _ = v65
                                                                                                    } else {
                                                                                                        if v61 {
                                                                                                            v66 := GenS7{F0: v36, F1: v2}
                                                                                                            switch 32768 {
                                                                                                            case 9223372036854775807:
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = (^(fi24))
                                                                                                                }
                                                                                                            default:
                                                                                                                var v67 bool = (42 != 76)
                                                                                                                _ = v67
                                                                                                            }
                                                                                                            _ = v66
                                                                                                        }
                                                                                                        v68 := GenS3{F0: (v27 ^ v3), F1: 1}
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    v61 = (v61 && v61)
                                                                                                    var v69 uint64 = 18446744073709551615
                                                                                                    _ = v69
                                                                                                }
                                                                                                v70 := GenS8{F0: v1}
                                                                                                var v71 uint32 = (v62 | v62)
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                            }
                                                                                            _ = v62
                                                                                        }
                                                                                        v19 = v19
                                                                                        var v72 uint64 = 42
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            v1 -= (fi24 | 1)
                                                                                        }
                                                                                        var v73 int8 = -128
                                                                                        _ = v61
                                                                                        _ = v72
                                                                                        _ = v73
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v74 uint64 = 868
                                                                            var v75 uint32 = 2
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                } else {
                                                    v76 := GenS9{F0: 441}
                                                    _ = len(v4)
                                                    _ = v76
                                                }
                                                fmt.Println(v27)
                                            }
                                        default:
                                            switch (fi23 ^ (-(fi24))) {
                                            case 638:
                                                var v77 float64 = 0.1
                                                _ = v77
                                            case 4294967295:
                                                v3 += v27
                                            }
                                        }
                                        _ = v27
                                    }
                                default:
                                    switch v2 {
                                    case "foo bar":
                                        if (((v19 / v19) - v19) == v19) {
                                            v78 := GenS10{F0: 100}
                                            v5 = v5
                                            _ = v78
                                        }
                                    case "\n\t":
                                        var v79 int32 = 1
                                        _ = v79
                                    }
                                }
                            }
                        } else {
                            var v80 int16 = -830
                            var v81 float32 = v19
                            _ = v80
                            _ = v81
                        }
                        v3 += 42
                    }
                } else {
                    v5 = v5
                }
            case "foo bar":
                v1 = (fi23 + (fi23 * (fi23 >> ((42) & 63))))
            case "":
                v9 -= 0
            }
            v82 := GenS5{F0: 0, F1: 108, F2: 2147483648}
            _ = fi24
            _ = v82
        }
        v83 := GenS9{F0: -1}
        switch v2 {
        case "foo bar":
            v84 := GenS7{F0: 2, F1: v2}
            _ = v84
        case "\n\t":
            switch (("" + (((v2 + (v2 + v2)) + (v2 + "foo bar")) + v2)) + "hello") {
            case "\n\t":
                v85 := GenS11{F0: (v9 | v9), F1: int8(v5), F2: 1}
                _ = v85
            case "":
                v3 = v3
            case "hello":
                _ = len(v4)
            }
        default:
            if true {
                if true {
                    var v86 int8 = 46
                    if false {
                        if (true || (v19 == (v19 / float32(1e10)))) {
                            switch v2 {
                            case "hello":
                                var v87 uint = (v5 & 442)
                                _ = v87
                            case "\n\t":
                                fmt.Println(v9, v5)
                            case "":
                                fmt.Println(v9, v1)
                            }
                            var v88 int8 = (v86 ^ (v86 << (((v5 - (v5 ^ v5))) & 63)))
                            v2 = v2
                            _ = v88
                        }
                        switch 0 {
                        case 42:
                            v89 := GenS12{F0: (^(v9))}
                            _ = v89
                        case 0:
                            fmt.Println(v9)
                        case -820:
                            _ = len(v4)
                        }
                    } else {
                        fmt.Println(v5)
                    }
                    _ = v86
                } else {
                    var v90 int = v1
                    v91 := GenS13{F0: 32767, F1: 76, F2: 0}
                    v5 = (v5 / (((v5 / (((v5 | (v5 >> ((65535) & 63)))) | 1))) | 1))
                    _ = v90
                    _ = v91
                }
                var v92 uint64 = 1
                if (132 < uint32(v5)) {
                    _ = len(v4)
                } else {
                    var v93 uint8 = (v3 + 128)
                    if false {
                        if (!(true || false)) {
                            v94 := GenS11{F0: (v9 << ((v5) & 63)), F1: 1, F2: 0}
                            _ = v94
                        } else {
                            switch v1 {
                            case -464:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            for fi95 := 0; fi95 < 20; fi95++ {
                                v93 = v3
                                _ = fi95
                            }
                        }
                        if (255 != 2) {
                            switch v2 {
                            case "":
                                if ((v5 | (v5 / ((897) | 1))) >= 4294967295) {
                                    var v96 int64 = 1
                                    var v97 int = (-(fi23))
                                    for fi98 := 0; fi98 < 1; fi98++ {
                                        fmt.Println(v5, v2)
                                        v99 := GenS14{F0: 1, F1: -9223372036854775808, F2: 255}
                                        fmt.Println(v92)
                                        if false {
                                            switch v2 {
                                            case "a":
                                                v100 := GenS15{F0: (v93 - v93)}
                                                v101 := GenS16{F0: 159}
                                                _ = v100
                                                _ = v101
                                            case "":
                                                v102 := GenS7{F0: 42, F1: "hello"}
                                                _ = v102
                                            case "\n\t":
                                                if (2 > (v5 + 612)) {
                                                    fmt.Println(v19, v5)
                                                } else {
                                                    var v103 uint8 = 42
                                                    var v104 bool = true
                                                    fmt.Println(v96, fi23)
                                                    _ = v103
                                                    _ = v104
                                                }
                                            default:
                                                var v105 string = v2
                                                _ = v105
                                            }
                                            v2 = (v2 + v2)
                                            switch (fi23 + v97) {
                                            case 128:
                                                _ = len(v4)
                                            }
                                        }
                                        var v106 uint = uint(v92)
                                        _ = fi98
                                        _ = v106
                                        _ = v99
                                    }
                                    _ = v96
                                    _ = v97
                                } else {
                                    var v107 int64 = (v9 ^ v9)
                                    switch (v2 + v2) {
                                    case "a":
                                        if true {
                                            if false {
                                                switch "a" {
                                                case "hello":
                                                    if false {
                                                        v108 := GenS1{F0: 1}
                                                        v3 += (v3 >> (((^(v5))) & 63))
                                                        _ = v108
                                                    } else {
                                                        var v109 uint16 = 32767
                                                        fmt.Println(v93, v92)
                                                        fmt.Println(fi23, v1)
                                                        _ = v109
                                                    }
                                                default:
                                                    var v110 uint8 = (v3 >> ((v5) & 63))
                                                    _ = v110
                                                }
                                                var v111 uint64 = (^(v92))
                                                _ = v111
                                            }
                                        }
                                    case "\n\t":
                                        switch v2 {
                                        case "\n\t":
                                            v5 -= 255
                                        case "hello":
                                            for fi112 := 0; fi112 < 1; fi112++ {
                                                if true {
                                                    var v113 uint32 = 303
                                                    _ = v113
                                                }
                                                fmt.Println(v19, v1, fi23)
                                                fmt.Println(fi112)
                                                _ = fi112
                                            }
                                        case "":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            var v114 uint32 = 65535
                                            _ = v114
                                        }
                                    default:
                                        var v115 string = "hello"
                                        _ = v115
                                    }
                                    _ = v107
                                }
                            case "a":
                                for fi116 := 0; fi116 < 17; fi116++ {
                                    var v117 int = 4294967295
                                    v118 := GenS17{F0: 0}
                                    fmt.Println(v92)
                                    _ = fi116
                                    _ = v117
                                    _ = v118
                                }
                            default:
                                fmt.Println(v93, v19)
                            }
                            v119 := GenS17{F0: 1}
                            _ = v119
                        } else {
                            for fi120 := 0; fi120 < 5; fi120++ {
                                if (false || false) {
                                    fmt.Println(fi120, fi120, v9)
                                } else {
                                    if false {
                                        switch v2 {
                                        case "":
                                            v19 = v19
                                        case "foo bar":
                                            fmt.Println(v93)
                                        case "\n\t":
                                            switch fi120 {
                                            case 4294967295:
                                                var v121 int64 = v9
                                                _ = v121
                                            default:
                                                if ((64 != 42) && (1.5 < (float64(v19) + 1.5))) {
                                                    var v122 int32 = 1
                                                    if false {
                                                        v123 := GenS18{F0: (v9 / ((v9) | 1)), F1: v9}
                                                        _ = v123
                                                    }
                                                    _ = v122
                                                }
                                            }
                                        }
                                        var v124 int32 = 558
                                        var v125 bool = (true || (("\n\t" + (v2 + v2)) == "hello"))
                                        _ = v124
                                        _ = v125
                                    } else {
                                        v4 = append(v4, 1)
                                        var v126 int16 = -32768
                                        _ = v126
                                    }
                                    var v127 int32 = 353
                                    v1 = (fi120 % (((v1 + v1)) | 1))
                                    _ = v127
                                }
                                _ = fi120
                            }
                        }
                    }
                    var v128 uint16 = 255
                    _ = v128
                    _ = v93
                }
                _ = v92
            } else {
                var v129 int16 = 1
                if false {
                    v130 := GenS12{F0: (v9 * v9)}
                    v5 += (v5 + (v5 >> ((v5) & 63)))
                    v4 = append(v4, fi23)
                    _ = v130
                }
                _ = v129
            }
        }
        _ = fi23
        _ = v83
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v131 := GenS19{F0: 2, F1: v1, F2: true}
    var v132 uint64 = 4294967295
    switch (v2 + "hello") {
    case "foo bar":
        if true {
            var v133 uint64 = (v132 / ((v132) | 1))
            v134 := GenS9{F0: -2147483648}
            _ = v133
            _ = v134
        }
    case "a":
        var v135 float64 = 3.14159
        _ = v135
    case "":
        v2 = (v2 + v2)
    default:
        if (!false) {
            if (v5 > v5) {
                var v136 float64 = 0.0
                switch v2 {
                case "foo bar":
                    switch "\n\t" {
                    case "hello":
                        v132 = (^(v132))
                    default:
                        var v137 int32 = -2147483648
                        _ = v137
                    }
                default:
                    var v138 string = v2
                    _ = v138
                }
                var v139 string = (v2 + v2)
                _ = v136
                _ = v139
            } else {
                if false {
                    v2 = v2
                    var v140 uint32 = 42
                    v1 = (^(v1))
                    _ = v140
                } else {
                    var v141 int64 = v9
                    _ = v141
                }
                if false {
                    if len(v4) > 0 {
                        v4[0] = v1
                    }
                    if (false && true) {
                        for fi142 := 0; fi142 < 19; fi142++ {
                            v19 = -2.25
                            var v143 int64 = v9
                            for fi144 := 0; fi144 < 20; fi144++ {
                                if len(v4) > 0 {
                                    v4[0] = (fi144 >> (((v5 % (((v5 / (((v5 ^ v5)) | 1))) | 1))) & 63))
                                }
                                _ = fi144
                            }
                            _ = fi142
                            _ = v143
                        }
                        v2 = ""
                        switch v2 {
                        case "":
                            v3 = (v3 << ((v5) & 63))
                        case "hello":
                            v145 := GenS5{F0: 0, F1: (v3 << (((^(v5))) & 63)), F2: 2147483647}
                            _ = v145
                        case "\n\t":
                            if (!false) {
                                var v146 float32 = (v19 * v19)
                                _ = v146
                            }
                        }
                    }
                } else {
                    for fi147 := 0; fi147 < 12; fi147++ {
                        _ = len(v4)
                        for fi148 := 0; fi148 < 6; fi148++ {
                            if (v2 <= (((v2 + "foo bar") + (v2 + v2)) + v2)) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            fmt.Println(v1, v132)
                            _ = fi148
                        }
                        for fi149 := 0; fi149 < 3; fi149++ {
                            var v150 uint16 = 65535
                            if len(v4) > 0 {
                                v4[0] = 32768
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi149
                            _ = v150
                        }
                        _ = fi147
                    }
                    v151 := GenS20{F0: uint32(v3), F1: false}
                    switch (v2 + "\n\t") {
                    case "hello":
                        _ = len(v4)
                    case "":
                        if (42 == 42) {
                            if (!(!false)) {
                                if false {
                                    switch "" {
                                    case "hello":
                                        v1 = int(v5)
                                    case "a":
                                        v152 := GenS21{F0: (^(v5)), F1: -1, F2: 18446744073709551615}
                                        _ = v152
                                    case "":
                                        v1 -= (v1 >> ((v5) & 63))
                                    default:
                                        if true {
                                            v3 += v3
                                        }
                                    }
                                } else {
                                    v132 = (v132 + v132)
                                    var v153 int16 = 42
                                    _ = v153
                                }
                                var v154 int64 = v9
                                var v155 float32 = 0.0
                                _ = v154
                                _ = v155
                            } else {
                                var v156 uint = v5
                                _ = v156
                            }
                            var v157 int16 = 32767
                            v5 = 908
                            _ = v157
                        } else {
                            if false {
                                switch v1 {
                                case 127:
                                    var v158 uint16 = 694
                                    _ = v158
                                case -178:
                                    for fi159 := 0; fi159 < 10; fi159++ {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v160 := GenS22{F0: v132}
                                        switch -2147483648 {
                                        case 1:
                                            if (((1e10 + 1e10) / 0.1) <= -2.25) {
                                                v161 := GenS2{F0: 100, F1: v132, F2: 1}
                                                var v162 uint8 = v3
                                                var v163 string = (v2 + v2)
                                                _ = v161
                                                _ = v162
                                                _ = v163
                                            } else {
                                                if false {
                                                    var v164 int = fi159
                                                    v2 = ("hello" + "\n\t")
                                                    _ = v164
                                                }
                                            }
                                        case 846:
                                            v1 = fi159
                                        case -469:
                                            for fi165 := 0; fi165 < 7; fi165++ {
                                                v166 := GenS23{F0: v3}
                                                _ = fi165
                                                _ = v166
                                            }
                                        }
                                        _ = fi159
                                        _ = v160
                                    }
                                case -698:
                                    v167 := GenS17{F0: 0}
                                    _ = v167
                                default:
                                    var v168 uint = v5
                                    _ = v168
                                }
                                var v169 int32 = -2147483648
                                v170 := GenS24{F0: (v3 ^ (v3 / ((255) | 1))), F1: 1, F2: "foo bar"}
                                _ = v169
                                _ = v170
                            }
                        }
                    default:
                        if true {
                            for fi171 := 0; fi171 < 14; fi171++ {
                                for fi172 := 0; fi172 < 18; fi172++ {
                                    fmt.Println(v3, v132, v19)
                                    switch 4294967295 {
                                    case 732:
                                        fmt.Println(v5)
                                        if (v1 > -1) {
                                            v9 -= -869
                                            v173 := GenS15{F0: 1}
                                            _ = v173
                                        }
                                        v4 = append(v4, fi171)
                                        v174 := GenS20{F0: 2147483647, F1: (-32768 >= 32767)}
                                        _ = v174
                                    }
                                    _ = fi172
                                }
                                _ = fi171
                            }
                            v175 := GenS2{F0: 100, F1: 1, F2: 65535}
                            fmt.Println(v2, v3)
                            _ = v175
                        }
                    }
                    _ = v151
                }
            }
            v9 = v9
            v5 -= (v5 | 4294967295)
        }
    }
    if true {
        v5 += 4294967295
        if ((^(v1)) != (-(v1))) {
            v3 -= uint8(v5)
            v19 = v19
            switch v2 {
            case "":
                if true {
                    v176 := GenS25{F0: 42}
                    var v177 int8 = 100
                    _ = v176
                    _ = v177
                }
            case "a":
                fmt.Println(v132)
            default:
                for fi178 := 0; fi178 < 9; fi178++ {
                    for fi179 := 0; fi179 < 6; fi179++ {
                        fmt.Println(fi179, v1)
                        _ = fi179
                    }
                    _ = fi178
                }
            }
        }
        if len(v4) > 0 {
            v4[0] = (v1 | 9223372036854775807)
        }
    } else {
        var v180 int16 = -219
        if (false && ((true || false) || (!(v9 <= v9)))) {
            v181 := GenS4{F0: 1}
            var v182 int = -1
            _ = v181
            _ = v182
        } else {
            var v183 uint16 = 1
            v9 = (v9 / ((v9) | 1))
            _ = len(v4)
            _ = v183
        }
        _ = v180
    }
    if true {
        v4 = append(v4, 32768)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v131
    _ = v132
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
