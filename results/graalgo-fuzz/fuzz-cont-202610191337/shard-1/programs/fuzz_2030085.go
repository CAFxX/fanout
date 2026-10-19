package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int, p1 uint) int8 {
    p0 -= p0
    for fi4 := 0; fi4 < 3; fi4++ {
        if true {
            if true {
                v5 := GenS1{F0: 127}
                var v6 int32 = 1
                fmt.Println(p1, v6)
                _ = v5
                _ = v6
            } else {
                var v7 uint64 = 2
                v8 := []int{}
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: -1}
            var v10 int = (p0 % ((699) | 1))
            _ = v10
            _ = v9
        }
        _ = fi4
    }
    fmt.Println(p0, p1)
    p0 = 284
    p1 = (p1 % ((398) | 1))
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int8 = int8(v1)
    v3 := make(map[uint32]bool)
    v2 = v2
    v11 := GenS2{F0: 42, F1: 266, F2: 713}
    v2 = 127
    v12 := v3[(v1 / ((1) | 1))]
    if false {
        var v13 int64 = 912
        _ = v13
    } else {
        if v12 {
            _ = len(v3)
        }
        v2 = v2
    }
    v2 = v2
    switch "a" {
    case "a":
        v14 := GenS1{F0: 127}
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := []int{2147483647}
            fmt.Println(v1, v12)
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS3{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
}
