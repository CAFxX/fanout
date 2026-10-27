package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }

func main() {
    v1 := []string{"\n\t", "a", "a"}
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    v4 := []int{}
    v5 := []string{"\n\t", (("" + "foo bar") + "foo bar"), ("hello" + (((("hello" + "a") + "") + ("" + ("foo bar" + (("foo bar" + "") + "foo bar")))) + "a"))}
    var v6 int8 = 20
    var v7 float32 = float32(-2.25)
    v6 = (v6 & v6)
    v7 = 1e10
    var v8 int = 805
    switch (v8 * (^(v8))) {
    case 2147483647:
        var v9 int16 = -32768
        _ = v9
    case 128:
        var v10 int = (v8 + (v8 << ((18446744073709551615) & 63)))
        _ = v10
    case 1:
        if true {
            if (!true) {
                v11 := GenS1{F0: (v8 >> ((18446744073709551615) & 63)), F1: 2147483647, F2: false}
                for fi12 := 0; fi12 < 17; fi12++ {
                    v13 := GenS1{F0: 954, F1: -387, F2: (!(-1 >= -174))}
                    switch (v8 << ((1) & 63)) {
                    case -307:
                        switch "" {
                        case "hello":
                            v3["hello"] = "hello"
                        case "":
                            for fi14 := 0; fi14 < 10; fi14++ {
                                var v15 int64 = -903
                                _ = fi14
                                _ = v15
                            }
                        case "a":
                            v7 += (v7 * -2.25)
                        default:
                            switch (fi12 % ((v8) | 1)) {
                            case 127:
                                v8 += (v8 % (((v8 - 32768)) | 1))
                            case 9223372036854775807:
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            case 32767:
                                var v16 float32 = 0.0
                                _ = v16
                            default:
                                switch "a" {
                                case "a":
                                    v8 += 2147483647
                                case "hello":
                                    fmt.Println(v8, v7)
                                case "\n\t":
                                    for fi17 := 0; fi17 < 13; fi17++ {
                                        var v18 int = 652
                                        _ = fi17
                                        _ = v18
                                    }
                                }
                            }
                        }
                    default:
                        fmt.Println(fi12)
                    }
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
            var v19 int = v8
            _ = v19
        } else {
            v20 := GenS1{F0: (v8 | (v8 % ((v8) | 1))), F1: int32(v8), F2: true}
            if false {
                for fi21 := 0; fi21 < 11; fi21++ {
                    fmt.Println(v6, v6, fi21)
                    var v22 uint8 = 234
                    _ = fi21
                    _ = v22
                }
                var v23 uint32 = 987
                _ = v23
            }
            if (0.1 == 0.1) {
                var v24 uint64 = 0
                v8 -= v8
                _ = v24
            } else {
                var v25 int64 = 42
                fmt.Println(v25)
                v8 = v8
                _ = v25
            }
            _ = v20
        }
    default:
        for fi26 := 0; fi26 < 14; fi26++ {
            v7 = v7
            for fi27 := 0; fi27 < 16; fi27++ {
                v8 = v8
                fmt.Println(fi26, fi27)
                v28 := GenS1{F0: fi26, F1: -777, F2: true}
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    v6 += (v6 * -1)
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
