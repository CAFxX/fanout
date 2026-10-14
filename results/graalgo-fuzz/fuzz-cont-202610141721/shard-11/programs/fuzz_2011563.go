package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 string; F2 int }
type GenS9 struct { F0 uint16 }

func main() {
    var v1 string = "hello"
    var v2 uint64 = 0
    var v3 uint16 = 32767
    var v4 string = v1
    v2 += v2
    v4 = v1
    if false {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 float64 = -2.25
            fmt.Println(v2)
            v4 = "\n\t"
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint8 = 61
            v1 = v4
            var v9 uint = 4294967295
            _ = fi7
            _ = v8
            _ = v9
        }
        v4 = v1
    } else {
        if (!true) {
            v2 = 18446744073709551615
        }
        var v10 uint = 1
        v11 := []int{-2147483648, -696, 32767}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS1{F0: (^(v3)), F1: 0}
        if (3.14159 >= 0.1) {
            v14 := []int{(fi12 * 127), -365}
            _ = len(v14)
            _ = v14
        }
        v4 = ""
        _ = fi12
        _ = v13
    }
    v1 = (v1 + v1)
    var v15 string = v1
    if true {
        if (v15 < v4) {
            fmt.Println(v4, v3)
            v16 := []int{253}
            switch "a" {
            case "\n\t":
                var v17 uint16 = (^(v3))
                _ = v17
            case "foo bar":
                v1 = "a"
            case "hello":
                if true {
                    if (v2 > 1) {
                        var v18 float64 = (0.0 - (0.0 + (1e10 / -2.25)))
                        v3 += v3
                        if false {
                            fmt.Println(v4, v3, v2)
                            v16 = append(v16, 9223372036854775807)
                        }
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        for fi20 := 0; fi20 < 14; fi20++ {
                            if ((^(v2)) <= v2) {
                                var v21 uint64 = (v2 % ((v2) | 1))
                                v2 = 2
                                v16 = append(v16, fi20)
                                _ = v21
                            }
                            if false {
                                switch fi20 {
                                case -2147483648:
                                    if (-1 < -1) {
                                        v16 = append(v16, fi20)
                                    } else {
                                        v3 -= v3
                                        if false {
                                            fmt.Println(fi20)
                                            var v22 string = (v1 + "")
                                            switch (^(fi19)) {
                                            case 42:
                                                v23 := GenS1{F0: (v3 % ((v3) | 1)), F1: 2147483647}
                                                _ = v23
                                            case 9223372036854775807:
                                                var v24 int32 = 2147483647
                                                _ = v24
                                            case -32:
                                                if false {
                                                    var v25 bool = (fi19 == (fi20 << ((605) & 63)))
                                                    _ = v25
                                                }
                                            default:
                                                v4 = ((v15 + "\n\t") + v4)
                                            }
                                            _ = v22
                                        } else {
                                            if ((fi19 / (((^(fi20))) | 1)) < fi20) {
                                                v26 := GenS2{F0: v3, F1: (v2 >> ((171) & 63))}
                                                _ = v26
                                            } else {
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                            }
                                            v27 := GenS2{F0: 42, F1: 479}
                                            _ = len(v16)
                                            _ = v27
                                        }
                                        _ = len(v16)
                                        v28 := GenS3{F0: 856, F1: v2}
                                        if (uint16(v2) == (^(v3))) {
                                            var v29 uint64 = v2
                                            var v30 int16 = 32767
                                            switch (fi20 ^ (fi20 << ((1) & 63))) {
                                            case 32767:
                                                fmt.Println(v2, fi19, v30)
                                            default:
                                                if false {
                                                    var v31 bool = true
                                                    if len(v16) > 0 {
                                                        v16[0] = -1
                                                    }
                                                    _ = v31
                                                } else {
                                                    var v32 float64 = 1.5
                                                    if true {
                                                        switch (fi19 ^ fi20) {
                                                        case 128:
                                                            v33 := GenS1{F0: 2, F1: 194}
                                                            _ = v33
                                                        case 4294967295:
                                                            _ = len(v16)
                                                        }
                                                        v3 = (v3 + 1)
                                                    } else {
                                                        if false {
                                                            v16 = append(v16, fi20)
                                                        } else {
                                                            var v34 string = v1
                                                            fmt.Println(v1)
                                                            _ = v34
                                                        }
                                                        if (((fi19 >= fi20) || true) && (!true)) {
                                                            v30 = v30
                                                            var v35 float32 = float32((1e10 - float64(((float32(v32) * 1.5) * (1e10 + 3.14159)))))
                                                            switch (v15 + "a") {
                                                            case "hello":
                                                                v36 := GenS4{F0: (!false), F1: 673}
                                                                _ = v36
                                                            default:
                                                                switch (fi20 >> ((1) & 63)) {
                                                                case -9223372036854775808:
                                                                    if (4294967295 != 1) {
                                                                        v16 = append(v16, fi19)
                                                                        v37 := GenS4{F0: (-101 >= -128), F1: 0}
                                                                        _ = v37
                                                                    } else {
                                                                        _ = len(v16)
                                                                    }
                                                                case 42:
                                                                    if false {
                                                                        switch fi20 {
                                                                        case -1:
                                                                            v2 += (^(v29))
                                                                        case 1:
                                                                            fmt.Println(v15, fi20)
                                                                        }
                                                                        v38 := GenS3{F0: 1, F1: v29}
                                                                        var v39 int8 = 127
                                                                        v29 += v29
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                }
                                                            }
                                                            _ = v35
                                                        }
                                                    }
                                                    _ = v32
                                                }
                                            }
                                            _ = v29
                                            _ = v30
                                        }
                                        _ = v28
                                    }
                                }
                                var v40 int64 = 134
                                _ = v40
                            }
                            _ = fi20
                        }
                        v1 = v4
                        v41 := GenS3{F0: -78, F1: (^(v2))}
                        _ = fi19
                        _ = v41
                    }
                    for fi42 := 0; fi42 < 5; fi42++ {
                        for fi43 := 0; fi43 < 16; fi43++ {
                            v44 := GenS3{F0: 385, F1: v2}
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            v3 = (v3 + uint16(v2))
                            _ = fi43
                            _ = v44
                        }
                        v16 = append(v16, -723)
                        v2 += v2
                        _ = fi42
                    }
                }
            }
            _ = v16
        } else {
            if true {
                if (-477 <= 108) {
                    for fi45 := 0; fi45 < 19; fi45++ {
                        var v46 uint32 = 0
                        v46 -= v46
                        _ = fi45
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 10; fi47++ {
                        v48 := []int{(fi47 ^ 32768), 42, 0, (fi47 << ((255) & 63))}
                        _ = fi47
                        _ = v48
                    }
                    v15 = v15
                } else {
                    v4 = (v1 + (v15 + v1))
                    for fi49 := 0; fi49 < 18; fi49++ {
                        for fi50 := 0; fi50 < 1; fi50++ {
                            switch 2147483647 {
                            case 0:
                                if false {
                                    v4 = "foo bar"
                                } else {
                                    if false {
                                        v4 = v4
                                    }
                                    var v51 int = fi49
                                    _ = v51
                                }
                            }
                            switch fi50 {
                            case 128:
                                fmt.Println(fi50, v1, v4)
                                v52 := []int{4294967295, (fi50 - fi50), 1}
                                if true {
                                    switch v1 {
                                    case "a":
                                        v52 = append(v52, fi49)
                                    default:
                                        if true {
                                            var v53 float32 = -2.25
                                            var v54 uint8 = uint8(v2)
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                }
                                _ = v52
                            case 2147483647:
                                switch "" {
                                case "hello":
                                    var v55 uint8 = 255
                                    _ = v55
                                case "a":
                                    v2 -= v2
                                default:
                                    v56 := []int{fi49, 533, fi49}
                                    _ = v56
                                }
                            case -352:
                                v57 := []int{fi49}
                                _ = v57
                            }
                            switch 127 {
                            case -1:
                                v58 := GenS1{F0: v3, F1: 0}
                                _ = v58
                            }
                            _ = fi50
                        }
                        var v59 bool = false
                        _ = fi49
                        _ = v59
                    }
                    v2 += v2
                }
            } else {
                v60 := []int{386}
                v3 += v3
                var v61 int16 = -515
                _ = v60
                _ = v61
            }
            if (18446744073709551615 < 0) {
                for fi62 := 0; fi62 < 14; fi62++ {
                    v63 := GenS5{F0: (v2 >> ((uint(v3)) & 63)), F1: uint64(v3), F2: 979}
                    var v64 string = "hello"
                    v3 += (^(v3))
                    _ = fi62
                    _ = v63
                    _ = v64
                }
                fmt.Println(v15, v15)
            } else {
                var v65 uint = 4294967295
                for fi66 := 0; fi66 < 8; fi66++ {
                    if (!((false || false) && (2147483647 >= 1))) {
                        var v67 string = ""
                        v15 = ((("a" + v1) + v4) + v15)
                        _ = v67
                    }
                    var v68 int8 = 92
                    v1 = "foo bar"
                    _ = fi66
                    _ = v68
                }
                switch (v15 + ((((v15 + "a") + v4) + v1) + v1)) {
                case "hello":
                    for fi69 := 0; fi69 < 7; fi69++ {
                        for fi70 := 0; fi70 < 5; fi70++ {
                            if true {
                                v2 = 985
                                fmt.Println(v2)
                            }
                            v71 := []int{}
                            _ = fi70
                            _ = v71
                        }
                        _ = fi69
                    }
                default:
                    fmt.Println(v15, v15, v15)
                }
                _ = v65
            }
        }
        for fi72 := 0; fi72 < 13; fi72++ {
            for fi73 := 0; fi73 < 20; fi73++ {
                if (!false) {
                    var v74 int64 = -591
                    var v75 int = (-(fi72))
                    v15 = ((v4 + "\n\t") + "hello")
                    v76 := GenS5{F0: v2, F1: v2, F2: 294}
                    _ = v74
                    _ = v75
                    _ = v76
                } else {
                    var v77 uint64 = v2
                    v78 := GenS1{F0: v3, F1: 531}
                    var v79 bool = true
                    _ = v77
                    _ = v78
                    _ = v79
                }
                _ = fi73
            }
            fmt.Println(v1, v2)
            _ = fi72
        }
        if true {
            if (-235 == -1) {
                for fi80 := 0; fi80 < 12; fi80++ {
                    switch (fi80 | -991) {
                    case 2147483647:
                        v2 = 903
                    case -795:
                        switch -1 {
                        case 1:
                            switch ((v15 + "hello") + "") {
                            case "a":
                                v81 := []int{fi80, fi80}
                                _ = v81
                            case "\n\t":
                                v1 = (v1 + "hello")
                            }
                        default:
                            v3 -= (^(v3))
                        }
                    case 32767:
                        switch fi80 {
                        case -222:
                            v82 := GenS2{F0: (v3 % ((400) | 1)), F1: 2}
                            _ = v82
                        case 147:
                            v83 := []int{(fi80 - fi80)}
                            _ = v83
                        default:
                            if false {
                                v1 = v4
                            }
                        }
                    }
                    _ = fi80
                }
                fmt.Println(v1, v3)
            }
            var v84 uint8 = 2
            v85 := []int{}
            _ = v84
            _ = v85
        }
    }
    v2 -= v2
    if ((^(v3)) >= (v3 - (v3 - 32767))) {
        if false {
            if true {
                switch -387 {
                case -2147483648:
                    v86 := GenS1{F0: v3, F1: 2147483647}
                    _ = v86
                case 2147483647:
                    var v87 float32 = (-2.25 + (float32(0.1) + 1e10))
                    _ = v87
                case 32767:
                    switch int(v3) {
                    case -2147483648:
                        switch ((((("\n\t" + v1) + "a") + v15) + ("a" + "\n\t")) + (v4 + "")) {
                        case "a":
                            v88 := []int{32767}
                            _ = v88
                        case "foo bar":
                            v15 = v1
                        case "\n\t":
                            v89 := []int{-2147483648, 42, 1}
                            _ = v89
                        default:
                            v90 := []int{32768}
                            _ = v90
                        }
                    case -566:
                        v91 := GenS5{F0: v2, F1: (v2 ^ (v2 | v2)), F2: 2147483648}
                        _ = v91
                    default:
                        v4 = ""
                    }
                }
                for fi92 := 0; fi92 < 8; fi92++ {
                    var v93 uint = 290
                    switch v15 {
                    case "\n\t":
                        v93 += (v93 - (v93 | v93))
                    }
                    v94 := GenS4{F0: (!(1.5 > float64(-2.25))), F1: uint32(v2)}
                    _ = fi92
                    _ = v93
                    _ = v94
                }
                if false {
                    var v95 uint64 = v2
                    _ = v95
                }
            }
            switch 576 {
            case 42:
                var v96 int8 = -1
                _ = v96
            case 32767:
                v3 = (v3 & v3)
            default:
                v97 := GenS5{F0: (v2 | 2), F1: 42, F2: 2147483648}
                _ = v97
            }
            v15 = ("hello" + (("\n\t" + v4) + "foo bar"))
        } else {
            v3 -= (v3 ^ v3)
            var v98 int8 = -1
            _ = v98
        }
        v4 = ((("a" + "foo bar") + v4) + v4)
        var v99 uint16 = 42
        _ = v99
    } else {
        v100 := GenS6{F0: 842}
        if false {
            v2 = (v2 >> ((857) & 63))
            for fi101 := 0; fi101 < 3; fi101++ {
                switch "hello" {
                case "hello":
                    v102 := []int{}
                    _ = v102
                default:
                    fmt.Println(v2, v2)
                }
                if (true && true) {
                    fmt.Println(v4, v2)
                } else {
                    var v103 int8 = 26
                    v104 := []int{(fi101 / ((32768) | 1)), (fi101 >> ((2) & 63)), fi101, fi101}
                    v104 = append(v104, 9223372036854775807)
                    _ = v103
                    _ = v104
                }
                _ = fi101
            }
            if false {
                for fi105 := 0; fi105 < 7; fi105++ {
                    var v106 uint = 1
                    v1 = "foo bar"
                    _ = fi105
                    _ = v106
                }
            } else {
                for fi107 := 0; fi107 < 17; fi107++ {
                    v4 = ("foo bar" + v15)
                    var v108 uint16 = 1
                    switch 791 {
                    case 2147483647:
                        v2 = 858
                    case 850:
                        v109 := GenS7{F0: 38}
                        _ = v109
                    }
                    _ = fi107
                    _ = v108
                }
                var v110 float64 = -2.25
                _ = v110
            }
        } else {
            if (-70 >= -1) {
                if true {
                    fmt.Println(v3, v2, v1)
                } else {
                    v111 := GenS8{F0: 846, F1: ("foo bar" + ("hello" + "foo bar")), F2: 32768}
                    var v112 uint8 = 128
                    _ = v111
                    _ = v112
                }
                if true {
                    for fi113 := 0; fi113 < 3; fi113++ {
                        v114 := []int{-364, fi113}
                        if len(v114) > 0 {
                            v114[0] = (-(fi113))
                        }
                        v115 := GenS1{F0: (v3 ^ 464), F1: 380}
                        _ = fi113
                        _ = v114
                        _ = v115
                    }
                    for fi116 := 0; fi116 < 19; fi116++ {
                        if false {
                            var v117 uint32 = 65535
                            for fi118 := 0; fi118 < 4; fi118++ {
                                var v119 uint8 = 127
                                if true {
                                    v15 = v4
                                    v3 -= v3
                                    v3 -= (v3 & 65535)
                                    v120 := GenS9{F0: uint16(v117)}
                                    var v121 uint = 42
                                    _ = v120
                                    _ = v121
                                }
                                _ = fi118
                                _ = v119
                            }
                            v4 = v4
                            _ = v117
                        }
                        v122 := []int{(-(fi116))}
                        _ = fi116
                        _ = v122
                    }
                    if false {
                        v123 := []int{}
                        _ = v123
                    }
                }
            }
        }
        _ = v100
    }
    switch "foo bar" {
    case "foo bar":
        for fi124 := 0; fi124 < 13; fi124++ {
            v125 := []int{fi124, (-(fi124))}
            v126 := GenS9{F0: (v3 + (v3 | v3))}
            fmt.Println(v3, v1, fi124)
            _ = fi124
            _ = v125
            _ = v126
        }
    case "\n\t":
        if false {
            v4 = ("a" + ("hello" + "\n\t"))
            v15 = "\n\t"
        }
    default:
        v127 := []int{}
        _ = v127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
