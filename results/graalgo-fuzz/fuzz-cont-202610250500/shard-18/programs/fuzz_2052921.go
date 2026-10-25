package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 string }

func main() {
    var v1 float64 = (((float64(1.5) * (1.5 + (3.14159 - float64(1e10)))) + 1.5) * 1e10)
    var v2 float32 = 1.5
    v3 := make(map[int]uint32)
    v4 := []uint32{444, 0}
    var v5 float64 = ((1e10 + ((float64((0.0 / (v2 + (v2 * v2)))) / v1) - (0.0 + (v1 + v1)))) * float64(0.0))
    var v6 uint16 = 255
    v5 = v1
    v7 := GenS1{F0: 18446744073709551615, F1: uint(v6), F2: 104}
    var v8 uint16 = (^(v6))
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            switch "hello" {
            case "":
                v1 = v1
            case "foo bar":
                var v11 uint8 = 40
                v2 += float32((v5 + ((1e10 / 3.14159) + v5)))
                _ = v11
            case "\n\t":
                switch "hello" {
                case "\n\t":
                    var v12 uint8 = 1
                    _ = v12
                case "a":
                    var v13 uint64 = 4294967295
                    if false {
                        var v14 uint32 = 1
                        if ((!(((329 <= 681) && true) || false)) || false) {
                            fmt.Println(v2)
                            fmt.Println(v5)
                            var v15 uint64 = 18446744073709551615
                            _ = v15
                        } else {
                            switch (("hello" + "") + "hello") {
                            case "foo bar":
                                v5 -= (float64(v2) / ((1.5 - (float64((v2 + v2)) * (3.14159 / (float64((v2 * v2)) * (1e10 + (v5 - (0.0 - 0.0))))))) + 1.5))
                            case "\n\t":
                                switch -9223372036854775808 {
                                case 128:
                                    var v16 int64 = -1
                                    _ = v16
                                case -1:
                                    var v17 uint = 255
                                    _ = v17
                                case 1:
                                    var v18 float32 = v2
                                    _ = v18
                                default:
                                    var v19 int64 = 42
                                    fmt.Println(v13, v14, v1)
                                    _ = v19
                                }
                            case "a":
                                var v20 int16 = -12
                                _ = v20
                            }
                        }
                        fmt.Println(v5)
                        fmt.Println(fi10, fi10, v6)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v14
                    }
                    switch 61 {
                    case 9223372036854775807:
                        _ = len(v4)
                    case 127:
                        v21 := GenS2{F0: 4294967295, F1: 0, F2: v6}
                        _ = v21
                    case 4294967295:
                        var v22 uint = 2
                        _ = v22
                    }
                    _ = v13
                case "":
                    v2 = v2
                }
            }
            v1 -= (v5 - v5)
            v23 := GenS1{F0: 4294967295, F1: 937, F2: 42}
            _ = fi10
            _ = v23
        }
        v6 -= (v8 + (v8 >> ((63) & 63)))
        _ = fi9
    }
    if false {
        v24 := GenS3{F0: 4294967295, F1: 42}
        for fi25 := 0; fi25 < 9; fi25++ {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v26 := GenS4{F0: ""}
                _ = v26
            }
            var v27 float64 = 1e10
            _ = fi25
            _ = v27
        }
        if true {
            var v28 uint32 = 4294967295
            for fi29 := 0; fi29 < 12; fi29++ {
                v8 += 0
                var v30 uint64 = 748
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        _ = v24
    } else {
        var v31 int16 = -738
        _ = v31
    }
    if (true && (v2 == (float32(1.5) * v2))) {
        var v32 bool = (v8 < v6)
        var v33 int = -9223372036854775808
        delete(v3, -1)
        _ = v32
        _ = v33
    }
    fmt.Println(v1, v2)
    fmt.Println(v5, v2, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
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
