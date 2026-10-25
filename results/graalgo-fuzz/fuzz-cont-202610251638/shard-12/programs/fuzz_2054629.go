package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string; F1 uint64; F2 uint32 }

func main() {
    var v1 int32 = 42
    var v2 float64 = (3.14159 * 3.14159)
    var v3 int32 = -369
    v4 := make(map[string]int)
    v5 := []uint32{}
    var v6 int32 = (v3 - (v1 ^ -130))
    fmt.Println(v3)
    for fi7 := 0; fi7 < 8; fi7++ {
        v3 = v3
        if (127 > 1) {
            _ = len(v4)
        } else {
            v6 -= v6
        }
        if (0 == -128) {
            switch ("hello" + ("a" + ("foo bar" + ("hello" + ("foo bar" + ("hello" + "")))))) {
            case "a":
                switch (-(fi7)) {
                case 177:
                    v8 := v4["\n\t"]
                    _ = v8
                }
            }
        } else {
            switch fi7 {
            case -68:
                var v9 uint = 65535
                _ = v9
            }
        }
        _ = fi7
    }
    var v10 int16 = 0
    v2 -= (1.5 * (((float64(0.0) - 1e10) / (float64(((1.5 / (float32((v2 / v2)) + 3.14159)) * -2.25)) * v2)) / v2))
    v11 := GenS1{F0: (v3 - (-(v6))), F1: -275}
    v4[("a" + "\n\t")] = -9223372036854775808
    v12 := GenS1{F0: (^(v1)), F1: 863}
    v2 -= -2.25
    fmt.Println(v3, v6, v2)
    if (0.1 >= 0.0) {
        var v13 int16 = (v10 + (v10 / (((v10 & (-(v10)))) | 1)))
        _ = v13
    } else {
        fmt.Println(v1, v1)
        v14 := GenS1{F0: 350, F1: -123}
        _ = v14
    }
    v15 := GenS1{F0: (v1 & (^(v6))), F1: 432}
    fmt.Println(v6, v2, v3)
    v3 = -156
    v16 := GenS1{F0: 2147483647, F1: 32767}
    for fi17 := 0; fi17 < 16; fi17++ {
        if (((0 > 857) || (2 >= 0)) && (1 < 144)) {
            switch fi17 {
            case -1:
                switch ("\n\t" + ("hello" + (("\n\t" + "foo bar") + ("foo bar" + ("" + "foo bar"))))) {
                case "\n\t":
                    if ((((0.1 * float32(((3.14159 + 1e10) + v2))) + 1e10) - -2.25) != 3.14159) {
                        fmt.Println(v6, v3)
                    } else {
                        var v18 uint = 2
                        fmt.Println(v2)
                        v2 -= float64(0.0)
                        _ = v18
                    }
                case "foo bar":
                    var v19 int32 = v6
                    _ = v19
                case "hello":
                    v20 := GenS1{F0: v3, F1: (fi17 | fi17)}
                    _ = v20
                default:
                    var v21 uint64 = 1
                    _ = v21
                }
            case 880:
                if (!(-1 > 42)) {
                    var v22 float32 = ((3.14159 * 0.0) - 1e10)
                    var v23 string = "\n\t"
                    _ = v22
                    _ = v23
                } else {
                    if ((128 >= 12) || (true || (false || (fi17 < (fi17 & (fi17 << ((65535) & 63))))))) {
                        v24 := GenS1{F0: (v3 % ((v1) | 1)), F1: 72}
                        v25 := v4["hello"]
                        if ("" >= "\n\t") {
                            fmt.Println(v10, v25, v10)
                            for fi26 := 0; fi26 < 20; fi26++ {
                                v10 = (v10 - -222)
                                v3 += (v1 % (((v6 ^ 627)) | 1))
                                v27 := GenS1{F0: v3, F1: (fi17 << ((0) & 63))}
                                var v28 float64 = 3.14159
                                var v29 int64 = -1
                                _ = fi26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            if false {
                                fmt.Println(v3, v3)
                                var v30 int8 = -1
                                _ = v30
                            }
                        }
                        _ = v24
                        _ = v25
                    }
                }
            default:
                v4[((("a" + ((("\n\t" + ("foo bar" + "hello")) + "") + (("hello" + "") + ("a" + (("hello" + ("foo bar" + "a")) + (("foo bar" + "a") + ("\n\t" + ((("foo bar" + ("\n\t" + "hello")) + "") + "a")))))))) + "") + "hello")] = fi17
            }
        } else {
            if (!true) {
                var v31 int16 = 42
                var v32 uint8 = 0
                _ = v31
                _ = v32
            }
            fmt.Println(v3, v6, fi17)
        }
        _ = fi17
    }
    if ((-2.25 + float64(float32(v2))) <= 3.14159) {
        if (829 >= 2) {
            var v33 int32 = (v6 % (((v3 >> ((558) & 63))) | 1))
            v34 := GenS2{F0: "\n\t", F1: 42, F2: 506}
            v10 = (v10 << ((1) & 63))
            _ = v33
            _ = v34
        }
    } else {
        v35 := v4["hello"]
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v36 := GenS2{F0: "\n\t", F1: 18446744073709551615, F2: 0}
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
