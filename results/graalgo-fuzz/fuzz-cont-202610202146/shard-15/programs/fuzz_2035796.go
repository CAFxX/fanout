package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int32; F1 uint; F2 int16 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 int8; F1 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint; F1 uint32; F2 uint }
type GenS12 struct { F0 int; F1 uint8 }
func genFn0(p0 int64) string {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 string = "hello"
        v7 := GenS1{F0: ((!true) || true), F1: fi5}
        fmt.Println(v6)
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 32767 {
    case 127:
        v8 := GenS2{F0: p0, F1: 42}
        _ = v8
    case 9223372036854775807:
        if false {
            if false {
                if (!false) {
                    v9 := GenS1{F0: true, F1: 1}
                    p0 -= (^(p0))
                    _ = v9
                }
                v10 := []int{-2147483648}
                _ = v10
            } else {
                p0 = 0
            }
            var v11 string = (("\n\t" + ("a" + ("" + "a"))) + "foo bar")
            p0 += p0
            _ = v11
        }
    case -361:
        var v12 int64 = p0
        _ = v12
    default:
        v13 := []int{1, int(p0), 40}
        _ = v13
    }
    switch 9223372036854775807 {
    case 2147483647:
        p0 += p0
    case 0:
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 12; fi15++ {
                switch "\n\t" {
                case "hello":
                    if (!(false && true)) {
                        fmt.Println(p0, fi15, fi14)
                        p0 += -77
                        switch ("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) {
                        case "hello":
                            var v16 int16 = 107
                            var v17 uint16 = 2
                            if false {
                                fmt.Println(fi15)
                                var v18 uint = 0
                                _ = v18
                            }
                            _ = v16
                            _ = v17
                        }
                    }
                case "a":
                    if (!false) {
                        switch "hello" {
                        case "":
                            var v19 string = "\n\t"
                            _ = v19
                        case "foo bar":
                            p0 = 0
                        case "a":
                            v20 := GenS3{F0: p0}
                            _ = v20
                        default:
                            p0 += p0
                        }
                        p0 += (^(p0))
                    } else {
                        switch fi14 {
                        case 1:
                            fmt.Println(p0, fi14)
                        case 2147483647:
                            switch "a" {
                            case "hello":
                                v21 := GenS4{F0: 742, F1: 246}
                                _ = v21
                            case "foo bar":
                                p0 -= (^(p0))
                            case "a":
                                v22 := GenS5{F0: true, F1: ((244 != 1) && true)}
                                _ = v22
                            }
                        default:
                            if (!(!(!true))) {
                                v23 := GenS6{F0: 2147483647, F1: 255, F2: 350}
                                _ = v23
                            }
                        }
                        p0 += p0
                    }
                case "":
                    var v24 uint8 = 166
                    _ = v24
                default:
                    if (!false) {
                        if true {
                            var v25 int16 = 575
                            var v26 string = (("a" + "a") + "a")
                            if false {
                                v27 := []int{}
                                switch (-(fi14)) {
                                case 4294967295:
                                    var v28 string = v26
                                    _ = v28
                                case -309:
                                    var v29 int32 = 96
                                    _ = v29
                                case 421:
                                    if len(v27) > 0 {
                                        v27[0] = 404
                                    }
                                }
                                v30 := GenS7{F0: true, F1: 251}
                                _ = v27
                                _ = v30
                            } else {
                                if true {
                                    var v31 int = (-(fi14))
                                    fmt.Println(v26, fi15, v26)
                                    switch fi14 {
                                    case 2147483647:
                                        var v32 int16 = (v25 / ((42) | 1))
                                        _ = v32
                                    case -2147483648:
                                        p0 -= 779
                                    case 32767:
                                        if false {
                                            fmt.Println(p0, v26)
                                            v33 := GenS8{F0: (v25 | (v25 & (-(v25)))), F1: 4294967295}
                                            v34 := []int{}
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v35 := []int{-1, (fi15 ^ (fi15 / ((823) | 1))), (-(v31)), (^(v31))}
                                            _ = v35
                                        }
                                    }
                                    var v36 float32 = (0.0 - float32(0.1))
                                    v37 := []int{(v31 * int(v25)), (v31 << ((4294967295) & 63))}
                                    if len(v37) >= 2 {
                                        v37 = v37[1:len(v37)]
                                    }
                                    _ = v31
                                    _ = v36
                                    _ = v37
                                }
                                var v38 int8 = 127
                                switch (fi14 % (((-(fi15))) | 1)) {
                                case -583:
                                    v26 = v26
                                default:
                                    var v39 int32 = 0
                                    v40 := []int{fi14, -942}
                                    v41 := GenS9{F0: (v38 >> ((1) & 63)), F1: fi15}
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                }
                                _ = v38
                            }
                            _ = v25
                            _ = v26
                        } else {
                            p0 += p0
                            var v42 int32 = -2147483648
                            _ = v42
                        }
                        p0 = (-(p0))
                    } else {
                        var v43 int32 = 411
                        v44 := []int{}
                        v43 = v43
                        _ = v43
                        _ = v44
                    }
                    if (3.14159 >= 0.1) {
                        var v45 bool = false
                        _ = v45
                    } else {
                        if (1.5 != 1.5) {
                            var v46 int = 1
                            _ = v46
                        } else {
                            if (!true) {
                                var v47 int32 = 2147483647
                                _ = v47
                            }
                        }
                    }
                    if (((1.5 * 1e10) + ((((3.14159 + 1e10) / 1.5) * -2.25) / 0.1)) <= 0.1) {
                        if true {
                            var v48 uint = 1
                            var v49 int16 = 916
                            _ = v48
                            _ = v49
                        } else {
                            if false {
                                switch "hello" {
                                case "\n\t":
                                    v50 := GenS1{F0: false, F1: 128}
                                    _ = v50
                                }
                                if (false || (0.0 < 0.0)) {
                                    var v51 int16 = -1
                                    switch ("hello" + ("hello" + "a")) {
                                    case "\n\t":
                                        switch -9223372036854775808 {
                                        case 9223372036854775807:
                                            v52 := []int{(fi15 % ((fi15) | 1)), (^(fi15)), (-(fi14))}
                                            _ = v52
                                        case 2147483647:
                                            v53 := []int{(fi15 ^ 1), fi14}
                                            _ = v53
                                        default:
                                            var v54 uint8 = 90
                                            v55 := []int{}
                                            _ = v54
                                            _ = v55
                                        }
                                    case "":
                                        v56 := GenS10{F0: 645}
                                        _ = v56
                                    default:
                                        if true {
                                            v51 += v51
                                        } else {
                                            if false {
                                                var v57 int64 = p0
                                                v58 := []int{(^(fi14)), 42, 668, (-(fi15))}
                                                _ = v57
                                                _ = v58
                                            }
                                            if false {
                                                v59 := []int{(-(fi15)), fi14, (-(fi14)), -298}
                                                if len(v59) > 0 {
                                                    v59[0] = fi15
                                                }
                                                fmt.Println(p0)
                                                _ = v59
                                            }
                                            var v60 uint32 = 2
                                            _ = v60
                                        }
                                    }
                                    _ = v51
                                } else {
                                    fmt.Println(p0)
                                }
                                p0 -= p0
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    switch (("foo bar" + "hello") + "foo bar") {
                                    case "foo bar":
                                        p0 = p0
                                    case "hello":
                                        p0 += -353
                                        v61 := []int{(fi15 / (((^(fi14))) | 1))}
                                        _ = v61
                                    case "":
                                        var v62 int32 = 2147483647
                                        _ = v62
                                    }
                                default:
                                    v63 := GenS10{F0: 385}
                                    _ = v63
                                }
                                var v64 uint = 1
                                _ = v64
                            }
                            switch "" {
                            case "a":
                                p0 = p0
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    var v65 uint16 = 42
                                    _ = v65
                                case "foo bar":
                                    var v66 int16 = 1000
                                    _ = v66
                                case "hello":
                                    fmt.Println(fi15)
                                default:
                                    var v67 bool = false
                                    _ = v67
                                }
                            }
                            fmt.Println(p0, p0, p0)
                        }
                        switch "a" {
                        case "\n\t":
                            var v68 int8 = -128
                            _ = v68
                        case "foo bar":
                            var v69 uint16 = 255
                            _ = v69
                        }
                    } else {
                        v70 := []int{1, fi15, fi14}
                        p0 -= p0
                        switch "" {
                        case "foo bar":
                            p0 -= int64(fi14)
                        default:
                            if false {
                                _ = len(v70)
                            }
                        }
                        var v71 uint16 = 42
                        _ = v70
                        _ = v71
                    }
                    v72 := GenS11{F0: 263, F1: 65535, F2: 2}
                    _ = v72
                }
                var v73 uint16 = 2
                p0 -= (p0 / (((p0 | (-(p0)))) | 1))
                _ = fi15
                _ = v73
            }
            _ = fi14
        }
    case 1:
        var v74 uint64 = 130
        _ = v74
    }
    for fi75 := 0; fi75 < 20; fi75++ {
        switch (("hello" + "") + ("\n\t" + ("a" + "a"))) {
        case "\n\t":
            switch (fi75 % (((^(fi75))) | 1)) {
            case 834:
                v76 := []int{(fi75 + fi75), fi75, (fi75 - (-(fi75)))}
                _ = v76
            case -2147483648:
                v77 := []int{fi75, (^(fi75))}
                _ = v77
            case 983:
                v78 := []int{}
                _ = v78
            }
        case "":
            for fi79 := 0; fi79 < 2; fi79++ {
                fmt.Println(fi75)
                _ = fi79
            }
        default:
            v80 := []int{(fi75 ^ fi75), fi75, (fi75 >> ((65535) & 63))}
            _ = v80
        }
        fmt.Println(p0, fi75)
        _ = fi75
    }
    _ = p0
    return "a"
}
func genFn1(p0 uint16, p1 int8) int32 {
    var v81 bool = true
    v82 := []int{}
    _ = p0
    _ = p1
    _ = v81
    _ = v82
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{2147483648}
    var v3 float64 = -2.25
    v4 := make(map[int]bool)
    _ = len(v4)
    for fi83 := 0; fi83 < 8; fi83++ {
        var v84 string = ("a" + "a")
        v85 := GenS12{F0: 127, F1: 0}
        _ = fi83
        _ = v84
        _ = v85
    }
    if (-13 >= 702) {
        if len(v2) > 0 {
            v2[0] = 2
        }
    }
    v86 := GenS9{F0: 100, F1: -9223372036854775808}
    v3 = 1e10
    var v87 uint8 = 42
    var v88 uint32 = 2147483647
    switch -380 {
    case -1:
        var v89 int8 = -1
        _ = v89
    case 56:
        var v90 int32 = genFn1(42, 1)
        _ = v90
    case -157:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v87 = v87
    }
    var v91 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v86
    _ = v87
    _ = v88
    _ = v91
}
