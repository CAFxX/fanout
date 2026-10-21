package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int16; F1 int64; F2 uint16 }
type GenS4 struct { F0 int16; F1 bool; F2 string }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int32; F2 int64 }
type GenS8 struct { F0 uint; F1 int32 }
func genFn0(p0 int32, p1 uint32) uint64 {
    fmt.Println(p1)
    p1 -= (^(p1))
    _ = p0
    _ = p1
    return 18446744073709551615
}
func genFn1() int16 {
    v5 := GenS1{F0: 32767, F1: ("foo bar" + ""), F2: 4294967295}
    v6 := []int{-9223372036854775808, -2147483648}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if ((3.14159 - -2.25) > (-2.25 * 0.1)) {
        switch "foo bar" {
        case "foo bar":
            if len(v6) > 0 {
                v6[0] = 540
            }
        case "hello":
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: (fi7 ^ fi7), F1: "foo bar", F2: 1}
                _ = fi7
                _ = v8
            }
        default:
            if false {
                v6 = append(v6, 32768)
            }
        }
        var v9 uint8 = 127
        _ = len(v6)
        _ = v9
    } else {
        _ = len(v6)
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 float64 = (-2.25 + 0.1)
                if false {
                    var v13 int16 = 0
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        var v14 int16 = -32768
        _ = v14
    }
    _ = v5
    _ = v6
    return -416
}

