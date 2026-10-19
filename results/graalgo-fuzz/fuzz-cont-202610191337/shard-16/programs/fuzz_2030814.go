package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 string }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }
type GenS4 struct { F0 uint16; F1 string; F2 uint64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int8; F1 uint64; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 bool; F1 int64; F2 uint64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32; F1 int64 }
type GenS13 struct { F0 int16; F1 uint8; F2 int32 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint64; F1 int8 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int32; F1 int16 }
type GenS19 struct { F0 bool; F1 int32 }
type GenS20 struct { F0 int32; F1 uint }
type GenS21 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 int8, p1 int8) int16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        if ((false && (!true)) || (!((-2147483648 >= 1) && true))) {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 string = "foo bar"
                _ = fi8
                _ = v9
            }
        }
        _ = fi7
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 830
}
func genFn1() int64 {
    v10 := GenS1{F0: false, F1: -9223372036854775808, F2: "foo bar"}
    switch "hello" {
    case "a":
        for fi11 := 0; fi11 < 6; fi11++ {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(fi11)
            case "\n\t":
                if false {
                    var v12 int = 127
                    for fi13 := 0; fi13 < 6; fi13++ {
                        fmt.Println(v12, fi11)
                        _ = fi13
                    }
                    for fi14 := 0; fi14 < 4; fi14++ {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(fi14, fi14)
                        case "foo bar":
                            var v15 int8 = -128
                            v16 := GenS2{F0: 1, F1: 18446744073709551615}
                            v17 := GenS2{F0: 791, F1: 1}
                            v15 = v15
                            if (1e10 < 3.14159) {
                                v15 -= 127
                            }
                            _ = v15
                            _ = v16
                            _ = v17
                        case "\n\t":
                            v12 -= 302
                        }
                        v12 += 42
                        _ = fi14
                    }
                    _ = v12
                }
            }
            _ = fi11
        }
    case "":
        v18 := GenS2{F0: 1, F1: 1}
        _ = v18
    case "hello":
        if (!(true || true)) {
            switch "hello" {
            case "hello":
                var v19 int64 = -634
                _ = v19
            case "\n\t":
                switch 0 {
                case 32768:
                    for fi20 := 0; fi20 < 19; fi20++ {
                        var v21 int64 = -885
                        _ = fi20
                        _ = v21
                    }
                }
            default:
                for fi22 := 0; fi22 < 4; fi22++ {
                    fmt.Println(fi22)
                    if (255 <= 889) {
                        var v23 uint32 = 627
                        if true {
                            if true {
                                var v24 string = "\n\t"
                                switch v24 {
                                case "a":
                                    if false {
                                        fmt.Println(fi22)
                                        v25 := []int{(fi22 - -232), (fi22 / (((fi22 | fi22)) | 1))}
                                        v23 = (v23 / ((v23) | 1))
                                        _ = v25
                                    } else {
                                        fmt.Println(v24)
                                        if ((fi22 << ((1) & 63)) < (fi22 / (((fi22 / ((fi22) | 1))) | 1))) {
                                            switch fi22 {
                                            case -9223372036854775808:
                                                v23 -= v23
                                            case 0:
                                                v23 -= (v23 >> ((643) & 63))
                                            default:
                                                v26 := []int{(fi22 * 2147483647)}
                                                _ = v26
                                            }
                                            v23 -= (v23 & (v23 | v23))
                                            for fi27 := 0; fi27 < 12; fi27++ {
                                                v28 := []int{}
                                                if ((((float32(float64(0.1)) + 3.14159) + (0.0 - 3.14159)) * 0.0) == 1e10) {
                                                    if false {
                                                        var v29 uint64 = 2
                                                        fmt.Println(fi22, v23, fi22)
                                                        if (42 <= 100) {
                                                            if len(v28) >= 2 {
                                                                v28 = v28[1:len(v28)]
                                                            }
                                                            fmt.Println(fi27)
                                                            if ((32767 > -32768) || (true && true)) {
                                                                switch 32768 {
                                                                case 32767:
                                                                    fmt.Println(fi27, fi22)
                                                                case -1:
                                                                    var v30 string = v24
                                                                    _ = v30
                                                                }
                                                                var v31 uint32 = 363
                                                                v28 = append(v28, (fi27 << ((42) & 63)))
                                                                _ = v31
                                                            } else {
                                                                _ = len(v28)
                                                                v29 -= 1
                                                                _ = len(v28)
                                                                v24 = (("hello" + v24) + "a")
                                                            }
                                                        }
                                                        _ = v29
                                                    }
                                                }
                                                _ = fi27
                                                _ = v28
                                            }
                                        } else {
                                            v23 = v23
                                            switch v24 {
                                            case "foo bar":
                                                var v32 int = fi22
                                                _ = v32
                                            }
                                            var v33 uint16 = 230
                                            _ = v33
                                        }
                                        var v34 int32 = -2147483648
                                        _ = v34
                                    }
                                case "\n\t":
                                    v23 = (v23 * (^(v23)))
                                case "":
                                    switch (fi22 ^ 756) {
                                    case -2147483648:
                                        var v35 uint64 = 565
                                        _ = v35
                                    default:
                                        for fi36 := 0; fi36 < 6; fi36++ {
                                            var v37 uint = 582
                                            v23 = v23
                                            _ = fi36
                                            _ = v37
                                        }
                                    }
                                }
                                switch (-(fi22)) {
                                case -1:
                                    fmt.Println(v24)
                                }
                                _ = v24
                            } else {
                                if (true || false) {
                                    v38 := []int{fi22, -1}
                                    _ = v38
                                }
                                var v39 float32 = ((float32(3.14159) * (0.0 / 3.14159)) / ((((float32(0.0) + float32((0.1 - (3.14159 * (0.0 / 1.5))))) + 0.1) - float32(((float64(0.1) * -2.25) + float64(((0.0 / 1.5) * (float32(-2.25) + float32(1.5))))))) * float32(1.5)))
                                _ = v39
                            }
                            fmt.Println(v23, v23)
                            v23 = 327
                        }
                        var v40 bool = true
                        _ = v23
                        _ = v40
                    } else {
                        v41 := GenS1{F0: false, F1: fi22, F2: "foo bar"}
                        var v42 int16 = 1
                        _ = v41
                        _ = v42
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v43 uint32 = 42
                        _ = v43
                    case "a":
                        v44 := GenS1{F0: true, F1: fi22, F2: ("hello" + ("hello" + "a"))}
                        _ = v44
                    default:
                        if false {
                            v45 := GenS1{F0: false, F1: fi22, F2: "a"}
                            v46 := []int{}
                            _ = v45
                            _ = v46
                        }
                    }
                    _ = fi22
                }
            }
            if true {
                for fi47 := 0; fi47 < 15; fi47++ {
                    if (false || false) {
                        switch (fi47 << ((0) & 63)) {
                        case 1:
                            var v48 string = "\n\t"
                            _ = v48
                        }
                    } else {
                        for fi49 := 0; fi49 < 4; fi49++ {
                            var v50 float32 = (0.0 + 1e10)
                            v51 := []int{}
                            v52 := GenS2{F0: 441, F1: 849}
                            _ = fi49
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                        var v53 int32 = -250
                        _ = v53
                    }
                    var v54 uint16 = 700
                    for fi55 := 0; fi55 < 20; fi55++ {
                        v54 = 305
                        v56 := GenS2{F0: 179, F1: 42}
                        _ = fi55
                        _ = v56
                    }
                    _ = fi47
                    _ = v54
                }
            } else {
                var v57 int8 = -1
                v58 := GenS3{F0: 65535, F1: 1, F2: false}
                var v59 bool = false
                _ = v57
                _ = v58
                _ = v59
            }
            v60 := GenS3{F0: 42, F1: -211, F2: (!false)}
            _ = v60
        }
    default:
        var v61 int8 = 42
        _ = v61
    }
    if true {
        for fi62 := 0; fi62 < 10; fi62++ {
            fmt.Println(fi62)
            v63 := GenS3{F0: 65535, F1: -837, F2: ((502 > 655) && true)}
            _ = fi62
            _ = v63
        }
        for fi64 := 0; fi64 < 13; fi64++ {
            var v65 uint32 = 2
            _ = fi64
            _ = v65
        }
        v66 := GenS4{F0: 2, F1: ((((("a" + "foo bar") + "\n\t") + "a") + "\n\t") + "hello"), F2: 2}
        _ = v66
    } else {
        var v67 int32 = 42
        v68 := GenS5{F0: 1, F1: (^(v67))}
        _ = v67
        _ = v68
    }
    if (430 > 0) {
        for fi69 := 0; fi69 < 5; fi69++ {
            fmt.Println(fi69)
            _ = fi69
        }
        var v70 string = ("foo bar" + "a")
        _ = v70
    } else {
        switch "a" {
        case "foo bar":
            for fi71 := 0; fi71 < 10; fi71++ {
                v72 := GenS6{F0: 0, F1: 42}
                _ = fi71
                _ = v72
            }
        case "\n\t":
            if (false && ((true && false) && (!((1 >= -1) || true)))) {
                v73 := GenS5{F0: -586, F1: 254}
                _ = v73
            }
        default:
            v74 := GenS2{F0: 42, F1: 577}
            _ = v74
        }
    }
    _ = v10
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 0
    var v2 float32 = 0.1
    var v3 int16 = 1
    v4 := make(map[uint32]string)
    var v5 uint32 = 0
    v6 := make(map[int]int)
    switch (("foo bar" + "foo bar") + "a") {
    case "foo bar":
        var v75 int64 = 0
        _ = v75
    default:
        v1 -= (v1 | (v1 & v1))
    }
    var v76 int8 = 1
    v2 = v2
    v77 := v6[-1]
    v78 := []int{v77, 382}
    var v79 uint64 = 991
    for fi80 := 0; fi80 < 10; fi80++ {
        for fi81 := 0; fi81 < 3; fi81++ {
            var v82 float64 = (-2.25 * (0.1 - 3.14159))
            fmt.Println(v3, v82, v2)
            switch (("a" + ("foo bar" + "")) + "foo bar") {
            case "foo bar":
                v79 += v79
            case "":
                v77 = fi81
            case "\n\t":
                v77 -= 9223372036854775807
            default:
                var v83 int8 = -1
                _ = v83
            }
            _ = fi81
            _ = v82
        }
        if false {
            v76 += (v76 << ((65535) & 63))
        } else {
            if (!("\n\t" >= ("foo bar" + "hello"))) {
                v84 := v4[(v5 ^ (v5 ^ v5))]
                _ = v84
            } else {
                for fi85 := 0; fi85 < 8; fi85++ {
                    if false {
                        v86 := GenS7{F0: (v76 >> ((18446744073709551615) & 63)), F1: 4294967295, F2: fi85}
                        _ = v86
                    } else {
                        var v87 int64 = 0
                        _ = v87
                    }
                    var v88 int = (fi80 % (((^(v77))) | 1))
                    _ = fi85
                    _ = v88
                }
            }
        }
        _ = fi80
    }
    v4[(v5 % ((394) | 1))] = "foo bar"
    if true {
        v76 += v76
    }
    if (18446744073709551615 < 1) {
        v89 := GenS1{F0: false, F1: v77, F2: "hello"}
        v1 += v1
        _ = v89
    }
    switch (v77 + v77) {
    case 127:
        for fi90 := 0; fi90 < 3; fi90++ {
            v1 = v1
            v77 += 2147483647
            if len(v78) >= 2 {
                v78 = v78[1:len(v78)]
            }
            _ = fi90
        }
    }
    v4[(v5 | 0)] = ((("foo bar" + "") + ("a" + ("" + ("hello" + "foo bar")))) + "")
    var v91 float32 = 0.1
    if (776 != 65535) {
        v79 = (v79 | (v79 | (v79 >> ((358) & 63))))
    } else {
        switch (v77 / (((v77 ^ 42)) | 1)) {
        case 769:
            if false {
                v1 = 2
                var v92 bool = (false || ((v77 == (v77 * int(v1))) && true))
                if (!(v92 && (339 < 2))) {
                    if len(v78) > 0 {
                        v78[0] = (v77 - (-(v77)))
                    }
                    if (2 != 4294967295) {
                        fmt.Println(v3, v5)
                    }
                } else {
                    if v92 {
                        v93 := GenS5{F0: 0, F1: -2147483648}
                        if true {
                            v94 := GenS8{F0: v1}
                            _ = v94
                        }
                        _ = v93
                    } else {
                        v3 = (v3 << ((0) & 63))
                        var v95 int8 = v76
                        v92 = ((("a" + "foo bar") == "hello") || true)
                        _ = v95
                    }
                    v5 += (v5 & v5)
                    if true {
                        _ = len(v78)
                        if len(v78) > 0 {
                            v78[0] = (^(v77))
                        }
                    } else {
                        var v96 float64 = 0.0
                        if ((v77 - v77) <= v77) {
                            switch ("hello" + "foo bar") {
                            case "\n\t":
                                var v97 int64 = -9223372036854775808
                                _ = v97
                            case "":
                                switch "\n\t" {
                                case "foo bar":
                                    v92 = v92
                                default:
                                    var v98 string = "\n\t"
                                    _ = v98
                                }
                            case "a":
                                if false {
                                    switch 0 {
                                    case 42:
                                        v99 := GenS7{F0: (v76 ^ v76), F1: (v79 | 419), F2: v77}
                                        _ = v99
                                    case -2147483648:
                                        var v100 uint = 42
                                        _ = v100
                                    case -12:
                                        switch v77 {
                                        case -2147483648:
                                            fmt.Println(v79)
                                        case 4294967295:
                                            if true {
                                                v79 -= uint64(v1)
                                            } else {
                                                v96 -= (v96 / v96)
                                                if false {
                                                    if (v79 < 689) {
                                                        delete(v4, 97)
                                                    } else {
                                                        if ((^(v79)) <= (^(v79))) {
                                                            for fi101 := 0; fi101 < 7; fi101++ {
                                                                if (true || (1.5 == 1.5)) {
                                                                    v5 -= 65535
                                                                    delete(v4, 2147483647)
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        for fi102 := 0; fi102 < 1; fi102++ {
                                                                            fmt.Println(v5, v77, fi101)
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                var v103 uint32 = (^(v5))
                                                                                _ = v103
                                                                            }
                                                                            v78 = append(v78, 32768)
                                                                            if ((^(v5)) < 2147483648) {
                                                                                if (v96 != v96) {
                                                                                    if v92 {
                                                                                        if v92 {
                                                                                            v4[(v5 & v5)] = ("\n\t" + "foo bar")
                                                                                            fmt.Println(v92, fi102, v92)
                                                                                            var v104 bool = (0 > 1)
                                                                                            _ = v104
                                                                                        }
                                                                                    } else {
                                                                                        if false {
                                                                                            var v105 uint64 = 1
                                                                                            var v106 uint = 65535
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                        fmt.Println(v92)
                                                                                    }
                                                                                } else {
                                                                                    v91 += float32((v96 - (v96 + v96)))
                                                                                    fmt.Println(v79, v92, v79)
                                                                                }
                                                                                v77 = 127
                                                                                v107 := GenS9{F0: 0}
                                                                                _ = v107
                                                                            } else {
                                                                                v78 = append(v78, -282)
                                                                                var v108 uint = 18446744073709551615
                                                                                var v109 int16 = (-(v3))
                                                                                if v92 {
                                                                                    fmt.Println(v96, v5, v2)
                                                                                }
                                                                                if (537 == 730) {
                                                                                    var v110 string = "a"
                                                                                    _ = v110
                                                                                } else {
                                                                                    v5 -= v5
                                                                                    var v111 float32 = (0.1 / (v91 + (3.14159 + (-2.25 * v91))))
                                                                                    v4[(^(v5))] = ("foo bar" + "")
                                                                                    _ = v111
                                                                                }
                                                                                v91 = float32(((((v96 - ((v96 / (v96 + v96)) / 1e10)) * (v96 / v96)) * v96) + v96))
                                                                                v79 = (v79 ^ (v79 ^ v79))
                                                                                _ = v108
                                                                                _ = v109
                                                                            }
                                                                            if (!(195 != -9223372036854775808)) {
                                                                                if false {
                                                                                    fmt.Println(v96, v76, v76)
                                                                                } else {
                                                                                    var v112 bool = v92
                                                                                    _ = v112
                                                                                }
                                                                                if v92 {
                                                                                    switch (fi101 | 0) {
                                                                                    case -584:
                                                                                        if (uint(v5) > uint(v5)) {
                                                                                            var v113 uint32 = v5
                                                                                            switch -9223372036854775808 {
                                                                                            case 1:
                                                                                                if false {
                                                                                                    if false {
                                                                                                        _ = len(v6)
                                                                                                        switch fi102 {
                                                                                                        case -2147483648:
                                                                                                            if true {
                                                                                                                v113 = v5
                                                                                                                switch fi102 {
                                                                                                                case -744:
                                                                                                                    delete(v6, fi101)
                                                                                                                case 42:
                                                                                                                    v114 := GenS2{F0: 65535, F1: 928}
                                                                                                                    _ = v114
                                                                                                                case 0:
                                                                                                                    switch (fi101 >> ((0) & 63)) {
                                                                                                                    case -786:
                                                                                                                        if v92 {
                                                                                                                            v115 := GenS4{F0: 65535, F1: ("hello" + "a"), F2: (v79 * (v79 / ((v79) | 1)))}
                                                                                                                            var v116 uint8 = (v1 & (^(v1)))
                                                                                                                            v1 = (v116 / (((^(v116))) | 1))
                                                                                                                            _ = v115
                                                                                                                            _ = v116
                                                                                                                        }
                                                                                                                    case 2147483647:
                                                                                                                        if false {
                                                                                                                            var v117 int32 = 0
                                                                                                                            switch (("a" + "\n\t") + "\n\t") {
                                                                                                                            case "hello":
                                                                                                                                fmt.Println(v1, v5, v113)
                                                                                                                            case "foo bar":
                                                                                                                                v118 := GenS4{F0: 563, F1: ("hello" + ("a" + ("" + ("a" + "hello")))), F2: (v79 - 1)}
                                                                                                                                _ = v118
                                                                                                                            case "a":
                                                                                                                                switch v77 {
                                                                                                                                case -1:
                                                                                                                                    switch "foo bar" {
                                                                                                                                    case "hello":
                                                                                                                                        var v119 uint32 = 0
                                                                                                                                        _ = v119
                                                                                                                                    case "a":
                                                                                                                                        if true {
                                                                                                                                            if v92 {
                                                                                                                                                var v120 uint = 245
                                                                                                                                                v121 := GenS4{F0: uint16(v113), F1: "foo bar", F2: 346}
                                                                                                                                                v96 += ((0.1 / (float64((v91 - (v91 * v2))) * ((float64(0.0) + float64(v91)) + (((((-2.25 * (v96 / (0.1 + v96))) - (float64(((v2 / (v2 - v2)) * (v2 - float32(((v96 / (-2.25 + 0.1)) + v96))))) / (-2.25 * (v96 * v96)))) / 1.0) + 1e10) / v96)))) - ((v96 / (1.5 + (v96 + 0.0))) - 1e10))
                                                                                                                                                _ = v120
                                                                                                                                                _ = v121
                                                                                                                                            }
                                                                                                                                            v117 = (v117 * (v117 & (-(v117))))
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v2 += 1e10
                                                                                                                            var v122 int8 = v76
                                                                                                                            fmt.Println(v122, v1)
                                                                                                                            if v92 {
                                                                                                                                v77 = v77
                                                                                                                                var v123 bool = true
                                                                                                                                _ = v123
                                                                                                                            } else {
                                                                                                                                v124 := GenS2{F0: 323, F1: v79}
                                                                                                                                switch (("\n\t" + "\n\t") + (("" + "a") + "foo bar")) {
                                                                                                                                case "":
                                                                                                                                    var v125 int16 = (v3 & int16(v122))
                                                                                                                                    _ = v125
                                                                                                                                case "hello":
                                                                                                                                    var v126 uint16 = 714
                                                                                                                                    _ = v126
                                                                                                                                case "a":
                                                                                                                                    if v92 {
                                                                                                                                        fmt.Println(v1, v79)
                                                                                                                                        v76 = 127
                                                                                                                                        v127 := GenS9{F0: 925}
                                                                                                                                        _ = v127
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                v128 := v4[v5]
                                                                                                                                _ = v124
                                                                                                                                _ = v128
                                                                                                                            }
                                                                                                                            _ = v117
                                                                                                                            _ = v122
                                                                                                                        } else {
                                                                                                                            if (-279 >= 9223372036854775807) {
                                                                                                                                var v129 uint = 0
                                                                                                                                v113 = (v113 & (v113 + v5))
                                                                                                                                _ = v129
                                                                                                                            }
                                                                                                                            v78 = append(v78, 42)
                                                                                                                            var v130 string = "hello"
                                                                                                                            _ = v130
                                                                                                                        }
                                                                                                                    case -178:
                                                                                                                        if len(v78) >= 2 {
                                                                                                                            v78 = v78[1:len(v78)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                var v131 float32 = (1.5 * v2)
                                                                                                                _ = v131
                                                                                                            } else {
                                                                                                                v77 = fi101
                                                                                                                _ = len(v6)
                                                                                                                var v132 uint16 = 255
                                                                                                                _ = v132
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v1, v113)
                                                                                                        var v133 int = fi101
                                                                                                        _ = len(v4)
                                                                                                        _ = v133
                                                                                                    }
                                                                                                    var v134 uint32 = (v113 % (((v113 + (v5 | v113))) | 1))
                                                                                                    v135 := GenS10{F0: false, F1: -1, F2: 4294967295}
                                                                                                    _ = v134
                                                                                                    _ = v135
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                var v136 uint16 = uint16(v1)
                                                                                                _ = v136
                                                                                            }
                                                                                            switch ("hello" + ("" + "hello")) {
                                                                                            case "hello":
                                                                                                v76 += (v76 << ((42) & 63))
                                                                                            }
                                                                                            _ = v113
                                                                                        }
                                                                                    case 32768:
                                                                                        fmt.Println(v79)
                                                                                    case 0:
                                                                                        switch 363 {
                                                                                        case -345:
                                                                                            fmt.Println(v79, v1)
                                                                                        }
                                                                                    }
                                                                                    var v137 int16 = int16(fi102)
                                                                                    var v138 int64 = -9223372036854775808
                                                                                    _ = v137
                                                                                    _ = v138
                                                                                } else {
                                                                                    switch fi101 {
                                                                                    case 1:
                                                                                        v3 -= v3
                                                                                    case 4294967295:
                                                                                        var v139 string = ("hello" + "a")
                                                                                        _ = v139
                                                                                    case -659:
                                                                                        _ = len(v78)
                                                                                    }
                                                                                    var v140 int32 = 768
                                                                                    if false {
                                                                                        v6[(^(fi101))] = (fi101 % (((-(fi102))) | 1))
                                                                                        v3 = (v3 >> ((279) & 63))
                                                                                        if v92 {
                                                                                            if v92 {
                                                                                                var v141 int32 = (v140 / ((-2147483648) | 1))
                                                                                                _ = v141
                                                                                            } else {
                                                                                                if ((32767 < 255) && v92) {
                                                                                                    v2 = (3.14159 - 0.1)
                                                                                                    if false {
                                                                                                        switch (-(fi101)) {
                                                                                                        case 131:
                                                                                                            if len(v78) > 0 {
                                                                                                                v78[0] = fi101
                                                                                                            }
                                                                                                        case 32768:
                                                                                                            v78 = append(v78, v77)
                                                                                                        }
                                                                                                        var v142 int32 = (^(v140))
                                                                                                        v143 := GenS4{F0: 2, F1: "\n\t", F2: (v79 ^ 1)}
                                                                                                        _ = v142
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    var v144 uint32 = uint32(v79)
                                                                                                    var v145 uint = 18446744073709551615
                                                                                                    _ = v144
                                                                                                    _ = v145
                                                                                                }
                                                                                                v146 := GenS11{F0: "hello"}
                                                                                                switch v77 {
                                                                                                case -538:
                                                                                                    switch v77 {
                                                                                                    case -1:
                                                                                                        if (255 > 357) {
                                                                                                            fmt.Println(v5)
                                                                                                            v147 := GenS12{F0: v140, F1: -327}
                                                                                                            _ = v147
                                                                                                        }
                                                                                                    }
                                                                                                case 1:
                                                                                                    if (v92 || v92) {
                                                                                                        var v148 int8 = v76
                                                                                                        v140 += v140
                                                                                                        if false {
                                                                                                            delete(v6, -1)
                                                                                                            var v149 int32 = (v140 | (v140 << ((65535) & 63)))
                                                                                                            v150 := GenS12{F0: int32(v79), F1: 0}
                                                                                                            if false {
                                                                                                                if v92 {
                                                                                                                    v96 += v96
                                                                                                                    v4[(v5 % ((2147483648) | 1))] = ("a" + "hello")
                                                                                                                    if (v92 || (v2 <= (v91 + 3.14159))) {
                                                                                                                        switch (("hello" + "foo bar") + (("\n\t" + ((("foo bar" + "\n\t") + ("hello" + ("foo bar" + (("\n\t" + "hello") + "")))) + "\n\t")) + "foo bar")) {
                                                                                                                        case "hello":
                                                                                                                            if (((!v92) && v92) && (v92 || ((v148 + int8(v5)) != (v76 & (^(v76)))))) {
                                                                                                                                _ = len(v78)
                                                                                                                                v91 = 0.0
                                                                                                                                if v92 {
                                                                                                                                    v2 = v91
                                                                                                                                    v4[934] = "a"
                                                                                                                                    v151 := v4[(^(v5))]
                                                                                                                                    _ = v151
                                                                                                                                }
                                                                                                                                v140 -= -2147483648
                                                                                                                                var v152 bool = v92
                                                                                                                                _ = v152
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v153 uint16 = 273
                                                                                                                    v140 = v149
                                                                                                                    fmt.Println(v153, v3)
                                                                                                                    _ = v153
                                                                                                                }
                                                                                                                v154 := GenS12{F0: 0, F1: 9223372036854775807}
                                                                                                                var v155 string = ""
                                                                                                                v156 := GenS13{F0: (v3 & (v3 % ((-1) | 1))), F1: (v1 * 251), F2: (v140 | v140)}
                                                                                                                v78 = append(v78, 32767)
                                                                                                                _ = v154
                                                                                                                _ = v155
                                                                                                                _ = v156
                                                                                                            }
                                                                                                            _ = v149
                                                                                                            _ = v150
                                                                                                        }
                                                                                                        _ = v148
                                                                                                    } else {
                                                                                                        if v92 {
                                                                                                            if false {
                                                                                                                if v92 {
                                                                                                                    v140 = (-(v140))
                                                                                                                } else {
                                                                                                                    if v92 {
                                                                                                                        var v157 float64 = v96
                                                                                                                        _ = v157
                                                                                                                    }
                                                                                                                }
                                                                                                                delete(v6, fi101)
                                                                                                                fmt.Println(v2)
                                                                                                            }
                                                                                                            _ = len(v78)
                                                                                                            if (false || (genFn1() < 9223372036854775807)) {
                                                                                                                v158 := GenS12{F0: -2147483648, F1: 917}
                                                                                                                v79 += 2
                                                                                                                _ = v158
                                                                                                            } else {
                                                                                                                var v159 int16 = (v3 / ((int16(v77)) | 1))
                                                                                                                v160 := GenS14{F0: 4294967295}
                                                                                                                _ = v159
                                                                                                                _ = v160
                                                                                                            }
                                                                                                        } else {
                                                                                                            if (!v92) {
                                                                                                                var v161 int = (-(v77))
                                                                                                                var v162 int16 = -919
                                                                                                                _ = v161
                                                                                                                _ = v162
                                                                                                            } else {
                                                                                                                v78 = append(v78, 2147483647)
                                                                                                                var v163 float64 = (1e10 - float64(v2))
                                                                                                                _ = v163
                                                                                                            }
                                                                                                            v164 := GenS1{F0: false, F1: 90, F2: "hello"}
                                                                                                            v165 := GenS15{F0: (v76 / ((v76) | 1))}
                                                                                                            fmt.Println(v1, v76)
                                                                                                            var v166 uint16 = 42
                                                                                                            fmt.Println(v92)
                                                                                                            _ = v164
                                                                                                            _ = v165
                                                                                                            _ = v166
                                                                                                        }
                                                                                                        delete(v4, (v5 / (((^(v5))) | 1)))
                                                                                                        fmt.Println(v1, fi101, v1)
                                                                                                    }
                                                                                                default:
                                                                                                    fmt.Println(v2, v92)
                                                                                                }
                                                                                                _ = v146
                                                                                            }
                                                                                        } else {
                                                                                            v167 := GenS16{F0: v79, F1: 42}
                                                                                            _ = v167
                                                                                        }
                                                                                    }
                                                                                    _ = v140
                                                                                }
                                                                                if len(v78) >= 2 {
                                                                                    v78 = v78[1:len(v78)]
                                                                                }
                                                                            } else {
                                                                                if len(v78) >= 2 {
                                                                                    v78 = v78[1:len(v78)]
                                                                                }
                                                                                var v168 string = (("hello" + "\n\t") + "\n\t")
                                                                                var v169 uint64 = (v79 & v79)
                                                                                var v170 int = (fi101 >> ((65535) & 63))
                                                                                var v171 string = ("hello" + ("" + (v168 + "")))
                                                                                if len(v78) > 0 {
                                                                                    v78[0] = v170
                                                                                }
                                                                                switch 32768 {
                                                                                case 0:
                                                                                    v168 = v168
                                                                                }
                                                                                _ = v168
                                                                                _ = v169
                                                                                _ = v170
                                                                                _ = v171
                                                                            }
                                                                            var v172 uint64 = (v79 << ((767) & 63))
                                                                            _ = fi102
                                                                            _ = v172
                                                                        }
                                                                    }
                                                                    switch v77 {
                                                                    case 9223372036854775807:
                                                                        var v173 uint16 = uint16(v5)
                                                                        _ = v173
                                                                    case 367:
                                                                        _ = len(v4)
                                                                    case 127:
                                                                        var v174 int32 = -101
                                                                        _ = v174
                                                                    default:
                                                                        var v175 uint64 = (v79 | 42)
                                                                        _ = v175
                                                                    }
                                                                    v176 := GenS2{F0: 65535, F1: (v79 ^ v79)}
                                                                    _ = v176
                                                                }
                                                                var v177 float32 = v91
                                                                _ = fi101
                                                                _ = v177
                                                            }
                                                            switch ("a" + "foo bar") {
                                                            case "":
                                                                var v178 bool = (65535 <= 616)
                                                                _ = v178
                                                            case "foo bar":
                                                                fmt.Println(v76)
                                                            default:
                                                                for fi179 := 0; fi179 < 15; fi179++ {
                                                                    v180 := GenS5{F0: -202, F1: -358}
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v181 := GenS4{F0: 42, F1: "hello", F2: v79}
                                                                        _ = v181
                                                                    default:
                                                                        v182 := GenS1{F0: v92, F1: (v77 + (^(v77))), F2: "a"}
                                                                        _ = v182
                                                                    }
                                                                    _ = fi179
                                                                    _ = v180
                                                                }
                                                            }
                                                            for fi183 := 0; fi183 < 10; fi183++ {
                                                                v184 := v4[v5]
                                                                v185 := GenS17{F0: 42}
                                                                _ = fi183
                                                                _ = v184
                                                                _ = v185
                                                            }
                                                        } else {
                                                            switch v77 {
                                                            case -1:
                                                                for fi186 := 0; fi186 < 20; fi186++ {
                                                                    var v187 uint64 = v79
                                                                    v91 -= 0.1
                                                                    fmt.Println(v79)
                                                                    _ = fi186
                                                                    _ = v187
                                                                }
                                                            case -203:
                                                                var v188 int = (v77 ^ 305)
                                                                _ = v188
                                                            case 2147483647:
                                                                if false {
                                                                    v189 := GenS15{F0: (^(v76))}
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v91 += (float32(v96) / v2)
                                                                    case "":
                                                                        v190 := GenS7{F0: (-(v76)), F1: uint64(v1), F2: (v77 ^ (v77 % (((-(v77))) | 1)))}
                                                                        _ = v190
                                                                    case "foo bar":
                                                                        switch (v77 ^ v77) {
                                                                        case 128:
                                                                            v191 := GenS18{F0: 42, F1: v3}
                                                                            _ = v191
                                                                        case 512:
                                                                            var v192 uint64 = uint64(v1)
                                                                            _ = v192
                                                                        case -282:
                                                                            v193 := GenS14{F0: 4294967295}
                                                                            _ = v193
                                                                        default:
                                                                            v96 -= v96
                                                                        }
                                                                    default:
                                                                        if (v96 == float64(v91)) {
                                                                            fmt.Println(v1, v77, v92)
                                                                        } else {
                                                                            v1 += v1
                                                                            v194 := GenS19{F0: v92, F1: 152}
                                                                            _ = v194
                                                                        }
                                                                    }
                                                                    _ = v189
                                                                }
                                                            default:
                                                                switch ("a" + "\n\t") {
                                                                case "\n\t":
                                                                    v79 = (v79 % (((v79 / ((v79) | 1))) | 1))
                                                                case "":
                                                                    var v195 string = "foo bar"
                                                                    _ = v195
                                                                }
                                                            }
                                                            v2 = float32((v96 / v96))
                                                            v196 := GenS13{F0: genFn0((v76 >> ((489) & 63)), v76), F1: 127, F2: 1}
                                                            _ = v196
                                                        }
                                                        if (!(9223372036854775807 >= int64(v79))) {
                                                            var v197 int64 = -1
                                                            var v198 int32 = 904
                                                            _ = v197
                                                            _ = v198
                                                        }
                                                    }
                                                } else {
                                                    v199 := GenS20{F0: int32(v5), F1: 478}
                                                    v79 += 42
                                                    _ = v199
                                                }
                                            }
                                        case 42:
                                            v3 += (v3 - (v3 - v3))
                                        default:
                                            var v200 uint64 = (v79 ^ (v79 & 381))
                                            _ = v200
                                        }
                                    default:
                                        if v92 {
                                            _ = len(v78)
                                            switch v77 {
                                            case -9223372036854775808:
                                                if len(v78) >= 2 {
                                                    v78 = v78[1:len(v78)]
                                                }
                                            default:
                                                for fi201 := 0; fi201 < 12; fi201++ {
                                                    v202 := GenS21{F0: 255, F1: -2147483648, F2: (v5 << ((18446744073709551615) & 63))}
                                                    for fi203 := 0; fi203 < 9; fi203++ {
                                                        var v204 uint16 = 266
                                                        delete(v4, v5)
                                                        _ = fi203
                                                        _ = v204
                                                    }
                                                    var v205 int = (v77 + fi201)
                                                    _ = fi201
                                                    _ = v202
                                                    _ = v205
                                                }
                                            }
                                            var v206 uint32 = (^(v5))
                                            _ = v206
                                        }
                                    }
                                    v77 = (v77 << ((18446744073709551615) & 63))
                                    var v207 uint64 = (v79 >> ((42) & 63))
                                    _ = v207
                                }
                            }
                            var v208 int8 = -1
                            _ = v208
                        }
                        if ((^(v1)) >= (v1 / ((v1) | 1))) {
                            fmt.Println(v96)
                        }
                        _ = v96
                    }
                }
                _ = v92
            }
        case 4294967295:
            var v209 uint = 4294967295
            _ = v209
        case 157:
            var v210 int64 = -1
            _ = v210
        default:
            var v211 int32 = -555
            _ = v211
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("slicelen v78:", len(v78), v78)
    fmt.Println("v79:", v79)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v91
}
