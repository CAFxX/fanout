package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int16; F1 int64; F2 uint32 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int64; F2 int16 }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 int32; F1 bool; F2 int32 }
type GenS11 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]int)
    var v3 string = "a"
    var v4 int16 = 42
    fmt.Println(v3, v1, v4)
    var v5 int = 32768
    v6 := GenS1{F0: "hello"}
    if (true || (!false)) {
        switch v3 {
        case "\n\t":
            var v7 int64 = -1
            _ = v7
        default:
            v8 := []int{v5, (^(v5)), int(v1)}
            _ = v8
        }
        if true {
            v2[42] = v5
            switch v3 {
            case "":
                var v9 int64 = 60
                _ = v9
            case "foo bar":
                v10 := GenS2{F0: (v4 % ((v4) | 1))}
                _ = v10
            }
            for fi11 := 0; fi11 < 8; fi11++ {
                v1 = v1
                for fi12 := 0; fi12 < 2; fi12++ {
                    v13 := GenS3{F0: v5, F1: 0}
                    v14 := v2[42]
                    v15 := GenS4{F0: 1, F1: 42}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi11
            }
        } else {
            for fi16 := 0; fi16 < 20; fi16++ {
                fmt.Println(v5, v1)
                var v17 int = fi16
                _ = fi16
                _ = v17
            }
            if (-965 < -995) {
                if (-505 >= v4) {
                    v5 += v5
                    var v18 int8 = 114
                    _ = v18
                } else {
                    v19 := GenS5{F0: (v4 / ((-444) | 1)), F1: 9223372036854775807, F2: 95}
                    v1 = (v1 * 65535)
                    _ = v19
                }
                v20 := GenS5{F0: 0, F1: -704, F2: 65535}
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 uint32 = 2147483648
                    if true {
                        if ("" > v3) {
                            var v23 uint64 = 187
                            v24 := GenS6{F0: 2, F1: true}
                            v25 := []int{fi21, 4294967295, (^(fi21))}
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        if true {
                            switch v3 {
                            case "hello":
                                if (-2147483648 < 329) {
                                    for fi26 := 0; fi26 < 16; fi26++ {
                                        if (-8 >= 71) {
                                            var v27 string = v3
                                            _ = v27
                                        }
                                        _ = fi26
                                    }
                                    v22 += v22
                                    fmt.Println(v22, v1, v4)
                                } else {
                                    v28 := GenS7{F0: 100}
                                    var v29 string = ("a" + ("" + v3))
                                    fmt.Println(v3)
                                    _ = v28
                                    _ = v29
                                }
                            case "foo bar":
                                for fi30 := 0; fi30 < 18; fi30++ {
                                    v31 := GenS5{F0: (v4 * 42), F1: -652, F2: v22}
                                    if true {
                                        fmt.Println(fi21, v4, fi21)
                                        var v32 uint = 65535
                                        v33 := []int{(fi30 % ((-778) | 1)), 128, (fi21 ^ (fi21 - 0)), (-(fi21))}
                                        var v34 float32 = (3.14159 - 0.1)
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    } else {
                                        var v35 uint64 = 944
                                        var v36 float64 = 1.5
                                        v37 := []int{fi30, 4294967295, (fi30 >> ((4294967295) & 63)), (^(fi30))}
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS7{F0: 8}
                                    if (v22 == 569) {
                                        switch fi21 {
                                        case -1:
                                            v39 := GenS8{F0: 71, F1: -513, F2: 32767}
                                            _ = v39
                                        case 1:
                                            var v40 uint = 465
                                            _ = v40
                                        }
                                        v3 = "a"
                                        v22 = (v22 * (v22 + v22))
                                    }
                                    _ = fi30
                                    _ = v31
                                    _ = v38
                                }
                            default:
                                var v41 int16 = 0
                                _ = v41
                            }
                            v5 += int(v1)
                        }
                    }
                    v1 = (v1 << ((42) & 63))
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                fmt.Println(v1)
                v3 = ("\n\t" + (v3 + v3))
                delete(v2, 323)
            }
        }
        if false {
            for fi42 := 0; fi42 < 20; fi42++ {
                v43 := []int{fi42, v5, 0}
                v5 -= v5
                fmt.Println(v4, v1, v4)
                _ = fi42
                _ = v43
            }
        } else {
            v5 = v5
            v44 := GenS9{F0: ((!false) || true), F1: (v1 / ((42) | 1))}
            _ = v44
        }
    } else {
        var v45 int32 = -714
        _ = v45
    }
    if (((!false) && true) || true) {
        switch (v5 % ((v5) | 1)) {
        case -9223372036854775808:
            if (v4 <= v4) {
                v46 := GenS2{F0: (v4 >> ((1) & 63))}
                _ = v46
            } else {
                if (((-2.25 - (3.14159 - -2.25)) / (3.14159 * (float32((3.14159 / 1e10)) + -2.25))) < 1e10) {
                    v4 = (v4 + v4)
                    v1 = v1
                    if (v3 >= (v3 + v3)) {
                        for fi47 := 0; fi47 < 8; fi47++ {
                            for fi48 := 0; fi48 < 10; fi48++ {
                                v49 := []int{fi47}
                                if (false || true) {
                                    var v50 string = (("a" + ("a" + ((((v3 + (v3 + "\n\t")) + v3) + "") + v3))) + "\n\t")
                                    v51 := GenS10{F0: -1, F1: (!false), F2: 2147483647}
                                    _ = v50
                                    _ = v51
                                } else {
                                    v52 := v2[42]
                                    fmt.Println(fi48, v5)
                                    _ = v52
                                }
                                _ = fi48
                                _ = v49
                            }
                            v53 := GenS1{F0: v3}
                            v54 := []int{(v5 >> ((584) & 63)), (v5 & fi47), fi47, fi47}
                            _ = fi47
                            _ = v53
                            _ = v54
                        }
                    } else {
                        if (27 == 42) {
                            if false {
                                switch (v5 % (((v5 / (((v5 >> ((65535) & 63))) | 1))) | 1)) {
                                case -435:
                                    v4 = (^(v4))
                                default:
                                    v1 -= 1
                                }
                            } else {
                                var v55 uint8 = 42
                                v5 += v5
                                _ = v55
                            }
                            v56 := GenS10{F0: -1, F1: true, F2: 1}
                            _ = v56
                        } else {
                            switch int(v4) {
                            case 32767:
                                var v57 uint64 = 2
                                _ = v57
                            case 32768:
                                if (!(1e10 != (float64((3.14159 + float32((0.1 * 0.1)))) / (float64(1.5) - (0.1 / (0.1 * (((1e10 * (float64(0.0) + (((1e10 * -2.25) - ((0.1 + 3.14159) - (float64((((1e10 - 1.5) / float32(3.14159)) + (1e10 / ((1.5 / (1e10 / 1e10)) - (0.1 - (1e10 * 0.0)))))) + (((1.5 - -2.25) * ((1e10 / ((0.1 - 3.14159) * 1e10)) - (0.0 + (0.0 + ((1.5 / 1.0) * 1.5))))) * 0.1)))) * (3.14159 - 1.5)))) / 3.14159) * float64(-2.25)))))))) {
                                    switch v3 {
                                    case "":
                                        fmt.Println(v4)
                                    }
                                    if (true || true) {
                                        var v58 string = v3
                                        var v59 int = (v5 << ((119) & 63))
                                        _ = v58
                                        _ = v59
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            for fi60 := 0; fi60 < 14; fi60++ {
                                                fmt.Println(v4, v5)
                                                for fi61 := 0; fi61 < 9; fi61++ {
                                                    var v62 uint16 = v1
                                                    var v63 int32 = -1
                                                    v2[1] = (^(fi61))
                                                    _ = fi61
                                                    _ = v62
                                                    _ = v63
                                                }
                                                v64 := []int{(-(fi60)), 930, v5, (v5 + -9223372036854775808)}
                                                _ = fi60
                                                _ = v64
                                            }
                                        case "a":
                                            if (128 >= 2) {
                                                v2[29] = (v5 * v5)
                                                fmt.Println(v1)
                                                if false {
                                                    v65 := GenS11{F0: 738, F1: 18446744073709551615}
                                                    fmt.Println(v3, v1, v5)
                                                    _ = v65
                                                } else {
                                                    if (42 > (v1 ^ (v1 << ((18446744073709551615) & 63)))) {
                                                        v66 := GenS7{F0: 61}
                                                        v2[0] = (v5 - v5)
                                                        _ = v66
                                                    } else {
                                                        v1 = (^(v1))
                                                    }
                                                }
                                            } else {
                                                v3 = ("hello" + (v3 + v3))
                                                var v67 int16 = v4
                                                _ = v67
                                            }
                                        case "hello":
                                            v68 := GenS3{F0: (v5 / ((644) | 1)), F1: 683}
                                            _ = v68
                                        }
                                        v2[632] = 127
                                    }
                                }
                            case 127:
                                for fi69 := 0; fi69 < 12; fi69++ {
                                    v1 += v1
                                    v3 = ("" + v3)
                                    v3 = (v3 + v3)
                                    _ = fi69
                                }
                            }
                            v1 += (v1 & 222)
                            v5 = v5
                        }
                        v3 = (v3 + v3)
                    }
                } else {
                    var v70 string = v3
                    v71 := GenS6{F0: 202, F1: false}
                    _ = v70
                    _ = v71
                }
            }
        case 32768:
            var v72 uint = 2
            _ = v72
        }
    }
    fmt.Println(v1, v3, v4)
    for fi73 := 0; fi73 < 15; fi73++ {
        v74 := []int{fi73, v5, (v5 / ((v5) | 1)), v5}
        _ = fi73
        _ = v74
    }
    fmt.Println(v3)
    v75 := GenS9{F0: true, F1: (v1 | v1)}
    var v76 int64 = 326
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
}
