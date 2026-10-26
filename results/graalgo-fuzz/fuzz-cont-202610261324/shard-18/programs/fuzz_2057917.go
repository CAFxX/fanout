package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int8; F1 int; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 uint64; F2 bool }
type GenS6 struct { F0 uint; F1 int8; F2 uint16 }
type GenS7 struct { F0 uint8; F1 string; F2 uint8 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint32; F1 int; F2 int }
type GenS11 struct { F0 string }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int8; F1 int8 }
func genFn0(p0 int, p1 uint32, p2 int8) int64 {
    v4 := []int{p0, p0}
    v4 = append(v4, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}

func main() {
    var v1 bool = (173 >= 255)
    var v2 int = -2147483648
    var v3 uint32 = 544
    v1 = (!(!true))
    v3 = (v3 % (((v3 - (v3 - 354))) | 1))
    v5 := []int{v2, v2, (v2 | v2), v2}
    v6 := GenS1{F0: 127, F1: 1, F2: 4294967295}
    var v7 bool = ("a" <= ("a" + ("\n\t" + ("a" + "a"))))
    for fi8 := 0; fi8 < 9; fi8++ {
        if ((true || v1) || true) {
            v9 := GenS2{F0: -945, F1: 711}
            v7 = (-9223372036854775808 < -1)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v9
        } else {
            v2 += v2
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v10 int16 = 32767
        if (genFn0(9223372036854775807, 687, 42) <= 1) {
            var v11 int16 = v10
            _ = v11
        }
        _ = fi8
        _ = v10
    }
    if true {
        var v12 uint32 = (^(v3))
        if v7 {
            for fi13 := 0; fi13 < 12; fi13++ {
                switch fi13 {
                case -455:
                    if v1 {
                        v3 = (v3 % ((v3) | 1))
                        for fi14 := 0; fi14 < 17; fi14++ {
                            v7 = (!(-1 == 1000))
                            v3 += (v12 ^ v12)
                            var v15 string = (("foo bar" + "foo bar") + ((("hello" + (("a" + "a") + "foo bar")) + ("foo bar" + "")) + "hello"))
                            _ = fi14
                            _ = v15
                        }
                    }
                case 634:
                    var v16 int16 = 0
                    _ = v16
                }
                if true {
                    if v7 {
                        v2 = 4294967295
                        switch ("hello" + "hello") {
                        case "":
                            v5 = append(v5, (fi13 | 414))
                        default:
                            var v17 string = ("\n\t" + ("\n\t" + ("hello" + "\n\t")))
                            _ = v17
                        }
                        var v18 int32 = 534
                        _ = v18
                    }
                }
                v19 := GenS3{F0: 42, F1: fi13, F2: 2147483647}
                _ = fi13
                _ = v19
            }
            switch v2 {
            case 857:
                v20 := GenS4{F0: -1}
                _ = v20
            }
            var v21 uint16 = 95
            _ = v21
        }
        _ = v12
    } else {
        v5 = append(v5, v2)
        if ((!(-2147483648 == 2147483647)) && false) {
            v5 = append(v5, v2)
            var v22 int = v2
            _ = v22
        } else {
            var v23 int = (^(v2))
            if true {
                if (42 == 65535) {
                    var v24 int64 = 42
                    _ = v24
                }
                var v25 float32 = (-2.25 + 0.0)
                fmt.Println(v1, v23)
                _ = v25
            } else {
                for fi26 := 0; fi26 < 6; fi26++ {
                    fmt.Println(fi26, v23)
                    _ = fi26
                }
                v27 := GenS1{F0: 128, F1: 42, F2: 0}
                _ = v27
            }
            switch "a" {
            case "":
                for fi28 := 0; fi28 < 19; fi28++ {
                    var v29 int = v2
                    _ = fi28
                    _ = v29
                }
            case "hello":
                var v30 int = (v23 >> ((65535) & 63))
                _ = v30
            case "\n\t":
                v31 := GenS2{F0: 1, F1: 0}
                _ = v31
            }
            _ = v23
        }
        v7 = (0.1 < 0.1)
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 int = 836
            switch "\n\t" {
            case "hello":
                if false {
                    v2 += fi33
                    if (false || (0.1 == (float64(((1.5 - ((1.5 * float32(1e10)) + float32(3.14159))) / -2.25)) + float64(0.1)))) {
                        v3 = v3
                    } else {
                        switch "hello" {
                        case "foo bar":
                            if (v7 || v7) {
                                if ((v7 || v1) || (65535 <= 660)) {
                                    v35 := GenS5{F0: 69, F1: 42, F2: (true || v7)}
                                    v36 := GenS3{F0: 1, F1: (fi33 + (v2 ^ (fi33 | (^(fi32))))), F2: 824}
                                    _ = v35
                                    _ = v36
                                }
                                if v1 {
                                    v2 = (fi33 << ((992) & 63))
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = fi32
                                    }
                                }
                            } else {
                                var v37 int32 = 1
                                fmt.Println(v7)
                                var v38 uint8 = 127
                                v39 := GenS6{F0: 213, F1: 0, F2: 545}
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                        case "\n\t":
                            var v40 uint64 = 619
                            _ = v40
                        default:
                            v3 = 240
                        }
                        if (0.0 == float64(0.1)) {
                            if v7 {
                                if v1 {
                                    var v41 uint64 = 1
                                    v2 = (fi32 % (((v2 / (((-(fi33))) | 1))) | 1))
                                    switch (-(v34)) {
                                    case 4294967295:
                                        if v7 {
                                            switch "hello" {
                                            case "":
                                                var v42 bool = (!v7)
                                                _ = v42
                                            case "a":
                                                var v43 int32 = 0
                                                _ = v43
                                            default:
                                                switch ("hello" + "hello") {
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "a":
                                                        v2 -= v2
                                                    }
                                                default:
                                                    var v44 int16 = 176
                                                    _ = v44
                                                }
                                            }
                                        } else {
                                            fmt.Println(v41, v7)
                                        }
                                        var v45 int32 = -264
                                        switch "foo bar" {
                                        case "\n\t":
                                            if len(v5) > 0 {
                                                v5[0] = fi33
                                            }
                                        case "":
                                            if false {
                                                if v7 {
                                                    if (!(((v45 | v45) > v45) || false)) {
                                                        var v46 uint8 = 74
                                                        var v47 int8 = 42
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        if (v34 < v34) {
                                                            if true {
                                                                var v48 int64 = int64(v41)
                                                                var v49 float64 = 0.1
                                                                var v50 int16 = -565
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                        } else {
                                                            v34 = fi33
                                                        }
                                                        v3 = (v3 ^ v3)
                                                        v51 := GenS7{F0: 0, F1: "foo bar", F2: 42}
                                                        fmt.Println(v2, v41, v3)
                                                        var v52 int64 = int64(v3)
                                                        v53 := GenS7{F0: 0, F1: ("\n\t" + "a"), F2: 42}
                                                        v54 := GenS8{F0: -32768, F1: v3}
                                                        switch -9223372036854775808 {
                                                        case 127:
                                                            v5 = append(v5, fi33)
                                                        case -9223372036854775808:
                                                            var v55 int = v34
                                                            _ = v55
                                                        case -640:
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                        var v56 string = ("" + ((("foo bar" + (("\n\t" + ("hello" + ("\n\t" + "\n\t"))) + ("\n\t" + "hello"))) + "foo bar") + "\n\t"))
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                        _ = v56
                                                    }
                                                    switch 127 {
                                                    case 32767:
                                                        v57 := GenS9{F0: ""}
                                                        _ = v57
                                                    case 32768:
                                                        if (!((v45 << ((255) & 63)) == (v45 + v45))) {
                                                            fmt.Println(v34)
                                                            v2 = (-(fi32))
                                                        } else {
                                                            switch "" {
                                                            case "a":
                                                                if len(v5) > 0 {
                                                                    v5[0] = fi33
                                                                }
                                                            case "\n\t":
                                                                switch 32767 {
                                                                case 0:
                                                                    if v7 {
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                case 612:
                                                                    v58 := GenS10{F0: v3, F1: (v34 % ((v2) | 1)), F2: (fi33 | fi32)}
                                                                    _ = v58
                                                                }
                                                            default:
                                                                v59 := GenS11{F0: ((("foo bar" + (("a" + "") + "")) + ("" + "\n\t")) + (((("hello" + "hello") + "") + ((("hello" + "\n\t") + "a") + ("a" + "hello"))) + (("" + (("a" + ("hello" + "a")) + "foo bar")) + "a")))}
                                                                _ = v59
                                                            }
                                                            v60 := GenS3{F0: 100, F1: 9223372036854775807, F2: v45}
                                                            _ = v60
                                                        }
                                                    case 0:
                                                        if v1 {
                                                            v7 = (!v7)
                                                            var v61 bool = v7
                                                            fmt.Println(v61)
                                                            _ = v61
                                                        }
                                                    }
                                                    if (0 < -323) {
                                                        v62 := GenS7{F0: 128, F1: "a", F2: 230}
                                                        _ = v62
                                                    } else {
                                                        v63 := GenS9{F0: (("foo bar" + "a") + "\n\t")}
                                                        var v64 uint16 = 987
                                                        v1 = (43 < int8(v41))
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                } else {
                                                    v65 := GenS1{F0: 255, F1: 0, F2: 1}
                                                    v66 := GenS4{F0: -32768}
                                                    _ = v65
                                                    _ = v66
                                                }
                                                var v67 float64 = 3.14159
                                                if v7 {
                                                    v2 = (-(v2))
                                                }
                                                _ = v67
                                            }
                                        case "a":
                                            var v68 int32 = -1
                                            _ = v68
                                        default:
                                            var v69 bool = v7
                                            if (("hello" + ("foo bar" + "")) != "foo bar") {
                                                var v70 int8 = 42
                                                _ = v70
                                            }
                                            _ = v69
                                        }
                                        v5 = append(v5, v2)
                                        v3 -= (^(v3))
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v71 := GenS8{F0: 32767, F1: (v3 >> ((440) & 63))}
                                        if (v3 == (v3 + (^(v3)))) {
                                            v34 += (v2 | 32768)
                                            v41 -= v41
                                        } else {
                                            if (!(!(v2 >= -614))) {
                                                if (v45 < v45) {
                                                    v7 = v7
                                                } else {
                                                    var v72 string = "hello"
                                                    v73 := GenS3{F0: -88, F1: (v34 << ((0) & 63)), F2: v45}
                                                    var v74 uint16 = 32767
                                                    v7 = (v1 || v7)
                                                    v75 := GenS10{F0: 57, F1: (-(fi33)), F2: 128}
                                                    _ = v72
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                }
                                            }
                                            v76 := GenS4{F0: 1}
                                            _ = v76
                                        }
                                        _ = v45
                                        _ = v71
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    _ = v41
                                }
                                v3 = v3
                            }
                            v1 = (1 != 255)
                        }
                        switch ("foo bar" + ("\n\t" + "hello")) {
                        case "foo bar":
                            switch "" {
                            case "foo bar":
                                v5 = append(v5, fi32)
                            case "\n\t":
                                var v77 uint = 2
                                _ = v77
                            default:
                                _ = len(v5)
                            }
                        }
                    }
                } else {
                    var v78 int64 = 104
                    switch fi33 {
                    case -68:
                        if v7 {
                            _ = len(v5)
                            v1 = v7
                            if len(v5) > 0 {
                                v5[0] = v2
                            }
                        } else {
                            v79 := GenS12{F0: v1}
                            v2 = (fi33 >> ((42) & 63))
                            var v80 string = "\n\t"
                            fmt.Println(fi33)
                            if false {
                                v7 = (42 == 1000)
                                if (!(-743 != v2)) {
                                    var v81 int8 = -43
                                    fmt.Println(v2)
                                    _ = v81
                                }
                                if (42 > 2) {
                                    var v82 uint = 81
                                    v7 = ((fi32 * -1) != 9223372036854775807)
                                    v34 = (fi33 << (((v82 / ((481) | 1))) & 63))
                                    if (!v1) {
                                        var v83 int = -1
                                        fmt.Println(v3, v83)
                                        v84 := GenS13{F0: 42, F1: -101}
                                        if v7 {
                                            v85 := GenS2{F0: -1, F1: v78}
                                            v86 := GenS1{F0: 42, F1: 44, F2: 2}
                                            _ = v85
                                            _ = v86
                                        }
                                        _ = v83
                                        _ = v84
                                    }
                                    _ = v82
                                }
                            }
                            if (v7 || (209 >= 38)) {
                                v2 = v34
                            } else {
                                var v87 uint8 = 255
                                _ = v87
                            }
                            if v7 {
                                v7 = (!(!(!v1)))
                                var v88 uint32 = v3
                                _ = v88
                            }
                            if true {
                                v2 = v2
                                v89 := GenS10{F0: (v3 % ((1) | 1)), F1: 32768, F2: fi33}
                                v34 = 513
                                _ = v89
                            }
                            v90 := GenS6{F0: 42, F1: 0, F2: 0}
                            if len(v5) > 0 {
                                v5[0] = (-(v34))
                            }
                            _ = v79
                            _ = v80
                            _ = v90
                        }
                    case 1:
                        var v91 string = ""
                        _ = v91
                    }
                    v92 := GenS2{F0: 1000, F1: v78}
                    _ = v78
                    _ = v92
                }
            case "\n\t":
                if (512 <= 1) {
                    switch (fi32 % (((fi33 << ((487) & 63))) | 1)) {
                    case 1:
                        v1 = (v3 <= v3)
                    case 127:
                        v93 := GenS12{F0: (-2147483648 > (^(fi32)))}
                        _ = v93
                    case -490:
                        var v94 uint64 = 18446744073709551615
                        _ = v94
                    default:
                        var v95 int8 = -128
                        _ = v95
                    }
                } else {
                    _ = len(v5)
                    var v96 uint = 78
                    _ = v96
                }
            }
            switch ("hello" + "\n\t") {
            case "\n\t":
                if v1 {
                    var v97 uint32 = (v3 % (((v3 * (^(v3)))) | 1))
                    switch fi32 {
                    case -1:
                        if true {
                            var v98 int16 = -101
                            var v99 uint16 = 438
                            v100 := GenS8{F0: v98, F1: (^(v97))}
                            _ = v100
                            _ = v98
                            _ = v99
                        }
                    case 639:
                        v7 = ((1 == -2147483648) || (!v1))
                    default:
                        v5 = append(v5, (v2 + 129))
                    }
                    _ = v97
                } else {
                    v101 := GenS12{F0: (42 != 622)}
                    _ = v101
                }
            case "a":
                v102 := GenS1{F0: 0, F1: 42, F2: 42}
                _ = v102
            }
            _ = fi33
            _ = v34
        }
        v1 = (1 <= int32(v3))
        var v103 int8 = 42
        _ = fi32
        _ = v103
    }
    var v104 uint16 = 65535
    for fi105 := 0; fi105 < 7; fi105++ {
        var v106 int64 = 9223372036854775807
        fmt.Println(v7)
        switch (fi105 >> ((4294967295) & 63)) {
        case 9223372036854775807:
            fmt.Println(v7, v2, v104)
        }
        _ = fi105
        _ = v106
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v104
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
