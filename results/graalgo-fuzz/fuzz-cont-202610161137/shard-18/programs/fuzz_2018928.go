package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 int }
type GenS3 struct { F0 uint8; F1 uint64; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS7 struct { F0 string; F1 uint; F2 uint64 }
type GenS8 struct { F0 int64; F1 uint32; F2 bool }
type GenS9 struct { F0 uint64; F1 uint8; F2 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32; F1 uint16; F2 int32 }
type GenS12 struct { F0 uint8; F1 int; F2 int32 }
func genFn0(p0 int8, p1 int16, p2 uint) int32 {
    var v7 float64 = 0.0
    v8 := GenS1{F0: (false && true), F1: 127, F2: 18446744073709551615}
    if false {
        fmt.Println(p2)
        var v9 int16 = (^(p1))
        v10 := GenS1{F0: false, F1: 128, F2: p2}
        _ = v10
        _ = v9
    } else {
        switch -9223372036854775808 {
        case -269:
            v7 = ((0.1 + 0.1) * float64(((0.1 / (-2.25 / ((0.1 * (1e10 / -2.25)) * -2.25))) - 0.0)))
        case 128:
            var v11 float64 = float64(0.1)
            _ = v11
        case 4294967295:
            v12 := GenS1{F0: true, F1: 42, F2: (p2 << ((42) & 63))}
            _ = v12
        }
        switch "a" {
        case "foo bar":
            var v13 uint16 = 1
            _ = v13
        default:
            switch "" {
            case "hello":
                var v14 int = 127
                _ = v14
            case "a":
                fmt.Println(p1, p1, p0)
            }
        }
    }
    v15 := GenS1{F0: (-605 != -964), F1: 255, F2: p2}
    if (true && (p1 >= 1)) {
        switch "\n\t" {
        case "":
            v16 := []int{}
            _ = v16
        case "a":
            fmt.Println(v7, p2)
        default:
            if (p0 <= p0) {
                var v17 bool = true
                _ = v17
            }
        }
        v7 = float64((-2.25 + 1.5))
        v18 := []int{4294967295, 205, -964}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    _ = v8
    return int32(p0)
}
func genFn1(p0 uint32) uint8 {
    v19 := GenS2{F0: 100, F1: (^(p0)), F2: 42}
    fmt.Println(p0)
    _ = p0
    _ = v19
    return 128
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = true
    var v3 int32 = 42
    var v4 float64 = 0.1
    var v5 string = ((("a" + "foo bar") + "") + "hello")
    v6 := []int{}
    fmt.Println(v2, v4)
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    _ = len(v6)
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int64 = -9223372036854775808
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        fmt.Println(fi22, v4, v1)
        switch fi22 {
        case 2147483647:
            if true {
                if false {
                    switch fi22 {
                    case 4294967295:
                        if v2 {
                            v5 = "hello"
                        }
                    case 127:
                        if (4294967295 > 526) {
                            fmt.Println(v4, v2, v4)
                            v3 += (v3 / ((v3) | 1))
                        } else {
                            v6 = append(v6, fi22)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            for fi23 := 0; fi23 < 10; fi23++ {
                                if (0 < 0) {
                                    v5 = v5
                                } else {
                                    switch 127 {
                                    case 923:
                                        var v24 string = (v5 + "foo bar")
                                        _ = v24
                                    case 0:
                                        fmt.Println(fi22, v3)
                                    case -2147483648:
                                        v4 = float64(0.1)
                                    default:
                                        switch (v5 + (v5 + "hello")) {
                                        case "hello":
                                            switch int(v3) {
                                            case 1:
                                                switch -2147483648 {
                                                case -399:
                                                    v5 = v5
                                                }
                                            }
                                        case "\n\t":
                                            var v25 string = ("foo bar" + v5)
                                            _ = v25
                                        default:
                                            var v26 uint16 = 2
                                            _ = v26
                                        }
                                    }
                                    if v2 {
                                        v3 = v3
                                    }
                                }
                                v4 += v4
                                _ = fi23
                            }
                        }
                    case 1:
                        v27 := GenS3{F0: 255, F1: (v1 / ((2) | 1)), F2: (^(fi22))}
                        _ = v27
                    default:
                        v1 -= (^(v1))
                    }
                } else {
                    v6 = append(v6, fi22)
                }
                v1 -= (v1 / ((v1) | 1))
            }
        case 9223372036854775807:
            if v2 {
                var v28 uint16 = 2
                _ = v28
            } else {
                switch (((v5 + v5) + (v5 + "foo bar")) + v5) {
                case "\n\t":
                    for fi29 := 0; fi29 < 1; fi29++ {
                        v30 := GenS3{F0: 108, F1: 501, F2: fi22}
                        if false {
                            v3 = (v3 * v3)
                            if (2 < uint16(v1)) {
                                fmt.Println(v5)
                                switch (fi29 >> ((2) & 63)) {
                                case 4294967295:
                                    if len(v6) > 0 {
                                        v6[0] = fi29
                                    }
                                case 9223372036854775807:
                                    var v31 uint32 = 2
                                    _ = v31
                                case 0:
                                    v5 = v5
                                }
                                var v32 uint32 = 623
                                fmt.Println(v4, fi29)
                                _ = v32
                            }
                        } else {
                            v1 = (v1 ^ v1)
                        }
                        if (18446744073709551615 != 2) {
                            v4 = v4
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            switch "a" {
                            case "a":
                                if (1 == 1) {
                                    if true {
                                        v33 := GenS4{F0: fi29}
                                        v34 := GenS2{F0: 127, F1: 1, F2: fi29}
                                        v6 = append(v6, 424)
                                        _ = v33
                                        _ = v34
                                    }
                                    if true {
                                        v35 := GenS5{F0: 233, F1: v1}
                                        switch fi29 {
                                        case 42:
                                            switch fi22 {
                                            case 637:
                                                v36 := GenS6{F0: 1, F1: 837, F2: 788}
                                                _ = v36
                                            }
                                        case 32768:
                                            fmt.Println(fi22)
                                        case -187:
                                            _ = len(v6)
                                        default:
                                            v6 = append(v6, -1)
                                        }
                                        v37 := GenS6{F0: 0, F1: 1, F2: 2147483648}
                                        _ = v35
                                        _ = v37
                                    } else {
                                        v38 := GenS7{F0: (("" + (v5 + v5)) + v5), F1: 65535, F2: v1}
                                        _ = v38
                                    }
                                    v39 := GenS2{F0: -1, F1: 0, F2: -2147483648}
                                    _ = v39
                                }
                            }
                            if (true || v2) {
                                v6 = append(v6, -1)
                                v40 := GenS8{F0: -660, F1: 2147483647, F2: false}
                                v3 -= v3
                                _ = v40
                            }
                            fmt.Println(fi29)
                        }
                        _ = fi29
                        _ = v30
                    }
                case "":
                    for fi41 := 0; fi41 < 7; fi41++ {
                        v6 = append(v6, (-(fi41)))
                        var v42 int32 = v3
                        v43 := GenS9{F0: (^(v1)), F1: 128, F2: (v5 + "foo bar")}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                case "a":
                    if v2 {
                        var v44 uint16 = 421
                        var v45 float32 = (float32((v4 + v4)) + 1.5)
                        _ = v44
                        _ = v45
                    } else {
                        if v2 {
                            v46 := GenS9{F0: (v1 & (v1 ^ (v1 % (((v1 * v1)) | 1)))), F1: 183, F2: v5}
                            switch v5 {
                            case "":
                                var v47 uint16 = 65535
                                _ = v47
                            case "a":
                                switch (-(fi22)) {
                                case 32767:
                                    fmt.Println(v3)
                                case 1:
                                    fmt.Println(v3)
                                case 828:
                                    if (!(v2 && (fi22 > (fi22 + (fi22 >> ((255) & 63)))))) {
                                        var v48 uint = 0
                                        fmt.Println(v5, v3, v2)
                                        _ = v48
                                    } else {
                                        fmt.Println(v3)
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v49 := GenS10{F0: (!true)}
                                        _ = v49
                                    }
                                }
                            }
                            v50 := GenS1{F0: (1 <= 127), F1: 42, F2: 2}
                            _ = v46
                            _ = v50
                        } else {
                            v3 = (v3 * -571)
                            v6 = append(v6, (fi22 / ((fi22) | 1)))
                            v5 = ((v5 + "hello") + "foo bar")
                        }
                    }
                }
                v3 = v3
            }
        default:
            v1 = 18446744073709551615
        }
        _ = fi22
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        fmt.Println(v2, v4)
        _ = fi51
    }
    for fi52 := 0; fi52 < 18; fi52++ {
        fmt.Println(fi52, v5, v3)
        _ = fi52
    }
    switch "hello" {
    case "a":
        switch 9223372036854775807 {
        case 128:
            v53 := GenS11{F0: v3, F1: 32767, F2: -2147483648}
            _ = v53
        case 926:
            v1 = v1
        case -9223372036854775808:
            var v54 float64 = (v4 + v4)
            _ = v54
        }
    }
    v1 += (^(v1))
    var v55 int64 = -1
    fmt.Println(v55, v3)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v2 = v2
    for fi56 := 0; fi56 < 1; fi56++ {
        var v57 uint32 = 0
        var v58 float32 = ((3.14159 + 1.5) - float32(v4))
        v5 = v5
        _ = fi56
        _ = v57
        _ = v58
    }
    v2 = (98 > int8(v1))
    if v2 {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        switch ((((v5 + v5) + v5) + v5) + v5) {
        case "\n\t":
            switch "foo bar" {
            case "foo bar":
                var v59 int = 724
                _ = v59
            }
        case "a":
            if (-47 == -1) {
                v60 := GenS12{F0: 2, F1: -813, F2: v3}
                var v61 bool = false
                _ = v60
                _ = v61
            } else {
                v62 := GenS12{F0: 2, F1: 32767, F2: (v3 >> ((161) & 63))}
                if v2 {
                    var v63 uint8 = 127
                    v3 = v3
                    _ = v63
                }
                _ = v62
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
