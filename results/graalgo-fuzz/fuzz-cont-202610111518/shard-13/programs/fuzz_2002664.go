package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 int64; F2 bool }
type GenS7 struct { F0 int64 }

func main() {
    var v1 uint8 = 127
    v2 := []int{-811, 127, 48}
    var v3 uint = 17
    if false {
        switch -9223372036854775808 {
        case 855:
            if (false || true) {
                for fi4 := 0; fi4 < 14; fi4++ {
                    v1 -= v1
                    for fi5 := 0; fi5 < 2; fi5++ {
                        var v6 float32 = 3.14159
                        _ = fi5
                        _ = v6
                    }
                    _ = fi4
                }
            } else {
                if true {
                    v1 -= 0
                    fmt.Println(v3)
                    for fi7 := 0; fi7 < 1; fi7++ {
                        v1 -= (^(v1))
                        if true {
                            switch (("hello" + ("hello" + "a")) + ("" + "\n\t")) {
                            case "a":
                                for fi8 := 0; fi8 < 2; fi8++ {
                                    var v9 float64 = 0.1
                                    v9 = 0.1
                                    _ = fi8
                                    _ = v9
                                }
                            }
                            var v10 uint = (v3 * (^(v3)))
                            v10 -= 255
                            _ = v10
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "a":
                                switch (fi7 | 32767) {
                                case -647:
                                    v3 += 0
                                default:
                                    for fi11 := 0; fi11 < 12; fi11++ {
                                        if (!("" >= "\n\t")) {
                                            fmt.Println(fi11)
                                            v12 := GenS1{F0: 255, F1: -1, F2: v1}
                                            if true {
                                                var v13 uint32 = 902
                                                v1 = (v1 % ((v1) | 1))
                                                _ = v13
                                            }
                                            _ = v12
                                        }
                                        _ = fi11
                                    }
                                }
                            default:
                                if ((!false) || true) {
                                    var v14 float64 = -2.25
                                    _ = v14
                                }
                            }
                            for fi15 := 0; fi15 < 10; fi15++ {
                                fmt.Println(v1)
                                v3 = 547
                                if false {
                                    if (261 != -2147483648) {
                                        v16 := GenS1{F0: 2, F1: -848, F2: (v1 << ((v3) & 63))}
                                        _ = v16
                                    } else {
                                        if (-66 <= 100) {
                                            fmt.Println(v3, fi15, v3)
                                            fmt.Println(fi7, fi7)
                                        } else {
                                            if true {
                                                v3 -= (v3 >> ((v3) & 63))
                                                v1 = v1
                                                fmt.Println(fi7, fi15)
                                            }
                                            fmt.Println(v1, fi15, v1)
                                        }
                                        var v17 uint8 = v1
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v17
                                    }
                                    var v18 uint8 = (v1 + (v1 * (v1 ^ v1)))
                                    fmt.Println(fi7)
                                    _ = v18
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = fi7
                                    }
                                    var v19 float32 = -2.25
                                    var v20 int32 = 2147483647
                                    _ = v19
                                    _ = v20
                                }
                                _ = fi15
                            }
                            fmt.Println(v1, fi7)
                        }
                        v21 := GenS2{F0: 199, F1: 123}
                        _ = fi7
                        _ = v21
                    }
                }
                var v22 uint32 = 2
                _ = v22
            }
        case -2147483648:
            var v23 string = (("a" + "hello") + "\n\t")
            _ = v23
        }
        if true {
            if false {
                var v24 uint = 0
                v1 += v1
                if (4294967295 != 635) {
                    if (!(false && false)) {
                        v24 -= 840
                        for fi25 := 0; fi25 < 10; fi25++ {
                            v26 := GenS3{F0: (v3 >> ((v24) & 63)), F1: 255, F2: 18446744073709551615}
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        v1 = (v1 & v1)
                    }
                } else {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        var v28 uint64 = 2
                        _ = fi27
                        _ = v28
                    }
                    if true {
                        v2 = append(v2, 1)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if ((-2.25 != 0.0) || (true && true)) {
                            if true {
                                var v29 float32 = 0.1
                                if len(v2) > 0 {
                                    v2[0] = 0
                                }
                                _ = v29
                            } else {
                                var v30 int = 4294967295
                                var v31 uint8 = v1
                                _ = v30
                                _ = v31
                            }
                        }
                    } else {
                        var v32 uint8 = v1
                        _ = v32
                    }
                }
                _ = v24
            }
            v1 = (v1 & v1)
        }
        v1 += 2
    }
    v3 = v3
    fmt.Println(v3, v1)
    v2 = append(v2, 32768)
    v33 := GenS4{F0: 1, F1: -83}
    v2 = append(v2, 2147483647)
    v34 := GenS5{F0: 31}
    v35 := GenS1{F0: 1, F1: 0, F2: v1}
    v1 += v1
    v3 -= v3
    v1 += v1
    v36 := GenS1{F0: 32767, F1: -264, F2: v1}
    v3 = v3
    for fi37 := 0; fi37 < 9; fi37++ {
        if false {
            v3 -= uint(v1)
            switch 2147483647 {
            case 878:
                var v38 string = ("foo bar" + "a")
                _ = v38
            case 4294967295:
                for fi39 := 0; fi39 < 4; fi39++ {
                    var v40 uint = (v3 << ((4294967295) & 63))
                    _ = fi39
                    _ = v40
                }
            default:
                if true {
                    switch (fi37 * -9223372036854775808) {
                    case 4294967295:
                        v1 = v1
                    case -2147483648:
                        v1 = v1
                    case 874:
                        var v41 bool = (!(fi37 < (-(fi37))))
                        _ = v41
                    default:
                        var v42 uint8 = v1
                        _ = v42
                    }
                    switch fi37 {
                    case 0:
                        fmt.Println(v1, v3)
                    case 1:
                        switch fi37 {
                        case -230:
                            if (("a" + ("" + "foo bar")) <= "a") {
                                if ((!false) || (745 == 0)) {
                                    for fi43 := 0; fi43 < 10; fi43++ {
                                        v44 := GenS6{F0: (v3 % (((v3 ^ 440)) | 1)), F1: 0, F2: ((((0.1 / float64((float32(float64(0.1)) / 3.14159))) / 1.0) + 0.0) >= 1e10)}
                                        v45 := GenS7{F0: 42}
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                    }
                                    v1 = v1
                                    switch (fi37 ^ fi37) {
                                    case 2147483647:
                                        v46 := GenS1{F0: 138, F1: -965, F2: (v1 / ((127) | 1))}
                                        _ = v46
                                    }
                                } else {
                                    for fi47 := 0; fi47 < 19; fi47++ {
                                        _ = len(v2)
                                        if false {
                                            if false {
                                                v1 += v1
                                            }
                                            if false {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(fi37, fi47)
                                                case "":
                                                    if (!true) {
                                                        _ = len(v2)
                                                    }
                                                case "a":
                                                    if false {
                                                        if (false && (!true)) {
                                                            var v48 float64 = 1e10
                                                            v48 -= v48
                                                            switch fi47 {
                                                            case 605:
                                                                v1 = (v1 % ((126) | 1))
                                                                var v49 uint32 = 1
                                                                _ = v49
                                                            default:
                                                                v50 := GenS6{F0: 2, F1: 997, F2: (!(!(((0 != 0) || false) || (true && (890 < 879)))))}
                                                                _ = v50
                                                            }
                                                            _ = v48
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    v51 := GenS6{F0: v3, F1: -94, F2: (42 >= (^(v1)))}
                                                                    _ = v51
                                                                } else {
                                                                    var v52 uint8 = (^(v1))
                                                                    _ = v52
                                                                }
                                                                v1 = (v1 ^ 2)
                                                                v1 += v1
                                                            } else {
                                                                var v53 string = "a"
                                                                var v54 string = "foo bar"
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            var v55 int32 = 2147483647
                                                            v2 = append(v2, fi47)
                                                            v2 = append(v2, (fi47 + fi47))
                                                            _ = v55
                                                        }
                                                        var v56 uint16 = 2
                                                        if ((v3 / (((v3 | uint(v1))) | 1)) >= 198) {
                                                            v56 = (v56 / ((v56) | 1))
                                                        } else {
                                                            v3 -= (v3 * v3)
                                                            v57 := GenS2{F0: (fi47 ^ (fi37 + (fi37 + -426))), F1: -128}
                                                            if false {
                                                                v1 = (v1 / ((42) | 1))
                                                                var v58 int = (-(fi47))
                                                                var v59 int16 = -778
                                                                switch fi37 {
                                                                case -326:
                                                                    if true {
                                                                        v3 = v3
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            if (("" + (("foo bar" + "\n\t") + "\n\t")) == "\n\t") {
                                                                                fmt.Println(v56)
                                                                            }
                                                                        case "":
                                                                            switch "hello" {
                                                                            case "":
                                                                                switch v58 {
                                                                                case 2147483647:
                                                                                    v1 -= v1
                                                                                case 82:
                                                                                    if (v1 >= 0) {
                                                                                        v2 = append(v2, 32767)
                                                                                    }
                                                                                }
                                                                            case "hello":
                                                                                var v60 int8 = 123
                                                                                _ = len(v2)
                                                                                _ = v60
                                                                            }
                                                                        case "a":
                                                                            switch (("hello" + "foo bar") + ("foo bar" + "\n\t")) {
                                                                            case "":
                                                                                v3 = (^(v3))
                                                                            }
                                                                            var v61 float32 = 1.5
                                                                            _ = v61
                                                                        }
                                                                        fmt.Println(v1)
                                                                    }
                                                                    var v62 uint64 = 368
                                                                    if (1.5 >= (1.5 + 0.0)) {
                                                                        var v63 string = "hello"
                                                                        if ((3.14159 >= 0.1) && false) {
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                if true {
                                                                                    fmt.Println(fi47)
                                                                                    v63 = v63
                                                                                    var v64 uint8 = 128
                                                                                    _ = v64
                                                                                } else {
                                                                                    var v65 int32 = 42
                                                                                    var v66 uint16 = (v56 - uint16(v3))
                                                                                    _ = v65
                                                                                    _ = v66
                                                                                }
                                                                            case "a":
                                                                                var v67 float32 = (0.0 - (float32(((1e10 * (((1.5 * (1e10 - -2.25)) - 0.0) / (float64(0.1) + -2.25))) * (((float64((1e10 * (1e10 * -2.25))) * (0.0 + 0.1)) - 1e10) - float64(1.5)))) * (-2.25 / 1.0)))
                                                                                _ = v67
                                                                            case "":
                                                                                var v68 int64 = 223
                                                                                _ = v68
                                                                            default:
                                                                                v69 := GenS3{F0: 18446744073709551615, F1: (v56 << (((^(v3))) & 63)), F2: (v62 * 18446744073709551615)}
                                                                                var v70 int32 = -1
                                                                                _ = v69
                                                                                _ = v70
                                                                            }
                                                                            v1 = 255
                                                                        } else {
                                                                            var v71 int = (v58 + int(v1))
                                                                            if (false && true) {
                                                                                v72 := GenS2{F0: fi47, F1: 42}
                                                                                v59 = v59
                                                                                if true {
                                                                                    v58 -= v58
                                                                                }
                                                                                var v73 string = v63
                                                                                _ = v72
                                                                                _ = v73
                                                                            } else {
                                                                                var v74 int = (fi37 * v71)
                                                                                var v75 uint32 = 245
                                                                                v62 = 0
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                            switch (v63 + (v63 + "foo bar")) {
                                                                            case "hello":
                                                                                var v76 uint16 = (v56 >> ((18446744073709551615) & 63))
                                                                                _ = v76
                                                                            case "foo bar":
                                                                                switch v63 {
                                                                                case "foo bar":
                                                                                    var v77 uint = v3
                                                                                    _ = v77
                                                                                case "":
                                                                                    _ = len(v2)
                                                                                default:
                                                                                    var v78 uint = 472
                                                                                    _ = v78
                                                                                }
                                                                                var v79 uint = v3
                                                                                _ = v79
                                                                            case "\n\t":
                                                                                fmt.Println(v56, v59)
                                                                            default:
                                                                                _ = len(v2)
                                                                            }
                                                                            _ = v71
                                                                        }
                                                                        _ = v63
                                                                    }
                                                                    _ = v62
                                                                case -9223372036854775808:
                                                                    v56 = 2
                                                                default:
                                                                    if false {
                                                                        var v80 int8 = -107
                                                                        _ = v80
                                                                    } else {
                                                                        var v81 uint = (v3 * v3)
                                                                        if (v59 >= (v59 << (((v81 << (((v3 ^ v81)) & 63))) & 63))) {
                                                                            fmt.Println(v1, fi47)
                                                                            if true {
                                                                                var v82 int32 = 0
                                                                                var v83 uint32 = 816
                                                                                var v84 uint32 = v83
                                                                                var v85 uint64 = 0
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v86 uint8 = v1
                                                                                _ = v82
                                                                                _ = v83
                                                                                _ = v84
                                                                                _ = v85
                                                                                _ = v86
                                                                            } else {
                                                                                v56 = v56
                                                                            }
                                                                            var v87 uint = uint(v1)
                                                                            _ = v87
                                                                        }
                                                                        _ = v81
                                                                    }
                                                                }
                                                                _ = v58
                                                                _ = v59
                                                            } else {
                                                                v88 := GenS4{F0: v3, F1: 1}
                                                                var v89 float64 = -2.25
                                                                if len(v2) > 0 {
                                                                    v2[0] = -9223372036854775808
                                                                }
                                                                switch ((("foo bar" + "a") + ("a" + ("foo bar" + "\n\t"))) + "hello") {
                                                                case "\n\t":
                                                                    v90 := GenS4{F0: uint(v56), F1: 1}
                                                                    _ = v90
                                                                default:
                                                                    v89 = (v89 + v89)
                                                                }
                                                                v91 := GenS6{F0: v3, F1: 0, F2: true}
                                                                _ = v88
                                                                _ = v89
                                                                _ = v91
                                                            }
                                                            _ = v57
                                                        }
                                                        _ = v56
                                                    }
                                                default:
                                                    var v92 int16 = -32768
                                                    _ = v92
                                                }
                                            }
                                        } else {
                                            var v93 uint16 = 857
                                            v2 = append(v2, (fi37 | fi37))
                                            v94 := GenS5{F0: 127}
                                            _ = v93
                                            _ = v94
                                        }
                                        _ = fi47
                                    }
                                    v3 += 255
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                v3 += (^(v3))
                            } else {
                                v3 = (v3 % ((4294967295) | 1))
                                fmt.Println(fi37)
                                fmt.Println(fi37)
                            }
                        case 488:
                            if false {
                                if false {
                                    var v95 uint = v3
                                    var v96 uint8 = (v1 ^ (v1 & v1))
                                    _ = v95
                                    _ = v96
                                }
                            } else {
                                v97 := GenS4{F0: (v3 + (v3 ^ 791)), F1: -1}
                                v3 = v3
                                _ = v97
                            }
                        }
                    default:
                        v3 = uint(v1)
                    }
                } else {
                    var v98 int32 = -1
                    if false {
                        _ = len(v2)
                        if (32767 > 812) {
                            v98 = v98
                            fmt.Println(v98, fi37, v3)
                        }
                        if true {
                            if ((-727 >= 42) || true) {
                                v98 = (^(v98))
                            } else {
                                v98 += (v98 * v98)
                                v99 := GenS1{F0: 42, F1: v98, F2: 221}
                                _ = v99
                            }
                        }
                    }
                    _ = v98
                }
            }
            v100 := GenS5{F0: 92}
            _ = v100
        }
        _ = fi37
    }
    if (!false) {
        fmt.Println(v1, v1)
        for fi101 := 0; fi101 < 12; fi101++ {
            for fi102 := 0; fi102 < 8; fi102++ {
                switch "a" {
                case "foo bar":
                    var v103 uint64 = 1
                    _ = v103
                case "":
                    v1 = v1
                }
                _ = fi102
            }
            var v104 int8 = 127
            _ = fi101
            _ = v104
        }
        v105 := GenS4{F0: v3, F1: 146}
        _ = v105
    } else {
        if (!(false && true)) {
            for fi106 := 0; fi106 < 4; fi106++ {
                var v107 uint = (^(v3))
                if false {
                    v3 = 18446744073709551615
                    var v108 bool = true
                    _ = v108
                }
                v107 = 18446744073709551615
                _ = fi106
                _ = v107
            }
        }
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        var v109 int16 = -342
        _ = v109
    case "hello":
        v1 -= (^(v1))
    case "foo bar":
        v3 += v3
    }
    v3 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
