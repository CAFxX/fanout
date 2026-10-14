package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 int; F1 uint; F2 int32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64; F1 int64; F2 int }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int32; F1 uint16; F2 uint }
type GenS13 struct { F0 uint; F1 uint; F2 uint64 }
type GenS14 struct { F0 int16 }
func genFn0(p0 string) int {
    for fi7 := 0; fi7 < 19; fi7++ {
        switch ("" + ((p0 + p0) + ((p0 + p0) + "foo bar"))) {
        case "hello":
            p0 = p0
        }
        v8 := []int{(fi7 ^ 127), fi7, (fi7 * (fi7 ^ fi7))}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: -1, F2: 2147483648}
    v10 := []int{32768}
    _ = p0
    _ = v10
    _ = v9
    return -963
}
func genFn1(p0 int16, p1 bool) uint64 {
    fmt.Println(p0, p0)
    var v11 bool = ((255 > 65535) && p1)
    var v12 string = ("a" + "")
    var v13 int32 = 43
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := []int{32768, 246, 40, 42}
    var v2 int = -1
    var v3 int16 = -1
    var v4 int8 = 42
    var v5 int8 = (v4 | 100)
    v6 := make(map[string]uint32)
    if true {
        switch "\n\t" {
        case "hello":
            if true {
                v2 -= 32768
            } else {
                if true {
                    v2 = v2
                    if (!(false || true)) {
                        v5 += v5
                    } else {
                        var v14 int32 = -893
                        _ = v14
                    }
                } else {
                    switch (v2 % ((v2) | 1)) {
                    case -1:
                        if (!(-60 == -95)) {
                            v15 := GenS2{F0: (v2 | (v2 % ((v2) | 1))), F1: 4294967295, F2: 42}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 3; fi16++ {
                                fmt.Println(v5, fi16, v3)
                                v4 += (v5 << ((18446744073709551615) & 63))
                                v1 = append(v1, (-(v2)))
                                _ = fi16
                            }
                            v17 := GenS1{F0: v4, F1: -1, F2: 812}
                            _ = v17
                        }
                    case 67:
                        v5 += 1
                    default:
                        switch (-(v2)) {
                        case 127:
                            v3 = -1
                        default:
                            if (0.1 <= -2.25) {
                                fmt.Println(v4, v5)
                                if ((v2 / ((-284) | 1)) < v2) {
                                    v18 := v6[""]
                                    v5 -= int8(v18)
                                    _ = v18
                                }
                                v3 -= (v3 | v3)
                            } else {
                                v19 := GenS1{F0: 1, F1: 42, F2: 2}
                                _ = v19
                            }
                        }
                    }
                }
                if (v4 != v4) {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        for fi21 := 0; fi21 < 4; fi21++ {
                            _ = len(v6)
                            var v22 uint = 65535
                            _ = fi21
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 3; fi23++ {
                            var v24 bool = true
                            var v25 int = -161
                            v26 := GenS3{F0: 42, F1: "hello"}
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 18; fi27++ {
                            var v28 uint32 = 0
                            fmt.Println(v2, v28, v28)
                            _ = fi27
                            _ = v28
                        }
                        _ = fi20
                    }
                    var v29 int32 = -953
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        if (!(true || false)) {
                            var v31 uint16 = 723
                            v32 := v6[("foo bar" + "hello")]
                            _ = v31
                            _ = v32
                        } else {
                            fmt.Println(fi30)
                            _ = len(v1)
                            v33 := v6["a"]
                            _ = v33
                        }
                        v6["foo bar"] = 567
                        fmt.Println(fi30, v4, v2)
                        _ = fi30
                    }
                    switch "a" {
                    case "a":
                        var v34 int8 = (v5 + (v4 << ((65535) & 63)))
                        _ = v34
                    case "":
                        fmt.Println(v5, v5, v5)
                    case "foo bar":
                        fmt.Println(v2, v3, v4)
                    }
                    delete(v6, "hello")
                }
            }
        default:
            switch v2 {
            case -80:
                fmt.Println(v2)
            case 4294967295:
                v3 = (v3 | 32767)
            }
        }
    }
    if ((!(((255 != 42) || (32767 >= 255)) && false)) || true) {
        fmt.Println(v5, v4, v5)
    } else {
        delete(v6, "")
        if false {
            var v35 float64 = (((1.5 + (0.0 - 0.1)) - 1.5) * 0.1)
            _ = v35
        }
        v36 := v6[(("a" + (("\n\t" + "hello") + (((("" + "a") + ("foo bar" + (("" + "hello") + ((("foo bar" + "hello") + ("a" + ("" + "hello"))) + ("a" + ""))))) + ("" + (("\n\t" + "") + "\n\t"))) + ((((("\n\t" + "\n\t") + (("hello" + "\n\t") + "\n\t")) + "a") + "\n\t") + "a")))) + "foo bar")]
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        if ((("\n\t" + "hello") + ("a" + "a")) <= ("a" + "a")) {
            for fi38 := 0; fi38 < 13; fi38++ {
                v2 -= fi37
                _ = fi38
            }
            v39 := GenS2{F0: -2147483648, F1: 1, F2: 283}
            v1 = append(v1, fi37)
            _ = v39
        }
        for fi40 := 0; fi40 < 9; fi40++ {
            if (!(!true)) {
                var v41 uint = 42
                v2 += 32767
                var v42 int64 = 0
                _ = v41
                _ = v42
            }
            if (1 <= 42) {
                v43 := GenS4{F0: 65, F1: false}
                v44 := GenS5{F0: int64(v3), F1: ""}
                fmt.Println(fi40, v3, v2)
                v45 := GenS6{F0: 0}
                _ = v43
                _ = v44
                _ = v45
            } else {
                fmt.Println(v5, fi40, v2)
                var v46 int = fi40
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                _ = v46
            }
            _ = fi40
        }
        _ = fi37
    }
    var v47 uint16 = 1
    v6["foo bar"] = 42
    v2 = v2
    switch (v2 ^ (-(v2))) {
    case 9223372036854775807:
        v4 = -125
    case 42:
        if true {
            if ((v3 + (v3 >> ((307) & 63))) < (-(v3))) {
                if (2 != 128) {
                    var v48 uint64 = 18446744073709551615
                    _ = v48
                }
                var v49 int16 = (v3 & v3)
                for fi50 := 0; fi50 < 16; fi50++ {
                    v2 += (v2 ^ genFn0("\n\t"))
                    var v51 uint = 290
                    if true {
                        fmt.Println(v5)
                        if (1.5 > 3.14159) {
                            if false {
                                for fi52 := 0; fi52 < 14; fi52++ {
                                    if true {
                                        var v53 bool = false
                                        switch int(v47) {
                                        case -1:
                                            if true {
                                                var v54 int32 = -1
                                                v2 += genFn0(((("a" + "a") + "a") + ("" + "\n\t")))
                                                _ = v54
                                            } else {
                                                v55 := GenS4{F0: 42, F1: ((^(v2)) != (^(fi52)))}
                                                fmt.Println(v5, v3, v4)
                                                if v53 {
                                                    v56 := GenS6{F0: 0}
                                                    fmt.Println(v47)
                                                    if (0.1 == 1e10) {
                                                        v53 = true
                                                        v6[("" + "")] = 65535
                                                        if (321 != (v51 - v51)) {
                                                            v57 := GenS7{F0: 860, F1: uint64(v51), F2: 893}
                                                            v4 += (v4 + v5)
                                                            if (v53 || false) {
                                                                v4 -= (v4 ^ 127)
                                                            }
                                                            _ = v57
                                                        } else {
                                                            if ("foo bar" != "foo bar") {
                                                                var v58 string = "hello"
                                                                _ = v58
                                                            } else {
                                                                v59 := GenS8{F0: 617}
                                                                v60 := GenS7{F0: 4294967295, F1: 1, F2: 297}
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                            var v61 uint16 = 2
                                                            _ = v61
                                                        }
                                                    }
                                                    _ = v56
                                                }
                                                _ = v55
                                            }
                                        default:
                                            fmt.Println(v4, fi52, v5)
                                        }
                                        var v62 int16 = (^(v3))
                                        _ = v53
                                        _ = v62
                                    } else {
                                        v63 := GenS9{F0: -1}
                                        v6["a"] = 0
                                        v1 = append(v1, genFn0("\n\t"))
                                        _ = v63
                                    }
                                    _ = fi52
                                }
                                v5 = v4
                            }
                            if (v4 > 12) {
                                fmt.Println(v2, v2, v4)
                                fmt.Println(v51, v5, v2)
                                if false {
                                    v64 := GenS4{F0: 1, F1: (255 > 2)}
                                    _ = v64
                                } else {
                                    if (true || ((true || (false && ((v47 << (((v51 * 4294967295)) & 63)) <= 2))) || (false || (3.14159 != 1.5)))) {
                                        var v65 float64 = 1.5
                                        _ = v65
                                    } else {
                                        v1 = append(v1, genFn0(("\n\t" + "")))
                                        fmt.Println(v51, v4)
                                        if len(v1) > 0 {
                                            v1[0] = (fi50 * -221)
                                        }
                                    }
                                    var v66 string = ("hello" + ("\n\t" + ("\n\t" + ("a" + "foo bar"))))
                                    var v67 float64 = 3.14159
                                    _ = v66
                                    _ = v67
                                }
                            } else {
                                v49 = v49
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch "foo bar" {
                                case "a":
                                    for fi68 := 0; fi68 < 10; fi68++ {
                                        if ((1e10 == 3.14159) && true) {
                                            var v69 int = genFn0(("hello" + "hello"))
                                            v70 := GenS10{F0: genFn1(int16(v2), true), F1: -1, F2: fi50}
                                            if false {
                                                _ = len(v6)
                                                switch 128 {
                                                case -608:
                                                    v47 = (v47 << (((v51 / ((v51) | 1))) & 63))
                                                case 32768:
                                                    var v71 uint = (v51 + 255)
                                                    _ = v71
                                                case 463:
                                                    var v72 int8 = int8(v47)
                                                    _ = v72
                                                default:
                                                    v6["foo bar"] = 42
                                                }
                                                v5 = v4
                                                fmt.Println(v47, v4, v49)
                                                if ((v5 & (v5 + v5)) <= v5) {
                                                    v73 := GenS4{F0: 1, F1: (((false && false) && true) || true)}
                                                    fmt.Println(v47)
                                                    v74 := GenS11{F0: 0}
                                                    v75 := GenS5{F0: -815, F1: ("foo bar" + "foo bar")}
                                                    fmt.Println(v4, v5, v49)
                                                    v3 += (v3 | v49)
                                                    switch (^(fi50)) {
                                                    case -926:
                                                        v47 -= v47
                                                    default:
                                                        var v76 int64 = 9223372036854775807
                                                        _ = v76
                                                    }
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                }
                                                v77 := GenS10{F0: 18446744073709551615, F1: 0, F2: fi50}
                                                if (2147483647 != 906) {
                                                    v2 += v2
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (false && (42 == 398)) {
                                                    var v78 int32 = -2147483648
                                                    _ = v78
                                                } else {
                                                    if true {
                                                        v4 -= (-(v4))
                                                        switch fi68 {
                                                        case 986:
                                                            var v79 string = "a"
                                                            _ = v79
                                                        case -2147483648:
                                                            v49 = (^(v49))
                                                            v69 += genFn0("a")
                                                        case 9223372036854775807:
                                                            var v80 int8 = v5
                                                            _ = v80
                                                        default:
                                                            if (1 >= 18446744073709551615) {
                                                                var v81 string = "a"
                                                                v47 = v47
                                                                fmt.Println(v4)
                                                                _ = v81
                                                            }
                                                        }
                                                    } else {
                                                        v82 := v6["a"]
                                                        var v83 uint16 = (v47 * v47)
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                    if false {
                                                        _ = len(v6)
                                                    } else {
                                                        if false {
                                                            _ = len(v6)
                                                            v84 := GenS5{F0: -1, F1: ""}
                                                            var v85 uint = 2
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                        fmt.Println(fi50, v69, v47)
                                                        var v86 int64 = -9223372036854775808
                                                        _ = v86
                                                    }
                                                }
                                                _ = v77
                                            }
                                            _ = v69
                                            _ = v70
                                        } else {
                                            var v87 int8 = v4
                                            if false {
                                                v88 := GenS12{F0: -2147483648, F1: (v47 & v47), F2: v51}
                                                _ = v88
                                            } else {
                                                v3 -= v3
                                            }
                                            _ = v87
                                        }
                                        _ = fi68
                                    }
                                case "":
                                    switch fi50 {
                                    case 875:
                                        var v89 uint = 65535
                                        _ = v89
                                    case -9223372036854775808:
                                        var v90 int16 = v3
                                        _ = v90
                                    case 128:
                                        v47 = 2
                                    default:
                                        if true {
                                            v51 += (v51 / ((255) | 1))
                                            var v91 int32 = 1
                                            _ = v91
                                        } else {
                                            v92 := GenS13{F0: v51, F1: v51, F2: 963}
                                            _ = v92
                                        }
                                    }
                                default:
                                    v93 := GenS13{F0: (^(v51)), F1: v51, F2: 0}
                                    _ = v93
                                }
                            }
                            for fi94 := 0; fi94 < 18; fi94++ {
                                v4 += (^(v4))
                                _ = fi94
                            }
                        }
                        v95 := GenS14{F0: (v3 / ((-223) | 1))}
                        _ = v95
                    }
                    _ = fi50
                    _ = v51
                }
                _ = v49
            }
            var v96 uint8 = 130
            _ = v96
        }
    case 32768:
        for fi97 := 0; fi97 < 13; fi97++ {
            v98 := GenS9{F0: v3}
            _ = fi97
            _ = v98
        }
    default:
        var v99 int64 = 1
        _ = v99
    }
    if (!false) {
        var v100 uint = 4
        for fi101 := 0; fi101 < 4; fi101++ {
            v102 := GenS6{F0: 973}
            var v103 float32 = 3.14159
            _ = fi101
            _ = v102
            _ = v103
        }
        _ = v100
    } else {
        switch genFn0(("hello" + "hello")) {
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 1:
            if true {
                v4 = (v5 >> ((307) & 63))
            } else {
                v5 -= v5
                var v104 int = (v2 * (-(v2)))
                v1 = append(v1, -2147483648)
                _ = v104
            }
        default:
            fmt.Println(v3)
        }
        for fi105 := 0; fi105 < 19; fi105++ {
            var v106 int16 = v3
            fmt.Println(v3, v4)
            var v107 string = ("a" + "hello")
            _ = fi105
            _ = v106
            _ = v107
        }
    }
    for fi108 := 0; fi108 < 12; fi108++ {
        for fi109 := 0; fi109 < 10; fi109++ {
            v110 := GenS7{F0: 2147483648, F1: 1, F2: 0}
            _ = fi109
            _ = v110
        }
        for fi111 := 0; fi111 < 4; fi111++ {
            var v112 int16 = v3
            fmt.Println(v4, v47)
            _ = fi111
            _ = v112
        }
        if true {
            v47 += (v47 + v47)
        }
        _ = fi108
    }
    for fi113 := 0; fi113 < 16; fi113++ {
        v4 = (^(v5))
        if len(v1) > 0 {
            v1[0] = genFn0(("foo bar" + ("foo bar" + "foo bar")))
        }
        _ = fi113
    }
    v47 = v47
    var v114 uint8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
