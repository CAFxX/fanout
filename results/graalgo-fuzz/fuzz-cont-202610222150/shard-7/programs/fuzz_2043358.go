package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8; F2 uint16 }
type GenS3 struct { F0 string }
func genFn0() int64 {
    if false {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 float32 = -2.25
            if true {
                v9 := GenS1{F0: 4294967295}
                _ = v9
            } else {
                if (100 != 0) {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        v11 := []int{32768, (^(fi7)), -243, (fi10 + fi10)}
                        _ = fi10
                        _ = v11
                    }
                }
                v12 := []int{fi7, 4294967295, fi7}
                for fi13 := 0; fi13 < 10; fi13++ {
                    v12 = append(v12, fi7)
                    _ = fi13
                }
                _ = v12
            }
            _ = fi7
            _ = v8
        }
        var v14 bool = true
        _ = v14
    } else {
        var v15 uint = 707
        var v16 int = 127
        fmt.Println(v16, v16)
        _ = v15
        _ = v16
    }
    var v17 int64 = 562
    _ = v17
    return -828
}
func genFn1(p0 uint8, p1 int, p2 uint32) uint64 {
    if (0.0 >= (0.1 + (((0.0 * -2.25) - (1.5 / 1.0)) + float32(float64((-2.25 + -2.25)))))) {
        p2 = p2
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        p1 = (fi18 ^ 0)
        p2 += (p2 << ((18446744073709551615) & 63))
        p0 += 1
        _ = fi18
    }
    v19 := GenS1{F0: 2}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 4294967295
}

func main() {
    var v1 bool = true
    v2 := make(map[int]int)
    v3 := []uint32{42, 2147483648, 693, 42}
    var v4 uint = 18446744073709551615
    var v5 uint16 = 42
    var v6 uint = (v4 & (v4 >> (((v4 << (((v4 / (((^(v4))) | 1))) & 63))) & 63)))
    for fi20 := 0; fi20 < 2; fi20++ {
        if v1 {
            v21 := GenS2{F0: (fi20 + (fi20 + 526)), F1: 34, F2: v5}
            var v22 int64 = int64(fi20)
            var v23 int32 = -2147483648
            _ = v21
            _ = v22
            _ = v23
        } else {
            v1 = v1
        }
        _ = fi20
    }
    v1 = v1
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 float32 = 0.1
        _ = fi24
        _ = v25
    }
    var v26 float32 = -2.25
    if true {
        fmt.Println(v4, v26, v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v27 int8 = 42
        fmt.Println(v4, v4)
        _ = v27
    }
    fmt.Println(v26, v26, v26)
    for fi28 := 0; fi28 < 13; fi28++ {
        v4 -= (v6 ^ v6)
        switch (fi28 ^ fi28) {
        case 42:
            var v29 string = ""
            _ = v29
        case 0:
            v30 := v2[fi28]
            _ = v30
        default:
            v2[-9223372036854775808] = (fi28 - (fi28 * -229))
        }
        v31 := GenS3{F0: "foo bar"}
        _ = fi28
        _ = v31
    }
    var v32 uint32 = 2
    v32 = 2
    fmt.Println(v4)
    fmt.Println(v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
