package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint32; F1 string; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 bool }

func main() {
    var v1 int64 = 560
    v2 := []string{("" + (("hello" + "hello") + "foo bar")), ("\n\t" + (((("hello" + "") + "\n\t") + ("a" + ("" + "foo bar"))) + "foo bar")), ("hello" + "a")}
    var v3 uint8 = 127
    v4 := make(map[string]bool)
    var v5 int32 = -727
    v6 := []uint32{203}
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            switch fi8 {
            case 32767:
                v3 = v3
            case -2147483648:
                _ = len(v2)
                v9 := GenS1{F0: 42, F1: -9223372036854775808, F2: fi7}
                _ = len(v4)
                _ = v9
            default:
                switch "\n\t" {
                case "":
                    v10 := GenS2{F0: "hello", F1: 108}
                    _ = v10
                }
            }
            if true {
                var v11 string = "hello"
                if (!true) {
                    switch ((v11 + v11) + v11) {
                    case "\n\t":
                        v12 := GenS3{F0: 945, F1: "\n\t", F2: 620}
                        _ = v12
                    case "foo bar":
                        if false {
                            if true {
                                if false {
                                    fmt.Println(fi8)
                                } else {
                                    _ = len(v4)
                                    var v13 int8 = 100
                                    var v14 uint8 = v3
                                    fmt.Println(v13)
                                    v15 := GenS4{F0: v1}
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                }
                            } else {
                                v4["\n\t"] = false
                                fmt.Println(fi8)
                                v11 = v11
                            }
                        } else {
                            v1 -= (v1 - (v1 % ((v1) | 1)))
                        }
                    default:
                        if (2147483647 <= 94) {
                            v16 := GenS4{F0: int64(v5)}
                            _ = v16
                        } else {
                            if true {
                                fmt.Println(fi8, v11, v3)
                                v17 := GenS2{F0: v11, F1: 1}
                                v18 := GenS3{F0: 1, F1: v11, F2: 397}
                                _ = v17
                                _ = v18
                            }
                            if ((v1 == (v1 >> ((4294967295) & 63))) && (false || ((!false) && (100 >= 0)))) {
                                var v19 string = "a"
                                var v20 int32 = (^(v5))
                                v21 := GenS3{F0: 0, F1: "\n\t", F2: 2}
                                _ = v19
                                _ = v20
                                _ = v21
                            } else {
                                v6 = append(v6, 0)
                            }
                        }
                    }
                    var v22 float32 = 1e10
                    _ = v22
                }
                _ = v11
            }
            fmt.Println(fi8, v1)
            _ = fi8
        }
        _ = fi7
    }
    v23 := GenS2{F0: "hello", F1: 94}
    v24 := GenS1{F0: 2, F1: v1, F2: -499}
    var v25 uint64 = 18446744073709551615
    v3 -= v3
    var v26 string = ""
    v5 -= (v5 ^ (^(v5)))
    for fi27 := 0; fi27 < 11; fi27++ {
        switch fi27 {
        case -2147483648:
            if true {
                v28 := GenS5{F0: 0}
                var v29 string = "hello"
                _ = v28
                _ = v29
            } else {
                if ((1e10 <= -2.25) && true) {
                    v26 = ((v26 + "foo bar") + ("" + v26))
                } else {
                    v5 = (v5 / (((v5 << ((543) & 63))) | 1))
                    fmt.Println(v26, v5)
                }
                if len(v6) > 0 {
                    v6[0] = 42
                }
            }
        case 267:
            switch v26 {
            case "a":
                var v30 uint16 = 2
                _ = v30
            case "foo bar":
                var v31 bool = true
                _ = v31
            case "\n\t":
                v32 := GenS6{F0: (^(v3)), F1: false}
                _ = v32
            }
        case -136:
            v6 = append(v6, 2147483647)
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
