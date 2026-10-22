package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint8; F1 string; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint; F2 int16 }
type GenS8 struct { F0 uint16; F1 string }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 uint; F1 string; F2 string }
type GenS11 struct { F0 uint32; F1 uint64 }
type GenS12 struct { F0 int8; F1 bool }
type GenS13 struct { F0 int16; F1 uint; F2 int }
func genFn0() int32 {
    var v6 string = ("hello" + ("" + ("hello" + "a")))
    switch v6 {
    case "foo bar":
        fmt.Println(v6)
    case "":
        if true {
            switch "foo bar" {
            case "":
                if (false && true) {
                    v6 = ""
                    v7 := GenS1{F0: ((-1 <= 2147483647) && true)}
                    _ = v7
                } else {
                    if (808 >= 255) {
                        var v8 int16 = -1
                        v8 = v8
                        if false {
                            v9 := GenS2{F0: 2, F1: 42}
                            _ = v9
                        } else {
                            var v10 int32 = 2147483647
                            _ = v10
                        }
                        _ = v8
                    }
                    v11 := []int{492}
                    v6 = ""
                    _ = v11
                }
            case "a":
                v12 := GenS3{F0: 4294967295, F1: -128}
                _ = v12
            default:
                if (0.0 < (0.0 - 0.0)) {
                    fmt.Println(v6)
                } else {
                    if false {
                        v6 = "hello"
                    }
                    for fi13 := 0; fi13 < 19; fi13++ {
                        var v14 string = "hello"
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS4{F0: 123, F1: "hello", F2: 0}
                    _ = v15
                }
            }
            var v16 int = 42
            if (false || true) {
                v17 := []int{}
                _ = v17
            }
            _ = v16
        }
    case "\n\t":
        if ((1.5 / 1.0) < 3.14159) {
            for fi18 := 0; fi18 < 12; fi18++ {
                for fi19 := 0; fi19 < 15; fi19++ {
                    switch (v6 + "hello") {
                    case "a":
                        if (((((float64((1.5 - float32(1.5))) / (-2.25 / 1e10)) - ((3.14159 * -2.25) / float64(0.1))) * (1.5 - float64(0.1))) < 1.5) && false) {
                            var v20 uint = 1
                            v21 := GenS1{F0: true}
                            if false {
                                v22 := GenS3{F0: 205, F1: 1}
                                switch fi18 {
                                case 1:
                                    v23 := []int{(fi19 | fi19), (-(fi19)), (fi19 / ((fi18) | 1))}
                                    _ = v23
                                default:
                                    var v24 uint8 = 42
                                    _ = v24
                                }
                                v25 := GenS5{F0: "\n\t"}
                                v6 = v6
                                v26 := []int{}
                                v6 = (("\n\t" + (v6 + "hello")) + "hello")
                                if (-1 >= 9223372036854775807) {
                                    fmt.Println(v20)
                                    if (!true) {
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        fmt.Println(fi18, v20, fi18)
                                    } else {
                                        var v27 int = fi19
                                        _ = v27
                                    }
                                }
                                _ = v22
                                _ = v25
                                _ = v26
                            }
                            _ = v20
                            _ = v21
                        }
                    case "hello":
                        var v28 string = v6
                        _ = v28
                    default:
                        if false {
                            switch fi18 {
                            case 128:
                                var v29 bool = (false || true)
                                _ = v29
                            case 906:
                                v6 = v6
                            default:
                                v30 := GenS5{F0: v6}
                                _ = v30
                            }
                            switch 0 {
                            case 127:
                                if true {
                                    switch "" {
                                    case "hello":
                                        if (!false) {
                                            if false {
                                                switch (v6 + (v6 + (v6 + v6))) {
                                                case "foo bar":
                                                    fmt.Println(fi18, v6, v6)
                                                default:
                                                    v6 = "\n\t"
                                                }
                                                var v31 uint32 = 4294967295
                                                v6 = v6
                                                v32 := GenS1{F0: (v31 != 428)}
                                                _ = v31
                                                _ = v32
                                            }
                                            fmt.Println(fi19)
                                            v6 = (v6 + "foo bar")
                                        } else {
                                            switch (v6 + v6) {
                                            case "\n\t":
                                                var v33 string = (v6 + "\n\t")
                                                var v34 uint16 = 266
                                                var v35 uint8 = 222
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            case "a":
                                                switch fi19 {
                                                case 4294967295:
                                                    var v36 float64 = -2.25
                                                    _ = v36
                                                case 461:
                                                    v6 = v6
                                                case 335:
                                                    var v37 int16 = -564
                                                    _ = v37
                                                default:
                                                    v38 := GenS1{F0: true}
                                                    _ = v38
                                                }
                                            case "":
                                                v39 := []int{(fi18 & fi19), (fi18 * 128), fi19, fi18}
                                                fmt.Println(v6)
                                                v6 = (v6 + "")
                                                _ = v39
                                            }
                                            var v40 uint64 = 703
                                            _ = v40
                                        }
                                    case "foo bar":
                                        v6 = v6
                                    case "":
                                        switch fi19 {
                                        case 607:
                                            v41 := []int{}
                                            _ = v41
                                        case 127:
                                            var v42 int8 = 100
                                            _ = v42
                                        default:
                                            v6 = v6
                                        }
                                    }
                                }
                            case 128:
                                fmt.Println(v6)
                            case 568:
                                switch fi18 {
                                case 128:
                                    v6 = "foo bar"
                                default:
                                    v43 := GenS5{F0: v6}
                                    _ = v43
                                }
                            default:
                                var v44 int64 = 174
                                v44 = (^(v44))
                                _ = v44
                            }
                            v6 = v6
                            if false {
                                v45 := []int{fi18}
                                _ = v45
                            }
                        }
                        var v46 uint16 = 65535
                        v47 := []int{(fi18 << ((2) & 63)), (fi19 - 127), fi18, (fi18 ^ 42)}
                        _ = v46
                        _ = v47
                    }
                    switch (fi19 - (^(fi18))) {
                    case 42:
                        v48 := GenS4{F0: 2, F1: v6, F2: -1}
                        _ = v48
                    case -1:
                        fmt.Println(v6, fi19, fi18)
                    case -286:
                        if false {
                            fmt.Println(fi19, fi18)
                            v6 = ("a" + v6)
                        } else {
                            switch (fi18 % ((fi19) | 1)) {
                            case 32767:
                                v6 = v6
                            }
                        }
                    default:
                        v49 := GenS4{F0: 42, F1: "foo bar", F2: 1}
                        _ = v49
                    }
                    var v50 uint = 165
                    _ = fi19
                    _ = v50
                }
                var v51 uint64 = 1
                v52 := []int{(fi18 | fi18)}
                _ = fi18
                _ = v51
                _ = v52
            }
        } else {
            switch 4294967295 {
            case -546:
                v53 := []int{9223372036854775807}
                _ = v53
            default:
                fmt.Println(v6)
            }
        }
    default:
        v6 = ""
    }
    _ = v6
    return -1
}

func main() {
    v1 := []uint32{1, 180, 326, 248}
    var v2 float32 = (1e10 - 0.1)
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    v5 := []uint32{220}
    v54 := GenS5{F0: "a"}
    v55 := GenS4{F0: 0, F1: ("a" + ("" + ("hello" + "\n\t"))), F2: -1}
    var v56 int64 = -1
    if (255 > 42) {
        for fi57 := 0; fi57 < 5; fi57++ {
            v2 = 0.1
            v58 := GenS1{F0: (2147483647 >= 2147483647)}
            v59 := GenS6{F0: 1000}
            _ = fi57
            _ = v58
            _ = v59
        }
        _ = len(v3)
    } else {
        fmt.Println(v56, v2)
        v56 = v56
        v60 := GenS7{F0: -128, F1: 750, F2: -32768}
        _ = v60
    }
    if true {
        v56 += (^(v56))
        var v61 int64 = v56
        _ = v61
    } else {
        v2 = 0.1
    }
    v2 -= (3.14159 * 3.14159)
    switch -628 {
    case 530:
        fmt.Println(v2, v56)
    case 0:
        v62 := GenS8{F0: 42, F1: "a"}
        _ = v62
    default:
        if false {
            for fi63 := 0; fi63 < 10; fi63++ {
                delete(v4, ("a" + "a"))
                if (!true) {
                    v56 = -907
                    v2 -= v2
                    fmt.Println(v2, fi63, fi63)
                }
                fmt.Println(v56, v2, v2)
                _ = fi63
            }
            v5 = append(v5, 1)
            var v64 int32 = genFn0()
            _ = v64
        } else {
            for fi65 := 0; fi65 < 13; fi65++ {
                _ = len(v3)
                v2 -= ((3.14159 / 1.0) - float32((0.0 + (1e10 - ((-2.25 / 1.0) - float64(1e10))))))
                v66 := GenS1{F0: false}
                _ = fi65
                _ = v66
            }
            if (!true) {
                if true {
                    var v67 int16 = 863
                    if true {
                        fmt.Println(v67)
                        v2 = 0.0
                    }
                    _ = v67
                }
            }
            var v68 int16 = 1000
            _ = v68
        }
    }
    v2 = v2
    var v69 int64 = (v56 + 9223372036854775807)
    if true {
        v69 = (v69 | v69)
        v70 := GenS4{F0: 42, F1: ("a" + ((("\n\t" + (("a" + "\n\t") + ("" + "a"))) + ("a" + "a")) + "foo bar")), F2: -1}
        var v71 uint16 = 32767
        _ = v70
        _ = v71
    } else {
        var v72 uint32 = 1
        var v73 uint64 = 483
        for fi74 := 0; fi74 < 3; fi74++ {
            for fi75 := 0; fi75 < 13; fi75++ {
                var v76 uint = 65535
                var v77 int16 = 0
                if false {
                    var v78 float32 = v2
                    _ = v78
                }
                _ = fi75
                _ = v76
                _ = v77
            }
            v5 = append(v5, (v72 ^ (^(v72))))
            _ = fi74
        }
        _ = v72
        _ = v73
    }
    if true {
        if true {
            v79 := GenS6{F0: 32767}
            _ = v79
        } else {
            switch ("" + "foo bar") {
            case "hello":
                var v80 uint = 46
                _ = v80
            case "":
                v2 = -2.25
            case "foo bar":
                for fi81 := 0; fi81 < 20; fi81++ {
                    if true {
                        v82 := GenS3{F0: 645, F1: -46}
                        var v83 float64 = (1e10 * (1e10 / (float64((v2 / (v2 / v2))) * float64((1.5 * v2)))))
                        switch (fi81 & fi81) {
                        case -1:
                            v56 = 255
                        }
                        _ = v82
                        _ = v83
                    } else {
                        if true {
                            switch "a" {
                            case "hello":
                                if (79 >= 457) {
                                    if (fi81 > 9223372036854775807) {
                                        v84 := GenS9{F0: 2147483648, F1: 42}
                                        fmt.Println(fi81, v69, v69)
                                        _ = v84
                                    } else {
                                        v56 += v56
                                    }
                                } else {
                                    for fi85 := 0; fi85 < 13; fi85++ {
                                        var v86 uint16 = 2
                                        v87 := GenS10{F0: 614, F1: "\n\t", F2: "\n\t"}
                                        _ = fi85
                                        _ = v86
                                        _ = v87
                                    }
                                    v56 = 0
                                    var v88 int16 = 32767
                                    _ = v88
                                }
                            }
                        }
                    }
                    for fi89 := 0; fi89 < 1; fi89++ {
                        var v90 int = -2147483648
                        var v91 uint8 = 2
                        fmt.Println(v56)
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                    v92 := GenS7{F0: 127, F1: 244, F2: 42}
                    _ = fi81
                    _ = v92
                }
            default:
                var v93 int16 = -1
                _ = v93
            }
            v5 = append(v5, 528)
        }
        switch "\n\t" {
        case "a":
            v94 := GenS11{F0: 1, F1: 4294967295}
            _ = v94
        case "hello":
            var v95 float32 = (float32((((-2.25 * 0.1) - float64((3.14159 + ((1.5 * float32(3.14159)) * 1e10)))) - ((1e10 - 1e10) * 0.1))) - (v2 - v2))
            _ = v95
        case "":
            if (genFn0() > 2147483647) {
                if true {
                    var v96 bool = (115 >= 42)
                    if v96 {
                        v97 := GenS7{F0: 1, F1: 42, F2: 42}
                        fmt.Println(v2, v96)
                        _ = v97
                    }
                    var v98 uint64 = 42
                    _ = v96
                    _ = v98
                }
                var v99 int16 = 0
                _ = v99
            } else {
                if false {
                    v56 += -1
                }
                fmt.Println(v2)
                var v100 uint32 = 2
                _ = v100
            }
        }
    } else {
        if (179 == 368) {
            var v101 int16 = 1
            _ = v101
        } else {
            if true {
                var v102 uint8 = 255
                _ = v102
            }
            v103 := GenS5{F0: "a"}
            var v104 int16 = 1000
            _ = v103
            _ = v104
        }
        v56 -= (v69 << ((776) & 63))
    }
    if (true && false) {
        v105 := GenS12{F0: 34, F1: true}
        v106 := GenS3{F0: 42, F1: 100}
        fmt.Println(v56, v2)
        _ = v105
        _ = v106
    } else {
        v69 -= v69
        var v107 uint16 = 42
        var v108 int16 = -775
        _ = v107
        _ = v108
    }
    if (42 == 582) {
        switch "foo bar" {
        case "a":
            for fi109 := 0; fi109 < 9; fi109++ {
                var v110 bool = (1 != 1)
                _ = fi109
                _ = v110
            }
        case "hello":
            for fi111 := 0; fi111 < 7; fi111++ {
                fmt.Println(v56, v69, v56)
                var v112 uint32 = 0
                _ = fi111
                _ = v112
            }
        case "":
            for fi113 := 0; fi113 < 7; fi113++ {
                if ((-776 > 0) || false) {
                    v56 = (v69 & (v56 % ((1) | 1)))
                    v56 -= (v56 ^ (-(v69)))
                    fmt.Println(fi113, fi113, v56)
                } else {
                    switch 2147483647 {
                    case 42:
                        v56 += (v56 / ((-1) | 1))
                    case 128:
                        var v114 uint16 = 42
                        _ = v114
                    case -9223372036854775808:
                        v56 += v56
                    }
                }
                fmt.Println(v69, v56, fi113)
                v4["a"] = 2147483648
                _ = fi113
            }
        default:
            v2 -= ((v2 - v2) + float32(0.0))
        }
    } else {
        _ = len(v4)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v2 -= v2
    if false {
        var v115 string = (("hello" + "foo bar") + "")
        var v116 uint16 = 32767
        _ = v115
        _ = v116
    } else {
        var v117 uint32 = 2147483648
        fmt.Println(v69)
        v117 = (v117 - 65535)
        _ = v117
    }
    v118 := GenS13{F0: 42, F1: 0, F2: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v69
}
