package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int32 = 1
    v2 := []uint32{872, 2147483648}
    v3 := []uint32{184, 1}
    v4 := []uint32{}
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: (v1 % ((-1) | 1)), F1: 65535}
    if false {
        v7 := GenS1{F0: v1, F1: 255}
        var v8 float32 = 0.0
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint8 = 146
            var v11 float32 = -2.25
            v8 = v11
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
        _ = v8
    } else {
        v12 := GenS2{F0: 32767, F1: 86}
        _ = v12
    }
    v1 = 500
    for fi13 := 0; fi13 < 9; fi13++ {
        switch (fi13 / (((fi13 | (fi13 | (fi13 + (fi13 ^ (^(fi13))))))) | 1)) {
        case -1:
            _ = len(v5)
        }
        var v14 int = fi13
        var v15 uint64 = 0
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v1)
    var v16 uint16 = 32767
    delete(v5, 2147483648)
    delete(v5, 4294967295)
    var v17 uint32 = 4294967295
    v1 -= (v1 % ((42) | 1))
    var v18 uint8 = 255
    if true {
        var v19 uint8 = 216
        if (!true) {
            var v20 int16 = 288
            _ = v20
        } else {
            v1 = 530
        }
        _ = v19
    } else {
        delete(v5, (v17 + (^(v17))))
        v17 = 2147483648
    }
    fmt.Println(v1, v17, v17)
    v1 += v1
    fmt.Println(v17, v18)
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 int64 = -1
        v22 -= (v22 - -9223372036854775808)
        var v23 bool = ((v22 << ((uint(v16)) & 63)) != v22)
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: v1, F1: v16}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
