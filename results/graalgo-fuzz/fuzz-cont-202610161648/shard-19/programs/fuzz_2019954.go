package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int; F2 uint8 }
type GenS5 struct { F0 string }

func main() {
    var v1 int = 32767
    v2 := []uint32{151, 2147483647, 130}
    v3 := make(map[string]bool)
    var v4 int16 = 1
    var v5 bool = ((!false) || true)
    v6 := v3["a"]
    v3["foo bar"] = ((!v5) && false)
    v7 := GenS1{F0: 949, F1: v1}
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: 811, F1: v1}
        if (!false) {
            v5 = v5
            fmt.Println(fi8, v4, v4)
            switch v1 {
            case 127:
                fmt.Println(fi8, v6, v1)
            }
        }
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 663}
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS1{F0: 1, F1: (v1 | v1)}
        _ = fi11
        _ = v12
    }
    var v13 int32 = 332
    for fi14 := 0; fi14 < 6; fi14++ {
        switch 127 {
        case -326:
            v4 += (v4 * -880)
        case 431:
            switch v1 {
            case -1:
                fmt.Println(v13)
            case -53:
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS1{F0: 0, F1: fi14}
                    _ = fi15
                    _ = v16
                }
            case 127:
                if v6 {
                    if len(v2) > 0 {
                        v2[0] = 802
                    }
                    v17 := GenS2{F0: 40}
                    _ = v17
                }
            }
        default:
            switch "hello" {
            case "a":
                if (1 < 4294967295) {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v19 := GenS3{F0: -9223372036854775808}
                        _ = fi18
                        _ = v19
                    }
                    if (("\n\t" + "a") == ("\n\t" + "foo bar")) {
                        v20 := GenS2{F0: 542}
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 12; fi21++ {
                            if (251 < 255) {
                                switch "hello" {
                                case "":
                                    var v22 string = "foo bar"
                                    _ = v22
                                case "hello":
                                    var v23 int64 = -462
                                    _ = v23
                                case "\n\t":
                                    _ = len(v3)
                                default:
                                    var v24 uint = 4294967295
                                    _ = v24
                                }
                                v4 -= 405
                            } else {
                                var v25 uint32 = 1
                                v26 := GenS4{F0: 9223372036854775807, F1: (v1 >> ((1) & 63)), F2: 72}
                                switch "" {
                                case "a":
                                    delete(v3, ("" + (("" + "hello") + "hello")))
                                case "foo bar":
                                    var v27 string = "foo bar"
                                    _ = v27
                                case "":
                                    var v28 bool = (!v6)
                                    _ = v28
                                default:
                                    v4 = v4
                                }
                                _ = v25
                                _ = v26
                            }
                            _ = fi21
                        }
                    }
                    switch (fi14 - 601) {
                    case 2147483647:
                        if v5 {
                            var v29 float64 = 1e10
                            _ = v29
                        }
                    case 247:
                        v2 = append(v2, 1)
                    }
                } else {
                    _ = len(v3)
                }
            }
        }
        fmt.Println(v1)
        v30 := GenS5{F0: ("a" + "foo bar")}
        _ = fi14
        _ = v30
    }
    v13 = (v13 * (v13 - (v13 & (v13 >> ((655) & 63)))))
    v31 := GenS3{F0: 820}
    var v32 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
