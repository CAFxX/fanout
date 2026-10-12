package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS13 struct { F0 string; F1 int; F2 int16 }
func genFn0(p0 bool, p1 uint16, p2 int32) uint8 {
    if (true || true) {
        v6 := GenS1{F0: 2}
        v7 := GenS1{F0: 369}
        _ = v6
        _ = v7
    }
    switch int(p2) {
    case 32767:
        if (0 == uint32(p1)) {
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := []int{}
                _ = fi8
                _ = v9
            }
        } else {
            v10 := GenS2{F0: 42}
            _ = v10
        }
    default:
        fmt.Println(p2, p1, p2)
    }
    if (0 == 899) {
        p2 = (p2 ^ (p2 / ((p2) | 1)))
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := GenS2{F0: 2}
            fmt.Println(p2, p2, p2)
            var v13 int8 = 127
            _ = fi11
            _ = v12
            _ = v13
        }
        if (!p0) {
            if p0 {
                var v14 int32 = (p2 ^ (^(p2)))
                for fi15 := 0; fi15 < 16; fi15++ {
                    if (42 <= 255) {
                        var v16 string = "foo bar"
                        var v17 float32 = -2.25
                        p2 -= (v14 / ((p2) | 1))
                        _ = v16
                        _ = v17
                    }
                    var v18 uint32 = 0
                    p1 += p1
                    _ = fi15
                    _ = v18
                }
                p1 += (p1 % (((p1 / (((p1 / (((p1 / ((42) | 1))) | 1))) | 1))) | 1))
                _ = v14
            }
        }
    }
    switch (("foo bar" + ("" + "")) + ("foo bar" + (("\n\t" + "foo bar") + ((((("\n\t" + "foo bar") + ((("foo bar" + "a") + "foo bar") + "")) + "foo bar") + ("a" + ("hello" + ""))) + "hello")))) {
    case "\n\t":
        if (true && p0) {
            v19 := GenS2{F0: 255}
            _ = v19
        }
    case "foo bar":
        for fi20 := 0; fi20 < 8; fi20++ {
            fmt.Println(p2, p1)
            var v21 float32 = 1e10
            _ = fi20
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 134
}
func genFn1(p0 int8) uint64 {
    for fi22 := 0; fi22 < 20; fi22++ {
        p0 = p0
        p0 -= p0
        _ = fi22
    }
    if false {
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint16 = 42
            if (!(32767 == -632)) {
                v25 := []int{}
                _ = v25
            }
            if false {
                v26 := GenS2{F0: 0}
                _ = v26
            } else {
                v27 := []int{(fi23 * (fi23 << ((65535) & 63))), (fi23 | (fi23 ^ fi23)), (fi23 / ((fi23) | 1))}
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        v28 := GenS3{F0: 2}
        _ = v28
    }
    v29 := []int{684, -2147483648, -177, -808}
    _ = len(v29)
    _ = p0
    _ = v29
    return 18446744073709551615
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 1
    var v3 float32 = (0.0 * (float32(-2.25) + 0.0))
    v4 := []int{764, -1, 0}
    var v5 int = 0
    v1 -= v1
    v5 = 2147483647
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (v5 % (((v5 << ((453) & 63))) | 1)) {
    case 127:
        for fi30 := 0; fi30 < 4; fi30++ {
            v2 = 42
            v3 = (v3 * (v3 * (1e10 + (v3 / v3))))
            if (((0.0 + v3) <= (0.0 - (3.14159 + v3))) || false) {
                v2 = v2
            } else {
                if true {
                    if (-2.25 < 0.1) {
                        v31 := GenS4{F0: -32768, F1: 275}
                        _ = v31
                    } else {
                        if false {
                            var v32 string = "\n\t"
                            var v33 uint16 = 488
                            v32 = v32
                            _ = v32
                            _ = v33
                        }
                        if len(v4) > 0 {
                            v4[0] = v5
                        }
                    }
                    var v34 int32 = -2147483648
                    _ = v34
                } else {
                    if false {
                        if (!((v3 > v3) || true)) {
                            v35 := GenS4{F0: 32767, F1: 65535}
                            switch 9223372036854775807 {
                            case 0:
                                v1 += 1
                            }
                            _ = v35
                        }
                        v36 := GenS3{F0: 0}
                        _ = v36
                    }
                }
                v1 = 273
            }
            _ = fi30
        }
    case -696:
        var v37 int32 = -1
        _ = v37
    }
    v1 = 42
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v38 bool = (false && true)
        v39 := GenS3{F0: 255}
        _ = v38
        _ = v39
    }
    if true {
        if (0 < 227) {
            for fi40 := 0; fi40 < 18; fi40++ {
                var v41 int8 = 42
                var v42 uint = 1
                _ = len(v4)
                _ = fi40
                _ = v41
                _ = v42
            }
            fmt.Println(v2, v5)
        }
        var v43 float64 = 1e10
        switch 956 {
        case -17:
            if true {
                if false {
                    fmt.Println(v1, v43)
                } else {
                    v43 += v43
                    fmt.Println(v1, v43)
                }
                var v44 uint8 = 255
                _ = v44
            } else {
                v45 := GenS1{F0: 2}
                v5 = 128
                v3 += float32(v43)
                _ = v45
            }
        case 127:
            v46 := GenS4{F0: 42, F1: 4294967295}
            _ = v46
        }
        _ = v43
    }
    v4 = append(v4, (v5 | (v5 ^ (v5 / (((v5 ^ v5)) | 1)))))
    for fi47 := 0; fi47 < 20; fi47++ {
        if true {
            _ = len(v4)
        } else {
            fmt.Println(fi47, v1, v5)
        }
        v2 -= (v2 | (v2 % (((v2 * 159)) | 1)))
        for fi48 := 0; fi48 < 14; fi48++ {
            v2 = v2
            var v49 uint = 0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    switch ("foo bar" + "a") {
    case "":
        for fi50 := 0; fi50 < 19; fi50++ {
            if len(v4) > 0 {
                v4[0] = (v5 * (v5 - -2147483648))
            }
            v51 := GenS5{F0: 2, F1: 749, F2: -1}
            v52 := GenS6{F0: 42}
            _ = fi50
            _ = v51
            _ = v52
        }
    case "\n\t":
        var v53 int8 = 42
        _ = v53
    }
    if false {
        switch "hello" {
        case "a":
            if true {
                fmt.Println(v1, v3, v3)
                for fi54 := 0; fi54 < 2; fi54++ {
                    v4 = append(v4, 119)
                    if ((-128 < -128) && ("\n\t" > "foo bar")) {
                        fmt.Println(v2)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v55 := GenS2{F0: (v2 >> ((2) & 63))}
                        _ = v55
                    }
                    for fi56 := 0; fi56 < 20; fi56++ {
                        if true {
                            if (-1 == 505) {
                                switch (fi54 | (fi56 | v5)) {
                                case -2147483648:
                                    v1 = v1
                                case 798:
                                    switch (("hello" + "foo bar") + "") {
                                    case "hello":
                                        v5 = v5
                                    case "foo bar":
                                        var v57 float32 = v3
                                        _ = v57
                                    case "a":
                                        fmt.Println(v3)
                                    default:
                                        v3 = ((0.0 + v3) * (v3 * (float32(((3.14159 + 3.14159) - ((((1e10 / 1.0) - -2.25) / 1e10) / ((float64(v3) + float64((v3 / 1.0))) + (1e10 + float64((3.14159 - (float32(((0.1 * 0.1) - (1e10 - 0.0))) - float32(0.1))))))))) - ((v3 + float32(3.14159)) - (v3 / v3)))))
                                    }
                                case -212:
                                    var v58 int64 = 42
                                    _ = v58
                                }
                                var v59 int = v5
                                switch (v59 & fi56) {
                                case 128:
                                    if (false && (false || (!false))) {
                                        if (!(0.1 >= ((v3 - v3) - v3))) {
                                            if false {
                                                var v60 int32 = 0
                                                _ = v60
                                            }
                                            if true {
                                                if (2147483647 > 2) {
                                                    v5 = (-(fi54))
                                                    v61 := GenS7{F0: v2}
                                                    _ = v61
                                                }
                                            } else {
                                                v3 -= 3.14159
                                                switch ("" + ("" + ("hello" + ""))) {
                                                case "a":
                                                    v62 := GenS1{F0: 137}
                                                    _ = v62
                                                case "\n\t":
                                                    if (2147483648 <= 4294967295) {
                                                        fmt.Println(fi56, v5, v1)
                                                        v63 := GenS8{F0: 20}
                                                        switch fi54 {
                                                        case 673:
                                                            if true {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            } else {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    switch ("\n\t" + "foo bar") {
                                                                    case "hello":
                                                                        var v64 uint16 = 1
                                                                        _ = v64
                                                                    }
                                                                case "hello":
                                                                    v5 = (fi56 ^ -2147483648)
                                                                }
                                                                v4 = append(v4, (-(v59)))
                                                                var v65 float64 = (0.1 + (1.5 - 1e10))
                                                                _ = v65
                                                            }
                                                        case 0:
                                                            v2 -= (v2 / ((2) | 1))
                                                        default:
                                                            v59 = (fi54 | fi56)
                                                        }
                                                        _ = v63
                                                    } else {
                                                        var v66 int = 0
                                                        _ = v66
                                                    }
                                                default:
                                                    fmt.Println(v1)
                                                    _ = len(v4)
                                                    v67 := GenS9{F0: "", F1: 1}
                                                    _ = v67
                                                }
                                                v68 := GenS10{F0: (3.14159 <= 1e10)}
                                                _ = v68
                                            }
                                        }
                                        switch ("" + "") {
                                        case "\n\t":
                                            var v69 int64 = (v1 ^ -9223372036854775808)
                                            _ = v69
                                        default:
                                            if false {
                                                var v70 uint64 = 0
                                                if true {
                                                    var v71 uint32 = 2147483648
                                                    if (v70 != v70) {
                                                        var v72 string = "a"
                                                        v73 := GenS9{F0: (("\n\t" + ("\n\t" + "hello")) + (v72 + "")), F1: 1}
                                                        fmt.Println(v3, v71)
                                                        _ = v72
                                                        _ = v73
                                                    } else {
                                                        v2 = genFn0(false, 42, 2147483647)
                                                        var v74 uint8 = (v2 + (v2 ^ v2))
                                                        switch -328 {
                                                        case -2147483648:
                                                            v1 = -1
                                                        case 42:
                                                            v71 = 2147483648
                                                        case 918:
                                                            var v75 int = -740
                                                            _ = v75
                                                        }
                                                        _ = v74
                                                    }
                                                    var v76 float32 = 1e10
                                                    _ = v71
                                                    _ = v76
                                                }
                                                if true {
                                                    var v77 uint16 = 42
                                                    _ = v77
                                                } else {
                                                    switch "hello" {
                                                    case "hello":
                                                        var v78 int16 = 1
                                                        _ = v78
                                                    case "":
                                                        v2 = (v2 >> ((1) & 63))
                                                    case "foo bar":
                                                        _ = len(v4)
                                                    default:
                                                        fmt.Println(v5, v1)
                                                    }
                                                }
                                                var v79 bool = false
                                                var v80 int32 = 42
                                                _ = v70
                                                _ = v79
                                                _ = v80
                                            } else {
                                                var v81 int64 = (v1 & v1)
                                                var v82 uint64 = 18446744073709551615
                                                _ = v81
                                                _ = v82
                                            }
                                        }
                                    }
                                case -2147483648:
                                    v83 := GenS11{F0: 0}
                                    _ = v83
                                default:
                                    v3 -= v3
                                }
                                _ = v59
                            }
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = v5
                                }
                            }
                            var v84 int = 32768
                            _ = v84
                        }
                        switch 128 {
                        case 42:
                            _ = len(v4)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 127:
                            switch ("a" + (("hello" + "\n\t") + "\n\t")) {
                            case "\n\t":
                                v85 := GenS12{F0: 32767, F1: genFn0((!((3.14159 * 0.1) == float64(v3))), 3, -1), F2: 214}
                                _ = v85
                            default:
                                v2 = (v2 * v2)
                            }
                        case 0:
                            if (634 <= 0) {
                                v2 = (v2 * 0)
                                fmt.Println(v3)
                                var v86 float32 = (v3 + v3)
                                _ = v86
                            } else {
                                v87 := GenS10{F0: (v5 == (^(v5)))}
                                _ = v87
                            }
                        }
                        if (!false) {
                            switch "\n\t" {
                            case "":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                var v88 string = ("" + ("hello" + ""))
                                _ = v88
                            }
                        }
                        _ = fi56
                    }
                    _ = fi54
                }
            }
        case "\n\t":
            for fi89 := 0; fi89 < 14; fi89++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v4 = append(v4, -2147483648)
                _ = fi89
            }
        case "":
            var v90 string = "foo bar"
            _ = v90
        }
        var v91 uint8 = (v2 << ((255) & 63))
        if (65535 > 2) {
            v91 += genFn0((((true && (true || (("" + (("\n\t" + "hello") + ("" + "a"))) > "a"))) && (2147483648 != 42)) && true), 570, -476)
        } else {
            var v92 int = (^(v5))
            switch 32768 {
            case 9223372036854775807:
                switch 0 {
                case 9223372036854775807:
                    var v93 uint = 18446744073709551615
                    _ = v93
                case 610:
                    fmt.Println(v3, v5)
                default:
                    v94 := GenS12{F0: 1000, F1: v91, F2: 709}
                    _ = v94
                }
            default:
                fmt.Println(v2, v3)
            }
            _ = v92
        }
        _ = v91
    } else {
        _ = len(v4)
        if false {
            v3 += v3
            v95 := GenS13{F0: "", F1: (v5 % (((v5 * (-(v5)))) | 1)), F2: 42}
            _ = v95
        }
    }
    v3 -= (v3 / 1e10)
    fmt.Println(v2)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
