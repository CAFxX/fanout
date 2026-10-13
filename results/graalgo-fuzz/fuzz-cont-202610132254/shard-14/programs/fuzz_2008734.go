package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 int16; F1 string; F2 int }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string; F1 int32; F2 string }

func main() {
    var v1 int8 = 42
    v2 := []uint32{1}
    var v3 int64 = 9223372036854775807
    var v4 uint8 = 1
    var v5 uint32 = 1
    v6 := make(map[string]string)
    switch 1 {
    case 2147483647:
        if (-1 < 128) {
            v3 = v3
            if true {
                var v7 int64 = 1
                fmt.Println(v3)
                _ = v7
            }
            if (127 <= -598) {
                switch "a" {
                case "":
                    if (1 != -1) {
                        var v8 uint = 907
                        switch ((("" + "") + "foo bar") + (((("hello" + "") + "") + "a") + (("a" + ("\n\t" + ("\n\t" + "a"))) + "\n\t"))) {
                        case "foo bar":
                            v8 = 18446744073709551615
                        case "\n\t":
                            for fi9 := 0; fi9 < 20; fi9++ {
                                var v10 uint8 = v4
                                if false {
                                    v11 := GenS1{F0: (v8 >> (((v8 % ((255) | 1))) & 63))}
                                    var v12 float64 = 0.1
                                    v4 = 115
                                    _ = v11
                                    _ = v12
                                }
                                if false {
                                    var v13 uint = v8
                                    _ = v13
                                } else {
                                    var v14 uint16 = 1
                                    _ = v14
                                }
                                _ = fi9
                                _ = v10
                            }
                        case "a":
                            if true {
                                _ = len(v2)
                            }
                        default:
                            v2 = append(v2, (v5 ^ (^(v5))))
                        }
                        _ = v8
                    } else {
                        var v15 uint32 = (v5 + (^(v5)))
                        v2 = append(v2, (^(v5)))
                        _ = v15
                    }
                case "foo bar":
                    switch 0 {
                    case 352:
                        for fi16 := 0; fi16 < 17; fi16++ {
                            var v17 uint = 548
                            _ = fi16
                            _ = v17
                        }
                    }
                case "a":
                    if true {
                        var v18 int = 901
                        _ = v18
                    }
                default:
                    var v19 bool = true
                    _ = v19
                }
                var v20 string = "foo bar"
                _ = v20
            }
        } else {
            var v21 int8 = -58
            _ = v21
        }
    case 32768:
        if (!true) {
            switch int(v4) {
            case 9223372036854775807:
                v4 += (^(v4))
            default:
                if false {
                    if true {
                        switch 9223372036854775807 {
                        case 90:
                            v22 := GenS1{F0: 4294967295}
                            _ = v22
                        }
                        if false {
                            v23 := GenS1{F0: 18446744073709551615}
                            _ = v23
                        } else {
                            v5 -= 2147483647
                            fmt.Println(v5, v4, v1)
                        }
                    }
                }
            }
            var v24 int16 = 42
            v25 := GenS2{F0: (v3 % ((v3) | 1)), F1: 2, F2: ("a" + "a")}
            _ = v24
            _ = v25
        }
    case 103:
        v1 -= v1
    }
    v3 += -854
    var v26 uint64 = 2
    v27 := GenS2{F0: -9223372036854775808, F1: (v5 + (v5 * v5)), F2: "a"}
    for fi28 := 0; fi28 < 6; fi28++ {
        _ = len(v6)
        var v29 uint = 588
        _ = fi28
        _ = v29
    }
    var v30 float64 = float64((float32(-2.25) - 0.0))
    v5 += v5
    if len(v2) > 0 {
        v2[0] = v5
    }
    var v31 string = "hello"
    if ((v3 >> ((65535) & 63)) > v3) {
        fmt.Println(v30)
        var v32 uint = 546
        delete(v6, v31)
        _ = v32
    }
    v4 = 127
    if (-2147483648 > -1) {
        switch "a" {
        case "hello":
            var v33 float64 = -2.25
            _ = v33
        case "a":
            fmt.Println(v1)
        case "":
            var v34 uint64 = (v26 * v26)
            _ = v34
        default:
            for fi35 := 0; fi35 < 11; fi35++ {
                var v36 bool = (32767 > 196)
                v37 := GenS3{F0: 42, F1: v31, F2: (fi35 >> ((65535) & 63))}
                for fi38 := 0; fi38 < 12; fi38++ {
                    fmt.Println(fi35, v4)
                    v39 := GenS4{F0: v3, F1: (^(v26))}
                    _ = fi38
                    _ = v39
                }
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        if (395 != v26) {
            v40 := v6["hello"]
            var v41 float32 = ((float32(float64(0.1)) * -2.25) / ((((0.1 / 1.5) + (((float32(((v30 - (float64(1e10) / v30)) / v30)) * 1e10) * 1.5) * -2.25)) - (0.0 / 1.0)) + 0.0))
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 8; fi42++ {
            v43 := GenS1{F0: 1}
            v1 += v1
            v44 := GenS5{F0: v5, F1: -12, F2: (v4 ^ v4)}
            _ = fi42
            _ = v43
            _ = v44
        }
    } else {
        var v45 int64 = (v3 >> ((574) & 63))
        for fi46 := 0; fi46 < 2; fi46++ {
            var v47 int64 = v45
            if (true || ((32767 > -838) && ((uint16(v4) == 2) && true))) {
                _ = len(v2)
                v45 -= (v47 ^ v45)
            } else {
                var v48 uint64 = 18446744073709551615
                v49 := GenS6{F0: 0}
                _ = v48
                _ = v49
            }
            _ = fi46
            _ = v47
        }
        v3 -= (v3 | (v45 + v3))
        _ = v45
    }
    for fi50 := 0; fi50 < 1; fi50++ {
        switch fi50 {
        case -728:
            for fi51 := 0; fi51 < 8; fi51++ {
                v4 = v4
                var v52 float32 = -2.25
                var v53 uint32 = (v5 / ((v5) | 1))
                _ = fi51
                _ = v52
                _ = v53
            }
        default:
            if true {
                for fi54 := 0; fi54 < 20; fi54++ {
                    fmt.Println(v1, fi54)
                    if false {
                        if true {
                            fmt.Println(v3, fi54)
                            v6[(v31 + v31)] = v31
                        } else {
                            v55 := GenS5{F0: v5, F1: 1000, F2: (^(v4))}
                            v56 := GenS7{F0: false, F1: 2147483647}
                            _ = v55
                            _ = v56
                        }
                    }
                    _ = fi54
                }
                var v57 string = (v31 + v31)
                _ = v57
            }
        }
        v58 := GenS8{F0: "foo bar", F1: -2147483648, F2: "\n\t"}
        _ = fi50
        _ = v58
    }
    var v59 uint16 = 13
    switch "a" {
    case "a":
        var v60 float64 = v30
        _ = v60
    case "":
        v61 := v6[v31]
        _ = v61
    default:
        var v62 float32 = 0.0
        _ = v62
    }
    var v63 int = 1
    v31 = "hello"
    switch (^(v63)) {
    case 120:
        v64 := GenS3{F0: -32768, F1: ("\n\t" + ((v31 + "a") + (((v31 + v31) + v31) + ""))), F2: -1}
        _ = v64
    case -1:
        fmt.Println(v5, v3, v1)
    default:
        v31 = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v63
}
