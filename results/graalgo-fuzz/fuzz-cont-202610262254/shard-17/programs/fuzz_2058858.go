package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS6 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 int = 9223372036854775807
    v3 := []uint32{1, 2147483648, 0, 2}
    var v4 float64 = 1.5
    var v5 int16 = 32767
    var v6 int32 = 42
    var v7 uint = 18446744073709551615
    v2 = -1
    v8 := GenS1{F0: int32(v7)}
    if (!false) {
        v9 := GenS1{F0: (v6 * 773)}
        if false {
            var v10 int = -1
            var v11 uint64 = 1
            _ = v10
            _ = v11
        }
        _ = v9
    }
    switch "\n\t" {
    case "a":
        var v12 uint16 = 42
        _ = v12
    case "\n\t":
        if (36 == 100) {
            v6 = (v6 % ((715) | 1))
            var v13 int64 = -9223372036854775808
            _ = v13
        } else {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int16 = v5
                _ = fi14
                _ = v15
            }
            if (int(v7) >= (v2 % (((v2 | (v2 >> ((255) & 63)))) | 1))) {
                v4 += ((float64(3.14159) / v4) * 0.1)
                _ = len(v1)
                if (65535 < 1) {
                    var v16 string = "foo bar"
                    v7 += v7
                    var v17 int8 = -1
                    _ = v16
                    _ = v17
                }
            }
            if true {
                var v18 int32 = (v6 ^ v6)
                _ = len(v3)
                _ = v18
            }
        }
    case "foo bar":
        v19 := GenS2{F0: 1000}
        _ = v19
    default:
        if false {
            if ((v4 * v4) == 3.14159) {
                var v20 int8 = -87
                _ = v20
            } else {
                var v21 uint32 = uint32(v7)
                switch v2 {
                case 42:
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v2)
                    case "":
                        for fi22 := 0; fi22 < 13; fi22++ {
                            _ = len(v3)
                            v3 = append(v3, v21)
                            _ = fi22
                        }
                    }
                case 982:
                    switch 9223372036854775807 {
                    case 42:
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v21 += (v21 >> (((v7 & 255)) & 63))
                            _ = fi23
                        }
                    }
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        _ = len(v3)
                        if true {
                            v21 -= v21
                            v25 := GenS2{F0: v5}
                            _ = v25
                        }
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        _ = fi24
                    }
                default:
                    var v26 int = (v2 | (v2 << (((v7 / ((v7) | 1))) & 63)))
                    _ = v26
                }
                v27 := GenS3{F0: -822, F1: -1, F2: (v2 / ((4294967295) | 1))}
                _ = v21
                _ = v27
            }
        }
    }
    if false {
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(v4, v6)
            v4 -= v4
            _ = fi28
        }
        if (true || ((^(v7)) == v7)) {
            var v29 uint64 = 1
            var v30 int = v2
            var v31 float32 = 0.1
            _ = v29
            _ = v30
            _ = v31
        }
        _ = len(v1)
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        if false {
            switch ("foo bar" + ("\n\t" + "")) {
            case "a":
                _ = len(v1)
            case "\n\t":
                if (v4 != v4) {
                    switch (^(fi32)) {
                    case 128:
                        switch v2 {
                        case 751:
                            for fi33 := 0; fi33 < 3; fi33++ {
                                if true {
                                    if true {
                                        var v34 float32 = 0.0
                                        _ = v34
                                    }
                                    v35 := GenS1{F0: v6}
                                    if false {
                                        v4 -= v4
                                        if false {
                                            v2 += (fi32 % (((fi32 >> ((v7) & 63))) | 1))
                                            if false {
                                                var v36 uint8 = 0
                                                _ = len(v3)
                                                switch fi32 {
                                                case 42:
                                                    v7 = v7
                                                case -2147483648:
                                                    v36 -= v36
                                                case 1:
                                                    v37 := GenS4{F0: 127}
                                                    _ = v37
                                                }
                                                if false {
                                                    var v38 uint16 = 114
                                                    v6 = 236
                                                    if (true || (false || ((4294967295 == 1) || (false && (("a" + "") >= ("hello" + "a")))))) {
                                                        var v39 uint = 2
                                                        _ = v39
                                                    } else {
                                                        fmt.Println(v7, v5)
                                                        switch "a" {
                                                        case "hello":
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v40 int32 = (v6 >> ((v7) & 63))
                                                                _ = v40
                                                            }
                                                        case "a":
                                                            v4 = v4
                                                        default:
                                                            fmt.Println(v2)
                                                        }
                                                        v5 += 1000
                                                    }
                                                    _ = v38
                                                } else {
                                                    v41 := GenS5{F0: 0, F1: 42, F2: v5}
                                                    _ = v41
                                                }
                                                _ = v36
                                            } else {
                                                v7 = 266
                                                switch "\n\t" {
                                                case "a":
                                                    switch v2 {
                                                    case -2147483648:
                                                        if ((false || ((v2 & fi32) < (fi33 / (((v2 - -2147483648)) | 1)))) || ((fi33 << ((v7) & 63)) != v2)) {
                                                            var v42 uint16 = 2
                                                            _ = v42
                                                        } else {
                                                            v43 := GenS1{F0: (-(v6))}
                                                            var v44 uint32 = 2147483647
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                    }
                                                case "\n\t":
                                                    v7 -= 255
                                                }
                                            }
                                        }
                                    } else {
                                        switch "a" {
                                        case "a":
                                            v6 = v6
                                            fmt.Println(v2, v6)
                                            var v45 float64 = v4
                                            _ = v45
                                        case "hello":
                                            v46 := GenS1{F0: (^(v6))}
                                            _ = v46
                                        default:
                                            v1[""] = -2147483648
                                        }
                                    }
                                    _ = v35
                                }
                                var v47 uint = 65535
                                switch "" {
                                case "foo bar":
                                    v47 = v7
                                case "\n\t":
                                    v48 := GenS1{F0: (v6 / (((-(v6))) | 1))}
                                    _ = v48
                                case "":
                                    v5 -= v5
                                }
                                _ = fi33
                                _ = v47
                            }
                        case 32768:
                            var v49 uint = v7
                            _ = v49
                        case 731:
                            var v50 int32 = v6
                            _ = v50
                        default:
                            if ((v6 >> (((v7 | v7)) & 63)) <= (v6 % (((-(v6))) | 1))) {
                                v5 -= v5
                                var v51 uint32 = 4294967295
                                v52 := GenS6{F0: v7, F1: (v5 % (((v5 + v5)) | 1))}
                                _ = v51
                                _ = v52
                            }
                        }
                    default:
                        for fi53 := 0; fi53 < 7; fi53++ {
                            var v54 float32 = 3.14159
                            v6 = (v6 & (^(v6)))
                            _ = fi53
                            _ = v54
                        }
                    }
                    v55 := GenS4{F0: 127}
                    _ = v55
                }
            default:
                fmt.Println(v5, v4)
            }
        }
        _ = fi32
    }
    var v56 int16 = (v5 ^ 42)
    _ = len(v1)
    if false {
        var v57 float32 = 1e10
        if (93 < 105) {
            var v58 int = v2
            v7 = 264
            for fi59 := 0; fi59 < 20; fi59++ {
                _ = len(v3)
                _ = fi59
            }
            _ = v58
        }
        _ = v57
    }
    _ = len(v1)
    v60 := GenS6{F0: 65535, F1: v5}
    if false {
        v61 := GenS4{F0: 212}
        var v62 uint32 = 1
        for fi63 := 0; fi63 < 13; fi63++ {
            for fi64 := 0; fi64 < 1; fi64++ {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                _ = fi64
            }
            fmt.Println(fi63, fi63)
            _ = fi63
        }
        _ = v61
        _ = v62
    }
    for fi65 := 0; fi65 < 18; fi65++ {
        fmt.Println(v6)
        _ = fi65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v60
    _ = v7
    _ = v8
}
