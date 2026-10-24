package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 string; F1 string }

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]int)
    var v3 uint8 = 9
    v4 := []uint32{4294967295, 1, 806}
    v5 := make(map[uint32]int)
    v3 += (^(v3))
    v6 := GenS1{F0: (v1 + v1), F1: 705}
    if false {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi7, fi7)
            switch ((("\n\t" + ("foo bar" + ("foo bar" + ("\n\t" + "hello")))) + ("a" + "hello")) + "hello") {
            case "":
                var v8 int32 = 1
                _ = v8
            case "\n\t":
                if false {
                    fmt.Println(v1, v1)
                }
            case "hello":
                if (1 >= 555) {
                    _ = len(v4)
                }
            default:
                v3 -= (^(v3))
            }
            _ = fi7
        }
    } else {
        fmt.Println(v1, v1)
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint64 = 0
        if (1 == v1) {
            if ((("a" + "\n\t") <= "a") && true) {
                v11 := GenS2{F0: (v3 | v3), F1: 9223372036854775807}
                var v12 int16 = 0
                _ = v11
                _ = v12
            }
        } else {
            if (!(-1 == 381)) {
                v1 = (v1 * 330)
                if (359 > 0) {
                    v1 = (v10 & (v1 & (v1 % ((v10) | 1))))
                    var v13 float32 = 0.0
                    v3 = (v3 ^ (v3 / (((^(v3))) | 1)))
                    _ = v13
                } else {
                    switch ("foo bar" + (("" + ((("\n\t" + ("a" + "hello")) + "\n\t") + ("" + "hello"))) + "hello")) {
                    case "":
                        if false {
                            _ = len(v4)
                        }
                    default:
                        var v14 uint8 = 233
                        _ = v14
                    }
                }
            }
        }
        _ = fi9
        _ = v10
    }
    if (!(!(884 <= 381))) {
        if true {
            var v15 uint8 = v3
            if false {
                var v16 int8 = 0
                v15 = uint8(v1)
                _ = v16
            } else {
                var v17 int8 = -128
                if true {
                    if true {
                        fmt.Println(v3, v15)
                        v1 -= 42
                    }
                    v15 += v15
                }
                _ = v17
            }
            v1 += (v1 * 561)
            _ = v15
        } else {
            var v18 int32 = -1
            v3 = v3
            if true {
                var v19 uint8 = 42
                v2[("hello" + "\n\t")] = 432
                _ = v19
            }
            _ = v18
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (42 > 1) {
            for fi20 := 0; fi20 < 13; fi20++ {
                v1 = (v1 * v1)
                _ = fi20
            }
            if len(v4) > 0 {
                v4[0] = 2
            }
            fmt.Println(v3)
        }
    }
    v1 = v1
    if (!(v3 == (v3 + (v3 >> ((548) & 63))))) {
        v1 = 899
        for fi21 := 0; fi21 < 13; fi21++ {
            fmt.Println(fi21, fi21, v1)
            v22 := v5[65535]
            fmt.Println(v1, v1, v1)
            _ = fi21
            _ = v22
        }
    } else {
        fmt.Println(v1, v1)
        v23 := GenS1{F0: v1, F1: uint32(v1)}
        _ = v23
    }
    v24 := GenS2{F0: v3, F1: 42}
    if (0 == -1) {
        fmt.Println(v3, v1)
        v25 := GenS1{F0: v1, F1: 2147483648}
        _ = v25
    } else {
        v1 = 0
        v1 = v1
    }
    var v26 float32 = (0.0 + (-2.25 + 1e10))
    v27 := GenS2{F0: v3, F1: 1}
    v3 = (v3 - 127)
    fmt.Println(v1, v26)
    v28 := GenS3{F0: ("\n\t" + "hello"), F1: ((("foo bar" + (("a" + ("a" + "foo bar")) + ("\n\t" + "foo bar"))) + "foo bar") + "")}
    for fi29 := 0; fi29 < 15; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            if true {
                v3 = 128
            }
            var v31 float64 = 0.0
            if len(v4) > 0 {
                v4[0] = 0
            }
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    _ = len(v2)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
