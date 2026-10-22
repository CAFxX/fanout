package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int }

func main() {
    v1 := []uint32{65535, 65535, 795}
    var v2 uint = 2
    v3 := make(map[uint32]string)
    v4 := []string{"hello", ("" + "")}
    v2 = (v2 - (v2 * (^(v2))))
    if (!(((-2.25 * (0.0 / 3.14159)) < ((1.5 * -2.25) - 1.5)) && true)) {
        v2 += v2
        var v5 int32 = -1
        _ = v5
    } else {
        for fi6 := 0; fi6 < 20; fi6++ {
            for fi7 := 0; fi7 < 7; fi7++ {
                var v8 float32 = 1e10
                var v9 int16 = -1
                if false {
                    v3[83] = ("\n\t" + (("a" + "hello") + "\n\t"))
                    if (2 <= 4294967295) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 += (v2 - v2)
                    } else {
                        switch fi7 {
                        case -2147483648:
                            var v10 uint32 = 777
                            _ = v10
                        case -1:
                            _ = len(v3)
                            var v11 string = ("\n\t" + ("a" + ""))
                            if true {
                                var v12 float64 = (1.5 * 0.0)
                                var v13 int64 = -9223372036854775808
                                v12 += ((v12 * 0.1) + 0.1)
                                var v14 int64 = (v13 | (v13 | 0))
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                            _ = v11
                        case -414:
                            v9 = v9
                        }
                        v15 := GenS1{F0: 9223372036854775807, F1: 4294967295}
                        var v16 int = fi6
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v2)
                }
                var v17 int16 = v9
                v18 := GenS2{F0: 5, F1: 1}
                var v19 int64 = 9223372036854775807
                _ = fi7
                _ = v17
                _ = v18
                _ = v19
                _ = v8
                _ = v9
            }
            var v20 int64 = 1
            fmt.Println(v2, v2, v20)
            _ = fi6
            _ = v20
        }
    }
    v2 = v2
    if (127 == 996) {
        var v21 float64 = 1e10
        var v22 int = -2147483648
        var v23 float64 = v21
        _ = v21
        _ = v22
        _ = v23
    }
    switch "hello" {
    case "":
        fmt.Println(v2)
    case "\n\t":
        if (true && false) {
            var v24 int = 4294967295
            v24 = v24
            _ = v24
        } else {
            fmt.Println(v2)
        }
    case "hello":
        fmt.Println(v2)
    default:
        var v25 uint8 = 194
        _ = v25
    }
    var v26 uint8 = 223
    for fi27 := 0; fi27 < 4; fi27++ {
        fmt.Println(v2, v2, fi27)
        _ = fi27
    }
    _ = len(v3)
    for fi28 := 0; fi28 < 7; fi28++ {
        if (865 <= 65535) {
            if false {
                if true {
                    v29 := GenS3{F0: 911, F1: (fi28 % ((fi28) | 1))}
                    v26 -= (v26 * v26)
                    _ = v29
                } else {
                    var v30 int16 = 306
                    _ = v30
                }
                var v31 uint64 = 42
                v26 = v26
                _ = v31
            } else {
                v26 = (v26 / (((^(v26))) | 1))
            }
        }
        _ = fi28
    }
    v3[65535] = "a"
    v4 = append(v4, "foo bar")
    v32 := GenS1{F0: -9223372036854775808, F1: v2}
    var v33 int8 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
