package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int8; F2 bool }
type GenS5 struct { F0 int32; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 int32 }

func main() {
    var v1 int16 = -1
    var v2 uint32 = 125
    var v3 float64 = 3.14159
    v4 := []string{((((("foo bar" + "\n\t") + ((("" + "a") + "\n\t") + ("foo bar" + ("hello" + "")))) + "") + "hello") + "\n\t"), "a"}
    var v5 float32 = 0.1
    if true {
        fmt.Println(v2)
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        var v6 int32 = 42
        _ = v6
    }
    v5 -= (((float32(v3) + -2.25) * (v5 - (float32(0.1) / 1e10))) * v5)
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint16 = 60
        var v9 int64 = -9223372036854775808
        v5 = (1e10 - v5)
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 = 0.0
    if (100 != -55) {
        v3 = (v3 / 3.14159)
    } else {
        var v10 int8 = 1
        v11 := GenS1{F0: 255, F1: 0}
        var v12 int16 = (v1 % ((1) | 1))
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 bool = false
    if v13 {
        var v14 uint = 4294967295
        _ = v14
    }
    if true {
        v1 -= -1
        v15 := GenS1{F0: 1, F1: 42}
        _ = v15
    }
    v2 = 306
    if false {
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := GenS2{F0: (v2 % ((v2) | 1)), F1: ""}
                if ("a" < "foo bar") {
                    v19 := GenS3{F0: "\n\t"}
                    _ = v19
                } else {
                    switch "foo bar" {
                    case "a":
                        v20 := GenS2{F0: 961, F1: "a"}
                        _ = v20
                    case "":
                        if (572 < 777) {
                            if false {
                                fmt.Println(v1, v2, v13)
                                fmt.Println(v5)
                                fmt.Println(v2, v1)
                            }
                            if (!(v13 || (float64(-2.25) != float64(1e10)))) {
                                fmt.Println(v5)
                                v21 := GenS1{F0: 1, F1: 0}
                                if len(v4) > 0 {
                                    v4[0] = "foo bar"
                                }
                                _ = v21
                            } else {
                                v22 := GenS3{F0: ("hello" + "hello")}
                                v23 := GenS4{F0: true, F1: 127, F2: v13}
                                _ = v22
                                _ = v23
                            }
                            var v24 int64 = -1
                            _ = v24
                        }
                    default:
                        if (18446744073709551615 <= 0) {
                            var v25 int64 = 211
                            _ = v25
                        }
                    }
                }
                if (18446744073709551615 != 84) {
                    var v26 int16 = (v1 / (((-(v1))) | 1))
                    if len(v4) > 0 {
                        v4[0] = ("a" + "hello")
                    }
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                    v27 := GenS5{F0: 1, F1: 2, F2: 59}
                    _ = v26
                    _ = v27
                }
                _ = fi17
                _ = v18
            }
            v28 := GenS6{F0: 4294967295}
            var v29 string = ""
            _ = fi16
            _ = v28
            _ = v29
        }
        var v30 bool = (255 < 42)
        v31 := GenS7{F0: true, F1: -1}
        _ = v30
        _ = v31
    } else {
        var v32 int = 42
        v33 := GenS3{F0: (((("" + "foo bar") + ("a" + "hello")) + "") + ("\n\t" + (("hello" + "a") + "")))}
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
