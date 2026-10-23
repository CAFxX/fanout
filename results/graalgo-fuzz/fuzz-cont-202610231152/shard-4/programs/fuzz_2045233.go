package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = -868
    v2 := []uint32{0}
    v3 := []int{}
    var v4 bool = (-1 <= 0)
    v5 := []uint32{1}
    if (((-2.25 / 1.0) >= (-2.25 / 1.0)) && (0 != -9223372036854775808)) {
        var v6 int8 = -1
        _ = v6
    } else {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 string = "hello"
            _ = fi7
            _ = v8
        }
        switch "hello" {
        case "\n\t":
            v4 = false
        case "hello":
            for fi9 := 0; fi9 < 16; fi9++ {
                fmt.Println(fi9, v1)
                _ = fi9
            }
        default:
            switch "a" {
            case "hello":
                v1 -= v1
            case "\n\t":
                fmt.Println(v1)
            case "foo bar":
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 uint16 = 42
                    _ = fi10
                    _ = v11
                }
            default:
                switch 552 {
                case 42:
                    v12 := GenS1{F0: 35, F1: -137, F2: 53}
                    _ = v12
                case -1:
                    var v13 float32 = float32(0.1)
                    _ = v13
                case 1:
                    var v14 string = ""
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 4; fi15++ {
                        for fi16 := 0; fi16 < 1; fi16++ {
                            v4 = v4
                            v17 := GenS1{F0: 53, F1: fi15, F2: -128}
                            v4 = v4
                            var v18 float64 = -2.25
                            var v19 uint8 = 2
                            _ = fi16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        switch 32767 {
                        case -577:
                            if (v4 && false) {
                                var v20 int8 = 106
                                v1 -= v1
                                _ = v20
                            }
                        }
                        var v21 string = ("\n\t" + "a")
                        _ = fi15
                        _ = v21
                    }
                }
            }
        }
        v22 := GenS2{F0: v1, F1: 930}
        _ = v22
    }
    if (525 > 4294967295) {
        v4 = v4
    } else {
        v4 = (255 < 1)
    }
    v4 = false
    if len(v2) > 0 {
        v2[0] = 14
    }
    fmt.Println(v4, v4)
    switch 128 {
    case 2147483647:
        v23 := GenS2{F0: v1, F1: 1000}
        _ = v23
    case 128:
        var v24 float32 = 0.0
        _ = v24
    case 9223372036854775807:
        var v25 int8 = 113
        _ = v25
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v1 += v1
    var v26 float32 = ((1.5 + (((((1e10 * -2.25) / 1e10) / 0.1) / float32(0.1)) + 1.5)) * 1e10)
    fmt.Println(v26)
    for fi27 := 0; fi27 < 7; fi27++ {
        fmt.Println(v1)
        _ = fi27
    }
    fmt.Println(v26, v4, v26)
    for fi28 := 0; fi28 < 3; fi28++ {
        v4 = (1 == 255)
        v29 := GenS2{F0: (^(v1)), F1: 42}
        _ = fi28
        _ = v29
    }
    switch "a" {
    case "hello":
        if ((v1 & (v1 | (v1 / ((42) | 1)))) > 0) {
            for fi30 := 0; fi30 < 5; fi30++ {
                v4 = (((!(float64(v26) > 1e10)) && true) || v4)
                _ = fi30
            }
        }
    case "\n\t":
        var v31 float32 = (v26 + v26)
        _ = v31
    case "a":
        v32 := GenS3{F0: 13}
        _ = v32
    default:
        for fi33 := 0; fi33 < 13; fi33++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
            if ((1 != 65535) || (504 <= 165)) {
                for fi34 := 0; fi34 < 16; fi34++ {
                    v26 -= v26
                    if (72 <= 127) {
                        if v4 {
                            var v35 uint = 42
                            if len(v5) > 0 {
                                v5[0] = 2147483647
                            }
                            _ = v35
                        } else {
                            var v36 uint32 = 2147483647
                            v26 = v26
                            _ = v36
                        }
                        var v37 uint16 = 19
                        _ = v37
                    }
                    fmt.Println(v26, fi33)
                    _ = fi34
                }
                v4 = (!((4294967295 >= 65535) && (2 >= 4294967295)))
                var v38 string = "\n\t"
                _ = v38
            } else {
                for fi39 := 0; fi39 < 9; fi39++ {
                    v1 = (v1 ^ v1)
                    _ = fi39
                }
                v40 := GenS1{F0: 128, F1: fi33, F2: -72}
                var v41 uint32 = 2
                _ = v40
                _ = v41
            }
            _ = fi33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
