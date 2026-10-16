package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 string; F2 int32 }
type GenS8 struct { F0 uint32; F1 int8; F2 string }
type GenS9 struct { F0 int8; F1 int16; F2 uint16 }
type GenS10 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 int8 = -1
    v2 := []int{-340, 266}
    var v3 uint = 0
    v4 := make(map[uint32]uint32)
    var v5 int16 = -32768
    var v6 int = 128
    var v7 uint64 = 2
    v2 = append(v2, (v6 % ((42) | 1)))
    v8 := GenS1{F0: 100, F1: (^(v7)), F2: v1}
    var v9 uint64 = (^(v7))
    fmt.Println(v6, v9)
    var v10 float64 = ((1.5 + (1.5 / (float64(((0.0 - ((3.14159 / (3.14159 * 1e10)) - 1.5)) + float32(-2.25))) / (float64(3.14159) + 0.1)))) - (1e10 - (((0.1 / -2.25) - (((1.5 + ((-2.25 * 3.14159) + 1e10)) - 3.14159) / (float64(3.14159) - (1.5 - 0.1)))) * 0.1)))
    var v11 bool = ((!(true || (!false))) && true)
    v12 := GenS1{F0: (-(v1)), F1: v7, F2: (v1 / ((0) | 1))}
    if (!v11) {
        v10 = 0.0
        if true {
            var v13 float64 = v10
            v11 = false
            var v14 uint64 = (v9 | v7)
            _ = v13
            _ = v14
        } else {
            if true {
                if (!(2 < 255)) {
                    v15 := GenS2{F0: 10, F1: (v9 - (v7 - (v9 - 4294967295))), F2: int(v5)}
                    v16 := GenS2{F0: 128, F1: v9, F2: (v6 >> (((v3 / ((791) | 1))) & 63))}
                    _ = v15
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        for fi18 := 0; fi18 < 15; fi18++ {
                            switch "\n\t" {
                            case "a":
                                if v11 {
                                    var v19 uint = v3
                                    _ = v19
                                } else {
                                    switch 519 {
                                    case 2147483647:
                                        fmt.Println(v10, fi17)
                                    case 996:
                                        v20 := GenS2{F0: 255, F1: v7, F2: int(v1)}
                                        _ = v20
                                    default:
                                        if ((-2.25 / 1.5) > ((0.1 / 3.14159) / -2.25)) {
                                            var v21 uint32 = 65535
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v7 += 55
                                            _ = v21
                                        }
                                    }
                                    if (35 != 83) {
                                        var v22 int32 = -1
                                        if ((360 > 2) && true) {
                                            v23 := GenS2{F0: 42, F1: v7, F2: (v6 / (((-(v6))) | 1))}
                                            _ = v23
                                        } else {
                                            v24 := GenS3{F0: 858}
                                            _ = v24
                                        }
                                        v25 := v4[2]
                                        _ = v22
                                        _ = v25
                                    }
                                }
                            case "foo bar":
                                var v26 float64 = (float64(0.1) / (((v10 - (((0.0 - -2.25) + v10) / v10)) - v10) + ((v10 + v10) / v10)))
                                _ = v26
                            case "\n\t":
                                v27 := GenS4{F0: 42, F1: 127}
                                _ = v27
                            default:
                                switch fi18 {
                                case 9223372036854775807:
                                    var v28 float64 = v10
                                    _ = v28
                                case -78:
                                    switch int(v3) {
                                    case -612:
                                        switch (fi17 >> (((^(v3))) & 63)) {
                                        case 791:
                                            v10 = v10
                                            var v29 bool = (0 <= -395)
                                            _ = v29
                                        case 127:
                                            if v11 {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v30 := GenS5{F0: 42}
                                                    _ = v30
                                                default:
                                                    var v31 float32 = 3.14159
                                                    _ = v31
                                                }
                                                switch fi18 {
                                                case 717:
                                                    if v11 {
                                                        var v32 int16 = int16(v7)
                                                        fmt.Println(v5)
                                                        v33 := GenS6{F0: (^(v1))}
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                case -9223372036854775808:
                                                    fmt.Println(v9)
                                                case 2147483647:
                                                    v11 = (124 >= 26)
                                                }
                                                v10 += v10
                                            }
                                            var v34 uint64 = 381
                                            var v35 int = -9223372036854775808
                                            if v11 {
                                                fmt.Println(v6, v9)
                                                v36 := GenS7{F0: uint64(v3), F1: "", F2: 0}
                                                switch "hello" {
                                                case "a":
                                                    if (v11 && v11) {
                                                        v37 := GenS8{F0: 2147483647, F1: (v1 * 69), F2: "a"}
                                                        switch -9223372036854775808 {
                                                        case 127:
                                                            var v38 int16 = (v5 << (((v3 - 2)) & 63))
                                                            _ = v38
                                                        case 9223372036854775807:
                                                            if (("a" + (("hello" + ((("hello" + "a") + ("\n\t" + "a")) + ("foo bar" + ("foo bar" + ("\n\t" + ("a" + ("foo bar" + "hello"))))))) + (((("foo bar" + "") + "") + ("foo bar" + ("a" + "a"))) + "foo bar"))) != "a") {
                                                                v2 = append(v2, 4294967295)
                                                                var v39 uint = v3
                                                                if (0 > 0) {
                                                                    fmt.Println(v10)
                                                                } else {
                                                                    v40 := GenS3{F0: 65535}
                                                                    v10 = v10
                                                                    if (!(3.14159 != (0.0 / ((((3.14159 + float32(v10)) + -2.25) + float32(v10)) - float32((v10 / v10)))))) {
                                                                        var v41 uint64 = v34
                                                                        fmt.Println(v6, v39, v34)
                                                                        _ = v41
                                                                    } else {
                                                                        v34 = v34
                                                                    }
                                                                    if ((9223372036854775807 <= 0) && false) {
                                                                        v34 -= (v7 + v9)
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v1 = 38
                                                                        case "a":
                                                                            if (v7 != 0) {
                                                                                v4[1] = 2147483648
                                                                            }
                                                                        case "\n\t":
                                                                            v4[0] = 930
                                                                        default:
                                                                            v9 = v7
                                                                        }
                                                                        var v42 int16 = (v5 >> (((v39 / (((v39 - (v39 - 2))) | 1))) & 63))
                                                                        v43 := GenS9{F0: v1, F1: v5, F2: 863}
                                                                        _ = v42
                                                                        _ = v43
                                                                    }
                                                                    v5 += v5
                                                                    _ = v40
                                                                }
                                                                _ = v39
                                                            }
                                                        default:
                                                            v44 := GenS6{F0: (^(v1))}
                                                            _ = v44
                                                        }
                                                        v45 := GenS1{F0: v1, F1: 568, F2: 42}
                                                        var v46 uint32 = 42
                                                        _ = v37
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                default:
                                                    switch "a" {
                                                    case "hello":
                                                        v11 = (!v11)
                                                    }
                                                }
                                                _ = v36
                                            } else {
                                                var v47 string = "a"
                                                _ = v47
                                            }
                                            v11 = v11
                                            switch "\n\t" {
                                            case "\n\t":
                                                if (v3 < 920) {
                                                    v10 += v10
                                                }
                                            case "hello":
                                                v5 = (v5 << (((^(v3))) & 63))
                                            default:
                                                _ = len(v4)
                                            }
                                            v48 := GenS7{F0: 2, F1: "a", F2: int32(fi17)}
                                            var v49 string = ("a" + "")
                                            var v50 uint32 = 42
                                            var v51 string = ("a" + v49)
                                            if v11 {
                                                if false {
                                                    v2 = append(v2, 127)
                                                }
                                            } else {
                                                var v52 int = (fi18 >> ((v3) & 63))
                                                fmt.Println(v9, v3)
                                                _ = v52
                                            }
                                            _ = v34
                                            _ = v35
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        case 128:
                                            if v11 {
                                                var v53 float64 = v10
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v53
                                            }
                                        default:
                                            var v54 uint = 1
                                            _ = v54
                                        }
                                    case -2147483648:
                                        switch "foo bar" {
                                        case "hello":
                                            v6 = v6
                                        case "":
                                            v55 := v4[484]
                                            _ = v55
                                        case "a":
                                            fmt.Println(v7, v5)
                                        }
                                    case -960:
                                        fmt.Println(v10, v10)
                                    }
                                }
                            }
                            _ = fi18
                        }
                        v5 -= -5
                        _ = fi17
                    }
                    if ((v3 & v3) == (v3 % ((255) | 1))) {
                        v10 = 0.1
                        for fi56 := 0; fi56 < 14; fi56++ {
                            for fi57 := 0; fi57 < 18; fi57++ {
                                v1 = (v1 << (((v3 >> ((0) & 63))) & 63))
                                _ = len(v2)
                                var v58 uint16 = 776
                                fmt.Println(v5)
                                _ = fi57
                                _ = v58
                            }
                            fmt.Println(v5)
                            v59 := GenS10{F0: 0, F1: 1}
                            _ = fi56
                            _ = v59
                        }
                    }
                    if (!true) {
                        switch -984 {
                        case 32768:
                            v60 := GenS8{F0: 383, F1: v1, F2: ""}
                            _ = v60
                        case 128:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v5 = v5
                    }
                }
                fmt.Println(v1, v5)
            } else {
                v61 := GenS8{F0: 0, F1: (v1 << ((42) & 63)), F2: "a"}
                _ = v61
            }
            v62 := GenS7{F0: v9, F1: "foo bar", F2: 396}
            fmt.Println(v10)
            _ = v62
        }
    } else {
        var v63 string = "\n\t"
        _ = v63
    }
    fmt.Println(v6, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
