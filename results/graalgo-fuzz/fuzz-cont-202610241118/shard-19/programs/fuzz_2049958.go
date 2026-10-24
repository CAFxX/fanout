package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint; F1 int; F2 uint64 }
type GenS4 struct { F0 int16; F1 int; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8 }
func genFn0(p0 bool, p1 int16) bool {
    switch ("" + ("\n\t" + "foo bar")) {
    case "foo bar":
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := []int{fi6, (fi6 | 127)}
            for fi8 := 0; fi8 < 10; fi8++ {
                fmt.Println(p1, p0)
                fmt.Println(fi6, p1, p0)
                _ = fi8
            }
            var v9 int8 = 71
            _ = fi6
            _ = v7
            _ = v9
        }
    case "":
        v10 := GenS1{F0: 0, F1: 50}
        _ = v10
    case "a":
        fmt.Println(p1)
    default:
        if (!(true || (!(42 >= 2)))) {
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := []int{(fi11 | fi11), fi11, fi11}
                v13 := GenS1{F0: 255, F1: 1}
                fmt.Println(p1, p0)
                _ = fi11
                _ = v12
                _ = v13
            }
        } else {
            fmt.Println(p0, p1)
        }
    }
    var v14 float32 = (3.14159 - 1.5)
    switch (("" + "") + "\n\t") {
    case "\n\t":
        v15 := GenS1{F0: 255, F1: 255}
        _ = v15
    case "foo bar":
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := GenS1{F0: 42, F1: 190}
            _ = fi16
            _ = v17
        }
    case "hello":
        switch ((("a" + ("" + "a")) + (("hello" + ((("" + "a") + ("" + "")) + ("hello" + "a"))) + "\n\t")) + ("foo bar" + "\n\t")) {
        case "\n\t":
            v18 := []int{}
            _ = v18
        case "":
            if ((true && false) && (2 != 205)) {
                for fi19 := 0; fi19 < 17; fi19++ {
                    if (true || ((-288 > 0) || p0)) {
                        if true {
                            for fi20 := 0; fi20 < 13; fi20++ {
                                v21 := GenS1{F0: 2, F1: 42}
                                var v22 bool = false
                                p1 = p1
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            v14 = 0.0
                            if false {
                                for fi23 := 0; fi23 < 6; fi23++ {
                                    v24 := []int{-220, fi23, (fi19 % (((fi23 >> ((2) & 63))) | 1)), (fi19 & fi23)}
                                    _ = fi23
                                    _ = v24
                                }
                                v25 := []int{127, (fi19 * -1), (fi19 << ((358) & 63))}
                                for fi26 := 0; fi26 < 7; fi26++ {
                                    v25 = append(v25, fi19)
                                    var v27 float64 = (((0.0 / 1.0) * ((0.1 - (0.0 * 0.1)) * -2.25)) + (1.5 + (1.5 + (0.1 / 3.14159))))
                                    _ = fi26
                                    _ = v27
                                }
                                _ = v25
                            }
                        }
                        var v28 uint32 = 324
                        for fi29 := 0; fi29 < 11; fi29++ {
                            if p0 {
                                var v30 uint64 = 881
                                if p0 {
                                    v31 := []int{0}
                                    v14 = v14
                                    switch fi29 {
                                    case 9223372036854775807:
                                        if p0 {
                                            p1 = 1
                                            var v32 uint = 4294967295
                                            _ = v32
                                        } else {
                                            var v33 int = fi29
                                            v34 := GenS2{F0: 18446744073709551615, F1: ("hello" + ("foo bar" + ("a" + "")))}
                                            v28 = 2147483647
                                            _ = v33
                                            _ = v34
                                        }
                                    case -699:
                                        var v35 uint32 = v28
                                        _ = v35
                                    case 32767:
                                        switch "a" {
                                        case "a":
                                            var v36 uint16 = 32767
                                            _ = v36
                                        case "hello":
                                            var v37 uint8 = 148
                                            _ = v37
                                        default:
                                            if len(v31) >= 2 {
                                                v31 = v31[1:len(v31)]
                                            }
                                        }
                                    default:
                                        _ = len(v31)
                                    }
                                    v31 = append(v31, (fi29 + (fi19 % ((2147483647) | 1))))
                                    v38 := GenS2{F0: v30, F1: "\n\t"}
                                    _ = v31
                                    _ = v38
                                }
                                _ = v30
                            }
                            v28 = 2
                            v39 := []int{fi19, (fi29 << ((499) & 63)), 4294967295}
                            _ = fi29
                            _ = v39
                        }
                        _ = v28
                    } else {
                        v40 := GenS3{F0: 42, F1: 2147483647, F2: 1}
                        fmt.Println(v14, p0, v14)
                        _ = v40
                    }
                    _ = fi19
                }
            } else {
                v41 := []int{-9223372036854775808, 0, 42, -645}
                _ = v41
            }
        default:
            p1 += 0
        }
    }
    var v42 uint8 = 0
    _ = p0
    _ = p1
    _ = v14
    _ = v42
    return (636 >= 313)
}
func genFn1(p0 int32, p1 int16, p2 uint64) int {
    if false {
        fmt.Println(p1)
    }
    var v43 uint64 = p2
    v43 = (^(p2))
    var v44 int64 = 0
    var v45 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    _ = v44
    _ = v45
    return -641
}

