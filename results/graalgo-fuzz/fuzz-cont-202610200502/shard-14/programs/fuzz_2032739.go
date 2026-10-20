package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int16; F2 string }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }
func genFn0(p0 uint16, p1 string) int16 {
    p0 = 65535
    switch -142 {
    case 2147483647:
        var v4 uint64 = 1
        _ = v4
    }
    _ = p0
    _ = p1
    return -32768
}

func main() {
    var v1 float64 = float64(-2.25)
    v2 := make(map[int]string)
    var v3 int = -59
    switch (v3 << ((59) & 63)) {
    case 32768:
        v5 := GenS1{F0: (v3 + 4294967295)}
        _ = v5
    case 1:
        fmt.Println(v1, v3)
    case 4294967295:
        delete(v2, v3)
    }
    var v6 float64 = (v1 - v1)
    if ((!(42 != 0)) || (65535 == 2147483648)) {
        for fi7 := 0; fi7 < 19; fi7++ {
            if true {
                var v8 float64 = v6
                v6 = -2.25
                _ = v8
            } else {
                for fi9 := 0; fi9 < 6; fi9++ {
                    if (((((0.0 * (1.5 * 1.5)) + 0.0) * 0.1) / -2.25) >= float32(v1)) {
                        var v10 bool = ((v3 % ((fi9) | 1)) > (fi9 >> ((18446744073709551615) & 63)))
                        _ = v10
                    }
                    v1 -= v6
                    _ = fi9
                }
            }
            _ = fi7
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            v1 -= v6
            _ = fi11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 bool = (65535 < 255)
            v14 := GenS1{F0: v3}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v2[127] = ""
    v15 := []int{v3, (v3 % (((v3 * -828)) | 1))}
    var v16 int = v3
    if (!true) {
        v17 := GenS1{F0: v3}
        v18 := GenS1{F0: (v3 / ((v16) | 1))}
        _ = v17
        _ = v18
    } else {
        v16 -= (v3 & (v16 ^ v16))
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                var v20 uint8 = 128
                for fi21 := 0; fi21 < 3; fi21++ {
                    v22 := GenS2{F0: v3}
                    v20 = (v20 ^ v20)
                    _ = len(v15)
                    var v23 string = ("foo bar" + "a")
                    v20 -= 42
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v24 := GenS3{F0: 42, F1: genFn0(1, (("\n\t" + "") + "\n\t")), F2: "foo bar"}
                _ = v20
                _ = v24
            }
            _ = fi19
        }
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(v3, v1, v3)
        _ = fi25
    }
    switch (v3 + (v3 & -9223372036854775808)) {
    case -639:
        v16 = v3
    case 128:
        if (213 > 1) {
            if true {
                switch "" {
                case "hello":
                    v1 += v6
                case "foo bar":
                    fmt.Println(v16, v3, v3)
                case "":
                    var v26 float64 = v6
                    _ = v26
                }
                switch v3 {
                case -9223372036854775808:
                    for fi27 := 0; fi27 < 18; fi27++ {
                        fmt.Println(v3, v3, v6)
                        v28 := GenS4{F0: 42, F1: -521}
                        _ = fi27
                        _ = v28
                    }
                case 4294967295:
                    v29 := GenS5{F0: -445}
                    _ = v29
                case -289:
                    v30 := GenS6{F0: (("a" + "hello") + (("\n\t" + ("hello" + "foo bar")) + "foo bar")), F1: 883}
                    _ = v30
                default:
                    if ((2 > 442) || ((4294967295 > 1) || (-69 > -128))) {
                        v1 -= v6
                    } else {
                        v6 = v6
                        for fi31 := 0; fi31 < 6; fi31++ {
                            for fi32 := 0; fi32 < 7; fi32++ {
                                var v33 int16 = 0
                                v34 := GenS6{F0: "\n\t", F1: 0}
                                v33 += genFn0(42, "\n\t")
                                if (!true) {
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    v35 := GenS3{F0: 212, F1: v33, F2: "\n\t"}
                                    if true {
                                        v36 := GenS7{F0: 0}
                                        _ = v36
                                    } else {
                                        v3 -= -9223372036854775808
                                        if (0.1 == 3.14159) {
                                            fmt.Println(v3)
                                            switch "\n\t" {
                                            case "\n\t":
                                                if len(v15) >= 2 {
                                                    v15 = v15[1:len(v15)]
                                                }
                                            default:
                                                fmt.Println(v33, fi31, v33)
                                            }
                                            v37 := GenS8{F0: 1}
                                            v33 += (v33 | (v33 - v33))
                                            if ((!(42 <= 643)) || false) {
                                                v1 -= v1
                                                fmt.Println(v6)
                                            }
                                            _ = v37
                                        } else {
                                            var v38 int32 = -1
                                            _ = v38
                                        }
                                    }
                                    _ = v35
                                }
                                fmt.Println(v6, v33, fi31)
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                            _ = fi31
                        }
                        _ = len(v15)
                    }
                }
                v39 := GenS5{F0: -293}
                _ = v39
            }
            delete(v2, v16)
            if true {
                v3 -= (v16 << ((18446744073709551615) & 63))
                if ((!(true || (false || false))) || (!false)) {
                    switch v3 {
                    case 29:
                        delete(v2, 9223372036854775807)
                    default:
                        var v40 int = v16
                        _ = v40
                    }
                    var v41 string = "foo bar"
                    v1 = (1.5 / (float64((0.0 * -2.25)) - 1e10))
                    _ = v41
                } else {
                    if (42 < v3) {
                        for fi42 := 0; fi42 < 8; fi42++ {
                            switch ("\n\t" + "") {
                            case "":
                                _ = len(v15)
                            default:
                                var v43 int16 = -647
                                _ = v43
                            }
                            v3 += fi42
                            delete(v2, fi42)
                            _ = fi42
                        }
                        fmt.Println(v3, v6)
                        var v44 float64 = (1.5 + (v6 - v1))
                        _ = v44
                    } else {
                        var v45 uint = 42
                        _ = v45
                    }
                    if true {
                        _ = len(v15)
                        fmt.Println(v3)
                        v15 = append(v15, (v3 ^ (-(v16))))
                    }
                }
            } else {
                var v46 float64 = (3.14159 + v6)
                v3 += (v16 ^ v16)
                _ = v46
            }
        } else {
            v16 = (v16 / (((v3 >> ((1) & 63))) | 1))
            var v47 int = (v16 << ((18446744073709551615) & 63))
            _ = v47
        }
    case -9223372036854775808:
        v15 = append(v15, (v3 + (v3 | 9223372036854775807)))
    }
    fmt.Println(v1)
    var v48 uint32 = 314
    v48 += v48
    v49 := v2[-926]
    var v50 uint64 = 2
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    switch (v49 + ("hello" + v49)) {
    case "":
        v51 := GenS2{F0: 579}
        _ = v51
    default:
        fmt.Println(v49, v49, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v6
}
