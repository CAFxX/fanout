package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 string; F1 uint8; F2 int32 }
type GenS5 struct { F0 bool }

func main() {
    v1 := []uint32{121, 1}
    var v2 int16 = -567
    v3 := make(map[string]bool)
    var v4 uint16 = 42
    v5 := make(map[string]bool)
    v6 := []string{"a", ((("a" + ((("" + "") + "hello") + "foo bar")) + ("a" + (("\n\t" + ("\n\t" + (("foo bar" + "") + ("\n\t" + ("" + ""))))) + "a"))) + "a"), "\n\t"}
    switch -718 {
    case 127:
        for fi7 := 0; fi7 < 13; fi7++ {
            v2 -= (v2 | v2)
            if (255 < 255) {
                var v8 int32 = 42
                _ = v8
            } else {
                v9 := v3[("a" + (("a" + "hello") + "foo bar"))]
                _ = v9
            }
            if (-1 < 1) {
                v6 = append(v6, ("\n\t" + "a"))
                switch ((("" + ("" + "foo bar")) + "\n\t") + "hello") {
                case "hello":
                    v10 := GenS1{F0: v2, F1: 1, F2: 2147483648}
                    _ = v10
                case "":
                    if len(v1) > 0 {
                        v1[0] = 788
                    }
                }
            }
            _ = fi7
        }
    case 975:
        delete(v5, "")
    }
    _ = len(v5)
    v11 := GenS1{F0: 42, F1: -1, F2: 65535}
    var v12 float32 = 0.0
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    if (2147483647 != 851) {
        _ = len(v5)
        if (((v12 + v12) == (float32(0.0) / v12)) && true) {
            var v13 uint = uint(v4)
            fmt.Println(v12, v2, v4)
            _ = v13
        }
        for fi14 := 0; fi14 < 13; fi14++ {
            if (!(9223372036854775807 <= 42)) {
                var v15 bool = (18446744073709551615 < 274)
                v15 = (0.1 < 3.14159)
                switch 1 {
                case 4294967295:
                    var v16 uint32 = 1
                    _ = v16
                case -1:
                    v2 = -347
                case 274:
                    fmt.Println(v12, v12)
                default:
                    switch (fi14 & 9223372036854775807) {
                    case 127:
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v15 = (-9223372036854775808 <= 1)
                            if v15 {
                                var v18 int64 = 9223372036854775807
                                var v19 bool = (!v15)
                                switch ("\n\t" + "hello") {
                                case "foo bar":
                                    v18 = (v18 * 659)
                                case "\n\t":
                                    switch (fi17 ^ fi17) {
                                    case 1:
                                        v6 = append(v6, "foo bar")
                                    case 127:
                                        v15 = (4294967295 != 2147483647)
                                    }
                                default:
                                    v20 := GenS2{F0: 0, F1: 259}
                                    _ = v20
                                }
                                _ = v18
                                _ = v19
                            } else {
                                var v21 string = "hello"
                                _ = v21
                            }
                            _ = fi17
                        }
                    case 756:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 42:
                        if v15 {
                            fmt.Println(v4)
                        } else {
                            v4 = 585
                            v22 := GenS3{F0: fi14, F1: "\n\t"}
                            _ = v22
                        }
                    }
                }
                _ = v15
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                if true {
                    if true {
                        if (true && (-1 > -352)) {
                            var v24 uint16 = v4
                            v25 := v3[""]
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v24
                            _ = v25
                        }
                    } else {
                        v12 = 3.14159
                        var v26 int32 = 1
                        _ = v26
                    }
                }
                var v27 uint = uint(v4)
                fmt.Println(v12, v27)
                _ = fi23
                _ = v27
            }
            v28 := GenS4{F0: ("foo bar" + "hello"), F1: 43, F2: -895}
            _ = fi14
            _ = v28
        }
    }
    v29 := GenS5{F0: (!true)}
    v12 = ((-2.25 + (v12 + float32(0.1))) * (((3.14159 / 3.14159) - ((v12 + v12) - 0.0)) + ((0.0 - 0.1) - 1.5)))
    for fi30 := 0; fi30 < 9; fi30++ {
        v6 = append(v6, "\n\t")
        fmt.Println(v4)
        _ = fi30
    }
    var v31 uint = 4294967295
    v32 := GenS5{F0: false}
    _ = len(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
