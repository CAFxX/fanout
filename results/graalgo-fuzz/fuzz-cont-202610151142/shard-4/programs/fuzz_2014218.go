package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int; F2 int16 }
type GenS7 struct { F0 string; F1 uint16; F2 uint16 }
type GenS8 struct { F0 uint }
func genFn0(p0 uint64) uint64 {
    var v7 int16 = 1000
    v7 += (v7 * v7)
    var v8 string = "hello"
    if (false && (-1 >= (v7 & (v7 + 32767)))) {
        var v9 uint8 = 42
        var v10 bool = (p0 <= p0)
        _ = v10
        _ = v9
    } else {
        fmt.Println(v8, v8, p0)
        var v11 int64 = 0
        fmt.Println(v7)
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return (p0 + (p0 + (p0 / ((759) | 1))))
}
func genFn1() uint8 {
    v12 := []int{-519, -1, 326, 128}
    var v13 uint32 = 2147483647
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int16 = 1
        _ = fi14
        _ = v15
    }
    fmt.Println(v13)
    _ = v12
    _ = v13
    return 255
}

func main() {
    var v1 float64 = (float64(1e10) / float64((0.0 + 1e10)))
    var v2 float64 = v1
    v3 := make(map[string]string)
    var v4 uint64 = 0
    v5 := make(map[int]bool)
    var v6 uint8 = 80
    v16 := GenS1{F0: uint16(v6)}
    switch -2147483648 {
    case -975:
        if false {
            v2 = (v2 * v2)
        } else {
            if false {
                if false {
                    switch "" {
                    case "hello":
                        var v17 uint16 = 2
                        _ = v17
                    case "\n\t":
                        var v18 int8 = 95
                        _ = v18
                    case "foo bar":
                        var v19 int16 = -1
                        _ = v19
                    }
                    v1 = 0.1
                    v20 := GenS1{F0: 1}
                    _ = v20
                }
                fmt.Println(v2, v2, v1)
            } else {
                fmt.Println(v2)
            }
            switch -9223372036854775808 {
            case 32767:
                fmt.Println(v2, v6)
            case 174:
                switch -851 {
                case 773:
                    if (-39 != 1) {
                        if (v4 < v4) {
                            v4 += (v4 | 18446744073709551615)
                            fmt.Println(v2, v2, v4)
                        } else {
                            var v21 bool = ((v4 << ((2) & 63)) == v4)
                            _ = v21
                        }
                        if false {
                            v1 += v1
                            if true {
                                v22 := []int{}
                                v6 -= genFn1()
                                _ = v22
                            }
                        }
                        for fi23 := 0; fi23 < 10; fi23++ {
                            for fi24 := 0; fi24 < 14; fi24++ {
                                if (-297 != 9223372036854775807) {
                                    switch (fi24 & fi24) {
                                    case 2147483647:
                                        v25 := v3["hello"]
                                        _ = v25
                                    case 0:
                                        switch ("a" + "foo bar") {
                                        case "a":
                                            if (false && ((("" + ("\n\t" + "foo bar")) != (("\n\t" + "\n\t") + "hello")) && (true && (1 <= 1)))) {
                                                v26 := GenS2{F0: (745 > 18446744073709551615), F1: 772, F2: 322}
                                                v6 = v6
                                                fmt.Println(fi24)
                                                _ = v26
                                            }
                                            fmt.Println(v1)
                                        case "hello":
                                            var v27 float32 = 1e10
                                            _ = v27
                                        case "foo bar":
                                            if ((-(fi24)) <= 42) {
                                                if true {
                                                    v28 := GenS1{F0: 2}
                                                    switch fi24 {
                                                    case 4294967295:
                                                        v6 = v6
                                                    default:
                                                        v29 := GenS1{F0: 97}
                                                        _ = v29
                                                    }
                                                    var v30 uint8 = v6
                                                    _ = v28
                                                    _ = v30
                                                }
                                                if true {
                                                    var v31 string = "\n\t"
                                                    _ = v31
                                                }
                                                if true {
                                                    v32 := GenS2{F0: false, F1: 2147483648, F2: 609}
                                                    var v33 int = fi24
                                                    _ = v32
                                                    _ = v33
                                                }
                                            }
                                        default:
                                            v2 += v2
                                        }
                                        if ((fi23 + (fi24 / (((fi24 + (fi24 * -259))) | 1))) != 42) {
                                            v34 := GenS3{F0: fi23, F1: ("\n\t" + ("" + "a"))}
                                            _ = v34
                                        } else {
                                            fmt.Println(v6)
                                        }
                                        var v35 uint32 = 254
                                        _ = v35
                                    case 127:
                                        var v36 uint8 = (v6 ^ v6)
                                        _ = v36
                                    default:
                                        if false {
                                            fmt.Println(v1)
                                            fmt.Println(v1, fi23)
                                            var v37 bool = false
                                            _ = v37
                                        }
                                    }
                                }
                                if ((fi23 >> ((18446744073709551615) & 63)) <= (fi23 * fi23)) {
                                    v38 := GenS4{F0: 42, F1: -981}
                                    v1 = (v2 / v2)
                                    var v39 uint8 = v6
                                    _ = v38
                                    _ = v39
                                } else {
                                    _ = len(v5)
                                }
                                v40 := GenS5{F0: (!true)}
                                _ = fi24
                                _ = v40
                            }
                            var v41 int8 = -90
                            for fi42 := 0; fi42 < 13; fi42++ {
                                v43 := []int{2147483647}
                                _ = fi42
                                _ = v43
                            }
                            _ = fi23
                            _ = v41
                        }
                    }
                case 4294967295:
                    fmt.Println(v1, v2)
                }
            }
            var v44 int32 = 42
            _ = v44
        }
    case 672:
        fmt.Println(v4, v2, v4)
    case -667:
        for fi45 := 0; fi45 < 12; fi45++ {
            if true {
                v1 = float64((1e10 / 3.14159))
            }
            _ = fi45
        }
    default:
        v46 := GenS2{F0: ((false || (65535 > 41)) || ((((-90 == -1) && true) && ((-2.25 * float32(v1)) <= float32(v2))) && ((v1 != 3.14159) && true))), F1: 2147483647, F2: 0}
        _ = v46
    }
    if ((0 <= 282) && (32767 < 32767)) {
        if false {
            v4 -= uint64(v6)
            for fi47 := 0; fi47 < 2; fi47++ {
                v3[(("foo bar" + "foo bar") + "\n\t")] = ""
                var v48 int = (^(fi47))
                v49 := v5[(-(v48))]
                _ = fi47
                _ = v48
                _ = v49
            }
            v50 := GenS6{F0: -1, F1: 2147483647, F2: -32768}
            _ = v50
        } else {
            v6 = 143
            v51 := GenS7{F0: "hello", F1: 65535, F2: 1}
            v52 := GenS7{F0: "a", F1: 32767, F2: 0}
            _ = v51
            _ = v52
        }
    }
    fmt.Println(v1, v2)
    fmt.Println(v2, v4)
    v53 := []int{-2147483648, 1, -9223372036854775808, 296}
    v6 -= (^(v6))
    if len(v53) > 0 {
        v53[0] = 128
    }
    var v54 bool = (348 < 198)
    var v55 int16 = 1000
    if (!false) {
        fmt.Println(v54, v6, v1)
    }
    fmt.Println(v2)
    v56 := GenS8{F0: 164}
    var v57 int8 = -80
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
