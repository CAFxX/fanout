package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 uint; F2 uint16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int16; F1 uint64; F2 uint16 }

func main() {
    var v1 int64 = 42
    var v2 int32 = 0
    var v3 uint32 = 2147483648
    fmt.Println(v2, v2)
    v2 += v2
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(v2, v1)
        _ = fi4
    }
    v3 = v3
    if false {
        v5 := GenS1{F0: false, F1: 970, F2: 1}
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 float32 = 1e10
                v1 += (v1 & (v1 | v1))
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        v9 := GenS2{F0: "\n\t"}
        _ = v5
        _ = v9
    } else {
        v10 := GenS1{F0: true, F1: 252, F2: 128}
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v3 = 661
        fmt.Println(fi11, fi11, v3)
        _ = fi11
    }
    switch (("a" + "a") + ("" + "a")) {
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        v12 := []int{1}
        _ = v12
    }
    var v13 uint8 = 2
    v14 := []int{684, 9223372036854775807, 182, 2147483647}
    for fi15 := 0; fi15 < 5; fi15++ {
        if false {
            if true {
                fmt.Println(fi15, v13)
                switch (("hello" + "hello") + "\n\t") {
                case "foo bar":
                    if (0.0 > 1e10) {
                        for fi16 := 0; fi16 < 7; fi16++ {
                            var v17 uint16 = uint16(v3)
                            if (v1 != (v1 % (((v1 | v1)) | 1))) {
                                v17 += v17
                                var v18 uint8 = v13
                                if (277 > 42) {
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                } else {
                                    v3 -= (^(v3))
                                    var v19 bool = true
                                    var v20 uint64 = 1
                                    _ = v19
                                    _ = v20
                                }
                                _ = v18
                            }
                            _ = fi16
                            _ = v17
                        }
                        var v21 int8 = -74
                        var v22 uint64 = 970
                        _ = v21
                        _ = v22
                    }
                case "a":
                    if (v3 <= v3) {
                        fmt.Println(fi15, v2, v3)
                        v23 := GenS3{F0: "foo bar", F1: ("\n\t" + ("hello" + ("" + "\n\t"))), F2: 2}
                        _ = v23
                    } else {
                        v24 := GenS4{F0: -128, F1: 1, F2: v13}
                        _ = v24
                    }
                default:
                    var v25 uint32 = (^(v3))
                    _ = v25
                }
            } else {
                v13 = v13
                fmt.Println(fi15)
            }
        } else {
            var v26 uint = 255
            _ = v26
        }
        var v27 uint = uint(v3)
        v2 = v2
        _ = fi15
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            if len(v14) > 0 {
                v14[0] = -9223372036854775808
            }
            _ = fi29
        }
        v30 := GenS2{F0: ("foo bar" + "hello")}
        _ = fi28
        _ = v30
    }
    v31 := GenS5{F0: v13}
    var v32 int16 = 42
    if (0 <= 1) {
        var v33 uint32 = v3
        switch 0 {
        case 128:
            var v34 uint = 44
            _ = v34
        case -324:
            switch int(v2) {
            case -926:
                fmt.Println(v3, v13, v13)
            case 4294967295:
                v2 += (v2 ^ -811)
            }
        case -2147483648:
            if false {
                fmt.Println(v32)
                v3 = (v3 << ((2) & 63))
                if true {
                    switch int(v2) {
                    case -170:
                        switch 9223372036854775807 {
                        case 0:
                            if (false || false) {
                                v35 := GenS6{F0: (v2 + v2), F1: -1}
                                var v36 string = (((("foo bar" + (("\n\t" + "") + "hello")) + (("hello" + "hello") + "")) + (("a" + "\n\t") + "foo bar")) + "\n\t")
                                _ = v35
                                _ = v36
                            } else {
                                switch 1 {
                                case 32767:
                                    if len(v14) > 0 {
                                        v14[0] = 4294967295
                                    }
                                default:
                                    v33 -= v33
                                }
                                switch "hello" {
                                case "a":
                                    v37 := GenS7{F0: 2147483647}
                                    _ = v37
                                case "":
                                    switch "foo bar" {
                                    case "hello":
                                        switch ("hello" + ("a" + "\n\t")) {
                                        case "a":
                                            switch 128 {
                                            case 2147483647:
                                                var v38 uint32 = 65535
                                                _ = v38
                                            default:
                                                for fi39 := 0; fi39 < 2; fi39++ {
                                                    _ = len(v14)
                                                    if (!true) {
                                                        if len(v14) > 0 {
                                                            v14[0] = fi39
                                                        }
                                                        for fi40 := 0; fi40 < 3; fi40++ {
                                                            fmt.Println(fi40)
                                                            v3 = (v33 & 2147483647)
                                                            switch 1 {
                                                            case 127:
                                                                v13 = v13
                                                            case 2147483647:
                                                                v41 := GenS3{F0: "a", F1: ("" + ("" + "")), F2: 42}
                                                                _ = v41
                                                            }
                                                            _ = fi40
                                                        }
                                                        if true {
                                                            for fi42 := 0; fi42 < 14; fi42++ {
                                                                if len(v14) > 0 {
                                                                    v14[0] = (fi39 + (^(fi39)))
                                                                }
                                                                _ = fi42
                                                            }
                                                            for fi43 := 0; fi43 < 15; fi43++ {
                                                                v33 += (v3 & (^(v33)))
                                                                v44 := GenS8{F0: (v3 >> ((0) & 63)), F1: 143, F2: 65535}
                                                                var v45 uint = 437
                                                                fmt.Println(fi39, fi39)
                                                                switch fi43 {
                                                                case 961:
                                                                    var v46 string = ("foo bar" + "a")
                                                                    _ = v46
                                                                case 9223372036854775807:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v47 int64 = (v1 / ((-9223372036854775808) | 1))
                                                                        _ = v47
                                                                    case "\n\t":
                                                                        switch ("" + "\n\t") {
                                                                        case "":
                                                                            fmt.Println(v2)
                                                                        case "foo bar":
                                                                            v14 = append(v14, fi39)
                                                                        case "\n\t":
                                                                            var v48 uint = (v45 >> (((v45 << ((v45) & 63))) & 63))
                                                                            _ = v48
                                                                        default:
                                                                            v49 := GenS6{F0: (v2 / (((v2 * v2)) | 1)), F1: (fi43 ^ fi39)}
                                                                            v1 = -122
                                                                            var v50 uint8 = (^(v13))
                                                                            _ = v49
                                                                            _ = v50
                                                                        }
                                                                    case "hello":
                                                                        v3 = v33
                                                                    }
                                                                case 127:
                                                                    if true {
                                                                        if (32767 == fi43) {
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v14 = append(v14, fi39)
                                                                            case "":
                                                                                var v51 uint8 = v13
                                                                                _ = v51
                                                                            default:
                                                                                switch (fi39 << (((v45 >> ((v45) & 63))) & 63)) {
                                                                                case 485:
                                                                                    fmt.Println(v33, v13)
                                                                                case 716:
                                                                                    switch 2147483647 {
                                                                                    case 128:
                                                                                        var v52 uint16 = 65535
                                                                                        _ = v52
                                                                                    case -705:
                                                                                        v13 = (^(v13))
                                                                                    case -83:
                                                                                        if false {
                                                                                            v13 = uint8(v3)
                                                                                            v13 -= (v13 * v13)
                                                                                            fmt.Println(v3)
                                                                                            var v53 uint8 = 42
                                                                                            _ = v53
                                                                                        }
                                                                                        v54 := GenS9{F0: v3}
                                                                                        _ = v54
                                                                                    default:
                                                                                        if len(v14) >= 2 {
                                                                                            v14 = v14[1:len(v14)]
                                                                                        }
                                                                                    }
                                                                                    v3 = 844
                                                                                case 2147483647:
                                                                                    var v55 int64 = v1
                                                                                    _ = v55
                                                                                }
                                                                            }
                                                                            v45 += v45
                                                                        } else {
                                                                            switch (fi39 | -284) {
                                                                            case 152:
                                                                                if (true || (!true)) {
                                                                                    if false {
                                                                                        v33 -= 2
                                                                                    } else {
                                                                                        var v56 int32 = (^(v2))
                                                                                        v57 := GenS8{F0: 2147483647, F1: 65535, F2: 42}
                                                                                        if len(v14) >= 2 {
                                                                                            v14 = v14[1:len(v14)]
                                                                                        }
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                    }
                                                                                    var v58 float64 = 1e10
                                                                                    if true {
                                                                                        var v59 uint = 18446744073709551615
                                                                                        _ = v59
                                                                                    }
                                                                                    _ = v58
                                                                                }
                                                                            default:
                                                                                fmt.Println(v32, fi39)
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                var v60 int16 = (v32 << (((^(v45))) & 63))
                                                                                _ = v60
                                                                            default:
                                                                                v33 = 395
                                                                            }
                                                                        }
                                                                        v32 = (v32 >> ((255) & 63))
                                                                        v1 += -1
                                                                    } else {
                                                                        v13 = 128
                                                                        switch 365 {
                                                                        case -497:
                                                                            var v61 uint8 = (v13 << ((v45) & 63))
                                                                            _ = v61
                                                                        case 1:
                                                                            var v62 uint64 = 42
                                                                            if true {
                                                                                if len(v14) > 0 {
                                                                                    v14[0] = fi39
                                                                                }
                                                                                fmt.Println(v1)
                                                                                v63 := GenS1{F0: ((v2 % ((int32(v62)) | 1)) > (v2 << ((v45) & 63))), F1: 32767, F2: 2147483647}
                                                                                _ = v63
                                                                            }
                                                                            _ = len(v14)
                                                                            _ = v62
                                                                        }
                                                                        v33 -= v3
                                                                    }
                                                                    _ = len(v14)
                                                                    var v64 float64 = (((0.1 * (0.0 / -2.25)) + (3.14159 - (((3.14159 - 1e10) + 0.0) / (1.5 + 0.0)))) * 1.5)
                                                                    _ = v64
                                                                default:
                                                                    v1 = (v1 + 1)
                                                                }
                                                                _ = fi43
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        } else {
                                                            if ((1.5 / (1e10 - (1.5 + -2.25))) != ((1.5 / (0.1 / 1.0)) + 1e10)) {
                                                                var v65 uint16 = uint16(v33)
                                                                _ = v65
                                                            }
                                                            v66 := GenS2{F0: "\n\t"}
                                                            _ = v66
                                                        }
                                                    }
                                                    _ = fi39
                                                }
                                            }
                                        case "hello":
                                            var v67 uint = 0
                                            _ = v67
                                        case "foo bar":
                                            if len(v14) > 0 {
                                                v14[0] = -265
                                            }
                                        default:
                                            var v68 int64 = v1
                                            _ = v68
                                        }
                                    case "a":
                                        if false {
                                            var v69 uint64 = 18446744073709551615
                                            if ((!(false && true)) && false) {
                                                for fi70 := 0; fi70 < 9; fi70++ {
                                                    if (true || (0 <= (fi70 << ((42) & 63)))) {
                                                        if (false && (((v33 - 605) != v33) && (!(v1 <= v1)))) {
                                                            v71 := GenS10{F0: v32, F1: 18446744073709551615, F2: 42}
                                                            _ = v71
                                                        } else {
                                                            switch 2147483647 {
                                                            case 32767:
                                                                fmt.Println(v3, v3)
                                                            case -915:
                                                                if false {
                                                                    v72 := GenS8{F0: (v3 << ((406) & 63)), F1: 255, F2: 0}
                                                                    _ = len(v14)
                                                                    v14 = append(v14, (fi70 - (fi70 << ((219) & 63))))
                                                                    _ = v72
                                                                } else {
                                                                    for fi73 := 0; fi73 < 19; fi73++ {
                                                                        _ = len(v14)
                                                                        _ = fi73
                                                                    }
                                                                }
                                                            case -2147483648:
                                                                _ = len(v14)
                                                            default:
                                                                switch fi70 {
                                                                case -9223372036854775808:
                                                                    v13 = uint8(v33)
                                                                case 1:
                                                                    for fi74 := 0; fi74 < 3; fi74++ {
                                                                        v2 -= -770
                                                                        _ = fi74
                                                                    }
                                                                case -1:
                                                                    v75 := GenS11{F0: v3}
                                                                    _ = v75
                                                                }
                                                            }
                                                        }
                                                        v1 += int64(v32)
                                                        v76 := GenS12{F0: v13}
                                                        _ = v76
                                                    }
                                                    switch "hello" {
                                                    case "a":
                                                        if (0 >= -1) {
                                                            var v77 uint32 = v3
                                                            _ = len(v14)
                                                            v13 -= (^(v13))
                                                            _ = v77
                                                        } else {
                                                            v13 = v13
                                                            var v78 uint = 42
                                                            _ = v78
                                                        }
                                                    default:
                                                        v69 += v69
                                                    }
                                                    _ = fi70
                                                }
                                                for fi79 := 0; fi79 < 1; fi79++ {
                                                    var v80 int32 = (v2 & v2)
                                                    v13 = v13
                                                    var v81 bool = true
                                                    _ = fi79
                                                    _ = v80
                                                    _ = v81
                                                }
                                            } else {
                                                v1 = 1
                                            }
                                            _ = v69
                                        } else {
                                            for fi82 := 0; fi82 < 18; fi82++ {
                                                fmt.Println(fi82)
                                                v83 := GenS3{F0: "foo bar", F1: "\n\t", F2: 677}
                                                var v84 int32 = (v2 / (((v2 >> ((4294967295) & 63))) | 1))
                                                _ = fi82
                                                _ = v83
                                                _ = v84
                                            }
                                        }
                                    case "foo bar":
                                        for fi85 := 0; fi85 < 2; fi85++ {
                                            v86 := GenS13{F0: 1000, F1: uint64(v33), F2: 2}
                                            _ = fi85
                                            _ = v86
                                        }
                                    }
                                case "foo bar":
                                    if true {
                                        for fi87 := 0; fi87 < 13; fi87++ {
                                            if true {
                                                var v88 uint = 100
                                                for fi89 := 0; fi89 < 7; fi89++ {
                                                    var v90 int64 = -9223372036854775808
                                                    v1 += int64(v13)
                                                    v14 = append(v14, -1)
                                                    _ = fi89
                                                    _ = v90
                                                }
                                                _ = v88
                                            }
                                            if ("hello" == "a") {
                                                v91 := GenS3{F0: "", F1: "a", F2: 4294967295}
                                                for fi92 := 0; fi92 < 2; fi92++ {
                                                    v3 -= (v33 >> ((255) & 63))
                                                    fmt.Println(v33)
                                                    _ = fi92
                                                }
                                                fmt.Println(v3, v13, v33)
                                                _ = v91
                                            }
                                            _ = fi87
                                        }
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    }
                                default:
                                    v13 += (v13 & v13)
                                }
                                if (493 >= 65535) {
                                    if ((v32 | (v32 - v32)) != (v32 % (((^(v32))) | 1))) {
                                        v1 += 9223372036854775807
                                        switch "a" {
                                        case "hello":
                                            for fi93 := 0; fi93 < 1; fi93++ {
                                                fmt.Println(v13)
                                                v13 = 168
                                                var v94 bool = ((-(v32)) >= v32)
                                                _ = fi93
                                                _ = v94
                                            }
                                        case "foo bar":
                                            var v95 bool = false
                                            _ = v95
                                        }
                                    }
                                }
                            }
                        case 747:
                            var v96 int32 = (v2 + v2)
                            _ = v96
                        case -953:
                            if (false && (2 != 1)) {
                                v32 = (v32 * v32)
                                var v97 int32 = (v2 & 2147483647)
                                _ = v97
                            } else {
                                for fi98 := 0; fi98 < 1; fi98++ {
                                    v99 := GenS9{F0: (v3 - (v33 / ((v3) | 1)))}
                                    v100 := GenS9{F0: (v33 | (v33 / ((v3) | 1)))}
                                    _ = fi98
                                    _ = v100
                                    _ = v99
                                }
                            }
                        }
                    default:
                        v33 = 65535
                    }
                    var v101 int64 = 9223372036854775807
                    _ = v101
                }
            }
        default:
            fmt.Println(v2)
        }
        var v102 uint64 = 0
        _ = v102
        _ = v33
    } else {
        v2 = (v2 * v2)
    }
    v13 += (v13 - (v13 % (((v13 / (((v13 * (v13 - v13))) | 1))) | 1)))
    for fi103 := 0; fi103 < 9; fi103++ {
        for fi104 := 0; fi104 < 5; fi104++ {
            v1 += (v1 ^ -9223372036854775808)
            if len(v14) > 0 {
                v14[0] = (^(fi103))
            }
            var v105 int64 = (v1 & v1)
            _ = fi104
            _ = v105
        }
        _ = fi103
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
}
