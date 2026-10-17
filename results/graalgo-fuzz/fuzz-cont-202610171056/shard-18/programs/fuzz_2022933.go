package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 bool }

func main() {
    var v1 float64 = (0.0 + 3.14159)
    var v2 int8 = 42
    var v3 int64 = 9223372036854775807
    v4 := []string{(("foo bar" + "a") + "hello"), "foo bar", (("hello" + "\n\t") + ("hello" + "\n\t"))}
    v5 := []string{"hello", "foo bar", "\n\t"}
    var v6 uint32 = 882
    switch -107 {
    case -1:
        if (v2 >= (v2 % (((v2 << ((255) & 63))) | 1))) {
            if (255 != 255) {
                fmt.Println(v3, v3)
                fmt.Println(v1, v6, v3)
            } else {
                for fi7 := 0; fi7 < 16; fi7++ {
                    fmt.Println(fi7, v2, v6)
                    _ = fi7
                }
                if (255 != 42) {
                    v6 = (v6 & 65535)
                }
                v8 := GenS1{F0: -1}
                _ = v8
            }
            if false {
                var v9 uint64 = 0
                var v10 int16 = 1
                _ = v10
                _ = v9
            }
        } else {
            v1 = (v1 + v1)
            v1 += float64((float32(0.1) - ((-2.25 / 3.14159) / -2.25)))
        }
    case 127:
        fmt.Println(v1, v2)
    default:
        for fi11 := 0; fi11 < 11; fi11++ {
            v4 = append(v4, "hello")
            _ = fi11
        }
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    fmt.Println(v2, v1, v3)
    v4 = append(v4, "foo bar")
    v2 = (v2 << ((42) & 63))
    for fi12 := 0; fi12 < 13; fi12++ {
        if false {
            fmt.Println(v2, v3)
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int64 = int64(v2)
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                v1 += 0.1
                _ = fi15
            }
        } else {
            v16 := GenS2{F0: -952}
            for fi17 := 0; fi17 < 17; fi17++ {
                switch fi12 {
                case 42:
                    switch (fi12 + 32767) {
                    case 42:
                        _ = len(v4)
                    case 651:
                        v2 = 127
                    default:
                        if (1 > 390) {
                            v18 := GenS2{F0: -32768}
                            v3 -= -549
                            _ = v18
                        } else {
                            v5 = append(v5, "hello")
                        }
                    }
                case -594:
                    v2 += v2
                case -2147483648:
                    fmt.Println(v6)
                }
                if len(v5) > 0 {
                    v5[0] = (("\n\t" + "hello") + "hello")
                }
                _ = fi17
            }
            _ = v16
        }
        _ = fi12
    }
    v2 = 1
    switch -244 {
    case -316:
        var v19 float32 = ((1e10 / ((((0.1 - 0.1) / 1.5) + 0.0) - 0.1)) * float32(v1))
        _ = v19
    case 127:
        var v20 int8 = (v2 ^ (v2 % (((-(v2))) | 1)))
        _ = v20
    case -847:
        switch "hello" {
        case "\n\t":
            v3 = v3
        case "foo bar":
            if false {
                v6 -= (v6 << ((4294967295) & 63))
                v21 := GenS3{F0: 554, F1: 982}
                _ = v21
            }
        }
    default:
        v1 = v1
    }
    v1 = v1
    if false {
        v3 += (-(v3))
        if false {
            v22 := GenS4{F0: uint(v6)}
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS5{F0: v2, F1: (v3 ^ 0)}
                if ("\n\t" >= "a") {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        fmt.Println(fi25, fi25, fi23)
                        if (!true) {
                            var v26 int = fi25
                            v26 -= (v26 << ((65535) & 63))
                            var v27 int = fi23
                            _ = v26
                            _ = v27
                        } else {
                            v3 += v3
                            if (255 < 163) {
                                fmt.Println(v6, v1, v1)
                            }
                            switch ("a" + "") {
                            case "":
                                var v28 string = "hello"
                                _ = v28
                            case "foo bar":
                                switch fi25 {
                                case 4294967295:
                                    fmt.Println(v3)
                                }
                                var v29 bool = (true && ((v2 << ((0) & 63)) > v2))
                                v3 = (v3 ^ 1)
                                _ = v29
                            case "\n\t":
                                v6 += (v6 / ((v6) | 1))
                            }
                        }
                        _ = fi25
                    }
                    var v30 uint8 = 0
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 10; fi31++ {
                        var v32 uint16 = 241
                        _ = fi31
                        _ = v32
                    }
                    if false {
                        v33 := GenS6{F0: fi23}
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        switch ("foo bar" + ("foo bar" + "hello")) {
                        case "\n\t":
                            v34 := GenS3{F0: -778, F1: 65535}
                            _ = v34
                        case "hello":
                            v1 -= (v1 * 3.14159)
                        case "foo bar":
                            for fi35 := 0; fi35 < 10; fi35++ {
                                v4 = append(v4, "")
                                fmt.Println(fi35, v3)
                                v3 += v3
                                v36 := GenS3{F0: (fi35 / ((-966) | 1)), F1: 42}
                                v37 := GenS1{F0: (fi35 >> ((4294967295) & 63))}
                                switch (fi35 + fi35) {
                                case 336:
                                    switch -314 {
                                    case 267:
                                        v1 = 1.5
                                    default:
                                        if (!false) {
                                            var v38 int64 = (^(v3))
                                            fmt.Println(v2, fi23, v1)
                                            switch "" {
                                            case "foo bar":
                                                if len(v4) > 0 {
                                                    v4[0] = "\n\t"
                                                }
                                            default:
                                                if true {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v39 float64 = 3.14159
                                                    if (-851 > 2147483647) {
                                                        _ = len(v4)
                                                        var v40 uint16 = 809
                                                        _ = v40
                                                    } else {
                                                        var v41 int32 = 42
                                                        if (false && true) {
                                                            var v42 int = (^(fi35))
                                                            if len(v4) > 0 {
                                                                v4[0] = "foo bar"
                                                            }
                                                            _ = v42
                                                        } else {
                                                            v1 = v1
                                                        }
                                                        _ = v41
                                                    }
                                                    _ = v39
                                                }
                                            }
                                            _ = v38
                                        }
                                    }
                                case 943:
                                    v1 = v1
                                default:
                                    v3 -= v3
                                }
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        default:
                            v43 := GenS2{F0: 491}
                            _ = v43
                        }
                        _ = v33
                    } else {
                        for fi44 := 0; fi44 < 9; fi44++ {
                            switch "a" {
                            case "\n\t":
                                switch ("a" + "\n\t") {
                                case "a":
                                    _ = len(v4)
                                case "foo bar":
                                    switch "a" {
                                    case "\n\t":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "foo bar":
                                        var v45 int = fi44
                                        _ = v45
                                    default:
                                        v3 = v3
                                    }
                                case "\n\t":
                                    var v46 float32 = (3.14159 - (float32(1e10) - float32(v1)))
                                    _ = v46
                                }
                            case "hello":
                                if (111 < (v2 % ((-128) | 1))) {
                                    v6 = v6
                                    fmt.Println(v1, v6)
                                    var v47 int16 = 42
                                    _ = v47
                                }
                            case "a":
                                v5 = append(v5, ("foo bar" + ""))
                                v1 = float64(float32(3.14159))
                                var v48 uint = 622
                                var v49 uint8 = 148
                                if true {
                                    v6 += (v6 % (((v6 % (((v6 >> ((v48) & 63))) | 1))) | 1))
                                }
                                v1 = 0.0
                                _ = v48
                                _ = v49
                            default:
                                if ((!(32767 > 835)) && (v3 > int64(v2))) {
                                    var v50 uint32 = v6
                                    _ = v50
                                } else {
                                    if (!(false && ((519 != 72) && (false || true)))) {
                                        var v51 uint = 2
                                        _ = v51
                                    } else {
                                        var v52 uint16 = 42
                                        if false {
                                            switch ("\n\t" + "") {
                                            case "a":
                                                var v53 uint8 = 128
                                                _ = v53
                                            case "\n\t":
                                                fmt.Println(fi23, fi23)
                                            }
                                        }
                                        _ = v52
                                    }
                                    if true {
                                        v54 := GenS7{F0: 1}
                                        _ = len(v4)
                                        v2 = (v2 ^ v2)
                                        _ = v54
                                    } else {
                                        if (v6 < (^(v6))) {
                                            fmt.Println(fi44)
                                            fmt.Println(v1, fi23, v2)
                                            var v55 int32 = 1
                                            _ = v55
                                        } else {
                                            _ = len(v4)
                                            switch fi44 {
                                            case -154:
                                                var v56 string = ("" + "\n\t")
                                                _ = v56
                                            case 4294967295:
                                                _ = len(v4)
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                v2 = (v2 + 127)
                                            }
                                        }
                                        v1 += 1e10
                                        v1 -= -2.25
                                    }
                                }
                            }
                            switch "a" {
                            case "a":
                                if false {
                                    v2 -= v2
                                    v57 := GenS1{F0: 32767}
                                    v58 := GenS8{F0: 645, F1: 161, F2: v6}
                                    _ = v57
                                    _ = v58
                                }
                            case "\n\t":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            default:
                                v59 := GenS9{F0: (v3 | (v3 + v3)), F1: 255}
                                _ = v59
                            }
                            v3 -= -9223372036854775808
                            _ = fi44
                        }
                        v6 -= 42
                        for fi60 := 0; fi60 < 11; fi60++ {
                            v2 -= v2
                            var v61 int32 = int32(v3)
                            switch fi23 {
                            case 1:
                                var v62 uint = 715
                                _ = v62
                            case 9223372036854775807:
                                v2 = (^(v2))
                            }
                            _ = fi60
                            _ = v61
                        }
                    }
                }
                var v63 uint64 = 197
                _ = fi23
                _ = v24
                _ = v63
            }
            var v64 int64 = (v3 * v3)
            _ = v22
            _ = v64
        }
    }
    fmt.Println(v2)
    if (-9223372036854775808 >= -1) {
        v2 -= v2
        fmt.Println(v3)
        var v65 int32 = 2147483647
        _ = v65
    }
    v2 += v2
    if true {
        if true {
            var v66 int = 1
            fmt.Println(v3, v6, v1)
            _ = v66
        }
        switch 127 {
        case 1:
            v6 += v6
        case -185:
            if ((0.0 < 3.14159) && ((-1 < 354) && ("hello" >= "\n\t"))) {
                for fi67 := 0; fi67 < 5; fi67++ {
                    fmt.Println(v6, v1)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if (285 != 32767) {
                        var v68 float64 = v1
                        for fi69 := 0; fi69 < 19; fi69++ {
                            if ((42 > 42) && (1.5 < v1)) {
                                var v70 int32 = -1
                                _ = v70
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v6 = (v6 ^ (v6 * (v6 << ((530) & 63))))
                                v68 += 0.0
                            }
                            fmt.Println(fi69, v3, v6)
                            _ = len(v4)
                            _ = fi69
                        }
                        var v71 int16 = 0
                        _ = v68
                        _ = v71
                    } else {
                        var v72 int8 = (v2 << ((255) & 63))
                        v73 := GenS10{F0: true}
                        var v74 uint16 = 255
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                    _ = fi67
                }
            }
        default:
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
        }
    } else {
        var v75 int64 = v3
        v76 := GenS9{F0: v75, F1: 1}
        for fi77 := 0; fi77 < 5; fi77++ {
            v2 = v2
            v78 := GenS6{F0: (fi77 ^ 0)}
            _ = fi77
            _ = v78
        }
        _ = v75
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
