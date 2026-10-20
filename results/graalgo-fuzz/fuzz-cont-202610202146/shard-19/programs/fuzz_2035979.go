package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 string; F2 int }
type GenS3 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint64; F1 string; F2 int16 }
func genFn0() uint {
    if false {
        var v6 int = -670
        var v7 uint32 = 2
        if true {
            v6 += 1
            if true {
                v6 += v6
                var v8 int = v6
                for fi9 := 0; fi9 < 5; fi9++ {
                    if true {
                        fmt.Println(v8, v6)
                    }
                    v6 = 628
                    _ = fi9
                }
                _ = v8
            }
            v10 := GenS1{F0: v6}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                if (((!false) || true) || false) {
                    switch v6 {
                    case 2147483647:
                        fmt.Println(v7, v7, fi11)
                    case 32767:
                        v6 += fi11
                    default:
                        fmt.Println(v7, v7)
                    }
                    v6 = (fi11 - (fi11 & v6))
                    switch 42 {
                    case 32768:
                        switch v6 {
                        case 32767:
                            for fi12 := 0; fi12 < 14; fi12++ {
                                fmt.Println(v6)
                                fmt.Println(v7, v7)
                                _ = fi12
                            }
                        default:
                            fmt.Println(v6)
                        }
                    case -575:
                        v7 = 2147483647
                    case -9223372036854775808:
                        fmt.Println(fi11)
                    }
                }
                if true {
                    v13 := []int{v6, (v6 >> ((0) & 63)), 128}
                    _ = v13
                }
                _ = fi11
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int64 = 42
                fmt.Println(v15, v7)
                if (262 < 32767) {
                    var v16 uint16 = 42
                    fmt.Println(v6, v15)
                    if true {
                        v15 -= (v15 - v15)
                        switch "hello" {
                        case "hello":
                            if false {
                                for fi17 := 0; fi17 < 12; fi17++ {
                                    switch "\n\t" {
                                    case "":
                                        v18 := GenS2{F0: (0 > 42), F1: "hello", F2: (v6 >> ((0) & 63))}
                                        _ = v18
                                    case "hello":
                                        var v19 bool = (true && false)
                                        _ = v19
                                    }
                                    _ = fi17
                                }
                            }
                        }
                    } else {
                        v20 := []int{(^(fi14))}
                        _ = v20
                    }
                    _ = v16
                } else {
                    var v21 int64 = v15
                    v7 = (v7 / (((v7 & v7)) | 1))
                    var v22 int32 = -174
                    _ = v21
                    _ = v22
                }
                _ = fi14
                _ = v15
            }
        }
        _ = v6
        _ = v7
    } else {
        var v23 int32 = 1
        _ = v23
    }
    var v24 float32 = float32(3.14159)
    if ("" != "a") {
        v25 := []int{128}
        _ = v25
    } else {
        var v26 bool = false
        v26 = v26
        switch -2147483648 {
        case -1:
            v27 := GenS3{F0: 625, F1: -570, F2: 1}
            _ = v27
        default:
            for fi28 := 0; fi28 < 17; fi28++ {
                switch (fi28 / ((fi28) | 1)) {
                case 32768:
                    v24 = v24
                case -2147483648:
                    v24 -= v24
                default:
                    var v29 bool = ((268 >= 1) || v26)
                    _ = v29
                }
                v24 += ((v24 / v24) - v24)
                _ = fi28
            }
        }
        _ = v26
    }
    v24 += float32(3.14159)
    v24 = (1e10 / 3.14159)
    _ = v24
    return 2
}
func genFn1(p0 uint16, p1 int32, p2 int) int8 {
    var v30 uint32 = uint32(p0)
    switch "foo bar" {
    case "a":
        for fi31 := 0; fi31 < 7; fi31++ {
            switch "foo bar" {
            case "":
                switch (-(fi31)) {
                case -9223372036854775808:
                    v32 := []int{}
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 2; fi33++ {
                        fmt.Println(p1)
                        fmt.Println(fi33, v30, fi31)
                        if (670 > 18446744073709551615) {
                            p0 -= 32767
                            switch "hello" {
                            case "":
                                p2 -= p2
                            case "foo bar":
                                var v34 int = (-(fi31))
                                _ = v34
                            case "hello":
                                var v35 uint16 = (p0 * 32767)
                                _ = v35
                            default:
                                v36 := []int{p2}
                                _ = v36
                            }
                            var v37 uint16 = (p0 | (p0 / ((p0) | 1)))
                            var v38 uint32 = v30
                            v39 := []int{32767}
                            _ = v37
                            _ = v38
                            _ = v39
                        } else {
                            v40 := []int{(fi31 % (((p2 / ((32768) | 1))) | 1)), (fi31 % ((p2) | 1)), 930, -2147483648}
                            if false {
                                v30 -= (v30 & v30)
                            }
                            _ = v40
                        }
                        _ = fi33
                    }
                }
            case "hello":
                v41 := GenS2{F0: (!(-781 != 2147483647)), F1: "foo bar", F2: p2}
                _ = v41
            }
            if (2147483647 == (p1 ^ (p1 - 42))) {
                fmt.Println(p0)
                v30 -= v30
                fmt.Println(p1)
            } else {
                if (403 != 30) {
                    fmt.Println(p2, p0, p2)
                    var v42 int = 42
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 3; fi43++ {
                        fmt.Println(p1, v30)
                        v44 := []int{(p2 % ((fi43) | 1))}
                        var v45 uint32 = v30
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    var v46 float32 = 0.0
                    _ = v46
                }
            }
            _ = fi31
        }
    case "hello":
        for fi47 := 0; fi47 < 6; fi47++ {
            v48 := GenS4{F0: (v30 / (((v30 / ((v30) | 1))) | 1)), F1: 60}
            switch "" {
            case "foo bar":
                for fi49 := 0; fi49 < 2; fi49++ {
                    p1 -= 2147483647
                    _ = fi49
                }
            }
            _ = fi47
            _ = v48
        }
    default:
        p2 -= int(v30)
    }
    v50 := []int{121, (p2 >> ((748) & 63)), p2, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v50
    return 1
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]string)
    var v3 uint64 = 4294967295
    v4 := []string{("a" + "a")}
    var v5 int16 = 32767
    switch -2147483648 {
    case 128:
        var v51 int16 = (v5 - v5)
        _ = v51
    case 32767:
        v52 := GenS2{F0: ("a" < ("" + "")), F1: ("" + "foo bar"), F2: 9223372036854775807}
        _ = v52
    case -2147483648:
        v5 -= v5
    default:
        if (127 == genFn1(1, 42, 9)) {
            fmt.Println(v1)
            var v53 float64 = (0.1 / 1.5)
            for fi54 := 0; fi54 < 2; fi54++ {
                v5 -= 173
                fmt.Println(v1, v53)
                _ = fi54
            }
            _ = v53
        } else {
            fmt.Println(v1, v3)
            v1 = (^(v1))
        }
    }
    if true {
        v55 := GenS2{F0: true, F1: ("foo bar" + ""), F2: 127}
        _ = v55
    } else {
        v3 -= v3
    }
    _ = len(v2)
    v56 := GenS1{F0: 32767}
    fmt.Println(v3, v1, v1)
    for fi57 := 0; fi57 < 5; fi57++ {
        v1 -= v1
        _ = fi57
    }
    var v58 int32 = -1
    switch -9223372036854775808 {
    case 9223372036854775807:
        if (!true) {
            for fi59 := 0; fi59 < 11; fi59++ {
                switch 822 {
                case 608:
                    v5 = (v5 / (((v5 | v5)) | 1))
                case 545:
                    for fi60 := 0; fi60 < 20; fi60++ {
                        fmt.Println(v58)
                        v5 += v5
                        _ = fi60
                    }
                }
                v1 = 778
                v61 := GenS4{F0: v1, F1: -120}
                _ = fi59
                _ = v61
            }
            var v62 float32 = 0.1
            if false {
                v3 -= (v3 + v3)
            } else {
                v63 := GenS1{F0: 2147483647}
                _ = v63
            }
            _ = v62
        }
    case 0:
        for fi64 := 0; fi64 < 2; fi64++ {
            for fi65 := 0; fi65 < 5; fi65++ {
                v66 := GenS5{F0: 5}
                _ = fi65
                _ = v66
            }
            var v67 string = "\n\t"
            _ = fi64
            _ = v67
        }
    case -9223372036854775808:
        var v68 float64 = -2.25
        _ = v68
    default:
        v69 := GenS6{F0: v1, F1: 0}
        _ = v69
    }
    if (((("foo bar" + ("a" + "hello")) + "a") + "foo bar") != "") {
        var v70 int8 = genFn1(42, -1, 128)
        _ = v70
    }
    if true {
        var v71 int32 = (v58 ^ v58)
        if true {
            var v72 float32 = 1.5
            _ = v72
        } else {
            fmt.Println(v3, v71)
        }
        _ = v71
    } else {
        fmt.Println(v58)
        v73 := GenS7{F0: (^(v3)), F1: "foo bar", F2: 1000}
        switch "a" {
        case "foo bar":
            v58 += v58
        }
        _ = v73
    }
    v58 = (v58 % (((v58 % (((v58 ^ -2147483648)) | 1))) | 1))
    var v74 float32 = (0.0 + 1.5)
    v75 := GenS4{F0: (v1 >> ((778) & 63)), F1: 127}
    v5 += 33
    if true {
        v76 := v2[2147483647]
        _ = v76
    } else {
        var v77 int32 = (v58 ^ v58)
        v3 -= 4294967295
        _ = v77
    }
    for fi78 := 0; fi78 < 5; fi78++ {
        v5 += (v5 / ((1) | 1))
        var v79 float32 = 3.14159
        _ = fi78
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v58
    _ = v74
    _ = v75
}
