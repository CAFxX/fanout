package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint; F2 uint8 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint; F1 bool; F2 int64 }
type GenS9 struct { F0 int8; F1 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int; F1 int32; F2 bool }
type GenS13 struct { F0 int8 }

func main() {
    var v1 float64 = (float64(0.0) - 0.1)
    var v2 uint = 941
    v3 := make(map[uint32]string)
    var v4 int32 = 2147483647
    var v5 int8 = 9
    var v6 uint32 = 0
    if (42 >= 232) {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 float32 = float32(v1)
            v3[uint32(v2)] = ("a" + "hello")
            var v9 bool = (!false)
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int8 = -128
        _ = v10
    }
    v11 := GenS1{F0: -834, F1: 128}
    v2 = 0
    var v12 uint64 = 427
    var v13 uint32 = (v6 << ((v2) & 63))
    var v14 uint32 = (v6 & v6)
    v15 := v3[v6]
    if false {
        delete(v3, 4294967295)
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(v13, v14, v14)
            switch v15 {
            case "":
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 int = (fi17 | (fi16 / (((fi17 ^ (fi17 & fi16))) | 1)))
                    var v19 float32 = 3.14159
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            default:
                if (false && (false && true)) {
                    switch (fi16 & 128) {
                    case 4294967295:
                        for fi20 := 0; fi20 < 20; fi20++ {
                            fmt.Println(fi16, fi20)
                            if false {
                                delete(v3, v13)
                            } else {
                                var v21 float32 = 3.14159
                                _ = v21
                            }
                            v22 := GenS1{F0: (-(v4)), F1: 174}
                            _ = fi20
                            _ = v22
                        }
                    case 32768:
                        var v23 uint32 = v6
                        _ = v23
                    }
                    var v24 uint8 = 242
                    _ = v24
                } else {
                    v25 := []int{int(v14), 0, (-(fi16))}
                    fmt.Println(v5, v2, v15)
                    _ = v25
                }
            }
            _ = fi16
        }
        v13 = 2147483647
    }
    v14 -= 128
    if (v2 < (^(v2))) {
        for fi26 := 0; fi26 < 9; fi26++ {
            v27 := GenS2{F0: v5}
            if (false || false) {
                if true {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        var v29 uint8 = 255
                        if false {
                            var v30 int16 = 42
                            v31 := []int{(fi26 - (fi28 - fi28))}
                            if (42 > 50) {
                                v32 := GenS3{F0: 1, F1: false}
                                fmt.Println(v1)
                                _ = v32
                            } else {
                                var v33 string = v15
                                var v34 int = fi28
                                v2 -= v2
                                _ = v33
                                _ = v34
                            }
                            _ = v30
                            _ = v31
                        }
                        _ = fi28
                        _ = v29
                    }
                } else {
                    fmt.Println(v2)
                    if (!(-173 != 554)) {
                        switch ((v15 + "foo bar") + "") {
                        case "\n\t":
                            v2 = 18446744073709551615
                        case "a":
                            var v35 int16 = 1
                            _ = v35
                        }
                        if true {
                            for fi36 := 0; fi36 < 20; fi36++ {
                                v14 -= v13
                                if (false || (true || (!false))) {
                                    v12 = (v12 * 18446744073709551615)
                                    fmt.Println(v4)
                                    v14 -= v14
                                }
                                _ = fi36
                            }
                            for fi37 := 0; fi37 < 6; fi37++ {
                                if false {
                                    v14 += 42
                                    fmt.Println(v5, v2)
                                    switch (fi26 << (((v2 | 18446744073709551615)) & 63)) {
                                    case 1:
                                        if true {
                                            if (23 > 1) {
                                                var v38 uint = (v2 / ((42) | 1))
                                                _ = v38
                                            } else {
                                                fmt.Println(v5)
                                                v39 := GenS2{F0: v5}
                                                v40 := GenS4{F0: fi26}
                                                if false {
                                                    fmt.Println(v12)
                                                }
                                                _ = v39
                                                _ = v40
                                            }
                                            v1 = 0.0
                                            if true {
                                                var v41 int = fi26
                                                var v42 uint = (v2 ^ v2)
                                                var v43 float64 = v1
                                                var v44 uint32 = v14
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            } else {
                                                v12 = v12
                                                switch "" {
                                                case "a":
                                                    v5 = v5
                                                case "hello":
                                                    fmt.Println(v2, v6)
                                                }
                                            }
                                        } else {
                                            v3[(v13 | v6)] = ("a" + "foo bar")
                                            v5 += (v5 % ((v5) | 1))
                                            v45 := GenS2{F0: 97}
                                            _ = v45
                                        }
                                    }
                                } else {
                                    if (!false) {
                                        if ((^(v2)) != (v2 ^ (v2 + (^(v2))))) {
                                            fmt.Println(v13, v6)
                                            delete(v3, 2147483647)
                                            delete(v3, (v13 >> ((18446744073709551615) & 63)))
                                        } else {
                                            v46 := GenS5{F0: (v4 * v4), F1: uint(v14), F2: uint8(v2)}
                                            v6 = (v6 / (((v13 & v14)) | 1))
                                            if false {
                                                var v47 string = v15
                                                if (((v12 ^ (^(v12))) == v12) && (((v4 ^ v4) > 42) && true)) {
                                                    var v48 int16 = 1
                                                    fmt.Println(fi37, v2, v6)
                                                    switch (v15 + v15) {
                                                    case "a":
                                                        if false {
                                                            v49 := GenS6{F0: v5, F1: 126}
                                                            v4 -= 2147483647
                                                            v50 := []int{fi26}
                                                            if true {
                                                                switch int(v2) {
                                                                case -1:
                                                                    switch (v47 + "foo bar") {
                                                                    case "\n\t":
                                                                        switch -1 {
                                                                        case -2147483648:
                                                                            if len(v50) >= 2 {
                                                                                v50 = v50[1:len(v50)]
                                                                            }
                                                                        case 32767:
                                                                            v15 = "foo bar"
                                                                            if len(v50) > 0 {
                                                                                v50[0] = fi26
                                                                            }
                                                                        default:
                                                                            switch 32768 {
                                                                            case -511:
                                                                                fmt.Println(v4)
                                                                                switch v47 {
                                                                                case "hello":
                                                                                    v47 = "\n\t"
                                                                                case "foo bar":
                                                                                    v50 = append(v50, fi37)
                                                                                case "":
                                                                                    fmt.Println(v1, v15)
                                                                                    v3[v6] = v47
                                                                                    var v51 string = (v47 + ((v47 + v15) + ("\n\t" + v47)))
                                                                                    _ = v51
                                                                                }
                                                                            default:
                                                                                var v52 int8 = (v5 / ((v5) | 1))
                                                                                _ = v52
                                                                            }
                                                                            v48 = 0
                                                                        }
                                                                    case "hello":
                                                                        var v53 bool = (v5 < v5)
                                                                        _ = v53
                                                                    case "":
                                                                        if len(v50) >= 2 {
                                                                            v50 = v50[1:len(v50)]
                                                                        }
                                                                    default:
                                                                        if len(v50) > 0 {
                                                                            v50[0] = fi37
                                                                        }
                                                                    }
                                                                }
                                                                v14 -= v14
                                                                var v54 float32 = 1.5
                                                                if true {
                                                                    switch v47 {
                                                                    case "hello":
                                                                        v55 := GenS2{F0: v5}
                                                                        _ = v55
                                                                    case "foo bar":
                                                                        switch v15 {
                                                                        case "hello":
                                                                            if (v1 < 1e10) {
                                                                                var v56 int64 = -1
                                                                                _ = v56
                                                                            }
                                                                        case "":
                                                                            if len(v50) > 0 {
                                                                                v50[0] = int(v4)
                                                                            }
                                                                        }
                                                                    default:
                                                                        v13 += v13
                                                                    }
                                                                    switch (v47 + v47) {
                                                                    case "hello":
                                                                        if ((v48 ^ (^(v48))) == -1) {
                                                                            if true {
                                                                                var v57 uint8 = 255
                                                                                if true {
                                                                                    if (true && true) {
                                                                                        fmt.Println(v2)
                                                                                        v58 := GenS4{F0: fi26}
                                                                                        _ = v58
                                                                                    } else {
                                                                                        _ = len(v50)
                                                                                        _ = len(v3)
                                                                                    }
                                                                                    switch fi26 {
                                                                                    case -969:
                                                                                        var v59 int64 = 1
                                                                                        _ = v59
                                                                                    case 32768:
                                                                                        var v60 int = 68
                                                                                        _ = v60
                                                                                    default:
                                                                                        var v61 int32 = v4
                                                                                        _ = v61
                                                                                    }
                                                                                }
                                                                                var v62 int32 = (v4 ^ v4)
                                                                                _ = v57
                                                                                _ = v62
                                                                            } else {
                                                                                _ = len(v50)
                                                                                var v63 uint32 = 65535
                                                                                _ = v63
                                                                            }
                                                                            var v64 string = ((v47 + v15) + ("" + v15))
                                                                            _ = v64
                                                                        } else {
                                                                            switch "" {
                                                                            case "a":
                                                                                v54 += 3.14159
                                                                            case "\n\t":
                                                                                v54 += ((((v54 / 0.1) - 0.1) - 1e10) / -2.25)
                                                                                if ("a" < ("hello" + v15)) {
                                                                                    fmt.Println(fi26, v6, v13)
                                                                                    var v65 bool = true
                                                                                    fmt.Println(fi37, v6)
                                                                                    _ = v65
                                                                                }
                                                                            case "hello":
                                                                                if false {
                                                                                    switch (fi26 >> ((v2) & 63)) {
                                                                                    case -2147483648:
                                                                                        if true {
                                                                                            v66 := GenS2{F0: 1}
                                                                                            if len(v50) > 0 {
                                                                                                v50[0] = (fi26 << (((v2 / ((v2) | 1))) & 63))
                                                                                            }
                                                                                            v3[v13] = v15
                                                                                            _ = v66
                                                                                        }
                                                                                    case 42:
                                                                                        var v67 string = "a"
                                                                                        _ = v67
                                                                                    default:
                                                                                        v54 += 0.0
                                                                                    }
                                                                                } else {
                                                                                    var v68 int16 = -32768
                                                                                    var v69 bool = false
                                                                                    var v70 uint8 = 0
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                }
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        var v71 float64 = float64(float32(-2.25))
                                                                        _ = v71
                                                                    case "a":
                                                                        if false {
                                                                            if (2147483648 <= v14) {
                                                                                var v72 float64 = 0.1
                                                                                v4 -= (v4 << ((1) & 63))
                                                                                _ = v72
                                                                            } else {
                                                                                if true {
                                                                                    var v73 uint8 = 2
                                                                                    v50 = append(v50, (-(fi37)))
                                                                                    _ = v73
                                                                                }
                                                                                if true {
                                                                                    fmt.Println(v5)
                                                                                } else {
                                                                                    if (true && ((v12 & (v12 / ((v12) | 1))) == 42)) {
                                                                                        v74 := GenS1{F0: v4, F1: 222}
                                                                                        _ = v74
                                                                                    }
                                                                                    _ = len(v3)
                                                                                    v75 := GenS7{F0: v4, F1: 255}
                                                                                    if (!((v47 <= v47) || true)) {
                                                                                        switch (v47 + "hello") {
                                                                                        case "a":
                                                                                            v12 = uint64(v2)
                                                                                        case "foo bar":
                                                                                            v50 = append(v50, 32768)
                                                                                        case "\n\t":
                                                                                            if (42 >= v12) {
                                                                                                if (!false) {
                                                                                                    if true {
                                                                                                        fmt.Println(v2, v6)
                                                                                                        v1 = v1
                                                                                                        switch (fi37 % ((fi37) | 1)) {
                                                                                                        case 2147483647:
                                                                                                            var v76 uint32 = v14
                                                                                                            _ = v76
                                                                                                        case 1:
                                                                                                            v50 = append(v50, (fi37 + fi37))
                                                                                                        case 9223372036854775807:
                                                                                                            if len(v50) > 0 {
                                                                                                                v50[0] = 128
                                                                                                            }
                                                                                                        }
                                                                                                        var v77 uint64 = v12
                                                                                                        _ = v77
                                                                                                    } else {
                                                                                                        var v78 bool = true
                                                                                                        switch fi26 {
                                                                                                        case -70:
                                                                                                            v3[v14] = v47
                                                                                                        default:
                                                                                                            v13 -= (v14 / ((v13) | 1))
                                                                                                        }
                                                                                                        _ = v78
                                                                                                    }
                                                                                                    _ = len(v50)
                                                                                                }
                                                                                                v1 = v1
                                                                                                v1 += float64(((v54 + v54) - v54))
                                                                                            } else {
                                                                                                if true {
                                                                                                    if (42 <= 32767) {
                                                                                                        v79 := GenS8{F0: 65535, F1: (42 == (fi26 / ((9223372036854775807) | 1))), F2: -890}
                                                                                                        _ = v79
                                                                                                    }
                                                                                                    var v80 float32 = (v54 * float32(0.1))
                                                                                                    if len(v50) > 0 {
                                                                                                        v50[0] = fi26
                                                                                                    }
                                                                                                    _ = v80
                                                                                                } else {
                                                                                                    v14 += v13
                                                                                                    v48 -= v48
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v47 = v15
                                                                                    }
                                                                                    if false {
                                                                                        v3[(v14 * (v14 * (v14 ^ 42)))] = ("" + ((v47 + v15) + "foo bar"))
                                                                                    } else {
                                                                                        v14 = 2147483648
                                                                                        var v81 bool = ((v14 % ((v6) | 1)) < 1)
                                                                                        _ = v81
                                                                                    }
                                                                                    _ = v75
                                                                                }
                                                                                var v82 string = v47
                                                                                var v83 uint = v2
                                                                                var v84 int64 = 0
                                                                                v6 += uint32(v83)
                                                                                _ = v82
                                                                                _ = v83
                                                                                _ = v84
                                                                            }
                                                                            var v85 int8 = (-(v5))
                                                                            _ = v85
                                                                        }
                                                                    default:
                                                                        v86 := GenS7{F0: 42, F1: 65535}
                                                                        _ = v86
                                                                    }
                                                                } else {
                                                                    v3[(^(v6))] = v47
                                                                }
                                                                switch (fi26 | (fi26 >> (((^(v2))) & 63))) {
                                                                case -1:
                                                                    var v87 float64 = v1
                                                                    _ = v87
                                                                default:
                                                                    if (!false) {
                                                                        if (793 > 32767) {
                                                                            var v88 uint = (v2 & v2)
                                                                            _ = v88
                                                                        }
                                                                        v12 = 1
                                                                        v6 = (^(v14))
                                                                    } else {
                                                                        if (!true) {
                                                                            v89 := GenS9{F0: v5, F1: 231}
                                                                            v4 -= int32(v5)
                                                                            _ = v89
                                                                        }
                                                                        v2 = v2
                                                                        v90 := GenS5{F0: v4, F1: v2, F2: 255}
                                                                        _ = v90
                                                                    }
                                                                }
                                                                _ = v54
                                                            }
                                                            v47 = v47
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            delete(v3, v13)
                                                            switch -756 {
                                                            case 42:
                                                                v91 := []int{fi37, fi26, (fi26 & (fi26 / (((fi26 >> (((v2 + v2)) & 63))) | 1))), fi26}
                                                                _ = v91
                                                            case 128:
                                                                v4 = (v4 - (v4 % ((v4) | 1)))
                                                            case 127:
                                                                fmt.Println(v48, v5)
                                                            default:
                                                                if false {
                                                                    var v92 uint64 = v12
                                                                    fmt.Println(fi26, v6, v13)
                                                                    v12 -= 2
                                                                    _ = v92
                                                                }
                                                            }
                                                        }
                                                    case "":
                                                        var v93 string = ""
                                                        _ = v93
                                                    case "foo bar":
                                                        v94 := []int{32768, (fi26 | 32767)}
                                                        _ = v94
                                                    default:
                                                        var v95 int8 = (v5 + v5)
                                                        _ = v95
                                                    }
                                                    _ = v48
                                                }
                                                v6 = v13
                                                _ = v47
                                            }
                                            _ = v46
                                        }
                                        v96 := GenS1{F0: (v4 - 2147483647), F1: 255}
                                        var v97 bool = (1 <= (v12 * (v12 % ((18446744073709551615) | 1))))
                                        v98 := GenS10{F0: 589}
                                        v6 = 380
                                        _ = v96
                                        _ = v97
                                        _ = v98
                                    }
                                }
                                v1 = (float64(0.0) + v1)
                                _ = fi37
                            }
                            if (!false) {
                                v12 = 4294967295
                            } else {
                                if (false || (995 == (v4 & v4))) {
                                    for fi99 := 0; fi99 < 19; fi99++ {
                                        v14 += 42
                                        _ = fi99
                                    }
                                    v100 := GenS10{F0: v14}
                                    for fi101 := 0; fi101 < 1; fi101++ {
                                        fmt.Println(v13, v2, fi101)
                                        v4 += (v4 - v4)
                                        _ = fi101
                                    }
                                    _ = v100
                                }
                                var v102 uint8 = 1
                                _ = v102
                            }
                        } else {
                            for fi103 := 0; fi103 < 8; fi103++ {
                                var v104 uint = (^(v2))
                                if true {
                                    v6 = v14
                                    v2 -= v104
                                    if false {
                                        v105 := GenS7{F0: v4, F1: 2}
                                        if true {
                                            var v106 int16 = -1
                                            var v107 int64 = int64(v12)
                                            _ = len(v3)
                                            _ = v106
                                            _ = v107
                                        }
                                        v104 -= (v2 * (v2 & (v2 - 0)))
                                        v108 := []int{(^(fi103)), fi103, 2147483647}
                                        _ = v105
                                        _ = v108
                                    }
                                } else {
                                    v109 := GenS4{F0: fi26}
                                    switch (("\n\t" + "\n\t") + v15) {
                                    case "a":
                                        var v110 bool = ((v12 + (v12 << (((v2 * (v104 | v2))) & 63))) != v12)
                                        _ = v110
                                    case "\n\t":
                                        var v111 uint32 = 578
                                        _ = v111
                                    }
                                    v112 := v3[(v13 >> ((v2) & 63))]
                                    _ = v109
                                    _ = v112
                                }
                                v6 += 1
                                fmt.Println(v15, fi26, v4)
                                _ = fi103
                                _ = v104
                            }
                        }
                        for fi113 := 0; fi113 < 16; fi113++ {
                            v6 -= (v13 ^ (v6 - (^(v13))))
                            v4 -= (v4 - (v4 ^ (v4 & (v4 & v4))))
                            var v114 string = ""
                            v115 := GenS11{F0: 28}
                            _ = fi113
                            _ = v114
                            _ = v115
                        }
                    } else {
                        v116 := GenS12{F0: -734, F1: v4, F2: ((!true) && false)}
                        v12 += v12
                        for fi117 := 0; fi117 < 9; fi117++ {
                            v118 := []int{(fi117 - 1), 0, fi26, fi26}
                            _ = fi117
                            _ = v118
                        }
                        _ = v116
                    }
                }
                var v119 int64 = 398
                if (true && true) {
                    v1 -= (1.5 - v1)
                    v120 := GenS11{F0: (v119 - v119)}
                    if true {
                        switch fi26 {
                        case 42:
                            var v121 uint = 42
                            _ = v121
                        case 947:
                            if false {
                                v122 := v3[1]
                                _ = v122
                            }
                        case 127:
                            fmt.Println(fi26, v5)
                        }
                        delete(v3, (^(v6)))
                        for fi123 := 0; fi123 < 7; fi123++ {
                            v3[v6] = v15
                            v124 := []int{(fi26 ^ fi26), (fi26 % ((fi123) | 1)), fi123, (-(fi26))}
                            _ = len(v3)
                            _ = fi123
                            _ = v124
                        }
                    }
                    _ = v120
                }
                _ = v119
            }
            for fi125 := 0; fi125 < 4; fi125++ {
                v12 = (v12 * v12)
                _ = fi125
            }
            _ = fi26
            _ = v27
        }
        if true {
            switch 128 {
            case 128:
                v6 = v6
            default:
                if ((v5 << ((v2) & 63)) != -1) {
                    v126 := GenS1{F0: v4, F1: uint8(v13)}
                    if (true || false) {
                        v4 += 518
                        v127 := GenS13{F0: -128}
                        _ = v127
                    }
                    var v128 uint = 2
                    _ = v126
                    _ = v128
                } else {
                    if true {
                        fmt.Println(v6, v15)
                    }
                }
            }
            var v129 int32 = (v4 ^ (-(v4)))
            _ = v129
        }
        var v130 int64 = -1
        _ = v130
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
