package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int; F2 uint32 }
type GenS8 struct { F0 uint64; F1 uint16; F2 int32 }

func main() {
    v1 := []string{"", "\n\t", (("hello" + ("foo bar" + "foo bar")) + "foo bar"), ("\n\t" + "")}
    var v2 int16 = 657
    var v3 uint = 1
    var v4 float64 = (-2.25 / (float64(((0.0 * 1.5) + 3.14159)) - -2.25))
    var v5 int64 = -1
    v6 := GenS1{F0: 2, F1: 1, F2: 1}
    var v7 uint32 = 403
    v8 := GenS1{F0: 771, F1: (v5 * -9223372036854775808), F2: 100}
    for fi9 := 0; fi9 < 2; fi9++ {
        switch "" {
        case "a":
            switch (fi9 >> (((v3 % ((v3) | 1))) & 63)) {
            case 1:
                switch "hello" {
                case "\n\t":
                    if ((v3 & v3) >= v3) {
                        v10 := GenS1{F0: 7, F1: (-(v5)), F2: 42}
                        v2 += v2
                        fmt.Println(v4, v4)
                        _ = v10
                    }
                default:
                    fmt.Println(v2)
                }
            case -9223372036854775808:
                if true {
                    fmt.Println(fi9, v3)
                    _ = len(v1)
                    fmt.Println(fi9, v2, v2)
                } else {
                    v3 = v3
                }
            default:
                for fi11 := 0; fi11 < 14; fi11++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if ((fi9 & fi9) > 0) {
                        v5 += (v5 << ((v3) & 63))
                    } else {
                        switch fi9 {
                        case 2147483647:
                            var v12 uint16 = 65535
                            _ = v12
                        case 42:
                            v1 = append(v1, ("foo bar" + "a"))
                        default:
                            _ = len(v1)
                        }
                        if false {
                            v13 := GenS2{F0: 616, F1: 1}
                            fmt.Println(v3)
                            _ = v13
                        }
                        v14 := GenS3{F0: fi9, F1: (v5 & (v5 & (v5 + v5)))}
                        _ = v14
                    }
                    v7 = (v7 ^ v7)
                    _ = fi11
                }
            }
        case "":
            if ("hello" > "a") {
                switch "a" {
                case "\n\t":
                    for fi15 := 0; fi15 < 19; fi15++ {
                        v7 = (v7 & (v7 % ((v7) | 1)))
                        var v16 bool = true
                        fmt.Println(v4, v7, v2)
                        _ = fi15
                        _ = v16
                    }
                case "foo bar":
                    if true {
                        _ = len(v1)
                        _ = len(v1)
                    }
                }
                v4 += (((v4 * (v4 / v4)) - (float64(1e10) + 0.1)) * ((3.14159 * v4) / v4))
            } else {
                switch (-(fi9)) {
                case 4294967295:
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := GenS4{F0: true, F1: false}
                        v4 += (v4 * 0.0)
                        if false {
                            var v19 bool = (42 != 42)
                            if (18446744073709551615 != 42) {
                                fmt.Println(v5)
                                if false {
                                    v19 = true
                                    fmt.Println(v3)
                                    v20 := GenS3{F0: (fi17 & (fi17 - 4294967295)), F1: v5}
                                    _ = v20
                                } else {
                                    var v21 int32 = int32(v3)
                                    _ = v21
                                }
                                var v22 float32 = 0.0
                                _ = v22
                            } else {
                                fmt.Println(v3, v19, v3)
                            }
                            _ = v19
                        } else {
                            var v23 int16 = 1000
                            if (!false) {
                                _ = len(v1)
                            } else {
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        if len(v1) > 0 {
                                            v1[0] = "a"
                                        }
                                    case "foo bar":
                                        v4 -= v4
                                    }
                                }
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    var v24 string = "hello"
                                    _ = v24
                                default:
                                    v23 -= 32767
                                }
                            }
                            _ = v23
                        }
                        _ = fi17
                        _ = v18
                    }
                case 42:
                    if (true && true) {
                        v7 += (^(v7))
                    } else {
                        v1 = append(v1, ("\n\t" + "foo bar"))
                        if ((-9223372036854775808 != v5) && false) {
                            var v25 uint = 4294967295
                            _ = v25
                        } else {
                            _ = len(v1)
                        }
                    }
                default:
                    var v26 float32 = 0.1
                    _ = v26
                }
                for fi27 := 0; fi27 < 2; fi27++ {
                    v3 = v3
                    var v28 float64 = v4
                    _ = fi27
                    _ = v28
                }
                v3 = (v3 << ((42) & 63))
            }
        case "hello":
            v5 = 9223372036854775807
        }
        var v29 int8 = 100
        _ = fi9
        _ = v29
    }
    fmt.Println(v7, v7, v4)
    fmt.Println(v7)
    v30 := GenS5{F0: -210, F1: (v7 << ((v3) & 63))}
    switch "foo bar" {
    case "":
        if true {
            var v31 uint64 = uint64(v7)
            fmt.Println(v2, v4)
            _ = v31
        } else {
            v7 = v7
            if false {
                fmt.Println(v2, v4, v4)
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            } else {
                v7 -= 4294967295
            }
            var v32 uint64 = 556
            _ = v32
        }
    case "hello":
        _ = len(v1)
    default:
        switch -832 {
        case -443:
            v33 := GenS6{F0: -233}
            _ = v33
        case 32768:
            if ((^(v7)) >= (v7 * v7)) {
                v34 := GenS7{F0: 42, F1: 32767, F2: v7}
                if false {
                    fmt.Println(v7, v5, v4)
                    var v35 uint64 = 99
                    _ = v35
                }
                var v36 bool = ((v7 / (((v7 / (((^(v7))) | 1))) | 1)) == (^(v7)))
                _ = v34
                _ = v36
            }
        case 436:
            fmt.Println(v4, v3, v7)
        default:
            switch 42 {
            case 61:
                if (!false) {
                    v1 = append(v1, "foo bar")
                } else {
                    v5 = v5
                }
            }
        }
    }
    v37 := GenS8{F0: 0, F1: 65535, F2: 2147483647}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
