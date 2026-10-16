package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int; F2 uint32 }
type GenS8 struct { F0 int16; F1 int64; F2 int64 }
func genFn0(p0 int32, p1 int8, p2 bool) string {
    p1 = 55
    v6 := GenS1{F0: 0}
    var v7 int8 = -5
    v8 := GenS2{F0: p0, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ("\n\t" + "hello")
}
func genFn1(p0 uint16, p1 int16) string {
    v9 := GenS3{F0: 1}
    switch (("\n\t" + "hello") + "a") {
    case "\n\t":
        p1 = p1
    case "a":
        v10 := GenS4{F0: -110, F1: 287}
        _ = v10
    default:
        if (65535 >= 4294967295) {
            if (1.5 <= float32(3.14159)) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        p0 = (^(p0))
                        p0 -= p0
                        _ = fi12
                    }
                    _ = fi11
                }
            } else {
                switch ("hello" + (("\n\t" + "\n\t") + "a")) {
                case "\n\t":
                    p1 -= p1
                case "foo bar":
                    switch ("" + ("\n\t" + "a")) {
                    case "a":
                        if (-2.25 >= 0.1) {
                            v13 := []int{}
                            for fi14 := 0; fi14 < 9; fi14++ {
                                p0 = (p0 ^ p0)
                                if ((p1 + p1) != p1) {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                } else {
                                    if true {
                                        fmt.Println(p0)
                                    } else {
                                        fmt.Println(p0)
                                    }
                                    if (!false) {
                                        var v15 int32 = -293
                                        for fi16 := 0; fi16 < 19; fi16++ {
                                            v15 = 198
                                            fmt.Println(p1)
                                            _ = fi16
                                        }
                                        var v17 int16 = p1
                                        _ = v15
                                        _ = v17
                                    } else {
                                        var v18 uint = 1
                                        var v19 float64 = 3.14159
                                        _ = v18
                                        _ = v19
                                    }
                                }
                                _ = fi14
                            }
                            _ = v13
                        }
                    case "foo bar":
                        fmt.Println(p1)
                    default:
                        for fi20 := 0; fi20 < 11; fi20++ {
                            var v21 uint = 255
                            _ = fi20
                            _ = v21
                        }
                    }
                case "hello":
                    var v22 uint8 = 0
                    _ = v22
                default:
                    var v23 string = "hello"
                    _ = v23
                }
                p0 = p0
                var v24 uint16 = p0
                _ = v24
            }
            v25 := []int{int(p1), 127, 9223372036854775807}
            var v26 int64 = 0
            _ = v25
            _ = v26
        } else {
            fmt.Println(p0)
        }
    }
    p1 = (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = v9
    return ""
}

