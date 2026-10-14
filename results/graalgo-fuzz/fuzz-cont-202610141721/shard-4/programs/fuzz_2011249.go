package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint16; F1 int8; F2 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 0.1
    v3 := []uint32{42, 970, 42}
    var v4 int8 = 57
    if true {
        v2 = 1.5
        for fi5 := 0; fi5 < 3; fi5++ {
            fmt.Println(v4, fi5)
            _ = fi5
        }
    } else {
        v6 := GenS1{F0: 2, F1: ("hello" + "hello"), F2: 42}
        _ = v6
    }
    var v7 int64 = -278
    var v8 string = "\n\t"
    v9 := GenS1{F0: 42, F1: v8, F2: 252}
    switch v8 {
    case "hello":
        var v10 int = 4294967295
        _ = v10
    case "a":
        if ((false && false) && false) {
            v1[128] = 196
            fmt.Println(v8)
            v7 += v7
        } else {
            v2 = v2
        }
    }
    fmt.Println(v7)
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 uint32 = 654
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        switch "a" {
        case "":
            var v14 int16 = 42
            _ = v14
        case "\n\t":
            v3 = append(v3, 386)
        case "hello":
            if (63 < 1) {
                var v15 int8 = -1
                for fi16 := 0; fi16 < 1; fi16++ {
                    v8 = ("hello" + v8)
                    if (v7 > 42) {
                        fmt.Println(v7, v15)
                    }
                    v2 += float32((1.5 / (1.5 / float64(float32(float64((((v2 - 1.5) - (float32(float64(v2)) / v2)) * v2)))))))
                    _ = fi16
                }
                for fi17 := 0; fi17 < 18; fi17++ {
                    switch fi13 {
                    case 4294967295:
                        v7 -= (v7 & (v7 & -332))
                    case 2147483647:
                        fmt.Println(v8, v8, v7)
                    case 32767:
                        var v18 uint64 = 4294967295
                        _ = v18
                    }
                    switch "a" {
                    case "":
                        var v19 float32 = v2
                        _ = v19
                    }
                    v20 := GenS1{F0: 620, F1: v8, F2: 2}
                    v21 := GenS1{F0: 42, F1: "a", F2: 18446744073709551615}
                    var v22 int = 4294967295
                    _ = fi17
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = v15
            }
        default:
            fmt.Println(v8, v7)
        }
        fmt.Println(v7, fi13)
        _ = fi13
    }
    v23 := GenS2{F0: (-(v4))}
    if false {
        switch v8 {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            if (42 < 4294967295) {
                v24 := GenS1{F0: 51, F1: ("a" + ""), F2: 1}
                switch (v8 + (v8 + (((v8 + (v8 + v8)) + v8) + "hello"))) {
                case "":
                    switch "a" {
                    case "a":
                        v25 := GenS2{F0: (v4 / ((v4) | 1))}
                        _ = v25
                    case "hello":
                        if (!true) {
                            var v26 bool = (171 > -250)
                            _ = v26
                        } else {
                            v2 += -2.25
                            fmt.Println(v4, v4)
                            fmt.Println(v2, v4)
                        }
                    case "foo bar":
                        v2 = v2
                    }
                case "\n\t":
                    v27 := GenS1{F0: 18446744073709551615, F1: v8, F2: 382}
                    _ = v27
                case "foo bar":
                    fmt.Println(v7, v7)
                default:
                    if (((((v8 + v8) + ((v8 + ((((v8 + v8) + v8) + v8) + v8)) + v8)) + "a") + (("foo bar" + v8) + (v8 + ""))) >= (v8 + (((v8 + "\n\t") + (("\n\t" + v8) + "a")) + ("" + (v8 + "hello"))))) {
                        if (v7 < v7) {
                            if (!true) {
                                for fi28 := 0; fi28 < 1; fi28++ {
                                    var v29 int8 = 1
                                    v30 := GenS2{F0: v29}
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            }
                        } else {
                            for fi31 := 0; fi31 < 16; fi31++ {
                                var v32 int64 = v7
                                fmt.Println(v2, v8)
                                if ((v32 << ((65535) & 63)) < -1) {
                                    for fi33 := 0; fi33 < 8; fi33++ {
                                        _ = len(v3)
                                        var v34 int = (^(fi33))
                                        _ = fi33
                                        _ = v34
                                    }
                                    v35 := GenS3{F0: 15}
                                    _ = v35
                                } else {
                                    fmt.Println(v7)
                                    for fi36 := 0; fi36 < 6; fi36++ {
                                        var v37 int32 = 0
                                        var v38 uint32 = 42
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                }
                                _ = fi31
                                _ = v32
                            }
                            for fi39 := 0; fi39 < 4; fi39++ {
                                v40 := GenS4{F0: 739}
                                if true {
                                    for fi41 := 0; fi41 < 2; fi41++ {
                                        var v42 bool = (v7 > 42)
                                        _ = fi41
                                        _ = v42
                                    }
                                    var v43 uint16 = 0
                                    v4 = (v4 >> ((256) & 63))
                                    _ = v43
                                }
                                v44 := GenS1{F0: 158, F1: v8, F2: 4294967295}
                                _ = fi39
                                _ = v40
                                _ = v44
                            }
                        }
                        fmt.Println(v2, v8, v8)
                    }
                }
                _ = v24
            }
        }
        v45 := GenS2{F0: v4}
        _ = v45
    } else {
        fmt.Println(v2, v2)
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v4, v7, v8)
        case "\n\t":
            switch v8 {
            case "hello":
                for fi46 := 0; fi46 < 4; fi46++ {
                    v2 = (v2 * 1.5)
                    _ = fi46
                }
            default:
                switch "\n\t" {
                case "":
                    v7 += (v7 / (((v7 - (v7 - v7))) | 1))
                case "\n\t":
                    v4 = v4
                case "a":
                    v7 = v7
                default:
                    v8 = v8
                }
            }
        case "":
            v47 := GenS1{F0: 18446744073709551615, F1: v8, F2: 1}
            _ = v47
        default:
            var v48 string = v8
            _ = v48
        }
    }
    v49 := GenS5{F0: v7, F1: 772}
    for fi50 := 0; fi50 < 11; fi50++ {
        switch fi50 {
        case -2147483648:
            v2 = v2
        default:
            for fi51 := 0; fi51 < 4; fi51++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi51
            }
        }
        var v52 float64 = -2.25
        var v53 int16 = 0
        _ = fi50
        _ = v52
        _ = v53
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v54 uint16 = 255
    v55 := GenS6{F0: v54, F1: (v4 & (v4 & v4)), F2: 127}
    var v56 int8 = (v4 - (v4 - v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v49
    _ = v54
    _ = v55
    _ = v56
    _ = v7
    _ = v8
    _ = v9
}
