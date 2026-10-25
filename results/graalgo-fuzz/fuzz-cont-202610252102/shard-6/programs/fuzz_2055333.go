package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 int = 4294967295
    for fi5 := 0; fi5 < 17; fi5++ {
        switch (fi5 ^ fi5) {
        case 0:
            if false {
                var v6 string = "a"
                v7 := []int{}
                _ = v6
                _ = v7
            }
        case -712:
            v4 = v4
        }
        v8 := GenS1{F0: 9223372036854775807, F1: 83, F2: 1}
        _ = fi5
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 bool = (286 != 1)
        _ = fi9
        _ = v10
    }
    v4 = -941
    if true {
        v11 := []int{(v4 & (v4 + v4)), v4}
        v4 += v4
        _ = v11
    }
    var v12 int16 = 1000
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            v4 = -181
            v4 -= 42
            _ = fi14
        }
        v15 := []int{v4, v4, fi13, 128}
        _ = fi13
        _ = v15
    }
    v4 -= (v4 % ((v4) | 1))
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint16 = 32767
        _ = fi16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
