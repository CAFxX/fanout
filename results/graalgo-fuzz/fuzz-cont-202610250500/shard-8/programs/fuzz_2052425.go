package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint32; F1 string }

func main() {
    v1 := []string{}
    v2 := []string{("a" + ""), "foo bar", ("hello" + "\n\t"), (("\n\t" + "foo bar") + "")}
    var v3 int64 = -400
    var v4 int8 = 0
    v3 += v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        if len(v1) > 0 {
            v1[0] = (((("hello" + "a") + ("hello" + "a")) + "\n\t") + "foo bar")
        }
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 float64 = (-2.25 / 0.1)
            _ = fi6
            _ = v7
        }
        var v8 float64 = 0.0
        _ = fi5
        _ = v8
    }
    fmt.Println(v4, v3)
    switch "" {
    case "hello":
        switch -9223372036854775808 {
        case 553:
            v4 += (-(v4))
        case 933:
            v4 = v4
        }
    case "foo bar":
        if (true && false) {
            v3 = v3
            v9 := GenS1{F0: v3, F1: 32767}
            _ = v9
        } else {
            v4 += 0
            v10 := GenS2{F0: v4}
            _ = v10
        }
    case "\n\t":
        switch 127 {
        case 0:
            fmt.Println(v4)
        case 313:
            switch "a" {
            case "\n\t":
                var v11 int16 = 1000
                _ = v11
            case "":
                var v12 int16 = 616
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int64 = v3
        for fi15 := 0; fi15 < 13; fi15++ {
            _ = len(v2)
            v2 = append(v2, ("a" + "a"))
            _ = fi15
        }
        if true {
            switch "hello" {
            case "a":
                for fi16 := 0; fi16 < 7; fi16++ {
                    v14 -= v14
                    v17 := GenS3{F0: (990 > 792)}
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    _ = fi16
                    _ = v17
                }
            case "":
                fmt.Println(fi13, fi13, fi13)
            case "foo bar":
                v14 -= (v14 << ((0) & 63))
            default:
                var v18 int16 = -971
                _ = v18
            }
            var v19 int16 = -147
            _ = v19
        } else {
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 uint64 = 4294967295
                v3 -= v14
                _ = fi20
                _ = v21
            }
            var v22 uint32 = 2147483647
            v3 = v3
            _ = v22
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v4)
    _ = len(v2)
    switch 714 {
    case 9223372036854775807:
        var v23 uint16 = 65535
        _ = v23
    case 32768:
        var v24 string = "hello"
        _ = v24
    case -1:
        v3 = v3
    default:
        if (42 < 255) {
            if (false || (true && false)) {
                var v25 uint16 = 684
                _ = v25
            }
            v26 := GenS4{F0: v3, F1: 330}
            var v27 string = ("foo bar" + ("\n\t" + "foo bar"))
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 8; fi28++ {
                v3 -= 0
                var v29 uint = 0
                _ = fi28
                _ = v29
            }
            if (65535 > 235) {
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 string = ""
                    _ = fi30
                    _ = v31
                }
            } else {
                if (!(0.0 >= 3.14159)) {
                    for fi32 := 0; fi32 < 14; fi32++ {
                        if (2 > 2) {
                            v4 += 42
                            v33 := GenS2{F0: (v4 << ((913) & 63))}
                            _ = v33
                        } else {
                            switch ("" + (("\n\t" + "") + ("\n\t" + "a"))) {
                            case "foo bar":
                                v1 = append(v1, "a")
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = ("" + "")
                                }
                            default:
                                for fi34 := 0; fi34 < 9; fi34++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v3, v3, v3)
                                    _ = fi34
                                }
                            }
                            v3 = (^(v3))
                            switch ((("" + "hello") + ((("hello" + ("foo bar" + "a")) + (((("" + "\n\t") + "") + "hello") + "")) + (("a" + "") + "foo bar"))) + "a") {
                            case "a":
                                v3 = (^(v3))
                            case "foo bar":
                                var v35 int = 901
                                _ = v35
                            case "hello":
                                v4 = -128
                            }
                        }
                        v3 -= 42
                        _ = fi32
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    for fi36 := 0; fi36 < 10; fi36++ {
                        if (539 > 1) {
                            var v37 float32 = (3.14159 - 0.0)
                            _ = v37
                        } else {
                            for fi38 := 0; fi38 < 1; fi38++ {
                                v4 = (v4 + v4)
                                _ = fi38
                            }
                        }
                        _ = fi36
                    }
                }
            }
            var v39 uint = 1
            _ = v39
        }
    }
    v3 -= (v3 - (v3 & 89))
    fmt.Println(v4)
    for fi40 := 0; fi40 < 16; fi40++ {
        fmt.Println(v3, v3, v3)
        _ = fi40
    }
    _ = len(v2)
    switch "\n\t" {
    case "":
        for fi41 := 0; fi41 < 9; fi41++ {
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (42 < 255) {
                    v42 := GenS5{F0: 2, F1: ("\n\t" + ("foo bar" + "\n\t"))}
                    var v43 uint = 886
                    _ = v42
                    _ = v43
                }
            } else {
                if (((828 <= 4294967295) && true) || (65535 > 1)) {
                    fmt.Println(fi41, v3, v3)
                }
            }
            _ = fi41
        }
    case "\n\t":
        var v44 int64 = (v3 - 257)
        _ = v44
    case "hello":
        v4 += v4
    default:
        v45 := GenS2{F0: (v4 / (((v4 << ((0) & 63))) | 1))}
        _ = v45
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        v2 = append(v2, "\n\t")
        _ = fi46
    }
    var v47 int64 = v3
    var v48 int16 = -815
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
}
