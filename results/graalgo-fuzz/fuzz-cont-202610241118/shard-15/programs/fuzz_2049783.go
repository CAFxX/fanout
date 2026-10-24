package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int64; F2 uint16 }
func genFn0() uint {
    var v5 uint64 = 1
    fmt.Println(v5)
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{909, 32768, 710, 674}
    var v3 int = 0
    var v4 int64 = 1
    var v6 bool = true
    v2 = append(v2, (v3 & v3))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 uint8 = 255
    var v8 int8 = 97
    v9 := GenS1{F0: ("a" + "hello")}
    v7 = v7
    switch "a" {
    case "a":
        if true {
            if (!((v3 + v3) == (v3 << ((320) & 63)))) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 int = -9223372036854775808
                    _ = fi10
                    _ = v11
                }
                switch v3 {
                case -583:
                    v12 := GenS1{F0: "hello"}
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 2; fi13++ {
                        if (!v6) {
                            _ = len(v1)
                            if (true || (!v6)) {
                                v14 := GenS1{F0: "foo bar"}
                                _ = v14
                            } else {
                                switch (-(fi13)) {
                                case 9223372036854775807:
                                    v1[4294967295] = v6
                                case -9223372036854775808:
                                    if (!(true || v6)) {
                                        var v15 float64 = (float64(1e10) / float64(1e10))
                                        var v16 uint16 = 2
                                        _ = v15
                                        _ = v16
                                    }
                                case 128:
                                    v1[2147483647] = (0.1 <= (1.5 - 0.0))
                                default:
                                    _ = len(v2)
                                }
                                if ((!false) || false) {
                                    for fi17 := 0; fi17 < 12; fi17++ {
                                        var v18 string = "hello"
                                        _ = fi17
                                        _ = v18
                                    }
                                    var v19 uint = 18446744073709551615
                                    for fi20 := 0; fi20 < 14; fi20++ {
                                        v7 = (v7 & v7)
                                        _ = fi20
                                    }
                                    _ = v19
                                } else {
                                    switch fi13 {
                                    case 643:
                                        var v21 uint8 = v7
                                        _ = v21
                                    case 4294967295:
                                        v7 = (v7 / (((v7 * 126)) | 1))
                                    case 32767:
                                        if (42 > v4) {
                                            var v22 bool = (-2147483648 <= 2147483647)
                                            _ = v22
                                        } else {
                                            var v23 int64 = -637
                                            v4 = (^(v23))
                                            v6 = false
                                            _ = v23
                                        }
                                    default:
                                        switch ("a" + "") {
                                        case "foo bar":
                                            v24 := GenS1{F0: (((("\n\t" + "hello") + "foo bar") + "foo bar") + "foo bar")}
                                            _ = v24
                                        }
                                    }
                                    var v25 uint32 = 65535
                                    _ = v25
                                }
                            }
                            var v26 uint = 255
                            _ = v26
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                if len(v2) > 0 {
                                    v2[0] = (^(fi13))
                                }
                            case "a":
                                var v27 float32 = ((0.1 + (0.0 + 0.1)) + (0.1 / 1.0))
                                _ = v27
                            case "foo bar":
                                v6 = (v6 && ((float32(float64(((((-2.25 + float32(((0.1 * 3.14159) * 3.14159))) - 0.1) - 3.14159) + (((1e10 + ((float32(3.14159) + 1.5) - -2.25)) - (-2.25 - 0.0)) + 1.5)))) - float32((1.5 + 0.0))) != -2.25))
                            }
                            var v28 int = fi13
                            _ = v28
                        }
                        v4 = v4
                        _ = fi13
                    }
                }
                v29 := GenS2{F0: 1, F1: (v4 + v4), F2: 65535}
                _ = v29
            }
        }
    case "foo bar":
        for fi30 := 0; fi30 < 6; fi30++ {
            fmt.Println(v7, v4, v3)
            _ = fi30
        }
    case "":
        var v31 uint64 = 338
        _ = v31
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v32 int64 = (v4 / (((v4 >> ((234) & 63))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
