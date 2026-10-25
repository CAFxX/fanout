package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int8; F1 int8; F2 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint32; F1 bool; F2 int32 }
type GenS10 struct { F0 uint16; F1 int16; F2 int64 }
func genFn0(p0 uint64) uint64 {
    v7 := GenS1{F0: 237, F1: 87}
    var v8 int = -9223372036854775808
    _ = p0
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int8, p1 uint32) uint64 {
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS1{F0: 0, F1: (p0 << ((1) & 63))}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS1{F0: 0, F1: p0}
    switch 42 {
    case 802:
        var v13 int = -799
        _ = v13
    case 2147483647:
        if (!((p1 << ((65535) & 63)) < 1)) {
            var v14 int16 = 90
            if (366 > 1) {
                v14 = (v14 / (((^(v14))) | 1))
                v15 := GenS1{F0: 0, F1: (p0 ^ (p0 >> ((65535) & 63)))}
                v16 := []int{}
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 6; fi17++ {
                    v14 = -32768
                    var v18 uint16 = 835
                    _ = fi17
                    _ = v18
                }
                v19 := []int{}
                for fi20 := 0; fi20 < 2; fi20++ {
                    if (42 == 131) {
                        var v21 uint8 = 48
                        _ = v21
                    } else {
                        var v22 uint32 = 0
                        if true {
                            v23 := GenS1{F0: 4294967295, F1: 127}
                            _ = v23
                        } else {
                            for fi24 := 0; fi24 < 18; fi24++ {
                                var v25 int8 = 42
                                var v26 string = "a"
                                p1 += 470
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            if (127 >= 12) {
                                switch ("a" + ("a" + "\n\t")) {
                                case "":
                                    for fi27 := 0; fi27 < 19; fi27++ {
                                        var v28 uint32 = (p1 ^ 2)
                                        switch (fi27 ^ fi27) {
                                        case 42:
                                            if ((0.1 - float64((((0.0 - (1.5 + float32((1e10 / -2.25)))) / ((1e10 * ((0.1 * 0.1) - -2.25)) / float32((1.5 * 0.1)))) + (1.5 - float32(0.1))))) == (0.1 + 1e10)) {
                                                if false {
                                                    if ("hello" >= "\n\t") {
                                                        var v29 int = (fi20 / ((int(v14)) | 1))
                                                        v30 := GenS1{F0: 603, F1: (p0 % ((-1) | 1))}
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        p1 = 65535
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v31 int = (fi20 | -2147483648)
                                                            _ = v31
                                                        default:
                                                            var v32 string = ""
                                                            _ = v32
                                                        }
                                                        if ((!true) && true) {
                                                            if len(v19) > 0 {
                                                                v19[0] = fi20
                                                            }
                                                            p1 -= (p1 >> ((42) & 63))
                                                        } else {
                                                            v33 := GenS1{F0: genFn0(102), F1: (p0 % (((p0 << ((0) & 63))) | 1))}
                                                            v28 = (v22 ^ p1)
                                                            v22 = v28
                                                            _ = v33
                                                        }
                                                    }
                                                }
                                                v19 = append(v19, fi27)
                                            } else {
                                                var v34 uint32 = v28
                                                _ = v34
                                            }
                                        case 919:
                                            p1 = v28
                                        default:
                                            switch ("\n\t" + ("foo bar" + "hello")) {
                                            case "foo bar":
                                                var v35 uint32 = v28
                                                _ = v35
                                            case "hello":
                                                switch (fi20 | fi27) {
                                                case -9223372036854775808:
                                                    switch "a" {
                                                    case "a":
                                                        var v36 bool = true
                                                        _ = v36
                                                    case "\n\t":
                                                        var v37 int = 1
                                                        _ = v37
                                                    default:
                                                        v19 = append(v19, (fi27 + fi27))
                                                    }
                                                case 0:
                                                    v38 := GenS2{F0: (p0 - (p0 % ((p0) | 1))), F1: "foo bar", F2: 887}
                                                    _ = v38
                                                case 42:
                                                    if (311 == -384) {
                                                        var v39 float32 = (0.0 - 3.14159)
                                                        if len(v19) > 0 {
                                                            v19[0] = (fi27 + -1)
                                                        }
                                                        if ((!false) && (("foo bar" + "a") != "")) {
                                                            fmt.Println(v22)
                                                        }
                                                        _ = v39
                                                    } else {
                                                        var v40 uint64 = 289
                                                        v22 = 1
                                                        var v41 string = "foo bar"
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi27
                                        _ = v28
                                    }
                                case "a":
                                    switch ("a" + ("foo bar" + "a")) {
                                    case "foo bar":
                                        if (true || (true && (true && true))) {
                                            switch "a" {
                                            case "":
                                                v42 := GenS2{F0: (-(p0)), F1: (("foo bar" + "") + "foo bar"), F2: 42}
                                                _ = v42
                                            case "hello":
                                                for fi43 := 0; fi43 < 5; fi43++ {
                                                    p1 = (p1 | v22)
                                                    var v44 uint64 = 771
                                                    var v45 uint = 75
                                                    _ = fi43
                                                    _ = v44
                                                    _ = v45
                                                }
                                            }
                                        }
                                    case "":
                                        v22 -= (v22 >> ((2) & 63))
                                    default:
                                        for fi46 := 0; fi46 < 15; fi46++ {
                                            if true {
                                                if (!true) {
                                                    p1 += 0
                                                    v47 := GenS2{F0: 1, F1: ("a" + ("foo bar" + (("" + "a") + ("hello" + "a")))), F2: 0}
                                                    _ = v47
                                                }
                                                v22 += v22
                                                var v48 uint16 = 65535
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                                fmt.Println(fi20)
                                                _ = v48
                                            } else {
                                                fmt.Println(fi46, p0, v14)
                                            }
                                            var v49 string = "\n\t"
                                            _ = fi46
                                            _ = v49
                                        }
                                    }
                                default:
                                    fmt.Println(v22)
                                }
                            } else {
                                v14 = 1
                                v14 += v14
                            }
                        }
                        v50 := GenS3{F0: ""}
                        _ = v22
                        _ = v50
                    }
                    v51 := GenS4{F0: 1}
                    if false {
                        var v52 int = fi20
                        var v53 int16 = (-(v14))
                        _ = v52
                        _ = v53
                    } else {
                        switch "a" {
                        case "foo bar":
                            var v54 uint16 = 881
                            _ = v54
                        default:
                            switch "a" {
                            case "hello":
                                for fi55 := 0; fi55 < 19; fi55++ {
                                    switch ("a" + "foo bar") {
                                    case "\n\t":
                                        if (!true) {
                                            if false {
                                                if len(v19) > 0 {
                                                    v19[0] = (fi20 ^ (fi20 * (-(fi20))))
                                                }
                                            } else {
                                                p0 += p0
                                            }
                                            var v56 uint16 = 215
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                            _ = v56
                                        } else {
                                            switch "\n\t" {
                                            case "\n\t":
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                            }
                                            var v57 int32 = -2147483648
                                            _ = v57
                                        }
                                    default:
                                        if len(v19) > 0 {
                                            v19[0] = fi20
                                        }
                                    }
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    _ = fi55
                                }
                            case "\n\t":
                                var v58 uint8 = 70
                                _ = v58
                            case "foo bar":
                                var v59 int32 = 42
                                _ = v59
                            default:
                                fmt.Println(fi20, v14)
                            }
                        }
                        p1 += 4294967295
                        if (!(true || (42 > 2))) {
                            for fi60 := 0; fi60 < 5; fi60++ {
                                fmt.Println(p0)
                                var v61 int64 = 9223372036854775807
                                _ = fi60
                                _ = v61
                            }
                            v62 := GenS2{F0: p0, F1: ("" + "hello"), F2: -1}
                            _ = v62
                        }
                    }
                    _ = fi20
                    _ = v51
                }
                _ = v19
            }
            _ = v14
        } else {
            if false {
                p1 = 2
            }
            p0 += p0
            p0 = (p0 - p0)
        }
    case -234:
        for fi63 := 0; fi63 < 13; fi63++ {
            v64 := GenS5{F0: p0, F1: 92}
            for fi65 := 0; fi65 < 13; fi65++ {
                p1 -= (p1 | 2147483647)
                v66 := GenS1{F0: genFn0(291), F1: p0}
                if ("hello" != "a") {
                    v67 := []int{32767, (fi63 / ((-545) | 1))}
                    _ = len(v67)
                    _ = v67
                } else {
                    if false {
                        v68 := GenS1{F0: 4294967295, F1: p0}
                        switch fi63 {
                        case -701:
                            p0 += 37
                        case -577:
                            var v69 int16 = 32767
                            var v70 uint8 = 224
                            _ = v69
                            _ = v70
                        case 42:
                            switch ((("hello" + "foo bar") + ("a" + "")) + "hello") {
                            case "\n\t":
                                v71 := []int{fi65, 32767, fi63}
                                _ = v71
                            }
                        default:
                            v72 := GenS6{F0: 658, F1: 42}
                            _ = v72
                        }
                        var v73 uint = 1
                        _ = v68
                        _ = v73
                    }
                    p1 = (p1 << ((4294967295) & 63))
                    var v74 int32 = 2147483647
                    _ = v74
                }
                if (867 != 428) {
                    p1 -= p1
                    var v75 uint8 = 193
                    _ = v75
                } else {
                    v76 := GenS7{F0: (^(p0)), F1: 1, F2: 1000}
                    p1 += p1
                    p0 = p0
                    _ = v76
                }
                _ = fi65
                _ = v66
            }
            _ = fi63
            _ = v64
        }
    }
    p0 = p0
    p0 += (p0 + 127)
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]int)
    var v3 int = 127
    var v4 int = 4294967295
    v5 := []string{"a", "foo bar"}
    v6 := make(map[uint32]string)
    var v77 uint64 = 42
    if (42 < 42) {
        v78 := GenS8{F0: -1}
        _ = v78
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi79 := 0; fi79 < 8; fi79++ {
        v80 := GenS3{F0: ("foo bar" + "")}
        v81 := v2[("" + ("a" + ""))]
        _ = fi79
        _ = v80
        _ = v81
    }
    v77 -= v77
    var v82 float32 = 0.0
    switch "\n\t" {
    case "a":
        v83 := GenS9{F0: 2147483648, F1: false, F2: -2147483648}
        _ = v83
    case "\n\t":
        var v84 uint = 65535
        _ = v84
    case "foo bar":
        v85 := GenS10{F0: 255, F1: 0, F2: -499}
        _ = v85
    }
    if true {
        switch "\n\t" {
        case "a":
            var v86 uint16 = 42
            _ = v86
        case "":
            if (9223372036854775807 <= -189) {
                var v87 int8 = 117
                if (2 != v77) {
                    for fi88 := 0; fi88 < 3; fi88++ {
                        for fi89 := 0; fi89 < 6; fi89++ {
                            switch (^(v3)) {
                            case -270:
                                var v90 uint8 = 1
                                _ = v90
                            case 127:
                                var v91 string = "foo bar"
                                _ = v91
                            default:
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(fi89)
                                }
                            }
                            v92 := GenS9{F0: 689, F1: (2 != 127), F2: -1}
                            _ = fi89
                            _ = v92
                        }
                        _ = fi88
                    }
                    var v93 float64 = 3.14159
                    _ = v93
                }
                _ = v87
            } else {
                v94 := GenS2{F0: 0, F1: (("" + "foo bar") + "\n\t"), F2: 582}
                if false {
                    v77 = genFn1(1, 2)
                    fmt.Println(v4, v4)
                    for fi95 := 0; fi95 < 8; fi95++ {
                        var v96 uint8 = 147
                        _ = fi95
                        _ = v96
                    }
                } else {
                    v77 = (^(v77))
                    switch "" {
                    case "\n\t":
                        var v97 uint = 1
                        _ = v97
                    default:
                        switch "hello" {
                        case "":
                            v3 += 32767
                        default:
                            var v98 int = (^(v4))
                            _ = v98
                        }
                    }
                }
                _ = v94
            }
        default:
            v82 -= (v82 / 1.5)
        }
    } else {
        var v99 uint16 = 859
        _ = v99
    }
    fmt.Println(v4)
    var v100 uint16 = 2
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v77:", v77)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v82
}
