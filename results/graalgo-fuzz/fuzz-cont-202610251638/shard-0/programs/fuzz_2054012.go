package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 string; F1 int8; F2 uint }
type GenS10 struct { F0 uint; F1 int; F2 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint64) int16 {
    switch "\n\t" {
    case "a":
        if ((!true) || false) {
            var v7 float32 = (float32(-2.25) * ((((3.14159 + (3.14159 / 1.0)) / -2.25) - float32((1e10 * (((1.5 + -2.25) - -2.25) + -2.25)))) - (0.0 + (((1e10 - (float32((0.1 + 0.0)) * (1.5 - float32((3.14159 * ((-2.25 / (((-2.25 + ((((float64((float32(((-2.25 - 3.14159) * 0.1)) / ((1.5 + float32(1e10)) + (1.5 + (1e10 + -2.25))))) * 0.0) * -2.25) / 1e10) + ((0.0 + 1.5) + 0.0))) / 1e10) - (0.0 + float64((3.14159 / 1.0))))) + (1.5 + 1.5))))))) + float32(float64(0.1))) / (((3.14159 + 1e10) + -2.25) * 0.1)))))
            p0 = p1
            v8 := []int{183, -1, 404}
            _ = v7
            _ = v8
        }
    case "foo bar":
        if true {
            switch "" {
            case "\n\t":
                v9 := GenS1{F0: p0, F1: 2147483647}
                _ = v9
            default:
                p1 += 2
            }
            p1 = 18446744073709551615
        }
    case "\n\t":
        var v10 int16 = 0
        _ = v10
    }
    v11 := GenS2{F0: 112, F1: 712}
    p0 = p0
    v12 := GenS3{F0: int32(p0)}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := []string{"a", "hello"}
    v2 := make(map[string]int)
    var v3 string = "foo bar"
    v4 := []uint32{2147483647}
    var v5 int16 = -32768
    var v6 float32 = 3.14159
    if (-111 > 100) {
        v13 := GenS4{F0: 2, F1: (v5 ^ v5)}
        switch 32767 {
        case 42:
            switch v3 {
            case "a":
                if true {
                    v5 -= (v5 + (v5 >> ((249) & 63)))
                    v3 = ""
                } else {
                    switch -9223372036854775808 {
                    case -1:
                        var v14 uint = 42
                        _ = v14
                    case 32767:
                        v15 := GenS5{F0: 127, F1: 255}
                        _ = v15
                    default:
                        var v16 int8 = 26
                        _ = v16
                    }
                }
            default:
                if true {
                    delete(v2, v3)
                    v17 := GenS3{F0: -1}
                    _ = v17
                } else {
                    if false {
                        fmt.Println(v5, v6)
                    } else {
                        fmt.Println(v6, v3, v3)
                        v18 := GenS6{F0: 9223372036854775807}
                        var v19 uint8 = 42
                        _ = v18
                        _ = v19
                    }
                    delete(v2, "foo bar")
                }
            }
        case 2147483647:
            var v20 int16 = genFn0(4294967295, 0)
            _ = v20
        }
        fmt.Println(v5, v6)
        _ = v13
    } else {
        for fi21 := 0; fi21 < 10; fi21++ {
            if false {
                if len(v1) > 0 {
                    v1[0] = v3
                }
                for fi22 := 0; fi22 < 11; fi22++ {
                    fmt.Println(v3, fi22)
                    var v23 bool = (1 >= -1)
                    switch -2147483648 {
                    case 347:
                        var v24 string = v3
                        _ = v24
                    case 70:
                        v3 = v3
                    case -736:
                        v3 = v3
                    }
                    _ = fi22
                    _ = v23
                }
            }
            _ = fi21
        }
        if (128 != -2147483648) {
            _ = len(v2)
            var v25 float32 = (0.1 - v6)
            v26 := GenS2{F0: -1, F1: 0}
            _ = v25
            _ = v26
        }
        var v27 float64 = (float64((1.5 - (v6 * ((v6 + float32(float64(v6))) + (1.5 / -2.25))))) - -2.25)
        _ = v27
    }
    var v28 float32 = v6
    v29 := GenS7{F0: 2147483648}
    var v30 int16 = v5
    v28 -= (1.5 / v28)
    if true {
        var v31 string = ""
        switch -867 {
        case -678:
            var v32 int64 = 1
            _ = v32
        case 128:
            v6 = (v6 / 3.14159)
        }
        _ = v31
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        switch "" {
        case "hello":
            switch v3 {
            case "hello":
                v34 := GenS7{F0: 42}
                _ = v34
            case "a":
                switch fi33 {
                case 42:
                    var v35 uint = 18446744073709551615
                    _ = v35
                default:
                    v36 := GenS1{F0: 2, F1: 2}
                    _ = v36
                }
            case "\n\t":
                for fi37 := 0; fi37 < 4; fi37++ {
                    fmt.Println(fi33)
                    var v38 uint8 = 2
                    _ = fi37
                    _ = v38
                }
            default:
                var v39 int8 = 100
                _ = v39
            }
        default:
            var v40 int16 = genFn0(831, 286)
            _ = v40
        }
        v41 := GenS3{F0: 2147483647}
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 int64 = -1
            fmt.Println(v30)
            delete(v2, "a")
            _ = fi42
            _ = v43
        }
        _ = fi33
        _ = v41
    }
    v6 = 0.0
    for fi44 := 0; fi44 < 7; fi44++ {
        var v45 int8 = -1
        if (v3 <= (((v3 + (("" + v3) + v3)) + v3) + v3)) {
            var v46 int32 = 1
            var v47 int = (fi44 % (((fi44 ^ fi44)) | 1))
            v46 += (v46 ^ 2147483647)
            _ = v46
            _ = v47
        }
        for fi48 := 0; fi48 < 8; fi48++ {
            v30 -= (v30 % ((v30) | 1))
            fmt.Println(fi48)
            _ = fi48
        }
        _ = fi44
        _ = v45
    }
    var v49 int = 128
    fmt.Println(v30, v49)
    if ((42 <= 1) && true) {
        v50 := GenS4{F0: 0, F1: v30}
        _ = v50
    } else {
        v1 = append(v1, v3)
        if (-88 != -1) {
            for fi51 := 0; fi51 < 16; fi51++ {
                if true {
                    for fi52 := 0; fi52 < 15; fi52++ {
                        v6 = 1.5
                        _ = fi52
                    }
                } else {
                    v49 -= v49
                    fmt.Println(v28, v5)
                    for fi53 := 0; fi53 < 6; fi53++ {
                        switch fi51 {
                        case 763:
                            fmt.Println(v6, v3, v5)
                        }
                        var v54 bool = false
                        var v55 float32 = ((v28 + (v28 * 1.5)) / ((1.5 - 1.5) / v28))
                        switch v3 {
                        case "foo bar":
                            switch "hello" {
                            case "foo bar":
                                v55 = -2.25
                            case "a":
                                v56 := GenS6{F0: fi51}
                                _ = v56
                            case "\n\t":
                                v57 := v2[((("a" + v3) + ("hello" + ((("\n\t" + v3) + v3) + v3))) + v3)]
                                _ = v57
                            default:
                                var v58 uint32 = 0
                                _ = v58
                            }
                        case "a":
                            v59 := GenS5{F0: 255, F1: 0}
                            _ = v59
                        }
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                }
                if false {
                    if false {
                        if false {
                            var v60 uint8 = 2
                            for fi61 := 0; fi61 < 17; fi61++ {
                                if true {
                                    var v62 uint64 = 18446744073709551615
                                    var v63 uint8 = (v60 << ((274) & 63))
                                    _ = v62
                                    _ = v63
                                } else {
                                    v64 := GenS8{F0: (v30 * v5), F1: -128}
                                    _ = len(v2)
                                    _ = v64
                                }
                                _ = fi61
                            }
                            switch (fi51 | 27) {
                            case 9223372036854775807:
                                _ = len(v4)
                            }
                            _ = v60
                        } else {
                            v65 := GenS4{F0: 289, F1: v5}
                            if true {
                                _ = len(v1)
                                switch (v3 + ("hello" + v3)) {
                                case "":
                                    var v66 float32 = (-2.25 - float32((float64(float32(float64((((-2.25 - v6) / (float32(0.0) * ((((v28 - 0.0) + (1e10 + (1e10 * 3.14159))) / 1.5) + 1e10))) + 0.1)))) * 3.14159)))
                                    _ = v66
                                case "foo bar":
                                    if (false || (fi51 == fi51)) {
                                        var v67 uint32 = 58
                                        v49 = fi51
                                        if false {
                                            var v68 int32 = -2147483648
                                            v30 = genFn0(977, 4294967295)
                                            var v69 uint16 = 32767
                                            _ = v68
                                            _ = v69
                                        }
                                        _ = v67
                                    }
                                default:
                                    var v70 int = fi51
                                    _ = v70
                                }
                                switch v3 {
                                case "a":
                                    for fi71 := 0; fi71 < 5; fi71++ {
                                        var v72 int32 = -410
                                        v30 += 1000
                                        if false {
                                            switch (fi71 / ((fi71) | 1)) {
                                            case 1:
                                                var v73 uint = 0
                                                _ = v73
                                            case 0:
                                                v5 = v5
                                            default:
                                                v74 := GenS4{F0: 42, F1: genFn0(706, 332)}
                                                _ = v74
                                            }
                                            _ = len(v2)
                                        } else {
                                            fmt.Println(v6, v6, v3)
                                            var v75 uint64 = 18446744073709551615
                                            var v76 string = v3
                                            _ = v75
                                            _ = v76
                                        }
                                        _ = fi71
                                        _ = v72
                                    }
                                case "":
                                    switch (v49 + -19) {
                                    case 42:
                                        v5 = 1
                                    default:
                                        if true {
                                            var v77 int8 = -1
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v77
                                        } else {
                                            v78 := GenS6{F0: (fi51 * (fi51 * 1))}
                                            for fi79 := 0; fi79 < 15; fi79++ {
                                                v80 := GenS2{F0: fi51, F1: 255}
                                                fmt.Println(v5, v3, fi79)
                                                _ = fi79
                                                _ = v80
                                            }
                                            _ = v78
                                        }
                                    }
                                }
                            }
                            var v81 uint8 = 0
                            _ = v65
                            _ = v81
                        }
                        var v82 int8 = 4
                        _ = v82
                    } else {
                        var v83 float32 = ((v6 / v6) * (v28 + (0.0 * (float32(1e10) * v6))))
                        var v84 string = v3
                        switch v84 {
                        case "\n\t":
                            var v85 uint16 = 255
                            _ = v85
                        case "hello":
                            fmt.Println(v3)
                        case "":
                            v86 := GenS9{F0: "a", F1: 50, F2: 2}
                            _ = v86
                        default:
                            var v87 uint8 = 127
                            _ = v87
                        }
                        _ = v83
                        _ = v84
                    }
                } else {
                    var v88 int64 = 0
                    v4 = append(v4, 2147483648)
                    _ = v88
                }
                var v89 int = 32767
                _ = fi51
                _ = v89
            }
            v49 += (v49 - v49)
            switch (v3 + "a") {
            case "":
                if (1 != -20) {
                    switch ((v3 + "foo bar") + "foo bar") {
                    case "foo bar":
                        v28 -= v6
                    }
                    for fi90 := 0; fi90 < 9; fi90++ {
                        if true {
                            var v91 bool = (false && (1 >= fi90))
                            v30 = (^(v30))
                            var v92 int = fi90
                            _ = v91
                            _ = v92
                        } else {
                            if (!false) {
                                fmt.Println(v5)
                                var v93 int32 = 2147483647
                                _ = v93
                            } else {
                                if ((-(v5)) != genFn0(4294967295, 0)) {
                                    if true {
                                        _ = len(v2)
                                        if false {
                                            if true {
                                                for fi94 := 0; fi94 < 15; fi94++ {
                                                    v95 := GenS10{F0: 4294967295, F1: (^(fi90)), F2: false}
                                                    fmt.Println(v30, v28, v3)
                                                    _ = fi94
                                                    _ = v95
                                                }
                                                var v96 int8 = 1
                                                _ = v96
                                            }
                                            v28 = (v6 / float32((float64(v28) - 1.5)))
                                            v28 -= -2.25
                                        } else {
                                            var v97 float64 = 3.14159
                                            for fi98 := 0; fi98 < 6; fi98++ {
                                                switch v49 {
                                                case 297:
                                                    if false {
                                                        var v99 int16 = 621
                                                        var v100 bool = (fi90 >= fi98)
                                                        v101 := GenS5{F0: 128, F1: 128}
                                                        _ = v100
                                                        _ = v101
                                                        _ = v99
                                                    }
                                                case -1:
                                                    var v102 uint8 = 255
                                                    _ = v102
                                                }
                                                _ = fi98
                                            }
                                            var v103 int64 = 9223372036854775807
                                            _ = v103
                                            _ = v97
                                        }
                                        v2[v3] = 42
                                    } else {
                                        v104 := GenS11{F0: -1}
                                        switch ("" + ("a" + "")) {
                                        case "hello":
                                            _ = len(v2)
                                        case "\n\t":
                                            v28 = 0.1
                                        default:
                                            v5 = (v5 / ((v5) | 1))
                                        }
                                        if (true && true) {
                                            v6 -= 1e10
                                        }
                                        _ = v104
                                    }
                                    var v105 int8 = -1
                                    if true {
                                        v3 = ("hello" + v3)
                                    } else {
                                        v28 -= 0.0
                                        for fi106 := 0; fi106 < 2; fi106++ {
                                            v28 -= v28
                                            v49 -= (^(fi106))
                                            v49 = fi90
                                            _ = fi106
                                        }
                                    }
                                    _ = v105
                                } else {
                                    _ = len(v4)
                                }
                                for fi107 := 0; fi107 < 14; fi107++ {
                                    switch 986 {
                                    case 4294967295:
                                        switch (fi90 << ((18446744073709551615) & 63)) {
                                        case -2147483648:
                                            v108 := GenS8{F0: genFn0(42, 1), F1: -128}
                                            _ = v108
                                        }
                                    case 76:
                                        if true {
                                            v49 -= (fi107 % ((fi90) | 1))
                                            fmt.Println(v28, v49, fi90)
                                        } else {
                                            if false {
                                                v6 = float32(1.5)
                                                var v109 int = 0
                                                _ = v109
                                            } else {
                                                var v110 float64 = (-2.25 / -2.25)
                                                _ = v110
                                            }
                                        }
                                    default:
                                        fmt.Println(fi90)
                                    }
                                    var v111 uint = 65535
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi107
                                    _ = v111
                                }
                            }
                        }
                        _ = fi90
                    }
                    if (!(false || (!true))) {
                        v112 := GenS4{F0: 4294967295, F1: 42}
                        for fi113 := 0; fi113 < 14; fi113++ {
                            var v114 uint16 = 65535
                            var v115 uint64 = 1
                            _ = fi113
                            _ = v114
                            _ = v115
                        }
                        _ = v112
                    } else {
                        for fi116 := 0; fi116 < 15; fi116++ {
                            v117 := GenS12{F0: 947}
                            _ = fi116
                            _ = v117
                        }
                    }
                }
            case "a":
                var v118 float64 = -2.25
                _ = v118
            default:
                v119 := GenS10{F0: 255, F1: 9223372036854775807, F2: (1 < 185)}
                _ = v119
            }
        }
        var v120 uint = 0
        _ = v120
    }
    var v121 string = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v121:", v121)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v121
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
