package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
func genFn0(p0 int8, p1 int16) uint8 {
    v6 := GenS1{F0: 1, F1: 2}
    var v7 uint = 23
    v8 := []int{127}
    if len(v8) > 0 {
        v8[0] = 343
    }
    if ((("\n\t" + "\n\t") < ("" + "foo bar")) || (true && ((p1 & p1) > p1))) {
        var v9 uint = v7
        _ = v9
    } else {
        v10 := GenS1{F0: -9223372036854775808, F1: 128}
        for fi11 := 0; fi11 < 7; fi11++ {
            if false {
                var v12 string = "foo bar"
                _ = v12
            } else {
                fmt.Println(fi11, fi11)
                p0 -= (p0 * p0)
            }
            var v13 uint64 = 18446744073709551615
            _ = len(v8)
            _ = fi11
            _ = v13
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 134
}

func main() {
    var v1 int64 = -1
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    v4 := []uint32{0, 0, 621}
    v5 := make(map[int]string)
    switch "a" {
    case "\n\t":
        fmt.Println(v1)
    case "hello":
        v1 = v1
    }
    v1 += 0
    if (false || false) {
        v3[83] = 42
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        v1 += v1
        if (255 == 65535) {
            var v14 int8 = 47
            var v15 int16 = 0
            var v16 int16 = (v15 << ((2) & 63))
            _ = v14
            _ = v15
            _ = v16
        } else {
            switch "" {
            case "a":
                v17 := GenS1{F0: v1, F1: 127}
                _ = v17
            case "":
                v1 -= 657
            case "foo bar":
                v1 -= (v1 - -9223372036854775808)
            }
            if false {
                v1 = (v1 % ((0) | 1))
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := GenS1{F0: 9223372036854775807, F1: 127}
                    if len(v4) > 0 {
                        v4[0] = 65535
                    }
                    if ((false || false) || false) {
                        var v20 float32 = 1e10
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 4; fi21++ {
                            fmt.Println(v1, v1)
                            _ = len(v4)
                            if false {
                                if true {
                                    v1 -= v1
                                }
                            }
                            _ = fi21
                        }
                    }
                    _ = fi18
                    _ = v19
                }
                v4 = append(v4, 65535)
            } else {
                var v22 int8 = -1
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                v23 := GenS1{F0: -1, F1: genFn0(v22, -32768)}
                _ = v22
                _ = v23
            }
        }
    }
    fmt.Println(v1)
    v24 := GenS1{F0: v1, F1: genFn0(100, 1)}
    var v25 uint32 = 2147483648
    if (18446744073709551615 >= 172) {
        v1 += (v1 >> ((65535) & 63))
    } else {
        if false {
            var v26 float32 = float32(0.0)
            _ = v26
        } else {
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 int32 = -199
                _ = fi27
                _ = v28
            }
        }
    }
    switch "a" {
    case "\n\t":
        v1 = (v1 & 1)
    default:
        var v29 uint8 = 2
        _ = v29
    }
    v30 := GenS1{F0: (v1 ^ v1), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
