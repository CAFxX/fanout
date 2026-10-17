package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
func genFn0(p0 int32, p1 int, p2 bool) uint64 {
    v7 := GenS1{F0: 1}
    if (1 > 32767) {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := []int{fi8, (fi8 % (((fi8 & -796)) | 1)), (fi8 & (p1 << ((18446744073709551615) & 63)))}
            p2 = p2
            _ = fi8
            _ = v9
        }
        if (p2 || p2) {
            if true {
                v10 := GenS1{F0: int8(p1)}
                var v11 int16 = 1000
                v12 := []int{(p1 % ((1) | 1)), 32768, 42}
                _ = v10
                _ = v11
                _ = v12
            }
        }
    }
    v13 := []int{(p1 & p1), (p1 << ((4294967295) & 63))}
    var v14 int8 = 76
    var v15 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    _ = v7
    return 4294967295
}
func genFn1(p0 int32, p1 uint, p2 int16) int64 {
    v16 := []int{}
    var v17 int16 = (p2 | p2)
    v17 = (p2 / (((v17 << ((p1) & 63))) | 1))
    v17 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return -1
}

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    var v3 uint = 2
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    v6 := make(map[int]string)
    v18 := GenS2{F0: ("\n\t" + "hello"), F1: 2147483648, F2: 42}
    var v19 float64 = -2.25
    for fi20 := 0; fi20 < 12; fi20++ {
        v3 += v3
        v6[fi20] = "foo bar"
        _ = fi20
    }
    v21 := []int{32767, 32768}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v19 -= v19
    switch -436 {
    case -2147483648:
        var v22 float64 = v19
        _ = v22
    default:
        if ((v3 < v3) && true) {
            v23 := GenS2{F0: "\n\t", F1: 42, F2: -288}
            fmt.Println(v3, v1)
            fmt.Println(v1)
            _ = v23
        }
    }
    if (255 <= 0) {
        var v24 int8 = 81
        v25 := GenS1{F0: v24}
        _ = v24
        _ = v25
    }
    switch 127 {
    case 9223372036854775807:
        for fi26 := 0; fi26 < 18; fi26++ {
            v3 = v3
            _ = fi26
        }
    case 2147483647:
        var v27 int32 = 2147483647
        _ = v27
    case 1:
        v28 := GenS2{F0: ("foo bar" + (("\n\t" + "") + ("hello" + "a"))), F1: 42, F2: -9223372036854775808}
        _ = v28
    }
    var v29 uint64 = 42
    v30 := v4[uint32(v3)]
    var v31 int64 = -986
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
