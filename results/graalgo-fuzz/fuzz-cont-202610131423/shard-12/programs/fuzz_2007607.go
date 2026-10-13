package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 uint16; F2 uint }

func main() {
    var v1 uint8 = 42
    var v2 int16 = -195
    var v3 uint64 = 1
    v4 := []string{"foo bar", "a"}
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int32 = 2147483647
        _ = fi5
        _ = v6
    }
    fmt.Println(v3, v1, v1)
    switch 2147483647 {
    case -598:
        var v7 uint32 = 1
        _ = v7
    }
    var v8 int8 = 76
    if (false || true) {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: v8, F1: -9223372036854775808, F2: (false || ((!(!true)) && false))}
            fmt.Println(v2, v8)
            _ = fi9
            _ = v10
        }
        var v11 int32 = 42
        for fi12 := 0; fi12 < 5; fi12++ {
            switch "\n\t" {
            case "a":
                if (v3 <= v3) {
                    fmt.Println(v8)
                }
            case "\n\t":
                var v13 uint8 = v1
                _ = v13
            case "":
                switch (fi12 << ((1) & 63)) {
                case 120:
                    if len(v4) > 0 {
                        v4[0] = (("\n\t" + "\n\t") + "foo bar")
                    }
                case 9223372036854775807:
                    var v14 int64 = -1
                    _ = v14
                case -541:
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                }
            }
            fmt.Println(fi12, v8, fi12)
            _ = fi12
        }
        _ = v11
    }
    v2 = v2
    for fi15 := 0; fi15 < 2; fi15++ {
        _ = len(v4)
        if (65535 <= 2) {
            switch "hello" {
            case "a":
                if false {
                    for fi16 := 0; fi16 < 7; fi16++ {
                        var v17 int8 = v8
                        v18 := GenS1{F0: (v17 & 127), F1: (fi15 + fi16), F2: false}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    var v19 float32 = 0.0
                    _ = len(v4)
                    _ = v19
                } else {
                    var v20 int32 = 42
                    v2 = int16(v8)
                    v3 = v3
                    _ = v20
                }
            case "":
                var v21 uint64 = (v3 + v3)
                _ = v21
            }
        }
        _ = fi15
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        switch (fi22 ^ fi22) {
        case 9223372036854775807:
            for fi23 := 0; fi23 < 15; fi23++ {
                v2 -= (v2 << ((42) & 63))
                _ = fi23
            }
        case -1:
            fmt.Println(fi22, v8)
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 uint32 = 4294967295
            switch -587 {
            case 0:
                var v26 float32 = (float32(0.1) + (1e10 * ((1e10 - (float32(3.14159) + float32(0.1))) / (1.5 - -2.25))))
                _ = v26
            default:
                if false {
                    if true {
                        if (!false) {
                            v27 := GenS2{F0: (fi24 ^ fi22), F1: (false || true)}
                            if false {
                                switch int(v1) {
                                case 127:
                                    var v28 uint = 909
                                    _ = v28
                                case -931:
                                    if false {
                                        v1 += (v1 >> ((255) & 63))
                                        fmt.Println(fi22, v8, v25)
                                    }
                                    v29 := GenS2{F0: 32767, F1: true}
                                    v2 = v2
                                    _ = v29
                                default:
                                    var v30 uint = 652
                                    _ = v30
                                }
                                v1 = v1
                            }
                            _ = v27
                        }
                        var v31 float32 = (0.1 * 3.14159)
                        switch "\n\t" {
                        case "foo bar":
                            switch "foo bar" {
                            case "":
                                var v32 bool = (false && (-1 >= 42))
                                _ = v32
                            case "foo bar":
                                if false {
                                    switch (fi22 % ((fi24) | 1)) {
                                    case 0:
                                        var v33 int8 = (-(v8))
                                        _ = v33
                                    default:
                                        var v34 int8 = v8
                                        _ = v34
                                    }
                                    v35 := GenS2{F0: (fi24 - fi24), F1: (((-(fi22)) <= fi22) || (!(63 <= -1)))}
                                    _ = len(v4)
                                    v4 = append(v4, "")
                                    switch (fi22 - (fi24 - 32768)) {
                                    case -946:
                                        v36 := GenS1{F0: (v8 ^ v8), F1: -663, F2: true}
                                        _ = v36
                                    case 0:
                                        var v37 int16 = v2
                                        _ = v37
                                    case 9223372036854775807:
                                        fmt.Println(v25, v31, v3)
                                    default:
                                        v8 = -1
                                    }
                                    _ = v35
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                    v3 = v3
                                    if (v1 != (v1 & 1)) {
                                        v38 := GenS2{F0: -297, F1: false}
                                        switch ("" + "a") {
                                        case "a":
                                            v2 = (v2 & (v2 / ((v2) | 1)))
                                        }
                                        _ = v38
                                    }
                                }
                            default:
                                if (false && false) {
                                    if (!true) {
                                        _ = len(v4)
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case "":
                                            switch fi22 {
                                            case 9223372036854775807:
                                                if ((1.5 < 3.14159) || false) {
                                                    if false {
                                                        var v39 uint64 = v3
                                                        switch fi22 {
                                                        case 32767:
                                                            v40 := GenS2{F0: fi22, F1: true}
                                                            var v41 uint64 = 1
                                                            _ = v40
                                                            _ = v41
                                                        case 4294967295:
                                                            v25 += (v25 / ((v25) | 1))
                                                        }
                                                        if (-9223372036854775808 < -90) {
                                                            var v42 int32 = -2147483648
                                                            var v43 int16 = v2
                                                            fmt.Println(v8, v1)
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            v2 += (v2 * (v2 + v2))
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        v31 = (((((v31 - v31) + 0.0) + v31) / v31) / float32(((0.1 + 3.14159) + 0.0)))
                                                        v44 := GenS3{F0: (v25 ^ (v25 / ((65535) | 1))), F1: fi22, F2: (v39 >> ((65535) & 63))}
                                                        v25 = (v25 & (v25 * v25))
                                                        switch (fi22 * (fi22 - (fi24 ^ fi22))) {
                                                        case 128:
                                                            fmt.Println(v1)
                                                        case -115:
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        default:
                                                            if true {
                                                                fmt.Println(v25, v2, v1)
                                                                fmt.Println(v2, v2, v2)
                                                                if false {
                                                                    v1 += (v1 * (v1 - 42))
                                                                    _ = len(v4)
                                                                }
                                                                if ((true && (2 == 137)) && (v2 > v2)) {
                                                                    v3 = (v39 | 0)
                                                                }
                                                            }
                                                        }
                                                        switch (^(fi22)) {
                                                        case 127:
                                                            var v45 uint16 = 0
                                                            _ = v45
                                                        default:
                                                            v8 = (^(v8))
                                                        }
                                                        _ = v39
                                                        _ = v44
                                                    } else {
                                                        if true {
                                                            v4 = append(v4, "foo bar")
                                                            _ = len(v4)
                                                        }
                                                        fmt.Println(v1)
                                                        v46 := GenS3{F0: (v25 >> ((18446744073709551615) & 63)), F1: (fi24 >> ((344) & 63)), F2: (v3 + v3)}
                                                        _ = v46
                                                    }
                                                } else {
                                                    var v47 uint = 255
                                                    if (127 < v1) {
                                                        var v48 int32 = int32(v47)
                                                        fmt.Println(v3, fi24, v31)
                                                        if (!true) {
                                                            v49 := GenS4{F0: (v2 << ((117) & 63)), F1: v47}
                                                            var v50 uint16 = 42
                                                            _ = len(v4)
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        if (((((0.1 + 0.1) / 1.5) * (1.5 * 0.0)) <= (3.14159 - -2.25)) && false) {
                                                            v51 := GenS3{F0: 42, F1: -157, F2: v3}
                                                            var v52 string = ("foo bar" + "")
                                                            v52 = (((v52 + ("a" + (((v52 + "") + ("a" + v52)) + v52))) + "a") + v52)
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        fmt.Println(v48)
                                                        _ = v48
                                                    }
                                                    if (false && true) {
                                                        var v53 int64 = 1
                                                        _ = v53
                                                    }
                                                    fmt.Println(fi24, fi22)
                                                    _ = v47
                                                }
                                            case -502:
                                                v31 = 3.14159
                                                if (("a" + "hello") < "a") {
                                                    v4 = append(v4, "hello")
                                                    switch ("foo bar" + "") {
                                                    case "foo bar":
                                                        v25 += v25
                                                    case "a":
                                                        v54 := GenS4{F0: v2, F1: 4294967295}
                                                        _ = v54
                                                    }
                                                    v4 = append(v4, "hello")
                                                }
                                                v1 = 255
                                            case -1:
                                                v2 -= (v2 >> ((255) & 63))
                                                v55 := GenS3{F0: 423, F1: (-(fi24)), F2: v3}
                                                _ = v55
                                            default:
                                                if ((-(fi22)) == fi24) {
                                                    v31 = (v31 - v31)
                                                    v56 := GenS5{F0: 2}
                                                    switch (fi24 | fi24) {
                                                    case 737:
                                                        v57 := GenS5{F0: 0}
                                                        _ = v57
                                                    }
                                                    _ = v56
                                                }
                                            }
                                        default:
                                            if true {
                                                v31 = v31
                                                if false {
                                                    switch "" {
                                                    case "foo bar":
                                                        v4 = append(v4, "")
                                                    case "a":
                                                        switch "hello" {
                                                        case "":
                                                            var v58 bool = false
                                                            _ = v58
                                                        default:
                                                            v31 = v31
                                                        }
                                                    case "hello":
                                                        if ((!true) && (1e10 >= 1.5)) {
                                                            fmt.Println(v3, v2, fi24)
                                                        }
                                                    default:
                                                        switch (fi24 - fi22) {
                                                        case 32768:
                                                            v59 := GenS5{F0: 168}
                                                            _ = v59
                                                        case 569:
                                                            fmt.Println(v31, fi24, v1)
                                                            switch (("" + ("" + "")) + "foo bar") {
                                                            case "\n\t":
                                                                var v60 uint64 = v3
                                                                _ = v60
                                                            }
                                                        case 745:
                                                            if (-2.25 == 3.14159) {
                                                                var v61 string = ("hello" + "foo bar")
                                                                _ = v61
                                                            } else {
                                                                v25 += v25
                                                            }
                                                        default:
                                                            var v62 float32 = float32(3.14159)
                                                            _ = v62
                                                        }
                                                        var v63 int = (-(fi22))
                                                        v4 = append(v4, "\n\t")
                                                        _ = v63
                                                    }
                                                    if (false || true) {
                                                        var v64 uint32 = (^(v25))
                                                        var v65 uint = 18446744073709551615
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    var v66 uint16 = 231
                                                    v1 = uint8(v66)
                                                    _ = v66
                                                } else {
                                                    switch (fi22 >> ((641) & 63)) {
                                                    case 32768:
                                                        v67 := GenS5{F0: 381}
                                                        _ = v67
                                                    case 4294967295:
                                                        if false {
                                                            var v68 uint64 = (v3 << ((1) & 63))
                                                            _ = v68
                                                        }
                                                    case 2147483647:
                                                        var v69 uint8 = (^(v1))
                                                        _ = v69
                                                    }
                                                    var v70 float64 = float64(float32(float64(0.1)))
                                                    _ = v70
                                                }
                                                var v71 int8 = 100
                                                _ = v71
                                            }
                                        }
                                    }
                                } else {
                                    var v72 uint16 = 1
                                    var v73 uint16 = v72
                                    _ = v72
                                    _ = v73
                                }
                            }
                        default:
                            v2 = (v2 | 1)
                        }
                        v8 = 106
                        if false {
                            var v74 int16 = 1
                            _ = v74
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = v31
                    } else {
                        var v75 uint64 = 619
                        v8 += (^(v8))
                        if false {
                            var v76 int = (fi24 << ((288) & 63))
                            _ = v76
                        } else {
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        }
                        _ = v75
                    }
                    var v77 float64 = (0.0 * 1e10)
                    if (true || (26 == -32768)) {
                        fmt.Println(v25, fi24)
                        v1 += 0
                    } else {
                        var v78 float64 = v77
                        v1 = v1
                        var v79 int64 = 844
                        _ = v78
                        _ = v79
                    }
                    _ = v77
                } else {
                    switch ("\n\t" + "hello") {
                    case "a":
                        fmt.Println(fi24, v8)
                    case "":
                        var v80 int16 = v2
                        var v81 uint32 = 42
                        _ = v80
                        _ = v81
                    case "\n\t":
                        var v82 float32 = 3.14159
                        _ = v82
                    }
                }
            }
            _ = len(v4)
            _ = fi24
            _ = v25
        }
        _ = fi22
    }
    var v83 int = 4294967295
    for fi84 := 0; fi84 < 14; fi84++ {
        var v85 uint = 42
        _ = fi84
        _ = v85
    }
    v8 -= (v8 - v8)
    v3 = v3
    if (!(!true)) {
        if (-505 < v83) {
            for fi86 := 0; fi86 < 9; fi86++ {
                if false {
                    for fi87 := 0; fi87 < 4; fi87++ {
                        v3 = v3
                        switch -9223372036854775808 {
                        case 9223372036854775807:
                            _ = len(v4)
                        }
                        switch v83 {
                        case 1:
                            v88 := GenS4{F0: v2, F1: 18446744073709551615}
                            _ = v88
                        default:
                            v4 = append(v4, "hello")
                        }
                        _ = fi87
                    }
                } else {
                    v89 := GenS6{F0: v2, F1: (^(v8))}
                    for fi90 := 0; fi90 < 2; fi90++ {
                        var v91 int64 = 635
                        if false {
                            _ = len(v4)
                            _ = len(v4)
                            v91 += 9223372036854775807
                        } else {
                            if len(v4) > 0 {
                                v4[0] = ("\n\t" + "hello")
                            }
                            v92 := GenS5{F0: 2}
                            _ = v92
                        }
                        if true {
                            fmt.Println(v2)
                        } else {
                            v2 = v2
                            v3 = 792
                            v91 += v91
                            if false {
                                if (false || (!false)) {
                                    fmt.Println(v83, v8)
                                    switch fi86 {
                                    case 621:
                                        var v93 int16 = (v2 % ((v2) | 1))
                                        _ = v93
                                    case 671:
                                        switch v83 {
                                        case 32768:
                                            _ = len(v4)
                                        default:
                                            switch "a" {
                                            case "foo bar":
                                                var v94 uint32 = 2147483647
                                                _ = v94
                                            case "hello":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "a":
                                                fmt.Println(v1, v8)
                                            }
                                        }
                                    case 0:
                                        var v95 uint8 = (v1 / (((^(v1))) | 1))
                                        _ = v95
                                    }
                                    _ = len(v4)
                                }
                            } else {
                                if true {
                                    if false {
                                        v96 := GenS7{F0: (-(v8))}
                                        _ = v96
                                    } else {
                                        var v97 int32 = 1
                                        v8 = 1
                                        _ = v97
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = ""
                                    }
                                }
                                switch ("foo bar" + ("foo bar" + "a")) {
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                default:
                                    fmt.Println(fi86, fi90)
                                }
                            }
                        }
                        _ = fi90
                        _ = v91
                    }
                    switch int(v1) {
                    case 42:
                        var v98 bool = false
                        _ = v98
                    case 127:
                        v8 = v8
                    case 821:
                        switch ("a" + ("" + ("hello" + ""))) {
                        case "a":
                            fmt.Println(v8, v8)
                        case "\n\t":
                            v3 -= (v3 - (^(v3)))
                        }
                    default:
                        _ = len(v4)
                    }
                    _ = v89
                }
                _ = fi86
            }
            _ = len(v4)
            var v99 int64 = -1
            _ = v99
        } else {
            v4 = append(v4, ("\n\t" + ((("\n\t" + "hello") + (("a" + "a") + (("foo bar" + "foo bar") + "hello"))) + ("foo bar" + ""))))
        }
        v2 = (^(v2))
        fmt.Println(v83)
    } else {
        if (false || true) {
            for fi100 := 0; fi100 < 1; fi100++ {
                v101 := GenS4{F0: v2, F1: 784}
                v1 = (v1 + 0)
                _ = fi100
                _ = v101
            }
            fmt.Println(v3, v83)
            switch "foo bar" {
            case "foo bar":
                switch 32768 {
                case 32768:
                    var v102 string = "hello"
                    _ = v102
                case 954:
                    if false {
                        var v103 uint64 = (v3 << ((42) & 63))
                        v1 = (v1 % ((v1) | 1))
                        _ = v103
                    }
                case 0:
                    v2 -= -259
                default:
                    _ = len(v4)
                }
            case "a":
                for fi104 := 0; fi104 < 3; fi104++ {
                    v83 = fi104
                    if false {
                        v105 := GenS1{F0: v8, F1: (^(fi104)), F2: (65535 != 18446744073709551615)}
                        for fi106 := 0; fi106 < 18; fi106++ {
                            v107 := GenS8{F0: "", F1: 1, F2: 113}
                            _ = fi106
                            _ = v107
                        }
                        _ = v105
                    } else {
                        v108 := GenS2{F0: (-(fi104)), F1: true}
                        _ = v108
                    }
                    _ = fi104
                }
            default:
                v109 := GenS7{F0: v8}
                _ = v109
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    if true {
        v4 = append(v4, "foo bar")
    }
    v2 = (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v83
}
