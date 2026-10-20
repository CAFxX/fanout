package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int32; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int; F2 int64 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS6 struct { F0 bool; F1 int8; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 int8; F2 uint8 }

func main() {
    v1 := []int{9223372036854775807, -2147483648, 697, 183}
    v2 := make(map[string]string)
    v3 := []int{-9223372036854775808, 669, 9223372036854775807}
    var v4 int8 = -128
    var v5 uint16 = 65535
    var v6 int = 42
    v5 = (v5 / ((32767) | 1))
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            v8 := GenS1{F0: 0}
            var v9 bool = false
            _ = v8
            _ = v9
        } else {
            v5 += (v5 ^ v5)
            for fi10 := 0; fi10 < 9; fi10++ {
                if ((v5 | v5) > v5) {
                    fmt.Println(fi7, fi7)
                    if (!true) {
                        var v11 int = fi7
                        _ = v11
                    }
                } else {
                    switch ("foo bar" + ("foo bar" + "foo bar")) {
                    case "a":
                        if false {
                            v12 := GenS1{F0: -2147483648}
                            fmt.Println(fi10, fi10)
                            v13 := GenS2{F0: 315, F1: 42, F2: v6}
                            _ = v12
                            _ = v13
                        } else {
                            v14 := v2["foo bar"]
                            switch v6 {
                            case -9223372036854775808:
                                switch v14 {
                                case "":
                                    if true {
                                        var v15 uint64 = 4294967295
                                        _ = v15
                                    } else {
                                        _ = len(v3)
                                    }
                                case "\n\t":
                                    if (((((0.1 - ((0.0 * 1.5) * -2.25)) - (-2.25 - float64((-2.25 + 0.0)))) + 0.0) * 3.14159) != (1.5 - (((float64(float32(float64((1.5 - 0.1)))) / (0.1 / float64((1e10 / ((0.0 + float32((0.1 + 1.5))) + -2.25))))) / 1.5) + (float64(float32(1.5)) + 0.0)))) {
                                        v6 += (^(fi10))
                                        v16 := GenS1{F0: 42}
                                        _ = v16
                                    } else {
                                        v4 += -99
                                        if false {
                                            var v17 uint8 = 2
                                            var v18 int = v6
                                            var v19 uint8 = v17
                                            v4 -= (v4 << ((18446744073709551615) & 63))
                                            _ = v17
                                            _ = v18
                                            _ = v19
                                        } else {
                                            switch v14 {
                                            case "":
                                                v20 := GenS1{F0: -2147483648}
                                                _ = v20
                                            case "hello":
                                                v21 := v2[((v14 + v14) + ("\n\t" + (v14 + v14)))]
                                                fmt.Println(fi10, v4)
                                                var v22 int32 = -2147483648
                                                _ = v21
                                                _ = v22
                                            default:
                                                v23 := GenS3{F0: 510}
                                                _ = v23
                                            }
                                        }
                                    }
                                }
                            case 967:
                                fmt.Println(fi7)
                                v6 = (fi7 >> ((0) & 63))
                                v4 = (v4 * v4)
                            case -1:
                                switch "hello" {
                                case "":
                                    v24 := GenS1{F0: 483}
                                    _ = v24
                                case "a":
                                    v25 := GenS1{F0: -1}
                                    _ = v25
                                case "hello":
                                    v4 = 100
                                }
                            default:
                                switch fi7 {
                                case 32768:
                                    fmt.Println(v4)
                                case -1:
                                    v2["\n\t"] = (v14 + v14)
                                case -911:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    v26 := GenS2{F0: 2, F1: 2147483647, F2: fi10}
                                    _ = v26
                                }
                            }
                            v2[("a" + "a")] = (v14 + ("a" + "hello"))
                            _ = v14
                        }
                    }
                    var v27 uint32 = 756
                    switch (-(v6)) {
                    case -1:
                        fmt.Println(v5)
                    case 520:
                        var v28 int16 = 0
                        _ = v28
                    case -2147483648:
                        var v29 int32 = int32(v27)
                        _ = v29
                    }
                    var v30 uint = 214
                    if (!(!((255 != 2) && (1.5 >= 1.5)))) {
                        var v31 float64 = 3.14159
                        v32 := GenS4{F0: 1000, F1: 1, F2: -707}
                        var v33 uint32 = v27
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    _ = v27
                    _ = v30
                }
                _ = fi10
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                switch ("foo bar" + "a") {
                case "hello":
                    fmt.Println(fi7, fi7)
                default:
                    if (-2.25 < 0.0) {
                        v6 += 277
                        fmt.Println(fi7)
                        fmt.Println(v6, v6, fi7)
                    }
                }
                _ = fi34
            }
        }
        _ = fi7
    }
    v5 -= v5
    if ((false && (false && (479 < 85))) && (581 < (v5 / (((v5 - (v5 - 287))) | 1)))) {
        if false {
            var v35 int32 = 1
            _ = v35
        }
    } else {
        if len(v1) > 0 {
            v1[0] = (v6 - 128)
        }
        var v36 float64 = 1.5
        _ = v36
    }
    if (-1 == -44) {
        for fi37 := 0; fi37 < 19; fi37++ {
            var v38 uint = 80
            _ = fi37
            _ = v38
        }
        var v39 string = ("foo bar" + ("\n\t" + (("\n\t" + "a") + "a")))
        v40 := GenS5{F0: 4294967295, F1: 1, F2: 42}
        _ = v39
        _ = v40
    }
    v41 := GenS6{F0: ("a" > "hello"), F1: (v4 >> ((0) & 63)), F2: (v4 / ((v4) | 1))}
    v42 := GenS7{F0: 1000}
    switch ("" + "a") {
    case "a":
        v6 = v6
    default:
        v5 += 793
    }
    v43 := GenS8{F0: -94, F1: 127, F2: 128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
