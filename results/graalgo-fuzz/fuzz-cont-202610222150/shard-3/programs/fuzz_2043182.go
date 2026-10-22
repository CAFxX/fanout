package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0(p0 int32) uint16 {
    for fi6 := 0; fi6 < 3; fi6++ {
        p0 = 1
        for fi7 := 0; fi7 < 5; fi7++ {
            switch fi6 {
            case 42:
                fmt.Println(fi6, p0, fi7)
            }
            if true {
                if (54 > -128) {
                    p0 -= -2147483648
                    p0 += 1
                    if false {
                        fmt.Println(fi6, p0, fi6)
                        v8 := []int{127}
                        _ = v8
                    }
                } else {
                    p0 -= (p0 * int32(fi7))
                    fmt.Println(p0)
                    var v9 int = fi7
                    _ = v9
                }
            }
            switch "foo bar" {
            case "a":
                var v10 bool = false
                _ = v10
            case "\n\t":
                p0 = (^(p0))
            }
            switch ((("\n\t" + "\n\t") + "a") + ("foo bar" + ("\n\t" + "foo bar"))) {
            case "a":
                p0 -= 237
            case "":
                if (1 > 208) {
                    var v11 uint8 = 188
                    _ = v11
                }
                if (!("hello" != "foo bar")) {
                    p0 = 2147483647
                }
                switch fi6 {
                case 32767:
                    p0 = p0
                case -55:
                    if false {
                        var v12 int16 = -32768
                        v13 := GenS1{F0: 42, F1: "hello"}
                        var v14 uint = 18446744073709551615
                        if true {
                            var v15 string = (("hello" + "foo bar") + "")
                            _ = v15
                        }
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(p0)
                        p0 -= p0
                    }
                }
                var v16 float32 = 3.14159
                _ = v16
            default:
                if true {
                    var v17 int = (fi7 << ((22) & 63))
                    var v18 int16 = 300
                    fmt.Println(fi7, fi6)
                    _ = v17
                    _ = v18
                }
            }
            _ = fi7
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := []int{(fi19 * (fi19 * 0)), (fi19 * fi19), (fi6 * 9223372036854775807)}
            if len(v20) > 0 {
                v20[0] = (-(fi19))
            }
            if false {
                fmt.Println(fi6, fi19, p0)
            } else {
                fmt.Println(p0)
            }
            switch ("a" + "hello") {
            case "a":
                v21 := GenS1{F0: p0, F1: ("hello" + ((("foo bar" + ("hello" + "hello")) + "hello") + "\n\t"))}
                _ = v21
            case "hello":
                p0 -= -2147483648
            case "":
                switch ("a" + "\n\t") {
                case "":
                    var v22 uint16 = 1
                    var v23 uint64 = 42
                    _ = v22
                    _ = v23
                case "\n\t":
                    p0 = (p0 | p0)
                case "hello":
                    switch "a" {
                    case "":
                        p0 = p0
                    }
                }
            default:
                p0 -= -849
                switch "" {
                case "a":
                    if (2 != 42) {
                        v24 := GenS1{F0: int32(fi6), F1: ((("\n\t" + "") + "\n\t") + (("\n\t" + "") + "\n\t"))}
                        if true {
                            if (true && true) {
                                v25 := GenS1{F0: p0, F1: ""}
                                v20 = append(v20, (fi19 << ((42) & 63)))
                                _ = v25
                            }
                        }
                        var v26 int = (^(fi19))
                        if false {
                            switch ("a" + (("" + "\n\t") + ("hello" + "\n\t"))) {
                            case "":
                                var v27 float64 = ((0.0 * -2.25) * (float64((1e10 * 3.14159)) * (3.14159 / 1e10)))
                                _ = v27
                            case "hello":
                                fmt.Println(fi6, v26, fi19)
                                if (true || false) {
                                    var v28 string = "a"
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    _ = v28
                                }
                            case "foo bar":
                                p0 = p0
                            default:
                                switch fi19 {
                                case -9223372036854775808:
                                    var v29 uint32 = 12
                                    _ = v29
                                case 127:
                                    fmt.Println(p0)
                                case -2147483648:
                                    v26 -= (-(v26))
                                }
                            }
                        } else {
                            if len(v20) > 0 {
                                v20[0] = (v26 * -1)
                            }
                        }
                        v26 -= (v26 % ((1) | 1))
                        _ = v24
                        _ = v26
                    }
                case "foo bar":
                    var v30 float32 = ((((1e10 - (1.5 + 0.0)) / -2.25) + (-2.25 / float32(1e10))) / ((3.14159 * (0.1 + (((0.1 * (1.5 - ((3.14159 * ((((1.5 + 3.14159) + (((-2.25 / 3.14159) - -2.25) + float32(1e10))) * -2.25) - (1.5 / 0.1))) * ((0.1 * (0.0 - (3.14159 / 1.0))) + 1e10)))) * 1.5) - 0.0))) / (3.14159 + (float32(((1e10 * 1e10) / 1.0)) / 1e10))))
                    _ = v30
                default:
                    p0 = (p0 >> ((2) & 63))
                }
            }
            var v31 bool = false
            fmt.Println(v31, v31)
            _ = fi19
            _ = v20
            _ = v31
        }
        _ = fi6
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        for fi33 := 0; fi33 < 5; fi33++ {
            p0 -= p0
            p0 -= 740
            _ = fi33
        }
        v34 := []int{fi32, fi32}
        _ = fi32
        _ = v34
    }
    p0 += (p0 ^ p0)
    for fi35 := 0; fi35 < 6; fi35++ {
        fmt.Println(fi35, p0)
        fmt.Println(fi35, fi35)
        _ = fi35
    }
    fmt.Println(p0)
    _ = p0
    return 32767
}
func genFn1(p0 uint64, p1 uint64) int8 {
    switch -2147483648 {
    case 0:
        for fi36 := 0; fi36 < 8; fi36++ {
            v37 := GenS1{F0: 1, F1: (("" + ("a" + ("foo bar" + "foo bar"))) + "\n\t")}
            for fi38 := 0; fi38 < 14; fi38++ {
                fmt.Println(fi36, fi36, fi38)
                _ = fi38
            }
            var v39 string = "a"
            _ = fi36
            _ = v37
            _ = v39
        }
    default:
        var v40 int8 = 0
        _ = v40
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        v42 := GenS2{F0: 864}
        _ = fi41
        _ = v42
    }
    v43 := []int{-2147483648, 9223372036854775807, 4294967295, 810}
    var v44 uint16 = genFn0(42)
    _ = p0
    _ = p1
    _ = v43
    _ = v44
    return -125
}

