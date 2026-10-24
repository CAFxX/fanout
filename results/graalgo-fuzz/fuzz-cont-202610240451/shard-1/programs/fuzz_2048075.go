package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int; F1 uint64; F2 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 uint16; F2 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 uint16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{}
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int = (fi7 ^ fi7)
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 4294967295, F1: -733, F2: 422}
        _ = fi5
        _ = v6
        _ = v9
    }
    var v10 int = 0
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint16 = 0
        v10 += 800
        _ = fi11
        _ = v12
    }
    fmt.Println(v10)
    _ = v10
    return 0
}
func genFn1() bool {
    var v13 uint16 = 32767
    fmt.Println(v13)
    switch 0 {
    case -665:
        fmt.Println(v13)
    case -932:
        v13 -= v13
    }
    v13 += v13
    if (true || false) {
        v14 := GenS1{F0: 563, F1: 42, F2: 1}
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 bool = false
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint = 42
                var v19 float32 = 3.14159
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi15
            _ = v16
        }
        var v20 float32 = -2.25
        _ = v14
        _ = v20
    }
    _ = v13
    return (-1 <= 42)
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{4294967295, 9223372036854775807, 4294967295, 131}
    var v3 int32 = 1
    var v4 float32 = float32(0.1)
    if ((v3 < (v3 << ((2) & 63))) || (!false)) {
        if (-9 > 42) {
            for fi21 := 0; fi21 < 12; fi21++ {
                v4 = v4
                for fi22 := 0; fi22 < 2; fi22++ {
                    if (false && (-19 >= int8(v3))) {
                        delete(v1, "\n\t")
                        var v23 int = fi21
                        fmt.Println(v4)
                        _ = v23
                    }
                    var v24 int64 = 42
                    fmt.Println(v24, v4)
                    _ = fi22
                    _ = v24
                }
                if (0 < 128) {
                    fmt.Println(fi21)
                } else {
                    switch (("a" + "a") + "foo bar") {
                    case "foo bar":
                        for fi25 := 0; fi25 < 1; fi25++ {
                            v26 := GenS1{F0: 425, F1: (v3 | (v3 ^ 1)), F2: 4294967295}
                            fmt.Println(fi25)
                            _ = fi25
                            _ = v26
                        }
                    case "a":
                        _ = len(v1)
                    case "\n\t":
                        v27 := GenS1{F0: 65, F1: 1, F2: 315}
                        _ = v27
                    default:
                        for fi28 := 0; fi28 < 19; fi28++ {
                            switch (("foo bar" + "") + "") {
                            case "hello":
                                fmt.Println(v4)
                            case "foo bar":
                                var v29 float32 = ((((-2.25 * v4) - float32(0.1)) * v4) - 1e10)
                                _ = v29
                            case "a":
                                if ((v3 + v3) != 536) {
                                    var v30 int16 = genFn0()
                                    _ = v30
                                } else {
                                    if (1 < -9223372036854775808) {
                                        var v31 int16 = -1
                                        var v32 float32 = float32(3.14159)
                                        fmt.Println(v32, fi21)
                                        if true {
                                            if false {
                                                v33 := GenS2{F0: "a", F1: 883}
                                                _ = v33
                                            }
                                            var v34 bool = ((fi28 - int(v31)) == fi28)
                                            v35 := GenS1{F0: 65535, F1: v3, F2: 1}
                                            _ = v34
                                            _ = v35
                                        }
                                        var v36 uint = 65535
                                        _ = v31
                                        _ = v32
                                        _ = v36
                                    }
                                    var v37 int8 = 1
                                    var v38 uint32 = 0
                                    _ = v37
                                    _ = v38
                                }
                            default:
                                _ = len(v1)
                            }
                            v4 += v4
                            _ = fi28
                        }
                    }
                }
                _ = fi21
            }
        } else {
            v39 := GenS1{F0: 850, F1: (v3 + v3), F2: 0}
            _ = v39
        }
    }
    var v40 uint = 943
    var v41 uint = 65535
    var v42 uint16 = 0
    for fi43 := 0; fi43 < 8; fi43++ {
        for fi44 := 0; fi44 < 1; fi44++ {
            if false {
                v3 -= v3
                var v45 string = (("hello" + "foo bar") + "a")
                v46 := GenS3{F0: (v42 / (((v42 / ((v42) | 1))) | 1)), F1: 42, F2: fi43}
                var v47 string = v45
                switch ((v47 + v47) + v47) {
                case "foo bar":
                    v42 = (v42 + 335)
                }
                _ = v45
                _ = v46
                _ = v47
            } else {
                v48 := GenS3{F0: (v42 * (^(v42))), F1: 18446744073709551615, F2: fi43}
                _ = v48
            }
            _ = fi44
        }
        if ((2 >= 12) && (0 >= -1)) {
            if true {
                var v49 float64 = (((1e10 * (0.1 - 3.14159)) * 3.14159) * -2.25)
                _ = v49
            } else {
                for fi50 := 0; fi50 < 7; fi50++ {
                    var v51 bool = (42 > 1)
                    if (false && (100 >= 0)) {
                        if (v51 && false) {
                            v40 = v40
                            v40 += 0
                            var v52 float64 = 1e10
                            switch (("foo bar" + ("hello" + "hello")) + "a") {
                            case "\n\t":
                                var v53 uint64 = 643
                                _ = v53
                            case "foo bar":
                                _ = len(v1)
                                v52 -= v52
                                v4 = ((3.14159 - v4) + float32((v52 / 1.0)))
                                switch "" {
                                case "\n\t":
                                    if true {
                                        var v54 int = fi43
                                        fmt.Println(v41)
                                        _ = v54
                                    }
                                case "foo bar":
                                    var v55 uint32 = 2147483647
                                    _ = v55
                                case "a":
                                    if true {
                                        v1[(("a" + "a") + "")] = (fi50 & (^(fi50)))
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v56 := GenS2{F0: ("foo bar" + (("hello" + "hello") + (("" + "hello") + "\n\t"))), F1: 158}
                                        _ = v56
                                    }
                                default:
                                    var v57 uint64 = 515
                                    _ = v57
                                }
                            case "hello":
                                v4 -= ((v4 - v4) * float32(v52))
                            }
                            var v58 uint = v41
                            _ = v52
                            _ = v58
                        } else {
                            v2 = append(v2, (fi50 ^ fi43))
                            fmt.Println(v3)
                        }
                        var v59 string = "a"
                        if (-2.25 == float64((v4 + ((v4 + v4) / (v4 + (v4 / v4)))))) {
                            if false {
                                var v60 uint64 = 910
                                _ = v60
                            }
                        }
                        _ = v59
                    }
                    _ = fi50
                    _ = v51
                }
            }
            for fi61 := 0; fi61 < 1; fi61++ {
                var v62 int16 = -1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi61
                _ = v62
            }
        } else {
            v63 := GenS4{F0: -1, F1: 198}
            if len(v2) > 0 {
                v2[0] = fi43
            }
            switch fi43 {
            case 149:
                _ = len(v2)
            }
            _ = v63
        }
        _ = fi43
    }
    for fi64 := 0; fi64 < 16; fi64++ {
        v40 = (v40 & (v41 - (v40 & v40)))
        var v65 uint32 = 65535
        _ = fi64
        _ = v65
    }
    if false {
        var v66 int32 = v3
        _ = v66
    }
    v41 += (v40 / ((42) | 1))
    switch 127 {
    case 290:
        switch 127 {
        case 4294967295:
            for fi67 := 0; fi67 < 20; fi67++ {
                var v68 int64 = -1
                delete(v1, "foo bar")
                v69 := GenS4{F0: v68, F1: 0}
                _ = fi67
                _ = v68
                _ = v69
            }
        case 32767:
            v4 -= v4
        case 2147483647:
            v70 := GenS5{F0: 4294967295, F1: 18446744073709551615, F2: int8(v3)}
            _ = v70
        }
    case 4294967295:
        switch "a" {
        case "a":
            var v71 uint16 = v42
            _ = v71
        case "":
            fmt.Println(v40)
        case "\n\t":
            fmt.Println(v40)
        default:
            for fi72 := 0; fi72 < 9; fi72++ {
                var v73 uint64 = 0
                v2 = append(v2, fi72)
                var v74 uint32 = 794
                _ = fi72
                _ = v73
                _ = v74
            }
        }
    case -313:
        var v75 uint16 = 32767
        _ = v75
    }
    v76 := GenS1{F0: uint(v42), F1: v3, F2: 42}
    v77 := GenS2{F0: "hello", F1: 1}
    for fi78 := 0; fi78 < 2; fi78++ {
        v42 = (v42 - (^(v42)))
        fmt.Println(v3, v4)
        v41 -= (v40 * 4294967295)
        _ = fi78
    }
    var v79 string = "\n\t"
    var v80 int16 = 1000
    v80 = (v80 >> ((v41) & 63))
    _ = len(v2)
    if false {
        v42 = 969
    } else {
        var v81 int32 = -2147483648
        v82 := GenS2{F0: "", F1: 21}
        if ((-(v3)) >= int32(v40)) {
            var v83 int32 = (v3 / (((^(v3))) | 1))
            for fi84 := 0; fi84 < 10; fi84++ {
                if false {
                    v85 := GenS3{F0: uint16(v40), F1: 0, F2: int(v80)}
                    if false {
                        fmt.Println(v42)
                        if len(v2) > 0 {
                            v2[0] = (fi84 ^ -682)
                        }
                    } else {
                        _ = len(v1)
                    }
                    _ = v85
                } else {
                    if false {
                        switch fi84 {
                        case 4294967295:
                            v86 := GenS2{F0: ((("foo bar" + (v79 + v79)) + ((v79 + ("" + ((("" + "foo bar") + (v79 + "\n\t")) + v79))) + ("a" + v79))) + v79), F1: 2}
                            _ = v86
                        case 391:
                            v79 = (v79 + ("hello" + (v79 + ("\n\t" + v79))))
                        case 32767:
                            for fi87 := 0; fi87 < 17; fi87++ {
                                v41 = v41
                                _ = fi87
                            }
                        default:
                            v83 += (^(v81))
                        }
                    } else {
                        _ = len(v2)
                        switch 0 {
                        case -380:
                            v83 = (v81 - (^(v83)))
                        case 42:
                            var v88 uint32 = 402
                            _ = v88
                        case 32767:
                            var v89 uint8 = 59
                            _ = v89
                        }
                    }
                    if (float64(v4) >= (1e10 / -2.25)) {
                        v80 -= (v80 >> ((v40) & 63))
                        if ((v4 - v4) > (1.5 + float32(0.0))) {
                            if (true && true) {
                                var v90 uint = (v41 + (^(v41)))
                                for fi91 := 0; fi91 < 12; fi91++ {
                                    v40 = v41
                                    var v92 uint64 = 2
                                    _ = fi91
                                    _ = v92
                                }
                                var v93 uint = (^(v41))
                                _ = v90
                                _ = v93
                            } else {
                                var v94 string = v79
                                _ = v94
                            }
                        } else {
                            switch "hello" {
                            case "foo bar":
                                v95 := GenS1{F0: 0, F1: v81, F2: 2}
                                _ = v95
                            case "a":
                                var v96 int = fi84
                                _ = v96
                            case "hello":
                                v83 -= 2147483647
                            default:
                                for fi97 := 0; fi97 < 12; fi97++ {
                                    v80 = v80
                                    fmt.Println(v81, v81)
                                    _ = fi97
                                }
                            }
                        }
                        var v98 int = fi84
                        _ = v98
                    } else {
                        var v99 int16 = (v80 ^ genFn0())
                        v79 = ""
                        _ = v99
                    }
                }
                _ = fi84
            }
            for fi100 := 0; fi100 < 17; fi100++ {
                switch ("hello" + (v79 + ("foo bar" + (v79 + (v79 + v79))))) {
                case "\n\t":
                    var v101 int64 = 42
                    _ = v101
                case "hello":
                    var v102 uint32 = 2147483648
                    _ = v102
                case "a":
                    v103 := GenS6{F0: v80}
                    _ = v103
                }
                var v104 int8 = 24
                _ = fi100
                _ = v104
            }
            _ = v83
        } else {
            if (true || (1 <= 18446744073709551615)) {
                v105 := GenS1{F0: v41, F1: int32(v41), F2: 271}
                for fi106 := 0; fi106 < 16; fi106++ {
                    for fi107 := 0; fi107 < 8; fi107++ {
                        v3 -= 0
                        switch "\n\t" {
                        case "hello":
                            v108 := GenS6{F0: (v80 % ((42) | 1))}
                            _ = v108
                        default:
                            v109 := v1["hello"]
                            _ = v109
                        }
                        v110 := GenS7{F0: -1, F1: (v42 & v42), F2: (v79 + "hello")}
                        _ = fi107
                        _ = v110
                    }
                    switch (v79 + v79) {
                    case "a":
                        for fi111 := 0; fi111 < 5; fi111++ {
                            var v112 int = int(v81)
                            v4 = 0.0
                            v80 -= 1
                            _ = fi111
                            _ = v112
                        }
                    case "":
                        switch 32767 {
                        case 2147483647:
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = (fi106 ^ fi106)
                                }
                                var v113 int = (fi106 * (fi106 * fi106))
                                _ = v113
                            }
                        case -739:
                            for fi114 := 0; fi114 < 14; fi114++ {
                                v81 = v3
                                var v115 int64 = 42
                                if len(v2) > 0 {
                                    v2[0] = int(v41)
                                }
                                _ = fi114
                                _ = v115
                            }
                        }
                    case "\n\t":
                        switch fi106 {
                        case 9223372036854775807:
                            v4 = 0.1
                        case -1:
                            for fi116 := 0; fi116 < 15; fi116++ {
                                v81 -= v81
                                switch "hello" {
                                case "\n\t":
                                    switch fi116 {
                                    case 127:
                                        v117 := GenS8{F0: 255}
                                        _ = v117
                                    }
                                case "foo bar":
                                    v3 = -2147483648
                                case "hello":
                                    if (1 == 0) {
                                        v118 := GenS4{F0: -9223372036854775808, F1: 0}
                                        v79 = (v79 + v79)
                                        _ = len(v1)
                                        _ = v118
                                    }
                                }
                                _ = fi116
                            }
                        default:
                            var v119 int = (fi106 * (fi106 ^ fi106))
                            _ = v119
                        }
                    default:
                        v120 := GenS9{F0: (v40 + 1), F1: v42}
                        _ = v120
                    }
                    _ = fi106
                }
                _ = v105
            }
            var v121 uint64 = 42
            v122 := v1["a"]
            _ = v121
            _ = v122
        }
        _ = v81
        _ = v82
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v76
    _ = v77
    _ = v79
    _ = v80
}
