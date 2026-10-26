package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
func genFn0(p0 int64) string {
    var v4 int16 = -172
    var v5 int8 = int8(p0)
    _ = p0
    _ = v4
    _ = v5
    return ("foo bar" + ("a" + "hello"))
}
func genFn1(p0 uint64, p1 bool) uint64 {
    var v6 int16 = -1
    p0 += 636
    var v7 int64 = -764
    v7 = (v7 << ((661) & 63))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 string = ""
    var v2 string = (v1 + v1)
    var v3 uint16 = 921
    if true {
        var v8 float64 = float64((1e10 / 0.1))
        _ = v8
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(v2)
            _ = fi9
        }
        v10 := GenS1{F0: false, F1: 930}
        if true {
            v11 := GenS1{F0: true, F1: 42}
            var v12 uint16 = v3
            if (!(((2147483647 == 645) && (0 <= 479)) || true)) {
                fmt.Println(v12, v1)
                v13 := GenS1{F0: false, F1: -1}
                switch 127 {
                case 128:
                    if (!(3.14159 >= ((1e10 * -2.25) + (1.5 * 0.0)))) {
                        var v14 int64 = -1
                        _ = v14
                    } else {
                        v15 := []int{-9223372036854775808}
                        var v16 uint32 = 176
                        fmt.Println(v2, v12)
                        _ = v15
                        _ = v16
                    }
                case 42:
                    fmt.Println(v1)
                default:
                    v1 = v2
                }
                _ = v13
            } else {
                if (-883 < -1) {
                    fmt.Println(v1, v3, v12)
                    if false {
                        fmt.Println(v3)
                        v3 -= 0
                        v17 := []int{-63}
                        _ = v17
                    } else {
                        v3 = v3
                    }
                    fmt.Println(v2, v2, v12)
                } else {
                    switch v2 {
                    case "hello":
                        for fi18 := 0; fi18 < 18; fi18++ {
                            var v19 uint64 = 202
                            fmt.Println(v3, v2, v12)
                            _ = fi18
                            _ = v19
                        }
                    default:
                        v2 = "a"
                    }
                    for fi20 := 0; fi20 < 6; fi20++ {
                        for fi21 := 0; fi21 < 15; fi21++ {
                            var v22 uint8 = 249
                            fmt.Println(v12, v12, v3)
                            var v23 uint16 = 0
                            if true {
                                v24 := []int{}
                                _ = v24
                            } else {
                                if (1 >= 42) {
                                    v25 := GenS2{F0: (v22 ^ v22), F1: 42, F2: -1}
                                    _ = v25
                                } else {
                                    v3 = v12
                                    if (false || ((^(v22)) < 1)) {
                                        fmt.Println(fi20, fi20, v12)
                                        fmt.Println(v2, v1)
                                    } else {
                                        if ((((float64(-2.25) + 0.1) - 1.5) * 1e10) >= (1.5 - 0.1)) {
                                            fmt.Println(v22, v22, v22)
                                            fmt.Println(v23, v2, v23)
                                            v2 = ("a" + "hello")
                                        } else {
                                            v26 := GenS1{F0: true, F1: -9223372036854775808}
                                            fmt.Println(v12, v1, v3)
                                            _ = v26
                                        }
                                        var v27 int = -1
                                        var v28 int = (^(fi20))
                                        _ = v27
                                        _ = v28
                                    }
                                }
                                var v29 int32 = 719
                                var v30 bool = false
                                _ = v29
                                _ = v30
                            }
                            v2 = (v2 + v1)
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        var v31 uint16 = 17
                        _ = fi20
                        _ = v31
                    }
                }
                v3 -= v3
            }
            _ = v11
            _ = v12
        } else {
            var v32 int32 = 42
            v33 := GenS2{F0: 118, F1: int8(v32), F2: (-(v32))}
            _ = v32
            _ = v33
        }
        _ = v10
    }
    var v34 int16 = 42
    fmt.Println(v3, v34, v2)
    v35 := []int{2147483647, 127, -929, 72}
    v36 := GenS2{F0: 2, F1: 100, F2: 2147483647}
    v34 = v34
    var v37 int = 2147483647
    v2 = (v2 + v1)
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    v38 := GenS2{F0: 255, F1: 76, F2: 953}
    if false {
        v37 += v37
    }
    var v39 int16 = (^(v34))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
