package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0(p0 int64) int64 {
    v6 := []int{4294967295, 42, -9223372036854775808}
    var v7 int8 = -47
    v8 := GenS1{F0: 42, F1: 1}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 uint64, p1 bool, p2 uint64) uint32 {
    var v9 int16 = 42
    v10 := []int{int(p0)}
    if p1 {
        v9 -= v9
        v10 = append(v10, 0)
        if ((-2.25 - ((float32(0.1) * 0.0) * -2.25)) >= 1.5) {
            fmt.Println(p0, p2)
            v11 := GenS2{F0: false, F1: (!false)}
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 11; fi12++ {
            v9 -= (v9 & v9)
            _ = fi12
        }
        var v13 uint64 = (p0 | p2)
        var v14 string = "hello"
        _ = v13
        _ = v14
    }
    _ = len(v10)
    var v15 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v15
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint = 1
    v2 := []uint32{}
    var v3 int64 = -244
    var v4 float64 = 1.5
    v5 := []uint32{}
    v4 = (((-2.25 + (((v4 - ((v4 * v4) * (1.5 - float64(float32((v4 + float64((0.0 + (3.14159 + 0.1))))))))) * ((((((v4 / v4) / (v4 / v4)) - float64(float32(v4))) / 0.1) + (v4 + (1e10 / float64(3.14159)))) - v4)) + 0.0)) - v4) * v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 uint = 0
    var v17 bool = true
    if ((3.14159 / 0.1) <= float32(v4)) {
        switch 438 {
        case -9223372036854775808:
            v1 -= (^(v16))
        case 32767:
            v16 = (v1 * v1)
        default:
            var v18 uint = v1
            _ = v18
        }
    }
    if (-2147483648 <= -697) {
        var v19 uint32 = 2147483647
        if (-1 == 1) {
            v19 -= (^(v19))
            v20 := GenS2{F0: ((-2147483648 < -2) && true), F1: v17}
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 float64 = (v4 / (v4 / (v4 * (v4 / (v4 - (-2.25 - 1e10))))))
                var v23 int8 = int8(v19)
                fmt.Println(v23, v19)
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        }
        v16 = v1
        _ = v19
    }
    v24 := GenS1{F0: 1, F1: 128}
    v17 = ("foo bar" < "\n\t")
    for fi25 := 0; fi25 < 8; fi25++ {
        if ("hello" >= ((("foo bar" + "a") + "hello") + "")) {
            v16 = v1
            switch "hello" {
            case "":
                v16 -= (v1 * (^(v1)))
            default:
                if (0 > (fi25 >> ((v16) & 63))) {
                    if (30 > 0) {
                        v26 := GenS3{F0: ("" + "a")}
                        _ = v26
                    } else {
                        v27 := GenS3{F0: ("" + "")}
                        if ((v3 * (v3 & v3)) < v3) {
                            if v17 {
                                var v28 uint8 = 0
                                var v29 uint8 = v28
                                _ = v28
                                _ = v29
                            }
                            switch fi25 {
                            case 1:
                                v16 -= v1
                            case 32768:
                                v30 := GenS4{F0: -847, F1: 4294967295}
                                _ = v30
                            }
                        } else {
                            v31 := GenS5{F0: 790, F1: -128}
                            v32 := GenS3{F0: ""}
                            _ = v31
                            _ = v32
                        }
                        v33 := GenS4{F0: 42, F1: 42}
                        _ = v27
                        _ = v33
                    }
                    v17 = true
                }
            }
        }
        _ = fi25
    }
    var v34 int32 = 504
    var v35 int8 = 49
    if (int64(v1) >= genFn0(v3)) {
        v16 = 255
        switch "a" {
        case "foo bar":
            _ = len(v2)
        default:
            v36 := GenS4{F0: v34, F1: 18446744073709551615}
            _ = v36
        }
        v4 -= (v4 - v4)
    }
    v4 = (float64(0.0) + 3.14159)
    var v37 uint = 116
    v38 := GenS3{F0: ("\n\t" + ("\n\t" + ("a" + "\n\t")))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
