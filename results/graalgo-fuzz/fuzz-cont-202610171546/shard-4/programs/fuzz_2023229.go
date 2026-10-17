package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 int; F2 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8; F1 uint16 }

func main() {
    var v1 string = ("\n\t" + "")
    v2 := make(map[int]uint32)
    v3 := []string{"a", "foo bar"}
    if (!(true || (!false))) {
        v1 = (v1 + v1)
        if len(v3) > 0 {
            v3[0] = v1
        }
        v1 = ("foo bar" + v1)
    }
    for fi4 := 0; fi4 < 2; fi4++ {
        fmt.Println(v1, fi4)
        _ = fi4
    }
    if true {
        if (18446744073709551615 >= 0) {
            if (!true) {
                if (0 == 2) {
                    v1 = v1
                } else {
                    fmt.Println(v1)
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    if true {
                        fmt.Println(v1)
                        v1 = v1
                        if true {
                            v1 = (v1 + "")
                            for fi5 := 0; fi5 < 4; fi5++ {
                                for fi6 := 0; fi6 < 3; fi6++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v1 = (((v1 + v1) + "a") + "foo bar")
                                    var v7 bool = (3.14159 <= (float64((float32(float64(3.14159)) - 0.0)) / (float64(float32((0.1 + 1e10))) + 1e10)))
                                    _ = fi6
                                    _ = v7
                                }
                                v8 := v2[(fi5 ^ (^(fi5)))]
                                _ = fi5
                                _ = v8
                            }
                        }
                    } else {
                        if (((v1 + (v1 + (v1 + v1))) + ("a" + (v1 + v1))) > "foo bar") {
                            for fi9 := 0; fi9 < 19; fi9++ {
                                switch (fi9 % (((fi9 - fi9)) | 1)) {
                                case 127:
                                    fmt.Println(fi9)
                                case -391:
                                    if (!false) {
                                        v1 = "\n\t"
                                    }
                                default:
                                    v1 = v1
                                }
                                for fi10 := 0; fi10 < 6; fi10++ {
                                    switch (((("foo bar" + "") + "") + ("hello" + ("" + ("" + "hello")))) + "hello") {
                                    case "foo bar":
                                        var v11 float32 = 0.0
                                        _ = v11
                                    default:
                                        var v12 float64 = ((1e10 - (-2.25 * ((0.0 * -2.25) + 0.0))) - -2.25)
                                        if len(v3) > 0 {
                                            v3[0] = (v1 + v1)
                                        }
                                        v12 = (v12 * v12)
                                        _ = v12
                                    }
                                    v13 := GenS1{F0: 128}
                                    fmt.Println(fi10, fi9)
                                    _ = fi10
                                    _ = v13
                                }
                                _ = fi9
                            }
                            if (!true) {
                                switch 42 {
                                case 969:
                                    var v14 uint32 = 0
                                    _ = v14
                                }
                            }
                            delete(v2, 127)
                        } else {
                            v1 = v1
                            if (!(1 > 1)) {
                                fmt.Println(v1)
                                var v15 uint8 = 1
                                _ = v15
                            }
                        }
                        var v16 uint16 = 498
                        for fi17 := 0; fi17 < 1; fi17++ {
                            fmt.Println(v16, v16, fi17)
                            switch (^(fi17)) {
                            case -234:
                                if true {
                                    fmt.Println(v1, v16, v1)
                                    for fi18 := 0; fi18 < 4; fi18++ {
                                        v2[9223372036854775807] = 2
                                        v1 = "hello"
                                        v16 = v16
                                        _ = fi18
                                    }
                                    for fi19 := 0; fi19 < 19; fi19++ {
                                        switch "hello" {
                                        case "hello":
                                            switch "hello" {
                                            case "foo bar":
                                                v16 -= v16
                                            default:
                                                if true {
                                                    if false {
                                                        var v20 uint = 4294967295
                                                        if false {
                                                            v1 = "\n\t"
                                                            v16 -= v16
                                                            switch "" {
                                                            case "\n\t":
                                                                v3 = append(v3, "\n\t")
                                                                v21 := GenS1{F0: 42}
                                                                _ = v21
                                                            case "foo bar":
                                                                if (uint64(v16) > 1) {
                                                                    v16 = (v16 | v16)
                                                                    v16 = v16
                                                                }
                                                            case "a":
                                                                if false {
                                                                    fmt.Println(fi19, v1, v16)
                                                                    var v22 uint = uint(v16)
                                                                    _ = v22
                                                                } else {
                                                                    v20 = (v20 * (v20 >> (((v20 % ((v20) | 1))) & 63)))
                                                                    var v23 uint32 = 759
                                                                    var v24 uint64 = 18446744073709551615
                                                                    var v25 uint32 = uint32(v24)
                                                                    switch -9223372036854775808 {
                                                                    case 9223372036854775807:
                                                                        if ((3.14159 + 1.5) < (1e10 - (-2.25 / 1.0))) {
                                                                            switch (fi19 << (((v20 + (v20 - (^(v20))))) & 63)) {
                                                                            case 127:
                                                                                v26 := GenS2{F0: 0}
                                                                                _ = v26
                                                                            case 1:
                                                                                v16 = (v16 + (v16 + (v16 ^ (^(v16)))))
                                                                                switch 163 {
                                                                                case 32768:
                                                                                    switch (^(fi17)) {
                                                                                    case 42:
                                                                                        fmt.Println(fi17)
                                                                                    case 127:
                                                                                        v27 := GenS2{F0: 42}
                                                                                        switch fi19 {
                                                                                        case 436:
                                                                                            v23 -= v25
                                                                                        case 9223372036854775807:
                                                                                            var v28 uint8 = 255
                                                                                            _ = v28
                                                                                        default:
                                                                                            var v29 float64 = -2.25
                                                                                            _ = v29
                                                                                        }
                                                                                        if (719 < 366) {
                                                                                            switch (fi19 + (fi17 ^ fi17)) {
                                                                                            case 4294967295:
                                                                                                var v30 int32 = 2147483647
                                                                                                _ = v30
                                                                                            case 128:
                                                                                                delete(v2, fi17)
                                                                                            case 32767:
                                                                                                v31 := GenS1{F0: 255}
                                                                                                _ = v31
                                                                                            }
                                                                                            fmt.Println(fi19, v20, v20)
                                                                                        } else {
                                                                                            v2[-747] = v25
                                                                                        }
                                                                                        _ = v27
                                                                                    default:
                                                                                        v16 = 1
                                                                                    }
                                                                                case 4294967295:
                                                                                    var v32 int16 = 813
                                                                                    _ = v32
                                                                                case 42:
                                                                                    if ((-32768 != 295) && false) {
                                                                                        var v33 string = (v1 + ((v1 + v1) + ((("\n\t" + "") + "") + ("\n\t" + ("hello" + ("\n\t" + "a"))))))
                                                                                        _ = v33
                                                                                    }
                                                                                    fmt.Println(v24, v1, v16)
                                                                                default:
                                                                                    switch (fi19 - (fi17 + 128)) {
                                                                                    case 128:
                                                                                        _ = len(v3)
                                                                                    case 9223372036854775807:
                                                                                        fmt.Println(v1, v16, v24)
                                                                                    case 127:
                                                                                        var v34 int8 = 127
                                                                                        _ = v34
                                                                                    }
                                                                                }
                                                                                v16 = 42
                                                                            case -579:
                                                                                v16 -= v16
                                                                            default:
                                                                                v35 := GenS3{F0: int8(v24)}
                                                                                _ = v35
                                                                            }
                                                                            if (42 <= 151) {
                                                                                switch (fi19 << (((v20 | v20)) & 63)) {
                                                                                case -2147483648:
                                                                                    var v36 int = (fi17 + (fi17 - fi17))
                                                                                    v1 = "\n\t"
                                                                                    fmt.Println(fi17, v25)
                                                                                    _ = v36
                                                                                case 327:
                                                                                    switch 9223372036854775807 {
                                                                                    case -1:
                                                                                        v1 = "foo bar"
                                                                                    default:
                                                                                        if false {
                                                                                            var v37 int16 = int16(v24)
                                                                                            v38 := GenS4{F0: uint64(v23), F1: 2147483647}
                                                                                            _ = v37
                                                                                            _ = v38
                                                                                        } else {
                                                                                            var v39 float32 = 3.14159
                                                                                            if (!(!false)) {
                                                                                                fmt.Println(v39, v24, v20)
                                                                                            }
                                                                                            v16 = 32767
                                                                                            _ = v39
                                                                                        }
                                                                                    }
                                                                                case 42:
                                                                                    if true {
                                                                                        if true {
                                                                                            if false {
                                                                                                v20 = (v20 ^ v20)
                                                                                                fmt.Println(v25, fi17)
                                                                                                if ((!((^(v25)) >= v23)) || ((v24 >> ((v20) & 63)) <= v24)) {
                                                                                                    v40 := GenS4{F0: v24, F1: 709}
                                                                                                    if false {
                                                                                                        switch "foo bar" {
                                                                                                        case "hello":
                                                                                                            _ = len(v2)
                                                                                                        case "":
                                                                                                            var v41 uint8 = 201
                                                                                                            _ = v41
                                                                                                        case "a":
                                                                                                            v42 := GenS4{F0: v24, F1: -237}
                                                                                                            v1 = v1
                                                                                                            _ = v42
                                                                                                        default:
                                                                                                            var v43 string = "hello"
                                                                                                            _ = v43
                                                                                                        }
                                                                                                        var v44 uint8 = 245
                                                                                                        fmt.Println(v20)
                                                                                                        _ = v44
                                                                                                    }
                                                                                                    _ = v40
                                                                                                }
                                                                                            } else {
                                                                                                delete(v2, (fi19 | fi19))
                                                                                                v23 = (^(v23))
                                                                                            }
                                                                                            v45 := GenS4{F0: 33, F1: -2147483648}
                                                                                            _ = v45
                                                                                        }
                                                                                        if (!(1e10 >= (-2.25 * 0.0))) {
                                                                                            v1 = "\n\t"
                                                                                        }
                                                                                        var v46 uint32 = (v23 * 542)
                                                                                        _ = v46
                                                                                    } else {
                                                                                        var v47 uint = v20
                                                                                        v16 = v16
                                                                                        _ = v47
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        var v48 bool = true
                                                                                        v25 = (^(v25))
                                                                                        v49 := GenS5{F0: (fi19 - (fi17 % ((fi17) | 1))), F1: 514}
                                                                                        v50 := GenS6{F0: (v23 | v25)}
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                        _ = v50
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v25, v25)
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v23
                                                                    _ = v24
                                                                    _ = v25
                                                                }
                                                            }
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = v1
                                                            }
                                                            v1 = v1
                                                        }
                                                        _ = v20
                                                    }
                                                } else {
                                                    fmt.Println(fi17, fi17)
                                                }
                                            }
                                        case "":
                                            switch v1 {
                                            case "foo bar":
                                                var v51 int16 = 1000
                                                _ = v51
                                            case "a":
                                                if false {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if (!(false && (1000 > -1))) {
                                                        var v52 bool = ((-41 < 0) && true)
                                                        _ = v52
                                                    } else {
                                                        var v53 float32 = 0.0
                                                        v16 += v16
                                                        switch (v1 + (((v1 + v1) + v1) + ((v1 + v1) + ((v1 + ((v1 + v1) + (v1 + ("hello" + v1)))) + v1)))) {
                                                        case "\n\t":
                                                            switch (fi19 & -10) {
                                                            case 0:
                                                                v54 := GenS7{F0: (-52 < 100), F1: (fi19 * 0), F2: 0}
                                                                _ = v54
                                                            case 969:
                                                                var v55 int16 = 1000
                                                                var v56 uint8 = 2
                                                                _ = v55
                                                                _ = v56
                                                            default:
                                                                switch (fi19 ^ (-(fi17))) {
                                                                case 0:
                                                                    v57 := GenS6{F0: 2}
                                                                    _ = v57
                                                                default:
                                                                    v1 = ""
                                                                }
                                                            }
                                                        case "hello":
                                                            if true {
                                                                _ = len(v3)
                                                                fmt.Println(fi19, v16)
                                                                if true {
                                                                    fmt.Println(fi19, v53, v1)
                                                                }
                                                            }
                                                        default:
                                                            var v58 bool = true
                                                            _ = v58
                                                        }
                                                        if (float64(((v53 - v53) + 1e10)) < 1e10) {
                                                            if (((((0.0 * 3.14159) - -2.25) * (0.1 + 1.5)) / (-2.25 + 1.5)) != 3.14159) {
                                                                v59 := GenS6{F0: 42}
                                                                var v60 int = fi17
                                                                fmt.Println(v53, fi19)
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            fmt.Println(v53, fi17, v1)
                                                            v16 = v16
                                                        }
                                                        _ = v53
                                                    }
                                                    if true {
                                                        if true {
                                                            var v61 int64 = 42
                                                            delete(v2, (fi19 ^ fi19))
                                                            _ = v61
                                                        }
                                                    } else {
                                                        if (fi19 >= -2147483648) {
                                                            v1 = ("" + v1)
                                                        }
                                                        v62 := GenS7{F0: (((42 > -1) || false) && false), F1: (fi19 * int(v16)), F2: 18446744073709551615}
                                                        fmt.Println(v16, fi17)
                                                        _ = v62
                                                    }
                                                }
                                            case "hello":
                                                var v63 string = (v1 + v1)
                                                _ = v63
                                            }
                                        }
                                        _ = fi19
                                    }
                                }
                            case 586:
                                var v64 float32 = (1.5 * float32(1.5))
                                _ = v64
                            case -9223372036854775808:
                                switch "foo bar" {
                                case "a":
                                    var v65 uint64 = 2
                                    _ = v65
                                case "\n\t":
                                    v16 = (v16 - 263)
                                default:
                                    for fi66 := 0; fi66 < 8; fi66++ {
                                        v67 := GenS6{F0: 2147483647}
                                        _ = fi66
                                        _ = v67
                                    }
                                }
                            default:
                                v1 = v1
                            }
                            _ = fi17
                        }
                        _ = v16
                    }
                }
                v68 := GenS4{F0: 0, F1: 1}
                _ = v68
            }
        } else {
            fmt.Println(v1)
            v2[9223372036854775807] = 2
        }
        v3 = append(v3, ("hello" + ""))
        v1 = ""
    } else {
        delete(v2, 32768)
        var v69 int64 = 9223372036854775807
        _ = v69
    }
    var v70 float32 = 3.14159
    switch (v1 + (v1 + "\n\t")) {
    case "a":
        for fi71 := 0; fi71 < 1; fi71++ {
            fmt.Println(fi71, v1)
            var v72 int16 = 0
            switch ("\n\t" + v1) {
            case "hello":
                v1 = ("a" + "\n\t")
            }
            _ = fi71
            _ = v72
        }
    }
    fmt.Println(v70, v70)
    v73 := GenS8{F0: 127}
    var v74 uint8 = 127
    var v75 uint32 = 2147483648
    if ("" > "") {
        if (!false) {
            var v76 float64 = 0.0
            v77 := GenS2{F0: 1}
            if (0 != (v75 | (v75 * 65535))) {
                if false {
                    var v78 float32 = float32(0.0)
                    v78 -= (float32((v76 * float64(-2.25))) + float32(((v76 * 0.1) + (float64(0.0) / (v76 / (v76 * (((0.0 + v76) - (0.0 / (v76 - v76))) * ((0.0 - v76) + (v76 * (v76 / float64(v78)))))))))))
                    switch 128 {
                    case -9223372036854775808:
                        var v79 uint16 = uint16(v75)
                        _ = v79
                    }
                    _ = v78
                }
                for fi80 := 0; fi80 < 15; fi80++ {
                    v81 := GenS3{F0: 0}
                    v74 -= 0
                    fmt.Println(v76)
                    _ = fi80
                    _ = v81
                }
            } else {
                v82 := GenS5{F0: 146, F1: 42}
                fmt.Println(v75, v76, v74)
                _ = v82
            }
            _ = v76
            _ = v77
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if false {
                v1 = "foo bar"
            } else {
                v83 := v2[2147483647]
                switch v1 {
                case "":
                    switch "\n\t" {
                    case "\n\t":
                        if (!false) {
                            _ = len(v2)
                            if true {
                                if (!false) {
                                    var v84 uint = 65535
                                    if true {
                                        v74 = 0
                                        v2[42] = v83
                                        for fi85 := 0; fi85 < 17; fi85++ {
                                            var v86 int64 = 0
                                            _ = fi85
                                            _ = v86
                                        }
                                    }
                                    v74 = v74
                                    _ = v84
                                }
                            } else {
                                var v87 int8 = -128
                                _ = v87
                            }
                            for fi88 := 0; fi88 < 15; fi88++ {
                                fmt.Println(v70, v83)
                                v89 := GenS9{F0: 9223372036854775807}
                                _ = fi88
                                _ = v89
                            }
                        }
                    case "":
                        switch 846 {
                        case 32767:
                            v75 = 1
                        default:
                            var v90 float64 = ((1.5 + -2.25) * 1.5)
                            _ = v90
                        }
                    case "foo bar":
                        var v91 int8 = -128
                        _ = v91
                    }
                default:
                    if true {
                        v92 := GenS10{F0: 1, F1: 0}
                        var v93 bool = (v75 == v83)
                        var v94 uint64 = 2
                        _ = v92
                        _ = v93
                        _ = v94
                    } else {
                        var v95 uint8 = (v74 & v74)
                        var v96 uint32 = (v83 % ((v75) | 1))
                        for fi97 := 0; fi97 < 10; fi97++ {
                            fmt.Println(v70)
                            _ = fi97
                        }
                        _ = v95
                        _ = v96
                    }
                }
                if true {
                    v98 := GenS6{F0: v83}
                    _ = v98
                }
                _ = v83
            }
        }
    }
    var v99 uint16 = 1
    if (18446744073709551615 > 4294967295) {
        v70 = 0.0
        var v100 int = 32768
        _ = v100
    }
    v70 += 3.14159
    if (false || (0 == -1)) {
        switch v1 {
        case "a":
            v74 = v74
        case "":
            switch 824 {
            case 510:
                if ((v99 * 0) <= 65535) {
                    if (248 < -2147483648) {
                        var v101 string = v1
                        v99 -= (v99 % ((v99) | 1))
                        _ = v101
                    } else {
                        var v102 bool = ((2147483647 >= v75) || (false && false))
                        for fi103 := 0; fi103 < 18; fi103++ {
                            v74 = (v74 ^ v74)
                            if v102 {
                                if len(v3) > 0 {
                                    v3[0] = v1
                                }
                                for fi104 := 0; fi104 < 17; fi104++ {
                                    if true {
                                        if (((0 != -621) || true) && (42 != -610)) {
                                            switch fi103 {
                                            case 127:
                                                var v105 int = 32768
                                                _ = v105
                                            default:
                                                fmt.Println(fi104, v70)
                                            }
                                            v75 = v75
                                        } else {
                                            v1 = v1
                                            v106 := GenS1{F0: v74}
                                            _ = v106
                                        }
                                        var v107 int8 = -1
                                        _ = v107
                                    }
                                    switch (^(fi104)) {
                                    case 542:
                                        var v108 int32 = 42
                                        _ = v108
                                    case 4294967295:
                                        if ((!v102) || (339 < 4294967295)) {
                                            v109 := GenS4{F0: 1, F1: 42}
                                            _ = v109
                                        }
                                    }
                                    v102 = (0 > 42)
                                    _ = fi104
                                }
                            }
                            _ = fi103
                        }
                        v110 := GenS5{F0: -2147483648, F1: 11}
                        _ = v102
                        _ = v110
                    }
                }
            case 42:
                _ = len(v3)
            default:
                v99 += v99
            }
        case "foo bar":
            var v111 int = 280
            _ = v111
        }
        var v112 int32 = -2147483648
        _ = v112
    } else {
        for fi113 := 0; fi113 < 5; fi113++ {
            for fi114 := 0; fi114 < 3; fi114++ {
                v115 := GenS1{F0: v74}
                _ = fi114
                _ = v115
            }
            _ = fi113
        }
        for fi116 := 0; fi116 < 12; fi116++ {
            v75 += v75
            if (9223372036854775807 != -431) {
                for fi117 := 0; fi117 < 3; fi117++ {
                    v70 = (v70 - -2.25)
                    var v118 int64 = 1
                    _ = fi117
                    _ = v118
                }
            }
            if (1.5 > 1.5) {
                v3 = append(v3, (v1 + (v1 + "hello")))
                if false {
                    v70 = v70
                    switch "foo bar" {
                    case "foo bar":
                        switch (fi116 & (-(fi116))) {
                        case -1:
                            v3 = append(v3, v1)
                        }
                    }
                }
            } else {
                v75 -= 4294967295
            }
            _ = fi116
        }
        v75 += uint32(v99)
    }
    v3 = append(v3, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v70:", v70)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v70
    _ = v73
    _ = v74
    _ = v75
    _ = v99
}
