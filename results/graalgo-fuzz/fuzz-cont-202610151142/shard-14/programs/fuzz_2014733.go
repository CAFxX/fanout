package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 int64; F1 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint; F1 uint; F2 bool }

func main() {
    v1 := []uint32{65535, 1}
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    switch "a" {
    case "\n\t":
        var v5 int8 = 1
        _ = v5
    case "hello":
        var v6 int8 = 100
        _ = v6
    case "a":
        v7 := GenS1{F0: 42}
        _ = v7
    default:
        var v8 uint8 = 1
        _ = v8
    }
    switch "" {
    case "hello":
        var v9 int64 = -628
        _ = v9
    case "a":
        v10 := GenS1{F0: -1}
        _ = v10
    default:
        var v11 int32 = -195
        _ = v11
    }
    v12 := GenS1{F0: 662}
    var v13 uint = 2
    var v14 int64 = 1
    if true {
        fmt.Println(v14)
    } else {
        v15 := GenS2{F0: 105}
        _ = v15
    }
    if true {
        v16 := GenS3{F0: (v14 << ((v13) & 63))}
        _ = v16
    }
    fmt.Println(v14, v14)
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            if false {
                var v19 uint8 = 127
                fmt.Println(v14, fi17)
                _ = v19
            } else {
                v13 -= (v13 % (((v13 | 4294967295)) | 1))
                var v20 int16 = 0
                v21 := GenS1{F0: (v20 - (v20 - v20))}
                if (!true) {
                    var v22 uint64 = 4294967295
                    v20 -= v20
                    _ = v22
                }
                v2[2] = "foo bar"
                switch ("\n\t" + "\n\t") {
                case "a":
                    fmt.Println(v13, fi18)
                case "hello":
                    if false {
                        var v23 int = fi17
                        switch (v23 % ((fi17) | 1)) {
                        case -641:
                            v14 += v14
                        case 9223372036854775807:
                            switch 653 {
                            case 2147483647:
                                v13 = 2
                            case 4294967295:
                                fmt.Println(fi18)
                            case 128:
                                switch v23 {
                                case -1:
                                    var v24 uint64 = 2
                                    _ = v24
                                }
                            }
                        case 32767:
                            if (2147483647 > 2147483647) {
                                fmt.Println(fi18, fi18)
                                v20 += (v20 << ((v13) & 63))
                                v25 := GenS2{F0: 42}
                                v26 := GenS4{F0: -2147483648}
                                if (true || ((187 != 0) || (2 >= 18446744073709551615))) {
                                    switch (-(fi18)) {
                                    case -358:
                                        v14 = v14
                                    default:
                                        v20 -= int16(v14)
                                    }
                                    v20 = (v20 + v20)
                                    var v27 uint64 = 2
                                    _ = v27
                                }
                                _ = v25
                                _ = v26
                            } else {
                                var v28 float64 = 1.5
                                v28 = v28
                                v28 += float64((3.14159 * -2.25))
                                if (65535 > 255) {
                                    switch 127 {
                                    case 9223372036854775807:
                                        var v29 int = (fi17 * 605)
                                        _ = v29
                                    case 127:
                                        fmt.Println(v28)
                                    }
                                    _ = len(v1)
                                    switch fi17 {
                                    case 127:
                                        if (0 != -119) {
                                            _ = len(v1)
                                        } else {
                                            _ = len(v1)
                                            var v30 int32 = -573
                                            var v31 string = (((("hello" + "\n\t") + "a") + ("a" + "\n\t")) + "a")
                                            _ = v30
                                            _ = v31
                                        }
                                        v20 -= (^(v20))
                                        switch ("foo bar" + "\n\t") {
                                        case "a":
                                            var v32 float64 = 3.14159
                                            _ = v32
                                        case "foo bar":
                                            v14 -= (^(v14))
                                        }
                                        if true {
                                            var v33 int16 = 618
                                            v34 := GenS5{F0: v23}
                                            _ = v33
                                            _ = v34
                                        }
                                        fmt.Println(fi18, v28, v23)
                                    default:
                                        if true {
                                            var v35 int16 = (v20 / ((32767) | 1))
                                            _ = v35
                                        } else {
                                            var v36 float32 = 0.1
                                            var v37 int8 = 36
                                            fmt.Println(v14, v36)
                                            fmt.Println(v36, fi17, v23)
                                            _ = v36
                                            _ = v37
                                        }
                                    }
                                } else {
                                    var v38 uint = v13
                                    v39 := v4[0]
                                    v40 := v2[0]
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                                fmt.Println(v23)
                                _ = v28
                            }
                            var v41 float32 = 1e10
                            switch "" {
                            case "hello":
                                v20 = v20
                            case "a":
                                v42 := v2[42]
                                _ = v42
                            default:
                                v23 = (v23 ^ (^(v23)))
                            }
                            _ = v41
                        }
                        v23 = -9223372036854775808
                        _ = v23
                    } else {
                        switch fi17 {
                        case -1:
                            v43 := GenS2{F0: 1}
                            if (false && (65535 < 159)) {
                                switch (fi17 + -29) {
                                case -2147483648:
                                    var v44 int64 = (v14 * v14)
                                    switch "a" {
                                    case "\n\t":
                                        var v45 float32 = -2.25
                                        _ = v45
                                    case "foo bar":
                                        v44 = v44
                                    }
                                    if (!("hello" <= "\n\t")) {
                                        if (!("foo bar" > ("" + "\n\t"))) {
                                            if (46 >= 127) {
                                                v46 := v4[0]
                                                if true {
                                                    v46 = "\n\t"
                                                    if true {
                                                        if false {
                                                            var v47 float32 = 1.5
                                                            _ = v47
                                                        }
                                                        v14 = 1
                                                        var v48 int = 4294967295
                                                        _ = v48
                                                    }
                                                } else {
                                                    v20 += v20
                                                }
                                                _ = v46
                                            }
                                            if false {
                                                v14 -= (v14 >> (((^(v13))) & 63))
                                                var v49 int16 = v20
                                                _ = len(v4)
                                                _ = v49
                                            } else {
                                                delete(v2, 65535)
                                                if true {
                                                    switch 32768 {
                                                    case 4294967295:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case -304:
                                                        v50 := GenS6{F0: 429}
                                                        _ = v50
                                                    case -630:
                                                        if (480 <= -2147483648) {
                                                            var v51 uint = 773
                                                            v52 := GenS3{F0: v14}
                                                            var v53 int64 = -697
                                                            v20 -= (v20 - (v20 & (v20 % (((v20 | -1)) | 1))))
                                                            if (2147483647 > 0) {
                                                                v54 := GenS3{F0: v44}
                                                                v44 -= v14
                                                                _ = v54
                                                            }
                                                            _ = v51
                                                            _ = v52
                                                            _ = v53
                                                        } else {
                                                            if false {
                                                                v20 += v20
                                                                v20 += 1000
                                                            } else {
                                                                v20 = (v20 + (v20 ^ (v20 * (v20 + v20))))
                                                                v55 := GenS5{F0: (fi18 << ((v13) & 63))}
                                                                _ = v55
                                                            }
                                                            v56 := GenS5{F0: fi18}
                                                            var v57 int8 = int8(v13)
                                                            if true {
                                                                var v58 uint64 = 42
                                                                _ = v58
                                                            } else {
                                                                v13 = 652
                                                                var v59 int8 = v57
                                                                if ((v13 + v13) > v13) {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v59 = v59
                                                                    switch (fi18 / (((fi17 | (fi18 - (fi18 >> ((v13) & 63))))) | 1)) {
                                                                    case 42:
                                                                        switch 1 {
                                                                        case 127:
                                                                            var v60 uint = v13
                                                                            v13 = v13
                                                                            v59 = -1
                                                                            if false {
                                                                                if (!(!(((!(((((0.1 + 1.5) * float32(float64(3.14159))) * (1.5 * 0.0)) >= 1.5) && true)) && false) || ((v44 / (((v14 / ((v44) | 1))) | 1)) == v14)))) {
                                                                                    v13 = (v13 * (v60 * (v60 % ((244) | 1))))
                                                                                }
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 2147483648
                                                                                }
                                                                            }
                                                                            var v61 uint16 = 255
                                                                            v20 = (^(v20))
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            _ = v60
                                                                            _ = v61
                                                                        default:
                                                                            var v62 uint16 = 0
                                                                            _ = v62
                                                                        }
                                                                    case 127:
                                                                        switch fi17 {
                                                                        case -409:
                                                                            v44 = -1
                                                                        case -732:
                                                                            switch fi18 {
                                                                            case 1:
                                                                                if false {
                                                                                    v20 = v20
                                                                                    _ = len(v2)
                                                                                }
                                                                            case 42:
                                                                                var v63 int32 = 630
                                                                                _ = v63
                                                                            case 4294967295:
                                                                                v64 := GenS7{F0: 42}
                                                                                _ = v64
                                                                            }
                                                                        case 2147483647:
                                                                            v14 = -1
                                                                        default:
                                                                            v14 += v44
                                                                        }
                                                                    case 2147483647:
                                                                        if ((v57 & (v59 * 42)) < v57) {
                                                                            v44 = v14
                                                                            var v65 int8 = v57
                                                                            _ = v65
                                                                        }
                                                                        v57 += (v59 * (v59 * v57))
                                                                        if (!(!false)) {
                                                                            if (v20 == v20) {
                                                                                if (580 >= 2147483647) {
                                                                                    if ((3.14159 / 3.14159) >= 0.0) {
                                                                                        if true {
                                                                                            if false {
                                                                                                v44 -= v44
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v66 int32 = 0
                                                                                                    v13 -= (v13 & 18446744073709551615)
                                                                                                    var v67 int = 85
                                                                                                    _ = v66
                                                                                                    _ = v67
                                                                                                } else {
                                                                                                    v68 := v4[765]
                                                                                                    if (!(true || (-2147483648 == -1))) {
                                                                                                        v69 := GenS3{F0: (v44 & 42)}
                                                                                                        _ = v69
                                                                                                    }
                                                                                                    _ = v68
                                                                                                }
                                                                                                if (v20 > 32767) {
                                                                                                    v70 := GenS3{F0: (v14 + (v14 >> (((v13 - (v13 + v13))) & 63)))}
                                                                                                    _ = v70
                                                                                                } else {
                                                                                                    v13 = v13
                                                                                                    switch int(v20) {
                                                                                                    case 9223372036854775807:
                                                                                                        var v71 bool = (!true)
                                                                                                        _ = v71
                                                                                                    default:
                                                                                                        fmt.Println(v14)
                                                                                                        fmt.Println(v14, v20, v14)
                                                                                                        v72 := v2[0]
                                                                                                        _ = v72
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    v20 = (v20 ^ 953)
                                                                                                    v1 = append(v1, 1)
                                                                                                    if true {
                                                                                                        var v73 string = ""
                                                                                                        _ = v73
                                                                                                    } else {
                                                                                                        if ((2 >= 864) && (-991 < -1)) {
                                                                                                            v74 := GenS8{F0: (^(v13)), F1: 1}
                                                                                                            fmt.Println(fi17, v44, v14)
                                                                                                            _ = v74
                                                                                                        }
                                                                                                    }
                                                                                                    switch -268 {
                                                                                                    case 1:
                                                                                                        v13 -= v13
                                                                                                    case 9223372036854775807:
                                                                                                        v59 = v57
                                                                                                    case 4294967295:
                                                                                                        v1 = append(v1, 2147483647)
                                                                                                    }
                                                                                                } else {
                                                                                                    switch (fi18 + fi18) {
                                                                                                    case 443:
                                                                                                        if (0.0 == float64(float32(float64(0.0)))) {
                                                                                                            var v75 int16 = v20
                                                                                                            v57 = -34
                                                                                                            _ = v75
                                                                                                        }
                                                                                                    case 9223372036854775807:
                                                                                                        v76 := v4[4294967295]
                                                                                                        _ = v76
                                                                                                    }
                                                                                                    fmt.Println(v57)
                                                                                                    fmt.Println(v14)
                                                                                                }
                                                                                                switch ("\n\t" + "\n\t") {
                                                                                                case "\n\t":
                                                                                                    if (((1e10 / 3.14159) + 3.14159) <= 1e10) {
                                                                                                        var v77 uint = 921
                                                                                                        delete(v4, 2)
                                                                                                        _ = v77
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    fmt.Println(fi17, v14, v14)
                                                                                                }
                                                                                                var v78 uint32 = 729
                                                                                                _ = v78
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v79 uint = 138
                                                                                        var v80 bool = false
                                                                                        switch fi17 {
                                                                                        case 127:
                                                                                            if true {
                                                                                                var v81 int32 = int32(fi17)
                                                                                                _ = v81
                                                                                            }
                                                                                        case -982:
                                                                                            var v82 uint8 = 63
                                                                                            _ = v82
                                                                                        }
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    }
                                                                                    var v83 int16 = -489
                                                                                    fmt.Println(v57, v83)
                                                                                    _ = v83
                                                                                } else {
                                                                                    if (1e10 > 0.1) {
                                                                                        switch "\n\t" {
                                                                                        case "foo bar":
                                                                                            var v84 uint64 = 2
                                                                                            _ = v84
                                                                                        }
                                                                                    }
                                                                                    v13 = 376
                                                                                    var v85 string = "\n\t"
                                                                                    _ = v85
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    v14 -= (v44 ^ v14)
                                                                                    if false {
                                                                                        var v86 uint16 = 32767
                                                                                        fmt.Println(fi17, fi18)
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        _ = v86
                                                                                    } else {
                                                                                        switch fi18 {
                                                                                        case 0:
                                                                                            fmt.Println(v13, v14, v14)
                                                                                        default:
                                                                                            v44 -= (v14 + (-(v14)))
                                                                                        }
                                                                                        var v87 uint16 = 801
                                                                                        var v88 float64 = 0.0
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                    }
                                                                                }
                                                                                fmt.Println(v59)
                                                                                switch (^(fi17)) {
                                                                                case 1:
                                                                                    v14 += (^(v14))
                                                                                    fmt.Println(v44, v20)
                                                                                }
                                                                            }
                                                                            v89 := GenS2{F0: -73}
                                                                            fmt.Println(fi18, v44, fi18)
                                                                            _ = v89
                                                                        } else {
                                                                            v90 := v2[586]
                                                                            v20 = v20
                                                                            _ = v90
                                                                        }
                                                                        v13 += 255
                                                                        switch ("\n\t" + "\n\t") {
                                                                        case "\n\t":
                                                                            delete(v2, 1)
                                                                        default:
                                                                            fmt.Println(v59, v44, v57)
                                                                        }
                                                                        v91 := GenS5{F0: fi17}
                                                                        if false {
                                                                            var v92 float32 = (1e10 + 1.5)
                                                                            var v93 int8 = v57
                                                                            switch 127 {
                                                                            case -89:
                                                                                v3[(fi17 & fi18)] = (1 >= 42)
                                                                            case 402:
                                                                                v94 := GenS9{F0: (v14 ^ -629), F1: 1}
                                                                                _ = v94
                                                                            default:
                                                                                switch fi18 {
                                                                                case 2147483647:
                                                                                    v95 := GenS7{F0: 2}
                                                                                    _ = v95
                                                                                }
                                                                            }
                                                                            _ = v92
                                                                            _ = v93
                                                                        } else {
                                                                            if (32767 != 831) {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                fmt.Println(fi17)
                                                                                var v96 int32 = -169
                                                                                _ = v96
                                                                            }
                                                                            v14 += int64(fi18)
                                                                            v44 -= (^(v14))
                                                                            v44 = 567
                                                                        }
                                                                        _ = v91
                                                                    default:
                                                                        v97 := GenS10{F0: (!true)}
                                                                        _ = v97
                                                                    }
                                                                }
                                                                _ = v59
                                                            }
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    default:
                                                        if (fi17 != (^(fi17))) {
                                                            var v98 uint64 = 18446744073709551615
                                                            var v99 int16 = (v20 ^ v20)
                                                            fmt.Println(v14)
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                    }
                                                    v100 := GenS11{F0: fi18, F1: ""}
                                                    _ = v100
                                                }
                                            }
                                            var v101 int = fi18
                                            _ = v101
                                        }
                                        _ = len(v4)
                                        if false {
                                            fmt.Println(v14)
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(fi17, v14)
                                        }
                                    }
                                    _ = v44
                                case 167:
                                    if true {
                                        delete(v4, 246)
                                        if false {
                                            v20 = -32768
                                        }
                                    }
                                case 911:
                                    switch 0 {
                                    case 42:
                                        var v102 uint32 = 4294967295
                                        _ = v102
                                    }
                                default:
                                    v103 := GenS1{F0: (v20 - v20)}
                                    _ = v103
                                }
                                v13 = (^(v13))
                            }
                            _ = v43
                        case 261:
                            switch (("foo bar" + "") + ("a" + "\n\t")) {
                            case "foo bar":
                                v104 := GenS12{F0: v13}
                                _ = v104
                            case "\n\t":
                                delete(v2, 2147483647)
                            default:
                                var v105 uint = (v13 | 4294967295)
                                _ = v105
                            }
                        }
                        if ((!(fi17 > 682)) && (true || (0 >= 18446744073709551615))) {
                            v106 := GenS6{F0: uint32(v13)}
                            _ = v106
                        } else {
                            switch ("hello" + "foo bar") {
                            case "hello":
                                var v107 bool = (v20 <= (v20 / ((951) | 1)))
                                _ = v107
                            case "":
                                if (4294967295 == 1) {
                                    var v108 float64 = 3.14159
                                    fmt.Println(fi18)
                                    _ = v108
                                }
                            case "a":
                                v13 = 233
                            }
                            fmt.Println(v14, v13)
                        }
                        _ = len(v1)
                        var v109 int16 = 1000
                        var v110 float32 = 0.0
                        _ = v109
                        _ = v110
                    }
                case "\n\t":
                    v13 = v13
                }
                v14 += v14
                _ = v20
                _ = v21
            }
            v14 = (v14 ^ -1)
            var v111 uint16 = 42
            fmt.Println(v14)
            _ = fi18
            _ = v111
        }
        v13 -= 0
        v13 = 255
        _ = fi17
    }
    fmt.Println(v14, v13)
    var v112 bool = true
    var v113 uint64 = 4294967295
    v14 = (v14 % (((v14 % ((-805) | 1))) | 1))
    switch int(v113) {
    case -1:
        for fi114 := 0; fi114 < 1; fi114++ {
            if true {
                v14 -= (v14 / ((0) | 1))
                v14 = v14
                v14 -= v14
            } else {
                var v115 int16 = 1000
                _ = v115
            }
            _ = fi114
        }
    }
    switch "a" {
    case "\n\t":
        if (true || (!v112)) {
            for fi116 := 0; fi116 < 8; fi116++ {
                var v117 bool = false
                var v118 int64 = (^(v14))
                _ = fi116
                _ = v117
                _ = v118
            }
            var v119 int64 = -1
            _ = v119
        } else {
            var v120 int64 = 1
            v113 = v113
            if (true || v112) {
                var v121 float32 = 1e10
                _ = v121
            }
            _ = v120
        }
    case "foo bar":
        var v122 int32 = -30
        _ = v122
    case "":
        if ((-2.25 + 3.14159) != ((3.14159 + 0.1) + 0.0)) {
            for fi123 := 0; fi123 < 19; fi123++ {
                if false {
                    fmt.Println(fi123)
                } else {
                    var v124 string = ""
                    _ = v124
                }
                switch fi123 {
                case -338:
                    var v125 uint = 1
                    _ = v125
                case 2147483647:
                    if ((((1e10 - float32(1e10)) + (float32(1e10) + 0.0)) == 3.14159) && true) {
                        for fi126 := 0; fi126 < 11; fi126++ {
                            var v127 int8 = 42
                            switch (-(fi123)) {
                            case -523:
                                var v128 int16 = 42
                                _ = v128
                            default:
                                v129 := GenS13{F0: (v14 | v14)}
                                if (v112 && v112) {
                                    fmt.Println(v112)
                                    var v130 string = "hello"
                                    v13 = (v13 << ((v13) & 63))
                                    _ = v130
                                } else {
                                    _ = len(v1)
                                    if v112 {
                                        v131 := GenS11{F0: (fi126 * fi126), F1: ("\n\t" + "")}
                                        v127 = (v127 & (v127 | v127))
                                        _ = v131
                                    }
                                    switch "foo bar" {
                                    case "":
                                        fmt.Println(v14, v14)
                                    case "hello":
                                        v132 := GenS14{F0: 793}
                                        _ = v132
                                    case "a":
                                        var v133 uint16 = 65535
                                        _ = v133
                                    }
                                    v14 += (v14 << ((v13) & 63))
                                }
                                v134 := GenS15{F0: (v13 + v13), F1: v13, F2: ((v14 >= v14) || (609 > 1))}
                                _ = v129
                                _ = v134
                            }
                            var v135 int = (fi123 - fi123)
                            _ = fi126
                            _ = v127
                            _ = v135
                        }
                    }
                case 953:
                    v14 -= (v14 % ((0) | 1))
                }
                if v112 {
                    v136 := GenS7{F0: v113}
                    _ = v136
                } else {
                    v113 -= 18446744073709551615
                    v13 = v13
                }
                _ = fi123
            }
            delete(v3, 9223372036854775807)
            v137 := GenS2{F0: 42}
            _ = v137
        } else {
            var v138 bool = v112
            _ = v138
        }
    }
    delete(v2, 4294967295)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v112
    _ = v113
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
