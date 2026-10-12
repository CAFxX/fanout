package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 float32 = ((float32(1.5) + float32(-2.25)) - 0.1)
    var v5 uint16 = 702
    var v6 uint16 = (v5 << ((42) & 63))
    v5 -= (v6 + 32767)
    v4 = -2.25
    var v7 int = 128
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 string = ""
            _ = fi9
            _ = v10
        }
        switch (fi8 * (fi8 - 0)) {
        case -2147483648:
            _ = len(v1)
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: "\n\t", F1: 870}
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    v13 := []int{}
    var v14 string = ("\n\t" + "a")
    v5 += 255
    switch v7 {
    case 32768:
        v6 += 389
    case 128:
        _ = len(v1)
    case -267:
        v14 = "\n\t"
    default:
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int = 32767
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS1{F0: v14, F1: 9223372036854775807}
    fmt.Println(v4)
    var v18 uint32 = 822
    v19 := GenS1{F0: "\n\t", F1: -558}
    var v20 int8 = 100
    var v21 int32 = -875
    v22 := v1[v18]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
