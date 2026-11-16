package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int16 }
func genFn0(p0 string) uint {
    p0 = p0
    v4 := []int{}
    v5 := GenS1{F0: 100, F1: 161, F2: 0}
    _ = p0
    _ = v4
    _ = v5
    return 634
}

func main() {
    var v1 uint32 = 738
    var v2 int32 = 2147483647
    var v3 int64 = 9223372036854775807
    v6 := GenS2{F0: 168}
    v7 := []int{764, -2147483648, -2147483648}
    fmt.Println(v2, v1, v1)
    v1 = (v1 * 42)
    switch ("foo bar" + "hello") {
    case "foo bar":
        v2 -= v2
    default:
        v8 := GenS2{F0: 128}
        _ = v8
    }
    _ = len(v7)
    if (-1 >= 127) {
        var v9 uint = 18446744073709551615
        _ = v9
    }
    v10 := GenS3{F0: genFn0(""), F1: (v1 % ((65535) | 1)), F2: 589}
    v1 += 2147483647
    var v11 int32 = (v2 / ((2147483647) | 1))
    var v12 string = "\n\t"
    var v13 int32 = 2147483647
    for fi14 := 0; fi14 < 18; fi14++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v15 := GenS1{F0: 1, F1: 348, F2: v3}
        fmt.Println(v11, fi14, v1)
        _ = fi14
        _ = v15
    }
    v2 = v2
    var v16 float64 = float64((float32(((0.1 + 1.5) * (0.1 * (1.5 * (3.14159 + (1.5 / 1e10)))))) * 1.5))
    for fi17 := 0; fi17 < 15; fi17++ {
        v16 -= 1.5
        if (false && (65535 != (v1 + v1))) {
            var v18 float64 = v16
            if true {
                var v19 uint = 4294967295
                _ = v19
            } else {
                v7 = append(v7, (fi17 * 4294967295))
            }
            if (42 > (v3 & v3)) {
                if false {
                    var v20 int32 = v13
                    _ = v20
                }
                switch ((v12 + (((("hello" + v12) + v12) + (v12 + v12)) + "foo bar")) + "a") {
                case "\n\t":
                    for fi21 := 0; fi21 < 13; fi21++ {
                        var v22 uint = 1
                        if ((false && true) || (-2147483648 >= 32768)) {
                            var v23 bool = (0 < 743)
                            if (!(-1 > 42)) {
                                v11 -= (^(v2))
                                v12 = (v12 + v12)
                                v1 -= v1
                                if true {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                } else {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    fmt.Println(v1)
                                    v2 = v2
                                }
                            } else {
                                v24 := GenS2{F0: 2}
                                _ = v24
                            }
                            var v25 uint64 = 160
                            switch "foo bar" {
                            case "\n\t":
                                if true {
                                    switch fi21 {
                                    case -667:
                                        var v26 int64 = (v3 ^ 0)
                                        _ = v26
                                    }
                                    var v27 int16 = 1
                                    _ = len(v7)
                                    var v28 uint16 = 32767
                                    v29 := GenS3{F0: 0, F1: (^(v1)), F2: 0}
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                } else {
                                    fmt.Println(v1, v13, v2)
                                    v11 = 42
                                }
                            default:
                                if len(v7) > 0 {
                                    v7[0] = fi21
                                }
                            }
                            _ = v23
                            _ = v25
                        }
                        _ = fi21
                        _ = v22
                    }
                case "":
                    var v30 int = fi17
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 15; fi31++ {
                        var v32 int32 = 2147483647
                        v2 = v13
                        var v33 uint32 = 4294967295
                        v34 := GenS2{F0: 2}
                        v35 := GenS1{F0: 42, F1: 44, F2: v3}
                        v13 = v2
                        fmt.Println(v1)
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
            }
            _ = v18
        } else {
            fmt.Println(v13)
        }
        _ = fi17
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
