package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 int8 }
func genFn0() uint8 {
    var v5 int = -904
    var v6 int8 = -20
    v7 := GenS1{F0: 675, F1: 4294967295, F2: 507}
    v8 := GenS1{F0: 65535, F1: 4294967295, F2: 0}
    var v9 float32 = 0.0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 2
    var v2 uint8 = (v1 << ((0) & 63))
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    v10 := []int{-349, 128, 997, int(v1)}
    switch 42 {
    case 1:
        var v11 int = -2147483648
        _ = v11
    case -1:
        fmt.Println(v2)
    case -721:
        v1 += 42
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if false {
            var v13 float64 = ((-2.25 * ((((-2.25 * ((0.1 / 1e10) / -2.25)) - 1.5) + 0.1) + (((((1e10 / 1e10) / float64(3.14159)) + (3.14159 - 3.14159)) + 3.14159) + (float64(0.0) / 0.1)))) - float64(float32(3.14159)))
            var v14 bool = false
            _ = v13
            _ = v14
        } else {
            if false {
                var v15 bool = ((false || false) || true)
                v15 = ("foo bar" >= "")
                _ = v15
            }
        }
        if true {
            v1 += (v2 + genFn0())
            if (((("" + "hello") + (("" + "a") + ("foo bar" + "foo bar"))) + "") < "foo bar") {
                v16 := GenS1{F0: 1, F1: 2, F2: 42}
                _ = v16
            }
            switch -4 {
            case 1:
                v1 = v1
            case -1:
                fmt.Println(v1, fi12, v2)
            case 32768:
                v1 = (v2 * v1)
            default:
                v1 -= 255
            }
        } else {
            var v17 int32 = -577
            _ = v17
        }
        var v18 float64 = float64((0.0 - (0.0 + (0.1 * 3.14159))))
        _ = fi12
        _ = v18
    }
    _ = len(v3)
    if false {
        v2 = (v2 + genFn0())
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            switch "a" {
            case "a":
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 int = 975
                    _ = fi20
                    _ = v21
                }
            }
            v22 := GenS2{F0: 32767, F1: v2, F2: (32767 == -7)}
            _ = fi19
            _ = v22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            delete(v4, "a")
            v1 += genFn0()
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi23
        }
    }
    v1 = (v1 % ((v1) | 1))
    v2 -= v1
    for fi24 := 0; fi24 < 1; fi24++ {
        if false {
            v2 -= 79
            for fi25 := 0; fi25 < 6; fi25++ {
                if (false && true) {
                    delete(v4, "")
                    v26 := GenS3{F0: fi25, F1: 101}
                    fmt.Println(fi24, v2, v2)
                    if (-1 <= -1) {
                        var v27 uint = 0
                        _ = v27
                    }
                    fmt.Println(v1, fi25)
                    _ = v26
                }
                v2 = genFn0()
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi25
            }
            var v28 int16 = 42
            _ = v28
        }
        v1 = v2
        _ = fi24
    }
    v29 := GenS4{F0: -1}
    for fi30 := 0; fi30 < 12; fi30++ {
        v31 := GenS2{F0: 42, F1: 255, F2: true}
        fmt.Println(v2, v2, v2)
        if (!true) {
            fmt.Println(v2)
        } else {
            v32 := GenS5{F0: 255, F1: -98}
            _ = v32
        }
        _ = fi30
        _ = v31
    }
    if (4294967295 == 42) {
        v3[1] = ((736 <= 2147483648) && true)
    }
    var v33 bool = false
    switch "" {
    case "\n\t":
        if v33 {
            v1 -= v2
            if v33 {
                fmt.Println(v1, v33, v2)
            } else {
                switch "\n\t" {
                case "a":
                    v2 += (v1 & (v1 - (^(v2))))
                case "\n\t":
                    if v33 {
                        v1 = v2
                        if true {
                            for fi34 := 0; fi34 < 5; fi34++ {
                                var v35 float32 = (-2.25 - 0.0)
                                _ = fi34
                                _ = v35
                            }
                            fmt.Println(v33, v1, v33)
                        } else {
                            v36 := GenS4{F0: -1}
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 12; fi37++ {
                            v2 -= v2
                            if false {
                                _ = len(v10)
                                for fi38 := 0; fi38 < 3; fi38++ {
                                    var v39 int64 = 42
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            _ = fi37
                        }
                    }
                }
                v10 = append(v10, 926)
            }
        } else {
            var v40 uint32 = 2147483647
            _ = v40
        }
    default:
        v33 = true
    }
    v2 += 186
    v1 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
}
