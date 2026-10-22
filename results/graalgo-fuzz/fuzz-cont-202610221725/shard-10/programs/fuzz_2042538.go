package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int; F1 int8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 string; F1 int8 }
func genFn0(p0 uint16, p1 uint32, p2 uint16) int16 {
    p2 = (^(p2))
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{(fi7 + (fi7 + fi7)), fi7}
        _ = fi7
        _ = v8
    }
    var v9 float32 = float32(0.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -32768
}

func main() {
    var v1 uint8 = 128
    var v2 bool = (false && false)
    v3 := make(map[string]uint32)
    var v4 uint = 2
    v5 := make(map[int]string)
    v6 := []uint32{uint32(v1), 2}
    fmt.Println(v4)
    for fi10 := 0; fi10 < 9; fi10++ {
        v4 -= (v4 / ((1) | 1))
        _ = fi10
    }
    v11 := GenS1{F0: ("hello" + "\n\t"), F1: 1, F2: 1}
    v12 := GenS1{F0: ("a" + "hello"), F1: 2147483647, F2: 32767}
    var v13 uint8 = v1
    if v2 {
        fmt.Println(v1, v2, v13)
        for fi14 := 0; fi14 < 8; fi14++ {
            _ = len(v6)
            if v2 {
                switch ((("a" + "\n\t") + "\n\t") + ("" + "\n\t")) {
                case "hello":
                    var v15 float64 = float64((0.1 - ((1e10 + 3.14159) - 1.5)))
                    _ = v15
                default:
                    v13 += (v13 + v13)
                }
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                if len(v6) > 0 {
                    v6[0] = 0
                }
                if false {
                    var v17 int64 = -1
                    var v18 uint16 = 0
                    _ = v17
                    _ = v18
                }
                if false {
                    var v19 float32 = 3.14159
                    _ = v19
                }
                var v20 int8 = 42
                _ = fi16
                _ = v20
            }
            _ = fi14
        }
    } else {
        switch 478 {
        case 127:
            v21 := GenS1{F0: "hello", F1: uint32(v4), F2: -159}
            _ = v21
        case 4294967295:
            v22 := GenS1{F0: ("" + ("foo bar" + "\n\t")), F1: 965, F2: 1}
            _ = v22
        }
        var v23 uint64 = 18446744073709551615
        switch ("\n\t" + "hello") {
        case "a":
            var v24 int16 = 42
            _ = v24
        }
        _ = v23
    }
    v25 := GenS2{F0: -1}
    var v26 string = "foo bar"
    v6 = append(v6, 42)
    if true {
        fmt.Println(v1, v4, v4)
        switch v26 {
        case "a":
            fmt.Println(v26, v2)
        case "hello":
            for fi27 := 0; fi27 < 20; fi27++ {
                v26 = ((v26 + (("\n\t" + v26) + "")) + (v26 + ""))
                if (0 >= -1) {
                    var v28 float64 = 3.14159
                    _ = v28
                }
                _ = fi27
            }
        default:
            var v29 int = 127
            _ = v29
        }
        if true {
            v1 += v13
        } else {
            var v30 string = ("\n\t" + ((v26 + "\n\t") + v26))
            if ((!((v2 || true) && ((2147483647 == 65535) || (-136 != genFn0(42, 2147483647, 795))))) || v2) {
                v31 := GenS2{F0: 42}
                _ = v31
            }
            _ = v30
        }
    } else {
        var v32 bool = (0 <= 32767)
        _ = v32
    }
    v33 := GenS3{F0: (v13 - (v1 - (v13 / ((v13) | 1))))}
    switch "" {
    case "hello":
        for fi34 := 0; fi34 < 1; fi34++ {
            for fi35 := 0; fi35 < 14; fi35++ {
                var v36 int = fi35
                if false {
                    v6 = append(v6, 65535)
                    var v37 uint = v4
                    v37 += 65535
                    _ = v37
                }
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    case "foo bar":
        var v38 int = 127
        _ = v38
    case "a":
        switch v26 {
        case "":
            switch 0 {
            case 831:
                for fi39 := 0; fi39 < 11; fi39++ {
                    v13 -= v13
                    if (true || (0 == 42)) {
                        var v40 uint32 = uint32(v4)
                        v4 = (v4 >> ((v4) & 63))
                        v2 = (fi39 <= -685)
                        _ = v40
                    }
                    var v41 bool = true
                    _ = fi39
                    _ = v41
                }
            case -1:
                if ((genFn0(1, 1, 0) != 1) && true) {
                    for fi42 := 0; fi42 < 2; fi42++ {
                        if v2 {
                            for fi43 := 0; fi43 < 1; fi43++ {
                                v6 = append(v6, 42)
                                _ = fi43
                            }
                        } else {
                            var v44 string = v26
                            if true {
                                v6 = append(v6, 674)
                                v45 := v5[(fi42 * (fi42 % (((fi42 & -9223372036854775808)) | 1)))]
                                _ = v45
                            }
                            switch (v44 + "foo bar") {
                            case "a":
                                v2 = ((!(!(1000 <= genFn0(65535, 2147483648, 140)))) || (!false))
                            case "":
                                fmt.Println(fi42)
                            default:
                                if false {
                                    switch "a" {
                                    case "foo bar":
                                        v26 = v26
                                    }
                                    if true {
                                        v2 = (1 != 1000)
                                    } else {
                                        v4 = 0
                                        switch v44 {
                                        case "foo bar":
                                            v26 = "a"
                                        default:
                                            v46 := GenS2{F0: 9223372036854775807}
                                            _ = v46
                                        }
                                    }
                                    if (!((v2 && v2) || (!v2))) {
                                        var v47 int = (fi42 << (((v4 - uint(v1))) & 63))
                                        for fi48 := 0; fi48 < 20; fi48++ {
                                            v49 := GenS2{F0: (v47 / ((v47) | 1))}
                                            v47 += 9223372036854775807
                                            v50 := GenS4{F0: v44, F1: 46}
                                            _ = fi48
                                            _ = v49
                                            _ = v50
                                        }
                                        v51 := v3[(v44 + "")]
                                        _ = v47
                                        _ = v51
                                    } else {
                                        for fi52 := 0; fi52 < 10; fi52++ {
                                            v13 = 1
                                            _ = fi52
                                        }
                                        for fi53 := 0; fi53 < 12; fi53++ {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v6 = append(v6, 2147483647)
                                            v1 = v13
                                            if true {
                                                v54 := GenS5{F0: fi53, F1: 42}
                                                v44 = v26
                                                if v2 {
                                                    delete(v5, fi53)
                                                    var v55 uint = v4
                                                    _ = v55
                                                } else {
                                                    if v2 {
                                                        v56 := GenS4{F0: v26, F1: -1}
                                                        _ = v56
                                                    }
                                                    v2 = (((!(true || (v2 && v2))) && v2) && v2)
                                                    v13 = v13
                                                }
                                                v1 -= v13
                                                v57 := GenS6{F0: v44, F1: 1}
                                                _ = v54
                                                _ = v57
                                            }
                                            _ = fi53
                                        }
                                    }
                                } else {
                                    var v58 int32 = -2147483648
                                    v2 = (genFn0(65535, 2, 65535) < 1)
                                    for fi59 := 0; fi59 < 11; fi59++ {
                                        _ = len(v6)
                                        _ = fi59
                                    }
                                    _ = v58
                                }
                            }
                            _ = v44
                        }
                        v60 := GenS7{F0: -422, F1: 1}
                        v61 := GenS6{F0: v26, F1: 9223372036854775807}
                        _ = fi42
                        _ = v60
                        _ = v61
                    }
                    if v2 {
                        v13 = (v13 >> ((1) & 63))
                        for fi62 := 0; fi62 < 11; fi62++ {
                            switch (fi62 - 9223372036854775807) {
                            case 127:
                                for fi63 := 0; fi63 < 5; fi63++ {
                                    v13 = (v1 + 1)
                                    v64 := GenS3{F0: 255}
                                    _ = fi63
                                    _ = v64
                                }
                            default:
                                v65 := GenS1{F0: "", F1: 65535, F2: 32767}
                                _ = v65
                            }
                            _ = len(v6)
                            fmt.Println(v1)
                            _ = fi62
                        }
                    } else {
                        switch ("foo bar" + ("a" + ("a" + (v26 + v26)))) {
                        case "a":
                            if (true || true) {
                                var v66 int64 = 9223372036854775807
                                v13 = (v1 / ((v1) | 1))
                                for fi67 := 0; fi67 < 14; fi67++ {
                                    for fi68 := 0; fi68 < 17; fi68++ {
                                        v26 = "\n\t"
                                        v13 = 123
                                        var v69 float32 = 0.0
                                        _ = fi68
                                        _ = v69
                                    }
                                    _ = fi67
                                }
                                _ = v66
                            }
                        case "":
                            v70 := GenS8{F0: 48}
                            _ = v70
                        default:
                            var v71 uint32 = 4294967295
                            _ = v71
                        }
                        if v2 {
                            if (-465 < 1000) {
                                switch v26 {
                                case "foo bar":
                                    if true {
                                        v2 = true
                                    }
                                }
                                if v2 {
                                    fmt.Println(v4, v26)
                                    for fi72 := 0; fi72 < 12; fi72++ {
                                        v73 := GenS9{F0: 2, F1: 65535}
                                        _ = fi72
                                        _ = v73
                                    }
                                } else {
                                    var v74 string = "a"
                                    switch "" {
                                    case "\n\t":
                                        for fi75 := 0; fi75 < 19; fi75++ {
                                            v2 = v2
                                            v2 = v2
                                            var v76 uint32 = 1
                                            _ = fi75
                                            _ = v76
                                        }
                                    case "hello":
                                        fmt.Println(v2)
                                    case "foo bar":
                                        fmt.Println(v74, v2, v4)
                                    }
                                    fmt.Println(v26, v74)
                                    _ = v74
                                }
                                switch 99 {
                                case 556:
                                    fmt.Println(v1, v4)
                                case -982:
                                    var v77 string = "hello"
                                    _ = v77
                                }
                            } else {
                                if ((v2 || (((v2 && (false || v2)) && v2) && ((((3.14159 - 0.0) / -2.25) * (0.0 + -2.25)) == 1.5))) || ((189 < uint64(v13)) && (128 < 0))) {
                                    v78 := GenS10{F0: -1}
                                    fmt.Println(v26, v2, v1)
                                    for fi79 := 0; fi79 < 9; fi79++ {
                                        if v2 {
                                            var v80 int64 = 0
                                            if v2 {
                                                var v81 uint64 = 4294967295
                                                v81 = (^(v81))
                                                switch fi79 {
                                                case 128:
                                                    if v2 {
                                                        var v82 int8 = -128
                                                        var v83 uint16 = 1
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        var v84 uint64 = (v81 / (((v81 * v81)) | 1))
                                                        _ = v84
                                                    }
                                                case 4294967295:
                                                    if false {
                                                        v26 = ""
                                                        switch fi79 {
                                                        case 2147483647:
                                                            fmt.Println(fi79)
                                                        case -1:
                                                            var v85 int = (fi79 * fi79)
                                                            _ = v85
                                                        }
                                                        v2 = v2
                                                    } else {
                                                        for fi86 := 0; fi86 < 12; fi86++ {
                                                            fmt.Println(v1, fi86)
                                                            _ = fi86
                                                        }
                                                        v87 := GenS11{F0: fi79, F1: 123}
                                                        _ = len(v6)
                                                        _ = v87
                                                    }
                                                case 9223372036854775807:
                                                    var v88 int16 = int16(fi79)
                                                    _ = v88
                                                }
                                                _ = v81
                                            }
                                            v1 -= (v13 * v13)
                                            _ = v80
                                        } else {
                                            fmt.Println(v13, v26, v4)
                                        }
                                        v89 := GenS12{F0: 4294967295}
                                        _ = fi79
                                        _ = v89
                                    }
                                    _ = v78
                                } else {
                                    v13 = (v1 - v1)
                                }
                                var v90 uint16 = 613
                                v13 += 0
                                _ = v90
                            }
                            for fi91 := 0; fi91 < 14; fi91++ {
                                if v2 {
                                    var v92 float32 = float32(0.1)
                                    _ = v92
                                }
                                for fi93 := 0; fi93 < 6; fi93++ {
                                    var v94 string = "foo bar"
                                    var v95 uint16 = 65535
                                    _ = fi93
                                    _ = v94
                                    _ = v95
                                }
                                v26 = "a"
                                _ = fi91
                            }
                            if v2 {
                                v3[(v26 + "\n\t")] = 2
                                v96 := GenS5{F0: 2147483647, F1: 100}
                                v4 += 0
                                _ = v96
                            }
                        }
                        v3[("hello" + v26)] = 273
                    }
                    for fi97 := 0; fi97 < 3; fi97++ {
                        for fi98 := 0; fi98 < 9; fi98++ {
                            v26 = ("foo bar" + (v26 + v26))
                            _ = fi98
                        }
                        _ = fi97
                    }
                } else {
                    if len(v6) > 0 {
                        v6[0] = uint32(v13)
                    }
                }
            case -550:
                fmt.Println(v1, v1)
            default:
                fmt.Println(v13, v13)
            }
        default:
            for fi99 := 0; fi99 < 9; fi99++ {
                v100 := GenS9{F0: 42, F1: 0}
                var v101 uint64 = 965
                _ = fi99
                _ = v100
                _ = v101
            }
        }
    default:
        v26 = v26
    }
    var v102 string = ("a" + v26)
    v2 = v2
    for fi103 := 0; fi103 < 18; fi103++ {
        v102 = v102
        for fi104 := 0; fi104 < 3; fi104++ {
            switch v26 {
            case "":
                v105 := GenS13{F0: ("hello" + v102), F1: 100}
                var v106 float64 = 1.5
                var v107 int8 = -128
                _ = v105
                _ = v106
                _ = v107
            case "hello":
                var v108 bool = ((-453 < 1) && (1 == 1000))
                _ = v108
            case "\n\t":
                var v109 float32 = 1.5
                _ = v109
            default:
                if ((!v2) || v2) {
                    var v110 int8 = -1
                    v13 = (v1 ^ 128)
                    _ = v110
                } else {
                    if v2 {
                        if ("a" <= (v102 + ((v102 + "foo bar") + ""))) {
                            if (433 < 2) {
                                v26 = v26
                                fmt.Println(v26)
                            } else {
                                v4 -= v4
                                switch fi104 {
                                case -2147483648:
                                    v13 = (^(v13))
                                default:
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v26, fi104, v4)
                                        v26 = (((v26 + (v102 + "a")) + ("hello" + v102)) + (v102 + ((v26 + (v26 + ((v102 + "hello") + v26))) + "")))
                                    case "hello":
                                        if (0.1 >= (0.1 - -2.25)) {
                                            v111 := GenS2{F0: 0}
                                            _ = v111
                                        }
                                    default:
                                        var v112 uint64 = 18446744073709551615
                                        _ = v112
                                    }
                                }
                                var v113 uint16 = 436
                                _ = v113
                            }
                            v26 = v102
                            var v114 uint8 = 128
                            _ = v114
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v102 = v102
                            v4 -= v4
                        }
                        v102 = v26
                        if false {
                            v13 -= v13
                            switch fi104 {
                            case 9223372036854775807:
                                _ = len(v6)
                            default:
                                v6 = append(v6, 956)
                            }
                        } else {
                            delete(v3, (v102 + ("a" + (v102 + "foo bar"))))
                            delete(v3, "")
                            if len(v6) > 0 {
                                v6[0] = 941
                            }
                        }
                    }
                    fmt.Println(v4)
                }
            }
            var v115 float32 = 0.1
            _ = fi104
            _ = v115
        }
        _ = len(v6)
        _ = fi103
    }
    if true {
        var v116 int64 = 9223372036854775807
        _ = v116
    }
    var v117 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v117:", v117)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v102
    _ = v11
    _ = v117
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
