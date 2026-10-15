package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 string; F2 uint32 }
type GenS8 struct { F0 int64 }

func main() {
    var v1 float64 = (1.5 + (1e10 * 0.1))
    v2 := []int{1, 4294967295, 0}
    v3 := make(map[int]bool)
    v4 := []uint32{1, 1, 592}
    var v5 int = 1
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 = (fi6 ^ -985)
        v7 := GenS1{F0: ((^(v5)) <= -127)}
        v1 -= v1
        _ = fi6
        _ = v7
    }
    v5 += -9223372036854775808
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    var v8 uint32 = 2147483647
    v1 = 0.1
    for fi9 := 0; fi9 < 12; fi9++ {
        if false {
            if (!(v8 == (v8 << ((18446744073709551615) & 63)))) {
                if ((("hello" + "foo bar") + "\n\t") != "a") {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        fmt.Println(v1)
                        switch 479 {
                        case -9223372036854775808:
                            var v11 uint8 = 190
                            _ = v11
                        }
                        fmt.Println(fi10, v1, fi10)
                        _ = fi10
                    }
                    v12 := GenS1{F0: false}
                    _ = v12
                } else {
                    if false {
                        v13 := GenS2{F0: uint16(v8)}
                        for fi14 := 0; fi14 < 6; fi14++ {
                            if true {
                                v15 := GenS3{F0: 65535, F1: 255, F2: 128}
                                _ = v15
                            } else {
                                v8 = 1
                                v16 := GenS4{F0: 1, F1: (-835 > 306)}
                                if true {
                                    switch 128 {
                                    case 4294967295:
                                        if false {
                                            var v17 uint8 = 12
                                            v5 = (-(v5))
                                            v8 += 2
                                            _ = v17
                                        }
                                    }
                                    var v18 uint64 = 244
                                    _ = v18
                                }
                                if len(v2) > 0 {
                                    v2[0] = fi9
                                }
                                var v19 uint16 = 2
                                var v20 uint64 = 18446744073709551615
                                v21 := GenS4{F0: (^(v8)), F1: true}
                                _ = v16
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            var v22 float32 = 1e10
                            v23 := GenS4{F0: 2, F1: false}
                            _ = fi14
                            _ = v22
                            _ = v23
                        }
                        var v24 int64 = 0
                        _ = v13
                        _ = v24
                    } else {
                        v25 := GenS5{F0: (fi9 + v5)}
                        v8 += (v8 % ((903) | 1))
                        _ = v25
                    }
                    switch -9223372036854775808 {
                    case -500:
                        _ = len(v4)
                    case 2147483647:
                        var v26 bool = (42 <= -1)
                        _ = v26
                    case 770:
                        switch fi9 {
                        case -2147483648:
                            for fi27 := 0; fi27 < 19; fi27++ {
                                var v28 uint32 = v8
                                var v29 int16 = 32767
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        case 921:
                            for fi30 := 0; fi30 < 7; fi30++ {
                                v5 += 4294967295
                                _ = fi30
                            }
                        }
                    default:
                        var v31 int = fi9
                        _ = v31
                    }
                }
            } else {
                for fi32 := 0; fi32 < 19; fi32++ {
                    var v33 uint32 = v8
                    _ = len(v2)
                    v2 = append(v2, fi9)
                    _ = fi32
                    _ = v33
                }
                v1 = float64(0.1)
            }
            var v34 uint8 = 2
            _ = v34
        }
        v5 = v5
        if (!true) {
            for fi35 := 0; fi35 < 17; fi35++ {
                _ = len(v2)
                fmt.Println(fi35, fi35, fi35)
                var v36 uint32 = (^(v8))
                _ = fi35
                _ = v36
            }
            if false {
                v5 += v5
                var v37 uint64 = 18446744073709551615
                _ = v37
            } else {
                _ = len(v3)
                for fi38 := 0; fi38 < 2; fi38++ {
                    v1 += 3.14159
                    switch "\n\t" {
                    case "hello":
                        switch "" {
                        case "foo bar":
                            if (true || (v1 > v1)) {
                                switch fi9 {
                                case 801:
                                    v1 = v1
                                case 4294967295:
                                    var v39 uint8 = 1
                                    _ = v39
                                case 0:
                                    var v40 uint16 = 421
                                    _ = v40
                                }
                                var v41 uint8 = 255
                                var v42 int16 = 42
                                _ = v41
                                _ = v42
                            }
                        case "\n\t":
                            if ((1.5 + ((3.14159 / 1.5) - 3.14159)) > ((float32((1e10 + (float64(float32(float64(3.14159))) - float64(float32(-2.25))))) * 3.14159) * 3.14159)) {
                                var v43 uint = 0
                                v44 := GenS4{F0: (^(v8)), F1: ((0.0 - 1e10) > (1e10 + 0.1))}
                                if (!(65535 <= (v43 | v43))) {
                                    v45 := GenS6{F0: 42}
                                    _ = v45
                                } else {
                                    v46 := GenS7{F0: (fi9 | fi9), F1: "\n\t", F2: v8}
                                    _ = v46
                                }
                                _ = v43
                                _ = v44
                            }
                        }
                    default:
                        var v47 bool = false
                        switch ("a" + "\n\t") {
                        case "":
                            var v48 float32 = (1e10 - 1.5)
                            _ = v48
                        default:
                            fmt.Println(v1)
                        }
                        fmt.Println(fi9, v8, fi9)
                        _ = v47
                    }
                    fmt.Println(v1)
                    _ = fi38
                }
                v8 = v8
            }
        }
        _ = fi9
    }
    v49 := GenS8{F0: 9223372036854775807}
    if false {
        var v50 float32 = -2.25
        v2 = append(v2, v5)
        _ = v50
    } else {
        v2 = append(v2, 128)
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        v8 = v8
        var v52 uint8 = 0
        _ = fi51
        _ = v52
    }
    fmt.Println(v5, v1, v8)
    var v53 float32 = (1.5 - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v53
    _ = v8
}
