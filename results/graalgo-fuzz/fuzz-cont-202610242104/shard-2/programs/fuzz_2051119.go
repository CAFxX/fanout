package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 bool; F2 int8 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int32; F1 int8; F2 bool }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint32; F1 int8; F2 int }
type GenS13 struct { F0 int8; F1 string; F2 uint32 }
type GenS14 struct { F0 bool; F1 int8; F2 bool }
type GenS15 struct { F0 bool; F1 int }
func genFn0(p0 uint16) int64 {
    v4 := []int{76}
    fmt.Println(p0)
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint64 = 2
        if false {
            switch "hello" {
            case "":
                _ = len(v4)
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "foo bar":
                fmt.Println(v6, p0, v6)
            }
        } else {
            for fi7 := 0; fi7 < 10; fi7++ {
                p0 = 65535
                _ = fi7
            }
            for fi8 := 0; fi8 < 3; fi8++ {
                fmt.Println(p0)
                _ = fi8
            }
        }
        if true {
            p0 += p0
            if false {
                if (false || (!false)) {
                    v9 := GenS1{F0: 1, F1: fi5}
                    _ = v9
                } else {
                    v10 := GenS1{F0: 0, F1: (fi5 << ((65535) & 63))}
                    var v11 float32 = -2.25
                    if ((((((float64((v11 * (v11 - v11))) * 3.14159) + (float64(v11) + (1.5 - float64(((v11 + v11) - 1e10))))) * 1.5) * 0.1) + 1e10) == 0.0) {
                        var v12 float64 = ((3.14159 / 1.0) - float64(v11))
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 uint8 = 128
                            switch 0 {
                            case 0:
                                if (v6 >= v6) {
                                    fmt.Println(fi5, fi13)
                                    fmt.Println(fi13, v11)
                                    if true {
                                        fmt.Println(fi13)
                                        fmt.Println(v11, p0)
                                        switch 42 {
                                        case 143:
                                            if len(v4) > 0 {
                                                v4[0] = fi13
                                            }
                                        }
                                        fmt.Println(fi5)
                                    } else {
                                        if (-564 > 509) {
                                            v11 = 1e10
                                        }
                                        var v15 uint64 = v6
                                        _ = v15
                                    }
                                }
                            }
                            _ = fi13
                            _ = v14
                        }
                        if false {
                            if len(v4) > 0 {
                                v4[0] = fi5
                            }
                            switch (("hello" + "\n\t") + "") {
                            case "hello":
                                v4 = append(v4, fi5)
                            case "\n\t":
                                p0 = (p0 + 0)
                            case "":
                                if (!true) {
                                    switch ("" + "a") {
                                    case "":
                                        if true {
                                            var v16 float64 = 3.14159
                                            var v17 bool = (2147483647 >= 2147483648)
                                            _ = v16
                                            _ = v17
                                        }
                                    case "foo bar":
                                        v4 = append(v4, (^(fi5)))
                                    }
                                    v18 := GenS2{F0: -2147483648}
                                    fmt.Println(v6)
                                    _ = v18
                                } else {
                                    var v19 uint32 = 4294967295
                                    for fi20 := 0; fi20 < 16; fi20++ {
                                        fmt.Println(v11, v19)
                                        _ = fi20
                                    }
                                    _ = v19
                                }
                            }
                        } else {
                            switch (-(fi5)) {
                            case 128:
                                if false {
                                    for fi21 := 0; fi21 < 18; fi21++ {
                                        v22 := GenS3{F0: 717}
                                        if (-1 <= -1) {
                                            v6 += (v6 + v6)
                                        } else {
                                            var v23 string = "foo bar"
                                            v24 := GenS3{F0: 4294967295}
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = fi21
                                        _ = v22
                                    }
                                }
                            case 42:
                                var v25 float32 = 0.1
                                _ = v25
                            case -48:
                                if (float64((v11 + ((0.1 / v11) - v11))) < v12) {
                                    var v26 uint = 923
                                    v27 := GenS2{F0: -1}
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 uint16 = p0
                                    v11 = (float32((((-2.25 / 0.1) * v12) / 1e10)) / v11)
                                    _ = v28
                                }
                            }
                        }
                        _ = v12
                    } else {
                        _ = len(v4)
                        switch (fi5 >> ((18446744073709551615) & 63)) {
                        case 0:
                            var v29 int16 = 317
                            _ = v29
                        case -935:
                            if true {
                                for fi30 := 0; fi30 < 16; fi30++ {
                                    if len(v4) > 0 {
                                        v4[0] = fi5
                                    }
                                    _ = fi30
                                }
                            } else {
                                switch 505 {
                                case 696:
                                    switch "\n\t" {
                                    case "a":
                                        _ = len(v4)
                                    case "hello":
                                        v31 := GenS3{F0: 722}
                                        _ = v31
                                    default:
                                        v32 := GenS1{F0: 42, F1: -9223372036854775808}
                                        _ = v32
                                    }
                                case 391:
                                    fmt.Println(p0, p0, v11)
                                case -9223372036854775808:
                                    for fi33 := 0; fi33 < 19; fi33++ {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v6 = 0
                                        v6 -= (v6 - 18446744073709551615)
                                        _ = fi33
                                    }
                                default:
                                    if (1 == -2147483648) {
                                        p0 = (p0 - (p0 & (^(p0))))
                                        fmt.Println(p0, v6, v11)
                                    } else {
                                        if (((!("a" == "")) && ((float64(v11) <= (1.5 / ((float64(v11) / float64(float32(float64(v11)))) - 0.1))) && ((^(v6)) <= v6))) || ((!false) || (fi5 < (^(fi5))))) {
                                            var v34 int8 = -1
                                            _ = v34
                                        } else {
                                            for fi35 := 0; fi35 < 14; fi35++ {
                                                switch (fi5 + -9223372036854775808) {
                                                case -591:
                                                    var v36 uint8 = 2
                                                    _ = v36
                                                case 4294967295:
                                                    var v37 int = (-(fi35))
                                                    _ = v37
                                                }
                                                _ = fi35
                                            }
                                            p0 = (p0 >> ((846) & 63))
                                            fmt.Println(v6, v11)
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            for fi38 := 0; fi38 < 19; fi38++ {
                                                if true {
                                                    switch ("\n\t" + "a") {
                                                    case "hello":
                                                        if ((v11 < (0.0 - 3.14159)) && ("hello" == "foo bar")) {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            fmt.Println(v6)
                                                            var v39 float64 = -2.25
                                                            _ = v39
                                                        } else {
                                                            var v40 uint = 18446744073709551615
                                                            _ = v40
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            if true {
                                                                var v41 uint32 = uint32(p0)
                                                                _ = v41
                                                            } else {
                                                                var v42 uint64 = 42
                                                                _ = v42
                                                            }
                                                            if false {
                                                                var v43 uint64 = v6
                                                                v44 := GenS2{F0: -1}
                                                                _ = v43
                                                                _ = v44
                                                            } else {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (fi38 | fi38)
                                                                    }
                                                                case "":
                                                                    var v45 uint16 = p0
                                                                    _ = v45
                                                                default:
                                                                    var v46 uint64 = (v6 % (((v6 ^ v6)) | 1))
                                                                    _ = v46
                                                                }
                                                                if true {
                                                                    if true {
                                                                        var v47 uint = 2
                                                                        v6 -= (v6 - (^(v6)))
                                                                        _ = v47
                                                                    } else {
                                                                        switch fi38 {
                                                                        case -1:
                                                                            p0 = 2
                                                                        }
                                                                        p0 += p0
                                                                    }
                                                                } else {
                                                                    fmt.Println(fi5)
                                                                }
                                                                v4 = append(v4, fi38)
                                                            }
                                                        }
                                                    case "\n\t":
                                                        p0 += 428
                                                    default:
                                                        v4 = append(v4, (fi5 - fi5))
                                                    }
                                                }
                                                _ = len(v4)
                                                fmt.Println(fi5, v11, v6)
                                                _ = fi38
                                            }
                                        }
                                        var v48 uint8 = 42
                                        _ = v48
                                    }
                                }
                            }
                        case 4294967295:
                            if (!(42 != 0)) {
                                v6 = 4294967295
                            } else {
                                v49 := GenS2{F0: 1}
                                switch "a" {
                                case "":
                                    v11 += v11
                                case "a":
                                    var v50 float32 = v11
                                    _ = v50
                                default:
                                    p0 -= (p0 ^ p0)
                                }
                                _ = v49
                            }
                        default:
                            _ = len(v4)
                        }
                        if (428 == -650) {
                            fmt.Println(p0)
                            if true {
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v11, v6)
                                case "\n\t":
                                    v6 -= v6
                                case "a":
                                    v11 += 0.0
                                }
                                v6 += v6
                                switch "" {
                                case "\n\t":
                                    for fi51 := 0; fi51 < 13; fi51++ {
                                        v52 := GenS4{F0: v6, F1: true, F2: 100}
                                        fmt.Println(v11)
                                        _ = fi51
                                        _ = v52
                                    }
                                case "hello":
                                    for fi53 := 0; fi53 < 6; fi53++ {
                                        v54 := GenS4{F0: (v6 << ((255) & 63)), F1: true, F2: 31}
                                        var v55 uint16 = p0
                                        _ = fi53
                                        _ = v54
                                        _ = v55
                                    }
                                default:
                                    switch fi5 {
                                    case -693:
                                        v56 := GenS5{F0: -9223372036854775808, F1: 508}
                                        _ = v56
                                    default:
                                        v57 := GenS3{F0: 42}
                                        _ = v57
                                    }
                                }
                            } else {
                                var v58 uint64 = (v6 * (v6 ^ (v6 & v6)))
                                var v59 uint16 = p0
                                switch fi5 {
                                case -990:
                                    var v60 int = fi5
                                    _ = v60
                                }
                                _ = v58
                                _ = v59
                            }
                            if (!true) {
                                switch 110 {
                                case 2147483647:
                                    if false {
                                        var v61 string = ("" + "")
                                        v11 = v11
                                        switch v61 {
                                        case "foo bar":
                                            v62 := GenS2{F0: 42}
                                            _ = v62
                                        case "a":
                                            if (v61 != v61) {
                                                v61 = "a"
                                                if len(v4) > 0 {
                                                    v4[0] = fi5
                                                }
                                                for fi63 := 0; fi63 < 16; fi63++ {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    fmt.Println(fi63, v6)
                                                    if (1 == 247) {
                                                        v61 = (v61 + "a")
                                                        if ((p0 | p0) <= 2) {
                                                            var v64 uint16 = 65535
                                                            _ = v64
                                                        }
                                                        var v65 string = ((("\n\t" + "foo bar") + v61) + v61)
                                                        _ = v65
                                                    } else {
                                                        v66 := GenS6{F0: 1, F1: 688, F2: 42}
                                                        _ = v66
                                                    }
                                                    v6 -= (v6 / ((0) | 1))
                                                    _ = fi63
                                                }
                                            }
                                        }
                                        _ = v61
                                    }
                                default:
                                    switch fi5 {
                                    case 4294967295:
                                        v6 = v6
                                    case -2147483648:
                                        v67 := GenS7{F0: 700}
                                        _ = v67
                                    default:
                                        v11 = -2.25
                                    }
                                }
                                v11 -= v11
                                var v68 int32 = 16
                                _ = v68
                            } else {
                                for fi69 := 0; fi69 < 14; fi69++ {
                                    if (false && true) {
                                        v6 += 2
                                        var v70 uint = 255
                                        fmt.Println(fi5, p0, fi5)
                                        _ = v70
                                    } else {
                                        v71 := GenS5{F0: -9223372036854775808, F1: 42}
                                        if len(v4) > 0 {
                                            v4[0] = fi69
                                        }
                                        _ = v71
                                    }
                                    var v72 int32 = 42
                                    _ = fi69
                                    _ = v72
                                }
                                if false {
                                    switch (^(fi5)) {
                                    case -673:
                                        for fi73 := 0; fi73 < 11; fi73++ {
                                            if false {
                                                if true {
                                                    switch (fi5 & (fi73 % ((-971) | 1))) {
                                                    case 169:
                                                        if (0 > -217) {
                                                            v74 := GenS3{F0: 477}
                                                            v75 := GenS8{F0: int32(v6), F1: -50, F2: true}
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                    }
                                                } else {
                                                    v76 := GenS9{F0: 42}
                                                    p0 = 123
                                                    _ = v76
                                                }
                                                var v77 string = ""
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v77
                                            }
                                            switch 2147483647 {
                                            case 42:
                                                if ((!true) && true) {
                                                    p0 += 106
                                                }
                                            }
                                            p0 -= p0
                                            _ = fi73
                                        }
                                    }
                                    var v78 int = (fi5 & 42)
                                    if false {
                                        fmt.Println(v78)
                                        v79 := GenS6{F0: 640, F1: 4294967295, F2: v6}
                                        if ((42 > 42) || false) {
                                            if (true && false) {
                                                v80 := GenS3{F0: 2}
                                                switch ("" + ("a" + "\n\t")) {
                                                case "a":
                                                    switch ((("a" + ("a" + "foo bar")) + "hello") + (("foo bar" + (("" + "hello") + "foo bar")) + "")) {
                                                    case "foo bar":
                                                        for fi81 := 0; fi81 < 18; fi81++ {
                                                            _ = len(v4)
                                                            var v82 uint64 = (v6 * (^(v6)))
                                                            _ = fi81
                                                            _ = v82
                                                        }
                                                    case "\n\t":
                                                        var v83 int64 = 9223372036854775807
                                                        _ = v83
                                                    default:
                                                        v84 := GenS5{F0: -588, F1: 42}
                                                        _ = v84
                                                    }
                                                case "hello":
                                                    switch "hello" {
                                                    case "a":
                                                        p0 = (^(p0))
                                                    case "hello":
                                                        if (0 > 9223372036854775807) {
                                                            var v85 uint64 = v6
                                                            _ = len(v4)
                                                            v85 = v6
                                                            _ = v85
                                                        } else {
                                                            var v86 int16 = -1
                                                            _ = v86
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v87 int16 = 1
                                                    _ = v87
                                                }
                                                _ = v80
                                            }
                                        }
                                        _ = v79
                                    } else {
                                        var v88 float64 = ((((float64(0.0) / ((1e10 + (1e10 * (1.5 - 3.14159))) + ((float64((float32((0.1 * (1e10 * 1.5))) * ((0.0 * v11) - 3.14159))) * -2.25) + 1e10))) - -2.25) * ((0.0 / 1.5) + 0.1)) / 0.1)
                                        _ = v88
                                    }
                                    _ = v78
                                } else {
                                    var v89 bool = (!false)
                                    p0 = 876
                                    var v90 int64 = 0
                                    _ = v89
                                    _ = v90
                                }
                                v6 -= v6
                            }
                        }
                    }
                    _ = v10
                    _ = v11
                }
            } else {
                v4 = append(v4, fi5)
                v91 := GenS10{F0: p0}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v91
            }
        }
        _ = fi5
        _ = v6
    }
    var v92 uint64 = 0
    _ = p0
    _ = v4
    _ = v92
    return -1
}
func genFn1(p0 bool, p1 string) int16 {
    if false {
        v93 := []int{983, 69, 32767, 127}
        if p0 {
            var v94 uint64 = 554
            var v95 int64 = 342
            _ = v94
            _ = v95
        } else {
            switch p1 {
            case "\n\t":
                var v96 string = ""
                _ = v96
            case "foo bar":
                var v97 float64 = 0.1
                _ = v97
            default:
                v93 = append(v93, 0)
            }
            if len(v93) >= 2 {
                v93 = v93[1:len(v93)]
            }
        }
        _ = v93
    } else {
        var v98 uint = 1
        _ = v98
    }
    var v99 uint16 = 29
    v100 := GenS5{F0: 0, F1: -846}
    for fi101 := 0; fi101 < 4; fi101++ {
        var v102 float32 = 1e10
        _ = fi101
        _ = v102
    }
    _ = p0
    _ = p1
    _ = v100
    _ = v99
    return 435
}

