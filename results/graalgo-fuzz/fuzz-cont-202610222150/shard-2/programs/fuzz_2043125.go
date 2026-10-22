package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 int32 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint16; F1 uint64; F2 int32 }

func main() {
    var v1 int8 = -101
    v2 := []uint32{}
    var v3 uint16 = 2
    var v4 uint32 = 1
    v5 := make(map[string]string)
    v6 := []int{-925, -1, 1, 9223372036854775807}
    var v7 float64 = (3.14159 - 0.1)
    v7 += float64(0.1)
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int8 = -112
            _ = fi9
            _ = v10
        }
        _ = len(v6)
        _ = fi8
    }
    var v11 uint32 = (v4 / (((v4 | v4)) | 1))
    if (true && (-2.25 != -2.25)) {
        fmt.Println(v7, v1)
        for fi12 := 0; fi12 < 7; fi12++ {
            if false {
                if (false || (-9223372036854775808 != 1)) {
                    if false {
                        var v13 uint64 = 4294967295
                        _ = v13
                    } else {
                        for fi14 := 0; fi14 < 18; fi14++ {
                            v11 -= v11
                            _ = fi14
                        }
                        for fi15 := 0; fi15 < 19; fi15++ {
                            v16 := GenS1{F0: 939, F1: 31}
                            fmt.Println(fi12)
                            if (false || (!false)) {
                                v7 = (float64((0.1 / 1.5)) / ((v7 / v7) + -2.25))
                            } else {
                                var v17 uint64 = 451
                                v17 = v17
                                switch "hello" {
                                case "\n\t":
                                    v4 += (v4 + 4294967295)
                                case "":
                                    v18 := GenS2{F0: 2, F1: 0, F2: (^(v17))}
                                    _ = v18
                                }
                                switch "foo bar" {
                                case "hello":
                                    var v19 uint8 = 0
                                    v20 := v5[(("foo bar" + ("" + (("\n\t" + "") + ""))) + (("hello" + "\n\t") + ""))]
                                    _ = v19
                                    _ = v20
                                default:
                                    v11 = v11
                                }
                                if true {
                                    fmt.Println(v7, v11)
                                    switch fi15 {
                                    case -1:
                                        var v21 uint64 = 42
                                        _ = v21
                                    case 678:
                                        _ = len(v6)
                                    case 4294967295:
                                        var v22 bool = true
                                        _ = v22
                                    default:
                                        fmt.Println(v3)
                                    }
                                }
                                if false {
                                    if (v7 >= (v7 - ((v7 / -2.25) + v7))) {
                                        var v23 float64 = ((v7 + v7) / v7)
                                        fmt.Println(fi15, fi12, fi15)
                                        if (true && true) {
                                            v24 := GenS3{F0: 501, F1: 504}
                                            v25 := GenS4{F0: 1000, F1: "hello", F2: int32(fi15)}
                                            if (3.14159 <= float64(1e10)) {
                                                fmt.Println(v3)
                                            } else {
                                                switch (fi12 | (fi15 >> ((42) & 63))) {
                                                case 32768:
                                                    switch (^(fi12)) {
                                                    case 46:
                                                        var v26 int32 = -2147483648
                                                        _ = v26
                                                    case 42:
                                                        if true {
                                                            var v27 int32 = 928
                                                            v4 -= (v11 / (((^(v11))) | 1))
                                                            _ = v27
                                                        } else {
                                                            v28 := GenS5{F0: false, F1: v1}
                                                            v29 := v5[(("" + ("hello" + ("foo bar" + ""))) + "\n\t")]
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                    case 2147483647:
                                                        switch ("" + "foo bar") {
                                                        case "foo bar":
                                                            if (-76 == -1) {
                                                                if false {
                                                                    v7 += (1.5 + ((v7 + v23) * v23))
                                                                    v1 = v1
                                                                } else {
                                                                    var v30 uint64 = (v17 * uint64(v3))
                                                                    var v31 uint = uint(v3)
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                                var v32 float64 = (v23 * v7)
                                                                _ = v32
                                                            }
                                                            v3 -= v3
                                                            switch (fi15 * 32768) {
                                                            case 2147483647:
                                                                v4 = (v11 ^ v4)
                                                            case 1:
                                                                fmt.Println(v17, v1, v3)
                                                            }
                                                        case "\n\t":
                                                            v3 -= v3
                                                        case "":
                                                            fmt.Println(v1, v1)
                                                        }
                                                    }
                                                }
                                                switch fi12 {
                                                case 2147483647:
                                                    v3 = v3
                                                case 127:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                            }
                                            _ = v24
                                            _ = v25
                                        } else {
                                            fmt.Println(fi12)
                                            v33 := GenS6{F0: 873}
                                            _ = v33
                                        }
                                        _ = v23
                                    }
                                }
                                _ = v17
                            }
                            v1 -= v1
                            _ = fi15
                            _ = v16
                        }
                    }
                } else {
                    var v34 string = "\n\t"
                    _ = v34
                }
                v11 += 2147483648
            }
            fmt.Println(v11, v4)
            v35 := GenS6{F0: 2147483647}
            _ = fi12
            _ = v35
        }
        v1 += 19
    } else {
        if (((2147483647 <= 1) || (!(int8(v11) < 42))) && (255 != 1)) {
            v1 = -124
        } else {
            var v36 float64 = -2.25
            var v37 uint = 255
            _ = v36
            _ = v37
        }
        switch "\n\t" {
        case "a":
            v7 = (1e10 - v7)
        case "":
            _ = len(v6)
        default:
            var v38 float32 = float32(float64((-2.25 + 0.1)))
            _ = v38
        }
    }
    if true {
        v1 = (v1 * v1)
        if true {
            v39 := GenS5{F0: true, F1: (^(v1))}
            switch 143 {
            case -1:
                switch "" {
                case "a":
                    switch "hello" {
                    case "":
                        fmt.Println(v11)
                    case "a":
                        v1 -= (^(v1))
                    }
                case "":
                    v11 -= (v4 % (((v4 % ((4294967295) | 1))) | 1))
                case "\n\t":
                    v40 := GenS7{F0: 127}
                    _ = v40
                }
            }
            _ = v39
        } else {
            var v41 int8 = v1
            var v42 uint16 = (v3 / ((v3) | 1))
            _ = v41
            _ = v42
        }
    }
    v11 += uint32(v3)
    v4 -= 276
    v43 := GenS8{F0: 2, F1: -449}
    for fi44 := 0; fi44 < 12; fi44++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v45 := GenS5{F0: false, F1: (v1 & (v1 ^ -1))}
        v46 := GenS6{F0: -2147483648}
        _ = fi44
        _ = v45
        _ = v46
    }
    switch ("hello" + (("foo bar" + "a") + "foo bar")) {
    case "foo bar":
        v47 := v5[""]
        _ = v47
    case "hello":
        if false {
            v3 = 32767
            fmt.Println(v1)
        } else {
            fmt.Println(v7)
        }
    case "a":
        var v48 int16 = 42
        _ = v48
    default:
        v49 := GenS9{F0: (v3 * uint16(v4)), F1: 2, F2: -1}
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
