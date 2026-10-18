package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 int32; F2 int32 }
type GenS7 struct { F0 int64; F1 int64; F2 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16; F1 bool; F2 int8 }
type GenS10 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int; F1 bool }
type GenS13 struct { F0 int8; F1 int32; F2 uint32 }
type GenS14 struct { F0 int; F1 int16; F2 uint8 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int32; F1 int64; F2 int32 }
func genFn0(p0 string) uint16 {
    p0 = ("foo bar" + p0)
    p0 = p0
    _ = p0
    return 32767
}
func genFn1(p0 int16, p1 uint32, p2 uint64) int32 {
    p1 = (p1 | p1)
    var v6 uint32 = p1
    var v7 uint8 = 127
    v8 := []int{655, 491}
    var v9 uint16 = genFn0(("\n\t" + "\n\t"))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 string = "foo bar"
    var v2 int16 = -32768
    var v3 uint32 = 2147483647
    v4 := []uint32{(v3 ^ 65535)}
    var v5 int64 = -9223372036854775808
    _ = len(v4)
    var v10 int64 = v5
    v11 := GenS1{F0: -128, F1: 255, F2: (v1 <= (v1 + v1))}
    var v12 uint8 = 4
    var v13 uint64 = 4294967295
    v14 := GenS1{F0: -128, F1: 255, F2: ((!(false || false)) || false)}
    if (v3 > (v3 & (v3 >> ((107) & 63)))) {
        for fi15 := 0; fi15 < 5; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 uint8 = v12
                v18 := GenS1{F0: -128, F1: 0, F2: true}
                var v19 uint = 129
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        v20 := GenS2{F0: 1}
        v21 := GenS2{F0: 42}
        _ = v20
        _ = v21
    }
    if (3 >= 32767) {
        for fi22 := 0; fi22 < 3; fi22++ {
            fmt.Println(v3, v1, v13)
            if false {
                v3 += v3
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (v12 < v12) {
                    var v23 bool = ((!true) || (!false))
                    _ = v23
                }
            }
            fmt.Println(v1)
            _ = fi22
        }
        var v24 string = "\n\t"
        if true {
            v1 = "hello"
            if (true && (!false)) {
                var v25 uint8 = (v12 | 0)
                if true {
                    var v26 bool = true
                    switch "\n\t" {
                    case "\n\t":
                        var v27 uint8 = (v12 >> ((1) & 63))
                        _ = v27
                    case "hello":
                        var v28 float64 = (0.1 / float64(0.1))
                        _ = v28
                    }
                    _ = v26
                }
                var v29 uint64 = v13
                _ = v25
                _ = v29
            }
        } else {
            v4 = append(v4, (^(v3)))
            if false {
                switch 32768 {
                case -2147483648:
                    v1 = ""
                }
                if false {
                    if len(v4) > 0 {
                        v4[0] = (^(v3))
                    }
                    v30 := GenS3{F0: 1, F1: 65535}
                    _ = v30
                } else {
                    if ((1 == 0) || true) {
                        var v31 int8 = 0
                        fmt.Println(v10, v31, v5)
                        _ = v31
                    } else {
                        var v32 int64 = -1
                        switch 2147483647 {
                        case 42:
                            var v33 int = -1
                            _ = v33
                        default:
                            if false {
                                if true {
                                    for fi34 := 0; fi34 < 13; fi34++ {
                                        var v35 string = ((v1 + (v24 + "\n\t")) + v24)
                                        _ = fi34
                                        _ = v35
                                    }
                                    for fi36 := 0; fi36 < 2; fi36++ {
                                        v12 = v12
                                        for fi37 := 0; fi37 < 16; fi37++ {
                                            v24 = "\n\t"
                                            _ = fi37
                                        }
                                        var v38 float64 = ((0.1 / (-2.25 / -2.25)) - 1e10)
                                        _ = fi36
                                        _ = v38
                                    }
                                    for fi39 := 0; fi39 < 17; fi39++ {
                                        v13 = (^(v13))
                                        _ = fi39
                                    }
                                }
                                switch 890 {
                                case 0:
                                    switch 42 {
                                    case 120:
                                        if (v24 < ((((("" + v1) + v1) + v1) + v1) + ("hello" + v1))) {
                                            if (32768 > 128) {
                                                if (342 <= 141) {
                                                    if false {
                                                        v40 := GenS3{F0: genFn0(v24), F1: 773}
                                                        switch 127 {
                                                        case 9223372036854775807:
                                                            if len(v4) > 0 {
                                                                v4[0] = (v3 % (((v3 + (v3 & v3))) | 1))
                                                            }
                                                        default:
                                                            switch ((v1 + v24) + (((v24 + ("foo bar" + v24)) + ("a" + (v24 + (v24 + "hello")))) + v1)) {
                                                            case "\n\t":
                                                                if true {
                                                                    for fi41 := 0; fi41 < 4; fi41++ {
                                                                        _ = len(v4)
                                                                        _ = fi41
                                                                    }
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                            default:
                                                                v4 = append(v4, 296)
                                                            }
                                                        }
                                                        _ = v40
                                                    } else {
                                                        fmt.Println(v3, v24, v32)
                                                        if (42 <= 811) {
                                                            v1 = "foo bar"
                                                        }
                                                    }
                                                    v32 = (v10 ^ (v32 * v10))
                                                } else {
                                                    for fi42 := 0; fi42 < 3; fi42++ {
                                                        if true {
                                                            v4 = append(v4, v3)
                                                        }
                                                        if (v24 <= "foo bar") {
                                                            if false {
                                                                v10 = (v10 >> ((1) & 63))
                                                            }
                                                            var v43 string = v1
                                                            _ = v43
                                                        }
                                                        _ = fi42
                                                    }
                                                }
                                                fmt.Println(v24)
                                                var v44 uint64 = 491
                                                _ = v44
                                            } else {
                                                for fi45 := 0; fi45 < 6; fi45++ {
                                                    var v46 uint8 = (v12 >> ((640) & 63))
                                                    switch fi45 {
                                                    case -597:
                                                        for fi47 := 0; fi47 < 1; fi47++ {
                                                            switch v1 {
                                                            case "a":
                                                                v24 = "a"
                                                            case "foo bar":
                                                                v48 := GenS2{F0: 255}
                                                                _ = v48
                                                            }
                                                            v49 := GenS3{F0: 255, F1: 1}
                                                            _ = fi47
                                                            _ = v49
                                                        }
                                                    }
                                                    _ = fi45
                                                    _ = v46
                                                }
                                                switch (v24 + "foo bar") {
                                                case "a":
                                                    v5 -= v10
                                                case "hello":
                                                    var v50 int16 = (v2 / ((-32768) | 1))
                                                    _ = v50
                                                }
                                                v51 := GenS3{F0: uint16(v3), F1: 65535}
                                                _ = v51
                                            }
                                        }
                                    case 128:
                                        if ((v2 << ((0) & 63)) != v2) {
                                            if (2 == v3) {
                                                if false {
                                                    if true {
                                                        v4 = append(v4, v3)
                                                        var v52 uint16 = 792
                                                        _ = v52
                                                    }
                                                    v10 -= (v32 / ((v5) | 1))
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = v3
                                                }
                                            }
                                        }
                                    case 184:
                                        var v53 int32 = 42
                                        _ = v53
                                    }
                                default:
                                    v54 := GenS3{F0: genFn0("a"), F1: 65535}
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 7; fi55++ {
                                    var v56 int8 = 127
                                    switch 128 {
                                    case 201:
                                        var v57 int16 = int16(v10)
                                        _ = v57
                                    case 42:
                                        if (!(true || true)) {
                                            for fi58 := 0; fi58 < 16; fi58++ {
                                                var v59 float32 = 1.5
                                                var v60 int = 128
                                                v13 = (v13 & v13)
                                                v4 = append(v4, (^(v3)))
                                                _ = fi58
                                                _ = v59
                                                _ = v60
                                            }
                                            v61 := GenS4{F0: 4294967295}
                                            if ("a" <= "hello") {
                                                var v62 float32 = float32(1.5)
                                                if false {
                                                    if (42 != 907) {
                                                        switch fi55 {
                                                        case -1:
                                                            var v63 float32 = float32(1e10)
                                                            _ = v63
                                                        case -2147483648:
                                                            switch 2147483647 {
                                                            case 42:
                                                                v56 += (v56 ^ v56)
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v56)
                                                }
                                                _ = v62
                                            }
                                            _ = v61
                                        } else {
                                            if ((!((v12 << ((202) & 63)) != (^(v12)))) && ((v2 / (((v2 - v2)) | 1)) > (v2 + (v2 | (-(v2)))))) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v64 bool = (!(!true))
                                                v65 := GenS5{F0: (!v64)}
                                                _ = v64
                                                _ = v65
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            v24 = ((v1 + ((("" + "hello") + "\n\t") + "a")) + ((v1 + (("" + v1) + v1)) + v24))
                                        }
                                    case 4294967295:
                                        if (!(!true)) {
                                            if ((^(v13)) == (v13 >> ((157) & 63))) {
                                                var v66 float32 = 3.14159
                                                switch fi55 {
                                                case 965:
                                                    for fi67 := 0; fi67 < 5; fi67++ {
                                                        var v68 string = ((("\n\t" + v24) + v1) + ((v1 + v1) + v24))
                                                        if ((v13 + v13) == v13) {
                                                            var v69 int = (fi67 ^ (fi55 | fi55))
                                                            _ = v69
                                                        }
                                                        var v70 string = (v24 + "\n\t")
                                                        _ = fi67
                                                        _ = v68
                                                        _ = v70
                                                    }
                                                case -2147483648:
                                                    v5 = v5
                                                default:
                                                    fmt.Println(v13)
                                                }
                                                _ = v66
                                            } else {
                                                v10 = (v10 | 1)
                                                if ((v10 + (v10 * 574)) > int64(fi55)) {
                                                    if (1.5 != (1e10 + 1.5)) {
                                                        switch 674 {
                                                        case 2147483647:
                                                            v71 := GenS6{F0: (v12 - v12), F1: 1, F2: 0}
                                                            _ = v71
                                                        case -2147483648:
                                                            for fi72 := 0; fi72 < 7; fi72++ {
                                                                v13 = v13
                                                                v1 = "\n\t"
                                                                var v73 float64 = (float64(1e10) * 1e10)
                                                                _ = fi72
                                                                _ = v73
                                                            }
                                                        default:
                                                            var v74 float64 = 1e10
                                                            _ = v74
                                                        }
                                                        fmt.Println(v5, v32)
                                                    } else {
                                                        switch (^(fi55)) {
                                                        case 663:
                                                            if false {
                                                                _ = len(v4)
                                                            }
                                                        default:
                                                            v12 = (^(v12))
                                                        }
                                                    }
                                                } else {
                                                    v10 += int64(v3)
                                                    if ((^(v3)) != (^(v3))) {
                                                        v75 := GenS4{F0: fi55}
                                                        v5 -= v5
                                                        var v76 uint8 = 2
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                }
                                            }
                                            if false {
                                                for fi77 := 0; fi77 < 14; fi77++ {
                                                    v4 = append(v4, v3)
                                                    fmt.Println(fi77)
                                                    v5 = (v10 / (((v10 | v10)) | 1))
                                                    v24 = v1
                                                    _ = len(v4)
                                                    _ = fi77
                                                }
                                            }
                                        }
                                    default:
                                        if ((0 != v13) && (((false && ((^(v3)) <= v3)) && (true && false)) || (false || true))) {
                                            _ = len(v4)
                                            switch (fi55 % (((-(fi55))) | 1)) {
                                            case -1:
                                                var v78 uint64 = (v13 * (v13 | v13))
                                                _ = v78
                                            case 32768:
                                                v79 := GenS5{F0: false}
                                                _ = v79
                                            }
                                        }
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                            }
                        }
                        _ = v32
                    }
                    switch v1 {
                    case "foo bar":
                        fmt.Println(v10)
                    case "a":
                        for fi80 := 0; fi80 < 10; fi80++ {
                            if false {
                                for fi81 := 0; fi81 < 12; fi81++ {
                                    switch v1 {
                                    case "hello":
                                        switch fi81 {
                                        case -815:
                                            v2 += (v2 >> ((255) & 63))
                                        case 731:
                                            if true {
                                                var v82 int32 = -2147483648
                                                v82 = v82
                                                _ = v82
                                            }
                                        case 2147483647:
                                            _ = len(v4)
                                            v13 = v13
                                            _ = len(v4)
                                        default:
                                            if (true || false) {
                                                if false {
                                                    fmt.Println(v10, v3, v10)
                                                    v10 = (v5 / ((v5) | 1))
                                                    if len(v4) > 0 {
                                                        v4[0] = (v3 - v3)
                                                    }
                                                    fmt.Println(fi80, v5)
                                                }
                                                _ = len(v4)
                                                var v83 uint16 = 65535
                                                _ = v83
                                            } else {
                                                v84 := GenS7{F0: (v10 & (v10 / (((v10 & v5)) | 1))), F1: (-(v5)), F2: (v13 | (v13 & (v13 << ((1) & 63))))}
                                                var v85 int32 = int32(v12)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v13 -= (v13 | v13)
                                                v4 = append(v4, (v3 & v3))
                                                _ = v84
                                                _ = v85
                                            }
                                        }
                                    case "":
                                        v3 -= v3
                                    default:
                                        v12 += (v12 << ((42) & 63))
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = fi81
                                }
                            } else {
                                v86 := GenS6{F0: 42, F1: 0, F2: 1}
                                var v87 int16 = v2
                                _ = v86
                                _ = v87
                            }
                            _ = fi80
                        }
                    case "\n\t":
                        switch 932 {
                        case -712:
                            if true {
                                switch v24 {
                                case "":
                                    if (false || false) {
                                        switch (v1 + v24) {
                                        case "":
                                            for fi88 := 0; fi88 < 7; fi88++ {
                                                switch (fi88 / ((int(v12)) | 1)) {
                                                case 20:
                                                    for fi89 := 0; fi89 < 19; fi89++ {
                                                        if false {
                                                            var v90 uint16 = 32767
                                                            var v91 uint = 49
                                                            fmt.Println(v5, fi88)
                                                            if true {
                                                                v92 := GenS8{F0: (v2 ^ (^(v2)))}
                                                                _ = v92
                                                            } else {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                var v93 uint8 = (v12 >> ((v91) & 63))
                                                                v3 += (v3 + v3)
                                                                _ = v93
                                                            }
                                                            v2 += 32767
                                                            var v94 uint64 = 18446744073709551615
                                                            v95 := GenS1{F0: 1, F1: v91, F2: ((!((113 > -2147483648) || (348 > 398))) || (v90 <= 65535))}
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v90
                                                            _ = v91
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                        _ = fi89
                                                    }
                                                }
                                                _ = fi88
                                            }
                                        case "foo bar":
                                            v96 := GenS9{F0: 588, F1: true, F2: 100}
                                            _ = v96
                                        default:
                                            _ = len(v4)
                                        }
                                    } else {
                                        if false {
                                            var v97 uint = 100
                                            _ = v97
                                        } else {
                                            var v98 string = v1
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v99 uint32 = 2147483647
                                                _ = v99
                                            case "a":
                                                var v100 int8 = -56
                                                _ = v100
                                            case "":
                                                if true {
                                                    if len(v4) > 0 {
                                                        v4[0] = v3
                                                    }
                                                    if false {
                                                        if (255 == 72) {
                                                            _ = len(v4)
                                                            if (true && true) {
                                                                var v101 uint = 0
                                                                _ = v101
                                                            }
                                                            var v102 bool = true
                                                            _ = v102
                                                        } else {
                                                            var v103 uint8 = 42
                                                            var v104 float32 = ((((0.1 / 1.0) + ((float32(3.14159) - 0.0) / 0.1)) - 0.1) / -2.25)
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                    } else {
                                                        var v105 uint64 = (v13 ^ v13)
                                                        if (255 == 1) {
                                                            if false {
                                                                v105 = (v105 + v13)
                                                                v105 = (v105 | 42)
                                                                fmt.Println(v2, v5, v24)
                                                            }
                                                            if (uint(v3) > 1) {
                                                                var v106 string = ""
                                                                var v107 bool = true
                                                                _ = v106
                                                                _ = v107
                                                            } else {
                                                                _ = len(v4)
                                                                if false {
                                                                    for fi108 := 0; fi108 < 8; fi108++ {
                                                                        v24 = "hello"
                                                                        if (true || (((("\n\t" + ("a" + v24)) + "a") + (v1 + "")) != "\n\t")) {
                                                                            var v109 int32 = 1
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 65535
                                                                            }
                                                                            _ = v109
                                                                        }
                                                                        _ = fi108
                                                                    }
                                                                    v24 = ((("foo bar" + v1) + v24) + v98)
                                                                }
                                                                if false {
                                                                    var v110 int8 = 100
                                                                    v12 = v12
                                                                    v10 -= (v10 ^ v5)
                                                                    _ = v110
                                                                } else {
                                                                    if (false || false) {
                                                                        var v111 int32 = 751
                                                                        switch v98 {
                                                                        case "a":
                                                                            v105 -= (v105 >> ((2) & 63))
                                                                        case "foo bar":
                                                                            for fi112 := 0; fi112 < 18; fi112++ {
                                                                                v3 = v3
                                                                                if (true && true) {
                                                                                    switch fi112 {
                                                                                    case -227:
                                                                                        fmt.Println(v10)
                                                                                    default:
                                                                                        var v113 string = (v1 + ("\n\t" + v98))
                                                                                        _ = v113
                                                                                    }
                                                                                }
                                                                                v114 := GenS6{F0: uint8(v3), F1: (v111 ^ v111), F2: v111}
                                                                                _ = fi112
                                                                                _ = v114
                                                                            }
                                                                        case "hello":
                                                                            for fi115 := 0; fi115 < 4; fi115++ {
                                                                                switch (v24 + (v24 + ("hello" + v24))) {
                                                                                case "a":
                                                                                    for fi116 := 0; fi116 < 6; fi116++ {
                                                                                        v10 = 14
                                                                                        v117 := GenS8{F0: -1000}
                                                                                        _ = fi116
                                                                                        _ = v117
                                                                                    }
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        var v118 int = 128
                                                                                        fmt.Println(v118, v2)
                                                                                        var v119 uint64 = (v13 >> ((65535) & 63))
                                                                                        _ = v118
                                                                                        _ = v119
                                                                                    } else {
                                                                                        v13 = v13
                                                                                        fmt.Println(v111, v98, v13)
                                                                                        if true {
                                                                                            for fi120 := 0; fi120 < 13; fi120++ {
                                                                                                var v121 int8 = 96
                                                                                                var v122 uint8 = (v12 >> ((18446744073709551615) & 63))
                                                                                                v121 += -101
                                                                                                _ = fi120
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                            }
                                                                                            fmt.Println(v3)
                                                                                        }
                                                                                    }
                                                                                case "\n\t":
                                                                                    v24 = v1
                                                                                default:
                                                                                    v5 -= 42
                                                                                }
                                                                                for fi123 := 0; fi123 < 5; fi123++ {
                                                                                    switch v98 {
                                                                                    case "a":
                                                                                        var v124 int32 = v111
                                                                                        _ = v124
                                                                                    case "foo bar":
                                                                                        v125 := GenS10{F0: (v105 & 18446744073709551615), F1: genFn1(int16(fi123), v3, uint64(v3)), F2: uint64(v12)}
                                                                                        _ = v125
                                                                                    }
                                                                                    if (false || true) {
                                                                                        var v126 int64 = (v10 << ((0) & 63))
                                                                                        var v127 int8 = 0
                                                                                        switch (v1 + ("foo bar" + v98)) {
                                                                                        case "foo bar":
                                                                                            v128 := GenS11{F0: 42}
                                                                                            _ = v128
                                                                                        default:
                                                                                            v111 = 2147483647
                                                                                        }
                                                                                        _ = v126
                                                                                        _ = v127
                                                                                    } else {
                                                                                        switch ("" + ((v24 + "\n\t") + "foo bar")) {
                                                                                        case "\n\t":
                                                                                            v129 := GenS4{F0: 2147483647}
                                                                                            _ = v129
                                                                                        case "":
                                                                                            switch fi123 {
                                                                                            case -2147483648:
                                                                                                var v130 uint = 4294967295
                                                                                                _ = v130
                                                                                            default:
                                                                                                v1 = v1
                                                                                            }
                                                                                        case "hello":
                                                                                            v5 = v10
                                                                                        default:
                                                                                            if false {
                                                                                                if (false && (true || true)) {
                                                                                                    if true {
                                                                                                        fmt.Println(v111, v13)
                                                                                                        var v131 int = fi115
                                                                                                        switch v131 {
                                                                                                        case 250:
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                            v10 = -542
                                                                                                            switch v131 {
                                                                                                            case 32768:
                                                                                                                v13 = v13
                                                                                                            case -666:
                                                                                                                var v132 uint8 = (v12 * v12)
                                                                                                                _ = v132
                                                                                                            case 42:
                                                                                                                v4 = append(v4, v3)
                                                                                                            default:
                                                                                                                if ((0 >= v3) || ((v3 ^ v3) == (v3 - v3))) {
                                                                                                                    v10 = 978
                                                                                                                    switch fi115 {
                                                                                                                    case -567:
                                                                                                                        v2 = v2
                                                                                                                    case -1:
                                                                                                                        fmt.Println(v3, v111)
                                                                                                                    default:
                                                                                                                        v111 -= 1
                                                                                                                    }
                                                                                                                    v3 += 65535
                                                                                                                    v133 := GenS12{F0: v131, F1: (false && false)}
                                                                                                                    _ = v133
                                                                                                                } else {
                                                                                                                    v3 = (v3 << ((42) & 63))
                                                                                                                    switch v24 {
                                                                                                                    case "foo bar":
                                                                                                                        var v134 string = ("hello" + v1)
                                                                                                                        if len(v4) > 0 {
                                                                                                                            v4[0] = v3
                                                                                                                        }
                                                                                                                        v5 = v10
                                                                                                                        _ = v134
                                                                                                                    }
                                                                                                                    v135 := GenS10{F0: 18446744073709551615, F1: (v111 + 68), F2: (^(v105))}
                                                                                                                    _ = v135
                                                                                                                }
                                                                                                                v4 = append(v4, 4294967295)
                                                                                                                _ = len(v4)
                                                                                                            }
                                                                                                        case -390:
                                                                                                            v136 := GenS13{F0: 127, F1: v111, F2: (v3 - 2147483647)}
                                                                                                            _ = v136
                                                                                                        case 42:
                                                                                                            v137 := GenS14{F0: (-(fi115)), F1: int16(v13), F2: v12}
                                                                                                            _ = v137
                                                                                                        }
                                                                                                        _ = v131
                                                                                                    }
                                                                                                    v1 = v98
                                                                                                    var v138 int32 = (v111 | (v111 + 2147483647))
                                                                                                    _ = v138
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v139 := GenS15{F0: 522}
                                                                                        switch v98 {
                                                                                        case "\n\t":
                                                                                            v140 := GenS11{F0: 718}
                                                                                            _ = v140
                                                                                        }
                                                                                        _ = v139
                                                                                    }
                                                                                    _ = fi123
                                                                                }
                                                                                if (v111 <= (-(v111))) {
                                                                                    v1 = v98
                                                                                }
                                                                                _ = fi115
                                                                            }
                                                                        default:
                                                                            for fi141 := 0; fi141 < 12; fi141++ {
                                                                                var v142 int32 = v111
                                                                                _ = fi141
                                                                                _ = v142
                                                                            }
                                                                        }
                                                                        v143 := GenS16{F0: (v2 + v2)}
                                                                        _ = v111
                                                                        _ = v143
                                                                    }
                                                                    var v144 int64 = v10
                                                                    v12 -= v12
                                                                    _ = v144
                                                                }
                                                            }
                                                            if false {
                                                                _ = len(v4)
                                                                v145 := GenS13{F0: 127, F1: -2147483648, F2: 721}
                                                                if false {
                                                                    fmt.Println(v1, v2)
                                                                    var v146 int16 = -31
                                                                    var v147 string = "a"
                                                                    _ = v146
                                                                    _ = v147
                                                                } else {
                                                                    var v148 int16 = (v2 - v2)
                                                                    _ = v148
                                                                }
                                                                _ = v145
                                                            }
                                                        }
                                                        _ = v105
                                                    }
                                                } else {
                                                    for fi149 := 0; fi149 < 10; fi149++ {
                                                        v4 = append(v4, 0)
                                                        _ = fi149
                                                    }
                                                    for fi150 := 0; fi150 < 1; fi150++ {
                                                        v3 = 2147483647
                                                        _ = fi150
                                                    }
                                                }
                                            default:
                                                var v151 uint16 = 32767
                                                _ = v151
                                            }
                                            _ = v98
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v24)
                                        case "\n\t":
                                            if len(v4) > 0 {
                                                v4[0] = v3
                                            }
                                        }
                                    }
                                case "foo bar":
                                    switch ((v1 + (((("\n\t" + ("" + "")) + "\n\t") + v24) + "hello")) + (("\n\t" + "\n\t") + v24)) {
                                    case "foo bar":
                                        if false {
                                            if false {
                                                var v152 int8 = 26
                                                if true {
                                                    v152 += v152
                                                    var v153 uint = 880
                                                    fmt.Println(v12)
                                                    _ = v153
                                                } else {
                                                    v4 = append(v4, v3)
                                                    v24 = v24
                                                    fmt.Println(v152)
                                                }
                                                for fi154 := 0; fi154 < 9; fi154++ {
                                                    for fi155 := 0; fi155 < 19; fi155++ {
                                                        v156 := GenS8{F0: (v2 * 0)}
                                                        v24 = ("a" + "\n\t")
                                                        v4 = append(v4, v3)
                                                        _ = fi155
                                                        _ = v156
                                                    }
                                                    v1 = v1
                                                    fmt.Println(v12, v152)
                                                    _ = fi154
                                                }
                                                _ = v152
                                            }
                                            v3 = v3
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    case "hello":
                                        var v157 int8 = -11
                                        _ = v157
                                    default:
                                        v158 := GenS1{F0: int8(v2), F1: 18446744073709551615, F2: false}
                                        _ = v158
                                    }
                                case "\n\t":
                                    for fi159 := 0; fi159 < 9; fi159++ {
                                        var v160 int = fi159
                                        var v161 int = fi159
                                        switch (v161 | -157) {
                                        case 792:
                                            v162 := GenS17{F0: 349}
                                            _ = v162
                                        case -641:
                                            v4 = append(v4, (v3 | 1))
                                        case 128:
                                            if (65535 < 194) {
                                                if false {
                                                    fmt.Println(v1, v24, v13)
                                                    v13 += v13
                                                }
                                                switch "a" {
                                                case "":
                                                    v4 = append(v4, v3)
                                                case "hello":
                                                    v24 = v1
                                                }
                                            }
                                        }
                                        _ = fi159
                                        _ = v160
                                        _ = v161
                                    }
                                default:
                                    var v163 uint16 = 32767
                                    _ = v163
                                }
                                v164 := GenS4{F0: -9223372036854775808}
                                _ = v164
                            }
                        case 101:
                            switch -9223372036854775808 {
                            case -2147483648:
                                var v165 uint16 = genFn0(("foo bar" + "hello"))
                                _ = v165
                            }
                        case -366:
                            v166 := GenS18{F0: 0, F1: v10, F2: 0}
                            _ = v166
                        }
                    }
                    v10 -= (v10 + v10)
                }
                var v167 int32 = -1
                _ = v167
            }
            for fi168 := 0; fi168 < 2; fi168++ {
                v4 = append(v4, v3)
                _ = fi168
            }
        }
        _ = v24
    }
    var v169 int64 = v10
    v13 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v169:", v169)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v169
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
