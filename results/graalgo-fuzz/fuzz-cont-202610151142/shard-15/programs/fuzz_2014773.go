package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 bool = true
    var v2 uint = 0
    var v3 int8 = 61
    v4 := []int{127}
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 4; fi6++ {
        delete(v5, 2147483648)
        _ = fi6
    }
    var v7 uint64 = 942
    fmt.Println(v7, v2, v7)
    if v1 {
        v8 := GenS1{F0: -32768, F1: 394, F2: 96}
        var v9 int16 = -32768
        fmt.Println(v2, v7, v3)
        _ = v8
        _ = v9
    } else {
        v2 = uint(v7)
        var v10 uint64 = (v7 << (((v2 + (^(v2)))) & 63))
        _ = v10
    }
    var v11 uint = 521
    if v1 {
        var v12 float64 = (-2.25 - 1.5)
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: 1}
            _ = len(v4)
            _ = fi13
            _ = v14
        }
        fmt.Println(v12)
        _ = v12
    } else {
        v1 = (!("\n\t" <= ("\n\t" + "foo bar")))
        if v1 {
            v15 := v5[0]
            for fi16 := 0; fi16 < 6; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    delete(v5, (v15 | (v15 / ((v15) | 1))))
                    _ = fi17
                }
                v2 = 614
                fmt.Println(v7, v11)
                _ = fi16
            }
            v11 = v11
            _ = v15
        }
    }
    _ = len(v5)
    v7 = (v7 & 788)
    v1 = v1
    var v18 bool = (42 != 1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
