package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int; F1 int8; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int16 = -122
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    var v4 uint8 = 0
    v5 := []int{127, 4294967295, -79, 127}
    v4 -= v4
    if ("foo bar" > "") {
        switch "a" {
        case "hello":
            v1 = v1
        case "\n\t":
            var v6 bool = (false && false)
            _ = v6
        default:
            if len(v5) > 0 {
                v5[0] = -9223372036854775808
            }
        }
    }
    v1 += 1
    v1 = v1
    var v7 float64 = ((0.0 + 1e10) + (1e10 - 1e10))
    for fi8 := 0; fi8 < 17; fi8++ {
        if (((!(42 == 778)) || ((1 >= 65535) || true)) || (false || false)) {
            for fi9 := 0; fi9 < 7; fi9++ {
                v4 = v4
                if (false && (true || true)) {
                    if true {
                        if true {
                            v4 = 127
                            var v10 string = "foo bar"
                            _ = v10
                        }
                        v11 := GenS1{F0: v4, F1: -1}
                        var v12 uint32 = 34
                        _ = v11
                        _ = v12
                    }
                    switch 4294967295 {
                    case 128:
                        v13 := v3[fi9]
                        _ = v13
                    case 1:
                        if false {
                            if (!(698 > 1)) {
                                if ((fi9 != 127) || (false && true)) {
                                    switch "foo bar" {
                                    case "hello":
                                        var v14 int8 = 100
                                        _ = v14
                                    case "":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "foo bar":
                                        fmt.Println(v1, v1)
                                    }
                                    if (!true) {
                                        var v15 uint64 = 18446744073709551615
                                        switch "\n\t" {
                                        case "hello":
                                            var v16 bool = false
                                            _ = v16
                                        case "a":
                                            v7 += 1.5
                                            var v17 float64 = -2.25
                                            _ = v17
                                        default:
                                            var v18 uint32 = 65535
                                            _ = v18
                                        }
                                        v19 := GenS2{F0: 127, F1: 42, F2: -1}
                                        _ = v15
                                        _ = v19
                                    } else {
                                        v4 = 99
                                    }
                                    fmt.Println(fi9, fi8)
                                } else {
                                    fmt.Println(fi8, fi8, fi9)
                                    var v20 uint8 = 62
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v20
                                }
                                switch fi8 {
                                case 128:
                                    if true {
                                        if (1.5 < ((1e10 + float32(v7)) + 0.1)) {
                                            if (false || true) {
                                                if true {
                                                    var v21 uint8 = (v4 << ((584) & 63))
                                                    _ = v21
                                                } else {
                                                    fmt.Println(fi8, v7, v4)
                                                }
                                            } else {
                                                switch "\n\t" {
                                                case "":
                                                    v22 := GenS3{F0: fi9}
                                                    _ = v22
                                                case "\n\t":
                                                    var v23 float64 = v7
                                                    _ = v23
                                                case "hello":
                                                    if false {
                                                        v1 = v1
                                                        if (!(!true)) {
                                                            var v24 uint64 = 309
                                                            v7 = 3.14159
                                                            _ = v24
                                                        }
                                                        v25 := GenS4{F0: 70, F1: 1}
                                                        _ = v25
                                                    } else {
                                                        fmt.Println(fi9, v7, v7)
                                                        if ((("\n\t" + "a") < "hello") && true) {
                                                            fmt.Println(v1, v1)
                                                            var v26 uint32 = 1
                                                            var v27 int = (fi8 << ((0) & 63))
                                                            v1 -= (v1 << ((65535) & 63))
                                                            v28 := GenS4{F0: int8(v4), F1: 65535}
                                                            _ = v26
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                    }
                                                default:
                                                    v29 := GenS5{F0: v4}
                                                    _ = v29
                                                }
                                                v7 = v7
                                            }
                                        }
                                        if true {
                                            if true {
                                                switch "hello" {
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v30 float32 = 3.14159
                                                        _ = v30
                                                    }
                                                default:
                                                    var v31 int16 = v1
                                                    _ = v31
                                                }
                                                var v32 bool = false
                                                _ = v32
                                            }
                                            fmt.Println(v4, fi9)
                                        } else {
                                            if ((42 == 2147483647) || (!false)) {
                                                var v33 int16 = (v1 / ((0) | 1))
                                                if false {
                                                    fmt.Println(fi9, v4)
                                                    v33 += v33
                                                    fmt.Println(v1)
                                                }
                                                v5 = append(v5, fi9)
                                                _ = v33
                                            }
                                        }
                                    }
                                case -218:
                                    var v34 uint32 = 2147483647
                                    _ = v34
                                case 32768:
                                    v35 := v3[-1]
                                    _ = v35
                                }
                                fmt.Println(fi8, v7, fi8)
                            } else {
                                v4 -= v4
                            }
                        } else {
                            v36 := GenS2{F0: -643, F1: 100, F2: -2147483648}
                            v1 -= 1
                            if true {
                                var v37 uint8 = 2
                                v38 := v2[fi8]
                                v4 -= 1
                                switch "\n\t" {
                                case "":
                                    var v39 uint64 = 4294967295
                                    _ = v39
                                case "\n\t":
                                    v1 = v1
                                case "a":
                                    v40 := GenS4{F0: 1, F1: 4294967295}
                                    _ = v40
                                default:
                                    v38 = ((!(("a" >= "\n\t") && true)) && v38)
                                    var v41 int8 = int8(v4)
                                    _ = v41
                                }
                                _ = v37
                                _ = v38
                            }
                            _ = v36
                        }
                    }
                    _ = len(v3)
                }
                _ = fi9
            }
        }
        _ = fi8
    }
    var v42 bool = true
    if v42 {
        v42 = (4294967295 > 4294967295)
    }
    fmt.Println(v42)
    switch 192 {
    case 32767:
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
    default:
        for fi43 := 0; fi43 < 14; fi43++ {
            v42 = (0 > 660)
            if v42 {
                _ = len(v5)
            }
            v3[895] = (!(!v42))
            _ = fi43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v7
}
