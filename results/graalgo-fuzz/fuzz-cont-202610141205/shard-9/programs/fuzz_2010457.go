package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 int32; F2 string }
func genFn0(p0 int16) int {
    p0 += (^(p0))
    v5 := GenS1{F0: -1, F1: 42, F2: 1}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int8 = 0
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    v1 := []int{1, -2147483648, 128, 2147483647}
    var v2 float32 = ((3.14159 * 1.5) * 0.0)
    var v3 uint = 1
    v4 := []uint32{}
    v8 := GenS1{F0: 100, F1: 42, F2: 1}
    v3 = v3
    if len(v4) > 0 {
        v4[0] = 42
    }
    v4 = append(v4, 4294967295)
    for fi9 := 0; fi9 < 9; fi9++ {
        if true {
            for fi10 := 0; fi10 < 3; fi10++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if true {
                    switch "a" {
                    case "a":
                        var v11 int8 = -49
                        _ = v11
                    case "\n\t":
                        var v12 uint16 = 255
                        _ = v12
                    default:
                        _ = len(v4)
                    }
                    fmt.Println(fi9, v2)
                    v13 := GenS2{F0: 2}
                    _ = v13
                } else {
                    fmt.Println(v3)
                    v3 = (v3 >> (((v3 + (^(v3)))) & 63))
                }
                switch (fi9 | fi10) {
                case 806:
                    switch "a" {
                    case "hello":
                        if true {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch genFn0(1000) {
                                case -2147483648:
                                    v2 = (v2 + v2)
                                    v1 = append(v1, 1)
                                case 42:
                                    v14 := GenS3{F0: 4294967295}
                                    _ = v14
                                case 885:
                                    v3 = v3
                                default:
                                    var v15 uint16 = 615
                                    _ = v15
                                }
                                if (!(159 <= 874)) {
                                    if ((17 <= 127) || false) {
                                        v3 += 4294967295
                                        v2 -= (v2 / v2)
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    fmt.Println(v3, v3, fi9)
                                }
                            } else {
                                if (false || (false || false)) {
                                    switch fi9 {
                                    case 32768:
                                        v3 = 255
                                        v3 = v3
                                    case 2147483647:
                                        if (-723 < 0) {
                                            v3 -= (v3 - v3)
                                            v16 := GenS4{F0: 899, F1: v3, F2: 927}
                                            var v17 uint8 = 25
                                            _ = v16
                                            _ = v17
                                        }
                                    default:
                                        switch fi10 {
                                        case -560:
                                            var v18 int16 = 32767
                                            _ = v18
                                        case 32767:
                                            if (((("\n\t" + "foo bar") + "\n\t") + (("foo bar" + "") + "foo bar")) != "hello") {
                                                v19 := GenS5{F0: 1, F1: 255}
                                                _ = v19
                                            }
                                            fmt.Println(fi10, v3, v2)
                                            var v20 int32 = 42
                                            fmt.Println(v3)
                                            v3 -= 207
                                            v21 := GenS2{F0: 249}
                                            _ = v20
                                            _ = v21
                                        case 127:
                                            v2 -= (v2 / float32(1.5))
                                        }
                                    }
                                } else {
                                    v22 := GenS5{F0: 127, F1: 255}
                                    _ = v22
                                }
                                v2 = v2
                                fmt.Println(fi10, fi9)
                            }
                            v1 = append(v1, (fi10 << ((2) & 63)))
                        } else {
                            v3 -= v3
                            v2 -= v2
                        }
                    case "foo bar":
                        _ = len(v4)
                    case "a":
                        if (!((((^(v3)) >= (v3 - 1)) && true) || true)) {
                            var v23 int8 = -1
                            _ = v23
                        }
                    }
                case 9223372036854775807:
                    v2 += v2
                default:
                    v2 += v2
                }
                _ = fi10
            }
            v3 = (v3 - (^(v3)))
        } else {
            v3 = 1
        }
        var v24 uint64 = 42
        _ = fi9
        _ = v24
    }
    v1 = append(v1, genFn0(32767))
    _ = len(v1)
    switch "foo bar" {
    case "hello":
        for fi25 := 0; fi25 < 10; fi25++ {
            v1 = append(v1, (fi25 * fi25))
            v26 := GenS6{F0: fi25}
            var v27 float64 = 1e10
            _ = fi25
            _ = v26
            _ = v27
        }
    case "":
        if true {
            v3 -= (v3 >> ((v3) & 63))
        }
    case "a":
        v2 += v2
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        v3 += 255
        fmt.Println(v2, v2)
        _ = fi28
    }
    switch ("hello" + "hello") {
    case "":
        v3 = 529
    }
    switch "a" {
    case "hello":
        var v29 uint64 = 18446744073709551615
        _ = v29
    case "foo bar":
        v2 = ((float32(0.0) - v2) / ((v2 - (v2 + -2.25)) - v2))
    case "a":
        var v30 uint8 = 252
        _ = v30
    }
    var v31 int32 = 1
    switch 32768 {
    case 128:
        fmt.Println(v2, v31, v3)
    case 621:
        if false {
            _ = len(v4)
            if (124 <= -121) {
                v2 = v2
                var v32 int16 = 1
                _ = v32
            }
            switch 329 {
            case 0:
                fmt.Println(v2, v31, v31)
            default:
                for fi33 := 0; fi33 < 3; fi33++ {
                    switch fi33 {
                    case 9223372036854775807:
                        var v34 float32 = v2
                        _ = v34
                    default:
                        var v35 int32 = 548
                        _ = v35
                    }
                    _ = fi33
                }
            }
        } else {
            v3 += 4294967295
            if false {
                v36 := GenS7{F0: v31, F1: 1, F2: (("hello" + "foo bar") + "a")}
                var v37 uint = v3
                _ = v36
                _ = v37
            }
        }
    }
    v2 -= 1.5
    var v38 uint64 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v8
}
