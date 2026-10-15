package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 string; F1 int64; F2 bool }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 uint; F1 uint64; F2 bool }
func genFn0(p0 int) uint32 {
    p0 = (p0 * (p0 | -1))
    switch "a" {
    case "":
        var v7 uint16 = 1
        _ = v7
    case "\n\t":
        var v8 uint16 = 1
        _ = v8
    case "hello":
        v9 := GenS1{F0: 65535, F1: 212}
        _ = v9
    default:
        p0 += (^(p0))
    }
    var v10 uint16 = 786
    _ = p0
    _ = v10
    return 42
}
func genFn1() int {
    v11 := GenS2{F0: 86}
    if true {
        var v12 string = ("hello" + "foo bar")
        var v13 int32 = 2147483647
        _ = v12
        _ = v13
    }
    var v14 float64 = (((1.5 + (((-2.25 / -2.25) / 1e10) / ((float64(1.5) + (float64(3.14159) + -2.25)) * ((((1.5 / (0.0 - -2.25)) * ((1.5 * 0.0) * ((-2.25 - 1.5) * 0.0))) / ((0.0 + 0.1) / 1.5)) + 1.5)))) + -2.25) - (-2.25 * 1.5))
    _ = v11
    _ = v14
    return 611
}

func main() {
    v1 := []string{("a" + ("" + "foo bar")), "foo bar", "a"}
    var v2 uint64 = 1
    var v3 string = "a"
    var v4 float32 = (float32(((0.0 + 0.1) + (0.1 + (-2.25 * (0.1 + ((-2.25 * 1e10) + 3.14159)))))) / (-2.25 - 0.0))
    v5 := []string{(v3 + ""), ("\n\t" + ("\n\t" + ((v3 + v3) + "\n\t"))), "\n\t", v3}
    var v6 bool = (!false)
    var v15 uint8 = 0
    if v6 {
        switch v3 {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            _ = len(v5)
        }
        switch v3 {
        case "\n\t":
            if v6 {
                v16 := GenS2{F0: -128}
                v1 = append(v1, (((v3 + v3) + v3) + ("foo bar" + "")))
                for fi17 := 0; fi17 < 14; fi17++ {
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v3 = "foo bar"
                        var v19 int8 = -55
                        var v20 float32 = v4
                        v20 = 1.5
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    v3 = v3
                    for fi21 := 0; fi21 < 2; fi21++ {
                        _ = len(v1)
                        _ = fi21
                    }
                    _ = fi17
                }
                _ = v16
            } else {
                var v22 int = -160
                if v6 {
                    v22 -= v22
                    if true {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            v24 := GenS3{F0: (v3 + (v3 + ("a" + v3)))}
                            for fi25 := 0; fi25 < 5; fi25++ {
                                v6 = (float32(1e10) == -2.25)
                                v26 := GenS1{F0: 1, F1: int64(fi25)}
                                v15 = (v15 % (((^(v15))) | 1))
                                v22 = genFn1()
                                fmt.Println(fi23, v6, v22)
                                v27 := GenS4{F0: 0}
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            var v28 uint = 478
                            _ = fi23
                            _ = v24
                            _ = v28
                        }
                    }
                    var v29 int64 = 1
                    _ = v29
                } else {
                    var v30 uint8 = v15
                    if v6 {
                        for fi31 := 0; fi31 < 16; fi31++ {
                            v3 = ("hello" + "hello")
                            v15 = 21
                            _ = fi31
                        }
                        if v6 {
                            if (42 < v2) {
                                for fi32 := 0; fi32 < 4; fi32++ {
                                    fmt.Println(v6, fi32, v30)
                                    v33 := GenS3{F0: ""}
                                    v4 += v4
                                    _ = fi32
                                    _ = v33
                                }
                                switch v3 {
                                case "\n\t":
                                    switch v3 {
                                    case "a":
                                        v5 = append(v5, v3)
                                    case "\n\t":
                                        var v34 float64 = 0.1
                                        _ = v34
                                    default:
                                        var v35 int16 = 0
                                        _ = v35
                                    }
                                default:
                                    _ = len(v1)
                                }
                                if (-1 >= -20) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    v6 = (!(19 != v2))
                                    v36 := GenS3{F0: v3}
                                    switch -47 {
                                    case 788:
                                        if ((!(!v6)) || (v6 && v6)) {
                                            v37 := GenS3{F0: v3}
                                            var v38 bool = v6
                                            var v39 int = genFn1()
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                    case 1:
                                        v30 -= v15
                                    default:
                                        var v40 uint = 65535
                                        _ = v40
                                    }
                                    _ = v36
                                }
                            }
                            v2 += (^(v2))
                        } else {
                            var v41 uint16 = 241
                            _ = v41
                        }
                    }
                    _ = v30
                }
                _ = v22
            }
        case "foo bar":
            v42 := GenS3{F0: (("foo bar" + v3) + v3)}
            _ = v42
        }
    }
    fmt.Println(v2)
    v4 -= 3.14159
    v43 := GenS5{F0: 0, F1: 42}
    var v44 int16 = 843
    if (!v6) {
        v45 := GenS6{F0: ((v3 + "\n\t") + "hello"), F1: 42, F2: v6}
        v4 += (v4 - ((0.1 - v4) - float32(-2.25)))
        var v46 int = -321
        _ = v45
        _ = v46
    }
    v44 = (v44 & v44)
    if (!false) {
        if (v6 && (42 > 522)) {
            switch v3 {
            case "":
                if (v6 || v6) {
                    v3 = v3
                    v15 -= (v15 * (v15 / ((255) | 1)))
                    switch v3 {
                    case "":
                        for fi47 := 0; fi47 < 18; fi47++ {
                            var v48 uint64 = v2
                            fmt.Println(v44, v3, v3)
                            if v6 {
                                for fi49 := 0; fi49 < 8; fi49++ {
                                    fmt.Println(fi49)
                                    v15 = v15
                                    _ = fi49
                                }
                            } else {
                                _ = len(v5)
                                _ = len(v5)
                                v44 = (v44 ^ 0)
                            }
                            _ = fi47
                            _ = v48
                        }
                    case "foo bar":
                        var v50 bool = v6
                        _ = v50
                    }
                }
            case "foo bar":
                var v51 int16 = v44
                _ = v51
            }
            var v52 string = v3
            _ = v52
        } else {
            if (true && (980 > 0)) {
                switch 1 {
                case -9223372036854775808:
                    v53 := GenS7{F0: -2147483648, F1: (v44 | v44)}
                    _ = v53
                case 9223372036854775807:
                    switch v3 {
                    case "a":
                        var v54 int8 = 0
                        _ = v54
                    case "hello":
                        var v55 uint32 = 65535
                        _ = v55
                    default:
                        switch (v3 + "hello") {
                        case "\n\t":
                            if (true || true) {
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    var v56 bool = ((!v6) && (100 > 97))
                                    _ = v56
                                case "a":
                                    var v57 uint32 = genFn0(4294967295)
                                    _ = v57
                                case "hello":
                                    v5 = append(v5, v3)
                                default:
                                    var v58 uint64 = (^(v2))
                                    _ = v58
                                }
                                for fi59 := 0; fi59 < 12; fi59++ {
                                    if (v3 > v3) {
                                        for fi60 := 0; fi60 < 9; fi60++ {
                                            v61 := GenS3{F0: v3}
                                            v15 += v15
                                            v44 = v44
                                            _ = fi60
                                            _ = v61
                                        }
                                    } else {
                                        v4 += 1e10
                                        if false {
                                            for fi62 := 0; fi62 < 7; fi62++ {
                                                v2 += (v2 << ((2) & 63))
                                                fmt.Println(fi62)
                                                _ = fi62
                                            }
                                            var v63 string = ""
                                            _ = v63
                                        }
                                        for fi64 := 0; fi64 < 16; fi64++ {
                                            var v65 uint = 172
                                            _ = fi64
                                            _ = v65
                                        }
                                    }
                                    v1 = append(v1, "a")
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi59
                                }
                            }
                        default:
                            var v66 uint = 65535
                            _ = v66
                        }
                    }
                }
                for fi67 := 0; fi67 < 17; fi67++ {
                    if v6 {
                        var v68 uint8 = (^(v15))
                        if v6 {
                            fmt.Println(v4, v44)
                            switch v3 {
                            case "":
                                if false {
                                    var v69 float32 = (v4 * 1e10)
                                    _ = v69
                                } else {
                                    var v70 int64 = -1
                                    _ = v70
                                }
                            default:
                                var v71 uint8 = (v15 - v15)
                                _ = v71
                            }
                        }
                        var v72 float64 = 1.5
                        _ = v68
                        _ = v72
                    }
                    if v6 {
                        var v73 float64 = 1.5
                        fmt.Println(v73, v15, v73)
                        v4 = v4
                        _ = v73
                    }
                    _ = fi67
                }
                fmt.Println(v4, v44)
            }
            fmt.Println(v2, v15)
            v5 = append(v5, v3)
        }
    }
    if len(v5) > 0 {
        v5[0] = ("a" + ("foo bar" + v3))
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v44 -= (v44 % ((int16(v2)) | 1))
    for fi74 := 0; fi74 < 18; fi74++ {
        var v75 float32 = v4
        _ = fi74
        _ = v75
    }
    fmt.Println(v2, v44)
    if (!v6) {
        v76 := GenS8{F0: 0, F1: 0, F2: v6}
        for fi77 := 0; fi77 < 1; fi77++ {
            var v78 uint64 = v2
            _ = fi77
            _ = v78
        }
        var v79 bool = v6
        _ = v76
        _ = v79
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            var v80 int32 = 761
            v6 = v6
            var v81 int64 = -1
            _ = v80
            _ = v81
        }
        switch v3 {
        case "hello":
            v6 = (0.0 == (((3.14159 / 3.14159) + 0.1) * float64((1.5 - v4))))
        case "foo bar":
            var v82 uint = 382
            _ = v82
        case "":
            var v83 int = 592
            _ = v83
        default:
            if (73 < -16) {
                v6 = v6
            } else {
                for fi84 := 0; fi84 < 13; fi84++ {
                    var v85 string = ""
                    switch "hello" {
                    case "":
                        var v86 uint = 733
                        _ = v86
                    default:
                        v3 = "\n\t"
                    }
                    v15 = (v15 & (v15 - 2))
                    _ = fi84
                    _ = v85
                }
                v87 := GenS7{F0: 0, F1: (v44 / ((v44) | 1))}
                _ = v87
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