func main() {
    var v1 float32 = ((1e10 / 1.0) + (0.0 - (float32(1e10) * (float32((3.14159 - 3.14159)) * 0.0))))
    var v2 bool = false
    var v3 float32 = 0.1
    for fi103 := 0; fi103 < 8; fi103++ {
        v3 += v3
        v1 -= 1e10
        _ = fi103
    }
    switch "a" {
    case "hello":
        var v104 uint16 = 255
        _ = v104
    case "foo bar":
        for fi105 := 0; fi105 < 7; fi105++ {
            for fi106 := 0; fi106 < 9; fi106++ {
                var v107 int32 = 0
                v108 := []int{-1, -565}
                v1 = ((v3 * v3) + v3)
                _ = fi106
                _ = v107
                _ = v108
            }
            v109 := GenS4{F0: 18446744073709551615, F1: (!false), F2: 1}
            _ = fi105
            _ = v109
        }
    default:
        var v110 int8 = -5
        _ = v110
    }
    v111 := GenS11{F0: 32767}
    for fi112 := 0; fi112 < 15; fi112++ {
        v113 := []int{fi112, (fi112 << ((255) & 63)), (fi112 + (fi112 % (((fi112 % (((fi112 << ((255) & 63))) | 1))) | 1))), fi112}
        _ = fi112
        _ = v113
    }
    fmt.Println(v1, v2)
    switch 426 {
    case 4294967295:
        fmt.Println(v3)
    default:
        var v114 uint32 = 42
        _ = v114
    }
    fmt.Println(v2, v3, v2)
    v115 := []int{127, 1, 463, 4294967295}
    if len(v115) > 0 {
        v115[0] = 128
    }
    if ((-128 <= 117) && v2) {
        switch -614 {
        case 1:
            fmt.Println(v3, v1)
        default:
            v116 := GenS12{F0: 1, F1: 20, F2: 42}
            _ = v116
        }
        var v117 int8 = 102
        _ = v117
    }
    if (845 <= 2147483647) {
        if v2 {
            v1 -= v3
            fmt.Println(v1, v2, v2)
        }
    } else {
        v118 := GenS13{F0: -1, F1: "a", F2: 2147483647}
        if v2 {
            var v119 uint32 = 2147483648
            _ = v119
        }
        _ = v118
    }
    var v120 uint = 899
    var v121 uint64 = 0
    switch 0 {
    case -2147483648:
        if v2 {
            var v122 uint8 = 126
            var v123 uint16 = 273
            _ = v122
            _ = v123
        } else {
            switch 4294967295 {
            case 1:
                if (1e10 != ((float64((v1 * v1)) * float64((((v1 + 1e10) * -2.25) - (v3 / v3)))) * 3.14159)) {
                    var v124 uint64 = (v121 * v121)
                    _ = v124
                } else {
                    fmt.Println(v121, v2, v2)
                    for fi125 := 0; fi125 < 20; fi125++ {
                        v121 += 935
                        v1 -= 1e10
                        v120 -= (v120 | v120)
                        _ = fi125
                    }
                    if true {
                        for fi126 := 0; fi126 < 17; fi126++ {
                            var v127 int8 = 43
                            if (42 == 1) {
                                var v128 int8 = v127
                                _ = v128
                            }
                            _ = fi126
                            _ = v127
                        }
                        if (float32(1.5) > ((0.1 / 0.1) * ((float32((-2.25 - 3.14159)) / (((float32((0.0 * float64(v1))) - v3) * float32(1.5)) / 1.0)) - (3.14159 * (3.14159 - v3))))) {
                            var v129 int16 = int16(v120)
                            fmt.Println(v1, v3)
                            _ = v129
                        }
                        fmt.Println(v1)
                    } else {
                        v130 := GenS14{F0: (v2 || (!(255 < 241))), F1: 100, F2: ((v2 || (v2 && (true && (!(!false))))) && (false || v2))}
                        var v131 float64 = (3.14159 - 0.1)
                        v131 = v131
                        _ = v130
                        _ = v131
                    }
                }
            case -159:
                for fi132 := 0; fi132 < 8; fi132++ {
                    v115 = append(v115, (fi132 ^ fi132))
                    _ = fi132
                }
            }
        }
    }
    for fi133 := 0; fi133 < 19; fi133++ {
        if (42 == 79) {
            if v2 {
                fmt.Println(v121, v121)
                v2 = false
                v115 = append(v115, (fi133 >> ((v120) & 63)))
            }
            var v134 int64 = -583
            v135 := GenS15{F0: ((fi133 % ((fi133) | 1)) >= 32768), F1: fi133}
            _ = v134
            _ = v135
        } else {
            v3 = 1e10
            var v136 int32 = 2147483647
            var v137 uint32 = 1
            _ = v136
            _ = v137
        }
        if (v2 || (!true)) {
            if v2 {
                var v138 uint64 = (v121 | (v121 * (v121 % ((2) | 1))))
                _ = v138
            }
        }
        if ((fi133 & fi133) >= 42) {
            v115 = append(v115, (-(fi133)))
            v139 := GenS9{F0: uint8(v120)}
            v121 -= v121
            _ = v139
        }
        _ = fi133
    }
    var v140 float32 = v1
    v121 -= 121
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v115:", len(v115), v115)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v111
    _ = v115
    _ = v120
    _ = v121
    _ = v140
    _ = v2
    _ = v3
}
