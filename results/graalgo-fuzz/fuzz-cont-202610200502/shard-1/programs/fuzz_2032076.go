package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 string }
type GenS3 struct { F0 int64; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 string; F1 uint }
func genFn0(p0 uint32, p1 int16) int16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            switch "\n\t" {
            case "":
                p1 -= (p1 * p1)
                var v7 int16 = (^(p1))
                v8 := []int{(fi5 + fi6), fi5, (^(fi5)), fi5}
                p0 -= 2147483647
                _ = v7
                _ = v8
            case "hello":
                var v9 float64 = 1.5
                _ = v9
            }
            v10 := GenS1{F0: 42, F1: 2}
            switch "\n\t" {
            case "foo bar":
                var v11 int64 = -9223372036854775808
                _ = v11
            }
            _ = fi6
            _ = v10
        }
        var v12 int8 = 100
        v13 := []int{2147483647, -1}
        _ = fi5
        _ = v12
        _ = v13
    }
    p0 += 362
    v14 := []int{-2147483648, 9223372036854775807, 32768}
    if (0 >= -2147483648) {
        p0 -= p0
        p1 += (p1 >> ((18446744073709551615) & 63))
        v15 := GenS2{F0: 97, F1: 100, F2: (("a" + "foo bar") + "a")}
        _ = v15
    } else {
        v16 := GenS3{F0: 0, F1: -1, F2: p0}
        if len(v14) > 0 {
            v14[0] = 127
        }
        _ = v16
    }
    var v17 int = -507
    _ = p0
    _ = p1
    _ = v14
    _ = v17
    return -32768
}

