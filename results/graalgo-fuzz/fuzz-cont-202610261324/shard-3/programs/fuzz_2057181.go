package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]bool)
    v3 := []int{42, -646}
    var v4 int32 = 84
    var v5 int32 = (v4 | (v4 ^ (v4 / ((v4) | 1))))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v5
    v6 := GenS1{F0: (^(v4)), F1: 255}
    for fi7 := 0; fi7 < 18; fi7++ {
        delete(v2, fi7)
        _ = fi7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int16 = -32768
        var v10 uint64 = 42
        v10 = 484
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: (^(v4)), F1: 437}
    if (0 > (v5 << ((18446744073709551615) & 63))) {
        v12 := GenS2{F0: 2, F1: 255, F2: -128}
        if (!(339 <= -180)) {
            for fi13 := 0; fi13 < 9; fi13++ {
                v1 += v1
                if false {
                    fmt.Println(v1)
                    for fi14 := 0; fi14 < 17; fi14++ {
                        v15 := GenS2{F0: 127, F1: 42, F2: -1}
                        if (false || false) {
                            if true {
                                v1 += v1
                                v16 := GenS2{F0: 2, F1: 2, F2: -128}
                                switch fi13 {
                                case 0:
                                    switch "\n\t" {
                                    case "":
                                        if (!(-2.25 == -2.25)) {
                                            if (242 <= 32767) {
                                                var v17 string = (("a" + ("\n\t" + ("foo bar" + "a"))) + "foo bar")
                                                v5 -= 42
                                                _ = v17
                                            }
                                            v3 = append(v3, (fi14 ^ (fi14 & 2147483647)))
                                            if (128 != 42) {
                                                var v18 uint8 = 127
                                                v5 = v5
                                                fmt.Println(fi13)
                                                _ = v18
                                            }
                                        }
                                    case "hello":
                                        var v19 uint32 = 42
                                        _ = v19
                                    case "\n\t":
                                        var v20 uint64 = 18446744073709551615
                                        _ = v20
                                    }
                                case 128:
                                    fmt.Println(v5, fi13)
                                }
                                _ = v16
                            } else {
                                fmt.Println(v4, v1, fi13)
                                fmt.Println(fi14, fi13)
                            }
                            switch "" {
                            case "foo bar":
                                var v21 string = ""
                                _ = v21
                            case "\n\t":
                                fmt.Println(v4, fi14, v5)
                            case "a":
                                fmt.Println(fi13, v1, v4)
                            default:
                                v4 += -2147483648
                            }
                            switch 75 {
                            case 0:
                                switch 451 {
                                case 644:
                                    if false {
                                        switch "\n\t" {
                                        case "":
                                            v1 += v1
                                        case "a":
                                            var v22 uint = 2
                                            _ = v22
                                        }
                                    }
                                case 9223372036854775807:
                                    v23 := GenS2{F0: 2, F1: 128, F2: 42}
                                    _ = v23
                                case 4294967295:
                                    if true {
                                        var v24 uint16 = 42
                                        if len(v3) > 0 {
                                            v3[0] = fi14
                                        }
                                        _ = v24
                                    } else {
                                        v1 = (v1 & (v1 * v1))
                                        fmt.Println(v1, v5, v1)
                                    }
                                }
                            case 127:
                                var v25 float64 = 0.1
                                _ = v25
                            default:
                                v1 = v1
                            }
                        }
                        _ = fi14
                        _ = v15
                    }
                } else {
                    switch fi13 {
                    case 407:
                        v1 += v1
                    case 0:
                        v26 := GenS3{F0: 46}
                        _ = v26
                    }
                    var v27 int16 = 1
                    _ = v27
                }
                switch (fi13 * -482) {
                case 4294967295:
                    if ((1 >= 42) && false) {
                        var v28 int = fi13
                        _ = v28
                    } else {
                        fmt.Println(fi13, v5)
                    }
                }
                _ = fi13
            }
            delete(v2, 592)
        } else {
            v2[0] = true
            var v29 int32 = -487
            v30 := GenS3{F0: 88}
            _ = v29
            _ = v30
        }
        _ = v12
    } else {
        v5 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
