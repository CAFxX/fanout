package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
func genFn0(p0 bool, p1 bool) uint16 {
    if (true || true) {
        p0 = false
        fmt.Println(p1, p0)
    }
    if (false || (1 <= 9223372036854775807)) {
        var v6 float32 = 3.14159
        _ = v6
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int16, p1 uint, p2 int32) int32 {
    fmt.Println(p2, p2, p1)
    var v7 float32 = 3.14159
    switch -9223372036854775808 {
    case -9223372036854775808:
        v8 := []int{4294967295, int(p2)}
        _ = v8
    default:
        var v9 int = 572
        _ = v9
    }
    v10 := []int{32768, 127, 4294967295}
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return int32(p1)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    var v3 float64 = (0.0 * -2.25)
    var v4 bool = (255 >= 255)
    var v5 float64 = v3
    fmt.Println(v5)
    v11 := GenS1{F0: 18446744073709551615, F1: -1, F2: 255}
    var v12 float32 = 1.5
    v12 += float32(1.5)
    v3 += v3
    v3 = v5
    v2 = false
    if v2 {
        if v4 {
            switch 0 {
            case 596:
                v13 := []int{-845, -799, 4294967295}
                _ = v13
            default:
                v5 = v3
            }
            v12 = 1.5
        }
        if (2 == 31) {
            v3 = (v3 / (v5 / 1.0))
            var v14 uint32 = 4294967295
            v5 -= v3
            _ = v14
        }
    } else {
        v3 = v5
        fmt.Println(v3, v2, v3)
        var v15 float64 = v3
        _ = v15
    }
    v3 -= ((v3 + (float64((v12 * -2.25)) / 0.1)) / v3)
    v5 = -2.25
    if (-2147483648 > 0) {
        if v2 {
            switch "" {
            case "":
                v12 -= 0.1
            case "hello":
                var v16 int = -1
                _ = v16
            case "\n\t":
                fmt.Println(v5, v12)
            default:
                fmt.Println(v4, v2, v4)
            }
        } else {
            v5 = 3.14159
        }
        if ((42 > 127) && (!v4)) {
            var v17 int8 = -34
            _ = v17
        }
    } else {
        v4 = (!v4)
        switch 32768 {
        case 4294967295:
            v12 = v12
        case -1:
            v18 := GenS1{F0: 18446744073709551615, F1: 42, F2: 42}
            _ = v18
        case 339:
            switch "" {
            case "hello":
                v4 = (662 <= 0)
            case "foo bar":
                v19 := []int{32768, 42, -761, 789}
                _ = v19
            case "":
                if (222 != 255) {
                    v20 := v1[1]
                    var v21 float32 = float32((1e10 * v5))
                    fmt.Println(v3, v4, v4)
                    _ = v20
                    _ = v21
                } else {
                    if v2 {
                        for fi22 := 0; fi22 < 8; fi22++ {
                            var v23 uint64 = 2
                            fmt.Println(v3, v23)
                            switch "foo bar" {
                            case "foo bar":
                                if (42 > -775) {
                                    v24 := []int{}
                                    _ = v24
                                } else {
                                    var v25 float32 = -2.25
                                    v25 += (v25 * v25)
                                    fmt.Println(v23, v23)
                                    _ = v25
                                }
                            case "\n\t":
                                switch ("\n\t" + "") {
                                case "foo bar":
                                    v3 = ((((float64(v12) * v5) * 0.0) + v5) * (v3 / v5))
                                default:
                                    if (255 > uint(v23)) {
                                        if v4 {
                                            v4 = ((fi22 - fi22) != fi22)
                                            _ = len(v1)
                                        } else {
                                            v26 := []int{9223372036854775807, (fi22 >> ((65535) & 63)), fi22}
                                            var v27 uint = 2
                                            v4 = v4
                                            _ = v26
                                            _ = v27
                                        }
                                    } else {
                                        var v28 string = "hello"
                                        _ = v28
                                    }
                                }
                            }
                            _ = fi22
                            _ = v23
                        }
                    }
                    v4 = (196 <= 203)
                    var v29 int32 = 156
                    _ = v29
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
