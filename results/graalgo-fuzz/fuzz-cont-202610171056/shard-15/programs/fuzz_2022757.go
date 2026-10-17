package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }
func genFn0(p0 uint16) int8 {
    v5 := []int{32768}
    v6 := GenS1{F0: 1000, F1: 2}
    _ = p0
    _ = v5
    _ = v6
    return -24
}
func genFn1() uint32 {
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 6; fi8++ {
        v7 = v7
        if true {
            v7 = v7
        }
        var v9 int8 = 72
        _ = fi8
        _ = v9
    }
    v7 = (v7 + v7)
    var v10 float64 = 0.1
    _ = v10
    _ = v7
    return 1
}

func main() {
    var v1 bool = true
    v2 := []string{"", (("\n\t" + ("\n\t" + "\n\t")) + "hello"), (("\n\t" + "\n\t") + ("\n\t" + "a"))}
    var v3 bool = (132 > 1)
    v4 := make(map[string]bool)
    if (42 > 231) {
        var v11 string = ("" + (("foo bar" + ("\n\t" + "foo bar")) + ("" + "foo bar")))
        _ = v11
    } else {
        v12 := v4["foo bar"]
        _ = v12
    }
    fmt.Println(v3, v1)
    if ((2 >= 843) && false) {
        if (0 == -1) {
            v1 = (true && v1)
            delete(v4, ("hello" + ""))
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v13 := v4[((("a" + (("a" + ((("foo bar" + "hello") + "hello") + ("" + ("\n\t" + "\n\t")))) + "hello")) + "") + "foo bar")]
        _ = len(v4)
        _ = v13
    }
    v14 := GenS1{F0: 1000, F1: 0}
    if v1 {
        var v15 float32 = 3.14159
        _ = v15
    }
    v16 := GenS2{F0: 1, F1: 0, F2: -860}
    fmt.Println(v3, v1)
    switch ((("a" + "") + "foo bar") + "") {
    case "hello":
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int8 = 100
            _ = fi17
            _ = v18
        }
    case "":
        fmt.Println(v1)
    case "a":
        var v19 int8 = 67
        _ = v19
    default:
        fmt.Println(v3, v1)
    }
    v3 = v3
    if v1 {
        v20 := GenS2{F0: 1, F1: 482, F2: -32768}
        _ = v20
    } else {
        v21 := GenS3{F0: 353, F1: 1, F2: 42}
        for fi22 := 0; fi22 < 17; fi22++ {
            v1 = true
            var v23 int64 = int64(fi22)
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    _ = len(v4)
    var v24 int32 = -855
    switch -2147483648 {
    case 42:
        switch 128 {
        case 127:
            if false {
                switch 913 {
                case -9223372036854775808:
                    var v25 int8 = genFn0(267)
                    _ = v25
                case 734:
                    if (-1 == 127) {
                        switch "foo bar" {
                        case "foo bar":
                            if (18446744073709551615 > 18446744073709551615) {
                                v24 -= v24
                                v24 = -451
                            }
                        case "\n\t":
                            var v26 uint64 = 0
                            _ = v26
                        case "a":
                            fmt.Println(v24, v24, v24)
                        default:
                            for fi27 := 0; fi27 < 12; fi27++ {
                                if true {
                                    if (fi27 > -655) {
                                        fmt.Println(v24, v3, v3)
                                        var v28 uint = 4294967295
                                        _ = v28
                                    }
                                    v29 := v4[("foo bar" + "hello")]
                                    for fi30 := 0; fi30 < 13; fi30++ {
                                        fmt.Println(v3, v1)
                                        switch (fi30 ^ fi30) {
                                        case 32768:
                                            var v31 string = ("\n\t" + "")
                                            _ = v31
                                        case -902:
                                            var v32 float64 = 0.1
                                            _ = v32
                                        case -9223372036854775808:
                                            switch "\n\t" {
                                            case "":
                                                var v33 int8 = -128
                                                _ = v33
                                            default:
                                                if false {
                                                    v34 := GenS1{F0: 0, F1: 0}
                                                    var v35 bool = (((((1e10 + 1.5) - 0.1) + float64(1e10)) / (1.5 + ((((3.14159 + (3.14159 + (-2.25 / 1e10))) * (float64((1e10 / 1.5)) - (-2.25 + float64(3.14159)))) + 0.0) + 0.0))) < (3.14159 - 0.1))
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    v29 = (1 < 100)
                                                }
                                            }
                                            var v36 int = (fi27 | -9223372036854775808)
                                            _ = v36
                                        }
                                        _ = fi30
                                    }
                                    _ = v29
                                }
                                _ = fi27
                            }
                        }
                        if v3 {
                            _ = len(v4)
                            var v37 uint = 2
                            _ = len(v4)
                            _ = v37
                        }
                        v1 = v1
                    }
                case 42:
                    fmt.Println(v3, v24, v3)
                }
            } else {
                v38 := GenS3{F0: 0, F1: 0, F2: 0}
                _ = v38
            }
        }
    case 2147483647:
        v1 = v3
    case 1:
        v24 -= v24
    }
    var v39 int32 = (v24 + v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
}
