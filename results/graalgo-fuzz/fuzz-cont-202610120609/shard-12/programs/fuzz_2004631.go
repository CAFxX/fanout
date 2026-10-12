package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint8; F1 uint8; F2 bool }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 int8; F2 int16 }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 bool; F1 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 string; F1 int64; F2 int64 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint8; F1 uint16 }
type GenS16 struct { F0 uint8; F1 int8; F2 uint16 }

func main() {
    var v1 uint = 42
    v2 := []int{4294967295, 39}
    v3 := make(map[uint32]bool)
    var v4 uint = v1
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(v4, v1, fi5)
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 int8 = 0
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS1{F0: -1, F1: 2147483647}
    switch -135 {
    case 0:
        for fi9 := 0; fi9 < 11; fi9++ {
            if len(v2) > 0 {
                v2[0] = 9223372036854775807
            }
            _ = fi9
        }
    case 42:
        var v10 float32 = (0.0 + 0.0)
        _ = v10
    }
    var v11 uint32 = 192
    _ = len(v2)
    v12 := GenS2{F0: 2147483647}
    if (1.5 == ((3.14159 / float32(float64(1e10))) / 1e10)) {
        v1 -= (v1 * (^(v4)))
        if true {
            v13 := GenS1{F0: -128, F1: (^(v11))}
            if true {
                switch 9223372036854775807 {
                case 4294967295:
                    delete(v3, (v11 | (v11 / ((992) | 1))))
                default:
                    if false {
                        switch 719 {
                        case 42:
                            v14 := GenS3{F0: -1, F1: false}
                            _ = v14
                        default:
                            switch (("\n\t" + ("\n\t" + "foo bar")) + "a") {
                            case "a":
                                switch ("a" + ("hello" + "")) {
                                case "a":
                                    fmt.Println(v11)
                                case "":
                                    if false {
                                        if true {
                                            v11 = v11
                                        } else {
                                            v4 = v1
                                        }
                                    } else {
                                        fmt.Println(v11, v4, v4)
                                        fmt.Println(v11)
                                        var v15 string = "foo bar"
                                        _ = v15
                                    }
                                }
                            case "hello":
                                if true {
                                    if false {
                                        _ = len(v2)
                                        v16 := v3[352]
                                        _ = v16
                                    } else {
                                        var v17 uint = v1
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v17
                                    }
                                    for fi18 := 0; fi18 < 9; fi18++ {
                                        delete(v3, (v11 * v11))
                                        switch ("hello" + ("foo bar" + "\n\t")) {
                                        case "a":
                                            v3[(^(v11))] = false
                                        }
                                        _ = fi18
                                    }
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            case "foo bar":
                                for fi19 := 0; fi19 < 19; fi19++ {
                                    if false {
                                        v4 += (v4 * (^(v4)))
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch ((("\n\t" + "") + "\n\t") + "foo bar") {
                                        case "a":
                                            var v20 int = 2147483647
                                            _ = v20
                                        case "\n\t":
                                            v4 = 42
                                        case "hello":
                                            v4 = 0
                                        default:
                                            v11 -= v11
                                        }
                                        switch (fi19 & (^(fi19))) {
                                        case 32768:
                                            var v21 string = "foo bar"
                                            _ = v21
                                        case 2147483647:
                                            for fi22 := 0; fi22 < 18; fi22++ {
                                                var v23 bool = true
                                                fmt.Println(fi22)
                                                _ = fi22
                                                _ = v23
                                            }
                                        case -949:
                                            fmt.Println(v11)
                                        default:
                                            v11 = 146
                                        }
                                    }
                                    for fi24 := 0; fi24 < 1; fi24++ {
                                        if true {
                                            v1 = (v4 >> (((^(v1))) & 63))
                                            fmt.Println(v1, fi19, fi24)
                                        } else {
                                            v1 = 42
                                            v25 := v3[v11]
                                            _ = v25
                                        }
                                        v1 -= v4
                                        v4 -= v1
                                        fmt.Println(v11)
                                        v26 := GenS2{F0: int(v4)}
                                        _ = fi24
                                        _ = v26
                                    }
                                    var v27 uint = (v1 % ((uint(v11)) | 1))
                                    _ = fi19
                                    _ = v27
                                }
                            default:
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    var v29 uint8 = 0
                                    switch "a" {
                                    case "foo bar":
                                        var v30 uint32 = v11
                                        _ = v30
                                    case "":
                                        switch fi28 {
                                        case -652:
                                            v29 -= 2
                                        case -586:
                                            var v31 uint32 = uint32(v29)
                                            _ = v31
                                        default:
                                            if false {
                                                if false {
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi28 / ((2147483647) | 1))
                                                    }
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = (-(fi28))
                                                    }
                                                }
                                            } else {
                                                for fi32 := 0; fi32 < 8; fi32++ {
                                                    v29 = (v29 / ((v29) | 1))
                                                    _ = fi32
                                                }
                                                for fi33 := 0; fi33 < 6; fi33++ {
                                                    var v34 int32 = 42
                                                    var v35 uint64 = uint64(v11)
                                                    switch fi28 {
                                                    case 127:
                                                        if false {
                                                            fmt.Println(v11)
                                                            var v36 uint64 = (v35 / (((v35 % (((^(v35))) | 1))) | 1))
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            var v37 uint32 = v11
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            v29 -= (v29 ^ 2)
                                                            var v38 float64 = (1e10 + float64((0.0 - 1.5)))
                                                            v38 = (v38 + (v38 * -2.25))
                                                            _ = v38
                                                        }
                                                    case 4294967295:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        var v39 int16 = 42
                                                        _ = v39
                                                    }
                                                    _ = fi33
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                        }
                                    }
                                    _ = fi28
                                    _ = v29
                                }
                            }
                        }
                        if (!(("" + "\n\t") >= ("\n\t" + "\n\t"))) {
                            _ = len(v2)
                            var v40 int8 = 100
                            v11 = (v11 + 641)
                            _ = v40
                        }
                        switch "hello" {
                        case "\n\t":
                            if ("a" < (("hello" + "") + "hello")) {
                                for fi41 := 0; fi41 < 6; fi41++ {
                                    if false {
                                        v42 := GenS4{F0: 1, F1: 1, F2: true}
                                        v43 := GenS3{F0: -128, F1: (true || false)}
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v1 += 911
                                        v1 = v1
                                    }
                                    fmt.Println(v1)
                                    v3[v11] = (!(!false))
                                    _ = fi41
                                }
                                var v44 uint64 = 42
                                _ = v44
                            }
                        case "hello":
                            if true {
                                v11 = v11
                                v1 = v1
                                v11 -= 613
                            } else {
                                v45 := GenS2{F0: -77}
                                var v46 uint32 = 2147483647
                                if true {
                                    switch 128 {
                                    case 127:
                                        v2 = append(v2, -2147483648)
                                    case -71:
                                        v4 -= (v4 - v4)
                                    default:
                                        if (1 > 42) {
                                            var v47 int = 127
                                            v48 := GenS5{F0: -1, F1: 9223372036854775807}
                                            _ = v47
                                            _ = v48
                                        } else {
                                            for fi49 := 0; fi49 < 15; fi49++ {
                                                v11 = 65535
                                                switch "hello" {
                                                case "\n\t":
                                                    for fi50 := 0; fi50 < 8; fi50++ {
                                                        var v51 int8 = 0
                                                        v51 -= -128
                                                        v51 -= v51
                                                        _ = fi50
                                                        _ = v51
                                                    }
                                                case "hello":
                                                    var v52 int = 127
                                                    _ = v52
                                                default:
                                                    if true {
                                                        if true {
                                                            v11 = 4294967295
                                                            for fi53 := 0; fi53 < 7; fi53++ {
                                                                switch 0 {
                                                                case 601:
                                                                    _ = len(v2)
                                                                case -2147483648:
                                                                    switch fi49 {
                                                                    case 32767:
                                                                        v54 := GenS6{F0: 202}
                                                                        _ = v54
                                                                    case -989:
                                                                        v11 = (v11 ^ (v46 & v11))
                                                                    case 32768:
                                                                        switch fi53 {
                                                                        case -1:
                                                                            fmt.Println(v1)
                                                                        }
                                                                    default:
                                                                        v3[389] = false
                                                                    }
                                                                }
                                                                fmt.Println(v11)
                                                                var v55 int32 = -2147483648
                                                                _ = fi53
                                                                _ = v55
                                                            }
                                                            v1 = v4
                                                        }
                                                        if (33 >= 255) {
                                                            if ((66 == 255) || (-32768 >= 32767)) {
                                                                v4 = (v4 / ((255) | 1))
                                                                var v56 uint32 = (v11 ^ 60)
                                                                _ = len(v3)
                                                                _ = v56
                                                            }
                                                        }
                                                        switch fi49 {
                                                        case 958:
                                                            if (2147483647 >= -2147483648) {
                                                                _ = len(v3)
                                                                v1 += (v4 ^ (v4 - (v4 / ((65535) | 1))))
                                                            } else {
                                                                if true {
                                                                    v57 := GenS6{F0: v4}
                                                                    if ((v4 == 0) || false) {
                                                                        var v58 string = "a"
                                                                        switch ((v58 + (v58 + (v58 + v58))) + v58) {
                                                                        case "a":
                                                                            v59 := v3[2147483647]
                                                                            _ = v59
                                                                        case "\n\t":
                                                                            v4 += (^(v1))
                                                                        case "foo bar":
                                                                            switch (fi49 * 0) {
                                                                            case 1:
                                                                                v60 := GenS7{F0: int32(v46)}
                                                                                _ = v60
                                                                            case -121:
                                                                                v11 = (^(v46))
                                                                            default:
                                                                                if false {
                                                                                    _ = len(v3)
                                                                                    v1 += v1
                                                                                    fmt.Println(fi49)
                                                                                }
                                                                            }
                                                                        default:
                                                                            v11 = v11
                                                                        }
                                                                        var v61 int16 = 726
                                                                        _ = v58
                                                                        _ = v61
                                                                    }
                                                                    _ = v57
                                                                } else {
                                                                    v1 = (v1 & (v4 % ((2) | 1)))
                                                                    fmt.Println(v1, v11)
                                                                }
                                                                var v62 uint8 = 229
                                                                _ = v62
                                                            }
                                                        case -902:
                                                            if (true && (!false)) {
                                                                var v63 bool = true
                                                                v1 = v4
                                                                _ = v63
                                                            }
                                                        case -2147483648:
                                                            v64 := v3[v46]
                                                            _ = v64
                                                        default:
                                                            if ((fi49 - (^(fi49))) <= -9223372036854775808) {
                                                                switch fi49 {
                                                                case -77:
                                                                    if false {
                                                                        if true {
                                                                            if false {
                                                                                v1 = (v1 + (v4 >> ((v4) & 63)))
                                                                                for fi65 := 0; fi65 < 20; fi65++ {
                                                                                    var v66 int64 = 9223372036854775807
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    _ = fi65
                                                                                    _ = v66
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v4)
                                                                                switch (^(fi49)) {
                                                                                case 292:
                                                                                    for fi67 := 0; fi67 < 14; fi67++ {
                                                                                        v46 = 67
                                                                                        v46 -= 88
                                                                                        _ = fi67
                                                                                    }
                                                                                case 32767:
                                                                                    var v68 int = (^(fi49))
                                                                                    _ = v68
                                                                                default:
                                                                                    var v69 uint = 255
                                                                                    _ = v69
                                                                                }
                                                                                if true {
                                                                                    var v70 bool = true
                                                                                    fmt.Println(v1, v1)
                                                                                    v70 = (true && false)
                                                                                    _ = v70
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v71 := GenS8{F0: 2}
                                                                            fmt.Println(v4, v1)
                                                                            fmt.Println(v11, fi49)
                                                                            _ = v71
                                                                        }
                                                                        v1 = v4
                                                                        var v72 int16 = 1000
                                                                        _ = v72
                                                                    }
                                                                case -645:
                                                                    for fi73 := 0; fi73 < 10; fi73++ {
                                                                        v74 := GenS6{F0: (v4 & 4294967295)}
                                                                        switch ("hello" + ("" + ((("a" + ("a" + "a")) + "hello") + ""))) {
                                                                        case "":
                                                                            v75 := GenS9{F0: 0, F1: 78, F2: -32768}
                                                                            _ = v75
                                                                        case "hello":
                                                                            v76 := GenS3{F0: -1, F1: ((!(1 == 804)) && (true || false))}
                                                                            _ = v76
                                                                        case "foo bar":
                                                                            fmt.Println(v1, v4, fi49)
                                                                        }
                                                                        _ = fi73
                                                                        _ = v74
                                                                    }
                                                                default:
                                                                    switch fi49 {
                                                                    case -1:
                                                                        var v77 int8 = -1
                                                                        _ = v77
                                                                    }
                                                                }
                                                                v78 := GenS10{F0: v46, F1: 1000}
                                                                var v79 float64 = 3.14159
                                                                _ = v78
                                                                _ = v79
                                                            } else {
                                                                for fi80 := 0; fi80 < 4; fi80++ {
                                                                    v1 = 255
                                                                    v11 += 455
                                                                    v3[v11] = (-1 <= int64(v46))
                                                                    _ = fi80
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = (fi49 << (((v4 * 322)) & 63))
                                                }
                                                _ = fi49
                                            }
                                            fmt.Println(v1)
                                        }
                                    }
                                    v81 := GenS11{F0: true, F1: 127}
                                    _ = v81
                                }
                                _ = v45
                                _ = v46
                            }
                        case "a":
                            var v82 string = "foo bar"
                            _ = v82
                        default:
                            for fi83 := 0; fi83 < 4; fi83++ {
                                var v84 uint = v1
                                switch (fi83 | (fi83 & (fi83 << ((v84) & 63)))) {
                                case 128:
                                    v1 = v1
                                }
                                _ = fi83
                                _ = v84
                            }
                        }
                    }
                }
                delete(v3, v11)
                switch "a" {
                case "foo bar":
                    if true {
                        for fi85 := 0; fi85 < 11; fi85++ {
                            var v86 int32 = 1
                            if (1000 == -167) {
                                var v87 int32 = (v86 + (v86 / ((183) | 1)))
                                for fi88 := 0; fi88 < 20; fi88++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi88
                                }
                                v89 := GenS5{F0: 32767, F1: fi85}
                                _ = v87
                                _ = v89
                            } else {
                                var v90 string = ""
                                var v91 uint16 = 0
                                _ = v90
                                _ = v91
                            }
                            _ = fi85
                            _ = v86
                        }
                        v4 += 2
                    } else {
                        var v92 uint8 = 255
                        var v93 uint = v1
                        if ((!(true && true)) && true) {
                            if len(v2) > 0 {
                                v2[0] = 32768
                            }
                            if true {
                                for fi94 := 0; fi94 < 5; fi94++ {
                                    var v95 uint8 = 42
                                    _ = fi94
                                    _ = v95
                                }
                                if len(v2) > 0 {
                                    v2[0] = -1
                                }
                            } else {
                                var v96 int16 = -1
                                for fi97 := 0; fi97 < 10; fi97++ {
                                    for fi98 := 0; fi98 < 16; fi98++ {
                                        if len(v2) > 0 {
                                            v2[0] = -2147483648
                                        }
                                        _ = fi98
                                    }
                                    if (-1 == 1) {
                                        v93 -= (v4 - 334)
                                    } else {
                                        for fi99 := 0; fi99 < 10; fi99++ {
                                            switch (fi99 * (^(fi99))) {
                                            case -365:
                                                switch "\n\t" {
                                                case "hello":
                                                    v4 += uint(v92)
                                                case "foo bar":
                                                    if false {
                                                        var v100 bool = ((0.1 * 0.1) == ((0.0 * (0.0 + 0.0)) / 1.0))
                                                        _ = v100
                                                    } else {
                                                        v1 = (v93 | v93)
                                                        fmt.Println(v92)
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v101 int64 = 316
                                                            _ = v101
                                                        case "foo bar":
                                                            var v102 uint32 = v11
                                                            _ = v102
                                                        }
                                                    }
                                                case "":
                                                    v92 = (v92 % ((uint8(v11)) | 1))
                                                default:
                                                    v103 := GenS11{F0: true, F1: 255}
                                                    _ = v103
                                                }
                                            case 32767:
                                                switch "foo bar" {
                                                case "":
                                                    var v104 string = "foo bar"
                                                    _ = v104
                                                case "foo bar":
                                                    delete(v3, v11)
                                                default:
                                                    v105 := GenS12{F0: "\n\t"}
                                                    _ = v105
                                                }
                                            case 32768:
                                                var v106 bool = (!true)
                                                _ = v106
                                            default:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            _ = fi99
                                        }
                                    }
                                    _ = fi97
                                }
                                for fi107 := 0; fi107 < 4; fi107++ {
                                    if (("foo bar" + "a") <= "a") {
                                        v92 -= v92
                                    }
                                    v92 += 254
                                    for fi108 := 0; fi108 < 3; fi108++ {
                                        if true {
                                            v92 = v92
                                        } else {
                                            if (!(v4 <= 4294967295)) {
                                                var v109 int8 = 42
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v110 int = fi108
                                                _ = v109
                                                _ = v110
                                            } else {
                                                var v111 int8 = 1
                                                if ((("hello" + ("foo bar" + ("a" + "a"))) + "foo bar") == "\n\t") {
                                                    var v112 uint = v1
                                                    _ = v112
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = int(v111)
                                                }
                                                var v113 float64 = (-2.25 * 1.5)
                                                v113 = (1e10 + ((float64(-2.25) / (v113 * (v113 - (1e10 / -2.25)))) / v113))
                                                v113 -= v113
                                                _ = v111
                                                _ = v113
                                            }
                                            var v114 int8 = 42
                                            fmt.Println(v93)
                                            _ = v114
                                        }
                                        switch 4294967295 {
                                        case -2147483648:
                                            var v115 string = (("foo bar" + ("hello" + "")) + "\n\t")
                                            _ = v115
                                        }
                                        v96 = (-(v96))
                                        _ = fi108
                                    }
                                    _ = fi107
                                }
                                _ = v96
                            }
                            for fi116 := 0; fi116 < 19; fi116++ {
                                for fi117 := 0; fi117 < 18; fi117++ {
                                    var v118 int8 = -128
                                    var v119 int16 = 42
                                    if true {
                                        switch (fi116 & fi117) {
                                        case -2147483648:
                                            var v120 uint16 = 32767
                                            _ = v120
                                        case 32767:
                                            v92 = v92
                                        }
                                        if ((true || (fi116 >= fi117)) || ((int64(v118) >= 42) && false)) {
                                            v121 := GenS10{F0: v11, F1: 1000}
                                            _ = v121
                                        } else {
                                            switch ("hello" + ("foo bar" + ("" + "foo bar"))) {
                                            case "hello":
                                                var v122 float64 = 3.14159
                                                _ = v122
                                            case "\n\t":
                                                if (-2.25 > 3.14159) {
                                                    v93 += 1
                                                }
                                            case "foo bar":
                                                v1 = v1
                                            }
                                            v123 := GenS8{F0: 32767}
                                            _ = v123
                                        }
                                        var v124 int8 = v118
                                        _ = v124
                                    } else {
                                        var v125 uint8 = 1
                                        if false {
                                            var v126 string = ("hello" + "")
                                            v127 := GenS13{F0: (v126 + "hello"), F1: -9223372036854775808, F2: 196}
                                            _ = v126
                                            _ = v127
                                        } else {
                                            v128 := GenS6{F0: v93}
                                            v129 := GenS7{F0: 0}
                                            v130 := GenS14{F0: fi117}
                                            _ = v128
                                            _ = v129
                                            _ = v130
                                        }
                                        _ = v125
                                    }
                                    _ = fi117
                                    _ = v118
                                    _ = v119
                                }
                                v1 = v1
                                v1 = (v93 << ((1) & 63))
                                _ = fi116
                            }
                        }
                        _ = v92
                        _ = v93
                    }
                case "":
                    v131 := GenS15{F0: 120, F1: 109}
                    _ = v131
                case "\n\t":
                    v11 = (v11 & (v11 + v11))
                default:
                    if (3.14159 != 1.5) {
                        v132 := GenS13{F0: "a", F1: int64(v1), F2: 9223372036854775807}
                        var v133 int16 = 32767
                        _ = v132
                        _ = v133
                    }
                }
            } else {
                var v134 int64 = 42
                if true {
                    v1 = v4
                    v135 := GenS8{F0: 378}
                    _ = v135
                }
                if ((1 >= -128) && (!false)) {
                    var v136 string = "hello"
                    _ = v136
                }
                _ = v134
            }
            _ = v13
        }
    } else {
        var v137 int8 = -128
        _ = v137
    }
    v138 := GenS14{F0: -1}
    switch -155 {
    case 4294967295:
        v139 := GenS2{F0: 32768}
        _ = v139
    case -344:
        v140 := v3[(v11 & (v11 & (v11 - (v11 - 192))))]
        _ = v140
    case 128:
        var v141 uint = v1
        _ = v141
    }
    var v142 uint8 = 205
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        if false {
            var v143 uint32 = (v11 << ((v4) & 63))
            v142 -= v142
            _ = v143
        }
        _ = len(v2)
    }
    var v144 uint16 = 2
    var v145 int64 = -9223372036854775808
    var v146 int64 = v145
    for fi147 := 0; fi147 < 6; fi147++ {
        if (-1 <= 42) {
            v148 := GenS16{F0: (v142 % ((v142) | 1)), F1: 111, F2: 0}
            var v149 float32 = 1e10
            v145 -= (v145 + (v145 % ((42) | 1)))
            _ = v148
            _ = v149
        } else {
            var v150 uint64 = 1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v150 = (v150 & (v150 >> (((v4 | (^(v1)))) & 63)))
            _ = v150
        }
        var v151 int64 = (v146 - (v145 + 42))
        _ = fi147
        _ = v151
    }
    v152 := GenS7{F0: 148}
    var v153 uint8 = uint8(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v142:", v142)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("v146:", v146)
    fmt.Println("v153:", v153)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v138
    _ = v142
    _ = v144
    _ = v145
    _ = v146
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
