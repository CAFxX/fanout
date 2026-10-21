package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -396
    var v3 int8 = -108
    v4 := []uint32{42}
    v5 := v1[2147483647]
    if true {
        if true {
            if len(v4) > 0 {
                v4[0] = 38
            }
        }
        var v6 float32 = 0.0
        _ = v6
    }
    switch (v5 ^ 127) {
    case 9223372036854775807:
        v3 = (v3 - -1)
    }
    _ = len(v4)
    var v7 uint32 = 0
    _ = len(v1)
    v7 -= (v7 + (v7 - 2))
    for fi8 := 0; fi8 < 4; fi8++ {
        if (false || true) {
            v5 = 2147483647
        }
        switch ("hello" + ("a" + "a")) {
        case "":
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(v5)
                switch "\n\t" {
                case "":
                    v10 := v1[(fi8 / (((fi8 * v5)) | 1))]
                    _ = v10
                case "foo bar":
                    var v11 uint = 4294967295
                    v3 += v3
                    _ = v11
                case "hello":
                    if (1 < 0) {
                        if true {
                            var v12 int = fi9
                            _ = v12
                        } else {
                            switch 32767 {
                            case 112:
                                v5 += 128
                            case 127:
                                v4 = append(v4, 2147483647)
                            case -2147483648:
                                _ = len(v4)
                            default:
                                var v13 uint16 = 134
                                _ = v13
                            }
                            fmt.Println(v5, v7)
                            v14 := GenS1{F0: 80, F1: (v7 | v7)}
                            _ = v14
                        }
                    } else {
                        var v15 uint64 = 2
                        v16 := GenS2{F0: 42, F1: 220, F2: 255}
                        _ = len(v4)
                        _ = v15
                        _ = v16
                    }
                }
                fmt.Println(fi9, fi9, v5)
                switch "\n\t" {
                case "hello":
                    v3 -= v3
                default:
                    switch "foo bar" {
                    case "foo bar":
                        v17 := GenS1{F0: -9223372036854775808, F1: (v7 - v7)}
                        _ = v17
                    case "\n\t":
                        var v18 bool = true
                        _ = v18
                    }
                }
                v2 = v2
                _ = fi9
            }
        case "\n\t":
            for fi19 := 0; fi19 < 2; fi19++ {
                if (2147483647 > (^(v7))) {
                    switch "hello" {
                    case "foo bar":
                        var v20 uint = 255
                        _ = v20
                    }
                    var v21 float64 = 0.0
                    var v22 bool = false
                    _ = v21
                    _ = v22
                }
                if true {
                    v3 -= (v3 / (((v3 >> ((4294967295) & 63))) | 1))
                    v23 := GenS3{F0: true}
                    _ = v23
                } else {
                    v7 = (v7 << ((13) & 63))
                    fmt.Println(v3)
                }
                if true {
                    var v24 float32 = 0.0
                    var v25 bool = ((!((v2 | (v2 / ((v2) | 1))) >= v2)) || (int32(fi8) <= -2147483648))
                    _ = v24
                    _ = v25
                } else {
                    v4 = append(v4, 2147483648)
                }
                _ = fi19
            }
        case "a":
            var v26 uint64 = 582
            _ = v26
        default:
            var v27 int = fi8
            _ = v27
        }
        _ = fi8
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        if ((!(false && (334 > 1))) && false) {
            v5 = 2147483647
            var v29 uint16 = 65535
            _ = len(v1)
            _ = v29
        } else {
            var v30 bool = (false || (820 < 42))
            fmt.Println(v5, v5, v7)
            _ = v30
        }
        v2 = v2
        _ = fi28
    }
    var v31 int16 = (-(v2))
    var v32 string = ("hello" + "\n\t")
    var v33 int16 = -32768
    v31 -= (v2 % ((-1) | 1))
    var v34 int64 = 0
    v5 = 4294967295
    var v35 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v7
}
