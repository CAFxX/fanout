package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 bool; F2 int64 }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 string; F1 uint }
func genFn0(p0 string, p1 uint8, p2 string) int {
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := GenS1{F0: 0, F1: 42, F2: 42}
            v7 := []int{fi4, 128, fi4}
            p2 = ""
            var v8 int64 = 0
            switch "a" {
            case "a":
                var v9 int8 = int8(v8)
                _ = v9
            default:
                v10 := GenS1{F0: 100, F1: 0, F2: 76}
                _ = v10
            }
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        fmt.Println(p2, p0)
        for fi11 := 0; fi11 < 17; fi11++ {
            switch (fi4 << ((18446744073709551615) & 63)) {
            case 9223372036854775807:
                v12 := GenS1{F0: -81, F1: 1, F2: 2}
                _ = v12
            case 2147483647:
                if (!((fi11 != -511) && true)) {
                    v13 := []int{}
                    v14 := GenS2{F0: 100, F1: 42, F2: 2}
                    p0 = p0
                    var v15 int32 = 42
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    p2 = p0
                }
            }
            _ = fi11
        }
        _ = fi4
    }
    switch 0 {
    case 9223372036854775807:
        v16 := GenS3{F0: 42}
        _ = v16
    case -187:
        fmt.Println(p2)
    }
    var v17 uint64 = 4294967295
    p1 = (p1 - 128)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1000
}
func genFn1(p0 uint8) int64 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -252
}

