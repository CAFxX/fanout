package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int; F2 int64 }
type GenS7 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 int32; F1 uint8 }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 int; F1 uint64 }
type GenS14 struct { F0 uint; F1 bool }
type GenS15 struct { F0 int; F1 int32; F2 string }

func main() {
    var v1 int32 = -953
    v2 := []string{((("" + "foo bar") + "a") + "\n\t"), ("hello" + "\n\t"), "\n\t"}
    var v3 string = "foo bar"
    v4 := make(map[string]int)
    if (3.14159 < -2.25) {
        switch -2147483648 {
        case 218:
            switch v3 {
            case "hello":
                var v5 int8 = -1
                _ = v5
            case "foo bar":
                if false {
                    var v6 uint32 = 746
                    fmt.Println(v1, v6, v3)
                    _ = v6
                } else {
                    var v7 string = "\n\t"
                    fmt.Println(v1, v3, v7)
                    _ = v7
                }
            case "":
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 uint8 = 169
                    if false {
                        v9 -= v9
                        v3 = "a"
                    } else {
                        switch 983 {
                        case 9223372036854775807:
                            v9 += v9
                        }
                        if true {
                            v10 := GenS1{F0: 947, F1: -696, F2: 65535}
                            var v11 float32 = 0.0
                            for fi12 := 0; fi12 < 11; fi12++ {
                                v13 := GenS2{F0: 2, F1: (!(!true))}
                                v2 = append(v2, "hello")
                                _ = fi12
                                _ = v13
                            }
                            _ = v10
                            _ = v11
                        }
                    }
                    _ = fi8
                    _ = v9
                }
            }
        case 0:
            v14 := GenS2{F0: 544, F1: (127 == 128)}
            _ = v14
        }
        delete(v4, v3)
        _ = len(v4)
    } else {
        var v15 string = v3
        if ((-(v1)) == (v1 ^ v1)) {
            _ = len(v2)
        } else {
            for fi16 := 0; fi16 < 10; fi16++ {
                fmt.Println(v15, fi16, v1)
                if (float64((0.1 - 1.5)) > 1.5) {
                    fmt.Println(v15, fi16, v15)
                    if (-1 != 1) {
                        fmt.Println(v15, v3, v15)
                        _ = len(v2)
                    } else {
                        var v17 uint16 = 255
                        if false {
                            if true {
                                var v18 int64 = 750
                                _ = v18
                            }
                        }
                        _ = v17
                    }
                    var v19 uint16 = 1
                    _ = v19
                }
                _ = fi16
            }
            if (2 < 42) {
                var v20 float32 = -2.25
                for fi21 := 0; fi21 < 17; fi21++ {
                    switch 32767 {
                    case 127:
                        v22 := GenS3{F0: 2147483648, F1: 0}
                        _ = v22
                    case -1:
                        switch -2147483648 {
                        case 1:
                            _ = len(v2)
                        default:
                            var v23 float64 = ((0.1 / 1.5) / 3.14159)
                            _ = v23
                        }
                    default:
                        v1 = 1
                    }
                    fmt.Println(v1, v1, v20)
                    _ = fi21
                }
                _ = v20
            }
        }
        _ = v15
    }
    v24 := GenS3{F0: 120, F1: 94}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v25 int32 = (-(v1))
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 int = (-(fi26))
        for fi28 := 0; fi28 < 7; fi28++ {
            if (!(688 != 456)) {
                v3 = v3
                switch (fi28 / ((4294967295) | 1)) {
                case 128:
                    v29 := GenS2{F0: 4294967295, F1: (((65535 < 4294967295) || true) && false)}
                    _ = v29
                case 32768:
                    var v30 uint64 = 792
                    _ = v30
                }
            }
            v1 = (v25 & (-(v1)))
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    v31 := v4[v3]
    var v32 float64 = ((-2.25 * (((1.5 - 1.5) - 1.5) * float64(1.5))) - -2.25)
    v32 = (v32 * 0.0)
    var v33 uint = 18446744073709551615
    var v34 int32 = (v1 - v25)
    switch int(v25) {
    case -1:
        var v35 float32 = 1.5
        _ = v35
    case -619:
        v34 += 482
    default:
        _ = len(v2)
    }
    switch ((((v3 + "a") + v3) + (v3 + v3)) + v3) {
    case "hello":
        v25 = v25
    case "\n\t":
        v33 = v33
    }
    switch (v3 + "foo bar") {
    case "a":
        if (!(true && true)) {
            v1 = 1
        }
    case "hello":
        if (!(-1 >= 60)) {
            var v36 float64 = v32
            switch "" {
            case "hello":
                fmt.Println(v3)
            case "":
                var v37 string = "hello"
                _ = v37
            }
            fmt.Println(v1)
            _ = v36
        }
    case "foo bar":
        if (!(1 != 255)) {
            switch v3 {
            case "":
                var v38 uint8 = 243
                _ = v38
            }
            switch (v31 * v31) {
            case 667:
                v39 := GenS3{F0: 65535, F1: 1}
                _ = v39
            case -790:
                if (-1 != -32768) {
                    for fi40 := 0; fi40 < 8; fi40++ {
                        var v41 uint16 = 1
                        if ((65535 >= 42) || false) {
                            v42 := GenS2{F0: 783, F1: false}
                            for fi43 := 0; fi43 < 19; fi43++ {
                                v2 = append(v2, ("foo bar" + v3))
                                v25 -= (v1 << ((v33) & 63))
                                _ = fi43
                            }
                            _ = v42
                        } else {
                            switch "" {
                            case "foo bar":
                                v1 = (v1 - 1)
                            case "":
                                switch -9223372036854775808 {
                                case -759:
                                    for fi44 := 0; fi44 < 19; fi44++ {
                                        if (271 <= 2147483647) {
                                            var v45 uint = (v33 | 65535)
                                            var v46 int32 = v34
                                            if (1000 != 32767) {
                                                v46 -= (^(v46))
                                                v3 = "\n\t"
                                                var v47 int32 = v1
                                                _ = v47
                                            }
                                            if ((v41 & v41) == v41) {
                                                switch (fi44 * (fi40 >> (((v45 & 91)) & 63))) {
                                                case 127:
                                                    fmt.Println(v45, v1, v1)
                                                    if true {
                                                        if true {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        v48 := GenS4{F0: (v33 != (^(v45))), F1: 69, F2: 42}
                                                        _ = v48
                                                    }
                                                case -2147483648:
                                                    fmt.Println(v3, v45)
                                                default:
                                                    var v49 uint32 = 4294967295
                                                    _ = v49
                                                }
                                            }
                                            v25 = v34
                                            _ = v45
                                            _ = v46
                                        }
                                        _ = fi44
                                    }
                                }
                            case "hello":
                                v50 := GenS3{F0: 42, F1: 42}
                                _ = v50
                            default:
                                v2 = append(v2, "a")
                            }
                        }
                        _ = fi40
                        _ = v41
                    }
                    if (v3 > ("\n\t" + (v3 + v3))) {
                        v32 = v32
                        switch v31 {
                        case -2147483648:
                            switch (v3 + v3) {
                            case "":
                                if false {
                                    var v51 int16 = 1
                                    _ = v51
                                } else {
                                    switch (v31 | (v31 ^ 1)) {
                                    case 32767:
                                        v33 = v33
                                    case -9223372036854775808:
                                        _ = len(v4)
                                    case 314:
                                        switch 2147483647 {
                                        case -2147483648:
                                            v52 := GenS5{F0: 2}
                                            _ = v52
                                        }
                                    default:
                                        switch (v3 + (((v3 + v3) + (v3 + v3)) + "a")) {
                                        case "foo bar":
                                            fmt.Println(v31, v34)
                                        case "a":
                                            v1 = (v34 | v25)
                                        default:
                                            var v53 uint8 = 42
                                            _ = v53
                                        }
                                    }
                                    v54 := v4["foo bar"]
                                    _ = v54
                                }
                            case "a":
                                var v55 int16 = 1
                                _ = v55
                            }
                        case 472:
                            switch -2147483648 {
                            case 0:
                                v34 = (-(v34))
                            default:
                                for fi56 := 0; fi56 < 7; fi56++ {
                                    if true {
                                        switch v3 {
                                        case "hello":
                                            v33 -= 1
                                        case "foo bar":
                                            v57 := GenS1{F0: uint16(v33), F1: -9223372036854775808, F2: 647}
                                            _ = v57
                                        default:
                                            if true {
                                                if true {
                                                    fmt.Println(v34, v1, v33)
                                                }
                                                for fi58 := 0; fi58 < 11; fi58++ {
                                                    var v59 uint8 = 237
                                                    v59 = v59
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = fi58
                                                    _ = v59
                                                }
                                                if (!false) {
                                                    _ = len(v2)
                                                    v60 := GenS3{F0: 65535, F1: -128}
                                                    _ = v60
                                                } else {
                                                    v32 = v32
                                                    v2 = append(v2, ("hello" + "hello"))
                                                }
                                            }
                                        }
                                        var v61 int64 = 42
                                        v62 := GenS1{F0: uint16(v33), F1: v61, F2: 2}
                                        _ = v61
                                        _ = v62
                                    }
                                    var v63 uint64 = 4294967295
                                    _ = fi56
                                    _ = v63
                                }
                            }
                        case 1:
                            var v64 int32 = v1
                            _ = v64
                        default:
                            if false {
                                _ = len(v4)
                                var v65 string = v3
                                _ = v65
                            } else {
                                var v66 uint8 = 128
                                if (!true) {
                                    v33 = 2
                                }
                                v32 -= v32
                                _ = v66
                            }
                        }
                        if (!(!(v33 < (v33 % ((0) | 1))))) {
                            var v67 uint64 = 4294967295
                            v68 := GenS6{F0: (v67 | 978), F1: v31, F2: 0}
                            if ((-1 <= 1) || ((4294967295 != 539) && ((^(v25)) != v1))) {
                                var v69 uint16 = 32767
                                for fi70 := 0; fi70 < 4; fi70++ {
                                    delete(v4, "")
                                    _ = fi70
                                }
                                if true {
                                    for fi71 := 0; fi71 < 7; fi71++ {
                                        if (true && true) {
                                            v3 = v3
                                        }
                                        for fi72 := 0; fi72 < 4; fi72++ {
                                            var v73 int = fi71
                                            v1 -= -2147483648
                                            _ = fi72
                                            _ = v73
                                        }
                                        var v74 string = ("a" + v3)
                                        _ = fi71
                                        _ = v74
                                    }
                                    v75 := GenS7{F0: 2, F1: 42, F2: 416}
                                    _ = v75
                                } else {
                                    if false {
                                        v69 += v69
                                        v76 := GenS8{F0: 0, F1: 1}
                                        v67 += v67
                                        _ = v76
                                    } else {
                                        var v77 int = 4294967295
                                        for fi78 := 0; fi78 < 9; fi78++ {
                                            if false {
                                                v79 := GenS3{F0: 2147483647, F1: -128}
                                                if false {
                                                    for fi80 := 0; fi80 < 2; fi80++ {
                                                        if len(v2) > 0 {
                                                            v2[0] = (v3 + v3)
                                                        }
                                                        if (false && (0.0 == 0.0)) {
                                                            if (0 >= (v34 >> ((v33) & 63))) {
                                                                var v81 uint8 = 128
                                                                if true {
                                                                    switch (v3 + "a") {
                                                                    case "a":
                                                                        if (((0.1 + 3.14159) + 1.5) <= 3.14159) {
                                                                            switch (fi80 & 42) {
                                                                            case -612:
                                                                                var v82 int32 = 0
                                                                                _ = v82
                                                                            case 127:
                                                                                var v83 int64 = -9223372036854775808
                                                                                _ = v83
                                                                            case -1:
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    if (!false) {
                                                                                        var v84 int8 = 0
                                                                                        v85 := GenS1{F0: (^(v69)), F1: -429, F2: (v69 & v69)}
                                                                                        if true {
                                                                                            v77 -= fi78
                                                                                            var v86 uint32 = 2
                                                                                            var v87 int16 = 0
                                                                                            _ = v86
                                                                                            _ = v87
                                                                                        }
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    } else {
                                                                                        fmt.Println(v32)
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        if ((v32 - 0.0) > v32) {
                                                                                            if (!false) {
                                                                                                fmt.Println(v25)
                                                                                            } else {
                                                                                                _ = len(v4)
                                                                                                var v88 int64 = -1
                                                                                                var v89 uint16 = v69
                                                                                                _ = v88
                                                                                                _ = v89
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v2 = append(v2, (v3 + v3))
                                                                                    v2 = append(v2, v3)
                                                                                default:
                                                                                    var v90 uint16 = (v69 / (((v69 & (v69 * (v69 << (((^(v33))) & 63))))) | 1))
                                                                                    _ = v90
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    switch (v3 + v3) {
                                                                                    case "foo bar":
                                                                                        switch fi80 {
                                                                                        case 128:
                                                                                            if (true || (62 <= 1)) {
                                                                                                v69 = v69
                                                                                                switch (v3 + v3) {
                                                                                                case "hello":
                                                                                                    _ = len(v2)
                                                                                                default:
                                                                                                    if true {
                                                                                                        v91 := GenS9{F0: v81}
                                                                                                        _ = v91
                                                                                                    }
                                                                                                }
                                                                                                v92 := GenS2{F0: 424, F1: ((v31 | v77) < fi80)}
                                                                                                _ = v92
                                                                                            } else {
                                                                                                v93 := GenS7{F0: uint32(v67), F1: 65535, F2: -1}
                                                                                                if (2 <= 42) {
                                                                                                    v94 := GenS10{F0: (^(v67)), F1: 556}
                                                                                                    fmt.Println(v25, v67, v33)
                                                                                                    if false {
                                                                                                        if (222 < (v81 / (((v81 ^ v81)) | 1))) {
                                                                                                            v95 := GenS4{F0: (-1 == 664), F1: 4294967295, F2: v67}
                                                                                                            v32 += v32
                                                                                                            _ = v95
                                                                                                        }
                                                                                                        v96 := GenS4{F0: true, F1: 2147483648, F2: 2}
                                                                                                        _ = len(v4)
                                                                                                        _ = v96
                                                                                                    } else {
                                                                                                        v67 += v67
                                                                                                    }
                                                                                                    v31 = fi80
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    var v97 int16 = 1000
                                                                                                    v4[(v3 + ((v3 + "") + "foo bar"))] = fi78
                                                                                                    var v98 float64 = (v32 + 1e10)
                                                                                                    _ = v97
                                                                                                    _ = v98
                                                                                                }
                                                                                                _ = v93
                                                                                            }
                                                                                        case 1:
                                                                                            fmt.Println(v34, v31)
                                                                                        case -250:
                                                                                            var v99 int = 297
                                                                                            _ = v99
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        switch (v77 ^ -346) {
                                                                                        case 418:
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = "\n\t"
                                                                                            }
                                                                                            v32 = v32
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = "\n\t"
                                                                                            }
                                                                                            var v100 int16 = 249
                                                                                            _ = v100
                                                                                        }
                                                                                        switch v3 {
                                                                                        case "hello":
                                                                                            switch (fi78 / ((4294967295) | 1)) {
                                                                                            case 642:
                                                                                                switch fi80 {
                                                                                                case 128:
                                                                                                    v67 = (v67 | (v67 << (((v33 + (^(v33)))) & 63)))
                                                                                                    var v101 uint32 = 2147483647
                                                                                                    if (v81 < 42) {
                                                                                                        var v102 int = v77
                                                                                                        _ = v102
                                                                                                    }
                                                                                                    v1 -= (v1 / ((v34) | 1))
                                                                                                    var v103 string = ("a" + "")
                                                                                                    _ = v101
                                                                                                    _ = v103
                                                                                                case 808:
                                                                                                    fmt.Println(fi78)
                                                                                                default:
                                                                                                    var v104 string = v3
                                                                                                    _ = v104
                                                                                                }
                                                                                            case -830:
                                                                                                v32 -= (0.0 * v32)
                                                                                            default:
                                                                                                v32 -= float64(float32((((float64(((-2.25 - 1.5) + (0.0 * 0.1))) * v32) * ((v32 - v32) / 1e10)) + -2.25)))
                                                                                            }
                                                                                        case "foo bar":
                                                                                            var v105 string = v3
                                                                                            _ = v105
                                                                                        }
                                                                                        v34 = (-(v34))
                                                                                    } else {
                                                                                        fmt.Println(v81)
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        fmt.Println(fi80, v32)
                                                                                        if (false || false) {
                                                                                            if (!((v33 % ((v33) | 1)) < (v33 + (^(v33))))) {
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = v3
                                                                                                }
                                                                                            } else {
                                                                                                delete(v4, v3)
                                                                                                var v106 int = -2147483648
                                                                                                var v107 int16 = 42
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            }
                                                                                        }
                                                                                        v81 -= v81
                                                                                    } else {
                                                                                        v2 = append(v2, v3)
                                                                                        var v108 uint16 = (v69 << (((v33 % ((uint(v81)) | 1))) & 63))
                                                                                        _ = v108
                                                                                    }
                                                                                    v4[v3] = v77
                                                                                    v109 := GenS4{F0: (-2.25 >= float64(float32(3.14159))), F1: 550, F2: (v67 * (v67 % (((v67 % (((v67 * 303)) | 1))) | 1)))}
                                                                                    var v110 bool = (v1 != v25)
                                                                                    if (!v110) {
                                                                                        switch v77 {
                                                                                        case -2147483648:
                                                                                            switch (-(fi80)) {
                                                                                            case 881:
                                                                                                v67 += (v67 & 4294967295)
                                                                                            }
                                                                                            var v111 int8 = -128
                                                                                            _ = v111
                                                                                        case 9223372036854775807:
                                                                                            var v112 string = (v3 + v3)
                                                                                            _ = v112
                                                                                        }
                                                                                        var v113 int32 = v25
                                                                                        switch -830 {
                                                                                        case 42:
                                                                                            var v114 uint32 = 0
                                                                                            _ = v114
                                                                                        case 464:
                                                                                            var v115 int32 = v1
                                                                                            _ = v115
                                                                                        case 0:
                                                                                            var v116 float64 = ((v32 - (v32 / (3.14159 - 1.5))) / 1.0)
                                                                                            _ = v116
                                                                                        }
                                                                                        _ = v113
                                                                                    }
                                                                                    _ = len(v4)
                                                                                    _ = v109
                                                                                    _ = v110
                                                                                }
                                                                            }
                                                                            v117 := GenS4{F0: (false || false), F1: 4294967295, F2: 4294967295}
                                                                            if (-9223372036854775808 < -9223372036854775808) {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (((v3 + "foo bar") + ("a" + v3)) + (v3 + (("" + v3) + "\n\t")))
                                                                                }
                                                                                v118 := GenS11{F0: 2147483647, F1: uint8(v67)}
                                                                                v32 = (3.14159 + (v32 * float64(1e10)))
                                                                                v2 = append(v2, v3)
                                                                                _ = v118
                                                                            } else {
                                                                                v31 -= (fi80 ^ v77)
                                                                            }
                                                                            _ = v117
                                                                        }
                                                                    default:
                                                                        switch ((v3 + "") + v3) {
                                                                        case "foo bar":
                                                                            v81 = v81
                                                                        case "hello":
                                                                            switch (-(v77)) {
                                                                            case -2147483648:
                                                                                var v119 int16 = 42
                                                                                _ = v119
                                                                            default:
                                                                                fmt.Println(v34, v25)
                                                                            }
                                                                            delete(v4, ("a" + "hello"))
                                                                            v120 := GenS12{F0: 1, F1: (v77 + v31)}
                                                                            _ = v120
                                                                        case "\n\t":
                                                                            var v121 bool = true
                                                                            _ = v121
                                                                        default:
                                                                            if (!(((948 < v33) && (((1e10 - 1.5) * 1e10) == 0.1)) && true)) {
                                                                                if (v33 < 65535) {
                                                                                    v67 += 18446744073709551615
                                                                                    if (false && false) {
                                                                                        v122 := GenS12{F0: 32767, F1: fi78}
                                                                                        if true {
                                                                                            switch ((v3 + v3) + v3) {
                                                                                            case "\n\t":
                                                                                                var v123 bool = ((false && (!(v33 >= v33))) && (((!false) && (!true)) && ((v69 & v69) <= uint16(v33))))
                                                                                                _ = v123
                                                                                            default:
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                            }
                                                                                            v3 = (v3 + "foo bar")
                                                                                            if (42 != 0) {
                                                                                                if (-912 <= -1) {
                                                                                                    v2 = append(v2, (("\n\t" + (v3 + "foo bar")) + v3))
                                                                                                }
                                                                                            }
                                                                                            v124 := GenS1{F0: 2, F1: 1, F2: 63}
                                                                                            var v125 int8 = 42
                                                                                            _ = v124
                                                                                            _ = v125
                                                                                        }
                                                                                        var v126 bool = (v31 > (fi78 & (v77 & fi78)))
                                                                                        _ = v122
                                                                                        _ = v126
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v1 -= (^(v1))
                                                                    v25 += v34
                                                                }
                                                                v127 := GenS13{F0: v31, F1: 800}
                                                                _ = v127
                                                                _ = v81
                                                            } else {
                                                                fmt.Println(fi78)
                                                                fmt.Println(v31)
                                                            }
                                                        }
                                                        switch "foo bar" {
                                                        case "":
                                                            if (!(42 <= 86)) {
                                                                var v128 uint8 = 128
                                                                var v129 int8 = 127
                                                                _ = v128
                                                                _ = v129
                                                            }
                                                        }
                                                        _ = fi80
                                                    }
                                                } else {
                                                    v1 = (-(v1))
                                                    v130 := GenS14{F0: 42, F1: true}
                                                    var v131 bool = false
                                                    _ = v130
                                                    _ = v131
                                                }
                                                _ = v79
                                            } else {
                                                if false {
                                                    var v132 uint16 = (v69 | (^(v69)))
                                                    var v133 uint16 = (v132 / ((42) | 1))
                                                    var v134 string = ((v3 + "") + ("\n\t" + "hello"))
                                                    _ = v132
                                                    _ = v133
                                                    _ = v134
                                                } else {
                                                    fmt.Println(v77, v32, v32)
                                                    var v135 float64 = (v32 - v32)
                                                    fmt.Println(v34)
                                                    _ = v135
                                                }
                                            }
                                            switch v3 {
                                            case "":
                                                v3 = ("foo bar" + v3)
                                            case "\n\t":
                                                if false {
                                                    var v136 int16 = 241
                                                    _ = v136
                                                } else {
                                                    if true {
                                                        if (1000 <= 1000) {
                                                            v137 := GenS15{F0: fi78, F1: v34, F2: v3}
                                                            v69 -= 0
                                                            v2 = append(v2, "a")
                                                            _ = v137
                                                        } else {
                                                            if true {
                                                                v25 = -772
                                                            }
                                                        }
                                                        fmt.Println(v33, v67, v3)
                                                    }
                                                }
                                            case "foo bar":
                                                fmt.Println(fi78)
                                            default:
                                                var v138 uint16 = 32767
                                                _ = v138
                                            }
                                            _ = fi78
                                        }
                                        _ = v77
                                    }
                                }
                                _ = v69
                            }
                            _ = v67
                            _ = v68
                        }
                    } else {
                        var v139 int = (^(v31))
                        _ = v139
                    }
                    var v140 float64 = (v32 * v32)
                    _ = v140
                }
            }
        }
    }
    switch (v31 + v31) {
    case 42:
        v33 -= v33
    case 4294967295:
        var v141 int32 = (v25 >> ((v33) & 63))
        _ = v141
    default:
        v3 = v3
    }
    v142 := GenS8{F0: 2147483647, F1: 0}
    delete(v4, v3)
    v143 := GenS6{F0: 0, F1: v31, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v142
    _ = v143
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
