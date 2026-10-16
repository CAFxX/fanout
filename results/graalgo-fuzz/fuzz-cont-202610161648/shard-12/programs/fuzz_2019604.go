package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint; F1 int; F2 uint }

func main() {
    var v1 uint = 611
    v2 := make(map[int]int)
    var v3 uint8 = 128
    var v4 uint64 = 2
    var v5 int32 = -647
    var v6 uint64 = uint64(v1)
    v3 = 2
    for fi7 := 0; fi7 < 7; fi7++ {
        switch "foo bar" {
        case "\n\t":
            v4 = (^(v4))
        case "hello":
            v8 := GenS1{F0: 561, F1: v5}
            _ = v8
        case "":
            var v9 int16 = 42
            _ = v9
        default:
            fmt.Println(fi7)
        }
        _ = fi7
    }
    var v10 bool = true
    var v11 uint32 = 1
    v12 := []int{-9223372036854775808}
    v3 = 33
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: 2, F1: v5}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 int = (fi13 - fi13)
            if (!false) {
                switch fi13 {
                case 607:
                    if len(v12) > 0 {
                        v12[0] = v17
                    }
                }
            } else {
                switch (v17 - fi16) {
                case 32767:
                    fmt.Println(v11, fi13, v10)
                case 4294967295:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                case 42:
                    v3 = v3
                default:
                    fmt.Println(v1)
                }
                if v10 {
                    switch "" {
                    case "":
                        var v18 string = "hello"
                        _ = v18
                    }
                } else {
                    var v19 int32 = 0
                    if v10 {
                        if len(v12) > 0 {
                            v12[0] = 32767
                        }
                    } else {
                        v10 = v10
                        v11 = 0
                        delete(v2, (fi16 << ((1) & 63)))
                    }
                    var v20 float64 = (0.0 - (0.0 * (1e10 / (1.5 + float64((float32((1e10 * 0.0)) / ((-2.25 - (((3.14159 + 0.1) + 1.5) * (-2.25 / 1e10))) / 1.0)))))))
                    var v21 uint = v1
                    if ("\n\t" < "hello") {
                        var v22 float64 = float64(((1e10 + 0.0) - float32(0.0)))
                        if false {
                            if (-21 != -128) {
                                v6 = 4294967295
                                switch (("foo bar" + "\n\t") + "\n\t") {
                                case "hello":
                                    fmt.Println(v3, v4)
                                default:
                                    _ = len(v12)
                                }
                                var v23 uint = uint(v6)
                                _ = v23
                            }
                            fmt.Println(v20)
                            v10 = (!true)
                        }
                        _ = v22
                    } else {
                        var v24 float64 = ((v20 - v20) - v20)
                        delete(v2, 9223372036854775807)
                        var v25 uint8 = v3
                        _ = v24
                        _ = v25
                    }
                    v4 -= v6
                    v3 = 127
                    fmt.Println(v5)
                    var v26 uint64 = (^(v6))
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v26
                }
                v6 = (v4 | (v4 + 4294967295))
            }
            _ = fi16
            _ = v17
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi13
    }
    if (v10 || v10) {
        var v27 uint16 = 1
        v28 := GenS2{F0: 4294967295, F1: -88, F2: v1}
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        var v30 float64 = 1.5
        if (false && v10) {
            if v10 {
                var v31 string = "a"
                if v10 {
                    var v32 uint8 = (v3 << ((uint(v11)) & 63))
                    _ = v32
                } else {
                    v33 := v2[fi29]
                    _ = v33
                }
                _ = v31
            } else {
                fmt.Println(v5, v1, v10)
                v12 = append(v12, (fi29 ^ (fi29 % (((-(fi29))) | 1))))
            }
            var v34 float32 = -2.25
            for fi35 := 0; fi35 < 4; fi35++ {
                if v10 {
                    v11 = (v11 | (v11 % ((v11) | 1)))
                    var v36 uint8 = uint8(v6)
                    _ = v36
                }
                _ = fi35
            }
            _ = v34
        }
        _ = fi29
        _ = v30
    }
    var v37 bool = (205 < 0)
    v11 = (v11 + v11)
    v10 = v10
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
