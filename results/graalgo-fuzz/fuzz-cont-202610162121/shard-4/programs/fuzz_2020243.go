package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint; F2 int }

func main() {
    var v1 uint32 = 2147483647
    var v2 int = -1
    var v3 uint64 = 684
    var v4 int16 = -1
    var v5 int8 = 21
    fmt.Println(v4, v2, v4)
    var v6 int32 = 697
    v7 := []int{}
    if true {
        if (!((v3 < 18446744073709551615) || ("" == "\n\t"))) {
            var v8 int16 = -1
            _ = v8
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint = 138
            v10 += (v10 % (((v10 * 42)) | 1))
            var v11 float64 = ((float64(-2.25) - -2.25) * 0.0)
            _ = fi9
            _ = v10
            _ = v11
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    switch "" {
    case "hello":
        v5 += int8(v1)
    case "a":
        fmt.Println(v6, v3, v6)
    case "foo bar":
        v12 := GenS1{F0: ""}
        _ = v12
    default:
        if (v5 >= -70) {
            v7 = append(v7, (v2 * (v2 ^ int(v6))))
        } else {
            switch v2 {
            case 9223372036854775807:
                var v13 int32 = 0
                _ = v13
            case -455:
                var v14 int64 = 608
                _ = v14
            case -831:
                _ = len(v7)
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int64 = 622
                for fi17 := 0; fi17 < 3; fi17++ {
                    if false {
                        var v18 int = fi15
                        _ = v18
                    } else {
                        _ = len(v7)
                        if false {
                            switch (fi17 >> ((4294967295) & 63)) {
                            case 42:
                                switch (-(fi15)) {
                                case 127:
                                    switch "a" {
                                    case "hello":
                                        v4 -= (v4 << ((370) & 63))
                                    }
                                case -9223372036854775808:
                                    v5 = (^(v5))
                                case -343:
                                    switch (fi15 % ((v2) | 1)) {
                                    case -9223372036854775808:
                                        switch "\n\t" {
                                        case "a":
                                            var v19 string = ("" + "hello")
                                            _ = v19
                                        case "":
                                            fmt.Println(v5, v4, v5)
                                            v20 := GenS2{F0: (v16 ^ (v16 / ((v16) | 1)))}
                                            _ = v20
                                        case "hello":
                                            v4 += (v4 + 1000)
                                        default:
                                            if (v3 <= 0) {
                                                v4 -= (v4 * v4)
                                            }
                                        }
                                    case 127:
                                        var v21 int = 4294967295
                                        _ = v21
                                    case 32767:
                                        v5 = v5
                                        v22 := GenS2{F0: (v16 + v16)}
                                        _ = v22
                                    }
                                }
                            case 158:
                                var v23 float32 = (1.5 * 3.14159)
                                _ = v23
                            case 1:
                                var v24 int8 = -128
                                _ = v24
                            }
                            v25 := GenS3{F0: v5, F1: 604, F2: (^(fi17))}
                            _ = v25
                        }
                        if len(v7) > 0 {
                            v7[0] = -807
                        }
                    }
                    fmt.Println(v4, v2)
                    _ = len(v7)
                    _ = fi17
                }
                _ = fi15
                _ = v16
            }
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        _ = len(v7)
        var v27 uint64 = v3
        if false {
            for fi28 := 0; fi28 < 9; fi28++ {
                fmt.Println(fi26, v5, v1)
                if false {
                    var v29 int64 = -103
                    var v30 int8 = v5
                    var v31 int = (fi26 / ((fi26) | 1))
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
        }
        _ = fi26
        _ = v27
    }
    var v32 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
