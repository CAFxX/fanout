package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
func genFn0(p0 uint, p1 string, p2 int32) uint {
    var v4 bool = ((false && (!(!((4294967295 <= p0) && (((1e10 - 3.14159) < 0.1) && (!true)))))) && true)
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            switch (^(fi5)) {
            case 4294967295:
                p2 = (p2 ^ p2)
            default:
                v7 := []int{fi5, (fi6 + fi5), 4294967295}
                _ = v7
            }
            _ = fi6
        }
        p1 = "hello"
        _ = fi5
    }
    fmt.Println(p0)
    p1 = "\n\t"
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{fi8, fi8, 128}
        if v4 {
            var v10 uint16 = 42
            v10 = (v10 >> (((p0 << ((p0) & 63))) & 63))
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}
func genFn1(p0 bool) uint32 {
    v11 := GenS1{F0: 1}
    v12 := []int{-938, -1, -362, -485}
    fmt.Println(p0)
    for fi13 := 0; fi13 < 19; fi13++ {
        if (!(42 == 3)) {
            if (4294967295 == 1) {
                p0 = (p0 || (32767 <= 65535))
            } else {
                var v14 uint = 65535
                _ = v14
            }
            var v15 float64 = ((float64(0.0) + 0.0) + 1.5)
            p0 = p0
            _ = v15
        } else {
            fmt.Println(p0)
            var v16 int32 = 42
            _ = v16
        }
        if (128 < 127) {
            var v17 float32 = 0.0
            var v18 int32 = -1
            _ = v17
            _ = v18
        } else {
            if p0 {
                var v19 int32 = 0
                _ = v19
            }
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
        _ = fi13
    }
    _ = p0
    _ = v11
    _ = v12
    return 65535
}

func main() {
    var v1 int32 = 754
    var v2 uint8 = 127
    v3 := make(map[int]uint32)
    var v20 uint8 = (^(v2))
    fmt.Println(v2, v20, v2)
    var v21 string = "\n\t"
    var v22 float64 = 3.14159
    for fi23 := 0; fi23 < 20; fi23++ {
        delete(v3, 790)
        if true {
            v22 += ((v22 + ((((v22 * (v22 / -2.25)) * v22) * v22) - float64(((1.5 - 0.0) + 1.5)))) * v22)
        } else {
            v22 -= ((v22 + v22) / (float64((1.5 * 0.0)) + (v22 / 1.5)))
        }
        _ = fi23
    }
    v24 := GenS2{F0: 0, F1: 4294967295, F2: 433}
    var v25 int32 = v1
    fmt.Println(v2, v22)
    v26 := []int{-9223372036854775808, -428, 4294967295}
    var v27 int16 = -32768
    v28 := GenS2{F0: 588, F1: 4294967295, F2: (^(v27))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
