package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int8; F1 int32; F2 int16 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 uint; F1 string; F2 int }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64 }
func genFn0(p0 uint64, p1 int64, p2 string) string {
    p1 = p1
    if (1 == -97) {
        for fi4 := 0; fi4 < 13; fi4++ {
            for fi5 := 0; fi5 < 15; fi5++ {
                switch fi4 {
                case -268:
                    v6 := GenS1{F0: (!((("foo bar" != "foo bar") && (false || (true || false))) || (65535 >= 0))), F1: 13}
                    _ = v6
                case 4294967295:
                    if false {
                        var v7 float64 = (1e10 + 0.1)
                        var v8 uint16 = uint16(p0)
                        fmt.Println(p0)
                        _ = v7
                        _ = v8
                    }
                case -9223372036854775808:
                    var v9 uint32 = 705
                    _ = v9
                }
                if ((p2 + (p2 + p2)) < p2) {
                    if (!true) {
                        v10 := GenS2{F0: 0}
                        var v11 int16 = 42
                        v12 := GenS2{F0: 659}
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        p0 = p0
                    }
                    if true {
                        switch "a" {
                        case "foo bar":
                            v13 := GenS3{F0: 32767, F1: 151}
                            switch p2 {
                            case "foo bar":
                                if (((^(p0)) >= (p0 >> ((146) & 63))) || false) {
                                    var v14 uint8 = 18
                                    _ = v14
                                } else {
                                    var v15 uint64 = (^(p0))
                                    v16 := []int{(fi4 | fi5), 32767, (fi5 % (((-(fi4))) | 1)), fi5}
                                    _ = len(v16)
                                    _ = v15
                                    _ = v16
                                }
                            case "":
                                if false {
                                    var v17 uint = 413
                                    fmt.Println(fi4, p1, v17)
                                    _ = v17
                                }
                            }
                            _ = v13
                        case "\n\t":
                            if true {
                                if false {
                                    p0 += 248
                                }
                                if ((-2.25 + (0.0 * 1e10)) != 3.14159) {
                                    p2 = ("\n\t" + p2)
                                    var v18 int32 = 42
                                    _ = v18
                                }
                                p2 = p2
                                if (3.14159 == -2.25) {
                                    v19 := GenS4{F0: -128, F1: 42, F2: -32768}
                                    var v20 uint = 255
                                    _ = v19
                                    _ = v20
                                }
                                var v21 int8 = 127
                                _ = v21
                            }
                        }
                        var v22 uint16 = 65535
                        if true {
                            var v23 int16 = -743
                            fmt.Println(v23)
                            _ = v23
                        }
                        _ = v22
                    }
                    var v24 int32 = -2147483648
                    _ = v24
                } else {
                    var v25 uint16 = 1
                    fmt.Println(p1, fi5, v25)
                    switch 127 {
                    case -1:
                        switch p2 {
                        case "a":
                            v26 := GenS3{F0: 255, F1: 0}
                            _ = v26
                        }
                        if true {
                            p1 = p1
                        } else {
                            v27 := GenS5{F0: ("\n\t" + p2), F1: -1}
                            if true {
                                var v28 float32 = (-2.25 * 0.0)
                                v29 := []int{(fi4 | (fi5 / ((fi4) | 1)))}
                                switch fi4 {
                                case 32768:
                                    fmt.Println(p2)
                                }
                                p0 -= (^(p0))
                                _ = v28
                                _ = v29
                            } else {
                                v30 := GenS3{F0: v25, F1: 128}
                                _ = v30
                            }
                            v31 := GenS1{F0: (false || true), F1: 128}
                            _ = v27
                            _ = v31
                        }
                    }
                    p1 = (p1 * 9223372036854775807)
                    _ = v25
                }
                if true {
                    var v32 uint16 = 255
                    p0 = (p0 | (^(p0)))
                    var v33 uint32 = 2
                    _ = v32
                    _ = v33
                }
                _ = fi5
            }
            if false {
                if ((!false) || false) {
                    p0 += p0
                    for fi34 := 0; fi34 < 8; fi34++ {
                        var v35 int32 = 0
                        if ((65535 >= uint(p0)) && (4294967295 >= 917)) {
                            v36 := []int{(fi4 + 32768), -632, 128, 9223372036854775807}
                            _ = v36
                        }
                        _ = fi34
                        _ = v35
                    }
                } else {
                    if true {
                        var v37 int16 = 32767
                        _ = v37
                    }
                    fmt.Println(fi4, p0)
                    p1 = 9223372036854775807
                }
                p2 = p2
            }
            var v38 uint32 = 65535
            _ = fi4
            _ = v38
        }
        var v39 float64 = (1.5 - (0.1 * 3.14159))
        _ = v39
    }
    var v40 uint64 = (p0 + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v40
    return ("hello" + p2)
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    if ((42 > 2) || (0 > 18446744073709551615)) {
        fmt.Println(v1)
        var v41 uint32 = 2
        _ = v41
    } else {
        fmt.Println(v1)
        if (100 < 100) {
            if ("hello" < (("foo bar" + "hello") + "a")) {
                for fi42 := 0; fi42 < 11; fi42++ {
                    v1 -= 1
                    if ("foo bar" <= "hello") {
                        if false {
                            v1 = v1
                            if (54 != 83) {
                                for fi43 := 0; fi43 < 15; fi43++ {
                                    var v44 int64 = 9223372036854775807
                                    _ = fi43
                                    _ = v44
                                }
                                var v45 int8 = -76
                                _ = v45
                            } else {
                                v46 := GenS4{F0: -128, F1: -1, F2: (v1 / ((32767) | 1))}
                                switch "foo bar" {
                                case "a":
                                    v47 := GenS5{F0: "a", F1: 1}
                                    _ = v47
                                case "":
                                    var v48 float32 = 1e10
                                    _ = v48
                                case "foo bar":
                                    var v49 uint64 = 4294967295
                                    _ = v49
                                }
                                v1 = (v1 & (v1 / ((v1) | 1)))
                                _ = v46
                            }
                            v1 = -32768
                        }
                        if (true && true) {
                            v1 = (v1 % ((v1) | 1))
                            v3[618] = 1
                        } else {
                            fmt.Println(v1, v1)
                            v1 -= v1
                        }
                        var v50 int16 = (v1 >> ((65535) & 63))
                        _ = v50
                    }
                    _ = fi42
                }
                var v51 int = 128
                var v52 uint16 = 65535
                _ = v51
                _ = v52
            }
            var v53 int = 42
            var v54 int32 = 578
            _ = v53
            _ = v54
        } else {
            fmt.Println(v1)
        }
        switch "foo bar" {
        case "foo bar":
            delete(v3, 2147483647)
        }
    }
    if (!true) {
        var v55 int64 = 42
        var v56 float64 = 1e10
        fmt.Println(v56, v1, v55)
        _ = v55
        _ = v56
    }
    if (!true) {
        for fi57 := 0; fi57 < 20; fi57++ {
            switch (fi57 + (fi57 + fi57)) {
            case 127:
                fmt.Println(v1, v1)
            }
            _ = fi57
        }
    } else {
        var v58 string = "a"
        var v59 string = "a"
        _ = v58
        _ = v59
    }
    if (!false) {
        v60 := GenS6{F0: 127, F1: false}
        v1 = -380
        _ = v60
    } else {
        v61 := []int{-1, 1}
        switch 1 {
        case -934:
            v1 = v1
        case 42:
            v62 := GenS2{F0: 4294967295}
            _ = v62
        case -197:
            for fi63 := 0; fi63 < 3; fi63++ {
                if ("" > "foo bar") {
                    switch (-(fi63)) {
                    case 127:
                        v1 = (v1 / (((v1 >> ((65535) & 63))) | 1))
                    case -555:
                        delete(v3, 711)
                    case 482:
                        if (540 > 65535) {
                            if (!(42 > 42)) {
                                if (1000 <= v1) {
                                    v1 = (v1 & (v1 * (-(v1))))
                                }
                                var v64 bool = (false && ((true || true) && true))
                                _ = v64
                            } else {
                                v65 := GenS7{F0: 203}
                                if len(v61) >= 2 {
                                    v61 = v61[1:len(v61)]
                                }
                                v1 -= v1
                                _ = v65
                            }
                        }
                    default:
                        fmt.Println(v1, v1)
                    }
                }
                _ = fi63
            }
        default:
            switch 261 {
            case 9223372036854775807:
                v66 := GenS8{F0: 32767, F1: 18446744073709551615}
                _ = v66
            case -921:
                for fi67 := 0; fi67 < 6; fi67++ {
                    var v68 int16 = (v1 - v1)
                    _ = fi67
                    _ = v68
                }
            case 946:
                v1 -= (v1 * v1)
            }
        }
        switch 128 {
        case 32767:
            if len(v61) > 0 {
                v61[0] = -2147483648
            }
        case 42:
            switch 2147483647 {
            case -1:
                if (true || (100 < 0)) {
                    if ((0.1 / (1e10 + float32(((((((1e10 + 0.1) / (1e10 - 0.0)) * (0.1 + (1e10 / 1e10))) - 3.14159) + (-2.25 - 3.14159)) + 0.1)))) > float32(1.5)) {
                        switch "" {
                        case "hello":
                            for fi69 := 0; fi69 < 11; fi69++ {
                                var v70 uint32 = 4294967295
                                v70 -= 867
                                if (((0.0 > 3.14159) && false) && true) {
                                    v70 = 4294967295
                                }
                                _ = fi69
                                _ = v70
                            }
                        case "foo bar":
                            switch (("a" + (("foo bar" + ("\n\t" + ("\n\t" + "\n\t"))) + ("" + "\n\t"))) + "\n\t") {
                            case "":
                                v71 := GenS7{F0: 42}
                                _ = v71
                            case "\n\t":
                                if true {
                                    v1 = 0
                                    fmt.Println(v1)
                                    v72 := GenS9{F0: 147, F1: "hello", F2: 2147483647}
                                    _ = v72
                                } else {
                                    fmt.Println(v1)
                                }
                            }
                        case "a":
                            v73 := GenS10{F0: 4294967295}
                            _ = v73
                        }
                        v1 = (v1 << ((65535) & 63))
                        v1 -= 733
                    } else {
                        var v74 int64 = int64(v1)
                        _ = v74
                    }
                    if ((v1 | (v1 ^ v1)) >= -32768) {
                        if (v1 != (v1 ^ 566)) {
                            var v75 uint8 = 1
                            _ = v75
                        }
                        if true {
                            if ((-2.25 >= 1e10) && (!(!(!true)))) {
                                v76 := v3[2147483648]
                                v77 := GenS3{F0: 1, F1: 127}
                                _ = v76
                                _ = v77
                            }
                            fmt.Println(v1)
                        } else {
                            var v78 int = 128
                            _ = v78
                        }
                        var v79 uint16 = 1
                        _ = v79
                    }
                }
            case -165:
                v1 = (v1 >> ((2) & 63))
            default:
                var v80 int8 = 100
                _ = v80
            }
        default:
            v1 = v1
        }
        _ = v61
    }
    var v81 int32 = 0
    var v82 int = -995
    var v83 int8 = -128
    for fi84 := 0; fi84 < 11; fi84++ {
        v85 := GenS11{F0: 42}
        _ = fi84
        _ = v85
    }
    v86 := []int{(v82 & (v82 % ((433) | 1))), v82, 42, 2147483647}
    var v87 int64 = 93
    for fi88 := 0; fi88 < 10; fi88++ {
        for fi89 := 0; fi89 < 15; fi89++ {
            var v90 string = ("\n\t" + "a")
            var v91 int64 = (v87 | (-(v87)))
            v81 = v81
            _ = fi89
            _ = v90
            _ = v91
        }
        var v92 int64 = (v87 % ((v87) | 1))
        v93 := v2[1]
        _ = fi88
        _ = v92
        _ = v93
    }
    var v94 uint64 = 88
    var v95 int16 = 1000
    v87 = (-(v87))
    if len(v86) >= 2 {
        v86 = v86[1:len(v86)]
    }
    v96 := GenS10{F0: 445}
    var v97 int64 = v87
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("slicelen v86:", len(v86), v86)
    fmt.Println("v87:", v87)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v81
    _ = v82
    _ = v83
    _ = v86
    _ = v87
    _ = v94
    _ = v95
    _ = v96
    _ = v97
}
