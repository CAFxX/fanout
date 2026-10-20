package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int64 }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 string; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 int64; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 int8; F1 int64 }
type GenS11 struct { F0 bool }
func genFn0(p0 uint8, p1 uint) uint16 {
    var v7 int64 = int64(p0)
    var v8 uint16 = 1
    var v9 bool = true
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1() int32 {
    var v10 string = "hello"
    fmt.Println(v10)
    v10 = ("a" + ("hello" + v10))
    for fi11 := 0; fi11 < 13; fi11++ {
        v10 = v10
        _ = fi11
    }
    var v12 float64 = (0.1 / ((((0.1 / 1.0) - 0.1) + 3.14159) + 0.0))
    _ = v10
    _ = v12
    return 2147483647
}

func main() {
    var v1 uint16 = 983
    v2 := make(map[string]bool)
    var v3 uint8 = 1
    var v4 bool = true
    v5 := make(map[int]int)
    v6 := []string{}
    switch "a" {
    case "a":
        v2["\n\t"] = (18446744073709551615 > 42)
    case "foo bar":
        for fi13 := 0; fi13 < 13; fi13++ {
            if len(v6) > 0 {
                v6[0] = "\n\t"
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                switch ("\n\t" + "a") {
                case "a":
                    v15 := v2[((((("a" + ("" + ("hello" + "\n\t"))) + "\n\t") + ("" + "foo bar")) + "a") + "a")]
                    _ = v15
                case "":
                    v16 := GenS1{F0: -547, F1: 1, F2: -952}
                    _ = v16
                case "foo bar":
                    if ((v3 - 1) <= (v3 >> ((4294967295) & 63))) {
                        var v17 string = "hello"
                        fmt.Println(v17)
                        _ = v17
                    } else {
                        var v18 float32 = ((((0.0 + 0.1) * 0.0) * 1e10) - (-2.25 - 1e10))
                        v19 := GenS1{F0: 0, F1: fi14, F2: -1}
                        var v20 int = (^(fi13))
                        v21 := GenS2{F0: (v20 / ((v20) | 1)), F1: 506, F2: v4}
                        var v22 int64 = 663
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                default:
                    var v23 uint8 = v3
                    _ = v23
                }
                switch "a" {
                case "hello":
                    v24 := GenS3{F0: "\n\t", F1: v3, F2: v3}
                    _ = v24
                case "\n\t":
                    delete(v5, fi13)
                }
                _ = fi14
            }
            _ = fi13
        }
    default:
        if (1e10 < 3.14159) {
            if (-9223372036854775808 == -1) {
                for fi25 := 0; fi25 < 5; fi25++ {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        v27 := GenS1{F0: 1, F1: fi25, F2: -9223372036854775808}
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v6 = append(v6, "")
                        var v28 int32 = -2147483648
                        v4 = true
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    var v29 float32 = 3.14159
                    for fi30 := 0; fi30 < 6; fi30++ {
                        fmt.Println(v4)
                        _ = fi30
                    }
                    _ = fi25
                    _ = v29
                }
            } else {
                var v31 uint32 = 0
                for fi32 := 0; fi32 < 10; fi32++ {
                    v3 += v3
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi32
                }
                v4 = ("" != "foo bar")
                _ = v31
            }
            v33 := v5[-2147483648]
            _ = v33
        }
    }
    var v34 int8 = 100
    v35 := GenS1{F0: 0, F1: 0, F2: -634}
    v36 := GenS1{F0: 0, F1: -530, F2: -631}
    v1 += (v1 % ((genFn0((v3 * (v3 % (((v3 << ((197) & 63))) | 1))), 0)) | 1))
    var v37 uint32 = 65535
    if v4 {
        for fi38 := 0; fi38 < 8; fi38++ {
            v39 := GenS4{F0: -1}
            v40 := v2[(("\n\t" + "") + (("foo bar" + ("" + "foo bar")) + ("hello" + "\n\t")))]
            _ = fi38
            _ = v39
            _ = v40
        }
        switch "a" {
        case "foo bar":
            for fi41 := 0; fi41 < 2; fi41++ {
                if ((!(42 != 18446744073709551615)) || (1 >= genFn1())) {
                    v42 := v2[("foo bar" + "a")]
                    fmt.Println(v37, v4, v42)
                    for fi43 := 0; fi43 < 4; fi43++ {
                        switch 32767 {
                        case -9223372036854775808:
                            switch (fi43 % (((fi41 | fi43)) | 1)) {
                            case 128:
                                switch fi43 {
                                case -406:
                                    v3 -= v3
                                case 32768:
                                    v6 = append(v6, ("a" + ""))
                                }
                            case 127:
                                v44 := GenS1{F0: 1000, F1: (fi43 >> ((2) & 63)), F2: 1}
                                _ = v44
                            case 42:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            default:
                                v3 = 42
                            }
                        case 32768:
                            if false {
                                v2["a"] = (0 != 1)
                                if v42 {
                                    v6 = append(v6, ((("foo bar" + "hello") + "a") + "\n\t"))
                                    if v42 {
                                        switch ((("a" + (("" + "foo bar") + "\n\t")) + ("a" + "a")) + "") {
                                        case "\n\t":
                                            var v45 int64 = -637
                                            _ = v45
                                        case "a":
                                            if len(v6) > 0 {
                                                v6[0] = "\n\t"
                                            }
                                        case "hello":
                                            if len(v6) > 0 {
                                                v6[0] = ("foo bar" + "")
                                            }
                                        default:
                                            switch "a" {
                                            case "hello":
                                                v46 := GenS1{F0: 42, F1: fi43, F2: 1}
                                                _ = v46
                                            case "\n\t":
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v37 = v37
                                                default:
                                                    var v47 uint8 = 153
                                                    _ = v47
                                                }
                                            default:
                                                var v48 int = -663
                                                _ = v48
                                            }
                                        }
                                        v4 = (22 <= (v34 - (^(v34))))
                                    } else {
                                        var v49 float64 = 0.0
                                        if v42 {
                                            var v50 int16 = -32768
                                            var v51 int16 = v50
                                            fmt.Println(v51, v37, v1)
                                            _ = v50
                                            _ = v51
                                        } else {
                                            v52 := GenS1{F0: -1, F1: -2147483648, F2: -398}
                                            v4 = (uint64(v3) <= 662)
                                            switch fi41 {
                                            case 4294967295:
                                                var v53 uint8 = 255
                                                _ = v53
                                            case -2147483648:
                                                v54 := GenS5{F0: 4294967295, F1: 0, F2: (v34 | v34)}
                                                _ = v54
                                            case 9223372036854775807:
                                                if (!true) {
                                                    if true {
                                                        if (v42 && v4) {
                                                            fmt.Println(fi43, fi41, fi41)
                                                            v34 = 0
                                                            if ((true || (65535 <= 77)) || (v37 < (v37 * 42))) {
                                                                switch fi43 {
                                                                case 235:
                                                                    if len(v6) > 0 {
                                                                        v6[0] = ""
                                                                    }
                                                                }
                                                                v49 += ((v49 * 1.5) - 0.0)
                                                            } else {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                var v55 int64 = 1
                                                                _ = v55
                                                            }
                                                        } else {
                                                            fmt.Println(v37, v3)
                                                            fmt.Println(v1, fi41)
                                                            v42 = (2147483647 == 1)
                                                        }
                                                        if true {
                                                            var v56 float32 = 0.1
                                                            _ = v56
                                                        }
                                                        switch "hello" {
                                                        case "foo bar":
                                                            delete(v5, -9223372036854775808)
                                                        default:
                                                            if v42 {
                                                                fmt.Println(fi43, v3, v1)
                                                                switch (fi41 ^ (fi41 ^ (fi43 / ((612) | 1)))) {
                                                                case 127:
                                                                    fmt.Println(v4, v4)
                                                                case 0:
                                                                    var v57 float64 = (float64((-2.25 * 0.1)) - v49)
                                                                    _ = v57
                                                                default:
                                                                    v6 = append(v6, "hello")
                                                                }
                                                            } else {
                                                                var v58 int32 = -1
                                                                v59 := GenS3{F0: "hello", F1: (v3 - v3), F2: (v3 * 42)}
                                                                v60 := GenS6{F0: ("" + "hello")}
                                                                var v61 uint = 4294967295
                                                                if ((v61 < 122) && false) {
                                                                    v62 := GenS7{F0: 15, F1: v3}
                                                                    fmt.Println(v61, v37, v61)
                                                                    switch (fi43 ^ 127) {
                                                                    case -9223372036854775808:
                                                                        var v63 uint8 = 255
                                                                        _ = v63
                                                                    case 778:
                                                                        if true {
                                                                            v1 -= (v1 | uint16(v37))
                                                                            v64 := GenS8{F0: genFn1(), F1: genFn0((v3 << ((v61) & 63)), uint(v1)), F2: v1}
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            _ = v64
                                                                        }
                                                                    case 127:
                                                                        v58 -= v58
                                                                    }
                                                                    _ = v62
                                                                }
                                                                _ = v58
                                                                _ = v59
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v52
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            v34 = 100
                                        case "":
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        case "\n\t":
                                            if true {
                                                v65 := GenS1{F0: 32767, F1: (fi41 % ((487) | 1)), F2: 42}
                                                fmt.Println(v1, v1, fi43)
                                                fmt.Println(v34, fi43, fi41)
                                                _ = v65
                                            }
                                        default:
                                            if true {
                                                var v66 int = (fi41 | int(v37))
                                                _ = v66
                                            } else {
                                                var v67 int8 = 1
                                                _ = v67
                                            }
                                        }
                                        _ = v49
                                    }
                                    switch "hello" {
                                    case "a":
                                        v5[127] = (fi43 >> ((18446744073709551615) & 63))
                                    }
                                    switch 518 {
                                    case 2147483647:
                                        v3 -= 127
                                    case -1:
                                        var v68 string = "foo bar"
                                        _ = v68
                                    }
                                } else {
                                    v4 = ((!((uint16(v37) > (v1 ^ 1)) || v42)) || (-1 == -32768))
                                }
                            }
                            v34 += v34
                            var v69 uint64 = 458
                            _ = v69
                        default:
                            delete(v2, "\n\t")
                        }
                        if v42 {
                            v70 := GenS3{F0: "hello", F1: 42, F2: v3}
                            v1 += (^(v1))
                            _ = v70
                        }
                        _ = fi43
                    }
                    _ = v42
                }
                for fi71 := 0; fi71 < 13; fi71++ {
                    v3 -= 127
                    var v72 int = 42
                    var v73 uint16 = uint16(v37)
                    v34 += 39
                    _ = fi71
                    _ = v72
                    _ = v73
                }
                switch (fi41 << ((65535) & 63)) {
                case -3:
                    var v74 uint64 = 145
                    _ = v74
                case 9223372036854775807:
                    v75 := GenS4{F0: 1}
                    _ = v75
                default:
                    for fi76 := 0; fi76 < 13; fi76++ {
                        switch (-(fi76)) {
                        case -321:
                            var v77 uint8 = 0
                            _ = v77
                        case 42:
                            v4 = ((uint16(v3) >= (v1 * uint16(v37))) || true)
                        case -967:
                            var v78 uint32 = (^(v37))
                            _ = v78
                        default:
                            _ = len(v5)
                        }
                        v79 := GenS2{F0: fi41, F1: 18446744073709551615, F2: true}
                        var v80 int8 = -1
                        _ = fi76
                        _ = v79
                        _ = v80
                    }
                }
                _ = fi41
            }
        case "a":
            var v81 string = "foo bar"
            _ = v81
        case "hello":
            v37 += v37
        default:
            delete(v5, -9223372036854775808)
        }
    }
    for fi82 := 0; fi82 < 1; fi82++ {
        v83 := GenS7{F0: v34, F1: uint8(v37)}
        _ = fi82
        _ = v83
    }
    for fi84 := 0; fi84 < 13; fi84++ {
        var v85 float32 = 0.1
        _ = fi84
        _ = v85
    }
    v86 := v5[1]
    if v4 {
        if v4 {
            var v87 float32 = 1.5
            for fi88 := 0; fi88 < 18; fi88++ {
                v87 += (0.0 + v87)
                v89 := GenS7{F0: (^(v34)), F1: 0}
                _ = fi88
                _ = v89
            }
            for fi90 := 0; fi90 < 5; fi90++ {
                v34 = v34
                v5[(fi90 << ((1) & 63))] = v86
                for fi91 := 0; fi91 < 1; fi91++ {
                    v34 = (^(v34))
                    v1 -= genFn0(v3, 21)
                    _ = fi91
                }
                _ = fi90
            }
            _ = v87
        } else {
            v4 = ((^(v3)) == (v3 & v3))
            v92 := GenS4{F0: -190}
            if v4 {
                var v93 int32 = 819
                switch ((("hello" + "hello") + ("" + ("hello" + ("a" + "\n\t")))) + "foo bar") {
                case "a":
                    v1 = (v1 - (v1 >> ((255) & 63)))
                default:
                    if false {
                        v37 = 2147483647
                        switch ("foo bar" + "") {
                        case "":
                            if v4 {
                                v37 = (v37 / (((v37 >> ((255) & 63))) | 1))
                                v1 = genFn0(1, 4294967295)
                                var v94 uint32 = (^(v37))
                                _ = v94
                            } else {
                                fmt.Println(v86)
                                v95 := GenS1{F0: int16(v1), F1: v86, F2: 9223372036854775807}
                                _ = v95
                            }
                        case "\n\t":
                            v96 := GenS4{F0: 0}
                            _ = v96
                        default:
                            v37 -= 42
                        }
                    } else {
                        v34 += 101
                        fmt.Println(v1)
                    }
                }
                var v97 uint16 = v1
                _ = v93
                _ = v97
            } else {
                v37 = v37
                if v4 {
                    if len(v6) > 0 {
                        v6[0] = "hello"
                    }
                } else {
                    v98 := GenS5{F0: 2, F1: 9223372036854775807, F2: (v34 ^ (-(v34)))}
                    for fi99 := 0; fi99 < 2; fi99++ {
                        fmt.Println(v37, v3)
                        _ = fi99
                    }
                    _ = v98
                }
            }
            _ = v92
        }
        switch (v86 ^ -2147483648) {
        case -2147483648:
            if false {
                var v100 string = ("foo bar" + "")
                if true {
                    v86 -= (v86 >> ((uint(v1)) & 63))
                }
                if false {
                    delete(v5, v86)
                    fmt.Println(v4)
                } else {
                    switch ("hello" + v100) {
                    case "":
                        for fi101 := 0; fi101 < 8; fi101++ {
                            v102 := GenS9{F0: v1, F1: (!v4)}
                            delete(v2, "\n\t")
                            v5[(v86 ^ (v86 | (fi101 / ((fi101) | 1))))] = -1
                            _ = fi101
                            _ = v102
                        }
                    case "foo bar":
                        fmt.Println(v1, v100, v86)
                    case "\n\t":
                        var v103 uint = 892
                        _ = v103
                    }
                }
                _ = v100
            } else {
                fmt.Println(v34, v34, v1)
            }
        }
    }
    var v104 uint16 = (v1 << ((4294967295) & 63))
    if false {
        v105 := GenS10{F0: (v34 ^ (v34 << ((18446744073709551615) & 63))), F1: -1}
        if len(v6) > 0 {
            v6[0] = "\n\t"
        }
        var v106 int32 = -1
        _ = v105
        _ = v106
    }
    switch "hello" {
    case "\n\t":
        for fi107 := 0; fi107 < 9; fi107++ {
            if v4 {
                v108 := v2[("a" + "a")]
                v104 = genFn0(v3, uint(v37))
                _ = v108
            }
            if (!false) {
                v109 := GenS11{F0: (!(-708 <= -1))}
                _ = v109
            }
            if v4 {
                v104 -= (v1 >> ((44) & 63))
                var v110 uint8 = 128
                _ = v110
            }
            _ = fi107
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v104
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v86
}
