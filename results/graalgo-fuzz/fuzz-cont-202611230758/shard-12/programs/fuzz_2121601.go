package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int16; F1 int16; F2 int32 }
func genFn0(p0 uint8) int64 {
    var v6 string = ""
    v7 := GenS1{F0: false}
    _ = p0
    _ = v6
    _ = v7
    return -801
}

func main() {
    var v1 uint = 255
    var v2 float32 = 3.14159
    v3 := make(map[int]bool)
    v4 := []int{2147483647, -2147483648, -616, -195}
    var v5 string = "a"
    v8 := GenS1{F0: (!false)}
    var v9 uint32 = 1
    if (0 != 632) {
        if len(v4) > 0 {
            v4[0] = 1
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            v5 = v5
            v5 = v5
            var v11 int64 = 42
            _ = fi10
            _ = v11
        }
        v2 = v2
    } else {
        v12 := GenS2{F0: 1, F1: 177}
        v3[4294967295] = false
        var v13 int64 = -1
        _ = v12
        _ = v13
    }
    switch 32768 {
    case -2147483648:
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint32 = 4294967295
            fmt.Println(v15, v15, v9)
            switch (^(fi14)) {
            case 1:
                v2 += 1e10
            case -144:
                if false {
                    if (true || ((3.14159 / (v2 + (v2 * v2))) != 1.5)) {
                        if false {
                            if (0 > 956) {
                                switch "" {
                                case "foo bar":
                                    var v16 uint64 = 0
                                    _ = v16
                                case "":
                                    for fi17 := 0; fi17 < 10; fi17++ {
                                        v3[42] = true
                                        _ = fi17
                                    }
                                case "a":
                                    fmt.Println(fi14, v15)
                                default:
                                    fmt.Println(v2)
                                }
                                var v18 uint32 = (^(v15))
                                var v19 string = ""
                                _ = v18
                                _ = v19
                            }
                            var v20 int32 = 846
                            _ = v20
                        } else {
                            v1 = v1
                        }
                    }
                } else {
                    v21 := GenS2{F0: -139, F1: 42}
                    if (!false) {
                        fmt.Println(v9, v15)
                        var v22 uint32 = (v9 * v9)
                        if ((fi14 & fi14) <= -680) {
                            fmt.Println(v2)
                        } else {
                            var v23 float32 = v2
                            v2 -= 1e10
                            v24 := GenS1{F0: (1 > 485)}
                            _ = v23
                            _ = v24
                        }
                        _ = v22
                    }
                    _ = v21
                }
            case 9223372036854775807:
                switch 128 {
                case -1:
                    fmt.Println(v15)
                case 9223372036854775807:
                    if (false || false) {
                        for fi25 := 0; fi25 < 18; fi25++ {
                            v26 := GenS2{F0: 0, F1: 218}
                            fmt.Println(v1)
                            _ = fi25
                            _ = v26
                        }
                    }
                case 1:
                    fmt.Println(v15, v9)
                }
            }
            _ = fi14
            _ = v15
        }
    case 1:
        if len(v4) > 0 {
            v4[0] = -2147483648
        }
    default:
        delete(v3, 42)
    }
    v27 := GenS1{F0: true}
    switch 127 {
    case 249:
        if (65535 < 961) {
            v1 = v1
            v2 -= v2
            if (2 != 4294967295) {
                fmt.Println(v5, v5)
            } else {
                switch 193 {
                case 14:
                    var v28 uint32 = (v9 & 42)
                    _ = v28
                case 32767:
                    if ((0.1 - v2) < ((v2 / (float32(-2.25) / v2)) - float32((-2.25 - (0.1 - 0.1))))) {
                        v1 += (v1 - 18446744073709551615)
                    } else {
                        v2 -= float32((float64(0.0) - (1.5 / 3.14159)))
                    }
                case 32768:
                    var v29 float64 = (-2.25 / 3.14159)
                    _ = v29
                default:
                    var v30 uint64 = 18446744073709551615
                    _ = v30
                }
                v2 -= v2
            }
        } else {
            v5 = ("" + v5)
        }
    case 446:
        if false {
            switch -2147483648 {
            case -2147483648:
                var v31 int64 = 42
                _ = v31
            case 4294967295:
                v2 -= ((v2 - float32(0.1)) + v2)
            }
        } else {
            fmt.Println(v2, v9, v2)
            v32 := GenS3{F0: -881, F1: 1, F2: 0}
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        for fi34 := 0; fi34 < 14; fi34++ {
            delete(v3, (fi33 >> (((^(v1))) & 63)))
            _ = fi34
        }
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 int32 = 669
            v36 -= v36
            _ = fi35
            _ = v36
        }
        _ = fi33
    }
    var v37 int32 = -451
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
