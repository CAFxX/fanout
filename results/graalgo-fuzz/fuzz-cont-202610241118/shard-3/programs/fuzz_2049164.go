package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint64 }
func genFn0() uint64 {
    var v4 uint = 42
    if (4294967295 == 2147483647) {
        v4 = (v4 * v4)
        v4 = v4
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 bool = true
            v7 := []int{}
            if v6 {
                fmt.Println(fi5)
                var v8 uint16 = 932
                _ = v8
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        var v9 uint = v4
        var v10 bool = true
        if v10 {
            var v11 int = -1
            v11 += (v11 << (((^(v4))) & 63))
            _ = v11
        }
        _ = v10
        _ = v9
    }
    _ = v4
    return 4294967295
}

func main() {
    var v1 uint16 = 1
    var v2 uint8 = 0
    var v3 int8 = -75
    v12 := GenS1{F0: 22}
    if false {
        v2 = (v2 - v2)
        v3 = v3
    } else {
        v2 += 128
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 uint32 = 410
            _ = fi14
            _ = v15
        }
        fmt.Println(fi13)
        _ = fi13
    }
    v16 := GenS2{F0: (v2 % ((108) | 1))}
    v1 = v1
    v17 := []int{127, 32767, 42, 127}
    var v18 int = 32767
    var v19 uint64 = genFn0()
    v20 := GenS3{F0: -148, F1: -538}
    switch (v18 << ((255) & 63)) {
    case 9223372036854775807:
        if false {
            fmt.Println(v1, v3, v19)
        }
    default:
        var v21 int16 = 32767
        _ = v21
    }
    v1 = 1
    v22 := GenS4{F0: (v1 * 1), F1: v19}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
