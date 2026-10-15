package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint; F1 int64; F2 int8 }
type GenS8 struct { F0 int32; F1 string; F2 int32 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 bool; F1 bool; F2 string }
type GenS12 struct { F0 int; F1 uint }
type GenS13 struct { F0 string; F1 string }
func genFn0() uint16 {
    var v4 uint8 = 42
    for fi5 := 0; fi5 < 14; fi5++ {
        switch -9223372036854775808 {
        case 127:
            v6 := []int{}
            _ = v6
        }
        switch -516 {
        case -1:
            v7 := []int{(fi5 << ((2) & 63))}
            _ = v7
        default:
            var v8 uint16 = 1
            _ = v8
        }
        fmt.Println(fi5, fi5)
        _ = fi5
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 13; fi10++ {
        switch ("hello" + ("\n\t" + "")) {
        case "":
            for fi11 := 0; fi11 < 14; fi11++ {
                v4 = v4
                if true {
                    fmt.Println(v4, fi10)
                } else {
                    var v12 int16 = 32
                    fmt.Println(fi10, fi11)
                    var v13 uint = 987
                    _ = v12
                    _ = v13
                }
                switch ("a" + "a") {
                case "hello":
                    var v14 int = (fi11 - 32768)
                    _ = v14
                case "a":
                    v4 += (v4 | 255)
                }
                _ = fi11
            }
        case "a":
            var v15 int16 = -723
            _ = v15
        }
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            if len(v9) > 0 {
                v9[0] = fi10
            }
        case "\n\t":
            for fi16 := 0; fi16 < 12; fi16++ {
                v4 = v4
                fmt.Println(v4, v4, fi10)
                v17 := GenS1{F0: 2, F1: 256}
                switch ("\n\t" + "foo bar") {
                case "hello":
                    fmt.Println(v4, v4, v4)
                default:
                    if len(v9) > 0 {
                        v9[0] = fi16
                    }
                }
                _ = fi16
                _ = v17
            }
        case "hello":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        default:
            fmt.Println(fi10, fi10)
        }
        _ = fi10
    }
    if (!true) {
        _ = len(v9)
    } else {
        var v18 int64 = -761
        v19 := GenS1{F0: 42, F1: 65535}
        var v20 uint32 = 65535
        _ = v18
        _ = v19
        _ = v20
    }
    _ = v4
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 839
    var v3 float32 = -2.25
    v21 := []int{808, -611, 796}
    v2 -= v2
    switch "a" {
    case "":
        fmt.Println(v2, v2)
    default:
        fmt.Println(v3)
    }
    fmt.Println(v3)
    var v22 uint32 = (v2 ^ 2147483647)
    v2 = v22
    switch "a" {
    case "\n\t":
        for fi23 := 0; fi23 < 3; fi23++ {
            v3 = 1.5
            for fi24 := 0; fi24 < 10; fi24++ {
                v25 := v1[fi24]
                v22 = (v25 + (v25 << ((2) & 63)))
                if (((-(fi24)) >= 226) || false) {
                    if (0 < 218) {
                        fmt.Println(v22, v3, fi23)
                    } else {
                        var v26 string = (("" + ("foo bar" + "hello")) + "\n\t")
                        _ = v26
                    }
                }
                if (float32((0.1 / float64(1.5))) > float32(3.14159)) {
                    fmt.Println(v25, v22)
                    if (53 < 0) {
                        var v27 int16 = -755
                        _ = v27
                    } else {
                        var v28 uint = 1
                        v29 := GenS2{F0: fi23, F1: 32767}
                        v3 = ((v3 / (v3 * v3)) * (v3 + 1.5))
                        _ = v28
                        _ = v29
                    }
                    v2 = v25
                } else {
                    switch (fi24 ^ 622) {
                    case -2147483648:
                        v30 := GenS3{F0: v25}
                        _ = v30
                    default:
                        if false {
                            v31 := GenS3{F0: 2147483647}
                            _ = v31
                        }
                    }
                    v25 = v22
                    v25 = v25
                }
                _ = fi24
                _ = v25
            }
            if len(v21) > 0 {
                v21[0] = -950
            }
            _ = fi23
        }
    default:
        switch (("" + "hello") + ("a" + ("\n\t" + "\n\t"))) {
        case "":
            if true {
                v32 := GenS1{F0: 42, F1: 65535}
                fmt.Println(v2, v22)
                _ = v32
            }
        }
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 12; fi34++ {
            if (4294967295 >= 2) {
                v22 = (^(v2))
                var v35 int16 = 0
                if true {
                    switch ((("foo bar" + ("a" + ("hello" + "\n\t"))) + "\n\t") + "hello") {
                    case "a":
                        v3 -= v3
                    }
                }
                var v36 uint8 = 1
                var v37 float64 = -2.25
                _ = v35
                _ = v36
                _ = v37
            } else {
                var v38 uint32 = (^(v22))
                v38 = (v22 % ((v2) | 1))
                _ = len(v21)
                _ = v38
            }
            fmt.Println(v2, v22, v3)
            v39 := GenS4{F0: 42, F1: 811}
            _ = fi34
            _ = v39
        }
        var v40 uint64 = 18446744073709551615
        _ = fi33
        _ = v40
    }
    v2 -= v22
    v22 = 65535
    v2 -= 906
    var v41 int8 = -1
    v42 := GenS4{F0: 214, F1: -772}
    if (1 == v2) {
        v43 := GenS2{F0: 32768, F1: 0}
        fmt.Println(v41, v22)
        for fi44 := 0; fi44 < 8; fi44++ {
            if false {
                if false {
                    var v45 bool = false
                    if v45 {
                        if ((v41 - -1) > v41) {
                            for fi46 := 0; fi46 < 16; fi46++ {
                                switch fi46 {
                                case 246:
                                    fmt.Println(fi46, v3)
                                case -9223372036854775808:
                                    switch fi44 {
                                    case 42:
                                        var v47 uint32 = (v22 * 65535)
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                        _ = v47
                                    }
                                }
                                _ = fi46
                            }
                            switch fi44 {
                            case 15:
                                for fi48 := 0; fi48 < 3; fi48++ {
                                    if ((false && (true && v45)) || ((128 > 42) && false)) {
                                        if (3.14159 <= ((3.14159 + 1e10) / 1.5)) {
                                            var v49 uint32 = (v2 % ((v2) | 1))
                                            if len(v21) > 0 {
                                                v21[0] = (fi48 + fi44)
                                            }
                                            _ = v49
                                        }
                                        var v50 uint8 = 2
                                        switch "hello" {
                                        case "foo bar":
                                            v51 := GenS5{F0: (v41 | (v41 - v41)), F1: 42}
                                            _ = v51
                                        case "\n\t":
                                            switch ("" + "a") {
                                            case "":
                                                fmt.Println(v2, v3)
                                            case "\n\t":
                                                v52 := GenS1{F0: 2, F1: 255}
                                                _ = v52
                                            case "a":
                                                fmt.Println(fi48)
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v22 -= v22
                                                case "":
                                                    if v45 {
                                                        v21 = append(v21, (fi48 | fi44))
                                                        delete(v1, 2147483647)
                                                        var v53 uint = 817
                                                        _ = v53
                                                    } else {
                                                        fmt.Println(v41)
                                                    }
                                                default:
                                                    v22 = (v2 / (((^(v22))) | 1))
                                                }
                                            }
                                            var v54 uint64 = 933
                                            switch (fi44 * (fi44 + fi44)) {
                                            case 1:
                                                var v55 int64 = 371
                                                _ = v55
                                            case 9223372036854775807:
                                                fmt.Println(fi44, v50, fi44)
                                            case -539:
                                                v45 = ((3.14159 == 0.0) || (536 <= 705))
                                            }
                                            var v56 int64 = 9223372036854775807
                                            switch "hello" {
                                            case "foo bar":
                                                v57 := GenS2{F0: (fi44 * fi44), F1: 896}
                                                _ = v57
                                            }
                                            _ = v54
                                            _ = v56
                                        default:
                                            if true {
                                                v58 := GenS3{F0: (v2 & v22)}
                                                var v59 uint32 = (v2 | (v2 >> ((65535) & 63)))
                                                switch (fi48 & fi44) {
                                                case 32767:
                                                    v41 = (v41 | 100)
                                                case 9223372036854775807:
                                                    if len(v21) > 0 {
                                                        v21[0] = fi48
                                                    }
                                                default:
                                                    var v60 float64 = 1.5
                                                    _ = v60
                                                }
                                                var v61 bool = v45
                                                v62 := v1[fi44]
                                                v63 := GenS6{F0: 127, F1: (fi48 ^ (fi44 >> ((1) & 63)))}
                                                _ = v58
                                                _ = v59
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            }
                                            var v64 int = -2147483648
                                            v65 := GenS3{F0: (v22 ^ 65535)}
                                            _ = v64
                                            _ = v65
                                        }
                                        fmt.Println(v45, fi48, v41)
                                        v66 := GenS7{F0: 2, F1: -465, F2: 1}
                                        _ = v50
                                        _ = v66
                                    }
                                    v41 = (-(v41))
                                    _ = fi48
                                }
                            case 127:
                                v41 = v41
                            case -9223372036854775808:
                                v2 = 566
                            }
                            fmt.Println(v45, v22, v41)
                        } else {
                            v22 = v22
                            if true {
                                if ((-(fi44)) != fi44) {
                                    _ = len(v1)
                                    for fi67 := 0; fi67 < 9; fi67++ {
                                        var v68 int = (fi44 / (((fi44 + (fi67 + fi67))) | 1))
                                        v2 += (^(v22))
                                        if (!(!(genFn0() == 164))) {
                                            var v69 string = ("\n\t" + "foo bar")
                                            var v70 bool = v45
                                            _ = v69
                                            _ = v70
                                        } else {
                                            v71 := GenS2{F0: fi67, F1: 534}
                                            v22 += (v22 >> ((4294967295) & 63))
                                            _ = v71
                                        }
                                        _ = fi67
                                        _ = v68
                                    }
                                } else {
                                    v45 = false
                                }
                                _ = len(v1)
                            }
                        }
                        var v72 int16 = -32768
                        fmt.Println(v41)
                        _ = v72
                    } else {
                        var v73 uint32 = 442
                        v74 := GenS8{F0: 0, F1: ("\n\t" + "hello"), F2: 283}
                        _ = v73
                        _ = v74
                    }
                    _ = v45
                }
                fmt.Println(v3, fi44)
                for fi75 := 0; fi75 < 12; fi75++ {
                    v76 := GenS3{F0: v2}
                    _ = fi75
                    _ = v76
                }
            }
            for fi77 := 0; fi77 < 20; fi77++ {
                v78 := v1[fi77]
                v3 -= (0.0 + v3)
                var v79 string = ("\n\t" + "foo bar")
                _ = fi77
                _ = v78
                _ = v79
            }
            fmt.Println(v22, v2)
            _ = fi44
        }
        _ = v43
    } else {
        switch "" {
        case "foo bar":
            switch 32768 {
            case -554:
                fmt.Println(v2)
            case 9223372036854775807:
                var v80 int32 = 107
                _ = v80
            case 127:
                v81 := GenS8{F0: 205, F1: "\n\t", F2: 2147483647}
                _ = v81
            default:
                if (355 <= -1) {
                    if false {
                        _ = len(v1)
                        if (255 >= 255) {
                            switch 4294967295 {
                            case 127:
                                var v82 int16 = 1
                                _ = v82
                            case 774:
                                v22 = v22
                            default:
                                v83 := GenS9{F0: 940, F1: 127}
                                _ = v83
                            }
                            v2 -= 831
                        }
                        for fi84 := 0; fi84 < 9; fi84++ {
                            switch (fi84 - 32767) {
                            case -455:
                                for fi85 := 0; fi85 < 17; fi85++ {
                                    var v86 bool = (-1 <= 1)
                                    v87 := GenS10{F0: 0, F1: 423}
                                    v2 = v22
                                    if true {
                                        var v88 int = fi84
                                        if ("" == "foo bar") {
                                            if v86 {
                                                if false {
                                                    if (0 == -9223372036854775808) {
                                                        fmt.Println(fi84, fi85, v88)
                                                        var v89 string = "\n\t"
                                                        _ = v89
                                                    } else {
                                                        if (!v86) {
                                                            v41 = (v41 << ((4294967295) & 63))
                                                            fmt.Println(v41, v3)
                                                        }
                                                    }
                                                    if (-2147483648 > 42) {
                                                        fmt.Println(fi84, v22)
                                                    } else {
                                                        v2 = (v2 * (v22 - 2147483647))
                                                    }
                                                    v90 := GenS1{F0: 123, F1: 42}
                                                    var v91 uint32 = v22
                                                    v92 := GenS2{F0: v88, F1: 0}
                                                    fmt.Println(v86)
                                                    _ = v90
                                                    _ = v91
                                                    _ = v92
                                                }
                                            } else {
                                                v1[fi85] = v22
                                                v93 := v1[32767]
                                                if ((!v86) && (false || true)) {
                                                    var v94 uint64 = 1
                                                    var v95 int8 = (^(v41))
                                                    _ = v94
                                                    _ = v95
                                                }
                                                _ = v93
                                            }
                                            var v96 float64 = -2.25
                                            fmt.Println(fi85, v96)
                                            _ = v96
                                        }
                                        var v97 uint = 65535
                                        _ = v88
                                        _ = v97
                                    }
                                    _ = len(v1)
                                    _ = fi85
                                    _ = v86
                                    _ = v87
                                }
                            case -2147483648:
                                var v98 float32 = v3
                                _ = v98
                            case 9223372036854775807:
                                var v99 uint = 2
                                _ = v99
                            default:
                                switch ("foo bar" + "\n\t") {
                                case "":
                                    var v100 int8 = 127
                                    _ = v100
                                case "a":
                                    v101 := v1[(fi84 | fi84)]
                                    _ = v101
                                case "\n\t":
                                    v102 := GenS11{F0: false, F1: true, F2: "a"}
                                    _ = v102
                                }
                            }
                            fmt.Println(v41, fi84)
                            _ = fi84
                        }
                    } else {
                        v103 := GenS6{F0: 103, F1: 1}
                        _ = v103
                    }
                } else {
                    v104 := GenS12{F0: 32767, F1: 18446744073709551615}
                    var v105 int8 = v41
                    for fi106 := 0; fi106 < 7; fi106++ {
                        v3 = (-2.25 + 3.14159)
                        for fi107 := 0; fi107 < 5; fi107++ {
                            v108 := GenS7{F0: 18446744073709551615, F1: 1, F2: (v105 & 0)}
                            var v109 uint = 4294967295
                            _ = fi107
                            _ = v108
                            _ = v109
                        }
                        _ = fi106
                    }
                    _ = v104
                    _ = v105
                }
            }
        case "\n\t":
            for fi110 := 0; fi110 < 7; fi110++ {
                for fi111 := 0; fi111 < 4; fi111++ {
                    if ("\n\t" <= (("" + "foo bar") + "\n\t")) {
                        v112 := GenS7{F0: 0, F1: 821, F2: v41}
                        _ = v112
                    } else {
                        v41 -= 124
                        var v113 uint32 = (v22 % (((v2 * v22)) | 1))
                        _ = v113
                    }
                    if false {
                        switch (^(fi111)) {
                        case -634:
                            var v114 int = (fi110 | (fi110 * fi110))
                            _ = v114
                        case -409:
                            v3 += float32(1e10)
                        case 0:
                            var v115 uint32 = 65535
                            _ = v115
                        default:
                            v41 += v41
                            var v116 uint = 18446744073709551615
                            _ = v116
                        }
                        var v117 int64 = -831
                        _ = v117
                    }
                    v21 = append(v21, fi110)
                    _ = fi111
                }
                v118 := GenS13{F0: "", F1: "hello"}
                _ = fi110
                _ = v118
            }
        default:
            v21 = append(v21, -9223372036854775808)
        }
        v21 = append(v21, -408)
        fmt.Println(v22, v41, v2)
    }
    if (!false) {
        for fi119 := 0; fi119 < 5; fi119++ {
            v2 -= (v22 % ((v22) | 1))
            for fi120 := 0; fi120 < 5; fi120++ {
                v2 = (^(v2))
                v21 = append(v21, (-(fi119)))
                _ = fi120
            }
            _ = fi119
        }
    } else {
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v41
    _ = v42
}
