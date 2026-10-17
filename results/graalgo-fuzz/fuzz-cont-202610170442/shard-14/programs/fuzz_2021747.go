package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 uint16 = 255
    var v2 int8 = 127
    var v3 uint64 = 4294967295
    var v4 uint32 = 65535
    v5 := make(map[string]int)
    var v6 uint16 = uint16(v3)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint64 = (v3 + v3)
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint16 = v6
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            if (false || true) {
                v12 := []int{-1, 1, (^(fi11)), 0}
                _ = v12
            } else {
                if false {
                    v1 = v6
                    v13 := GenS1{F0: (^(fi11))}
                    _ = v13
                }
                v14 := GenS2{F0: 42, F1: true}
                _ = v14
            }
            switch "hello" {
            case "a":
                var v15 uint8 = 42
                _ = v15
            }
            var v16 int64 = -314
            _ = fi11
            _ = v16
        }
        _ = fi7
        _ = v8
    }
    v17 := []int{}
    v18 := GenS2{F0: uint8(v4), F1: (!((-2.25 + 0.0) < (3.14159 - -2.25)))}
    var v19 int = 2147483647
    v6 = uint16(v4)
    var v20 int32 = 0
    var v21 uint32 = v4
    var v22 int8 = v2
    var v23 int32 = v20
    var v24 int32 = -1
    fmt.Println(v20, v20, v20)
    if false {
        var v25 int64 = -1
        var v26 int64 = v25
        v27 := GenS2{F0: 0, F1: true}
        _ = v25
        _ = v26
        _ = v27
    }
    if ((v20 != (^(v24))) && (!false)) {
        v28 := GenS2{F0: 127, F1: true}
        switch "foo bar" {
        case "hello":
            v22 -= v2
        case "foo bar":
            if false {
                for fi29 := 0; fi29 < 9; fi29++ {
                    v6 = v1
                    if (int64(v6) != 501) {
                        v3 = (v3 >> ((4294967295) & 63))
                        delete(v5, "foo bar")
                    } else {
                        v30 := GenS3{F0: 1}
                        switch "" {
                        case "\n\t":
                            switch "" {
                            case "a":
                                v31 := GenS4{F0: "foo bar", F1: -1}
                                _ = v31
                            }
                        case "":
                            v5[""] = (v19 * 42)
                        case "foo bar":
                            v32 := GenS4{F0: "foo bar", F1: (v22 << ((4294967295) & 63))}
                            _ = v32
                        default:
                            var v33 bool = true
                            _ = v33
                        }
                        _ = v30
                    }
                    if false {
                        v24 += 2147483647
                        if (!(false && (!(true || (0.1 < (-2.25 - 0.1)))))) {
                            fmt.Println(v24)
                        }
                    }
                    _ = fi29
                }
            } else {
                switch "\n\t" {
                case "\n\t":
                    if false {
                        var v34 int32 = v23
                        _ = v34
                    }
                case "hello":
                    for fi35 := 0; fi35 < 8; fi35++ {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            if (0 >= 255) {
                                switch ("a" + "") {
                                case "a":
                                    v20 -= 1
                                case "\n\t":
                                    v37 := GenS2{F0: 0, F1: false}
                                    _ = v37
                                case "":
                                    v6 += (^(v1))
                                }
                                var v38 int64 = 0
                                _ = len(v17)
                                var v39 uint8 = 203
                                var v40 float32 = 3.14159
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            var v41 uint16 = v1
                            v42 := v5["foo bar"]
                            _ = fi36
                            _ = v41
                            _ = v42
                        }
                        v43 := GenS4{F0: "foo bar", F1: (v2 << ((361) & 63))}
                        _ = fi35
                        _ = v43
                    }
                default:
                    for fi44 := 0; fi44 < 16; fi44++ {
                        if true {
                            switch fi44 {
                            case 97:
                                var v45 uint16 = v1
                                _ = v45
                            default:
                                if (false || ((v22 % ((1) | 1)) < v22)) {
                                    v17 = append(v17, v19)
                                    var v46 uint32 = (v4 + (v21 ^ v21))
                                    fmt.Println(v3, v20)
                                    _ = v46
                                }
                            }
                            switch ("\n\t" + ((((("foo bar" + (("foo bar" + "foo bar") + "hello")) + "") + "") + ("hello" + "")) + "a")) {
                            case "hello":
                                var v47 int = (fi44 / ((976) | 1))
                                _ = v47
                            case "a":
                                v48 := GenS5{F0: v6, F1: (v21 % ((2147483648) | 1)), F2: v4}
                                _ = v48
                            case "":
                                for fi49 := 0; fi49 < 8; fi49++ {
                                    v50 := GenS6{F0: (v6 - (v1 | (v6 + 2))), F1: ""}
                                    switch ("hello" + "foo bar") {
                                    case "hello":
                                        switch "" {
                                        case "":
                                            v51 := GenS1{F0: -281}
                                            _ = v51
                                        case "a":
                                            var v52 uint = 42
                                            _ = v52
                                        default:
                                            v24 = (^(v23))
                                        }
                                    default:
                                        var v53 uint16 = (v6 | (v6 % ((uint16(v3)) | 1)))
                                        _ = v53
                                    }
                                    fmt.Println(v24)
                                    _ = fi49
                                    _ = v50
                                }
                            default:
                                var v54 float64 = (1e10 + 3.14159)
                                _ = v54
                            }
                        }
                        v6 = v1
                        if true {
                            fmt.Println(v2)
                        }
                        _ = fi44
                    }
                }
                if false {
                    var v55 string = "foo bar"
                    switch "\n\t" {
                    case "hello":
                        var v56 int8 = (v2 - v2)
                        _ = v56
                    default:
                        if (!true) {
                            switch -2147483648 {
                            case 127:
                                delete(v5, "a")
                            case 1:
                                for fi57 := 0; fi57 < 12; fi57++ {
                                    v58 := GenS7{F0: 1, F1: v6}
                                    switch (fi57 + 586) {
                                    case 9223372036854775807:
                                        var v59 int = (v19 - (fi57 >> ((201) & 63)))
                                        _ = v59
                                    case 32768:
                                        v20 += -2147483648
                                    case -9223372036854775808:
                                        switch "" {
                                        case "foo bar":
                                            v60 := GenS8{F0: ("hello" + ""), F1: v20}
                                            _ = v60
                                        }
                                    }
                                    _ = fi57
                                    _ = v58
                                }
                            default:
                                var v61 int8 = 42
                                _ = v61
                            }
                        } else {
                            for fi62 := 0; fi62 < 14; fi62++ {
                                if (((47 > -1) || true) && true) {
                                    v63 := GenS6{F0: 605, F1: ""}
                                    _ = v63
                                } else {
                                    if len(v17) > 0 {
                                        v17[0] = 1
                                    }
                                }
                                _ = fi62
                            }
                            v55 = ("a" + v55)
                        }
                    }
                    _ = v55
                }
                for fi64 := 0; fi64 < 15; fi64++ {
                    fmt.Println(v21, v2)
                    v65 := GenS9{F0: v1, F1: 18446744073709551615}
                    _ = fi64
                    _ = v65
                }
            }
        case "a":
            for fi66 := 0; fi66 < 14; fi66++ {
                fmt.Println(fi66)
                for fi67 := 0; fi67 < 15; fi67++ {
                    var v68 int64 = 42
                    v69 := GenS5{F0: (^(v6)), F1: v21, F2: 4294967295}
                    _ = len(v5)
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                _ = fi66
            }
        }
        for fi70 := 0; fi70 < 9; fi70++ {
            v24 += v24
            _ = fi70
        }
        _ = v28
    } else {
        v71 := GenS4{F0: (((("hello" + "") + ("a" + "foo bar")) + (((("" + "foo bar") + "") + "\n\t") + ("foo bar" + ("\n\t" + "hello")))) + ("a" + "foo bar")), F1: v22}
        _ = v71
    }
    v72 := GenS4{F0: "hello", F1: v22}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
}
