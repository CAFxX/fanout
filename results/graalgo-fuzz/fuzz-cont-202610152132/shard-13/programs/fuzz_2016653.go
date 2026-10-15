package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 bool; F1 int16; F2 int32 }
func genFn0(p0 string) int32 {
    v5 := GenS1{F0: 42, F1: 18446744073709551615}
    if false {
        v6 := GenS2{F0: 422, F1: 1}
        var v7 bool = ((-2.25 < 1e10) || true)
        fmt.Println(p0)
        _ = v6
        _ = v7
    }
    var v8 float32 = (((1.5 / 1e10) + float32((float64(0.0) * 1e10))) - (0.1 * ((-2.25 / 1.5) * 0.1)))
    v9 := []int{0, 285}
    switch p0 {
    case "\n\t":
        p0 = p0
    case "foo bar":
        var v10 uint64 = 18446744073709551615
        _ = v10
    case "":
        switch -345 {
        case 9223372036854775807:
            fmt.Println(v8)
        }
    }
    _ = p0
    _ = v5
    _ = v8
    _ = v9
    return -112
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    v4 := []uint32{2147483648, 42, 532, 90}
    v1 = v1
    var v11 float32 = 0.1
    fmt.Println(v11, v1)
    v12 := GenS3{F0: false, F1: "a"}
    if false {
        if false {
            fmt.Println(v1, v1)
        }
        v13 := v2[42]
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        if (2 >= 346) {
            v15 := GenS1{F0: 2, F1: 18446744073709551615}
            var v16 float32 = -2.25
            fmt.Println(v16, fi14)
            _ = v15
            _ = v16
        } else {
            fmt.Println(v11)
            switch 32768 {
            case -2147483648:
                v17 := GenS1{F0: 255, F1: 0}
                _ = v17
            }
        }
        _ = fi14
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v1 = (-(v1))
        switch "" {
        case "hello":
            v1 = v1
        case "\n\t":
            fmt.Println(v1, fi18, v11)
        case "foo bar":
            v11 = ((-2.25 / v11) - (((v11 * v11) - v11) / v11))
        default:
            var v19 uint8 = 22
            _ = v19
        }
        v20 := GenS4{F0: 42}
        _ = fi18
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        switch "a" {
        case "hello":
            if true {
                if len(v4) > 0 {
                    v4[0] = 65535
                }
                switch (fi21 & fi21) {
                case 2147483647:
                    v4 = append(v4, 324)
                case -1:
                    v22 := GenS5{F0: -1, F1: 32767}
                    _ = v22
                default:
                    var v23 float64 = float64(float32(0.0))
                    _ = v23
                }
                if false {
                    var v24 float32 = (((-2.25 / v11) / 3.14159) - v11)
                    v25 := GenS3{F0: (((0.0 + 1.5) != float64(v24)) && (!false)), F1: "foo bar"}
                    _ = v24
                    _ = v25
                } else {
                    var v26 int64 = 424
                    switch 42 {
                    case 0:
                        switch (fi21 >> ((4294967295) & 63)) {
                        case -185:
                            var v27 uint16 = 288
                            _ = v27
                        case -1:
                            fmt.Println(v26)
                        case 127:
                            v26 += v26
                        default:
                            for fi28 := 0; fi28 < 1; fi28++ {
                                fmt.Println(v11, v1)
                                if (!false) {
                                    var v29 int32 = -2147483648
                                    v1 = 2147483647
                                    _ = v29
                                }
                                var v30 int32 = (v1 & 2147483647)
                                _ = fi28
                                _ = v30
                            }
                        }
                    default:
                        v2[fi21] = "a"
                    }
                    _ = v26
                }
            }
        case "a":
            v11 -= v11
        case "\n\t":
            for fi31 := 0; fi31 < 12; fi31++ {
                fmt.Println(fi21)
                v32 := GenS2{F0: 42, F1: 1}
                var v33 int16 = -32768
                v1 -= 42
                switch fi31 {
                case 32768:
                    fmt.Println(fi21, v11, v1)
                default:
                    v11 += (0.0 + -2.25)
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        default:
            var v34 uint = 42
            _ = v34
        }
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 int64 = 0
            var v37 uint32 = 907
            if false {
                v37 = 726
            }
            v1 = (v1 ^ v1)
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi21
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        v39 := GenS6{F0: true, F1: 42, F2: (^(v1))}
        v40 := GenS4{F0: -62}
        _ = fi38
        _ = v39
        _ = v40
    }
    if false {
        var v41 uint = 18446744073709551615
        var v42 uint32 = 448
        if len(v4) > 0 {
            v4[0] = (v42 - (v42 % ((v42) | 1)))
        }
        _ = v41
        _ = v42
    }
    v11 -= float32(3.14159)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
