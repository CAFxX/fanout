package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 816
    var v3 uint32 = 4294967295
    v4 := []uint32{v3, (^(v3))}
    v5 := GenS1{F0: 993, F1: true, F2: -2147483648}
    v6 := GenS1{F0: 32768, F1: (true && (v2 > v2)), F2: -1}
    fmt.Println(v3, v3)
    if (-1 > 1) {
        var v7 uint = 18446744073709551615
        _ = v7
    } else {
        var v8 string = "hello"
        var v9 string = v8
        var v10 int32 = 0
        _ = v10
        _ = v8
        _ = v9
    }
    if true {
        if true {
            _ = len(v1)
        } else {
            v11 := GenS2{F0: 4294967295, F1: -710, F2: 837}
            v2 = (v2 + 432)
            _ = v11
        }
    }
    if false {
        _ = len(v4)
    } else {
        v12 := GenS1{F0: -799, F1: (true || (false || false)), F2: 0}
        _ = v12
    }
    switch "\n\t" {
    case "\n\t":
        v2 = 32767
    case "hello":
        v3 -= 0
    }
    if true {
        switch 960 {
        case 127:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 uint16 = v2
                if true {
                    switch fi13 {
                    case -2147483648:
                        v15 := v1[("hello" + "\n\t")]
                        _ = v15
                    case 464:
                        _ = len(v4)
                    default:
                        v2 += 509
                    }
                    var v16 float32 = 0.0
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v16 = 0.0
                        fmt.Println(fi17)
                        if (197 != 0) {
                            v18 := GenS1{F0: (fi13 << ((108) & 63)), F1: (!(!(2 > 542))), F2: 2147483647}
                            _ = v18
                        } else {
                            if true {
                                if true {
                                    if ((1e10 > (1.5 + 0.0)) && true) {
                                        v2 += 927
                                        if (((fi17 << ((378) & 63)) >= (^(fi13))) || true) {
                                            fmt.Println(fi13, fi17, v3)
                                        }
                                        var v19 int = fi13
                                        _ = v19
                                    } else {
                                        v4 = append(v4, 0)
                                        v16 = 3.14159
                                        if true {
                                            fmt.Println(v3)
                                        } else {
                                            fmt.Println(v14)
                                            fmt.Println(fi17)
                                            switch (fi17 + (fi13 * fi13)) {
                                            case -827:
                                                var v20 int16 = 42
                                                _ = v20
                                            case 127:
                                                var v21 float64 = (0.1 - (1e10 - float64(v16)))
                                                _ = v21
                                            case -631:
                                                v16 = 1e10
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if false {
                                                        v22 := GenS3{F0: 100}
                                                        var v23 uint64 = 42
                                                        var v24 int32 = 2147483647
                                                        _ = v22
                                                        _ = v23
                                                        _ = v24
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        switch -614 {
                                                        case 1:
                                                            fmt.Println(v2, v14, v2)
                                                            v25 := GenS1{F0: 9223372036854775807, F1: (18446744073709551615 > 4294967295), F2: 2147483647}
                                                            _ = v25
                                                        default:
                                                            v16 += v16
                                                        }
                                                        var v26 int8 = 100
                                                        _ = v26
                                                    }
                                                default:
                                                    var v27 int = (fi17 / (((fi17 / ((fi13) | 1))) | 1))
                                                    _ = v27
                                                }
                                            }
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        fmt.Println(v16, fi13, v2)
                                    }
                                } else {
                                    v28 := GenS4{F0: 674, F1: (true && (747 > (v3 & v3)))}
                                    v2 = (v2 + (v2 + (v2 | v14)))
                                    var v29 int8 = 100
                                    _ = v28
                                    _ = v29
                                }
                            } else {
                                fmt.Println(v16)
                            }
                            var v30 uint16 = 42
                            v2 += (v2 >> ((65535) & 63))
                            _ = v30
                        }
                        if len(v4) > 0 {
                            v4[0] = (v3 | v3)
                        }
                        var v31 int64 = 9223372036854775807
                        v16 = 3.14159
                        _ = fi17
                        _ = v31
                    }
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
