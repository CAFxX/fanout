package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 bool; F2 int8 }
func genFn0(p0 uint32) int16 {
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    p0 += 0
    v4 := GenS1{F0: true, F1: 707}
    _ = p0
    _ = v4
    return -32768
}
func genFn1() uint8 {
    var v5 uint64 = 18446744073709551615
    var v6 string = "\n\t"
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 255
    var v3 int16 = 1
    v7 := []int{127, 32767, 1, -517}
    v8 := GenS2{F0: 128, F1: v1}
    var v9 uint32 = 42
    fmt.Println(v3, v2)
    var v10 uint64 = 4294967295
    var v11 uint64 = 672
    switch "\n\t" {
    case "\n\t":
        switch "a" {
        case "foo bar":
            switch 32768 {
            case -1:
                var v12 uint = v1
                _ = v12
            default:
                var v13 float64 = float64(1.5)
                _ = v13
            }
        case "\n\t":
            v14 := GenS1{F0: (0 > 32767), F1: -2147483648}
            _ = v14
        case "":
            v7 = append(v7, -2147483648)
        default:
            if false {
                var v15 int8 = 127
                if true {
                    if true {
                        var v16 int = -1
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v18 := GenS2{F0: v16, F1: v1}
                            var v19 uint64 = uint64(v9)
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = v16
                    } else {
                        v20 := GenS3{F0: (^(v10)), F1: "a"}
                        _ = v20
                    }
                    var v21 string = "a"
                    _ = v21
                } else {
                    v7 = append(v7, 32768)
                    if (true && (float32((3.14159 - 0.1)) <= 1.5)) {
                        if (32767 != 0) {
                            var v22 string = "\n\t"
                            for fi23 := 0; fi23 < 19; fi23++ {
                                if (!true) {
                                    switch v22 {
                                    case "a":
                                        if ((1 >= 0) && true) {
                                            for fi24 := 0; fi24 < 10; fi24++ {
                                                var v25 uint = uint(v9)
                                                var v26 bool = false
                                                var v27 int16 = v3
                                                _ = fi24
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            }
                                            for fi28 := 0; fi28 < 7; fi28++ {
                                                var v29 int16 = (v3 / ((-117) | 1))
                                                if true {
                                                    var v30 uint32 = v9
                                                    _ = v30
                                                }
                                                _ = fi28
                                                _ = v29
                                            }
                                        } else {
                                            fmt.Println(v1)
                                            var v31 int64 = 862
                                            _ = v31
                                        }
                                    case "hello":
                                        var v32 int64 = -1
                                        _ = v32
                                    default:
                                        var v33 uint64 = (v11 + uint64(v1))
                                        _ = v33
                                    }
                                }
                                _ = fi23
                            }
                            _ = v22
                        } else {
                            v34 := GenS4{F0: ((^(v2)) == uint8(v1))}
                            v10 -= (v11 + 1)
                            _ = v34
                        }
                    } else {
                        v1 = 2
                        var v35 uint16 = 451
                        _ = v35
                    }
                    switch "a" {
                    case "a":
                        v36 := GenS2{F0: -9223372036854775808, F1: v1}
                        _ = v36
                    default:
                        if len(v7) > 0 {
                            v7[0] = int(v3)
                        }
                    }
                }
                _ = v15
            } else {
                v37 := GenS5{F0: 9223372036854775807, F1: false, F2: 0}
                _ = v37
            }
        }
    case "":
        v1 += (v1 / ((v1) | 1))
    }
    if len(v7) > 0 {
        v7[0] = 42
    }
    switch ((("" + "a") + "\n\t") + ((("hello" + ("" + ("a" + "\n\t"))) + "foo bar") + "")) {
    case "\n\t":
        fmt.Println(v1, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
