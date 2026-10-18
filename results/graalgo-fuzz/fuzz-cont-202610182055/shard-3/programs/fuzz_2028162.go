package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint64 = 727
    v2 := []string{"\n\t", "foo bar", (("\n\t" + "hello") + "")}
    var v3 uint16 = 687
    v4 := []string{"a"}
    var v5 int8 = -1
    var v6 string = ("" + "a")
    if (float32((((3.14159 - 0.0) * 1e10) + (3.14159 + ((0.0 / (1.5 + (3.14159 + (1.5 * 0.1)))) + (float64(1.5) / float64(1.5)))))) != 0.1) {
        v2 = append(v2, "")
        v7 := GenS1{F0: v1, F1: "\n\t"}
        _ = v7
    } else {
        v1 = 0
    }
    var v8 bool = true
    v8 = true
    var v9 uint64 = (v1 - (v1 / (((v1 & 1)) | 1)))
    for fi10 := 0; fi10 < 17; fi10++ {
        if len(v2) > 0 {
            v2[0] = v6
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int = (-(fi11))
        fmt.Println(v9)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v9 += v1
        if false {
            v4 = append(v4, (v6 + v6))
            fmt.Println(fi13)
        } else {
            if (-2147483648 > 42) {
                switch (v6 + v6) {
                case "":
                    v5 += (v5 ^ -5)
                default:
                    if ((!true) && false) {
                        v1 = v1
                    } else {
                        switch (fi13 / (((fi13 & fi13)) | 1)) {
                        case -2147483648:
                            var v14 uint16 = 310
                            _ = v14
                        default:
                            fmt.Println(v8, v5)
                        }
                        v5 -= -128
                    }
                }
                fmt.Println(v9)
                switch (("a" + v6) + (((v6 + (((v6 + "hello") + v6) + (v6 + v6))) + v6) + v6)) {
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = (v6 + "a")
                    }
                case "\n\t":
                    v3 = (^(v3))
                case "a":
                    if v8 {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            fmt.Println(v5)
                            fmt.Println(fi15)
                            if (65535 < 18446744073709551615) {
                                var v16 uint32 = 1
                                if false {
                                    switch (("\n\t" + v6) + v6) {
                                    case "a":
                                        if len(v2) > 0 {
                                            v2[0] = v6
                                        }
                                    case "foo bar":
                                        v3 -= (v3 * 528)
                                    case "\n\t":
                                        var v17 uint32 = v16
                                        _ = v17
                                    default:
                                        if false {
                                            if len(v4) > 0 {
                                                v4[0] = "foo bar"
                                            }
                                        } else {
                                            if (349 > 998) {
                                                v16 = v16
                                                if v8 {
                                                    var v18 bool = true
                                                    _ = v18
                                                } else {
                                                    var v19 uint8 = 28
                                                    v9 -= (v9 * v9)
                                                    _ = v19
                                                }
                                            }
                                            var v20 int64 = -836
                                            if v8 {
                                                v5 = (v5 % (((v5 / (((v5 - v5)) | 1))) | 1))
                                            } else {
                                                fmt.Println(v9, v3, v1)
                                                v21 := GenS1{F0: v1, F1: ("a" + ((("\n\t" + v6) + v6) + v6))}
                                                v8 = v8
                                                _ = v21
                                            }
                                            _ = v20
                                        }
                                    }
                                    var v22 string = ((v6 + ((v6 + ("hello" + "\n\t")) + ("foo bar" + ""))) + ("foo bar" + v6))
                                    v16 += uint32(v3)
                                    _ = v22
                                }
                                _ = v16
                            }
                            _ = fi15
                        }
                    } else {
                        var v23 uint8 = 42
                        _ = v23
                    }
                }
            }
            var v24 uint16 = v3
            v2 = append(v2, ((v6 + v6) + "foo bar"))
            _ = v24
        }
        var v25 uint64 = 137
        _ = fi13
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        switch (fi26 + fi26) {
        case 128:
            fmt.Println(v8, v5)
        case 1:
            v27 := GenS2{F0: ((v6 + v6) + "hello"), F1: 1}
            _ = v27
        }
        v9 -= v1
        var v28 uint8 = 1
        _ = fi26
        _ = v28
    }
    var v29 int16 = 42
    v5 = int8(v3)
    fmt.Println(v29)
    v8 = (0 > 2147483647)
    v30 := GenS3{F0: 186, F1: uint(v3)}
    if true {
        fmt.Println(v3, v29, v3)
        var v31 int8 = (v5 * -1)
        var v32 int32 = -2147483648
        _ = v31
        _ = v32
    } else {
        var v33 float32 = 0.1
        var v34 uint64 = v1
        fmt.Println(v3)
        _ = v33
        _ = v34
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
