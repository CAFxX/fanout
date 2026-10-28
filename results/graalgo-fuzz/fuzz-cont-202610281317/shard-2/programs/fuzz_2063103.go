package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 string }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 string; F1 string }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 float32 = -2.25
    var v4 float64 = 0.1
    var v5 uint = 4294967295
    v4 -= v4
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(v4, v5, v4)
        _ = fi6
    }
    v7 := GenS1{F0: -2147483648, F1: (v5 << ((v5) & 63)), F2: ""}
    v5 -= v5
    v8 := []int{4294967295, -871, 219}
    if ((304 != 0) && (true || (127 >= 2))) {
        for fi9 := 0; fi9 < 10; fi9++ {
            switch "a" {
            case "foo bar":
                switch (fi9 | fi9) {
                case -32:
                    v10 := GenS1{F0: -2147483648, F1: v5, F2: "a"}
                    _ = v10
                case 0:
                    v11 := GenS1{F0: -2147483648, F1: v5, F2: (("hello" + "\n\t") + "foo bar")}
                    _ = v11
                case 32768:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        var v13 float64 = (v4 - ((-2.25 / 1e10) + -2.25))
                        var v14 string = ""
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                default:
                    var v15 int16 = 1
                    _ = v15
                }
            default:
                switch fi9 {
                case 94:
                    if (!(v3 != 0.0)) {
                        v3 = float32(v4)
                    } else {
                        switch fi9 {
                        case 527:
                            for fi16 := 0; fi16 < 2; fi16++ {
                                switch (fi9 & fi16) {
                                case -574:
                                    delete(v2, 1)
                                default:
                                    switch "" {
                                    case "":
                                        fmt.Println(fi9, v5, v3)
                                    case "\n\t":
                                        if (true || false) {
                                            if ((18446744073709551615 <= 807) && true) {
                                                if len(v8) > 0 {
                                                    v8[0] = 398
                                                }
                                                fmt.Println(fi16, fi16, fi9)
                                                switch ((("foo bar" + "") + "a") + "foo bar") {
                                                case "a":
                                                    var v17 uint16 = 1
                                                    _ = v17
                                                case "foo bar":
                                                    if (255 > 128) {
                                                        v4 += v4
                                                        if (255 == 42) {
                                                            var v18 bool = true
                                                            fmt.Println(fi16, v3)
                                                            var v19 uint64 = 42
                                                            v20 := GenS1{F0: -1, F1: v5, F2: "hello"}
                                                            _ = v18
                                                            _ = v19
                                                            _ = v20
                                                        }
                                                    }
                                                }
                                            } else {
                                                if false {
                                                    var v21 int64 = 42
                                                    v21 += v21
                                                    _ = v21
                                                } else {
                                                    var v22 int = (fi16 << ((v5) & 63))
                                                    if (0 > 581) {
                                                        switch 9223372036854775807 {
                                                        case -9223372036854775808:
                                                            fmt.Println(v22)
                                                        case 127:
                                                            if true {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    if false {
                                                                        var v23 float32 = v3
                                                                        var v24 bool = ((^(v5)) <= 39)
                                                                        var v25 float32 = 1.5
                                                                        _ = v23
                                                                        _ = v24
                                                                        _ = v25
                                                                    } else {
                                                                        switch -9223372036854775808 {
                                                                        case 369:
                                                                            if false {
                                                                                fmt.Println(v3, fi9)
                                                                            } else {
                                                                                var v26 int32 = 2147483647
                                                                                switch (^(fi16)) {
                                                                                case -361:
                                                                                    var v27 int = 42
                                                                                    var v28 int32 = v26
                                                                                    var v29 int16 = 42
                                                                                    _ = v27
                                                                                    _ = v28
                                                                                    _ = v29
                                                                                case 0:
                                                                                    v3 = -2.25
                                                                                case 32767:
                                                                                    fmt.Println(v5, v4, fi9)
                                                                                default:
                                                                                    if true {
                                                                                        v26 -= 42
                                                                                    } else {
                                                                                        var v30 uint32 = 2147483648
                                                                                        fmt.Println(v26, v3)
                                                                                        var v31 uint64 = 0
                                                                                        var v32 int32 = int32(v31)
                                                                                        _ = v30
                                                                                        _ = v31
                                                                                        _ = v32
                                                                                    }
                                                                                }
                                                                                var v33 uint64 = 4294967295
                                                                                _ = v26
                                                                                _ = v33
                                                                            }
                                                                        }
                                                                        var v34 uint8 = 13
                                                                        switch (fi9 / (((fi16 % ((fi16) | 1))) | 1)) {
                                                                        case 0:
                                                                            v34 += 0
                                                                        case -9223372036854775808:
                                                                            if ((true && false) && true) {
                                                                                delete(v1, fi16)
                                                                                switch ("\n\t" + "a") {
                                                                                case "foo bar":
                                                                                    var v35 int64 = 9223372036854775807
                                                                                    _ = v35
                                                                                case "hello":
                                                                                    v4 = (0.0 - 0.1)
                                                                                }
                                                                            }
                                                                        default:
                                                                            if (!false) {
                                                                                if false {
                                                                                    var v36 int64 = -1
                                                                                    _ = v36
                                                                                } else {
                                                                                    if ("a" <= ((((("foo bar" + "\n\t") + "") + "foo bar") + "hello") + "\n\t")) {
                                                                                        var v37 float32 = v3
                                                                                        var v38 uint = v5
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            fmt.Println(v34, v4)
                                                                                        case "foo bar":
                                                                                            v39 := GenS1{F0: 2147483647, F1: 1, F2: "a"}
                                                                                            delete(v1, fi9)
                                                                                            _ = v39
                                                                                        case "a":
                                                                                            switch -1 {
                                                                                            case -291:
                                                                                                switch ("a" + "foo bar") {
                                                                                                case "\n\t":
                                                                                                    v3 = (v37 - ((3.14159 - (v3 * v3)) * (v37 / float32((v4 * 3.14159)))))
                                                                                                }
                                                                                            case -1:
                                                                                                v3 = v3
                                                                                            default:
                                                                                                v34 -= v34
                                                                                            }
                                                                                        default:
                                                                                            if ((v34 ^ 128) == 2) {
                                                                                                v22 = (fi16 << ((v5) & 63))
                                                                                            }
                                                                                        }
                                                                                        _ = v37
                                                                                        _ = v38
                                                                                    }
                                                                                }
                                                                                v22 = (fi16 ^ (v22 / (((fi9 & fi9)) | 1)))
                                                                            }
                                                                        }
                                                                        _ = v34
                                                                    }
                                                                default:
                                                                    fmt.Println(fi9)
                                                                }
                                                                v3 += ((v3 - v3) * (0.0 + 1e10))
                                                                v3 = 3.14159
                                                            } else {
                                                                delete(v2, 65535)
                                                                v2[652] = v22
                                                                var v40 float32 = ((v3 * (1.5 + v3)) - (3.14159 - (v3 + v3)))
                                                                if false {
                                                                    fmt.Println(v3, v22, v4)
                                                                    var v41 uint = v5
                                                                    if true {
                                                                        if true {
                                                                            if true {
                                                                                delete(v1, (fi16 / ((v22) | 1)))
                                                                            } else {
                                                                                v4 -= ((-2.25 / v4) - (1.5 / v4))
                                                                            }
                                                                            switch "a" {
                                                                            case "":
                                                                                switch -1 {
                                                                                case 128:
                                                                                    var v42 string = ("\n\t" + "a")
                                                                                    _ = v42
                                                                                case 227:
                                                                                    switch (fi16 / (((fi16 << (((^(v41))) & 63))) | 1)) {
                                                                                    case 32767:
                                                                                        v43 := GenS2{F0: 0, F1: 2}
                                                                                        _ = v43
                                                                                    case -470:
                                                                                        switch v22 {
                                                                                        case 790:
                                                                                            fmt.Println(v4, fi16)
                                                                                        case 14:
                                                                                            var v44 float64 = v4
                                                                                            _ = v44
                                                                                        case 42:
                                                                                            var v45 int8 = 0
                                                                                            _ = v45
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v40 = v40
                                                                                }
                                                                            case "\n\t":
                                                                                var v46 int64 = 9223372036854775807
                                                                                _ = v46
                                                                            default:
                                                                                var v47 int16 = -32768
                                                                                _ = v47
                                                                            }
                                                                            if (1000 >= -1) {
                                                                                v48 := GenS1{F0: -559, F1: 18446744073709551615, F2: (("a" + (("foo bar" + "") + ("foo bar" + "a"))) + ("hello" + ("a" + "")))}
                                                                                _ = v48
                                                                            }
                                                                        }
                                                                        var v49 int16 = -496
                                                                        v50 := v1[fi9]
                                                                        _ = v49
                                                                        _ = v50
                                                                    } else {
                                                                        var v51 string = ""
                                                                        _ = v51
                                                                    }
                                                                    v40 -= (((1e10 - (1e10 + (v40 + -2.25))) / 0.1) * ((((float32(v4) - v40) + v40) + v3) + v40))
                                                                    v52 := GenS3{F0: 9223372036854775807, F1: -1}
                                                                    switch v22 {
                                                                    case 32767:
                                                                        var v53 uint32 = 238
                                                                        _ = v53
                                                                    case -349:
                                                                        var v54 bool = (false && true)
                                                                        _ = v54
                                                                    case 127:
                                                                        switch fi9 {
                                                                        case 17:
                                                                            v55 := v2[2147483648]
                                                                            _ = v55
                                                                        case 977:
                                                                            if true {
                                                                                v4 -= v4
                                                                                var v56 uint64 = 4294967295
                                                                                var v57 float64 = 1e10
                                                                                if false {
                                                                                    if (uint32(v56) <= uint32(v41)) {
                                                                                        v57 = v4
                                                                                        v5 -= v41
                                                                                        if ((!((v41 % ((v5) | 1)) < v41)) || (!false)) {
                                                                                            v58 := GenS2{F0: 1, F1: 0}
                                                                                            var v59 uint = v5
                                                                                            var v60 uint32 = 65535
                                                                                            _ = v58
                                                                                            _ = v59
                                                                                            _ = v60
                                                                                        }
                                                                                    }
                                                                                    _ = len(v8)
                                                                                    switch 127 {
                                                                                    case 32767:
                                                                                        if (((3.14159 >= (v3 * v3)) && true) || true) {
                                                                                            v61 := v2[2147483647]
                                                                                            _ = v61
                                                                                        } else {
                                                                                            if true {
                                                                                                switch -41 {
                                                                                                case -2147483648:
                                                                                                    v22 = (fi9 ^ (fi16 * (fi9 % (((fi16 & v22)) | 1))))
                                                                                                default:
                                                                                                    v40 -= -2.25
                                                                                                }
                                                                                                var v62 string = ("foo bar" + "a")
                                                                                                v4 += v4
                                                                                                _ = v62
                                                                                            } else {
                                                                                                v40 = (-2.25 + ((float32(float64(float32(((((3.14159 * 0.0) + (1e10 / (v4 + v57))) * (v4 * 0.1)) + ((v4 * v4) + v57))))) - v3) * 1e10))
                                                                                                fmt.Println(fi9)
                                                                                                v63 := GenS2{F0: 10, F1: 0}
                                                                                                switch fi16 {
                                                                                                case 1:
                                                                                                    v1[(fi16 * fi9)] = ("" + "")
                                                                                                case -465:
                                                                                                    v4 = (v4 + (-2.25 * v57))
                                                                                                default:
                                                                                                    var v64 uint64 = v56
                                                                                                    _ = v64
                                                                                                }
                                                                                                _ = v63
                                                                                            }
                                                                                            fmt.Println(v56, v40, fi9)
                                                                                            switch fi9 {
                                                                                            case 127:
                                                                                                if len(v8) > 0 {
                                                                                                    v8[0] = v22
                                                                                                }
                                                                                            case 602:
                                                                                                var v65 int64 = 9223372036854775807
                                                                                                _ = v65
                                                                                            default:
                                                                                                v66 := GenS3{F0: 1, F1: -2147483648}
                                                                                                _ = v66
                                                                                            }
                                                                                        }
                                                                                    case -9223372036854775808:
                                                                                        if false {
                                                                                            v3 -= (v40 * v40)
                                                                                            v57 = 3.14159
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                var v67 int = fi9
                                                                                                _ = v67
                                                                                            case "foo bar":
                                                                                                switch (v22 | (v22 & fi16)) {
                                                                                                case 32767:
                                                                                                    v57 -= v57
                                                                                                case 127:
                                                                                                    var v68 uint64 = 568
                                                                                                    _ = v68
                                                                                                case 42:
                                                                                                    v69 := GenS4{F0: -1}
                                                                                                    _ = v69
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v56)
                                                                                            }
                                                                                        } else {
                                                                                            v22 -= (fi9 ^ fi16)
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v70 uint64 = (^(v56))
                                                                                    if (100 <= 119) {
                                                                                        var v71 uint16 = 42
                                                                                        if true {
                                                                                            var v72 uint64 = (v56 * v56)
                                                                                            _ = v72
                                                                                        } else {
                                                                                            switch "foo bar" {
                                                                                            case "":
                                                                                                if ((v70 + 0) != v56) {
                                                                                                    v5 = v41
                                                                                                    switch "\n\t" {
                                                                                                    case "\n\t":
                                                                                                        if false {
                                                                                                            var v73 uint32 = 2147483648
                                                                                                            _ = v73
                                                                                                        } else {
                                                                                                            fmt.Println(v71, fi16)
                                                                                                            v70 = 2
                                                                                                        }
                                                                                                    case "":
                                                                                                        switch 333 {
                                                                                                        case 2147483647:
                                                                                                            switch fi16 {
                                                                                                            case 4294967295:
                                                                                                                v70 += v56
                                                                                                            case 32767:
                                                                                                                v41 += (v41 >> (((v5 | v41)) & 63))
                                                                                                            case 740:
                                                                                                                var v74 int8 = 42
                                                                                                                _ = v74
                                                                                                            default:
                                                                                                                v75 := GenS1{F0: 0, F1: 4294967295, F2: "hello"}
                                                                                                                _ = v75
                                                                                                            }
                                                                                                        default:
                                                                                                            if (true || (v40 > v40)) {
                                                                                                                switch "\n\t" {
                                                                                                                case "foo bar":
                                                                                                                    var v76 uint8 = 1
                                                                                                                    _ = v76
                                                                                                                default:
                                                                                                                    if (2147483648 > 0) {
                                                                                                                        v5 = (^(v41))
                                                                                                                    } else {
                                                                                                                        v5 = v5
                                                                                                                        v77 := v2[2]
                                                                                                                        _ = v77
                                                                                                                    }
                                                                                                                }
                                                                                                                v4 -= v4
                                                                                                                var v78 int32 = 455
                                                                                                                _ = v78
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        var v79 uint = v5
                                                                                                        _ = v79
                                                                                                    }
                                                                                                    var v80 int8 = 127
                                                                                                    _ = v80
                                                                                                }
                                                                                            default:
                                                                                                v3 -= (1e10 + (0.1 / v40))
                                                                                                v2[4294967295] = v22
                                                                                                v81 := GenS3{F0: -9223372036854775808, F1: 2147483647}
                                                                                                _ = v81
                                                                                            }
                                                                                            if ("\n\t" < ("foo bar" + "")) {
                                                                                                var v82 uint64 = v70
                                                                                                switch (fi9 << ((42) & 63)) {
                                                                                                case 4294967295:
                                                                                                    switch ("hello" + ("a" + ("\n\t" + "foo bar"))) {
                                                                                                    case "a":
                                                                                                        fmt.Println(v41, fi16)
                                                                                                    default:
                                                                                                        var v83 uint32 = 1
                                                                                                        _ = v83
                                                                                                    }
                                                                                                }
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    switch 127 {
                                                                                                    case 607:
                                                                                                        v56 = v56
                                                                                                    default:
                                                                                                        v82 = (^(v82))
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v8 = append(v8, 4294967295)
                                                                                                default:
                                                                                                    var v84 uint = v5
                                                                                                    _ = v84
                                                                                                }
                                                                                                _ = v82
                                                                                            }
                                                                                            var v85 float64 = 0.0
                                                                                            _ = v85
                                                                                        }
                                                                                        _ = v71
                                                                                    }
                                                                                    var v86 uint32 = 1
                                                                                    var v87 int8 = -94
                                                                                    _ = v70
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                }
                                                                                if len(v8) >= 2 {
                                                                                    v8 = v8[1:len(v8)]
                                                                                }
                                                                                _ = v56
                                                                                _ = v57
                                                                            } else {
                                                                                switch (fi16 * (fi16 / (((v22 * v22)) | 1))) {
                                                                                case 9223372036854775807:
                                                                                    v41 -= v5
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            if (1 == -1) {
                                                                                fmt.Println(fi16, v22, fi16)
                                                                            }
                                                                            v5 += (v5 % ((v5) | 1))
                                                                            if (!(!true)) {
                                                                                var v88 uint32 = 2147483648
                                                                                if (!false) {
                                                                                    v89 := GenS2{F0: 115, F1: 2147483648}
                                                                                    if true {
                                                                                        v4 = 1e10
                                                                                        v90 := GenS2{F0: 42, F1: (v88 | (v88 >> ((820) & 63)))}
                                                                                        if false {
                                                                                            var v91 bool = (true && true)
                                                                                            _ = v91
                                                                                        }
                                                                                        _ = v90
                                                                                    } else {
                                                                                        if (false || ((v5 ^ uint(v88)) != (^(v5)))) {
                                                                                            if false {
                                                                                                _ = len(v1)
                                                                                                v5 = 163
                                                                                            } else {
                                                                                                fmt.Println(v22)
                                                                                            }
                                                                                            _ = len(v1)
                                                                                            if false {
                                                                                                fmt.Println(fi9, v22, v40)
                                                                                                v92 := GenS3{F0: int64(v88), F1: -1}
                                                                                                var v93 uint8 = 90
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                            }
                                                                                        } else {
                                                                                            v94 := GenS4{F0: 1}
                                                                                            v95 := GenS4{F0: 0}
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        }
                                                                                    }
                                                                                    v96 := GenS2{F0: 42, F1: v88}
                                                                                    _ = v89
                                                                                    _ = v96
                                                                                } else {
                                                                                    v40 = v3
                                                                                    switch v22 {
                                                                                    case -802:
                                                                                        switch ("foo bar" + "foo bar") {
                                                                                        case "\n\t":
                                                                                            var v97 uint16 = 0
                                                                                            _ = v97
                                                                                        case "hello":
                                                                                            switch ("" + ("" + "")) {
                                                                                            case "":
                                                                                                v3 = 0.0
                                                                                            case "foo bar":
                                                                                                fmt.Println(v40, v40, v41)
                                                                                            default:
                                                                                                v41 -= (v41 ^ (v41 >> ((604) & 63)))
                                                                                            }
                                                                                            if (18446744073709551615 > 2) {
                                                                                                v3 = float32(v4)
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    switch "" {
                                                                                                    case "a":
                                                                                                        v4 = v4
                                                                                                    case "\n\t":
                                                                                                        var v98 int = 831
                                                                                                        _ = v98
                                                                                                    }
                                                                                                default:
                                                                                                    v99 := GenS3{F0: 0, F1: 0}
                                                                                                    _ = v99
                                                                                                }
                                                                                                if (!true) {
                                                                                                    if true {
                                                                                                        var v100 float32 = v40
                                                                                                        var v101 float64 = v4
                                                                                                        _ = v100
                                                                                                        _ = v101
                                                                                                    }
                                                                                                    fmt.Println(v41, v41, v88)
                                                                                                } else {
                                                                                                    switch "foo bar" {
                                                                                                    case "foo bar":
                                                                                                        v22 = (fi16 << ((2) & 63))
                                                                                                    case "":
                                                                                                        var v102 float64 = ((v4 * 0.1) - (v4 - v4))
                                                                                                        _ = v102
                                                                                                    case "\n\t":
                                                                                                        v3 = v40
                                                                                                    default:
                                                                                                        delete(v1, (^(v22)))
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v4)
                                                                                        }
                                                                                        var v103 uint = 65535
                                                                                        _ = v103
                                                                                    }
                                                                                    fmt.Println(v5, v88)
                                                                                }
                                                                                v4 -= ((v4 / (float64(1.5) + v4)) - float64(((((0.1 - -2.25) / 1e10) + v40) / v40)))
                                                                                _ = v88
                                                                            } else {
                                                                                fmt.Println(v41)
                                                                            }
                                                                            delete(v2, 281)
                                                                            v4 += 0.1
                                                                            if false {
                                                                                v104 := GenS1{F0: -888, F1: v41, F2: ("a" + ("" + ("foo bar" + "a")))}
                                                                                _ = v104
                                                                            } else {
                                                                                var v105 uint32 = 294
                                                                                _ = v105
                                                                            }
                                                                        }
                                                                    }
                                                                    v5 = (v41 << ((65535) & 63))
                                                                    v3 -= v3
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                    var v106 int = v22
                                                                    _ = v106
                                                                    _ = v41
                                                                    _ = v52
                                                                }
                                                                _ = v40
                                                            }
                                                        case -185:
                                                            fmt.Println(fi9, v22, v22)
                                                        default:
                                                            v107 := GenS5{F0: 32767}
                                                            _ = v107
                                                        }
                                                        v108 := GenS6{F0: -2147483648, F1: 53}
                                                        if (!false) {
                                                            var v109 uint16 = 32767
                                                            if len(v8) > 0 {
                                                                v8[0] = (fi16 | v22)
                                                            }
                                                            _ = len(v1)
                                                            if false {
                                                                v22 = (fi9 - int(v109))
                                                            }
                                                            _ = v109
                                                        } else {
                                                            if (1e10 <= 3.14159) {
                                                                v110 := v2[0]
                                                                v110 = (v110 | v110)
                                                                if false {
                                                                    fmt.Println(v110)
                                                                } else {
                                                                    switch v22 {
                                                                    case 128:
                                                                        v110 -= v22
                                                                    }
                                                                    if (true && false) {
                                                                        v111 := GenS1{F0: -2147483648, F1: v5, F2: ("\n\t" + ("hello" + "foo bar"))}
                                                                        _ = v111
                                                                    } else {
                                                                        v110 += v22
                                                                    }
                                                                }
                                                                v22 = (fi9 ^ -206)
                                                                delete(v1, (fi9 / (((fi9 << ((18446744073709551615) & 63))) | 1)))
                                                                _ = v110
                                                            }
                                                            var v112 uint16 = 717
                                                            var v113 float32 = ((v3 / float32((v4 / float64(3.14159)))) / float32((((v4 - v4) - v4) + (1.5 * ((v4 - 0.1) / float64(v3))))))
                                                            _ = v112
                                                            _ = v113
                                                        }
                                                        _ = v108
                                                    } else {
                                                        v114 := GenS7{F0: 4294967295}
                                                        var v115 uint32 = 304
                                                        _ = v114
                                                        _ = v115
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v116 int16 = 0
                                                        _ = v116
                                                    }
                                                    _ = v22
                                                }
                                            }
                                            var v117 float64 = v4
                                            fmt.Println(v4)
                                            _ = v117
                                        }
                                    default:
                                        var v118 float32 = float32(0.0)
                                        _ = v118
                                    }
                                }
                                _ = fi16
                            }
                        case 32768:
                            v4 -= v4
                        default:
                            var v119 uint8 = 177
                            _ = v119
                        }
                        if (1 == v5) {
                            switch "foo bar" {
                            case "foo bar":
                                v3 -= float32((float64(v3) / v4))
                            case "\n\t":
                                switch "hello" {
                                case "a":
                                    v120 := GenS6{F0: 2147483647, F1: 1}
                                    _ = v120
                                default:
                                    switch (fi9 + fi9) {
                                    case 4294967295:
                                        if (true || (-2.25 <= v4)) {
                                            v121 := GenS8{F0: "\n\t", F1: 42}
                                            var v122 int64 = 0
                                            fmt.Println(v4, v4)
                                            _ = v121
                                            _ = v122
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            for fi123 := 0; fi123 < 17; fi123++ {
                                                if (false && (207 == 97)) {
                                                    v4 -= 0.1
                                                    v124 := GenS9{F0: "", F1: "foo bar"}
                                                    _ = v124
                                                }
                                                v125 := v1[(fi9 + fi9)]
                                                var v126 int64 = 1
                                                v125 = v125
                                                v125 = ("\n\t" + v125)
                                                switch v125 {
                                                case "\n\t":
                                                    _ = len(v8)
                                                    v125 = v125
                                                default:
                                                    var v127 uint32 = 65535
                                                    _ = v127
                                                }
                                                v5 += (v5 / ((v5) | 1))
                                                _ = fi123
                                                _ = v125
                                                _ = v126
                                            }
                                        case "a":
                                            if len(v8) > 0 {
                                                v8[0] = (-(fi9))
                                            }
                                        default:
                                            switch 1 {
                                            case 9223372036854775807:
                                                v4 += v4
                                            case 485:
                                                switch fi9 {
                                                case 864:
                                                    if len(v8) > 0 {
                                                        v8[0] = (fi9 | fi9)
                                                    }
                                                case -263:
                                                    var v128 float64 = (v4 / (((1e10 - ((float64((v3 - 3.14159)) + v4) - float64(v3))) / -2.25) / (float64((float32(float64(1.5)) * (0.0 / v3))) / (3.14159 - (float64(1.5) - (float64(v3) - (v4 + v4)))))))
                                                    _ = v128
                                                }
                                            default:
                                                fmt.Println(v3, fi9, fi9)
                                            }
                                        }
                                    }
                                }
                            default:
                                if ((199 <= 2147483647) || true) {
                                    v3 = v3
                                    for fi129 := 0; fi129 < 16; fi129++ {
                                        fmt.Println(fi9, fi9, v5)
                                        _ = fi129
                                    }
                                } else {
                                    fmt.Println(v3)
                                }
                            }
                            if (0 >= 239) {
                                if true {
                                    v3 = float32(v4)
                                    delete(v2, 42)
                                    var v130 int8 = -128
                                    _ = v130
                                }
                                v8 = append(v8, 965)
                                v3 = (v3 + v3)
                            } else {
                                v3 += ((0.0 - 1e10) / float32((((v4 - v4) - (0.1 * ((float64(v3) / 1.0) * float64(float32(v4))))) * (1e10 / float64(v3)))))
                                var v131 uint = (v5 >> ((42) & 63))
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                _ = v131
                            }
                            if (!true) {
                                var v132 uint8 = 201
                                _ = v132
                            } else {
                                var v133 int8 = 42
                                switch ("\n\t" + "a") {
                                case "a":
                                    var v134 string = "a"
                                    _ = v134
                                case "hello":
                                    if false {
                                        v4 = (v4 + float64(v3))
                                        switch (fi9 % ((fi9) | 1)) {
                                        case 1:
                                            v3 -= (v3 + float32(v4))
                                        case -164:
                                            var v135 uint64 = 18446744073709551615
                                            _ = v135
                                        case -277:
                                            for fi136 := 0; fi136 < 2; fi136++ {
                                                fmt.Println(v5)
                                                v3 = (v3 * 0.1)
                                                if true {
                                                    v4 += (v4 - -2.25)
                                                } else {
                                                    var v137 int64 = 9223372036854775807
                                                    fmt.Println(v5, v3, v133)
                                                    var v138 int64 = v137
                                                    _ = v137
                                                    _ = v138
                                                }
                                                _ = fi136
                                            }
                                        default:
                                            v139 := GenS1{F0: 2147483647, F1: 42, F2: ("hello" + "")}
                                            _ = v139
                                        }
                                        v140 := GenS8{F0: "hello", F1: -778}
                                        _ = v140
                                    }
                                case "":
                                    v3 = float32((v4 - 1e10))
                                default:
                                    var v141 int8 = v133
                                    _ = v141
                                }
                                _ = v133
                            }
                        }
                    }
                default:
                    var v142 uint8 = 174
                    _ = v142
                }
            }
            _ = fi9
        }
        var v143 int16 = 127
        _ = v143
    } else {
        _ = len(v2)
        fmt.Println(v5)
        _ = len(v8)
    }
    var v144 int8 = 1
    if (128 >= uint8(v5)) {
        v145 := v1[-9223372036854775808]
        _ = v145
    } else {
        fmt.Println(v3, v4, v4)
    }
    v146 := GenS1{F0: int32(v5), F1: v5, F2: "a"}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v4 = v4
    fmt.Println(v3, v5, v3)
    v3 -= (float32(v4) + (float32(0.1) * float32((v4 - 1e10))))
    var v147 float64 = float64((-2.25 * v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v144:", v144)
    fmt.Println("v147:", v147)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v144
    _ = v146
    _ = v147
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
