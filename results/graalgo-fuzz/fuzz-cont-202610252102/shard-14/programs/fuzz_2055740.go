package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 uint; F1 int64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 int16; F1 bool }
func genFn0() string {
    var v7 string = "a"
    if false {
        v7 = v7
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 float64 = (1.5 / (0.1 + float64((-2.25 * -2.25))))
            _ = fi8
            _ = v9
        }
    }
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{2, 222}
    v3 := make(map[string]uint32)
    var v4 string = "\n\t"
    v5 := make(map[int]uint32)
    v6 := make(map[int]bool)
    var v10 int8 = 42
    if (!false) {
        v11 := GenS1{F0: 42, F1: (v4 + v4), F2: 2}
        _ = v11
    }
    if false {
        var v12 int8 = v10
        _ = v12
    } else {
        v10 = v10
        v10 -= v10
    }
    var v13 int = 42
    if false {
        var v14 bool = false
        if (1 == 42) {
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 int = 42
                v5[v16] = 878
                _ = fi15
                _ = v16
            }
        }
        v17 := GenS2{F0: v4, F1: 9223372036854775807, F2: 42}
        _ = v14
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        if (fi18 >= fi18) {
            if (!(127 <= 255)) {
                if len(v2) > 0 {
                    v2[0] = 573
                }
                v4 = v4
            }
            var v19 int16 = 954
            var v20 uint64 = 110
            _ = v19
            _ = v20
        } else {
            var v21 uint16 = 898
            _ = v21
        }
        _ = fi18
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        if (!false) {
            v4 = "foo bar"
            v23 := GenS3{F0: 891}
            if (!((42 == 412) || (1 >= 1))) {
                v24 := GenS4{F0: 782, F1: v13, F2: 335}
                for fi25 := 0; fi25 < 1; fi25++ {
                    var v26 float64 = 0.1
                    if true {
                        v27 := GenS1{F0: 724, F1: "hello", F2: 2147483648}
                        v13 = 2147483647
                        v28 := GenS2{F0: "foo bar", F1: -1, F2: 255}
                        _ = v27
                        _ = v28
                    }
                    var v29 bool = false
                    switch v4 {
                    case "\n\t":
                        v30 := GenS2{F0: v4, F1: 0, F2: 128}
                        _ = v30
                    default:
                        var v31 int8 = v10
                        _ = v31
                    }
                    if (2 != 0) {
                        v3[v4] = 4294967295
                    }
                    _ = fi25
                    _ = v26
                    _ = v29
                }
                _ = v24
            } else {
                v32 := GenS4{F0: int16(v10), F1: fi22, F2: 255}
                v33 := GenS1{F0: 2, F1: "\n\t", F2: 2}
                if true {
                    v34 := v5[v13]
                    if ((true || false) || (9223372036854775807 > v13)) {
                        var v35 int = fi22
                        _ = v35
                    } else {
                        v36 := GenS5{F0: 0, F1: 255, F2: 42}
                        v37 := GenS6{F0: 991, F1: "a"}
                        _ = v36
                        _ = v37
                    }
                    var v38 uint64 = 42
                    _ = v34
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v5[(fi22 % ((fi22) | 1))] = 1
                        var v40 uint32 = 2
                        v41 := GenS3{F0: (v40 << ((310) & 63))}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v42 int8 = (v10 - (v10 / ((v10) | 1)))
                        _ = fi39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                }
                _ = v32
                _ = v33
            }
            _ = v23
        }
        _ = fi22
    }
    v43 := GenS6{F0: 1, F1: v4}
    for fi44 := 0; fi44 < 3; fi44++ {
        var v45 int16 = 1000
        fmt.Println(fi44, v45)
        _ = fi44
        _ = v45
    }
    if true {
        v46 := GenS4{F0: -134, F1: (v13 / ((v13) | 1)), F2: 255}
        v47 := GenS5{F0: 255, F1: 255, F2: 119}
        _ = v46
        _ = v47
    } else {
        for fi48 := 0; fi48 < 14; fi48++ {
            v10 -= v10
            var v49 bool = false
            _ = fi48
            _ = v49
        }
        switch int(v10) {
        case 1:
            if true {
                if (-751 <= 1) {
                    v6[-345] = false
                    v50 := GenS2{F0: "hello", F1: 9223372036854775807, F2: 2}
                    _ = v50
                } else {
                    var v51 int8 = (v10 / ((-104) | 1))
                    v13 -= (v13 - (v13 - -2147483648))
                    var v52 int = (v13 & v13)
                    _ = v51
                    _ = v52
                }
                var v53 uint32 = 932
                switch (v4 + v4) {
                case "foo bar":
                    v4 = v4
                case "\n\t":
                    v10 = (v10 % ((-128) | 1))
                case "":
                    delete(v6, v13)
                }
                _ = v53
            } else {
                var v54 uint = 993
                if true {
                    var v55 int16 = -760
                    if (0.1 > 1e10) {
                        v10 = (-(v10))
                        if true {
                            var v56 uint16 = 2
                            v57 := GenS7{F0: 823}
                            fmt.Println(v10, v13, v54)
                            _ = v56
                            _ = v57
                        } else {
                            fmt.Println(v10)
                            if true {
                                var v58 int8 = v10
                                _ = v58
                            } else {
                                for fi59 := 0; fi59 < 2; fi59++ {
                                    v54 -= (v54 * v54)
                                    v13 += (fi59 + -774)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi59
                                }
                                v60 := GenS6{F0: 604, F1: v4}
                                var v61 uint = (v54 + 2)
                                _ = v60
                                _ = v61
                            }
                            if true {
                                for fi62 := 0; fi62 < 17; fi62++ {
                                    var v63 int = fi62
                                    if false {
                                        v63 = (v63 ^ -537)
                                        switch -814 {
                                        case 517:
                                            delete(v1, 352)
                                        case -902:
                                            switch v4 {
                                            case "\n\t":
                                                delete(v1, 745)
                                            case "hello":
                                                var v64 int64 = 9223372036854775807
                                                _ = v64
                                            default:
                                                switch fi62 {
                                                case -804:
                                                    fmt.Println(v4)
                                                case 128:
                                                    if false {
                                                        switch v4 {
                                                        case "\n\t":
                                                            var v65 uint32 = 2
                                                            _ = v65
                                                        case "":
                                                            var v66 uint16 = 32767
                                                            _ = v66
                                                        }
                                                        switch 367 {
                                                        case -424:
                                                            switch (v13 << ((65535) & 63)) {
                                                            case 4294967295:
                                                                fmt.Println(v54, v4)
                                                            case 554:
                                                                v55 = v55
                                                            case 270:
                                                                if true {
                                                                    v10 = 12
                                                                }
                                                            default:
                                                                for fi67 := 0; fi67 < 13; fi67++ {
                                                                    fmt.Println(v10, fi62, v4)
                                                                    var v68 int8 = 1
                                                                    v63 -= (v13 << ((v54) & 63))
                                                                    _ = fi67
                                                                    _ = v68
                                                                }
                                                            }
                                                        default:
                                                            var v69 uint8 = 0
                                                            _ = v69
                                                        }
                                                        if true {
                                                            fmt.Println(v54, v10, v55)
                                                        } else {
                                                            if true {
                                                                v70 := GenS6{F0: 142, F1: ""}
                                                                _ = v70
                                                            } else {
                                                                v4 = (((v4 + (((v4 + "\n\t") + ((v4 + v4) + ("foo bar" + (((v4 + ("a" + (v4 + "foo bar"))) + ("" + v4)) + "hello")))) + ("\n\t" + v4))) + v4) + "foo bar")
                                                            }
                                                            for fi71 := 0; fi71 < 7; fi71++ {
                                                                v55 += (v55 & (v55 / ((v55) | 1)))
                                                                _ = fi71
                                                            }
                                                        }
                                                    } else {
                                                        v72 := GenS8{F0: ("foo bar" + ("foo bar" + ("foo bar" + v4))), F1: 152}
                                                        _ = v72
                                                    }
                                                case 735:
                                                    if len(v2) > 0 {
                                                        v2[0] = 2147483647
                                                    }
                                                default:
                                                    v73 := GenS3{F0: 253}
                                                    _ = v73
                                                }
                                            }
                                        default:
                                            fmt.Println(v13, v55, v54)
                                        }
                                        v4 = ("hello" + v4)
                                    } else {
                                        var v74 float32 = 0.0
                                        var v75 uint = 65535
                                        _ = v74
                                        _ = v75
                                    }
                                    var v76 bool = false
                                    _ = fi62
                                    _ = v63
                                    _ = v76
                                }
                                for fi77 := 0; fi77 < 10; fi77++ {
                                    v54 = 255
                                    _ = fi77
                                }
                            } else {
                                var v78 uint64 = 1
                                if (false || (((1.5 - 3.14159) * 3.14159) >= (1.5 + 0.1))) {
                                    v79 := GenS9{F0: v54, F1: -564}
                                    var v80 uint64 = v78
                                    v81 := GenS10{F0: (v10 * (v10 ^ 76))}
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                }
                                _ = v78
                            }
                        }
                    }
                    fmt.Println(v54)
                    _ = v55
                } else {
                    for fi82 := 0; fi82 < 3; fi82++ {
                        switch (fi82 - (-(v13))) {
                        case 128:
                            if (-1 != 1) {
                                if ((3.14159 - (0.1 - (0.1 + (0.1 - 1.5)))) != (float32(3.14159) - 0.1)) {
                                    v10 -= 127
                                    switch "hello" {
                                    case "hello":
                                        var v83 int32 = -424
                                        _ = v83
                                    case "a":
                                        switch v13 {
                                        case 1:
                                            var v84 string = "a"
                                            _ = v84
                                        case 32767:
                                            var v85 uint64 = 4294967295
                                            _ = v85
                                        }
                                    case "foo bar":
                                        v4 = (v4 + v4)
                                    default:
                                        for fi86 := 0; fi86 < 7; fi86++ {
                                            var v87 int8 = (-(v10))
                                            v4 = v4
                                            switch (fi82 / ((1) | 1)) {
                                            case -1:
                                                switch int(v10) {
                                                case 648:
                                                    v10 += 100
                                                }
                                            }
                                            _ = fi86
                                            _ = v87
                                        }
                                    }
                                    for fi88 := 0; fi88 < 20; fi88++ {
                                        v89 := GenS11{F0: 128}
                                        switch "foo bar" {
                                        case "hello":
                                            if ((1 >= 31) || false) {
                                                v90 := GenS3{F0: 394}
                                                fmt.Println(v13)
                                                var v91 uint64 = uint64(v54)
                                                v10 += v10
                                                _ = v90
                                                _ = v91
                                            } else {
                                                var v92 uint8 = 235
                                                v93 := GenS1{F0: (v54 * v54), F1: "a", F2: 937}
                                                var v94 bool = true
                                                switch int(v92) {
                                                case 708:
                                                    v95 := GenS2{F0: "a", F1: -471, F2: v92}
                                                    var v96 int16 = 1
                                                    _ = v95
                                                    _ = v96
                                                case 539:
                                                    v97 := GenS10{F0: (v10 + v10)}
                                                    _ = v97
                                                default:
                                                    v98 := GenS12{F0: ((v54 - v54) < 0), F1: (v54 != 380)}
                                                    _ = v98
                                                }
                                                var v99 string = "\n\t"
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                                _ = v99
                                            }
                                        default:
                                            if (873 != 1) {
                                                v13 -= 2147483647
                                                var v100 uint = (v54 % (((v54 & v54)) | 1))
                                                if false {
                                                    if ((!true) || (true || (false && false))) {
                                                        _ = len(v5)
                                                    }
                                                    var v101 uint32 = 0
                                                    _ = len(v6)
                                                    _ = v101
                                                } else {
                                                    var v102 uint8 = 1
                                                    _ = v102
                                                }
                                                _ = v100
                                            } else {
                                                v103 := GenS13{F0: 1000, F1: false}
                                                v104 := GenS6{F0: 1, F1: (v4 + (("hello" + "a") + v4))}
                                                var v105 bool = false
                                                _ = v103
                                                _ = v104
                                                _ = v105
                                            }
                                        }
                                        fmt.Println(v13, v4)
                                        _ = fi88
                                        _ = v89
                                    }
                                } else {
                                    v106 := GenS13{F0: -1, F1: false}
                                    _ = len(v2)
                                    _ = v106
                                }
                                v107 := GenS8{F0: v4, F1: 451}
                                _ = v107
                            } else {
                                v13 -= (v13 % ((v13) | 1))
                                fmt.Println(fi82)
                            }
                        case 395:
                            fmt.Println(v4, v10)
                        default:
                            switch 1 {
                            case -989:
                                switch (fi82 - -9223372036854775808) {
                                case -797:
                                    v54 = (v54 & (v54 >> ((408) & 63)))
                                }
                            case 2147483647:
                                _ = len(v1)
                            }
                        }
                        _ = fi82
                    }
                }
                _ = v54
            }
        case -670:
            v10 += v10
        case -1:
            var v108 int8 = (v10 << ((556) & 63))
            _ = v108
        }
    }
    switch -659 {
    case 42:
        _ = len(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