func main() {
    var v1 uint8 = 127
    v2 := []int{-9223372036854775808}
    var v3 uint16 = 255
    var v4 int64 = 0
    v5 := []uint32{706, 502}
    var v45 uint64 = 18446744073709551615
    for fi46 := 0; fi46 < 18; fi46++ {
        v3 -= (v3 * 2)
        var v47 uint16 = 1
        _ = fi46
        _ = v47
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v48 bool = false
    v1 = (v1 - 2)
    v2 = append(v2, 32768)
    v45 += v45
    for fi49 := 0; fi49 < 14; fi49++ {
        if ((v4 ^ (v4 / ((v4) | 1))) != (v4 / ((42) | 1))) {
            fmt.Println(v45, v4, v3)
        }
        switch 42 {
        case 1:
            var v50 uint = 0
            _ = v50
        case 32768:
            v51 := GenS3{F0: (^(v3))}
            _ = v51
        }
        if len(v2) > 0 {
            v2[0] = fi49
        }
        _ = fi49
    }
    if ((!v48) && false) {
        fmt.Println(v48, v1)
        if (0 < 2) {
            fmt.Println(v48)
        } else {
            if v48 {
                if (v45 != (v45 + (v45 & (v45 << ((0) & 63))))) {
                    v3 = (v3 * (v3 & v3))
                    if true {
                        if v48 {
                            switch 42 {
                            case -2147483648:
                                v48 = true
                            }
                            var v52 uint8 = 255
                            v4 = v4
                            _ = v52
                        }
                        var v53 uint64 = (^(v45))
                        for fi54 := 0; fi54 < 2; fi54++ {
                            v55 := GenS2{F0: 1}
                            switch "\n\t" {
                            case "hello":
                                var v56 float64 = 3.14159
                                _ = v56
                            }
                            _ = fi54
                            _ = v55
                        }
                        _ = v53
                    } else {
                        v2 = append(v2, 42)
                    }
                } else {
                    var v57 int = -975
                    v48 = true
                    _ = v57
                }
                switch "hello" {
                case "":
                    for fi58 := 0; fi58 < 9; fi58++ {
                        fmt.Println(v4)
                        _ = fi58
                    }
                case "\n\t":
                    v45 -= 379
                case "a":
                    var v59 uint64 = (v45 % ((263) | 1))
                    _ = v59
                default:
                    var v60 string = ""
                    _ = v60
                }
            } else {
                if v48 {
                    var v61 float64 = (0.1 + (0.0 / -2.25))
                    fmt.Println(v61, v4)
                    for fi62 := 0; fi62 < 14; fi62++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi63 := 0; fi63 < 20; fi63++ {
                            if ((v4 % ((v4) | 1)) != v4) {
                                v1 += v1
                                v64 := GenS4{F0: -978, F1: 21}
                                _ = len(v5)
                                _ = v64
                            } else {
                                var v65 int16 = -1
                                var v66 int16 = -1
                                v5 = append(v5, 0)
                                var v67 int8 = 42
                                var v68 bool = ("hello" <= "")
                                if v68 {
                                    switch "a" {
                                    case "foo bar":
                                        v69 := GenS5{F0: (fi62 & (fi63 * fi63)), F1: (v1 - v1)}
                                        v70 := GenS1{F0: -1, F1: "\n\t"}
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v71 := GenS1{F0: -2147483648, F1: ("hello" + "")}
                                        v1 = (^(v1))
                                        var v72 uint16 = v3
                                        v2 = append(v2, fi62)
                                        if (!v68) {
                                            v48 = (v68 || ((-2147483648 >= 406) || v48))
                                        }
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    case "hello":
                                        switch ("" + "hello") {
                                        case "a":
                                            v66 = -605
                                        case "hello":
                                            var v73 float32 = ((((float32(v61) * 3.14159) * 0.1) / 3.14159) * (1e10 * float32(-2.25)))
                                            _ = v73
                                        case "foo bar":
                                            v74 := GenS6{F0: (^(v66)), F1: 927, F2: v1}
                                            v61 += v61
                                            _ = v74
                                        default:
                                            if v48 {
                                                if (117 < v67) {
                                                    var v75 uint8 = v1
                                                    if (true && v68) {
                                                        v76 := GenS3{F0: (^(v3))}
                                                        var v77 string = ("a" + ((("\n\t" + "") + ("" + "\n\t")) + "hello"))
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    _ = v75
                                                } else {
                                                    v78 := GenS1{F0: -2147483648, F1: "a"}
                                                    _ = v78
                                                }
                                            } else {
                                                v3 += (v3 & v3)
                                                if v68 {
                                                    switch ("a" + "\n\t") {
                                                    case "\n\t":
                                                        v68 = (!(((v1 ^ uint8(v3)) < (v1 + (^(v1)))) || v48))
                                                    case "hello":
                                                        var v79 string = "foo bar"
                                                        _ = v79
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = 2
                                                    }
                                                    v5 = append(v5, 42)
                                                    fmt.Println(fi63, fi63)
                                                } else {
                                                    v80 := GenS6{F0: (^(v65)), F1: 0, F2: 8}
                                                    v65 = v66
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(v1, v1, fi63)
                                                    case "a":
                                                        var v81 float64 = v61
                                                        _ = v81
                                                    case "":
                                                        v66 = -1
                                                    }
                                                    var v82 string = ""
                                                    _ = v80
                                                    _ = v82
                                                }
                                            }
                                        }
                                    case "a":
                                        switch ("hello" + "a") {
                                        case "\n\t":
                                            if v68 {
                                                v61 = v61
                                                v4 -= v4
                                                var v83 uint64 = v45
                                                _ = v83
                                            } else {
                                                v84 := GenS1{F0: 953, F1: "foo bar"}
                                                _ = v84
                                            }
                                        default:
                                            fmt.Println(v66, v61, v66)
                                        }
                                    default:
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v85 uint8 = (^(v1))
                                            v3 = (v3 % (((^(v3))) | 1))
                                            _ = v85
                                        }
                                    }
                                    v4 += v4
                                    if v48 {
                                        v48 = (v3 > (^(v3)))
                                    } else {
                                        fmt.Println(v45)
                                        fmt.Println(v1)
                                        var v86 float64 = (v61 + float64(0.1))
                                        _ = v86
                                    }
                                }
                                _ = v65
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                            var v87 uint8 = v1
                            _ = fi63
                            _ = v87
                        }
                        _ = fi62
                    }
                    _ = v61
                } else {
                    for fi88 := 0; fi88 < 15; fi88++ {
                        for fi89 := 0; fi89 < 12; fi89++ {
                            v45 += v45
                            var v90 float64 = 1.5
                            _ = fi89
                            _ = v90
                        }
                        _ = len(v2)
                        var v91 uint64 = (v45 ^ v45)
                        _ = fi88
                        _ = v91
                    }
                    var v92 uint8 = uint8(v3)
                    _ = v92
                }
                v48 = v48
                if v48 {
                    v93 := GenS6{F0: 42, F1: 0, F2: v1}
                    fmt.Println(v4, v45, v3)
                    v1 -= (v1 >> ((42) & 63))
                    _ = v93
                }
            }
            switch 128 {
            case 318:
                var v94 int16 = 1000
                _ = v94
            case -714:
                if true {
                    fmt.Println(v45, v4)
                    for fi95 := 0; fi95 < 10; fi95++ {
                        fmt.Println(v48, v1)
                        if ((float32((0.1 + float64((-2.25 * -2.25)))) + (((1.5 * (0.1 / 1.5)) + 0.0) + -2.25)) <= 3.14159) {
                            for fi96 := 0; fi96 < 19; fi96++ {
                                var v97 float64 = 1.5
                                _ = fi96
                                _ = v97
                            }
                            v48 = true
                            for fi98 := 0; fi98 < 9; fi98++ {
                                fmt.Println(fi95, v3, v4)
                                var v99 int8 = 0
                                var v100 uint32 = 65535
                                _ = fi98
                                _ = v100
                                _ = v99
                            }
                        }
                        var v101 uint32 = 2
                        _ = fi95
                        _ = v101
                    }
                    var v102 float32 = 1.5
                    _ = v102
                }
            case 317:
                switch 42 {
                case 4294967295:
                    _ = len(v5)
                }
            }
            v2 = append(v2, 128)
        }
    }
    var v103 int8 = -65
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
}
