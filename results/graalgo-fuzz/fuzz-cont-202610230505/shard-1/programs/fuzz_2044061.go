package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int; F1 uint16; F2 uint32 }

func main() {
    v1 := []string{"hello", (("a" + "a") + ("hello" + ((("a" + "a") + ("foo bar" + "foo bar")) + ""))), "foo bar"}
    v2 := make(map[uint32]int)
    var v3 uint16 = 684
    var v4 int16 = -711
    v5 := make(map[uint32]string)
    v6 := make(map[int]uint32)
    var v7 int16 = (-(v4))
    v4 -= 468
    _ = len(v1)
    v8 := v2[2]
    fmt.Println(v7)
    for fi9 := 0; fi9 < 6; fi9++ {
        switch fi9 {
        case 870:
            v10 := GenS1{F0: 42, F1: (v7 / ((v7) | 1)), F2: 336}
            _ = v10
        case 32768:
            v4 += (v7 - v7)
        }
        v11 := GenS2{F0: -583, F1: 11}
        _ = fi9
        _ = v11
    }
    var v12 int8 = -128
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS3{F0: 42, F1: "foo bar"}
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 137, F1: 1, F2: -552}
    if false {
        fmt.Println(v3, v3)
        var v16 bool = false
        switch 0 {
        case -434:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -2147483648:
            switch v8 {
            case 4294967295:
                v17 := GenS2{F0: 2147483647, F1: 1}
                _ = v17
            case 32768:
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS2{F0: 42, F1: 47}
                    _ = fi18
                    _ = v19
                }
            case 42:
                v8 = v8
            }
        }
        _ = v16
    }
    v4 = (v7 + (^(v4)))
    fmt.Println(v4, v4, v8)
    v20 := GenS4{F0: -1, F1: v3, F2: 2147483648}
    delete(v5, 780)
    v21 := GenS2{F0: 1, F1: 0}
    for fi22 := 0; fi22 < 18; fi22++ {
        if (1 > (v8 * fi22)) {
            for fi23 := 0; fi23 < 4; fi23++ {
                v4 = (^(v7))
                _ = fi23
            }
            fmt.Println(v12, v7)
        }
        _ = fi22
    }
    v24 := GenS3{F0: 24, F1: "\n\t"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
