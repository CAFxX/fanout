package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint8; F1 bool; F2 string }
type GenS5 struct { F0 int16 }

func main() {
    v1 := []string{"a", ""}
    var v2 bool = ((100 != 1) && true)
    var v3 uint64 = 440
    var v4 float32 = (-2.25 / -2.25)
    var v5 float32 = 1.5
    var v6 float64 = -2.25
    v2 = true
    fmt.Println(v2, v3, v5)
    var v7 int16 = 32767
    if len(v1) > 0 {
        v1[0] = ("" + ("\n\t" + "\n\t"))
    }
    var v8 uint8 = 128
    for fi9 := 0; fi9 < 8; fi9++ {
        v4 = v5
        fmt.Println(v4)
        _ = fi9
    }
    v2 = v2
    switch "a" {
    case "foo bar":
        var v10 int8 = 27
        _ = v10
    case "a":
        var v11 uint16 = 2
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 8; fi12++ {
            if (65535 < 0) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                switch "foo bar" {
                case "":
                    switch fi12 {
                    case 977:
                        for fi13 := 0; fi13 < 8; fi13++ {
                            var v14 int64 = 148
                            switch "" {
                            case "hello":
                                if ((v2 || false) && v2) {
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                } else {
                                    var v15 uint8 = v8
                                    _ = v15
                                }
                            case "\n\t":
                                switch 1 {
                                case 1:
                                    if false {
                                        v16 := GenS1{F0: -77, F1: v3, F2: 4294967295}
                                        v8 = (v8 - uint8(v3))
                                        _ = v16
                                    }
                                case -1:
                                    var v17 float64 = (v6 * v6)
                                    v18 := GenS1{F0: 127, F1: 42, F2: (v3 | v3)}
                                    if v2 {
                                        if ((v7 | (-(v7))) <= -32768) {
                                            switch "foo bar" {
                                            case "a":
                                                var v19 int8 = int8(v3)
                                                _ = v19
                                            case "foo bar":
                                                var v20 string = "\n\t"
                                                _ = v20
                                            }
                                            switch "" {
                                            case "hello":
                                                v1 = append(v1, (("a" + ("hello" + "\n\t")) + "hello"))
                                                v6 -= -2.25
                                            case "\n\t":
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v21 int = 32767
                                                    _ = v21
                                                case "":
                                                    fmt.Println(v4, fi13)
                                                case "hello":
                                                    v1 = append(v1, ("foo bar" + "\n\t"))
                                                default:
                                                    v4 += v5
                                                }
                                            default:
                                                v17 = v6
                                            }
                                        } else {
                                            v2 = ("foo bar" <= "hello")
                                            if (!v2) {
                                                if v2 {
                                                    fmt.Println(v5, v5, v5)
                                                    fmt.Println(v3, v7)
                                                }
                                                v6 = float64(float32((v17 - (v17 * -2.25))))
                                            } else {
                                                v4 = v4
                                                v6 = 0.0
                                            }
                                            v22 := GenS2{F0: v14}
                                            _ = v22
                                        }
                                    }
                                    switch 4294967295 {
                                    case -321:
                                        var v23 float32 = v4
                                        _ = v23
                                    case 155:
                                        v14 = int64(fi13)
                                    default:
                                        if (!("a" != "\n\t")) {
                                            var v24 float32 = v5
                                            var v25 uint16 = 65535
                                            if (100 != 42) {
                                                v6 = float64(float32(v17))
                                                v26 := GenS2{F0: (v14 | (v14 + (v14 % ((-797) | 1))))}
                                                _ = v26
                                            } else {
                                                switch (^(fi13)) {
                                                case 1:
                                                    var v27 uint = 42
                                                    _ = v27
                                                case 127:
                                                    fmt.Println(fi12, v8, v14)
                                                default:
                                                    v24 = float32(((v6 * float64((-2.25 + 3.14159))) / v17))
                                                }
                                            }
                                            _ = v24
                                            _ = v25
                                        }
                                    }
                                    _ = v17
                                    _ = v18
                                }
                            default:
                                _ = len(v1)
                            }
                            _ = fi13
                            _ = v14
                        }
                    case 0:
                        for fi28 := 0; fi28 < 7; fi28++ {
                            v29 := GenS3{F0: int16(v8), F1: 0}
                            _ = fi28
                            _ = v29
                        }
                    case 4294967295:
                        var v30 bool = (float64(-2.25) == 0.0)
                        _ = v30
                    default:
                        _ = len(v1)
                    }
                case "hello":
                    var v31 int16 = (v7 - 0)
                    _ = v31
                default:
                    fmt.Println(v8, v3, v7)
                }
            }
            switch "foo bar" {
            case "a":
                var v32 uint16 = 2
                _ = v32
            default:
                v1 = append(v1, "foo bar")
            }
            v33 := GenS3{F0: 1, F1: 4294967295}
            _ = fi12
            _ = v33
        }
        var v34 int32 = 0
        _ = v34
    } else {
        var v35 uint32 = 2147483647
        v6 = 0.0
        v4 -= 3.14159
        _ = v35
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        if (("foo bar" + "\n\t") != ("" + (("a" + "\n\t") + (("a" + ((((("hello" + "foo bar") + "\n\t") + "foo bar") + (("foo bar" + "\n\t") + ("\n\t" + "\n\t"))) + "a")) + "\n\t")))) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v37 uint = 2
            _ = v37
        } else {
            switch "\n\t" {
            case "hello":
                if v2 {
                    v38 := GenS4{F0: v8, F1: false, F2: ""}
                    v2 = v2
                    _ = v38
                }
            case "foo bar":
                switch fi36 {
                case 10:
                    var v39 float32 = v5
                    _ = v39
                default:
                    v8 += v8
                }
            case "\n\t":
                for fi40 := 0; fi40 < 15; fi40++ {
                    v6 = v6
                    v41 := GenS5{F0: v7}
                    switch ("a" + ("hello" + "\n\t")) {
                    case "hello":
                        var v42 int = (fi40 & fi36)
                        _ = v42
                    case "\n\t":
                        var v43 uint64 = (v3 & (v3 * (^(v3))))
                        _ = v43
                    }
                    _ = fi40
                    _ = v41
                }
            }
            var v44 float32 = v5
            fmt.Println(v44, v3, v6)
            _ = v44
        }
        var v45 uint32 = 2147483647
        _ = fi36
        _ = v45
    }
    v3 -= (v3 & v3)
    var v46 float32 = ((v5 - v5) / v5)
    fmt.Println(v2)
    var v47 int16 = 0
    v48 := GenS1{F0: 1, F1: 557, F2: (v3 ^ 18446744073709551615)}
    for fi49 := 0; fi49 < 8; fi49++ {
        var v50 int = (fi49 % (((fi49 + fi49)) | 1))
        _ = fi49
        _ = v50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
