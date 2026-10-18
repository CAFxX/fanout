package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 int64 }
type GenS3 struct { F0 uint; F1 int8; F2 int16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint32, p1 string, p2 uint8) uint8 {
    v6 := []int{-773, -1, -2147483648}
    p0 -= (p0 << ((65535) & 63))
    var v7 uint16 = 32767
    p0 = (p0 + (p0 / ((p0) | 1)))
    p2 -= (p2 >> ((841) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 uint, p1 int64) int64 {
    if false {
        fmt.Println(p0, p0)
    } else {
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v8 float64 = (float64((1.5 * (1e10 + 0.1))) * float64((((0.0 - 0.1) / 0.1) / 1e10)))
            _ = v8
        }
        var v9 uint64 = 714
        switch "" {
        case "a":
            var v10 string = "hello"
            _ = v10
        }
        _ = v9
    }
    var v11 uint32 = 738
    p1 = -9223372036854775808
    var v12 float64 = (0.1 / -2.25)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return (^(p1))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 2
    var v3 int64 = -9223372036854775808
    var v4 int32 = 0
    var v5 string = (("foo bar" + "foo bar") + "foo bar")
    v3 -= 42
    fmt.Println(v4)
    var v13 int64 = (v3 | (v3 & v3))
    v2 = v2
    fmt.Println(v13)
    v2 += v2
    v13 -= 351
    var v14 int8 = 100
    if (false && (v14 == v14)) {
        switch 128 {
        case 304:
            if (-2147483648 <= 1) {
                v1[2147483648] = 361
                v4 = (v4 * (v4 / ((v4) | 1)))
            } else {
                for fi15 := 0; fi15 < 15; fi15++ {
                    if true {
                        v16 := []int{}
                        _ = v16
                    } else {
                        switch (v5 + v5) {
                        case "hello":
                            _ = len(v1)
                        case "foo bar":
                            v14 = v14
                        case "a":
                            v17 := GenS1{F0: (0.0 < -2.25)}
                            _ = v17
                        default:
                            if true {
                                switch fi15 {
                                case 32768:
                                    v18 := GenS2{F0: (v2 + (^(v2))), F1: v14, F2: 1}
                                    _ = v18
                                case 32767:
                                    var v19 int32 = 0
                                    _ = v19
                                case 122:
                                    v20 := GenS1{F0: (false && (947 != 779))}
                                    _ = v20
                                default:
                                    var v21 uint16 = 255
                                    _ = v21
                                }
                            } else {
                                var v22 int32 = (v4 << ((297) & 63))
                                delete(v1, 4294967295)
                                v23 := v1[0]
                                _ = v22
                                _ = v23
                            }
                        }
                    }
                    v24 := GenS3{F0: 0, F1: -1, F2: 1}
                    _ = fi15
                    _ = v24
                }
                v5 = ""
            }
        default:
            v14 = 0
        }
        v13 -= v13
        var v25 string = ("" + "hello")
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        v27 := GenS4{F0: 0}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 uint = 867
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
