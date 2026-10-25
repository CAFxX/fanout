package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS10 struct { F0 uint16; F1 uint8 }
type GenS11 struct { F0 int64; F1 string }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string; F1 uint8; F2 int32 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int32; F1 int8; F2 int16 }
type GenS16 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS17 struct { F0 bool; F1 string }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 bool; F1 uint32 }
func genFn0() int8 {
    var v6 float64 = 1.5
    var v7 string = "hello"
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 string) string {
    var v8 float64 = -2.25
    switch -714 {
    case 470:
        fmt.Println(v8, v8)
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: 42, F1: true, F2: (p0 + p0)}
            v8 += 1e10
            fmt.Println(v8, v8, fi9)
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = v8
    return p0
}

func main() {
    var v1 uint64 = 0
    var v2 bool = ((true && ((0.1 >= 3.14159) && ((0 >= 65535) || false))) || false)
    v3 := []string{"foo bar", "", ("" + ("" + "a"))}
    v4 := make(map[int]int)
    var v5 int8 = 1
    if ((-1 <= -1) || ("hello" <= "foo bar")) {
        var v11 int8 = genFn0()
        _ = v11
    } else {
        var v12 bool = v2
        var v13 uint32 = 2147483648
        _ = v12
        _ = v13
    }
    v2 = (("hello" + "hello") != "hello")
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v1, fi14, v2)
        var v15 int = (fi14 & fi14)
        if (1 <= -220) {
            fmt.Println(fi14)
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int16 = 0
                if ((v1 < v1) && (!(-1 <= 318))) {
                    fmt.Println(fi16, v1, v2)
                    fmt.Println(v2)
                }
                v18 := GenS2{F0: 707}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 uint32 = 676
            _ = v19
        } else {
            var v20 int32 = 42
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := v4[(fi21 + v15)]
                var v23 int8 = 56
                fmt.Println(v20, v20, v5)
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        }
        _ = fi14
        _ = v15
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(fi24, v2)
        for fi25 := 0; fi25 < 17; fi25++ {
            if true {
                v26 := GenS3{F0: 202, F1: 9223372036854775807}
                var v27 uint16 = 2
                _ = v26
                _ = v27
            } else {
                fmt.Println(fi24)
                v28 := GenS4{F0: 0, F1: 2147483647}
                _ = len(v3)
                v5 = (v5 | v5)
                _ = v28
            }
            var v29 uint16 = 2
            v3 = append(v3, "a")
            _ = fi25
            _ = v29
        }
        _ = fi24
    }
    if v2 {
        v30 := GenS5{F0: 42}
        _ = v30
    } else {
        var v31 string = "a"
        if (!(false || (0 >= 252))) {
            v5 = (v5 - genFn0())
            v32 := GenS6{F0: v31, F1: 18446744073709551615}
            for fi33 := 0; fi33 < 14; fi33++ {
                v34 := GenS4{F0: 255, F1: 42}
                for fi35 := 0; fi35 < 13; fi35++ {
                    v3 = append(v3, "\n\t")
                    switch ("" + (v31 + v31)) {
                    case "\n\t":
                        var v36 uint8 = 127
                        _ = v36
                    }
                    var v37 uint64 = (v1 % ((v1) | 1))
                    fmt.Println(v5, fi35, v31)
                    v5 -= v5
                    fmt.Println(fi33, v31)
                    _ = fi35
                    _ = v37
                }
                switch v31 {
                case "hello":
                    for fi38 := 0; fi38 < 2; fi38++ {
                        var v39 bool = (-460 == -627)
                        _ = fi38
                        _ = v39
                    }
                case "foo bar":
                    switch 42 {
                    case -2147483648:
                        v2 = (!v2)
                    case 128:
                        var v40 uint8 = 127
                        _ = v40
                    case 643:
                        var v41 uint32 = 42
                        _ = v41
                    }
                }
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        for fi42 := 0; fi42 < 11; fi42++ {
            v43 := GenS7{F0: fi42, F1: v5}
            _ = fi42
            _ = v43
        }
        _ = v31
    }
    switch "" {
    case "foo bar":
        switch 0 {
        case 0:
            if (497 == 42) {
                v5 = v5
                v2 = (0 != 42)
            }
        case 739:
            var v44 uint64 = v1
            _ = v44
        default:
            if (1 <= -9223372036854775808) {
                var v45 uint64 = v1
                v1 = v1
                v4[2147483647] = 127
                _ = v45
            } else {
                var v46 int = 4294967295
                v5 = v5
                if (v1 < v1) {
                    v47 := GenS8{F0: 0}
                    for fi48 := 0; fi48 < 5; fi48++ {
                        switch "a" {
                        case "":
                            var v49 int = (v46 - fi48)
                            _ = v49
                        case "\n\t":
                            switch "" {
                            case "a":
                                var v50 uint = 740
                                _ = v50
                            case "\n\t":
                                var v51 uint32 = 4294967295
                                _ = v51
                            case "hello":
                                switch ("foo bar" + "") {
                                case "\n\t":
                                    v5 += (v5 | (-(v5)))
                                default:
                                    v5 -= int8(v46)
                                }
                            default:
                                fmt.Println(v1, v5)
                            }
                        case "a":
                            for fi52 := 0; fi52 < 3; fi52++ {
                                var v53 int64 = -652
                                v46 = (^(fi52))
                                v46 = (^(fi52))
                                _ = fi52
                                _ = v53
                            }
                        }
                        if (!(0 != 8)) {
                            var v54 bool = v2
                            _ = v54
                        }
                        _ = len(v3)
                        _ = fi48
                    }
                    _ = v47
                }
                _ = v46
            }
        }
    default:
        for fi55 := 0; fi55 < 5; fi55++ {
            for fi56 := 0; fi56 < 18; fi56++ {
                switch (fi55 >> ((4294967295) & 63)) {
                case 166:
                    v57 := GenS2{F0: 4294967295}
                    _ = v57
                default:
                    var v58 uint16 = 868
                    _ = v58
                }
                switch ("\n\t" + ("a" + "a")) {
                case "a":
                    switch "\n\t" {
                    case "a":
                        if v2 {
                            v59 := GenS9{F0: v1, F1: 0, F2: v5}
                            switch fi56 {
                            case 4294967295:
                                v60 := GenS6{F0: ("" + ("\n\t" + "foo bar")), F1: (v1 ^ 2)}
                                _ = v60
                            case 715:
                                fmt.Println(v5, v2)
                            default:
                                var v61 bool = v2
                                _ = v61
                            }
                            _ = v59
                        } else {
                            if v2 {
                                v2 = v2
                                var v62 uint64 = (v1 * v1)
                                v4[(fi55 << ((2) & 63))] = (-(fi55))
                                _ = v62
                            }
                        }
                    default:
                        v63 := GenS10{F0: 269, F1: 42}
                        _ = v63
                    }
                case "hello":
                    var v64 int16 = int16(fi55)
                    _ = v64
                }
                if v2 {
                    fmt.Println(fi55)
                } else {
                    var v65 int16 = 0
                    _ = v65
                }
                _ = fi56
            }
            v1 = v1
            v5 = v5
            _ = fi55
        }
    }
    v2 = (0 > -9223372036854775808)
    fmt.Println(v2, v2)
    switch 32768 {
    case 42:
        _ = len(v4)
    case 759:
        _ = len(v3)
    case 0:
        v66 := GenS9{F0: v1, F1: 2, F2: v5}
        _ = v66
    default:
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    }
    fmt.Println(v1)
    for fi67 := 0; fi67 < 14; fi67++ {
        switch "hello" {
        case "\n\t":
            v1 -= v1
        case "foo bar":
            v68 := GenS11{F0: 0, F1: ("foo bar" + "\n\t")}
            _ = v68
        default:
            v1 -= v1
        }
        v2 = ((1e10 == 0.1) || ((2 != 2) && v2))
        for fi69 := 0; fi69 < 3; fi69++ {
            _ = len(v4)
            _ = fi69
        }
        _ = fi67
    }
    if true {
        if v2 {
            if true {
                v70 := GenS12{F0: 128}
                for fi71 := 0; fi71 < 14; fi71++ {
                    if v2 {
                        var v72 int64 = -9223372036854775808
                        fmt.Println(v1, v1, v1)
                        _ = v72
                    } else {
                        var v73 bool = v2
                        _ = v73
                    }
                    _ = fi71
                }
                if false {
                    fmt.Println(v2, v2, v5)
                    v5 -= 127
                } else {
                    fmt.Println(v1, v1)
                }
                _ = v70
            } else {
                if v2 {
                    for fi74 := 0; fi74 < 16; fi74++ {
                        switch fi74 {
                        case 1:
                            switch ((("\n\t" + ("a" + (("\n\t" + "hello") + "hello"))) + "\n\t") + "") {
                            case "foo bar":
                                var v75 float32 = ((0.1 - (0.1 / 1.0)) - 1.5)
                                _ = v75
                            }
                        case 138:
                            v76 := GenS2{F0: 42}
                            _ = v76
                        case -344:
                            v77 := GenS11{F0: -9223372036854775808, F1: ""}
                            _ = v77
                        }
                        _ = fi74
                    }
                    v2 = v2
                } else {
                    if v2 {
                        var v78 int8 = genFn0()
                        v5 += 100
                        v2 = (("hello" > "hello") || (0 != 42))
                        _ = v78
                    }
                }
                switch 4294967295 {
                case 32767:
                    if ((892 != 1) || v2) {
                        if (!false) {
                            v79 := GenS13{F0: ("foo bar" + "foo bar"), F1: 0, F2: -232}
                            var v80 float32 = 0.0
                            _ = v79
                            _ = v80
                        }
                        v81 := GenS14{F0: 146}
                        for fi82 := 0; fi82 < 5; fi82++ {
                            fmt.Println(fi82)
                            switch 42 {
                            case -595:
                                v1 = 0
                            default:
                                v83 := v4[(fi82 ^ (^(fi82)))]
                                _ = v83
                            }
                            if (0 < -9223372036854775808) {
                                v84 := GenS15{F0: 0, F1: (v5 - (v5 & (v5 | (v5 ^ (v5 % (((^(v5))) | 1)))))), F2: 893}
                                _ = v84
                            } else {
                                switch ("" + ("\n\t" + "hello")) {
                                case "":
                                    v85 := GenS11{F0: 0, F1: "\n\t"}
                                    _ = v85
                                }
                                v2 = ((("a" + ("" + "\n\t")) >= ("hello" + "foo bar")) && v2)
                            }
                            _ = fi82
                        }
                        _ = v81
                    } else {
                        if (42 != -104) {
                            for fi86 := 0; fi86 < 4; fi86++ {
                                switch ("\n\t" + "a") {
                                case "hello":
                                    for fi87 := 0; fi87 < 7; fi87++ {
                                        if (2 > 42) {
                                            v88 := GenS4{F0: 4294967295, F1: 42}
                                            fmt.Println(v2, fi86, v2)
                                            _ = v88
                                        } else {
                                            v1 = (^(v1))
                                            v89 := GenS10{F0: 32767, F1: 164}
                                            _ = v89
                                        }
                                        _ = fi87
                                    }
                                case "\n\t":
                                    var v90 float64 = 0.1
                                    _ = v90
                                }
                                v4[-194] = fi86
                                v1 = 42
                                _ = fi86
                            }
                            if (v1 == (v1 * v1)) {
                                v2 = true
                                var v91 uint8 = 0
                                if (1 != -599) {
                                    switch 2147483647 {
                                    case 0:
                                        v1 -= v1
                                    default:
                                        fmt.Println(v91, v2)
                                    }
                                    v92 := GenS1{F0: 1, F1: (((1 < 1) || v2) && v2), F2: ("\n\t" + (("foo bar" + "a") + (("\n\t" + ("hello" + "hello")) + "\n\t")))}
                                    _ = v92
                                } else {
                                    v2 = (!(!v2))
                                }
                                _ = v91
                            }
                        }
                        if (!(("\n\t" + "") != "a")) {
                            for fi93 := 0; fi93 < 20; fi93++ {
                                if (-2147483648 != -1) {
                                    if (2147483647 < 759) {
                                        fmt.Println(v2)
                                    }
                                }
                                var v94 bool = (!v2)
                                _ = fi93
                                _ = v94
                            }
                            v95 := GenS16{F0: 0, F1: v5, F2: 42}
                            _ = v95
                        }
                        delete(v4, 32767)
                    }
                }
            }
            if true {
                fmt.Println(v5, v5, v2)
            } else {
                var v96 int16 = 32767
                _ = v96
            }
            v5 += (v5 % ((v5) | 1))
        } else {
            switch 128 {
            case 127:
                for fi97 := 0; fi97 < 16; fi97++ {
                    v4[(^(fi97))] = -9223372036854775808
                    var v98 int = fi97
                    var v99 float64 = ((0.1 - (0.0 + ((-2.25 * 1e10) / 1.5))) * -2.25)
                    _ = fi97
                    _ = v98
                    _ = v99
                }
            case -943:
                switch "\n\t" {
                case "":
                    v100 := GenS15{F0: 42, F1: v5, F2: -338}
                    _ = v100
                case "hello":
                    var v101 string = (((("foo bar" + "a") + "foo bar") + "\n\t") + "")
                    _ = v101
                }
            case 219:
                var v102 float64 = 1.5
                _ = v102
            default:
                _ = len(v3)
            }
        }
        var v103 uint8 = 72
        var v104 uint8 = v103
        _ = v103
        _ = v104
    } else {
        v2 = (138 > 1)
        for fi105 := 0; fi105 < 19; fi105++ {
            v106 := GenS17{F0: true, F1: ("hello" + "")}
            _ = fi105
            _ = v106
        }
        if v2 {
            switch 2147483647 {
            case 32767:
                v107 := GenS18{F0: 1000}
                _ = v107
            case -9223372036854775808:
                v108 := GenS19{F0: (!(v1 == (v1 & v1))), F1: 299}
                _ = v108
            }
            var v109 int32 = -1
            _ = v109
        }
    }
    var v110 uint = 1
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
