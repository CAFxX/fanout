package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 uint, p1 int16, p2 uint64) bool {
    p1 = (-(p1))
    var v4 uint32 = 42
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: -74, F1: (^(p2)), F2: (p1 - p1)}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return false
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 29
    for fi7 := 0; fi7 < 3; fi7++ {
        delete(v1, "hello")
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    v3 += v3
    v8 := []int{-1, 0, 42, 127}
    for fi9 := 0; fi9 < 7; fi9++ {
        switch "hello" {
        case "\n\t":
            var v10 uint = 0
            _ = v10
        }
        switch "a" {
        case "hello":
            fmt.Println(fi9, v3)
        case "\n\t":
            var v11 uint64 = 0
            _ = v11
        }
        v3 = v3
        _ = fi9
    }
    v3 += (v3 % ((v3) | 1))
    v12 := GenS1{F0: 0, F1: 18446744073709551615, F2: 1}
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            if (4294967295 < 42) {
                var v14 int64 = -1
                _ = v14
            }
            _ = len(v8)
            _ = fi13
        }
        v3 = 2
    } else {
        v3 -= (v3 % (((v3 | (v3 / (((v3 / ((1) | 1))) | 1)))) | 1))
        v15 := GenS1{F0: 28, F1: 18446744073709551615, F2: 0}
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        v3 = 42
        var v17 uint16 = 2
        _ = fi16
        _ = v17
    }
    var v18 string = ("hello" + "foo bar")
    fmt.Println(v3)
    v3 -= (v3 << ((2) & 63))
    v19 := GenS2{F0: 0, F1: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
