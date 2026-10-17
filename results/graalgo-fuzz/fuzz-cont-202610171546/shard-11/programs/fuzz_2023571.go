package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 bool; F2 int8 }
type GenS5 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 198
    var v3 int16 = 1
    v4 := []int{}
    var v5 float64 = (-2.25 / (0.0 + (-2.25 + 1.5)))
    if true {
        if true {
            var v6 uint16 = 32767
            _ = v6
        } else {
            var v7 float32 = ((0.0 * 1.5) * 3.14159)
            _ = v7
        }
        fmt.Println(v2)
    } else {
        v5 -= 0.1
        var v8 uint64 = 18446744073709551615
        for fi9 := 0; fi9 < 2; fi9++ {
            switch (fi9 * 32767) {
            case -2147483648:
                if true {
                    if false {
                        var v10 float64 = (v5 - v5)
                        if false {
                            if (((v10 >= 3.14159) || true) && true) {
                                v11 := GenS1{F0: true}
                                v12 := GenS1{F0: (false || (fi9 == fi9))}
                                _ = v11
                                _ = v12
                            }
                            if (!true) {
                                var v13 uint64 = (v8 * (v8 / ((2) | 1)))
                                if false {
                                    var v14 int16 = v3
                                    if true {
                                        if false {
                                            if (!true) {
                                                var v15 int16 = (v3 % ((1) | 1))
                                                v1[4294967295] = fi9
                                                _ = v15
                                            }
                                            var v16 float64 = 3.14159
                                            _ = v16
                                        }
                                        var v17 string = ""
                                        _ = v17
                                    }
                                    _ = v14
                                }
                                var v18 uint64 = (v8 * 613)
                                _ = v13
                                _ = v18
                            } else {
                                switch (fi9 - int(v3)) {
                                case -2147483648:
                                    switch "" {
                                    case "foo bar":
                                        v19 := GenS1{F0: ((!((42 >= 128) && (!true))) && true)}
                                        _ = v19
                                    case "hello":
                                        fmt.Println(v3)
                                    case "a":
                                        _ = len(v4)
                                    default:
                                        var v20 float64 = v5
                                        _ = v20
                                    }
                                case 32767:
                                    fmt.Println(fi9, v8)
                                default:
                                    fmt.Println(v8, v10, fi9)
                                }
                                var v21 bool = ((!(127 >= 127)) || (1e10 >= 1.5))
                                _ = v21
                            }
                        }
                        _ = v10
                    } else {
                        switch (fi9 - (fi9 ^ (fi9 / (((-(fi9))) | 1)))) {
                        case -54:
                            if false {
                                v2 = (-(v2))
                                if (112 >= 2) {
                                    v8 = (v8 - (v8 & (v8 % ((2) | 1))))
                                    var v22 uint8 = 255
                                    _ = v22
                                } else {
                                    v23 := GenS1{F0: (!(!(42 <= 255)))}
                                    _ = v23
                                }
                            } else {
                                v24 := GenS2{F0: v3, F1: (^(fi9))}
                                v5 = v5
                                _ = v24
                            }
                        case -9223372036854775808:
                            switch ("" + "") {
                            case "":
                                var v25 bool = (false || false)
                                _ = v25
                            case "\n\t":
                                fmt.Println(v8)
                            case "a":
                                v5 -= 3.14159
                            }
                        default:
                            v26 := GenS2{F0: (v3 % (((v3 + v3)) | 1)), F1: -1}
                            _ = v26
                        }
                        v27 := GenS3{F0: 168}
                        _ = v27
                    }
                } else {
                    v4 = append(v4, fi9)
                }
            }
            fmt.Println(v3, v5, v3)
            _ = fi9
        }
        _ = v8
    }
    var v28 float64 = (0.0 - v5)
    v29 := v1[505]
    v5 = 3.14159
    for fi30 := 0; fi30 < 13; fi30++ {
        v28 += float64(((-2.25 + (0.0 / 1.0)) / 3.14159))
        _ = fi30
    }
    v28 += float64(0.0)
    var v31 uint8 = 255
    var v32 int = (v29 + v29)
    if false {
        var v33 int = v32
        v34 := GenS4{F0: 1, F1: false, F2: 0}
        var v35 uint16 = 42
        _ = v33
        _ = v34
        _ = v35
    } else {
        var v36 float32 = 3.14159
        _ = v36
    }
    v37 := GenS5{F0: v29, F1: v2}
    v38 := v1[2147483648]
    var v39 int16 = 1000
    if true {
        v40 := GenS3{F0: v31}
        v4 = append(v4, (v32 << ((613) & 63)))
        _ = v40
    }
    v3 += v39
    var v41 int16 = v3
    v29 = (v29 % (((v32 ^ -203)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
}
