package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int8; F2 int32 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 uint = 4294967295
    var v2 int8 = -128
    v3 := []int{-9223372036854775808, 128, -534}
    v4 := make(map[string]string)
    var v5 bool = (-9223372036854775808 != 1)
    if (32767 < 816) {
        _ = len(v4)
    } else {
        v5 = v5
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                if ((!(false && (float64((0.0 / 1.0)) >= 0.1))) || (v5 || true)) {
                    v2 -= (v2 - v2)
                }
                if (!v5) {
                    v2 = (v2 << (((v1 % ((v1) | 1))) & 63))
                    v1 -= (v1 ^ (v1 / (((v1 - (v1 % ((v1) | 1)))) | 1)))
                    v1 = 1
                }
                var v8 int16 = 1000
                _ = fi7
                _ = v8
            }
            if v5 {
                switch fi6 {
                case 1:
                    if len(v3) > 0 {
                        v3[0] = fi6
                    }
                case -849:
                    for fi9 := 0; fi9 < 8; fi9++ {
                        if v5 {
                            var v10 float64 = (-2.25 * 1.5)
                            fmt.Println(fi9, fi6, v10)
                            var v11 int64 = 42
                            _ = v10
                            _ = v11
                        }
                        var v12 int32 = 0
                        fmt.Println(v2, v1, v1)
                        _ = fi9
                        _ = v12
                    }
                }
                fmt.Println(v5)
            } else {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v3 = append(v3, -9223372036854775808)
                    _ = fi13
                }
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := GenS1{F0: v5}
                    fmt.Println(v1, v1, v5)
                    v2 += 4
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 19; fi16++ {
                    v2 += v2
                    _ = fi16
                }
            }
            _ = fi6
        }
        v5 = (!(!(0.1 != float64(1e10))))
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        switch "a" {
        case "a":
            v5 = (v5 && v5)
        case "":
            v18 := GenS2{F0: 42, F1: v2}
            _ = v18
        }
        v19 := GenS3{F0: v5, F1: 18446744073709551615}
        _ = fi17
        _ = v19
    }
    if false {
        if true {
            v20 := GenS3{F0: (0 > 449), F1: 1}
            if ("foo bar" > "\n\t") {
                fmt.Println(v1)
                if ("" != "\n\t") {
                    v21 := GenS4{F0: 2}
                    if v5 {
                        v5 = (0 <= -2147483648)
                        if (1 == 255) {
                            var v22 uint16 = 32767
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 20; fi23++ {
                            v24 := GenS5{F0: (v2 ^ (v2 << ((95) & 63))), F1: v2, F2: 623}
                            _ = fi23
                            _ = v24
                        }
                    }
                    var v25 uint8 = 255
                    _ = v21
                    _ = v25
                } else {
                    switch "hello" {
                    case "foo bar":
                        if (("a" == ((("" + "foo bar") + ("a" + (("foo bar" + ("\n\t" + "hello")) + ("hello" + "foo bar")))) + "")) || v5) {
                            fmt.Println(v2, v5, v2)
                        }
                    }
                }
                if (314 == 42) {
                    var v26 int16 = -32768
                    var v27 int = 128
                    fmt.Println(v5, v2)
                    _ = v26
                    _ = v27
                }
            } else {
                fmt.Println(v5)
                switch int(v2) {
                case 2147483647:
                    v5 = ("hello" != "a")
                case -1:
                    for fi28 := 0; fi28 < 12; fi28++ {
                        var v29 int32 = 42
                        if (v29 != -1) {
                            v2 -= (v2 << (((v1 + (^(v1)))) & 63))
                            v30 := GenS4{F0: 42}
                            _ = v30
                        } else {
                            var v31 int64 = -269
                            _ = v31
                        }
                        var v32 uint8 = 201
                        _ = fi28
                        _ = v29
                        _ = v32
                    }
                }
            }
            _ = v20
        }
        if false {
            var v33 bool = v5
            v33 = (42 <= 42)
            _ = v33
        }
        if false {
            switch "foo bar" {
            case "\n\t":
                var v34 float64 = -2.25
                _ = v34
            case "a":
                v5 = v5
            }
        } else {
            v35 := GenS1{F0: false}
            var v36 int = 2147483647
            _ = v35
            _ = v36
        }
    }
    v1 -= v1
    switch "foo bar" {
    case "\n\t":
        var v37 bool = v5
        _ = v37
    default:
        if true {
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 uint32 = 0
                _ = fi38
                _ = v39
            }
            if (255 > (v1 % (((v1 | (v1 + v1))) | 1))) {
                _ = len(v4)
                v40 := GenS5{F0: 94, F1: -1, F2: -2147483648}
                v41 := GenS2{F0: 745, F1: v2}
                _ = v40
                _ = v41
            }
            for fi42 := 0; fi42 < 9; fi42++ {
                if (v5 || v5) {
                    for fi43 := 0; fi43 < 20; fi43++ {
                        fmt.Println(v5, v2)
                        var v44 float64 = 1.5
                        _ = fi43
                        _ = v44
                    }
                } else {
                    for fi45 := 0; fi45 < 1; fi45++ {
                        if ((1000 == -1) && v5) {
                            switch "a" {
                            case "":
                                v1 -= (^(v1))
                            case "foo bar":
                                fmt.Println(v2, v5, v1)
                            }
                            fmt.Println(v1, v1)
                        } else {
                            v1 = 65535
                            v2 = (v2 % ((v2) | 1))
                            var v46 uint16 = 2
                            _ = v46
                        }
                        if (65535 != 32767) {
                            v4["a"] = "foo bar"
                        } else {
                            switch ((("" + ("\n\t" + "hello")) + "\n\t") + "foo bar") {
                            case "hello":
                                var v47 uint = 653
                                _ = v47
                            case "\n\t":
                                v5 = v5
                            }
                        }
                        v48 := GenS5{F0: 1, F1: (v2 ^ v2), F2: 723}
                        _ = fi45
                        _ = v48
                    }
                    v49 := GenS6{F0: 127, F1: fi42}
                    _ = v49
                }
                _ = fi42
            }
        }
    }
    var v50 float32 = 1e10
    fmt.Println(v2)
    v5 = v5
    switch "" {
    case "foo bar":
        v51 := GenS7{F0: 42}
        _ = v51
    case "hello":
        var v52 int64 = -1
        _ = v52
    default:
        v2 = v2
    }
    v2 -= -1
    if len(v3) > 0 {
        v3[0] = 127
    }
    fmt.Println(v5, v2, v50)
    var v53 uint64 = 1
    var v54 int32 = -48
    v53 += (^(v53))
    v53 += (v53 & v53)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v53
    _ = v54
}
