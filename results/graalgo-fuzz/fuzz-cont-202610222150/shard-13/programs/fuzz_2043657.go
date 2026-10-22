package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int; F1 uint16; F2 int }
func genFn0(p0 uint16, p1 uint) int16 {
    fmt.Println(p0)
    var v4 float32 = (float32(((-2.25 / (0.1 * 0.1)) + -2.25)) - float32((1.5 / (((0.0 + 1e10) + 1.5) - 0.0))))
    v5 := []int{-9223372036854775808, -1, 2147483647}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 215
}
func genFn1(p0 string) int64 {
    switch (p0 + (p0 + p0)) {
    case "foo bar":
        switch p0 {
        case "foo bar":
            v6 := []int{1, -9223372036854775808, 604}
            _ = v6
        case "hello":
            fmt.Println(p0)
        default:
            v7 := []int{32767}
            _ = v7
        }
    }
    v8 := []int{128, 776}
    p0 = ((p0 + "hello") + "")
    switch 9223372036854775807 {
    case 42:
        if (-128 < 1) {
            fmt.Println(p0)
        } else {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
    case 32767:
        for fi9 := 0; fi9 < 17; fi9++ {
            switch "\n\t" {
            case "":
                switch 1 {
                case 604:
                    switch "a" {
                    case "a":
                        v10 := GenS1{F0: true, F1: 100}
                        _ = v10
                    case "hello":
                        var v11 float64 = 1.5
                        _ = v11
                    case "\n\t":
                        for fi12 := 0; fi12 < 4; fi12++ {
                            switch (p0 + (p0 + p0)) {
                            case "hello":
                                if (!false) {
                                    p0 = "a"
                                    v13 := GenS1{F0: (!(18446744073709551615 >= 4294967295)), F1: 100}
                                    fmt.Println(p0, fi9, p0)
                                    _ = v13
                                } else {
                                    if false {
                                        p0 = "\n\t"
                                    }
                                    p0 = p0
                                }
                                p0 = ("foo bar" + ("\n\t" + p0))
                                fmt.Println(fi12)
                            case "a":
                                var v14 uint32 = 2147483647
                                _ = v14
                            case "":
                                switch p0 {
                                case "foo bar":
                                    p0 = p0
                                }
                            default:
                                fmt.Println(fi12)
                            }
                            _ = fi12
                        }
                    }
                case -1:
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v8 = append(v8, (fi15 | fi15))
                        _ = fi15
                    }
                default:
                    if (false && false) {
                        v16 := GenS2{F0: 769}
                        switch -377 {
                        case 32767:
                            if false {
                                fmt.Println(fi9, p0)
                                fmt.Println(fi9, fi9)
                                p0 = p0
                            }
                        case 127:
                            var v17 uint32 = 65535
                            _ = v17
                        case 369:
                            if (true && (783 != 118)) {
                                p0 = "foo bar"
                                p0 = p0
                                switch p0 {
                                case "\n\t":
                                    v18 := GenS1{F0: true, F1: 0}
                                    _ = v18
                                case "foo bar":
                                    v19 := GenS2{F0: 76}
                                    _ = v19
                                case "":
                                    fmt.Println(fi9, fi9)
                                }
                            } else {
                                switch fi9 {
                                case 128:
                                    if false {
                                        var v20 bool = true
                                        _ = v20
                                    }
                                case 32767:
                                    var v21 uint8 = 127
                                    _ = v21
                                default:
                                    if (2147483647 <= -1) {
                                        var v22 int16 = 42
                                        if (127 >= 42) {
                                            if len(v8) > 0 {
                                                v8[0] = (fi9 * 32767)
                                            }
                                        }
                                        fmt.Println(fi9)
                                        _ = v22
                                    }
                                }
                                fmt.Println(p0, fi9)
                                fmt.Println(fi9, fi9)
                            }
                        default:
                            if len(v8) > 0 {
                                v8[0] = fi9
                            }
                        }
                        if len(v8) > 0 {
                            v8[0] = fi9
                        }
                        _ = v16
                    } else {
                        fmt.Println(fi9)
                        v8 = append(v8, fi9)
                        if ("a" < ("\n\t" + "foo bar")) {
                            for fi23 := 0; fi23 < 16; fi23++ {
                                fmt.Println(fi23)
                                p0 = "foo bar"
                                v8 = append(v8, -2147483648)
                                _ = fi23
                            }
                            v24 := GenS2{F0: 42}
                            p0 = "a"
                            _ = v24
                        }
                    }
                }
            case "hello":
                if false {
                    switch p0 {
                    case "\n\t":
                        fmt.Println(p0, p0)
                    }
                    if len(v8) > 0 {
                        v8[0] = -1
                    }
                    _ = len(v8)
                } else {
                    if false {
                        p0 = (p0 + p0)
                        _ = len(v8)
                        fmt.Println(fi9, p0)
                    }
                    if true {
                        var v25 float64 = 3.14159
                        _ = v25
                    }
                }
            }
            _ = fi9
        }
    }
    v8 = append(v8, 127)
    _ = p0
    _ = v8
    return 823
}

func main() {
    var v1 uint = 65535
    v2 := make(map[uint32]string)
    v3 := []uint32{2, 2147483648, 42, 42}
    for fi26 := 0; fi26 < 12; fi26++ {
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 uint32 = 1
            _ = fi27
            _ = v28
        }
        var v29 int = (fi26 % (((fi26 - fi26)) | 1))
        var v30 int16 = -599
        _ = fi26
        _ = v29
        _ = v30
    }
    if false {
        v31 := GenS2{F0: 2}
        var v32 uint32 = 2147483647
        v33 := GenS3{F0: (32767 == genFn0(308, 0))}
        _ = v31
        _ = v32
        _ = v33
    }
    var v34 float32 = -2.25
    var v35 float64 = 3.14159
    v34 += v34
    for fi36 := 0; fi36 < 14; fi36++ {
        v2[1] = "foo bar"
        for fi37 := 0; fi37 < 14; fi37++ {
            if (false && true) {
                _ = len(v3)
                v34 -= v34
                v35 -= v35
                var v38 uint32 = 42
                _ = v38
            } else {
                v39 := GenS4{F0: "\n\t"}
                if true {
                    v40 := GenS2{F0: 255}
                    _ = v40
                } else {
                    v41 := GenS4{F0: "a"}
                    var v42 uint64 = 4294967295
                    v43 := GenS2{F0: 666}
                    _ = v41
                    _ = v42
                    _ = v43
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v39
            }
            _ = len(v3)
            _ = fi37
        }
        _ = fi36
    }
    v2[0] = ("" + "hello")
    v44 := GenS5{F0: -2147483648, F1: 9223372036854775807, F2: 945}
    fmt.Println(v34, v34)
    var v45 uint8 = 255
    v46 := GenS6{F0: 4294967295, F1: false}
    var v47 string = "hello"
    v48 := GenS6{F0: 1, F1: false}
    v49 := GenS7{F0: -2147483648, F1: 1, F2: 2147483647}
    v45 -= 127
    var v50 string = v47
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
