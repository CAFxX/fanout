package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0(p0 int64) int {
    v7 := []int{2147483647, 32768}
    switch "foo bar" {
    case "":
        for fi8 := 0; fi8 < 18; fi8++ {
            if len(v7) > 0 {
                v7[0] = (fi8 << ((258) & 63))
            }
            v7 = append(v7, (fi8 ^ (fi8 ^ (fi8 * fi8))))
            _ = fi8
        }
    }
    var v9 int32 = 0
    switch "foo bar" {
    case "a":
        for fi10 := 0; fi10 < 20; fi10++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v11 int = fi10
            var v12 bool = true
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        if false {
            v9 = v9
        }
    }
    _ = p0
    _ = v7
    _ = v9
    return -571
}

func main() {
    var v1 int8 = -1
    v2 := []int{-9223372036854775808, 128, -908}
    var v3 int64 = 0
    var v4 bool = (42 != 1)
    var v5 string = ""
    v6 := make(map[string]string)
    fmt.Println(v5, v4)
    v4 = v4
    v13 := GenS1{F0: 100, F1: 1}
    v1 = 100
    var v14 int32 = -2147483648
    var v15 int = -38
    for fi16 := 0; fi16 < 10; fi16++ {
        if (1 != 18446744073709551615) {
            v14 += (v14 & -1)
        } else {
            v5 = "a"
        }
        _ = fi16
    }
    v14 = v14
    for fi17 := 0; fi17 < 8; fi17++ {
        if (v14 != v14) {
            for fi18 := 0; fi18 < 2; fi18++ {
                v14 = 0
                _ = fi18
            }
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            v1 = (-(v1))
            v2 = append(v2, -367)
            _ = fi19
        }
        _ = fi17
    }
    var v20 uint64 = 619
    switch genFn0(v3) {
    case 74:
        if (65535 == 42) {
            fmt.Println(v3)
            if (0.1 != 1e10) {
                if ((v20 - v20) < 4294967295) {
                    fmt.Println(v3, v5)
                    v2 = append(v2, (v15 & (v15 + genFn0((-(v3))))))
                    if true {
                        var v21 bool = (genFn0((v3 * v3)) >= v15)
                        if (-2.25 > 0.1) {
                            var v22 int8 = (v1 + (v1 + v1))
                            var v23 int8 = v22
                            delete(v6, "foo bar")
                            _ = v22
                            _ = v23
                        }
                        v15 = (v15 & 42)
                        _ = v21
                    }
                }
            } else {
                if v4 {
                    v4 = v4
                    var v24 uint64 = (^(v20))
                    _ = v24
                }
            }
            switch ("a" + v5) {
            case "\n\t":
                v14 = (v14 | v14)
            case "":
                if v4 {
                    fmt.Println(v4, v4, v14)
                    for fi25 := 0; fi25 < 9; fi25++ {
                        switch v5 {
                        case "foo bar":
                            fmt.Println(v15, v4)
                        }
                        v3 -= (v3 * (-(v3)))
                        _ = fi25
                    }
                }
            default:
                v26 := GenS1{F0: v1, F1: 0}
                _ = v26
            }
        }
    case 0:
        v27 := GenS1{F0: (v1 - v1), F1: 0}
        _ = v27
    case 421:
        v28 := GenS1{F0: (v1 / ((-1) | 1)), F1: 803}
        _ = v28
    default:
        v29 := GenS1{F0: (-(v1)), F1: 42}
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 int8 = (v1 | (^(v1)))
        _ = fi30
        _ = v31
    }
    switch "\n\t" {
    case "a":
        var v32 int16 = 1
        _ = v32
    case "\n\t":
        if (((-2.25 / ((-2.25 - 0.0) - 0.0)) + (1.5 - 0.0)) < 0.1) {
            v20 -= (v20 ^ (v20 / ((v20) | 1)))
        } else {
            v14 = (v14 >> ((40) & 63))
        }
    default:
        var v33 float64 = float64((0.0 / 1.0))
        _ = v33
    }
    v34 := GenS1{F0: int8(v3), F1: uint(v20)}
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
