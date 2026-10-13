package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint, p1 int) uint32 {
    if (959 != 42) {
        if true {
            v6 := []int{(^(p1)), p1, 32768}
            switch p1 {
            case 32767:
                v7 := GenS1{F0: 54, F1: 1, F2: 1}
                _ = v7
            case 4294967295:
                for fi8 := 0; fi8 < 7; fi8++ {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        p1 = (fi8 - -9223372036854775808)
                        _ = len(v6)
                        _ = fi9
                    }
                    fmt.Println(p0, p0)
                    _ = fi8
                }
            default:
                fmt.Println(p1)
            }
            if (2 != 0) {
                for fi10 := 0; fi10 < 6; fi10++ {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v12 := GenS1{F0: 128, F1: 255, F2: 18446744073709551615}
                        _ = fi11
                        _ = v12
                    }
                    v13 := GenS1{F0: 255, F1: 42, F2: uint64(p0)}
                    if len(v6) > 0 {
                        v6[0] = (fi10 - p1)
                    }
                    _ = fi10
                    _ = v13
                }
                if true {
                    p0 = 576
                    var v14 int32 = -59
                    switch "\n\t" {
                    case "hello":
                        var v15 int32 = v14
                        _ = v15
                    }
                    _ = v14
                } else {
                    v16 := GenS1{F0: 38, F1: 64, F2: 0}
                    _ = len(v6)
                    v6 = append(v6, (p1 + (-(p1))))
                    _ = v16
                }
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 int8 = 100
                    p1 = fi17
                    switch "\n\t" {
                    case "\n\t":
                        for fi19 := 0; fi19 < 5; fi19++ {
                            var v20 uint8 = 0
                            p0 = p0
                            var v21 int32 = 382
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            }
            _ = v6
        }
        p1 += (-(p1))
    } else {
        v22 := []int{p1}
        if true {
            var v23 int16 = 1
            var v24 uint32 = 830
            var v25 float32 = -2.25
            _ = v23
            _ = v24
            _ = v25
        }
        _ = v22
    }
    v26 := GenS2{F0: 1}
    v27 := []int{128, (p1 * p1), (p1 ^ p1)}
    p0 += p0
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 4294967295
}

func main() {
    var v1 float32 = float32((0.0 / 1.0))
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    v5 := []int{4294967295, 128, 9223372036854775807, 128}
    if (!(42 != 42)) {
        var v28 float64 = 1.5
        var v29 int8 = 1
        _ = v28
        _ = v29
    }
    var v30 int64 = 106
    var v31 uint32 = 0
    v31 = (v31 ^ (^(v31)))
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 uint32 = (v31 << ((376) & 63))
        _ = fi32
        _ = v33
    }
    v34 := GenS1{F0: 42, F1: 1, F2: 2}
    var v35 uint = 660
    fmt.Println(v1, v35)
    if false {
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 uint8 = 124
            _ = fi36
            _ = v37
        }
        _ = len(v2)
    } else {
        var v38 int = -51
        _ = v38
    }
    v35 = v35
    var v39 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v5
}
