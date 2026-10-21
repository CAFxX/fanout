package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 string; F1 uint8; F2 bool }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 string; F1 string }
type GenS10 struct { F0 int32; F1 uint32; F2 bool }

func main() {
    var v1 int16 = 826
    v2 := []uint32{924, 2147483648, 4294967295, 2}
    var v3 string = ((("" + "\n\t") + ((("hello" + "") + "\n\t") + ("\n\t" + ("foo bar" + (("a" + "\n\t") + "foo bar"))))) + "a")
    v3 = "hello"
    v4 := GenS1{F0: 4294967295, F1: 0}
    var v5 uint8 = 255
    if len(v2) > 0 {
        v2[0] = 0
    }
    switch -9223372036854775808 {
    case -2147483648:
        var v6 int = -1
        _ = v6
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        if true {
            v8 := GenS1{F0: 1, F1: 1}
            _ = v8
        } else {
            if (!false) {
                if len(v2) > 0 {
                    v2[0] = 577
                }
            }
        }
        switch v3 {
        case "a":
            v3 = ("hello" + ((v3 + (((v3 + v3) + "foo bar") + v3)) + v3))
        default:
            v2 = append(v2, 65535)
        }
        if (!(0.1 <= 0.0)) {
            fmt.Println(v5)
        } else {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 string = (v3 + (("a" + (v3 + v3)) + "\n\t"))
                _ = fi9
                _ = v10
            }
        }
        _ = fi7
    }
    v5 -= (^(v5))
    v11 := GenS1{F0: 999, F1: 288}
    if (!(!(v3 <= "hello"))) {
        fmt.Println(v1, v5, v1)
        var v12 uint32 = 65535
        _ = v12
    }
    switch v3 {
    case "foo bar":
        v13 := GenS1{F0: 2, F1: 42}
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 20; fi14++ {
            v5 += v5
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(fi14, v3, v3)
                _ = fi15
            }
            _ = fi14
        }
    case "":
        var v16 int32 = 1
        _ = v16
    }
    if (!((-(v1)) <= v1)) {
        fmt.Println(v1, v1)
    }
    if false {
        var v17 int16 = 932
        if (!false) {
            var v18 string = (v3 + v3)
            _ = v18
        } else {
            if (9223372036854775807 >= 9223372036854775807) {
                for fi19 := 0; fi19 < 7; fi19++ {
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 float64 = 1.5
                        if (true || (100 <= 127)) {
                            if (float32(v21) < float32(1.5)) {
                                if (v5 < v5) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if (true || (1 < -869)) {
                                        fmt.Println(fi20)
                                    } else {
                                        var v22 int64 = 815
                                        v22 += -1
                                        v1 += (v1 << ((42) & 63))
                                        _ = v22
                                    }
                                    var v23 int16 = (v1 * (v1 / ((v17) | 1)))
                                    _ = v23
                                }
                            }
                            var v24 int16 = v17
                            v25 := GenS1{F0: 0, F1: 0}
                            _ = v24
                            _ = v25
                        } else {
                            v26 := GenS1{F0: 339, F1: 255}
                            if (v3 <= v3) {
                                var v27 bool = false
                                var v28 bool = true
                                if (v27 && true) {
                                    var v29 int8 = -1
                                    v27 = ((-(v29)) > 0)
                                    _ = v29
                                } else {
                                    if (v3 != v3) {
                                        if (int32(v5) == 1) {
                                            if len(v2) > 0 {
                                                v2[0] = 259
                                            }
                                            switch fi20 {
                                            case 9223372036854775807:
                                                v2 = append(v2, 4294967295)
                                            }
                                            fmt.Println(v27, fi20)
                                        } else {
                                            v28 = (!(((v21 / (v21 + v21)) * float64(-2.25)) < 0.0))
                                            v5 -= v5
                                            if ((!(!v27)) || (700 > 273)) {
                                                v30 := GenS1{F0: 2, F1: 2}
                                                _ = v30
                                            }
                                        }
                                    } else {
                                        if (false || (v27 || v27)) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            var v31 uint32 = 2
                                            v32 := GenS1{F0: 43, F1: 18446744073709551615}
                                            _ = v31
                                            _ = v32
                                        }
                                        if (527 < 2) {
                                            v3 = "a"
                                            if v28 {
                                                fmt.Println(fi19, v5, v28)
                                                var v33 int64 = int64(v17)
                                                _ = v33
                                            } else {
                                                v3 = (v3 + (v3 + (((v3 + v3) + v3) + ("\n\t" + (v3 + v3)))))
                                                v5 += (v5 * v5)
                                                v27 = (476 != 0)
                                                _ = len(v2)
                                            }
                                            if (750 != 2147483648) {
                                                fmt.Println(v17)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if (881 <= 870) {
                                                    v21 = (1e10 - -2.25)
                                                    var v34 int64 = -9223372036854775808
                                                    fmt.Println(fi19)
                                                    if (!v27) {
                                                        v35 := GenS2{F0: 104, F1: 34}
                                                        v27 = (65535 != 0)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v35
                                                    }
                                                    _ = v34
                                                } else {
                                                    switch fi19 {
                                                    case 2147483647:
                                                        var v36 float64 = v21
                                                        _ = v36
                                                    }
                                                    v3 = (v3 + v3)
                                                    v27 = false
                                                }
                                                v37 := GenS3{F0: 430, F1: v28}
                                                _ = v37
                                            } else {
                                                v38 := GenS4{F0: 42, F1: 885}
                                                if (uint(v5) < 255) {
                                                    var v39 uint32 = 2147483647
                                                    switch fi20 {
                                                    case -129:
                                                        if (((v1 >> ((4294967295) & 63)) >= v17) && v27) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v21 = v21
                                                            var v40 int8 = 0
                                                            switch (fi20 >> ((2) & 63)) {
                                                            case -308:
                                                                _ = len(v2)
                                                            case -272:
                                                                v41 := GenS1{F0: 0, F1: 601}
                                                                _ = v41
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = 0
                                                            }
                                                            _ = v40
                                                        }
                                                    default:
                                                        v42 := GenS4{F0: 1, F1: 50}
                                                        _ = v42
                                                    }
                                                    var v43 bool = false
                                                    var v44 uint32 = (v39 & (^(v39)))
                                                    v45 := GenS4{F0: 1, F1: 2}
                                                    var v46 int32 = -2147483648
                                                    var v47 int8 = 0
                                                    v47 = -17
                                                    v44 += v39
                                                    _ = v39
                                                    _ = v43
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                }
                                                _ = v38
                                            }
                                        } else {
                                            if v27 {
                                                var v48 uint32 = 284
                                                v21 = (float64(3.14159) * 0.1)
                                                if (-929 >= -1) {
                                                    switch "hello" {
                                                    case "":
                                                        var v49 int16 = 42
                                                        _ = v49
                                                    case "hello":
                                                        v50 := GenS2{F0: 984, F1: 2}
                                                        _ = v50
                                                    }
                                                    v51 := GenS2{F0: 1, F1: 275}
                                                    _ = v51
                                                }
                                                _ = v48
                                            }
                                        }
                                        v27 = v28
                                    }
                                    var v52 int16 = v1
                                    if (65535 > 0) {
                                        if v27 {
                                            if v27 {
                                                switch (v3 + "hello") {
                                                case "\n\t":
                                                    _ = len(v2)
                                                case "a":
                                                    v53 := GenS3{F0: 0, F1: v28}
                                                    v17 -= v1
                                                    v5 = v5
                                                    _ = v53
                                                }
                                                switch (-(fi20)) {
                                                case 128:
                                                    v1 = 258
                                                case -2147483648:
                                                    if (v21 > (1e10 - 3.14159)) {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v54 := GenS3{F0: 2, F1: v27}
                                                        _ = v54
                                                    }
                                                    v17 = v1
                                                    var v55 float64 = v21
                                                    if (-75 <= 0) {
                                                        v56 := GenS3{F0: 42, F1: (-840 > 0)}
                                                        _ = v56
                                                    }
                                                    switch (v3 + "\n\t") {
                                                    case "\n\t":
                                                        v1 = v1
                                                    case "a":
                                                        v27 = (!((v27 || v28) || ((42 == 1) && false)))
                                                    }
                                                    _ = v55
                                                case 127:
                                                    var v57 uint = 0
                                                    var v58 uint32 = 149
                                                    v2 = append(v2, (v58 >> ((v57) & 63)))
                                                    _ = v57
                                                    _ = v58
                                                default:
                                                    if (!(!(42 > 240))) {
                                                        _ = len(v2)
                                                    }
                                                }
                                                v59 := GenS1{F0: 18446744073709551615, F1: 758}
                                                _ = v59
                                            }
                                            fmt.Println(v27, v17)
                                        }
                                        switch v3 {
                                        case "a":
                                            v3 = "foo bar"
                                        }
                                        switch fi20 {
                                        case 42:
                                            switch "hello" {
                                            case "\n\t":
                                                if (v27 && ((v3 > "a") || (546 >= 2))) {
                                                    switch 0 {
                                                    case -442:
                                                        v60 := GenS2{F0: 2, F1: uint64(v5)}
                                                        _ = v60
                                                    case -1:
                                                        var v61 uint16 = 65535
                                                        _ = v61
                                                    }
                                                } else {
                                                    switch "\n\t" {
                                                    case "":
                                                        var v62 int = fi20
                                                        _ = v62
                                                    }
                                                    v21 = 0.0
                                                }
                                            }
                                        }
                                    }
                                    _ = v52
                                }
                                switch (fi20 - fi20) {
                                case 2147483647:
                                    if true {
                                        v2 = append(v2, 224)
                                    }
                                case 128:
                                    v3 = v3
                                }
                                fmt.Println(v21)
                                _ = v27
                                _ = v28
                            } else {
                                v21 = (-2.25 / 1e10)
                                var v63 int32 = -2147483648
                                var v64 float64 = (1.5 - ((1.5 / 1.5) + 0.0))
                                _ = v63
                                _ = v64
                            }
                            switch -292 {
                            case 189:
                                var v65 float64 = (v21 + 1.5)
                                _ = v65
                            case 32767:
                                switch (fi19 + (fi19 + (fi19 & (fi20 / (((fi20 / ((fi20) | 1))) | 1))))) {
                                case 128:
                                    switch "\n\t" {
                                    case "a":
                                        v2 = append(v2, 750)
                                    default:
                                        var v66 int32 = 144
                                        _ = v66
                                    }
                                default:
                                    v21 = 1.5
                                }
                            default:
                                fmt.Println(v21, fi19)
                            }
                            v3 = v3
                            v2 = append(v2, 65535)
                            if true {
                                switch v3 {
                                case "\n\t":
                                    v2 = append(v2, 177)
                                case "foo bar":
                                    v67 := GenS3{F0: 1, F1: true}
                                    _ = v67
                                }
                            } else {
                                v2 = append(v2, 944)
                                if (42 > 42) {
                                    _ = len(v2)
                                    var v68 int = (fi20 * 681)
                                    if (0 == 514) {
                                        v69 := GenS1{F0: 255, F1: 18446744073709551615}
                                        v1 += 1
                                        _ = v69
                                    } else {
                                        if false {
                                            var v70 int64 = 0
                                            v70 = v70
                                            _ = v70
                                        }
                                        v71 := GenS3{F0: 42, F1: false}
                                        if false {
                                            if false {
                                                v72 := GenS5{F0: 255, F1: -1}
                                                if true {
                                                    v73 := GenS6{F0: 0, F1: 42}
                                                    _ = v73
                                                }
                                                _ = v72
                                            } else {
                                                v74 := GenS6{F0: -2147483648, F1: 42}
                                                _ = v74
                                            }
                                            var v75 uint32 = 231
                                            _ = v75
                                        }
                                        _ = v71
                                    }
                                    _ = v68
                                }
                                switch fi19 {
                                case 4294967295:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                            _ = v26
                        }
                        _ = fi20
                        _ = v21
                    }
                    var v76 float32 = 0.1
                    v1 = (v17 % ((v1) | 1))
                    _ = fi19
                    _ = v76
                }
            } else {
                if true {
                    switch 267 {
                    case 668:
                        v77 := GenS4{F0: 42, F1: 160}
                        _ = v77
                    case 42:
                        fmt.Println(v1, v1, v5)
                    case -157:
                        v17 = (v1 - (v1 * (v1 >> ((1) & 63))))
                    default:
                        v1 -= v17
                    }
                }
            }
            var v78 int8 = -1
            _ = v78
        }
        switch "hello" {
        case "foo bar":
            v1 = v1
        case "a":
            for fi79 := 0; fi79 < 9; fi79++ {
                var v80 int = (fi79 ^ (fi79 / (((fi79 / (((fi79 - -84)) | 1))) | 1)))
                for fi81 := 0; fi81 < 2; fi81++ {
                    if false {
                        v1 = 318
                        fmt.Println(v17, v80)
                        if len(v2) > 0 {
                            v2[0] = 137
                        }
                    } else {
                        fmt.Println(v3, fi79, v17)
                        fmt.Println(v17)
                        switch v3 {
                        case "foo bar":
                            v82 := GenS7{F0: v3, F1: v5, F2: false}
                            _ = v82
                        case "a":
                            switch (v3 + "\n\t") {
                            case "":
                                v83 := GenS8{F0: 494, F1: 18446744073709551615}
                                _ = v83
                            case "a":
                                switch (v80 * -2147483648) {
                                case 9223372036854775807:
                                    if true {
                                        if (!(9223372036854775807 < -347)) {
                                            var v84 uint64 = 18446744073709551615
                                            _ = v84
                                        } else {
                                            if ((42 < -96) || false) {
                                                var v85 int8 = 0
                                                _ = v85
                                            } else {
                                                fmt.Println(v3, v80, fi81)
                                                var v86 string = "foo bar"
                                                switch ("a" + "a") {
                                                case "a":
                                                    v17 -= (v1 * v1)
                                                default:
                                                    var v87 uint64 = 0
                                                    v86 = "hello"
                                                    fmt.Println(v87, v3, v17)
                                                    _ = v87
                                                }
                                                var v88 float64 = (float64(0.0) * 0.1)
                                                _ = v86
                                                _ = v88
                                            }
                                        }
                                        v5 += (^(v5))
                                        var v89 uint8 = 42
                                        _ = v89
                                    }
                                case 0:
                                    if (18446744073709551615 < 769) {
                                        v1 = v17
                                        v2 = append(v2, 1)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                            }
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v2 = append(v2, 905)
                    _ = fi81
                }
                for fi90 := 0; fi90 < 13; fi90++ {
                    v3 = ((v3 + ((v3 + v3) + (v3 + v3))) + v3)
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi90
                }
                _ = fi79
                _ = v80
            }
        case "hello":
            fmt.Println(v1, v3)
        }
        _ = v17
    } else {
        v91 := GenS9{F0: "\n\t", F1: v3}
        var v92 uint16 = 1
        _ = v91
        _ = v92
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= (v1 << ((0) & 63))
    switch v3 {
    case "hello":
        for fi93 := 0; fi93 < 15; fi93++ {
            switch "" {
            case "a":
                for fi94 := 0; fi94 < 6; fi94++ {
                    v1 -= v1
                    v1 = (v1 ^ (^(v1)))
                    var v95 int16 = (v1 - 673)
                    _ = fi94
                    _ = v95
                }
            case "hello":
                var v96 int16 = v1
                _ = v96
            case "":
                fmt.Println(v1, fi93, fi93)
            default:
                fmt.Println(v5, v5, v1)
            }
            _ = fi93
        }
    case "":
        v97 := GenS10{F0: -2147483648, F1: 65535, F2: ((9223372036854775807 >= 42) || true)}
        _ = v97
    default:
        if (!true) {
            fmt.Println(v5)
            if ((false || ((1e10 / 3.14159) <= (0.1 - 3.14159))) && false) {
                if true {
                    switch 4294967295 {
                    case -2147483648:
                        for fi98 := 0; fi98 < 7; fi98++ {
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            if ("" <= (((v3 + v3) + v3) + v3)) {
                                for fi99 := 0; fi99 < 4; fi99++ {
                                    _ = len(v2)
                                    switch (fi98 << ((1) & 63)) {
                                    case 0:
                                        if true {
                                            if false {
                                                var v100 int8 = -77
                                                v2 = append(v2, 65535)
                                                v100 += v100
                                                _ = v100
                                            } else {
                                                switch v3 {
                                                case "\n\t":
                                                    v3 = (v3 + v3)
                                                case "":
                                                    v2 = append(v2, 0)
                                                }
                                                v2 = append(v2, 2)
                                                fmt.Println(fi98, fi99, v5)
                                                if (true && (!(!(false && (-9223372036854775808 <= 0))))) {
                                                    switch (fi99 | 2147483647) {
                                                    case -9223372036854775808:
                                                        var v101 bool = true
                                                        _ = v101
                                                    }
                                                    if false {
                                                        var v102 string = ""
                                                        var v103 int64 = 9223372036854775807
                                                        fmt.Println(fi98, v102)
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    v104 := GenS2{F0: 51, F1: 2}
                                                    _ = v104
                                                }
                                                var v105 int8 = 117
                                                _ = v105
                                            }
                                            v5 = (v5 << ((131) & 63))
                                            v2 = append(v2, 60)
                                            var v106 string = "a"
                                            var v107 uint32 = 1
                                            _ = v106
                                            _ = v107
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = 2
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = 951
                                            }
                                        }
                                    case 407:
                                        var v108 uint16 = 2
                                        var v109 string = v3
                                        switch "hello" {
                                        case "hello":
                                            switch v3 {
                                            case "a":
                                                v2 = append(v2, 42)
                                            case "\n\t":
                                                var v110 int = (fi99 << ((1) & 63))
                                                _ = v110
                                            default:
                                                v109 = "foo bar"
                                            }
                                        case "":
                                            v5 -= (v5 * v5)
                                        }
                                        _ = v108
                                        _ = v109
                                    case -463:
                                        fmt.Println(v3, fi99, v5)
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = 2
                                        }
                                    }
                                    v111 := GenS4{F0: 42, F1: 65535}
                                    _ = fi99
                                    _ = v111
                                }
                            } else {
                                var v112 string = v3
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v112
                            }
                            _ = fi98
                        }
                    case 42:
                        var v113 uint64 = 248
                        _ = v113
                    }
                } else {
                    var v114 string = v3
                    _ = v114
                }
                v2 = append(v2, 2147483647)
            }
            switch ((v3 + v3) + (v3 + ((v3 + v3) + ""))) {
            case "foo bar":
                v3 = v3
            }
        } else {
            var v115 bool = (1e10 != 3.14159)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v115 = v115
            _ = v115
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
