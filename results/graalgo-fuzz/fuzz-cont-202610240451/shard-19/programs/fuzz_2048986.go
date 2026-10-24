package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int64; F1 string; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 int8; F1 uint8; F2 uint }
type GenS8 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{4294967295, 127, 42, 32767}
    var v3 uint32 = 172
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{(^(v3)), 42}
    if ((float64((float32(0.0) * 1e10)) < 0.0) || (-2.25 >= 0.0)) {
        switch "" {
        case "foo bar":
            v3 = 48
        case "hello":
            v3 += (v3 / ((v3) | 1))
        case "\n\t":
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 int32 = 1
                if true {
                    var v8 int = (fi6 / ((fi6) | 1))
                    var v9 uint8 = 2
                    v7 = 619
                    _ = v8
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            v5 = append(v5, 4294967295)
            var v11 uint = 0
            switch "\n\t" {
            case "hello":
                if (true || (-25 < -128)) {
                    var v12 int64 = 9223372036854775807
                    if ((!(!true)) || false) {
                        for fi13 := 0; fi13 < 18; fi13++ {
                            v14 := GenS1{F0: -1, F1: (v11 & (v11 + v11))}
                            if ((false || (!false)) && true) {
                                var v15 uint = (v11 * v11)
                                _ = v15
                            }
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        var v16 int = (fi10 >> ((v11) & 63))
                        v2 = append(v2, v16)
                        var v17 uint64 = (v4 / ((18446744073709551615) | 1))
                        _ = v16
                        _ = v17
                    }
                    _ = v12
                } else {
                    var v18 int16 = -1
                    v19 := GenS1{F0: 1, F1: v11}
                    _ = v18
                    _ = v19
                }
            case "foo bar":
                switch "hello" {
                case "a":
                    _ = len(v5)
                case "foo bar":
                    switch "hello" {
                    case "":
                        if false {
                            for fi20 := 0; fi20 < 19; fi20++ {
                                _ = len(v2)
                                var v21 int8 = 1
                                v11 -= v11
                                _ = fi20
                                _ = v21
                            }
                            if true {
                                fmt.Println(v11, v3)
                            }
                            for fi22 := 0; fi22 < 18; fi22++ {
                                v23 := GenS1{F0: 1, F1: 1}
                                v4 = v4
                                _ = fi22
                                _ = v23
                            }
                        }
                    case "hello":
                        if (3.14159 <= 1e10) {
                            var v24 float64 = ((1e10 + 1e10) + float64((0.0 + (1.5 - 0.0))))
                            var v25 bool = true
                            var v26 int32 = -878
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        v27 := GenS2{F0: uint8(v4), F1: true, F2: (^(v3))}
                        _ = v27
                    }
                case "\n\t":
                    if (v11 != 255) {
                        if false {
                            v1[""] = v3
                            fmt.Println(fi10, v11)
                        } else {
                            if false {
                                v4 = v4
                            } else {
                                var v28 float64 = 0.1
                                var v29 string = "foo bar"
                                _ = v28
                                _ = v29
                            }
                            v30 := GenS1{F0: 0, F1: v11}
                            _ = v30
                        }
                    } else {
                        var v31 int = (-(fi10))
                        v3 = 65535
                        v2 = append(v2, 852)
                        _ = v31
                    }
                }
            default:
                fmt.Println(v4, v4, v11)
            }
            _ = fi10
            _ = v11
        }
    }
    var v32 float32 = 0.0
    switch "hello" {
    case "hello":
        v4 += (^(v4))
    default:
        var v33 uint64 = (^(v4))
        _ = v33
    }
    v34 := GenS3{F0: -372, F1: 0}
    for fi35 := 0; fi35 < 15; fi35++ {
        var v36 uint64 = (v4 / ((4294967295) | 1))
        var v37 string = ""
        switch v37 {
        case "foo bar":
            fmt.Println(fi35, v37)
        case "\n\t":
            v4 -= 477
        case "a":
            v38 := GenS4{F0: 0, F1: (v37 + ("" + "foo bar")), F2: (v3 & v3)}
            _ = v38
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    v32 = (v32 * v32)
    v32 = (float32(((float64((float32(0.1) / v32)) / 0.1) * 0.0)) - v32)
    var v39 int = 397
    for fi40 := 0; fi40 < 13; fi40++ {
        if (!false) {
            if true {
                var v41 int32 = -2147483648
                fmt.Println(fi40)
                _ = v41
            } else {
                for fi42 := 0; fi42 < 8; fi42++ {
                    v43 := GenS5{F0: (18446744073709551615 < 255)}
                    switch v39 {
                    case 2147483647:
                        fmt.Println(v4)
                    }
                    delete(v1, "")
                    _ = fi42
                    _ = v43
                }
            }
        } else {
            delete(v1, "")
        }
        _ = fi40
    }
    fmt.Println(v32, v32)
    v44 := GenS6{F0: uint64(v3), F1: 42}
    if (!false) {
        v3 = (v3 / ((v3) | 1))
        v3 -= (v3 * (v3 & uint32(v4)))
    }
    for fi45 := 0; fi45 < 20; fi45++ {
        switch "a" {
        case "hello":
            fmt.Println(v3)
        case "a":
            var v46 uint8 = 242
            _ = v46
        case "foo bar":
            var v47 string = "\n\t"
            _ = v47
        default:
            v3 = v3
        }
        v48 := GenS7{F0: 42, F1: 2, F2: 255}
        _ = fi45
        _ = v48
    }
    v49 := GenS6{F0: (^(v4)), F1: -842}
    v50 := GenS8{F0: v39, F1: -32768}
    fmt.Println(v3, v3)
    v4 += (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v39
    _ = v4
    _ = v44
    _ = v49
    _ = v5
    _ = v50
}
