package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int8; F1 uint32; F2 int8 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 uint; F1 uint }

func main() {
    v1 := make(map[int]int)
    var v2 string = "a"
    var v3 float32 = 0.1
    var v4 uint32 = 2147483647
    var v5 float64 = 1e10
    v6 := make(map[string]bool)
    v1[9223372036854775807] = -9223372036854775808
    var v7 float64 = v5
    var v8 uint32 = v4
    v9 := GenS1{F0: ((!true) || (42 >= 0)), F1: 4294967295, F2: -1}
    var v10 uint64 = 171
    v4 = (v8 * v4)
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            fmt.Println(v8, v10, v4)
            v5 += float64(((0.1 / v3) * (v3 / float32(float64(v3)))))
            v5 = (v7 * 0.0)
            _ = fi12
        }
        delete(v6, (v2 + (v2 + v2)))
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := []int{(fi13 >> ((683) & 63))}
            _ = fi13
            _ = v14
        }
        _ = fi11
    }
    if false {
        if (!false) {
            if false {
                fmt.Println(v5, v3, v4)
                if (129 == -778) {
                    if true {
                        if true {
                            v15 := v6[("" + ((v2 + (v2 + "")) + ("\n\t" + (v2 + v2))))]
                            for fi16 := 0; fi16 < 9; fi16++ {
                                switch (^(fi16)) {
                                case 171:
                                    v1[570] = fi16
                                case 2147483647:
                                    v7 -= v7
                                default:
                                    var v17 bool = v15
                                    _ = v17
                                }
                                switch (fi16 * fi16) {
                                case -2147483648:
                                    for fi18 := 0; fi18 < 17; fi18++ {
                                        v15 = ((!(true || false)) || true)
                                        switch fi18 {
                                        case 9223372036854775807:
                                            v3 -= ((v3 + v3) - (0.1 * ((v3 / float32(v7)) + (float32(-2.25) / v3))))
                                        case -9223372036854775808:
                                            v19 := GenS1{F0: ((1 < -32768) || v15), F1: 65535, F2: -32768}
                                            _ = v19
                                        }
                                        if v15 {
                                            fmt.Println(v8)
                                            switch v2 {
                                            case "\n\t":
                                                if false {
                                                    var v20 uint8 = 128
                                                    switch fi16 {
                                                    case 338:
                                                        v8 = v4
                                                    default:
                                                        var v21 uint = 322
                                                        _ = v21
                                                    }
                                                    v22 := GenS1{F0: (!true), F1: 4294967295, F2: -1}
                                                    _ = v20
                                                    _ = v22
                                                } else {
                                                    v15 = (v15 && (592 == 42))
                                                    var v23 float64 = v7
                                                    _ = v23
                                                }
                                            case "a":
                                                var v24 bool = ((v2 >= "hello") || false)
                                                _ = v24
                                            case "foo bar":
                                                v25 := GenS1{F0: (!false), F1: 560, F2: 42}
                                                _ = v25
                                            }
                                            if (0 == uint(v10)) {
                                                v26 := GenS1{F0: (v15 && true), F1: uint(v10), F2: -1}
                                                fmt.Println(v15)
                                                v4 = (v4 + v8)
                                                _ = v26
                                            }
                                        } else {
                                            if (0 >= 0) {
                                                v27 := GenS2{F0: (fi18 ^ 1)}
                                                v28 := GenS1{F0: (v2 != (v2 + ("a" + ("\n\t" + v2)))), F1: 255, F2: -32768}
                                                v29 := GenS1{F0: v15, F1: 18446744073709551615, F2: 1}
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            } else {
                                                var v30 int8 = 42
                                                var v31 float32 = 3.14159
                                                switch (fi16 / ((127) | 1)) {
                                                case -52:
                                                    v10 += 2
                                                case 128:
                                                    if true {
                                                        var v32 float32 = 1.5
                                                        if (!(0 > -32768)) {
                                                            _ = len(v6)
                                                            v7 = ((-2.25 - v5) + 3.14159)
                                                            _ = len(v6)
                                                        } else {
                                                            if true {
                                                                v32 -= v3
                                                                v31 = v3
                                                            }
                                                            var v33 int = (fi16 >> ((65535) & 63))
                                                            var v34 uint16 = 0
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        _ = v32
                                                    } else {
                                                        _ = len(v6)
                                                    }
                                                case -9223372036854775808:
                                                    delete(v6, v2)
                                                }
                                                _ = v30
                                                _ = v31
                                            }
                                            fmt.Println(v8, v8, v15)
                                        }
                                        v35 := GenS2{F0: int(v4)}
                                        if (!true) {
                                            var v36 int = (fi18 << ((1) & 63))
                                            _ = v36
                                        } else {
                                            delete(v1, (fi18 + 127))
                                            var v37 string = (("a" + "\n\t") + v2)
                                            if (114 <= -54) {
                                                v10 -= v10
                                            }
                                            _ = v37
                                        }
                                        var v38 uint = 255
                                        _ = fi18
                                        _ = v35
                                        _ = v38
                                    }
                                default:
                                    v39 := GenS1{F0: v15, F1: 255, F2: -264}
                                    _ = v39
                                }
                                _ = fi16
                            }
                            _ = v15
                        }
                        v40 := []int{int(v10), -991, 32767, 9223372036854775807}
                        _ = v40
                    } else {
                        v41 := []int{2147483647}
                        var v42 float64 = v7
                        _ = v41
                        _ = v42
                    }
                    switch v2 {
                    case "":
                        if true {
                            if (42 == 42) {
                                var v43 uint8 = 2
                                fmt.Println(v7, v43, v10)
                                _ = v43
                            }
                            var v44 string = ((v2 + v2) + v2)
                            for fi45 := 0; fi45 < 8; fi45++ {
                                for fi46 := 0; fi46 < 7; fi46++ {
                                    v7 += v7
                                    _ = fi46
                                }
                                _ = fi45
                            }
                            _ = v44
                        } else {
                            v2 = v2
                            var v47 uint16 = 2
                            _ = v47
                        }
                    case "foo bar":
                        v3 = v3
                    case "hello":
                        v48 := GenS3{F0: 4294967295, F1: 8}
                        _ = v48
                    default:
                        switch v2 {
                        case "a":
                            for fi49 := 0; fi49 < 4; fi49++ {
                                for fi50 := 0; fi50 < 8; fi50++ {
                                    if (float64((3.14159 - ((1e10 / 0.1) * v3))) >= v7) {
                                        var v51 int8 = 127
                                        v52 := []int{128, fi50, (fi50 ^ fi50), fi49}
                                        if len(v52) >= 2 {
                                            v52 = v52[1:len(v52)]
                                        }
                                        _ = v51
                                        _ = v52
                                    }
                                    var v53 float32 = v3
                                    _ = fi50
                                    _ = v53
                                }
                                if (v10 == 535) {
                                    fmt.Println(v2)
                                    for fi54 := 0; fi54 < 19; fi54++ {
                                        switch v2 {
                                        case "hello":
                                            if false {
                                                v10 += (v10 * 18446744073709551615)
                                            } else {
                                                v10 = (v10 << ((18446744073709551615) & 63))
                                            }
                                        default:
                                            var v55 uint8 = 42
                                            _ = v55
                                        }
                                        v56 := GenS1{F0: true, F1: 666, F2: 1000}
                                        v4 = (v8 << ((877) & 63))
                                        _ = fi54
                                        _ = v56
                                    }
                                    var v57 int16 = 934
                                    _ = v57
                                }
                                _ = fi49
                            }
                        case "\n\t":
                            v8 += (v8 - (v8 | 2))
                        default:
                            if (!(1 < 1)) {
                                switch 1 {
                                case -431:
                                    fmt.Println(v2, v3)
                                case 2147483647:
                                    if false {
                                        v10 -= 174
                                        switch 32768 {
                                        case 289:
                                            for fi58 := 0; fi58 < 9; fi58++ {
                                                for fi59 := 0; fi59 < 9; fi59++ {
                                                    if true {
                                                        if false {
                                                            if false {
                                                                if false {
                                                                    if (108 != (v10 - (v10 % ((v10) | 1)))) {
                                                                        fmt.Println(v2, fi59, fi59)
                                                                        v3 -= v3
                                                                    } else {
                                                                        v60 := GenS4{F0: 12, F1: v4, F2: 0}
                                                                        _ = v60
                                                                    }
                                                                }
                                                            }
                                                            delete(v1, (fi58 - (fi59 >> ((255) & 63))))
                                                            if ((fi58 % ((fi59) | 1)) == fi58) {
                                                                if (!(!(!(false && true)))) {
                                                                    v61 := GenS4{F0: 100, F1: (v4 - (v4 % ((v8) | 1))), F2: 1}
                                                                    _ = v61
                                                                } else {
                                                                    var v62 bool = (182 >= 255)
                                                                    v4 += 0
                                                                    v63 := []int{160, fi58}
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                                v5 = 1e10
                                                            } else {
                                                                switch "" {
                                                                case "foo bar":
                                                                    if (-32768 <= 232) {
                                                                        if (!true) {
                                                                            var v64 float32 = 0.0
                                                                            v65 := v6[v2]
                                                                            v66 := GenS5{F0: 0, F1: 65535}
                                                                            _ = v64
                                                                            _ = v65
                                                                            _ = v66
                                                                        } else {
                                                                            switch "a" {
                                                                            case "hello":
                                                                                v4 = v8
                                                                            case "":
                                                                                var v67 uint8 = 1
                                                                                _ = v67
                                                                            default:
                                                                                var v68 int8 = -93
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                        v10 = v10
                                                                    } else {
                                                                        var v69 uint32 = v8
                                                                        var v70 int32 = -574
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                case "hello":
                                                                    v71 := []int{(-(fi59))}
                                                                    _ = v71
                                                                default:
                                                                    v10 += v10
                                                                }
                                                                if false {
                                                                    var v72 string = v2
                                                                    fmt.Println(v8, v2)
                                                                    _ = v72
                                                                }
                                                                delete(v1, fi59)
                                                            }
                                                        } else {
                                                            v73 := GenS4{F0: 100, F1: (v4 - v8), F2: 91}
                                                            v2 = ""
                                                            v3 = (-2.25 * v3)
                                                            _ = v73
                                                        }
                                                        var v74 bool = (v2 <= (("" + "") + "foo bar"))
                                                        v75 := []int{4294967295, fi59, 1, 9223372036854775807}
                                                        _ = v74
                                                        _ = v75
                                                    } else {
                                                        var v76 int8 = -128
                                                        v2 = (("foo bar" + (("\n\t" + v2) + "hello")) + (("hello" + "a") + "foo bar"))
                                                        if true {
                                                            switch (v2 + (v2 + "\n\t")) {
                                                            case "":
                                                                v8 -= 375
                                                                if ((v76 & -5) <= (v76 * -1)) {
                                                                    v76 += (^(v76))
                                                                    v5 = 1e10
                                                                    fmt.Println(v8)
                                                                }
                                                            case "a":
                                                                v76 = 42
                                                            }
                                                        }
                                                        fmt.Println(v5, v7, v5)
                                                        v5 = 0.0
                                                        v77 := v1[fi58]
                                                        v3 += v3
                                                        delete(v1, fi59)
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    if false {
                                                        fmt.Println(v10, v2, fi58)
                                                        v3 = 3.14159
                                                        var v78 int32 = 1
                                                        _ = v78
                                                    }
                                                    _ = fi59
                                                }
                                                var v79 int64 = 42
                                                _ = fi58
                                                _ = v79
                                            }
                                        }
                                    } else {
                                        v80 := v1[643]
                                        delete(v6, "")
                                        _ = v80
                                    }
                                }
                                for fi81 := 0; fi81 < 11; fi81++ {
                                    var v82 bool = ((v5 + -2.25) > v5)
                                    _ = fi81
                                    _ = v82
                                }
                            }
                        }
                    }
                    v83 := GenS5{F0: -2147483648, F1: 255}
                    _ = v83
                } else {
                    v6["a"] = true
                    fmt.Println(v8, v3)
                }
            }
        } else {
            v84 := GenS6{F0: 42}
            _ = v84
        }
    } else {
        v3 += (v3 / v3)
        v10 -= v10
    }
    v85 := []int{637}
    for fi86 := 0; fi86 < 1; fi86++ {
        if (-1 >= -2147483648) {
            fmt.Println(v8, v10, fi86)
            v2 = "foo bar"
            v3 = (((v3 - 0.1) / (v3 / 1.5)) * (1e10 + v3))
        } else {
            if true {
                v3 -= v3
                if false {
                    v4 = v8
                    v10 -= (v10 >> ((29) & 63))
                    var v87 int32 = -58
                    _ = v87
                } else {
                    if false {
                        if len(v85) >= 2 {
                            v85 = v85[1:len(v85)]
                        }
                        if false {
                            v3 = ((v3 / 1.0) - 0.1)
                            var v88 float64 = (0.0 - v5)
                            _ = v88
                        } else {
                            v2 = ("" + "")
                            v7 = v5
                            for fi89 := 0; fi89 < 6; fi89++ {
                                fmt.Println(v10, v8, fi86)
                                v1[(fi86 + fi89)] = fi86
                                if false {
                                    v8 -= 951
                                }
                                _ = fi89
                            }
                        }
                    } else {
                        var v90 int64 = 671
                        v5 = 1e10
                        for fi91 := 0; fi91 < 20; fi91++ {
                            var v92 int16 = -230
                            v93 := GenS7{F0: (v10 - (^(v10))), F1: v8}
                            _ = fi91
                            _ = v92
                            _ = v93
                        }
                        _ = v90
                    }
                    v94 := GenS4{F0: -1, F1: (v4 << ((42) & 63)), F2: -128}
                    _ = v94
                }
            } else {
                fmt.Println(v10, v5)
            }
            v5 += (3.14159 / 1e10)
            var v95 float64 = (1.5 + v7)
            _ = v95
        }
        v2 = v2
        _ = fi86
    }
    v6[v2] = (true && (v2 > v2))
    v2 = v2
    var v96 int16 = -358
    var v97 int64 = -9223372036854775808
    if (-128 >= -128) {
        v2 = ("hello" + (("hello" + v2) + v2))
        var v98 float64 = ((v7 - v7) + ((v5 - 1e10) * 1.5))
        var v99 int32 = -2147483648
        _ = v98
        _ = v99
    } else {
        v85 = append(v85, -112)
        if len(v85) > 0 {
            v85[0] = -2147483648
        }
        for fi100 := 0; fi100 < 11; fi100++ {
            var v101 int8 = 0
            var v102 uint64 = (v10 - (v10 * 2))
            v103 := GenS8{F0: 42, F1: 1}
            _ = fi100
            _ = v101
            _ = v102
            _ = v103
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v85:", len(v85), v85)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v85
    _ = v9
    _ = v96
    _ = v97
}
