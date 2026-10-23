package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int8 }
type GenS5 struct { F0 int16; F1 int16; F2 uint16 }
type GenS6 struct { F0 bool; F1 string; F2 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int8; F1 bool }
func genFn0(p0 uint16, p1 int8, p2 bool) string {
    v7 := []int{-1, 2147483647, -822}
    v8 := GenS1{F0: (("" + "\n\t") + "a"), F1: -533, F2: 255}
    _ = len(v7)
    switch "hello" {
    case "\n\t":
        fmt.Println(p2, p1)
    case "hello":
        switch -311 {
        case 2147483647:
            v9 := GenS1{F0: "\n\t", F1: 42, F2: 839}
            _ = v9
        case 0:
            var v10 float64 = ((0.1 - 0.0) / (((1.5 * 0.0) - -2.25) - 1.5))
            _ = v10
        case 9223372036854775807:
            if p2 {
                if len(v7) > 0 {
                    v7[0] = -581
                }
            } else {
                var v11 string = ("" + "")
                var v12 int64 = -684
                var v13 int = -425
                _ = v11
                _ = v12
                _ = v13
            }
        default:
            fmt.Println(p2)
        }
    case "a":
        p0 += 0
    default:
        if (!p2) {
            var v14 int32 = -2147483648
            var v15 float32 = (float32(1.5) + (3.14159 / -2.25))
            _ = v14
            _ = v15
        }
    }
    var v16 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v7
    _ = v8
    return (("foo bar" + "a") + (("foo bar" + ("a" + "hello")) + (("a" + "") + ("" + "hello"))))
}
func genFn1() int8 {
    var v17 int8 = -1
    if (true || (554 >= 42)) {
        fmt.Println(v17)
        v18 := GenS2{F0: 255, F1: 761}
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                v17 = (v17 >> ((280) & 63))
                if ((0 < 0) && (128 != 217)) {
                    switch (-(fi19)) {
                    case 32767:
                        v21 := GenS3{F0: 495, F1: 0, F2: 2}
                        _ = v21
                    case 1:
                        var v22 float32 = (0.1 * ((1.5 / 3.14159) - -2.25))
                        _ = v22
                    }
                    v17 -= (^(v17))
                }
                _ = fi20
            }
            var v23 float32 = 1e10
            _ = fi19
            _ = v23
        }
        _ = v18
    } else {
        switch ("a" + "") {
        case "":
            v24 := []int{-287, 2147483647}
            _ = v24
        default:
            fmt.Println(v17)
        }
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 int8 = v17
            v27 := []int{(fi25 | (fi25 >> ((2) & 63))), fi25, fi25}
            fmt.Println(v17, fi25, v26)
            _ = fi25
            _ = v26
            _ = v27
        }
        if true {
            v17 -= 100
        }
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS1{F0: "", F1: 1000, F2: 65535}
        _ = fi28
        _ = v29
    }
    fmt.Println(v17)
    _ = v17
    return v17
}

