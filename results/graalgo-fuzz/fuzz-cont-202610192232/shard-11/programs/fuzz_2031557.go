package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint8, p1 uint, p2 int32) int8 {
    var v7 float64 = (float64(3.14159) * (0.0 + 0.0))
    p1 += (p1 % (((p1 << ((p1) & 63))) | 1))
    switch "a" {
    case "":
        p1 = p1
    }
    switch "" {
    case "hello":
        if false {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := []int{fi8, fi8}
                _ = fi8
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int32 = (p2 << ((p1) & 63))
                v12 := []int{1, 2147483647}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    case "\n\t":
        for fi13 := 0; fi13 < 20; fi13++ {
            p0 = 199
            if false {
                var v14 int64 = 134
                p1 += p1
                v15 := []int{(fi13 | (^(fi13)))}
                _ = v14
                _ = v15
            } else {
                v16 := []int{}
                if len(v16) > 0 {
                    v16[0] = -798
                }
                p2 = p2
                _ = v16
            }
            _ = fi13
        }
    default:
        v17 := GenS1{F0: 1, F1: 42}
        _ = v17
    }
    v7 -= v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -108
}
func genFn1(p0 int8) uint64 {
    var v18 int64 = 665
    v19 := []int{-534, -315, 127, -927}
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            if (true || (127 >= 0)) {
                _ = len(v19)
                v19 = append(v19, 2147483647)
            } else {
                var v22 uint16 = 714
                fmt.Println(v22, v18)
                _ = v22
            }
            if (false && (0 > 42)) {
                var v23 uint = 42
                _ = v23
            } else {
                v24 := GenS2{F0: 0, F1: 1}
                _ = v24
            }
            _ = fi21
        }
        var v25 int = fi20
        _ = fi20
        _ = v25
    }
    switch -830 {
    case -239:
        var v26 uint64 = 0
        _ = v26
    case -9223372036854775808:
        var v27 int16 = 1000
        _ = v27
    }
    _ = p0
    _ = v18
    _ = v19
    return 951
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "\n\t"
    v3 := []uint32{2147483648, 603, 1, 4294967295}
    v4 := []uint32{}
    var v5 int8 = -128
    var v6 int16 = 1000
    for fi28 := 0; fi28 < 7; fi28++ {
        v2 = ((v2 + ((v2 + "") + "foo bar")) + (v2 + "hello"))
        v29 := v1[(fi28 + -9223372036854775808)]
        var v30 uint = 2
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        v5 = (v5 ^ (v5 % ((int8(v6)) | 1)))
        _ = fi31
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 uint8 = 127
        var v34 string = v2
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 int8 = v5
    v5 += (v35 ^ 1)
    var v36 uint32 = 2147483647
    var v37 uint = 42
    v2 = v2
    var v38 uint = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
