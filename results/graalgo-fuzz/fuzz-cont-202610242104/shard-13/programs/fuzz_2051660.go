package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 string; F2 int16 }
type GenS7 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int, p1 uint32, p2 int) bool {
    var v7 uint = 18446744073709551615
    var v8 uint32 = (^(p1))
    v7 += (v7 >> ((v7) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return true
}
func genFn1() string {
    var v9 int64 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        if ((!false) || (false && true)) {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 uint32 = 4294967295
                _ = fi11
                _ = v12
            }
            if ("a" >= "hello") {
                var v13 uint8 = 2
                var v14 int16 = 969
                var v15 uint = 42
                _ = v13
                _ = v14
                _ = v15
            } else {
                v16 := []int{-1, fi10, -933}
                _ = v16
            }
        } else {
            v9 = 42
            switch "" {
            case "":
                fmt.Println(v9, v9)
            case "hello":
                var v17 int32 = -715
                _ = v17
            case "\n\t":
                v18 := []int{fi10}
                _ = v18
            }
        }
        fmt.Println(v9, v9)
        switch 32767 {
        case 42:
            if (!(1.5 <= 1.5)) {
                fmt.Println(fi10)
                v19 := GenS1{F0: "hello", F1: 42}
                var v20 bool = ((32767 < 0) && false)
                _ = v19
                _ = v20
            }
        case 9223372036854775807:
            for fi21 := 0; fi21 < 6; fi21++ {
                switch fi10 {
                case -470:
                    var v22 float64 = float64(3.14159)
                    _ = v22
                case 42:
                    var v23 string = ("hello" + ("foo bar" + ("foo bar" + "")))
                    _ = v23
                default:
                    v24 := GenS1{F0: "a", F1: 1}
                    _ = v24
                }
                _ = fi21
            }
        case -759:
            fmt.Println(fi10, v9)
        }
        _ = fi10
    }
    _ = v9
    return "a"
}

