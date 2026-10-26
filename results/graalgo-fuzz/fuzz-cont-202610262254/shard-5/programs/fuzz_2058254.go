package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint; F1 uint16; F2 int }
func genFn0() int8 {
    var v5 uint32 = 719
    var v6 string = (("\n\t" + "\n\t") + "a")
    v5 = (v5 * (v5 & v5))
    _ = v5
    _ = v6
    return 3
}
func genFn1(p0 uint64) uint8 {
    var v7 uint32 = 2147483648
    v8 := []int{}
    switch -9223372036854775808 {
    case 4294967295:
        var v9 int16 = 42
        _ = v9
    case -303:
        var v10 int8 = 127
        _ = v10
    case 2147483647:
        p0 = p0
    }
    _ = p0
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 int32 = -2147483648
    v4 := make(map[int]int)
    fmt.Println(v3)
    v11 := []int{128, 9223372036854775807, 1}
    v3 = -1
    switch ("\n\t" + "foo bar") {
    case "":
        v3 = -2147483648
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        switch ((((("hello" + "\n\t") + "\n\t") + "hello") + "a") + ("" + "a")) {
        case "a":
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int8 = 42
                if (true && false) {
                    v15 := GenS1{F0: -1, F1: (-(fi12))}
                    fmt.Println(fi12, fi13, v14)
                    fmt.Println(fi12, v14, fi13)
                    _ = v15
                }
                switch "" {
                case "":
                    if (false || false) {
                        fmt.Println(fi13, v3)
                    } else {
                        var v16 uint64 = 106
                        _ = len(v1)
                        if (!((v14 ^ v14) > v14)) {
                            v14 = -128
                        } else {
                            v4[fi12] = (fi13 - (fi13 % (((fi13 | 32768)) | 1)))
                            var v17 float32 = 3.14159
                            _ = v17
                        }
                        _ = v16
                    }
                case "\n\t":
                    if (!false) {
                        if true {
                            if (((0.1 + 3.14159) + ((1.5 * 0.0) - float32(-2.25))) == 0.1) {
                                var v18 int64 = 492
                                v18 = v18
                                _ = v18
                            }
                        }
                    } else {
                        var v19 string = "\n\t"
                        _ = v19
                    }
                case "hello":
                    fmt.Println(v14, v3, v14)
                }
                _ = fi13
                _ = v14
            }
        case "hello":
            if (true && (2 != 2)) {
                fmt.Println(v3, v3)
                fmt.Println(fi12, fi12)
                var v20 uint = 42
                _ = v20
            } else {
                switch "a" {
                case "a":
                    v21 := GenS1{F0: 9223372036854775807, F1: (fi12 + 2147483647)}
                    _ = v21
                default:
                    var v22 string = "a"
                    _ = v22
                }
            }
        case "foo bar":
            var v23 int64 = 9223372036854775807
            _ = v23
        }
        v24 := GenS2{F0: ("foo bar" + "\n\t"), F1: 2}
        if true {
            v3 += -2147483648
            var v25 bool = false
            if (1 != 800) {
                var v26 uint64 = 42
                _ = v26
            } else {
                for fi27 := 0; fi27 < 1; fi27++ {
                    switch fi12 {
                    case 596:
                        if len(v11) > 0 {
                            v11[0] = fi12
                        }
                    }
                    _ = fi27
                }
                v28 := GenS2{F0: "", F1: 2}
                if ((-2.25 + ((3.14159 - (float32(0.1) * 0.0)) - 3.14159)) >= 1e10) {
                    if v25 {
                        v25 = ((2147483647 > 4294967295) || false)
                        var v29 string = (("" + ("hello" + ((("\n\t" + "a") + "foo bar") + "\n\t"))) + "")
                        v3 = (v3 / ((v3) | 1))
                        _ = v29
                    } else {
                        v30 := GenS2{F0: "", F1: 42}
                        if (v3 <= v3) {
                            var v31 bool = (3.14159 != 0.1)
                            v31 = (float64(0.1) > 3.14159)
                            if false {
                                v32 := v4[(^(fi12))]
                                var v33 int64 = -1
                                _ = v32
                                _ = v33
                            }
                            _ = v31
                        } else {
                            for fi34 := 0; fi34 < 1; fi34++ {
                                fmt.Println(fi12)
                                _ = fi34
                            }
                            if (!(!(127 != genFn1(753)))) {
                                var v35 float32 = 0.1
                                var v36 int64 = 1
                                _ = v35
                                _ = v36
                            } else {
                                v37 := GenS1{F0: 9223372036854775807, F1: (fi12 & fi12)}
                                v38 := GenS2{F0: ("a" + ("a" + ("foo bar" + "hello"))), F1: 0}
                                _ = v37
                                _ = v38
                            }
                        }
                        _ = v30
                    }
                    fmt.Println(v25)
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v39 uint = 1
                        _ = v39
                    case "hello":
                        v40 := GenS2{F0: ("\n\t" + "hello"), F1: 65535}
                        _ = v40
                    case "foo bar":
                        var v41 int = fi12
                        _ = v41
                    }
                    fmt.Println(v3, v25)
                }
                _ = v28
            }
            _ = v25
        } else {
            var v42 float64 = (-2.25 / 1.0)
            if len(v11) > 0 {
                v11[0] = (fi12 % (((-(fi12))) | 1))
            }
            v42 = ((0.0 / 1.5) - 0.1)
            _ = v42
        }
        _ = fi12
        _ = v24
    }
    if (!(("foo bar" + "") > "\n\t")) {
        _ = len(v11)
        if ((((-2.25 * 0.0) + 0.1) / 1.0) < (((1.5 + 1e10) + (3.14159 - 3.14159)) / float64(1e10))) {
            v43 := v1[-108]
            v44 := GenS1{F0: -20, F1: 11}
            delete(v1, 1)
            _ = v43
            _ = v44
        } else {
            delete(v4, -1)
            var v45 bool = (127 != genFn0())
            _ = v45
        }
    }
    var v46 uint64 = 18446744073709551615
    if (-32768 == 32767) {
        fmt.Println(v3, v3)
        switch 800 {
        case -171:
            v3 -= v3
        case 1:
            fmt.Println(v3, v3)
        case 370:
            if false {
                switch "" {
                case "\n\t":
                    fmt.Println(v46)
                case "foo bar":
                    for fi47 := 0; fi47 < 2; fi47++ {
                        fmt.Println(v46)
                        var v48 int8 = 42
                        _ = fi47
                        _ = v48
                    }
                case "":
                    if false {
                        if len(v11) > 0 {
                            v11[0] = -672
                        }
                        v49 := GenS2{F0: ("hello" + "a"), F1: 42}
                        fmt.Println(v3, v46)
                        _ = v49
                    } else {
                        for fi50 := 0; fi50 < 19; fi50++ {
                            for fi51 := 0; fi51 < 12; fi51++ {
                                fmt.Println(v46)
                                var v52 bool = true
                                v3 += 0
                                _ = fi51
                                _ = v52
                            }
                            if true {
                                v46 = v46
                            } else {
                                for fi53 := 0; fi53 < 19; fi53++ {
                                    if (int64(v46) != 288) {
                                        var v54 uint = 255
                                        _ = v54
                                    } else {
                                        v55 := v4[2147483647]
                                        if (!(!(!true))) {
                                            v46 = 0
                                            v46 = 4294967295
                                            delete(v2, 1)
                                        }
                                        _ = v55
                                    }
                                    v3 = v3
                                    v46 = (v46 | 1)
                                    _ = fi53
                                }
                                var v56 uint8 = 42
                                _ = v56
                            }
                            var v57 int = (-(fi50))
                            _ = fi50
                            _ = v57
                        }
                        switch 1 {
                        case 42:
                            fmt.Println(v46, v3)
                        }
                        v58 := GenS2{F0: "", F1: 32767}
                        _ = v58
                    }
                default:
                    var v59 int8 = 127
                    _ = v59
                }
                var v60 float32 = (0.1 - ((1.5 + 3.14159) + 3.14159))
                var v61 uint64 = (^(v46))
                _ = v60
                _ = v61
            } else {
                v46 += 438
                if false {
                    switch 0 {
                    case -544:
                        v62 := GenS1{F0: 79, F1: 0}
                        _ = v62
                    case -360:
                        var v63 bool = false
                        _ = v63
                    default:
                        fmt.Println(v46, v46)
                    }
                } else {
                    fmt.Println(v46, v46)
                }
            }
        default:
            var v64 float64 = 0.0
            _ = v64
        }
    }
    var v65 uint32 = 924
    var v66 int = -1
    for fi67 := 0; fi67 < 4; fi67++ {
        for fi68 := 0; fi68 < 15; fi68++ {
            v69 := GenS3{F0: ((^(v65)) >= 2147483647), F1: (("hello" + ("foo bar" + "hello")) + "hello")}
            var v70 int16 = 1000
            if true {
                if ((v70 % (((-(v70))) | 1)) == (v70 - -32768)) {
                    if (0.0 < (0.0 / (0.1 * 3.14159))) {
                        var v71 float32 = ((3.14159 + 0.1) - (float32(3.14159) / float32(((((float64((0.1 / 1e10)) + -2.25) / -2.25) / 3.14159) + 3.14159))))
                        if false {
                            if (65535 >= 0) {
                                switch (v66 * (fi67 | (v66 / ((v66) | 1)))) {
                                case 127:
                                    fmt.Println(v65, v70, fi67)
                                case -847:
                                    if len(v11) > 0 {
                                        v11[0] = 128
                                    }
                                case -9223372036854775808:
                                    var v72 float64 = ((0.0 / 1.0) / float64((0.0 - -2.25)))
                                    _ = v72
                                }
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                            }
                            var v73 int64 = 0
                            switch ("\n\t" + "a") {
                            case "foo bar":
                                var v74 uint8 = 42
                                _ = v74
                            case "hello":
                                fmt.Println(v46)
                            case "a":
                                v75 := GenS1{F0: (v73 * -9223372036854775808), F1: int(v46)}
                                _ = v75
                            }
                            _ = v73
                        } else {
                            v76 := v2[v65]
                            v70 += (v70 / (((v70 | (v70 % ((v70) | 1)))) | 1))
                            _ = v76
                        }
                        switch "a" {
                        case "":
                            var v77 int32 = (-(v3))
                            _ = v77
                        case "foo bar":
                            if true {
                                v78 := GenS2{F0: "", F1: 2}
                                _ = v78
                            }
                        case "a":
                            v79 := GenS3{F0: false, F1: "\n\t"}
                            _ = v79
                        default:
                            v65 = (v65 / ((v65) | 1))
                        }
                        v65 = v65
                        _ = v71
                    } else {
                        switch "foo bar" {
                        case "hello":
                            var v80 string = ((("a" + "foo bar") + ("hello" + "")) + "hello")
                            _ = v80
                        default:
                            switch 9223372036854775807 {
                            case 4294967295:
                                var v81 int = (-(v66))
                                _ = v81
                            default:
                                switch fi68 {
                                case 9223372036854775807:
                                    v82 := GenS2{F0: "foo bar", F1: 2}
                                    _ = v82
                                case 32767:
                                    var v83 int64 = 42
                                    _ = v83
                                case -597:
                                    var v84 float64 = (float64(-2.25) / -2.25)
                                    _ = v84
                                default:
                                    _ = len(v11)
                                }
                            }
                        }
                    }
                    v70 = 0
                    if (34 == 0) {
                        v85 := GenS3{F0: (-9223372036854775808 > 684), F1: ("" + ("a" + ("foo bar" + "\n\t")))}
                        v86 := GenS2{F0: "foo bar", F1: 65535}
                        v87 := GenS4{F0: v65, F1: 1}
                        v70 = v70
                        _ = v85
                        _ = v86
                        _ = v87
                    } else {
                        var v88 int16 = (-(v70))
                        fmt.Println(v70, v46, fi67)
                        _ = v88
                    }
                    v1[fi68] = ((false || true) && (!false))
                } else {
                    var v89 int = (fi68 | v66)
                    _ = v89
                }
                fmt.Println(fi68)
                if ((!(-1 != v3)) && (!(fi67 <= 0))) {
                    v90 := GenS4{F0: 2147483648, F1: (-(v70))}
                    if (32767 == v70) {
                        delete(v1, (v66 ^ (^(fi67))))
                    }
                    v46 -= v46
                    _ = v90
                } else {
                    var v91 int64 = 535
                    v92 := GenS4{F0: uint32(v46), F1: (v70 % ((v70) | 1))}
                    _ = v91
                    _ = v92
                }
            }
            if false {
                var v93 int8 = 1
                var v94 string = ("" + ("hello" + "a"))
                _ = v93
                _ = v94
            }
            _ = fi68
            _ = v69
            _ = v70
        }
        v95 := GenS5{F0: 924, F1: uint16(v65), F2: v66}
        v96 := GenS4{F0: v65, F1: 0}
        _ = fi67
        _ = v95
        _ = v96
    }
    v3 -= 1
    var v97 bool = false
    var v98 float32 = 1e10
    if len(v11) > 0 {
        v11[0] = 32768
    }
    switch (v66 % (((v66 >> ((65535) & 63))) | 1)) {
    case 9223372036854775807:
        var v99 int = (v66 * -654)
        _ = v99
    default:
        delete(v1, -1)
    }
    if len(v11) > 0 {
        v11[0] = (v66 << ((135) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v65
    _ = v66
    _ = v97
    _ = v98
}
