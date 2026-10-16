package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int; F1 uint64; F2 uint }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 int32; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16 }
func genFn0(p0 int) uint32 {
    var v5 int32 = 774
    v6 := GenS1{F0: 1, F1: -1}
    v7 := GenS1{F0: -104, F1: -2147483648}
    var v8 float64 = 3.14159
    if true {
        var v9 bool = (!((0 > 65535) && (9 > 0)))
        v10 := GenS1{F0: 1, F1: (v5 ^ -1)}
        _ = v10
        _ = v9
    } else {
        if true {
            var v11 uint16 = 1
            var v12 int64 = 1
            _ = v11
            _ = v12
        }
        v13 := GenS1{F0: 1, F1: v5}
        _ = v13
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int64, p1 int64) int32 {
    var v14 float32 = (3.14159 + (-2.25 - 1.5))
    v15 := GenS2{F0: -1}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 2147483647
}

func main() {
    v1 := []string{("\n\t" + (("\n\t" + "") + "hello")), "foo bar", "foo bar", "\n\t"}
    var v2 uint16 = 42
    v3 := []string{"foo bar"}
    var v4 float64 = 1e10
    v16 := GenS1{F0: -1, F1: 1}
    fmt.Println(v4, v4)
    if true {
        v2 -= v2
        if (v2 != 0) {
            fmt.Println(v2, v4)
            switch "foo bar" {
            case "":
                v17 := GenS3{F0: 255, F1: v2, F2: 42}
                _ = v17
            case "\n\t":
                for fi18 := 0; fi18 < 7; fi18++ {
                    if ((false || false) && (2 > 257)) {
                        v19 := GenS4{F0: 128, F1: 2, F2: 2}
                        _ = v19
                    }
                    _ = fi18
                }
            default:
                for fi20 := 0; fi20 < 10; fi20++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch "hello" {
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "foo bar":
                        for fi21 := 0; fi21 < 16; fi21++ {
                            v4 = v4
                            _ = fi21
                        }
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi20
                }
            }
            var v22 int64 = -9223372036854775808
            _ = v22
        }
    }
    v2 = v2
    if true {
        fmt.Println(v2)
    } else {
        if true {
            v23 := GenS3{F0: 85, F1: 2, F2: 766}
            v4 -= float64(float32(v4))
            switch "a" {
            case "a":
                var v24 int64 = int64(v2)
                _ = v24
            default:
                for fi25 := 0; fi25 < 5; fi25++ {
                    if true {
                        var v26 int32 = genFn1(-794, 0)
                        _ = v26
                    } else {
                        var v27 int16 = 130
                        _ = v27
                    }
                    _ = fi25
                }
            }
            _ = v23
        }
        var v28 int64 = 898
        switch "hello" {
        case "a":
            _ = len(v3)
        }
        _ = v28
    }
    var v29 int8 = 127
    for fi30 := 0; fi30 < 19; fi30++ {
        v29 = (v29 | 42)
        v1 = append(v1, "hello")
        _ = fi30
    }
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "\n\t")
    }
    if false {
        v4 = ((v4 * v4) + (((v4 - (1.5 / 1e10)) / v4) - float64(((float32(v4) + -2.25) * 1.5))))
        for fi31 := 0; fi31 < 12; fi31++ {
            v2 -= 2
            switch 1000 {
            case 42:
                if (v2 >= (v2 & (v2 << ((80) & 63)))) {
                    var v32 uint = 42
                    v1 = append(v1, ("\n\t" + "hello"))
                    _ = v32
                } else {
                    _ = len(v1)
                    if (1 == 266) {
                        switch fi31 {
                        case 552:
                            var v33 bool = (648 == 18446744073709551615)
                            _ = v33
                        case 647:
                            v34 := GenS3{F0: 128, F1: (v2 >> ((65535) & 63)), F2: 2147483648}
                            _ = v34
                        }
                        v35 := GenS5{F0: 0, F1: -1}
                        _ = v35
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ("\n\t" + "")
                        }
                    }
                }
            case 1:
                v36 := GenS2{F0: 1}
                _ = v36
            default:
                v37 := GenS6{F0: -1, F1: 2, F2: 150}
                _ = v37
            }
            fmt.Println(v4, v29, fi31)
            _ = fi31
        }
        var v38 int16 = 1000
        _ = v38
    } else {
        var v39 uint16 = v2
        _ = v39
    }
    switch 32768 {
    case 2147483647:
        v40 := GenS7{F0: 1, F1: v2, F2: (v29 | int8(v2))}
        _ = v40
    case -9223372036854775808:
        v29 = (v29 ^ (v29 * (v29 + (-(v29)))))
    case -367:
        fmt.Println(v2, v2)
    }
    var v41 int64 = int64(v2)
    v42 := GenS4{F0: 9223372036854775807, F1: 851, F2: 2}
    v2 -= 1
    for fi43 := 0; fi43 < 11; fi43++ {
        if true {
            v44 := GenS4{F0: fi43, F1: 0, F2: 619}
            if (277 < 2) {
                if (838 > 295) {
                    if (fi43 < (fi43 ^ (fi43 * (fi43 / ((fi43) | 1))))) {
                        for fi45 := 0; fi45 < 3; fi45++ {
                            v46 := GenS1{F0: 42, F1: 42}
                            var v47 int64 = (-(v41))
                            switch fi43 {
                            case 32768:
                                v48 := GenS8{F0: v47}
                                _ = v48
                            case 9223372036854775807:
                                if false {
                                    var v49 int = 1
                                    _ = v49
                                } else {
                                    if (int32(v41) != 1) {
                                        switch fi45 {
                                        case 34:
                                            fmt.Println(fi43, fi43, fi45)
                                        case 127:
                                            v41 -= -1
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v50 uint16 = v2
                                        var v51 string = "a"
                                        _ = v50
                                        _ = v51
                                    } else {
                                        if (v41 <= 42) {
                                            var v52 int8 = 42
                                            var v53 int = -9223372036854775808
                                            if len(v1) > 0 {
                                                v1[0] = (("hello" + "hello") + "hello")
                                            }
                                            _ = v52
                                            _ = v53
                                        }
                                        v54 := GenS3{F0: 42, F1: v2, F2: 0}
                                        _ = v54
                                    }
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    var v55 bool = (!(-547 == -997))
                                    if (!false) {
                                        v56 := GenS9{F0: 908}
                                        v29 = 127
                                        if true {
                                            v41 = 203
                                            v57 := GenS3{F0: 5, F1: v2, F2: 2147483647}
                                            _ = v57
                                        } else {
                                            v58 := GenS6{F0: -2147483648, F1: 55, F2: int64(v29)}
                                            _ = v58
                                        }
                                        _ = v56
                                    }
                                    v29 = (v29 / ((81) | 1))
                                    _ = v55
                                }
                            }
                            var v59 bool = (((0.0 + float32((1e10 - v4))) < 3.14159) && true)
                            v2 += (^(v2))
                            var v60 int16 = 42
                            _ = fi45
                            _ = v46
                            _ = v47
                            _ = v59
                            _ = v60
                        }
                    }
                    v61 := GenS8{F0: (v41 % ((v41) | 1))}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v61
                } else {
                    if len(v1) > 0 {
                        v1[0] = (("foo bar" + ((((("" + ("a" + "\n\t")) + "a") + "") + "") + ("foo bar" + "foo bar"))) + "\n\t")
                    }
                    fmt.Println(v4, v2, v2)
                    var v62 int = fi43
                    _ = v62
                }
                for fi63 := 0; fi63 < 9; fi63++ {
                    var v64 int16 = 1000
                    var v65 uint16 = v2
                    switch ("hello" + (("" + ("hello" + ("" + ("" + ("hello" + "foo bar"))))) + (("a" + "foo bar") + (("" + "") + "")))) {
                    case "\n\t":
                        var v66 int = (^(fi43))
                        _ = v66
                    case "":
                        var v67 float32 = (1e10 / 1.0)
                        _ = v67
                    case "a":
                        v41 = v41
                    default:
                        if (32767 != v2) {
                            v4 = (1.5 - (((v4 + v4) - v4) + (0.1 / 3.14159)))
                        }
                    }
                    switch (fi63 % (((fi63 * fi63)) | 1)) {
                    case 642:
                        var v68 uint32 = 1
                        _ = v68
                    case -1:
                        v2 = v65
                    default:
                        if true {
                            var v69 int8 = -113
                            var v70 uint64 = 1
                            v65 -= v65
                            _ = v69
                            _ = v70
                        }
                    }
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                _ = len(v1)
            }
            _ = v44
        } else {
            v41 -= (v41 & v41)
        }
        _ = fi43
    }
    if false {
        v41 += 9223372036854775807
        var v71 int32 = 2147483647
        _ = v71
    } else {
        v29 = v29
        var v72 float64 = 0.0
        _ = v72
    }
    switch 4294967295 {
    case 676:
        var v73 uint = 4294967295
        _ = v73
    case -101:
        if len(v3) > 0 {
            v3[0] = ""
        }
    case 42:
        v74 := GenS6{F0: 1, F1: 2, F2: (v41 >> ((919) & 63))}
        _ = v74
    }
    if false {
        var v75 uint16 = (v2 >> ((4294967295) & 63))
        var v76 float64 = (1.5 - -2.25)
        _ = v75
        _ = v76
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
