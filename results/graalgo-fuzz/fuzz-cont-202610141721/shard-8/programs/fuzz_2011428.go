package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint; F2 uint64 }
type GenS8 struct { F0 string; F1 uint16; F2 uint8 }

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 0
    var v3 bool = (false && true)
    v4 := []int{}
    v4 = append(v4, 389)
    var v5 uint32 = 0
    v3 = v3
    fmt.Println(v2, v1)
    v6 := GenS1{F0: 86, F1: v3}
    if ((v3 || (-1 == 0)) && (!false)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch "\n\t" {
        case "a":
            v7 := GenS2{F0: false, F1: 550}
            _ = v7
        case "foo bar":
            v5 += 1
        default:
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(v2, v3, v1)
                _ = fi8
            }
        }
        switch -517 {
        case 0:
            switch "a" {
            case "foo bar":
                v9 := GenS3{F0: 1}
                _ = v9
            case "a":
                if v3 {
                    var v10 float64 = float64(-2.25)
                    _ = v10
                } else {
                    var v11 int32 = 545
                    _ = v11
                }
            default:
                v1 = ((((1e10 / float32((0.0 + (1.5 * 0.1)))) - ((v1 + (float32((3.14159 * 1e10)) + 3.14159)) - 0.1)) - 1.5) * (((v1 - v1) * (float32(1.5) + v1)) / (float32(0.1) / (float32(3.14159) - v1))))
            }
        case 42:
            fmt.Println(v3)
        case -9223372036854775808:
            var v12 float64 = 0.0
            _ = v12
        default:
            if v3 {
                var v13 bool = (0 != 1)
                _ = v13
            } else {
                var v14 float32 = v1
                if v3 {
                    var v15 uint8 = 127
                    v16 := GenS4{F0: -32768, F1: v2, F2: 4294967295}
                    fmt.Println(v5, v2, v5)
                    _ = v15
                    _ = v16
                }
                v5 += 2147483648
                _ = v14
            }
        }
    } else {
        switch -409 {
        case -9223372036854775808:
            var v17 uint16 = 1
            _ = v17
        case 322:
            fmt.Println(v2, v1)
        case 127:
            v5 = v5
        default:
            v3 = (v3 || true)
        }
    }
    var v18 int8 = 1
    var v19 int64 = 9223372036854775807
    var v20 uint16 = (v2 ^ v2)
    v21 := GenS2{F0: ((v18 | (-(v18))) == 127), F1: 18446744073709551615}
    if true {
        v2 = 0
        fmt.Println(v1, v18, v5)
    }
    v22 := GenS5{F0: (237 != 0)}
    if (0.0 < (float64((v1 / float32(0.1))) * 0.0)) {
        var v23 uint16 = (^(v20))
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 bool = true
            _ = fi24
            _ = v25
        }
        _ = len(v4)
        _ = v23
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            if (-9223372036854775808 >= -9223372036854775808) {
                if v3 {
                    if v3 {
                        v28 := GenS4{F0: 0, F1: v20, F2: 18446744073709551615}
                        _ = len(v4)
                        if true {
                            if v3 {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v29 uint = 90
                                var v30 int = (fi27 - (fi27 + (fi26 * (fi27 + fi27))))
                                _ = v29
                                _ = v30
                            } else {
                                v31 := GenS5{F0: v3}
                                var v32 int16 = -1
                                var v33 uint32 = (v5 & uint32(v2))
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            if (false && false) {
                                v34 := GenS4{F0: 0, F1: 1, F2: 18446744073709551615}
                                _ = v34
                            }
                        } else {
                            v35 := GenS1{F0: -76, F1: true}
                            switch ((("" + "\n\t") + (("foo bar" + "foo bar") + "hello")) + "\n\t") {
                            case "":
                                if false {
                                    v4 = append(v4, 0)
                                }
                            }
                            v36 := GenS6{F0: 67}
                            _ = v35
                            _ = v36
                        }
                        fmt.Println(v3, v3)
                        v37 := GenS7{F0: (1 == (v5 << ((18446744073709551615) & 63))), F1: uint(v5), F2: 832}
                        _ = v28
                        _ = v37
                    }
                    var v38 uint64 = 42
                    _ = v38
                } else {
                    var v39 float32 = v1
                    if len(v4) > 0 {
                        v4[0] = fi26
                    }
                    if (584 == 2147483647) {
                        fmt.Println(v3, fi27)
                        var v40 int8 = (^(v18))
                        _ = v40
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            v20 += v20
                        case "hello":
                            var v41 uint64 = 0
                            _ = v41
                        default:
                            fmt.Println(v20, v20)
                        }
                        v19 = v19
                    }
                    _ = v39
                }
                v42 := GenS7{F0: true, F1: 18446744073709551615, F2: 18446744073709551615}
                _ = v42
            }
            var v43 float64 = 1.5
            if (true || v3) {
                v44 := GenS8{F0: "a", F1: v2, F2: 1}
                _ = v44
            } else {
                if len(v4) > 0 {
                    v4[0] = -693
                }
            }
            v45 := GenS5{F0: (v3 && true)}
            _ = fi27
            _ = v43
            _ = v45
        }
        _ = fi26
    }
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    v1 = ((float32(3.14159) + v1) * v1)
    v2 = (v2 ^ 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
