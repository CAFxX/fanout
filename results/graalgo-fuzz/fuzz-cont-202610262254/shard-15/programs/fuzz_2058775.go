package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS7 struct { F0 int32; F1 int; F2 bool }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint }
func genFn0(p0 int32) bool {
    fmt.Println(p0)
    p0 = (^(p0))
    var v6 int64 = 405
    switch -1 {
    case 581:
        var v7 uint = 4294967295
        _ = v7
    default:
        v8 := GenS1{F0: 255, F1: 1, F2: 127}
        _ = v8
    }
    _ = p0
    _ = v6
    return true
}
func genFn1() uint8 {
    v9 := []int{}
    var v10 uint32 = 850
    var v11 bool = (!true)
    v9 = append(v9, 128)
    for fi12 := 0; fi12 < 13; fi12++ {
        v10 = v10
        if v11 {
            fmt.Println(v10, v10)
            v10 = v10
        } else {
            if (v11 || false) {
                v9 = append(v9, (fi12 % ((fi12) | 1)))
            } else {
                for fi13 := 0; fi13 < 16; fi13++ {
                    switch fi12 {
                    case 4294967295:
                        var v14 bool = (!v11)
                        _ = v14
                    case -341:
                        if (0.0 > 1.5) {
                            fmt.Println(fi12, v10)
                            v15 := GenS1{F0: 65535, F1: uint(v10), F2: 0}
                            v9 = append(v9, (^(fi13)))
                            _ = v15
                        }
                    default:
                        if (!(9223372036854775807 > 9223372036854775807)) {
                            v10 -= 65535
                        } else {
                            v10 = 0
                            fmt.Println(v11, fi13, fi12)
                            if (0 < 2) {
                                v10 = (v10 ^ (^(v10)))
                                v9 = append(v9, fi13)
                            } else {
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    if len(v9) > 0 {
                                        v9[0] = (fi13 + fi13)
                                    }
                                case "hello":
                                    var v16 uint8 = 2
                                    _ = v16
                                }
                                if (!v11) {
                                    var v17 int8 = -1
                                    _ = v17
                                } else {
                                    var v18 uint64 = 42
                                    _ = v18
                                }
                                v19 := GenS2{F0: 9223372036854775807, F1: 18446744073709551615}
                                switch 127 {
                                case 1:
                                    var v20 string = "hello"
                                    _ = v20
                                case -2147483648:
                                    switch "" {
                                    case "\n\t":
                                        _ = len(v9)
                                    }
                                case 2147483647:
                                    var v21 bool = (((v11 || false) && v11) && (!(9223372036854775807 >= 156)))
                                    _ = v21
                                default:
                                    if v11 {
                                        switch fi13 {
                                        case 9223372036854775807:
                                            fmt.Println(fi13, fi13, v11)
                                        case 42:
                                            var v22 float32 = (3.14159 - 1e10)
                                            _ = v22
                                        }
                                    } else {
                                        v23 := GenS1{F0: 42, F1: 437, F2: 127}
                                        switch (fi13 >> ((255) & 63)) {
                                        case -9223372036854775808:
                                            v10 -= v10
                                            if len(v9) > 0 {
                                                v9[0] = 106
                                            }
                                        }
                                        var v24 uint32 = 977
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                if true {
                                    v9 = append(v9, (fi13 >> ((42) & 63)))
                                    if (68 == 684) {
                                        var v25 float32 = float32(1e10)
                                        var v26 uint32 = v10
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v27 uint32 = 735
                                        _ = v27
                                    }
                                    var v28 uint16 = 263
                                    _ = v28
                                }
                                _ = v19
                            }
                        }
                    }
                    v11 = v11
                    var v29 uint16 = 255
                    v29 = (^(v29))
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v30 int = (^(fi13))
                    _ = len(v9)
                    _ = fi13
                    _ = v29
                    _ = v30
                }
            }
            switch (fi12 % ((fi12) | 1)) {
            case 4294967295:
                _ = len(v9)
            case 1:
                v9 = append(v9, fi12)
            default:
                v31 := GenS2{F0: 42, F1: 0}
                _ = v31
            }
        }
        fmt.Println(v11, v10, v10)
        _ = fi12
    }
    _ = v10
    _ = v11
    _ = v9
    return 127
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"\n\t", "foo bar", ("foo bar" + (("hello" + "foo bar") + ("foo bar" + "")))}
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    var v32 uint16 = 65535
    var v33 float32 = ((0.1 * 0.1) + (1.5 / (0.0 - 1.5)))
    switch 128 {
    case 2147483647:
        v34 := GenS2{F0: -382, F1: 445}
        _ = v34
    case 637:
        for fi35 := 0; fi35 < 5; fi35++ {
            for fi36 := 0; fi36 < 19; fi36++ {
                if false {
                    if (true && true) {
                        fmt.Println(v32, v33)
                        if (fi36 != (fi35 | 253)) {
                            switch "" {
                            case "hello":
                                v1[fi35] = fi36
                            case "\n\t":
                                var v37 uint8 = 255
                                _ = v37
                            case "foo bar":
                                if len(v2) > 0 {
                                    v2[0] = "\n\t"
                                }
                            default:
                                var v38 uint32 = 42
                                _ = v38
                            }
                            v39 := GenS3{F0: 1}
                            v32 = (v32 & 32767)
                            fmt.Println(v33, fi36, v32)
                            var v40 uint64 = 981
                            _ = v39
                            _ = v40
                        } else {
                            v32 = 32767
                            v33 -= ((v33 * v33) * 3.14159)
                        }
                    } else {
                        fmt.Println(v32, fi35, fi36)
                        var v41 int8 = 0
                        if (!(18446744073709551615 > 119)) {
                            v33 = float32(0.1)
                            _ = len(v4)
                        }
                        _ = v41
                    }
                    switch fi36 {
                    case 127:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 371:
                        if false {
                            v33 += v33
                            v33 += ((v33 - v33) - 0.0)
                        } else {
                            fmt.Println(fi35)
                            v42 := GenS4{F0: ("hello" + "\n\t")}
                            if ("\n\t" >= "") {
                                var v43 uint32 = 4294967295
                                switch "a" {
                                case "":
                                    v43 = v43
                                case "\n\t":
                                    var v44 int64 = -269
                                    _ = v44
                                default:
                                    var v45 uint8 = 42
                                    _ = v45
                                }
                                var v46 bool = (!(v32 < 636))
                                v3[fi36] = "\n\t"
                                _ = v43
                                _ = v46
                            } else {
                                v47 := GenS5{F0: v32, F1: 0}
                                _ = v47
                            }
                            _ = v42
                        }
                    case -2147483648:
                        if len(v2) > 0 {
                            v2[0] = ("a" + "")
                        }
                    default:
                        v48 := GenS4{F0: ""}
                        _ = v48
                    }
                    fmt.Println(v32, v33)
                }
                if (0 <= 127) {
                    v49 := GenS6{F0: (v32 + 65535), F1: 1, F2: 127}
                    v50 := GenS7{F0: -1, F1: -593, F2: (true && (true && true))}
                    v5[316] = (("" + ("\n\t" + ("foo bar" + ("foo bar" + "hello")))) + ("foo bar" + "a"))
                    _ = v49
                    _ = v50
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi36
            }
            _ = fi35
        }
    case 42:
        v32 = 593
    default:
        _ = len(v5)
    }
    v32 -= v32
    var v51 bool = (("" + "hello") < "foo bar")
    v52 := GenS8{F0: 9223372036854775807}
    for fi53 := 0; fi53 < 16; fi53++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v54 := GenS9{F0: 0}
        _ = fi53
        _ = v54
    }
    for fi55 := 0; fi55 < 7; fi55++ {
        if v51 {
            if false {
                if (!true) {
                    var v56 int64 = -1
                    var v57 uint8 = 0
                    if (42 <= 1) {
                        v56 += -9223372036854775808
                        var v58 int32 = 1
                        fmt.Println(v33)
                        _ = v58
                    }
                    _ = v56
                    _ = v57
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v33 = 1e10
                }
                switch (fi55 << ((2) & 63)) {
                case -631:
                    v51 = true
                }
            } else {
                for fi59 := 0; fi59 < 8; fi59++ {
                    var v60 uint16 = v32
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v51 = (false || (false && v51))
                    v60 = (v60 & v32)
                    _ = fi59
                    _ = v60
                }
            }
        } else {
            fmt.Println(v32, v33, v51)
            v61 := GenS3{F0: 703}
            v32 = (v32 / (((v32 << ((4294967295) & 63))) | 1))
            _ = v61
        }
        _ = fi55
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        v63 := GenS9{F0: 42}
        v32 = (v32 / (((v32 ^ (v32 & 42))) | 1))
        for fi64 := 0; fi64 < 9; fi64++ {
            if len(v2) > 0 {
                v2[0] = ("\n\t" + "foo bar")
            }
            _ = fi64
        }
        _ = fi62
        _ = v63
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v32 = (v32 | (v32 | v32))
    if true {
        if true {
            var v65 int16 = -514
            if (42 >= -128) {
                v3[-2147483648] = (("a" + "\n\t") + "hello")
            }
            fmt.Println(v32, v33, v51)
            _ = v65
        }
    } else {
        v51 = (v51 && v51)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v51
    _ = v52
}
