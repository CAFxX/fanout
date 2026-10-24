package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint8; F1 string; F2 bool }
type GenS7 struct { F0 uint32; F1 string; F2 uint8 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 uint; F1 string; F2 int8 }
type GenS11 struct { F0 int64; F1 uint32 }
type GenS12 struct { F0 int64; F1 int16 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int64; F1 uint; F2 bool }
type GenS15 struct { F0 uint; F1 uint32 }
type GenS16 struct { F0 uint64; F1 int32 }
type GenS17 struct { F0 int64; F1 int16; F2 uint64 }
type GenS18 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 0
    var v3 int64 = -904
    if (!((9223372036854775807 == (^(v3))) && (v3 > (v3 << ((65535) & 63))))) {
        var v4 int64 = (-(v3))
        v2 = v2
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 10; fi6++ {
                v4 -= v4
                v7 := []int{32767, 993}
                _ = fi6
                _ = v7
            }
            if true {
                if true {
                    var v8 int32 = 126
                    var v9 uint32 = 525
                    var v10 uint16 = 42
                    _ = v10
                    _ = v8
                    _ = v9
                }
                fmt.Println(v4, v3, v3)
                var v11 uint64 = (v2 * 0)
                _ = v11
            } else {
                var v12 uint = 772
                if true {
                    if (65535 != 221) {
                        v13 := GenS1{F0: -1}
                        if ("\n\t" > "foo bar") {
                            v14 := []int{(fi5 * fi5), (fi5 % ((fi5) | 1))}
                            var v15 int64 = 9223372036854775807
                            var v16 int8 = 1
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            var v17 uint8 = 255
                            if (((((-72 < 94) || false) && (!true)) || false) && true) {
                                if false {
                                    v2 += (v2 & (v2 + (v2 - v2)))
                                    v3 -= -1
                                }
                            } else {
                                if false {
                                    switch -452 {
                                    case 28:
                                        if false {
                                            switch "foo bar" {
                                            case "":
                                                switch "foo bar" {
                                                case "foo bar":
                                                    switch ((("a" + ("hello" + "foo bar")) + "\n\t") + "hello") {
                                                    case "a":
                                                        fmt.Println(v2, v3, fi5)
                                                    }
                                                case "":
                                                    if false {
                                                        v18 := GenS1{F0: 662}
                                                        var v19 bool = true
                                                        fmt.Println(fi5, v12)
                                                        _ = v18
                                                        _ = v19
                                                    } else {
                                                        for fi20 := 0; fi20 < 2; fi20++ {
                                                            fmt.Println(v3, v12)
                                                            if true {
                                                                if false {
                                                                    if true {
                                                                        v21 := []int{204, fi5}
                                                                        _ = v21
                                                                    }
                                                                    var v22 uint16 = uint16(v12)
                                                                    if false {
                                                                        v23 := GenS1{F0: -346}
                                                                        var v24 int32 = -481
                                                                        v3 = (v3 * v4)
                                                                        var v25 uint = v12
                                                                        _ = v23
                                                                        _ = v24
                                                                        _ = v25
                                                                    }
                                                                    _ = v22
                                                                } else {
                                                                    v3 -= 177
                                                                    switch int(v17) {
                                                                    case 112:
                                                                        v4 = 9223372036854775807
                                                                    case 366:
                                                                        v4 -= v4
                                                                        v26 := []int{(fi20 << ((4294967295) & 63)), (fi20 << (((^(v12))) & 63)), (fi5 & (fi5 << ((v12) & 63)))}
                                                                        v3 -= (v4 | v3)
                                                                        var v27 int16 = 1000
                                                                        _ = v26
                                                                        _ = v27
                                                                    case 365:
                                                                        if true {
                                                                            var v28 int = -1
                                                                            v29 := GenS1{F0: -2147483648}
                                                                            var v30 bool = (true || (995 <= 2))
                                                                            var v31 uint8 = (v17 & v17)
                                                                            switch -9223372036854775808 {
                                                                            case -2147483648:
                                                                                fmt.Println(v30)
                                                                            }
                                                                            var v32 int8 = 0
                                                                            _ = v28
                                                                            _ = v29
                                                                            _ = v30
                                                                            _ = v31
                                                                            _ = v32
                                                                        }
                                                                    }
                                                                    var v33 uint64 = (v2 ^ uint64(v17))
                                                                    v2 -= v33
                                                                    fmt.Println(v17, v3)
                                                                    v12 += v12
                                                                    _ = v33
                                                                }
                                                                v34 := GenS2{F0: 1, F1: uint(v17)}
                                                                _ = v34
                                                            }
                                                            v35 := GenS3{F0: 1, F1: 2147483647}
                                                            _ = fi20
                                                            _ = v35
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                v2 -= (^(v2))
                                            }
                                            fmt.Println(v17, v4, v2)
                                            if (true || (1 != 595)) {
                                                _ = len(v1)
                                            }
                                        }
                                    default:
                                        v3 += -608
                                    }
                                } else {
                                    v12 = (v12 & (v12 & uint(v17)))
                                    var v36 uint8 = (v17 * 42)
                                    v3 = int64(v2)
                                    _ = v36
                                }
                                if (v12 > v12) {
                                    _ = len(v1)
                                }
                                if (!false) {
                                    var v37 bool = (!(1e10 == (1.5 - (((0.1 + float64(0.1)) - (0.0 - 0.0)) + (1e10 + 0.0)))))
                                    v38 := GenS4{F0: -1, F1: 0, F2: 692}
                                    _ = v37
                                    _ = v38
                                } else {
                                    var v39 int64 = int64(v17)
                                    _ = v39
                                }
                            }
                            _ = v17
                        }
                        _ = v13
                    } else {
                        var v40 uint8 = 244
                        v41 := []int{(-(fi5))}
                        var v42 uint64 = (v2 | v2)
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                }
                _ = v12
            }
            switch fi5 {
            case 52:
                for fi43 := 0; fi43 < 9; fi43++ {
                    v4 = (v3 ^ 42)
                    if (!((1 < -696) && (2 >= 1))) {
                        if (!false) {
                            if ("a" < "hello") {
                                v3 += v3
                            }
                        } else {
                            var v44 float32 = (0.0 * (3.14159 * 0.0))
                            if false {
                                var v45 uint32 = 165
                                if (1 == 429) {
                                    if true {
                                        if false {
                                            v46 := []int{}
                                            var v47 uint = 554
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                    if true {
                                        v48 := []int{(^(fi43)), fi5}
                                        switch "" {
                                        case "hello":
                                            v45 = v45
                                        case "a":
                                            v44 -= 1e10
                                        default:
                                            var v49 uint = 4294967295
                                            _ = v49
                                        }
                                        _ = v48
                                    }
                                }
                                switch -1 {
                                case -821:
                                    v50 := GenS3{F0: -240, F1: 0}
                                    _ = v50
                                case -2147483648:
                                    v1[1] = v45
                                case -766:
                                    switch "" {
                                    case "hello":
                                        if true {
                                            v51 := []int{fi43, fi5, fi5}
                                            delete(v1, 9223372036854775807)
                                            _ = v51
                                        } else {
                                            fmt.Println(v44, v4)
                                        }
                                    case "a":
                                        v52 := GenS5{F0: v45, F1: -1}
                                        _ = v52
                                    default:
                                        v3 -= (v3 << ((255) & 63))
                                    }
                                }
                                if (false || false) {
                                    if false {
                                        var v53 int8 = -128
                                        _ = v53
                                    }
                                    v54 := GenS6{F0: 1, F1: "a", F2: true}
                                    _ = v54
                                } else {
                                    v2 = 4294967295
                                    v4 += (v4 % (((v3 << ((815) & 63))) | 1))
                                    fmt.Println(v3, v4)
                                    v2 -= v2
                                    switch "hello" {
                                    case "":
                                        v55 := GenS5{F0: v45, F1: -408}
                                        _ = v55
                                    case "a":
                                        v56 := []int{(fi5 % ((fi5) | 1)), fi5, fi43, fi5}
                                        _ = v56
                                    default:
                                        if (1e10 <= (1.5 + (3.14159 / 0.1))) {
                                            v44 = v44
                                            switch (^(fi43)) {
                                            case 888:
                                                var v57 string = "hello"
                                                _ = v57
                                            default:
                                                v58 := v1[990]
                                                _ = v58
                                            }
                                        }
                                    }
                                }
                                v59 := GenS6{F0: 0, F1: (("" + "\n\t") + "\n\t"), F2: (0 > 1)}
                                _ = v45
                                _ = v59
                            }
                            _ = v44
                        }
                        v60 := GenS5{F0: 1, F1: int64(v2)}
                        if (!(((3.14159 - (1.5 / (float64(1.5) - float64(-2.25)))) * (1.5 + -2.25)) == (-2.25 / ((float64((1e10 - 0.1)) / -2.25) / float64(0.1))))) {
                            if (true && false) {
                                if false {
                                    v61 := GenS3{F0: -262, F1: -1}
                                    v62 := []int{}
                                    v63 := GenS7{F0: 577, F1: "hello", F2: 1}
                                    var v64 uint = 807
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                if false {
                                    switch ("\n\t" + "hello") {
                                    case "hello":
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (float64(3.14159) == 3.14159) {
                                                var v65 float32 = 1.5
                                                var v66 uint32 = 2147483648
                                                _ = v65
                                                _ = v66
                                            } else {
                                                var v67 int16 = -32768
                                                v2 -= (^(v2))
                                                v68 := GenS8{F0: 709, F1: ("a" + "hello")}
                                                v69 := GenS9{F0: 42, F1: 561}
                                                v70 := GenS2{F0: 556, F1: 255}
                                                v71 := []int{128}
                                                _ = v67
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                            }
                                        case "\n\t":
                                            if (864 <= 272) {
                                                fmt.Println(fi5, fi5)
                                                if false {
                                                    if true {
                                                        switch "\n\t" {
                                                        case "a":
                                                            if (1 > 42) {
                                                                delete(v1, 1)
                                                            }
                                                        case "":
                                                            _ = len(v1)
                                                            var v72 string = "\n\t"
                                                            _ = v72
                                                        case "\n\t":
                                                            v1[(^(fi43))] = 0
                                                        default:
                                                            switch "" {
                                                            case "\n\t":
                                                                switch ("\n\t" + "") {
                                                                case "a":
                                                                    if (!false) {
                                                                        fmt.Println(fi43, v4, fi43)
                                                                        v73 := v1[(fi43 ^ (fi5 ^ 42))]
                                                                        var v74 string = ""
                                                                        _ = v73
                                                                        _ = v74
                                                                    } else {
                                                                        delete(v1, fi43)
                                                                        var v75 string = "\n\t"
                                                                        _ = v75
                                                                    }
                                                                case "foo bar":
                                                                    v1[fi43] = 2147483647
                                                                case "hello":
                                                                    var v76 float64 = (-2.25 * (1.5 + ((0.0 * (0.0 * float64(1e10))) + -2.25)))
                                                                    _ = v76
                                                                }
                                                            default:
                                                                switch "a" {
                                                                case "hello":
                                                                    var v77 int16 = -612
                                                                    _ = v77
                                                                case "\n\t":
                                                                    v78 := GenS4{F0: 1, F1: 32767, F2: 1}
                                                                    _ = v78
                                                                case "a":
                                                                    switch fi43 {
                                                                    case -1:
                                                                        switch fi43 {
                                                                        case 42:
                                                                            v79 := GenS9{F0: 1, F1: 255}
                                                                            _ = v79
                                                                        case 128:
                                                                            var v80 int64 = -722
                                                                            _ = v80
                                                                        case -2147483648:
                                                                            fmt.Println(v4)
                                                                        }
                                                                    case 79:
                                                                        fmt.Println(v2)
                                                                    case 32767:
                                                                        var v81 float64 = 1e10
                                                                        _ = v81
                                                                    }
                                                                default:
                                                                    v2 += 838
                                                                }
                                                                if true {
                                                                    v82 := GenS10{F0: 452, F1: "a", F2: 34}
                                                                    var v83 uint = 18446744073709551615
                                                                    var v84 uint64 = 18446744073709551615
                                                                    var v85 uint8 = 1
                                                                    v86 := GenS9{F0: -1, F1: 42}
                                                                    _ = v82
                                                                    _ = v83
                                                                    _ = v84
                                                                    _ = v85
                                                                    _ = v86
                                                                } else {
                                                                    v4 -= 9223372036854775807
                                                                    v2 = v2
                                                                    switch 530 {
                                                                    case 995:
                                                                        v4 -= v3
                                                                        if true {
                                                                            _ = len(v1)
                                                                        }
                                                                        var v87 int32 = -997
                                                                        _ = v87
                                                                    }
                                                                    v88 := GenS11{F0: (^(v3)), F1: 42}
                                                                    switch (("hello" + ("foo bar" + "")) + "hello") {
                                                                    case "hello":
                                                                        v89 := []int{487, fi5, (fi5 / (((fi43 | (fi5 / ((int(v3)) | 1)))) | 1))}
                                                                        _ = v89
                                                                    case "a":
                                                                        var v90 uint32 = 91
                                                                        _ = v90
                                                                    case "foo bar":
                                                                        v4 = (v3 + v3)
                                                                    default:
                                                                        switch "a" {
                                                                        case "":
                                                                            v4 += (^(v4))
                                                                        case "\n\t":
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                fmt.Println(fi5, v2)
                                                                            case "":
                                                                                v1[(-(fi5))] = 65535
                                                                            default:
                                                                                v1[fi5] = 0
                                                                            }
                                                                        case "foo bar":
                                                                            var v91 uint32 = 4294967295
                                                                            _ = v91
                                                                        default:
                                                                            var v92 int32 = -2147483648
                                                                            var v93 float32 = -2.25
                                                                            if (!false) {
                                                                                switch 32768 {
                                                                                case -2147483648:
                                                                                    var v94 float64 = 0.0
                                                                                    _ = v94
                                                                                default:
                                                                                    v3 = (v3 << ((255) & 63))
                                                                                }
                                                                                v95 := GenS12{F0: v4, F1: 142}
                                                                                _ = v95
                                                                            } else {
                                                                                var v96 int32 = v92
                                                                                _ = v96
                                                                            }
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                    }
                                                                    v3 = (v3 * (v3 % (((^(v4))) | 1)))
                                                                    fmt.Println(v4, v3, v4)
                                                                    switch (("foo bar" + "a") + ("\n\t" + "foo bar")) {
                                                                    case "\n\t":
                                                                        if (42 <= 0) {
                                                                            v2 += v2
                                                                        }
                                                                    case "foo bar":
                                                                        if ((((("\n\t" + "") + "") + ("hello" + (("hello" + "hello") + ("a" + "")))) + "foo bar") == "\n\t") {
                                                                            var v97 int32 = -1
                                                                            v98 := GenS13{F0: 255}
                                                                            switch (("hello" + "hello") + (("" + (("" + "\n\t") + "foo bar")) + "a")) {
                                                                            case "foo bar":
                                                                                var v99 string = "\n\t"
                                                                                v99 = v99
                                                                                v100 := GenS3{F0: -21, F1: v97}
                                                                                _ = v100
                                                                                _ = v99
                                                                            }
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                    case "a":
                                                                        v1[fi5] = 2
                                                                    default:
                                                                        v101 := GenS14{F0: v3, F1: 0, F2: ((!(96 <= 0)) && false)}
                                                                        _ = v101
                                                                    }
                                                                    v102 := GenS11{F0: (v4 | (v3 % (((v4 & (v3 ^ (v4 / (((v3 + (v4 ^ 42))) | 1))))) | 1))), F1: 506}
                                                                    _ = v102
                                                                    _ = v88
                                                                }
                                                                v2 = (v2 % (((v2 ^ v2)) | 1))
                                                            }
                                                        }
                                                        v103 := GenS6{F0: 128, F1: "foo bar", F2: (v4 > v4)}
                                                        if (!true) {
                                                            v4 = 1
                                                        } else {
                                                            var v104 int64 = -1
                                                            _ = v104
                                                        }
                                                        var v105 string = ((("foo bar" + ("a" + (("\n\t" + ("\n\t" + "hello")) + "a"))) + "foo bar") + "foo bar")
                                                        var v106 float64 = -2.25
                                                        v4 += v3
                                                        _ = v103
                                                        _ = v105
                                                        _ = v106
                                                    }
                                                }
                                                switch (((("\n\t" + "\n\t") + ("foo bar" + "hello")) + "a") + "") {
                                                case "foo bar":
                                                    v4 = -1
                                                case "hello":
                                                    fmt.Println(v4)
                                                case "\n\t":
                                                    v2 = v2
                                                default:
                                                    v2 = (v2 >> ((255) & 63))
                                                }
                                            }
                                            switch "" {
                                            case "a":
                                                v107 := GenS7{F0: 0, F1: "a", F2: 127}
                                                _ = v107
                                            }
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(v4, v4)
                                            case "a":
                                                switch (fi5 ^ fi43) {
                                                case 1:
                                                    v4 += v4
                                                case -172:
                                                    var v108 uint64 = v2
                                                    _ = v108
                                                case -994:
                                                    switch "hello" {
                                                    case "a":
                                                        if false {
                                                            v2 += v2
                                                        } else {
                                                            v109 := []int{127, fi5, fi43, 538}
                                                            if (false && (443 <= 255)) {
                                                                v3 = (^(v4))
                                                                if true {
                                                                    _ = len(v109)
                                                                    v110 := GenS15{F0: 2, F1: 2147483648}
                                                                    _ = v110
                                                                } else {
                                                                    v3 += v3
                                                                    v111 := GenS9{F0: -2147483648, F1: 0}
                                                                    _ = v111
                                                                }
                                                                var v112 int8 = int8(fi5)
                                                                _ = v112
                                                            }
                                                            if (((288 != 0) && false) && ("" > "")) {
                                                                if true {
                                                                    v113 := GenS4{F0: -1, F1: -608, F2: 255}
                                                                    _ = v113
                                                                }
                                                            } else {
                                                                v114 := GenS16{F0: (^(v2)), F1: -730}
                                                                _ = v114
                                                            }
                                                            v115 := GenS15{F0: 4294967295, F1: 4294967295}
                                                            var v116 bool = true
                                                            if (255 == 42) {
                                                                var v117 int = fi43
                                                                fmt.Println(fi5, fi5)
                                                                _ = v117
                                                            }
                                                            var v118 uint16 = 42
                                                            _ = v109
                                                            _ = v115
                                                            _ = v116
                                                            _ = v118
                                                        }
                                                    default:
                                                        v4 = v4
                                                    }
                                                default:
                                                    switch ("\n\t" + ("" + "\n\t")) {
                                                    case "":
                                                        var v119 int64 = (-(v4))
                                                        _ = v119
                                                    case "hello":
                                                        delete(v1, fi5)
                                                    default:
                                                        switch ("\n\t" + "a") {
                                                        case "hello":
                                                            var v120 bool = (99 > 234)
                                                            _ = v120
                                                        default:
                                                            v1[(-(fi5))] = 65535
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                v121 := GenS3{F0: 885, F1: -2147483648}
                                                _ = v121
                                            default:
                                                switch "" {
                                                case "hello":
                                                    v2 += 4294967295
                                                case "a":
                                                    fmt.Println(v2, v3)
                                                default:
                                                    v122 := []int{(^(fi43)), fi5, fi43}
                                                    _ = v122
                                                }
                                            }
                                        default:
                                            if false {
                                                var v123 uint = 0
                                                _ = v123
                                            }
                                        }
                                        switch fi43 {
                                        case 0:
                                            var v124 uint32 = 0
                                            _ = v124
                                        default:
                                            _ = len(v1)
                                        }
                                    case "\n\t":
                                        var v125 bool = true
                                        _ = v125
                                    }
                                }
                                var v126 float64 = -2.25
                                _ = v126
                            }
                        }
                        _ = v60
                    }
                    v4 -= (-(v4))
                    _ = fi43
                }
            default:
                delete(v1, (^(fi5)))
            }
            _ = fi5
        }
        _ = v4
    } else {
        v127 := GenS17{F0: (^(v3)), F1: 42, F2: v2}
        _ = v127
    }
    v2 -= 1
    v128 := []int{42, -786}
    switch -2147483648 {
    case 2147483647:
        var v129 uint = 1
        _ = v129
    }
    fmt.Println(v3, v3)
    v130 := GenS18{F0: 448}
    for fi131 := 0; fi131 < 8; fi131++ {
        v2 -= v2
        _ = fi131
    }
    v3 -= v3
    fmt.Println(v3, v2)
    v2 = 42
    fmt.Println(v3, v3)
    v132 := GenS13{F0: 1}
    v2 = (v2 + v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v128:", len(v128), v128)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v128
    _ = v130
    _ = v132
    _ = v2
    _ = v3
}
