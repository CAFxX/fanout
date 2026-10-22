package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 string; F1 string; F2 int64 }
type GenS4 struct { F0 string; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := []uint32{2147483647, 2147483647, 0, 708}
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 int8 = 127
    var v5 uint = 453
    if (73 <= 4294967295) {
        switch 127 {
        case 1:
            for fi6 := 0; fi6 < 20; fi6++ {
                v7 := GenS1{F0: 1, F1: 0, F2: (v4 ^ (v4 % ((-128) | 1)))}
                _ = fi6
                _ = v7
            }
        case -9223372036854775808:
            v5 = v5
        }
        for fi8 := 0; fi8 < 6; fi8++ {
            if true {
                switch (fi8 >> ((v5) & 63)) {
                case -2147483648:
                    var v9 int = fi8
                    _ = v9
                case 9223372036854775807:
                    var v10 uint32 = 1
                    _ = v10
                case 2147483647:
                    v11 := GenS2{F0: (^(v5)), F1: v5}
                    _ = v11
                default:
                    v4 = (^(v4))
                }
                v5 = (v5 + v5)
            } else {
                if (-662 >= 1) {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        v13 := GenS3{F0: "hello", F1: "hello", F2: 0}
                        var v14 uint8 = 255
                        switch "\n\t" {
                        case "\n\t":
                            v5 = v5
                        case "foo bar":
                            var v15 uint = 18446744073709551615
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
            }
            v16 := GenS2{F0: (v5 ^ 42), F1: v5}
            v4 = v4
            _ = fi8
            _ = v16
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            switch (("" + ("\n\t" + ("\n\t" + ("foo bar" + "")))) + ("hello" + (((("hello" + (("a" + ("" + "foo bar")) + "")) + "foo bar") + ("a" + "\n\t")) + "hello"))) {
            case "\n\t":
                var v18 float32 = ((3.14159 - 1e10) - float32(((1e10 + 1e10) * (((1.5 - 3.14159) * 0.1) * 1.5))))
                _ = v18
            }
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := GenS4{F0: "foo bar", F1: (^(v5)), F2: 4294967295}
                _ = fi19
                _ = v20
            }
            v4 = v4
            _ = fi17
        }
    } else {
        switch "" {
        case "hello":
            var v21 int = 128
            _ = v21
        case "":
            switch ("a" + "\n\t") {
            case "foo bar":
                fmt.Println(v5, v4)
            case "":
                fmt.Println(v4, v5)
            default:
                if (!(!true)) {
                    switch 32768 {
                    case 4294967295:
                        v5 -= (v5 & v5)
                    case 1:
                        var v22 float32 = ((1.5 * -2.25) / 1.5)
                        _ = v22
                    case -772:
                        v5 = (v5 & v5)
                    }
                    v23 := GenS4{F0: ("" + "a"), F1: v5, F2: 35}
                    _ = v23
                } else {
                    v5 = v5
                    for fi24 := 0; fi24 < 9; fi24++ {
                        for fi25 := 0; fi25 < 13; fi25++ {
                            var v26 int8 = v4
                            var v27 uint = 4294967295
                            switch (fi24 + fi25) {
                            case 42:
                                var v28 bool = true
                                _ = v28
                            case 1:
                                fmt.Println(v4)
                            }
                            var v29 string = ""
                            _ = fi25
                            _ = v26
                            _ = v27
                            _ = v29
                        }
                        v4 -= int8(fi24)
                        _ = fi24
                    }
                }
            }
        case "foo bar":
            fmt.Println(v5, v5)
        default:
            v30 := GenS2{F0: (^(v5)), F1: v5}
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        fmt.Println(v5, fi31, v5)
        v4 += v4
        switch (-(fi31)) {
        case 876:
            v5 -= (v5 ^ v5)
        case 42:
            for fi32 := 0; fi32 < 3; fi32++ {
                var v33 bool = (0.1 == 1.5)
                _ = fi32
                _ = v33
            }
        default:
            switch "hello" {
            case "a":
                fmt.Println(v5)
            case "":
                var v34 int64 = -968
                _ = v34
            default:
                if len(v1) > 0 {
                    v1[0] = 1
                }
            }
        }
        _ = fi31
    }
    _ = len(v1)
    var v35 uint8 = 2
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 uint32 = uint32(v5)
        if ((!(true && ((v35 / ((v35) | 1)) > 127))) && false) {
            v38 := GenS5{F0: (v35 | (v35 ^ (v35 % ((uint8(v5)) | 1)))), F1: uint64(v35)}
            _ = v38
        }
        fmt.Println(v4, v4)
        _ = fi36
        _ = v37
    }
    fmt.Println(v5, v4)
    v39 := v3[0]
    v40 := GenS4{F0: "", F1: (v5 * v5), F2: 332}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
