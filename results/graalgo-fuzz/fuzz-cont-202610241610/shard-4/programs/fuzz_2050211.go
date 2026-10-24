package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint32; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int16; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS8 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{85, -9223372036854775808, -922}
    v3 := make(map[int]string)
    if false {
        v4 := GenS1{F0: 418, F1: 1}
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v5 int = 42
            _ = v5
        case -1:
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 int16 = -32768
                for fi8 := 0; fi8 < 4; fi8++ {
                    v7 += 42
                    var v9 uint64 = 0
                    _ = fi8
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        case -2147483648:
            if ((true || (4294967295 > 42)) && (((0.1 / 1.5) * 1.5) <= ((-2.25 / float64(3.14159)) * 0.0))) {
                if true {
                    switch -1 {
                    case 2147483647:
                        v10 := GenS2{F0: 823, F1: 2, F2: "a"}
                        _ = v10
                    default:
                        for fi11 := 0; fi11 < 3; fi11++ {
                            v3[128] = "foo bar"
                            var v12 uint16 = 363
                            if (32767 >= 42) {
                                var v13 string = "\n\t"
                                _ = v13
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                    switch "hello" {
                    case "foo bar":
                        var v14 int16 = -32768
                        _ = v14
                    case "hello":
                        for fi15 := 0; fi15 < 11; fi15++ {
                            v16 := GenS1{F0: 1, F1: 42}
                            var v17 float32 = float32((1.5 + 0.1))
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case "\n\t":
                        var v18 uint = 18446744073709551615
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 19; fi19++ {
                            var v20 uint = 65535
                            _ = fi19
                            _ = v20
                        }
                    }
                } else {
                    var v21 uint = 0
                    _ = v21
                }
            } else {
                var v22 int8 = -1
                var v23 float32 = 1e10
                v24 := GenS1{F0: 758, F1: 65535}
                _ = v22
                _ = v23
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 3; fi25++ {
                if (1 < -1) {
                    var v26 int = (fi25 & (fi25 * fi25))
                    if true {
                        var v27 int16 = 1
                        _ = v27
                    }
                    _ = v26
                } else {
                    var v28 string = ""
                    var v29 uint8 = 1
                    if len(v2) > 0 {
                        v2[0] = (fi25 - (fi25 - 0))
                    }
                    _ = v28
                    _ = v29
                }
                _ = fi25
            }
        }
        if false {
            switch 9223372036854775807 {
            case 312:
                switch "a" {
                case "\n\t":
                    for fi30 := 0; fi30 < 9; fi30++ {
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v32 := GenS2{F0: 4294967295, F1: 2, F2: "a"}
                            v33 := GenS1{F0: 2, F1: 411}
                            var v34 int64 = -12
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS1{F0: 32767, F1: 65}
                        var v36 string = "\n\t"
                        _ = fi30
                        _ = v35
                        _ = v36
                    }
                case "":
                    v2 = append(v2, 9223372036854775807)
                case "hello":
                    var v37 int64 = -9223372036854775808
                    _ = v37
                }
            case -969:
                if false {
                    var v38 int32 = 1
                    _ = v38
                } else {
                    if (0 == 520) {
                        var v39 int8 = 72
                        _ = v39
                    }
                }
            case 42:
                if (-1 <= -1) {
                    v40 := GenS1{F0: 1, F1: 2}
                    var v41 int64 = -132
                    v41 = v41
                    _ = v40
                    _ = v41
                } else {
                    if (false || false) {
                        var v42 bool = false
                        _ = v42
                    }
                }
            default:
                v43 := v3[2147483647]
                _ = v43
            }
            v44 := GenS3{F0: 2}
            _ = v44
        } else {
            v45 := GenS4{F0: 0, F1: 1, F2: -1}
            for fi46 := 0; fi46 < 10; fi46++ {
                fmt.Println(fi46)
                var v47 uint16 = 474
                for fi48 := 0; fi48 < 5; fi48++ {
                    v49 := GenS2{F0: 903, F1: 42, F2: "a"}
                    if true {
                        v47 = (v47 * 268)
                        var v50 bool = true
                        _ = v50
                    } else {
                        v47 += (v47 / ((v47) | 1))
                    }
                    _ = fi48
                    _ = v49
                }
                _ = fi46
                _ = v47
            }
            if (2147483647 <= 657) {
                switch -351 {
                case 0:
                    v1[65535] = 32767
                case 2147483647:
                    var v51 uint = 65535
                    _ = v51
                }
                var v52 uint64 = 18446744073709551615
                _ = v52
            }
            _ = v45
        }
        _ = v4
    } else {
        if (!(!(1e10 != 1e10))) {
            switch 25 {
            case -179:
                var v53 uint64 = 4294967295
                _ = v53
            case 42:
                v54 := GenS1{F0: 32767, F1: 4294967295}
                _ = v54
            case 772:
                switch ((("hello" + ("\n\t" + "a")) + "\n\t") + "foo bar") {
                case "":
                    for fi55 := 0; fi55 < 1; fi55++ {
                        for fi56 := 0; fi56 < 10; fi56++ {
                            v57 := GenS4{F0: 1, F1: 32767, F2: -2147483648}
                            var v58 string = (("" + ("" + "hello")) + "hello")
                            v58 = "\n\t"
                            _ = len(v2)
                            if false {
                                var v59 uint64 = 4294967295
                                if true {
                                    var v60 uint8 = 255
                                    var v61 int = 4294967295
                                    v62 := GenS1{F0: 2, F1: 4294967295}
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                                _ = v59
                            } else {
                                var v63 string = (v58 + ((v58 + (v58 + "\n\t")) + v58))
                                _ = v63
                            }
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                        _ = fi55
                    }
                case "foo bar":
                    var v64 int64 = -1
                    _ = v64
                default:
                    switch "\n\t" {
                    case "hello":
                        for fi65 := 0; fi65 < 9; fi65++ {
                            var v66 uint = 378
                            _ = fi65
                            _ = v66
                        }
                    case "\n\t":
                        var v67 int32 = 1
                        _ = v67
                    default:
                        var v68 uint32 = 121
                        _ = v68
                    }
                }
            default:
                if (1 < 18446744073709551615) {
                    var v69 int16 = 42
                    var v70 bool = true
                    _ = v69
                    _ = v70
                }
            }
            v71 := GenS2{F0: 888, F1: 4294967295, F2: "\n\t"}
            _ = v71
        }
    }
    if (!true) {
        var v72 string = "\n\t"
        delete(v3, 333)
        _ = v72
    } else {
        for fi73 := 0; fi73 < 4; fi73++ {
            switch (fi73 >> ((1) & 63)) {
            case 369:
                var v74 float64 = (0.0 - -2.25)
                _ = v74
            case 1:
                v2 = append(v2, 491)
            }
            _ = fi73
        }
        var v75 string = "foo bar"
        _ = v75
    }
    var v76 uint8 = 2
    switch 32768 {
    case 1:
        v76 = (v76 / ((v76) | 1))
    case 730:
        _ = len(v1)
    case 0:
        var v77 int16 = -794
        _ = v77
    }
    v78 := GenS2{F0: 1, F1: 42, F2: "foo bar"}
    if false {
        var v79 uint32 = 65535
        v2 = append(v2, 0)
        var v80 string = ("hello" + "hello")
        _ = v79
        _ = v80
    }
    var v81 int32 = 708
    for fi82 := 0; fi82 < 19; fi82++ {
        if true {
            if true {
                var v83 uint = 255
                v76 -= (v76 ^ v76)
                _ = v83
            }
            if (891 < 65535) {
                v2 = append(v2, fi82)
                fmt.Println(v81)
            } else {
                _ = len(v3)
            }
            if true {
                v81 += (^(v81))
            } else {
                v76 = 128
            }
        }
        var v84 float32 = (float32((3.14159 + -2.25)) / 3.14159)
        fmt.Println(fi82)
        _ = fi82
        _ = v84
    }
    for fi85 := 0; fi85 < 18; fi85++ {
        fmt.Println(v76, fi85)
        if false {
            v76 += v76
            if true {
                for fi86 := 0; fi86 < 12; fi86++ {
                    if (false && (false && (9223372036854775807 >= 9223372036854775807))) {
                        if (((!false) || true) || (1.5 == -2.25)) {
                            v87 := GenS4{F0: v81, F1: -1, F2: (v81 << ((2) & 63))}
                            fmt.Println(v81)
                            _ = v87
                        } else {
                            v76 -= (v76 & (v76 << ((2) & 63)))
                            if true {
                                var v88 uint32 = 2147483648
                                _ = v88
                            }
                            v76 = 2
                            v76 = (v76 + 21)
                            switch "\n\t" {
                            case "\n\t":
                                var v89 uint = 445
                                _ = v89
                            case "hello":
                                var v90 string = "a"
                                _ = v90
                            case "":
                                var v91 int64 = 165
                                _ = v91
                            }
                        }
                        v81 = v81
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v92 := GenS1{F0: 42, F1: 1}
                    var v93 string = ("" + "hello")
                    var v94 float32 = 0.1
                    _ = fi86
                    _ = v92
                    _ = v93
                    _ = v94
                }
                for fi95 := 0; fi95 < 10; fi95++ {
                    switch "" {
                    case "foo bar":
                        v76 = (^(v76))
                    case "a":
                        v76 += (v76 >> ((65535) & 63))
                    case "hello":
                        fmt.Println(v76, fi85)
                    default:
                        if true {
                            var v96 bool = (1 < 797)
                            _ = v96
                        } else {
                            var v97 string = ((("" + "") + "hello") + "hello")
                            v97 = v97
                            _ = len(v1)
                            v97 = "a"
                            switch fi95 {
                            case -319:
                                v97 = v97
                            case 172:
                                if (!false) {
                                    v76 = v76
                                } else {
                                    var v98 bool = (2147483648 >= 1)
                                    v2 = append(v2, 4294967295)
                                    fmt.Println(v81, v76, v97)
                                    _ = v98
                                }
                            case 32768:
                                switch "\n\t" {
                                case "a":
                                    if false {
                                        v97 = v97
                                        _ = len(v2)
                                        var v99 int = (fi85 + 9223372036854775807)
                                        _ = v99
                                    } else {
                                        v76 = (^(v76))
                                        v97 = v97
                                        fmt.Println(fi85, fi95, fi95)
                                    }
                                case "foo bar":
                                    v97 = v97
                                default:
                                    v97 = "hello"
                                }
                                v100 := GenS5{F0: (v81 & (v81 / ((v81) | 1)))}
                                _ = v100
                            default:
                                switch v97 {
                                case "":
                                    var v101 string = v97
                                    _ = v101
                                case "\n\t":
                                    var v102 bool = true
                                    if len(v2) > 0 {
                                        v2[0] = 127
                                    }
                                    v76 -= (v76 << ((4294967295) & 63))
                                    _ = v102
                                case "foo bar":
                                    v97 = (v97 + "foo bar")
                                }
                            }
                            _ = v97
                        }
                    }
                    _ = fi95
                }
            } else {
                var v103 bool = (float64(float32(float64(float32(1.5)))) < 1e10)
                v81 += (-(v81))
                fmt.Println(v76)
                _ = v103
            }
            v104 := GenS1{F0: 2, F1: 685}
            _ = v104
        } else {
            var v105 string = "a"
            if ((^(fi85)) != (fi85 | (fi85 * int(v81)))) {
                v105 = v105
                v106 := GenS2{F0: 2, F1: 2147483647, F2: v105}
                _ = v106
            } else {
                var v107 string = v105
                v76 = v76
                for fi108 := 0; fi108 < 4; fi108++ {
                    if false {
                        var v109 uint16 = 481
                        var v110 int64 = 1
                        v1[42] = (fi108 | fi108)
                        v111 := GenS6{F0: fi85, F1: -1}
                        _ = v109
                        _ = v110
                        _ = v111
                    }
                    if (false && true) {
                        v1[65535] = 32768
                    } else {
                        var v112 bool = false
                        switch fi85 {
                        case 42:
                            v112 = ((1 < 42) || (v107 > (v107 + "a")))
                        }
                        _ = v112
                    }
                    _ = fi108
                }
                _ = v107
            }
            _ = v105
        }
        for fi113 := 0; fi113 < 12; fi113++ {
            switch "\n\t" {
            case "\n\t":
                v114 := GenS7{F0: 502, F1: 73, F2: -9223372036854775808}
                _ = v114
            }
            switch "foo bar" {
            case "\n\t":
                v115 := GenS8{F0: 42, F1: fi85}
                _ = v115
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = 9223372036854775807
                }
            case "":
                if true {
                    if false {
                        fmt.Println(fi113, v81)
                    } else {
                        var v116 bool = false
                        v117 := GenS4{F0: v81, F1: -332, F2: (v81 / (((v81 ^ 2147483647)) | 1))}
                        _ = v116
                        _ = v117
                    }
                }
            }
            _ = fi113
        }
        _ = fi85
    }
    v76 = (v76 & v76)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v76:", v76)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v76
    _ = v78
    _ = v81
}
