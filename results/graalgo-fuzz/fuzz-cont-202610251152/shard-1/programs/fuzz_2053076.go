package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 bool; F2 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 bool; F1 uint16; F2 int32 }
type GenS13 struct { F0 int8; F1 string; F2 int32 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint; F1 uint8; F2 int8 }
type GenS18 struct { F0 int32; F1 int64 }

func main() {
    v1 := []int{}
    var v2 uint = 889
    var v3 uint = 65535
    for fi4 := 0; fi4 < 6; fi4++ {
        v5 := GenS1{F0: 2147483647, F1: "hello", F2: -1}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            if false {
                v3 = v2
                switch "foo bar" {
                case "a":
                    var v8 int64 = -456
                    _ = v8
                default:
                    v3 = 18446744073709551615
                }
                var v9 string = ("a" + ("foo bar" + ("foo bar" + (("\n\t" + ("" + "\n\t")) + ""))))
                switch (^(fi6)) {
                case 32767:
                    var v10 uint8 = 127
                    _ = v10
                case -2147483648:
                    var v11 uint16 = 42
                    _ = v11
                default:
                    var v12 uint = v3
                    v13 := GenS2{F0: 0}
                    v14 := GenS3{F0: 1, F1: 9223372036854775807, F2: (^(v12))}
                    switch -1 {
                    case -9223372036854775808:
                        var v15 uint64 = 168
                        _ = v15
                    default:
                        if (false || true) {
                            v3 += 18446744073709551615
                            v1 = append(v1, fi7)
                            fmt.Println(fi7, fi7)
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                            fmt.Println(fi6)
                        }
                    }
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v16 string = "hello"
                _ = v16
                _ = v9
            } else {
                var v17 int = fi7
                _ = v17
            }
            var v18 float32 = 0.1
            v3 = (v2 & (v2 / ((v2) | 1)))
            _ = fi7
            _ = v18
        }
        _ = fi6
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if (((((3.14159 * 3.14159) * ((0.0 - (0.1 - 1.5)) - float32(1e10))) / float32(-2.25)) + float32(float64(float32(-2.25)))) >= float32(-2.25)) {
            v20 := GenS4{F0: (fi19 % (((fi19 % ((fi19) | 1))) | 1))}
            _ = len(v1)
            _ = v20
        } else {
            v21 := GenS4{F0: (^(fi19))}
            _ = v21
        }
        v3 -= v3
        _ = fi19
    }
    v3 = (v3 | v3)
    v2 = v3
    if (282 <= -1) {
        var v22 float64 = 0.0
        _ = v22
    }
    v2 += 2
    if true {
        for fi23 := 0; fi23 < 6; fi23++ {
            if (32767 == -555) {
                if ((v3 ^ v2) > (v3 >> ((v2) & 63))) {
                    var v24 bool = ((3.14159 > 0.0) || (0 > 42))
                    _ = v24
                } else {
                    v3 = (v3 << (((v2 % ((2) | 1))) & 63))
                    for fi25 := 0; fi25 < 4; fi25++ {
                        if (1 >= 65) {
                            if true {
                                switch (^(fi25)) {
                                case 127:
                                    v2 = 797
                                case 128:
                                    if len(v1) > 0 {
                                        v1[0] = (fi23 << ((255) & 63))
                                    }
                                case 9223372036854775807:
                                    v26 := GenS4{F0: fi25}
                                    _ = v26
                                default:
                                    fmt.Println(fi25)
                                }
                                if false {
                                    var v27 float32 = 0.1
                                    v27 += (3.14159 / 3.14159)
                                    switch fi23 {
                                    case -769:
                                        var v28 int8 = 1
                                        _ = v28
                                    }
                                    fmt.Println(v2, v27)
                                    _ = len(v1)
                                    if (!false) {
                                        v29 := GenS5{F0: (("a" + "foo bar") + "hello")}
                                        switch 638 {
                                        case -2147483648:
                                            if (("foo bar" + "") >= "\n\t") {
                                                var v30 int = fi23
                                                v31 := GenS2{F0: 286}
                                                v27 += 0.1
                                                _ = v30
                                                _ = v31
                                            }
                                        case 127:
                                            var v32 int16 = -143
                                            _ = v32
                                        case -472:
                                            switch ("hello" + "") {
                                            case "":
                                                if (17 > 1) {
                                                    if ((v27 + (3.14159 - (v27 + v27))) != (v27 * v27)) {
                                                        v33 := GenS2{F0: v2}
                                                        if ((^(v3)) >= 18446744073709551615) {
                                                            _ = len(v1)
                                                        } else {
                                                            v34 := GenS6{F0: v3, F1: false, F2: 69}
                                                            v1 = append(v1, 127)
                                                            fmt.Println(v27, v3, v27)
                                                            _ = v34
                                                        }
                                                        v3 += 42
                                                        var v35 uint32 = 0
                                                        _ = v33
                                                        _ = v35
                                                    }
                                                    var v36 float64 = (0.1 * 0.0)
                                                    var v37 bool = true
                                                    v27 = (float32(v36) / 1e10)
                                                    _ = v36
                                                    _ = v37
                                                }
                                            case "foo bar":
                                                if (true || ((0.1 + v27) >= (v27 - v27))) {
                                                    v3 -= 255
                                                    switch fi25 {
                                                    case 32:
                                                        var v38 uint16 = 42
                                                        _ = v38
                                                    }
                                                    var v39 int16 = -1
                                                    var v40 string = "foo bar"
                                                    if len(v1) > 0 {
                                                        v1[0] = fi23
                                                    }
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    v3 = 1
                                                    v3 = (v2 + 65535)
                                                }
                                            default:
                                                v27 = 3.14159
                                            }
                                        }
                                        var v41 int = fi25
                                        _ = v29
                                        _ = v41
                                    } else {
                                        var v42 string = "foo bar"
                                        _ = v42
                                    }
                                    if true {
                                        v43 := GenS2{F0: v3}
                                        v1 = append(v1, -23)
                                        _ = v43
                                    }
                                    v3 += v2
                                    _ = v27
                                }
                                v44 := GenS1{F0: 0, F1: "\n\t", F2: int16(fi23)}
                                if false {
                                    var v45 int16 = 42
                                    _ = v45
                                }
                                var v46 int8 = 0
                                _ = v44
                                _ = v46
                            } else {
                                _ = len(v1)
                                var v47 string = "\n\t"
                                var v48 bool = (true && true)
                                _ = v47
                                _ = v48
                            }
                        }
                        _ = fi25
                    }
                }
                for fi49 := 0; fi49 < 7; fi49++ {
                    var v50 uint64 = 0
                    v50 -= (v50 * (v50 + (v50 & (v50 ^ (v50 - (v50 >> ((v2) & 63)))))))
                    var v51 float64 = (1.5 * float64(3.14159))
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            }
            _ = fi23
        }
        v3 += (v3 & v2)
        v2 = 4294967295
    } else {
        v2 = (^(v3))
        switch 32768 {
        case -2147483648:
            for fi52 := 0; fi52 < 17; fi52++ {
                var v53 float64 = 1e10
                fmt.Println(v2, v3)
                for fi54 := 0; fi54 < 8; fi54++ {
                    var v55 float32 = 1.5
                    _ = fi54
                    _ = v55
                }
                _ = fi52
                _ = v53
            }
        default:
            for fi56 := 0; fi56 < 17; fi56++ {
                for fi57 := 0; fi57 < 19; fi57++ {
                    var v58 int = fi56
                    v3 = (v2 + v3)
                    _ = fi57
                    _ = v58
                }
                v2 = 65535
                v3 = 255
                _ = fi56
            }
        }
        v2 = (^(v2))
    }
    v59 := GenS1{F0: -1, F1: "\n\t", F2: -1}
    v1 = append(v1, 127)
    if false {
        v60 := GenS7{F0: true}
        _ = v60
    }
    var v61 float64 = 1e10
    if true {
        switch "a" {
        case "\n\t":
            v2 = v3
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        for fi62 := 0; fi62 < 2; fi62++ {
            switch (("hello" + "hello") + "\n\t") {
            case "hello":
                v61 = 3.14159
            case "\n\t":
                switch -175 {
                case 42:
                    for fi63 := 0; fi63 < 12; fi63++ {
                        fmt.Println(v2)
                        if ((-2147483648 < -2147483648) && false) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v1 = append(v1, fi63)
                        } else {
                            if ((v2 < 2) || true) {
                                fmt.Println(fi63)
                            } else {
                                v61 += (v61 - v61)
                                var v64 int8 = 91
                                _ = v64
                            }
                            if true {
                                switch (fi63 / (((-(fi63))) | 1)) {
                                case -1:
                                    if (42 > 214) {
                                        v3 -= v3
                                    }
                                case 32767:
                                    fmt.Println(v61)
                                case 128:
                                    v65 := GenS6{F0: (^(v3)), F1: (v2 >= 452), F2: -1}
                                    _ = v65
                                default:
                                    v2 -= (v3 / ((v3) | 1))
                                }
                            }
                        }
                        if (("hello" < "a") && (2 > 255)) {
                            v1 = append(v1, (fi62 | (fi62 ^ (fi62 * (fi63 | (fi63 & fi63))))))
                        } else {
                            v2 = v2
                            var v66 string = ("foo bar" + ("\n\t" + "hello"))
                            if true {
                                var v67 uint32 = 1
                                var v68 string = v66
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v67
                                _ = v68
                            }
                            _ = v66
                        }
                        _ = fi63
                    }
                case 128:
                    fmt.Println(v2, fi62, v61)
                case 32767:
                    for fi69 := 0; fi69 < 3; fi69++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v70 := GenS7{F0: (0 >= -32768)}
                        _ = fi69
                        _ = v70
                    }
                default:
                    for fi71 := 0; fi71 < 1; fi71++ {
                        fmt.Println(v61, v2)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi71
                    }
                }
            case "a":
                v72 := GenS8{F0: fi62}
                _ = v72
            default:
                var v73 uint64 = 1
                _ = v73
            }
            _ = fi62
        }
        if len(v1) > 0 {
            v1[0] = 722
        }
    }
    if false {
        fmt.Println(v3, v61)
        switch ("" + "") {
        case "":
            var v74 int8 = 1
            _ = v74
        case "hello":
            switch "a" {
            case "hello":
                var v75 int = -2147483648
                _ = v75
            case "\n\t":
                var v76 float32 = 1.5
                _ = v76
            case "a":
                for fi77 := 0; fi77 < 10; fi77++ {
                    switch ("a" + "hello") {
                    case "a":
                        var v78 float64 = v61
                        _ = v78
                    case "hello":
                        if true {
                            var v79 uint32 = 289
                            if true {
                                if (!(246 >= 2)) {
                                    fmt.Println(v3, v2)
                                    switch fi77 {
                                    case 151:
                                        fmt.Println(fi77, fi77)
                                    default:
                                        for fi80 := 0; fi80 < 1; fi80++ {
                                            v3 += 49
                                            var v81 uint = (v2 + (^(v3)))
                                            _ = fi80
                                            _ = v81
                                        }
                                    }
                                } else {
                                    if (1 > 1) {
                                        for fi82 := 0; fi82 < 13; fi82++ {
                                            var v83 int16 = 1000
                                            _ = fi82
                                            _ = v83
                                        }
                                    } else {
                                        if (v79 >= (^(v79))) {
                                            v3 -= (v2 << ((v3) & 63))
                                            switch (fi77 * 1) {
                                            case 42:
                                                if false {
                                                    var v84 float32 = -2.25
                                                    if false {
                                                        v85 := GenS5{F0: ("a" + "foo bar")}
                                                        _ = v85
                                                    } else {
                                                        switch fi77 {
                                                        case 788:
                                                            v61 = v61
                                                        }
                                                        for fi86 := 0; fi86 < 12; fi86++ {
                                                            if (!false) {
                                                                fmt.Println(fi77, v61, v3)
                                                                var v87 uint64 = 42
                                                                var v88 int16 = 42
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                            _ = fi86
                                                        }
                                                        if true {
                                                            var v89 int64 = 0
                                                            for fi90 := 0; fi90 < 6; fi90++ {
                                                                v89 += v89
                                                                if true {
                                                                    if true {
                                                                        if true {
                                                                            if (1 <= 32767) {
                                                                                var v91 float64 = ((v61 + 1e10) / v61)
                                                                                _ = v91
                                                                            } else {
                                                                                if ((((fi90 >> (((v3 - (v3 - (v2 >> (((^(v2))) & 63))))) & 63)) == (fi77 + fi90)) && ((v61 * float64(v84)) != v61)) || ((v89 / ((v89) | 1)) < (v89 * v89))) {
                                                                                    var v92 uint16 = 1
                                                                                    var v93 uint = 4294967295
                                                                                    var v94 uint8 = 127
                                                                                    _ = v92
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                } else {
                                                                                    switch (fi77 >> ((v3) & 63)) {
                                                                                    case 32768:
                                                                                        var v95 uint64 = 42
                                                                                        _ = v95
                                                                                    }
                                                                                    v1 = append(v1, (fi77 * (fi77 >> ((4294967295) & 63))))
                                                                                }
                                                                                fmt.Println(v79, fi77, v2)
                                                                                if false {
                                                                                    if (1 >= -128) {
                                                                                        var v96 int8 = 42
                                                                                        v97 := GenS7{F0: true}
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                    }
                                                                                } else {
                                                                                    v84 = (v84 + (v84 * (float32(float64(v84)) / (0.0 - v84))))
                                                                                    v89 -= -319
                                                                                }
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                            fmt.Println(fi90)
                                                                            var v98 uint = (v3 << ((v3) & 63))
                                                                            _ = v98
                                                                        } else {
                                                                            switch ("\n\t" + "foo bar") {
                                                                            case "":
                                                                                var v99 int32 = int32(v79)
                                                                                _ = v99
                                                                            default:
                                                                                switch fi90 {
                                                                                case 4294967295:
                                                                                    var v100 int64 = v89
                                                                                    v61 = v61
                                                                                    _ = v100
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(v89, v3)
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (fi90 / ((fi90) | 1))
                                                                                    }
                                                                                case 1:
                                                                                    var v101 uint8 = 37
                                                                                    v102 := GenS9{F0: 841, F1: (fi90 + fi90)}
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                }
                                                                            }
                                                                        }
                                                                        var v103 bool = true
                                                                        v89 = 174
                                                                        _ = v103
                                                                    }
                                                                    if (((1.5 / (v84 + v84)) * v84) <= float32(v61)) {
                                                                        v89 += v89
                                                                        var v104 int16 = 1
                                                                        _ = v104
                                                                    } else {
                                                                        fmt.Println(v84, v89)
                                                                    }
                                                                } else {
                                                                    var v105 uint16 = 175
                                                                    if false {
                                                                        v84 -= 0.0
                                                                        v106 := GenS4{F0: fi90}
                                                                        fmt.Println(v3)
                                                                        v89 = v89
                                                                        _ = v106
                                                                    } else {
                                                                        if (939 == 825) {
                                                                            switch (-(fi77)) {
                                                                            case 1:
                                                                                if ((^(v3)) < v2) {
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    v105 = 928
                                                                                    v107 := GenS7{F0: false}
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = fi90
                                                                                    }
                                                                                    var v108 int16 = 0
                                                                                    var v109 int8 = -1
                                                                                    _ = v107
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                } else {
                                                                                    switch int(v105) {
                                                                                    case 0:
                                                                                        fmt.Println(v105, v3)
                                                                                    case -990:
                                                                                        v89 += (v89 % ((9223372036854775807) | 1))
                                                                                    default:
                                                                                        v110 := GenS10{F0: (fi77 & fi90)}
                                                                                        _ = v110
                                                                                    }
                                                                                    v3 = (v3 | v3)
                                                                                }
                                                                            case -963:
                                                                                v89 = v89
                                                                                fmt.Println(fi90, v84)
                                                                            case 614:
                                                                                var v111 uint32 = (v79 - 4294967295)
                                                                                _ = v111
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v84, v2)
                                                                            if ((("hello" + ("" + "")) != "a") || ((v61 >= 3.14159) && ((fi90 >> ((v3) & 63)) != (fi90 - 127)))) {
                                                                                v112 := GenS11{F0: -1, F1: 208}
                                                                                var v113 uint8 = 1
                                                                                var v114 int = (fi77 ^ fi77)
                                                                                v79 += (v79 << ((4294967295) & 63))
                                                                                switch (fi77 << (((v3 % ((uint(v113)) | 1))) & 63)) {
                                                                                case 127:
                                                                                    var v115 int8 = 100
                                                                                    _ = v115
                                                                                case 128:
                                                                                    v89 -= int64(v114)
                                                                                case 0:
                                                                                    fmt.Println(v114, fi77)
                                                                                default:
                                                                                    if false {
                                                                                        var v116 int8 = int8(v105)
                                                                                        v61 = v61
                                                                                        _ = v116
                                                                                    }
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    var v117 uint = v2
                                                                                    _ = v117
                                                                                case "":
                                                                                    v118 := GenS8{F0: v114}
                                                                                    _ = v118
                                                                                }
                                                                                var v119 bool = (4294967295 == (v3 + (v2 % ((uint(v113)) | 1))))
                                                                                _ = v112
                                                                                _ = v113
                                                                                _ = v114
                                                                                _ = v119
                                                                            }
                                                                        }
                                                                    }
                                                                    v120 := GenS1{F0: 1, F1: "", F2: -32768}
                                                                    _ = v105
                                                                    _ = v120
                                                                }
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v121 float32 = (1.5 * 1e10)
                                                                    _ = v121
                                                                case "foo bar":
                                                                    v122 := GenS4{F0: 599}
                                                                    var v123 float32 = ((v84 * (3.14159 / ((1e10 / 1.0) * float32(0.1)))) - float32(v61))
                                                                    v1 = append(v1, (fi90 + int(v79)))
                                                                    _ = v122
                                                                    _ = v123
                                                                default:
                                                                    if true {
                                                                        v89 -= (v89 ^ (v89 << (((v2 + 2)) & 63)))
                                                                        v1 = append(v1, fi90)
                                                                    }
                                                                }
                                                                v124 := GenS12{F0: false, F1: 255, F2: -313}
                                                                _ = fi90
                                                                _ = v124
                                                            }
                                                            fmt.Println(v3, v84, v3)
                                                            _ = v89
                                                        }
                                                    }
                                                    _ = v84
                                                } else {
                                                    for fi125 := 0; fi125 < 5; fi125++ {
                                                        var v126 int16 = 675
                                                        _ = len(v1)
                                                        fmt.Println(v61)
                                                        _ = fi125
                                                        _ = v126
                                                    }
                                                    fmt.Println(fi77, v61)
                                                    v3 -= 4294967295
                                                }
                                            case 4294967295:
                                                var v127 uint64 = 33
                                                _ = v127
                                            case 0:
                                                for fi128 := 0; fi128 < 14; fi128++ {
                                                    v129 := GenS8{F0: fi128}
                                                    v79 = v79
                                                    switch (fi77 + fi128) {
                                                    case 1:
                                                        var v130 bool = true
                                                        _ = v130
                                                    }
                                                    v3 -= (v2 & v2)
                                                    if len(v1) > 0 {
                                                        v1[0] = -1
                                                    }
                                                    if false {
                                                        v3 += (v3 << (((v3 % ((v2) | 1))) & 63))
                                                        if (2 <= 255) {
                                                            var v131 int64 = 42
                                                            fmt.Println(fi128, fi128)
                                                            var v132 int32 = 27
                                                            _ = v131
                                                            _ = v132
                                                        }
                                                    } else {
                                                        var v133 bool = (-141 < 2147483647)
                                                        _ = v133
                                                    }
                                                    _ = fi128
                                                    _ = v129
                                                }
                                            default:
                                                for fi134 := 0; fi134 < 11; fi134++ {
                                                    switch fi134 {
                                                    case -738:
                                                        if (!false) {
                                                            v1 = append(v1, fi77)
                                                            var v135 int16 = 32767
                                                            _ = v135
                                                        }
                                                    }
                                                    var v136 uint32 = 942
                                                    _ = fi134
                                                    _ = v136
                                                }
                                            }
                                        } else {
                                            for fi137 := 0; fi137 < 16; fi137++ {
                                                if len(v1) > 0 {
                                                    v1[0] = (fi137 - (fi137 / ((fi77) | 1)))
                                                }
                                                _ = fi137
                                            }
                                            v79 -= v79
                                            if len(v1) > 0 {
                                                v1[0] = -558
                                            }
                                        }
                                        v138 := GenS12{F0: (!false), F1: 42, F2: 0}
                                        _ = v138
                                    }
                                    for fi139 := 0; fi139 < 2; fi139++ {
                                        v140 := GenS3{F0: 42, F1: -9223372036854775808, F2: 629}
                                        v79 = v79
                                        _ = fi139
                                        _ = v140
                                    }
                                }
                                for fi141 := 0; fi141 < 18; fi141++ {
                                    v61 += v61
                                    if true {
                                        v142 := GenS13{F0: -53, F1: "", F2: 0}
                                        _ = v142
                                    }
                                    _ = fi141
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            var v143 float64 = v61
                            _ = v143
                            _ = v79
                        } else {
                            v61 -= 1.5
                            if (!true) {
                                fmt.Println(v2)
                                v1 = append(v1, (fi77 << (((v3 / (((v3 % ((v3) | 1))) | 1))) & 63)))
                                if (-1 <= 42) {
                                    var v144 float32 = (((-2.25 * 3.14159) + 1e10) - 3.14159)
                                    for fi145 := 0; fi145 < 10; fi145++ {
                                        _ = len(v1)
                                        v61 = 1e10
                                        _ = fi145
                                    }
                                    for fi146 := 0; fi146 < 19; fi146++ {
                                        v61 = 0.1
                                        v144 = v144
                                        var v147 int32 = 0
                                        _ = fi146
                                        _ = v147
                                    }
                                    _ = v144
                                } else {
                                    v2 = v2
                                    v148 := GenS14{F0: 127}
                                    for fi149 := 0; fi149 < 19; fi149++ {
                                        v150 := GenS15{F0: 42}
                                        if ((!true) && false) {
                                            var v151 uint8 = 2
                                            _ = v151
                                        } else {
                                            var v152 bool = false
                                            _ = v152
                                        }
                                        _ = fi149
                                        _ = v150
                                    }
                                    _ = v148
                                }
                            }
                            switch 858 {
                            case 127:
                                v153 := GenS2{F0: (v3 & 255)}
                                _ = v153
                            case 0:
                                for fi154 := 0; fi154 < 16; fi154++ {
                                    v2 = v3
                                    v61 -= v61
                                    _ = fi154
                                }
                            }
                        }
                    case "":
                        if (248 <= 0) {
                            v155 := GenS6{F0: (v3 % (((v2 | 2)) | 1)), F1: (42 >= 208), F2: 100}
                            v156 := GenS16{F0: (v2 >> (((^(v3))) & 63))}
                            _ = v155
                            _ = v156
                        } else {
                            v1 = append(v1, fi77)
                            switch "foo bar" {
                            case "foo bar":
                                switch (fi77 << ((18446744073709551615) & 63)) {
                                case 42:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            if false {
                                v157 := GenS6{F0: (^(v2)), F1: (!(false || false)), F2: 100}
                                if false {
                                    fmt.Println(v61)
                                }
                                if (0 > 42) {
                                    if len(v1) > 0 {
                                        v1[0] = (fi77 >> (((v2 / ((v3) | 1))) & 63))
                                    }
                                    for fi158 := 0; fi158 < 6; fi158++ {
                                        v159 := GenS17{F0: (v3 & 42), F1: 33, F2: 127}
                                        if (true || (true || false)) {
                                            var v160 uint32 = 2147483648
                                            _ = v160
                                        } else {
                                            var v161 int16 = 1000
                                            fmt.Println(fi77, v2)
                                            switch (fi77 % ((fi77) | 1)) {
                                            case -857:
                                                var v162 uint = (v3 ^ v2)
                                                _ = v162
                                            case 4294967295:
                                                fmt.Println(v2, fi77, v61)
                                            case 127:
                                                v161 = -564
                                            }
                                            v163 := GenS18{F0: 0, F1: -9223372036854775808}
                                            _ = v161
                                            _ = v163
                                        }
                                        _ = fi158
                                        _ = v159
                                    }
                                }
                                _ = v157
                            } else {
                                var v164 uint32 = 44
                                fmt.Println(fi77, v2, v61)
                                v2 = (v3 & v2)
                                _ = v164
                            }
                        }
                    }
                    switch "\n\t" {
                    case "":
                        v2 -= v3
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi77
                }
            }
        default:
            var v165 uint = 482
            _ = v165
        }
    }
    for fi166 := 0; fi166 < 8; fi166++ {
        fmt.Println(v61)
        fmt.Println(v2)
        v1 = append(v1, (^(fi166)))
        _ = fi166
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v61
}
