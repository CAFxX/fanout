package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint; F2 int32 }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 uint64; F2 uint }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0(p0 int16) int {
    v6 := []int{}
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 bool = (-917 <= fi7)
        for fi9 := 0; fi9 < 15; fi9++ {
            p0 -= (-(p0))
            p0 = (p0 ^ (p0 << ((235) & 63)))
            _ = fi9
        }
        fmt.Println(fi7, p0)
        _ = fi7
        _ = v8
    }
    if true {
        if ((42 >= 450) || false) {
            fmt.Println(p0)
            v10 := GenS1{F0: "", F1: 42}
            _ = v10
        }
        if (42 <= 127) {
            v6 = append(v6, 0)
            switch 2147483647 {
            case -259:
                p0 += (p0 ^ (p0 - -771))
            case 4294967295:
                for fi11 := 0; fi11 < 9; fi11++ {
                    for fi12 := 0; fi12 < 14; fi12++ {
                        var v13 uint8 = 242
                        if ((v13 * v13) <= v13) {
                            switch -2147483648 {
                            case 606:
                                var v14 int = (fi11 * 32768)
                                _ = v14
                            case -111:
                                v15 := GenS2{F0: 1, F1: 42}
                                _ = v15
                            case -1:
                                v13 -= (v13 * (^(v13)))
                            }
                            switch (fi12 - (fi12 ^ fi11)) {
                            case 32768:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            case -1:
                                var v16 int8 = 89
                                _ = v16
                            case 1:
                                switch fi11 {
                                case 32768:
                                    fmt.Println(fi11, p0, p0)
                                case -1:
                                    v17 := GenS3{F0: 2147483647}
                                    _ = v17
                                default:
                                    fmt.Println(v13)
                                }
                            default:
                                if len(v6) > 0 {
                                    v6[0] = 164
                                }
                            }
                            v13 = 1
                            if false {
                                switch ((("\n\t" + "\n\t") + (("" + "a") + "foo bar")) + "\n\t") {
                                case "\n\t":
                                    v13 += 127
                                }
                                fmt.Println(fi11)
                                switch (-(fi12)) {
                                case 548:
                                    if len(v6) > 0 {
                                        v6[0] = 9223372036854775807
                                    }
                                }
                            }
                        } else {
                            var v18 int8 = -81
                            var v19 int32 = int32(fi12)
                            _ = v18
                            _ = v19
                        }
                        _ = fi12
                        _ = v13
                    }
                    fmt.Println(fi11)
                    if len(v6) > 0 {
                        v6[0] = fi11
                    }
                    _ = fi11
                }
            case -24:
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 bool = false
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 string = "foo bar"
                        _ = len(v6)
                        v23 = v23
                        _ = fi22
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
            default:
                v24 := GenS4{F0: 18446744073709551615}
                _ = v24
            }
            var v25 string = (("" + ("foo bar" + "hello")) + "foo bar")
            _ = v25
        }
        if true {
            if true {
                var v26 uint64 = 1
                v6 = append(v6, -544)
                var v27 uint16 = 42
                _ = v26
                _ = v27
            } else {
                fmt.Println(p0)
            }
            switch "" {
            case "hello":
                _ = len(v6)
            case "":
                if ((1e10 - (1.5 * 0.0)) < float64(1.5)) {
                    var v28 uint8 = 88
                    _ = len(v6)
                    _ = v28
                }
            case "a":
                for fi29 := 0; fi29 < 9; fi29++ {
                    if true {
                        v6 = append(v6, fi29)
                        v6 = append(v6, 2147483647)
                        for fi30 := 0; fi30 < 5; fi30++ {
                            p0 -= p0
                            v31 := GenS5{F0: ("\n\t" + ""), F1: 4294967295, F2: 1}
                            v32 := GenS4{F0: 255}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    } else {
                        var v33 int16 = p0
                        _ = v33
                    }
                    p0 += p0
                    _ = fi29
                }
            default:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        } else {
            var v34 int64 = 422
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 uint = 1
        _ = fi35
        _ = v36
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 42
}
func genFn1() uint8 {
    if (2 != 2) {
        v37 := []int{0, -1, 9223372036854775807, 42}
        _ = v37
    } else {
        var v38 uint16 = 738
        _ = v38
    }
    var v39 string = ("hello" + "foo bar")
    v40 := GenS6{F0: 2, F1: 663, F2: -668}
    v41 := GenS4{F0: 293}
    _ = v39
    _ = v40
    _ = v41
    return 2
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{"hello", ("\n\t" + "foo bar"), "foo bar"}
    v3 := make(map[uint32]bool)
    var v4 int32 = 42
    v5 := make(map[string]string)
    if (!false) {
        if false {
            var v42 bool = false
            switch "hello" {
            case "hello":
                v43 := GenS2{F0: 42, F1: 0}
                _ = v43
            }
            var v44 uint32 = 2147483647
            _ = v42
            _ = v44
        }
        for fi45 := 0; fi45 < 20; fi45++ {
            v46 := GenS7{F0: -23, F1: 2}
            _ = fi45
            _ = v46
        }
        for fi47 := 0; fi47 < 3; fi47++ {
            v48 := GenS1{F0: ("\n\t" + (("foo bar" + ((("a" + "foo bar") + ("" + "\n\t")) + "a")) + "\n\t")), F1: (v1 % (((v1 + (v1 * v1))) | 1))}
            v4 = (v4 << ((65535) & 63))
            _ = fi47
            _ = v48
        }
    } else {
        var v49 float64 = -2.25
        v1 = (^(v1))
        switch "hello" {
        case "a":
            for fi50 := 0; fi50 < 14; fi50++ {
                v51 := v5[""]
                fmt.Println(v4, v49)
                if ((v1 ^ v1) == 8) {
                    var v52 uint = 255
                    v2 = append(v2, v51)
                    v51 = v51
                    _ = v52
                }
                _ = fi50
                _ = v51
            }
        }
        _ = v49
    }
    var v53 int16 = -1
    v4 = v4
    var v54 float64 = 0.0
    switch "foo bar" {
    case "a":
        v55 := GenS6{F0: 255, F1: 18446744073709551615, F2: 296}
        _ = v55
    case "":
        switch genFn0(v53) {
        case 2147483647:
            if (0.1 > 1e10) {
                var v56 uint = 65535
                switch "a" {
                case "hello":
                    v57 := GenS2{F0: 879, F1: 2}
                    _ = v57
                case "":
                    var v58 int64 = v1
                    _ = v58
                default:
                    _ = len(v5)
                }
                switch 118 {
                case 1:
                    v1 += (v1 + 0)
                case -24:
                    for fi59 := 0; fi59 < 16; fi59++ {
                        _ = len(v3)
                        _ = fi59
                    }
                default:
                    for fi60 := 0; fi60 < 20; fi60++ {
                        v61 := GenS8{F0: ""}
                        _ = len(v3)
                        v53 += (v53 + 42)
                        _ = fi60
                        _ = v61
                    }
                }
                _ = v56
            }
        case 4294967295:
            v54 = 0.1
        case 42:
            _ = len(v2)
        default:
            var v62 int8 = -122
            _ = v62
        }
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = ""
        }
    }
    var v63 int8 = 0
    v64 := GenS9{F0: "foo bar", F1: 0, F2: 2}
    switch ((("hello" + "hello") + ("hello" + "a")) + "") {
    case "a":
        var v65 int32 = (v4 - (v4 - v4))
        _ = v65
    case "":
        var v66 float32 = (3.14159 * 0.1)
        _ = v66
    case "hello":
        for fi67 := 0; fi67 < 4; fi67++ {
            if false {
                var v68 uint16 = 32767
                _ = v68
            }
            var v69 uint = 0
            for fi70 := 0; fi70 < 19; fi70++ {
                var v71 int = (fi70 - (fi67 / (((fi70 % ((genFn0(-32768)) | 1))) | 1)))
                _ = fi70
                _ = v71
            }
            _ = fi67
            _ = v69
        }
    }
    if true {
        switch -9223372036854775808 {
        case 32767:
            var v72 int8 = -1
            _ = v72
        case -1:
            if (349 <= 904) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
                switch 2147483647 {
                case 833:
                    v63 = v63
                case 9223372036854775807:
                    var v73 uint64 = 0
                    _ = v73
                }
            } else {
                fmt.Println(v53, v63)
                if true {
                    v74 := GenS10{F0: 42, F1: 140}
                    var v75 string = ("hello" + (("" + "\n\t") + "hello"))
                    _ = v74
                    _ = v75
                } else {
                    var v76 bool = true
                    var v77 int32 = v4
                    var v78 float32 = (-2.25 - 1e10)
                    _ = v76
                    _ = v77
                    _ = v78
                }
            }
        case 32768:
            if true {
                var v79 int8 = (v63 + -128)
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = v79
            } else {
                var v80 uint32 = 2
                _ = v80
            }
        default:
            if len(v2) > 0 {
                v2[0] = ("hello" + "\n\t")
            }
        }
        for fi81 := 0; fi81 < 1; fi81++ {
            var v82 int64 = (v1 >> ((605) & 63))
            for fi83 := 0; fi83 < 15; fi83++ {
                fmt.Println(v1, v1, fi81)
                var v84 uint = 902
                var v85 int8 = 100
                _ = fi83
                _ = v84
                _ = v85
            }
            _ = fi81
            _ = v82
        }
    } else {
        for fi86 := 0; fi86 < 18; fi86++ {
            v54 = v54
            v87 := GenS11{F0: 289, F1: ((v63 > v63) && true), F2: 41}
            _ = fi86
            _ = v87
        }
        fmt.Println(v53, v63, v54)
        var v88 int32 = (v4 & -2147483648)
        _ = v88
    }
    fmt.Println(v4)
    var v89 float32 = float32(v54)
    fmt.Println(v89)
    var v90 int = 2147483647
    switch v90 {
    case -588:
        switch genFn0(-498) {
        case 0:
            if false {
                v63 = v63
            }
        default:
            v53 = -32768
        }
    default:
        if false {
            if false {
                v91 := GenS4{F0: 0}
                if true {
                    var v92 uint8 = 1
                    v92 = v92
                    _ = v92
                } else {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    v89 -= float32(v54)
                }
                _ = v91
            }
        } else {
            var v93 int16 = v53
            if (1 == -32768) {
                v4 = 42
                switch (v90 - (v90 * v90)) {
                case -542:
                    if ((v53 % ((0) | 1)) == 1000) {
                        if true {
                            for fi94 := 0; fi94 < 5; fi94++ {
                                v53 -= (v53 | v93)
                                fmt.Println(v1, v53, v53)
                                _ = fi94
                            }
                            for fi95 := 0; fi95 < 13; fi95++ {
                                var v96 bool = ((v63 - int8(v1)) < -1)
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    if false {
                                        var v97 float32 = -2.25
                                        var v98 bool = (0 > 65535)
                                        _ = v97
                                        _ = v98
                                    }
                                case "\n\t":
                                    if false {
                                        fmt.Println(v54)
                                    } else {
                                        v99 := v5[""]
                                        v100 := GenS11{F0: 0, F1: v96, F2: 255}
                                        if (v96 && v96) {
                                            var v101 int32 = 514
                                            _ = v101
                                        }
                                        _ = v100
                                        _ = v99
                                    }
                                }
                                _ = fi95
                                _ = v96
                            }
                            v93 -= (v53 ^ (v53 % ((v93) | 1)))
                        } else {
                            if false {
                                fmt.Println(v54, v63, v1)
                            }
                            var v102 int = v90
                            _ = v102
                        }
                        if len(v2) > 0 {
                            v2[0] = ("a" + "foo bar")
                        }
                        fmt.Println(v90)
                    }
                case 32768:
                    if true {
                        v90 = v90
                        v103 := v3[65535]
                        _ = v103
                    }
                default:
                    v4 = v4
                }
                var v104 uint32 = 748
                _ = v104
            }
            v105 := v3[4294967295]
            _ = v105
            _ = v93
        }
    }
    var v106 float32 = float32(((v54 + v54) * (v54 - (1.5 - 1.5))))
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v63:", v63)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v63
    _ = v64
    _ = v89
    _ = v90
}
