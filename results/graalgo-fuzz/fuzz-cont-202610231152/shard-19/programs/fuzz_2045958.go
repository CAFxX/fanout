package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS9 struct { F0 uint; F1 int32 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]int)
    v3 := []string{("foo bar" + "foo bar"), "\n\t"}
    var v4 uint64 = 401
    var v5 int32 = 809
    v6 := make(map[string]bool)
    if false {
        if len(v3) > 0 {
            v3[0] = ("hello" + ("foo bar" + "foo bar"))
        }
        switch "foo bar" {
        case "a":
            v4 = (v4 % ((v4) | 1))
        case "foo bar":
            for fi7 := 0; fi7 < 15; fi7++ {
                switch 150 {
                case 9223372036854775807:
                    v8 := GenS1{F0: true, F1: 37, F2: 255}
                    _ = v8
                default:
                    var v9 uint8 = 2
                    _ = v9
                }
                var v10 uint32 = 2147483647
                v1 += 3.14159
                _ = fi7
                _ = v10
            }
        default:
            v11 := GenS2{F0: 42, F1: 11}
            _ = v11
        }
    }
    var v12 bool = true
    v13 := GenS3{F0: ((v4 ^ v4) >= v4)}
    v12 = v12
    switch 1 {
    case -9223372036854775808:
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                switch "\n\t" {
                case "hello":
                    var v16 int = fi15
                    _ = v16
                case "a":
                    v12 = ((true && (v12 || false)) && false)
                case "\n\t":
                    fmt.Println(fi15, v1, v5)
                default:
                    if true {
                        v17 := GenS4{F0: -667, F1: 174}
                        if (("hello" + (("hello" + (("hello" + "") + "\n\t")) + ("" + ("a" + "foo bar")))) == "\n\t") {
                            v18 := GenS5{F0: 396, F1: 2, F2: fi15}
                            _ = v18
                        } else {
                            v19 := GenS6{F0: (v5 + v5), F1: 65535}
                            _ = v19
                        }
                        v12 = (v12 || false)
                        if v12 {
                            if ((v4 == v4) && true) {
                                var v20 int8 = 42
                                if true {
                                    v21 := GenS2{F0: 119, F1: 2}
                                    var v22 int64 = 42
                                    _ = v21
                                    _ = v22
                                } else {
                                    v5 -= (v5 - v5)
                                }
                                var v23 int = fi15
                                _ = v20
                                _ = v23
                            } else {
                                switch fi14 {
                                case -451:
                                    v24 := GenS3{F0: (!v12)}
                                    _ = v24
                                }
                                v12 = (((!v12) && (1 < 42)) || ((!(!((v12 || (1 >= 42)) && ((4294967295 >= 1) && true)))) || (false && v12)))
                            }
                            var v25 int8 = 0
                            if len(v3) > 0 {
                                v3[0] = ("foo bar" + "")
                            }
                            _ = v25
                        } else {
                            v5 -= v5
                            v26 := GenS7{F0: 127}
                            _ = v26
                        }
                        if (!v12) {
                            v27 := GenS8{F0: -300, F1: 32767, F2: 871}
                            var v28 string = ""
                            v5 = (^(v5))
                            _ = v27
                            _ = v28
                        }
                        v5 -= (v5 + 2147483647)
                        _ = v17
                    }
                }
                if true {
                    fmt.Println(v1, v12)
                    if (!v12) {
                        if v12 {
                            delete(v2, (^(fi15)))
                            var v29 float64 = (0.0 + float64(((((0.1 / (1.5 + v1)) / (v1 - (0.0 + (-2.25 + v1)))) * v1) * (1.5 * (-2.25 + v1)))))
                            switch (-(fi14)) {
                            case 32767:
                                var v30 float64 = 0.0
                                _ = v30
                            }
                            _ = v29
                        } else {
                            var v31 string = "foo bar"
                            var v32 float64 = 0.1
                            v31 = "hello"
                            _ = v31
                            _ = v32
                        }
                        switch (fi15 >> ((255) & 63)) {
                        case -14:
                            _ = len(v3)
                        default:
                            v4 = v4
                        }
                    }
                    fmt.Println(v12, v4)
                    v4 = 576
                }
                _ = fi15
            }
            v33 := GenS7{F0: 255}
            switch "a" {
            case "hello":
                v34 := GenS2{F0: 1, F1: 227}
                _ = v34
            default:
                var v35 uint8 = 2
                _ = v35
            }
            _ = fi14
            _ = v33
        }
    case 32767:
        v6[""] = v12
    default:
        fmt.Println(v12, v1, v12)
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v5 = (v5 ^ (v5 % (((v5 >> ((2) & 63))) | 1)))
        if true {
            for fi37 := 0; fi37 < 4; fi37++ {
                v38 := v6["\n\t"]
                _ = fi37
                _ = v38
            }
            v39 := GenS9{F0: 65535, F1: v5}
            fmt.Println(v12, v5, v5)
            _ = v39
        }
        var v40 uint = 588
        _ = fi36
        _ = v40
    }
    fmt.Println(v1, v5, v5)
    v5 += 2147483647
    if v12 {
        var v41 float32 = v1
        var v42 uint = 255
        _ = v41
        _ = v42
    } else {
        if ((1000 == -32768) || false) {
            _ = len(v3)
            if true {
                for fi43 := 0; fi43 < 14; fi43++ {
                    v4 -= v4
                    v44 := GenS5{F0: v4, F1: 127, F2: -1}
                    _ = fi43
                    _ = v44
                }
            } else {
                for fi45 := 0; fi45 < 1; fi45++ {
                    var v46 int = (fi45 - 42)
                    fmt.Println(v4, fi45, fi45)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi45
                    _ = v46
                }
                fmt.Println(v12)
                if true {
                    v47 := GenS9{F0: 368, F1: (-(v5))}
                    var v48 uint16 = 32767
                    _ = v47
                    _ = v48
                } else {
                    v1 = 1e10
                }
            }
        }
        v5 += v5
    }
    for fi49 := 0; fi49 < 13; fi49++ {
        v12 = (0 < 187)
        fmt.Println(v12, v5)
        fmt.Println(v12, fi49, v1)
        _ = fi49
    }
    v50 := GenS2{F0: 42, F1: 53}
    var v51 int32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
