package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 int32 = 0
    v2 := []uint32{}
    var v3 uint8 = 1
    if true {
        var v4 uint32 = 561
        fmt.Println(v1, v1)
        _ = v4
    }
    if (1 >= 582) {
        for fi5 := 0; fi5 < 15; fi5++ {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 uint16 = uint16(v3)
                if true {
                    v1 = (v1 << ((0) & 63))
                    v7 -= v7
                } else {
                    if true {
                        v8 := GenS1{F0: (v7 * (^(v7))), F1: (v3 + v3)}
                        fmt.Println(v7, v1, fi6)
                        _ = v8
                    } else {
                        _ = len(v2)
                        fmt.Println(v3, fi6)
                    }
                }
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        var v9 uint32 = 4294967295
        switch (("\n\t" + ("\n\t" + ("foo bar" + (("" + (("foo bar" + ("hello" + ("" + (("hello" + "") + "foo bar")))) + "")) + (((("a" + "hello") + "hello") + (("a" + "hello") + ((((("\n\t" + "") + "\n\t") + "foo bar") + "\n\t") + "foo bar"))) + "\n\t"))))) + ("foo bar" + "a")) {
        case "foo bar":
            var v10 int = 127
            _ = v10
        default:
            v3 -= v3
        }
        _ = v9
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint64 = 846
        _ = fi11
        _ = v12
    }
    v3 += (v3 / (((v3 + 128)) | 1))
    v1 = 943
    _ = len(v2)
    var v13 uint8 = (v3 ^ v3)
    v14 := GenS2{F0: 180}
    for fi15 := 0; fi15 < 8; fi15++ {
        switch fi15 {
        case 898:
            fmt.Println(fi15, v13, fi15)
        }
        _ = fi15
    }
    switch "hello" {
    case "\n\t":
        for fi16 := 0; fi16 < 19; fi16++ {
            fmt.Println(v1, v1, v3)
            var v17 float64 = (1.5 / (3.14159 - 0.1))
            v2 = append(v2, 0)
            _ = fi16
            _ = v17
        }
    case "hello":
        var v18 string = "foo bar"
        _ = v18
    case "":
        if (73 > 42) {
            v1 -= v1
        }
    }
    v1 = -621
    var v19 bool = (!(!(((1 >= -899) || (42 <= 0)) && true)))
    for fi20 := 0; fi20 < 2; fi20++ {
        if true {
            if (2 >= 285) {
                switch "" {
                case "\n\t":
                    if (2 == 42) {
                        v21 := GenS3{F0: (v19 && ((!false) && v19)), F1: 584}
                        for fi22 := 0; fi22 < 6; fi22++ {
                            v13 -= v13
                            var v23 uint32 = 0
                            var v24 uint32 = (v23 + v23)
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v21
                    }
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "hello":
                    var v25 uint64 = 1
                    _ = v25
                default:
                    if v19 {
                        v26 := GenS2{F0: (v3 - 42)}
                        _ = v26
                    } else {
                        if false {
                            var v27 float64 = (0.0 * 1.5)
                            _ = v27
                        }
                    }
                }
                v3 -= v3
            } else {
                v28 := GenS4{F0: 293, F1: 22}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v29 := GenS3{F0: v19, F1: 32767}
                _ = v28
                _ = v29
            }
            v19 = v19
            var v30 int32 = 2147483647
            _ = v30
        }
        _ = fi20
    }
    v1 = v1
    v3 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
}
