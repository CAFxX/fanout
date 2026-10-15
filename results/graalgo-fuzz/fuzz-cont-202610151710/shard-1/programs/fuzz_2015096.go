package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int64; F1 uint32; F2 int16 }
type GenS10 struct { F0 int32; F1 string; F2 uint }
type GenS11 struct { F0 uint; F1 int32; F2 int }
type GenS12 struct { F0 uint32; F1 int16; F2 bool }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint32; F1 int32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int; F1 uint16; F2 int32 }
type GenS17 struct { F0 int64; F1 uint8 }
type GenS18 struct { F0 uint8; F1 int; F2 int }
type GenS19 struct { F0 uint8; F1 string }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 uint32; F1 int; F2 int16 }
func genFn0(p0 uint32, p1 uint, p2 string) uint {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int16 = -644
        p0 = p0
        v7 = -1
        _ = fi6
        _ = v7
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 uint16 = 65535
    var v2 uint32 = 42
    var v3 float32 = (-2.25 + -2.25)
    var v4 int64 = 370
    v5 := []string{"foo bar"}
    var v8 string = ("\n\t" + "hello")
    switch v8 {
    case "\n\t":
        if (4294967295 != 0) {
            var v9 int16 = -1
            if len(v5) > 0 {
                v5[0] = v8
            }
            v9 = -32768
            _ = v9
        } else {
            var v10 uint8 = 127
            _ = v10
        }
    case "a":
        for fi11 := 0; fi11 < 13; fi11++ {
            v8 = "a"
            for fi12 := 0; fi12 < 13; fi12++ {
                switch (fi11 * (fi11 | 2147483647)) {
                case -950:
                    if len(v5) > 0 {
                        v5[0] = ""
                    }
                default:
                    var v13 uint = genFn0((v2 + (v2 * (v2 & (v2 >> ((721) & 63))))), 2, v8)
                    _ = v13
                }
                v8 = v8
                switch fi11 {
                case 2147483647:
                    v14 := GenS1{F0: (v2 + 4294967295), F1: fi11, F2: (^(v1))}
                    _ = v14
                case 0:
                    if false {
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        v15 := GenS1{F0: v2, F1: fi12, F2: (v1 / (((v1 << ((228) & 63))) | 1))}
                        var v16 uint8 = 1
                        _ = v15
                        _ = v16
                    }
                    if (42 > 42) {
                        v17 := GenS2{F0: true, F1: false, F2: -413}
                        switch (-(fi12)) {
                        case 0:
                            var v18 uint64 = 1
                            _ = v18
                        default:
                            v3 -= v3
                        }
                        _ = v17
                    }
                    var v19 int32 = -1
                    v20 := GenS3{F0: 1}
                    v21 := GenS1{F0: (v2 / (((v2 * v2)) | 1)), F1: (fi11 | int(v2)), F2: v1}
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = fi12
            }
            var v22 int32 = 2147483647
            _ = fi11
            _ = v22
        }
    default:
        var v23 int = 0
        _ = v23
    }
    var v24 int16 = 679
    var v25 float32 = (v3 - v3)
    for fi26 := 0; fi26 < 5; fi26++ {
        if false {
            _ = len(v5)
            for fi27 := 0; fi27 < 5; fi27++ {
                _ = len(v5)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v28 int64 = 96
                _ = fi27
                _ = v28
            }
        }
        _ = fi26
    }
    fmt.Println(v8, v24)
    var v29 string = "a"
    for fi30 := 0; fi30 < 8; fi30++ {
        fmt.Println(v25, v29, v29)
        var v31 uint8 = 245
        _ = fi30
        _ = v31
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v32 int = 32767
    if len(v5) > 0 {
        v5[0] = (v8 + v29)
    }
    v5 = append(v5, ((v29 + v8) + ("" + ("a" + ("hello" + "\n\t")))))
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 uint8 = 42
        if false {
            v32 = (v32 ^ (fi33 - fi33))
            switch v8 {
            case "\n\t":
                fmt.Println(v2)
            case "foo bar":
                var v35 uint64 = 42
                _ = v35
            case "hello":
                if false {
                    if true {
                        switch v32 {
                        case -695:
                            var v36 float64 = 1.5
                            _ = v36
                        default:
                            switch fi33 {
                            case 1:
                                if ((2 < 486) && true) {
                                    var v37 int16 = -1
                                    _ = v37
                                }
                            case -1:
                                var v38 int16 = (v24 & v24)
                                _ = v38
                            case 2147483647:
                                switch v8 {
                                case "\n\t":
                                    switch v32 {
                                    case 2147483647:
                                        var v39 int = v32
                                        _ = v39
                                    case 128:
                                        var v40 string = (v8 + ("hello" + "\n\t"))
                                        _ = v40
                                    case 127:
                                        v34 = (v34 ^ uint8(v1))
                                    }
                                default:
                                    switch v29 {
                                    case "\n\t":
                                        switch ((v8 + v8) + (((("a" + ((v8 + "foo bar") + (v29 + v8))) + v29) + v8) + "a")) {
                                        case "foo bar":
                                            v8 = v8
                                        default:
                                            switch "hello" {
                                            case "a":
                                                v41 := GenS3{F0: 18446744073709551615}
                                                _ = v41
                                            }
                                        }
                                    case "a":
                                        if true {
                                            fmt.Println(v32, v32, v25)
                                            for fi42 := 0; fi42 < 8; fi42++ {
                                                v1 = (v1 >> ((255) & 63))
                                                var v43 float32 = (0.1 - ((v25 * ((v3 / 1e10) + v25)) - v3))
                                                if len(v5) > 0 {
                                                    v5[0] = v8
                                                }
                                                _ = fi42
                                                _ = v43
                                            }
                                        }
                                    case "hello":
                                        if false {
                                            var v44 uint8 = v34
                                            switch 9223372036854775807 {
                                            case 0:
                                                fmt.Println(v8, v8)
                                            case -2147483648:
                                                if (608 != 714) {
                                                    var v45 int32 = -1
                                                    switch v29 {
                                                    case "hello":
                                                        var v46 uint64 = uint64(v2)
                                                        _ = v46
                                                    case "a":
                                                        fmt.Println(v3)
                                                    case "foo bar":
                                                        for fi47 := 0; fi47 < 15; fi47++ {
                                                            v44 -= v34
                                                            _ = fi47
                                                        }
                                                    }
                                                    _ = v45
                                                } else {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    if (-2.25 > (3.14159 / 3.14159)) {
                                                        fmt.Println(v8, v24)
                                                    }
                                                    var v48 int8 = int8(v2)
                                                    _ = v48
                                                }
                                            case 1:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            default:
                                                switch 32767 {
                                                case -9223372036854775808:
                                                    if ((false && (false && (!true))) || ((v24 & v24) != (v24 >> ((18446744073709551615) & 63)))) {
                                                        v5 = append(v5, ((v8 + ((v29 + "hello") + v8)) + (v8 + "foo bar")))
                                                        var v49 uint32 = v2
                                                        _ = len(v5)
                                                        _ = v49
                                                    } else {
                                                        if true {
                                                            var v50 int32 = -2147483648
                                                            _ = v50
                                                        } else {
                                                            switch v29 {
                                                            case "hello":
                                                                switch 9223372036854775807 {
                                                                case 4294967295:
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                default:
                                                                    fmt.Println(v32, v44)
                                                                }
                                                            default:
                                                                switch (fi33 * fi33) {
                                                                case 4294967295:
                                                                    fmt.Println(v44, v34, v4)
                                                                case 0:
                                                                    v5 = append(v5, v29)
                                                                default:
                                                                    var v51 int16 = 1000
                                                                    _ = v51
                                                                }
                                                            }
                                                        }
                                                        var v52 uint = 65535
                                                        v24 -= v24
                                                        _ = v52
                                                    }
                                                case 0:
                                                    if ((v4 & v4) != 0) {
                                                        v53 := GenS4{F0: (v2 << ((1) & 63))}
                                                        if len(v5) > 0 {
                                                            v5[0] = ("\n\t" + v29)
                                                        }
                                                        var v54 uint64 = 334
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                case -97:
                                                    for fi55 := 0; fi55 < 7; fi55++ {
                                                        var v56 uint64 = 18446744073709551615
                                                        var v57 int64 = (v4 % ((v4) | 1))
                                                        var v58 uint = 42
                                                        _ = fi55
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                }
                                            }
                                            _ = v44
                                        }
                                    default:
                                        for fi59 := 0; fi59 < 10; fi59++ {
                                            if false {
                                                if false {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                } else {
                                                    var v60 int16 = v24
                                                    switch ((v29 + (v8 + ("hello" + (v8 + v29)))) + "\n\t") {
                                                    case "\n\t":
                                                        switch (v32 - v32) {
                                                        case 4294967295:
                                                            v4 = v4
                                                            var v61 int8 = 127
                                                            _ = v61
                                                        case 42:
                                                            v34 -= (^(v34))
                                                        }
                                                    case "":
                                                        if false {
                                                            v32 = fi59
                                                            var v62 float64 = -2.25
                                                            _ = v62
                                                        }
                                                    default:
                                                        var v63 uint64 = 18446744073709551615
                                                        _ = v63
                                                    }
                                                    var v64 uint8 = 34
                                                    _ = v60
                                                    _ = v64
                                                }
                                                if (!false) {
                                                    v65 := GenS4{F0: v2}
                                                    v3 = 0.1
                                                    if ((9223372036854775807 == (^(v4))) || (v25 >= (0.1 + (v3 * (float32((((((((1e10 / 1.0) + (0.0 + 1.5)) / -2.25) / 0.1) + (float64(float32(1e10)) / 1.5)) / 1.0) - (0.1 - 0.0))) / v25))))) {
                                                        fmt.Println(v2, v4, v34)
                                                        var v66 string = v29
                                                        _ = v66
                                                    } else {
                                                        v67 := GenS5{F0: (255 <= 65535), F1: 915}
                                                        v3 -= v3
                                                        switch v8 {
                                                        case "":
                                                            v5 = append(v5, ("hello" + "foo bar"))
                                                        }
                                                        _ = v67
                                                    }
                                                    if (false || true) {
                                                        fmt.Println(v29, v2)
                                                    }
                                                    fmt.Println(v25, v24, fi59)
                                                    _ = v65
                                                } else {
                                                    if (v2 <= v2) {
                                                        v5 = append(v5, (v8 + "foo bar"))
                                                        switch v32 {
                                                        case -509:
                                                            switch (v32 + 32767) {
                                                            case 2147483647:
                                                                v68 := GenS6{F0: v2}
                                                                v2 = (v2 + (v2 % ((v2) | 1)))
                                                                _ = v68
                                                            default:
                                                                var v69 uint8 = 128
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v29 = ("foo bar" + v8)
                                                                    _ = len(v5)
                                                                    var v70 uint8 = v34
                                                                    var v71 float64 = 3.14159
                                                                    _ = v70
                                                                    _ = v71
                                                                case "\n\t":
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                case "":
                                                                    v72 := GenS7{F0: true, F1: ((v1 ^ v1) == (^(v1)))}
                                                                    _ = v72
                                                                default:
                                                                    v73 := GenS8{F0: (true && true), F1: int8(v1)}
                                                                    _ = v73
                                                                }
                                                                _ = len(v5)
                                                                _ = v69
                                                            }
                                                        default:
                                                            var v74 int32 = -1
                                                            if true {
                                                                if (!(true && true)) {
                                                                    var v75 int16 = (^(v24))
                                                                    v5 = append(v5, (v8 + v8))
                                                                    v76 := GenS9{F0: 42, F1: v2, F2: int16(fi33)}
                                                                    v29 = ("\n\t" + "foo bar")
                                                                    var v77 uint = 606
                                                                    _ = v75
                                                                    _ = v76
                                                                    _ = v77
                                                                }
                                                                v25 += (0.0 - 1e10)
                                                                if ((!true) || false) {
                                                                    if len(v5) > 0 {
                                                                        v5[0] = v29
                                                                    }
                                                                } else {
                                                                    switch (v8 + "foo bar") {
                                                                    case "foo bar":
                                                                        v78 := GenS8{F0: (uint32(v1) <= (v2 * (v2 + 591))), F1: 1}
                                                                        _ = v78
                                                                    }
                                                                    if false {
                                                                        if true {
                                                                            v79 := GenS5{F0: false, F1: 0}
                                                                            if (!(((v1 | (v1 % ((v1) | 1))) >= v1) && true)) {
                                                                                if false {
                                                                                    switch (v29 + ("a" + v29)) {
                                                                                    case "hello":
                                                                                        var v80 uint32 = 2147483648
                                                                                        _ = v80
                                                                                    case "foo bar":
                                                                                        switch v8 {
                                                                                        case "hello":
                                                                                            if true {
                                                                                                var v81 float32 = v25
                                                                                                var v82 uint16 = v1
                                                                                                var v83 bool = false
                                                                                                _ = v81
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                            }
                                                                                        default:
                                                                                            switch (-(v32)) {
                                                                                            case 4294967295:
                                                                                                var v84 float32 = ((v3 / v3) / 1e10)
                                                                                                _ = v84
                                                                                            case -1:
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = v8
                                                                                                }
                                                                                            default:
                                                                                                var v85 int32 = (v74 << ((42) & 63))
                                                                                                _ = v85
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        _ = len(v5)
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    if ((v2 - (v2 * 438)) > 2147483648) {
                                                                                        v8 = ((v8 + ("\n\t" + v29)) + v8)
                                                                                    }
                                                                                    v86 := GenS1{F0: v2, F1: (v32 >> ((1) & 63)), F2: uint16(v34)}
                                                                                    _ = v86
                                                                                } else {
                                                                                    v87 := GenS6{F0: 42}
                                                                                    _ = v87
                                                                                }
                                                                                v3 = v3
                                                                                fmt.Println(v3)
                                                                                fmt.Println(v29)
                                                                                var v88 uint32 = 315
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                _ = v88
                                                                            }
                                                                            if ((v29 + ("foo bar" + "hello")) < v29) {
                                                                                v1 = (v1 * v1)
                                                                                var v89 uint64 = uint64(v1)
                                                                                _ = v89
                                                                            }
                                                                            _ = v79
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "":
                                                                                if ((v24 * v24) != -893) {
                                                                                    v32 += 127
                                                                                    switch v8 {
                                                                                    case "hello":
                                                                                        v90 := GenS10{F0: 206, F1: v29, F2: 462}
                                                                                        _ = v90
                                                                                    case "\n\t":
                                                                                        v34 = uint8(v1)
                                                                                    case "foo bar":
                                                                                        fmt.Println(fi33, v2)
                                                                                    }
                                                                                    var v91 int64 = (^(v4))
                                                                                    _ = v91
                                                                                } else {
                                                                                    if false {
                                                                                        if true {
                                                                                            v92 := GenS8{F0: (!false), F1: -128}
                                                                                            var v93 float32 = 1.5
                                                                                            _ = v92
                                                                                            _ = v93
                                                                                        }
                                                                                        switch fi33 {
                                                                                        case -238:
                                                                                            var v94 uint = 679
                                                                                            _ = v94
                                                                                        case -163:
                                                                                            fmt.Println(v24, fi33)
                                                                                        case 104:
                                                                                            if (genFn0((v2 + 2), 715, v29) != 42) {
                                                                                                var v95 int = fi59
                                                                                                var v96 float64 = 0.0
                                                                                                var v97 int32 = -7
                                                                                                switch int(v1) {
                                                                                                case 0:
                                                                                                    var v98 int = fi59
                                                                                                    _ = v98
                                                                                                case -1:
                                                                                                    var v99 uint32 = (v2 & (v2 / ((v2) | 1)))
                                                                                                    _ = v99
                                                                                                case 712:
                                                                                                    if (v8 == "foo bar") {
                                                                                                        if ((v29 >= v29) || (2 >= genFn0((v2 * v2), 0, ""))) {
                                                                                                            v95 += 0
                                                                                                        } else {
                                                                                                            v100 := GenS11{F0: 65535, F1: (v97 << ((493) & 63)), F2: 1}
                                                                                                            var v101 int8 = 1
                                                                                                            _ = v100
                                                                                                            _ = v101
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v25)
                                                                                                        v97 = 907
                                                                                                    }
                                                                                                    var v102 int64 = v4
                                                                                                    var v103 int32 = (v97 + (v97 + 42))
                                                                                                    switch (v95 >> ((255) & 63)) {
                                                                                                    case 107:
                                                                                                        v5 = append(v5, v8)
                                                                                                    default:
                                                                                                        v74 += (v74 | v103)
                                                                                                    }
                                                                                                    _ = v102
                                                                                                    _ = v103
                                                                                                default:
                                                                                                    if false {
                                                                                                        v104 := GenS3{F0: 1}
                                                                                                        v105 := GenS12{F0: (v2 >> ((204) & 63)), F1: (v24 % (((v24 + v24)) | 1)), F2: true}
                                                                                                        switch -404 {
                                                                                                        case 249:
                                                                                                            if (((!false) && (18446744073709551615 >= 0)) || ((v74 % (((v74 + v97)) | 1)) != v97)) {
                                                                                                                _ = len(v5)
                                                                                                            } else {
                                                                                                                fmt.Println(v97, fi33, fi59)
                                                                                                            }
                                                                                                        case 643:
                                                                                                            if false {
                                                                                                                v106 := GenS2{F0: false, F1: true, F2: v74}
                                                                                                                if ((1.5 >= (((v96 * v96) * -2.25) + (1e10 * v96))) || false) {
                                                                                                                    v107 := GenS5{F0: false, F1: 214}
                                                                                                                    if (-76 < int8(v1)) {
                                                                                                                        v5 = append(v5, "hello")
                                                                                                                    } else {
                                                                                                                        var v108 string = v8
                                                                                                                        _ = v108
                                                                                                                    }
                                                                                                                    fmt.Println(v1)
                                                                                                                    _ = v107
                                                                                                                }
                                                                                                                _ = v106
                                                                                                            }
                                                                                                            v109 := GenS13{F0: (v95 | fi33)}
                                                                                                            fmt.Println(v8)
                                                                                                            _ = v109
                                                                                                        }
                                                                                                        switch v32 {
                                                                                                        case 1:
                                                                                                            var v110 string = v8
                                                                                                            if ((v24 & v24) > v24) {
                                                                                                                var v111 int = -132
                                                                                                                _ = v111
                                                                                                            } else {
                                                                                                                v25 -= (float32(v96) / ((v3 + 0.1) - v25))
                                                                                                                var v112 uint32 = (v2 * 83)
                                                                                                                _ = v112
                                                                                                            }
                                                                                                            _ = v110
                                                                                                        case 0:
                                                                                                            v74 = int32(v1)
                                                                                                        }
                                                                                                        _ = v104
                                                                                                        _ = v105
                                                                                                    } else {
                                                                                                        if len(v5) >= 2 {
                                                                                                            v5 = v5[1:len(v5)]
                                                                                                        }
                                                                                                        var v113 int32 = v74
                                                                                                        _ = v113
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v29, v34, fi33)
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                            } else {
                                                                                                v5 = append(v5, v29)
                                                                                                if (158 >= 756) {
                                                                                                    if true {
                                                                                                        var v114 string = v29
                                                                                                        _ = v114
                                                                                                    } else {
                                                                                                        var v115 int8 = -1
                                                                                                        v25 = (0.0 / 1e10)
                                                                                                        switch 127 {
                                                                                                        case -9223372036854775808:
                                                                                                            var v116 uint = 18446744073709551615
                                                                                                            _ = v116
                                                                                                        case 32767:
                                                                                                            switch -2147483648 {
                                                                                                            case 0:
                                                                                                                v5 = append(v5, v29)
                                                                                                            case 1:
                                                                                                                var v117 uint = 65535
                                                                                                                var v118 int32 = 42
                                                                                                                var v119 int = -979
                                                                                                                _ = v117
                                                                                                                _ = v118
                                                                                                                _ = v119
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v115
                                                                                                    }
                                                                                                    v25 = v3
                                                                                                }
                                                                                                var v120 int16 = v24
                                                                                                _ = v120
                                                                                            }
                                                                                        default:
                                                                                            v5 = append(v5, "")
                                                                                        }
                                                                                        v121 := GenS8{F0: (182 >= (v34 ^ (v34 + (v34 << ((0) & 63))))), F1: 97}
                                                                                        _ = v121
                                                                                    }
                                                                                    v122 := GenS1{F0: v2, F1: (-(fi33)), F2: uint16(v2)}
                                                                                    if (!false) {
                                                                                        v123 := GenS7{F0: false, F1: true}
                                                                                        _ = v123
                                                                                    }
                                                                                    _ = v122
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    switch v29 {
                                                                                    case "a":
                                                                                        if (v8 < ((v8 + (((("a" + v29) + (v8 + "hello")) + v8) + "hello")) + (v29 + "\n\t"))) {
                                                                                            fmt.Println(fi59)
                                                                                            switch int(v34) {
                                                                                            case -2147483648:
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = v8
                                                                                                }
                                                                                            default:
                                                                                                if (v29 < "foo bar") {
                                                                                                    var v124 int64 = 42
                                                                                                    if false {
                                                                                                        switch "a" {
                                                                                                        case "\n\t":
                                                                                                            var v125 int = fi33
                                                                                                            var v126 int16 = (v24 / ((v24) | 1))
                                                                                                            var v127 int8 = -1
                                                                                                            if true {
                                                                                                                v127 = 127
                                                                                                                if (uint64(v2) < 0) {
                                                                                                                    var v128 float64 = (1e10 + 3.14159)
                                                                                                                    v34 = v34
                                                                                                                    _ = v128
                                                                                                                }
                                                                                                                var v129 float64 = (float64((float32(((-2.25 - 3.14159) * 0.0)) + v3)) - 0.1)
                                                                                                                _ = v129
                                                                                                            }
                                                                                                            var v130 uint32 = (v2 - (v2 >> ((2) & 63)))
                                                                                                            _ = v125
                                                                                                            _ = v126
                                                                                                            _ = v127
                                                                                                            _ = v130
                                                                                                        case "a":
                                                                                                            v25 -= v3
                                                                                                        case "foo bar":
                                                                                                            if false {
                                                                                                                fmt.Println(v25, v34)
                                                                                                            } else {
                                                                                                                var v131 uint32 = (^(v2))
                                                                                                                v132 := GenS2{F0: false, F1: ((v3 < v3) && (v3 <= float32(0.1))), F2: 445}
                                                                                                                var v133 int = v32
                                                                                                                switch int(v131) {
                                                                                                                case 42:
                                                                                                                    v131 -= (^(v131))
                                                                                                                case -9223372036854775808:
                                                                                                                    var v134 float64 = float64((v25 / (v25 / (v25 / float32(-2.25)))))
                                                                                                                    _ = v134
                                                                                                                case 796:
                                                                                                                    if len(v5) >= 2 {
                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                    }
                                                                                                                }
                                                                                                                if false {
                                                                                                                    v34 = (v34 + v34)
                                                                                                                }
                                                                                                                v74 = v74
                                                                                                                _ = len(v5)
                                                                                                                var v135 uint8 = v34
                                                                                                                _ = v131
                                                                                                                _ = v132
                                                                                                                _ = v133
                                                                                                                _ = v135
                                                                                                            }
                                                                                                        }
                                                                                                        switch fi33 {
                                                                                                        case -1:
                                                                                                            v136 := GenS14{F0: (v2 / ((v2) | 1)), F1: (-(v74))}
                                                                                                            _ = v136
                                                                                                        case -159:
                                                                                                            v137 := GenS15{F0: uint64(v1)}
                                                                                                            _ = v137
                                                                                                        }
                                                                                                        var v138 float64 = 0.0
                                                                                                        v5 = append(v5, ("\n\t" + v29))
                                                                                                        v1 -= v1
                                                                                                        _ = v138
                                                                                                    }
                                                                                                    v124 += (v124 ^ (v4 % (((v124 / ((v124) | 1))) | 1)))
                                                                                                    if false {
                                                                                                        switch "" {
                                                                                                        case "hello":
                                                                                                            v5 = append(v5, (((v8 + ((((v8 + (v29 + v8)) + ("a" + v29)) + v29) + (v8 + ((((v8 + v8) + (("hello" + v29) + "\n\t")) + (v8 + ((("" + (v29 + v29)) + v8) + v8))) + v8)))) + ((v29 + "") + v29)) + ((v29 + "hello") + "hello")))
                                                                                                        case "foo bar":
                                                                                                            v5 = append(v5, "a")
                                                                                                        }
                                                                                                    } else {
                                                                                                        v74 -= v74
                                                                                                        v1 -= v1
                                                                                                    }
                                                                                                    _ = v124
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = v29
                                                                                                        }
                                                                                                        v8 = v29
                                                                                                        if (!(v1 < v1)) {
                                                                                                            switch v29 {
                                                                                                            case "a":
                                                                                                                if ((127 == 100) || false) {
                                                                                                                    v139 := GenS16{F0: (fi59 >> ((916) & 63)), F1: (^(v1)), F2: (^(v74))}
                                                                                                                    var v140 bool = (fi59 != fi33)
                                                                                                                    fmt.Println(v32)
                                                                                                                    _ = v139
                                                                                                                    _ = v140
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v4 = v4
                                                                                                                if (!(true && true)) {
                                                                                                                    if true {
                                                                                                                        v34 -= 236
                                                                                                                        v141 := GenS11{F0: 255, F1: 42, F2: (v32 | 739)}
                                                                                                                        var v142 int = 42
                                                                                                                        _ = v141
                                                                                                                        _ = v142
                                                                                                                    } else {
                                                                                                                        v4 = (v4 ^ (v4 / (((-(v4))) | 1)))
                                                                                                                        if true {
                                                                                                                            v143 := GenS17{F0: 485, F1: 2}
                                                                                                                            var v144 float32 = v25
                                                                                                                            _ = v143
                                                                                                                            _ = v144
                                                                                                                        } else {
                                                                                                                            v145 := GenS10{F0: v74, F1: (v29 + (v29 + (v29 + "\n\t"))), F2: 255}
                                                                                                                            v32 = fi33
                                                                                                                            v74 = int32(fi59)
                                                                                                                            _ = v145
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v3 -= (-2.25 * v25)
                                                                                                                }
                                                                                                                var v146 string = v29
                                                                                                                _ = v146
                                                                                                            case "":
                                                                                                                switch v8 {
                                                                                                                case "hello":
                                                                                                                    if true {
                                                                                                                        v8 = v29
                                                                                                                        if len(v5) > 0 {
                                                                                                                            v5[0] = v8
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v147 int = fi33
                                                                                                                        _ = v147
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    if len(v5) > 0 {
                                                                                                                        v5[0] = "foo bar"
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    var v148 int32 = (v74 | v74)
                                                                                                                    _ = v148
                                                                                                                }
                                                                                                                v25 = (v25 / (1e10 * 1e10))
                                                                                                                if (!(true || (!(!(1 != 4294967295))))) {
                                                                                                                    var v149 string = v8
                                                                                                                    _ = v149
                                                                                                                }
                                                                                                                v150 := GenS3{F0: 4294967295}
                                                                                                                _ = v150
                                                                                                            default:
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                            }
                                                                                                            v74 = (v74 / (((v74 - v74)) | 1))
                                                                                                            v24 = v24
                                                                                                            _ = len(v5)
                                                                                                        }
                                                                                                    }
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            fmt.Println(v25, v25, v74)
                                                                                                        }
                                                                                                        fmt.Println(v8)
                                                                                                    }
                                                                                                    switch (((v29 + v29) + v29) + (("" + ("" + (v8 + v29))) + (("" + v29) + "hello"))) {
                                                                                                    case "\n\t":
                                                                                                        var v151 float64 = 0.1
                                                                                                        _ = v151
                                                                                                    case "":
                                                                                                        if len(v5) >= 2 {
                                                                                                            v5 = v5[1:len(v5)]
                                                                                                        }
                                                                                                    default:
                                                                                                        var v152 uint64 = 747
                                                                                                        _ = v152
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if true {
                                                                                                var v153 int16 = (v24 - v24)
                                                                                                if (uint64(v1) == 42) {
                                                                                                    v154 := GenS3{F0: 672}
                                                                                                    v155 := GenS6{F0: (v2 % ((993) | 1))}
                                                                                                    if (!false) {
                                                                                                        v156 := GenS18{F0: (v34 % (((v34 % (((v34 << ((18446744073709551615) & 63))) | 1))) | 1)), F1: (^(fi59)), F2: fi59}
                                                                                                        v32 = (fi59 ^ (v32 >> ((4294967295) & 63)))
                                                                                                        v25 += (v3 + float32(1.5))
                                                                                                        _ = v156
                                                                                                    } else {
                                                                                                        if (true && ((v74 % ((int32(v1)) | 1)) > 0)) {
                                                                                                            fmt.Println(v74, v25)
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v4 += (v4 % (((^(v4))) | 1))
                                                                                                                switch (fi33 * 2147483647) {
                                                                                                                case 2147483647:
                                                                                                                    var v157 int = (fi59 << ((65535) & 63))
                                                                                                                    switch "a" {
                                                                                                                    case "\n\t":
                                                                                                                        switch v29 {
                                                                                                                        case "a":
                                                                                                                            v29 = "a"
                                                                                                                        case "hello":
                                                                                                                            v158 := GenS12{F0: v2, F1: (^(v153)), F2: (true || (!true))}
                                                                                                                            _ = v158
                                                                                                                        case "\n\t":
                                                                                                                            var v159 int = v157
                                                                                                                            _ = v159
                                                                                                                        }
                                                                                                                    case "":
                                                                                                                        var v160 int32 = -1
                                                                                                                        _ = v160
                                                                                                                    case "hello":
                                                                                                                        if len(v5) >= 2 {
                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        switch 127 {
                                                                                                                        case 42:
                                                                                                                            var v161 int32 = int32(v157)
                                                                                                                            _ = v161
                                                                                                                        case 127:
                                                                                                                            v162 := GenS9{F0: v4, F1: v2, F2: 978}
                                                                                                                            _ = v162
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v163 := GenS7{F0: false, F1: true}
                                                                                                                    _ = v157
                                                                                                                    _ = v163
                                                                                                                case 1:
                                                                                                                    v3 += v3
                                                                                                                }
                                                                                                            } else {
                                                                                                                v164 := GenS17{F0: v4, F1: v34}
                                                                                                                switch 0 {
                                                                                                                case 846:
                                                                                                                    fmt.Println(v2)
                                                                                                                case 9223372036854775807:
                                                                                                                    _ = len(v5)
                                                                                                                default:
                                                                                                                    if len(v5) >= 2 {
                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v164
                                                                                                            }
                                                                                                        }
                                                                                                        v153 = int16(v34)
                                                                                                        if false {
                                                                                                            v3 = 0.1
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v154
                                                                                                    _ = v155
                                                                                                }
                                                                                                var v165 uint8 = uint8(v1)
                                                                                                _ = v153
                                                                                                _ = v165
                                                                                            }
                                                                                            v74 -= v74
                                                                                            v2 = (v2 & v2)
                                                                                            fmt.Println(fi59, v3)
                                                                                            _ = len(v5)
                                                                                            v166 := GenS8{F0: ((int8(v1) != -128) && (1 < v24)), F1: -83}
                                                                                            v8 = v29
                                                                                            _ = v166
                                                                                        }
                                                                                    case "":
                                                                                        var v167 int16 = (^(v24))
                                                                                        _ = v167
                                                                                    case "\n\t":
                                                                                        v29 = "foo bar"
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (9223372036854775807 == int(v1)) {
                                                                                fmt.Println(v24, v74, v32)
                                                                            } else {
                                                                                fmt.Println(v3, v34)
                                                                            }
                                                                            v34 += 255
                                                                        }
                                                                        if false {
                                                                            if true {
                                                                                var v168 string = v29
                                                                                _ = v168
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v25, v34)
                                                                    var v169 float64 = (0.1 + 1.5)
                                                                    v170 := GenS15{F0: 247}
                                                                    _ = v169
                                                                    _ = v170
                                                                }
                                                            }
                                                            v29 = v8
                                                            if (!((fi33 | 9223372036854775807) < int(v2))) {
                                                                v171 := GenS19{F0: v34, F1: v29}
                                                                v34 = (v34 % (((v34 ^ 1)) | 1))
                                                                _ = v171
                                                            }
                                                            _ = v74
                                                        }
                                                    } else {
                                                        switch (^(fi33)) {
                                                        case 32767:
                                                            fmt.Println(v24)
                                                        default:
                                                            fmt.Println(fi33, v1, v34)
                                                        }
                                                        if (127 < uint8(v2)) {
                                                            v1 -= 989
                                                        }
                                                    }
                                                    if false {
                                                        fmt.Println(v4, fi59)
                                                        if (4294967295 > 42) {
                                                            if ((v8 < (v29 + v8)) && (false && true)) {
                                                                v24 -= v24
                                                                v172 := GenS20{F0: (2 >= 4294967295)}
                                                                _ = v172
                                                            }
                                                            v24 += (v24 ^ 1)
                                                            v173 := GenS1{F0: v2, F1: (fi33 + (v32 & fi33)), F2: uint16(v2)}
                                                            switch 2147483647 {
                                                            case 2147483647:
                                                                switch -973 {
                                                                case -636:
                                                                    var v174 uint64 = 42
                                                                    _ = v174
                                                                }
                                                            }
                                                            _ = v173
                                                        }
                                                    }
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            } else {
                                                var v175 uint64 = 42
                                                fmt.Println(v29)
                                                v5 = append(v5, "hello")
                                                _ = v175
                                            }
                                            _ = fi59
                                        }
                                    }
                                }
                            default:
                                switch "foo bar" {
                                case "hello":
                                    v1 = 32767
                                default:
                                    switch 128 {
                                    case -1:
                                        fmt.Println(v32)
                                    }
                                }
                            }
                        }
                    } else {
                        v5 = append(v5, v8)
                        v176 := GenS17{F0: 9223372036854775807, F1: (v34 / (((v34 - v34)) | 1))}
                        _ = v176
                    }
                }
            }
            v1 -= v1
        }
        v177 := GenS2{F0: true, F1: ((false && true) || false), F2: 0}
        _ = fi33
        _ = v177
        _ = v34
    }
    v178 := GenS14{F0: (v2 + v2), F1: -1}
    v24 = (v24 / ((v24) | 1))
    if (42 == 1) {
        v8 = "a"
    }
    if true {
        for fi179 := 0; fi179 < 10; fi179++ {
            v180 := GenS21{F0: (v2 % (((v2 + 4294967295)) | 1)), F1: v32, F2: -626}
            for fi181 := 0; fi181 < 7; fi181++ {
                var v182 float64 = -2.25
                _ = fi181
                _ = v182
            }
            v8 = "\n\t"
            _ = fi179
            _ = v180
        }
        fmt.Println(v1, v4)
        fmt.Println(v1, v32)
    }
    var v183 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v183:", v183)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v178
    _ = v183
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v8
}
