package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 bool; F2 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 int32; F1 uint64 }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 string }
func genFn0(p0 int8) int64 {
    v7 := GenS1{F0: 0, F1: 0, F2: 9223372036854775807}
    if true {
        if true {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 int64 = -9223372036854775808
                if true {
                    var v10 bool = false
                    switch fi8 {
                    case 32767:
                        var v11 bool = (v10 && v10)
                        _ = v11
                    case 42:
                        if (42 >= 1) {
                            for fi12 := 0; fi12 < 14; fi12++ {
                                v13 := GenS1{F0: -1, F1: 42, F2: (v9 % (((v9 | 1)) | 1))}
                                _ = fi12
                                _ = v13
                            }
                            v14 := []int{fi8, fi8}
                            for fi15 := 0; fi15 < 14; fi15++ {
                                if ((^(v9)) > -651) {
                                    var v16 uint16 = 728
                                    p0 = (-(p0))
                                    if (0.1 > float64(((((1e10 + ((3.14159 * 1e10) - (1.5 + (-2.25 + 3.14159)))) - 3.14159) * (float32((3.14159 - -2.25)) * 3.14159)) + 1e10))) {
                                        v14 = append(v14, (^(fi15)))
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                        v17 := GenS2{F0: 1}
                                        var v18 float32 = 3.14159
                                        _ = v17
                                        _ = v18
                                    } else {
                                        fmt.Println(fi15, p0)
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                        fmt.Println(p0)
                                        v9 = v9
                                    }
                                    _ = v16
                                } else {
                                    v9 = -1
                                }
                                _ = fi15
                            }
                            _ = v14
                        } else {
                            var v19 int32 = 0
                            _ = v19
                        }
                    case -1:
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v9 = v9
                            var v21 bool = false
                            _ = fi20
                            _ = v21
                        }
                    default:
                        v10 = (-2147483648 >= -1)
                    }
                    _ = v10
                } else {
                    v22 := GenS1{F0: 1, F1: 2, F2: int64(p0)}
                    v23 := []int{fi8}
                    _ = v22
                    _ = v23
                }
                _ = fi8
                _ = v9
            }
        }
        switch "a" {
        case "hello":
            var v24 uint16 = 32767
            _ = v24
        default:
            var v25 float32 = (1e10 / (((0.1 * 3.14159) * 1.5) - ((0.0 / 3.14159) / float32(1e10))))
            _ = v25
        }
        fmt.Println(p0)
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := []int{}
        _ = fi26
        _ = v27
    }
    fmt.Println(p0)
    p0 += p0
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 int32 = 1
    var v2 float64 = float64((0.1 - 0.1))
    var v3 uint16 = 32767
    v4 := make(map[uint32]bool)
    var v5 uint64 = 2
    v6 := []string{}
    fmt.Println(v1, v5, v1)
    fmt.Println(v3, v2)
    v2 = v2
    var v28 uint = 1
    for fi29 := 0; fi29 < 19; fi29++ {
        v6 = append(v6, "\n\t")
        _ = fi29
    }
    v30 := GenS1{F0: 0, F1: v28, F2: -752}
    v5 += (v5 & (^(v5)))
    v31 := GenS3{F0: 9223372036854775807, F1: true, F2: v1}
    switch ("" + (("a" + (("\n\t" + (((("\n\t" + ("" + "")) + ("hello" + "")) + ("" + ("\n\t" + "\n\t"))) + "")) + ("hello" + "hello"))) + "a")) {
    case "\n\t":
        for fi32 := 0; fi32 < 6; fi32++ {
            for fi33 := 0; fi33 < 2; fi33++ {
                switch (fi32 ^ 42) {
                case 127:
                    v34 := v4[1]
                    _ = v34
                case 1:
                    var v35 int8 = 53
                    _ = v35
                default:
                    fmt.Println(v2, v5)
                }
                _ = fi33
            }
            _ = fi32
        }
    }
    switch 127 {
    case 127:
        fmt.Println(v5)
    case 2147483647:
        v36 := GenS4{F0: 230}
        _ = v36
    case -640:
        _ = len(v4)
    }
    if (18446744073709551615 >= (v5 ^ (v5 * uint64(v28)))) {
        v37 := GenS5{F0: -179}
        v38 := GenS2{F0: 255}
        _ = v37
        _ = v38
    } else {
        if false {
            var v39 int = 179
            _ = v39
        }
        if true {
            v3 += 241
        } else {
            v4[2] = (!(-2.25 < (float64(0.0) / v2)))
            var v40 uint16 = (^(v3))
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 int32 = (v1 - 1)
                switch "a" {
                case "foo bar":
                    fmt.Println(fi41)
                case "\n\t":
                    v2 -= (0.1 * 1.5)
                case "a":
                    v43 := GenS2{F0: 255}
                    _ = v43
                }
                _ = fi41
                _ = v42
            }
            _ = v40
        }
        if (!false) {
            fmt.Println(v28, v28)
            if false {
                v3 += 42
            } else {
                v2 = v2
            }
            if false {
                v44 := v4[2147483647]
                _ = v44
            } else {
                delete(v4, 42)
            }
        }
    }
    for fi45 := 0; fi45 < 12; fi45++ {
        switch "" {
        case "":
            v46 := GenS2{F0: 42}
            _ = v46
        default:
            if (((v2 + -2.25) > v2) || (-2 != 1000)) {
                if false {
                    for fi47 := 0; fi47 < 11; fi47++ {
                        switch fi47 {
                        case 2147483647:
                            v48 := GenS5{F0: -32768}
                            _ = len(v6)
                            if (v1 < v1) {
                                fmt.Println(fi45, fi45)
                            } else {
                                switch "\n\t" {
                                case "foo bar":
                                    v49 := GenS5{F0: -149}
                                    _ = v49
                                case "hello":
                                    _ = len(v4)
                                    v2 -= (0.1 + v2)
                                }
                                if (18446744073709551615 > 18446744073709551615) {
                                    v3 -= (v3 + 72)
                                    if (v2 >= v2) {
                                        v50 := GenS6{F0: (!false)}
                                        _ = v50
                                    }
                                    v51 := GenS1{F0: 1000, F1: (v28 >> ((v28) & 63)), F2: 853}
                                    _ = v51
                                } else {
                                    var v52 uint16 = (v3 ^ 2)
                                    if (!(!(v28 > (v28 * (v28 % (((v28 - v28)) | 1)))))) {
                                        if false {
                                            v53 := GenS1{F0: 42, F1: 65535, F2: genFn0(42)}
                                            var v54 int32 = v1
                                            var v55 int64 = -1
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v52 = v3
                                            if true {
                                                v2 -= (((v2 - v2) * 1e10) * ((v2 - v2) + -2.25))
                                                v28 = (v28 * (v28 + (v28 * 255)))
                                                v1 += v1
                                            } else {
                                                v56 := GenS4{F0: 62}
                                                _ = v56
                                            }
                                        }
                                        fmt.Println(fi47)
                                        switch (^(fi45)) {
                                        case 917:
                                            var v57 string = ("\n\t" + "")
                                            _ = v57
                                        case 9223372036854775807:
                                            v58 := GenS5{F0: int16(v1)}
                                            _ = v58
                                        }
                                    } else {
                                        var v59 float32 = -2.25
                                        if (!true) {
                                            v3 = 1
                                            _ = len(v4)
                                        } else {
                                            if true {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                            v4[2147483647] = true
                                            v2 = v2
                                        }
                                        _ = v59
                                    }
                                    _ = v52
                                }
                                var v60 uint8 = 189
                                _ = v60
                            }
                            fmt.Println(v1)
                            _ = len(v4)
                            switch fi47 {
                            case -206:
                                var v61 float32 = (-2.25 + 3.14159)
                                _ = v61
                            case -663:
                                _ = len(v6)
                            case -680:
                                v5 = (v5 >> ((v28) & 63))
                            default:
                                var v62 float64 = ((v2 - 0.1) - v2)
                                if true {
                                    v2 -= (-2.25 * v2)
                                } else {
                                    fmt.Println(v3, v5, v1)
                                }
                                _ = v62
                            }
                            _ = v48
                        default:
                            if true {
                                var v63 bool = false
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v64 float32 = -2.25
                                _ = v63
                                _ = v64
                            } else {
                                _ = len(v4)
                                v65 := GenS6{F0: false}
                                _ = v65
                            }
                        }
                        _ = fi47
                    }
                    switch (("foo bar" + (((("" + "") + "hello") + "hello") + "foo bar")) + "foo bar") {
                    case "foo bar":
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            if false {
                                var v66 string = ""
                                v66 = "a"
                                _ = v66
                            } else {
                                _ = len(v4)
                            }
                        case "":
                            v67 := GenS7{F0: 0, F1: (v28 ^ (v28 ^ 2))}
                            _ = v67
                        case "a":
                            v1 += (v1 - (v1 % (((v1 & int32(fi45))) | 1)))
                        default:
                            if false {
                                fmt.Println(v5)
                            } else {
                                for fi68 := 0; fi68 < 4; fi68++ {
                                    v69 := GenS7{F0: (v5 ^ (^(v5))), F1: v28}
                                    _ = fi68
                                    _ = v69
                                }
                                if (124 != 117) {
                                    switch "foo bar" {
                                    case "":
                                        var v70 bool = true
                                        _ = v70
                                    case "foo bar":
                                        var v71 int32 = (v1 - (-(v1)))
                                        _ = v71
                                    default:
                                        v3 = (^(v3))
                                    }
                                    if false {
                                        var v72 string = ("hello" + "hello")
                                        var v73 int16 = -600
                                        if (false && (!(false && false))) {
                                            var v74 uint = (v28 / (((v28 / ((v28) | 1))) | 1))
                                            _ = v74
                                        }
                                        _ = v72
                                        _ = v73
                                    } else {
                                        v2 += float64(((1e10 / 1.5) - (3.14159 - 0.0)))
                                    }
                                    if true {
                                        fmt.Println(fi45, v5, v2)
                                        var v75 uint32 = 926
                                        switch ((("hello" + "") + "foo bar") + "foo bar") {
                                        case "foo bar":
                                            v75 -= (v75 >> ((2) & 63))
                                        default:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        _ = v75
                                    } else {
                                        fmt.Println(v3, v3, v1)
                                        if (1 <= -1) {
                                            if len(v6) > 0 {
                                                v6[0] = "hello"
                                            }
                                            for fi76 := 0; fi76 < 15; fi76++ {
                                                var v77 uint = 2
                                                v2 -= 1.5
                                                var v78 float64 = 3.14159
                                                _ = fi76
                                                _ = v77
                                                _ = v78
                                            }
                                            for fi79 := 0; fi79 < 5; fi79++ {
                                                if (1e10 != 1.5) {
                                                    fmt.Println(fi45, fi79)
                                                    var v80 uint16 = (v3 / ((v3) | 1))
                                                    if ((v2 - v2) >= v2) {
                                                        v1 = 42
                                                        switch "foo bar" {
                                                        case "hello":
                                                            fmt.Println(v3, v28)
                                                        }
                                                        v3 = (^(v80))
                                                    } else {
                                                        v3 = v3
                                                        v81 := v4[4294967295]
                                                        _ = len(v4)
                                                        var v82 int32 = v1
                                                        v83 := GenS2{F0: 2}
                                                        v28 = (^(v28))
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                    _ = v80
                                                } else {
                                                    var v84 int64 = 233
                                                    delete(v4, 2147483648)
                                                    v85 := GenS7{F0: 4294967295, F1: (^(v28))}
                                                    _ = v84
                                                    _ = v85
                                                }
                                                if true {
                                                    if (int8(fi79) == 42) {
                                                        fmt.Println(v5, fi45, fi45)
                                                    } else {
                                                        switch fi45 {
                                                        case -1:
                                                            if (!(42 == -1)) {
                                                                var v86 uint32 = 2147483647
                                                                fmt.Println(v2)
                                                                _ = v86
                                                            }
                                                        case 128:
                                                            switch fi79 {
                                                            case 0:
                                                                switch (("a" + (("" + ("a" + ("foo bar" + "hello"))) + "foo bar")) + ("\n\t" + "foo bar")) {
                                                                case "":
                                                                    v87 := GenS8{F0: fi79, F1: int32(v5)}
                                                                    _ = v87
                                                                }
                                                            }
                                                        case -2147483648:
                                                            var v88 uint = v28
                                                            v89 := GenS9{F0: int8(fi45), F1: 0}
                                                            if true {
                                                                _ = len(v6)
                                                            } else {
                                                                v2 = 0.0
                                                            }
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                    }
                                                    v90 := GenS3{F0: -1, F1: false, F2: (v1 | (v1 << ((uint(v5)) & 63)))}
                                                    var v91 uint = (^(v28))
                                                    var v92 uint32 = 2
                                                    _ = v90
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    v4[65535] = true
                                                }
                                                _ = fi79
                                            }
                                        } else {
                                            v28 -= 18446744073709551615
                                            v93 := GenS10{F0: (v1 | 42), F1: (^(v5))}
                                            var v94 uint64 = v5
                                            _ = v93
                                            _ = v94
                                        }
                                        v1 = v1
                                    }
                                }
                            }
                        }
                    case "a":
                        for fi95 := 0; fi95 < 17; fi95++ {
                            if false {
                                if true {
                                    if ((fi95 << (((^(v28))) & 63)) != (fi45 * (^(fi95)))) {
                                        var v96 float64 = ((((1.5 - v2) * 1.5) / v2) / float64((float32(0.0) - 3.14159)))
                                        _ = len(v6)
                                        _ = v96
                                    } else {
                                        fmt.Println(v2, v5, v3)
                                        if (42 < 2) {
                                            v97 := GenS10{F0: int32(v5), F1: (v5 % (((v5 | 4294967295)) | 1))}
                                            _ = v97
                                        } else {
                                            v5 -= (v5 | 1)
                                            var v98 uint8 = 1
                                            _ = v98
                                        }
                                        fmt.Println(v1, v5)
                                    }
                                    if true {
                                        v1 = (v1 % (((v1 << (((v28 & (v28 - (v28 << ((v28) & 63))))) & 63))) | 1))
                                        var v99 uint32 = 422
                                        delete(v4, v99)
                                        _ = v99
                                    } else {
                                        if false {
                                            var v100 uint32 = 0
                                            v101 := GenS5{F0: -838}
                                            v100 -= 1
                                            _ = v100
                                            _ = v101
                                        }
                                        fmt.Println(fi45, fi45, v5)
                                    }
                                    if true {
                                        var v102 int = (^(fi45))
                                        if len(v6) > 0 {
                                            v6[0] = ("" + "a")
                                        }
                                        _ = v102
                                    }
                                } else {
                                    switch ("a" + "hello") {
                                    case "foo bar":
                                        v4[2147483647] = (((1e10 + (0.1 * 1e10)) > 0.1) || false)
                                    case "a":
                                        v6 = append(v6, ("a" + "hello"))
                                    case "\n\t":
                                        v103 := GenS11{F0: fi45, F1: (("\n\t" + "a") + "\n\t")}
                                        _ = v103
                                    default:
                                        if (true && false) {
                                            if len(v6) > 0 {
                                                v6[0] = "foo bar"
                                            }
                                            v6 = append(v6, "\n\t")
                                        } else {
                                            v1 = (v1 - v1)
                                        }
                                    }
                                    v3 = v3
                                }
                            }
                            var v104 uint64 = 1
                            v105 := GenS3{F0: 42, F1: (-114 < 63), F2: (-(v1))}
                            v106 := GenS12{F0: (("a" + ("hello" + "")) + ("a" + (("a" + ("" + "hello")) + "\n\t")))}
                            _ = fi95
                            _ = v104
                            _ = v105
                            _ = v106
                        }
                    default:
                        if (1000 <= 42) {
                            v107 := v4[65535]
                            _ = v107
                        } else {
                            v108 := GenS11{F0: (fi45 << ((v28) & 63)), F1: ""}
                            var v109 uint8 = uint8(v28)
                            _ = v108
                            _ = v109
                        }
                    }
                }
            }
        }
        var v110 int64 = 9223372036854775807
        _ = fi45
        _ = v110
    }
    v3 = v3
    var v111 int64 = genFn0(-1)
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v111
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
