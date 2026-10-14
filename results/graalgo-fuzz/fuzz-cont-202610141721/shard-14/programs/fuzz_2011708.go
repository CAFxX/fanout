package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 string; F2 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 uint32; F2 int32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8; F1 int; F2 int16 }
type GenS10 struct { F0 int32 }
func genFn0() int64 {
    var v7 int = -2147483648
    var v8 uint8 = 0
    var v9 int8 = 99
    var v10 int64 = 13
    var v11 uint64 = 942
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return v10
}
func genFn1(p0 int16, p1 uint32, p2 uint8) uint8 {
    for fi12 := 0; fi12 < 12; fi12++ {
        p2 += 255
        _ = fi12
    }
    var v13 int64 = -9223372036854775808
    v14 := GenS1{F0: -1, F1: int16(p2), F2: (p0 ^ (p0 + 838))}
    v13 = int64(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return (p2 & 174)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]int)
    var v4 int32 = -408
    var v5 int64 = 42
    v6 := make(map[string]uint32)
    var v15 float32 = (0.0 * 3.14159)
    v5 += v5
    v16 := GenS2{F0: v2, F1: 2, F2: 255}
    switch "hello" {
    case "\n\t":
        for fi17 := 0; fi17 < 1; fi17++ {
            if (!true) {
                v4 -= (v4 % ((v4) | 1))
            } else {
                if false {
                    switch "hello" {
                    case "":
                        for fi18 := 0; fi18 < 17; fi18++ {
                            var v19 uint8 = 255
                            var v20 int16 = 1
                            v19 = v19
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    case "a":
                        v3[v2] = fi17
                    case "hello":
                        v2 += (v2 >> ((433) & 63))
                    default:
                        fmt.Println(v2, v5)
                    }
                    v2 -= (v2 / ((42) | 1))
                }
            }
            if (131 == 134) {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v22 := GenS1{F0: v4, F1: -605, F2: int16(fi21)}
                    if (1 < 1) {
                        v23 := []int{}
                        v2 += (v2 % (((v2 ^ (v2 << ((18446744073709551615) & 63)))) | 1))
                        _ = v23
                    } else {
                        var v24 uint8 = 0
                        if false {
                            delete(v1, v2)
                        } else {
                            v25 := []int{32768}
                            var v26 uint = 42
                            _ = v25
                            _ = v26
                        }
                        fmt.Println(v24)
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
                v4 = v4
            } else {
                var v27 int32 = 575
                _ = len(v1)
                _ = v27
            }
            for fi28 := 0; fi28 < 19; fi28++ {
                var v29 uint8 = 127
                v30 := GenS2{F0: (^(v2)), F1: v29, F2: (v29 | 127)}
                if true {
                    switch (fi28 - (fi17 & (fi17 ^ fi28))) {
                    case -324:
                        fmt.Println(fi17, fi28, fi28)
                        v3[(v2 & v2)] = fi17
                    case -1:
                        v29 -= (v29 >> ((1) & 63))
                    case 841:
                        var v31 string = ""
                        _ = v31
                    }
                }
                fmt.Println(v5, v4, fi17)
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = fi17
        }
    }
    v15 += v15
    if ("hello" != "foo bar") {
        var v32 uint64 = 0
        switch (("foo bar" + "a") + "") {
        case "":
            for fi33 := 0; fi33 < 2; fi33++ {
                v5 = (v5 + v5)
                v34 := GenS3{F0: v32, F1: 255, F2: 32767}
                switch "\n\t" {
                case "hello":
                    fmt.Println(v32, v15)
                case "a":
                    v35 := []int{}
                    _ = v35
                case "foo bar":
                    v36 := GenS2{F0: 894, F1: 42, F2: 42}
                    _ = v36
                }
                _ = fi33
                _ = v34
            }
        default:
            var v37 uint16 = 65535
            _ = v37
        }
        var v38 int32 = v4
        _ = v32
        _ = v38
    }
    v39 := GenS4{F0: 1}
    v4 -= -2147483648
    v2 = (^(v2))
    fmt.Println(v15)
    switch ("\n\t" + "hello") {
    case "\n\t":
        switch "" {
        case "a":
            var v40 int64 = genFn0()
            _ = v40
        case "foo bar":
            v5 = v5
        default:
            if false {
                var v41 uint8 = 2
                _ = v41
            }
        }
    case "foo bar":
        switch 128 {
        case -2147483648:
            if (!true) {
                v5 = (-(v5))
                var v42 int = 96
                _ = v42
            } else {
                v43 := GenS5{F0: 1000, F1: "", F2: 0}
                _ = v43
            }
        case -495:
            if false {
                if false {
                    v5 = v5
                    v4 = 0
                    v44 := GenS6{F0: (("a" + ("hello" + "a")) < (("\n\t" + "hello") + ""))}
                    _ = v44
                }
                v4 += v4
            } else {
                v2 = 2147483647
            }
        }
    default:
        switch ("hello" + "\n\t") {
        case "hello":
            var v45 float64 = (float64((0.1 / 0.1)) - float64(1.5))
            _ = v45
        case "":
            var v46 float32 = (v15 + (3.14159 + v15))
            _ = v46
        default:
            var v47 float32 = (float32((1.5 * 0.1)) - ((v15 / v15) / 0.1))
            _ = v47
        }
    }
    v4 = (v4 - 2147483647)
    for fi48 := 0; fi48 < 17; fi48++ {
        v49 := GenS3{F0: 2, F1: 128, F2: 659}
        _ = fi48
        _ = v49
    }
    v5 -= 42
    if true {
        if false {
            if (("a" + "") == "\n\t") {
                switch 42 {
                case 1000:
                    var v50 float32 = 0.0
                    _ = v50
                case -2147483648:
                    _ = len(v6)
                default:
                    if true {
                        for fi51 := 0; fi51 < 11; fi51++ {
                            for fi52 := 0; fi52 < 2; fi52++ {
                                v53 := []int{}
                                switch (fi52 + (fi52 - 1)) {
                                case -2147483648:
                                    v2 = (v2 ^ v2)
                                }
                                v2 = v2
                                v54 := v3[0]
                                var v55 string = (("" + "a") + (("a" + "foo bar") + "hello"))
                                if (638 <= uint(v2)) {
                                    if false {
                                        var v56 uint8 = genFn1(0, 2147483648, 1)
                                        if false {
                                            v3[(v2 >> ((0) & 63))] = v54
                                        } else {
                                            var v57 int64 = genFn0()
                                            fmt.Println(fi52, v56)
                                            v58 := GenS2{F0: v2, F1: (^(v56)), F2: (v56 / (((v56 + v56)) | 1))}
                                            _ = v57
                                            _ = v58
                                        }
                                        fmt.Println(v55, v4)
                                        _ = v56
                                    }
                                }
                                fmt.Println(fi52, fi51)
                                v54 = fi52
                                v15 = float32(((float64((3.14159 + (v15 * (0.0 / v15)))) + float64(v15)) - ((0.1 / ((0.1 / 1.5) / 1.0)) * 0.1)))
                                v59 := GenS1{F0: -1, F1: -1, F2: int16(v2)}
                                if false {
                                    v54 -= -458
                                } else {
                                    v5 += (v5 ^ genFn0())
                                    var v60 float64 = (float64(v15) * float64((v15 - -2.25)))
                                    _ = v60
                                }
                                if true {
                                    var v61 bool = ((v55 + (v55 + "")) == "hello")
                                    v62 := GenS5{F0: 0, F1: v55, F2: 1}
                                    _ = v61
                                    _ = v62
                                } else {
                                    v15 = v15
                                    var v63 float32 = -2.25
                                    _ = v63
                                }
                                v4 += (v4 & v4)
                                _ = fi52
                                _ = v53
                                _ = v54
                                _ = v55
                                _ = v59
                            }
                            _ = fi51
                        }
                        for fi64 := 0; fi64 < 20; fi64++ {
                            fmt.Println(v2)
                            if (v5 < v5) {
                                for fi65 := 0; fi65 < 7; fi65++ {
                                    v4 += v4
                                    fmt.Println(v4, v2, v4)
                                    switch ("foo bar" + ("\n\t" + ("\n\t" + (("" + "\n\t") + "a")))) {
                                    case "a":
                                        v3[(v2 ^ (v2 ^ 0))] = (fi64 / (((fi65 - fi64)) | 1))
                                    }
                                    if (!false) {
                                        if true {
                                            v5 = genFn0()
                                            v15 += float32((0.1 - ((((((-2.25 / 1.5) - ((3.14159 / -2.25) * (0.1 * 1e10))) / 3.14159) - 0.0) / (-2.25 - ((float64(float32((1e10 / (-2.25 / 1e10)))) + 0.0) * 1.5))) - 0.0)))
                                            fmt.Println(v15, v2)
                                        } else {
                                            v15 = (v15 - v15)
                                            v66 := []int{}
                                            if (!((false || (!(-29 < int8(v2)))) || (!true))) {
                                                v5 -= v5
                                                if false {
                                                    if (-128 < 1) {
                                                        var v67 uint8 = 2
                                                        v15 = float32(1.5)
                                                        var v68 int16 = 884
                                                        _ = v67
                                                        _ = v68
                                                    } else {
                                                        var v69 float64 = (0.1 / 3.14159)
                                                        _ = v69
                                                    }
                                                    v4 += (^(v4))
                                                    v4 = v4
                                                } else {
                                                    var v70 float64 = -2.25
                                                    _ = v70
                                                }
                                                fmt.Println(v5, v4, v2)
                                            }
                                            _ = v66
                                        }
                                        if (-2.25 != (-2.25 + -2.25)) {
                                            switch "" {
                                            case "\n\t":
                                                switch -9223372036854775808 {
                                                case -641:
                                                    v71 := []int{fi65}
                                                    _ = v71
                                                default:
                                                    v72 := GenS7{F0: ("\n\t" + ""), F1: v2, F2: 1}
                                                    _ = v72
                                                }
                                                v4 = (v4 / ((v4) | 1))
                                            }
                                            switch (fi64 % ((fi65) | 1)) {
                                            case 32767:
                                                var v73 uint32 = v2
                                                _ = v73
                                            case 2147483647:
                                                switch (("foo bar" + "") + (((("foo bar" + "a") + (("a" + "\n\t") + "foo bar")) + "\n\t") + ("hello" + ("foo bar" + "")))) {
                                                case "foo bar":
                                                    var v74 int8 = 1
                                                    _ = v74
                                                }
                                            case -1:
                                                v4 += v4
                                            }
                                        } else {
                                            v75 := v6[("a" + ((("hello" + "") + "hello") + (("\n\t" + "") + ((("foo bar" + "\n\t") + "foo bar") + ""))))]
                                            var v76 bool = (!true)
                                            var v77 int8 = -128
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                    switch (fi65 & (^(fi64))) {
                                    case 9223372036854775807:
                                        fmt.Println(v5, v5)
                                    default:
                                        v78 := GenS5{F0: int16(fi64), F1: "\n\t", F2: 42}
                                        _ = v78
                                    }
                                    delete(v1, (v2 * v2))
                                    _ = fi65
                                }
                                var v79 int64 = (v5 + int64(v2))
                                if false {
                                    var v80 int32 = v4
                                    _ = v80
                                }
                                _ = v79
                            } else {
                                delete(v1, (v2 << ((2) & 63)))
                                if ((fi64 << ((0) & 63)) >= (fi64 | (^(fi64)))) {
                                    var v81 int = (fi64 >> ((255) & 63))
                                    _ = v81
                                } else {
                                    for fi82 := 0; fi82 < 1; fi82++ {
                                        v4 = v4
                                        v2 -= v2
                                        v5 = (v5 - int64(v4))
                                        _ = fi82
                                    }
                                    v2 -= 0
                                    for fi83 := 0; fi83 < 10; fi83++ {
                                        if (!(4294967295 >= 246)) {
                                            v1[v2] = "foo bar"
                                        }
                                        switch fi64 {
                                        case -2147483648:
                                            if (4294967295 <= 4294967295) {
                                                var v84 string = "\n\t"
                                                v85 := GenS8{F0: uint8(v2)}
                                                _ = v84
                                                _ = v85
                                            }
                                        case -831:
                                            switch "" {
                                            case "\n\t":
                                                var v86 string = ""
                                                _ = v86
                                            case "":
                                                if (!true) {
                                                    v87 := []int{(fi83 >> ((1) & 63)), (fi64 | fi64), 982}
                                                    v15 = v15
                                                    v2 = (v2 ^ (v2 & (v2 - v2)))
                                                    v15 += (v15 + 0.0)
                                                    _ = v87
                                                }
                                            }
                                            var v88 int64 = 9223372036854775807
                                            var v89 int8 = 42
                                            _ = v88
                                            _ = v89
                                        default:
                                            var v90 int = (-(fi83))
                                            _ = v90
                                        }
                                        _ = fi83
                                    }
                                }
                            }
                            v91 := []int{32768, (fi64 - fi64), fi64}
                            _ = fi64
                            _ = v91
                        }
                        var v92 int16 = 0
                        _ = v92
                    }
                }
                v93 := GenS8{F0: genFn1(-1, (v2 * v2), 0)}
                _ = v93
            } else {
                v5 = (^(v5))
                switch 4294967295 {
                case -2147483648:
                    if true {
                        var v94 uint8 = 127
                        _ = v94
                    } else {
                        v15 -= (1.5 / v15)
                        if false {
                            fmt.Println(v4, v2, v15)
                        }
                        var v95 int8 = -88
                        _ = v95
                    }
                case -369:
                    fmt.Println(v15)
                case 42:
                    if true {
                        v96 := GenS9{F0: 74, F1: int(v5), F2: -1}
                        for fi97 := 0; fi97 < 12; fi97++ {
                            v15 = (((-2.25 / float32(((((0.0 / 0.1) / 1e10) - (1e10 * float64((v15 * (-2.25 - ((v15 / (float32(0.0) - v15)) - float32(1.5))))))) * -2.25))) * 1.5) / v15)
                            for fi98 := 0; fi98 < 15; fi98++ {
                                switch (fi98 / ((fi97) | 1)) {
                                case -9223372036854775808:
                                    v99 := []int{(^(fi97)), fi97, 9223372036854775807}
                                    _ = v99
                                default:
                                    v2 = (v2 / (((v2 ^ (v2 * v2))) | 1))
                                }
                                _ = fi98
                            }
                            var v100 int32 = (v4 - 42)
                            _ = fi97
                            _ = v100
                        }
                        v101 := v3[v2]
                        _ = v101
                        _ = v96
                    }
                default:
                    v4 = v4
                }
                if true {
                    if (18446744073709551615 > 4294967295) {
                        v102 := GenS3{F0: 42, F1: 229, F2: 32767}
                        delete(v1, 2147483648)
                        var v103 bool = false
                        _ = v102
                        _ = v103
                    } else {
                        var v104 uint = 0
                        _ = v104
                    }
                    var v105 uint16 = 1
                    for fi106 := 0; fi106 < 15; fi106++ {
                        for fi107 := 0; fi107 < 15; fi107++ {
                            if (false || true) {
                                var v108 int64 = v5
                                v5 -= (v108 ^ v108)
                                fmt.Println(fi106, v5)
                                _ = v108
                            }
                            _ = fi107
                        }
                        var v109 string = ("a" + "foo bar")
                        _ = fi106
                        _ = v109
                    }
                    _ = v105
                } else {
                    v15 = v15
                    fmt.Println(v5, v4, v2)
                    for fi110 := 0; fi110 < 4; fi110++ {
                        fmt.Println(v5)
                        _ = fi110
                    }
                }
            }
        } else {
            if (287 < -692) {
                if (-846 < 720) {
                    v2 = 4294967295
                    for fi111 := 0; fi111 < 6; fi111++ {
                        v4 += -862
                        v112 := GenS2{F0: (v2 % ((v2) | 1)), F1: 1, F2: 127}
                        _ = fi111
                        _ = v112
                    }
                    var v113 uint64 = 0
                    _ = v113
                }
                v15 = 1e10
            } else {
                var v114 uint64 = 368
                fmt.Println(v2)
                _ = v114
            }
            v15 += 0.1
        }
        v115 := GenS10{F0: v4}
        _ = v115
    } else {
        v2 = v2
    }
    v4 = -717
    fmt.Println(v5, v5, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
