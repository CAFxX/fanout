package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0(p0 int32, p1 string, p2 uint16) uint16 {
    v6 := []int{}
    for fi7 := 0; fi7 < 14; fi7++ {
        switch (p1 + "foo bar") {
        case "a":
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 int8 = 127
                _ = fi8
                _ = v9
            }
        case "\n\t":
            v10 := GenS1{F0: 532, F1: (^(p0)), F2: 127}
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 11; fi11++ {
                v12 := GenS2{F0: 0, F1: 68, F2: 2}
                if true {
                    v13 := GenS2{F0: 69, F1: 0, F2: 255}
                    _ = v13
                }
                if len(v6) > 0 {
                    v6[0] = (fi11 >> ((2) & 63))
                }
                _ = fi11
                _ = v12
            }
        default:
            p0 += p0
        }
        p0 = -1
        _ = fi7
    }
    v14 := GenS1{F0: 42, F1: -2147483648, F2: 240}
    for fi15 := 0; fi15 < 17; fi15++ {
        switch (p1 + "a") {
        case "foo bar":
            if (!false) {
                v16 := GenS1{F0: 0, F1: -1, F2: 70}
                switch ((p1 + ((p1 + p1) + "a")) + p1) {
                case "hello":
                    p0 += (p0 >> ((18446744073709551615) & 63))
                case "\n\t":
                    if (128 < (fi15 * (fi15 % ((fi15) | 1)))) {
                        var v17 int16 = 1000
                        _ = v17
                    } else {
                        var v18 uint64 = 134
                        _ = v18
                    }
                case "a":
                    switch (p1 + "hello") {
                    case "a":
                        if ((((1e10 * 3.14159) * 3.14159) / 1.0) == 0.1) {
                            if (239 >= 42) {
                                fmt.Println(p1)
                            } else {
                                switch 361 {
                                case 690:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                case 2147483647:
                                    for fi19 := 0; fi19 < 20; fi19++ {
                                        var v20 uint = 141
                                        _ = fi19
                                        _ = v20
                                    }
                                default:
                                    if (61 > 65535) {
                                        if (!false) {
                                            for fi21 := 0; fi21 < 3; fi21++ {
                                                var v22 uint64 = 736
                                                v23 := GenS3{F0: (p2 / ((p2) | 1))}
                                                _ = fi21
                                                _ = v22
                                                _ = v23
                                            }
                                            var v24 int = fi15
                                            _ = v24
                                        } else {
                                            v25 := GenS4{F0: 0}
                                            if (!true) {
                                                var v26 float64 = 0.0
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                _ = v26
                                            }
                                            fmt.Println(p0)
                                            _ = v25
                                        }
                                        fmt.Println(p1, p2, p2)
                                        if false {
                                            switch fi15 {
                                            case 9223372036854775807:
                                                var v27 int32 = (p0 & 578)
                                                _ = v27
                                            default:
                                                switch (fi15 << ((446) & 63)) {
                                                case 222:
                                                    p0 = -1
                                                case -67:
                                                    p0 += p0
                                                case -1:
                                                    fmt.Println(p1)
                                                default:
                                                    if false {
                                                        var v28 int8 = -74
                                                        p0 = (p0 % (((p0 & (p0 / (((p0 & (p0 % (((^(p0))) | 1)))) | 1)))) | 1))
                                                        _ = v28
                                                    } else {
                                                        v29 := GenS5{F0: -1, F1: 65535}
                                                        _ = v29
                                                    }
                                                }
                                            }
                                            v30 := GenS2{F0: 42, F1: 32767, F2: 2}
                                            _ = v30
                                        } else {
                                            switch fi15 {
                                            case 2147483647:
                                                var v31 bool = (!(2 >= 1))
                                                _ = v31
                                            case -9223372036854775808:
                                                v6 = append(v6, -9223372036854775808)
                                            default:
                                                var v32 int = (fi15 - (fi15 | fi15))
                                                _ = v32
                                            }
                                            p1 = "foo bar"
                                            p1 = (p1 + "foo bar")
                                        }
                                    } else {
                                        p1 = p1
                                        for fi33 := 0; fi33 < 7; fi33++ {
                                            if (1e10 >= (((0.1 - 1e10) + 0.0) - (0.1 + (-2.25 - 1.5)))) {
                                                var v34 uint32 = 844
                                                v35 := GenS3{F0: (^(p2))}
                                                if ((-(fi33)) > int(p0)) {
                                                    p2 = p2
                                                    v36 := GenS5{F0: 42, F1: 394}
                                                    _ = v36
                                                } else {
                                                    if (393 <= 1) {
                                                        fmt.Println(p1, p2)
                                                        v37 := GenS2{F0: 42, F1: -605, F2: 765}
                                                        fmt.Println(p2)
                                                        p2 += (p2 % ((1) | 1))
                                                        _ = v37
                                                    }
                                                }
                                                _ = v34
                                                _ = v35
                                            } else {
                                                p0 -= p0
                                                p1 = (("\n\t" + p1) + (p1 + "\n\t"))
                                            }
                                            if (1e10 != 1e10) {
                                                fmt.Println(p0, p2, p1)
                                            } else {
                                                var v38 uint = uint(p2)
                                                p1 = p1
                                                _ = len(v6)
                                                _ = v38
                                            }
                                            p1 = p1
                                            var v39 uint32 = 672
                                            v40 := GenS2{F0: 0, F1: 0, F2: 1}
                                            _ = fi33
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                }
                            }
                        } else {
                            var v41 uint16 = p2
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v41
                        }
                    }
                default:
                    var v42 uint = 18446744073709551615
                    _ = v42
                }
                _ = v16
            } else {
                var v43 float64 = -2.25
                _ = v43
            }
        case "\n\t":
            _ = len(v6)
        case "a":
            p2 -= (p2 * (p2 & (^(p2))))
        }
        _ = fi15
    }
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    return 1
}

func main() {
    var v1 uint = 646
    v2 := []string{("foo bar" + "a"), "", ("" + "hello")}
    v3 := make(map[int]bool)
    var v4 uint16 = 2
    var v5 float64 = 0.0
    delete(v3, 572)
    v1 = v1
    for fi44 := 0; fi44 < 17; fi44++ {
        switch "" {
        case "\n\t":
            if false {
                for fi45 := 0; fi45 < 18; fi45++ {
                    v46 := GenS5{F0: 57, F1: 2}
                    var v47 int64 = 0
                    v48 := GenS5{F0: 24, F1: 4294967295}
                    _ = fi45
                    _ = v46
                    _ = v47
                    _ = v48
                }
                v49 := GenS5{F0: 42, F1: (v1 + 42)}
                v50 := GenS2{F0: 137, F1: 32767, F2: 255}
                _ = v49
                _ = v50
            }
        }
        if true {
            v51 := v3[(fi44 >> ((391) & 63))]
            _ = v51
        } else {
            for fi52 := 0; fi52 < 14; fi52++ {
                v1 = 768
                var v53 string = ((((((("\n\t" + "\n\t") + ("\n\t" + "\n\t")) + "a") + "foo bar") + "\n\t") + "a") + "hello")
                switch int(v4) {
                case 127:
                    v54 := v3[(fi52 | (fi52 - fi44))]
                    _ = v54
                case 42:
                    v5 = (v5 * v5)
                case 395:
                    v55 := GenS6{F0: (fi44 & (fi44 / ((fi44) | 1)))}
                    var v56 string = "foo bar"
                    _ = v55
                    _ = v56
                default:
                    var v57 int32 = 1
                    _ = v57
                }
                _ = fi52
                _ = v53
            }
            v58 := GenS7{F0: genFn0(42, "foo bar", genFn0(-1, "a", v4)), F1: 2147483647}
            if (v5 != v5) {
                for fi59 := 0; fi59 < 12; fi59++ {
                    if true {
                        var v60 int16 = -510
                        switch "hello" {
                        case "hello":
                            v61 := GenS8{F0: 0, F1: -2147483648}
                            _ = v61
                        }
                        var v62 uint = v1
                        _ = v60
                        _ = v62
                    }
                    var v63 int32 = 995
                    fmt.Println(fi44, v5)
                    _ = fi59
                    _ = v63
                }
                v64 := GenS9{F0: 2, F1: 160, F2: 353}
                switch fi44 {
                case -2147483648:
                    var v65 int = (fi44 << (((^(v1))) & 63))
                    _ = v65
                case -839:
                    fmt.Println(fi44)
                case 9223372036854775807:
                    var v66 uint32 = 992
                    _ = v66
                }
                _ = v64
            } else {
                v1 = v1
            }
            _ = v58
        }
        _ = fi44
    }
    switch "a" {
    case "":
        var v67 uint16 = (v4 * (v4 & v4))
        _ = v67
    }
    var v68 int8 = 0
    var v69 float64 = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v70 int = -1
    v69 += 1.5
    fmt.Println(v4, v70, v1)
    v68 += v68
    var v71 float32 = 0.0
    v5 = v5
    v71 = v71
    switch (("\n\t" + (("" + "") + ("" + "a"))) + "") {
    case "\n\t":
        v68 = (v68 / ((1) | 1))
    default:
        var v72 int32 = -1
        _ = v72
    }
    var v73 int16 = -32768
    var v74 uint64 = 4294967295
    v75 := GenS5{F0: (-(v68)), F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v73
    _ = v74
    _ = v75
}
