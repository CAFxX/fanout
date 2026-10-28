package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool; F1 uint }
type GenS11 struct { F0 int32; F1 uint16; F2 int }
type GenS12 struct { F0 int16; F1 int8; F2 uint64 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int16 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 int8 = -48
    v5 := make(map[int]string)
    var v6 bool = false
    if v6 {
        var v7 int64 = 9223372036854775807
        var v8 uint = 4294967295
        fmt.Println(v7, v1)
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: 0}
    switch "a" {
    case "":
        var v10 float64 = 3.14159
        _ = v10
    case "a":
        fmt.Println(v1, v6, v4)
    case "hello":
        switch 161 {
        case 0:
            v4 += -128
        case -512:
            if true {
                v6 = true
                switch 752 {
                case 0:
                    if v6 {
                        v11 := []int{9223372036854775807, -400}
                        _ = v11
                    }
                case 9223372036854775807:
                    v12 := GenS1{F0: 1000, F1: 42}
                    _ = v12
                case 128:
                    if (v1 <= v1) {
                        for fi13 := 0; fi13 < 16; fi13++ {
                            v14 := GenS2{F0: -1, F1: 5}
                            if v6 {
                                var v15 bool = (!v6)
                                var v16 uint32 = 666
                                v15 = false
                                _ = v15
                                _ = v16
                            }
                            _ = fi13
                            _ = v14
                        }
                        v17 := GenS3{F0: 1000, F1: 1}
                        var v18 float32 = v1
                        _ = v17
                        _ = v18
                    } else {
                        v19 := GenS2{F0: v4, F1: 255}
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v21 := GenS4{F0: 0}
                            switch (("\n\t" + "a") + "hello") {
                            case "hello":
                                var v22 int64 = -69
                                _ = v22
                            case "foo bar":
                                v23 := GenS3{F0: 42, F1: 0}
                                _ = v23
                            case "\n\t":
                                delete(v3, (("" + "\n\t") + "a"))
                            default:
                                v4 = (^(v4))
                            }
                            _ = fi20
                            _ = v21
                        }
                        if (v6 || true) {
                            for fi24 := 0; fi24 < 6; fi24++ {
                                if true {
                                    for fi25 := 0; fi25 < 15; fi25++ {
                                        var v26 float32 = -2.25
                                        _ = fi25
                                        _ = v26
                                    }
                                    if ((true && (!v6)) && (!(!v6))) {
                                        v27 := v5[(fi24 / ((fi24) | 1))]
                                        v28 := GenS5{F0: 42, F1: (!v6)}
                                        _ = v27
                                        _ = v28
                                    }
                                } else {
                                    fmt.Println(v4, v6)
                                    v29 := GenS1{F0: 1000, F1: 18446744073709551615}
                                    v30 := []int{}
                                    _ = v29
                                    _ = v30
                                }
                                _ = fi24
                            }
                            for fi31 := 0; fi31 < 1; fi31++ {
                                var v32 int64 = 0
                                v4 = 41
                                if (("a" == (((("hello" + "hello") + "a") + ("hello" + "")) + "\n\t")) && (-939 > 42)) {
                                    if v6 {
                                        for fi33 := 0; fi33 < 17; fi33++ {
                                            v34 := GenS6{F0: 42, F1: -2147483648}
                                            v35 := GenS4{F0: 881}
                                            _ = fi33
                                            _ = v34
                                            _ = v35
                                        }
                                    } else {
                                        v36 := []int{fi31, (^(fi31)), (^(fi31))}
                                        fmt.Println(fi31, v6)
                                        delete(v5, fi31)
                                        _ = v36
                                    }
                                } else {
                                    if (false && true) {
                                        var v37 int32 = 1
                                        for fi38 := 0; fi38 < 13; fi38++ {
                                            var v39 int16 = -1
                                            var v40 int64 = v32
                                            _ = fi38
                                            _ = v39
                                            _ = v40
                                        }
                                        for fi41 := 0; fi41 < 1; fi41++ {
                                            v42 := []int{fi31, fi31, fi31}
                                            if (!v6) {
                                                v42 = append(v42, fi31)
                                                v43 := v2[(fi41 * (fi31 - fi41))]
                                                switch "foo bar" {
                                                case "":
                                                    var v44 int8 = (v4 & (v4 << ((1) & 63)))
                                                    _ = v44
                                                case "hello":
                                                    delete(v2, fi31)
                                                default:
                                                    v45 := GenS1{F0: 1000, F1: 400}
                                                    _ = v45
                                                }
                                                _ = v43
                                            }
                                            _ = fi41
                                            _ = v42
                                        }
                                        _ = v37
                                    } else {
                                        var v46 float32 = (v1 / (((((v1 + -2.25) * v1) * v1) + v1) * (0.1 + 0.1)))
                                        _ = v46
                                    }
                                    v47 := v3[""]
                                    for fi48 := 0; fi48 < 14; fi48++ {
                                        var v49 int16 = 1
                                        v50 := []int{1, fi31, -448}
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                    _ = v47
                                }
                                _ = fi31
                                _ = v32
                            }
                            for fi51 := 0; fi51 < 5; fi51++ {
                                v4 = v4
                                for fi52 := 0; fi52 < 9; fi52++ {
                                    _ = len(v5)
                                    if true {
                                        if (fi51 <= (-(fi52))) {
                                            v53 := []int{}
                                            v54 := GenS7{F0: 42, F1: ("hello" + "hello")}
                                            v55 := GenS4{F0: 1}
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                        v4 -= v4
                                        v56 := v5[fi52]
                                        _ = v56
                                    } else {
                                        fmt.Println(v4, fi51)
                                        v6 = true
                                        if false {
                                            var v57 int16 = 42
                                            v1 += 0.0
                                            fmt.Println(fi51, v57)
                                            _ = v57
                                        }
                                        switch -2147483648 {
                                        case -1:
                                            var v58 int64 = 0
                                            _ = v58
                                        default:
                                            v59 := GenS8{F0: 1, F1: (-(fi51))}
                                            _ = v59
                                        }
                                    }
                                    var v60 int64 = -205
                                    _ = fi52
                                    _ = v60
                                }
                                v61 := []int{770, (fi51 / (((fi51 - (-(fi51)))) | 1)), 32767}
                                _ = fi51
                                _ = v61
                            }
                        }
                        _ = v19
                    }
                default:
                    v62 := []int{}
                    _ = v62
                }
            } else {
                fmt.Println(v1, v1)
            }
        case 127:
            v6 = (true || ((3.14159 <= -2.25) || (1 != 97)))
        default:
            for fi63 := 0; fi63 < 9; fi63++ {
                var v64 string = "hello"
                switch "foo bar" {
                case "a":
                    var v65 int = (fi63 | 0)
                    _ = v65
                }
                _ = fi63
                _ = v64
            }
        }
    }
    var v66 int8 = (-(v4))
    var v67 float64 = 1e10
    var v68 int8 = v66
    v6 = ((404 == 0) || (!true))
    v69 := v2[128]
    var v70 int16 = 32767
    var v71 bool = (!(42 < 152))
    v67 -= (1e10 + (1.5 / float64(v1)))
    for fi72 := 0; fi72 < 18; fi72++ {
        var v73 int32 = -1
        for fi74 := 0; fi74 < 18; fi74++ {
            v69 = (v70 < v70)
            if v71 {
                if (v6 || v69) {
                    v75 := GenS9{F0: -1}
                    v76 := []int{}
                    switch fi72 {
                    case 328:
                        var v77 float32 = (1.5 - 0.0)
                        _ = v77
                    case 128:
                        v6 = false
                    default:
                        v69 = v6
                    }
                    if ((v71 || false) && true) {
                        if (!true) {
                            switch 32767 {
                            case -163:
                                if (fi72 <= -158) {
                                    v78 := GenS8{F0: 255, F1: fi72}
                                    if ((v66 + (v68 | v68)) >= (^(v66))) {
                                        var v79 int8 = 1
                                        var v80 uint32 = 201
                                        v81 := GenS10{F0: v69, F1: 65535}
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                    }
                                    _ = v78
                                }
                            case 383:
                                v71 = (((fi74 / ((4294967295) | 1)) > fi72) || v71)
                            case 289:
                                v5[fi72] = "\n\t"
                            }
                        }
                        switch (fi72 / ((-10) | 1)) {
                        case 307:
                            fmt.Println(v1, v70)
                        case 42:
                            v82 := GenS11{F0: (v73 / ((-168) | 1)), F1: 738, F2: fi74}
                            _ = v82
                        case -933:
                            if v6 {
                                var v83 bool = true
                                v76 = append(v76, (fi74 / ((892) | 1)))
                                _ = v83
                            } else {
                                v3["\n\t"] = "a"
                                switch "a" {
                                case "":
                                    _ = len(v3)
                                case "foo bar":
                                    v4 += 127
                                case "\n\t":
                                    if v71 {
                                        switch "foo bar" {
                                        case "hello":
                                            if true {
                                                v84 := v2[fi74]
                                                var v85 float64 = (float64(v1) - ((1.5 + v67) + ((v67 + 3.14159) * (1.5 + 1.5))))
                                                switch "\n\t" {
                                                case "a":
                                                    switch "\n\t" {
                                                    case "":
                                                        _ = len(v2)
                                                    case "\n\t":
                                                        v86 := GenS11{F0: (-(v73)), F1: 0, F2: 228}
                                                        _ = v86
                                                    case "foo bar":
                                                        v66 = (v66 ^ (v68 << ((4294967295) & 63)))
                                                    }
                                                case "foo bar":
                                                    switch ("foo bar" + "hello") {
                                                    case "a":
                                                        v3[("a" + ("a" + "hello"))] = "\n\t"
                                                        fmt.Println(v85)
                                                        v87 := GenS10{F0: false, F1: 852}
                                                        fmt.Println(v4, v71)
                                                        if v6 {
                                                            var v88 int16 = (v70 - (v70 + v70))
                                                            v89 := v3["a"]
                                                            if v84 {
                                                                v73 += -1
                                                                var v90 int8 = v4
                                                                _ = v90
                                                            }
                                                            _ = v88
                                                            _ = v89
                                                        } else {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                fmt.Println(v4)
                                                                v1 = (v1 * 1.5)
                                                            case "a":
                                                                var v91 int16 = (-(v70))
                                                                _ = v91
                                                            }
                                                            v76 = append(v76, (fi72 + fi74))
                                                            fmt.Println(v68, fi74)
                                                        }
                                                        _ = v87
                                                    case "":
                                                        var v92 bool = (-2.25 != -2.25)
                                                        _ = v92
                                                    default:
                                                        switch fi74 {
                                                        case -1:
                                                            var v93 bool = (127 <= 2)
                                                            _ = v93
                                                        case -839:
                                                            v94 := GenS12{F0: (v70 / ((v70) | 1)), F1: (v68 & 0), F2: 2}
                                                            _ = v94
                                                        case 0:
                                                            v95 := GenS13{F0: ((("hello" + "foo bar") + "a") + ((("" + "a") + "") + "\n\t"))}
                                                            _ = v95
                                                        }
                                                    }
                                                case "hello":
                                                    v68 = 127
                                                default:
                                                    fmt.Println(v85, v70, v68)
                                                }
                                                fmt.Println(v67)
                                                _ = v84
                                                _ = v85
                                            }
                                        default:
                                            var v96 string = ("" + ("a" + ("" + "hello")))
                                            _ = v96
                                        }
                                    } else {
                                        v67 -= 1.5
                                    }
                                }
                                v97 := GenS14{F0: 42}
                                _ = v97
                            }
                        }
                        v76 = append(v76, (fi74 * (fi72 % ((fi74) | 1))))
                    } else {
                        if (((fi74 / ((0) | 1)) == fi74) || true) {
                            fmt.Println(v73)
                            v98 := GenS10{F0: v69, F1: 65535}
                            var v99 float64 = (v67 * (3.14159 - -2.25))
                            _ = v98
                            _ = v99
                        } else {
                            var v100 int8 = 42
                            v67 = float64(v1)
                            var v101 uint8 = 117
                            _ = v100
                            _ = v101
                        }
                    }
                    _ = v75
                    _ = v76
                }
                v1 = v1
                fmt.Println(v70, v73, v70)
            }
            v102 := GenS15{F0: (^(v70))}
            _ = fi74
            _ = v102
        }
        _ = fi72
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v9
}
