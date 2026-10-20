package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int64; F1 uint; F2 int32 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint32; F1 int; F2 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 int16; F2 int32 }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 int; F1 uint; F2 uint16 }
type GenS11 struct { F0 int8; F1 uint64; F2 int8 }
type GenS12 struct { F0 bool }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint16 = 42
            var v9 uint32 = 976
            var v10 int32 = -1
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        switch fi6 {
        case 32767:
            v11 := GenS1{F0: 65535, F1: 42}
            _ = v11
        case 42:
            v12 := []int{(-(fi6)), fi6, (fi6 >> ((65535) & 63))}
            _ = v12
        default:
            for fi13 := 0; fi13 < 20; fi13++ {
                fmt.Println(fi13)
                _ = fi13
            }
        }
        _ = fi6
    }
    v14 := []int{127, -9223372036854775808}
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint = 665
        if false {
            switch 42 {
            case 128:
                if false {
                    v16 += 1
                }
            }
            if false {
                v16 -= v16
                for fi17 := 0; fi17 < 1; fi17++ {
                    v16 = (v16 | (v16 - 2))
                    v16 = v16
                    _ = fi17
                }
            } else {
                if (2147483647 < 605) {
                    fmt.Println(v16)
                    var v18 float64 = 3.14159
                    var v19 uint8 = 2
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS1{F0: (v16 >> ((65535) & 63)), F1: int64(v16)}
                    _ = v20
                }
                if len(v14) > 0 {
                    v14[0] = (^(fi15))
                }
                fmt.Println(v16)
            }
        } else {
            var v21 float32 = float32(((1e10 + float64(1.5)) * 1e10))
            _ = v21
        }
        if ((-885 > 534) && true) {
            switch "" {
            case "foo bar":
                fmt.Println(fi15, v16)
            }
            for fi22 := 0; fi22 < 11; fi22++ {
                if true {
                    v16 -= (v16 - v16)
                    var v23 int64 = -9223372036854775808
                    _ = v23
                } else {
                    switch (fi22 - 4294967295) {
                    case 32767:
                        var v24 uint8 = 1
                        _ = v24
                    default:
                        fmt.Println(fi15)
                    }
                    v25 := GenS1{F0: (v16 + (v16 * v16)), F1: 581}
                    _ = v25
                }
                switch "a" {
                case "":
                    v26 := GenS1{F0: v16, F1: 208}
                    _ = v26
                case "foo bar":
                    if (!(1 >= (v16 << (((v16 * v16)) & 63)))) {
                        v27 := GenS2{F0: 676, F1: 4294967295, F2: (false && (9223372036854775807 > (-(fi15))))}
                        _ = v27
                    } else {
                        if (!true) {
                            var v28 uint8 = 128
                            v29 := GenS3{F0: -128, F1: -91}
                            _ = v28
                            _ = v29
                        } else {
                            var v30 uint32 = 65535
                            _ = v30
                        }
                    }
                }
                if ((fi15 > (fi22 * fi22)) || ((165 <= 238) || (-148 < 42))) {
                    v16 = v16
                    var v31 float32 = (1.5 + -2.25)
                    if len(v14) > 0 {
                        v14[0] = 127
                    }
                    _ = v31
                } else {
                    if true {
                        if (!true) {
                            v32 := GenS2{F0: 2147483647, F1: 42, F2: ((-2.25 + ((0.1 / -2.25) / float64(1.5))) < 1.5)}
                            if (2 == 42) {
                                var v33 uint64 = 18446744073709551615
                                v16 = (v16 + v16)
                                if true {
                                    v16 = v16
                                    var v34 bool = (0 > 42)
                                    fmt.Println(v34, v16, fi15)
                                    var v35 uint = 18446744073709551615
                                    _ = v34
                                    _ = v35
                                } else {
                                    var v36 uint = 725
                                    _ = v36
                                }
                                if false {
                                    switch fi15 {
                                    case 127:
                                        if (0.0 == 0.0) {
                                            var v37 string = ("foo bar" + "foo bar")
                                            fmt.Println(v16)
                                            _ = v37
                                        } else {
                                            switch (fi15 - fi15) {
                                            case 0:
                                                v33 = uint64(v16)
                                            case -2147483648:
                                                if len(v14) >= 2 {
                                                    v14 = v14[1:len(v14)]
                                                }
                                            case 460:
                                                fmt.Println(v33, v16, fi22)
                                            default:
                                                var v38 bool = true
                                                _ = v38
                                            }
                                            switch ("\n\t" + "") {
                                            case "hello":
                                                fmt.Println(fi22)
                                            case "":
                                                if len(v14) >= 2 {
                                                    v14 = v14[1:len(v14)]
                                                }
                                            case "foo bar":
                                                v39 := GenS4{F0: 723, F1: uint(v33), F2: -1}
                                                _ = v39
                                            }
                                        }
                                    default:
                                        var v40 uint = 338
                                        _ = v40
                                    }
                                    v41 := GenS3{F0: -114, F1: 1000}
                                    switch 42 {
                                    case -2147483648:
                                        var v42 uint16 = 65535
                                        var v43 int = (^(fi22))
                                        fmt.Println(v42)
                                        _ = v42
                                        _ = v43
                                    default:
                                        var v44 int8 = 1
                                        _ = v44
                                    }
                                    var v45 float64 = -2.25
                                    v45 = (((v45 - 1.5) / 1.0) * ((1.5 - -2.25) - float64((-2.25 - 0.1))))
                                    _ = v41
                                    _ = v45
                                } else {
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    switch 128 {
                                    case 0:
                                        v46 := GenS3{F0: 100, F1: 97}
                                        _ = v46
                                    case 42:
                                        v33 = v33
                                    case 4294967295:
                                        var v47 uint64 = (v33 << ((v16) & 63))
                                        _ = v47
                                    default:
                                        fmt.Println(fi22)
                                    }
                                }
                                _ = v33
                            } else {
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            }
                            _ = v32
                        }
                    }
                    v16 += (v16 >> ((368) & 63))
                }
                _ = fi22
            }
            for fi48 := 0; fi48 < 9; fi48++ {
                switch ("foo bar" + ("a" + "a")) {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        var v49 int64 = 0
                        _ = v49
                    case "hello":
                        v50 := GenS5{F0: true, F1: 2}
                        _ = v50
                    }
                case "":
                    v51 := GenS6{F0: 2, F1: -2147483648, F2: -169}
                    _ = v51
                case "hello":
                    v16 = (v16 << ((v16) & 63))
                }
                v16 += v16
                _ = len(v14)
                v52 := GenS3{F0: 127, F1: 999}
                v53 := GenS7{F0: true}
                v16 = v16
                _ = fi48
                _ = v52
                _ = v53
            }
        } else {
            fmt.Println(v16, fi15)
            var v54 int8 = 100
            _ = v54
        }
        _ = fi15
        _ = v16
    }
    var v55 bool = false
    _ = v14
    _ = v55
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{"a"}
    var v3 float32 = 3.14159
    var v4 uint = 0
    var v5 int32 = -2147483648
    v1 = v1
    var v56 int16 = -486
    if (0.1 > v3) {
        for fi57 := 0; fi57 < 20; fi57++ {
            for fi58 := 0; fi58 < 14; fi58++ {
                if false {
                    var v59 int16 = (v56 << (((v4 - (v4 * 855))) & 63))
                    fmt.Println(v5)
                    fmt.Println(fi57, v1, v5)
                    _ = v59
                }
                _ = fi58
            }
            _ = fi57
        }
        if (((4294967295 <= 18446744073709551615) && (18446744073709551615 != 699)) && false) {
            _ = len(v2)
            v3 -= -2.25
            if ((v1 & (v1 - -9223372036854775808)) < 42) {
                v3 -= v3
                var v60 int = -546
                _ = v60
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                var v61 int8 = 1
                _ = v61
            case "hello":
                switch 127 {
                case 507:
                    var v62 float32 = float32(0.1)
                    _ = v62
                case 127:
                    v63 := GenS1{F0: 18446744073709551615, F1: 42}
                    _ = v63
                case 4294967295:
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                }
            case "a":
                v64 := GenS8{F0: true, F1: v56, F2: 668}
                _ = v64
            default:
                var v65 uint64 = 919
                _ = v65
            }
            v66 := GenS9{F0: "foo bar", F1: 100}
            _ = v66
        }
    } else {
        if ((!true) && false) {
            for fi67 := 0; fi67 < 12; fi67++ {
                var v68 uint16 = 424
                _ = fi67
                _ = v68
            }
            for fi69 := 0; fi69 < 2; fi69++ {
                v56 += (v56 % (((-(v56))) | 1))
                _ = fi69
            }
        }
        v4 += v4
    }
    if false {
        v2 = append(v2, ("" + ""))
        switch 696 {
        case 932:
            fmt.Println(v5, v4, v1)
        default:
            switch 9223372036854775807 {
            case 500:
                if len(v2) > 0 {
                    v2[0] = ""
                }
            case 128:
                fmt.Println(v5)
            case -92:
                switch "hello" {
                case "":
                    var v70 float64 = float64(v3)
                    _ = v70
                case "foo bar":
                    var v71 string = ("a" + "a")
                    _ = v71
                case "hello":
                    v1 = (^(v1))
                }
            }
        }
        v72 := GenS10{F0: 128, F1: 2, F2: 32767}
        _ = v72
    }
    switch (("" + "hello") + "\n\t") {
    case "a":
        if true {
            var v73 int64 = v1
            v74 := GenS4{F0: 552, F1: v4, F2: (v5 * (v5 / (((v5 % (((v5 % (((v5 / (((^(v5))) | 1))) | 1))) | 1))) | 1)))}
            _ = v73
            _ = v74
        } else {
            fmt.Println(v56, v5)
            for fi75 := 0; fi75 < 20; fi75++ {
                var v76 int32 = (v5 % ((-2147483648) | 1))
                v56 += (v56 & (v56 + 32767))
                var v77 uint = 255
                _ = fi75
                _ = v76
                _ = v77
            }
            var v78 int32 = v5
            _ = v78
        }
    case "":
        var v79 bool = (false && true)
        _ = v79
    case "foo bar":
        v1 += (v1 * (^(v1)))
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    fmt.Println(v5, v56, v4)
    v5 += -1
    if false {
        for fi80 := 0; fi80 < 15; fi80++ {
            switch ("hello" + "hello") {
            case "foo bar":
                switch (fi80 << ((v4) & 63)) {
                case 0:
                    if false {
                        v81 := GenS11{F0: 42, F1: 674, F2: int8(fi80)}
                        for fi82 := 0; fi82 < 1; fi82++ {
                            v83 := GenS2{F0: (^(v5)), F1: genFn0(), F2: (!((!(-1 != 0)) || false))}
                            var v84 float64 = 0.1
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                        _ = v81
                    }
                case 1:
                    v1 += (v1 << ((v4) & 63))
                case -9223372036854775808:
                    if (!false) {
                        v4 -= (v4 ^ (v4 / ((v4) | 1)))
                    } else {
                        fmt.Println(v56, v3, fi80)
                        for fi85 := 0; fi85 < 16; fi85++ {
                            var v86 uint32 = 42
                            v86 = 197
                            if true {
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ("\n\t" + "foo bar")
                                    }
                                    v56 += (v56 * (-(v56)))
                                    v56 -= (^(v56))
                                    var v87 float32 = ((-2.25 - v3) + 1e10)
                                    _ = v87
                                } else {
                                    v88 := GenS5{F0: ((v4 + 4294967295) <= v4), F1: 255}
                                    v89 := GenS12{F0: false}
                                    _ = v88
                                    _ = v89
                                }
                                if false {
                                    fmt.Println(v1, v5, v4)
                                } else {
                                    var v90 uint = v4
                                    var v91 string = (("" + "hello") + ("foo bar" + (("foo bar" + "\n\t") + "a")))
                                    _ = v90
                                    _ = v91
                                }
                            } else {
                                if true {
                                    if true {
                                        v4 = v4
                                        fmt.Println(v1, v5)
                                    }
                                }
                            }
                            _ = fi85
                            _ = v86
                        }
                    }
                default:
                    _ = len(v2)
                }
            case "hello":
                switch "" {
                case "":
                    for fi92 := 0; fi92 < 19; fi92++ {
                        v1 += v1
                        if (!((false && false) || (v4 <= v4))) {
                            v5 -= (v5 * 0)
                            v2 = append(v2, ("\n\t" + ""))
                        }
                        _ = fi92
                    }
                case "hello":
                    for fi93 := 0; fi93 < 2; fi93++ {
                        v94 := GenS8{F0: false, F1: 32767, F2: (v5 & 42)}
                        _ = len(v2)
                        _ = fi93
                        _ = v94
                    }
                case "foo bar":
                    v95 := GenS11{F0: 42, F1: 2, F2: 100}
                    _ = v95
                default:
                    var v96 float32 = 0.0
                    _ = v96
                }
            }
            switch "\n\t" {
            case "hello":
                var v97 uint = (v4 % (((v4 ^ (v4 / ((65535) | 1)))) | 1))
                _ = v97
            }
            for fi98 := 0; fi98 < 19; fi98++ {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                var v99 int8 = 1
                v1 = (v1 & v1)
                _ = fi98
                _ = v99
            }
            _ = fi80
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
