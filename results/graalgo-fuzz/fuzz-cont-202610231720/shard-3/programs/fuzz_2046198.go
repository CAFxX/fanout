package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int; F2 int }
type GenS7 struct { F0 uint64; F1 uint32; F2 int8 }
func genFn0(p0 int) uint {
    v5 := GenS1{F0: -128}
    if (!false) {
        p0 = (p0 + (p0 + p0))
        if true {
            for fi6 := 0; fi6 < 14; fi6++ {
                v7 := GenS2{F0: (((((("foo bar" + "") + "hello") + "a") + ("hello" + ((((("foo bar" + "a") + "\n\t") + "foo bar") + "\n\t") + (("foo bar" + (("a" + ("a" + "a")) + ("a" + (("foo bar" + ("\n\t" + ("foo bar" + "hello"))) + "a")))) + "")))) + "hello") + "foo bar"), F1: 2, F2: 155}
                var v8 int64 = -9223372036854775808
                _ = fi6
                _ = v7
                _ = v8
            }
            p0 -= (p0 | p0)
            p0 = 1
        } else {
            for fi9 := 0; fi9 < 10; fi9++ {
                p0 = fi9
                fmt.Println(fi9, p0)
                v10 := GenS1{F0: -58}
                _ = fi9
                _ = v10
            }
            if ("a" != "foo bar") {
                p0 = (p0 | p0)
                p0 += 9223372036854775807
                var v11 uint8 = 60
                _ = v11
            }
            var v12 int64 = 42
            _ = v12
        }
        p0 = (p0 % ((p0) | 1))
    } else {
        v13 := GenS1{F0: 0}
        if true {
            v14 := GenS1{F0: 72}
            v15 := GenS1{F0: 30}
            for fi16 := 0; fi16 < 6; fi16++ {
                switch ("hello" + "foo bar") {
                case "a":
                    fmt.Println(fi16)
                default:
                    p0 += fi16
                }
                _ = fi16
            }
            _ = v14
            _ = v15
        } else {
            switch "\n\t" {
            case "":
                for fi17 := 0; fi17 < 12; fi17++ {
                    if (false && true) {
                        p0 = p0
                        var v18 bool = ((true && (true && true)) || false)
                        _ = v18
                    }
                    v19 := []int{fi17, (fi17 - p0), (fi17 >> ((1) & 63)), 711}
                    _ = fi17
                    _ = v19
                }
            case "foo bar":
                fmt.Println(p0)
            }
            fmt.Println(p0)
            var v20 float64 = 1.5
            _ = v20
        }
        _ = v13
    }
    _ = p0
    _ = v5
    return 881
}
func genFn1() int {
    var v21 int64 = -319
    var v22 int64 = -25
    v22 += v21
    v23 := GenS3{F0: 1, F1: -9223372036854775808}
    _ = v21
    _ = v22
    _ = v23
    return 42
}

