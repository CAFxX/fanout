package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 string; F1 int }
func genFn0(p0 string) int16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint32 = 42
        v8 -= 0
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    _ = p0
    return 32767
}

func main() {
    var v1 int8 = 42
    var v2 uint = 42
    var v3 bool = true
    var v4 bool = v3
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]string)
    v1 += v1
    v6[1] = ""
    var v11 float32 = ((-2.25 - 3.14159) - 0.0)
    var v12 float64 = float64(v11)
    var v13 uint32 = 42
    v14 := []int{-47, -9223372036854775808}
    var v15 uint32 = (v13 % ((752) | 1))
    fmt.Println(v3, v2)
    switch 4294967295 {
    case 128:
        fmt.Println(v1, v2, v11)
    case 32768:
        if (9223372036854775807 >= -1) {
            v16 := GenS1{F0: 915, F1: -841}
            fmt.Println(v1, v11, v13)
            if v3 {
                for fi17 := 0; fi17 < 5; fi17++ {
                    if false {
                        var v18 string = ""
                        v4 = (!(255 >= (^(v2))))
                        _ = v18
                    } else {
                        switch "hello" {
                        case "hello":
                            var v19 string = ("foo bar" + "foo bar")
                            _ = v19
                        case "a":
                            if (1.5 > v11) {
                                var v20 int64 = 1
                                for fi21 := 0; fi21 < 14; fi21++ {
                                    v14 = append(v14, fi21)
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    _ = fi21
                                }
                                _ = v20
                            }
                        default:
                            for fi22 := 0; fi22 < 19; fi22++ {
                                var v23 int16 = -1
                                v6[(v15 / ((uint32(v2)) | 1))] = ("foo bar" + "hello")
                                v12 += (v12 + v12)
                                _ = fi22
                                _ = v23
                            }
                        }
                    }
                    var v24 bool = (v3 || v3)
                    fmt.Println(v2)
                    _ = fi17
                    _ = v24
                }
            } else {
                v25 := GenS1{F0: 42, F1: -8}
                if (!false) {
                    switch "\n\t" {
                    case "\n\t":
                        v26 := GenS2{F0: ("foo bar" + "hello"), F1: -2147483648}
                        _ = v26
                    default:
                        for fi27 := 0; fi27 < 13; fi27++ {
                            v14 = append(v14, fi27)
                            _ = fi27
                        }
                    }
                } else {
                    v28 := GenS1{F0: 268, F1: -901}
                    _ = v28
                }
                if (v12 < -2.25) {
                    if (!(-76 < -9223372036854775808)) {
                        fmt.Println(v4)
                        fmt.Println(v12, v3)
                        var v29 int32 = -165
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 11; fi30++ {
                        var v31 int32 = 1
                        if ((v13 & v13) == (^(v13))) {
                            var v32 float32 = float32(1e10)
                            _ = v32
                        }
                        switch (fi30 / ((fi30) | 1)) {
                        case 952:
                            for fi33 := 0; fi33 < 15; fi33++ {
                                if v3 {
                                    fmt.Println(fi30, v31)
                                } else {
                                    if ((v13 - v15) == v15) {
                                        if v4 {
                                            var v34 string = "hello"
                                            v13 -= (v13 - 0)
                                            _ = v34
                                        }
                                        var v35 float32 = (3.14159 * v11)
                                        _ = v35
                                    }
                                }
                                v11 -= v11
                                _ = fi33
                            }
                        default:
                            var v36 int32 = 703
                            _ = v36
                        }
                        _ = fi30
                        _ = v31
                    }
                }
                _ = v25
            }
            _ = v16
        } else {
            v37 := GenS1{F0: 611, F1: -758}
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        _ = len(v5)
        v13 -= (v13 ^ 813)
        v1 = v1
        _ = fi38
    }
    var v39 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
