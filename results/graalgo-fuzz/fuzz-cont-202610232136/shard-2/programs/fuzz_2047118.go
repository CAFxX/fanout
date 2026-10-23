package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint; F2 uint8 }
func genFn0(p0 int64, p1 uint, p2 string) int32 {
    v5 := []int{85}
    v6 := GenS1{F0: 18446744073709551615, F1: 326}
    switch "a" {
    case "a":
        v7 := GenS2{F0: -79}
        _ = v7
    case "":
        var v8 uint8 = 11
        _ = v8
    default:
        p1 = (p1 + p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int = 4294967295
    var v2 bool = (81 == 2147483647)
    v3 := make(map[string]string)
    var v4 int32 = 0
    for fi9 := 0; fi9 < 9; fi9++ {
        v1 -= (fi9 & v1)
        _ = fi9
    }
    v2 = (v1 == v1)
    v4 += (v4 / ((495) | 1))
    v10 := []int{(v1 + v1)}
    v11 := GenS2{F0: -70}
    if v2 {
        v12 := GenS2{F0: 32}
        _ = len(v3)
        fmt.Println(v4, v1)
        _ = v12
    } else {
        v1 -= (v1 >> ((65535) & 63))
        var v13 int16 = -197
        v2 = (-1 <= -1)
        _ = v13
    }
    v2 = (1 <= -654)
    if true {
        v4 = 2147483647
        v14 := v3["foo bar"]
        _ = v14
    }
    v10 = append(v10, (-(v1)))
    fmt.Println(v4)
    v2 = v2
    var v15 uint32 = 286
    if (v2 || true) {
        if v2 {
            var v16 string = "hello"
            if v2 {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v2 = v2
                switch v16 {
                case "":
                    v10 = append(v10, (v1 | 128))
                case "hello":
                    var v17 uint8 = 255
                    _ = v17
                case "\n\t":
                    if false {
                        fmt.Println(v2)
                        if false {
                            for fi18 := 0; fi18 < 17; fi18++ {
                                fmt.Println(fi18)
                                var v19 uint32 = (v15 * v15)
                                v2 = true
                                _ = fi18
                                _ = v19
                            }
                        } else {
                            v2 = (float32(0.0) != 0.0)
                            v2 = (!v2)
                        }
                    } else {
                        for fi20 := 0; fi20 < 14; fi20++ {
                            if (32767 >= 1) {
                                if true {
                                    switch "a" {
                                    case "foo bar":
                                        _ = len(v3)
                                    }
                                    switch (fi20 << ((18446744073709551615) & 63)) {
                                    case 127:
                                        v3[(v16 + "\n\t")] = v16
                                    default:
                                        if (!v2) {
                                            v4 += -42
                                            delete(v3, v16)
                                        }
                                    }
                                    fmt.Println(v2, v16)
                                }
                            } else {
                                v21 := GenS1{F0: 271, F1: (v15 * (v15 / ((42) | 1)))}
                                _ = v21
                            }
                            v10 = append(v10, (v1 ^ v1))
                            _ = fi20
                        }
                    }
                default:
                    _ = len(v3)
                }
            }
            _ = v16
        } else {
            v22 := GenS3{F0: -104, F1: 2, F2: 255}
            _ = v22
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            var v24 bool = (-2147483648 <= -272)
            _ = fi23
            _ = v24
        }
        var v25 int = v1
        _ = v25
    }
    if (1000 >= 519) {
        v26 := GenS1{F0: 2, F1: v15}
        v1 -= 355
        v1 += (v1 ^ 127)
        _ = v26
    } else {
        var v27 uint16 = 1
        for fi28 := 0; fi28 < 4; fi28++ {
            v4 = (v4 ^ (v4 / ((genFn0(721, 4294967295, "foo bar")) | 1)))
            v2 = v2
            _ = fi28
        }
        _ = v27
    }
    _ = len(v3)
    v29 := v3[""]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
