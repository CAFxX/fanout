package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 bool; F1 int16; F2 int32 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint64; F1 int32; F2 int32 }
func genFn0(p0 uint8, p1 int16) int64 {
    v5 := GenS1{F0: p1, F1: (!false), F2: (p1 - 0)}
    switch ("a" + "a") {
    case "foo bar":
        if (65535 != 42) {
            v6 := []int{-263, 2147483647, 307, 881}
            _ = v6
        }
    }
    v7 := []int{32767}
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return -706
}
func genFn1(p0 uint8, p1 bool) int16 {
    v8 := GenS1{F0: 1, F1: (42 == 350), F2: -1}
    p0 -= p0
    if ((true || true) && false) {
        v9 := []int{128, -2147483648, -940, 127}
        p0 -= p0
        var v10 uint8 = p0
        _ = v10
        _ = v9
    } else {
        var v11 int = -471
        var v12 uint32 = 591
        p0 = p0
        _ = v11
        _ = v12
    }
    v13 := []int{688, -2147483648}
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    return 782
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 0.0
    var v3 int32 = 0
    v4 := []string{"hello"}
    _ = len(v1)
    var v14 float32 = 3.14159
    _ = len(v4)
    var v15 float32 = v14
    var v16 int = 748
    v1[2147483647] = "foo bar"
    fmt.Println(v14, v16)
    var v17 int16 = -1
    v4 = append(v4, "\n\t")
    if false {
        if (v3 < v3) {
            for fi18 := 0; fi18 < 9; fi18++ {
                if true {
                    fmt.Println(v3, v2, v15)
                    _ = len(v1)
                    v14 += (-2.25 / 1.5)
                } else {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 int16 = 32767
                        v21 := GenS2{F0: v17, F1: 992}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                v4 = append(v4, "foo bar")
                if false {
                    var v22 uint = 4294967295
                    if (!(true && (!true))) {
                        v23 := GenS3{F0: 65535, F1: 128}
                        _ = v23
                    } else {
                        var v24 int = 678
                        _ = v24
                    }
                    _ = v22
                } else {
                    v15 = 1e10
                }
                _ = fi18
            }
        } else {
            for fi25 := 0; fi25 < 4; fi25++ {
                for fi26 := 0; fi26 < 13; fi26++ {
                    v27 := GenS4{F0: true, F1: -32768, F2: (^(v3))}
                    var v28 string = (("hello" + ("" + "")) + "hello")
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v16 -= (fi25 ^ (fi25 << ((18446744073709551615) & 63)))
                _ = fi25
            }
            v29 := GenS5{F0: (v3 | 1), F1: (^(v16))}
            if (!false) {
                if true {
                    v17 = -32768
                    switch (v16 | (v16 >> ((0) & 63))) {
                    case -467:
                        if true {
                            if true {
                                for fi30 := 0; fi30 < 19; fi30++ {
                                    var v31 float64 = v2
                                    v32 := GenS4{F0: ("a" <= ""), F1: v17, F2: 42}
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                            }
                            var v33 uint = 497
                            _ = v33
                        } else {
                            v16 -= (v16 / ((v16) | 1))
                        }
                    case 663:
                        v3 += (^(v3))
                    case -71:
                        if true {
                            switch v16 {
                            case 4294967295:
                                fmt.Println(v17)
                            case 2147483647:
                                var v34 uint32 = 42
                                _ = v34
                            case 184:
                                v35 := GenS4{F0: ((v3 % ((1) | 1)) <= (v3 % ((42) | 1))), F1: v17, F2: (v3 * v3)}
                                _ = v35
                            }
                            switch "foo bar" {
                            case "a":
                                var v36 int = 0
                                _ = v36
                            case "\n\t":
                                if (true && true) {
                                    var v37 float64 = v2
                                    v17 -= genFn1(2, (v2 != (-2.25 + v37)))
                                    _ = v37
                                } else {
                                    v14 += 0.1
                                    v16 -= v16
                                    var v38 string = "\n\t"
                                    _ = v38
                                }
                            default:
                                switch (v16 + 755) {
                                case 1:
                                    switch (v16 / (((v16 * 9223372036854775807)) | 1)) {
                                    case 32767:
                                        var v39 float32 = (v14 * v15)
                                        _ = v39
                                    case -595:
                                        v17 = (v17 ^ 599)
                                    case 128:
                                        for fi40 := 0; fi40 < 3; fi40++ {
                                            switch "a" {
                                            case "a":
                                                v41 := GenS1{F0: (v17 & -859), F1: (!((!(true || true)) || (32767 != 255))), F2: v17}
                                                _ = v41
                                            }
                                            v3 += v3
                                            _ = len(v4)
                                            _ = fi40
                                        }
                                    }
                                case 127:
                                    v2 = 1.5
                                case -678:
                                    v15 = ((v14 + -2.25) - v14)
                                default:
                                    if (215 < 42) {
                                        v2 += float64(0.0)
                                        v42 := GenS2{F0: v17, F1: 2}
                                        var v43 uint8 = 62
                                        _ = v42
                                        _ = v43
                                    } else {
                                        if true {
                                            v44 := GenS5{F0: v3, F1: 32768}
                                            _ = v44
                                        }
                                        if (!false) {
                                            var v45 uint16 = 0
                                            _ = v45
                                        } else {
                                            fmt.Println(v16, v17, v15)
                                            v16 += -499
                                        }
                                        _ = len(v1)
                                    }
                                }
                            }
                            delete(v1, 303)
                        } else {
                            v46 := GenS6{F0: 2}
                            var v47 bool = false
                            _ = v46
                            _ = v47
                        }
                    default:
                        v15 += v14
                    }
                } else {
                    if (false || true) {
                        v48 := GenS1{F0: genFn1(182, (!(v2 > v2))), F1: true, F2: -32768}
                        _ = v48
                    } else {
                        v4 = append(v4, ("a" + ""))
                        for fi49 := 0; fi49 < 4; fi49++ {
                            v17 += 973
                            fmt.Println(v14, v17, v3)
                            _ = fi49
                        }
                        switch (v16 / ((v16) | 1)) {
                        case 1:
                            v50 := GenS7{F0: (869 < 955), F1: 42}
                            _ = v50
                        case -1:
                            var v51 uint8 = 160
                            _ = v51
                        case -238:
                            for fi52 := 0; fi52 < 7; fi52++ {
                                v53 := GenS8{F0: 319, F1: -2147483648, F2: v3}
                                v2 = v2
                                var v54 float32 = v14
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                        default:
                            var v55 int64 = genFn0(0, (v17 >> ((0) & 63)))
                            _ = v55
                        }
                    }
                    var v56 bool = (((2 == 1) && true) || false)
                    for fi57 := 0; fi57 < 10; fi57++ {
                        var v58 uint16 = 0
                        _ = fi57
                        _ = v58
                    }
                    _ = v56
                }
                v17 += (v17 ^ v17)
            } else {
                switch v16 {
                case -1:
                    if (!(!(1 == 254))) {
                        var v59 uint16 = 32767
                        _ = v59
                    }
                case -362:
                    for fi60 := 0; fi60 < 18; fi60++ {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        v61 := GenS2{F0: (v17 | genFn1(0, false)), F1: 46}
                        if (0 != 42) {
                            v2 += 1.5
                        }
                        _ = fi60
                        _ = v61
                    }
                }
                for fi62 := 0; fi62 < 7; fi62++ {
                    fmt.Println(v14, v17, v14)
                    var v63 int = (fi62 - (^(v16)))
                    for fi64 := 0; fi64 < 11; fi64++ {
                        v17 = v17
                        v65 := GenS3{F0: 511, F1: 128}
                        v66 := GenS3{F0: 0, F1: 184}
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    _ = fi62
                    _ = v63
                }
                v67 := GenS6{F0: 2}
                _ = v67
            }
            _ = v29
        }
        for fi68 := 0; fi68 < 4; fi68++ {
            v3 = (-(v3))
            _ = fi68
        }
    }
    v1[2] = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
