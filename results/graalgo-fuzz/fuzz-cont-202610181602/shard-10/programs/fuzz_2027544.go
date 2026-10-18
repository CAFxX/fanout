package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 string; F2 string }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 uint64; F1 int32; F2 bool }
type GenS9 struct { F0 int16; F1 int16; F2 int }
func genFn0(p0 uint8, p1 int, p2 int64) bool {
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            if false {
                var v9 int16 = -1
                _ = v9
            } else {
                var v10 float32 = 0.0
                v11 := GenS1{F0: 1000, F1: 703, F2: 2147483648}
                _ = v10
                _ = v11
            }
            if false {
                if true {
                    var v12 uint32 = 684
                    var v13 float32 = 1.5
                    if false {
                        var v14 int = (fi7 + 4294967295)
                        if true {
                            v15 := GenS2{F0: -1, F1: (v14 ^ p1)}
                            fmt.Println(p1, v12)
                            _ = v15
                        }
                        if false {
                            var v16 int = (fi8 | 567)
                            if true {
                                fmt.Println(fi7)
                                if (p2 <= (p2 / (((p2 | p2)) | 1))) {
                                    var v17 int32 = -745
                                    p2 -= p2
                                    _ = v17
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        v13 += -2.25
                                    case "a":
                                        if (v12 <= v12) {
                                            if true {
                                                fmt.Println(p0, v16)
                                                if ((p2 % ((p2) | 1)) >= (p2 ^ 0)) {
                                                    p0 = 255
                                                    fmt.Println(v13)
                                                    v18 := GenS3{F0: ("hello" + (("a" + "\n\t") + "foo bar"))}
                                                    _ = v18
                                                } else {
                                                    if false {
                                                        var v19 uint8 = (^(p0))
                                                        _ = v19
                                                    }
                                                    fmt.Println(v12, p0)
                                                    var v20 int8 = 42
                                                    _ = v20
                                                }
                                            }
                                            var v21 bool = true
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v22 uint = 0
                                                _ = v22
                                            case "":
                                                v23 := []int{(fi8 & 2147483647)}
                                                _ = v23
                                            case "hello":
                                                p2 = p2
                                            default:
                                                var v24 uint = 18446744073709551615
                                                _ = v24
                                            }
                                            _ = v21
                                        }
                                    default:
                                        var v25 int64 = p2
                                        _ = v25
                                    }
                                    v26 := GenS2{F0: 32767, F1: (fi7 - 4294967295)}
                                    var v27 int16 = 353
                                    _ = v26
                                    _ = v27
                                }
                                var v28 float64 = -2.25
                                _ = v28
                            } else {
                                var v29 uint = 18446744073709551615
                                switch p1 {
                                case 42:
                                    v30 := []int{fi7, (fi7 % ((p1) | 1))}
                                    _ = v30
                                case -677:
                                    var v31 uint16 = 42
                                    _ = v31
                                default:
                                    var v32 uint16 = 2
                                    _ = v32
                                }
                                p0 -= uint8(v12)
                                _ = v29
                            }
                            _ = v16
                        } else {
                            v33 := GenS4{F0: (false && true)}
                            if (!false) {
                                p0 = p0
                                if (1 != 4294967295) {
                                    var v34 uint64 = 2
                                    var v35 float64 = 0.0
                                    v35 += (3.14159 / (v35 - v35))
                                    _ = v34
                                    _ = v35
                                }
                                v36 := GenS1{F0: 1, F1: 255, F2: (^(v12))}
                                _ = v36
                            } else {
                                if true {
                                    v12 = 1
                                    p1 += (fi8 ^ (^(v14)))
                                } else {
                                    v12 += 0
                                }
                                v37 := GenS5{F0: 2}
                                v13 += v13
                                _ = v37
                            }
                            v38 := []int{(fi7 + -224), 127, (fi7 * p1), fi7}
                            _ = v33
                            _ = v38
                        }
                        _ = v14
                    }
                    _ = v12
                    _ = v13
                }
                if false {
                    p0 = p0
                    if true {
                        switch 128 {
                        case 1:
                            if true {
                                p0 += p0
                            }
                        case -9223372036854775808:
                            var v39 int32 = 953
                            _ = v39
                        default:
                            v40 := GenS6{F0: ("foo bar" + "\n\t"), F1: (("a" + ("\n\t" + "\n\t")) + "a"), F2: "a"}
                            _ = v40
                        }
                        v41 := GenS7{F0: 4294967295, F1: 0}
                        _ = v41
                    } else {
                        p0 -= p0
                        var v42 int64 = p2
                        v43 := []int{(fi8 << ((65535) & 63)), fi8, (fi8 | fi7)}
                        v44 := GenS5{F0: 2}
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                } else {
                    p0 += (p0 & (p0 << ((65535) & 63)))
                    v45 := GenS6{F0: ("a" + ""), F1: "", F2: ""}
                    _ = v45
                }
                var v46 string = "hello"
                if false {
                    if true {
                        switch 425 {
                        case 952:
                            switch (fi8 - fi7) {
                            case 128:
                                p0 = p0
                            case 0:
                                var v47 uint16 = 42
                                _ = v47
                            default:
                                p1 = fi7
                            }
                        default:
                            var v48 float32 = (0.1 - 0.1)
                            _ = v48
                        }
                        v49 := GenS1{F0: -149, F1: 65535, F2: 2147483648}
                        if (570 >= 0) {
                            switch 146 {
                            case -9223372036854775808:
                                p2 = p2
                            case 745:
                                if true {
                                    v46 = v46
                                    fmt.Println(p1)
                                } else {
                                    fmt.Println(p1)
                                    var v50 string = v46
                                    if ((((v46 + v46) + v46) < "\n\t") || (65535 <= 0)) {
                                        switch (v46 + (v50 + ((("foo bar" + "a") + v50) + ("foo bar" + (("\n\t" + "foo bar") + v46))))) {
                                        case "foo bar":
                                            v51 := GenS3{F0: v50}
                                            _ = v51
                                        case "hello":
                                            var v52 int32 = 2147483647
                                            _ = v52
                                        default:
                                            v53 := []int{(p1 / ((42) | 1)), (fi8 * fi7)}
                                            _ = v53
                                        }
                                        var v54 float64 = float64(((-2.25 + 3.14159) + 1.5))
                                        v55 := []int{fi8, fi8, -2147483648, fi7}
                                        p2 = (p2 | (p2 - -1))
                                        _ = v54
                                        _ = v55
                                    }
                                    _ = v50
                                }
                            case -2147483648:
                                v46 = ((v46 + (v46 + (("hello" + "\n\t") + v46))) + ("hello" + (v46 + v46)))
                            }
                            if true {
                                v46 = "a"
                                var v56 uint8 = 255
                                _ = v56
                            }
                        }
                        v57 := GenS6{F0: "foo bar", F1: ((v46 + "") + (((v46 + v46) + v46) + ("" + v46))), F2: ("" + v46)}
                        var v58 int32 = -2147483648
                        _ = v49
                        _ = v57
                        _ = v58
                    }
                    var v59 uint64 = 42
                    var v60 int = fi7
                    _ = v59
                    _ = v60
                } else {
                    switch fi8 {
                    case 1:
                        var v61 uint = 18446744073709551615
                        _ = v61
                    case 127:
                        switch v46 {
                        case "foo bar":
                            v62 := []int{(^(p1)), p1}
                            _ = v62
                        }
                    case -1:
                        if (-739 <= 0) {
                            p1 = (p1 | fi8)
                            fmt.Println(p0, p2)
                        } else {
                            if false {
                                v63 := GenS1{F0: -889, F1: 18446744073709551615, F2: 789}
                                var v64 float32 = 0.1
                                switch (fi8 | fi7) {
                                case 974:
                                    p2 += -1
                                default:
                                    p1 = p1
                                }
                                _ = v63
                                _ = v64
                            }
                        }
                    }
                    switch "hello" {
                    case "foo bar":
                        if false {
                            v65 := GenS3{F0: (v46 + (v46 + ((("\n\t" + (v46 + v46)) + v46) + v46)))}
                            _ = v65
                        }
                    default:
                        var v66 float64 = float64((1e10 / 1.0))
                        _ = v66
                    }
                    var v67 uint8 = 42
                    _ = v67
                }
                fmt.Println(p1)
                _ = v46
            }
            v68 := []int{p1, 2147483647, p1, (fi7 << ((4294967295) & 63))}
            _ = fi8
            _ = v68
        }
        p1 -= fi7
        _ = fi7
    }
    for fi69 := 0; fi69 < 10; fi69++ {
        switch "\n\t" {
        case "foo bar":
            if ((p1 >> ((18446744073709551615) & 63)) >= fi69) {
                p2 += (p2 - p2)
                fmt.Println(fi69, p2, fi69)
                v70 := GenS7{F0: 1, F1: -93}
                _ = v70
            } else {
                p1 += -45
                var v71 uint = 2
                _ = v71
            }
        case "":
            fmt.Println(p2, p0, p0)
        case "\n\t":
            if ((fi69 >> ((18446744073709551615) & 63)) == 32767) {
                var v72 int32 = 100
                _ = v72
            } else {
                switch (p1 + fi69) {
                case -9223372036854775808:
                    if ((float32(0.1) != (0.1 + 1.5)) && (1 <= 0)) {
                        var v73 int32 = -2147483648
                        var v74 float64 = 1e10
                        _ = v73
                        _ = v74
                    } else {
                        if false {
                            switch (p1 / ((2147483647) | 1)) {
                            case 885:
                                for fi75 := 0; fi75 < 17; fi75++ {
                                    p1 -= p1
                                    _ = fi75
                                }
                            }
                        }
                        v76 := GenS1{F0: 0, F1: 0, F2: 42}
                        if true {
                            v77 := []int{}
                            var v78 uint32 = 0
                            _ = v77
                            _ = v78
                        } else {
                            switch (fi69 % ((168) | 1)) {
                            case -9223372036854775808:
                                v79 := []int{fi69, p1, 1}
                                _ = v79
                            case -13:
                                switch p1 {
                                case 127:
                                    p1 = (fi69 * 2147483647)
                                case 2147483647:
                                    p0 += p0
                                case -1:
                                    var v80 uint = 65535
                                    _ = v80
                                default:
                                    p2 = p2
                                }
                            default:
                                v81 := GenS8{F0: 2, F1: -753, F2: ((853 >= -2147483648) && (42 == 482))}
                                _ = v81
                            }
                            if false {
                                p2 = 1
                                p2 -= 42
                            } else {
                                var v82 uint64 = 258
                                v83 := []int{(p1 ^ (fi69 / (((-(p1))) | 1))), p1, -9223372036854775808, (fi69 << ((328) & 63))}
                                v84 := GenS3{F0: "\n\t"}
                                _ = v82
                                _ = v83
                                _ = v84
                            }
                            v85 := []int{4294967295, (fi69 % ((p1) | 1))}
                            _ = v85
                        }
                        _ = v76
                    }
                case 0:
                    v86 := GenS1{F0: -1, F1: 414, F2: 0}
                    _ = v86
                default:
                    var v87 uint16 = 0
                    _ = v87
                }
            }
        }
        _ = fi69
    }
    v88 := []int{(-(p1))}
    if len(v88) >= 2 {
        v88 = v88[1:len(v88)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v88
    return false
}
func genFn1(p0 uint8, p1 int8, p2 uint64) string {
    v89 := []int{42, 287, 0, -782}
    switch 4294967295 {
    case 0:
        if ((false && true) && false) {
            var v90 float64 = 3.14159
            _ = v90
        } else {
            for fi91 := 0; fi91 < 15; fi91++ {
                for fi92 := 0; fi92 < 8; fi92++ {
                    if len(v89) >= 2 {
                        v89 = v89[1:len(v89)]
                    }
                    if len(v89) >= 2 {
                        v89 = v89[1:len(v89)]
                    }
                    if false {
                        var v93 int16 = -32768
                        if len(v89) >= 2 {
                            v89 = v89[1:len(v89)]
                        }
                        v93 = -600
                        v94 := GenS5{F0: (p2 ^ (p2 << ((4294967295) & 63)))}
                        _ = v93
                        _ = v94
                    }
                    _ = fi92
                }
                for fi95 := 0; fi95 < 7; fi95++ {
                    if false {
                        fmt.Println(p1)
                    }
                    _ = fi95
                }
                _ = fi91
            }
            var v96 uint32 = 1
            _ = v96
        }
    case 2147483647:
        _ = len(v89)
    case 31:
        var v97 int64 = 606
        _ = v97
    default:
        for fi98 := 0; fi98 < 18; fi98++ {
            var v99 uint8 = p0
            v100 := GenS4{F0: ((!true) || (330 != 42))}
            if true {
                for fi101 := 0; fi101 < 8; fi101++ {
                    switch "hello" {
                    case "":
                        var v102 string = ""
                        _ = v102
                    case "foo bar":
                        v99 = p0
                    case "hello":
                        _ = len(v89)
                    }
                    _ = fi101
                }
            }
            _ = fi98
            _ = v100
            _ = v99
        }
    }
    if true {
        p1 -= (^(p1))
    }
    var v103 int32 = -927
    _ = p0
    _ = p1
    _ = p2
    _ = v103
    _ = v89
    return "hello"
}

func main() {
    v1 := []int{9223372036854775807, 9223372036854775807, -1}
    var v2 int = 32768
    var v3 int32 = -392
    var v4 string = "hello"
    var v5 int32 = (v3 / (((v3 >> ((0) & 63))) | 1))
    v6 := make(map[int]string)
    var v104 uint64 = 1
    v105 := v6[(^(v2))]
    v106 := GenS8{F0: v104, F1: -1, F2: (v2 == v2)}
    delete(v6, v2)
    v105 = v105
    for fi107 := 0; fi107 < 6; fi107++ {
        for fi108 := 0; fi108 < 17; fi108++ {
            if false {
                if true {
                    var v109 float32 = 1.5
                    _ = v109
                }
            }
            v6[fi107] = "foo bar"
            v3 = (^(v3))
            _ = fi108
        }
        v110 := GenS9{F0: 0, F1: 912, F2: (fi107 - fi107)}
        _ = fi107
        _ = v110
    }
    v2 += (v2 * (-(v2)))
    fmt.Println(v5, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v104
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
