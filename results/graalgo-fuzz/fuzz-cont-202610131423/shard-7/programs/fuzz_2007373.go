package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int; F1 uint8; F2 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 bool; F1 uint32; F2 int32 }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 string; F1 int16; F2 uint32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int8 }

func main() {
    var v1 uint8 = 128
    var v2 int8 = 100
    v3 := make(map[string]uint32)
    v4 := []string{("a" + ("hello" + "\n\t")), "foo bar"}
    for fi5 := 0; fi5 < 11; fi5++ {
        v2 -= 1
        _ = fi5
    }
    _ = len(v3)
    if (42 > v1) {
        if ((true && true) || (0 == 833)) {
            for fi6 := 0; fi6 < 5; fi6++ {
                for fi7 := 0; fi7 < 17; fi7++ {
                    delete(v3, ("a" + "foo bar"))
                    switch "hello" {
                    case "\n\t":
                        if true {
                            var v8 uint8 = 128
                            fmt.Println(v8)
                            v1 -= v8
                            var v9 uint64 = 927
                            _ = v8
                            _ = v9
                        } else {
                            if (!true) {
                                v1 += (v1 & 219)
                            } else {
                                fmt.Println(v1)
                                var v10 int64 = -120
                                v10 += (v10 * v10)
                                _ = v10
                            }
                            fmt.Println(fi7)
                        }
                    case "foo bar":
                        var v11 int16 = 695
                        _ = v11
                    case "hello":
                        if ("foo bar" < ("\n\t" + "")) {
                            v2 += v2
                            v1 -= v1
                            if true {
                                if (627 >= 42) {
                                    switch 2147483647 {
                                    case 107:
                                        delete(v3, "foo bar")
                                    case 2147483647:
                                        switch 0 {
                                        case -790:
                                            v2 += -1
                                        default:
                                            switch fi7 {
                                            case 9223372036854775807:
                                                _ = len(v3)
                                            case -2147483648:
                                                v1 = (v1 ^ (v1 + v1))
                                            case -1:
                                                if true {
                                                    v1 -= (v1 / (((^(v1))) | 1))
                                                    v2 += v2
                                                }
                                            }
                                        }
                                    default:
                                        v2 += v2
                                    }
                                } else {
                                    v1 = 255
                                    if false {
                                        v12 := v3["a"]
                                        _ = v12
                                    } else {
                                        if false {
                                            if (42 == 295) {
                                                v13 := GenS1{F0: 0}
                                                switch "\n\t" {
                                                case "":
                                                    var v14 uint32 = 2147483648
                                                    _ = v14
                                                }
                                                _ = v13
                                            } else {
                                                v15 := GenS2{F0: 42, F1: 42}
                                                var v16 float32 = 1.5
                                                _ = v15
                                                _ = v16
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = "a"
                                            }
                                            v2 = (v2 / (((v2 - v2)) | 1))
                                        }
                                        fmt.Println(v1, fi7, fi6)
                                    }
                                    fmt.Println(fi7)
                                }
                                var v17 uint32 = 106
                                v1 = (v1 & (^(v1)))
                                _ = v17
                            } else {
                                v2 += v2
                                if ((v1 * (v1 - (v1 << ((255) & 63)))) >= v1) {
                                    v18 := GenS3{F0: false}
                                    switch "hello" {
                                    case "foo bar":
                                        v2 = (v2 & (v2 / ((42) | 1)))
                                    case "":
                                        v3[""] = 493
                                    default:
                                        switch ("\n\t" + "foo bar") {
                                        case "foo bar":
                                            fmt.Println(fi7, v2, fi6)
                                        default:
                                            switch "\n\t" {
                                            case "a":
                                                v19 := GenS4{F0: 18446744073709551615, F1: ((3.14159 - 0.0) <= 0.0)}
                                                _ = v19
                                            case "\n\t":
                                                var v20 int64 = 9223372036854775807
                                                _ = v20
                                            case "hello":
                                                if (585 >= 124) {
                                                    var v21 int16 = 791
                                                    _ = len(v3)
                                                    switch ("\n\t" + ("\n\t" + (("\n\t" + "a") + "\n\t"))) {
                                                    case "":
                                                        if (((((v1 + v1) != 42) || true) || (42 < -1)) && true) {
                                                            v22 := GenS4{F0: 873, F1: (v21 == v21)}
                                                            _ = v22
                                                        } else {
                                                            if (!((1.5 - 0.1) > -2.25)) {
                                                                var v23 float32 = 0.0
                                                                v21 -= v21
                                                                var v24 uint8 = 0
                                                                switch ((("foo bar" + ((((("foo bar" + "hello") + (("\n\t" + "foo bar") + "\n\t")) + ((("hello" + "") + ("hello" + (("a" + (("foo bar" + "a") + ("foo bar" + "\n\t"))) + ""))) + "a")) + (((("foo bar" + "a") + ("foo bar" + "hello")) + "\n\t") + (("hello" + ("hello" + (("" + ("" + (("\n\t" + ("" + "")) + "\n\t"))) + "a"))) + ("foo bar" + "a")))) + "hello")) + ("foo bar" + "")) + "") {
                                                                case "a":
                                                                    var v25 uint32 = 4294967295
                                                                    switch -1 {
                                                                    case 127:
                                                                        v21 -= v21
                                                                        fmt.Println(fi7, fi7, v24)
                                                                        var v26 uint32 = v25
                                                                        _ = v26
                                                                    default:
                                                                        v27 := GenS5{F0: -2147483648, F1: 53}
                                                                        _ = v27
                                                                    }
                                                                    v24 += (v1 & uint8(v25))
                                                                    _ = v25
                                                                case "hello":
                                                                    fmt.Println(v23, v21)
                                                                case "":
                                                                    v4 = append(v4, "\n\t")
                                                                default:
                                                                    fmt.Println(v24)
                                                                }
                                                                v3["\n\t"] = 3
                                                                _ = v23
                                                                _ = v24
                                                            } else {
                                                                v28 := GenS6{F0: fi7, F1: (^(v1)), F2: (2 == 4294967295)}
                                                                var v29 int32 = 392
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                            v1 = v1
                                                        }
                                                    default:
                                                        v30 := GenS7{F0: 255}
                                                        _ = v30
                                                    }
                                                    _ = v21
                                                }
                                            }
                                            v2 += v2
                                            var v31 uint8 = (v1 + (v1 & v1))
                                            _ = v31
                                        }
                                    }
                                    v32 := GenS8{F0: v2}
                                    var v33 float32 = (0.1 - 0.1)
                                    var v34 bool = (!true)
                                    v35 := GenS9{F0: (v1 + (v1 | v1)), F1: ("" + "hello")}
                                    _ = v18
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                v36 := GenS9{F0: (^(v1)), F1: "hello"}
                                _ = v36
                            }
                        } else {
                            delete(v3, "foo bar")
                            var v37 uint16 = 0
                            _ = v37
                        }
                    }
                    switch "a" {
                    case "":
                        fmt.Println(fi7, fi7)
                    case "foo bar":
                        v2 += (^(v2))
                    case "hello":
                        if ((v1 | v1) != v1) {
                            switch fi7 {
                            case 128:
                                var v38 float64 = 1.5
                                _ = v38
                            case -2147483648:
                                if false {
                                    var v39 int16 = 1
                                    v2 = 42
                                    v2 -= int8(fi7)
                                    _ = v39
                                }
                            case 32768:
                                var v40 uint64 = 0
                                _ = v40
                            default:
                                fmt.Println(fi7, v1)
                            }
                            var v41 uint = 255
                            _ = v41
                        } else {
                            v2 = v2
                            if false {
                                v1 -= (v1 - (v1 * v1))
                            } else {
                                var v42 uint16 = 32767
                                v43 := GenS10{F0: (985 != 0), F1: 42, F2: 84}
                                _ = v42
                                _ = v43
                            }
                            v44 := GenS11{F0: 2147483647, F1: v2}
                            _ = v44
                        }
                        v1 = (v1 / ((42) | 1))
                        var v45 uint = 618
                        _ = v45
                    default:
                        var v46 float64 = 0.1
                        _ = v46
                    }
                    var v47 int16 = 0
                    _ = fi7
                    _ = v47
                }
                fmt.Println(fi6, fi6)
                _ = fi6
            }
            var v48 int16 = 0
            v48 -= 32767
            _ = v48
        }
        switch 219 {
        case 0:
            var v49 int32 = 591
            _ = v49
        case -901:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 128:
            for fi50 := 0; fi50 < 11; fi50++ {
                var v51 uint16 = 65535
                v1 += 1
                if (!false) {
                    v3["hello"] = 0
                } else {
                    v51 = v51
                    var v52 uint = 0
                    var v53 float32 = (3.14159 * (((1e10 - (-2.25 + (0.1 / 0.1))) + 1e10) / 1e10))
                    _ = v52
                    _ = v53
                }
                _ = fi50
                _ = v51
            }
        }
        fmt.Println(v2)
    } else {
        var v54 uint16 = 1
        _ = v54
    }
    v1 = (^(v1))
    if (!false) {
        switch 0 {
        case 127:
            if true {
                v55 := GenS12{F0: "foo bar", F1: -32768, F2: 42}
                switch 128 {
                case 0:
                    v4 = append(v4, ((("a" + ("" + (("a" + ("\n\t" + "hello")) + (("a" + "foo bar") + "foo bar")))) + "") + ("hello" + "a")))
                }
                _ = v55
            } else {
                _ = len(v4)
                switch "\n\t" {
                case "foo bar":
                    v56 := GenS2{F0: 1, F1: 0}
                    _ = v56
                default:
                    for fi57 := 0; fi57 < 6; fi57++ {
                        v2 = (v2 >> ((0) & 63))
                        v58 := GenS13{F0: 0}
                        _ = fi57
                        _ = v58
                    }
                }
                v1 -= v1
            }
        case 32768:
            for fi59 := 0; fi59 < 9; fi59++ {
                fmt.Println(v2, fi59)
                if true {
                    var v60 int8 = v2
                    for fi61 := 0; fi61 < 19; fi61++ {
                        fmt.Println(v1, v2)
                        v1 += 159
                        _ = fi61
                    }
                    for fi62 := 0; fi62 < 14; fi62++ {
                        v63 := GenS14{F0: false}
                        _ = fi62
                        _ = v63
                    }
                    _ = v60
                } else {
                    fmt.Println(fi59, fi59)
                    v3[(("hello" + "") + "a")] = 0
                    for fi64 := 0; fi64 < 19; fi64++ {
                        var v65 int32 = 1
                        if len(v4) > 0 {
                            v4[0] = ("foo bar" + ("\n\t" + ""))
                        }
                        _ = fi64
                        _ = v65
                    }
                }
                _ = fi59
            }
        default:
            var v66 uint64 = 18446744073709551615
            _ = v66
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            var v67 int32 = 1
            _ = v67
        case "":
            switch 1 {
            case -556:
                v4 = append(v4, "hello")
            default:
                v1 = (v1 * (v1 | v1))
            }
        case "foo bar":
            for fi68 := 0; fi68 < 9; fi68++ {
                for fi69 := 0; fi69 < 5; fi69++ {
                    v2 -= v2
                    _ = fi69
                }
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                var v70 bool = false
                _ = fi68
                _ = v70
            }
        }
        var v71 uint16 = 903
        _ = v71
    }
    delete(v3, "\n\t")
    v72 := GenS15{F0: (v2 | (v2 << ((42) & 63)))}
    var v73 uint16 = 42
    var v74 float32 = -2.25
    if true {
        for fi75 := 0; fi75 < 8; fi75++ {
            if false {
                v74 += ((float32(0.0) / 1e10) / 0.1)
                v76 := GenS14{F0: ((true || true) || true)}
                _ = v76
            } else {
                if false {
                    var v77 uint32 = 1
                    _ = v77
                }
                fmt.Println(v73)
            }
            _ = fi75
        }
        var v78 uint16 = 930
        v2 += v2
        _ = v78
    }
    v79 := v3["a"]
    v80 := v3["a"]
    v80 -= (v80 % ((v80) | 1))
    v81 := GenS2{F0: -1, F1: 9223372036854775807}
    v73 = v73
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v72
    _ = v73
    _ = v74
    _ = v79
    _ = v80
    _ = v81
}
