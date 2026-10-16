package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint }
type GenS4 struct { F0 string; F1 uint64; F2 uint }
type GenS5 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 int32; F2 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int64; F1 int64; F2 uint8 }
type GenS13 struct { F0 int; F1 string; F2 int }
type GenS14 struct { F0 string; F1 uint16; F2 uint8 }
type GenS15 struct { F0 bool; F1 int8 }
func genFn0(p0 uint8) bool {
    p0 = 2
    var v5 float64 = 1.5
    if false {
        p0 = p0
        for fi6 := 0; fi6 < 14; fi6++ {
            if false {
                v7 := GenS1{F0: (291 > 1), F1: "foo bar", F2: -9223372036854775808}
                v8 := []int{}
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 uint = 255
                    if ((float64(1.5) == v5) && true) {
                        var v11 float64 = v5
                        p0 -= (p0 << ((v10) & 63))
                        fmt.Println(v10, fi6)
                        _ = v11
                    } else {
                        var v12 float64 = (float64((((((1e10 + 1e10) - 3.14159) + 0.1) + 0.1) * ((1e10 - 1e10) + (1.5 + 0.1)))) / v5)
                        _ = v12
                    }
                    if (0.1 == float32((3.14159 * (v5 + v5)))) {
                        v13 := GenS2{F0: 100}
                        _ = v13
                    } else {
                        v14 := []int{1}
                        switch ("" + "") {
                        case "":
                            fmt.Println(v10, p0, p0)
                        case "foo bar":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        case "\n\t":
                            v15 := GenS1{F0: (!false), F1: ("" + "hello"), F2: (fi6 % (((fi9 + (fi9 - fi9))) | 1))}
                            switch (fi9 % ((fi6) | 1)) {
                            case 128:
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            case 657:
                                v16 := GenS2{F0: -74}
                                _ = v16
                            case 951:
                                v5 -= (v5 - 1.5)
                            }
                            _ = v15
                        }
                        _ = v14
                    }
                    _ = fi9
                    _ = v10
                }
            }
            _ = fi6
        }
        switch ("hello" + "\n\t") {
        case "hello":
            var v17 uint = 0
            _ = v17
        case "a":
            var v18 int8 = 83
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        switch (fi19 | -1) {
        case 2147483647:
            fmt.Println(fi19, v5, p0)
        default:
            v20 := []int{fi19, (-(fi19))}
            _ = v20
        }
        fmt.Println(fi19, v5)
        _ = fi19
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v5 += ((v5 * (-2.25 + (float64(1.5) + v5))) + 0.0)
        var v22 int16 = -1
        _ = fi21
        _ = v22
    }
    _ = p0
    _ = v5
    return false
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    var v4 uint8 = 55
    var v23 int = -708
    if false {
        var v24 int64 = 1
        _ = v24
    }
    _ = len(v1)
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 int64 = -9223372036854775808
        var v27 uint32 = 0
        v28 := GenS3{F0: 127, F1: (!(1e10 > (0.1 - 1e10))), F2: uint(v4)}
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    v29 := []int{v23, -1, (-(v23))}
    var v30 int16 = 1
    fmt.Println(v23)
    for fi31 := 0; fi31 < 1; fi31++ {
        if (v23 <= (^(fi31))) {
            switch (fi31 * fi31) {
            case 0:
                var v32 uint32 = 48
                _ = v32
            case 128:
                v23 += v23
            default:
                v33 := GenS2{F0: 127}
                _ = v33
            }
            if true {
                fmt.Println(v4)
                for fi34 := 0; fi34 < 13; fi34++ {
                    if true {
                        v4 = (v4 | v4)
                    } else {
                        if false {
                            var v35 uint8 = (^(v4))
                            v35 += v35
                            _ = v35
                        } else {
                            if true {
                                v36 := GenS4{F0: "foo bar", F1: 1, F2: 255}
                                v23 = fi34
                                var v37 uint64 = 18446744073709551615
                                _ = v36
                                _ = v37
                            }
                            var v38 float64 = (float64((((float32(((1.5 / -2.25) * 1.5)) - float32(float64((((0.1 * (1.5 * (0.0 + float32(0.0)))) / 1e10) / (((((float32(0.1) * 3.14159) - (3.14159 * (-2.25 - 1e10))) - 1.5) - (float32((3.14159 * 0.1)) / 3.14159)) * float32(1.5)))))) / 1.0) - ((((1.5 + 0.0) / (3.14159 / 3.14159)) * (1.5 / ((1e10 - 1.5) / 1e10))) / 0.1))) / (0.0 + (3.14159 * 1e10)))
                            v39 := v2[(("hello" + "a") + ("a" + "a"))]
                            if (!(18446744073709551615 != 18446744073709551615)) {
                                var v40 uint32 = 312
                                if false {
                                    if false {
                                        if v39 {
                                            fmt.Println(fi34, fi34)
                                            if v39 {
                                                switch fi34 {
                                                case 0:
                                                    v4 -= v4
                                                case 2147483647:
                                                    var v41 string = ""
                                                    _ = v41
                                                default:
                                                    var v42 int16 = v30
                                                    _ = v42
                                                }
                                                var v43 uint32 = v40
                                                switch "a" {
                                                case "\n\t":
                                                    v29 = append(v29, -2147483648)
                                                case "hello":
                                                    if v39 {
                                                        var v44 bool = (0.0 <= float64(0.0))
                                                        v45 := GenS4{F0: "foo bar", F1: 820, F2: 4294967295}
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        var v46 bool = v39
                                                        if v39 {
                                                            fmt.Println(v23, v23)
                                                            v47 := GenS5{F0: 1, F1: 750, F2: 69}
                                                            if len(v29) > 0 {
                                                                v29[0] = fi34
                                                            }
                                                            _ = v47
                                                        }
                                                        v48 := GenS4{F0: "hello", F1: 42, F2: 42}
                                                        _ = v46
                                                        _ = v48
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v38, v38)
                                                default:
                                                    if len(v29) > 0 {
                                                        v29[0] = (fi31 & (fi34 >> ((1) & 63)))
                                                    }
                                                }
                                                fmt.Println(v30)
                                                var v49 float32 = (0.1 * 0.1)
                                                v50 := GenS6{F0: 514}
                                                _ = v43
                                                _ = v49
                                                _ = v50
                                            } else {
                                                v39 = v39
                                                var v51 uint16 = 1
                                                _ = v51
                                            }
                                        }
                                        switch (("" + "hello") + "hello") {
                                        case "\n\t":
                                            var v52 uint8 = (v4 ^ (v4 - v4))
                                            _ = v52
                                        case "hello":
                                            v4 = (^(v4))
                                        case "foo bar":
                                            switch -328 {
                                            case 520:
                                                if v39 {
                                                    if (v30 <= v30) {
                                                        var v53 bool = (v38 != 1e10)
                                                        if len(v29) > 0 {
                                                            v29[0] = 289
                                                        }
                                                        v4 = (v4 - v4)
                                                        _ = v53
                                                    } else {
                                                        var v54 uint64 = 465
                                                        switch "\n\t" {
                                                        case "a":
                                                            v54 -= uint64(v40)
                                                        }
                                                        var v55 string = "foo bar"
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                    v38 -= 3.14159
                                                }
                                            case 4294967295:
                                                var v56 string = "a"
                                                _ = v56
                                            default:
                                                v57 := GenS4{F0: "foo bar", F1: 4294967295, F2: 2}
                                                _ = v57
                                            }
                                        }
                                        v30 -= (v30 >> ((0) & 63))
                                    }
                                    var v58 uint64 = 294
                                    _ = len(v29)
                                    _ = v58
                                } else {
                                    v59 := GenS7{F0: 127, F1: 0, F2: v4}
                                    _ = v59
                                }
                                v4 += (^(v4))
                                _ = v40
                            }
                            _ = v38
                            _ = v39
                        }
                    }
                    switch (("foo bar" + "foo bar") + (("foo bar" + "foo bar") + ("hello" + "hello"))) {
                    case "a":
                        fmt.Println(fi31, v4, v23)
                    case "hello":
                        var v60 int64 = 42
                        _ = v60
                    case "":
                        var v61 int16 = (v30 / (((v30 % ((v30) | 1))) | 1))
                        _ = v61
                    }
                    _ = fi34
                }
            } else {
                var v62 int64 = 0
                _ = v62
            }
        }
        _ = len(v29)
        var v63 int64 = -9223372036854775808
        _ = fi31
        _ = v63
    }
    var v64 int8 = -1
    v65 := GenS8{F0: v23}
    v66 := GenS9{F0: "a", F1: v64}
    var v67 float64 = (((-2.25 + 1e10) * 1e10) / 3.14159)
    v68 := GenS10{F0: 1, F1: -834}
    var v69 int8 = v64
    v69 = -1
    if false {
        var v70 uint = 1
        v69 += v69
        switch int(v4) {
        case 2147483647:
            if true {
                fmt.Println(v70, v4)
                for fi71 := 0; fi71 < 9; fi71++ {
                    fmt.Println(v30, v30, v4)
                    _ = fi71
                }
                for fi72 := 0; fi72 < 19; fi72++ {
                    for fi73 := 0; fi73 < 6; fi73++ {
                        v74 := GenS11{F0: v30}
                        _ = fi73
                        _ = v74
                    }
                    v30 = 32767
                    for fi75 := 0; fi75 < 15; fi75++ {
                        v30 = v30
                        var v76 uint = v70
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        _ = fi75
                        _ = v76
                    }
                    _ = fi72
                }
            } else {
                var v77 float32 = (1.5 * ((-2.25 + -2.25) / 1.0))
                if (2147483648 < 2) {
                    for fi78 := 0; fi78 < 5; fi78++ {
                        var v79 int32 = 42
                        _ = fi78
                        _ = v79
                    }
                    for fi80 := 0; fi80 < 20; fi80++ {
                        fmt.Println(fi80)
                        v81 := GenS5{F0: (v70 & 677), F1: 18446744073709551615, F2: 0}
                        v82 := v1["hello"]
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                }
                _ = v77
            }
        case 42:
            var v83 uint16 = 32767
            _ = v83
        }
        _ = v70
    }
    for fi84 := 0; fi84 < 8; fi84++ {
        switch (fi84 << ((42) & 63)) {
        case -2147483648:
            for fi85 := 0; fi85 < 13; fi85++ {
                v86 := GenS12{F0: -290, F1: 0, F2: v4}
                var v87 int8 = v69
                if false {
                    if false {
                        switch "hello" {
                        case "hello":
                            if (fi85 < (fi85 / ((-169) | 1))) {
                                if (false || (((0 != 0) && false) || (false || true))) {
                                    v88 := GenS13{F0: v23, F1: "hello", F2: v23}
                                    if len(v29) > 0 {
                                        v29[0] = fi85
                                    }
                                    _ = v88
                                } else {
                                    var v89 string = ("\n\t" + "hello")
                                    switch v89 {
                                    case "hello":
                                        v69 = (-(v69))
                                    default:
                                        fmt.Println(v64)
                                    }
                                    _ = v89
                                }
                            }
                        case "foo bar":
                            v90 := GenS14{F0: "\n\t", F1: 255, F2: v4}
                            _ = v90
                        default:
                            v91 := GenS11{F0: -32768}
                            _ = v91
                        }
                        if true {
                            fmt.Println(v23, v67, v87)
                            if false {
                                switch ("hello" + ("" + ("" + "hello"))) {
                                case "foo bar":
                                    fmt.Println(v64, fi85, fi85)
                                default:
                                    v64 = v64
                                }
                                var v92 uint8 = (v4 | 89)
                                _ = v92
                            } else {
                                switch (fi85 << ((65535) & 63)) {
                                case 32768:
                                    v64 = v87
                                case -681:
                                    v30 = -548
                                }
                                v93 := GenS4{F0: "a", F1: 640, F2: 1}
                                switch "" {
                                case "a":
                                    if ((fi84 ^ (^(fi85))) > (v23 + 32768)) {
                                        if (42 >= -1) {
                                            v64 -= v69
                                        }
                                        var v94 float32 = 1.5
                                        v69 = 100
                                        _ = v94
                                    } else {
                                        var v95 uint8 = (v4 >> ((481) & 63))
                                        _ = v95
                                    }
                                case "":
                                    v96 := GenS10{F0: 42, F1: 0}
                                    _ = v96
                                }
                                _ = v93
                            }
                            if (v4 < v4) {
                                v67 += float64((3.14159 / (1e10 - float32(0.0))))
                                if len(v29) > 0 {
                                    v29[0] = (fi85 - 32767)
                                }
                                if (0.0 < 0.1) {
                                    v23 += 2147483647
                                }
                                if false {
                                    v4 -= 0
                                    if (1 > 65535) {
                                        var v97 string = "hello"
                                        if false {
                                            var v98 bool = (!false)
                                            _ = v98
                                        }
                                        v30 -= v30
                                        _ = v97
                                    } else {
                                        v99 := GenS14{F0: "foo bar", F1: 2, F2: (^(v4))}
                                        v100 := GenS11{F0: v30}
                                        fmt.Println(v30)
                                        _ = v100
                                        _ = v99
                                    }
                                }
                                v23 += -1
                            }
                        } else {
                            if (!false) {
                                var v101 float32 = 1.5
                                v87 += 42
                                _ = v101
                            }
                            if len(v29) > 0 {
                                v29[0] = (^(fi85))
                            }
                        }
                    }
                    switch 128 {
                    case 772:
                        v67 += 0.1
                        var v102 int32 = 42
                        _ = v102
                    case -9223372036854775808:
                        v29 = append(v29, fi85)
                    }
                } else {
                    _ = len(v1)
                }
                if true {
                    fmt.Println(v4)
                    var v103 float32 = ((3.14159 * 0.0) * 3.14159)
                    _ = v103
                } else {
                    if len(v29) > 0 {
                        v29[0] = (fi84 / (((v23 / ((fi85) | 1))) | 1))
                    }
                }
                var v104 uint16 = 1
                v105 := GenS15{F0: (!false), F1: (^(v69))}
                if true {
                    var v106 float64 = float64(1e10)
                    _ = v106
                } else {
                    v4 = v4
                    if false {
                        var v107 int = 127
                        _ = v107
                    }
                }
                var v108 int64 = -889
                _ = fi85
                _ = v104
                _ = v105
                _ = v108
                _ = v86
                _ = v87
            }
        case 9223372036854775807:
            for fi109 := 0; fi109 < 2; fi109++ {
                var v110 uint8 = v4
                _ = fi109
                _ = v110
            }
        }
        _ = fi84
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
