package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 bool, p1 int32) uint16 {
    fmt.Println(p0)
    var v6 float64 = (0.1 / -2.25)
    v7 := []int{}
    v6 = (v6 - 1e10)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int32) int32 {
    var v8 uint = 42
    p0 += (p0 ^ p0)
    v9 := GenS1{F0: true, F1: 66}
    _ = p0
    _ = v8
    _ = v9
    return (p0 % ((p0) | 1))
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 1
    var v3 uint16 = 32767
    v4 := []int{-2147483648, 0, -927}
    var v5 int8 = 100
    if len(v4) > 0 {
        v4[0] = 32767
    }
    var v10 int64 = 659
    var v11 int64 = (v10 << ((255) & 63))
    if false {
        var v12 uint8 = 1
        fmt.Println(v2)
        v4 = append(v4, 9223372036854775807)
        _ = v12
    } else {
        fmt.Println(v2, v10, v10)
        if (!(-453 < 2147483647)) {
            v4 = append(v4, 251)
            v1 -= 0.0
        }
    }
    var v13 float64 = 3.14159
    switch "a" {
    case "foo bar":
        v1 += (3.14159 - v1)
    }
    v4 = append(v4, 42)
    v11 = v11
    for fi14 := 0; fi14 < 8; fi14++ {
        v1 = 0.1
        _ = fi14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(v10)
        var v16 int64 = v10
        switch ("foo bar" + "foo bar") {
        case "":
            v17 := GenS2{F0: 1}
            _ = v17
        case "foo bar":
            v18 := GenS3{F0: (^(v5))}
            _ = v18
        default:
            switch fi15 {
            case -590:
                _ = len(v4)
            case -665:
                if false {
                    var v19 int16 = -32768
                    if true {
                        v2 = (v2 / ((v2) | 1))
                    } else {
                        if len(v4) > 0 {
                            v4[0] = 625
                        }
                        for fi20 := 0; fi20 < 9; fi20++ {
                            if len(v4) > 0 {
                                v4[0] = (fi15 >> ((1) & 63))
                            }
                            v4 = append(v4, (fi15 | 9223372036854775807))
                            v10 -= v11
                            _ = fi20
                        }
                        if ((v2 / ((v2) | 1)) == v2) {
                            if false {
                                if (int16(fi15) < (v19 & 1)) {
                                    v11 -= (v11 & -9223372036854775808)
                                    var v21 uint64 = (v2 ^ (v2 % ((v2) | 1)))
                                    if false {
                                        for fi22 := 0; fi22 < 8; fi22++ {
                                            v1 = v1
                                            if ((v10 > (v16 + (v10 | -1))) || true) {
                                                v10 += (-(v11))
                                            } else {
                                                var v23 int64 = v11
                                                v24 := GenS1{F0: (!true), F1: 98}
                                                v19 += (v19 * (v19 - v19))
                                                if len(v4) > 0 {
                                                    v4[0] = fi15
                                                }
                                                if true {
                                                    v10 -= (-(v23))
                                                    fmt.Println(fi15, v10)
                                                }
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = fi22
                                        }
                                        for fi25 := 0; fi25 < 12; fi25++ {
                                            var v26 int8 = -1
                                            _ = fi25
                                            _ = v26
                                        }
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v21 = v21
                                    }
                                    _ = v21
                                }
                                v3 = 42
                                var v27 int16 = 1
                                _ = v27
                            } else {
                                v5 -= (v5 << ((579) & 63))
                                v16 = 0
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    if true {
                                        v13 = v13
                                        if (v19 < 360) {
                                            var v29 float32 = v1
                                            v19 = (v19 / (((v19 << ((42) & 63))) | 1))
                                            _ = v29
                                        }
                                    } else {
                                        var v30 bool = false
                                        _ = v30
                                    }
                                    _ = len(v4)
                                    var v31 int64 = (v10 ^ (-(v10)))
                                    _ = fi28
                                    _ = v31
                                }
                            }
                        } else {
                            v5 = (v5 & v5)
                            v2 -= 1
                        }
                    }
                    _ = v19
                }
            default:
                v32 := GenS1{F0: true, F1: 173}
                _ = v32
            }
        }
        _ = fi15
        _ = v16
    }
    v2 -= (v2 | v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v33 int32 = 2147483647
    var v34 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
