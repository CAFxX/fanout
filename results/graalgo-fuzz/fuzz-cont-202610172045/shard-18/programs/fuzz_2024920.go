package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }

func main() {
    var v1 int64 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    v4 := []int{-590, -249, 32768}
    fmt.Println(v1)
    v5 := GenS1{F0: 870, F1: 2147483647}
    switch 112 {
    case 703:
        for fi6 := 0; fi6 < 11; fi6++ {
            if (!(false || false)) {
                v2[""] = 65535
                if true {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        if false {
                            v1 = (-(v1))
                            var v8 int16 = 0
                            _ = v8
                        }
                        var v9 int32 = -341
                        var v10 string = ("a" + (((("a" + ("foo bar" + "a")) + "\n\t") + "a") + "\n\t"))
                        _ = fi7
                        _ = v10
                        _ = v9
                    }
                }
                for fi11 := 0; fi11 < 6; fi11++ {
                    if (((float32(1.5) / float32(0.1)) - (1e10 * (0.0 * (1e10 * ((1e10 * (0.0 - float32(-2.25))) + 0.0))))) == 0.1) {
                        var v12 int16 = -783
                        _ = v12
                    }
                    _ = fi11
                }
            } else {
                if (-246 != (v1 - v1)) {
                    var v13 int16 = -628
                    if true {
                        v13 = (v13 + v13)
                        var v14 int8 = -60
                        var v15 uint16 = 255
                        _ = v14
                        _ = v15
                    }
                    if true {
                        var v16 float32 = (((-2.25 / 1.0) + float32(3.14159)) + 0.1)
                        var v17 int16 = -32768
                        _ = v16
                        _ = v17
                    }
                    _ = v13
                } else {
                    var v18 bool = false
                    v19 := GenS2{F0: 0, F1: "\n\t", F2: 65535}
                    for fi20 := 0; fi20 < 20; fi20++ {
                        if (-128 != 42) {
                            var v21 int8 = 127
                            var v22 int = fi20
                            _ = v21
                            _ = v22
                        }
                        switch fi6 {
                        case -2147483648:
                            if (!v18) {
                                fmt.Println(v18)
                                v4 = append(v4, (fi20 / ((32768) | 1)))
                                var v23 uint8 = 92
                                v24 := GenS2{F0: 1, F1: "foo bar", F2: 65535}
                                v18 = ((!(274 > 353)) && (!false))
                                _ = len(v2)
                                v1 += v1
                                v4 = append(v4, (-(fi20)))
                                v23 = 1
                                var v25 int32 = 4
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        case 32767:
                            var v26 uint64 = 2
                            _ = v26
                        case 32768:
                            var v27 float64 = float64(((-2.25 / 1.0) / 3.14159))
                            _ = v27
                        }
                        if (!false) {
                            v1 = (v1 ^ -9223372036854775808)
                            v1 = v1
                        } else {
                            var v28 uint = 29
                            _ = v28
                        }
                        var v29 int = fi6
                        _ = len(v2)
                        _ = fi20
                        _ = v29
                    }
                    _ = v18
                    _ = v19
                }
                for fi30 := 0; fi30 < 2; fi30++ {
                    var v31 float32 = float32(1.5)
                    _ = len(v4)
                    _ = fi30
                    _ = v31
                }
            }
            v32 := GenS3{F0: false, F1: -128, F2: v1}
            _ = fi6
            _ = v32
        }
    default:
        var v33 int32 = 42
        _ = v33
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        v3[2] = false
        var v35 float64 = 3.14159
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi34
        _ = v35
    }
    var v36 int8 = 100
    fmt.Println(v1, v36)
    var v37 float64 = -2.25
    var v38 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
