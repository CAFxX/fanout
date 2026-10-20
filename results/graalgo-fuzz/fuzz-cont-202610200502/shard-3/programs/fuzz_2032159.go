package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 uint8; F2 bool }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64 }
func genFn0(p0 int64, p1 uint64, p2 int64) int64 {
    v7 := []int{200, 2147483647, 32767}
    p0 -= 9223372036854775807
    v8 := GenS1{F0: "a"}
    var v9 int64 = (p2 << ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p2
}

func main() {
    v1 := []string{"\n\t", "a", "hello", "a"}
    var v2 bool = false
    var v3 int16 = 1000
    var v4 float64 = -2.25
    var v5 int64 = 0
    v6 := make(map[uint32]bool)
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: ""}
        v12 := GenS1{F0: "a"}
        var v13 int64 = (v5 + -9223372036854775808)
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 bool = ((v3 + (v3 << ((601) & 63))) <= 32767)
    if ("\n\t" >= ("foo bar" + "hello")) {
        var v15 int32 = 2147483647
        var v16 uint64 = 1
        if v2 {
            v17 := GenS1{F0: (("\n\t" + "foo bar") + ("foo bar" + ""))}
            v1 = append(v1, ("hello" + ("hello" + "foo bar")))
            _ = v17
        } else {
            switch 128 {
            case 127:
                v2 = (!(v5 >= genFn0(9223372036854775807, (v16 << ((2) & 63)), v5)))
            case 453:
                for fi18 := 0; fi18 < 7; fi18++ {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v4 += (((float64(0.1) / (1e10 + (float64(-2.25) - v4))) + v4) - (1e10 + ((1e10 * v4) * v4)))
                        _ = fi19
                    }
                    _ = fi18
                }
            default:
                v2 = false
            }
            for fi20 := 0; fi20 < 12; fi20++ {
                v1 = append(v1, "foo bar")
                var v21 int32 = -1
                fmt.Println(v16, v16)
                _ = fi20
                _ = v21
            }
        }
        _ = v15
        _ = v16
    }
    fmt.Println(v5, v2)
    switch -9223372036854775808 {
    case -925:
        var v22 int8 = 1
        _ = v22
    }
    var v23 uint = 838
    v3 -= 42
    switch -1 {
    case 9223372036854775807:
        v24 := GenS2{F0: 164, F1: 758}
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := GenS3{F0: 42, F1: 42, F2: v3}
        var v27 int32 = 0
        switch ("" + ("foo bar" + "hello")) {
        case "a":
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 string = ""
                _ = fi28
                _ = v29
            }
        case "":
            v30 := GenS4{F0: (-(v3))}
            _ = v30
        default:
            switch ("\n\t" + ("\n\t" + "")) {
            case "foo bar":
                if true {
                    switch fi25 {
                    case -9223372036854775808:
                        v23 -= (v23 * (v23 - 0))
                    case 32767:
                        for fi31 := 0; fi31 < 11; fi31++ {
                            var v32 int32 = v27
                            _ = fi31
                            _ = v32
                        }
                    }
                } else {
                    switch "foo bar" {
                    case "a":
                        v33 := GenS5{F0: 42, F1: ("hello" + "a")}
                        _ = v33
                    case "foo bar":
                        switch (("hello" + "hello") + ("hello" + ("hello" + "hello"))) {
                        case "foo bar":
                            v34 := v6[54]
                            _ = v34
                        case "a":
                            v35 := GenS6{F0: 100}
                            _ = v35
                        case "hello":
                            var v36 int8 = -128
                            _ = v36
                        default:
                            if (!(int(v3) <= -9223372036854775808)) {
                                switch (-(fi25)) {
                                case 4294967295:
                                    v3 += -96
                                }
                                v37 := GenS1{F0: ""}
                                _ = v37
                            } else {
                                fmt.Println(v2)
                                var v38 float64 = ((v4 + (((-2.25 * ((v4 * (1e10 - (0.0 - ((0.0 / 0.1) + v4)))) * -2.25)) + (0.0 - (1.5 - 3.14159))) * (3.14159 - (((float64(3.14159) - v4) / -2.25) - 1e10)))) - v4)
                                if v14 {
                                    v38 = float64((float32(((((v4 * v4) + v4) - v4) + (((3.14159 * ((v38 - (1e10 - v4)) - 3.14159)) * float64(float32(v38))) / v38))) / float32(v38)))
                                    fmt.Println(v4)
                                    for fi39 := 0; fi39 < 4; fi39++ {
                                        var v40 int32 = (v27 >> (((^(v23))) & 63))
                                        fmt.Println(fi25)
                                        if ((v38 - v38) != v4) {
                                            v3 += v3
                                            switch "" {
                                            case "":
                                                var v41 uint16 = 1
                                                _ = v41
                                            case "\n\t":
                                                switch (((((("foo bar" + "hello") + "a") + "\n\t") + "") + "hello") + "\n\t") {
                                                case "\n\t":
                                                    var v42 string = ("" + "foo bar")
                                                    if (((float32(v38) + ((0.1 - ((((((0.1 * (float32((v4 + v4)) + 3.14159)) - 0.0) / 3.14159) - -2.25) + ((0.0 * 1e10) * (float32(v4) + (-2.25 * (1.5 - (1.5 / 1.0)))))) * 1e10)) / (float32((float64(3.14159) - v4)) / (3.14159 / 1.5)))) / 1.0) != ((3.14159 - 0.0) + 0.0)) {
                                                        if (!true) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v43 float32 = (0.0 * 0.1)
                                                            _ = v43
                                                        } else {
                                                            if (0 == 127) {
                                                                var v44 int16 = (v3 << (((v23 * 255)) & 63))
                                                                if (v2 && true) {
                                                                    var v45 uint = v23
                                                                    _ = v45
                                                                }
                                                                _ = v44
                                                            } else {
                                                                v46 := GenS7{F0: (fi25 | fi25), F1: 0, F2: v2}
                                                                v27 += -179
                                                                _ = v46
                                                            }
                                                        }
                                                        v27 -= v40
                                                        v23 -= v23
                                                        v47 := GenS8{F0: 2147483647, F1: -59}
                                                        switch "" {
                                                        case "\n\t":
                                                            var v48 int = fi39
                                                            _ = v48
                                                        }
                                                        _ = v47
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        fmt.Println(v2, v38)
                                                        v3 = (v3 | (v3 << (((v23 << (((v23 ^ v23)) & 63))) & 63)))
                                                    }
                                                    _ = v42
                                                case "":
                                                    if v2 {
                                                        switch (("a" + "a") + "hello") {
                                                        case "foo bar":
                                                            switch fi39 {
                                                            case -9223372036854775808:
                                                                switch (fi25 / (((^(fi39))) | 1)) {
                                                                case 4294967295:
                                                                    var v49 uint = v23
                                                                    _ = v49
                                                                case 2147483647:
                                                                    var v50 float64 = (float64(1e10) - (3.14159 * v4))
                                                                    _ = v50
                                                                }
                                                            case 0:
                                                                v51 := v6[65535]
                                                                _ = v51
                                                            case 9223372036854775807:
                                                                v3 = (v3 << ((v23) & 63))
                                                            default:
                                                                fmt.Println(v27, fi39)
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                switch (fi25 >> ((2) & 63)) {
                                                                case 32768:
                                                                    switch 0 {
                                                                    case -9223372036854775808:
                                                                        v6[2147483647] = true
                                                                    case 943:
                                                                        var v52 int = -2147483648
                                                                        _ = v52
                                                                    }
                                                                case 4294967295:
                                                                    switch (fi39 + (fi39 - fi39)) {
                                                                    case -557:
                                                                        var v53 uint32 = 2147483647
                                                                        _ = v53
                                                                    }
                                                                case 388:
                                                                    if (65535 > 0) {
                                                                        if true {
                                                                            if ("hello" == ("foo bar" + "")) {
                                                                                v23 = 0
                                                                                if (!false) {
                                                                                    v2 = (((v14 && v2) && (!(!v2))) || v14)
                                                                                    v5 = 786
                                                                                } else {
                                                                                    var v54 float64 = v4
                                                                                    if false {
                                                                                        if v2 {
                                                                                            _ = len(v1)
                                                                                        }
                                                                                    } else {
                                                                                        var v55 uint32 = 42
                                                                                        _ = v55
                                                                                    }
                                                                                    fmt.Println(v3, fi25)
                                                                                    _ = v54
                                                                                }
                                                                                v40 -= (v40 ^ (v27 | v40))
                                                                            }
                                                                            fmt.Println(v27, v3, v27)
                                                                            fmt.Println(v38, fi39, v5)
                                                                        }
                                                                        _ = len(v6)
                                                                        var v56 uint = 784
                                                                        _ = v56
                                                                    } else {
                                                                        v5 -= (v5 ^ (v5 / ((genFn0((v5 * -9223372036854775808), 198, 9223372036854775807)) | 1)))
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v40 = -492
                                                                switch "\n\t" {
                                                                case "":
                                                                    var v57 int16 = v3
                                                                    _ = v57
                                                                default:
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v58 := GenS8{F0: (^(fi25)), F1: (v5 / ((genFn0((v5 / ((v5) | 1)), 1, v5)) | 1))}
                                                                        _ = v58
                                                                    default:
                                                                        v1 = append(v1, "")
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                v59 := GenS4{F0: v3}
                                                                _ = v59
                                                            }
                                                        case "":
                                                            v60 := GenS9{F0: v14}
                                                            _ = v60
                                                        case "a":
                                                            var v61 int64 = (v5 * v5)
                                                            _ = v61
                                                        default:
                                                            switch ("a" + "") {
                                                            case "hello":
                                                                var v62 uint64 = 43
                                                                _ = v62
                                                            }
                                                        }
                                                        v63 := GenS4{F0: (v3 & (v3 >> ((v23) & 63)))}
                                                        if (int(v3) == 2147483647) {
                                                            var v64 int8 = 76
                                                            var v65 float64 = float64(1e10)
                                                            if v14 {
                                                                v66 := GenS8{F0: (fi25 << ((v23) & 63)), F1: v5}
                                                                _ = v66
                                                            } else {
                                                                if (332 <= 65535) {
                                                                    fmt.Println(fi25)
                                                                    fmt.Println(v64, v27, v65)
                                                                } else {
                                                                    var v67 int16 = v3
                                                                    _ = v67
                                                                }
                                                                if v14 {
                                                                    var v68 uint16 = 164
                                                                    switch fi25 {
                                                                    case -782:
                                                                        var v69 uint = 668
                                                                        fmt.Println(v4, v23, v4)
                                                                        _ = v69
                                                                    case -984:
                                                                        fmt.Println(v68, v68, v27)
                                                                        delete(v6, 2147483647)
                                                                        var v70 uint16 = (v68 ^ (^(v68)))
                                                                        _ = v70
                                                                    case 9223372036854775807:
                                                                        switch (("a" + "foo bar") + "") {
                                                                        case "":
                                                                            v71 := GenS5{F0: 2, F1: "\n\t"}
                                                                            _ = v71
                                                                        }
                                                                    default:
                                                                        v72 := v6[393]
                                                                        _ = v72
                                                                    }
                                                                    _ = v68
                                                                } else {
                                                                    var v73 float32 = 1e10
                                                                    if (965 != int16(fi25)) {
                                                                        var v74 int8 = v64
                                                                        _ = v74
                                                                    }
                                                                    var v75 float64 = ((v4 + (v4 / float64((float32((float64(v73) * v38)) - v73)))) + (1.5 - 0.1))
                                                                    _ = v73
                                                                    _ = v75
                                                                }
                                                                _ = len(v1)
                                                            }
                                                            if ("" == "foo bar") {
                                                                v76 := GenS6{F0: v64}
                                                                _ = v76
                                                            } else {
                                                                if v14 {
                                                                    v77 := GenS9{F0: (v2 || false)}
                                                                    _ = v77
                                                                }
                                                            }
                                                            _ = v64
                                                            _ = v65
                                                        } else {
                                                            if (!false) {
                                                                delete(v6, 4294967295)
                                                                v78 := GenS5{F0: 42, F1: "a"}
                                                                v14 = v14
                                                                _ = v78
                                                            }
                                                            fmt.Println(fi25, v5)
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                        _ = v63
                                                    } else {
                                                        v79 := v6[33]
                                                        _ = v79
                                                    }
                                                case "a":
                                                    fmt.Println(v38, v2)
                                                default:
                                                    v80 := GenS10{F0: genFn0((v5 % (((v5 - 96)) | 1)), 646, v5)}
                                                    _ = v80
                                                }
                                            case "a":
                                                if v2 {
                                                    var v81 int8 = 38
                                                    _ = v81
                                                } else {
                                                    v82 := GenS9{F0: v14}
                                                    if ((float32((0.0 - v38)) != 3.14159) && (!false)) {
                                                        v40 -= v27
                                                    } else {
                                                        v4 -= (v38 * v4)
                                                        v23 = v23
                                                        fmt.Println(v27, v3, v3)
                                                        var v83 uint8 = 127
                                                        _ = v83
                                                    }
                                                    var v84 string = ""
                                                    _ = v82
                                                    _ = v84
                                                }
                                                var v85 int = (fi25 + fi25)
                                                var v86 uint32 = 2147483647
                                                v2 = (v5 == v5)
                                                v2 = v2
                                                var v87 float64 = (0.1 * -2.25)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v85
                                                _ = v86
                                                _ = v87
                                            }
                                        }
                                        _ = fi39
                                        _ = v40
                                    }
                                } else {
                                    if v14 {
                                        switch (fi25 ^ fi25) {
                                        case -2147483648:
                                            switch "a" {
                                            case "hello":
                                                for fi88 := 0; fi88 < 16; fi88++ {
                                                    if v2 {
                                                        var v89 bool = (((3.14159 / 1.0) * ((((0.1 / (float32((-2.25 - 1e10)) / 1.0)) - float32((1e10 - (3.14159 * 1.5)))) + (0.1 / 1.5)) * 3.14159)) < 1e10)
                                                        v3 = v3
                                                        fmt.Println(v5, v89)
                                                        _ = v89
                                                    }
                                                    var v90 uint32 = 377
                                                    _ = fi88
                                                    _ = v90
                                                }
                                            case "foo bar":
                                                var v91 string = (("\n\t" + "foo bar") + "")
                                                _ = v91
                                            default:
                                                if ((v5 << (((v23 * v23)) & 63)) == (v5 % ((v5) | 1))) {
                                                    v14 = true
                                                }
                                            }
                                        default:
                                            v4 = 0.0
                                        }
                                    } else {
                                        v38 -= (((v38 / v38) * 1.5) / 1e10)
                                        v92 := GenS1{F0: ""}
                                        _ = v92
                                    }
                                }
                                _ = v38
                            }
                        }
                    default:
                        var v93 float32 = 3.14159
                        _ = v93
                    }
                }
            case "\n\t":
                switch ("" + "foo bar") {
                case "":
                    var v94 int16 = (v3 ^ (v3 % ((v3) | 1)))
                    _ = v94
                default:
                    var v95 string = "hello"
                    _ = v95
                }
            case "hello":
                switch (fi25 + (fi25 / ((-786) | 1))) {
                case -768:
                    if ((2147483647 < int32(v23)) && v2) {
                        if v2 {
                            _ = len(v6)
                        } else {
                            var v96 uint8 = 127
                            _ = v96
                        }
                    } else {
                        v97 := v6[2147483647]
                        var v98 string = ("" + "foo bar")
                        _ = v97
                        _ = v98
                    }
                case 4294967295:
                    _ = len(v1)
                default:
                    if (true || v14) {
                        fmt.Println(v2)
                    }
                }
            }
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v14, v2)
    var v99 int64 = 473
    var v100 float64 = v4
    var v101 int16 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
