package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int; F1 int; F2 bool }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 int16; F1 int16; F2 uint }
type GenS10 struct { F0 int8; F1 int8; F2 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32; F1 int64; F2 int8 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint64; F1 bool }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int32; F1 string }
type GenS19 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    var v3 float64 = 3.14159
    v2 = v2
    v1[32767] = 456
    v2 = (!v2)
    var v4 uint64 = 953
    v2 = v2
    v2 = true
    if false {
        switch 564 {
        case 749:
            delete(v1, -135)
        case 127:
            var v5 uint = 42
            _ = v5
        }
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 int64 = 0
            _ = fi6
            _ = v7
        }
    } else {
        v8 := GenS1{F0: -529, F1: "a"}
        fmt.Println(v2, v2, v3)
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if v2 {
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS1{F0: 42, F1: "hello"}
                _ = fi10
                _ = v11
            }
            var v12 uint16 = 1
            v13 := GenS2{F0: (fi9 == fi9), F1: 415, F2: v12}
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            if v2 {
                var v15 uint = 65535
                v16 := v1[2147483647]
                _ = v15
                _ = v16
            }
            if (42 <= uint8(v4)) {
                var v17 bool = (1 <= 128)
                _ = v17
            }
            _ = fi14
        }
        v3 = 1e10
        _ = fi9
    }
    v18 := GenS3{F0: 2147483647, F1: 70, F2: v2}
    v1[-120] = 475
    v3 = v3
    v19 := GenS4{F0: -9223372036854775808, F1: (2 >= 0)}
    fmt.Println(v3, v2, v4)
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := []int{127}
        _ = fi20
        _ = v21
    }
    if (1000 == -767) {
        for fi22 := 0; fi22 < 15; fi22++ {
            if (-31 > 0) {
                v2 = v2
            } else {
                var v23 uint8 = 2
                var v24 string = ("foo bar" + ((("\n\t" + ("a" + "foo bar")) + "a") + ("\n\t" + ("" + ("hello" + "hello")))))
                _ = v23
                _ = v24
            }
            v4 = v4
            _ = fi22
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := GenS3{F0: (fi25 - (fi25 % (((fi25 | fi25)) | 1))), F1: -745, F2: (9223372036854775807 < -9223372036854775808)}
            switch 4294967295 {
            case 9223372036854775807:
                v4 = (v4 | 4294967295)
            case -9223372036854775808:
                for fi27 := 0; fi27 < 10; fi27++ {
                    fmt.Println(v3, fi27)
                    _ = fi27
                }
            case 575:
                var v28 int32 = 0
                _ = v28
            default:
                for fi29 := 0; fi29 < 5; fi29++ {
                    fmt.Println(v4)
                    if v2 {
                        switch "\n\t" {
                        case "a":
                            switch ("" + "") {
                            case "a":
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(fi29, v4)
                                case "\n\t":
                                    var v30 uint64 = v4
                                    _ = v30
                                case "":
                                    switch fi25 {
                                    case -278:
                                        switch "a" {
                                        case "hello":
                                            if v2 {
                                                if ((v4 <= v4) || (-2147483648 == 385)) {
                                                    var v31 int32 = 0
                                                    v32 := v1[(fi29 << ((1) & 63))]
                                                    _ = v31
                                                    _ = v32
                                                }
                                                var v33 uint8 = 95
                                                if v2 {
                                                    v34 := GenS5{F0: -2147483648}
                                                    if false {
                                                        v33 -= 63
                                                        v4 += v4
                                                        if (!(1000 == 1)) {
                                                            v4 += (v4 / (((^(v4))) | 1))
                                                        }
                                                    }
                                                    var v35 uint32 = 2147483647
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    if (2 >= 18446744073709551615) {
                                                        var v36 uint8 = 127
                                                        var v37 uint = 986
                                                        var v38 string = (("\n\t" + "foo bar") + "hello")
                                                        var v39 int = (fi29 % ((int(v4)) | 1))
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                    var v40 uint64 = (v4 << ((255) & 63))
                                                    fmt.Println(v2)
                                                    _ = v40
                                                }
                                                v41 := GenS3{F0: fi25, F1: 32767, F2: ((v4 > 1) || ((false || v2) || (!(("a" <= "foo bar") || (!(255 >= 0))))))}
                                                v2 = (("foo bar" + ((((("\n\t" + "\n\t") + ("\n\t" + "a")) + ("" + "")) + ("a" + "")) + ("foo bar" + ("a" + "hello")))) > "foo bar")
                                                v2 = v2
                                                v33 += 255
                                                var v42 string = ""
                                                v43 := GenS6{F0: 127}
                                                _ = v33
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                            } else {
                                                if (v2 && false) {
                                                    if (-2.25 != 1e10) {
                                                        v44 := []int{-1, (fi29 - (fi25 / (((fi25 + (fi25 | fi25))) | 1))), fi29}
                                                        _ = v44
                                                    } else {
                                                        v45 := []int{(-(fi29)), fi25, (fi25 - (fi25 / ((42) | 1))), -319}
                                                        if (110 > 739) {
                                                            var v46 uint64 = v4
                                                            _ = v46
                                                        } else {
                                                            v4 -= 983
                                                            v4 -= v4
                                                            v2 = (100 > 42)
                                                        }
                                                        if len(v45) >= 2 {
                                                            v45 = v45[1:len(v45)]
                                                        }
                                                        if v2 {
                                                            var v47 int = (^(fi29))
                                                            switch (("hello" + "a") + ("a" + "\n\t")) {
                                                            case "hello":
                                                                if v2 {
                                                                    switch (fi25 << ((913) & 63)) {
                                                                    case 189:
                                                                        v4 = 18446744073709551615
                                                                    case -2147483648:
                                                                        v3 = (float64(((0.0 / -2.25) * (0.1 / -2.25))) - v3)
                                                                    case 0:
                                                                        v47 = (fi25 & v47)
                                                                    }
                                                                    delete(v1, fi25)
                                                                } else {
                                                                    var v48 uint16 = 0
                                                                    _ = v48
                                                                }
                                                            case "\n\t":
                                                                v3 = ((((1.5 / float64(float32(-2.25))) / v3) + ((-2.25 * (3.14159 - (3.14159 * (v3 + v3)))) * 3.14159)) / v3)
                                                            case "foo bar":
                                                                if false {
                                                                    if false {
                                                                        v49 := GenS7{F0: 0, F1: 894}
                                                                        v3 -= 0.0
                                                                        _ = v49
                                                                    }
                                                                    if v2 {
                                                                        switch (-(fi29)) {
                                                                        case -531:
                                                                            v50 := GenS1{F0: -593, F1: ("" + ("" + ""))}
                                                                            _ = v50
                                                                        case 32767:
                                                                            v47 = 2147483647
                                                                        }
                                                                    }
                                                                } else {
                                                                    if len(v45) >= 2 {
                                                                        v45 = v45[1:len(v45)]
                                                                    }
                                                                }
                                                            }
                                                            v51 := GenS6{F0: -89}
                                                            _ = v47
                                                            _ = v51
                                                        } else {
                                                            switch (fi29 ^ fi29) {
                                                            case 32768:
                                                                v4 = v4
                                                            case -234:
                                                                v52 := GenS4{F0: (fi25 << ((489) & 63)), F1: false}
                                                                _ = v52
                                                            case -658:
                                                                var v53 int = (fi25 ^ fi25)
                                                                _ = v53
                                                            }
                                                        }
                                                        _ = v45
                                                    }
                                                } else {
                                                    v2 = ("hello" == "foo bar")
                                                    v2 = (-1 < -1)
                                                    switch "a" {
                                                    case "a":
                                                        var v54 uint64 = v4
                                                        v55 := GenS1{F0: -623, F1: ""}
                                                        _ = v54
                                                        _ = v55
                                                    case "\n\t":
                                                        v4 += 933
                                                    }
                                                    var v56 string = "foo bar"
                                                    var v57 int = 247
                                                    if v2 {
                                                        v58 := []int{127, fi29, (fi25 - -1), 32767}
                                                        if true {
                                                            switch (v56 + "hello") {
                                                            case "foo bar":
                                                                v59 := GenS4{F0: (fi29 | (fi29 & fi25)), F1: (-32768 == 0)}
                                                                _ = v59
                                                            case "hello":
                                                                var v60 int32 = -611
                                                                _ = v60
                                                            default:
                                                                if (2147483647 > (fi25 + -122)) {
                                                                    if (32767 >= -827) {
                                                                        var v61 uint8 = 143
                                                                        v56 = v56
                                                                        var v62 bool = (42 >= 42)
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                    if (156 > 2) {
                                                                        v3 -= v3
                                                                        var v63 float32 = ((((float32((0.0 - v3)) + (float32(0.0) + ((3.14159 - 0.1) - -2.25))) * float32(float64(3.14159))) * 0.1) * 1e10)
                                                                        _ = v63
                                                                    }
                                                                }
                                                            }
                                                            v64 := GenS4{F0: v57, F1: (2 == 255)}
                                                            _ = v64
                                                        }
                                                        delete(v1, fi29)
                                                        if ((v2 || (1 <= 285)) && v2) {
                                                            var v65 uint32 = 0
                                                            switch "hello" {
                                                            case "":
                                                                if (!v2) {
                                                                    if (fi25 >= -209) {
                                                                        if len(v58) >= 2 {
                                                                            v58 = v58[1:len(v58)]
                                                                        }
                                                                        if true {
                                                                            if (v2 || v2) {
                                                                                v58 = append(v58, -1)
                                                                                fmt.Println(v4)
                                                                                v66 := GenS5{F0: -293}
                                                                                _ = v66
                                                                            }
                                                                        }
                                                                    }
                                                                    switch v57 {
                                                                    case 42:
                                                                        switch ((v56 + v56) + "hello") {
                                                                        case "":
                                                                            v57 -= -954
                                                                        default:
                                                                            var v67 float64 = v3
                                                                            _ = v67
                                                                        }
                                                                    case -833:
                                                                        _ = len(v58)
                                                                    case -9223372036854775808:
                                                                        if v2 {
                                                                            if v2 {
                                                                                v68 := GenS8{F0: 127, F1: -788}
                                                                                switch 1 {
                                                                                case -1:
                                                                                    var v69 int8 = -1
                                                                                    _ = v69
                                                                                case -204:
                                                                                    v2 = v2
                                                                                default:
                                                                                    if v2 {
                                                                                        var v70 int64 = 431
                                                                                        _ = v70
                                                                                    }
                                                                                    v56 = v56
                                                                                    if v2 {
                                                                                        v57 = (fi29 ^ (fi29 & (v57 * (-(v57)))))
                                                                                        var v71 string = v56
                                                                                        _ = v71
                                                                                    }
                                                                                }
                                                                                _ = v68
                                                                            } else {
                                                                                var v72 float32 = ((3.14159 / (1.5 / 3.14159)) + 0.1)
                                                                                var v73 int64 = 42
                                                                                v73 += (-(v73))
                                                                                _ = v72
                                                                                _ = v73
                                                                            }
                                                                            var v74 uint8 = 205
                                                                            _ = v74
                                                                        }
                                                                    }
                                                                    switch (fi25 ^ (fi25 & 32767)) {
                                                                    case 921:
                                                                        var v75 bool = true
                                                                        _ = v75
                                                                    case -769:
                                                                        var v76 string = ("a" + "foo bar")
                                                                        _ = v76
                                                                    case -501:
                                                                        var v77 uint32 = 294
                                                                        _ = v77
                                                                    default:
                                                                        switch ((v56 + v56) + v56) {
                                                                        case "\n\t":
                                                                            var v78 int = fi25
                                                                            _ = v78
                                                                        default:
                                                                            var v79 int32 = 2147483647
                                                                            _ = v79
                                                                        }
                                                                        v58 = append(v58, int(v4))
                                                                        v57 = fi29
                                                                        v1[9223372036854775807] = v65
                                                                        var v80 uint8 = 128
                                                                        delete(v1, (fi25 << ((uint(v80)) & 63)))
                                                                        fmt.Println(fi29)
                                                                        _ = v80
                                                                    }
                                                                    var v81 int8 = 42
                                                                    v65 = (v65 << ((18446744073709551615) & 63))
                                                                    _ = v81
                                                                }
                                                            case "\n\t":
                                                                if true {
                                                                    switch int(v65) {
                                                                    case 2147483647:
                                                                        var v82 float64 = float64(0.1)
                                                                        _ = v82
                                                                    default:
                                                                        switch 1 {
                                                                        case -9223372036854775808:
                                                                            v2 = v2
                                                                        case -541:
                                                                            v83 := GenS1{F0: 0, F1: v56}
                                                                            _ = v83
                                                                        case 32767:
                                                                            var v84 int8 = -1
                                                                            _ = v84
                                                                        }
                                                                    }
                                                                } else {
                                                                    v85 := v1[v57]
                                                                    if (!true) {
                                                                        fmt.Println(fi25, v2, v4)
                                                                        switch 1 {
                                                                        case 1:
                                                                            fmt.Println(fi29, v56, v85)
                                                                        }
                                                                        if (v2 || ((v4 & 920) == v4)) {
                                                                            var v86 bool = v2
                                                                            switch (fi25 ^ (^(fi25))) {
                                                                            case 1:
                                                                                fmt.Println(v57, v57, fi25)
                                                                            case 128:
                                                                                if len(v58) > 0 {
                                                                                    v58[0] = fi29
                                                                                }
                                                                            case 42:
                                                                                if ((^(v4)) <= (v4 | (v4 & uint64(v65)))) {
                                                                                    switch (fi25 & (-(v57))) {
                                                                                    case -9223372036854775808:
                                                                                        _ = len(v58)
                                                                                        v4 -= (v4 / ((v4) | 1))
                                                                                    case -65:
                                                                                        fmt.Println(v2, fi25)
                                                                                    default:
                                                                                        switch (v57 + (fi25 % (((v57 & (-(fi25)))) | 1))) {
                                                                                        case 4294967295:
                                                                                            fmt.Println(fi25, fi25)
                                                                                            v56 = (v56 + v56)
                                                                                            if v86 {
                                                                                                var v87 float64 = v3
                                                                                                if (false && v86) {
                                                                                                    var v88 int8 = -112
                                                                                                    var v89 int64 = 979
                                                                                                    v56 = (v56 + v56)
                                                                                                    _ = v88
                                                                                                    _ = v89
                                                                                                } else {
                                                                                                    if (v86 || v86) {
                                                                                                        v57 += fi25
                                                                                                    }
                                                                                                    var v90 int32 = -1
                                                                                                    _ = v90
                                                                                                }
                                                                                                var v91 float64 = 3.14159
                                                                                                _ = v87
                                                                                                _ = v91
                                                                                            } else {
                                                                                                _ = len(v1)
                                                                                                v92 := GenS6{F0: -1}
                                                                                                v85 += (v85 / ((uint32(v4)) | 1))
                                                                                                _ = v92
                                                                                            }
                                                                                            var v93 int64 = -1
                                                                                            _ = v93
                                                                                        case -477:
                                                                                            v3 -= v3
                                                                                        default:
                                                                                            if len(v58) > 0 {
                                                                                                v58[0] = 0
                                                                                            }
                                                                                            var v94 int32 = -663
                                                                                            _ = v94
                                                                                        }
                                                                                    }
                                                                                    v95 := GenS7{F0: -1, F1: (v65 ^ v65)}
                                                                                    _ = v95
                                                                                }
                                                                            default:
                                                                                v65 = v85
                                                                                switch (((((v56 + v56) + v56) + v56) + "a") + "foo bar") {
                                                                                case "":
                                                                                    v58 = append(v58, fi25)
                                                                                case "foo bar":
                                                                                    v86 = v2
                                                                                }
                                                                                if v2 {
                                                                                    if false {
                                                                                        var v96 uint64 = (v4 | v4)
                                                                                        var v97 uint8 = 125
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                    }
                                                                                    switch v56 {
                                                                                    case "hello":
                                                                                        v57 -= -2147483648
                                                                                    default:
                                                                                        switch "a" {
                                                                                        case "foo bar":
                                                                                            v86 = v2
                                                                                        case "hello":
                                                                                            if v2 {
                                                                                                var v98 float32 = -2.25
                                                                                                fmt.Println(v65, v98)
                                                                                                _ = v98
                                                                                            }
                                                                                        default:
                                                                                            var v99 uint8 = 53
                                                                                            _ = v99
                                                                                        }
                                                                                    }
                                                                                    v100 := GenS9{F0: 1000, F1: 32767, F2: 4294967295}
                                                                                    _ = v100
                                                                                } else {
                                                                                    fmt.Println(v56, v65)
                                                                                }
                                                                                var v101 int64 = 749
                                                                                _ = v101
                                                                            }
                                                                            delete(v1, (fi25 * 921))
                                                                            _ = v86
                                                                        } else {
                                                                            fmt.Println(v2)
                                                                            var v102 int8 = 1
                                                                            _ = v102
                                                                        }
                                                                        v2 = ((42 != 9223372036854775807) || (!v2))
                                                                    } else {
                                                                        switch -1 {
                                                                        case 127:
                                                                            fmt.Println(fi25)
                                                                        case -748:
                                                                            if ((v3 >= v3) && v2) {
                                                                                var v103 int = fi25
                                                                                if true {
                                                                                    v85 = uint32(v4)
                                                                                }
                                                                                _ = v103
                                                                            } else {
                                                                                v2 = false
                                                                                var v104 bool = (!false)
                                                                                v105 := GenS2{F0: (!(!(v104 && v104))), F1: v65, F2: 32767}
                                                                                var v106 int16 = 451
                                                                                _ = v104
                                                                                _ = v105
                                                                                _ = v106
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v85
                                                                }
                                                            }
                                                            var v107 uint8 = 128
                                                            _ = v107
                                                            _ = v65
                                                        }
                                                        _ = v58
                                                    } else {
                                                        v4 = 450
                                                    }
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        case "\n\t":
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v108 int8 = -128
                                                _ = v108
                                            }
                                            if v2 {
                                                switch 127 {
                                                case -256:
                                                    fmt.Println(v4, fi25, fi25)
                                                case 127:
                                                    v4 += (v4 - 42)
                                                }
                                            } else {
                                                var v109 int16 = -32768
                                                var v110 uint32 = 2
                                                var v111 bool = ((^(v4)) <= v4)
                                                _ = v109
                                                _ = v110
                                                _ = v111
                                            }
                                            var v112 bool = true
                                            switch "" {
                                            case "a":
                                                v112 = v112
                                            default:
                                                v113 := []int{(^(fi29)), (fi29 >> ((4294967295) & 63))}
                                                _ = v113
                                            }
                                            v3 -= v3
                                            switch (^(fi25)) {
                                            case 2147483647:
                                                var v114 int = (^(fi29))
                                                _ = v114
                                            case 9223372036854775807:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v115 int16 = -914
                                                    _ = v115
                                                case "\n\t":
                                                    v116 := GenS10{F0: 53, F1: 0, F2: 127}
                                                    v117 := []int{(^(fi29)), 128, (fi29 | (fi29 / ((644) | 1))), (-(fi25))}
                                                    _ = v116
                                                    _ = v117
                                                case "hello":
                                                    fmt.Println(fi25, v3)
                                                }
                                            }
                                            v118 := []int{fi29, fi25, 128}
                                            _ = v112
                                            _ = v118
                                        case "foo bar":
                                            v2 = (true || (!(v4 <= 2)))
                                        }
                                    case -30:
                                        v3 = 3.14159
                                    case -1:
                                        if false {
                                            if v2 {
                                                if (42 <= -9223372036854775808) {
                                                    delete(v1, fi29)
                                                    if v2 {
                                                        var v119 uint32 = 2
                                                        if v2 {
                                                            v120 := GenS8{F0: 168, F1: -2147483648}
                                                            var v121 uint = 42
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                        if (!true) {
                                                            switch fi25 {
                                                            case 128:
                                                                v3 -= (1.5 + float64(0.1))
                                                            case -1:
                                                                fmt.Println(v2, fi25)
                                                            case 32768:
                                                                if (v2 || v2) {
                                                                    switch ("\n\t" + ("a" + "")) {
                                                                    case "\n\t":
                                                                        var v122 float32 = (1e10 - 3.14159)
                                                                        _ = v122
                                                                    default:
                                                                        fmt.Println(v2)
                                                                    }
                                                                }
                                                            default:
                                                                if v2 {
                                                                    v3 -= (float64(1e10) + 1.5)
                                                                } else {
                                                                    var v123 uint16 = 310
                                                                    _ = v123
                                                                }
                                                            }
                                                            var v124 int = (^(fi25))
                                                            if ((32767 > 223) && (-2.25 != 0.0)) {
                                                                if v2 {
                                                                    if (1 >= -872) {
                                                                        var v125 uint64 = 415
                                                                        v126 := GenS11{F0: -128}
                                                                        v127 := []int{(fi25 << ((946) & 63)), (v124 << ((2) & 63))}
                                                                        _ = v125
                                                                        _ = v126
                                                                        _ = v127
                                                                    } else {
                                                                        var v128 int = v124
                                                                        v1[(fi29 >> ((uint(v4)) & 63))] = 710
                                                                        v129 := GenS2{F0: (0 < 42), F1: (v119 / ((v119) | 1)), F2: 42}
                                                                        _ = v128
                                                                        _ = v129
                                                                    }
                                                                } else {
                                                                    v130 := GenS12{F0: (^(v119))}
                                                                    var v131 uint = 265
                                                                    _ = v130
                                                                    _ = v131
                                                                }
                                                            } else {
                                                                v132 := GenS13{F0: (v119 % ((uint32(v4)) | 1)), F1: 1, F2: -1}
                                                                _ = v132
                                                            }
                                                            _ = v124
                                                        }
                                                        fmt.Println(fi29, fi25, v2)
                                                        v3 = float64(0.1)
                                                        _ = v119
                                                    }
                                                    var v133 bool = (95 > 42)
                                                    _ = v133
                                                } else {
                                                    if false {
                                                        switch fi29 {
                                                        case 2147483647:
                                                            var v134 float32 = -2.25
                                                            _ = v134
                                                        }
                                                        if v2 {
                                                            v135 := []int{(fi25 << ((83) & 63)), fi25, 0}
                                                            v136 := GenS8{F0: 1, F1: -68}
                                                            var v137 int8 = -63
                                                            _ = v135
                                                            _ = v136
                                                            _ = v137
                                                        } else {
                                                            v138 := []int{529, (fi29 & 873), fi25, fi25}
                                                            switch (-(fi25)) {
                                                            case 9223372036854775807:
                                                                if ("\n\t" > ("foo bar" + (((("a" + "\n\t") + "hello") + ("hello" + "\n\t")) + "\n\t"))) {
                                                                    if len(v138) > 0 {
                                                                        v138[0] = fi25
                                                                    }
                                                                    v3 -= (v3 * v3)
                                                                }
                                                            default:
                                                                v139 := GenS14{F0: (v2 && v2)}
                                                                _ = v139
                                                            }
                                                            var v140 int64 = -962
                                                            _ = v138
                                                            _ = v140
                                                        }
                                                    }
                                                    v2 = (uint(v4) > 430)
                                                    v141 := GenS15{F0: (v4 + (v4 & (v4 >> ((65535) & 63)))), F1: (false || (v3 < float64(float32(v3))))}
                                                    _ = v141
                                                }
                                                fmt.Println(v3)
                                                var v142 bool = v2
                                                _ = v142
                                            }
                                            switch "foo bar" {
                                            case "":
                                                if false {
                                                    v4 = (v4 % (((v4 / ((1) | 1))) | 1))
                                                    var v143 uint8 = 0
                                                    v144 := GenS16{F0: v4}
                                                    _ = v143
                                                    _ = v144
                                                }
                                            case "a":
                                                delete(v1, fi25)
                                            case "hello":
                                                var v145 uint16 = 614
                                                _ = v145
                                            }
                                            fmt.Println(fi29, v3, fi25)
                                        } else {
                                            _ = len(v1)
                                        }
                                    }
                                }
                            default:
                                if (-2.25 > -2.25) {
                                    v146 := GenS17{F0: fi29}
                                    v2 = (!true)
                                    if false {
                                        switch "a" {
                                        case "":
                                            var v147 uint16 = 0
                                            _ = v147
                                        default:
                                            v148 := GenS17{F0: (fi25 ^ (fi25 * 9223372036854775807))}
                                            if v2 {
                                                var v149 uint64 = v4
                                                var v150 bool = v2
                                                _ = v149
                                                _ = v150
                                            } else {
                                                var v151 bool = v2
                                                v3 = -2.25
                                                v151 = v151
                                                v1[2147483647] = 2
                                                if v151 {
                                                    var v152 float32 = (1.5 + (3.14159 * 0.1))
                                                    _ = v152
                                                }
                                                _ = v151
                                            }
                                            _ = v148
                                        }
                                    } else {
                                        fmt.Println(fi25, fi25, v3)
                                        fmt.Println(v2)
                                        var v153 int8 = 96
                                        v154 := GenS8{F0: 42, F1: 42}
                                        _ = v153
                                        _ = v154
                                    }
                                    switch (fi25 | 127) {
                                    case 32768:
                                        delete(v1, (-(fi29)))
                                    case -2147483648:
                                        var v155 uint8 = 47
                                        _ = v155
                                    case 128:
                                        fmt.Println(v4)
                                    }
                                    fmt.Println(fi25, v2)
                                    v4 = 42
                                    _ = v146
                                } else {
                                    fmt.Println(fi25)
                                    if false {
                                        var v156 uint64 = v4
                                        var v157 int32 = -850
                                        fmt.Println(fi25, v156)
                                        _ = v156
                                        _ = v157
                                    } else {
                                        v158 := []int{668, fi29, fi29}
                                        if (42 != 481) {
                                            v159 := GenS18{F0: 2147483647, F1: ("\n\t" + "")}
                                            _ = v159
                                        }
                                        fmt.Println(v4, v4, fi25)
                                        _ = v158
                                    }
                                    v160 := []int{(fi25 & fi29)}
                                    _ = v160
                                }
                            }
                        case "foo bar":
                            fmt.Println(v2, v2)
                            var v161 int8 = 1
                            _ = v161
                        default:
                            var v162 uint8 = 0
                            _ = v162
                        }
                    }
                    v163 := GenS19{F0: fi25}
                    _ = fi29
                    _ = v163
                }
            }
            fmt.Println(v4, v2)
            _ = fi25
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