func main() {
    v1 := []string{((("\n\t" + "hello") + "a") + ("hello" + "foo bar"))}
    var v2 int64 = 9223372036854775807
    var v3 int32 = 2147483647
    var v4 float32 = 0.1
    var v5 int32 = (v3 - v3)
    var v46 bool = false
    for fi47 := 0; fi47 < 2; fi47++ {
        if v46 {
            for fi48 := 0; fi48 < 14; fi48++ {
                v4 -= v4
                _ = fi48
            }
            var v49 string = "hello"
            _ = v49
        }
        _ = fi47
    }
    var v50 uint32 = 0
    for fi51 := 0; fi51 < 6; fi51++ {
        if v46 {
            for fi52 := 0; fi52 < 2; fi52++ {
                var v53 bool = (!false)
                switch ("a" + "a") {
                case "\n\t":
                    if v53 {
                        if v53 {
                            var v54 bool = v46
                            v3 = (v5 >> ((0) & 63))
                            if v54 {
                                fmt.Println(v5, v46, v5)
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            }
                            _ = v54
                        }
                        var v55 int8 = 4
                        fmt.Println(v5, v4)
                        var v56 uint16 = 255
                        if (v46 && ("a" != ("a" + "foo bar"))) {
                            v56 = 2
                            if true {
                                _ = len(v1)
                            } else {
                                if (int8(v50) > v55) {
                                    v1 = append(v1, "")
                                    var v57 uint = 4294967295
                                    v53 = true
                                    _ = v57
                                }
                                switch 127 {
                                case 128:
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                    v58 := GenS4{F0: 0, F1: (-(fi51)), F2: 128}
                                    if ((1 != -32768) || v46) {
                                        var v59 float64 = float64(v4)
                                        if len(v1) > 0 {
                                            v1[0] = ""
                                        }
                                        _ = v59
                                    } else {
                                        v60 := GenS2{F0: 1, F1: "a"}
                                        _ = len(v1)
                                        _ = v60
                                    }
                                    _ = v58
                                case 1:
                                    v3 = v5
                                case -2147483648:
                                    var v61 string = ("a" + "foo bar")
                                    _ = v61
                                default:
                                    var v62 uint = 579
                                    _ = v62
                                }
                            }
                        }
                        var v63 int16 = 42
                        _ = v55
                        _ = v56
                        _ = v63
                    }
                case "a":
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                case "foo bar":
                    var v64 bool = (82 != 4294967295)
                    v5 = (v3 & (-(v5)))
                    _ = v64
                }
                _ = fi52
                _ = v53
            }
            var v65 float64 = 1.5
            _ = v65
        }
        _ = fi51
    }
    var v66 int32 = v3
    fmt.Println(v2)
    for fi67 := 0; fi67 < 20; fi67++ {
        if false {
            var v68 uint = 255
            v50 -= v50
            for fi69 := 0; fi69 < 8; fi69++ {
                var v70 int64 = (v2 ^ (v2 * v2))
                _ = fi69
                _ = v70
            }
            _ = v68
        } else {
            var v71 int32 = (v3 * 2147483647)
            v66 = (v3 << ((670) & 63))
            switch (fi67 * (fi67 / ((fi67) | 1))) {
            case -9223372036854775808:
                if false {
                    if (32767 != 891) {
                        for fi72 := 0; fi72 < 1; fi72++ {
                            var v73 string = ("foo bar" + ("" + ""))
                            v74 := GenS5{F0: 128, F1: 65535, F2: 65}
                            var v75 int32 = (v66 | (^(v5)))
                            _ = fi72
                            _ = v73
                            _ = v74
                            _ = v75
                        }
                        v50 = v50
                        v76 := GenS5{F0: 175, F1: 65535, F2: 255}
                        _ = v76
                    } else {
                        v77 := GenS6{F0: false}
                        var v78 bool = (!(v4 == ((((v4 / (v4 - float32(-2.25))) + 0.1) * v4) / v4)))
                        _ = v77
                        _ = v78
                    }
                } else {
                    switch (fi67 / (((fi67 * (^(fi67)))) | 1)) {
                    case -867:
                        v46 = (!(42 >= 1))
                    case 0:
                        if v46 {
                            fmt.Println(v71, v2, v5)
                            for fi79 := 0; fi79 < 11; fi79++ {
                                var v80 int8 = -1
                                var v81 bool = true
                                _ = fi79
                                _ = v80
                                _ = v81
                            }
                            var v82 string = "\n\t"
                            _ = v82
                        } else {
                            v4 = ((v4 - 1e10) + 0.1)
                        }
                    case -795:
                        v46 = v46
                    }
                }
            }
            _ = v71
        }
        _ = fi67
    }
    var v83 uint32 = v50
    switch -147 {
    case -736:
        if (v46 || (0.1 == 1e10)) {
            v50 = 310
            var v84 uint = 255
            v50 = v50
            _ = v84
        } else {
            if v46 {
                var v85 int = 127
                for fi86 := 0; fi86 < 10; fi86++ {
                    v85 -= (v85 % ((378) | 1))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi86
                }
                var v87 uint8 = 1
                _ = v85
                _ = v87
            } else {
                v3 = v3
            }
        }
    default:
        switch ("" + ("" + ("" + ("\n\t" + "foo bar")))) {
        case "a":
            if (653 != 0) {
                if (-590 >= 1) {
                    var v88 int = 2147483647
                    v89 := GenS4{F0: 402, F1: int(v50), F2: 255}
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    _ = v88
                    _ = v89
                } else {
                    if false {
                        fmt.Println(v5)
                    }
                    v90 := GenS2{F0: 42, F1: "a"}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v90
                }
                var v91 uint = 2
                _ = v91
            }
        }
    }
    v92 := GenS7{F0: ("foo bar" + "foo bar")}
    for fi93 := 0; fi93 < 10; fi93++ {
        var v94 float32 = 0.0
        v3 -= v3
        v95 := GenS8{F0: "hello"}
        _ = fi93
        _ = v94
        _ = v95
    }
    fmt.Println(v66)
    switch "" {
    case "hello":
        var v96 uint64 = 1
        _ = v96
    }
    for fi97 := 0; fi97 < 8; fi97++ {
        if v46 {
            var v98 int = genFn1((v66 & (-(v66))), 42, 1)
            _ = v98
        } else {
            var v99 uint32 = (^(v83))
            v100 := GenS9{F0: -128}
            v1 = append(v1, "\n\t")
            _ = v100
            _ = v99
        }
        var v101 int8 = -1
        _ = fi97
        _ = v101
    }
    switch ("" + ((("\n\t" + (("\n\t" + "") + "hello")) + ("" + "")) + "")) {
    case "foo bar":
        switch -9223372036854775808 {
        case -173:
            for fi102 := 0; fi102 < 1; fi102++ {
                if v46 {
                    v4 = ((0.1 + (float32(float64(v4)) / float32(1.5))) * ((1e10 / (v4 * 3.14159)) / (float32((1.5 * 3.14159)) - v4)))
                    var v103 int16 = 1000
                    var v104 uint8 = 42
                    _ = v103
                    _ = v104
                }
                _ = fi102
            }
        case -1:
            fmt.Println(v83, v2)
        case -9223372036854775808:
            if (32767 == int16(v50)) {
                switch "a" {
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    var v105 float64 = (((3.14159 / ((3.14159 / ((((1e10 + 0.1) - float64(0.0)) / float64(1.5)) - 1e10)) / ((0.1 / 1e10) / float64((-2.25 * v4))))) + 1.5) + 3.14159)
                    _ = v105
                }
            }
        }
    case "a":
        fmt.Println(v2, v50, v66)
    case "\n\t":
        if (-776 < 42) {
            _ = len(v1)
            fmt.Println(v4, v46, v3)
            v50 -= v50
        }
    default:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v66:", v66)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
    _ = v66
    _ = v83
    _ = v92
}
