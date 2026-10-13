package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint8 = 42
    v2 := []string{"hello", "foo bar", "", ""}
    v3 := []int{-9223372036854775808, 669}
    v4 := make(map[string]uint32)
    var v5 int32 = 2147483647
    v1 -= (v1 % (((v1 * 1)) | 1))
    v6 := GenS1{F0: 2147483647, F1: 42}
    v1 = v1
    _ = len(v2)
    _ = len(v4)
    fmt.Println(v5, v1)
    v5 = v5
    v5 = (v5 + (v5 / ((v5) | 1)))
    var v7 uint16 = 518
    var v8 uint16 = v7
    switch 32767 {
    case -752:
        for fi9 := 0; fi9 < 20; fi9++ {
            if false {
                var v10 bool = false
                _ = v10
            }
            fmt.Println(v1, v1)
            _ = fi9
        }
    default:
        var v11 int64 = 987
        _ = v11
    }
    v7 = (v7 | (^(v8)))
    v1 += 127
    if (!false) {
        var v12 float64 = -2.25
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if false {
            var v14 int32 = 42
            if ((0.1 < 1.5) && (true || false)) {
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 string = (("foo bar" + "a") + "")
                    var v17 uint8 = (v1 * 127)
                    if false {
                        if (-32768 != 1000) {
                            if (!false) {
                                v1 -= (v1 / (((^(v17))) | 1))
                            } else {
                                v18 := GenS1{F0: 2, F1: uint(v17)}
                                if true {
                                    switch fi13 {
                                    case 1:
                                        fmt.Println(v1, v14)
                                    }
                                    v19 := GenS2{F0: 0, F1: -128}
                                    var v20 int16 = 32767
                                    _ = v19
                                    _ = v20
                                }
                                var v21 uint8 = v17
                                _ = v18
                                _ = v21
                            }
                            v16 = (((v16 + v16) + v16) + (v16 + v16))
                            fmt.Println(v5, v17, v14)
                        }
                        v22 := GenS2{F0: 158, F1: -37}
                        v1 += v17
                        _ = v22
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = ("foo bar" + "foo bar")
                }
                var v23 int16 = 0
                _ = v23
            }
            _ = v14
        } else {
            v1 = 186
            switch "\n\t" {
            case "foo bar":
                v8 = (v8 % ((42) | 1))
            case "\n\t":
                var v24 int32 = -93
                _ = v24
            case "a":
                v5 += (^(v5))
            default:
                for fi25 := 0; fi25 < 6; fi25++ {
                    _ = len(v4)
                    switch "foo bar" {
                    case "foo bar":
                        switch (fi13 / (((fi13 << ((44) & 63))) | 1)) {
                        case -2147483648:
                            var v26 uint64 = 0
                            _ = v26
                        case 247:
                            if len(v3) > 0 {
                                v3[0] = int(v8)
                            }
                            if true {
                                v5 += (v5 + v5)
                                switch "" {
                                case "hello":
                                    v5 += v5
                                }
                            } else {
                                fmt.Println(v7, v8)
                            }
                            var v27 float32 = 1e10
                            _ = v27
                        case -43:
                            v2 = append(v2, "a")
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "hello":
                        var v28 int = fi13
                        _ = v28
                    case "\n\t":
                        switch "\n\t" {
                        case "hello":
                            v5 = (v5 + (v5 & (-(v5))))
                        case "":
                            if (504 != -549) {
                                v8 = (v8 * 1)
                            } else {
                                v29 := GenS1{F0: 2, F1: 255}
                                _ = v29
                            }
                        default:
                            delete(v4, ("hello" + "a"))
                        }
                    default:
                        if false {
                            v4["a"] = 45
                            if (100 > 100) {
                                v7 = 2
                                var v30 int64 = -353
                                _ = v30
                            }
                        }
                    }
                    var v31 string = ("hello" + "")
                    _ = fi25
                    _ = v31
                }
            }
        }
        _ = fi13
    }
    switch ("foo bar" + "a") {
    case "hello":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
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
