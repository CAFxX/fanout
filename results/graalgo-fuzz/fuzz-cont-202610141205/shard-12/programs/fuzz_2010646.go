package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 bool }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint32, p1 int16, p2 int32) uint32 {
    switch "a" {
    case "foo bar":
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := GenS1{F0: (p0 % ((p0) | 1)), F1: -69, F2: (!(false && false))}
            _ = fi5
            _ = v6
        }
    case "\n\t":
        v7 := []int{}
        _ = v7
    case "a":
        v8 := GenS1{F0: (p0 | (p0 + p0)), F1: -1, F2: (255 == 732)}
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{2147483647}
        p0 = 1
        _ = fi9
        _ = v10
    }
    p2 = (p2 >> ((793) & 63))
    _ = p0
    _ = p1
    _ = p2
    return (p0 + p0)
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]int)
    var v3 string = ""
    v4 := make(map[int]bool)
    _ = len(v2)
    fmt.Println(v1)
    fmt.Println(v3, v1)
    var v11 uint64 = 737
    var v12 string = ""
    var v13 int8 = -128
    v14 := GenS1{F0: 0, F1: (v13 << ((731) & 63)), F2: (!(868 <= 1))}
    v15 := GenS1{F0: 576, F1: v13, F2: ((1.5 + 0.1) != ((1.5 + 0.1) / 1e10))}
    fmt.Println(v11)
    v3 = v12
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v12 = ("foo bar" + v12)
            v18 := GenS1{F0: 536, F1: v13, F2: false}
            v19 := GenS2{F0: 2, F1: 0}
            _ = fi17
            _ = v18
            _ = v19
        }
        v13 += v13
        v2[genFn0(16, -896, -2147483648)] = fi16
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
