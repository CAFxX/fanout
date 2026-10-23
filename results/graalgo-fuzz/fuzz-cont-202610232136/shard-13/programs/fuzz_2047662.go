package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int64; F2 uint }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32; F1 int16; F2 uint32 }
type GenS12 struct { F0 int8; F1 int16 }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 int8; F1 int8; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 2
    v3 := make(map[int]string)
    var v4 int = 607
    var v5 uint64 = (v2 % (((v2 % (((v2 / ((827) | 1))) | 1))) | 1))
    var v6 bool = true
    var v7 int32 = -190
    var v8 uint = 0
    if false {
        _ = len(v3)
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 0, F1: "\n\t"}
            fmt.Println(v5, v7, v6)
            if v6 {
                var v11 uint32 = 4294967295
                _ = v11
            } else {
                if false {
                    switch "" {
                    case "":
                        if v6 {
                            switch "" {
                            case "foo bar":
                                for fi12 := 0; fi12 < 13; fi12++ {
                                    if true {
                                        if ((^(v8)) < (v8 + v8)) {
                                            switch 2147483647 {
                                            case 127:
                                                v13 := []int{4294967295, 127, 42}
                                                _ = v13
                                            case 2147483647:
                                                v14 := []int{(v4 ^ (fi9 & fi12))}
                                                v15 := GenS1{F0: 127, F1: "hello"}
                                                v16 := GenS1{F0: -128, F1: "\n\t"}
                                                _ = v14
                                                _ = v15
                                                _ = v16
                                            default:
                                                v17 := []int{(^(fi9))}
                                                _ = v17
                                            }
                                            switch v4 {
                                            case 143:
                                                v18 := GenS2{F0: (("" + (("a" + "\n\t") + "a")) + ("" + "hello"))}
                                                _ = v18
                                            case 42:
                                                fmt.Println(v8)
                                            }
                                            switch ("" + ("\n\t" + ("foo bar" + "a"))) {
                                            case "a":
                                                fmt.Println(fi12, v5, v8)
                                            case "":
                                                if (!true) {
                                                    var v19 int16 = 1
                                                    _ = v19
                                                }
                                            case "foo bar":
                                                var v20 uint = v8
                                                _ = v20
                                            }
                                        } else {
                                            fmt.Println(fi12)
                                            if ((1 >= 622) || (v8 < v8)) {
                                                v21 := GenS3{F0: 18446744073709551615, F1: 0}
                                                _ = v21
                                            } else {
                                                v22 := GenS1{F0: -128, F1: ""}
                                                _ = v22
                                            }
                                            fmt.Println(fi9, v2)
                                        }
                                        v23 := GenS1{F0: 42, F1: "a"}
                                        v24 := GenS4{F0: -875, F1: 255, F2: -823}
                                        _ = v23
                                        _ = v24
                                    } else {
                                        v25 := GenS5{F0: -1, F1: v7}
                                        v26 := []int{(fi12 + (v4 ^ (fi9 % ((42) | 1)))), (fi12 * v4), 32768, int(v8)}
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = fi12
                                }
                            case "a":
                                switch (v4 % (((v4 % (((v4 ^ (fi9 >> ((v8) & 63)))) | 1))) | 1)) {
                                case 0:
                                    v5 += (v5 ^ v2)
                                case 127:
                                    v2 -= (v2 & v5)
                                case -9223372036854775808:
                                    var v27 uint64 = 1
                                    _ = v27
                                default:
                                    for fi28 := 0; fi28 < 20; fi28++ {
                                        var v29 float32 = (0.1 - 1e10)
                                        v30 := GenS6{F0: 510, F1: 42}
                                        if ((v29 > 3.14159) && v6) {
                                            v31 := GenS7{F0: (("foo bar" + "foo bar") + "a")}
                                            _ = v31
                                        }
                                        _ = fi28
                                        _ = v29
                                        _ = v30
                                    }
                                }
                            case "hello":
                                v4 = 9223372036854775807
                            }
                        } else {
                            v3[127] = "hello"
                            fmt.Println(v8)
                        }
                    default:
                        for fi32 := 0; fi32 < 5; fi32++ {
                            var v33 uint32 = 2147483648
                            if false {
                                v33 += v33
                                v34 := GenS8{F0: 42, F1: 9223372036854775807, F2: 65535}
                                if (v8 <= v8) {
                                    v35 := GenS1{F0: -62, F1: ""}
                                    v3[2147483647] = (("\n\t" + ("hello" + "a")) + "\n\t")
                                    v6 = (!((!(!((1 <= 42) || v6))) || true))
                                    var v36 string = "hello"
                                    v37 := GenS2{F0: ""}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    v33 -= (v33 % ((v33) | 1))
                                    var v38 int = fi32
                                    _ = v38
                                }
                                _ = v34
                            } else {
                                v5 += (v5 | v5)
                                switch "hello" {
                                case "":
                                    switch ("foo bar" + "\n\t") {
                                    case "hello":
                                        if v6 {
                                            v8 += v8
                                            switch 127 {
                                            case 128:
                                                var v39 uint32 = 987
                                                _ = v39
                                            case 32768:
                                                var v40 string = ("foo bar" + "foo bar")
                                                _ = v40
                                            case 385:
                                                v41 := GenS2{F0: (((("foo bar" + "") + ("foo bar" + ("hello" + (("hello" + "a") + (("" + "\n\t") + "foo bar"))))) + "") + ("a" + (("\n\t" + "") + ("\n\t" + "foo bar"))))}
                                                _ = v41
                                            }
                                        } else {
                                            fmt.Println(v6)
                                            switch (("hello" + "a") + "a") {
                                            case "hello":
                                                fmt.Println(v4, v2)
                                                if v6 {
                                                    switch "\n\t" {
                                                    case "":
                                                        switch (-(fi9)) {
                                                        case 9223372036854775807:
                                                            v6 = ((-2.25 + (1e10 * -2.25)) >= 0.1)
                                                        case 4294967295:
                                                            if true {
                                                                var v42 string = ((("a" + ("\n\t" + "hello")) + "foo bar") + ("hello" + "\n\t"))
                                                                _ = v42
                                                            }
                                                        }
                                                    case "a":
                                                        switch "foo bar" {
                                                        case "":
                                                            var v43 uint64 = (v2 - 4294967295)
                                                            _ = v43
                                                        case "foo bar":
                                                            v44 := []int{v4, fi9, (fi9 & (fi32 | fi32)), (fi32 ^ fi32)}
                                                            _ = v44
                                                        case "hello":
                                                            v45 := GenS6{F0: -32768, F1: 191}
                                                            _ = v45
                                                        default:
                                                            var v46 uint16 = uint16(v33)
                                                            _ = v46
                                                        }
                                                    case "\n\t":
                                                        if v6 {
                                                            var v47 bool = v6
                                                            v47 = v47
                                                            v3[121] = "foo bar"
                                                            _ = v47
                                                        } else {
                                                            var v48 string = ("foo bar" + ((("a" + "a") + ("a" + "")) + "foo bar"))
                                                            v49 := GenS6{F0: -32768, F1: 2}
                                                            v50 := []int{fi9, v4, fi32}
                                                            var v51 uint64 = (v5 / ((2) | 1))
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    }
                                                    if ((!true) && v6) {
                                                        switch fi9 {
                                                        case 127:
                                                            v2 -= v5
                                                        case -299:
                                                            var v52 uint64 = (^(v5))
                                                            if ("\n\t" > ("\n\t" + "\n\t")) {
                                                                var v53 uint64 = v52
                                                                _ = v53
                                                            }
                                                            fmt.Println(v4)
                                                            fmt.Println(v6)
                                                            _ = v52
                                                        case 766:
                                                            if (!(-1 <= 1)) {
                                                                var v54 int64 = -171
                                                                _ = v54
                                                            } else {
                                                                var v55 float32 = -2.25
                                                                if ((fi32 + v4) < (fi32 % ((-195) | 1))) {
                                                                    v56 := GenS8{F0: 42, F1: 918, F2: 255}
                                                                    v2 -= 4294967295
                                                                    _ = len(v3)
                                                                    v57 := GenS9{F0: v8, F1: 127}
                                                                    _ = v56
                                                                    _ = v57
                                                                }
                                                                _ = v55
                                                            }
                                                        }
                                                        v8 = (v8 >> ((1) & 63))
                                                    }
                                                    v5 = v2
                                                } else {
                                                    v8 += v8
                                                    v58 := GenS10{F0: (^(v8))}
                                                    _ = v58
                                                }
                                                var v59 int32 = (v7 - -2147483648)
                                                _ = v59
                                            }
                                        }
                                    case "\n\t":
                                        fmt.Println(v7, v2, fi9)
                                    default:
                                        var v60 uint = 65535
                                        _ = v60
                                    }
                                }
                                switch "\n\t" {
                                case "hello":
                                    v61 := GenS7{F0: "a"}
                                    _ = v61
                                }
                                v4 = (fi9 - (v4 % (((^(fi9))) | 1)))
                                fmt.Println(v7, v5, fi9)
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                    switch (v4 * (v4 / (((v4 & fi9)) | 1))) {
                    case -526:
                        v2 -= 18446744073709551615
                    }
                }
            }
            _ = fi9
            _ = v10
        }
        var v62 int16 = 74
        _ = v62
    }
    var v63 uint64 = (v5 - (v2 / (((v5 & 42)) | 1)))
    switch (v4 >> (((v8 % ((984) | 1))) & 63)) {
    case 127:
        var v64 uint32 = uint32(v8)
        _ = v64
    case 32767:
        v63 += (v5 ^ v63)
    case 1:
        for fi65 := 0; fi65 < 16; fi65++ {
            fmt.Println(v7)
            for fi66 := 0; fi66 < 11; fi66++ {
                switch "a" {
                case "":
                    v8 = 0
                case "a":
                    var v67 uint64 = (v5 % (((v2 | v63)) | 1))
                    _ = v67
                case "foo bar":
                    v68 := []int{1, (v4 & (v4 & (fi66 | v4))), (fi66 / ((117) | 1)), fi66}
                    _ = v68
                }
                _ = fi66
            }
            _ = fi65
        }
    }
    for fi69 := 0; fi69 < 18; fi69++ {
        v70 := []int{}
        v4 += (fi69 * fi69)
        fmt.Println(v5, v7, v2)
        _ = fi69
        _ = v70
    }
    v71 := GenS11{F0: (^(v7)), F1: int16(v7), F2: 2}
    v72 := GenS2{F0: "\n\t"}
    var v73 float64 = ((-2.25 - 1e10) - ((0.1 + 0.1) / 3.14159))
    v74 := GenS12{F0: 97, F1: -1}
    for fi75 := 0; fi75 < 10; fi75++ {
        v63 = v2
        var v76 int64 = 145
        v77 := GenS13{F0: fi75, F1: 2}
        _ = fi75
        _ = v76
        _ = v77
    }
    fmt.Println(v2, v73)
    var v78 uint8 = 47
    switch (v4 & (v4 << ((65535) & 63))) {
    case 32768:
        v4 -= -1
    case -9223372036854775808:
        var v79 uint8 = (v78 << ((441) & 63))
        _ = v79
    case 145:
        v80 := GenS14{F0: 42, F1: 0, F2: (65535 != 0)}
        _ = v80
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    fmt.Println("v73:", v73)
    fmt.Println("v78:", v78)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v7
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v78
    _ = v8
}
