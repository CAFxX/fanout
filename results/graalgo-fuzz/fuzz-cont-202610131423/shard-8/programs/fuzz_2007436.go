package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
func genFn0(p0 uint16, p1 uint, p2 uint64) int {
    v5 := []int{}
    var v6 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return int(p2)
}

func main() {
    var v1 uint8 = 211
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483648
    var v4 int32 = 2147483647
    var v7 uint16 = 255
    fmt.Println(v1, v3, v1)
    switch "\n\t" {
    case "foo bar":
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 float32 = 0.0
            v10 := []int{2147483647, (^(fi8))}
            _ = fi8
            _ = v10
            _ = v9
        }
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int32 = (v4 - v4)
            if (!true) {
                v3 += 2147483648
                v13 := GenS1{F0: 32767, F1: -1}
                _ = v13
            } else {
                v14 := []int{int(v7), -2147483648, 42, fi11}
                for fi15 := 0; fi15 < 18; fi15++ {
                    if ((!false) || (false || true)) {
                        if false {
                            var v16 uint8 = v1
                            v3 = (v3 % ((v3) | 1))
                            _ = v16
                        }
                        var v17 int = (fi11 & fi11)
                        v18 := v2[0]
                        _ = v17
                        _ = v18
                    } else {
                        v4 = int32(fi15)
                        _ = len(v14)
                        if (!false) {
                            switch "hello" {
                            case "a":
                                if false {
                                    var v19 bool = ((!false) && true)
                                    v12 = int32(v7)
                                    _ = v19
                                } else {
                                    fmt.Println(v4, fi11)
                                    fmt.Println(fi11)
                                    v12 += (v12 * (-(v4)))
                                }
                            }
                        }
                    }
                    _ = fi15
                }
                _ = v14
            }
            switch fi11 {
            case 24:
                var v20 bool = (fi11 == (^(fi11)))
                _ = v20
            default:
                if (false || ((1e10 / (((0.0 * (1.5 - ((1.5 * 0.1) + 1e10))) - 1e10) + 0.1)) > 1e10)) {
                    if false {
                        if ((v12 | -1) < (v12 << ((uint(v7)) & 63))) {
                            fmt.Println(fi11, v7)
                        }
                        if (215 > 42) {
                            v12 -= 764
                            var v21 uint8 = 255
                            v22 := []int{(-(fi11)), genFn0((v7 << ((0) & 63)), 255, 4294967295), genFn0(v7, 18446744073709551615, 0)}
                            _ = v21
                            _ = v22
                        } else {
                            v12 = (v4 / ((2147483647) | 1))
                        }
                        v1 += (v1 & v1)
                    } else {
                        v1 = 127
                    }
                    var v23 int64 = -577
                    _ = v23
                } else {
                    if (false && false) {
                        var v24 int32 = v12
                        _ = v24
                    } else {
                        switch -2147483648 {
                        case 32767:
                            v25 := GenS2{F0: ((-2.25 < 0.0) || true)}
                            _ = v25
                        default:
                            var v26 int8 = 100
                            _ = v26
                        }
                        switch genFn0(32767, 65535, 295) {
                        case 42:
                            fmt.Println(v3, v3, v3)
                        case 0:
                            fmt.Println(v1)
                        }
                    }
                    v2[4294967295] = (677 <= 1)
                    v12 += (v4 * -529)
                }
            }
            _ = fi11
            _ = v12
        }
    case "a":
        v27 := GenS2{F0: false}
        _ = v27
    }
    v28 := GenS3{F0: true}
    v29 := GenS2{F0: (true && (!(true && (v3 <= 2))))}
    v1 += 1
    v7 -= 308
    var v30 bool = true
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v7
}
