package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int }

func main() {
    var v1 bool = ("hello" > "")
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 1
    v4 := make(map[string]bool)
    var v5 uint32 = 330
    v6 := []string{"", "foo bar"}
    if (!true) {
        v3 = uint16(v5)
        var v7 float64 = 1e10
        fmt.Println(v3, v5, v7)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            v5 = v5
            v2 = (v2 & (v2 | v2))
            _ = fi8
        }
        var v9 string = "hello"
        _ = v9
    }
    var v10 int64 = 0
    var v11 float64 = 0.0
    if (true || true) {
        v12 := GenS1{F0: v10}
        v5 -= v5
        switch "foo bar" {
        case "hello":
            var v13 int16 = 1000
            _ = v13
        case "foo bar":
            if ((^(v3)) != v3) {
                var v14 int16 = 1000
                switch ("" + "hello") {
                case "":
                    var v15 int32 = -1
                    _ = v15
                default:
                    v2 = v2
                }
                _ = v14
            }
        default:
            if true {
                _ = len(v6)
                v16 := GenS2{F0: 2147483647}
                switch ("" + "hello") {
                case "":
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 uint32 = (v5 + (v5 & (v5 + (v5 * 2147483647))))
                        var v19 uint64 = 2
                        switch "\n\t" {
                        case "":
                            switch (fi17 & fi17) {
                            case 2147483647:
                                v20 := GenS1{F0: (v10 / ((int64(v18)) | 1))}
                                _ = v20
                            default:
                                fmt.Println(v5, v3, v3)
                            }
                        case "hello":
                            switch 127 {
                            case 32768:
                                v10 -= (v2 | v10)
                            }
                        case "\n\t":
                            fmt.Println(v1, v2)
                        default:
                            for fi21 := 0; fi21 < 11; fi21++ {
                                var v22 string = "\n\t"
                                var v23 int16 = -32768
                                _ = fi21
                                _ = v22
                                _ = v23
                            }
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                case "hello":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        switch (fi24 / ((fi24) | 1)) {
                        case 2147483647:
                            if v1 {
                                switch "" {
                                case "\n\t":
                                    if (false || v1) {
                                        var v25 int32 = 1
                                        if v1 {
                                            v3 = v3
                                            var v26 uint8 = 229
                                            _ = v26
                                        }
                                        for fi27 := 0; fi27 < 3; fi27++ {
                                            fmt.Println(v5)
                                            fmt.Println(v3, v11, fi24)
                                            if v1 {
                                                var v28 uint64 = 2
                                                if len(v6) > 0 {
                                                    v6[0] = ("hello" + (("" + "hello") + ("foo bar" + "foo bar")))
                                                }
                                                _ = v28
                                            }
                                            _ = fi27
                                        }
                                        _ = v25
                                    } else {
                                        v11 -= -2.25
                                        var v29 uint32 = v5
                                        _ = v29
                                    }
                                case "foo bar":
                                    v5 = v5
                                case "":
                                    if (-268 < -2147483648) {
                                        if ((^(v2)) == (v2 + v10)) {
                                            v10 -= (-(v2))
                                            switch fi24 {
                                            case -338:
                                                v11 = (3.14159 + v11)
                                            case -9223372036854775808:
                                                for fi30 := 0; fi30 < 13; fi30++ {
                                                    delete(v4, (("" + "a") + "\n\t"))
                                                    if (!((1e10 * ((1.5 / (float64(-2.25) * v11)) - float64(1.5))) >= (((v11 / (-2.25 - v11)) - v11) * v11))) {
                                                        var v31 uint8 = 128
                                                        if (!(true || true)) {
                                                            v32 := GenS3{F0: (v31 + 0), F1: (v5 + (v5 % (((v5 % ((788) | 1))) | 1))), F2: 1}
                                                            if len(v6) > 0 {
                                                                v6[0] = ("" + (("\n\t" + "\n\t") + ("" + "\n\t")))
                                                            }
                                                            _ = v32
                                                        }
                                                        var v33 int64 = 0
                                                        fmt.Println(v33, v11)
                                                        var v34 uint32 = v5
                                                        _ = v31
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    var v35 uint32 = 2147483647
                                                    _ = fi30
                                                    _ = v35
                                                }
                                            default:
                                                _ = len(v4)
                                            }
                                        }
                                    }
                                default:
                                    v6 = append(v6, "foo bar")
                                }
                            } else {
                                var v36 uint8 = 1
                                _ = len(v4)
                                _ = v36
                            }
                        case 128:
                            if len(v6) > 0 {
                                v6[0] = ("" + (("foo bar" + (("\n\t" + ("a" + (("a" + (((("\n\t" + "") + ((((("a" + ("" + "")) + "\n\t") + "\n\t") + "a") + "\n\t")) + "a") + (("a" + "a") + (("a" + "\n\t") + ("a" + "foo bar"))))) + ""))) + "")) + ("foo bar" + ("hello" + ""))))
                            }
                        }
                        _ = fi24
                    }
                default:
                    v5 = v5
                }
                _ = v16
            } else {
                if false {
                    if false {
                        for fi37 := 0; fi37 < 20; fi37++ {
                            if (0 < 18446744073709551615) {
                                switch (fi37 >> ((1) & 63)) {
                                case 127:
                                    delete(v4, "foo bar")
                                case 128:
                                    v38 := GenS2{F0: 1}
                                    _ = v38
                                }
                            } else {
                                var v39 int64 = 9223372036854775807
                                fmt.Println(fi37, fi37, v1)
                                _ = v39
                            }
                            v40 := GenS3{F0: 42, F1: (v5 + v5), F2: -32768}
                            _ = fi37
                            _ = v40
                        }
                    }
                }
            }
        }
        _ = v12
    }
    if (!(191 > 0)) {
        if v1 {
            var v41 float32 = -2.25
            if v1 {
                fmt.Println(v3, v3, v11)
                if v1 {
                    v11 = (1.5 / (1.5 * -2.25))
                    fmt.Println(v3)
                    if true {
                        var v42 int8 = -1
                        v43 := GenS4{F0: -1, F1: v3}
                        fmt.Println(v42, v3)
                        _ = v42
                        _ = v43
                    } else {
                        v44 := GenS1{F0: v10}
                        var v45 float64 = v11
                        _ = v44
                        _ = v45
                    }
                } else {
                    v11 = ((v11 + v11) - float64(v41))
                    _ = len(v6)
                }
            }
            v1 = false
            _ = v41
        }
    }
    v5 += v5
    v46 := GenS5{F0: 42, F1: 0}
    if (v1 && false) {
        v47 := GenS5{F0: 42, F1: v3}
        switch 42 {
        case 42:
            v48 := GenS6{F0: (v5 >> ((624) & 63))}
            _ = v48
        case -1:
            switch 361 {
            case 148:
                v49 := GenS7{F0: 42, F1: 128}
                _ = v49
            case -646:
                if v1 {
                    v1 = v1
                }
            }
        default:
            _ = len(v4)
        }
        v11 = v11
        _ = v47
    }
    fmt.Println(v5, v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
