package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 int32; F1 bool; F2 int64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32; F1 bool; F2 uint16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS14 struct { F0 uint64; F1 int16; F2 int32 }

func main() {
    v1 := []uint32{4294967295, 2147483648}
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    v4 := GenS1{F0: 1, F1: 4294967295, F2: 2147483647}
    v2 = "a"
    if (15 > 127) {
        for fi5 := 0; fi5 < 7; fi5++ {
            if (0.1 < 1.5) {
                v2 = "hello"
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v6 := GenS1{F0: -2147483648, F1: 50, F2: -298}
                _ = v6
            } else {
                v1 = append(v1, 0)
                var v7 uint32 = 0
                v7 = v7
                _ = v7
            }
            v8 := GenS1{F0: -281, F1: 42, F2: 574}
            _ = fi5
            _ = v8
        }
        var v9 uint16 = 0
        v10 := GenS2{F0: -1, F1: 79}
        _ = v10
        _ = v9
    } else {
        if (v2 >= ("" + (((v2 + (v2 + v2)) + v2) + ("hello" + "")))) {
            var v11 bool = true
            v3[("\n\t" + v2)] = v2
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 int16 = -1
            var v14 uint64 = 2
            _ = fi12
            _ = v13
            _ = v14
        }
        v2 = (v2 + v2)
    }
    v2 = "foo bar"
    fmt.Println(v2)
    fmt.Println(v2)
    v2 = v2
    v2 = (v2 + v2)
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint64 = 42
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 bool = (!false)
            fmt.Println(v18, v16)
            _ = fi17
            _ = v18
        }
        v19 := GenS3{F0: (2147483647 == 102), F1: 393}
        _ = fi15
        _ = v16
        _ = v19
    }
    v2 = "\n\t"
    if true {
        fmt.Println(v2)
        if (710 != 946) {
            switch (v2 + ((v2 + "foo bar") + (((v2 + "hello") + v2) + (v2 + v2)))) {
            case "hello":
                if (4294967295 < 42) {
                    v20 := v3[(("hello" + v2) + "hello")]
                    for fi21 := 0; fi21 < 9; fi21++ {
                        var v22 uint64 = 62
                        fmt.Println(v22, v22)
                        _ = fi21
                        _ = v22
                    }
                    if (950 != 65535) {
                        var v23 uint64 = 42
                        _ = v23
                    }
                    _ = v20
                } else {
                    v24 := GenS1{F0: 1, F1: 2147483647, F2: -1}
                    if (-2147483648 == 32768) {
                        if false {
                            var v25 string = "\n\t"
                            var v26 int32 = 87
                            _ = v25
                            _ = v26
                        } else {
                            fmt.Println(v2)
                            switch 42 {
                            case 9223372036854775807:
                                v2 = v2
                            default:
                                switch v2 {
                                case "":
                                    v2 = v2
                                default:
                                    switch v2 {
                                    case "":
                                        if (606 != 433) {
                                            for fi27 := 0; fi27 < 15; fi27++ {
                                                var v28 int64 = -1
                                                _ = fi27
                                                _ = v28
                                            }
                                        } else {
                                            v29 := GenS4{F0: 4294967295}
                                            switch (("\n\t" + (("hello" + "") + v2)) + "hello") {
                                            case "hello":
                                                v30 := GenS5{F0: 42, F1: 127}
                                                _ = v30
                                            default:
                                                switch 32768 {
                                                case -64:
                                                    v1 = append(v1, 2147483648)
                                                case -643:
                                                    for fi31 := 0; fi31 < 18; fi31++ {
                                                        _ = len(v3)
                                                        _ = fi31
                                                    }
                                                default:
                                                    for fi32 := 0; fi32 < 1; fi32++ {
                                                        switch (v2 + ("" + v2)) {
                                                        case "":
                                                            if (false || false) {
                                                                v2 = v2
                                                                var v33 int16 = -1
                                                                for fi34 := 0; fi34 < 16; fi34++ {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 4294967295
                                                                    }
                                                                    var v35 bool = (!(fi34 == 1))
                                                                    v35 = v35
                                                                    _ = fi34
                                                                    _ = v35
                                                                }
                                                                _ = v33
                                                            }
                                                        default:
                                                            var v36 int16 = 0
                                                            _ = v36
                                                        }
                                                        var v37 float32 = 1e10
                                                        switch v2 {
                                                        case "hello":
                                                            for fi38 := 0; fi38 < 19; fi38++ {
                                                                fmt.Println(fi38, fi38, v2)
                                                                _ = fi38
                                                            }
                                                        }
                                                        _ = fi32
                                                        _ = v37
                                                    }
                                                }
                                            }
                                            _ = v29
                                        }
                                    case "foo bar":
                                        if (4294967295 >= 2) {
                                            fmt.Println(v2)
                                            v2 = (v2 + "a")
                                        } else {
                                            var v39 uint16 = 0
                                            v40 := GenS4{F0: 255}
                                            _ = v39
                                            _ = v40
                                        }
                                    case "hello":
                                        if false {
                                            fmt.Println(v2)
                                        }
                                    }
                                }
                            }
                            switch -718 {
                            case 128:
                                v3[v2] = "\n\t"
                            case 1:
                                _ = len(v3)
                            case -95:
                                if true {
                                    var v41 uint16 = 255
                                    _ = v41
                                } else {
                                    for fi42 := 0; fi42 < 6; fi42++ {
                                        var v43 bool = false
                                        _ = fi42
                                        _ = v43
                                    }
                                }
                            default:
                                v2 = (v2 + v2)
                            }
                        }
                        v2 = "a"
                    }
                    _ = v24
                }
            case "foo bar":
                if (0.0 >= 1.5) {
                    for fi44 := 0; fi44 < 11; fi44++ {
                        v2 = ""
                        _ = fi44
                    }
                    var v45 uint8 = 75
                    _ = v45
                }
            case "":
                v46 := GenS6{F0: 65535}
                _ = v46
            default:
                for fi47 := 0; fi47 < 2; fi47++ {
                    var v48 bool = false
                    _ = fi47
                    _ = v48
                }
            }
            var v49 float64 = float64(((0.1 * (-2.25 + (3.14159 * 0.0))) * ((0.1 + 0.0) * 3.14159)))
            switch -9223372036854775808 {
            case 32767:
                _ = len(v3)
            default:
                v49 = v49
            }
            _ = v49
        }
    } else {
        v50 := GenS2{F0: 2147483647, F1: 100}
        var v51 int16 = 32767
        var v52 int32 = 524
        _ = v50
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 uint64 = 159
        switch "a" {
        case "hello":
            _ = len(v3)
        default:
            if true {
                if (-2147483648 <= 2147483647) {
                    v55 := GenS4{F0: 18446744073709551615}
                    if true {
                        v2 = (((v2 + v2) + (v2 + (("hello" + (("foo bar" + v2) + "foo bar")) + ""))) + "a")
                        if ((v54 + 4294967295) > 715) {
                            switch fi53 {
                            case 598:
                                v2 = v2
                            case -931:
                                switch "" {
                                case "hello":
                                    var v56 uint16 = 202
                                    _ = v56
                                case "foo bar":
                                    for fi57 := 0; fi57 < 17; fi57++ {
                                        v2 = "\n\t"
                                        v2 = v2
                                        var v58 int8 = 100
                                        _ = fi57
                                        _ = v58
                                    }
                                default:
                                    var v59 uint32 = 65535
                                    _ = v59
                                }
                            case -203:
                                var v60 string = "a"
                                _ = v60
                            default:
                                v61 := GenS7{F0: 2, F1: 42}
                                _ = v61
                            }
                            var v62 uint64 = 1
                            var v63 uint8 = 144
                            _ = v62
                            _ = v63
                        }
                    } else {
                        var v64 bool = true
                        var v65 float64 = (0.1 - 3.14159)
                        var v66 uint8 = 103
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    _ = v55
                }
            } else {
                v67 := GenS2{F0: -1, F1: 100}
                if false {
                    v1 = append(v1, 2)
                }
                _ = v67
            }
        }
        var v68 int = 46
        _ = fi53
        _ = v54
        _ = v68
    }
    var v69 int = -9223372036854775808
    fmt.Println(v2, v2)
    v69 = (v69 / ((v69) | 1))
    if ((^(v69)) < (v69 / ((v69) | 1))) {
        if (!(18446744073709551615 == 1)) {
            if true {
                for fi70 := 0; fi70 < 1; fi70++ {
                    v71 := GenS7{F0: 62, F1: 32767}
                    if true {
                        for fi72 := 0; fi72 < 18; fi72++ {
                            v69 += -214
                            _ = fi72
                        }
                        v2 = ""
                    }
                    _ = fi70
                    _ = v71
                }
                if (!(42 <= 2147483647)) {
                    v2 = v2
                } else {
                    v73 := GenS8{F0: "", F1: 2}
                    v74 := GenS7{F0: 255, F1: -41}
                    switch (v69 >> ((255) & 63)) {
                    case 128:
                        fmt.Println(v2, v2)
                    case -2147483648:
                        v69 += v69
                    default:
                        delete(v3, (v2 + v2))
                    }
                    _ = v73
                    _ = v74
                }
            }
            switch v69 {
            case -250:
                var v75 uint32 = 0
                _ = v75
            case 1:
                v2 = ((v2 + "") + (((("hello" + v2) + ("hello" + "\n\t")) + v2) + v2))
            default:
                switch "hello" {
                case "hello":
                    fmt.Println(v2, v69)
                case "foo bar":
                    var v76 int32 = 1
                    _ = v76
                default:
                    if (4294967295 <= 1) {
                        switch (v69 / ((986) | 1)) {
                        case 127:
                            var v77 uint8 = 128
                            _ = v77
                        case -290:
                            var v78 int64 = -879
                            _ = v78
                        }
                        if (!false) {
                            v69 += (v69 >> ((956) & 63))
                        }
                        v69 = (^(v69))
                    } else {
                        switch -9223372036854775808 {
                        case 4294967295:
                            var v79 int8 = 1
                            _ = v79
                        case 32768:
                            v2 = (v2 + v2)
                        case -2147483648:
                            var v80 int32 = 985
                            _ = v80
                        default:
                            var v81 bool = (false && true)
                            _ = v81
                        }
                        for fi82 := 0; fi82 < 18; fi82++ {
                            if true {
                                v2 = "foo bar"
                            } else {
                                v83 := GenS8{F0: ("\n\t" + v2), F1: 627}
                                if true {
                                    var v84 uint8 = 0
                                    v2 = v2
                                    if (2147483647 <= 1) {
                                        for fi85 := 0; fi85 < 12; fi85++ {
                                            var v86 bool = (9223372036854775807 < 0)
                                            switch v69 {
                                            case 127:
                                                var v87 int = v69
                                                _ = v87
                                            }
                                            _ = fi85
                                            _ = v86
                                        }
                                        delete(v3, v2)
                                    }
                                    _ = v84
                                } else {
                                    var v88 int32 = -147
                                    for fi89 := 0; fi89 < 7; fi89++ {
                                        v2 = "hello"
                                        var v90 uint64 = 518
                                        switch (v2 + v2) {
                                        case "\n\t":
                                            fmt.Println(v90, fi82, v88)
                                        case "hello":
                                            var v91 int32 = 42
                                            _ = v91
                                        case "":
                                            var v92 string = v2
                                            _ = v92
                                        }
                                        _ = fi89
                                        _ = v90
                                    }
                                    _ = v88
                                }
                                _ = v83
                            }
                            v93 := GenS1{F0: int32(v69), F1: 2147483647, F2: 2147483647}
                            _ = fi82
                            _ = v93
                        }
                    }
                }
            }
            if (1 == 42) {
                if (837 >= -32768) {
                    switch (v69 - 127) {
                    case 574:
                        var v94 int32 = -2147483648
                        _ = v94
                    }
                    var v95 uint = 531
                    var v96 bool = true
                    _ = v95
                    _ = v96
                } else {
                    for fi97 := 0; fi97 < 18; fi97++ {
                        var v98 int = (v69 - -9223372036854775808)
                        if true {
                            v69 = (v98 | (^(fi97)))
                            v99 := GenS9{F0: 42, F1: true, F2: int64(v69)}
                            var v100 float32 = float32((1.5 * -2.25))
                            _ = v100
                            _ = v99
                        }
                        _ = fi97
                        _ = v98
                    }
                    for fi101 := 0; fi101 < 18; fi101++ {
                        if true {
                            v102 := GenS10{F0: 42}
                            _ = v102
                        }
                        v69 += -735
                        switch v69 {
                        case 1:
                            v69 = v69
                        case 321:
                            v103 := GenS7{F0: 127, F1: 0}
                            _ = v103
                        case 127:
                            switch v2 {
                            case "":
                                var v104 uint16 = 42
                                _ = v104
                            case "\n\t":
                                fmt.Println(v69, v2, fi101)
                            case "foo bar":
                                if (1 <= -204) {
                                    v69 = v69
                                } else {
                                    delete(v3, "a")
                                    if (true && ((42 < -734) || false)) {
                                        if (9223372036854775807 <= 1) {
                                            fmt.Println(v69, v69, v2)
                                            for fi105 := 0; fi105 < 6; fi105++ {
                                                var v106 string = "\n\t"
                                                fmt.Println(v69)
                                                var v107 int16 = 919
                                                fmt.Println(v106, fi105)
                                                _ = fi105
                                                _ = v106
                                                _ = v107
                                            }
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if (true && true) {
                                                v108 := GenS6{F0: 42}
                                                switch (((v2 + ((("a" + (v2 + (("a" + v2) + v2))) + "foo bar") + "foo bar")) + v2) + "hello") {
                                                case "a":
                                                    delete(v3, (((v2 + v2) + "a") + v2))
                                                case "":
                                                    for fi109 := 0; fi109 < 2; fi109++ {
                                                        if (true || true) {
                                                            v69 = v69
                                                        } else {
                                                            switch (fi101 ^ (fi101 / ((128) | 1))) {
                                                            case -155:
                                                                fmt.Println(v69, fi101)
                                                            default:
                                                                v69 = fi109
                                                            }
                                                            v69 = fi101
                                                        }
                                                        fmt.Println(v2)
                                                        switch v2 {
                                                        case "\n\t":
                                                            var v110 uint64 = 0
                                                            _ = v110
                                                        case "":
                                                            fmt.Println(fi101)
                                                        case "hello":
                                                            if false {
                                                                v2 = ((v2 + v2) + (("hello" + (v2 + v2)) + "\n\t"))
                                                                fmt.Println(fi101)
                                                                var v111 int8 = 100
                                                                _ = v111
                                                            }
                                                        default:
                                                            if true {
                                                                var v112 int = 0
                                                                var v113 float32 = 0.0
                                                                _ = v112
                                                                _ = v113
                                                            }
                                                        }
                                                        _ = fi109
                                                    }
                                                case "foo bar":
                                                    var v114 uint64 = 18446744073709551615
                                                    _ = v114
                                                default:
                                                    if (!(!false)) {
                                                        if (true || (!false)) {
                                                            switch "a" {
                                                            case "foo bar":
                                                                if false {
                                                                    for fi115 := 0; fi115 < 15; fi115++ {
                                                                        v69 += 0
                                                                        var v116 uint16 = 335
                                                                        _ = fi115
                                                                        _ = v116
                                                                    }
                                                                }
                                                            case "hello":
                                                                var v117 int = (fi101 / ((198) | 1))
                                                                _ = v117
                                                            case "\n\t":
                                                                if (-9223372036854775808 != (v69 ^ (v69 << ((0) & 63)))) {
                                                                    v69 = 10
                                                                }
                                                            }
                                                            if (!true) {
                                                                v69 += (v69 - (v69 >> ((2) & 63)))
                                                            } else {
                                                                v69 = -373
                                                                delete(v3, "foo bar")
                                                            }
                                                            switch fi101 {
                                                            case 1:
                                                                v69 += (-(fi101))
                                                            case 128:
                                                                var v118 uint32 = 882
                                                                _ = v118
                                                            default:
                                                                if false {
                                                                    v119 := GenS11{F0: 42, F1: (52 == 0), F2: 42}
                                                                    fmt.Println(v69, v2)
                                                                    v3[("hello" + "hello")] = "hello"
                                                                    _ = v119
                                                                }
                                                            }
                                                        }
                                                        switch fi101 {
                                                        case 856:
                                                            v69 = (^(v69))
                                                        case 32768:
                                                            for fi120 := 0; fi120 < 7; fi120++ {
                                                                if ((1 >= 255) || ((-2.25 != (((0.1 / ((-2.25 / 1.0) * (-2.25 / 0.1))) - float64(0.0)) + 3.14159)) || false)) {
                                                                    v2 = "a"
                                                                }
                                                                if true {
                                                                    v69 = (^(v69))
                                                                    if false {
                                                                        fmt.Println(v69, fi120)
                                                                        v121 := GenS12{F0: 201}
                                                                        _ = v121
                                                                    } else {
                                                                        v69 += v69
                                                                    }
                                                                }
                                                                switch (fi101 % ((fi120) | 1)) {
                                                                case 128:
                                                                    v2 = v2
                                                                case -88:
                                                                    v69 += fi120
                                                                default:
                                                                    switch (fi120 >> ((243) & 63)) {
                                                                    case 42:
                                                                        switch (v69 << ((719) & 63)) {
                                                                        case 1:
                                                                            v69 = 32767
                                                                        case 42:
                                                                            if true {
                                                                                var v122 uint32 = 2147483647
                                                                                _ = len(v3)
                                                                                _ = v122
                                                                            } else {
                                                                                fmt.Println(v69)
                                                                                v69 = 426
                                                                                if ((0.1 + 0.1) > 0.1) {
                                                                                    fmt.Println(fi120)
                                                                                    var v123 uint32 = 2147483647
                                                                                    _ = v123
                                                                                }
                                                                            }
                                                                        case 310:
                                                                            v2 = v2
                                                                        default:
                                                                            v69 += (fi120 % ((fi120) | 1))
                                                                        }
                                                                    case 711:
                                                                        if (2147483648 >= 231) {
                                                                            if false {
                                                                                v69 += fi120
                                                                                v69 = 32768
                                                                                v124 := GenS10{F0: 2}
                                                                                _ = v124
                                                                            }
                                                                            v69 = 1
                                                                            fmt.Println(v69, fi120, v2)
                                                                        } else {
                                                                            if false {
                                                                                var v125 int8 = 125
                                                                                v2 = v2
                                                                                if (2 < 1) {
                                                                                    var v126 uint = 673
                                                                                    fmt.Println(v125, v2, fi120)
                                                                                    _ = v126
                                                                                }
                                                                                _ = v125
                                                                            }
                                                                            var v127 float32 = float32(1e10)
                                                                            v2 = ""
                                                                            v128 := GenS13{F0: 65535, F1: 255, F2: 1}
                                                                            _ = v127
                                                                            _ = v128
                                                                        }
                                                                    }
                                                                }
                                                                if ((128 > 255) && true) {
                                                                    if false {
                                                                        var v129 uint8 = 1
                                                                        v1 = append(v1, 952)
                                                                        _ = v129
                                                                    }
                                                                    v130 := GenS2{F0: 850, F1: 101}
                                                                    if ((!(0.0 <= ((((((3.14159 - (-2.25 + float32(3.14159))) - float32(-2.25)) + (1e10 + (float32(1.5) * (float32(0.1) * 1.5)))) - (0.1 - 0.1)) * 1.5) - 1.5))) && (!(-358 <= -9223372036854775808))) {
                                                                        fmt.Println(v69, v2)
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 2147483648
                                                                        }
                                                                        var v131 uint = 1
                                                                        _ = v131
                                                                    }
                                                                    _ = v130
                                                                } else {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 0
                                                                    }
                                                                    v132 := v3["foo bar"]
                                                                    fmt.Println(v69)
                                                                    _ = v132
                                                                }
                                                                if true {
                                                                    switch -237 {
                                                                    case 78:
                                                                        v69 = 323
                                                                    case 0:
                                                                        var v133 uint16 = 500
                                                                        _ = v133
                                                                    case 2147483647:
                                                                        v69 += -316
                                                                    default:
                                                                        var v134 uint16 = 255
                                                                        _ = v134
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                                _ = fi120
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v108
                                            }
                                        }
                                    }
                                }
                            default:
                                v69 -= (v69 >> ((149) & 63))
                            }
                        default:
                            for fi135 := 0; fi135 < 3; fi135++ {
                                var v136 int64 = -1
                                _ = fi135
                                _ = v136
                            }
                        }
                        _ = fi101
                    }
                }
                var v137 int32 = -2147483648
                _ = v137
            } else {
                if len(v1) > 0 {
                    v1[0] = 1
                }
                for fi138 := 0; fi138 < 2; fi138++ {
                    for fi139 := 0; fi139 < 10; fi139++ {
                        delete(v3, "")
                        _ = fi139
                    }
                    var v140 int8 = 0
                    for fi141 := 0; fi141 < 10; fi141++ {
                        v140 += v140
                        v142 := GenS14{F0: 2, F1: -32768, F2: 720}
                        fmt.Println(v2, v69)
                        _ = fi141
                        _ = v142
                    }
                    _ = fi138
                    _ = v140
                }
                fmt.Println(v69, v69)
            }
        }
        var v143 int = (v69 + 128)
        _ = v143
    }
    delete(v3, ("a" + ((v2 + ((v2 + "foo bar") + v2)) + v2)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v69
}
