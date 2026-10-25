package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 string; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint8; F2 bool }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 int16; F2 uint64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 bool; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (float32(1.5) / (3.14159 / float32(1.5)))
    v3 := []string{"\n\t", ("hello" + "a"), ""}
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    var v6 float64 = (1.5 * 3.14159)
    v7 := GenS1{F0: 1, F1: 2, F2: 100}
    v2 = v2
    v8 := GenS1{F0: 128, F1: 126, F2: 42}
    v2 -= (v2 * v2)
    fmt.Println(v2)
    if (377 > 194) {
        if true {
            v6 -= v6
            var v9 int = 993
            switch (v9 >> ((304) & 63)) {
            case 127:
                _ = len(v3)
            case 32768:
                v1[-1] = (("" + "a") + "a")
            case 1:
                v6 = 3.14159
            }
            _ = v9
        } else {
            v10 := GenS2{F0: -1}
            var v11 float64 = 0.0
            _ = v10
            _ = v11
        }
        fmt.Println(v6)
    } else {
        v12 := GenS3{F0: 2, F1: "a", F2: 18446744073709551615}
        if (false || (1 > 1)) {
            v6 = float64(v2)
            var v13 int64 = 9223372036854775807
            var v14 uint = 65535
            _ = v13
            _ = v14
        } else {
            if false {
                var v15 uint64 = 1
                _ = v15
            }
            if ((0 == 445) || true) {
                v6 = float64((((v2 / 0.1) + (float32(v6) / 1e10)) / -2.25))
            } else {
                delete(v1, 32768)
                v6 += 0.0
                v3 = append(v3, ("hello" + "a"))
            }
            if (0.1 == v6) {
                v16 := GenS1{F0: 222, F1: 980, F2: 100}
                if false {
                    v2 += 0.0
                    var v17 uint32 = 840
                    _ = v17
                }
                switch 42 {
                case 0:
                    if false {
                        v18 := GenS4{F0: 4294967295}
                        fmt.Println(v6, v2)
                        v19 := GenS5{F0: 1, F1: 2, F2: (-311 > 818)}
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS4{F0: 1}
                        _ = v20
                    }
                default:
                    v6 = (v6 / v6)
                }
                _ = v16
            }
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            if false {
                var v22 uint16 = 45
                var v23 uint8 = 1
                var v24 uint16 = (v22 << ((2) & 63))
                _ = v22
                _ = v23
                _ = v24
            } else {
                v2 -= v2
            }
            var v25 int32 = -36
            var v26 int16 = 1
            _ = fi21
            _ = v25
            _ = v26
        }
        _ = v12
    }
    v2 = (((v2 * 1e10) + (v2 - (v2 - float32((v6 * v6))))) * 0.1)
    var v27 string = "foo bar"
    if false {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (32767 >= -772) {
                fmt.Println(v27, v2, v6)
                v6 -= (((0.0 * 0.0) / v6) - v6)
                _ = len(v3)
            }
        } else {
            var v28 uint32 = 1
            _ = v28
        }
    }
    var v29 uint = 128
    for fi30 := 0; fi30 < 13; fi30++ {
        switch "" {
        case "hello":
            var v31 uint32 = 2
            _ = v31
        default:
            if (!(1 <= 436)) {
                v27 = (((v27 + v27) + v27) + v27)
            }
        }
        v6 = 0.1
        _ = fi30
    }
    var v32 float64 = v6
    var v33 uint = 18446744073709551615
    v34 := GenS6{F0: false, F1: 239}
    if true {
        if (!false) {
            switch v27 {
            case "foo bar":
                v29 -= v33
            default:
                v32 = (3.14159 * v32)
            }
            switch v27 {
            case "\n\t":
                var v35 int16 = -32768
                _ = v35
            default:
                for fi36 := 0; fi36 < 18; fi36++ {
                    v37 := GenS3{F0: 455, F1: "foo bar", F2: 939}
                    v38 := GenS7{F0: 801}
                    var v39 int = fi36
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v39
                }
            }
        }
    } else {
        var v40 float32 = v2
        switch v27 {
        case "foo bar":
            switch v27 {
            case "a":
                switch ("\n\t" + v27) {
                case "":
                    for fi41 := 0; fi41 < 9; fi41++ {
                        v32 += (-2.25 / v32)
                        var v42 int = (-(fi41))
                        _ = fi41
                        _ = v42
                    }
                default:
                    for fi43 := 0; fi43 < 13; fi43++ {
                        v44 := GenS8{F0: (65535 != uint32(v29)), F1: 1000, F2: 244}
                        _ = fi43
                        _ = v44
                    }
                }
            case "":
                if false {
                    for fi45 := 0; fi45 < 10; fi45++ {
                        fmt.Println(v33, v6, v32)
                        switch fi45 {
                        case -9223372036854775808:
                            v2 -= v40
                        case 138:
                            v40 = 0.1
                        case 2147483647:
                            v33 = 255
                        }
                        _ = fi45
                    }
                }
            case "hello":
                switch 128 {
                case 32767:
                    _ = len(v5)
                case -1:
                    v29 = (v29 | 42)
                case 126:
                    for fi46 := 0; fi46 < 20; fi46++ {
                        v47 := GenS9{F0: 4294967295}
                        _ = fi46
                        _ = v47
                    }
                default:
                    fmt.Println(v6)
                }
            }
        case "\n\t":
            v33 = 255
        case "":
            v48 := GenS10{F0: true, F1: "\n\t"}
            _ = v48
        default:
            fmt.Println(v40)
        }
        _ = v40
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
