package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int64) int {
    var v5 uint8 = 255
    var v6 uint8 = 2
    var v7 uint64 = (p0 | 1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return int(v7)
}
func genFn1(p0 uint64) uint16 {
    switch -9223372036854775808 {
    case -2147483648:
        var v8 uint64 = p0
        _ = v8
    }
    v9 := []int{-9223372036854775808}
    if len(v9) > 0 {
        v9[0] = -2147483648
    }
    _ = p0
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    v3 := []int{32767}
    var v4 float64 = ((-2.25 * float64(1e10)) * 1.5)
    fmt.Println(v4)
    var v10 uint16 = 42
    v10 -= v10
    if false {
        var v11 bool = false
        if len(v2) > 0 {
            v2[0] = 42
        }
        if (1 != 88) {
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: 2, F1: 670}
                if (42 >= 1) {
                    var v14 int64 = 0
                    for fi15 := 0; fi15 < 4; fi15++ {
                        fmt.Println(fi12, fi12, v10)
                        _ = fi15
                    }
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
        } else {
            v11 = (0 >= 42)
        }
        _ = v11
    }
    var v16 bool = true
    var v17 bool = (v16 || v16)
    v4 = v4
    v18 := GenS1{F0: 235, F1: 18446744073709551615}
    var v19 int64 = 23
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int16 = -1
        _ = len(v1)
        _ = fi20
        _ = v21
    }
    switch "\n\t" {
    case "":
        for fi22 := 0; fi22 < 18; fi22++ {
            v10 = (^(v10))
            var v23 float64 = v4
            if v17 {
                fmt.Println(v23, v10)
                v17 = v16
                switch "foo bar" {
                case "a":
                    for fi24 := 0; fi24 < 10; fi24++ {
                        var v25 int8 = -128
                        _ = fi24
                        _ = v25
                    }
                case "hello":
                    v26 := GenS2{F0: 16}
                    _ = v26
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v27 uint = 558
                    _ = v27
                }
            } else {
                v16 = ((^(v19)) > -9223372036854775808)
                var v28 bool = (!(0 == 933))
                _ = v28
            }
            _ = fi22
            _ = v23
        }
    default:
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 int16 = 42
            _ = fi29
            _ = v30
        }
    }
    if v17 {
        fmt.Println(v10, v4, v19)
        v17 = ("a" != "hello")
        v2 = append(v2, 946)
    }
    var v31 uint64 = 77
    for fi32 := 0; fi32 < 13; fi32++ {
        for fi33 := 0; fi33 < 17; fi33++ {
            v1[genFn0(v31, (v19 & 1))] = (v17 && v16)
            var v34 float32 = (0.1 + (1e10 - (1.5 + 1.5)))
            _ = fi33
            _ = v34
        }
        v35 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615}
        fmt.Println(fi32)
        _ = fi32
        _ = v35
    }
    v4 += v4
    v36 := GenS2{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v4
}
