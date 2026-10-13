package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            switch "a" {
            case "foo bar":
                var v6 int64 = -9223372036854775808
                _ = v6
            case "a":
                var v7 string = "hello"
                _ = v7
            default:
                v8 := []int{}
                _ = v8
            }
            _ = fi5
        }
        var v9 string = (("" + "foo bar") + "")
        _ = fi4
        _ = v9
    }
    v10 := []int{4294967295}
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: -607, F1: 2}
        _ = fi11
        _ = v12
    }
    v10 = append(v10, 2147483647)
    _ = v10
    return 2147483647
}
func genFn1(p0 int16, p1 uint64) int64 {
    var v13 uint = 309
    var v14 uint = (^(v13))
    v15 := []int{-609, -673}
    v15 = append(v15, 128)
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := []string{("a" + "a"), ("foo bar" + "a")}
    v2 := []int{752, -826}
    var v3 uint64 = 0
    for fi16 := 0; fi16 < 14; fi16++ {
        v3 += (v3 << ((552) & 63))
        _ = fi16
    }
    switch (("hello" + "foo bar") + (("" + "") + "a")) {
    case "":
        var v17 int8 = 102
        _ = v17
    case "\n\t":
        switch 294 {
        case 32767:
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int = fi18
                _ = fi18
                _ = v19
            }
        case 128:
            if ("a" < "") {
                for fi20 := 0; fi20 < 3; fi20++ {
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 uint16 = 65535
                        var v23 uint8 = 2
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    var v24 bool = false
                    _ = fi20
                    _ = v24
                }
                _ = len(v1)
            } else {
                var v25 uint = 2
                _ = v25
            }
        case 469:
            if (true && true) {
                if ((true || false) && (2 != 255)) {
                    var v26 int16 = 42
                    _ = v26
                }
                if (42 == 255) {
                    fmt.Println(v3)
                } else {
                    v27 := GenS1{F0: 1, F1: (v3 >> ((18446744073709551615) & 63))}
                    v28 := GenS1{F0: 0, F1: 18446744073709551615}
                    _ = v27
                    _ = v28
                }
                switch "a" {
                case "a":
                    v3 = (^(v3))
                case "hello":
                    if (-52 != 1) {
                        for fi29 := 0; fi29 < 11; fi29++ {
                            for fi30 := 0; fi30 < 15; fi30++ {
                                v3 += v3
                                _ = fi30
                            }
                            _ = fi29
                        }
                        v3 = (v3 ^ (v3 ^ v3))
                        switch "\n\t" {
                        case "":
                            if false {
                                var v31 uint64 = (v3 % (((^(v3))) | 1))
                                if len(v2) > 0 {
                                    v2[0] = 103
                                }
                                _ = v31
                            } else {
                                fmt.Println(v3)
                            }
                        case "foo bar":
                            if true {
                                for fi32 := 0; fi32 < 16; fi32++ {
                                    v33 := GenS1{F0: 9223372036854775807, F1: 4294967295}
                                    _ = fi32
                                    _ = v33
                                }
                                v34 := GenS1{F0: 650, F1: 18446744073709551615}
                                v3 += v3
                                _ = v34
                            }
                        }
                    }
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    v35 := GenS1{F0: 764, F1: v3}
                    _ = v35
                }
            } else {
                var v36 uint32 = 0
                var v37 uint = 586
                var v38 uint = 255
                _ = v36
                _ = v37
                _ = v38
            }
        }
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if true {
        v39 := GenS2{F0: (v3 % ((v3) | 1)), F1: (v3 & 1)}
        v3 += v3
        fmt.Println(v3)
        _ = v39
    } else {
        var v40 string = "a"
        v41 := GenS1{F0: 0, F1: v3}
        _ = v40
        _ = v41
    }
    v3 = v3
    var v42 int = 4294967295
    for fi43 := 0; fi43 < 20; fi43++ {
        v44 := GenS3{F0: 90, F1: (v3 >> ((742) & 63)), F2: v3}
        _ = fi43
        _ = v44
    }
    v42 -= (v42 << ((741) & 63))
    if ((0.1 - 3.14159) == 0.1) {
        for fi45 := 0; fi45 < 1; fi45++ {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            for fi46 := 0; fi46 < 15; fi46++ {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                switch v42 {
                case 359:
                    var v47 int = (^(v42))
                    _ = v47
                case 2147483647:
                    _ = len(v2)
                default:
                    switch "hello" {
                    case "hello":
                        v1 = append(v1, "")
                    case "a":
                        var v48 int32 = 371
                        _ = v48
                    }
                }
                v42 += fi46
                _ = fi46
            }
            if true {
                var v49 uint = 18446744073709551615
                v50 := GenS3{F0: 255, F1: (v3 + v3), F2: v3}
                _ = v49
                _ = v50
            }
            _ = fi45
        }
        if false {
            switch "foo bar" {
            case "\n\t":
                var v51 int32 = 146
                _ = v51
            case "hello":
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if true {
                        v42 = v42
                    } else {
                        var v52 uint64 = (v3 + v3)
                        v52 = 2
                        _ = v52
                    }
                }
            default:
                v1 = append(v1, ("hello" + ("\n\t" + "\n\t")))
            }
            switch (v42 - v42) {
            case 127:
                for fi53 := 0; fi53 < 9; fi53++ {
                    var v54 float32 = (1e10 * (float32(1e10) - (0.1 * 3.14159)))
                    var v55 int8 = 42
                    if true {
                        v56 := GenS3{F0: 128, F1: v3, F2: v3}
                        _ = v56
                    } else {
                        switch "\n\t" {
                        case "hello":
                            switch ((("\n\t" + "a") + (("foo bar" + "hello") + "hello")) + "") {
                            case "a":
                                if false {
                                    fmt.Println(v55, v55, v42)
                                }
                            case "\n\t":
                                if (((^(v3)) == v3) && true) {
                                    v55 = v55
                                    v42 -= (fi53 * 447)
                                    var v57 int64 = 42
                                    _ = v57
                                }
                            case "hello":
                                if (false && false) {
                                    v54 = 1.5
                                    if (!false) {
                                        v58 := GenS2{F0: 369, F1: (^(v3))}
                                        fmt.Println(v54, v3, v3)
                                        _ = v58
                                    }
                                }
                            default:
                                v2 = append(v2, fi53)
                            }
                        case "foo bar":
                            var v59 string = ((("\n\t" + "a") + "") + ("" + "hello"))
                            _ = v59
                        default:
                            v3 -= v3
                        }
                        var v60 float32 = 1.5
                        _ = v60
                    }
                    _ = fi53
                    _ = v54
                    _ = v55
                }
            case 318:
                if true {
                    var v61 int64 = -451
                    _ = v61
                } else {
                    fmt.Println(v3, v42)
                }
            default:
                v42 += v42
            }
            var v62 int = (v42 / ((v42) | 1))
            _ = v62
        } else {
            v3 = v3
            var v63 uint32 = 4294967295
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v63
        }
        var v64 int64 = 42
        _ = v64
    }
    if ((v3 - v3) != 1) {
        for fi65 := 0; fi65 < 6; fi65++ {
            var v66 uint32 = 357
            _ = fi65
            _ = v66
        }
    } else {
        switch (v42 ^ (^(v42))) {
        case -2147483648:
            v67 := GenS4{F0: 42, F1: genFn1(-1, (^(v3))), F2: 1000}
            _ = v67
        }
    }
    v42 -= v42
    v3 += (^(v3))
    v68 := GenS1{F0: -9223372036854775808, F1: 110}
    for fi69 := 0; fi69 < 17; fi69++ {
        if len(v2) > 0 {
            v2[0] = 127
        }
        var v70 int16 = 232
        _ = fi69
        _ = v70
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v68
}
