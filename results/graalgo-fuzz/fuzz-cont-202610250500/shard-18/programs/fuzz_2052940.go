package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint8) int64 {
    v7 := GenS1{F0: 0, F1: 175}
    v8 := []int{576, -964}
    var v9 uint32 = 4294967295
    switch "foo bar" {
    case "":
        for fi10 := 0; fi10 < 7; fi10++ {
            if false {
                _ = len(v8)
                v11 := GenS2{F0: (^(p1))}
                _ = v11
            }
            var v12 float64 = ((-2.25 * (0.1 + 3.14159)) * 1.5)
            v9 -= 4294967295
            _ = fi10
            _ = v12
        }
    case "hello":
        if (true && (false && (1 == 1))) {
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS2{F0: p1}
                if (42 <= 565) {
                    if (p1 != (p1 - (p1 ^ p1))) {
                        if ((128 > (p1 / (((p1 | (^(p1)))) | 1))) || (!(1 != 1000))) {
                            v9 = v9
                        } else {
                            v15 := GenS2{F0: p1}
                            var v16 int = -1
                            v9 = (^(v9))
                            _ = v15
                            _ = v16
                        }
                        p0 = 530
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                }
                _ = fi13
                _ = v14
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        } else {
            v17 := GenS3{F0: -32768}
            v9 += 0
            _ = v17
        }
    default:
        if (!false) {
            p0 = (p0 * (p0 << ((2) & 63)))
            if len(v8) > 0 {
                v8[0] = -9223372036854775808
            }
            if (!false) {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                fmt.Println(v9, p1, v9)
            } else {
                if len(v8) > 0 {
                    v8[0] = int(v9)
                }
                var v18 int8 = -1
                _ = v18
            }
        } else {
            var v19 uint64 = 858
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 bool = true
    var v2 string = "a"
    v3 := []int{9223372036854775807, 0, 637, 4294967295}
    var v4 int8 = 42
    var v5 int16 = 1
    v6 := make(map[int]string)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 int64 = 9223372036854775807
    v21 := GenS2{F0: 0}
    fmt.Println(v2, v2)
    if len(v3) > 0 {
        v3[0] = 1
    }
    v6[0] = v2
    v22 := GenS4{F0: -1}
    if v1 {
        var v23 string = "a"
        if v1 {
            v4 -= (-(v4))
        } else {
            var v24 uint32 = 745
            var v25 string = v23
            _ = v24
            _ = v25
        }
        _ = v23
    } else {
        v2 = "hello"
        fmt.Println(v2, v1, v2)
        for fi26 := 0; fi26 < 6; fi26++ {
            fmt.Println(v1, v5, v1)
            for fi27 := 0; fi27 < 15; fi27++ {
                var v28 int64 = (v20 - 0)
                v5 -= (v5 ^ (v5 * (v5 - 0)))
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
