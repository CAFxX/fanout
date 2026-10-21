package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool; F1 string; F2 int8 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 string; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint32 = 524
        var v9 float64 = 1.5
        v9 -= (v9 * v9)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: 21}
    var v11 uint = 65535
    switch "a" {
    case "hello":
        if false {
            v12 := GenS2{F0: (false || (!false)), F1: "hello", F2: 42}
            _ = v12
        } else {
            v13 := []int{}
            var v14 int = -9223372036854775808
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int16 = -1
        _ = fi15
        _ = v16
    }
    _ = v10
    _ = v11
    return 2
}

func main() {
    var v1 string = "hello"
    var v2 string = ""
    var v3 float32 = 0.0
    v4 := make(map[uint32]bool)
    var v5 uint32 = 42
    v6 := make(map[uint32]int)
    v17 := GenS1{F0: 0, F1: 4294967295}
    switch 9223372036854775807 {
    case 4294967295:
        switch 0 {
        case -418:
            var v18 int16 = 32767
            _ = v18
        case 4294967295:
            for fi19 := 0; fi19 < 6; fi19++ {
                if false {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        var v21 uint16 = 44
                        switch (v1 + (v1 + "")) {
                        case "":
                            v22 := GenS2{F0: true, F1: (v2 + v2), F2: 42}
                            _ = v22
                        case "hello":
                            if false {
                                v2 = v1
                                if ((true && false) && (2 >= 794)) {
                                    if false {
                                        v23 := GenS1{F0: 2, F1: 4294967295}
                                        _ = v23
                                    } else {
                                        switch (v1 + v2) {
                                        case "":
                                            v24 := []int{fi19, fi20, fi19, (fi19 >> ((426) & 63))}
                                            _ = v24
                                        case "a":
                                            switch fi19 {
                                            case 42:
                                                v1 = "hello"
                                            case 9223372036854775807:
                                                if (3.14159 < 0.1) {
                                                    v25 := GenS1{F0: uint8(v5), F1: 920}
                                                    var v26 float32 = (0.0 * v3)
                                                    v21 -= (v21 / ((v21) | 1))
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    v1 = "hello"
                                                    v5 = 690
                                                    fmt.Println(v21)
                                                }
                                            case 1:
                                                v27 := GenS3{F0: 42, F1: 680}
                                                _ = v27
                                            }
                                        case "foo bar":
                                            fmt.Println(v5)
                                        }
                                        var v28 int32 = 1
                                        v6[42] = 755
                                        _ = v28
                                    }
                                }
                                var v29 bool = (false && true)
                                _ = v29
                            }
                        }
                        switch v2 {
                        case "":
                            v21 -= (^(v21))
                        case "hello":
                            var v30 bool = (false && false)
                            _ = v30
                        case "foo bar":
                            v3 += 1.5
                        default:
                            var v31 uint = 2
                            if true {
                                switch (fi20 >> ((v31) & 63)) {
                                case 828:
                                    v32 := GenS4{F0: v1}
                                    _ = v32
                                case 530:
                                    if (!(573 >= 481)) {
                                        var v33 uint8 = 0
                                        v34 := GenS2{F0: false, F1: ("foo bar" + ((v1 + (("hello" + v2) + (v2 + "a"))) + v1)), F2: -37}
                                        v31 -= v31
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            }
                            fmt.Println(v21, fi19)
                            _ = v31
                        }
                        if (false && true) {
                            v1 = "foo bar"
                            v1 = v2
                            if ((1.5 - 1.5) == ((-2.25 - ((3.14159 - 0.1) + ((float64((1.5 / 0.1)) * (-2.25 + 1e10)) + 1.5))) - 0.0)) {
                                if (977 <= int64(fi19)) {
                                    if (float64(3.14159) != 0.0) {
                                        v35 := GenS5{F0: 448}
                                        fmt.Println(fi19, v3)
                                        var v36 uint = 249
                                        _ = v35
                                        _ = v36
                                    }
                                    delete(v4, (v5 << ((65535) & 63)))
                                    var v37 string = ((v2 + v1) + "")
                                    _ = v37
                                } else {
                                    switch -1 {
                                    case 2147483647:
                                        v3 = v3
                                    case -874:
                                        if false {
                                            if (!(uint64(v5) != genFn0())) {
                                                var v38 int8 = -68
                                                v39 := GenS6{F0: (fi20 & (fi20 & 0)), F1: (("\n\t" + ((("a" + v1) + v2) + v1)) + ("foo bar" + "")), F2: 127}
                                                _ = v38
                                                _ = v39
                                            } else {
                                                switch v2 {
                                                case "a":
                                                    v1 = ("" + v1)
                                                case "foo bar":
                                                    v40 := GenS2{F0: ((false && (!(!("foo bar" >= v1)))) || true), F1: (v2 + ""), F2: 0}
                                                    _ = v40
                                                case "":
                                                    var v41 float64 = 1.5
                                                    _ = v41
                                                }
                                                v1 = (("hello" + "\n\t") + ("hello" + ("\n\t" + ((v2 + "\n\t") + v2))))
                                                var v42 uint64 = 212
                                                _ = v42
                                            }
                                            v1 = (((((v2 + v2) + (("hello" + v1) + v2)) + v1) + "hello") + "hello")
                                        }
                                    case 351:
                                        v43 := GenS5{F0: 450}
                                        _ = v43
                                    }
                                    v3 = v3
                                    v21 = v21
                                    if (true || false) {
                                        var v44 int8 = 100
                                        if (uint(v5) <= 187) {
                                            v45 := GenS5{F0: 1}
                                            v46 := GenS1{F0: 1, F1: 861}
                                            var v47 float64 = 1e10
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        } else {
                                            v1 = "a"
                                            var v48 string = (v1 + ((v2 + v2) + v2))
                                            v2 = ("" + (v1 + ("" + (v2 + v1))))
                                            v49 := GenS5{F0: 2147483647}
                                            v50 := GenS4{F0: v1}
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                        }
                                        _ = v44
                                    } else {
                                        var v51 string = ("a" + v2)
                                        _ = v51
                                    }
                                    switch v2 {
                                    case "":
                                        v2 = ((((((v1 + ("foo bar" + v1)) + "") + "foo bar") + v2) + "hello") + "\n\t")
                                    case "foo bar":
                                        v52 := GenS3{F0: -128, F1: 255}
                                        _ = v52
                                    case "\n\t":
                                        var v53 int32 = 168
                                        _ = v53
                                    default:
                                        v4[(^(v5))] = (v5 == (v5 / (((v5 | (v5 << ((255) & 63)))) | 1)))
                                    }
                                }
                                v1 = ((v2 + (v1 + "")) + (v1 + "foo bar"))
                                v54 := GenS1{F0: 105, F1: 617}
                                var v55 bool = (1000 >= 1)
                                _ = v54
                                _ = v55
                            } else {
                                var v56 int16 = 1
                                v57 := GenS7{F0: -2147483648, F1: v21}
                                v58 := GenS1{F0: 255, F1: 118}
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                        }
                        switch fi19 {
                        case -77:
                            if (2147483648 >= (v5 | (v5 % ((432) | 1)))) {
                                switch ("foo bar" + ((v1 + "a") + v1)) {
                                case "\n\t":
                                    switch fi19 {
                                    case 42:
                                        v21 = (v21 ^ 472)
                                    case 9223372036854775807:
                                        v3 += ((v3 - 0.0) - 1e10)
                                    case 0:
                                        v5 += (v5 & 4294967295)
                                    }
                                case "":
                                    if (!(false || false)) {
                                        v59 := GenS8{F0: 1}
                                        var v60 uint64 = 995
                                        v5 = v5
                                        _ = v59
                                        _ = v60
                                    }
                                    switch "\n\t" {
                                    case "":
                                        v61 := GenS9{F0: false}
                                        _ = v61
                                    default:
                                        var v62 uint16 = v21
                                        _ = v62
                                    }
                                default:
                                    v3 -= v3
                                }
                            } else {
                                v21 = uint16(v5)
                                fmt.Println(v5, v21)
                                v3 = (v3 * v3)
                                switch (fi20 << ((348) & 63)) {
                                case 42:
                                    v6[(v5 - 42)] = fi20
                                case 306:
                                    v63 := GenS3{F0: -27, F1: (v21 - (v21 % (((v21 / ((803) | 1))) | 1)))}
                                    _ = v63
                                }
                            }
                        default:
                            var v64 uint = 2
                            _ = v64
                        }
                        v5 -= (v5 % (((v5 ^ (v5 >> ((900) & 63)))) | 1))
                        _ = fi20
                        _ = v21
                    }
                    if (-326 >= -2147483648) {
                        v65 := GenS8{F0: -1}
                        _ = v65
                    } else {
                        var v66 int64 = 42
                        for fi67 := 0; fi67 < 18; fi67++ {
                            switch (("\n\t" + ((v2 + v1) + "hello")) + ("\n\t" + (v1 + v2))) {
                            case "\n\t":
                                v68 := GenS4{F0: v1}
                                _ = v68
                            case "a":
                                var v69 int16 = 385
                                _ = v69
                            }
                            v70 := GenS5{F0: 42}
                            _ = fi67
                            _ = v70
                        }
                        _ = v66
                    }
                    v3 = 1e10
                } else {
                    for fi71 := 0; fi71 < 4; fi71++ {
                        var v72 int64 = 65
                        _ = fi71
                        _ = v72
                    }
                    for fi73 := 0; fi73 < 19; fi73++ {
                        v74 := GenS6{F0: 128, F1: (v1 + "hello"), F2: 2}
                        switch (-(fi73)) {
                        case -9223372036854775808:
                            v75 := []int{}
                            _ = v75
                        default:
                            v2 = "hello"
                        }
                        _ = fi73
                        _ = v74
                    }
                    v1 = v1
                }
                v2 = "\n\t"
                _ = fi19
            }
        case 604:
            var v76 int32 = -2147483648
            _ = v76
        }
    case -349:
        fmt.Println(v2)
    case -1:
        for fi77 := 0; fi77 < 1; fi77++ {
            var v78 uint64 = 2
            _ = fi77
            _ = v78
        }
    default:
        v2 = (("hello" + v1) + (v2 + ((v1 + v2) + "")))
    }
    _ = len(v6)
    if (2147483647 == 739) {
        for fi79 := 0; fi79 < 1; fi79++ {
            for fi80 := 0; fi80 < 1; fi80++ {
                v3 = v3
                _ = len(v4)
                v81 := []int{(fi79 & 4294967295)}
                _ = fi80
                _ = v81
            }
            v1 = "hello"
            v3 = ((v3 + (v3 * ((1.5 / v3) / ((v3 + (v3 * v3)) - v3)))) / v3)
            _ = fi79
        }
    }
    fmt.Println(v1, v3)
    v2 = v1
    v3 = float32((0.0 * (1.5 * -2.25)))
    var v82 int16 = 1000
    v5 -= (v5 | v5)
    var v83 string = ("a" + (((((v1 + v2) + (v2 + "")) + ("\n\t" + v2)) + (v1 + v2)) + v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v83
}
