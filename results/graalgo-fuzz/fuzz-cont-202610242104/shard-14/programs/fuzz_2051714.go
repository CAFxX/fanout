package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 int8; F1 int16; F2 uint }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int64; F1 int16; F2 uint }
type GenS7 struct { F0 uint16; F1 string; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int64; F1 uint64 }

func main() {
    var v1 int = 9223372036854775807
    var v2 float32 = 1e10
    v3 := []string{"a", ("" + "hello"), ("\n\t" + "\n\t")}
    var v4 uint16 = 1
    fmt.Println(v2, v4, v4)
    fmt.Println(v4, v4)
    v2 = v2
    var v5 uint16 = 1
    switch (v1 << ((42) & 63)) {
    case -9223372036854775808:
        fmt.Println(v5, v1, v4)
    case -363:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    case 42:
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 bool = true
            var v8 uint64 = 197
            if v7 {
                if (float64(v2) > 3.14159) {
                    var v9 uint8 = 128
                    _ = v9
                } else {
                    if true {
                        if v7 {
                            v10 := GenS1{F0: 990}
                            var v11 bool = (514 > 127)
                            _ = v10
                            _ = v11
                        } else {
                            var v12 int16 = -151
                            _ = v12
                        }
                        var v13 bool = (!v7)
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 506}
                        var v15 float64 = ((float64(1.5) + float64(v2)) + 1.5)
                        _ = v14
                        _ = v15
                    }
                    var v16 int8 = -128
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v16 = -121
                        if (!(0 == 238)) {
                            var v18 string = ("\n\t" + ("" + "\n\t"))
                            var v19 bool = v7
                            _ = v18
                            _ = v19
                        } else {
                            if (4294967295 == 2) {
                                switch (fi6 / (((fi6 ^ (fi17 << ((2) & 63)))) | 1)) {
                                case 2147483647:
                                    if true {
                                        switch fi17 {
                                        case -1:
                                            var v20 bool = true
                                            _ = v20
                                        case -372:
                                            var v21 int8 = v16
                                            _ = v21
                                        case -867:
                                            v16 += 1
                                        default:
                                            v22 := GenS1{F0: 2147483647}
                                            _ = v22
                                        }
                                    } else {
                                        v23 := GenS2{F0: "\n\t", F1: (v16 * (v16 / ((58) | 1))), F2: v16}
                                        switch -1 {
                                        case 9223372036854775807:
                                            v1 = v1
                                        case 2147483647:
                                            var v24 uint16 = 42
                                            _ = v24
                                        case 42:
                                            switch "a" {
                                            case "a":
                                                var v25 int = (-(fi6))
                                                _ = v25
                                            default:
                                                v16 = (v16 + (v16 / ((v16) | 1)))
                                            }
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        switch "" {
                                        case "":
                                            if (0 < 924) {
                                                if v7 {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                fmt.Println(v16)
                                            }
                                        case "foo bar":
                                            var v26 uint32 = 42
                                            _ = v26
                                        case "\n\t":
                                            v16 = (v16 % ((62) | 1))
                                        }
                                        _ = v23
                                    }
                                case 99:
                                    v5 -= uint16(v8)
                                case -1:
                                    v4 = v4
                                }
                                if ((128 <= 128) || true) {
                                    switch -318 {
                                    case -800:
                                        fmt.Println(fi17, fi17)
                                    case 296:
                                        v27 := GenS3{F0: 65535, F1: fi6, F2: 18446744073709551615}
                                        _ = v27
                                    case -1:
                                        v28 := GenS1{F0: 42}
                                        v29 := GenS3{F0: 0, F1: (fi6 >> ((2) & 63)), F2: 18446744073709551615}
                                        _ = v28
                                        _ = v29
                                    }
                                    switch "a" {
                                    case "":
                                        var v30 uint64 = v8
                                        _ = v30
                                    case "foo bar":
                                        fmt.Println(v7, fi6, v7)
                                        fmt.Println(fi6)
                                        if ("\n\t" >= "a") {
                                            if v7 {
                                                v31 := GenS3{F0: 308, F1: fi6, F2: 2}
                                                _ = v31
                                            } else {
                                                if true {
                                                    fmt.Println(v7, v5, v8)
                                                    fmt.Println(v2, fi6)
                                                    if (306 > 710) {
                                                        fmt.Println(v2, v8)
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                } else {
                                                    v2 -= 1e10
                                                    fmt.Println(fi17, v1, v16)
                                                }
                                                if v7 {
                                                    switch fi17 {
                                                    case -1:
                                                        switch (("a" + "hello") + ("foo bar" + "a")) {
                                                        case "hello":
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                switch ("\n\t" + ("" + ("a" + ("\n\t" + "a")))) {
                                                                case "foo bar":
                                                                    v32 := GenS4{F0: int8(v5), F1: int16(v8), F2: 1}
                                                                    _ = v32
                                                                }
                                                                var v33 string = "a"
                                                                _ = v33
                                                            case "hello":
                                                                v5 = 843
                                                            case "a":
                                                                var v34 string = ("foo bar" + ("" + ((("a" + "foo bar") + ("a" + ("hello" + "hello"))) + "a")))
                                                                var v35 float64 = -2.25
                                                                v8 += 18446744073709551615
                                                                v1 = fi6
                                                                _ = v34
                                                                _ = v35
                                                            default:
                                                                switch 32768 {
                                                                case -9223372036854775808:
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v7 = v7
                                                                    case "\n\t":
                                                                        v36 := GenS5{F0: (v5 ^ (v5 % (((v4 * 269)) | 1))), F1: (v1 - (fi6 ^ 0))}
                                                                        _ = v36
                                                                    case "foo bar":
                                                                        var v37 uint64 = (v8 * 42)
                                                                        _ = v37
                                                                    }
                                                                default:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v38 := GenS6{F0: 1, F1: 42, F2: 2}
                                                            _ = v38
                                                        case "\n\t":
                                                            var v39 bool = (-46 > v16)
                                                            _ = v39
                                                        default:
                                                            var v40 bool = (!((("\n\t" + "a") + ("\n\t" + ("hello" + "hello"))) > ("foo bar" + "foo bar")))
                                                            _ = v40
                                                        }
                                                        var v41 int8 = (v16 / (((v16 >> ((955) & 63))) | 1))
                                                        _ = v41
                                                    case 128:
                                                        v7 = (383 < -1)
                                                    case 369:
                                                        if v7 {
                                                            var v42 uint = 1
                                                            _ = v42
                                                        } else {
                                                            var v43 int32 = 265
                                                            if (true && (!(!false))) {
                                                                switch "hello" {
                                                                case "a":
                                                                    var v44 int8 = v16
                                                                    _ = v44
                                                                default:
                                                                    var v45 float64 = -2.25
                                                                    _ = v45
                                                                }
                                                                v5 -= (^(v4))
                                                            } else {
                                                                var v46 float32 = 0.1
                                                                _ = v46
                                                            }
                                                            _ = v43
                                                        }
                                                    }
                                                } else {
                                                    v3 = append(v3, ("hello" + "foo bar"))
                                                    v7 = v7
                                                    var v47 int = (fi17 / (((-(fi6))) | 1))
                                                    _ = v47
                                                }
                                                v48 := GenS7{F0: (v5 + 65535), F1: "hello", F2: 1}
                                                var v49 uint8 = 128
                                                fmt.Println(v7, fi6)
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                        var v50 int32 = -2147483648
                                        _ = v50
                                    default:
                                        var v51 uint8 = 255
                                        _ = v51
                                    }
                                    v52 := GenS1{F0: 2}
                                    _ = v52
                                }
                                var v53 string = ("\n\t" + "hello")
                                _ = v53
                            } else {
                                if ((-2.25 + 0.0) <= 1.5) {
                                    switch fi17 {
                                    case -53:
                                        var v54 string = ("" + "hello")
                                        _ = v54
                                    }
                                    v7 = false
                                }
                            }
                            if false {
                                if v7 {
                                    var v55 int64 = -1
                                    var v56 float64 = 1e10
                                    _ = v55
                                    _ = v56
                                }
                            } else {
                                switch "" {
                                case "foo bar":
                                    var v57 uint16 = v5
                                    _ = v57
                                default:
                                    v58 := GenS5{F0: v4, F1: fi6}
                                    _ = v58
                                }
                            }
                            switch fi17 {
                            case 438:
                                v59 := GenS7{F0: v5, F1: "a", F2: 32767}
                                _ = v59
                            default:
                                v60 := GenS4{F0: int8(fi17), F1: 1, F2: 42}
                                _ = v60
                            }
                            var v61 bool = v7
                            _ = v61
                        }
                        var v62 uint8 = 2
                        _ = fi17
                        _ = v62
                    }
                    _ = v16
                }
                var v63 int32 = 2147483647
                _ = v63
            } else {
                v8 = v8
                for fi64 := 0; fi64 < 16; fi64++ {
                    var v65 float64 = 1e10
                    _ = fi64
                    _ = v65
                }
                fmt.Println(v2)
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    if len(v3) > 0 {
        v3[0] = ("a" + "foo bar")
    }
    for fi66 := 0; fi66 < 13; fi66++ {
        for fi67 := 0; fi67 < 1; fi67++ {
            v5 = 0
            fmt.Println(v5, v1, v1)
            v68 := GenS6{F0: 0, F1: 683, F2: 42}
            _ = fi67
            _ = v68
        }
        for fi69 := 0; fi69 < 19; fi69++ {
            v5 -= v5
            v70 := GenS8{F0: 65535, F1: v4, F2: (fi66 ^ (v1 ^ (^(fi69))))}
            var v71 uint64 = 2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi69
            _ = v70
            _ = v71
        }
        _ = fi66
    }
    v72 := GenS9{F0: v4, F1: 42}
    v3 = append(v3, "a")
    if (1 <= -9223372036854775808) {
        v73 := GenS10{F0: 42, F1: 42}
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v5, v5, v5)
        }
        _ = v73
    }
    var v74 int8 = 1
    _ = len(v3)
    fmt.Println(v5, v74)
    switch "a" {
    case "foo bar":
        v2 += 0.1
    case "a":
        var v75 int64 = int64(v5)
        _ = v75
    }
    v74 = (v74 | (-(v74)))
    for fi76 := 0; fi76 < 2; fi76++ {
        v3 = append(v3, "\n\t")
        _ = fi76
    }
    var v77 float32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v74
    _ = v77
}
