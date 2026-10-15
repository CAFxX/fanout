package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint }
type GenS4 struct { F0 bool; F1 string; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 uint8; F1 uint; F2 int32 }

func main() {
    var v1 uint = 786
    v2 := []uint32{692}
    v3 := make(map[uint32]int)
    v4 := []uint32{4294967295, 826, 854, 1}
    var v5 float32 = 0.1
    v3[2147483648] = 212
    v6 := v3[2]
    if (float64(v5) != ((((-2.25 * float64((0.0 / (v5 - v5)))) - 1e10) - ((0.1 / float64((v5 - v5))) * ((0.1 - 3.14159) * 0.1))) + 0.1)) {
        v7 := GenS1{F0: (("hello" + ("" + "")) + "\n\t"), F1: ""}
        v5 = 0.1
        fmt.Println(v6, v6, v5)
        _ = v7
    } else {
        fmt.Println(v5)
        v1 += (^(v1))
        v8 := GenS1{F0: (("\n\t" + "hello") + "hello"), F1: "foo bar"}
        _ = v8
    }
    fmt.Println(v5, v1)
    switch (v6 ^ -2147483648) {
    case 42:
        v6 -= (v6 / (((^(v6))) | 1))
    case 786:
        _ = len(v3)
    case 1:
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(fi9, v1)
            if (126 < 1) {
                v6 = -796
            } else {
                v10 := GenS2{F0: -699, F1: 1}
                for fi11 := 0; fi11 < 11; fi11++ {
                    v6 -= 2147483647
                    if true {
                        v12 := GenS3{F0: 2, F1: 4294967295, F2: v1}
                        _ = v12
                    } else {
                        var v13 float32 = ((1.5 - v5) * -2.25)
                        var v14 uint64 = 2
                        _ = v13
                        _ = v14
                    }
                    switch (fi9 * 1) {
                    case -251:
                        v15 := GenS1{F0: ((("" + "") + "hello") + ("" + ((("\n\t" + "a") + "hello") + ""))), F1: ""}
                        _ = v15
                    case -2147483648:
                        if true {
                            if true {
                                fmt.Println(v1, fi11)
                                if (-902 >= 63) {
                                    var v16 float64 = 1e10
                                    if true {
                                        if true {
                                            v17 := GenS1{F0: "a", F1: ("a" + ("\n\t" + ("\n\t" + (("hello" + "\n\t") + "foo bar"))))}
                                            _ = v17
                                        }
                                        _ = len(v4)
                                        fmt.Println(v16, fi9)
                                    }
                                    _ = v16
                                }
                                switch (("foo bar" + ("a" + "\n\t")) + (("\n\t" + "") + "foo bar")) {
                                case "":
                                    v1 -= (v1 / ((v1) | 1))
                                    v18 := GenS3{F0: 2, F1: 614, F2: (v1 * v1)}
                                    _ = v18
                                }
                                v19 := GenS4{F0: (202 <= 127), F1: ("a" + "\n\t"), F2: 0}
                                _ = v19
                            }
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                            v1 = v1
                            if (-128 < -75) {
                                fmt.Println(v1, fi11)
                            }
                            switch ((("" + "\n\t") + "a") + ("a" + ("hello" + (("foo bar" + "hello") + ("foo bar" + "hello"))))) {
                            case "hello":
                                if false {
                                    if true {
                                        switch fi11 {
                                        case 1:
                                            fmt.Println(v1, v1, v6)
                                        case -1:
                                            if (2 != 18446744073709551615) {
                                                var v20 int = (fi9 | 128)
                                                _ = v20
                                            } else {
                                                switch -2147483648 {
                                                case 127:
                                                    var v21 int16 = -1
                                                    _ = v21
                                                default:
                                                    fmt.Println(v5, v6, v5)
                                                }
                                                v22 := GenS5{F0: "a"}
                                                var v23 int16 = 1000
                                                _ = v22
                                                _ = v23
                                            }
                                        case 0:
                                            fmt.Println(fi11)
                                        }
                                        var v24 int32 = 2147483647
                                        fmt.Println(v1, v6)
                                        _ = v24
                                    } else {
                                        v25 := GenS6{F0: (!(true && true)), F1: 2}
                                        _ = len(v3)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v25
                                    }
                                }
                            case "":
                                v1 = 383
                            case "a":
                                var v26 int8 = 100
                                _ = v26
                            }
                        } else {
                            v5 = v5
                            var v27 int8 = 127
                            if (!(892 <= 517)) {
                                v5 -= (((0.0 + 0.1) / v5) - v5)
                            }
                            _ = v27
                        }
                    case 42:
                        switch "hello" {
                        case "":
                            var v28 float64 = 0.0
                            _ = v28
                        case "a":
                            v1 = (v1 % (((v1 ^ v1)) | 1))
                        }
                        var v29 int8 = 42
                        fmt.Println(v5, fi11, fi11)
                        _ = v29
                    }
                    _ = fi11
                }
                _ = v10
            }
            _ = fi9
        }
    default:
        v30 := GenS3{F0: 2, F1: 810, F2: 566}
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        if (42 <= 2147483648) {
            var v32 int8 = 98
            v33 := GenS3{F0: 363, F1: 18446744073709551615, F2: (v1 + 256)}
            _ = v32
            _ = v33
        }
        _ = fi31
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        switch v6 {
        case 32767:
            v35 := GenS7{F0: 256, F1: 0}
            _ = v35
        default:
            v36 := GenS8{F0: 0, F1: (^(v1)), F2: -1}
            _ = v36
        }
        var v37 int16 = 0
        _ = fi34
        _ = v37
    }
    delete(v3, 2)
    v38 := GenS6{F0: (!true), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
