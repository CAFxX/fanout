package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int; F1 int16; F2 int }
type GenS7 struct { F0 string; F1 int16 }
type GenS8 struct { F0 int32; F1 int8 }
type GenS9 struct { F0 int32; F1 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int; F1 uint }
type GenS12 struct { F0 uint; F1 int; F2 uint32 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]int)
    v3 := []string{((("" + "") + ("a" + ("foo bar" + ("a" + "")))) + ""), "a", "\n\t", ("\n\t" + "foo bar")}
    var v4 int = 380
    v5 := make(map[string]string)
    var v6 int8 = int8(v1)
    if (true || false) {
        v7 := GenS1{F0: 128, F1: 0}
        _ = v7
    }
    v8 := GenS1{F0: 128, F1: (v4 & (^(v4)))}
    v6 = -128
    fmt.Println(v4, v6)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2[v1] = v4
    var v9 uint16 = 255
    _ = len(v3)
    v10 := GenS2{F0: v1}
    switch "hello" {
    case "hello":
        var v11 string = (("\n\t" + ("hello" + "foo bar")) + ("" + "hello"))
        _ = v11
    case "foo bar":
        var v12 int64 = -1
        _ = v12
    }
    switch "foo bar" {
    case "a":
        for fi13 := 0; fi13 < 4; fi13++ {
            if true {
                switch -9223372036854775808 {
                case 523:
                    switch v4 {
                    case -2147483648:
                        switch "" {
                        case "":
                            v14 := GenS3{F0: (((("\n\t" + "foo bar") + "a") + ((("hello" + "hello") + "\n\t") + "a")) + "a")}
                            _ = v14
                        case "foo bar":
                            switch 32768 {
                            case 2147483647:
                                v9 = v9
                            case -9223372036854775808:
                                switch (fi13 << ((255) & 63)) {
                                case 127:
                                    var v15 float64 = -2.25
                                    _ = v15
                                case -262:
                                    v6 -= v6
                                default:
                                    fmt.Println(v1, v6, fi13)
                                }
                            }
                        case "a":
                            var v16 float64 = 3.14159
                            _ = v16
                        default:
                            if (42 != 32767) {
                                var v17 int64 = int64(v9)
                                v6 -= 100
                                _ = v17
                            }
                        }
                    case 9223372036854775807:
                        delete(v5, "\n\t")
                    case 2147483647:
                        v18 := GenS2{F0: v1}
                        _ = v18
                    default:
                        var v19 uint = uint(v1)
                        _ = v19
                    }
                case -358:
                    for fi20 := 0; fi20 < 9; fi20++ {
                        v4 -= 9223372036854775807
                        _ = fi20
                    }
                default:
                    fmt.Println(v1, v4, v4)
                }
                switch "a" {
                case "hello":
                    var v21 int8 = 1
                    _ = v21
                case "":
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v23 := v2[v1]
                        _ = len(v3)
                        _ = fi22
                        _ = v23
                    }
                case "a":
                    fmt.Println(v4, v1)
                }
            }
            if (!(false && (17 > 1))) {
                switch fi13 {
                case 2147483647:
                    v6 += 127
                case -966:
                    if true {
                        if false {
                            if false {
                                if (true || false) {
                                    fmt.Println(fi13)
                                    var v24 uint32 = v1
                                    _ = v24
                                }
                            }
                        } else {
                            v4 += (fi13 + v4)
                            v25 := GenS4{F0: 919, F1: ""}
                            v4 = v4
                            _ = v25
                        }
                        var v26 float64 = 0.1
                        v4 = fi13
                        _ = v26
                    } else {
                        var v27 int64 = -851
                        var v28 uint8 = 127
                        switch "a" {
                        case "a":
                            switch "foo bar" {
                            case "":
                                v29 := GenS4{F0: (^(v4)), F1: "\n\t"}
                                _ = v29
                            case "foo bar":
                                v1 -= 96
                            }
                        case "hello":
                            v27 = v27
                        default:
                            v28 = v28
                        }
                        _ = v27
                        _ = v28
                    }
                default:
                    var v30 uint64 = 0
                    _ = v30
                }
                for fi31 := 0; fi31 < 15; fi31++ {
                    if (!(false && (true && false))) {
                        v3 = append(v3, "\n\t")
                        if (42 > v6) {
                            var v32 int32 = -2147483648
                            v32 += (v32 << ((499) & 63))
                            v32 = v32
                            _ = v32
                        }
                        v1 = (^(v1))
                    } else {
                        if (!(0.1 <= ((0.0 + 0.1) - 1e10))) {
                            var v33 uint64 = 0
                            _ = v33
                        } else {
                            v1 += v1
                        }
                        v2[(v1 | v1)] = fi13
                    }
                    _ = fi31
                }
            } else {
                for fi34 := 0; fi34 < 19; fi34++ {
                    if false {
                        v5[("\n\t" + "")] = (("" + (("\n\t" + ("foo bar" + "hello")) + "hello")) + "hello")
                    } else {
                        v35 := GenS4{F0: (fi13 / (((v4 ^ fi34)) | 1)), F1: ""}
                        _ = v35
                    }
                    if false {
                        var v36 int8 = v6
                        if false {
                            fmt.Println(v6, v9)
                            if false {
                                var v37 uint16 = v9
                                if len(v3) > 0 {
                                    v3[0] = (("" + ("hello" + "hello")) + "a")
                                }
                                _ = v37
                            }
                            var v38 uint32 = v1
                            _ = v38
                        }
                        v39 := GenS4{F0: (^(fi34)), F1: (("" + "a") + "foo bar")}
                        _ = v36
                        _ = v39
                    } else {
                        v40 := GenS3{F0: ("a" + "")}
                        if (false && (2 < 42)) {
                            fmt.Println(v4, fi34)
                            var v41 float32 = 3.14159
                            _ = len(v5)
                            var v42 int = -2147483648
                            if true {
                                var v43 int32 = 42
                                if true {
                                    v44 := GenS5{F0: 522, F1: -258}
                                    switch v4 {
                                    case -2147483648:
                                        v4 -= (fi34 - (fi34 % (((-(fi34))) | 1)))
                                    case -31:
                                        _ = len(v5)
                                    }
                                    var v45 int8 = v6
                                    _ = v44
                                    _ = v45
                                } else {
                                    v3 = append(v3, "")
                                    if (float32(0.1) != 1.5) {
                                        var v46 float64 = (float64((3.14159 * v41)) * float64(v41))
                                        var v47 int64 = 1
                                        _ = v46
                                        _ = v47
                                    }
                                    fmt.Println(fi13, v4, v43)
                                }
                                var v48 uint64 = 200
                                _ = v43
                                _ = v48
                            } else {
                                v49 := GenS6{F0: (fi13 % ((fi13) | 1)), F1: -936, F2: fi34}
                                var v50 uint16 = v9
                                v51 := v5["hello"]
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            _ = v41
                            _ = v42
                        }
                        _ = v40
                    }
                    switch "hello" {
                    case "":
                        v52 := GenS4{F0: v4, F1: (("a" + ("" + "hello")) + "")}
                        _ = v52
                    case "\n\t":
                        v1 = uint32(v9)
                    default:
                        if false {
                            if true {
                                switch "a" {
                                case "":
                                    if false {
                                        v6 = v6
                                    } else {
                                        if (42 < 331) {
                                            fmt.Println(v9, v6)
                                        } else {
                                            v9 = 1
                                            v9 -= (v9 % ((255) | 1))
                                            switch (("\n\t" + "") + ("foo bar" + ("\n\t" + (("\n\t" + "") + "foo bar")))) {
                                            case "":
                                                v53 := GenS3{F0: ""}
                                                _ = v53
                                            case "hello":
                                                v54 := GenS2{F0: v1}
                                                _ = v54
                                            }
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                case "\n\t":
                                    var v55 float32 = 0.0
                                    _ = v55
                                case "foo bar":
                                    v1 = (v1 << ((2) & 63))
                                default:
                                    if ((^(v1)) < 196) {
                                        v6 = 1
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if false {
                                            v56 := GenS7{F0: (("a" + "") + "foo bar"), F1: 1000}
                                            v9 += v9
                                            v57 := GenS6{F0: 9223372036854775807, F1: 0, F2: (-(fi13))}
                                            _ = v56
                                            _ = v57
                                        }
                                        if true {
                                            var v58 string = ""
                                            var v59 int = v4
                                            _ = v58
                                            _ = v59
                                        }
                                        v60 := GenS8{F0: -2147483648, F1: (v6 | v6)}
                                        _ = v60
                                    }
                                }
                                v61 := GenS4{F0: (-(fi34)), F1: "\n\t"}
                                _ = v61
                            } else {
                                v62 := GenS9{F0: 932, F1: uint64(v1)}
                                v1 = v1
                                v6 = -1
                                _ = v62
                            }
                            fmt.Println(v4)
                            var v63 int64 = -339
                            fmt.Println(v63)
                            _ = v63
                        } else {
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            var v64 float64 = ((1.5 * 0.1) - 0.0)
                            _ = v64
                        }
                    }
                    _ = fi34
                }
                var v65 int32 = 277
                v6 = -128
                _ = v65
            }
            fmt.Println(v6, v9, fi13)
            _ = fi13
        }
    case "foo bar":
        switch "\n\t" {
        case "":
            var v66 int32 = 183
            _ = v66
        case "foo bar":
            v9 += v9
        default:
            var v67 int32 = 1
            _ = v67
        }
    default:
        v9 -= 0
    }
    var v68 int = -229
    switch "a" {
    case "a":
        v4 += -9223372036854775808
    case "hello":
        fmt.Println(v68, v6)
    default:
        v69 := GenS10{F0: 2147483647}
        _ = v69
    }
    for fi70 := 0; fi70 < 5; fi70++ {
        var v71 bool = (true || (!false))
        _ = fi70
        _ = v71
    }
    v72 := GenS11{F0: v68, F1: 631}
    v73 := GenS12{F0: 255, F1: v4, F2: (v1 * v1)}
    fmt.Println(v68)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v72
    _ = v73
    _ = v8
    _ = v9
}
