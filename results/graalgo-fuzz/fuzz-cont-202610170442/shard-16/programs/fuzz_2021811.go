package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    v1 := []string{"hello"}
    v2 := make(map[uint32]int)
    var v3 uint32 = 42
    var v4 uint32 = (v3 / ((v3) | 1))
    v5 := []string{"hello", ("" + "hello"), "a"}
    v6 := GenS1{F0: -1, F1: (-1 >= -759)}
    var v7 int64 = -1
    for fi8 := 0; fi8 < 4; fi8++ {
        v3 = (^(v3))
        _ = fi8
    }
    var v9 uint8 = 183
    v4 += (v3 | (v3 ^ v4))
    v7 = (v7 ^ -1)
    var v10 uint8 = v9
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int = -2147483648
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 int64 = (v7 * (v7 ^ v7))
        switch (^(fi13)) {
        case 389:
            v7 = 42
        case 32767:
            if (!(true && false)) {
                if false {
                    for fi15 := 0; fi15 < 10; fi15++ {
                        var v16 string = (("a" + ("hello" + ("foo bar" + ("hello" + (("" + "a") + "a"))))) + ("" + ""))
                        if true {
                            var v17 uint16 = 268
                            _ = v17
                        }
                        _ = fi15
                        _ = v16
                    }
                } else {
                    v18 := GenS1{F0: 197, F1: ((131 < v14) && (1 == -985))}
                    _ = v18
                }
                for fi19 := 0; fi19 < 17; fi19++ {
                    var v20 int8 = 1
                    if (true || (-2.25 < 0.0)) {
                        var v21 int32 = -1
                        _ = v21
                    } else {
                        v10 -= (^(v9))
                        v22 := GenS2{F0: (fi13 | (fi13 | fi19)), F1: v20}
                        _ = v22
                    }
                    var v23 float32 = float32((0.0 / 1.0))
                    _ = fi19
                    _ = v20
                    _ = v23
                }
                v9 = (v9 & (^(v10)))
            }
        case 1:
            if false {
                v2[(^(v3))] = -325
            }
        }
        v10 = v9
        _ = fi13
        _ = v14
    }
    _ = len(v2)
    v7 = v7
    if true {
        v24 := GenS1{F0: -535, F1: false}
        fmt.Println(v4, v3, v3)
        var v25 int64 = v7
        _ = v24
        _ = v25
    } else {
        var v26 uint16 = 553
        _ = v26
    }
    var v27 int64 = v7
    v4 += (v4 % ((2) | 1))
    if false {
        for fi28 := 0; fi28 < 9; fi28++ {
            v1 = append(v1, "foo bar")
            if (-2.25 >= (3.14159 - (float64((1e10 / (float32((-2.25 / 0.1)) + -2.25))) / (-2.25 / -2.25)))) {
                v29 := GenS2{F0: (^(fi28)), F1: 42}
                _ = v29
            } else {
                for fi30 := 0; fi30 < 18; fi30++ {
                    switch fi30 {
                    case 32768:
                        switch ("" + "foo bar") {
                        case "\n\t":
                            v9 -= v10
                        case "hello":
                            if (false || true) {
                                v10 -= 153
                            } else {
                                fmt.Println(fi30, v4, fi28)
                            }
                        default:
                            v10 = (v10 >> ((415) & 63))
                        }
                    case 741:
                        switch "a" {
                        case "hello":
                            v9 = v9
                        }
                    case 2147483647:
                        _ = len(v2)
                    default:
                        v7 = (v7 >> ((1) & 63))
                    }
                    v3 += uint32(v10)
                    v10 = (v10 ^ 120)
                    _ = fi30
                }
                for fi31 := 0; fi31 < 2; fi31++ {
                    fmt.Println(fi31, v10)
                    var v32 bool = true
                    v32 = ((-32768 < -32768) && (uint(v4) < 1))
                    _ = fi31
                    _ = v32
                }
            }
            _ = fi28
        }
        v3 = v3
    } else {
        var v33 uint = 18446744073709551615
        fmt.Println(v7, v27)
        if true {
            for fi34 := 0; fi34 < 6; fi34++ {
                v33 = 54
                v3 += v4
                _ = fi34
            }
        }
        _ = v33
    }
    v27 -= v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
