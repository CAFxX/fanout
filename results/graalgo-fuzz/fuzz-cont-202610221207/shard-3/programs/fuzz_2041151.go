package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int64; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 0.1
    v3 := []uint32{744, 2, 4294967295}
    var v4 int = -9223372036854775808
    var v5 float32 = (v2 * 1e10)
    var v6 int16 = 32767
    var v7 string = "foo bar"
    if false {
        if true {
            for fi8 := 0; fi8 < 16; fi8++ {
                fmt.Println(v5, v4)
                _ = fi8
            }
            v7 = v7
            fmt.Println(v4)
        }
    }
    if (0.0 != v5) {
        fmt.Println(v2, v5, v4)
        var v9 int8 = 1
        v10 := GenS1{F0: "a", F1: 308}
        _ = v10
        _ = v9
    } else {
        fmt.Println(v2)
        var v11 int64 = -9223372036854775808
        _ = v11
    }
    v12 := GenS2{F0: (v4 % ((v4) | 1)), F1: int16(v4)}
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(v2, fi13)
        _ = fi13
    }
    v14 := GenS1{F0: ("\n\t" + v7), F1: 2147483648}
    switch -2147483648 {
    case -2147483648:
        if false {
            var v15 uint8 = 2
            _ = v15
        } else {
            var v16 bool = false
            _ = v16
        }
    }
    switch v7 {
    case "hello":
        var v17 float64 = 1.5
        _ = v17
    default:
        v7 = v7
    }
    if false {
        v18 := GenS3{F0: false}
        _ = v18
    } else {
        switch v7 {
        case "":
            for fi19 := 0; fi19 < 18; fi19++ {
                switch "hello" {
                case "\n\t":
                    var v20 string = ("\n\t" + (v7 + (v7 + v7)))
                    _ = v20
                case "hello":
                    switch fi19 {
                    case -1:
                        v21 := GenS1{F0: ((v7 + (("a" + v7) + v7)) + v7), F1: 242}
                        _ = v21
                    default:
                        v7 = "\n\t"
                    }
                default:
                    v22 := GenS4{F0: 1, F1: 2147483648}
                    _ = v22
                }
                _ = fi19
            }
        case "\n\t":
            switch 127 {
            case 1:
                for fi23 := 0; fi23 < 18; fi23++ {
                    v2 += -2.25
                    v7 = v7
                    v2 = v5
                    _ = fi23
                }
            default:
                var v24 uint64 = 4294967295
                _ = v24
            }
        }
        switch v7 {
        case "":
            for fi25 := 0; fi25 < 12; fi25++ {
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 uint64 = 18446744073709551615
                    if true {
                        var v28 int16 = v6
                        _ = len(v3)
                        var v29 uint16 = 182
                        v30 := GenS3{F0: true}
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    if (-752 >= 9223372036854775807) {
                        var v31 int32 = 493
                        if (false || true) {
                            var v32 uint = 42
                            _ = v32
                        } else {
                            fmt.Println(v2, fi25, v7)
                            v31 = -2147483648
                            fmt.Println(v7, v5, v31)
                        }
                        _ = v31
                    }
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
        case "foo bar":
            switch (-(v4)) {
            case 127:
                v3 = append(v3, 1)
            case 60:
                for fi33 := 0; fi33 < 9; fi33++ {
                    for fi34 := 0; fi34 < 9; fi34++ {
                        switch 1 {
                        case -848:
                            v6 = (v6 + (v6 - (v6 | v6)))
                        case 2147483647:
                            v35 := v1[((v7 + "foo bar") + v7)]
                            _ = v35
                        case -9223372036854775808:
                            _ = len(v3)
                            fmt.Println(v6, v7)
                            var v36 uint8 = 128
                            _ = v36
                        default:
                            var v37 float64 = 0.0
                            _ = v37
                        }
                        _ = fi34
                    }
                    switch (fi33 >> ((42) & 63)) {
                    case -2147483648:
                        delete(v1, (v7 + (v7 + v7)))
                    case -127:
                        v38 := GenS5{F0: 887, F1: -9223372036854775808, F2: (((v7 + (("hello" + v7) + "a")) < v7) && true)}
                        _ = v38
                    case 9223372036854775807:
                        fmt.Println(v7, v6, v2)
                    default:
                        switch 42 {
                        case 1:
                            var v39 float64 = 0.0
                            _ = v39
                        case -429:
                            var v40 uint8 = 194
                            _ = v40
                        case 159:
                            var v41 uint16 = 42
                            _ = v41
                        }
                    }
                    _ = fi33
                }
            case 0:
                var v42 int32 = 2147483647
                _ = v42
            default:
                switch v7 {
                case "\n\t":
                    switch (v4 ^ (-(v4))) {
                    case 678:
                        var v43 uint32 = 0
                        _ = v43
                    default:
                        if ((!true) || (false || true)) {
                            delete(v1, ("\n\t" + ("a" + "")))
                        }
                    }
                }
            }
        }
        var v44 uint32 = 42
        _ = v44
    }
    v2 -= (3.14159 + float32(((((1.5 / ((float64(0.0) + -2.25) * (0.0 + 3.14159))) * (1e10 / ((0.1 - ((-2.25 - 1.5) / 0.1)) * 0.1))) - ((0.0 - -2.25) - 0.1)) / 1.0)))
    fmt.Println(v7, v2, v2)
    var v45 int64 = 159
    var v46 uint16 = 740
    for fi47 := 0; fi47 < 12; fi47++ {
        fmt.Println(v45)
        _ = fi47
    }
    v45 = v45
    var v48 uint16 = (v46 - 65535)
    delete(v1, "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
