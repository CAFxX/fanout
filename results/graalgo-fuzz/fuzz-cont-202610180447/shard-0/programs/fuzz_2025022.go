package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int; F1 int64 }
func genFn0() uint8 {
    v5 := []int{32768, 32767}
    switch ("" + "\n\t") {
    case "\n\t":
        switch -2147483648 {
        case 183:
            for fi6 := 0; fi6 < 19; fi6++ {
                v7 := GenS1{F0: "\n\t"}
                v8 := GenS1{F0: "hello"}
                fmt.Println(fi6)
                _ = fi6
                _ = v7
                _ = v8
            }
        case 868:
            if (false || (!true)) {
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(fi9)
                    _ = fi9
                }
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 int64 = 9223372036854775807
                    _ = fi10
                    _ = v11
                }
                switch 347 {
                case 483:
                    var v12 bool = false
                    _ = v12
                default:
                    var v13 uint32 = 864
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 float32 = 1e10
                    for fi16 := 0; fi16 < 5; fi16++ {
                        v15 = ((3.14159 / (v15 - v15)) / (float32(0.0) + 1.5))
                        _ = len(v5)
                        _ = fi16
                    }
                    var v17 int16 = 32767
                    _ = fi14
                    _ = v15
                    _ = v17
                }
                var v18 float64 = float64(float32(-2.25))
                _ = v18
            }
        default:
            var v19 float64 = 0.0
            _ = v19
        }
    default:
        if true {
            switch ((("hello" + ("\n\t" + "foo bar")) + (("\n\t" + "\n\t") + "hello")) + "foo bar") {
            case "\n\t":
                v20 := GenS1{F0: ("hello" + ("a" + "foo bar"))}
                _ = v20
            case "foo bar":
                switch "a" {
                case "hello":
                    var v21 int16 = 108
                    _ = v21
                case "\n\t":
                    var v22 bool = (0 != 136)
                    _ = v22
                }
            case "a":
                var v23 float64 = (1e10 + 1.5)
                _ = v23
            default:
                for fi24 := 0; fi24 < 12; fi24++ {
                    for fi25 := 0; fi25 < 14; fi25++ {
                        var v26 int16 = -138
                        v26 -= v26
                        v26 += (v26 / (((v26 * (v26 * v26))) | 1))
                        _ = fi25
                        _ = v26
                    }
                    var v27 int8 = -128
                    _ = fi24
                    _ = v27
                }
            }
        } else {
            if (true || false) {
                var v28 uint = 65535
                if false {
                    switch "\n\t" {
                    case "hello":
                        if (0 != 1) {
                            fmt.Println(v28)
                            var v29 int32 = 42
                            switch ("hello" + "hello") {
                            case "\n\t":
                                v30 := GenS2{F0: -128, F1: 656}
                                _ = v30
                            case "hello":
                                var v31 int16 = -32768
                                _ = v31
                            default:
                                v28 += (v28 & v28)
                            }
                            _ = v29
                        }
                    case "\n\t":
                        var v32 int32 = 2147483647
                        _ = v32
                    case "a":
                        if ((v28 | (^(v28))) >= v28) {
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            var v33 int = 1
                            _ = v33
                        } else {
                            if (!true) {
                                v28 = 4294967295
                            }
                        }
                    }
                } else {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        v28 = v28
                        v28 = 1
                        switch 32767 {
                        case 0:
                            v35 := GenS3{F0: 609, F1: 2147483647}
                            _ = v35
                        default:
                            v28 = v28
                        }
                        _ = fi34
                    }
                }
                v28 += v28
                _ = v28
            }
            v5 = append(v5, 628)
        }
    }
    var v36 int8 = -128
    var v37 float64 = (-2.25 + (float64(1e10) / 1.0))
    switch 209 {
    case -590:
        _ = len(v5)
    default:
        v36 -= v36
    }
    _ = v36
    _ = v37
    _ = v5
    return 255
}

func main() {
    var v1 float64 = 0.0
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    var v4 int16 = -772
    for fi38 := 0; fi38 < 20; fi38++ {
        v1 = v1
        _ = fi38
    }
    v39 := GenS1{F0: "foo bar"}
    v2 = (((3.14159 / 1.0) / 0.1) > float32((float64(0.1) + v1)))
    var v40 int8 = 127
    var v41 int8 = 15
    v2 = true
    if v2 {
        v42 := GenS4{F0: 0, F1: -1}
        v43 := v3[2147483647]
        _ = v42
        _ = v43
    } else {
        v44 := GenS4{F0: -1, F1: -9223372036854775808}
        _ = v44
    }
    for fi45 := 0; fi45 < 4; fi45++ {
        var v46 uint8 = 1
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
