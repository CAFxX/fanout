package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int; F1 int64 }

func main() {
    var v1 float32 = (3.14159 + float32(((1.5 - ((1.5 / 1.5) * 1.5)) - (0.1 + (-2.25 - 0.1)))))
    v2 := []uint32{42}
    v3 := []int{}
    v4 := make(map[int]int)
    v1 = (v1 * v1)
    switch 42 {
    case 125:
        fmt.Println(v1)
    default:
        var v5 bool = false
        _ = v5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint32 = 0
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int32 = -1
            var v10 int32 = (-(v9))
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: 100}
        _ = fi6
        _ = v11
        _ = v7
    }
    if false {
        if false {
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 int16 = -1
                _ = fi12
                _ = v13
            }
            var v14 uint = 381
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS2{F0: false}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
    } else {
        v2 = append(v2, 2147483647)
        if true {
            v17 := GenS1{F0: 127}
            _ = v17
        }
        if (62 <= 100) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if false {
                switch "a" {
                case "a":
                    v2 = append(v2, 42)
                case "\n\t":
                    v4[414] = -1
                }
                fmt.Println(v1)
                if ((v1 * 1.5) != v1) {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        fmt.Println(v1)
                        _ = fi18
                    }
                    var v19 int64 = -377
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 float64 = 3.14159
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                    var v23 uint8 = 215
                    _ = v23
                }
            } else {
                var v24 string = (("" + "\n\t") + ("" + "foo bar"))
                _ = v24
            }
            if (false || false) {
                v4[4294967295] = -1
            }
        } else {
            _ = len(v4)
            if true {
                var v25 int32 = -1
                _ = v25
            } else {
                if false {
                    switch "hello" {
                    case "":
                        if len(v2) > 0 {
                            v2[0] = 767
                        }
                    case "\n\t":
                        for fi26 := 0; fi26 < 4; fi26++ {
                            v1 = (float32(3.14159) - v1)
                            for fi27 := 0; fi27 < 3; fi27++ {
                                if true {
                                    var v28 uint = 973
                                    v29 := GenS1{F0: 1}
                                    _ = v28
                                    _ = v29
                                }
                                v30 := GenS2{F0: (42 != 2147483647)}
                                _ = fi27
                                _ = v30
                            }
                            var v31 string = "a"
                            _ = fi26
                            _ = v31
                        }
                    }
                    for fi32 := 0; fi32 < 14; fi32++ {
                        var v33 int64 = 610
                        var v34 uint = 232
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    delete(v4, 0)
                } else {
                    if (18446744073709551615 != 0) {
                        var v35 bool = ((false || (((("" + (("\n\t" + "foo bar") + "a")) + "hello") >= (("hello" + ("hello" + ("a" + ""))) + ("hello" + "a"))) && true)) && false)
                        if (0.0 == (0.1 + (-2.25 * 1.5))) {
                            fmt.Println(v35, v1)
                            switch 340 {
                            case 2147483647:
                                switch -9223372036854775808 {
                                case 1:
                                    switch ("hello" + "\n\t") {
                                    case "foo bar":
                                        var v36 bool = true
                                        _ = v36
                                    case "":
                                        fmt.Println(v35)
                                    case "hello":
                                        v35 = (!(true && v35))
                                    }
                                default:
                                    if v35 {
                                        for fi37 := 0; fi37 < 3; fi37++ {
                                            v3 = append(v3, fi37)
                                            for fi38 := 0; fi38 < 8; fi38++ {
                                                fmt.Println(fi38)
                                                if (!false) {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if v35 {
                                                            v1 = v1
                                                            var v39 uint8 = 155
                                                            v40 := GenS1{F0: -128}
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                        var v41 uint16 = 2
                                                        _ = v41
                                                    case "a":
                                                        if false {
                                                            if v35 {
                                                                if false {
                                                                    v35 = (v35 || (-2.25 < 1e10))
                                                                    var v42 int8 = -115
                                                                    fmt.Println(v35, v1)
                                                                    _ = v42
                                                                } else {
                                                                    v43 := GenS1{F0: 0}
                                                                    v35 = ((!v35) && true)
                                                                    switch ("foo bar" + (("a" + ((("a" + "foo bar") + ("hello" + "\n\t")) + "foo bar")) + "")) {
                                                                    case "":
                                                                        v44 := GenS2{F0: (!false)}
                                                                        _ = v44
                                                                    case "foo bar":
                                                                        if v35 {
                                                                            v45 := GenS3{F0: 42, F1: ("" + "")}
                                                                            fmt.Println(v35)
                                                                            _ = v45
                                                                        } else {
                                                                            _ = len(v3)
                                                                            if true {
                                                                                v46 := v4[(fi38 / (((^(fi37))) | 1))]
                                                                                v35 = ((fi37 << ((444) & 63)) <= (-(fi37)))
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                var v47 uint16 = 42
                                                                                v48 := GenS2{F0: (2 >= 600)}
                                                                                var v49 uint32 = 2147483648
                                                                                _ = v46
                                                                                _ = v47
                                                                                _ = v48
                                                                                _ = v49
                                                                            }
                                                                            var v50 int16 = -1
                                                                            _ = v50
                                                                        }
                                                                    case "a":
                                                                        v35 = (("" < "") && (!(11 > 596)))
                                                                    default:
                                                                        fmt.Println(v35)
                                                                    }
                                                                    _ = v43
                                                                }
                                                                if ((!v35) || v35) {
                                                                    var v51 uint = 1
                                                                    var v52 int16 = 1
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 4294967295
                                                                    }
                                                                    var v53 int8 = 1
                                                                    _ = v53
                                                                }
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v1 = v1
                                                                case "foo bar":
                                                                    var v54 uint64 = 1
                                                                    _ = v54
                                                                }
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v55 float64 = (((1.5 * (1e10 + ((float64(v1) / (((0.0 * (float64(v1) * -2.25)) + float64(v1)) + 1e10)) + (((0.0 + 1.5) - -2.25) * float64(float32((3.14159 / ((1.5 - (3.14159 / 0.1)) / 1.5)))))))) + 1e10) + 3.14159)
                                                        _ = v55
                                                    default:
                                                        v56 := GenS4{F0: 590}
                                                        _ = v56
                                                    }
                                                } else {
                                                    v57 := GenS5{F0: 65535}
                                                    _ = v57
                                                }
                                                v35 = ((1.5 / v1) >= (v1 - float32((float64(v1) * 0.1))))
                                                _ = fi38
                                            }
                                            var v58 string = "foo bar"
                                            _ = fi37
                                            _ = v58
                                        }
                                    }
                                }
                            case 176:
                                v59 := GenS1{F0: 127}
                                _ = v59
                            case 32768:
                                if v35 {
                                    v60 := GenS6{F0: 398, F1: "foo bar"}
                                    fmt.Println(v35)
                                    _ = v60
                                }
                            default:
                                v35 = false
                            }
                            v1 = ((v1 * (float32((1.5 / ((3.14159 * (1e10 * (0.0 / (-2.25 / 3.14159)))) - -2.25))) + 0.0)) - v1)
                        }
                        _ = v35
                    }
                    if (42 <= 2147483648) {
                        if (((("" + "\n\t") + "\n\t") + "hello") > ("" + "a")) {
                            if true {
                                var v61 string = ((("\n\t" + "a") + "") + ((("hello" + "foo bar") + (("" + "foo bar") + "\n\t")) + "\n\t"))
                                _ = v61
                            } else {
                                if (0 > 847) {
                                    v1 = 1e10
                                } else {
                                    switch "a" {
                                    case "a":
                                        var v62 int = 9223372036854775807
                                        _ = v62
                                    }
                                    var v63 uint32 = 690
                                    v64 := GenS7{F0: true, F1: (^(v63)), F2: 613}
                                    _ = v63
                                    _ = v64
                                }
                                v1 -= float32((1e10 / 3.14159))
                            }
                        }
                        for fi65 := 0; fi65 < 12; fi65++ {
                            for fi66 := 0; fi66 < 7; fi66++ {
                                v1 = 0.1
                                _ = fi66
                            }
                            var v67 uint64 = 793
                            _ = fi65
                            _ = v67
                        }
                        v1 = v1
                    } else {
                        var v68 float64 = 3.14159
                        _ = v68
                    }
                }
                delete(v4, 4294967295)
            }
        }
    }
    var v69 bool = true
    v70 := GenS7{F0: (v69 && ("\n\t" == ("foo bar" + "a"))), F1: 1, F2: 4294967295}
    switch "foo bar" {
    case "":
        var v71 int8 = 118
        _ = v71
    }
    v72 := v4[-2147483648]
    v72 -= v72
    switch 32768 {
    case 127:
        v1 = 1.5
    case 0:
        if v69 {
            switch v72 {
            case 328:
                var v73 int16 = 0
                _ = v73
            case 42:
                fmt.Println(v72, v69, v69)
            default:
                fmt.Println(v69, v1, v69)
            }
            for fi74 := 0; fi74 < 17; fi74++ {
                fmt.Println(v72)
                _ = len(v4)
                _ = fi74
            }
        }
    case 404:
        if (2 <= 0) {
            var v75 bool = (!(-94 > -1))
            _ = v75
        } else {
            switch "" {
            case "foo bar":
                for fi76 := 0; fi76 < 19; fi76++ {
                    for fi77 := 0; fi77 < 14; fi77++ {
                        switch v72 {
                        case 518:
                            var v78 float32 = v1
                            _ = v78
                        case -1:
                            v79 := GenS8{F0: 52}
                            _ = v79
                        }
                        var v80 int16 = 456
                        switch "a" {
                        case "hello":
                            v2 = append(v2, 4294967295)
                        }
                        _ = fi77
                        _ = v80
                    }
                    _ = fi76
                }
            default:
                var v81 float64 = 1.5
                _ = v81
            }
            if v69 {
                v69 = (1.5 == 0.0)
                fmt.Println(v1, v72)
                if (false || (9223372036854775807 < 0)) {
                    var v82 uint = 18446744073709551615
                    v72 += -1
                    _ = v82
                }
            }
            v1 = (1e10 / (1e10 - v1))
        }
    default:
        v83 := v4[0]
        _ = v83
    }
    for fi84 := 0; fi84 < 19; fi84++ {
        if (881 >= -2147483648) {
            var v85 int8 = 42
            _ = v85
        } else {
            v86 := GenS9{F0: 2}
            v69 = v69
            for fi87 := 0; fi87 < 7; fi87++ {
                var v88 uint8 = 1
                v89 := GenS10{F0: fi84, F1: 0}
                fmt.Println(v69, v1, v1)
                var v90 int = v72
                var v91 uint = 238
                _ = fi87
                _ = v88
                _ = v89
                _ = v90
                _ = v91
            }
            _ = v86
        }
        _ = fi84
    }
    if (!v69) {
        delete(v4, (v72 & 1))
        switch "a" {
        case "a":
            v69 = v69
        case "":
            if len(v3) > 0 {
                v3[0] = (v72 | v72)
            }
        }
        _ = len(v3)
    }
    var v92 int32 = -766
    v72 = 32
    v1 = 1.5
    var v93 uint32 = 35
    if false {
        if (v69 || false) {
            for fi94 := 0; fi94 < 7; fi94++ {
                v93 -= v93
                v92 = 2147483647
                var v95 bool = v69
                _ = fi94
                _ = v95
            }
            switch (v72 - (v72 & v72)) {
            case -796:
                var v96 uint = 4294967295
                _ = v96
            case 637:
                var v97 float64 = (0.0 * 3.14159)
                _ = v97
            case -618:
                if ((^(v92)) <= (v92 / ((v92) | 1))) {
                    v98 := GenS4{F0: 9223372036854775807}
                    if len(v3) > 0 {
                        v3[0] = v72
                    }
                    _ = v98
                } else {
                    for fi99 := 0; fi99 < 3; fi99++ {
                        v92 = int32(v93)
                        _ = fi99
                    }
                    if v69 {
                        var v100 uint = 255
                        _ = v100
                    }
                }
            default:
                var v101 int32 = -110
                _ = v101
            }
            v1 = v1
        }
    }
    if v69 {
        delete(v4, (v72 / (((v72 | 127)) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v69:", v69)
    fmt.Println("v72:", v72)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
    _ = v72
    _ = v92
    _ = v93
}
