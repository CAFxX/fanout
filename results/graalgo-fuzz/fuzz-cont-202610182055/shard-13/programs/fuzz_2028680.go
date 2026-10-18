package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 uint8; F2 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8; F1 uint32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint32; F1 bool }
type GenS14 struct { F0 bool; F1 uint16 }
type GenS15 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 uint }
func genFn0(p0 uint8) bool {
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{9223372036854775807, (^(fi7))}
        fmt.Println(p0, p0)
        _ = fi7
        _ = v8
    }
    var v9 int16 = 857
    _ = p0
    _ = v6
    _ = v9
    return false
}
func genFn1(p0 uint64) string {
    var v10 int8 = 68
    v11 := GenS2{F0: 4294967295, F1: v10, F2: 2147483647}
    if false {
        if true {
            var v12 int16 = 1000
            v12 = v12
            if true {
                var v13 float32 = (1.5 / float32((1e10 - (3.14159 / 1e10))))
                fmt.Println(v12)
                _ = v13
            } else {
                for fi14 := 0; fi14 < 11; fi14++ {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v16 := GenS3{F0: ("foo bar" + ("a" + "hello"))}
                        _ = fi15
                        _ = v16
                    }
                    var v17 uint16 = 1
                    _ = fi14
                    _ = v17
                }
                v10 = 43
                var v18 int64 = -1
                _ = v18
            }
            _ = v12
        }
        p0 = p0
    }
    _ = p0
    _ = v10
    _ = v11
    return "a"
}

