package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint16; F1 int16; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 uint8; F1 bool; F2 uint8 }

func main() {
    var v1 uint8 = 0
    var v2 uint = 2
    v3 := []string{"\n\t", "hello", "hello", ""}
    var v4 int16 = -788
    if (62 < -90) {
        v3 = append(v3, "")
        v5 := GenS1{F0: 65535}
        _ = v5
    }
    if (!((false || true) || false)) {
        fmt.Println(v2, v1, v4)
        for fi6 := 0; fi6 < 10; fi6++ {
            if false {
                switch "" {
                case "hello":
                    var v7 int64 = 397
                    _ = v7
                case "foo bar":
                    switch "foo bar" {
                    case "hello":
                        if ((false || true) && true) {
                            v8 := GenS2{F0: (^(v2)), F1: -128, F2: "foo bar"}
                            var v9 int32 = 0
                            for fi10 := 0; fi10 < 15; fi10++ {
                                switch (fi10 * (fi10 & -522)) {
                                case 32767:
                                    v1 += 127
                                case 4294967295:
                                    v9 -= v9
                                case 42:
                                    if len(v3) > 0 {
                                        v3[0] = "hello"
                                    }
                                }
                                var v11 uint16 = 0
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi10
                                _ = v11
                            }
                            _ = v8
                            _ = v9
                        } else {
                            v4 = 32767
                        }
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                case "a":
                    fmt.Println(v1, v1, v4)
                default:
                    _ = len(v3)
                }
                var v12 int = fi6
                _ = v12
            } else {
                switch "" {
                case "hello":
                    fmt.Println(v2, v4, fi6)
                default:
                    if false {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v3 = append(v3, "hello")
                            if false {
                                v1 -= 42
                                switch (fi6 - (fi6 ^ (^(fi6)))) {
                                case -1:
                                    v1 -= (v1 * v1)
                                case -9223372036854775808:
                                    switch "\n\t" {
                                    case "foo bar":
                                        v1 = v1
                                    case "":
                                        if true {
                                            switch "foo bar" {
                                            case "hello":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case "foo bar":
                                                if (!(true || true)) {
                                                    fmt.Println(fi6)
                                                }
                                            case "":
                                                switch 916 {
                                                case 831:
                                                    var v14 int8 = 31
                                                    _ = v14
                                                }
                                            default:
                                                v15 := GenS3{F0: ("foo bar" + "\n\t")}
                                                _ = v15
                                            }
                                            var v16 int32 = 2147483647
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(fi6, v16, v2)
                                                var v17 uint8 = v1
                                                if true {
                                                    var v18 int16 = v4
                                                    if true {
                                                        v19 := GenS4{F0: "hello", F1: (-(fi13))}
                                                        _ = v19
                                                    } else {
                                                        v20 := GenS4{F0: "a", F1: -398}
                                                        v21 := GenS1{F0: 1}
                                                        switch (fi6 ^ (fi13 / (((-(fi6))) | 1))) {
                                                        case -447:
                                                            _ = len(v3)
                                                        }
                                                        _ = v20
                                                        _ = v21
                                                    }
                                                    var v22 uint32 = 1
                                                    _ = v18
                                                    _ = v22
                                                } else {
                                                    fmt.Println(v4, v4, fi13)
                                                    switch "a" {
                                                    case "a":
                                                        v23 := GenS5{F0: uint16(v17), F1: v4, F2: 128}
                                                        _ = v23
                                                    case "":
                                                        v1 = 255
                                                    case "\n\t":
                                                        var v24 uint = (v2 << (((^(v2))) & 63))
                                                        _ = v24
                                                    }
                                                    switch ("\n\t" + "a") {
                                                    case "\n\t":
                                                        var v25 float64 = 0.1
                                                        _ = v25
                                                    case "hello":
                                                        _ = len(v3)
                                                    case "a":
                                                        v26 := GenS6{F0: ((("\n\t" + "\n\t") + ("foo bar" + "")) + "")}
                                                        _ = v26
                                                    default:
                                                        var v27 int16 = -1
                                                        _ = v27
                                                    }
                                                }
                                                _ = v17
                                            default:
                                                var v28 uint64 = 500
                                                _ = v28
                                            }
                                            _ = v16
                                        }
                                        v29 := GenS2{F0: (v2 % ((978) | 1)), F1: -128, F2: (("foo bar" + ("a" + ("foo bar" + ("hello" + "")))) + "\n\t")}
                                        if true {
                                            var v30 uint32 = 297
                                            var v31 bool = false
                                            _ = len(v3)
                                            _ = v30
                                            _ = v31
                                        }
                                        if (313 > 474) {
                                            _ = len(v3)
                                            v2 = 42
                                        }
                                        _ = v29
                                    case "\n\t":
                                        v32 := GenS7{F0: 42, F1: "foo bar"}
                                        _ = v32
                                    }
                                }
                            } else {
                                fmt.Println(v4, fi6, fi13)
                                v1 += v1
                                v1 = (v1 ^ v1)
                            }
                            _ = fi13
                        }
                        switch (fi6 << ((v2) & 63)) {
                        case -9223372036854775808:
                            var v33 uint16 = 961
                            _ = v33
                        default:
                            switch ("" + "") {
                            case "foo bar":
                                fmt.Println(v4)
                            case "\n\t":
                                v34 := GenS8{F0: 351, F1: 669}
                                _ = v34
                            case "":
                                var v35 uint8 = v1
                                _ = v35
                            }
                        }
                    } else {
                        var v36 uint64 = 1
                        v4 -= v4
                        _ = v36
                    }
                }
            }
            v4 -= (v4 << (((v2 + (v2 ^ v2))) & 63))
            _ = fi6
        }
        fmt.Println(v1)
    }
    var v37 int32 = -1
    v38 := GenS8{F0: 255, F1: 42}
    var v39 int = -1
    for fi40 := 0; fi40 < 1; fi40++ {
        _ = len(v3)
        if false {
            v39 += fi40
            for fi41 := 0; fi41 < 4; fi41++ {
                if (((3.14159 + ((1e10 - -2.25) * 3.14159)) - float32(float64(1e10))) == (0.1 + 3.14159)) {
                    var v42 int16 = (v4 << (((v2 ^ (^(v2)))) & 63))
                    _ = v42
                }
                _ = fi41
            }
            if false {
                for fi43 := 0; fi43 < 15; fi43++ {
                    if (true || false) {
                        var v44 uint16 = 42
                        v2 = v2
                        _ = v44
                    } else {
                        switch "\n\t" {
                        case "a":
                            fmt.Println(fi43)
                        case "foo bar":
                            if false {
                                var v45 int64 = -1
                                v37 = v37
                                _ = v45
                            }
                        }
                        var v46 uint16 = 42
                        _ = v46
                    }
                    if true {
                        v4 -= (v4 % ((32767) | 1))
                    }
                    var v47 int8 = 5
                    _ = fi43
                    _ = v47
                }
                switch ("" + "foo bar") {
                case "a":
                    fmt.Println(fi40)
                case "hello":
                    if false {
                        var v48 float32 = ((1e10 / (1e10 - -2.25)) * (1.5 * 0.1))
                        _ = v48
                    }
                case "foo bar":
                    for fi49 := 0; fi49 < 12; fi49++ {
                        var v50 uint = v2
                        v37 += 1
                        switch "foo bar" {
                        case "foo bar":
                            switch v39 {
                            case -9223372036854775808:
                                v51 := GenS9{F0: 100, F1: (!false)}
                                _ = v51
                            case 4294967295:
                                var v52 float64 = 1.5
                                _ = v52
                            default:
                                v3 = append(v3, "")
                            }
                        case "\n\t":
                            v37 = v37
                        case "":
                            if false {
                                if true {
                                    v1 = 42
                                    if false {
                                        switch "\n\t" {
                                        case "":
                                            var v53 int32 = (v37 << ((v2) & 63))
                                            v1 = 1
                                            fmt.Println(v50, v53)
                                            _ = v53
                                        case "hello":
                                            var v54 int32 = (^(v37))
                                            _ = v54
                                        case "a":
                                            v37 -= 42
                                        default:
                                            switch ("" + "\n\t") {
                                            case "foo bar":
                                                fmt.Println(v1, v50, v50)
                                            }
                                        }
                                    } else {
                                        v2 -= v50
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                }
                                v55 := GenS6{F0: "foo bar"}
                                _ = v55
                            } else {
                                v39 = (fi49 | v39)
                                var v56 uint64 = 42
                                _ = v56
                            }
                        }
                        _ = fi49
                        _ = v50
                    }
                default:
                    v1 = v1
                }
                var v57 uint8 = 210
                _ = v57
            } else {
                v37 = v37
                v2 += (v2 % ((v2) | 1))
                v4 = v4
            }
        } else {
            v4 = (v4 / (((v4 & v4)) | 1))
            v2 = (v2 >> ((v2) & 63))
            v39 = -312
        }
        for fi58 := 0; fi58 < 14; fi58++ {
            if (-1 < 417) {
                var v59 int8 = 100
                var v60 float32 = -2.25
                var v61 int32 = (v37 * v37)
                v39 += (fi40 * 9223372036854775807)
                _ = len(v3)
                var v62 uint32 = 2147483648
                v63 := GenS10{F0: v1, F1: 32767}
                _ = v59
                _ = v60
                _ = v61
                _ = v62
                _ = v63
            }
            var v64 int64 = 9223372036854775807
            _ = fi58
            _ = v64
        }
        _ = fi40
    }
    v4 = int16(v2)
    v65 := GenS11{F0: 182, F1: false, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v65
}
