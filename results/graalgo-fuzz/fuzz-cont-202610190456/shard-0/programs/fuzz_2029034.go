package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 int; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 bool; F1 int32; F2 uint8 }
type GenS13 struct { F0 uint64; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = -2.25
    v3 := []int{369, 42, 1}
    v2 -= 0.0
    v2 = (v2 + (v2 - v2))
    v4 := GenS1{F0: -650, F1: false, F2: 32767}
    v2 = (v2 + 3.14159)
    var v5 int8 = 127
    v5 = -31
    if (127 != 155) {
        v6 := GenS1{F0: 127, F1: (1 >= 1), F2: 0}
        var v7 int64 = 0
        _ = v6
        _ = v7
    } else {
        if (1 < -386) {
            var v8 uint64 = 204
            v9 := GenS1{F0: 0, F1: false, F2: 32768}
            v10 := GenS1{F0: 127, F1: true, F2: 9223372036854775807}
            _ = v10
            _ = v8
            _ = v9
        } else {
            if (34 < 0) {
                v11 := GenS1{F0: 127, F1: true, F2: -9223372036854775808}
                _ = v11
            } else {
                var v12 uint64 = 4294967295
                switch 128 {
                case 32768:
                    v13 := GenS1{F0: 127, F1: (!(65535 == 1)), F2: 0}
                    _ = v13
                }
                _ = v12
            }
            if true {
                for fi14 := 0; fi14 < 20; fi14++ {
                    delete(v1, 913)
                    _ = len(v1)
                    v15 := GenS2{F0: 0}
                    _ = fi14
                    _ = v15
                }
                switch "foo bar" {
                case "a":
                    if true {
                        switch 128 {
                        case 9223372036854775807:
                            if true {
                                var v16 uint8 = 42
                                if ((-646 >= 0) || (!((1.5 > v2) && (!false)))) {
                                    fmt.Println(v2, v2, v5)
                                    var v17 float64 = v2
                                    if (1 != 65535) {
                                        for fi18 := 0; fi18 < 15; fi18++ {
                                            v19 := GenS1{F0: fi18, F1: (42 < 0), F2: (fi18 * 126)}
                                            var v20 int16 = 1
                                            v21 := GenS3{F0: 4294967295}
                                            _ = fi18
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        }
                                    }
                                    _ = v17
                                } else {
                                    if true {
                                        if false {
                                            var v22 float64 = v2
                                            if (!false) {
                                                for fi23 := 0; fi23 < 19; fi23++ {
                                                    var v24 uint32 = 60
                                                    v16 = v16
                                                    for fi25 := 0; fi25 < 3; fi25++ {
                                                        var v26 bool = true
                                                        var v27 int16 = -1
                                                        v27 = (v27 + -1)
                                                        var v28 int8 = (-(v5))
                                                        _ = fi25
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                    _ = fi23
                                                    _ = v24
                                                }
                                            } else {
                                                v29 := GenS1{F0: -1, F1: false, F2: -10}
                                                fmt.Println(v16, v5)
                                                _ = v29
                                            }
                                            _ = v22
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v30 := GenS1{F0: 1, F1: false, F2: 165}
                                            if (false && false) {
                                                for fi31 := 0; fi31 < 7; fi31++ {
                                                    for fi32 := 0; fi32 < 6; fi32++ {
                                                        fmt.Println(v5, fi32)
                                                        _ = fi32
                                                    }
                                                    v33 := GenS4{F0: (v16 - v16)}
                                                    _ = fi31
                                                    _ = v33
                                                }
                                            }
                                            _ = v30
                                        }
                                        var v34 uint8 = v16
                                        _ = v34
                                    }
                                }
                                _ = v16
                            } else {
                                for fi35 := 0; fi35 < 9; fi35++ {
                                    _ = len(v1)
                                    for fi36 := 0; fi36 < 15; fi36++ {
                                        v37 := GenS5{F0: 42, F1: 153, F2: -267}
                                        _ = fi36
                                        _ = v37
                                    }
                                    v5 = v5
                                    _ = fi35
                                }
                                switch "hello" {
                                case "":
                                    switch 4294967295 {
                                    case 493:
                                        v2 = v2
                                    case 9223372036854775807:
                                        fmt.Println(v2)
                                    }
                                case "a":
                                    v2 += 1.5
                                }
                            }
                        case -667:
                            var v38 int = int(v5)
                            _ = v38
                        case 206:
                            v2 = 0.0
                        }
                        switch -2147483648 {
                        case 369:
                            v5 = 116
                        case -9223372036854775808:
                            v2 = v2
                        default:
                            for fi39 := 0; fi39 < 3; fi39++ {
                                for fi40 := 0; fi40 < 7; fi40++ {
                                    v2 = ((((v2 - 1.5) * v2) - 1.5) - v2)
                                    v2 -= -2.25
                                    _ = fi40
                                }
                                var v41 uint16 = 587
                                v42 := GenS4{F0: 1}
                                _ = fi39
                                _ = v41
                                _ = v42
                            }
                        }
                    } else {
                        if (float32(v2) == float32((v2 / (1e10 / float64(1.5))))) {
                            if (908 < 7) {
                                var v43 int32 = 2147483647
                                var v44 int8 = v5
                                _ = v43
                                _ = v44
                            } else {
                                v5 += v5
                                fmt.Println(v5, v5)
                            }
                            v45 := GenS6{F0: (v5 << ((255) & 63)), F1: false}
                            _ = v45
                        }
                    }
                case "hello":
                    switch "\n\t" {
                    case "a":
                        v1[2147483647] = 559
                    case "hello":
                        for fi46 := 0; fi46 < 17; fi46++ {
                            v47 := GenS7{F0: 267}
                            _ = fi46
                            _ = v47
                        }
                    }
                default:
                    var v48 bool = false
                    _ = v48
                }
                var v49 float64 = ((v2 - v2) + 3.14159)
                _ = v49
            } else {
                fmt.Println(v5, v5)
            }
        }
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        if false {
            var v50 int32 = 2147483647
            switch (("hello" + (("\n\t" + ("hello" + "\n\t")) + "")) + "\n\t") {
            case "\n\t":
                v1[0] = 1
            case "hello":
                var v51 uint8 = 1
                _ = v51
            default:
                if true {
                    v3 = append(v3, 9223372036854775807)
                    fmt.Println(v5, v5, v2)
                    v2 += (3.14159 * v2)
                } else {
                    for fi52 := 0; fi52 < 19; fi52++ {
                        if (v2 <= ((0.1 - v2) * v2)) {
                            if false {
                                var v53 uint64 = 0
                                _ = v53
                            } else {
                                switch fi52 {
                                case 1:
                                    fmt.Println(fi52, v2)
                                }
                            }
                        } else {
                            v5 -= (v5 - v5)
                            v2 = v2
                        }
                        _ = len(v3)
                        if false {
                            if false {
                                v2 -= (-2.25 + v2)
                            } else {
                                if false {
                                    switch (fi52 << ((65535) & 63)) {
                                    case 42:
                                        var v54 uint32 = 175
                                        _ = v54
                                    case 0:
                                        var v55 int8 = 2
                                        _ = v55
                                    }
                                    var v56 uint8 = 0
                                    _ = v56
                                } else {
                                    v57 := GenS8{F0: 9223372036854775807, F1: (fi52 + fi52), F2: v5}
                                    _ = v57
                                }
                                v5 = v5
                                v58 := GenS9{F0: 128, F1: 65535}
                                _ = v58
                            }
                            if (false && (2 >= 2147483647)) {
                                var v59 string = "\n\t"
                                v59 = v59
                                _ = v59
                            }
                            if len(v3) > 0 {
                                v3[0] = -501
                            }
                        } else {
                            fmt.Println(fi52, v2, v50)
                            for fi60 := 0; fi60 < 16; fi60++ {
                                switch "hello" {
                                case "":
                                    var v61 uint16 = 42
                                    _ = v61
                                case "a":
                                    var v62 float32 = 0.1
                                    _ = v62
                                default:
                                    if true {
                                        v63 := GenS8{F0: 42, F1: fi60, F2: (v5 ^ (v5 % (((v5 + v5)) | 1)))}
                                        fmt.Println(v50, fi52, fi60)
                                        var v64 uint32 = 510
                                        v65 := GenS2{F0: (v5 >> ((0) & 63))}
                                        v66 := GenS9{F0: 0, F1: 251}
                                        v64 += 42
                                        v50 = v50
                                        if (255 > 0) {
                                            v5 = (v5 / (((-(v5))) | 1))
                                        }
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                    v50 += v50
                                    fmt.Println(v5)
                                    v67 := GenS6{F0: (^(v5)), F1: true}
                                    v68 := GenS3{F0: 2}
                                    switch "" {
                                    case "":
                                        if false {
                                            if ((fi60 + (fi60 | (fi60 + fi52))) < 32767) {
                                                if true {
                                                    var v69 uint32 = 568
                                                    var v70 uint16 = 42
                                                    fmt.Println(v69)
                                                    var v71 float32 = 0.0
                                                    _ = v69
                                                    _ = v70
                                                    _ = v71
                                                }
                                                v1[42] = 0
                                                v72 := GenS5{F0: 255, F1: 2, F2: -1}
                                                _ = v72
                                            }
                                        } else {
                                            v5 += v5
                                            v50 += (v50 & (v50 / ((1) | 1)))
                                        }
                                    case "foo bar":
                                        var v73 string = "hello"
                                        _ = v73
                                    case "a":
                                        if true {
                                            v74 := GenS3{F0: 827}
                                            v75 := GenS10{F0: 907}
                                            _ = v74
                                            _ = v75
                                        } else {
                                            if true {
                                                v3 = append(v3, 1)
                                                v50 -= (v50 ^ v50)
                                            }
                                            v50 += 581
                                            v76 := GenS5{F0: 604, F1: 42, F2: int32(v5)}
                                            _ = v76
                                        }
                                        var v77 uint = 237
                                        _ = v77
                                    }
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi60
                            }
                        }
                        _ = fi52
                    }
                    v3 = append(v3, 32768)
                }
            }
            _ = v50
        }
    }
    _ = len(v3)
    if ((!true) || ((float64(1.5) * v2) == (v2 + v2))) {
        if (2 != 4294967295) {
            var v78 int32 = 731
            _ = v78
        } else {
            var v79 uint16 = 32767
            for fi80 := 0; fi80 < 17; fi80++ {
                v5 = v5
                v1[4294967295] = 1
                _ = fi80
            }
            _ = v79
        }
    }
    for fi81 := 0; fi81 < 19; fi81++ {
        if false {
            var v82 float64 = (v2 - 1.5)
            if (-503 == 0) {
                for fi83 := 0; fi83 < 19; fi83++ {
                    if (((!(float32((float64(1e10) / 1.0)) > 1.5)) && (true && ((-1 > -2147483648) || (!(false || false))))) || (v5 <= (v5 ^ 127))) {
                        v84 := GenS10{F0: 0}
                        var v85 uint = 522
                        v82 += v2
                        _ = v84
                        _ = v85
                    }
                    _ = fi83
                }
                for fi86 := 0; fi86 < 5; fi86++ {
                    v5 += 0
                    _ = fi86
                }
                v82 = 1.5
            }
            _ = v82
        } else {
            for fi87 := 0; fi87 < 20; fi87++ {
                if (true || false) {
                    fmt.Println(fi87)
                    switch "foo bar" {
                    case "foo bar":
                        var v88 bool = ((!true) && true)
                        _ = v88
                    case "a":
                        var v89 uint64 = 105
                        _ = v89
                    case "hello":
                        fmt.Println(v5, fi87)
                    }
                    if true {
                        v90 := GenS4{F0: 2}
                        switch fi87 {
                        case 974:
                            v91 := GenS10{F0: 80}
                            _ = v91
                        case -9223372036854775808:
                            fmt.Println(v2, v5, v2)
                        default:
                            v2 -= (v2 * v2)
                        }
                        v2 = v2
                        _ = v90
                    } else {
                        fmt.Println(fi81, fi87)
                        var v92 int64 = 42
                        v92 += v92
                        if ((973 > 833) && true) {
                            _ = len(v1)
                            v5 -= (^(v5))
                        } else {
                            v1[0] = 582
                            v93 := GenS5{F0: 65535, F1: 1, F2: -1}
                            _ = v93
                        }
                        var v94 uint16 = 2
                        _ = v92
                        _ = v94
                    }
                    var v95 int64 = -221
                    _ = v95
                } else {
                    _ = len(v3)
                    switch "hello" {
                    case "foo bar":
                        v5 = (v5 / ((0) | 1))
                    case "a":
                        v5 = (v5 & v5)
                    default:
                        _ = len(v1)
                    }
                    v96 := GenS1{F0: (fi87 << ((278) & 63)), F1: true, F2: -2147483648}
                    _ = v96
                }
                var v97 string = "hello"
                v98 := GenS1{F0: fi81, F1: false, F2: -508}
                switch (fi81 / ((fi87) | 1)) {
                case -173:
                    v97 = "\n\t"
                }
                _ = fi87
                _ = v97
                _ = v98
            }
            v99 := GenS1{F0: (fi81 ^ (fi81 / ((128) | 1))), F1: true, F2: fi81}
            _ = v99
        }
        _ = fi81
    }
    for fi100 := 0; fi100 < 15; fi100++ {
        for fi101 := 0; fi101 < 8; fi101++ {
            if true {
                v5 += (v5 << ((2) & 63))
                if false {
                    v1[253] = 263
                    if true {
                        _ = len(v3)
                        delete(v1, 356)
                    } else {
                        var v102 uint16 = 32767
                        v5 -= v5
                        _ = v102
                    }
                }
                if ((v5 <= (v5 + 0)) && (false || false)) {
                    var v103 int = 32768
                    _ = v103
                }
                switch fi100 {
                case 637:
                    if true {
                        v3 = append(v3, 128)
                        switch (("hello" + "\n\t") + ("a" + "")) {
                        case "\n\t":
                            var v104 uint8 = 42
                            _ = v104
                        case "hello":
                            var v105 float32 = float32(1.5)
                            _ = v105
                        case "foo bar":
                            switch ("" + "a") {
                            case "hello":
                                var v106 string = "hello"
                                _ = v106
                            }
                        default:
                            var v107 float64 = float64(float32(1e10))
                            _ = v107
                        }
                        v2 = v2
                    } else {
                        var v108 uint64 = 1
                        _ = v108
                    }
                default:
                    v5 = (v5 >> ((2) & 63))
                }
                var v109 int64 = -1
                switch (-(fi100)) {
                case -1:
                    var v110 int8 = v5
                    _ = v110
                case 0:
                    var v111 uint8 = 128
                    _ = v111
                case 2147483647:
                    v1[0] = 1
                }
                if (false && (42 == 42)) {
                    v109 = (v109 - 0)
                } else {
                    var v112 uint = 305
                    var v113 bool = false
                    var v114 int64 = v109
                    _ = v112
                    _ = v113
                    _ = v114
                }
                _ = v109
            } else {
                v5 = (v5 >> ((1) & 63))
                fmt.Println(fi101, fi101)
            }
            switch "\n\t" {
            case "":
                fmt.Println(v2, v5)
            case "foo bar":
                v115 := GenS11{F0: (!((v2 != v2) || false))}
                _ = v115
            default:
                var v116 float64 = 1e10
                _ = v116
            }
            v3 = append(v3, fi100)
            _ = fi101
        }
        var v117 int = (fi100 / ((fi100) | 1))
        if false {
            v118 := GenS11{F0: (28 >= 548)}
            var v119 int8 = (v5 * int8(fi100))
            _ = v118
            _ = v119
        } else {
            var v120 int16 = -317
            var v121 int64 = 293
            fmt.Println(v121, v117, v2)
            _ = v120
            _ = v121
        }
        _ = fi100
        _ = v117
    }
    for fi122 := 0; fi122 < 2; fi122++ {
        var v123 bool = (0 >= 42)
        _ = fi122
        _ = v123
    }
    for fi124 := 0; fi124 < 3; fi124++ {
        var v125 uint16 = 1
        var v126 uint64 = uint64(v125)
        v126 = 2
        _ = fi124
        _ = v125
        _ = v126
    }
    switch int(v5) {
    case 1:
        v127 := GenS11{F0: (!(!true))}
        _ = v127
    case 32767:
        for fi128 := 0; fi128 < 7; fi128++ {
            if (!((!false) || true)) {
                for fi129 := 0; fi129 < 8; fi129++ {
                    fmt.Println(fi129, v5, fi129)
                    v2 -= ((float64(((-2.25 / -2.25) - 0.1)) - v2) + v2)
                    v2 += (v2 + (v2 + (v2 - 3.14159)))
                    _ = fi129
                }
                v2 = (3.14159 * v2)
            }
            if len(v3) > 0 {
                v3[0] = (fi128 / ((4294967295) | 1))
            }
            v130 := GenS12{F0: false, F1: 42, F2: 1}
            _ = fi128
            _ = v130
        }
    default:
        for fi131 := 0; fi131 < 15; fi131++ {
            v132 := GenS13{F0: 42, F1: 0}
            _ = fi131
            _ = v132
        }
    }
    fmt.Println(v2)
    if ((1 != 451) && false) {
        v133 := GenS3{F0: 42}
        _ = v133
    } else {
        v5 -= -128
        var v134 int = 4294967295
        _ = v134
    }
    v1[65535] = 927
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
