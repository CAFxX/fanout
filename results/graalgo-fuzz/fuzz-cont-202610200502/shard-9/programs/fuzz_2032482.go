package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 float32 = ((0.1 / 0.1) + 3.14159)
    v2 := make(map[int]int)
    v3 := []string{}
    fmt.Println(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 := GenS1{F0: 42}
    v1 = v1
    v5 := v2[-2147483648]
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS2{F0: 42, F1: 104, F2: -1}
        if (2147483647 < 1) {
            switch "foo bar" {
            case "foo bar":
                for fi8 := 0; fi8 < 20; fi8++ {
                    if (float32(0.1) == 0.1) {
                        if true {
                            fmt.Println(fi8, v1)
                            if (311 >= -961) {
                                var v9 uint8 = 128
                                v9 += (^(v9))
                                var v10 bool = true
                                _ = v10
                                _ = v9
                            } else {
                                if true {
                                    var v11 int64 = 107
                                    var v12 uint = 2
                                    _ = v11
                                    _ = v12
                                }
                                var v13 bool = (false && false)
                                v1 -= ((v1 - v1) * v1)
                                _ = v13
                            }
                            var v14 int64 = 970
                            _ = v14
                        } else {
                            v15 := GenS1{F0: 9223372036854775807}
                            if (false && true) {
                                switch fi6 {
                                case 1:
                                    if true {
                                        v1 = v1
                                        v16 := GenS3{F0: 801, F1: -90}
                                        _ = v16
                                    } else {
                                        v5 = (fi8 << ((0) & 63))
                                    }
                                }
                            } else {
                                fmt.Println(v5, v1, v5)
                                if (0 <= 0) {
                                    fmt.Println(v5, v5)
                                } else {
                                    v5 += (fi8 / (((v5 ^ fi6)) | 1))
                                    v2[fi8] = fi8
                                }
                            }
                            if false {
                                v5 -= -2147483648
                                v17 := GenS3{F0: 2, F1: 100}
                                var v18 float32 = (((v1 - v1) - 1e10) - v1)
                                _ = v17
                                _ = v18
                            }
                            _ = v15
                        }
                        fmt.Println(v5)
                        v19 := GenS1{F0: 9223372036854775807}
                        _ = v19
                    }
                    fmt.Println(v1)
                    _ = fi8
                }
            default:
                for fi20 := 0; fi20 < 10; fi20++ {
                    if (!false) {
                        var v21 int64 = 42
                        v22 := GenS2{F0: -259, F1: 1, F2: -32768}
                        var v23 uint = 979
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (("hello" + "hello") + "foo bar")
                        }
                        switch ("hello" + "foo bar") {
                        case "foo bar":
                            fmt.Println(fi6, v5)
                        case "a":
                            v24 := GenS2{F0: v5, F1: 0, F2: -1}
                            _ = v24
                        default:
                            fmt.Println(fi6)
                        }
                        fmt.Println(v1)
                        v25 := GenS1{F0: 174}
                        _ = v25
                    }
                    v26 := GenS1{F0: 42}
                    _ = fi20
                    _ = v26
                }
            }
            _ = len(v2)
            fmt.Println(v1, fi6, v1)
        }
        _ = fi6
        _ = v7
    }
    v5 = (v5 >> ((18446744073709551615) & 63))
    var v27 uint16 = 1
    var v28 string = ""
    switch ("a" + "foo bar") {
    case "a":
        var v29 int64 = 1
        _ = v29
    case "hello":
        delete(v2, v5)
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = v28
        }
    default:
        v5 = (v5 >> ((42) & 63))
    }
    if true {
        v1 = (0.0 + float32(1e10))
        for fi30 := 0; fi30 < 10; fi30++ {
            var v31 int = v5
            _ = fi30
            _ = v31
        }
    }
    v32 := GenS4{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
