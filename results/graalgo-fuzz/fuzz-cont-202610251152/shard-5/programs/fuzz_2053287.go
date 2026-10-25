package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 int; F2 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16; F1 int8; F2 string }
func genFn0(p0 int64, p1 int64) uint32 {
    p0 = (^(p1))
    for fi7 := 0; fi7 < 18; fi7++ {
        p1 -= (p0 >> ((828) & 63))
        _ = fi7
    }
    switch ("a" + (((((("\n\t" + ("foo bar" + ("" + "foo bar"))) + "") + "") + "") + "a") + "foo bar")) {
    case "hello":
        switch -1 {
        case 1:
            var v8 uint64 = 341
            _ = v8
        default:
            v9 := []int{400}
            _ = v9
        }
    }
    v10 := GenS1{F0: 42, F1: 1}
    var v11 string = "foo bar"
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 426
}
func genFn1(p0 bool, p1 uint, p2 string) uint8 {
    var v12 int8 = 11
    v12 = (v12 % ((v12) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}

func main() {
    v1 := []uint32{2, 4294967295, 445, 0}
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    var v4 float64 = 1e10
    var v5 uint32 = 2147483647
    v6 := make(map[string]string)
    v4 = float64((1e10 * -2.25))
    var v13 int8 = 1
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            if true {
                if (false || true) {
                    fmt.Println(v4)
                }
                if (!(18446744073709551615 >= 2)) {
                    if false {
                        if (!(!(!(true && false)))) {
                            var v16 uint32 = 73
                            _ = v16
                        }
                        var v17 uint = 4294967295
                        _ = v17
                    }
                    var v18 int8 = (v13 % (((v13 ^ v13)) | 1))
                    _ = v18
                } else {
                    if (3.14159 < 0.1) {
                        v5 -= 0
                        if true {
                            var v19 int64 = 804
                            v20 := GenS2{F0: 42}
                            v21 := GenS2{F0: 1}
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            var v22 int64 = -400
                            var v23 int64 = int64(fi15)
                            if (genFn0(v22, v22) <= v5) {
                                v24 := GenS3{F0: 0, F1: 4294967295, F2: (^(v5))}
                                delete(v3, "hello")
                                _ = v24
                            } else {
                                v25 := GenS1{F0: 149, F1: -256}
                                if (!((false && (v5 != 79)) || ((!(-819 > 0)) && true))) {
                                    if (523 <= 2) {
                                        switch fi15 {
                                        case -2147483648:
                                            _ = len(v1)
                                        case 554:
                                            v5 -= genFn0(v23, v22)
                                        default:
                                            var v26 int = fi15
                                            _ = v26
                                        }
                                        if true {
                                            if false {
                                                _ = len(v6)
                                            }
                                        } else {
                                            v27 := GenS2{F0: 2}
                                            v28 := GenS4{F0: 1}
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                }
                                var v29 string = (((("hello" + "") + "") + "a") + "")
                                v23 = (v22 - v22)
                                v1 = append(v1, (v5 << ((145) & 63)))
                                _ = len(v2)
                                switch v29 {
                                case "\n\t":
                                    v30 := GenS5{F0: 0, F1: 128, F2: -2147483648}
                                    _ = v30
                                case "foo bar":
                                    v29 = (v29 + "\n\t")
                                }
                                _ = v25
                                _ = v29
                            }
                            if (1.5 < ((0.1 * (float32((v4 + v4)) - (0.1 * 0.1))) - -2.25)) {
                                fmt.Println(fi15)
                                var v31 int8 = v13
                                fmt.Println(v13, v22, fi14)
                                _ = v31
                            }
                            if (!(128 > 118)) {
                                v6["a"] = "a"
                            }
                            _ = v22
                            _ = v23
                        }
                    }
                    fmt.Println(v4)
                }
            } else {
                v13 -= (v13 % ((0) | 1))
                var v32 uint64 = 18446744073709551615
                _ = v32
            }
            _ = fi15
        }
        v13 = v13
        var v33 int8 = -30
        _ = fi14
        _ = v33
    }
    v34 := GenS1{F0: 1, F1: 32767}
    if false {
        fmt.Println(v13, v5, v13)
        if (("a" >= "hello") && (1 >= 2)) {
            v35 := GenS4{F0: 2147483647}
            v4 -= v4
            for fi36 := 0; fi36 < 7; fi36++ {
                for fi37 := 0; fi37 < 12; fi37++ {
                    var v38 float32 = ((float32(v4) / 1.0) - 0.1)
                    fmt.Println(v5)
                    _ = fi37
                    _ = v38
                }
                _ = fi36
            }
            _ = v35
        } else {
            var v39 uint8 = 191
            _ = v39
        }
        for fi40 := 0; fi40 < 13; fi40++ {
            var v41 int64 = -9223372036854775808
            _ = fi40
            _ = v41
        }
    } else {
        v6[("foo bar" + ((("" + (("\n\t" + "") + ("hello" + ""))) + ("a" + "a")) + "a"))] = "hello"
        for fi42 := 0; fi42 < 8; fi42++ {
            var v43 int32 = 887
            _ = fi42
            _ = v43
        }
    }
    var v44 float64 = ((3.14159 * float64(((1e10 - float32(1.5)) + float32(v4)))) * 1e10)
    if (4294967295 >= 642) {
        if true {
            switch 128 {
            case 128:
                if (42 > 186) {
                    _ = len(v1)
                    for fi45 := 0; fi45 < 3; fi45++ {
                        if (32767 >= 42) {
                            if len(v1) > 0 {
                                v1[0] = genFn0(-206, 0)
                            }
                            for fi46 := 0; fi46 < 7; fi46++ {
                                fmt.Println(fi45)
                                if false {
                                    v47 := GenS1{F0: 2, F1: int16(fi45)}
                                    if true {
                                        v48 := GenS5{F0: 873, F1: 2, F2: 593}
                                        v49 := GenS2{F0: 42}
                                        var v50 bool = true
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v51 := v6["hello"]
                                        _ = v51
                                    }
                                    fmt.Println(fi45, fi46, fi46)
                                    _ = v47
                                } else {
                                    var v52 uint32 = (v5 / (((v5 % ((2) | 1))) | 1))
                                    if (!(65535 < (v5 << ((255) & 63)))) {
                                        switch -1 {
                                        case 128:
                                            delete(v3, "\n\t")
                                        default:
                                            if false {
                                                if true {
                                                    var v53 uint16 = 758
                                                    var v54 uint8 = 127
                                                    if (!(!true)) {
                                                        fmt.Println(v54, v44, v4)
                                                        v44 += 1.5
                                                    } else {
                                                        if true {
                                                            v55 := GenS3{F0: 228, F1: 42, F2: 4294967295}
                                                            var v56 float32 = 0.1
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                        var v57 int64 = -9223372036854775808
                                                        _ = v57
                                                    }
                                                    _ = v53
                                                    _ = v54
                                                }
                                                var v58 uint64 = 255
                                                _ = v58
                                            }
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            var v59 float32 = (1e10 / float32((float64(1e10) + float64(3.14159))))
                                            _ = v59
                                        case "hello":
                                            if false {
                                                v60 := GenS2{F0: 2}
                                                _ = v60
                                            } else {
                                                fmt.Println(v52, fi46, v13)
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v61 uint32 = v5
                                            _ = v61
                                        case "\n\t":
                                            if true {
                                                var v62 float32 = 3.14159
                                                var v63 string = ("a" + (("hello" + "foo bar") + ""))
                                                fmt.Println(v44, v44, fi45)
                                                _ = v62
                                                _ = v63
                                            }
                                        default:
                                            fmt.Println(fi45)
                                        }
                                    } else {
                                        v64 := v3["hello"]
                                        switch "hello" {
                                        case "":
                                            switch fi46 {
                                            case 32767:
                                                _ = len(v3)
                                            }
                                        default:
                                            v5 -= (v52 / (((^(v5))) | 1))
                                        }
                                        v4 -= 1e10
                                        _ = v64
                                    }
                                    if true {
                                        v5 = v5
                                    }
                                    _ = v52
                                }
                                fmt.Println(v13, v44)
                                _ = fi46
                            }
                            delete(v3, "a")
                        } else {
                            fmt.Println(fi45)
                        }
                        var v65 uint8 = 2
                        var v66 string = ("foo bar" + ("\n\t" + ""))
                        _ = fi45
                        _ = v65
                        _ = v66
                    }
                } else {
                    v67 := GenS3{F0: 218, F1: 1, F2: (v5 << ((2) & 63))}
                    fmt.Println(v13)
                    _ = v67
                }
            default:
                fmt.Println(v44, v4, v44)
            }
            switch -134 {
            case -9223372036854775808:
                for fi68 := 0; fi68 < 2; fi68++ {
                    v69 := v6[""]
                    v70 := GenS6{F0: 42, F1: 32767}
                    var v71 int64 = -640
                    _ = fi68
                    _ = v69
                    _ = v70
                    _ = v71
                }
            }
        }
        fmt.Println(v44, v4, v13)
    }
    v5 -= v5
    var v72 string = "a"
    for fi73 := 0; fi73 < 14; fi73++ {
        switch fi73 {
        case 2147483647:
            v74 := GenS7{F0: 42}
            _ = v74
        case 159:
            var v75 int64 = 1
            _ = v75
        case 127:
            v72 = v72
        default:
            switch (fi73 / ((42) | 1)) {
            case 4294967295:
                if true {
                    for fi76 := 0; fi76 < 4; fi76++ {
                        var v77 int32 = -2147483648
                        _ = fi76
                        _ = v77
                    }
                    fmt.Println(v5, v72)
                } else {
                    for fi78 := 0; fi78 < 18; fi78++ {
                        fmt.Println(v5, v4, v13)
                        v72 = ""
                        _ = fi78
                    }
                    if true {
                        var v79 uint16 = 514
                        _ = v79
                    }
                    if (!(!((((((((float32(0.0) / 3.14159) / 1.0) / 1e10) + (-2.25 / 0.1)) * float32((v4 * (0.0 + (v4 - 0.1))))) - (0.0 / 0.1)) + 1.5) == (1e10 / 1e10)))) {
                        var v80 uint = 65535
                        v13 = v13
                        _ = v80
                    }
                }
            }
        }
        if ((!true) && false) {
            if len(v1) > 0 {
                v1[0] = v5
            }
        } else {
            for fi81 := 0; fi81 < 13; fi81++ {
                v5 += 4294967295
                var v82 uint = 402
                _ = fi81
                _ = v82
            }
        }
        v72 = (v72 + v72)
        _ = fi73
    }
    v72 = v72
    v72 = v72
    if true {
        switch (v72 + v72) {
        case "a":
            if false {
                var v83 uint32 = v5
                v13 += (v13 / ((v13) | 1))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v83
            } else {
                switch v72 {
                case "foo bar":
                    var v84 int16 = -1
                    _ = v84
                }
            }
        case "":
            v85 := GenS8{F0: -2147483648, F1: 2147483647, F2: 116}
            _ = v85
        case "hello":
            if (406 >= 42) {
                for fi86 := 0; fi86 < 14; fi86++ {
                    for fi87 := 0; fi87 < 12; fi87++ {
                        if false {
                            v88 := GenS9{F0: 0}
                            _ = v88
                        } else {
                            v44 -= (1.5 - (((v4 / v4) - 3.14159) * ((v44 + (v44 + v4)) - v44)))
                        }
                        _ = fi87
                    }
                    _ = fi86
                }
                if ((!false) || ((-1 < 9223372036854775807) && (4294967295 > 64))) {
                    var v89 uint8 = 1
                    if (v4 == 0.0) {
                        fmt.Println(v89, v13, v89)
                    }
                    switch 42 {
                    case 128:
                        v90 := GenS10{F0: 578}
                        _ = v90
                    case -968:
                        if false {
                            v5 = (v5 ^ 65535)
                        }
                    }
                    _ = v89
                }
                v13 = (v13 % ((100) | 1))
            } else {
                var v91 bool = false
                _ = v91
            }
        }
    } else {
        switch (v72 + "hello") {
        case "a":
            v44 = 1e10
        case "\n\t":
            v44 -= (v44 / ((((v4 + -2.25) * (v44 / (-2.25 - float64(0.1)))) + v44) * (v4 / float64(float32(v4)))))
        case "foo bar":
            if false {
                var v92 int8 = v13
                for fi93 := 0; fi93 < 2; fi93++ {
                    fmt.Println(fi93)
                    for fi94 := 0; fi94 < 16; fi94++ {
                        v44 = 3.14159
                        var v95 int64 = 9223372036854775807
                        v95 += (^(v95))
                        _ = fi94
                        _ = v95
                    }
                    _ = fi93
                }
                _ = v92
            }
        default:
            var v96 uint16 = 1
            _ = v96
        }
        v72 = v72
    }
    v97 := GenS11{F0: 2, F1: (v13 >> ((2) & 63)), F2: (v72 + ("\n\t" + ((("foo bar" + (v72 + v72)) + v72) + v72)))}
    v5 = v5
    switch 32768 {
    case -9223372036854775808:
        if (!true) {
            var v98 uint8 = 127
            var v99 int = 656
            _ = v98
            _ = v99
        }
    case -1:
        var v100 bool = false
        _ = v100
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v72
    _ = v97
}
