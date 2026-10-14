package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int32 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS10 struct { F0 uint16; F1 int; F2 uint8 }
type GenS11 struct { F0 int16 }

func main() {
    var v1 int32 = 649
    var v2 float64 = ((0.1 / ((1.5 - 3.14159) * 3.14159)) / 0.1)
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 244, F1: 42, F2: 1}
    switch "a" {
    case "":
        if false {
            v6 := []int{-543}
            if false {
                var v7 uint = 658
                _ = v7
            }
            _ = v6
        }
    default:
        if (false || false) {
            if false {
                if (4294967295 == 2147483647) {
                    v1 -= v1
                } else {
                    fmt.Println(v2)
                }
                switch "a" {
                case "\n\t":
                    fmt.Println(v1, v1)
                case "hello":
                    if true {
                        var v8 uint = 4294967295
                        _ = v8
                    }
                case "foo bar":
                    var v9 int = -2147483648
                    _ = v9
                default:
                    if (42 < 255) {
                        var v10 uint8 = 1
                        fmt.Println(v1)
                        _ = v10
                    } else {
                        if (!(51 == 127)) {
                            fmt.Println(v1, v2)
                            delete(v4, 4294967295)
                        } else {
                            if (!false) {
                                var v11 int32 = 1
                                for fi12 := 0; fi12 < 19; fi12++ {
                                    if (((1.5 + 1.5) + 1.5) < 1.5) {
                                        v13 := GenS1{F0: 255, F1: fi12, F2: (v11 >> ((0) & 63))}
                                        for fi14 := 0; fi14 < 15; fi14++ {
                                            v2 = (((1.5 - -2.25) + v2) - (float64(-2.25) / ((1.5 - v2) / 0.1)))
                                            var v15 int8 = 42
                                            fmt.Println(v15, fi12, v2)
                                            _ = fi14
                                            _ = v15
                                        }
                                        _ = v13
                                    } else {
                                        switch fi12 {
                                        case 597:
                                            for fi16 := 0; fi16 < 14; fi16++ {
                                                v17 := []int{fi12, fi16, (fi12 ^ fi16), fi16}
                                                v18 := GenS2{F0: 0, F1: -715}
                                                _ = fi16
                                                _ = v17
                                                _ = v18
                                            }
                                        default:
                                            v19 := []int{}
                                            _ = v19
                                        }
                                        if (1e10 <= (0.1 * 1.5)) {
                                            v20 := GenS3{F0: int16(fi12)}
                                            for fi21 := 0; fi21 < 3; fi21++ {
                                                switch (fi21 >> ((274) & 63)) {
                                                case 0:
                                                    if false {
                                                        var v22 int = -9223372036854775808
                                                        if ((-(fi12)) >= fi21) {
                                                            fmt.Println(v2, fi12)
                                                            v23 := GenS1{F0: 1, F1: (fi12 % ((793) | 1)), F2: v1}
                                                            _ = v23
                                                        }
                                                        _ = v22
                                                    }
                                                case -2147483648:
                                                    v24 := GenS4{F0: 1}
                                                    _ = v24
                                                }
                                                v25 := GenS5{F0: "hello", F1: -542}
                                                var v26 float64 = v2
                                                _ = fi21
                                                _ = v25
                                                _ = v26
                                            }
                                            _ = v20
                                        } else {
                                            v1 = 114
                                        }
                                        v11 = v1
                                    }
                                    v27 := []int{-2147483648, fi12}
                                    _ = fi12
                                    _ = v27
                                }
                                _ = len(v4)
                                _ = v11
                            } else {
                                v28 := GenS6{F0: 2, F1: -9223372036854775808}
                                _ = v28
                            }
                        }
                    }
                }
            } else {
                if (!(32767 != -680)) {
                    v29 := GenS1{F0: 107, F1: 806, F2: (v1 + (v1 << ((42) & 63)))}
                    for fi30 := 0; fi30 < 8; fi30++ {
                        for fi31 := 0; fi31 < 17; fi31++ {
                            switch ("hello" + "") {
                            case "a":
                                v2 -= v2
                            case "foo bar":
                                if true {
                                    if true {
                                        fmt.Println(v1)
                                        fmt.Println(v1, v2)
                                        v1 = v1
                                    }
                                } else {
                                    var v32 float64 = v2
                                    v33 := GenS1{F0: 1, F1: 790, F2: (v1 & v1)}
                                    v32 -= v2
                                    var v34 int16 = 32767
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            case "hello":
                                v35 := v4[756]
                                _ = v35
                            }
                            if (!(42 >= 503)) {
                                if false {
                                    var v36 uint16 = 677
                                    _ = v36
                                }
                            }
                            _ = fi31
                        }
                        _ = fi30
                    }
                    _ = v29
                } else {
                    v2 -= 3.14159
                    v2 += float64(1.5)
                    var v37 int8 = -1
                    _ = v37
                }
            }
        }
    }
    if ((65535 > 65535) && false) {
        if (4 == 978) {
            var v38 float32 = (0.0 - 1.5)
            _ = v38
        }
    } else {
        for fi39 := 0; fi39 < 3; fi39++ {
            v2 -= (float64(0.1) + ((v2 * v2) - v2))
            switch fi39 {
            case 127:
                v40 := GenS1{F0: 168, F1: (fi39 + fi39), F2: v1}
                _ = v40
            case -2147483648:
                var v41 uint16 = 65535
                _ = v41
            }
            _ = fi39
        }
        switch "a" {
        case "a":
            v1 = 0
        case "\n\t":
            if false {
                var v42 uint = 0
                _ = v42
            } else {
                for fi43 := 0; fi43 < 11; fi43++ {
                    if true {
                        var v44 int = fi43
                        v45 := []int{}
                        fmt.Println(v2)
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint8 = 213
                        var v47 int64 = 295
                        v47 += v47
                        _ = v46
                        _ = v47
                    }
                    _ = fi43
                }
                fmt.Println(v2)
            }
        case "foo bar":
            v2 = ((((float64(1e10) * v2) / (((v2 + (0.1 + v2)) - v2) - v2)) + (1e10 + -2.25)) * ((float64(1.5) * v2) + v2))
        default:
            switch "hello" {
            case "hello":
                v2 += (v2 * v2)
            case "foo bar":
                var v48 int64 = -500
                _ = v48
            case "\n\t":
                for fi49 := 0; fi49 < 1; fi49++ {
                    var v50 float32 = -2.25
                    v50 = 3.14159
                    v51 := []int{fi49}
                    _ = fi49
                    _ = v50
                    _ = v51
                }
            }
        }
        switch -244 {
        case 191:
            fmt.Println(v2, v2)
        }
    }
    if false {
        var v52 uint32 = 108
        delete(v4, v52)
        for fi53 := 0; fi53 < 7; fi53++ {
            var v54 uint16 = 617
            v55 := []int{fi53, 24}
            v52 = (v52 * (v52 ^ (v52 << ((0) & 63))))
            _ = fi53
            _ = v54
            _ = v55
        }
        _ = v52
    } else {
        fmt.Println(v2, v2)
    }
    var v56 string = (("foo bar" + "a") + "")
    fmt.Println(v2)
    v57 := GenS6{F0: 149, F1: -1}
    for fi58 := 0; fi58 < 13; fi58++ {
        v56 = ""
        v56 = (("foo bar" + (("" + (v56 + ("hello" + "foo bar"))) + ("\n\t" + ((((v56 + v56) + v56) + "") + "hello")))) + (v56 + v56))
        _ = fi58
    }
    var v59 bool = false
    var v60 float64 = -2.25
    for fi61 := 0; fi61 < 8; fi61++ {
        switch v56 {
        case "foo bar":
            for fi62 := 0; fi62 < 17; fi62++ {
                if ((float64(0.1) / 3.14159) < 1e10) {
                    v63 := GenS7{F0: v56, F1: 2}
                    var v64 uint16 = 0
                    _ = v63
                    _ = v64
                }
                delete(v3, (fi61 / (((fi62 + fi62)) | 1)))
                switch "" {
                case "":
                    var v65 uint32 = 2147483647
                    _ = v65
                }
                var v66 uint16 = 32767
                _ = fi62
                _ = v66
            }
        case "":
            var v67 uint64 = 4294967295
            _ = v67
        }
        _ = fi61
    }
    if (-604 >= 32767) {
        if (true && ((v2 == v60) && v59)) {
            fmt.Println(v2, v2)
            v68 := v4[42]
            _ = v68
        } else {
            for fi69 := 0; fi69 < 18; fi69++ {
                v59 = v59
                v70 := GenS2{F0: 499, F1: (v1 + (v1 << ((554) & 63)))}
                _ = fi69
                _ = v70
            }
        }
        if true {
            v71 := GenS8{F0: 128, F1: (v56 + "\n\t")}
            _ = v71
        } else {
            if v59 {
                for fi72 := 0; fi72 < 13; fi72++ {
                    if (int64(fi72) > -9223372036854775808) {
                        switch fi72 {
                        case -44:
                            v73 := GenS8{F0: 1, F1: v56}
                            _ = v73
                        default:
                            var v74 float64 = v2
                            _ = v74
                        }
                        for fi75 := 0; fi75 < 3; fi75++ {
                            v4[2147483648] = fi75
                            v76 := GenS2{F0: 65535, F1: (v1 >> ((1) & 63))}
                            delete(v3, -1)
                            _ = fi75
                            _ = v76
                        }
                    }
                    if true {
                        switch "" {
                        case "hello":
                            v59 = (23 != 651)
                        default:
                            for fi77 := 0; fi77 < 5; fi77++ {
                                if (true && (1 < 42)) {
                                    var v78 float32 = 0.1
                                    if v59 {
                                        switch (fi77 ^ -1) {
                                        case -9223372036854775808:
                                            v79 := []int{(-(fi72))}
                                            _ = v79
                                        default:
                                            if (!v59) {
                                                v80 := GenS9{F0: 576, F1: 1, F2: 255}
                                                _ = v80
                                            } else {
                                                var v81 uint64 = 42
                                                var v82 float32 = 0.0
                                                _ = v81
                                                _ = v82
                                            }
                                        }
                                        _ = len(v4)
                                        v83 := []int{(^(fi72)), (-(fi77)), 128}
                                        _ = v83
                                    } else {
                                        switch (fi77 - 128) {
                                        case -505:
                                            v84 := v3[(^(fi77))]
                                            _ = v84
                                        case 4294967295:
                                            if v59 {
                                                switch fi77 {
                                                case 366:
                                                    v78 -= v78
                                                }
                                                v1 = (-(v1))
                                                switch 32768 {
                                                case 9223372036854775807:
                                                    if (v59 || (-19 != 836)) {
                                                        delete(v4, 4294967295)
                                                        v85 := GenS8{F0: 130, F1: (("foo bar" + v56) + "\n\t")}
                                                        fmt.Println(v56, v2, v1)
                                                        _ = v85
                                                    } else {
                                                        fmt.Println(v60, v60)
                                                        v78 -= 0.1
                                                        var v86 uint8 = 255
                                                        _ = v86
                                                    }
                                                }
                                            }
                                        default:
                                            switch (-(fi77)) {
                                            case 128:
                                                if false {
                                                    var v87 int16 = 1000
                                                    if (-2.25 > 3.14159) {
                                                        v88 := []int{fi77, fi77, fi72, 0}
                                                        v60 += float64(float32(1.5))
                                                        _ = v88
                                                    } else {
                                                        var v89 int32 = -1
                                                        if (2 <= 42) {
                                                            v1 += (v1 << ((2) & 63))
                                                            v89 -= v89
                                                        }
                                                        v90 := GenS10{F0: 366, F1: fi77, F2: 128}
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                    _ = v87
                                                }
                                            case -512:
                                                fmt.Println(v60, v78)
                                            case 854:
                                                fmt.Println(v56, v78, fi77)
                                            default:
                                                var v91 uint = 2
                                                _ = v91
                                            }
                                        }
                                    }
                                    v2 = float64((v78 * v78))
                                    _ = v78
                                } else {
                                    v60 += v2
                                    if (65535 != 18446744073709551615) {
                                        v56 = (v56 + v56)
                                    }
                                }
                                _ = fi77
                            }
                        }
                        var v92 bool = false
                        _ = v92
                    }
                    v93 := GenS7{F0: v56, F1: 2}
                    _ = fi72
                    _ = v93
                }
                var v94 int16 = 0
                _ = v94
            }
            switch -1 {
            case 2147483647:
                v60 += v60
            case 4294967295:
                v95 := GenS5{F0: v56, F1: -489}
                _ = v95
            case 0:
                v59 = false
            }
            for fi96 := 0; fi96 < 3; fi96++ {
                fmt.Println(v59, v1)
                v59 = v59
                switch v56 {
                case "\n\t":
                    for fi97 := 0; fi97 < 4; fi97++ {
                        switch (-(fi96)) {
                        case 92:
                            delete(v3, fi97)
                        case 32767:
                            v98 := GenS11{F0: 1}
                            _ = v98
                        default:
                            var v99 int32 = v1
                            _ = v99
                        }
                        if (v56 == v56) {
                            fmt.Println(fi97, v2)
                        }
                        _ = fi97
                    }
                case "":
                    var v100 uint16 = 937
                    _ = v100
                }
                _ = fi96
            }
        }
    } else {
        v101 := GenS10{F0: 255, F1: 42, F2: 233}
        var v102 int = -883
        _ = v101
        _ = v102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v59
    _ = v60
}
