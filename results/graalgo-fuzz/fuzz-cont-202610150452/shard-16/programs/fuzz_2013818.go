package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint16 }
func genFn0(p0 uint16, p1 uint16, p2 int8) uint32 {
    if (107 <= 194) {
        if (-128 >= (p2 * 42)) {
            var v5 uint64 = 233
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 float64 = 0.1
                _ = fi6
                _ = v7
            }
            _ = v5
        }
    }
    if true {
        var v8 int32 = -144
        v9 := GenS1{F0: true}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int16 = 558
    v2 := make(map[int]uint32)
    var v3 uint32 = 2147483647
    v4 := make(map[int]string)
    if (32767 == 0) {
        v10 := []int{-401, 2147483647, -1, 127}
        _ = v10
    }
    var v11 uint8 = 128
    v11 += 15
    var v12 float64 = 3.14159
    for fi13 := 0; fi13 < 20; fi13++ {
        switch 9223372036854775807 {
        case -922:
            v14 := v2[fi13]
            _ = v14
        case 9:
            v11 = v11
        }
        v11 += 127
        if (-1 < 42) {
            v15 := GenS1{F0: true}
            v3 += (^(v3))
            if false {
                if (0.1 != 1.5) {
                    v16 := GenS2{F0: "hello"}
                    v17 := GenS1{F0: (128 <= (v11 >> ((42) & 63)))}
                    for fi18 := 0; fi18 < 8; fi18++ {
                        fmt.Println(fi13)
                        _ = fi18
                    }
                    _ = v16
                    _ = v17
                }
                v19 := GenS3{F0: -412, F1: 2}
                _ = v19
            }
            _ = v15
        } else {
            v11 = 122
            var v20 int32 = 42
            _ = v20
        }
        _ = fi13
    }
    var v21 uint32 = (v3 ^ (^(v3)))
    v1 += v1
    v11 -= v11
    v21 = (v21 & v21)
    var v22 int64 = -708
    v23 := []int{-9223372036854775808, -346, 42, 127}
    fmt.Println(v11, v22, v21)
    v1 = (-(v1))
    if (v12 >= 0.1) {
        for fi24 := 0; fi24 < 11; fi24++ {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            v3 -= (v3 << ((4294967295) & 63))
            _ = fi24
        }
    } else {
        if false {
            v22 += 0
            v25 := GenS2{F0: ("a" + ("\n\t" + "a"))}
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = v25
        } else {
            var v26 int8 = -48
            _ = v26
        }
        var v27 uint8 = v11
        _ = v27
    }
    _ = len(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
