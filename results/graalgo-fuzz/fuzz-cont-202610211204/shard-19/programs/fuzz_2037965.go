package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 int16; F2 bool }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 string; F1 int16; F2 uint8 }

func main() {
    v1 := []int{-2147483648}
    var v2 int32 = 42
    var v3 string = "foo bar"
    var v4 string = "\n\t"
    var v5 int64 = 700
    var v6 uint = 0
    v3 = ((v4 + "foo bar") + (("foo bar" + v4) + (("hello" + ("a" + "hello")) + "hello")))
    v4 = (v3 + v4)
    fmt.Println(v4, v3)
    var v7 float32 = 3.14159
    if len(v1) > 0 {
        v1[0] = 128
    }
    switch -74 {
    case 2147483647:
        for fi8 := 0; fi8 < 7; fi8++ {
            switch v3 {
            case "\n\t":
                var v9 uint8 = 128
                _ = v9
            case "a":
                if true {
                    fmt.Println(v7)
                    switch fi8 {
                    case 1:
                        if len(v1) > 0 {
                            v1[0] = (fi8 * fi8)
                        }
                    case 2147483647:
                        v3 = v3
                    case -814:
                        fmt.Println(v5, v6)
                    default:
                        v7 += (((v7 + (float32(0.0) - ((v7 * v7) * (float32(1.5) - float32(-2.25))))) + v7) + v7)
                    }
                } else {
                    if ((v7 + float32(1e10)) >= 1e10) {
                        var v10 float32 = (v7 * 0.1)
                        fmt.Println(v4, v6)
                        _ = v10
                    }
                    var v11 int16 = 32767
                    _ = v11
                }
            case "foo bar":
                v12 := GenS1{F0: 65535, F1: ("\n\t" + v3)}
                _ = v12
            }
            _ = fi8
        }
    case -2147483648:
        fmt.Println(v7, v5)
    case 9223372036854775807:
        var v13 int64 = (v5 | v5)
        _ = v13
    }
    var v14 uint = v6
    if false {
        switch -2147483648 {
        case 32768:
            switch "\n\t" {
            case "hello":
                _ = len(v1)
            case "":
                v1 = append(v1, -9223372036854775808)
            }
        case -2147483648:
            v1 = append(v1, 2147483647)
        case -1:
            v3 = v3
        default:
            var v15 float32 = 0.0
            _ = v15
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            v7 += v7
            v17 := GenS2{F0: (v14 % ((745) | 1)), F1: 0, F2: false}
            v18 := GenS3{F0: 100, F1: 151}
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 int = (fi19 / ((-426) | 1))
        _ = fi19
        _ = v20
    }
    fmt.Println(v2, v5, v4)
    var v21 int16 = 266
    v7 += v7
    if (1 > 255) {
        switch -36 {
        case 32768:
            var v22 uint64 = 0
            _ = v22
        case 2147483647:
            v23 := GenS1{F0: 1, F1: v4}
            _ = v23
        case 0:
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 string = v3
                switch fi24 {
                case 9223372036854775807:
                    switch fi24 {
                    case 128:
                        fmt.Println(v3)
                    case -153:
                        if (339 != 42) {
                            var v26 uint32 = 745
                            var v27 int8 = 91
                            _ = v26
                            _ = v27
                        } else {
                            fmt.Println(v2)
                            if true {
                                for fi28 := 0; fi28 < 8; fi28++ {
                                    var v29 int64 = (v5 >> (((v6 >> ((v14) & 63))) & 63))
                                    var v30 uint8 = 2
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                v31 := GenS1{F0: 255, F1: "foo bar"}
                                v7 -= v7
                                _ = v31
                            } else {
                                v25 = v3
                                v32 := GenS4{F0: 876, F1: -2147483648}
                                _ = v32
                            }
                            v33 := GenS4{F0: 0, F1: v2}
                            _ = v33
                        }
                    }
                case -9223372036854775808:
                    if (42 == 2) {
                        fmt.Println(v21, v7)
                        for fi34 := 0; fi34 < 18; fi34++ {
                            _ = len(v1)
                            v21 = v21
                            v4 = (v4 + (v3 + v4))
                            _ = fi34
                        }
                        v2 = int32(v14)
                    } else {
                        if (255 < 255) {
                            switch v25 {
                            case "a":
                                var v35 string = v25
                                _ = v35
                            case "":
                                v36 := GenS5{F0: (v3 + ((v3 + ((("\n\t" + "\n\t") + v3) + v3)) + "foo bar")), F1: v21, F2: 128}
                                _ = v36
                            case "hello":
                                if true {
                                    var v37 uint16 = 32767
                                    _ = v37
                                } else {
                                    for fi38 := 0; fi38 < 6; fi38++ {
                                        if true {
                                            var v39 uint32 = 4294967295
                                            switch (-(fi38)) {
                                            case 1:
                                                fmt.Println(v39)
                                            default:
                                                fmt.Println(v39, v25)
                                            }
                                            fmt.Println(v5, v3, v4)
                                            _ = v39
                                        } else {
                                            var v40 float64 = (3.14159 * 0.0)
                                            if (false || false) {
                                                var v41 uint16 = 0
                                                _ = len(v1)
                                                _ = v41
                                            }
                                            _ = v40
                                        }
                                        v6 = (v14 & v6)
                                        v21 = (v21 | (v21 / ((v21) | 1)))
                                        _ = fi38
                                    }
                                    v14 -= 65535
                                    var v42 int64 = -874
                                    _ = v42
                                }
                            default:
                                var v43 int = fi24
                                _ = v43
                            }
                            v1 = append(v1, fi24)
                            if (0.1 < v7) {
                                v44 := GenS1{F0: 746, F1: v3}
                                v45 := GenS2{F0: (^(v6)), F1: -312, F2: ((4294967295 >= 42) || false)}
                                fmt.Println(v25)
                                _ = v44
                                _ = v45
                            }
                        } else {
                            v4 = v3
                            for fi46 := 0; fi46 < 10; fi46++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v1 = append(v1, 817)
                                var v47 int8 = -128
                                _ = fi46
                                _ = v47
                            }
                        }
                    }
                case 4294967295:
                    if true {
                        v14 -= v6
                        v7 = -2.25
                        fmt.Println(v14)
                    }
                default:
                    fmt.Println(v5)
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi24
                _ = v25
            }
        }
    }
    if (9223372036854775807 > 32767) {
        v5 -= -979
    }
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
