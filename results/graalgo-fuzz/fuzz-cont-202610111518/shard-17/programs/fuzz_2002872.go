package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int32; F2 string }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint32; F1 int; F2 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 bool; F2 int16 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 int8; F1 string }
func genFn0(p0 string, p1 int64) int16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(fi4, p1)
        fmt.Println(fi4, p1, p1)
        var v5 float32 = 1.5
        _ = fi4
        _ = v5
    }
    if false {
        p1 = -1
        var v6 uint = 65535
        var v7 uint32 = 65535
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(fi8, p0)
        v9 := GenS1{F0: 16}
        _ = fi8
        _ = v9
    }
    var v10 int = -765
    _ = p0
    _ = p1
    _ = v10
    return -940
}
func genFn1(p0 uint32) uint64 {
    p0 -= (p0 << ((255) & 63))
    p0 += p0
    p0 -= p0
    fmt.Println(p0)
    var v11 uint16 = 2
    _ = p0
    _ = v11
    return 314
}

func main() {
    var v1 uint8 = 2
    var v2 bool = false
    v3 := []int{42, 42}
    for fi12 := 0; fi12 < 13; fi12++ {
        if len(v3) > 0 {
            v3[0] = 128
        }
        var v13 uint32 = 707
        switch ("\n\t" + ("a" + "a")) {
        case "a":
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint32 = v13
                var v16 uint = 255
                _ = fi14
                _ = v15
                _ = v16
            }
        case "hello":
            v17 := GenS2{F0: 32767}
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v1, v2)
    switch -900 {
    case 859:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 1:
        _ = len(v3)
    }
    var v18 uint64 = 4294967295
    for fi19 := 0; fi19 < 18; fi19++ {
        v2 = (-9223372036854775808 <= (^(fi19)))
        _ = len(v3)
        v2 = (v2 || (0.1 <= float32(0.0)))
        _ = fi19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 uint8 = (v1 % ((v1) | 1))
        var v22 int8 = 0
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        v18 = v18
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS1{F0: 127}
            v26 := GenS1{F0: -128}
            var v27 bool = (!(v2 || v2))
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi23
    }
    if true {
        if (v2 && false) {
            v28 := GenS3{F0: -2147483648, F1: -534, F2: ""}
            var v29 float32 = 0.0
            var v30 bool = (!((1 < 42) && ((false && (-2.25 != (0.1 + 0.0))) || (2 < 95))))
            _ = v28
            _ = v29
            _ = v30
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 int32 = 0
                _ = fi31
                _ = v32
            }
        }
        fmt.Println(v18)
    } else {
        fmt.Println(v18)
        fmt.Println(v2, v2, v1)
    }
    v3 = append(v3, 4294967295)
    if false {
        v1 -= (v1 + 49)
        fmt.Println(v18, v2, v2)
    }
    if v2 {
        for fi33 := 0; fi33 < 9; fi33++ {
            for fi34 := 0; fi34 < 11; fi34++ {
                switch "a" {
                case "a":
                    var v35 string = ("foo bar" + "\n\t")
                    switch "" {
                    case "a":
                        v36 := GenS3{F0: -2147483648, F1: -2147483648, F2: "hello"}
                        _ = v36
                    default:
                        var v37 bool = v2
                        _ = v37
                    }
                    if len(v3) > 0 {
                        v3[0] = 575
                    }
                    _ = v35
                }
                v1 = 255
                _ = fi34
            }
            v38 := GenS3{F0: 573, F1: -1, F2: "hello"}
            _ = fi33
            _ = v38
        }
        v3 = append(v3, 32767)
    } else {
        if v2 {
            switch -241 {
            case 42:
                for fi39 := 0; fi39 < 3; fi39++ {
                    v40 := GenS1{F0: 0}
                    v41 := GenS3{F0: -491, F1: 2147483647, F2: "a"}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            }
            if (-896 != 2147483647) {
                fmt.Println(v1, v2)
                var v42 uint32 = 2
                _ = v42
            }
        }
        for fi43 := 0; fi43 < 20; fi43++ {
            switch "a" {
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                var v44 int8 = 42
                _ = v44
            case "foo bar":
                v2 = (3.14159 <= -2.25)
            default:
                v45 := GenS3{F0: -604, F1: -2147483648, F2: "foo bar"}
                _ = v45
            }
            _ = fi43
        }
    }
    switch "foo bar" {
    case "foo bar":
        if (!v2) {
            fmt.Println(v2, v18)
            var v46 int8 = 0
            _ = v46
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v47 bool = v2
            _ = v47
        }
    case "\n\t":
        if v2 {
            var v48 uint = 637
            var v49 uint64 = 1
            if len(v3) > 0 {
                v3[0] = 128
            }
            _ = v48
            _ = v49
        } else {
            var v50 int8 = 70
            _ = len(v3)
            _ = v50
        }
    default:
        v1 += (v1 ^ v1)
    }
    _ = len(v3)
    if (uint32(v1) == 855) {
        v2 = (v2 || false)
        v18 = genFn1(65535)
        if v2 {
            if false {
                v18 = v18
                for fi51 := 0; fi51 < 12; fi51++ {
                    v52 := GenS1{F0: -1}
                    if true {
                        for fi53 := 0; fi53 < 11; fi53++ {
                            var v54 uint = 798
                            v54 = 616
                            _ = fi53
                            _ = v54
                        }
                    }
                    _ = fi51
                    _ = v52
                }
            } else {
                v3 = append(v3, -2147483648)
                switch "\n\t" {
                case "":
                    if v2 {
                        switch "" {
                        case "hello":
                            var v55 bool = (!((uint16(v18) != 1) && v2))
                            _ = v55
                        default:
                            var v56 uint32 = 2147483648
                            _ = v56
                        }
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            var v57 float32 = (3.14159 * 0.1)
                            _ = v57
                        case "\n\t":
                            v1 += 127
                        default:
                            if (0 == -361) {
                                var v58 uint32 = 65535
                                fmt.Println(v2)
                                _ = v58
                            }
                        }
                        if v2 {
                            if (v2 || true) {
                                if v2 {
                                    switch 32767 {
                                    case -9223372036854775808:
                                        var v59 float64 = 3.14159
                                        _ = v59
                                    }
                                    if true {
                                        v1 += 42
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "hello":
                                        switch 1 {
                                        case -2147483648:
                                            v1 -= v1
                                        case 2147483647:
                                            switch "hello" {
                                            case "a":
                                                if ((v2 && (-73 >= 1)) && (!(255 == 285))) {
                                                    if v2 {
                                                        v60 := GenS1{F0: 25}
                                                        v18 = 0
                                                        _ = v60
                                                    }
                                                    fmt.Println(v18, v2, v2)
                                                }
                                            case "":
                                                switch int(v18) {
                                                case 32768:
                                                    v61 := GenS4{F0: "a", F1: -32768}
                                                    _ = v61
                                                default:
                                                    var v62 float32 = 0.0
                                                    _ = v62
                                                }
                                            default:
                                                v63 := GenS1{F0: 127}
                                                _ = v63
                                            }
                                        default:
                                            var v64 uint16 = 2
                                            _ = v64
                                        }
                                    case "":
                                        v1 = (v1 / (((^(v1))) | 1))
                                    }
                                }
                                v3 = append(v3, 32767)
                            } else {
                                if (!false) {
                                    switch -986 {
                                    case 128:
                                        for fi65 := 0; fi65 < 12; fi65++ {
                                            v66 := GenS5{F0: 1, F1: 378}
                                            _ = fi65
                                            _ = v66
                                        }
                                    case 32767:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        fmt.Println(v1, v2, v1)
                                    }
                                    fmt.Println(v18, v18, v1)
                                    _ = len(v3)
                                } else {
                                    _ = len(v3)
                                    for fi67 := 0; fi67 < 7; fi67++ {
                                        if (("" <= "") && (!false)) {
                                            _ = len(v3)
                                        }
                                        _ = fi67
                                    }
                                    for fi68 := 0; fi68 < 7; fi68++ {
                                        var v69 string = ""
                                        v18 = 948
                                        _ = fi68
                                        _ = v69
                                    }
                                }
                            }
                            var v70 string = ""
                            _ = v70
                        } else {
                            v71 := GenS6{F0: 1, F1: -9223372036854775808, F2: 640}
                            fmt.Println(v2, v18, v18)
                            if false {
                                switch "foo bar" {
                                case "a":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "hello":
                                    if v2 {
                                        var v72 uint64 = (v18 * genFn1(0))
                                        _ = v72
                                    }
                                default:
                                    var v73 uint32 = 2147483647
                                    _ = v73
                                }
                                for fi74 := 0; fi74 < 19; fi74++ {
                                    var v75 int32 = -2147483648
                                    v75 = v75
                                    for fi76 := 0; fi76 < 7; fi76++ {
                                        var v77 int64 = 1
                                        var v78 int32 = (v75 ^ (v75 - (v75 ^ v75)))
                                        v79 := GenS4{F0: "foo bar", F1: -32768}
                                        _ = fi76
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                    }
                                    _ = fi74
                                    _ = v75
                                }
                            } else {
                                for fi80 := 0; fi80 < 9; fi80++ {
                                    if len(v3) > 0 {
                                        v3[0] = fi80
                                    }
                                    var v81 string = "a"
                                    _ = fi80
                                    _ = v81
                                }
                            }
                            _ = v71
                        }
                    }
                case "a":
                    v1 = uint8(v18)
                case "foo bar":
                    v82 := GenS1{F0: 1}
                    _ = v82
                }
            }
            for fi83 := 0; fi83 < 3; fi83++ {
                v2 = ((((float32(0.1) - (0.1 * float32(1.5))) / 1.5) * (1e10 * float32(1e10))) > 1e10)
                var v84 float32 = -2.25
                var v85 string = "hello"
                _ = fi83
                _ = v84
                _ = v85
            }
        } else {
            var v86 uint = 65535
            if v2 {
                var v87 string = "\n\t"
                for fi88 := 0; fi88 < 12; fi88++ {
                    for fi89 := 0; fi89 < 19; fi89++ {
                        v1 = 128
                        if (true || (v87 == v87)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v90 int64 = 42
                        _ = fi89
                        _ = v90
                    }
                    for fi91 := 0; fi91 < 14; fi91++ {
                        if v2 {
                            var v92 float64 = 3.14159
                            fmt.Println(v2, v1)
                            _ = v92
                        }
                        switch "\n\t" {
                        case "\n\t":
                            _ = len(v3)
                        case "foo bar":
                            fmt.Println(v87, v86, v1)
                        case "":
                            var v93 uint64 = v18
                            _ = v93
                        }
                        v86 = v86
                        v94 := GenS7{F0: v18}
                        v95 := GenS8{F0: 42, F1: v2, F2: 1}
                        _ = fi91
                        _ = v94
                        _ = v95
                    }
                    _ = fi88
                }
                if v2 {
                    v96 := GenS9{F0: -13, F1: (v1 % ((v1) | 1))}
                    fmt.Println(v18, v86, v2)
                    var v97 uint8 = 137
                    _ = v96
                    _ = v97
                }
                _ = v87
            }
            _ = v86
        }
    }
    switch ("foo bar" + "hello") {
    case "a":
        v18 -= 368
    case "hello":
        v3 = append(v3, 9223372036854775807)
    case "\n\t":
        v98 := GenS10{F0: -1, F1: ("hello" + "a")}
        _ = v98
    default:
        for fi99 := 0; fi99 < 20; fi99++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v100 := GenS9{F0: 562, F1: 112}
            _ = fi99
            _ = v100
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
}
