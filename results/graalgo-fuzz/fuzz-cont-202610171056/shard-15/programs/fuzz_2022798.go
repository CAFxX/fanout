package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 string; F1 int16; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64; F1 uint64; F2 int64 }
type GenS10 struct { F0 int64; F1 uint8 }
type GenS11 struct { F0 int; F1 string; F2 string }
func genFn0() uint {
    v6 := []int{}
    var v7 int8 = -1
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int64 = 0
        _ = fi8
        _ = v9
    }
    v7 = (-(v7))
    _ = v6
    _ = v7
    return 2
}
func genFn1() int64 {
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 42, F1: -1}
        switch (fi10 % (((-(fi10))) | 1)) {
        case 1:
            var v12 bool = (!true)
            _ = v12
        case -1:
            switch fi10 {
            case 2147483647:
                var v13 int32 = -861
                _ = v13
            default:
                var v14 int16 = -1
                _ = v14
            }
        case 295:
            v15 := GenS2{F0: 42, F1: 1}
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 uint = 409
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS3{F0: (("hello" + (("hello" + "hello") + "\n\t")) + "hello"), F1: -32768, F2: 806}
        switch fi17 {
        case 27:
            var v19 uint = 42
            _ = v19
        case 1:
            fmt.Println(v16)
        default:
            v20 := []int{(-(fi17))}
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    var v21 float64 = -2.25
    _ = v16
    _ = v21
    return -1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 float64 = (-2.25 - 1.5)
    var v4 int = 32768
    v5 := []uint32{}
    v22 := GenS1{F0: 9223372036854775807, F1: 0}
    delete(v1, "\n\t")
    v2[9223372036854775807] = v4
    var v23 uint32 = 2147483648
    if (255 >= 0) {
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(v3)
            if (false || (0 == 1)) {
                v23 = (v23 | (v23 >> ((18446744073709551615) & 63)))
            } else {
                if len(v5) > 0 {
                    v5[0] = v23
                }
                _ = len(v5)
                v25 := GenS4{F0: 42}
                _ = v25
            }
            _ = fi24
        }
        var v26 bool = false
        _ = v26
    }
    var v27 int64 = -1
    if true {
        var v28 int64 = (v27 - (v27 / (((-(v27))) | 1)))
        switch "hello" {
        case "hello":
            switch ("hello" + "a") {
            case "a":
                if (!(!true)) {
                    v29 := GenS5{F0: (v23 + 2), F1: (v27 + -1), F2: (v23 ^ (v23 * 1))}
                    _ = v29
                } else {
                    if (((v27 ^ (v27 & (v27 ^ v27))) < (^(v27))) || true) {
                        fmt.Println(v23, v4, v4)
                        v23 = (v23 >> ((1) & 63))
                    } else {
                        var v30 bool = (845 != 2)
                        _ = v30
                    }
                    v28 += (v28 ^ v28)
                    fmt.Println(v4, v27)
                }
            case "hello":
                v3 += v3
            }
        case "foo bar":
            if true {
                for fi31 := 0; fi31 < 16; fi31++ {
                    if (32767 <= 322) {
                        v32 := GenS2{F0: 4294967295, F1: 382}
                        for fi33 := 0; fi33 < 16; fi33++ {
                            v5 = append(v5, (v23 ^ (v23 << ((65535) & 63))))
                            var v34 int64 = v27
                            fmt.Println(v28, v34, v3)
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    } else {
                        for fi35 := 0; fi35 < 15; fi35++ {
                            fmt.Println(v23, v3, v27)
                            _ = fi35
                        }
                        if ((v23 / (((v23 - v23)) | 1)) == 4294967295) {
                            var v36 uint64 = 4294967295
                            _ = v36
                        } else {
                            for fi37 := 0; fi37 < 4; fi37++ {
                                var v38 int32 = -831
                                v1[""] = -2147483648
                                if (42 >= 32767) {
                                    var v39 uint64 = 18446744073709551615
                                    var v40 string = ("a" + ((("" + ("a" + "a")) + ("foo bar" + ((("a" + "\n\t") + "a") + ""))) + "foo bar"))
                                    _ = v39
                                    _ = v40
                                }
                                _ = fi37
                                _ = v38
                            }
                            v23 -= v23
                        }
                    }
                    _ = fi31
                }
                v41 := GenS4{F0: 32767}
                fmt.Println(v28)
                _ = v41
            } else {
                switch "foo bar" {
                case "\n\t":
                    if (-259 > 1) {
                        switch "foo bar" {
                        case "":
                            delete(v2, (v4 * v4))
                        case "a":
                            fmt.Println(v28, v4, v3)
                        case "hello":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    }
                case "hello":
                    for fi42 := 0; fi42 < 16; fi42++ {
                        fmt.Println(fi42, v28)
                        if true {
                            fmt.Println(v4)
                        }
                        v23 += (v23 | v23)
                        _ = fi42
                    }
                case "":
                    var v43 uint = 42
                    _ = v43
                }
            }
        }
        _ = v28
    } else {
        fmt.Println(v4, v4, v3)
    }
    if false {
        if (v23 <= (v23 * v23)) {
            v23 += 1
            switch ("a" + "hello") {
            case "a":
                _ = len(v5)
            case "":
                switch "" {
                case "\n\t":
                    if (273 >= v4) {
                        v44 := GenS6{F0: true, F1: "hello"}
                        var v45 bool = ((1 > 465) || (4294967295 <= 18446744073709551615))
                        _ = v44
                        _ = v45
                    } else {
                        v46 := GenS1{F0: (v27 + int64(v4)), F1: 84}
                        if (128 != (v4 - v4)) {
                            var v47 uint16 = 42
                            if false {
                                v27 = (v27 >> ((727) & 63))
                            } else {
                                if false {
                                    for fi48 := 0; fi48 < 13; fi48++ {
                                        v27 = (v27 + -9223372036854775808)
                                        v4 = fi48
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = fi48
                                    }
                                    fmt.Println(v23)
                                } else {
                                    v49 := GenS7{F0: 2, F1: v23}
                                    _ = v49
                                }
                                if true {
                                    var v50 uint = 4294967295
                                    v51 := GenS1{F0: 42, F1: -13}
                                    var v52 uint8 = 255
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                } else {
                                    if true {
                                        switch ("" + "a") {
                                        case "foo bar":
                                            v47 -= v47
                                        default:
                                            switch (v4 % (((v4 * -1)) | 1)) {
                                            case 42:
                                                for fi53 := 0; fi53 < 15; fi53++ {
                                                    v54 := GenS8{F0: 1000}
                                                    _ = fi53
                                                    _ = v54
                                                }
                                            }
                                        }
                                        _ = len(v2)
                                        fmt.Println(v47, v47, v4)
                                    } else {
                                        v55 := GenS3{F0: "\n\t", F1: -1, F2: 389}
                                        _ = v55
                                    }
                                    v3 = (v3 / v3)
                                }
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v47
                        }
                        _ = v46
                    }
                }
            default:
                switch (^(v4)) {
                case 2147483647:
                    for fi56 := 0; fi56 < 4; fi56++ {
                        var v57 int = (-(v4))
                        v4 += (fi56 << ((18446744073709551615) & 63))
                        delete(v1, "a")
                        _ = fi56
                        _ = v57
                    }
                default:
                    var v58 uint = 4294967295
                    _ = v58
                }
            }
            if len(v5) > 0 {
                v5[0] = (v23 | 1)
            }
        } else {
            v4 = -754
            for fi59 := 0; fi59 < 4; fi59++ {
                if true {
                    var v60 bool = (!(0 == 1))
                    if (172 != 775) {
                        if v60 {
                            v5 = append(v5, 2)
                            var v61 int8 = 100
                            var v62 int8 = v61
                            _ = v61
                            _ = v62
                        }
                    } else {
                        v63 := v1["foo bar"]
                        if v60 {
                            var v64 float64 = 0.0
                            _ = v64
                        } else {
                            v65 := v2[fi59]
                            switch ((("\n\t" + "hello") + "hello") + "hello") {
                            case "":
                                v60 = (!v60)
                            case "hello":
                                switch "foo bar" {
                                case "a":
                                    var v66 uint32 = 2147483647
                                    _ = v66
                                default:
                                    if v60 {
                                        v67 := GenS8{F0: -32768}
                                        var v68 int16 = 42
                                        _ = v67
                                        _ = v68
                                    } else {
                                        var v69 float32 = -2.25
                                        for fi70 := 0; fi70 < 2; fi70++ {
                                            v69 -= (v69 - ((v69 - v69) - (v69 - v69)))
                                            _ = fi70
                                        }
                                        v71 := GenS5{F0: v23, F1: -1, F2: (v23 / ((2) | 1))}
                                        _ = v69
                                        _ = v71
                                    }
                                }
                            case "foo bar":
                                v72 := GenS9{F0: -427, F1: 0, F2: -247}
                                _ = v72
                            }
                            _ = v65
                        }
                        for fi73 := 0; fi73 < 20; fi73++ {
                            var v74 uint8 = 255
                            _ = fi73
                            _ = v74
                        }
                        _ = v63
                    }
                    _ = v60
                } else {
                    var v75 uint16 = 255
                    switch "\n\t" {
                    case "foo bar":
                        v27 -= 176
                    case "hello":
                        v76 := GenS9{F0: (v27 - (^(v27))), F1: 18446744073709551615, F2: genFn1()}
                        _ = v76
                    case "":
                        for fi77 := 0; fi77 < 11; fi77++ {
                            v5 = append(v5, v23)
                            v78 := GenS10{F0: genFn1(), F1: 2}
                            _ = fi77
                            _ = v78
                        }
                    }
                    if (((("a" + "hello") + "\n\t") + "foo bar") >= ("\n\t" + "")) {
                        var v79 int32 = 42
                        _ = v79
                    }
                    _ = v75
                }
                if (true || (54 < 42)) {
                    var v80 uint = 2
                    v80 = (v80 / (((^(v80))) | 1))
                    if true {
                        var v81 float32 = 0.1
                        v23 = v23
                        _ = v81
                    } else {
                        if false {
                            v82 := v1["foo bar"]
                            v4 -= v82
                            v2[(fi59 << ((255) & 63))] = (v4 % (((v82 ^ (^(v82)))) | 1))
                            _ = v82
                        }
                        for fi83 := 0; fi83 < 16; fi83++ {
                            v23 = 4294967295
                            _ = fi83
                        }
                        v84 := GenS9{F0: (-(v27)), F1: 18446744073709551615, F2: (v27 / (((v27 * -802)) | 1))}
                        _ = v84
                    }
                    _ = v80
                } else {
                    v85 := GenS10{F0: genFn1(), F1: 127}
                    _ = v85
                }
                v23 = (v23 >> ((4294967295) & 63))
                _ = fi59
            }
            var v86 int = -673
            _ = v86
        }
    } else {
        for fi87 := 0; fi87 < 6; fi87++ {
            for fi88 := 0; fi88 < 6; fi88++ {
                fmt.Println(fi87, fi88, v23)
                if (1.5 == (0.1 / 1.0)) {
                    var v89 int8 = 42
                    _ = v89
                } else {
                    _ = len(v2)
                    var v90 uint16 = 174
                    v91 := GenS2{F0: 1, F1: 2}
                    _ = v90
                    _ = v91
                }
                _ = fi88
            }
            v27 = genFn1()
            _ = fi87
        }
        if ((v23 << ((228) & 63)) == (v23 * (v23 * v23))) {
            v4 = (v4 ^ 128)
        } else {
            if ("\n\t" <= (("" + "\n\t") + "foo bar")) {
                switch 42 {
                case 2147483647:
                    v92 := GenS9{F0: genFn1(), F1: 290, F2: (v27 << ((4294967295) & 63))}
                    _ = v92
                case -457:
                    v93 := GenS9{F0: (v27 | (v27 ^ genFn1())), F1: 4294967295, F2: (v27 * (v27 >> ((0) & 63)))}
                    _ = v93
                default:
                    for fi94 := 0; fi94 < 2; fi94++ {
                        for fi95 := 0; fi95 < 13; fi95++ {
                            v96 := v1[("hello" + "foo bar")]
                            v3 = 1.5
                            _ = fi95
                            _ = v96
                        }
                        v4 -= v4
                        _ = fi94
                    }
                }
            } else {
                switch "hello" {
                case "foo bar":
                    var v97 string = ("hello" + ("a" + (("" + "") + ("a" + "\n\t"))))
                    _ = v97
                case "a":
                    var v98 int8 = 1
                    _ = v98
                }
                fmt.Println(v3)
                for fi99 := 0; fi99 < 2; fi99++ {
                    if true {
                        v3 += v3
                        switch 1 {
                        case -2147483648:
                            v3 -= 0.0
                        case -281:
                            switch v4 {
                            case -372:
                                var v100 int = 128
                                _ = v100
                            default:
                                for fi101 := 0; fi101 < 13; fi101++ {
                                    var v102 int64 = 1
                                    _ = fi101
                                    _ = v102
                                }
                            }
                        default:
                            v2[v4] = (fi99 % (((v4 % (((fi99 / ((42) | 1))) | 1))) | 1))
                        }
                        v4 -= fi99
                    }
                    _ = fi99
                }
            }
            v103 := GenS6{F0: (((127 <= 42) || (!true)) && false), F1: "hello"}
            _ = v103
        }
        var v104 int16 = -693
        _ = v104
    }
    if ((0 == (v27 * v27)) && true) {
        var v105 int64 = v27
        _ = len(v5)
        fmt.Println(v4, v23, v23)
        _ = v105
    } else {
        if ("hello" <= "\n\t") {
            _ = len(v5)
            var v106 int64 = (v27 << ((genFn0()) & 63))
            var v107 uint8 = 0
            _ = v106
            _ = v107
        }
    }
    v108 := GenS11{F0: (v4 - v4), F1: ("a" + ("hello" + ("" + "\n\t"))), F2: ("" + ("a" + ("" + ("a" + "foo bar"))))}
    v109 := v1[""]
    v2[32768] = (-(v109))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
