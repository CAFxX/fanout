package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int8; F1 int; F2 int32 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32; F1 uint; F2 int }
type GenS8 struct { F0 bool; F1 int32; F2 uint32 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 int64 }
func genFn0(p0 int64) string {
    var v6 uint64 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7, p0)
        v8 := GenS1{F0: 1, F1: (3.14159 <= (1.5 + float64(1e10)))}
        fmt.Println(v6)
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: -561, F1: 0}
    v10 := GenS3{F0: -1, F1: 514}
    var v11 string = ("hello" + (("\n\t" + "hello") + "a"))
    _ = p0
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int64 = -380
    var v2 uint = 183
    v3 := make(map[string]bool)
    v4 := []uint32{65535, 252}
    var v5 uint16 = 32767
    for fi12 := 0; fi12 < 2; fi12++ {
        if len(v4) > 0 {
            v4[0] = 0
        }
        _ = fi12
    }
    v2 = v2
    v5 = 65535
    for fi13 := 0; fi13 < 1; fi13++ {
        v1 += v1
        _ = fi13
    }
    if ((!true) || (!true)) {
        switch "" {
        case "a":
            switch 128 {
            case -1:
                for fi14 := 0; fi14 < 15; fi14++ {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 int8 = 0
                        v17 := GenS2{F0: v1, F1: -590}
                        switch fi15 {
                        case 128:
                            v5 = v5
                        case 127:
                            fmt.Println(fi15)
                        case -2147483648:
                            v4 = append(v4, 2147483647)
                        default:
                            v18 := GenS3{F0: 1, F1: (fi15 * -9223372036854775808)}
                            _ = v18
                        }
                        v16 = v16
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    fmt.Println(v2, fi14, v5)
                    v19 := GenS2{F0: v1, F1: 1}
                    _ = fi14
                    _ = v19
                }
            case -2147483648:
                for fi20 := 0; fi20 < 1; fi20++ {
                    if true {
                        var v21 int8 = 70
                        switch ((("foo bar" + "a") + ((("foo bar" + "foo bar") + "") + (((("a" + ("hello" + "foo bar")) + "hello") + "") + "a"))) + "") {
                        case "":
                            var v22 int16 = 1
                            _ = v22
                        }
                        _ = v21
                    } else {
                        if ((false && true) && true) {
                            v1 = -1
                            delete(v3, ("" + ""))
                            var v23 float64 = -2.25
                            _ = v23
                        } else {
                            if true {
                                var v24 string = ("" + "")
                                _ = v24
                            } else {
                                v1 = v1
                                fmt.Println(v1, v2, v5)
                                v2 += 1
                            }
                            v5 -= (v5 << (((v2 ^ v2)) & 63))
                            switch "\n\t" {
                            case "":
                                for fi25 := 0; fi25 < 19; fi25++ {
                                    v26 := GenS4{F0: 0, F1: -9223372036854775808, F2: 1}
                                    _ = fi25
                                    _ = v26
                                }
                            default:
                                var v27 float32 = -2.25
                                _ = v27
                            }
                        }
                    }
                    _ = fi20
                }
            case 9223372036854775807:
                _ = len(v3)
            }
        case "":
            v5 -= 239
        case "foo bar":
            for fi28 := 0; fi28 < 16; fi28++ {
                if true {
                    v2 = (v2 | (v2 / (((v2 << ((4294967295) & 63))) | 1)))
                    v29 := GenS1{F0: 2, F1: ((false && ("\n\t" <= "a")) || true)}
                    var v30 int32 = 285
                    _ = v29
                    _ = v30
                }
                var v31 string = ""
                _ = fi28
                _ = v31
            }
        default:
            fmt.Println(v5, v2, v5)
        }
    } else {
        v32 := GenS2{F0: (v1 % ((-9223372036854775808) | 1)), F1: 880}
        _ = v32
    }
    if ((^(v5)) < (v5 >> ((21) & 63))) {
        v2 = v2
    }
    var v33 uint64 = 697
    var v34 uint16 = (v5 | v5)
    switch ("foo bar" + "a") {
    case "\n\t":
        if ((false || false) && false) {
            var v35 uint64 = v33
            v36 := GenS5{F0: ("" + "a"), F1: (v35 + 101)}
            _ = v35
            _ = v36
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v3["a"] = (2 == 98)
    for fi37 := 0; fi37 < 11; fi37++ {
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 float64 = (1e10 * 0.1)
            v1 -= (v1 << ((v2) & 63))
            v33 = (v33 - v33)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch "" {
            case "hello":
                switch "foo bar" {
                case "hello":
                    v34 += v34
                case "a":
                    v5 += 32767
                case "foo bar":
                    if (!(!(127 == -1))) {
                        switch (fi37 ^ (fi38 + (fi38 - fi37))) {
                        case -756:
                            v4 = append(v4, 777)
                        case 128:
                            v5 += (^(v5))
                        default:
                            v40 := GenS6{F0: fi38}
                            _ = v40
                        }
                    }
                default:
                    v5 = v5
                }
            case "":
                if false {
                    v41 := GenS7{F0: 42, F1: (^(v2)), F2: (fi37 | (fi37 | 489))}
                    _ = len(v3)
                    _ = len(v3)
                    _ = v41
                } else {
                    fmt.Println(v2)
                    v42 := GenS7{F0: 65535, F1: v2, F2: fi38}
                    switch (fi37 ^ (fi38 - (fi37 << ((127) & 63)))) {
                    case 2147483647:
                        switch ((("" + "foo bar") + "foo bar") + "\n\t") {
                        case "hello":
                            v4 = append(v4, 2147483648)
                        case "":
                            v43 := GenS3{F0: 32767, F1: (fi37 & int(v34))}
                            _ = v43
                        }
                    }
                    var v44 uint8 = 42
                    var v45 string = "foo bar"
                    _ = v42
                    _ = v44
                    _ = v45
                }
            case "foo bar":
                switch "hello" {
                case "a":
                    if ((v2 * v2) != (v2 + 18446744073709551615)) {
                        v46 := GenS6{F0: 42}
                        v47 := GenS8{F0: (true || true), F1: int32(v34), F2: 1}
                        _ = v46
                        _ = v47
                    } else {
                        fmt.Println(v34, fi37)
                        v1 -= -9223372036854775808
                        v48 := GenS1{F0: 286, F1: (!false)}
                        _ = v48
                    }
                case "hello":
                    fmt.Println(v1, fi37, fi37)
                case "":
                    var v49 float32 = -2.25
                    _ = v49
                }
                var v50 bool = true
                v34 += (v5 / ((v34) | 1))
                _ = v50
            default:
                var v51 int16 = int16(fi38)
                _ = v51
            }
            if false {
                v33 -= (v33 & (v33 | v33))
                switch -9223372036854775808 {
                case -155:
                    var v52 uint8 = uint8(v33)
                    _ = v52
                case 128:
                    v2 -= v2
                    v4 = append(v4, 79)
                    var v53 uint16 = (^(v5))
                    _ = v53
                case -2147483648:
                    v54 := GenS4{F0: 1, F1: -2147483648, F2: 297}
                    _ = v54
                default:
                    if true {
                        var v55 int8 = 100
                        var v56 uint8 = 38
                        _ = v55
                        _ = v56
                    } else {
                        v1 = 1
                        fmt.Println(v1)
                        var v57 int32 = 788
                        _ = v57
                    }
                    v39 = v39
                    var v58 bool = true
                    _ = v58
                }
            } else {
                var v59 uint64 = v33
                if true {
                    v59 = v33
                    var v60 int8 = -128
                    v61 := GenS9{F0: 4294967295, F1: 1}
                    _ = v60
                    _ = v61
                } else {
                    var v62 uint16 = v5
                    var v63 uint16 = v62
                    var v64 uint8 = uint8(v59)
                    _ = v62
                    _ = v63
                    _ = v64
                }
                _ = v59
            }
            v65 := GenS6{F0: (fi38 + fi37)}
            _ = fi38
            _ = v39
            _ = v65
        }
        if false {
            v66 := GenS10{F0: v1}
            _ = v66
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi37
    }
    var v67 int64 = v1
    fmt.Println(v2)
    v68 := GenS1{F0: 140, F1: true}
    v67 = 50
    if (0.1 == (-2.25 + 3.14159)) {
        v34 -= v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v67
    _ = v68
}
