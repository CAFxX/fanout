package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -93
    v3 := []int{-9223372036854775808, -1, -49}
    v4 := make(map[uint32]uint32)
    var v5 int = 710
    var v6 int = -1
    v2 += (v2 & v2)
    v7 := GenS1{F0: "a", F1: -2147483648, F2: 0}
    switch ("foo bar" + "a") {
    case "a":
        _ = len(v4)
    default:
        if (342 != 42) {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 int64 = 0
                _ = fi8
                _ = v9
            }
        } else {
            fmt.Println(v5)
            var v10 int32 = 110
            _ = v10
        }
    }
    if false {
        fmt.Println(v5, v5, v2)
    } else {
        switch "a" {
        case "foo bar":
            if true {
                var v11 string = ("\n\t" + "\n\t")
                if true {
                    var v12 uint64 = 4294967295
                    if ((!(0.1 > 1e10)) || true) {
                        v6 = (v6 + v5)
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v14 := GenS2{F0: 42, F1: 2}
                        switch v6 {
                        case -340:
                            var v15 bool = true
                            _ = v15
                        default:
                            if (42 < 1) {
                                if true {
                                    for fi16 := 0; fi16 < 2; fi16++ {
                                        if true {
                                            if false {
                                                v17 := GenS3{F0: 540}
                                                v18 := GenS1{F0: v11, F1: 20, F2: 797}
                                                switch -942 {
                                                case -9223372036854775808:
                                                    if (true || false) {
                                                        v5 = (v6 % ((fi16) | 1))
                                                        switch 128 {
                                                        case -783:
                                                            var v19 int16 = 32767
                                                            _ = v19
                                                        case 128:
                                                            v11 = v11
                                                        case -2147483648:
                                                            if false {
                                                                v2 += v2
                                                            } else {
                                                                var v20 string = ("\n\t" + (((v11 + v11) + v11) + (("foo bar" + ((v11 + "") + v11)) + ("a" + ""))))
                                                                _ = v20
                                                            }
                                                        }
                                                    } else {
                                                        if false {
                                                            v21 := GenS4{F0: (v11 + (v11 + v11))}
                                                            _ = v21
                                                        }
                                                        v22 := v4[522]
                                                        _ = v22
                                                    }
                                                case 128:
                                                    switch v11 {
                                                    case "hello":
                                                        var v23 uint16 = 848
                                                        _ = v23
                                                    case "":
                                                        var v24 int8 = v2
                                                        _ = v24
                                                    }
                                                default:
                                                    if (false && (42 < 99)) {
                                                        var v25 string = ""
                                                        var v26 int16 = 42
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                }
                                                _ = v17
                                                _ = v18
                                            } else {
                                                v5 = (fi13 + 32768)
                                                v6 = v6
                                            }
                                            var v27 uint32 = 430
                                            _ = v27
                                        }
                                        if (473 > 0) {
                                            fmt.Println(v5)
                                            if (fi13 < (fi16 << ((0) & 63))) {
                                                fmt.Println(fi13)
                                                var v28 uint64 = 508
                                                _ = v28
                                            }
                                            v2 = v2
                                        }
                                        _ = fi16
                                    }
                                }
                                for fi29 := 0; fi29 < 17; fi29++ {
                                    v5 = (^(fi13))
                                    v5 += fi13
                                    fmt.Println(fi13, v11)
                                    _ = fi29
                                }
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                }
                fmt.Println(v6, v11, v6)
                _ = v11
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    v5 += -9223372036854775808
                    for fi31 := 0; fi31 < 13; fi31++ {
                        var v32 uint16 = 2
                        var v33 bool = ("foo bar" != ((("\n\t" + ("a" + (("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) + "foo bar"))) + "a") + "a"))
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 1; fi34++ {
                        v35 := GenS5{F0: 42, F1: 9223372036854775807}
                        if true {
                            v1[801] = 2147483648
                            v6 = (-(v5))
                        }
                        var v36 int = 128
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    _ = fi30
                }
            }
        }
    }
    v6 += v5
    fmt.Println(v2)
    fmt.Println(v2, v6, v2)
    switch v6 {
    case 32767:
        v2 = 42
    case 54:
        v37 := GenS1{F0: "\n\t", F1: 999, F2: 2147483648}
        _ = v37
    }
    v5 = (v5 % (((v5 << ((64) & 63))) | 1))
    v6 = 994
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
