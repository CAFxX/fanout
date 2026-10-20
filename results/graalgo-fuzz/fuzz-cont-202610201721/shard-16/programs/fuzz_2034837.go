package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := []string{"", "", "foo bar"}
    var v2 uint64 = 1
    var v3 uint32 = 711
    v4 := []string{(((("a" + ("" + "hello")) + "foo bar") + "") + "\n\t"), (("\n\t" + "a") + "hello"), "hello", ("hello" + "foo bar")}
    v3 += (v3 ^ (v3 ^ 96))
    v1 = append(v1, "")
    var v5 int64 = -1
    var v6 int64 = v5
    fmt.Println(v3, v6)
    switch 1 {
    case 127:
        fmt.Println(v2, v6)
    case 9223372036854775807:
        var v7 int32 = -1
        _ = v7
    default:
        v5 = (v5 >> ((0) & 63))
    }
    v2 = v2
    var v8 float32 = 1e10
    for fi9 := 0; fi9 < 9; fi9++ {
        if (!true) {
            var v10 int16 = -32768
            _ = v10
        } else {
            v11 := GenS1{F0: 253, F1: "\n\t", F2: -9223372036854775808}
            if (v3 <= v3) {
                switch "a" {
                case "a":
                    if (0 >= 1) {
                        var v12 uint8 = 1
                        _ = v12
                    }
                }
            }
            _ = v11
        }
        _ = fi9
    }
    v13 := GenS2{F0: 128}
    switch "\n\t" {
    case "a":
        if (983 <= 255) {
            switch -1 {
            case -9223372036854775808:
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
            }
            fmt.Println(v2, v8)
            v14 := GenS3{F0: false}
            _ = v14
        } else {
            var v15 int32 = 42
            v6 = (^(v6))
            _ = v15
        }
    }
    fmt.Println(v6, v2, v6)
    v8 = (v8 * float32((-2.25 * ((0.0 - (((-2.25 * float64(v8)) * 1.5) / (0.1 + 1.5))) * -2.25))))
    v16 := GenS1{F0: 119, F1: ("a" + ("\n\t" + "")), F2: (v5 % ((v6) | 1))}
    switch "" {
    case "hello":
        fmt.Println(v5)
    case "\n\t":
        v8 += v8
    default:
        if ((1.5 / (0.1 + 1.5)) != (3.14159 / (1.5 + float64(-2.25)))) {
            v17 := GenS4{F0: 255}
            switch 2147483647 {
            case 902:
                if true {
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                    if (!(!(v2 == (v2 * v2)))) {
                        var v18 int64 = v5
                        var v19 uint32 = (v3 + (v3 + 42))
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 uint32 = (v3 ^ v19)
                            if false {
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    if (!true) {
                                        var v23 string = "\n\t"
                                        var v24 float32 = (0.1 - (v8 - 3.14159))
                                        _ = v23
                                        _ = v24
                                    } else {
                                        var v25 float32 = float32(1.5)
                                        switch "a" {
                                        case "hello":
                                            var v26 float64 = 1e10
                                            _ = v26
                                        case "\n\t":
                                            v18 = int64(v3)
                                        case "":
                                            var v27 uint16 = 255
                                            _ = v27
                                        }
                                        _ = v25
                                    }
                                    _ = fi22
                                }
                                if false {
                                    v21 -= 4294967295
                                    v21 = v21
                                }
                            }
                            var v28 float64 = 1e10
                            _ = fi20
                            _ = v21
                            _ = v28
                        }
                        _ = v18
                        _ = v19
                    } else {
                        v29 := GenS5{F0: (v6 - -929), F1: -64}
                        for fi30 := 0; fi30 < 11; fi30++ {
                            var v31 int8 = 127
                            _ = fi30
                            _ = v31
                        }
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                        _ = v29
                    }
                } else {
                    var v32 int64 = (-(v5))
                    var v33 uint64 = v2
                    _ = v32
                    _ = v33
                }
            default:
                v34 := GenS6{F0: 255}
                _ = v34
            }
            _ = v17
        }
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        v2 -= v2
        var v36 int = fi35
        var v37 bool = false
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
