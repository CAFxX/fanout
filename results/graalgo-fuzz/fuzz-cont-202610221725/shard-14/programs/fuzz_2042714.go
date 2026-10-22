package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1e10
    v3 += v3
    for fi4 := 0; fi4 < 12; fi4++ {
        fmt.Println(v1)
        _ = fi4
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        if (42 <= -128) {
            v3 -= -2.25
        } else {
            for fi6 := 0; fi6 < 13; fi6++ {
                v1 -= (^(v1))
                if ((0.0 / 1.5) >= (float64((float32(((1e10 / 1.0) * 0.1)) / (v3 + 1.5))) / 1.0)) {
                    fmt.Println(fi5, v1, v3)
                }
                _ = fi6
            }
            if true {
                switch fi5 {
                case -1:
                    var v7 uint8 = 0
                    _ = v7
                default:
                    v1 += v1
                }
                var v8 uint16 = 32767
                fmt.Println(v8)
                _ = v8
            } else {
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 int8 = -107
                    fmt.Println(v3, v1)
                    v11 := []int{(fi5 | (fi5 >> ((0) & 63))), fi9}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            v1 = v1
        }
        _ = fi5
    }
    var v12 float32 = v3
    fmt.Println(v3, v12)
    fmt.Println(v3, v12, v3)
    v12 += v12
    v1 -= 1
    fmt.Println(v12, v3, v12)
    if (0 == 4294967295) {
        if (0.1 != v3) {
            v13 := []int{}
            fmt.Println(v3)
            _ = v13
        } else {
            if true {
                var v14 uint64 = 1
                fmt.Println(v3, v12, v3)
                _ = v14
            }
        }
        v15 := []int{127, -908, 567, 655}
        v16 := GenS1{F0: (-1 <= -1), F1: ((152 != 43) || false)}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        fmt.Println(v3, v3, v3)
        _ = fi17
    }
    v1 = 4294967295
    v12 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
