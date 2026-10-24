package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint }
type GenS2 struct { F0 int8; F1 int64; F2 uint }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    v5 := []string{"foo bar"}
    var v6 bool = (335 >= 0)
    switch "foo bar" {
    case "foo bar":
        switch -6 {
        case -959:
            v7 := v2["a"]
            _ = v7
        case 9223372036854775807:
            switch 9223372036854775807 {
            case 4294967295:
                v8 := GenS1{F0: 0, F1: 0, F2: 2}
                _ = v8
            default:
                v1 += (((v1 + v1) - v1) - v1)
            }
        default:
            var v9 float64 = -2.25
            _ = v9
        }
    case "a":
        var v10 int8 = 50
        _ = v10
    default:
        for fi11 := 0; fi11 < 3; fi11++ {
            v1 = v1
            _ = fi11
        }
    }
    v5 = append(v5, "a")
    _ = len(v4)
    if v6 {
        v2[("foo bar" + "")] = 4294967295
        var v12 uint = 2
        _ = v12
    } else {
        var v13 uint = 18446744073709551615
        switch "\n\t" {
        case "a":
            switch "a" {
            case "\n\t":
                v6 = (true && true)
            case "foo bar":
                v13 = 65535
            default:
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 float64 = 1.5
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := GenS2{F0: 0, F1: 0, F2: (v13 / ((v13) | 1))}
                        var v18 uint8 = 1
                        var v19 uint8 = v18
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 6; fi20++ {
                        fmt.Println(v13, fi14)
                        v1 = (v1 + (v1 * v1))
                        fmt.Println(v1, v1)
                        _ = fi20
                    }
                    _ = fi14
                    _ = v15
                }
            }
        case "\n\t":
            v21 := GenS2{F0: 127, F1: 627, F2: (^(v13))}
            _ = v21
        case "foo bar":
            var v22 string = ("a" + "foo bar")
            _ = v22
        default:
            switch 32767 {
            case 32768:
                switch 961 {
                case -589:
                    if (!true) {
                        v6 = false
                        if false {
                            if (false && v6) {
                                if ((-2.25 / (float64(1.5) - ((0.1 + 0.1) + -2.25))) >= -2.25) {
                                    v23 := GenS1{F0: -1, F1: 471, F2: v13}
                                    switch 621 {
                                    case 32768:
                                        switch "" {
                                        case "hello":
                                            var v24 string = "hello"
                                            _ = v24
                                        case "":
                                            v25 := GenS2{F0: -128, F1: -244, F2: v13}
                                            _ = v25
                                        case "a":
                                            fmt.Println(v1, v6)
                                        default:
                                            for fi26 := 0; fi26 < 14; fi26++ {
                                                if v6 {
                                                    var v27 float64 = 0.1
                                                    _ = v27
                                                } else {
                                                    var v28 bool = (1 != 255)
                                                    _ = v28
                                                }
                                                fmt.Println(v6, v1)
                                                _ = fi26
                                            }
                                        }
                                    case 42:
                                        var v29 float64 = 1.5
                                        _ = v29
                                    case 529:
                                        for fi30 := 0; fi30 < 16; fi30++ {
                                            var v31 float64 = (0.1 - float64(float32(float64((v1 - 0.0)))))
                                            v6 = false
                                            v32 := GenS1{F0: 2147483647, F1: 4294967295, F2: v13}
                                            _ = fi30
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    var v33 int64 = -1
                                    _ = v23
                                    _ = v33
                                } else {
                                    if v6 {
                                        v34 := GenS3{F0: 32767, F1: -1}
                                        _ = v34
                                    } else {
                                        if v6 {
                                            if (!(v13 >= (v13 + 42))) {
                                                v13 += v13
                                                for fi35 := 0; fi35 < 12; fi35++ {
                                                    fmt.Println(v13, v6)
                                                    fmt.Println(v1, v1, v6)
                                                    _ = fi35
                                                }
                                                v1 = 0.0
                                            } else {
                                                v5 = append(v5, "")
                                            }
                                            v2[("a" + "")] = 2
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            if v6 {
                                                if (985 < 104) {
                                                    var v36 int = 128
                                                    v1 += v1
                                                    v6 = v6
                                                    _ = v36
                                                }
                                                v37 := GenS4{F0: 4294967295}
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v1 = 0.0
                                                case "\n\t":
                                                    v1 = 1.5
                                                }
                                                _ = v37
                                            }
                                        case "hello":
                                            v2["hello"] = 65535
                                        case "a":
                                            v2[""] = 677
                                        default:
                                            switch "\n\t" {
                                            case "hello":
                                                for fi38 := 0; fi38 < 5; fi38++ {
                                                    var v39 uint = v13
                                                    _ = fi38
                                                    _ = v39
                                                }
                                            case "a":
                                                switch (("hello" + "hello") + ("a" + "hello")) {
                                                case "a":
                                                    if ((128 >= 569) && false) {
                                                        for fi40 := 0; fi40 < 15; fi40++ {
                                                            var v41 uint32 = 65535
                                                            _ = fi40
                                                            _ = v41
                                                        }
                                                    } else {
                                                        v1 = (float32(-2.25) - v1)
                                                        if len(v5) > 0 {
                                                            v5[0] = ""
                                                        }
                                                        if (true && (!(!false))) {
                                                            if len(v5) > 0 {
                                                                v5[0] = "\n\t"
                                                            }
                                                        } else {
                                                            v42 := GenS5{F0: 132}
                                                            _ = len(v5)
                                                            var v43 uint16 = 227
                                                            _ = v42
                                                            _ = v43
                                                        }
                                                    }
                                                }
                                            default:
                                                var v44 float64 = 1.5
                                                _ = v44
                                            }
                                        }
                                        for fi45 := 0; fi45 < 6; fi45++ {
                                            v1 = 0.1
                                            _ = fi45
                                        }
                                    }
                                    if (0 >= 100) {
                                        fmt.Println(v6, v1, v13)
                                    } else {
                                        if v6 {
                                            var v46 int64 = -1
                                            _ = v46
                                        }
                                    }
                                }
                                v13 = 0
                            }
                            switch "hello" {
                            case "a":
                                fmt.Println(v6)
                            case "foo bar":
                                fmt.Println(v1, v1)
                            case "\n\t":
                                for fi47 := 0; fi47 < 18; fi47++ {
                                    if (!v6) {
                                        var v48 uint16 = 255
                                        if v6 {
                                            v6 = (v6 || v6)
                                            v1 = ((v1 / ((1e10 - v1) / 1e10)) - (1e10 / 1e10))
                                        }
                                        var v49 int16 = 1
                                        _ = v48
                                        _ = v49
                                    } else {
                                        v13 = v13
                                    }
                                    v6 = ((("hello" + ("foo bar" + "a")) + "") == "foo bar")
                                    _ = fi47
                                }
                            }
                        }
                        v13 += 65535
                    }
                }
            case -1:
                var v50 float64 = 1.5
                _ = v50
            }
        }
        for fi51 := 0; fi51 < 15; fi51++ {
            if (-2147483648 != -1) {
                v5 = append(v5, ("foo bar" + "foo bar"))
                v6 = false
                v1 = (1.5 + ((1.5 - (v1 * 0.0)) + (1.5 * float32(float64(v1)))))
            }
            var v52 int32 = 1
            _ = fi51
            _ = v52
        }
        _ = v13
    }
    fmt.Println(v6)
    var v53 int = -1
    switch "foo bar" {
    case "foo bar":
        _ = len(v5)
    case "":
        v54 := GenS2{F0: 1, F1: 9223372036854775807, F2: 1}
        _ = v54
    default:
        v1 = ((v1 + v1) - v1)
    }
    for fi55 := 0; fi55 < 18; fi55++ {
        if ((float32(3.14159) - (1.5 + 0.0)) > v1) {
            switch "foo bar" {
            case "foo bar":
                v56 := GenS1{F0: -145, F1: 686, F2: 482}
                _ = v56
            case "":
                for fi57 := 0; fi57 < 16; fi57++ {
                    var v58 int16 = 1
                    v6 = (2147483647 != -547)
                    v5 = append(v5, "a")
                    v59 := GenS3{F0: 272, F1: -35}
                    _ = fi57
                    _ = v58
                    _ = v59
                }
            }
            v5 = append(v5, "foo bar")
            v5 = append(v5, ("hello" + ("a" + ("foo bar" + "a"))))
        }
        _ = fi55
    }
    var v60 bool = true
    for fi61 := 0; fi61 < 3; fi61++ {
        if ("foo bar" <= (("hello" + ((("hello" + "a") + (("a" + "") + "")) + "a")) + "a")) {
            fmt.Println(v53, v1)
            var v62 uint = 2
            _ = v62
        }
        v1 -= 0.1
        _ = len(v3)
        _ = fi61
    }
    switch "" {
    case "foo bar":
        v6 = ((0 < 32767) || (42 != 255))
    case "a":
        var v63 int16 = 1000
        _ = v63
    case "":
        fmt.Println(v53, v6, v53)
    default:
        for fi64 := 0; fi64 < 7; fi64++ {
            v6 = v60
            _ = fi64
        }
    }
    switch (v53 * 42) {
    case 109:
        var v65 int = v53
        _ = v65
    case -956:
        if v60 {
            for fi66 := 0; fi66 < 13; fi66++ {
                v67 := GenS1{F0: 2147483647, F1: 65535, F2: 4294967295}
                _ = fi66
                _ = v67
            }
            v68 := GenS5{F0: 42}
            _ = v68
        }
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v60
}
