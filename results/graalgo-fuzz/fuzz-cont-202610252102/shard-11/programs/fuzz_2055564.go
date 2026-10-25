package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int32 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 int16 = -916
    var v5 uint16 = v1
    v5 += v1
    if (0.1 > 1e10) {
        var v6 uint = 42
        var v7 int16 = v4
        var v8 uint8 = 2
        _ = v6
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int16 = (v4 / (((-(v4))) | 1))
            var v11 uint64 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
        switch 1 {
        case -594:
            v12 := GenS1{F0: (v4 % ((v4) | 1)), F1: 0, F2: 1}
            _ = v12
        case -2147483648:
            for fi13 := 0; fi13 < 11; fi13++ {
                if (v4 == (v4 - -1)) {
                    v14 := []int{}
                    _ = v14
                }
                _ = fi13
            }
        case 32767:
            v4 = 42
        }
        v15 := v2[0]
        _ = v15
    }
    v16 := GenS2{F0: 127, F1: (^(v1))}
    v17 := GenS3{F0: ("\n\t" + "a")}
    v18 := []int{}
    var v19 bool = false
    switch 128 {
    case -641:
        v20 := GenS4{F0: -170, F1: 4294967295}
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v3[fi21] = 2147483647
        var v22 int32 = -1
        _ = fi21
        _ = v22
    }
    v1 -= 1
    v19 = v19
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    v23 := GenS4{F0: 1, F1: 712}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
