package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 float64 = 3.14159
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: true}
            switch "foo bar" {
            case "\n\t":
                v7 = v7
            case "":
                v7 = (v7 / v7)
            case "a":
                v7 += float64(-2.25)
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 uint32 = 634
    if true {
        switch 42 {
        case -743:
            for fi11 := 0; fi11 < 20; fi11++ {
                v10 += v10
                _ = fi11
            }
        case 420:
            v12 := []int{}
            _ = v12
        case -2147483648:
            fmt.Println(v10)
        default:
            v10 = v10
        }
        fmt.Println(v10)
        v13 := []int{128, -736, 278}
        _ = v13
    } else {
        v14 := GenS1{F0: true}
        _ = v14
    }
    if (951 >= 42) {
        fmt.Println(v10)
        v10 -= 2147483647
        v10 = (v10 << ((602) & 63))
    }
    switch (("foo bar" + "hello") + ("hello" + ("\n\t" + "foo bar"))) {
    case "a":
        v10 += (v10 >> ((0) & 63))
    case "":
        v15 := GenS1{F0: (false || (false || true))}
        _ = v15
    }
    _ = v10
    return 39
}

func main() {
    var v1 int16 = 686
    var v2 float32 = 1.5
    var v3 int64 = 705
    v4 := []string{"", "hello", "hello", ""}
    var v5 int8 = 1
    switch (((("" + "") + "foo bar") + "") + ("" + "hello")) {
    case "hello":
        v16 := GenS2{F0: -786, F1: "foo bar", F2: "\n\t"}
        _ = v16
    case "a":
        var v17 int16 = v1
        _ = v17
    case "foo bar":
        if ((("" + (("a" + "hello") + (("\n\t" + ((((("" + "hello") + ("" + "hello")) + "foo bar") + "hello") + (("" + "a") + "hello"))) + "foo bar"))) + "foo bar") >= "") {
            if false {
                for fi18 := 0; fi18 < 15; fi18++ {
                    v19 := GenS3{F0: 100, F1: 0, F2: 2}
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 14; fi20++ {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v3 = (v3 | v3)
                        var v22 int64 = (v3 - (v3 & 42))
                        _ = fi21
                        _ = v22
                    }
                    _ = fi20
                }
                var v23 uint32 = 4294967295
                _ = v23
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                fmt.Println(fi24, v1)
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 bool = false
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
        }
    }
    v27 := GenS1{F0: false}
    switch 128 {
    case -119:
        switch "" {
        case "":
            if true {
                switch 32767 {
                case -2147483648:
                    switch (("" + "\n\t") + ("hello" + "\n\t")) {
                    case "hello":
                        _ = len(v4)
                    }
                case 1:
                    for fi28 := 0; fi28 < 15; fi28++ {
                        fmt.Println(fi28, fi28, v1)
                        _ = fi28
                    }
                default:
                    switch ("foo bar" + "") {
                    case "hello":
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                    }
                }
                v5 = -128
            }
        case "\n\t":
            if true {
                for fi29 := 0; fi29 < 2; fi29++ {
                    v5 = (v5 & v5)
                    var v30 uint = 0
                    _ = fi29
                    _ = v30
                }
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
            } else {
                for fi31 := 0; fi31 < 3; fi31++ {
                    switch ("a" + "\n\t") {
                    case "a":
                        fmt.Println(v5)
                    default:
                        switch fi31 {
                        case -1:
                            v2 += v2
                        case 80:
                            var v32 string = ("foo bar" + "a")
                            _ = v32
                        default:
                            v1 = (v1 % ((v1) | 1))
                        }
                    }
                    for fi33 := 0; fi33 < 14; fi33++ {
                        fmt.Println(v2, fi33, v2)
                        v34 := GenS3{F0: v5, F1: 80, F2: 267}
                        fmt.Println(fi33, v3, v1)
                        _ = fi33
                        _ = v34
                    }
                    if false {
                        v35 := GenS2{F0: 9223372036854775807, F1: "a", F2: ((("a" + "foo bar") + (("\n\t" + "foo bar") + "\n\t")) + "a")}
                        _ = v35
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v36 uint32 = 580
                        _ = v36
                    }
                    _ = fi31
                }
                switch 127 {
                case -1:
                    v3 = (v3 | v3)
                case 0:
                    var v37 int8 = (^(v5))
                    _ = v37
                }
            }
        case "foo bar":
            v5 += 123
        default:
            fmt.Println(v1)
        }
    default:
        var v38 int64 = (v3 % ((-9223372036854775808) | 1))
        _ = v38
    }
    var v39 string = ("" + ("\n\t" + "hello"))
    if true {
        fmt.Println(v5, v2, v3)
        v40 := GenS3{F0: (-(v5)), F1: genFn0(), F2: 255}
        var v41 float64 = (-2.25 + ((1e10 / 1.0) + 0.0))
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        fmt.Println(v5, fi42)
        v43 := GenS4{F0: 42}
        for fi44 := 0; fi44 < 19; fi44++ {
            v2 = 0.1
            _ = fi44
        }
        _ = fi42
        _ = v43
    }
    if (32767 >= (v1 | (v1 % ((v1) | 1)))) {
        v45 := GenS5{F0: v3, F1: 2147483647}
        v39 = v39
        _ = v45
    } else {
        v39 = v39
        var v46 uint = 65535
        fmt.Println(v39, v46)
        _ = v46
    }
    v47 := GenS6{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v47
    _ = v5
}
