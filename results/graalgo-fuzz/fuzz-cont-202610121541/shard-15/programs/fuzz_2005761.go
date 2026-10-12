package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 bool = false
    v1 -= v1
    v1 += (v1 | 158)
    if v4 {
        if true {
            var v5 string = "hello"
            _ = v5
        }
        if (v4 || true) {
            v4 = (42 < 0)
            if (!(176 == 4)) {
                var v6 int64 = -1
                var v7 float32 = -2.25
                _ = v6
                _ = v7
            } else {
                v1 = 42
            }
            v1 = (v1 * v1)
        } else {
            if (42 >= 42) {
                v1 = 0
                v8 := GenS1{F0: 0, F1: -447}
                _ = v8
            }
            var v9 float64 = -2.25
            _ = v9
        }
    } else {
        if (((65535 < 4294967295) && false) && v4) {
            var v10 int8 = 1
            _ = v10
        } else {
            v4 = false
            if v4 {
                v1 = 127
                if (3.14159 > (((1e10 + (0.1 * -2.25)) + 0.0) * (3.14159 - 1e10))) {
                    v4 = v4
                }
                v11 := []int{9223372036854775807, -2147483648, -1, 127}
                _ = v11
            } else {
                v12 := GenS1{F0: 4294967295, F1: -163}
                switch 2147483647 {
                case 128:
                    var v13 int16 = -420
                    _ = v13
                }
                v1 = (v1 / (((v1 >> ((65535) & 63))) | 1))
                _ = v12
            }
        }
        v1 -= v1
        if false {
            if ((0 == 1) || v4) {
                fmt.Println(v4)
                if (2147483647 != 42) {
                    for fi14 := 0; fi14 < 4; fi14++ {
                        v15 := GenS2{F0: fi14}
                        if (!(947 <= 2)) {
                            v16 := []int{(^(fi14)), -248, (fi14 & -2147483648)}
                            _ = v16
                        }
                        if true {
                            for fi17 := 0; fi17 < 13; fi17++ {
                                v4 = (uint(v1) == 184)
                                fmt.Println(fi14, v1, v4)
                                _ = fi17
                            }
                            v4 = (!true)
                            for fi18 := 0; fi18 < 20; fi18++ {
                                var v19 uint = 955
                                if (-4 == 127) {
                                    v20 := GenS3{F0: 1, F1: 2147483647}
                                    switch ("" + "hello") {
                                    case "hello":
                                        var v21 uint8 = (v1 - v1)
                                        switch "hello" {
                                        case "hello":
                                            v22 := GenS1{F0: 289, F1: 1000}
                                            _ = v22
                                        case "foo bar":
                                            v4 = v4
                                        }
                                        var v23 int8 = -81
                                        _ = v21
                                        _ = v23
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        v24 := []int{fi14, (fi14 % (((-(fi18))) | 1)), fi18}
                                        _ = v24
                                    }
                                    v4 = v4
                                    v1 += (^(v1))
                                    _ = v20
                                } else {
                                    var v25 uint8 = v1
                                    v26 := GenS4{F0: 2}
                                    _ = v25
                                    _ = v26
                                }
                                var v27 uint64 = 4294967295
                                _ = fi18
                                _ = v19
                                _ = v27
                            }
                        }
                        _ = fi14
                        _ = v15
                    }
                    v28 := GenS1{F0: 1, F1: -730}
                    _ = v28
                }
            }
            fmt.Println(v1, v1)
            var v29 string = "a"
            _ = v29
        }
    }
    fmt.Println(v4, v1)
    v1 -= (v1 + 243)
    v30 := GenS2{F0: 584}
    if true {
        v31 := GenS5{F0: 255}
        v2[42] = 2
        var v32 uint = 977
        _ = v31
        _ = v32
    }
    if (!(2147483647 < 2147483647)) {
        var v33 int = int(v1)
        _ = v33
    }
    switch 42 {
    case 0:
        v34 := []int{173, 4294967295}
        _ = v34
    case 237:
        fmt.Println(v1, v1)
    }
    var v35 int8 = -128
    if v4 {
        v4 = v4
        switch ("foo bar" + "hello") {
        case "hello":
            var v36 uint8 = (^(v1))
            _ = v36
        case "":
            var v37 string = "\n\t"
            _ = v37
        }
        v38 := []int{-901, 279, 9223372036854775807}
        _ = v38
    }
    switch 9223372036854775807 {
    case 489:
        var v39 string = "foo bar"
        _ = v39
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        if (v4 && (v1 > 0)) {
            v41 := []int{}
            fmt.Println(v1, v1, v4)
            _ = v41
        } else {
            v42 := GenS1{F0: 118, F1: -1}
            if v4 {
                var v43 uint64 = 0
                var v44 int64 = -444
                var v45 int32 = -6
                _ = v43
                _ = v44
                _ = v45
            }
            _ = v42
        }
        _ = fi40
    }
    v35 += (v35 | -51)
    v46 := GenS5{F0: (^(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v46
}
