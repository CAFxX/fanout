package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int32; F2 string }
type GenS3 struct { F0 int64; F1 uint64; F2 bool }
type GenS4 struct { F0 bool; F1 string; F2 uint8 }

func main() {
    var v1 int8 = 42
    v2 := []string{"a", "hello", "foo bar", "foo bar"}
    var v3 uint16 = 2
    var v4 int64 = -1
    v1 = (-(v1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 1; fi6++ {
        v3 = (v3 & v3)
        var v7 int16 = 0
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 2
    for fi9 := 0; fi9 < 2; fi9++ {
        if (("" + ("a" + "hello")) <= (((("\n\t" + (((("hello" + "hello") + "\n\t") + "hello") + "a")) + "a") + ("hello" + "\n\t")) + "hello")) {
            var v10 uint8 = 2
            if true {
                fmt.Println(v1)
                for fi11 := 0; fi11 < 6; fi11++ {
                    switch -9223372036854775808 {
                    case 127:
                        v12 := GenS1{F0: uint32(v8)}
                        _ = v12
                    case 128:
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    default:
                        var v13 bool = false
                        _ = v13
                    }
                    _ = fi11
                }
                for fi14 := 0; fi14 < 4; fi14++ {
                    if (((1e10 + ((1.5 * 3.14159) + 1e10)) * 0.0) != float64((0.0 + -2.25))) {
                        if true {
                            v15 := GenS1{F0: 4294967295}
                            _ = v15
                        }
                        var v16 int64 = v4
                        switch (fi14 % ((-522) | 1)) {
                        case 32767:
                            if false {
                                v17 := GenS2{F0: v4, F1: -1, F2: ""}
                                _ = v17
                            } else {
                                v2 = append(v2, "hello")
                            }
                        case 42:
                            var v18 float32 = 1e10
                            _ = v18
                        }
                        _ = v16
                    }
                    var v19 float32 = 3.14159
                    _ = fi14
                    _ = v19
                }
            } else {
                for fi20 := 0; fi20 < 10; fi20++ {
                    v1 = int8(v10)
                    var v21 uint = 4294967295
                    var v22 int16 = 1
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 uint32 = 2
                switch "foo bar" {
                case "":
                    var v24 float32 = -2.25
                    _ = v24
                case "a":
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + (("a" + "hello") + (("\n\t" + "\n\t") + "foo bar")))
                    }
                default:
                    var v25 bool = false
                    _ = v25
                }
                _ = v23
            }
            v10 += v10
            _ = v10
        } else {
            fmt.Println(fi9, v3)
            for fi26 := 0; fi26 < 13; fi26++ {
                switch ((("a" + "hello") + "hello") + "") {
                case "foo bar":
                    _ = len(v2)
                default:
                    var v27 bool = ((((v4 - v4) <= (v4 % ((-9223372036854775808) | 1))) && false) && ((3.14159 < 0.1) || false))
                    _ = v27
                }
                if false {
                    switch (("hello" + "") + (("" + ("foo bar" + "a")) + "foo bar")) {
                    case "\n\t":
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                    case "":
                        var v28 int = fi9
                        _ = v28
                    }
                }
                _ = fi26
            }
        }
        _ = fi9
    }
    v29 := GenS2{F0: (v4 << ((2) & 63)), F1: -1, F2: ("a" + "a")}
    fmt.Println(v8, v4)
    var v30 float32 = (1e10 - 3.14159)
    var v31 int = -544
    var v32 uint64 = (v8 / ((v8) | 1))
    v33 := GenS3{F0: (-(v4)), F1: v8, F2: false}
    if len(v2) > 0 {
        v2[0] = (("hello" + ("foo bar" + "\n\t")) + "a")
    }
    var v34 uint32 = 472
    v35 := GenS4{F0: true, F1: ("foo bar" + (("" + "\n\t") + "a")), F2: 1}
    if false {
        if (float32(1e10) >= float32((-2.25 / (3.14159 / float64(-2.25))))) {
            v31 = (^(v31))
        }
        v36 := GenS2{F0: 683, F1: 874, F2: "hello"}
        v4 += (-(v4))
        _ = v36
    } else {
        for fi37 := 0; fi37 < 2; fi37++ {
            switch "\n\t" {
            case "":
                v34 += v34
            }
            _ = fi37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v8
}
