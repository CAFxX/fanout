package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 bool; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 18446744073709551615
    v3 := make(map[int]int)
    v4 := []uint32{}
    v2 = v2
    switch (("a" + "a") + "hello") {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            if (!(true && (!(1 >= 255)))) {
                if false {
                    switch -418 {
                    case 128:
                        delete(v3, 42)
                    case 42:
                        v2 += v2
                    case 2147483647:
                        var v5 int32 = -557
                        _ = v5
                    default:
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                    }
                }
                var v6 uint64 = 42
                _ = v6
            } else {
                v4 = append(v4, 0)
                v7 := GenS1{F0: 255, F1: -2147483648, F2: v2}
                v2 = (v2 - 1)
                _ = v7
            }
        case "a":
            for fi8 := 0; fi8 < 5; fi8++ {
                switch (-(fi8)) {
                case 1:
                    v2 += v2
                case 32767:
                    fmt.Println(fi8, fi8)
                case 198:
                    v2 = v2
                default:
                    fmt.Println(v2)
                }
                if true {
                    switch fi8 {
                    case 2147483647:
                        v2 -= (^(v2))
                    case 42:
                        v9 := GenS1{F0: 93, F1: 42, F2: v2}
                        _ = v9
                    case 0:
                        var v10 float32 = 3.14159
                        _ = v10
                    default:
                        v11 := GenS2{F0: 207, F1: -1}
                        _ = v11
                    }
                    v2 = (v2 / (((v2 + v2)) | 1))
                }
                v2 += (^(v2))
                _ = fi8
            }
        default:
            switch (("\n\t" + ((("foo bar" + "hello") + ("\n\t" + "hello")) + "")) + ("hello" + "")) {
            case "a":
                _ = len(v4)
            case "foo bar":
                if (-1 != -1) {
                    switch "hello" {
                    case "":
                        var v12 int8 = 42
                        _ = v12
                    case "a":
                        if ((false || (!(("a" != "hello") || (true || true)))) || false) {
                            v2 = v2
                        } else {
                            switch "\n\t" {
                            case "a":
                                if true {
                                    v13 := GenS2{F0: 2, F1: -32768}
                                    fmt.Println(v2)
                                    _ = v13
                                }
                            default:
                                v2 -= v2
                            }
                            var v14 uint8 = 128
                            var v15 string = (("\n\t" + (((("\n\t" + ("" + "a")) + "foo bar") + "\n\t") + ("hello" + ("hello" + "a")))) + (("a" + "") + ("\n\t" + "")))
                            _ = v14
                            _ = v15
                        }
                    case "hello":
                        var v16 int = 1
                        _ = v16
                    default:
                        if (false || (-1 == 42)) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            for fi17 := 0; fi17 < 4; fi17++ {
                                v18 := GenS1{F0: 255, F1: 42, F2: v2}
                                _ = fi17
                                _ = v18
                            }
                        }
                    }
                    v2 += (v2 / ((v2) | 1))
                    var v19 uint64 = 0
                    _ = v19
                }
            }
        }
    case "hello":
        switch (("hello" + ((("\n\t" + "hello") + "hello") + "a")) + "a") {
        case "foo bar":
            switch "" {
            case "hello":
                for fi20 := 0; fi20 < 4; fi20++ {
                    v2 += (^(v2))
                    _ = fi20
                }
            default:
                if (false && false) {
                    v2 = v2
                    fmt.Println(v2)
                }
            }
        }
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2)
    }
    v2 = v2
    v3[878] = 1
    var v21 uint32 = 0
    fmt.Println(v21, v21)
    var v22 uint32 = (v21 | v21)
    v2 = (v2 * v2)
    v23 := GenS2{F0: 319, F1: -32768}
    var v24 uint16 = 42
    switch 32768 {
    case 32767:
        fmt.Println(v21)
    default:
        var v25 bool = (false || true)
        _ = v25
    }
    v26 := GenS3{F0: 0, F1: 9223372036854775807, F2: 32767}
    switch "hello" {
    case "\n\t":
        v27 := v3[9223372036854775807]
        _ = v27
    case "a":
        fmt.Println(v22)
    case "hello":
        var v28 float64 = -2.25
        _ = v28
    default:
        fmt.Println(v21, v22, v2)
    }
    if ((845 == -344) && true) {
        _ = len(v1)
        var v29 int16 = 34
        v1[uint32(v24)] = 416
        _ = v29
    } else {
        for fi30 := 0; fi30 < 10; fi30++ {
            v2 -= v2
            var v31 int16 = -36
            _ = fi30
            _ = v31
        }
        if (51 > 255) {
            v32 := GenS3{F0: 0, F1: -55, F2: -35}
            fmt.Println(v2)
            if true {
                var v33 int64 = 0
                var v34 uint16 = (v24 >> (((v2 ^ (^(v2)))) & 63))
                var v35 int16 = 32767
                _ = v33
                _ = v34
                _ = v35
            } else {
                switch "hello" {
                case "a":
                    if (true && false) {
                        v22 = 42
                        v36 := GenS4{F0: 18446744073709551615, F1: v24}
                        _ = v36
                    } else {
                        v24 = 32767
                        var v37 int = 9223372036854775807
                        for fi38 := 0; fi38 < 14; fi38++ {
                            v39 := GenS5{F0: 2, F1: true, F2: false}
                            if false {
                                switch fi38 {
                                case 42:
                                    var v40 int16 = 32767
                                    _ = v40
                                case 288:
                                    for fi41 := 0; fi41 < 6; fi41++ {
                                        if len(v4) > 0 {
                                            v4[0] = (^(v21))
                                        }
                                        _ = fi41
                                    }
                                case 955:
                                    _ = len(v4)
                                }
                                if false {
                                    fmt.Println(v22)
                                    var v42 uint8 = 1
                                    _ = v42
                                }
                                _ = len(v4)
                            }
                            _ = fi38
                            _ = v39
                        }
                        _ = v37
                    }
                case "foo bar":
                    var v43 int8 = 1
                    _ = v43
                case "hello":
                    var v44 uint = v2
                    _ = v44
                default:
                    fmt.Println(v24, v24, v22)
                }
            }
            _ = v32
        }
    }
    var v45 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v45
}
