package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 int64 = -1
    v2 := []int{int(v1), 438}
    var v3 uint8 = 1
    v4 := make(map[uint32]int)
    var v5 int64 = 1
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 bool = true
        if (!true) {
            for fi8 := 0; fi8 < 18; fi8++ {
                switch fi8 {
                case 4294967295:
                    v2 = append(v2, (fi6 / ((4294967295) | 1)))
                case 919:
                    fmt.Println(fi6)
                default:
                    fmt.Println(v1, fi6, v1)
                }
                v9 := GenS1{F0: v3, F1: fi8}
                _ = len(v4)
                _ = fi8
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    v5 = v1
    if (false || true) {
        var v10 uint16 = 42
        _ = v10
    } else {
        v2 = append(v2, -2147483648)
    }
    v11 := GenS2{F0: 189, F1: (2 == 0)}
    v12 := GenS2{F0: 31, F1: false}
    var v13 string = "\n\t"
    delete(v4, 1)
    for fi14 := 0; fi14 < 10; fi14++ {
        if false {
            var v15 uint = 433
            _ = v15
        } else {
            var v16 float32 = 1e10
            _ = v16
        }
        v3 = 128
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
