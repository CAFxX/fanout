package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int32; F2 string }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int }
func genFn0() string {
    if (2 > 2) {
        v6 := GenS1{F0: 2147483648, F1: 2147483647}
        v7 := GenS1{F0: 0, F1: 2147483647}
        var v8 uint64 = 18446744073709551615
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 4294967295, F1: 1}
    _ = v9
    return ((("hello" + "hello") + ("hello" + "a")) + (("a" + ("foo bar" + ("\n\t" + "a"))) + "\n\t"))
}

func main() {
    var v1 uint16 = 124
    var v2 int16 = 0
    v3 := []uint32{728, 0, 4294967295}
    v4 := make(map[uint32]uint32)
    var v5 float32 = -2.25
    switch "hello" {
    case "hello":
        switch -145 {
        case -9223372036854775808:
            for fi10 := 0; fi10 < 13; fi10++ {
                for fi11 := 0; fi11 < 14; fi11++ {
                    fmt.Println(fi11)
                    v5 -= (v5 * v5)
                    _ = fi11
                }
                v12 := GenS2{F0: 241}
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                _ = fi10
                _ = v12
            }
        }
    }
    v1 = v1
    switch "\n\t" {
    case "a":
        if true {
            for fi13 := 0; fi13 < 5; fi13++ {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                v5 = v5
                _ = fi13
            }
            switch ("hello" + "a") {
            case "a":
                switch (("\n\t" + "foo bar") + "") {
                case "":
                    if false {
                        v14 := GenS3{F0: "\n\t", F1: 42, F2: ("" + ((((("hello" + "\n\t") + ("" + "")) + ("a" + "foo bar")) + "\n\t") + "hello"))}
                        _ = v14
                    } else {
                        v5 += (0.1 - v5)
                    }
                case "hello":
                    if (float64(v5) <= (1.5 + ((0.0 * (1.5 + ((3.14159 + 1.5) * -2.25))) + (((1.5 * 3.14159) - 1.5) * (float64(v5) + (3.14159 + -2.25)))))) {
                        var v15 uint8 = 1
                        _ = v15
                    } else {
                        delete(v4, 2147483648)
                    }
                default:
                    if (!(4294967295 >= 2147483648)) {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            fmt.Println(v2, v1, v5)
                            if (!(true || (4294967295 < 2147483648))) {
                                fmt.Println(fi16)
                                v5 = v5
                            } else {
                                v5 += v5
                                for fi17 := 0; fi17 < 3; fi17++ {
                                    var v18 uint64 = 1
                                    fmt.Println(fi16, v5)
                                    _ = fi17
                                    _ = v18
                                }
                                v19 := GenS3{F0: "\n\t", F1: -2147483648, F2: ""}
                                _ = v19
                            }
                            _ = fi16
                        }
                        v2 += (v2 - 742)
                        v3 = append(v3, 370)
                    } else {
                        var v20 int64 = 158
                        v21 := GenS1{F0: 2, F1: 119}
                        switch -1 {
                        case 2147483647:
                            v20 = -539
                        case -462:
                            v1 += v1
                        }
                        _ = v20
                        _ = v21
                    }
                }
            case "foo bar":
                var v22 string = ("hello" + "")
                _ = v22
            default:
                for fi23 := 0; fi23 < 4; fi23++ {
                    if true {
                        v24 := GenS2{F0: (v1 >> ((4294967295) & 63))}
                        var v25 float64 = -2.25
                        _ = v24
                        _ = v25
                    }
                    _ = fi23
                }
            }
        } else {
            if false {
                var v26 int8 = -118
                _ = v26
            } else {
                var v27 int32 = -2147483648
                v28 := v4[2147483648]
                _ = v27
                _ = v28
            }
        }
    default:
        fmt.Println(v2, v5, v2)
    }
    switch ((((("\n\t" + "a") + ("hello" + ((("foo bar" + "") + ("a" + "hello")) + "foo bar"))) + "\n\t") + "\n\t") + ("foo bar" + ("" + "a"))) {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            v29 := GenS1{F0: 2147483647, F1: 611}
            _ = v29
        case "":
            switch 1 {
            case -2147483648:
                v30 := GenS4{F0: 325, F1: 801}
                _ = v30
            case 1:
                if true {
                    var v31 string = ("a" + "a")
                    var v32 int64 = 0
                    v5 += 0.0
                    _ = v31
                    _ = v32
                } else {
                    v5 -= 1.5
                }
            case -9223372036854775808:
                v4[1] = 4294967295
            default:
                switch "hello" {
                case "hello":
                    for fi33 := 0; fi33 < 14; fi33++ {
                        var v34 int = (^(fi33))
                        switch (v34 << ((7) & 63)) {
                        case 621:
                            for fi35 := 0; fi35 < 11; fi35++ {
                                switch ("foo bar" + (("" + "foo bar") + (("" + (("hello" + "a") + "a")) + ("a" + "a")))) {
                                case "foo bar":
                                    if true {
                                        switch ("" + "\n\t") {
                                        case "":
                                            v36 := GenS1{F0: 2147483647, F1: 1}
                                            _ = v36
                                        default:
                                            _ = len(v4)
                                        }
                                        v37 := GenS5{F0: 1, F1: -173}
                                        fmt.Println(v1, fi33, v2)
                                        _ = v37
                                    }
                                case "hello":
                                    var v38 int64 = -9223372036854775808
                                    _ = v38
                                default:
                                    var v39 float32 = v5
                                    _ = v39
                                }
                                if (42 > 0) {
                                    v40 := GenS2{F0: v1}
                                    var v41 bool = (18446744073709551615 < 375)
                                    switch ("foo bar" + "foo bar") {
                                    case "":
                                        v42 := GenS3{F0: ("a" + "a"), F1: 1, F2: (("foo bar" + "foo bar") + "hello")}
                                        _ = v42
                                    default:
                                        var v43 uint16 = 2
                                        _ = v43
                                    }
                                    _ = v40
                                    _ = v41
                                }
                                var v44 float64 = (-2.25 + 1e10)
                                if (505 == 2) {
                                    fmt.Println(v34, v2, fi35)
                                    var v45 uint64 = 668
                                    var v46 uint16 = 42
                                    _ = v45
                                    _ = v46
                                } else {
                                    v34 -= (fi35 / ((fi33) | 1))
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi35
                                _ = v44
                            }
                        case 2147483647:
                            v2 += 1000
                        }
                        _ = fi33
                        _ = v34
                    }
                case "\n\t":
                    v3 = append(v3, 2)
                case "foo bar":
                    for fi47 := 0; fi47 < 13; fi47++ {
                        var v48 int64 = -9223372036854775808
                        v49 := GenS2{F0: (v1 - (v1 - v1))}
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                default:
                    _ = len(v4)
                }
            }
        case "a":
            switch ("foo bar" + (("" + ("\n\t" + "foo bar")) + "hello")) {
            case "":
                switch "\n\t" {
                case "":
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                }
            case "hello":
                if true {
                    if (-1 > 632) {
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                    }
                    v4[65535] = 88
                } else {
                    var v50 float64 = 1e10
                    v51 := GenS1{F0: 2147483647, F1: uint32(v1)}
                    _ = v50
                    _ = v51
                }
            default:
                for fi52 := 0; fi52 < 11; fi52++ {
                    v53 := GenS6{F0: 2, F1: v1, F2: fi52}
                    fmt.Println(v1, v5, v5)
                    _ = fi52
                    _ = v53
                }
            }
        }
    }
    v2 = (^(v2))
    var v54 int = 616
    v2 = (v2 / ((-928) | 1))
    v55 := GenS4{F0: 2, F1: int64(v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
}
