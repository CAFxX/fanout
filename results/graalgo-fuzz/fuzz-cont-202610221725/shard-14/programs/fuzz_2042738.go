package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint64 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8; F1 uint32; F2 int32 }
func genFn0() uint32 {
    var v5 int32 = 1
    v6 := GenS1{F0: 100, F1: 1, F2: 1}
    var v7 uint16 = 65535
    _ = v5
    _ = v6
    _ = v7
    return 2147483648
}
func genFn1(p0 uint8) uint16 {
    v8 := GenS2{F0: 1000, F1: 2147483647}
    for fi9 := 0; fi9 < 10; fi9++ {
        p0 = p0
        var v10 int32 = 42
        _ = fi9
        _ = v10
    }
    p0 = (p0 >> ((65535) & 63))
    v11 := []int{}
    var v12 uint64 = 4294967295
    _ = p0
    _ = v11
    _ = v12
    _ = v8
    return 467
}

func main() {
    v1 := []uint32{1, 2147483648}
    v2 := make(map[uint32]bool)
    var v3 uint32 = 755
    var v4 string = ("" + ("" + "\n\t"))
    v3 = genFn0()
    var v13 int16 = 1
    fmt.Println(v3, v4, v4)
    v3 -= v3
    var v14 uint16 = 54
    for fi15 := 0; fi15 < 9; fi15++ {
        v4 = ("foo bar" + "a")
        _ = fi15
    }
    var v16 uint8 = 42
    switch -1 {
    case -693:
        v4 = v4
    default:
        for fi17 := 0; fi17 < 7; fi17++ {
            v3 -= (v3 / ((genFn0()) | 1))
            _ = fi17
        }
    }
    var v18 uint16 = v14
    fmt.Println(v13)
    v13 = -32768
    v19 := v2[genFn0()]
    v20 := v2[v3]
    v21 := GenS3{F0: 100, F1: (v3 * v3), F2: 179}
    var v22 int8 = 1
    delete(v2, 2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v23 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
