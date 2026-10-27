package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint64; F2 int }
type GenS3 struct { F0 uint16; F1 int; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 int16 }

func main() {
    v1 := []int{127, 128, 128}
    v2 := make(map[int]int)
    var v3 uint8 = 1
    v4 := []uint32{2147483648, 2147483648, 65535, 0}
    var v5 uint16 = 32767
    v6 := GenS1{F0: 2147483647}
    v3 = uint8(v5)
    v3 = v3
    v3 += 127
    var v7 uint8 = 42
    var v8 float32 = -2.25
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(fi9)
        v7 += (v3 - v7)
        _ = fi9
    }
    var v10 int64 = 679
    var v11 float32 = (float32((1.5 * ((1e10 * 1.5) * ((((1.5 / float64(v8)) + (3.14159 + (1.5 * float64(v8)))) * 0.0) - 1.5)))) + (((0.0 + 1e10) + (1.5 / v8)) / v8))
    fmt.Println(v8, v5, v7)
    switch ("" + ("a" + "")) {
    case "a":
        v11 = 1.5
    case "foo bar":
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 string = "foo bar"
                switch (^(fi12)) {
                case 9223372036854775807:
                    if true {
                        var v15 uint64 = 4294967295
                        _ = v15
                    } else {
                        v11 = 3.14159
                        switch ("hello" + "") {
                        case "":
                            v10 = (v10 * 232)
                        case "a":
                            v14 = v14
                        case "foo bar":
                            switch fi12 {
                            case -756:
                                var v16 string = (v14 + v14)
                                _ = v16
                            case 722:
                                fmt.Println(v8, v7, fi12)
                            }
                        default:
                            _ = len(v2)
                        }
                        var v17 bool = false
                        _ = v17
                    }
                case 32767:
                    var v18 int32 = 2147483647
                    _ = v18
                case 1:
                    switch fi12 {
                    case 4294967295:
                        v11 -= 1e10
                    case 42:
                        if ((v8 != 0.1) || false) {
                            v19 := GenS2{F0: 65535, F1: 2, F2: (fi12 << ((18446744073709551615) & 63))}
                            v20 := GenS3{F0: (v5 >> ((98) & 63)), F1: fi12, F2: v7}
                            _ = v19
                            _ = v20
                        } else {
                            var v21 uint32 = 65535
                            v10 = (-(v10))
                            _ = v21
                        }
                        switch (v14 + ((v14 + v14) + v14)) {
                        case "a":
                            var v22 int64 = (^(v10))
                            _ = v22
                        case "foo bar":
                            var v23 int8 = -16
                            _ = v23
                        case "":
                            v24 := GenS4{F0: 338, F1: (v5 ^ v5), F2: 0}
                            _ = v24
                        }
                        v4 = append(v4, 2)
                        v25 := GenS3{F0: (v5 * (v5 >> ((0) & 63))), F1: fi13, F2: v7}
                        if false {
                            var v26 uint16 = 2
                            _ = v26
                        } else {
                            v2[fi12] = (fi12 % (((fi13 - fi13)) | 1))
                            var v27 int32 = -602
                            var v28 bool = true
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    case -912:
                        if len(v1) > 0 {
                            v1[0] = -731
                        }
                    default:
                        var v29 uint = uint(v5)
                        _ = v29
                    }
                default:
                    var v30 int = fi12
                    _ = v30
                }
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    default:
        var v31 float32 = (v8 + 3.14159)
        _ = v31
    }
    fmt.Println(v7, v5)
    switch "hello" {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        v7 = (v3 % (((^(v3))) | 1))
    }
    v8 = (0.0 - (v8 * ((1e10 - (v11 * -2.25)) - v8)))
    for fi32 := 0; fi32 < 7; fi32++ {
        fmt.Println(v7, v3, v5)
        v8 = float32(float64(v8))
        _ = fi32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        for fi34 := 0; fi34 < 14; fi34++ {
            switch fi34 {
            case 555:
                fmt.Println(v10, v5, v10)
            }
            delete(v2, (fi34 % ((42) | 1)))
            _ = fi34
        }
        switch fi33 {
        case 4294967295:
            for fi35 := 0; fi35 < 1; fi35++ {
                v36 := GenS5{F0: (v10 % ((-286) | 1))}
                if len(v4) > 0 {
                    v4[0] = 2
                }
                _ = fi35
                _ = v36
            }
        case -269:
            v37 := GenS6{F0: fi33, F1: -32768}
            _ = v37
        case 669:
            v4 = append(v4, 42)
        default:
            fmt.Println(v3)
        }
        _ = fi33
    }
    v2[130] = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
