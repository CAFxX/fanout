package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int; F1 int64; F2 int }
type GenS7 struct { F0 int16; F1 string; F2 string }
func genFn0() int {
    var v4 uint64 = 1
    for fi5 := 0; fi5 < 19; fi5++ {
        if (!true) {
            fmt.Println(fi5, fi5)
            if (-164 < 0) {
                v4 += v4
            }
            fmt.Println(v4, v4)
        }
        v4 += (v4 + (v4 / ((1) | 1)))
        if ((((-(fi5)) >= fi5) && (58 > 65535)) && (0 < 579)) {
            v6 := []int{(fi5 / ((fi5) | 1)), (fi5 & (fi5 + fi5)), (fi5 ^ fi5)}
            if (float32(3.14159) != 1e10) {
                for fi7 := 0; fi7 < 5; fi7++ {
                    fmt.Println(v4, fi5, v4)
                    if (false && (100 > -63)) {
                        if (0 > -32768) {
                            fmt.Println(fi5, fi5, fi7)
                        }
                    } else {
                        v4 += v4
                        var v8 string = ((("a" + "") + ("" + "")) + "hello")
                        _ = v8
                    }
                    v6 = append(v6, -779)
                    _ = fi7
                }
                var v9 float64 = 3.14159
                _ = v9
            } else {
                v4 += v4
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 uint8 = 128
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v4, v4)
            }
            v12 := GenS1{F0: fi5, F1: ("a" + ("hello" + "hello"))}
            _ = v12
            _ = v6
        } else {
            switch "\n\t" {
            case "foo bar":
                v13 := []int{}
                _ = v13
            }
            if (32767 <= 1) {
                v4 -= (^(v4))
                fmt.Println(fi5, fi5)
                var v14 float32 = (1.5 / 3.14159)
                _ = v14
            } else {
                var v15 int8 = 127
                _ = v15
            }
        }
        _ = fi5
    }
    v16 := []int{128}
    _ = v16
    _ = v4
    return 127
}
func genFn1() uint32 {
    if false {
        var v17 int = 127
        var v18 uint = 1
        var v19 int64 = 0
        _ = v17
        _ = v18
        _ = v19
    } else {
        var v20 int32 = 1
        _ = v20
    }
    switch "hello" {
    case "a":
        v21 := []int{-9223372036854775808, -2147483648, 32768}
        _ = v21
    case "":
        var v22 string = (("hello" + "\n\t") + "\n\t")
        _ = v22
    }
    if (false && (915 == 542)) {
        var v23 int = -190
        fmt.Println(v23)
        _ = v23
    }
    return 42
}

