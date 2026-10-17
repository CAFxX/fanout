package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 bool; F1 int64; F2 int }
type GenS6 struct { F0 string; F1 int16; F2 int8 }

func main() {
    var v1 uint16 = 694
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v4 bool = (-2.25 == 3.14159)
    v5 := []uint32{65535, 2147483648, 210}
    v6 := make(map[string]uint32)
    v4 = v4
    v7 := GenS1{F0: 127}
    var v8 int8 = 1
    if v4 {
        _ = len(v5)
    }
    if ((("a" + ((("\n\t" + "foo bar") + ("hello" + "")) + "")) < (((("hello" + "foo bar") + "") + "") + "foo bar")) && false) {
        v9 := v3[269]
        _ = v9
    }
    if (255 >= 255) {
        if true {
            v4 = (!((("hello" + "foo bar") + "hello") == "foo bar"))
        }
        var v10 int = -636
        var v11 bool = v4
        _ = v10
        _ = v11
    } else {
        v2[2147483647] = 4294967295
        v1 = v1
        fmt.Println(v8, v8, v1)
    }
    v12 := GenS2{F0: 2147483647, F1: "foo bar"}
    var v13 int32 = 42
    switch 201 {
    case 797:
        switch -6 {
        case 2147483647:
            switch "" {
            case "foo bar":
                for fi14 := 0; fi14 < 9; fi14++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v15 float64 = (-2.25 * ((-2.25 + ((-2.25 * float64(0.1)) + -2.25)) + 0.0))
                    _ = fi14
                    _ = v15
                }
            case "hello":
                v16 := GenS3{F0: v4, F1: 127}
                _ = v16
            }
        case -2147483648:
            switch -2147483648 {
            case -2147483648:
                var v17 bool = (!(255 >= 18446744073709551615))
                _ = v17
            default:
                var v18 string = ("" + ("a" + "foo bar"))
                _ = v18
            }
        }
    default:
        v19 := GenS1{F0: 0}
        _ = v19
    }
    switch "" {
    case "foo bar":
        switch ("" + "a") {
        case "a":
            v20 := GenS3{F0: v4, F1: v8}
            _ = v20
        default:
            switch int(v13) {
            case 190:
                v21 := GenS4{F0: (v1 - 754), F1: -495}
                _ = v21
            }
        }
    case "\n\t":
        if v4 {
            var v22 bool = v4
            if (!v4) {
                v4 = v22
                var v23 uint16 = (v1 ^ (v1 << ((1) & 63)))
                _ = v23
            } else {
                var v24 int8 = 0
                v13 += 1
                if (-1 >= 0) {
                    v24 -= v8
                    if v4 {
                        if (249 >= 18446744073709551615) {
                            v4 = (false && (!true))
                        } else {
                            fmt.Println(v24, v8, v24)
                            for fi25 := 0; fi25 < 15; fi25++ {
                                var v26 float32 = (((((0.1 * (-2.25 * float32(float64(((-2.25 + (0.0 * 0.0)) + 0.0))))) + 0.0) * 0.0) * 3.14159) + 0.1)
                                for fi27 := 0; fi27 < 10; fi27++ {
                                    v3[462] = fi27
                                    _ = fi27
                                }
                                _ = fi25
                                _ = v26
                            }
                        }
                    } else {
                        v24 = (-(v24))
                        switch "hello" {
                        case "":
                            switch int(v1) {
                            case -2147483648:
                                fmt.Println(v24, v8)
                            case 32767:
                                var v28 int8 = (v24 | 1)
                                _ = v28
                            case 9223372036854775807:
                                var v29 int32 = (-(v13))
                                _ = v29
                            }
                        }
                    }
                    v22 = v4
                } else {
                    v30 := GenS4{F0: (v1 << ((33) & 63)), F1: 128}
                    delete(v2, 4294967295)
                    if ((v4 || false) || true) {
                        if v4 {
                            v31 := GenS5{F0: false, F1: 902, F2: 128}
                            fmt.Println(v4, v8)
                            _ = v31
                        }
                        switch 32768 {
                        case 9223372036854775807:
                            for fi32 := 0; fi32 < 17; fi32++ {
                                var v33 int = fi32
                                if false {
                                    var v34 int16 = -1
                                    fmt.Println(v22)
                                    var v35 uint = 2
                                    _ = v34
                                    _ = v35
                                } else {
                                    _ = len(v2)
                                }
                                _ = fi32
                                _ = v33
                            }
                        case -2147483648:
                            var v36 float64 = -2.25
                            _ = v36
                        case -9223372036854775808:
                            var v37 uint64 = 440
                            _ = v37
                        }
                        if ((1e10 == (0.0 / 1.5)) && v22) {
                            fmt.Println(v4)
                        } else {
                            _ = len(v5)
                            v38 := GenS6{F0: "hello", F1: 879, F2: (v8 << ((668) & 63))}
                            if v22 {
                                v24 = v8
                            } else {
                                fmt.Println(v4)
                                var v39 uint8 = 24
                                v5 = append(v5, 2)
                                _ = v39
                            }
                            _ = v38
                        }
                    }
                    _ = v30
                }
                _ = v24
            }
            _ = v22
        }
    case "a":
        fmt.Println(v8)
    }
    var v40 uint8 = 128
    if v4 {
        fmt.Println(v13)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
