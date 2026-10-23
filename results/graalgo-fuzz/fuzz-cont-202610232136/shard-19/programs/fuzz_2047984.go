package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 string; F1 uint8; F2 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"foo bar", "\n\t", "\n\t", "foo bar"}
    var v3 int = 0
    v4 := make(map[int]string)
    v3 -= v3
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS1{F0: 2, F1: 0}
            _ = fi6
            _ = v7
        }
        fmt.Println(v3, fi5)
        switch (v3 >> ((2) & 63)) {
        case -213:
            for fi8 := 0; fi8 < 17; fi8++ {
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = len(v1)
                _ = fi8
            }
        }
        _ = fi5
    }
    var v9 uint16 = 255
    switch ("hello" + "a") {
    case "a":
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int32 = -2147483648
            _ = fi10
            _ = v11
        }
    case "foo bar":
        fmt.Println(v3, v9)
    case "":
        v12 := GenS2{F0: 1, F1: 42}
        _ = v12
    default:
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS3{F0: 508, F1: 4294967295, F2: -1}
            _ = fi13
            _ = v14
        }
    }
    v15 := GenS4{F0: 32767}
    fmt.Println(v9, v3)
    v9 += 678
    var v16 int32 = 1
    if (1.5 > 0.1) {
        fmt.Println(v3, v16, v16)
        if true {
            if false {
                if (false && false) {
                    fmt.Println(v16, v9, v9)
                    v17 := GenS5{F0: 0, F1: 686}
                    v9 += v9
                    _ = v17
                }
                v18 := GenS6{F0: 0, F1: v3}
                fmt.Println(v9, v9, v9)
                _ = v18
            } else {
                v3 = (v3 | 42)
                fmt.Println(v9)
                v9 = (v9 | 42)
            }
            v16 = (v16 * (v16 - (v16 << ((2) & 63))))
        } else {
            if (32767 <= v9) {
                if (v3 != 1) {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v3 -= (v3 ^ v3)
                        _ = fi19
                    }
                    if (50 <= -9223372036854775808) {
                        fmt.Println(v9, v3)
                        v16 = v16
                        v3 = v3
                    }
                }
            }
            if true {
                switch "a" {
                case "hello":
                    v16 -= v16
                }
            } else {
                v1[2147483648] = "\n\t"
            }
        }
        if true {
            if (!(0.0 >= -2.25)) {
                if (!true) {
                    delete(v4, v3)
                    var v20 int64 = int64(v16)
                    for fi21 := 0; fi21 < 19; fi21++ {
                        var v22 int64 = -1
                        var v23 uint8 = 1
                        v3 = (-(fi21))
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    _ = v20
                }
            }
        } else {
            v3 = (v3 << ((0) & 63))
            var v24 float32 = (0.1 / -2.25)
            _ = v24
        }
    }
    v3 = (v3 & v3)
    var v25 uint32 = 42
    var v26 int8 = 1
    v27 := GenS6{F0: -1, F1: 701}
    switch v3 {
    case 42:
        fmt.Println(v9, v26, v9)
    case 880:
        _ = len(v1)
    case 9223372036854775807:
        var v28 string = "\n\t"
        _ = v28
    default:
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := v4[1]
            v31 := GenS2{F0: 1, F1: 686}
            var v32 int64 = 42
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
    }
    fmt.Println(v9, v3)
    if (true && true) {
        var v33 string = "\n\t"
        var v34 float32 = 0.1
        for fi35 := 0; fi35 < 14; fi35++ {
            v1[42] = "hello"
            if true {
                if (!((^(v25)) == (v25 / (((v25 << ((uint(v25)) & 63))) | 1)))) {
                    if ((false || (!(540 < 42))) || false) {
                        v36 := GenS7{F0: v33, F1: 255, F2: true}
                        v34 = ((v34 / float32(((0.1 + ((1.5 / (0.1 * 0.1)) - -2.25)) / -2.25))) / ((v34 - (0.1 / 1.5)) / v34))
                        for fi37 := 0; fi37 < 15; fi37++ {
                            if false {
                                fmt.Println(fi35, v34, v33)
                                if (!true) {
                                    v38 := GenS8{F0: v26}
                                    if ((1000 != 466) && (1 >= 2)) {
                                        v26 += (v26 - -67)
                                        var v39 uint16 = (v9 * (v9 & 0))
                                        fmt.Println(v16, fi37, v16)
                                        fmt.Println(v16, v25, v25)
                                        _ = v39
                                    }
                                    _ = v38
                                }
                            } else {
                                var v40 string = v33
                                if (1 != (v9 + v9)) {
                                    v9 += (^(v9))
                                    v34 += (((float32(-2.25) - 1e10) - 1e10) + -2.25)
                                }
                                _ = v40
                            }
                            fmt.Println(v16, fi35)
                            _ = fi37
                        }
                        _ = v36
                    } else {
                        v41 := GenS3{F0: -826, F1: 255, F2: (v26 & 14)}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v16 += 39
                        _ = v41
                    }
                    v2 = append(v2, "")
                } else {
                    if ((v16 ^ v16) != (v16 - (v16 / (((v16 - -197)) | 1)))) {
                        var v42 int32 = (v16 << ((2) & 63))
                        for fi43 := 0; fi43 < 10; fi43++ {
                            var v44 int = fi43
                            _ = fi43
                            _ = v44
                        }
                        v16 += v42
                        _ = v42
                    }
                }
            }
            switch "hello" {
            case "foo bar":
                _ = len(v4)
            case "":
                v1[998] = (((v33 + ("a" + (((v33 + v33) + (v33 + "foo bar")) + "hello"))) + ((((("a" + v33) + v33) + "") + v33) + v33)) + ("" + "a"))
            default:
                if (!true) {
                    for fi45 := 0; fi45 < 4; fi45++ {
                        switch v33 {
                        case "a":
                            if true {
                                switch "" {
                                case "a":
                                    if true {
                                        v2 = append(v2, "a")
                                        var v46 uint16 = 933
                                        _ = v46
                                    }
                                    var v47 uint32 = (^(v25))
                                    var v48 float32 = -2.25
                                    _ = v47
                                    _ = v48
                                case "foo bar":
                                    _ = len(v2)
                                default:
                                    var v49 int16 = -32768
                                    v50 := GenS9{F0: v9, F1: 4294967295}
                                    fmt.Println(v25, fi45, v9)
                                    _ = v49
                                    _ = v50
                                }
                                v3 = (fi45 + fi35)
                                fmt.Println(v3, fi35)
                            } else {
                                var v51 bool = false
                                _ = v51
                            }
                        case "\n\t":
                            v52 := GenS10{F0: (v25 | v25)}
                            _ = v52
                        case "foo bar":
                            v26 = (v26 | v26)
                        default:
                            v53 := GenS8{F0: v26}
                            _ = v53
                        }
                        _ = fi45
                    }
                    for fi54 := 0; fi54 < 5; fi54++ {
                        fmt.Println(v3, fi54)
                        var v55 uint32 = (v25 | (^(v25)))
                        _ = fi54
                        _ = v55
                    }
                    delete(v1, (v25 >> ((uint(v25)) & 63)))
                } else {
                    v33 = ("foo bar" + (((((v33 + ("" + v33)) + ((v33 + "hello") + ((v33 + (v33 + (v33 + (v33 + ((((v33 + v33) + (v33 + "a")) + "hello") + v33))))) + v33))) + "foo bar") + "\n\t") + v33))
                }
            }
            _ = fi35
        }
        _ = v33
        _ = v34
    }
    v56 := v4[v3]
    var v57 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v9
}
