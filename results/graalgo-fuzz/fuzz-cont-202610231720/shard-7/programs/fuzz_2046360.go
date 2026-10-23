package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int16; F1 bool; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 uint32; F1 uint32; F2 string }
type GenS7 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS8 struct { F0 uint; F1 uint16; F2 uint }
type GenS9 struct { F0 string }
func genFn0(p0 int8) uint {
    if ((32767 != 32767) && (844 <= 1000)) {
        if false {
            v7 := GenS1{F0: 2147483647}
            _ = v7
        } else {
            if true {
                switch -787 {
                case -9223372036854775808:
                    if ((true && true) || (false || (32767 != 32767))) {
                        for fi8 := 0; fi8 < 13; fi8++ {
                            if (-2.25 <= 0.1) {
                                if false {
                                    for fi9 := 0; fi9 < 1; fi9++ {
                                        p0 = 127
                                        if (!true) {
                                            p0 += p0
                                        } else {
                                            if (2147483647 <= -2147483648) {
                                                v10 := GenS1{F0: -2147483648}
                                                switch "" {
                                                case "foo bar":
                                                    p0 = p0
                                                case "a":
                                                    var v11 bool = (42 < 2147483647)
                                                    _ = v11
                                                }
                                                _ = v10
                                            } else {
                                                v12 := []int{589, (fi8 & fi9), (fi8 << ((65535) & 63))}
                                                var v13 uint = 563
                                                p0 -= p0
                                                _ = v12
                                                _ = v13
                                            }
                                            p0 = (p0 >> ((255) & 63))
                                            var v14 uint = 255
                                            _ = v14
                                        }
                                        if false {
                                            if (!true) {
                                                var v15 int = (fi8 >> ((255) & 63))
                                                _ = v15
                                            } else {
                                                var v16 int32 = 2147483647
                                                _ = v16
                                            }
                                        } else {
                                            p0 += 9
                                            p0 = p0
                                            var v17 bool = (0 != 42)
                                            _ = v17
                                        }
                                        _ = fi9
                                    }
                                    p0 -= 42
                                } else {
                                    if (748 < 0) {
                                        v18 := []int{-645, 128}
                                        if (!true) {
                                            var v19 int32 = -2147483648
                                            _ = v19
                                        } else {
                                            if false {
                                                fmt.Println(p0)
                                                var v20 float64 = 3.14159
                                                _ = len(v18)
                                                _ = v20
                                            }
                                            p0 += (p0 % ((0) | 1))
                                            for fi21 := 0; fi21 < 18; fi21++ {
                                                v22 := GenS1{F0: -2147483648}
                                                fmt.Println(p0, fi8)
                                                p0 += (^(p0))
                                                _ = fi21
                                                _ = v22
                                            }
                                        }
                                        _ = v18
                                    }
                                }
                                fmt.Println(fi8)
                            } else {
                                if (2 >= 4294967295) {
                                    for fi23 := 0; fi23 < 4; fi23++ {
                                        var v24 uint64 = 18446744073709551615
                                        v25 := GenS2{F0: 10, F1: (v24 << ((255) & 63))}
                                        v26 := []int{(fi23 * 127), fi8, fi8}
                                        _ = fi23
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                    for fi27 := 0; fi27 < 9; fi27++ {
                                        switch "" {
                                        case "\n\t":
                                            if (true || true) {
                                                v28 := GenS2{F0: 130, F1: 0}
                                                _ = v28
                                            }
                                        }
                                        p0 = p0
                                        var v29 uint16 = 26
                                        _ = fi27
                                        _ = v29
                                    }
                                }
                            }
                            _ = fi8
                        }
                    }
                default:
                    p0 = (p0 * (^(p0)))
                }
                fmt.Println(p0)
                var v30 int = 205
                _ = v30
            } else {
                p0 = -128
            }
            v31 := GenS2{F0: 127, F1: 42}
            p0 -= p0
            _ = v31
        }
    }
    v32 := GenS3{F0: "\n\t", F1: 75}
    p0 -= (p0 >> ((2) & 63))
    _ = p0
    _ = v32
    return 18446744073709551615
}

func main() {
    v1 := []int{32768, 128, 469, 695}
    v2 := []int{128, -1, 9223372036854775807, -252}
    v3 := []int{32768, -9223372036854775808}
    v4 := []string{}
    v5 := make(map[uint32]bool)
    v6 := []string{"foo bar", ("hello" + "hello")}
    var v33 uint64 = 105
    for fi34 := 0; fi34 < 2; fi34++ {
        if false {
            var v35 uint16 = 255
            if false {
                v35 -= (v35 & (v35 << ((uint(v33)) & 63)))
                if (false && (42 == (v35 >> ((2) & 63)))) {
                    v36 := GenS3{F0: "foo bar", F1: 42}
                    _ = v36
                }
                v1 = append(v1, 127)
            }
            var v37 float64 = 1e10
            _ = v35
            _ = v37
        }
        _ = fi34
    }
    v33 -= 954
    switch "foo bar" {
    case "":
        for fi38 := 0; fi38 < 9; fi38++ {
            switch "foo bar" {
            case "":
                v39 := GenS4{F0: 1000, F1: true, F2: 42}
                _ = v39
            case "hello":
                delete(v5, 2)
            default:
                v40 := GenS3{F0: "", F1: 42}
                _ = v40
            }
            v41 := GenS5{F0: (v33 + v33), F1: 1}
            v33 -= (v33 | 4294967295)
            _ = fi38
            _ = v41
        }
    case "hello":
        var v42 bool = false
        _ = v42
    }
    fmt.Println(v33)
    if false {
        if true {
            v33 = v33
            var v43 uint16 = 32767
            var v44 int64 = 9223372036854775807
            _ = v43
            _ = v44
        }
        fmt.Println(v33)
        if true {
            for fi45 := 0; fi45 < 4; fi45++ {
                for fi46 := 0; fi46 < 16; fi46++ {
                    if false {
                        fmt.Println(fi46)
                        v47 := GenS3{F0: "", F1: -241}
                        fmt.Println(fi45, fi46)
                        _ = v47
                    } else {
                        if (true && (false && (((301 < 493) || (-2147483648 > -252)) || false))) {
                            var v48 bool = false
                            _ = v48
                        }
                        v33 += (v33 * 18446744073709551615)
                        if false {
                            v49 := GenS6{F0: 2, F1: 2147483647, F2: "foo bar"}
                            var v50 int16 = 0
                            _ = v49
                            _ = v50
                        }
                        var v51 int16 = 1000
                        _ = v51
                    }
                    switch (("foo bar" + ("foo bar" + "\n\t")) + "a") {
                    case "\n\t":
                        var v52 int64 = 1
                        _ = v52
                    case "":
                        switch 815 {
                        case 0:
                            if (-2.25 == 1e10) {
                                if (v33 <= v33) {
                                    v5[0] = (42 != -863)
                                    var v53 int16 = 0
                                    delete(v5, 0)
                                    _ = v53
                                }
                                v54 := GenS1{F0: 1}
                                var v55 uint16 = 704
                                _ = v54
                                _ = v55
                            }
                        case 204:
                            v33 = (^(v33))
                        default:
                            if (!(true && (((((-2.25 * (3.14159 - 1e10)) + -2.25) / 1.0) - float64(((float32(0.0) + float32((0.1 / ((1.5 / (0.1 + 1e10)) + ((1e10 - -2.25) + 1e10))))) * 0.1))) > (float64((1.5 + (0.0 - 1.5))) * 3.14159)))) {
                                var v56 uint32 = 4294967295
                                v5[v56] = (57 >= 1)
                                _ = v56
                            }
                        }
                    case "hello":
                        fmt.Println(fi46, fi45)
                    }
                    fmt.Println(fi45, fi46)
                    switch fi45 {
                    case 4294967295:
                        if (127 == -128) {
                            var v57 uint32 = 2147483647
                            _ = v57
                        }
                    case -2147483648:
                        v33 = 744
                    case 9223372036854775807:
                        if false {
                            if true {
                                if false {
                                    var v58 int64 = 1
                                    _ = v58
                                }
                                switch "\n\t" {
                                case "":
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "\n\t":
                                            v5[1] = false
                                        case "a":
                                            _ = len(v3)
                                        }
                                    case "":
                                        v33 = (v33 & 1)
                                    case "a":
                                        v33 += (^(v33))
                                    default:
                                        var v59 int32 = 128
                                        _ = v59
                                    }
                                case "foo bar":
                                    if false {
                                        delete(v5, 65535)
                                    } else {
                                        v33 = 4294967295
                                        v60 := GenS2{F0: 1, F1: v33}
                                        _ = v60
                                    }
                                }
                            } else {
                                v33 -= 0
                            }
                            fmt.Println(v33, fi46)
                            var v61 int8 = 48
                            _ = v61
                        }
                    default:
                        v33 = (^(v33))
                    }
                    if true {
                        v33 -= v33
                    }
                    switch fi45 {
                    case 32767:
                        v62 := v5[2147483647]
                        _ = v62
                    case -9223372036854775808:
                        var v63 float64 = (0.1 - (0.0 + (1e10 * 0.0)))
                        _ = v63
                    case 210:
                        if (255 == 132) {
                            v64 := GenS4{F0: -979, F1: (1 < 273), F2: 700}
                            var v65 int32 = int32(fi45)
                            v33 = v33
                            v65 = int32(fi46)
                            _ = v64
                            _ = v65
                        } else {
                            v33 += 2
                            if false {
                                switch (^(fi45)) {
                                case 9223372036854775807:
                                    var v66 int = fi46
                                    _ = v66
                                case 127:
                                    if (fi46 >= -818) {
                                        v33 -= v33
                                    }
                                case -847:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                if (false && false) {
                                    var v67 int = (fi46 << ((47) & 63))
                                    switch (("foo bar" + "") + ("\n\t" + "")) {
                                    case "a":
                                        var v68 int16 = 1000
                                        if (false && true) {
                                            v69 := GenS7{F0: (!(604 != (v33 & (v33 % ((v33) | 1))))), F1: 146, F2: 4294967295}
                                            _ = v69
                                        } else {
                                            v70 := GenS8{F0: uint(v33), F1: 0, F2: 1}
                                            _ = v70
                                        }
                                        _ = v68
                                    case "hello":
                                        if true {
                                            if true {
                                                v67 += -1
                                                fmt.Println(fi46, v33)
                                            }
                                            delete(v5, 2)
                                        } else {
                                            if true {
                                                v33 = v33
                                                v67 -= 574
                                            }
                                        }
                                    }
                                    delete(v5, 42)
                                    _ = v67
                                } else {
                                    _ = len(v5)
                                    _ = len(v4)
                                    v3 = append(v3, fi46)
                                }
                            }
                            var v71 int8 = 100
                            _ = v71
                        }
                    }
                    _ = fi46
                }
                switch "hello" {
                case "hello":
                    if (false && ((-2147483648 >= -1) || true)) {
                        v33 = v33
                    } else {
                        switch (("hello" + "a") + ((("foo bar" + ("" + "")) + "foo bar") + ("hello" + (((("" + "") + "foo bar") + "") + (("a" + ("a" + ("\n\t" + "a"))) + (("\n\t" + (("hello" + "a") + "foo bar")) + "")))))) {
                        case "hello":
                            if false {
                                v72 := GenS9{F0: "foo bar"}
                                v2 = append(v2, fi45)
                                _ = v72
                            }
                        }
                        for fi73 := 0; fi73 < 10; fi73++ {
                            fmt.Println(fi45, fi73, fi45)
                            if false {
                                fmt.Println(fi45, fi73)
                            }
                            _ = fi73
                        }
                        v33 = v33
                    }
                default:
                    var v74 uint32 = 2147483647
                    _ = v74
                }
                _ = fi45
            }
        }
    }
    var v75 uint16 = 489
    var v76 float32 = float32(-2.25)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    v5[776] = (1 < (^(v33)))
    fmt.Println(v75, v33)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
}
