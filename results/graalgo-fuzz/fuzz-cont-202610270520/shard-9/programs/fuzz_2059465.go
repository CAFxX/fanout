package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint32; F1 int; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }

func main() {
    v1 := []uint32{2147483647, 2147483648}
    var v2 string = "a"
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 1000, F1: -1}
    v2 = (v2 + "a")
    fmt.Println(v2)
    if true {
        var v5 float32 = 3.14159
        for fi6 := 0; fi6 < 12; fi6++ {
            delete(v3, 0)
            switch v2 {
            case "a":
                v2 = "foo bar"
            case "foo bar":
                switch "" {
                case "foo bar":
                    for fi7 := 0; fi7 < 3; fi7++ {
                        fmt.Println(fi7)
                        v2 = "foo bar"
                        switch (v2 + v2) {
                        case "\n\t":
                            v2 = v2
                        case "foo bar":
                            var v8 uint64 = 942
                            _ = v8
                        case "hello":
                            if false {
                                switch fi6 {
                                case 0:
                                    switch (v2 + ("\n\t" + v2)) {
                                    case "\n\t":
                                        v9 := GenS2{F0: 745, F1: 937, F2: 88}
                                        v10 := GenS1{F0: 511, F1: 42}
                                        _ = v10
                                        _ = v9
                                    case "a":
                                        v2 = "a"
                                    }
                                default:
                                    var v11 string = v2
                                    _ = v11
                                }
                                if (65535 >= 303) {
                                    if (int64(fi6) == 42) {
                                        switch (v2 + v2) {
                                        case "foo bar":
                                            var v12 string = ("foo bar" + v2)
                                            _ = v12
                                        default:
                                            if true {
                                                v13 := GenS2{F0: 2, F1: 18446744073709551615, F2: 128}
                                                var v14 int = fi7
                                                if true {
                                                    delete(v3, 65535)
                                                    if len(v1) > 0 {
                                                        v1[0] = 885
                                                    }
                                                }
                                                fmt.Println(fi6)
                                                _ = v13
                                                _ = v14
                                            }
                                        }
                                        if true {
                                            v15 := GenS3{F0: 32767, F1: (fi7 + fi7)}
                                            var v16 float64 = ((0.0 + (1.5 + 0.0)) - (-2.25 / 0.1))
                                            v17 := GenS4{F0: 73, F1: (fi6 / ((fi7) | 1)), F2: ("\n\t" + v2)}
                                            if ((("" + "foo bar") + ("a" + "foo bar")) > v2) {
                                                if true {
                                                    v5 = (3.14159 / v5)
                                                    var v18 string = v2
                                                    v5 -= 0.0
                                                    _ = v18
                                                } else {
                                                    var v19 int8 = -1
                                                    if ((v19 > (v19 & v19)) || (fi7 != fi7)) {
                                                        v20 := GenS5{F0: (fi7 << ((966) & 63))}
                                                        _ = len(v1)
                                                        _ = len(v1)
                                                        _ = v20
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = 502
                                                    }
                                                    _ = v19
                                                }
                                                v16 = float64(v5)
                                            }
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                        }
                                        v5 -= v5
                                    }
                                    fmt.Println(fi7, v2)
                                }
                            }
                        }
                        _ = fi7
                    }
                case "a":
                    if false {
                        if (!(!(127 > 42))) {
                            var v21 int16 = 1000
                            fmt.Println(fi6)
                            var v22 uint32 = 527
                            _ = v21
                            _ = v22
                        }
                        v2 = v2
                        var v23 uint8 = 42
                        _ = v23
                    }
                }
            case "":
                for fi24 := 0; fi24 < 7; fi24++ {
                    switch 42 {
                    case 1:
                        fmt.Println(v2, v5, fi6)
                    default:
                        switch fi24 {
                        case 2147483647:
                            if ((-532 > -32768) && (false || false)) {
                                v2 = "\n\t"
                            }
                        case -9223372036854775808:
                            var v25 int16 = -1
                            _ = v25
                        default:
                            v5 += v5
                        }
                    }
                    var v26 int64 = 0
                    var v27 int16 = 1000
                    _ = fi24
                    _ = v26
                    _ = v27
                }
            }
            _ = fi6
        }
        v28 := GenS6{F0: 588}
        _ = v28
        _ = v5
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 uint64 = 18446744073709551615
        _ = fi29
        _ = v30
    }
    v2 = ("\n\t" + v2)
    if false {
        _ = len(v3)
    } else {
        v2 = "a"
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 string = ("hello" + "hello")
            fmt.Println(v2, fi31)
            for fi33 := 0; fi33 < 16; fi33++ {
                v2 = ((v32 + ("a" + (("hello" + (v32 + v32)) + v32))) + v2)
                v32 = "foo bar"
                _ = fi33
            }
            _ = fi31
            _ = v32
        }
        var v34 uint8 = 127
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        fmt.Println(fi35, v2)
        v2 = (("a" + v2) + (v2 + "a"))
        v1 = append(v1, 181)
        _ = fi35
    }
    v36 := GenS4{F0: 42, F1: 1, F2: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
