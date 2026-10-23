package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{65535, 314}
    var v3 int64 = -1
    fmt.Println(v1)
    var v4 int = -1
    for fi5 := 0; fi5 < 3; fi5++ {
        if false {
            v6 := GenS1{F0: 2}
            fmt.Println(v1)
            switch (((("a" + ((("a" + "a") + "foo bar") + "")) + "hello") + ((("" + "a") + "") + "a")) + (("\n\t" + "\n\t") + "hello")) {
            case "hello":
                v7 := GenS2{F0: "hello", F1: (v1 | (^(v1)))}
                _ = v7
            case "foo bar":
                for fi8 := 0; fi8 < 1; fi8++ {
                    if true {
                        fmt.Println(v4, fi5)
                        v1 = (^(v1))
                    }
                    v4 -= int(v1)
                    if true {
                        if false {
                            v9 := GenS2{F0: "foo bar", F1: (v1 << ((18446744073709551615) & 63))}
                            _ = v9
                        } else {
                            if (0.0 <= -2.25) {
                                var v10 uint32 = 65535
                                fmt.Println(v4)
                                _ = v10
                            }
                            if true {
                                var v11 int8 = 42
                                v4 -= fi5
                                _ = v11
                            } else {
                                fmt.Println(v1)
                                if true {
                                    v3 = v3
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(fi5)
                                    var v12 uint64 = (v1 << ((65535) & 63))
                                    var v13 uint32 = 2147483648
                                    _ = v12
                                    _ = v13
                                }
                                switch (-(fi5)) {
                                case 32767:
                                    v1 += 1
                                case 1:
                                    fmt.Println(v3, fi5, fi5)
                                case 853:
                                    if false {
                                        switch (fi5 + fi8) {
                                        case 42:
                                            v14 := GenS1{F0: 115}
                                            _ = v14
                                        }
                                    } else {
                                        var v15 float64 = float64(1e10)
                                        _ = v15
                                    }
                                default:
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(fi8, v1)
                                    default:
                                        var v16 uint8 = 223
                                        _ = v16
                                    }
                                }
                                v17 := GenS1{F0: 138}
                                _ = v17
                            }
                        }
                    } else {
                        v4 += (-(fi5))
                    }
                    var v18 float32 = 1.5
                    var v19 uint8 = 42
                    _ = fi8
                    _ = v18
                    _ = v19
                }
            case "\n\t":
                switch (("" + "foo bar") + "") {
                case "":
                    var v20 string = "foo bar"
                    _ = v20
                case "a":
                    switch (("foo bar" + "\n\t") + ("\n\t" + "foo bar")) {
                    case "hello":
                        v1 = v1
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "foo bar":
                    v21 := GenS1{F0: 255}
                    _ = v21
                }
            }
            _ = v6
        }
        _ = fi5
    }
    var v22 bool = (-434 > 1)
    v23 := GenS3{F0: "hello", F1: 0}
    v1 -= v1
    switch "hello" {
    case "hello":
        v22 = (v22 && true)
    case "\n\t":
        var v24 uint16 = 2
        _ = v24
    case "":
        if v22 {
            for fi25 := 0; fi25 < 18; fi25++ {
                v1 = v1
                _ = fi25
            }
            switch -9223372036854775808 {
            case -2147483648:
                for fi26 := 0; fi26 < 20; fi26++ {
                    _ = len(v2)
                    v4 += (fi26 / ((fi26) | 1))
                    v27 := GenS3{F0: (((((((("" + "foo bar") + "foo bar") + "a") + "hello") + "foo bar") + "") + "a") + "hello"), F1: 130}
                    _ = fi26
                    _ = v27
                }
            case -900:
                var v28 int64 = (v3 | (v3 % (((v3 / ((-1) | 1))) | 1)))
                _ = v28
            case -454:
                if (1 > 403) {
                    if false {
                        var v29 uint16 = 2
                        var v30 uint8 = 128
                        var v31 string = ""
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        var v32 string = "\n\t"
                        var v33 uint16 = 1
                        var v34 int64 = 9223372036854775807
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 8; fi35++ {
                        v4 -= (v4 >> ((2) & 63))
                        _ = fi35
                    }
                } else {
                    fmt.Println(v22)
                }
            default:
                v36 := GenS3{F0: "foo bar", F1: 2147483647}
                _ = v36
            }
        } else {
            v37 := GenS4{F0: v1}
            if len(v2) > 0 {
                v2[0] = 521
            }
            var v38 uint8 = uint8(v1)
            _ = v37
            _ = v38
        }
    }
    if (2147483647 < 2147483647) {
        v4 = (v4 ^ (^(v4)))
    } else {
        v3 += 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
