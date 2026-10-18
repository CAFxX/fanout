package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64; F1 int16; F2 int }
type GenS3 struct { F0 uint }

func main() {
    var v1 float64 = (-2.25 / float64(0.1))
    var v2 int = 899
    v3 := []string{"", (((("foo bar" + "") + "") + ((("\n\t" + ("a" + "a")) + "hello") + "hello")) + ("\n\t" + ("foo bar" + "\n\t")))}
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: "foo bar", F1: 127}
    v2 = (v2 - v2)
    var v6 uint = 766
    fmt.Println(v2, v6, v6)
    var v7 uint64 = 18446744073709551615
    if (!false) {
        switch v2 {
        case 9223372036854775807:
            v2 = (v2 >> ((0) & 63))
        }
        if true {
            switch (^(v2)) {
            case 127:
                if (-793 < 1000) {
                    switch v2 {
                    case 32768:
                        if true {
                            fmt.Println(v1, v2, v2)
                        } else {
                            switch ("" + "hello") {
                            case "hello":
                                _ = len(v3)
                            case "foo bar":
                                var v8 uint32 = 42
                                _ = v8
                            }
                            for fi9 := 0; fi9 < 8; fi9++ {
                                for fi10 := 0; fi10 < 15; fi10++ {
                                    var v11 uint16 = 519
                                    v1 += v1
                                    _ = fi10
                                    _ = v11
                                }
                                var v12 float32 = 0.1
                                switch "foo bar" {
                                case "a":
                                    v7 = 4294967295
                                case "foo bar":
                                    for fi13 := 0; fi13 < 16; fi13++ {
                                        v14 := GenS1{F0: "\n\t", F1: (fi9 + 127)}
                                        _ = fi13
                                        _ = v14
                                    }
                                }
                                _ = fi9
                                _ = v12
                            }
                            for fi15 := 0; fi15 < 7; fi15++ {
                                v16 := GenS1{F0: "a", F1: (v2 / ((v2) | 1))}
                                _ = fi15
                                _ = v16
                            }
                        }
                    case -475:
                        var v17 uint16 = 97
                        _ = v17
                    default:
                        if true {
                            switch ("" + ("" + "\n\t")) {
                            case "hello":
                                var v18 uint32 = 2147483648
                                _ = v18
                            case "\n\t":
                                if (!(!false)) {
                                    if (false || true) {
                                        v19 := GenS2{F0: v7, F1: -1, F2: int(v7)}
                                        var v20 uint8 = 255
                                        _ = v19
                                        _ = v20
                                    }
                                } else {
                                    _ = len(v3)
                                    switch "foo bar" {
                                    case "a":
                                        v6 = (v6 | (v6 >> (((v6 % (((^(v6))) | 1))) & 63)))
                                    case "foo bar":
                                        var v21 int16 = 0
                                        _ = v21
                                    case "\n\t":
                                        v4[861] = 4294967295
                                    }
                                }
                            case "foo bar":
                                switch ("hello" + "") {
                                case "hello":
                                    var v22 int32 = -1
                                    _ = v22
                                case "a":
                                    fmt.Println(v6)
                                case "":
                                    v7 = (v7 & (v7 | v7))
                                default:
                                    fmt.Println(v1)
                                }
                            }
                        }
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case 0:
                switch v2 {
                case 2147483647:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 32767:
                    v23 := v4[0]
                    _ = v23
                default:
                    var v24 uint8 = 42
                    _ = v24
                }
            case 923:
                v25 := GenS3{F0: (^(v6))}
                _ = v25
            }
        }
        fmt.Println(v1, v7)
    }
    v1 = v1
    if (((v7 * v7) < v7) || true) {
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int8 = 121
            switch "a" {
            case "foo bar":
                v1 = v1
            case "\n\t":
                v2 = (^(fi26))
            case "":
                switch (fi26 * int(v7)) {
                case 9223372036854775807:
                    v28 := GenS1{F0: ("\n\t" + "hello"), F1: (-(fi26))}
                    _ = v28
                default:
                    v29 := GenS1{F0: "\n\t", F1: (v2 - int(v7))}
                    _ = v29
                }
            }
            _ = fi26
            _ = v27
        }
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        fmt.Println(v2, v7)
        v31 := GenS2{F0: (v7 << ((v6) & 63)), F1: -32768, F2: v2}
        _ = len(v3)
        _ = fi30
        _ = v31
    }
    var v32 uint32 = 534
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
