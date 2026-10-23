package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 bool; F2 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 int8; F1 uint16; F2 int16 }

func main() {
    v1 := []int{9223372036854775807, 128}
    var v2 uint16 = 249
    var v3 uint = 4294967295
    v4 := GenS1{F0: ("foo bar" + "\n\t"), F1: v3, F2: (v3 & 187)}
    var v5 uint16 = v2
    if true {
        for fi6 := 0; fi6 < 18; fi6++ {
            v1 = append(v1, (fi6 >> ((915) & 63)))
            var v7 uint32 = 4294967295
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 bool = true
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int = (fi10 >> (((^(v3))) & 63))
            v3 = (v3 >> ((255) & 63))
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v2, v2, v3)
    var v12 int8 = 1
    for fi13 := 0; fi13 < 10; fi13++ {
        v2 = v2
        v14 := GenS1{F0: "foo bar", F1: 930, F2: v3}
        _ = fi13
        _ = v14
    }
    v2 = 42
    if (true && false) {
        v1 = append(v1, 32767)
        switch "" {
        case "a":
            var v15 bool = false
            _ = v15
        case "hello":
            switch 32767 {
            case 9223372036854775807:
                for fi16 := 0; fi16 < 16; fi16++ {
                    v3 -= v3
                    _ = fi16
                }
            case 42:
                v17 := GenS1{F0: ((("a" + ("a" + (("" + "a") + (("foo bar" + (("hello" + "\n\t") + "hello")) + "hello")))) + ("" + ("a" + ((("hello" + "hello") + "") + ("hello" + ("\n\t" + "foo bar")))))) + "hello"), F1: v3, F2: v3}
                _ = v17
            case -1:
                v5 -= v5
            }
        case "":
            switch ("\n\t" + ("" + "hello")) {
            case "":
                v5 -= v2
            case "a":
                if false {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v19 := GenS2{F0: 255}
                        fmt.Println(fi18)
                        _ = fi18
                        _ = v19
                    }
                    fmt.Println(v3, v3)
                } else {
                    switch 1 {
                    case 32767:
                        fmt.Println(v3, v5, v3)
                    case 1:
                        switch "\n\t" {
                        case "":
                            if true {
                                for fi20 := 0; fi20 < 8; fi20++ {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v21 int = (fi20 + fi20)
                                        _ = v21
                                    case "":
                                        _ = len(v1)
                                    case "a":
                                        v1 = append(v1, fi20)
                                    default:
                                        _ = len(v1)
                                    }
                                    _ = fi20
                                }
                                var v22 bool = false
                                v23 := GenS3{F0: 715}
                                _ = v22
                                _ = v23
                            } else {
                                v24 := GenS1{F0: "", F1: v3, F2: (^(v3))}
                                v12 -= -24
                                v25 := GenS4{F0: -1}
                                _ = v24
                                _ = v25
                            }
                        case "a":
                            fmt.Println(v2, v5, v2)
                        default:
                            for fi26 := 0; fi26 < 17; fi26++ {
                                if true {
                                    if false {
                                        if false {
                                            v5 = v2
                                            for fi27 := 0; fi27 < 7; fi27++ {
                                                fmt.Println(fi26, fi27)
                                                v28 := GenS1{F0: "a", F1: (v3 ^ v3), F2: (v3 >> ((v3) & 63))}
                                                v5 = 32767
                                                _ = fi27
                                                _ = v28
                                            }
                                            if (3.14159 != (((3.14159 * (float32(3.14159) / float32((1.5 * 1e10)))) * 0.1) + 0.1)) {
                                                for fi29 := 0; fi29 < 14; fi29++ {
                                                    var v30 int = fi26
                                                    _ = fi29
                                                    _ = v30
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v31 := GenS1{F0: "hello", F1: 42, F2: v3}
                                                _ = v31
                                            }
                                        } else {
                                            if true {
                                                if (142 >= 4294967295) {
                                                    for fi32 := 0; fi32 < 13; fi32++ {
                                                        var v33 float32 = (0.1 / (3.14159 - 0.1))
                                                        v33 = ((1.5 / 1e10) / v33)
                                                        _ = len(v1)
                                                        _ = fi32
                                                        _ = v33
                                                    }
                                                    for fi34 := 0; fi34 < 16; fi34++ {
                                                        var v35 uint16 = 0
                                                        if ((1 > int64(fi26)) && (true && (0 <= 0))) {
                                                            var v36 uint16 = (^(v35))
                                                            _ = v36
                                                        } else {
                                                            var v37 uint16 = (v5 ^ (^(v5)))
                                                            var v38 int16 = -134
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                        _ = fi34
                                                        _ = v35
                                                    }
                                                } else {
                                                    if (true || ((v5 + v5) <= v5)) {
                                                        var v39 bool = false
                                                        _ = v39
                                                    } else {
                                                        v40 := GenS5{F0: 42}
                                                        var v41 uint64 = 266
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                }
                                                var v42 int32 = int32(v3)
                                                if true {
                                                    v12 += (v12 >> (((^(v3))) & 63))
                                                    v1 = append(v1, fi26)
                                                }
                                                _ = v42
                                            }
                                        }
                                        v43 := GenS6{F0: (^(v3)), F1: false, F2: (("\n\t" + "a") + "a")}
                                        _ = v43
                                    }
                                    v12 -= 125
                                }
                                _ = fi26
                            }
                        }
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            if (v3 < (v3 % (((v3 << (((v3 + 2)) & 63))) | 1))) {
                                v1 = append(v1, 32768)
                                for fi44 := 0; fi44 < 15; fi44++ {
                                    var v45 uint = 350
                                    for fi46 := 0; fi46 < 17; fi46++ {
                                        v47 := GenS7{F0: 504}
                                        _ = fi46
                                        _ = v47
                                    }
                                    for fi48 := 0; fi48 < 5; fi48++ {
                                        v49 := GenS8{F0: true, F1: ""}
                                        _ = fi48
                                        _ = v49
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                                for fi50 := 0; fi50 < 13; fi50++ {
                                    var v51 string = ""
                                    if true {
                                        v52 := GenS6{F0: (v3 + (^(v3))), F1: (true && false), F2: ((v51 + v51) + ("" + ("foo bar" + ((v51 + "\n\t") + (v51 + "foo bar")))))}
                                        _ = v52
                                    }
                                    _ = fi50
                                    _ = v51
                                }
                            }
                        case "hello":
                            for fi53 := 0; fi53 < 12; fi53++ {
                                var v54 uint8 = 143
                                var v55 uint8 = v54
                                if ((v3 >> (((v3 / (((v3 >> ((v3) & 63))) | 1))) & 63)) >= (v3 << (((v3 << ((1) & 63))) & 63))) {
                                    v12 = (v12 % (((v12 / (((v12 + v12)) | 1))) | 1))
                                }
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            default:
                fmt.Println(v3)
            }
        default:
            var v56 int8 = v12
            _ = v56
        }
    } else {
        switch 1 {
        case 1:
            v57 := GenS8{F0: (1 == 42), F1: "foo bar"}
            _ = v57
        case 0:
            fmt.Println(v12)
        case 9223372036854775807:
            var v58 int8 = v12
            _ = v58
        }
    }
    for fi59 := 0; fi59 < 18; fi59++ {
        for fi60 := 0; fi60 < 16; fi60++ {
            fmt.Println(v2, v12)
            fmt.Println(v2)
            var v61 int8 = -124
            _ = fi60
            _ = v61
        }
        _ = fi59
    }
    v62 := GenS1{F0: "\n\t", F1: v3, F2: v3}
    v3 -= (v3 >> (((v3 % (((v3 / (((v3 % ((924) | 1))) | 1))) | 1))) & 63))
    fmt.Println(v2, v2, v12)
    v63 := GenS9{F0: v12, F1: (^(v2)), F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
}
