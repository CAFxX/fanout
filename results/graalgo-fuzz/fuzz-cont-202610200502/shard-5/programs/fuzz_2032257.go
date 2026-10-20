package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int16 }
type GenS2 struct { F0 uint32; F1 string; F2 uint8 }
type GenS3 struct { F0 uint }
func genFn0(p0 bool) int32 {
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := []int{127, fi5, fi5, (fi5 / (((fi5 * fi5)) | 1))}
        var v7 bool = (p0 || p0)
        fmt.Println(p0)
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = p0
    _ = p0
    return 0
}
func genFn1() uint {
    var v8 int = 1
    var v9 int32 = -854
    if (42 >= 2) {
        v9 = v9
        var v10 float32 = ((1e10 - (1e10 / float32(((0.0 + (-2.25 - (((0.0 + -2.25) / (3.14159 * ((float64(0.1) + ((0.0 * 1.5) / 3.14159)) * 1.5))) - 3.14159))) / 1.5)))) - 0.1)
        var v11 uint8 = 255
        _ = v10
        _ = v11
    }
    var v12 int8 = 24
    _ = v12
    _ = v8
    _ = v9
    return 974
}

func main() {
    var v1 uint32 = 42
    var v2 bool = true
    var v3 float64 = float64((-2.25 * (0.0 - ((0.1 / 1.0) / 1.0))))
    v4 := make(map[int]bool)
    v13 := v4[-837]
    v1 -= 65535
    v14 := []int{0}
    switch "\n\t" {
    case "\n\t":
        v2 = true
    case "a":
        if v13 {
            if len(v14) > 0 {
                v14[0] = -2147483648
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                switch "\n\t" {
                case "hello":
                    if (402 >= 32767) {
                        v16 := GenS1{F0: (v1 ^ v1), F1: 117, F2: -1}
                        var v17 string = ("hello" + "\n\t")
                        v18 := GenS1{F0: 4294967295, F1: -99, F2: -32768}
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        if true {
                            if v2 {
                                fmt.Println(v2)
                            } else {
                                v19 := GenS1{F0: 2147483648, F1: 114, F2: 42}
                                _ = v19
                            }
                        } else {
                            var v20 uint16 = 314
                            var v21 int16 = -1
                            _ = v20
                            _ = v21
                        }
                        v2 = v13
                    }
                default:
                    var v22 uint16 = 2
                    _ = v22
                }
                v3 += 0.0
                var v23 int32 = -2147483648
                _ = fi15
                _ = v23
            }
            if (true || (v1 <= 65535)) {
                var v24 int8 = -72
                _ = v24
            } else {
                switch "foo bar" {
                case "a":
                    var v25 string = "foo bar"
                    _ = v25
                }
                v2 = v13
                for fi26 := 0; fi26 < 16; fi26++ {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v3)
                        }
                        var v28 int32 = 42
                        _ = fi27
                        _ = v28
                    }
                    _ = fi26
                }
            }
        }
    default:
        v1 = v1
    }
    var v29 int64 = 42
    v30 := GenS2{F0: v1, F1: (("a" + ("hello" + ("" + "foo bar"))) + ""), F2: 255}
    var v31 string = ("hello" + ("" + "\n\t"))
    if false {
        var v32 uint16 = 801
        _ = v32
    } else {
        var v33 int = 310
        var v34 uint64 = 703
        v35 := GenS3{F0: 876}
        _ = v33
        _ = v34
        _ = v35
    }
    var v36 int = 9223372036854775807
    v36 = v36
    fmt.Println(v31, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v4
}
