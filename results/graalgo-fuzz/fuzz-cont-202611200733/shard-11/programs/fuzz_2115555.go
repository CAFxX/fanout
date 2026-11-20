package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int; F1 string; F2 string }
type GenS6 struct { F0 int64; F1 int64 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    v4 := []int{-1, -2147483648, -2147483648}
    var v5 float64 = 0.1
    if false {
        var v6 uint32 = 42
        v5 -= 1.5
        for fi7 := 0; fi7 < 4; fi7++ {
            switch (fi7 >> ((65535) & 63)) {
            case 32767:
                switch fi7 {
                case -261:
                    v8 := GenS1{F0: 286, F1: true, F2: 65535}
                    _ = v8
                case 4294967295:
                    if true {
                        fmt.Println(v5)
                    } else {
                        v5 -= v5
                    }
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case 133:
                var v9 uint32 = 201
                _ = v9
            default:
                if (0.0 == 1.5) {
                    switch (fi7 | fi7) {
                    case 307:
                        var v10 uint = 730
                        _ = v10
                    default:
                        v11 := GenS1{F0: (fi7 + 1), F1: true, F2: uint16(v1)}
                        _ = v11
                    }
                    switch -347 {
                    case -631:
                        var v12 uint64 = v1
                        _ = v12
                    case -9223372036854775808:
                        v6 -= (v6 * 2)
                    case -832:
                        for fi13 := 0; fi13 < 5; fi13++ {
                            v14 := v2[""]
                            v6 = 2147483648
                            v15 := GenS2{F0: 17, F1: -306, F2: 32767}
                            var v16 uint = 42
                            v14 = fi13
                            var v17 uint32 = 0
                            if (32767 < 1) {
                                var v18 uint32 = (v6 - v6)
                                v18 = 42
                                v16 = (v16 + (v16 & 2))
                                var v19 uint32 = (v17 << ((v16) & 63))
                                v5 = 0.0
                                _ = v18
                                _ = v19
                            }
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    default:
                        fmt.Println(fi7, v1, v5)
                    }
                    var v20 int64 = 0
                    _ = v20
                } else {
                    v5 += (v5 - ((v5 / 1.5) + v5))
                }
            }
            var v21 string = ((("" + (("a" + "foo bar") + ((("a" + "") + (("hello" + (((((("" + ("a" + ("foo bar" + ""))) + "a") + "") + "hello") + "a") + "\n\t")) + "hello")) + "foo bar"))) + "foo bar") + ("\n\t" + (("\n\t" + "foo bar") + "")))
            _ = fi7
            _ = v21
        }
        _ = v6
    }
    v1 = (v1 % ((v1) | 1))
    v1 = 18446744073709551615
    v1 -= (v1 % ((v1) | 1))
    v5 = -2.25
    v1 = (^(v1))
    v5 = (v5 - v5)
    v5 -= 0.0
    if (687 < 255) {
        v22 := GenS2{F0: 2, F1: 42, F2: 1}
        for fi23 := 0; fi23 < 17; fi23++ {
            fmt.Println(v5, v5, fi23)
            _ = fi23
        }
        var v24 bool = true
        _ = v22
        _ = v24
    } else {
        v5 = -2.25
    }
    switch 9223372036854775807 {
    case 85:
        var v25 uint64 = (v1 | v1)
        _ = v25
    default:
        if true {
            var v26 uint8 = 241
            var v27 string = (("" + "") + "")
            _ = v26
            _ = v27
        }
    }
    if false {
        for fi28 := 0; fi28 < 19; fi28++ {
            fmt.Println(fi28)
            for fi29 := 0; fi29 < 9; fi29++ {
                v4 = append(v4, fi29)
                v1 = v1
                v5 = (v5 / ((v5 / ((v5 * 0.1) + v5)) / v5))
                v30 := GenS3{F0: v1, F1: 127, F2: 0}
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 12; fi31++ {
                v5 = v5
                _ = fi31
            }
            _ = fi28
        }
    } else {
        v32 := GenS4{F0: int64(v1), F1: 2147483648, F2: 745}
        v33 := GenS5{F0: -1, F1: "foo bar", F2: "foo bar"}
        _ = v32
        _ = v33
    }
    v34 := GenS6{F0: 0, F1: -867}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