func main() {
    var v1 float32 = float32(float64(3.14159))
    var v2 int16 = 0
    v3 := make(map[uint32]bool)
    switch ("" + "a") {
    case "":
        if (true && (229 >= 0)) {
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 int16 = 192
                v26 := []int{fi24, (fi24 / ((2147483647) | 1)), (-(fi24)), (fi24 >> ((0) & 63))}
                var v27 int16 = (^(v25))
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            var v28 uint8 = 128
            v29 := []int{32767}
            _ = v28
            _ = v29
        } else {
            if (false && true) {
                v2 += v2
                v1 -= (3.14159 + ((v1 + ((v1 - (v1 * v1)) + v1)) - v1))
                v1 -= 3.14159
            }
        }
    default:
        v1 -= float32(float64(v1))
    }
    v1 = ((v1 + 1.5) - float32((0.0 + 0.1)))
    for fi30 := 0; fi30 < 18; fi30++ {
        v1 += 0.0
        var v31 string = ("\n\t" + "hello")
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 int8 = -1
            var v35 int = 9223372036854775807
            delete(v3, 2147483647)
            switch "\n\t" {
            case "\n\t":
                v2 += (v2 >> ((399) & 63))
            }
            fmt.Println(fi33, v2, v2)
            _ = fi33
            _ = v34
            _ = v35
        }
        fmt.Println(v1)
        _ = fi32
    }
    v36 := GenS1{F0: 32768, F1: ("a" + ("hello" + "hello"))}
    for fi37 := 0; fi37 < 16; fi37++ {
        for fi38 := 0; fi38 < 18; fi38++ {
            var v39 int16 = 0
            var v40 uint32 = 2147483648
            v39 = v2
            v2 -= (-(v2))
            var v41 uint64 = 1
            v39 = v2
            if (128 > 1) {
                var v42 float64 = 3.14159
                v39 -= -32768
                _ = v42
            }
            if ((v39 / ((1) | 1)) < (v2 - v39)) {
                v1 -= float32(float64(0.0))
            }
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
        for fi43 := 0; fi43 < 20; fi43++ {
            fmt.Println(v1)
            fmt.Println(v1, v1, fi43)
            _ = fi43
        }
        _ = fi37
    }
    v44 := []int{32767, 9223372036854775807, -716}
    for fi45 := 0; fi45 < 8; fi45++ {
        for fi46 := 0; fi46 < 2; fi46++ {
            if true {
                var v47 int64 = 42
                switch fi46 {
                case -1:
                    switch 2147483647 {
                    case 32767:
                        var v48 float32 = (((0.0 - v1) + (v1 * float32(-2.25))) - float32(0.0))
                        _ = v48
                    default:
                        if len(v44) >= 2 {
                            v44 = v44[1:len(v44)]
                        }
                    }
                default:
                    var v49 string = (("\n\t" + ("\n\t" + "\n\t")) + (("\n\t" + "hello") + ""))
                    _ = v49
                }
                _ = v47
            } else {
                v2 = (v2 ^ 32767)
                var v50 int = (fi45 / (((fi46 | genFn0())) | 1))
                v51 := GenS2{F0: 127, F1: (v2 >> ((65535) & 63))}
                _ = v50
                _ = v51
            }
            switch "a" {
            case "\n\t":
                if (!false) {
                    fmt.Println(fi45, v2)
                    var v52 int8 = 7
                    _ = v52
                } else {
                    if false {
                        if false {
                            switch (fi46 - (-(fi46))) {
                            case -1:
                                v53 := GenS2{F0: 80, F1: v2}
                                _ = v53
                            case 42:
                                if true {
                                    v2 = -583
                                    switch (("a" + ("hello" + ("foo bar" + "hello"))) + "") {
                                    case "a":
                                        var v54 int32 = 2147483647
                                        _ = v54
                                    }
                                } else {
                                    var v55 int16 = 0
                                    if ((true && false) || false) {
                                        fmt.Println(fi46, fi45)
                                    } else {
                                        v56 := v3[genFn1()]
                                        if ((!(0 >= 722)) && v56) {
                                            delete(v3, genFn1())
                                            v57 := GenS3{F0: v56}
                                            _ = v57
                                        } else {
                                            _ = len(v44)
                                            if (!(v56 || true)) {
                                                switch "foo bar" {
                                                case "hello":
                                                    if false {
                                                        var v58 bool = v56
                                                        v2 += v2
                                                        fmt.Println(v1, v58, v1)
                                                        _ = v58
                                                    }
                                                case "\n\t":
                                                    if false {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            fmt.Println(fi46, fi46, v2)
                                                        case "hello":
                                                            var v59 float32 = v1
                                                            _ = v59
                                                        default:
                                                            v60 := GenS4{F0: 65535}
                                                            _ = v60
                                                        }
                                                    }
                                                default:
                                                    var v61 bool = (((2147483647 <= 2) && false) || (338 > 1))
                                                    _ = v61
                                                }
                                                _ = len(v44)
                                                v2 -= v2
                                            } else {
                                                if v56 {
                                                    v62 := v3[4294967295]
                                                    v55 -= (v55 & v55)
                                                    v56 = (!v62)
                                                    _ = v62
                                                }
                                                var v63 string = ""
                                                var v64 int8 = -128
                                                _ = v63
                                                _ = v64
                                            }
                                            v56 = v56
                                            var v65 int32 = 42
                                            _ = v65
                                        }
                                        _ = v56
                                    }
                                    _ = v55
                                }
                            case 1:
                                var v66 uint = 654
                                _ = v66
                            }
                            v2 = (v2 - 573)
                        }
                        _ = len(v44)
                        v1 = -2.25
                    }
                }
            case "hello":
                v67 := GenS2{F0: -57, F1: -1}
                _ = v67
            default:
                if (42 > 4294967295) {
                    var v68 uint16 = 693
                    v1 += 3.14159
                    _ = v68
                } else {
                    switch (fi46 ^ genFn0()) {
                    case 128:
                        v1 -= 0.0
                    case -9223372036854775808:
                        v69 := v3[784]
                        _ = v69
                    case 2147483647:
                        switch (fi45 / (((fi46 + (-(fi45)))) | 1)) {
                        case 1:
                            if (true || (true && (!false))) {
                                v1 = float32(-2.25)
                            }
                        }
                    }
                    v2 = (v2 % ((-32768) | 1))
                }
            }
            _ = fi46
        }
        var v70 float32 = (v1 / ((float32(1.5) * float32((3.14159 - 3.14159))) + 1.5))
        _ = fi45
        _ = v70
    }
    v71 := GenS4{F0: 65535}
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    v2 = 0
    if (true && false) {
        v72 := GenS5{F0: 82, F1: ("" + ("foo bar" + "\n\t"))}
        v73 := v3[2147483648]
        for fi74 := 0; fi74 < 3; fi74++ {
            for fi75 := 0; fi75 < 1; fi75++ {
                v76 := GenS6{F0: (fi74 / (((fi74 | (fi74 << ((42) & 63)))) | 1)), F1: -1, F2: (fi75 << ((18446744073709551615) & 63))}
                var v77 bool = v73
                if (481 != 18446744073709551615) {
                    var v78 int32 = 0
                    _ = v78
                } else {
                    if (true && (v1 < (0.1 * v1))) {
                        var v79 uint = 65535
                        var v80 string = "hello"
                        _ = v79
                        _ = v80
                    } else {
                        if true {
                            v2 -= (v2 - (-(v2)))
                            var v81 float32 = 3.14159
                            var v82 uint32 = 822
                            _ = v81
                            _ = v82
                        }
                        var v83 int = 9223372036854775807
                        v84 := GenS7{F0: (v2 % ((-1) | 1)), F1: ("a" + "hello"), F2: "a"}
                        _ = v83
                        _ = v84
                    }
                    var v85 float32 = (((1e10 - 0.0) * v1) - v1)
                    switch (fi75 % (((fi75 % ((fi74) | 1))) | 1)) {
                    case -1:
                        delete(v3, 192)
                    default:
                        switch "a" {
                        case "\n\t":
                            v73 = ((0.1 - ((-2.25 * float64((1e10 / 1.0))) + ((((3.14159 * (((0.1 - 1e10) * (0.1 + float64(1.5))) + 3.14159)) / (0.0 + 0.1)) - (1.5 - -2.25)) / ((3.14159 + 1.5) - (-2.25 - (((0.1 + (0.0 * (1e10 + float64(v1)))) / (3.14159 + 1.5)) * (1e10 - float64(0.1)))))))) >= 1e10)
                        case "hello":
                            if false {
                                if (true && v77) {
                                    if (v73 || true) {
                                        delete(v3, 734)
                                    }
                                    v44 = append(v44, (^(fi75)))
                                    if len(v44) >= 2 {
                                        v44 = v44[1:len(v44)]
                                    }
                                }
                                if v73 {
                                    var v86 int32 = 42
                                    switch "" {
                                    case "foo bar":
                                        v1 = 0.0
                                    case "":
                                        var v87 float64 = (1e10 / (3.14159 + -2.25))
                                        _ = v87
                                    case "hello":
                                        v2 = (v2 - (v2 << ((65535) & 63)))
                                    }
                                    if v73 {
                                        v88 := GenS3{F0: (42 > 0)}
                                        _ = v88
                                    } else {
                                        if (1e10 < 0.0) {
                                            v1 = 0.0
                                            v89 := GenS1{F0: genFn0(), F1: "hello"}
                                            fmt.Println(v85, v73)
                                            _ = v89
                                        } else {
                                            var v90 uint = 65535
                                            switch fi74 {
                                            case 0:
                                                v2 += v2
                                            }
                                            v44 = append(v44, (fi75 % ((fi75) | 1)))
                                            var v91 int = fi75
                                            var v92 uint8 = 0
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                        }
                                        var v93 bool = (((v2 >> ((0) & 63)) != (v2 | v2)) || v77)
                                        _ = v93
                                    }
                                    _ = v86
                                }
                                v1 += v1
                            } else {
                                var v94 int64 = 0
                                _ = v94
                            }
                        }
                    }
                    _ = v85
                }
                _ = fi75
                _ = v76
                _ = v77
            }
            v95 := GenS6{F0: 1, F1: 0, F2: 42}
            var v96 uint16 = 0
            _ = fi74
            _ = v95
            _ = v96
        }
        _ = v72
        _ = v73
    } else {
        switch 4294967295 {
        case -697:
            _ = len(v3)
        case 128:
            for fi97 := 0; fi97 < 2; fi97++ {
                var v98 int = genFn0()
                v99 := GenS2{F0: -32, F1: v2}
                _ = fi97
                _ = v98
                _ = v99
            }
        }
    }
    v2 = (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v44
    _ = v71
}
