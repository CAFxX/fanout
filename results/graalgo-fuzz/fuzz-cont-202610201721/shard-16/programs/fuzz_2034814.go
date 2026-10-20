package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }

func main() {
    var v1 int32 = 42
    var v2 uint32 = 42
    var v3 uint32 = 2147483647
    v4 := make(map[int]string)
    v1 -= (v1 * v1)
    if (9223372036854775807 != -985) {
        v1 = 2147483647
        for fi5 := 0; fi5 < 12; fi5++ {
            v1 -= v1
            _ = fi5
        }
        v1 = (v1 << ((uint(v2)) & 63))
    } else {
        v3 -= (v2 | v2)
        v4[624] = "foo bar"
    }
    if false {
        for fi6 := 0; fi6 < 13; fi6++ {
            if (!false) {
                switch fi6 {
                case 32768:
                    var v7 uint = 255
                    _ = v7
                case 127:
                    for fi8 := 0; fi8 < 12; fi8++ {
                        v9 := GenS1{F0: (fi8 << ((65535) & 63))}
                        v10 := []int{0, (fi6 * 0)}
                        v2 += (v2 & v2)
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case 32767:
                    fmt.Println(v1, v1, v3)
                default:
                    for fi11 := 0; fi11 < 14; fi11++ {
                        v12 := []int{}
                        if ((!true) && true) {
                            var v13 bool = (float64(-2.25) != 1.5)
                            v2 += v3
                            var v14 int16 = -915
                            if len(v12) > 0 {
                                v12[0] = 9223372036854775807
                            }
                            _ = v13
                            _ = v14
                        }
                        _ = fi11
                        _ = v12
                    }
                }
                var v15 string = "\n\t"
                _ = v15
            } else {
                v16 := GenS2{F0: -95, F1: true, F2: 32767}
                fmt.Println(v3)
                var v17 float32 = 1.5
                _ = v16
                _ = v17
            }
            _ = fi6
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS2{F0: -70, F1: (4294967295 == 444), F2: 32767}
            _ = fi18
            _ = v19
        }
        v3 += v2
    } else {
        v1 = v1
    }
    switch 960 {
    case 32767:
        var v20 int32 = v1
        _ = v20
    default:
        var v21 int32 = (-(v1))
        _ = v21
    }
    v22 := []int{-987, 32767, 1}
    fmt.Println(v1, v1)
    v3 += (v3 - 42)
    if true {
        var v23 int32 = v1
        var v24 uint64 = 4294967295
        if (true && (false && (81 > (v3 ^ (v3 * (v2 + (v2 << ((449) & 63)))))))) {
            if len(v22) > 0 {
                v22[0] = -600
            }
            fmt.Println(v2)
        }
        _ = v23
        _ = v24
    }
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
