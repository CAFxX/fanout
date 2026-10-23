package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 uint8; F2 int32 }
type GenS9 struct { F0 int64; F1 int16; F2 uint8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint64; F1 uint32 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint32; F1 int; F2 int8 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int8; F1 int16; F2 string }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 int8; F1 int32 }
type GenS20 struct { F0 int; F1 uint64 }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 uint64; F1 int; F2 int }
func genFn0(p0 int) int8 {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{(p0 % (((fi5 / ((4294967295) | 1))) | 1))}
        _ = fi5
        _ = v6
    }
    var v7 uint = 255
    if false {
        fmt.Println(v7, p0)
        fmt.Println(v7, v7)
    }
    _ = p0
    _ = v7
    return -90
}

func main() {
    var v1 float64 = (0.1 + 0.1)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 2
    var v4 int32 = -1
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: 100}
        var v10 int16 = -1
        if (546 > 945) {
            _ = len(v2)
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 42}
    v4 = -1
    if (0 >= -214) {
        var v12 float32 = 0.1
        _ = v12
    }
    switch ("\n\t" + "\n\t") {
    case "":
        if false {
            var v13 uint32 = 65535
            for fi14 := 0; fi14 < 4; fi14++ {
                v3 += v3
                fmt.Println(v4, v3)
                _ = fi14
            }
            for fi15 := 0; fi15 < 11; fi15++ {
                for fi16 := 0; fi16 < 15; fi16++ {
                    if (false || (891 < 1)) {
                        var v17 uint32 = v13
                        if (false || false) {
                            v18 := v2[2147483647]
                            v19 := GenS3{F0: v13, F1: 0, F2: 42}
                            _ = v18
                            _ = v19
                        }
                        v4 += v4
                        var v20 uint64 = 24
                        switch "hello" {
                        case "hello":
                            v13 -= (v17 * (v13 - v17))
                        }
                        v1 -= 1.5
                        v21 := v2[v17]
                        v22 := []int{fi16, fi15}
                        v1 -= ((float64((3.14159 - (1e10 + (0.1 + float32(((((-2.25 * v1) + v1) * v1) * v1)))))) + v1) * v1)
                        _ = v17
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        if false {
                            v23 := []int{fi15}
                            _ = v23
                        } else {
                            if ((26 <= (v13 + v13)) || false) {
                                var v24 bool = false
                                v1 += v1
                                _ = v24
                            } else {
                                if ((v4 >> ((42) & 63)) >= (v4 << ((18446744073709551615) & 63))) {
                                    if (0 <= 1) {
                                        fmt.Println(v3, v13, fi15)
                                    } else {
                                        v4 -= v4
                                        fmt.Println(v1, fi16)
                                        v4 -= v4
                                    }
                                }
                            }
                            var v25 int64 = 42
                            switch "a" {
                            case "":
                                var v26 uint8 = 16
                                _ = v26
                            case "foo bar":
                                if ((fi16 < (fi16 % ((fi15) | 1))) || false) {
                                    v27 := GenS4{F0: v25}
                                    fmt.Println(fi16, v25)
                                    switch fi16 {
                                    case -1:
                                        switch (fi15 | fi15) {
                                        case -876:
                                            v4 = (v4 % (((v4 * (v4 / (((v4 % ((-2147483648) | 1))) | 1)))) | 1))
                                        case 127:
                                            v28 := GenS5{F0: (v3 + v3), F1: -52}
                                            v3 = v3
                                            _ = v28
                                        default:
                                            v29 := []int{(-(fi15)), fi15, (fi16 << ((1) & 63)), fi16}
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                            if ((0.0 + 1e10) >= (1.5 + ((0.1 + (0.0 / 0.1)) - 0.0))) {
                                                v13 += v13
                                                var v30 string = "foo bar"
                                                _ = v30
                                            }
                                            _ = v29
                                        }
                                    case 32768:
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            if (true && (!false)) {
                                                v31 := []int{}
                                                v4 -= v4
                                                v3 = (v3 & v3)
                                                _ = v31
                                            }
                                        }
                                    default:
                                        if false {
                                            if (255 <= 614) {
                                                fmt.Println(fi16, v25)
                                            }
                                        } else {
                                            v4 += int32(v13)
                                        }
                                    }
                                    _ = v27
                                }
                            }
                            _ = v25
                        }
                        if true {
                            v3 += uint64(v13)
                            v1 -= v1
                        } else {
                            var v32 int8 = 0
                            fmt.Println(v1, v3, fi15)
                            _ = v32
                        }
                    }
                    v33 := GenS6{F0: 42, F1: 128}
                    v13 += v13
                    _ = fi16
                    _ = v33
                }
                v34 := GenS7{F0: 2}
                var v35 uint8 = 127
                _ = fi15
                _ = v34
                _ = v35
            }
            _ = v13
        }
    case "foo bar":
        for fi36 := 0; fi36 < 17; fi36++ {
            v37 := GenS3{F0: 0, F1: 2, F2: 127}
            _ = fi36
            _ = v37
        }
    default:
        v1 += v1
    }
    var v38 uint64 = 2
    var v39 float32 = (0.1 + 0.1)
    var v40 int8 = 0
    if false {
        delete(v2, 240)
        v41 := []int{-2147483648}
        var v42 uint8 = 0
        _ = v41
        _ = v42
    }
    if true {
        v39 = float32(v1)
        v1 = v1
        if true {
            v39 = (1.5 * v39)
        }
    } else {
        if (false && (((false || (1e10 <= v1)) && false) || false)) {
            v43 := GenS8{F0: 576, F1: 42, F2: (v4 / (((-(v4))) | 1))}
            _ = v43
        }
        if true {
            fmt.Println(v38, v1, v38)
            v38 -= v38
            if (2 >= uint32(v3)) {
                v1 += v1
                switch "" {
                case "hello":
                    var v44 float32 = ((v39 + 0.0) + v39)
                    _ = v44
                case "a":
                    _ = len(v2)
                case "":
                    var v45 float64 = -2.25
                    _ = v45
                default:
                    delete(v2, 2)
                }
            } else {
                v4 = v4
                v4 -= v4
            }
        } else {
            if true {
                for fi46 := 0; fi46 < 12; fi46++ {
                    v47 := GenS2{F0: 0}
                    var v48 float32 = v39
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            } else {
                if (!(!false)) {
                    v49 := GenS3{F0: 370, F1: 65535, F2: 255}
                    if false {
                        v50 := []int{32767, 827, 711}
                        var v51 int8 = v40
                        switch "foo bar" {
                        case "\n\t":
                            switch "a" {
                            case "":
                                if false {
                                    fmt.Println(v3, v38)
                                }
                            case "\n\t":
                                var v52 int = 127
                                _ = v52
                            case "a":
                                if true {
                                    var v53 int64 = 501
                                    if (1e10 > 1e10) {
                                        for fi54 := 0; fi54 < 3; fi54++ {
                                            var v55 int = 1
                                            if (1 != 180) {
                                                for fi56 := 0; fi56 < 20; fi56++ {
                                                    if (552 > 745) {
                                                        switch "hello" {
                                                        case "a":
                                                            v39 -= (1.5 + v39)
                                                        case "foo bar":
                                                            v50 = append(v50, fi56)
                                                        case "\n\t":
                                                            if ("a" != ("" + "")) {
                                                                v57 := GenS9{F0: 648, F1: 32767, F2: 180}
                                                                if false {
                                                                    fmt.Println(v38)
                                                                    var v58 string = ("a" + "foo bar")
                                                                    var v59 int16 = 1
                                                                    _ = v58
                                                                    _ = v59
                                                                } else {
                                                                    var v60 int64 = (v53 & v53)
                                                                    v1 = (v1 / (float64(((v39 / (((0.1 + (1e10 + (v39 * (v39 * v39)))) * (float32(float64(v39)) / 1.0)) / v39)) / (v39 + (v39 - float32((((0.1 / ((((-2.25 * ((0.0 + (v1 / v1)) + v1)) - v1) * v1) * v1)) - (0.0 - v1)) + v1)))))) - v1))
                                                                    _ = v60
                                                                }
                                                                _ = v57
                                                            }
                                                        }
                                                        v4 -= 1
                                                        fmt.Println(v40, v53, v40)
                                                    }
                                                    if ((0 == 42) && (!(true || true))) {
                                                        switch 4294967295 {
                                                        case -941:
                                                            v40 += 100
                                                        default:
                                                            v1 -= 3.14159
                                                        }
                                                        var v61 uint32 = 2147483647
                                                        _ = v61
                                                    }
                                                    if true {
                                                        delete(v2, 42)
                                                        v4 = 1
                                                        v40 = v40
                                                        if false {
                                                            v62 := GenS6{F0: 0, F1: 128}
                                                            _ = v62
                                                        }
                                                        switch (fi56 | (fi54 | 1)) {
                                                        case 4294967295:
                                                            var v63 float32 = float32((float64(v39) + ((v1 + 1e10) / v1)))
                                                            _ = v63
                                                        case -1:
                                                            if true {
                                                                delete(v2, 2147483647)
                                                                v50 = append(v50, (v55 ^ 938))
                                                            } else {
                                                                var v64 uint16 = 65535
                                                                fmt.Println(v38)
                                                                _ = v64
                                                            }
                                                        case 74:
                                                            var v65 uint64 = (v38 & (^(v3)))
                                                            _ = v65
                                                        }
                                                        switch v55 {
                                                        case 32768:
                                                            v66 := GenS10{F0: 2147483648}
                                                            _ = v66
                                                        case 42:
                                                            _ = len(v50)
                                                        }
                                                    } else {
                                                        v4 += (v4 | (v4 | v4))
                                                    }
                                                    _ = fi56
                                                }
                                                for fi67 := 0; fi67 < 11; fi67++ {
                                                    var v68 uint32 = 0
                                                    v68 += 65535
                                                    _ = fi67
                                                    _ = v68
                                                }
                                            } else {
                                                var v69 uint64 = v38
                                                _ = v69
                                            }
                                            switch int(v51) {
                                            case 32768:
                                                for fi70 := 0; fi70 < 8; fi70++ {
                                                    delete(v2, 989)
                                                    var v71 int32 = v4
                                                    _ = fi70
                                                    _ = v71
                                                }
                                            case 128:
                                                for fi72 := 0; fi72 < 16; fi72++ {
                                                    var v73 int8 = (v40 % (((v51 - -1)) | 1))
                                                    if true {
                                                        fmt.Println(v40, v1, fi54)
                                                        fmt.Println(v38, v4)
                                                        var v74 string = "foo bar"
                                                        _ = v74
                                                    }
                                                    v40 -= v73
                                                    _ = fi72
                                                    _ = v73
                                                }
                                            case 680:
                                                v75 := GenS8{F0: (v53 + (v53 | 0)), F1: 127, F2: (v4 ^ v4)}
                                                _ = v75
                                            }
                                            _ = fi54
                                            _ = v55
                                        }
                                        switch "\n\t" {
                                        case "":
                                            for fi76 := 0; fi76 < 7; fi76++ {
                                                v77 := GenS11{F0: v4}
                                                _ = len(v50)
                                                _ = fi76
                                                _ = v77
                                            }
                                        case "foo bar":
                                            if (42 <= 4294967295) {
                                                for fi78 := 0; fi78 < 19; fi78++ {
                                                    v39 = v39
                                                    _ = fi78
                                                }
                                            }
                                        default:
                                            if (v53 <= v53) {
                                                v79 := GenS12{F0: (v3 | v3), F1: 2147483648}
                                                if (false || ("foo bar" <= ("foo bar" + ""))) {
                                                    var v80 int16 = 1000
                                                    var v81 uint8 = 141
                                                    _ = v80
                                                    _ = v81
                                                }
                                                _ = v79
                                            }
                                        }
                                        for fi82 := 0; fi82 < 11; fi82++ {
                                            if (128 <= (fi82 % ((fi82) | 1))) {
                                                var v83 bool = true
                                                switch (fi82 >> ((65535) & 63)) {
                                                case -9223372036854775808:
                                                    if true {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v83 = true
                                                        case "a":
                                                            var v84 uint8 = 96
                                                            _ = v84
                                                        case "":
                                                            fmt.Println(fi82, v51)
                                                        default:
                                                            var v85 int8 = v40
                                                            _ = v85
                                                        }
                                                        var v86 int16 = 604
                                                        _ = v86
                                                    }
                                                case 9223372036854775807:
                                                    var v87 float64 = -2.25
                                                    _ = v87
                                                }
                                                _ = v83
                                            } else {
                                                var v88 int64 = (v53 & 42)
                                                v89 := GenS4{F0: int64(v4)}
                                                for fi90 := 0; fi90 < 12; fi90++ {
                                                    v91 := v2[42]
                                                    v38 -= (v38 << ((18446744073709551615) & 63))
                                                    _ = fi90
                                                    _ = v91
                                                }
                                                _ = v88
                                                _ = v89
                                            }
                                            v2[1] = ((v3 * v3) == v38)
                                            fmt.Println(v4, v38, v38)
                                            _ = fi82
                                        }
                                    } else {
                                        if (!false) {
                                            v92 := GenS8{F0: 0, F1: 255, F2: 222}
                                            _ = v92
                                        }
                                        if true {
                                            switch "\n\t" {
                                            case "foo bar":
                                                if true {
                                                    v93 := GenS12{F0: 1, F1: 2}
                                                    _ = v93
                                                } else {
                                                    if false {
                                                        switch ("foo bar" + ((("foo bar" + ("hello" + "a")) + "foo bar") + "\n\t")) {
                                                        case "a":
                                                            v38 = (v3 + v3)
                                                        case "hello":
                                                            for fi94 := 0; fi94 < 13; fi94++ {
                                                                for fi95 := 0; fi95 < 1; fi95++ {
                                                                    v96 := GenS2{F0: 866}
                                                                    _ = len(v50)
                                                                    _ = fi95
                                                                    _ = v96
                                                                }
                                                                v4 += v4
                                                                v50 = append(v50, fi94)
                                                                _ = fi94
                                                            }
                                                        }
                                                        v97 := GenS9{F0: v53, F1: -453, F2: 42}
                                                        _ = v97
                                                    } else {
                                                        if true {
                                                            v98 := GenS13{F0: 2}
                                                            v38 += 2
                                                            _ = v98
                                                        } else {
                                                            v38 = (v38 & v38)
                                                            for fi99 := 0; fi99 < 16; fi99++ {
                                                                fmt.Println(v38, fi99)
                                                                _ = fi99
                                                            }
                                                        }
                                                        var v100 int8 = (v51 * 1)
                                                        _ = v100
                                                    }
                                                }
                                            case "":
                                                if true {
                                                    var v101 int32 = v4
                                                    var v102 float64 = 1e10
                                                    if (4294967295 >= 18446744073709551615) {
                                                        if false {
                                                            if (380 != (v38 << ((4294967295) & 63))) {
                                                                var v103 float64 = (v1 + v102)
                                                                _ = v103
                                                            }
                                                            v102 = (float64(0.0) + v1)
                                                        } else {
                                                            switch ("\n\t" + ("hello" + "foo bar")) {
                                                            case "":
                                                                v40 += v40
                                                            case "hello":
                                                                if true {
                                                                    if len(v50) > 0 {
                                                                        v50[0] = 32768
                                                                    }
                                                                    _ = len(v2)
                                                                    var v104 bool = true
                                                                    _ = v104
                                                                } else {
                                                                    v40 += -128
                                                                    var v105 int = 127
                                                                    _ = v105
                                                                }
                                                            case "a":
                                                                if false {
                                                                    v106 := GenS8{F0: (v53 | (-(v53))), F1: 1, F2: int32(v53)}
                                                                    fmt.Println(v3)
                                                                    _ = v106
                                                                } else {
                                                                    var v107 int32 = (v101 & 734)
                                                                    for fi108 := 0; fi108 < 18; fi108++ {
                                                                        fmt.Println(v101, v51, v40)
                                                                        for fi109 := 0; fi109 < 9; fi109++ {
                                                                            fmt.Println(v101, v40)
                                                                            switch fi108 {
                                                                            case 40:
                                                                                v39 += v39
                                                                            case -2147483648:
                                                                                var v110 int16 = -420
                                                                                _ = v110
                                                                            default:
                                                                                switch fi109 {
                                                                                case 9223372036854775807:
                                                                                    v1 = float64((0.0 + (v39 * (((0.1 + v39) - 0.0) + 1.5))))
                                                                                case 2147483647:
                                                                                    fmt.Println(v102, v38)
                                                                                case 32767:
                                                                                    v50 = append(v50, (fi108 | 32767))
                                                                                }
                                                                            }
                                                                            _ = fi109
                                                                        }
                                                                        var v111 uint8 = 127
                                                                        _ = fi108
                                                                        _ = v111
                                                                    }
                                                                    for fi112 := 0; fi112 < 1; fi112++ {
                                                                        for fi113 := 0; fi113 < 17; fi113++ {
                                                                            v101 = (v4 << ((481) & 63))
                                                                            _ = fi113
                                                                        }
                                                                        switch (fi112 - int(v53)) {
                                                                        case -9223372036854775808:
                                                                            for fi114 := 0; fi114 < 18; fi114++ {
                                                                                var v115 int8 = v51
                                                                                v51 -= (-(v115))
                                                                                _ = fi114
                                                                                _ = v115
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                _ = len(v2)
                                                                            } else {
                                                                                var v116 int64 = v53
                                                                                _ = v116
                                                                            }
                                                                        }
                                                                        if (!false) {
                                                                            v117 := GenS14{F0: ("a" + "\n\t")}
                                                                            if ((2147483648 >= 2147483648) || ((3.14159 + v1) != float64(0.0))) {
                                                                                switch ("a" + ("hello" + "hello")) {
                                                                                case "\n\t":
                                                                                    fmt.Println(v38)
                                                                                case "":
                                                                                    v107 += -833
                                                                                }
                                                                                v118 := v2[2147483648]
                                                                                v107 -= v101
                                                                                _ = v118
                                                                            } else {
                                                                                if true {
                                                                                    _ = len(v2)
                                                                                    var v119 string = (("\n\t" + "") + "\n\t")
                                                                                    _ = v119
                                                                                } else {
                                                                                    for fi120 := 0; fi120 < 13; fi120++ {
                                                                                        _ = len(v2)
                                                                                        v121 := GenS14{F0: (("hello" + "\n\t") + ((((("\n\t" + ("a" + "hello")) + ("hello" + ("" + "foo bar"))) + ("" + "")) + "hello") + "a"))}
                                                                                        v51 = v51
                                                                                        if (false || false) {
                                                                                            if (!(494 >= 65535)) {
                                                                                                if true {
                                                                                                    _ = len(v50)
                                                                                                    fmt.Println(v4, v101)
                                                                                                }
                                                                                            }
                                                                                            if ((false || (fi120 != int(v4))) || false) {
                                                                                                v102 = v102
                                                                                            } else {
                                                                                                if (!(false && false)) {
                                                                                                    if len(v50) > 0 {
                                                                                                        v50[0] = -2147483648
                                                                                                    }
                                                                                                    v122 := GenS15{F0: 4294967295, F1: fi120, F2: (v51 ^ v40)}
                                                                                                    v3 += 0
                                                                                                    var v123 int16 = -1
                                                                                                    if false {
                                                                                                        switch "foo bar" {
                                                                                                        case "foo bar":
                                                                                                            switch fi112 {
                                                                                                            case 128:
                                                                                                                v3 -= 1
                                                                                                            case -2147483648:
                                                                                                                switch "foo bar" {
                                                                                                                case "a":
                                                                                                                    v124 := GenS10{F0: uint32(v38)}
                                                                                                                    _ = v124
                                                                                                                }
                                                                                                            case 127:
                                                                                                                if (v3 == 0) {
                                                                                                                    fmt.Println(v40, fi112, fi120)
                                                                                                                    v1 = ((0.0 - 1.5) + v102)
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        _ = len(v2)
                                                                                                                        var v125 float64 = v102
                                                                                                                        _ = v125
                                                                                                                    }
                                                                                                                    v53 = (^(v53))
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            switch (^(fi112)) {
                                                                                                            case 1:
                                                                                                                v4 += (v4 & int32(v38))
                                                                                                            case 128:
                                                                                                                switch "" {
                                                                                                                case "":
                                                                                                                    v123 = v123
                                                                                                                case "\n\t":
                                                                                                                    fmt.Println(v4)
                                                                                                                default:
                                                                                                                    var v126 uint16 = 648
                                                                                                                    _ = v126
                                                                                                                }
                                                                                                            default:
                                                                                                                var v127 uint64 = v38
                                                                                                                _ = v127
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        v51 = (-(v51))
                                                                                                        v128 := GenS5{F0: v38, F1: v51}
                                                                                                        fmt.Println(v51)
                                                                                                        _ = v128
                                                                                                    }
                                                                                                    _ = v122
                                                                                                    _ = v123
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = fi120
                                                                                        _ = v121
                                                                                    }
                                                                                    if false {
                                                                                        v129 := v2[4294967295]
                                                                                        switch -869 {
                                                                                        case 127:
                                                                                            var v130 int = fi112
                                                                                            _ = v130
                                                                                        case 9223372036854775807:
                                                                                            for fi131 := 0; fi131 < 3; fi131++ {
                                                                                                var v132 uint64 = v38
                                                                                                _ = fi131
                                                                                                _ = v132
                                                                                            }
                                                                                        case -9223372036854775808:
                                                                                            v1 += 3.14159
                                                                                        default:
                                                                                            v133 := GenS9{F0: (v53 ^ (v53 + -604)), F1: 32767, F2: 66}
                                                                                            _ = v133
                                                                                        }
                                                                                        for fi134 := 0; fi134 < 12; fi134++ {
                                                                                            fmt.Println(v102)
                                                                                            _ = fi134
                                                                                        }
                                                                                        _ = v129
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v117
                                                                        }
                                                                        _ = fi112
                                                                    }
                                                                    _ = v107
                                                                }
                                                            default:
                                                                switch -451 {
                                                                case 127:
                                                                    if (65535 < 335) {
                                                                        for fi135 := 0; fi135 < 5; fi135++ {
                                                                            _ = len(v50)
                                                                            _ = fi135
                                                                        }
                                                                        v4 = v4
                                                                    }
                                                                case 749:
                                                                    v101 += (v4 * (v4 | (v101 >> ((18446744073709551615) & 63))))
                                                                case 2147483647:
                                                                    switch -1 {
                                                                    case 32768:
                                                                        if true {
                                                                            switch 128 {
                                                                            case -739:
                                                                                v51 = (v40 / ((genFn0(127)) | 1))
                                                                            case -2147483648:
                                                                                for fi136 := 0; fi136 < 18; fi136++ {
                                                                                    if (false && false) {
                                                                                        fmt.Println(v39, v40, v3)
                                                                                        v2[65535] = false
                                                                                    }
                                                                                    _ = fi136
                                                                                }
                                                                            }
                                                                            v137 := GenS4{F0: (v53 >> ((545) & 63))}
                                                                            if true {
                                                                                _ = len(v2)
                                                                                var v138 int16 = 305
                                                                                _ = v138
                                                                            }
                                                                            _ = v137
                                                                        } else {
                                                                            if (!false) {
                                                                                var v139 uint64 = (v38 + v38)
                                                                                if true {
                                                                                    fmt.Println(v40, v139)
                                                                                    var v140 int32 = 2147483647
                                                                                    v53 += -1
                                                                                    _ = v140
                                                                                } else {
                                                                                    for fi141 := 0; fi141 < 4; fi141++ {
                                                                                        if true {
                                                                                            v142 := GenS14{F0: "hello"}
                                                                                            var v143 float64 = (v102 * ((((float64(v39) / 3.14159) - (3.14159 / v1)) + v102) / ((v1 - 0.1) - v1)))
                                                                                            var v144 int16 = -152
                                                                                            _ = v142
                                                                                            _ = v143
                                                                                            _ = v144
                                                                                        }
                                                                                        _ = fi141
                                                                                    }
                                                                                }
                                                                                _ = v139
                                                                            }
                                                                            var v145 float32 = 0.0
                                                                            var v146 uint8 = 128
                                                                            _ = v145
                                                                            _ = v146
                                                                        }
                                                                    case 588:
                                                                        if false {
                                                                            var v147 bool = false
                                                                            v4 += v4
                                                                            v50 = append(v50, 0)
                                                                            _ = v147
                                                                        } else {
                                                                            fmt.Println(v101, v51, v4)
                                                                        }
                                                                    case 473:
                                                                        fmt.Println(v51, v3, v51)
                                                                    default:
                                                                        fmt.Println(v51, v40, v38)
                                                                    }
                                                                }
                                                            }
                                                            v148 := GenS15{F0: 380, F1: -683, F2: v51}
                                                            _ = v148
                                                        }
                                                        var v149 uint64 = 140
                                                        v150 := v2[2147483647]
                                                        _ = v149
                                                        _ = v150
                                                    } else {
                                                        for fi151 := 0; fi151 < 6; fi151++ {
                                                            var v152 uint8 = 52
                                                            fmt.Println(v101, v40, v51)
                                                            v153 := GenS16{F0: (("" + "hello") + "hello")}
                                                            _ = fi151
                                                            _ = v152
                                                            _ = v153
                                                        }
                                                    }
                                                    _ = v101
                                                    _ = v102
                                                }
                                            default:
                                                v4 = v4
                                            }
                                            var v154 float64 = v1
                                            _ = v154
                                        } else {
                                            var v155 uint16 = 198
                                            var v156 uint16 = v155
                                            var v157 string = "foo bar"
                                            _ = v155
                                            _ = v156
                                            _ = v157
                                        }
                                        switch int(v53) {
                                        case 1:
                                            v3 = v38
                                        case -433:
                                            if (!(true || false)) {
                                                fmt.Println(v3, v40)
                                            } else {
                                                if (42 <= 1) {
                                                    var v158 float64 = float64(v39)
                                                    _ = v158
                                                }
                                                if (v4 != 42) {
                                                    fmt.Println(v51, v51, v4)
                                                    _ = len(v50)
                                                } else {
                                                    v53 -= v53
                                                }
                                                if ((("\n\t" + "") >= "") || false) {
                                                    for fi159 := 0; fi159 < 2; fi159++ {
                                                        var v160 int = (fi159 | (-(fi159)))
                                                        switch "foo bar" {
                                                        case "":
                                                            for fi161 := 0; fi161 < 14; fi161++ {
                                                                switch ("\n\t" + "") {
                                                                case "hello":
                                                                    var v162 float32 = (0.0 - -2.25)
                                                                    _ = v162
                                                                }
                                                                v50 = append(v50, fi159)
                                                                _ = fi161
                                                            }
                                                        default:
                                                            if true {
                                                                fmt.Println(v51, v51, v51)
                                                                switch 127 {
                                                                case -1:
                                                                    v4 = (v4 % ((515) | 1))
                                                                case 32768:
                                                                    var v163 float64 = 1.5
                                                                    _ = v163
                                                                case -9223372036854775808:
                                                                    switch (^(fi159)) {
                                                                    case 42:
                                                                        v164 := GenS9{F0: (v53 / (((v53 & (v53 | int64(v51)))) | 1)), F1: -650, F2: 1}
                                                                        _ = v164
                                                                    case -780:
                                                                        if ((fi159 ^ v160) == (^(v160))) {
                                                                            if ((v1 * v1) >= v1) {
                                                                                if false {
                                                                                    if (!true) {
                                                                                        switch (((("hello" + "hello") + "a") + ("\n\t" + "a")) + "hello") {
                                                                                        case "a":
                                                                                            if ((v1 * v1) > ((v1 / float64(float32((v1 / 1.0)))) * (v1 + v1))) {
                                                                                                v3 = 1
                                                                                                v165 := GenS14{F0: "\n\t"}
                                                                                                var v166 uint16 = 631
                                                                                                _ = v165
                                                                                                _ = v166
                                                                                            } else {
                                                                                                v167 := GenS14{F0: ""}
                                                                                                _ = v167
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v51 = (-(v40))
                                                                                        default:
                                                                                            for fi168 := 0; fi168 < 16; fi168++ {
                                                                                                if true {
                                                                                                    if (70 == 0) {
                                                                                                        var v169 int64 = v53
                                                                                                        if len(v50) > 0 {
                                                                                                            v50[0] = 717
                                                                                                        }
                                                                                                        _ = v169
                                                                                                    } else {
                                                                                                        v4 = 1
                                                                                                        v3 += v38
                                                                                                        v3 = 1
                                                                                                    }
                                                                                                    v160 -= (fi159 * 1)
                                                                                                } else {
                                                                                                    v53 -= 0
                                                                                                    v170 := GenS17{F0: (v51 ^ (v40 | genFn0((fi168 & fi159)))), F1: 0, F2: "hello"}
                                                                                                    if ("foo bar" < "foo bar") {
                                                                                                        var v171 int32 = v4
                                                                                                        v160 = v160
                                                                                                        v4 -= (v171 << ((0) & 63))
                                                                                                        _ = v171
                                                                                                    }
                                                                                                    _ = v170
                                                                                                }
                                                                                                _ = fi168
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v160 -= v160
                                                                                }
                                                                                for fi172 := 0; fi172 < 13; fi172++ {
                                                                                    v160 += fi159
                                                                                    if (128 < fi172) {
                                                                                        delete(v2, 4294967295)
                                                                                        if ((fi172 & (fi159 & 1)) == fi159) {
                                                                                            fmt.Println(v1)
                                                                                            if false {
                                                                                                _ = len(v50)
                                                                                                v38 = v38
                                                                                            } else {
                                                                                                if (("a" != "\n\t") || true) {
                                                                                                    v51 -= v40
                                                                                                    var v173 float64 = (1e10 - 1e10)
                                                                                                    _ = len(v2)
                                                                                                    if false {
                                                                                                        if true {
                                                                                                            if (1 >= 2) {
                                                                                                                fmt.Println(v38)
                                                                                                            }
                                                                                                            v39 = v39
                                                                                                        } else {
                                                                                                            v1 += 1.5
                                                                                                            fmt.Println(fi172)
                                                                                                            if len(v50) >= 2 {
                                                                                                                v50 = v50[1:len(v50)]
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v38)
                                                                                                        if (0 > 2) {
                                                                                                            if len(v50) > 0 {
                                                                                                                v50[0] = int(v4)
                                                                                                            }
                                                                                                            var v174 float32 = float32(v173)
                                                                                                            if len(v50) > 0 {
                                                                                                                v50[0] = (v160 | (fi172 ^ (-(v160))))
                                                                                                            }
                                                                                                            _ = v174
                                                                                                        } else {
                                                                                                            v4 = int32(v53)
                                                                                                            fmt.Println(v38, v160)
                                                                                                            v40 -= 0
                                                                                                        }
                                                                                                        var v175 bool = true
                                                                                                        var v176 uint16 = 2
                                                                                                        _ = len(v2)
                                                                                                        _ = v175
                                                                                                        _ = v176
                                                                                                    }
                                                                                                    v177 := v2[2147483647]
                                                                                                    if len(v50) >= 2 {
                                                                                                        v50 = v50[1:len(v50)]
                                                                                                    }
                                                                                                    var v178 int64 = v53
                                                                                                    _ = v173
                                                                                                    _ = v177
                                                                                                    _ = v178
                                                                                                }
                                                                                                var v179 float32 = (v39 + 1e10)
                                                                                                v160 = -808
                                                                                                var v180 uint8 = 1
                                                                                                _ = v179
                                                                                                _ = v180
                                                                                            }
                                                                                            v40 = (v51 - v40)
                                                                                        }
                                                                                    }
                                                                                    v4 = -1
                                                                                    _ = fi172
                                                                                }
                                                                                var v181 int = (fi159 + (^(v160)))
                                                                                _ = v181
                                                                            } else {
                                                                                if (127 < (v160 * v160)) {
                                                                                    fmt.Println(v51, v3, fi159)
                                                                                    switch (v160 + (fi159 + (v160 % ((v160) | 1)))) {
                                                                                    case -466:
                                                                                        var v182 float64 = (0.0 / 3.14159)
                                                                                        _ = v182
                                                                                    case -2147483648:
                                                                                        if (!false) {
                                                                                            if true {
                                                                                                fmt.Println(v160)
                                                                                                for fi183 := 0; fi183 < 15; fi183++ {
                                                                                                    v1 = 0.1
                                                                                                    _ = fi183
                                                                                                }
                                                                                            }
                                                                                            v184 := GenS14{F0: (("hello" + "a") + "hello")}
                                                                                            _ = v184
                                                                                        } else {
                                                                                            var v185 uint16 = 730
                                                                                            for fi186 := 0; fi186 < 8; fi186++ {
                                                                                                fmt.Println(v160)
                                                                                                var v187 uint = 4294967295
                                                                                                _ = fi186
                                                                                                _ = v187
                                                                                            }
                                                                                            if false {
                                                                                                for fi188 := 0; fi188 < 13; fi188++ {
                                                                                                    v4 -= (v4 | (v4 % ((v4) | 1)))
                                                                                                    _ = fi188
                                                                                                }
                                                                                                switch (v160 - fi159) {
                                                                                                case -1:
                                                                                                    switch v160 {
                                                                                                    case -9223372036854775808:
                                                                                                        fmt.Println(v51, v4)
                                                                                                    case 0:
                                                                                                        var v189 bool = (true || (v4 != -1))
                                                                                                        _ = v189
                                                                                                    default:
                                                                                                        var v190 float32 = (v39 / 1.5)
                                                                                                        _ = v190
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v185
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v53, v51)
                                                                                }
                                                                                for fi191 := 0; fi191 < 7; fi191++ {
                                                                                    switch fi191 {
                                                                                    case 32768:
                                                                                        v192 := GenS17{F0: v40, F1: -1, F2: "hello"}
                                                                                        _ = v192
                                                                                    case 4294967295:
                                                                                        var v193 uint32 = 926
                                                                                        _ = v193
                                                                                    default:
                                                                                        _ = len(v50)
                                                                                    }
                                                                                    _ = fi191
                                                                                }
                                                                            }
                                                                            switch (v160 >> ((130) & 63)) {
                                                                            case 128:
                                                                                var v194 int = fi159
                                                                                _ = v194
                                                                            default:
                                                                                for fi195 := 0; fi195 < 18; fi195++ {
                                                                                    fmt.Println(v4, v1, v51)
                                                                                    _ = fi195
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    var v196 uint32 = 941
                                                                    _ = v196
                                                                }
                                                                _ = len(v50)
                                                            } else {
                                                                if false {
                                                                    if (255 != 127) {
                                                                        fmt.Println(v40)
                                                                    } else {
                                                                        if false {
                                                                            v39 -= (((float32(((3.14159 + v1) * (((v1 * float64(-2.25)) / 1.5) / (0.1 / v1)))) / (v39 * float32((v1 - (((3.14159 / v1) + 3.14159) + 0.0))))) / v39) / v39)
                                                                            for fi197 := 0; fi197 < 6; fi197++ {
                                                                                if (v38 >= v38) {
                                                                                    switch int(v4) {
                                                                                    case -2147483648:
                                                                                        v198 := GenS18{F0: v4}
                                                                                        _ = v198
                                                                                    case 4294967295:
                                                                                        var v199 int = (v160 >> ((0) & 63))
                                                                                        _ = v199
                                                                                    default:
                                                                                        if len(v50) > 0 {
                                                                                            v50[0] = 0
                                                                                        }
                                                                                    }
                                                                                    v39 = 3.14159
                                                                                    v50 = append(v50, int(v4))
                                                                                } else {
                                                                                    v200 := GenS19{F0: v51, F1: v4}
                                                                                    _ = v200
                                                                                }
                                                                                _ = fi197
                                                                            }
                                                                        }
                                                                        for fi201 := 0; fi201 < 13; fi201++ {
                                                                            var v202 float64 = ((float64((v39 + 3.14159)) - 0.0) / v1)
                                                                            fmt.Println(fi159)
                                                                            var v203 uint32 = 2
                                                                            _ = fi201
                                                                            _ = v202
                                                                            _ = v203
                                                                        }
                                                                        var v204 uint = 1
                                                                        _ = v204
                                                                    }
                                                                    if len(v50) > 0 {
                                                                        v50[0] = (fi159 & 4294967295)
                                                                    }
                                                                    switch v160 {
                                                                    case -779:
                                                                        v205 := GenS20{F0: fi159, F1: (v3 % (((v3 ^ v3)) | 1))}
                                                                        _ = v205
                                                                    }
                                                                }
                                                                if ("foo bar" != "hello") {
                                                                    for fi206 := 0; fi206 < 6; fi206++ {
                                                                        fmt.Println(v38, v51, v38)
                                                                        _ = fi206
                                                                    }
                                                                } else {
                                                                    var v207 uint8 = 128
                                                                    fmt.Println(fi159, v1)
                                                                    _ = v207
                                                                }
                                                            }
                                                        }
                                                        v53 -= v53
                                                        _ = fi159
                                                        _ = v160
                                                    }
                                                } else {
                                                    v2[20] = ((^(v53)) > v53)
                                                    fmt.Println(v53)
                                                }
                                            }
                                        case 2147483647:
                                            v208 := GenS15{F0: 65535, F1: 2147483647, F2: v51}
                                            _ = v208
                                        default:
                                            fmt.Println(v40)
                                        }
                                    }
                                    if false {
                                        v209 := GenS10{F0: 1}
                                        for fi210 := 0; fi210 < 18; fi210++ {
                                            fmt.Println(v39, v4)
                                            if true {
                                                var v211 float64 = v1
                                                v212 := GenS21{F0: 1}
                                                v213 := GenS22{F0: (v3 | (v3 + (v38 * 18446744073709551615))), F1: fi210, F2: fi210}
                                                _ = v211
                                                _ = v212
                                                _ = v213
                                            }
                                            v40 = (v40 | (v51 / (((v40 & 26)) | 1)))
                                            _ = fi210
                                        }
                                        _ = v209
                                    }
                                    _ = v53
                                }
                            }
                        case "foo bar":
                            if (-132 < -102) {
                                v4 -= v4
                            }
                        }
                        _ = v50
                        _ = v51
                    }
                    if ((42 >= 2) || true) {
                        var v214 uint16 = 1
                        _ = v214
                    } else {
                        switch 152 {
                        case -1:
                            v39 += ((float32(0.0) - v39) + float32(1.5))
                        }
                        var v215 int = 1
                        var v216 int64 = 42
                        _ = v215
                        _ = v216
                    }
                    _ = v49
                } else {
                    for fi217 := 0; fi217 < 14; fi217++ {
                        v218 := GenS21{F0: 1}
                        var v219 int16 = 1000
                        if (v1 != 1e10) {
                            v220 := []int{fi217, fi217, -891, fi217}
                            var v221 uint = 354
                            _ = v220
                            _ = v221
                        }
                        _ = fi217
                        _ = v218
                        _ = v219
                    }
                    v40 = v40
                }
                var v222 int32 = 1
                _ = v222
            }
            var v223 int16 = 42
            var v224 bool = true
            _ = v223
            _ = v224
        }
    }
    fmt.Println(v38, v38, v40)
    if true {
        for fi225 := 0; fi225 < 3; fi225++ {
            if (false || (!(!(!(!(genFn0((fi225 >> ((65535) & 63))) != genFn0(fi225))))))) {
                v226 := []int{fi225, (^(fi225)), 127}
                if len(v226) >= 2 {
                    v226 = v226[1:len(v226)]
                }
                _ = v226
            } else {
                for fi227 := 0; fi227 < 3; fi227++ {
                    v1 += (float64(1.5) - v1)
                    v228 := GenS3{F0: 42, F1: 783, F2: 39}
                    if false {
                        var v229 int = fi225
                        _ = len(v2)
                        _ = v229
                    }
                    _ = fi227
                    _ = v228
                }
                switch fi225 {
                case -329:
                    v230 := GenS13{F0: 18446744073709551615}
                    _ = v230
                case 182:
                    v2[1] = false
                default:
                    v231 := GenS6{F0: 42, F1: 128}
                    _ = v231
                }
                switch fi225 {
                case -238:
                    var v232 uint16 = 255
                    _ = v232
                default:
                    var v233 uint8 = 2
                    _ = v233
                }
            }
            var v234 bool = (255 >= 42)
            fmt.Println(v3)
            _ = fi225
            _ = v234
        }
        for fi235 := 0; fi235 < 9; fi235++ {
            fmt.Println(v4, v39)
            var v236 uint16 = 1
            if false {
                for fi237 := 0; fi237 < 5; fi237++ {
                    v1 = 1.5
                    _ = fi237
                }
                v238 := GenS10{F0: 956}
                _ = v238
            }
            _ = fi235
            _ = v236
        }
    } else {
        v239 := v2[2]
        _ = v239
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
