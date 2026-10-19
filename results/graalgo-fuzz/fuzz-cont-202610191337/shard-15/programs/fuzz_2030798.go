package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 uint64; F2 int16 }
func genFn0(p0 uint16, p1 uint8) bool {
    for fi4 := 0; fi4 < 1; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := GenS1{F0: true, F1: -722}
            fmt.Println(p0)
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v7 := GenS2{F0: 4294967295, F1: true, F2: 1}
    var v8 int = -749
    v8 = v8
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (32767 == 1)
}

func main() {
    var v1 float64 = 1.5
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    v1 = 1e10
    v1 -= 1.5
    var v9 int16 = 32767
    switch "\n\t" {
    case "a":
        var v10 int = 42
        _ = v10
    case "hello":
        var v11 uint16 = 84
        _ = v11
    default:
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 18; fi13++ {
                v9 += v9
                delete(v3, 4294967295)
                switch fi13 {
                case 914:
                    var v14 int16 = 42
                    _ = v14
                case 2147483647:
                    _ = len(v3)
                default:
                    v1 += v1
                    v1 -= float64(1.5)
                    v1 = ((v1 / v1) / 3.14159)
                    v1 += (-2.25 / (((((v1 + float64(0.0)) + v1) - v1) * 0.1) + v1))
                }
                _ = fi13
            }
            var v15 int64 = 0
            _ = fi12
            _ = v15
        }
    }
    switch "a" {
    case "":
        v1 -= 0.1
    default:
        if false {
            for fi16 := 0; fi16 < 12; fi16++ {
                if len(v2) > 0 {
                    v2[0] = fi16
                }
                var v17 uint16 = 0
                _ = fi16
                _ = v17
            }
            var v18 float64 = 1.5
            _ = v18
        }
    }
    var v19 uint16 = 255
    switch 4294967295 {
    case -1:
        v20 := GenS3{F0: -51, F1: 2, F2: (v9 * (v9 + (-(v9))))}
        _ = v20
    default:
        v9 = (v9 >> ((0) & 63))
    }
    fmt.Println(v19, v19, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