func main() {
    var v1 string = ""
    var v2 float64 = 0.0
    v3 := []string{(v1 + (("hello" + v1) + (v1 + (((v1 + "foo bar") + (v1 + "")) + v1)))), ("foo bar" + "hello")}
    var v4 int8 = -51
    v18 := GenS4{F0: 218, F1: 9223372036854775807}
    v19 := GenS2{F0: (v4 | (v4 % ((-128) | 1))), F1: 0, F2: ((((("foo bar" + (v1 + v1)) + v1) + (v1 + v1)) + v1) + ("" + ("\n\t" + v1)))}
    if (true && false) {
        if true {
            if true {
                v2 = 1e10
                var v20 int32 = 873
                _ = v20
            }
            v21 := GenS3{F0: 0, F1: 42, F2: 586}
            _ = v21
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            if (v2 < v2) {
                var v23 int8 = (v4 << ((18446744073709551615) & 63))
                v2 = v2
                _ = v23
            } else {
                v4 = (v4 * v4)
                if false {
                    var v24 int32 = -138
                    v24 -= (v24 - v24)
                    if false {
                        v25 := GenS2{F0: v4, F1: v4, F2: "hello"}
                        if false {
                            fmt.Println(v24)
                            v26 := GenS2{F0: v4, F1: (v4 & v4), F2: ((v1 + v1) + ((v1 + "\n\t") + ("\n\t" + v1)))}
                            if ((!false) && ((!(1 != 1)) || true)) {
                                var v27 string = (v1 + ("a" + "hello"))
                                _ = v27
                            }
                            _ = v26
                        } else {
                            v28 := GenS2{F0: v4, F1: (v4 * (v4 - (-(v4)))), F2: (((((v1 + v1) + "foo bar") + v1) + "\n\t") + ((v1 + v1) + ("hello" + (v1 + v1))))}
                            _ = v28
                        }
                        _ = v25
                    }
                    _ = v24
                } else {
                    if true {
                        if len(v3) > 0 {
                            v3[0] = v1
                        }
                        switch 32767 {
                        case 42:
                            v1 = v1
                        default:
                            var v29 float32 = (0.0 + 1e10)
                            _ = v29
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v4, v4)
                        fmt.Println(v4, fi22)
                    }
                }
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 int16 = 1000
                    _ = fi30
                    _ = v31
                }
            }
            _ = fi22
        }
        v32 := GenS5{F0: 0}
        _ = v32
    }
    v4 -= v4
    fmt.Println(v1, v1)
    v2 = float64(1e10)
    var v33 float32 = 0.0
    for fi34 := 0; fi34 < 17; fi34++ {
        fmt.Println(v1)
        v35 := GenS2{F0: v4, F1: v4, F2: "\n\t"}
        _ = fi34
        _ = v35
    }
    v36 := GenS5{F0: 65535}
    if (v1 == (("" + ("" + "hello")) + ("foo bar" + "a"))) {
        for fi37 := 0; fi37 < 17; fi37++ {
            v3 = append(v3, v1)
            if (1 <= 1) {
                var v38 bool = (385 <= 42)
                switch 128 {
                case 0:
                    if true {
                        if v38 {
                            v4 -= (v4 | (v4 % ((42) | 1)))
                            for fi39 := 0; fi39 < 20; fi39++ {
                                v40 := GenS2{F0: 1, F1: v4, F2: v1}
                                _ = fi39
                                _ = v40
                            }
                            var v41 float32 = v33
                            _ = v41
                        } else {
                            _ = len(v3)
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v42 string = v1
                        _ = v42
                    } else {
                        switch v1 {
                        case "foo bar":
                            var v43 int64 = 42
                            _ = v43
                        case "hello":
                            v4 = v4
                        default:
                            var v44 uint32 = 2147483647
                            _ = v44
                        }
                        if v38 {
                            for fi45 := 0; fi45 < 20; fi45++ {
                                v2 += v2
                                fmt.Println(v4)
                                _ = fi45
                            }
                            for fi46 := 0; fi46 < 6; fi46++ {
                                if (376 >= 0) {
                                    switch (fi46 - (fi37 ^ (fi37 * fi46))) {
                                    case 453:
                                        fmt.Println(v2, v4)
                                    default:
                                        if ((v33 + v33) < -2.25) {
                                            if v38 {
                                                if v38 {
                                                    v47 := GenS1{F0: int32(fi46), F1: 0}
                                                    _ = len(v3)
                                                    v48 := GenS2{F0: v4, F1: -1, F2: (((v1 + v1) + "foo bar") + ("hello" + "a"))}
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    v33 = v33
                                                }
                                                if v38 {
                                                    fmt.Println(v2, fi37, v38)
                                                    v4 += (^(v4))
                                                } else {
                                                    v49 := GenS6{F0: 0, F1: -1}
                                                    _ = v49
                                                }
                                                v3 = append(v3, (v1 + v1))
                                                switch v1 {
                                                case "hello":
                                                    v1 = v1
                                                default:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                var v50 uint64 = 42
                                                _ = v50
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v51 := GenS2{F0: -1, F1: v4, F2: (v1 + v1)}
                                                _ = v51
                                            }
                                        }
                                        switch v1 {
                                        case "foo bar":
                                            if false {
                                                var v52 string = (((v1 + v1) + ((v1 + v1) + v1)) + v1)
                                                if ((-(v4)) == v4) {
                                                    v52 = (v1 + "hello")
                                                    if (-1 == 0) {
                                                        if v38 {
                                                            v33 += v33
                                                            var v53 float64 = v2
                                                            var v54 uint = 1
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                    }
                                                }
                                                _ = v52
                                            }
                                        case "\n\t":
                                            var v55 int32 = 68
                                            _ = v55
                                        case "":
                                            fmt.Println(fi37, fi37, v38)
                                        default:
                                            if true {
                                                var v56 int16 = 506
                                                _ = v56
                                            }
                                        }
                                    }
                                    v2 = v2
                                } else {
                                    if (0 >= 42) {
                                        v3 = append(v3, (v1 + v1))
                                    } else {
                                        if (v38 && v38) {
                                            var v57 uint32 = 229
                                            _ = v57
                                        }
                                        v33 += float32(v2)
                                        if len(v3) > 0 {
                                            v3[0] = (v1 + ((v1 + v1) + (v1 + "a")))
                                        }
                                        v4 = (v4 % ((v4) | 1))
                                    }
                                    v3 = append(v3, "hello")
                                    var v58 float32 = (((0.1 - v33) + 1.5) / v33)
                                    _ = v58
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi46
                            }
                            v4 = (v4 / ((v4) | 1))
                        } else {
                            v2 += v2
                            var v59 float32 = float32(((((v2 / v2) * v2) / 3.14159) - ((v2 - v2) * float64((float32(float64(-2.25)) / v33)))))
                            if (2147483647 == 256) {
                                v60 := GenS7{F0: "a", F1: 1}
                                var v61 float64 = 0.1
                                _ = v60
                                _ = v61
                            }
                            _ = v59
                        }
                    }
                }
                _ = v38
            }
            _ = fi37
        }
        _ = len(v3)
        var v62 int16 = 42
        _ = v62
    }
    if ((float32(0.1) - (1e10 / v33)) == v33) {
        var v63 float64 = (0.1 * v2)
        v1 = (v1 + (v1 + ("\n\t" + (v1 + v1))))
        _ = v63
    }
    _ = len(v3)
    var v64 int = 723
    var v65 int32 = 2147483647
    if (255 == 1) {
        fmt.Println(v33)
        if false {
            v65 = (-(v65))
        }
        var v66 float64 = v2
        _ = v66
    }
    switch v64 {
    case -2147483648:
        if true {
            v3 = append(v3, v1)
        } else {
            for fi67 := 0; fi67 < 9; fi67++ {
                v1 = ((v1 + "") + "foo bar")
                v1 = "foo bar"
                v2 -= v2
                _ = fi67
            }
        }
    default:
        for fi68 := 0; fi68 < 3; fi68++ {
            switch "foo bar" {
            case "\n\t":
                for fi69 := 0; fi69 < 6; fi69++ {
                    v70 := GenS4{F0: 122, F1: 9223372036854775807}
                    var v71 uint16 = 0
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            }
            fmt.Println(fi68, v64)
            switch (-(v64)) {
            case 653:
                var v72 int8 = v4
                _ = v72
            case 127:
                if ((3.14159 + (0.0 + ((0.0 - 1.5) / (-2.25 + float64(1e10))))) != ((-2.25 - v2) / float64(v33))) {
                    v2 = ((v2 / (v2 + (v2 / float64(float32(0.1))))) * 0.0)
                } else {
                    v33 = v33
                    for fi73 := 0; fi73 < 2; fi73++ {
                        fmt.Println(v33, v65)
                        v74 := GenS4{F0: 128, F1: 9223372036854775807}
                        _ = fi73
                        _ = v74
                    }
                }
            default:
                v2 = v2
            }
            _ = fi68
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v4
    _ = v64
    _ = v65
}
