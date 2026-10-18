package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 uint }
type GenS4 struct { F0 uint32; F1 uint64; F2 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 string; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint; F1 uint8; F2 int16 }
type GenS12 struct { F0 uint; F1 int8 }
type GenS13 struct { F0 string }
func genFn0(p0 int32, p1 bool) uint8 {
    v6 := []int{-2147483648, -9223372036854775808, 32767}
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 int64) int {
    var v7 int64 = (p0 & (-(p0)))
    switch 9223372036854775807 {
    case -768:
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{(fi8 * 9223372036854775807), -9223372036854775808, 127}
            _ = fi8
            _ = v9
        }
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            if (true && (9223372036854775807 < int64(fi10))) {
                v11 := GenS1{F0: "a", F1: 9223372036854775807, F2: 2}
                var v12 int16 = 32767
                _ = v11
                _ = v12
            } else {
                if (false || true) {
                    v7 = (v7 * (p0 << ((4294967295) & 63)))
                } else {
                    var v13 int = fi10
                    _ = v13
                }
            }
            v7 = v7
            _ = fi10
        }
    }
    var v14 uint = 1
    switch -478 {
    case 199:
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                fmt.Println(fi16)
                v14 = 697
                if true {
                    v14 = v14
                    var v17 uint32 = 400
                    _ = v17
                }
                switch "" {
                case "a":
                    var v18 int16 = 1
                    _ = v18
                default:
                    var v19 int8 = 127
                    _ = v19
                }
                var v20 float32 = 0.1
                _ = fi16
                _ = v20
            }
            if false {
                var v21 int32 = 141
                _ = v21
            }
            _ = fi15
        }
    }
    v22 := GenS1{F0: ("hello" + ""), F1: 0, F2: 205}
    _ = p0
    _ = v14
    _ = v22
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    v3 := []uint32{}
    var v4 uint64 = 616
    v5 := []uint32{88, 65535}
    v23 := GenS1{F0: ("\n\t" + "hello"), F1: 945, F2: 128}
    v4 = v4
    v24 := GenS2{F0: 32767, F1: -42}
    switch 134 {
    case -17:
        for fi25 := 0; fi25 < 19; fi25++ {
            if (false || true) {
                _ = len(v1)
                for fi26 := 0; fi26 < 12; fi26++ {
                    if false {
                        v4 += 0
                        var v27 int32 = 2147483647
                        _ = v27
                    } else {
                        v28 := v1[2]
                        _ = v28
                    }
                    fmt.Println(fi26, fi26)
                    v29 := GenS3{F0: (^(fi26)), F1: 2147483647, F2: 42}
                    _ = fi26
                    _ = v29
                }
            } else {
                var v30 int32 = 1
                switch (fi25 & fi25) {
                case 4294967295:
                    if (!(65535 == 32767)) {
                        switch (-(fi25)) {
                        case -349:
                            if (!false) {
                                for fi31 := 0; fi31 < 6; fi31++ {
                                    if ((!(!(!true))) && (!(274 > 2))) {
                                        var v32 int64 = 42
                                        var v33 uint16 = 209
                                        if true {
                                            v34 := GenS3{F0: (fi31 / (((fi31 >> ((65535) & 63))) | 1)), F1: 42, F2: 4294967295}
                                            var v35 int16 = 42
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = v32
                                        _ = v33
                                    } else {
                                        var v36 int16 = 42
                                        _ = v36
                                    }
                                    _ = fi31
                                }
                                for fi37 := 0; fi37 < 9; fi37++ {
                                    v38 := GenS1{F0: (("\n\t" + ("hello" + ("hello" + "a"))) + "a"), F1: -784, F2: 64}
                                    v1[405] = 2147483648
                                    v39 := GenS4{F0: 723, F1: v4, F2: ("\n\t" + "hello")}
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                                fmt.Println(fi25, fi25)
                            }
                        case 32767:
                            var v40 int = -463
                            _ = v40
                        case -60:
                            v30 = 0
                        default:
                            v41 := GenS3{F0: genFn1(0), F1: 2147483647, F2: 2}
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 20; fi42++ {
                            v43 := GenS1{F0: (("a" + "\n\t") + "\n\t"), F1: 42, F2: 183}
                            _ = fi42
                            _ = v43
                        }
                    }
                }
                _ = v30
            }
            v44 := GenS5{F0: true}
            _ = fi25
            _ = v44
        }
    case -9223372036854775808:
        if ((3.14159 + 0.0) != 0.0) {
            v4 += v4
            v4 += (^(v4))
        } else {
            _ = len(v1)
            if ((0 >= -128) || (((690 != 1) || true) && true)) {
                switch int(v4) {
                case 128:
                    fmt.Println(v4)
                case -696:
                    for fi45 := 0; fi45 < 18; fi45++ {
                        var v46 int16 = -32768
                        _ = fi45
                        _ = v46
                    }
                case -489:
                    var v47 bool = ((!("" != ("hello" + ("" + (("a" + "hello") + ("foo bar" + "")))))) && true)
                    _ = v47
                default:
                    for fi48 := 0; fi48 < 11; fi48++ {
                        for fi49 := 0; fi49 < 11; fi49++ {
                            var v50 float64 = 1.5
                            var v51 bool = true
                            _ = len(v1)
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 12; fi52++ {
                            fmt.Println(v4, fi52)
                            _ = fi52
                        }
                        v53 := GenS6{F0: 0, F1: "hello", F2: v4}
                        _ = fi48
                        _ = v53
                    }
                }
            } else {
                v54 := GenS6{F0: 495, F1: "a", F2: (v4 | (v4 << ((4294967295) & 63)))}
                _ = v54
            }
            var v55 float32 = 3.14159
            _ = v55
        }
    }
    if (1000 < 1) {
        _ = len(v3)
    }
    fmt.Println(v4)
    v4 += (v4 / ((4294967295) | 1))
    var v56 int16 = 0
    for fi57 := 0; fi57 < 20; fi57++ {
        var v58 uint16 = 42
        v59 := GenS5{F0: ((448 == 1) || (false || (true || (1 == 0))))}
        _ = fi57
        _ = v58
        _ = v59
    }
    if (0 != 65535) {
        fmt.Println(v56)
        v60 := GenS7{F0: 127, F1: 42}
        _ = v60
    }
    var v61 uint8 = 2
    if ((1.5 - 0.1) <= float32((3.14159 * (3.14159 - 1e10)))) {
        v62 := GenS8{F0: -2147483648, F1: 1, F2: (v4 | (v4 % (((v4 * v4)) | 1)))}
        v61 = 212
        switch "foo bar" {
        case "hello":
            v63 := GenS6{F0: v56, F1: "foo bar", F2: v4}
            _ = v63
        case "":
            v4 = v4
        case "foo bar":
            fmt.Println(v56, v56, v4)
        default:
            if false {
                if false {
                    fmt.Println(v4, v56)
                    var v64 float32 = float32(3.14159)
                    _ = v64
                }
            } else {
                if false {
                    for fi65 := 0; fi65 < 10; fi65++ {
                        var v66 uint8 = v61
                        var v67 int8 = -1
                        v68 := GenS1{F0: ("a" + ((("\n\t" + "\n\t") + "a") + ("hello" + ""))), F1: 15, F2: v66}
                        _ = fi65
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                } else {
                    v69 := GenS9{F0: 840}
                    _ = v69
                }
                v61 = uint8(v4)
            }
        }
        _ = v62
    } else {
        switch 1 {
        case 32768:
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        }
        _ = len(v1)
    }
    v70 := GenS10{F0: "a"}
    if false {
        v4 = 967
    } else {
        switch 9223372036854775807 {
        case -427:
            var v71 int8 = -9
            _ = v71
        case 1:
            v72 := GenS11{F0: 2, F1: v61, F2: (v56 * (v56 % ((752) | 1)))}
            _ = v72
        case 656:
            for fi73 := 0; fi73 < 20; fi73++ {
                var v74 int32 = 2147483647
                switch (fi73 + 2147483647) {
                case 2147483647:
                    if (false || true) {
                        var v75 bool = (!(((0.1 + 3.14159) / (float32((float64(1e10) * float64(((0.1 / 1.0) * 0.1)))) / (3.14159 / (1.5 * (3.14159 - (1e10 - float32((3.14159 + (0.0 * -2.25))))))))) != float32(3.14159)))
                        for fi76 := 0; fi76 < 7; fi76++ {
                            var v77 int8 = 0
                            if true {
                                switch (fi73 - genFn1(9223372036854775807)) {
                                case 575:
                                    if v75 {
                                        switch fi73 {
                                        case 32768:
                                            if false {
                                                _ = len(v3)
                                            } else {
                                                switch int(v61) {
                                                case 32768:
                                                    v74 = v74
                                                }
                                                _ = len(v1)
                                                switch (-(fi73)) {
                                                case 0:
                                                    if ((-(v74)) >= (^(v74))) {
                                                        v4 = 654
                                                        var v78 int16 = v56
                                                        _ = v78
                                                    }
                                                default:
                                                    if v75 {
                                                        v1[65535] = 2
                                                    }
                                                }
                                            }
                                        case -1:
                                            fmt.Println(v75)
                                        }
                                    }
                                    fmt.Println(v56, v75, v74)
                                case 18:
                                    var v79 uint16 = 408
                                    v61 = v61
                                    fmt.Println(v61, v61, fi76)
                                    _ = v79
                                }
                            }
                            fmt.Println(v61, v56, v61)
                            _ = fi76
                            _ = v77
                        }
                        _ = v75
                    } else {
                        fmt.Println(v56, v74)
                        v74 -= (v74 + v74)
                        v80 := GenS12{F0: uint(v61), F1: -66}
                        _ = v80
                    }
                case 32767:
                    var v81 int8 = -1
                    _ = v81
                case -9223372036854775808:
                    fmt.Println(fi73, v61, fi73)
                default:
                    switch (("" + "foo bar") + "a") {
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        for fi82 := 0; fi82 < 8; fi82++ {
                            var v83 string = ("foo bar" + ("" + ("a" + ("" + ("hello" + ("\n\t" + "a"))))))
                            if false {
                                v61 += genFn0(-798, (!true))
                            }
                            _ = fi82
                            _ = v83
                        }
                    default:
                        fmt.Println(fi73, fi73)
                    }
                }
                v84 := GenS6{F0: (v56 << ((880) & 63)), F1: "hello", F2: 308}
                _ = fi73
                _ = v74
                _ = v84
            }
        }
        v85 := GenS13{F0: ("a" + "hello")}
        _ = v85
    }
    if (1.5 > 1.5) {
        var v86 uint16 = 32767
        v87 := GenS13{F0: "foo bar"}
        _ = v86
        _ = v87
    }
    v56 += v56
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v61
    _ = v70
}
