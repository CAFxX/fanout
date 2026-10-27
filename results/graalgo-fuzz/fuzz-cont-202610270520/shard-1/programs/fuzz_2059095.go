package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8; F1 int; F2 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 bool }
func genFn0(p0 int32, p1 string, p2 int32) int64 {
    p0 += 2147483647
    var v7 uint = 536
    p2 -= 398
    p2 -= -1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 uint32, p2 int32) bool {
    var v8 int32 = -749
    v8 -= v8
    fmt.Println(p1)
    if (true && (!true)) {
        var v9 uint8 = 2
        var v10 int = 0
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            if false {
                if true {
                    v12 := GenS1{F0: -1}
                    if ((150 != 1) || (true || (32767 >= 927))) {
                        if (p2 == (p2 << ((2) & 63))) {
                            for fi13 := 0; fi13 < 4; fi13++ {
                                fmt.Println(fi11)
                                _ = fi13
                            }
                        } else {
                            fmt.Println(p1, fi11)
                            v14 := []int{}
                            _ = v14
                        }
                    } else {
                        v15 := []int{(fi11 + 9223372036854775807), 710}
                        var v16 string = (("a" + ((("hello" + "") + "\n\t") + "\n\t")) + "")
                        _ = v15
                        _ = v16
                    }
                    _ = v12
                }
                v17 := GenS2{F0: genFn0(-2147483648, "hello", p2), F1: v8}
                _ = v17
            } else {
                switch "" {
                case "foo bar":
                    v18 := []int{(fi11 & (fi11 & 9223372036854775807))}
                    _ = v18
                case "":
                    v19 := []int{32768, (fi11 ^ (fi11 % ((fi11) | 1)))}
                    _ = v19
                case "a":
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        if (3.14159 > ((1.5 / 1.0) / 1.5)) {
                            for fi20 := 0; fi20 < 6; fi20++ {
                                v21 := []int{}
                                _ = fi20
                                _ = v21
                            }
                            fmt.Println(p0, v8)
                            switch (fi11 * 9223372036854775807) {
                            case 0:
                                var v22 uint8 = 1
                                _ = v22
                            default:
                                switch "\n\t" {
                                case "\n\t":
                                    switch "hello" {
                                    case "hello":
                                        v23 := []int{(fi11 & fi11), -1, int(p0), fi11}
                                        _ = v23
                                    case "foo bar":
                                        v24 := []int{}
                                        _ = v24
                                    case "":
                                        v25 := GenS1{F0: (p2 >> ((42) & 63))}
                                        _ = v25
                                    default:
                                        switch fi11 {
                                        case 492:
                                            v26 := GenS3{F0: 2147483647, F1: 1, F2: 18446744073709551615}
                                            _ = v26
                                        default:
                                            p1 += 65535
                                        }
                                    }
                                case "":
                                    p0 -= (p0 - (p0 ^ 50))
                                default:
                                    for fi27 := 0; fi27 < 4; fi27++ {
                                        p0 -= p0
                                        v28 := []int{fi11, fi27}
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                            }
                        } else {
                            for fi29 := 0; fi29 < 5; fi29++ {
                                v30 := []int{(fi11 | (^(fi29))), -243, (fi11 & fi11)}
                                v31 := GenS4{F0: 127, F1: fi11, F2: "a"}
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            if false {
                                switch "a" {
                                case "":
                                    v32 := GenS5{F0: 127}
                                    _ = v32
                                case "foo bar":
                                    v33 := []int{}
                                    _ = v33
                                case "\n\t":
                                    p0 += (p0 >> ((65535) & 63))
                                default:
                                    if true {
                                        v34 := GenS5{F0: 128}
                                        v35 := GenS6{F0: 42, F1: 531, F2: p1}
                                        _ = v34
                                        _ = v35
                                    } else {
                                        for fi36 := 0; fi36 < 3; fi36++ {
                                            fmt.Println(fi11, p0)
                                            _ = fi36
                                        }
                                    }
                                }
                                var v37 uint16 = 32767
                                _ = v37
                            }
                        }
                    case "":
                        var v38 bool = true
                        _ = v38
                    }
                default:
                    if true {
                        p1 = p1
                    }
                }
                var v39 int16 = 42
                _ = v39
            }
            var v40 bool = true
            v41 := []int{}
            _ = fi11
            _ = v40
            _ = v41
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return false
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]string)
    var v3 uint8 = 2
    var v4 bool = (true && true)
    var v5 uint8 = 128
    v6 := make(map[uint32]string)
    v42 := []int{-221}
    switch 525 {
    case 9223372036854775807:
        var v43 uint64 = 4294967295
        _ = v43
    case 379:
        v44 := GenS3{F0: 4294967295, F1: 4, F2: 18446744073709551615}
        _ = v44
    case 128:
        fmt.Println(v3, v5, v5)
    }
    switch 9223372036854775807 {
    case 1:
        v45 := GenS3{F0: 743, F1: 1, F2: uint64(v3)}
        _ = v45
    case -1:
        v1 -= v1
    }
    v5 += 0
    if (("hello" + "") <= ("" + "")) {
        _ = len(v42)
        var v46 int64 = 914
        _ = v46
    }
    switch -2147483648 {
    case -307:
        v3 = 2
    case 128:
        _ = len(v2)
    default:
        var v47 string = (("hello" + "foo bar") + ((("hello" + "hello") + "\n\t") + "\n\t"))
        _ = v47
    }
    if (42 == 772) {
        if (1.5 < (((float32((0.0 - 1e10)) / (0.1 - 1.5)) + (((0.0 - 0.1) / float32(3.14159)) / (-2.25 / 1e10))) / 1.0)) {
            fmt.Println(v1, v5)
        }
    }
    var v48 int = -2147483648
    _ = len(v2)
    v1 += (-(v1))
    switch "foo bar" {
    case "foo bar":
        v42 = append(v42, v48)
    case "\n\t":
        switch ((("hello" + ("foo bar" + "\n\t")) + "") + "a") {
        case "":
            if ((v5 * v5) == v5) {
                var v49 int32 = 1
                var v50 uint64 = 167
                _ = v49
                _ = v50
            } else {
                if false {
                    var v51 uint8 = v5
                    _ = v51
                }
            }
        case "a":
            fmt.Println(v48, v5)
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                for fi52 := 0; fi52 < 6; fi52++ {
                    if (v4 || (v1 == v1)) {
                        var v53 int8 = -121
                        if v4 {
                            v3 -= v3
                        } else {
                            var v54 bool = (!v4)
                            _ = v54
                        }
                        _ = v53
                    }
                    _ = fi52
                }
            default:
                if v4 {
                    if false {
                        v55 := GenS7{F0: "hello"}
                        v48 = 743
                        _ = v55
                    } else {
                        v56 := GenS8{F0: 42, F1: "hello"}
                        _ = v56
                    }
                    var v57 int = v48
                    v3 = (v5 ^ v5)
                    _ = v57
                } else {
                    v1 += (v1 << ((324) & 63))
                    if (!(v5 >= (v5 | v5))) {
                        var v58 float64 = (0.1 / float64((0.0 - ((3.14159 + 0.0) / 1.5))))
                        switch (v48 & (v48 - v48)) {
                        case -1:
                            var v59 uint64 = 42
                            _ = v59
                        }
                        _ = v58
                    } else {
                        v60 := GenS9{F0: (v48 >> ((18446744073709551615) & 63))}
                        for fi61 := 0; fi61 < 8; fi61++ {
                            for fi62 := 0; fi62 < 12; fi62++ {
                                v4 = v4
                                v63 := v6[979]
                                v48 = 337
                                var v64 int = (v48 + (fi61 / ((42) | 1)))
                                var v65 int16 = v1
                                _ = fi62
                                _ = v63
                                _ = v64
                                _ = v65
                            }
                            v66 := GenS10{F0: v1, F1: (v5 / ((151) | 1))}
                            _ = fi61
                            _ = v66
                        }
                        _ = v60
                    }
                }
            }
        default:
            v1 = 1
        }
    default:
        for fi67 := 0; fi67 < 17; fi67++ {
            for fi68 := 0; fi68 < 11; fi68++ {
                if v4 {
                    v5 += v3
                    v69 := GenS5{F0: v5}
                    if len(v42) > 0 {
                        v42[0] = 138
                    }
                    _ = v69
                }
                if len(v42) >= 2 {
                    v42 = v42[1:len(v42)]
                }
                if (255 <= 2) {
                    var v70 int32 = 2147483647
                    v42 = append(v42, fi68)
                    _ = v70
                }
                _ = fi68
            }
            _ = fi67
        }
    }
    if v4 {
        for fi71 := 0; fi71 < 6; fi71++ {
            v72 := GenS1{F0: 0}
            var v73 int32 = 42
            v48 -= v48
            _ = fi71
            _ = v72
            _ = v73
        }
        for fi74 := 0; fi74 < 8; fi74++ {
            switch "foo bar" {
            case "a":
                if v4 {
                    v5 += 255
                    var v75 float64 = 1.5
                    v48 += (v48 - (v48 | -537))
                    _ = v75
                }
            }
            v3 = (^(v5))
            if true {
                if false {
                    for fi76 := 0; fi76 < 14; fi76++ {
                        v6[2147483648] = (("a" + "\n\t") + "foo bar")
                        v77 := GenS7{F0: "foo bar"}
                        v6[65535] = "\n\t"
                        var v78 int64 = genFn0(2147483647, "hello", -573)
                        _ = fi76
                        _ = v77
                        _ = v78
                    }
                    for fi79 := 0; fi79 < 8; fi79++ {
                        v5 += v5
                        _ = fi79
                    }
                    var v80 int16 = -1
                    _ = v80
                }
                if v4 {
                    var v81 string = ("hello" + (("" + "a") + (("hello" + ("a" + "hello")) + (("\n\t" + (("foo bar" + "hello") + "")) + ((((("a" + "a") + "\n\t") + "\n\t") + (("" + "") + ("foo bar" + "foo bar"))) + "foo bar")))))
                    _ = v81
                }
                v82 := v6[4294967295]
                _ = v82
            } else {
                var v83 int8 = 42
                if v4 {
                    var v84 int32 = int32(fi74)
                    v85 := GenS1{F0: (v84 - (v84 / ((v84) | 1)))}
                    if (!(((("" + "") + ((("hello" + "hello") + "foo bar") + ("hello" + (("a" + (("" + ("" + "foo bar")) + ("\n\t" + "hello"))) + ("\n\t" + ("\n\t" + "hello")))))) + "\n\t") > "a")) {
                        v48 = v48
                    }
                    _ = v84
                    _ = v85
                } else {
                    if v4 {
                        if false {
                            v6[793] = ("a" + ("foo bar" + (("hello" + ((("foo bar" + ("a" + ("" + "a"))) + "hello") + "hello")) + "hello")))
                            v86 := v6[0]
                            fmt.Println(v83)
                            _ = v86
                        }
                        var v87 uint32 = 2147483648
                        _ = v87
                    } else {
                        v88 := v6[2]
                        var v89 int16 = v1
                        _ = v88
                        _ = v89
                    }
                    switch (fi74 + 792) {
                    case -1:
                        if v4 {
                            var v90 float32 = -2.25
                            v1 = (v1 << ((212) & 63))
                            v91 := GenS11{F0: v4}
                            _ = v90
                            _ = v91
                        }
                    case 32768:
                        fmt.Println(v4)
                    case 4294967295:
                        if ((fi74 << ((18446744073709551615) & 63)) > 32768) {
                            v92 := v2[830]
                            _ = v92
                        }
                    }
                    v93 := GenS8{F0: 0, F1: "hello"}
                    _ = v93
                }
                var v94 uint8 = (v3 * v5)
                _ = v83
                _ = v94
            }
            _ = fi74
        }
        var v95 float64 = (float64(3.14159) * (1.5 - 3.14159))
        _ = v95
    }
    for fi96 := 0; fi96 < 7; fi96++ {
        var v97 string = ((("\n\t" + "hello") + "hello") + "foo bar")
        v5 += (v3 >> ((1) & 63))
        v98 := GenS3{F0: 0, F1: 632, F2: 4294967295}
        _ = fi96
        _ = v97
        _ = v98
    }
    v5 -= 1
    for fi99 := 0; fi99 < 1; fi99++ {
        v100 := GenS9{F0: 2147483647}
        _ = fi99
        _ = v100
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v48
    _ = v5
    _ = v6
}
