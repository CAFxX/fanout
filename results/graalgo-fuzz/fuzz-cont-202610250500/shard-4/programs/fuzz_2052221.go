package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint; F1 int32 }
func genFn0(p0 uint8, p1 uint16, p2 uint32) int64 {
    if true {
        p2 = (p2 ^ (p2 % (((p2 >> ((3) & 63))) | 1)))
        if (!((p0 << ((831) & 63)) != 1)) {
            if (416 >= 0) {
                v6 := GenS1{F0: 4294967295, F1: 1}
                _ = v6
            } else {
                var v7 uint = 4294967295
                v8 := GenS2{F0: p0}
                v9 := GenS1{F0: uint(p2), F1: 42}
                _ = v7
                _ = v8
                _ = v9
            }
            p0 += 127
        } else {
            fmt.Println(p1, p0)
            if false {
                var v10 uint64 = 1
                for fi11 := 0; fi11 < 18; fi11++ {
                    fmt.Println(p0)
                    v12 := []int{(fi11 | (fi11 ^ fi11)), 40}
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                v13 := GenS3{F0: 1, F1: (p1 & p1)}
                for fi14 := 0; fi14 < 18; fi14++ {
                    if (23 > 4294967295) {
                        p0 = (p0 | (p0 & 255))
                    }
                    for fi15 := 0; fi15 < 2; fi15++ {
                        var v16 int32 = -2147483648
                        v17 := GenS2{F0: 128}
                        var v18 string = ("hello" + "foo bar")
                        v16 = v16
                        var v19 int = (-(fi14))
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    fmt.Println(fi14, fi14)
                    _ = fi14
                }
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 float64 = float64(-2.25)
                    _ = fi20
                    _ = v21
                }
                _ = v13
            }
            p2 -= 2147483648
        }
    } else {
        switch 32767 {
        case 4294967295:
            v22 := GenS1{F0: 2, F1: 9223372036854775807}
            _ = v22
        case 12:
            p2 -= 65535
        }
        p0 += 61
    }
    v23 := GenS3{F0: 42, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 uint16) uint32 {
    var v24 float64 = 1.5
    v25 := []int{32767, 32768, 896, 1}
    v26 := GenS1{F0: 1, F1: -1}
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v26
    return 754
}

func main() {
    var v1 int32 = 1
    var v2 bool = false
    v3 := make(map[int]uint32)
    v4 := []int{}
    v5 := []string{"hello", "", (((("\n\t" + (("\n\t" + "") + "")) + "") + "") + ""), ("hello" + "")}
    var v27 uint64 = 42
    var v28 uint64 = (v27 ^ (v27 ^ 18446744073709551615))
    if (-325 <= -787) {
        v2 = true
        switch "hello" {
        case "":
            for fi29 := 0; fi29 < 2; fi29++ {
                v30 := GenS1{F0: 239, F1: 9223372036854775807}
                var v31 uint16 = 32767
                _ = fi29
                _ = v30
                _ = v31
            }
        case "foo bar":
            var v32 float64 = 0.0
            _ = v32
        case "hello":
            fmt.Println(v1, v2, v2)
        }
    } else {
        fmt.Println(v28, v2)
        v28 = v27
    }
    _ = len(v4)
    for fi33 := 0; fi33 < 17; fi33++ {
        v34 := GenS2{F0: 255}
        _ = len(v3)
        v5 = append(v5, "\n\t")
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 string = ("a" + ("hello" + ("\n\t" + "")))
        _ = fi35
        _ = v36
    }
    _ = len(v3)
    for fi37 := 0; fi37 < 20; fi37++ {
        _ = len(v3)
        fmt.Println(v2, fi37, v1)
        _ = fi37
    }
    if ((32767 < 2) || true) {
        var v38 int8 = -128
        _ = v38
    } else {
        if v2 {
            if true {
                v1 = 982
                var v39 int16 = -986
                if v2 {
                    var v40 int64 = 42
                    v41 := GenS4{F0: 65535, F1: (-(v40)), F2: (v28 & (v27 + v28))}
                    for fi42 := 0; fi42 < 17; fi42++ {
                        var v43 uint8 = 1
                        _ = fi42
                        _ = v43
                    }
                    _ = v40
                    _ = v41
                } else {
                    var v44 int32 = v1
                    _ = v44
                }
                _ = v39
            }
            v2 = false
        } else {
            if v2 {
                var v45 bool = v2
                var v46 uint64 = v27
                var v47 uint = 0
                _ = v45
                _ = v46
                _ = v47
            }
            v27 = (^(v28))
        }
    }
    switch -759 {
    case 2147483647:
        var v48 int64 = -718
        _ = v48
    case -1:
        for fi49 := 0; fi49 < 5; fi49++ {
            for fi50 := 0; fi50 < 14; fi50++ {
                v51 := GenS2{F0: 80}
                var v52 uint8 = 2
                _ = fi50
                _ = v51
                _ = v52
            }
            _ = fi49
        }
    }
    for fi53 := 0; fi53 < 15; fi53++ {
        fmt.Println(v28, v28, v28)
        _ = fi53
    }
    if true {
        switch "hello" {
        case "a":
            for fi54 := 0; fi54 < 15; fi54++ {
                switch (fi54 ^ fi54) {
                case 1:
                    v5 = append(v5, ("\n\t" + ("a" + ("\n\t" + ("\n\t" + "")))))
                case 0:
                    var v55 uint8 = 150
                    _ = v55
                default:
                    if (!false) {
                        var v56 string = "a"
                        v57 := GenS1{F0: 255, F1: -1}
                        v56 = v56
                        _ = v56
                        _ = v57
                    } else {
                        for fi58 := 0; fi58 < 1; fi58++ {
                            v5 = append(v5, "a")
                            var v59 int16 = -253
                            _ = fi58
                            _ = v59
                        }
                        v5 = append(v5, "hello")
                        switch "a" {
                        case "":
                            v2 = v2
                        default:
                            if false {
                                v27 = v27
                                fmt.Println(v27, fi54)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                switch (("" + "foo bar") + "foo bar") {
                                case "foo bar":
                                    if (v2 && ((!true) && true)) {
                                        if (25 > 127) {
                                            fmt.Println(v28, fi54)
                                        }
                                        switch -9223372036854775808 {
                                        case 128:
                                            if true {
                                                delete(v3, -2147483648)
                                                fmt.Println(fi54, fi54)
                                            } else {
                                                v2 = (v2 || true)
                                                var v60 uint8 = 1
                                                _ = v60
                                            }
                                        case 127:
                                            var v61 float64 = (0.0 - (-2.25 - (3.14159 / 1e10)))
                                            _ = v61
                                        }
                                        if v2 {
                                            v3[(fi54 + fi54)] = 2
                                            var v62 int32 = (v1 >> ((998) & 63))
                                            _ = v62
                                        }
                                    } else {
                                        _ = len(v3)
                                        fmt.Println(v2)
                                        for fi63 := 0; fi63 < 3; fi63++ {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi63
                                        }
                                    }
                                default:
                                    var v64 bool = true
                                    _ = v64
                                }
                                v65 := v3[(fi54 * fi54)]
                                fmt.Println(fi54, v28, v65)
                                _ = v65
                            }
                        }
                    }
                }
                var v66 int8 = 100
                v67 := v3[(fi54 * (fi54 << ((4294967295) & 63)))]
                _ = fi54
                _ = v66
                _ = v67
            }
        case "\n\t":
            if v2 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v68 int32 = (v1 - (v1 & v1))
                _ = v68
            }
        }
    }
    fmt.Println(v1, v1, v2)
    v69 := GenS5{F0: (v1 ^ v1), F1: v28, F2: (v28 + 983)}
    v70 := GenS6{F0: 42, F1: (v1 + -2147483648)}
    switch 801 {
    case 955:
        var v71 uint8 = 42
        _ = v71
    case -9223372036854775808:
        v1 = v1
    case -1:
        if (454 > genFn1(v27, uint16(v27))) {
            var v72 uint64 = 180
            delete(v3, -9223372036854775808)
            var v73 int32 = v1
            _ = v72
            _ = v73
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
}
