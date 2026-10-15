package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 bool }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]uint32)
    var v3 float64 = 0.0
    v4 := []string{(("foo bar" + "a") + "a"), "\n\t"}
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 6; fi6++ {
        switch "a" {
        case "":
            switch "" {
            case "foo bar":
                for fi7 := 0; fi7 < 12; fi7++ {
                    if (2147483647 <= fi6) {
                        var v8 uint32 = 341
                        v8 -= v8
                        _ = v8
                    }
                    if (!true) {
                        v4 = append(v4, "")
                        if (!(42 != 0)) {
                            if false {
                                v9 := GenS1{F0: 696, F1: -1, F2: (42 >= 751)}
                                if false {
                                    fmt.Println(fi7, fi7, v3)
                                }
                                v10 := GenS1{F0: 519, F1: 127, F2: true}
                                _ = v10
                                _ = v9
                            } else {
                                if (!(186 > 58)) {
                                    v3 -= float64((3.14159 * v1))
                                }
                                var v11 string = "foo bar"
                                v12 := GenS2{F0: (0 >= 42)}
                                var v13 uint8 = 42
                                v14 := GenS2{F0: (((v1 - v1) - 1e10) <= 1e10)}
                                v15 := GenS3{F0: -1, F1: (42 < 0)}
                                v1 = v1
                                v11 = "hello"
                                v3 = v3
                                if (!(true && (false && false))) {
                                    v16 := GenS4{F0: 1}
                                    v1 += v1
                                    fmt.Println(v1)
                                    var v17 bool = true
                                    if v17 {
                                        delete(v2, (fi6 << ((2) & 63)))
                                        fmt.Println(v3, v13)
                                    }
                                    _ = v16
                                    _ = v17
                                }
                                switch "\n\t" {
                                case "":
                                    var v18 int = fi6
                                    _ = v18
                                default:
                                    switch (fi6 % (((fi7 ^ 128)) | 1)) {
                                    case 9223372036854775807:
                                        switch "hello" {
                                        case "a":
                                            if false {
                                                v13 -= (v13 & (v13 << ((1) & 63)))
                                                if false {
                                                    v2[fi7] = 2
                                                }
                                                v19 := GenS5{F0: 0, F1: 0}
                                                _ = v19
                                            } else {
                                                if false {
                                                    fmt.Println(v3, v3, v11)
                                                    switch v11 {
                                                    case "hello":
                                                        fmt.Println(v1)
                                                    case "":
                                                        v1 = (1.5 / 1.0)
                                                    case "foo bar":
                                                        var v20 string = (v11 + v11)
                                                        _ = v20
                                                    }
                                                    v3 -= float64(v1)
                                                } else {
                                                    v21 := GenS2{F0: false}
                                                    if false {
                                                        if (!true) {
                                                            if (830 >= uint64(v13)) {
                                                                v4 = append(v4, "")
                                                                var v22 uint16 = 1
                                                                var v23 int = (fi7 * 807)
                                                                _ = v22
                                                                _ = v23
                                                            }
                                                            v11 = (v11 + "")
                                                            if true {
                                                                v3 = 3.14159
                                                                v1 = v1
                                                            }
                                                            v24 := GenS4{F0: -269}
                                                            fmt.Println(fi6, v1, v13)
                                                            v25 := GenS1{F0: 0, F1: 113, F2: true}
                                                            var v26 int = 1
                                                            switch (v26 & 2147483647) {
                                                            case 42:
                                                                if (true || (1 != 2147483648)) {
                                                                    if (127 == -1) {
                                                                        var v27 uint64 = 224
                                                                        _ = v27
                                                                    } else {
                                                                        if true {
                                                                            v28 := GenS6{F0: 255, F1: 1}
                                                                            v26 -= v26
                                                                            _ = v28
                                                                        }
                                                                    }
                                                                    v11 = ((v11 + v11) + (("foo bar" + "hello") + v11))
                                                                }
                                                            case -105:
                                                                switch "" {
                                                                case "\n\t":
                                                                    if true {
                                                                        if ((v13 - (v13 << ((42) & 63))) >= (v13 / ((42) | 1))) {
                                                                            var v29 uint = 417
                                                                            _ = v29
                                                                        }
                                                                        v3 = v3
                                                                    } else {
                                                                        fmt.Println(fi6)
                                                                        var v30 int16 = -325
                                                                        if true {
                                                                            v31 := GenS6{F0: 678, F1: 9223372036854775807}
                                                                            var v32 int16 = (-(v30))
                                                                            v1 = (v1 + 0.0)
                                                                            _ = v31
                                                                            _ = v32
                                                                        }
                                                                        _ = v30
                                                                    }
                                                                case "":
                                                                    var v33 float64 = float64((v1 - (v1 * ((1e10 - v1) + (v1 - v1)))))
                                                                    _ = v33
                                                                }
                                                            }
                                                            _ = v24
                                                            _ = v25
                                                            _ = v26
                                                        }
                                                        var v34 int16 = 343
                                                        fmt.Println(v1)
                                                        _ = v34
                                                    } else {
                                                        v11 = v11
                                                        var v35 uint32 = 0
                                                        var v36 int16 = 979
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    _ = v21
                                                }
                                            }
                                        case "hello":
                                            var v37 int16 = 1
                                            _ = v37
                                        case "":
                                            v11 = "foo bar"
                                        }
                                    case 585:
                                        var v38 int = (fi6 - (fi6 * fi7))
                                        _ = v38
                                    case 42:
                                        if (false && false) {
                                            if true {
                                                fmt.Println(fi6, v13, v11)
                                                var v39 uint64 = 42
                                                v40 := GenS5{F0: (v39 << ((0) & 63)), F1: 0}
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                    }
                                }
                                _ = v11
                                _ = v12
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                            fmt.Println(fi7, v3, v1)
                        } else {
                            switch fi6 {
                            case -917:
                                v4 = append(v4, ("foo bar" + "foo bar"))
                            case 128:
                                var v41 uint32 = 1
                                _ = v41
                            }
                            if ((!false) || true) {
                                if false {
                                    switch "foo bar" {
                                    case "":
                                        var v42 uint32 = 42
                                        _ = v42
                                    case "foo bar":
                                        delete(v5, 4294967295)
                                    }
                                } else {
                                    v43 := GenS7{F0: false}
                                    _ = v43
                                }
                            }
                        }
                    }
                    _ = fi7
                }
            case "":
                if ((0.1 * v1) != 0.0) {
                    var v44 uint32 = 65535
                    _ = len(v5)
                    _ = v44
                }
            case "hello":
                fmt.Println(v3)
            }
        case "\n\t":
            _ = len(v2)
        }
        _ = fi6
    }
    v1 += float32(-2.25)
    for fi45 := 0; fi45 < 2; fi45++ {
        v3 -= (v3 * v3)
        var v46 uint16 = 473
        _ = fi45
        _ = v46
    }
    v1 = 3.14159
    for fi47 := 0; fi47 < 4; fi47++ {
        var v48 uint64 = 18446744073709551615
        fmt.Println(v1, v3, v1)
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 18; fi49++ {
        var v50 int64 = -359
        _ = fi49
        _ = v50
    }
    v1 = (v1 / v1)
    v51 := GenS4{F0: 398}
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = 1.5
    var v52 uint64 = 20
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
