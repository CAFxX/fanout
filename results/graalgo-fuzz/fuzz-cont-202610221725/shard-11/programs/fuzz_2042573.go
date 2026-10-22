package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 string; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 int64; F1 int16; F2 int8 }
type GenS8 struct { F0 uint32; F1 string; F2 int8 }
type GenS9 struct { F0 bool; F1 int64; F2 uint }
func genFn0(p0 string, p1 uint, p2 uint8) int {
    var v7 float64 = 1e10
    v8 := []int{-329, 781, -9223372036854775808, -166}
    var v9 float32 = float32(v7)
    v9 = 1.5
    var v10 float32 = ((3.14159 - v9) + v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 string = ("" + "\n\t")
    v2 := make(map[int]int)
    var v3 bool = false
    v4 := []uint32{0, 2147483648}
    v5 := make(map[string]int)
    v6 := make(map[int]string)
    fmt.Println(v3, v3)
    if (0 >= -128) {
        if v3 {
            v3 = (!(0 <= -979))
            v11 := GenS1{F0: 255}
            _ = v11
        }
        v1 = (v1 + v1)
        switch v1 {
        case "hello":
            for fi12 := 0; fi12 < 10; fi12++ {
                if (((genFn0("\n\t", 105, 12) < (fi12 << ((1) & 63))) && false) && (((!(!(472 > 1))) && (-2147483648 < -2147483648)) || v3)) {
                    v1 = v1
                } else {
                    v3 = true
                    v1 = v1
                    var v13 float32 = 1e10
                    _ = v13
                }
                _ = fi12
            }
        case "":
            for fi14 := 0; fi14 < 10; fi14++ {
                v3 = (107 > -1)
                v1 = (v1 + v1)
                _ = fi14
            }
        }
    }
    v15 := GenS2{F0: -32768, F1: (v1 + "hello"), F2: 0}
    for fi16 := 0; fi16 < 10; fi16++ {
        fmt.Println(fi16, v1)
        _ = fi16
    }
    switch "\n\t" {
    case "hello":
        var v17 uint16 = 0
        _ = v17
    case "":
        switch "\n\t" {
        case "\n\t":
            var v18 string = v1
            _ = v18
        default:
            for fi19 := 0; fi19 < 14; fi19++ {
                switch genFn0(("" + ((("" + ("foo bar" + v1)) + (((v1 + (("a" + v1) + "hello")) + v1) + v1)) + (((((v1 + "foo bar") + ((v1 + ("" + v1)) + ("hello" + v1))) + ("a" + v1)) + ((("foo bar" + v1) + "a") + v1)) + ""))), 255, 128) {
                case -1:
                    v20 := GenS3{F0: 128}
                    _ = v20
                case 2147483647:
                    switch genFn0("\n\t", 2, 127) {
                    case 4294967295:
                        if (211 < 255) {
                            v1 = "hello"
                        } else {
                            var v21 uint64 = 0
                            _ = v21
                        }
                    case 0:
                        if v3 {
                            v3 = true
                            switch (v1 + v1) {
                            case "\n\t":
                                if (1 >= 151) {
                                    if v3 {
                                        for fi22 := 0; fi22 < 20; fi22++ {
                                            switch fi22 {
                                            case -380:
                                                var v23 int8 = -20
                                                _ = v23
                                            default:
                                                fmt.Println(fi22, fi22)
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            if (!(v3 || v3)) {
                                                _ = len(v4)
                                                v3 = v3
                                            } else {
                                                _ = len(v4)
                                                var v24 int64 = -758
                                                v24 = (v24 % ((-1) | 1))
                                                _ = v24
                                            }
                                            _ = len(v4)
                                            _ = len(v4)
                                            _ = fi22
                                        }
                                    } else {
                                        switch (fi19 + (fi19 | fi19)) {
                                        case 0:
                                            var v25 int8 = 42
                                            _ = v25
                                        default:
                                            v1 = v1
                                        }
                                        for fi26 := 0; fi26 < 16; fi26++ {
                                            fmt.Println(v3, fi19)
                                            v3 = v3
                                            var v27 string = v1
                                            _ = fi26
                                            _ = v27
                                        }
                                    }
                                } else {
                                    for fi28 := 0; fi28 < 13; fi28++ {
                                        v3 = ((fi19 / ((genFn0(v1, 963, 127)) | 1)) > (fi19 & genFn0(v1, 18446744073709551615, 155)))
                                        if len(v4) > 0 {
                                            v4[0] = 394
                                        }
                                        _ = fi28
                                    }
                                    var v29 int32 = 1
                                    _ = v29
                                }
                            case "a":
                                var v30 uint16 = 0
                                _ = v30
                            default:
                                for fi31 := 0; fi31 < 6; fi31++ {
                                    switch fi19 {
                                    case 4294967295:
                                        v1 = ((v1 + ((("a" + v1) + (("\n\t" + ("a" + v1)) + v1)) + ("hello" + ((((v1 + v1) + "foo bar") + (((v1 + (("" + v1) + "hello")) + "hello") + v1)) + (((v1 + v1) + v1) + v1))))) + "")
                                    default:
                                        v1 = "hello"
                                    }
                                    _ = fi31
                                }
                            }
                            if v3 {
                                v32 := GenS1{F0: 2}
                                v33 := GenS4{F0: 1, F1: -1}
                                _ = v32
                                _ = v33
                            } else {
                                v4 = append(v4, 65535)
                                v1 = (v1 + v1)
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    _ = len(v5)
                                    _ = fi34
                                }
                            }
                        } else {
                            if (!v3) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                if (!(((0.1 + float64(1e10)) - -2.25) < 0.0)) {
                                    switch v1 {
                                    case "\n\t":
                                        v1 = v1
                                    }
                                }
                                var v35 string = ("a" + v1)
                                _ = v35
                            }
                            var v36 uint64 = 1
                            _ = v36
                        }
                    }
                }
                v3 = v3
                if (700 >= -559) {
                    var v37 string = v1
                    v3 = v3
                    v1 = "a"
                    _ = v37
                } else {
                    v1 = (v1 + ((("hello" + (v1 + v1)) + ((((v1 + "\n\t") + v1) + v1) + "a")) + v1))
                }
                _ = fi19
            }
        }
    case "a":
        var v38 int = -9223372036854775808
        _ = v38
    default:
        v1 = ("foo bar" + "a")
    }
    if (-32768 != 1000) {
        var v39 int8 = 1
        v40 := GenS5{F0: 458, F1: "\n\t"}
        var v41 uint16 = 32767
        _ = v39
        _ = v40
        _ = v41
    }
    switch 2147483647 {
    case -62:
        if v3 {
            v3 = v3
            if v3 {
                v3 = (v3 || (-82 >= 0))
                if (!((2 > 1) || true)) {
                    if true {
                        if (-941 != -281) {
                            var v42 uint64 = 42
                            _ = v42
                        } else {
                            if (!v3) {
                                v3 = (1 < 888)
                            }
                            var v43 float32 = 3.14159
                            _ = v43
                        }
                    }
                    for fi44 := 0; fi44 < 11; fi44++ {
                        v1 = v1
                        fmt.Println(v1, fi44)
                        _ = fi44
                    }
                    v45 := v5[v1]
                    _ = v45
                }
                var v46 int8 = -1
                _ = v46
            }
            for fi47 := 0; fi47 < 3; fi47++ {
                v48 := GenS4{F0: 123, F1: 914}
                v49 := GenS6{F0: 438, F1: 1}
                _ = fi47
                _ = v48
                _ = v49
            }
        }
    case -1:
        if (721 > -9223372036854775808) {
            switch 636 {
            case 2147483647:
                var v50 int64 = -280
                _ = v50
            case -384:
                v51 := GenS7{F0: -1, F1: -1, F2: 1}
                _ = v51
            case 32767:
                if v3 {
                    if (130 == 2) {
                        var v52 float32 = 1.5
                        v52 = (1.5 + (((v52 / 1.0) * float32(float64(0.1))) + v52))
                        _ = v52
                    } else {
                        for fi53 := 0; fi53 < 20; fi53++ {
                            var v54 int = (fi53 + -307)
                            _ = fi53
                            _ = v54
                        }
                        fmt.Println(v1, v1)
                        v55 := GenS7{F0: -37, F1: -32768, F2: 126}
                        _ = v55
                    }
                    if (4294967295 <= 0) {
                        var v56 uint8 = 128
                        for fi57 := 0; fi57 < 5; fi57++ {
                            var v58 bool = (0 >= 0)
                            var v59 string = "a"
                            v60 := GenS7{F0: -9223372036854775808, F1: -32768, F2: -128}
                            _ = fi57
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        _ = v56
                    }
                    v61 := GenS8{F0: 2, F1: v1, F2: 127}
                    _ = v61
                }
            default:
                delete(v6, genFn0((v1 + "hello"), 639, 128))
            }
            var v62 uint16 = 1
            var v63 bool = (false || false)
            _ = v62
            _ = v63
        }
    case 184:
        var v64 float64 = 3.14159
        _ = v64
    default:
        for fi65 := 0; fi65 < 15; fi65++ {
            var v66 int32 = -2147483648
            for fi67 := 0; fi67 < 1; fi67++ {
                fmt.Println(v3, v1)
                var v68 uint8 = 1
                _ = fi67
                _ = v68
            }
            _ = fi65
            _ = v66
        }
    }
    v1 = v1
    fmt.Println(v1, v3)
    if false {
        v1 = v1
        switch "a" {
        case "\n\t":
            if v3 {
                for fi69 := 0; fi69 < 13; fi69++ {
                    var v70 int8 = -60
                    v71 := v5[v1]
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                delete(v6, -9223372036854775808)
                v1 = v1
            }
        case "hello":
            v1 = v1
        case "a":
            for fi72 := 0; fi72 < 4; fi72++ {
                for fi73 := 0; fi73 < 8; fi73++ {
                    if false {
                        v3 = (-1 == 835)
                        if v3 {
                            v3 = (1e10 <= 0.1)
                        }
                    } else {
                        fmt.Println(fi73, v3, fi73)
                        v1 = "hello"
                    }
                    if v3 {
                        fmt.Println(fi73, fi72)
                    } else {
                        v3 = true
                        v74 := GenS9{F0: ((^(fi73)) == (fi73 - fi72)), F1: 42, F2: 18446744073709551615}
                        switch fi72 {
                        case 596:
                            var v75 uint32 = 1
                            _ = v75
                        case 0:
                            v1 = v1
                        }
                        _ = v74
                    }
                    var v76 int64 = 0
                    v3 = false
                    _ = fi73
                    _ = v76
                }
                _ = fi72
            }
        default:
            v77 := GenS7{F0: 42, F1: 215, F2: -128}
            _ = v77
        }
        v78 := v2[70]
        _ = v78
    } else {
        for fi79 := 0; fi79 < 11; fi79++ {
            switch fi79 {
            case -2147483648:
                if (1 < 1) {
                    v3 = (1 == 2)
                } else {
                    var v80 float64 = 1.5
                    _ = v80
                }
            case -9223372036854775808:
                v1 = v1
            default:
                var v81 float32 = (1e10 * 3.14159)
                _ = v81
            }
            for fi82 := 0; fi82 < 13; fi82++ {
                if (-844 <= 42) {
                    fmt.Println(fi79)
                }
                v1 = ("" + (((("hello" + v1) + "a") + v1) + (v1 + (v1 + v1))))
                _ = fi82
            }
            _ = fi79
        }
    }
    v1 = "a"
    v83 := GenS7{F0: -1, F1: -32768, F2: 56}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
}