func main() {
    v1 := []int{}
    var v2 int64 = -190
    v3 := make(map[uint32]int)
    var v4 uint64 = 42
    v5 := make(map[int]int)
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := GenS4{F0: 220, F1: 32767}
        switch fi19 {
        case 127:
            v4 -= v4
        case 42:
            fmt.Println(v4)
        case 32767:
            if (false || (((((0.1 * 1e10) - -2.25) <= 0.0) || (v4 >= (^(v4)))) && (0 < 911))) {
                var v21 int8 = -128
                var v22 float64 = -2.25
                if (false && false) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v21
                _ = v22
            } else {
                v23 := GenS1{F0: 160}
                var v24 float64 = (1.5 - (((1e10 - (float64((1e10 + 1e10)) + 0.0)) / float64(float32((1e10 - (1.5 - -2.25))))) + -2.25))
                _ = v23
                _ = v24
            }
        default:
            var v25 int16 = -32768
            _ = v25
        }
        _ = fi19
        _ = v20
    }
    v4 = v4
    v2 += (v2 & v2)
    if true {
        if (32767 <= -2147483648) {
            for fi26 := 0; fi26 < 18; fi26++ {
                if false {
                    var v27 int32 = 42
                    var v28 uint = 626
                    if (!false) {
                        if (32767 >= 65535) {
                            v28 = (v28 - v28)
                            var v29 uint = v28
                            if true {
                                for fi30 := 0; fi30 < 3; fi30++ {
                                    if true {
                                        var v31 uint = (v29 ^ (v28 ^ 18446744073709551615))
                                        _ = v31
                                    }
                                    _ = fi30
                                }
                                v28 -= 255
                                if false {
                                    fmt.Println(v28, v27)
                                } else {
                                    v32 := GenS1{F0: 872}
                                    if true {
                                        var v33 uint = v28
                                        fmt.Println(fi26)
                                        _ = v33
                                    }
                                    fmt.Println(v29, v29)
                                    _ = v32
                                }
                            } else {
                                v29 = (v29 ^ v28)
                                var v34 int = 2147483647
                                _ = v34
                            }
                            _ = v29
                        }
                        if false {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v4)
                            }
                            for fi35 := 0; fi35 < 3; fi35++ {
                                v2 -= v2
                                v2 += (v2 >> ((v28) & 63))
                                var v36 string = ""
                                _ = fi35
                                _ = v36
                            }
                            switch ("a" + "foo bar") {
                            case "hello":
                                v28 = v28
                            case "\n\t":
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        for fi37 := 0; fi37 < 4; fi37++ {
                                            fmt.Println(fi37)
                                            _ = fi37
                                        }
                                    case "foo bar":
                                        for fi38 := 0; fi38 < 8; fi38++ {
                                            _ = len(v5)
                                            _ = fi38
                                        }
                                    }
                                    if false {
                                        v39 := GenS4{F0: 2147483647, F1: 1}
                                        _ = v39
                                    }
                                    var v40 int16 = 32767
                                    _ = v40
                                } else {
                                    if true {
                                        v4 = (v4 ^ v4)
                                        var v41 float32 = 0.1
                                        switch ("foo bar" + "") {
                                        case "a":
                                            if true {
                                                var v42 uint64 = (v4 * (v4 & v4))
                                                v43 := GenS3{F0: (("foo bar" + "hello") + (("" + ("foo bar" + (("hello" + "") + "foo bar"))) + ("a" + ("\n\t" + "a"))))}
                                                if (!(24 == 127)) {
                                                    v44 := GenS5{F0: false, F1: (v27 + (v27 | -2147483648))}
                                                    _ = v44
                                                } else {
                                                    if true {
                                                        fmt.Println(v28)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v45 string = "hello"
                                                        _ = v45
                                                    } else {
                                                        v41 += -2.25
                                                        var v46 int64 = (v2 ^ v2)
                                                        _ = v46
                                                    }
                                                    switch 2147483647 {
                                                    case 62:
                                                        var v47 float64 = 0.1
                                                        _ = v47
                                                    case 776:
                                                        fmt.Println(v2)
                                                    case -457:
                                                        fmt.Println(v42, fi26)
                                                    }
                                                }
                                                _ = v42
                                                _ = v43
                                            }
                                        case "hello":
                                            switch (fi26 << ((v28) & 63)) {
                                            case 790:
                                                switch fi26 {
                                                case 9223372036854775807:
                                                    var v48 int = (fi26 ^ (^(fi26)))
                                                    _ = v48
                                                case 128:
                                                    var v49 float64 = 1e10
                                                    _ = v49
                                                default:
                                                    for fi50 := 0; fi50 < 14; fi50++ {
                                                        v1 = append(v1, fi26)
                                                        _ = len(v1)
                                                        _ = fi50
                                                    }
                                                }
                                            case -764:
                                                for fi51 := 0; fi51 < 20; fi51++ {
                                                    if false {
                                                        v52 := v3[2147483648]
                                                        _ = v52
                                                    }
                                                    if true {
                                                        v2 -= -702
                                                    }
                                                    v53 := v5[(^(fi26))]
                                                    _ = fi51
                                                    _ = v53
                                                }
                                            case 42:
                                                v5[(-(fi26))] = fi26
                                            }
                                        case "":
                                            var v54 int32 = (v27 | v27)
                                            _ = v54
                                        default:
                                            v28 = (v28 & v28)
                                        }
                                        _ = v41
                                    } else {
                                        for fi55 := 0; fi55 < 9; fi55++ {
                                            if (42 != v27) {
                                                if (((v2 % ((v2) | 1)) < -956) && false) {
                                                    v4 += v4
                                                } else {
                                                    var v56 uint64 = (^(v4))
                                                    _ = v56
                                                }
                                                fmt.Println(fi26, v4)
                                            } else {
                                                if false {
                                                    if ((v4 == (v4 % (((v4 ^ (v4 % ((v4) | 1)))) | 1))) && (!true)) {
                                                        if (1 == 789) {
                                                            v57 := GenS1{F0: 2}
                                                            var v58 int = 32768
                                                            switch fi55 {
                                                            case 9223372036854775807:
                                                                if (true || (42 >= 42)) {
                                                                    delete(v3, 2147483648)
                                                                    if (false && false) {
                                                                        v2 = v2
                                                                        v59 := GenS1{F0: 606}
                                                                        var v60 uint16 = 1
                                                                        _ = v59
                                                                        _ = v60
                                                                    }
                                                                }
                                                            default:
                                                                var v61 int = 4294967295
                                                                _ = v61
                                                            }
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        v62 := GenS6{F0: 0}
                                                        v63 := GenS6{F0: 1}
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    v2 += (v2 ^ v2)
                                                }
                                                switch "hello" {
                                                case "a":
                                                    switch "" {
                                                    case "":
                                                        switch (fi55 / (((fi55 + (-(fi55)))) | 1)) {
                                                        case -2147483648:
                                                            v28 += (v28 | 838)
                                                        default:
                                                            if (2147483647 != int32(v28)) {
                                                                var v64 int8 = 89
                                                                _ = v64
                                                            } else {
                                                                var v65 int8 = 1
                                                                var v66 float32 = float32((float64((3.14159 / 1.0)) * ((0.1 * (3.14159 / float64(1.5))) / 1e10)))
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                        }
                                                    case "\n\t":
                                                        if false {
                                                            if false {
                                                                switch ("" + "hello") {
                                                                case "a":
                                                                    var v67 int32 = -1
                                                                    _ = v67
                                                                default:
                                                                    v68 := GenS7{F0: -1, F1: 255, F2: v27}
                                                                    _ = v68
                                                                }
                                                                var v69 float64 = 3.14159
                                                                _ = v69
                                                            }
                                                            var v70 bool = true
                                                            fmt.Println(v2)
                                                            _ = v70
                                                        } else {
                                                            switch fi55 {
                                                            case 346:
                                                                v28 -= 18446744073709551615
                                                            case 2147483647:
                                                                v4 -= 18446744073709551615
                                                            default:
                                                                var v71 int16 = int16(fi26)
                                                                _ = v71
                                                            }
                                                            v72 := GenS5{F0: ((v28 + v28) < (^(v28))), F1: 1}
                                                            var v73 int8 = 100
                                                            switch "a" {
                                                            case "hello":
                                                                v74 := GenS8{F0: 42}
                                                                _ = v74
                                                            case "a":
                                                                v27 -= -5
                                                            default:
                                                                fmt.Println(fi26)
                                                            }
                                                            switch "" {
                                                            case "\n\t":
                                                                v4 += 0
                                                            case "foo bar":
                                                                fmt.Println(v28)
                                                            case "":
                                                                var v75 uint8 = 1
                                                                _ = v75
                                                            default:
                                                                var v76 bool = false
                                                                _ = v76
                                                            }
                                                            fmt.Println(fi55, v27, v73)
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    }
                                                case "foo bar":
                                                    if ((v27 - (v27 ^ -1)) <= 0) {
                                                        v77 := GenS9{F0: (fi55 / (((-(fi55))) | 1))}
                                                        fmt.Println(v2, v2, v4)
                                                        if true {
                                                            var v78 int32 = (v27 % ((-2147483648) | 1))
                                                            if (1.5 < (3.14159 / 3.14159)) {
                                                                var v79 int64 = (^(v2))
                                                                _ = v79
                                                            }
                                                            _ = v78
                                                        }
                                                        _ = v77
                                                    }
                                                default:
                                                    var v80 float32 = 0.1
                                                    _ = v80
                                                }
                                            }
                                            v28 = (v28 % ((v28) | 1))
                                            v81 := GenS10{F0: 0, F1: 65535}
                                            _ = fi55
                                            _ = v81
                                        }
                                        var v82 int = -1
                                        _ = v82
                                    }
                                }
                            }
                        }
                    } else {
                        v2 -= (v2 | v2)
                        v3[151] = fi26
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi26
            }
            for fi83 := 0; fi83 < 12; fi83++ {
                for fi84 := 0; fi84 < 1; fi84++ {
                    fmt.Println(v4, v2)
                    if (4294967295 == 965) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v4 = v4
                        if true {
                            var v85 int16 = 946
                            if (-32768 > v85) {
                                if true {
                                    if true {
                                        v4 -= v4
                                        fmt.Println(v2, fi84, fi84)
                                        switch "" {
                                        case "a":
                                            var v86 string = "hello"
                                            _ = v86
                                        case "\n\t":
                                            fmt.Println(fi84, fi84, fi83)
                                        default:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        v3[455] = fi84
                                        fmt.Println(fi84, fi83)
                                    }
                                }
                                v87 := GenS11{F0: 2}
                                _ = len(v1)
                                _ = v87
                            } else {
                                fmt.Println(fi83, v2, v4)
                            }
                            if ("" < "foo bar") {
                                v88 := GenS3{F0: ("hello" + "\n\t")}
                                _ = v88
                            } else {
                                v85 -= v85
                                v4 -= (v4 + v4)
                            }
                            v89 := GenS10{F0: 76, F1: 2}
                            v2 -= (v2 / ((9223372036854775807) | 1))
                            fmt.Println(fi83, v4)
                            _ = v85
                            _ = v89
                        } else {
                            v4 = 42
                            switch (fi84 * 1) {
                            case 32767:
                                fmt.Println(fi83)
                                v90 := GenS7{F0: 1000, F1: 27, F2: 2147483647}
                                var v91 string = "a"
                                _ = v90
                                _ = v91
                            }
                        }
                    }
                    if true {
                        switch fi84 {
                        case 1:
                            var v92 int = (^(fi84))
                            _ = v92
                        case 0:
                            var v93 int = 32767
                            _ = v93
                        }
                        var v94 bool = true
                        _ = v94
                    }
                    if true {
                        var v95 uint32 = 513
                        v96 := GenS4{F0: v95, F1: 2}
                        if (127 == 2147483647) {
                            var v97 int16 = -1
                            _ = v97
                        } else {
                            fmt.Println(v2, fi83, fi83)
                            _ = len(v1)
                            if (255 <= 1) {
                                v98 := GenS5{F0: true, F1: 0}
                                v95 = 2
                                if (2 <= 567) {
                                    if (2 != 255) {
                                        if ((fi83 + 127) == fi83) {
                                            switch int(v4) {
                                            case 42:
                                                fmt.Println(v95, fi83, fi84)
                                            default:
                                                v3[(^(v95))] = (fi83 - fi83)
                                            }
                                            switch "" {
                                            case "\n\t":
                                                var v99 int8 = -128
                                                _ = v99
                                            case "a":
                                                v100 := GenS12{F0: v4}
                                                _ = v100
                                            }
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                v2 += (v2 - -1)
                                            case "a":
                                                v101 := GenS10{F0: 100, F1: (v95 + 65535)}
                                                _ = v101
                                            case "foo bar":
                                                fmt.Println(fi83)
                                            default:
                                                v2 = 42
                                                switch "a" {
                                                case "foo bar":
                                                    var v102 int64 = (v2 | (^(v2)))
                                                    _ = v102
                                                case "\n\t":
                                                    v4 = 531
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    v103 := GenS13{F0: 2147483647, F1: (((668 < v95) || false) || (-32768 <= 42))}
                                                    _ = v103
                                                }
                                            }
                                            switch 4294967295 {
                                            case 590:
                                                fmt.Println(fi83, v4)
                                            case -9223372036854775808:
                                                v2 += v2
                                            case -405:
                                                fmt.Println(v95, v2)
                                            }
                                        }
                                        if (-213 > 0) {
                                            var v104 float64 = 0.1
                                            var v105 uint32 = 65535
                                            _ = v104
                                            _ = v105
                                        } else {
                                            v95 += (^(v95))
                                            if false {
                                                var v106 int8 = 59
                                                if (!true) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v107 uint16 = 255
                                                    _ = v107
                                                }
                                                v1 = append(v1, (fi84 & fi84))
                                                _ = v106
                                            }
                                        }
                                    }
                                    v108 := GenS3{F0: ""}
                                    v109 := GenS12{F0: v4}
                                    _ = v108
                                    _ = v109
                                }
                                v95 += 4294967295
                                _ = v98
                            } else {
                                fmt.Println(v95, v95)
                                v2 += v2
                            }
                            switch ("\n\t" + "a") {
                            case "hello":
                                if true {
                                    if (true && true) {
                                        v110 := GenS2{F0: (fi83 % (((^(fi83))) | 1)), F1: -12, F2: 42}
                                        switch (-(fi83)) {
                                        case 128:
                                            v5[(^(fi83))] = fi83
                                            v3[65535] = int(v95)
                                        default:
                                            if true {
                                                var v111 float64 = 1e10
                                                _ = v111
                                            }
                                        }
                                        _ = v110
                                    }
                                }
                            case "a":
                                v95 = (v95 + (v95 | 2147483648))
                            }
                        }
                        _ = v95
                        _ = v96
                    } else {
                        fmt.Println(v2)
                        var v112 uint32 = 42
                        _ = v112
                    }
                    switch fi83 {
                    case 127:
                        v4 -= v4
                        v4 = (v4 + (v4 ^ 895))
                        v4 = v4
                        v4 = (v4 + v4)
                    default:
                        v4 += (v4 & v4)
                    }
                    _ = fi84
                }
                var v113 int = 453
                _ = fi83
                _ = v113
            }
        }
    } else {
        delete(v5, 0)
        v4 = 1
    }
    var v114 int16 = 42
    v115 := GenS2{F0: -1, F1: 100, F2: -327}
    v116 := v3[65535]
    if false {
        for fi117 := 0; fi117 < 17; fi117++ {
            if (!true) {
                v114 = v114
                v114 = 32767
                var v118 uint32 = 464
                _ = v118
            } else {
                switch ("hello" + "") {
                case "hello":
                    v119 := GenS14{F0: true, F1: 1}
                    _ = v119
                case "\n\t":
                    if false {
                        if ((180 < (v4 / ((1) | 1))) && (1 > -128)) {
                            v120 := GenS8{F0: (v114 ^ v114)}
                            v121 := GenS15{F0: 255, F1: -1, F2: (v4 << ((0) & 63))}
                            _ = v120
                            _ = v121
                        }
                    }
                case "":
                    v114 = v114
                default:
                    switch v116 {
                    case 2147483647:
                        switch "a" {
                        case "a":
                            var v122 int8 = 119
                            _ = v122
                        case "foo bar":
                            var v123 int32 = 2147483647
                            _ = v123
                        default:
                            v2 += v2
                        }
                    default:
                        for fi124 := 0; fi124 < 15; fi124++ {
                            v4 += v4
                            if false {
                                v125 := GenS8{F0: (v114 | v114)}
                                if ((v114 <= (v114 | v114)) && false) {
                                    v126 := GenS16{F0: -1}
                                    v4 -= v4
                                    var v127 uint = 2
                                    _ = v126
                                    _ = v127
                                } else {
                                    v116 = 1
                                    fmt.Println(fi124, v116, fi117)
                                }
                                _ = v125
                            } else {
                                fmt.Println(v2, v4)
                                if false {
                                    switch ((((("\n\t" + "hello") + ("foo bar" + "a")) + ("" + "")) + "foo bar") + ((("\n\t" + "") + (("foo bar" + ("hello" + (("" + ((("\n\t" + ("\n\t" + "\n\t")) + "foo bar") + "foo bar")) + "foo bar"))) + "")) + ("" + "foo bar"))) {
                                    case "foo bar":
                                        v128 := GenS17{F0: "a"}
                                        _ = v128
                                    }
                                }
                            }
                            if false {
                                v129 := GenS10{F0: 0, F1: 0}
                                _ = v129
                            }
                            _ = len(v1)
                            _ = fi124
                        }
                    }
                }
            }
            fmt.Println(fi117, v2)
            _ = fi117
        }
        v130 := GenS18{F0: uint(v4)}
        _ = v130
    } else {
        v131 := GenS4{F0: 4294967295, F1: 335}
        var v132 int = v116
        v116 = (v116 >> ((324) & 63))
        _ = v131
        _ = v132
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v114:", v114)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v114
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
