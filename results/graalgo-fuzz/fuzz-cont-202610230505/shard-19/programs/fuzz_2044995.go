package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 bool; F2 string }
type GenS4 struct { F0 int16; F1 uint64; F2 string }

func main() {
    v1 := []string{("a" + "foo bar"), "foo bar"}
    v2 := make(map[string]uint32)
    var v3 bool = (42 > 232)
    var v4 int = 455
    v5 := make(map[string]bool)
    v5["hello"] = (true && (((v4 << ((1) & 63)) == (v4 ^ v4)) && false))
    if (!v3) {
        if ((!((121 != 2) && (true && v3))) || v3) {
            for fi6 := 0; fi6 < 8; fi6++ {
                if true {
                    v7 := GenS1{F0: 255}
                    _ = len(v5)
                    _ = v7
                } else {
                    switch (fi6 % (((v4 | (^(fi6)))) | 1)) {
                    case -632:
                        fmt.Println(fi6)
                    case 4294967295:
                        if v3 {
                            var v8 float32 = 0.1
                            for fi9 := 0; fi9 < 18; fi9++ {
                                fmt.Println(v8, fi9, v4)
                                var v10 int16 = 823
                                v3 = false
                                _ = fi9
                                _ = v10
                            }
                            v3 = v3
                            _ = v8
                        }
                    default:
                        v11 := GenS1{F0: 128}
                        _ = v11
                    }
                    v12 := GenS2{F0: 0}
                    if (!false) {
                        var v13 int16 = -32768
                        fmt.Println(v3, v4, v4)
                        v4 += (v4 % ((fi6) | 1))
                        _ = v13
                    }
                    _ = v12
                }
                if (0 != -127) {
                    v3 = v3
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    switch fi6 {
                    case 127:
                        if (-114 > 580) {
                            delete(v2, ("hello" + ("\n\t" + ("hello" + "foo bar"))))
                            fmt.Println(fi6)
                            if v3 {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v14 float32 = (-2.25 + float32(1e10))
                                _ = v14
                            } else {
                                v4 = -2147483648
                            }
                        }
                    case 128:
                        var v15 float32 = (3.14159 - (-2.25 / 1e10))
                        _ = v15
                    case 197:
                        v16 := GenS1{F0: 48}
                        _ = v16
                    default:
                        var v17 int64 = -9223372036854775808
                        _ = v17
                    }
                } else {
                    _ = len(v1)
                    switch (-(fi6)) {
                    case 4294967295:
                        v4 = 952
                    case 9223372036854775807:
                        v3 = ((25 > 0) || ((-9223372036854775808 > -1) && false))
                    case 0:
                        var v18 uint = 261
                        _ = v18
                    default:
                        var v19 bool = false
                        _ = v19
                    }
                    var v20 uint8 = 146
                    _ = v20
                }
                v21 := v5[((("foo bar" + "") + "\n\t") + (("hello" + ("foo bar" + ("" + "a"))) + ("" + "hello")))]
                _ = fi6
                _ = v21
            }
            v22 := GenS3{F0: ("" + "a"), F1: v3, F2: "\n\t"}
            v3 = (v3 || v3)
            _ = v22
        }
    } else {
        _ = len(v1)
        fmt.Println(v3, v3)
    }
    v3 = false
    var v23 bool = (!v3)
    var v24 int64 = -1
    fmt.Println(v23, v3, v24)
    v25 := GenS4{F0: 645, F1: 18446744073709551615, F2: ""}
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 uint8 = 231
        if true {
            v4 = (fi26 & v4)
            v3 = (681 < 2)
            v28 := GenS2{F0: 42}
            _ = v28
        }
        _ = fi26
        _ = v27
    }
    var v29 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