func main() {
    v1 := []string{"a", "\n\t"}
    v2 := make(map[string]int)
    var v3 int32 = -1
    v3 = v3
    var v18 float32 = 1.5
    var v19 uint8 = 2
    switch "\n\t" {
    case "":
        v20 := v2[("a" + "foo bar")]
        _ = v20
    default:
        var v21 int32 = v3
        _ = v21
    }
    v22 := v2["a"]
    for fi23 := 0; fi23 < 12; fi23++ {
        if true {
            v24 := GenS3{F0: -1}
            v2["hello"] = fi23
            _ = v24
        } else {
            v3 -= v3
        }
        _ = fi23
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 uint16 = 1
        var v27 int64 = -1
        _ = fi25
        _ = v26
        _ = v27
    }
    if true {
        v2[(("foo bar" + "foo bar") + ((("a" + "hello") + (("" + "\n\t") + ("" + "hello"))) + "foo bar"))] = (v22 - genFn0("\n\t", v19, ""))
        var v28 int8 = 0
        _ = v28
    } else {
        if (234 >= 1) {
            for fi29 := 0; fi29 < 7; fi29++ {
                if (715 > 32767) {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        if (18446744073709551615 != 0) {
                            var v31 int8 = 1
                            v32 := GenS3{F0: -42}
                            _ = v31
                            _ = v32
                        } else {
                            var v33 int = (v22 + v22)
                            if false {
                                v19 = v19
                                if (0 <= 204) {
                                    if (1 <= 65535) {
                                        fmt.Println(v22)
                                        v33 += fi30
                                        var v34 uint8 = v19
                                        _ = v34
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                            } else {
                                var v35 string = "foo bar"
                                v36 := GenS2{F0: 1, F1: 2, F2: 592}
                                _ = v35
                                _ = v36
                            }
                            switch genFn0((((("" + "") + "foo bar") + "a") + "hello"), 127, ("foo bar" + ("hello" + ("" + "a")))) {
                            case -1:
                                fmt.Println(v18, v19, fi29)
                                _ = len(v2)
                            case 93:
                                if (((((-2.25 + (-2.25 + -2.25)) + (0.1 - (3.14159 * (((((1.5 / 1.5) + 0.1) - 3.14159) / -2.25) / 0.1)))) / 0.1) - -2.25) != (0.1 - ((0.0 - -2.25) + float64(v18)))) {
                                    v37 := GenS3{F0: 0}
                                    if true {
                                        if true {
                                            switch fi30 {
                                            case 32767:
                                                switch -305 {
                                                case 2147483647:
                                                    if (("a" + ((("foo bar" + ("foo bar" + "foo bar")) + "hello") + "")) < "\n\t") {
                                                        fmt.Println(v3, v18, v33)
                                                    }
                                                case -2147483648:
                                                    v22 += (fi29 & -9223372036854775808)
                                                    v19 -= (v19 - (v19 * 255))
                                                    v33 -= v33
                                                    if (!(-1 < 1)) {
                                                        v22 = (fi30 >> ((92) & 63))
                                                    }
                                                    v3 += v3
                                                    v38 := GenS4{F0: -91, F1: v19, F2: 127}
                                                    v3 += (v3 >> ((1) & 63))
                                                    v39 := GenS2{F0: 1, F1: 18446744073709551615, F2: 1}
                                                    var v40 uint = 255
                                                    _ = v38
                                                    _ = v39
                                                    _ = v40
                                                case 32767:
                                                    if (false && (((v19 >> ((1) & 63)) > v19) && (!true))) {
                                                        fmt.Println(v22)
                                                        if true {
                                                            v1 = append(v1, "a")
                                                            v41 := GenS2{F0: 42, F1: 768, F2: 803}
                                                            var v42 int64 = 111
                                                            var v43 uint16 = 32767
                                                            _ = v41
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                        var v44 uint32 = 65535
                                                        _ = v44
                                                    }
                                                default:
                                                    if false {
                                                        switch fi29 {
                                                        case 128:
                                                            var v45 int64 = 9223372036854775807
                                                            _ = v45
                                                        case -303:
                                                            var v46 uint = 603
                                                            _ = v46
                                                        case 9223372036854775807:
                                                            v18 = (v18 - 3.14159)
                                                        default:
                                                            var v47 string = ((("" + "\n\t") + ("foo bar" + "a")) + ("a" + "a"))
                                                            _ = v47
                                                        }
                                                        if false {
                                                            v3 -= 0
                                                            var v48 int8 = 13
                                                            var v49 uint16 = 1
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                        switch "hello" {
                                                        case "":
                                                            v3 = v3
                                                        case "a":
                                                            fmt.Println(v22)
                                                        }
                                                    } else {
                                                        switch "" {
                                                        case "":
                                                            v33 += 128
                                                        case "hello":
                                                            _ = len(v1)
                                                        case "foo bar":
                                                            v3 += v3
                                                        }
                                                        if (42 != 42) {
                                                            v19 = v19
                                                            var v50 uint = 65535
                                                            if false {
                                                                var v51 int16 = -1
                                                                switch fi29 {
                                                                case 1:
                                                                    fmt.Println(v18, fi29, fi30)
                                                                case 127:
                                                                    v1 = append(v1, (("foo bar" + "foo bar") + ""))
                                                                default:
                                                                    fmt.Println(v18)
                                                                }
                                                                _ = v51
                                                            }
                                                            _ = v50
                                                        }
                                                    }
                                                }
                                            case -1:
                                                var v52 uint = 1
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (true || true) {
                                                    switch genFn0("", (v19 - (v19 - (v19 % ((52) | 1)))), (("" + (("" + "\n\t") + ("foo bar" + "\n\t"))) + (((("\n\t" + "hello") + ("hello" + ("a" + (("foo bar" + ("foo bar" + "foo bar")) + "foo bar")))) + ("\n\t" + "")) + ("\n\t" + ("\n\t" + ("" + "hello")))))) {
                                                    case 42:
                                                        if true {
                                                            _ = len(v2)
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                    case 2147483647:
                                                        if true {
                                                            var v53 uint16 = 309
                                                            v22 = (fi29 - -437)
                                                            _ = v53
                                                        } else {
                                                            v18 += 1e10
                                                            if ((!(false || (!true))) || false) {
                                                                var v54 uint8 = (v19 << ((582) & 63))
                                                                var v55 float64 = 0.1
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            if false {
                                                                var v56 float64 = float64(v18)
                                                                _ = v56
                                                            } else {
                                                                v3 += -2147483648
                                                            }
                                                        }
                                                    case 128:
                                                        delete(v2, ((("a" + "a") + "a") + ((("hello" + "foo bar") + ("a" + "foo bar")) + "foo bar")))
                                                    default:
                                                        fmt.Println(fi29, fi29)
                                                    }
                                                    var v57 uint32 = 95
                                                    _ = v57
                                                } else {
                                                    v3 += (v3 * (v3 >> (((^(v52))) & 63)))
                                                    var v58 int32 = v3
                                                    var v59 float32 = (0.0 * 0.0)
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v52
                                            case -147:
                                                v18 += (float32(3.14159) * (v18 * v18))
                                            default:
                                                if (false && false) {
                                                    var v60 int8 = 42
                                                    _ = v60
                                                } else {
                                                    v22 = (v33 & (v33 + (fi30 - fi29)))
                                                    if (42 < 199) {
                                                        if true {
                                                            var v61 string = "hello"
                                                            v33 += 9223372036854775807
                                                            v62 := GenS1{F0: -103, F1: 117, F2: 41}
                                                            fmt.Println(fi30, fi29)
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            v63 := GenS3{F0: 1000}
                                                            if true {
                                                                v18 = 0.0
                                                                fmt.Println(v19, v22, fi30)
                                                                if false {
                                                                    v18 -= float32(0.1)
                                                                } else {
                                                                    if (-1 <= 42) {
                                                                        fmt.Println(v22)
                                                                        var v64 uint = 299
                                                                        _ = v64
                                                                    }
                                                                }
                                                            }
                                                            v3 += (v3 % ((v3) | 1))
                                                            _ = v63
                                                        }
                                                        v33 += -298
                                                        v65 := GenS1{F0: 16, F1: 426, F2: 0}
                                                        _ = v65
                                                    } else {
                                                        if (-128 >= 127) {
                                                            v33 += (fi30 ^ 1)
                                                            fmt.Println(fi30)
                                                        } else {
                                                            v66 := v2["foo bar"]
                                                            _ = v66
                                                        }
                                                        if false {
                                                            v18 += v18
                                                            if (!(42 != 255)) {
                                                                switch v22 {
                                                                case 32767:
                                                                    if (!(((-1 == 337) || false) && (18446744073709551615 > 42))) {
                                                                        fmt.Println(v33, v33)
                                                                        v22 -= 32768
                                                                    }
                                                                default:
                                                                    var v67 uint64 = 853
                                                                    _ = v67
                                                                }
                                                                v3 = (-(v3))
                                                                v1 = append(v1, (("\n\t" + "foo bar") + ("a" + "foo bar")))
                                                            }
                                                            v68 := GenS5{F0: 65535, F1: "hello"}
                                                            _ = v68
                                                        }
                                                        var v69 uint8 = 1
                                                        _ = v69
                                                    }
                                                }
                                            }
                                            v3 -= (v3 % (((v3 << ((255) & 63))) | 1))
                                            if false {
                                                var v70 uint16 = 1
                                                var v71 uint32 = 2147483647
                                                _ = v70
                                                _ = v71
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            v3 -= v3
                                            var v72 uint = 0
                                            var v73 float32 = (float32(3.14159) * (float32(float64(0.1)) * ((v18 + ((float32((-2.25 * 1.5)) * (float32(-2.25) - (v18 + 0.1))) / v18)) - v18)))
                                            _ = v72
                                            _ = v73
                                        }
                                        var v74 int8 = -1
                                        v2[""] = (fi30 - (^(fi29)))
                                        _ = v74
                                    }
                                    var v75 int64 = 9223372036854775807
                                    _ = v37
                                    _ = v75
                                } else {
                                    fmt.Println(v33, fi30)
                                    v1 = append(v1, ("hello" + ("a" + "a")))
                                }
                            }
                            _ = v33
                        }
                        _ = fi30
                    }
                    var v76 uint8 = v19
                    var v77 float64 = ((0.1 + 1.5) / (0.1 + 3.14159))
                    _ = v76
                    _ = v77
                } else {
                    v78 := GenS3{F0: -36}
                    _ = v78
                }
                if (((("a" + "\n\t") + "\n\t") > ((("a" + "hello") + "foo bar") + "")) && (false || true)) {
                    switch 9223372036854775807 {
                    case 42:
                        switch "hello" {
                        case "hello":
                            v79 := GenS2{F0: -41, F1: 993, F2: 0}
                            _ = v79
                        case "":
                            v3 = (v3 - (v3 << ((65535) & 63)))
                        }
                    }
                    var v80 int8 = 0
                    fmt.Println(v18, v22)
                    _ = v80
                } else {
                    var v81 float32 = v18
                    var v82 int8 = 42
                    _ = v81
                    _ = v82
                }
                _ = fi29
            }
        }
    }
    v18 -= v18
    fmt.Println(v22, v22)
    v1 = append(v1, "hello")
    v83 := GenS3{F0: -1}
    _ = len(v2)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    if (970 == 0) {
        switch (v22 + (v22 >> ((776) & 63))) {
        case -1:
            for fi84 := 0; fi84 < 7; fi84++ {
                if len(v1) > 0 {
                    v1[0] = ""
                }
                _ = fi84
            }
        default:
            var v85 uint64 = 0
            _ = v85
        }
        v18 -= v18
    } else {
        var v86 float32 = (v18 - 3.14159)
        if (false && false) {
            var v87 uint64 = 4294967295
            delete(v2, "")
            for fi88 := 0; fi88 < 14; fi88++ {
                var v89 int8 = 99
                _ = len(v2)
                fmt.Println(fi88, v19, fi88)
                _ = fi88
                _ = v89
            }
            _ = v87
        }
        _ = v86
    }
    v22 = v22
    if true {
        v2["hello"] = (v22 % ((v22) | 1))
        var v90 uint8 = (v19 % ((v19) | 1))
        _ = v90
    } else {
        var v91 int16 = 0
        _ = v91
    }
    for fi92 := 0; fi92 < 5; fi92++ {
        v93 := GenS4{F0: 239, F1: 212, F2: -119}
        for fi94 := 0; fi94 < 11; fi94++ {
            if (!false) {
                switch (^(v22)) {
                case 0:
                    switch 128 {
                    case -348:
                        v22 -= fi92
                    case 105:
                        v22 += (v22 + fi92)
                    default:
                        switch "hello" {
                        case "hello":
                            v3 = 2147483647
                        case "":
                            v2[""] = (fi94 & (fi92 / (((fi94 >> ((789) & 63))) | 1)))
                        case "foo bar":
                            if (false && false) {
                                if len(v1) > 0 {
                                    v1[0] = "hello"
                                }
                                fmt.Println(v18)
                                v22 -= 42
                            }
                        default:
                            if ((((("\n\t" + ("\n\t" + ("" + ("" + ("foo bar" + "foo bar"))))) + ("foo bar" + "foo bar")) + "a") + (("\n\t" + "foo bar") + ("a" + "\n\t"))) > "a") {
                                v95 := GenS6{F0: 9223372036854775807}
                                if ((1 != -128) && false) {
                                    var v96 uint16 = 2
                                    fmt.Println(v19, fi94, fi94)
                                    fmt.Println(v96, fi94, v19)
                                    _ = v96
                                }
                                if false {
                                    v97 := GenS7{F0: 1, F1: (((v18 + ((float32(0.0) - v18) * v18)) == (1.5 - float32(0.1))) || true), F2: 763}
                                    _ = v97
                                }
                                if true {
                                    switch 162 {
                                    case -25:
                                        v22 += v22
                                    }
                                    v19 = (^(v19))
                                } else {
                                    var v98 int8 = -1
                                    switch (fi92 - fi94) {
                                    case 9223372036854775807:
                                        if false {
                                            _ = len(v1)
                                            fmt.Println(v22)
                                        } else {
                                            var v99 uint16 = 0
                                            v98 = (v98 >> ((255) & 63))
                                            var v100 int = (v22 + fi92)
                                            _ = v100
                                            _ = v99
                                        }
                                    case 42:
                                        var v101 float64 = ((0.0 + (-2.25 + 0.1)) - 0.1)
                                        _ = v101
                                    }
                                    _ = v98
                                }
                                v22 += fi94
                                delete(v2, "foo bar")
                                _ = v95
                            } else {
                                var v102 int16 = int16(v22)
                                _ = v102
                            }
                        }
                    }
                case -736:
                    if (652 > 238) {
                        if false {
                            var v103 int16 = 997
                            v3 = v3
                            _ = v103
                        }
                        v22 += genFn0(("hello" + "a"), (v19 & (v19 / ((47) | 1))), "foo bar")
                        var v104 uint16 = 255
                        fmt.Println(fi92, v18, v22)
                        _ = v104
                    } else {
                        v105 := GenS4{F0: 1, F1: 2, F2: 98}
                        _ = v105
                    }
                default:
                    switch 127 {
                    case 0:
                        v3 = v3
                    case 128:
                        v106 := GenS8{F0: -1, F1: 42}
                        _ = v106
                    case -1:
                        var v107 int64 = -941
                        _ = v107
                    default:
                        if (!(100 != int8(v3))) {
                            v108 := GenS9{F0: (("a" + (((("a" + (("\n\t" + "foo bar") + "")) + ("foo bar" + "hello")) + "foo bar") + "a")) + "\n\t"), F1: 2}
                            _ = v108
                        } else {
                            var v109 int = (fi94 >> ((65535) & 63))
                            _ = v109
                        }
                    }
                }
                var v110 uint = 607
                _ = v110
            }
            fmt.Println(v18, v3, v22)
            _ = fi94
        }
        v111 := GenS2{F0: 100, F1: 4294967295, F2: 42}
        _ = fi92
        _ = v111
        _ = v93
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v83
}
