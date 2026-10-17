package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 string; F1 int64; F2 int16 }
type GenS10 struct { F0 int8; F1 uint32; F2 int8 }
type GenS11 struct { F0 int; F1 uint; F2 uint }
type GenS12 struct { F0 uint8; F1 uint16; F2 string }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 int64 = 1
    v4 := []int{9223372036854775807, 128, -2147483648, -145}
    v5 := []int{32767}
    if (101 != 250) {
        for fi6 := 0; fi6 < 10; fi6++ {
            if true {
                v7 := GenS1{F0: (fi6 & fi6), F1: 0}
                _ = v7
            }
            var v8 uint8 = 202
            _ = fi6
            _ = v8
        }
        v3 = 42
        if (32768 > 128) {
            delete(v2, 9223372036854775807)
        } else {
            _ = len(v1)
        }
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS2{F0: fi9, F1: 0}
        var v11 string = "\n\t"
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS3{F0: 2147483648, F1: 2}
    var v13 uint = 4294967295
    _ = len(v2)
    switch 32767 {
    case -72:
        if (1 > v3) {
            var v14 int8 = -76
            var v15 uint8 = 0
            _ = v14
            _ = v15
        }
    }
    v13 += 65535
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(fi16)
        v17 := GenS2{F0: fi16, F1: 255}
        v3 = (^(v3))
        _ = fi16
        _ = v17
    }
    if false {
        switch "\n\t" {
        case "a":
            for fi18 := 0; fi18 < 17; fi18++ {
                if (2 <= 18446744073709551615) {
                    if (2147483648 != 795) {
                        v3 -= (^(v3))
                        switch (("\n\t" + "hello") + "a") {
                        case "hello":
                            switch int(v3) {
                            case 827:
                                switch "\n\t" {
                                case "hello":
                                    fmt.Println(v13, fi18)
                                case "foo bar":
                                    v19 := v2[(fi18 & -9223372036854775808)]
                                    _ = v19
                                default:
                                    for fi20 := 0; fi20 < 2; fi20++ {
                                        v21 := GenS4{F0: -4, F1: 4294967295}
                                        var v22 int32 = -1
                                        _ = fi20
                                        _ = v21
                                        _ = v22
                                    }
                                }
                            case 42:
                                v3 += v3
                            case 2147483647:
                                var v23 float64 = ((((1e10 / 1e10) - (0.1 * float64((-2.25 * 1e10)))) / -2.25) - ((1.5 - (1.5 * (3.14159 / 3.14159))) + ((-2.25 + ((1e10 + (3.14159 / ((((-2.25 - ((0.1 + float64(1e10)) + (0.1 + 1e10))) + -2.25) + 0.1) - -2.25))) * 0.1)) * 0.1)))
                                _ = v23
                            default:
                                if (163 != 17) {
                                    v3 = (v3 / ((-1) | 1))
                                } else {
                                    v24 := GenS1{F0: -935, F1: 190}
                                    _ = v24
                                }
                            }
                        case "":
                            if ("" != ("foo bar" + (("foo bar" + "") + (("foo bar" + "") + (("hello" + "hello") + "\n\t"))))) {
                                var v25 uint64 = 18446744073709551615
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    if (3.14159 > (1e10 / 1e10)) {
                                        if false {
                                            v27 := GenS5{F0: -96}
                                            v28 := GenS6{F0: (v25 - v25), F1: 2}
                                            var v29 bool = (0 == 65535)
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                        if (4 != 0) {
                                            v13 += (v13 >> (((^(v13))) & 63))
                                            var v30 uint16 = 1
                                            _ = v30
                                        } else {
                                            v31 := GenS3{F0: uint32(v13), F1: v25}
                                            var v32 string = ""
                                            _ = v31
                                            _ = v32
                                        }
                                        switch (("\n\t" + "\n\t") + ("\n\t" + ("foo bar" + ("" + "\n\t")))) {
                                        case "a":
                                            v1[fi26] = "foo bar"
                                        case "\n\t":
                                            if len(v5) > 0 {
                                                v5[0] = fi18
                                            }
                                        case "foo bar":
                                            v3 = (v3 - v3)
                                        default:
                                            if false {
                                                if false {
                                                    v3 = v3
                                                    v3 -= (v3 >> (((v13 & (v13 - (v13 * 500)))) & 63))
                                                }
                                            }
                                        }
                                    } else {
                                        v33 := v1[128]
                                        var v34 string = v33
                                        _ = v33
                                        _ = v34
                                    }
                                    switch (fi26 ^ (fi18 + -2147483648)) {
                                    case 32767:
                                        if (!false) {
                                            if ((float32(float64(0.0)) + float32(0.1)) <= -2.25) {
                                                v25 += v25
                                                if (1 >= 395) {
                                                    v3 = -704
                                                    v35 := GenS3{F0: 42, F1: (v25 | 42)}
                                                    _ = v35
                                                }
                                                delete(v1, 9223372036854775807)
                                                if false {
                                                    v36 := GenS3{F0: 2, F1: v25}
                                                    if ((212 <= 1) && true) {
                                                        var v37 int32 = 0
                                                        v25 += v25
                                                        _ = v37
                                                    }
                                                    fmt.Println(fi18, v25)
                                                    _ = v36
                                                }
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = fi26
                                                }
                                                if ((-32768 <= 42) && (300 == v3)) {
                                                    var v38 int16 = 1
                                                    var v39 uint16 = 1
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                            switch 9223372036854775807 {
                                            case 214:
                                                if ((230 >= 222) && true) {
                                                    if (((-128 <= -1) || false) || (0 < 934)) {
                                                        var v40 int8 = int8(v25)
                                                        v13 = v13
                                                        if (true && false) {
                                                            var v41 float32 = (-2.25 / 1.0)
                                                            fmt.Println(v13)
                                                            _ = v41
                                                        } else {
                                                            v3 = v3
                                                        }
                                                        _ = v40
                                                    } else {
                                                        v1[(fi26 - fi18)] = ""
                                                    }
                                                    v5 = append(v5, fi18)
                                                    switch (("a" + "hello") + (("hello" + "a") + (("foo bar" + "") + "\n\t"))) {
                                                    case "hello":
                                                        var v42 int64 = (v3 & (v3 >> ((v13) & 63)))
                                                        _ = v42
                                                    case "\n\t":
                                                        var v43 string = ""
                                                        _ = v43
                                                    case "":
                                                        fmt.Println(fi18, v25)
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            case 4294967295:
                                                if len(v5) > 0 {
                                                    v5[0] = fi18
                                                }
                                            case -98:
                                                switch 32768 {
                                                case 0:
                                                    v44 := GenS3{F0: 2147483648, F1: (v25 << ((65535) & 63))}
                                                    _ = v44
                                                case 32767:
                                                    v45 := GenS7{F0: (v25 | (v25 >> (((v13 - 255)) & 63))), F1: (!(2 >= 692))}
                                                    switch "a" {
                                                    case "a":
                                                        switch "foo bar" {
                                                        case "a":
                                                            v46 := GenS8{F0: 0, F1: 0}
                                                            _ = v46
                                                        case "hello":
                                                            var v47 uint8 = 128
                                                            _ = v47
                                                        }
                                                    case "foo bar":
                                                        switch "a" {
                                                        case "foo bar":
                                                            v3 += (v3 & v3)
                                                        default:
                                                            var v48 float64 = 0.0
                                                            _ = v48
                                                        }
                                                    case "":
                                                        var v49 uint = 357
                                                        _ = v49
                                                    }
                                                    var v50 float32 = 1.5
                                                    _ = v45
                                                    _ = v50
                                                case 9223372036854775807:
                                                    v51 := GenS8{F0: 163, F1: (v3 * (v3 | 0))}
                                                    _ = v51
                                                default:
                                                    if false {
                                                        if len(v4) > 0 {
                                                            v4[0] = fi26
                                                        }
                                                        var v52 int32 = 0
                                                        fmt.Println(v13, v25)
                                                        _ = v52
                                                    } else {
                                                        v25 += 712
                                                        v13 = (v13 & (v13 / ((544) | 1)))
                                                    }
                                                }
                                            default:
                                                var v53 uint8 = 43
                                                _ = v53
                                            }
                                            v25 += v25
                                        } else {
                                            if (true && true) {
                                                v54 := GenS9{F0: ("a" + ("foo bar" + "hello")), F1: int64(v25), F2: int16(v13)}
                                                _ = v54
                                            }
                                            fmt.Println(v13, fi26, fi26)
                                            v25 = v25
                                        }
                                    default:
                                        fmt.Println(v13, v13, v13)
                                    }
                                    v25 -= (v25 / ((v25) | 1))
                                    v25 = 630
                                    v55 := GenS10{F0: -57, F1: 181, F2: -1}
                                    var v56 uint64 = (v25 << ((v13) & 63))
                                    _ = len(v1)
                                    if (true && (!true)) {
                                        v56 -= v25
                                        fmt.Println(fi18)
                                        delete(v2, (fi18 << ((0) & 63)))
                                        v2[-836] = 32767
                                        v57 := GenS11{F0: 2147483647, F1: (^(v13)), F2: (v13 | 4294967295)}
                                        _ = v57
                                    } else {
                                        var v58 int64 = 9223372036854775807
                                        _ = v58
                                    }
                                    _ = fi26
                                    _ = v55
                                    _ = v56
                                }
                                v59 := v2[fi18]
                                _ = v25
                                _ = v59
                            }
                        default:
                            for fi60 := 0; fi60 < 6; fi60++ {
                                switch "" {
                                case "\n\t":
                                    if (!("a" >= ("foo bar" + (("foo bar" + ((("foo bar" + "hello") + "") + "")) + "hello")))) {
                                        v13 = (v13 ^ (^(v13)))
                                    }
                                case "a":
                                    var v61 int16 = -1
                                    _ = v61
                                case "foo bar":
                                    v62 := GenS12{F0: 24, F1: 2, F2: "a"}
                                    fmt.Println(v13, fi60, fi18)
                                    _ = v62
                                default:
                                    v13 = (v13 / ((v13) | 1))
                                }
                                v13 = v13
                                _ = fi60
                            }
                        }
                        _ = len(v1)
                    } else {
                        v13 = v13
                    }
                    v63 := v1[-1]
                    _ = v63
                } else {
                    v13 = (v13 * v13)
                    _ = len(v2)
                }
                for fi64 := 0; fi64 < 6; fi64++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi64
                }
                _ = fi18
            }
        case "hello":
            var v65 float32 = ((0.1 + 1.5) / 1e10)
            _ = v65
        case "foo bar":
            v66 := GenS5{F0: 127}
            _ = v66
        }
        var v67 uint16 = 42
        _ = v67
    }
    switch 128 {
    case 128:
        var v68 uint32 = 4294967295
        _ = v68
    default:
        var v69 uint = (v13 % (((v13 / ((v13) | 1))) | 1))
        _ = v69
    }
    var v70 int8 = -86
    fmt.Println(v70)
    for fi71 := 0; fi71 < 3; fi71++ {
        fmt.Println(fi71, v3)
        var v72 int64 = 42
        if ((true || (!true)) && false) {
            var v73 int = (fi71 / (((fi71 * 42)) | 1))
            for fi74 := 0; fi74 < 5; fi74++ {
                var v75 int16 = 1000
                _ = fi74
                _ = v75
            }
            var v76 int16 = 1000
            _ = v73
            _ = v76
        } else {
            for fi77 := 0; fi77 < 2; fi77++ {
                switch fi77 {
                case -1:
                    var v78 float64 = 0.0
                    _ = v78
                case -2147483648:
                    if ((^(v13)) > 1) {
                        var v79 string = ""
                        _ = v79
                    }
                default:
                    var v80 int8 = v70
                    _ = v80
                }
                _ = fi77
            }
            v4 = append(v4, (fi71 % (((fi71 + (^(fi71)))) | 1)))
            if true {
                _ = len(v5)
                var v81 int16 = 32767
                v70 = 1
                _ = v81
            } else {
                var v82 float64 = float64(-2.25)
                _ = v82
            }
        }
        _ = fi71
        _ = v72
    }
    if true {
        if true {
            var v83 float64 = (0.1 + 0.0)
            _ = v83
        }
    }
    v70 = v70
    if true {
        for fi84 := 0; fi84 < 4; fi84++ {
            v13 = 1
            _ = fi84
        }
        for fi85 := 0; fi85 < 17; fi85++ {
            v5 = append(v5, (fi85 ^ fi85))
            v70 = (v70 | 0)
            v70 += (v70 & v70)
            _ = fi85
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
}
