package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int32; F1 bool; F2 bool }
type GenS5 struct { F0 int; F1 uint8 }

func main() {
    v1 := []string{("foo bar" + "a"), "foo bar"}
    var v2 int64 = 0
    var v3 uint32 = 42
    v4 := make(map[string]string)
    v5 := []uint32{}
    fmt.Println(v3, v3)
    v1 = append(v1, "")
    var v6 int = -9223372036854775808
    var v7 int64 = 823
    v6 = v6
    var v8 uint16 = 255
    var v9 uint64 = 182
    v4["\n\t"] = ("" + ("a" + ""))
    v6 += v6
    v7 += -9223372036854775808
    if (false || false) {
        v10 := GenS1{F0: 242, F1: (v3 | v3)}
        _ = v10
    }
    var v11 uint8 = 146
    v8 -= 1
    if (v6 > -2147483648) {
        switch "a" {
        case "\n\t":
            var v12 int = int(v8)
            _ = v12
        case "a":
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 int = v6
                _ = fi13
                _ = v14
            }
        case "foo bar":
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 uint8 = 1
                _ = fi15
                _ = v16
            }
        }
        v6 = (v6 % ((0) | 1))
    } else {
        if (v9 < (v9 - v9)) {
            var v17 int16 = 32767
            fmt.Println(v2, v3, v3)
            v17 -= -543
            _ = v17
        }
    }
    if false {
        if true {
            v18 := GenS2{F0: v3, F1: 1}
            var v19 uint16 = (^(v8))
            v20 := v4[((("\n\t" + "foo bar") + ("a" + ("\n\t" + "a"))) + "a")]
            _ = v18
            _ = v19
            _ = v20
        }
        v21 := GenS2{F0: 2, F1: 4294967295}
        _ = v21
    } else {
        v5 = append(v5, (^(v3)))
        for fi22 := 0; fi22 < 4; fi22++ {
            if (42 <= v3) {
                fmt.Println(v11)
            } else {
                var v23 uint8 = v11
                _ = v23
            }
            _ = fi22
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if false {
        if true {
            switch "\n\t" {
            case "a":
                fmt.Println(v7, v6, v11)
            case "hello":
                if false {
                    v8 += (v8 & v8)
                }
            case "foo bar":
                var v24 int32 = -2147483648
                _ = v24
            default:
                v6 = (v6 / ((v6) | 1))
            }
        } else {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v2, v11)
            case "\n\t":
                if (!false) {
                    fmt.Println(v11, v8, v3)
                    var v25 int8 = 0
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 float64 = -2.25
                        var v28 int16 = 1
                        var v29 uint16 = (v8 / ((v8) | 1))
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = v25
                }
            case "hello":
                for fi30 := 0; fi30 < 13; fi30++ {
                    if (int8(v11) >= 1) {
                        v3 += (v3 & 1)
                    }
                    if ("hello" == ("foo bar" + "hello")) {
                        v31 := GenS2{F0: (v3 << ((18446744073709551615) & 63)), F1: 351}
                        for fi32 := 0; fi32 < 6; fi32++ {
                            if false {
                                var v33 uint64 = v9
                                if true {
                                    var v34 uint64 = 4294967295
                                    v35 := v4["\n\t"]
                                    _ = v34
                                    _ = v35
                                } else {
                                    var v36 int64 = v2
                                    var v37 float64 = 1.5
                                    fmt.Println(fi30, v37)
                                    _ = v36
                                    _ = v37
                                }
                                v38 := GenS3{F0: "foo bar", F1: 140}
                                _ = v33
                                _ = v38
                            } else {
                                v9 = (v9 >> ((18446744073709551615) & 63))
                                if (v7 <= -995) {
                                    v3 = v3
                                    var v39 uint64 = v9
                                    if (-1 != (v2 | 9223372036854775807)) {
                                        v40 := GenS2{F0: v3, F1: 18446744073709551615}
                                        v41 := GenS4{F0: 42, F1: (false || (!((v8 | v8) == uint16(v3)))), F2: ((v8 ^ (v8 / ((255) | 1))) <= 42)}
                                        v2 += 1
                                        _ = v40
                                        _ = v41
                                    } else {
                                        var v42 uint8 = v11
                                        v43 := GenS5{F0: 4294967295, F1: v11}
                                        var v44 uint8 = v42
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v39
                                }
                                var v45 int32 = 803
                                switch "\n\t" {
                                case "foo bar":
                                    v46 := GenS4{F0: (v45 ^ v45), F1: (0.0 >= -2.25), F2: false}
                                    _ = v46
                                case "a":
                                    var v47 int16 = 478
                                    _ = v47
                                default:
                                    fmt.Println(fi30)
                                }
                                _ = v45
                            }
                            _ = fi32
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v31
                    }
                    v9 += (^(v9))
                    _ = fi30
                }
            default:
                v48 := v4["a"]
                _ = v48
            }
            var v49 uint32 = v3
            _ = v49
        }
        var v50 int8 = -1
        _ = v50
    } else {
        v9 += (v9 % ((1) | 1))
        v2 -= 42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
