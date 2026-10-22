package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16; F1 bool; F2 int16 }
type GenS3 struct { F0 uint8; F1 int8; F2 int8 }
type GenS4 struct { F0 int; F1 int8; F2 uint }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint32; F1 uint; F2 string }
type GenS7 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -957
    var v3 float32 = (0.1 + 3.14159)
    v4 := []uint32{224, 42, 2147483648}
    v5 := make(map[int]bool)
    v6 := GenS1{F0: 2, F1: "\n\t"}
    v4 = append(v4, 314)
    v2 = v2
    v2 = 4294967295
    v7 := GenS2{F0: 746, F1: false, F2: -913}
    v2 -= v2
    var v8 uint8 = 142
    fmt.Println(v3, v3)
    switch (v2 ^ (-(v2))) {
    case -9223372036854775808:
        var v9 int32 = -83
        _ = v9
    case 194:
        if (!false) {
            switch (v2 ^ v2) {
            case 676:
                switch 4294967295 {
                case 42:
                    v10 := GenS2{F0: 229, F1: (0.0 == ((-2.25 * v3) - 3.14159)), F2: 1000}
                    _ = v10
                case 468:
                    fmt.Println(v3, v3)
                case 127:
                    v11 := GenS3{F0: v8, F1: -1, F2: 1}
                    _ = v11
                }
            default:
                v12 := GenS4{F0: v2, F1: -1, F2: 447}
                _ = v12
            }
            v1[65535] = 1
        } else {
            switch "\n\t" {
            case "\n\t":
                var v13 bool = false
                _ = v13
            default:
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := GenS5{F0: -189, F1: v2}
                    _ = fi14
                    _ = v15
                }
            }
        }
    default:
        v3 += v3
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    v3 += ((1.5 + v3) + (v3 + v3))
    if true {
        for fi16 := 0; fi16 < 4; fi16++ {
            if (-605 != 42) {
                v8 += 255
                for fi17 := 0; fi17 < 6; fi17++ {
                    if true {
                        if true {
                            delete(v1, 2147483647)
                            var v18 uint = 4294967295
                            _ = v18
                        }
                        fmt.Println(fi17)
                    }
                    v3 -= -2.25
                    _ = fi17
                }
            } else {
                fmt.Println(v3)
                if false {
                    _ = len(v4)
                    var v19 int8 = 6
                    if ((v19 + 42) > v19) {
                        switch (("a" + (("foo bar" + "foo bar") + "a")) + "\n\t") {
                        case "foo bar":
                            var v20 bool = true
                            _ = v20
                        default:
                            _ = len(v1)
                        }
                        v21 := GenS2{F0: 42, F1: (65535 > 18446744073709551615), F2: 32767}
                        switch ("hello" + ("foo bar" + "foo bar")) {
                        case "\n\t":
                            delete(v5, (fi16 & (fi16 * 627)))
                        case "":
                            var v22 float32 = v3
                            _ = v22
                        case "a":
                            if true {
                                v19 -= (v19 >> ((18446744073709551615) & 63))
                                v23 := GenS1{F0: (v8 ^ v8), F1: ("hello" + (("hello" + (("hello" + ("a" + ("a" + "a"))) + ("" + ""))) + "foo bar"))}
                                var v24 float64 = float64(v3)
                                _ = v23
                                _ = v24
                            } else {
                                v25 := GenS6{F0: 65535, F1: 8, F2: ("hello" + "a")}
                                var v26 uint16 = 747
                                _ = v25
                                _ = v26
                            }
                        default:
                            if (!(false && true)) {
                                delete(v1, 0)
                            } else {
                                v3 = (v3 / 1.5)
                                v4 = append(v4, 2147483647)
                                v27 := GenS4{F0: fi16, F1: (-(v19)), F2: 1}
                                _ = v27
                            }
                        }
                        _ = v21
                    } else {
                        var v28 uint32 = 48
                        _ = v28
                    }
                    _ = v19
                } else {
                    if false {
                        fmt.Println(fi16, fi16)
                        if false {
                            for fi29 := 0; fi29 < 18; fi29++ {
                                var v30 uint64 = 18446744073709551615
                                var v31 uint8 = v8
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                        }
                    } else {
                        if true {
                            fmt.Println(v3, v3, v2)
                        } else {
                            if (763 <= -1) {
                                delete(v1, 2147483648)
                            } else {
                                v8 = (v8 | 42)
                                _ = len(v4)
                                if (0 < 18446744073709551615) {
                                    v8 -= v8
                                    var v32 int32 = -1
                                    if (!(42 == 47)) {
                                        switch 128 {
                                        case -2147483648:
                                            var v33 int32 = v32
                                            _ = v33
                                        case 9223372036854775807:
                                            v3 += (v3 + (v3 - (((v3 + (v3 + (3.14159 / v3))) / 1.0) + v3)))
                                        case 674:
                                            fmt.Println(v3, v2, v32)
                                        default:
                                            fmt.Println(v3, v8, v2)
                                        }
                                        var v34 string = "\n\t"
                                        _ = v34
                                    }
                                    _ = v32
                                } else {
                                    var v35 int64 = 9223372036854775807
                                    _ = v35
                                }
                            }
                        }
                    }
                }
                var v36 int16 = 1000
                _ = v36
            }
            var v37 uint8 = (^(v8))
            v37 = v37
            _ = fi16
            _ = v37
        }
        switch 4294967295 {
        case 504:
            v3 -= v3
        case -830:
            for fi38 := 0; fi38 < 10; fi38++ {
                switch "a" {
                case "foo bar":
                    switch "a" {
                    case "":
                        v39 := GenS7{F0: -370}
                        _ = v39
                    }
                }
                if (v8 != 1) {
                    switch (fi38 % ((-1) | 1)) {
                    case 0:
                        var v40 float64 = 0.1
                        _ = v40
                    case 128:
                        var v41 float64 = 0.1
                        _ = v41
                    case -825:
                        fmt.Println(fi38, v2)
                    }
                } else {
                    fmt.Println(v3, v3)
                    for fi42 := 0; fi42 < 2; fi42++ {
                        var v43 int8 = int8(v8)
                        var v44 int = (^(v2))
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                }
                _ = fi38
            }
        }
    } else {
        fmt.Println(v8, v3, v8)
    }
    fmt.Println(v2, v2)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
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
