package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint8 }

func main() {
    v1 := []string{("hello" + "\n\t")}
    var v2 uint64 = 42
    v3 := []string{"\n\t", ("hello" + ("hello" + "")), "", ("" + "")}
    v4 := make(map[uint32]uint32)
    switch 602 {
    case 645:
        v5 := GenS1{F0: 2147483647}
        _ = v5
    case 9223372036854775807:
        var v6 string = ("a" + "hello")
        _ = v6
    case -9223372036854775808:
        switch 0 {
        case -519:
            v7 := GenS2{F0: 128, F1: 1}
            _ = v7
        case -9223372036854775808:
            delete(v4, 2)
        case 32768:
            switch 128 {
            case 32768:
                if (42 <= 127) {
                    switch ("a" + "hello") {
                    case "a":
                        fmt.Println(v2)
                    case "":
                        for fi8 := 0; fi8 < 14; fi8++ {
                            v9 := GenS2{F0: 2, F1: 700}
                            v2 += 1
                            switch ("hello" + "") {
                            case "\n\t":
                                if (42 >= 85) {
                                    v2 -= 669
                                    if false {
                                        for fi10 := 0; fi10 < 4; fi10++ {
                                            _ = len(v4)
                                            v11 := GenS3{F0: 345, F1: 42, F2: -936}
                                            switch "hello" {
                                            case "hello":
                                                var v12 uint32 = 2147483647
                                                _ = v12
                                            case "foo bar":
                                                switch fi8 {
                                                case 42:
                                                    var v13 int16 = -333
                                                    _ = v13
                                                }
                                            case "":
                                                switch "foo bar" {
                                                case "\n\t":
                                                    _ = len(v4)
                                                }
                                            default:
                                                if (!false) {
                                                    v3 = append(v3, "a")
                                                    v14 := GenS2{F0: 0, F1: 1}
                                                    v15 := v4[0]
                                                    _ = v14
                                                    _ = v15
                                                } else {
                                                    v2 += v2
                                                    var v16 int16 = 32767
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v17 := GenS3{F0: 2147483647, F1: 825, F2: (v16 - (v16 + v16))}
                                                    switch 32767 {
                                                    case 628:
                                                        v2 = v2
                                                        v18 := GenS3{F0: 590, F1: 9, F2: v16}
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if ((v2 / (((v2 % (((^(v2))) | 1))) | 1)) >= v2) {
                                                            var v19 bool = ((1e10 + (1e10 * 3.14159)) <= 0.1)
                                                            v19 = (true && (((fi10 % (((fi10 + fi8)) | 1)) <= (fi10 - (fi8 - fi8))) || false))
                                                            v20 := GenS4{F0: 32767, F1: 42, F2: 65535}
                                                            _ = v19
                                                            _ = v20
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v18
                                                    case -2147483648:
                                                        var v21 uint16 = 255
                                                        _ = v21
                                                    case -791:
                                                        v4[42] = 2
                                                    default:
                                                        fmt.Println(v2, fi8)
                                                    }
                                                    _ = v16
                                                    _ = v17
                                                }
                                                v22 := GenS5{F0: -1}
                                                _ = v22
                                            }
                                            _ = fi10
                                            _ = v11
                                        }
                                    } else {
                                        fmt.Println(fi8, v2)
                                        v2 -= 2
                                    }
                                } else {
                                    var v23 int = fi8
                                    _ = v23
                                }
                            case "":
                                v24 := GenS2{F0: 128, F1: 42}
                                _ = v24
                            case "hello":
                                for fi25 := 0; fi25 < 14; fi25++ {
                                    if false {
                                        v26 := GenS3{F0: 517, F1: 745, F2: 0}
                                        _ = v26
                                    }
                                    _ = fi25
                                }
                            }
                            _ = fi8
                            _ = v9
                        }
                    }
                    var v27 float64 = 0.1
                    v1 = append(v1, "foo bar")
                    _ = v27
                }
            case 2147483647:
                fmt.Println(v2)
            default:
                fmt.Println(v2)
            }
        }
    }
    var v28 uint = 0
    if true {
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := GenS6{F0: 284, F1: ("foo bar" + "\n\t")}
            _ = fi29
            _ = v30
        }
        v2 += 18446744073709551615
        var v31 bool = true
        _ = v31
    }
    switch "hello" {
    case "hello":
        v32 := GenS5{F0: 990}
        _ = v32
    }
    v2 = v2
    for fi33 := 0; fi33 < 3; fi33++ {
        v34 := GenS7{F0: -9223372036854775808, F1: -128}
        v2 = (^(v2))
        _ = fi33
        _ = v34
    }
    v28 = v28
    for fi35 := 0; fi35 < 3; fi35++ {
        if ((!(255 < 950)) && (2 < 2)) {
            var v36 uint64 = (^(v2))
            v37 := GenS8{F0: 2, F1: uint16(v36), F2: 1}
            _ = v36
            _ = v37
        } else {
            var v38 int8 = -1
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                fmt.Println(fi35, v2, v28)
            case "a":
                v38 -= v38
            default:
                v38 += (-(v38))
            }
            _ = v38
        }
        var v39 uint64 = (^(v2))
        v39 += (v39 & 1)
        _ = fi35
        _ = v39
    }
    switch ("\n\t" + "") {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    case "":
        var v40 uint8 = 211
        _ = v40
    case "\n\t":
        var v41 string = "a"
        _ = v41
    }
    if (true || ((-110 > 127) && true)) {
        for fi42 := 0; fi42 < 7; fi42++ {
            var v43 string = (("" + "foo bar") + "a")
            _ = fi42
            _ = v43
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
