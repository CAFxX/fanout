package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 float64 = (0.1 * 3.14159)
    var v2 float32 = float32(v1)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]int)
    v5 := []int{878, 128, 519}
    v6 := GenS1{F0: 42, F1: 2147483648}
    v4[42] = -900
    v1 += v1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (((v1 * ((1.5 + 3.14159) - ((v1 + (1e10 + v1)) + (v1 - v1)))) - 3.14159) == v1) {
        var v7 float32 = ((v2 + 0.0) * ((3.14159 + (3.14159 * v2)) + v2))
        if ((false && false) || (("hello" >= ("\n\t" + "foo bar")) || ((!(true && true)) && false))) {
            v1 -= 1.5
        } else {
            var v8 uint8 = 127
            v2 = ((v7 / (0.0 - ((-2.25 - v7) / 1.0))) - (1.5 + v7))
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            delete(v4, 348)
            _ = fi9
        }
        _ = v7
    }
    fmt.Println(v1, v1)
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint32 = 423
            v11 -= 4294967295
            _ = fi10
            _ = v11
        }
    }
    var v12 uint64 = 2
    switch ("a" + (("foo bar" + ("\n\t" + "hello")) + ("\n\t" + (("a" + "hello") + "foo bar")))) {
    case "foo bar":
        var v13 int16 = -32768
        _ = v13
    case "":
        if (!(!(!(!(true && ((((false || false) || (!false)) || (false && ("foo bar" <= "\n\t"))) || true)))))) {
            v14 := GenS1{F0: v12, F1: 318}
            _ = v14
        } else {
            v1 = (v1 + (v1 / float64((v2 - (v2 - v2)))))
        }
    default:
        v1 = v1
    }
    if (false || true) {
        if false {
            for fi15 := 0; fi15 < 6; fi15++ {
                fmt.Println(v2, v12, v1)
                _ = fi15
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                fmt.Println(v1, fi16)
                for fi17 := 0; fi17 < 2; fi17++ {
                    switch fi16 {
                    case 32768:
                        var v18 uint8 = 0
                        switch "a" {
                        case "\n\t":
                            _ = len(v4)
                        case "a":
                            switch ("foo bar" + "\n\t") {
                            case "\n\t":
                                v2 -= v2
                            case "hello":
                                v19 := v4[943]
                                _ = v19
                            case "foo bar":
                                if ((1 < (v18 - (v18 % ((v18) | 1)))) && false) {
                                    v20 := GenS1{F0: 42, F1: 0}
                                    _ = v20
                                } else {
                                    v18 = 128
                                    var v21 uint8 = v18
                                    v22 := GenS1{F0: 2, F1: 4294967295}
                                    _ = v21
                                    _ = v22
                                }
                            }
                        default:
                            v2 += float32(v1)
                        }
                        _ = v18
                    case 2147483647:
                        fmt.Println(fi16, fi17)
                    }
                    _ = fi17
                }
                fmt.Println(fi16, fi16, v2)
                _ = fi16
            }
            switch 128 {
            case -181:
                switch 32767 {
                case 151:
                    if (272 != 1) {
                        var v23 uint8 = 193
                        v2 = v2
                        _ = v23
                    }
                }
            case -9223372036854775808:
                var v24 float64 = -2.25
                _ = v24
            case 42:
                v2 = (float32(v1) * 0.0)
            }
        } else {
            switch "" {
            case "a":
                delete(v4, uint32(v12))
            }
            v5 = append(v5, 32767)
            v12 -= v12
        }
        fmt.Println(v12, v12)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
