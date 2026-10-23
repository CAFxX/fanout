package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0() string {
    v7 := []int{-784, -9223372036854775808, 333, -114}
    var v8 uint = 255
    v8 -= v8
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 int, p1 int8, p2 int16) string {
    if false {
        fmt.Println(p1, p2, p1)
        fmt.Println(p2)
        var v9 int16 = 32767
        _ = v9
    }
    var v10 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "a"
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (float64(((0.0 - 1.5) - ((0.0 * 0.1) * 3.14159))) / ((0.0 * 3.14159) + 3.14159))
    v3 := make(map[string]uint32)
    var v4 float32 = ((1e10 / 1.0) - 1.5)
    v5 := []string{"\n\t", "\n\t", ""}
    var v6 float64 = v2
    if (false || true) {
        var v11 string = ""
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint16 = 32767
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        v14 := GenS1{F0: -2147483648, F1: 42, F2: -9223372036854775808}
        _ = v14
    }
    v6 = v2
    if (2 == 0) {
        v2 = 0.1
        var v15 float32 = v4
        if ((false && false) || true) {
            if (!(-32768 >= -1)) {
                if (!true) {
                    var v16 int = 32767
                    _ = v16
                } else {
                    v15 += ((v4 * ((((v15 - v4) * float32(v6)) + v15) / float32(((v6 - (0.0 - v2)) + float64((v15 / v15)))))) * (v4 + ((v4 + ((float32(3.14159) + ((0.1 - (v15 * v15)) / 3.14159)) / float32(((0.0 + 0.0) - ((float64(v15) - (v6 + v6)) + v2))))) + v15)))
                    if (false && true) {
                        var v17 int32 = -2147483648
                        for fi18 := 0; fi18 < 10; fi18++ {
                            var v19 int16 = 1000
                            fmt.Println(v17, v19, v6)
                            var v20 float64 = 1e10
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        v15 = v4
                        _ = v17
                    } else {
                        switch 127 {
                        case 32768:
                            var v21 uint64 = 4294967295
                            _ = v21
                        case 128:
                            switch 4294967295 {
                            case -1:
                                fmt.Println(v2)
                            }
                        case 2147483647:
                            var v22 uint32 = 246
                            _ = v22
                        default:
                            fmt.Println(v15, v15)
                        }
                        if true {
                            var v23 string = ("hello" + "foo bar")
                            var v24 uint64 = 4294967295
                            v25 := GenS2{F0: -128, F1: 2147483648, F2: -1}
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            v4 = -2.25
                        }
                        fmt.Println(v2)
                    }
                }
                v15 = v15
            } else {
                if true {
                    var v26 float32 = (float32(v2) * ((v4 - 3.14159) + v4))
                    v4 = v4
                    _ = v26
                }
            }
            if ("\n\t" >= "a") {
                var v27 float64 = float64(v15)
                _ = v27
            }
        }
        _ = v15
    } else {
        if false {
            v28 := GenS2{F0: 106, F1: 4294967295, F2: 41}
            _ = v28
        }
        v29 := GenS2{F0: 115, F1: 848, F2: 100}
        _ = v29
    }
    fmt.Println(v2, v2, v6)
    v6 -= ((v6 - (float64(float32((-2.25 / (((0.1 * (v2 * 1.5)) - 0.0) - (((float64(1.5) + float64(float32(((1.5 + (-2.25 - (v6 / (v2 / (((-2.25 * (((1e10 * (0.1 / (3.14159 / -2.25))) - (v6 - float64(v4))) * v2)) + 0.0) + ((1.5 + v6) + ((float64(v4) - (1.5 + 3.14159)) - (v6 + float64(((0.0 / v4) * 3.14159)))))))))) + v2)))) * 0.1) + ((0.1 + v6) + -2.25)))))) / (1.5 - (float64(v4) - (v6 + v6))))) * ((0.0 / v6) + ((1e10 - (v2 + v2)) * 0.1)))
    v30 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: -284}
    var v31 int32 = 1
    var v32 string = "a"
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int32 = v31
        _ = fi33
        _ = v34
    }
    var v35 uint64 = 4294967295
    v36 := GenS3{F0: (255 < 128), F1: (-(v31))}
    fmt.Println(v31, v32, v35)
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
