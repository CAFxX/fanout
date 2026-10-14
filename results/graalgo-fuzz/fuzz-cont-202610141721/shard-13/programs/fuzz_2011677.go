package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 int16; F2 int16 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 int16; F1 uint64 }

func main() {
    v1 := []string{"a", "", ("foo bar" + (("hello" + ("" + "hello")) + "hello")), "\n\t"}
    v2 := []int{}
    var v3 float32 = 3.14159
    v4 := make(map[int]string)
    var v5 int64 = -9223372036854775808
    v6 := GenS1{F0: (v5 << ((377) & 63)), F1: true, F2: 4294967295}
    fmt.Println(v5)
    if (-595 > 549) {
        fmt.Println(v3, v3)
        v7 := GenS2{F0: 639, F1: (((("\n\t" + "\n\t") + (("hello" + "a") + ("a" + (("hello" + "\n\t") + "foo bar")))) + "\n\t") + "")}
        v3 = v3
        _ = v7
    } else {
        if (-24 <= 42) {
            var v8 int8 = 0
            _ = v8
        }
        for fi9 := 0; fi9 < 12; fi9++ {
            fmt.Println(v3, v3)
            switch ("" + "\n\t") {
            case "hello":
                if (42 < 1) {
                    fmt.Println(fi9, v5, v3)
                    var v10 uint32 = 2147483647
                    var v11 uint16 = 812
                    _ = v10
                    _ = v11
                } else {
                    var v12 int = -9223372036854775808
                    _ = v12
                }
            default:
                var v13 float64 = 1e10
                _ = v13
            }
            _ = fi9
        }
    }
    if (100 == 77) {
        if (v3 >= v3) {
            var v14 int = 42
            v15 := GenS3{F0: 0, F1: 1}
            _ = v14
            _ = v15
        } else {
            v16 := GenS2{F0: 128, F1: "foo bar"}
            if (((103 > -380) && ((1 == 589) && false)) || true) {
                v17 := GenS1{F0: -1, F1: (0 <= -653), F2: 0}
                var v18 uint = 1
                _ = v17
                _ = v18
            } else {
                v3 = ((3.14159 / (float32(1.5) * ((v3 * (0.0 / 1.0)) - 3.14159))) + (v3 / v3))
            }
            _ = v16
        }
        v5 += (v5 | 42)
        v19 := GenS4{F0: (!(!(!true)))}
        _ = v19
    } else {
        fmt.Println(v3)
        var v20 int8 = 100
        _ = v20
    }
    fmt.Println(v5)
    if false {
        var v21 int64 = 9223372036854775807
        v22 := GenS5{F0: (626 != 4294967295)}
        var v23 uint32 = 65535
        _ = v21
        _ = v22
        _ = v23
    } else {
        v5 -= (v5 << ((255) & 63))
        var v24 bool = false
        fmt.Println(v5, v5, v5)
        _ = v24
    }
    var v25 uint8 = 127
    switch -1 {
    case -986:
        v3 -= (float32((0.0 * (1.5 - (1.5 / ((1e10 / (-2.25 / -2.25)) * 1.5))))) - v3)
    case -234:
        v3 -= (0.1 + v3)
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if (false && false) {
            switch "\n\t" {
            case "foo bar":
                if true {
                    for fi27 := 0; fi27 < 9; fi27++ {
                        var v28 string = "foo bar"
                        v25 = (v25 << ((1) & 63))
                        switch fi27 {
                        case -975:
                            v29 := GenS6{F0: 701, F1: 1, F2: 1}
                            _ = v29
                        case -9223372036854775808:
                            v28 = v28
                        default:
                            if true {
                                v2 = append(v2, 0)
                            } else {
                                v5 = v5
                            }
                        }
                        _ = fi27
                        _ = v28
                    }
                    v25 -= v25
                    if (!(1 == 1)) {
                        v5 = (v5 << ((1) & 63))
                    } else {
                        v30 := GenS7{F0: 0, F1: 647}
                        _ = v30
                    }
                } else {
                    var v31 int32 = 0
                    _ = v31
                }
            default:
                for fi32 := 0; fi32 < 8; fi32++ {
                    v5 -= -588
                    if (true || ((v25 - v25) <= 55)) {
                        if false {
                            v4[fi26] = ("\n\t" + (("foo bar" + "a") + ("foo bar" + "\n\t")))
                            if false {
                                v33 := GenS8{F0: (fi32 - fi26), F1: -128}
                                switch ("hello" + "a") {
                                case "hello":
                                    v3 += 0.1
                                default:
                                    var v34 int8 = -1
                                    _ = v34
                                }
                                v5 = (v5 >> ((699) & 63))
                                v25 -= 255
                                v4[fi32] = (("foo bar" + "hello") + ("foo bar" + ("foo bar" + "\n\t")))
                                v25 = v25
                                _ = v33
                            }
                        } else {
                            var v35 uint = 812
                            v5 -= (v5 - (v5 >> ((v35) & 63)))
                            _ = v35
                        }
                        v3 = 0.0
                        v36 := GenS8{F0: (fi26 | (fi32 * fi32)), F1: 42}
                        _ = v36
                    } else {
                        var v37 int8 = 42
                        if ("a" >= "a") {
                            _ = len(v2)
                        }
                        _ = v37
                    }
                    fmt.Println(v3)
                    _ = fi32
                }
            }
            v25 = v25
            fmt.Println(v25)
        }
        v38 := GenS9{F0: 0, F1: 42}
        _ = fi26
        _ = v38
    }
    v39 := GenS1{F0: (v5 | v5), F1: (!true), F2: 739}
    v40 := GenS9{F0: 32767, F1: 759}
    v3 = v3
    if false {
        var v41 int8 = 1
        _ = v41
    } else {
        var v42 bool = (!false)
        v3 = (v3 + v3)
        _ = v42
    }
    switch "\n\t" {
    case "foo bar":
        switch ("" + ("hello" + ("a" + "foo bar"))) {
        case "hello":
            if len(v2) > 0 {
                v2[0] = 32768
            }
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            for fi43 := 0; fi43 < 17; fi43++ {
                v3 += ((v3 * (((v3 * v3) + -2.25) + 3.14159)) + 0.0)
                _ = fi43
            }
        default:
            var v44 float64 = 0.1
            _ = v44
        }
    default:
        var v45 int16 = 42
        _ = v45
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    if (805 != 0) {
        var v46 int8 = 88
        _ = v46
    } else {
        if false {
            v25 = (v25 << ((1) & 63))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            switch 32768 {
            case 1:
                v5 = v5
            case 128:
                var v47 uint8 = 127
                _ = v47
            case -9223372036854775808:
                delete(v4, 32768)
            default:
                var v48 int32 = -1
                _ = v48
            }
            var v49 float32 = v3
            v25 -= 42
            _ = v49
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
