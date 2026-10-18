package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS9 struct { F0 string; F1 int8; F2 string }
type GenS10 struct { F0 int64; F1 uint16; F2 int32 }
type GenS11 struct { F0 uint8; F1 uint; F2 int8 }
type GenS12 struct { F0 uint; F1 int16; F2 bool }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint8; F1 int16 }
type GenS15 struct { F0 uint64; F1 int8 }

func main() {
    v1 := []int{-2147483648, 326}
    var v2 uint64 = 341
    var v3 uint = 255
    var v4 bool = true
    var v5 int32 = 665
    fmt.Println(v2, v2, v3)
    v5 += int32(v2)
    v4 = (!(false && (!(!v4))))
    var v6 int = 538
    _ = len(v1)
    fmt.Println(v6)
    fmt.Println(v2)
    switch "foo bar" {
    case "a":
        switch "\n\t" {
        case "":
            for fi7 := 0; fi7 < 7; fi7++ {
                var v8 uint8 = 128
                switch "\n\t" {
                case "\n\t":
                    v3 -= 386
                }
                if v4 {
                    v1 = append(v1, int(v5))
                }
                _ = fi7
                _ = v8
            }
        case "foo bar":
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(v4, v5)
                    var v11 int = v6
                    v12 := GenS1{F0: ((("foo bar" + "foo bar") + "") + "hello")}
                    v11 = -231
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
        case "\n\t":
            switch ("\n\t" + ("hello" + ("foo bar" + "a"))) {
            case "a":
                for fi13 := 0; fi13 < 5; fi13++ {
                    switch "a" {
                    case "foo bar":
                        fmt.Println(fi13)
                    default:
                        if (!false) {
                            switch "" {
                            case "a":
                                switch "foo bar" {
                                case "":
                                    for fi14 := 0; fi14 < 5; fi14++ {
                                        v1 = append(v1, fi13)
                                        _ = fi14
                                    }
                                case "foo bar":
                                    var v15 string = ""
                                    _ = v15
                                case "hello":
                                    fmt.Println(v5)
                                default:
                                    for fi16 := 0; fi16 < 6; fi16++ {
                                        v17 := GenS1{F0: ((("\n\t" + "") + "hello") + ("foo bar" + ("\n\t" + "hello")))}
                                        _ = fi16
                                        _ = v17
                                    }
                                }
                            default:
                                if v4 {
                                    var v18 string = "foo bar"
                                    for fi19 := 0; fi19 < 13; fi19++ {
                                        if v4 {
                                            switch (fi13 / (((-(fi19))) | 1)) {
                                            case 9223372036854775807:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v20 string = "hello"
                                                _ = v20
                                            case 0:
                                                switch -2147483648 {
                                                case 128:
                                                    v3 -= (v3 | 1)
                                                case -658:
                                                    switch v18 {
                                                    case "\n\t":
                                                        v1 = append(v1, (fi13 | fi13))
                                                    case "":
                                                        v21 := GenS1{F0: ((v18 + ((("a" + (v18 + v18)) + v18) + "foo bar")) + v18)}
                                                        _ = v21
                                                    default:
                                                        v22 := GenS2{F0: v2, F1: 229}
                                                        _ = v22
                                                    }
                                                default:
                                                    v1 = append(v1, (fi19 * fi19))
                                                }
                                            default:
                                                var v23 int16 = 1000
                                                _ = v23
                                            }
                                            switch "hello" {
                                            case "hello":
                                                if false {
                                                    if v4 {
                                                        var v24 uint16 = 2
                                                        var v25 int = 0
                                                        if (0.0 != ((0.0 - -2.25) + 1.5)) {
                                                            if len(v1) > 0 {
                                                                v1[0] = 1
                                                            }
                                                            var v26 uint32 = 2147483647
                                                            v6 = int(v2)
                                                            _ = v26
                                                        } else {
                                                            v24 = v24
                                                            var v27 uint64 = v2
                                                            _ = v27
                                                        }
                                                        v28 := GenS2{F0: (v2 ^ v2), F1: 2}
                                                        _ = v24
                                                        _ = v25
                                                        _ = v28
                                                    } else {
                                                        v29 := GenS2{F0: v2, F1: 2}
                                                        v30 := GenS3{F0: v5, F1: 65535}
                                                        v31 := GenS4{F0: v3, F1: v5}
                                                        var v32 int64 = 364
                                                        _ = len(v1)
                                                        fmt.Println(v6)
                                                        _ = v29
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                }
                                            case "a":
                                                if true {
                                                    var v33 uint = (v3 << (((v3 & (v3 | (^(v3))))) & 63))
                                                    _ = v33
                                                }
                                                if true {
                                                    var v34 int32 = -2147483648
                                                    _ = v34
                                                }
                                                if (100 < -91) {
                                                    if v4 {
                                                        v35 := GenS5{F0: v5}
                                                        v18 = v18
                                                        switch (fi19 / ((fi13) | 1)) {
                                                        case 127:
                                                            switch (^(fi13)) {
                                                            case 847:
                                                                if (!false) {
                                                                    fmt.Println(v6, fi13, v4)
                                                                }
                                                            case 680:
                                                                v36 := GenS6{F0: true, F1: 42}
                                                                if true {
                                                                    v5 += v5
                                                                    fmt.Println(v4, v2, fi13)
                                                                } else {
                                                                    v4 = true
                                                                    switch (((("" + v18) + (("hello" + (("" + v18) + v18)) + "\n\t")) + (v18 + "foo bar")) + "\n\t") {
                                                                    case "a":
                                                                        var v37 uint8 = uint8(v3)
                                                                        _ = v37
                                                                    }
                                                                }
                                                                if false {
                                                                    switch (fi19 % ((v6) | 1)) {
                                                                    case 2147483647:
                                                                        switch (v18 + (((("\n\t" + "foo bar") + v18) + "foo bar") + "")) {
                                                                        case "hello":
                                                                            switch (fi13 << ((v3) & 63)) {
                                                                            case 42:
                                                                                v5 = 1
                                                                            default:
                                                                                switch (("hello" + v18) + v18) {
                                                                                case "\n\t":
                                                                                    v38 := GenS7{F0: (^(v2)), F1: v6}
                                                                                    _ = v38
                                                                                case "":
                                                                                    v6 -= (fi19 + (v6 - -476))
                                                                                case "a":
                                                                                    fmt.Println(v6)
                                                                                }
                                                                                var v39 float32 = (-2.25 * -2.25)
                                                                                var v40 uint16 = 0
                                                                                _ = v39
                                                                                _ = v40
                                                                            }
                                                                            if v4 {
                                                                                if (!v4) {
                                                                                    var v41 int64 = 619
                                                                                    _ = v41
                                                                                }
                                                                            }
                                                                            v42 := GenS8{F0: 795, F1: v3, F2: 2147483648}
                                                                            _ = v42
                                                                        }
                                                                        fmt.Println(fi13)
                                                                        v43 := GenS2{F0: (v2 >> ((v3) & 63)), F1: uint32(v2)}
                                                                        _ = len(v1)
                                                                        v44 := GenS7{F0: v2, F1: v6}
                                                                        _ = v43
                                                                        _ = v44
                                                                    case 1:
                                                                        if v4 {
                                                                            var v45 uint64 = (v2 / ((v2) | 1))
                                                                            var v46 bool = (1 != 124)
                                                                            var v47 float64 = 1e10
                                                                            v6 = (fi19 & (^(fi13)))
                                                                            if v46 {
                                                                                v47 = v47
                                                                            } else {
                                                                                if (!v46) {
                                                                                    var v48 uint64 = (v2 ^ 42)
                                                                                    _ = v48
                                                                                } else {
                                                                                    var v49 uint = 435
                                                                                    _ = v49
                                                                                }
                                                                                if v4 {
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v46)
                                                                                        if v4 {
                                                                                            v47 = v47
                                                                                            v50 := GenS9{F0: ("hello" + (("hello" + ("hello" + v18)) + "hello")), F1: -1, F2: (v18 + (v18 + "a"))}
                                                                                            fmt.Println(v3, v5)
                                                                                            _ = v50
                                                                                        } else {
                                                                                            switch (((v18 + "foo bar") + v18) + (v18 + "foo bar")) {
                                                                                            case "foo bar":
                                                                                                var v51 uint64 = 18446744073709551615
                                                                                                _ = v51
                                                                                            case "a":
                                                                                                var v52 bool = (!(100 >= -1))
                                                                                                _ = v52
                                                                                            case "hello":
                                                                                                v4 = ((0.0 > v47) && (!v4))
                                                                                            }
                                                                                            var v53 int32 = v5
                                                                                            _ = v53
                                                                                        }
                                                                                        switch v18 {
                                                                                        case "foo bar":
                                                                                            fmt.Println(v6, v4, v5)
                                                                                        }
                                                                                    default:
                                                                                        v1 = append(v1, -9223372036854775808)
                                                                                    }
                                                                                } else {
                                                                                    v54 := GenS10{F0: 0, F1: 0, F2: (v5 * (v5 ^ (v5 - v5)))}
                                                                                    var v55 float32 = -2.25
                                                                                    v45 = v2
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                }
                                                                                v18 = ("\n\t" + v18)
                                                                            }
                                                                            v4 = (-1 > -32768)
                                                                            fmt.Println(fi19)
                                                                            _ = v45
                                                                            _ = v46
                                                                            _ = v47
                                                                        } else {
                                                                            v56 := GenS8{F0: 483, F1: v3, F2: 906}
                                                                            if true {
                                                                                _ = len(v1)
                                                                                if (((0.0 * (0.0 - 1.5)) / 1.0) > 3.14159) {
                                                                                    fmt.Println(fi19, v4, v5)
                                                                                }
                                                                                v57 := GenS9{F0: "foo bar", F1: -128, F2: ""}
                                                                                if (int(v3) < fi13) {
                                                                                    var v58 bool = v4
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    var v59 bool = true
                                                                                    var v60 bool = v4
                                                                                    if v60 {
                                                                                        if (!(((v18 + v18) >= "foo bar") && v4)) {
                                                                                            if v58 {
                                                                                                switch (-(v6)) {
                                                                                                case -718:
                                                                                                    v3 -= v3
                                                                                                }
                                                                                                v60 = false
                                                                                            } else {
                                                                                                if v60 {
                                                                                                    fmt.Println(v58, v60)
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                } else {
                                                                                                    v4 = true
                                                                                                    var v61 int32 = v5
                                                                                                    _ = v61
                                                                                                }
                                                                                                if ((0.1 - ((1e10 / 1e10) + 0.0)) >= -2.25) {
                                                                                                    var v62 float32 = -2.25
                                                                                                    _ = v62
                                                                                                }
                                                                                                var v63 uint16 = 278
                                                                                                _ = v63
                                                                                            }
                                                                                        } else {
                                                                                            v64 := GenS11{F0: 113, F1: (^(v3)), F2: 0}
                                                                                            if (-2.25 >= float32(3.14159)) {
                                                                                                switch v18 {
                                                                                                case "":
                                                                                                    v65 := GenS5{F0: (v5 & -487)}
                                                                                                    _ = v65
                                                                                                case "foo bar":
                                                                                                    v66 := GenS12{F0: (^(v3)), F1: 1, F2: (128 != 0)}
                                                                                                    _ = v66
                                                                                                default:
                                                                                                    var v67 int64 = 326
                                                                                                    _ = v67
                                                                                                }
                                                                                                v68 := GenS8{F0: uint32(v2), F1: 42, F2: 65535}
                                                                                                v69 := GenS13{F0: 65535}
                                                                                                _ = v68
                                                                                                _ = v69
                                                                                            } else {
                                                                                                var v70 string = v18
                                                                                                v71 := GenS3{F0: (v5 & 2147483647), F1: 32767}
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                            }
                                                                                            _ = v64
                                                                                        }
                                                                                    }
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                }
                                                                                _ = v57
                                                                            }
                                                                            _ = v56
                                                                        }
                                                                        fmt.Println(fi19, v6)
                                                                        var v72 int16 = -863
                                                                        _ = v72
                                                                    default:
                                                                        var v73 string = "foo bar"
                                                                        _ = v73
                                                                    }
                                                                    v3 -= 0
                                                                } else {
                                                                    v18 = v18
                                                                    _ = len(v1)
                                                                }
                                                                _ = v36
                                                            default:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        default:
                                                            if true {
                                                                v3 = v3
                                                                v18 = v18
                                                                v74 := GenS10{F0: 9223372036854775807, F1: 255, F2: (-(v5))}
                                                                fmt.Println(v6)
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                var v75 int16 = int16(v5)
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                        }
                                                        v76 := GenS4{F0: (v3 + 65535), F1: -727}
                                                        v2 -= (v2 << ((908) & 63))
                                                        var v77 uint64 = v2
                                                        _ = v35
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        v18 = v18
                                                    }
                                                    var v78 uint16 = 255
                                                    v79 := GenS10{F0: 1, F1: v78, F2: -461}
                                                    _ = v78
                                                    _ = v79
                                                }
                                            }
                                        }
                                        _ = fi19
                                    }
                                    v4 = ((0 < 2) || (1 > 2))
                                    _ = v18
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v80 := GenS7{F0: 0, F1: (fi13 - 472)}
                                        _ = v80
                                    case "a":
                                        _ = len(v1)
                                    default:
                                        for fi81 := 0; fi81 < 1; fi81++ {
                                            v5 = (v5 & (v5 + 0))
                                            _ = fi81
                                        }
                                    }
                                }
                            }
                            if v4 {
                                if (1 <= -128) {
                                    fmt.Println(v3, v2)
                                    var v82 int = fi13
                                    fmt.Println(fi13, v3, v3)
                                    _ = v82
                                }
                            }
                            v83 := GenS9{F0: "hello", F1: 52, F2: "\n\t"}
                            _ = v83
                        } else {
                            var v84 int16 = 42
                            v85 := GenS6{F0: true, F1: (v84 + (^(v84)))}
                            _ = v84
                            _ = v85
                        }
                    }
                    var v86 uint8 = 156
                    v87 := GenS14{F0: v86, F1: -1}
                    _ = fi13
                    _ = v86
                    _ = v87
                }
            case "foo bar":
                switch ("foo bar" + "hello") {
                case "foo bar":
                    for fi88 := 0; fi88 < 4; fi88++ {
                        v89 := GenS15{F0: v2, F1: 127}
                        _ = fi88
                        _ = v89
                    }
                case "":
                    var v90 uint8 = 255
                    _ = v90
                case "\n\t":
                    fmt.Println(v3)
                default:
                    var v91 int32 = 42
                    _ = v91
                }
            }
        }
    case "":
        fmt.Println(v2, v2)
    case "foo bar":
        var v92 bool = true
        _ = v92
    default:
        v3 = (v3 - (v3 * v3))
    }
    var v93 int64 = 42
    for fi94 := 0; fi94 < 5; fi94++ {
        for fi95 := 0; fi95 < 13; fi95++ {
            var v96 int16 = -1
            _ = fi95
            _ = v96
        }
        v1 = append(v1, fi94)
        v97 := GenS13{F0: 42}
        _ = fi94
        _ = v97
    }
    v98 := GenS4{F0: (v3 - v3), F1: -1}
    v4 = v4
    fmt.Println(v3, v5)
    if ((2147483647 != 811) && v4) {
        if v4 {
            fmt.Println(v2, v3)
            var v99 int16 = 42
            v2 = v2
            _ = v99
        } else {
            fmt.Println(v3, v93, v4)
        }
    }
    var v100 string = ""
    v100 = v100
    v5 += (v5 | v5)
    var v101 uint8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
    _ = v98
}
