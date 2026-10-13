package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 uint64 = 995
    var v4 uint32 = 4294967295
    var v5 uint64 = (v3 / ((v3) | 1))
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: ((100 <= 42) || false), F1: 1}
        v2[fi6] = int(v5)
        if (!true) {
            fmt.Println(v3, v5)
            v1[(v4 / ((uint32(v5)) | 1))] = fi6
        }
        _ = fi6
        _ = v7
    }
    v3 -= 4294967295
    for fi8 := 0; fi8 < 13; fi8++ {
        if (((-2.25 / 0.1) / 1e10) > float32(0.1)) {
            for fi9 := 0; fi9 < 1; fi9++ {
                if (!true) {
                    fmt.Println(fi9)
                    var v10 int16 = 0
                    _ = v10
                }
                switch (fi9 * (fi8 ^ (fi8 + 737))) {
                case 656:
                    var v11 uint32 = (v4 % ((v4) | 1))
                    _ = v11
                case 128:
                    if (128 == 2) {
                        var v12 int16 = 32767
                        _ = v12
                    }
                }
                _ = fi9
            }
            fmt.Println(fi8, fi8, v5)
        } else {
            v13 := []int{}
            _ = v13
        }
        _ = fi8
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int16 = 1000
        v3 = v5
        _ = fi14
        _ = v15
    }
    v1[4294967295] = -520
    var v16 uint = 957
    v16 -= (v16 & v16)
    v17 := []int{127, 9223372036854775807}
    v18 := GenS2{F0: (!true)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
