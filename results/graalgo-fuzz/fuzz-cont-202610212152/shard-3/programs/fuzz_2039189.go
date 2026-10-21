package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 bool, p2 int64) int16 {
    if (((false || (-263 == 1)) || false) || p1) {
        for fi4 := 0; fi4 < 19; fi4++ {
            var v5 bool = true
            _ = fi4
            _ = v5
        }
        var v6 uint8 = 252
        _ = v6
    }
    v7 := GenS1{F0: 1}
    v8 := []int{(p0 | (p0 >> ((1) & 63))), (p0 >> ((42) & 63)), p0, p0}
    if len(v8) > 0 {
        v8[0] = int(p2)
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 288
}

func main() {
    var v1 uint8 = 127
    v2 := []uint32{1, 608, 726}
    var v3 int32 = 567
    fmt.Println(v3, v1)
    v2 = append(v2, 1)
    v9 := GenS2{F0: 2147483647}
    var v10 uint8 = (^(v1))
    for fi11 := 0; fi11 < 18; fi11++ {
        switch "hello" {
        case "a":
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint32 = 2147483647
                _ = fi12
                _ = v13
            }
        case "":
            var v14 float64 = 0.1
            _ = v14
        case "foo bar":
            var v15 int8 = -128
            _ = v15
        }
        v10 += (v10 << ((223) & 63))
        _ = fi11
    }
    if false {
        v16 := GenS2{F0: -910}
        v1 += v10
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v16
    }
    v3 -= -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "\n\t" {
    case "a":
        _ = len(v2)
    case "hello":
        v10 += (v1 << ((354) & 63))
    }
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
