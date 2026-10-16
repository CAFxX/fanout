package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint16 }
func genFn0(p0 uint16, p1 string, p2 uint8) int16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{fi6}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0, p0, p1)
    var v8 string = (p1 + "foo bar")
    var v9 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -32768
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 42
    v3 := []uint32{1}
    var v4 float32 = (0.1 / 1.0)
    var v5 uint8 = 128
    switch ("\n\t" + ("" + "foo bar")) {
    case "":
        v10 := GenS1{F0: -607, F1: 42, F2: 79}
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS2{F0: 4294967295, F1: 255}
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v5)
    var v13 uint = 720
    _ = len(v3)
    v3 = append(v3, 463)
    var v14 int64 = -74
    v2 = (^(v2))
    for fi15 := 0; fi15 < 18; fi15++ {
        _ = len(v3)
        _ = fi15
    }
    var v16 int64 = (-(v14))
    if false {
        var v17 int = (v2 << (((^(v13))) & 63))
        fmt.Println(v4, v4)
        _ = v17
    }
    v13 -= (v13 & v13)
    if (0.0 < (((v4 - v4) + v4) - v4)) {
        fmt.Println(v5, v16, v2)
        v18 := GenS3{F0: (^(v13)), F1: 1, F2: 173}
        _ = v18
    } else {
        v14 += v14
        if (!(v2 < 32767)) {
            v5 -= v5
        }
    }
    v14 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
