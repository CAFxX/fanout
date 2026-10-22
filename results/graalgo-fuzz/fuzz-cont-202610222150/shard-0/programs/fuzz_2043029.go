package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint8 }
type GenS3 struct { F0 string; F1 string; F2 int16 }
func genFn0(p0 uint, p1 uint16) uint64 {
    p1 = p1
    var v7 int = 2147483647
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 0
    var v2 int = -665
    var v3 uint8 = uint8(v1)
    var v4 float32 = ((3.14159 + 0.0) / 1.0)
    v5 := make(map[int]bool)
    v6 := []uint32{2, 4294967295}
    v3 = (^(v3))
    delete(v5, -1)
    v4 += (v4 * ((v4 - (v4 / 1.0)) * ((1.5 + (1.5 - v4)) - v4)))
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int32 = 297
        _ = fi8
        _ = v9
    }
    v6 = append(v6, 110)
    v10 := GenS1{F0: 1, F1: "hello", F2: 0}
    v11 := GenS1{F0: 170, F1: "a", F2: 1}
    _ = len(v5)
    var v12 int64 = 9223372036854775807
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int32 = -783
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: -935, F1: false, F2: (v3 % (((v3 & uint8(v1))) | 1))}
        v17 := GenS3{F0: "foo bar", F1: "a", F2: 42}
        _ = fi13
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        _ = len(v6)
        fmt.Println(v2, v1, v2)
        if ((!(v4 >= v4)) && false) {
            v1 = 1
        } else {
            var v19 uint = 42
            _ = v19
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
