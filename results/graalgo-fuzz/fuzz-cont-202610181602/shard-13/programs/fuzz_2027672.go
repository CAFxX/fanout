package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int }
func genFn0() int8 {
    v6 := GenS1{F0: 2}
    if false {
        var v7 int32 = 1
        _ = v7
    }
    _ = v6
    return 1
}
func genFn1(p0 int8) uint64 {
    if true {
        p0 -= genFn0()
    }
    if (true || (0.1 < -2.25)) {
        p0 += (p0 / ((42) | 1))
        var v8 int = 136
        _ = v8
    }
    _ = p0
    return 0
}

func main() {
    var v1 float64 = float64(0.0)
    v2 := []int{128, 716, -290, 805}
    v3 := make(map[string]int)
    var v4 int = 128
    var v5 string = "hello"
    if true {
        if true {
            for fi9 := 0; fi9 < 9; fi9++ {
                v1 = -2.25
                _ = fi9
            }
            v1 = v1
        } else {
            v10 := GenS2{F0: 1000}
            v11 := GenS3{F0: "foo bar", F1: 0}
            switch v4 {
            case 42:
                if true {
                    _ = len(v2)
                    var v12 int16 = -32768
                    _ = len(v3)
                    _ = v12
                }
            default:
                fmt.Println(v1, v1, v4)
            }
            _ = v10
            _ = v11
        }
        v5 = "a"
    } else {
        fmt.Println(v5)
        for fi13 := 0; fi13 < 6; fi13++ {
            switch v5 {
            case "foo bar":
                for fi14 := 0; fi14 < 14; fi14++ {
                    fmt.Println(fi14, fi14, fi14)
                    switch fi13 {
                    case 127:
                        var v15 int16 = -32768
                        _ = v15
                    default:
                        var v16 int = (fi14 >> ((550) & 63))
                        v17 := GenS1{F0: 749}
                        _ = v16
                        _ = v17
                    }
                    if true {
                        if ((((0.0 / (1.5 * ((v1 + v1) - float64(1e10)))) * float64(0.0)) + -2.25) != (float64(float32(1e10)) / ((v1 / float64(float32(v1))) * v1))) {
                            if true {
                                var v18 bool = true
                                if (!v18) {
                                    v1 += (v1 - v1)
                                    var v19 uint32 = 14
                                    _ = v19
                                }
                                var v20 uint8 = 128
                                _ = v18
                                _ = v20
                            }
                            fmt.Println(fi14, v4)
                        } else {
                            var v21 uint = 2
                            _ = v21
                        }
                        v22 := v3[v5]
                        switch "hello" {
                        case "":
                            var v23 int = v22
                            fmt.Println(fi13, v5, v22)
                            _ = v23
                        case "\n\t":
                            if (2147483647 != 42) {
                                v24 := GenS2{F0: 1}
                                _ = v24
                            } else {
                                fmt.Println(v4)
                            }
                            v4 -= (v22 % ((v22) | 1))
                        }
                        _ = v22
                    } else {
                        v2 = append(v2, (fi13 / ((42) | 1)))
                    }
                    _ = fi14
                }
            case "hello":
                var v25 float64 = (0.0 / float64(1.5))
                _ = v25
            }
            _ = fi13
        }
        if len(v2) > 0 {
            v2[0] = (v4 + (-(v4)))
        }
    }
    var v26 uint16 = 2
    v27 := GenS4{F0: (v4 / ((9223372036854775807) | 1))}
    var v28 int16 = 836
    for fi29 := 0; fi29 < 10; fi29++ {
        v30 := v3[v5]
        _ = fi29
        _ = v30
    }
    if false {
        switch (v5 + "foo bar") {
        case "\n\t":
            var v31 uint8 = 255
            _ = v31
        default:
            v32 := v3[v5]
            _ = v32
        }
        if (!(!true)) {
            var v33 int32 = 1
            _ = v33
        }
    }
    switch v4 {
    case 1:
        v4 = (v4 & (v4 & -1))
    case -609:
        v4 -= (v4 << ((255) & 63))
    case -1:
        switch "foo bar" {
        case "hello":
            var v34 int8 = -87
            _ = v34
        default:
            var v35 int64 = 722
            _ = v35
        }
    }
    var v36 int16 = 0
    v1 += v1
    var v37 float32 = (1e10 / 1.5)
    var v38 int16 = (^(v36))
    v39 := GenS3{F0: v5, F1: (v4 << ((18446744073709551615) & 63))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += v1
    if true {
        fmt.Println(v28)
    } else {
        fmt.Println(v36, v38, v37)
        var v40 int = (v4 / ((v4) | 1))
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
