package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 1
    var v3 uint64 = 574
    var v4 uint8 = v2
    var v5 int = 127
    if ((false || false) || false) {
        v5 += v5
        switch "foo bar" {
        case "hello":
            switch "" {
            case "foo bar":
                v5 = v5
            case "":
                v1["foo bar"] = (1 == -353)
            default:
                if false {
                    if false {
                        fmt.Println(v4, v5)
                        _ = len(v1)
                    }
                    for fi6 := 0; fi6 < 8; fi6++ {
                        var v7 float64 = -2.25
                        var v8 float64 = v7
                        if false {
                            for fi9 := 0; fi9 < 5; fi9++ {
                                switch (-(fi6)) {
                                case 128:
                                    if (!true) {
                                        v10 := GenS1{F0: v5, F1: "hello"}
                                        fmt.Println(v5, v2)
                                        if ((-2.25 * 1e10) < -2.25) {
                                            v2 -= (v4 << ((65535) & 63))
                                            if true {
                                                _ = len(v1)
                                            }
                                            switch (fi6 | -2147483648) {
                                            case 2147483647:
                                                var v11 int64 = 9223372036854775807
                                                _ = v11
                                            case 127:
                                                switch -9223372036854775808 {
                                                case 42:
                                                    v12 := GenS2{F0: true, F1: 127}
                                                    _ = v12
                                                case 32768:
                                                    v4 -= v4
                                                default:
                                                    fmt.Println(v3)
                                                }
                                            case 32767:
                                                var v13 bool = (false || (0 > 53))
                                                _ = v13
                                            }
                                        }
                                        _ = v10
                                    }
                                    v14 := GenS2{F0: (v2 > (v4 << ((987) & 63))), F1: 0}
                                    _ = v14
                                case -1:
                                    var v15 uint64 = 259
                                    _ = v15
                                }
                                var v16 int = fi6
                                v17 := v1[("foo bar" + ((("" + "foo bar") + (("\n\t" + ("foo bar" + ("hello" + "\n\t"))) + "\n\t")) + "a"))]
                                _ = fi9
                                _ = v16
                                _ = v17
                            }
                            if true {
                                v3 -= v3
                                var v18 uint = 65535
                                switch (((("foo bar" + "hello") + (((("a" + ("foo bar" + (("" + "foo bar") + "hello"))) + "\n\t") + "a") + (("hello" + "\n\t") + ""))) + "") + "") {
                                case "a":
                                    var v19 uint8 = 1
                                    _ = v19
                                case "hello":
                                    v20 := GenS2{F0: (1 > -848), F1: -107}
                                    _ = v20
                                case "":
                                    for fi21 := 0; fi21 < 6; fi21++ {
                                        v22 := v1[("hello" + "a")]
                                        _ = fi21
                                        _ = v22
                                    }
                                default:
                                    for fi23 := 0; fi23 < 7; fi23++ {
                                        var v24 int64 = 1
                                        _ = fi23
                                        _ = v24
                                    }
                                }
                                _ = v18
                            }
                            v25 := GenS1{F0: v5, F1: ""}
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 13; fi26++ {
                                v2 = v4
                                fmt.Println(v2, v3)
                                _ = fi26
                            }
                            v2 -= uint8(v3)
                        }
                        _ = fi6
                        _ = v7
                        _ = v8
                    }
                } else {
                    v27 := []int{(-(v5)), -374, (v5 * (v5 << ((65535) & 63)))}
                    switch ("hello" + "hello") {
                    case "":
                        var v28 uint64 = (v3 * (v3 % ((2) | 1)))
                        _ = v28
                    case "\n\t":
                        v29 := GenS1{F0: 2147483647, F1: ("" + "foo bar")}
                        _ = v29
                    }
                    v30 := GenS3{F0: 42}
                    _ = v27
                    _ = v30
                }
            }
        }
        var v31 int64 = 1
        _ = v31
    }
    v5 -= 352
    v32 := GenS4{F0: 0, F1: (("\n\t" + "hello") < ("\n\t" + "\n\t"))}
    var v33 float64 = (1.5 + float64(0.1))
    var v34 float64 = 3.14159
    v5 = v5
    var v35 string = ""
    _ = len(v1)
    for fi36 := 0; fi36 < 12; fi36++ {
        v37 := v1[((v35 + (("hello" + v35) + v35)) + v35)]
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