func main() {
    var v1 bool = false
    var v2 uint16 = 32767
    v3 := []string{("\n\t" + ("hello" + ""))}
    v4 := make(map[string]int)
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 string = ("" + "")
        _ = fi24
        _ = v25
    }
    if ((0.1 + 1e10) <= 3.14159) {
        if (-16 > 42) {
            fmt.Println(v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        switch 2147483647 {
        case 32767:
            if (1 > 255) {
                v26 := GenS3{F0: 948, F1: 2147483647}
                _ = v26
            } else {
                switch "foo bar" {
                case "":
                    delete(v4, "foo bar")
                case "foo bar":
                    var v27 uint16 = (^(v2))
                    _ = v27
                case "hello":
                    for fi28 := 0; fi28 < 14; fi28++ {
                        v1 = v1
                        _ = fi28
                    }
                default:
                    for fi29 := 0; fi29 < 10; fi29++ {
                        if (((65535 == 2147483647) && v1) && v1) {
                            v30 := GenS1{F0: 100}
                            v2 = (v2 >> ((genFn0(0)) & 63))
                            for fi31 := 0; fi31 < 8; fi31++ {
                                switch -2147483648 {
                                case 32767:
                                    if v1 {
                                        var v32 float32 = ((0.0 - (-2.25 * (1e10 - ((3.14159 + (3.14159 + 0.1)) * 1.5)))) + 0.1)
                                        if (!v1) {
                                            v2 -= v2
                                        }
                                        var v33 uint = 0
                                        _ = v32
                                        _ = v33
                                    }
                                case 97:
                                    if v1 {
                                        v34 := GenS1{F0: -128}
                                        if v1 {
                                            var v35 uint = 421
                                            _ = v35
                                        }
                                        v36 := GenS3{F0: (fi31 * fi29), F1: 9223372036854775807}
                                        _ = v34
                                        _ = v36
                                    } else {
                                        if (v1 && v1) {
                                            var v37 int = -9223372036854775808
                                            fmt.Println(fi29, v37, v1)
                                            _ = v37
                                        }
                                    }
                                default:
                                    switch ("a" + "\n\t") {
                                    case "hello":
                                        v38 := GenS1{F0: 0}
                                        _ = v38
                                    }
                                }
                                _ = fi31
                            }
                            _ = v30
                        }
                        _ = fi29
                    }
                }
            }
        case 481:
            switch ("foo bar" + "") {
            case "hello":
                v2 = v2
            case "foo bar":
                v39 := GenS4{F0: (("foo bar" + "foo bar") + "\n\t")}
                _ = v39
            case "":
                if true {
                    var v40 int = -605
                    v41 := GenS1{F0: 1}
                    _ = v40
                    _ = v41
                } else {
                    var v42 float64 = float64(((0.0 * 0.1) + float32((1e10 * (0.0 - 0.0)))))
                    var v43 float32 = 0.0
                    _ = v42
                    _ = v43
                }
            }
        default:
            for fi44 := 0; fi44 < 7; fi44++ {
                v1 = (v1 && v1)
                _ = len(v3)
                var v45 float32 = -2.25
                _ = fi44
                _ = v45
            }
        }
        v46 := v4[""]
        _ = v46
    } else {
        v47 := GenS2{F0: "\n\t", F1: 1, F2: 230}
        switch genFn1() {
        case -651:
            var v48 uint16 = (v2 >> ((65535) & 63))
            _ = v48
        case 1:
            v3 = append(v3, "a")
        default:
            for fi49 := 0; fi49 < 18; fi49++ {
                v2 = v2
                var v50 int64 = -9223372036854775808
                var v51 float64 = 1.5
                _ = fi49
                _ = v50
                _ = v51
            }
        }
        _ = v47
    }
    for fi52 := 0; fi52 < 10; fi52++ {
        delete(v4, ("hello" + ("\n\t" + "hello")))
        _ = fi52
    }
    v1 = true
    _ = len(v4)
    if (((v2 % ((v2) | 1)) <= 255) && true) {
        v2 = (v2 % (((v2 & (v2 & (^(v2))))) | 1))
        v53 := GenS3{F0: genFn1(), F1: 127}
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = v53
    } else {
        for fi54 := 0; fi54 < 17; fi54++ {
            var v55 int16 = 739
            _ = fi54
            _ = v55
        }
        switch "" {
        case "\n\t":
            for fi56 := 0; fi56 < 13; fi56++ {
                v57 := GenS1{F0: 127}
                _ = fi56
                _ = v57
            }
        default:
            var v58 uint16 = v2
            _ = v58
        }
    }
    v1 = (2 < 1)
    v1 = (v1 || v1)
    for fi59 := 0; fi59 < 5; fi59++ {
        v60 := GenS2{F0: "", F1: 0, F2: 42}
        if v1 {
            switch (^(fi59)) {
            case 25:
                switch genFn1() {
                case 1:
                    for fi61 := 0; fi61 < 1; fi61++ {
                        switch (fi59 >> ((2) & 63)) {
                        case 636:
                            fmt.Println(fi61)
                        case 1:
                            if (-143 <= (fi61 * fi61)) {
                                if false {
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch fi61 {
                                        case 9223372036854775807:
                                            var v62 int16 = 1000
                                            _ = v62
                                        }
                                    case "hello":
                                        var v63 bool = (128 > 0)
                                        _ = v63
                                    default:
                                        v64 := GenS5{F0: 2147483648}
                                        _ = v64
                                    }
                                } else {
                                    var v65 int32 = 2147483647
                                    _ = v65
                                }
                                v66 := GenS3{F0: fi59, F1: fi61}
                                _ = v66
                            }
                        case 127:
                            if v1 {
                                v2 = 253
                                var v67 int8 = 29
                                var v68 int = (fi61 & 101)
                                _ = v67
                                _ = v68
                            }
                        default:
                            switch ("\n\t" + ("\n\t" + ("a" + ""))) {
                            case "a":
                                if (1 < 2) {
                                    v69 := GenS6{F0: 407, F1: (fi61 & (-(fi59))), F2: (fi61 + (fi59 >> ((18446744073709551615) & 63)))}
                                    _ = v69
                                }
                                v2 -= v2
                                var v70 float32 = -2.25
                                _ = v70
                            case "":
                                v71 := GenS7{F0: 1, F1: 2147483647, F2: -8}
                                _ = v71
                            }
                        }
                        _ = fi61
                    }
                default:
                    v2 -= (^(v2))
                }
            case -300:
                switch (fi59 + 42) {
                case -1:
                    var v72 uint32 = 64
                    _ = v72
                }
            }
        } else {
            v73 := GenS6{F0: 2, F1: -450, F2: (fi59 & fi59)}
            for fi74 := 0; fi74 < 20; fi74++ {
                var v75 string = ("hello" + "")
                if (v2 < 674) {
                    switch v75 {
                    case "foo bar":
                        var v76 uint64 = 1
                        _ = v76
                    default:
                        v77 := GenS6{F0: 2, F1: (fi59 * fi74), F2: 32767}
                        _ = v77
                    }
                }
                _ = fi74
                _ = v75
            }
            _ = v73
        }
        var v78 int = (^(fi59))
        _ = fi59
        _ = v60
        _ = v78
    }
    if (42 > 0) {
        switch 127 {
        case 32767:
            var v79 string = ((("\n\t" + "") + "hello") + "\n\t")
            _ = v79
        case 1:
            var v80 uint16 = (v2 - (v2 % (((v2 - (v2 | v2))) | 1)))
            _ = v80
        }
        var v81 float32 = float32(1e10)
        _ = len(v3)
        _ = v81
    } else {
        for fi82 := 0; fi82 < 8; fi82++ {
            v1 = (0.1 < 1.5)
            switch (fi82 >> ((genFn0(fi82)) & 63)) {
            case 910:
                fmt.Println(v1, v1, fi82)
            case -1:
                _ = len(v3)
            case 9223372036854775807:
                switch ("a" + "hello") {
                case "\n\t":
                    v1 = (562 <= 65535)
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            default:
                for fi83 := 0; fi83 < 5; fi83++ {
                    delete(v4, (("\n\t" + "") + "hello"))
                    if ((0 == 208) || v1) {
                        switch "a" {
                        case "a":
                            if (v1 && (v1 && v1)) {
                                delete(v4, "")
                            } else {
                                if (0 > 1) {
                                    v84 := GenS1{F0: -128}
                                    if (false && true) {
                                        v2 = (^(v2))
                                    }
                                    var v85 uint16 = v2
                                    _ = v84
                                    _ = v85
                                }
                                v4[""] = (fi83 | fi83)
                                var v86 uint8 = 0
                                _ = v86
                            }
                        case "\n\t":
                            if true {
                                v2 -= v2
                                var v87 uint8 = 0
                                switch 313 {
                                case 0:
                                    if v1 {
                                        var v88 float64 = 3.14159
                                        v1 = v1
                                        v89 := GenS2{F0: "\n\t", F1: 1, F2: 1}
                                        v88 += v88
                                        _ = v88
                                        _ = v89
                                    }
                                default:
                                    v1 = ("hello" > ("a" + ("\n\t" + "hello")))
                                }
                                _ = v87
                            }
                            fmt.Println(v2, fi82)
                        default:
                            _ = len(v4)
                        }
                        v1 = (192 > 0)
                    } else {
                        v2 = (v2 * (^(v2)))
                        v1 = v1
                        v90 := GenS7{F0: 4294967295, F1: 460, F2: 0}
                        _ = v90
                    }
                    var v91 uint8 = 128
                    _ = fi83
                    _ = v91
                }
            }
            _ = fi82
        }
        for fi92 := 0; fi92 < 9; fi92++ {
            v93 := GenS4{F0: ""}
            _ = fi92
            _ = v93
        }
        var v94 bool = (true && (!false))
        _ = v94
    }
    v2 += 2
    var v95 bool = v1
    var v96 uint8 = 2
    fmt.Println(v95, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v95
    _ = v96
}
