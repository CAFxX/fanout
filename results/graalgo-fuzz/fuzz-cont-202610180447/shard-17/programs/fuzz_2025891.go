package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int32; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint16 }
func genFn0() uint64 {
    if (!false) {
        if (!false) {
            var v5 int = 32767
            if (!(563 > -1)) {
                v5 = v5
                fmt.Println(v5)
            }
            v5 = v5
            _ = v5
        }
        var v6 int8 = 42
        _ = v6
    } else {
        for fi7 := 0; fi7 < 5; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
    }
    var v8 uint16 = 1
    var v9 int = -782
    v10 := []int{v9, (v9 + v9), (v9 >> ((4294967295) & 63))}
    v9 = (v9 | -959)
    _ = v10
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 string, p1 int64) uint64 {
    fmt.Println(p0)
    v11 := GenS1{F0: p1, F1: 65535}
    p0 = "a"
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    var v1 uint8 = 42
    var v2 uint = 668
    v3 := make(map[uint32]bool)
    var v4 bool = false
    v4 = (!("foo bar" >= "hello"))
    v12 := GenS1{F0: 796, F1: 0}
    v2 -= 0
    v4 = v4
    v4 = (91 >= v1)
    switch (("\n\t" + "hello") + "") {
    case "hello":
        v2 = v2
    case "foo bar":
        v4 = (true || (2 < 1))
    }
    v2 = 1
    v13 := []int{}
    var v14 float64 = 1.5
    switch 128 {
    case -396:
        fmt.Println(v2, v2, v14)
    case -1:
        var v15 uint16 = 133
        _ = v15
    default:
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS1{F0: -1, F1: 423}
            v18 := GenS2{F0: 1}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v19 := GenS3{F0: (4294967295 != 32767), F1: -1, F2: (v1 ^ (^(v1)))}
    for fi20 := 0; fi20 < 16; fi20++ {
        v4 = (-1 >= -637)
        for fi21 := 0; fi21 < 8; fi21++ {
            switch 1 {
            case 110:
                fmt.Println(fi21, v2)
            default:
                v14 = (0.1 - float64(1.5))
                if (!true) {
                    var v22 int8 = -43
                    var v23 uint32 = 42
                    v14 = ((float64(float32(float64((0.1 + 1.5)))) / v14) - v14)
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS4{F0: 42, F1: 2}
                    _ = v24
                }
                if v4 {
                    if (-1 <= 9223372036854775807) {
                        v13 = append(v13, 9223372036854775807)
                        fmt.Println(v2)
                    }
                    var v25 int64 = -380
                    v14 = 1.5
                    _ = v25
                }
            }
            v4 = (-1 == 93)
            v1 = (v1 * (v1 & v1))
            _ = fi21
        }
        delete(v3, 60)
        _ = fi20
    }
    if true {
        fmt.Println(v14, v2)
        if false {
            v1 += v1
            v14 += (1.5 * (v14 * -2.25))
            v4 = (!(-2147483648 == -2147483648))
        }
    }
    switch -705 {
    case -596:
        v13 = append(v13, -1)
    case 128:
        switch (("\n\t" + ("foo bar" + "foo bar")) + "\n\t") {
        case "foo bar":
            v1 = v1
        }
    case -690:
        if v4 {
            if false {
                var v26 int8 = 100
                switch int(v26) {
                case -519:
                    fmt.Println(v14, v2, v14)
                default:
                    v3[1] = (false || v4)
                }
                _ = v26
            }
        } else {
            v27 := v3[2147483648]
            _ = v27
        }
    default:
        if true {
            switch -2147483648 {
            case 9223372036854775807:
                if v4 {
                    fmt.Println(v4)
                } else {
                    switch "foo bar" {
                    case "hello":
                        v1 -= (v1 + (v1 << ((v2) & 63)))
                    }
                    v28 := GenS3{F0: (!true), F1: 42, F2: v1}
                    v14 = -2.25
                    _ = v28
                }
            case 4294967295:
                switch ("" + "") {
                case "":
                    if (1000 < 42) {
                        _ = len(v13)
                    }
                default:
                    var v29 uint = 4294967295
                    _ = v29
                }
            default:
                v14 = (v14 - (0.1 * v14))
            }
            v30 := GenS5{F0: false, F1: -1}
            _ = v30
        } else {
            if ((2147483647 != -427) || ((9223372036854775807 != -1) || true)) {
                fmt.Println(v1, v1, v14)
                v31 := GenS3{F0: (0 == 2147483647), F1: int32(v2), F2: v1}
                _ = v31
            } else {
                var v32 uint32 = 2147483648
                if (0 < 1000) {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v1 -= 127
                        _ = fi33
                    }
                    delete(v3, v32)
                    var v34 int64 = -1
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 5; fi35++ {
                        if true {
                            if v4 {
                                switch "" {
                                case "hello":
                                    v1 += (v1 & v1)
                                case "\n\t":
                                    v1 -= 127
                                case "foo bar":
                                    for fi36 := 0; fi36 < 8; fi36++ {
                                        v2 = uint(v1)
                                        if true {
                                            var v37 uint8 = v1
                                            if len(v13) > 0 {
                                                v13[0] = -626
                                            }
                                            _ = v37
                                        } else {
                                            v32 = v32
                                            var v38 int16 = -375
                                            if len(v13) > 0 {
                                                v13[0] = (-(fi35))
                                            }
                                            _ = v38
                                        }
                                        _ = fi36
                                    }
                                default:
                                    v32 = v32
                                }
                                v1 += (v1 / ((77) | 1))
                                v4 = (!v4)
                            } else {
                                if (!(true && true)) {
                                    v39 := GenS5{F0: v4, F1: 0}
                                    _ = v39
                                }
                                v14 = 1e10
                                if v4 {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    v1 += uint8(v32)
                                }
                            }
                            var v40 uint8 = (^(v1))
                            var v41 bool = v4
                            _ = v40
                            _ = v41
                        }
                        v1 = v1
                        var v42 uint64 = 4294967295
                        _ = fi35
                        _ = v42
                    }
                    if v4 {
                        switch ("foo bar" + "hello") {
                        case "hello":
                            if true {
                                for fi43 := 0; fi43 < 1; fi43++ {
                                    v44 := GenS6{F0: 2}
                                    _ = fi43
                                    _ = v44
                                }
                                for fi45 := 0; fi45 < 15; fi45++ {
                                    var v46 int64 = -151
                                    var v47 uint16 = 972
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                }
                                var v48 float64 = (((v14 - 0.0) / v14) + v14)
                                _ = v48
                            } else {
                                for fi49 := 0; fi49 < 13; fi49++ {
                                    var v50 int8 = -1
                                    _ = fi49
                                    _ = v50
                                }
                                if (v4 || ((("a" != "\n\t") || false) && (v4 || false))) {
                                    var v51 uint64 = 1
                                    if v4 {
                                        for fi52 := 0; fi52 < 18; fi52++ {
                                            fmt.Println(fi52, v32)
                                            _ = fi52
                                        }
                                    } else {
                                        v53 := GenS5{F0: v4, F1: 909}
                                        v54 := GenS3{F0: v4, F1: -2147483648, F2: v1}
                                        _ = v53
                                        _ = v54
                                    }
                                    _ = v51
                                } else {
                                    fmt.Println(v32)
                                    v4 = false
                                    v13 = append(v13, 32767)
                                }
                                var v55 int8 = -118
                                _ = v55
                            }
                        }
                        v56 := GenS7{F0: (v32 ^ v32), F1: v32}
                        for fi57 := 0; fi57 < 18; fi57++ {
                            v58 := GenS7{F0: 1, F1: 873}
                            switch "\n\t" {
                            case "":
                                for fi59 := 0; fi59 < 4; fi59++ {
                                    if v4 {
                                        var v60 int32 = 490
                                        var v61 uint8 = 245
                                        if v4 {
                                            v62 := GenS3{F0: (!(0.0 <= 0.1)), F1: v60, F2: (^(v1))}
                                            if (281 > 18446744073709551615) {
                                                fmt.Println(v32, v32, v14)
                                                var v63 uint16 = 1
                                                _ = v63
                                            } else {
                                                _ = len(v3)
                                                v64 := GenS8{F0: 1}
                                                var v65 uint16 = 42
                                                _ = v64
                                                _ = v65
                                            }
                                            _ = v62
                                        } else {
                                            var v66 uint64 = 2
                                            v67 := GenS6{F0: 4294967295}
                                            var v68 int8 = 42
                                            v1 = uint8(v32)
                                            _ = v66
                                            _ = v67
                                            _ = v68
                                        }
                                        if len(v13) > 0 {
                                            v13[0] = -2147483648
                                        }
                                        _ = v60
                                        _ = v61
                                    } else {
                                        _ = len(v3)
                                    }
                                    _ = fi59
                                }
                            }
                            fmt.Println(v1)
                            _ = fi57
                            _ = v58
                        }
                        _ = v56
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            switch ("foo bar" + (((("\n\t" + "hello") + "a") + "") + ("hello" + ""))) {
                            case "foo bar":
                                var v69 int64 = 1
                                _ = v69
                            case "\n\t":
                                v1 = v1
                            case "":
                                if v4 {
                                    var v70 uint16 = 38
                                    fmt.Println(v14, v2)
                                    switch -208 {
                                    case 9223372036854775807:
                                        for fi71 := 0; fi71 < 2; fi71++ {
                                            switch (("\n\t" + "hello") + "\n\t") {
                                            case "foo bar":
                                                fmt.Println(v1, v14)
                                            }
                                            _ = fi71
                                        }
                                    default:
                                        if true {
                                            var v72 string = "a"
                                            switch 42 {
                                            case 0:
                                                var v73 float32 = 0.1
                                                _ = v73
                                            case 42:
                                                v2 = v2
                                            }
                                            v72 = ("" + v72)
                                            _ = v72
                                        } else {
                                            fmt.Println(v32)
                                        }
                                    }
                                    _ = v70
                                } else {
                                    var v74 int16 = 1000
                                    _ = len(v13)
                                    _ = v74
                                }
                            default:
                                v75 := GenS9{F0: -966}
                                _ = v75
                            }
                        case "a":
                            for fi76 := 0; fi76 < 2; fi76++ {
                                switch (fi76 & -9223372036854775808) {
                                case 270:
                                    for fi77 := 0; fi77 < 7; fi77++ {
                                        v3[(^(v32))] = v4
                                        var v78 int8 = -128
                                        if v4 {
                                            if v4 {
                                                fmt.Println(v14, v14, v2)
                                                v79 := GenS7{F0: v32, F1: (^(v32))}
                                                _ = v79
                                            }
                                        }
                                        fmt.Println(fi76, v32, v2)
                                        _ = fi77
                                        _ = v78
                                    }
                                case -9223372036854775808:
                                    var v80 int64 = -9223372036854775808
                                    _ = v80
                                case 465:
                                    v2 -= v2
                                default:
                                    if v4 {
                                        var v81 float64 = v14
                                        _ = v81
                                    }
                                }
                                _ = fi76
                            }
                        case "":
                            if true {
                                if ((^(v2)) == 42) {
                                    v82 := GenS8{F0: 100}
                                    _ = v82
                                } else {
                                    v2 += 2
                                    var v83 uint32 = v32
                                    var v84 bool = (90 != -1)
                                    _ = v83
                                    _ = v84
                                }
                                if true {
                                    var v85 uint = 181
                                    for fi86 := 0; fi86 < 10; fi86++ {
                                        v87 := GenS10{F0: (fi86 + fi86), F1: 1}
                                        _ = fi86
                                        _ = v87
                                    }
                                    if v4 {
                                        switch "foo bar" {
                                        case "":
                                            var v88 int64 = 140
                                            _ = v88
                                        default:
                                            if v4 {
                                                var v89 bool = false
                                                if (750 >= 932) {
                                                    fmt.Println(v14, v32, v85)
                                                    fmt.Println(v32, v1)
                                                } else {
                                                    v13 = append(v13, 127)
                                                    var v90 int8 = -15
                                                    _ = v90
                                                }
                                                var v91 int32 = 33
                                                _ = v89
                                                _ = v91
                                            }
                                        }
                                        for fi92 := 0; fi92 < 14; fi92++ {
                                            switch ((("\n\t" + ("a" + "hello")) + "foo bar") + "a") {
                                            case "\n\t":
                                                v14 -= (v14 + 1e10)
                                            }
                                            v1 = v1
                                            var v93 uint16 = 65535
                                            _ = fi92
                                            _ = v93
                                        }
                                        v94 := GenS9{F0: 1}
                                        _ = v94
                                    } else {
                                        var v95 float64 = (3.14159 - (v14 + (0.1 * (v14 + v14))))
                                        _ = v95
                                    }
                                    _ = v85
                                }
                                var v96 int8 = -62
                                _ = v96
                            } else {
                                v4 = (v4 && v4)
                                fmt.Println(v14, v32)
                            }
                        }
                    }
                    var v97 int32 = 42
                    _ = v97
                }
                if v4 {
                    _ = len(v13)
                    if len(v13) > 0 {
                        v13[0] = -2147483648
                    }
                } else {
                    var v98 string = "\n\t"
                    for fi99 := 0; fi99 < 13; fi99++ {
                        var v100 uint64 = 1
                        var v101 uint32 = (v32 >> (((v2 % ((v2) | 1))) & 63))
                        _ = fi99
                        _ = v100
                        _ = v101
                    }
                    if v4 {
                        if v4 {
                            var v102 int32 = -401
                            switch "a" {
                            case "hello":
                                delete(v3, uint32(v1))
                            default:
                                v98 = ("\n\t" + v98)
                            }
                            _ = v102
                        } else {
                            var v103 bool = (!(2 <= 65535))
                            if v103 {
                                var v104 int16 = -32768
                                v32 += 0
                                _ = v104
                            }
                            _ = v103
                        }
                    }
                    _ = v98
                }
                _ = v32
            }
            var v105 int32 = -52
            var v106 uint64 = 1
            _ = v105
            _ = v106
        }
    }
    for fi107 := 0; fi107 < 20; fi107++ {
        delete(v3, 2)
        _ = fi107
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
