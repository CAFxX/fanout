package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 string; F1 uint64; F2 int64 }
type GenS5 struct { F0 int; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint8; F1 uint32; F2 string }
type GenS7 struct { F0 string; F1 uint32; F2 int }
type GenS8 struct { F0 bool }
func genFn0() int {
    if true {
        v4 := []int{4294967295, 42, -1}
        if (true || false) {
            var v5 bool = (2 <= 0)
            _ = v5
        } else {
            var v6 int16 = 32767
            if false {
                v7 := GenS1{F0: "a", F1: 2}
                v8 := GenS1{F0: ("hello" + (("hello" + "hello") + ("\n\t" + "hello"))), F1: 303}
                _ = v7
                _ = v8
            }
            for fi9 := 0; fi9 < 18; fi9++ {
                v6 -= 42
                v6 = v6
                var v10 uint32 = 470
                _ = fi9
                _ = v10
            }
            _ = v6
        }
        switch "hello" {
        case "":
            if false {
                if len(v4) > 0 {
                    v4[0] = 543
                }
            }
        case "a":
            v11 := GenS1{F0: "", F1: 0}
            _ = v11
        case "\n\t":
            var v12 int = 1
            _ = v12
        default:
            if true {
                if (!true) {
                    var v13 uint = 65535
                    _ = v13
                } else {
                    var v14 int64 = 484
                    if true {
                        fmt.Println(v14)
                        var v15 uint = 4294967295
                        _ = v15
                    } else {
                        var v16 uint = 990
                        v14 = 1
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 15; fi17++ {
                        switch (fi17 / ((fi17) | 1)) {
                        case 181:
                            v14 = (v14 * v14)
                        case 127:
                            v18 := GenS2{F0: 2, F1: v14, F2: -2147483648}
                            _ = v18
                        }
                        _ = fi17
                    }
                    _ = v14
                }
                for fi19 := 0; fi19 < 20; fi19++ {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 string = ""
                        if (false || (!(65535 > 0))) {
                            fmt.Println(fi19)
                            if (-82 <= 0) {
                                v21 = v21
                                var v22 uint32 = 602
                                _ = v22
                            }
                            v23 := GenS1{F0: "hello", F1: 4294967295}
                            _ = v23
                        } else {
                            v21 = ((((((v21 + "") + ((v21 + v21) + v21)) + ("a" + (((v21 + (v21 + v21)) + v21) + "foo bar"))) + v21) + v21) + (("foo bar" + "\n\t") + ("hello" + (((v21 + v21) + "a") + (v21 + "")))))
                            v21 = "foo bar"
                            v21 = v21
                            if (42 != 213) {
                                if false {
                                    if true {
                                        switch (((((v21 + (v21 + "a")) + (v21 + "a")) + v21) + "\n\t") + v21) {
                                        case "a":
                                            var v24 float64 = -2.25
                                            _ = v24
                                        case "":
                                            v21 = v21
                                        case "hello":
                                            switch v21 {
                                            case "foo bar":
                                                var v25 uint16 = 861
                                                _ = v25
                                            case "\n\t":
                                                var v26 int16 = 32767
                                                _ = v26
                                            default:
                                                fmt.Println(fi20)
                                            }
                                            v21 = v21
                                            if (false || ((293 == 678) && false)) {
                                                var v27 float32 = 0.0
                                                var v28 uint32 = 0
                                                fmt.Println(v21)
                                                _ = v27
                                                _ = v28
                                            }
                                        }
                                        if false {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        switch ("\n\t" + (v21 + v21)) {
                                        case "foo bar":
                                            if (false && true) {
                                                var v29 float32 = (1e10 + (3.14159 - (0.0 - ((1e10 - 3.14159) + (1e10 / 1e10)))))
                                                var v30 float32 = 0.0
                                                v31 := GenS3{F0: 2, F1: ((true && false) || false)}
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                        case "a":
                                            var v32 string = v21
                                            _ = v32
                                        case "hello":
                                            var v33 bool = false
                                            _ = v33
                                        default:
                                            if (1.5 < 1.5) {
                                                switch 127 {
                                                case -1:
                                                    var v34 int64 = -9223372036854775808
                                                    _ = v34
                                                }
                                                var v35 float32 = (1e10 * -2.25)
                                                _ = v35
                                            }
                                        }
                                    }
                                }
                                fmt.Println(fi19)
                                v21 = (v21 + "foo bar")
                            }
                            v21 = "foo bar"
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
                if true {
                    var v36 int16 = 42
                    _ = v36
                } else {
                    var v37 uint8 = 42
                    v37 = (v37 >> ((18446744073709551615) & 63))
                    v38 := GenS2{F0: 42, F1: -473, F2: 0}
                    _ = v37
                    _ = v38
                }
            } else {
                v39 := GenS2{F0: 1, F1: 515, F2: -48}
                var v40 int = 2147483647
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v39
                _ = v40
            }
        }
        _ = v4
    } else {
        var v41 uint16 = 255
        _ = v41
    }
    v42 := GenS4{F0: "\n\t", F1: 892, F2: 836}
    var v43 float32 = 1e10
    v44 := GenS1{F0: "", F1: 18446744073709551615}
    _ = v42
    _ = v43
    _ = v44
    return 128
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 2147483647
    var v3 float64 = (3.14159 + 0.0)
    v2 = v2
    if (2147483648 != 610) {
        for fi45 := 0; fi45 < 17; fi45++ {
            if true {
                for fi46 := 0; fi46 < 20; fi46++ {
                    var v47 uint32 = 2147483647
                    var v48 uint64 = 0
                    if (!(-2.25 == 0.0)) {
                        v3 -= 3.14159
                        fmt.Println(fi46, fi46)
                        switch "" {
                        case "\n\t":
                            var v49 bool = true
                            _ = v49
                        default:
                            var v50 int = fi45
                            _ = v50
                        }
                    } else {
                        fmt.Println(v3)
                        v51 := []int{(-(fi46)), (fi45 >> ((705) & 63))}
                        _ = v51
                    }
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                switch "foo bar" {
                case "foo bar":
                    v52 := GenS5{F0: fi45, F1: 0, F2: -1}
                    _ = v52
                default:
                    var v53 uint16 = 1
                    _ = v53
                }
            }
            _ = fi45
        }
        var v54 string = ("a" + "foo bar")
        _ = v54
    }
    var v55 int16 = 322
    var v56 uint8 = 255
    if (!(0 == -110)) {
        if ("foo bar" == "\n\t") {
            for fi57 := 0; fi57 < 1; fi57++ {
                var v58 float64 = 1e10
                switch "\n\t" {
                case "hello":
                    if true {
                        for fi59 := 0; fi59 < 9; fi59++ {
                            v60 := []int{(-(fi57)), genFn0(), fi57, 245}
                            _ = fi59
                            _ = v60
                        }
                        fmt.Println(v3, v2, fi57)
                        var v61 uint16 = 255
                        _ = v61
                    }
                default:
                    v62 := []int{-623, 2147483647}
                    _ = v62
                }
                v55 += v55
                _ = fi57
                _ = v58
            }
        } else {
            for fi63 := 0; fi63 < 11; fi63++ {
                v64 := []int{1, fi63, 128, (-(fi63))}
                _ = fi63
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 16; fi65++ {
            for fi66 := 0; fi66 < 13; fi66++ {
                v67 := GenS3{F0: 42, F1: ((true || false) && true)}
                if ((!(-1 < -1)) && (!(true && (!true)))) {
                    v68 := []int{genFn0(), genFn0(), 32767}
                    _ = v68
                } else {
                    var v69 int16 = v55
                    _ = v69
                }
                v3 += v3
                v3 += v3
                v3 = 1e10
                _ = fi66
                _ = v67
            }
            v1["\n\t"] = true
            for fi70 := 0; fi70 < 17; fi70++ {
                v71 := GenS2{F0: 18446744073709551615, F1: 715, F2: v2}
                var v72 uint64 = 318
                _ = fi70
                _ = v71
                _ = v72
            }
            _ = fi65
        }
        var v73 string = "\n\t"
        _ = v73
    } else {
        for fi74 := 0; fi74 < 4; fi74++ {
            for fi75 := 0; fi75 < 14; fi75++ {
                v2 = (^(v2))
                _ = fi75
            }
            for fi76 := 0; fi76 < 4; fi76++ {
                v77 := []int{genFn0(), (^(fi74))}
                if false {
                    var v78 int32 = -556
                    if false {
                        v78 = (^(v78))
                        switch (("\n\t" + (("" + "foo bar") + "foo bar")) + ("\n\t" + "\n\t")) {
                        case "a":
                            v78 -= v2
                        case "":
                            fmt.Println(v55, fi76, fi76)
                        default:
                            _ = len(v1)
                        }
                    }
                    v78 += v2
                    v56 = v56
                    v77 = append(v77, (fi74 + fi74))
                    _ = v78
                } else {
                    v3 = float64((-2.25 * (3.14159 * float32(v3))))
                    switch (fi76 & (fi74 % ((-43) | 1))) {
                    case 2147483647:
                        if (127 >= (v56 / (((v56 >> ((1) & 63))) | 1))) {
                            if (!((v2 % ((v2) | 1)) > (^(v2)))) {
                                v55 += (^(v55))
                                var v79 string = (("" + "a") + ("hello" + ""))
                                _ = v79
                            } else {
                                fmt.Println(fi76)
                                var v80 int = fi76
                                v81 := GenS3{F0: 1, F1: ((0 >= 0) && false)}
                                _ = v80
                                _ = v81
                            }
                            v82 := GenS4{F0: ("hello" + ("" + "")), F1: 18446744073709551615, F2: int64(v2)}
                            _ = v82
                        }
                    default:
                        var v83 uint8 = v56
                        _ = v83
                    }
                    var v84 uint16 = 65535
                    _ = v84
                }
                if false {
                    v3 = -2.25
                    var v85 bool = (2147483648 == 493)
                    _ = v85
                }
                _ = fi76
                _ = v77
            }
            if true {
                if false {
                    fmt.Println(v3, v2, fi74)
                    v56 = (v56 & (v56 + (v56 / ((v56) | 1))))
                    v3 -= v3
                }
                var v86 int16 = v55
                fmt.Println(v55)
                _ = v86
            } else {
                v2 = v2
            }
            _ = fi74
        }
        for fi87 := 0; fi87 < 16; fi87++ {
            for fi88 := 0; fi88 < 3; fi88++ {
                v89 := v1["foo bar"]
                if v89 {
                    v2 -= int32(v56)
                    v2 = v2
                    switch (("hello" + (("a" + "") + "hello")) + (("hello" + "\n\t") + "hello")) {
                    case "\n\t":
                        v2 -= (v2 >> ((1) & 63))
                    }
                } else {
                    v2 -= v2
                    v2 = 42
                }
                v90 := GenS5{F0: -466, F1: 240, F2: 127}
                _ = fi88
                _ = v89
                _ = v90
            }
            if ((v55 & int16(v2)) >= v55) {
                var v91 bool = (!true)
                v92 := v1["hello"]
                v92 = (false || (true || v91))
                _ = v91
                _ = v92
            } else {
                v56 += (v56 | (v56 + (v56 / ((v56) | 1))))
                var v93 int64 = 0
                _ = v93
            }
            v94 := GenS6{F0: (v56 + (v56 & v56)), F1: 65535, F2: "foo bar"}
            _ = fi87
            _ = v94
        }
    }
    switch 0 {
    case 9223372036854775807:
        fmt.Println(v55, v3, v55)
    default:
        if true {
            var v95 int32 = (v2 / ((-1) | 1))
            _ = v95
        }
    }
    v96 := GenS5{F0: 4294967295, F1: v56, F2: 117}
    v56 -= (^(v56))
    delete(v1, "a")
    if (42 <= -128) {
        var v97 int = 895
        _ = v97
    } else {
        fmt.Println(v3, v2)
        v98 := []int{-9223372036854775808}
        var v99 int8 = 54
        _ = v98
        _ = v99
    }
    if true {
        var v100 int8 = -37
        _ = v100
    }
    v56 = v56
    switch genFn0() {
    case 42:
        if ((v56 / ((0) | 1)) >= v56) {
            fmt.Println(v56)
            v101 := []int{1}
            _ = v101
        } else {
            v102 := GenS5{F0: 32768, F1: v56, F2: 14}
            _ = v102
        }
    }
    if (128 == -1) {
        v103 := GenS7{F0: "\n\t", F1: 65535, F2: -868}
        switch 1 {
        case 9223372036854775807:
            var v104 int32 = v2
            _ = v104
        case 42:
            v105 := []int{2147483647}
            _ = v105
        case 127:
            v3 -= v3
        }
        fmt.Println(v3, v3)
        _ = v103
    }
    if (-2147483648 < 128) {
        for fi106 := 0; fi106 < 18; fi106++ {
            if true {
                switch fi106 {
                case 1:
                    if true {
                        for fi107 := 0; fi107 < 11; fi107++ {
                            v55 = (v55 + (^(v55)))
                            if ((543 != 1) || (1 >= 65535)) {
                                var v108 float64 = 0.0
                                v1[""] = ((!true) && ((-(v55)) == 0))
                                var v109 int32 = v2
                                _ = v108
                                _ = v109
                            }
                            _ = fi107
                        }
                        switch ("\n\t" + "foo bar") {
                        case "\n\t":
                            fmt.Println(v55, v55, v55)
                        default:
                            fmt.Println(v56)
                        }
                    }
                }
                switch (("foo bar" + "\n\t") + "foo bar") {
                case "a":
                    v110 := GenS1{F0: "a", F1: 364}
                    _ = v110
                case "\n\t":
                    v56 -= v56
                }
                v3 += v3
            } else {
                if true {
                    v55 = v55
                    v3 = v3
                    v3 += v3
                }
            }
            _ = fi106
        }
        v3 -= v3
    }
    v111 := GenS8{F0: (!true)}
    var v112 uint8 = 192
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v112:", v112)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v111
    _ = v112
    _ = v2
    _ = v3
    _ = v55
    _ = v56
    _ = v96
}
