package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 uint32; F1 bool; F2 int16 }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 int32 }
type GenS10 struct { F0 int32; F1 uint16 }
type GenS11 struct { F0 int8; F1 bool }
type GenS12 struct { F0 int; F1 uint16 }
func genFn0(p0 uint32) uint16 {
    p0 = (p0 & 579)
    var v6 uint16 = 0
    _ = p0
    _ = v6
    return uint16(p0)
}
func genFn1() string {
    for fi7 := 0; fi7 < 19; fi7++ {
        if (!(-485 >= 139)) {
            var v8 string = ("\n\t" + "hello")
            _ = v8
        } else {
            switch fi7 {
            case 9223372036854775807:
                switch "" {
                case "hello":
                    v9 := []int{fi7, fi7, (fi7 % ((-1) | 1)), 4294967295}
                    _ = v9
                case "":
                    if false {
                        for fi10 := 0; fi10 < 6; fi10++ {
                            var v11 int = (fi7 >> ((255) & 63))
                            if true {
                                if (false && false) {
                                    var v12 float64 = ((-2.25 * 1.5) * 1e10)
                                    _ = v12
                                } else {
                                    v13 := GenS1{F0: 1, F1: false}
                                    _ = v13
                                }
                            } else {
                                var v14 float32 = (-2.25 * ((0.0 / (1e10 + float32(3.14159))) * ((1.5 - float32((1.5 + -2.25))) + (3.14159 / 1e10))))
                                _ = v14
                            }
                            var v15 int64 = 9223372036854775807
                            switch (v11 ^ v11) {
                            case -478:
                                v16 := GenS2{F0: 42}
                                _ = v16
                            default:
                                v11 = 127
                            }
                            _ = fi10
                            _ = v11
                            _ = v15
                        }
                        v17 := GenS3{F0: 0, F1: 1}
                        _ = v17
                    }
                case "\n\t":
                    v18 := GenS3{F0: 65535, F1: 42}
                    _ = v18
                default:
                    if false {
                        var v19 uint32 = 4294967295
                        for fi20 := 0; fi20 < 16; fi20++ {
                            if (true || true) {
                                if (!(!true)) {
                                    if false {
                                        var v21 int16 = -788
                                        fmt.Println(fi7, v19, v19)
                                        if false {
                                            fmt.Println(fi20, v21, v19)
                                            fmt.Println(v21, v19, v19)
                                            v22 := []int{797, fi7}
                                            _ = v22
                                        }
                                        _ = v21
                                    }
                                    if true {
                                        fmt.Println(fi20)
                                        var v23 float64 = 1e10
                                        _ = v23
                                    } else {
                                        var v24 int8 = 127
                                        _ = v24
                                    }
                                    var v25 int32 = 1
                                    _ = v25
                                }
                            }
                            v26 := []int{fi7, (fi7 & fi20)}
                            v19 += 2147483648
                            _ = fi20
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 14; fi27++ {
                            var v28 uint32 = (v19 % (((v19 << ((1) & 63))) | 1))
                            _ = fi27
                            _ = v28
                        }
                        _ = v19
                    }
                }
            case 705:
                for fi29 := 0; fi29 < 17; fi29++ {
                    if true {
                        v30 := GenS4{F0: 129}
                        var v31 int16 = 0
                        _ = v30
                        _ = v31
                    } else {
                        fmt.Println(fi7)
                        v32 := []int{(fi7 >> ((550) & 63)), 32767, -632}
                        fmt.Println(fi29)
                        _ = v32
                    }
                    var v33 uint64 = 861
                    _ = fi29
                    _ = v33
                }
            case 128:
                if false {
                    fmt.Println(fi7)
                }
            }
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 int32 = -1
                v36 := GenS2{F0: (v35 % ((2147483647) | 1))}
                _ = fi34
                _ = v35
                _ = v36
            }
        }
        _ = fi7
    }
    var v37 string = ""
    if false {
        for fi38 := 0; fi38 < 2; fi38++ {
            v39 := GenS3{F0: 2, F1: 1}
            _ = fi38
            _ = v39
        }
    } else {
        v37 = v37
        for fi40 := 0; fi40 < 14; fi40++ {
            var v41 bool = false
            v42 := []int{4294967295, 312}
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    switch ("a" + v37) {
    case "a":
        v43 := GenS5{F0: 4294967295, F1: -80}
        _ = v43
    }
    v37 = "\n\t"
    _ = v37
    return (((v37 + (("foo bar" + v37) + v37)) + v37) + ("foo bar" + (v37 + ("a" + "a"))))
}

func main() {
    var v1 int16 = 362
    var v2 uint8 = 255
    var v3 int32 = int32(v2)
    var v4 bool = false
    var v5 uint32 = 2147483648
    v44 := []int{9223372036854775807, -934, 2147483647, 42}
    switch ("a" + ("hello" + "")) {
    case "a":
        fmt.Println(v3, v5)
    case "":
        v45 := GenS6{F0: v5, F1: v4, F2: (v1 & (v1 & 22))}
        _ = v45
    }
    for fi46 := 0; fi46 < 11; fi46++ {
        switch "foo bar" {
        case "hello":
            if true {
                fmt.Println(fi46, v3)
                v47 := GenS7{F0: (v5 >> ((0) & 63)), F1: 62}
                _ = v47
            } else {
                v5 -= (v5 ^ v5)
            }
        case "a":
            var v48 uint16 = 240
            _ = v48
        case "":
            var v49 int16 = v1
            _ = v49
        }
        _ = fi46
    }
    v1 -= (v1 - v1)
    var v50 uint8 = v2
    v44 = append(v44, -9223372036854775808)
    if ((v5 / (((v5 & v5)) | 1)) > uint32(v2)) {
        for fi51 := 0; fi51 < 17; fi51++ {
            if len(v44) > 0 {
                v44[0] = 32768
            }
            v52 := GenS8{F0: ""}
            _ = fi51
            _ = v52
        }
        v2 = (v50 - (v2 ^ 42))
    }
    if ((120 <= (v2 ^ 0)) || ((v5 <= 65535) || ((v50 & v50) != 2))) {
        v44 = append(v44, 32768)
        if ((-9223372036854775808 < 320) || ((0.0 / 3.14159) == 1.5)) {
            if len(v44) > 0 {
                v44[0] = -1
            }
            switch "a" {
            case "a":
                v50 -= (v50 / ((v2) | 1))
            case "hello":
                switch -89 {
                case -994:
                    switch -9223372036854775808 {
                    case -2147483648:
                        if len(v44) > 0 {
                            v44[0] = -9223372036854775808
                        }
                    case 4294967295:
                        switch (("a" + "hello") + "\n\t") {
                        case "\n\t":
                            var v53 uint = 4294967295
                            _ = v53
                        case "foo bar":
                            v1 = (v1 | 1000)
                        case "":
                            if (uint64(v2) <= 2) {
                                if true {
                                    switch 1 {
                                    case -570:
                                        v50 = v50
                                    }
                                    if ((2 == 1) || (42 > 0)) {
                                        v54 := GenS9{F0: 255, F1: 153}
                                        switch -1 {
                                        case 32768:
                                            v5 += (v5 + 548)
                                        case 32767:
                                            v55 := GenS10{F0: v3, F1: 255}
                                            _ = v55
                                        }
                                        v56 := GenS3{F0: uint16(v50), F1: 32767}
                                        _ = v54
                                        _ = v56
                                    } else {
                                        var v57 int16 = (^(v1))
                                        v57 = (^(v57))
                                        for fi58 := 0; fi58 < 3; fi58++ {
                                            switch fi58 {
                                            case 128:
                                                for fi59 := 0; fi59 < 19; fi59++ {
                                                    var v60 uint32 = v5
                                                    fmt.Println(v4)
                                                    _ = fi59
                                                    _ = v60
                                                }
                                            }
                                            var v61 uint = 255
                                            _ = fi58
                                            _ = v61
                                        }
                                        _ = v57
                                    }
                                    v3 = v3
                                }
                            }
                        default:
                            switch -1 {
                            case -2147483648:
                                v62 := GenS2{F0: int32(v1)}
                                _ = v62
                            default:
                                for fi63 := 0; fi63 < 2; fi63++ {
                                    if len(v44) > 0 {
                                        v44[0] = 255
                                    }
                                    _ = fi63
                                }
                            }
                        }
                    case 9223372036854775807:
                        for fi64 := 0; fi64 < 6; fi64++ {
                            switch ("hello" + "") {
                            case "hello":
                                if (!((!(v4 && v4)) || false)) {
                                    _ = len(v44)
                                    v2 = v2
                                    if (v4 || v4) {
                                        if false {
                                            _ = len(v44)
                                        }
                                        fmt.Println(v3, v50)
                                    }
                                }
                            case "foo bar":
                                v65 := GenS11{F0: 118, F1: true}
                                _ = v65
                            }
                            _ = fi64
                        }
                    }
                }
            case "\n\t":
                if (!v4) {
                    if (!true) {
                        fmt.Println(v1)
                        for fi66 := 0; fi66 < 15; fi66++ {
                            v67 := GenS1{F0: (v1 - v1), F1: (4294967295 != (v5 >> ((255) & 63)))}
                            v68 := GenS12{F0: (-(fi66)), F1: 106}
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                    }
                    for fi69 := 0; fi69 < 13; fi69++ {
                        if true {
                            var v70 int8 = 127
                            v70 -= v70
                            fmt.Println(v4, fi69, v3)
                            _ = v70
                        }
                        _ = fi69
                    }
                    v3 += (-(v3))
                } else {
                    if len(v44) >= 2 {
                        v44 = v44[1:len(v44)]
                    }
                    for fi71 := 0; fi71 < 17; fi71++ {
                        var v72 int32 = v3
                        _ = fi71
                        _ = v72
                    }
                    var v73 int16 = (v1 | -1)
                    _ = v73
                }
            default:
                fmt.Println(v2)
            }
        } else {
            if (!false) {
                switch "foo bar" {
                case "a":
                    fmt.Println(v5, v2, v2)
                }
                fmt.Println(v4, v5)
                var v74 int32 = (^(v3))
                _ = v74
            } else {
                var v75 int16 = v1
                _ = v75
            }
        }
        for fi76 := 0; fi76 < 8; fi76++ {
            if v4 {
                switch (-(fi76)) {
                case -270:
                    var v77 int64 = -842
                    _ = v77
                case -9223372036854775808:
                    var v78 string = ((("hello" + "foo bar") + "") + ("\n\t" + ((("hello" + (("a" + ("" + "")) + "a")) + "foo bar") + (("\n\t" + "") + "hello"))))
                    _ = v78
                default:
                    for fi79 := 0; fi79 < 18; fi79++ {
                        if false {
                            if ((v3 | -2147483648) < -1) {
                                if (v4 || (((459 > v1) || (718 < int16(v50))) || v4)) {
                                    switch (("\n\t" + ("a" + "foo bar")) + ("hello" + ("" + "hello"))) {
                                    case "hello":
                                        v80 := GenS2{F0: v3}
                                        _ = v80
                                    default:
                                        _ = len(v44)
                                    }
                                }
                            } else {
                                v50 -= (v2 / ((v50) | 1))
                                v81 := GenS7{F0: (v5 | (v5 - v5)), F1: 996}
                                _ = v81
                            }
                            var v82 int64 = -9223372036854775808
                            _ = v82
                        }
                        _ = fi79
                    }
                }
            }
            switch fi76 {
            case -9223372036854775808:
                var v83 bool = true
                _ = v83
            case -2147483648:
                if (0.0 < 0.0) {
                    if true {
                        fmt.Println(v4, v3, fi76)
                    }
                    var v84 bool = false
                    v85 := GenS6{F0: 42, F1: (1.5 <= 0.0), F2: (v1 | (v1 * 1))}
                    _ = v84
                    _ = v85
                }
            case 4294967295:
                if v4 {
                    var v86 int64 = 559
                    _ = v86
                } else {
                    if len(v44) >= 2 {
                        v44 = v44[1:len(v44)]
                    }
                    if (18446744073709551615 > 855) {
                        v4 = (v50 == v50)
                        v2 = 80
                        for fi87 := 0; fi87 < 20; fi87++ {
                            fmt.Println(v1, fi76)
                            var v88 uint64 = 18446744073709551615
                            var v89 int16 = (-(v1))
                            _ = fi87
                            _ = v88
                            _ = v89
                        }
                    } else {
                        if ((!false) && v4) {
                            if ((^(v2)) >= v50) {
                                v2 = (v2 - 42)
                            } else {
                                for fi90 := 0; fi90 < 9; fi90++ {
                                    v91 := GenS7{F0: 542, F1: 255}
                                    _ = fi90
                                    _ = v91
                                }
                            }
                        } else {
                            v3 += v3
                        }
                        var v92 uint64 = 4294967295
                        v1 = v1
                        _ = v92
                    }
                }
            default:
                for fi93 := 0; fi93 < 5; fi93++ {
                    v94 := GenS6{F0: 959, F1: (1e10 <= (float32(((1.5 / 1.0) / float64(1.5))) * 1e10)), F2: v1}
                    v95 := GenS10{F0: (v3 * (v3 | int32(v5))), F1: 225}
                    var v96 uint32 = 0
                    _ = fi93
                    _ = v94
                    _ = v95
                    _ = v96
                }
            }
            if v4 {
                fmt.Println(v50)
            }
            _ = fi76
        }
    } else {
        switch ("hello" + ("hello" + "\n\t")) {
        case "\n\t":
            v5 = (v5 | (v5 >> ((101) & 63)))
        case "":
            var v97 int8 = -128
            _ = v97
        }
        fmt.Println(v4, v4, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v50
}
