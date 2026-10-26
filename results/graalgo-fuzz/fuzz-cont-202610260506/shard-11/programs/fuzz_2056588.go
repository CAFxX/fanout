package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int; F2 uint64 }
type GenS5 struct { F0 string; F1 int32; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 int64; F1 int32 }
func genFn0(p0 int16, p1 int16) string {
    if false {
        var v4 float32 = (1.5 + -2.25)
        var v5 int32 = 2147483647
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 bool = ((("a" > "foo bar") && true) || (0 < 2))
            v7 = (!true)
            _ = fi6
            _ = v7
        }
        _ = v4
        _ = v5
    }
    v8 := []int{}
    if ((p1 + (p0 ^ 915)) < 32767) {
        v8 = append(v8, 2147483647)
    }
    _ = p0
    _ = p1
    _ = v8
    return "hello"
}
func genFn1(p0 int8) string {
    v9 := GenS1{F0: 88, F1: (p0 / ((p0) | 1)), F2: 0}
    v10 := GenS1{F0: 99, F1: (p0 / ((p0) | 1)), F2: -32768}
    p0 += 97
    switch "a" {
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            v11 := []int{336}
            _ = v11
        }
    default:
        p0 = -12
    }
    v12 := GenS2{F0: 4294967295}
    _ = p0
    _ = v10
    _ = v12
    _ = v9
    return "a"
}

func main() {
    v1 := []int{798, 2147483647}
    var v2 float64 = 0.0
    v3 := make(map[string]bool)
    var v13 int64 = 0
    _ = len(v3)
    var v14 uint64 = 18446744073709551615
    v13 = (v13 + (v13 << ((uint(v14)) & 63)))
    var v15 uint = 110
    switch "a" {
    case "foo bar":
        var v16 int = -939
        _ = v16
    default:
        if true {
            v13 -= v13
            var v17 uint16 = 1
            switch -1 {
            case 565:
                v15 -= v15
            case -1:
                for fi18 := 0; fi18 < 8; fi18++ {
                    if false {
                        if (v17 == v17) {
                            if false {
                                var v19 uint64 = (v14 | v14)
                                _ = v19
                            } else {
                                for fi20 := 0; fi20 < 20; fi20++ {
                                    var v21 int32 = 0
                                    v22 := GenS3{F0: (!(!false))}
                                    v23 := GenS4{F0: (^(v17)), F1: int(v21), F2: 0}
                                    v2 -= (-2.25 + 0.1)
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                v24 := GenS5{F0: "hello", F1: 42, F2: -2147483648}
                                _ = v24
                            }
                            fmt.Println(v13, v17, v15)
                            v14 = v14
                        } else {
                            fmt.Println(fi18, v2, v15)
                        }
                    }
                    v2 = -2.25
                    switch "foo bar" {
                    case "":
                        v2 += 3.14159
                    }
                    _ = fi18
                }
            case -829:
                switch "\n\t" {
                case "\n\t":
                    var v25 uint32 = 65535
                    _ = v25
                case "":
                    _ = len(v1)
                default:
                    fmt.Println(v2, v14, v17)
                }
            default:
                switch "" {
                case "hello":
                    v26 := GenS6{F0: (^(v17))}
                    _ = v26
                case "a":
                    v15 += (^(v15))
                case "":
                    var v27 int16 = 1
                    _ = v27
                default:
                    fmt.Println(v14, v13)
                }
            }
            _ = v17
        }
    }
    v15 = v15
    _ = len(v1)
    switch "a" {
    case "":
        if len(v1) > 0 {
            v1[0] = 15
        }
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v28 := GenS3{F0: (("a" + ((("\n\t" + "\n\t") + "hello") + "foo bar")) > "foo bar")}
    v13 = v13
    for fi29 := 0; fi29 < 1; fi29++ {
        switch "hello" {
        case "foo bar":
            for fi30 := 0; fi30 < 2; fi30++ {
                switch (fi29 + fi30) {
                case 956:
                    switch (("a" + "hello") + ("hello" + ("hello" + "\n\t"))) {
                    case "a":
                        var v31 int64 = v13
                        _ = v31
                    case "":
                        v15 -= v15
                    case "foo bar":
                        switch fi30 {
                        case 1:
                            var v32 float32 = -2.25
                            _ = v32
                        }
                    default:
                        var v33 int32 = 0
                        _ = v33
                    }
                case -1:
                    v34 := GenS7{F0: 127}
                    _ = v34
                case -818:
                    v15 = (v15 >> ((393) & 63))
                }
                v14 += (^(v14))
                v2 = 0.0
                _ = fi30
            }
        case "":
            for fi35 := 0; fi35 < 7; fi35++ {
                v36 := GenS8{F0: ("foo bar" + "\n\t"), F1: 127}
                if false {
                    v15 += 18446744073709551615
                }
                var v37 uint8 = 126
                _ = fi35
                _ = v36
                _ = v37
            }
        default:
            v15 = (^(v15))
        }
        var v38 int = fi29
        var v39 int8 = 100
        _ = fi29
        _ = v38
        _ = v39
    }
    if (false || (("a" < "\n\t") && (!(-2147483648 > 2147483647)))) {
        delete(v3, "\n\t")
        switch -623 {
        case 42:
            for fi40 := 0; fi40 < 13; fi40++ {
                v14 = (v14 + 42)
                fmt.Println(v13)
                for fi41 := 0; fi41 < 20; fi41++ {
                    v2 = (float64((1e10 / (1.5 * 0.1))) + ((float64((0.1 / 1.5)) / v2) / v2))
                    var v42 uint = (v15 | v15)
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
        case 4294967295:
            var v43 float32 = 1e10
            _ = v43
        default:
            for fi44 := 0; fi44 < 16; fi44++ {
                switch "" {
                case "foo bar":
                    for fi45 := 0; fi45 < 20; fi45++ {
                        var v46 uint8 = uint8(v14)
                        var v47 uint64 = (v14 + v14)
                        v48 := GenS9{F0: v13, F1: 0}
                        v1 = append(v1, (fi45 << ((42) & 63)))
                        var v49 uint = (^(v15))
                        _ = fi45
                        _ = v46
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                }
                v14 += (v14 % (((v14 / (((v14 | 4294967295)) | 1))) | 1))
                v14 += (v14 ^ (v14 % ((682) | 1)))
                _ = fi44
            }
        }
        v1 = append(v1, 9223372036854775807)
    }
    fmt.Println(v13)
    var v50 uint32 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v14 -= v14
    for fi51 := 0; fi51 < 14; fi51++ {
        var v52 uint32 = v50
        _ = fi51
        _ = v52
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v50
}
