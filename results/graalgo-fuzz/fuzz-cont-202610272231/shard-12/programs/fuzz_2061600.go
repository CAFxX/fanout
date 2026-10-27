package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint }
type GenS5 struct { F0 uint; F1 int; F2 int16 }
type GenS6 struct { F0 uint16; F1 int64 }

func main() {
    var v1 bool = (1 == 255)
    var v2 float64 = 3.14159
    var v3 int64 = 9223372036854775807
    v4 := []int{-461}
    v5 := make(map[int]string)
    fmt.Println(v2)
    var v6 float64 = (1e10 - 1e10)
    v7 := GenS1{F0: 503, F1: 4294967295, F2: 4294967295}
    var v8 uint16 = 42
    fmt.Println(v2)
    delete(v5, 61)
    v8 -= (v8 + v8)
    v4 = append(v4, -157)
    v8 = v8
    for fi9 := 0; fi9 < 13; fi9++ {
        v1 = (v1 && v1)
        if (!v1) {
            v6 -= 0.0
            var v10 float32 = -2.25
            _ = v10
        } else {
            if (2147483647 < 1) {
                v2 -= (((((v2 / v2) + 3.14159) * v2) + v6) * 0.1)
                switch (-(fi9)) {
                case 0:
                    fmt.Println(v3)
                }
                if v1 {
                    v11 := GenS2{F0: 2, F1: 2147483647, F2: 100}
                    fmt.Println(v1, fi9)
                    var v12 uint = 42
                    _ = v11
                    _ = v12
                } else {
                    if false {
                        _ = len(v4)
                    } else {
                        for fi13 := 0; fi13 < 8; fi13++ {
                            var v14 float64 = 3.14159
                            if ("" != ("hello" + "")) {
                                v1 = true
                            } else {
                                var v15 int64 = 449
                                if v1 {
                                    var v16 int8 = 42
                                    switch ("\n\t" + "hello") {
                                    case "\n\t":
                                        v17 := GenS3{F0: int64(v8), F1: v15, F2: 42}
                                        _ = v17
                                    case "hello":
                                        if false {
                                            var v18 uint8 = 128
                                            _ = v18
                                        }
                                    }
                                    if (uint8(v8) != 2) {
                                        fmt.Println(v15, v6, v2)
                                        var v19 uint8 = 0
                                        _ = len(v5)
                                        _ = v19
                                    } else {
                                        var v20 int16 = -1
                                        _ = v20
                                    }
                                    _ = v16
                                }
                                _ = v15
                            }
                            if (1 <= (fi9 % (((fi9 & (fi9 % ((32768) | 1)))) | 1))) {
                                v3 -= v3
                            }
                            if v1 {
                                v21 := GenS4{F0: 0, F1: 1, F2: 65535}
                                _ = v21
                            } else {
                                _ = len(v5)
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi13
                            _ = v14
                        }
                        switch fi9 {
                        case 783:
                            if v1 {
                                switch "\n\t" {
                                case "foo bar":
                                    for fi22 := 0; fi22 < 11; fi22++ {
                                        if true {
                                            v1 = true
                                            var v23 int64 = -9223372036854775808
                                            _ = v23
                                        }
                                        var v24 float64 = float64(float32(((v6 - (0.0 - v6)) + v6)))
                                        var v25 uint = 2
                                        fmt.Println(v24)
                                        _ = fi22
                                        _ = v24
                                        _ = v25
                                    }
                                case "":
                                    v8 = 255
                                case "\n\t":
                                    for fi26 := 0; fi26 < 6; fi26++ {
                                        switch "hello" {
                                        case "hello":
                                            v4 = append(v4, fi26)
                                        case "a":
                                            var v27 uint = 0
                                            _ = v27
                                        default:
                                            var v28 float64 = v6
                                            _ = v28
                                        }
                                        var v29 int64 = -571
                                        var v30 float64 = 3.14159
                                        _ = fi26
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                v31 := GenS2{F0: 4294967295, F1: 2, F2: 0}
                                _ = v31
                            } else {
                                v2 -= 0.1
                            }
                        default:
                            v32 := GenS3{F0: v3, F1: -1, F2: 42}
                            _ = v32
                        }
                    }
                }
            }
        }
        _ = fi9
    }
    switch -2147483648 {
    case 4294967295:
        v33 := GenS4{F0: 65535, F1: 0, F2: 207}
        _ = v33
    case -1:
        fmt.Println(v2, v2)
    case 128:
        v2 -= v2
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        _ = len(v4)
        for fi35 := 0; fi35 < 11; fi35++ {
            v5[fi34] = "a"
            if (!(!("\n\t" >= ""))) {
                v36 := GenS3{F0: 0, F1: v3, F2: 1}
                v8 = v8
                _ = v36
            } else {
                v37 := GenS3{F0: v3, F1: (v3 ^ v3), F2: 42}
                var v38 int32 = -2147483648
                _ = v37
                _ = v38
            }
            if v1 {
                v6 += 0.1
                fmt.Println(v3)
                v39 := GenS5{F0: 255, F1: fi35, F2: 42}
                _ = v39
            }
            _ = fi35
        }
        _ = fi34
    }
    fmt.Println(v1, v3, v2)
    switch "hello" {
    case "a":
        v8 += v8
    default:
        var v40 uint64 = 0
        _ = v40
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 uint32 = 42
        v43 := GenS6{F0: v8, F1: (v3 / ((v3) | 1))}
        _ = fi41
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
