package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool }

func main() {
    var v1 int8 = 127
    var v2 uint16 = 2
    var v3 int32 = -1
    v4 := []string{}
    var v5 uint8 = 2
    var v6 bool = (-1 >= (v1 ^ v1))
    fmt.Println(v3, v3, v1)
    _ = len(v4)
    switch "" {
    case "hello":
        v2 += (v2 >> ((2) & 63))
    case "":
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(v5)
            var v8 float64 = 0.0
            _ = fi7
            _ = v8
        }
    case "a":
        var v9 float64 = -2.25
        _ = v9
    }
    v10 := GenS1{F0: 127, F1: (32767 != -212)}
    v4 = append(v4, "")
    v11 := GenS1{F0: v1, F1: ((v1 * v1) != (^(v1)))}
    fmt.Println(v1)
    if v6 {
        fmt.Println(v5, v5)
        switch 873 {
        case -457:
            for fi12 := 0; fi12 < 17; fi12++ {
                v1 = (v1 % ((1) | 1))
                switch fi12 {
                case 128:
                    v13 := GenS1{F0: (^(v1)), F1: v6}
                    _ = v13
                case 707:
                    v14 := GenS2{F0: "a"}
                    _ = v14
                case 0:
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "hello")
                    }
                }
                _ = fi12
            }
        case -102:
            v4 = append(v4, "foo bar")
        default:
            v5 = (v5 << ((255) & 63))
        }
    } else {
        var v15 float32 = float32(1.5)
        v16 := GenS3{F0: -227, F1: 0}
        var v17 int16 = 744
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS4{F0: 461, F1: 479}
    var v19 int64 = -9223372036854775808
    var v20 bool = ((v1 / (((^(v1))) | 1)) <= v1)
    if v6 {
        v21 := GenS5{F0: v19}
        _ = v21
    }
    v1 += (-(v1))
    v22 := GenS4{F0: 2, F1: uint(v5)}
    fmt.Println(v20, v2)
    for fi23 := 0; fi23 < 17; fi23++ {
        if v20 {
            switch fi23 {
            case 128:
                var v24 int8 = (v1 % (((v1 & v1)) | 1))
                _ = v24
            case -1:
                var v25 uint16 = (v2 - v2)
                _ = v25
            case 1:
                v1 = int8(v19)
            default:
                var v26 bool = true
                _ = v26
            }
            v27 := GenS3{F0: (^(fi23)), F1: 2147483647}
            switch (("a" + ("" + "a")) + "\n\t") {
            case "hello":
                switch "foo bar" {
                case "":
                    for fi28 := 0; fi28 < 20; fi28++ {
                        _ = len(v4)
                        if (0 > 1) {
                            if (!(160 <= 1)) {
                                v2 = 666
                            }
                            v19 = v19
                        } else {
                            _ = len(v4)
                        }
                        _ = fi28
                    }
                case "hello":
                    for fi29 := 0; fi29 < 5; fi29++ {
                        if len(v4) > 0 {
                            v4[0] = ("foo bar" + "a")
                        }
                        if (501 <= 95) {
                            fmt.Println(v5, v3, v19)
                            v30 := GenS6{F0: v5}
                            _ = v30
                        } else {
                            if (false || false) {
                                v4 = append(v4, "foo bar")
                                switch "hello" {
                                case "hello":
                                    switch ("" + "\n\t") {
                                    case "foo bar":
                                        var v31 uint8 = v5
                                        _ = v31
                                    case "\n\t":
                                        v4 = append(v4, "foo bar")
                                    default:
                                        v1 -= (v1 & 42)
                                    }
                                }
                                if ((false || v20) && true) {
                                    v19 -= v19
                                }
                                var v32 int16 = 1
                                _ = v32
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch (fi29 / (((fi29 | (fi29 - fi23))) | 1)) {
                        case 128:
                            if v20 {
                                v4 = append(v4, "")
                            }
                        case -9223372036854775808:
                            v33 := GenS7{F0: "a"}
                            v34 := GenS1{F0: v1, F1: v20}
                            _ = v33
                            _ = v34
                        }
                        _ = fi29
                    }
                }
            default:
                v35 := GenS3{F0: 59, F1: 889}
                _ = v35
            }
            _ = v27
        }
        v36 := GenS8{F0: true}
        _ = fi23
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
