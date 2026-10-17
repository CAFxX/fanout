package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int32; F2 int32 }

func main() {
    var v1 int32 = -1
    var v2 uint16 = 588
    v3 := make(map[uint32]string)
    v4 := []uint32{528}
    v5 := []int{}
    fmt.Println(v2)
    v2 = v2
    v6 := GenS1{F0: (!(-1 < 32768)), F1: (("hello" != "a") && false), F2: (^(v2))}
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 bool = false
        var v10 float64 = (3.14159 / (0.1 - 0.0))
        v3[587] = "\n\t"
        _ = fi8
        _ = v10
        _ = v9
    }
    switch "" {
    case "foo bar":
        v2 -= 0
    case "a":
        v3[4294967295] = ((("foo bar" + "a") + ("" + (("hello" + ("\n\t" + ("" + "hello"))) + "foo bar"))) + "hello")
    }
    var v11 bool = true
    fmt.Println(v7, v11, v2)
    var v12 uint8 = v7
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 float32 = (-2.25 / 1.0)
        for fi15 := 0; fi15 < 7; fi15++ {
            switch (fi15 + 1) {
            case -702:
                switch -201 {
                case 4294967295:
                    var v16 float64 = (0.0 + ((3.14159 - -2.25) - (-2.25 / 1.0)))
                    _ = v16
                default:
                    v14 = 1.5
                }
            case 42:
                if true {
                    switch 4294967295 {
                    case -1:
                        switch fi13 {
                        case 4294967295:
                            switch ("" + "hello") {
                            case "foo bar":
                                _ = len(v4)
                            case "hello":
                                switch ("\n\t" + (("a" + ((("hello" + "foo bar") + ("a" + "a")) + "\n\t")) + "a")) {
                                case "\n\t":
                                    switch -421 {
                                    case -2147483648:
                                        if ((v11 || (1000 > 42)) && (824 > 4294967295)) {
                                            switch ((("hello" + ("a" + "\n\t")) + "\n\t") + "a") {
                                            case "hello":
                                                v17 := GenS2{F0: v1, F1: 65535, F2: (^(v12))}
                                                _ = v17
                                            case "":
                                                if v11 {
                                                    v7 = 0
                                                    var v18 bool = (true || (((^(fi13)) >= 4294967295) && ((471 != 1) && (v14 < v14))))
                                                    _ = v18
                                                }
                                            case "\n\t":
                                                var v19 float64 = 1e10
                                                _ = v19
                                            default:
                                                v14 = ((1.5 * (1.5 - float32(1e10))) - float32((0.0 + -2.25)))
                                            }
                                        } else {
                                            v11 = v11
                                            v2 += v2
                                            if (2 == 48) {
                                                fmt.Println(fi13, v7, v2)
                                            }
                                            fmt.Println(v2)
                                            v20 := GenS1{F0: true, F1: (v11 && true), F2: v2}
                                            _ = v20
                                        }
                                    case -569:
                                        var v21 uint32 = 4294967295
                                        _ = v21
                                    case 127:
                                        v22 := GenS2{F0: (v1 - v1), F1: 1, F2: 255}
                                        _ = v22
                                    }
                                case "foo bar":
                                    if v11 {
                                        v12 -= (^(v12))
                                        v2 = v2
                                    }
                                case "":
                                    switch 9223372036854775807 {
                                    case -1:
                                        v1 = -1
                                    case 32767:
                                        v4 = append(v4, 42)
                                    case -2147483648:
                                        if ((!(-32768 != 0)) || (!(v7 == v7))) {
                                            v12 = 144
                                            _ = len(v5)
                                        } else {
                                            fmt.Println(fi15)
                                            _ = len(v4)
                                            if len(v4) > 0 {
                                                v4[0] = 1
                                            }
                                            v14 = v14
                                            fmt.Println(fi15)
                                        }
                                    }
                                    fmt.Println(fi15)
                                    v12 += (^(v7))
                                default:
                                    v12 = (v12 + (^(v12)))
                                }
                            case "":
                                _ = len(v3)
                            }
                        case 0:
                            var v23 int8 = 116
                            _ = v23
                        case 1:
                            delete(v3, uint32(v2))
                        }
                    case 1:
                        if v11 {
                            if v11 {
                                var v24 uint = 26
                                _ = v24
                            } else {
                                v25 := GenS3{F0: fi15}
                                v14 += ((v14 * 1.5) - v14)
                                fmt.Println(v14)
                                _ = v25
                            }
                        } else {
                            if (v12 > v12) {
                                v7 = v7
                                if (207 <= 1) {
                                    var v26 int8 = 42
                                    fmt.Println(v14)
                                    _ = v26
                                }
                            }
                            v27 := GenS2{F0: 0, F1: 1, F2: (v7 / ((2) | 1))}
                            if len(v5) > 0 {
                                v5[0] = (-(fi15))
                            }
                            _ = v27
                        }
                    case 2147483647:
                        if v11 {
                            var v28 int = fi15
                            var v29 bool = (v7 > (v7 & 0))
                            _ = v28
                            _ = v29
                        }
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    v30 := GenS4{F0: 0, F1: (v2 ^ v2)}
                    var v31 int32 = 25
                    _ = v30
                    _ = v31
                } else {
                    v11 = ("foo bar" >= "")
                    fmt.Println(v12, v12)
                }
            case -267:
                v32 := v3[65535]
                _ = v32
            default:
                v33 := GenS5{F0: 4294967295, F1: (v1 & v1), F2: v1}
                _ = v33
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    if (((v7 | 1) == (v7 / ((v7) | 1))) && ((1 < 531) && v11)) {
        fmt.Println(v1, v1, v12)
        if (18446744073709551615 != 0) {
            var v34 uint8 = (^(v12))
            v3[888] = ""
            _ = v34
        }
    }
    var v35 int = 32767
    v36 := GenS4{F0: int8(v12), F1: 783}
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 2; fi38++ {
            if ((fi38 + int(v7)) < -1) {
                var v39 uint16 = (v2 & v2)
                _ = v39
            }
            v40 := GenS1{F0: (!v11), F1: v11, F2: 32767}
            _ = fi38
            _ = v40
        }
        var v41 int = fi37
        _ = fi37
        _ = v41
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
