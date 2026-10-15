package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int32; F1 int8; F2 int8 }
type GenS4 struct { F0 int }
func genFn0(p0 int16) int64 {
    v4 := GenS1{F0: -2147483648, F1: 502, F2: 1}
    if true {
        for fi5 := 0; fi5 < 14; fi5++ {
            p0 = (p0 >> ((687) & 63))
            _ = fi5
        }
    }
    p0 = (p0 << ((255) & 63))
    p0 -= (p0 & 32767)
    v6 := []int{}
    _ = p0
    _ = v4
    _ = v6
    return 1
}
func genFn1(p0 bool, p1 uint32) uint32 {
    v7 := GenS2{F0: 9223372036854775807, F1: 2147483647}
    var v8 int = 128
    p0 = p0
    var v9 bool = (0 < 2)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p1 / ((p1) | 1))
}

func main() {
    var v1 bool = false
    v2 := []int{1, 354, 127, 1}
    var v3 uint64 = 843
    var v10 float32 = float32((1e10 - 0.0))
    if (493 == 9223372036854775807) {
        var v11 bool = (v1 || (1 > 1))
        for fi12 := 0; fi12 < 17; fi12++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi12
        }
        if len(v2) > 0 {
            v2[0] = 1
        }
        _ = v11
    }
    v2 = append(v2, 0)
    var v13 uint64 = (v3 % ((v3) | 1))
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    switch 128 {
    case -26:
        if len(v2) > 0 {
            v2[0] = -291
        }
    case -699:
        if (-2147483648 == -9223372036854775808) {
            if (891 == 1) {
                if ((1.5 != 1.5) || v1) {
                    var v14 int16 = -32768
                    _ = v14
                } else {
                    fmt.Println(v1)
                    _ = len(v2)
                }
            } else {
                fmt.Println(v13)
                var v15 uint = 18446744073709551615
                _ = v15
            }
        } else {
            v13 -= (v3 * (v3 & (v3 - (v3 | v13))))
            switch "a" {
            case "hello":
                var v16 uint8 = 128
                _ = v16
            case "a":
                v10 += v10
            case "foo bar":
                if (((1.5 * (float32(1.5) - v10)) - 0.1) == v10) {
                    v10 = v10
                    var v17 float32 = -2.25
                    if true {
                        var v18 uint32 = 0
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v18
                    }
                    _ = v17
                } else {
                    v19 := GenS1{F0: -2147483648, F1: 4294967295, F2: genFn0(1000)}
                    _ = v19
                }
            default:
                if (2 < 1) {
                    switch 42 {
                    case -2147483648:
                        var v20 string = ("\n\t" + (((("hello" + "\n\t") + "a") + "a") + "foo bar"))
                        _ = v20
                    case -689:
                        v21 := GenS3{F0: 0, F1: -128, F2: 100}
                        _ = v21
                    case 32767:
                        v10 = 1e10
                    }
                    fmt.Println(v13)
                    if (2147483647 == 42) {
                        _ = len(v2)
                        for fi22 := 0; fi22 < 13; fi22++ {
                            var v23 uint64 = (^(v13))
                            _ = fi22
                            _ = v23
                        }
                        var v24 int = -655
                        _ = v24
                    }
                }
            }
        }
    default:
        fmt.Println(v3, v1, v10)
    }
    fmt.Println(v3)
    v3 += (v13 * (^(v13)))
    v13 = 164
    fmt.Println(v3, v3, v10)
    fmt.Println(v3, v1, v1)
    switch -9223372036854775808 {
    case 32767:
        var v25 float32 = v10
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        fmt.Println(v10, v10, v10)
        if len(v2) > 0 {
            v2[0] = fi26
        }
        var v27 bool = (v1 || v1)
        _ = fi26
        _ = v27
    }
    if ((((float64(v10) + ((1e10 / 1e10) + 1e10)) * 0.1) + (3.14159 / 1e10)) >= 0.0) {
        v3 = 4294967295
        for fi28 := 0; fi28 < 13; fi28++ {
            fmt.Println(v13)
            fmt.Println(v3, v1, v3)
            v3 = 37
            _ = fi28
        }
    } else {
        var v29 uint = 65535
        _ = v29
    }
    v30 := GenS3{F0: 794, F1: -98, F2: 42}
    v31 := GenS4{F0: 46}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
