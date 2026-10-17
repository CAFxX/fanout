package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint8 }
func genFn0(p0 bool, p1 int32, p2 int) int64 {
    var v4 uint8 = 203
    fmt.Println(p2, v4, v4)
    v5 := []int{(p2 >> ((4294967295) & 63)), int(p1), (p2 / (((p2 >> ((797) & 63))) | 1))}
    for fi6 := 0; fi6 < 19; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if ((888 < 1) && true) {
            var v7 int = (-(p2))
            _ = v7
        } else {
            var v8 float32 = float32((float64(1.5) + ((0.1 + (3.14159 - ((0.1 - float64(((((3.14159 + (0.1 - 0.1)) - 0.1) / -2.25) / 1.0))) - 0.1))) / (float64(1e10) - 0.1))))
            var v9 int8 = 0
            var v10 int16 = -32768
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}
func genFn1() int64 {
    v11 := []int{}
    for fi12 := 0; fi12 < 17; fi12++ {
        v11 = append(v11, (fi12 % ((283) | 1)))
        var v13 uint64 = 162
        _ = fi12
        _ = v13
    }
    _ = v11
    return 0
}

func main() {
    var v1 int32 = -112
    var v2 uint64 = 4294967295
    v3 := make(map[int]string)
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := []int{fi14}
        fmt.Println(v1, v2)
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: (v1 / ((2147483647) | 1)), F1: (v1 | v1), F2: false}
    switch "a" {
    case "a":
        var v17 float64 = 0.1
        _ = v17
    default:
        v18 := GenS1{F0: v1, F1: v1, F2: (42 != 1000)}
        _ = v18
    }
    _ = len(v3)
    v2 = v2
    v19 := []int{-988, 1, 2147483647}
    fmt.Println(v1, v2)
    switch "foo bar" {
    case "a":
        var v20 int8 = 39
        _ = v20
    case "foo bar":
        if (205 == 183) {
            if false {
                var v21 int64 = -9223372036854775808
                switch (("" + "a") + "\n\t") {
                case "":
                    var v22 int32 = 2147483647
                    _ = v22
                case "foo bar":
                    if true {
                        for fi23 := 0; fi23 < 7; fi23++ {
                            if len(v19) > 0 {
                                v19[0] = fi23
                            }
                            _ = fi23
                        }
                        switch "foo bar" {
                        case "hello":
                            v24 := GenS2{F0: -250}
                            _ = v24
                        case "\n\t":
                            fmt.Println(v2, v1, v1)
                        case "foo bar":
                            v21 = v21
                        }
                    }
                }
                var v25 uint = 718
                _ = v21
                _ = v25
            } else {
                if (!true) {
                    v26 := GenS3{F0: 2, F1: 151, F2: v1}
                    switch ("hello" + "hello") {
                    case "\n\t":
                        v2 = 160
                    case "hello":
                        var v27 uint16 = 42
                        _ = v27
                    default:
                        if false {
                            v1 -= (v1 >> ((18446744073709551615) & 63))
                            if true {
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    _ = len(v3)
                                    v2 -= 42
                                    _ = fi28
                                }
                                v29 := GenS4{F0: -1, F1: 442}
                                v19 = append(v19, -831)
                                _ = v29
                            }
                            if (!true) {
                                var v30 float32 = 0.1
                                v1 = (v1 / (((v1 ^ v1)) | 1))
                                _ = v30
                            }
                        }
                    }
                    _ = v26
                }
                if len(v19) > 0 {
                    v19[0] = 2147483647
                }
            }
        } else {
            var v31 uint32 = 467
            var v32 uint64 = (v2 * (v2 + v2))
            if len(v19) > 0 {
                v19[0] = 353
            }
            _ = v31
            _ = v32
        }
    case "":
        _ = len(v19)
    }
    fmt.Println(v2, v1)
    if true {
        for fi33 := 0; fi33 < 17; fi33++ {
            v2 = v2
            v34 := GenS5{F0: 32767, F1: 1}
            v2 -= v2
            _ = fi33
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
