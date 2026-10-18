package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 uint8; F1 int64; F2 bool }
type GenS6 struct { F0 bool; F1 string; F2 uint64 }
func genFn0() uint16 {
    var v7 uint8 = 0
    if false {
        fmt.Println(v7)
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := []int{9223372036854775807}
            v7 += (v7 - v7)
            _ = fi8
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            fmt.Println(v7)
            _ = fi10
        }
        v7 += (v7 * v7)
    }
    _ = v7
    return 304
}
func genFn1() int8 {
    var v11 uint32 = 0
    fmt.Println(v11)
    var v12 float32 = (0.1 + (0.0 / (1e10 * 1e10)))
    v13 := []int{-563}
    _ = v11
    _ = v12
    _ = v13
    return 127
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{32767}
    v3 := make(map[string]bool)
    v4 := []int{-62, -2147483648, 127}
    v5 := []string{"foo bar"}
    v6 := make(map[int]bool)
    var v14 int8 = -66
    switch (("foo bar" + ("hello" + (("\n\t" + "\n\t") + ""))) + "hello") {
    case "hello":
        for fi15 := 0; fi15 < 2; fi15++ {
            fmt.Println(v14, fi15)
            for fi16 := 0; fi16 < 15; fi16++ {
                if ((false || (255 > 2)) && (42 >= -1)) {
                    v14 = v14
                    _ = len(v2)
                }
                if false {
                    v3["a"] = true
                    if len(v2) > 0 {
                        v2[0] = 584
                    }
                    v14 -= genFn1()
                } else {
                    if (9223372036854775807 == fi16) {
                        v17 := GenS1{F0: 1}
                        var v18 float64 = 0.1
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(fi15)
                }
                if len(v2) > 0 {
                    v2[0] = (fi15 * fi16)
                }
                _ = fi16
            }
            v14 = 101
            _ = fi15
        }
    case "foo bar":
        if (613 != 0) {
            var v19 string = ((((("" + ("\n\t" + "a")) + (("foo bar" + "hello") + "hello")) + "a") + "a") + "")
            if (!(((1 <= 2) || false) || false)) {
                v20 := GenS1{F0: 1}
                var v21 uint64 = 0
                v14 = v14
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 11; fi22++ {
                    if len(v2) > 0 {
                        v2[0] = (^(fi22))
                    }
                    var v23 int8 = v14
                    v19 = "hello"
                    _ = fi22
                    _ = v23
                }
                delete(v1, 1)
            }
            _ = v19
        }
    }
    var v24 uint = 255
    v25 := GenS2{F0: 255, F1: 42, F2: true}
    if true {
        fmt.Println(v14, v24)
    }
    switch 42 {
    case 9223372036854775807:
        var v26 uint32 = 556
        _ = v26
    case -248:
        fmt.Println(v14)
    default:
        for fi27 := 0; fi27 < 12; fi27++ {
            if (!false) {
                switch ((("foo bar" + "\n\t") + "\n\t") + ((("foo bar" + "foo bar") + ("a" + ("" + ""))) + "")) {
                case "foo bar":
                    v28 := GenS3{F0: "foo bar", F1: 233}
                    _ = v28
                case "a":
                    if (!(-2.25 > 3.14159)) {
                        v29 := GenS3{F0: "hello", F1: 882}
                        var v30 uint = v24
                        for fi31 := 0; fi31 < 9; fi31++ {
                            if (992 >= 4294967295) {
                                v14 = (v14 + (v14 / ((v14) | 1)))
                                var v32 int64 = -1
                                _ = v32
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    if false {
                                        v33 := GenS4{F0: 4294967295, F1: ("foo bar" + "")}
                                        var v34 uint8 = 127
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v2 = append(v2, (fi31 | 482))
                                        v14 = (-(v14))
                                        v35 := GenS4{F0: 2147483648, F1: "\n\t"}
                                        var v36 float64 = float64(float32(0.0))
                                        _ = v35
                                        _ = v36
                                    }
                                default:
                                    var v37 uint = v24
                                    _ = v37
                                }
                                switch "foo bar" {
                                case "hello":
                                    var v38 string = "\n\t"
                                    _ = v38
                                case "a":
                                    var v39 string = "\n\t"
                                    _ = v39
                                case "\n\t":
                                    v40 := GenS2{F0: 42, F1: 0, F2: (4294967295 == 332)}
                                    _ = v40
                                default:
                                    v14 = v14
                                }
                            }
                            var v41 int = -969
                            var v42 uint = v24
                            _ = fi31
                            _ = v41
                            _ = v42
                        }
                        _ = v29
                        _ = v30
                    } else {
                        if true {
                            fmt.Println(v14, v24)
                            if (100 <= 33) {
                                v3[("a" + "foo bar")] = false
                            } else {
                                var v43 float32 = 1e10
                                _ = v43
                            }
                        }
                        for fi44 := 0; fi44 < 8; fi44++ {
                            if false {
                                v24 = 0
                                var v45 bool = false
                                fmt.Println(v24, fi27)
                                _ = v45
                            } else {
                                v14 = (v14 ^ int8(v24))
                            }
                            var v46 int64 = 0
                            _ = fi44
                            _ = v46
                        }
                    }
                default:
                    v47 := v1[(fi27 * fi27)]
                    _ = v47
                }
                var v48 uint16 = 1
                _ = v48
            }
            var v49 float32 = (0.1 / 1.0)
            for fi50 := 0; fi50 < 18; fi50++ {
                if ("\n\t" > ("foo bar" + (("\n\t" + ("foo bar" + "a")) + "hello"))) {
                    switch "" {
                    case "":
                        if true {
                            v6[(fi27 << ((v24) & 63))] = true
                            v6[(fi27 % ((fi50) | 1))] = true
                        }
                        v14 = -66
                        switch 799 {
                        case 0:
                            var v51 int64 = 9223372036854775807
                            _ = v51
                        }
                    default:
                        v49 += v49
                    }
                    v52 := GenS5{F0: 67, F1: -9223372036854775808, F2: true}
                    _ = v52
                }
                if (!(2147483648 != 570)) {
                    v14 = (v14 >> (((v24 | 18446744073709551615)) & 63))
                    var v53 int64 = 9223372036854775807
                    _ = v53
                }
                _ = fi50
            }
            _ = fi27
            _ = v49
        }
    }
    v14 += (v14 / (((v14 | v14)) | 1))
    fmt.Println(v14, v24)
    v14 = (v14 >> (((v24 / ((v24) | 1))) & 63))
    v54 := GenS6{F0: false, F1: ("hello" + ""), F2: 654}
    var v55 int8 = 89
    v24 += (v24 % (((v24 ^ (^(v24)))) | 1))
    if (!false) {
        v56 := GenS2{F0: 2, F1: 143, F2: true}
        v24 = (v24 / (((v24 % (((v24 + 18446744073709551615)) | 1))) | 1))
        _ = v56
    }
    v57 := v6[679]
    fmt.Println(v24, v14)
    v3[""] = (542 >= 2)
    for fi58 := 0; fi58 < 20; fi58++ {
        if (!v57) {
            var v59 int16 = 1
            _ = v59
        } else {
            v57 = v57
        }
        _ = fi58
    }
    for fi60 := 0; fi60 < 11; fi60++ {
        v57 = v57
        if len(v4) > 0 {
            v4[0] = fi60
        }
        _ = fi60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v57
    _ = v6
}
