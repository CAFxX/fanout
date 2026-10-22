package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 int16; F2 bool }
type GenS8 struct { F0 bool; F1 uint }

func main() {
    var v1 float32 = (0.0 - 0.1)
    v2 := []int{}
    v3 := make(map[string]bool)
    if (true && (-32768 <= 1)) {
        switch "" {
        case "hello":
            switch 32767 {
            case 32768:
                fmt.Println(v1)
            case 0:
                for fi4 := 0; fi4 < 1; fi4++ {
                    for fi5 := 0; fi5 < 5; fi5++ {
                        switch (fi4 ^ fi4) {
                        case -501:
                            if len(v2) > 0 {
                                v2[0] = fi5
                            }
                        case 401:
                            var v6 string = "foo bar"
                            _ = v6
                        case -739:
                            if true {
                                v7 := GenS1{F0: 255, F1: (2147483647 < 0), F2: 2}
                                _ = v7
                            } else {
                                fmt.Println(v1, fi4, v1)
                                v1 += 1e10
                                fmt.Println(fi4, fi5, fi4)
                            }
                        }
                        _ = fi5
                    }
                    _ = fi4
                }
            case 9223372036854775807:
                switch 42 {
                case 42:
                    _ = len(v3)
                case 1:
                    var v8 uint16 = 2
                    _ = v8
                case 127:
                    if true {
                        var v9 int8 = 0
                        _ = v9
                    }
                }
            default:
                v10 := v3[((("\n\t" + (("hello" + "\n\t") + ("\n\t" + ("foo bar" + ("hello" + "foo bar"))))) + "") + "\n\t")]
                _ = v10
            }
        case "\n\t":
            var v11 int8 = -10
            _ = v11
        default:
            if true {
                for fi12 := 0; fi12 < 15; fi12++ {
                    if false {
                        v1 += float32(0.0)
                    }
                    _ = fi12
                }
                v1 -= v1
                v1 += ((3.14159 / v1) - 3.14159)
            } else {
                var v13 uint32 = 1
                var v14 uint64 = 0
                _ = v13
                _ = v14
            }
        }
    } else {
        v15 := GenS1{F0: 127, F1: false, F2: 4294967295}
        switch "hello" {
        case "":
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS1{F0: 128, F1: false, F2: 18446744073709551615}
                v1 += v1
                if (0 <= 406) {
                    var v18 float64 = (3.14159 / (3.14159 * float64(-2.25)))
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
        case "\n\t":
            var v19 uint = 760
            _ = v19
        case "foo bar":
            for fi20 := 0; fi20 < 13; fi20++ {
                fmt.Println(v1)
                var v21 uint16 = 65535
                v1 += 3.14159
                _ = fi20
                _ = v21
            }
        default:
            if false {
                if (false || (!true)) {
                    var v22 int16 = -1
                    v1 -= v1
                    switch "hello" {
                    case "foo bar":
                        v23 := GenS2{F0: 4294967295, F1: true}
                        _ = v23
                    case "\n\t":
                        if (4294967295 > 18446744073709551615) {
                            v1 = 3.14159
                            v1 -= 1e10
                            v22 += -1
                        } else {
                            switch 2147483647 {
                            case -414:
                                for fi24 := 0; fi24 < 1; fi24++ {
                                    if true {
                                        if (3.14159 >= v1) {
                                            var v25 uint8 = 127
                                            switch "hello" {
                                            case "a":
                                                for fi26 := 0; fi26 < 9; fi26++ {
                                                    fmt.Println(fi24, fi24)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v27 int8 = 42
                                                    v27 += (^(v27))
                                                    v27 = (v27 - (v27 + v27))
                                                    if true {
                                                        var v28 int = 9223372036854775807
                                                        v29 := GenS2{F0: 4294967295, F1: (1 < 32767)}
                                                        switch ("" + "foo bar") {
                                                        case "hello":
                                                            if len(v2) > 0 {
                                                                v2[0] = fi24
                                                            }
                                                            v30 := GenS2{F0: 112, F1: (!((v27 ^ 127) >= (v27 % (((v27 - v27)) | 1))))}
                                                            v28 += (fi26 % (((fi24 | v28)) | 1))
                                                            _ = v30
                                                        default:
                                                            if ((!true) || true) {
                                                                var v31 float32 = (v1 * v1)
                                                                _ = v31
                                                            }
                                                        }
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                    v22 += -1
                                                    _ = fi26
                                                    _ = v27
                                                }
                                            default:
                                                switch (("a" + "a") + ("\n\t" + "a")) {
                                                case "a":
                                                    var v32 float64 = 1.5
                                                    _ = v32
                                                case "hello":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case "foo bar":
                                                    v33 := GenS1{F0: (v25 ^ (v25 ^ v25)), F1: false, F2: 4294967295}
                                                    _ = v33
                                                default:
                                                    var v34 int64 = -9223372036854775808
                                                    _ = v34
                                                }
                                            }
                                            if false {
                                                var v35 uint16 = 2
                                                _ = v35
                                            } else {
                                                if true {
                                                    delete(v3, "hello")
                                                    fmt.Println(v1, v1)
                                                    if (245 >= -367) {
                                                        for fi36 := 0; fi36 < 14; fi36++ {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if ((-546 != -1) || true) {
                                                                var v37 uint64 = 4294967295
                                                                _ = v37
                                                            } else {
                                                                v1 = 0.1
                                                            }
                                                            if false {
                                                                v22 = (v22 ^ (v22 - v22))
                                                                v22 = (v22 | v22)
                                                                v38 := v3["foo bar"]
                                                                _ = v38
                                                            } else {
                                                                fmt.Println(v1, v1)
                                                            }
                                                            var v39 float32 = float32(-2.25)
                                                            v25 -= v25
                                                            _ = fi36
                                                            _ = v39
                                                        }
                                                        var v40 uint16 = 850
                                                        v40 = v40
                                                        _ = v40
                                                    } else {
                                                        v22 = v22
                                                        for fi41 := 0; fi41 < 1; fi41++ {
                                                            if (fi41 != (fi24 % ((fi41) | 1))) {
                                                                v42 := v3["hello"]
                                                                _ = v42
                                                            } else {
                                                                var v43 uint = 95
                                                                var v44 int32 = -1
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                            if (false || (561 > 2)) {
                                                                if (("a" + ("hello" + ("\n\t" + "hello"))) < ("a" + "foo bar")) {
                                                                    fmt.Println(v1)
                                                                }
                                                                v45 := GenS1{F0: (^(v25)), F1: false, F2: 18446744073709551615}
                                                                v22 += v22
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    if ((2 > 231) && (!true)) {
                                                                        v46 := GenS2{F0: 65535, F1: (1000 <= v22)}
                                                                        if (("foo bar" + "foo bar") > "") {
                                                                            if false {
                                                                                var v47 int = fi41
                                                                                _ = v47
                                                                            }
                                                                        }
                                                                        var v48 int64 = 42
                                                                        fmt.Println(v25, v1, v48)
                                                                        v25 += 160
                                                                        v3["\n\t"] = false
                                                                        v22 += -779
                                                                        v22 -= int16(v48)
                                                                        _ = v46
                                                                        _ = v48
                                                                    } else {
                                                                        v49 := GenS3{F0: 0, F1: (4294967295 >= 18446744073709551615)}
                                                                        var v50 uint16 = 216
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                case "hello":
                                                                    var v51 int8 = -128
                                                                    _ = v51
                                                                case "\n\t":
                                                                    if (!(-1 == -318)) {
                                                                        var v52 bool = false
                                                                        _ = v52
                                                                    } else {
                                                                        var v53 uint8 = v25
                                                                        _ = v53
                                                                    }
                                                                default:
                                                                    v22 -= -32768
                                                                }
                                                                v25 -= v25
                                                                _ = v45
                                                            }
                                                            var v54 int = fi24
                                                            _ = fi41
                                                            _ = v54
                                                        }
                                                        fmt.Println(v1, fi24, fi24)
                                                    }
                                                }
                                            }
                                            _ = v25
                                        }
                                    }
                                    v22 = v22
                                    switch "foo bar" {
                                    case "foo bar":
                                        if (!(!false)) {
                                            v55 := GenS4{F0: 2147483648}
                                            if ((false || false) || (int32(fi24) == int32(v22))) {
                                                var v56 uint8 = 24
                                                _ = v56
                                            } else {
                                                if false {
                                                    for fi57 := 0; fi57 < 13; fi57++ {
                                                        v22 = v22
                                                        if (0.0 >= 1.5) {
                                                            var v58 float64 = 3.14159
                                                            _ = v58
                                                        }
                                                        _ = fi57
                                                    }
                                                    v59 := GenS5{F0: 334}
                                                    _ = v59
                                                } else {
                                                    v60 := GenS5{F0: 361}
                                                    v2 = append(v2, 128)
                                                    switch fi24 {
                                                    case 119:
                                                        for fi61 := 0; fi61 < 7; fi61++ {
                                                            _ = len(v3)
                                                            _ = fi61
                                                        }
                                                    case 127:
                                                        v1 = (0.1 * ((v1 * (0.0 + v1)) + 1.5))
                                                    }
                                                    _ = v60
                                                }
                                            }
                                            _ = v55
                                        } else {
                                            fmt.Println(v22)
                                            v2 = append(v2, (fi24 / ((fi24) | 1)))
                                        }
                                    case "a":
                                        _ = len(v2)
                                    }
                                    _ = fi24
                                }
                            }
                            v62 := GenS6{F0: 74}
                            var v63 int8 = 114
                            _ = v62
                            _ = v63
                        }
                    case "hello":
                        var v64 uint8 = 128
                        _ = v64
                    default:
                        var v65 uint = 4294967295
                        _ = v65
                    }
                    _ = v22
                }
                _ = len(v3)
                var v66 float32 = v1
                _ = v66
            }
        }
        v1 = (v1 - ((v1 - v1) / v1))
        _ = v15
    }
    switch 4294967295 {
    case 4294967295:
        v1 -= (3.14159 - v1)
    default:
        v67 := GenS2{F0: 1, F1: true}
        _ = v67
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi68 := 0; fi68 < 1; fi68++ {
        switch (fi68 >> ((1) & 63)) {
        case 0:
            var v69 float64 = (((0.1 - 1.5) / 1.0) / 1.0)
            _ = v69
        }
        _ = fi68
    }
    v1 = 1.5
    v1 = 1.5
    fmt.Println(v1)
    var v70 uint16 = 125
    v1 = (float32(float64((v1 + -2.25))) + 0.1)
    delete(v3, "foo bar")
    for fi71 := 0; fi71 < 7; fi71++ {
        for fi72 := 0; fi72 < 17; fi72++ {
            var v73 int32 = 0
            if true {
                fmt.Println(fi71)
                v74 := GenS1{F0: 1, F1: false, F2: 42}
                switch "foo bar" {
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = 128
                    }
                default:
                    fmt.Println(v70, fi72)
                }
                _ = v74
            } else {
                switch ("" + "hello") {
                case "":
                    switch "" {
                    case "foo bar":
                        delete(v3, "\n\t")
                    case "a":
                        var v75 int8 = 42
                        _ = v75
                    case "\n\t":
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            switch "foo bar" {
                            case "a":
                                _ = len(v3)
                            case "hello":
                                v73 -= (-(v73))
                            }
                        default:
                            fmt.Println(v73, v70)
                        }
                        v1 += v1
                        if (!(42 > 1)) {
                            var v76 int8 = -1
                            _ = v76
                        }
                    default:
                        v1 = (3.14159 - v1)
                    }
                case "\n\t":
                    v77 := GenS6{F0: 42}
                    _ = v77
                }
            }
            _ = fi72
            _ = v73
        }
        _ = fi71
    }
    v70 -= (v70 / ((v70) | 1))
    for fi78 := 0; fi78 < 4; fi78++ {
        if len(v2) > 0 {
            v2[0] = (fi78 * (fi78 % ((fi78) | 1)))
        }
        if true {
            v1 -= (v1 * (((v1 / (v1 + float32(float64(((float32((1e10 * 1.5)) / v1) + 0.0))))) * -2.25) / (((1e10 + (v1 - (1e10 * 0.0))) / 1.5) - 1.5)))
            if true {
                for fi79 := 0; fi79 < 16; fi79++ {
                    switch fi79 {
                    case 9223372036854775807:
                        v3["a"] = false
                    case -683:
                        _ = len(v3)
                    case -1:
                        v1 += 1.5
                    default:
                        if (0 <= 42) {
                            if (42 == 1) {
                                v70 = 65535
                                v70 -= 466
                                delete(v3, ("\n\t" + "foo bar"))
                            } else {
                                v80 := GenS7{F0: 42, F1: 1000, F2: false}
                                switch (^(fi78)) {
                                case 1:
                                    v81 := GenS8{F0: true, F1: 0}
                                    _ = v81
                                default:
                                    v1 = (v1 * (v1 - v1))
                                    if (-9223372036854775808 < -832) {
                                        v1 -= (((((float32((1e10 + 1e10)) / 1e10) * 1e10) * float32(0.0)) / v1) * 0.0)
                                        switch "hello" {
                                        case "":
                                            var v82 int32 = -1
                                            _ = v82
                                        }
                                        switch (^(fi78)) {
                                        case -987:
                                            fmt.Println(v1, v1)
                                        case 127:
                                            var v83 uint8 = 255
                                            _ = v83
                                        default:
                                            if (!(-1 <= 127)) {
                                                switch "foo bar" {
                                                case "":
                                                    var v84 int8 = 26
                                                    _ = v84
                                                }
                                                v85 := GenS5{F0: (^(v70))}
                                                _ = v85
                                            }
                                        }
                                    }
                                }
                                _ = v80
                            }
                        } else {
                            fmt.Println(v1, v70)
                            var v86 uint = 2
                            _ = v86
                        }
                    }
                    fmt.Println(v70)
                    _ = fi79
                }
                for fi87 := 0; fi87 < 1; fi87++ {
                    var v88 float32 = (v1 / v1)
                    switch (fi78 & fi78) {
                    case 759:
                        fmt.Println(v70, v1, v1)
                    default:
                        if ("hello" <= "\n\t") {
                            _ = len(v3)
                        } else {
                            switch fi78 {
                            case 2147483647:
                                switch 0 {
                                case 4294967295:
                                    var v89 string = "\n\t"
                                    _ = v89
                                case 9223372036854775807:
                                    if len(v2) > 0 {
                                        v2[0] = fi87
                                    }
                                case 32768:
                                    v70 -= (^(v70))
                                }
                            case 32768:
                                var v90 string = "hello"
                                _ = v90
                            default:
                                v91 := GenS1{F0: 127, F1: true, F2: 18446744073709551615}
                                _ = v91
                            }
                            switch fi87 {
                            case 795:
                                v1 -= (v88 * v1)
                                var v92 string = "\n\t"
                                _ = v92
                            }
                            var v93 int64 = -157
                            _ = v93
                        }
                        fmt.Println(v70, v70, v88)
                    }
                    v70 = 255
                    _ = fi87
                    _ = v88
                }
                v1 = ((-2.25 - ((v1 - -2.25) + 1.5)) - v1)
            }
        } else {
            v3[("foo bar" + ("foo bar" + (("a" + "") + "foo bar")))] = ((42 <= 42) || (true || ("hello" < ("foo bar" + "\n\t"))))
            fmt.Println(v70)
            var v94 uint32 = 194
            _ = v94
        }
        _ = fi78
    }
    for fi95 := 0; fi95 < 18; fi95++ {
        v96 := GenS3{F0: v70, F1: false}
        if true {
            fmt.Println(v70, v70, v1)
            if len(v2) > 0 {
                v2[0] = fi95
            }
            _ = len(v2)
        }
        _ = fi95
        _ = v96
    }
    var v97 float64 = (1e10 / float64(-2.25))
    if false {
        var v98 float32 = v1
        delete(v3, (("a" + "foo bar") + "a"))
        for fi99 := 0; fi99 < 18; fi99++ {
            var v100 uint16 = (v70 >> ((42) & 63))
            _ = fi99
            _ = v100
        }
        _ = v98
    }
    v97 += (v97 - v97)
    v70 = (v70 - 2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v70:", v70)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v70
    _ = v97
}
