package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []string{("" + ("a" + "hello")), "\n\t", ""}
    var v2 uint64 = 0
    var v3 uint32 = 2147483648
    if (-2147483648 <= -362) {
        fmt.Println(v2, v2)
    }
    for fi4 := 0; fi4 < 7; fi4++ {
        if false {
            var v5 uint32 = 291
            _ = v5
        }
        if false {
            v3 += v3
            if true {
                switch "\n\t" {
                case "a":
                    var v6 int = 577
                    _ = v6
                case "hello":
                    var v7 int16 = -1
                    _ = v7
                case "\n\t":
                    switch ("" + ("foo bar" + "hello")) {
                    case "foo bar":
                        for fi8 := 0; fi8 < 13; fi8++ {
                            if (!(42 != 42)) {
                                v9 := GenS1{F0: 1}
                                var v10 int16 = 287
                                v10 = (v10 * (v10 & v10))
                                fmt.Println(v10)
                                v10 = (v10 + v10)
                                _ = v10
                                _ = v9
                            } else {
                                var v11 uint32 = v3
                                _ = v11
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (false && true) {
                                v2 = 42
                            }
                            v2 = (v2 * v2)
                            switch ("" + (("" + "hello") + "hello")) {
                            case "":
                                var v12 bool = (1 < 100)
                                _ = v12
                            default:
                                var v13 int = fi4
                                _ = v13
                            }
                            _ = fi8
                        }
                    case "":
                        var v14 uint32 = v3
                        _ = v14
                    case "hello":
                        if false {
                            fmt.Println(v3)
                            var v15 uint = 0
                            _ = v15
                        } else {
                            var v16 int16 = -32768
                            v2 = (v2 >> ((18446744073709551615) & 63))
                            _ = v16
                        }
                    }
                default:
                    switch fi4 {
                    case 32768:
                        var v17 float64 = ((-2.25 / 3.14159) * 0.0)
                        _ = v17
                    case 9223372036854775807:
                        switch (fi4 ^ fi4) {
                        case 128:
                            fmt.Println(fi4)
                        case -503:
                            v18 := GenS1{F0: 327}
                            _ = v18
                        default:
                            v19 := GenS2{F0: 1, F1: 0, F2: (v3 % ((v3) | 1))}
                            _ = v19
                        }
                    case -2147483648:
                        switch (fi4 << ((669) & 63)) {
                        case 2147483647:
                            var v20 uint8 = 1
                            _ = v20
                        case 128:
                            v3 = v3
                        case 148:
                            var v21 int = fi4
                            _ = v21
                        default:
                            var v22 int64 = 1
                            _ = v22
                        }
                    default:
                        var v23 string = "hello"
                        _ = v23
                    }
                }
            }
        }
        _ = fi4
    }
    _ = len(v1)
    v24 := GenS2{F0: 252, F1: 49, F2: v3}
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            if true {
                switch fi26 {
                case -402:
                    if (true || (1.5 <= -2.25)) {
                        var v27 string = "foo bar"
                        _ = v27
                    } else {
                        if true {
                            v28 := GenS3{F0: (-818 != 0), F1: 434}
                            _ = v28
                        }
                    }
                }
            }
            _ = fi26
        }
        _ = fi25
    }
    var v29 int32 = 42
    var v30 bool = (false && true)
    v29 = v29
    var v31 uint64 = (v2 - (v2 & 18446744073709551615))
    var v32 int32 = v29
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
