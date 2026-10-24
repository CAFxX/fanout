package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int8 }
func genFn0() uint {
    v7 := []int{32767, 32768}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint64 = 0
        v9 -= (v9 ^ 1)
        _ = fi8
        _ = v9
    }
    _ = v7
    return 255
}

func main() {
    var v1 bool = false
    var v2 uint64 = 67
    v3 := make(map[string]uint32)
    var v4 int8 = -76
    v5 := make(map[int]uint32)
    v6 := []uint32{2}
    _ = len(v6)
    var v10 int32 = 42
    switch "hello" {
    case "":
        v2 = (v2 + v2)
    case "\n\t":
        for fi11 := 0; fi11 < 6; fi11++ {
            if true {
                _ = len(v6)
                fmt.Println(fi11)
                v4 -= (v4 * (v4 % (((-(v4))) | 1)))
            }
            _ = fi11
        }
    case "hello":
        var v12 int64 = -1
        _ = v12
    }
    v13 := GenS1{F0: "", F1: (-(v10))}
    var v14 bool = (!((42 < 32767) || v1))
    v6 = append(v6, 734)
    fmt.Println(v4)
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS2{F0: int8(fi15)}
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
