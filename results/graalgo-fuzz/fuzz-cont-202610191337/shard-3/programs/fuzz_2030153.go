package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int; F1 uint16 }
type GenS12 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS13 struct { F0 uint32; F1 string }
type GenS14 struct { F0 uint16; F1 uint; F2 int32 }
type GenS15 struct { F0 uint64; F1 uint }
type GenS16 struct { F0 int64; F1 uint32 }
type GenS17 struct { F0 int8; F1 int32; F2 int8 }
type GenS18 struct { F0 int32; F1 uint16 }
type GenS19 struct { F0 int64; F1 string; F2 int }
type GenS20 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS21 struct { F0 uint32; F1 string; F2 int32 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    v3 := make(map[uint32]int)
    v4 := []int{-9223372036854775808}
    v5 := GenS1{F0: 100, F1: 32767, F2: 2}
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int32 = -2147483648
        fmt.Println(v7, v2, v7)
        _ = fi6
        _ = v7
    }
    fmt.Println(v2)
    v2 = v2
    switch ("a" + "\n\t") {
    case "foo bar":
        var v8 float32 = (0.1 * (0.0 * ((1.5 * 1e10) * (((3.14159 * 0.0) * float32(((0.0 + float64((float32(v2) / -2.25))) - ((v2 * ((float64(3.14159) + 0.0) / float64(0.1))) - v2)))) + 1.5))))
        _ = v8
    case "\n\t":
        if (65535 <= 690) {
            var v9 float64 = (v2 * ((v2 - v2) + v2))
            _ = v9
        } else {
            var v10 float64 = 0.1
            for fi11 := 0; fi11 < 1; fi11++ {
                v2 -= (v2 / 3.14159)
                if true {
                    v12 := GenS2{F0: -90}
                    fmt.Println(fi11, v10, fi11)
                    _ = v12
                } else {
                    _ = len(v4)
                    v13 := GenS1{F0: -115, F1: 1000, F2: 2}
                    fmt.Println(v2)
                    _ = v13
                }
                switch 127 {
                case 9223372036854775807:
                    fmt.Println(v10, v2)
                }
                _ = fi11
            }
            _ = v10
        }
    case "":
        for fi14 := 0; fi14 < 18; fi14++ {
            _ = len(v1)
            _ = fi14
        }
    default:
        v2 = v2
    }
    if false {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int = 302
            switch (("" + ("hello" + "")) + "") {
            case "a":
                var v17 uint32 = 65535
                _ = v17
            case "\n\t":
                fmt.Println(fi15, v16, v16)
            case "foo bar":
                var v18 uint16 = 1
                _ = v18
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi15
            _ = v16
        }
        switch "" {
        case "hello":
            switch "" {
            case "\n\t":
                if true {
                    v2 = 0.1
                } else {
                    var v19 int16 = 989
                    fmt.Println(v2)
                    _ = v19
                }
            case "":
                for fi20 := 0; fi20 < 14; fi20++ {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if (2147483647 > 2147483647) {
                            fmt.Println(v2, fi21)
                            v1[2147483648] = false
                        } else {
                            fmt.Println(fi20, v2, fi21)
                            var v22 float32 = -2.25
                            _ = v22
                        }
                        v2 = (v2 / v2)
                        switch "hello" {
                        case "":
                            var v23 int8 = 42
                            _ = v23
                        case "a":
                            if (!false) {
                                fmt.Println(fi20)
                            }
                        case "hello":
                            if true {
                                if false {
                                    var v24 int16 = -568
                                    var v25 string = (((("" + "\n\t") + ("" + "a")) + "foo bar") + "foo bar")
                                    v4 = append(v4, fi20)
                                    var v26 string = ("hello" + v25)
                                    switch ("\n\t" + v25) {
                                    case "foo bar":
                                        if false {
                                            v4 = append(v4, (fi21 / ((fi21) | 1)))
                                        }
                                    case "":
                                        var v27 uint32 = 4294967295
                                        _ = v27
                                    case "hello":
                                        v25 = v26
                                    default:
                                        switch v26 {
                                        case "foo bar":
                                            switch fi20 {
                                            case -2147483648:
                                                v24 -= 32767
                                            case 9223372036854775807:
                                                delete(v1, 2147483648)
                                            case -1:
                                                var v28 float64 = ((0.0 - ((0.1 / v2) * ((v2 * v2) * v2))) - v2)
                                                _ = v28
                                            default:
                                                v24 -= (v24 ^ 42)
                                            }
                                        }
                                    }
                                    fmt.Println(fi21, v2, v24)
                                    _ = len(v4)
                                    fmt.Println(fi21)
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                } else {
                                    fmt.Println(v2, v2, v2)
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(fi20, fi20)
                                    default:
                                        var v29 int64 = 673
                                        _ = v29
                                    }
                                    if (-1 == 42) {
                                        v30 := GenS3{F0: 343, F1: fi21}
                                        _ = v30
                                    }
                                }
                                if (true || false) {
                                    v31 := GenS2{F0: 0}
                                    switch (("a" + (("a" + "foo bar") + "\n\t")) + "hello") {
                                    case "hello":
                                        fmt.Println(fi21, fi20)
                                    case "foo bar":
                                        if true {
                                            if false {
                                                v2 += 3.14159
                                            } else {
                                                fmt.Println(fi20, fi21)
                                                if (!true) {
                                                    switch fi20 {
                                                    case -9223372036854775808:
                                                        v32 := GenS1{F0: 1, F1: 1000, F2: 1}
                                                        _ = v32
                                                    case 2147483647:
                                                        if (42 <= 2) {
                                                            v2 = 0.1
                                                            var v33 int64 = -9223372036854775808
                                                            v33 += (v33 % (((-(v33))) | 1))
                                                            _ = v33
                                                        } else {
                                                            delete(v1, 1)
                                                        }
                                                    default:
                                                        fmt.Println(v2, fi21)
                                                    }
                                                    v2 += ((3.14159 / v2) + (((0.1 + 3.14159) + ((v2 * v2) + float64(-2.25))) - ((((((v2 / v2) + (1.5 * (float64((0.1 - float32((v2 - float64(1e10))))) * float64(((1e10 * 3.14159) + ((1e10 * (0.0 + (float32((-2.25 + 1e10)) / (1e10 - 1.5)))) + float32(v2))))))) + 0.1) - (v2 / 1.5)) - (1e10 - 1e10)) / v2)))
                                                    var v34 uint = 2
                                                    _ = v34
                                                } else {
                                                    fmt.Println(v2, fi21)
                                                }
                                                v35 := GenS2{F0: 2147483647}
                                                var v36 uint = 920
                                                _ = v35
                                                _ = v36
                                            }
                                        } else {
                                            var v37 bool = false
                                            v37 = (16 >= 20)
                                            _ = v37
                                        }
                                    default:
                                        v2 += ((v2 - 1e10) / 1.0)
                                    }
                                    _ = v31
                                } else {
                                    switch fi20 {
                                    case 42:
                                        if (false && true) {
                                            switch "\n\t" {
                                            case "a":
                                                var v38 uint8 = 255
                                                _ = v38
                                            case "hello":
                                                fmt.Println(fi21)
                                            case "foo bar":
                                                v2 += (((float64(1.5) / (v2 + 0.0)) / v2) - 0.0)
                                                var v39 uint = 255
                                                _ = v39
                                            }
                                            v2 += ((-2.25 * -2.25) - v2)
                                            fmt.Println(fi20, fi21, fi20)
                                        }
                                    case 192:
                                        var v40 string = "a"
                                        _ = v40
                                    case 128:
                                        fmt.Println(v2, fi20, v2)
                                    default:
                                        v2 = float64(3.14159)
                                    }
                                }
                            } else {
                                v2 = v2
                                v2 += (((v2 / v2) - 3.14159) + v2)
                            }
                        }
                        _ = fi21
                    }
                    _ = fi20
                }
            }
        case "\n\t":
            if (!(1 != 42)) {
                fmt.Println(v2)
                delete(v3, 92)
            }
        }
    }
    if true {
        for fi41 := 0; fi41 < 7; fi41++ {
            fmt.Println(v2)
            v2 -= 0.0
            if false {
                var v42 int16 = -32768
                _ = v42
            }
            _ = fi41
        }
        v43 := GenS2{F0: -1}
        _ = v43
    }
    v44 := GenS3{F0: -369, F1: -9223372036854775808}
    v45 := GenS1{F0: 0, F1: -421, F2: 4294967295}
    switch -21 {
    case -2147483648:
        for fi46 := 0; fi46 < 3; fi46++ {
            if len(v4) > 0 {
                v4[0] = (fi46 & (fi46 + fi46))
            }
            switch (fi46 % ((fi46) | 1)) {
            case 0:
                if false {
                    var v47 int = fi46
                    for fi48 := 0; fi48 < 17; fi48++ {
                        fmt.Println(fi46, v47, v2)
                        if true {
                            v2 = 1.5
                            v47 = 128
                        } else {
                            v47 = 305
                        }
                        if false {
                            var v49 uint16 = 334
                            v2 = (3.14159 + (v2 - v2))
                            _ = v49
                        }
                        if true {
                            if (false || ("\n\t" >= "")) {
                                var v50 float32 = 3.14159
                                v2 = 1.5
                                if true {
                                    var v51 uint32 = 2147483648
                                    switch (fi46 ^ v47) {
                                    case 994:
                                        v50 += v50
                                    case -9223372036854775808:
                                        v2 = v2
                                    }
                                    switch ("hello" + (("hello" + "foo bar") + "hello")) {
                                    case "\n\t":
                                        var v52 string = "\n\t"
                                        _ = v52
                                    }
                                    _ = v51
                                } else {
                                    v50 = v50
                                }
                                _ = v50
                            } else {
                                v53 := v1[2147483648]
                                var v54 float32 = 1.5
                                _ = v53
                                _ = v54
                            }
                        }
                        _ = len(v1)
                        fmt.Println(v47)
                        switch "a" {
                        case "\n\t":
                            if false {
                                v55 := GenS3{F0: 538, F1: v47}
                                switch (fi48 / (((fi46 << ((247) & 63))) | 1)) {
                                case 4294967295:
                                    v47 = fi48
                                case 1:
                                    if false {
                                        var v56 bool = false
                                        _ = v56
                                    }
                                    v47 = fi48
                                    v47 = (fi48 % ((v47) | 1))
                                default:
                                    fmt.Println(v47, fi48, fi46)
                                }
                                v2 = v2
                                _ = v55
                            } else {
                                v47 -= (fi48 >> ((18446744073709551615) & 63))
                                if true {
                                    v2 += (float64(0.1) * v2)
                                    switch ("" + "a") {
                                    case "":
                                        v47 = -905
                                    default:
                                        v47 += v47
                                    }
                                    switch (("foo bar" + ("\n\t" + ((((((("foo bar" + (("" + "") + (("hello" + ("foo bar" + "\n\t")) + ""))) + "\n\t") + "") + "hello") + "a") + "") + ""))) + ("" + "hello")) {
                                    case "foo bar":
                                        var v57 uint64 = 18446744073709551615
                                        _ = v57
                                    default:
                                        switch v47 {
                                        case 42:
                                            var v58 string = "foo bar"
                                            _ = v58
                                        case 0:
                                            v2 += float64(float32(v2))
                                        }
                                    }
                                } else {
                                    var v59 int = v47
                                    v60 := GenS2{F0: -1}
                                    _ = v59
                                    _ = v60
                                }
                            }
                        case "hello":
                            v61 := GenS4{F0: 42, F1: 189}
                            _ = v61
                        case "a":
                            if (1e10 < 3.14159) {
                                switch fi46 {
                                case -814:
                                    switch (^(fi46)) {
                                    case 2147483647:
                                        v62 := GenS2{F0: -2147483648}
                                        _ = v62
                                    case -336:
                                        fmt.Println(fi46, v47, fi48)
                                    default:
                                        var v63 uint = 111
                                        _ = v63
                                    }
                                case 1:
                                    switch 32767 {
                                    case -9223372036854775808:
                                        if (((2147483647 != 1) || ((false || false) && (true && (832 < 42)))) || false) {
                                            switch (-(v47)) {
                                            case 127:
                                                v2 += v2
                                            case 2147483647:
                                                if true {
                                                    v47 = 4294967295
                                                    var v64 int = fi48
                                                    if false {
                                                        switch v64 {
                                                        case -1:
                                                            if (!false) {
                                                                var v65 bool = true
                                                                _ = v65
                                                            }
                                                        default:
                                                            v66 := GenS5{F0: 426}
                                                            _ = v66
                                                        }
                                                        if true {
                                                            if false {
                                                                switch "" {
                                                                case "a":
                                                                    _ = len(v1)
                                                                case "":
                                                                    if false {
                                                                        v3[80] = v64
                                                                        v47 = 9223372036854775807
                                                                    } else {
                                                                        v64 += (-(v64))
                                                                    }
                                                                    _ = len(v4)
                                                                }
                                                                var v67 float32 = ((0.0 + 3.14159) / ((1e10 - -2.25) + (0.1 * 1e10)))
                                                                _ = v67
                                                            }
                                                            v2 += (v2 / ((v2 + (v2 / ((0.0 * 3.14159) + 1.5))) + (1e10 - (v2 * (float64(1.5) * ((1.5 / float64(1.5)) + (0.1 - ((float64(-2.25) * 1e10) * 1e10))))))))
                                                            v64 -= v47
                                                        } else {
                                                            if (9223372036854775807 <= 0) {
                                                                var v68 int32 = -279
                                                                _ = len(v3)
                                                                if ((!(-128 < -128)) && (128 <= 118)) {
                                                                    v64 = fi46
                                                                } else {
                                                                    v69 := v3[921]
                                                                    v70 := GenS1{F0: 105, F1: 878, F2: 0}
                                                                    if false {
                                                                        var v71 uint64 = 89
                                                                        _ = v71
                                                                    } else {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        v72 := GenS6{F0: v69, F1: true}
                                                                        _ = v72
                                                                    }
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                _ = v68
                                                            }
                                                            v73 := GenS3{F0: 1, F1: v64}
                                                            _ = len(v1)
                                                            var v74 uint32 = 0
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                        var v75 int = v47
                                                        _ = v75
                                                    }
                                                    _ = v64
                                                } else {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v76 int32 = 279
                                                        _ = v76
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi48 << ((1) & 63))
                                                    }
                                                    v77 := GenS7{F0: 1, F1: 1}
                                                    _ = v77
                                                }
                                            default:
                                                v2 = 1e10
                                            }
                                            v2 += 3.14159
                                            v2 = v2
                                            var v78 int8 = -26
                                            switch "foo bar" {
                                            case "":
                                                switch (v47 * (fi46 % ((-1) | 1))) {
                                                case -274:
                                                    v2 -= ((v2 / (v2 * (v2 / float64(((0.0 / (0.1 - 1e10)) + -2.25))))) - v2)
                                                case 4294967295:
                                                    switch "hello" {
                                                    case "":
                                                        v2 = (((v2 / -2.25) + v2) + v2)
                                                    case "foo bar":
                                                        v78 -= -123
                                                    case "hello":
                                                        v2 -= -2.25
                                                    }
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(v2)
                                                    case "foo bar":
                                                        v47 += (fi46 * fi48)
                                                    case "\n\t":
                                                        if true {
                                                            v47 -= 0
                                                            switch fi48 {
                                                            case -2147483648:
                                                                fmt.Println(v78, fi46)
                                                            }
                                                            if (3.14159 < 1e10) {
                                                                v47 += (fi46 + fi46)
                                                                var v79 uint16 = 255
                                                                v80 := GenS8{F0: 707, F1: (!(!false)), F2: 0}
                                                                var v81 float32 = (0.0 / (float32((((v2 - v2) / v2) / v2)) / ((3.14159 - 0.1) / float32((((v2 + ((v2 + 1.5) + v2)) * 0.0) * 0.1)))))
                                                                _ = len(v3)
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v82 float32 = -2.25
                                                                    _ = v82
                                                                }
                                                                if false {
                                                                    var v83 int32 = -2147483648
                                                                    _ = v83
                                                                } else {
                                                                    if false {
                                                                        var v84 int16 = 32767
                                                                        switch ("foo bar" + "foo bar") {
                                                                        case "foo bar":
                                                                            if false {
                                                                                if true {
                                                                                    if (851 == 65535) {
                                                                                        v85 := GenS8{F0: 18446744073709551615, F1: (!(1 < 889)), F2: 0}
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        v81 = float32(v2)
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            v81 += (1e10 * v81)
                                                                                        case "":
                                                                                            switch fi48 {
                                                                                            case 42:
                                                                                                fmt.Println(v78, v79, v78)
                                                                                            case 475:
                                                                                                v81 += v81
                                                                                            default:
                                                                                                v86 := GenS9{F0: v78, F1: false}
                                                                                                _ = v86
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v79 += v79
                                                                                        }
                                                                                        _ = v85
                                                                                    }
                                                                                    if true {
                                                                                        delete(v1, 329)
                                                                                    } else {
                                                                                        var v87 uint = 736
                                                                                        if false {
                                                                                            fmt.Println(v81)
                                                                                        }
                                                                                        _ = v87
                                                                                    }
                                                                                } else {
                                                                                    v88 := GenS1{F0: -1, F1: (v84 >> ((392) & 63)), F2: 207}
                                                                                    v79 = v79
                                                                                    v89 := GenS10{F0: (v79 / (((^(v79))) | 1))}
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                }
                                                                                v90 := GenS9{F0: v78, F1: false}
                                                                                v91 := GenS7{F0: -2147483648, F1: 186}
                                                                                _ = v90
                                                                                _ = v91
                                                                            } else {
                                                                                delete(v3, 589)
                                                                            }
                                                                        case "hello":
                                                                            if (!((v81 * (v81 / (1.5 + v81))) <= (0.0 / v81))) {
                                                                                var v92 uint32 = 2147483647
                                                                                if true {
                                                                                    var v93 int32 = 1
                                                                                    var v94 uint64 = 115
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                }
                                                                                delete(v3, (^(v92)))
                                                                                _ = v92
                                                                            }
                                                                        case "":
                                                                            delete(v3, 145)
                                                                        default:
                                                                            if (251 <= 2) {
                                                                                v84 = (v84 >> ((2) & 63))
                                                                                v47 += (^(fi46))
                                                                                var v95 string = "a"
                                                                                _ = v95
                                                                            }
                                                                        }
                                                                        if true {
                                                                            if (false || (0 < 42)) {
                                                                                var v96 int16 = 42
                                                                                var v97 int32 = 1
                                                                                _ = v96
                                                                                _ = v97
                                                                            }
                                                                            v47 -= fi48
                                                                            if (false || (true && (v79 >= (v79 + v79)))) {
                                                                                var v98 int = fi48
                                                                                _ = v98
                                                                            } else {
                                                                                var v99 uint = 2
                                                                                switch ("hello" + "foo bar") {
                                                                                case "":
                                                                                    v79 += 255
                                                                                case "hello":
                                                                                    v78 -= -123
                                                                                default:
                                                                                    var v100 uint64 = 579
                                                                                    fmt.Println(v2, v78)
                                                                                    _ = v100
                                                                                }
                                                                                _ = v99
                                                                            }
                                                                            var v101 uint8 = 2
                                                                            v101 = 255
                                                                            v81 -= v81
                                                                            _ = v101
                                                                        }
                                                                        if (65535 <= 42) {
                                                                            if false {
                                                                                switch (fi48 % ((fi48) | 1)) {
                                                                                case 42:
                                                                                    v102 := GenS11{F0: (fi46 * v47), F1: (v79 / ((v79) | 1))}
                                                                                    _ = v102
                                                                                }
                                                                            } else {
                                                                                fmt.Println(fi46, v79, v47)
                                                                                if (!false) {
                                                                                    v103 := GenS11{F0: v47, F1: 255}
                                                                                    _ = v103
                                                                                }
                                                                            }
                                                                            v3[811] = 32767
                                                                            v78 -= (v78 & 100)
                                                                            v104 := GenS12{F0: (false || true), F1: 2, F2: 1}
                                                                            v84 += (v84 & (v84 - (v84 << ((778) & 63))))
                                                                            if false {
                                                                                fmt.Println(v84, fi46, fi48)
                                                                                fmt.Println(v84, v79, v79)
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    var v105 int8 = (v78 >> ((99) & 63))
                                                                                    _ = v105
                                                                                case "":
                                                                                    v106 := GenS11{F0: 0, F1: (v79 % ((v79) | 1))}
                                                                                    _ = v106
                                                                                case "\n\t":
                                                                                    var v107 float32 = v81
                                                                                    _ = v107
                                                                                default:
                                                                                    v78 = v78
                                                                                }
                                                                            }
                                                                            _ = v104
                                                                        } else {
                                                                            v108 := GenS13{F0: 218, F1: (("\n\t" + ("a" + ("hello" + "\n\t"))) + (("foo bar" + ("\n\t" + "a")) + "foo bar"))}
                                                                            delete(v1, 2147483647)
                                                                            _ = v108
                                                                        }
                                                                        v109 := GenS14{F0: (v79 << ((1) & 63)), F1: 65535, F2: -763}
                                                                        _ = v109
                                                                        _ = v84
                                                                    } else {
                                                                        fmt.Println(v79, v81, v79)
                                                                    }
                                                                    v79 -= 255
                                                                    v110 := GenS15{F0: 0, F1: 683}
                                                                    _ = v110
                                                                }
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                            } else {
                                                                v111 := GenS4{F0: 2, F1: 212}
                                                                fmt.Println(fi48, fi46, fi46)
                                                                _ = v111
                                                            }
                                                        }
                                                        fmt.Println(fi48)
                                                    default:
                                                        fmt.Println(v47, fi46, fi48)
                                                    }
                                                    if true {
                                                        var v112 int = (fi48 % ((fi46) | 1))
                                                        var v113 int = v112
                                                        var v114 int16 = 0
                                                        v115 := GenS16{F0: -1, F1: 2147483648}
                                                        _ = len(v1)
                                                        var v116 bool = false
                                                        v117 := GenS17{F0: (v78 % (((v78 * v78)) | 1)), F1: -1, F2: (v78 - (-(v78)))}
                                                        _ = v112
                                                        _ = v113
                                                        _ = v114
                                                        _ = v115
                                                        _ = v116
                                                        _ = v117
                                                    }
                                                default:
                                                    v2 = float64(float32(0.1))
                                                    v2 += v2
                                                    switch ("a" + "\n\t") {
                                                    case "\n\t":
                                                        if (false || false) {
                                                            switch "hello" {
                                                            case "":
                                                                v78 = v78
                                                            case "a":
                                                                v118 := GenS18{F0: 0, F1: 255}
                                                                var v119 int32 = 2147483647
                                                                _ = v118
                                                                _ = v119
                                                            case "\n\t":
                                                                if (403 <= 533) {
                                                                    v120 := GenS18{F0: 2147483647, F1: 65535}
                                                                    var v121 uint32 = 2
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    _ = v120
                                                                    _ = v121
                                                                } else {
                                                                    fmt.Println(v2, v47, fi46)
                                                                    if true {
                                                                        if (2 >= 42) {
                                                                            v122 := GenS10{F0: 255}
                                                                            _ = v122
                                                                        }
                                                                        v47 = fi46
                                                                    }
                                                                }
                                                            }
                                                            if (false || (714 <= 42)) {
                                                                delete(v1, 1)
                                                                switch (-(fi48)) {
                                                                case 9223372036854775807:
                                                                    v123 := GenS6{F0: (fi48 * fi46), F1: false}
                                                                    _ = v123
                                                                }
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            } else {
                                                                fmt.Println(fi46)
                                                                var v124 int8 = 127
                                                                _ = v124
                                                            }
                                                        } else {
                                                            v125 := GenS19{F0: 67, F1: "", F2: v47}
                                                            if false {
                                                                var v126 int64 = -9223372036854775808
                                                                var v127 uint32 = 9
                                                                var v128 int = (fi48 * fi48)
                                                                _ = v126
                                                                _ = v127
                                                                _ = v128
                                                            } else {
                                                                v2 = (1e10 * v2)
                                                            }
                                                            v3[2147483647] = fi48
                                                            _ = v125
                                                        }
                                                    case "foo bar":
                                                        v78 = (^(v78))
                                                    }
                                                }
                                            case "foo bar":
                                                v47 = 149
                                            }
                                            _ = v78
                                        } else {
                                            v47 = -9223372036854775808
                                        }
                                    case -171:
                                        v129 := v1[953]
                                        _ = v129
                                    case 764:
                                        v130 := GenS9{F0: -1, F1: (-432 <= -851)}
                                        _ = v130
                                    default:
                                        var v131 int16 = -269
                                        if (((v131 / ((32767) | 1)) <= (v131 - (v131 % ((-32768) | 1)))) || true) {
                                            if len(v4) > 0 {
                                                v4[0] = fi48
                                            }
                                            switch ("\n\t" + "a") {
                                            case "":
                                                fmt.Println(v47, fi48)
                                            case "hello":
                                                switch ("foo bar" + "a") {
                                                case "a":
                                                    v2 -= (v2 - v2)
                                                case "\n\t":
                                                    var v132 int = fi46
                                                    _ = v132
                                                default:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            case "foo bar":
                                                v47 = (v47 << ((2) & 63))
                                            }
                                            _ = len(v4)
                                            if (-408 >= -9223372036854775808) {
                                                fmt.Println(v131, v47)
                                                var v133 int = fi46
                                                var v134 int8 = -128
                                                _ = v133
                                                _ = v134
                                            } else {
                                                v2 = 0.0
                                                v2 = (v2 - (3.14159 - 1e10))
                                            }
                                            v135 := GenS5{F0: 2}
                                            if true {
                                                switch (((((((("foo bar" + ("" + "a")) + "a") + "hello") + "foo bar") + "a") + "foo bar") + "hello") + "a") {
                                                case "hello":
                                                    if true {
                                                        fmt.Println(v47, v2)
                                                        var v136 int32 = -1
                                                        v1[972] = false
                                                        _ = v136
                                                    }
                                                default:
                                                    v137 := GenS18{F0: 42, F1: 0}
                                                    _ = v137
                                                }
                                                var v138 float32 = -2.25
                                                _ = v138
                                            } else {
                                                fmt.Println(fi46, fi48, fi48)
                                            }
                                            _ = v135
                                        }
                                        delete(v1, 2147483648)
                                        _ = v131
                                    }
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = -179
                                    }
                                }
                                _ = len(v4)
                                fmt.Println(v47, fi48, fi48)
                                v47 += fi46
                                if (125 > 1) {
                                    v139 := GenS20{F0: -1, F1: 42, F2: 0}
                                    v140 := GenS21{F0: 4294967295, F1: "", F2: 2147483647}
                                    _ = v139
                                    _ = v140
                                }
                            }
                        }
                        var v141 uint16 = 255
                        _ = fi48
                        _ = v141
                    }
                    _ = v47
                } else {
                    switch 0 {
                    case 42:
                        v142 := GenS14{F0: 779, F1: 422, F2: 2147483647}
                        _ = v142
                    case 32768:
                        v2 -= v2
                    case -2147483648:
                        if false {
                            switch ("" + ("foo bar" + "a")) {
                            case "foo bar":
                                switch "" {
                                case "":
                                    switch "" {
                                    case "":
                                        if false {
                                            for fi143 := 0; fi143 < 17; fi143++ {
                                                switch (-(fi143)) {
                                                case -9223372036854775808:
                                                    switch 128 {
                                                    case 127:
                                                        v2 -= v2
                                                    case -640:
                                                        v2 += (v2 + (v2 - v2))
                                                    default:
                                                        v2 += 0.0
                                                    }
                                                case 0:
                                                    if false {
                                                        fmt.Println(fi46, fi46, v2)
                                                        switch "" {
                                                        case "foo bar":
                                                            v1[624] = true
                                                        }
                                                    } else {
                                                        fmt.Println(fi143, v2)
                                                        switch fi143 {
                                                        case -277:
                                                            if true {
                                                                var v144 int16 = -1
                                                                var v145 int = fi46
                                                                _ = v144
                                                                _ = v145
                                                            }
                                                        default:
                                                            fmt.Println(v2)
                                                        }
                                                    }
                                                default:
                                                    if true {
                                                        var v146 uint16 = 0
                                                        _ = v146
                                                    } else {
                                                        if (!true) {
                                                            var v147 uint = 730
                                                            _ = v147
                                                        }
                                                        v148 := GenS10{F0: 749}
                                                        _ = v148
                                                    }
                                                }
                                                var v149 uint64 = 526
                                                v150 := GenS17{F0: 42, F1: 0, F2: -1}
                                                _ = fi143
                                                _ = v149
                                                _ = v150
                                            }
                                            fmt.Println(fi46, v2)
                                        } else {
                                            v2 = -2.25
                                            fmt.Println(v2, fi46)
                                            var v151 uint32 = 0
                                            _ = v151
                                        }
                                    default:
                                        v2 += (1e10 - (((0.0 * ((v2 / v2) * (1.5 / v2))) / (float64((0.0 * -2.25)) - float64(3.14159))) - (float64(float32(1e10)) / 1e10)))
                                    }
                                default:
                                    var v152 string = ""
                                    _ = v152
                                }
                            default:
                                v2 += v2
                            }
                            v153 := v3[2]
                            for fi154 := 0; fi154 < 12; fi154++ {
                                v153 += fi154
                                var v155 uint16 = 280
                                if true {
                                    var v156 int8 = 0
                                    _ = v156
                                }
                                _ = fi154
                                _ = v155
                            }
                            _ = v153
                        }
                    }
                }
            case -9223372036854775808:
                var v157 uint16 = 42
                _ = v157
            case -1:
                v158 := GenS22{F0: (1 < 474)}
                _ = v158
            default:
                for fi159 := 0; fi159 < 19; fi159++ {
                    fmt.Println(v2, fi46, fi159)
                    _ = fi159
                }
            }
            _ = fi46
        }
    case 0:
        var v160 int8 = -1
        _ = v160
    }
    for fi161 := 0; fi161 < 10; fi161++ {
        var v162 uint64 = 0
        if true {
            v162 += (^(v162))
            var v163 int32 = -1
            v164 := GenS23{F0: 0, F1: 0}
            _ = v163
            _ = v164
        }
        _ = fi161
        _ = v162
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
