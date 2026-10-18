package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint32; F1 bool; F2 int8 }
type GenS7 struct { F0 string }
func genFn0(p0 uint) int64 {
    v7 := GenS1{F0: -176, F1: 42, F2: 32767}
    switch -108 {
    case -1:
        for fi8 := 0; fi8 < 9; fi8++ {
            if (!(255 == 127)) {
                var v9 int = 128
                p0 = (p0 << ((p0) & 63))
                v10 := GenS1{F0: -586, F1: 254, F2: 0}
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    case -821:
        var v11 uint = 833
        _ = v11
    case 32768:
        var v12 uint8 = 252
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 string = "a"
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(fi15)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := []int{340, 0}
            var v18 int32 = -2147483648
            switch "foo bar" {
            case "hello":
                switch fi16 {
                case -9223372036854775808:
                    fmt.Println(fi16, fi16, p0)
                case 135:
                    v18 -= v18
                default:
                    if false {
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v18, p0)
                        case "foo bar":
                            v17 = append(v17, -2147483648)
                        case "":
                            v18 -= 1
                        }
                        v19 := GenS2{F0: uint8(p0), F1: 2147483647}
                        _ = v19
                    } else {
                        var v20 uint8 = 1
                        for fi21 := 0; fi21 < 3; fi21++ {
                            p0 = (^(p0))
                            _ = fi21
                        }
                        _ = v20
                    }
                }
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        var v22 int64 = 9223372036854775807
        _ = v22
    } else {
        p0 += (p0 ^ p0)
        if false {
            p0 = (p0 % ((799) | 1))
            var v23 uint16 = 1
            _ = v23
        }
        p0 += 4294967295
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    v1 := []int{-195}
    v2 := make(map[uint32]uint32)
    v3 := []string{"foo bar"}
    v4 := []int{42, -2147483648, 1}
    var v5 int32 = -2147483648
    var v6 int8 = 100
    if false {
        fmt.Println(v5, v6)
        if len(v3) > 0 {
            v3[0] = ("a" + "")
        }
        if true {
            var v24 bool = (("foo bar" < "hello") || true)
            v5 = (v5 ^ v5)
            if v24 {
                v25 := GenS3{F0: 349, F1: 105, F2: 507}
                v26 := GenS1{F0: -9223372036854775808, F1: 758, F2: 65535}
                _ = v25
                _ = v26
            }
            _ = v24
        }
    } else {
        fmt.Println(v5)
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v5)
    case "\n\t":
        for fi27 := 0; fi27 < 14; fi27++ {
            v6 += 125
            for fi28 := 0; fi28 < 20; fi28++ {
                _ = len(v2)
                _ = fi28
            }
            if true {
                v29 := GenS3{F0: v5, F1: 97, F2: 255}
                v30 := GenS1{F0: 9223372036854775807, F1: 201, F2: 1}
                v31 := GenS1{F0: -1, F1: 2, F2: 736}
                _ = v29
                _ = v30
                _ = v31
            } else {
                v6 += (v6 ^ -128)
                switch (fi27 % ((662) | 1)) {
                case 728:
                    var v32 float32 = float32(3.14159)
                    _ = v32
                case 723:
                    v6 -= 127
                }
            }
            _ = fi27
        }
    case "":
        v6 -= (v6 >> ((657) & 63))
    }
    v33 := GenS2{F0: 64, F1: 2147483647}
    fmt.Println(v6)
    fmt.Println(v5, v5)
    var v34 float64 = 1e10
    v5 = 608
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 int64 = -9223372036854775808
        for fi37 := 0; fi37 < 10; fi37++ {
            switch "hello" {
            case "a":
                v38 := GenS4{F0: ((626 != (v5 << ((18446744073709551615) & 63))) || false), F1: 4294967295}
                _ = len(v2)
                _ = v38
            case "foo bar":
                if false {
                    switch ("a" + "hello") {
                    case "a":
                        var v39 string = "hello"
                        _ = v39
                    case "":
                        v40 := GenS4{F0: true, F1: 200}
                        _ = v40
                    case "hello":
                        v5 -= v5
                    }
                    var v41 uint = 65535
                    _ = v41
                }
            default:
                if ((v6 * v6) >= v6) {
                    fmt.Println(fi35)
                    v5 = -647
                    if true {
                        var v42 int8 = v6
                        if (false && (float32(v34) != 1e10)) {
                            switch (fi35 >> ((65535) & 63)) {
                            case -1:
                                fmt.Println(v42, fi35, v42)
                            case 32768:
                                _ = len(v3)
                            }
                        } else {
                            fmt.Println(v6, v34)
                        }
                        if true {
                            v43 := GenS2{F0: 86, F1: (-(v5))}
                            v44 := v2[2147483647]
                            var v45 uint8 = 42
                            _ = v43
                            _ = v44
                            _ = v45
                        } else {
                            fmt.Println(fi37)
                            switch (fi37 / (((fi37 * (fi35 >> ((980) & 63)))) | 1)) {
                            case 1:
                                if false {
                                    v46 := GenS5{F0: (fi35 % (((fi37 | 32767)) | 1)), F1: "hello"}
                                    var v47 uint32 = 761
                                    var v48 float32 = ((1.5 + -2.25) * 0.1)
                                    switch (fi35 | (-(fi35))) {
                                    case 32767:
                                        fmt.Println(v6, v47)
                                    case 300:
                                        if (true && false) {
                                            v49 := GenS3{F0: 2147483647, F1: 127, F2: 0}
                                            switch "\n\t" {
                                            case "":
                                                v50 := GenS3{F0: (v5 | (v5 & (v5 >> ((2) & 63)))), F1: uint8(v47), F2: 51}
                                                _ = v50
                                            case "a":
                                                var v51 bool = (!(("a" + "") < ("foo bar" + "\n\t")))
                                                var v52 uint32 = (v47 | (v47 >> ((4294967295) & 63)))
                                                v42 += v42
                                                _ = v51
                                                _ = v52
                                            case "\n\t":
                                                var v53 uint32 = (^(v47))
                                                _ = v53
                                            }
                                            _ = v49
                                        } else {
                                            switch fi35 {
                                            case 2147483647:
                                                var v54 int = (fi37 * (fi37 - (-(fi37))))
                                                _ = v54
                                            default:
                                                v55 := GenS5{F0: 32767, F1: "a"}
                                                _ = v55
                                            }
                                        }
                                    default:
                                        v34 = float64(0.1)
                                    }
                                    var v56 bool = (true && (!false))
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                    _ = v56
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + ("a" + "hello"))
                                    }
                                    _ = len(v2)
                                    _ = len(v3)
                                }
                            default:
                                switch (fi35 - (fi35 % (((fi35 * (-(fi35)))) | 1))) {
                                case 573:
                                    if (true && true) {
                                        if ((0.1 + 1e10) <= 0.0) {
                                            if (false || true) {
                                                v36 = (v36 % ((genFn0(18446744073709551615)) | 1))
                                                var v57 bool = false
                                                _ = v57
                                            } else {
                                                v5 += -956
                                                var v58 int8 = v6
                                                v34 -= v34
                                                var v59 float64 = v34
                                                var v60 bool = (!(!(!((1 <= 1) || false))))
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                            v5 = v5
                                            fmt.Println(fi37)
                                        } else {
                                            fmt.Println(v5, v34, fi35)
                                            var v61 float64 = 1.5
                                            _ = v61
                                        }
                                        var v62 int64 = (v36 >> ((0) & 63))
                                        _ = v62
                                    }
                                case 1:
                                    fmt.Println(v34, v42, v5)
                                case 128:
                                    fmt.Println(v34, v6, fi35)
                                    v5 -= v5
                                default:
                                    v63 := GenS6{F0: 4294967295, F1: (fi35 >= (fi35 ^ fi35)), F2: -128}
                                    _ = v63
                                }
                            }
                            v64 := GenS7{F0: "hello"}
                            _ = v64
                        }
                        _ = v42
                    }
                } else {
                    v65 := GenS5{F0: (fi35 >> ((18446744073709551615) & 63)), F1: "hello"}
                    v34 = (v34 + float64(((0.1 - (1.5 + float32(-2.25))) + ((float32(1e10) - (float32(v34) * 1.5)) / (0.1 * 0.1)))))
                    _ = v65
                }
            }
            v66 := GenS7{F0: "foo bar"}
            _ = fi37
            _ = v66
        }
        if (518 < 65535) {
            var v67 uint16 = 65535
            v2[960] = 2
            v68 := GenS5{F0: -1, F1: "foo bar"}
            _ = v67
            _ = v68
        } else {
            _ = len(v2)
            if ("a" < ("\n\t" + "a")) {
                for fi69 := 0; fi69 < 17; fi69++ {
                    v34 += (-2.25 / v34)
                    _ = fi69
                }
                fmt.Println(fi35, v5, v34)
            } else {
                switch fi35 {
                case 1:
                    _ = len(v2)
                default:
                    fmt.Println(v5, v6)
                }
            }
        }
        _ = fi35
        _ = v36
    }
    v4 = append(v4, -404)
    var v70 int8 = v6
    var v71 bool = (((("\n\t" + "hello") + "a") <= "\n\t") || (65535 < 42))
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v6 -= (v6 >> ((703) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
}
