package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint; F2 uint64 }
type GenS5 struct { F0 int; F1 uint8; F2 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 int8; F2 string }
type GenS12 struct { F0 uint64; F1 int32; F2 int64 }
type GenS13 struct { F0 uint32; F1 int }
type GenS14 struct { F0 uint32; F1 int16; F2 int8 }
type GenS15 struct { F0 uint64; F1 int8 }
type GenS16 struct { F0 uint32; F1 uint64; F2 uint }
type GenS17 struct { F0 int }
type GenS18 struct { F0 string; F1 int32 }
type GenS19 struct { F0 string; F1 uint64 }
type GenS20 struct { F0 uint32; F1 string }
type GenS21 struct { F0 int16; F1 uint }
type GenS22 struct { F0 uint64; F1 int8 }
type GenS23 struct { F0 int; F1 uint32; F2 uint }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 string; F1 bool; F2 int16 }
type GenS26 struct { F0 uint }
type GenS27 struct { F0 int64 }

func main() {
    v1 := []uint32{}
    var v2 uint8 = 178
    v3 := []int{32768, 1, 127}
    var v4 int = 563
    var v5 int16 = 52
    v5 = (v5 ^ 1000)
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            _ = len(v1)
            switch "hello" {
            case "foo bar":
                v3 = append(v3, fi6)
            case "a":
                v1 = append(v1, 2)
            case "hello":
                var v8 uint64 = 4294967295
                _ = v8
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi7
        }
        _ = fi6
    }
    var v9 uint64 = 2
    var v10 int64 = -1
    var v11 uint64 = (v9 << ((65535) & 63))
    v4 -= (v4 - -544)
    switch "\n\t" {
    case "":
        var v12 bool = (!false)
        _ = v12
    case "a":
        for fi13 := 0; fi13 < 17; fi13++ {
            _ = len(v3)
            _ = fi13
        }
    case "foo bar":
        switch v4 {
        case 127:
            if true {
                var v14 string = "hello"
                switch "\n\t" {
                case "a":
                    v15 := GenS1{F0: 304}
                    _ = v15
                case "\n\t":
                    if false {
                        fmt.Println(v14, v11, v2)
                    }
                default:
                    var v16 int16 = (v5 * (v5 + (v5 & 32767)))
                    _ = v16
                }
                if ("a" == ((v14 + v14) + ("a" + "a"))) {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v18 := GenS1{F0: 2}
                        _ = fi17
                        _ = v18
                    }
                } else {
                    v19 := GenS2{F0: 2147483648, F1: -1}
                    switch (v14 + v14) {
                    case "\n\t":
                        v20 := GenS2{F0: 452, F1: -1}
                        _ = v20
                    case "hello":
                        var v21 int = v4
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 5; fi22++ {
                            if (0 > 1) {
                                if (((1e10 + -2.25) - float32(-2.25)) <= (float32(1.5) / (float32(((1.5 + 0.0) / float64(float32(((1e10 / (0.1 / float64(1.5))) - float64(3.14159)))))) + ((((-2.25 - (((((3.14159 / ((1.5 - (3.14159 - 0.1)) + 0.1)) + 3.14159) + 0.1) + (0.0 * -2.25)) + 3.14159)) * (1.5 * 0.1)) - float32((3.14159 * (float64(((-2.25 * 3.14159) / 1e10)) / 1e10)))) * (-2.25 / (((3.14159 - (0.0 - 0.1)) + (3.14159 - 0.0)) - 1.5)))))) {
                                    v10 += v10
                                    fmt.Println(v2, fi22, v10)
                                }
                                for fi23 := 0; fi23 < 18; fi23++ {
                                    v14 = (v14 + (v14 + "a"))
                                    _ = fi23
                                }
                                for fi24 := 0; fi24 < 10; fi24++ {
                                    var v25 uint8 = (v2 + (v2 / (((v2 + (v2 + (v2 << ((233) & 63))))) | 1)))
                                    if len(v1) > 0 {
                                        v1[0] = 560
                                    }
                                    v26 := GenS3{F0: "a"}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                            }
                            var v27 uint = 659
                            if true {
                                v1 = append(v1, 1)
                                if false {
                                    switch fi22 {
                                    case 4294967295:
                                        switch (-(fi22)) {
                                        case 32767:
                                            v27 = v27
                                        }
                                    case -9223372036854775808:
                                        v28 := GenS1{F0: 749}
                                        _ = v28
                                    case 316:
                                        switch (-(fi22)) {
                                        case -9223372036854775808:
                                            v11 -= v11
                                        default:
                                            v29 := GenS4{F0: 2147483647, F1: (v27 + (v27 / ((18446744073709551615) | 1))), F2: v9}
                                            _ = v29
                                        }
                                    default:
                                        v30 := GenS5{F0: (v4 * int(v2)), F1: 127, F2: fi22}
                                        _ = v30
                                    }
                                    for fi31 := 0; fi31 < 18; fi31++ {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v10 -= (v10 | (v10 << (((^(v27))) & 63)))
                                        _ = fi31
                                    }
                                }
                                switch v14 {
                                case "":
                                    var v32 uint32 = 4294967295
                                    _ = v32
                                case "hello":
                                    v1 = append(v1, 2147483647)
                                default:
                                    fmt.Println(v27, v5, v27)
                                }
                            } else {
                                var v33 bool = true
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    v35 := GenS6{F0: 9223372036854775807}
                                    v36 := GenS7{F0: (v27 * uint(v9))}
                                    if v33 {
                                        v11 = (v11 * v11)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        v37 := GenS5{F0: (v4 - fi34), F1: (v2 / ((127) | 1)), F2: (-(fi34))}
                                        _ = v37
                                    }
                                    fmt.Println(v4)
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v33
                            }
                            _ = fi22
                            _ = v27
                        }
                    }
                    _ = v19
                }
                _ = v14
            }
        default:
            v5 += v5
        }
    default:
        v10 = v10
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        if false {
            if (((-1 >= v5) || false) && true) {
                var v39 int = v4
                _ = v39
            } else {
                fmt.Println(v11)
                var v40 uint32 = 42
                if ((v10 % ((655) | 1)) != v10) {
                    if (209 == 65535) {
                        v41 := GenS8{F0: v5, F1: (0.0 == 1.5)}
                        if (-128 != -1) {
                            var v42 uint32 = (v40 & v40)
                            v2 = (v2 - v2)
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 5; fi43++ {
                            var v44 int32 = -1
                            v45 := GenS7{F0: 0}
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        _ = v41
                    } else {
                        v46 := GenS9{F0: true}
                        if true {
                            v3 = append(v3, -1)
                            switch "hello" {
                            case "a":
                                v11 = 151
                            case "foo bar":
                                var v47 uint32 = (^(v40))
                                _ = v47
                            case "\n\t":
                                var v48 int64 = 1
                                _ = v48
                            }
                            if (!true) {
                                if (((0.1 + (1.5 / 1.0)) == 0.0) && (1.5 <= (-2.25 * 3.14159))) {
                                    for fi49 := 0; fi49 < 7; fi49++ {
                                        v9 = v9
                                        _ = fi49
                                    }
                                    v11 = v9
                                    v4 = fi38
                                } else {
                                    v5 = v5
                                    v40 -= 558
                                }
                            } else {
                                var v50 string = ""
                                _ = v50
                            }
                        } else {
                            if ((v5 & (v5 | v5)) >= (v5 / ((133) | 1))) {
                                v2 = v2
                                v9 += (v9 * (v11 - (^(v9))))
                            }
                            for fi51 := 0; fi51 < 5; fi51++ {
                                if (false || ((!false) || false)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v40 = (^(v40))
                                }
                                if false {
                                    v9 = (v11 + (v11 - (v9 - (v11 | (v11 + v9)))))
                                }
                                _ = fi51
                            }
                        }
                        if true {
                            switch -1 {
                            case 0:
                                v52 := GenS10{F0: 702}
                                _ = v52
                            case -891:
                                switch (("hello" + "foo bar") + "") {
                                case "hello":
                                    v53 := GenS9{F0: (2 > v40)}
                                    _ = v53
                                case "\n\t":
                                    for fi54 := 0; fi54 < 16; fi54++ {
                                        fmt.Println(v2)
                                        v4 += 1
                                        v55 := GenS11{F0: v40, F1: 111, F2: "foo bar"}
                                        _ = fi54
                                        _ = v55
                                    }
                                default:
                                    v5 = (v5 * v5)
                                }
                            case 35:
                                v56 := GenS4{F0: -2147483648, F1: 42, F2: (^(v11))}
                                _ = v56
                            }
                        }
                        _ = v46
                    }
                }
                _ = v40
            }
        } else {
            if true {
                v57 := GenS12{F0: (v11 - v9), F1: 0, F2: 9223372036854775807}
                _ = v57
            }
            fmt.Println(v10, v10, v5)
            for fi58 := 0; fi58 < 12; fi58++ {
                switch "" {
                case "":
                    switch 4294967295 {
                    case 242:
                        var v59 string = (((("" + "\n\t") + "foo bar") + ("\n\t" + "foo bar")) + "\n\t")
                        _ = v59
                    case 42:
                        switch fi58 {
                        case 4294967295:
                            v11 += v11
                        case 1:
                            v60 := GenS13{F0: 2147483647, F1: (fi38 / (((v4 % (((fi58 / ((-41) | 1))) | 1))) | 1))}
                            _ = v60
                        default:
                            v11 += 149
                        }
                    case -9223372036854775808:
                        fmt.Println(v4)
                    }
                default:
                    fmt.Println(fi38, v9)
                }
                _ = fi58
            }
        }
        _ = fi38
    }
    if (v10 >= int64(v4)) {
        if (uint32(v2) == 2) {
            if (true && true) {
                v61 := GenS8{F0: 32767, F1: ((128 != 0) && ((v5 + -877) > (v5 >> ((42) & 63))))}
                for fi62 := 0; fi62 < 6; fi62++ {
                    v10 = (-(v10))
                    fmt.Println(v5, v10)
                    v2 -= v2
                    _ = fi62
                }
                if (18446744073709551615 >= v9) {
                    switch (("a" + "hello") + (("\n\t" + "") + "foo bar")) {
                    case "a":
                        var v63 int8 = 0
                        _ = v63
                    case "hello":
                        v1 = append(v1, 602)
                    case "foo bar":
                        var v64 bool = false
                        _ = v64
                    default:
                        v4 += (v4 * (v4 >> ((0) & 63)))
                    }
                    if false {
                        v65 := GenS14{F0: 42, F1: v5, F2: 100}
                        switch v4 {
                        case 128:
                            if true {
                                switch "foo bar" {
                                case "\n\t":
                                    var v66 uint64 = (v9 % (((^(v9))) | 1))
                                    _ = v66
                                case "":
                                    if true {
                                        for fi67 := 0; fi67 < 17; fi67++ {
                                            if ((!true) && true) {
                                                v68 := GenS6{F0: 225}
                                                for fi69 := 0; fi69 < 15; fi69++ {
                                                    v11 += v9
                                                    fmt.Println(v10, fi69, v2)
                                                    fmt.Println(v5, fi69)
                                                    _ = fi69
                                                }
                                                _ = v68
                                            } else {
                                                if (4294967295 < 2147483647) {
                                                    switch "" {
                                                    case "a":
                                                        v9 -= 923
                                                    case "hello":
                                                        v5 += (v5 - v5)
                                                    case "":
                                                        if (true && true) {
                                                            var v70 int64 = (v10 + v10)
                                                            _ = v70
                                                        } else {
                                                            switch ("\n\t" + "foo bar") {
                                                            case "hello":
                                                                var v71 string = "a"
                                                                _ = v71
                                                            case "":
                                                                for fi72 := 0; fi72 < 9; fi72++ {
                                                                    v5 += 177
                                                                    _ = fi72
                                                                }
                                                            }
                                                            switch "foo bar" {
                                                            case "hello":
                                                                for fi73 := 0; fi73 < 4; fi73++ {
                                                                    v2 = (v2 << ((645) & 63))
                                                                    switch (fi67 & (v4 & (^(v4)))) {
                                                                    case -2147483648:
                                                                        v74 := GenS7{F0: 255}
                                                                        _ = v74
                                                                    case 128:
                                                                        var v75 uint64 = 42
                                                                        _ = v75
                                                                    case -957:
                                                                        v5 = v5
                                                                    }
                                                                    if (-30 <= 42) {
                                                                        fmt.Println(fi73, fi67, fi73)
                                                                        fmt.Println(v9, fi67, v9)
                                                                        v76 := GenS4{F0: -1, F1: 4294967295, F2: (v11 << ((1) & 63))}
                                                                        v77 := GenS15{F0: uint64(v2), F1: -128}
                                                                        _ = v76
                                                                        _ = v77
                                                                    } else {
                                                                        var v78 int64 = (-(v10))
                                                                        _ = v78
                                                                    }
                                                                    _ = fi73
                                                                }
                                                            case "foo bar":
                                                                var v79 float64 = 0.1
                                                                _ = v79
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            v11 -= v11
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = fi67
                                        }
                                    }
                                case "hello":
                                    for fi80 := 0; fi80 < 13; fi80++ {
                                        var v81 uint8 = 15
                                        for fi82 := 0; fi82 < 16; fi82++ {
                                            fmt.Println(v2, v4, fi80)
                                            var v83 float64 = (0.0 / 1.0)
                                            if len(v1) > 0 {
                                                v1[0] = 929
                                            }
                                            var v84 int = (fi80 ^ (v4 ^ (fi80 - (v4 & -366))))
                                            var v85 int64 = 0
                                            _ = len(v1)
                                            if (!true) {
                                                var v86 int = 127
                                                if (!false) {
                                                    switch ((("a" + ("a" + "\n\t")) + "") + "\n\t") {
                                                    case "foo bar":
                                                        if (2147483647 == 1) {
                                                            fmt.Println(fi80)
                                                            v5 -= (^(v5))
                                                        } else {
                                                            if true {
                                                                v87 := GenS16{F0: 2, F1: v9, F2: 1}
                                                                if false {
                                                                    v85 -= (v85 ^ v85)
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                                var v88 uint32 = 631
                                                                _ = v87
                                                                _ = v88
                                                            } else {
                                                                v3 = append(v3, -512)
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v9 -= (^(v9))
                                                                case "":
                                                                    var v89 string = ("hello" + "\n\t")
                                                                    _ = v89
                                                                }
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = 318
                                                            }
                                                            if (false || false) {
                                                                if ((v5 & v5) < (v5 ^ 1)) {
                                                                    v2 = (^(v2))
                                                                    if false {
                                                                        switch 425 {
                                                                        case 4294967295:
                                                                            v2 = v2
                                                                        case 32768:
                                                                            if (!false) {
                                                                                v90 := GenS2{F0: 42, F1: -22}
                                                                                _ = v90
                                                                            }
                                                                        case 0:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v86
                                                                            }
                                                                        }
                                                                        v5 += (v5 | v5)
                                                                        var v91 bool = (42 <= -9223372036854775808)
                                                                        switch fi80 {
                                                                        case 388:
                                                                            fmt.Println(fi82)
                                                                        case 128:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v84
                                                                            }
                                                                        case 890:
                                                                            var v92 uint16 = 2
                                                                            _ = v92
                                                                        }
                                                                        _ = v91
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = int(v85)
                                                                    }
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (v4 << ((18446744073709551615) & 63))
                                                                    }
                                                                    if (v9 <= (v9 * v11)) {
                                                                        var v93 uint32 = 65535
                                                                        _ = v93
                                                                    }
                                                                }
                                                                v4 -= 42
                                                            }
                                                        }
                                                    case "a":
                                                        if false {
                                                            if false {
                                                                v2 = 128
                                                                _ = len(v3)
                                                                v5 = 1
                                                            }
                                                        } else {
                                                            switch "a" {
                                                            case "hello":
                                                                var v94 bool = true
                                                                _ = v94
                                                            }
                                                            var v95 uint16 = 496
                                                            if (((-2.25 + (v83 / v83)) * (0.0 * float64(3.14159))) > (float64((0.1 * -2.25)) - 1.5)) {
                                                                var v96 int = 457
                                                                _ = v96
                                                            }
                                                            _ = v95
                                                        }
                                                    case "hello":
                                                        v97 := GenS5{F0: (fi82 * (^(v86))), F1: uint8(v9), F2: (v86 | v84)}
                                                        _ = v97
                                                    }
                                                    var v98 uint64 = v9
                                                    _ = v98
                                                } else {
                                                    if (v83 > v83) {
                                                        switch (("hello" + "hello") + "hello") {
                                                        case "":
                                                            fmt.Println(v10, v81, fi82)
                                                        case "a":
                                                            var v99 string = ("" + "")
                                                            _ = v99
                                                        case "\n\t":
                                                            fmt.Println(v9, v83)
                                                        }
                                                    } else {
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v100 := GenS16{F0: 0, F1: 4294967295, F2: 100}
                                                            v85 -= -1
                                                            _ = v100
                                                        }
                                                        v5 += (v5 + 32767)
                                                        v101 := GenS11{F0: 4294967295, F1: 0, F2: (("hello" + "hello") + "")}
                                                        _ = v101
                                                    }
                                                }
                                                if (true || false) {
                                                    if (980 != 65535) {
                                                        v102 := GenS17{F0: fi80}
                                                        v2 -= v81
                                                        v81 = v2
                                                        _ = v102
                                                    }
                                                } else {
                                                    v103 := GenS3{F0: "\n\t"}
                                                    fmt.Println(v86, v84)
                                                    _ = v103
                                                }
                                                if (!((v5 & (v5 | (v5 ^ -1))) <= 432)) {
                                                    v104 := GenS6{F0: v85}
                                                    switch ("a" + (("" + ((("" + "foo bar") + "") + "hello")) + "a")) {
                                                    case "\n\t":
                                                        fmt.Println(v85, v85, v2)
                                                    case "hello":
                                                        v105 := GenS18{F0: "hello", F1: -1}
                                                        _ = v105
                                                    default:
                                                        if true {
                                                            switch ("a" + "\n\t") {
                                                            case "a":
                                                                var v106 float64 = ((float64(3.14159) + ((v83 - v83) * ((float64(float32((((v83 - 1.5) + v83) + -2.25))) / 0.1) * 3.14159))) / v83)
                                                                _ = v106
                                                            case "hello":
                                                                var v107 string = "a"
                                                                if len(v3) > 0 {
                                                                    v3[0] = (^(v84))
                                                                }
                                                                if true {
                                                                    switch fi82 {
                                                                    case -1:
                                                                        if true {
                                                                            v2 = v81
                                                                            v108 := GenS2{F0: 2, F1: 42}
                                                                            v10 = v10
                                                                            if (true && (false && (32767 == (v4 >> ((255) & 63))))) {
                                                                                switch 9223372036854775807 {
                                                                                case 630:
                                                                                    var v109 uint = 65535
                                                                                    _ = v109
                                                                                case 731:
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = -9223372036854775808
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v110 uint8 = 1
                                                                                v86 += fi82
                                                                                if ((v5 - 1000) > (v5 + (-(v5)))) {
                                                                                    if true {
                                                                                        var v111 int32 = -495
                                                                                        _ = v111
                                                                                    } else {
                                                                                        if false {
                                                                                            v112 := GenS13{F0: 2147483648, F1: (fi82 | fi80)}
                                                                                            v113 := GenS2{F0: 2147483648, F1: 100}
                                                                                            v1 = append(v1, 422)
                                                                                            _ = v112
                                                                                            _ = v113
                                                                                        } else {
                                                                                            v84 = (-(fi80))
                                                                                        }
                                                                                        var v114 float32 = ((1e10 + ((1.5 / -2.25) / 0.1)) + ((-2.25 - 1.5) - 1e10))
                                                                                        v9 = (v11 * (v9 << ((740) & 63)))
                                                                                        _ = v114
                                                                                    }
                                                                                    v115 := GenS19{F0: (v107 + v107), F1: v9}
                                                                                    _ = len(v1)
                                                                                    _ = v115
                                                                                }
                                                                                _ = v110
                                                                            }
                                                                            if true {
                                                                                if (0.1 != 3.14159) {
                                                                                    v116 := GenS11{F0: 0, F1: 0, F2: (v107 + v107)}
                                                                                    v107 = "a"
                                                                                    v4 = 1
                                                                                    _ = v116
                                                                                } else {
                                                                                    if (1.5 >= (1.5 * -2.25)) {
                                                                                        v86 = (v84 << ((65535) & 63))
                                                                                        v117 := GenS12{F0: v9, F1: 0, F2: (v85 % ((466) | 1))}
                                                                                        v5 = 1
                                                                                        _ = v117
                                                                                    } else {
                                                                                        v118 := GenS20{F0: 871, F1: ((v107 + ("a" + v107)) + (v107 + (((v107 + ("a" + (v107 + v107))) + v107) + v107)))}
                                                                                        var v119 string = (v107 + v107)
                                                                                        var v120 uint64 = 179
                                                                                        var v121 float64 = ((((v83 / v83) * v83) - float64(0.1)) + v83)
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        _ = v118
                                                                                        _ = v119
                                                                                        _ = v120
                                                                                        _ = v121
                                                                                    }
                                                                                    if false {
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            v83 += (-2.25 - -2.25)
                                                                                        case "a":
                                                                                            v122 := GenS16{F0: 103, F1: 1, F2: 0}
                                                                                            _ = v122
                                                                                        default:
                                                                                            v107 = v107
                                                                                        }
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        v81 -= 42
                                                                                    }
                                                                                }
                                                                                var v123 int8 = 0
                                                                                v123 += (v123 / ((v123) | 1))
                                                                                _ = v123
                                                                            }
                                                                            if false {
                                                                                v3 = append(v3, v86)
                                                                                var v124 uint = 255
                                                                                _ = v124
                                                                            } else {
                                                                                switch ((v107 + v107) + v107) {
                                                                                case "hello":
                                                                                    v2 -= (v81 - v81)
                                                                                    fmt.Println(v9)
                                                                                    v86 = v84
                                                                                    var v125 bool = (42 <= 942)
                                                                                    _ = v125
                                                                                }
                                                                                if true {
                                                                                    v10 = v85
                                                                                }
                                                                                _ = len(v1)
                                                                            }
                                                                            var v126 float64 = 1e10
                                                                            _ = v108
                                                                            _ = v126
                                                                        } else {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v127 uint8 = 98
                                                                            if false {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = -95
                                                                                }
                                                                                switch -9223372036854775808 {
                                                                                case -419:
                                                                                    v86 = fi80
                                                                                default:
                                                                                    v107 = "foo bar"
                                                                                    v128 := GenS21{F0: (^(v5)), F1: 989}
                                                                                    fmt.Println(v4, v9)
                                                                                    _ = v128
                                                                                }
                                                                                v107 = "\n\t"
                                                                            } else {
                                                                                switch ("\n\t" + v107) {
                                                                                case "hello":
                                                                                    var v129 int32 = 544
                                                                                    _ = v129
                                                                                }
                                                                                v11 += uint64(v2)
                                                                                v130 := GenS22{F0: 74, F1: 76}
                                                                                _ = v130
                                                                            }
                                                                            _ = v127
                                                                        }
                                                                    case 4294967295:
                                                                        v131 := GenS23{F0: -2147483648, F1: 2147483647, F2: 291}
                                                                        _ = v131
                                                                    case -479:
                                                                        if (false || true) {
                                                                            v86 = fi80
                                                                            var v132 uint64 = v11
                                                                            fmt.Println(v84)
                                                                            _ = v132
                                                                        }
                                                                    default:
                                                                        var v133 float64 = (1e10 * v83)
                                                                        _ = v133
                                                                    }
                                                                    var v134 uint16 = 0
                                                                    var v135 string = ((v107 + "hello") + v107)
                                                                    fmt.Println(fi80, v11, v107)
                                                                    v136 := GenS24{F0: int8(v5)}
                                                                    _ = v134
                                                                    _ = v135
                                                                    _ = v136
                                                                } else {
                                                                    switch ("foo bar" + v107) {
                                                                    case "":
                                                                        var v137 uint = 4294967295
                                                                        _ = v137
                                                                    case "foo bar":
                                                                        var v138 int64 = v10
                                                                        _ = v138
                                                                    case "hello":
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                }
                                                                var v139 int64 = v85
                                                                if true {
                                                                    v10 -= v139
                                                                    v107 = v107
                                                                }
                                                                _ = v107
                                                                _ = v139
                                                            case "\n\t":
                                                                v10 += (v85 * v85)
                                                            default:
                                                                if (1 != -16) {
                                                                    var v140 uint64 = v9
                                                                    if true {
                                                                        fmt.Println(v81, v86, v140)
                                                                    }
                                                                    var v141 uint8 = v2
                                                                    _ = v140
                                                                    _ = v141
                                                                } else {
                                                                    v142 := GenS9{F0: true}
                                                                    v11 -= (v9 % ((v11) | 1))
                                                                    _ = v142
                                                                }
                                                            }
                                                            v81 += 2
                                                            var v143 int = 4294967295
                                                            if ((v9 < v11) || ((fi80 | v84) > 32768)) {
                                                                v144 := GenS25{F0: "a", F1: (!(false || true)), F2: v5}
                                                                v1 = append(v1, 2147483648)
                                                                _ = v144
                                                            }
                                                            v2 = v81
                                                            _ = v143
                                                        } else {
                                                            v85 -= (v85 - v10)
                                                            _ = len(v3)
                                                        }
                                                    }
                                                    var v145 uint = 1
                                                    _ = v104
                                                    _ = v145
                                                }
                                                _ = v86
                                            } else {
                                                v10 += v85
                                                v10 -= v85
                                            }
                                            _ = fi82
                                            _ = v83
                                            _ = v84
                                            _ = v85
                                        }
                                        _ = fi80
                                        _ = v81
                                    }
                                default:
                                    var v146 uint64 = (v11 + 4294967295)
                                    _ = v146
                                }
                                if true {
                                    v4 = (v4 << ((18446744073709551615) & 63))
                                    for fi147 := 0; fi147 < 4; fi147++ {
                                        v5 -= 0
                                        for fi148 := 0; fi148 < 6; fi148++ {
                                            v149 := GenS21{F0: v5, F1: 2}
                                            v11 += (v11 / (((v9 ^ (^(v11)))) | 1))
                                            _ = fi148
                                            _ = v149
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = 42
                                        }
                                        _ = fi147
                                    }
                                }
                            } else {
                                v150 := GenS12{F0: v11, F1: int32(v2), F2: -1}
                                v5 -= v5
                                _ = v150
                            }
                        case -962:
                            var v151 float64 = (3.14159 / 1.0)
                            _ = v151
                        case 533:
                            v9 = (v9 + v11)
                        default:
                            switch "hello" {
                            case "\n\t":
                                var v152 uint64 = (v9 - (v9 % (((^(v11))) | 1)))
                                _ = v152
                            case "foo bar":
                                switch v4 {
                                case 128:
                                    var v153 string = (("a" + "\n\t") + "hello")
                                    _ = v153
                                case 1000:
                                    v154 := GenS26{F0: 35}
                                    _ = v154
                                case -2147483648:
                                    for fi155 := 0; fi155 < 15; fi155++ {
                                        if len(v3) > 0 {
                                            v3[0] = v4
                                        }
                                        switch ((("" + "a") + ("foo bar" + "a")) + "a") {
                                        case "":
                                            fmt.Println(v5)
                                        case "hello":
                                            v156 := GenS9{F0: false}
                                            _ = v156
                                        default:
                                            if ((^(fi155)) <= (v4 / (((^(v4))) | 1))) {
                                                v2 += (v2 / (((v2 / ((42) | 1))) | 1))
                                            }
                                        }
                                        v2 += v2
                                        _ = fi155
                                    }
                                }
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        }
                        _ = v65
                    } else {
                        var v157 uint = 65535
                        if (v4 < (v4 - 42)) {
                            fmt.Println(v11, v4)
                        }
                        switch ("foo bar" + "\n\t") {
                        case "foo bar":
                            var v158 int32 = 0
                            _ = v158
                        case "\n\t":
                            v159 := GenS27{F0: -1}
                            _ = v159
                        case "hello":
                            v4 = (^(v4))
                        }
                        _ = v157
                    }
                    v5 = -1
                }
                _ = v61
            } else {
                var v160 uint64 = (v9 - v11)
                v10 -= (v10 / (((v10 << ((0) & 63))) | 1))
                _ = v160
            }
        } else {
            var v161 int32 = -929
            _ = v161
        }
        if false {
            if len(v1) > 0 {
                v1[0] = 0
            }
            switch 42 {
            case 4294967295:
                switch ("a" + ("hello" + "hello")) {
                case "hello":
                    v162 := GenS15{F0: v11, F1: -1}
                    _ = v162
                case "foo bar":
                    var v163 uint = 623
                    _ = v163
                case "a":
                    var v164 int8 = -1
                    _ = v164
                }
            }
            var v165 uint = 0
            _ = v165
        } else {
            fmt.Println(v5, v9)
            _ = len(v1)
            v166 := GenS27{F0: (-(v10))}
            _ = v166
        }
        var v167 uint64 = 1
        _ = v167
    } else {
        v11 -= 42
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (-1 >= int16(v4)) {
        v10 += (^(v10))
        v2 -= 138
        v9 += 2
    } else {
        switch "hello" {
        case "hello":
            if true {
                v168 := GenS6{F0: -9223372036854775808}
                _ = v168
            }
        default:
            var v169 int32 = -1
            _ = v169
        }
    }
    v4 = v4
    var v170 string = "\n\t"
    v5 = (v5 - 32767)
    if true {
        v171 := GenS27{F0: (v10 ^ (^(v10)))}
        v170 = (((v170 + v170) + "") + v170)
        _ = v171
    }
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v170:", v170)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v170
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
