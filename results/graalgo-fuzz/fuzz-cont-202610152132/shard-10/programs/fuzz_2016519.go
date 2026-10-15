package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 bool; F1 string; F2 int }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint8; F2 uint }
type GenS8 struct { F0 int64; F1 int32; F2 int }
func genFn0(p0 string, p1 int32, p2 string) int64 {
    switch p0 {
    case "hello":
        var v7 int64 = -9223372036854775808
        _ = v7
    case "foo bar":
        v8 := GenS1{F0: 1}
        _ = v8
    case "\n\t":
        v9 := GenS1{F0: 2147483648}
        _ = v9
    default:
        v10 := []int{-718, 42, 32767}
        _ = v10
    }
    fmt.Println(p2, p1, p0)
    p0 = "\n\t"
    v11 := []int{-2147483648, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    v1 := []string{("a" + "foo bar"), ""}
    var v2 string = "a"
    v3 := []int{42}
    v4 := make(map[string]bool)
    v5 := []int{42, 524}
    v6 := make(map[int]uint32)
    fmt.Println(v2)
    if (!((v2 + (("a" + v2) + (v2 + v2))) >= ("hello" + v2))) {
        var v12 uint32 = 4294967295
        v12 = v12
        _ = v12
    } else {
        v13 := v6[32768]
        fmt.Println(v2, v13)
        var v14 uint = 65535
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        if (true && false) {
            v16 := GenS2{F0: -2147483648, F1: 255}
            for fi17 := 0; fi17 < 19; fi17++ {
                v2 = v2
                if len(v5) > 0 {
                    v5[0] = (fi17 + 127)
                }
                v2 = v2
                _ = fi17
            }
            if (!((-2147483648 > -763) && (!(748 > 0)))) {
                if (1e10 <= 1e10) {
                    var v18 uint64 = 1
                    var v19 float64 = 0.0
                    if true {
                        if (!false) {
                            v20 := v6[-672]
                            for fi21 := 0; fi21 < 2; fi21++ {
                                fmt.Println(fi21)
                                _ = fi21
                            }
                            if (-32768 == -1) {
                                if len(v5) > 0 {
                                    v5[0] = fi15
                                }
                                v18 += v18
                            }
                            _ = v20
                        }
                        switch (v2 + (((((v2 + ("hello" + "a")) + v2) + (v2 + (v2 + ("foo bar" + ((v2 + "foo bar") + ((v2 + ("a" + v2)) + (((v2 + "") + "a") + "foo bar"))))))) + v2) + ("" + v2))) {
                        case "hello":
                            v2 = ((v2 + (v2 + v2)) + v2)
                        case "\n\t":
                            _ = len(v4)
                        default:
                            delete(v4, (v2 + "\n\t"))
                        }
                    } else {
                        if (0 == 42) {
                            v6[(fi15 & (-(fi15)))] = 42
                            v18 = v18
                        } else {
                            if false {
                                switch v2 {
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = (fi15 & 127)
                                    }
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                            _ = len(v1)
                            var v22 string = v2
                            _ = v22
                        }
                    }
                    _ = v18
                    _ = v19
                }
            } else {
                fmt.Println(fi15, v2)
                var v23 uint8 = 110
                v24 := GenS3{F0: -9, F1: (("hello" + v2) + v2)}
                _ = v23
                _ = v24
            }
            _ = v16
        }
        v2 = ((((("a" + (v2 + ((v2 + "a") + v2))) + v2) + (("\n\t" + "foo bar") + "hello")) + v2) + (v2 + v2))
        v25 := GenS4{F0: true, F1: v2, F2: (fi15 ^ (fi15 + (^(fi15))))}
        _ = fi15
        _ = v25
    }
    var v26 int32 = 2147483647
    switch -67 {
    case 969:
        fmt.Println(v2)
    default:
        v5 = append(v5, 544)
    }
    var v27 int = -2147483648
    v26 -= 472
    switch 0 {
    case 641:
        var v28 bool = (v27 < (v27 << ((4294967295) & 63)))
        _ = v28
    }
    v29 := GenS3{F0: 127, F1: ((v2 + v2) + v2)}
    switch "\n\t" {
    case "hello":
        v30 := v6[v27]
        _ = v30
    case "\n\t":
        for fi31 := 0; fi31 < 10; fi31++ {
            var v32 string = (v2 + "\n\t")
            _ = fi31
            _ = v32
        }
    default:
        if true {
            for fi33 := 0; fi33 < 3; fi33++ {
                v34 := GenS1{F0: 4294967295}
                v2 = "foo bar"
                _ = fi33
                _ = v34
            }
        } else {
            for fi35 := 0; fi35 < 18; fi35++ {
                v27 += 4294967295
                if ((v26 / (((^(v26))) | 1)) == v26) {
                    if (!false) {
                        fmt.Println(v27)
                    } else {
                        switch (v27 % (((fi35 | -9223372036854775808)) | 1)) {
                        case 9223372036854775807:
                            var v36 string = v2
                            _ = v36
                        case 128:
                            if (!true) {
                                if true {
                                    fmt.Println(v2, fi35)
                                    var v37 uint8 = 1
                                    _ = v37
                                } else {
                                    if true {
                                        for fi38 := 0; fi38 < 14; fi38++ {
                                            var v39 int32 = v26
                                            _ = fi38
                                            _ = v39
                                        }
                                        var v40 uint16 = 255
                                        _ = v40
                                    }
                                }
                            }
                        case 42:
                            switch 0 {
                            case 1:
                                if (v27 <= (v27 * fi35)) {
                                    for fi41 := 0; fi41 < 19; fi41++ {
                                        _ = len(v4)
                                        if false {
                                            v26 = v26
                                            if len(v5) > 0 {
                                                v5[0] = 520
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            switch 42 {
                                            case 128:
                                                if true {
                                                    var v42 int32 = -758
                                                    _ = v42
                                                } else {
                                                    v26 = v26
                                                }
                                            case 0:
                                                if (2 <= 2147483647) {
                                                    if true {
                                                        v2 = v2
                                                        v43 := GenS2{F0: v26, F1: 1}
                                                        v26 -= -527
                                                        _ = v43
                                                    } else {
                                                        v27 = v27
                                                        if (884 == 18446744073709551615) {
                                                            var v44 float64 = float64((0.1 + 1.5))
                                                            fmt.Println(fi41, v27, v2)
                                                            fmt.Println(v44, fi35)
                                                            _ = v44
                                                        }
                                                    }
                                                }
                                            }
                                            var v45 string = v2
                                            if true {
                                                fmt.Println(v2, fi35, fi41)
                                            } else {
                                                if false {
                                                    var v46 int8 = -1
                                                    v1 = append(v1, "a")
                                                    var v47 int32 = v26
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    var v48 float32 = 1.5
                                                    fmt.Println(v48, fi35)
                                                    _ = v48
                                                }
                                            }
                                            switch (v27 ^ fi41) {
                                            case 556:
                                                v45 = (v45 + v45)
                                            case 918:
                                                v4[v2] = (1 == 255)
                                                fmt.Println(v45)
                                                v49 := v6[-684]
                                                _ = v49
                                            }
                                            _ = v45
                                        }
                                        switch ("" + v2) {
                                        case "":
                                            _ = len(v1)
                                            switch (-(fi41)) {
                                            case 32767:
                                                v2 = (("hello" + (v2 + "")) + v2)
                                            case -150:
                                                fmt.Println(v2)
                                            }
                                            fmt.Println(v2, v27, v27)
                                        case "foo bar":
                                            v50 := GenS5{F0: 1, F1: -559}
                                            _ = v50
                                        default:
                                            switch 32767 {
                                            case 2147483647:
                                                v51 := GenS6{F0: 4294967295}
                                                _ = v51
                                            default:
                                                if false {
                                                    v52 := GenS4{F0: true, F1: "\n\t", F2: fi41}
                                                    delete(v6, (-(fi35)))
                                                    _ = v52
                                                } else {
                                                    fmt.Println(v27, fi41)
                                                    v53 := v6[842]
                                                    var v54 int32 = (-(v26))
                                                    _ = v53
                                                    _ = v54
                                                }
                                            }
                                        }
                                        _ = fi41
                                    }
                                    v5 = append(v5, 9223372036854775807)
                                } else {
                                    fmt.Println(v26, v26)
                                    switch v27 {
                                    case 42:
                                        var v55 int = (^(fi35))
                                        _ = v55
                                    default:
                                        var v56 uint64 = 4294967295
                                        _ = v56
                                    }
                                }
                            case -9223372036854775808:
                                switch (v27 >> ((18446744073709551615) & 63)) {
                                case 1:
                                    _ = len(v6)
                                case 128:
                                    v57 := v4[""]
                                    _ = v57
                                default:
                                    fmt.Println(v26, v26, v26)
                                }
                            default:
                                if (false || (true && (!false))) {
                                    v26 -= (v26 / ((152) | 1))
                                    var v58 int8 = -5
                                    _ = v58
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        var v59 int64 = -9223372036854775808
                                        _ = v59
                                    case "a":
                                        for fi60 := 0; fi60 < 10; fi60++ {
                                            v61 := GenS2{F0: (^(v26)), F1: 538}
                                            var v62 int16 = 0
                                            _ = fi60
                                            _ = v61
                                            _ = v62
                                        }
                                    case "":
                                        var v63 uint = 18446744073709551615
                                        _ = v63
                                    }
                                }
                            }
                        }
                    }
                } else {
                    if false {
                        fmt.Println(v27)
                    } else {
                        for fi64 := 0; fi64 < 14; fi64++ {
                            v26 -= v26
                            _ = fi64
                        }
                        for fi65 := 0; fi65 < 2; fi65++ {
                            var v66 int16 = 1
                            v27 = v27
                            _ = fi65
                            _ = v66
                        }
                        fmt.Println(fi35, v26)
                    }
                }
                var v67 uint = 1
                _ = fi35
                _ = v67
            }
            for fi68 := 0; fi68 < 10; fi68++ {
                v69 := v4[(v2 + "foo bar")]
                var v70 float64 = -2.25
                _ = fi68
                _ = v69
                _ = v70
            }
        }
    }
    for fi71 := 0; fi71 < 1; fi71++ {
        for fi72 := 0; fi72 < 3; fi72++ {
            v73 := GenS7{F0: 1, F1: 166, F2: 42}
            if (!(828 <= 0)) {
                v1 = append(v1, (v2 + ""))
                v4["foo bar"] = ((755 != 42) || (479 > 18446744073709551615))
                v74 := v4[v2]
                _ = v74
            }
            v75 := GenS8{F0: genFn0("a", -2147483648, "foo bar"), F1: v26, F2: 9223372036854775807}
            _ = fi72
            _ = v73
            _ = v75
        }
        v26 -= (v26 + (v26 % ((v26) | 1)))
        var v76 int64 = genFn0(v2, 0, "")
        _ = fi71
        _ = v76
    }
    fmt.Println(v26, v27)
    v77 := GenS3{F0: 13, F1: v2}
    var v78 float32 = 3.14159
    for fi79 := 0; fi79 < 8; fi79++ {
        var v80 string = (v2 + v2)
        delete(v6, (v27 / (((-(v27))) | 1)))
        _ = fi79
        _ = v80
    }
    delete(v4, (v2 + v2))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
}
