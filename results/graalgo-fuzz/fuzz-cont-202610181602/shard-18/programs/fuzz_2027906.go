package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int32 = 42
    var v2 int32 = (v1 >> ((0) & 63))
    var v3 int8 = 0
    v4 := make(map[int]uint32)
    var v5 bool = false
    v6 := []int{0, 32768, 127, 385}
    var v7 string = "foo bar"
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    if true {
        switch ("" + "hello") {
        case "hello":
            fmt.Println(v3, v5)
        case "a":
            for fi8 := 0; fi8 < 15; fi8++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = len(v4)
                _ = fi8
            }
        default:
            delete(v4, 375)
        }
    }
    var v9 int8 = -72
    if (!v5) {
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 uint64 = 388
                var v13 uint = 662
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(v3, v3)
            fmt.Println(v9, v5)
            _ = fi10
        }
        v5 = false
    } else {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int16 = 32767
            _ = fi14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            if ((v9 % ((v3) | 1)) <= v3) {
                v2 = v1
            } else {
                if v5 {
                    var v18 uint8 = 42
                    var v19 uint8 = v18
                    fmt.Println(fi16, v5, v7)
                    _ = v18
                    _ = v19
                }
                if false {
                    if (1 == 42) {
                        if (!v5) {
                            v4[(fi16 | fi16)] = 859
                            v20 := v4[fi16]
                            _ = v20
                        }
                        if (float64(3.14159) >= 3.14159) {
                            var v21 bool = true
                            _ = v21
                        } else {
                            v22 := GenS1{F0: fi16, F1: 1000}
                            _ = v22
                        }
                    }
                } else {
                    v7 = ("foo bar" + "a")
                }
                v23 := GenS1{F0: (^(fi16)), F1: 282}
                _ = v23
            }
            v24 := GenS1{F0: 435, F1: 42}
            _ = fi17
            _ = v24
        }
        _ = fi16
    }
    var v25 int32 = 473
    if len(v6) > 0 {
        v6[0] = 404
    }
    if v5 {
        if v5 {
            if (false || false) {
                if false {
                    var v26 float32 = ((1.5 * (float32(0.1) + 0.0)) - float32((1e10 - (1e10 + float64(float32((1.5 - 3.14159)))))))
                    var v27 int8 = (-(v9))
                    var v28 bool = ((0 != 764) || v5)
                    _ = v26
                    _ = v27
                    _ = v28
                }
            }
            v7 = "a"
            v29 := GenS2{F0: (!true)}
            _ = v29
        } else {
            if v5 {
                if (128 == 32767) {
                    switch 4294967295 {
                    case -301:
                        if v5 {
                            fmt.Println(v2, v25)
                        }
                    default:
                        switch 32767 {
                        case -2147483648:
                            v2 = -214
                        case 32768:
                            fmt.Println(v9, v3)
                        case 128:
                            _ = len(v6)
                        default:
                            v5 = ((1 >= (v2 | 1)) && v5)
                        }
                    }
                }
            } else {
                var v30 float64 = -2.25
                v1 = -362
                var v31 float32 = (1.5 / (float32(v30) + 0.0))
                _ = v30
                _ = v31
            }
        }
    }
    switch v7 {
    case "\n\t":
        v32 := GenS1{F0: 32767, F1: -1}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
