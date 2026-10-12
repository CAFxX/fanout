package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint }

func main() {
    var v1 int64 = -427
    v2 := make(map[string]bool)
    v3 := []int{}
    v4 := []string{"", (("foo bar" + "hello") + "")}
    fmt.Println(v1)
    switch "hello" {
    case "\n\t":
        v1 = 42
    case "hello":
        switch 4294967295 {
        case 32767:
            v1 -= v1
        case 1:
            v1 -= v1
        case 127:
            for fi5 := 0; fi5 < 13; fi5++ {
                var v6 int8 = 74
                _ = fi5
                _ = v6
            }
        }
    case "":
        var v7 int8 = -49
        _ = v7
    }
    var v8 int64 = v1
    switch "hello" {
    case "\n\t":
        for fi9 := 0; fi9 < 7; fi9++ {
            if true {
                var v10 int = (fi9 - (-(fi9)))
                _ = v10
            } else {
                fmt.Println(v1, v1)
                v11 := GenS1{F0: 1, F1: -991, F2: false}
                var v12 uint16 = 604
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    case "a":
        var v13 int16 = 1
        _ = v13
    case "hello":
        v1 = (v1 + v8)
    default:
        var v14 int = 2147483647
        _ = v14
    }
    switch 42 {
    case 4294967295:
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 int32 = 42
            _ = fi15
            _ = v16
        }
    case 2147483647:
        switch "" {
        case "":
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 uint16 = 42
                switch 0 {
                case -1:
                    v19 := GenS2{F0: 127, F1: 1}
                    _ = v19
                }
                if true {
                    v8 -= (v8 * (v8 + -289))
                    var v20 float64 = (0.0 + 1.5)
                    _ = v20
                } else {
                    var v21 int = fi17
                    for fi22 := 0; fi22 < 8; fi22++ {
                        switch v21 {
                        case 128:
                            var v23 int8 = 77
                            _ = v23
                        case 32768:
                            v24 := GenS3{F0: 918, F1: -1, F2: (v18 / ((v18) | 1))}
                            _ = v24
                        case 42:
                            if true {
                                v3 = append(v3, (fi17 / ((-484) | 1)))
                            } else {
                                var v25 int8 = 42
                                _ = v25
                            }
                        default:
                            if (2147483647 == 887) {
                                fmt.Println(v21, v8, fi22)
                                v26 := v2["\n\t"]
                                fmt.Println(v18, fi17)
                                _ = v26
                            }
                        }
                        fmt.Println(v8, v18)
                        _ = fi22
                    }
                    v21 += fi17
                    _ = v21
                }
                _ = fi17
                _ = v18
            }
        case "hello":
            if (0 == 42) {
                _ = len(v2)
            }
        }
    default:
        var v27 int16 = 32767
        _ = v27
    }
    v28 := GenS4{F0: -32768, F1: 42}
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 int16 = 146
        v1 = -1
        _ = fi29
        _ = v30
    }
    v31 := GenS5{F0: 18446744073709551615}
    if true {
        fmt.Println(v1, v8)
        v1 = 0
        var v32 string = ("a" + ("a" + "\n\t"))
        _ = v32
    }
    v1 = 1
    for fi33 := 0; fi33 < 19; fi33++ {
        v8 += (^(v8))
        _ = fi33
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        v1 += v8
        var v35 string = ((("a" + "foo bar") + ((("" + ("hello" + "hello")) + "\n\t") + "")) + ((("" + "\n\t") + "a") + "\n\t"))
        _ = fi34
        _ = v35
    }
    v8 += (-(v8))
    var v36 int16 = 42
    var v37 uint8 = 255
    if ("" >= "a") {
        v8 = v1
        fmt.Println(v36, v37, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v4
    _ = v8
}
