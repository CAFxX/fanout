package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 uint8) string {
    p0 -= (^(p0))
    v4 := GenS1{F0: true, F1: 42, F2: "\n\t"}
    _ = p0
    _ = v4
    return "foo bar"
}

func main() {
    v1 := []int{127, 0}
    var v2 float64 = (0.0 / 1.0)
    v3 := make(map[uint32]bool)
    var v5 uint = 65
    v5 = 4294967295
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(v5, v5)
        if (v5 >= v5) {
            v7 := GenS1{F0: (-8 > 127), F1: -14, F2: "a"}
            _ = len(v3)
            for fi8 := 0; fi8 < 6; fi8++ {
                switch fi8 {
                case 4294967295:
                    _ = len(v1)
                    v5 += (v5 | (v5 ^ v5))
                    v5 = (v5 * 1)
                    v2 = 0.1
                    fmt.Println(fi8, v2, fi8)
                case 353:
                    v9 := GenS2{F0: 42, F1: "foo bar", F2: 266}
                    _ = v9
                case 235:
                    switch (fi6 & fi8) {
                    case 32767:
                        v5 = 42
                    }
                default:
                    _ = len(v1)
                }
                v1 = append(v1, -2147483648)
                var v10 uint = 0
                v10 = (v5 % ((v10) | 1))
                _ = fi8
                _ = v10
            }
            _ = v7
        } else {
            var v11 uint64 = 969
            v11 += uint64(v5)
            _ = v11
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS3{F0: fi13, F1: 1}
            v5 = v5
            _ = fi13
            _ = v14
        }
        fmt.Println(fi12, v2)
        fmt.Println(fi12, v2)
        _ = fi12
    }
    var v15 uint32 = 1
    _ = len(v1)
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(fi16)
        _ = fi16
    }
    var v17 float64 = v2
    v3[(^(v15))] = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v5
}
