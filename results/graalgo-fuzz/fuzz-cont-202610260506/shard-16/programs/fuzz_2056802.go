package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int }

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 128
    var v3 int16 = (v1 & v1)
    v4 := []int{127, -61, 1, 79}
    v4 = append(v4, 32767)
    var v5 int16 = v3
    fmt.Println(v1)
    switch "\n\t" {
    case "":
        v2 -= v2
    case "hello":
        var v6 int8 = 1
        _ = v6
    default:
        var v7 int8 = -128
        _ = v7
    }
    var v8 uint8 = (^(v2))
    _ = len(v4)
    if false {
        switch "\n\t" {
        case "foo bar":
            v2 = (v8 / ((255) | 1))
        case "hello":
            fmt.Println(v1)
        }
    } else {
        for fi9 := 0; fi9 < 13; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS1{F0: 0, F1: v3}
                switch "hello" {
                case "\n\t":
                    var v12 string = ("\n\t" + "foo bar")
                    _ = v12
                case "hello":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "":
                    fmt.Println(fi10)
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                fmt.Println(fi13)
                v15 := GenS2{F0: (v8 << ((4294967295) & 63)), F1: 42}
                _ = fi14
                _ = v15
            }
            fmt.Println(v1, v5)
            var v16 uint8 = v8
            _ = fi13
            _ = v16
        }
        var v17 float32 = (1.5 * ((0.0 - float32((((0.0 * 1.5) - 1.5) - (-2.25 / 1.5)))) * 0.1))
        _ = v17
    }
    if true {
        v4 = append(v4, 0)
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 float32 = 0.1
            _ = fi18
            _ = v19
        }
        v20 := GenS3{F0: -484}
        _ = v20
    }
    var v21 float64 = 1.5
    for fi22 := 0; fi22 < 5; fi22++ {
        if ("\n\t" == "a") {
            var v23 float64 = 1.5
            _ = v23
        } else {
            switch "\n\t" {
            case "foo bar":
                var v24 int64 = 705
                _ = v24
            case "":
                fmt.Println(v8, v8, v5)
            }
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        switch (fi25 * (^(fi25))) {
        case 128:
            if len(v4) > 0 {
                v4[0] = fi25
            }
        case -1:
            for fi26 := 0; fi26 < 13; fi26++ {
                if ("\n\t" <= "hello") {
                    switch "foo bar" {
                    case "":
                        fmt.Println(v5)
                    case "\n\t":
                        switch (fi25 - 2147483647) {
                        case 1:
                            var v27 uint = 0
                            if false {
                                fmt.Println(v27, v5, v3)
                                var v28 uint32 = 42
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v28
                            }
                            v3 = v3
                            _ = v27
                        case 127:
                            v29 := GenS3{F0: fi26}
                            _ = v29
                        case 32767:
                            fmt.Println(fi26, v21, fi26)
                        default:
                            v4 = append(v4, (fi25 + fi26))
                        }
                    case "a":
                        v30 := GenS4{F0: 1}
                        _ = v30
                    default:
                        v2 += v8
                    }
                    var v31 float32 = 0.1
                    if true {
                        v3 += (^(v5))
                        fmt.Println(v2, v8)
                    } else {
                        switch ((("hello" + "a") + ("" + ((("a" + ("hello" + "")) + ("hello" + ("foo bar" + "\n\t"))) + "\n\t"))) + "\n\t") {
                        case "\n\t":
                            switch (fi25 << ((802) & 63)) {
                            case 32768:
                                v3 = v3
                            case -2147483648:
                                if ((!true) || true) {
                                    var v32 string = "a"
                                    fmt.Println(v5)
                                    _ = v32
                                } else {
                                    v33 := GenS4{F0: 2}
                                    var v34 int = (fi26 / (((fi25 / (((fi26 - 32767)) | 1))) | 1))
                                    v21 = (v21 * float64(v31))
                                    _ = v33
                                    _ = v34
                                }
                            case -394:
                                v8 -= 128
                            default:
                                if (1 == 49) {
                                    var v35 string = "hello"
                                    v21 += ((0.0 / 1.0) / (((v21 / (-2.25 / (v21 + v21))) - v21) + v21))
                                    var v36 uint8 = (v2 >> ((84) & 63))
                                    _ = v35
                                    _ = v36
                                } else {
                                    fmt.Println(fi25)
                                    v37 := GenS4{F0: 255}
                                    _ = v37
                                }
                            }
                        default:
                            fmt.Println(v31, fi26, v31)
                        }
                        var v38 uint16 = 533
                        _ = v38
                    }
                    _ = v31
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch fi26 {
                    case 4294967295:
                        var v39 int64 = 9223372036854775807
                        _ = v39
                    case 128:
                        if (42 > 2147483647) {
                            fmt.Println(v8, v2)
                            if (255 != 0) {
                                if false {
                                    v40 := GenS3{F0: fi26}
                                    _ = v40
                                } else {
                                    v21 += 3.14159
                                    switch fi26 {
                                    case 596:
                                        var v41 uint32 = 42
                                        _ = v41
                                    case 4294967295:
                                        if false {
                                            var v42 string = (("foo bar" + "a") + "\n\t")
                                            if true {
                                                v4 = append(v4, fi25)
                                            } else {
                                                var v43 uint32 = 1
                                                _ = v43
                                            }
                                            var v44 float64 = v21
                                            _ = v42
                                            _ = v44
                                        } else {
                                            v1 = (v5 - v3)
                                        }
                                    case 495:
                                        v45 := GenS5{F0: -1, F1: -1}
                                        _ = v45
                                    default:
                                        switch "hello" {
                                        case "foo bar":
                                            switch fi26 {
                                            case 32767:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case 1:
                                                switch fi26 {
                                                case 127:
                                                    v3 -= v1
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v46 int = 9223372036854775807
                                                _ = v46
                                            default:
                                                v47 := GenS6{F0: "\n\t", F1: 1}
                                                _ = v47
                                            }
                                        case "hello":
                                            if len(v4) > 0 {
                                                v4[0] = fi26
                                            }
                                        case "a":
                                            _ = len(v4)
                                        default:
                                            var v48 int32 = 2147483647
                                            _ = v48
                                        }
                                    }
                                }
                                var v49 int64 = 9223372036854775807
                                var v50 bool = (v49 <= v49)
                                if false {
                                    v49 = v49
                                    fmt.Println(v5, v3)
                                    v1 = v1
                                } else {
                                    v51 := GenS3{F0: (fi25 + fi26)}
                                    var v52 uint8 = v8
                                    v1 += v5
                                    v53 := GenS4{F0: 1}
                                    v1 -= 698
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                                _ = v49
                                _ = v50
                            } else {
                                v1 = v5
                                if (true || false) {
                                    if false {
                                        v21 = float64((1e10 + (-2.25 + ((-2.25 + ((float32(v21) - 3.14159) - ((0.1 / 1.0) - -2.25))) / 1e10))))
                                        v2 -= 2
                                        var v54 uint64 = 1
                                        _ = v54
                                    } else {
                                        v4 = append(v4, (-(fi25)))
                                        switch fi26 {
                                        case -2147483648:
                                            switch (^(fi25)) {
                                            case 9223372036854775807:
                                                switch "a" {
                                                case "hello":
                                                    v4 = append(v4, fi25)
                                                }
                                            case 0:
                                                fmt.Println(fi26)
                                            case 42:
                                                v5 = v1
                                                var v55 uint = 342
                                                _ = v55
                                            }
                                        case 271:
                                            fmt.Println(v5, v3)
                                        case 1:
                                            var v56 float64 = 0.0
                                            v57 := GenS4{F0: 255}
                                            fmt.Println(v56, v1, v8)
                                            _ = v56
                                            _ = v57
                                        default:
                                            v5 = (v3 << ((0) & 63))
                                            fmt.Println(v21, v3)
                                        }
                                        v58 := GenS5{F0: 100, F1: 561}
                                        v1 = 1000
                                        var v59 uint64 = 0
                                        var v60 int = 411
                                        switch "a" {
                                        case "\n\t":
                                            var v61 int64 = 9223372036854775807
                                            _ = v61
                                        case "hello":
                                            var v62 int32 = 42
                                            _ = v62
                                        case "a":
                                            fmt.Println(v60)
                                        default:
                                            if false {
                                                if (v60 < (fi26 / ((v60) | 1))) {
                                                    v59 = (v59 / ((0) | 1))
                                                    var v63 string = ""
                                                    switch (fi25 | (^(fi25))) {
                                                    case 669:
                                                        if true {
                                                            if (v63 < v63) {
                                                                if true {
                                                                    switch v63 {
                                                                    case "a":
                                                                        v8 -= 2
                                                                    }
                                                                    v3 += (v3 ^ v1)
                                                                    v1 += (v3 + v1)
                                                                } else {
                                                                    if true {
                                                                        switch (^(fi25)) {
                                                                        case 331:
                                                                            if (false || ((false && false) || (1000 >= (v1 & -54)))) {
                                                                                var v64 float64 = float64((0.1 / 1.0))
                                                                                _ = v64
                                                                            }
                                                                        case 128:
                                                                            v65 := GenS7{F0: v59, F1: -1, F2: (^(v59))}
                                                                            _ = v65
                                                                        case 42:
                                                                            v66 := GenS2{F0: 128, F1: 2147483647}
                                                                            _ = v66
                                                                        default:
                                                                            v67 := GenS5{F0: 100, F1: 144}
                                                                            _ = v67
                                                                        }
                                                                        var v68 int = v60
                                                                        if false {
                                                                            v2 = v8
                                                                            var v69 int64 = -9223372036854775808
                                                                            switch -1 {
                                                                            case 570:
                                                                                v8 = (v8 + v8)
                                                                            }
                                                                            v21 = v21
                                                                            _ = v69
                                                                        } else {
                                                                            if (true || (65535 <= 0)) {
                                                                                switch (v68 % (((v68 >> ((0) & 63))) | 1)) {
                                                                                case -1:
                                                                                    v4 = append(v4, fi25)
                                                                                    v21 = v21
                                                                                    v1 -= (v3 - (v3 >> ((0) & 63)))
                                                                                case 32767:
                                                                                    var v70 string = "\n\t"
                                                                                    _ = v70
                                                                                default:
                                                                                    var v71 uint16 = 32767
                                                                                    _ = v71
                                                                                }
                                                                                _ = len(v4)
                                                                                if (true || true) {
                                                                                    var v72 uint8 = v8
                                                                                    var v73 uint64 = 42
                                                                                    v8 -= v8
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                }
                                                                            }
                                                                            var v74 int16 = (v3 * (v1 ^ 1000))
                                                                            if true {
                                                                                if true {
                                                                                    v21 = (1.5 * (float64(((1e10 + 3.14159) / 0.1)) + (1e10 / v21)))
                                                                                    v59 += v59
                                                                                    if false {
                                                                                        if true {
                                                                                            var v75 float64 = ((v21 * 1.5) / v21)
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = (v68 + (fi25 / ((fi25) | 1)))
                                                                                            }
                                                                                            _ = v75
                                                                                        }
                                                                                        if true {
                                                                                            var v76 float64 = (v21 / v21)
                                                                                            _ = v76
                                                                                        } else {
                                                                                            v77 := GenS8{F0: (^(v59))}
                                                                                            _ = v77
                                                                                        }
                                                                                        var v78 int64 = 1
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                                fmt.Println(v8)
                                                                            }
                                                                            _ = v74
                                                                        }
                                                                        _ = v68
                                                                    }
                                                                    fmt.Println(v59, fi26)
                                                                }
                                                                fmt.Println(v5, v63, v60)
                                                            }
                                                        }
                                                    case 127:
                                                        v3 += -32768
                                                        if ((fi25 * (fi25 ^ fi26)) <= fi26) {
                                                            if false {
                                                                fmt.Println(v8)
                                                                if (int16(v8) > 1) {
                                                                    v63 = "a"
                                                                } else {
                                                                    v79 := GenS9{F0: v60}
                                                                    var v80 string = v63
                                                                    if true {
                                                                        if true {
                                                                            v60 = fi26
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            v4 = append(v4, fi26)
                                                                        }
                                                                    }
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                v81 := GenS6{F0: "", F1: (^(v59))}
                                                                if (29 > -112) {
                                                                    v8 = (v8 * v8)
                                                                    v60 = (fi26 ^ (v60 * (v60 / ((int(v8)) | 1))))
                                                                }
                                                                _ = v81
                                                            }
                                                            var v82 string = "foo bar"
                                                            _ = v82
                                                        } else {
                                                            switch fi26 {
                                                            case 128:
                                                                var v83 uint8 = (v2 ^ (v2 / (((v2 << ((4294967295) & 63))) | 1)))
                                                                _ = v83
                                                            case -9223372036854775808:
                                                                var v84 int16 = (v5 >> ((18446744073709551615) & 63))
                                                                _ = v84
                                                            case 32767:
                                                                var v85 string = v63
                                                                _ = v85
                                                            }
                                                        }
                                                    default:
                                                        if (v3 > -32768) {
                                                            if (true || false) {
                                                                if (!(!false)) {
                                                                    if (float64(((-2.25 / float32(1e10)) * (float32(((float64(-2.25) + v21) * v21)) + 1e10))) != (0.1 / (v21 * v21))) {
                                                                        switch fi25 {
                                                                        case 32768:
                                                                            var v86 int32 = 437
                                                                            _ = v86
                                                                        case 0:
                                                                            var v87 uint8 = uint8(v59)
                                                                            _ = v87
                                                                        case -1:
                                                                            if len(v4) > 0 {
                                                                                v4[0] = -178
                                                                            }
                                                                        default:
                                                                            fmt.Println(v3, fi26)
                                                                        }
                                                                        fmt.Println(v60)
                                                                    } else {
                                                                        v63 = v63
                                                                    }
                                                                    fmt.Println(v5)
                                                                    v88 := GenS4{F0: 42}
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (^(fi26))
                                                                    }
                                                                    _ = v88
                                                                }
                                                                v60 -= 0
                                                            } else {
                                                                var v89 int32 = -2147483648
                                                                _ = v89
                                                            }
                                                            if len(v4) > 0 {
                                                                v4[0] = fi26
                                                            }
                                                            var v90 uint64 = (v59 / ((v59) | 1))
                                                            _ = v90
                                                        } else {
                                                            _ = len(v4)
                                                            var v91 int8 = -128
                                                            fmt.Println(v1, v3)
                                                            _ = v91
                                                        }
                                                    }
                                                    _ = v63
                                                } else {
                                                    v4 = append(v4, (v60 % (((fi25 & (v60 >> ((65535) & 63)))) | 1)))
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = (fi26 ^ 1)
                                                }
                                                v8 = 128
                                            }
                                        }
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    }
                                    v3 = (-(v1))
                                }
                                var v92 uint16 = 677
                                _ = v92
                            }
                        } else {
                            v2 -= (v8 / ((4) | 1))
                            fmt.Println(v3, v2)
                        }
                    case 726:
                        switch ("hello" + ("a" + "")) {
                        case "a":
                            v8 = v2
                        }
                    }
                }
                _ = fi26
            }
        default:
            fmt.Println(v1, v3)
        }
        _ = fi25
    }
    var v93 uint = 1
    v94 := GenS1{F0: 2, F1: v3}
    if (!(false && false)) {
        if true {
            v21 = ((0.0 * (v21 * (((v21 / (v21 + (float64(((1e10 - float32(v21)) * 1e10)) + 1.5))) / (v21 / (v21 * 0.0))) - v21))) + (0.0 - (v21 - v21)))
        }
        v5 -= -1
    } else {
        for fi95 := 0; fi95 < 7; fi95++ {
            for fi96 := 0; fi96 < 5; fi96++ {
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v97 := GenS8{F0: 2}
                v4 = append(v4, fi95)
                if false {
                    fmt.Println(v8, v2)
                    v8 -= v8
                    fmt.Println(v93)
                }
                _ = fi96
                _ = v97
            }
            _ = fi95
        }
        v21 -= 0.0
        var v98 uint64 = 86
        _ = v98
    }
    switch "hello" {
    case "hello":
        for fi99 := 0; fi99 < 12; fi99++ {
            v21 += ((v21 + (v21 + v21)) + ((v21 - 0.1) + v21))
            var v100 float32 = (1.5 / ((-2.25 - (1e10 / 1.0)) + float32((-2.25 - ((v21 / float64(1e10)) / ((float64(3.14159) + 0.1) - (1.5 / 3.14159)))))))
            v100 -= float32(((v21 / 0.1) + float64((3.14159 + 0.1))))
            _ = fi99
            _ = v100
        }
    default:
        v5 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v93
    _ = v94
}
