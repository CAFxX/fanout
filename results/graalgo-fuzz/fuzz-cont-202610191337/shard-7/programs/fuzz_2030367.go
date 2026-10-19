package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string; F1 uint; F2 int8 }
type GenS9 struct { F0 uint8; F1 string; F2 int64 }
type GenS10 struct { F0 int; F1 uint8; F2 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int8 }

func main() {
    v1 := []int{127, -9223372036854775808, -52}
    var v2 int64 = 0
    var v3 uint16 = 1
    var v4 uint8 = 140
    var v5 uint64 = 42
    var v6 int = -122
    for fi7 := 0; fi7 < 5; fi7++ {
        v5 -= 4294967295
        if (!false) {
            var v8 float64 = (float64((1e10 + float32(0.1))) + (1.5 * 1e10))
            _ = v8
        }
        _ = fi7
    }
    v1 = append(v1, v6)
    var v9 int16 = 42
    switch ("" + ("\n\t" + "foo bar")) {
    case "a":
        switch int(v2) {
        case 35:
            if (float64(0.0) >= (1e10 / (float64(float32(1e10)) - 0.1))) {
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := GenS1{F0: "", F1: -874, F2: v4}
                    v5 = (v5 + v5)
                    _ = fi10
                    _ = v11
                }
            } else {
                v2 = v2
                var v12 int8 = -96
                _ = v12
            }
        case 4294967295:
            v13 := GenS2{F0: -128}
            _ = v13
        default:
            v14 := GenS2{F0: 49}
            _ = v14
        }
    case "\n\t":
        for fi15 := 0; fi15 < 7; fi15++ {
            if false {
                if true {
                    var v16 uint = 4294967295
                    for fi17 := 0; fi17 < 4; fi17++ {
                        var v18 float32 = ((float32((0.0 - 1e10)) * (0.1 * (float32(1e10) / float32(-2.25)))) - 0.0)
                        v5 = (^(v5))
                        v16 = v16
                        var v19 uint32 = 25
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = v16
                } else {
                    v4 = 0
                    if ((154 >= (v4 + v4)) || (!(true || (true || (162 != (v5 % (((v5 << ((0) & 63))) | 1))))))) {
                        var v20 uint32 = 65535
                        switch ("foo bar" + "") {
                        case "hello":
                            var v21 uint32 = (v20 / (((v20 * (v20 >> ((uint(v5)) & 63)))) | 1))
                            _ = v21
                        case "":
                            for fi22 := 0; fi22 < 2; fi22++ {
                                if (false || false) {
                                    fmt.Println(v3)
                                } else {
                                    var v23 float32 = -2.25
                                    v24 := GenS1{F0: "a", F1: -1, F2: (v4 / ((v4) | 1))}
                                    fmt.Println(fi22, v3, v6)
                                    v25 := GenS2{F0: 1}
                                    if (-90 == 100) {
                                        v20 += v20
                                        if true {
                                            if (uint64(v20) != (v5 + v5)) {
                                                fmt.Println(v5)
                                            }
                                            v20 -= (v20 / ((1) | 1))
                                        }
                                        v26 := GenS3{F0: 42}
                                        _ = len(v1)
                                        v27 := GenS3{F0: -779}
                                        _ = v26
                                        _ = v27
                                    } else {
                                        if (v3 < (v3 | 42)) {
                                            v28 := GenS4{F0: v20, F1: (v20 >> ((65535) & 63)), F2: v3}
                                            var v29 uint16 = 2
                                            if len(v1) > 0 {
                                                v1[0] = fi15
                                            }
                                            _ = v28
                                            _ = v29
                                        } else {
                                            fmt.Println(v3)
                                            fmt.Println(v20)
                                            if false {
                                                if len(v1) > 0 {
                                                    v1[0] = v6
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 127
                                                }
                                                v6 = (fi15 | (fi15 % ((4294967295) | 1)))
                                                v30 := GenS5{F0: (v3 * v3), F1: "a"}
                                                _ = v30
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                switch ("\n\t" + ("foo bar" + "foo bar")) {
                                                case "":
                                                    var v31 uint64 = 18446744073709551615
                                                    _ = v31
                                                case "a":
                                                    var v32 uint = 783
                                                    _ = v32
                                                }
                                            }
                                            v9 += (v9 >> ((0) & 63))
                                            fmt.Println(v4)
                                        }
                                        if false {
                                            fmt.Println(fi15, v3, v5)
                                            v9 += int16(v2)
                                            if len(v1) > 0 {
                                                v1[0] = (fi22 - fi15)
                                            }
                                        }
                                        var v33 int = fi22
                                        _ = v33
                                    }
                                    var v34 float64 = (1e10 + float64(v23))
                                    v1 = append(v1, (-(v6)))
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                    _ = v34
                                }
                                fmt.Println(v5, v4)
                                v20 += v20
                                _ = fi22
                            }
                        case "a":
                            var v35 float32 = 3.14159
                            _ = v35
                        }
                        fmt.Println(v5, v2)
                        _ = v20
                    }
                    var v36 int16 = v9
                    _ = v36
                }
                v37 := GenS6{F0: v4, F1: "foo bar"}
                var v38 float32 = 0.0
                _ = v37
                _ = v38
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi15
        }
    case "hello":
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 string = "\n\t"
            v4 = v4
            _ = fi39
            _ = v40
        }
    }
    if true {
        v3 = v3
        _ = len(v1)
        switch "hello" {
        case "a":
            v41 := GenS6{F0: v4, F1: ""}
            _ = v41
        case "":
            var v42 float64 = 0.0
            _ = v42
        case "foo bar":
            if false {
                for fi43 := 0; fi43 < 13; fi43++ {
                    v44 := GenS3{F0: -925}
                    v45 := GenS4{F0: 65535, F1: 1, F2: 65535}
                    if ((!(uint8(v5) == 42)) && false) {
                        for fi46 := 0; fi46 < 13; fi46++ {
                            var v47 int = fi43
                            _ = fi46
                            _ = v47
                        }
                        v5 += v5
                        v48 := GenS1{F0: "\n\t", F1: v6, F2: v4}
                        _ = v48
                    }
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            }
        default:
            var v49 bool = ((!(-32768 == v9)) && true)
            _ = v49
        }
    } else {
        fmt.Println(v5)
        var v50 float64 = 0.1
        for fi51 := 0; fi51 < 7; fi51++ {
            v5 = uint64(v4)
            _ = fi51
        }
        _ = v50
    }
    switch "\n\t" {
    case "a":
        if (!false) {
            v52 := GenS5{F0: v3, F1: "a"}
            switch int(v4) {
            case -9223372036854775808:
                var v53 bool = (v5 != uint64(v3))
                _ = v53
            case 427:
                var v54 uint8 = uint8(v5)
                _ = v54
            case 1:
                if ((v9 % ((440) | 1)) > (v9 & v9)) {
                    var v55 uint64 = uint64(v3)
                    fmt.Println(v9, v5)
                    switch -155 {
                    case 128:
                        for fi56 := 0; fi56 < 3; fi56++ {
                            v57 := GenS6{F0: v4, F1: (("hello" + "") + "\n\t")}
                            v2 = (-(v2))
                            _ = fi56
                            _ = v57
                        }
                    }
                    _ = v55
                }
            default:
                v5 += (v5 * (v5 - 2))
            }
            var v58 bool = true
            _ = v52
            _ = v58
        } else {
            if len(v1) > 0 {
                v1[0] = 722
            }
            for fi59 := 0; fi59 < 7; fi59++ {
                _ = len(v1)
                if false {
                    _ = len(v1)
                    _ = len(v1)
                    v60 := GenS7{F0: v5}
                    _ = v60
                }
                _ = fi59
            }
            for fi61 := 0; fi61 < 2; fi61++ {
                v62 := GenS8{F0: "foo bar", F1: 485, F2: 100}
                switch (fi61 / ((v6) | 1)) {
                case -1:
                    fmt.Println(v9, v4, v3)
                case 9223372036854775807:
                    for fi63 := 0; fi63 < 9; fi63++ {
                        v1 = append(v1, (v6 | (v6 - (-(fi63)))))
                        if false {
                            var v64 int16 = v9
                            var v65 uint16 = 750
                            _ = v64
                            _ = v65
                        }
                        _ = fi63
                    }
                case -115:
                    var v66 uint = 670
                    _ = v66
                default:
                    v5 += (v5 >> ((18446744073709551615) & 63))
                }
                v67 := GenS5{F0: v3, F1: ("hello" + "hello")}
                _ = fi61
                _ = v62
                _ = v67
            }
        }
    default:
        if (4294967295 < 1) {
            var v68 int64 = -318
            var v69 int64 = (v68 - (v68 | (v2 << ((2) & 63))))
            for fi70 := 0; fi70 < 9; fi70++ {
                if (v4 != 178) {
                    v9 += (v9 & (v9 * (v9 % ((int16(v2)) | 1))))
                    v69 += int64(v5)
                }
                if false {
                    if (42 > 42) {
                        fmt.Println(v9, v3)
                        if true {
                            v71 := GenS3{F0: 186}
                            _ = v71
                        } else {
                            if false {
                                var v72 bool = true
                                var v73 int16 = v9
                                _ = v72
                                _ = v73
                            }
                        }
                        switch "hello" {
                        case "\n\t":
                            for fi74 := 0; fi74 < 6; fi74++ {
                                if true {
                                    _ = len(v1)
                                } else {
                                    if ((v2 & (v2 | 85)) <= v69) {
                                        switch int(v3) {
                                        case -9223372036854775808:
                                            switch (fi70 & fi74) {
                                            case 32767:
                                                switch fi70 {
                                                case 1:
                                                    if ((!(v3 < (^(v3)))) && true) {
                                                        if false {
                                                            if (-2.25 < 0.0) {
                                                                v75 := GenS1{F0: ("foo bar" + ("" + "\n\t")), F1: int(v5), F2: 206}
                                                                var v76 bool = (v9 != (v9 + v9))
                                                                if v76 {
                                                                    v1 = append(v1, (^(fi70)))
                                                                } else {
                                                                    v76 = ("" == (((("a" + "foo bar") + ("" + "foo bar")) + "hello") + ("a" + "hello")))
                                                                }
                                                                if false {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        switch v6 {
                                                                        case -544:
                                                                            var v77 uint = 2
                                                                            _ = v77
                                                                        default:
                                                                            var v78 uint8 = (v4 / ((v4) | 1))
                                                                            _ = v78
                                                                        }
                                                                    case "\n\t":
                                                                        var v79 bool = (!v76)
                                                                        _ = v79
                                                                    }
                                                                }
                                                                v9 -= (v9 - v9)
                                                                v80 := GenS3{F0: 1}
                                                                _ = v75
                                                                _ = v76
                                                                _ = v80
                                                            }
                                                        } else {
                                                            v81 := GenS7{F0: (v5 * 796)}
                                                            fmt.Println(v6, v3, v3)
                                                            _ = v81
                                                        }
                                                    }
                                                case 9223372036854775807:
                                                    var v82 int64 = 9223372036854775807
                                                    _ = v82
                                                case 2147483647:
                                                    var v83 uint64 = v5
                                                    _ = v83
                                                default:
                                                    v6 = fi74
                                                }
                                            case 1:
                                                v68 = v69
                                                v4 += (^(v4))
                                                v5 = 18446744073709551615
                                                v9 -= (v9 << ((90) & 63))
                                                fmt.Println(v69, v9, v3)
                                            case 42:
                                                var v84 uint32 = uint32(v3)
                                                v68 = v69
                                                v3 = 42
                                                _ = v84
                                            default:
                                                v4 = v4
                                            }
                                        default:
                                            if true {
                                                switch -1 {
                                                case -1:
                                                    if (-1 <= 42) {
                                                        var v85 uint16 = v3
                                                        _ = v85
                                                    }
                                                case 32767:
                                                    var v86 int = fi70
                                                    _ = v86
                                                case -617:
                                                    var v87 int16 = (-(v9))
                                                    if (!false) {
                                                        var v88 uint = 919
                                                        switch (-(v6)) {
                                                        case 327:
                                                            var v89 uint8 = v4
                                                            _ = v89
                                                        case 127:
                                                            v69 = (v2 >> ((uint(v4)) & 63))
                                                        case -724:
                                                            var v90 uint64 = 0
                                                            _ = v90
                                                        }
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            _ = len(v1)
                                                        default:
                                                            v4 += 1
                                                        }
                                                        var v91 uint8 = 94
                                                        _ = v88
                                                        _ = v91
                                                    } else {
                                                        v2 += -836
                                                        fmt.Println(v87, v87, v68)
                                                        if len(v1) > 0 {
                                                            v1[0] = 4294967295
                                                        }
                                                    }
                                                    _ = v87
                                                }
                                            }
                                        }
                                        var v92 uint64 = uint64(v4)
                                        var v93 int64 = (v69 >> ((482) & 63))
                                        _ = v92
                                        _ = v93
                                    } else {
                                        v94 := GenS8{F0: ("a" + "\n\t"), F1: 4294967295, F2: 108}
                                        _ = v94
                                    }
                                    switch (("foo bar" + "") + "foo bar") {
                                    case "\n\t":
                                        switch "a" {
                                        case "":
                                            v95 := GenS8{F0: "a", F1: 11, F2: 1}
                                            _ = v95
                                        case "hello":
                                            if (false && false) {
                                                v68 = (v68 * (v68 & int64(v6)))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (!(false || ((v9 + 42) < v9))) {
                                                    v68 = (v69 - v69)
                                                    v1 = append(v1, 32768)
                                                    switch (^(fi74)) {
                                                    case 4294967295:
                                                        if true {
                                                            if false {
                                                                if true {
                                                                    v68 = int64(v3)
                                                                } else {
                                                                    v96 := GenS9{F0: (v4 * 0), F1: "\n\t", F2: v68}
                                                                    v69 = (v2 ^ v69)
                                                                    var v97 bool = ((v5 < v5) && true)
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (v6 + (fi70 + (v6 & fi74)))
                                                                    }
                                                                    v98 := GenS9{F0: v4, F1: ("foo bar" + (("foo bar" + "foo bar") + "foo bar")), F2: 348}
                                                                    _ = v96
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                                var v99 int32 = -2147483648
                                                                _ = v99
                                                            } else {
                                                                v100 := GenS3{F0: 42}
                                                                fmt.Println(v68)
                                                                switch fi74 {
                                                                case 268:
                                                                    var v101 string = "a"
                                                                    _ = v101
                                                                case 1:
                                                                    var v102 float64 = 0.0
                                                                    _ = v102
                                                                default:
                                                                    v6 += (fi74 << ((4294967295) & 63))
                                                                }
                                                                _ = v100
                                                            }
                                                            v2 = v68
                                                        } else {
                                                            switch (fi70 << ((1) & 63)) {
                                                            case -1:
                                                                fmt.Println(v69, v2)
                                                            case -797:
                                                                if len(v1) > 0 {
                                                                    v1[0] = v6
                                                                }
                                                            }
                                                            if false {
                                                                var v103 bool = ("hello" != "\n\t")
                                                                if v103 {
                                                                    switch (-(v6)) {
                                                                    case 32767:
                                                                        var v104 uint32 = 939
                                                                        _ = v104
                                                                    case -223:
                                                                        fmt.Println(v2)
                                                                    case -339:
                                                                        fmt.Println(v4)
                                                                    }
                                                                    var v105 uint64 = 0
                                                                    var v106 uint8 = v4
                                                                    _ = v105
                                                                    _ = v106
                                                                }
                                                                v5 += 4294967295
                                                                v3 -= (v3 % ((v3) | 1))
                                                                _ = v103
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v5 += (v5 | (v5 << ((18446744073709551615) & 63)))
                                                }
                                            } else {
                                                fmt.Println(v3, v69)
                                            }
                                            v3 = 255
                                        case "foo bar":
                                            if false {
                                                v107 := GenS7{F0: v5}
                                                fmt.Println(fi70)
                                                v108 := GenS1{F0: "", F1: 32767, F2: (v4 & (^(v4)))}
                                                _ = v107
                                                _ = v108
                                            }
                                        }
                                        var v109 float64 = 0.0
                                        var v110 int = (v6 + (fi74 & int(v4)))
                                        _ = v109
                                        _ = v110
                                    }
                                }
                                v4 = (v4 | (v4 & (v4 + v4)))
                                if false {
                                    var v111 int8 = -128
                                    _ = v111
                                } else {
                                    fmt.Println(v9)
                                    var v112 float32 = float32(1e10)
                                    _ = v112
                                }
                                _ = fi74
                            }
                        case "hello":
                            var v113 int64 = 9223372036854775807
                            _ = v113
                        default:
                            switch (v6 << ((uint(v5)) & 63)) {
                            case 128:
                                v68 -= (v68 * v69)
                            default:
                                for fi114 := 0; fi114 < 10; fi114++ {
                                    fmt.Println(v9, v6, v4)
                                    fmt.Println(v2)
                                    v115 := GenS8{F0: ("foo bar" + "\n\t"), F1: 42, F2: 5}
                                    _ = fi114
                                    _ = v115
                                }
                            }
                        }
                    } else {
                        v69 -= v2
                        v1 = append(v1, (fi70 << ((0) & 63)))
                        v68 += (v2 - v2)
                    }
                } else {
                    v3 = (v3 << ((255) & 63))
                }
                switch "hello" {
                case "a":
                    v116 := GenS3{F0: 2147483647}
                    _ = v116
                }
                _ = fi70
            }
            _ = v68
            _ = v69
        }
    }
    v1 = append(v1, 4294967295)
    fmt.Println(v3, v9)
    for fi117 := 0; fi117 < 10; fi117++ {
        fmt.Println(v3, v9)
        var v118 uint8 = uint8(v5)
        v119 := GenS7{F0: v5}
        _ = fi117
        _ = v118
        _ = v119
    }
    var v120 uint = 18446744073709551615
    for fi121 := 0; fi121 < 9; fi121++ {
        if len(v1) > 0 {
            v1[0] = (v6 - v6)
        }
        if len(v1) > 0 {
            v1[0] = v6
        }
        v122 := GenS10{F0: (-(v6)), F1: 1, F2: 100}
        _ = fi121
        _ = v122
    }
    v123 := GenS11{F0: (^(v4))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi124 := 0; fi124 < 1; fi124++ {
        var v125 float32 = (((float32(float64(float32(0.0))) * -2.25) - 1e10) * 0.0)
        _ = fi124
        _ = v125
    }
    if len(v1) > 0 {
        v1[0] = -49
    }
    switch 0 {
    case 42:
        if (true || (((v3 << ((455) & 63)) == v3) && false)) {
            if len(v1) > 0 {
                v1[0] = (v6 * (v6 & v6))
            }
            var v126 int16 = (v9 / (((v9 | -480)) | 1))
            v126 = v9
            _ = v126
        } else {
            var v127 string = "foo bar"
            _ = v127
        }
    case 32768:
        if ((230 != 935) && (!(true && (1.5 <= 1.5)))) {
            v6 = v6
        }
    case 127:
        v128 := GenS12{F0: 100}
        _ = v128
    }
    if len(v1) > 0 {
        v1[0] = (v6 / ((v6) | 1))
    }
    for fi129 := 0; fi129 < 16; fi129++ {
        v5 = (v5 % (((v5 / ((0) | 1))) | 1))
        v130 := GenS1{F0: (("" + (("" + ("a" + "hello")) + (("\n\t" + "\n\t") + "foo bar"))) + "foo bar"), F1: fi129, F2: 1}
        _ = fi129
        _ = v130
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v120
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