func main() {
    var v1 uint16 = 652
    var v2 int8 = -31
    v3 := []string{"\n\t", "hello", "\n\t"}
    v4 := make(map[string]string)
    for fi15 := 0; fi15 < 1; fi15++ {
        if true {
            for fi16 := 0; fi16 < 16; fi16++ {
                if false {
                    if false {
                        v17 := GenS2{F0: 420, F1: -1}
                        _ = v17
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                        var v18 uint64 = 4294967295
                        v19 := GenS2{F0: 255, F1: -164}
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS1{F0: fi16, F1: "a", F2: 576}
                    v3 = append(v3, (("a" + (("hello" + "") + "\n\t")) + ""))
                    _ = v20
                } else {
                    v2 -= v2
                }
                if true {
                    v2 = v2
                    v21 := GenS2{F0: 255, F1: 1}
                    if false {
                        delete(v4, "\n\t")
                        var v22 uint = 443
                        v22 = (v22 + 1)
                        _ = v22
                    }
                    _ = v21
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        if (true && (!true)) {
                            if false {
                                v23 := GenS3{F0: 234, F1: -1, F2: v1}
                                _ = v23
                            } else {
                                if (32767 <= 0) {
                                    fmt.Println(fi16, fi16)
                                }
                                switch -9223372036854775808 {
                                case -625:
                                    if false {
                                        v1 = 2
                                        if false {
                                            fmt.Println(v1, v2, v1)
                                        } else {
                                            if true {
                                                v2 -= (v2 * -111)
                                                if false {
                                                    v3 = append(v3, ("" + ("foo bar" + "foo bar")))
                                                } else {
                                                    v1 = v1
                                                    v3 = append(v3, "")
                                                    if false {
                                                        v1 += (v1 + v1)
                                                        v2 -= (v2 + (-(v2)))
                                                        v1 += (^(v1))
                                                    }
                                                }
                                            }
                                            switch (fi15 - (fi15 >> ((1) & 63))) {
                                            case -9223372036854775808:
                                                switch -9223372036854775808 {
                                                case 32768:
                                                    if ((4294967295 > 2147483647) || (945 >= 0)) {
                                                        v24 := GenS2{F0: 255, F1: 42}
                                                        delete(v4, "a")
                                                        _ = v24
                                                    } else {
                                                        v1 -= v1
                                                    }
                                                case 485:
                                                    v4[("" + ((("a" + "\n\t") + ("hello" + "a")) + "foo bar"))] = ""
                                                case -666:
                                                    fmt.Println(fi16, fi16, v1)
                                                }
                                            case 127:
                                                var v25 uint64 = genFn0(-2147483648, 632)
                                                _ = v25
                                            case 61:
                                                if (-41 != (v2 * (v2 % ((v2) | 1)))) {
                                                    v2 = (v2 / ((v2) | 1))
                                                    if ((^(v1)) > (^(v1))) {
                                                        var v26 float64 = 1e10
                                                        _ = v26
                                                    }
                                                    v2 -= 1
                                                } else {
                                                    var v27 float64 = (-2.25 / 3.14159)
                                                    var v28 int64 = -345
                                                    _ = v27
                                                    _ = v28
                                                }
                                            default:
                                                if (true && false) {
                                                    var v29 uint8 = 0
                                                    if (!false) {
                                                        if false {
                                                            var v30 int32 = 42
                                                            fmt.Println(v1)
                                                            _ = v30
                                                        } else {
                                                            var v31 int = fi16
                                                            if true {
                                                                switch (v31 - (fi15 - 70)) {
                                                                case 2147483647:
                                                                    fmt.Println(v1, v31, fi16)
                                                                case 32767:
                                                                    if false {
                                                                        delete(v4, "a")
                                                                        v3 = append(v3, (("hello" + "\n\t") + ("a" + "")))
                                                                    }
                                                                default:
                                                                    if (9223372036854775807 < 194) {
                                                                        v29 -= 2
                                                                        fmt.Println(v29, v29, v31)
                                                                    }
                                                                }
                                                                v2 = v2
                                                            } else {
                                                                v29 -= uint8(v1)
                                                                var v32 float64 = (0.0 - 0.1)
                                                                v33 := GenS1{F0: (fi16 << ((uint(v1)) & 63)), F1: ("a" + ("foo bar" + (("hello" + "hello") + "\n\t"))), F2: 1}
                                                                _ = v32
                                                                _ = v33
                                                            }
                                                            v34 := GenS3{F0: -32768, F1: 1, F2: (v1 << ((0) & 63))}
                                                            _ = v31
                                                            _ = v34
                                                        }
                                                    } else {
                                                        v2 = (v2 * (v2 << ((18446744073709551615) & 63)))
                                                        var v35 int16 = 1
                                                        _ = v35
                                                    }
                                                    var v36 uint64 = 1
                                                    _ = v29
                                                    _ = v36
                                                }
                                            }
                                        }
                                    }
                                case 42:
                                    v1 += (^(v1))
                                case -943:
                                    v1 = (v1 << ((102) & 63))
                                }
                                v37 := GenS3{F0: -32768, F1: -900, F2: v1}
                                var v38 string = "hello"
                                _ = v37
                                _ = v38
                            }
                            fmt.Println(v2, v1, fi16)
                            switch "hello" {
                            case "":
                                v39 := GenS3{F0: 0, F1: 0, F2: 32767}
                                _ = v39
                            default:
                                v2 += (v2 / ((100) | 1))
                            }
                        } else {
                            if (v2 == -128) {
                                var v40 bool = true
                                _ = v40
                            } else {
                                v2 = (v2 >> ((18446744073709551615) & 63))
                                if true {
                                    if (("" + "") < (("" + "a") + "hello")) {
                                        v2 -= (v2 | v2)
                                        fmt.Println(v1, fi15, fi15)
                                    }
                                    if (255 <= 59) {
                                        v1 = v1
                                        fmt.Println(v1, v2, v2)
                                    }
                                    v1 = v1
                                }
                                switch (fi15 - -759) {
                                case -870:
                                    var v41 int8 = v2
                                    _ = v41
                                case 9223372036854775807:
                                    v1 = (^(v1))
                                case 419:
                                    switch 42 {
                                    case 9223372036854775807:
                                        var v42 uint32 = 754
                                        _ = v42
                                    default:
                                        var v43 string = ("hello" + "\n\t")
                                        _ = v43
                                    }
                                default:
                                    var v44 int32 = 42
                                    _ = v44
                                }
                            }
                        }
                    case "":
                        if true {
                            if true {
                                fmt.Println(v2, fi16)
                            }
                            v3 = append(v3, ("foo bar" + "hello"))
                            if true {
                                var v45 uint16 = v1
                                if false {
                                    var v46 uint64 = 42
                                    _ = v46
                                }
                                v45 = v45
                                v47 := v4[""]
                                v45 -= (^(v45))
                                _ = v45
                                _ = v47
                            }
                        }
                    case "\n\t":
                        v2 = v2
                    default:
                        var v48 int16 = 42
                        _ = v48
                    }
                    if (!false) {
                        v1 = (v1 - (v1 + (v1 ^ v1)))
                        var v49 int8 = (v2 & (v2 ^ v2))
                        v50 := GenS4{F0: 32767, F1: false, F2: ""}
                        fmt.Println(v1, fi16)
                        switch (^(fi16)) {
                        case -9223372036854775808:
                            var v51 int8 = v49
                            _ = v51
                        case 32768:
                            var v52 string = ""
                            _ = v52
                        }
                        _ = v49
                        _ = v50
                    }
                    var v53 float32 = ((1e10 + (1e10 * -2.25)) - 1e10)
                    _ = v53
                }
                v1 += v1
                _ = fi16
            }
            var v54 uint64 = 2
            fmt.Println(v2)
            _ = v54
        }
        _ = fi15
    }
    v55 := GenS3{F0: -1, F1: -88, F2: (v1 + (v1 / (((v1 << ((0) & 63))) | 1)))}
    var v56 uint8 = 255
    var v57 int = -682
    v3 = append(v3, "\n\t")
    for fi58 := 0; fi58 < 7; fi58++ {
        var v59 bool = (1e10 > 3.14159)
        for fi60 := 0; fi60 < 9; fi60++ {
            v57 = (fi60 << ((255) & 63))
            _ = fi60
        }
        v61 := GenS4{F0: int16(v2), F1: true, F2: "foo bar"}
        _ = fi58
        _ = v59
        _ = v61
    }
    var v62 int = -1
    switch "a" {
    case "\n\t":
        var v63 float32 = float32(0.1)
        _ = v63
    case "foo bar":
        v1 = (v1 & (v1 >> ((0) & 63)))
    case "a":
        v64 := GenS5{F0: 914, F1: (v56 | (v56 / ((v56) | 1)))}
        _ = v64
    }
    v65 := GenS6{F0: v2, F1: 0}
    for fi66 := 0; fi66 < 20; fi66++ {
        if true {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            var v67 bool = (65535 > 4294967295)
            v67 = v67
            _ = v67
        } else {
            var v68 int64 = -9223372036854775808
            for fi69 := 0; fi69 < 15; fi69++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v3 = append(v3, ("hello" + ""))
                _ = fi69
            }
            _ = v68
        }
        var v70 int16 = -1
        _ = fi66
        _ = v70
    }
    for fi71 := 0; fi71 < 18; fi71++ {
        switch (v57 - v57) {
        case 2147483647:
            for fi72 := 0; fi72 < 6; fi72++ {
                var v73 int16 = 32767
                if false {
                    fmt.Println(v1, v73)
                    if (genFn0(0, 841) == uint64(v1)) {
                        var v74 uint16 = v1
                        _ = v74
                    } else {
                        v2 -= v2
                    }
                    v62 += v57
                }
                _ = fi72
                _ = v73
            }
        case -9223372036854775808:
            fmt.Println(v57, v2)
        case 866:
            v75 := GenS5{F0: 0, F1: 78}
            _ = v75
        }
        v1 = (v1 << ((65535) & 63))
        _ = fi71
    }
    switch (v57 ^ 9223372036854775807) {
    case 42:
        if false {
            var v76 float64 = 1.5
            _ = v76
        } else {
            v3 = append(v3, "\n\t")
            _ = len(v3)
        }
    case -9223372036854775808:
        switch ("a" + "hello") {
        case "a":
            v4["a"] = ""
        case "":
            v56 -= (v56 >> ((42) & 63))
        case "\n\t":
            if false {
                var v77 float32 = 1e10
                v1 -= v1
                fmt.Println(v56, v1)
                _ = v77
            } else {
                v78 := GenS3{F0: genFn1(), F1: 644, F2: (v1 / (((v1 + v1)) | 1))}
                if false {
                    if ((^(v56)) < v56) {
                        v79 := GenS7{F0: 255, F1: 2147483647, F2: 538}
                        v80 := GenS2{F0: 370, F1: 288}
                        _ = v79
                        _ = v80
                    } else {
                        if ((v1 | (v1 / ((255) | 1))) > v1) {
                            v2 = v2
                        }
                    }
                    v62 += -503
                    var v81 int64 = 642
                    _ = v81
                }
                delete(v4, "a")
                _ = v78
            }
        }
    default:
        _ = len(v3)
    }
    switch ("foo bar" + ("" + "foo bar")) {
    case "a":
        if false {
            for fi82 := 0; fi82 < 6; fi82++ {
                if ((!(!(true && false))) && true) {
                    v62 += int(v1)
                } else {
                    v62 = (v57 << ((196) & 63))
                    v83 := GenS4{F0: -1, F1: true, F2: "foo bar"}
                    _ = v83
                }
                _ = fi82
            }
        } else {
            var v84 int32 = -1
            _ = v84
        }
    default:
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v85 int = v62
            _ = v85
        case 32768:
            for fi86 := 0; fi86 < 14; fi86++ {
                v62 -= int(v1)
                v2 -= v2
                v87 := GenS8{F0: 1, F1: -1}
                _ = fi86
                _ = v87
            }
        case -241:
            switch v62 {
            case 127:
                v2 = v2
            case -641:
                var v88 int8 = 104
                _ = v88
            }
        default:
            var v89 float32 = 3.14159
            _ = v89
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v62
    _ = v65
}
