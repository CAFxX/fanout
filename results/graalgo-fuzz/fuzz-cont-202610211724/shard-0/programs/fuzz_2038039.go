package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 int32; F1 uint64; F2 uint16 }

func main() {
    var v1 int64 = -315
    v2 := make(map[int]bool)
    v3 := []string{"hello", "", ("foo bar" + "\n\t"), "\n\t"}
    var v4 int16 = 1
    var v5 uint64 = 1
    fmt.Println(v1, v1)
    switch "" {
    case "foo bar":
        _ = len(v2)
    case "\n\t":
        var v6 bool = true
        _ = v6
    default:
        if true {
            delete(v2, 4294967295)
            fmt.Println(v5, v1)
        }
    }
    if (643 != 133) {
        v5 -= 1
        switch "foo bar" {
        case "a":
            v7 := GenS1{F0: 744, F1: 0}
            _ = v7
        case "foo bar":
            switch "" {
            case "foo bar":
                if true {
                    if (!true) {
                        switch ("hello" + ("a" + "\n\t")) {
                        case "foo bar":
                            v3 = append(v3, "a")
                        }
                        if false {
                            switch ("" + "") {
                            case "hello":
                                if ((v4 * (v4 - (v4 - (v4 - v4)))) < v4) {
                                    for fi8 := 0; fi8 < 17; fi8++ {
                                        for fi9 := 0; fi9 < 11; fi9++ {
                                            v1 = v1
                                            _ = fi9
                                        }
                                        v10 := GenS2{F0: 393, F1: 996}
                                        _ = fi8
                                        _ = v10
                                    }
                                    fmt.Println(v4)
                                } else {
                                    v11 := GenS1{F0: 595, F1: (v4 * (^(v4)))}
                                    var v12 uint8 = 15
                                    _ = v11
                                    _ = v12
                                }
                            case "a":
                                for fi13 := 0; fi13 < 1; fi13++ {
                                    var v14 int64 = v1
                                    v1 += -1
                                    fmt.Println(fi13, v5)
                                    _ = fi13
                                    _ = v14
                                }
                            default:
                                _ = len(v3)
                            }
                            for fi15 := 0; fi15 < 6; fi15++ {
                                var v16 int = (fi15 & -393)
                                _ = fi15
                                _ = v16
                            }
                            fmt.Println(v1, v1)
                        }
                    } else {
                        if true {
                            v5 -= v5
                        }
                    }
                } else {
                    fmt.Println(v5, v5)
                    v5 += 448
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                }
            case "hello":
                for fi17 := 0; fi17 < 18; fi17++ {
                    v18 := GenS2{F0: 0, F1: 261}
                    v5 -= v5
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v5 += (v5 << ((0) & 63))
                        var v20 int16 = (v4 | v4)
                        var v21 uint = 42
                        v21 = (v21 << ((v21) & 63))
                        switch "" {
                        case "hello":
                            var v22 uint16 = 255
                            _ = v22
                        default:
                            var v23 string = "\n\t"
                            _ = v23
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = fi17
                    _ = v18
                }
            case "\n\t":
                for fi24 := 0; fi24 < 9; fi24++ {
                    v5 = v5
                    fmt.Println(fi24, v1, v4)
                    _ = fi24
                }
            default:
                switch ("foo bar" + ("\n\t" + ("\n\t" + "\n\t"))) {
                case "a":
                    for fi25 := 0; fi25 < 10; fi25++ {
                        switch "foo bar" {
                        case "foo bar":
                            var v26 string = "hello"
                            _ = v26
                        case "\n\t":
                            v1 = (v1 % ((v1) | 1))
                        default:
                            v4 = int16(v5)
                        }
                        if false {
                            v27 := GenS2{F0: 22, F1: 753}
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 15; fi28++ {
                                if false {
                                    if true {
                                        var v29 int16 = (v4 - v4)
                                        v5 = v5
                                        _ = v29
                                    }
                                }
                                v5 = (v5 << ((0) & 63))
                                v30 := GenS3{F0: 255}
                                _ = fi28
                                _ = v30
                            }
                            var v31 uint = 91
                            v31 += v31
                            _ = v31
                        }
                        v5 -= (v5 >> ((1) & 63))
                        _ = fi25
                    }
                case "hello":
                    if false {
                        v32 := GenS2{F0: 1, F1: 118}
                        switch "hello" {
                        case "foo bar":
                            v1 = (v1 | v1)
                        default:
                            switch "foo bar" {
                            case "":
                                v4 = (v4 + -217)
                            case "a":
                                fmt.Println(v5)
                            }
                        }
                        _ = v32
                    } else {
                        v1 = v1
                        for fi33 := 0; fi33 < 9; fi33++ {
                            _ = len(v3)
                            _ = fi33
                        }
                    }
                }
            }
        case "\n\t":
            if true {
                var v34 uint64 = v5
                fmt.Println(v34)
                fmt.Println(v5, v34, v4)
                _ = v34
            } else {
                _ = len(v3)
                if (-2147483648 > 0) {
                    v5 -= (v5 / ((18446744073709551615) | 1))
                    v1 = (v1 | v1)
                    if (v4 > -32768) {
                        for fi35 := 0; fi35 < 11; fi35++ {
                            var v36 int8 = 127
                            v2[(fi35 ^ (fi35 ^ -317))] = true
                            fmt.Println(v1, fi35)
                            _ = fi35
                            _ = v36
                        }
                        switch int(v5) {
                        case 1:
                            if (true && ((508 < 13) && false)) {
                                v1 = (v1 / (((v1 / ((v1) | 1))) | 1))
                                _ = len(v2)
                            }
                        case 32767:
                            for fi37 := 0; fi37 < 20; fi37++ {
                                var v38 uint = 901
                                _ = fi37
                                _ = v38
                            }
                        }
                    }
                }
            }
        }
        switch "a" {
        case "":
            var v39 uint8 = 27
            _ = v39
        case "a":
            if len(v3) > 0 {
                v3[0] = (((((("hello" + "") + "hello") + "a") + ("a" + "a")) + "a") + ("foo bar" + ("foo bar" + "")))
            }
        case "\n\t":
            v40 := GenS4{F0: 127}
            _ = v40
        default:
            if true {
                fmt.Println(v4, v1, v5)
                var v41 int = 0
                var v42 int32 = 42
                _ = v41
                _ = v42
            }
        }
    }
    var v43 float64 = float64(3.14159)
    for fi44 := 0; fi44 < 20; fi44++ {
        v45 := GenS2{F0: 952, F1: 1}
        var v46 uint64 = 2
        _ = fi44
        _ = v45
        _ = v46
    }
    v47 := GenS5{F0: -458}
    fmt.Println(v43)
    if (!(true && (2147483647 == 9223372036854775807))) {
        if (2 <= 494) {
            var v48 uint8 = 1
            v49 := GenS6{F0: (v5 << ((uint(v5)) & 63)), F1: 164}
            v50 := GenS7{F0: -2147483648, F1: 860, F2: 990}
            _ = v48
            _ = v49
            _ = v50
        }
        v43 -= 0.1
    }
    v4 = (v4 | v4)
    var v51 float64 = 1e10
    v43 = (1.5 - v43)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v5
    _ = v51
}
