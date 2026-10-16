package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int; F2 int32 }
type GenS6 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ""
    v3 := []string{v2, v2, v2, "a"}
    var v4 float64 = -2.25
    v5 := []string{v2, ("\n\t" + "")}
    v1[232] = 4294967295
    v4 = 0.0
    v5 = append(v5, "foo bar")
    v6 := GenS1{F0: 1}
    v7 := GenS1{F0: 32767}
    v8 := GenS2{F0: "\n\t", F1: 1, F2: 2}
    var v9 uint16 = 1
    v10 := GenS1{F0: 42}
    if false {
        v4 -= v4
        v2 = ""
    }
    v11 := GenS3{F0: 629, F1: 42}
    v2 = (v2 + "a")
    v2 = (("" + ("\n\t" + "foo bar")) + v2)
    switch 9223372036854775807 {
    case -610:
        switch -906 {
        case 119:
            var v12 uint32 = 776
            _ = v12
        case -2147483648:
            if true {
                fmt.Println(v2, v2, v4)
                v13 := GenS1{F0: -32768}
                _ = v13
            }
        default:
            if (835 == -816) {
                var v14 int64 = -1
                _ = v14
            } else {
                if false {
                    v9 -= v9
                    fmt.Println(v2, v2, v9)
                } else {
                    _ = len(v1)
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v9 -= (v9 & v9)
                        fmt.Println(v9, v4)
                        _ = fi15
                    }
                }
                if ((1 > 1) && true) {
                    v2 = v2
                    v16 := GenS4{F0: 18446744073709551615}
                    _ = v16
                } else {
                    if false {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            switch fi17 {
                            case 1:
                                var v18 int32 = 0
                                _ = v18
                            case -2147483648:
                                if ((fi17 + (fi17 >> ((0) & 63))) != fi17) {
                                    var v19 int8 = -1
                                    _ = v19
                                } else {
                                    v20 := GenS4{F0: 0}
                                    if false {
                                        var v21 bool = false
                                        _ = v21
                                    } else {
                                        v22 := GenS4{F0: 2}
                                        _ = v22
                                    }
                                    _ = v20
                                }
                            case -1:
                                v2 = v2
                            default:
                                switch fi17 {
                                case 32768:
                                    var v23 int = (fi17 * (fi17 - (fi17 * 42)))
                                    _ = v23
                                case 1:
                                    v4 = v4
                                case 0:
                                    v2 = "hello"
                                default:
                                    v2 = v2
                                }
                            }
                            v24 := v1[561]
                            if false {
                                v25 := v1[887]
                                v25 -= v25
                                v4 = v4
                                _ = v25
                            }
                            _ = fi17
                            _ = v24
                        }
                        v26 := GenS5{F0: -2147483648, F1: 0, F2: -2147483648}
                        v27 := GenS6{F0: 0}
                        _ = v26
                        _ = v27
                    }
                    v9 = 2
                }
                v28 := GenS5{F0: -33, F1: 148, F2: 42}
                _ = v28
            }
        }
    case -9223372036854775808:
        var v29 int = 4294967295
        _ = v29
    default:
        v2 = ("a" + v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
