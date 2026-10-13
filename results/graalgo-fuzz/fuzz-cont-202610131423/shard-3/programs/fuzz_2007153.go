package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int8; F1 string; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint }
type GenS6 struct { F0 int8; F1 int16; F2 uint8 }

func main() {
    var v1 uint16 = 0
    var v2 int8 = -128
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]int)
    switch -1000 {
    case -2147483648:
        v5 := GenS1{F0: ((((("" + ("hello" + "hello")) + "foo bar") + ("" + (("a" + "") + ("a" + "hello")))) + (("\n\t" + "foo bar") + "foo bar")) + ""), F1: 70}
        _ = v5
    case 32768:
        v6 := []int{-2147483648, 1}
        _ = v6
    case 2147483647:
        if (-1 > 0) {
            for fi7 := 0; fi7 < 7; fi7++ {
                v2 = (v2 % (((v2 * 68)) | 1))
                var v8 uint8 = 42
                _ = fi7
                _ = v8
            }
            var v9 uint32 = 0
            _ = v9
        } else {
            v10 := []int{2147483647}
            if false {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                switch "" {
                case "a":
                    var v11 float32 = 1.5
                    _ = v11
                }
            }
            if ((0 >= 1) && true) {
                switch "hello" {
                case "a":
                    v12 := GenS2{F0: (v2 + (v2 & -1)), F1: "a", F2: 62}
                    _ = v12
                case "hello":
                    v1 = v1
                case "":
                    var v13 int16 = 1000
                    _ = v13
                }
                fmt.Println(v2)
            } else {
                fmt.Println(v3, v2, v2)
            }
            _ = v10
        }
    default:
        v1 -= 474
    }
    fmt.Println(v3, v2)
    for fi14 := 0; fi14 < 7; fi14++ {
        if (v2 > (v2 + 1)) {
            for fi15 := 0; fi15 < 13; fi15++ {
                if true {
                    switch "" {
                    case "\n\t":
                        if true {
                            v16 := GenS2{F0: v2, F1: "a", F2: 0}
                            delete(v4, 422)
                            switch "a" {
                            case "":
                                var v17 uint8 = 128
                                _ = v17
                            case "a":
                                if (1 >= 1000) {
                                    v18 := []int{-1, fi14, (fi14 | fi14), (fi14 + -523)}
                                    _ = v18
                                } else {
                                    fmt.Println(fi14, v2, fi14)
                                    var v19 int64 = -9223372036854775808
                                    _ = v19
                                }
                            }
                            if true {
                                v20 := []int{fi15, fi15}
                                v21 := GenS1{F0: "", F1: -620}
                                _ = v20
                                _ = v21
                            } else {
                                if (!(true && (0 < 648))) {
                                    fmt.Println(fi14, v2, v1)
                                    v2 = 127
                                }
                            }
                            _ = v16
                        } else {
                            v22 := GenS3{F0: 0}
                            _ = v22
                        }
                    case "foo bar":
                        if (3.14159 > 3.14159) {
                            var v23 uint32 = 645
                            _ = v23
                        } else {
                            var v24 uint16 = (v1 << ((142) & 63))
                            v4[42] = (fi14 * (-(fi14)))
                            _ = v24
                        }
                    default:
                        var v25 float32 = (0.0 - ((0.1 * 1.5) / float32(3.14159)))
                        _ = v25
                    }
                    switch fi14 {
                    case -2147483648:
                        var v26 int = fi14
                        _ = v26
                    default:
                        v1 = (^(v1))
                    }
                }
                _ = fi15
            }
            fmt.Println(v2)
            var v27 bool = (false || true)
            _ = v27
        }
        var v28 uint16 = 42
        fmt.Println(fi14)
        _ = fi14
        _ = v28
    }
    switch 32768 {
    case -2147483648:
        v1 += (v1 << ((0) & 63))
    case 721:
        v2 = 127
    }
    v3 = v3
    var v29 uint64 = 7
    var v30 uint8 = 128
    switch 4294967295 {
    case 1:
        v31 := GenS2{F0: (v2 >> ((820) & 63)), F1: "", F2: -32768}
        _ = v31
    default:
        fmt.Println(v1)
    }
    v32 := []int{}
    var v33 uint16 = 32767
    var v34 int8 = 127
    if (true || false) {
        v3 = v3
    } else {
        if false {
            var v35 uint64 = 18446744073709551615
            _ = v35
        } else {
            var v36 int32 = 0
            _ = v36
        }
        switch 42 {
        case -2147483648:
            var v37 float32 = -2.25
            _ = v37
        case 1:
            for fi38 := 0; fi38 < 6; fi38++ {
                if false {
                    for fi39 := 0; fi39 < 15; fi39++ {
                        if len(v32) > 0 {
                            v32[0] = (fi39 % ((int(v30)) | 1))
                        }
                        v2 -= v2
                        switch "a" {
                        case "":
                            fmt.Println(v34, fi39, v2)
                        default:
                            var v40 bool = (!(true && (v34 > v34)))
                            _ = v40
                        }
                        v29 = (v29 - (v29 ^ (v29 * (v29 ^ 42))))
                        if (((-2.25 / 1.5) - 0.0) <= (0.0 / ((0.0 * 3.14159) - 1e10))) {
                            v1 = v33
                            v3 = -1
                        } else {
                            v29 = v29
                            if true {
                                if true {
                                    v41 := GenS4{F0: 4294967295, F1: 2147483647}
                                    v30 = (v30 >> ((430) & 63))
                                    v32 = append(v32, 4294967295)
                                    _ = v41
                                } else {
                                    _ = len(v32)
                                    switch fi38 {
                                    case 32768:
                                        var v42 uint16 = v1
                                        _ = v42
                                    case -155:
                                        v32 = append(v32, fi38)
                                    case 729:
                                        var v43 uint8 = (v30 * (v30 + v30))
                                        _ = v43
                                    }
                                    v30 = v30
                                    v3 = 688
                                }
                            }
                            var v44 uint8 = v30
                            _ = v44
                        }
                        _ = fi39
                    }
                    var v45 uint64 = v29
                    v2 = v34
                    _ = v45
                }
                var v46 float32 = -2.25
                if (true && (0.1 >= v46)) {
                    _ = len(v4)
                } else {
                    v47 := GenS2{F0: v2, F1: "hello", F2: int16(v1)}
                    _ = v47
                }
                _ = fi38
                _ = v46
            }
        case 882:
            if true {
                for fi48 := 0; fi48 < 20; fi48++ {
                    var v49 string = "foo bar"
                    _ = fi48
                    _ = v49
                }
                for fi50 := 0; fi50 < 13; fi50++ {
                    v1 -= v1
                    _ = fi50
                }
                var v51 float64 = 1e10
                _ = v51
            } else {
                var v52 uint16 = (v1 / ((32767) | 1))
                var v53 bool = (!true)
                _ = v52
                _ = v53
            }
        }
    }
    if (!(!("" >= ("hello" + "a")))) {
        var v54 int64 = -1
        _ = v54
    } else {
        delete(v4, 2147483647)
        var v55 int = 127
        _ = v55
    }
    var v56 uint64 = (v29 / ((18446744073709551615) | 1))
    for fi57 := 0; fi57 < 10; fi57++ {
        if false {
            if true {
                if false {
                    var v58 float32 = 1.5
                    fmt.Println(v2, v58, v34)
                    _ = v58
                } else {
                    v29 = (^(v29))
                }
                var v59 bool = false
                _ = v59
            } else {
                v60 := GenS5{F0: v1, F1: 42, F2: 18446744073709551615}
                _ = v60
            }
        } else {
            delete(v4, 0)
            v61 := v4[914]
            v62 := GenS6{F0: int8(v29), F1: -573, F2: (v30 ^ v30)}
            _ = v61
            _ = v62
        }
        if true {
            if (0.1 >= float64(((1.5 - (3.14159 * 0.0)) * 1e10))) {
                v33 += v1
                fmt.Println(v3)
                for fi63 := 0; fi63 < 5; fi63++ {
                    delete(v4, uint32(v56))
                    _ = fi63
                }
            }
            for fi64 := 0; fi64 < 10; fi64++ {
                if (0 != 127) {
                    v3 = (-(v3))
                    v1 = uint16(v30)
                    switch "\n\t" {
                    case "":
                        v33 -= 32767
                    case "foo bar":
                        v29 = (v56 >> ((0) & 63))
                    case "a":
                        var v65 uint16 = (v1 * 2)
                        _ = v65
                    }
                }
                var v66 float64 = 3.14159
                _ = fi64
                _ = v66
            }
        }
        _ = fi57
    }
    var v67 float32 = ((0.0 - 0.0) / -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v56:", v56)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v56
    _ = v67
}
