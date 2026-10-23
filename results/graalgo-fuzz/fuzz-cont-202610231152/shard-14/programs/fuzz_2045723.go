package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 int64; F1 int64; F2 uint8 }
type GenS12 struct { F0 uint8; F1 uint16 }
type GenS13 struct { F0 uint32; F1 int; F2 uint }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int8; F1 uint; F2 uint64 }
func genFn0(p0 uint8) uint8 {
    var v7 int = -2147483648
    var v8 uint16 = 0
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(p0, v7)
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return (p0 ^ p0)
}
func genFn1(p0 int16) int8 {
    switch "hello" {
    case "\n\t":
        var v10 uint = 65535
        _ = v10
    case "":
        p0 = (^(p0))
    case "hello":
        p0 += p0
    default:
        var v11 int8 = 0
        _ = v11
    }
    if true {
        v12 := []int{42, 156, 483}
        _ = v12
    }
    p0 = p0
    _ = p0
    return 127
}

func main() {
    v1 := []string{"", "a"}
    var v2 uint32 = 2
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 766
    v5 := make(map[int]bool)
    v6 := make(map[int]string)
    v13 := GenS1{F0: 0, F1: 42, F2: v4}
    v4 = (v4 << ((1) & 63))
    var v14 int16 = 1
    for fi15 := 0; fi15 < 8; fi15++ {
        v2 = (^(v2))
        v16 := GenS1{F0: -9223372036854775808, F1: (^(v14)), F2: (v4 + v4)}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            v2 += v2
            _ = fi18
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        if true {
            v20 := GenS2{F0: fi19, F1: 36}
            v21 := v6[fi19]
            _ = v20
            _ = v21
        }
        _ = len(v1)
        if false {
            v4 = 42
            for fi22 := 0; fi22 < 9; fi22++ {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                _ = fi22
            }
            if (!(-2.25 >= 3.14159)) {
                v23 := GenS2{F0: fi19, F1: 1}
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    var v24 string = "\n\t"
                    _ = v24
                }
                var v25 uint = 437
                _ = v23
                _ = v25
            } else {
                var v26 uint32 = (v2 - (v2 + (v2 - v2)))
                _ = v26
            }
        } else {
            switch "a" {
            case "a":
                if true {
                    if (-2.25 <= 0.0) {
                        if false {
                            fmt.Println(v14)
                            v14 -= (v14 >> ((0) & 63))
                        } else {
                            for fi27 := 0; fi27 < 2; fi27++ {
                                v1 = append(v1, "\n\t")
                                var v28 int = fi19
                                switch fi27 {
                                case 855:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case -1:
                                    v29 := GenS3{F0: (^(v2)), F1: 2, F2: 413}
                                    v4 -= v4
                                    _ = v29
                                }
                                if (42 >= -9223372036854775808) {
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    var v30 int64 = -1
                                    v30 = (v30 / ((1) | 1))
                                    _ = v30
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = ((("" + ("" + "")) + ("" + "hello")) + ("" + (("foo bar" + ("" + ("" + "foo bar"))) + "a")))
                                    }
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(fi27, fi19)
                                    default:
                                        if (!true) {
                                            var v31 int32 = 1
                                            _ = v31
                                        }
                                    }
                                    v4 += 173
                                }
                                _ = fi27
                                _ = v28
                            }
                            var v32 float32 = (1.5 * 0.1)
                            fmt.Println(v14)
                            _ = v32
                        }
                        var v33 int32 = -325
                        _ = v33
                    } else {
                        for fi34 := 0; fi34 < 8; fi34++ {
                            if false {
                                fmt.Println(v2)
                                v2 = (v2 + 7)
                            } else {
                                var v35 uint8 = 127
                                fmt.Println(fi19)
                                if true {
                                    v36 := GenS4{F0: (v35 | (v35 | v35))}
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        if (!true) {
                                            var v37 int32 = -246
                                            v38 := GenS5{F0: "foo bar"}
                                            _ = v37
                                            _ = v38
                                        }
                                    default:
                                        fmt.Println(fi34, fi34, v14)
                                    }
                                    var v39 bool = (!(!(false || false)))
                                    _ = v36
                                    _ = v39
                                }
                                var v40 string = "hello"
                                fmt.Println(fi19, v4)
                                _ = v35
                                _ = v40
                            }
                            v14 = v14
                            _ = fi34
                        }
                        var v41 int8 = genFn1(-559)
                        _ = v41
                    }
                    var v42 uint = 65535
                    _ = v42
                } else {
                    switch "a" {
                    case "\n\t":
                        var v43 int32 = int32(fi19)
                        _ = v43
                    }
                    delete(v6, fi19)
                    for fi44 := 0; fi44 < 13; fi44++ {
                        v2 = v2
                        if (fi19 != (fi19 + 0)) {
                            var v45 bool = true
                            _ = v45
                        }
                        _ = fi44
                    }
                }
            case "hello":
                v6[-9223372036854775808] = (("a" + "a") + "a")
            default:
                v14 += (v14 & 601)
            }
            v46 := GenS6{F0: (18446744073709551615 > 18446744073709551615)}
            fmt.Println(fi19, v2)
            _ = v46
        }
        _ = fi19
    }
    switch -9223372036854775808 {
    case 32767:
        for fi47 := 0; fi47 < 13; fi47++ {
            v2 = 1
            var v48 int16 = v14
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            _ = fi47
            _ = v48
        }
    case -9223372036854775808:
        for fi49 := 0; fi49 < 7; fi49++ {
            switch "" {
            case "\n\t":
                for fi50 := 0; fi50 < 16; fi50++ {
                    if ((0.1 + ((((0.1 * float32(0.0)) / float32((1e10 - float64((1.5 - (((float32(float64(0.1)) * 0.0) + 0.0) + float32(-2.25))))))) - 3.14159) + (1.5 / 0.1))) > -2.25) {
                        var v51 uint64 = v4
                        delete(v3, 0)
                        v52 := GenS7{F0: 255, F1: (fi50 | fi50)}
                        _ = v51
                        _ = v52
                    } else {
                        if true {
                            switch "a" {
                            case "foo bar":
                                if true {
                                    if true {
                                        v53 := GenS5{F0: ("" + "foo bar")}
                                        _ = v53
                                    }
                                    if (291 == 0) {
                                        fmt.Println(v4)
                                    } else {
                                        var v54 uint16 = 255
                                        v4 -= (^(v4))
                                        switch ("a" + "") {
                                        case "":
                                            fmt.Println(fi50, fi49)
                                        }
                                        _ = v54
                                    }
                                    v55 := GenS8{F0: (v14 | v14), F1: genFn0(207)}
                                    _ = v55
                                } else {
                                    if true {
                                        v14 = v14
                                    } else {
                                        fmt.Println(v4)
                                        v56 := GenS2{F0: fi49, F1: 1}
                                        v14 -= 42
                                        _ = v56
                                    }
                                }
                                _ = len(v6)
                            case "":
                                if (0 <= 0) {
                                    switch "hello" {
                                    case "foo bar":
                                        var v57 int64 = -9223372036854775808
                                        _ = v57
                                    case "\n\t":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "hello":
                                        v1 = append(v1, (("foo bar" + "") + ("a" + "a")))
                                    }
                                    var v58 uint8 = 212
                                    _ = v58
                                }
                            case "a":
                                switch "a" {
                                case "hello":
                                    var v59 int32 = 163
                                    _ = v59
                                case "a":
                                    var v60 int = fi50
                                    _ = v60
                                }
                            default:
                                v61 := GenS7{F0: 18446744073709551615, F1: (fi50 ^ 581)}
                                _ = v61
                            }
                            var v62 float32 = 3.14159
                            _ = v62
                        } else {
                            if (0 > v2) {
                                delete(v5, fi50)
                                v63 := GenS3{F0: v2, F1: 205, F2: 889}
                                fmt.Println(fi49)
                                _ = v63
                            }
                            v64 := GenS9{F0: int(v4)}
                            v65 := v5[fi49]
                            _ = v64
                            _ = v65
                        }
                        var v66 int8 = 1
                        _ = v66
                    }
                    switch 127 {
                    case 205:
                        v1 = append(v1, ("foo bar" + ""))
                    default:
                        if true {
                            fmt.Println(v4, v14)
                            _ = len(v1)
                        } else {
                            v67 := GenS10{F0: v14, F1: 272}
                            if true {
                                var v68 float32 = 1e10
                                var v69 string = "foo bar"
                                var v70 float32 = 3.14159
                                switch v69 {
                                case "foo bar":
                                    if (!true) {
                                        var v71 string = ""
                                        var v72 int16 = v14
                                        _ = v71
                                        _ = v72
                                    }
                                case "hello":
                                    var v73 int64 = -9223372036854775808
                                    _ = v73
                                default:
                                    v74 := GenS11{F0: -910, F1: -9223372036854775808, F2: 2}
                                    _ = v74
                                }
                                if true {
                                    switch -100 {
                                    case 2147483647:
                                        fmt.Println(v14, v69, fi49)
                                    default:
                                        v4 += (^(v4))
                                    }
                                    var v75 uint16 = 7
                                    fmt.Println(v4, v14, v68)
                                    _ = v75
                                }
                                v2 += v2
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                            _ = v67
                        }
                    }
                    _ = fi50
                }
            case "hello":
                delete(v5, fi49)
            default:
                switch -993 {
                case -1:
                    for fi76 := 0; fi76 < 14; fi76++ {
                        fmt.Println(fi49)
                        v4 -= (v4 | v4)
                        v14 = (v14 ^ v14)
                        _ = fi76
                    }
                default:
                    var v77 int32 = -2147483648
                    _ = v77
                }
            }
            var v78 uint8 = 128
            var v79 bool = true
            _ = fi49
            _ = v78
            _ = v79
        }
    case 9223372036854775807:
        for fi80 := 0; fi80 < 19; fi80++ {
            switch 127 {
            case -2147483648:
                if true {
                    v81 := GenS9{F0: (-(fi80))}
                    _ = v81
                } else {
                    v82 := v3[uint32(v4)]
                    v1 = append(v1, ("" + ("" + "foo bar")))
                    v14 = (v14 + v14)
                    _ = v82
                }
            case 128:
                v83 := GenS12{F0: 42, F1: 1}
                _ = v83
            case 127:
                for fi84 := 0; fi84 < 14; fi84++ {
                    var v85 int = fi84
                    v5[0] = false
                    _ = fi84
                    _ = v85
                }
            }
            fmt.Println(v4, v4, v4)
            _ = fi80
        }
    }
    v14 -= v14
    switch "" {
    case "\n\t":
        v2 += (^(v2))
    case "":
        fmt.Println(v4, v2, v14)
    default:
        v4 -= v4
    }
    for fi86 := 0; fi86 < 19; fi86++ {
        if false {
            v87 := GenS13{F0: 0, F1: fi86, F2: 1}
            _ = v87
        } else {
            var v88 int64 = 9223372036854775807
            switch (fi86 ^ (fi86 & (fi86 ^ fi86))) {
            case -810:
                for fi89 := 0; fi89 < 11; fi89++ {
                    var v90 bool = (65535 >= 32767)
                    v88 = v88
                    v91 := GenS14{F0: 32767}
                    _ = fi89
                    _ = v90
                    _ = v91
                }
            case 108:
                switch fi86 {
                case 32767:
                    fmt.Println(v2)
                case -9223372036854775808:
                    v92 := GenS2{F0: fi86, F1: (v88 / ((v88) | 1))}
                    _ = v92
                case -463:
                    switch 2147483647 {
                    case 0:
                        v4 = 1
                    case 42:
                        switch 128 {
                        case 42:
                            delete(v6, fi86)
                        default:
                            v4 += v4
                        }
                    default:
                        for fi93 := 0; fi93 < 5; fi93++ {
                            fmt.Println(v2, fi86)
                            _ = fi93
                        }
                    }
                }
            }
            _ = v88
        }
        _ = fi86
    }
    var v94 int32 = 1
    if true {
        for fi95 := 0; fi95 < 8; fi95++ {
            var v96 uint64 = (v4 - (v4 << ((1) & 63)))
            switch (fi95 >> ((1) & 63)) {
            case 127:
                if false {
                    switch (("\n\t" + "a") + "") {
                    case "foo bar":
                        var v97 uint8 = 127
                        _ = v97
                    default:
                        v96 += (v4 - 0)
                    }
                } else {
                    v98 := GenS6{F0: false}
                    _ = v98
                }
            case 42:
                fmt.Println(v14, v14)
            case -20:
                switch "hello" {
                case "a":
                    fmt.Println(fi95, v94)
                case "":
                    for fi99 := 0; fi99 < 3; fi99++ {
                        v100 := GenS15{F0: 127, F1: 188, F2: v96}
                        _ = fi99
                        _ = v100
                    }
                case "hello":
                    v4 = (v4 << ((329) & 63))
                }
            default:
                v94 -= (v94 - (v94 ^ (v94 >> ((4294967295) & 63))))
            }
            _ = fi95
            _ = v96
        }
    }
    v101 := GenS11{F0: 1, F1: 898, F2: 127}
    v94 += (v94 + -724)
    v102 := GenS2{F0: -58, F1: 1}
    for fi103 := 0; fi103 < 8; fi103++ {
        for fi104 := 0; fi104 < 11; fi104++ {
            var v105 bool = true
            _ = fi104
            _ = v105
        }
        v106 := GenS6{F0: true}
        _ = fi103
        _ = v106
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v94:", v94)
    _ = v1
    _ = v101
    _ = v102
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
}
