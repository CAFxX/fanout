package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 int32; F2 uint }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 uint; F1 int16 }
type GenS10 struct { F0 bool }

func main() {
    var v1 uint = 42
    var v2 string = ("foo bar" + "")
    var v3 int = 84
    if (!false) {
        v4 := []int{}
        for fi5 := 0; fi5 < 5; fi5++ {
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := GenS1{F0: "\n\t"}
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        var v8 int64 = -69
        _ = v4
        _ = v8
    } else {
        v9 := GenS2{F0: false, F1: v3}
        var v10 float64 = 3.14159
        v10 -= 0.1
        _ = v10
        _ = v9
    }
    var v11 uint = v1
    if (128 > (-(v3))) {
        v12 := []int{0}
        var v13 int8 = 1
        _ = v12
        _ = v13
    } else {
        v3 = v3
    }
    if (!(1 < 4294967295)) {
        v2 = (v2 + v2)
        switch "\n\t" {
        case "":
            v14 := []int{}
            _ = v14
        case "foo bar":
            fmt.Println(v3)
        default:
            if (!(1 == 238)) {
                fmt.Println(v3)
            }
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            if false {
                v11 -= (v1 / ((v11) | 1))
                if true {
                    var v16 int32 = 0
                    v1 = (v1 - v11)
                    fmt.Println(v1)
                    _ = v16
                }
                v1 = (v11 >> ((18446744073709551615) & 63))
            } else {
                if ("\n\t" != "hello") {
                    if (-95 >= 2147483647) {
                        v2 = v2
                        switch fi15 {
                        case 128:
                            v17 := []int{fi15, fi15, v3}
                            _ = v17
                        }
                        v3 = 127
                    } else {
                        if true {
                            for fi18 := 0; fi18 < 17; fi18++ {
                                if (!true) {
                                    switch v3 {
                                    case 0:
                                        v19 := []int{fi15, (fi18 * (v3 & (v3 - (v3 << ((v11) & 63))))), (v3 % ((fi15) | 1)), -1}
                                        _ = v19
                                    case -9223372036854775808:
                                        var v20 uint8 = 0
                                        _ = v20
                                    case -2147483648:
                                        v11 -= 18446744073709551615
                                    }
                                    var v21 int = 4294967295
                                    var v22 uint = 1
                                    _ = v21
                                    _ = v22
                                }
                                fmt.Println(v3, v11, fi15)
                                _ = fi18
                            }
                            for fi23 := 0; fi23 < 15; fi23++ {
                                fmt.Println(fi23, v2, v11)
                                var v24 uint16 = 1
                                _ = fi23
                                _ = v24
                            }
                        } else {
                            if false {
                                v25 := []int{fi15, v3, v3}
                                if (42 <= 0) {
                                    v26 := GenS1{F0: v2}
                                    _ = v26
                                }
                                _ = v25
                            }
                        }
                        for fi27 := 0; fi27 < 20; fi27++ {
                            fmt.Println(v1, v3, fi15)
                            _ = fi27
                        }
                        if (!true) {
                            v28 := GenS3{F0: (v1 % (((^(v1))) | 1)), F1: 0}
                            v2 = v2
                            _ = v28
                        } else {
                            v29 := GenS4{F0: ("foo bar" + v2)}
                            for fi30 := 0; fi30 < 15; fi30++ {
                                v31 := GenS5{F0: 2, F1: v2}
                                var v32 uint = (v1 >> ((v11) & 63))
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            if ((42 > 1) && ((1 > 621) && false)) {
                                switch (v3 - 1) {
                                case -1:
                                    var v33 int8 = 100
                                    _ = v33
                                case 1:
                                    if false {
                                        v34 := GenS6{F0: v11}
                                        _ = v34
                                    } else {
                                        var v35 uint8 = 0
                                        v36 := GenS5{F0: 2, F1: "\n\t"}
                                        _ = v35
                                        _ = v36
                                    }
                                case -483:
                                    switch v2 {
                                    case "hello":
                                        var v37 int64 = 1
                                        _ = v37
                                    case "":
                                        v11 = (v11 << (((v1 + 42)) & 63))
                                    }
                                }
                                var v38 int8 = -49
                                _ = v38
                            } else {
                                switch -1 {
                                case 127:
                                    var v39 int32 = 42
                                    _ = v39
                                }
                                var v40 uint32 = 2
                                switch fi15 {
                                case -770:
                                    if false {
                                        for fi41 := 0; fi41 < 17; fi41++ {
                                            switch v2 {
                                            case "foo bar":
                                                var v42 int8 = 29
                                                v42 += v42
                                                fmt.Println(v40)
                                                if (3.14159 >= 1.5) {
                                                    if (true || (v3 == fi15)) {
                                                        v42 -= v42
                                                        fmt.Println(v11)
                                                        var v43 float64 = ((3.14159 * 1e10) / 1.5)
                                                        _ = v43
                                                    }
                                                    v44 := GenS7{F0: (v40 ^ (^(v40))), F1: 42, F2: v11}
                                                    v2 = "hello"
                                                    _ = v44
                                                }
                                                v45 := []int{4294967295, (fi41 - v3), 42}
                                                var v46 int8 = 42
                                                _ = v42
                                                _ = v45
                                                _ = v46
                                            case "hello":
                                                if false {
                                                    switch (v2 + (v2 + v2)) {
                                                    case "hello":
                                                        var v47 float64 = ((0.0 * (((((((1e10 + 3.14159) * 1e10) - 0.0) / 0.1) / 3.14159) * (0.0 + 0.0)) - (1.5 / 1e10))) / -2.25)
                                                        _ = v47
                                                    case "":
                                                        var v48 uint = (v11 & v1)
                                                        _ = v48
                                                    default:
                                                        v49 := []int{}
                                                        _ = v49
                                                    }
                                                    switch v2 {
                                                    case "hello":
                                                        switch fi41 {
                                                        case 32767:
                                                            if false {
                                                                v50 := []int{fi41, fi15, fi41, (fi41 / ((v3) | 1))}
                                                                var v51 uint = 4294967295
                                                                var v52 int = 32768
                                                                _ = v50
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                v11 += 65535
                                                                v53 := GenS2{F0: (v40 < 339), F1: fi41}
                                                                _ = v53
                                                            }
                                                        case -2147483648:
                                                            v54 := GenS7{F0: (^(v40)), F1: 0, F2: (v1 >> ((255) & 63))}
                                                            v1 = 0
                                                            var v55 uint16 = 1
                                                            _ = v54
                                                            _ = v55
                                                        case 0:
                                                            if (!(!false)) {
                                                                v2 = (v2 + ("\n\t" + v2))
                                                                v11 = (v1 - v1)
                                                            } else {
                                                                v11 += (v11 % (((v1 & (v1 % (((v1 << ((v1) & 63))) | 1)))) | 1))
                                                                var v56 int16 = 649
                                                                _ = v56
                                                            }
                                                        default:
                                                            fmt.Println(v3, v3)
                                                        }
                                                    }
                                                    v57 := []int{(fi15 % (((fi41 - fi15)) | 1)), -2147483648, (v3 * 193)}
                                                    _ = v57
                                                } else {
                                                    v58 := GenS1{F0: ("a" + v2)}
                                                    _ = v58
                                                }
                                            }
                                            v59 := []int{(^(fi41))}
                                            _ = fi41
                                            _ = v59
                                        }
                                    } else {
                                        v11 += 2
                                        var v60 uint16 = 32767
                                        _ = v60
                                    }
                                case -746:
                                    var v61 int = -571
                                    _ = v61
                                case 127:
                                    v3 += 4294967295
                                }
                                _ = v40
                            }
                            _ = v29
                        }
                    }
                    v2 = v2
                    if (int32(v11) != 0) {
                        if true {
                            for fi62 := 0; fi62 < 13; fi62++ {
                                var v63 uint8 = 1
                                switch (v3 - v3) {
                                case 714:
                                    if true {
                                        var v64 bool = (true || true)
                                        _ = v64
                                    }
                                    v65 := []int{}
                                    switch fi15 {
                                    case 712:
                                        v66 := GenS2{F0: false, F1: (v3 + (v3 * (fi62 - (fi15 + 32767))))}
                                        _ = v66
                                    case -201:
                                        v63 = v63
                                    case 42:
                                        if ((1.5 <= (float32((((-2.25 / 1e10) + (0.1 * (-2.25 - ((((1.5 * 0.0) - -2.25) * 0.1) / 3.14159)))) * (0.0 / (-2.25 * (1e10 - 0.0))))) * 3.14159)) && (-831 < 856)) {
                                            switch fi62 {
                                            case -2147483648:
                                                v65 = append(v65, 4294967295)
                                            case -530:
                                                var v67 int8 = -108
                                                _ = v67
                                            default:
                                                var v68 float32 = 3.14159
                                                _ = v68
                                            }
                                        }
                                    }
                                    _ = v65
                                default:
                                    var v69 uint = 65535
                                    _ = v69
                                }
                                _ = fi62
                                _ = v63
                            }
                            v70 := GenS8{F0: "foo bar", F1: false}
                            _ = v70
                        }
                        if (!true) {
                            fmt.Println(v3)
                            for fi71 := 0; fi71 < 1; fi71++ {
                                if (255 <= 922) {
                                    v72 := GenS6{F0: (v11 * (^(v11)))}
                                    v11 = v11
                                    v3 += -6
                                    var v73 bool = (!((v2 + "foo bar") != v2))
                                    _ = v72
                                    _ = v73
                                }
                                v74 := []int{(v3 + 1), v3}
                                _ = fi71
                                _ = v74
                            }
                        }
                        if ((18446744073709551615 == 67) || false) {
                            for fi75 := 0; fi75 < 2; fi75++ {
                                var v76 int = (v3 - (-(v3)))
                                if ((("\n\t" + ((("\n\t" + v2) + v2) + v2)) + "\n\t") >= (("" + v2) + v2)) {
                                    switch "a" {
                                    case "hello":
                                        v2 = (("a" + v2) + v2)
                                    case "\n\t":
                                        switch v3 {
                                        case 351:
                                            v76 += -9223372036854775808
                                        case 0:
                                            fmt.Println(fi75, fi75, v76)
                                        }
                                    case "a":
                                        v77 := []int{4294967295}
                                        _ = v77
                                    }
                                    v78 := []int{v3, v3}
                                    fmt.Println(v76, v11)
                                    _ = v78
                                }
                                switch "" {
                                case "a":
                                    var v79 uint32 = 2147483647
                                    _ = v79
                                }
                                if true {
                                    v2 = v2
                                    v3 -= (v3 + v3)
                                    fmt.Println(v11)
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        switch "hello" {
                                        case "a":
                                            v3 = (v76 % (((fi15 | (fi75 / (((^(v3))) | 1)))) | 1))
                                        case "\n\t":
                                            v76 = fi75
                                        }
                                        v80 := []int{(fi15 + fi75), -657}
                                        v81 := GenS9{F0: v11, F1: -784}
                                        _ = v80
                                        _ = v81
                                    default:
                                        var v82 string = "a"
                                        _ = v82
                                    }
                                    if false {
                                        switch -103 {
                                        case -190:
                                            v3 = v3
                                        default:
                                            fmt.Println(fi75, fi75, v11)
                                        }
                                    }
                                }
                                var v83 float64 = 0.1
                                _ = fi75
                                _ = v76
                                _ = v83
                            }
                            v84 := GenS10{F0: (!true)}
                            _ = v84
                        }
                    } else {
                        var v85 int16 = -1
                        var v86 uint = v11
                        _ = v85
                        _ = v86
                    }
                }
            }
            _ = fi15
        }
    }
    var v87 uint = v1
    var v88 int64 = 42
    for fi89 := 0; fi89 < 20; fi89++ {
        for fi90 := 0; fi90 < 2; fi90++ {
            if (!(float64((-2.25 - 3.14159)) < 1.5)) {
                fmt.Println(fi90)
            }
            _ = fi90
        }
        _ = fi89
    }
    for fi91 := 0; fi91 < 1; fi91++ {
        v3 = (v3 + v3)
        var v92 uint32 = 0
        var v93 float32 = -2.25
        _ = fi91
        _ = v92
        _ = v93
    }
    if false {
        v94 := []int{(-(v3)), 9223372036854775807, (v3 ^ (v3 + v3))}
        var v95 int8 = -30
        for fi96 := 0; fi96 < 10; fi96++ {
            if (234 > v3) {
                v1 -= v1
                var v97 uint64 = 919
                _ = v97
            }
            for fi98 := 0; fi98 < 4; fi98++ {
                if (fi98 >= fi96) {
                    fmt.Println(v3, v3)
                    v99 := GenS5{F0: 255, F1: v2}
                    _ = v99
                }
                v2 = v2
                v95 = (v95 ^ (v95 ^ v95))
                _ = fi98
            }
            _ = fi96
        }
        _ = v94
        _ = v95
    } else {
        var v100 string = "hello"
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v87
    _ = v88
}
