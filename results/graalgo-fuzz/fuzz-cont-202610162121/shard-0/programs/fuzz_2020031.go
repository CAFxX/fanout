package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int8; F1 int32; F2 string }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 uint8; F2 int8 }
type GenS8 struct { F0 uint; F1 uint; F2 uint32 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []string{"", "hello", ("foo bar" + "hello"), "a"}
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    v5 := []string{("\n\t" + "foo bar")}
    v6 := []int{-99, 32767, 42}
    fmt.Println(v1)
    switch "hello" {
    case "\n\t":
        v1 += (^(v1))
    default:
        if false {
            var v7 uint8 = 255
            _ = v7
        } else {
            if false {
                v1 -= v1
                v8 := GenS1{F0: false, F1: true, F2: (false && (true && ((!true) || true)))}
                _ = v8
            }
            for fi9 := 0; fi9 < 17; fi9++ {
                v6 = append(v6, fi9)
                _ = fi9
            }
        }
    }
    var v10 int = 128
    v11 := GenS2{F0: 87, F1: (1 <= 2)}
    var v12 float32 = (-2.25 * 3.14159)
    if ((801 >= 24) || ((3 <= 1) && false)) {
        v13 := GenS3{F0: 127, F1: 2147483647, F2: "foo bar"}
        if true {
            switch "\n\t" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "hello":
                for fi14 := 0; fi14 < 8; fi14++ {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        switch "\n\t" {
                        case "a":
                            fmt.Println(fi14)
                            v16 := GenS3{F0: -122, F1: 2147483647, F2: "hello"}
                            switch "foo bar" {
                            case "a":
                                var v17 uint32 = 664
                                _ = v17
                            default:
                                var v18 int64 = int64(v1)
                                _ = v18
                            }
                            switch (fi15 * fi15) {
                            case 42:
                                fmt.Println(v10)
                            case -9223372036854775808:
                                v19 := GenS3{F0: 1, F1: -795, F2: ("" + "\n\t")}
                                _ = v19
                            }
                            _ = v16
                        case "hello":
                            v10 -= -1
                        case "":
                            v10 = (fi15 + v10)
                        default:
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                        }
                        var v20 int = fi14
                        fmt.Println(v1, v20, v10)
                        v21 := GenS3{F0: -125, F1: -1, F2: ""}
                        _ = fi15
                        _ = v20
                        _ = v21
                    }
                    _ = fi14
                }
            default:
                switch (v10 ^ (^(v10))) {
                case 128:
                    fmt.Println(v10)
                case 0:
                    var v22 uint8 = 2
                    _ = v22
                default:
                    v23 := GenS3{F0: -30, F1: 0, F2: "a"}
                    _ = v23
                }
            }
            v24 := GenS2{F0: 42, F1: true}
            if true {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v26 := GenS4{F0: 4294967295, F1: 2}
                    _ = fi25
                    _ = v26
                }
                var v27 int64 = 0
                v28 := GenS3{F0: 0, F1: 42, F2: "hello"}
                _ = v27
                _ = v28
            }
            _ = v24
        } else {
            var v29 int64 = -397
            var v30 bool = (v10 < v10)
            _ = v29
            _ = v30
        }
        var v31 float32 = 3.14159
        _ = v13
        _ = v31
    }
    fmt.Println(v1, v12, v12)
    if true {
        v10 = (v10 & v10)
        v12 = 0.0
        for fi32 := 0; fi32 < 11; fi32++ {
            switch "a" {
            case "a":
                v5 = append(v5, ("\n\t" + ""))
            case "":
                v10 += fi32
            }
            _ = fi32
        }
    } else {
        switch "a" {
        case "\n\t":
            fmt.Println(v1)
        }
    }
    if (v12 <= v12) {
        if (v1 == (v1 % ((18446744073709551615) | 1))) {
            v33 := GenS2{F0: 428, F1: false}
            var v34 uint64 = (^(v1))
            v1 -= v1
            _ = v33
            _ = v34
        }
        var v35 float32 = v12
        var v36 float64 = 0.1
        _ = v35
        _ = v36
    } else {
        v37 := GenS5{F0: 42}
        if true {
            v3[42] = true
            for fi38 := 0; fi38 < 6; fi38++ {
                switch (^(v10)) {
                case 42:
                    v39 := GenS4{F0: 202, F1: 32767}
                    _ = v39
                case -603:
                    if ((1.5 - float64((v12 - ((-2.25 - 0.0) - v12)))) <= (1e10 + -2.25)) {
                        var v40 float64 = (0.1 - 3.14159)
                        _ = v40
                    } else {
                        _ = len(v4)
                        var v41 string = ("\n\t" + "")
                        v12 = v12
                        _ = v41
                    }
                default:
                    var v42 int64 = 405
                    _ = v42
                }
                fmt.Println(v10, v10)
                _ = fi38
            }
            for fi43 := 0; fi43 < 6; fi43++ {
                for fi44 := 0; fi44 < 7; fi44++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v45 string = ("a" + "")
                    v10 += int(v1)
                    var v46 uint = 0
                    v3[(v10 ^ (v10 | fi44))] = true
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                switch -9223372036854775808 {
                case 32767:
                    v47 := GenS6{F0: 631}
                    _ = v47
                case 42:
                    switch (^(fi43)) {
                    case -9223372036854775808:
                        v10 = fi43
                    }
                default:
                    switch v10 {
                    case 32768:
                        v4[0] = ((v1 * (v1 + (^(v1)))) == (v1 | (v1 * v1)))
                    default:
                        switch 194 {
                        case -2147483648:
                            v48 := GenS7{F0: -32768, F1: 46, F2: -61}
                            _ = v48
                        case 2147483647:
                            for fi49 := 0; fi49 < 20; fi49++ {
                                v50 := GenS5{F0: 2147483647}
                                _ = fi49
                                _ = v50
                            }
                        }
                    }
                }
                _ = fi43
            }
        }
        v51 := GenS1{F0: ((v1 | 951) < (v1 + v1)), F1: false, F2: true}
        _ = v37
        _ = v51
    }
    v52 := GenS8{F0: 18446744073709551615, F1: 0, F2: 2}
    if ((665 == 1) && true) {
        v53 := v4[963]
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        _ = v53
    }
    v12 -= v12
    v1 += (v1 - 284)
    v12 = 1.5
    for fi54 := 0; fi54 < 14; fi54++ {
        v10 += v10
        fmt.Println(v12, fi54, fi54)
        v1 = v1
        _ = fi54
    }
    v10 -= v10
    if (!false) {
        if (!(1 >= -794)) {
            v12 = float32(float64(((0.1 + v12) * 0.0)))
            v55 := GenS1{F0: false, F1: (1.5 != 1.5), F2: false}
            v12 = ((v12 / (v12 / v12)) + v12)
            _ = v55
        }
        for fi56 := 0; fi56 < 4; fi56++ {
            var v57 int8 = -128
            v1 = (^(v1))
            v58 := v4[2147483647]
            _ = fi56
            _ = v57
            _ = v58
        }
        v59 := GenS1{F0: true, F1: false, F2: true}
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
