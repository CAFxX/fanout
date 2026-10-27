package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 uint64 }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 string; F1 uint16; F2 uint64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint32; F1 bool }
type GenS14 struct { F0 bool; F1 int16; F2 uint8 }
type GenS15 struct { F0 int; F1 uint16 }
type GenS16 struct { F0 int16; F1 int64 }
type GenS17 struct { F0 uint64; F1 uint32; F2 uint }
type GenS18 struct { F0 int32; F1 uint64 }
type GenS19 struct { F0 int64; F1 uint64; F2 uint }
type GenS20 struct { F0 bool; F1 int64; F2 int16 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint8; F1 int; F2 uint16 }
type GenS23 struct { F0 string; F1 int8 }
type GenS24 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS25 struct { F0 uint32 }

func main() {
    var v1 float64 = -2.25
    var v2 int = 9223372036854775807
    var v3 string = "hello"
    var v4 int8 = -51
    var v5 int64 = -355
    v6 := make(map[int]bool)
    v1 = (((v1 - v1) / 3.14159) + v1)
    v4 -= v4
    v1 += v1
    fmt.Println(v1)
    v7 := GenS1{F0: 2, F1: v2, F2: 65535}
    switch (v2 ^ (v2 & (v2 + (v2 + v2)))) {
    case -232:
        var v8 string = v3
        _ = v8
    case -391:
        if (true || (true || true)) {
            if (!(42 < 0)) {
                var v9 int32 = -2147483648
                _ = v9
            } else {
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 int64 = (v5 & (v5 - 44))
                    _ = fi10
                    _ = v11
                }
                v5 += v5
                v6[4294967295] = true
            }
            v4 = (^(v4))
        }
    case 32767:
        switch (v2 << ((255) & 63)) {
        case 32767:
            if (false || (v3 > v3)) {
                if (false || (!true)) {
                    if true {
                        var v12 int64 = (v5 + v5)
                        v4 = v4
                        _ = v12
                    }
                    var v13 int32 = 1
                    var v14 float32 = (0.1 * (1e10 - (-2.25 / (3.14159 / 1.0))))
                    _ = v13
                    _ = v14
                }
                switch (v2 | v2) {
                case -9223372036854775808:
                    v15 := GenS2{F0: (v5 >= (v5 | v5)), F1: 185, F2: 0}
                    _ = v15
                case 2147483647:
                    switch (v2 - int(v4)) {
                    case 127:
                        var v16 int = (v2 & v2)
                        _ = v16
                    case 32768:
                        if ((4294967295 < 113) && ((65535 >= 2) && true)) {
                            if (289 == -608) {
                                for fi17 := 0; fi17 < 15; fi17++ {
                                    var v18 int64 = (v5 / ((v5) | 1))
                                    var v19 int64 = (v5 - v18)
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                                v5 -= (v5 % (((^(v5))) | 1))
                            } else {
                                if false {
                                    for fi20 := 0; fi20 < 3; fi20++ {
                                        var v21 uint64 = 0
                                        _ = fi20
                                        _ = v21
                                    }
                                    v2 = (v2 << ((2) & 63))
                                } else {
                                    for fi22 := 0; fi22 < 18; fi22++ {
                                        switch v2 {
                                        case 32767:
                                            _ = len(v6)
                                        case 128:
                                            v23 := GenS3{F0: false, F1: 1}
                                            _ = v23
                                        default:
                                            for fi24 := 0; fi24 < 6; fi24++ {
                                                if ((((float64(float32(float64((0.1 + 1e10)))) - (float64(float32((v1 * (v1 * ((3.14159 / ((v1 * v1) + float64(((1e10 - 1.5) - ((-2.25 * float32(1.5)) * 3.14159))))) * float64(3.14159)))))) + v1)) / ((float64(-2.25) * 1.5) * (v1 - v1))) / v1) < (v1 + (v1 - 0.1))) {
                                                    switch v3 {
                                                    case "hello":
                                                        var v25 uint32 = 353
                                                        _ = v25
                                                    case "":
                                                        var v26 uint64 = 0
                                                        _ = v26
                                                    case "a":
                                                        v27 := []int{v2, fi22}
                                                        _ = v27
                                                    }
                                                    switch ("hello" + "hello") {
                                                    case "a":
                                                        delete(v6, (fi22 << ((0) & 63)))
                                                    default:
                                                        fmt.Println(fi22)
                                                    }
                                                }
                                                var v28 int = fi22
                                                var v29 int = int(v5)
                                                _ = fi24
                                                _ = v28
                                                _ = v29
                                            }
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            v30 := GenS1{F0: 0, F1: (fi22 | v2), F2: 42}
                                            _ = v30
                                        case "\n\t":
                                            if true {
                                                fmt.Println(fi22)
                                                v2 = (fi22 * fi22)
                                            } else {
                                                if (v5 > (^(v5))) {
                                                    v1 -= (v1 + v1)
                                                    v31 := v6[v2]
                                                    _ = v31
                                                } else {
                                                    if true {
                                                        fmt.Println(v4, v1, v4)
                                                        switch ((("hello" + v3) + "a") + (v3 + ("a" + ("a" + "hello")))) {
                                                        case "a":
                                                            for fi32 := 0; fi32 < 11; fi32++ {
                                                                v33 := []int{-9223372036854775808, (v2 % ((v2) | 1)), (-(fi22)), fi32}
                                                                var v34 uint16 = 0
                                                                _ = fi32
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                        case "\n\t":
                                                            for fi35 := 0; fi35 < 7; fi35++ {
                                                                v4 = 31
                                                                var v36 uint16 = 32767
                                                                _ = fi35
                                                                _ = v36
                                                            }
                                                        case "":
                                                            switch v2 {
                                                            case -426:
                                                                switch (v2 - 42) {
                                                                case 0:
                                                                    if ((((914 <= 1) || (!(1e10 > 1.5))) && false) || (!(0.1 > 0.0))) {
                                                                        if (!(611 >= 42)) {
                                                                            var v37 int = (v2 % ((v2) | 1))
                                                                            if (32767 >= -1) {
                                                                                v5 = v5
                                                                                if false {
                                                                                    v5 = (v5 / (((v5 % ((v5) | 1))) | 1))
                                                                                }
                                                                                for fi38 := 0; fi38 < 12; fi38++ {
                                                                                    var v39 uint8 = 0
                                                                                    v6[v2] = (!(!false))
                                                                                    _ = fi38
                                                                                    _ = v39
                                                                                }
                                                                            } else {
                                                                                var v40 float64 = -2.25
                                                                                _ = v40
                                                                            }
                                                                            v41 := GenS4{F0: 1, F1: 1, F2: 128}
                                                                            _ = v37
                                                                            _ = v41
                                                                        } else {
                                                                            v5 = 101
                                                                            v42 := GenS1{F0: 2, F1: v2, F2: 1}
                                                                            _ = v42
                                                                        }
                                                                        v43 := []int{(v2 ^ 1), v2, fi22, (fi22 >> ((42) & 63))}
                                                                        switch v2 {
                                                                        case -740:
                                                                            fmt.Println(v4)
                                                                        default:
                                                                            if true {
                                                                                var v44 float64 = 0.1
                                                                                var v45 float64 = 0.1
                                                                                v44 = -2.25
                                                                                _ = v44
                                                                                _ = v45
                                                                            } else {
                                                                                v46 := GenS2{F0: true, F1: 34, F2: 4294967295}
                                                                                var v47 int16 = 42
                                                                                _ = v46
                                                                                _ = v47
                                                                            }
                                                                        }
                                                                        _ = v43
                                                                    } else {
                                                                        v48 := GenS5{F0: v3}
                                                                        var v49 string = v3
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                }
                                                            case 0:
                                                                if (!true) {
                                                                    var v50 float32 = -2.25
                                                                    if true {
                                                                        v5 += -9223372036854775808
                                                                        for fi51 := 0; fi51 < 12; fi51++ {
                                                                            var v52 int16 = 1
                                                                            fmt.Println(v1, v3, v4)
                                                                            v53 := GenS4{F0: 128, F1: 0, F2: 1}
                                                                            _ = fi51
                                                                            _ = v52
                                                                            _ = v53
                                                                        }
                                                                        v54 := []int{v2}
                                                                        _ = v54
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(v5, v5, v3)
                                                                            for fi55 := 0; fi55 < 20; fi55++ {
                                                                                v50 = v50
                                                                                var v56 int64 = (v5 / ((v5) | 1))
                                                                                v57 := GenS6{F0: 2, F1: 42, F2: 42}
                                                                                _ = fi55
                                                                                _ = v56
                                                                                _ = v57
                                                                            }
                                                                            for fi58 := 0; fi58 < 15; fi58++ {
                                                                                if (v3 >= v3) {
                                                                                    var v59 uint8 = 245
                                                                                    var v60 uint64 = uint64(v59)
                                                                                    v61 := v6[(fi22 ^ 32767)]
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                } else {
                                                                                    if true {
                                                                                        if (((!((^(v5)) >= (v5 % (((v5 ^ v5)) | 1)))) && true) && (!((0.0 - (v1 + 1.5)) >= (v1 + (1.5 / v1))))) {
                                                                                            switch (fi58 << ((586) & 63)) {
                                                                                            case -2147483648:
                                                                                                if true {
                                                                                                    fmt.Println(v2, fi58, v50)
                                                                                                } else {
                                                                                                    v62 := []int{fi58, (fi58 / ((v2) | 1)), 560, fi58}
                                                                                                    _ = v62
                                                                                                }
                                                                                            default:
                                                                                                var v63 uint8 = 128
                                                                                                _ = v63
                                                                                            }
                                                                                            var v64 int16 = 0
                                                                                            _ = v64
                                                                                        }
                                                                                    }
                                                                                    v65 := v6[v2]
                                                                                    _ = v65
                                                                                }
                                                                                v66 := GenS3{F0: ((461 > 4294967295) && (-2147483648 < -1)), F1: 747}
                                                                                _ = fi58
                                                                                _ = v66
                                                                            }
                                                                        } else {
                                                                            if ((false || (65535 >= 65535)) || false) {
                                                                                var v67 uint8 = 42
                                                                                v68 := []int{(v2 >> ((255) & 63)), 1, (fi22 / (((^(v2))) | 1)), (v2 / ((v2) | 1))}
                                                                                var v69 uint16 = 853
                                                                                _ = v67
                                                                                _ = v68
                                                                                _ = v69
                                                                            }
                                                                            for fi70 := 0; fi70 < 18; fi70++ {
                                                                                v71 := GenS4{F0: 0, F1: 2, F2: 200}
                                                                                _ = fi70
                                                                                _ = v71
                                                                            }
                                                                            v1 = ((1.5 / float64(((0.0 - 3.14159) + (1e10 / v50)))) + 1e10)
                                                                        }
                                                                    }
                                                                    v72 := []int{(^(fi22))}
                                                                    _ = v50
                                                                    _ = v72
                                                                }
                                                            }
                                                        default:
                                                            v73 := []int{}
                                                            _ = v73
                                                        }
                                                        for fi74 := 0; fi74 < 1; fi74++ {
                                                            delete(v6, (fi74 & (^(v2))))
                                                            if (true || true) {
                                                                if (640 != 2) {
                                                                    delete(v6, v2)
                                                                    if false {
                                                                        fmt.Println(v1, v1)
                                                                        var v75 int64 = -330
                                                                        _ = v75
                                                                    } else {
                                                                        fmt.Println(v4)
                                                                        fmt.Println(v4, v4)
                                                                    }
                                                                    var v76 float64 = (-2.25 + (1.5 - ((0.0 + v1) + (float64(float32(v1)) - float64(0.1)))))
                                                                    _ = v76
                                                                } else {
                                                                    var v77 float64 = v1
                                                                    if false {
                                                                        if (!true) {
                                                                            var v78 string = "foo bar"
                                                                            _ = v78
                                                                        } else {
                                                                            var v79 int = 42
                                                                            _ = v79
                                                                        }
                                                                        v1 += float64(1.5)
                                                                    }
                                                                    var v80 string = "foo bar"
                                                                    _ = v77
                                                                    _ = v80
                                                                }
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if (!true) {
                                                                        delete(v6, v2)
                                                                    } else {
                                                                        switch ((v3 + "\n\t") + "foo bar") {
                                                                        case "\n\t":
                                                                            var v81 uint = 255
                                                                            _ = v81
                                                                        default:
                                                                            var v82 uint32 = 42
                                                                            _ = v82
                                                                        }
                                                                        if ((0 == 18446744073709551615) || (!(0 > 2147483647))) {
                                                                            v83 := GenS1{F0: 2, F1: -9223372036854775808, F2: 0}
                                                                            var v84 uint8 = 255
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            v6[(v2 % ((fi74) | 1))] = (v3 > v3)
                                                                        }
                                                                    }
                                                                case "":
                                                                    var v85 string = (((v3 + "") + v3) + v3)
                                                                    _ = v85
                                                                case "a":
                                                                    var v86 int = fi22
                                                                    _ = v86
                                                                default:
                                                                    var v87 float32 = 0.0
                                                                    _ = v87
                                                                }
                                                                var v88 int8 = int8(v5)
                                                                v89 := v6[fi74]
                                                                if v89 {
                                                                    v89 = (42 >= v2)
                                                                    v90 := GenS6{F0: 131, F1: 122, F2: 42}
                                                                    var v91 bool = true
                                                                    if false {
                                                                        if (42 <= 65535) {
                                                                            var v92 uint16 = 651
                                                                            v93 := []int{}
                                                                            v1 = v1
                                                                            _ = v92
                                                                            _ = v93
                                                                        }
                                                                        if v89 {
                                                                            if (1000 == int16(v2)) {
                                                                                v89 = v89
                                                                                var v94 int = v2
                                                                                switch v3 {
                                                                                case "hello":
                                                                                    var v95 uint8 = 2
                                                                                    _ = v95
                                                                                default:
                                                                                    v1 = 3.14159
                                                                                }
                                                                                v88 -= v88
                                                                                _ = v94
                                                                            }
                                                                            if v91 {
                                                                                if (65535 > 255) {
                                                                                    v96 := []int{}
                                                                                    var v97 int8 = v88
                                                                                    v3 = "foo bar"
                                                                                    var v98 float32 = float32(float64(0.0))
                                                                                    if (v89 && ((1.5 / 0.1) < 1.5)) {
                                                                                        var v99 uint32 = 2147483647
                                                                                        v2 = -849
                                                                                        _ = v99
                                                                                    } else {
                                                                                        fmt.Println(v91, v97)
                                                                                        if (2 > 678) {
                                                                                            if v89 {
                                                                                                var v100 int64 = -1
                                                                                                _ = v100
                                                                                            }
                                                                                            v101 := GenS7{F0: 868}
                                                                                            _ = v101
                                                                                        }
                                                                                        if (43 != (fi74 >> ((1) & 63))) {
                                                                                            var v102 float32 = v98
                                                                                            if v89 {
                                                                                                switch v2 {
                                                                                                case -176:
                                                                                                    v88 = v4
                                                                                                case -612:
                                                                                                    if v89 {
                                                                                                        if true {
                                                                                                            switch (v3 + v3) {
                                                                                                            case "\n\t":
                                                                                                                if true {
                                                                                                                    switch ("\n\t" + v3) {
                                                                                                                    case "hello":
                                                                                                                        v2 += 591
                                                                                                                        switch v3 {
                                                                                                                        case "hello":
                                                                                                                            switch (fi74 - (^(fi74))) {
                                                                                                                            case 125:
                                                                                                                                if (2 != 17) {
                                                                                                                                    if (1 == 1) {
                                                                                                                                        switch fi22 {
                                                                                                                                        case 127:
                                                                                                                                            switch (v2 + (v2 / (((fi22 & v2)) | 1))) {
                                                                                                                                            case -2147483648:
                                                                                                                                                v1 = float64((0.0 + v98))
                                                                                                                                                v3 = ("" + (v3 + v3))
                                                                                                                                                v103 := GenS7{F0: 453}
                                                                                                                                                _ = v103
                                                                                                                                            case 128:
                                                                                                                                                switch (v3 + v3) {
                                                                                                                                                case "\n\t":
                                                                                                                                                    _ = len(v6)
                                                                                                                                                default:
                                                                                                                                                    v5 -= v5
                                                                                                                                                }
                                                                                                                                                v96 = append(v96, fi22)
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            var v104 uint64 = 532
                                                                                                                                            _ = v104
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = len(v6)
                                                                                                                                } else {
                                                                                                                                    var v105 uint64 = 42
                                                                                                                                    v106 := GenS2{F0: v91, F1: 42, F2: 255}
                                                                                                                                    _ = v105
                                                                                                                                    _ = v106
                                                                                                                                }
                                                                                                                            case -1:
                                                                                                                                _ = len(v6)
                                                                                                                            }
                                                                                                                        case "foo bar":
                                                                                                                            var v107 int32 = -109
                                                                                                                            _ = v107
                                                                                                                        case "\n\t":
                                                                                                                            _ = len(v6)
                                                                                                                        default:
                                                                                                                            if len(v96) >= 2 {
                                                                                                                                v96 = v96[1:len(v96)]
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v1 -= v1
                                                                                                                    case "foo bar":
                                                                                                                        v108 := GenS8{F0: (fi74 - (fi74 + (-(fi22))))}
                                                                                                                        _ = v108
                                                                                                                    case "\n\t":
                                                                                                                        switch v3 {
                                                                                                                        case "\n\t":
                                                                                                                            var v109 uint16 = 32767
                                                                                                                            _ = v109
                                                                                                                        default:
                                                                                                                            v89 = true
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch fi74 {
                                                                                                                        case 4294967295:
                                                                                                                            v110 := GenS9{F0: 4294967295, F1: 549}
                                                                                                                            _ = v110
                                                                                                                        case 1:
                                                                                                                            _ = len(v6)
                                                                                                                        default:
                                                                                                                            if v89 {
                                                                                                                                if v89 {
                                                                                                                                    var v111 float64 = v1
                                                                                                                                    v112 := GenS5{F0: ("hello" + v3)}
                                                                                                                                    _ = v111
                                                                                                                                    _ = v112
                                                                                                                                }
                                                                                                                                v113 := GenS6{F0: 551, F1: 545, F2: 1}
                                                                                                                                _ = v113
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v2 += 128
                                                                                                                    fmt.Println(v102)
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v98 += v98
                                                                                                            default:
                                                                                                                var v114 float64 = v1
                                                                                                                _ = v114
                                                                                                            }
                                                                                                            var v115 int = v2
                                                                                                            if (((890 < 25) && ((1 != 656) && true)) && (v1 > (v1 * (float64(v98) - float64(v98))))) {
                                                                                                                switch (v2 | v2) {
                                                                                                                case -850:
                                                                                                                    var v116 uint = 316
                                                                                                                    _ = v116
                                                                                                                case 1:
                                                                                                                    v117 := GenS5{F0: (v3 + (v3 + v3))}
                                                                                                                    _ = v117
                                                                                                                default:
                                                                                                                    v97 = v4
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(fi22, v97)
                                                                                                            }
                                                                                                            var v118 int32 = 951
                                                                                                            _ = v115
                                                                                                            _ = v118
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    switch v2 {
                                                                                                    case 0:
                                                                                                        v119 := GenS4{F0: 42, F1: 253, F2: 2}
                                                                                                        _ = v119
                                                                                                    case 42:
                                                                                                        var v120 bool = (!false)
                                                                                                        _ = v120
                                                                                                    case 470:
                                                                                                        v121 := GenS7{F0: 1000}
                                                                                                        _ = v121
                                                                                                    default:
                                                                                                        var v122 int = fi74
                                                                                                        _ = v122
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v3)
                                                                                            _ = v102
                                                                                        }
                                                                                        v96 = append(v96, 42)
                                                                                        var v123 int8 = -1
                                                                                        v124 := GenS10{F0: v3, F1: -1}
                                                                                        _ = v123
                                                                                        _ = v124
                                                                                    }
                                                                                    _ = v96
                                                                                    _ = v97
                                                                                    _ = v98
                                                                                } else {
                                                                                    v125 := []int{fi22}
                                                                                    var v126 int = v2
                                                                                    v5 = v5
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                }
                                                                                switch v3 {
                                                                                case "":
                                                                                    var v127 string = (((v3 + "\n\t") + v3) + v3)
                                                                                    _ = v127
                                                                                case "foo bar":
                                                                                    if v91 {
                                                                                        fmt.Println(v3)
                                                                                        v4 -= (^(v4))
                                                                                        switch "\n\t" {
                                                                                        case "a":
                                                                                            v91 = (false || true)
                                                                                        case "hello":
                                                                                            v91 = (!true)
                                                                                        case "":
                                                                                            if v89 {
                                                                                                switch v3 {
                                                                                                case "a":
                                                                                                    var v128 float64 = v1
                                                                                                    _ = v128
                                                                                                }
                                                                                                var v129 int16 = 32767
                                                                                                _ = v129
                                                                                            } else {
                                                                                                v91 = v91
                                                                                                delete(v6, 1)
                                                                                                v3 = v3
                                                                                            }
                                                                                        default:
                                                                                            switch -584 {
                                                                                            case -574:
                                                                                                switch 9223372036854775807 {
                                                                                                case 870:
                                                                                                    v130 := GenS11{F0: (v3 + v3), F1: 197, F2: 300}
                                                                                                    _ = v130
                                                                                                default:
                                                                                                    if (v5 >= v5) {
                                                                                                        var v131 float64 = float64(0.1)
                                                                                                        var v132 int = 717
                                                                                                        _ = v131
                                                                                                        _ = v132
                                                                                                    }
                                                                                                }
                                                                                            case 527:
                                                                                                switch (fi74 ^ 9223372036854775807) {
                                                                                                case -1:
                                                                                                    v133 := GenS12{F0: 545}
                                                                                                    _ = v133
                                                                                                case 32767:
                                                                                                    var v134 uint64 = 2
                                                                                                    _ = v134
                                                                                                }
                                                                                            default:
                                                                                                v1 = float64(float32(v1))
                                                                                            }
                                                                                        }
                                                                                        var v135 string = v3
                                                                                        _ = v135
                                                                                    } else {
                                                                                        v5 = 0
                                                                                    }
                                                                                case "a":
                                                                                    v3 = ((v3 + ((v3 + v3) + "")) + (v3 + v3))
                                                                                default:
                                                                                    var v136 bool = v91
                                                                                    _ = v136
                                                                                }
                                                                                var v137 float64 = ((v1 / (v1 * v1)) - (v1 - v1))
                                                                                _ = v137
                                                                            }
                                                                        }
                                                                        var v138 int32 = 400
                                                                        _ = v138
                                                                    }
                                                                    fmt.Println(fi22)
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                                v139 := GenS13{F0: 2, F1: v89}
                                                                _ = v139
                                                                _ = v88
                                                                _ = v89
                                                            } else {
                                                                fmt.Println(fi74)
                                                                fmt.Println(v5)
                                                                var v140 uint32 = 681
                                                                _ = v140
                                                            }
                                                            v141 := GenS14{F0: false, F1: 1, F2: 233}
                                                            _ = fi74
                                                            _ = v141
                                                        }
                                                    }
                                                    for fi142 := 0; fi142 < 13; fi142++ {
                                                        if false {
                                                            if true {
                                                                if false {
                                                                    _ = len(v6)
                                                                    if false {
                                                                        if (false && false) {
                                                                            var v143 int8 = -128
                                                                            if (255 >= 127) {
                                                                                switch (fi22 >> ((0) & 63)) {
                                                                                case 448:
                                                                                    var v144 uint64 = 42
                                                                                    _ = v144
                                                                                case 32767:
                                                                                    v145 := []int{(fi142 & (fi22 << ((4294967295) & 63)))}
                                                                                    _ = v145
                                                                                case -681:
                                                                                    var v146 int16 = -893
                                                                                    _ = v146
                                                                                default:
                                                                                    v143 -= (v4 / (((-(v4))) | 1))
                                                                                }
                                                                            }
                                                                            switch 488 {
                                                                            case 128:
                                                                                v2 = (-(v2))
                                                                            case -510:
                                                                                v3 = (v3 + "foo bar")
                                                                            case 32767:
                                                                                var v147 uint32 = 522
                                                                                _ = v147
                                                                            }
                                                                            if ((!(18446744073709551615 != 2)) || (!false)) {
                                                                                v3 = ""
                                                                            } else {
                                                                                v148 := []int{(fi142 >> ((0) & 63)), (^(fi22))}
                                                                                _ = v148
                                                                            }
                                                                            v143 -= v143
                                                                            _ = v143
                                                                        } else {
                                                                            switch (fi142 | (v2 * fi22)) {
                                                                            case 9223372036854775807:
                                                                                var v149 int8 = (v4 + 42)
                                                                                _ = v149
                                                                            case 127:
                                                                                v2 = (fi22 * (-(fi22)))
                                                                            case 128:
                                                                                var v150 uint16 = 255
                                                                                v5 = v5
                                                                                if (!(!false)) {
                                                                                    v151 := GenS15{F0: (v2 / ((-286) | 1)), F1: 42}
                                                                                    var v152 uint32 = 838
                                                                                    switch (fi142 * (^(v2))) {
                                                                                    case -441:
                                                                                        v6[v2] = (18446744073709551615 == 550)
                                                                                    case -732:
                                                                                        var v153 int8 = v4
                                                                                        _ = v153
                                                                                    }
                                                                                    _ = v151
                                                                                    _ = v152
                                                                                } else {
                                                                                    if (true && (v4 > -74)) {
                                                                                        switch 128 {
                                                                                        case 707:
                                                                                            if (fi142 > 376) {
                                                                                                if true {
                                                                                                    v150 += 350
                                                                                                    v154 := []int{(fi142 + fi142), (^(v2))}
                                                                                                    var v155 uint16 = v150
                                                                                                    _ = v154
                                                                                                    _ = v155
                                                                                                }
                                                                                            }
                                                                                            switch v3 {
                                                                                            case "foo bar":
                                                                                                switch (fi22 | v2) {
                                                                                                case 1:
                                                                                                    v156 := v6[(fi22 / ((fi142) | 1))]
                                                                                                    _ = v156
                                                                                                case 127:
                                                                                                    v157 := GenS16{F0: 42, F1: int64(fi22)}
                                                                                                    _ = v157
                                                                                                default:
                                                                                                    v158 := GenS9{F0: 1, F1: 359}
                                                                                                    _ = v158
                                                                                                }
                                                                                            case "hello":
                                                                                                switch 0 {
                                                                                                case 42:
                                                                                                    v5 -= v5
                                                                                                case -109:
                                                                                                    v5 = v5
                                                                                                default:
                                                                                                    switch (v3 + "hello") {
                                                                                                    case "":
                                                                                                        v159 := GenS17{F0: 42, F1: 552, F2: 42}
                                                                                                        _ = v159
                                                                                                    case "\n\t":
                                                                                                        v160 := []int{fi22}
                                                                                                        var v161 int = fi142
                                                                                                        _ = v160
                                                                                                        _ = v161
                                                                                                    case "foo bar":
                                                                                                        v162 := []int{-9223372036854775808, fi22}
                                                                                                        _ = v162
                                                                                                    }
                                                                                                }
                                                                                            case "a":
                                                                                                v163 := GenS6{F0: v150, F1: (v150 / (((v150 + v150)) | 1)), F2: 312}
                                                                                                _ = v163
                                                                                            }
                                                                                        default:
                                                                                            var v164 int16 = -798
                                                                                            _ = v164
                                                                                        }
                                                                                    }
                                                                                    delete(v6, v2)
                                                                                }
                                                                                _ = v150
                                                                            default:
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    var v165 uint64 = 18446744073709551615
                                                                                    _ = v165
                                                                                }
                                                                            }
                                                                        }
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            fmt.Println(fi142, v2, v5)
                                                                        case "":
                                                                            v2 = 128
                                                                        case "foo bar":
                                                                            var v166 uint32 = 1
                                                                            _ = v166
                                                                        default:
                                                                            v1 = ((1e10 * (1.5 - v1)) * v1)
                                                                        }
                                                                    } else {
                                                                        switch (^(v2)) {
                                                                        case 1:
                                                                            v1 = 3.14159
                                                                        case -1:
                                                                            v167 := GenS9{F0: 694, F1: 516}
                                                                            _ = v167
                                                                        default:
                                                                            if (1 > -1) {
                                                                                var v168 uint8 = 2
                                                                                _ = v168
                                                                            }
                                                                        }
                                                                        var v169 float32 = ((0.0 * 0.0) + (((-2.25 + (0.1 * -2.25)) * 1e10) / 3.14159))
                                                                        _ = v169
                                                                    }
                                                                    var v170 float64 = v1
                                                                    _ = v170
                                                                }
                                                                v2 = (v2 | (v2 - (v2 & (fi22 >> ((954) & 63)))))
                                                            }
                                                            var v171 int = (-(v2))
                                                            if ((-767 != (fi22 * -1)) && (fi142 <= v171)) {
                                                                _ = len(v6)
                                                            } else {
                                                                var v172 int = fi142
                                                                var v173 uint16 = 0
                                                                var v174 int32 = 1
                                                                if (int(v4) <= fi142) {
                                                                    switch v172 {
                                                                    case 42:
                                                                        if false {
                                                                            v173 += v173
                                                                            var v175 uint16 = (v173 + (v173 | v173))
                                                                            _ = v175
                                                                        } else {
                                                                            if (!((v4 >> ((2) & 63)) < v4)) {
                                                                                v176 := []int{}
                                                                                var v177 uint64 = 42
                                                                                _ = v176
                                                                                _ = v177
                                                                            } else {
                                                                                v5 = (v5 - (-(v5)))
                                                                                var v178 int16 = 1
                                                                                if true {
                                                                                    v178 += v178
                                                                                    fmt.Println(v178)
                                                                                    v179 := GenS18{F0: v174, F1: 483}
                                                                                    v178 = -1
                                                                                    _ = v179
                                                                                }
                                                                                _ = v178
                                                                            }
                                                                            if false {
                                                                                if (true && (v4 == (v4 & (v4 << ((65535) & 63))))) {
                                                                                    v3 = ""
                                                                                }
                                                                            } else {
                                                                                v180 := GenS19{F0: -439, F1: 823, F2: 800}
                                                                                var v181 int16 = int16(v174)
                                                                                _ = v180
                                                                                _ = v181
                                                                            }
                                                                        }
                                                                    default:
                                                                        fmt.Println(v173, v3, v171)
                                                                    }
                                                                    if false {
                                                                        v182 := GenS20{F0: true, F1: (-(v5)), F2: 1000}
                                                                        _ = v182
                                                                    } else {
                                                                        v183 := GenS15{F0: v171, F1: (v173 ^ v173)}
                                                                        v174 = -1
                                                                        _ = v183
                                                                    }
                                                                    if false {
                                                                        v184 := GenS12{F0: 65535}
                                                                        _ = v184
                                                                    } else {
                                                                        if (!true) {
                                                                            var v185 uint = 807
                                                                            fmt.Println(v1, v1, fi142)
                                                                            _ = v185
                                                                        }
                                                                    }
                                                                } else {
                                                                    if (false || false) {
                                                                        var v186 string = "\n\t"
                                                                        switch v186 {
                                                                        case "hello":
                                                                            v171 = 9223372036854775807
                                                                        case "a":
                                                                            switch (v2 - 214) {
                                                                            case 732:
                                                                                v6[(-(fi22))] = true
                                                                            case 2147483647:
                                                                                v187 := GenS18{F0: (v174 % (((v174 | -1)) | 1)), F1: uint64(v173)}
                                                                                _ = v187
                                                                            default:
                                                                                if (!true) {
                                                                                    v171 -= (-(fi142))
                                                                                    v173 -= v173
                                                                                    if (v4 == v4) {
                                                                                        if false {
                                                                                            var v188 float32 = (-2.25 * (3.14159 / 3.14159))
                                                                                            _ = v188
                                                                                        } else {
                                                                                            v189 := []int{}
                                                                                            if false {
                                                                                                v190 := GenS21{F0: v4}
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    v191 := GenS15{F0: (v171 | (fi142 * (fi142 + fi22))), F1: v173}
                                                                                                    _ = v191
                                                                                                case "a":
                                                                                                    v192 := GenS1{F0: 0, F1: v171, F2: v173}
                                                                                                    _ = v192
                                                                                                case "":
                                                                                                    if len(v189) >= 2 {
                                                                                                        v189 = v189[1:len(v189)]
                                                                                                    }
                                                                                                }
                                                                                                _ = v190
                                                                                            }
                                                                                            var v193 float32 = (0.1 - float32((v1 * (v1 / (1e10 + v1)))))
                                                                                            fmt.Println(v174, v193)
                                                                                            _ = v189
                                                                                            _ = v193
                                                                                        }
                                                                                        var v194 uint = 255
                                                                                        v195 := GenS6{F0: (v173 << (((^(v194))) & 63)), F1: (v173 / ((v173) | 1)), F2: 1}
                                                                                        _ = v194
                                                                                        _ = v195
                                                                                    }
                                                                                } else {
                                                                                    v196 := GenS5{F0: ("hello" + (v186 + v186))}
                                                                                    _ = v196
                                                                                }
                                                                                switch (v186 + (((("\n\t" + "\n\t") + (v3 + v186)) + v186) + (((((v3 + (v3 + v3)) + v186) + "\n\t") + (v186 + (v3 + v186))) + ("a" + v3)))) {
                                                                                case "hello":
                                                                                    _ = len(v6)
                                                                                case "foo bar":
                                                                                    var v197 uint = 0
                                                                                    _ = v197
                                                                                case "a":
                                                                                    v171 = (fi22 % ((v171) | 1))
                                                                                default:
                                                                                    v3 = v186
                                                                                    fmt.Println(v2, v172, v173)
                                                                                }
                                                                                v198 := []int{1, -1, 128, v172}
                                                                                _ = v198
                                                                            }
                                                                        case "":
                                                                            v199 := GenS22{F0: 250, F1: v171, F2: v173}
                                                                            _ = v199
                                                                        }
                                                                        v3 = v186
                                                                        _ = v186
                                                                    }
                                                                    v200 := []int{fi142, -2147483648}
                                                                    v201 := GenS16{F0: 733, F1: 371}
                                                                    _ = v200
                                                                    _ = v201
                                                                }
                                                                switch (^(fi142)) {
                                                                case -823:
                                                                    var v202 uint8 = 255
                                                                    _ = v202
                                                                default:
                                                                    var v203 int8 = (v4 | (v4 >> ((255) & 63)))
                                                                    _ = v203
                                                                }
                                                                _ = v172
                                                                _ = v173
                                                                _ = v174
                                                            }
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v204 uint = 1
                                                                _ = v204
                                                            case "a":
                                                                switch v3 {
                                                                case "":
                                                                    var v205 string = v3
                                                                    _ = v205
                                                                }
                                                            }
                                                            if false {
                                                                if (!false) {
                                                                    v5 = (v5 & -9223372036854775808)
                                                                    v1 = ((1e10 + float64(float32((float64(float32(-2.25)) * v1)))) / v1)
                                                                }
                                                            }
                                                            _ = v171
                                                        } else {
                                                            v3 = "\n\t"
                                                            var v206 int8 = v4
                                                            switch v3 {
                                                            case "a":
                                                                if (128 != 226) {
                                                                    v207 := GenS14{F0: false, F1: 163, F2: 0}
                                                                    v208 := GenS23{F0: v3, F1: 32}
                                                                    var v209 uint = 178
                                                                    v1 = (v1 - ((float64(-2.25) + float64((3.14159 + -2.25))) / (v1 / (-2.25 / 3.14159))))
                                                                    _ = v207
                                                                    _ = v208
                                                                    _ = v209
                                                                }
                                                            case "hello":
                                                                v210 := v6[(v2 * v2)]
                                                                _ = v210
                                                            }
                                                            v211 := v6[290]
                                                            _ = v206
                                                            _ = v211
                                                        }
                                                        var v212 uint8 = 2
                                                        var v213 int = (v2 | fi22)
                                                        _ = fi142
                                                        _ = v212
                                                        _ = v213
                                                    }
                                                }
                                                fmt.Println(v1)
                                            }
                                        }
                                        _ = fi22
                                    }
                                }
                                v214 := GenS20{F0: false, F1: (v5 & (v5 + -1)), F2: -1}
                                _ = v214
                            }
                            _ = len(v6)
                        } else {
                            v1 = (v1 + (float64(float32((((v1 / 1e10) * v1) / float64(1.5)))) + (1e10 + 1.5)))
                        }
                    default:
                        if (!(0.0 > (0.0 - 3.14159))) {
                            fmt.Println(v1, v3, v2)
                            v215 := v6[(v2 * (-(v2)))]
                            switch int(v4) {
                            case 0:
                                var v216 int8 = v4
                                _ = v216
                            case -1:
                                var v217 int = 32767
                                _ = v217
                            }
                            _ = v215
                        } else {
                            v4 = 75
                        }
                    }
                default:
                    v6[(^(v2))] = (true && (0 > 717))
                }
            }
        case -933:
            v3 = v3
        case 9223372036854775807:
            if (-229 >= -395) {
                v3 = "\n\t"
                v218 := GenS24{F0: 328, F1: 1, F2: 123}
                _ = v218
            } else {
                for fi219 := 0; fi219 < 13; fi219++ {
                    for fi220 := 0; fi220 < 8; fi220++ {
                        if ((((3.14159 + v1) * float64(3.14159)) * 1e10) != v1) {
                            fmt.Println(v3)
                        } else {
                            var v221 float64 = (v1 + v1)
                            _ = v221
                        }
                        switch ("hello" + "foo bar") {
                        case "hello":
                            _ = len(v6)
                        case "foo bar":
                            if (557 < 32767) {
                                v5 = v5
                            }
                        case "\n\t":
                            v3 = (v3 + v3)
                        default:
                            _ = len(v6)
                        }
                        v1 -= (0.1 * v1)
                        _ = fi220
                    }
                    _ = fi219
                }
            }
        default:
            delete(v6, -2147483648)
        }
    }
    v1 -= 1.5
    fmt.Println(v4, v5, v5)
    var v222 int16 = 977
    fmt.Println(v4, v4)
    v223 := GenS25{F0: 4294967295}
    var v224 uint16 = 2
    v225 := GenS23{F0: (((v3 + v3) + "") + v3), F1: int8(v2)}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v222:", v222)
    fmt.Println("v224:", v224)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v222
    _ = v223
    _ = v224
    _ = v225
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
