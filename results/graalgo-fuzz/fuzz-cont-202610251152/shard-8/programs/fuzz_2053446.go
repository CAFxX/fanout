package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }

func main() {
    var v1 int16 = -1
    var v2 int32 = -831
    v3 := []string{((((("foo bar" + ("\n\t" + "hello")) + ("a" + "foo bar")) + ("foo bar" + (("foo bar" + "\n\t") + "a"))) + "a") + ("foo bar" + ("a" + ("foo bar" + "\n\t")))), "\n\t"}
    switch 128 {
    case 9223372036854775807:
        var v4 int32 = v2
        _ = v4
    case -9223372036854775808:
        v1 += v1
    case 32767:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v5 uint32 = 42
        _ = v5
    }
    var v6 float64 = float64(0.1)
    v1 = (v1 % (((v1 & -440)) | 1))
    var v7 int8 = 127
    v8 := GenS1{F0: 41, F1: -686, F2: ("a" + ("a" + "hello"))}
    for fi9 := 0; fi9 < 4; fi9++ {
        if ("\n\t" > "hello") {
            var v10 int = fi9
            switch "hello" {
            case "foo bar":
                for fi11 := 0; fi11 < 4; fi11++ {
                    v2 = -270
                    switch "hello" {
                    case "hello":
                        var v12 float64 = 3.14159
                        _ = v12
                    case "foo bar":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "":
                        v7 = (v7 / ((v7) | 1))
                    }
                    _ = fi11
                }
            case "hello":
                v13 := GenS2{F0: -9223372036854775808}
                _ = v13
            case "\n\t":
                v1 = (^(v1))
            }
            if (false && (!false)) {
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS3{F0: -263, F1: (v7 | (v7 | v7))}
                    if (207 == 191) {
                        v16 := GenS4{F0: 2}
                        fmt.Println(v6, v2)
                        _ = v16
                    } else {
                        v17 := GenS3{F0: 42, F1: (v7 / ((v7) | 1))}
                        if (4294967295 != 2) {
                            if ("foo bar" != "foo bar") {
                                if false {
                                    var v18 uint = 538
                                    _ = v18
                                } else {
                                    var v19 uint8 = 127
                                    fmt.Println(v10, v2)
                                    v19 += 218
                                    _ = v19
                                }
                                v7 -= v7
                                _ = len(v3)
                                var v20 uint32 = 893
                                _ = v20
                            } else {
                                switch (fi14 / (((fi9 % ((fi14) | 1))) | 1)) {
                                case 32768:
                                    var v21 uint8 = 2
                                    _ = v21
                                case -2147483648:
                                    var v22 int32 = 48
                                    _ = v22
                                default:
                                    var v23 uint = 65535
                                    _ = v23
                                }
                                var v24 uint16 = 2
                                v25 := GenS2{F0: 909}
                                _ = v24
                                _ = v25
                            }
                            if ("\n\t" == "") {
                                _ = len(v3)
                            } else {
                                var v26 string = ("\n\t" + "\n\t")
                                _ = v26
                            }
                        } else {
                            var v27 uint64 = 42
                            _ = v27
                        }
                        var v28 int64 = 593
                        _ = v17
                        _ = v28
                    }
                    switch fi14 {
                    case 127:
                        v3 = append(v3, "\n\t")
                    default:
                        v10 += fi9
                    }
                    _ = fi14
                    _ = v15
                }
                v7 += (v7 ^ 127)
                switch v10 {
                case 32767:
                    fmt.Println(v7, v2)
                case -844:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            _ = v10
        }
        var v29 int32 = (-(v2))
        v30 := GenS5{F0: v7}
        _ = fi9
        _ = v29
        _ = v30
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
