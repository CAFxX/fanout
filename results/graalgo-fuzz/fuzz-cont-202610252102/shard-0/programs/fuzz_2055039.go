package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int8; F1 int16; F2 string }
type GenS9 struct { F0 uint64; F1 int16; F2 uint }
type GenS10 struct { F0 string; F1 int64; F2 uint64 }
type GenS11 struct { F0 uint; F1 int; F2 int64 }
type GenS12 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int32, p1 int64, p2 int32) bool {
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 float32 = (-2.25 - 3.14159)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    switch 128 {
    case 1:
        v10 := GenS1{F0: -128, F1: 255, F2: 4294967295}
        _ = v10
    default:
        v11 := GenS1{F0: 1, F1: 1, F2: 1}
        _ = v11
    }
    v12 := GenS2{F0: 255, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return (!false)
}
func genFn1(p0 int64, p1 string, p2 bool) int16 {
    v13 := GenS2{F0: 0, F1: 255}
    var v14 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = (v1 * v1)
    var v3 uint64 = v2
    var v4 uint16 = 1
    var v5 int = -2147483648
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint16 = v4
        var v17 uint16 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 255, F1: 0}
    if (true && false) {
        switch "foo bar" {
        case "\n\t":
            if false {
                for fi19 := 0; fi19 < 3; fi19++ {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        if true {
                            v21 := GenS3{F0: ((("" + ("hello" + ("a" + ("" + "a")))) + ("hello" + ("foo bar" + "\n\t"))) + "a")}
                            fmt.Println(v1, fi19, v2)
                            v5 += int(v4)
                            _ = v21
                        } else {
                            v22 := []int{fi19, fi20, (fi20 + 0), 9223372036854775807}
                            v23 := GenS1{F0: -128, F1: 42, F2: 423}
                            _ = v22
                            _ = v23
                        }
                        v24 := GenS4{F0: (fi20 | 32767)}
                        v25 := []int{32767}
                        _ = fi20
                        _ = v24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 13; fi26++ {
                        var v27 int64 = -108
                        v28 := GenS1{F0: 82, F1: v4, F2: 104}
                        switch "" {
                        case "":
                            if ("\n\t" >= "a") {
                                v29 := []int{}
                                var v30 int = (v5 ^ v5)
                                v4 -= (v4 / ((v4) | 1))
                                v1 = v3
                                if true {
                                    if (true && false) {
                                        var v31 float32 = 1.5
                                        v30 += (^(fi26))
                                        _ = v31
                                    } else {
                                        if len(v29) >= 2 {
                                            v29 = v29[1:len(v29)]
                                        }
                                        v32 := GenS1{F0: 100, F1: v4, F2: 758}
                                        fmt.Println(v5)
                                        _ = v32
                                    }
                                } else {
                                    if len(v29) >= 2 {
                                        v29 = v29[1:len(v29)]
                                    }
                                    switch (^(v5)) {
                                    case -493:
                                        if true {
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                        } else {
                                            switch -1 {
                                            case 898:
                                                v2 = v2
                                            case 511:
                                                v33 := GenS5{F0: fi19}
                                                _ = v33
                                            default:
                                                if len(v29) > 0 {
                                                    v29[0] = (-(v5))
                                                }
                                            }
                                            switch (^(v5)) {
                                            case 1:
                                                var v34 int32 = 257
                                                _ = v34
                                            }
                                            if (v27 < v27) {
                                                v35 := GenS6{F0: "foo bar", F1: 1}
                                                _ = len(v29)
                                                if (!false) {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        _ = len(v29)
                                                    case "a":
                                                        if false {
                                                            v5 -= 32767
                                                        }
                                                    default:
                                                        v27 = (-(v27))
                                                    }
                                                    switch ("\n\t" + (("" + "") + "hello")) {
                                                    case "foo bar":
                                                        v36 := GenS2{F0: 42, F1: 1}
                                                        fmt.Println(v27)
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v37 int = (fi19 >> ((65535) & 63))
                                                            _ = v37
                                                        default:
                                                            if false {
                                                                v38 := GenS7{F0: -128, F1: v1}
                                                                v3 = v1
                                                                v27 -= 0
                                                                _ = v38
                                                            } else {
                                                                switch (("hello" + "\n\t") + "") {
                                                                case "hello":
                                                                    var v39 int16 = 529
                                                                    switch -1 {
                                                                    case 1:
                                                                        var v40 uint64 = v3
                                                                        var v41 uint = 892
                                                                        fmt.Println(v40)
                                                                        _ = v40
                                                                        _ = v41
                                                                    case -336:
                                                                        var v42 float32 = float32((1e10 / 1.5))
                                                                        _ = v42
                                                                    default:
                                                                        if ((2 > 0) && true) {
                                                                            _ = len(v29)
                                                                            if len(v29) > 0 {
                                                                                v29[0] = (v30 - fi26)
                                                                            }
                                                                            if (((128 <= 2) || false) && (((1.5 + 1e10) >= 3.14159) && (!((v27 >= (-(v27))) || false)))) {
                                                                                v43 := GenS5{F0: v30}
                                                                                v44 := GenS6{F0: "hello", F1: int16(v4)}
                                                                                v45 := GenS5{F0: int(v39)}
                                                                                _ = v43
                                                                                _ = v44
                                                                                _ = v45
                                                                            }
                                                                            v46 := GenS4{F0: v5}
                                                                            var v47 float64 = -2.25
                                                                            _ = v46
                                                                            _ = v47
                                                                        }
                                                                    }
                                                                    if ((v39 & (v39 - v39)) >= genFn1(9223372036854775807, (("" + "\n\t") + "\n\t"), true)) {
                                                                        if false {
                                                                            v3 += v1
                                                                            v2 += uint64(v4)
                                                                        } else {
                                                                            v48 := GenS8{F0: -128, F1: v39, F2: "\n\t"}
                                                                            _ = v48
                                                                        }
                                                                        if false {
                                                                            var v49 int64 = (v27 - 708)
                                                                            _ = v49
                                                                        } else {
                                                                            v50 := GenS7{F0: int8(v4), F1: v2}
                                                                            v27 -= (v27 ^ 9223372036854775807)
                                                                            _ = v50
                                                                        }
                                                                        var v51 uint = 1
                                                                        _ = v51
                                                                    }
                                                                    _ = v39
                                                                case "\n\t":
                                                                    if false {
                                                                        var v52 int64 = (^(v27))
                                                                        _ = v52
                                                                    }
                                                                case "":
                                                                    var v53 uint64 = 1
                                                                    _ = v53
                                                                default:
                                                                    v54 := GenS1{F0: 42, F1: v4, F2: 42}
                                                                    _ = v54
                                                                }
                                                                fmt.Println(v2, v3, v5)
                                                                if (526 >= (v2 % ((v3) | 1))) {
                                                                    var v55 uint64 = v1
                                                                    _ = v55
                                                                }
                                                            }
                                                        }
                                                        _ = v36
                                                    case "\n\t":
                                                        var v56 uint = 479
                                                        _ = v56
                                                    case "":
                                                        if len(v29) > 0 {
                                                            v29[0] = v5
                                                        }
                                                    }
                                                    var v57 float32 = 3.14159
                                                    _ = v57
                                                }
                                                _ = v35
                                            }
                                        }
                                    case 42:
                                        if true {
                                            switch v30 {
                                            case -2147483648:
                                                v27 = v27
                                            case 9223372036854775807:
                                                var v58 uint16 = (v4 & v4)
                                                _ = v58
                                            }
                                            v1 = 1
                                            var v59 uint8 = 193
                                            _ = v59
                                        } else {
                                            fmt.Println(fi26, v2)
                                            if (1 != 42) {
                                                var v60 uint16 = 1
                                                var v61 float32 = (0.0 * -2.25)
                                                _ = v60
                                                _ = v61
                                            }
                                            v62 := GenS7{F0: -1, F1: v1}
                                            _ = v62
                                        }
                                    default:
                                        var v63 int8 = 31
                                        _ = v63
                                    }
                                }
                                _ = v29
                                _ = v30
                            } else {
                                var v64 int16 = genFn1(-633, "\n\t", (true && (false && (fi26 < fi26))))
                                _ = v64
                            }
                        case "a":
                            if true {
                                switch "\n\t" {
                                case "":
                                    v65 := []int{32767, -9223372036854775808, (v5 % ((4294967295) | 1)), 42}
                                    _ = v65
                                case "a":
                                    var v66 int = (fi19 & (^(fi19)))
                                    _ = v66
                                case "foo bar":
                                    v67 := []int{}
                                    _ = v67
                                default:
                                    fmt.Println(v5, v4, v4)
                                }
                            } else {
                                switch 4294967295 {
                                case 1:
                                    if true {
                                        switch 4294967295 {
                                        case -2147483648:
                                            var v68 string = "hello"
                                            _ = v68
                                        case 32767:
                                            v69 := GenS8{F0: 0, F1: -191, F2: "foo bar"}
                                            _ = v69
                                        }
                                        var v70 int64 = v27
                                        _ = v70
                                    }
                                default:
                                    fmt.Println(fi19, v3, v4)
                                }
                                if (false && false) {
                                    fmt.Println(v5, fi19)
                                    var v71 uint = 255
                                    _ = v71
                                } else {
                                    var v72 bool = (-1 != 1)
                                    v2 = 189
                                    _ = v72
                                }
                            }
                        case "\n\t":
                            switch ("" + ((("\n\t" + "foo bar") + "a") + "\n\t")) {
                            case "":
                                v3 += (v2 << ((1) & 63))
                            case "a":
                                fmt.Println(v3)
                            case "foo bar":
                                var v73 float32 = 0.0
                                _ = v73
                            default:
                                if ((v4 % ((42) | 1)) < 1) {
                                    v2 = (v2 << ((18446744073709551615) & 63))
                                } else {
                                    var v74 int16 = 1000
                                    v2 += 1
                                    switch v5 {
                                    case 42:
                                        v74 += 1
                                    case 1:
                                        v75 := []int{(fi19 % (((fi26 + fi19)) | 1)), (^(v5))}
                                        _ = v75
                                    default:
                                        var v76 uint = 65535
                                        _ = v76
                                    }
                                    _ = v74
                                }
                            }
                        }
                        if false {
                            var v77 bool = true
                            _ = v77
                        }
                        var v78 uint = 1
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v78
                    }
                    _ = fi19
                }
            } else {
                if false {
                    if (957 != -1) {
                        if true {
                            v1 += v1
                            v79 := []int{v5}
                            v80 := GenS4{F0: (v5 / ((-1) | 1))}
                            _ = v79
                            _ = v80
                        }
                        v81 := GenS1{F0: 100, F1: 42, F2: 554}
                        _ = v81
                    } else {
                        for fi82 := 0; fi82 < 7; fi82++ {
                            fmt.Println(v2)
                            _ = fi82
                        }
                        v2 = v2
                    }
                    v3 = (v3 | (v2 - v1))
                }
                v83 := []int{v5, v5}
                _ = v83
            }
        case "":
            v4 += (v4 & v4)
        }
        var v84 float32 = (0.0 + 0.1)
        _ = v84
    } else {
        v1 = v3
        v85 := GenS9{F0: (v2 >> ((1) & 63)), F1: 1, F2: 866}
        _ = v85
    }
    v86 := GenS10{F0: "", F1: 42, F2: v1}
    v87 := GenS4{F0: v5}
    var v88 uint8 = 184
    v1 += (v1 >> ((390) & 63))
    for fi89 := 0; fi89 < 10; fi89++ {
        v90 := GenS4{F0: fi89}
        _ = fi89
        _ = v90
    }
    if ((true && true) || true) {
        v3 = (v2 << ((255) & 63))
        if false {
            v91 := GenS5{F0: -930}
            _ = v91
        } else {
            fmt.Println(v4)
            switch "\n\t" {
            case "foo bar":
                v92 := GenS8{F0: 100, F1: genFn1(9223372036854775807, "foo bar", (42 != 1)), F2: "\n\t"}
                _ = v92
            case "\n\t":
                v4 = v4
            case "a":
                switch ("a" + "hello") {
                case "\n\t":
                    for fi93 := 0; fi93 < 14; fi93++ {
                        v94 := GenS8{F0: 110, F1: 1, F2: "hello"}
                        _ = fi93
                        _ = v94
                    }
                case "hello":
                    if (genFn1(-9223372036854775808, "a", (!false)) <= genFn1(42, "\n\t", ("\n\t" > ((("hello" + "hello") + "foo bar") + "\n\t")))) {
                        var v95 uint = 255
                        _ = v95
                    }
                case "":
                    var v96 uint16 = v4
                    _ = v96
                }
            default:
                var v97 int = v5
                _ = v97
            }
        }
        for fi98 := 0; fi98 < 7; fi98++ {
            for fi99 := 0; fi99 < 8; fi99++ {
                var v100 float32 = 3.14159
                v5 -= 128
                _ = fi99
                _ = v100
            }
            var v101 string = ""
            v102 := GenS11{F0: 4294967295, F1: (fi98 + (fi98 << ((237) & 63))), F2: 9223372036854775807}
            _ = fi98
            _ = v101
            _ = v102
        }
    } else {
        if (true || (v4 > (v4 & (v4 / (((v4 >> ((105) & 63))) | 1))))) {
            if (65535 > 72) {
                v103 := GenS12{F0: v1, F1: (v4 << ((1) & 63))}
                v5 = v5
                _ = v103
            } else {
                fmt.Println(v88)
            }
        } else {
            var v104 int = (v5 ^ (v5 + (v5 * v5)))
            _ = v104
        }
        v1 = (v2 ^ (v2 | (v3 << ((1) & 63))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v87
    _ = v88
}
