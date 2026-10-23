package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    var v1 float64 = 0.1
    var v2 uint16 = 322
    v3 := make(map[uint32]int)
    var v4 string = "foo bar"
    var v5 int8 = -1
    var v6 string = "\n\t"
    v2 = v2
    var v7 bool = (!true)
    v3[2147483647] = -922
    v8 := []int{0, -1, 416, 2147483647}
    if false {
        if true {
            v9 := GenS1{F0: v4, F1: 1}
            _ = v9
        }
    } else {
        v10 := GenS2{F0: 171, F1: 42}
        for fi11 := 0; fi11 < 19; fi11++ {
            fmt.Println(v1)
            if (937 < 0) {
                v8 = append(v8, fi11)
            }
            _ = fi11
        }
        var v12 uint32 = 0
        _ = v10
        _ = v12
    }
    switch v4 {
    case "foo bar":
        v8 = append(v8, 32768)
    default:
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint32 = uint32(v2)
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 uint64 = 0
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    }
    v6 = "hello"
    fmt.Println(v4, v5, v6)
    v2 -= (v2 % ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