func main() {
    var v1 float32 = (0.0 / (-2.25 * (3.14159 / 1.0)))
    var v2 uint8 = 115
    var v3 int = 32767
    v4 := []uint32{4294967295, 699, 2147483648}
    var v5 float64 = float64(v1)
    v6 := make(map[uint32]uint32)
    v30 := GenS4{F0: 737, F1: -1, F2: 127}
    v5 = 3.14159
    var v31 uint8 = 1
    if true {
        fmt.Println(v3)
        switch "\n\t" {
        case "a":
            if false {
                v32 := GenS5{F0: 1000, F1: 0, F2: 317}
                _ = v32
            } else {
                var v33 uint32 = 0
                var v34 int64 = 42
                var v35 float64 = (v5 * v5)
                _ = v33
                _ = v34
                _ = v35
            }
        case "":
            fmt.Println(v2, v2, v1)
        }
        delete(v6, 1)
    }
    var v36 float64 = v5
    v5 += float64((float32(v5) * 0.1))
    v37 := GenS6{F0: (v5 != v36), F1: (("a" + ((("hello" + (((("a" + "foo bar") + "foo bar") + ("foo bar" + ((("a" + "hello") + "\n\t") + ""))) + "hello")) + "") + "foo bar")) + (("" + "hello") + "\n\t")), F2: -41}
    v38 := GenS1{F0: "a", F1: 32767, F2: 42}
    for fi39 := 0; fi39 < 1; fi39++ {
        v1 = float32((((((v36 * ((v36 - v36) / v5)) * v36) * ((v36 / 1.0) / 1.0)) + (v36 + float64((0.1 - v1)))) / (v5 / v5)))
        switch "" {
        case "":
            v1 = v1
        case "hello":
            switch "a" {
            case "\n\t":
                _ = len(v4)
            case "hello":
                for fi40 := 0; fi40 < 13; fi40++ {
                    var v41 uint8 = 2
                    var v42 uint = 489
                    v31 = v41
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
        }
        if ((-2.25 >= 3.14159) || (v2 < 205)) {
            switch (fi39 | 1) {
            case 1:
                v4 = append(v4, 0)
            case 2147483647:
                var v43 uint16 = 0
                _ = v43
            default:
                v4 = append(v4, 65535)
            }
        }
        _ = fi39
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        v45 := v6[934]
        _ = fi44
        _ = v45
    }
    if true {
        v2 -= v2
        if len(v4) > 0 {
            v4[0] = 2
        }
        for fi46 := 0; fi46 < 8; fi46++ {
            v1 -= 0.0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi46
        }
    } else {
        for fi47 := 0; fi47 < 7; fi47++ {
            fmt.Println(v5, v5, v31)
            _ = fi47
        }
        for fi48 := 0; fi48 < 10; fi48++ {
            switch (("" + "") + "hello") {
            case "a":
                v31 += v2
            case "hello":
                v49 := GenS7{F0: (v3 - (fi48 & 790))}
                _ = v49
            case "":
                if (!(!false)) {
                    if (!true) {
                        v50 := GenS4{F0: 735, F1: -9223372036854775808, F2: 100}
                        fmt.Println(v31, v1)
                        _ = v50
                    } else {
                        v51 := GenS5{F0: -1, F1: 42, F2: 0}
                        var v52 int64 = 1
                        v53 := GenS8{F0: -1}
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                } else {
                    if (((792 >= 42) && false) || true) {
                        if (false && false) {
                            for fi54 := 0; fi54 < 14; fi54++ {
                                _ = len(v6)
                                var v55 uint = 1
                                switch v3 {
                                case 2147483647:
                                    v56 := v6[2147483647]
                                    _ = v56
                                case 0:
                                    var v57 uint8 = 31
                                    _ = v57
                                case -1:
                                    v2 = v31
                                default:
                                    var v58 uint8 = 116
                                    _ = v58
                                }
                                _ = fi54
                                _ = v55
                            }
                            for fi59 := 0; fi59 < 13; fi59++ {
                                v3 = fi59
                                var v60 uint = 1
                                _ = fi59
                                _ = v60
                            }
                        }
                        for fi61 := 0; fi61 < 2; fi61++ {
                            _ = len(v4)
                            _ = fi61
                        }
                    }
                }
            default:
                if (!((!(!(-32768 >= 32767))) && false)) {
                    v62 := GenS9{F0: genFn1(), F1: true}
                    _ = v62
                } else {
                    fmt.Println(v5, v31)
                }
            }
            v5 = v5
            if true {
                v1 += 0.1
                v31 = 95
            } else {
                for fi63 := 0; fi63 < 4; fi63++ {
                    v31 = v31
                    v64 := GenS4{F0: 792, F1: 42, F2: -1}
                    v5 += float64(v1)
                    switch ("foo bar" + "a") {
                    case "":
                        v2 += v2
                    default:
                        fmt.Println(v36)
                    }
                    switch 128 {
                    case 1:
                        var v65 string = ((((("a" + "hello") + "hello") + "\n\t") + "a") + "")
                        _ = v65
                    }
                    _ = fi63
                    _ = v64
                }
            }
            _ = fi48
        }
        var v66 uint8 = 1
        _ = v66
    }
    var v67 uint8 = v31
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v67
}