func main() {
    var v1 uint = 52
    var v2 string = "a"
    var v3 bool = true
    var v4 int64 = 0
    var v5 int32 = 0
    if ((^(v1)) <= 0) {
        if (int(v4) == 4294967295) {
            v27 := GenS3{F0: 417}
            switch 128 {
            case 2147483647:
                switch 1 {
                case 0:
                    v3 = (-128 != 1)
                case 42:
                    if (!(!v3)) {
                        for fi28 := 0; fi28 < 14; fi28++ {
                            var v29 int16 = 42
                            _ = fi28
                            _ = v29
                        }
                        v30 := []int{-9223372036854775808, -308, -9223372036854775808}
                        _ = v30
                    }
                default:
                    for fi31 := 0; fi31 < 9; fi31++ {
                        fmt.Println(v4)
                        v4 -= v4
                        var v32 float64 = 1e10
                        _ = fi31
                        _ = v32
                    }
                }
            case 136:
                v33 := GenS2{F0: v5, F1: 4294967295}
                _ = v33
            case 128:
                var v34 bool = (((1 <= (v5 | v5)) && (!(1 >= 1))) || ((v1 != (v1 >> ((42) & 63))) || (true || true)))
                _ = v34
            }
            if true {
                if ((156 < 128) && (446 <= 2147483647)) {
                    fmt.Println(v1)
                    v35 := GenS4{F0: -6, F1: 65535}
                    _ = v35
                }
                if v3 {
                    var v36 int16 = 32767
                    for fi37 := 0; fi37 < 5; fi37++ {
                        fmt.Println(v2, v4)
                        _ = fi37
                    }
                    v38 := GenS5{F0: 9223372036854775807, F1: (v36 * v36)}
                    _ = v36
                    _ = v38
                } else {
                    switch "" {
                    case "foo bar":
                        switch 1 {
                        case 0:
                            fmt.Println(v4, v5)
                        default:
                            v5 -= (-(v5))
                        }
                    default:
                        if v3 {
                            fmt.Println(v5, v3, v2)
                            for fi39 := 0; fi39 < 18; fi39++ {
                                var v40 float64 = ((-2.25 + 0.0) + 0.1)
                                switch v2 {
                                case "a":
                                    switch (v2 + v2) {
                                    case "foo bar":
                                        for fi41 := 0; fi41 < 14; fi41++ {
                                            v40 = (v40 - (1.5 - v40))
                                            if (!false) {
                                                var v42 uint = (v1 % ((v1) | 1))
                                                v5 = (v5 / (((v5 % (((v5 * int32(v4))) | 1))) | 1))
                                                v43 := GenS6{F0: v42}
                                                _ = v42
                                                _ = v43
                                            }
                                            _ = fi41
                                        }
                                    default:
                                        if true {
                                            if (!(0 != 100)) {
                                                if (1e10 == 0.0) {
                                                    v44 := []int{fi39, (fi39 | (fi39 | (fi39 / ((fi39) | 1)))), (fi39 + (^(fi39)))}
                                                    _ = v44
                                                }
                                                for fi45 := 0; fi45 < 13; fi45++ {
                                                    v46 := []int{}
                                                    _ = fi45
                                                    _ = v46
                                                }
                                            }
                                            var v47 string = v2
                                            _ = v47
                                        }
                                    }
                                case "\n\t":
                                    for fi48 := 0; fi48 < 14; fi48++ {
                                        fmt.Println(v40, fi48, v5)
                                        _ = fi48
                                    }
                                case "foo bar":
                                    if v3 {
                                        if v3 {
                                            v49 := []int{}
                                            _ = v49
                                        } else {
                                            v50 := []int{}
                                            _ = v50
                                        }
                                        fmt.Println(v3, v5, v3)
                                    } else {
                                        if v3 {
                                            v1 -= v1
                                            v51 := GenS6{F0: (v1 >> (((^(v1))) & 63))}
                                            if true {
                                                fmt.Println(v1, v1)
                                                for fi52 := 0; fi52 < 4; fi52++ {
                                                    var v53 uint = 18446744073709551615
                                                    v53 -= (v53 >> ((65535) & 63))
                                                    _ = fi52
                                                    _ = v53
                                                }
                                            }
                                            _ = v51
                                        }
                                        for fi54 := 0; fi54 < 12; fi54++ {
                                            v55 := GenS7{F0: 0, F1: (fi54 ^ fi39), F2: uint32(v1)}
                                            var v56 string = v2
                                            _ = fi54
                                            _ = v55
                                            _ = v56
                                        }
                                        var v57 uint = 2
                                        _ = v57
                                    }
                                }
                                _ = fi39
                                _ = v40
                            }
                        }
                    }
                    v58 := []int{42, 127, 0, 32767}
                    _ = v58
                }
            } else {
                if v3 {
                    for fi59 := 0; fi59 < 19; fi59++ {
                        fmt.Println(v2, v4, v2)
                        var v60 float64 = (3.14159 + 1.5)
                        _ = fi59
                        _ = v60
                    }
                    fmt.Println(v1, v1, v1)
                } else {
                    v5 -= (v5 + v5)
                    v61 := []int{-323, 586}
                    _ = v61
                }
                v62 := GenS1{F0: v5}
                v63 := []int{32768, 0}
                _ = v62
                _ = v63
            }
            _ = v27
        }
        for fi64 := 0; fi64 < 20; fi64++ {
            var v65 int = (^(fi64))
            switch (^(fi64)) {
            case 32767:
                if v3 {
                    v66 := GenS1{F0: (v5 - 1)}
                    fmt.Println(v2, v4, v2)
                    _ = v66
                } else {
                    v67 := GenS1{F0: 1}
                    var v68 int8 = 81
                    v5 = 639
                    _ = v67
                    _ = v68
                }
            case 0:
                for fi69 := 0; fi69 < 8; fi69++ {
                    fmt.Println(v5, fi64, v4)
                    _ = fi69
                }
            case 1:
                var v70 int32 = (-(v5))
                _ = v70
            }
            if false {
                var v71 bool = v3
                v3 = (v71 || ((1000 > -407) || v3))
                _ = v71
            }
            _ = fi64
            _ = v65
        }
        v72 := GenS1{F0: v5}
        _ = v72
    }
    fmt.Println(v3, v1, v3)
    v73 := []int{}
    v2 = ("foo bar" + v2)
    v74 := GenS5{F0: -9223372036854775808, F1: 580}
    for fi75 := 0; fi75 < 18; fi75++ {
        v76 := GenS7{F0: -117, F1: fi75, F2: 975}
        switch (fi75 * fi75) {
        case 32768:
            var v77 uint = (v1 & (v1 + 42))
            _ = v77
        case -688:
            var v78 int = int(v1)
            _ = v78
        default:
            var v79 int = (fi75 ^ (fi75 | -2147483648))
            _ = v79
        }
        _ = fi75
        _ = v76
    }
    if (-746 != -9223372036854775808) {
        _ = len(v73)
    } else {
        v80 := GenS6{F0: v1}
        if (!(3.14159 != (0.0 * float32(-2.25)))) {
            v4 = (v4 - v4)
            for fi81 := 0; fi81 < 1; fi81++ {
                v3 = true
                _ = fi81
            }
        }
        _ = v80
    }
    v82 := GenS8{F0: 306, F1: v4, F2: (v4 - (v4 & v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v73:", len(v73), v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v73
    _ = v74
    _ = v82
}
