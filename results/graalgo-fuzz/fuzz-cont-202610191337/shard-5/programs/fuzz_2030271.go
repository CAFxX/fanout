package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0(p0 int32) int16 {
    v4 := []int{464, 1}
    v5 := GenS1{F0: -409, F1: 2147483647}
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return 793
}
func genFn1(p0 uint16, p1 int8, p2 uint8) int32 {
    p1 = (p1 >> ((65535) & 63))
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint = 154
        fmt.Println(p1, p1, fi6)
        v8 := []int{(fi6 ^ (fi6 >> ((v7) & 63))), (^(fi6)), fi6, -86}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 0
    p2 = 224
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 305
    var v3 bool = ((2147483648 != 519) || false)
    v10 := GenS1{F0: 42, F1: 1}
    v2 += (v2 - 0)
    var v11 string = "hello"
    var v12 float64 = 0.0
    for fi13 := 0; fi13 < 2; fi13++ {
        _ = len(v1)
        var v14 uint32 = 415
        _ = fi13
        _ = v14
    }
    v15 := []int{-927, -9223372036854775808, 2147483647}
    v3 = (!false)
    v16 := GenS2{F0: 109, F1: 42}
    if v3 {
        var v17 uint8 = 224
        fmt.Println(v17)
        _ = v17
    } else {
        if false {
            if len(v15) > 0 {
                v15[0] = -665
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                delete(v1, fi18)
                _ = fi18
            }
            v1[32768] = 1
        }
    }
    var v19 int8 = 100
    var v20 uint32 = 564
    var v21 int32 = (v2 % (((v2 % ((v2) | 1))) | 1))
    v12 += (0.0 - (((float64(1.5) - float64(float32(v12))) * (-2.25 * v12)) + float64(-2.25)))
    for fi22 := 0; fi22 < 20; fi22++ {
        if ((-32768 != -1) || (!(145 == 1))) {
            var v23 uint64 = 1
            var v24 uint16 = 1
            _ = v23
            _ = v24
        }
        _ = len(v15)
        _ = fi22
    }
    fmt.Println(v20, v2)
    if (!false) {
        if v3 {
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 uint16 = 63
                var v27 uint16 = (v26 - v26)
                _ = fi25
                _ = v26
                _ = v27
            }
        } else {
            v28 := GenS1{F0: genFn0(v2), F1: v20}
            _ = v28
        }
        var v29 uint16 = 178
        var v30 int8 = v19
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
