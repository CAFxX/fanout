package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 bool }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 int }

func main() {
    v1 := []int{42, -9223372036854775808, 2147483647}
    v2 := make(map[string]uint32)
    var v3 uint16 = 130
    v4 := make(map[uint32]bool)
    v5 := []uint32{168, 514}
    var v6 int32 = 2147483647
    if true {
        fmt.Println(v3)
    } else {
        switch "\n\t" {
        case "foo bar":
            if (65535 >= 156) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v6 = (v6 * (v6 | (v6 - (v6 + (v6 / ((v6) | 1))))))
                _ = len(v1)
            }
        case "hello":
            if true {
                if false {
                    switch 32768 {
                    case -168:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        var v7 float32 = ((float32(-2.25) - float32(1e10)) + 0.1)
                        _ = v7
                    }
                    if false {
                        v8 := GenS1{F0: true}
                        var v9 int32 = v6
                        fmt.Println(v9, v3, v3)
                        _ = v8
                        _ = v9
                    } else {
                        _ = len(v5)
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (!(-32768 >= 32767)) {
                                var v10 float32 = 3.14159
                                var v11 int = 128
                                if (v11 == (v11 & 32768)) {
                                    if true {
                                        for fi12 := 0; fi12 < 17; fi12++ {
                                            v3 += v3
                                            if false {
                                                for fi13 := 0; fi13 < 11; fi13++ {
                                                    fmt.Println(v3, v11)
                                                    v14 := GenS2{F0: 2}
                                                    v15 := GenS2{F0: 0}
                                                    v16 := GenS1{F0: (false || (!(true || (!false))))}
                                                    _ = fi13
                                                    _ = v14
                                                    _ = v15
                                                    _ = v16
                                                }
                                                var v17 uint32 = 1
                                                switch 127 {
                                                case 32767:
                                                    var v18 int64 = -9223372036854775808
                                                    _ = v18
                                                case 2147483647:
                                                    if true {
                                                        switch 0 {
                                                        case 42:
                                                            v11 = (fi12 >> ((1) & 63))
                                                        case 18:
                                                            var v19 int64 = 876
                                                            _ = v19
                                                        case 1:
                                                            var v20 float32 = v10
                                                            _ = v20
                                                        default:
                                                            v1 = append(v1, (fi12 | (v11 % (((v11 + (v11 >> ((829) & 63)))) | 1))))
                                                        }
                                                        for fi21 := 0; fi21 < 14; fi21++ {
                                                            v10 -= (v10 / v10)
                                                            fmt.Println(fi12, v3)
                                                            _ = fi21
                                                        }
                                                    }
                                                case -283:
                                                    v22 := GenS1{F0: (((1000 != -534) || ((2 == 255) && true)) && true)}
                                                    _ = v22
                                                }
                                                _ = v17
                                            } else {
                                                fmt.Println(v10)
                                                var v23 uint64 = 4294967295
                                                _ = v23
                                            }
                                            _ = fi12
                                        }
                                    }
                                    if false {
                                        if (779 > (v3 & v3)) {
                                            if false {
                                                v3 = v3
                                                v24 := GenS1{F0: (!(!true))}
                                                _ = v24
                                            } else {
                                                _ = len(v2)
                                                v25 := GenS3{F0: "\n\t", F1: 728, F2: ((-919 > (v11 + v11)) && false)}
                                                v26 := GenS4{F0: 2, F1: 255}
                                                _ = v25
                                                _ = v26
                                            }
                                        } else {
                                            for fi27 := 0; fi27 < 13; fi27++ {
                                                v1 = append(v1, (fi27 + (v11 | (^(v11)))))
                                                var v28 uint = 65535
                                                _ = fi27
                                                _ = v28
                                            }
                                            var v29 float64 = 0.1
                                            var v30 bool = true
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                }
                                _ = v10
                                _ = v11
                            } else {
                                switch -9223372036854775808 {
                                case 0:
                                    for fi31 := 0; fi31 < 15; fi31++ {
                                        _ = len(v4)
                                        _ = len(v2)
                                        v32 := GenS2{F0: 2}
                                        _ = fi31
                                        _ = v32
                                    }
                                case 23:
                                    var v33 float64 = (-2.25 + 0.1)
                                    _ = v33
                                case 42:
                                    var v34 uint8 = 0
                                    _ = v34
                                default:
                                    var v35 int8 = -128
                                    _ = v35
                                }
                                v6 = (v6 | -2147483648)
                            }
                            delete(v2, "a")
                        } else {
                            if (42 <= 0) {
                                v3 -= v3
                            }
                            fmt.Println(v3)
                            v3 = (v3 & 763)
                        }
                    }
                } else {
                    var v36 int32 = v6
                    _ = v36
                }
                if (2 < 969) {
                    fmt.Println(v6, v3)
                    v3 -= (v3 | v3)
                    var v37 float64 = 0.0
                    _ = v37
                }
                for fi38 := 0; fi38 < 10; fi38++ {
                    v39 := GenS5{F0: -66}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v6 += 126
                    _ = fi38
                    _ = v39
                }
            } else {
                fmt.Println(v6, v6)
                if (250 != -995) {
                    v6 -= (v6 % ((v6) | 1))
                }
            }
        }
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        v41 := GenS6{F0: (fi40 & -998), F1: 4294967295}
        var v42 int8 = 100
        _ = fi40
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        switch fi43 {
        case 945:
            v3 -= (^(v3))
        case -639:
            fmt.Println(fi43, v6)
        default:
            fmt.Println(v6)
        }
        var v44 string = ("\n\t" + (("" + ("" + (("" + "\n\t") + ("\n\t" + "\n\t")))) + (((("\n\t" + "hello") + "\n\t") + "a") + "a")))
        for fi45 := 0; fi45 < 19; fi45++ {
            v3 = (^(v3))
            v46 := v4[4294967295]
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi45
            _ = v46
        }
        _ = fi43
        _ = v44
    }
    var v47 uint16 = v3
    v48 := GenS4{F0: 2147483648, F1: 643}
    fmt.Println(v47)
    for fi49 := 0; fi49 < 15; fi49++ {
        var v50 uint8 = 6
        _ = fi49
        _ = v50
    }
    if len(v1) > 0 {
        v1[0] = 127
    }
    fmt.Println(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
