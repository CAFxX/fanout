package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    var v1 uint = 621
    var v2 uint32 = 631
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    v5 := []string{""}
    var v6 int64 = -370
    for fi7 := 0; fi7 < 10; fi7++ {
        if true {
            v8 := GenS1{F0: -1, F1: -206}
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 uint8 = 127
                fmt.Println(fi7)
                v6 = v6
                _ = fi9
                _ = v10
            }
            var v11 int64 = (^(v6))
            _ = v11
            _ = v8
        }
        if false {
            delete(v3, "a")
            var v12 uint16 = 803
            _ = v12
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                if false {
                    if true {
                        var v14 int64 = (v6 * v6)
                        v14 = v6
                        v2 -= (v2 * (v2 << ((v1) & 63)))
                        _ = v14
                    } else {
                        v2 = (v2 | v2)
                        var v15 int16 = -1
                        _ = v15
                    }
                    v2 -= (^(v2))
                    if true {
                        if false {
                            v16 := GenS2{F0: "a"}
                            var v17 uint = 4294967295
                            _ = v16
                            _ = v17
                        } else {
                            v2 = (v2 - 407)
                        }
                        if (!false) {
                            switch ("a" + "hello") {
                            case "a":
                                _ = len(v3)
                            case "foo bar":
                                var v18 int32 = 1
                                _ = v18
                            case "hello":
                                v5 = append(v5, "\n\t")
                            default:
                                v5 = append(v5, ("foo bar" + "foo bar"))
                            }
                            var v19 uint = (v1 + v1)
                            _ = v19
                        } else {
                            v2 = (v2 + v2)
                            v20 := GenS1{F0: -1, F1: fi13}
                            switch "\n\t" {
                            case "foo bar":
                                v2 = uint32(v1)
                            default:
                                v6 += (^(v6))
                            }
                            _ = len(v4)
                            _ = v20
                        }
                        v6 += v6
                    }
                }
                v2 -= v2
                v4[(fi7 % ((fi13) | 1))] = (2 == 1)
                _ = fi13
            }
            var v21 uint16 = 32767
            v5 = append(v5, "")
            _ = v21
        }
        _ = fi7
    }
    var v22 int64 = v6
    if (-2147483648 >= 678) {
        v2 = v2
        switch ("hello" + "foo bar") {
        case "foo bar":
            switch -697 {
            case -170:
                var v23 int8 = -128
                _ = v23
            default:
                var v24 int32 = 42
                _ = v24
            }
        default:
            v2 -= (v2 >> ((0) & 63))
        }
    } else {
        _ = len(v3)
    }
    var v25 string = "a"
    if false {
        _ = len(v5)
        if false {
            fmt.Println(v25)
            if (255 != 127) {
                fmt.Println(v2)
            }
            v26 := GenS3{F0: 89, F1: true}
            _ = v26
        }
    } else {
        _ = len(v5)
        var v27 float32 = (0.1 * -2.25)
        fmt.Println(v25)
        _ = v27
    }
    v25 = v25
    fmt.Println(v22, v2)
    v4[128] = true
    switch ("" + ((v25 + (v25 + "\n\t")) + "foo bar")) {
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "hello":
        v28 := GenS1{F0: 322, F1: 42}
        _ = v28
    }
    if len(v5) > 0 {
        v5[0] = ((("a" + v25) + v25) + ((v25 + (("hello" + ((((v25 + "foo bar") + v25) + (v25 + v25)) + (v25 + (v25 + "")))) + v25)) + (v25 + v25)))
    }
    if false {
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 int = 42
            switch v25 {
            case "":
                fmt.Println(v22, v2)
            case "a":
                var v31 uint32 = (v2 >> (((^(v1))) & 63))
                _ = v31
            default:
                var v32 float64 = ((1e10 / 3.14159) / (-2.25 + 0.0))
                _ = v32
            }
            _ = fi29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
