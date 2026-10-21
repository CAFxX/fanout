package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 uint; F1 bool; F2 int }
type GenS8 struct { F0 int64; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint32; F1 uint64; F2 string }
type GenS11 struct { F0 uint; F1 int8 }
type GenS12 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 bool; F1 uint16; F2 int64 }
type GenS15 struct { F0 uint64; F1 string }
type GenS16 struct { F0 bool; F1 int16 }
type GenS17 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS18 struct { F0 int64; F1 string }
type GenS19 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS20 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS21 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint32, p2 int64) bool {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{32767, fi6}
        v7 = append(v7, (fi6 + (-(fi6))))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int32 = 1
        v10 := []int{}
        switch "a" {
        case "foo bar":
            if false {
                fmt.Println(p1, p1, p2)
                p1 -= p1
            }
        case "\n\t":
            var v11 uint = 4294967295
            _ = v11
        case "a":
            switch -1 {
            case 682:
                for fi12 := 0; fi12 < 7; fi12++ {
                    v9 += (v9 * (v9 / (((v9 * v9)) | 1)))
                    fmt.Println(p0)
                    _ = fi12
                }
            case 9223372036854775807:
                v9 = v9
            }
        default:
            var v13 int32 = v9
            _ = v13
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    switch 42 {
    case -2147483648:
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := []int{}
            p0 = ((^(p1)) != p1)
            _ = fi14
            _ = v15
        }
    case 9223372036854775807:
        switch ("hello" + ("foo bar" + ("\n\t" + "a"))) {
        case "hello":
            v16 := []int{2147483647, -155, 31}
            _ = v16
        case "\n\t":
            v17 := []int{-2147483648, -9223372036854775808}
            _ = v17
        case "foo bar":
            var v18 string = ("hello" + "a")
            _ = v18
        }
    default:
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1e10
    v3 := make(map[string]string)
    var v4 int16 = -1
    var v5 int8 = 100
    v19 := GenS1{F0: 18446744073709551615, F1: -32768}
    if (36 == 4294967295) {
        if ((255 <= 980) || (true && false)) {
            if (2 != 37) {
                v5 = (v5 * -128)
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 uint64 = 1
                _ = fi20
                _ = v21
            }
        }
        var v22 int32 = -1
        v23 := GenS1{F0: 73, F1: v4}
        _ = v22
        _ = v23
    }
    if (!(!true)) {
        if ((1 < 255) || (42 == 792)) {
            switch "\n\t" {
            case "\n\t":
                v24 := GenS1{F0: 89, F1: (v4 * (v4 & v4))}
                _ = v24
            case "hello":
                switch "" {
                case "a":
                    for fi25 := 0; fi25 < 18; fi25++ {
                        v26 := GenS2{F0: 2, F1: 42, F2: 0}
                        if true {
                            var v27 float32 = (-2.25 / (1.5 + 3.14159))
                            _ = v27
                        } else {
                            v2 += v2
                        }
                        if (2 == 491) {
                            v28 := GenS3{F0: 1, F1: -1}
                            if (true && (0 <= 32767)) {
                                var v29 int32 = 2147483647
                                fmt.Println(v4)
                                var v30 uint64 = 1
                                _ = v29
                                _ = v30
                            }
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                    }
                }
            default:
                v2 = v2
            }
            v31 := GenS1{F0: 18446744073709551615, F1: v4}
            _ = v31
        } else {
            if (42 >= -329) {
                var v32 int8 = (v5 | v5)
                switch 128 {
                case -1:
                    v2 = float64((float32(v2) * float32(3.14159)))
                case -9223372036854775808:
                    v5 = v5
                default:
                    for fi33 := 0; fi33 < 16; fi33++ {
                        v34 := []int{(fi33 % ((269) | 1)), (fi33 / ((734) | 1)), (fi33 & (fi33 % ((127) | 1)))}
                        _ = fi33
                        _ = v34
                    }
                }
                fmt.Println(v2, v32)
                _ = v32
            } else {
                var v35 bool = false
                var v36 string = "foo bar"
                v37 := []int{9223372036854775807, -1, 138, 128}
                _ = v35
                _ = v36
                _ = v37
            }
            for fi38 := 0; fi38 < 10; fi38++ {
                v39 := GenS1{F0: 65535, F1: v4}
                _ = fi38
                _ = v39
            }
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            fmt.Println(fi40)
            var v41 bool = true
            _ = fi40
            _ = v41
        }
        switch 4294967295 {
        case -260:
            switch "" {
            case "hello":
                v4 = (v4 + 42)
            case "a":
                if true {
                    v42 := []int{}
                    _ = v42
                }
            case "\n\t":
                var v43 uint = 4294967295
                _ = v43
            }
        default:
            var v44 float32 = (0.1 - 1.5)
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 5; fi45++ {
        for fi46 := 0; fi46 < 8; fi46++ {
            var v47 uint16 = 1
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    var v48 uint16 = 65535
    var v49 uint8 = 1
    v2 = ((v2 + ((v2 * v2) + (v2 - v2))) + (v2 / v2))
    v50 := []int{32768}
    var v51 uint64 = 105
    if len(v50) >= 2 {
        v50 = v50[1:len(v50)]
    }
    switch 223 {
    case 4294967295:
        switch -554 {
        case 0:
            v52 := GenS2{F0: 2, F1: 4294967295, F2: 4294967295}
            _ = v52
        case 283:
            var v53 uint16 = 255
            _ = v53
        }
    case 32767:
        v4 += (-(v4))
    }
    switch 0 {
    case 4294967295:
        for fi54 := 0; fi54 < 5; fi54++ {
            v48 = v48
            _ = fi54
        }
    case 9223372036854775807:
        switch -787 {
        case 2147483647:
            v55 := GenS2{F0: 31, F1: v51, F2: 401}
            _ = v55
        case 627:
            var v56 float32 = float32(3.14159)
            _ = v56
        case 0:
            for fi57 := 0; fi57 < 5; fi57++ {
                for fi58 := 0; fi58 < 14; fi58++ {
                    switch (fi57 & (fi58 - (fi57 - 32767))) {
                    case 1:
                        if true {
                            var v59 uint64 = 18446744073709551615
                            v48 = (v48 & 0)
                            _ = v59
                        } else {
                            var v60 uint32 = uint32(v51)
                            fmt.Println(fi58)
                            fmt.Println(v2, fi58)
                            _ = v60
                        }
                    default:
                        var v61 float64 = ((v2 + ((v2 + (v2 / (-2.25 + v2))) / (0.1 - v2))) * (((-2.25 * v2) / ((0.0 / (v2 / -2.25)) / 1e10)) * 1.5))
                        _ = v61
                    }
                    switch -721 {
                    case 128:
                        v62 := GenS4{F0: fi58, F1: (v49 / ((11) | 1)), F2: (v48 ^ 32767)}
                        _ = v62
                    }
                    fmt.Println(v49, v49, v49)
                    _ = fi58
                }
                if ((!false) && (!false)) {
                    for fi63 := 0; fi63 < 6; fi63++ {
                        if ((false && true) || (1 != -9223372036854775808)) {
                            v64 := GenS5{F0: v4}
                            switch "foo bar" {
                            case "a":
                                v3["foo bar"] = "foo bar"
                            case "":
                                v65 := GenS5{F0: (v4 | (^(v4)))}
                                _ = v65
                            case "\n\t":
                                var v66 int8 = v5
                                _ = v66
                            default:
                                switch (fi63 % (((fi57 * 1)) | 1)) {
                                case 788:
                                    switch "foo bar" {
                                    case "foo bar":
                                        if (((float32(v2) + float32((float64(1e10) - 0.0))) / 0.1) != ((3.14159 + 1e10) / ((1.5 / 1.0) - (1e10 / 3.14159)))) {
                                            var v67 string = "a"
                                            _ = v67
                                        } else {
                                            switch (fi63 / (((fi57 >> ((2) & 63))) | 1)) {
                                            case 32767:
                                                delete(v1, (("\n\t" + "a") + "foo bar"))
                                            case -9223372036854775808:
                                                var v68 uint16 = (v48 * 255)
                                                _ = v68
                                            case -2147483648:
                                                v69 := GenS6{F0: (-(v5)), F1: 42}
                                                _ = v69
                                            }
                                            if (!((v4 << ((2) & 63)) == -1)) {
                                                switch -495 {
                                                case 128:
                                                    var v70 int64 = -1
                                                    _ = v70
                                                case -9223372036854775808:
                                                    if ((898 <= 2) && true) {
                                                        var v71 uint64 = (v51 % ((2) | 1))
                                                        var v72 int64 = -1
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        v73 := GenS7{F0: uint(v48), F1: true, F2: 264}
                                                        _ = v73
                                                    }
                                                case -664:
                                                    switch (fi57 - 0) {
                                                    case 219:
                                                        var v74 bool = true
                                                        _ = v74
                                                    default:
                                                        switch ("" + "a") {
                                                        case "hello":
                                                            var v75 int32 = 2147483647
                                                            _ = v75
                                                        case "a":
                                                            switch fi57 {
                                                            case 32768:
                                                                delete(v1, ("a" + ("foo bar" + "a")))
                                                            case 0:
                                                                v1["hello"] = 2
                                                            case -228:
                                                                if (false && (!true)) {
                                                                    v49 -= 42
                                                                    v4 = 0
                                                                    var v76 uint32 = 0
                                                                    switch 1 {
                                                                    case -2147483648:
                                                                        if (776 < 816) {
                                                                            if (65535 < 850) {
                                                                                delete(v3, "a")
                                                                                switch (fi57 >> ((4294967295) & 63)) {
                                                                                case -412:
                                                                                    var v77 uint16 = (v48 % (((^(v48))) | 1))
                                                                                    _ = v77
                                                                                }
                                                                                fmt.Println(fi57, v4)
                                                                            }
                                                                            v78 := GenS4{F0: 624, F1: v49, F2: 217}
                                                                            var v79 int32 = 1
                                                                            _ = v78
                                                                            _ = v79
                                                                        } else {
                                                                            switch fi57 {
                                                                            case -207:
                                                                                v5 -= v5
                                                                            case -197:
                                                                                if (!true) {
                                                                                    if (-233 > 0) {
                                                                                        v4 = int16(v51)
                                                                                        var v80 int16 = 32767
                                                                                        v4 += 32767
                                                                                        _ = v80
                                                                                    } else {
                                                                                        fmt.Println(fi63, v4)
                                                                                        if (v51 != (^(v51))) {
                                                                                            v49 = v49
                                                                                            var v81 uint32 = v76
                                                                                            _ = v81
                                                                                        } else {
                                                                                            v48 = (^(v48))
                                                                                            fmt.Println(v5, v4)
                                                                                        }
                                                                                    }
                                                                                    _ = len(v50)
                                                                                    if (!((1 == 9223372036854775807) || ((^(v76)) > (v76 / (((v76 >> ((2) & 63))) | 1))))) {
                                                                                        if false {
                                                                                            if len(v50) > 0 {
                                                                                                v50[0] = int(v51)
                                                                                            }
                                                                                            var v82 uint = 347
                                                                                            v49 = 170
                                                                                            _ = v82
                                                                                        } else {
                                                                                            _ = len(v50)
                                                                                            v50 = append(v50, fi57)
                                                                                            var v83 bool = (((false && false) && (493 != -1)) || false)
                                                                                            _ = v83
                                                                                        }
                                                                                        if len(v50) > 0 {
                                                                                            v50[0] = fi63
                                                                                        }
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            var v84 float64 = (3.14159 + v2)
                                                                                            _ = v84
                                                                                        case "\n\t":
                                                                                            if (!false) {
                                                                                                if len(v50) >= 2 {
                                                                                                    v50 = v50[1:len(v50)]
                                                                                                }
                                                                                                v4 -= 32767
                                                                                                var v85 int64 = 9223372036854775807
                                                                                                _ = v85
                                                                                            } else {
                                                                                                _ = len(v50)
                                                                                                _ = len(v1)
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    v86 := GenS8{F0: -1, F1: 65535, F2: (v48 >> ((65535) & 63))}
                                                                                                    _ = v86
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v5, v4, v2)
                                                                                                case "hello":
                                                                                                    switch (fi57 % (((^(fi57))) | 1)) {
                                                                                                    case 32768:
                                                                                                        v49 = v49
                                                                                                    }
                                                                                                }
                                                                                                switch fi57 {
                                                                                                case -241:
                                                                                                    if true {
                                                                                                        v50 = append(v50, 2147483647)
                                                                                                        var v87 uint = 392
                                                                                                        v2 = v2
                                                                                                        _ = v87
                                                                                                    } else {
                                                                                                        var v88 string = ""
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            v89 := GenS1{F0: 42, F1: v4}
                                                                                                            _ = v89
                                                                                                        default:
                                                                                                            v88 = v88
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v90 int16 = (v4 ^ v4)
                                                                                                            v91 := GenS3{F0: 1, F1: -124}
                                                                                                            _ = v90
                                                                                                            _ = v91
                                                                                                        }
                                                                                                        v49 = (v49 ^ v49)
                                                                                                        _ = v88
                                                                                                    }
                                                                                                }
                                                                                                var v92 int16 = (v4 * v4)
                                                                                                _ = v92
                                                                                            }
                                                                                            _ = len(v50)
                                                                                            switch ("hello" + "\n\t") {
                                                                                            case "a":
                                                                                                if false {
                                                                                                    v93 := GenS3{F0: 65535, F1: int64(v76)}
                                                                                                    _ = v93
                                                                                                }
                                                                                            case "hello":
                                                                                                v94 := v1["\n\t"]
                                                                                                _ = v94
                                                                                            case "foo bar":
                                                                                                switch 32767 {
                                                                                                case -2147483648:
                                                                                                    v48 += (v48 / (((v48 - uint16(v49))) | 1))
                                                                                                }
                                                                                            default:
                                                                                                v2 += v2
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            if len(v50) >= 2 {
                                                                                v50 = v50[1:len(v50)]
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v76
                                                                }
                                                                var v95 uint32 = 684
                                                                v51 += 348
                                                                switch fi63 {
                                                                case 32767:
                                                                    v48 = 1
                                                                default:
                                                                    v95 = v95
                                                                }
                                                                switch fi57 {
                                                                case 4294967295:
                                                                    v95 += (v95 * v95)
                                                                case 904:
                                                                    v2 -= v2
                                                                default:
                                                                    var v96 uint8 = v49
                                                                    _ = v96
                                                                }
                                                                if true {
                                                                    var v97 int16 = (v4 ^ 666)
                                                                    _ = v97
                                                                } else {
                                                                    var v98 uint16 = (v48 / ((v48) | 1))
                                                                    v98 = (v98 & v98)
                                                                    if false {
                                                                        if true {
                                                                            switch (fi57 | (fi57 + fi63)) {
                                                                            case 421:
                                                                                fmt.Println(fi57, v51)
                                                                                if ((fi57 - fi63) != (fi57 + fi63)) {
                                                                                    if false {
                                                                                        v98 -= (v48 - v98)
                                                                                        switch 9223372036854775807 {
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(fi57, v98, v95)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if (0.0 > -2.25) {
                                                                                    v50 = append(v50, fi63)
                                                                                    fmt.Println(v95, v51, v4)
                                                                                    v49 = uint8(v95)
                                                                                }
                                                                            default:
                                                                                var v99 int = -41
                                                                                _ = v99
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                v5 += v5
                                                                                var v100 float32 = float32(v2)
                                                                                _ = v100
                                                                            }
                                                                            fmt.Println(v2, v95)
                                                                            v51 -= 18446744073709551615
                                                                        }
                                                                        v49 += (v49 | (v49 - v49))
                                                                        var v101 uint64 = (v51 << ((14) & 63))
                                                                        var v102 int32 = -1
                                                                        if true {
                                                                            switch int(v95) {
                                                                            case 9223372036854775807:
                                                                                if true {
                                                                                    v51 -= (v101 << ((466) & 63))
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        if (false && ("foo bar" < "")) {
                                                                                            var v103 int64 = -9223372036854775808
                                                                                            _ = v103
                                                                                        } else {
                                                                                            if (v51 >= uint64(v49)) {
                                                                                                switch (fi57 / (((fi63 + fi63)) | 1)) {
                                                                                                case -260:
                                                                                                    var v104 float64 = (v2 / float64((1.5 * float32(v2))))
                                                                                                    v51 -= (^(v51))
                                                                                                    switch fi57 {
                                                                                                    case -937:
                                                                                                        if false {
                                                                                                            if (!(-736 == fi63)) {
                                                                                                                if true {
                                                                                                                    v49 = (v49 ^ 0)
                                                                                                                    var v105 int8 = (^(v5))
                                                                                                                    v102 += 42
                                                                                                                    _ = v105
                                                                                                                }
                                                                                                                var v106 uint8 = v49
                                                                                                                var v107 bool = (true && true)
                                                                                                                v51 = uint64(v98)
                                                                                                                v107 = v107
                                                                                                                _ = v106
                                                                                                                _ = v107
                                                                                                            } else {
                                                                                                                var v108 int32 = (^(v102))
                                                                                                                v3[""] = ""
                                                                                                                var v109 uint = 475
                                                                                                                v109 -= 660
                                                                                                                _ = v108
                                                                                                                _ = v109
                                                                                                            }
                                                                                                            var v110 uint16 = (v98 + uint16(v49))
                                                                                                            var v111 bool = true
                                                                                                            _ = v110
                                                                                                            _ = v111
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v104
                                                                                                case 32767:
                                                                                                    fmt.Println(fi57, v2, v2)
                                                                                                }
                                                                                                v5 += (v5 * (v5 % ((-128) | 1)))
                                                                                                fmt.Println(v4, v2, v98)
                                                                                                if len(v50) > 0 {
                                                                                                    v50[0] = (fi63 << ((796) & 63))
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case "":
                                                                                        var v112 float64 = ((3.14159 + float64(1e10)) + (v2 * (1e10 / ((v2 - v2) / v2))))
                                                                                        _ = v112
                                                                                    case "a":
                                                                                        v4 = (v4 ^ v4)
                                                                                        v113 := GenS7{F0: 2, F1: (0.1 <= 0.1), F2: (fi63 << ((uint(v49)) & 63))}
                                                                                        if (false || true) {
                                                                                            if ((v49 < (v49 - v49)) || (4294967295 == 1)) {
                                                                                                v101 -= v51
                                                                                                if (894 >= uint(v49)) {
                                                                                                    v114 := GenS5{F0: 42}
                                                                                                    v115 := GenS2{F0: 713, F1: 1, F2: 218}
                                                                                                    _ = v114
                                                                                                    _ = v115
                                                                                                } else {
                                                                                                    fmt.Println(v5)
                                                                                                    if true {
                                                                                                        if true {
                                                                                                            v50 = append(v50, (^(fi63)))
                                                                                                            switch "" {
                                                                                                            case "\n\t":
                                                                                                                v98 -= 0
                                                                                                            case "a":
                                                                                                                switch fi63 {
                                                                                                                case 127:
                                                                                                                    v116 := GenS9{F0: (^(v98))}
                                                                                                                    _ = v116
                                                                                                                case 0:
                                                                                                                    switch fi57 {
                                                                                                                    case 114:
                                                                                                                        var v117 uint = 65535
                                                                                                                        _ = v117
                                                                                                                    case -1:
                                                                                                                        if false {
                                                                                                                            v95 += uint32(v48)
                                                                                                                        } else {
                                                                                                                            v51 -= (v51 - (^(v51)))
                                                                                                                        }
                                                                                                                    case 42:
                                                                                                                        fmt.Println(v49, v2, v2)
                                                                                                                    }
                                                                                                                    v5 -= (v5 | v5)
                                                                                                                case -9223372036854775808:
                                                                                                                    v48 -= (v48 ^ v48)
                                                                                                                default:
                                                                                                                    var v118 float64 = (v2 - v2)
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                var v119 float32 = 1e10
                                                                                                                _ = v119
                                                                                                            }
                                                                                                            _ = len(v50)
                                                                                                        }
                                                                                                        fmt.Println(v102, v48, v5)
                                                                                                    } else {
                                                                                                        fmt.Println(v95, v101, v51)
                                                                                                        if len(v50) > 0 {
                                                                                                            v50[0] = 128
                                                                                                        }
                                                                                                        v48 = v98
                                                                                                        switch fi57 {
                                                                                                        case 4294967295:
                                                                                                            if (!true) {
                                                                                                                fmt.Println(v95)
                                                                                                                if (fi57 >= fi63) {
                                                                                                                    v120 := GenS4{F0: int(v95), F1: 1, F2: (v48 | v98)}
                                                                                                                    if (!true) {
                                                                                                                        v51 = 112
                                                                                                                        v49 = (v49 * (v49 >> ((233) & 63)))
                                                                                                                    }
                                                                                                                    fmt.Println(v5, v95)
                                                                                                                    v121 := GenS10{F0: (v95 << ((0) & 63)), F1: 969, F2: "\n\t"}
                                                                                                                    _ = v120
                                                                                                                    _ = v121
                                                                                                                }
                                                                                                                switch (fi57 % (((fi57 + (fi57 >> ((255) & 63)))) | 1)) {
                                                                                                                case 32767:
                                                                                                                    var v122 uint = 599
                                                                                                                    _ = v122
                                                                                                                case -796:
                                                                                                                    if false {
                                                                                                                        if false {
                                                                                                                            if len(v50) >= 2 {
                                                                                                                                v50 = v50[1:len(v50)]
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                var v123 uint64 = v101
                                                                                                                                v124 := GenS4{F0: fi63, F1: 2, F2: v98}
                                                                                                                                var v125 uint16 = v48
                                                                                                                                var v126 uint16 = v98
                                                                                                                                v127 := v1["\n\t"]
                                                                                                                                _ = v123
                                                                                                                                _ = v124
                                                                                                                                _ = v125
                                                                                                                                _ = v126
                                                                                                                                _ = v127
                                                                                                                            }
                                                                                                                            v128 := GenS9{F0: v98}
                                                                                                                            _ = v128
                                                                                                                        }
                                                                                                                        if true {
                                                                                                                            var v129 uint64 = (v101 % ((547) | 1))
                                                                                                                            _ = v129
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        switch (-(fi63)) {
                                                                                                                        case 1:
                                                                                                                            v130 := GenS5{F0: int16(v5)}
                                                                                                                            _ = v130
                                                                                                                        case 992:
                                                                                                                            v48 += 131
                                                                                                                        case 0:
                                                                                                                            var v131 int = (fi57 << ((2) & 63))
                                                                                                                            _ = v131
                                                                                                                        default:
                                                                                                                            var v132 float64 = (v2 / float64(0.1))
                                                                                                                            _ = v132
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        var v133 uint16 = v98
                                                                                                                        var v134 int = int(v101)
                                                                                                                        _ = v133
                                                                                                                        _ = v134
                                                                                                                    }
                                                                                                                    v95 = v95
                                                                                                                }
                                                                                                                delete(v3, "a")
                                                                                                                var v135 int32 = 42
                                                                                                                _ = v135
                                                                                                            } else {
                                                                                                                var v136 int16 = v4
                                                                                                                _ = v136
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v137 uint8 = v49
                                                                                                var v138 uint64 = 2
                                                                                                v98 = uint16(v95)
                                                                                                var v139 uint = 0
                                                                                                _ = v137
                                                                                                _ = v138
                                                                                                _ = v139
                                                                                            }
                                                                                            var v140 uint64 = 251
                                                                                            _ = v140
                                                                                        }
                                                                                        _ = v113
                                                                                    }
                                                                                    fmt.Println(v98)
                                                                                    v50 = append(v50, 42)
                                                                                    v49 += 154
                                                                                }
                                                                            case 341:
                                                                                v141 := GenS11{F0: 18446744073709551615, F1: v5}
                                                                                if (v5 > v5) {
                                                                                    fmt.Println(fi57, fi57, v49)
                                                                                }
                                                                                _ = v141
                                                                            }
                                                                            if true {
                                                                                v2 += v2
                                                                                v142 := GenS2{F0: uint(v48), F1: (v101 >> ((2) & 63)), F2: 4294967295}
                                                                                _ = v142
                                                                            }
                                                                            v143 := v3[("" + "hello")]
                                                                            _ = v143
                                                                        } else {
                                                                            if len(v50) > 0 {
                                                                                v50[0] = fi63
                                                                            }
                                                                            if true {
                                                                                v144 := GenS7{F0: 565, F1: (int16(v49) > (v4 + 32767)), F2: (fi57 ^ fi63)}
                                                                                v98 = 2
                                                                                _ = len(v50)
                                                                                _ = v144
                                                                            } else {
                                                                                v102 = v102
                                                                                v48 -= (v48 % ((v48) | 1))
                                                                                fmt.Println(v49, v4, fi57)
                                                                                var v145 float64 = (v2 / 0.1)
                                                                                _ = v145
                                                                            }
                                                                        }
                                                                        _ = v101
                                                                        _ = v102
                                                                    }
                                                                    _ = v98
                                                                }
                                                                var v146 uint16 = (v48 >> ((42) & 63))
                                                                _ = v146
                                                                _ = v95
                                                            }
                                                        case "":
                                                            v147 := GenS5{F0: v4}
                                                            _ = v147
                                                        default:
                                                            switch ("a" + "\n\t") {
                                                            case "a":
                                                                if true {
                                                                    if true {
                                                                        if len(v50) > 0 {
                                                                            v50[0] = fi63
                                                                        }
                                                                    } else {
                                                                        var v148 int8 = v5
                                                                        v51 = 42
                                                                        v149 := v3["a"]
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                    if false {
                                                                        v3[""] = ("a" + "")
                                                                    }
                                                                    fmt.Println(fi57, v51, v5)
                                                                }
                                                            default:
                                                                v150 := GenS8{F0: 42, F1: 18446744073709551615, F2: (v48 / ((255) | 1))}
                                                                _ = v150
                                                            }
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v4, v48, v5)
                                                }
                                            }
                                        }
                                    case "a":
                                        _ = len(v3)
                                    }
                                case 312:
                                    if false {
                                        v48 -= v48
                                        v2 -= 0.1
                                        var v151 int = int(v5)
                                        _ = len(v50)
                                        v152 := GenS6{F0: -1, F1: 65535}
                                        v153 := v3[("a" + ("hello" + ""))]
                                        var v154 int16 = v4
                                        _ = v151
                                        _ = v152
                                        _ = v153
                                        _ = v154
                                    } else {
                                        v51 += 114
                                        if true {
                                            fmt.Println(fi63, v2, fi57)
                                            var v155 float64 = float64(-2.25)
                                            if (0 != 65535) {
                                                if ((!false) && true) {
                                                    v4 -= (-(v4))
                                                    var v156 int8 = v5
                                                    var v157 float64 = float64(1e10)
                                                    v157 += float64(0.1)
                                                    switch 32767 {
                                                    case 32768:
                                                        if false {
                                                            v158 := v1["a"]
                                                            if true {
                                                                if len(v50) >= 2 {
                                                                    v50 = v50[1:len(v50)]
                                                                }
                                                                v3["a"] = "hello"
                                                                var v159 uint16 = v48
                                                                var v160 float32 = 0.0
                                                                v51 -= 1
                                                                if true {
                                                                    v160 = (v160 - 1.5)
                                                                    var v161 bool = (v5 < v5)
                                                                    _ = v161
                                                                }
                                                                _ = v159
                                                                _ = v160
                                                            } else {
                                                                if ((-576 <= 42) || true) {
                                                                    v51 = (v51 + (v51 << ((1) & 63)))
                                                                    switch "" {
                                                                    case "hello":
                                                                        var v162 float32 = float32((float64((float32(((((((v155 - 1.5) + v157) * -2.25) / 0.1) + v2) * ((v2 / (v155 / float64(3.14159))) / (3.14159 / v155)))) - float32(v2))) * v2))
                                                                        var v163 string = ("\n\t" + "foo bar")
                                                                        v164 := GenS4{F0: 9223372036854775807, F1: 42, F2: (v48 % (((v48 + 441)) | 1))}
                                                                        _ = v162
                                                                        _ = v163
                                                                        _ = v164
                                                                    case "foo bar":
                                                                        v48 -= (v48 << ((12) & 63))
                                                                    }
                                                                }
                                                            }
                                                            if ((-(v4)) < (-(v4))) {
                                                                v165 := GenS1{F0: 65535, F1: (v4 + (-(v4)))}
                                                                var v166 int64 = -1
                                                                v49 = v49
                                                                v167 := GenS8{F0: (v166 >> ((1) & 63)), F1: 18446744073709551615, F2: (v48 / ((v48) | 1))}
                                                                if false {
                                                                    v51 = (v51 | v51)
                                                                    var v168 uint32 = v158
                                                                    if (v155 == (v2 - 1.5)) {
                                                                        v50 = append(v50, (fi63 % (((fi57 ^ fi63)) | 1)))
                                                                        switch fi57 {
                                                                        case -1:
                                                                            var v169 int16 = v4
                                                                            _ = v169
                                                                        }
                                                                        v156 = v5
                                                                    } else {
                                                                        var v170 bool = false
                                                                        _ = v170
                                                                    }
                                                                    if (true && (845 <= (v166 >> ((666) & 63)))) {
                                                                        if (!((1.5 <= -2.25) || true)) {
                                                                            v48 = 1
                                                                        }
                                                                    }
                                                                    _ = v168
                                                                }
                                                                v171 := GenS1{F0: 4, F1: 1000}
                                                                if (0 < (v166 + v166)) {
                                                                    switch fi63 {
                                                                    case 419:
                                                                        v155 = v2
                                                                    case -2147483648:
                                                                        v172 := GenS10{F0: v158, F1: 2, F2: ""}
                                                                        _ = v172
                                                                    case -9223372036854775808:
                                                                        fmt.Println(v51)
                                                                    default:
                                                                        var v173 bool = true
                                                                        _ = v173
                                                                    }
                                                                    var v174 uint32 = v158
                                                                    _ = v174
                                                                }
                                                                _ = v165
                                                                _ = v166
                                                                _ = v167
                                                                _ = v171
                                                            }
                                                            _ = v158
                                                        } else {
                                                            v157 += v2
                                                            var v175 uint = 198
                                                            _ = len(v50)
                                                            v157 -= -2.25
                                                            _ = v175
                                                        }
                                                    }
                                                    _ = v156
                                                    _ = v157
                                                }
                                                if (494 < 65535) {
                                                    v51 = (v51 * v51)
                                                    v176 := GenS2{F0: 2, F1: v51, F2: 16}
                                                    _ = v176
                                                }
                                                _ = len(v50)
                                            } else {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch ("\n\t" + "a") {
                                                    case "foo bar":
                                                        v48 -= 2
                                                    case "a":
                                                        var v177 uint32 = uint32(v48)
                                                        _ = v177
                                                    default:
                                                        var v178 float32 = -2.25
                                                        _ = v178
                                                    }
                                                case "a":
                                                    v48 = uint16(v49)
                                                default:
                                                    v4 = v4
                                                }
                                            }
                                            _ = v155
                                        } else {
                                            v49 -= 136
                                        }
                                    }
                                }
                            }
                            if len(v50) > 0 {
                                v50[0] = fi63
                            }
                            _ = v64
                        }
                        _ = fi63
                    }
                    fmt.Println(v5)
                    var v179 string = ""
                    _ = v179
                } else {
                    if (614 <= 397) {
                        v48 = 822
                        v5 = -16
                    } else {
                        v49 += 255
                        switch "foo bar" {
                        case "":
                            if true {
                                var v180 int8 = -99
                                var v181 float64 = v2
                                _ = v180
                                _ = v181
                            }
                        case "a":
                            if (true || true) {
                                var v182 float64 = ((3.14159 / (float64(3.14159) * v2)) * (v2 / 1e10))
                                _ = v182
                            }
                        case "\n\t":
                            v5 = -1
                        default:
                            if false {
                                if ((v4 ^ v4) != int16(fi57)) {
                                    v183 := GenS3{F0: 518, F1: int64(v5)}
                                    fmt.Println(v51, v51, v5)
                                    _ = v183
                                } else {
                                    for fi184 := 0; fi184 < 16; fi184++ {
                                        _ = len(v50)
                                        _ = fi184
                                    }
                                    if false {
                                        fmt.Println(fi57, v5)
                                    }
                                }
                            } else {
                                var v185 int32 = -479
                                v186 := GenS2{F0: 42, F1: v51, F2: 931}
                                delete(v3, "")
                                _ = v185
                                _ = v186
                            }
                        }
                        var v187 uint64 = v51
                        _ = v187
                    }
                    for fi188 := 0; fi188 < 6; fi188++ {
                        v2 += 1.5
                        var v189 int16 = (v4 * (-(v4)))
                        v190 := GenS7{F0: 98, F1: false, F2: fi188}
                        if false {
                            var v191 string = "hello"
                            v192 := GenS10{F0: 2, F1: (v51 ^ 4294967295), F2: ((v191 + ("foo bar" + ("foo bar" + v191))) + v191)}
                            fmt.Println(v51)
                            _ = v191
                            _ = v192
                        } else {
                            fmt.Println(v5)
                        }
                        _ = fi188
                        _ = v189
                        _ = v190
                    }
                }
                _ = fi57
            }
        }
    case 127:
        for fi193 := 0; fi193 < 2; fi193++ {
            var v194 uint8 = (v49 ^ 255)
            var v195 float32 = -2.25
            _ = fi193
            _ = v194
            _ = v195
        }
    default:
        v51 = 0
    }
    for fi196 := 0; fi196 < 18; fi196++ {
        v5 += (v5 ^ v5)
        _ = fi196
    }
    v48 = 0
    if (!true) {
        if ("hello" == ("foo bar" + "foo bar")) {
            if ((("foo bar" + "") == "foo bar") && true) {
                v51 = (v51 & (v51 / ((0) | 1)))
            } else {
                var v197 int = 0
                _ = v197
            }
            if (false || ((205 != 1) || false)) {
                if ((v48 - (v48 + 32767)) >= (v48 / ((32767) | 1))) {
                    delete(v1, ("a" + ""))
                    for fi198 := 0; fi198 < 8; fi198++ {
                        v48 = (^(v48))
                        if false {
                            if true {
                                for fi199 := 0; fi199 < 5; fi199++ {
                                    var v200 int32 = -2147483648
                                    _ = fi199
                                    _ = v200
                                }
                            }
                            switch -981 {
                            case 127:
                                fmt.Println(v4, v5)
                            case 42:
                                var v201 bool = (!(((-(v4)) < 736) || true))
                                _ = v201
                            case 569:
                                var v202 int32 = 17
                                _ = v202
                            }
                        } else {
                            if false {
                                if true {
                                    v48 = uint16(v51)
                                    v203 := GenS9{F0: (v48 / ((v48) | 1))}
                                    _ = v203
                                }
                                v4 -= (v4 | 42)
                                var v204 float64 = float64(3.14159)
                                _ = v204
                            }
                        }
                        _ = fi198
                    }
                    switch 4294967295 {
                    case 745:
                        var v205 int64 = 1
                        _ = v205
                    }
                }
            }
        } else {
            if false {
                for fi206 := 0; fi206 < 5; fi206++ {
                    fmt.Println(v2, v5)
                    var v207 uint = 4294967295
                    _ = fi206
                    _ = v207
                }
                var v208 uint64 = (v51 ^ 4294967295)
                _ = v208
            } else {
                v209 := GenS12{F0: v48, F1: v51, F2: (v49 >> ((42) & 63))}
                v48 -= (v48 << ((2) & 63))
                _ = v209
            }
            for fi210 := 0; fi210 < 13; fi210++ {
                for fi211 := 0; fi211 < 8; fi211++ {
                    var v212 string = ("\n\t" + "")
                    var v213 string = (v212 + v212)
                    fmt.Println(v2, v51, fi210)
                    _ = fi211
                    _ = v212
                    _ = v213
                }
                if true {
                    v214 := GenS13{F0: 1}
                    var v215 float32 = 0.0
                    _ = v214
                    _ = v215
                }
                fmt.Println(fi210, v49, v51)
                _ = fi210
            }
            fmt.Println(v48, v5, v49)
        }
        for fi216 := 0; fi216 < 18; fi216++ {
            for fi217 := 0; fi217 < 19; fi217++ {
                if (!(325 == 1)) {
                    var v218 uint = 18446744073709551615
                    var v219 uint16 = v48
                    _ = v218
                    _ = v219
                }
                var v220 uint8 = (v49 - (v49 + 42))
                _ = fi217
                _ = v220
            }
            switch "a" {
            case "":
                if false {
                    fmt.Println(fi216)
                    v221 := GenS1{F0: uint(v49), F1: v4}
                    for fi222 := 0; fi222 < 13; fi222++ {
                        if true {
                            var v223 uint32 = 2147483648
                            _ = v223
                        }
                        _ = len(v50)
                        delete(v1, "foo bar")
                        var v224 float64 = v2
                        if (-1 > v5) {
                            switch (fi222 | (-(fi216))) {
                            case 32768:
                                var v225 int16 = v4
                                _ = v225
                            case -1:
                                v50 = append(v50, fi222)
                            case 447:
                                if true {
                                    if ((true && (true || false)) && false) {
                                        v5 = v5
                                        if ((!((v48 ^ 184) <= (v48 << ((255) & 63)))) && true) {
                                            if true {
                                                v2 = 0.1
                                                v224 = (3.14159 * 1e10)
                                                fmt.Println(v4)
                                            } else {
                                                switch "hello" {
                                                case "":
                                                    var v226 uint = 71
                                                    _ = v226
                                                case "hello":
                                                    v227 := GenS14{F0: false, F1: v48, F2: -1}
                                                    _ = v227
                                                default:
                                                    var v228 bool = true
                                                    _ = v228
                                                }
                                            }
                                            if (42 == 1) {
                                                var v229 float64 = v2
                                                _ = len(v3)
                                                fmt.Println(fi222, v5, fi222)
                                                _ = v229
                                            } else {
                                                v230 := GenS2{F0: 0, F1: v51, F2: 504}
                                                _ = len(v1)
                                                v2 = float64(1.5)
                                                _ = len(v3)
                                                var v231 bool = (!(v5 > v5))
                                                switch fi222 {
                                                case 2147483647:
                                                    fmt.Println(v49, v231)
                                                case 128:
                                                    v5 -= int8(v4)
                                                }
                                                fmt.Println(v49)
                                                _ = v230
                                                _ = v231
                                            }
                                            v232 := GenS3{F0: 364, F1: 9223372036854775807}
                                            _ = v232
                                        } else {
                                            v233 := GenS11{F0: 0, F1: (v5 ^ int8(fi222))}
                                            v234 := GenS5{F0: -242}
                                            _ = v233
                                            _ = v234
                                        }
                                        v235 := v3[("foo bar" + "a")]
                                        var v236 uint64 = (^(v51))
                                        _ = v235
                                        _ = v236
                                    }
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        var v237 float32 = (1e10 * 3.14159)
                                        _ = v237
                                    case "":
                                        v49 = v49
                                    case "hello":
                                        v224 -= float64((1e10 - 1e10))
                                    }
                                } else {
                                    var v238 int = (-(fi216))
                                    _ = v238
                                }
                            }
                            switch int(v51) {
                            case 719:
                                v239 := GenS15{F0: (^(v51)), F1: "hello"}
                                _ = v239
                            default:
                                v224 = 0.0
                            }
                            switch (fi216 - 4294967295) {
                            case 128:
                                v240 := GenS16{F0: false, F1: v4}
                                _ = v240
                            case -205:
                                fmt.Println(fi216, fi216, v224)
                            case 0:
                                v4 = v4
                            default:
                                switch 981 {
                                case 32768:
                                    fmt.Println(v49, v224, fi216)
                                    if (!false) {
                                        var v241 bool = (v224 == -2.25)
                                        _ = v241
                                    }
                                case -2147483648:
                                    v242 := GenS17{F0: (v48 % ((694) | 1)), F1: 2, F2: 48}
                                    _ = v242
                                default:
                                    if false {
                                        v4 = 13
                                        fmt.Println(v49, v48, fi216)
                                    }
                                }
                            }
                            switch "a" {
                            case "\n\t":
                                v48 = v48
                            case "":
                                v2 -= 3.14159
                            default:
                                var v243 uint16 = (v48 & 0)
                                v244 := GenS18{F0: 1, F1: ""}
                                var v245 int32 = 2147483647
                                if false {
                                    v246 := GenS19{F0: (v245 / (((^(v245))) | 1)), F1: 849, F2: v243}
                                    fmt.Println(fi216, v4, v5)
                                    _ = v246
                                }
                                _ = v243
                                _ = v244
                                _ = v245
                            }
                        } else {
                            switch "hello" {
                            case "a":
                                var v247 int16 = (v4 + v4)
                                _ = v247
                            case "foo bar":
                                v248 := GenS20{F0: uint8(v48), F1: v51, F2: (v48 * (v48 + v48))}
                                _ = v248
                            case "\n\t":
                                v4 = v4
                            default:
                                var v249 int = -2147483648
                                _ = v249
                            }
                            v224 += v2
                            if false {
                                v250 := GenS21{F0: v51}
                                _ = v250
                            } else {
                                v224 = v2
                                if false {
                                    switch fi216 {
                                    case 42:
                                        v49 -= (v49 * 128)
                                    }
                                } else {
                                    if false {
                                        if len(v50) >= 2 {
                                            v50 = v50[1:len(v50)]
                                        }
                                    }
                                }
                                var v251 int8 = 1
                                _ = v251
                            }
                        }
                        var v252 uint32 = 4294967295
                        v4 = 32767
                        _ = fi222
                        _ = v224
                        _ = v252
                    }
                    _ = v221
                }
            }
            _ = fi216
        }
    } else {
        for fi253 := 0; fi253 < 1; fi253++ {
            v2 += 3.14159
            _ = fi253
        }
        switch "" {
        case "":
            fmt.Println(v49, v5)
        case "foo bar":
            if (!(1.5 == float32(v2))) {
                switch "foo bar" {
                case "hello":
                    for fi254 := 0; fi254 < 10; fi254++ {
                        var v255 int8 = (^(v5))
                        if len(v50) > 0 {
                            v50[0] = 302
                        }
                        _ = fi254
                        _ = v255
                    }
                default:
                    for fi256 := 0; fi256 < 7; fi256++ {
                        v48 = (^(v48))
                        fmt.Println(v49, v48, fi256)
                        v2 += v2
                        _ = fi256
                    }
                }
                switch "" {
                case "hello":
                    v48 += (v48 % ((v48) | 1))
                case "a":
                    delete(v1, ("a" + ("foo bar" + "a")))
                }
            }
        case "a":
            v2 -= v2
        default:
            v257 := GenS1{F0: 0, F1: v4}
            _ = v257
        }
        if false {
            var v258 uint8 = (v49 & (v49 % (((v49 & v49)) | 1)))
            _ = v258
        }
    }
    var v259 uint16 = (v48 - (v48 >> ((42) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v259:", v259)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v19
    _ = v2
    _ = v259
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
