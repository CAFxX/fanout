package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint32; F2 string }

func main() {
    var v1 float32 = (-2.25 * (0.1 - -2.25))
    var v2 uint64 = 520
    v3 := []uint32{2147483648}
    v4 := []int{-2147483648}
    var v5 int = -9223372036854775808
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint8 = 81
        if false {
            _ = len(v4)
            switch fi6 {
            case 128:
                var v8 uint = 777
                _ = v8
            }
        } else {
            switch "" {
            case "hello":
                v9 := GenS1{F0: int8(v7), F1: 9223372036854775807}
                _ = v9
            case "":
                var v10 string = "a"
                _ = v10
            case "foo bar":
                if true {
                    fmt.Println(v7, v1)
                    fmt.Println(v5, fi6)
                } else {
                    var v11 int = 224
                    switch (fi6 | v11) {
                    case 1:
                        if (1 > 1) {
                            for fi12 := 0; fi12 < 16; fi12++ {
                                fmt.Println(v1, fi6)
                                v1 += float32(-2.25)
                                switch "a" {
                                case "\n\t":
                                    if false {
                                        v5 = v11
                                        v1 = (0.1 + float32(((-2.25 + 3.14159) / 0.1)))
                                        v11 += (v11 + int(v7))
                                    } else {
                                        v4 = append(v4, 381)
                                        v2 = 4294967295
                                        if (0.0 < (-2.25 - 0.1)) {
                                            v1 = v1
                                            var v13 uint64 = v2
                                            _ = v13
                                        }
                                        v14 := GenS1{F0: 2, F1: 1}
                                        v15 := GenS1{F0: -128, F1: -715}
                                        _ = v14
                                        _ = v15
                                    }
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = int(v7)
                                    }
                                }
                                _ = fi12
                            }
                        }
                    case 9223372036854775807:
                        fmt.Println(v5)
                    }
                    _ = v11
                }
            default:
                switch -270 {
                case 32768:
                    v7 = (^(v7))
                case -2147483648:
                    v5 -= (fi6 - int(v2))
                case 42:
                    v16 := GenS2{F0: 595, F1: -538}
                    _ = v16
                }
            }
            v17 := GenS3{F0: (v7 | v7), F1: 51, F2: -32768}
            if false {
                v2 = v2
                v18 := GenS1{F0: 100, F1: -1}
                switch 793 {
                case 4294967295:
                    for fi19 := 0; fi19 < 2; fi19++ {
                        switch "hello" {
                        case "\n\t":
                            v20 := GenS3{F0: v7, F1: 1, F2: -322}
                            _ = v20
                        default:
                            v1 = 0.1
                        }
                        v21 := GenS1{F0: 42, F1: 1}
                        _ = fi19
                        _ = v21
                    }
                case 128:
                    fmt.Println(v5)
                case 53:
                    if (true || (false || false)) {
                        v1 = v1
                    } else {
                        v22 := GenS2{F0: -1, F1: 42}
                        v23 := GenS1{F0: int8(v2), F1: 9223372036854775807}
                        _ = v22
                        _ = v23
                    }
                default:
                    for fi24 := 0; fi24 < 18; fi24++ {
                        v25 := GenS4{F0: "hello", F1: 1}
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                        _ = fi24
                        _ = v25
                    }
                }
                _ = v18
            }
            _ = v17
        }
        var v26 uint64 = 375
        _ = fi6
        _ = v26
        _ = v7
    }
    v27 := GenS1{F0: -1, F1: -321}
    switch v5 {
    case 42:
        v1 = v1
    }
    v28 := GenS5{F0: -128, F1: 0, F2: ("a" + (("foo bar" + "foo bar") + ""))}
    fmt.Println(v5)
    v4 = append(v4, (v5 << ((216) & 63)))
    fmt.Println(v5, v2)
    fmt.Println(v2)
    if true {
        v4 = append(v4, v5)
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := GenS5{F0: 102, F1: 1, F2: "\n\t"}
            v5 = 828
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 10; fi31++ {
            v2 = 767
            v5 -= 42
            _ = fi31
        }
    }
    var v32 int = (v5 / ((583) | 1))
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 uint64 = (^(v2))
        var v35 float64 = 1e10
        var v36 uint32 = 2147483648
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println(v5, v5, v2)
    v5 += v32
    var v37 int8 = 75
    var v38 uint64 = (v2 - v2)
    v37 = 42
    v5 -= v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
