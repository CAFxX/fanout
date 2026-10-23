package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 uint32; F2 int8 }
type GenS10 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS11 struct { F0 string; F1 int64 }
type GenS12 struct { F0 uint8; F1 uint; F2 string }
type GenS13 struct { F0 uint32; F1 uint; F2 uint }
func genFn0(p0 uint16, p1 bool) int8 {
    var v4 int64 = -1
    p1 = (-652 == -736)
    v5 := GenS1{F0: 2147483647, F1: p1, F2: 128}
    v6 := []int{388, -9223372036854775808, -735, -608}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int32, p1 int8, p2 uint64) string {
    var v7 uint16 = 314
    if ((true && true) || (1 > 42)) {
        var v8 int = 2147483647
        v7 -= 974
        switch (v8 ^ (v8 ^ v8)) {
        case 127:
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 int64 = 9223372036854775807
                if false {
                    p0 = (-(p0))
                }
                var v11 uint = 4294967295
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        _ = v8
    }
    var v12 int8 = 42
    v13 := []int{9223372036854775807, 32767, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v7
    return (("hello" + "a") + "")
}

func main() {
    var v1 int16 = 471
    v2 := []string{"a", "hello", "hello", "foo bar"}
    var v3 uint = 255
    switch int(v3) {
    case 418:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v14 int32 = -459
        _ = v14
    }
    fmt.Println(v3, v3)
    v15 := GenS1{F0: 42, F1: false, F2: 2}
    fmt.Println(v3)
    var v16 int32 = 0
    var v17 int8 = 100
    var v18 uint8 = 43
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v19 uint8 = v18
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int32 = -287
        _ = fi20
        _ = v21
    }
    switch "hello" {
    case "hello":
        if false {
            for fi22 := 0; fi22 < 7; fi22++ {
                fmt.Println(v19, v17, v17)
                if (0.1 != (float32((-2.25 - 1.5)) / 1e10)) {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        fmt.Println(v19, fi22, fi22)
                        v3 = 600
                        if true {
                            v24 := GenS2{F0: "foo bar"}
                            v18 = v18
                            v25 := GenS2{F0: ("foo bar" + ("hello" + "foo bar"))}
                            switch (fi23 & fi22) {
                            case -2147483648:
                                v18 -= (v18 * 128)
                            case 127:
                                switch int(v17) {
                                case 160:
                                    switch (fi23 - (fi22 << (((v3 & 18446744073709551615)) & 63))) {
                                    case -36:
                                        v3 = v3
                                    }
                                case -665:
                                    v17 = 100
                                    if ((^(v16)) < (v16 ^ int32(v3))) {
                                        v16 = (v16 >> ((v3) & 63))
                                        v26 := GenS3{F0: 65535}
                                        _ = v26
                                    } else {
                                        if false {
                                            var v27 int16 = 348
                                            v28 := GenS4{F0: (v1 * (^(v27))), F1: 65535}
                                            _ = v27
                                            _ = v28
                                        }
                                        v29 := GenS5{F0: 18446744073709551615, F1: 636, F2: v17}
                                        _ = len(v2)
                                        _ = v29
                                    }
                                    if (((fi23 % ((fi22) | 1)) <= fi22) || (!true)) {
                                        var v30 int16 = -657
                                        fmt.Println(v17)
                                        _ = v30
                                    } else {
                                        switch fi23 {
                                        case -322:
                                            _ = len(v2)
                                        case 42:
                                            v16 = (-(v16))
                                        }
                                    }
                                default:
                                    if true {
                                        v1 += -270
                                    }
                                }
                            }
                            _ = v24
                            _ = v25
                        }
                        _ = fi23
                    }
                    switch (fi22 << ((0) & 63)) {
                    case 2147483647:
                        v3 -= (^(v3))
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    for fi31 := 0; fi31 < 5; fi31++ {
                        if ("hello" <= "a") {
                            v32 := GenS3{F0: 32767}
                            v3 = (v3 & 1)
                            _ = v32
                        } else {
                            var v33 uint = (^(v3))
                            var v34 int = (fi31 ^ fi22)
                            v19 = v19
                            _ = v33
                            _ = v34
                        }
                        v17 += genFn0(2, false)
                        _ = fi31
                    }
                }
                _ = fi22
            }
        }
    case "foo bar":
        var v35 int8 = v17
        _ = v35
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v1 = v1
    var v36 uint = (v3 - v3)
    _ = len(v2)
    if true {
        v3 = 65535
        var v37 float32 = (((1.5 - ((3.14159 / 1.0) + 1.5)) - float32(3.14159)) + float32(1.5))
        _ = v37
    }
    var v38 float64 = (float64(((0.0 - 0.1) - (float32(float64(0.1)) + float32(1e10)))) - 3.14159)
    if false {
        for fi39 := 0; fi39 < 15; fi39++ {
            if false {
                fmt.Println(v19)
                var v40 int = fi39
                v41 := GenS3{F0: 32767}
                _ = v40
                _ = v41
            }
            if false {
                for fi42 := 0; fi42 < 1; fi42++ {
                    v43 := GenS6{F0: ("hello" + "a"), F1: 1}
                    var v44 int8 = 42
                    v18 += v18
                    var v45 string = ("\n\t" + ("hello" + (("foo bar" + "hello") + ("foo bar" + ("\n\t" + "\n\t")))))
                    var v46 uint16 = 649
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v45
                    _ = v46
                }
                v1 -= int16(fi39)
            } else {
                switch (("\n\t" + "foo bar") + "\n\t") {
                case "foo bar":
                    v47 := GenS7{F0: (v16 | 938)}
                    _ = v47
                case "\n\t":
                    var v48 string = "\n\t"
                    _ = v48
                default:
                    _ = len(v2)
                }
                v49 := GenS8{F0: v3}
                _ = v49
            }
            _ = fi39
        }
        v36 += v3
        var v50 int = 128
        _ = v50
    } else {
        if true {
            var v51 string = "foo bar"
            _ = v51
        } else {
            for fi52 := 0; fi52 < 5; fi52++ {
                var v53 int16 = v1
                switch fi52 {
                case 591:
                    if false {
                        if false {
                            if false {
                                for fi54 := 0; fi54 < 17; fi54++ {
                                    v19 += v18
                                    if len(v2) > 0 {
                                        v2[0] = "foo bar"
                                    }
                                    fmt.Println(v3, v17)
                                    _ = fi54
                                }
                                if (!false) {
                                    v53 -= (^(v1))
                                    var v55 float64 = v38
                                    var v56 uint32 = 4294967295
                                    _ = v55
                                    _ = v56
                                } else {
                                    var v57 int16 = v1
                                    var v58 float32 = 1e10
                                    _ = v57
                                    _ = v58
                                }
                                if true {
                                    switch "a" {
                                    case "a":
                                        var v59 int16 = (v1 + (v53 >> ((v3) & 63)))
                                        _ = v59
                                    }
                                    if true {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        var v60 uint = v36
                                        if true {
                                            for fi61 := 0; fi61 < 11; fi61++ {
                                                v62 := GenS3{F0: 2}
                                                _ = fi61
                                                _ = v62
                                            }
                                            if false {
                                                var v63 uint16 = 1
                                                for fi64 := 0; fi64 < 1; fi64++ {
                                                    var v65 int32 = v16
                                                    var v66 float64 = v38
                                                    if false {
                                                        _ = len(v2)
                                                        if true {
                                                            if false {
                                                                v67 := GenS4{F0: (v53 * int16(v16)), F1: 441}
                                                                _ = v67
                                                            } else {
                                                                v68 := GenS5{F0: 4294967295, F1: 216, F2: v17}
                                                                var v69 uint64 = 526
                                                                v1 = -1
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            v66 += v38
                                                        }
                                                    } else {
                                                        v38 += float64(1e10)
                                                    }
                                                    _ = fi64
                                                    _ = v65
                                                    _ = v66
                                                }
                                                v70 := GenS9{F0: 0, F1: 253, F2: v17}
                                                _ = v63
                                                _ = v70
                                            }
                                        }
                                        _ = v60
                                    }
                                }
                            }
                            for fi71 := 0; fi71 < 10; fi71++ {
                                fmt.Println(v17, v3)
                                if (!(true || (true && ((v38 * ((v38 * ((v38 / v38) - (((v38 * (v38 - (v38 * v38))) + -2.25) / v38))) * (float64(0.1) - 1e10))) >= float64((-2.25 / 1e10)))))) {
                                    var v72 int = int(v17)
                                    v38 = v38
                                    _ = v72
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                    switch (fi71 << ((v36) & 63)) {
                                    case 4294967295:
                                        switch "foo bar" {
                                        case "foo bar":
                                            v2 = append(v2, "")
                                        default:
                                            var v73 uint16 = 0
                                            _ = v73
                                        }
                                    case -2147483648:
                                        fmt.Println(v36, fi52, fi52)
                                    case 9223372036854775807:
                                        if ((false || (42 > 938)) || false) {
                                            if true {
                                                if len(v2) > 0 {
                                                    v2[0] = "hello"
                                                }
                                                switch (fi71 ^ fi52) {
                                                case -576:
                                                    var v74 int64 = 992
                                                    _ = v74
                                                }
                                                v18 -= v18
                                                v75 := GenS10{F0: 2, F1: 0, F2: 2}
                                                var v76 int16 = (^(v1))
                                                _ = v75
                                                _ = v76
                                            } else {
                                                switch ("foo bar" + "a") {
                                                case "foo bar":
                                                    _ = len(v2)
                                                case "\n\t":
                                                    v2 = append(v2, "")
                                                case "a":
                                                    v77 := GenS11{F0: "hello", F1: -958}
                                                    _ = v77
                                                }
                                            }
                                            var v78 uint64 = 573
                                            var v79 float32 = float32((v38 / (v38 - (1e10 * 0.0))))
                                            _ = v78
                                            _ = v79
                                        }
                                    }
                                }
                                if len(v2) > 0 {
                                    v2[0] = ("" + "hello")
                                }
                                _ = fi71
                            }
                            var v80 int32 = 2147483647
                            _ = v80
                        }
                        v17 = (v17 >> ((v36) & 63))
                    } else {
                        v2 = append(v2, "")
                        switch "a" {
                        case "a":
                            for fi81 := 0; fi81 < 13; fi81++ {
                                _ = len(v2)
                                _ = fi81
                            }
                        case "hello":
                            fmt.Println(v1, v18)
                        case "foo bar":
                            var v82 float32 = ((-2.25 + float32(((float64(float32((float64(1.5) + ((((v38 / v38) * v38) / v38) + 1e10)))) * v38) - v38))) / 3.14159)
                            _ = v82
                        }
                        fmt.Println(v16, fi52)
                    }
                default:
                    switch ("a" + (((("hello" + "foo bar") + "\n\t") + "\n\t") + ("hello" + "\n\t"))) {
                    case "a":
                        v38 = 1.5
                    default:
                        v83 := GenS12{F0: (v19 << ((v36) & 63)), F1: (v36 & v3), F2: "hello"}
                        _ = v83
                    }
                }
                _ = fi52
                _ = v53
            }
            v84 := GenS13{F0: 42, F1: (^(v3)), F2: (v3 / ((v36) | 1))}
            _ = v84
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v38
}
