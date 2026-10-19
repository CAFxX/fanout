package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16 }
func genFn0() int {
    var v5 uint16 = 265
    if true {
        var v6 uint = 4294967295
        _ = v6
    }
    v5 = (^(v5))
    fmt.Println(v5)
    _ = v5
    return 4294967295
}
func genFn1(p0 int64, p1 uint32, p2 int64) uint {
    if (2 < 4294967295) {
        var v7 int32 = 42
        p2 = 0
        _ = v7
    } else {
        switch "foo bar" {
        case "":
            for fi8 := 0; fi8 < 4; fi8++ {
                p1 = (p1 ^ 0)
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint64 = 4294967295
                    if ((fi8 & fi9) < (-(fi8))) {
                        v11 := []int{(fi9 & 9223372036854775807), (-(fi9)), (fi8 % ((42) | 1)), 753}
                        var v12 float64 = 1.5
                        _ = v11
                        _ = v12
                    } else {
                        p2 = (-(p2))
                        v13 := GenS1{F0: -519, F1: 2147483647}
                        var v14 uint8 = 2
                        _ = v13
                        _ = v14
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        case "\n\t":
            v15 := []int{42, 42, 32767, 127}
            _ = v15
        }
        if false {
            p0 += p2
            if true {
                p1 = (p1 << ((255) & 63))
                p0 = (p2 - p2)
                v16 := GenS1{F0: -2147483648, F1: 42}
                _ = v16
            } else {
                p1 += (p1 >> ((418) & 63))
                var v17 int8 = 46
                _ = v17
            }
            for fi18 := 0; fi18 < 9; fi18++ {
                fmt.Println(p0, p1)
                _ = fi18
            }
        } else {
            p2 += p2
            p0 -= 9223372036854775807
            v19 := GenS1{F0: 42, F1: -2147483648}
            _ = v19
        }
    }
    var v20 int16 = 1
    if true {
        v21 := []int{0, 1}
        v22 := GenS1{F0: -2147483648, F1: 2147483647}
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 1, F1: 797}
    var v24 uint = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v23
    _ = v24
    return v24
}

func main() {
    var v1 uint64 = 541
    var v2 bool = true
    v3 := []uint32{2, 4294967295, 65535}
    var v4 int8 = -107
    v4 += v4
    var v25 bool = v2
    var v26 uint16 = 2
    v3 = append(v3, 65535)
    var v27 uint64 = v1
    for fi28 := 0; fi28 < 15; fi28++ {
        v2 = v25
        var v29 string = "a"
        _ = fi28
        _ = v29
    }
    if v25 {
        fmt.Println(v1, v27)
        switch "a" {
        case "":
            for fi30 := 0; fi30 < 12; fi30++ {
                v31 := GenS1{F0: -771, F1: -1}
                var v32 uint32 = 506
                for fi33 := 0; fi33 < 18; fi33++ {
                    if (2 == 18446744073709551615) {
                        v34 := GenS1{F0: -1, F1: 1}
                        if (42 == 32767) {
                            v35 := GenS2{F0: -390}
                            v3 = append(v3, (v32 ^ 1))
                            v26 += v26
                            _ = v35
                        }
                        _ = v34
                    }
                    var v36 int64 = 0
                    _ = fi33
                    _ = v36
                }
                _ = fi30
                _ = v31
                _ = v32
            }
        case "foo bar":
            v2 = v25
        }
    }
    var v37 int8 = v4
    var v38 int8 = (v37 / (((-(v4))) | 1))
    v26 += 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
