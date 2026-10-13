package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64; F1 string; F2 int }
type GenS4 struct { F0 string; F1 uint64; F2 int32 }
type GenS5 struct { F0 string; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS7 struct { F0 uint }
func genFn0(p0 bool) bool {
    var v6 int16 = 1000
    if p0 {
        v7 := GenS1{F0: p0}
        var v8 int8 = 42
        var v9 int8 = (v8 + -26)
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int32 = 2147483647
    _ = p0
    _ = v10
    _ = v6
    return p0
}
func genFn1() uint64 {
    var v11 uint32 = 2147483647
    v12 := GenS2{F0: 469, F1: 4294967295, F2: v11}
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (3.14159 / ((3.14159 + (1e10 + 0.1)) / float64(1.5)))
    v3 := []string{"a"}
    v4 := make(map[string]int)
    var v5 uint8 = 128
    v2 = v2
    switch "\n\t" {
    case "a":
        if (-1 <= 1) {
            var v13 string = "foo bar"
            _ = v13
        } else {
            if true {
                var v14 uint32 = 65535
                if true {
                    if (1 > -1) {
                        var v15 int = 2147483647
                        var v16 int16 = 1
                        fmt.Println(v15)
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS1{F0: (!("hello" > "foo bar"))}
                        var v18 uint = 357
                        _ = v17
                        _ = v18
                    }
                    var v19 float32 = -2.25
                    _ = v19
                }
                if (42 < -1) {
                    if len(v3) > 0 {
                        v3[0] = ((("foo bar" + "hello") + "foo bar") + "a")
                    }
                    switch "" {
                    case "\n\t":
                        var v20 string = "\n\t"
                        _ = v20
                    case "a":
                        v5 = v5
                    }
                }
                _ = v14
            }
            var v21 uint32 = 4294967295
            _ = v21
        }
    case "":
        if true {
            var v22 int16 = 1000
            v23 := GenS2{F0: 32767, F1: -1, F2: 2147483648}
            _ = v22
            _ = v23
        }
    case "hello":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    default:
        var v24 uint32 = 2147483648
        _ = v24
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v5, v5)
    default:
        v25 := GenS1{F0: (628 > 127)}
        _ = v25
    }
    v2 = 1e10
    for fi26 := 0; fi26 < 10; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            v2 += 3.14159
            v5 -= v5
            _ = fi27
        }
        _ = fi26
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int64 = -96
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := GenS1{F0: (false || (!(49 != 127)))}
            v32 := GenS3{F0: 122, F1: "", F2: (fi30 << ((71) & 63))}
            _ = fi30
            _ = v31
            _ = v32
        }
        v1[375] = ("foo bar" + ("" + "\n\t"))
        _ = fi28
        _ = v29
    }
    var v33 uint8 = 70
    for fi34 := 0; fi34 < 19; fi34++ {
        if true {
            switch "hello" {
            case "foo bar":
                var v35 float32 = (1e10 * (1e10 + -2.25))
                _ = v35
            case "a":
                switch fi34 {
                case 127:
                    if false {
                        _ = len(v4)
                    } else {
                        switch "a" {
                        case "\n\t":
                            v36 := GenS2{F0: 882, F1: (-(fi34)), F2: 2}
                            _ = v36
                        default:
                            if true {
                                var v37 uint16 = 65535
                                _ = v37
                            }
                        }
                        switch -1 {
                        case 157:
                            var v38 int64 = 503
                            _ = v38
                        default:
                            var v39 float64 = v2
                            _ = v39
                        }
                        switch (fi34 & fi34) {
                        case -342:
                            v5 = v33
                        default:
                            v5 -= v5
                        }
                    }
                case 128:
                    for fi40 := 0; fi40 < 9; fi40++ {
                        v41 := v4["hello"]
                        switch "" {
                        case "":
                            v3 = append(v3, "\n\t")
                        case "\n\t":
                            var v42 string = "hello"
                            _ = v42
                        default:
                            v43 := GenS1{F0: (v5 >= v5)}
                            _ = v43
                        }
                        _ = fi40
                        _ = v41
                    }
                }
            }
            if (!false) {
                var v44 int = fi34
                _ = v44
            } else {
                v3 = append(v3, ("foo bar" + ""))
                for fi45 := 0; fi45 < 2; fi45++ {
                    var v46 uint = 25
                    var v47 int = 894
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            if (!true) {
                if (true || true) {
                    for fi48 := 0; fi48 < 10; fi48++ {
                        v2 = 0.0
                        v2 = v2
                        _ = fi48
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    fmt.Println(v5)
                }
                if false {
                    for fi49 := 0; fi49 < 17; fi49++ {
                        var v50 uint = 1
                        switch (("foo bar" + ("\n\t" + "\n\t")) + ("foo bar" + "hello")) {
                        case "a":
                            var v51 float32 = 0.0
                            _ = v51
                        }
                        _ = fi49
                        _ = v50
                    }
                    var v52 uint32 = 42
                    _ = v52
                } else {
                    switch "foo bar" {
                    case "a":
                        if (128 >= 762) {
                            switch (fi34 * (fi34 & -9223372036854775808)) {
                            case 42:
                                fmt.Println(fi34, fi34)
                            case -646:
                                if false {
                                    var v53 string = ("a" + "\n\t")
                                    var v54 string = (v53 + ("" + v53))
                                    v33 = (^(v5))
                                    _ = v53
                                    _ = v54
                                }
                            default:
                                var v55 uint = 0
                                _ = v55
                            }
                            for fi56 := 0; fi56 < 3; fi56++ {
                                var v57 int = (fi34 % ((32768) | 1))
                                if (32767 == fi34) {
                                    fmt.Println(fi34, fi34)
                                    v58 := GenS1{F0: (!true)}
                                    v57 = (v57 % (((-(v57))) | 1))
                                    _ = v58
                                } else {
                                    v59 := GenS1{F0: false}
                                    switch fi56 {
                                    case 4294967295:
                                        v33 -= (v5 ^ (v5 ^ v5))
                                    case 1:
                                        fmt.Println(fi56, v33, fi56)
                                    case 127:
                                        var v60 int8 = -128
                                        _ = v60
                                    }
                                    if true {
                                        fmt.Println(v2, fi56, fi34)
                                        var v61 int16 = 942
                                        fmt.Println(fi56, v2)
                                        _ = v61
                                    }
                                    _ = v59
                                }
                                _ = fi56
                                _ = v57
                            }
                        } else {
                            var v62 float64 = v2
                            switch "foo bar" {
                            case "":
                                var v63 uint = 581
                                _ = v63
                            }
                            _ = v62
                        }
                    }
                    var v64 int8 = 0
                    for fi65 := 0; fi65 < 2; fi65++ {
                        var v66 uint16 = 42
                        var v67 int = (-(fi65))
                        v3 = append(v3, "foo bar")
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    _ = v64
                }
                for fi68 := 0; fi68 < 13; fi68++ {
                    v5 = (v33 % ((0) | 1))
                    v33 -= v5
                    _ = fi68
                }
            } else {
                var v69 int32 = -938
                for fi70 := 0; fi70 < 20; fi70++ {
                    var v71 int32 = (v69 | -198)
                    v72 := GenS3{F0: 4294967295, F1: "a", F2: (fi34 >> ((1) & 63))}
                    fmt.Println(v2, v69, fi70)
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                var v73 int32 = v69
                _ = v69
                _ = v73
            }
        } else {
            _ = len(v1)
            var v74 int16 = 32767
            v75 := v1[629]
            _ = v74
            _ = v75
        }
        switch (fi34 + fi34) {
        case -9223372036854775808:
            var v76 uint32 = 65535
            _ = v76
        case 219:
            fmt.Println(fi34, fi34)
        case -2147483648:
            v77 := GenS3{F0: 4294967295, F1: "\n\t", F2: (fi34 ^ 4294967295)}
            _ = v77
        default:
            v2 += v2
        }
        for fi78 := 0; fi78 < 15; fi78++ {
            v5 = v33
            var v79 uint = 4294967295
            if (v79 <= v79) {
                _ = len(v1)
                v79 -= 42
            } else {
                if true {
                    v80 := GenS2{F0: 2, F1: fi78, F2: 2147483648}
                    v79 = v79
                    var v81 float64 = v2
                    _ = v80
                    _ = v81
                } else {
                    _ = len(v4)
                    v2 += -2.25
                    v79 -= v79
                }
                switch (("a" + ("\n\t" + "foo bar")) + "\n\t") {
                case "a":
                    if (!false) {
                        if true {
                            switch fi78 {
                            case -9223372036854775808:
                                var v82 int32 = -193
                                _ = v82
                            case -623:
                                v5 -= (^(v33))
                                delete(v4, ("hello" + "\n\t"))
                            }
                        } else {
                            v83 := GenS4{F0: ("foo bar" + (("a" + ("foo bar" + "foo bar")) + "")), F1: 0, F2: 2147483647}
                            fmt.Println(v79)
                            var v84 int64 = -1
                            _ = v83
                            _ = v84
                        }
                        v85 := v4["\n\t"]
                        fmt.Println(fi78)
                        _ = v85
                    } else {
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v86 uint = (^(v79))
                            if ("\n\t" == (("foo bar" + "hello") + "\n\t")) {
                                if (false && (0 < 0)) {
                                    v87 := GenS4{F0: (("hello" + "\n\t") + "a"), F1: uint64(v86), F2: 1}
                                    var v88 float64 = (3.14159 * v2)
                                    v86 += 255
                                    _ = v87
                                    _ = v88
                                } else {
                                    v89 := GenS5{F0: "a", F1: 2, F2: 404}
                                    _ = v89
                                }
                                v90 := v4["\n\t"]
                                _ = v90
                            }
                            _ = v86
                        }
                        v91 := GenS2{F0: 42, F1: (fi34 >> ((v79) & 63)), F2: 2147483648}
                        var v92 uint8 = v33
                        _ = v91
                        _ = v92
                    }
                case "":
                    var v93 string = ""
                    _ = v93
                case "foo bar":
                    fmt.Println(fi34, fi78)
                }
                if (!false) {
                    var v94 uint = (v79 >> (((v79 / (((v79 + (v79 - v79))) | 1))) & 63))
                    fmt.Println(fi34, v5)
                    if false {
                        v95 := GenS6{F0: 255, F1: 42, F2: 65535}
                        fmt.Println(v2, v5)
                        _ = v95
                    }
                    _ = v94
                } else {
                    _ = len(v3)
                    v3 = append(v3, "a")
                }
            }
            _ = fi78
            _ = v79
        }
        _ = fi34
    }
    v96 := GenS3{F0: 18446744073709551615, F1: ("foo bar" + ("" + "foo bar")), F2: 184}
    v3 = append(v3, (("foo bar" + "hello") + "a"))
    if (-215 < 9223372036854775807) {
        if (9223372036854775807 != -9223372036854775808) {
            var v97 int = 9223372036854775807
            if false {
                var v98 string = ("foo bar" + "")
                v99 := GenS5{F0: "foo bar", F1: 18446744073709551615, F2: 42}
                _ = v98
                _ = v99
            }
            _ = v97
        }
    }
    if true {
        v100 := GenS7{F0: 42}
        v101 := GenS3{F0: 18446744073709551615, F1: "", F2: 1}
        _ = v100
        _ = v101
    } else {
        var v102 uint64 = 42
        v33 = v5
        _ = v102
    }
    var v103 uint16 = 881
    for fi104 := 0; fi104 < 17; fi104++ {
        if (128 > fi104) {
            switch fi104 {
            case 1:
                var v105 int16 = 1000
                _ = v105
            default:
                v106 := GenS5{F0: "hello", F1: 0, F2: -489}
                _ = v106
            }
            var v107 int16 = 97
            v107 += (-(v107))
            _ = v107
        } else {
            _ = len(v1)
            for fi108 := 0; fi108 < 18; fi108++ {
                v33 += 2
                v4[""] = fi108
                var v109 int = (fi104 << ((uint(v5)) & 63))
                delete(v4, "a")
                v110 := GenS1{F0: true}
                _ = fi108
                _ = v109
                _ = v110
            }
        }
        switch "hello" {
        case "hello":
            if len(v3) > 0 {
                v3[0] = "a"
            }
        }
        fmt.Println(v33, fi104, v5)
        _ = fi104
    }
    for fi111 := 0; fi111 < 9; fi111++ {
        v2 = (float64(1e10) * 3.14159)
        _ = fi111
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v96
}
