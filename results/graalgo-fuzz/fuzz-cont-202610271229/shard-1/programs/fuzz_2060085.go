package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int8 }
type GenS2 struct { F0 int8; F1 int32; F2 int }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []int{}
    var v2 uint8 = 127
    var v3 int = 32767
    var v4 int16 = -1
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            if len(v1) > 0 {
                v1[0] = -5
            }
            if (!((^(v2)) < (v2 * v2))) {
                var v7 uint = 65535
                v3 = 128
                fmt.Println(v4, v7)
                _ = v7
            } else {
                v8 := GenS1{F0: 255, F1: (-(fi5)), F2: -1}
                _ = v8
            }
            _ = fi6
        }
        var v9 uint16 = 0
        v9 = 65535
        _ = fi5
        _ = v9
    }
    v2 = (v2 ^ 164)
    fmt.Println(v2, v2)
    if false {
        var v10 float32 = (((-2.25 * (3.14159 * 1.5)) + 0.1) + 3.14159)
        v11 := GenS2{F0: -128, F1: 1, F2: -728}
        if true {
            v3 += v3
            var v12 float32 = float32(1.5)
            _ = v12
        }
        _ = v10
        _ = v11
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 bool = (true || ("a" > "\n\t"))
            fmt.Println(v4, v14, v14)
            for fi15 := 0; fi15 < 18; fi15++ {
                v1 = append(v1, fi15)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
        var v16 float64 = 0.0
        if false {
            var v17 int16 = (v4 / (((-(v4))) | 1))
            fmt.Println(v2, v17, v3)
            _ = v17
        }
        _ = v16
    }
    v2 = (v2 >> ((1) & 63))
    v18 := GenS3{F0: 42}
    var v19 uint8 = (v2 - (v2 << ((105) & 63)))
    v20 := GenS1{F0: 1, F1: v3, F2: -128}
    v21 := GenS2{F0: 100, F1: 42, F2: (v3 % ((127) | 1))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