func main() {
    var v1 int16 = 32767
    var v2 uint = 588
    v3 := []string{("\n\t" + "hello"), ("\n\t" + ("foo bar" + ("hello" + "hello"))), "a", "foo bar"}
    v4 := []uint32{0, 4294967295, 42, 484}
    v5 := []string{"foo bar", "", "\n\t"}
    v6 := []uint32{2}
    if (false && true) {
        var v25 float32 = ((0.1 - -2.25) * -2.25)
        var v26 uint64 = 290
        if (v2 > (^(v2))) {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v27 := GenS1{F0: ("\n\t" + ""), F1: 42}
            _ = v27
        }
        _ = v25
        _ = v26
    }
    var v28 bool = (0 >= 4294967295)
    var v29 int64 = 9223372036854775807
    var v30 uint16 = 0
    v31 := GenS2{F0: 2}
    if (!v28) {
        var v32 int64 = (v29 << (((v2 / ((2) | 1))) & 63))
        v4 = append(v4, 2)
        var v33 bool = v28
        _ = v32
        _ = v33
    } else {
        v30 = v30
        v30 += 125
        v28 = v28
    }
    var v34 uint32 = 2147483648
    if v28 {
        v35 := GenS3{F0: 42, F1: 65535, F2: (v29 ^ (v29 * v29))}
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 uint8 = 255
            for fi38 := 0; fi38 < 19; fi38++ {
                if v28 {
                    v29 += v29
                    switch (fi36 - fi38) {
                    case 42:
                        v34 += 2147483647
                    case 32768:
                        fmt.Println(v1, v29, v1)
                    default:
                        if (int32(fi36) >= 1) {
                            if v28 {
                                switch (^(fi36)) {
                                case 32768:
                                    fmt.Println(v34, v2, fi38)
                                case 566:
                                    var v39 float64 = 3.14159
                                    _ = v39
                                case -1:
                                    var v40 uint = v2
                                    _ = v40
                                }
                                var v41 int32 = 2147483647
                                var v42 uint = v2
                                _ = v41
                                _ = v42
                            } else {
                                v28 = true
                                var v43 uint = v2
                                if v28 {
                                    var v44 int8 = -128
                                    _ = v44
                                }
                                switch int(v37) {
                                case 964:
                                    v37 = v37
                                case 2147483647:
                                    var v45 uint32 = v34
                                    _ = v45
                                }
                                _ = v43
                            }
                        }
                    }
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v37, v29, fi38)
                    case "a":
                        v46 := GenS2{F0: 4294967295}
                        _ = v46
                    }
                    v47 := GenS1{F0: "\n\t", F1: 97}
                    var v48 int64 = (v29 | v29)
                    v49 := GenS2{F0: 231}
                    switch (fi38 & (-(fi36))) {
                    case 4294967295:
                        v29 = v29
                    }
                    _ = v47
                    _ = v48
                    _ = v49
                }
                v30 += 974
                if false {
                    v1 = int16(fi38)
                    var v50 uint64 = 612
                    if v28 {
                        if v28 {
                            switch "" {
                            case "foo bar":
                                v29 = v29
                            case "a":
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    var v51 bool = ((v34 % ((v34) | 1)) >= v34)
                                    _ = v51
                                case "foo bar":
                                    switch ("" + ("foo bar" + "")) {
                                    case "foo bar":
                                        switch ("a" + ("a" + "\n\t")) {
                                        case "a":
                                            var v52 int = int(v50)
                                            _ = v52
                                        case "hello":
                                            v50 = v50
                                        case "":
                                            var v53 int8 = 1
                                            _ = v53
                                        default:
                                            v1 = (v1 << ((804) & 63))
                                        }
                                    case "\n\t":
                                        v34 = uint32(v2)
                                    default:
                                        v37 += (v37 & uint8(v30))
                                    }
                                case "":
                                    switch (fi36 << (((v2 ^ (v2 << (((v2 ^ v2)) & 63)))) & 63)) {
                                    case 845:
                                        fmt.Println(v2, v29)
                                    case 688:
                                        v54 := GenS1{F0: "", F1: 0}
                                        _ = v54
                                    }
                                    var v55 uint = v2
                                    _ = v55
                                }
                            case "\n\t":
                                v56 := GenS4{F0: (v34 % (((v34 + v34)) | 1)), F1: (v29 % (((v29 % (((^(v29))) | 1))) | 1))}
                                _ = v56
                            default:
                                v50 -= v50
                            }
                            fmt.Println(v50, fi38, v30)
                            var v57 int = -574
                            _ = v57
                        } else {
                            if v28 {
                                switch 4294967295 {
                                case 127:
                                    switch (fi36 << (((v2 ^ v2)) & 63)) {
                                    case 2147483647:
                                        var v58 uint16 = 143
                                        _ = v58
                                    case -1:
                                        var v59 uint32 = (v34 | (v34 % ((42) | 1)))
                                        _ = v59
                                    default:
                                        v60 := GenS2{F0: v50}
                                        _ = v60
                                    }
                                case -9223372036854775808:
                                    if ((v1 & (^(v1))) == (v1 / (((v1 / ((v1) | 1))) | 1))) {
                                        if false {
                                            var v61 int32 = -903
                                            var v62 int = (fi38 | fi38)
                                            _ = v61
                                            _ = v62
                                        } else {
                                            fmt.Println(v2)
                                        }
                                        var v63 int32 = 42
                                        v64 := GenS4{F0: (^(v34)), F1: 0}
                                        _ = v63
                                        _ = v64
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            var v65 int8 = 0
                                            _ = v65
                                        case "hello":
                                            v28 = true
                                        default:
                                            v3 = append(v3, ("foo bar" + "a"))
                                        }
                                        switch int(v50) {
                                        case 32767:
                                            if v28 {
                                                switch ((("a" + "\n\t") + ("\n\t" + ("\n\t" + "hello"))) + "") {
                                                case "":
                                                    v66 := GenS5{F0: 1}
                                                    _ = v66
                                                case "foo bar":
                                                    if v28 {
                                                        if v28 {
                                                            fmt.Println(v37, fi38, v1)
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v67 int8 = -127
                                                        fmt.Println(v2, v1, v34)
                                                        if v28 {
                                                            v68 := GenS4{F0: v34, F1: v29}
                                                            v69 := GenS1{F0: (("\n\t" + "hello") + "foo bar"), F1: v67}
                                                            v67 -= (v67 | (-(v67)))
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                        _ = v67
                                                    }
                                                case "hello":
                                                    v2 += (v2 & 966)
                                                default:
                                                    v30 = (v30 & (v30 % (((^(v30))) | 1)))
                                                }
                                                var v70 int64 = v29
                                                _ = v70
                                            } else {
                                                var v71 uint64 = v50
                                                v1 = (v1 >> ((2) & 63))
                                                v72 := GenS5{F0: 100}
                                                v50 = (v50 * v50)
                                                v28 = v28
                                                _ = v71
                                                _ = v72
                                            }
                                        default:
                                            v73 := GenS2{F0: v50}
                                            _ = v73
                                        }
                                    }
                                    v37 = 235
                                    v5 = append(v5, "hello")
                                case -1:
                                    var v74 uint16 = (v30 << ((v2) & 63))
                                    _ = v74
                                default:
                                    var v75 uint32 = 2
                                    _ = v75
                                }
                                fmt.Println(fi36, v29, v29)
                                fmt.Println(v30)
                                v76 := GenS4{F0: (v34 & (v34 - 2147483648)), F1: v29}
                                var v77 int = (fi36 / ((fi38) | 1))
                                _ = v76
                                _ = v77
                            } else {
                                switch -9223372036854775808 {
                                case -9223372036854775808:
                                    var v78 uint8 = (v37 ^ (v37 + (v37 & v37)))
                                    _ = v78
                                case -366:
                                    _ = len(v5)
                                case 4294967295:
                                    v37 = uint8(v50)
                                }
                                fmt.Println(v34)
                            }
                        }
                        v2 = 18446744073709551615
                    }
                    _ = v50
                } else {
                    v30 = v30
                }
                fmt.Println(v1)
                _ = fi38
            }
            var v79 int8 = 0
            _ = fi36
            _ = v37
            _ = v79
        }
        _ = v35
    } else {
        for fi80 := 0; fi80 < 9; fi80++ {
            for fi81 := 0; fi81 < 11; fi81++ {
                v82 := GenS2{F0: 18446744073709551615}
                v83 := GenS3{F0: v34, F1: v30, F2: (v29 ^ 0)}
                var v84 int = fi80
                _ = fi81
                _ = v82
                _ = v83
                _ = v84
            }
            _ = fi80
        }
        var v85 uint64 = 4294967295
        if (!(!(v30 != (v30 / ((42) | 1))))) {
            switch "\n\t" {
            case "":
                fmt.Println(v1)
            case "foo bar":
                var v86 uint8 = 2
                _ = v86
            case "a":
                v87 := GenS2{F0: (^(v85))}
                _ = v87
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            fmt.Println(v29, v29)
            v29 += v29
        }
        _ = v85
    }
    v1 += v1
    var v88 float64 = 1.5
    switch "foo bar" {
    case "foo bar":
        if (!(!(v28 || v28))) {
            for fi89 := 0; fi89 < 2; fi89++ {
                v30 -= (v30 + (v30 / ((v30) | 1)))
                switch -1 {
                case 0:
                    v90 := GenS3{F0: 65535, F1: 255, F2: (v29 ^ (-(v29)))}
                    _ = v90
                case -522:
                    for fi91 := 0; fi91 < 13; fi91++ {
                        switch "\n\t" {
                        case "foo bar":
                            v88 = v88
                        case "hello":
                            fmt.Println(fi89, fi91)
                        }
                        _ = fi91
                    }
                case 128:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                default:
                    switch "\n\t" {
                    case "foo bar":
                        for fi92 := 0; fi92 < 13; fi92++ {
                            v93 := GenS2{F0: 1}
                            v28 = ((((("foo bar" + ("foo bar" + ("hello" + ""))) + ("a" + ("a" + "hello"))) + "hello") == ((((("a" + "foo bar") + "a") + "foo bar") + ("\n\t" + ("\n\t" + "foo bar"))) + "")) || (true && true))
                            var v94 bool = (!(!((!v28) || (116 != 127))))
                            _ = fi92
                            _ = v93
                            _ = v94
                        }
                    case "hello":
                        switch (fi89 % ((fi89) | 1)) {
                        case -74:
                            var v95 int8 = 127
                            _ = v95
                        case 893:
                            fmt.Println(v34, v29, v29)
                        }
                    case "":
                        var v96 float32 = 1.5
                        _ = v96
                    }
                }
                switch (fi89 % ((-1) | 1)) {
                case 1:
                    _ = len(v5)
                case 2147483647:
                    var v97 string = "a"
                    _ = v97
                case -1:
                    var v98 string = ("hello" + "hello")
                    _ = v98
                }
                _ = fi89
            }
        }
    default:
        switch 775 {
        case -542:
            var v99 int32 = -294
            _ = v99
        default:
            fmt.Println(v1, v2, v1)
        }
    }
    for fi100 := 0; fi100 < 17; fi100++ {
        _ = len(v3)
        v101 := GenS4{F0: 2147483648, F1: 0}
        _ = fi100
        _ = v101
    }
    v88 = float64(3.14159)
    if v28 {
        v4 = append(v4, (v34 ^ (v34 | uint32(v30))))
        if len(v3) > 0 {
            v3[0] = "a"
        }
        for fi102 := 0; fi102 < 17; fi102++ {
            for fi103 := 0; fi103 < 5; fi103++ {
                switch "foo bar" {
                case "a":
                    switch "\n\t" {
                    case "hello":
                        if len(v6) > 0 {
                            v6[0] = (v34 | v34)
                        }
                    case "\n\t":
                        var v104 string = "\n\t"
                        _ = v104
                    case "":
                        v1 = (v1 ^ 992)
                    }
                default:
                    switch "hello" {
                    case "hello":
                        v2 = v2
                    case "a":
                        v3 = append(v3, "\n\t")
                    }
                }
                var v105 int64 = -9223372036854775808
                _ = fi103
                _ = v105
            }
            _ = fi102
        }
    } else {
        v106 := GenS6{F0: 128, F1: ("a" + ("a" + "\n\t")), F2: (v1 % (((v1 - (v1 % (((-(v1))) | 1)))) | 1))}
        v28 = (v28 || v28)
        _ = len(v4)
        _ = v106
    }
    v107 := GenS7{F0: 4294967295, F1: 63}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v107
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v88
}
