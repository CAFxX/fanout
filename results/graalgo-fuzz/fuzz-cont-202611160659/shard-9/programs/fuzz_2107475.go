package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
type GenS3 struct { F0 bool; F1 uint8; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int; F1 uint16; F2 uint }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 int16; F1 int8 }

func main() {
    var v1 int32 = -885
    v2 := []string{("hello" + "hello"), "hello", ""}
    var v3 int = 4294967295
    var v4 float64 = 0.0
    v3 += v3
    v1 = -1
    fmt.Println(v4, v3, v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(v1)
        fmt.Println(v1, fi5, v4)
        _ = fi5
    }
    v1 += (v1 / (((v1 >> ((4294967295) & 63))) | 1))
    _ = len(v2)
    var v6 uint32 = 65535
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint64 = 18446744073709551615
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: (v8 | (^(v8))), F1: -37}
            v1 = v1
            switch "a" {
            case "a":
                v11 := GenS2{F0: 128, F1: (-1 > 978), F2: ("" + (((("foo bar" + "\n\t") + ("a" + "a")) + "") + ("\n\t" + ("a" + ("hello" + "hello")))))}
                _ = v11
            case "\n\t":
                v8 += (v8 * (v8 | 42))
            case "":
                if false {
                    v12 := GenS1{F0: (v8 + v8), F1: -9223372036854775808}
                    v6 = (v6 << ((4294967295) & 63))
                    if true {
                        switch v3 {
                        case 128:
                            if true {
                                if (true && false) {
                                    v13 := GenS3{F0: false, F1: 2, F2: 42}
                                    _ = v13
                                } else {
                                    v6 = 42
                                }
                                if true {
                                    v6 = (v6 % (((v6 % ((2147483647) | 1))) | 1))
                                } else {
                                    switch ("a" + "foo bar") {
                                    case "\n\t":
                                        var v14 int16 = 32767
                                        _ = v14
                                    case "a":
                                        v15 := GenS4{F0: 9223372036854775807}
                                        _ = v15
                                    case "hello":
                                        switch "a" {
                                        case "a":
                                            var v16 uint64 = v8
                                            _ = v16
                                        case "":
                                            var v17 int16 = int16(v6)
                                            _ = v17
                                        default:
                                            fmt.Println(v6, v6)
                                        }
                                    }
                                }
                                switch "hello" {
                                case "\n\t":
                                    var v18 uint = 255
                                    _ = v18
                                case "foo bar":
                                    fmt.Println(v6, v1, v1)
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        case 1:
                            v3 = (^(fi7))
                        case -9223372036854775808:
                            var v19 float32 = -2.25
                            fmt.Println(v8, v4)
                            var v20 uint16 = 65535
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v21 := GenS2{F0: 142, F1: false, F2: ("foo bar" + "\n\t")}
                            _ = v19
                            _ = v20
                            _ = v21
                        default:
                            if false {
                                v22 := GenS5{F0: uint32(v8), F1: 32767}
                                var v23 int = fi9
                                var v24 uint = 42
                                _ = v22
                                _ = v23
                                _ = v24
                            } else {
                                v4 += ((v4 * 0.0) - (float64(((3.14159 + 1.5) - (3.14159 / 3.14159))) * (v4 - v4)))
                            }
                        }
                        v6 += (v6 >> ((0) & 63))
                        if (v1 > 2147483647) {
                            v3 = int(v6)
                            var v25 uint64 = v8
                            _ = len(v2)
                            _ = v25
                        }
                    } else {
                        if (false || false) {
                            _ = len(v2)
                            v26 := GenS2{F0: 0, F1: false, F2: ""}
                            v8 = (^(v8))
                            _ = v26
                        }
                        switch "hello" {
                        case "hello":
                            if true {
                                switch "a" {
                                case "\n\t":
                                    v4 = float64(float32((1.5 * 1.5)))
                                }
                            } else {
                                var v27 int64 = -1
                                v1 -= v1
                                if (!(v8 >= v8)) {
                                    var v28 uint = 0
                                    if false {
                                        v2 = append(v2, ("foo bar" + (("\n\t" + "hello") + "hello")))
                                        v29 := GenS6{F0: fi9, F1: 255, F2: 255}
                                        _ = v29
                                    } else {
                                        v30 := GenS7{F0: v6, F1: (v28 >> (((v28 | (^(v28)))) & 63))}
                                        v1 += -2147483648
                                        if true {
                                            v4 = (-2.25 * v4)
                                        } else {
                                            var v31 int32 = (v1 | (v1 / (((v1 + (v1 * v1))) | 1)))
                                            v2 = append(v2, "foo bar")
                                            var v32 uint64 = (^(v8))
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v30
                                    }
                                    _ = v28
                                }
                                _ = v27
                            }
                        default:
                            v2 = append(v2, "foo bar")
                        }
                        v33 := GenS1{F0: v8, F1: -9223372036854775808}
                        v4 += v4
                        _ = v33
                    }
                    _ = len(v2)
                    if false {
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v4 = (0.0 / 1e10)
                        default:
                            switch (^(v3)) {
                            case 9223372036854775807:
                                fmt.Println(v1, v1)
                            }
                        }
                        var v34 int64 = -9223372036854775808
                        switch "hello" {
                        case "hello":
                            var v35 string = ("hello" + "a")
                            _ = v35
                        }
                        _ = v34
                    }
                    _ = v12
                } else {
                    var v36 int16 = 0
                    var v37 uint32 = 46
                    switch fi7 {
                    case 32768:
                        fmt.Println(v36)
                    default:
                        v4 = v4
                    }
                    _ = v36
                    _ = v37
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v4 = (1e10 / v4)
    v1 = -1
    switch (v3 ^ (v3 << ((1) & 63))) {
    case -1:
        var v38 float32 = 0.1
        _ = v38
    default:
        v3 = -496
    }
    v39 := GenS8{F0: 0, F1: 94}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v6
}
