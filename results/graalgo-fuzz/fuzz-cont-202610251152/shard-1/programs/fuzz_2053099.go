package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{("" + "foo bar"), ("foo bar" + "foo bar"), (((("" + "foo bar") + "hello") + "hello") + "a"), ("a" + "")}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    switch 127 {
    case -2147483648:
        var v5 int64 = 42
        _ = v5
    case 42:
        var v6 bool = true
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    fmt.Println(v7)
    switch "\n\t" {
    case "foo bar":
        var v8 int = 481
        _ = v8
    case "\n\t":
        delete(v3, 0)
    case "hello":
        for fi9 := 0; fi9 < 17; fi9++ {
            _ = len(v3)
            v7 = (v7 * (-(v7)))
            _ = fi9
        }
    default:
        fmt.Println(v7)
    }
    var v10 int = 0
    v11 := GenS1{F0: 65535, F1: "hello", F2: (255 != 32767)}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint16 = 805
        _ = fi12
        _ = v13
    }
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "foo bar")
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        switch 747 {
        case 127:
            _ = len(v3)
        case -104:
            v15 := GenS2{F0: 1, F1: 431, F2: ((847 < 0) || true)}
            _ = v15
        }
        if false {
            var v16 int16 = 42
            _ = v16
        }
        if (541 < 4294967295) {
            v7 = v7
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v17 int8 = 127
            _ = v17
        } else {
            fmt.Println(fi14, v10)
        }
        _ = fi14
    }
    if true {
        _ = len(v1)
        var v18 uint64 = 374
        v19 := GenS2{F0: 42, F1: (^(v18)), F2: false}
        _ = v18
        _ = v19
    }
    switch v10 {
    case 594:
        var v20 int = -261
        _ = v20
    case 9223372036854775807:
        var v21 uint = 2
        _ = v21
    }
    switch int(v7) {
    case -1:
        v7 += (v7 << ((18446744073709551615) & 63))
    case 42:
        v7 -= 42
    case 0:
        v22 := v1["foo bar"]
        _ = v22
    default:
        for fi23 := 0; fi23 < 15; fi23++ {
            if false {
                v24 := GenS3{F0: 2, F1: 0}
                _ = v24
            }
            if ((797 == 0) || (!false)) {
                v1["hello"] = (v10 / ((fi23) | 1))
                if true {
                    var v25 int64 = v7
                    var v26 float64 = 0.0
                    v27 := GenS2{F0: 1, F1: 1, F2: false}
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 uint16 = 771
                        switch "" {
                        case "":
                            if (-2147483648 == 293) {
                                var v30 int16 = -266
                                _ = v30
                            } else {
                                v29 -= v29
                            }
                        case "a":
                            var v31 int32 = -1
                            v31 -= (v31 >> ((453) & 63))
                            _ = v31
                        case "foo bar":
                            v10 -= v10
                        default:
                            if true {
                                v7 -= (-(v7))
                            }
                        }
                        _ = fi28
                        _ = v29
                    }
                    var v32 uint = 717
                    _ = v32
                }
            } else {
                fmt.Println(v7, v10, fi23)
                if true {
                    for fi33 := 0; fi33 < 4; fi33++ {
                        var v34 float32 = (0.0 + 0.1)
                        v35 := GenS4{F0: 42, F1: (v7 | (v7 | (v7 - v7)))}
                        switch "a" {
                        case "":
                            v34 = v34
                        }
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    if false {
                        fmt.Println(fi23, v7)
                    } else {
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        var v36 string = "foo bar"
                        switch (fi23 * v10) {
                        case 32767:
                            if true {
                                v36 = v36
                            }
                        case 32768:
                            v7 = 1
                        case -9223372036854775808:
                            for fi37 := 0; fi37 < 18; fi37++ {
                                fmt.Println(v7)
                                v36 = "\n\t"
                                v2 = append(v2, ("\n\t" + "hello"))
                                _ = fi37
                            }
                        default:
                            v38 := GenS4{F0: 4294967295, F1: 0}
                            _ = v38
                        }
                        _ = v36
                    }
                }
                v10 -= -361
            }
            _ = fi23
        }
    }
    if (509 >= 65535) {
        for fi39 := 0; fi39 < 10; fi39++ {
            v7 = -9223372036854775808
            var v40 uint16 = 65535
            for fi41 := 0; fi41 < 10; fi41++ {
                if ((!((!false) || true)) && false) {
                    v4[842] = 2147483647
                    fmt.Println(v7, fi39)
                    var v42 uint64 = 2
                    var v43 uint = 18446744073709551615
                    fmt.Println(v43, v40, v10)
                    _ = v42
                    _ = v43
                }
                v44 := GenS2{F0: 62, F1: 18446744073709551615, F2: true}
                v10 -= fi39
                _ = fi41
                _ = v44
            }
            _ = fi39
            _ = v40
        }
        fmt.Println(v7, v7)
        if true {
            for fi45 := 0; fi45 < 6; fi45++ {
                if (32767 < 729) {
                    v4[4294967295] = 564
                    fmt.Println(v7)
                    var v46 uint8 = 147
                    _ = v46
                }
                _ = fi45
            }
        }
    }
    var v47 float32 = (((1.5 - 0.0) - ((3.14159 - 0.1) + 1e10)) * 1.5)
    switch "\n\t" {
    case "foo bar":
        for fi48 := 0; fi48 < 11; fi48++ {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            v7 += (-(v7))
            v4[0] = (fi48 - v10)
            _ = fi48
        }
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = "a"
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v7
}
