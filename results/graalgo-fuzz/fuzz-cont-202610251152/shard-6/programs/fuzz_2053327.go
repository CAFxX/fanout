package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 int16 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]uint32)
    var v3 float32 = float32((0.0 - (3.14159 / ((1.5 * 1e10) + (3.14159 + 3.14159)))))
    var v4 uint = 0
    var v5 int64 = -24
    switch ("hello" + "") {
    case "hello":
        v6 := []int{2147483647, 946, 127, 127}
        _ = v6
    case "":
        if true {
            for fi7 := 0; fi7 < 2; fi7++ {
                v3 += (((float32(1e10) / v3) + (1.5 / 1.0)) / float32(0.1))
                switch "a" {
                case "foo bar":
                    v1 = 13
                case "a":
                    v5 = (-(v5))
                default:
                    for fi8 := 0; fi8 < 15; fi8++ {
                        v4 -= (^(v4))
                        _ = fi8
                    }
                }
                _ = fi7
            }
            fmt.Println(v5)
            for fi9 := 0; fi9 < 4; fi9++ {
                if ((v5 < (v5 << ((4294967295) & 63))) && true) {
                    var v10 int8 = 1
                    _ = v10
                } else {
                    if (false || true) {
                        v3 += float32(0.1)
                    } else {
                        fmt.Println(v5, v5)
                        fmt.Println(v5, fi9, v3)
                    }
                    for fi11 := 0; fi11 < 13; fi11++ {
                        v5 = (-(v5))
                        if (!(!(uint16(v4) <= 736))) {
                            var v12 int64 = (v5 | 199)
                            _ = v12
                        } else {
                            if true {
                                var v13 string = "\n\t"
                                _ = v13
                            } else {
                                v3 = float32((3.14159 / 1e10))
                                var v14 float32 = v3
                                v5 = (v5 >> ((v4) & 63))
                                _ = v14
                            }
                        }
                        _ = fi11
                    }
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        for fi15 := 0; fi15 < 6; fi15++ {
                            if false {
                                v16 := []int{(fi9 ^ fi15), (fi15 | (fi9 ^ fi15)), fi9, 2147483647}
                                _ = v16
                            } else {
                                if (false && ((v5 * v5) > (v5 & v5))) {
                                    var v17 float64 = (-2.25 + float64(v3))
                                    v18 := GenS1{F0: 1, F1: ("foo bar" != ("hello" + "a"))}
                                    v1 = v1
                                    fmt.Println(fi15)
                                    _ = v17
                                    _ = v18
                                }
                                var v19 int8 = -1
                                _ = v19
                            }
                            v20 := GenS1{F0: 42, F1: (!((v3 + -2.25) == (float32(float64(((((v3 * (0.0 * float32(1e10))) / (((v3 + -2.25) - v3) + v3)) * (v3 + (v3 + -2.25))) * -2.25))) - v3)))}
                            var v21 int = 9223372036854775807
                            _ = fi15
                            _ = v20
                            _ = v21
                        }
                    default:
                        v22 := v2[256]
                        _ = v22
                    }
                }
                v4 += 4294967295
                _ = fi9
            }
        }
    case "foo bar":
        v23 := []int{42, 128}
        _ = v23
    }
    v3 = ((-2.25 / (v3 - 1e10)) - v3)
    v1 = v1
    var v24 float32 = (float32((1e10 - 1.5)) * 3.14159)
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := GenS1{F0: 815, F1: true}
        v4 -= 1
        _ = fi25
        _ = v26
    }
    if (1 < -494) {
        var v27 bool = false
        var v28 uint = v4
        _ = v27
        _ = v28
    }
    var v29 uint8 = 128
    var v30 string = ""
    v31 := GenS2{F0: 65535, F1: 827, F2: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
